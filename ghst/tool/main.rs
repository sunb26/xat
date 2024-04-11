fn main() -> eyre::Result<()> {
  use clap::Parser;
  tracing_subscriber::fmt()
    .with_env_filter(
      tracing_subscriber::EnvFilter::builder()
        .with_default_directive(tracing::Level::DEBUG.into())
        .from_env_lossy(),
    )
    .try_init()
    .map_err(|err| eyre::eyre!("{err:?}"))?;
  color_eyre::install()?;
  let args::Args { form, out } = args::Args::parse();
  tracing::info!(?form, ?out);
  ghst::fill_form(&mut pdf::file::FileOptions::cached().open(form)?)?;
  Ok(())
}
