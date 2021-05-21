# TDD using struct in Go
We are going to find perimeter of rectangle here.
1.	So first write the test to call perimeter function with two arguments which are of type float64.
1.	Now test the function, we need to add mod file, after adding mod file, test the function again, it will show compilation error as Perimeter function not found.
1.	Add appropriate function which takes two arguments and return 0.
1.	Now again test by using go test command, we will get error because of unexpected output which is 0.
1.	Now add the required return statement to return output
1.	Our code will pass now. Add another test to call Area of rectangle, Add the required function with return statement to get proper output.
1.	Now our complete code will pass, but it is not a clean code, as we are not specifying anything about rectangle.
1.	So the best solution for this is to create a struct of type Rectangle.
1.	Struct: Struct is a named collection of fields where we can store data.
1.	Create a struct like below
> type Rectangle struct {
-    W float64
-    H float64
> }
1.	Change the test file to access struct and assign values, now test the function, it will show error as not enough arguments in call.
1.	Now change the functions to access struct and its values by using dot(.) operator
1.	Now add a test to get area of circle, and test it, we will get Compilation error as Circle struct undefined.
1.	Change the code by adding Struct to circle, but when we test we will get an error saying 
“Cannot use values (type Circle) as type Rectangle in argument to Area”
1.	In other programming languages we can define same method name with different parameters, but in go it`s not possible, we can create same method name in other package, but that’s not good, so here we will be using methods.
##	Method
 A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.
####	Syntax:
>	func (receiverName ReceiverType) MethodName(args)
1.	Change the test to call method by using an instance
1.	If we test our code now, we will get error as rectangle. Area undefined (type Rectangle has no field or method Area) same for circle also.
1.	So let’s create related methods which returns 0.
1.	Again we will get error if we test, because of unrelated output.
1.	Add proper return statements and test it once again, it will pass.
1.	When your method is called on a variable of that type, you get your reference to its data via the receiverName  variable. In many other programming languages this is done implicitly and you access the receiver via this.
1.	It is a convention in Go to have the receiver variable be the first letter of the type.
1.	Here we are having some duplication in our code.
1.	So now we have to create a function which will have collection of shapes, so we can call our rectangle and circle from here.
1.	For this we need to use interface
##	Interface: 
These are a very powerful concept in statically typed languages like Go because they allow you to make functions that can be used with different types and create highly-decoupled code while still maintaining type-safety
#### Syntax: 
> type Shape interface {}
1. 	In test file, we need to create a helper function which takes Shape, If we try to call this with something that isn't a shape, then it will not compile.
1.	 Now create a shape interface which has method Area that return float64.
1.	So according to shape the area method will be called from interface.
1.	Here rectangle has method Area so it satisfies Shape interface, as well for circle also, but if pass any other shape which doesn’t contain Area method with float64 return type, then we will encounter an error.
1.	Now test it again, it will pass.
1.	Now we will go through Table driven tests
1.	Change our test file into table driven tests, by creating an anonymous struct which will have all the fields that we need, and then iterate over the struct using range loop, and check for expected output.
1.	The table driven tests makes it very easy to add a new test case, also Table based tests to make your assertions(errors) clearer and your suites easier to extend & maintain
1.	We can also test each test by using  “go test -run TestArea/Rectangle”
