use crate::problems::Problem;

pub struct Day01 {
    input: Vec<u64>,
    solution1: u64,
    solution2: u64,
}

impl Day01 {
    pub fn new() -> Day01 {
        Day01 {
            input: Vec::new(),
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day01 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/01-test.txt");
        let lines = crate::read_lines("input-data/01-data.txt");
        for s in lines {
            if s.trim().len() == 0 {
                self.input.push(0);
            } else {
                self.input.push(str::parse(s.as_str()).unwrap());
            }
        }
        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("01 - Calorie Counting")
    }

    fn solve_problem1(&mut self) {
        let mut elve_sum: u64 = 0;
        let mut max_sum: u64 = 0;
        for nr in self.input.iter() {
            if *nr > 0 {
                elve_sum += *nr;
            } else {
                if elve_sum > max_sum {
                    max_sum = elve_sum;
                }
                elve_sum = 0;
            }
        }
        // last one, to collect them all:
        if elve_sum > max_sum {
            max_sum = elve_sum;
        }
        self.solution1 = max_sum;
    }
    fn solve_problem2(&mut self) {
        let mut elve_sum: u64 = 0;
        let mut totals: Vec<u64> = Vec::new();

        for nr in self.input.iter() {
            if *nr > 0 {
                elve_sum += *nr;
            } else {
                totals.push(elve_sum);
                elve_sum = 0;
            }
        }
        // collect last one, too:
        totals.push(elve_sum);


        totals.sort();
        let sum = totals.iter().rev().take(3).sum();
        self.solution2 = sum;
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
