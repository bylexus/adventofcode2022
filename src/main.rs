///
/// Advent of Code 2022 - Rust Edition
///
/// As always, I participate in Adventofcode (https://adventofcode.com/),
/// this time I will use Rust - a new year, a new language :-)
///
use adventofcode2022::problems::{Day1, Day2021_01, Problem, Day2021_02};

use std::{collections::HashMap, env, time::SystemTime};

///
/// Define problems here - set a key (e.g. the day name), and instantiate th
/// problem struct.
fn create_problems() -> HashMap<String, Box<dyn Problem>> {
    let mut problems: HashMap<String, Box<dyn Problem>> = HashMap::new();
    // define all problems:
    // Test problems
    problems.insert(String::from("2021-01"), Box::new(Day2021_01::new()));
    problems.insert(String::from("2021-02"), Box::new(Day2021_02::new()));

    // AoC 2022 problems
    problems.insert(String::from("01"), Box::new(Day1::new()));

    return problems;
}

fn main() {
    let args: Vec<String> = env::args().collect();
    let mut problems = create_problems();

    let mut running_problems: Vec<String> = problems.keys().map(|nr| format!("{}", nr)).collect();

    tannenbaum();

    if args.len() > 1 {
        running_problems.clear();
        for arg in 1..args.len() {
            let key: String = match str::parse(&args[arg]) {
                Ok(key) => key,
                Err(_) => continue,
            };
            running_problems.push(key);
        }
    }
    running_problems.sort();

    let global_start = SystemTime::now();
    for key in running_problems {
        let p = problems.get_mut(&key).expect("Oops - unknown problem.");
        println!("\n\n{}: {}", key, p.title());

        let mut start = SystemTime::now();
        p.setup();
        let mut duration = SystemTime::now().duration_since(start).unwrap();
        println!("    Setup time: took: {:?}", duration);

        print!("    Solving Problem 1... ");
        start = SystemTime::now();
        p.solve_problem1();
        duration = SystemTime::now().duration_since(start).unwrap();
        println!("took: {:?}", duration);

        print!("    Solving Problem 2... ");
        start = SystemTime::now();
        p.solve_problem2();
        duration = SystemTime::now().duration_since(start).unwrap();
        println!("took: {:?}", duration);

        println!("    \x1B[1;97mSolution\x1B[0m to Problem 1: \x1B[1;97m{}\x1B[0m", p.solution_problem1());
        println!("    \x1B[1;97mSolution\x1B[0m to Problem 2: \x1B[1;97m{}\x1B[0m", p.solution_problem2());
    }

    let global_duration = SystemTime::now().duration_since(global_start).unwrap();
    println!(
        "\n\n\x1B[0;32mDone! All in all, it took {:?}\x1B[0m",
        global_duration
    );
}

fn tannenbaum() {
    println!("
\x1B[1;97m
Advent of Code 2022
--------------------

        \x1B[1;93m*   *
         \\ /
         AoC
         -\x1B[1;91m*\x1B[1;93m-
          \x1B[1;37m|\x1B[0;32m
          *
         /*\\
        /\x1B[1;94m*\x1B[0;32m*\x1B[1;93m*\x1B[0;32m\\
       /\x1B[1;91m*\x1B[0;32m***\x1B[1;94m*\x1B[0;32m\\
      /**\x1B[1;93m*\x1B[0;32m****\\
     /**\x1B[1;94m*\x1B[0;32m***\x1B[1;91m*\x1B[0;32m**\\
    /********\x1B[1;93m*\x1B[0;32m**\\
   /**\x1B[1;91m*\x1B[0;32m*****\x1B[1;94m*\x1B[0;32m****\\
  /**\x1B[1;94m*\x1B[0;32m*\x1B[1;93m*\x1B[0;32m**********\\
 /**\x1B[1;94m*\x1B[0;32m*****\x1B[1;93m*\x1B[0;32m**\x1B[1;91m*\x1B[0;32m****\x1B[1;93m*\x1B[0;32m\\
          #
          #
       \x1B[1;97m2-0-2-2
       #######
\x1B[0m");
}
