package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{1, 65535}
    v3 := []string{("" + ("" + "hello")), "", "\n\t", (("foo bar" + "\n\t") + "foo bar")}
    var v4 uint16 = 42
    var v5 uint16 = (^(v4))
    _ = len(v1)
    var v6 int32 = 42
    v7 := GenS1{F0: 1, F1: uint8(v4), F2: 1}
    if ((true || (!false)) && true) {
        v8 := GenS1{F0: 32767, F1: 42, F2: -496}
        fmt.Println(v5, v5, v6)
        fmt.Println(v6)
        _ = v8
    }
    v9 := GenS1{F0: 0, F1: 208, F2: 869}
    var v10 uint64 = 2
    v2 = append(v2, 2147483647)
    var v11 int64 = -9223372036854775808
    var v12 float64 = float64(float32(3.14159))
    if false {
        for fi13 := 0; fi13 < 16; fi13++ {
            fmt.Println(v4)
            v14 := GenS1{F0: 1, F1: 1, F2: fi13}
            _ = fi13
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 10; fi15++ {
            for fi16 := 0; fi16 < 9; fi16++ {
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                v17 := GenS1{F0: 1000, F1: 42, F2: (fi15 << ((1) & 63))}
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        v3 = append(v3, "hello")
    }
    switch "" {
    case "a":
        var v18 int64 = v11
        _ = v18
    case "":
        var v19 float64 = 1e10
        _ = v19
    default:
        var v20 int64 = 463
        _ = v20
    }
    v5 = (^(v5))
    var v21 uint64 = (v10 % ((4294967295) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
