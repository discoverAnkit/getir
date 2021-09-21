# getir case study

A REST API to store and retrieve key value pairs on the fly and to interact with mongoDB.


## Running on Local

```
git@github.com:discoverAnkit/getir.git
cd getir
./gradlew build
go run main.go
```

This will download all dependencies and produce a working binary in the `./build/bin` subdirectory.

However, if gradle is not set up or is not working for some reason, we can always default to commands provided by Go itself to run tests and create a binary etc 

Example 
```
go build
go test ./...
go run main.go
```

## Local Machine Requirements:

* You need Go 1.12.7 or newer (its not test for earlier versions)
* You also need JRE 8 or later to run Gradle build system (the gradle jar files are already included in repo so they won't need to be downloaded)


## Dependency management

* [Go modules](https://github.com/golang/go/wiki/Modules) are used to manage dependencies. Clone the project outside of the `$GOPATH` or set `export GO111MODULE=on` in your shell to enable the modules.

#API Information

##GetKeyValueRecords (Mongo)
```
POST https://nameless-garden-48450.herokuapp.com/getKeyValueRecords

{
    "startDate" : "2015-01-01",
    "endDate" : "2015-02-20",
    "minCount" : 2500,
    "maxCount" : 3000
}
```
Sample Response 

```
{
    "code": 0,
    "msg": "Success",
    "records": [
        {
            "key": "PVCjLOnI",
            "createdAt": "1970-01-01T00:00:00Z",
            "totalCount": 2969
        },
        {
            "key": "qtUoRRTg",
            "createdAt": "1970-01-01T00:00:00Z",
            "totalCount": 2763
        },
        {
            "key": "qtUoRRTg",
            "createdAt": "1970-01-01T00:00:00Z",
            "totalCount": 2763
        },
        {
            "key": "ARpRHoHo",
            "createdAt": "1970-01-01T00:00:00Z",
            "totalCount": 2769
        },
        {
            "key": "ARpRHoHo",
            "createdAt": "1970-01-01T00:00:00Z",
            "totalCount": 2769
        },
        {
            "key": "XSujaXkY",
            "createdAt": "1970-01-01T00:00:00Z",
            "totalCount": 2952
        }
    ]
}
```

##SetKeyValue (InMemory)
```
POST https://nameless-garden-48450.herokuapp.com/setKeyValue

{
    "key" : "Ankit",
    "value" : "Goyal"
}
```
Sample Response

```
{
    "key": "Ankit",
    "value": "Goyal"
}
```

##GetValue (InMemory)
```
GET https://nameless-garden-48450.herokuapp.com/getValue?key=Ankit
```
Sample Response

```
{
    "key": "Ankit",
    "value": "Goyal"
}
```