package main

//go:generate go tool github.com/zyedidia/unionize -output=union.go Template template.go
type Template struct {
	i1 uint32
	i2 uint16
}
