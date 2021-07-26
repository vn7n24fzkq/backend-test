package config

type Database struct {
	Host     string
	Type     string
	User     string
	Password string
	Port     string
	Database string
}

// TODO use environmont variable
var DatabaseConfig Database = Database{
	Host:     "127.0.0.1",
	Type:     "mysql",
	User:     "admin",
	Password: "password",
	Port:     "3306",
	Database: "backend_test",
}

type Server struct {
	Host string
	Port string
}

var ServerConfig Server = Server{}

func InitConfig() {

}
