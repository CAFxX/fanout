package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint32; F2 uint }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 uint64; F1 int64; F2 bool }
type GenS7 struct { F0 uint32; F1 int8; F2 int16 }
func genFn0() string {
    v4 := GenS1{F0: 89, F1: 4294967295, F2: 42}
    v5 := []int{}
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: -128, F1: 305, F2: 18446744073709551615}
            var v9 uint16 = 65535
            fmt.Println(v9, v9, fi6)
            if true {
                var v10 bool = (733 >= -548)
                if ((float64((0.1 + 0.1)) / 1.5) >= ((-2.25 - 0.0) / ((-2.25 / (float64(((0.0 * ((0.0 * 0.0) * 1e10)) - (1.5 * -2.25))) / 1e10)) / 3.14159))) {
                    switch "hello" {
                    case "hello":
                        _ = len(v5)
                    case "\n\t":
                        switch "\n\t" {
                        case "foo bar":
                            v9 += v9
                        default:
                            var v11 int8 = 56
                            _ = v11
                        }
                    case "a":
                        v10 = v10
                    }
                } else {
                    var v12 float32 = float32((3.14159 / 0.1))
                    _ = v12
                }
                if (v10 && v10) {
                    v13 := GenS1{F0: -1, F1: 380, F2: 275}
                    var v14 uint64 = 4294967295
                    _ = v13
                    _ = v14
                } else {
                    v10 = (!true)
                    switch 959 {
                    case -488:
                        var v15 uint16 = v9
                        var v16 uint8 = 2
                        _ = v15
                        _ = v16
                    }
                    v17 := GenS2{F0: 2, F1: 1}
                    v9 = 2
                    _ = v17
                }
                _ = v10
            } else {
                fmt.Println(fi6, fi7, fi7)
                _ = len(v5)
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        var v18 float32 = 1.5
        _ = fi6
        _ = v18
    }
    if true {
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 int16 = 492
            for fi21 := 0; fi21 < 8; fi21++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if false {
                    var v22 uint16 = 255
                    v23 := GenS2{F0: 255, F1: 0}
                    if len(v5) > 0 {
                        v5[0] = 115
                    }
                    _ = v22
                    _ = v23
                } else {
                    fmt.Println(fi21, fi19)
                }
                if true {
                    fmt.Println(v20)
                    var v24 uint = 255
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = v24
                }
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
        var v25 string = ("\n\t" + "")
        var v26 uint8 = 2
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v5 = append(v5, fi27)
        _ = fi27
    }
    _ = v4
    _ = v5
    return "hello"
}
func genFn1(p0 uint, p1 bool, p2 uint) int {
    p2 = p0
    for fi28 := 0; fi28 < 20; fi28++ {
        p0 = p0
        _ = fi28
    }
    if p1 {
        fmt.Println(p1)
        switch "\n\t" {
        case "hello":
            fmt.Println(p1, p2)
        default:
            var v29 uint16 = 172
            _ = v29
        }
    }
    fmt.Println(p0)
    v30 := GenS2{F0: 1, F1: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 127
}

func main() {
    var v1 uint = 42
    v2 := make(map[int]bool)
    var v3 uint = v1
    for fi31 := 0; fi31 < 18; fi31++ {
        fmt.Println(v3, v3, fi31)
        v32 := []int{(fi31 * (fi31 ^ (^(fi31)))), (fi31 >> ((v3) & 63)), 128}
        v3 -= (v3 & 18446744073709551615)
        _ = fi31
        _ = v32
    }
    v3 = (v3 << ((255) & 63))
    var v33 int32 = 1
    if (183 == 922) {
        v1 = v3
    }
    delete(v2, genFn1((v3 ^ 42), false, v1))
    var v34 string = (("" + "foo bar") + "\n\t")
    v35 := GenS1{F0: 100, F1: 2, F2: 688}
    v36 := GenS2{F0: 2, F1: 42}
    _ = len(v2)
    switch v34 {
    case "\n\t":
        v37 := GenS1{F0: -1, F1: 2147483648, F2: 4294967295}
        _ = v37
    case "hello":
        switch -342 {
        case -9223372036854775808:
            var v38 string = v34
            _ = v38
        case 0:
            fmt.Println(v3)
        case -2147483648:
            v3 -= (v1 ^ 65535)
        default:
            if true {
                if true {
                    v39 := []int{32768}
                    v3 = (v1 & 18446744073709551615)
                    _ = v39
                } else {
                    v40 := []int{}
                    if true {
                        _ = len(v2)
                        var v41 int = 31
                        if (0 < 1) {
                            var v42 int64 = -1000
                            var v43 string = "hello"
                            for fi44 := 0; fi44 < 2; fi44++ {
                                for fi45 := 0; fi45 < 5; fi45++ {
                                    if len(v40) > 0 {
                                        v40[0] = fi44
                                    }
                                    _ = fi45
                                }
                                _ = len(v2)
                                var v46 int32 = (v33 | v33)
                                _ = fi44
                                _ = v46
                            }
                            _ = v42
                            _ = v43
                        }
                        _ = v41
                    } else {
                        v47 := GenS3{F0: v3, F1: 428}
                        for fi48 := 0; fi48 < 13; fi48++ {
                            switch 32768 {
                            case 127:
                                if (true && false) {
                                    v49 := v2[127]
                                    _ = v49
                                }
                            case 4294967295:
                                v50 := GenS3{F0: v1, F1: 18446744073709551615}
                                _ = v50
                            case 490:
                                if (false && (v3 >= v1)) {
                                    v51 := GenS4{F0: v3, F1: 42, F2: 161}
                                    v52 := GenS4{F0: v3, F1: 42, F2: v3}
                                    v1 = 377
                                    _ = v51
                                    _ = v52
                                } else {
                                    if true {
                                        for fi53 := 0; fi53 < 14; fi53++ {
                                            v54 := GenS2{F0: 218, F1: -9223372036854775808}
                                            switch (fi53 / ((fi48) | 1)) {
                                            case -1:
                                                var v55 float64 = 0.1
                                                _ = v55
                                            case 1:
                                                var v56 int = (^(fi48))
                                                _ = v56
                                            case 42:
                                                v2[fi53] = (!true)
                                            }
                                            _ = fi53
                                            _ = v54
                                        }
                                        if true {
                                            v57 := GenS4{F0: v3, F1: 99, F2: 0}
                                            _ = v57
                                        } else {
                                            if true {
                                                fmt.Println(v3, v3, v3)
                                            }
                                            v3 = v3
                                        }
                                    } else {
                                        if (!true) {
                                            for fi58 := 0; fi58 < 20; fi58++ {
                                                v59 := GenS5{F0: 32767, F1: 128}
                                                fmt.Println(fi58, v1)
                                                _ = fi58
                                                _ = v59
                                            }
                                            v33 -= (^(v33))
                                        }
                                    }
                                    switch (fi48 / ((int(v33)) | 1)) {
                                    case -10:
                                        v1 = 4294967295
                                    }
                                    fmt.Println(fi48, fi48, v3)
                                }
                            }
                            v60 := GenS4{F0: 0, F1: 1, F2: v1}
                            _ = fi48
                            _ = v60
                        }
                        fmt.Println(v3)
                        _ = v47
                    }
                    if len(v40) > 0 {
                        v40[0] = 913
                    }
                    _ = v40
                }
                for fi61 := 0; fi61 < 1; fi61++ {
                    switch fi61 {
                    case -247:
                        if true {
                            v1 = 311
                            fmt.Println(v1)
                            if (true && ((128 >= 127) || false)) {
                                for fi62 := 0; fi62 < 4; fi62++ {
                                    v3 += (v3 * (^(v3)))
                                    v3 = v1
                                    var v63 int64 = int64(v33)
                                    _ = fi62
                                    _ = v63
                                }
                                delete(v2, (^(fi61)))
                            } else {
                                v1 = v3
                            }
                        }
                    case -286:
                        v64 := GenS2{F0: 137, F1: -1}
                        _ = v64
                    case 42:
                        fmt.Println(v1, v33)
                    }
                    if ((((float64((0.1 + ((1e10 * 1e10) * float32(-2.25)))) * 3.14159) - 0.0) / -2.25) >= 0.0) {
                        fmt.Println(v3, v3, v34)
                        v34 = ("foo bar" + "foo bar")
                        v34 = (v34 + (("a" + v34) + v34))
                    } else {
                        fmt.Println(v34, v3, v1)
                        var v65 string = "a"
                        v33 += 42
                        _ = v65
                    }
                    _ = fi61
                }
            }
        }
    case "":
        var v66 int8 = -124
        _ = v66
    default:
        fmt.Println(v3, v1)
    }
    var v67 int8 = 25
    var v68 uint16 = uint16(v1)
    for fi69 := 0; fi69 < 6; fi69++ {
        for fi70 := 0; fi70 < 13; fi70++ {
            v1 = 0
            v71 := GenS1{F0: v67, F1: 2147483647, F2: 53}
            v67 += v67
            _ = fi70
            _ = v71
        }
        for fi72 := 0; fi72 < 12; fi72++ {
            _ = len(v2)
            _ = fi72
        }
        if (507 < -1) {
            switch 1 {
            case 128:
                if true {
                    v73 := GenS6{F0: 2, F1: 9223372036854775807, F2: (2 <= 0)}
                    _ = v73
                }
            case 4294967295:
                _ = len(v2)
            case 136:
                v74 := []int{9223372036854775807, (fi69 + 2147483647), (fi69 >> (((v1 % (((^(v1))) | 1))) & 63))}
                _ = v74
            default:
                v3 += 46
            }
            if true {
                for fi75 := 0; fi75 < 1; fi75++ {
                    v67 = v67
                    v3 += v1
                    v76 := GenS6{F0: 18446744073709551615, F1: -1, F2: ((fi69 ^ (fi75 << ((65535) & 63))) > 1)}
                    _ = fi75
                    _ = v76
                }
                v3 = v1
                v77 := GenS4{F0: v1, F1: 4294967295, F2: v3}
                _ = v77
            } else {
                v78 := []int{fi69}
                for fi79 := 0; fi79 < 8; fi79++ {
                    v80 := GenS7{F0: 2, F1: (-(v67)), F2: 0}
                    v78 = append(v78, fi79)
                    v68 -= 741
                    _ = fi79
                    _ = v80
                }
                _ = v78
            }
            for fi81 := 0; fi81 < 8; fi81++ {
                var v82 uint8 = 2
                _ = fi81
                _ = v82
            }
        }
        _ = fi69
    }
    if (-2.25 == (1e10 * 1e10)) {
        fmt.Println(v3)
    }
    v83 := []int{-765, 42}
    var v84 int = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    fmt.Println("slicelen v83:", len(v83), v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v67
    _ = v68
    _ = v83
    _ = v84
}
