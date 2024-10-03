use ::safer_ffi::prelude::*;
use safer_ffi::option::TaggedOption;

#[derive_ReprC]
#[repr(u8)] 
pub enum SongError {
    InvalidTitle = 0,
    InvalidArtist = 1,
    InvalidReleaseYear = 2,
}

#[derive_ReprC]
#[repr(C)]
pub struct GoResult<T, E> {
    pub val: TaggedOption<T>,
    pub err: TaggedOption<E>,
}

impl<T, E> GoResult<T, E> {
    pub fn error(e: E) -> Self {
        GoResult {
            val: TaggedOption::None,
            err: TaggedOption::Some(e),
        }
    }

    pub fn ok(t: T) -> Self {
        GoResult {
            val: TaggedOption::Some(t),
            err: TaggedOption::None,
        }
    }
}

#[derive_ReprC]
#[repr(C)]      // <- defined C layout is mandatory!
pub struct Song {
    title: char_p::Box,
    artist: char_p::Box,
    release_year: u32,
}

#[ffi_export]
fn new_song(title: char_p::Ref<'_>, artist: char_p::Ref<'_>, release_year: u32) -> Song {
    Song{
        title: title.to_owned(),
        artist: artist.to_owned(),
        release_year
    }
}

#[ffi_export]
fn try_new_song(title: char_p::Ref<'_>, artist: char_p::Ref<'_>, release_year: u32) -> GoResult<Song, SongError> {
    if release_year < 1400 || release_year > 2100 {
        return GoResult::error(SongError::InvalidReleaseYear);
    }
    
    GoResult::ok(Song{
        title: title.to_owned(),
        artist: artist.to_owned(),
        release_year
    })
}

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
