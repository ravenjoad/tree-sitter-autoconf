import XCTest
import SwiftTreeSitter
import TreeSitterAutoconf

final class TreeSitterAutoconfTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_autoconf())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Autoconf grammar")
    }
}
