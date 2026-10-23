package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint32; F2 string }
type GenS3 struct { F0 bool; F1 uint32; F2 uint }
type GenS4 struct { F0 string; F1 uint; F2 int }

func main() {
    var v1 uint = 0
    var v2 string = (("a" + "\n\t") + "a")
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: v2, F1: 0}
    var v5 string = ""
    v6 := []int{-2147483648}
    var v7 uint8 = 127
    if (-1 >= -9223372036854775808) {
        switch v2 {
        case "foo bar":
            v7 -= (v7 | (v7 << (((^(v1))) & 63)))
        case "":
            fmt.Println(v5, v5, v1)
        default:
            v1 = v1
        }
    }
    var v8 bool = (!false)
    v8 = true
    switch 10 {
    case -2147483648:
        for fi9 := 0; fi9 < 19; fi9++ {
            v3[804] = v5
            for fi10 := 0; fi10 < 8; fi10++ {
                if (!(0 >= 2147483648)) {
                    v5 = (v2 + "a")
                    v11 := GenS1{F0: ((v5 + v2) + v2), F1: -2147483648}
                    v12 := v3[2147483648]
                    v7 = v7
                    if false {
                        v13 := GenS2{F0: v7, F1: 347, F2: v12}
                        var v14 uint8 = 127
                        var v15 int64 = -9223372036854775808
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    _ = v11
                    _ = v12
                }
                v16 := v3[1]
                _ = fi10
                _ = v16
            }
            v17 := GenS3{F0: ((2 >= 1) && v8), F1: 2147483647, F2: 345}
            _ = fi9
            _ = v17
        }
    case 9223372036854775807:
        if v8 {
            fmt.Println(v7)
        }
    case -1:
        v18 := GenS3{F0: v8, F1: uint32(v1), F2: 42}
        _ = v18
    default:
        var v19 int64 = 0
        _ = v19
    }
    var v20 int8 = 127
    v8 = v8
    var v21 bool = v8
    v20 = -1
    v22 := GenS4{F0: v5, F1: 0, F2: 42}
    if v8 {
        var v23 float64 = 1.5
        _ = v23
    }
    v8 = ((!(uint8(v1) <= v7)) || (82 == 927))
    for fi24 := 0; fi24 < 14; fi24++ {
        if (v21 || (42 < 42)) {
            v5 = "foo bar"
            _ = len(v3)
        }
        fmt.Println(v21)
        v3[2147483648] = v5
        _ = fi24
    }
    if false {
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
        for fi25 := 0; fi25 < 17; fi25++ {
            for fi26 := 0; fi26 < 16; fi26++ {
                v6 = append(v6, fi25)
                _ = fi26
            }
            _ = fi25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
