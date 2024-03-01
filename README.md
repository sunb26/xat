# xat

- GST/HST calculator
- upload receipts
- sum up HST across all the receipts
- subtract the expense hst from the revenue hst

Tax Form: https://www.canada.ca/content/dam/cra-arc/migration/cra-arc/tx/bsnss/tpcs/gst-tps/bspsbch/rtrns/wrkngcp-eng.pdf

## run

1. install `aspect` build tool per https://github.com/aspect-build/aspect-cli?tab=readme-ov-file#installation
2. generate build files `aspect run //:gazelle`
3. inspect generated build files `git diff`
4. commit generated build files
5. select a target to run (e.g. in `//:BUILD.bazel` there is `:xat`)
6. run the target `aspect run //:xat`
>>>>>>>
