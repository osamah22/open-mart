package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"

	"github.com/go-playground/form/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Server struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	sessionManager *scs.SessionManager
	formDecoder    *form.Decoder
	templateCache  map[string]*template.Template
}

// loading the enviroment variables for testing
func loadEnv() {
	viper.SetConfigFile(".env") // the path to the parent project (where the env file exists)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error while loading the config files:", err)
	}
}

func main() {
	loadEnv()

	addr := flag.String("address", ":7070", "Network Address\tDefault: :7070")
	flag.Parse()

	loadEnv()

	dbUrl := viper.GetString("DB_URL")
	dbDriver := viper.GetString("DB_DRIVER")
	if dbUrl == "" || dbDriver == "" {
		log.Fatalf("dbUrl='%s' or dbDriver='%s' is empty", dbUrl, dbDriver)
	}

	conn, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		log.Fatal("Could not connect to DB:", err)
	}
	defer conn.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	sessionManager := scs.New()
	sessionManager.Store = postgresstore.New(conn)

	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.IdleTimeout = 30 * time.Minute
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.Secure = true // only if using HTTPS
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	app := &Server{
		infoLog:        infoLog,
		errorLog:       errorLog,
		templateCache:  templateCache,
		formDecoder:    form.NewDecoder(),
		sessionManager: sessionManager,
	}
	server := app.routes()

	log.Printf("Starting server on address %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, server))
}
