use std::collections::HashMap;

use regex::Regex;

use crate::problems::Problem;

#[derive(Debug)]
struct Instr {
    amount: usize,
    from: usize,
    to: usize,
}

pub struct Day05 {
    stacks1: HashMap<usize, Vec<char>>,
    stacks2: HashMap<usize, Vec<char>>,
    instructions: Vec<Instr>,
    solution1: String,
    solution2: String,
}

impl Day05 {
    pub fn new() -> Day05 {
        Day05 {
            stacks1: HashMap::new(),
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
        let nr_of_stacks = (lines[idx].len() + 1) / 4;

        for i in 1..=nr_of_stacks {
            // create 2 separate storates for problem 1 + 2, as we mofify the input
            // during the solution:
            self.stacks1.insert(i, Vec::with_capacity(100));
            self.stacks2.insert(i, Vec::with_capacity(100));
        }

        // extract stacks
        for line_ids in (0..idx).rev() {
            let line = lines.get(line_ids).unwrap();
            // parse lines from bottom to up, add it to the stacks:
            for stack_idx in 0..nr_of_stacks {
                let chars: Vec<char> = line.chars().collect();
                let chr = chars.get(stack_idx * 4 + 1).unwrap();
                if *chr != '\u{0020}' {
                    self.stacks1.get_mut(&(stack_idx + 1)).unwrap().push(*chr);
                    self.stacks2.get_mut(&(stack_idx + 1)).unwrap().push(*chr);
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
    }

    fn title(&self) -> String {
        String::from("05 - Supply Stacks")
    }

    fn solve_problem1(&mut self) {
        for instr in &self.instructions {
            for _ in 0..instr.amount {
                let item = self.stacks1.get_mut(&instr.from).unwrap().pop().unwrap();
                self.stacks1.get_mut(&instr.to).unwrap().push(item);
            }
        }

        for i in 1..=(self.stacks1.len()) {
            self.solution1
                .push(*self.stacks1.get(&i).unwrap().last().unwrap());
        }
    }

    fn solve_problem2(&mut self) {
        let mut tmp: Vec<char> = Vec::with_capacity(100);
        for instr in &self.instructions {
            for _ in 0..instr.amount {
                let item = self.stacks2.get_mut(&instr.from).unwrap().pop().unwrap();
                tmp.push(item);
            }
            while tmp.len() > 0 {
                self.stacks2
                    .get_mut(&instr.to)
                    .unwrap()
                    .push(tmp.pop().unwrap());
            }
            tmp.clear();
        }

        for i in 1..=(self.stacks2.len()) {
            self.solution2
                .push(*self.stacks2.get(&i).unwrap().last().unwrap());
        }
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
