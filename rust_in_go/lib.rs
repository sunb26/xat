#[allow(unsafe_code, clippy::missing_panics_doc)]
#[no_mangle]
pub extern "C" fn hello_message() -> *const libc::c_char {
  use tap::Pipe;
  #[allow(clippy::unwrap_used)]
  "Hello, world!"
    .to_string()
    .pipe(std::ffi::CString::new)
    .unwrap()
    .into_raw()
}
