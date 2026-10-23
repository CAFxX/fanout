package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint, p1 string) int {
    switch 9223372036854775807 {
    case 813:
        v4 := GenS1{F0: 255, F1: 1, F2: 112}
        _ = v4
    case 0:
        p1 = (p1 + p1)
    }
    var v5 uint = 679
    var v6 int32 = 767
    v5 -= (^(v5))
    p0 -= (^(p0))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 bool, p1 uint64, p2 uint) uint64 {
    v7 := GenS2{F0: 355}
    p2 += 18446744073709551615
    switch 621 {
    case 2147483647:
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint = (p2 >> ((p2) & 63))
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 float64 = 0.1
                var v12 int16 = -32768
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
    case 0:
        v13 := GenS1{F0: 175, F1: 215, F2: p1}
        _ = v13
    case 9223372036854775807:
        var v14 bool = true
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 633
}

func main() {
    var v1 int8 = 42
    var v2 float64 = 0.1
    var v3 uint64 = 2
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 uint64 = 18446744073709551615
        if ((v1 & (v1 & v1)) >= 42) {
            v17 := GenS1{F0: 65, F1: (v16 * (v16 + genFn1(((false && false) || true), (v3 ^ (v3 >> ((0) & 63))), 700))), F2: (v3 << ((65535) & 63))}
            v1 -= (v1 << ((18446744073709551615) & 63))
            _ = v17
        } else {
            var v18 bool = (!(42 == 1))
            _ = v18
        }
        v3 += (v3 % ((v3) | 1))
        _ = fi15
        _ = v16
    }
    var v19 uint = 18446744073709551615
    v20 := []int{}
    switch 128 {
    case 32768:
        var v21 uint8 = 255
        _ = v21
    case 714:
        if false {
            fmt.Println(v2)
            v22 := GenS2{F0: 65535}
            _ = v22
        }
    }
    var v23 string = ""
    var v24 uint8 = 90
    v25 := GenS1{F0: (v24 >> (((^(v19))) & 63)), F1: (v3 >> (((v19 << ((v19) & 63))) & 63)), F2: (v3 * 1)}
    for fi26 := 0; fi26 < 2; fi26++ {
        var v27 uint16 = 255
        v28 := GenS1{F0: 0, F1: (v3 | genFn1(false, uint64(v27), 982)), F2: v3}
        _ = fi26
        _ = v27
        _ = v28
    }
    var v29 string = v23
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
}
