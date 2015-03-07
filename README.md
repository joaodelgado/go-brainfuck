# GO-BRAINFUCK

Simple [brainfuck](http://www.muppetlabs.com/~breadbox/bf/) interpreter
written in [Go](https://golang.org/) as an exercise to get my hands dirty with
the language.

This is just an adaptation of the [post](http://howistart.org/posts/nim/1)
written by Dennis Felsing on how to start Nim projects.


## Usage

If run without arguments, it interprets the first input line:

    $ go run brainfuck.go
    ++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++.
    Hello World!

It also accepts the name of a file as the first argument:

    $ go run brainfuck.go helloworld.b
    Hello World!
