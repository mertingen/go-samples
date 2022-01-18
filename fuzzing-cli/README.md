# Project: CLI app for custom fuzzing
It gets the following parameters and provides HTTP Get requests. 
Thus, it checks the response code. It will use **./assets/letter_wl.txt** as letter and word resource.

**url:** wanted URL to be fuzzing with a parameter.
**respCode:** wanted HTTP response code to understand expected case.

### Usage

For instance, once it requests to "github.com/w", it will be returned "404" because Github did allocate some of letters and words.
However we can find some available profile nicknames in this way. If it returns 404 response code, we understand that letters are available to sign-up.  
```
go run main.go -url=github.com -respCode=404
```
