package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 string; F2 int32 }
type GenS4 struct { F0 int16; F1 uint16; F2 int16 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 int16; F1 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int32; F1 int32; F2 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint, p1 uint64) uint32 {
    v7 := []int{}
    var v8 uint = p0
    v9 := GenS1{F0: p1, F1: 65535, F2: 401}
    for fi10 := 0; fi10 < 9; fi10++ {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 2147483648
}
func genFn1() int {
    switch "" {
    case "\n\t":
        for fi11 := 0; fi11 < 7; fi11++ {
            if (((0.0 + 1e10) / 1e10) > ((1.5 / 0.1) * ((1e10 * (0.0 - (((3.14159 - (1e10 * ((((-2.25 * 1e10) / 1e10) - 3.14159) * 3.14159))) / ((3.14159 / (3.14159 + 1.5)) + 1.5)) * (3.14159 / float64(3.14159))))) + 0.1))) {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v13 := GenS1{F0: 2, F1: 2, F2: 4294967295}
                    _ = fi12
                    _ = v13
                }
                var v14 uint64 = 2
                _ = v14
            } else {
                if true {
                    var v15 string = ("" + "a")
                    _ = v15
                } else {
                    var v16 float64 = (3.14159 * ((float64(3.14159) * 0.0) + 1.5))
                    switch "\n\t" {
                    case "hello":
                        switch ("a" + ("" + "hello")) {
                        case "foo bar":
                            fmt.Println(fi11)
                        }
                    }
                    v16 = v16
                    _ = v16
                }
                var v17 uint8 = 15
                var v18 float64 = -2.25
                _ = v17
                _ = v18
            }
            _ = fi11
        }
    case "a":
        if false {
            v19 := GenS1{F0: 42, F1: 127, F2: 807}
            _ = v19
        } else {
            var v20 int = 4294967295
            v20 -= v20
            _ = v20
        }
    }
    if (42 <= 42) {
        v21 := []int{32768, 42, -469, -586}
        var v22 uint64 = 18446744073709551615
        var v23 int8 = -66
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 int8 = 100
    v25 := []int{}
    _ = v24
    _ = v25
    return 32768
}

func main() {
    var v1 uint = 1
    var v2 uint32 = 84
    var v3 uint32 = v2
    v4 := []string{"foo bar", "hello", "\n\t", ("" + (("a" + ("hello" + "hello")) + ("" + "a")))}
    var v5 int = int(v3)
    var v6 uint16 = uint16(v3)
    v6 += 96
    fmt.Println(v5, v6, v5)
    if len(v4) > 0 {
        v4[0] = "a"
    }
    var v26 float64 = 0.0
    var v27 string = ""
    switch (("\n\t" + v27) + v27) {
    case "hello":
        for fi28 := 0; fi28 < 1; fi28++ {
            if len(v4) > 0 {
                v4[0] = "a"
            }
            v1 += 2
            _ = fi28
        }
    case "\n\t":
        switch v27 {
        case "\n\t":
            v5 -= (v5 + -2147483648)
        case "foo bar":
            if ((v1 | (^(v1))) < (v1 >> ((2) & 63))) {
                var v29 uint64 = 42
                if len(v4) > 0 {
                    v4[0] = v27
                }
                var v30 int64 = 776
                _ = v29
                _ = v30
            } else {
                switch (v5 * (v5 - (v5 / ((2147483647) | 1)))) {
                case 4294967295:
                    v31 := GenS2{F0: v5}
                    _ = v31
                case 832:
                    v5 -= v5
                default:
                    v3 = genFn0(v1, 1)
                }
            }
        case "":
            if true {
                switch v5 {
                case 127:
                    fmt.Println(v27, v27)
                case -355:
                    var v32 int16 = 1000
                    _ = v32
                }
            }
        }
    case "":
        v33 := GenS3{F0: v27, F1: v27, F2: 0}
        _ = v33
    default:
        v34 := GenS2{F0: 0}
        _ = v34
    }
    if (!(-1 > -651)) {
        v5 -= genFn1()
        switch ((v27 + "hello") + (((v27 + (v27 + v27)) + (("foo bar" + v27) + (v27 + v27))) + "hello")) {
        case "":
            v35 := GenS4{F0: 32767, F1: (v6 ^ uint16(v1)), F2: 42}
            _ = v35
        case "foo bar":
            for fi36 := 0; fi36 < 2; fi36++ {
                if false {
                    var v37 int64 = -1
                    _ = v37
                } else {
                    for fi38 := 0; fi38 < 8; fi38++ {
                        if (!((1e10 + float32(float64(-2.25))) >= (((1e10 + 0.0) - float32(1.5)) / float32(v26)))) {
                            v39 := GenS5{F0: ((v1 / ((651) | 1)) != 407), F1: false}
                            _ = v39
                        } else {
                            v27 = ("" + ((((v27 + v27) + v27) + v27) + ((v27 + ("\n\t" + "\n\t")) + v27)))
                        }
                        v40 := GenS6{F0: "\n\t", F1: (v6 + (^(v6)))}
                        v26 -= (0.1 - v26)
                        _ = fi38
                        _ = v40
                    }
                    var v41 uint32 = 843
                    _ = v41
                }
                _ = fi36
            }
        case "\n\t":
            _ = len(v4)
        }
    }
    var v42 float64 = 1e10
    var v43 int64 = 0
    if true {
        var v44 int16 = -210
        _ = v44
    } else {
        if ((^(v1)) <= 4294967295) {
            v45 := GenS1{F0: 642, F1: v6, F2: (v2 % ((2147483648) | 1))}
            var v46 int8 = 71
            _ = v45
            _ = v46
        } else {
            var v47 int16 = 32767
            _ = v47
        }
    }
    v26 = ((((v42 / v42) / v42) / 3.14159) - 0.1)
    for fi48 := 0; fi48 < 17; fi48++ {
        v49 := GenS5{F0: false, F1: (!false)}
        if false {
            if false {
                switch (^(v5)) {
                case -2147483648:
                    for fi50 := 0; fi50 < 5; fi50++ {
                        v51 := GenS5{F0: (!false), F1: (2 > 70)}
                        v4 = append(v4, "")
                        _ = fi50
                        _ = v51
                    }
                }
            }
            var v52 int32 = 2147483647
            if (3.14159 != 0.1) {
                if true {
                    v53 := GenS3{F0: ("\n\t" + v27), F1: "a", F2: (v52 ^ -911)}
                    _ = v53
                } else {
                    switch fi48 {
                    case 926:
                        v27 = v27
                    case 9223372036854775807:
                        switch v27 {
                        case "foo bar":
                            if false {
                                switch v5 {
                                case -9223372036854775808:
                                    var v54 int = v5
                                    _ = v54
                                default:
                                    for fi55 := 0; fi55 < 3; fi55++ {
                                        if ((v6 >> ((v1) & 63)) != (v6 ^ v6)) {
                                            switch ((((v27 + "\n\t") + v27) + v27) + "hello") {
                                            case "foo bar":
                                                if (v2 == v3) {
                                                    var v56 uint16 = v6
                                                    fmt.Println(v6, v43, v6)
                                                    var v57 int8 = -1
                                                    var v58 int64 = (v43 << ((18446744073709551615) & 63))
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                }
                                            case "\n\t":
                                                v43 = 42
                                            default:
                                                var v59 uint16 = v6
                                                _ = v59
                                            }
                                            var v60 string = v27
                                            if len(v4) > 0 {
                                                v4[0] = v60
                                            }
                                            _ = v60
                                        }
                                        if (-2.25 <= (3.14159 * 1.5)) {
                                            var v61 uint8 = uint8(v1)
                                            var v62 uint16 = v6
                                            _ = v61
                                            _ = v62
                                        }
                                        _ = fi55
                                    }
                                }
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = ((v27 + (v27 + (v27 + v27))) + "\n\t")
                                }
                            }
                        case "":
                            switch (fi48 | (fi48 | (fi48 % ((v5) | 1)))) {
                            case -2147483648:
                                if len(v4) > 0 {
                                    v4[0] = ((v27 + v27) + (("foo bar" + (v27 + v27)) + ("foo bar" + "a")))
                                }
                            case 631:
                                v63 := GenS1{F0: 18446744073709551615, F1: v6, F2: (v3 << ((2) & 63))}
                                _ = v63
                            default:
                                fmt.Println(v2, v3, v43)
                            }
                        default:
                            var v64 int16 = 32767
                            _ = v64
                        }
                    default:
                        if len(v4) > 0 {
                            v4[0] = v27
                        }
                    }
                    for fi65 := 0; fi65 < 5; fi65++ {
                        if true {
                            switch v27 {
                            case "":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case "\n\t":
                                if true {
                                    if false {
                                        switch genFn1() {
                                        case 987:
                                            fmt.Println(v2, v1)
                                        case 128:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case -9223372036854775808:
                                            if false {
                                                v66 := GenS7{F0: fi65, F1: 0}
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                var v67 float64 = v26
                                                var v68 string = v27
                                                _ = v66
                                                _ = v67
                                                _ = v68
                                            } else {
                                                v69 := GenS5{F0: false, F1: ((v43 >= -63) || (v42 >= v42))}
                                                switch v27 {
                                                case "a":
                                                    v2 = (v3 * 2147483648)
                                                case "\n\t":
                                                    v42 -= (v26 / 3.14159)
                                                    var v70 string = v27
                                                    _ = v70
                                                case "foo bar":
                                                    var v71 uint32 = v3
                                                    if false {
                                                        if ((v6 - (v6 + (v6 ^ v6))) < 627) {
                                                            fmt.Println(v42)
                                                        }
                                                        switch (v5 | (v5 % ((fi65) | 1))) {
                                                        case 9223372036854775807:
                                                            var v72 int64 = v43
                                                            _ = v72
                                                        default:
                                                            var v73 uint64 = 835
                                                            _ = v73
                                                        }
                                                    }
                                                    var v74 int32 = v52
                                                    _ = v71
                                                    _ = v74
                                                }
                                                _ = v69
                                            }
                                        }
                                        v75 := GenS8{F0: -1, F1: (v27 + ("\n\t" + (v27 + ((("a" + "hello") + "hello") + "foo bar"))))}
                                        v52 = -1
                                        v52 = int32(v43)
                                        _ = v75
                                    } else {
                                        if false {
                                            var v76 float64 = v26
                                            v1 += uint(v6)
                                            v1 += v1
                                            v77 := GenS9{F0: (v27 >= (v27 + v27))}
                                            _ = v76
                                            _ = v77
                                        }
                                        v78 := GenS8{F0: 42, F1: "a"}
                                        v2 -= (^(v3))
                                        _ = v78
                                    }
                                }
                            case "hello":
                                if (3.14159 >= float32(3.14159)) {
                                    _ = len(v4)
                                    var v79 int64 = (v43 >> ((v1) & 63))
                                    if (2 != 1) {
                                        var v80 uint8 = 127
                                        _ = v80
                                    }
                                    switch "a" {
                                    case "\n\t":
                                        if (!((v79 | (-(v43))) <= (-(v43)))) {
                                            if len(v4) > 0 {
                                                v4[0] = ""
                                            }
                                            var v81 int = v5
                                            if true {
                                                v26 -= v26
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                fmt.Println(v42, v43, v42)
                                                if (!true) {
                                                    var v82 int = 9223372036854775807
                                                    _ = len(v4)
                                                    _ = v82
                                                } else {
                                                    v83 := GenS10{F0: 98}
                                                    if true {
                                                        if true {
                                                            var v84 int16 = 32767
                                                            v85 := GenS9{F0: (v52 <= (v52 >> ((v1) & 63)))}
                                                            var v86 uint64 = 257
                                                            v42 -= 1e10
                                                            switch (fi65 >> ((v1) & 63)) {
                                                            case -2147483648:
                                                                fmt.Println(v26, fi48)
                                                            case -949:
                                                                var v87 int16 = v84
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                _ = v87
                                                            case -378:
                                                                var v88 float64 = -2.25
                                                                _ = v88
                                                            }
                                                            _ = v84
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                        v79 = (^(v43))
                                                    } else {
                                                        var v89 int64 = v79
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        _ = v89
                                                    }
                                                    v90 := GenS4{F0: 32767, F1: 177, F2: -32768}
                                                    _ = v83
                                                    _ = v90
                                                }
                                                v91 := GenS11{F0: v52, F1: (v52 / ((v52) | 1)), F2: (fi48 >> ((18446744073709551615) & 63))}
                                                _ = v91
                                            } else {
                                                v92 := GenS12{F0: v27}
                                                _ = v92
                                            }
                                            switch (fi48 % ((128) | 1)) {
                                            case 387:
                                                v3 = genFn0(255, 4294967295)
                                            case 418:
                                                var v93 bool = false
                                                _ = v93
                                            default:
                                                var v94 uint32 = (v3 * 65535)
                                                _ = v94
                                            }
                                            fmt.Println(v5, v3)
                                            _ = v81
                                        } else {
                                            switch int(v6) {
                                            case -83:
                                                v4 = append(v4, "a")
                                            case 9223372036854775807:
                                                v1 = (v1 & 18446744073709551615)
                                            case -1:
                                                if true {
                                                    switch genFn1() {
                                                    case 32767:
                                                        var v95 float32 = 1e10
                                                        _ = v95
                                                    default:
                                                        var v96 uint = (v1 | 42)
                                                        _ = v96
                                                    }
                                                    fmt.Println(v42, fi65, v2)
                                                    fmt.Println(v1)
                                                    var v97 int8 = 0
                                                    _ = v97
                                                } else {
                                                    v98 := GenS13{F0: false}
                                                    _ = v98
                                                }
                                            default:
                                                var v99 uint8 = 42
                                                _ = v99
                                            }
                                        }
                                    case "a":
                                        var v100 int64 = (v79 % ((v43) | 1))
                                        _ = v100
                                    }
                                    fmt.Println(v79, v43)
                                    if len(v4) > 0 {
                                        v4[0] = ("foo bar" + "a")
                                    }
                                    switch v27 {
                                    case "\n\t":
                                        if true {
                                            if true {
                                                var v101 float64 = (0.1 - float64(1.5))
                                                if false {
                                                    v26 += 3.14159
                                                }
                                                var v102 uint32 = genFn0((v1 >> ((v1) & 63)), 42)
                                                switch v27 {
                                                case "":
                                                    var v103 string = "a"
                                                    _ = v103
                                                }
                                                _ = v101
                                                _ = v102
                                            } else {
                                                var v104 string = v27
                                                v105 := GenS14{F0: 626, F1: 449}
                                                _ = v104
                                                _ = v105
                                            }
                                            v4 = append(v4, (v27 + "hello"))
                                        } else {
                                            if (!(int32(v79) == v52)) {
                                                if false {
                                                    var v106 uint8 = 2
                                                    var v107 int8 = 58
                                                    if false {
                                                        fmt.Println(fi65, v3)
                                                        var v108 uint64 = 18446744073709551615
                                                        v52 -= v52
                                                        _ = v108
                                                    }
                                                    _ = v106
                                                    _ = v107
                                                }
                                                var v109 int8 = 0
                                                _ = v109
                                            } else {
                                                fmt.Println(fi48)
                                                fmt.Println(v52)
                                                v27 = (v27 + v27)
                                                v110 := GenS4{F0: 32767, F1: (v6 + 42), F2: -1}
                                                _ = v110
                                            }
                                        }
                                        if ((v52 ^ (v52 * (v52 / ((42) | 1)))) >= 0) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v111 := GenS12{F0: v27}
                                            _ = v111
                                        }
                                    }
                                    _ = v79
                                } else {
                                    v27 = v27
                                    switch v27 {
                                    case "\n\t":
                                        v112 := GenS11{F0: v52, F1: (-(v52)), F2: fi65}
                                        _ = v112
                                    case "hello":
                                        v43 += v43
                                    case "foo bar":
                                        if (v2 >= (v3 ^ (^(v2)))) {
                                            v1 += (v1 + 952)
                                        } else {
                                            if false {
                                                var v113 float64 = v26
                                                switch -1 {
                                                case 32768:
                                                    if len(v4) > 0 {
                                                        v4[0] = "a"
                                                    }
                                                }
                                                _ = v113
                                            }
                                        }
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                }
                            default:
                                var v114 bool = true
                                _ = v114
                            }
                        } else {
                            v27 = v27
                        }
                        v5 += genFn1()
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        switch (^(v5)) {
                        case -858:
                            _ = len(v4)
                        case 128:
                            v42 += (v42 + 1.5)
                        default:
                            v42 += v42
                        }
                        if (((((v27 + (v27 + ("foo bar" + v27))) + v27) + "hello") + "hello") <= v27) {
                            var v115 uint = 414
                            _ = v115
                        } else {
                            var v116 int64 = 1
                            _ = v116
                        }
                        _ = fi65
                    }
                }
                for fi117 := 0; fi117 < 6; fi117++ {
                    fmt.Println(v43)
                    _ = fi117
                }
                switch v27 {
                case "\n\t":
                    fmt.Println(v3, v2, fi48)
                default:
                    fmt.Println(fi48, v1, v2)
                }
            }
            _ = v52
        }
        _ = fi48
        _ = v49
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    v43 = (v43 / (((v43 & v43)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
