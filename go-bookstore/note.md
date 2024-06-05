In Go, the `init` functions are executed before the `main` function runs. This is a fundamental part of the Go initialization process. When you run your Go program with `go run`, the Go runtime performs the following steps:

1. **Imports**: It loads and initializes all packages that are imported in your code.
2. **Initialization**: Within each package, any variables with package-level scope are initialized.
3. **`init` Functions**: It then executes the `init` functions for each package, if they exist.
4. **Main Package**: Finally, it executes the `main` function from the `main` package.

### Sequence of Execution

1. **Load Imports**: When `main.go` is run, Go starts by importing the necessary packages:

    ```go
    import (
        "log"
        "net/http"
        "fmt"
        "github.com/etharrra/go-bookstore/pkg/routes"
        "github.com/gorilla/mux"
        _ "github.com/jinzhu/gorm/dialects/mysql"
    )
    ```

2. **Initialize Packages**: Each imported package is initialized. For example, when `routes` is imported, it might import other packages, including your `models` package.

3. **Run `init` Functions**: Any `init` function within the imported packages is executed. In your case, the `init` function in `models/book.go` is executed:

    ```go
    func init() {
        config.Connect()
        db = config.GetDB()
        db.AutoMigrate(&Book{})
    }
    ```

4. **Run `main` Function**: After all `init` functions have been executed, the `main` function is finally called:
    ```go
    func main() {
        r := mux.NewRouter()
        routes.RegisterBookStoreRoutes(r)
        http.Handle("/", r)
        fmt.Println("Starting server at localhost:1010")
        log.Fatal(http.ListenAndServe("localhost:1010", r))
    }
    ```

### How It Works in Your Application

When you run `main.go`, here is the sequence of events:

1. **Import Packages**:

    - `main.go` imports `routes` which in turn might import `models`.

2. **Initialize and Run `init` Functions**:

    - The `models` package is imported, triggering its `init` function.
    - The `init` function in `models/book.go` calls `config.Connect()`, which establishes the database connection, and `db.AutoMigrate(&Book{})`, which migrates the `Book` model schema.

3. **Run `main`**:
    - The `main` function sets up the router, registers routes, and starts the HTTP server.

### In Summary

The `init` function in the `models` package is executed before the `main` function due to Go's initialization process. This ensures that the database connection is established and the schema is migrated before the main logic of your application starts running. This is why when you run `main.go`, the database connection is set up as expected.
