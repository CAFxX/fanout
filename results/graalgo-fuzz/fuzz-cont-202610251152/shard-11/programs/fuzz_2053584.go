package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64, p1 uint64) int64 {
    p0 = 0
    v5 := GenS1{F0: 1, F1: 2147483648, F2: -418}
    v6 := []int{-9223372036854775808}
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 -= p0
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        p1 -= (p1 - (p1 * (^(p1))))
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 0
    v3 := make(map[string]uint32)
    v4 := make(map[int]string)
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := GenS1{F0: 268, F1: 2147483648, F2: fi8}
        var v10 float32 = 1.5
        for fi11 := 0; fi11 < 2; fi11++ {
            v2 -= v2
            _ = fi11
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 -= (v2 & v2)
    _ = len(v4)
    var v12 uint = 381
    v13 := []int{2147483647, 9223372036854775807, 32768, -9223372036854775808}
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 float64 = 1e10
        v16 := GenS2{F0: 1}
        _ = fi14
        _ = v15
        _ = v16
    }
    if false {
        fmt.Println(v12)
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 20; fi18++ {
                fmt.Println(v12, v12)
                switch fi17 {
                case 603:
                    v12 = v12
                case -2147483648:
                    var v19 float32 = 1e10
                    _ = v19
                case -1:
                    v13 = append(v13, -1)
                }
                switch "a" {
                case "":
                    var v20 int64 = 9223372036854775807
                    _ = v20
                default:
                    var v21 int64 = 1
                    _ = v21
                }
                _ = fi18
            }
            v22 := GenS1{F0: 2, F1: 2147483647, F2: int(v2)}
            _ = fi17
            _ = v22
        }
        var v23 float32 = 0.1
        _ = v23
    }
    var v24 uint16 = 1
    v25 := GenS2{F0: 0}
    if (!true) {
        v12 -= v12
        fmt.Println(v2, v12)
    }
    var v26 float32 = (1.5 + 1e10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
