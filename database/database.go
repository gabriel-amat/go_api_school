package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gabriel-amat/go_api_school/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataBase() {
	log.Println("--Init DataBase--")
	dbName := "schooldb"
	defaultConnStr := "host=localhost port=5432 user=amat sslmode=disable"

	err := createDatabase(defaultConnStr)

	if err != nil {
		log.Fatalf("createDatabaseIfNotExists error: %v", err)
	}

	// Atualiza a string de conexão com o banco de dados criado
	conn := fmt.Sprintf("host=localhost port=5432 user=admin dbname=%s password=admin sslmode=disable", dbName)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		panic("Falha ao conectar ao banco de dados")
	}

	log.Println("--DataBase OK--")

	log.Println("--Running migrations--")
	db.AutoMigrate(
		&models.AlunoModel{},
		&models.CursoAluno{},
		&models.CursoModel{})

	log.Println("--Migrations ok--")
	DB = db
}

func createDatabase(connStr string) error {
	log.Println("--Create Database If Not Exists--")

	connStr = fmt.Sprintf("%s dbname=postgres", connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening data base: %v", err)
		return err
	}
	defer db.Close()

	// SQL script to create user admin
	err = executeSQLScript(db, "sql/create_user.sql")
	if err != nil {
		return fmt.Errorf("error executing create_user.sql: %v", err)
	}

	//Check if DB already exists
	var exists bool
	query := "SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'schooldb');"
	result := db.QueryRow(query).Scan(&exists)
	fmt.Printf("A tabela existe?:%t ", exists)
	if result != nil {
		log.Fatalf("Erro ao validar se tabela existe: %v", result)
		return result
	}
	if !exists {
		// SQL script to create database
		err = executeSQLScript(db, "sql/create_db.sql")
		if err != nil {
			return fmt.Errorf("error executing create_db.sql: %v", err)
		}
	}

	// Conect to created data base to grant privileges
	connStr = fmt.Sprintf("%s dbname=schooldb", connStr)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}
	defer db.Close()

	// SQL script to grant privileges
	err = executeSQLScript(db, "sql/grant_privileges.sql")
	if err != nil {
		return fmt.Errorf("error executing grant_privileges.sql: %v", err)
	}
	return nil
}

func executeSQLScript(db *sql.DB, scriptPath string) error {
	script, err := os.ReadFile(scriptPath)
	if err != nil {
		return fmt.Errorf("could not read SQL script file: %v", err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		return fmt.Errorf("could not execute SQL script: %v", err)
	}

	return nil
}
