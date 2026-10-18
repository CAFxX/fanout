package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint; F1 int8; F2 int }
type GenS3 struct { F0 bool; F1 uint32; F2 uint16 }

func main() {
    var v1 string = ("\n\t" + ("hello" + ("" + "")))
    var v2 int = 128
    var v3 int8 = -128
    var v4 bool = true
    v5 := GenS1{F0: 65535, F1: (v1 + v1)}
    if (-32768 > 150) {
        fmt.Println(v1)
        switch 0 {
        case -2147483648:
            if (437 <= 2147483647) {
                for fi6 := 0; fi6 < 9; fi6++ {
                    for fi7 := 0; fi7 < 19; fi7++ {
                        v1 = "foo bar"
                        switch "\n\t" {
                        case "hello":
                            v3 = (v3 / ((-32) | 1))
                        default:
                            if v4 {
                                if (42 >= 42) {
                                    v8 := GenS1{F0: 70, F1: v1}
                                    _ = v8
                                }
                            } else {
                                if false {
                                    switch v1 {
                                    case "hello":
                                        v9 := GenS1{F0: 556, F1: v1}
                                        _ = v9
                                    case "":
                                        if (!(false || (-1 != 350))) {
                                            if (1 >= 18446744073709551615) {
                                                if (1 < 0) {
                                                    v10 := []int{-2147483648, -79}
                                                    var v11 bool = v4
                                                    _ = v10
                                                    _ = v11
                                                } else {
                                                    v4 = v4
                                                    fmt.Println(v3, v1)
                                                    v2 = (fi7 + (fi7 * (fi7 / ((fi6) | 1))))
                                                }
                                            } else {
                                                v4 = v4
                                                v3 += 56
                                            }
                                        } else {
                                            switch v1 {
                                            case "foo bar":
                                                v3 += v3
                                            case "hello":
                                                v3 += -128
                                            case "a":
                                                v12 := []int{(fi6 % (((v2 | (fi7 << ((0) & 63)))) | 1)), fi6, 42, fi6}
                                                _ = v12
                                            default:
                                                switch fi7 {
                                                case -9223372036854775808:
                                                    v4 = false
                                                    switch v2 {
                                                    case -760:
                                                        var v13 float64 = 3.14159
                                                        _ = v13
                                                    default:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v3 -= v3
                                                        default:
                                                            if (84 < 100) {
                                                                switch v1 {
                                                                case "a":
                                                                    var v14 uint32 = 821
                                                                    _ = v14
                                                                case "hello":
                                                                    var v15 float64 = (-2.25 + 1e10)
                                                                    v2 = 70
                                                                    _ = v15
                                                                default:
                                                                    v2 = 4294967295
                                                                }
                                                                var v16 float64 = -2.25
                                                                if v4 {
                                                                    v17 := GenS2{F0: 268, F1: 5, F2: fi7}
                                                                    _ = v17
                                                                }
                                                                _ = v16
                                                            } else {
                                                                v1 = v1
                                                            }
                                                        }
                                                    }
                                                case 563:
                                                    fmt.Println(fi7)
                                                }
                                                v4 = (!(813 > 2))
                                                var v18 uint32 = 431
                                                v2 = fi6
                                                v4 = (0 > 46)
                                                v19 := []int{fi6, fi6, 793}
                                                _ = v18
                                                _ = v19
                                            }
                                        }
                                    }
                                    switch v1 {
                                    case "":
                                        v1 = "foo bar"
                                        v4 = v4
                                    default:
                                        var v20 float64 = 1.5
                                        _ = v20
                                    }
                                }
                            }
                        }
                        _ = fi7
                    }
                    var v21 uint32 = 65535
                    _ = fi6
                    _ = v21
                }
                if ((1.5 + 0.1) >= 3.14159) {
                    for fi22 := 0; fi22 < 8; fi22++ {
                        v23 := []int{}
                        fmt.Println(v1, fi22, v4)
                        _ = fi22
                        _ = v23
                    }
                    v24 := []int{}
                    if len(v24) > 0 {
                        v24[0] = v2
                    }
                    _ = v24
                }
            }
        case 1:
            fmt.Println(v2)
        }
        v4 = v4
    }
    v1 = v1
    v25 := GenS2{F0: 42, F1: v3, F2: (v2 * (v2 - (v2 % ((-371) | 1))))}
    var v26 int16 = -955
    v27 := GenS3{F0: ((!(2 >= 24)) && (int32(v2) != 90)), F1: 564, F2: 1}
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 bool = v4
        v4 = true
        var v30 uint32 = 878
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
