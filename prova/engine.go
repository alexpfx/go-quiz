package prova

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func FromJson(file string) *Prova {
	bs, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	var prova *Prova

	err = json.Unmarshal(bs, &prova)
	if err != nil {
		log.Println(string(bs))
		log.Fatal(err.Error())
	}
	return prova
}

func ToJson(prova *Prova) string {
	bs, err := json.MarshalIndent(prova, "", "    ")
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(bs)
}

type Alternativa struct {
	Label string
	Texto string
}

type Questao struct {
	Enunciado    string
	Alternativas []Alternativa
	Correta      Alternativa
}

type Prova struct {
	Questoes []Questao
	atual    int
}

func (p *Prova) Go(index int) error {
	if index < 0 {
		return fmt.Errorf("índice passado é negativo. index: %d", index)
	}
	if index > len(p.Questoes)-1 {
		return fmt.Errorf("índice passado é maior que a quantidade de Questoes. index: %d", index)
	}
	p.atual = index
	return nil
}
func (p *Prova) Next() error {
	index := p.atual + 1
	if index > len(p.Questoes)-1 {
		return fmt.Errorf("índice passado é maior que a quantidade de Questoes. index: %d", index)
	}
	p.atual = index

	return nil
}

func (p *Prova) Prev() error {
	index := p.atual - 1

	if index < 0 {
		return fmt.Errorf("índice passado é negativo. index: %d", index)
	}
	p.atual = index
	return nil
}

func (p *Prova) Get() *Questao {
	return &p.Questoes[p.atual]
}

func (p *Prova) Check(a Alternativa) bool {
	q := p.Get()
	return q.Correta == a
}

type Engine interface {
	Go(index int) error
	Get() *Questao
	Next() error
	Prev() error
	Check(alt Alternativa) bool
}
