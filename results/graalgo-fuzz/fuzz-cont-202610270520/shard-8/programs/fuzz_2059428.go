package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int }

func main() {
    var v1 bool = (!true)
    v2 := make(map[int]bool)
    var v3 uint32 = 4294967295
    var v4 int64 = 537
    v5 := GenS1{F0: 42, F1: 2}
    var v6 int = 127
    if (1e10 < 1.5) {
        fmt.Println(v4, v6, v3)
    } else {
        switch "a" {
        case "":
            switch "\n\t" {
            case "":
                v7 := []int{v6, v6, (v6 - (v6 / ((v6) | 1)))}
                _ = v7
            case "\n\t":
                v8 := GenS1{F0: 32767, F1: 255}
                _ = v8
            default:
                var v9 uint64 = uint64(v3)
                _ = v9
            }
        }
    }
    if (!true) {
        switch (v6 << ((1) & 63)) {
        case 32768:
            if (0 == uint(v3)) {
                if (v1 || v1) {
                    var v10 uint = 469
                    _ = v10
                }
                if false {
                    switch "" {
                    case "\n\t":
                        for fi11 := 0; fi11 < 4; fi11++ {
                            v12 := GenS1{F0: 1, F1: 0}
                            v6 -= -405
                            v6 = (v6 / ((v6) | 1))
                            _ = fi11
                            _ = v12
                        }
                    case "a":
                        for fi13 := 0; fi13 < 19; fi13++ {
                            var v14 float64 = float64((1e10 / ((float32(((0.1 / -2.25) + 0.1)) / 3.14159) - 3.14159)))
                            _ = fi13
                            _ = v14
                        }
                    case "foo bar":
                        var v15 int32 = 2147483647
                        _ = v15
                    default:
                        var v16 float64 = -2.25
                        _ = v16
                    }
                }
                v17 := GenS1{F0: 65535, F1: 0}
                _ = v17
            }
        case 2147483647:
            var v18 uint16 = 65535
            _ = v18
        case 42:
            if v1 {
                v19 := GenS2{F0: v6}
                _ = v19
            }
        }
    } else {
        v2[v6] = true
        var v20 uint = 18446744073709551615
        _ = v20
    }
    v21 := GenS1{F0: 32767, F1: 4294967295}
    var v22 int64 = v4
    var v23 uint64 = 460
    var v24 uint64 = (v23 & v23)
    v25 := []int{}
    v25 = append(v25, (v6 << ((255) & 63)))
    v26 := GenS1{F0: 674, F1: 1}
    var v27 float64 = float64((1.5 / 1e10))
    var v28 float64 = v27
    v25 = append(v25, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
