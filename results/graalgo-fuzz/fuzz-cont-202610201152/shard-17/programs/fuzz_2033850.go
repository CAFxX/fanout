package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int; F1 uint64; F2 uint8 }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]string)
    var v3 float64 = (0.0 / 0.1)
    v4 := make(map[string]uint32)
    v5 := GenS1{F0: 0, F1: -1}
    switch int(v1) {
    case 1:
        switch 982 {
        case 32767:
            v6 := GenS1{F0: 1, F1: -9223372036854775808}
            _ = v6
        }
    case -632:
        v7 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
        _ = v7
    default:
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint16 = 541
            _ = fi8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 float64 = (float64(1.5) * v3)
        switch 723 {
        case -17:
            v12 := GenS1{F0: -538, F1: 9223372036854775807}
            _ = v12
        case 398:
            fmt.Println(fi10, fi10)
        case 32767:
            v3 = (((0.1 * (v3 / 3.14159)) - float64(-2.25)) - 0.0)
        default:
            var v13 bool = false
            _ = v13
        }
        fmt.Println(v1, v3)
        _ = fi10
        _ = v11
    }
    _ = len(v2)
    v14 := []int{127}
    if false {
        v15 := GenS2{F0: 1, F1: 18446744073709551615, F2: 0}
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 string = ("" + "")
            v3 = (v3 + v3)
            _ = fi16
            _ = v17
        }
        delete(v2, "a")
        _ = v15
    } else {
        var v18 uint32 = 564
        _ = v18
    }
    fmt.Println(v3, v3)
    var v19 int64 = -9223372036854775808
    v20 := GenS1{F0: 42, F1: (v19 + (v19 + v19))}
    var v21 float32 = 3.14159
    fmt.Println(v19)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
