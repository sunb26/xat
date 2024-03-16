# xat

- GST/HST calculator
- upload receipts
- sum up HST across all the receipts
- subtract the expense hst from the revenue hst

Tax Form: https://www.canada.ca/content/dam/cra-arc/migration/cra-arc/tx/bsnss/tpcs/gst-tps/bspsbch/rtrns/wrkngcp-eng.pdf

## run
### MacOS
1. install `aspect` build tool [here](https://github.com/aspect-build/aspect-cli?tab=readme-ov-file#installation)
2. generate build files `aspect run //:gazelle`
3. inspect generated build files `git diff`
4. commit generated build files
5. select a target to run (e.g. in `//:BUILD.bazel` there is `:xat_go`)
6. run the target `aspect run //:xat_go`

### Linux
1. Install `bazelisk` as a wrapper for `aspect cli` [here](https://github.com/bazelbuild/bazelisk?tab=readme-ov-file#installation).
2. The rest follows from MacOS setup:
    - generate build files `bazel run //:gazelle`
    - inspect generated build files `git diff`
    - commit generated build files
    - select a target to run (e.g. in `//:BUILD.bazel` there is `:xat_go`)
    - run the target `bazel run //:xat_go`

## deploy

```bash
fly auth login
podman init
podman start
podman login registry.fly.io -u x --password $(fly auth token) \
  --authfile ~/.docker/config.json
aspect run //cmd/serve:push --config=deploy
fly deploy --config cmd/serve/fly.toml
```

## go

- avoid using your local go toolchain as it may lead to incompatibility
- use the builtin toolchain instead `aspect run @rules_go//go` as a replacement for `go`
- see [guide](https://github.com/bazelbuild/rules_go/blob/master/docs/go/core/bzlmod.md) for adding dependencies and managing `go.mod`

## rust

- enable rust-analyzer LSP support by generating `rust-project.json`
  ```bash
  aspect run @rules_rust//tools/rust_analyzer:gen_rust_project
  ```
- run hello world binary `aspect run //:xat_rust`
