package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/hrily/go-typechat"
	"github.com/hrily/go-typechat/examples"
	"github.com/hrily/go-typechat/examples/openlibrary/client"
	"github.com/hrily/go-typechat/examples/openlibrary/models"
)

const (
	maxBooksToPrint = 5
)

var (
	openlibrary = client.New()
	translator  = newTranslator(1)
	ctx         = context.Background()
)

func main() {
	examples.Interactive(process, "📖> ")
}

func process(request string) {
	searchRequest := &models.SearchRequest{}
	if err := translator.Translate(
		ctx, request, searchRequest,
	); err != nil {
		fmt.Println(err)
		return
	}

	response, err := openlibrary.Search(ctx, searchRequest)
	if err != nil {
		panic(err)
	}

	printResponse(response)
	fmt.Println()
}

func printResponse(response *models.SearchResponse) {
	if len(response.Books) == 0 {
		fmt.Println("No books found")
		return
	}
	if len(response.Books) > maxBooksToPrint {
		response.Books = response.Books[:maxBooksToPrint]
	}

	for _, book := range response.Books {
		fmt.Println("+ ", book.Title, " - ", strings.Join(book.Authors, ", "))
	}
}

func newTranslator(repairAttempts int) typechat.Translator {
	definitions, err := typechat.NewTypeDefinitionsFromFile("models/search_request.go")
	if err != nil {
		panic(err)
	}

	return typechat.NewTranslator(&typechat.TranslatorParams{
		ChatModel:       typechat.NewOpenAIChatModel(),
		RepairAttempts:  repairAttempts,
		TypeDefinitions: definitions,
	})
}
