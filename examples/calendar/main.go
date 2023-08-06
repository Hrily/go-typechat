package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hrily/go-typechat"
	"github.com/hrily/go-typechat/examples"
	"github.com/hrily/go-typechat/examples/calendar/models"
)

var (
	translator = newTranslator(1)
	ctx        = context.Background()
)

func main() {
	examples.Interactive(process, "📆> ")
}

func process(request string) {
	response := &models.CalendarActions{}
	if err := translator.Translate(
		ctx, request, response,
	); err != nil {
		fmt.Println(err)
		return
	}

	j, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(j))
}

func newTranslator(repairAttempts int) typechat.Translator {
	definitions, err := typechat.NewTypeDefinitionsFromFile("models/calendarActions.go")
	if err != nil {
		panic(err)
	}

	return typechat.NewTranslator(&typechat.TranslatorParams{
		ChatModel:       typechat.NewOpenAIChatModel(),
		RepairAttempts:  repairAttempts,
		TypeDefinitions: definitions,
	})
}
