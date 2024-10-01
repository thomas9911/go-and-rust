use ::safer_ffi::prelude::*;

pub fn count_numbers_inner(s: &str) -> u32 {
    s.chars().filter(|c| c.is_numeric()).count() as u32
}

#[ffi_export]
fn count_numbers(text_pointer: char_p::Ref<'_>) -> u32 {
    let text = text_pointer.to_str();
    count_numbers_inner(text)
}

#[cfg(feature = "headers")] // c.f. the `Cargo.toml` section
pub fn generate_headers() -> ::std::io::Result<()> {
    // we can do .ok() because the next call will error with a good message
    std::fs::create_dir("out/").ok();

    ::safer_ffi::headers::builder()
        .to_file("out/go_and_rust.h")?
        .generate()
}
