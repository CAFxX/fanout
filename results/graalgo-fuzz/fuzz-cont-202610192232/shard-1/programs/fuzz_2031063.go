package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int; F2 uint8 }

func main() {
    v1 := []int{}
    var v2 uint32 = 42
    v3 := make(map[uint32]uint32)
    var v4 float64 = 0.1
    switch 0 {
    case 42:
        v2 = v2
    case 32767:
        fmt.Println(v2)
    case 833:
        var v5 int = 128
        _ = v5
    }
    switch "foo bar" {
    case "\n\t":
        for fi6 := 0; fi6 < 13; fi6++ {
            if false {
                var v7 int = (fi6 % (((fi6 / (((fi6 % ((902) | 1))) | 1))) | 1))
                _ = v7
            }
            _ = fi6
        }
    default:
        if true {
            v8 := GenS1{F0: (v2 >> ((2) & 63))}
            for fi9 := 0; fi9 < 13; fi9++ {
                _ = len(v3)
                fmt.Println(v4, v4)
                _ = fi9
            }
            _ = v8
        } else {
            v4 -= float64(float32(0.0))
        }
    }
    v1 = append(v1, 1)
    var v10 bool = false
    var v11 float32 = 0.0
    var v12 float32 = (float32((float64(((((0.0 * v11) + ((v11 * 0.0) - (((3.14159 + float32((float64(v11) * 0.1))) / v11) - -2.25))) * v11) * (-2.25 / (1e10 / 0.1)))) - v4)) * v11)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    for fi13 := 0; fi13 < 8; fi13++ {
        v10 = v10
        _ = fi13
    }
    if ((!("" != ("hello" + ("a" + "\n\t")))) && false) {
        v12 = v12
        v14 := GenS2{F0: 127, F1: 0, F2: 65535}
        if v10 {
            var v15 float64 = -2.25
            _ = len(v3)
            switch 749 {
            case -2147483648:
                fmt.Println(v2)
            case -1:
                v16 := GenS3{F0: 230, F1: 32768, F2: 0}
                _ = v16
            }
            _ = v15
        }
        _ = v14
    }
    v3[2147483647] = (v2 << ((708) & 63))
    if (-2147483648 > -1) {
        if (0 != 594) {
            for fi17 := 0; fi17 < 12; fi17++ {
                v18 := GenS2{F0: 1, F1: 32767, F2: 255}
                var v19 int16 = 1000
                var v20 float32 = 0.0
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        }
    }
    var v21 uint32 = (v2 | 2147483647)
    if len(v1) > 0 {
        v1[0] = -127
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
