package services_test

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/AdemarTellecher/pos-web-go/core/entities"
	"github.com/AdemarTellecher/pos-web-go/core/services"
)

func TestStore(t *testing.T) {
	b := &entities.Beer{
		ID:    1,
		Name:  "Heiniken",
		Type:  entities.TypeAle,
		Style: entities.StylePale,
	}
	// Teste de conexão com o banco de dados
	//dbPath := "../../database/"
	db, err := sql.Open("sqlite3", "../../database/beer_test.db")
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
	saved, err := service.Get(1)
	if err != nil {
		t.Fatalf("Erro ao fazer busca no banco de dados...%s", err.Error())
	}
	if saved.ID != 1 {
		t.Fatalf("Dados invalido. Esperando %d, e foi recebido %d", 1, saved.ID)
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
