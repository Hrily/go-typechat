# go-typechat

GoLang port of [microsoft/TypeChat](https://github.com/microsoft/TypeChat)

## What is TypeChat?

From [microsoft/TypeChat](https://github.com/microsoft/TypeChat/tree/main#typechat)

> TypeChat is a library that makes it easy to build natural language interfaces using types.
>
> Building natural language interfaces has traditionally been difficult. These apps often relied on complex decision trees to determine intent and collect the required inputs to take action. Large language models (LLMs) have made this easier by enabling us to take natural language input from a user and match to intent. This has introduced its own challenges including the need to constrain the model's reply for safety, structure responses from the model for further processing, and ensuring that the reply from the model is valid. Prompt engineering aims to solve these problems, but comes with a steep learning curve and increased fragility as the prompt increases in size.
>
> TypeChat replaces _prompt engineering_ with _schema engineering_.
>
> Simply define types that represent the intents supported in your natural language application. That could be as simple as an interface for categorizing sentiment or more complex examples like types for a shopping cart or music application. For example, to add additional intents to a schema, a developer can add additional types into a discriminated union. To make schemas hierarchical, a developer can use a "meta-schema" to choose one or more sub-schemas based on user input.
>
> After defining your types, TypeChat takes care of the rest by:
>
> 1. Constructing a prompt to the LLM using types.
> 2. Validating the LLM response conforms to the schema. If the validation fails, repair the non-conforming output through further language model interaction.
> 3. Summarizing succinctly (without use of a LLM) the instance and confirm that it aligns with user intent.
>
> Types are all you need!

## Getting Started

Get the module:

```bash
go get github.com/hrily/go-typechat
```

### Usage

<details>
  <summary><b>Translate Book Search Request</b></summary>

  ```golang
  package main

  import (
  	"context"
  	"fmt"

  	"github.com/hrily/go-typechat"
  )

  // SearchRequest to search for books
  // The user can specify one or more of the following fields
  type SearchRequest struct {
  	// Title will find any books with the given title
  	Title string `json:"title,omitempty"`
  	// Author will find any books with the given author
  	Author string `json:"author,omitempty"`
  	// Subject will find any books about the given subject
  	// eg: "tennis rules" will find books about "tennis" and "rules"
  	Subject string `json:"subject,omitempty"`
  	// Query will find any books with the given query
  	// Is used when the user input does not correspond to any of the other fields
  	Query string `json:"query,omitempty"`
  }

  const searchRequestDefinition = "" +
  	"// SearchRequest to search for books" +
  	"// The user can specify one or more of the following fields" +
  	"type SearchRequest struct {" +
  	"	// Title will find any books with the given title" +
  	"	Title string `json:\"title,omitempty\"`" +
  	"	// Author will find any books with the given author" +
  	"	Author string `json:\"author,omitempty\"`" +
  	"	// Subject will find any books about the given subject" +
  	"	// eg: \"tennis rules\" will find books about \"tennis\" and \"rules\"" +
  	"	Subject string `json:\"subject,omitempty\"`" +
  	"	// Query will find any books with the given query" +
  	"	// Is used when the user input does not correspond to any of the other fields" +
  	"	Query string `json:\"query,omitempty\"`" +
  	"}"

  func main() {
  	translator := typechat.NewTranslator(&typechat.TranslatorParams{
  		ChatModel:       typechat.NewOpenAIChatModel(),
  		RepairAttempts:  0,
  		TypeDefinitions: typechat.NewTypeDefinitions(searchRequestDefinition),
  	})

  	ctx := context.Background()
  	request := "by JK Rowling"

  	searchRequest := &SearchRequest{}
  	if err := translator.Translate(
  		ctx, request, searchRequest,
  	); err != nil {
  		fmt.Println(err)
  		return
  	}

  	// Will print:
  	// &main.SearchRequest{Title:"", Author:"JK Rowling", Subject:"", Query:""}
  	fmt.Printf("%#v\n", searchRequest)
  }
  ```
</details>

<br/>

Check [examples](examples/README.md) for more usage examples
