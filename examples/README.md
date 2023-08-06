# Examples

This directory contains examples to use `go-typechat`

## List of examples

- [Calendar](calendar/README.md): This sample translates user intent into a sequence of actions to modify a calendar. Based on [microsoft/TypeChat/examples/calendar](https://github.com/microsoft/TypeChat/tree/main/examples/calendar).
- [OpenLibrary](openlibrary/README.md): This sample translates user intent into [OpenLibrary Search API](https://openlibrary.org/dev/docs/api/search) request and calls it to get the result.

## Running examples

1. Export the following env variables

```bash
export OPENAI_API_KEY="..."
export OPENAI_MODEL="gpt-3.5-turbo"
# if not using default organisation
export OPENAI_ORGANIZATION="org-..."
```

2. Run the example from it's directory

```bash
go run main.go
```
