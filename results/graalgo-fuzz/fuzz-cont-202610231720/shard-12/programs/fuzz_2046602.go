package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int16; F1 int32; F2 int64 }
type GenS3 struct { F0 int64 }
func genFn0() int16 {
    var v4 int8 = 69
    var v5 int32 = -1
    v4 = (v4 ^ v4)
    _ = v4
    _ = v5
    return -1
}

func main() {
    var v1 uint16 = 0
    var v2 uint16 = 32767
    v3 := []int{}
    v6 := GenS1{F0: 100, F1: 42}
    switch "hello" {
    case "":
        switch ("hello" + "foo bar") {
        case "":
            _ = len(v3)
        case "\n\t":
            v1 += v1
        default:
            v7 := GenS1{F0: 100, F1: 4294967295}
            _ = v7
        }
    case "foo bar":
        var v8 float32 = (1e10 - 0.1)
        _ = v8
    case "a":
        if true {
            for fi9 := 0; fi9 < 7; fi9++ {
                if true {
                    v10 := GenS1{F0: -128, F1: 4294967295}
                    _ = v10
                }
                if true {
                    v11 := GenS2{F0: -330, F1: 0, F2: -777}
                    fmt.Println(v2, v2, fi9)
                    _ = v11
                }
                var v12 int16 = 1
                _ = fi9
                _ = v12
            }
            v1 -= 136
        }
    }
    fmt.Println(v1, v2)
    fmt.Println(v2, v1)
    if len(v3) > 0 {
        v3[0] = 715
    }
    v3 = append(v3, -674)
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint8 = 42
        _ = fi13
        _ = v14
    }
    switch -334 {
    case -375:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case -102:
        fmt.Println(v2, v2)
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println(v2, v1)
    var v15 int16 = -1
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := GenS2{F0: v15, F1: -2147483648, F2: int64(fi16)}
        v18 := GenS2{F0: v15, F1: 2147483647, F2: 0}
        var v19 int8 = 122
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := GenS3{F0: 272}
    switch -341 {
    case -65:
        for fi21 := 0; fi21 < 13; fi21++ {
            fmt.Println(v1, fi21, fi21)
            _ = fi21
        }
    case -9223372036854775808:
        var v22 int = 623
        _ = v22
    }
    fmt.Println(v2)
    var v23 uint16 = v2
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v6
}
