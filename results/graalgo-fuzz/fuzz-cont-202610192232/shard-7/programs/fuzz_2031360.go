package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 string }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int32; F1 bool; F2 int }
type GenS6 struct { F0 uint; F1 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS13 struct { F0 int16; F1 int; F2 int32 }
type GenS14 struct { F0 uint8; F1 uint8 }
type GenS15 struct { F0 int16; F1 int64; F2 uint64 }
func genFn0(p0 uint64) int {
    p0 += 0
    switch 791 {
    case 128:
        var v5 string = "foo bar"
        _ = v5
    case 127:
        var v6 int16 = -365
        _ = v6
    case 32768:
        var v7 bool = false
        _ = v7
    default:
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint16 = 65535
            switch -2147483648 {
            case 32767:
                var v10 int8 = 42
                _ = v10
            case -657:
                v11 := GenS1{F0: 87, F1: -3, F2: (^(v9))}
                _ = v11
            }
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    return 42
}
func genFn1() int {
    if true {
        var v12 int16 = -568
        v13 := GenS2{F0: 1, F1: 72, F2: "hello"}
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := GenS1{F0: 0, F1: -1, F2: 867}
        var v16 uint16 = 32767
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := GenS2{F0: -243, F1: 292, F2: ""}
    var v18 uint16 = 0
    var v19 uint16 = (v18 / ((v18) | 1))
    _ = v17
    _ = v18
    _ = v19
    return 127
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]uint32)
    v3 := []int{191, -2147483648}
    var v4 float32 = float32(float64(((0.1 / ((-2.25 * (1.5 / 1e10)) * 3.14159)) / ((float32(float64(0.0)) - 0.1) - float32(1.5)))))
    var v20 uint16 = 42
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 uint16 = 42
        for fi23 := 0; fi23 < 15; fi23++ {
            fmt.Println(fi23, v22)
            var v24 int64 = -1
            var v25 int64 = 882
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 string = "a"
        _ = fi21
        _ = v22
        _ = v26
    }
    v4 += 0.0
    switch 2147483647 {
    case 0:
        if ((!false) && (32767 > 65535)) {
            var v27 string = "a"
            _ = v27
        } else {
            v1 = v1
            v28 := v2[4294967295]
            _ = v28
        }
    case 2147483647:
        switch -200 {
        case 2147483647:
            v1 = (v1 % ((v1) | 1))
        case 32768:
            fmt.Println(v1)
        }
    case -9223372036854775808:
        v1 -= v1
    default:
        var v29 int32 = (v1 * v1)
        _ = v29
    }
    if true {
        fmt.Println(v20, v1, v20)
    }
    switch ("" + (("\n\t" + "foo bar") + (((("hello" + "a") + ((("a" + "foo bar") + "") + ("hello" + "hello"))) + "\n\t") + "\n\t"))) {
    case "foo bar":
        v1 = (v1 & -531)
    case "\n\t":
        for fi30 := 0; fi30 < 3; fi30++ {
            for fi31 := 0; fi31 < 20; fi31++ {
                v4 += v4
                v20 = (v20 / (((^(v20))) | 1))
                fmt.Println(v20, fi30)
                _ = fi31
            }
            _ = fi30
        }
    case "":
        if false {
            v32 := GenS2{F0: (v1 & v1), F1: 4294967295, F2: (("foo bar" + "a") + ("hello" + "\n\t"))}
            for fi33 := 0; fi33 < 10; fi33++ {
                delete(v2, (^(fi33)))
                _ = fi33
            }
            _ = v32
        }
    }
    var v34 uint64 = 2
    if false {
        _ = len(v3)
    }
    if (!false) {
        for fi35 := 0; fi35 < 17; fi35++ {
            if (42 > -9223372036854775808) {
                var v36 int = genFn1()
                var v37 uint8 = uint8(v20)
                v38 := GenS1{F0: 2147483648, F1: 42, F2: (v20 * (v20 | (v20 ^ v20)))}
                _ = v36
                _ = v37
                _ = v38
            }
            for fi39 := 0; fi39 < 15; fi39++ {
                var v40 uint = 42
                _ = fi39
                _ = v40
            }
            _ = fi35
        }
        v41 := GenS3{F0: 9223372036854775807, F1: (-(v1))}
        if (v20 != v20) {
            switch "" {
            case "\n\t":
                v42 := GenS2{F0: (v1 ^ v1), F1: (v34 >> ((63) & 63)), F2: "foo bar"}
                _ = v42
            case "":
                switch 2147483647 {
                case 628:
                    v4 -= ((v4 * float32(0.1)) / (1.5 - (0.1 - ((v4 - v4) + 0.1))))
                case 32768:
                    v34 = (^(v34))
                }
            default:
                var v43 string = "\n\t"
                _ = v43
            }
            fmt.Println(v20, v4, v4)
            switch -1 {
            case 1:
                v20 -= v20
            case -1:
                if false {
                    var v44 int64 = 42
                    if (v20 <= uint16(v34)) {
                        v45 := GenS2{F0: (v1 | v1), F1: (v34 / (((v34 / (((v34 - v34)) | 1))) | 1)), F2: "hello"}
                        _ = v45
                    }
                    _ = v44
                }
            default:
                for fi46 := 0; fi46 < 15; fi46++ {
                    v34 -= (v34 - (v34 >> ((1) & 63)))
                    v2[0] = 4294967295
                    fmt.Println(fi46)
                    _ = fi46
                }
            }
        } else {
            switch 0 {
            case -5:
                v2[-2147483648] = 42
            case 128:
                var v47 float64 = ((1e10 * float64(v4)) + float64((((float32(float64(0.1)) + v4) + 0.0) - v4)))
                _ = v47
            case 728:
                fmt.Println(v20, v34)
            default:
                v48 := v2[-9223372036854775808]
                _ = v48
            }
            if false {
                switch ("a" + "") {
                case "\n\t":
                    for fi49 := 0; fi49 < 14; fi49++ {
                        v4 = v4
                        if true {
                            if false {
                                v50 := GenS1{F0: 2147483648, F1: 100, F2: (v20 * v20)}
                                _ = v50
                            }
                            var v51 uint32 = 42
                            if (-1 >= 9223372036854775807) {
                                if false {
                                    var v52 int64 = -973
                                    switch 9223372036854775807 {
                                    case 2147483647:
                                        if (!true) {
                                            for fi53 := 0; fi53 < 9; fi53++ {
                                                var v54 float32 = v4
                                                v1 += (v1 ^ (v1 << ((453) & 63)))
                                                v4 += 0.0
                                                _ = fi53
                                                _ = v54
                                            }
                                            fmt.Println(v1, fi49, v34)
                                        }
                                    default:
                                        v55 := GenS4{F0: v1, F1: (^(v20))}
                                        _ = v55
                                    }
                                    v56 := GenS5{F0: (-(v1)), F1: (990 <= -440), F2: fi49}
                                    _ = v52
                                    _ = v56
                                } else {
                                    if (!((v34 + (v34 << ((42) & 63))) >= v34)) {
                                        switch "a" {
                                        case "foo bar":
                                            switch fi49 {
                                            case 128:
                                                if false {
                                                    var v57 float32 = -2.25
                                                    if ("hello" >= "foo bar") {
                                                        var v58 int16 = -1
                                                        if false {
                                                            switch (fi49 >> ((42) & 63)) {
                                                            case -9223372036854775808:
                                                                fmt.Println(fi49, v1)
                                                            case 0:
                                                                v59 := GenS6{F0: 255, F1: 255}
                                                                _ = v59
                                                            default:
                                                                v1 = 1
                                                            }
                                                        }
                                                        _ = v58
                                                    }
                                                    _ = v57
                                                }
                                            case 0:
                                                v34 = v34
                                            case 693:
                                                switch "foo bar" {
                                                case "foo bar":
                                                    fmt.Println(fi49, v20)
                                                case "":
                                                    for fi60 := 0; fi60 < 20; fi60++ {
                                                        fmt.Println(v51, v34, v34)
                                                        _ = fi60
                                                    }
                                                }
                                            }
                                        case "hello":
                                            if true {
                                                if (!(true && ((-2.25 * (1e10 + 1.5)) < 1e10))) {
                                                    if false {
                                                        v61 := GenS7{F0: -128}
                                                        v1 = (v1 * (v1 / ((v1) | 1)))
                                                        delete(v2, (fi49 * int(v1)))
                                                        _ = v61
                                                    }
                                                } else {
                                                    v20 = (v20 % ((65535) | 1))
                                                    var v62 int16 = 0
                                                    _ = v62
                                                }
                                                var v63 uint = 0
                                                if ((v1 - 0) != v1) {
                                                    v20 -= (v20 >> ((v63) & 63))
                                                }
                                                _ = v63
                                            } else {
                                                var v64 int = fi49
                                                delete(v2, fi49)
                                                _ = v64
                                            }
                                        case "":
                                            v65 := GenS1{F0: v51, F1: -18, F2: (v20 * (v20 * v20))}
                                            _ = v65
                                        default:
                                            switch fi49 {
                                            case -808:
                                                v1 += -579
                                            }
                                        }
                                    }
                                    v51 -= v51
                                    switch ("" + "a") {
                                    case "":
                                        if (!false) {
                                            v1 = (-(v1))
                                        }
                                    case "\n\t":
                                        var v66 uint32 = v51
                                        _ = v66
                                    }
                                }
                            }
                            _ = v51
                        } else {
                            for fi67 := 0; fi67 < 17; fi67++ {
                                if false {
                                    switch "" {
                                    case "a":
                                        v3 = append(v3, int(v34))
                                    default:
                                        v68 := GenS7{F0: -128}
                                        _ = v68
                                    }
                                    if (!(((((1e10 * -2.25) - (float64(1.5) - ((1e10 - (1e10 / -2.25)) * 1.5))) * -2.25) - (1.5 / 1e10)) < 3.14159)) {
                                        v69 := GenS3{F0: (^(fi49)), F1: (v1 & (v1 << ((65535) & 63)))}
                                        switch "\n\t" {
                                        case "":
                                            v70 := GenS8{F0: v1, F1: genFn1()}
                                            _ = v70
                                        case "foo bar":
                                            v1 = (v1 - 602)
                                        }
                                        _ = v69
                                    } else {
                                        _ = len(v2)
                                        v3 = append(v3, (fi49 % (((fi67 - (fi49 / (((fi67 >> ((999) & 63))) | 1)))) | 1)))
                                    }
                                } else {
                                    var v71 int16 = -1
                                    if false {
                                        v72 := GenS9{F0: (false && false)}
                                        if (0 > 4294967295) {
                                            v4 -= (((v4 / 1.0) / (0.1 - 3.14159)) + (((v4 * (((1.5 * (0.0 / ((v4 * v4) / ((v4 / (0.0 - 0.1)) + float32((-2.25 / 0.1)))))) - float32(-2.25)) / ((v4 * v4) / v4))) * v4) - 0.0))
                                        } else {
                                            v1 += (v1 - (v1 & v1))
                                            _ = len(v3)
                                        }
                                        switch 907 {
                                        case 595:
                                            v73 := GenS10{F0: (fi67 - fi67)}
                                            _ = v73
                                        default:
                                            if true {
                                                v74 := GenS11{F0: "a"}
                                                _ = v74
                                            } else {
                                                var v75 uint64 = 0
                                                _ = len(v3)
                                                if (true && true) {
                                                    if false {
                                                        var v76 int32 = v1
                                                        switch ("\n\t" + (("" + "\n\t") + (("foo bar" + "\n\t") + "hello"))) {
                                                        case "":
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        default:
                                                            _ = len(v3)
                                                        }
                                                        var v77 int8 = -29
                                                        _ = v76
                                                        _ = v77
                                                    } else {
                                                        fmt.Println(v75, v4, fi49)
                                                    }
                                                    var v78 uint16 = v20
                                                    if len(v3) > 0 {
                                                        v3[0] = (^(fi49))
                                                    }
                                                    _ = v78
                                                } else {
                                                    if ((!(true || true)) && ((v75 - v75) == (v75 ^ (v75 + v34)))) {
                                                        if true {
                                                            if true {
                                                                v3 = append(v3, 0)
                                                                if ((^(v75)) <= v34) {
                                                                    var v79 float64 = (3.14159 - (0.1 / 3.14159))
                                                                    _ = v79
                                                                } else {
                                                                    v80 := GenS12{F0: v20, F1: (genFn1() == (fi49 * 9223372036854775807)), F2: 6}
                                                                    fmt.Println(v34, v4)
                                                                    fmt.Println(v34, fi49, v75)
                                                                    _ = v80
                                                                }
                                                            } else {
                                                                var v81 int64 = -667
                                                                var v82 int32 = (v1 / ((374) | 1))
                                                                _ = v81
                                                                _ = v82
                                                            }
                                                            if (v34 >= v75) {
                                                                var v83 uint8 = 2
                                                                _ = v83
                                                            }
                                                            v84 := GenS13{F0: 624, F1: fi67, F2: (^(v1))}
                                                            _ = v84
                                                        }
                                                    }
                                                    switch fi67 {
                                                    case 106:
                                                        delete(v2, genFn1())
                                                    case 2147483647:
                                                        var v85 uint32 = 534
                                                        _ = v85
                                                    case -1:
                                                        var v86 int64 = 1
                                                        _ = v86
                                                    default:
                                                        v87 := GenS14{F0: 0, F1: 160}
                                                        _ = v87
                                                    }
                                                }
                                                _ = v75
                                            }
                                        }
                                        _ = v72
                                    }
                                    var v88 float32 = ((1e10 - float32(3.14159)) / ((((0.1 + (v4 * (3.14159 / 1e10))) + v4) / v4) - float32(3.14159)))
                                    _ = v71
                                    _ = v88
                                }
                                var v89 uint = 42
                                v34 -= v34
                                _ = fi67
                                _ = v89
                            }
                        }
                        _ = fi49
                    }
                }
            }
        }
        _ = v41
    }
    v90 := GenS14{F0: 255, F1: 127}
    if true {
        if false {
            v2[0] = 42
            v91 := GenS15{F0: 1000, F1: -9223372036854775808, F2: 2}
            _ = v91
        } else {
            v1 += v1
            fmt.Println(v4, v4)
        }
        if true {
            if (1 >= 20) {
                if (uint(v20) < 42) {
                    v92 := GenS13{F0: 1, F1: -928, F2: v1}
                    for fi93 := 0; fi93 < 18; fi93++ {
                        v20 += v20
                        v94 := GenS14{F0: 214, F1: 127}
                        v34 -= 905
                        _ = fi93
                        _ = v94
                    }
                    _ = v92
                }
                v2[4294967295] = uint32(v34)
            }
            for fi95 := 0; fi95 < 19; fi95++ {
                if (true && (v20 != 32767)) {
                    switch "foo bar" {
                    case "\n\t":
                        for fi96 := 0; fi96 < 14; fi96++ {
                            v3 = append(v3, fi95)
                            var v97 float64 = -2.25
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi96
                            _ = v97
                        }
                    default:
                        fmt.Println(v20, v4)
                    }
                }
                _ = fi95
            }
        }
        var v98 int64 = 415
        _ = v98
    } else {
        v4 = float32((((1e10 * 3.14159) / 1e10) + 0.0))
        var v99 int = 1
        if true {
            var v100 float64 = ((1e10 / 1.0) + float64(float32((float64(v4) - float64(((v4 + 1e10) + v4))))))
            v3 = append(v3, (v99 + v99))
            if len(v3) > 0 {
                v3[0] = v99
            }
            _ = v100
        }
        _ = v99
    }
    for fi101 := 0; fi101 < 1; fi101++ {
        for fi102 := 0; fi102 < 11; fi102++ {
            var v103 int32 = v1
            _ = fi102
            _ = v103
        }
        for fi104 := 0; fi104 < 1; fi104++ {
            delete(v2, fi101)
            var v105 string = ""
            _ = fi104
            _ = v105
        }
        _ = fi101
    }
    v20 -= v20
    for fi106 := 0; fi106 < 11; fi106++ {
        v1 += v1
        v34 = 441
        var v107 uint32 = 420
        _ = fi106
        _ = v107
    }
    switch "foo bar" {
    case "hello":
        if true {
            for fi108 := 0; fi108 < 11; fi108++ {
                for fi109 := 0; fi109 < 11; fi109++ {
                    if (v1 > -94) {
                        v110 := GenS1{F0: 65535, F1: -106, F2: (v20 / ((65535) | 1))}
                        v4 = ((float32(3.14159) + (1.5 * 0.0)) * v4)
                        _ = v110
                    } else {
                        v3 = append(v3, -44)
                    }
                    _ = fi109
                }
                v3 = append(v3, fi108)
                if len(v3) > 0 {
                    v3[0] = (fi108 - (fi108 * (fi108 + genFn0((v34 ^ (v34 * uint64(v20)))))))
                }
                _ = fi108
            }
        }
    case "a":
        var v111 int16 = -32768
        _ = v111
    case "":
        for fi112 := 0; fi112 < 7; fi112++ {
            for fi113 := 0; fi113 < 20; fi113++ {
                v34 = (^(v34))
                _ = fi113
            }
            _ = fi112
        }
    }
    _ = len(v2)
    var v114 int64 = -299
    fmt.Println("v1:", v1)
    fmt.Println("v114:", v114)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v114
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v4
    _ = v90
}
