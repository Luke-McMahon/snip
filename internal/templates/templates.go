/*
Copyright © 2025 Luke McMahon <me@lmc.id.au>
*/
package templates

import (
	"fmt"
	"strings"
	"time"

	"github.com/Luke-McMahon/snip/internal/snippets"
)

// Template represents a snippet template
type Template struct {
	Name        string
	Description string
	Language    string
	Tags        []string
	Content     string
}

// GetBuiltinTemplates returns a list of built-in templates
func GetBuiltinTemplates() map[string]Template {
	return map[string]Template{
		"go-func": {
			Name:        "go-func",
			Description: "Go function template",
			Language:    "go",
			Tags:        []string{"go", "function", "template"},
			Content: `func {{.FuncName}}({{.Params}}) {{.ReturnType}} {
	{{.Body}}
	return {{.ReturnValue}}
}`,
		},
		"go-struct": {
			Name:        "go-struct",
			Description: "Go struct template",
			Language:    "go",
			Tags:        []string{"go", "struct", "template"},
			Content: `type {{.StructName}} struct {
	{{.Fields}}
}`,
		},
		"py-func": {
			Name:        "py-func",
			Description: "Python function template",
			Language:    "python",
			Tags:        []string{"python", "function", "template"},
			Content: `def {{.FuncName}}({{.Params}}):
    """{{.Docstring}}"""
    {{.Body}}
    return {{.ReturnValue}}`,
		},
		"py-class": {
			Name:        "py-class",
			Description: "Python class template",
			Language:    "python",
			Tags:        []string{"python", "class", "template"},
			Content: `class {{.ClassName}}:
    """{{.Docstring}}"""
    
    def __init__(self{{.InitParams}}):
        {{.InitBody}}
    
    def {{.MethodName}}(self{{.MethodParams}}):
        {{.MethodBody}}`,
		},
		"js-func": {
			Name:        "js-func", 
			Description: "JavaScript function template",
			Language:    "javascript",
			Tags:        []string{"javascript", "function", "template"},
			Content: `function {{.FuncName}}({{.Params}}) {
    {{.Body}}
    return {{.ReturnValue}};
}`,
		},
		"js-arrow": {
			Name:        "js-arrow",
			Description: "JavaScript arrow function template",
			Language:    "javascript",
			Tags:        []string{"javascript", "arrow", "function", "template"},
			Content: `const {{.FuncName}} = ({{.Params}}) => {
    {{.Body}}
    return {{.ReturnValue}};
};`,
		},
		"bash-script": {
			Name:        "bash-script",
			Description: "Bash script template",
			Language:    "bash",
			Tags:        []string{"bash", "script", "template"},
			Content: `#!/bin/bash

# {{.Description}}

set -e  # Exit on error

{{.Variables}}

{{.Body}}

echo "Script completed successfully"`,
		},
		"dockerfile": {
			Name:        "dockerfile",
			Description: "Dockerfile template",
			Language:    "dockerfile", 
			Tags:        []string{"docker", "dockerfile", "template"},
			Content: `FROM {{.BaseImage}}

WORKDIR {{.WorkDir}}

COPY {{.CopySource}} {{.CopyDest}}

RUN {{.InstallCommands}}

EXPOSE {{.Port}}

CMD {{.Command}}`,
		},
	}
}

// RenderTemplate renders a template with given variables
func RenderTemplate(template Template, vars map[string]string) string {
	content := template.Content
	
	// Simple template variable replacement
	for key, value := range vars {
		placeholder := "{{." + key + "}}"
		content = strings.ReplaceAll(content, placeholder, value)
	}
	
	return content
}

// CreateSnippetFromTemplate creates a snippet from a template
func CreateSnippetFromTemplate(templateName string, title string, vars map[string]string) (*snippets.Snippet, error) {
	templates := GetBuiltinTemplates()
	template, exists := templates[templateName]
	if !exists {
		return nil, fmt.Errorf("template '%s' not found", templateName)
	}
	
	content := RenderTemplate(template, vars)
	
	snippet := &snippets.Snippet{
		Title:     title,
		Content:   content,
		Language:  template.Language,
		Tags:      append(template.Tags, "generated"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	return snippet, nil
}

// ListTemplates returns a list of available template names and descriptions
func ListTemplates() []Template {
	templates := GetBuiltinTemplates()
	var list []Template
	for _, t := range templates {
		list = append(list, t)
	}
	return list
}
