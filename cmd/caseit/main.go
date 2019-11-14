package main

import (
	"fmt"
	"os"

	"github.com/runeimp/caseit"
)

func main() {
	// fmt.Println("CaseIt Main!")

	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Println("USAGE caseit 'string to_case CONVERT'")
		} else if os.Args[1] == "-v" || os.Args[1] == "--version" {
			fmt.Println("CaseIt v0.1.0")
		} else {
			fmt.Printf("Source: %q\n", os.Args[1])
			fmt.Println("    NoDilimiter:")
			fmt.Printf("lower: %s\n", caseit.NoDelimiterLowerCase(os.Args[1]))
			fmt.Printf("camel: %s\n", caseit.CamelCase(os.Args[1]))
			fmt.Printf("Title: %s (PascalCase)\n", caseit.PascalCase(os.Args[1]))
			fmt.Printf("UPPER: %s\n", caseit.NoDelimiterUpperCase(os.Args[1]))
			fmt.Println("    dot.case:")
			fmt.Printf("lower: %s\n", caseit.DotLowerCase(os.Args[1]))
			fmt.Printf("Title: %s\n", caseit.DotTitleCase(os.Args[1]))
			fmt.Printf("UPPER: %s\n", caseit.DotUpperCase(os.Args[1]))
			fmt.Println("    kebab-case:")
			fmt.Printf("lower: %s\n", caseit.KebabCase(os.Args[1]))
			fmt.Printf("Title: %s\n", caseit.KebabTitleCase(os.Args[1]))
			fmt.Printf("UPPER: %s\n", caseit.KebabUpperCase(os.Args[1]))
			fmt.Println("    snake_case:")
			fmt.Printf("lower: %s\n", caseit.SnakeCase(os.Args[1]))
			fmt.Printf("Title: %s\n", caseit.SnakeTitleCase(os.Args[1]))
			fmt.Printf("UPPER: %s\n", caseit.SnakeUpperCase(os.Args[1]))
			fmt.Println("    Space Case:")
			fmt.Printf("lower: %s\n", caseit.SpaceLowerCase(os.Args[1]))
			fmt.Printf("Title: %s\n", caseit.SpaceTitleCase(os.Args[1]))
			fmt.Printf("Smart: %s\n", caseit.SpaceSmartCase(os.Args[1]))
			fmt.Printf("UPPER: %s\n", caseit.SpaceUpperCase(os.Args[1]))
		}
	}
}
