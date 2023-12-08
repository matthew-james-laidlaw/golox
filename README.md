# GoLox
![Unit Tests](https://github.com/matthew-james-laidlaw/golox/actions/workflows/unit-tests.yml/badge.svg)

An interpreter for the Lox language as described in Robert Nystrom's [Crafting Interpreters](https://craftinginterpreters.com/). Written in Go.

# Purpose
I developed this application as an exercise in learning how programming languages work under the hood and as an opportunity to get better with the Go programming language.
This application can parse the [Lox grammar](https://craftinginterpreters.com/appendix-i.html) into an [abstract syntax tree](https://craftinginterpreters.com/appendix-ii.html) and implements a tree-walk interpreter to execute Lox programs.
The tool has two modes of operation. If passed a filename, it will execute code written in that file. If nothing is passed in, the tool drives an interactive REPL.

# Build & Test

```
> go build golox/cmd/golox
> go test -v golox/pkg/lox
```

# Run

```
# script (file extension doesn't matter)
> go run golox/cmd/golox "/path/to/something.lox"

# REPL
> go run golox/cmd/golox
> ...
```

# Usage

###### Hello World
```
> print "Hello, World!";
Hello, World!
```

###### Variable Declarations
```
> var a = 42;
> print a;
42
```

###### If Statements
```
> var a = 42;
> if (a == 42) { print "Good!"; } else { print "Bad"; }
Good!
```

###### While Loops
```
> var a = 0
> while (a <= 3) { print a; a = a + 1; }
0
1
2
3
```

###### For Loops
```
> for (var a = 0; a <= 3; a = a + 1) { print a; }
0
1
2
3
```

###### Functions
```
> fun Hello() { print "Hello, World!"; }
> Hello();
Hello, World!
```

# Examples

###### Iterative Fibonacci
A terrible iterative algorithm for calculating fibonacci numbers. This implementation of Lox does not yet support returning from functions.
Thus, recursion is impossible and I must use this as a proxy.
```
// prog.lox
fun Fibonacci(n) {
  var a = 0;
  var b = 1;
  for (var i = 1; i < n; i = i + 1) {
    var temp = a + b;
    a = b;
    b = temp;  
  }
  print b;
}

Fibonacci(10);
```
```
> go run golox/cmd/golox "path/to/prog.lox"
55
```
