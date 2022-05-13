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
    | `float64` (single precision)| 64 bits or 8 bytes
- String
    - They are declared using `string` keyword and are always enclosed in `""` (double quotes). A `string` data type takes 16 bytes of memory.
- Boolean
    - They are declared using `bool` keyword and are either `true` or `false` not `0` or `1`. A `bool` data type takes 1 byte of memory.
- Some other data types are Arrays, Slices and Maps.

## Variables and their Declarations
- Syntax for declaring variables
    - `var <variable name> <data type> = <value>`
- Examples
    - `var s string = "Hello there"`
    - `var i int = 2468`
    - `var b bool = true`
    - `var f float64 = 987.123`
## Printing Variables
