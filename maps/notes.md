# TDD with Maps in Go
1.	Here we are going to learn maps with dictionary problem.
1.	First, assuming we already have some words with their definitions in the dictionary, if we search for a word, it should return the definition of it.
1.	Write a test where it will call a Search function which takes map of key value as strings, and it should return a string
1.	Test it, it will fail, as function is undefined.
1.	Write a function now, which takes a map[string]string, and one more string for comparing, also return an empty string. Declaring a Map is somewhat similar to an array. Except, it starts with the map keyword and requires two types. The first is the key type, which is written inside the []. The second is the value type, which goes right after the [}.
1.	Here key should be unique and value can be duplicated.
1.	Now test it, it will fail because of unexpected output.
1.	Add proper return statement to return the string
1.  Getting a value out of a Map is the same as getting a value out of Array map[key]
1.	Now make the test file clean by adding a function to check got and want results(assertStrings)
1.	Also change the test file by Calling Dictionary(of type map[string]string) with the data, and with this reference variable call Search function.
1.	Change the search function to a method by adding Dictionary type as method to it.
1.	Now write one more test, where we have to write test case if word is not there, The way to handle this scenario in Go is to return a second argument which is an Error type.  and Errors can be converted to a string with the .Error() method in go. check for error, it should be there beacuase the word is not present,Then test it.
1.	Add error return statement in the code, and return nil.
1.	In order to make this pass, we are using an interesting property of the map lookup. It can return 2 values. The second value is a boolean which indicates if the key was found successfully.
1.	This property allows us to differentiate between a word that doesn't exist and a word that just doesn't have a definition
1.	Write a new error by using “errors.New” if the function return true.
1.	Add this new error in a variable and return it
1.	Write one more assertError function to check the error output of “unknown word”
1.	By creating a new helper we were able to simplify our test, and start using our ErrNotFound variable so our test doesn't fail if we change the error text in the future.
1.	Now we are going to add new word to dictionary.
1.	Write one more test function to add new words to dictionary.
1.	Add code and return type to make it pass. Adding to a map is also similar to an array. You just need to specify a key and set it equal to a value
1.	An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
1.	We can create maps by using
>	var dictionary = map[string]string{}
>	var dictionary = make(map[string]string)
1.	Both approaches create an empty hash map and point dictionary at it. Which ensures that you will never get a runtime panic
1.	Write one helper function to Add Test, to make the code cleaner.
1.	Map will not throw an error if the value already exists. Instead, they will go ahead and overwrite the value with the newly provided value. This can be convenient in practice, but makes our function name less than accurate. Add should not modify existing values. It should only add new words to our dictionary.
1.	we need to modify Add to return an error, which we are validating against a new error variable, ErrWordExists. We also modified the previous test to check for a nil error, as well as the assertError function.
1.	Now write enough code to keep the errors in varaiables and use switch case to return any other errors.
1.	Here we are using a switch statement to match on the error. Having a switch like this provides an extra safety net, in case Search returns an error other than ErrNotFound
1.	We made the errors constant; this required us to create our own DictionaryErr type which implements the error interface , it makes the errors more reusable and immutable.
1.	Next, let's create a test function to Update the definition of a word
1.	Execute the test, we will get error as Update is undefined
1.	Add Update function and test it, we will get error as we require new definition which is defined in test.
1.	we now have the same issue as with Add. If we pass in a new word, Update will add it to the dictionary
1.	Write tests if word already exists for update test, and also add error return statement in update function
1.	Now we need to add  our own error type and are returning a nil error.
1.	In update function we need to search for word if it is already exists, we will update or we will return wordnotexists err.
1.	Now test it, it will pass.
1.	Again write a test function to delete a word in dictionary
1.	Write a code to delete the word, Go has a built-in function delete that works on maps. It takes two arguments. The first is the map and the second is the key to be removed.

## Learned from this Chapter
* Create maps
* Search for items in maps
* Add new items to maps
* Update items in maps
* Delete items from a map
* Learned more about errors
* How to create errors that are constants
* Writing error wrappers

