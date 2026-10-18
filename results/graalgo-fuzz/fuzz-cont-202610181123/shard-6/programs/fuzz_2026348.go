package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint8; F2 uint32 }
func genFn0(p0 int32) uint8 {
    if (0.1 > 0.1) {
        fmt.Println(p0)
    } else {
        v5 := GenS1{F0: 325, F1: p0}
        _ = v5
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{fi6, (fi6 % ((fi6) | 1))}
        for fi8 := 0; fi8 < 13; fi8++ {
            p0 += (p0 & p0)
            v9 := GenS2{F0: -128}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            if (true || false) {
                var v11 string = "foo bar"
                var v12 bool = (!false)
                _ = v11
                _ = v12
            }
            p0 -= (p0 - int32(fi6))
            fmt.Println(p0, p0)
            _ = fi10
        }
        _ = fi6
        _ = v7
    }
    p0 += p0
    _ = p0
    return 2
}

func main() {
    var v1 string = ("" + "")
    var v2 uint8 = 127
    v3 := []int{}
    var v4 float64 = 3.14159
    var v13 uint32 = uint32(v2)
    _ = len(v3)
    v14 := GenS1{F0: 2, F1: 42}
    v4 = ((float64((((((1e10 - 1e10) / 1e10) + (((0.1 - ((((-2.25 - 1.5) - 1e10) / 1.0) - (1.5 * float32((float64((1.5 - 0.1)) / v4))))) / (0.0 + 3.14159)) * 1e10)) - (float32((float64(1e10) - 0.1)) * 1.5)) * (float32(v4) + -2.25))) - (3.14159 / ((3.14159 * v4) - 3.14159))) * 0.0)
    for fi15 := 0; fi15 < 12; fi15++ {
        v2 = (v2 / (((^(v2))) | 1))
        switch "hello" {
        case "foo bar":
            v2 = 255
        case "\n\t":
            if (!false) {
                fmt.Println(fi15)
            } else {
                if true {
                    if (!(1 <= 1)) {
                        v1 = v1
                        v4 += float64(1.5)
                        switch v1 {
                        case "":
                            for fi16 := 0; fi16 < 20; fi16++ {
                                if (103 == -48) {
                                    v13 += (^(v13))
                                    v1 = "foo bar"
                                    v17 := GenS1{F0: 1, F1: -2147483648}
                                    fmt.Println(fi16)
                                    _ = v17
                                }
                                var v18 int32 = 2147483647
                                if len(v3) > 0 {
                                    v3[0] = (fi15 % ((fi15) | 1))
                                }
                                _ = fi16
                                _ = v18
                            }
                        case "foo bar":
                            if (v13 > v13) {
                                var v19 int = 32767
                                for fi20 := 0; fi20 < 11; fi20++ {
                                    var v21 int8 = -14
                                    var v22 uint16 = 0
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                                _ = v19
                            }
                        }
                    } else {
                        v2 -= 67
                        v23 := GenS3{F0: v1, F1: (^(v2)), F2: (^(v13))}
                        var v24 float32 = 0.1
                        _ = v23
                        _ = v24
                    }
                }
                var v25 uint32 = 44
                v4 -= v4
                _ = v25
            }
        case "hello":
            fmt.Println(fi15, fi15)
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi15
    }
    var v26 uint16 = 2
    var v27 int16 = 42
    var v28 int16 = 1000
    switch "a" {
    case "foo bar":
        v3 = append(v3, -9223372036854775808)
    default:
        var v29 int32 = -1
        _ = v29
    }
    v30 := GenS3{F0: "a", F1: (v2 + 1), F2: 4294967295}
    v26 = (v26 % (((v26 - v26)) | 1))
    fmt.Println(v2, v28)
    fmt.Println(v13, v13, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
}
