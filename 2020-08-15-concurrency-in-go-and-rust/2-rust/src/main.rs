use std::env;
use std::fs;
use std::collections::{HashMap};

fn main() {
    let mut words = HashMap::<String, u32>::new();

    for arg in env::args().skip(1) {
        println!("{}", arg);
        tally_words(arg.to_string(), &mut words);
    }

    for (word, count) in words.iter() {
        if *count > 1 {
            println!("{} {}", count, word)
        }
    }
}


fn tally_words(filename: String, words: &mut HashMap<String, u32>) {
    let contents = fs::read_to_string(filename).expect("Cannot read file");

    for s in contents.split_whitespace() {
        let key = s.to_lowercase();
        *words.entry(key).or_insert(0) += 1;

    }
}
