package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint }
func genFn0(p0 string) int32 {
    p0 = ("a" + ("foo bar" + (p0 + (("\n\t" + "a") + p0))))
    fmt.Println(p0)
    _ = p0
    return -257
}

func main() {
    v1 := []uint32{}
    var v2 int = 754
    v3 := []uint32{1, 276}
    var v4 uint = 1
    var v5 int32 = 995
    v2 += v2
    v5 = -1
    var v6 int = v2
    _ = len(v1)
    v7 := GenS1{F0: (-(v5)), F1: -9223372036854775808}
    if (4294967295 > (v6 ^ v2)) {
        for fi8 := 0; fi8 < 18; fi8++ {
            v4 += 633
            v4 += (^(v4))
            fmt.Println(v5)
            _ = fi8
        }
        v4 -= (v4 ^ (v4 | v4))
    }
    if true {
        var v9 int = v6
        v1 = append(v1, 42)
        v10 := GenS2{F0: v4}
        _ = v10
        _ = v9
    }
    _ = len(v1)
    if (!((-(v2)) < v2)) {
        v11 := GenS2{F0: (v4 << (((v4 + (v4 / ((v4) | 1)))) & 63))}
        _ = v11
    } else {
        v4 = (v4 << (((v4 ^ 2)) & 63))
    }
    var v12 uint = 1
    switch v6 {
    case 655:
        v13 := GenS1{F0: (v5 - (v5 >> ((v12) & 63))), F1: (v6 << ((255) & 63))}
        _ = v13
    default:
        v5 += genFn0("foo bar")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
