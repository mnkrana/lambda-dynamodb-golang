- Make sure to use primary key as "uuid" in dynamodb table
- Make sure to add env vars for table name and region in lambda env vars

```
KEY_REGION = "dynamodb_region"
KEY_TABLE  = "dynamodb_table"
```

- First function call should be Init() to get the evn vars and create a session

