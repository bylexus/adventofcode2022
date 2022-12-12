use std::{
    io::{stdout, Write},
    time::Duration,
};

use crate::problems::Problem;

const CRT_WIDTH: u32 = 40;
const _CRT_HEIGHT: u32 = 6;

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
            solution1: 0,
            solution2: String::new(),
        }
    }

    fn _solve_problem2_animated(&mut self) {
        self.cpu.iptr = 0;
        self.cpu.reg_x = 1;

        println!("\n");
        println!("\n");
        while self.cpu.iptr < self.cpu.mem.len() as i64 {
            let x = self.cpu.iptr % (CRT_WIDTH as i64);
            std::thread::sleep(Duration::from_millis(30));
            if x == 0 {
                println!("");
            }
            if self.cpu.reg_x >= x - 1 && self.cpu.reg_x <= x + 1 {
                if rand::random() {
                    print!("\x1b[38;5;40m█");
                } else {
                    print!("\x1b[38;5;42m█");
                }
            } else {
                print!("\x1b[38;5;29m·");
            }
            stdout().flush().unwrap();
            self.cpu.advance();
        }
        println!("\n");
        println!("\x1B[0m\n");
    }

    fn _solve_problem2_ocr(&mut self) {
        self.cpu.iptr = 0;
        self.cpu.reg_x = 1;

        let mut img: image::RgbImage = image::ImageBuffer::new(CRT_WIDTH, _CRT_HEIGHT);

        while self.cpu.iptr < self.cpu.mem.len() as i64 {
            let x = self.cpu.iptr % (CRT_WIDTH as i64);
            let y = self.cpu.iptr / (CRT_WIDTH as i64);
            if self.cpu.reg_x >= x - 1 && self.cpu.reg_x <= x + 1 {
                img.put_pixel(x as u32, y as u32, image::Rgb { 0: [0, 0, 0] });
            } else {
                img.put_pixel(x as u32, y as u32, image::Rgb { 0: [255, 255, 255] });
            }
            self.cpu.advance();
        }
        img.save("day10-img.png").unwrap();
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
        // ******* Uncomment to get an animated version of the solution: *****************3
        // return self._solve_problem2_animated();

        // ******* Uncomment to get an OCR version of the solution: *****************3
        // return self._solve_problem2_ocr();

        self.cpu.iptr = 0;
        self.cpu.reg_x = 1;

        let mut output = String::from("\n");
        while self.cpu.iptr < self.cpu.mem.len() as i64 {
            let x = self.cpu.iptr % (CRT_WIDTH as i64);
            if x == 0 {
                output += "\n";
            }
            if self.cpu.reg_x >= x - 1 && self.cpu.reg_x <= x + 1 {
                output += "█";
            } else {
                output += " ";
            }
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
