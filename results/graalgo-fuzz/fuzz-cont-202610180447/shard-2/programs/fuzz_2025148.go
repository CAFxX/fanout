package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int16; F1 int64; F2 int64 }
type GenS3 struct { F0 int32; F1 uint32; F2 int }
func genFn0(p0 int8) uint {
    p0 += p0
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 uint8 = 106
        _ = fi4
        _ = v5
    }
    _ = p0
    return 667
}
func genFn1(p0 uint8) int32 {
    var v6 string = "foo bar"
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int16 = 32767
        _ = fi7
        _ = v8
    }
    if false {
        var v9 int64 = 9223372036854775807
        fmt.Println(v9)
        _ = v9
    }
    p0 = p0
    p0 -= p0
    _ = p0
    _ = v6
    return 2147483647
}

func main() {
    var v1 float32 = (float32(0.1) / 1.5)
    var v2 int8 = -1
    v3 := []int{128, 9223372036854775807, 192}
    v10 := GenS1{F0: 42, F1: 166}
    var v11 int64 = 0
    if (0 <= 1) {
        var v12 string = "foo bar"
        v13 := GenS2{F0: -1, F1: 1, F2: (v11 % (((v11 % (((v11 * v11)) | 1))) | 1))}
        switch v12 {
        case "hello":
            for fi14 := 0; fi14 < 11; fi14++ {
                switch fi14 {
                case 9223372036854775807:
                    v1 += (float32(3.14159) * v1)
                }
                fmt.Println(v12, v2, v12)
                switch "" {
                case "\n\t":
                    fmt.Println(v12, fi14, v11)
                case "hello":
                    v2 -= -1
                default:
                    var v15 int16 = -32768
                    _ = v15
                }
                _ = fi14
            }
        case "":
            v3 = append(v3, 2147483647)
        }
        _ = v12
        _ = v13
    } else {
        v1 += (float32(((0.1 * 3.14159) * 0.1)) / ((v1 + (v1 / (v1 / (v1 - v1)))) + v1))
    }
    var v16 int = -290
    v17 := GenS2{F0: -1, F1: v11, F2: (v11 << ((986) & 63))}
    v18 := GenS3{F0: 1, F1: 2147483647, F2: (v16 + -639)}
    v1 = float32((-2.25 / (-2.25 - (1e10 / (0.1 - 3.14159)))))
    if (0 >= 42) {
        var v19 int16 = 1
        fmt.Println(v19, v1)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v19
    } else {
        var v20 uint8 = 127
        if false {
            switch 128 {
            case -2147483648:
                v16 = v16
            default:
                var v21 int16 = 39
                _ = v21
            }
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
