package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int }
func genFn0(p0 uint16, p1 int64) uint16 {
    v5 := []int{4294967295}
    v6 := GenS1{F0: 128, F1: 810, F2: 277}
    p0 = p0
    var v7 uint16 = p0
    v8 := GenS1{F0: 255, F1: 32767, F2: 2147483648}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return (v7 / ((1) | 1))
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 125
    var v4 int32 = 2147483647
    v9 := []int{32767}
    fmt.Println(v3)
    var v10 uint8 = (v3 * 2)
    for fi11 := 0; fi11 < 1; fi11++ {
        if false {
            v9 = append(v9, fi11)
            if true {
                v4 -= (v4 + (v4 - (v4 | v4)))
            }
            switch ("hello" + ("a" + ("hello" + "\n\t"))) {
            case "foo bar":
                var v12 uint16 = 0
                _ = v12
            default:
                v3 = v3
            }
        }
        for fi13 := 0; fi13 < 20; fi13++ {
            v4 += v4
            if (3.14159 > -2.25) {
                v14 := v2[2]
                v15 := GenS1{F0: v10, F1: 950, F2: 2}
                _ = v14
                _ = v15
            }
            fmt.Println(fi11, v10, v10)
            _ = fi13
        }
        var v16 bool = false
        _ = fi11
        _ = v16
    }
    if len(v9) > 0 {
        v9[0] = 0
    }
    v10 += (^(v10))
    switch "\n\t" {
    case "":
        fmt.Println(v10, v10)
    case "\n\t":
        v4 = 700
    case "a":
        if true {
            switch "hello" {
            case "hello":
                var v17 uint16 = 0
                _ = v17
            case "":
                v18 := GenS1{F0: v3, F1: 42, F2: uint32(v3)}
                _ = v18
            case "a":
                if (uint(v10) > 276) {
                    var v19 int = -255
                    _ = v19
                } else {
                    var v20 uint = 617
                    _ = v20
                }
            }
            v21 := GenS2{F0: 2, F1: 1}
            var v22 float64 = -2.25
            _ = v21
            _ = v22
        }
    default:
        fmt.Println(v10, v4)
    }
    switch ("" + "a") {
    case "foo bar":
        var v23 int8 = -128
        _ = v23
    default:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    v10 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
