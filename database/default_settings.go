package database

import "github.com/AH-dark/random-donate/pkg/conf"

var defaultSettings = []Setting{
	{Name: "app_version", Type: "system", Value: conf.AppVersion},
	{Name: "db_version", Type: "system", Value: conf.DbVersion},
	{Name: "site_url", Type: "basic", Value: "http://localhost:8080/"},
	{Name: "site_name", Type: "basic", Value: "v我50"},
	{Name: "site_description", Type: "basic", Value: "v我50，为你一生做牛做马"},
}
