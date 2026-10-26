package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint8; F2 uint }

func main() {
    v1 := []int{32768, 32767, 9223372036854775807, 42}
    var v2 float64 = -2.25
    var v3 int8 = 1
    var v4 uint64 = 644
    var v5 int16 = 134
    v6 := []uint32{}
    var v7 int = 2147483647
    v3 = v3
    var v8 int = (v7 ^ v7)
    if (244 < 2) {
        switch "\n\t" {
        case "a":
            _ = len(v6)
        case "hello":
            for fi9 := 0; fi9 < 16; fi9++ {
                if (42 <= 4294967295) {
                    var v10 uint64 = v4
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v10
                }
                for fi11 := 0; fi11 < 8; fi11++ {
                    v8 += v7
                    fmt.Println(v5)
                    _ = fi11
                }
                _ = fi9
            }
        }
    } else {
        fmt.Println(v8, v5, v4)
        fmt.Println(v4, v5, v3)
        if (14 < 2147483648) {
            var v12 int8 = v3
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        if len(v6) > 0 {
            v6[0] = 1
        }
        _ = fi13
    }
    switch v8 {
    case 813:
        switch v7 {
        case -2147483648:
            var v14 float32 = (0.0 - (((-2.25 / 0.1) - 1e10) * -2.25))
            _ = v14
        }
    case -736:
        v3 = -1
    }
    fmt.Println(v2)
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v1 = append(v1, v7)
        fmt.Println(v8, v8)
        var v17 int8 = -1
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 15
    v3 -= (v3 % (((v3 | v3)) | 1))
    v8 -= v8
    if true {
        var v19 uint32 = uint32(v4)
        fmt.Println(v18, v3, v5)
        v20 := GenS1{F0: false}
        _ = v19
        _ = v20
    }
    var v21 float32 = 0.0
    v22 := GenS2{F0: v7, F1: 1, F2: 4294967295}
    var v23 uint16 = 765
    v18 -= (v18 >> ((255) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
