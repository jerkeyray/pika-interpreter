# Pika Programming Language

```
       .__ __
______ |__|  | _______
\____ \|  |  |/ /\__  \
|  |_> >  |    <  / __ \_
|   __/|__|__|_ \(____  /
|__|           \/     \/
```

Pika is a tree-walking interpreter for a dynamically typed programming language written in Go. It features a complete lexer, parser, and evaluator supporting multiple data types, functions, closures, and built-in operations.

## Features

- **Mathematical expressions** with standard operators (+, -, \*, /, <, >, ==, !=)
- **Variable bindings** using `let` statements
- **Functions** with parameters and closures
- **Conditionals** with if/else expressions
- **Return statements** for early function exits
- **Higher-order functions** supporting functional programming patterns
- **Lexical scoping** with proper environment handling
- **Multiple data types**: integers, booleans, strings, arrays, and hash maps
- **Built-in functions** for common operations

## Data Types

### Integers

```javascript
let x = 42;
let y = -10;
```

### Booleans

```javascript
let isTrue = true;
let isFalse = false;
```

### Strings

```javascript
let greeting = "Hello, World!";
let name = "Pika";
```

### Arrays

```javascript
let numbers = [1, 2, 3, 4, 5];
let mixed = [1, "hello", true, [1, 2]];
```

### Hash Maps

```javascript
let person = { name: "Alice", age: 30, active: true };
let empty = {};
```

## Language Syntax

### Variable Declaration

```javascript
let x = 10;
let name = "Pika";
let isActive = true;
```

### Functions

```javascript
// Function declaration
let add = fn(x, y) {
    return x + y;
};

// Function call
let result = add(5, 3);

// Closures
let makeCounter = fn() {
    let count = 0;
    return fn() {
        count = count + 1;
        return count;
    };
};

let counter = makeCounter();
counter(); // 1
counter(); // 2
```

### Conditionals

```javascript
let max = fn(x, y) {
    if (x > y) {
        return x;
    } else {
        return y;
    }
};
```

### Array Operations

```javascript
let arr = [1, 2, 3];
arr[0]; // Access first element
arr[1] = 10; // Arrays are mutable through indexing
```

### Hash Map Operations

```javascript
let person = { name: "Bob", age: 25 };
person["name"]; // Access value by key
person["city"] = "New York"; // Add new key-value pair
```

## Built-in Functions

### `len(object)`

Returns the length of strings and arrays.

```javascript
len("hello"); // 5
len([1, 2, 3]); // 3
```

### `first(array)`

Returns the first element of an array.

```javascript
first([1, 2, 3]); // 1
```

### `last(array)`

Returns the last element of an array.

```javascript
last([1, 2, 3]); // 3
```

### `rest(array)`

Returns a new array with all elements except the first.

```javascript
rest([1, 2, 3]); // [2, 3]
```

### `push(array, element)`

Returns a new array with the element appended.

```javascript
push([1, 2], 3); // [1, 2, 3]
```

### `print(...args)`

Prints values to stdout.

```javascript
print("Hello", 42, true);
```

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:

```bash
git clone https://github.com/jerkeyray/pika-interpreter.git
cd pika-interpreter
```

2. Navigate to the pika directory:

```bash
cd pika
```

3. Run the interpreter:

```bash
go run main.go
```

### Using the REPL

The Pika REPL (Read-Eval-Print Loop) provides an interactive environment:

```
>> let x = 10;
>> let y = 20;
>> x + y
30
>> let greet = fn(name) { return "Hello, " + name; };
>> greet("World")
Hello, World
```

## Project Structure

```
pika/
├── ast/           # Abstract Syntax Tree definitions
├── evaluator/     # Expression evaluation and built-ins
├── lexer/         # Tokenization of source code
├── object/        # Object system and environment
├── parser/        # Recursive descent parser
├── repl/          # Read-Eval-Print Loop
├── token/         # Token definitions and keywords
└── main.go        # Entry point
```

## Examples

### Fibonacci Function

```javascript
let fibonacci = fn(n) {
    if (n < 2) {
        return n;
    } else {
        return fibonacci(n - 1) + fibonacci(n - 2);
    }
};

fibonacci(10); // 55
```

### Array Processing

```javascript
let map = fn(arr, f) {
    let iter = fn(arr, accumulated) {
        if (len(arr) == 0) {
            return accumulated;
        } else {
            return iter(rest(arr), push(accumulated, f(first(arr))));
        }
    };
    return iter(arr, []);
};

let double = fn(x) { return x * 2; };
map([1, 2, 3, 4], double); // [2, 4, 6, 8]
```

### Working with Hash Maps

```javascript
let person = {"name": "Alice", "age": 30};
let getName = fn(p) { return p["name"]; };
let incrementAge = fn(p) {
    p["age"] = p["age"] + 1;
    return p;
};

getName(person); // "Alice"
```

## Implementation Details

Pika uses a tree-walking interpreter architecture:

1. **Lexer**: Converts source code into tokens
2. **Parser**: Builds an Abstract Syntax Tree (AST) using recursive descent parsing
3. **Evaluator**: Walks the AST and evaluates expressions in the given environment

The interpreter supports proper lexical scoping through environment chaining and implements closures by capturing the defining environment within function objects.

## Current Limitations

- Only supports integers (no floats)
- No postfix operators
- Basic error handling
- Limited standard library

## License

This project is part of a learning exercise in building interpreters.

## License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for details.
