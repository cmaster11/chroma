package g

import (
	. "github.com/alecthomas/chroma" // nolint
	"github.com/alecthomas/chroma/lexers/internal"
)

func makeGoIfTemplateRules() Rules{
	goIfTemplateRules := make(Rules)

	// Use all go template rules, but stipping the {{ area
	for key, rules := range goTemplateRules {
		if key == "root" {
			goIfTemplateRules[key] = []Rule{
				{`\(`, Operator, Push("subexpression")},
				{`"(\\\\|\\"|[^"])*"`, LiteralString, nil},
				Include("expression"),
			}
			continue
		}

		if key == "template" {
			// It is now the root
			continue
		}

		goIfTemplateRules[key] = rules
	}

	return goIfTemplateRules
}

var GoIfTemplate = internal.Register(MustNewLexer(
	&Config{
		Name:    "Go if-template",
		Aliases: []string{"go-if-template"},
	},
	makeGoIfTemplateRules(),
))
