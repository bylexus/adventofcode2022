use crate::problems::Problem;

#[derive(Debug)]
struct Cpu {
    mem: Vec<Instruction>,
    iptr: i64,
    reg_x: i64,
}

impl Cpu {
    fn advance(&mut self) {
        match self.mem[self.iptr as usize] {
            Instruction::Addx(nr) => {
                self.reg_x += nr;
            }
            _ => {}
        };
        self.iptr += 1;
    }
}

#[derive(Debug)]
enum Instruction {
    Noop,
    Calculating,
    Addx(i64),
}

pub struct Day10 {
    cpu: Cpu,
    crt: Vec<Vec<u8>>,
    solution1: i64,
    solution2: String,
}

impl Day10 {
    pub fn new() -> Day10 {
        Day10 {
            cpu: Cpu {
                iptr: 0,
                mem: Vec::new(),
                reg_x: 1,
            },
            crt: Vec::new(),
            solution1: 0,
            solution2: String::new(),
        }
    }
}

impl Problem for Day10 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/10-test.txt");
        let lines = crate::read_lines("input-data/10-data.txt");

        crate::split_lines(&lines, " ").iter().for_each(|parts| {
            if parts.len() > 0 {
                if parts[0] == "noop" {
                    self.cpu.mem.push(Instruction::Noop)
                } else if parts[0] == "addx" {
                    self.cpu.mem.push(Instruction::Calculating);
                    self.cpu
                        .mem
                        .push(Instruction::Addx(str::parse(&parts[1]).unwrap()))
                }
            }
        });

        // init crt:
        for _ in 0..6 {
            self.crt.push(vec![0; 40]);
        }
    }

    fn title(&self) -> String {
        String::from("10 - Cathode-Ray Tube")
    }

    fn solve_problem1(&mut self) {
        let ptr_measure = vec![19, 59, 99, 139, 179, 219];
        let mut sum = 0;
        self.cpu.iptr = 0;
        self.cpu.reg_x = 1;

        while self.cpu.iptr < self.cpu.mem.len() as i64 {
            if ptr_measure.contains(&(self.cpu.iptr as i32)) {
                sum += (self.cpu.iptr + 1) * self.cpu.reg_x;
            }
            self.cpu.advance();
        }

        self.solution1 = sum;
    }
    fn solve_problem2(&mut self) {
        self.cpu.iptr = 0;
        self.cpu.reg_x = 1;

        let mut output = String::from("\n");
        while self.cpu.iptr < self.cpu.mem.len() as i64 {
            let y = self.cpu.iptr / 40;
            let x = self.cpu.iptr % 40;
            if self.cpu.reg_x >= x - 1 && self.cpu.reg_x <= x + 1 {
                self.crt[y as usize][x as usize] = 1;
            }
            if x == 0 {
                output += "\n";
            }
            output += match self.crt[y as usize][x as usize] {
                0 => " ",
                1 => "#",
                _ => "",
            };
            self.cpu.advance();
        }

        self.solution2 = output;
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
