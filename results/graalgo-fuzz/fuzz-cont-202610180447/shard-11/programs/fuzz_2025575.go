package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -9223372036854775808
    var v3 int = 998
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    v6 := []int{(^(v3)), 9223372036854775807, v3}
    v3 += (-(v3))
    delete(v5, 2147483647)
    switch "" {
    case "foo bar":
        var v7 uint64 = 404
        _ = v7
    }
    fmt.Println(v3)
    for fi8 := 0; fi8 < 20; fi8++ {
        if (2 >= 2) {
            fmt.Println(fi8, fi8, v3)
            if len(v6) > 0 {
                v6[0] = 331
            }
            for fi9 := 0; fi9 < 2; fi9++ {
                v2 -= -408
                v3 = (fi8 | 4294967295)
                v3 += (v3 % (((v3 >> ((0) & 63))) | 1))
                _ = fi9
            }
        } else {
            fmt.Println(v2)
            var v10 uint16 = 179
            _ = v10
        }
        var v11 uint = 42
        var v12 int16 = 1
        _ = fi8
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS1{F0: (fi13 ^ (v3 ^ fi13))}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        if true {
            v2 = -993
        }
        _ = fi15
    }
    if len(v6) > 0 {
        v6[0] = v3
    }
    v3 = v3
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        fmt.Println(v2)
        _ = fi16
    }
    if len(v6) > 0 {
        v6[0] = v3
    }
    v3 += v3
    for fi17 := 0; fi17 < 3; fi17++ {
        v2 = (v2 & (v2 % ((-864) | 1)))
        var v18 uint = 18446744073709551615
        var v19 float32 = 0.0
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
