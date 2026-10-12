package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    v1 := []uint32{42}
    var v2 uint16 = 65535
    var v3 float32 = float32(float64(-2.25))
    var v4 bool = false
    var v5 float64 = 0.0
    v6 := make(map[uint32]uint32)
    v4 = (765 <= 347)
    fmt.Println(v3, v5, v5)
    var v7 uint64 = uint64(v2)
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 6; fi9++ {
            v3 = (v3 + -2.25)
            _ = fi9
        }
        switch "a" {
        case "foo bar":
            for fi10 := 0; fi10 < 19; fi10++ {
                if v4 {
                    var v11 int8 = 1
                    fmt.Println(v4)
                    fmt.Println(v2, v7)
                    delete(v6, 247)
                    _ = v11
                }
                v12 := GenS1{F0: 18446744073709551615, F1: "a"}
                _ = fi10
                _ = v12
            }
        default:
            var v13 int16 = 1
            _ = v13
        }
        var v14 uint = 42
        _ = fi8
        _ = v14
    }
    v7 = (v7 + v7)
    if v4 {
        v15 := GenS1{F0: 65535, F1: ""}
        delete(v6, 42)
        fmt.Println(v7)
        _ = v15
    } else {
        if (v4 && (((0 == -598) && v4) || (v4 || (!(-672 > 0))))) {
            for fi16 := 0; fi16 < 2; fi16++ {
                v5 = 1e10
                v7 -= (v7 - (v7 / ((v7) | 1)))
                if v4 {
                    v17 := GenS2{F0: (fi16 / ((-1) | 1)), F1: 652}
                    _ = v17
                }
                _ = fi16
            }
            fmt.Println(v3, v7)
        }
        v4 = (255 != v2)
        _ = len(v1)
    }
    var v18 uint32 = 600
    v7 = v7
    if v4 {
        var v19 uint = 42
        fmt.Println(v2, v4)
        _ = v19
    } else {
        if len(v1) > 0 {
            v1[0] = (v18 | (v18 ^ uint32(v2)))
        }
        v5 = 0.0
        var v20 bool = true
        _ = v20
    }
    var v21 int = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
