package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 int; F1 uint16 }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 uint32; F1 uint64; F2 int }
type GenS8 struct { F0 int64; F1 int32; F2 uint }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint16; F1 uint64 }
type GenS12 struct { F0 uint32 }
func genFn0(p0 int, p1 uint, p2 uint16) bool {
    var v6 uint = (p1 ^ 255)
    v7 := []int{(p0 >> ((v6) & 63)), p0}
    var v8 int64 = 9223372036854775807
    p2 = 633
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 uint32, p1 string) int8 {
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 1, F1: 1, F2: -591}
        p1 = "hello"
        var v11 uint8 = 42
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(p0, p1)
    for fi12 := 0; fi12 < 19; fi12++ {
        if (166 != 0) {
            var v13 uint32 = (p0 ^ (p0 & 65535))
            _ = v13
        } else {
            var v14 int64 = -142
            var v15 uint = 42
            _ = v14
            _ = v15
        }
        var v16 int8 = -39
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 uint = 517
            _ = fi17
            _ = v18
        }
        _ = fi12
        _ = v16
    }
    if true {
        var v19 uint8 = 0
        _ = v19
    } else {
        for fi20 := 0; fi20 < 12; fi20++ {
            if (!(float64((-2.25 + (float32(-2.25) * (-2.25 - 3.14159)))) <= 0.0)) {
                for fi21 := 0; fi21 < 8; fi21++ {
                    p1 = (((p1 + p1) + (p1 + "\n\t")) + (p1 + "a"))
                    var v22 uint = 1
                    p1 = p1
                    _ = fi21
                    _ = v22
                }
            } else {
                v23 := []int{}
                switch p1 {
                case "\n\t":
                    for fi24 := 0; fi24 < 17; fi24++ {
                        p0 = 1
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                        _ = fi24
                    }
                }
                _ = v23
            }
            switch fi20 {
            case 42:
                if true {
                    v25 := GenS2{F0: fi20}
                    _ = v25
                }
            case 2147483647:
                var v26 uint32 = (^(p0))
                _ = v26
            case -970:
                var v27 uint = 42
                _ = v27
            }
            _ = fi20
        }
    }
    var v28 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v28
    return 127
}

