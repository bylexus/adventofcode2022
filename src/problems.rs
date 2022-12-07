
pub trait Problem {

	fn title(&self) -> String;
	fn setup(&mut self);
	fn solve_problem1(&mut self);
	fn solve_problem2(&mut self);
	fn solution_problem1(&self) -> String;
	fn solution_problem2(&self) -> String;
}

mod day01;
pub use day01::Day01;

mod day02;
pub use day02::Day02;

mod day03;
pub use day03::Day03;

mod day04;
pub use day04::Day04;

mod day05;
pub use day05::Day05;

mod day06;
pub use day06::Day06;

mod day07;
pub use day07::Day07;

mod day_2021_01;
pub use day_2021_01::Day2021_01;

mod day_2021_02;
pub use day_2021_02::Day2021_02;

