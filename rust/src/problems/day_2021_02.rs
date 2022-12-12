use crate::problems::Problem;

pub struct Day2021_02 {
    input: Vec<(String, i64)>,
    solution1: i64,
    solution2: i64,

    x: i64,
    y: i64,
    aim: i64,
}

impl Day2021_02 {
    pub fn new() -> Day2021_02 {
        Day2021_02 {
            input: Vec::new(),
            x: 0,
            y: 0,
            aim: 0,
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day2021_02 {
    fn setup(&mut self) {
        // parsing each line into a (String, i64) into a Vec:
        self.input = crate::split_lines(&crate::read_lines("input-data/2021-002.txt"), " ")
            .iter()
            .map(|el| match el.len() == 2 {
                true => (String::from(&el[0]), str::parse::<i64>(&el[1]).unwrap()),
                false => panic!("Cannot parse"),
            })
            .collect();
        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("2021 - Day 2: Dive!")
    }

    fn solve_problem1(&mut self) {
        self.x = 0;
        self.y = 0;
        self.aim = 0;
        self.input
            .iter()
            .for_each(|(command, amount)| match command.as_str() {
                "forward" => self.x += amount,
                "up" => self.y -= amount,
                "down" => self.y += amount,
                _ => return,
            });
        self.solution1 = self.x * self.y;
    }

    fn solve_problem2(&mut self) {
        self.x = 0;
        self.y = 0;
        self.aim = 0;
        self.input
            .iter()
            .for_each(|(command, amount)| match command.as_str() {
                "forward" => {
                    self.x += amount;
                    self.y += amount * self.aim;
                }
                "up" => self.aim -= amount,
                "down" => self.aim += amount,
                _ => return,
            });
        self.solution2 = self.x * self.y;
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
