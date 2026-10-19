package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }

func main() {
    var v1 uint16 = 51
    v2 := make(map[string]uint32)
    var v3 uint64 = 18446744073709551615
    var v4 bool = ((0 > -2147483648) && false)
    v5 := GenS1{F0: -182, F1: "\n\t"}
    v6 := GenS1{F0: 893, F1: "\n\t"}
    var v7 int8 = -128
    if v4 {
        for fi8 := 0; fi8 < 16; fi8++ {
            v7 = (v7 << ((4294967295) & 63))
            if v4 {
                var v9 uint = 0
                if (1000 >= 22) {
                    v1 += v1
                    var v10 int32 = -2147483648
                    _ = v10
                }
                _ = v9
            }
            var v11 uint8 = 128
            _ = fi8
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 int64 = -641
            for fi14 := 0; fi14 < 1; fi14++ {
                v15 := GenS2{F0: 32768, F1: (v1 + v1)}
                _ = fi14
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        v16 := []int{-1}
        switch -712 {
        case 4294967295:
            var v17 uint = 2
            _ = v17
        case -1:
            v7 = 1
        case -482:
            v18 := GenS3{F0: v1, F1: 1}
            _ = v18
        default:
            if ((0 > 9223372036854775807) && (42 == 1000)) {
                if v4 {
                    var v19 bool = (255 <= 255)
                    _ = v19
                }
                v1 = 596
            }
        }
        _ = v16
    }
    v4 = (v1 != (v1 - v1))
    v20 := GenS2{F0: -1, F1: v1}
    var v21 int16 = 762
    var v22 uint16 = v1
    v4 = v4
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 int8 = -128
        var v25 uint8 = 1
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
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
}
