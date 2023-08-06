package typechat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

const (
	requestPromptFormat = `You are a service that translates user requests into JSON objects of type "CalendarActions" according to the following GoLang definitions:
"""
%s
"""

Return the user request translated into a JSON object with 2 spaces of indentation and no properties with the value undefined or null or empty
`
	repairPromptFormat = `The JSON object is invalid for the following reason:
"""
%s
"""
Return revised JSON object`
)

type Translator interface {
	Translate(ctx context.Context, request string, target interface{}) error
}

type translator struct {
	*TranslatorParams
}

type TranslatorParams struct {
	ChatModel       ChatModel
	RepairAttempts  int
	TypeDefinitions *TypeDefinitions
}

func NewTranslator(params *TranslatorParams) Translator {
	return &translator{
		TranslatorParams: params,
	}
}

func (t *translator) Translate(
	ctx context.Context, request string, target interface{},
) (err error) {
	requestPrompt := t.createRequestPrompt(t.TypeDefinitions.definitions)
	messages := []string{request}

	for attempts := t.RepairAttempts + 1; attempts > 0; attempts-- {
		response, err := t.ChatModel.Send(ctx, requestPrompt, messages)
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(response), target); err != nil {
			messages = append(messages, t.createRepairPrompt(
				errors.Wrap(err, "malformed json"),
			))
			continue
		}

		validator, ok := target.(Validator)
		if !ok {
			return nil
		}

		err = validator.Validate()
		if err == nil {
			return nil
		}
		messages = append(messages, t.createRepairPrompt(err))
	}

	return
}

func (t *translator) createRequestPrompt(request string) string {
	return fmt.Sprintf(requestPromptFormat, request)
}

func (t *translator) createRepairPrompt(err error) string {
	return fmt.Sprintf(repairPromptFormat, err.Error())
}
