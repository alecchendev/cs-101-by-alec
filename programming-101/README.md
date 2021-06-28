# Programming 101

*Status: To do, drafted*

## Motivation

Introduction to real programming. Get exposed to the basic concepts of a programming language and how development works for a functional, compiled language. Why Golang? It is simple, garbage-collected, statically typed, compiled, and used for various use cases such as backend web development, microservices, and distributed systems. When being introduced to programming, I don't want to abstract away too much, i.e. I think they should see the general structure of a package, the process of compiling and running your code, and understanding data types. I don't think they should need to worry about memory management or even object orientation. Golang fit the role, and I think it's cool to learn something different from the typical python javascript stack beginners are into nowadays. Also I just want to learn Go myself so it was convenient to incorporate it into this...

## Topics

- Basic command line
- Basic programming
- Make a command line program

## Content

- Basic command line - git bash
    - cd - directory, .., ., absolute paths
    - pwd
    - ls - alone, -a, directory
    - touch filename
    - mkdir filename
    - rm - file, empty folder, -r, -f, -rf
    - mv - rename, or move
- Basic functional programming in Go
    - Data types - int, string, bool
    - Variables - store and change data, var name type = value

        `var name type = value`

        `var x int = 3`

		`x := 3`

    - Conditionals - if, else, else if, while, for (while), for syntactic sugar (traditional for)

        `if expression {}`

        `if x > 4 {}`

        ```go
        if x > 4 {
        	// do something
        } else {
        	// do a different something
        }
        ```

        `for expression {}`

        `for init; expression; post {}`

        ```go
        // version 1
        var x int = 0
        for x < 5 {
        	fmt.Println(x)
        	x = x + 1
        }

        // version 2
        for x := 0; x < 5; x = x + 1 {
        	fmt.Println(x)
        }
        ```

    - Functions - container for code to run

        `func functionName(argName type, argName type) returnType {}`

        `func add(numOne int, numTwo int) int { return numOne + numTwo }`

    - Packages - collection of functions, runs code in package main
- Make a command line program
    - Show example - hangman
    - Have them create their own
        - Take the user's input
        - Use an if statement
        - Use a for loop
        - Print something

## Plan

- Show example - test, hangman
- Install git bash
- Learn command line
- Install VS Code
- Setup go module
- Learn basic programming
	- Have them look up and figure out parts of it
- Create basic command line program