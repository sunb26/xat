#[tracing::instrument(err, skip(form), level = tracing::Level::INFO)]
pub fn fill_form(
  mut form: pdf::file::File<
    Vec<u8>,
    pdf::file::ObjectCache,
    pdf::file::StreamCache,
    pdf::file::NoLog,
  >,
  fields: &[Field],
) -> eyre::Result<
  pdf::file::File<Vec<u8>, pdf::file::ObjectCache, pdf::file::StreamCache, pdf::file::NoLog>,
> {
  use pdf::object::Updater;
  let Some(ref forms) = form.get_root().forms else {
    eyre::bail!("some forms");
  };
  for field in &forms.fields {
    tracing::info!(?field);
  }
  for replace in fields {
    let Field { selector, value } = replace;
    let mut to_replace = None;
    let Some(ref forms) = form.get_root().forms else {
      eyre::bail!("some forms");
    };
    for field in &forms.fields {
      if !(selector.is_match(
        &field
          .name
          .as_ref()
          .map_or_else(|| Ok(String::new()), pdf::primitive::PdfString::to_string)?,
      ) || selector.is_match(
        &field
          .alt_name
          .as_ref()
          .map_or_else(|| Ok(String::new()), pdf::primitive::PdfString::to_string)?,
      )) {
        continue;
      }
      to_replace.replace((*field).clone());
    }
    let Some(to_replace) = to_replace.inspect(|found| tracing::info!(?found)) else {
      break;
    };
    let mut replaced = (*to_replace).clone();
    replaced.value = value.clone();
    form.update(to_replace.get_ref().get_inner(), replaced)?;
  }
  Ok(form)
}

#[derive(Debug, Clone)]
pub struct Field {
  selector: lazy_regex::Regex,
  value: pdf::primitive::Primitive,
}

impl std::str::FromStr for Field {
  type Err = eyre::Error;

  #[tracing::instrument(err, level = tracing::Level::DEBUG)]
  fn from_str(s: &str) -> Result<Self, Self::Err> {
    use itertools::Itertools;
    let Some((sel, val)) = s.split('=').collect_tuple() else {
      eyre::bail!("exactly one `=` separating selector and value");
    };
    Ok(Self {
      selector: lazy_regex::Regex::new(sel)?,
      value: pdf::primitive::Primitive::String(pdf::primitive::PdfString::new(val.into())),
    })
  }
}
