package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int }
type GenS2 struct { F0 int; F1 uint64 }
type GenS3 struct { F0 string; F1 int; F2 int64 }
type GenS4 struct { F0 uint32; F1 string; F2 uint32 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint; F1 int16 }
type GenS8 struct { F0 uint64; F1 int8 }
type GenS9 struct { F0 int8; F1 uint8 }
type GenS10 struct { F0 int16; F1 uint16 }
type GenS11 struct { F0 int; F1 uint16; F2 bool }
type GenS12 struct { F0 bool; F1 uint32 }
type GenS13 struct { F0 uint64; F1 int }
type GenS14 struct { F0 bool }
func genFn0() uint32 {
    switch 42 {
    case -722:
        v7 := GenS1{F0: -1, F1: 255, F2: 849}
        _ = v7
    case 0:
        var v8 uint16 = 255
        _ = v8
    case 1:
        v9 := GenS2{F0: 82, F1: 512}
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        if (false || false) {
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := []int{fi11, (fi11 / (((fi10 & (^(fi10)))) | 1)), fi11}
                _ = fi11
                _ = v12
            }
            v13 := GenS3{F0: "a", F1: 32767, F2: -1}
            _ = v13
        }
        var v14 float32 = (3.14159 - (float32((0.0 / (0.0 - -2.25))) - 1.5))
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 int32 = 2147483647
            switch ("" + "hello") {
            case "foo bar":
                v14 = (1.5 / ((0.0 + (float32(3.14159) / v14)) + float32(0.1)))
            case "\n\t":
                var v17 bool = true
                _ = v17
            case "a":
                if (!true) {
                    switch (fi10 | (fi10 >> ((876) & 63))) {
                    case 4294967295:
                        var v18 float32 = v14
                        _ = v18
                    default:
                        v19 := GenS4{F0: 1, F1: "hello", F2: 825}
                        _ = v19
                    }
                    if false {
                        v16 += v16
                        fmt.Println(v14, fi10)
                    }
                    switch ((((("a" + ("\n\t" + "a")) + ("foo bar" + "")) + "") + ("a" + "foo bar")) + "\n\t") {
                    case "\n\t":
                        v20 := GenS4{F0: 773, F1: "", F2: 46}
                        _ = v20
                    case "a":
                        v21 := []int{}
                        _ = v21
                    default:
                        v16 -= (v16 % (((v16 & v16)) | 1))
                        v22 := []int{(fi15 ^ -9223372036854775808)}
                        switch (("a" + "a") + "\n\t") {
                        case "":
                            switch 163 {
                            case 1:
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                            case 128:
                                switch (fi15 & fi15) {
                                case 127:
                                    var v23 int64 = -1
                                    _ = v23
                                case 32767:
                                    switch 9223372036854775807 {
                                    case 128:
                                        var v24 int16 = 1
                                        _ = v24
                                    }
                                default:
                                    fmt.Println(v14)
                                    v14 = ((v14 + float32(0.1)) * v14)
                                    var v25 int16 = 1
                                    var v26 uint = 187
                                    v27 := GenS4{F0: 42, F1: ("hello" + "a"), F2: 847}
                                    fmt.Println(v16, v14)
                                    var v28 float64 = 1.5
                                    v14 = (v14 + 0.0)
                                    fmt.Println(v25)
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            default:
                                if (!true) {
                                    var v29 float32 = 1e10
                                    _ = v29
                                } else {
                                    var v30 uint32 = 748
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                    v14 = v14
                                    v31 := GenS1{F0: (v16 * 0), F1: 18446744073709551615, F2: (fi10 - fi15)}
                                    if false {
                                        fmt.Println(fi15, fi15)
                                        v16 += -881
                                        v32 := GenS4{F0: (v30 << ((42) & 63)), F1: "", F2: (v30 / (((v30 & v30)) | 1))}
                                        _ = v32
                                    } else {
                                        var v33 uint16 = 65535
                                        v30 = v30
                                        var v34 uint16 = (v33 + (^(v33)))
                                        _ = v33
                                        _ = v34
                                    }
                                    _ = v30
                                    _ = v31
                                }
                            }
                        case "a":
                            var v35 bool = true
                            _ = v35
                        case "\n\t":
                            if len(v22) > 0 {
                                v22[0] = (-(fi10))
                            }
                        }
                        _ = v22
                    }
                } else {
                    var v36 uint8 = 1
                    var v37 int32 = 160
                    var v38 int64 = 42
                    _ = v36
                    _ = v37
                    _ = v38
                }
            default:
                v16 = 1
            }
            v14 += (float32(-2.25) - v14)
            _ = fi15
            _ = v16
        }
        _ = fi10
        _ = v14
    }
    return 0
}
func genFn1(p0 int64, p1 uint16, p2 bool) int {
    var v39 string = "a"
    if (p1 != p1) {
        p1 = 32767
        for fi40 := 0; fi40 < 14; fi40++ {
            for fi41 := 0; fi41 < 16; fi41++ {
                var v42 int64 = p0
                p0 -= v42
                _ = fi41
                _ = v42
            }
            var v43 float32 = (1e10 - 1.5)
            _ = fi40
            _ = v43
        }
    } else {
        v44 := GenS5{F0: uint32(p1), F1: 0}
        for fi45 := 0; fi45 < 11; fi45++ {
            if false {
                p0 -= p0
                var v46 int8 = 127
                _ = v46
            } else {
                if (false && (((v39 < v39) && p2) || p2)) {
                    p2 = p2
                } else {
                    v39 = v39
                }
            }
            var v47 float32 = 1e10
            _ = fi45
            _ = v47
        }
        var v48 float32 = 1e10
        _ = v44
        _ = v48
    }
    if p2 {
        p0 -= (p0 + -1)
        if (0 <= 426) {
            v49 := []int{-1, 259, 895, 42}
            switch 9223372036854775807 {
            case 127:
                if (-949 >= 0) {
                    v50 := GenS6{F0: 202}
                    _ = v50
                }
            }
            if ((p1 >> ((65535) & 63)) != (p1 >> ((41) & 63))) {
                p1 += p1
                for fi51 := 0; fi51 < 13; fi51++ {
                    fmt.Println(fi51, fi51, p0)
                    p0 = 0
                    if true {
                        var v52 float32 = 1e10
                        _ = v52
                    }
                    _ = fi51
                }
                fmt.Println(p0)
            } else {
                v49 = append(v49, 263)
            }
            _ = v49
        }
    } else {
        v53 := GenS5{F0: 0, F1: 4294967295}
        v39 = v39
        _ = v53
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v39
    return -9223372036854775808
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    v4 := make(map[string]uint32)
    var v5 float32 = (1e10 * float32(0.1))
    var v6 int64 = -1
    switch "foo bar" {
    case "\n\t":
        var v54 string = "a"
        _ = v54
    case "foo bar":
        fmt.Println(v5, v5)
    }
    fmt.Println(v5)
    if (!(-1 != -32768)) {
        switch "hello" {
        case "foo bar":
            v55 := GenS7{F0: 425, F1: 32767}
            _ = v55
        default:
            fmt.Println(v5, v5)
        }
        v56 := []int{}
        v1 -= (v1 % (((v1 | 952)) | 1))
        _ = v56
    } else {
        if true {
            fmt.Println(v6, v5)
            _ = len(v4)
            v57 := []int{0, 255}
            _ = v57
        }
    }
    for fi58 := 0; fi58 < 9; fi58++ {
        v1 += (v1 - (v1 | v1))
        _ = fi58
    }
    switch "a" {
    case "a":
        var v59 int64 = (v6 ^ v6)
        _ = v59
    case "hello":
        for fi60 := 0; fi60 < 9; fi60++ {
            v61 := GenS3{F0: "a", F1: fi60, F2: (v6 << ((347) & 63))}
            v6 += v6
            v62 := GenS2{F0: (fi60 / ((genFn1(653, v1, (v5 <= v5))) | 1)), F1: 1}
            _ = fi60
            _ = v61
            _ = v62
        }
    }
    var v63 int8 = 42
    var v64 float32 = 0.1
    if false {
        for fi65 := 0; fi65 < 11; fi65++ {
            v63 = v63
            _ = fi65
        }
        for fi66 := 0; fi66 < 6; fi66++ {
            v67 := GenS2{F0: (fi66 ^ fi66), F1: 0}
            _ = fi66
            _ = v67
        }
    } else {
        var v68 int = 1
        v64 += float32(-2.25)
        v69 := GenS8{F0: 0, F1: v63}
        _ = v68
        _ = v69
    }
    v70 := GenS9{F0: (v63 + v63), F1: 255}
    if (190 == 0) {
        if false {
            switch ("foo bar" + (("" + "") + "")) {
            case "foo bar":
                for fi71 := 0; fi71 < 8; fi71++ {
                    var v72 uint = 42
                    for fi73 := 0; fi73 < 18; fi73++ {
                        if false {
                            switch ("" + (("a" + "\n\t") + (("" + "\n\t") + "foo bar"))) {
                            case "hello":
                                if ((2 <= 23) && false) {
                                    v1 = v1
                                } else {
                                    v1 += v1
                                    var v74 uint16 = (v1 | (^(v1)))
                                    _ = v74
                                }
                            case "a":
                                var v75 float64 = 0.1
                                _ = v75
                            default:
                                if false {
                                    var v76 int16 = 0
                                    fmt.Println(v64, fi71, v1)
                                    var v77 uint = 4294967295
                                    _ = v76
                                    _ = v77
                                }
                            }
                            var v78 string = "\n\t"
                            _ = v78
                        }
                        fmt.Println(v63, v1, v64)
                        _ = fi73
                    }
                    _ = fi71
                    _ = v72
                }
            default:
                if (("foo bar" + "foo bar") != "") {
                    fmt.Println(v63, v63, v5)
                    var v79 uint32 = 4294967295
                    _ = v79
                } else {
                    if (!(!true)) {
                        fmt.Println(v5)
                        fmt.Println(v1, v64)
                    } else {
                        fmt.Println(v1, v5)
                    }
                    var v80 int8 = 42
                    if ((true || (((1 <= 1) && true) && false)) && (false && true)) {
                        switch 42 {
                        case -2147483648:
                            fmt.Println(v1, v63)
                        case 635:
                            delete(v3, ("foo bar" + ("foo bar" + "\n\t")))
                        }
                        if ((v1 / ((65535) | 1)) >= 2) {
                            var v81 int32 = 2147483647
                            if (true && ((("hello" + "hello") + ("a" + "a")) > "hello")) {
                                if false {
                                    var v82 int64 = v6
                                    _ = v82
                                } else {
                                    v83 := []int{9223372036854775807, 128}
                                    if (!false) {
                                        v84 := GenS2{F0: genFn1(v6, (v1 - (v1 | 255)), (514 <= v81)), F1: 102}
                                        var v85 bool = false
                                        _ = v84
                                        _ = v85
                                    }
                                    v86 := GenS10{F0: -32768, F1: 42}
                                    _ = v83
                                    _ = v86
                                }
                                v64 += (-2.25 / (float32((0.0 - 1.5)) + (1e10 / (3.14159 + v64))))
                                v81 -= v81
                            } else {
                                var v87 bool = (v1 > v1)
                                _ = v87
                            }
                            _ = v81
                        }
                        v88 := GenS3{F0: (("\n\t" + "foo bar") + ((("\n\t" + "foo bar") + "") + ("" + ("a" + ("hello" + ("" + "foo bar")))))), F1: 2147483647, F2: (v6 | v6)}
                        _ = v88
                    } else {
                        v89 := GenS3{F0: "hello", F1: 127, F2: v6}
                        for fi90 := 0; fi90 < 10; fi90++ {
                            v5 += v5
                            _ = fi90
                        }
                        _ = v89
                    }
                    _ = v80
                }
            }
            var v91 int = -2147483648
            if (v6 > v6) {
                for fi92 := 0; fi92 < 9; fi92++ {
                    var v93 uint8 = 42
                    _ = fi92
                    _ = v93
                }
            } else {
                var v94 int8 = (v63 * 113)
                v95 := GenS11{F0: v91, F1: (v1 - 409), F2: true}
                var v96 int8 = v63
                _ = v94
                _ = v95
                _ = v96
            }
            _ = v91
        }
    }
    v97 := GenS11{F0: -373, F1: (v1 * v1), F2: true}
    var v98 uint8 = 0
    v99 := GenS12{F0: (32767 <= 1), F1: 0}
    if (42 <= -202) {
        v63 = v63
    } else {
        if false {
            v100 := v2[-1]
            for fi101 := 0; fi101 < 3; fi101++ {
                v102 := []int{genFn1(v6, 1, v100), 32768}
                if (v5 != v64) {
                    v103 := GenS13{F0: 42, F1: fi101}
                    if len(v102) >= 2 {
                        v102 = v102[1:len(v102)]
                    }
                    _ = v103
                }
                var v104 uint32 = 42
                _ = fi101
                _ = v102
                _ = v104
            }
            _ = v100
        } else {
            v105 := GenS10{F0: 534, F1: 539}
            v5 += (v64 / -2.25)
            _ = v105
        }
        if true {
            for fi106 := 0; fi106 < 15; fi106++ {
                v64 = 1e10
                _ = fi106
            }
            var v107 uint8 = (v98 ^ 145)
            for fi108 := 0; fi108 < 7; fi108++ {
                v109 := []int{(^(fi108)), fi108, genFn1((v6 * (v6 & (v6 - 42))), (v1 + 42), false)}
                for fi110 := 0; fi110 < 8; fi110++ {
                    var v111 float32 = v64
                    switch (fi108 - (fi110 / ((int(v98)) | 1))) {
                    case 9223372036854775807:
                        v107 += 1
                    case 380:
                        var v112 int64 = -1
                        _ = v112
                    default:
                        switch 128 {
                        case -327:
                            var v113 int32 = 316
                            _ = v113
                        case 127:
                            v5 = 0.0
                        default:
                            fmt.Println(v6, v6)
                        }
                    }
                    _ = len(v109)
                    _ = fi110
                    _ = v111
                }
                _ = fi108
                _ = v109
            }
            _ = v107
        } else {
            switch "hello" {
            case "\n\t":
                for fi114 := 0; fi114 < 11; fi114++ {
                    for fi115 := 0; fi115 < 1; fi115++ {
                        v63 = int8(v1)
                        var v116 uint8 = (^(v98))
                        v117 := GenS13{F0: 1, F1: fi115}
                        _ = fi115
                        _ = v116
                        _ = v117
                    }
                    _ = fi114
                }
            case "":
                v98 -= v98
            case "a":
                v118 := GenS14{F0: ((0.1 + 3.14159) <= (float64((v64 * float32(-2.25))) - 0.0))}
                _ = v118
            }
            v119 := []int{genFn1((v6 << ((42) & 63)), 334, false), -404}
            var v120 float64 = 0.0
            _ = v119
            _ = v120
        }
        fmt.Println(v64, v6)
    }
    v98 += (v98 / ((v98) | 1))
    for fi121 := 0; fi121 < 3; fi121++ {
        if (1 < 2) {
            var v122 float32 = -2.25
            var v123 string = ""
            _ = v122
            _ = v123
        }
        _ = fi121
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v70
    _ = v97
    _ = v98
    _ = v99
}
