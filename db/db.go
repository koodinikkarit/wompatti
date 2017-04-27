package db

import "fmt"

type context struct{}

func (c *context) Add() {

}

func Moi() {
	fmt.Println("Moikkamoi")
}

func New() context {
	return context{}
}
