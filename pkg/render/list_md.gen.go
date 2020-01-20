/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package render

// list_mdTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func list_mdTemplate() string {
	var tmpl = "\n" +
		"{{- range .Data }}\n" +
		"{{ .Name }}\n" +
		"{{ end -}}\n" +
		""
	return tmpl
}
