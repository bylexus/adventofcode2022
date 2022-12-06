use std::collections::HashSet;

use crate::problems::Problem;

pub struct Day06 {
    input: String,
    solution1: usize,
    solution2: usize,
}

impl Day06 {
    pub fn new() -> Day06 {
        Day06 {
            input: String::new(),
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day06 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/06-test.txt");
        let lines = crate::read_lines("input-data/06-data.txt");
        self.input = String::from(lines.get(0).unwrap().as_str());

        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("06 - Tuning Trouble")
    }

    fn solve_problem1(&mut self) {
        let bytes = self.input.as_bytes();
        let mut unique_set = HashSet::new();
        for i in 3..bytes.len() {
            unique_set.clear();
            unique_set.insert(bytes[i]);
            unique_set.insert(bytes[i - 1]);
            unique_set.insert(bytes[i - 2]);
            unique_set.insert(bytes[i - 3]);
            if unique_set.len() == 4 {
                self.solution1 = i + 1;
                break;
            }
        }
    }
    fn solve_problem2(&mut self) {
        let bytes = self.input.as_bytes();
        let mut unique_set = HashSet::new();
        let unique_len = 14;
        let len_range:Vec<usize> = (0..unique_len).rev().collect();
        for i in (unique_len - 1)..bytes.len() {
            unique_set.clear();
            for n in len_range.iter() {
                if unique_set.contains(&bytes[i - n]) {
                    break;
                } else {
                    unique_set.insert(bytes[i - n]);
                }
            }
            if unique_set.len() == unique_len {
                self.solution2 = i + 1;
                break;
            }
        }
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
