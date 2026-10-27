package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 int; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 1
    v3 := []string{("\n\t" + "foo bar"), ("a" + ""), "\n\t", (("\n\t" + ("foo bar" + ("foo bar" + "\n\t"))) + ("\n\t" + "a"))}
    if (!false) {
        var v4 int64 = v2
        var v5 uint64 = 307
        _ = v4
        _ = v5
    }
    fmt.Println(v2)
    var v6 float32 = -2.25
    for fi7 := 0; fi7 < 8; fi7++ {
        _ = len(v3)
        v8 := GenS1{F0: -233, F1: fi7}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v2, v2)
        _ = fi9
    }
    if false {
        v6 += (v6 + v6)
        if len(v3) > 0 {
            v3[0] = ("a" + "\n\t")
        }
        var v10 uint16 = 627
        _ = v10
    }
    v2 -= (v2 | v2)
    if false {
        var v11 int8 = 61
        var v12 uint = 42
        for fi13 := 0; fi13 < 14; fi13++ {
            if true {
                v2 -= (v2 + (-(v2)))
                var v14 uint = 4294967295
                var v15 int = fi13
                _ = v14
                _ = v15
            }
            _ = fi13
        }
        _ = v11
        _ = v12
    }
    fmt.Println(v6, v2)
    v16 := GenS2{F0: 32768, F1: false}
    v2 += v2
    v6 = v6
    var v17 uint = 331
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
}
