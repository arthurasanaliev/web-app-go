### Nov 4
- net/http module -> internet stuff
- http.HandleFunc(path, page handler function) -> page on path
- w http.ResponseWriter -> to write on page (used to fmt.Fprintf(w, text))
- r *http.Request -> 
- http.ListenAndServe(port, nil) -> run wep app on port

- Encapsulation -> Encapsulation is defined as the wrapping up of data under a single unit:
```go
// addTwoValues is accessible within current package
func addTwoValues(x, y int) int {
}
// AddTwoValues is accessible outside current package as well
func AddTwoValues(x, y int) int {
}
```

- Just above the function, specify what it does following this convention:
```go
// About is about page handler
func About(w http.ResponseWriter, r *http.Request) {
}
```
