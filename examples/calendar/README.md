# Calendar go-typechat example

This sample translates user intent into a sequence of actions to modify a calendar. Based on [microsoft/TypeChat/examples/calendar](https://github.com/microsoft/TypeChat/tree/main/examples/calendar).

## Try Calendar example

To run the Calendar example, follow the instructions in the [examples README](https://github.com/Hrily/go-typechat/blob/main/examples/README.md#running-examples)

## Sample Runs

```
$ go run main.go
📆> add meeting with Gavin at 1pm and remove meetings with Sasha
{
  "actions": [
    {
      "addEvent": {
        "event": {
          "timeRange": {
            "startTime": "13:00"
          },
          "participants": [
            "Gavin"
          ]
        }
      }
    },
    {
      "removeEvent": {
        "eventReference": {
          "participants": [
            "Sasha"
          ]
        }
      }
    }
  ]
}
```
