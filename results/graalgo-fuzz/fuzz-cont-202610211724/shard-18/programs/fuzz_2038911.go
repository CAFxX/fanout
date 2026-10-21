package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint; F1 int16 }
type GenS5 struct { F0 uint32; F1 int32 }
type GenS6 struct { F0 int32; F1 string; F2 uint64 }
type GenS7 struct { F0 uint16; F1 string }
func genFn0() string {
    if (42 <= 249) {
        v4 := GenS1{F0: -135, F1: 2, F2: 0}
        var v5 bool = (759 != 42)
        for fi6 := 0; fi6 < 4; fi6++ {
            for fi7 := 0; fi7 < 18; fi7++ {
                if (!(42 != 42)) {
                    v8 := GenS2{F0: 42}
                    v9 := []int{32768, fi7, (fi6 + 383), (fi6 << ((370) & 63))}
                    var v10 uint8 = 42
                    _ = v10
                    _ = v8
                    _ = v9
                } else {
                    v5 = true
                }
                fmt.Println(v5, v5, v5)
                _ = fi7
            }
            v5 = v5
            _ = fi6
        }
        _ = v4
        _ = v5
    }
    if true {
        v11 := []int{-103, -2147483648, -1, 947}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int32 = -223
            switch "foo bar" {
            case "a":
                if (297 == 745) {
                    var v14 int = (fi12 - fi12)
                    var v15 int32 = (v13 | v13)
                    fmt.Println(v15, v15)
                    _ = v14
                    _ = v15
                } else {
                    fmt.Println(fi12)
                }
            case "\n\t":
                v13 = (v13 * (v13 ^ v13))
            }
            v13 -= (v13 % ((36) | 1))
            _ = fi12
            _ = v13
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            switch "foo bar" {
            case "":
                v17 := []int{(^(fi16))}
                _ = v17
            }
            fmt.Println(fi16)
            _ = fi16
        }
    }
    if (false && true) {
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS1{F0: 42, F1: 694, F2: -128}
            var v20 uint8 = 130
            var v21 int16 = -533
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        if (42 >= 49) {
            var v22 string = "a"
            _ = v22
        }
        var v23 int16 = -32768
        _ = v23
    }
    var v24 int8 = 100
    switch ("" + "\n\t") {
    case "":
        v24 += v24
    case "\n\t":
        for fi25 := 0; fi25 < 6; fi25++ {
            switch fi25 {
            case 2147483647:
                var v26 int16 = 641
                _ = v26
            case 1:
                if (-1 != 771) {
                    switch fi25 {
                    case -556:
                        var v27 float32 = 0.1
                        _ = v27
                    }
                    if (850 <= 1000) {
                        v24 = v24
                        v28 := GenS3{F0: (!true), F1: (fi25 << ((0) & 63))}
                        _ = v28
                    } else {
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v24)
                        case "":
                            var v29 uint = 2
                            _ = v29
                        }
                        v24 -= 127
                        v30 := GenS4{F0: 213, F1: -32768}
                        _ = v30
                    }
                }
            case 42:
                var v31 string = "a"
                _ = v31
            default:
                v24 += v24
            }
            if false {
                if false {
                    v32 := GenS4{F0: 4294967295, F1: 1000}
                    _ = v32
                } else {
                    var v33 int16 = 1000
                    var v34 int16 = -1
                    _ = v33
                    _ = v34
                }
                var v35 int64 = 1
                _ = v35
            } else {
                v36 := GenS2{F0: 527}
                v24 = -128
                v37 := GenS5{F0: 871, F1: -2147483648}
                _ = v36
                _ = v37
            }
            for fi38 := 0; fi38 < 19; fi38++ {
                var v39 int32 = 2147483647
                var v40 uint8 = 7
                fmt.Println(v39, fi38, v24)
                fmt.Println(v39, v39, v24)
                _ = fi38
                _ = v39
                _ = v40
            }
            _ = fi25
        }
    case "a":
        v41 := []int{9223372036854775807}
        _ = v41
    }
    _ = v24
    return "\n\t"
}

