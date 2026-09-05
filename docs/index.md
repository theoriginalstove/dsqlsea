# dsqlsea Documentation

> And lo, the Great One looked down upon the people and proclaimed:
> "SQL is actually pretty great"

dsqlsea generates **fully type-safe idiomatic Go code** from SQL. Here's how it
works:

1. You write your SQL queries.
2. You run `dsql` to generate Go code that presents type-safe interfaces to those
   queries
3. You write application code that calls the methods `dsql` generated

Seriously, it's that easy. You don't have to write any boilerplate SQL querying
code ever again.
