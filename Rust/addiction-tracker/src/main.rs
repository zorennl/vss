use std::env::args;
use std::fs::{File, create_dir, read, write};
fn push(file: &str, value: u16) {
    write(file, value.to_le_bytes()).expect("Failed to write to file");
}
fn pull(file: &str) -> std::io::Result<u16> {
    let bytes = read(file).expect("Failed to read file please run 'goon init' first");
    Ok(u16::from_le_bytes([bytes[0], bytes[1]]))
}

fn main() {
    let input = args().nth(1);
    const STREAK_PATH: &str = "/home/zmynz/.local/share/addiction/streak.bin";
    const RESETS_PATH: &str = "/home/zmynz/.local/share/addiction/resets.bin";

    let mut streak: u16 = 0;
    let mut resets: u16 = 0;

    if input == Some("init".to_string()) {
        println!("Instantiting required files");
        create_dir("/home/zmynz/.local/addiction/goon/").expect("Failed to create folder");
        File::create_new(STREAK_PATH).expect("Failed to create file");
        push(STREAK_PATH, 0);
        File::create_new(RESETS_PATH).expect("Failed to create file");
        push(RESETS_PATH, 0);
    } else {
        streak = pull(STREAK_PATH).unwrap();
        resets = pull(RESETS_PATH).unwrap();
    }

    if input == Some("reset".to_string()) {
        println!("Fucking loser");
        push(STREAK_PATH, 0);
        push(RESETS_PATH, resets + 1)
    }
    if input == Some("add".to_string()) {
        println!("HELL YEAH!");
        println!("{}", streak + 1);
        push(STREAK_PATH, streak + 1);
    }
    if input == Some("info".to_string()) {
        println!("streak: {streak}, resets: {resets}");

        if streak >= 100 {
            println!("Why are you even using this lol");
        } else if streak >= 75 {
            println!("Almost there");
        } else if streak >= 50 {
            println!("Halfway there keep going");
        } else if streak >= 25 {
            println!("Hell Yeah!");
        }
    }
}
