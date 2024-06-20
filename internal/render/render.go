package render

import (
	"html/template"
	"invest-tracker/pkg/config"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) error {
	htmlTemplatePath, err := getHtmlTemplatePath()
	if err != nil {
		return err
	}

	parsedTemplate, err := template.ParseFiles(filepath.Join(htmlTemplatePath, tmpl))
	if err != nil {
		return err
	}
	return parsedTemplate.Execute(w, nil)
}

func getHtmlTemplatePath() (string, error) {
	cfg, err := config.Read()
	if err != nil {
		return "", err
	}
	return cfg.HtmlTemplatePath, nil
}
