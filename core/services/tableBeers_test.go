package services_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/AdemarTellecher/pos-web-go/core/entities"
	"github.com/AdemarTellecher/pos-web-go/core/services"
)

func TestStore(t *testing.T) {
	// ID utilizado para fazer teste de consulta por ID no banco de dados
	dbPath := "../../database/beer_test.db"
	var getID int64 = 0

	b := &entities.Beer{
		Name:  "Heiniken",
		Type:  entities.TypeAle,
		Style: entities.StylePale,
	}
	// Teste de conexão com o banco de dados
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatalf("Erro ao se conectar com o bando de dados %s", err.Error())
	}

	defer db.Close() // Fecha conexão com banco de dados, depois de executar as funções ou retornar algum erro
	err = clearDB(db)
	if err != nil {
		t.Fatalf("Erro ao limpar o banco de dados %s ", err.Error())

	}

	// Teste de gravação no banco de dados
	service := services.NewService(db)
	err = service.Store(b)
	if err != nil {
		t.Fatalf("Erro ao gravar os dados no banco: %s", err.Error())
	}
	//Teste de busca no banco de dados
	saved, err := service.Get(getID)
	if err != nil {
		t.Fatalf("Erro ao fazer busca no banco de dados...%s", err.Error())
	}
	if saved.ID != getID {
		t.Fatalf("ID invalido ou não existe. Esperando %d, e foi recebido %d", saved.ID, getID)
	}
}

func clearDB(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from beers")
	tx.Commit()
	return err
}
