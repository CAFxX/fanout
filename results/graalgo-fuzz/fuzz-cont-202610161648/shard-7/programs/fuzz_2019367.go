package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint64 = 1
    var v2 uint64 = 4294967295
    var v3 bool = (1e10 >= (1e10 / 1e10))
    var v4 float32 = ((0.0 / 1e10) * 3.14159)
    switch 32768 {
    case 32768:
        var v5 int8 = 127
        _ = v5
    case 0:
        var v6 float64 = 0.1
        _ = v6
    case 1000:
        v4 = v4
    default:
        var v7 uint64 = (v1 / (((v2 | (v1 * v1))) | 1))
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        if v3 {
            v9 := GenS1{F0: uint(v1), F1: int16(v1)}
            _ = v9
        } else {
            v10 := GenS1{F0: 316, F1: 1000}
            _ = v10
        }
        v11 := []int{127, (fi8 % (((fi8 << ((0) & 63))) | 1)), fi8, (fi8 - 128)}
        fmt.Println(v1)
        _ = fi8
        _ = v11
    }
    var v12 int32 = 2147483647
    v13 := []int{-1, 917, 478, 2147483647}
    fmt.Println(v3, v1, v1)
    v1 = 4294967295
    if (v2 == 42) {
        if true {
            if false {
                var v14 bool = (32767 != 2)
                v14 = v3
                _ = v14
            } else {
                v12 = -1
                switch "" {
                case "":
                    var v15 int16 = -476
                    _ = v15
                case "a":
                    v1 = v2
                }
                var v16 int8 = 36
                _ = v16
            }
            for fi17 := 0; fi17 < 3; fi17++ {
                for fi18 := 0; fi18 < 4; fi18++ {
                    var v19 string = ""
                    var v20 int8 = 22
                    var v21 int64 = 9223372036854775807
                    v19 = "hello"
                    fmt.Println(v19, fi17, v2)
                    fmt.Println(v2, fi17)
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                var v22 int32 = (-(v12))
                if v3 {
                    fmt.Println(v12, v12)
                    var v23 uint8 = 42
                    _ = v23
                }
                _ = fi17
                _ = v22
            }
            v13 = append(v13, -1)
        }
        var v24 int = 128
        for fi25 := 0; fi25 < 1; fi25++ {
            v1 = 0
            fmt.Println(v2)
            _ = fi25
        }
        _ = v24
    } else {
        if len(v13) > 0 {
            v13[0] = 42
        }
        var v26 int = -812
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
