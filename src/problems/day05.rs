use std::{
    collections::{HashMap, VecDeque},
    slice::SliceIndex,
};

use regex::Regex;

use crate::problems::Problem;

#[derive(Debug)]
struct Instr {
    amount: usize,
    from: usize,
    to: usize,
}

pub struct Day05 {
    stacks: HashMap<usize, VecDeque<char>>,
    stacks2: HashMap<usize, VecDeque<char>>,
    instructions: Vec<Instr>,
    solution1: String,
    solution2: String,
}

impl Day05 {
    pub fn new() -> Day05 {
        Day05 {
            stacks: HashMap::new(),
            stacks2: HashMap::new(),
            instructions: Vec::new(),
            solution1: String::new(),
            solution2: String::new(),
        }
    }
}

impl Problem for Day05 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/05-test.txt");
        let lines = crate::read_lines("input-data/05-data.txt");
        // find empty line:
        let mut idx = 0;
        for i in &lines {
            if i.trim() == "" {
                break;
            }
            idx += 1;
        }
        idx -= 1;
        // println!("idx: {}", idx);
        // count nr of stacks: last line
        let nr_of_stacks = (lines[idx].len() + 1) / 4;
        // println!("Nr of stacks: {}", nr_of_stacks);

        for i in 1..=nr_of_stacks {
            self.stacks.insert(i, VecDeque::new());
            self.stacks2.insert(i, VecDeque::new());
        }

        // extract stacks
        for line_ids in 0..idx {
            let line = lines.get(line_ids).unwrap();
            for stack_idx in 0..nr_of_stacks {
                let chars: Vec<char> = line.chars().collect();
                let chr = chars.get(stack_idx * 4 + 1).unwrap();
                if *chr != '\u{0020}' {
                    self.stacks
                        .get_mut(&(stack_idx + 1))
                        .unwrap()
                        .push_back(*chr);
                    self.stacks2
                        .get_mut(&(stack_idx + 1))
                        .unwrap()
                        .push_back(*chr);
                }
            }
        }
        // println!("{:?}", self.stacks);

        // part 2: instructions
        idx += 2;
        let re = Regex::new(r"move (\d+) from (\d+) to (\d+)").unwrap();
        for i in idx..lines.len() {
            let line = lines.get(i).unwrap();
            let groups = match re.captures(&line) {
                Some(g) => g,
                None => continue,
            };
            self.instructions.push(Instr {
                amount: str::parse(groups.get(1).unwrap().as_str()).unwrap(),
                from: str::parse(groups.get(2).unwrap().as_str()).unwrap(),
                to: str::parse(groups.get(3).unwrap().as_str()).unwrap(),
            });
        }
        // println!("Instructions: {:?}", self.instructions);

        // let lines = crate::read_lines("input-data/05-data.txt");

        crate::split_lines(&lines, " ").iter().for_each(|parts| {});
        self.solution1 = String::new();
        self.solution2 = String::new();
    }

    fn title(&self) -> String {
        String::from("05 - Supply Stacks")
    }

    fn solve_problem1(&mut self) {
        for instr in &self.instructions {
            // println!("\n\nStacks: {:?}", self.stacks);
            // println!("Instr: {:?}", instr);
            for _ in 0..instr.amount {
                let item = self
                    .stacks
                    .get_mut(&instr.from)
                    .unwrap()
                    .pop_front()
                    .unwrap();
                self.stacks.get_mut(&instr.to).unwrap().push_front(item);
            }
            // println!("Stacks after: {:?}", self.stacks);
        }

        for i in 1..=(self.stacks.len()) {
            self.solution1
                .push(*self.stacks.get(&i).unwrap().get(0).unwrap());
        }
    }

    fn solve_problem2(&mut self) {
        let mut tmp: VecDeque<char> = VecDeque::with_capacity(100);
        for instr in &self.instructions {
            tmp.clear();
            // println!("\n\nStacks: {:?}", self.stacks2);
            // println!("Instr: {:?}", instr);
            for _ in 0..instr.amount {
                let item = self
                    .stacks2
                    .get_mut(&instr.from)
                    .unwrap()
                    .pop_front()
                    .unwrap();
                tmp.push_back(item);
            }
            for i in tmp.iter().rev() {
                self.stacks2.get_mut(&instr.to).unwrap().push_front(*i);
            }
            // println!("Stacks after: {:?}", self.stacks2);
        }

        for i in 1..=(self.stacks2.len()) {
            self.solution2
                .push(*self.stacks2.get(&i).unwrap().get(0).unwrap());
        }
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
