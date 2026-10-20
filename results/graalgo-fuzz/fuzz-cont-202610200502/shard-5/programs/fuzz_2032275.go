package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint16; F1 int64; F2 int8 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 float32 = float32(0.1)
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint64 = 1
        v7 += v7
        _ = fi6
        _ = v7
    }
    v8 := []int{0, 32767, 9223372036854775807, 0}
    v8 = append(v8, 127)
    _ = v8
    return 116
}

func main() {
    var v1 int = 773
    v2 := make(map[int]int)
    v3 := []string{"hello", "a", "", "foo bar"}
    var v9 int = 1
    var v10 float64 = 3.14159
    v11 := GenS1{F0: (!true), F1: 945}
    if (0 <= 0) {
        v12 := GenS2{F0: (v9 ^ v1), F1: 2147483647, F2: 1}
        v1 = 315
        _ = v12
    } else {
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        switch v9 {
        case 9223372036854775807:
            if ((735 < 629) || false) {
                var v13 int16 = 1000
                _ = v13
            }
        case 621:
            v14 := v2[(v1 | v9)]
            _ = v14
        case -239:
            v1 = 1
        default:
            var v15 float64 = (v10 + 3.14159)
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v17 := GenS2{F0: v1, F1: 456, F2: -3}
        _ = fi16
        _ = v17
    }
    switch ("" + "a") {
    case "hello":
        for fi18 := 0; fi18 < 1; fi18++ {
            switch "hello" {
            case "a":
                v2[fi18] = 105
            }
            _ = fi18
        }
    default:
        v19 := GenS3{F0: 2, F1: -1, F2: -1}
        _ = v19
    }
    var v20 uint8 = 1
    v20 -= v20
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
