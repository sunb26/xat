load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
  name = "rules_rust",
  integrity = "sha256-ww398ehv1QZQp26mRbOkXy8AZnsGGHpoXpVU4WfKl+4=",
  urls = ["https://github.com/bazelbuild/rules_rust/releases/download/0.40.0/rules_rust-v0.40.0.tar.gz"],
)
load("@rules_rust//rust:repositories.bzl", "rules_rust_dependencies", "rust_register_toolchains", "rust_repository_set")
rules_rust_dependencies()
RUST_VERSION = "1.76.0"
rust_register_toolchains(edition = "2021", versions = [RUST_VERSION])
rust_repository_set(
  name = "rust_x86_64-unknown-linux-musl",
  edition = "2021",
  exec_triple = "aarch64-apple-darwin",
  extra_target_triples = ["x86_64-unknown-linux-musl"],
  versions = [RUST_VERSION],
)
load("@rules_rust//crate_universe:defs.bzl", "crate", "crates_repository", "render_config")
crates_repository(
  name = "crates",
  cargo_lockfile = "//:Cargo.lock",
  lockfile = "//:Cargo.bazel.lock",
  packages = {
    "color-eyre": crate.spec(version = "0.6.2"),
    "eyre": crate.spec(version = "0.6.12"),
    "libc": crate.spec(version = "0.2.153"),
    "tap": crate.spec(version = "1.0.1"),
  },
)
load("@crates//:defs.bzl", "crate_repositories")
crate_repositories()
load("@rules_rust//tools/rust_analyzer:deps.bzl", "rust_analyzer_dependencies")
rust_analyzer_dependencies()
