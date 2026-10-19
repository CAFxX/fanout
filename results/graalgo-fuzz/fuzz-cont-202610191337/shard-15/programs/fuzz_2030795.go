package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }

func main() {
    v1 := []uint32{2147483648, 2147483647}
    var v2 uint = 2
    v3 := make(map[int]string)
    v4 := make(map[uint32]int)
    for fi5 := 0; fi5 < 11; fi5++ {
        fmt.Println(v2)
        _ = fi5
    }
    v4[244] = 749
    var v6 int = -572
    var v7 int8 = 127
    v7 += 100
    var v8 uint16 = 32767
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(v6, v6, v8)
        fmt.Println(fi9)
        _ = fi9
    }
    v2 = v2
    var v10 float32 = 3.14159
    if (!false) {
        v10 += ((v10 + (v10 / v10)) * v10)
    }
    fmt.Println(v6, v6, v2)
    v10 -= (v10 - (v10 / -2.25))
    var v11 bool = true
    var v12 float64 = (1.5 + 3.14159)
    var v13 float32 = v10
    v10 = (3.14159 * v10)
    switch ("foo bar" + ("a" + ("a" + "\n\t"))) {
    case "foo bar":
        v14 := GenS1{F0: 42, F1: 706}
        _ = v14
    case "":
        v6 -= v6
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
