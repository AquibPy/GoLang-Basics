# Go Basics

![Go Gopher][gopher_image]

[gopher_image]: https://blog.golang.org/go-brand/Go-Logo/SVG/Go-Logo_Blue.svg

This repository is designed to help beginners learn the fundamental concepts and syntax of the Go programming language. It contains a collection of examples and explanations covering various topics, including:

- Hello World
- Variables and Data Types
- Control Structures (if-else, switch, loops)
- Functions
- Arrays and Slices
- Maps
- Structs
- Interfaces
- Concurrency (Goroutines and Channels)
- Error Handling
- File I/O
- Packages and Modules

## Getting Started

### Installing Go

Before you can start working with this repository, you'll need to have Go installed on your machine. Follow these steps to install Go:

1. Download the latest version of Go from the official website: https://golang.org/dl/
2. Follow the installation instructions for your operating system.
3. After the installation is complete, verify that Go is installed by running the following command in your terminal or command prompt:

```go
go version
```
This should print the installed version of Go.

### Cloning the Repository

Once you have Go installed, you can clone this repository using the following command:

```
git clone https://github.com/your-username/go-basics.git
```

After cloning the repository, navigate to the respective folder for each topic to explore the examples and explanations.

## Usage

Each folder in this repository contains one or more Go source files (`.go` files) that demonstrate a specific concept or feature of the Go programming language. You can run these examples by navigating to the corresponding folder and executing the following command:

```go
go run hello/hello.go
```
This command will compile and run the `hello.go` file in the hello directory.

## Lexer

The lexer (short for lexical analyzer or scanner) is a component of the Go compiler or interpreter that breaks down the source code into a sequence of tokens. These tokens are the smallest meaningful units that the compiler or interpreter can understand, such as keywords, identifiers, operators, literals, and punctuation marks.

