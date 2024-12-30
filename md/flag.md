# The `flag` package in Go 
The `flag` package in Go is used to parse command-line flags. Here are some of the most commonly used functions from the `flag` package:

1. **`flag.String()`**  
   Defines a string flag with a default value and a description.
   ```go
   var name = flag.String("name", "default", "your name")
   ```

2. **`flag.Int()`**  
   Defines an integer flag with a default value and a description.
   ```go
   var age = flag.Int("age", 30, "your age")
   ```

3. **`flag.Bool()`**  
   Defines a boolean flag with a default value and a description.
   ```go
   var verbose = flag.Bool("verbose", false, "enable verbose output")
   ```

4. **`flag.Parse()`**  
   Parses the command-line flags. This must be called after all flags are defined.
   ```go
   flag.Parse()
   ```

5. **`flag.Args()`**  
   Returns the non-flag arguments after parsing.
   ```go
   args := flag.Args()
   ```

6. **`flag.NArg()`**  
   Returns the number of non-flag arguments.
   ```go
   numArgs := flag.NArg()
   ```

7. **`flag.NFlag()`**  
   Returns the number of flags that have been set.
   ```go
   numFlags := flag.NFlag()
   ```

8. **`flag.Float64()`**  
   Defines a float64 flag with a default value and description.
   ```go
   var weight = flag.Float64("weight", 75.5, "your weight in kg")
   ```

9. **`flag.Duration()`**  
   Defines a duration flag (time.Duration) with a default value and description.
   ```go
   var timeout = flag.Duration("timeout", 10*time.Second, "timeout duration")
   ```

10. **`flag.Set()`**  
    Allows setting a flag programmatically.
    ```go
    flag.Set("name", "John")
    ```

11. **`flag.Lookup()`**  
    Returns the flag with the given name (if it exists).
    ```go
    f := flag.Lookup("name")
    ```

These functions allow you to define and manage flags in your Go programs, making it easy to handle user input from the command line.