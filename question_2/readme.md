# Question 2

Consider a database table Users with the following columns:
- id (Primary Key)
- username
- email
- created_at

Your task is to design an indexing strategy to optimize the following queries:
1. Fetch a user by username.
2. Fetch users who signed up after a certain date (created_at > "2023-01-01").
3. Fetch a user by email.

Explain which columns you would index, and whether composite indexes or individual indexes would be appropriate for each query. Discuss trade-offs in terms of read and write performance.

# Answer
To optimize the given queries, the following indexing strategy can be used:

1. Fetch a user by username:
  - Index: Individual index on the `username` column.
  - Reason: This will allow the database to quickly locate the user by their username.

2. Fetch users who signed up after a certain date (created_at > "2023-01-01"):
  - Index: Individual index on the `created_at` column.
  - Reason: This index will speed up range queries on the `created_at` column, since it only use for search range, it only need individual index. If there is other query like `username` and `created_at` as order, we dont need index for `created_at`,  only index on column `username` is enough.

3. Fetch a user by email:
  - Index: Individual index on the `email` column.
  - Reason: Similar to the `username` index.

## Trade off 
- Index can significanly increase speed of query, but it also decrease the performance of write query because it need to reindex after create. Also index takes additional disk space, so that the more index it have, the more storage it used.
- For the trade-off, we can consider split the process between read and write. For example we can use read replicas for read process, and main database for write only process. So that write process will not decrease the performance, and read process can perform and adjusted as well
- For other trade-off, we can sharding the database, we can split the data based on created_at, we can split based on year and month, so that the read query will scan small part of database, also write query performace is also optimize as well.
- In this case composite index is not necessary, because the required query is targeting only single target column.
