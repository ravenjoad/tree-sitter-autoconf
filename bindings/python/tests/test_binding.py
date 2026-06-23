from unittest import TestCase

import tree_sitter
import tree_sitter_autoconf


class TestLanguage(TestCase):
    def test_can_load_grammar(self):
        try:
            tree_sitter.Language(tree_sitter_autoconf.language())
        except Exception:
            self.fail("Error loading Autoconf grammar")
