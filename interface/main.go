package  main

import (
	"fmt"
)

type svc interface {
	Start()
}

type ServiceDef struct {
	Name string
	Def string
}

type ServiceProps struct {
	Desc string
}


func (sd ServiceDef) Start(){
	fmt.Println("starting the service")
}

func (sp ServiceProps) Start(){
	fmt.Println("starting the service prop")
}


func NewService(s svc){
	s.Start()
}

func main(){
	sd := ServiceProps{Desc: "marlon"}
	NewService(sd)

}