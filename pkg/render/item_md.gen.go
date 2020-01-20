/*
 * CODE GENERATED AUTOMATICALLY WITH
 *    github.com/wlbr/templify
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package render

// item_mdTemplate is a generated function returning the template as a string.
// That string should be parsed by the functions of the golang's template package.
func item_mdTemplate() string {
	var tmpl = "{{- with .Data -}}\n" +
		"# [{{ .Name }}]({{ .PublicUrl }})\n" +
		"\n" +
		"By: [{{ .Creator.Name }}]({{ .Creator.PublicUrl }}) {{ .Added }}\n" +
		"\n" +
		"## Summary: \n" +
		"\n" +
		"{{ .Description }}\n" +
		"\n" +
		"## Print Settings\n" +
		"\n" +
		"{{ with printSettings . }}\n" +
		"\n" +
		"### Printer Brand\n" +
		"\n" +
		"{{ .PrinterBrand }}\n" +
		"\n" +
		"### Printer\n" +
		"\n" +
		"{{ .Printer }}\n" +
		"\n" +
		"### Resolution\n" +
		"\n" +
		"{{ .Resolution }}\n" +
		"\n" +
		"### Infill\n" +
		"\n" +
		"{{ .Infill }}\n" +
		"\n" +
		"### Filament\n" +
		"\n" +
		"{{ .FilamentBrand }}\n" +
		"{{ .FilamentColor }}\n" +
		"{{ .FilamentMaterial }}\n" +
		"\n" +
		"### Notes\n" +
		"\n" +
		"{{ .Notes }}\n" +
		"\n" +
		"{{ end }}\n" +
		"\n" +
		"## Custom Section\n" +
		"\n" +
		"{{ range (customSections .) }}\n" +
		"{{ .Video }}\n" +
		"{{ end }}\n" +
		"\n" +
		"## Tags\n" +
		"\n" +
		"## Design Tools\n" +
		"\n" +
		"## License\n" +
		"\n" +
		"{{ .License }}\n" +
		"\n" +
		"## Thing Statistics\n" +
		"\n" +
		"{{ .ViewCount }} Views\n" +
		"{{ .DownloadCount }} Downloads\n" +
		"\n" +
		"{{ end -}}\n" +
		""
	return tmpl
}
