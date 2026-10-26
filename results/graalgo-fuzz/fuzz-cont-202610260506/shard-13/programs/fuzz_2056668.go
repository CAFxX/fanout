package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int32; F2 uint }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 string; F1 uint }
func genFn0(p0 int16, p1 uint16, p2 uint) int32 {
    switch "hello" {
    case "a":
        v6 := GenS1{F0: 439}
        _ = v6
    case "\n\t":
        var v7 float64 = 0.0
        _ = v7
    case "":
        var v8 float32 = -2.25
        _ = v8
    }
    fmt.Println(p0, p0, p2)
    fmt.Println(p1, p2)
    p0 += (p0 % (((p0 / ((1) | 1))) | 1))
    for fi9 := 0; fi9 < 17; fi9++ {
        if false {
            v10 := GenS1{F0: 0}
            var v11 float32 = float32(0.1)
            v12 := []int{fi9, -480, 1, 826}
            _ = v10
            _ = v11
            _ = v12
        } else {
            if true {
                v13 := []int{1, fi9}
                fmt.Println(p0, p2)
                _ = v13
            }
            switch fi9 {
            case 550:
                if false {
                    v14 := GenS1{F0: 1}
                    var v15 int64 = 0
                    var v16 float32 = float32(3.14159)
                    _ = v14
                    _ = v15
                    _ = v16
                }
            case 32767:
                p1 += (p1 / ((p1) | 1))
            }
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[int]string)
    var v3 string = "hello"
    v4 := []uint32{0}
    var v5 float64 = ((1.5 - 1e10) + 1.5)
    if true {
        v2[119] = ((("hello" + (v3 + "hello")) + v3) + ("hello" + (v3 + "\n\t")))
        fmt.Println(v1, v1)
    } else {
        var v17 int32 = 42
        v18 := GenS2{F0: -1, F1: v17, F2: 2}
        fmt.Println(v5, v3)
        _ = v17
        _ = v18
    }
    var v19 uint = 4294967295
    switch 32768 {
    case 27:
        v5 += v5
    case 127:
        if ((1.5 - 0.1) < ((3.14159 * float32((3.14159 - (v5 - (float64(((1.5 - -2.25) * (0.1 - (float32((float64(((float32(0.0) - (float32(0.0) / (1e10 / 1.0))) + (0.0 / (float32(v5) + (0.1 / 1e10))))) - 0.1)) + 1e10)))) + (v5 - v5)))))) * (1e10 * 3.14159))) {
            var v20 uint = (v19 / (((v19 % ((18446744073709551615) | 1))) | 1))
            v5 = v5
            _ = v20
        }
    case -920:
        if (2 < 0) {
            v19 = (^(v19))
        }
    default:
        if true {
            if (false || (42 > 127)) {
                v1 = 94
                fmt.Println(v3, v1)
                switch v3 {
                case "hello":
                    v21 := v2[0]
                    _ = v21
                case "a":
                    v22 := GenS3{F0: v1, F1: -614}
                    _ = v22
                default:
                    if (-921 == -32768) {
                        var v23 uint = v19
                        v4 = append(v4, 2)
                        _ = v23
                    } else {
                        v3 = "a"
                    }
                }
            }
            v24 := v2[127]
            _ = v24
        }
    }
    fmt.Println(v5, v19, v1)
    for fi25 := 0; fi25 < 1; fi25++ {
        for fi26 := 0; fi26 < 14; fi26++ {
            v19 = v19
            _ = fi26
        }
        _ = fi25
    }
    v27 := GenS3{F0: 127, F1: -9223372036854775808}
    if ((v19 <= v19) || (-604 >= -217)) {
        if true {
            _ = len(v4)
            var v28 float32 = float32(3.14159)
            _ = v28
        }
        if (65535 == 869) {
            v4 = append(v4, 0)
        } else {
            fmt.Println(v19, v5, v19)
            var v29 float32 = (((3.14159 - (-2.25 / -2.25)) - 1e10) + 0.0)
            _ = v29
        }
        v30 := GenS4{F0: v3, F1: v19}
        _ = v30
    }
    v5 -= 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
