use std::io::{self, Read};

fn main() {
    let mut input = String::new();
    io::stdin()
        .read_to_string(&mut input)
        .expect("failed to read stdin");

    for line in input.lines() {
        println!("{}", decode(line))
    }
}

fn decode(line: &str) -> String {
    let mut acc: u32 = 0;
    let mut bits: u32 = 0;
    let mut out: Vec<u8> = Vec::new();

    for &b in line.as_bytes() {
        if b == b'=' {
            continue;
        }
        let v = char_to_val(b).expect("invalid base64");

        acc = (acc << 6) | v as u32;
        bits = bits + 6;

        if bits >= 8 {
            bits = bits - 8;
            out.push((acc >> bits) as u8);
        }
    }
    String::from_utf8(out).expect("Not valid utf-8")
}

fn char_to_val(b: u8) -> Option<u8> {
    match b {
        b'A'..=b'Z' => Some(b - b'A'),
        b'a'..=b'z' => Some(b - b'a' + 26),
        b'0'..=b'9' => Some(b - b'0' + 52),
        b'+' => Some(62),
        b'/' => Some(63),
        _ => None,
    }
}
