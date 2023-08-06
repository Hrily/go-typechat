# OpenLibrary go-typechat example

This sample translates user intent into [OpenLibrary Search API](https://openlibrary.org/dev/docs/api/search) request and calls it to get the result.

## Try OpenLibrary example

To run the OpenLibrary example, follow the instructions in the [examples README](https://github.com/Hrily/go-typechat/blob/main/examples/README.md#running-examples)

## Sample Runs

```
$ go run main.go
📖> books about self help
Searching:  https://openlibrary.org/search.json?sort=rating&subject=self+help
+  A child called "it"  -  Dave Pelzer
+  Atomic Habits  -  James Clear, Àlex Guàrdia i Berdiell
+  The 48 Laws of Power  -  Robert Greene
+  How to Win Friends and Influence People  -  Dale Carnegie
+  The Fault in Our Stars  -  John Green

📖> by robert green
Searching:  https://openlibrary.org/search.json?author=robert+green&sort=rating
+  The 48 Laws of Power  -  Robert Greene
+  AS 48 LEIS DO PODER  -  Robert Greene
+  The Laws of Human Nature  -  Robert Greene
+  Las 48 leyes del poder  -  Robert Greene
+  Las 48 Leyes del Poder  -  Robert Greene

📖> book named human nature
Searching:  https://openlibrary.org/search.json?sort=rating&title=human+nature
+  The Laws of Human Nature  -  Robert Greene
+  The Nature Of Human Values  -  Milton Rokeach
+  A Treatise of Human Nature  -  David Hume
+  The Red Queen: Sex and the Evolution of Human Nature  -
+  On human nature  -  Edward Osborne Wilson


```
