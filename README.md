# Golang Notes
Golang is a fast, statically typed, compiled language
that feels like a dynamically typed, interpreted language.

Go has a concept of types that is either explicitly declared
by the programmer or is infered by the compiler.

## Data Types in Golang
- Integer
    - `int` => It means signed integer.
    - `uint` => It means unsigned integer.

    | Data Type     | Memory        
    | ------------- | -------------
    | `int`      | 4 bytes for 32-bit machines and 8 bytes for 64 bit    
    | `int8`  / `uint8`       | 8 bits or 1 bytes
    | `int16` / `uint16`      | 16 bits or 2 bytes
    | `int32` / `uint32`      | 32 bits or 4 bytes
    | `int64` / `uint64`      | 64 bits or 8 bytes
- Float
    | Data Type     | Memory        
    | ------------- | -------------
    | `float32` (single precision)| 32 bits or 4 bytes
    | `float64` (double precision)| 64 bits or 8 bytes
- String
    - They are declared using `string` keyword and are always enclosed in `""` (double quotes). A `string` data type takes 16 bytes of memory.
- Boolean
    - They are declared using `bool` keyword and are either `true` or `false` not `0` or `1`. A `bool` data type takes 1 byte of memory.
- Some other data types are Arrays, Slices and Maps.

## Variables and their Declarations
- Syntax for declaring variables
    - ```golang
        var <variable name> <data type> = <value>
    ```
- Examples
    ```golang
    var s string = "Hello there"
    var i int = 2468
    var b bool = true
    var f float64 = 987.123
    ```
## Printing Variables
- There are three methods available in "fmt" package for printing variables.
    - `fmt.Print()`: This simply prints the value inside it and does not add a new line at the end after printing.
    - `fmt.Println()`: This prints the value inside it and adds a new line at the end after printing. Also adds a space when multiple values are passed to it.
    - `fmt.Printf()`: This function can print formatted output using format specifiers same as in C/C++ language.
        - Format Specifiers: Ther are used to format different kinds of data types and are always preceded with a `%` sign.
            | Format Specifier     | Description        
            | ------------- | -------------
            | %v | Prints the value in default format
            | %t | Used for boolean values
            | %d | Used for printing integers
            | %f | Used for decimal numbers
            | %.3f | Prints the decimal value upto 3 decimal places
            | %c | Used for printing characters
            | %s | Used for printing plain string
            | %q | Used for printing characters or string in quotes
            | %T | Prints the type of value
    - Examples
        ```golang
        fmt.Printf("Template String with %d %s", num, str)
        ```
- Scope of variables in Golang is same as in C/C++ languages.
- Zero Values: This is the default value given to an uninitialised variable in Golang.
    - `bool` -> `false`
    - `int` -> `0`
    - `float64` -> `0.0`
    - `string` -> `""`
- Taking User Input
    ```golang
    fmt.Scanf("%<format specifier> (s)", <variables>)
    ```
    - `fmt.Scanf` returns 2 arguments `count` and `err`
        - `count` : the number of arguments that the function writes to
        - `err` : any error thrown during the execution of the function.
- Type Casting
    - Data types can be converted to other data types, but this does not guarantee that the value will remain intact.
- Constants: They are variables whose value after initialization can't be changed.
    - Example: `const <const name> <data type> = <value>`
    - Here data type is optional when declaring constants
    - Theu are of two types:
        - Untyped Constant: Constants declared without specifying data type.
        - Typed Constant: Constants declared with data type.
## Operators
- There are 5 types of operators in Golang:
    - Arithmatic Operators: `+, -, *, /, %, ++, --`
    - Comparison Operators: `==, <=, >=, <, >, !=`
    - Assignment Operators: `=, +=, -=, *=, /=, %=`
    - Bitwise Operators: `&, |, <<, >>, ^`
    - Logical Operators: `&&, ||, !`
## Control Statements
```golang
    if condition { // parenthesis around condition is not necessary
        // executes when condition is true
    } else { // This else statement must be on the same line where the if block ends
        // executes when condition is false
    }


    if condition {
        // executes when condition is true
    } else if condition { // This else-if statement must be on the same line where the if block ends
        // executes when condition is true
    } else { 
        // executes when all above conditions are false
    }
```
- Switch Case in golang has implicit break statement unlike in languages like C/C++.
- `fallthrough` : This keyword in switch block forces the execution to fall through all the conditions uptill default condition is reached or a condition without this keywoord is reached.
- switch with conditions, does not need ```expression```
```golang
    switch expression {
    case value_1 / condition:
        // executes when expression matches to this value
    case value_2 / condition:
        // executes when expression matches to this value
    case value_3 / condition:
        // executes when expression matches to this value
    default:
        // executes when none of the above values matches to expression
    }
```
## Loops
```golang
for initialization; condition; post {
    // your code
}
```
Example
```golang
for i := 1; i < 3; i++ {
    fmt.Println(i)
}
```
Infinite for loop
```golang
for {
    // your code
}
```
## Arrays
- The elements are stored in contigous memory.
- Array can store elements of same data type.
- Arrays are of fixed length. Length cannot be changed after their declaration.
    - length : It denotes the no of elements in the array.
    - capacity : It denotes the no of elements that it can contain in the array.
    - For Arrays the above two properties are same.
- Array elements can be accessed and can be modified from their indices.
- `range` keyword is used for looping through an array, slices and maps.