The lexer is the first stage of the compilation or interpretation process, and it is responsible for recognizing and classifying the input characters into meaningful tokens. It reads the source code character by character and applies a set of rules (defined by the language's lexical grammar) to identify and categorize the tokens.

In Go, the lexer is implemented in the `go/scanner` package, which provides a scanner that reads the source code and produces a stream of tokens. This repository includes examples and explanations on how to use the `go/scanner` package to tokenize Go source code.

When you write a program in a programming language like Go, you essentially write it in a human-readable form using words, symbols, and syntax rules defined by the language. However, computers can only understand machine-readable instructions, which are typically represented as binary code.
The lexer is the first step in the process of translating your human-readable code into a form that the computer can understand. Its job is to break down your source code into smaller, meaningful units called tokens.

Think of it like this: when you read a sentence in English, your brain automatically separates the sentence into individual words and punctuation marks. The lexer does a similar thing, but instead of separating words, it separates the source code into tokens like keywords (e.g., func, package), identifiers (e.g., variable and function names), operators (e.g., +, -, *), literals (e.g., numbers and strings), and punctuation marks (e.g., {, }, ;).

## Types in Go lang

Basic Types

* Boolean: bool (values are true or false)
* Numeric:

   * Integers: int8, int16, int32, int64, uint8, uint16, uint32, uint64, int (32 or 64 bits), uint (32 or 64 bits), uintptr (an unsigned integer large enough to store the uninterpreted bits of a pointer value)
   * Floating-point: float32, float64
   * Complex: complex64, complex128


* String: string (immutable sequence of bytes)
* Rune: rune (an alias for int32, represents a Unicode code point)

* Aggregate Types

    * Array: A numbered sequence of elements of a single type with a fixed length.
    * Slice: A reference to an underlying array with added functionality. Slices are dynamically-sized and can grow or shrink as needed.
    * Struct: A sequence of named elements called fields, each of which has a name and type.
    * Map: An unordered group of key-value pairs of the same type.
    * Channel: A channel is a communication mechanism used for sending and receiving values between goroutines (lightweight threads).

* Reference Types

    * Pointer: A pointer holds the memory address of a value of the same type.
    * Function: Functions are first-class values in Go and can be assigned to variables, passed as arguments, and returned from other functions.
    * Interface: An interface is a type that defines a set of method signatures. Any type that provides the methods defined by the interface can be considered an implementation of that interface.

## Variable Shadowing

Variable shadowing occurs when a variable declared in an inner scope has the same name as a variable declared in an outer scope, causing the inner variable to hide or "shadow" the outer variable within the inner scope.

## Comma ok or Error ok Syntax

In Go, the "comma ok" or "error ok" syntax is a way to check for successful operations and handle potential errors when using built-in functions or methods that return two values. The second value is typically a boolean or an error indicating whether the operation was successful or if an error occurred.
This is because in Go we don't have try catch block.

## Random Numbers

**math/rand** -  It provides a pseudorandom number generator (PRNG) for generating non-cryptographically secure random numbers, suitable for general-purpose applications.

**crypto/rand** - It provides a cryptographically secure pseudorandom number generator (CSPRNG) for generating unpredictable random numbers, suitable for security-sensitive applications like cryptography.

## Memory Management in Go

Go uses a tracing garbage collector to automatically manage memory. The garbage collector runs periodically in the background to free up unused memory. This means developers don't need to manually allocate or free memory.

There are three main functions for memory allocation and initialization in Go:

- new() allocates memory for a type and returns a pointer. It is similar to malloc() in C.

- make() allocates and initializes memory for maps, slices, and channels by preallocating the underlying data structure.

- delete() removes references to an object to allow it to be collected earlier by the garbage collector.

The garbage collector runs concurrently so it doesn't block the program during collection. This helps ensure responsiveness. Developers can control memory by removing references with delete().

## Pointers in Go

Pointers in Go allow us to reference a memory location that stores a value of a particular type. We use pointers because sometimes we need to pass large structures by reference instead of by value for efficiency.

In simple terms, a pointer contains the memory address of a value. We define a pointer using an asterisk (*) before the type. For example, a pointer to an int would be defined as *int. 

We can create a pointer by using the address of (&) operator on a variable, which gives us the memory location of that variable. Then we can pass that pointer to functions to modify the original variable.

Pointers are useful in Go because they let us modify values more efficiently for large data structures like slices, maps and channels that are passed by reference instead of copying by value. This avoids expensive copies of large amounts of data.

So in summary, pointers allow passing references to data instead of copies, improving performance for large or complex values.

## Arrays

Arrays in Go are fixed-length sequences of elements of the same type. Some key properties of arrays in Go include:

    - The length is part of the array's type, so [5]int and [10]int are different types. Arrays must be initialized to a fixed length.

    - You can declare and initialize an array like var a [5]int or a := [5]int{1, 2, 3, 4, 5}. Each element is initialized to the zero value by default. 

    - Unlike slices, arrays are values and can be directly assigned or passed to functions. However, their size is immutable so they cannot be appended to or resized like slices.

    - Common array operations include accessing elements via their index (a[0]), iterating with a for loop, and passing to functions to work with the elements.

## Slices

1. **Arrays vs. Slices**:
   - Arrays are like fixed-size boxes where you put things in. Once you define the size of an array, you can't change it.
   - Slices are like resizable boxes. They can grow or shrink as needed.

2. **Making a Slice**:
   - You can make a slice using the `make()` function or by slicing an existing array or slice.
   - Example: `s := make([]int, 5)` creates a slice of integers with a length of 5.

3. **Slicing**:
   - Slicing is like taking a piece of a pie. You specify a range of elements you want from an existing slice or array.
   - Example: `sliced := s[1:4]` creates a new slice containing elements from index 1 to index 3 of the original slice.

4. **Appending**:
   - You can add more items to a slice using the `append()` function. It makes the slice bigger.
   - Example: `s = append(s, 4)` adds the number 4 to the end of the slice.

5. **Length and Capacity**:
   - Every slice has a length, which is how many items it currently holds, and a capacity, which is how many items it can hold without resizing.
   - Example: `s := make([]int, 3, 5)` creates a slice with a length of 3 and a capacity of 5.

Slices are really handy because they let you work with collections of data easily, and they can change size to fit your needs.

# Maps in Go

In Go, a map is a built-in data structure that stores a collection of key-value pairs. It's like a dictionary or an associative array in other programming languages.

## Key-Value Pairs

- A map holds a collection of key-value pairs, where each key is unique within the map, and each key is associated with exactly one value.
- Keys and values can be of any data type, as long as the key types and value types are consistent within the map. For example, you can have a map where keys are strings and values are integers.

## Declaration and Initialization

```go
// Declaration and initialization using make()
var myMap map[string]int
myMap = make(map[string]int)

// Declaration and initialization using a map literal
anotherMap := map[string]int{"apple": 5, "banana": 3, "orange": 2}
```
## Accessing and Modifying Values

```go
// Accessing values
value := myMap["key"]

// Modifying values or adding new key-value pairs
myMap["newKey"] = 10

// Checking for key existence
value, ok := myMap["key"]
if ok {
    fmt.Println("Value exists:", value)
} else {
    fmt.Println("Key does not exist")
}

// Deleting key-value pair
delete(myMap, "key")
```

## Structs in Go

In Go, a struct is a composite data type that groups together zero or more fields of different data types under a single name. It's like a record or a structure in other programming languages.

## Definition

- You define a struct using the `type` keyword followed by the struct's name and a list of its fields enclosed in curly braces `{}`.
- Each field in the struct has a name and a data type.

```go
type Person struct {
    name string
    age  int
}
```

## Creating Instances

- You create instances (also known as objects or structs) of a struct by declaring variables of that struct type and initializing their fields.

```go
var p1 Person
p1.name = "Alice"
p1.age = 30

p2 := Person{name: "Bob", age: 25}
```

## Accessing Fields

- You can access fields of a struct using the dot (`.`) operator followed by the field name.

```go
fmt.Println(p1.name) // Output: Alice
fmt.Println(p2.age)  // Output: 25
```

## Anonymous Structs

- You can also create anonymous structs without specifying a name. These are useful for temporary data structures.

```go
p3 := struct {
    name string
    age  int
}{"Charlie", 35}
```

## Embedding

- You can embed one struct within another to create a hierarchy of data structures.

```go
type Address struct {
    city    string
    country string
}

type Person struct {
    name    string
    age     int
    address Address
}
```

Structs are commonly used in Go to represent complex data types, such as users, products, or any other entity that has multiple properties. They provide a way to organize related data and make the code more readable and maintainable.

## Ellipsis

In Go, the ellipsis (`...`) is used in a few different contexts:

1. **Variadic Functions**: When defining a function, you can use the ellipsis before the type of the last parameter to indicate that the function can accept a variable number of arguments of that type. These are called variadic functions.

    ```go
    func exampleFunc(nums ...int) {
        // nums is a slice of int
        for _, num := range nums {
            fmt.Println(num)
        }
    }
    ```

    Here, `nums ...int` indicates that `exampleFunc` can accept zero or more `int` arguments.

2. **Variadic Arguments**: When calling a variadic function, you can pass any number of arguments of the specified type, and Go will pack them into a slice.

    ```go
    exampleFunc(1, 2, 3, 4, 5)
    ```

    In this call, `1, 2, 3, 4, 5` will be packed into a slice of `int` and passed to `exampleFunc`.

3. **Slices**: The ellipsis can also be used when passing a slice as arguments to a variadic function. It tells Go to unpack the slice and treat its elements as individual arguments.

    ```go
    nums := []int{1, 2, 3, 4, 5}
    exampleFunc(nums...)
    ```

    Here, `nums...` unpacks the `nums` slice into individual arguments and passes them to `exampleFunc`.

In all cases, the ellipsis (`...`) is a handy tool for dealing with variable numbers of arguments in functions and for manipulating slices.

## Methods

In Go, methods are functions that are associated with a particular type. They allow you to define behavior or actions that can be performed on values of that type. Here's a simple explanation:

1. **Definition**:
   - Methods are defined by specifying the receiver type before the function name.
   - The receiver type is the type on which the method operates.

2. **Usage**:
   - Methods are called on values of the receiver type using the dot (`.`) notation.

3. **Receiver Types**:
   - Methods can be associated with any named type in Go, including built-in types, structs, and user-defined types.

### Example:

Let's say we have a struct called `Rectangle`:

```go
type Rectangle struct {
    width  float64
    height float64
}
```

We can define a method called `Area` for the `Rectangle` type:

```go
func (r Rectangle) Area() float64 {
    return r.width * r.height
}
```

Here's what's happening in this code:

- `func (r Rectangle) Area() float64 { ... }`:
  - This defines a method named `Area` for the `Rectangle` type.
  - `(r Rectangle)` specifies that `Area` is a method of the `Rectangle` type.
  - `r` is the receiver variable, which represents the instance of `Rectangle` on which the method is called.
  - `float64` is the return type of the method.

Now, we can create a `Rectangle` instance and call its `Area` method:

```go
func main() {
    rectangle := Rectangle{width: 10, height: 5}
    area := rectangle.Area()
    fmt.Println("Area:", area) // Output: Area: 50
}
```

In this example, `rectangle.Area()` calculates the area of the `rectangle` instance using the `Area` method we defined earlier.

Methods in Go provide a way to attach behaviors to types, making the code more organized and readable. They are widely used in Go for defining functionalities associated with structs and other types.

## Defer

In Go, the `defer` statement is used to ensure that a function call is performed later in a program's execution, typically for purposes of cleanup. The deferred call is executed when the surrounding function returns, regardless of whether the function returns normally or because of a panic.

### Key Points:

1. **When to Use `defer`**:
   - Often used for resource cleanup tasks, such as closing files or releasing locks.
   - Ensures that the cleanup code is executed even if the function encounters an error or returns early.

2. **Syntax**:
   - The `defer` statement is placed before the function call that you want to defer.

3. **Execution Order**:
   - Deferred calls are executed in last-in, first-out (LIFO) order.

### Example:

Let's look at an example where `defer` is used to close a file:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }

    // Ensure the file is closed when main function returns
    defer file.Close()

    // Do some operations with the file
    fmt.Println("File opened successfully")
}
```

### Explanation:

1. **Opening a File**:
   - `file, err := os.Open("example.txt")` attempts to open a file.
   - If there is an error opening the file, it prints the error and returns early.

2. **Deferring a Function Call**:
   - `defer file.Close()` ensures that `file.Close()` is called when the `main` function returns, no matter where the return happens in the function.

3. **Order of Execution**:
   - If multiple `defer` statements are used, they are executed in reverse order of their appearance.

### Multiple `defer` Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Start")

    defer fmt.Println("Deferred 1")
    defer fmt.Println("Deferred 2")
    defer fmt.Println("Deferred 3")

    fmt.Println("End")
}
```

