use std::fs;

fn main() {
    let mut file = get_file_contents("../../test.pr");
    remove_comments(&mut file);
    println!("{}", &file);


}

fn remove_comments(file: &mut String) {
    let mut new_file = String::new();
    for line in file.lines() {
        let new_line;
        if let Some(i) = line.find("//") {
            new_line = line[0..i].to_string() + "\n";
            new_file.push_str(&new_line)
        }
        else {
            new_line = line.to_string() + "\n";
            new_file.push_str(&new_line);
        }
    }
    *file = new_file;
}

fn get_file_contents(path: &str) -> String {
    let contents = fs::read_to_string(path);
    match contents {
        Ok(file) => file,
        Err(_)   => String::new(),
    }
}

#[allow(dead_code)]
fn greet() {
    println!(
        "Welcome to the main compilation environment for the `Prog` Programming Language!"
    );
    println!("Rust will eventually be the primary language that Prog is implemented in,
so this will serve as the central hub for the project.");
}
