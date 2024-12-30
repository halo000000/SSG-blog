# `fmt` Package

The `fmt` package provides functions for formatted I/O, string formatting, and error handling. It is part of the standard library in Go.

---

### **Output Functions**

#### **1. `fmt.Println`**
Prints values with a newline at the end.
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Go is awesome!")
}
```

**Output:**
```
Hello, World!
Go is awesome!
```

#### **2. `fmt.Print`**
Prints values without adding a newline.
```go
package main

import "fmt"

func main() {
    fmt.Print("Hello, ")
    fmt.Print("World!")
}
```

**Output:**
```
Hello, World!
```

#### **3. `fmt.Printf`**
Prints formatted values using format specifiers.
```go
package main

import "fmt"

func main() {
    name := "Areen"
    age := 16
    fmt.Printf("My name is %s and I am %d years old.\n", name, age)
}
```

**Output:**
```
My name is Areen and I am 16 years old.
```

---

### **String Formatting Functions**

#### **4. `fmt.Sprintln`**
Formats values and returns them as a string with a newline.
```go
package main

import "fmt"

func main() {
    str := fmt.Sprintln("Hello,", "World!")
    fmt.Print(str)
}
```

**Output:**
```
Hello, World!
```

#### **5. `fmt.Sprint`**
Formats values and returns them as a string without a newline.
```go
package main

import "fmt"

func main() {
    str := fmt.Sprint("Hello, ", "World!")
    fmt.Print(str)
}
```

**Output:**
```
Hello, World!
```

#### **6. `fmt.Sprintf`**
Formats values and returns them as a formatted string.
```go
package main

import "fmt"

func main() {
    name := "Areen"
    age := 16
    str := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
    fmt.Println(str)
}
```

**Output:**
```
My name is Areen and I am 16 years old.
```

---

### **Error Formatting**

#### **7. `fmt.Errorf`**
Creates a formatted error message of type `error`.
```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    err := fmt.Errorf("this is an error: %w", errors.New("underlying issue"))
    fmt.Println(err)
}
```

**Output:**
```
this is an error: underlying issue
```

---

### **Input Functions**

#### **8. `fmt.Scan`**
Reads space-separated values from standard input.
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scan(&name)
    fmt.Println("Hello,", name)
}
```

**Example Input:**
```
Enter your name: Areen
```

**Output:**
```
Hello, Areen
```

#### **9. `fmt.Scanln`**
Reads values from standard input until a newline is encountered.
```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello,", name)
}
```

**Example Input:**
```
Enter your name: Areen
```

**Output:**
```
Hello, Areen
```

#### **10. `fmt.Scanf`**
Reads formatted input from standard input.
```go
package main

import "fmt"

func main() {
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("Hello %s, you are %d years old.\n", name, age)
}
```

**Example Input:**
```
Enter your name and age: Areen 16
```

**Output:**
```
Hello Areen, you are 16 years old.
```

---

These examples showcase how to use each function in the `fmt` package. Let me know if you want further clarification!