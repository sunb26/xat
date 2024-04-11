#[derive(Debug, clap::Parser)]
pub struct Args {
  /// Path to the blank GST/HST Return Working Copy form.
  #[arg(long, env = "GHST_FORM")]
  pub form: camino::Utf8PathBuf,
  /// Path to write filled GST/HST Return Working Copy form.
  #[arg(long, env = "GHST_OUT")]
  pub out: camino::Utf8PathBuf,
}
