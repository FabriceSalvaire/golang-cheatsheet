# Go

https://golang.org

Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.

Go was publicly announced in November 2009,[29] and version 1.0 was released in March 2012.

## Trend

From https://www.tiobe.com/tiobe-index

*   1 C
*   2 Python /
*   3 Java \
*   4 C++
*   5     C# — @Microsoft
*   6         Visual Basic — @Microsoft
*   7 JavaScript
*   8     PHP
*   9         R
*  10         SQL
*  11     Groovy — @Java https://en.wikipedia.org/wiki/Apache_Groovy
*  12         Perl
*  13 Go /
*  14     Swift — @Apple
*  15 Ruby \
*  16         ASM
*  17         MATLAB
*  18         Delphi — ~Pascal @Windows
*  19     Objective-C — @Apple / @NextStep

## Book

* https://www.digitalocean.com/community/books/how-to-code-in-go-ebook

# How to Run

```
Usage:

        go <command> [arguments]

The commands are:

        bug         start a bug report
        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         add dependencies to current module and install them
        install     compile and install packages and dependencies
        list        list packages or modules
        mod         module maintenance
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
        vet         report likely mistakes in packages

Use "go help <command>" for more information about a command.

Additional help topics:

        buildmode   build modes
        c           calling between Go and C
        cache       build and test caching
        environment environment variables
        filetype    file types
        go.mod      the go.mod file
        gopath      GOPATH environment variable
        gopath-get  legacy GOPATH go get
        goproxy     module proxy protocol
        importpath  import path syntax
        modules     modules, module versions, and more
        module-get  module-aware go get
        module-auth module authentication using go.sum
        module-private module configuration for non-public modules
        packages    package lists and patterns
        testflag    testing flags
        testfunc    testing functions

Use "go help <topic>" for more information about that topic.
```

To run a go file:
```
go run FILENAME.go
```

## Module

* a name is exported if it begins with a capital letter
