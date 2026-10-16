package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 int }
func genFn0(p0 uint64, p1 string, p2 int32) int8 {
    var v4 uint32 = 2147483647
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -99
}

func main() {
    var v1 string = "\n\t"
    var v2 float64 = 1.5
    var v3 int16 = 564
    v1 = v1
    v5 := GenS1{F0: 255, F1: v1, F2: (!true)}
    v2 = (float64(-2.25) - ((v2 / (1e10 * v2)) * v2))
    if false {
        v6 := []int{}
        if true {
            var v7 int = 127
            switch "\n\t" {
            case "\n\t":
                v8 := GenS1{F0: 29, F1: v1, F2: false}
                _ = v8
            case "":
                var v9 uint8 = 0
                _ = v9
            default:
                v10 := GenS1{F0: 255, F1: v1, F2: true}
                _ = v10
            }
            _ = v7
        }
        _ = v6
    } else {
        if true {
            v11 := GenS2{F0: 1}
            v12 := []int{2147483647, 526, 32767}
            _ = v11
            _ = v12
        }
        v13 := []int{-1, 128, -2147483648}
        _ = v13
    }
    v2 = (v2 - ((0.0 + (((float64(0.0) / v2) - float64((1e10 / 1.0))) + float64(1e10))) / v2))
    v14 := GenS1{F0: 1, F1: ("hello" + ""), F2: true}
    fmt.Println(v3, v3)
    v3 -= v3
    v15 := []int{}
    var v16 int16 = v3
    if ((2147483648 <= 4294967295) || true) {
        var v17 float64 = (v2 * 1e10)
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
}