### Output:

```
Start
End
Deferred 3
Deferred 2
Deferred 1
```

### Explanation:

- The `defer` statements are executed in reverse order when the `main` function returns, ensuring a predictable and controlled way to clean up resources or perform necessary final steps.

# JSON Handling in Go:

1. **Importing the JSON Package:**
   In Go, you start by importing the `"encoding/json"` package. This package provides functions for encoding (converting Go data structures to JSON) and decoding (converting JSON back to Go data structures).

   ```go
   import "encoding/json"
   ```

2. **Defining Go Structs:**
   Before you can encode or decode JSON, you need to define Go structs that represent your data. Each field in the struct corresponds to a key in the JSON object.

   ```go
   type Person struct {
       Name string
       Age  int
   }
   ```

3. **Encoding to JSON:**
   To convert a Go data structure (like a struct) into JSON, you use the `json.Marshal()` function.

   ```go
   person := Person{Name: "Alice", Age: 30}
   jsonBytes, err := json.Marshal(person)
   if err != nil {
       // Handle error
   }
   fmt.Println(string(jsonBytes))
   ```

   This will print `{"Name":"Alice","Age":30}`.

4. **Decoding from JSON:**
   To convert JSON back into a Go data structure, you use the `json.Unmarshal()` function.

   ```go
   var newPerson Person
   err := json.Unmarshal(jsonBytes, &newPerson)
   if err != nil {
       // Handle error
   }
   fmt.Println(newPerson)
   ```

   This will print `{Alice 30}`.

