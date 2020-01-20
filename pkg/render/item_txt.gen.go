/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package render

// item_txtTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func item_txtTemplate() string {
	var tmpl = "{{- with .Data -}}\n" +
		"Name: {{ .Name }}\n" +
		"URL: {{ .PublicUrl }}\n" +
		"By: {{ .Creator.Name }} \n" +
		"Created: {{ .Added }}\n" +
		"Summary: {{ .Description }}\n" +
		"{{ end -}}"
	return tmpl
}
