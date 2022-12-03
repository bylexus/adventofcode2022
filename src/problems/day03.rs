use std::collections::HashSet;

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
        let mut sums: Vec<u64> = Vec::new();
        for line in &self.input {
            let mut itemset: HashSet<u8> = HashSet::new();
            let (s1, s2) = line.split_at(line.len() / 2);
            for c in Vec::from(s1) {
                itemset.insert(c);
            }
            for c in Vec::from(s2) {
                if itemset.contains(&c) {
                    if c >= 97 {
                        // lower case:
                        sums.push((c - 96).into());
                    } else {
                        // upper case:
                        sums.push((c - 38).into());
                    }
                    break;
                }
            }
        }
        self.solution1 = sums.iter().sum();
    }

    fn solve_problem2(&mut self) {
        let mut sums: Vec<u64> = Vec::new();
        let mut idx = 0;
        let mut itemset1: HashSet<u8> = HashSet::new();
        let mut itemset2: HashSet<u8> = HashSet::new();

        while idx < self.input.len()-2 {
            itemset1.clear();
            itemset2.clear();
            for c in Vec::from(self.input.get(idx).unwrap().as_str()) {
                itemset1.insert(c);
            }
            for c in Vec::from(self.input.get(idx+1).unwrap().as_str()) {
                itemset2.insert(c);
            }
            for c in Vec::from(self.input.get(idx+2).unwrap().as_str()) {
                if itemset1.contains(&c) && itemset2.contains(&c) {
                    if c >= 97 {
                        // lower case:
                        sums.push((c - 96).into());
                    } else {
                        // upper case:
                        sums.push((c - 38).into());
                    }
                    break;
                }
            }
            idx += 3;
        }

        self.solution2 = sums.iter().sum();
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