### Explanation:

- **Encoding:** `json.Marshal()` converts Go data structures into JSON bytes.
- **Decoding:** `json.Unmarshal()` converts JSON bytes into Go data structures.
- **Structs:** Define Go structs to represent your data, with fields corresponding to JSON keys.
- **Error Handling:** Always check for errors when encoding or decoding JSON.

Absolutely! Let's extend the explanation to include JSON tags for aliasing fields, omitting empty fields, and handling sensitive data like passwords.

## JSON Tags in Go:

1. **Alias Fields:**
   You can specify custom JSON keys for struct fields using tags. This is helpful when the field names in your struct don't match the keys you want in the JSON.

   ```go
   type Person struct {
       FirstName string `json:"first_name"`
       LastName  string `json:"last_name"`
   }
   ```

   When encoded to JSON, `FirstName` will be represented as `"first_name"`.

2. **Omit Empty Fields:**
   You can specify that a field should be omitted from the JSON if it has a zero value (like an empty string or zero integer) by using the `omitempty` option.

   ```go
   type Person struct {
       Name string `json:"name,omitempty"`
       Age  int    `json:"age,omitempty"`
   }
   ```

   If `Name` or `Age` is empty or zero, they won't appear in the JSON output.

3. **Password Handling:**
   When dealing with sensitive data like passwords, you might want to prevent them from being encoded to JSON at all.

   ```go
   type User struct {
       Username string `json:"username"`
       Password string `json:"-"`
   }
   ```

   The `-` tag means that `Password` field will be ignored during encoding.

