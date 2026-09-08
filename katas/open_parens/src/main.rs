use std::io;

fn main() {
    for line in io::stdin().lines() {
        let line = line.expect("failed to read line");
        println!("{}", if is_balanced(&line) { "yes" } else { "no" });
    }
}

fn is_balanced(parens: &str) -> bool {
    let mut counter = 0;
    for c in parens.chars() {
        match c {
            '(' => counter += 1,
            ')' => counter -= 1,
            _ => (),
        }
        if counter < 0 {
            return false;
        }
    }
    counter == 0
}
