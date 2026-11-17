package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
func genFn0(p0 int8, p1 int64, p2 int) int16 {
    var v5 float32 = 3.14159
    v5 = (v5 - v5)
    var v6 float64 = 3.14159
    fmt.Println(p2, v5, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -32768
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 18446744073709551615
    var v3 float32 = ((1e10 - -2.25) - 1.5)
    var v4 int32 = 398
    for fi7 := 0; fi7 < 15; fi7++ {
        v3 -= v3
        _ = fi7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := []int{-9223372036854775808, (fi8 / (((fi9 << ((255) & 63))) | 1))}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v1[1] = ""
    switch "foo bar" {
    case "":
        switch 4294967295 {
        case 4294967295:
            var v11 int8 = 1
            _ = v11
        default:
            v3 = 0.0
        }
    case "foo bar":
        v12 := []int{-1}
        _ = v12
    }
    v3 = 0.0
    var v13 int8 = -94
    v14 := GenS1{F0: v2, F1: true}
    var v15 uint32 = 65535
    switch 128 {
    case 121:
        if (!((v13 << ((252) & 63)) < v13)) {
            v15 += v15
        }
    case -497:
        v16 := GenS1{F0: (^(v2)), F1: false}
        _ = v16
    default:
        v13 -= (v13 >> (((v2 - 2)) & 63))
    }
    v17 := GenS1{F0: (v2 | v2), F1: false}
    v3 = (v3 + v3)
    for fi18 := 0; fi18 < 12; fi18++ {
        v13 = v13
        var v19 uint32 = v15
        _ = fi18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
