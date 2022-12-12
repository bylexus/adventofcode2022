use regex::Regex;

use crate::problems::Problem;

struct Section {
    begin: u64,
    end: u64,
}

impl Section {
    fn contains(&self, other: &Section) -> bool {
        (self.begin >= other.begin && self.end <= other.end)
            || (other.begin >= self.begin && other.end <= self.end)
    }
    fn overlaps(&self, other: &Section) -> bool {
        (self.begin <= other.end && self.end >= other.begin)
            || (other.begin <= self.end && other.end >= self.begin)
    }
}

pub struct Day04 {
    input: Vec<(Section, Section)>,
    solution1: usize,
    solution2: usize,
}

impl Day04 {
    pub fn new() -> Day04 {
        Day04 {
            input: Vec::new(),
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day04 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/04-test.txt");
        let lines = crate::read_lines("input-data/04-data.txt");
        let re = Regex::new(r"(\d+)-(\d+),(\d+)-(\d+)").unwrap();
        crate::split_groups(&lines, &re).iter().for_each(|groups| {
            self.input.push((
                Section {
                    begin: str::parse::<u64>(groups.get(1).unwrap().as_str()).unwrap(),
                    end: str::parse::<u64>(groups.get(2).unwrap().as_str()).unwrap(),
                },
                Section {
                    begin: str::parse::<u64>(groups.get(3).unwrap().as_str()).unwrap(),
                    end: str::parse::<u64>(groups.get(4).unwrap().as_str()).unwrap(),
                },
            ));
        });
        // let lines = crate::read_lines("input-data/04-data.txt");

        // crate::split_lines(&lines, " ").iter().for_each(|parts| {
        // });
        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("04 - Camp Cleanup")
    }

    fn solve_problem1(&mut self) {
        self.solution1 = self
            .input
            .iter()
            .filter(|entry| entry.0.contains(&entry.1) || entry.1.contains(&entry.0))
            .count();
    }
    fn solve_problem2(&mut self) {
        self.solution2 = self
            .input
            .iter()
            .filter(|entry| entry.0.overlaps(&entry.1) || entry.1.overlaps(&entry.0))
            .count();
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
