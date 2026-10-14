package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 bool }
func genFn0(p0 bool, p1 uint, p2 int8) uint16 {
    var v4 uint8 = 128
    v4 = v4
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := []int{}
        p1 = (p1 << (((p1 ^ (p1 & 496))) & 63))
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}
func genFn1(p0 int) int16 {
    var v7 uint8 = 2
    var v8 int16 = 0
    v7 -= (^(v7))
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            v7 = (v7 + v7)
            _ = fi10
        }
        _ = fi9
    }
    _ = p0
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"\n\t"}
    var v3 float32 = 0.1
    v3 -= float32(0.1)
    delete(v1, 2147483647)
    if true {
        v11 := GenS1{F0: 0, F1: 1, F2: false}
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        if true {
            for fi13 := 0; fi13 < 5; fi13++ {
                fmt.Println(fi13, fi12, v3)
                var v14 bool = (!true)
                _ = fi13
                _ = v14
            }
            v15 := GenS1{F0: 2, F1: -128, F2: (347 == 42)}
            v3 = v3
            _ = v15
        } else {
            switch "a" {
            case "\n\t":
                v3 += (0.1 - (0.1 - -2.25))
            }
            _ = len(v1)
        }
        var v16 uint16 = 1
        _ = fi12
        _ = v16
    }
    v17 := GenS1{F0: 2, F1: -32, F2: (3.14159 != (0.0 - float64((float32(0.0) / v3))))}
    var v18 float64 = 0.0
    switch ("foo bar" + "hello") {
    case "\n\t":
        v18 += v18
    case "a":
        v18 -= v18
    default:
        var v19 int64 = -9223372036854775808
        _ = v19
    }
    var v20 int16 = -8
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
}
