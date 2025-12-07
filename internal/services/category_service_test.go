package services

import (
	"fmt"
	"log"
	"testing"

	"github.com/osamah22/open-mart/internal/models"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

func loadEnv() {
	viper.SetConfigName(".env") // name of file (no path)
	viper.SetConfigType("env")  // dotenv format

	viper.AddConfigPath(".")         // current folder
	viper.AddConfigPath("../../")    // go up two levels
	viper.AddConfigPath("../../../") // useful if running tests from deeper dirs

	viper.AutomaticEnv() // env vars override the file

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error loading config:", err)
	}
}

func TestGetCategory(t *testing.T) {
	loadEnv()
	fmt.Println("viper:", viper.GetString("DB_URL"))

	db, err := models.NewDatabase()
	require.NoError(t, err)

	queries := models.New(db)
	s := NewCategoryService(queries)

	categories, err := s.ListCategories(t.Context())
	require.NoError(t, err)
	require.NotEmpty(t, categories)
}
