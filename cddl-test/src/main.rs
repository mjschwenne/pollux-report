use cddl::{self, validate_cbor_from_slice};

fn main() {
    let test_f64: f64 = 32.0;

    let mut encoded: Vec<u8> = Vec::new();
    let _ = ciborium::into_writer(&test_f64, &mut encoded);

    let cddl_f16 = r#"testrule = float16"#;
    let cddl_f32 = r#"testrule = float32"#;
    let cddl_f64 = r#"testrule = float64"#;
    assert!(validate_cbor_from_slice(cddl_f16, &encoded, Some(&["cbor"])).is_ok());
    assert!(validate_cbor_from_slice(cddl_f32, &encoded, Some(&["cbor"])).is_ok());
    assert!(validate_cbor_from_slice(cddl_f64, &encoded, Some(&["cbor"])).is_ok());
}
