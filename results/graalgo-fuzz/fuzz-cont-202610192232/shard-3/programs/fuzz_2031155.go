package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
func genFn0(p0 uint64, p1 string, p2 uint64) uint {
    var v6 int8 = 1
    if (!(!false)) {
        var v7 uint8 = 1
        if (-1 >= 1) {
            fmt.Println(p1, v6)
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 float32 = (3.14159 - (0.0 - (1e10 / 1.5)))
                p0 += p2
                _ = fi8
                _ = v9
            }
        } else {
            v6 = (v6 / (((v6 & (v6 | -128))) | 1))
        }
        v10 := []int{}
        _ = v10
        _ = v7
    } else {
        if false {
            p1 = p1
            switch 995 {
            case -1:
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := []int{}
                    _ = fi11
                    _ = v12
                }
            }
        }
    }
    fmt.Println(p0)
    v13 := []int{-1, -586, 9223372036854775807}
    var v14 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    _ = v6
    return 2
}

func main() {
    var v1 float32 = (1e10 / 3.14159)
    v2 := []uint32{523, 4294967295, 2147483647}
    v3 := []int{}
    var v4 int32 = 42
    v5 := []uint32{4294967295}
    v15 := GenS1{F0: 1, F1: 294, F2: 162}
    var v16 int = 128
    var v17 float64 = (0.1 * -2.25)
    fmt.Println(v17, v16, v4)
    switch "hello" {
    case "foo bar":
        switch "" {
        case "foo bar":
            for fi18 := 0; fi18 < 16; fi18++ {
                if len(v2) > 0 {
                    v2[0] = 65535
                }
                if (true && true) {
                    fmt.Println(v17, fi18, v16)
                }
                _ = fi18
            }
        default:
            switch "foo bar" {
            case "foo bar":
                var v19 int = (v16 * v16)
                _ = v19
            }
        }
    default:
        v2 = append(v2, 2)
    }
    var v20 float32 = v1
    v20 = ((-2.25 - v20) - (float32((v17 + ((3.14159 - v17) - v17))) + (float32(1e10) - v1)))
    var v21 float64 = v17
    if true {
        v22 := GenS1{F0: 4294967295, F1: 50, F2: 176}
        var v23 int32 = -603
        v17 = 1e10
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 uint64 = 2
        _ = fi24
        _ = v25
    }
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v4 += (v4 << ((255) & 63))
        v20 = 1.5
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
