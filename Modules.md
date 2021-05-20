# Go Modules:
    Go module introduced in Go version 1.11. Go mod provides access to operations on modules, Modules are how Go manages dependencies, Go projects required to be in Gopath without go modules.
    Go code is grouped into packages, and packages are grouped into modules. Your module specifies dependencies needed to run your code, including the Go version and the set of other modules it requires.
    Using Modules is pretty straightforward. Select any directory outside GOPATH as the root of your project, and create a new module with the go mod init / go mod init [module path] command.
## go mod init command
    The go mod init command initializes and writes a new go.mod file in the current directory, The go.mod file must not already exist, init accepts one optional argument which is a module path, if we omit this then init will attempt to infer the module path using import comments in .go files
    After we execute the above command , a new go.mod file will be created which contains dependencies and version as below.

module hello

go 1.16

    Now, we can add dependencies by using go get command, after executing, our go.mod file will look like below.

module hello

go 1.16

require golang.org/x/tools v0.1.1 // indirect

## Go mod tidy
    go mod tidy ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module's packages and dependencies, and it removes requirements on modules that don't provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.
### We can also run go build or go test command to install all the missing dependencies automatically. 

## References:
#####	[Modules](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world#go-modules)
#####	[Create go Module](https://golang.org/doc/tutorial/create-module)
#####	[Go module references](https://golang.org/ref/mod#go-mod-init)
#####	[Why go mod instead of gopath](https://insujang.github.io/2020-04-04/go-modules/)

