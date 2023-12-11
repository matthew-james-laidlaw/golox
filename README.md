# GoLox
![Unit Tests](https://github.com/matthew-james-laidlaw/golox/actions/workflows/unit-tests.yml/badge.svg)

An interpreter for the Lox language as described in Robert Nystrom's [Crafting Interpreters](https://craftinginterpreters.com/). Written in Go.

# Purpose
I developed this application as an exercise in learning how programming languages work under the hood and as an opportunity to get better with the Go programming language.
This application can parse the [Lox grammar](https://craftinginterpreters.com/appendix-i.html) into an [abstract syntax tree](https://craftinginterpreters.com/appendix-ii.html) and implements a tree-walk interpreter to execute Lox programs.
The tool has two modes of operation. If passed a filename, it will execute code written in that file. If nothing is passed in, the tool drives an interactive REPL.

# Test

```
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

###### Returns
```
> fun Foo() { return 42; }
> var n = Foo();
> print n;
42
```

# Examples

###### Fibonacci
This program defines a recursive function that can calculate fibonacci numbers and uses it to print the first twenty fibonacci values. 
```
// prog.lox
fun fib(n) {
  if (n <= 1) return n;
  return fib(n - 2) + fib(n - 1);
}

for (var i = 0; i < 20; i = i + 1) {
  print fib(i);
}
```
```
> go run golox/cmd/golox "path/to/prog.lox"
0
1
1
2
3
5
8
13
21
34
55
89
144
233
377
610
987
1597
2584
4181
```
