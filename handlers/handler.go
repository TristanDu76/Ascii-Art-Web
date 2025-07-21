package handlers

import (
	"ascii-art-web/functions"
	"ascii-art-web/utils"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		utils.ServeError(w, r, http.StatusInternalServerError, "500.html")
		return
	}

	data := struct {
		Result template.HTML
		Raw    string
		Style  []string
		Color  []string
	}{
		Result: "",
		Raw:    "",
		Style:  []string{"Standard", "Shadow", "Thinkertoy"},
		Color:  []string{"White", "Red", "Green", "Blue", "Yellow", "Purple", "Cyan"},
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			utils.ServeError(w, r, http.StatusBadRequest, "400.html")
			return
		}

		text := r.FormValue("text")
		style := strings.ToLower(r.FormValue("Style"))
		color := strings.ToLower(r.FormValue("Color"))

		if text == "" || style == "" {
			utils.ServeError(w, r, http.StatusBadRequest, "400.html")
			return
		}

		// Call BuildResult and handle error
		result, err := functions.BuildResult(text, style, color)
		if err != nil {
			utils.ServeError(w, r, http.StatusBadRequest, "400.html")
			return
		}

		rawOutput := result.String()
		coloredOutput := fmt.Sprintf(`<span class="ascii-%s">%s</span>`, color, rawOutput)

		data.Result = template.HTML(coloredOutput)
		data.Raw = rawOutput // For download
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		utils.ServeError(w, r, http.StatusInternalServerError, "500.html")
	}
}

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		utils.ServeError(w, r, http.StatusBadRequest, "400.html")

		return
	}

	ascii := r.FormValue("ascii")
	if ascii == "" {
		utils.ServeError(w, r, http.StatusBadRequest, "400.html")

		return
	}
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii.art.txt")
	w.Write([]byte(ascii))

}