func main() {
    var v1 int16 = -235
    var v2 int = -9223372036854775808
    v3 := make(map[int]bool)
    var v4 string = "a"
    v5 := []uint32{}
    for fi29 := 0; fi29 < 19; fi29++ {
        var v30 int8 = -79
        var v31 uint = 65535
        for fi32 := 0; fi32 < 4; fi32++ {
            switch v4 {
            case "\n\t":
                v33 := GenS1{F0: 0, F1: v30, F2: v1}
                _ = v33
            default:
                _ = len(v5)
            }
            var v34 uint32 = 42
            _ = fi32
            _ = v34
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    if (128 < 127) {
        var v35 int16 = v1
        _ = v35
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        fmt.Println(v1, v4)
        for fi37 := 0; fi37 < 11; fi37++ {
            var v38 bool = (0.1 >= 1.5)
            _ = fi37
            _ = v38
        }
        for fi39 := 0; fi39 < 15; fi39++ {
            var v40 int32 = -2147483648
            switch (fi36 * fi39) {
            case 42:
                v5 = append(v5, 2147483647)
            default:
                if true {
                    var v41 uint = 65535
                    v42 := GenS2{F0: fi36}
                    v43 := GenS3{F0: 203, F1: 296}
                    _ = v41
                    _ = v42
                    _ = v43
                }
            }
            _ = fi39
            _ = v40
        }
        _ = fi36
    }
    switch -37 {
    case 128:
        var v44 int = (^(v2))
        _ = v44
    case 32767:
        if true {
            fmt.Println(v2)
            for fi45 := 0; fi45 < 13; fi45++ {
                for fi46 := 0; fi46 < 4; fi46++ {
                    v47 := GenS4{F0: -458, F1: -171}
                    v5 = append(v5, 42)
                    var v48 uint8 = 127
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                _ = fi45
            }
        } else {
            v1 = v1
        }
    case 32768:
        var v49 int8 = 100
        _ = v49
    default:
        delete(v3, (^(v2)))
    }
    _ = len(v3)
    v50 := GenS5{F0: (v2 % (((v2 / ((v2) | 1))) | 1)), F1: 255}
    v4 = (v4 + (v4 + ("a" + ((((v4 + v4) + (v4 + v4)) + (v4 + v4)) + ("foo bar" + "hello")))))
    v51 := GenS4{F0: 226, F1: -745}
    fmt.Println(v1)
    fmt.Println(v2)
    if false {
        var v52 float64 = (float64(float32(1e10)) * 0.0)
        var v53 float64 = 0.0
        var v54 string = ("\n\t" + "a")
        _ = v52
        _ = v53
        _ = v54
    } else {
        var v55 bool = true
        if v55 {
            v56 := v3[v2]
            var v57 bool = ((float32(1e10) * ((0.1 - 1e10) * (1e10 / 1.5))) == (0.1 + float32((0.1 * 0.1))))
            var v58 uint16 = 58
            _ = v56
            _ = v57
            _ = v58
        }
        for fi59 := 0; fi59 < 3; fi59++ {
            v60 := GenS2{F0: 32767}
            if v55 {
                v4 = "hello"
                if v55 {
                    _ = len(v3)
                }
                fmt.Println(v2)
            } else {
                fmt.Println(v4, v2, v4)
            }
            _ = fi59
            _ = v60
        }
        _ = v55
    }
    var v61 int64 = 1
    v62 := GenS1{F0: 83, F1: 42, F2: (v1 ^ (^(v1)))}
    var v63 string = "hello"
    var v64 float32 = ((0.0 - (3.14159 - (-2.25 * ((((3.14159 - (3.14159 + 3.14159)) - -2.25) / 3.14159) * (1e10 * 0.0))))) * 3.14159)
    var v65 uint16 = 32767
    switch (-(v2)) {
    case 135:
        if (!(float32(0.0) < ((v64 - ((v64 / (1.5 - (((((1.5 - 0.0) / v64) / (3.14159 + v64)) - ((float32(float64(v64)) - float32(1e10)) * 1e10)) * v64))) - ((1e10 - ((v64 - (((v64 + 0.0) * ((((v64 + v64) - float32(-2.25)) + v64) / 3.14159)) * v64)) - v64)) + (1.5 / 1.0)))) * (0.0 * v64)))) {
            for fi66 := 0; fi66 < 9; fi66++ {
                if len(v5) > 0 {
                    v5[0] = 2147483647
                }
                if (false || true) {
                    v2 = (v2 << ((1) & 63))
                    if false {
                        if true {
                            fmt.Println(fi66, v63, fi66)
                            switch -1 {
                            case 127:
                                if (v63 < v4) {
                                    switch (v4 + v4) {
                                    case "foo bar":
                                        fmt.Println(v63, v4, v65)
                                    case "a":
                                        var v67 int = (v2 | fi66)
                                        _ = v67
                                    case "":
                                        v63 = v63
                                    }
                                    v64 = 0.0
                                }
                            case 0:
                                v68 := GenS6{F0: int8(v65), F1: (42 > (v1 / (((^(v1))) | 1)))}
                                _ = v68
                            case 1:
                                v61 = v61
                            }
                        } else {
                            v69 := GenS3{F0: 0, F1: v65}
                            switch v63 {
                            case "":
                                if true {
                                    for fi70 := 0; fi70 < 1; fi70++ {
                                        v1 = (v1 % (((v1 + -514)) | 1))
                                        switch (v2 >> ((465) & 63)) {
                                        case 268:
                                            v63 = "a"
                                        case 32768:
                                            v1 = -253
                                        default:
                                            fmt.Println(v2, fi66)
                                        }
                                        v5 = append(v5, 797)
                                        _ = fi70
                                    }
                                    var v71 float64 = float64(-2.25)
                                    if false {
                                        if false {
                                            fmt.Println(v63, v1, v1)
                                        } else {
                                            _ = len(v5)
                                            switch (v4 + v4) {
                                            case "hello":
                                                fmt.Println(fi66)
                                            default:
                                                v63 = "\n\t"
                                            }
                                            var v72 int = -972
                                            _ = v72
                                        }
                                    } else {
                                        for fi73 := 0; fi73 < 18; fi73++ {
                                            var v74 int8 = 42
                                            _ = fi73
                                            _ = v74
                                        }
                                        switch fi66 {
                                        case 666:
                                            for fi75 := 0; fi75 < 18; fi75++ {
                                                switch fi75 {
                                                case 1:
                                                    var v76 uint16 = 784
                                                    _ = v76
                                                }
                                                v3[v2] = false
                                                var v77 int32 = 791
                                                _ = fi75
                                                _ = v77
                                            }
                                        default:
                                            for fi78 := 0; fi78 < 19; fi78++ {
                                                switch fi66 {
                                                case 2147483647:
                                                    v79 := GenS7{F0: 2147483648, F1: 4294967295, F2: fi66}
                                                    _ = v79
                                                case 1:
                                                    if true {
                                                        if false {
                                                            var v80 float64 = v71
                                                            if true {
                                                                if true {
                                                                    var v81 uint32 = 417
                                                                    var v82 uint8 = 240
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                            }
                                                            _ = v80
                                                        } else {
                                                            v83 := GenS8{F0: v61, F1: -2147483648, F2: 445}
                                                            v61 += (v61 >> ((65535) & 63))
                                                            _ = v83
                                                        }
                                                        v5 = append(v5, 790)
                                                        fmt.Println(fi78, v65)
                                                    } else {
                                                        v61 -= (v61 & -1)
                                                    }
                                                case 86:
                                                    v84 := GenS9{F0: 1}
                                                    _ = v84
                                                }
                                                fmt.Println(fi66, v4, v4)
                                                var v85 int64 = int64(v65)
                                                _ = fi78
                                                _ = v85
                                            }
                                        }
                                        for fi86 := 0; fi86 < 19; fi86++ {
                                            if len(v5) > 0 {
                                                v5[0] = 4294967295
                                            }
                                            _ = fi86
                                        }
                                    }
                                    _ = v71
                                }
                            }
                            if ((v64 == -2.25) && (128 != 255)) {
                                fmt.Println(v4, v64)
                                for fi87 := 0; fi87 < 11; fi87++ {
                                    v88 := v3[fi66]
                                    if (!v88) {
                                        var v89 float64 = 1e10
                                        v63 = v4
                                        fmt.Println(v63)
                                        _ = v89
                                    } else {
                                        var v90 bool = v88
                                        if len(v5) > 0 {
                                            v5[0] = 2
                                        }
                                        switch (v2 % ((-407) | 1)) {
                                        case 42:
                                            v3[fi87] = false
                                        case 4294967295:
                                            switch fi87 {
                                            case 2147483647:
                                                if (!v90) {
                                                    v91 := GenS1{F0: 1, F1: -1, F2: v1}
                                                    var v92 float64 = -2.25
                                                    v4 = "a"
                                                    var v93 int = (^(fi87))
                                                    switch v4 {
                                                    case "":
                                                        var v94 uint16 = v65
                                                        _ = v94
                                                    case "a":
                                                        v92 -= 3.14159
                                                        v63 = v63
                                                    }
                                                    _ = v91
                                                    _ = v92
                                                    _ = v93
                                                }
                                            case 9223372036854775807:
                                                if (v65 >= 65535) {
                                                    switch v4 {
                                                    case "\n\t":
                                                        switch fi87 {
                                                        case 366:
                                                            v95 := GenS10{F0: -1}
                                                            _ = v95
                                                        case 0:
                                                            if (v63 > (((v63 + v4) + ((v63 + v4) + "hello")) + v4)) {
                                                                if v90 {
                                                                    var v96 int64 = (v61 + -9223372036854775808)
                                                                    v1 = (v1 ^ 1)
                                                                    var v97 uint = 18446744073709551615
                                                                    var v98 int = fi66
                                                                    v99 := GenS9{F0: 0}
                                                                    _ = v96
                                                                    _ = v97
                                                                    _ = v98
                                                                    _ = v99
                                                                }
                                                            }
                                                        default:
                                                            v61 += int64(v1)
                                                        }
                                                    }
                                                    var v100 uint = 817
                                                    _ = v100
                                                }
                                            }
                                        }
                                        _ = v90
                                    }
                                    v61 -= v61
                                    v101 := GenS11{F0: 193, F1: 1}
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if v88 {
                                        fmt.Println(v2)
                                        v65 = (v65 ^ 42)
                                        var v102 int16 = int16(fi87)
                                        _ = v102
                                    }
                                    var v103 float32 = -2.25
                                    _ = fi87
                                    _ = v101
                                    _ = v103
                                    _ = v88
                                }
                            }
                            _ = v69
                        }
                        var v104 float64 = 0.0
                        _ = v104
                    }
                }
                fmt.Println(v4, v64)
                _ = fi66
            }
            var v105 int = (-(v2))
            _ = v105
        } else {
            v61 -= v61
            for fi106 := 0; fi106 < 15; fi106++ {
                if true {
                    for fi107 := 0; fi107 < 6; fi107++ {
                        v108 := GenS3{F0: 2, F1: (^(v65))}
                        var v109 int = 32767
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi107
                        _ = v108
                        _ = v109
                    }
                    v1 -= -1
                    fmt.Println(v2)
                } else {
                    if true {
                        v65 += 32767
                    }
                    for fi110 := 0; fi110 < 11; fi110++ {
                        v2 += (^(v2))
                        _ = fi110
                    }
                }
                v111 := GenS12{F0: 90}
                _ = fi106
                _ = v111
            }
            if true {
                var v112 uint64 = 18446744073709551615
                _ = v112
            } else {
                v61 -= v61
                v63 = v63
                for fi113 := 0; fi113 < 18; fi113++ {
                    if (!false) {
                        var v114 float64 = ((3.14159 / 1e10) / 1.0)
                        v65 = (^(v65))
                        var v115 float64 = v114
                        _ = v114
                        _ = v115
                    }
                    _ = fi113
                }
            }
        }
    case 990:
        fmt.Println(v63)
    case 50:
        _ = len(v5)
    default:
        v116 := GenS6{F0: -28, F1: false}
        _ = v116
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
