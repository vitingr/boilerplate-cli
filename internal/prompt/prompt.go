package prompt

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/vitingr/boilerplate-cli/internal/config"
)

func Run(projectName string) (*config.ScaffoldConfig, error) {
	cfg := &config.ScaffoldConfig{}

	if projectName != "" {
		cfg.ProjectName = projectName
	} else {
		if err := survey.AskOne(&survey.Input{
			Message: "Project name:",
			Default: "my-project",
		}, &cfg.ProjectName, survey.WithValidator(survey.Required)); err != nil {
			return nil, err
		}
	}

	languages := sortedKeys(config.Catalog)
	if err := survey.AskOne(&survey.Select{
		Message: "Language:",
		Options: languages,
	}, &cfg.Language); err != nil {
		return nil, err
	}

	frameworks := sortedKeys(config.Catalog[cfg.Language])
	if err := survey.AskOne(&survey.Select{
		Message: "Framework:",
		Options: frameworks,
	}, &cfg.Framework); err != nil {
		return nil, err
	}

	templates := config.Catalog[cfg.Language][cfg.Framework]
	if err := survey.AskOne(&survey.Select{
		Message: "Architecture / template:",
		Options: templates,
	}, &cfg.Template); err != nil {
		return nil, err
	}

	cwd, _ := os.Getwd()
	defaultOut := filepath.Join(cwd, cfg.ProjectName)
	if err := survey.AskOne(&survey.Input{
		Message: "Output directory:",
		Default: defaultOut,
	}, &cfg.OutputDir); err != nil {
		return nil, err
	}

	if err := survey.AskOne(&survey.Confirm{
		Message: "Initialize git repository?",
		Default: true,
	}, &cfg.InitGit); err != nil {
		return nil, err
	}

	depLabel := dependencyLabel(cfg.Language)
	if depLabel != "" {
		if err := survey.AskOne(&survey.Confirm{
			Message: fmt.Sprintf("Run %s after scaffold?", depLabel),
			Default: true,
		}, &cfg.InstallDeps); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func dependencyLabel(lang string) string {
	switch lang {
	case "node":
		return "npm install"
	case "golang":
		return "go mod tidy"
	case "python":
		return "pip install -r requirements.txt"
	}
	return ""
}
