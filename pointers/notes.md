# TDD using pointers in Go
1. Pointers will point to a value, and we can change them, so rather than taking copy of a value, we will directly take the value and make changes according to our requirement.
1. We can represent with *, for example *Wallet, we can call this as pointer to wallet.
1. Here we are going to develop a bitcoin project, where we will 1st write the test and we will pass the amount to Deposit method, and we will get the balance by using balance method, and compare the result with expected output and test it, we will get error as methods are undefined.
1. Add required methods and test it.it will pass
1. Now add one more test case for withdraw, and test it.
1. it will fail as withdraw is undefined, add required method and test it.
1. We will get error as “got 0 want 10”, because the data is not updating, to update it we are using pointers, where all changed data will be updated.
1. Now we will define an Stringer interface, This interface is defined in the fmt package and lets you define how your type is printed when used with the %s format string in prints, we also need to change the format to %s in test to have this stringer working.
1. Change the testing file, so that we can have helper function for deposit and withdraw.
1. Now we need to return error if the withdrawing amount is more than balance, so for the write one test which is for insufficient balance, and call withdraw function with more amount.
1. Now test it, we will get error, so write error return statement in withdraw function, also write condition for withdrawing more amount which is more than balance.
1. errors.New creates a new error with a message of your choosing.
1. t.Fatal which will stop the test if it is called.

## Learned from this Chapter
### Pointers
Go copies values when you pass them to functions/methods so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples could be very large data or perhaps things you intend only to have one instance of (like database connection pools).
### nil
Pointers can be nil
When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception, the compiler won't help you here.
Useful for when you want to describe a value that could be missing
### Errors
Errors are the way to signify failure when calling a function/method.
By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.
This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.