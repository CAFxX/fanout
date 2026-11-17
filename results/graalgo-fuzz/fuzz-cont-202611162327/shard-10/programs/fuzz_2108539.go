package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }

func main() {
    var v1 uint32 = 675
    v2 := make(map[int]bool)
    v3 := []string{"hello", "hello"}
    var v4 int = 24
    v5 := GenS1{F0: false, F1: 1, F2: (("a" + "") + "hello")}
    v4 = (v4 % (((v4 ^ (v4 | v4))) | 1))
    var v6 uint16 = 1
    var v7 int8 = -128
    var v8 uint = 548
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 float64 = (1e10 / 3.14159)
        v10 = v10
        _ = fi9
        _ = v10
    }
    v1 -= v1
    v1 = (^(v1))
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 bool = true
        delete(v2, v4)
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v4 -= v4
        v6 = (^(v6))
        v14 := GenS1{F0: true, F1: (v6 ^ (v6 >> (((v8 % ((v8) | 1))) & 63))), F2: ""}
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
