package main

type vehicle interface {
	GetAutonomy() float32
}
type car struct {
	tyres uint8
	seats uint8
	fuel  float32
	kml   float32
}
type truck struct {
	tyres  uint8
	weight float32
	fuel   float32
	kml    float32
}

func (c car) GetAutonomy() float32 {
	return c.fuel * c.kml
}

func (t truck) GetAutonomy() float32 {
	return t.fuel * t.kml
}

func MaxAvailableDistance(v1 vehicle, v2 vehicle) vehicle {
	if v1.GetAutonomy() >= v2.GetAutonomy() {
		return v1
	}
	return v2
}

func Add(n1, n2 int) int {
	return n1 + n2
}

func Mul(n1, n2 int) int {
	return n1 * n2
}

func Run(n1, n2, n3 int, Function func(int, int) int) int {
	return Function(Function(n1, n2), n3)
}

func main() {

}
