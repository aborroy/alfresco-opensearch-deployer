package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/aborroy/alf-opensearch/pkg"
	"github.com/spf13/cobra"
)

// Default values
const OutputRootPath string = "output"
const TemplateRootPath string = "templates"

// Parameters mapping
var interactive bool = true
var outputDirectory string
var version string

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create assets to deploy Alfresco and OpenSearch",
	Run: func(cmd *cobra.Command, args []string) {

		var outputRoot = OutputRootPath
		if outputDirectory != "" {
			outputRoot = outputDirectory
		}

		var values pkg.Parameters

		if !interactive {
			values = pkg.Parameters{
				Version: version,
			}
		} else {
			values = pkg.GetPromptValues()
		}

		templateList, err := pkg.EmbedWalk("templates")
		if err != nil {
			panic(err)
		}

		os.MkdirAll(outputRoot, os.ModePerm)

		for _, t := range templateList {
			outputFile := outputRoot + "/" + t
			position := strings.Index(outputFile, "/templates/")
			outputFile = outputFile[position+len("/templates/"):]
			os.MkdirAll(filepath.Dir(outputRoot+"/"+outputFile), os.ModePerm)
			f, _ := os.Create(outputRoot + "/" + outputFile)
			name := filepath.Base(t)
			tpl, _ := template.New(name).ParseFS(pkg.TemplateFs, t)
			err = tpl.ExecuteTemplate(f, name, values)
			if err != nil {
				panic(err)
			}
			err = pkg.VerifyOutputFile(f.Name())
			if err != nil {
				panic(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(&interactive, "interactive", "i", true, "Input values replying to command line prompts instead of using command line parameters")
	createCmd.Flags().StringVarP(&version, "version", "v", "", "Version of OpenSearch to be deployed (2.13.0 or 1.3.16)")
	createCmd.Flags().StringVarP(&outputDirectory, "output", "o", "", "Local Directory to write produced assets, 'output' by default")
}
