package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 700
    var v3 int16 = 1
    var v4 float64 = float64(1.5)
    v5 := []uint32{2}
    if len(v5) > 0 {
        v5[0] = 42
    }
    v6 := GenS1{F0: 65535, F1: 963, F2: 65535}
    var v7 uint16 = 255
    v3 += -1
    var v8 int64 = (^(v2))
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            if false {
                v3 += (v3 >> ((18446744073709551615) & 63))
                v1[(fi10 >> ((1) & 63))] = (42 >= 787)
                v11 := v1[fi9]
                _ = v11
            } else {
                if false {
                    v12 := GenS1{F0: 42, F1: 255, F2: 42}
                    _ = v12
                }
                switch -9223372036854775808 {
                case 357:
                    var v13 uint16 = v7
                    _ = v13
                case 149:
                    if (true || ((v2 / (((v2 << ((1) & 63))) | 1)) < 0)) {
                        v5 = append(v5, 42)
                        v14 := GenS1{F0: 0, F1: 128, F2: 2147483648}
                        v2 += (v2 | (v8 * 887))
                        _ = v14
                    } else {
                        v3 = -727
                    }
                case 32768:
                    v15 := GenS1{F0: 485, F1: 475, F2: 65535}
                    _ = v15
                }
            }
            var v16 float64 = float64(3.14159)
            _ = len(v1)
            _ = fi10
            _ = v16
        }
        switch -335 {
        case -9223372036854775808:
            for fi17 := 0; fi17 < 3; fi17++ {
                v2 += (v8 / (((v8 % (((v8 - 1)) | 1))) | 1))
                _ = len(v5)
                v18 := GenS1{F0: 323, F1: 2, F2: 65535}
                _ = fi17
                _ = v18
            }
        case 373:
            var v19 float64 = ((v4 - (v4 / v4)) / v4)
            _ = v19
        case 2147483647:
            fmt.Println(v8, v4, v7)
        default:
            switch "a" {
            case "\n\t":
                for fi20 := 0; fi20 < 11; fi20++ {
                    v2 = v8
                    var v21 uint = 42
                    fmt.Println(v8, v3)
                    _ = fi20
                    _ = v21
                }
            case "hello":
                v22 := GenS2{F0: 4294967295}
                _ = v22
            case "":
                var v23 float32 = 3.14159
                _ = v23
            }
        }
        var v24 string = ("\n\t" + "")
        _ = fi9
        _ = v24
    }
    fmt.Println(v8, v3)
    v3 -= (v3 * (v3 - v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
