# TDD with Dependency injection
1. In Go, interfaces enable us to use dependency inversion. We are able to use different implementations in our code, as long as they satisfy the interface we have defined. We use dependency injection to tell the application which implementation to use.
1.	We want to write a function that greets someone, just like we did in the hello-world chapter but this time we are going to be testing the actual printing.
1.	But how can we test this? Calling fmt.Printf prints to stdout, which is pretty hard for us to capture using the testing framework.
1.	What we need to do is to be able to inject (which is just a fancy word for pass in) the dependency of printing
1.	Our function doesn't need to care where or how the printing happens, so we should accept an interface rather than a concrete type
1.	If we do that, we can then change the implementation to print to something we control so that we can test it. In "real life" you would inject in something that writes to stdout
1.	If you look at the source code of fmt.Printf you can see a way for us to hook in, Under the hood Printf just calls Fprintf passing in os.Stdout.
1.	See the below example
a.	func Printf(format string, a ...interface{}) (n int, err error) {
b.	return Fprintf(os.Stdout, format, a...)
c.	}
1.	Fprintf will take os.Stdout, what it means?What is the 1st argument for Fprintf?
1.	func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
1.	Writer is an interface, that will have write method which takes []byte variable and returns number of bytes and an error.
1.	So we know  we're ultimately using Writer to send our greeting somewhere.
1.	Now lets write the test for it, So here we are using Buffer from bytes package which implements writer interface.
1.	So we'll use it in our test to send in as our Writer and then we can check what was written to it after we invoke Greet
1.	Call the Greet function by using address of buffer(reference variable that holds bytes.Buffer{})
1.	Check the wanted output and test it, it will fail, as method is undefined.
1.	Write some code that takes pointer to bytes.Buffer and a string for name.
1.	Now test it, it will fail asit didn’t got required output.
1.	Now change the fmt.Printf to fmt.Fprintf which will use writer.
1.	fmt.Fprintf is like fmt.Printf but instead takes a Writer to send the string to, whereas fmt.Printf defaults to stdout.
1.	Now the test will pass.
1.	Earlier the compiler told us to pass in a pointer to a bytes.Buffer. This is technically correct but not very useful.
1.	To demonstrate this, try wiring up the Greet function into a Go application where we want it to print to stdout.
1.	fmt.Fprintf allows us to pass in an io.Writer which we know both os.Stdout and bytes.Buffer implement.
1.	So if we change *bytes.Buffer to io.Writer in greet function, we can use it in our main function also, where it calls greet function with os.Stdout, name arguments.
1.	We can now also call our greet from a handler function where it will display on browser.



