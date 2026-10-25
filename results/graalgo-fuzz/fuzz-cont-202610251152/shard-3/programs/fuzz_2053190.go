package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 string }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }

func main() {
    var v1 uint16 = 32767
    v2 := []string{"hello", "foo bar", "\n\t", "foo bar"}
    v3 := make(map[int]uint32)
    var v4 float32 = 1.5
    v5 := make(map[string]bool)
    v6 := make(map[int]int)
    v7 := GenS1{F0: 9223372036854775807, F1: 0, F2: ("\n\t" + "")}
    v4 += (v4 * 0.1)
    if true {
        fmt.Println(v1)
        var v8 float64 = (-2.25 + 1e10)
        v4 = v4
        _ = v8
    } else {
        var v9 int8 = 1
        if ((18446744073709551615 < 658) && (42 <= 1)) {
            fmt.Println(v4, v1, v9)
            var v10 float32 = float32(1e10)
            _ = v10
        }
        _ = v9
    }
    fmt.Println(v1)
    v1 += v1
    if true {
        v1 = v1
    } else {
        var v11 float32 = 1.5
        var v12 bool = true
        v12 = (("foo bar" + "foo bar") != "a")
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: false, F1: 9223372036854775807, F2: 0}
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(fi14, fi14, v4)
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
