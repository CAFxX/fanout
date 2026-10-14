package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }

func main() {
    var v1 uint = 1
    v2 := make(map[int]bool)
    v3 := []int{-99}
    var v4 uint16 = 0
    var v5 uint32 = 2
    v6 := make(map[int]bool)
    if (!false) {
        v1 -= 65535
        var v7 uint32 = v5
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        switch fi8 {
        case -9223372036854775808:
            var v9 float64 = 0.1
            _ = v9
        default:
            var v10 string = ""
            _ = v10
        }
        v4 = 762
        var v11 float64 = 3.14159
        _ = fi8
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v3 = append(v3, fi12)
        switch ("" + ("hello" + "hello")) {
        case "a":
            fmt.Println(v5, v1)
        }
        if len(v3) > 0 {
            v3[0] = fi12
        }
        _ = fi12
    }
    if (0 == 127) {
        var v13 int64 = -1
        switch "a" {
        case "a":
            fmt.Println(v1)
        case "\n\t":
            var v14 string = ("a" + ("foo bar" + ""))
            _ = v14
        case "foo bar":
            v1 = (v1 | v1)
        default:
            for fi15 := 0; fi15 < 20; fi15++ {
                v3 = append(v3, fi15)
                v16 := GenS1{F0: 0, F1: ("hello" + "foo bar")}
                _ = fi15
                _ = v16
            }
        }
        _ = v13
    }
    switch "" {
    case "\n\t":
        _ = len(v3)
    default:
        v5 -= 2
    }
    switch 127 {
    case 9223372036854775807:
        v4 = 356
    case 32767:
        if len(v3) > 0 {
            v3[0] = 32768
        }
    default:
        v4 = (v4 / (((v4 << (((^(v1))) & 63))) | 1))
    }
    var v17 string = ((("hello" + "foo bar") + "foo bar") + "a")
    switch v17 {
    case "foo bar":
        v4 += v4
    case "hello":
        fmt.Println(v17)
    default:
        var v18 int32 = 352
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 string = v17
        fmt.Println(v17, v17, v5)
        v21 := GenS1{F0: 42, F1: v17}
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
