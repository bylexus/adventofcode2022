use crate::problems::Problem;

pub struct Day2021_01 {
    input: Vec<i64>,
    solution1: u64,
    solution2: i64,
}

impl Day2021_01 {
    pub fn new() -> Day2021_01 {
        Day2021_01 {
            input: Vec::new(),
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day2021_01 {
    fn setup(&mut self) {
        self.input = crate::lines_to_numbers(&crate::read_lines("input-data/2021-001.txt"));
        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("2021 - Day 1: Sonar Sweep")
    }

    fn solve_problem1(&mut self) {
        for i in 1..(self.input.len()) {
            if self.input[i - 1] < self.input[i] {
                self.solution1 += 1;
            }
        }
    }

    fn solve_problem2(&mut self) {
        let mut prev: i64 = 0;
        let mut count: i64 = 0;
        for i in 0..=(self.input.len() - 3) {
            let sum = self.input[i] + self.input[i + 1] + self.input[i + 2];
            if prev > 0 && sum > prev {
                count += 1
            }
            prev = sum;
        }
        self.solution2 = count;
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
