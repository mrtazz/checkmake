# Formatting

checkmake provides a custom formatter class aptly named "`CustomFormatter`".
This formatter supports passing in a text/template that will be used to format
the output for each violation. When rendering the template, each
`RuleViolation` struct is rendered individually, followed by a `\r\n` carriage
return. This means you can access its fields in the template, e.g.:

```
"{{.LineNumber}}:{{.Rule}}:{{.Violation}}"
```

The custom formatter can be enabled either via the `--format=` command line
option or the `default.format` option in the configuration file.
