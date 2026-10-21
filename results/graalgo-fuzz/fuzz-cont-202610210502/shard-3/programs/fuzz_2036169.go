package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint; F2 int }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -128
    var v3 float64 = 0.1
    var v4 float64 = ((v3 * 0.0) * -2.25)
    v5 := make(map[string]uint32)
    v6 := make(map[string]bool)
    var v7 bool = (-1 != -1)
    var v8 uint64 = 702
    var v9 float32 = float32(float64(1.5))
    v1[-666] = 0
    v8 += (v8 & 454)
    switch -914 {
    case 2147483647:
        v10 := GenS1{F0: 0}
        _ = v10
    case 9:
        var v11 int64 = 9223372036854775807
        _ = v11
    case 9223372036854775807:
        v4 = v3
    }
    if (v8 != (v8 % ((v8) | 1))) {
        if v7 {
            v12 := GenS1{F0: 646}
            v7 = false
            switch -1 {
            case 733:
                var v13 uint = 2
                _ = v13
            default:
                switch "" {
                case "a":
                    v14 := []int{127, 4294967295, 42}
                    _ = v14
                default:
                    var v15 float64 = (0.0 + v3)
                    _ = v15
                }
            }
            _ = v12
        } else {
            if v7 {
                v1[128] = 32768
                fmt.Println(v8)
            } else {
                if (1 <= 928) {
                    v16 := GenS2{F0: 100, F1: -128}
                    v9 += (v9 * float32((float64(float32(3.14159)) - (0.0 + float64((float32(v3) * float32((((v4 + (-2.25 / 1.0)) - 3.14159) - 1e10))))))))
                    var v17 int16 = 1000
                    _ = v16
                    _ = v17
                } else {
                    var v18 float32 = (v9 * (((v9 * v9) / 1.5) / (float32(((float64(v9) + v3) + (0.0 + v4))) + float32(v4))))
                    _ = v18
                }
            }
            v7 = v7
            v19 := GenS2{F0: (v2 & 0), F1: 91}
            _ = v19
        }
    } else {
        if ((v2 % ((v2) | 1)) <= v2) {
            v4 -= v3
            switch 127 {
            case -227:
                delete(v5, "\n\t")
            case -9223372036854775808:
                var v20 uint64 = (v8 << ((4294967295) & 63))
                _ = v20
            case -933:
                v21 := GenS3{F0: (v8 * 1), F1: 943, F2: 4294967295}
                _ = v21
            }
        } else {
            v22 := []int{}
            var v23 float64 = v4
            _ = v22
            _ = v23
        }
    }
    var v24 uint16 = 497
    for fi25 := 0; fi25 < 13; fi25++ {
        var v26 int64 = 9223372036854775807
        for fi27 := 0; fi27 < 16; fi27++ {
            switch fi27 {
            case -499:
                v28 := []int{(fi25 + (fi25 << ((1) & 63))), (fi27 % (((fi25 / (((-(fi27))) | 1))) | 1)), 1}
                _ = v28
            default:
                v2 -= v2
            }
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
