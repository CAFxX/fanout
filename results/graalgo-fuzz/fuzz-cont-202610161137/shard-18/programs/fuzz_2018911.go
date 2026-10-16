package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint }
func genFn0() uint64 {
    v6 := []int{-1, 4294967295, 9223372036854775807, -2147483648}
    var v7 int64 = -1
    v7 += (v7 | v7)
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int64) uint {
    v8 := []int{}
    fmt.Println(p0)
    p0 += (p0 | (^(p0)))
    for fi9 := 0; fi9 < 6; fi9++ {
        p0 += (p0 % ((p0) | 1))
        p0 = (-(p0))
        switch int(p0) {
        case -362:
            p0 -= -9223372036854775808
        case 127:
            var v10 float32 = 1e10
            _ = v10
        case 0:
            if (654 < 1) {
                fmt.Println(fi9, fi9)
                switch "a" {
                case "hello":
                    _ = len(v8)
                case "\n\t":
                    p0 += (p0 << ((42) & 63))
                default:
                    for fi11 := 0; fi11 < 8; fi11++ {
                        var v12 float64 = 3.14159
                        _ = fi11
                        _ = v12
                    }
                }
            }
        }
        _ = fi9
    }
    _ = p0
    _ = v8
    return 0
}

func main() {
    v1 := []int{363, 0}
    v2 := []uint32{249, 1, 447, 0}
    var v3 int64 = -296
    var v4 int = -938
    var v5 uint16 = 65535
    var v13 float32 = 1.5
    var v14 float64 = 3.14159
    var v15 uint64 = 99
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ("" + ("a" + ("a" + "foo bar"))) {
    case "":
        var v16 uint16 = 0
        _ = v16
    case "\n\t":
        var v17 uint8 = 0
        _ = v17
    case "foo bar":
        v3 += (v3 | v3)
    default:
        v5 -= v5
    }
    v18 := GenS1{F0: 0, F1: 596, F2: 2}
    switch (v4 ^ (v4 << ((26) & 63))) {
    case -100:
        v13 = float32((1.5 - v14))
    }
    fmt.Println(v3)
    v4 = (v4 >> ((genFn1((^(v3)))) & 63))
    var v19 uint16 = (v5 * 2)
    if ((false || (!true)) || (true || (!(!((776 == (v3 >> ((255) & 63))) || true))))) {
        var v20 uint64 = v15
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
