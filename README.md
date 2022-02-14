## Context
After implementing database backups we need to ensure that the data is restored correctly and no corruption happens. To verify equality, we compare the exported data from the before and after the backup. Exporting data from the DB produces a list of complex JSON objects. The objects and keys inside can be sorted randomly due to the specificity of the database.

## Task
Write a small program in Go that accepts two JSON files and prints out if they are equal.

Notes
1) Files consist of a list of JSON objects.
2) The order of the JSON objects or the keys inside the JSON object does not matter for equality.
3) You have the freedom to implement and structure it as you wish.
4) Sizes of input files can vary from a couple of lines to gigabytes.

## Simple example of equal JSON files
file1
```json
[
    {
        "id": "jhasdad",
        "name": "test json"
    },
    {
        "id": "wqweq",
        "name": "test json 2"
    }
]
```
file2
```json
[
    {
        "id": "wqweq",
        "name": "test json 2"
    },
    {
        "name": "test json",
        "id": "jhasdad"
    }
]
```

## How to use it
To run the program `go run main.go -first "./test_files/file1.json" -second "./test_files/file2.json"`