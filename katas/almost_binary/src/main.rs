use std::io;
fn main() {
    for line in io::stdin().lines() {
        let line = line.expect("Failed to read line");
        println!("{}", parse_binary(&line));
    }
}

fn parse_binary(line: &str) -> u16 {
    let mut result: u16 = 0;

    for c in line.chars() {
        result <<= 1;
        if c == '#' {
            result |= 1;
        }
    }
    result
}
