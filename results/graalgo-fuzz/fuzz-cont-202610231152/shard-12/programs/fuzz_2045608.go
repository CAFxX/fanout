package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []int{32768}
    var v2 float64 = (1e10 / 1.5)
    v3 := make(map[uint32]int)
    var v4 uint16 = 130
    var v5 uint8 = 111
    fmt.Println(v2)
    if true {
        v4 = (^(v4))
    }
    v6 := GenS1{F0: ""}
    var v7 uint32 = 2147483648
    var v8 uint16 = v4
    v1 = append(v1, 839)
    var v9 uint8 = (v5 >> ((4294967295) & 63))
    var v10 bool = false
    v11 := GenS1{F0: "\n\t"}
    v10 = (!((v10 || ((v10 || v10) || v10)) || v10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
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
