package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 bool; F1 int32 }
type GenS8 struct { F0 uint64; F1 int; F2 uint }
type GenS9 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS10 struct { F0 bool; F1 uint32; F2 int8 }
func genFn0(p0 uint32, p1 uint32) bool {
    v7 := []int{32768, -97, -2147483648, -1}
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS1{F0: 600, F1: 0}
            _ = fi9
            _ = v10
        }
        p0 += 42
        p0 -= 4294967295
        _ = fi8
    }
    var v11 int16 = 824
    var v12 uint32 = (^(p0))
    var v13 float64 = float64((-2.25 - (0.0 * float32(1.5))))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v7
    return false
}
func genFn1(p0 int64) uint16 {
    var v14 float64 = 3.14159
    v15 := GenS2{F0: 18446744073709551615, F1: 1}
    p0 = 0
    _ = p0
    _ = v14
    _ = v15
    return 255
}

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 447
    var v3 int = -564
    v4 := make(map[string]int)
    v5 := make(map[uint32]string)
    var v6 int32 = -116
    v6 += v6
    var v16 float32 = 1.5
    if (!false) {
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 int32 = -1
            fmt.Println(v2, fi17, fi17)
            v19 := GenS3{F0: "a"}
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        v20 := GenS4{F0: v2, F1: 2}
        var v21 uint32 = (v2 * 2)
        var v22 uint = 42
        _ = v20
        _ = v21
        _ = v22
    }
    if true {
        switch (v3 ^ v3) {
        case 2147483647:
            v16 -= v16
        case 941:
            var v23 bool = (("hello" + ("foo bar" + ("" + ("foo bar" + ((("\n\t" + "foo bar") + "hello") + ("foo bar" + ("" + (((((("foo bar" + "foo bar") + ("\n\t" + "foo bar")) + "\n\t") + "foo bar") + "foo bar") + (((("a" + "foo bar") + "") + "") + "foo bar"))))))))) <= ("" + ""))
            _ = v23
        case 32767:
            for fi24 := 0; fi24 < 9; fi24++ {
                if (950 < 42) {
                    v16 = (v16 * v16)
                    var v25 int = (^(fi24))
                    switch int(v6) {
                    case -9223372036854775808:
                        fmt.Println(fi24, v2, fi24)
                    case 0:
                        var v26 uint64 = 934
                        _ = v26
                    default:
                        var v27 int = (fi24 % (((v3 - v3)) | 1))
                        _ = v27
                    }
                    _ = v25
                } else {
                    var v28 int8 = -128
                    var v29 uint8 = 42
                    v3 += fi24
                    _ = v28
                    _ = v29
                }
                _ = fi24
            }
        }
    }
    var v30 int32 = -2147483648
    for fi31 := 0; fi31 < 16; fi31++ {
        fmt.Println(v30, v3)
        for fi32 := 0; fi32 < 15; fi32++ {
            v33 := []int{-9223372036854775808, 42, 4294967295, (fi31 >> ((18446744073709551615) & 63))}
            switch (("\n\t" + "\n\t") + ("foo bar" + "a")) {
            case "":
                if true {
                    var v34 float64 = 0.1
                    var v35 uint8 = 255
                    _ = v34
                    _ = v35
                } else {
                    switch "foo bar" {
                    case "hello":
                        if (42 < 516) {
                            var v36 int64 = 42
                            _ = v36
                        } else {
                            if (v2 <= 22) {
                                if (677 < 32767) {
                                    v37 := GenS1{F0: 263, F1: genFn1(0)}
                                    v33 = append(v33, 887)
                                    _ = v37
                                }
                                var v38 uint16 = 1
                                _ = len(v5)
                                _ = v38
                            }
                            if ((fi32 - (fi31 % ((fi32) | 1))) < 1) {
                                var v39 bool = (540 >= 876)
                                fmt.Println(v3, v6)
                                v6 -= (v6 & -2147483648)
                                _ = v39
                            }
                        }
                    }
                    v33 = append(v33, fi32)
                    var v40 uint64 = 63
                    _ = v40
                }
            case "hello":
                switch -9223372036854775808 {
                case 1:
                    if (4294967295 > (fi32 & v3)) {
                        if false {
                            switch (^(fi32)) {
                            case 32767:
                                _ = len(v33)
                            case -981:
                                var v41 bool = true
                                _ = v41
                            case 127:
                                var v42 bool = true
                                v16 += (0.1 - 1e10)
                                _ = len(v33)
                                var v43 uint32 = 2147483648
                                _ = v42
                                _ = v43
                            default:
                                v44 := GenS3{F0: "hello"}
                                fmt.Println(v6)
                                var v45 uint64 = 18446744073709551615
                                _ = v44
                                _ = v45
                            }
                            var v46 int64 = -675
                            var v47 bool = (206 < 504)
                            _ = v46
                            _ = v47
                        }
                        v16 = v16
                    }
                default:
                    switch (("a" + "") + (("foo bar" + "foo bar") + "foo bar")) {
                    case "hello":
                        if true {
                            _ = len(v1)
                        }
                        if len(v33) >= 2 {
                            v33 = v33[1:len(v33)]
                        }
                    case "foo bar":
                        var v48 int = (-(v3))
                        _ = v48
                    case "\n\t":
                        var v49 bool = false
                        _ = v49
                    default:
                        if true {
                            v3 = (v3 * (v3 * 539))
                            v50 := GenS2{F0: 786, F1: 127}
                            if true {
                                if (-2.25 < (-2.25 + ((3.14159 - float64(3.14159)) + 3.14159))) {
                                    if true {
                                        v6 = v6
                                        switch "a" {
                                        case "foo bar":
                                            var v51 int16 = 779
                                            _ = v51
                                        case "a":
                                            v30 = (^(v6))
                                        case "hello":
                                            if (v2 < v2) {
                                                if true {
                                                    v3 = fi31
                                                    fmt.Println(v2)
                                                }
                                                switch 128 {
                                                case -1:
                                                    if len(v33) >= 2 {
                                                        v33 = v33[1:len(v33)]
                                                    }
                                                case -9223372036854775808:
                                                    fmt.Println(v3)
                                                    v52 := GenS1{F0: 42, F1: 438}
                                                    _ = v52
                                                case -930:
                                                    var v53 uint64 = 0
                                                    _ = v53
                                                default:
                                                    v54 := GenS2{F0: 0, F1: 156}
                                                    _ = v54
                                                }
                                            } else {
                                                if true {
                                                    var v55 int = (fi31 >> ((426) & 63))
                                                    var v56 string = "a"
                                                    var v57 string = v56
                                                    switch "a" {
                                                    case "":
                                                        delete(v1, 1)
                                                    case "\n\t":
                                                        var v58 int8 = 127
                                                        _ = v58
                                                    }
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                }
                                                if (-335 > 32767) {
                                                    switch (^(v3)) {
                                                    case 332:
                                                        v3 += 32767
                                                    case -364:
                                                        switch (fi31 % (((^(fi31))) | 1)) {
                                                        case 829:
                                                            v6 = v30
                                                        case 1:
                                                            switch "\n\t" {
                                                            case "":
                                                                v16 = v16
                                                            case "hello":
                                                                v59 := GenS1{F0: 164, F1: 65535}
                                                                _ = v59
                                                            }
                                                            var v60 float64 = -2.25
                                                            v30 -= (v6 / ((v30) | 1))
                                                            _ = v60
                                                        }
                                                    case 32768:
                                                        v6 += v30
                                                    }
                                                    if true {
                                                        var v61 float64 = 0.1
                                                        v3 += fi31
                                                        _ = len(v33)
                                                        var v62 float32 = ((float32(float64((v16 + (v16 / 1e10)))) * float32(0.1)) * 0.0)
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v63 := GenS2{F0: 1, F1: 2}
                                                            v62 = -2.25
                                                            switch "a" {
                                                            case "a":
                                                                switch ("" + ("hello" + "\n\t")) {
                                                                case "a":
                                                                    if (1 >= 0) {
                                                                        if (!true) {
                                                                            v64 := GenS1{F0: 255, F1: 504}
                                                                            v16 = ((v16 - (1.5 - v62)) + (((-2.25 - (1.5 / v62)) / (v62 * 0.0)) * v16))
                                                                            _ = v64
                                                                        }
                                                                        v4["hello"] = v3
                                                                        v65 := GenS5{F0: 127}
                                                                        var v66 float64 = 1e10
                                                                        fmt.Println(fi32, v6, v2)
                                                                        _ = v65
                                                                        _ = v66
                                                                    }
                                                                case "hello":
                                                                    if false {
                                                                        fmt.Println(v61, v16, v2)
                                                                        fmt.Println(v2)
                                                                        v16 += float32(float64(float32((v61 / 0.1))))
                                                                        v61 = float64(1.5)
                                                                    } else {
                                                                        switch fi31 {
                                                                        case 2147483647:
                                                                            if ((4294967295 != 18446744073709551615) || (!false)) {
                                                                                var v67 float32 = 1e10
                                                                                _ = v67
                                                                            } else {
                                                                                v68 := GenS3{F0: (("\n\t" + "hello") + "\n\t")}
                                                                                if true {
                                                                                    var v69 int = int(v6)
                                                                                    _ = v69
                                                                                } else {
                                                                                    v16 += (1e10 - v16)
                                                                                    var v70 float64 = (1e10 / (v61 + ((v61 / v61) / v61)))
                                                                                    v61 = (v61 * float64(float32(v70)))
                                                                                    _ = v70
                                                                                }
                                                                                _ = v68
                                                                            }
                                                                        default:
                                                                            v71 := GenS6{F0: 823, F1: 1}
                                                                            _ = v71
                                                                        }
                                                                        v72 := GenS1{F0: 0, F1: genFn1(9223372036854775807)}
                                                                        _ = v72
                                                                    }
                                                                case "":
                                                                    fmt.Println(v61, v62, v16)
                                                                }
                                                            }
                                                            _ = v63
                                                        case "":
                                                            var v73 uint = 1
                                                            _ = v73
                                                        case "\n\t":
                                                            fmt.Println(v2, v16)
                                                        default:
                                                            switch fi31 {
                                                            case -1:
                                                                var v74 int = v3
                                                                v4["a"] = v3
                                                                _ = v74
                                                            case 1:
                                                                var v75 string = "foo bar"
                                                                _ = v75
                                                            default:
                                                                if len(v33) >= 2 {
                                                                    v33 = v33[1:len(v33)]
                                                                }
                                                            }
                                                        }
                                                        var v76 uint = 980
                                                        _ = v61
                                                        _ = v62
                                                        _ = v76
                                                    }
                                                } else {
                                                    _ = len(v33)
                                                    var v77 bool = false
                                                    v3 = -2147483648
                                                    switch (fi32 / ((fi31) | 1)) {
                                                    case 32767:
                                                        switch 127 {
                                                        case 872:
                                                            v16 = -2.25
                                                        case 127:
                                                            v78 := GenS7{F0: (991 == 2), F1: v6}
                                                            _ = v78
                                                        }
                                                        v2 += (v2 & 42)
                                                    case 9223372036854775807:
                                                        var v79 int32 = (v30 ^ v6)
                                                        _ = v79
                                                    }
                                                    if v77 {
                                                        v30 = 0
                                                        if ("a" >= "hello") {
                                                            v33 = append(v33, (v3 | fi31))
                                                        }
                                                        if (324 > -1) {
                                                            if v77 {
                                                                v80 := GenS1{F0: 42, F1: 0}
                                                                _ = v80
                                                            } else {
                                                                var v81 bool = v77
                                                                v2 = 4294967295
                                                                if (v81 && v81) {
                                                                    v82 := GenS1{F0: 255, F1: 42}
                                                                    if ((-(v6)) > v30) {
                                                                        switch ("foo bar" + ("" + "\n\t")) {
                                                                        case "a":
                                                                            v83 := GenS3{F0: ("\n\t" + "hello")}
                                                                            v77 = ((true || v81) && (v2 > v2))
                                                                            switch -1 {
                                                                            case 138:
                                                                                var v84 uint16 = 65535
                                                                                var v85 uint64 = 1
                                                                                _ = v84
                                                                                _ = v85
                                                                            default:
                                                                                if len(v33) >= 2 {
                                                                                    v33 = v33[1:len(v33)]
                                                                                }
                                                                            }
                                                                            _ = v83
                                                                        default:
                                                                            v77 = false
                                                                        }
                                                                        fmt.Println(v77, v81)
                                                                    } else {
                                                                        var v86 int = (^(fi32))
                                                                        _ = v86
                                                                    }
                                                                    v2 = (v2 & (v2 & 2))
                                                                    _ = v82
                                                                }
                                                                _ = v81
                                                            }
                                                        } else {
                                                            var v87 uint64 = 1
                                                            if v77 {
                                                                var v88 bool = v77
                                                                v77 = (false || ((v30 | -1) > 42))
                                                                v89 := GenS8{F0: (v87 << ((4294967295) & 63)), F1: fi32, F2: 0}
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                            var v90 int = 462
                                                            _ = v87
                                                            _ = v90
                                                        }
                                                    } else {
                                                        v91 := GenS4{F0: v2, F1: 18446744073709551615}
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            if v77 {
                                                                v92 := GenS1{F0: 386, F1: 32767}
                                                                _ = v92
                                                            }
                                                        case "foo bar":
                                                            var v93 uint32 = (^(v2))
                                                            _ = v93
                                                        case "a":
                                                            v94 := GenS8{F0: 18446744073709551615, F1: (fi31 >> ((uint(v2)) & 63)), F2: 18446744073709551615}
                                                            _ = v94
                                                        }
                                                        _ = v91
                                                    }
                                                    var v95 float32 = v16
                                                    v30 -= (-(v6))
                                                    _ = v77
                                                    _ = v95
                                                }
                                                if false {
                                                    var v96 bool = true
                                                    var v97 uint32 = (^(v2))
                                                    _ = v96
                                                    _ = v97
                                                } else {
                                                    switch "hello" {
                                                    case "":
                                                        if (((0.1 - (1e10 * -2.25)) >= 1.5) && (0 < genFn1(-9223372036854775808))) {
                                                            v30 = 990
                                                            fmt.Println(v16, v16, v6)
                                                            v6 += -1
                                                        }
                                                    case "\n\t":
                                                        switch "" {
                                                        case "a":
                                                            var v98 uint8 = 151
                                                            _ = v98
                                                        }
                                                    case "a":
                                                        v99 := GenS9{F0: 0, F1: (1 == 15), F2: uint64(v2)}
                                                        _ = v99
                                                    default:
                                                        if true {
                                                            if len(v33) >= 2 {
                                                                v33 = v33[1:len(v33)]
                                                            }
                                                            v2 += v2
                                                        } else {
                                                            v2 += (v2 * 2147483648)
                                                        }
                                                    }
                                                }
                                            }
                                        default:
                                            var v100 int64 = 972
                                            _ = v100
                                        }
                                        v2 = (v2 << ((2) & 63))
                                    } else {
                                        switch ("hello" + "") {
                                        case "a":
                                            if true {
                                                fmt.Println(v3, fi31)
                                            } else {
                                                var v101 int32 = (v30 & (v6 << ((2) & 63)))
                                                var v102 uint16 = 42
                                                _ = v101
                                                _ = v102
                                            }
                                        }
                                    }
                                }
                            }
                            _ = v50
                        } else {
                            fmt.Println(v30, v30)
                        }
                    }
                }
                var v103 int16 = 602
                var v104 uint8 = 0
                _ = v103
                _ = v104
            case "\n\t":
                if (v2 >= (v2 & v2)) {
                    fmt.Println(fi31, v6, fi32)
                    var v105 uint16 = 828
                    fmt.Println(v2, v2)
                    var v106 int64 = 9223372036854775807
                    switch v3 {
                    case 32767:
                        if (v106 == v106) {
                            v6 = (^(v30))
                        }
                        delete(v5, 864)
                    case -9223372036854775808:
                        v2 = (v2 / ((v2) | 1))
                    case 128:
                        v6 = (v30 - v30)
                    default:
                        if len(v33) >= 2 {
                            v33 = v33[1:len(v33)]
                        }
                    }
                    _ = v105
                    _ = v106
                } else {
                    switch fi31 {
                    case 4294967295:
                        v16 -= ((((float32(0.0) * (1e10 + v16)) + float32((((-2.25 * 1e10) + 0.1) + 0.1))) - 3.14159) * v16)
                    case 127:
                        v1[fi31] = fi31
                    case -716:
                        v33 = append(v33, v3)
                    default:
                        v30 = (v6 * (v30 - v30))
                    }
                    fmt.Println(fi32, v6, fi32)
                    if (267 == 255) {
                        if false {
                            fmt.Println(v6, v2, fi32)
                            v33 = append(v33, (v3 % (((fi31 & fi31)) | 1)))
                            v107 := v1[(fi32 ^ -118)]
                            _ = v107
                        }
                        v2 -= 42
                    } else {
                        v3 = (v3 * 591)
                    }
                }
            }
            _ = fi32
            _ = v33
        }
        _ = fi31
    }
    if false {
        v3 += v3
    }
    for fi108 := 0; fi108 < 6; fi108++ {
        _ = len(v4)
        _ = fi108
    }
    var v109 float32 = float32(1e10)
    for fi110 := 0; fi110 < 20; fi110++ {
        v2 += (v2 | v2)
        if ((!(!false)) && false) {
            v6 = (v30 & (v6 / ((-2147483648) | 1)))
        }
        _ = fi110
    }
    fmt.Println(v16, v109, v109)
    _ = len(v5)
    v111 := GenS10{F0: (1 <= 2), F1: 1, F2: 64}
    if (!(false || true)) {
        var v112 bool = (false || (-9223372036854775808 == 9223372036854775807))
        _ = v112
    }
    for fi113 := 0; fi113 < 16; fi113++ {
        var v114 int = fi113
        for fi115 := 0; fi115 < 13; fi115++ {
            var v116 float32 = 1.5
            _ = fi115
            _ = v116
        }
        _ = fi113
        _ = v114
    }
    var v117 string = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v109:", v109)
    fmt.Println("v117:", v117)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v109
    _ = v111
    _ = v117
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
