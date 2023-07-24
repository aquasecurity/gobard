package man7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	BaseURL     = "https://man7.org/linux/man-pages/"
	AllManPages = "dir_all_alphabetic.html"
)

// ManPage

// ManPage describes metadata for a man7.org manual page
type ManPage struct {
	section     int
	name        string
	description string
	url         string
}

// NewManPage creates a new ManPage
func NewManPage(section int, name, url, description string) *ManPage {
	return &ManPage{
		section:     section,
		name:        name,
		description: description,
		url:         url,
	}
}

// String returns a string representation of a ManPage
func (man *ManPage) String() string {
	return fmt.Sprintf("%v: %v(%v) - %v", man.url, man.name, man.section, man.description)
}

// GetSection returns the URL of the ManPage
func (man *ManPage) GetSection() int {
	return man.section
}

// GetName returns the name of the ManPage
func (man *ManPage) GetName() string {
	return man.name
}

// GetDescription returns the description of the ManPage
func (man *ManPage) GetDescription() string {
	return man.description
}

// GetURL returns the URL of the ManPage
func (man *ManPage) GetURL() string {
	return man.url
}

// GetFullURL returns the full URL of the ManPage
func (man *ManPage) GetFULLURL() string {
	return BaseURL + strings.TrimLeft(man.url, "./")
}

// ManPages

// ManPages is a collection of ManPage
type ManPages struct {
	manPages map[string]*ManPage
}

// NewManPages creates a new ManPages
func NewManPages() *ManPages {
	return &ManPages{
		manPages: make(map[string]*ManPage),
	}
}

// Add adds a ManPage to the ManPages
func (manPgs *ManPages) Add(man *ManPage) {
	manPgs.manPages[man.GetName()] = man
}

// Get returns a ManPage by name
func (manPgs *ManPages) Get(name string) *ManPage {
	return manPgs.manPages[name]
}

// CreateManPages creates man pages from the man7.org website
func CreateManPages() (*ManPages, error) {
	manPages := NewManPages()

	client := resty.New()
	client.SetBaseURL(BaseURL)

	resp, err := client.R().Get(AllManPages)
	if err != nil {
		return manPages, fmt.Errorf("failed to get the URL: %v", err)
	}
	if resp.StatusCode() != 200 {
		return manPages, fmt.Errorf("status code: %v", resp.StatusCode())
	}

	// e.g.: nbsp; &nbsp; <a href="./man8/yum-copr.8.html">yum-copr(8)</a> - YUM copr Plugin
	pattern := regexp.MustCompile(".*<a href=\"(.*)\">(.*)\\((\\d)\\)</a> - (.*)")

	lines := pattern.FindAllStringSubmatch(string(resp.Body()), -1)
	for _, line := range lines {
		section, _ := strconv.Atoi(line[3])
		manPages.Add(
			NewManPage(
				section, // section =D
				line[2], // name
				line[1], // url
				line[4], // description
			),
		)
	}

	return manPages, nil
}
