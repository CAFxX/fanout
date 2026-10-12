package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := []string{"foo bar", "foo bar"}
    var v2 int64 = 1
    var v3 int16 = -539
    v4 := make(map[int]int)
    v5 := []string{"", "", "hello"}
    var v6 uint32 = 42
    var v7 uint8 = 172
    v8 := GenS1{F0: 467}
    _ = len(v4)
    switch 519 {
    case -28:
        for fi9 := 0; fi9 < 7; fi9++ {
            v6 += 42
            _ = fi9
        }
    default:
        var v10 int8 = 100
        _ = v10
    }
    if (!(0 < 2)) {
        v11 := v4[0]
        _ = v11
    } else {
        var v12 bool = (1e10 < 1.5)
        fmt.Println(v6)
        if (!(1 == 375)) {
            v13 := v4[-1]
            v3 -= v3
            _ = v13
        } else {
            v14 := GenS1{F0: uint(v6)}
            _ = v14
        }
        _ = v12
    }
    fmt.Println(v6, v7, v7)
    for fi15 := 0; fi15 < 18; fi15++ {
        if (-128 == 42) {
            if true {
                v2 = (v2 / ((v2) | 1))
                var v16 int32 = 1
                _ = v16
            } else {
                var v17 int = fi15
                _ = v17
            }
        }
        _ = fi15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
