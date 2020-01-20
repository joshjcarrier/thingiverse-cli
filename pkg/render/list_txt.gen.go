/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package render

// list_txtTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func list_txtTemplate() string {
	var tmpl = "\n" +
		"{{- if $.Options.LongFormat -}}\n" +
		"\n" +
		"{{- printf \"%-7s|%-41s|%-40s\" \"ID\" \"URL\" \"Name\" }}\n" +
		"{{ range .Data -}}\n" +
		"{{- printf \"%-7d\" .Id -}}|{{- .PublicUrl -}}|{{- printf \"%-40s\" .Name }}\n" +
		"{{ end -}}\n" +
		"\n" +
		"{{- else -}}\n" +
		"\n" +
		"{{- range .Data -}}\n" +
		"{{- .Id }}\n" +
		"{{ end -}}\n" +
		"\n" +
		"{{ end -}}\n" +
		""
	return tmpl
}
