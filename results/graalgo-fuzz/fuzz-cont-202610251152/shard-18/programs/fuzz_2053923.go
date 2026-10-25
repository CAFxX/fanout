package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int8; F1 uint64; F2 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 string; F1 int64; F2 uint64 }
func genFn0(p0 int) bool {
    var v4 bool = true
    var v5 bool = v4
    switch p0 {
    case -1:
        if (v4 || v5) {
            if (-534 != -9223372036854775808) {
                var v6 int16 = 1000
                v7 := GenS1{F0: -2147483648, F1: 18446744073709551615}
                for fi8 := 0; fi8 < 14; fi8++ {
                    v9 := GenS2{F0: 255, F1: (^(v6))}
                    _ = fi8
                    _ = v9
                }
                _ = v6
                _ = v7
            } else {
                v10 := []int{p0, (p0 - (p0 % ((-845) | 1))), (-(p0)), p0}
                for fi11 := 0; fi11 < 17; fi11++ {
                    var v12 int32 = -2147483648
                    var v13 int64 = -655
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = v10
            }
        }
    case 1:
        var v14 int32 = -1
        _ = v14
    }
    var v15 bool = true
    _ = p0
    _ = v15
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 uint, p1 int64) uint16 {
    if (((p1 + (p1 << ((65535) & 63))) >= p1) && (!true)) {
        var v16 float64 = 3.14159
        var v17 uint16 = 1
        p1 = 42
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(fi18, p1)
            var v19 float64 = 1.5
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(p1)
    switch 9223372036854775807 {
    case 128:
        v20 := GenS2{F0: 424, F1: 688}
        _ = v20
    case 127:
        v21 := []int{232, 32767, 128, 348}
        _ = v21
    default:
        if false {
            fmt.Println(p0, p0)
            if true {
                p0 += (^(p0))
            } else {
                if true {
                    v22 := []int{42}
                    if (((!true) && true) || false) {
                        if false {
                            if (42 < 39) {
                                fmt.Println(p1, p1)
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            } else {
                                if true {
                                    var v23 uint32 = 472
                                    fmt.Println(p1, v23)
                                    v24 := GenS1{F0: -2147483648, F1: 42}
                                    _ = v23
                                    _ = v24
                                }
                            }
                            p1 -= p1
                        } else {
                            p1 = 589
                        }
                        for fi25 := 0; fi25 < 20; fi25++ {
                            if len(v22) > 0 {
                                v22[0] = 2147483647
                            }
                            var v26 int8 = 1
                            _ = fi25
                            _ = v26
                        }
                    } else {
                        p0 = (p0 - (p0 | 0))
                        var v27 uint32 = 2147483647
                        var v28 int16 = 32767
                        _ = v27
                        _ = v28
                    }
                    _ = v22
                } else {
                    fmt.Println(p1, p1)
                }
                p1 = -76
            }
            for fi29 := 0; fi29 < 19; fi29++ {
                if true {
                    for fi30 := 0; fi30 < 19; fi30++ {
                        p1 -= 1
                        switch -9223372036854775808 {
                        case 267:
                            var v31 uint32 = 0
                            _ = v31
                        }
                        var v32 bool = (42 <= 32767)
                        _ = fi30
                        _ = v32
                    }
                }
                _ = fi29
            }
        }
    }
    _ = p0
    _ = p1
    return 32767
}

func main() {
    var v1 string = "hello"
    v2 := []int{42, -1, 32767, -213}
    var v3 int8 = 127
    for fi33 := 0; fi33 < 9; fi33++ {
        fmt.Println(fi33, fi33)
        _ = fi33
    }
    var v34 uint32 = 753
    if false {
        v35 := GenS1{F0: 0, F1: 4294967295}
        var v36 int32 = 0
        _ = v35
        _ = v36
    } else {
        v37 := GenS3{F0: 2, F1: genFn1(uint(v34), 1), F2: (^(v34))}
        var v38 uint8 = 128
        _ = v37
        _ = v38
    }
    if (((!(v3 < 42)) && (-2.25 == ((-2.25 / 1.5) - 0.0))) || (!(9223372036854775807 >= 42))) {
        switch "foo bar" {
        case "hello":
            for fi39 := 0; fi39 < 17; fi39++ {
                v3 += -128
                var v40 int64 = 0
                _ = fi39
                _ = v40
            }
        case "":
            var v41 uint64 = 0
            _ = v41
        case "foo bar":
            v34 = v34
        default:
            switch "" {
            case "\n\t":
                fmt.Println(v34, v34, v3)
            default:
                var v42 int64 = 872
                _ = v42
            }
        }
        if true {
            v3 += v3
            if (0 == 675) {
                for fi43 := 0; fi43 < 17; fi43++ {
                    for fi44 := 0; fi44 < 6; fi44++ {
                        v45 := GenS1{F0: 0, F1: 1}
                        v46 := GenS4{F0: 42, F1: (67 <= 255)}
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                    v47 := GenS5{F0: 1000, F1: 1}
                    var v48 int16 = int16(fi43)
                    _ = fi43
                    _ = v47
                    _ = v48
                }
            } else {
                v49 := GenS6{F0: (v3 ^ (v3 << ((1) & 63))), F1: 42, F2: -1}
                var v50 int64 = 117
                v3 = (^(v3))
                _ = v49
                _ = v50
            }
        }
        v34 -= 4294967295
    } else {
        for fi51 := 0; fi51 < 3; fi51++ {
            var v52 uint32 = 42
            if true {
                fmt.Println(v34)
                v52 = (v34 >> ((2) & 63))
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    var v53 string = (v1 + "a")
                    if false {
                        if ((!true) && false) {
                            for fi54 := 0; fi54 < 19; fi54++ {
                                v55 := GenS7{F0: false}
                                var v56 uint32 = v34
                                var v57 int = (fi54 - (fi51 % (((-(fi54))) | 1)))
                                _ = fi54
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                            for fi58 := 0; fi58 < 18; fi58++ {
                                _ = len(v2)
                                switch fi58 {
                                case -9223372036854775808:
                                    v34 = (^(v52))
                                case 127:
                                    v3 -= -128
                                case 1:
                                    var v59 int32 = 1
                                    _ = v59
                                default:
                                    fmt.Println(v52, v34)
                                }
                                _ = fi58
                            }
                        } else {
                            fmt.Println(fi51, fi51)
                            switch v53 {
                            case "a":
                                var v60 int64 = 0
                                _ = v60
                            default:
                                fmt.Println(v1, v34, v53)
                            }
                        }
                    } else {
                        fmt.Println(v53, v34, v1)
                    }
                    _ = v53
                }
            }
            _ = fi51
            _ = v52
        }
    }
    if false {
        if false {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    var v61 bool = (!(float64(-2.25) != 1.5))
    if (-1 < -1) {
        v62 := GenS4{F0: 1, F1: (!v61)}
        _ = len(v2)
        _ = len(v2)
        _ = v62
    }
    v34 -= v34
    var v63 int32 = 495
    v61 = false
    for fi64 := 0; fi64 < 9; fi64++ {
        switch 0 {
        case 32768:
            switch 128 {
            case -1:
                for fi65 := 0; fi65 < 17; fi65++ {
                    if (v63 == (v63 & int32(v3))) {
                        v66 := GenS5{F0: -32768, F1: v63}
                        _ = v66
                    }
                    var v67 uint64 = 80
                    if v61 {
                        v68 := GenS4{F0: 0, F1: (!v61)}
                        var v69 int64 = -68
                        v61 = true
                        _ = v68
                        _ = v69
                    }
                    var v70 int16 = 1000
                    _ = fi65
                    _ = v67
                    _ = v70
                }
            default:
                if v61 {
                    if true {
                        switch fi64 {
                        case -1:
                            for fi71 := 0; fi71 < 3; fi71++ {
                                v63 -= v63
                                _ = fi71
                            }
                        default:
                            var v72 float64 = 1.5
                            _ = v72
                        }
                        if v61 {
                            var v73 uint16 = 435
                            _ = v73
                        } else {
                            _ = len(v2)
                            switch (fi64 / (((fi64 - (fi64 ^ 2147483647))) | 1)) {
                            case -973:
                                for fi74 := 0; fi74 < 3; fi74++ {
                                    v1 = v1
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi74
                                }
                            default:
                                v75 := GenS8{F0: (v61 && false), F1: (v1 + ("foo bar" + v1))}
                                _ = v75
                            }
                            for fi76 := 0; fi76 < 19; fi76++ {
                                var v77 uint32 = v34
                                if (255 < 42) {
                                    v34 += 627
                                    var v78 int8 = v3
                                    fmt.Println(v77, v3)
                                    var v79 int = fi76
                                    switch v1 {
                                    case "":
                                        var v80 uint = 65535
                                        _ = v80
                                    case "\n\t":
                                        if (!((-659 >= -1) && v61)) {
                                            fmt.Println(v1, v34)
                                            switch fi64 {
                                            case 32767:
                                                var v81 uint32 = v77
                                                _ = v81
                                            case 32768:
                                                v61 = false
                                            }
                                        } else {
                                            switch fi76 {
                                            case 18:
                                                v82 := GenS7{F0: (v61 || v61)}
                                                _ = v82
                                            case 32768:
                                                v61 = (4294967295 <= 4294967295)
                                            case -2147483648:
                                                var v83 string = v1
                                                _ = v83
                                            }
                                            fmt.Println(v3)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v84 := GenS8{F0: (v61 || (!false)), F1: ("foo bar" + (v1 + (v1 + ("hello" + v1))))}
                                            switch v1 {
                                            case "\n\t":
                                                _ = len(v2)
                                            case "a":
                                                var v85 uint64 = 719
                                                v63 = -2147483648
                                                _ = v85
                                            default:
                                                var v86 int64 = 1
                                                _ = v86
                                            }
                                            v78 = (v3 ^ -128)
                                            _ = v84
                                        }
                                    default:
                                        if (!(255 != 643)) {
                                            v1 = "\n\t"
                                        } else {
                                            if v61 {
                                                var v87 int8 = (v78 * (v3 >> ((42) & 63)))
                                                v63 = v63
                                                _ = v87
                                            } else {
                                                if (!false) {
                                                    _ = len(v2)
                                                    fmt.Println(fi76)
                                                }
                                                var v88 uint8 = 128
                                                v2 = append(v2, (fi76 * (^(fi76))))
                                                _ = v88
                                            }
                                            switch v79 {
                                            case 9223372036854775807:
                                                v3 += v78
                                            case 182:
                                                var v89 bool = v61
                                                _ = v89
                                            case 0:
                                                var v90 int16 = -32768
                                                _ = v90
                                            default:
                                                v3 = (-(v3))
                                            }
                                            v63 += (v63 * v63)
                                            v77 = v34
                                            if (1.5 < (1.5 - 1e10)) {
                                                v61 = (1.5 <= ((-2.25 * 0.0) / 3.14159))
                                            }
                                        }
                                    }
                                    _ = v78
                                    _ = v79
                                }
                                _ = fi76
                                _ = v77
                            }
                        }
                    } else {
                        var v91 float64 = -2.25
                        v2 = append(v2, (fi64 + fi64))
                        _ = v91
                    }
                } else {
                    switch -246 {
                    case -1:
                        fmt.Println(v63, v63, v3)
                    }
                    _ = len(v2)
                    for fi92 := 0; fi92 < 17; fi92++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v3 = 127
                        _ = fi92
                    }
                }
            }
        default:
            if (!(!v61)) {
                v61 = (-505 >= 1)
            }
        }
        _ = fi64
    }
    v3 = (v3 / ((v3) | 1))
    v1 = "a"
    for fi93 := 0; fi93 < 2; fi93++ {
        if v61 {
            if (v3 <= 42) {
                v94 := GenS6{F0: 60, F1: 1, F2: (fi93 & (fi93 * fi93))}
                switch v1 {
                case "\n\t":
                    if (((v63 << ((42) & 63)) > -554) || v61) {
                        fmt.Println(v1)
                        if v61 {
                            switch (fi93 - fi93) {
                            case -753:
                                var v95 bool = v61
                                _ = v95
                            case 32767:
                                fmt.Println(v61, v3)
                            case 0:
                                if v61 {
                                    fmt.Println(v61, v63)
                                    fmt.Println(v1, v34)
                                }
                            }
                            v96 := GenS9{F0: (v1 + ("\n\t" + v1)), F1: 0, F2: 2}
                            _ = v96
                        } else {
                            v2 = append(v2, -9223372036854775808)
                        }
                    } else {
                        v2 = append(v2, 42)
                        for fi97 := 0; fi97 < 6; fi97++ {
                            if len(v2) > 0 {
                                v2[0] = (fi93 | (fi97 & 2147483647))
                            }
                            var v98 uint = 42
                            var v99 uint = (v98 ^ (v98 | v98))
                            _ = fi97
                            _ = v98
                            _ = v99
                        }
                        if v61 {
                            var v100 int64 = -205
                            _ = v100
                        } else {
                            var v101 bool = true
                            v3 = -128
                            _ = v101
                        }
                    }
                default:
                    for fi102 := 0; fi102 < 17; fi102++ {
                        if len(v2) > 0 {
                            v2[0] = fi102
                        }
                        var v103 bool = v61
                        _ = fi102
                        _ = v103
                    }
                }
                _ = v94
            }
            var v104 uint = 2
            var v105 uint32 = v34
            _ = v104
            _ = v105
        } else {
            for fi106 := 0; fi106 < 20; fi106++ {
                var v107 int32 = v63
                _ = fi106
                _ = v107
            }
            v63 = 0
            v1 = "hello"
        }
        v61 = (!((v61 && (fi93 >= (fi93 / ((493) | 1)))) || v61))
        for fi108 := 0; fi108 < 19; fi108++ {
            var v109 uint8 = 255
            switch (fi108 / ((fi108) | 1)) {
            case 9223372036854775807:
                var v110 float64 = 0.0
                _ = v110
            }
            _ = fi108
            _ = v109
        }
        _ = fi93
    }
    for fi111 := 0; fi111 < 11; fi111++ {
        switch fi111 {
        case 42:
            var v112 int16 = -1
            _ = v112
        }
        _ = fi111
    }
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    for fi113 := 0; fi113 < 17; fi113++ {
        var v114 float32 = 0.0
        _ = fi113
        _ = v114
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v61
    _ = v63
}