### Putting it Together:

Here's an example struct incorporating all these concepts:

```go
type User struct {
    Username string `json:"username"`
    Password string `json:"-"`
    Email    string `json:"email,omitempty"`
}
```

- `Username` will be encoded as `"username"` in JSON.
- `Password` won't be included in JSON at all.
- `Email` will only be included in JSON if it's not empty.

This way, you can control how your Go structs are represented in JSON, including aliasing, omitting empty fields, and handling sensitive data.

## What is `go mod`?

`go mod` is a tool in the Go programming language used to manage dependencies (external packages or libraries your code relies on) and modules (a collection of related Go packages). It helps you keep track of which versions of dependencies you are using and makes it easier to share and collaborate on Go projects.

## Basic Concepts

- **Module**: A collection of Go packages stored in a directory with a `go.mod` file at its root.
- **Dependencies**: Other Go modules your module needs to work.

## Common `go mod` Commands

Here are the essential `go mod` commands you'll use:

1. **Initialize a New Module**

   ```sh
   go mod init <module-path>
   ```

   Creates a new `go.mod` file in the current directory. `<module-path>` is typically the URL of your module (e.g., `github.com/user/repo`).

2. **Add a Dependency**

   ```sh
   go get <dependency>
   ```

   Adds a new dependency to your project. For example, `go get github.com/pkg/errors`.

3. **Tidy Up Dependencies**

   ```sh
   go mod tidy
   ```

   Removes any dependencies that are not used in your code and adds any dependencies that are needed but not listed in `go.mod`.

4. **Download Dependencies**

   ```sh
   go mod download
   ```

   Downloads the modules required by your project into the local cache.

5. **Verify Dependencies**

   ```sh
   go mod verify
   ```

   Checks that the dependencies in your local cache have not been modified since they were downloaded.

6. **Vendor Dependencies**

   ```sh
   go mod vendor
   ```

   Copies all dependencies into a `vendor` directory. This can be useful for projects that need to ensure all dependencies are stored locally (e.g., for deployment).

7. **List Dependencies**

   ```sh
   go list -m all
   ```

   Lists all the modules your project depends on.

8. **Graph Dependencies**

   ```sh
   go mod graph
   ```

   Displays a graph of module dependencies. This can help you understand how modules are related.

