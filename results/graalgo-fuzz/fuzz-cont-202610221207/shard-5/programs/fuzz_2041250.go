package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[uint32]bool)
    v3 := []int{-397, 1, -9223372036854775808}
    var v4 uint16 = 42
    var v5 uint32 = 0
    for fi6 := 0; fi6 < 10; fi6++ {
        _ = len(v2)
        _ = fi6
    }
    delete(v2, v5)
    v5 = uint32(v4)
    v1 = (v1 % ((v1) | 1))
    v7 := GenS1{F0: -1, F1: uint32(v1)}
    if (!true) {
        var v8 uint8 = (v1 & (v1 + (v1 << ((uint(v5)) & 63))))
        v4 = (v4 * v4)
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: 100, F1: v5}
        _ = fi9
        _ = v10
    }
    v4 = 0
    if false {
        for fi11 := 0; fi11 < 15; fi11++ {
            v5 += (v5 / (((v5 & v5)) | 1))
            _ = fi11
        }
    }
    v12 := GenS1{F0: 127, F1: v5}
    v5 -= (v5 << ((2) & 63))
    if (-9223372036854775808 >= -699) {
        switch 599 {
        case -9223372036854775808:
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS1{F0: 23, F1: 2147483647}
                _ = fi13
                _ = v14
            }
        case 47:
            v1 = v1
        default:
            var v15 uint32 = v5
            _ = v15
        }
        var v16 string = ("hello" + "a")
        v4 = v4
        _ = v16
    } else {
        v17 := v2[(v5 * v5)]
        var v18 int16 = 0
        switch ("hello" + ("\n\t" + ("a" + "hello"))) {
        case "hello":
            v19 := GenS1{F0: 127, F1: v5}
            _ = v19
        case "a":
            v4 -= v4
        case "\n\t":
            if v17 {
                var v20 uint8 = v1
                v18 += v18
                for fi21 := 0; fi21 < 19; fi21++ {
                    fmt.Println(v18, fi21)
                    _ = fi21
                }
                _ = v20
            } else {
                if len(v3) > 0 {
                    v3[0] = -858
                }
                var v22 int8 = 42
                _ = v22
            }
        default:
            v17 = v17
        }
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
