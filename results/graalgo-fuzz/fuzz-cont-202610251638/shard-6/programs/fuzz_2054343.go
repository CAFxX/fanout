package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint64 }

func main() {
    var v1 float32 = (0.1 * 1.5)
    v2 := []int{42}
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 1; fi6++ {
        switch "a" {
        case "foo bar":
            v1 = v1
        case "":
            switch "foo bar" {
            case "":
                v7 := GenS1{F0: 788}
                _ = v7
            case "hello":
                fmt.Println(v1, v1)
            case "\n\t":
                v8 := v3[""]
                _ = v8
            default:
                v9 := GenS1{F0: -851}
                _ = v9
            }
        default:
            var v10 int64 = 146
            _ = v10
        }
        _ = fi6
    }
    v1 += v1
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := GenS2{F0: 2}
        v1 -= float32(((0.0 * 3.14159) - (1.5 / 1.5)))
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        v14 := GenS3{F0: 683, F1: 2, F2: 42}
        _ = fi13
        _ = v14
    }
    v1 = (v1 * v1)
    switch -185 {
    case 807:
        fmt.Println(v1)
    case 42:
        fmt.Println(v1)
    case -9223372036854775808:
        fmt.Println(v1)
    }
    v1 = 0.0
    switch "hello" {
    case "":
        var v15 uint64 = 349
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 14; fi16++ {
            switch (fi16 / ((fi16) | 1)) {
            case 4294967295:
                _ = len(v2)
            }
            v17 := GenS2{F0: 287}
            var v18 uint64 = 4294967295
            _ = fi16
            _ = v17
            _ = v18
        }
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = 32768
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