9. **Edit go.mod Directly**

   ```sh
   go mod edit
   ```

   Provides commands to edit the `go.mod` file programmatically (e.g., adding or removing dependencies).

10. **View Module Cache**

    ```sh
    go clean -modcache
    ```

    Cleans the module cache, removing all downloaded modules.

## Example Workflow

1. **Create a New Module**

   ```sh
   go mod init github.com/yourname/yourproject
   ```

2. **Add Dependencies**

   ```sh
   go get github.com/some/dependency
   ```

3. **Tidy Up Dependencies**

   ```sh
   go mod tidy
   ```

4. **Download All Dependencies**

   ```sh
   go mod download
   ```

5. **Verify Dependencies**

   ```sh
   go mod verify
   ```



### Go Concurrency Vs Parallelism

| Feature       | Concurrency                                         | Parallelism                                         |
|---------------|-----------------------------------------------------|-----------------------------------------------------|
| **Definition**| Managing multiple tasks at once, pausing and resuming| Executing multiple tasks simultaneously             |
| **Example**   | One chef cooking by alternating between tasks        | Two chefs cooking different tasks at the same time  |
| **Focus**     | Structure and design of tasks                       | Actual execution of tasks                           |
| **Mechanism** | Goroutines in Go                                     | Multiple CPU cores                                  |
| **Goal**      | Improve responsiveness and structure                | Improve execution speed                             |

</br>

