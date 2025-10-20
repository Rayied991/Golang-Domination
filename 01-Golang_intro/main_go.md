# Explanation of Go Program: Hello, World!

## 🧩 Program Overview

This Go program prints "Hello, World!" to the console. It demonstrates the basic structure of a Go program, including packages, imports, and the main function.

## 📝 Detailed Explanation

### 1. `package main`

Every Go file must belong to a package.  
The `main` package is special — it indicates that this file will produce an **executable program**, not a library.

### 2. `import "fmt"`

`fmt` stands for **format**.  
It is part of Go's standard library and provides functions for formatted input and output.  
For example, `fmt.Println()` prints output followed by a newline.

### 3. `func main()`

This defines the **entry point** of your Go program.  
When you execute the Go program, the code inside `main()` runs first.

### 4. `fmt.Println("Hello, World!")`

This line prints the text `Hello, World!` to the console.  
`Println()` adds a newline automatically after printing the text.

---

## 🧠 Summary

✅ This is the simplest Go program and a perfect starting point to understand the Go syntax and program structure.

---

## ▶️ How to Run the Program

1. Save your code in a file named `main.go`.
2. Open your terminal in the same directory as `main.go`.
3. Run the following command:

   ```bash
   go run main.go
   ```
