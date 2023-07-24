package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aquasecurity/askamanpage/pkg/chat"
	"github.com/aquasecurity/askamanpage/pkg/man7"
	"github.com/aquasecurity/askamanpage/pkg/recipe"
	"github.com/urfave/cli/v2"
)

const (
	outDirectory = "./out"
)

func main() {
	app := &cli.App{
		Name:  "AskAManPage",
		Usage: "Create content from man7.org manual pages",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "Path to the YAML file",
				Required: true,
			},
		},
		Action: AskAManPage,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func AskAManPage(c *cli.Context) error {
	// Sanity Check for BARD_COOKIE
	if os.Getenv("BARD_COOKIE") == "" {
		fmt.Fprintf(os.Stderr, "BARD_COOKIE is not set\n")
		os.Exit(1)
	}

	// Read the recipe YAML file
	recipe, err := recipe.ReadYAML(c.String("file"))
	if err != nil {
		return fmt.Errorf("failed to read the YAML file: %v", err)
	}

	// Create man7.org manual pages metadata
	manPages, err := man7.CreateManPages()
	if err != nil {
		return fmt.Errorf("failed to create man pages: %v", err)
	}

	// TODO: thread pool for parallel processing

	for _, manPageWanted := range recipe.GetManPages() { // for each man page I want to ask questions...
		// Create a Bard Chat to each man page
		manPageChat, err := chat.New(manPages.Get(manPageWanted), recipe.GetTitlesAndQuestions())
		if err != nil {
			return fmt.Errorf("failed to create chat: %v", err)
		}
		// Ask all questions to each man page, one by one
		manPageAnswers, err := manPageChat.AskQuestions()
		if err != nil {
			return fmt.Errorf("failed to ask questions: %v", err)
		}
		// Append the answers to a file
		for title, answer := range manPageAnswers {
			writeFile(manPageWanted, title, answer)
		}
		break // only a single entry for now
	}

	return nil
}

func writeFile(manpage, title, content string) error {
	filePath := outDirectory + "/" + manpage + ".md"

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		_, err = file.WriteString(fmt.Sprintf("# Event: %s\n\n", manpage))
		if err != nil {
			return fmt.Errorf("failed to write header to file: %v", err)
		}
		file.Close()
	} else if err != nil {
		return fmt.Errorf("failed to check file: %v", err)
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	what := fmt.Sprintf("## %s\n%s", title, content)
	_, err = file.WriteString(what)
	if err != nil {
		return fmt.Errorf("failed to append content to file: %v", err)
	}

	return nil
}
