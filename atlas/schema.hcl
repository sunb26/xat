table "expense_snapshot_v1" {
  schema = schema.public
  column "expense_id" {
    null = false
    type = bigint
  }
  column "snapshot_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  column "scan_id" {
    null = true
    type = bigint
  }
  column "title" {
    null = false
    type = text
    comment = "name of item/service expense"
  }
  column "amount" {
    null = false
    type = money
  }
  column "deductible" {
    null = false
    type = double_precision
    comment = "percentage of expense that is deductible"
  }
  column "create_time" {
    null = false
    type = timestamptz
  }
  primary_key {
    columns = [column.expense_id, column.snapshot_id]
  }
  foreign_key "expense_id" {
    columns     = [column.expense_id]
    ref_columns = [table.expense_v1.column.expense_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  foreign_key "scan_id" {
    columns     = [column.scan_id]
    ref_columns = [table.scan_v1.column.scan_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
table "expense_v1" {
  schema = schema.public
  column "project_id" {
    null = false
    type = bigint
  }
  column "expense_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  primary_key {
    columns = [column.project_id, column.expense_id]
  }
  foreign_key "project_id" {
    columns     = [column.project_id]
    ref_columns = [table.project_v1.column.project_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  index "expense_id_idx" {
    unique  = true
    columns = [column.expense_id]
  }
}
table "income_snapshot_v1" {
  schema = schema.public
  column "income_id" {
    null = false
    type = bigint
  }
  column "snapshot_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  column "scan_id" {
    null = true
    type = bigint
  }
  column "title" {
    null = false
    type = text
    comment = "name of product/service source of income"
  }
  column "amount" {
    null = false
    type = money
  }
  column "create_time" {
    null = false
    type = time
  }
  primary_key {
    columns = [column.income_id, column.snapshot_id]
  }
  foreign_key "income_id" {
    columns     = [column.income_id]
    ref_columns = [table.income_v1.column.income_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  foreign_key "scan_id" {
    columns     = [column.scan_id]
    ref_columns = [table.scan_v1.column.scan_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
table "income_v1" {
  schema = schema.public
  column "project_id" {
    null = false
    type = bigint
  }
  column "income_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  primary_key {
    columns = [column.project_id, column.income_id]
  }
  foreign_key "project_id" {
    columns     = [column.project_id]
    ref_columns = [table.project_v1.column.project_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  index "income_id_idx" {
    unique  = true
    columns = [column.income_id]
  }
}

# Refer to the link for explanations of the organization table fields
# https://help.wealthsimple.com/hc/en-ca/articles/4408339655323-How-do-I-report-my-self-employment-income-on-a-T2125

table "organization_snapshot_v1" {
  schema = schema.public
  column "organization_id" {
    null = false
    type = bigint
  }
  column "snapshot_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  column "title" {
    null = false
    type = text
    comment = "name of business"
  }
  column "income_type" {
    null = false
    type = text
  }
  column "industry_code" {
    null = false
    type = text
  }
  column "main_product" {
    null = true
    type = text
    comment = "main product/service sold by the business"
  }
  column "partnership" {
    null = false
    type = bool
    default = false
  }
  column "address" {
    null = true
    type = text
    comment = "business physical address if applicable"
  }
  column "create_time" {
    null = false
    type = time
  }
  primary_key {
    columns = [column.organization_id, column.snapshot_id]
  }
  foreign_key "organization_id" {
    columns     = [column.organization_id]
    ref_columns = [table.organization_v1.column.organization_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
table "organization_v1" {
  schema = schema.public
  column "organization_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
 column "title" {
    null = false
    type = text
    comment = "name of business"
  }
  column "income_type" {
    null = false
    type = text
  }
  column "industry_code" {
    null = false
    type = text
  }
  column "main_product" {
    null = true
    type = text
    comment = "main product/service sold by the business"
  }
  column "partnership" {
    null = false
    type = bool
    default = false
  }
  column "address" {
    null = true
    type = text
    comment = "business physical address if applicable"
  }
  column "create_time" {
    null = false
    type = time
  }
  primary_key {
    columns = [column.organization_id]
  }
}
table "project_v1" {
  schema = schema.public
  column "user_id" {
    null = false
    type = bigint
  }
  column "project_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  primary_key {
    columns = [column.user_id, column.project_id]
  }
  foreign_key "user_id" {
    columns     = [column.user_id]
    ref_columns = [table.user_v1.column.user_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  index "project_id_idx" {
    unique  = true
    columns = [column.project_id]
  }
}
table "scan_inference_v1" {
  schema = schema.public
  column "scan_id" {
    null = false
    type = bigint
  }
  column "scan_inference_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  column "inference_result" {
    null = false
    type = bytea
    comment = "JSON of inference output"
  }
  primary_key {
    columns = [column.scan_id, column.scan_inference_id]
  }
  foreign_key "scan_id" {
    columns     = [column.scan_id]
    ref_columns = [table.scan_v1.column.scan_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
table "scan_v1" {
  schema = schema.public
  column "project_id" {
    null = false
    type = bigint
  }
  column "scan_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  column "image" {
    null = false
    type = bytea
  }
  column "create_time" {
    null = false
    type = time
  }
  primary_key {
    columns = [column.project_id, column.scan_id]
  }
  foreign_key "project_id" {
    columns     = [column.project_id]
    ref_columns = [table.project_v1.column.project_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
  index "scan_id_idx" {
    unique  = true
    columns = [column.scan_id]
  }
}
table "user_snapshot_v1" {
  schema = schema.public
  column "user_id" {
    null = false
    type = bigint
  }
  column "snapshot_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  column "given_name" {
    null = false
    type = text
  }
  column "middle_name" {
    null = false
    type = text
  }
  column "family_name" {
    null = false
    type = text
  }
  column "date_of_birth" {
    null = false
    type = date
  }
  column "marital_status" {
    null = false
    type = text
  }
  column "citizenship" {
    null = false
    type = text
  }
  column "email" {
    null = false
    type = text
  }
  column "create_time" {
    null = false
    type = time
  }
  primary_key {
    columns = [column.user_id, column.snapshot_id]
  }
  foreign_key "user_id" {
    columns     = [column.user_id]
    ref_columns = [table.user_v1.column.user_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
table "user_v1" {
  schema = schema.public
  column "organization_id" {
    null = false
    type = bigint
  }
  column "user_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  column "given_name" {
    null = false
    type = text
  }
  column "middle_name" {
    null = false
    type = text
  }
  column "family_name" {
    null = false
    type = text
  }
  column "date_of_birth" {
    null = false
    type = date
  }
  column "marital_status" {
    null = false
    type = text
  }
  column "citizenship" {
    null = false
    type = text
  }
  column "email" {
    null = false
    type = text
  }
  column "create_time" {
    null = false
    type = time
  }
  primary_key {
    columns = [column.user_id]
  }
  foreign_key "organization_id" {
    columns     = [column.organization_id]
    ref_columns = [table.organization_v1.column.organization_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
schema "public" {
  comment = "standard public schema"
}
