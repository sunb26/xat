#[tracing::instrument(err, skip(form), level = tracing::Level::INFO)]
pub fn fill_form(
  form: &mut pdf::file::File<
    Vec<u8>,
    pdf::file::ObjectCache,
    pdf::file::StreamCache,
    pdf::file::NoLog,
  >,
) -> eyre::Result<()> {
  tracing::info!("Hello, world!");
  let Some(ref forms) = form.get_root().forms else {
    eyre::bail!("some forms");
  };
  for field in &forms.fields {
    tracing::info!(?field);
    // TODO: accept config and update the form field.
    // https://github.com/pdf-rs/pdf/blob/c5481318fa4405567d914bb42009f33693a85127/examples/src/bin/form.rs#L96
  }
  Ok(())
}
