# Atlas

## Inspect

Reads database schema from database at `<DSN>` and writes to `./schema.hcl`.

```bash
bazel run //atlas:inspect --action_env="DSN=<DSN>"
```

## Apply

Applies `schema.hcl` to database at `<DSN>`.

```bash
DSN="<DSN>" bazel run //atlas:apply
```
