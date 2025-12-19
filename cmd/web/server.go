package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	CategoryService services.CategoryService
	OAuthConfig     *oauth2.Config
	authService     services.AuthService
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

	authSvc := services.NewService(q)
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
	store := cookie.NewStore([]byte(cfg.SESSION_KEY), []byte(cfg.SESSION_KEY))
	store.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400 * 1, // 24 hours
	})
	r.Use(sessions.Sessions("openmart", store))
	r.Use(
		gin.Recovery(), // or RecoverPanic()
		flashMiddleware(),
		secureHeaders(),
		AuthContext(srv.authService),
	)

	r.LoadHTMLGlob("ui/html/**/*.tmpl")
	// r.LoadHTMLGlob("ui/html/pages/*")
	r.Static("/static", "./ui/static")

	registerRoutes(r, srv)

	return srv, nil
}
