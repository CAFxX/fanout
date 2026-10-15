package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int8; F2 int16 }
type GenS4 struct { F0 uint16; F1 int64 }

func main() {
    var v1 uint8 = 236
    var v2 float64 = 3.14159
    var v3 int32 = -1
    var v4 float64 = (0.0 / -2.25)
    var v5 uint16 = 42
    var v6 uint64 = 42
    v6 = v6
    switch 127 {
    case 0:
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 int8 = 1
                v4 += float64(0.1)
                v10 := GenS1{F0: 177, F1: (("foo bar" + "foo bar") > ("\n\t" + "a"))}
                _ = fi8
                _ = v10
                _ = v9
            }
            switch fi7 {
            case 127:
                var v11 int32 = v3
                _ = v11
            case 0:
                var v12 float64 = (v4 / v2)
                _ = v12
            default:
                v5 = (^(v5))
            }
            v13 := GenS1{F0: (v5 % ((uint16(v6)) | 1)), F1: false}
            _ = fi7
            _ = v13
        }
    }
    var v14 int = -9223372036854775808
    switch "foo bar" {
    case "foo bar":
        var v15 int8 = 127
        _ = v15
    case "hello":
        v16 := GenS2{F0: 1}
        _ = v16
    case "a":
        fmt.Println(v3)
    default:
        var v17 uint32 = 2147483647
        _ = v17
    }
    var v18 float64 = v4
    var v19 float64 = float64((3.14159 - (float32(((v4 / float64(-2.25)) / float64((0.1 * (1.5 / 1e10))))) - 0.0)))
    var v20 float64 = 3.14159
    fmt.Println(v6, v19, v5)
    var v21 int64 = 42
    var v22 int = 0
    v23 := GenS3{F0: (^(v1)), F1: 104, F2: 1}
    v24 := GenS4{F0: v5, F1: v21}
    v25 := []int{(-(v22))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
