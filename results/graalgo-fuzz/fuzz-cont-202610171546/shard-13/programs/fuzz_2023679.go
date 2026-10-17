package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }

func main() {
    var v1 uint = 65535
    v2 := []int{32768, -9223372036854775808, -877}
    v3 := make(map[int]int)
    if false {
        if false {
            var v4 int = 996
            var v5 bool = (-631 != -1)
            fmt.Println(v5, v5)
            _ = v4
            _ = v5
        } else {
            v1 += (v1 >> ((v1) & 63))
        }
        if (!(!(true && (2147483647 > 65535)))) {
            if false {
                v1 += v1
                for fi6 := 0; fi6 < 15; fi6++ {
                    v1 += 42
                    v7 := GenS1{F0: 65535, F1: 2147483647}
                    _ = fi6
                    _ = v7
                }
            }
            fmt.Println(v1)
        }
    } else {
        _ = len(v2)
    }
    if (-835 > 32768) {
        v1 += (v1 * v1)
        var v8 string = ("\n\t" + "hello")
        v8 = v8
        _ = v8
    }
    v1 += v1
    v1 += (v1 / ((117) | 1))
    var v9 uint16 = 65535
    v10 := GenS1{F0: 983, F1: 346}
    var v11 int = 378
    var v12 int32 = -1
    for fi13 := 0; fi13 < 13; fi13++ {
        if (!false) {
            if (!false) {
                v14 := GenS2{F0: 1, F1: 56, F2: -76}
                var v15 uint = v1
                _ = v14
                _ = v15
            }
            var v16 uint32 = 2147483647
            _ = v16
        }
        _ = fi13
    }
    v9 += v9
    switch v11 {
    case 591:
        switch "a" {
        case "foo bar":
            fmt.Println(v12, v1, v12)
        case "a":
            if len(v2) > 0 {
                v2[0] = (v11 & (v11 % ((-2147483648) | 1)))
            }
        }
    }
    v17 := GenS1{F0: 308, F1: 4294967295}
    var v18 int64 = 0
    if false {
        if len(v2) > 0 {
            v2[0] = (v11 / (((v11 % ((v11) | 1))) | 1))
        }
        v9 += (v9 - (v9 ^ v9))
    } else {
        var v19 float32 = float32(1.5)
        var v20 uint8 = 221
        _ = v19
        _ = v20
    }
    if (0.0 != 3.14159) {
        v21 := GenS1{F0: 296, F1: 65535}
        _ = v21
    } else {
        v2 = append(v2, (v11 >> ((496) & 63)))
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := GenS3{F0: (v12 ^ 2147483647), F1: (v9 / (((v9 & (v9 << ((v1) & 63)))) | 1)), F2: v1}
        fmt.Println(fi22, v9)
        v9 += v9
        _ = fi22
        _ = v23
    }
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
