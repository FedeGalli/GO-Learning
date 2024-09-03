package main

import "fmt"

type engine interface {
	getKmLeft() uint16
}

type gasEngine struct {
	kml    uint16
	liters uint16
	owner  owner
}

type elettricEngine struct {
	kmkwh uint16
	kwh   uint16
	owner owner
}

type owner struct {
	name    string
	surname string
}

func (e gasEngine) getKmLeft() uint16 {
	return e.kml * e.liters
}

func (e elettricEngine) getKmLeft() uint16 {
	return e.kmkwh * e.kwh
}

func canMekeIt(e engine, kms uint16) bool {
	if e.getKmLeft() < kms {
		return false
	}
	return true
}

func main() {
	var myGasEngine gasEngine = gasEngine{kml: 5, liters: 30, owner: owner{"Gino", "Pasticcino"}}
	fmt.Println(myGasEngine.kml, myGasEngine.liters, myGasEngine.owner)

	var myAbstarctGasEngine = struct {
		kml    uint8
		liters int8
	}{10, 100}

	fmt.Println(myAbstarctGasEngine.kml, myAbstarctGasEngine.liters)

	fmt.Println(myGasEngine.getKmLeft())

	myElettricEngine := elettricEngine{10, 10, owner{"Coglio", "Ne"}}

	fmt.Println(canMekeIt(myGasEngine, 30))
	fmt.Println(canMekeIt(myElettricEngine, 255))
}
