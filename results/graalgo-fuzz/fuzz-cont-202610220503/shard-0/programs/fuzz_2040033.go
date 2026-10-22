package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }

func main() {
    var v1 int64 = 42
    var v2 bool = ((984 < 128) && false)
    var v3 bool = (18446744073709551615 != 42)
    v4 := GenS1{F0: 377, F1: 42}
    fmt.Println(v2, v2, v1)
    if ((true || (v2 && v2)) || (!false)) {
        var v5 uint64 = 4294967295
        v5 = (^(v5))
        var v6 float32 = 0.1
        _ = v5
        _ = v6
    } else {
        switch "\n\t" {
        case "":
            var v7 uint64 = 34
            _ = v7
        }
        var v8 uint = 2
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 uint32 = 65535
            var v11 int64 = (^(v1))
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    }
    v1 = (^(v1))
    fmt.Println(v3)
    for fi12 := 0; fi12 < 3; fi12++ {
        switch "" {
        case "foo bar":
            var v13 int = (^(fi12))
            _ = v13
        case "hello":
            v14 := []int{(fi12 ^ fi12), fi12, (^(fi12)), (fi12 % (((fi12 / ((fi12) | 1))) | 1))}
            _ = v14
        case "":
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 int8 = -1
                v17 := []int{-2147483648, 4294967295, 1, -1}
                var v18 int8 = (^(v16))
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        _ = fi12
    }
    v3 = v2
    v19 := []int{-596, 42, 9223372036854775807, -1}
    v3 = true
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 int8 = 127
        v22 := GenS1{F0: 2, F1: v21}
        var v23 int = fi20
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
