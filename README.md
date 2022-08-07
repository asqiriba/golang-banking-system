# Banking System

Create and manage bank accounts and money transactions.

## Installation

We launch a local image of Postgres

```bash
make postgres
```

Then, we run the migration script (Install from [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate))

```bash
make createdb
make migrateup
```

## DB Schema

![](db/provision/db-schema.png)

```sql
Table accounts as A {
  id bigserial [pk]
  owner varchar [not null]
  balance bigint [not null]
  currency varchar [not null]
  created_at timestamptz [not null, default: `now()`]

  Indexes{
    owner
  }
}

Table entries {
  id bigserial [pk]
  account_id bigint [ref: > A.id, not null]
  amount bigint [not null]
  created_at timestamptz [not null, default: `now()`]

  Indexes{
    account_id
  }
}

Table transfers {
  id bigserial [pk]
  from_account_id bigint [ref: > A.id, not null]
  to_account_id bigint [ref: > A.id, not null]
  amount bigint [not null, note: 'can only be positive']
  created_at timestamptz [not null, default: `now()`]

  Indexes{
    from_account_id
    to_account_id
    (from_account_id, to_account_id)
  }
}
```

We create the migration scripts for the database schema.

```bash
migrate create -ext sql -dir db/migration -seq init_schema
```

SSH into the Postgres image and run the following commands to create the database.

```bash
createdb --username=root --owner=root simple_bank
```
