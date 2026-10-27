package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int64; F1 int; F2 int16 }
type GenS3 struct { F0 uint8 }

func main() {
    v1 := []uint32{496, 191, 601}
    var v2 uint32 = 452
    v3 := make(map[string]bool)
    var v4 int32 = 1
    switch 2147483647 {
    case -443:
        if (((1e10 + 0.1) * 0.1) < (1.5 + (1.5 / 1e10))) {
            v5 := GenS1{F0: (255 <= 176), F1: true}
            v6 := GenS2{F0: 1, F1: 0, F2: -483}
            v4 = v4
            _ = v5
            _ = v6
        }
    default:
        v4 = (v4 + v4)
    }
    switch "a" {
    case "foo bar":
        v4 = (v4 >> ((464) & 63))
    case "hello":
        v7 := GenS1{F0: true, F1: (true && true)}
        _ = v7
    default:
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(fi8)
            _ = fi8
        }
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v2 = (v2 & (v2 % ((v2) | 1)))
        _ = fi9
    }
    if true {
        for fi10 := 0; fi10 < 3; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                v4 -= v4
                if (((-2.25 / 1.5) / (((3.14159 / 1.0) * 1e10) * 1e10)) != float64(-2.25)) {
                    v12 := GenS2{F0: 645, F1: fi11, F2: int16(v4)}
                    _ = v12
                }
                _ = fi11
            }
            _ = fi10
        }
    }
    if (!false) {
        v13 := GenS1{F0: (!(!true)), F1: (!(18446744073709551615 != 823))}
        v4 += 42
        _ = v13
    }
    v14 := GenS1{F0: (!false), F1: (true && true)}
    var v15 int8 = 47
    var v16 string = "a"
    var v17 int8 = v15
    v15 = 106
    fmt.Println(v17)
    var v18 bool = false
    v19 := GenS3{F0: 127}
    switch 545 {
    case -132:
        fmt.Println(v16, v2, v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
