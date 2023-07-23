# GOBARD

Unofficial Golang API for Google BARD ChatBOT.

## OBTAIN A BARD COOKIE

1. Visit https://bard.google.com/ (login with your account).
2. F12 for console.
3. Session: Application → Cookies → `__Secure-1PSID` cookie value.

> ATTENTION: Do not share your auth cookie.

## HOWTO

- Create a GOBARD object.
- `.Ask("something")`
- `.GetAnswer()`
- Did not like this answer ? `.Next()` and `.GetAnswer()` (or `.NextAnswer()`)
- Want to go back to previous answer ? `.Prev()` and `.GetAnswer()` (or `.PrevAnswer()`)
- Next question (`.Ask()`) will keep Chat reference until `.Reset()` is called.
- `.Ask()` will use the "current answer" as a reference for the question.
- Have Fun!

## QUICK EXAMPLE

```go
package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/aquasecurity/gobard"
	"github.com/charmbracelet/glamour"
)

var render *glamour.TermRenderer

func init() {
	render, _ = glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(80),
	)
}

func main() {
	cookie := os.Getenv("BARD_COOKIE")
	if cookie == "" {
		fmt.Fprintf(os.Stderr, "BARD_COOKIE is not set\n")
		os.Exit(1)
	}

	gobard.New(cookie)

	bard01 := gobard.New(cookie)

	prompt := "Act as a simple calculator and calculate 2 + 2. Give the result only, no more words."

	err := bard01.Ask(url.QueryEscape(prompt))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for i := 0; i < bard01.GetNumOfAnswers(); i++ {
		answerMD, _ := render.Render(bard01.GetAnswer())
		fmt.Printf("%s\n", answerMD)
		fmt.Printf("----\n")
		bard01.Next()
	}

	bard01.Next() // will continue the conversation using the first answer as a base

	err = bard01.Ask(url.QueryEscape("What if I add 3 to the result ?"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for i := 0; i < bard01.GetNumOfAnswers(); i++ {
		answerMD, _ := render.Render(bard01.GetAnswer())
		fmt.Printf("%s\n", answerMD)
		fmt.Printf("----\n")
		bard01.Next()
	}
}
```
