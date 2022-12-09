use std::collections::HashMap;

use crate::problems::Problem;

/// x, y, means: Point.0 => x, Point.1 => y
type Point = (i64, i64);

#[derive(Debug)]
struct Instruction {
    dir: char,
    steps: u64,
}

pub struct Day09 {
    instructions: Vec<Instruction>,
    visited: HashMap<Point, u64>,
    solution1: u64,
    solution2: u64,
}

impl Day09 {
    pub fn new() -> Day09 {
        Day09 {
            instructions: Vec::new(),
            visited: HashMap::new(),
            solution1: 0,
            solution2: 0,
        }
    }

    fn print_visited(&self, head: (i64, i64), tail: (i64, i64)) {
        let mut top_left = (0, 0);
        let mut bottom_right = (0, 0);

        for key in self.visited.keys() {
            if key.0 < top_left.0 {
                top_left.0 = key.0;
            }
            if key.1 < top_left.1 {
                top_left.1 = key.1;
            }
            if key.0 > bottom_right.0 {
                bottom_right.0 = key.0;
            }
            if key.1 > bottom_right.1 {
                bottom_right.1 = key.1;
            }
        }
        top_left.0 -= 2;
        top_left.1 -= 2;
        bottom_right.0 += 2;
        bottom_right.1 += 2;

        println!("head: {:?}, tail: {:?}", head, tail);
        for y in top_left.1..=bottom_right.1 {
            for x in top_left.0..=bottom_right.0 {
                let value = match self.visited.get(&(x, y)) {
                    Some(n) => *n,
                    None => 0,
                };
                if (x, y) == head {
                    print!("{}", 'H');
                } else if (x, y) == tail {
                    print!("{}", 'T');
                } else {
                    print!(
                        "{}",
                        match value {
                            0 => ".",
                            1.. => "#",
                        }
                    );
                }
            }
            println!("");
        }
        println!("");
    }

    fn walk_instructions(&mut self, rope_length: usize) -> u64 {
        self.visited = HashMap::new();
        let mut rope: Vec<(i64, i64)> = Vec::new();

        for _ in 0..rope_length {
            rope.push((0, 0));
        }

        for instr in self.instructions.iter() {
            for _ in 0..instr.steps {
                // Move head first:
                if instr.dir == 'U' {
                    rope[0].1 -= 1;
                } else if instr.dir == 'R' {
                    rope[0].0 += 1;
                } else if instr.dir == 'D' {
                    rope[0].1 += 1;
                } else if instr.dir == 'L' {
                    rope[0].0 -= 1;
                }

                // ---- process tails -----
                for i in 1..rope.len() {
                    let head = rope[i - 1];
                    let mut tail = rope[i];
                    if (head.0 - tail.0).abs() <= 1 && (head.1 - tail.1).abs() <= 1 {
                        // head is only 1 step away from tail: no tail move needed
                        continue;
                    }
                    // follow diagonally: (move up to 1 place in one step in the dir of the head)
                    tail.0 += match head.0 - tail.0 {
                        1.. => 1,
                        0 => 0,
                        _ => -1,
                    };
                    tail.1 -= match head.1 - tail.1 {
                        1.. => -1,
                        0 => 0,
                        _ => 1,
                    };
                    rope[i] = tail;
                }

                let tail = rope.last().unwrap();
                if !self.visited.contains_key(tail) {
                    self.visited.insert(*tail, 0);
                }
                self.visited.entry(*tail).and_modify(|entry| {
                    *entry += 1;
                });

                // println!("Tail move");
                // self.print_visited(head, tail);
            }
        }
        // let head = rope[0];
        // let tail = rope.last().unwrap();
        // self.print_visited(head, *tail);
        return self.visited.len() as u64;
    }
}

impl Problem for Day09 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/09-test.txt");
        // let lines = crate::read_lines("input-data/09-test2.txt");
        let lines = crate::read_lines("input-data/09-data.txt");

        crate::split_lines(&lines, " ").iter().for_each(|parts| {
            if parts.len() == 2 {
                let instr = Instruction {
                    dir: parts[0].chars().nth(0).unwrap(),
                    steps: str::parse(parts[1].as_str()).unwrap(),
                };
                self.instructions.push(instr);
            }
        });
        // println!("{:?}", self.instructions);
        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("09 - Rope Bridge")
    }

    fn solve_problem1(&mut self) {
        self.solution1 = self.walk_instructions(2);
    }

    fn solve_problem2(&mut self) {
        self.solution2 = self.walk_instructions(10);
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
