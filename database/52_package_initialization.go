package database

var connection string

// automatically run when package is imported
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
