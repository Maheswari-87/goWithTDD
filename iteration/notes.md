# Iteration with TDD
## Here in Golang there are no while, do, until keywords, we can only use for.
1.	Here we need to iterate a loop for 5 times to print the same character.
1.	Write a test function which repeats a character 5 times
1.	Test it with go test, it will fail, and reports a compilation error which says, Repeat function undefined.
1.	Add a function for iterating the character and returns 0
1.	When we test the code, it will fail because of unexpected output.
1.	Refactor the code to return proper output.
1.	Now the test will pass.
1.	Write a benchmarking test which runs b.N times and measures how long it will take( The amount of times the code is run shouldn't matter to us, the framework will determine what is a "good" value for that to let you have some decent results.)
1.	Add more test cases as per requirement and refactor the code to pass.