func main() {
    var v1 float64 = 0.0
    var v2 bool = (((0.1 * (-2.25 - 0.0)) / (-2.25 / 0.1)) == (0.0 / (0.1 + (3.14159 / 0.1))))
    var v3 int16 = 1000
    v2 = (873 >= 18446744073709551615)
    v42 := []int{}
    for fi43 := 0; fi43 < 8; fi43++ {
        var v44 uint64 = 2
        fmt.Println(v3, v1, v2)
        for fi45 := 0; fi45 < 17; fi45++ {
            _ = len(v42)
            switch (fi43 + (fi45 >> ((714) & 63))) {
            case 220:
                v44 += v44
            case 127:
                v46 := GenS3{F0: false, F1: fi43}
                _ = v46
            default:
                v44 = 1
                fmt.Println(v3, v1, v1)
            }
            v47 := GenS6{F0: -1, F1: "hello", F2: (v44 ^ (^(v44)))}
            var v48 uint = 4294967295
            _ = fi45
            _ = v47
            _ = v48
        }
        _ = fi43
        _ = v44
    }
    if v2 {
        if len(v42) > 0 {
            v42[0] = 0
        }
    } else {
        if (((v2 && v2) && false) && (804 < 529)) {
            for fi49 := 0; fi49 < 13; fi49++ {
                switch fi49 {
                case 2147483647:
                    switch "" {
                    case "hello":
                        v1 -= (0.0 * 1.5)
                    case "\n\t":
                        for fi50 := 0; fi50 < 10; fi50++ {
                            v3 = v3
                            switch fi49 {
                            case 4294967295:
                                v1 -= float64(float32(v1))
                            case -9223372036854775808:
                                fmt.Println(fi50)
                            default:
                                v42 = append(v42, 2147483647)
                            }
                            _ = fi50
                        }
                    }
                case 0:
                    var v51 uint64 = 1
                    _ = v51
                case -2147483648:
                    var v52 int8 = 0
                    _ = v52
                }
                if (((!(-2147483648 != fi49)) || v2) || v2) {
                    for fi53 := 0; fi53 < 15; fi53++ {
                        var v54 int8 = 0
                        fmt.Println(v54)
                        _ = fi53
                        _ = v54
                    }
                    if len(v42) >= 2 {
                        v42 = v42[1:len(v42)]
                    }
                    switch ("\n\t" + "foo bar") {
                    case "a":
                        if len(v42) >= 2 {
                            v42 = v42[1:len(v42)]
                        }
                    case "":
                        var v55 int64 = -380
                        _ = v55
                    }
                }
                _ = fi49
            }
            var v56 uint64 = 833
            var v57 uint = 321
            _ = v56
            _ = v57
        } else {
            switch 127 {
            case 574:
                var v58 uint8 = 127
                _ = v58
            }
        }
        var v59 uint = 0
        fmt.Println(v59, v2)
        _ = v59
    }
    switch 4294967295 {
    case -9223372036854775808:
        v2 = v2
    default:
        var v60 int32 = 1
        _ = v60
    }
    var v61 uint8 = 0
    if false {
        v3 -= 967
        switch -1 {
        case 32767:
            if false {
                v62 := GenS4{F0: 4294967295, F1: (v3 << ((18446744073709551615) & 63))}
                _ = len(v42)
                _ = v62
            }
        default:
            switch "\n\t" {
            case "\n\t":
                for fi63 := 0; fi63 < 8; fi63++ {
                    v1 = 1.5
                    _ = fi63
                }
            case "":
                var v64 int32 = -97
                _ = v64
            case "a":
                for fi65 := 0; fi65 < 6; fi65++ {
                    if ((0 >= 0) && true) {
                        if (true || (!(9223372036854775807 > -210))) {
                            var v66 int = (fi65 % (((fi65 - (fi65 & 42))) | 1))
                            v61 = 2
                            switch "foo bar" {
                            case "foo bar":
                                var v67 uint8 = 1
                                _ = v67
                            default:
                                if (1 <= -58) {
                                    switch "hello" {
                                    case "\n\t":
                                        if len(v42) >= 2 {
                                            v42 = v42[1:len(v42)]
                                        }
                                    case "a":
                                        if (!(!(!(9223372036854775807 > fi65)))) {
                                            var v68 uint64 = 0
                                            _ = v68
                                        } else {
                                            v42 = append(v42, (fi65 << ((2) & 63)))
                                        }
                                    default:
                                        var v69 int16 = v3
                                        _ = v69
                                    }
                                    v70 := GenS1{F0: 9223372036854775807, F1: uint16(v61), F2: 0}
                                    _ = len(v42)
                                    _ = v70
                                } else {
                                    v1 = v1
                                }
                            }
                            _ = v66
                        }
                    }
                    _ = fi65
                }
            default:
                if ((1.5 <= 0.0) && true) {
                    if len(v42) > 0 {
                        v42[0] = 1
                    }
                    if true {
                        v2 = true
                        if len(v42) > 0 {
                            v42[0] = 32767
                        }
                        if len(v42) > 0 {
                            v42[0] = 2147483647
                        }
                    }
                    v71 := GenS7{F0: 2, F1: ("a" + "foo bar")}
                    _ = v71
                }
            }
        }
    }
    if v2 {
        fmt.Println(v3)
        v61 = v61
        fmt.Println(v3)
    } else {
        v61 = v61
        v72 := GenS3{F0: v2, F1: 32768}
        fmt.Println(v1, v3, v2)
        _ = v72
    }
    fmt.Println(v61, v3, v3)
    var v73 float32 = (float32(v1) / (0.0 + float32(v1)))
    v74 := GenS7{F0: 255, F1: "foo bar"}
    v75 := GenS5{F0: 2, F1: 2147483647}
    var v76 int32 = 533
    v42 = append(v42, 239)
    var v77 int32 = -1
    if v2 {
        var v78 uint64 = 0
        v61 += 13
        if len(v42) >= 2 {
            v42 = v42[1:len(v42)]
        }
        _ = v78
    } else {
        switch -9223372036854775808 {
        case 2147483647:
            v2 = v2
        case 128:
            if (v2 || (false && ((v2 && v2) && true))) {
                switch 204 {
                case 4294967295:
                    switch -61 {
                    case -72:
                        if (false || false) {
                            v73 += v73
                        }
                    case 128:
                        fmt.Println(v3, v2, v76)
                    case 32768:
                        switch 936 {
                        case -9223372036854775808:
                            var v79 int32 = (v76 + v76)
                            _ = v79
                        }
                    default:
                        v80 := GenS7{F0: 1, F1: "foo bar"}
                        _ = v80
                    }
                case 0:
                    v1 = (float64((v73 * v73)) - (v1 + float64((float32(v1) + (v73 * (v73 + (float32((v1 + float64(3.14159))) / ((1.5 - ((v73 / (v73 / v73)) + v73)) * (v73 * v73)))))))))
                case 383:
                    var v81 int8 = 127
                    _ = v81
                default:
                    fmt.Println(v77, v2)
                }
                var v82 float64 = v1
                var v83 string = "hello"
                _ = v82
                _ = v83
            }
        case 707:
            var v84 uint32 = 4294967295
            _ = v84
        }
        for fi85 := 0; fi85 < 14; fi85++ {
            fmt.Println(v3)
            var v86 int64 = 905
            v87 := GenS6{F0: (v77 * v77), F1: ("hello" + ("hello" + "foo bar")), F2: 42}
            _ = fi85
            _ = v86
            _ = v87
        }
        v42 = append(v42, -9223372036854775808)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v61:", v61)
    fmt.Println("v73:", v73)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v61
    _ = v73
    _ = v74
    _ = v75
    _ = v76
    _ = v77
}
