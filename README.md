# Advent of Code 2022

Eagerly, we're all awaiting [Advent of Code, Edition 2022!](https://adventofcode.com/2022/). Finally, it's here.
What a magical time of the year!

This year I choose [Rust](https://rust-lang.org) as my weapon of choice, because:

* it is new for me
* It is said to be blazzzing fast<sup>TM</sup>
* It is hyped
* I want to do something useful (hrmmm...) with the language

For such a challenge, choosing a language that is irritatingly limiting your freedom (I look at you, Type System, Single-Object-Ownership, and the likes), makes it even more fun! Using a non-limiting language is for wimps,
Rust for real bearded people only! Let's see how this turns out.

So this repository just contains my naïve solutions / tries for this year's code challenge.
Expect nothing fancy, expect ugly code, expect one-time-hacks, expect solutions that do NOT follow "The Rust Way". It's the way, my way! Don't blame me, you're warned now.

Alex

## Run problems

All problems can be run by its day index, e.g:

```
$ cargo run 01
```

or all together:

```
$ cargo run
```

## Define Problems

1) Create a struct that implements the `Problem` trait:

```rs
// src/problems/day01.rs:
use crate::problems::Problem;

pub struct Day01 {
    // .....
}

impl Day01 {
    pub fn new() -> Day01 {
        Day01 {
			// ...
        }
    }
}

impl Problem for Day01 {
    fn setup(&mut self) {
        // setup - read data, parse input etc.
    }

    fn title(&self) -> String {
        String::from("01 - Calorie Counting")
    }

    fn solve_problem1(&mut self) {
		// solve problem 1, store result e.g. in struct
    }
    fn solve_problem2(&mut self) {
		// solve problem 2, store result e.g. in struct
    }

    fn solution_problem1(&self) -> String {
		// Return a string that represents solution 1, e.g.:
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
		// Return a string that represents solution 2, e.g.:
        String::from(format!("{}", self.solution2))
    }
}
```

2) expose the type in the `problems` module:

```rs
// src/problems.rs:
mod day01;
pub use day01::Day01;
```

3) include and instantiate it in `main.rs`:

```rs
// main.rs:
use adventofcode2022::problems::{Day01};

//..
problems.insert(String::from("01"), Box::new(Day01::new()));
```
