use crate::problems::Problem;

#[derive(Debug, PartialEq, Eq)]
/// The Hand enum represents the possible hands: rock, paper, or scissor
/// It also implements the value function, which returns the hand's value,
/// as well as the fight_against method, which returns the result when fighting against another hand.
enum Hand {
    Rock,
    Paper,
    Scissor,
    Unknown,
}

impl Hand {
    /// Returns the value of the hand
    fn value(&self) -> u64 {
        match self {
            Self::Rock => 1,
            Self::Paper => 2,
            Self::Scissor => 3,
            Self::Unknown => 0,
        }
    }

    /// returns the result of the fight against another hand
    /// (from self point of view)
    fn fight_against(&self, opponent: &Hand) -> PlayResult {
        match self {
            Self::Rock => match opponent {
                Self::Rock => PlayResult::Draw,
                Self::Paper => PlayResult::Loose,
                Self::Scissor => PlayResult::Win,
                _ => PlayResult::INVALID,
            },
            Self::Paper => match opponent {
                Self::Rock => PlayResult::Win,
                Self::Paper => PlayResult::Draw,
                Self::Scissor => PlayResult::Loose,
                _ => PlayResult::INVALID,
            },
            Self::Scissor => match opponent {
                Self::Rock => PlayResult::Loose,
                Self::Paper => PlayResult::Win,
                Self::Scissor => PlayResult::Draw,
                _ => PlayResult::INVALID,
            },
            _ => PlayResult::INVALID,
        }
    }

    /// returns the hand needed to reach a specific result
    /// if play agains self:
    /// for example:
    /// - self is a Rock
    /// - if you want to win,
    /// - select_card(self, PlayResult::Win) will return Hand::Paper
    /// as paper wins against rock.
    fn select_card(&self, whish: &PlayResult) -> Hand {
        match self {
            Self::Rock => match whish {
                PlayResult::Win => Hand::Scissor,
                PlayResult::Draw => Hand::Rock,
                PlayResult::Loose => Hand::Paper,
                _ => Hand::Unknown,
            },
            Self::Paper => match whish {
                PlayResult::Win => Hand::Rock,
                PlayResult::Draw => Hand::Paper,
                PlayResult::Loose => Hand::Scissor,
                _ => Hand::Unknown,
            },
            Self::Scissor => match whish {
                PlayResult::Win => Hand::Paper,
                PlayResult::Draw => Hand::Scissor,
                PlayResult::Loose => Hand::Rock,
                _ => Hand::Unknown,
            },
            _ => Hand::Unknown,
        }
    }
}

#[derive(Debug, PartialEq, Eq)]
/// Represents the result of a fight hand vs hand
enum PlayResult {
    Win,
    Loose,
    Draw,
    INVALID,
}

impl PlayResult {
    /// Returns the points gained for the specific result
    fn points(&self) -> u64 {
        match self {
            Self::Win => 6,
            Self::Draw => 3,
            _ => 0,
        }
    }
}

pub struct Day02 {
    input1: Vec<(Hand, Hand)>,
    input2: Vec<(Hand, PlayResult)>,
    solution1: u64,
    solution2: u64,
}

impl Day02 {
    pub fn new() -> Day02 {
        Day02 {
            input1: Vec::new(),
            input2: Vec::new(),
            solution1: 0,
            solution2: 0,
        }
    }
}

impl Problem for Day02 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/02-test.txt");
        let lines = crate::read_lines("input-data/02-data.txt");

        // From each line, we create 2 input vectors,
        // one containing (Hand, Hand), for problem 1,
        // one containing (Hand, Strategy), for problem 2
        crate::split_lines(&lines, " ").iter().for_each(|pair| {
            let draw1 = (
                match pair[0].chars().nth(0).unwrap() {
                    'A' => Hand::Rock,
                    'B' => Hand::Paper,
                    'C' => Hand::Scissor,
                    _ => Hand::Unknown,
                },
                match pair[1].chars().nth(0).unwrap() {
                    'X' => Hand::Rock,
                    'Y' => Hand::Paper,
                    'Z' => Hand::Scissor,
                    _ => Hand::Unknown,
                },
            );
            if draw1.0 != Hand::Unknown && draw1.1 != Hand::Unknown {
                self.input1.push(draw1);
            }

            let draw2 = (
                match pair[0].chars().nth(0).unwrap() {
                    'A' => Hand::Rock,
                    'B' => Hand::Paper,
                    'C' => Hand::Scissor,
                    _ => Hand::Unknown,
                },
                match pair[1].chars().nth(0).unwrap() {
                    'X' => PlayResult::Loose,
                    'Y' => PlayResult::Draw,
                    'Z' => PlayResult::Win,
                    _ => PlayResult::INVALID,
                },
            );
            if draw2.0 != Hand::Unknown && draw2.1 != PlayResult::INVALID {
                self.input2.push(draw2);
            }
        });
        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("02 - Rock Paper Scissors")
    }

    fn solve_problem1(&mut self) {
        self.solution1 = self
            .input1
            .iter()
            .map(|draw| 
                // draw hand score:
                draw.1.value()
                
                // draw battle score:
                + draw.1.fight_against(&draw.0).points())
            .sum();
    }
    fn solve_problem2(&mut self) {
        self.solution2 = self
            .input2
            .iter()
            .map(|draw| {
                // add fight score:
                draw.1.points()
                    // add draw score:
                    // we need to select the opposite result against the whished result,
                    // as we only know the OTHER's hand:
                    // So we fight as opponent of the elve:
                    + (match draw.1 {
                        PlayResult::Draw => draw.0.select_card(&PlayResult::Draw),
                        PlayResult::Win => draw.0.select_card(&PlayResult::Loose),
                        PlayResult::Loose => draw.0.select_card(&PlayResult::Win),
                        _ => Hand::Unknown,
                    })
                    .value()
            })
            .sum();
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
