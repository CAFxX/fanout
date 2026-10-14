package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint16; F1 int32; F2 int16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16 }
func genFn0(p0 uint64, p1 uint32) uint32 {
    if (1 != 100) {
        switch -1 {
        case 2147483647:
            if true {
                if (!(-1 < 32767)) {
                    switch 9223372036854775807 {
                    case -753:
                        switch "hello" {
                        case "a":
                            if ((p1 - (p1 + (p1 * p1))) != (p1 ^ p1)) {
                                if true {
                                    v6 := GenS1{F0: 0}
                                    _ = v6
                                }
                            }
                        case "hello":
                            var v7 int32 = -2147483648
                            _ = v7
                        }
                    }
                } else {
                    for fi8 := 0; fi8 < 14; fi8++ {
                        var v9 int64 = 42
                        _ = fi8
                        _ = v9
                    }
                    if false {
                        var v10 string = "hello"
                        v11 := []int{0}
                        _ = v10
                        _ = v11
                    } else {
                        var v12 int16 = 408
                        _ = v12
                    }
                    fmt.Println(p1)
                }
                fmt.Println(p0, p1)
            } else {
                var v13 int = 127
                var v14 uint8 = 127
                _ = v13
                _ = v14
            }
        default:
            var v15 uint = 2
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 int = fi16
            p1 -= 42
            _ = fi16
            _ = v17
        }
    } else {
        v18 := GenS1{F0: 1}
        _ = v18
    }
    fmt.Println(p0, p1)
    v19 := []int{int(p1), 128, 1, 32767}
    v20 := GenS2{F0: true, F1: 360}
    switch "a" {
    case "foo bar":
        for fi21 := 0; fi21 < 2; fi21++ {
            for fi22 := 0; fi22 < 9; fi22++ {
                var v23 uint32 = (p1 & p1)
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                v19 = append(v19, fi22)
                p0 -= 4294967295
                v23 = p1
                _ = fi22
                _ = v23
            }
            switch (fi21 - fi21) {
            case 127:
                if false {
                    v19 = append(v19, (fi21 * (fi21 * fi21)))
                    v19 = append(v19, 9223372036854775807)
                }
            case -896:
                fmt.Println(fi21, p1)
            }
            if false {
                _ = len(v19)
            }
            _ = fi21
        }
    case "\n\t":
        v24 := GenS1{F0: 2}
        _ = v24
    default:
        var v25 uint16 = 788
        _ = v25
    }
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 388
}
func genFn1(p0 string, p1 uint, p2 int64) uint {
    switch ("a" + (("\n\t" + (("" + "a") + (p0 + ("" + p0)))) + p0)) {
    case "foo bar":
        var v26 uint64 = 1
        _ = v26
    case "":
        var v27 int8 = 1
        _ = v27
    case "\n\t":
        v28 := GenS2{F0: (!(-103 != -72)), F1: 937}
        _ = v28
    default:
        switch 127 {
        case -224:
            v29 := []int{-2147483648, 0, 559, 269}
            _ = v29
        }
    }
    var v30 int8 = -128
    switch (p0 + "a") {
    case "":
        switch "\n\t" {
        case "a":
            v31 := GenS2{F0: false, F1: 42}
            _ = v31
        case "":
            v30 = (v30 | v30)
        case "hello":
            for fi32 := 0; fi32 < 19; fi32++ {
                p1 = 18446744073709551615
                for fi33 := 0; fi33 < 9; fi33++ {
                    var v34 float64 = 0.0
                    p0 = p0
                    _ = fi33
                    _ = v34
                }
                if (genFn0(239, 208) != 42) {
                    var v35 uint = p1
                    _ = v35
                } else {
                    var v36 uint32 = 42
                    _ = v36
                }
                _ = fi32
            }
        }
    }
    var v37 int8 = v30
    for fi38 := 0; fi38 < 10; fi38++ {
        v37 -= (v30 | v30)
        _ = fi38
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    _ = v37
    return (p1 | 295)
}

func main() {
    var v1 int32 = -54
    var v2 int = -2147483648
    var v3 uint64 = 718
    var v4 uint16 = 0
    v5 := make(map[int]string)
    for fi39 := 0; fi39 < 14; fi39++ {
        for fi40 := 0; fi40 < 12; fi40++ {
            v41 := GenS2{F0: true, F1: (v4 >> ((uint(v3)) & 63))}
            var v42 int8 = -1
            v1 = (v1 ^ (v1 & 816))
            _ = fi40
            _ = v41
            _ = v42
        }
        if true {
            for fi43 := 0; fi43 < 15; fi43++ {
                if ((v4 & 1) < (v4 + (^(v4)))) {
                    v3 = (v3 & (v3 / ((v3) | 1)))
                }
                if true {
                    v1 += -1
                    v2 += (fi39 / ((300) | 1))
                }
                _ = fi43
            }
        }
        _ = fi39
    }
    for fi44 := 0; fi44 < 1; fi44++ {
        v3 = (v3 + (^(v3)))
        _ = fi44
    }
    var v45 uint = 664
    var v46 uint32 = 0
    if (-2.25 != -2.25) {
        fmt.Println(v1)
        var v47 string = "\n\t"
        _ = v47
    } else {
        for fi48 := 0; fi48 < 18; fi48++ {
            for fi49 := 0; fi49 < 14; fi49++ {
                v50 := GenS3{F0: -502, F1: (v46 / ((genFn0((v3 | 2), 65535)) | 1)), F2: 255}
                _ = fi49
                _ = v50
            }
            if (0.0 >= -2.25) {
                switch ("foo bar" + "\n\t") {
                case "foo bar":
                    for fi51 := 0; fi51 < 15; fi51++ {
                        delete(v5, (fi48 / ((fi48) | 1)))
                        v2 -= fi51
                        var v52 int16 = 1
                        _ = fi51
                        _ = v52
                    }
                case "hello":
                    if false {
                        v5[(fi48 ^ (-(v2)))] = "\n\t"
                        v53 := GenS2{F0: true, F1: v4}
                        if false {
                            v46 -= v46
                            var v54 int64 = -9223372036854775808
                            for fi55 := 0; fi55 < 8; fi55++ {
                                delete(v5, -745)
                                fmt.Println(v45, v1, fi55)
                                v56 := GenS4{F0: genFn1("foo bar", v45, -9223372036854775808)}
                                _ = fi55
                                _ = v56
                            }
                            _ = v54
                        }
                        _ = v53
                    }
                case "a":
                    fmt.Println(v46, fi48, v3)
                }
                fmt.Println(v46, v46)
            } else {
                switch ("a" + ("" + "a")) {
                case "foo bar":
                    var v57 int16 = int16(v46)
                    _ = v57
                case "":
                    for fi58 := 0; fi58 < 3; fi58++ {
                        if true {
                            if true {
                                v59 := GenS5{F0: "hello", F1: (v1 + v1), F2: v46}
                                var v60 uint32 = 2147483648
                                v2 = 31
                                _ = v59
                                _ = v60
                            }
                        } else {
                            if true {
                                var v61 uint16 = v4
                                _ = v61
                            } else {
                                v62 := GenS6{F0: (v4 << ((v45) & 63)), F1: 0, F2: 0}
                                _ = v62
                            }
                        }
                        _ = fi58
                    }
                case "\n\t":
                    v63 := GenS6{F0: (v4 >> (((v45 << (((v45 * v45)) & 63))) & 63)), F1: v1, F2: 0}
                    _ = v63
                default:
                    v64 := []int{}
                    _ = v64
                }
                v1 = (^(v1))
                v1 -= v1
            }
            _ = fi48
        }
        v65 := GenS7{F0: 0}
        _ = v65
    }
    for fi66 := 0; fi66 < 1; fi66++ {
        switch -9223372036854775808 {
        case 9223372036854775807:
            if true {
                switch fi66 {
                case -2147483648:
                    v5[9223372036854775807] = ""
                case 43:
                    fmt.Println(fi66, v2)
                }
            }
        case -660:
            v67 := []int{127}
            _ = v67
        case -19:
            switch ("hello" + ((("" + ("hello" + ("a" + ""))) + "foo bar") + "")) {
            case "a":
                var v68 int = fi66
                _ = v68
            case "hello":
                for fi69 := 0; fi69 < 18; fi69++ {
                    v45 = v45
                    switch ("" + "a") {
                    case "hello":
                        v46 = 4294967295
                    case "\n\t":
                        switch (fi69 % (((fi69 | (fi66 % ((-9223372036854775808) | 1)))) | 1)) {
                        case -253:
                            fmt.Println(fi66, fi66, v4)
                        default:
                            if ((v45 - (v45 - 183)) > 927) {
                                switch int(v4) {
                                case -2147483648:
                                    v1 -= (v1 ^ -632)
                                case -1:
                                    v70 := []int{int(v4)}
                                    _ = v70
                                }
                                v71 := []int{}
                                v71 = append(v71, 4294967295)
                                var v72 uint32 = (v46 % ((v46) | 1))
                                _ = v71
                                _ = v72
                            } else {
                                v73 := GenS7{F0: 0}
                                _ = v73
                            }
                        }
                    case "a":
                        switch "\n\t" {
                        case "a":
                            v74 := []int{fi66}
                            _ = v74
                        case "":
                            switch "hello" {
                            case "":
                                v3 -= (v3 << (((v45 << ((v45) & 63))) & 63))
                            case "foo bar":
                                v75 := GenS8{F0: 42}
                                _ = v75
                            }
                        default:
                            fmt.Println(v46, v3)
                        }
                    default:
                        v76 := GenS2{F0: true, F1: v4}
                        _ = v76
                    }
                    if false {
                        var v77 float64 = 1e10
                        v78 := []int{(v2 & fi69)}
                        _ = v77
                        _ = v78
                    }
                    v79 := GenS3{F0: 37, F1: genFn0(18446744073709551615, (v46 >> (((v45 | v45)) & 63))), F2: 128}
                    _ = fi69
                    _ = v79
                }
            }
        default:
            fmt.Println(v3)
        }
        _ = fi66
    }
    var v80 int = -109
    for fi81 := 0; fi81 < 12; fi81++ {
        v82 := GenS7{F0: 9223372036854775807}
        v83 := v5[128]
        v84 := GenS1{F0: 255}
        _ = fi81
        _ = v82
        _ = v83
        _ = v84
    }
    for fi85 := 0; fi85 < 12; fi85++ {
        v86 := GenS3{F0: 1000, F1: v46, F2: 206}
        for fi87 := 0; fi87 < 20; fi87++ {
            v88 := GenS5{F0: "a", F1: v1, F2: 2147483648}
            _ = fi87
            _ = v88
        }
        v80 -= (v80 / ((959) | 1))
        _ = fi85
        _ = v86
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v80
}
