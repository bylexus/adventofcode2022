use crate::problems::Problem;

pub struct Day02 {
    input: Vec<(char, char)>,
    solution1: u64,
    solution2: u64,
}

impl Day02 {
    pub fn new() -> Day02 {
        Day02 {
            input: Vec::new(),
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day02 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/02-test.txt");
        let lines = crate::read_lines("input-data/02-data.txt");
        self.input = crate::split_lines(&lines, " ")
            .iter()
            .map(|pair| {
                let mut draw = ('_', '_');
                if pair[0] == "A" {
                    draw.0 = 'r';
                } else if pair[0] == "B" {
                    draw.0 = 'p';
                } else if pair[0] == "C" {
                    draw.0 = 's';
                }
                if pair[1] == "X" {
                    draw.1 = 'r';
                } else if pair[1] == "Y" {
                    draw.1 = 'p';
                } else if pair[1] == "Z" {
                    draw.1 = 's';
                }

                return draw;
            })
            .filter(|pair| pair.0 != '_' && pair.1 != '_')
            .collect();
        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("02 - Rock Paper Scissors")
    }

    fn solve_problem1(&mut self) {
        let mut score = 0;
        for draw in &self.input {
            let draw = *draw;
            // draw score:
            if draw.1 == 'r' {
                score += 1;
            }
            if draw.1 == 'p' {
                score += 2;
            }
            if draw.1 == 's' {
                score += 3;
            }
            // If I win:
            if draw.0 == 's' && draw.1 == 'r' {
                score += 6;
            } else if draw.0 == 'r' && draw.1 == 'p' {
                score += 6;
            } else if draw.0 == 'p' && draw.1 == 's' {
                score += 6;
            // a draw:
            } else if draw.0 == draw.1 {
                score += 3;
            }
        }
        self.solution1 = score;
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
