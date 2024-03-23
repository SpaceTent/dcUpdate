package web

import (
	"fmt"
	l "log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"dcupdate/app/environment"
	"dcupdate/plugins"
)

func findAndParseTemplates(templateDirectory string, funcMap template.FuncMap) (*template.Template, error) {

	cleanRoot := filepath.Clean(templateDirectory)
	pfx := len(cleanRoot) + 1
	AllTemplates := template.New("")

	// l.With("templateDirectory", templateDirectory).With("cleanRoot", cleanRoot).Debug("findAndParseTemplates")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if !info.IsDir() && strings.HasSuffix(path, ".gohtml") {
			if e1 != nil {
				return e1
			}

			b, e2 := os.ReadFile(path)
			if e2 != nil {
				return e2
			}

			name := path[pfx:]
			// l.With("name", name).Debug("findAndParseTemplates")
			t := AllTemplates.New(name).Funcs(funcMap)
			_, e2 = t.Parse(string(b))
			if e2 != nil {
				l.With("error", e2).With("template", name).Error("Error parsing template")
				return e2
			}
		}

		return nil
	})

	return AllTemplates, err
}

func Render(w http.ResponseWriter, r *http.Request, pageData map[string]any, templateFile string) {

	templateDirectory := environment.GetEnvString("TEMPLATE_PREFIX", "") + "templates"

	// does this Exist?
	if _, err := os.Stat(templateDirectory); os.IsNotExist(err) {
		l.With("error", err).With("templateDirectory", templateDirectory).Error("Template Directory does not exist")
		return
	}

	templates, err := findAndParseTemplates(templateDirectory, plugins.Plugin)
	if err != nil {
		l.With("error", err).Error("Error loading templates")
		_ = WriteJSON(w, 500, err.Error())
		return
	}

	if pageData == nil {
		pageData = make(map[string]any)
	}

	pageData["RAND"] = "NoCache"
	if environment.GetEnvString("CACHE", "on") == "off" {
		if environment.VERSION == "Development" {
			pageData["RAND"] = time.Now().UnixNano()
		}
	}
	// pageData["RAND"] = "CacheMe"
	pageData["VERSION"] = environment.VERSION

	DebugInfo := ""
	for key, value := range pageData { // range over map
		DebugInfo = DebugInfo + fmt.Sprintf(" K:[F-MAGENTA]%s[F-NORMAL] V:[F-MAGENTA]%v[F-NORMAL]", key, value)
	}

	l.With("template", fmt.Sprintf("%s", templateFile), "URL", r.URL.String()).Debug("Render")

	err = templates.ExecuteTemplate(w, templateFile, pageData)
	if err != nil {
		l.With("error", err).With("templatefile", templateFile).Error("Error loading templates")
		_ = WriteJSON(w, 500, err.Error())
		return
	}
}
