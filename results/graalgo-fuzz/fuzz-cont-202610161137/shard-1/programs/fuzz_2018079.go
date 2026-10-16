package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int16 }
func genFn0(p0 int8) int8 {
    fmt.Println(p0)
    var v4 int = 780
    v5 := GenS1{F0: v4, F1: (p0 + (p0 | p0)), F2: 32767}
    v4 = v4
    if (!((int(p0) <= v4) && (!false))) {
        v6 := []int{v4, (-(v4)), v4}
        _ = v6
    } else {
        v7 := GenS1{F0: (^(v4)), F1: (p0 * p0), F2: 886}
        _ = v7
    }
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 int8 = -101
    var v2 uint = 65535
    var v3 int = -9223372036854775808
    fmt.Println(v2, v1, v1)
    v8 := []int{v3}
    v1 += (v1 | v1)
    v3 += (v3 % (((v3 & (v3 % ((9223372036854775807) | 1)))) | 1))
    if (float64(((3.14159 * 1e10) / (float32(0.0) + 0.1))) == (0.1 - -2.25)) {
        if (!(2 < 4294967295)) {
            v1 += (v1 + (v1 >> ((v2) & 63)))
            fmt.Println(v3, v3, v2)
        }
    } else {
        var v9 uint64 = 2
        if true {
            v3 += 127
        }
        var v10 int64 = -1
        _ = v10
        _ = v9
    }
    _ = len(v8)
    var v11 bool = (!true)
    var v12 string = (("a" + "foo bar") + "foo bar")
    if len(v8) > 0 {
        v8[0] = -9223372036854775808
    }
    var v13 uint = 205
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
}
