# TDD in Go
Test-Driven development (TDD) is a process of developing and running automated test before actual development of the application. It is a skill that needs practice to develop but by being able to break problems down into smaller components that you can test, you will have a much easier time writing software.

## Writing a test is just like writing a function, we have to follow some rules to write a test:
1. It needs to be in a file with a name like xxx_test.go
1. The test function must start with the word Test
1. The test function takes one argument only t *testing.T
1. In order to use the *testing.T type, you need to import "testing".

## Hello world program with TDD
  Flow of Writing Unit test Case on Hello World Program: 
1. Write test Case [on Hello World which return Hello World as a String]
1. Check with go test command, Test will fail
1. We will get a compile time error on console saying Hello function undefined.
1. Add Functionality of Hello World function which return empty string
1. Run go test command, Again unit test will fail because of unexpected output (empty string).
1. Check and Write more code to make the test pass (add return statement which is a hello world string)
1. Unit test will pass, and we will get required output (Hello world string)
1. Add more test cases for new requirements
1. Refactor the code to fulfill the requirements
1. Repeat the same procedure to add more requirements from step1

## functions used
### t.Helper()
It is needed to tell the test suite that this method is a helper. By doing this when it fails the line number reported will be in our function call rather than inside our test helper. This will help other developers track down problems easier

### t.Errof()
We are calling the Errorf method on our t which will print out a message and fail the test. The f stands for format which allows us to build a string with values inserted into the placeholder values %q. When you made the test fail it should be clear how it works.
### godoc
1. we can have documentation fo our package by using this godoc.
1. we can install it manually by using "go get golang.org/x/tools/cmd/godoc"
1. Now run "godoc -http :8000"
1. So now we can acces our documentation in browser by using "http://localhost:8000/pkg/goWithTesting/integers/"