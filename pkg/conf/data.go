package conf

var SystemConfig = &system{
	Port:          ":8080",
	Debug:         false,
	SessionSecret: "",
	HashIDSalt:    "",
}

var DatabaseConfig = &database{
	Type:        "sqlite3",
	Host:        "",
	Port:        3306,
	User:        "root",
	Password:    "root",
	Database:    "random_donate",
	Charset:     "utf8",
	File:        "random_donate.db",
	TablePrefix: "rd_",
}

var CORSConfig = &cors{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"PUT", "POST", "GET", "OPTIONS"},
	AllowHeaders:     []string{"Cookie", "Authorization", "Content-Length", "Content-Type"},
	AllowCredentials: false,
	ExposeHeaders:    nil,
}

var OptionOverwrite = map[string]interface{}{}
