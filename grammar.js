/**
 * @file Tree-sitter grammar for GNU Autoconf
 * @author Raven Hallsby <raven@hallsby.com>
 * @license LGPL-3.0-or-later
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "autoconf",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
