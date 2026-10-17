package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint32, p1 uint8) int {
    var v5 uint16 = 953
    if (2 == 2) {
        v6 := []int{32768, -717, 32767, 2147483647}
        v5 = (^(v5))
        if (!(true && (0 >= 275))) {
            v7 := GenS1{F0: (!(-324 >= 158)), F1: 42, F2: 42}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 18; fi8++ {
                if (((3.14159 * float64(3.14159)) - 0.0) < 0.1) {
                    v6 = append(v6, fi8)
                }
                switch int(v5) {
                case 2147483647:
                    if len(v6) > 0 {
                        v6[0] = fi8
                    }
                case -1:
                    for fi9 := 0; fi9 < 6; fi9++ {
                        var v10 int16 = 1
                        p0 += (p0 - p0)
                        _ = fi9
                        _ = v10
                    }
                case -2147483648:
                    v5 = v5
                default:
                    if ((v5 & v5) == (^(v5))) {
                        v11 := GenS1{F0: ((float64((((((0.1 + (3.14159 * 3.14159)) * (1e10 / 1e10)) - 3.14159) / (1.5 / (((float32((1.5 * 0.0)) + 0.0) + (3.14159 * (-2.25 + 1e10))) - 1e10))) - 3.14159)) - -2.25) != (-2.25 + (0.1 + -2.25))), F1: 175, F2: 98}
                        if len(v6) > 0 {
                            v6[0] = fi8
                        }
                        _ = v11
                    }
                }
                _ = fi8
            }
        }
        _ = v6
    }
    switch 2147483647 {
    case -9223372036854775808:
        if (false && false) {
            v12 := []int{0}
            _ = v12
        }
    case 32768:
        v13 := []int{0, 0, 32768, -2147483648}
        _ = v13
    }
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v5
    return -2147483648
}

func main() {
    var v1 bool = true
    var v2 float64 = (0.1 - -2.25)
    var v3 float64 = v2
    var v4 uint16 = 255
    v2 = v2
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint = 1
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := []int{}
            fmt.Println(v2, v3)
            v18 := GenS2{F0: 65535, F1: (v15 / (((v15 >> ((2) & 63))) | 1))}
            if ("" == "a") {
                if len(v17) > 0 {
                    v17[0] = fi14
                }
                switch "foo bar" {
                case "":
                    if v1 {
                        var v19 int16 = 1000
                        var v20 int32 = 1
                        _ = len(v17)
                        _ = v19
                        _ = v20
                    }
                }
                v4 += uint16(v15)
            }
            v1 = (1.5 >= ((3.14159 * ((-2.25 * float32(float64((0.1 / ((3.14159 / float32(v2)) * -2.25))))) + float32((1.5 + (v3 / float64(0.1)))))) - 1.5))
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS3{F0: 32767, F1: 1}
            v23 := []int{fi14, (fi21 + (fi21 % (((fi14 + (fi14 / (((-(fi21))) | 1)))) | 1))), fi21}
            fmt.Println(v4, v2, v2)
            v2 = 3.14159
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi14
        _ = v15
    }
    v24 := GenS3{F0: (v4 % (((v4 - 65535)) | 1)), F1: 0}
    v25 := GenS3{F0: v4, F1: -2147483648}
    v2 = ((v2 / (v3 - v2)) - float64((0.0 * 0.1)))
    v26 := GenS3{F0: (v4 & v4), F1: 671}
    var v27 float64 = 0.1
    fmt.Println(v4, v4, v4)
    v4 = (^(v4))
    var v28 int32 = -1
    switch 32768 {
    case 114:
        v4 = (^(v4))
    default:
        v3 = (1e10 - float64(float32((v27 * v2))))
    }
    v29 := GenS4{F0: (-(v28))}
    switch ("hello" + ("a" + ("foo bar" + ("a" + "")))) {
    case "":
        v30 := GenS2{F0: 16, F1: uint(v4)}
        _ = v30
    case "hello":
        fmt.Println(v27)
    case "foo bar":
        v31 := []int{1, 1}
        _ = v31
    }
    v32 := []int{1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
}
