use std::{
    fs,
    io::{BufRead, BufReader},
};

pub mod problems;

pub fn read_lines(file: &str) -> Vec<String> {
    let fh = fs::File::open(file).unwrap();
    let buffered_reader = BufReader::new(fh);
    buffered_reader.lines().map(|res| res.unwrap()).collect()
}

pub fn lines_to_numbers(lines: &Vec<String>) -> Vec<i64> {
    let lines: Vec<i64> = lines
        .iter()
        .filter(|l| l.trim().len() > 0)
        .map(|s| str::parse::<i64>(&s).unwrap())
        .collect();
    lines
}
