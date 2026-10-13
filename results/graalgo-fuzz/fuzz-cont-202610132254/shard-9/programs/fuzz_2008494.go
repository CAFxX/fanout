package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int16 }

func main() {
    v1 := []uint32{4294967295, 348, 200}
    var v2 uint32 = 2147483647
    var v3 uint32 = 2
    v4 := make(map[string]bool)
    var v5 uint64 = 839
    switch "" {
    case "a":
        var v6 int8 = 42
        _ = v6
    case "":
        if true {
            if false {
                v4[(("a" + "a") + "")] = false
            }
            if ((1.5 + 1e10) <= 0.0) {
                var v7 int16 = 1
                var v8 uint = 720
                var v9 uint64 = v5
                _ = v7
                _ = v8
                _ = v9
            } else {
                fmt.Println(v3)
            }
        }
    default:
        fmt.Println(v3)
    }
    var v10 int64 = 791
    if ((false && false) && ((792 <= 0) && (uint(v5) == 0))) {
        v4["a"] = true
        v5 = v5
        fmt.Println(v2, v2)
    } else {
        v10 -= v10
        v3 = uint32(v5)
        if true {
            v11 := GenS1{F0: v10, F1: v10, F2: 42}
            if (!false) {
                var v12 float32 = 0.1
                _ = v12
            } else {
                fmt.Println(v10, v2, v3)
                v3 -= v3
                var v13 int64 = v10
                _ = v13
            }
            _ = len(v4)
            _ = v11
        }
    }
    v10 = (v10 % ((9223372036854775807) | 1))
    var v14 bool = true
    var v15 bool = v14
    v15 = v15
    fmt.Println(v14, v5)
    if len(v1) > 0 {
        v1[0] = (v3 % (((v3 % ((v2) | 1))) | 1))
    }
    v15 = (v14 || v15)
    v10 -= -9223372036854775808
    v10 -= (v10 & (v10 << ((312) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
