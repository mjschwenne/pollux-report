use std::fs::File;

use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
struct PersonalData {
    first_name: String,
    family_name: String,
    display_name: Option<String>,
    age: Option<u64>,
    // test: String,
    shoesize: u64,
}

#[derive(Serialize, Deserialize, Debug)]
struct MinimalPersonalData {
    first_name: String,
    family_name: String,
}

fn main() {
    let person = PersonalData {
        first_name: "Matthew".to_string(),
        family_name: "Schwennesen".to_string(),
        display_name: Some("Matt".to_string()),
        age: Some(24),
        // test: "test".to_string(),
        shoesize: 11,
    };
    let min_person = MinimalPersonalData {
        first_name: "Matthew".to_string(),
        family_name: "Schwennesen".to_string(),
    };

    let output_file = File::create("person.cbor").unwrap();
    let _ = ciborium::into_writer(&person, output_file);
    let serialized = serde_json::to_string(&person).unwrap();
    println!("JSON = {}", serialized);

    let m_output_file = File::create("mperson.cbor").unwrap();
    let _ = ciborium::into_writer(&min_person, m_output_file);
}
