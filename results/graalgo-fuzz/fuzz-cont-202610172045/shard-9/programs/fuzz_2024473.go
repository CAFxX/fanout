package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int; F2 int16 }

func main() {
    v1 := []uint32{2147483647}
    var v2 uint16 = 0
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v5 := GenS1{F0: -182, F1: 1}
    v2 -= 32767
    if true {
        v6 := v3["foo bar"]
        var v7 uint = 185
        _ = v6
        _ = v7
    } else {
        var v8 uint8 = 127
        _ = v8
    }
    if false {
        _ = len(v4)
    }
    v9 := GenS1{F0: 0, F1: 2147483648}
    if (2 > 242) {
        delete(v3, "hello")
        v2 = 421
        for fi10 := 0; fi10 < 12; fi10++ {
            if len(v1) > 0 {
                v1[0] = 431
            }
            _ = fi10
        }
    } else {
        fmt.Println(v2)
        var v11 bool = ((!((3.14159 * 1e10) <= 1.5)) && (0 < -1))
        fmt.Println(v11)
        _ = v11
    }
    v3[""] = (!true)
    fmt.Println(v2)
    switch "\n\t" {
    case "a":
        switch "hello" {
        case "hello":
            var v12 uint64 = 181
            _ = v12
        case "a":
            v2 += 32767
        case "\n\t":
            for fi13 := 0; fi13 < 17; fi13++ {
                v2 = 0
                v2 = v2
                for fi14 := 0; fi14 < 17; fi14++ {
                    v15 := GenS2{F0: 65535, F1: fi13, F2: 0}
                    _ = fi14
                    _ = v15
                }
                _ = fi13
            }
        default:
            if true {
                var v16 uint8 = 110
                _ = v16
            }
        }
    case "\n\t":
        var v17 int64 = 42
        _ = v17
    case "hello":
        v2 += (v2 % ((v2) | 1))
    }
    v18 := GenS2{F0: 128, F1: 127, F2: 0}
    var v19 int64 = -129
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 float32 = float32((0.1 - -2.25))
        var v22 bool = false
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        v2 += (v2 ^ v2)
        v24 := GenS2{F0: 1, F1: (fi23 << ((255) & 63)), F2: -798}
        _ = fi23
        _ = v24
    }
    v19 = (v19 * (-(v19)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
