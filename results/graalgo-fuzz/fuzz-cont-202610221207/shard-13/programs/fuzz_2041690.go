package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8, p1 int8, p2 int) uint {
    v5 := GenS1{F0: 2147483648, F1: 0}
    p0 = -60
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 375
}
func genFn1(p0 int16) uint32 {
    p0 = p0
    v6 := GenS2{F0: 127}
    _ = p0
    _ = v6
    return 693
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    var v3 uint32 = 826
    var v4 float32 = -2.25
    v7 := GenS1{F0: v3, F1: 255}
    v8 := GenS1{F0: v3, F1: 245}
    var v9 float64 = (1.5 * 1e10)
    for fi10 := 0; fi10 < 7; fi10++ {
        switch "foo bar" {
        case "hello":
            v4 += v4
        }
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    fmt.Println(v4, v9)
    switch ((("\n\t" + "foo bar") + ("a" + "\n\t")) + ((("\n\t" + "\n\t") + ((("a" + ("" + ("\n\t" + "foo bar"))) + "") + "\n\t")) + "a")) {
    case "hello":
        var v12 uint8 = 255
        _ = v12
    default:
        v4 = (float32(v9) - float32(0.0))
    }
    v13 := GenS1{F0: genFn1(-677), F1: 2}
    for fi14 := 0; fi14 < 1; fi14++ {
        var v15 int8 = 0
        switch "\n\t" {
        case "\n\t":
            fmt.Println(fi14)
        case "hello":
            var v16 int8 = v15
            _ = v16
        case "":
            if (-1 < 1000) {
                v17 := GenS1{F0: (v3 * v3), F1: 125}
                _ = v17
            }
        default:
            fmt.Println(v4, v9)
        }
        _ = fi14
        _ = v15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
