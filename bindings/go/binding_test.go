package tree_sitter_autoconf_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_autoconf "github.com/ravenjoad/tree-sitter-autoconf/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_autoconf.Language())
	if language == nil {
		t.Errorf("Error loading Autoconf grammar")
	}
}
