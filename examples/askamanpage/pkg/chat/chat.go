package chat

import (
	"os"
	"strings"

	"github.com/aquasecurity/gobard"
	"github.com/aquasecurity/askamanpage/pkg/man7"
)

type Chat struct {
	manPage   *man7.ManPage
	questions map[string]string
	bard      *gobard.Bard
}

func New(manPage *man7.ManPage, questions map[string]string) (*Chat, error) {
	bard := gobard.New(os.Getenv("BARD_COOKIE")) // at this point we know it's set

	return &Chat{
		manPage:   manPage,
		questions: questions,
		bard:      bard,
	}, nil
}

func (c *Chat) AskQuestions() (map[string]string, error) {
	r := make(map[string]string)

	for title, question := range c.questions {
		err := c.bard.Ask(c.normQuestion(question))
		if err != nil {
			return r, err
		}
		r[title] = c.bard.GetAnswer()
	}

	return r, nil
}

func (c *Chat) normQuestion(question string) string {
	question = strings.Replace(question, "%URL%", c.manPage.GetFULLURL(), 1)
	question = strings.Replace(question, "%MANPAGE%", c.manPage.GetName(), 1)

	return question
}
