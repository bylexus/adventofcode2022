use std::collections::VecDeque;

use regex::Regex;

use crate::problems::Problem;

#[derive(Debug, Clone, Copy)]
enum Operator {
    Add(OpValue),
    Multiply(OpValue),
}

#[derive(Debug, Clone, Copy)]
enum OpValue {
    Number(u64),
    Old,
}

#[derive(Debug)]
struct Monkey {
    items: VecDeque<u64>,
    inspect_times: u64,
    op: Operator,
    div_test: u64,
    on_true: usize,
    on_false: usize,
}

pub struct Day11 {
    monkeys1: Vec<Monkey>,
    monkeys2: Vec<Monkey>,
    solution1: u64,
    solution2: u64,
}

impl Day11 {
    pub fn new() -> Day11 {
        Day11 {
            monkeys1: Vec::new(),
            monkeys2: Vec::new(),
            solution1: 0,
            solution2: 0,
        }
    }

    fn _print_monkeys(&self, monkeys: &Vec<Monkey>, round: usize) {
        println!("\n\nAfter Round {}", round);
        for (i, monkey) in monkeys.iter().enumerate() {
            print!("Monkey {}: ", i);
            for item in monkey.items.iter() {
                print!("{}, ", item);
            }
            println!(", Inspections: {}", monkey.inspect_times);
            println!("");
        }
    }
}

impl Problem for Day11 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/11-test.txt");
        let lines = crate::read_lines("input-data/11-data.txt");

        let mut line = 0;
        let starting_items_re = Regex::new(r"Starting items: (.*)").unwrap();
        let op_re = Regex::new(r"Operation: new = old (.) (old|\d+)").unwrap();
        let test_re = Regex::new(r"Test: divisible by (\d+)").unwrap();
        let true_re = Regex::new(r"If true: throw to monkey (\d+)").unwrap();
        let false_re = Regex::new(r"If false: throw to monkey (\d+)").unwrap();

        while line + 5 < lines.len() {
            // skip monkey number line:
            line += 1;

            // starting items:
            let mut items1: VecDeque<u64> = VecDeque::new();
            let mut items2: VecDeque<u64> = VecDeque::new();
            let si_group = starting_items_re.captures(&lines[line]).unwrap();
            si_group[1]
                .split(',')
                .map(|i| str::parse(&i.trim()).unwrap())
                .for_each(|i| {
                    items1.push_back(i);
                    items2.push_back(i);
                });
            line += 1;

            // op:
            let op_group = op_re.captures(&lines[line]).unwrap();
            let op_str = op_group[2].trim();
            let op_value = match op_str {
                "old" => OpValue::Old,
                _ => OpValue::Number(str::parse(op_str).unwrap()),
            };
            let opg1 = &op_group[1];
            let op = match opg1 {
                "*" => Operator::Multiply(op_value),
                "+" => Operator::Add(op_value),
                _ => panic!("Unknown operator"),
            };
            line += 1;

            // test:
            let test_group = test_re.captures(&lines[line]).unwrap();
            let div_by: u64 = str::parse(&test_group[1]).unwrap();
            line += 1;

            // if true:
            let true_group = true_re.captures(&lines[line]).unwrap();
            let true_monkey: usize = str::parse(&true_group[1]).unwrap();
            line += 1;

            // if false:
            let false_group = false_re.captures(&lines[line]).unwrap();
            let false_monkey: usize = str::parse(&false_group[1]).unwrap();
            line += 1;

            // end, skip line
            line += 1;

            self.monkeys1.push(Monkey {
                items: items1,
                inspect_times: 0,
                op: op,
                div_test: div_by,
                on_true: true_monkey,
                on_false: false_monkey,
            });
            self.monkeys2.push(Monkey {
                items: items2,
                inspect_times: 0,
                op: op,
                div_test: div_by,
                on_true: true_monkey,
                on_false: false_monkey,
            });
        }

        // println!("{:?}", self.monkeys);

        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("11 - Monkey in the Middle")
    }

    fn solve_problem1(&mut self) {
        for _ in 0..20 {
            for i in 0..self.monkeys1.len() {
                while self.monkeys1[i].items.len() > 0 {
                    self.monkeys1[i].inspect_times += 1;
                    let old_item = self.monkeys1[i].items.pop_front().unwrap();
                    let new_item = match &self.monkeys1[i].op {
                        Operator::Add(op_value) => {
                            old_item
                                + match op_value {
                                    OpValue::Old => old_item,
                                    OpValue::Number(nr) => *nr,
                                }
                        }
                        Operator::Multiply(op_value) => {
                            old_item
                                * match op_value {
                                    OpValue::Old => old_item,
                                    OpValue::Number(nr) => *nr,
                                }
                        }
                    } / 3;
                    if new_item % self.monkeys1[i].div_test == 0 {
                        let other_idx = self.monkeys1[i].on_true;
                        self.monkeys1[other_idx].items.push_back(new_item);
                    } else {
                        let other_idx = self.monkeys1[i].on_false;
                        self.monkeys1[other_idx].items.push_back(new_item);
                    }
                }
            }
            // self.print_monkeys(round + 1);
        }

        let mut inspections: Vec<u64> = self.monkeys1.iter().map(|m| m.inspect_times).collect();
        inspections.sort();

        self.solution1 = inspections[inspections.len() - 2] * inspections[inspections.len() - 1];
    }

    fn solve_problem2(&mut self) {
        // Main Idea: because the worry values get too large soon, we need to reduce them, but
        // in a way that does not harm the divisor checks: If we take the common divisor to reduce them (modulus),
        // we can avoid that:
        // The common divisor: We take the modulo of the common multiply of all divisors:
        // this way we can reduce the value in each round so that the single modulos still work,
        // but the number does not get too large:
        let divisor = self.monkeys2.iter().map(|m| m.div_test).reduce(|prod, item | prod*item).unwrap();

        for _ in 0..10000 {
            for i in 0..self.monkeys2.len() {
                while self.monkeys2[i].items.len() > 0 {
                    self.monkeys2[i].inspect_times += 1;
                    let old_item = self.monkeys2[i].items.pop_front().unwrap();
                    let mut new_item = match &self.monkeys2[i].op {
                        Operator::Add(op_value) => {
                            old_item
                                + match op_value {
                                    OpValue::Old => old_item,
                                    OpValue::Number(nr) => *nr,
                                }
                        }
                        Operator::Multiply(op_value) => {
                            old_item
                                * match op_value {
                                    OpValue::Old => old_item,
                                    OpValue::Number(nr) => *nr,
                                }
                        }
                    };
                    new_item %= divisor;
                    if new_item % self.monkeys2[i].div_test == 0 {
                        let other_idx = self.monkeys2[i].on_true;
                        self.monkeys2[other_idx].items.push_back(new_item);
                    } else {
                        let other_idx = self.monkeys2[i].on_false;
                        self.monkeys2[other_idx].items.push_back(new_item);
                    }
                }
            }
            // self.print_monkeys(&self.monkeys2, round + 1);
        }

        let mut inspections: Vec<u64> = self.monkeys2.iter().map(|m| m.inspect_times).collect();
        inspections.sort();

        self.solution2 = inspections[inspections.len() - 2] * inspections[inspections.len() - 1];
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
