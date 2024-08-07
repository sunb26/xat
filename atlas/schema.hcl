table "receipt_snapshot_v1" {
  schema = schema.public
  comment = "The cumulative totals associated with a receipt such as subtotal, gratuity, tax, etc. The breakdown of the subtotal is tracked as individual expenses."
  column "receipt_id" {
    null = false
    type = bigint
  }
  column "snapshot_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 1
      increment = 1
    }
  }
  column "scan_id" {
    null = true
    type = bigint
    comment = "If scan_id is null, this means the snapshot is a result of manual alterations to the entry. If it is not null, that means this snapshot is based on the inference result of an image scan."
  }
  column "receipt_date" {
    null = false
    type = date
    comment = "The date of purchase for the receipt."
  }
  column "create_time" {
    null = false
    type = timestamptz
  }
  primary_key {
    columns = [column.receipt_id, column.snapshot_id]
  }
  foreign_key "receipt_id" {
    columns     = [column.receipt_id]
    ref_columns = [table.receipt_v1.column.receipt_id]
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
table "receipt_v1" {
  schema = schema.public
  comment = "This table links receipts back to their project. No data should be stored here."
  column "project_id" {
    null = false
    type = bigint
  }
  column "receipt_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 1
      increment = 1
    }
  }
  index "receipt_id_idx" {
    unique  = true
    columns = [column.receipt_id]
  }
  primary_key {
    columns = [column.project_id, column.receipt_id]
  }
  foreign_key "project_id" {
    columns     = [column.project_id]
    ref_columns = [table.project_v1.column.project_id]
    on_update   = CASCADE
    on_delete   = CASCADE
  }
}
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
      start = 1
      increment = 1
    }
  }
  column "scan_id" {
    null = true
    type = bigint
    comment = "If scan_id is null, this means the snapshot is a result of manual alterations to the entry. If it is not null, that means this snapshot is based on the inference result of an image scan."
  }
  column "tag" {
    null = false
    type = enum.expense_tag
    comment = "The tag associated with an expense to differentiate between expense types."
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
    comment = "The percentage of the expense that is deductible. If an expense is deductible, then it can be subtracted from taxable income so that you get taxed less. See link: <https://www.investopedia.com/articles/tax/09/self-employed-tax-deductions.asp>"
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
  comment = "This table links expenses back to its receipt. No data should be stored here."
  column "receipt_id" {
    null = false
    type = bigint
  }
  column "expense_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 1
      increment = 1
    }
  }
  primary_key {
    columns = [column.receipt_id, column.expense_id]
  }
  foreign_key "receipt_id" {
    columns     = [column.receipt_id]
    ref_columns = [table.receipt_v1.column.receipt_id]
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
      start = 1
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
      start = 1
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
      start = 1
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
    comment = "This is either Commission or Business/Professional. See link <https://help.wealthsimple.com/hc/en-ca/articles/4408339655323-How-do-I-report-my-self-employment-income-on-a-T2125> for more information."
  }
  column "industry_code" {
    null = false
    type = text
    comment = "The code pertaining to the industry the business is related to. These are mostly used for statistical purposes. See link <https://www.canada.ca/en/revenue-agency/services/tax/businesses/topics/sole-proprietorships-partnerships/report-business-income-expenses/industry-codes.html>."
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
    comment = "An attribute of the T2125 form - additional info is required if the business is a partnership. See link <https://help.wealthsimple.com/hc/en-ca/articles/4408339655323-How-do-I-report-my-self-employment-income-on-a-T2125>."
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
      start = 1
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
    type = uuid
  }
  column "project_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 1
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
      start = 1
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
      start = 1
      increment = 1
    }
  }
  column "image_url" {
    null = false
    type = text
    comment = "The url of where the image is stored on Google Drive."
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
    type = uuid
  }
  column "snapshot_id" {
    null = false
    type = bigint
    identity {
      generated = ALWAYS
      start = 1
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
    type = uuid
    comment = "The UUID generated from Clerk <https://clerk.com/>."
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
view "receipt_data_v1" {
  schema = schema.public
  column "receipt_id" {
    type = bigint
  }
  column "subtotal" {
    type = money
    comment = "The subtotal of all expenses before tax on a receipt."
  }
  column "tax" {
    type = money
    comment = "The tax on the given receipt."
  }
  column "gratuity" {
    type = money
    comment = "The gratuity on the given receipt."
  }
  column "total" {
    type = money
    comment = "The total after tax on the given receipt."
  }
  column "receipt_date" {
    type = date
  }
  column "scan_id" {
    type = bigint
  }
  column "project_id" {
    type = bigint
  }
  as  = <<-SQL
    WITH LatestReceiptSnapshots AS (
        SELECT DISTINCT ON (rs.receipt_id)
          rs.receipt_id,
          rc.project_id,
          rs.receipt_date,
          rs.scan_id
        FROM receipt_snapshot_v1 rs
        JOIN receipt_v1 rc ON rs.receipt_id = rc.receipt_id
        ORDER BY rs.receipt_id, rs.create_time DESC
    ),
    LatestExpenseSnapshots AS (
        SELECT DISTINCT ON (es.expense_id)
            es.expense_id,
            es.tag,
            es.amount
        FROM expense_snapshot_v1 es
        ORDER BY es.expense_id, es.create_time DESC
    ),
    AggregatedExpenses AS (
        SELECT
            e.receipt_id,
            SUM(CASE WHEN les.tag = 'purchase' OR les.tag = 'gratuity' THEN les.amount ELSE 0::money END) AS subtotal,
            SUM(CASE WHEN les.tag = 'tax' THEN les.amount ELSE 0::money END) AS tax,
            SUM(CASE WHEN les.tag = 'gratuity' THEN les.amount ELSE 0::money END) AS gratuity
        FROM LatestExpenseSnapshots les
        JOIN expense_v1 e ON les.expense_id = e.expense_id
        GROUP BY e.receipt_id
    )
    SELECT
        lrs.receipt_id,
        ae.subtotal,
        ae.tax,
        ae.gratuity,
        (ae.subtotal + ae.tax) AS total,
        lrs.receipt_date,
        lrs.scan_id,
        lrs.project_id
    FROM LatestReceiptSnapshots lrs
    JOIN AggregatedExpenses ae ON lrs.receipt_id = ae.receipt_id;
  SQL
  depends_on = [table.expense_v1, table.expense_snapshot_v1, table.receipt_snapshot_v1]
  comment = "A view to store the information displayed on the receipt form of a given receipt id. Note that gratuity is included in the subtotal already but tax is not."
}

enum "expense_tag" {
  schema = schema.public
  values = ["purchase", "gratuity", "tax"]
  comment = "The tags associated with expenses for filtering during calculations."
}
schema "public" {
  comment = "standard public schema"
}
