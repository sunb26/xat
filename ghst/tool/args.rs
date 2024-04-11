#[derive(Debug, clap::Parser)]
pub struct Args {
  /// Path to the blank GST/HST Return Working Copy form.
  #[arg(long, env = "GHST_FORM")]
  pub form: camino::Utf8PathBuf,
  /// Path to write filled GST/HST Return Working Copy form.
  #[arg(long, env = "GHST_OUT")]
  pub out: Option<camino::Utf8PathBuf>,
  /// Form fields to fill out in `<selector>=<value>` flags.
  #[arg(long, env = "GHST_FIELDS")]
  pub fields: Vec<ghst::Field>,
}
