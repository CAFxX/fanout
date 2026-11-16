package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int; F1 uint16; F2 int }

func main() {
    var v1 string = "hello"
    var v2 int32 = 869
    var v3 float32 = 0.1
    var v4 uint8 = 0
    v5 := []uint32{}
    v5 = append(v5, 1)
    v4 = (v4 & (^(v4)))
    for fi6 := 0; fi6 < 19; fi6++ {
        switch v1 {
        case "a":
            if (-2.25 < v3) {
                v7 := GenS1{F0: 100}
                if (fi6 != (fi6 % (((fi6 * (fi6 - (fi6 % (((fi6 * fi6)) | 1))))) | 1))) {
                    if false {
                        v2 += 1
                    }
                    v3 = ((float32((0.1 * -2.25)) / v3) + float32((((3.14159 - -2.25) * (1e10 * (((((3.14159 + 1e10) / (1e10 * float64(v3))) * float64(1.5)) / (1e10 * (1e10 / 3.14159))) / 0.1))) * float64(v3))))
                } else {
                    v2 -= v2
                    var v8 uint32 = 2147483648
                    var v9 bool = false
                    _ = v8
                    _ = v9
                }
                var v10 int8 = 127
                _ = v10
                _ = v7
            }
        case "hello":
            var v11 int8 = -47
            _ = v11
        default:
            v1 = (v1 + "hello")
        }
        _ = fi6
    }
    if (18446744073709551615 == 800) {
        v2 += v2
    }
    fmt.Println(v2, v1, v4)
    fmt.Println(v4, v2)
    if (!(!false)) {
        v5 = append(v5, uint32(v4))
        var v12 uint32 = 1
        v4 = (v4 | 132)
        _ = v12
    }
    var v13 float32 = (v3 - ((v3 * v3) - v3))
    switch "\n\t" {
    case "":
        var v14 bool = false
        _ = v14
    case "a":
        v13 = v3
    }
    v2 = 42
    for fi15 := 0; fi15 < 2; fi15++ {
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if true {
                if (!(2 >= 42)) {
                    v16 := GenS2{F0: 42, F1: 61}
                    if true {
                        v17 := GenS1{F0: -128}
                        var v18 int = (fi15 % ((905) | 1))
                        var v19 uint8 = (v4 / ((v4) | 1))
                        _ = v17
                        _ = v18
                        _ = v19
                    } else {
                        v20 := GenS1{F0: 100}
                        var v21 float32 = float32((-2.25 * 1e10))
                        switch 128 {
                        case 811:
                            if false {
                                var v22 string = v1
                                switch 4294967295 {
                                case -9223372036854775808:
                                    v22 = (v22 + "\n\t")
                                }
                                _ = v22
                            }
                        case 573:
                            fmt.Println(v4)
                        }
                        _ = v20
                        _ = v21
                    }
                    _ = v16
                } else {
                    if (1 <= 785) {
                        v23 := GenS3{F0: (fi15 - 4294967295), F1: 65535, F2: (-(fi15))}
                        v4 = (v4 - 0)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v23
                    }
                    if len(v5) > 0 {
                        v5[0] = 2147483647
                    }
                }
            } else {
                fmt.Println(v13)
            }
            for fi24 := 0; fi24 < 5; fi24++ {
                _ = len(v5)
                _ = fi24
            }
        } else {
            fmt.Println(fi15, v1)
            var v25 int64 = 394
            _ = v25
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
