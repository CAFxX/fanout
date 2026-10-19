package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int8 }
type GenS2 struct { F0 uint32; F1 int64; F2 int16 }
func genFn0() bool {
    if false {
        var v7 uint8 = 255
        _ = v7
    } else {
        var v8 uint16 = 42
        _ = v8
    }
    switch "hello" {
    case "":
        var v9 int64 = 0
        _ = v9
    default:
        v10 := []int{-1, 176}
        _ = v10
    }
    v11 := []int{-1, 32768, 42}
    v12 := GenS1{F0: 18446744073709551615, F1: 128, F2: 100}
    _ = v11
    _ = v12
    return false
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{308, 1}
    var v3 bool = (42 > -661)
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    var v6 int64 = 9223372036854775807
    if v3 {
        if true {
            switch "\n\t" {
            case "":
                v3 = v3
            default:
                v13 := GenS1{F0: 42, F1: 1, F2: -1}
                _ = v13
            }
            v4["\n\t"] = -2147483648
            var v14 uint = 255
            _ = v14
        }
        for fi15 := 0; fi15 < 14; fi15++ {
            switch "hello" {
            case "hello":
                if len(v2) > 0 {
                    v2[0] = 2
                }
            }
            fmt.Println(fi15)
            var v16 uint16 = 0
            _ = fi15
            _ = v16
        }
        switch -259 {
        case 32767:
            v3 = (true || (!v3))
        case 9223372036854775807:
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS1{F0: 369, F1: (fi17 % ((fi17) | 1)), F2: 100}
                v6 += (-(v6))
                _ = fi17
                _ = v18
            }
        case 1:
            switch 633 {
            case -9223372036854775808:
                v2 = append(v2, 374)
            default:
                fmt.Println(v3)
            }
        default:
            var v19 int32 = -1
            _ = v19
        }
    } else {
        v6 = (v6 & (v6 << ((65535) & 63)))
    }
    v20 := GenS1{F0: 2, F1: 127, F2: -15}
    fmt.Println(v6)
    v3 = true
    var v21 int = -1
    var v22 uint32 = 4294967295
    v23 := GenS2{F0: 2147483647, F1: -718, F2: -1}
    fmt.Println(v22)
    for fi24 := 0; fi24 < 9; fi24++ {
        v6 = 9223372036854775807
        var v25 uint16 = 42
        v1[(fi24 % ((0) | 1))] = v21
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
