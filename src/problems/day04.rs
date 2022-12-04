use regex::Regex;

use crate::problems::Problem;

pub struct Day04 {
    input: Vec<((u64, u64), (u64, u64))>,
    solution1: usize,
    solution2: usize,
}

impl Day04 {
    pub fn new() -> Day04 {
        Day04 {
            input: Vec::new(),
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day04 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/04-test.txt");
        let lines = crate::read_lines("input-data/04-data.txt");
        let re = Regex::new(r"(\d+)-(\d+),(\d+)-(\d+)").unwrap();
        lines.iter().for_each(|line| {
            match re.captures(line) {
                Some(groups) => self.input.push((
                    (
                        str::parse::<u64>(groups.get(1).unwrap().as_str()).unwrap(),
                        str::parse::<u64>(groups.get(2).unwrap().as_str()).unwrap(),
                    ),
                    (
                        str::parse::<u64>(groups.get(3).unwrap().as_str()).unwrap(),
                        str::parse::<u64>(groups.get(4).unwrap().as_str()).unwrap(),
                    ),
                )),
                None => return,
            };
        });
        // let lines = crate::read_lines("input-data/04-data.txt");

        // crate::split_lines(&lines, " ").iter().for_each(|parts| {
        // });
        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("04 - Camp Cleanup")
    }

    fn solve_problem1(&mut self) {
        self.solution1 = self
            .input
            .iter()
            .filter(|entry| {
                (entry.0 .0 >= entry.1 .0 && entry.0 .1 <= entry.1 .1)
                    || (entry.1 .0 >= entry.0 .0 && entry.1 .1 <= entry.0 .1)
            })
            .count();
    }
    fn solve_problem2(&mut self) {
        self.solution2 = self
            .input
            .iter()
            .filter(|entry| {
                (entry.0 .0 <= entry.1 .1 && entry.0 .1 >= entry.1 .0)
                    || (entry.1 .0 <= entry.0 .1 && entry.1 .1 >= entry.0 .0)
            })
            .count();
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
