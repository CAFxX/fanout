package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int64 }
type GenS2 struct { F0 int64; F1 int64; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 434
    v3 := make(map[string]int)
    var v4 uint16 = 42
    v2 += 0
    var v5 bool = (!true)
    v2 = (v2 | (v2 << ((255) & 63)))
    var v6 float32 = (float32((((1.5 / 1.0) - 1.5) / 0.1)) - 0.1)
    if ((127 != -128) || ((^(v4)) < v4)) {
        fmt.Println(v4, v4, v2)
        switch 2147483647 {
        case 4294967295:
            if v5 {
                var v7 int16 = 32767
                _ = len(v3)
                _ = v7
            }
        case 332:
            _ = len(v1)
        case 0:
            delete(v1, 4294967295)
        default:
            switch 32768 {
            case 760:
                var v8 int8 = -1
                _ = v8
            case -9223372036854775808:
                v9 := GenS1{F0: 4294967295, F1: 1, F2: 42}
                _ = v9
            case -465:
                for fi10 := 0; fi10 < 15; fi10++ {
                    fmt.Println(v2, v6)
                    _ = fi10
                }
            }
        }
        var v11 float64 = (1.5 / 1e10)
        _ = v11
    } else {
        v12 := GenS2{F0: -9223372036854775808, F1: 560, F2: 42}
        for fi13 := 0; fi13 < 20; fi13++ {
            v6 -= v6
            _ = fi13
        }
        switch ("\n\t" + "") {
        case "foo bar":
            _ = len(v3)
        case "":
            var v14 float32 = float32((((1.5 / 1.5) + 0.1) - 1.5))
            _ = v14
        }
        _ = v12
    }
    fmt.Println(v6, v2, v2)
    if (true && v5) {
        v15 := []int{-1, -989, 127, 0}
        _ = v15
    } else {
        v3["foo bar"] = 2147483647
        var v16 int8 = 127
        v16 = (v16 + v16)
        _ = v16
    }
    v6 = v6
    var v17 float32 = (v6 / (float32(1.5) + v6))
    if (false || (546 == 32767)) {
        v6 = 0.0
    }
    v4 -= (v4 << ((1) & 63))
    var v18 int64 = -1
    var v19 float64 = 1e10
    var v20 uint16 = (v4 * uint16(v2))
    switch "\n\t" {
    case "a":
        fmt.Println(v5)
    case "":
        v21 := GenS1{F0: 32767, F1: 1000, F2: v18}
        _ = v21
    default:
        var v22 uint64 = (v2 % (((v2 + (v2 >> ((0) & 63)))) | 1))
        _ = v22
    }
    var v23 int32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
