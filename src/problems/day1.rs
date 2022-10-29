use crate::problems::Problem;

pub struct Day1 {
    input: Vec<i64>,
    solution1: u64,
    solution2: u64,
}

impl Day1 {
    pub fn new() -> Day1 {
        Day1 {
            input: Vec::new(),
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day1 {
    fn setup(&mut self) {
        // self.input = crate::lines_to_numbers(&crate::read_lines("input-data/001-test.txt"));
        // self.solution1 = 0;
        // self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("001 - The Start")
    }

    fn solve_problem1(&mut self) {
        // for i in 1..(self.input.len()) {
        //     if self.input[i-1] < self.input[i] {
        //         self.solution1 += 1;
        //     }
        // }
    }
    fn solve_problem2(&mut self) {}

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
