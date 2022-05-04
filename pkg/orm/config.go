package orm

type Config struct {
	Type        string `default:"sqlite3"`
	Host        string
	Port        int
	Dbname      string
	Username    string
	Password    string
	DBFile      string `yaml:"db_file" env:"DB_FILE"`
	ShowLog     bool
	Migration   bool   `default:"false"`
}
