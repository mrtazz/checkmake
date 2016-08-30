# Rules

Rules for checking Makefiles are written in Go code. They are simple types
satisfying the following interface:

```
type Rule interface {
	Name() string
	Description() string
	Run(parser.Makefile, RuleConfig) RuleViolationList
}
```

They are added as subpackages of `rules` and mostly define an empty struct
type which provides the functions to satisfy the interface and then gets
registered in the rule registry like so:

```
func init() {
	rules.RegisterRule(&Rule1{})
}
```
