package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 0
    var v3 float64 = (1.5 - 1.5)
    v4 := []uint32{541, 0, 4294967295}
    var v5 bool = (!true)
    if (372 < 325) {
        v6 := GenS1{F0: v2}
        _ = v6
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v1)
    v7 := GenS2{F0: 65535, F1: "hello", F2: -1}
    var v8 uint16 = 638
    fmt.Println(v3, v3)
    if v5 {
        var v9 int32 = -265
        v3 += v3
        _ = v9
    } else {
        v10 := GenS2{F0: (v8 >> ((4294967295) & 63)), F1: "foo bar", F2: 9223372036854775807}
        if (!true) {
            v8 = v8
            fmt.Println(v8)
        }
        _ = v10
    }
    v11 := GenS1{F0: (v2 * (v2 ^ (v2 % ((-1) | 1))))}
    v5 = (v5 && false)
    v3 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
