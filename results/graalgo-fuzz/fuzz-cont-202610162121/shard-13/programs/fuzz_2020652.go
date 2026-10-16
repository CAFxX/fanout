package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int64 }
func genFn0(p0 uint, p1 int64, p2 int16) int64 {
    p1 = (^(p1))
    var v5 uint16 = 0
    var v6 float32 = (1e10 * 0.1)
    var v7 int32 = 150
    var v8 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = -813
    var v3 uint64 = 975
    var v4 uint = 852
    switch v2 {
    case -640:
        v2 = v2
    case 128:
        var v9 int16 = -1
        _ = v9
    }
    v10 := []int{}
    switch "a" {
    case "":
        var v11 int = v2
        _ = v11
    case "a":
        v12 := GenS1{F0: -244, F1: 100, F2: -195}
        _ = v12
    }
    fmt.Println(v4)
    var v13 uint8 = 1
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 13; fi15++ {
            v13 = v13
            switch ("\n\t" + "a") {
            case "":
                v2 += 4294967295
            case "hello":
                _ = len(v1)
            case "a":
                v13 -= v13
            default:
                if (uint(v13) != v4) {
                    v2 -= (fi15 & (-(fi14)))
                    v4 = (^(v4))
                }
            }
            fmt.Println(fi14)
            _ = fi15
        }
        _ = fi14
    }
    v4 = (v4 / ((v4) | 1))
    _ = len(v10)
    v13 -= 2
    var v16 int8 = -70
    fmt.Println(v13, v4, v13)
    var v17 int = v2
    if (0 <= 32767) {
        v4 = (v4 >> ((300) & 63))
        var v18 uint8 = (v13 ^ (v13 << (((v4 << ((857) & 63))) & 63)))
        if len(v10) > 0 {
            v10[0] = 9223372036854775807
        }
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
