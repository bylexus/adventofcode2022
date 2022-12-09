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
    solution1: usize,
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
}

impl Problem for Day09 {
    fn setup(&mut self) {
        let lines = crate::read_lines("input-data/09-test.txt");
        // let lines = crate::read_lines("input-data/09-data.txt");

        crate::split_lines(&lines, " ").iter().for_each(|parts| {
            if parts.len() == 2 {
                let instr = Instruction {
                    dir: parts[0].chars().nth(0).unwrap(),
                    steps: str::parse(parts[1].as_str()).unwrap(),
                };
                self.instructions.push(instr);
            }
        });
        println!("{:?}", self.instructions);
        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("09 - xxx")
    }

    fn solve_problem1(&mut self) {
        let mut head = (0, 0);
        let mut tail = (0, 0);

        // rope: head is at the FRONT (index 0),
        //       tail is at the BACK (index len()-1)
        let mut rope = Vec::new();

        for i in 0..2 {
            rope.push((0, 0));
        }

        self.visited.insert(head, 1);
        println!("");

        for instr in self.instructions.iter() {
            println!("Instr: {:?}", instr);
            for nr in 0..instr.steps {
                // Move head first:
                if instr.dir == 'U' {
                    head.1 -= 1;
                }
                if instr.dir == 'R' {
                    head.0 += 1;
                }
                if instr.dir == 'D' {
                    head.1 += 1;
                }
                if instr.dir == 'L' {
                    head.0 -= 1;
                }
                // println!("Head move");
                // self.print_visited(head, tail);

                // ---- process tail -----
                if (head.0 - tail.0).abs() <= 1 && (head.1 - tail.1).abs() <= 1 {
                    // head is only 1 step away from tail: no tail move needed
                    continue;
                }
                // head is above tail
                if (head.1 - tail.1) == -2 {
                    // follow diagonally:
                    tail.1 -= 1;
                    tail.0 += head.0 - tail.0;
                }
                // head is below tail
                else if (head.1 - tail.1) == 2 {
                    // follow diagonally:
                    tail.1 += 1;
                    tail.0 += head.0 - tail.0;
                }

                // head is right of tail
                if (head.0 - tail.0) == 2 {
                    // follow diagonally:
                    tail.0 += 1;
                    tail.1 += head.1 - tail.1;
                }
                // head is left of tail
                else if (head.0 - tail.0) == -2 {
                    // follow diagonally:
                    tail.0 -= 1;
                    tail.1 += head.1 - tail.1;
                }

                if !self.visited.contains_key(&tail) {
                    self.visited.insert(tail, 0);
                }
                self.visited.entry(tail).and_modify(|entry| {
                    *entry += 1;
                });

                // println!("Tail move");
                // self.print_visited(head, tail);
            }
        }
        self.print_visited(head, tail);
        self.solution1 = self.visited.len();
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
