# go-c-library-example
Example of using C libraries in Go.
Test the use the ncurses C libraries to access keyboard functions.

## Credits
Based on [this](https://www.goinggo.net/2013/08/using-c-dynamic-libraries-in-go-programs.html) great post by William Kennedy.

## Build Steps

#### Build C keyboard library:
There are two build options in Makefile, `dynamic` for Mac OS X and `shared` for linux. This can be set editing the `all` make rule.
```bash
cd Lib/
make
cd ..
```

### Build and install the GO package:
In project folder:
```go
go build
go install
```

### Run test command:
In project folder:
```bash
$GOPATH/bin/gockeyboard
```