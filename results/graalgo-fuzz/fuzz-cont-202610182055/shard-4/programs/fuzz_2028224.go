package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint32 }

func main() {
    var v1 int16 = 42
    var v2 float32 = 0.0
    v3 := []uint32{2147483647, 42, 692}
    if false {
        var v4 int8 = 127
        var v5 float32 = v2
        v6 := GenS1{F0: 65535, F1: v4, F2: 2147483648}
        _ = v4
        _ = v5
        _ = v6
    } else {
        _ = len(v3)
        var v7 float64 = 1.5
        _ = v7
    }
    var v8 uint32 = 4294967295
    v9 := GenS1{F0: (v8 - v8), F1: 71, F2: v8}
    fmt.Println(v1)
    switch "foo bar" {
    case "":
        switch 32768 {
        case -145:
            if len(v3) > 0 {
                v3[0] = (^(v8))
            }
        case -2147483648:
            var v10 float64 = 1e10
            _ = v10
        default:
            if len(v3) > 0 {
                v3[0] = 42
            }
        }
    case "\n\t":
        v1 = (v1 | v1)
    }
    fmt.Println(v2, v1)
    if ((0 != 42) && (!(-84 <= 42))) {
        v2 -= v2
        fmt.Println(v2)
    } else {
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 uint8 = 232
                switch ("a" + ("a" + (("a" + "") + "\n\t"))) {
                case "a":
                    switch fi12 {
                    case 1:
                        var v14 int = fi11
                        _ = v14
                    case 32767:
                        var v15 float32 = 3.14159
                        _ = v15
                    case 2147483647:
                        v2 = (v2 / v2)
                    }
                case "hello":
                    v1 += -516
                default:
                    var v16 uint16 = 687
                    _ = v16
                }
                var v17 string = ""
                _ = fi12
                _ = v13
                _ = v17
            }
            _ = len(v3)
            v18 := GenS1{F0: 2147483648, F1: int8(fi11), F2: (v8 + (v8 / (((v8 - v8)) | 1)))}
            _ = fi11
            _ = v18
        }
        var v19 uint8 = 42
        _ = v19
    }
    var v20 uint32 = v8
    var v21 int = 2147483647
    v22 := GenS2{F0: -902, F1: 292}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v8
    _ = v9
}
