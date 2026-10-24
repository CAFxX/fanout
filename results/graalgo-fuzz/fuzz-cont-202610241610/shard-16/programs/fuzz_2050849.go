package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }
type GenS2 struct { F0 int32; F1 bool; F2 int32 }
func genFn0(p0 int64, p1 int32) string {
    v5 := GenS1{F0: (p0 + p0), F1: 127, F2: 715}
    if false {
        p1 -= p1
        fmt.Println(p1, p0)
    }
    var v6 uint16 = 2
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    var v1 uint16 = 82
    var v2 uint16 = (v1 * v1)
    var v3 int64 = -1
    var v4 uint = 1
    if true {
        v7 := []int{-606, 32767}
        var v8 uint = v4
        _ = v7
        _ = v8
    }
    v2 = (v2 - (v2 - (v2 % (((v2 / ((v2) | 1))) | 1))))
    fmt.Println(v4)
    if true {
        var v9 uint16 = (v1 ^ v2)
        var v10 uint32 = 0
        v11 := []int{926}
        _ = v10
        _ = v11
        _ = v9
    } else {
        switch "a" {
        case "foo bar":
            fmt.Println(v2, v2)
        default:
            var v12 float32 = 0.0
            _ = v12
        }
    }
    v13 := GenS1{F0: (v3 ^ 9223372036854775807), F1: 1, F2: v4}
    fmt.Println(v1, v3, v3)
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            if ((^(v4)) == 65535) {
                v3 += int64(v1)
            } else {
                v16 := []int{(fi15 | (fi15 / ((fi14) | 1))), (fi15 * fi14), 1, 4294967295}
                _ = v16
            }
            var v17 float32 = 1.5
            _ = fi15
            _ = v17
        }
        v1 -= v1
        _ = fi14
    }
    fmt.Println(v2)
    v18 := GenS2{F0: 42, F1: (1.5 <= ((1e10 + 0.1) + 0.0)), F2: 1}
    if false {
        v4 += v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
