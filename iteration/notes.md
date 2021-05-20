# Iteration with TDD
## Here in Golang there are no while, do, until keywords, we can only use for.
For loop doesnot require parentheses surrounding the three components of the for statement and the braces { } are always required.
1.	Here we need to iterate a loop for 5 times to print the same character.
1.	Write a test function which repeats a character 5 times
1.	Test it with go test, it will fail, and reports a compilation error which says, Repeat function undefined.
1.	Add a function for iterating the character and returns 0
1.	When we test the code, it will fail because of unexpected output.
1.	Refactor the code by adding proper return statement to return proper output.
1.	Now the test will pass.
1.	Write a benchmarking test which runs b.N times(the testing.B gives access to b.N) and measures how long it will take( The amount of times the code is run shouldn't matter to us, the framework will determine what is a "good" value for that to let us have some decent results.), we can test benchmark time by using following command "go test -bench="."", and we will get output as follows
        
                PS C:\Users\SRS\gocode\src\goWithTesting\iteration> go test -bench="."
                goos: windows
                goarch: amd64
                pkg: iteration
                cpu: Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
                BenchmarkRepeat-8       16110712                72.05 ns/op
                PASS
                ok      iteration       1.316s
1.  Now as an exercise Change the test so a caller can specify how many times the character is repeated and then fix the code
1.	Add more test cases as per requirement and refactor the code to pass.
