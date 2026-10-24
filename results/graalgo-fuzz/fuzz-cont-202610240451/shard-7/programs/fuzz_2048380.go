package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := []int{52}
    v3 := []int{}
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    var v6 uint16 = 1
    var v7 int8 = -128
    var v8 int64 = 42
    if true {
        for fi9 := 0; fi9 < 12; fi9++ {
            var v10 int32 = 1
            _ = fi9
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            v12 := GenS1{F0: (v7 << ((4294967295) & 63)), F1: 811}
            _ = fi11
            _ = v12
        }
        var v13 uint32 = 2
        _ = v13
    }
    fmt.Println(v8, v8)
    switch "" {
    case "hello":
        _ = len(v1)
    case "\n\t":
        v8 -= v8
    case "a":
        var v14 int16 = 1000
        _ = v14
    }
    switch 2147483647 {
    case 127:
        v15 := v1[32767]
        _ = v15
    case -549:
        v6 = (v6 % ((v6) | 1))
    case 9223372036854775807:
        var v16 uint = 0
        _ = v16
    default:
        if true {
            var v17 uint8 = 42
            _ = v17
        } else {
            for fi18 := 0; fi18 < 8; fi18++ {
                v6 -= v6
                _ = fi18
            }
            v8 = v8
        }
    }
    v8 = (v8 / ((v8) | 1))
    v19 := GenS1{F0: (v7 - v7), F1: 795}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
