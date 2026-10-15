package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }

func main() {
    v1 := []int{}
    var v2 int8 = 100
    var v3 uint = 255
    var v4 uint = v3
    v4 += 65535
    if false {
        var v5 float32 = 0.1
        _ = v5
    } else {
        v3 = (v4 ^ 18446744073709551615)
        if (float32((0.1 / 1e10)) > -2.25) {
            switch -9223372036854775808 {
            case 267:
                v6 := GenS1{F0: v2, F1: 65535, F2: ("a" + "\n\t")}
                _ = v6
            default:
                v1 = append(v1, 32767)
            }
        } else {
            fmt.Println(v2, v4, v2)
            var v7 float32 = float32(float64(0.0))
            _ = v7
        }
        v2 = v2
    }
    if len(v1) > 0 {
        v1[0] = 769
    }
    fmt.Println(v2)
    var v8 uint32 = 853
    if len(v1) > 0 {
        v1[0] = 32768
    }
    fmt.Println(v8, v4)
    var v9 float32 = (1.5 - 1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
