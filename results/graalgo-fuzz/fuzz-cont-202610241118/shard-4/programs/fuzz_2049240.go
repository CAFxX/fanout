package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint64) uint {
    switch "foo bar" {
    case "foo bar":
        p0 = (^(p0))
    case "hello":
        v7 := GenS1{F0: 1, F1: 701}
        _ = v7
    }
    p0 = (p0 * 42)
    var v8 int8 = 1
    _ = p0
    _ = v8
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint64 = 1
    var v3 float64 = -2.25
    var v4 uint16 = 0
    var v5 uint8 = 2
    var v6 int8 = 77
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            v3 += v3
            _ = fi10
        }
        var v11 uint32 = 4294967295
        switch fi9 {
        case -1:
            var v12 uint16 = (v4 % ((370) | 1))
            _ = v12
        }
        _ = fi9
        _ = v11
    }
    if (v3 != 1.5) {
        v2 = v2
    }
    if (true || true) {
        v6 += v6
    }
    switch (((("a" + ("foo bar" + "a")) + ("a" + "\n\t")) + "") + "foo bar") {
    case "":
        v13 := GenS2{F0: 42}
        _ = v13
    default:
        for fi14 := 0; fi14 < 19; fi14++ {
            switch "a" {
            case "\n\t":
                v5 += 86
            case "a":
                v6 -= (v6 % ((-1) | 1))
            default:
                v1[0] = fi14
            }
            _ = fi14
        }
    }
    if (((v6 & v6) < 33) || false) {
        var v15 int32 = 0
        var v16 int32 = -2147483648
        _ = v15
        _ = v16
    }
    var v17 float32 = 0.0
    fmt.Println(v17, v6, v2)
    v1[uint32(v5)] = 1
    v18 := []int{0, 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
