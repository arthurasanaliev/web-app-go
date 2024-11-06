### Nov 4
- net/http package -> internet stuff
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

- errors package -> error handling
- err := errors.New(message) -> declaring new error

- hmtl/template package -> package for go-template files
- parsed, _ := template.ParseFiles(file) -> recieves parsed file
- err := parsed.Execute(w, nil) -> executes tmpl file

- go packages -> pkg/package/package.go (organizing code using packages)
`import mod-name/package`

### Nov 5
- go-layouts -> templates that can be wrapped in single tmpl file and used later in different tmpl files
- `{{block "name" .}} {{end}} -> {{define "name"}} {{end}}` -> block that can be modified in different files
- `{{define "name"}} {{end}} -> {{template "name" .}}` -> whole template

- path/filepath package -> access to files, etc
`pages, err := filepath.Glob(pattern)` -> returns slice of strings
`name := filepath.Base(filePath)` -> returns base (everything after last /) name as string
