package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 uint; F1 string; F2 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int32; F1 bool; F2 int16 }
type GenS7 struct { F0 int32; F1 bool; F2 int32 }
func genFn0(p0 int8, p1 uint8, p2 int) int8 {
    var v6 int32 = 741
    v6 = -784
    if (!(true || (true && (float64(1.5) == 0.1)))) {
        var v7 uint8 = 42
        v8 := GenS1{F0: uint(v7)}
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}
func genFn1(p0 int16) bool {
    var v9 uint32 = 42
    v10 := GenS1{F0: 1}
    if true {
        v9 -= v9
        fmt.Println(v9)
    }
    switch -712 {
    case 2147483647:
        v11 := []int{-1, 0, 128, -787}
        _ = v11
    default:
        p0 -= p0
    }
    _ = p0
    _ = v10
    _ = v9
    return (true || true)
}

func main() {
    var v1 int32 = -1
    var v2 float32 = 0.1
    v3 := []uint32{1}
    var v4 uint = 2
    var v5 float32 = (0.1 + v2)
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS1{F0: 997}
        v5 = -2.25
        if (!true) {
            for fi14 := 0; fi14 < 14; fi14++ {
                switch 837 {
                case 0:
                    v2 = v2
                }
                fmt.Println(v1, fi14, fi14)
                fmt.Println(fi14, fi12, v4)
                _ = fi14
            }
            v5 = 0.1
        } else {
            for fi15 := 0; fi15 < 14; fi15++ {
                v4 = v4
                v16 := GenS1{F0: 942}
                fmt.Println(v5, v5)
                _ = fi15
                _ = v16
            }
        }
        _ = fi12
        _ = v13
    }
    v2 = float32((0.1 * 3.14159))
    v3 = append(v3, 65535)
    v1 -= v1
    var v17 int16 = 0
    for fi18 := 0; fi18 < 18; fi18++ {
        v2 += v2
        _ = fi18
    }
    if true {
        var v19 uint32 = 533
        if false {
            if (v17 != -789) {
                fmt.Println(v2, v1)
                v5 = ((v5 * float32(((((((0.1 * ((float64(((float32(0.0) * float32(-2.25)) - v5)) * float64(v5)) - 0.1)) * 0.1) / 3.14159) + 1e10) * float64(float32(3.14159))) - 1.5))) * v2)
                v20 := GenS1{F0: (^(v4))}
                _ = v20
            }
            fmt.Println(v19, v5)
        } else {
            if (true || (0 <= -1)) {
                var v21 uint32 = v19
                _ = v21
            }
        }
        _ = v19
    }
    if ("a" >= "\n\t") {
        v22 := GenS2{F0: 65535, F1: 65535}
        _ = v22
    } else {
        _ = len(v3)
        for fi23 := 0; fi23 < 16; fi23++ {
            v5 -= v5
            if true {
                var v24 float64 = 1.5
                if true {
                    if false {
                        var v25 int64 = -9223372036854775808
                        switch fi23 {
                        case 32767:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case 1:
                            v26 := GenS3{F0: 0, F1: -272}
                            _ = v26
                        case 2147483647:
                            if false {
                                var v27 uint16 = 129
                                v25 += (-(v25))
                                switch "\n\t" {
                                case "a":
                                    fmt.Println(v27, v4)
                                default:
                                    var v28 uint = (v4 - v4)
                                    _ = v28
                                }
                                _ = v27
                            }
                        }
                        v25 += 42
                        _ = v25
                    } else {
                        v29 := GenS2{F0: 65535, F1: 0}
                        switch "" {
                        case "":
                            v17 -= (^(v17))
                        case "a":
                            _ = len(v3)
                        case "foo bar":
                            for fi30 := 0; fi30 < 16; fi30++ {
                                switch "" {
                                case "":
                                    v3 = append(v3, 1)
                                case "hello":
                                    v17 += v17
                                case "foo bar":
                                    var v31 string = "\n\t"
                                    _ = v31
                                default:
                                    switch fi23 {
                                    case 794:
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case 1:
                                        if false {
                                            var v32 uint = v4
                                            v33 := GenS4{F0: (v4 * 0), F1: "", F2: 127}
                                            _ = v32
                                            _ = v33
                                        }
                                    }
                                }
                                _ = fi30
                            }
                        }
                        v17 = v17
                        _ = v29
                    }
                    switch ("foo bar" + "\n\t") {
                    case "":
                        v24 -= (v24 + (float64((3.14159 * (((0.1 + v5) + float32(v24)) * v5))) + (v24 / 1.5)))
                    case "hello":
                        if (65535 < 2) {
                            if (2 < 1) {
                                if len(v3) > 0 {
                                    v3[0] = 445
                                }
                                if true {
                                    _ = len(v3)
                                    v34 := GenS5{F0: v1}
                                    switch "" {
                                    case "a":
                                        for fi35 := 0; fi35 < 10; fi35++ {
                                            v24 = 0.1
                                            if len(v3) > 0 {
                                                v3[0] = 978
                                            }
                                            v36 := GenS3{F0: 99, F1: (^(v17))}
                                            _ = fi35
                                            _ = v36
                                        }
                                    case "":
                                        v17 = v17
                                    default:
                                        var v37 uint32 = 65535
                                        _ = v37
                                    }
                                    _ = v34
                                }
                                for fi38 := 0; fi38 < 3; fi38++ {
                                    if false {
                                        if len(v3) > 0 {
                                            v3[0] = 415
                                        }
                                        switch "foo bar" {
                                        case "":
                                            v39 := GenS1{F0: v4}
                                            v17 += 42
                                            fmt.Println(v24, v1)
                                            _ = v39
                                        default:
                                            v2 += v2
                                        }
                                    }
                                    _ = fi38
                                }
                            } else {
                                v24 -= (v24 / v24)
                            }
                            for fi40 := 0; fi40 < 18; fi40++ {
                                fmt.Println(v1, fi40)
                                _ = fi40
                            }
                        }
                    case "a":
                        for fi41 := 0; fi41 < 11; fi41++ {
                            var v42 uint8 = 0
                            _ = fi41
                            _ = v42
                        }
                    default:
                        for fi43 := 0; fi43 < 16; fi43++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (428 <= 9223372036854775807) {
                                var v44 uint16 = 127
                                var v45 int64 = 603
                                _ = v44
                                _ = v45
                            }
                            v3 = append(v3, 879)
                            _ = fi43
                        }
                    }
                    if false {
                        v3 = append(v3, 2147483647)
                    } else {
                        v46 := GenS6{F0: v1, F1: false, F2: v17}
                        var v47 int16 = 1
                        v3 = append(v3, 985)
                        _ = v46
                        _ = v47
                    }
                }
                _ = v24
            }
            v1 = (v1 - 0)
            _ = fi23
        }
        switch 847 {
        case -153:
            switch ("\n\t" + ("hello" + "a")) {
            case "hello":
                if (990 < 169) {
                    v17 = v17
                    v48 := GenS3{F0: 100, F1: v17}
                    _ = v48
                } else {
                    v3 = append(v3, 1)
                    v49 := GenS5{F0: -556}
                    var v50 int32 = v1
                    _ = v49
                    _ = v50
                }
            case "":
                v1 = (v1 | (v1 / (((v1 ^ v1)) | 1)))
            }
        case -751:
            v51 := GenS3{F0: -128, F1: v17}
            _ = v51
        case -872:
            v52 := GenS7{F0: 255, F1: ((313 != 2147483647) && true), F2: v1}
            _ = v52
        }
    }
    var v53 uint = (v4 / (((v4 << ((65535) & 63))) | 1))
    if true {
        fmt.Println(v5, v5, v2)
    } else {
        v4 = (v4 & 18446744073709551615)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
}
