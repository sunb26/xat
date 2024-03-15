table "test" {
  schema = schema.public
  column "id" {
    null = false
    type = integer
  }
  primary_key {
    columns = [column.id]
  }
}
schema "public" {
  comment = "standard public schema"
}
