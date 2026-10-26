package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int16 = 42
    v2 := []int{}
    var v3 float64 = 1.5
    var v4 uint32 = 2147483647
    v5 := GenS1{F0: 2147483647, F1: 2147483647, F2: 361}
    if false {
        v6 := GenS2{F0: 902}
        v4 = (v4 + 0)
        _ = v6
    } else {
        v7 := GenS1{F0: -1, F1: (v4 >> ((255) & 63)), F2: 42}
        _ = v7
    }
    switch "a" {
    case "hello":
        v3 += 0.1
    case "foo bar":
        switch "\n\t" {
        case "a":
            var v8 float64 = 1.5
            _ = v8
        case "foo bar":
            v2 = append(v2, -1)
        case "":
            v3 -= v3
        default:
            if false {
                v3 -= ((v3 / v3) * v3)
                v1 -= (v1 << ((42) & 63))
            }
        }
    default:
        var v9 uint = 52
        _ = v9
    }
    var v10 int = 1
    v3 = v3
    var v11 int8 = 1
    v1 -= 32767
    switch "foo bar" {
    case "hello":
        if (!(((1e10 + 1e10) * (float32(1e10) + 1e10)) >= ((float32(1.5) * 1.5) / 1.0))) {
            _ = len(v2)
            v12 := GenS2{F0: 1}
            _ = v12
        }
    case "foo bar":
        fmt.Println(v1, v11, v4)
    default:
        for fi13 := 0; fi13 < 3; fi13++ {
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 uint64 = 2
                if (1.5 != -2.25) {
                    fmt.Println(v10, v4)
                }
                v16 := GenS2{F0: (^(v15))}
                _ = fi14
                _ = v15
                _ = v16
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                fmt.Println(v10, v3)
                _ = fi17
            }
            _ = fi13
        }
    }
    v3 = (0.1 * v3)
    v3 += (float64(0.0) / 1.0)
    fmt.Println(v11)
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 uint16 = 271
        _ = len(v2)
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
