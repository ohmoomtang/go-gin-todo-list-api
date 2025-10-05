package config

import "os"

var MONGODB_URI = os.Getenv("MONGODB_URI")
var MONGODB_DB_NAME = "todo_list"