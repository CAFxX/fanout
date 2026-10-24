package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []string{"foo bar", "\n\t"}
    v2 := make(map[string]string)
    var v3 string = "hello"
    var v4 float32 = 3.14159
    v5 := []uint32{}
    var v6 uint16 = 979
    var v7 uint8 = 1
    var v8 uint8 = (v7 | v7)
    v9 := GenS1{F0: true}
    v8 = v7
    v4 = float32((-2.25 - ((float64(float32(1.5)) / 0.1) * 3.14159)))
    var v10 uint64 = 648
    var v11 int = 127
    v8 = v7
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
