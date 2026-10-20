package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint }
func genFn0(p0 uint, p1 int8) uint64 {
    p0 += (p0 ^ (p0 - (^(p0))))
    p1 = p1
    var v5 uint8 = 1
    p0 = p0
    if true {
        v6 := []int{299, 42, 1, 9223372036854775807}
        var v7 uint = (p0 | p0)
        v8 := GenS1{F0: (true || ((32767 >= 0) || false)), F1: (p0 / ((4294967295) | 1))}
        _ = v6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return 18446744073709551615
}
func genFn1(p0 uint) uint64 {
    fmt.Println(p0)
    v9 := []int{-549, 128, -1, 1}
    var v10 int64 = -9223372036854775808
    v10 = v10
    v11 := GenS2{F0: p0}
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 2
}

func main() {
    var v1 int64 = 1
    var v2 float64 = 3.14159
    var v3 uint16 = 42
    v4 := []uint32{}
    v1 += v1
    if false {
        switch 9223372036854775807 {
        case -1:
            v1 = -9223372036854775808
        case 1:
            if false {
                _ = len(v4)
                if len(v4) > 0 {
                    v4[0] = 626
                }
                v2 -= 1.5
            } else {
                var v12 float32 = 1.5
                _ = v12
            }
        case 128:
            v2 = (v2 * v2)
        }
        if false {
            var v13 float32 = -2.25
            var v14 uint64 = 0
            _ = v13
            _ = v14
        }
        v3 = 255
    }
    fmt.Println(v2)
    var v15 float64 = 0.0
    var v16 int32 = 0
    _ = len(v4)
    switch "a" {
    case "foo bar":
        var v17 int64 = (v1 | v1)
        _ = v17
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
