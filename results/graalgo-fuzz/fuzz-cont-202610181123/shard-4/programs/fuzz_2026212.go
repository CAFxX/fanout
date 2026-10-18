package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
func genFn0(p0 string) uint16 {
    if ((((1 <= 1) || (!(467 != 2))) && true) || (4294967295 == 935)) {
        p0 = (p0 + p0)
    } else {
        if (-477 > 1) {
            p0 = p0
        } else {
            switch "foo bar" {
            case "foo bar":
                v4 := []int{-9223372036854775808}
                _ = v4
            }
        }
    }
    var v5 float64 = float64((0.0 - 0.0))
    var v6 uint64 = 4294967295
    _ = p0
    _ = v5
    _ = v6
    return 235
}
func genFn1(p0 uint64) int {
    p0 = p0
    p0 += p0
    p0 = 1
    v7 := GenS1{F0: (-560 <= -2147483648), F1: 127}
    fmt.Println(p0)
    _ = p0
    _ = v7
    return -173
}

func main() {
    var v1 int = -1
    var v2 int16 = 467
    var v3 int16 = v2
    var v8 uint64 = 42
    if (v1 < (v1 >> ((uint(v8)) & 63))) {
        fmt.Println(v3)
        fmt.Println(v1)
        var v9 bool = (((-1 != 0) || (-2.25 > 1e10)) && (false || false))
        _ = v9
    }
    v10 := []int{}
    if (uint32(v8) >= 42) {
        v3 = (v3 << ((421) & 63))
        var v11 string = "\n\t"
        fmt.Println(v1, v11, v2)
        _ = v11
    }
    v8 += v8
    v8 = v8
    v3 = (v2 ^ (v2 ^ v2))
    if (true || true) {
        v12 := GenS1{F0: (("foo bar" + ("\n\t" + ("" + "a"))) == ""), F1: -105}
        for fi13 := 0; fi13 < 13; fi13++ {
            v1 -= genFn1(2)
            fmt.Println(v1, v3, fi13)
            _ = fi13
        }
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v8
}
