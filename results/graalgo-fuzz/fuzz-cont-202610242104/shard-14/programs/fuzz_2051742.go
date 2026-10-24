package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
func genFn0(p0 int8, p1 uint16) string {
    p1 = (p1 ^ (p1 & 810))
    var v6 uint32 = 0
    v6 = 2147483648
    p0 += p0
    _ = p0
    _ = p1
    _ = v6
    return "foo bar"
}

func main() {
    var v1 int8 = -128
    var v2 int8 = -128
    v3 := []uint32{}
    v4 := make(map[string]bool)
    var v5 uint = 65535
    v7 := GenS1{F0: -9223372036854775808, F1: false}
    for fi8 := 0; fi8 < 13; fi8++ {
        switch "a" {
        case "\n\t":
            v4[""] = true
        case "a":
            var v9 float32 = (3.14159 - (0.0 + 3.14159))
            _ = v9
        default:
            v10 := GenS1{F0: 0, F1: true}
            _ = v10
        }
        _ = fi8
    }
    if false {
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    v11 := GenS1{F0: -662, F1: true}
    if true {
        v12 := GenS1{F0: 9223372036854775807, F1: (0 < 42)}
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 int64 = int64(v5)
            var v15 uint8 = 42
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = v12
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        v1 += (v2 << (((v5 / ((v5) | 1))) & 63))
    } else {
        v1 += v2
    }
    var v16 uint = v5
    v5 += v5
    var v17 int = -107
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
