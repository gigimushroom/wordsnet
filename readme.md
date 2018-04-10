# WordsNet 
a app that returns you random words every time you request it.
The English word is fetched from Oxford Dictionary and provide your the best explanation.

## How to run it:
```
execute reader/./reader
go to your browser: localhost:8000
```
New word with detail explanation every time

## How it works under the hood:
Input: a txt format words list
The app lookup each word using Oxford API
Save to embeddedDB boltdb
Provide a random access API to return new English word from db

## If you want to *build* everything from scratch:
### Build Storager:
Follow [this tool's instruction](https://github.com/patrickbucher/ox) to build your Go CLI for the Oxford Dictionary API
Put your Go CLI binary the same directory with main.go
Build and run main.go, which will use your input file, lookup each word using the Go CLI, save details as key/value to DB.
### Build WebServer:
build reader/reader.go
check it in browser
