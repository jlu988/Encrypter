resource "aws_dynamodb_table" "encryption" {
  name           = "Encryption"
  read_capacity  = 10
  write_capacity = 10
  hash_key       = "OriginalKey"

  attribute {
    name = "OriginalKey"
    type = "S"
  }
}
