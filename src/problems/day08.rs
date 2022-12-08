use crate::problems::Problem;

struct Data {
    map: Vec<Vec<u8>>,
}

pub struct Day08 {
    data: Data,
    solution1: u64,
    solution2: u64,
}

impl Day08 {
    pub fn new() -> Day08 {
        Day08 {
            data: Data { map: Vec::new() },
            solution1: 0,
            solution2: 0,
        }
    }

    fn is_visible(&self, tree: u8, row_idx: usize, col_idx: usize) -> bool {
        let map = &self.data.map;
        let mut visible = true;

        // check left border to tree:
        for act_col in 0..col_idx {
            if map[row_idx][act_col] >= tree {
                visible = false;
                break;
            }
        }
        if visible == true {
            return true;
        }

        // check tree to right border:
        visible = true;
        for act_col in (col_idx + 1)..map[row_idx].len() {
            if map[row_idx][act_col] >= tree {
                visible = false;
                break;
            }
        }
        if visible == true {
            return true;
        }

        // check top border to tree:
        visible = true;
        for act_row in 0..row_idx {
            if map[act_row][col_idx] >= tree {
                visible = false;
                break;
            }
        }
        if visible == true {
            return true;
        }

        // check tree to bottom border:
        visible = true;
        for act_row in (row_idx + 1)..map.len() {
            if map[act_row][col_idx] >= tree {
                visible = false;
                break;
            }
        }
        if visible == true {
            return true;
        }

        return false;
    }

    fn scenic_score(&self, row_idx: usize, col_idx: usize) -> u64 {
        let map = &self.data.map;
        let tree = map[row_idx][col_idx];
        let mut score: u64 = 1;
        let mut local_score = 0;

        // check from tree to the left:
        for act_col in (0..col_idx).rev() {
            local_score += 1;
            if map[row_idx][act_col] >= tree {
                break;
            }
        }
        score *= local_score;
        local_score = 0;

        // check from tree to the right:
        for act_col in (col_idx + 1)..map[row_idx].len() {
            local_score += 1;
            if map[row_idx][act_col] >= tree {
                break;
            }
        }
        score *= local_score;
        local_score = 0;

        // check from tree to the top:
        for act_row in (0..row_idx).rev() {
            local_score += 1;
            if map[act_row][col_idx] >= tree {
                break;
            }
        }
        score *= local_score;
        local_score = 0;

        // check from tree to the bottom:
        for act_row in (row_idx + 1)..map.len() {
            local_score += 1;
            if map[act_row][col_idx] >= tree {
                break;
            }
        }
        score *= local_score;

        return score;
    }
}

impl Problem for Day08 {
    fn setup(&mut self) {
        // let lines = crate::read_lines("input-data/08-test.txt");
        let lines = crate::read_lines("input-data/08-data.txt");
        lines.iter().for_each(|line| {
            if line.trim().len() == 0 {
                return;
            }
            let mut row: Vec<u8> = Vec::with_capacity(line.len());
            for b in line.as_bytes() {
                row.push(b - 48);
            }
            self.data.map.push(row);
        });
        // println!("{:?}", self.data.map);

        self.solution1 = 0;
        self.solution2 = 0;
    }

    fn title(&self) -> String {
        String::from("08 - Treetop Tree House")
    }

    fn solve_problem1(&mut self) {
        // start with counting the top/bottom row of trees:
        let mut nr_of_trees:u64 = (self.data.map[0].len()*2) as u64;

        // Process inset trees:
        for row_idx in 1..(self.data.map.len()-1) {
            let row = &self.data.map[row_idx];
            // count the edge trees per row:
            nr_of_trees += 2;

            // process each tree within the row:
            for col_idx in 1..(row.len()-1) {
                let tree = row[col_idx];
                if self.is_visible(tree, row_idx, col_idx) {
                    nr_of_trees += 1;
                }
            }
        }
        self.solution1 = nr_of_trees;
    }
    fn solve_problem2(&mut self) {
        let mut scenic_scores = Vec::new();

        for row_idx in 1..(self.data.map.len()-1) {
            let row = &self.data.map[row_idx];

            // process each tree within the row:
            for col_idx in 1..(row.len()-1) {
                scenic_scores.push(self.scenic_score(row_idx, col_idx));
            }
        }
        self.solution2 = *scenic_scores.iter().max().unwrap();
    }

    fn solution_problem1(&self) -> String {
        String::from(format!("{}", self.solution1))
    }

    fn solution_problem2(&self) -> String {
        String::from(format!("{}", self.solution2))
    }
}
