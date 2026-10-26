package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS3 struct { F0 int32; F1 uint; F2 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint; F1 uint }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int16; F1 string }

func main() {
    var v1 uint16 = 42
    var v2 string = "foo bar"
    var v3 uint32 = 2147483648
    v4 := []int{9223372036854775807, -9223372036854775808, 32767}
    v5 := []int{-882}
    if (-32768 < -14) {
        fmt.Println(v3)
        v1 = v1
        v6 := GenS1{F0: 846, F1: 1, F2: (632 >= 2147483647)}
        _ = v6
    }
    if (!(3.14159 <= float64(-2.25))) {
        for fi7 := 0; fi7 < 7; fi7++ {
            for fi8 := 0; fi8 < 6; fi8++ {
                v9 := GenS1{F0: 0, F1: 42, F2: ((!false) && (!true))}
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    }
    v1 -= (v1 & 606)
    if false {
        if true {
            if true {
                var v10 int16 = 517
                switch 290 {
                case 9223372036854775807:
                    fmt.Println(v1, v1)
                case -2147483648:
                    v11 := GenS2{F0: 128, F1: -343, F2: 70}
                    _ = v11
                case 4294967295:
                    if ((v10 ^ int16(v1)) == -32768) {
                        v12 := GenS1{F0: 774, F1: 69, F2: (0 >= 2147483647)}
                        var v13 int64 = 291
                        if false {
                            if (!(false || (int16(v3) <= v10))) {
                                for fi14 := 0; fi14 < 5; fi14++ {
                                    v15 := GenS1{F0: 0, F1: 42, F2: (!true)}
                                    v2 = v2
                                    _ = fi14
                                    _ = v15
                                }
                                v4 = append(v4, 858)
                                _ = len(v4)
                            } else {
                                if ((false || true) && true) {
                                    if (v3 <= v3) {
                                        _ = len(v5)
                                        var v16 uint32 = (v3 + (v3 & (v3 / ((42) | 1))))
                                        _ = v16
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = int(v13)
                                        }
                                        var v17 int32 = 824
                                        _ = v17
                                    }
                                } else {
                                    var v18 float64 = -2.25
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = v18
                                }
                                if (-1 >= 108) {
                                    if true {
                                        v19 := GenS1{F0: 347, F1: 0, F2: true}
                                        _ = v19
                                    }
                                    v5 = append(v5, 32767)
                                    for fi20 := 0; fi20 < 11; fi20++ {
                                        _ = len(v5)
                                        v2 = (v2 + v2)
                                        v3 = v3
                                        _ = fi20
                                    }
                                }
                            }
                            switch -831 {
                            case 650:
                                for fi21 := 0; fi21 < 7; fi21++ {
                                    for fi22 := 0; fi22 < 9; fi22++ {
                                        if false {
                                            fmt.Println(v2)
                                        } else {
                                            var v23 int16 = 1
                                            v24 := GenS3{F0: 2147483647, F1: 0, F2: 59}
                                            var v25 int32 = 833
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        v3 += 2147483648
                                        _ = fi22
                                    }
                                    _ = fi21
                                }
                            case 42:
                                v1 = 0
                            default:
                                _ = len(v4)
                            }
                        }
                        _ = v12
                        _ = v13
                    }
                default:
                    if true {
                        for fi26 := 0; fi26 < 16; fi26++ {
                            switch (fi26 & fi26) {
                            case -9223372036854775808:
                                v5 = append(v5, (-(fi26)))
                            default:
                                fmt.Println(fi26, v3)
                            }
                            if true {
                                var v27 bool = (0 < 348)
                                fmt.Println(v27, v27, v27)
                                _ = v27
                            }
                            var v28 int8 = 13
                            _ = fi26
                            _ = v28
                        }
                        if true {
                            if (!true) {
                                v1 = (v1 | (v1 / ((74) | 1)))
                                v3 -= v3
                                _ = len(v5)
                            }
                        } else {
                            v1 += v1
                        }
                    }
                }
                v29 := GenS4{F0: -777}
                _ = v10
                _ = v29
            } else {
                v4 = append(v4, -788)
            }
            var v30 int = 128
            if true {
                v3 += v3
                v1 += (v1 >> ((42) & 63))
            }
            _ = v30
        }
        v5 = append(v5, -633)
        v31 := GenS5{F0: 9223372036854775807}
        _ = v31
    }
    if (-825 == 892) {
        if (((v2 == v2) || (65535 <= 42)) && ((1.5 + 1.5) == 1.5)) {
            for fi32 := 0; fi32 < 5; fi32++ {
                v33 := GenS6{F0: true, F1: 4294967295, F2: uint32(v1)}
                for fi34 := 0; fi34 < 4; fi34++ {
                    v2 = (((v2 + (("\n\t" + "hello") + v2)) + (((v2 + "a") + (v2 + v2)) + (v2 + v2))) + v2)
                    if (200 >= 1000) {
                        v35 := GenS7{F0: ((!(0 > 69)) && (236 == -32768)), F1: 1}
                        v1 -= (^(v1))
                        if len(v5) > 0 {
                            v5[0] = fi34
                        }
                        fmt.Println(fi32, v2)
                        _ = v35
                    } else {
                        if false {
                            var v36 uint32 = 2147483647
                            _ = v36
                        }
                        var v37 int32 = 2147483647
                        _ = v37
                    }
                    _ = fi34
                }
                var v38 string = "\n\t"
                _ = fi32
                _ = v33
                _ = v38
            }
            switch 540 {
            case 946:
                for fi39 := 0; fi39 < 8; fi39++ {
                    v1 -= 65535
                    _ = fi39
                }
            default:
                var v40 uint16 = v1
                _ = v40
            }
        }
        if false {
            var v41 uint8 = 19
            _ = v41
        }
        for fi42 := 0; fi42 < 14; fi42++ {
            switch (fi42 & 762) {
            case 946:
                _ = len(v4)
            case 2147483647:
                v2 = (v2 + v2)
            }
            var v43 float32 = float32(float64((1e10 / 1.5)))
            _ = fi42
            _ = v43
        }
    } else {
        switch 507 {
        case 1:
            var v44 int = 4294967295
            _ = v44
        case 127:
            var v45 int64 = 319
            _ = v45
        default:
            fmt.Println(v1, v1, v2)
        }
        var v46 uint16 = v1
        for fi47 := 0; fi47 < 11; fi47++ {
            switch v2 {
            case "hello":
                v48 := GenS5{F0: (fi47 + -273)}
                _ = v48
            case "foo bar":
                var v49 uint = 0
                _ = v49
            case "":
                for fi50 := 0; fi50 < 1; fi50++ {
                    if len(v5) > 0 {
                        v5[0] = (-(fi50))
                    }
                    if len(v5) > 0 {
                        v5[0] = 9223372036854775807
                    }
                    var v51 uint16 = v1
                    if false {
                        var v52 int32 = int32(fi50)
                        if false {
                            switch v2 {
                            case "\n\t":
                                switch "\n\t" {
                                case "hello":
                                    v53 := GenS8{F0: 1}
                                    _ = v53
                                case "foo bar":
                                    v52 = (-(v52))
                                }
                            case "foo bar":
                                if ((866 < 1) || false) {
                                    if (255 > 52) {
                                        v54 := GenS9{F0: 18446744073709551615, F1: 1}
                                        _ = v54
                                    }
                                    v52 = (v52 % ((1) | 1))
                                } else {
                                    v55 := GenS4{F0: fi50}
                                    fmt.Println(v51, fi47)
                                    _ = v55
                                }
                            case "":
                                v52 = (v52 / (((v52 << ((0) & 63))) | 1))
                            default:
                                if (-1 < 9223372036854775807) {
                                    v3 += 2147483648
                                    var v56 int16 = -1
                                    switch ("hello" + "\n\t") {
                                    case "\n\t":
                                        if len(v5) > 0 {
                                            v5[0] = fi50
                                        }
                                    case "":
                                        switch "a" {
                                        case "hello":
                                            if false {
                                                if false {
                                                    v57 := GenS3{F0: v52, F1: uint(v1), F2: 1}
                                                    v1 = 42
                                                    v4 = append(v4, fi50)
                                                    if len(v4) > 0 {
                                                        v4[0] = (-(fi50))
                                                    }
                                                    _ = v57
                                                }
                                                v1 -= 0
                                                var v58 int8 = 0
                                                _ = v58
                                            } else {
                                                if (!true) {
                                                    v1 += v46
                                                    v51 += v51
                                                } else {
                                                    _ = len(v5)
                                                    if (!false) {
                                                        fmt.Println(v52, v52, fi47)
                                                        if true {
                                                            _ = len(v4)
                                                        } else {
                                                            v3 = 2
                                                        }
                                                        v59 := GenS8{F0: (fi50 + (fi47 | (fi47 ^ fi47)))}
                                                        _ = v59
                                                    } else {
                                                        var v60 bool = (42 == 62)
                                                        _ = v60
                                                    }
                                                    v61 := GenS7{F0: false, F1: (v3 >> ((527) & 63))}
                                                    _ = v61
                                                }
                                                _ = len(v4)
                                                v62 := GenS9{F0: 0, F1: uint(v3)}
                                                _ = len(v4)
                                                var v63 int32 = (v52 % ((v52) | 1))
                                                _ = v62
                                                _ = v63
                                            }
                                        }
                                    case "foo bar":
                                        v1 = v1
                                    }
                                    _ = v56
                                }
                            }
                        } else {
                            v2 = ""
                        }
                        if (false && (869 < uint64(v51))) {
                            var v64 uint32 = v3
                            _ = v64
                        }
                        v52 += v52
                        _ = v52
                    }
                    v2 = v2
                    _ = fi50
                    _ = v51
                }
            default:
                v1 += (v46 % (((^(v1))) | 1))
            }
            v65 := GenS1{F0: 0, F1: 127, F2: ((0.0 - 3.14159) <= -2.25)}
            if true {
                for fi66 := 0; fi66 < 14; fi66++ {
                    if false {
                        var v67 string = "foo bar"
                        var v68 int64 = 0
                        _ = v67
                        _ = v68
                    }
                    var v69 int64 = 1
                    fmt.Println(v2, v3, v2)
                    v3 = uint32(v1)
                    _ = fi66
                    _ = v69
                }
                var v70 int = (fi47 - 9223372036854775807)
                v3 = v3
                _ = v70
            }
            _ = fi47
            _ = v65
        }
        _ = v46
    }
    switch 1 {
    case 42:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 128:
        for fi71 := 0; fi71 < 14; fi71++ {
            fmt.Println(v2)
            _ = fi71
        }
    case -2147483648:
        var v72 float32 = 0.1
        _ = v72
    default:
        if false {
            _ = len(v5)
        } else {
            if (0 <= 1) {
                fmt.Println(v2, v3, v1)
                for fi73 := 0; fi73 < 13; fi73++ {
                    var v74 bool = true
                    _ = fi73
                    _ = v74
                }
            }
            var v75 int16 = 1000
            v76 := GenS1{F0: 232, F1: 0, F2: false}
            _ = v75
            _ = v76
        }
    }
    if len(v4) > 0 {
        v4[0] = 1
    }
    v77 := GenS10{F0: 514}
    if true {
        for fi78 := 0; fi78 < 10; fi78++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi78
        }
        if (false && true) {
            v79 := GenS2{F0: 128, F1: 30, F2: v1}
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v79
        }
    }
    v80 := GenS11{F0: 65535, F1: 0}
    var v81 uint16 = 1
    if (!(!((42 <= 1) || false))) {
        v81 += (^(v1))
    }
    for fi82 := 0; fi82 < 15; fi82++ {
        fmt.Println(v1)
        for fi83 := 0; fi83 < 19; fi83++ {
            var v84 int16 = 1
            _ = fi83
            _ = v84
        }
        _ = fi82
    }
    if true {
        v5 = append(v5, -151)
        switch v2 {
        case "\n\t":
            fmt.Println(v81, v1)
        }
    }
    if (1e10 != 1e10) {
        for fi85 := 0; fi85 < 16; fi85++ {
            if ((42 >= 127) || true) {
                var v86 bool = (false && true)
                fmt.Println(fi85, v3)
                v1 -= 254
                _ = v86
            } else {
                if true {
                    v1 -= v81
                    v81 = (v1 % ((2) | 1))
                    for fi87 := 0; fi87 < 19; fi87++ {
                        fmt.Println(v81)
                        _ = fi87
                    }
                }
            }
            v2 = (v2 + v2)
            v2 = v2
            _ = fi85
        }
        switch ((v2 + (v2 + "foo bar")) + (v2 + "")) {
        case "hello":
            v88 := GenS11{F0: 4294967295, F1: 0}
            _ = v88
        case "":
            v89 := GenS12{F0: 2}
            _ = v89
        case "a":
            var v90 string = v2
            _ = v90
        default:
            var v91 int32 = -538
            _ = v91
        }
        fmt.Println(v1)
    } else {
        if (false && false) {
            for fi92 := 0; fi92 < 19; fi92++ {
                switch int(v1) {
                case 220:
                    v2 = v2
                }
                _ = fi92
            }
            v3 = (v3 / ((v3) | 1))
        } else {
            v93 := GenS12{F0: 1}
            var v94 uint32 = 120
            _ = v93
            _ = v94
        }
        for fi95 := 0; fi95 < 20; fi95++ {
            for fi96 := 0; fi96 < 13; fi96++ {
                v97 := GenS13{F0: 42, F1: v2}
                v81 = v1
                if true {
                    var v98 int = 651
                    fmt.Println(v81, v1)
                    var v99 float64 = 3.14159
                    _ = v98
                    _ = v99
                }
                _ = fi96
                _ = v97
            }
            var v100 uint8 = 0
            _ = fi95
            _ = v100
        }
        v1 = (v1 % ((270) | 1))
    }
    for fi101 := 0; fi101 < 9; fi101++ {
        v3 -= (v3 | (v3 | v3))
        fmt.Println(v2)
        v5 = append(v5, (^(fi101)))
        _ = fi101
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v77
    _ = v80
    _ = v81
}