Syntax for declaration:
```golang
var <array name> [<size of the array>] <data type>
// Array declaration
var numbers [100] int
var words [50] string

// Array initialization
var nums [3]int = [3]int {1, 2, 3}
// OR
nums := [3]int{1, 2, 3}
// OR
nums := [...]int{1, 2, 3} // ... is called elipsis, compiler detects the length of the array
```

## Slices
- A slice is an continuous segment of an underlying array. It is more powerful, flexible, convenient than array as it is a variable length sequence and stores elements of same type.
- A slice has 3 components:
    - Pointer: It points to the first element of the array which is accessible through slice.
    - Length: It is the total no of elements present in the array. `len()`
    - Capacity: It is the maximum size upto which the slice can expand. `cap()`
- Declaration and initialization:
    
    ```golang
    <slice name> := []<data type>{values}
    ```
    In the above code, the compiler first creates the array and then returns the slice reference to it.
    
    ```golang
    <slice name> := <array name>[startIndex : endIndex]
    slice1 := array[0 : 3] // elements upto index 2 are included starting from 0
    slice2 := array[1 : 6] // elements upto index 5 are included starting from 1
    slice3 := array[ : 4] // elements upto index 3 are included starting from 0
    slice4 := array[ : ] // all elements are included
    ```
    In the above code, the compiler uses the created array and creates all the slices.

    ```golang
    <slice name> := make([]<data type>, length, capacity)
    slice := make([]int, 15, 100)
    ```
    In the above code, `make()` function is usd to create an empty slice. This is the common way of creating slices.
- `len(array)` and `cap(array)` are same for array but can be different from slice. As `capacity(slice)` starts from the starting index of slice.
- **When we change any value in a slice, it gets affected in the array as well, because slice is an reference to the array.**
- Appending values to a slice
    ```golang
    func append(s []T, values ...T)[]T
    // Example
    slice = append(slice, 11, 22, 33)
    ```
- Copying values from a slice: Both the slices must be of same data type.
    ```golang
    func copy(dst, src T[])int
    // Example
    num := copy(dest_slice, src_slice)
    ```
- Looping through a slice is same as looping through an array.

## Maps
- Map is a data structure which provides an unordered collection of key-value pairs. It is implemented using hash tables.
- They are same as dictionary in python, hash table in java.
- Syntax:
```golang
var <map name> map[key data type]<value data type>
// Example
var pair map[string]int
// This creates a nil map, a map with no key value pairs.
// A nil map does not allow to add key-value pairs to it.
```
- Declaration and initialization

    ```golang
    <map name> := map[key data type]<value data type>{<key-value-pairs>}
    // Example
    nums := map<string>int{"one":1, "two":2}
    ```
- Using `make()` function to create a map
    ```golang
    <map name> := make(map[key data type]<value data type>, <initial capacity>)
    // initial capacity is optional
    // Example
    nums := make(map[string]int)
    ```
- Adding and Updating the map
    ```golang
    // Adding key value pairs
    nums["one"] = 1
    nums["two"] = 2
    nums["three"] = 33
    // Updating key value pairs
    nums["three"] = 3
    ```
- Deletinga key value pair from the map
    ```golang
    delete(nums, "three")
    // nums => map[one:1 two:2]
    ```
- A map can be truncated or cleared by reinitialisation
    ```golang
    nums = make(map[string]int)
    ```
- Looping through a map
    ```golang
    for key, value := range nums {
		fmt.Println(key, "=>", value)
	}
    ```
- Accessing a key-value pair
    ```golang
    value, found := languages["one"]
    // value contains the key value if it exists, else zero value of the respective data type
    // found is true if the key value exists, else false
    ```

## Functions
- Syntax:
    ```golang
    func <function name>(<parameters>)<return type>{
        // function body
    }
    <function name>(<argument(s)>)
    ```

- Example
    ```golang
    func addition(int a, int b) int{
        sum := a + b
        return sum
    }
    addition(123, 456)
    sum := addition(100, 200)
    ```
- Naming convention
    - It must starts with a letter and cannot have spaces.
    - It can have any no of letters and symbols.
    - It is case sensitive.
- Return Types
    - A function can return multiple values.
    - A function can return values which are named at the function declaration.
    - A variadic function is a function that accepts variable no of arguments of same data type.
    
        ```golang
        func <func_name>(param-1 type, param-2 type, para-3 ...type)<return type>

        // This function expects 0 or more integers
        // Within the function nums will have slice of all arguments.
        func summation(nums ...int)int{

        }
        ```
    - Anonymous Function: It is a function without a name.
    - High Order Functions: It accepts a function as an argument or returns a function as its return type.
    - `defer` statement delays the execution of a function untill the surrounding function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

## Pointers
- The concept of pointers here is same as it is in C/C++.
- Declaration
    
    ```golang
    var <pointer_name> *<data_type>
    var ptr_i *int
    var ptr_s *string
    ```
- Initialization
    
    ```golang
    var <pointer name> *<data type> = &<variable name>
    <pointer name> := &<variable name>
    i := 5
    var ptr_i *int = &i
    s := "Golang"
    ptr_s := &s
    ```
- Dereferencing a pointer
    
    ```golang
    s := "Golang"
    ptr_s := &s
    fmt.Println(*ptr_s) // Golang
    *ptr_str = "Go"
	fmt.Println(*ptr_str) // Go
    ```
- Passing by value and reference: These are also same as in C/C++.
    - Slices and maps are passed by reference by default and all the other data types are passed by making their copy.