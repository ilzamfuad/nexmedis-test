# Question 4

Given a database table orders with the following schema:
CREATE TABLE orders (
    id INT PRIMARY KEY,
    customer_id INT,
    product_id INT,
    order_date TIMESTAMP,
    amount DECIMAL(10, 2)
);

Assume that customer_id is indexed, but amount and order_date are not indexed.

- Write an optimized SQL query to find the top 5 customers who spent the most money in the past month.
- How would you improve the performance of this query in a production environment?

# Answer

1. Write an optimized SQL query to find the top 5 customers who spent the most money in the past month.

SELECT customer_id, SUM(amount) as total_spent
FROM orders
WHERE order_date >= date_trunc('month', current_date - interval '1' month)
  AND order_date < date_trunc('month', current_date)
GROUP BY customer_id
ORDER BY total_spent DESC
LIMIT 5;

2. How would you improve the performance of this query in a production environment?

- Add composite index on `(order_data, customer_id)` to optimize the filtering, depend on the total data, if the data is so big index on column `customer_id` is needed, if the total data is small, index on column `customer_id` is not needed, but either way its not reducing its performance.
- Also, we can consider partitioning/sharding the table by date to reduce the amount of data scanned for query.
