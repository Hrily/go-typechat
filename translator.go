package typechat

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

const (
	requestPromptFormat = `You are a service that translates user requests into JSON objects of type "%s" according to the following GoLang definitions:
"""
%s
"""

Return the user request translated into a JSON object with 2 spaces of indentation

Ensure the following for JSON:
- omit any fields with null or undefined values
`
	repairPromptFormat = `The JSON object is invalid for the following reason:
"""
%s
"""
Return revised JSON object`
)

// Translator ...
type Translator interface {
	Translate(ctx context.Context, request string, target interface{}) error
}

type translator struct {
	*TranslatorParams
}

// TranslatorParams are the dependencies for Translator
type TranslatorParams struct {
	ChatModel       ChatModel
	RepairAttempts  int
	TypeDefinitions *TypeDefinitions
}

// NewTranslator ...
func NewTranslator(params *TranslatorParams) Translator {
	return &translator{
		TranslatorParams: params,
	}
}

// Translate a natural language `request` into JSON and unmarshal it into `target`
// if the JSON returned by the language model is not valid and max
// `RepairAttempts` will be made to repair the JSON using diagnostics produced
// in validation.
// If the JSON is not valid after `RepairAttempts` attempts, an error will be
// returned.
// `target` can optionally implement `Validator` to perform additional
// validation on the response, whose error will be used in diagnostics to
// repair the response.
func (t *translator) Translate(
	ctx context.Context, request string, target interface{},
) (err error) {
	requestPrompt := t.createRequestPrompt(t.TypeDefinitions.definitions, target)
	messages := []*ChatModelMessage{
		{System: &requestPrompt},
		{User: &request},
	}

	for attempts := t.RepairAttempts + 1; attempts > 0; attempts-- {
		response, err := t.ChatModel.Send(ctx, messages)
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(response), target); err != nil {
			messages = t.appendRepairDiagnostics(
				messages, response, errors.Wrap(err, "malformed json"),
			)
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

		messages = t.appendRepairDiagnostics(messages, response, err)
	}

	return
}

func (t *translator) createRequestPrompt(
	request string, target interface{},
) string {
	typeName := t.getTypeName(target)
	return fmt.Sprintf(requestPromptFormat, typeName, request)
}

func (t *translator) getTypeName(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

func (t *translator) appendRepairDiagnostics(
	messages []*ChatModelMessage, response string, err error,
) []*ChatModelMessage {
	repairPrompt := t.createRepairPrompt(err)
	messages = append(messages,
		&ChatModelMessage{AI: &response},
		&ChatModelMessage{User: &repairPrompt},
	)
	return messages
}

func (t *translator) createRepairPrompt(err error) string {
	return fmt.Sprintf(repairPromptFormat, err.Error())
}
