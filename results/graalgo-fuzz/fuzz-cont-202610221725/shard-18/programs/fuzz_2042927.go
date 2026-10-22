package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }

func main() {
    v1 := []string{("" + ("" + (("a" + "foo bar") + "foo bar"))), "", (("foo bar" + "\n\t") + (("\n\t" + "") + "a")), "foo bar"}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    if (0 >= 1) {
        v4 := v3[32767]
        _ = v4
    } else {
        if (!(-1 > 127)) {
            var v5 uint16 = 0
            _ = v5
        }
    }
    var v6 int32 = -890
    v6 = (v6 / ((v6) | 1))
    fmt.Println(v6)
    v7 := GenS1{F0: "hello", F1: 4294967295, F2: -595}
    for fi8 := 0; fi8 < 6; fi8++ {
        v6 -= (-(v6))
        _ = len(v3)
        var v9 int8 = 28
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: "foo bar", F1: 4294967295, F2: -18}
    fmt.Println(v6)
    var v11 uint = 42
    if (104 != -32768) {
        fmt.Println(v11, v11)
        v11 = v11
    }
    var v12 float32 = float32(float64((1.5 * 1.5)))
    if true {
        v12 = 1.5
        fmt.Println(v6, v11, v11)
        var v13 int32 = v6
        _ = v13
    } else {
        v14 := GenS2{F0: 42}
        v15 := GenS3{F0: -1}
        _ = v14
        _ = v15
    }
    v11 -= v11
    var v16 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
