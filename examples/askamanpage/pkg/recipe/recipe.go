package recipe

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Question struct {
	Title    string `yaml:"title"`
	Question string `yaml:"question"`
}

type Recipe struct {
	Manpages  []string   `yaml:"manpages"`
	Questions []Question `yaml:"questions"`
}

func (r Recipe) GetManPages() []string {
	return r.Manpages
}

func (r Recipe) GetTitlesAndQuestions() map[string]string {
	titlesAndQuestions := make(map[string]string)

	for _, q := range r.Questions {
		titlesAndQuestions[q.Title] = q.Question
	}

	return titlesAndQuestions
}

func ReadYAML(filePath string) (Recipe, error) {
	yamlData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to read the YAML file: %v", err)
	}

	var data Recipe
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return Recipe{}, fmt.Errorf("failed to unmarshal YAML: %v", err)
	}

	return data, nil
}
