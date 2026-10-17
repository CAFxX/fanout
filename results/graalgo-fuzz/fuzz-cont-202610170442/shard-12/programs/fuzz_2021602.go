package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    v1 := []int{-1, 128, -83, 32768}
    var v2 bool = (!true)
    var v3 int8 = 100
    var v4 uint64 = 1
    var v5 uint8 = 0
    var v6 bool = true
    v1 = append(v1, 151)
    v7 := GenS1{F0: "foo bar", F1: 1}
    var v8 uint8 = v5
    v9 := GenS1{F0: "foo bar", F1: -299}
    if len(v1) > 0 {
        v1[0] = 968
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v2 = (0.1 == (((-2.25 - 1.5) * ((0.1 + 1e10) + 1e10)) + (1.5 - 0.0)))
        var v11 int8 = (v3 + (v3 << ((4294967295) & 63)))
        _ = fi10
        _ = v11
    }
    if (320 > -101) {
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(v2)
            _ = len(v1)
            v3 += 63
            _ = fi12
        }
    } else {
        if ((11 <= 389) && false) {
            v8 -= 0
        } else {
            v1 = append(v1, int(v8))
            v2 = ((2 >= 210) || v2)
        }
        v3 = 1
        v8 = v5
    }
    v5 = (v8 * (v8 >> ((873) & 63)))
    fmt.Println(v3, v5, v3)
    switch -1 {
    case 42:
        var v13 int64 = -9223372036854775808
        _ = v13
    default:
        for fi14 := 0; fi14 < 1; fi14++ {
            switch fi14 {
            case 0:
                fmt.Println(v3, v4, v3)
            default:
                for fi15 := 0; fi15 < 3; fi15++ {
                    if len(v1) > 0 {
                        v1[0] = 32768
                    }
                    _ = fi15
                }
            }
            if (-2.25 < (float64(((float32(1.5) * float32((0.0 - (float64(3.14159) / 1.0)))) * 0.1)) * ((1e10 / (1e10 + 3.14159)) - -2.25))) {
                v8 += (v5 * (v5 % (((v5 * v5)) | 1)))
                var v16 int = fi14
                _ = v16
            }
            _ = fi14
        }
    }
    v17 := GenS2{F0: uint32(v8), F1: 381}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
