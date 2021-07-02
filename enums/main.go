package main

import (
	"fmt"
)

type Domain string

const (
	DataCentreManagment  Domain = "DataCentreManagment"
	NetworTowerManagment        = "NetworTowerManagment"
)

func (d Domain) isValid() bool {
	switch d {
	case DataCentreManagment, NetworTowerManagment:
		return true
	}
	return false
}

func main() {
	a := "DataCentreManagment"
	x := Domain(a)
	isValid := x.isValid()
	if isValid {
		fmt.Println("Valid")
	} else {
		fmt.Println("Not valid")
	}
}
