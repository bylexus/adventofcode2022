# Advent of Code 2022

Eagerly, we're all awaiting [Advent of Code, Edition 2022!](https://adventofcode.com/2022/). Finally, it's here.
What a magical time of the year!

I started with [Rust](https://rust-lang.org) this year, until Day 12: On that day, I pulled the plug:
I have to admit it, Rust is too annoying a programming language for my taste:

All in all Rust has some very cool and intriguing features like Enums with additional values, a very good
type system, no garbage collector, match expressions, no exceptions.
But all in all, the safety net for this kind of memory management - ownership and borrow checker - makes it really really
hard to do certain kind of tasks:

For example cyclic references (Parent has a list of childs, and child has a reference to the parent) are almost impossible to do.
But even small things are very hard to implement, e.g. accessing an array index while in a iterate loop over the same array.

So I lost most of the time not solving the actual puzzle, but to understand why certain code will not compile.

All those compile-time checks like object ownership, borrow checker, lifetimes etc are good things - but too limitating for
my taste.

So I switched to [GO](https://go.dev/), which is a very simple, small language, and much much easier to grasp.

Maybe I am just too stupid for Rust, who knows.

Anyway, you can find the first 11 solutions in Rust, while the rest is sovled in Go:

* Rust solutions: [./rust/](./rust/)
* GO solutions: [./go/](./go/)
