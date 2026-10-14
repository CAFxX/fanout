package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint8 }
func genFn0(p0 uint32) int8 {
    v6 := []int{541, 485, 599}
    p0 = p0
    if true {
        if true {
            if (32767 > 0) {
                fmt.Println(p0)
            } else {
                p0 = p0
                v7 := GenS1{F0: (p0 << ((255) & 63))}
                _ = v7
            }
            p0 = (p0 - 2147483648)
        }
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 float64 = 0.1
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            p0 = (p0 | (^(p0)))
            var v11 int8 = 100
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = v6
    return -128
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = float32((0.0 * 1e10))
    var v3 uint16 = 33
    v4 := make(map[int]int)
    var v5 int32 = 1
    var v12 uint64 = 18446744073709551615
    for fi13 := 0; fi13 < 9; fi13++ {
        _ = len(v1)
        var v14 float64 = 1e10
        switch (fi13 + 0) {
        case -618:
            var v15 uint32 = 42
            _ = v15
        case -351:
            var v16 int16 = -32768
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    var v17 uint32 = 1
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int64 = 1
        _ = fi18
        _ = v19
    }
    v2 = 0.0
    fmt.Println(v3, v2, v5)
    v20 := []int{-9223372036854775808, 32768}
    if (!true) {
        var v21 int32 = v5
        var v22 int = -541
        _ = v21
        _ = v22
    }
    var v23 uint8 = 2
    fmt.Println(v5, v5, v2)
    v24 := GenS2{F0: 438, F1: 65535, F2: 127}
    switch ("hello" + "foo bar") {
    case "a":
        fmt.Println(v2, v5)
    case "foo bar":
        var v25 int32 = (^(v5))
        _ = v25
    case "hello":
        switch "hello" {
        case "hello":
            var v26 uint8 = 127
            _ = v26
        case "\n\t":
            var v27 float32 = (1e10 * (((v2 - ((v2 / v2) - v2)) + (float32(-2.25) - 0.0)) - v2))
            _ = v27
        case "":
            v3 = v3
        }
    default:
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
