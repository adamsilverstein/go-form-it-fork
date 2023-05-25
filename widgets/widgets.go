// This package contains the base logic for the creation and rendering of field widgets. Base widgets are defined for most input fields,
// both in classic and Bootstrap3 style; custom widgets can be defined and associated to a field, provided that they implement the
// WidgetInterface interface.
package widgets

import (
	"bytes"
	"html/template"

	formcommon "github.com/adamsilverstein/go-form-it-fork/formcommon"
)

// Simple widget object that gets executed at render time.
type Widget struct {
	template *template.Template
}

// WidgetInterface defines the requirements for custom widgets.
type WidgetInterface interface {
	Render(data interface{}) string
}

// Render executes the internal template and returns the result as a template.HTML object.
func (w *Widget) Render(data interface{}) string {
	var s string
	buf := bytes.NewBufferString(s)
	w.template.ExecuteTemplate(buf, "main", data)
	return buf.String()
}

// BaseWidget creates a Widget based on style and inpuType parameters, both defined in the common package.
func BaseWidget(style, inputType string) *Widget {
	urls := []string{formcommon.CreateUrl("static/templates/forms/%s/generic.tmpl", style)}
	switch inputType {
	case formcommon.BUTTON:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/button.html", style))
	case formcommon.CHECKBOX:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/options/checkbox.html", style))
	case formcommon.TEXTAREA:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/text/textareainput.html", style))
	case formcommon.SELECT:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/options/select.html", style))
	case formcommon.PASSWORD:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/text/passwordinput.html", style))
	case formcommon.RADIO:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/options/radiobutton.html", style))
	case formcommon.TEXT:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/text/textinput.html", style))
	case formcommon.RANGE:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/number/range.html", style))
	case formcommon.NUMBER:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/number/number.html", style))
	case formcommon.RESET:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/button.html", style))
	case formcommon.SUBMIT:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/button.html", style))
	case formcommon.DATE:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/datetime/date.html", style))
	case formcommon.DATETIME:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/datetime/datetime.html", style))
	case formcommon.TIME:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/datetime/time.html", style))
	case formcommon.DATETIME_LOCAL:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/datetime/datetime.html", style))
	case formcommon.STATIC:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/static.html", style))
	case formcommon.SEARCH,
		formcommon.TEL,
		formcommon.URL,
		formcommon.WEEK,
		formcommon.COLOR,
		formcommon.EMAIL,
		formcommon.FILE,
		formcommon.HIDDEN,
		formcommon.IMAGE,
		formcommon.MONTH:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/input.html", style))
	default:
		urls = append(urls, formcommon.CreateUrl("static/templates/forms/%s/input.html", style))
	}
	templ := template.Must(template.ParseFiles(urls...))
	return &Widget{templ}
}
