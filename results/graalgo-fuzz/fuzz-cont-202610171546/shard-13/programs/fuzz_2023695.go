package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint32) int {
    v7 := []int{-9223372036854775808, 127, -1, 127}
    var v8 uint16 = 914
    if (0 > 1000) {
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS1{F0: -1}
            p0 = (p0 + (^(p0)))
            v8 = (v8 ^ v8)
            _ = fi9
            _ = v10
        }
        _ = len(v7)
    }
    _ = p0
    _ = v7
    _ = v8
    return 545
}

func main() {
    v1 := []uint32{2, 1, 65535}
    var v2 uint16 = 2
    var v3 bool = (!false)
    var v4 int = 42
    v5 := []int{}
    v6 := make(map[int]string)
    fmt.Println(v4, v3, v2)
    var v11 int16 = 0
    v12 := GenS1{F0: -42}
    var v13 uint8 = 128
    v6[2147483647] = "hello"
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 int8 = -92
        fmt.Println(v3, v2, v11)
        _ = fi14
        _ = v15
    }
    switch "\n\t" {
    case "hello":
        var v16 uint16 = (^(v2))
        _ = v16
    case "a":
        fmt.Println(v2)
    }
    var v17 float64 = (1.5 + -2.25)
    var v18 uint64 = 1
    v17 = float64((-2.25 / 1.0))
    switch "foo bar" {
    case "hello":
        v18 = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
