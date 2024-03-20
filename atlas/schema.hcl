table "expense_snapshot_v1" {
  schema = schema.public
  comment = "A snapshot in time of a given expense entry. A new snapshot is created every time the expense entry is modified."
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
    comment = "If scan_id is null, this means the snapshot is a result of manual alterations to the entry. If it is not null, that means this snapshot is based on the inference result of an image scan."
  }
  column "title" {
    null = false
    type = text
    comment = "The name of item/service expense."
  }
  column "amount" {
    null = false
    type = money
    comment = "The amount listed on the receipt for this entry."
  }
  column "deductible" {
    null = false
    type = double_precision
    comment = "The percentage of the expense that is deductible. If an expense is deductible, then it can be subtracted from taxable income so that you get taxed less. <https://www.investopedia.com/articles/tax/09/self-employed-tax-deductions.asp>"
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
  comment = "This table links expenses back to its project. No data should be stored here."
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
  comment = "A snapshot in time of a given income entry. A new snapshot is created every time the income entry is modified."
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
    comment = "If scan_id is null, this means the snapshot is a result of manual alterations to the entry. If it is not null, that means this snapshot is based on the inference result of an image scan."
  }
  column "title" {
    null = false
    type = text
    comment = "The name of the source of income."
  }
  column "amount" {
    null = false
    type = money
    comment = "The amount listed on the invoice for this source of income."
  }
  column "create_time" {
    null = false
    type = timestamptz
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
  comment = "This table links income back to its project. No data should be stored here. Note that income does not mean revenue."
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

table "organization_snapshot_v1" {
  schema = schema.public
  comment = "A snapshot in time of a given organization entry. A new snapshot is created every time the organization entry is modified. Refer to the link for explanations of the organization table fields: <https://help.wealthsimple.com/hc/en-ca/articles/4408339655323-How-do-I-report-my-self-employment-income-on-a-T2125>."
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
    comment = "The name of organization."
  }
  column "income_type" {
    null = false
    type = text
    comment = "This is either Commission or Business/Professional. See <https://help.wealthsimple.com/hc/en-ca/articles/4408339655323-How-do-I-report-my-self-employment-income-on-a-T2125> for more information."
  }
  column "industry_code" {
    null = false
    type = text
    comment = "The code pertaining to the industry the business is related to. These are mostly used for statistical purposes. <https://www.canada.ca/en/revenue-agency/services/tax/businesses/topics/sole-proprietorships-partnerships/report-business-income-expenses/industry-codes.html>"
  }
  column "main_product" {
    null = true
    type = text
    comment = "The main product/service sold by the organization. This field is optional on the T2125 form."
  }
  column "partnership" {
    null = false
    type = bool
    default = false
    comment = "An attribute of the T2125 form - additional info is required if the business is a partnership. <https://help.wealthsimple.com/hc/en-ca/articles/4408339655323-How-do-I-report-my-self-employment-income-on-a-T2125>"
  }
  column "address" {
    null = true
    type = text
    comment = "The organization's physical address if applicable."
  }
  column "create_time" {
    null = false
    type = timestamptz
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
  comment = "An organization refers to the business that is being filed under. All organization data lives in the organization_snapshot table."
  column "organization_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 0
      increment = 1
    }
  }
  primary_key {
    columns = [column.organization_id]
  }
}
table "project_v1" {
  schema = schema.public
  comment = "A project can be any time a user starts a new form. A user can have many projects."
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
  comment = "This table stores the results of inferencing on the image scans."
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
    type = jsonb
    comment = "JSON of image scan inference output."
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
  comment = "Stores the compressed images to be scanned and inferenced."
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
  column "image_content" {
    null = false
    type = bytea
    comment = "The content of the compressed image file."
  }
  column "image_compression_algorithm" {
    null = false
    type = text
  }
  column "create_time" {
    null = false
    type = timestamptz
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
  comment = "A snapshot in time of a user profile. A new snapshot is generated everytime a user modifies information related to their profile."
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
    type = timestamptz
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
  comment = "A user can open many projects and belongs to an organization. All user data lives in the user_snapshot table."
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
