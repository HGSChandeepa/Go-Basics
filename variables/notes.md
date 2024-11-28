# Go Variable Naming Rules

A variable can have a short name (like `x` and `y`) or a more descriptive name (e.g., `age`, `price`, `carname`).

## Go Variable Naming Rules

- A variable name must start with a letter or an underscore character (`_`).
- A variable name cannot start with a digit.
- A variable name can only contain alpha-numeric characters and underscores (`a-z`, `A-Z`, `0-9`, and `_`).
- Variable names are case-sensitive (`age`, `Age`, and `AGE` are three different variables).
- There is no limit on the length of the variable name.
- A variable name cannot contain spaces.
- The variable name cannot be any Go keywords.

## Multi-Word Variable Names

Variable names with more than one word can be difficult to read. There are several techniques you can use to make them more readable:

### Camel Case

Each word, except the first, starts with a capital letter:
```go
myVariableName = "John"
```

### Pascal Case

Each word starts with a capital letter:
```go
MyVariableName = "John"
```

### Snake Case

Each word is separated by an underscore character:
```go
my_variable_name = "John"
```
