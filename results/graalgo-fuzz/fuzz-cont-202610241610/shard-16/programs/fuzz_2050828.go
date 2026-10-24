package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int8; F1 int64; F2 uint64 }
func genFn0(p0 uint16, p1 string, p2 int32) uint {
    if true {
        var v7 int8 = 72
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{fi8, int(p2), (-(fi8)), -501}
            switch -2147483648 {
            case -442:
                p2 = p2
            case -825:
                var v10 uint64 = 18446744073709551615
                _ = v10
            case 32767:
                v11 := GenS1{F0: -32768, F1: 264}
                _ = v11
            default:
                if (!false) {
                    p0 = p0
                } else {
                    v12 := GenS1{F0: 1, F1: (fi8 % (((fi8 >> ((971) & 63))) | 1))}
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    var v13 int32 = p2
                    _ = v12
                    _ = v13
                }
            }
            _ = fi8
            _ = v9
        }
        p2 = (^(p2))
        _ = v7
    } else {
        var v14 int = 32767
        _ = v14
    }
    if false {
        fmt.Println(p0)
        v15 := GenS1{F0: -47, F1: 9223372036854775807}
        if (!false) {
            fmt.Println(p1, p1)
            p1 = p1
            for fi16 := 0; fi16 < 20; fi16++ {
                var v17 int16 = 42
                if true {
                    for fi18 := 0; fi18 < 14; fi18++ {
                        p0 -= (p0 + (^(p0)))
                        _ = fi18
                    }
                    var v19 float32 = 0.1
                    _ = v19
                }
                _ = fi16
                _ = v17
            }
        }
        _ = v15
    }
    var v20 float64 = 1.5
    var v21 int32 = (p2 - p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 65535
}

func main() {
    var v1 int64 = -605
    v2 := []int{128}
    var v3 float64 = 0.0
    v4 := []int{}
    var v5 uint32 = 2
    var v6 uint = 16
    v5 = v5
    v22 := GenS2{F0: -128, F1: 63, F2: 42}
    v6 -= (v6 ^ (^(v6)))
    switch 729 {
    case 32767:
        v1 = (v1 << ((v6) & 63))
    }
    _ = len(v2)
    v5 += v5
    var v23 float64 = v3
    switch "\n\t" {
    case "":
        var v24 int8 = -128
        _ = v24
    case "\n\t":
        v1 = 42
    case "a":
        var v25 uint8 = 50
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
