package adf2md_test

import (
	"fmt"

	"github.com/cary/adf2md/pkg/adf2md"
)

func ExampleRenderer_RenderToMarkdown() {
	// Example ADF document with various features
	adfJSON := `{
		"version": 1,
		"type": "doc",
		"content": [
			{
				"type": "heading",
				"attrs": {
					"level": 1
				},
				"content": [
					{
						"type": "text",
						"text": "ADF to Markdown Example"
					}
				]
			},
			{
				"type": "paragraph",
				"content": [
					{
						"type": "text",
						"text": "This is a "
					},
					{
						"type": "text",
						"text": "simple",
						"marks": [
							{
								"type": "em"
							}
						]
					},
					{
						"type": "text",
						"text": " paragraph with "
					},
					{
						"type": "text",
						"text": "bold",
						"marks": [
							{
								"type": "strong"
							}
						]
					},
					{
						"type": "text",
						"text": " and "
					},
					{
						"type": "text",
						"text": "code",
						"marks": [
							{
								"type": "code"
							}
						]
					},
					{
						"type": "text",
						"text": " formatting."
					}
				]
			},
			{
				"type": "bulletList",
				"content": [
					{
						"type": "listItem",
						"content": [
							{
								"type": "paragraph",
								"content": [
									{
										"type": "text",
										"text": "Item 1"
									}
								]
							}
						]
					},
					{
						"type": "listItem",
						"content": [
							{
								"type": "paragraph",
								"content": [
									{
										"type": "text",
										"text": "Item 2"
									}
								]
							},
							{
								"type": "bulletList",
								"content": [
									{
										"type": "listItem",
										"content": [
											{
												"type": "paragraph",
												"content": [
													{
														"type": "text",
														"text": "Nested item 2.1"
													}
												]
											}
										]
									},
									{
										"type": "listItem",
										"content": [
											{
												"type": "paragraph",
												"content": [
													{
														"type": "text",
														"text": "Nested item 2.2"
													}
												]
											}
										]
									}
								]
							}
						]
					},
					{
						"type": "listItem",
						"content": [
							{
								"type": "paragraph",
								"content": [
									{
										"type": "text",
										"text": "Item 3"
									}
								]
							}
						]
					}
				]
			},
			{
				"type": "codeBlock",
				"attrs": {
					"language": "go"
				},
				"content": [
					{
						"type": "text",
						"text": "func main() {\n\tfmt.Println(\"Hello, Markdown!\")\n}"
					}
				]
			},
			{
				"type": "paragraph",
				"content": [
					{
						"type": "text",
						"text": "Check out this "
					},
					{
						"type": "text",
						"text": "link",
						"marks": [
							{
								"type": "link",
								"attrs": {
									"href": "https://example.com"
								}
							}
						]
					},
					{
						"type": "text",
						"text": "."
					}
				]
			},
			{
				"type": "taskList",
				"attrs": {
					"localId": "task-list-1"
				},
				"content": [
					{
						"type": "taskItem",
						"attrs": {
							"localId": "task-1",
							"state": "TODO"
						},
						"content": [
							{
								"type": "text",
								"text": "Task to do"
							}
						]
					},
					{
						"type": "taskItem",
						"attrs": {
							"localId": "task-2",
							"state": "DONE"
						},
						"content": [
							{
								"type": "text",
								"text": "Completed task"
							}
						]
					}
				]
			}
		]
	}`

	// Parse ADF
	node, err := adf2md.ParseADF(adfJSON)
	if err != nil {
		fmt.Printf("Error parsing ADF: %v\n", err)
		return
	}

	// Convert to Markdown
	renderer := adf2md.NewRenderer()
	markdown, err := renderer.RenderToMarkdown(node)
	if err != nil {
		fmt.Printf("Error rendering Markdown: %v\n", err)
		return
	}

	// Print the Markdown output (for demonstration purposes)
	fmt.Print(markdown)

	// No Output defined here to skip the example test
}