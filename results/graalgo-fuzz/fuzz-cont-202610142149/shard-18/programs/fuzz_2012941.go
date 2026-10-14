package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int32; F1 uint8; F2 int8 }
type GenS6 struct { F0 uint16; F1 uint8; F2 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 uint64 }
func genFn0(p0 int) int64 {
    var v6 int16 = 42
    var v7 float32 = (1e10 / (0.0 + (0.1 - -2.25)))
    _ = p0
    _ = v6
    _ = v7
    return 73
}
func genFn1(p0 int16, p1 uint16, p2 int32) int16 {
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: (fi8 << ((42) & 63))}
        if true {
            p0 = 1000
            if (false && (!false)) {
                for fi10 := 0; fi10 < 11; fi10++ {
                    v11 := GenS1{F0: fi8}
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: (fi8 / (((-(fi8))) | 1))}
                if false {
                    p1 -= p1
                    switch "hello" {
                    case "hello":
                        switch "\n\t" {
                        case "a":
                            v13 := GenS2{F0: 42, F1: 242}
                            _ = v13
                        case "":
                            for fi14 := 0; fi14 < 3; fi14++ {
                                p2 = p2
                                var v15 int32 = p2
                                _ = fi14
                                _ = v15
                            }
                        default:
                            var v16 float32 = float32(0.0)
                            _ = v16
                        }
                    case "foo bar":
                        if ((("hello" + "foo bar") <= "foo bar") && true) {
                            v17 := GenS3{F0: p1, F1: int32(p1)}
                            _ = v17
                        } else {
                            p2 -= p2
                        }
                    case "\n\t":
                        v18 := []int{-1}
                        _ = v18
                    }
                } else {
                    v19 := []int{42}
                    var v20 float64 = 0.1
                    _ = v19
                    _ = v20
                }
                _ = v12
            } else {
                if true {
                    v21 := GenS1{F0: (fi8 % (((fi8 + fi8)) | 1))}
                    if true {
                        var v22 uint16 = (p1 - p1)
                        p1 += 65535
                        v23 := []int{(fi8 + -173), -538, -2147483648, (fi8 & (fi8 ^ (fi8 + fi8)))}
                        _ = v22
                        _ = v23
                    }
                    _ = v21
                } else {
                    v24 := []int{-478}
                    var v25 uint16 = p1
                    _ = v24
                    _ = v25
                }
                v26 := GenS1{F0: (fi8 | (fi8 % ((fi8) | 1)))}
                _ = v26
            }
        }
        _ = fi8
        _ = v9
    }
    v27 := GenS1{F0: 128}
    for fi28 := 0; fi28 < 13; fi28++ {
        v29 := GenS1{F0: 191}
        var v30 int16 = (p0 * (p0 - (p0 >> ((18446744073709551615) & 63))))
        _ = fi28
        _ = v29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return (p0 % ((714) | 1))
}

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = -112
    var v3 int = 0
    var v4 uint32 = 0
    var v5 bool = (!(!(0 == 42)))
    for fi31 := 0; fi31 < 15; fi31++ {
        switch ("hello" + (("hello" + "foo bar") + "")) {
        case "foo bar":
            v32 := GenS1{F0: (v3 % ((v3) | 1))}
            _ = v32
        case "":
            v1["foo bar"] = (("" + "\n\t") >= "foo bar")
        case "hello":
            fmt.Println(v5, v3)
        default:
            if v5 {
                if v5 {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v34 := []int{}
                        var v35 uint8 = 127
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
                for fi36 := 0; fi36 < 19; fi36++ {
                    var v37 string = "\n\t"
                    _ = fi36
                    _ = v37
                }
            }
        }
        if v5 {
            if (!v5) {
                v5 = v5
                var v38 int64 = -373
                _ = v38
            } else {
                v3 = (-(v3))
                v39 := []int{(fi31 & fi31), fi31, int(v4), -9223372036854775808}
                v4 = 4294967295
                _ = v39
            }
            v40 := GenS4{F0: 0}
            v41 := GenS1{F0: v3}
            _ = v40
            _ = v41
        } else {
            for fi42 := 0; fi42 < 8; fi42++ {
                var v43 int32 = 2147483647
                var v44 string = "a"
                v45 := GenS4{F0: (v43 * v43)}
                _ = fi42
                _ = v43
                _ = v44
                _ = v45
            }
            for fi46 := 0; fi46 < 8; fi46++ {
                var v47 int16 = genFn1(977, 749, 0)
                _ = fi46
                _ = v47
            }
            if v5 {
                v48 := []int{fi31, (fi31 % ((v3) | 1)), fi31}
                delete(v1, (("" + "") + "a"))
                v3 += (v3 ^ (^(v3)))
                _ = v48
            } else {
                switch "" {
                case "foo bar":
                    v49 := []int{v3}
                    _ = v49
                case "":
                    fmt.Println(v2, fi31, v5)
                case "hello":
                    v2 -= (v2 / (((^(v2))) | 1))
                default:
                    var v50 string = "\n\t"
                    _ = v50
                }
                v51 := []int{(-(v3)), v3, (-(fi31))}
                _ = v51
            }
        }
        v52 := GenS3{F0: 133, F1: -1}
        _ = fi31
        _ = v52
    }
    for fi53 := 0; fi53 < 19; fi53++ {
        v54 := GenS2{F0: v3, F1: 2}
        fmt.Println(v5, v3)
        _ = fi53
        _ = v54
    }
    var v55 string = ""
    v2 = (-(v2))
    v5 = (((1e10 * 1e10) / -2.25) <= (1.5 / 1.5))
    switch v3 {
    case 128:
        fmt.Println(v5, v5)
    default:
        switch v3 {
        case 9223372036854775807:
            if v5 {
                v56 := []int{}
                for fi57 := 0; fi57 < 6; fi57++ {
                    v5 = v5
                    v2 = v2
                    v58 := GenS5{F0: 42, F1: 127, F2: v2}
                    _ = fi57
                    _ = v58
                }
                _ = v56
            }
        case -465:
            if true {
                switch v55 {
                case "hello":
                    v2 += (v2 & v2)
                case "foo bar":
                    var v59 float32 = 3.14159
                    _ = v59
                case "a":
                    if (((v55 < "\n\t") && ((0 >= int8(v3)) && (32768 <= (v3 - 127)))) && (-197 > 42)) {
                        var v60 int32 = 0
                        v2 = v2
                        _ = v60
                    } else {
                        for fi61 := 0; fi61 < 14; fi61++ {
                            v55 = v55
                            v4 -= v4
                            _ = fi61
                        }
                        if v5 {
                            var v62 float32 = -2.25
                            var v63 uint8 = 2
                            v64 := GenS3{F0: 131, F1: -1}
                            _ = v62
                            _ = v63
                            _ = v64
                        } else {
                            switch v55 {
                            case "a":
                                var v65 int16 = 492
                                _ = v65
                            }
                        }
                    }
                default:
                    v66 := GenS2{F0: (v3 | v3), F1: 128}
                    _ = v66
                }
                var v67 float64 = (0.0 - 3.14159)
                v2 = 1
                _ = v67
            }
        }
    }
    v2 -= (v2 % ((v2) | 1))
    if (219 <= 4294967295) {
        for fi68 := 0; fi68 < 1; fi68++ {
            for fi69 := 0; fi69 < 11; fi69++ {
                if ((fi69 * v3) >= fi69) {
                    var v70 int8 = v2
                    v71 := GenS3{F0: 42, F1: 153}
                    var v72 uint16 = 1
                    if (v70 != v2) {
                        v2 -= 63
                    } else {
                        v2 += -125
                        var v73 float32 = 0.1
                        switch 9223372036854775807 {
                        case -1:
                            v5 = v5
                        }
                        _ = v73
                    }
                    var v74 int32 = 756
                    switch (fi68 | 127) {
                    case 2147483647:
                        var v75 int = 2147483647
                        _ = v75
                    case 9223372036854775807:
                        v72 = v72
                    }
                    _ = v70
                    _ = v71
                    _ = v72
                    _ = v74
                }
                v76 := GenS6{F0: 299, F1: 2, F2: v55}
                _ = fi69
                _ = v76
            }
            if (v55 <= "") {
                switch v55 {
                case "":
                    v4 -= v4
                case "hello":
                    var v77 int = (v3 + -1)
                    _ = v77
                case "foo bar":
                    v55 = "\n\t"
                }
                var v78 float32 = 3.14159
                _ = v78
            } else {
                switch int(v4) {
                case -2147483648:
                    if false {
                        for fi79 := 0; fi79 < 11; fi79++ {
                            v80 := GenS7{F0: -128}
                            v81 := GenS5{F0: int32(fi68), F1: 0, F2: v2}
                            switch ((((v55 + (("\n\t" + (v55 + (v55 + v55))) + v55)) + ((("\n\t" + (v55 + (v55 + v55))) + v55) + ((v55 + ("hello" + (v55 + (v55 + "foo bar")))) + v55))) + v55) + (v55 + v55)) {
                            case "a":
                                switch int(v4) {
                                case 2147483647:
                                    v82 := GenS6{F0: 42, F1: 255, F2: "hello"}
                                    _ = v82
                                }
                            case "":
                                var v83 uint16 = 0
                                _ = v83
                            case "foo bar":
                                fmt.Println(v55)
                            default:
                                fmt.Println(v5, fi79, v5)
                            }
                            _ = fi79
                            _ = v80
                            _ = v81
                        }
                        v84 := []int{179}
                        _ = v84
                    } else {
                        fmt.Println(v55)
                    }
                case 127:
                    v85 := GenS2{F0: (-(fi68)), F1: 2}
                    _ = v85
                case -1:
                    for fi86 := 0; fi86 < 16; fi86++ {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v55, v4, fi86)
                        case "foo bar":
                            switch v55 {
                            case "foo bar":
                                var v87 bool = v5
                                _ = v87
                            case "":
                                if false {
                                    fmt.Println(fi68, v5)
                                    var v88 int8 = v2
                                    var v89 float32 = 0.1
                                    if (!v5) {
                                        v90 := GenS8{F0: ((((v55 + v55) + (v55 + v55)) + ("\n\t" + (v55 + "foo bar"))) < v55), F1: v4}
                                        _ = v90
                                    }
                                    _ = v88
                                    _ = v89
                                }
                            default:
                                fmt.Println(v4, v4, fi86)
                            }
                        default:
                            switch (v55 + "hello") {
                            case "foo bar":
                                v91 := GenS5{F0: -1, F1: 255, F2: v2}
                                _ = v91
                            }
                            switch (-(fi86)) {
                            case 612:
                                var v92 uint8 = 84
                                _ = v92
                            case -9223372036854775808:
                                var v93 float64 = -2.25
                                _ = v93
                            case -1:
                                var v94 uint8 = 127
                                _ = v94
                            default:
                                if true {
                                    v95 := []int{fi86, fi86, (^(v3)), fi86}
                                    _ = v95
                                }
                            }
                        }
                        _ = fi86
                    }
                }
            }
            _ = fi68
        }
        for fi96 := 0; fi96 < 14; fi96++ {
            v97 := GenS9{F0: 4294967295}
            v98 := []int{fi96, (^(v3))}
            _ = fi96
            _ = v97
            _ = v98
        }
        if v5 {
            switch v3 {
            case 492:
                if v5 {
                    v3 = -233
                    for fi99 := 0; fi99 < 6; fi99++ {
                        var v100 string = v55
                        if true {
                            var v101 string = (v55 + v55)
                            v102 := GenS7{F0: v2}
                            _ = v101
                            _ = v102
                        }
                        var v103 float64 = (1e10 + (1e10 - 3.14159))
                        _ = fi99
                        _ = v100
                        _ = v103
                    }
                }
            case -9223372036854775808:
                v3 = v3
            case 1:
                fmt.Println(v3, v5, v5)
            default:
                v5 = ("" != "a")
            }
            for fi104 := 0; fi104 < 7; fi104++ {
                fmt.Println(v5, fi104, v4)
                _ = fi104
            }
        }
    } else {
        v5 = v5
        v105 := []int{(v3 ^ v3), v3}
        for fi106 := 0; fi106 < 17; fi106++ {
            switch ((("" + "hello") + "hello") + v55) {
            case "\n\t":
                var v107 int32 = 384
                _ = v107
            default:
                for fi108 := 0; fi108 < 7; fi108++ {
                    var v109 int32 = 0
                    _ = fi108
                    _ = v109
                }
            }
            switch 0 {
            case 0:
                delete(v1, (((v55 + "\n\t") + "hello") + "foo bar"))
            case -421:
                var v110 bool = v5
                _ = v110
            case 767:
                for fi111 := 0; fi111 < 14; fi111++ {
                    if (!false) {
                        v112 := GenS7{F0: -1}
                        _ = v112
                    }
                    _ = fi111
                }
            }
            _ = fi106
        }
        _ = v105
    }
    v5 = v5
    for fi113 := 0; fi113 < 6; fi113++ {
        v55 = ("" + ((v55 + v55) + v55))
        if (42 != 134) {
            v2 += v2
            switch 4294967295 {
            case 128:
                var v114 int32 = 0
                _ = v114
            }
            v55 = v55
        }
        fmt.Println(fi113, fi113)
        _ = fi113
    }
    v115 := GenS8{F0: (0 != 90), F1: v4}
    var v116 uint32 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
}
