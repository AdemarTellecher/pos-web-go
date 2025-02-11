package services

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/AdemarTellecher/pos-web-go/core/entities"
)

// Define a interface com as funções que serão usadas pelo restante do projeto
type UseCase interface {
	GetAll() ([]*entities.Beer, error)
	Get(ID int64) (*entities.Beer, error)
	Store(b *entities.Beer) error
	Update(b *entities.Beer) error
	Remove(ID int64) error
}

// em Go qualquer coisa que implemente as funções de uma interface
// passa a ser uma implementação válida. Não existe a palavra "implements" como em Java ou PHP
// desta forma, uma struct, uma string, um inteiro, etc, qualquer coisa pode ser válido, desde
// que implemente todas as funções
type Service struct {
	DB *sql.DB
}

// Esta função retorna um ponteiro em memoria para uma estrutura pre-definida
func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

// Declarando as (funções/metodos) que implementam a interface UseCase{}, que serão implentadas posteriomente
func (s *Service) GetAll() ([]*entities.Beer, error) {
	// Result é um slice de ponteiros do tipo Beer
	// Se existir um erro, a função vai retorna-lo, e o erro deve ser tratado
	// pelo chamador de função.
	//Vamos sempre usar a conexão que esta dentro do Service

	var result []*entities.Beer

	rows, err := s.DB.Query("select id, name, type, style from beers")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var b entities.Beer

		err = rows.Scan(&b.ID, &b.Name, &b.Type, &b.Style)
		if err != nil {
			return nil, err
		}
		result = append(result, &b)
	}
	return result, nil
}

func (s *Service) Get(ID int64) (*entities.Beer, error) {
	// A variavel 'b' é do tipo Beer e recebe valores somente deste tipo
	var b entities.Beer

	// Verificando se a consulta esta valida
	stmt, err := s.DB.Prepare("select id, name, type, style from berrs where id=?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&b.ID, &b.Name, &b.Type, &b.Style)
	if err != nil {
		return nil, err
	}
	// O retorno final da função, deve retornar a posição da memoria de da variavel'b'
	return &b, nil
}

func (s *Service) Store(b *entities.Beer) error {
	// Iniciando uma transação com o banco de dados
	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("insert into beers(id, name, type, style) values(?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	//o comando Exec, retorna um Result, mas não temos interesse nele, por isso podemos ignorá-lo com o '_' underline
	_, err = stmt.Exec(b.ID, b.Name, b.Type, b.Style)
	if err != nil {
		tx.Rollback() // Se deu erro, desfaz todas as alterações ao estado inicial
		return err
	}
	tx.Commit() // Commita no banco de dados, ou seja, grava no banco de dados
	return nil
}

func (s *Service) Update(b *entities.Beer) error {
	if b.ID <= 0 {
		return fmt.Errorf("ID: %q é invalido ", b.ID)
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare("update bbers set name=?, type=?, style=? where id=?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	//o comando Exec retorna um Result, mas não temos interesse nele, por isso podemos ignorá-lo com o _
	_, err = stmt.Exec(b.ID, b.Name, b.Type, b.Style)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (s *Service) Remove(ID int64) error {
	if ID <= 0 {
		return fmt.Errorf("ID: %q é invalido", ID)
	}

	tx, err := s.DB.Begin()
	if err != nil {
		return err
	}

	//o comando Exec retorna um Result, mas não temos interesse nele, por isso podemos ignorá-lo com o _
	_, err = tx.Exec("delete from beers where id=?", ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
