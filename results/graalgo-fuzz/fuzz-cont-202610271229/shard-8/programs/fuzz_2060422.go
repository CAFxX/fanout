package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    for fi5 := 0; fi5 < 15; fi5++ {
        switch (fi5 << ((4294967295) & 63)) {
        case 469:
            var v6 int32 = v1
            _ = v6
        case 473:
            var v7 int32 = 2147483647
            _ = v7
        case 83:
            for fi8 := 0; fi8 < 11; fi8++ {
                fmt.Println(fi5)
                switch "\n\t" {
                case "foo bar":
                    var v9 int32 = 1
                    _ = v9
                case "hello":
                    switch "" {
                    case "hello":
                        var v10 uint16 = 1
                        _ = v10
                    case "":
                        v1 = v1
                    case "a":
                        v1 = 1
                    }
                case "":
                    var v11 int64 = int64(v1)
                    _ = v11
                default:
                    fmt.Println(v1)
                }
                _ = fi8
            }
        default:
            v1 = v1
        }
        v12 := []int{fi5, -940, fi5, (fi5 * fi5)}
        _ = fi5
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        if (true || false) {
            var v14 uint = 2
            v1 -= v1
            _ = v14
        }
        if (false || true) {
            v1 += 2147483647
            var v15 float64 = (((0.1 / 1e10) + (-2.25 + float64((0.0 / (1.5 * (0.1 - 1.5)))))) + 0.1)
            _ = v15
        }
        _ = fi13
    }
    if false {
        var v16 uint = 4294967295
        _ = v16
    } else {
        v1 -= v1
        v1 += 1
    }
    v17 := GenS1{F0: 0, F1: 2147483648, F2: 842}
    var v18 int32 = 42
    v18 = v1
    v19 := []int{-2147483648, 689, 128}
    switch 42 {
    case -2147483648:
        v18 += (v18 - (v1 ^ v1))
    case 42:
        v20 := GenS1{F0: 825, F1: 2147483647, F2: 2}
        _ = v20
    case 4294967295:
        v21 := GenS1{F0: 32767, F1: 2147483648, F2: 32767}
        _ = v21
    }
    v1 = v1
    v1 -= v18
    fmt.Println(v1)
    v22 := GenS1{F0: -789, F1: 42, F2: 255}
    v18 -= (v18 / ((v1) | 1))
    var v23 uint32 = 897
    var v24 uint16 = 991
    v25 := GenS2{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
