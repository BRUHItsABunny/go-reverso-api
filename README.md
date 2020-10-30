# go-reverso-api
Golang practice project for Reverso API regarding contexts and synonyms

### Introduction
This library is based on an older less complete private python library I made years ago, I just added some more features.
The services in this library where reversed from their respective Android clients.

Do not use this project in production environments unless you know what you are doing.

### Future of this project
I do not intend to maintain this library.

### Usage
You can install this library by using:
```
go get github.com/BRUHItsABunny/go-reverso-api
```
This library depends on my HTTP client abstraction (also not production ready yet) :
```
go get github.com/BRUHItsABunny/gOkHttp
```

You can find examples in the `main.go` inside the `_examples` directory.

These API's are free:
* Reverso Translate
* Reverso TTS
* Reverso Synonyms
* Reverso AutoComplete/Suggestions
* Reverso Context

Library provided as is, usage is at own risk and you will solely be responsible for whatever you do
