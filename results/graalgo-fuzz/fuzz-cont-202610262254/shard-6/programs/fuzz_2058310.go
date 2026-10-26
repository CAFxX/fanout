package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int; F2 int16 }
func genFn0(p0 uint8) int64 {
    v4 := GenS1{F0: "a", F1: "foo bar"}
    v5 := GenS2{F0: 0, F1: 2147483647}
    p0 += 128
    v6 := GenS3{F0: -109}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 int32 = -801
    v2 := []int{-9223372036854775808, 949, -1}
    var v3 uint16 = 42
    var v7 float32 = 0.0
    v3 += v3
    var v8 int = 169
    fmt.Println(v1, v8, v1)
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 uint64 = 68
        if false {
            var v11 float64 = 3.14159
            v1 += v1
            _ = v11
        } else {
            if true {
                fmt.Println(v8)
            } else {
                v7 -= v7
                if false {
                    switch (v8 - v8) {
                    case -1:
                        switch 2147483647 {
                        case -9223372036854775808:
                            v12 := GenS4{F0: 2}
                            _ = v12
                        case 32767:
                            fmt.Println(v1, v10)
                        }
                    case -2147483648:
                        var v13 int8 = 56
                        _ = v13
                    case 1:
                        var v14 uint64 = v10
                        _ = v14
                    default:
                        v15 := GenS2{F0: 65535, F1: v8}
                        _ = v15
                    }
                    v2 = append(v2, (-(fi9)))
                }
                switch fi9 {
                case 1:
                    switch fi9 {
                    case 9223372036854775807:
                        for fi16 := 0; fi16 < 4; fi16++ {
                            if len(v2) > 0 {
                                v2[0] = fi16
                            }
                            _ = fi16
                        }
                    case 0:
                        v17 := GenS5{F0: 956, F1: (v8 >> ((18446744073709551615) & 63)), F2: 182}
                        _ = v17
                    case 744:
                        var v18 float32 = (v7 / (((float32(1e10) - 0.0) - float32((3.14159 / -2.25))) * (((((v7 + (float32((1.5 / ((0.1 + 1.5) * 1e10))) - -2.25)) * 3.14159) + float32(-2.25)) * (float32(-2.25) + v7)) + ((1.5 * v7) + float32((0.0 + ((3.14159 + 1e10) + 0.0)))))))
                        _ = v18
                    default:
                        for fi19 := 0; fi19 < 13; fi19++ {
                            fmt.Println(v1, fi9, fi19)
                            var v20 uint64 = (v10 / (((v10 ^ (^(v10)))) | 1))
                            _ = fi19
                            _ = v20
                        }
                    }
                case 916:
                    v10 = 42
                }
            }
        }
        fmt.Println(v7)
        _ = fi9
        _ = v10
    }
    if (((v7 / 1.5) / 1.5) >= 1.5) {
        var v21 int = (v8 + v8)
        _ = v21
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3, v1, v7)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
