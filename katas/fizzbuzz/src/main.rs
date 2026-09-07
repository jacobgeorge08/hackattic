use std::io::{self, Read};

fn main() {
    let mut input = String::new();
    io::stdin()
        .read_to_string(&mut input)
        .expect("failed to read stdin");

    let mut pieces = input
        .split_whitespace()
        .map(|p| p.parse::<u32>().expect("not a number"));

    let n = pieces.next().expect("failed to read n");
    let m = pieces.next().expect("failed to read m");

    for num in n..=m {
        match (num % 3, num % 5) {
            (0, 0) => println!("FizzBuzz"),
            (0, _) => println!("Fizz"),
            (_, 0) => println!("Buzz"),
            (_, _) => println!("{num}"),
        }
    }
}
