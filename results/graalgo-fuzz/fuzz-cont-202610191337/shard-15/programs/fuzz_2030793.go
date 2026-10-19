package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    var v1 uint32 = 1
    var v2 uint = 429
    v3 := make(map[int]uint32)
    v4 := make(map[int]uint32)
    var v5 uint64 = 42
    if false {
        var v6 int32 = -1
        var v7 float64 = (3.14159 * -2.25)
        _ = v6
        _ = v7
    }
    var v8 int = 924
    switch (^(v8)) {
    case 32767:
        fmt.Println(v2, v8, v1)
    default:
        fmt.Println(v8)
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := []int{v8, fi10}
            _ = fi10
            _ = v11
        }
        var v12 int = (^(fi9))
        v5 = 1
        _ = fi9
        _ = v12
    }
    switch ((("\n\t" + ("hello" + "a")) + "") + "a") {
    case "":
        v1 = v1
    default:
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    if (0 == 42) {
        var v14 uint32 = 42
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 int8 = 127
            v17 := []int{}
            for fi18 := 0; fi18 < 4; fi18++ {
                fmt.Println(v1)
                var v19 uint64 = 42
                var v20 int32 = 1
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = v14
    }
    v21 := GenS1{F0: -1, F1: v8}
    var v22 uint64 = v5
    v4[(v8 / (((v8 | (v8 & (v8 ^ -823)))) | 1))] = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
