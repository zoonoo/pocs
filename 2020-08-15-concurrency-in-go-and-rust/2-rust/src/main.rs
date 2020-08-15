// extern crate rayon;

use std::env;
use std::fs;
use std::collections::HashMap;
use rayon::prelude::*;
use std::sync::Mutex;

fn main() {
    let mut words = Mutex::new(HashMap::<String, u32>::new());

    env::args().skip(1).collect::<Vec<String>>().par_iter().for_each(
        |arg| tally_words(arg.to_string(), &words));

    for (word, count) in words.lock().unwrap().iter() {
        if *count > 1 {
            println!("{} {}", count, word)
        }
    }
}


fn tally_words(filename: String, words: &Mutex<HashMap<String, u32>>) {
    let contents = fs::read_to_string(filename).expect("Cannot read file");

    for s in contents.split_whitespace() {
        let key = s.to_lowercase();
        *words.lock().unwrap().entry(key).or_insert(0) += 1;

    }
}
