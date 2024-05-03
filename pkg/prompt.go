package pkg

import (
	"github.com/AlecAivazis/survey/v2"
)

type Parameters struct {
	Version string // OpenSearch Version (1.x, 2.x)
}

var basicQuestions = []*survey.Question{
	{
		Name: "version",
		Prompt: &survey.Select{
			Message: "Which OpenSearch version do you want to use?",
			Options: []string{"2.13.0", "1.3.16"},
			Default: "2.13.0",
		},
	},
}

func GetPromptValues() Parameters {
	answers := Parameters{}
	err := survey.Ask(basicQuestions, &answers)
	if err != nil {
		panic(err)
	}
	return answers
}
