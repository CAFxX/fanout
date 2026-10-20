package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 string }

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]uint32)
    var v3 int32 = 974
    v1 -= (v1 ^ 1)
    v4 := GenS1{F0: 593}
    v5 := []int{4294967295, 2147483647, 42, 4294967295}
    var v6 uint32 = 65535
    if false {
        v7 := GenS2{F0: 42, F1: "a"}
        var v8 bool = true
        fmt.Println(v8, v6)
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        switch fi9 {
        case 128:
            for fi10 := 0; fi10 < 16; fi10++ {
                v6 = v6
                v6 = 650
                v3 = v3
                _ = fi10
            }
        case 127:
            var v11 uint8 = 162
            _ = v11
        case 962:
            delete(v2, "\n\t")
        default:
            var v12 int16 = 53
            _ = v12
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint16 = v1
            _ = fi13
            _ = v14
        }
        _ = fi9
    }
    v6 += (v6 ^ (v6 & (v6 / ((v6) | 1))))
    for fi15 := 0; fi15 < 18; fi15++ {
        switch "hello" {
        case "a":
            v3 = (^(v3))
        case "foo bar":
            delete(v2, "foo bar")
        case "\n\t":
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 float64 = 1.5
                _ = fi16
                _ = v17
            }
        }
        fmt.Println(v3, v6, fi15)
        _ = fi15
    }
    v3 -= v3
    for fi18 := 0; fi18 < 4; fi18++ {
        if false {
            var v19 float32 = (float32(3.14159) - 3.14159)
            _ = v19
        }
        _ = len(v5)
        fmt.Println(v1, v1)
        _ = fi18
    }
    var v20 uint16 = (v1 | v1)
    if false {
        v20 -= v1
        var v21 int64 = 9223372036854775807
        var v22 int8 = -27
        _ = v21
        _ = v22
    }
    if false {
        v3 += (v3 * (v3 ^ 0))
        v3 += v3
    }
    var v23 int16 = 898
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
