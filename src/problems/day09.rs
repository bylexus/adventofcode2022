use std::collections::{HashMap, HashSet};

use crate::problems::Problem;

#[derive(Debug, Clone, Copy, Hash, PartialEq, Eq)]
struct Point {
    x: i64,
    y: i64,
}

#[derive(Debug)]
struct Instruction {
    dir: char,
    steps: u64,
}

pub struct Day09 {
    instructions: Vec<Instruction>,
    visited: HashSet<Point>,
    dir_map: HashMap<char, Point>,
    solution1: u64,
    solution2: u64,
}

impl Day09 {
    pub fn new() -> Day09 {
        Day09 {
            instructions: Vec::new(),
            visited: HashSet::new(),
            dir_map: HashMap::new(),
            solution1: 0,
            solution2: 0,
        }
    }

    fn _print_visited(&self, head: (i64, i64), tail: (i64, i64)) {
        let mut top_left = Point { x: 0, y: 0 };
        let mut bottom_right = Point { x: 0, y: 0 };

        for point in self.visited.iter() {
            if point.x < top_left.x {
                top_left.x = point.x;
            }
            if point.y < top_left.y {
                top_left.y = point.y;
            }
            if point.x > bottom_right.x {
                bottom_right.x = point.x;
            }
            if point.y > bottom_right.y {
                bottom_right.y = point.y;
            }
        }
        top_left.x -= 2;
        top_left.y -= 2;
        bottom_right.x += 2;
        bottom_right.y += 2;

        println!("head: {:?}, tail: {:?}", head, tail);
        for y in top_left.y..=bottom_right.y {
            for x in top_left.x..=bottom_right.x {
                let value = self.visited.get(&Point { x, y });
                if (x, y) == head {
                    print!("{}", 'H');
                } else if (x, y) == tail {
                    print!("{}", 'T');
                } else {
                    print!(
                        "{}",
                        match value {
                            None => ".",
                            Some(_) => "#",
                        }
                    );
                }
            }
            println!("");
        }
        println!("");
    }

    fn walk_instructions(&mut self, rope_length: usize) -> u64 {
        self.visited = HashSet::new();
        let mut rope: Vec<Point> = vec![Point { x: 0, y: 0 }; rope_length];

        self.instructions.iter().for_each(|instr| {
            for _ in 0..instr.steps {
                // Move head first, according to direction map pointer:
                let move_ptr = self.dir_map.get(&instr.dir).unwrap();
                rope[0].x += move_ptr.x;
                rope[0].y += move_ptr.y;

                // ---- process tails, from head to toe :-) -----
                for i in 1..rope.len() {
                    let head = rope[i - 1];
                    let tail = rope.get_mut(i).unwrap();

                    // head is only 1 step away from tail: no tail move needed
                    if (head.x - tail.x).abs() <= 1 && (head.y - tail.y).abs() <= 1 {
                        continue;
                    }
                    // follow diagonally: (move up to 1 place in one step in the dir of the head)
                    tail.x += match head.x - tail.x {
                        1.. => 1,
                        0 => 0,
                        _ => -1,
                    };
                    tail.y -= match head.y - tail.y {
                        1.. => -1,
                        0 => 0,
                        _ => 1,
                    };
                }

                let tail = rope.last().unwrap();
                self.visited.insert(*tail);
            }
        });
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

        self.dir_map.insert('U', Point { x: 0, y: -1 });
        self.dir_map.insert('R', Point { x: 1, y: 0 });
        self.dir_map.insert('D', Point { x: 0, y: 1 });
        self.dir_map.insert('L', Point { x: -1, y: 0 });
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
