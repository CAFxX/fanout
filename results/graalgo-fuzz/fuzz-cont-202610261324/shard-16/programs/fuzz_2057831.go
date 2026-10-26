package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint16, p1 int8, p2 uint8) uint {
    if true {
        var v7 int64 = -1
        v8 := []int{}
        v9 := GenS1{F0: 431, F1: (p2 << ((4294967295) & 63))}
        _ = v7
        _ = v8
        _ = v9
    }
    switch "\n\t" {
    case "foo bar":
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int = fi10
            _ = fi10
            _ = v11
        }
    case "a":
        if true {
            p2 -= (p2 % ((p2) | 1))
        }
    }
    p1 -= (p1 % (((p1 + (p1 & (^(p1))))) | 1))
    for fi12 := 0; fi12 < 12; fi12++ {
        v13 := []int{}
        p0 = (p0 & p0)
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := []int{(fi14 - fi15), (fi15 >> ((347) & 63)), (fi15 + fi15), int(p0)}
            _ = len(v16)
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := []uint32{221, 65535, 1, 0}
    var v2 int64 = -43
    v3 := []uint32{65535, 574, 42, 0}
    v4 := []int{}
    var v5 uint64 = 18446744073709551615
    var v6 uint64 = (v5 / (((v5 & v5)) | 1))
    if false {
        fmt.Println(v6)
    } else {
        v2 = v2
        var v17 int32 = 2147483647
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            v5 += v5
            v6 = 2
            var v18 int16 = -32768
            _ = v18
        }
        _ = v17
    }
    var v19 float64 = (0.1 + (((0.0 * 1.5) / 1e10) + float64(0.1)))
    v20 := GenS1{F0: 42, F1: 127}
    for fi21 := 0; fi21 < 19; fi21++ {
        fmt.Println(v6, v5)
        var v22 uint8 = 127
        var v23 uint32 = 588
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: 32767, F1: 128}
    _ = len(v4)
    var v25 int = 9223372036854775807
    for fi26 := 0; fi26 < 11; fi26++ {
        v2 = (v2 - (v2 / (((v2 % (((-(v2))) | 1))) | 1)))
        for fi27 := 0; fi27 < 1; fi27++ {
            if ((false && false) && (560 < -835)) {
                var v28 uint32 = 1
                v25 -= (v25 % ((fi26) | 1))
                _ = v28
            }
            v6 = (v6 & (v5 - v6))
            _ = fi27
        }
        fmt.Println(v2)
        _ = fi26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