![Concurrency Vs Parallelism](https://www.freecodecamp.org/news/content/images/2022/12/1-1.png)

### Concurrency

- **Definition**: Concurrency is about dealing with multiple tasks at once by managing the execution of these tasks in a way that they can be paused and resumed. It's about the structure and design of a program.
- **Example**: Imagine you are cooking a meal. You can start boiling water, then while the water is boiling, you start chopping vegetables. You are not doing both tasks at the exact same time, but you are handling multiple tasks by switching between them.
- **In Go**: Go achieves concurrency using Goroutines. A Goroutine is a lightweight thread managed by the Go runtime. When you launch a Goroutine using the `go` keyword, it runs concurrently with other Goroutines.

   ```go
   package main

   import (
      "fmt"
      "time"
   )

   func sayHello() {
      for i := 0; i < 5; i++ {
         time.Sleep(100 * time.Millisecond)
         fmt.Println("Hello")
      }
   }

   func sayWorld() {
      for i := 0; i < 5; i++ {
         time.Sleep(100 * time.Millisecond)
         fmt.Println("World")
      }
   }

   func main() {
      go sayHello()
      go sayWorld()
      time.Sleep(1 * time.Second)
   }
   ```

In this example, `sayHello` and `sayWorld` run concurrently.

### Parallelism

- **Definition**: Parallelism is about executing multiple tasks at the exact same time. It's about the actual execution of tasks.
- **Example**: Imagine you have two chefs in the kitchen. One chef is boiling water while the other chef is chopping vegetables. Both tasks are being done simultaneously.
- **In Go**: Parallelism is achieved when the program runs on multiple processors or cores, allowing multiple Goroutines to run at the same time.

   ```go
   package main

   import (
      "fmt"
      "runtime"
      "sync"
   )

   func main() {
      runtime.GOMAXPROCS(2) // Set the maximum number of CPUs that can be executing simultaneously

      var wg sync.WaitGroup
      wg.Add(2)

      go func() {
         defer wg.Done()
         for i := 0; i < 5; i++ {
            fmt.Println("Hello")
         }
      }()

      go func() {
         defer wg.Done()
         for i := 0; i < 5; i++ {
            fmt.Println("World")
         }
      }()

      wg.Wait()
   }
   ```

In this example, by setting `GOMAXPROCS` to 2, we are allowing the program to use two CPU cores, which can run Goroutines in parallel.

### Summary

- **Concurrency** is about managing multiple tasks and making progress on them.
- **Parallelism** is about executing multiple tasks at the same time.
- Go uses **Goroutines** for concurrency.
- Go can achieve **parallelism** if there are multiple CPU cores available.

# Sync in Go Lang

Synchronization is a crucial concept in concurrent programming, including in Go. It ensures that multiple Goroutines can safely access shared resources without causing data races or other issues. Go provides several synchronization mechanisms in its `sync` package. Let's go through the most commonly used synchronization tools in Go:

## WaitGroup

A `WaitGroup` waits for a collection of Goroutines to finish. You can add the number of Goroutines to wait for using `Add`, and each Goroutine calls `Done` when it finishes. The `Wait` method blocks until all Goroutines have finished.

### Example:

```go
   package main

   import (
      "fmt"
      "sync"
   )

   func worker(id int, wg *sync.WaitGroup) {
      defer wg.Done() // Mark this Goroutine as done when it finishes
      fmt.Printf("Worker %d starting\n", id)
      // Simulate work with Sleep
      fmt.Printf("Worker %d done\n", id)
   }

   func main() {
      var wg sync.WaitGroup

      for i := 1; i <= 3; i++ {
         wg.Add(1) // Increment the WaitGroup counter
         go worker(i, &wg)
      }

      wg.Wait() // Wait for all Goroutines to finish
   }
```

### Mutex

A `Mutex` is used to provide mutual exclusion, allowing only one Goroutine to access a critical section of code at a time. This prevents data races when multiple Goroutines read and write shared variables.

#### Example:

```go
   package main

   import (
      "fmt"
      "sync"
   )

   var (
      counter int
      mu      sync.Mutex
   )

   func increment(wg *sync.WaitGroup) {
      defer wg.Done()
      mu.Lock()
      counter++
      mu.Unlock()
   }

   func main() {
      var wg sync.WaitGroup

      for i := 0; i < 1000; i++ {
         wg.Add(1)
         go increment(&wg)
      }

      wg.Wait()
      fmt.Println("Counter:", counter)
   }
```

### RWMutex

A `RWMutex` is a reader/writer mutual exclusion lock. It allows multiple readers or one writer, but not both. This is useful when you have frequent reads and infrequent writes.

#### Example:

```go
   package main

   import (
      "fmt"
      "sync"
   )

   var (
      counter int
      rwMu    sync.RWMutex
   )

   func read(wg *sync.WaitGroup) {
      defer wg.Done()
      rwMu.RLock()
      defer rwMu.RUnlock()
      fmt.Println("Counter value:", counter)
   }

   func write(wg *sync.WaitGroup) {
      defer wg.Done()
      rwMu.Lock()
      counter++
      rwMu.Unlock()
   }

   func main() {
      var wg sync.WaitGroup

      for i := 0; i < 10; i++ {
         wg.Add(1)
         go write(&wg)
      }

      for i := 0; i < 10; i++ {
         wg.Add(1)
         go read(&wg)
      }

      wg.Wait()
   }
```

### Cond

A `Cond` is used for broadcast signaling between Goroutines. It's useful for implementing higher-level synchronization patterns.

#### Example:

```go
   package main

   import (
      "fmt"
      "sync"
   )

   type Queue struct {
      items []int
      cond  *sync.Cond
   }

   func (q *Queue) Enqueue(item int) {
      q.cond.L.Lock()
      q.items = append(q.items, item)
      q.cond.Signal()
      q.cond.L.Unlock()
   }

   func (q *Queue) Dequeue() int {
      q.cond.L.Lock()
      for len(q.items) == 0 {
         q.cond.Wait()
      }
      item := q.items[0]
      q.items = q.items[1:]
      q.cond.L.Unlock()
      return item
   }

   func main() {
      queue := &Queue{
         cond: sync.NewCond(&sync.Mutex{}),
      }

      var wg sync.WaitGroup
      wg.Add(2)

      go func() {
         defer wg.Done()
         for i := 0; i < 5; i++ {
            queue.Enqueue(i)
            fmt.Printf("Enqueued: %d\n", i)
         }
      }()

      go func() {
         defer wg.Done()
         for i := 0; i < 5; i++ {
            item := queue.Dequeue()
            fmt.Printf("Dequeued: %d\n", item)
         }
      }()

      wg.Wait()
   }
```

### Summary

- **WaitGroup**: Waits for a group of Goroutines to finish.
- **Mutex**: Ensures mutual exclusion, allowing only one Goroutine to access a critical section.
- **RWMutex**: Allows multiple readers or one writer.
- **Cond**: Used for signaling and broadcasting conditions between Goroutines.

These synchronization tools help you manage concurrent Goroutines and ensure safe access to shared resources.

## Channels

In Go, channels are powerful concurrency primitives that allow Goroutines to communicate with each other and synchronize their execution. They enable safe data exchange between Goroutines without the need for explicit locks or condition variables.

### Basics of Channels

- **Declaration**: Channels are declared using the `chan` keyword.
- **Types**: Channels are typed, meaning you must specify the type of data the channel will transport.
- **Operations**: The two main operations on channels are sending data to a channel and receiving data from a channel.

### Creating and Using Channels

#### Declaring a Channel

```go
var ch chan int
```

#### Initializing a Channel

```go
ch = make(chan int)
```

#### Sending Data to a Channel

```go
ch <- 42
```

#### Receiving Data from a Channel

```go
value := <-ch
```

#### Example: Basic Channel Usage

```go
   package main

   import (
      "fmt"
   )

   func main() {
      ch := make(chan int)

      go func() {
         ch <- 42
      }()

      value := <-ch
      fmt.Println(value) // Output: 42
   }
```

### Buffered vs Unbuffered Channels

- **Unbuffered Channels**: Block until both the sender and receiver are ready.
- **Buffered Channels**: Have a capacity and allow sending or receiving a fixed number of elements without blocking.

#### Buffered Channel Example

```go
   package main

   import (
      "fmt"
   )

   func main() {
      ch := make(chan int, 2)

      ch <- 1
      ch <- 2

      fmt.Println(<-ch) // Output: 1
      fmt.Println(<-ch) // Output: 2
   }
```

### Closing Channels

Closing a channel indicates that no more values will be sent on it. This is useful to signal completion.

```go
   package main

   import (
      "fmt"
   )

   func main() {
      ch := make(chan int, 2)

      ch <- 1
      ch <- 2
      close(ch)

      for value := range ch {
         fmt.Println(value)
      }
   }
```

### Select Statement

The `select` statement allows a Goroutine to wait on multiple communication operations. It blocks until one of its cases can proceed, then it executes that case.

#### Example: Using Select

```go
   package main

   import (
      "fmt"
      "time"
   )

   func main() {
      ch1 := make(chan string)
      ch2 := make(chan string)

      go func() {
         time.Sleep(1 * time.Second)
         ch1 <- "one"
      }()

      go func() {
         time.Sleep(2 * time.Second)
         ch2 <- "two"
      }()

      for i := 0; i < 2; i++ {
         select {
         case msg1 := <-ch1:
            fmt.Println("Received", msg1)
         case msg2 := <-ch2:
            fmt.Println("Received", msg2)
         }
      }
   }
```

### Directional Channels

Channels can be directional, meaning they can be restricted to send-only or receive-only. This is useful for enforcing communication patterns.

#### Example: Directional Channels

```go
   package main

   import "fmt"

   // sendOnly is a channel that can only send data
   func sendOnly(ch chan<- int) {
      ch <- 42
   }

   // receiveOnly is a channel that can only receive data
   func receiveOnly(ch <-chan int) int {
      return <-ch
   }

   func main() {
      ch := make(chan int)

      go sendOnly(ch)
      fmt.Println(receiveOnly(ch)) // Output: 42
   }
```

### Summary

- **Channels** enable safe communication and synchronization between Goroutines.
- **Unbuffered Channels** block until both sending and receiving operations are ready.
- **Buffered Channels** allow a fixed number of values to be sent without blocking.
- **Closing Channels** signals that no more values will be sent.
- **Select Statement** allows waiting on multiple channel operations.
- **Directional Channels** can be restricted to send-only or receive-only to enforce communication patterns.

Using channels effectively can help you build robust concurrent programs in Go.