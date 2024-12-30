**Connecting to Databases and Writing Queries in Go**

Go (or Golang) is a powerful, statically typed language widely used for building scalable applications. In this tutorial, we'll cover how to connect to three popular databases—PostgreSQL, SQLite, and MySQL—using Go, along with examples of basic queries.

---

## **1. Setting Up Your Go Project**

Before working with databases, ensure you have Go installed. You can create a new project with the following commands:

```bash
mkdir go-database-tutorial
cd go-database-tutorial
go mod init go-database-tutorial
```

To interact with databases, we'll use the `database/sql` package along with specific drivers for each database:

- **PostgreSQL**: `github.com/lib/pq`
- **SQLite**: `modernc.org/sqlite`
- **MySQL**: `github.com/go-sql-driver/mysql`

Install the relevant drivers for your database:

```bash
go get github.com/lib/pq
# For SQLite
go get modernc.org/sqlite
# For MySQL
go get github.com/go-sql-driver/mysql
```

---

## **2. Connecting to Databases**

### **PostgreSQL Connection**

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=username password=yourpassword dbname=yourdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to PostgreSQL successfully!")
}
```

### **SQLite Connection**

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "file:test.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to SQLite successfully!")
}
```

### **MySQL Connection**

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connStr := "username:password@tcp(127.0.0.1:3306)/yourdb"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Connected to MySQL successfully!")
}
```

---

## **3. Writing Basic Queries**

### **Creating a Table**

For all databases, we can use the `Exec` method to execute SQL statements such as creating tables.

#### Example (for PostgreSQL, SQLite, or MySQL):

```go
query := `CREATE TABLE IF NOT EXISTS posts (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255),
	content TEXT,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)`

_, err := db.Exec(query)
if err != nil {
	log.Fatalf("Error creating table: %v", err)
}
fmt.Println("Table created successfully!")
```

### **Inserting Data**

```go
query := `INSERT INTO posts (title, content) VALUES ($1, $2)`
// Use ? instead of $1, $2 for MySQL and SQLite
_, err := db.Exec(query, "My First Post", "This is the content of my first post.")
if err != nil {
	log.Fatalf("Error inserting data: %v", err)
}
fmt.Println("Data inserted successfully!")
```

### **Fetching Data**

```go
rows, err := db.Query("SELECT id, title, content FROM posts")
if err != nil {
	log.Fatalf("Error fetching data: %v", err)
}
defer rows.Close()

for rows.Next() {
	var id int
	var title, content string
	if err := rows.Scan(&id, &title, &content); err != nil {
		log.Fatalf("Error scanning row: %v", err)
	}
	fmt.Printf("ID: %d, Title: %s, Content: %s\n", id, title, content)
}
```

### **Updating Data**

```go
query := `UPDATE posts SET content = $1 WHERE id = $2`
// Use ? for MySQL and SQLite
_, err := db.Exec(query, "Updated content", 1)
if err != nil {
	log.Fatalf("Error updating data: %v", err)
}
fmt.Println("Data updated successfully!")
```

### **Deleting Data**

```go
query := `DELETE FROM posts WHERE id = $1`
// Use ? for MySQL and SQLite
_, err := db.Exec(query, 1)
if err != nil {
	log.Fatalf("Error deleting data: %v", err)
}
fmt.Println("Data deleted successfully!")
```

---



### converting data to slice 

you can fetch data and convert it to slice and send it or work with the data


```go
type Note struct{
	Id string
	Title string
	Text string
}
getNotes:=func (w http.ResponseWriter, r *http.Request)  {
row,_ := db.Query("SELECT id, title, text FROM notes")
defer row.Close()

var notes []Note // we make a slice of Notes and each note is a struct 
for row.Next() {  // we can use the  next function to itrate over all the rows inside the qurey 
	var note Note
	row.Scan(&note.Id,&note.Title, &note.Text) // with the scan function we add all the data to each note inside the slice 
	notes = append(notes, note) // we append each note to the note slice 
}
html, _ := template.ParseFiles("index.html")
html.Execute(w,notes)

}
```



---

## **4. Wrapping Up**

You’ve now learned how to connect to PostgreSQL, SQLite, and MySQL databases and execute basic SQL queries using Go. To make your applications production-ready, consider the following best practices:

- Use connection pooling libraries like `github.com/jmoiron/sqlx` or `github.com/golang-migrate/migrate` for advanced database management.
- Sanitize inputs to prevent SQL injection.
- Use environment variables or a configuration file to manage database credentials.

With these fundamentals, you can start building Go applications that interact with databases efficiently!

