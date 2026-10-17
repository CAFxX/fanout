package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    v3 := make(map[uint32]bool)
    var v4 int8 = 113
    var v5 bool = false
    var v6 string = "hello"
    v7 := GenS1{F0: 143}
    v6 = v6
    fmt.Println(v6)
    var v8 uint = 18446744073709551615
    var v9 uint16 = 292
    v10 := GenS2{F0: (v5 || v5), F1: 42}
    for fi11 := 0; fi11 < 14; fi11++ {
        if len(v2) > 0 {
            v2[0] = v6
        }
        _ = fi11
    }
    var v12 uint8 = 0
    v13 := GenS2{F0: v5, F1: 255}
    v5 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
