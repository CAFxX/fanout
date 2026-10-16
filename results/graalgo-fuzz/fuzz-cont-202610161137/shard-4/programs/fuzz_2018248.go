package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]string)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[uint32]int)
    switch ("\n\t" + "") {
    case "a":
        fmt.Println(v3)
    }
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 uint16 = 587
        var v7 float64 = -2.25
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v3, v1)
    v8 := GenS1{F0: 42, F1: 42, F2: 0}
    var v9 bool = (!true)
    v1 += 2147483647
    v1 = (^(v1))
    v10 := GenS1{F0: 65535, F1: 2, F2: 65535}
    v11 := GenS2{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
