package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
func genFn0(p0 uint64, p1 uint16) bool {
    if true {
        for fi5 := 0; fi5 < 13; fi5++ {
            for fi6 := 0; fi6 < 17; fi6++ {
                p1 += p1
                _ = fi6
            }
            _ = fi5
        }
    } else {
        var v7 int8 = 42
        v8 := GenS1{F0: uint8(p1), F1: "\n\t", F2: true}
        _ = v7
        _ = v8
    }
    var v9 bool = (-1 == 42)
    _ = p0
    _ = p1
    _ = v9
    return false
}
func genFn1(p0 bool, p1 uint32, p2 int64) int64 {
    var v10 string = "foo bar"
    fmt.Println(v10, p2, p0)
    v11 := GenS1{F0: 114, F1: v10, F2: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 uint8 = 99
    var v2 float32 = 3.14159
    var v3 uint64 = 51
    var v4 int = 128
    v12 := []int{(v4 << ((65535) & 63)), 9223372036854775807, (v4 - (v4 - 128))}
    var v13 float32 = ((v2 * 0.0) * float32(-2.25))
    fmt.Println(v2, v1, v13)
    v2 -= v2
    if len(v12) > 0 {
        v12[0] = v4
    }
    var v14 int16 = 0
    var v15 uint32 = 4294967295
    if true {
        fmt.Println(v3, v3, v4)
        if true {
            var v16 int64 = 9223372036854775807
            var v17 int = (v4 | v4)
            _ = v16
            _ = v17
        }
        if true {
            v13 = (float32(0.1) / v13)
        }
    } else {
        if (1 == uint32(v1)) {
            v12 = append(v12, 4294967295)
            if true {
                var v18 int = (v4 & (v4 % ((v4) | 1)))
                _ = v18
            } else {
                fmt.Println(v2, v15)
                v19 := GenS1{F0: (v1 ^ v1), F1: "hello", F2: false}
                _ = v19
            }
        } else {
            v4 = (v4 - v4)
            for fi20 := 0; fi20 < 4; fi20++ {
                v15 = (v15 + v15)
                v2 -= (0.0 + 1.5)
                _ = fi20
            }
        }
        v15 += 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
