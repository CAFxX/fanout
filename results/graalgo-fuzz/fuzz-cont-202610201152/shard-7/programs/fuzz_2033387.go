package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "hello"
    v3 := []int{32768, 128}
    _ = len(v3)
    v4 := GenS1{F0: 2147483647, F1: ("" + v2)}
    var v5 uint = 255
    var v6 uint64 = 766
    fmt.Println(v6, v5)
    v2 = "a"
    v7 := GenS1{F0: -2147483648, F1: v2}
    for fi8 := 0; fi8 < 8; fi8++ {
        v1[2147483648] = 37
        v9 := GenS2{F0: -1, F1: (fi8 - (fi8 * fi8))}
        _ = fi8
        _ = v9
    }
    if false {
        var v10 uint8 = 93
        v6 -= (v6 + (v6 | 473))
        _ = v10
    } else {
        var v11 int64 = -9223372036854775808
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        if (128 == 128) {
            switch fi12 {
            case 128:
                v2 = v2
            case 1:
                v5 -= (v5 % ((255) | 1))
            case 0:
                if (!(-111 >= -1)) {
                    v2 = (v2 + v2)
                    var v13 int32 = 0
                    _ = v13
                } else {
                    v2 = v2
                }
            }
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := GenS3{F0: 65535}
                _ = fi14
                _ = v15
            }
            v2 = "a"
        }
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 int64 = 1
            if false {
                v5 = uint(v6)
                var v18 int8 = 127
                var v19 bool = (!false)
                _ = v18
                _ = v19
            } else {
                v17 += (v17 >> ((901) & 63))
                v20 := v1[436]
                v21 := GenS4{F0: ("hello" + v2), F1: (^(v6))}
                v6 += 543
                if len(v3) > 0 {
                    v3[0] = fi16
                }
                var v22 float32 = 3.14159
                delete(v1, uint32(v5))
                _ = v20
                _ = v21
                _ = v22
            }
            v23 := v1[729]
            var v24 string = v2
            _ = fi16
            _ = v17
            _ = v23
            _ = v24
        }
        v25 := v1[0]
        _ = fi12
        _ = v25
    }
    fmt.Println(v2, v2)
    v2 = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
