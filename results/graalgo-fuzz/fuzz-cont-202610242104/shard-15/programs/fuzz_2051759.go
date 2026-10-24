package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int; F2 uint64 }

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[uint32]uint32)
    var v3 int8 = 0
    v4 := make(map[uint32]bool)
    v5 := []int{-2147483648}
    v6 := []string{"hello"}
    fmt.Println(v3)
    v7 := GenS1{F0: 40}
    v8 := GenS2{F0: 4294967295, F1: 4294967295, F2: 0}
    for fi9 := 0; fi9 < 8; fi9++ {
        if (2 < 255) {
            v3 = v3
        }
        _ = fi9
    }
    if (1 == 2) {
        for fi10 := 0; fi10 < 8; fi10++ {
            v3 = v3
            v11 := v4[1]
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 int64 = 9223372036854775807
                _ = fi12
                _ = v13
            }
            _ = fi10
            _ = v11
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v14 int32 = 2147483647
        _ = v14
    } else {
        fmt.Println(v3)
        v5 = append(v5, 9223372036854775807)
    }
    switch "" {
    case "\n\t":
        var v15 int64 = 1
        _ = v15
    case "":
        v3 = (v3 >> ((4294967295) & 63))
    case "a":
        switch "a" {
        case "hello":
            v3 = (v3 ^ (v3 / ((v3) | 1)))
        default:
            v3 += v3
        }
    default:
        switch "hello" {
        case "a":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "":
            for fi16 := 0; fi16 < 7; fi16++ {
                for fi17 := 0; fi17 < 3; fi17++ {
                    fmt.Println(fi16)
                    var v18 float64 = 3.14159
                    var v19 uint = 390
                    v18 += v18
                    v19 = 136
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
        }
    }
    v3 += v3
    v20 := GenS1{F0: 255}
    for fi21 := 0; fi21 < 19; fi21++ {
        v3 = -1
        v2[71] = 348
        v22 := v4[929]
        _ = fi21
        _ = v22
    }
    fmt.Println(v3)
    v3 -= (v3 / ((v3) | 1))
    var v23 string = ""
    v24 := GenS2{F0: 4294967295, F1: 9223372036854775807, F2: 818}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
