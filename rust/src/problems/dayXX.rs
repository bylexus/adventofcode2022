use crate::problems::Problem;

pub struct Day03 {
    input: Vec<u64>,
    solution1: u64,
    solution2: u64,
}

impl Day03 {
    pub fn new() -> Day03 {
        Day03 {
            input: Vec::new(),
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day03 {
    fn setup(&mut self) {
        let lines = crate::read_lines("input-data/03-test.txt");
        // let lines = crate::read_lines("input-data/03-data.txt");

        crate::split_lines(&lines, " ").iter().for_each(|parts| {
        });
        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("03 - xxx")
    }

    fn solve_problem1(&mut self) {
        self.solution1 = 0;
    }
    fn solve_problem2(&mut self) {
        self.solution2 = 0;
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
