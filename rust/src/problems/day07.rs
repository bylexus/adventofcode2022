use std::{borrow::Borrow, cell::RefCell, rc::Rc};

use regex::Regex;

use crate::problems::Problem;

#[derive(Debug)]
struct DirInfo {
    name: String,
    parent_dir: Option<Rc<RefCell<DirInfo>>>,
    child_dirs: Vec<Rc<RefCell<DirInfo>>>,
    child_files: Vec<FileInfo>,
}

impl DirInfo {
    fn get_parent_dir(&mut self) -> Option<Rc<RefCell<DirInfo>>> {
        match self.parent_dir.borrow() {
            Some(d) => Some(Rc::clone(&d)),
            None => None,
        }
    }
    fn find_dir(&self, dir_name: &str) -> Option<Rc<RefCell<DirInfo>>> {
        for d in self.child_dirs.iter() {
            let dir = d.as_ref().borrow();
            if dir.name == dir_name {
                return Some(Rc::clone(d));
            }
        }
        return None;
    }

    fn _print(&self, indent: usize) {
        let mut spaces = String::new();
        for _ in 0..indent {
            spaces += " ";
        }

        // println!("{} - {}", spaces, self.name);
        for d in self.child_dirs.iter() {
            let dir = d.as_ref().borrow();
            dir._print(indent + 2);
        }
        for f in self.child_files.iter() {
            f._print(indent + 2);
        }
    }

    fn calc_size(&self) -> u64 {
        let mut sum = 0;
        for d in self.child_dirs.iter() {
            let dir = d.as_ref().borrow();
            sum += dir.calc_size();
        }
        for f in self.child_files.iter() {
            sum += f.size;
        }
        return sum;
    }
}

#[derive(Debug)]
struct FileInfo {
    _name: String,
    size: u64,
}

impl FileInfo {
    fn _print(&self, indent: usize) {
        let mut spaces = String::new();
        for _ in 0..indent {
            spaces += " ";
        }

        println!("{} - {}", spaces, self._name);
    }
}

struct DataStore {
    root_dir: Option<Rc<RefCell<DirInfo>>>,
    dirs: Vec<Rc<RefCell<DirInfo>>>,
}

impl DataStore {
    fn parse_input(&mut self, lines: &Vec<String>) {
        self.root_dir = Some(Rc::new(RefCell::new(DirInfo {
            name: String::from("/"),
            parent_dir: None,
            child_dirs: Vec::new(),
            child_files: Vec::new(),
        })));
        self
            .dirs
            .push(Rc::clone(self.root_dir.as_ref().unwrap()));
        let mut act_dir = Rc::clone(self.root_dir.as_ref().unwrap());
        let file_match = Regex::new(r"^(\d+) (.*)").unwrap();

        for line in lines.iter() {
            if line.trim().len() == 0 {
                continue;
            }
            // command found:
            if line.chars().nth(0).unwrap() == '$' {
                let cmd = line.get(2..4).unwrap();
                if cmd == "cd" {
                    // change dir: cd
                    let dir = line.get(5..).unwrap();
                    if dir == "/" {
                        // root dir found:
                        act_dir = Rc::clone(self.root_dir.as_ref().unwrap());
                    } else if dir == ".." {
                        // go one up:
                        let ad = match act_dir.as_ref().borrow_mut().get_parent_dir() {
                            Some(d) => Rc::clone(&d),
                            None => Rc::clone(self.root_dir.as_ref().unwrap()),
                        };
                        act_dir = ad;
                    } else {
                        // cd into subdir
                        let nd = act_dir.as_ref().borrow().find_dir(dir);
                        if let Some(d) = nd {
                            act_dir = Rc::clone(&d);
                        }
                    }
                }
            } else if line.get(0..3).unwrap() == "dir" {
                // found dir, attach to act dir
                let d = DirInfo {
                    name: String::from(line.get(4..).unwrap()),
                    parent_dir: Some(Rc::clone(&act_dir)),
                    child_dirs: Vec::new(),
                    child_files: Vec::new(),
                };
                let dref = Rc::new(RefCell::new(d));
                self.dirs.push(Rc::clone(&dref));
                act_dir.borrow_mut().child_dirs.push(dref);
            }
            if let Some(groups) = file_match.captures(&line) {
                // found file, attach to act dir:
                let size: u64 = str::parse(groups.get(1).unwrap().as_str()).unwrap();
                let name = groups.get(2).unwrap().as_str();
                let f = FileInfo {
                    _name: String::from(name),
                    size: size,
                };
                act_dir.borrow_mut().child_files.push(f);
            }
        }

    }
}

pub struct Day07 {
    data: DataStore,
    solution1: u64,
    solution2: u64,
}

impl Day07 {
    pub fn new() -> Day07 {
        Day07 {
            data: DataStore {
                root_dir: None,
                dirs: Vec::new(),
            },
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day07 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/07-test.txt");
        let lines = crate::read_lines("input-data/07-data.txt");
        self.data.parse_input(&lines);

        // self.root_dir.as_ref().unwrap().as_ref().borrow().print(0);

        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("07 - No Space Left On Device")
    }

    fn solve_problem1(&mut self) {
        let mut sum = 0;
        for d in self.data.dirs.iter() {
            let dir = d.as_ref().borrow();
            let size = dir.calc_size();
            if size <= 100000 {
                sum += size;
            }
        }

        self.solution1 = sum;
    }
    fn solve_problem2(&mut self) {
        let mut dirs: Vec<u64> = Vec::new();
        let total = self
            .data
            .root_dir
            .as_ref()
            .unwrap()
            .as_ref()
            .borrow()
            .calc_size();
        let to_delete = total - 40000000;

        for d in self.data.dirs.iter() {
            let dir = d.as_ref().borrow();
            let size = dir.calc_size();
            if size >= to_delete {
                dirs.push(size);
            }
        }
        dirs.sort();

        let smallest = dirs.iter().nth(0).unwrap();
        self.solution2 = *smallest;
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
