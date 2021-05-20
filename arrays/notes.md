# Arrays and slices

	Here we are going to test a function which takes an array of numbers and returns sum of them.
1. Write the test for adding sum of an array, and test it.
1. We will get compilation error as function undefined
1. Now write the required function to add sum of an array, also return 0
1. When we test by using go test, test will fail because of irrelevant output.
1. Write more code to iterate over the array and to return the sum.
1. Minimize the code by adding range in place of for loop.
1. Range lets you iterate over an array/slice/map/struct/channel. Every time it is called it returns two values, the index and the value. We are choosing to ignore the index value by using _
#### Disadvantage of Array: In array the size is fixed.
#### Advantage of Slice: Slice allows us to have collections of any size, so we are replacing array with slice now, syntax is same but we omit the size when using slice.
## Syntax for creating slice
1. mySlice := []int{1,2,3} rather than myArray := [3]int{1,2,3}
1. Slice with make:
   Ex: sums := make([]int, lengthOfNumbers)
   make will allows us to create a slice, so here sums is a slice variable, and the make function takes an empty array and a starting capacity (length)
1. Appending slice:
   We can use the append function which takes a slice and a new value, returning a new slice with all the items in it.
   Ex: sums = append(sums, Sum(numbers))
## Changing Array to slice.
1. So now make the array into slice and test it again, it will pass..
1. Now again create one more test which calls a function that takes multiple slices and adds each slice values, and returns the sum of each slice in a different slice.
1. Now test it, we will receive compilation error for not having function.
1. Write a variadic function that can take a variable number of arguments
        func SumAll(numbersToSum ...[]int) (sums []int) {
                return
        }

1. Now test it by using go test command.
1. Still our test fails, as we can’t use equality operator on slices, So here we can use reflect.DeepEqual which is useful for seeing if any two variables are the same like below.
        if !reflect.DeepEqual (got, want) {
            t.Errorf ("got %v want %v", got, want)
        }
1. we have to import reflect to access reflect.DeepEqual, and It's important to note that   reflect.DeepEqual is not "type safe", the code will compile even if you did something a bit silly
1. Now we need to do is iterate over the varargs, calculate the sum using our Sum function     from before and then add it to the slice we will return
1. We can also create slice by using make like below where sums is slice variable.
	sums := make ([] int, lengthOfNumbers)
1. You can use the append function which takes a slice and a new value, returning a new slice with all the items in it
	var sums [] int
        sums = append (sums, Sum(numbers))
1. we can test the coverage of our test by using " go test -cover "
1. Now our next requirement is to change SumAll to SumAllTails, where it now calculates the totals of the "tails" of each slice. The tail of a collection is all the items apart from the first one (the "head")
1. Write the test where it returns sum of tails of slices, and test it.
1. Now we will get compilation error saying ---undefined: SumAllTails
1. write the function and test it, it will fails because of unexpected output as it returns sum of all values not tails.
1. Make some changes to code and test it again, Slices can be sliced! The syntax is slice [low: high] If you omit the value on one of the sides of the: it captures everything to the side of it. In our case, we are saying "take from 1 to the end" with numbers[1:]. You might want to invest some time in writing other tests around slices and experimenting with the slice operator so you can be familiar with it.
1. What if we passed in an empty slice into our function? What is the "tail" of an empty slice? What happens when you tell Go to capture all elements from myEmptySlice [1:]?
1. Write the test for an empty slice and test it, we will get error saying:
panic: runtime error: slice bounds out of range [recovered], panic: runtime error: slice bounds out of range
1. Refactor the code to accept 0 slice and return 0.
