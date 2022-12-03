use std::{collections::HashSet};

use crate::problems::Problem;

pub struct Day03 {
    input: Vec<String>,
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
        // let lines = crate::read_lines("input-data/03-test.txt");
        let lines = crate::read_lines("input-data/03-data.txt");
        self.input = lines;

        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("03 - Rucksack Reorganization")
    }

    fn solve_problem1(&mut self) {
        let mut sum: u64 = 0;
        for line in &self.input {
            let mut itemset: HashSet<u8> = HashSet::new();
            let bytes = line.as_bytes();
            for i in 0..(line.len() / 2) {
                itemset.insert(bytes[i]);
            }
            for i in (line.len() / 2)..line.len() {
                let c = bytes[i];
                if itemset.contains(&c) {
                    if c >= 97 {
                        // lower case:
                        sum += u64::from(c - 96);
                    } else {
                        // upper case:
                        sum += u64::from(c - 38);
                    }
                    break;
                }
            }
        }
        self.solution1 = sum;
    }

    fn solve_problem2(&mut self) {
        let mut sum: u64 = 0;
        let mut idx = 0;
        let mut itemset1: HashSet<u8> = HashSet::new();
        let mut itemset2: HashSet<u8> = HashSet::new();

        // put line 1 + 2 items into a set each,
        // then check the 3rd for entries in the two sets:
        while idx < self.input.len() - 2 {
            itemset1.clear();
            itemset2.clear();
            for c in self.input.get(idx).unwrap().as_bytes() {
                itemset1.insert(*c);
            }
            for c in self.input.get(idx + 1).unwrap().as_bytes() {
                itemset2.insert(*c);
            }
            for c in self.input.get(idx + 2).unwrap().as_bytes() {
                if itemset1.contains(&c) && itemset2.contains(&c) {
                    if *c >= 97u8 {
                        // lower case:
                        sum += u64::from(c - 96);
                    } else {
                        // upper case:
                        sum += u64::from(c - 38);
                    }
                    break;
                }
            }
            idx += 3;
        }

        self.solution2 = sum;
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
