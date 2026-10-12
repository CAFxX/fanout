package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 int) int16 {
    var v4 string = "\n\t"
    p0 = (p0 - p0)
    switch v4 {
    case "a":
        switch "" {
        case "foo bar":
            switch v4 {
            case "\n\t":
                p0 += (p0 >> ((911) & 63))
            default:
                var v5 uint = 621
                _ = v5
            }
        default:
            v6 := GenS1{F0: 0, F1: 127}
            _ = v6
        }
    case "foo bar":
        v7 := []int{}
        _ = v7
    }
    if (!false) {
        var v8 int32 = 2147483647
        _ = v8
    } else {
        v4 = v4
        for fi9 := 0; fi9 < 19; fi9++ {
            if (1 >= 318) {
                p0 += (fi9 / ((9223372036854775807) | 1))
            } else {
                var v10 uint8 = 1
                fmt.Println(p0)
                v4 = v4
                _ = v10
            }
            fmt.Println(p0, v4, v4)
            _ = fi9
        }
    }
    _ = p0
    _ = v4
    return 1
}

func main() {
    var v1 int32 = 0
    var v2 float32 = ((-2.25 / 0.1) / 1.0)
    var v3 int64 = -1
    var v11 int64 = v3
    v12 := []int{}
    if ((!(65535 <= 1)) && (true || (!false))) {
        v13 := GenS1{F0: -91, F1: 1}
        _ = v13
    }
    fmt.Println(v3)
    v3 -= v11
    var v14 uint = 429
    fmt.Println(v2, v14, v2)
    if false {
        switch "a" {
        case "":
            v14 += (v14 + v14)
        }
    }
    var v15 uint8 = 42
    v11 -= 1
    if len(v12) > 0 {
        v12[0] = int(v14)
    }
    var v16 int = -215
    v3 = v3
    switch "foo bar" {
    case "a":
        v17 := GenS1{F0: int8(v1), F1: 42}
        _ = v17
    }
    switch v16 {
    case 4294967295:
        _ = len(v12)
    case 128:
        v11 -= -9223372036854775808
    case -9223372036854775808:
        v14 -= v14
    default:
        if (4294967295 < 569) {
            fmt.Println(v2, v15)
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        }
    }
    v15 -= v15
    var v18 uint16 = 922
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
