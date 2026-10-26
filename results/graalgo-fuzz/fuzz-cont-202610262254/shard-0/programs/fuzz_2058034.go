package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }

func main() {
    var v1 float32 = (-2.25 * 0.1)
    var v2 uint16 = 949
    v3 := make(map[uint32]uint32)
    v2 -= (^(v2))
    var v4 int64 = 0
    fmt.Println(v2, v2, v4)
    if false {
        fmt.Println(v1, v1)
    }
    v4 -= -690
    v5 := []int{9223372036854775807, -2147483648, 32767}
    v1 -= 0.0
    v1 = ((v1 - v1) / (v1 * (v1 / (1e10 / v1))))
    for fi6 := 0; fi6 < 7; fi6++ {
        v7 := GenS1{F0: 42, F1: (v4 % ((v4) | 1))}
        v1 = 1e10
        _ = fi6
        _ = v7
    }
    var v8 uint = 65535
    for fi9 := 0; fi9 < 6; fi9++ {
        v1 -= v1
        v10 := GenS1{F0: 127, F1: (v4 ^ (v4 % (((v4 & (v4 - v4))) | 1)))}
        _ = fi9
        _ = v10
    }
    switch "hello" {
    case "\n\t":
        if false {
            for fi11 := 0; fi11 < 20; fi11++ {
                v8 += 4294967295
                switch "foo bar" {
                case "hello":
                    var v12 uint64 = 1
                    _ = v12
                case "\n\t":
                    var v13 int32 = 42
                    _ = v13
                case "foo bar":
                    var v14 int64 = (^(v4))
                    _ = v14
                }
                _ = fi11
            }
        } else {
            var v15 uint16 = (v2 % (((v2 >> ((v8) & 63))) | 1))
            _ = v15
        }
    default:
        v3[292] = 1
    }
    var v16 uint64 = 1
    fmt.Println(v4)
    var v17 uint64 = 18446744073709551615
    v18 := GenS1{F0: 127, F1: (-(v4))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
