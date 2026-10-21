package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 int8; F2 int }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v4 int64 = 173
    v5 := make(map[uint32]bool)
    var v6 int8 = -128
    v7 := []int{128, 42, 264, 42}
    switch 42 {
    case -967:
        switch v1 {
        case "foo bar":
            v4 -= v4
        case "\n\t":
            v1 = "hello"
        case "":
            var v8 uint16 = 135
            _ = v8
        default:
            var v9 int64 = v4
            _ = v9
        }
    case -181:
        for fi10 := 0; fi10 < 16; fi10++ {
            v6 = v6
            if true {
                v11 := GenS1{F0: v6, F1: v4, F2: (v6 % ((int8(v4)) | 1))}
                var v12 uint = 18446744073709551615
                _ = v11
                _ = v12
            } else {
                delete(v2, 548)
            }
            _ = fi10
        }
    default:
        v4 -= (-(v4))
    }
    var v13 float64 = 0.1
    fmt.Println(v6, v1)
    v13 = 1.5
    v14 := GenS1{F0: (v6 + v6), F1: v4, F2: v6}
    fmt.Println(v13)
    v6 += -15
    for fi15 := 0; fi15 < 4; fi15++ {
        if ((fi15 | fi15) > fi15) {
            var v16 string = "foo bar"
            _ = v16
        } else {
            for fi17 := 0; fi17 < 4; fi17++ {
                var v18 int32 = -2147483648
                var v19 bool = true
                v20 := GenS1{F0: -128, F1: v4, F2: v6}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        }
        v21 := GenS2{F0: 421}
        _ = fi15
        _ = v21
    }
    v22 := GenS3{F0: true, F1: -128, F2: -9223372036854775808}
    switch -516 {
    case 322:
        if len(v7) > 0 {
            v7[0] = -444
        }
    default:
        fmt.Println(v6, v13)
    }
    var v23 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
