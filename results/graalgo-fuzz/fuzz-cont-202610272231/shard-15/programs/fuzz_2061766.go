package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 bool }

func main() {
    var v1 int64 = 335
    var v2 uint32 = 4294967295
    v3 := make(map[int]bool)
    var v4 float64 = (1e10 - ((1.5 + (1.5 + 1e10)) / 1e10))
    if (!(125 == 100)) {
        v5 := []int{42, -849, 42}
        v6 := GenS1{F0: -1}
        _ = v5
        _ = v6
    } else {
        if (v2 < (v2 ^ v2)) {
            if true {
                v4 += v4
            } else {
                switch 461 {
                case 9223372036854775807:
                    var v7 int64 = 0
                    _ = v7
                case -300:
                    v2 += 2147483647
                }
                _ = len(v3)
                fmt.Println(v1, v4, v1)
            }
        } else {
            if false {
                var v8 int16 = 32767
                v9 := v3[32767]
                _ = v8
                _ = v9
            } else {
                var v10 int16 = -32768
                fmt.Println(v4, v4, v2)
                _ = v10
            }
        }
        fmt.Println(v4, v1)
    }
    v11 := []int{824}
    v4 -= float64(0.1)
    var v12 uint8 = 178
    if (18446744073709551615 > 18446744073709551615) {
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 float64 = ((v4 / (0.1 + v4)) * v4)
            _ = fi13
            _ = v14
        }
        if false {
            switch -1 {
            case -9223372036854775808:
                v1 -= (v1 ^ (v1 - v1))
            }
        } else {
            if len(v11) > 0 {
                v11[0] = 42
            }
            switch int(v1) {
            case 41:
                var v15 uint32 = v2
                _ = v15
            case 32768:
                v3[32768] = true
            case -701:
                fmt.Println(v1, v1, v12)
            }
        }
    } else {
        fmt.Println(v1, v1)
        var v16 int8 = 42
        _ = v16
    }
    fmt.Println(v2, v2, v12)
    v12 -= (v12 + 0)
    v2 = (v2 + v2)
    if (2 >= 42) {
        var v17 float64 = 1.5
        v1 += 42
        _ = v17
    }
    if true {
        v18 := GenS1{F0: -128}
        v19 := GenS2{F0: v12, F1: int64(v2)}
        fmt.Println(v1, v4, v12)
        _ = v18
        _ = v19
    }
    if (!true) {
        var v20 bool = (v4 <= v4)
        v21 := GenS3{F0: (v2 > v2)}
        _ = v20
        _ = v21
    }
    switch 356 {
    case -404:
        fmt.Println(v1, v12)
    case 128:
        _ = len(v3)
    default:
        v22 := v3[-795]
        _ = v22
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
