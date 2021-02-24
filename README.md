# Tree Utility in Go

Given a path prints out the folders in a tree-like fashion. Optionally prints the files in the folders and their size in bytes.

## Usage
### Provide a path such as `.` or `/myfolder` without an argument to print only folders in a tree like structure.

Example:
`go run main.go testdata` 

```
├───project
├───static
│       ├───a_lorem
│       │       └───ipsum
│       ├───css
│       ├───html
│       ├───js
│       └───z_lorem
│               └───ipsum
└───zline
        └───lorem
                └───ipsum
```

### Optional add argument `-f` which will print the files and the size of the files in bytes.

Example:
`go run main.go testdata -f`
```
├───project
│       ├───file.txt (19b)
│       └───gopher.png (70372b)
├───static
│       ├───a_lorem
│       │       ├───dolor.txt (empty)
│       │       ├───gopher.png (70372b)
│       │       └───ipsum
│       │               └───gopher.png (70372b)
│       ├───css
│       │       └───body.css (28b)
│       ├───empty.txt (empty)
│       ├───html
│       │       └───index.html (57b)
│       ├───js
│       │       └───site.js (10b)
│       └───z_lorem
│               ├───dolor.txt (empty)
│               ├───gopher.png (70372b)
│               └───ipsum
│                       └───gopher.png (70372b)
├───zline
│       ├───empty.txt (empty)
│       └───lorem
│               ├───dolor.txt (empty)
│               ├───gopher.png (70372b)
│               └───ipsum
│                       └───gopher.png (70372b)
└───zzfile.txt (empty)
```

## Run Tests
Run the tests from root folder:
`go test -v`.




