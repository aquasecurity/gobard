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

	err := bard01.Ask(url.QueryEscape("Act as a simple calculator and calculate 2 + 2. Give the result only, no more words."))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	for i := 0; i < bard01.GetNumOfAnswers(); i++ {
		answerMD, _ := render.Render(bard01.GetAnswer())
		fmt.Printf("%s\n", answerMD)
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
		bard01.Next()
	}
}
```

## NOTES

1. Each Bard object, from `gobard.new()`, will have its own context until `Reset()` is called:

```
Bard01: Act as a simple calculator and calculate 2 + 2. Give the result only, no more words.

  Sure, I can help you with that.                                             
                                                                              
  2 + 2 = 4                                                                   
                                                                              
  I hope this is helpful. Let me know if you have any other questions.        

Bard02: Act as a simple calculator and calculate 4 + 8. Give the result only, no more words.

  Sure, I can help you with that.                                             
                                                                              
  4 + 8 = 12                                                                  
                                                                              
  Is there anything else I can help you with?                                 

Bard01: What if I add 5 to the result ?

  If you add 5 to the result of 2 + 2, you get 4 + 5 = 9.                     
                                                                              
  Is there anything else I can help you with today?                           

Bard02: What if I add 5 to the result ?

  If you add 5 to the result of 4 + 8, you get 12 + 5 = 17.                   
                                                                              
  Is there anything else I can help you with?                   
```
