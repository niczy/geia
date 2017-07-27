package tool
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"


func Connect() {
	_, err := sql.Open("mysql", "user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true")

	fmt.Printf("%s", err)
}
