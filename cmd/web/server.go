package main

import (
	"github.com/gin-gonic/gin"
	"github.com/osamah22/open-mart/internal/auth"
	"github.com/osamah22/open-mart/internal/models"

	"github.com/osamah22/open-mart/internal/services"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"go.uber.org/zap"
)

type Server struct {
	Router          *gin.Engine
	Logger          *zap.SugaredLogger
	Config          Config
	OAuthConfig     *oauth2.Config
	CategoryService services.CategoryService
	authService     auth.AuthService
}

func NewServer(cfg Config) (*Server, error) {
	logger := newLogger()

	// db
	db, err := models.NewDatabase(models.DBConfig{URL: cfg.DB_URL})
	if err != nil {
		return nil, err
	}
	// services
	q := models.New(db)

	authSvc := auth.NewService(q)
	categorySvc := services.NewCategoryService(q)

	oauthCfg := &oauth2.Config{
		ClientID:     cfg.GOOGLE_CLIENT_ID,
		ClientSecret: cfg.GOOGLE_CLIENT_SECRET,
		RedirectURL:  cfg.BASE_URL + "/auth/google/callback",
		Scopes: []string{
			"openid",
			"email",
			"profile",
		},
		Endpoint: google.Endpoint,
	}

	r := gin.Default()
	srv := &Server{
		Router:          r,
		Logger:          logger,
		Config:          cfg,
		OAuthConfig:     oauthCfg,
		authService:     authSvc,
		CategoryService: categorySvc,
	}
	// 1) sessions FIRST
	store := cookie.NewStore([]byte(cfg.SESSION_KEY))
	r.Use(sessions.Sessions("openmart", store))
	r.Use(
		// gin.Recovery(), // or RecoverPanic()
		recovery(logger),
		secureHeaders(),
		AuthContext(srv.authService),
		// RequestLogger(logger),
	)

	// r.LoadHTMLGlob("ui/html/*.html")
	r.LoadHTMLGlob("ui/html/**/*.html")
	// r.LoadHTMLGlob("ui/html/pages/*")
	r.Static("/static", "./ui/static")

	registerRoutes(r, srv)

	return srv, nil
}
