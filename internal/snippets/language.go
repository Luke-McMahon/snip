/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package snippets

import (
	"path/filepath"
	"strings"
)

// DetectLanguage attempts to detect the programming language from content
func DetectLanguage(content string) string {
	content = strings.TrimSpace(content)
	if content == "" {
		return ""
	}

	// Check for common patterns
	patterns := map[string][]string{
		"go": {
			"package ", "func ", "import (", "var ", "type ", "const ",
		},
		"python": {
			"def ", "import ", "from ", "class ", "if __name__", "#!/usr/bin/env python",
		},
		"javascript": {
			"function ", "const ", "let ", "var ", "=> ", "require(", "import {",
		},
		"typescript": {
			"interface ", "type ", ": string", ": number", ": boolean", "export ",
		},
		"rust": {
			"fn ", "let ", "mut ", "use ", "struct ", "impl ", "match ",
		},
		"java": {
			"public class", "private ", "public ", "static ", "import java",
		},
		"c": {
			"#include", "int main", "printf(", "scanf(", "malloc(",
		},
		"cpp": {
			"#include <iostream>", "std::", "namespace ", "class ", "template<",
		},
		"bash": {
			"#!/bin/bash", "#!/bin/sh", "echo ", "if [", "for ", "while ",
		},
		"sql": {
			"SELECT ", "INSERT ", "UPDATE ", "DELETE ", "CREATE TABLE", "FROM ",
		},
		"html": {
			"<!DOCTYPE", "<html", "<head", "<body", "<div", "<script",
		},
		"css": {
			"{", "}", "color:", "font-", "margin:", "padding:",
		},
		"json": {
			"{", "}", "[", "]", "\":", "\",",
		},
		"yaml": {
			"---", "- ", ": ", "  - ",
		},
		"dockerfile": {
			"FROM ", "RUN ", "COPY ", "ADD ", "WORKDIR ", "CMD ",
		},
	}

	// Convert to lowercase for pattern matching
	lowerContent := strings.ToLower(content)

	// Score each language
	scores := make(map[string]int)
	for lang, langPatterns := range patterns {
		for _, pattern := range langPatterns {
			if strings.Contains(lowerContent, strings.ToLower(pattern)) {
				scores[lang]++
			}
		}
	}

	// Find the language with the highest score
	var bestLang string
	var bestScore int
	for lang, score := range scores {
		if score > bestScore {
			bestScore = score
			bestLang = lang
		}
	}

	// Special cases for disambiguation
	if bestLang == "javascript" && scores["typescript"] > 0 {
		// If both JS and TS patterns found, prefer TS if it has specific TS patterns
		tsSpecific := []string{"interface ", "type ", ": string", ": number"}
		for _, pattern := range tsSpecific {
			if strings.Contains(lowerContent, pattern) {
				bestLang = "typescript"
				break
			}
		}
	}

	return bestLang
}

// DetectLanguageFromFilename detects language from a filename/extension
func DetectLanguageFromFilename(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	
	extMap := map[string]string{
		".go":         "go",
		".py":         "python",
		".js":         "javascript",
		".ts":         "typescript",
		".rs":         "rust",
		".java":       "java",
		".c":          "c",
		".cpp":        "cpp",
		".cc":         "cpp",
		".cxx":        "cpp",
		".sh":         "bash",
		".bash":       "bash",
		".zsh":        "bash",
		".sql":        "sql",
		".html":       "html",
		".htm":        "html",
		".css":        "css",
		".json":       "json",
		".yaml":       "yaml",
		".yml":        "yaml",
		".dockerfile": "dockerfile",
		".rb":         "ruby",
		".php":        "php",
		".swift":      "swift",
		".kt":         "kotlin",
		".scala":      "scala",
	}

	if lang, exists := extMap[ext]; exists {
		return lang
	}

	// Check for special filenames
	lower := strings.ToLower(filename)
	if lower == "dockerfile" || lower == "makefile" {
		return strings.ToLower(filename)
	}

	return ""
}
