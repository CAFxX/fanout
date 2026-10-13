package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = 1e10
    var v3 int8 = 100
    var v4 int16 = -992
    if (42 != 535) {
        v2 = float64(0.0)
        var v5 uint = 127
        for fi6 := 0; fi6 < 5; fi6++ {
            v4 -= (v4 | (v4 & (v4 + (v4 | v4))))
            fmt.Println(v4, v4)
            _ = fi6
        }
        _ = v5
    }
    if true {
        v3 = (v3 + v3)
    }
    v7 := []int{42, 127, 32767}
    v2 -= 1.5
    fmt.Println(v2)
    var v8 int64 = 0
    if true {
        fmt.Println(v2, v2, v8)
        var v9 int16 = v4
        v2 = v2
        _ = v9
    } else {
        var v10 int = 9223372036854775807
        v3 = v3
        _ = v10
    }
    if false {
        v4 = (-(v4))
        var v11 uint64 = 4294967295
        _ = v11
    }
    v8 = v8
    v12 := GenS1{F0: 194, F1: 0}
    v1[("hello" + ("a" + "foo bar"))] = (62 != 4294967295)
    fmt.Println(v2, v2, v4)
    v2 += 0.0
    for fi13 := 0; fi13 < 4; fi13++ {
        if false {
            v8 -= (v8 << ((0) & 63))
            var v14 uint = 2
            if false {
                v14 -= 65535
                var v15 int64 = v8
                _ = v15
            }
            _ = v14
        }
        var v16 uint32 = 4294967295
        v4 -= (^(v4))
        _ = fi13
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
