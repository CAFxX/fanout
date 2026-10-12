package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int16; F2 uint32 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint8; F1 int32; F2 int }
func genFn0() int32 {
    var v7 bool = (42 == -87)
    var v8 int8 = -127
    if (float64((1.5 - 0.0)) > 0.0) {
        v8 = 1
    }
    _ = v7
    _ = v8
    return 2147483647
}
func genFn1() int8 {
    var v9 int16 = 1000
    v9 = (v9 % ((-32768) | 1))
    fmt.Println(v9)
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            fmt.Println(v9, fi10, fi11)
            _ = fi11
        }
        for fi12 := 0; fi12 < 8; fi12++ {
            switch (fi10 ^ (^(fi10))) {
            case 1:
                switch "" {
                case "foo bar":
                    v13 := GenS1{F0: 572, F1: -1}
                    _ = v13
                }
            case -2147483648:
                var v14 float64 = 3.14159
                _ = v14
            default:
                fmt.Println(fi10)
            }
            _ = fi12
        }
        _ = fi10
    }
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]string)
    v3 := []string{(("foo bar" + "\n\t") + "foo bar"), "\n\t"}
    var v4 string = "foo bar"
    v5 := make(map[string]int)
    var v6 int16 = 1000
    switch ("foo bar" + v4) {
    case "hello":
        var v15 uint = 546
        _ = v15
    case "\n\t":
        var v16 int32 = 1
        _ = v16
    case "foo bar":
        var v17 bool = true
        _ = v17
    default:
        if (42 == 42) {
            var v18 float32 = (-2.25 + (3.14159 + (((float32((1e10 + (-2.25 - 0.1))) / 1.0) / 0.1) * 3.14159)))
            v18 += 0.1
            var v19 int16 = (v6 >> ((1) & 63))
            _ = v18
            _ = v19
        }
    }
    if len(v3) > 0 {
        v3[0] = (("a" + (v4 + "")) + v4)
    }
    switch 686 {
    case -375:
        for fi20 := 0; fi20 < 11; fi20++ {
            if len(v3) > 0 {
                v3[0] = ""
            }
            _ = fi20
        }
    case -9223372036854775808:
        fmt.Println(v4, v4)
    default:
        var v21 uint64 = 18446744073709551615
        _ = v21
    }
    fmt.Println(v6)
    if (-1 > 42) {
        fmt.Println(v6)
        var v22 string = (v4 + ((v4 + v4) + v4))
        for fi23 := 0; fi23 < 17; fi23++ {
            if (255 < 1) {
                v4 = (v4 + ("hello" + (v22 + v4)))
                fmt.Println(v22)
            } else {
                if (fi23 < fi23) {
                    for fi24 := 0; fi24 < 15; fi24++ {
                        var v25 float32 = (1e10 * float32(3.14159))
                        fmt.Println(fi23, v25, fi24)
                        var v26 float32 = (0.1 + (3.14159 * (3.14159 * float32((0.1 / -2.25)))))
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                    v22 = v4
                    switch ("\n\t" + "hello") {
                    case "\n\t":
                        var v27 string = v4
                        _ = v27
                    case "":
                        if false {
                            v22 = ((v22 + (v22 + (v22 + v4))) + v22)
                        }
                    case "hello":
                        var v28 bool = false
                        _ = v28
                    }
                }
                v29 := GenS1{F0: 2, F1: 117}
                v4 = "foo bar"
                _ = v29
            }
            _ = fi23
        }
        _ = v22
    } else {
        v4 = ("hello" + ("a" + ""))
    }
    var v30 int32 = 0
    switch v4 {
    case "hello":
        var v31 uint = 975
        _ = v31
    case "\n\t":
        for fi32 := 0; fi32 < 6; fi32++ {
            fmt.Println(v4)
            for fi33 := 0; fi33 < 6; fi33++ {
                v30 += 168
                var v34 int64 = 9223372036854775807
                _ = fi33
                _ = v34
            }
            v30 -= v30
            _ = fi32
        }
    default:
        var v35 int64 = int64(v30)
        _ = v35
    }
    v36 := GenS2{F0: 988}
    v3 = append(v3, (v4 + v4))
    if (false || (47 < 222)) {
        for fi37 := 0; fi37 < 5; fi37++ {
            switch v4 {
            case "":
                _ = len(v1)
            default:
                if false {
                    fmt.Println(fi37, v4, v4)
                    v38 := GenS1{F0: 269, F1: 1}
                    for fi39 := 0; fi39 < 4; fi39++ {
                        var v40 bool = (42 >= 4294967295)
                        if ((!(v40 || v40)) && false) {
                            v41 := GenS3{F0: 255, F1: 1, F2: true}
                            switch ("a" + (("" + ("hello" + v4)) + v4)) {
                            case "hello":
                                fmt.Println(v6)
                            }
                            v42 := GenS4{F0: 2}
                            _ = v41
                            _ = v42
                        }
                        switch -716 {
                        case 127:
                            switch v4 {
                            case "\n\t":
                                v43 := GenS2{F0: 0}
                                _ = v43
                            case "a":
                                if (v4 == v4) {
                                    v44 := GenS5{F0: 2, F1: (v6 << ((983) & 63)), F2: 2147483648}
                                    if ((v40 && v40) && false) {
                                        delete(v1, v4)
                                    }
                                    switch (v4 + "\n\t") {
                                    case "\n\t":
                                        if v40 {
                                            _ = len(v3)
                                            delete(v2, ((("hello" + (v4 + v4)) + "foo bar") + v4))
                                        }
                                        delete(v1, (v4 + (v4 + v4)))
                                    case "foo bar":
                                        var v45 float32 = 1.5
                                        _ = v45
                                    case "a":
                                        if (65535 >= 1) {
                                            v46 := GenS6{F0: (true || false), F1: int8(v6)}
                                            _ = v46
                                        } else {
                                            v40 = (!false)
                                            var v47 int64 = -326
                                            _ = v47
                                        }
                                        switch fi37 {
                                        case -319:
                                            var v48 uint16 = 778
                                            _ = v48
                                        case 2147483647:
                                            var v49 int8 = 0
                                            _ = v49
                                        case -920:
                                            var v50 int8 = 127
                                            _ = v50
                                        }
                                        v40 = (739 < 763)
                                    default:
                                        if (v40 && (false || v40)) {
                                            if (v4 == v4) {
                                                if (!v40) {
                                                    var v51 int32 = (v30 + v30)
                                                    _ = v51
                                                }
                                                switch v4 {
                                                case "hello":
                                                    v52 := GenS7{F0: fi39}
                                                    _ = v52
                                                case "foo bar":
                                                    fmt.Println(fi37, v6, v30)
                                                default:
                                                    if v40 {
                                                        v40 = true
                                                        var v53 int16 = v6
                                                        v54 := GenS7{F0: -9223372036854775808}
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        v4 = (("foo bar" + (v4 + ("\n\t" + "foo bar"))) + (v4 + v4))
                                                        var v55 int64 = -1
                                                        _ = v55
                                                    }
                                                }
                                                if true {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                } else {
                                                    var v56 uint16 = 1
                                                    _ = v56
                                                }
                                            } else {
                                                var v57 float32 = (0.1 - (float32((1e10 / (3.14159 + 1e10))) / 1e10))
                                                v58 := GenS2{F0: 2}
                                                v59 := v2[v4]
                                                _ = v57
                                                _ = v58
                                                _ = v59
                                            }
                                            v30 = (v30 - v30)
                                            v4 = ((v4 + (v4 + ((((v4 + "\n\t") + v4) + "\n\t") + ("a" + (v4 + v4))))) + ("" + "hello"))
                                        } else {
                                            delete(v5, v4)
                                            if v40 {
                                                delete(v2, ((v4 + (v4 + "hello")) + (v4 + "hello")))
                                                switch "hello" {
                                                case "a":
                                                    var v60 float32 = 1e10
                                                    _ = v60
                                                case "\n\t":
                                                    v30 -= v30
                                                case "":
                                                    v4 = ("hello" + ((v4 + "a") + v4))
                                                default:
                                                    v4 = v4
                                                }
                                            } else {
                                                fmt.Println(v40, v30)
                                                v4 = "hello"
                                                v61 := GenS8{F0: -1, F1: 0}
                                                _ = v61
                                            }
                                            if (fi37 >= fi39) {
                                                v62 := GenS7{F0: (fi39 + (-(fi37)))}
                                                if v40 {
                                                    if (42 == 18446744073709551615) {
                                                        v30 += -2147483648
                                                        var v63 int16 = v6
                                                        _ = v63
                                                    }
                                                    if (((!v40) || ((807 != 42) && false)) && v40) {
                                                        v64 := GenS1{F0: 255, F1: int8(v30)}
                                                        _ = v64
                                                    }
                                                } else {
                                                    v1["foo bar"] = v4
                                                    switch (fi37 * fi37) {
                                                    case 42:
                                                        if ((0.1 == 0.0) && ((v6 % ((v6) | 1)) != v6)) {
                                                            switch v4 {
                                                            case "\n\t":
                                                                v4 = (v4 + ((v4 + (v4 + ("a" + "a"))) + ("foo bar" + v4)))
                                                            case "foo bar":
                                                                var v65 float32 = -2.25
                                                                _ = v65
                                                            case "hello":
                                                                var v66 string = v4
                                                                _ = v66
                                                            default:
                                                                var v67 uint = 1
                                                                v40 = (fi37 == (fi37 % ((fi39) | 1)))
                                                                v67 -= 1
                                                                if (v67 > 18446744073709551615) {
                                                                    switch v4 {
                                                                    case "hello":
                                                                        switch v4 {
                                                                        case "foo bar":
                                                                            v6 = (-(v6))
                                                                        case "hello":
                                                                            v68 := GenS9{F0: (fi39 ^ fi37)}
                                                                            _ = v68
                                                                        case "":
                                                                            if false {
                                                                                switch (^(fi39)) {
                                                                                case 127:
                                                                                    var v69 int32 = (v30 | (v30 >> (((v67 | v67)) & 63)))
                                                                                    _ = v69
                                                                                case -2147483648:
                                                                                    var v70 uint = (v67 | (v67 ^ (v67 >> ((v67) & 63))))
                                                                                    _ = v70
                                                                                case -595:
                                                                                    switch fi39 {
                                                                                    case 32767:
                                                                                        v71 := GenS1{F0: v67, F1: 100}
                                                                                        _ = v71
                                                                                    default:
                                                                                        if (!true) {
                                                                                            switch (v4 + v4) {
                                                                                            case "a":
                                                                                                v67 += v67
                                                                                            case "hello":
                                                                                                v30 += v30
                                                                                            case "foo bar":
                                                                                                v5[((v4 + v4) + "")] = fi37
                                                                                            }
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v72 := GenS8{F0: -15, F1: 2}
                                                                                    _ = v72
                                                                                }
                                                                                v67 += 1
                                                                                v73 := GenS8{F0: 127, F1: 42}
                                                                                v74 := GenS6{F0: true, F1: genFn1()}
                                                                                if false {
                                                                                    var v75 int64 = 1
                                                                                    var v76 float32 = -2.25
                                                                                    v3 = append(v3, (v4 + ((v4 + (v4 + v4)) + (v4 + "a"))))
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                }
                                                                                var v77 int8 = -8
                                                                                _ = v73
                                                                                _ = v74
                                                                                _ = v77
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    v6 = 4
                                                                }
                                                                fmt.Println(fi39)
                                                                _ = v67
                                                            }
                                                            fmt.Println(v6, fi39)
                                                            if v40 {
                                                                v40 = (((((0.1 / 1.0) / (0.1 * 1e10)) < ((((1.5 / 0.1) + -2.25) * 3.14159) - (1e10 + 0.0))) || (!v40)) && (649 < 18446744073709551615))
                                                                fmt.Println(v40, v4, v6)
                                                                switch v4 {
                                                                case "\n\t":
                                                                    fmt.Println(fi39, v30)
                                                                case "foo bar":
                                                                    switch -2147483648 {
                                                                    case 18:
                                                                        v30 = (v30 / ((v30) | 1))
                                                                    case -438:
                                                                        var v78 uint64 = 2
                                                                        _ = v78
                                                                    case 9223372036854775807:
                                                                        switch fi39 {
                                                                        case 0:
                                                                            v6 -= 1
                                                                        case -899:
                                                                            _ = len(v2)
                                                                        default:
                                                                            var v79 float32 = ((1.5 + (float32(0.1) / 3.14159)) - 3.14159)
                                                                            _ = v79
                                                                        }
                                                                        fmt.Println(v30, v40)
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case -847:
                                                        var v80 uint64 = 2
                                                        _ = v80
                                                    case 4294967295:
                                                        var v81 bool = (!(!((v40 || v40) && (!((255 == 1) || false)))))
                                                        var v82 int64 = 9223372036854775807
                                                        var v83 int16 = v6
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                }
                                                v6 = (v6 - v6)
                                                var v84 int32 = (v30 % ((2147483647) | 1))
                                                switch (("a" + (v4 + v4)) + v4) {
                                                case "a":
                                                    if ((!(104 < 2)) && (2147483647 == 2147483648)) {
                                                        v4 = (((v4 + v4) + v4) + v4)
                                                        v85 := v5[v4]
                                                        _ = v85
                                                    } else {
                                                        if (!false) {
                                                            if (2 > 255) {
                                                                var v86 int16 = v6
                                                                if (!(127 == 42)) {
                                                                    v30 = -2147483648
                                                                } else {
                                                                    if true {
                                                                        switch v4 {
                                                                        case "\n\t":
                                                                            v86 = (v6 << ((332) & 63))
                                                                        case "":
                                                                            fmt.Println(v4, fi37, v6)
                                                                        case "hello":
                                                                            v30 += 1
                                                                        default:
                                                                            v6 -= v6
                                                                        }
                                                                        v40 = (!false)
                                                                        if v40 {
                                                                            if true {
                                                                                var v87 int32 = v30
                                                                                _ = v87
                                                                            }
                                                                            if ((("foo bar" + v4) + v4) != "a") {
                                                                                fmt.Println(v40, v4)
                                                                                var v88 int64 = 42
                                                                                v88 += v88
                                                                                _ = v88
                                                                            }
                                                                        }
                                                                    }
                                                                    v89 := GenS1{F0: 64, F1: -78}
                                                                    _ = v89
                                                                }
                                                                switch fi39 {
                                                                case -9223372036854775808:
                                                                    if v40 {
                                                                        if true {
                                                                            delete(v5, (v4 + ("foo bar" + v4)))
                                                                            v90 := v2[v4]
                                                                            v30 = v30
                                                                            _ = v90
                                                                        }
                                                                    } else {
                                                                        v6 -= (v6 & 13)
                                                                    }
                                                                case 825:
                                                                    _ = len(v5)
                                                                default:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                                _ = v86
                                                            }
                                                            _ = len(v3)
                                                        } else {
                                                            switch (v4 + ("" + "hello")) {
                                                            case "foo bar":
                                                                v1["a"] = ("hello" + v4)
                                                            case "\n\t":
                                                                v40 = (true && v40)
                                                            case "":
                                                                fmt.Println(v40, v4, v30)
                                                            default:
                                                                switch (((v4 + v4) + ("hello" + ((v4 + "") + v4))) + "\n\t") {
                                                                case "\n\t":
                                                                    var v91 int8 = -1
                                                                    _ = v91
                                                                case "a":
                                                                    if (2147483647 > 456) {
                                                                        switch 32767 {
                                                                        case 4294967295:
                                                                            fmt.Println(v40)
                                                                            if v40 {
                                                                                delete(v2, v4)
                                                                            } else {
                                                                                var v92 bool = (v4 < "hello")
                                                                                switch "\n\t" {
                                                                                case "hello":
                                                                                    var v93 string = v4
                                                                                    _ = v93
                                                                                case "":
                                                                                    v94 := GenS2{F0: 18446744073709551615}
                                                                                    _ = v94
                                                                                default:
                                                                                    v95 := v1[v4]
                                                                                    _ = v95
                                                                                }
                                                                                _ = v92
                                                                            }
                                                                            v84 = (v30 & 1)
                                                                        case 42:
                                                                            fmt.Println(v30, v30)
                                                                        case 127:
                                                                            var v96 uint16 = 948
                                                                            _ = v96
                                                                        }
                                                                    } else {
                                                                        switch 387 {
                                                                        case -1:
                                                                            v4 = v4
                                                                        case 9223372036854775807:
                                                                            v3 = append(v3, v4)
                                                                        case 32767:
                                                                            var v97 uint = 706
                                                                            _ = v97
                                                                        }
                                                                        if len(v3) > 0 {
                                                                            v3[0] = v4
                                                                        }
                                                                    }
                                                                case "":
                                                                    v30 = genFn0()
                                                                default:
                                                                    _ = len(v2)
                                                                }
                                                            }
                                                        }
                                                        v98 := GenS10{F0: v6}
                                                        switch (v4 + v4) {
                                                        case "":
                                                            if v40 {
                                                                v3 = append(v3, "a")
                                                                var v99 uint32 = 63
                                                                v6 = (-(v6))
                                                                _ = v99
                                                            }
                                                        default:
                                                            switch v4 {
                                                            case "\n\t":
                                                                var v100 uint32 = 42
                                                                _ = v100
                                                            default:
                                                                if v40 {
                                                                    v40 = v40
                                                                } else {
                                                                    v101 := GenS6{F0: v40, F1: -128}
                                                                    _ = len(v5)
                                                                    _ = v101
                                                                }
                                                            }
                                                        }
                                                        _ = v98
                                                    }
                                                }
                                                fmt.Println(v84)
                                                switch (fi39 | fi37) {
                                                case 2147483647:
                                                    var v102 bool = false
                                                    _ = v102
                                                case 1:
                                                    switch v4 {
                                                    case "foo bar":
                                                        v4 = v4
                                                    case "a":
                                                        v6 -= v6
                                                    }
                                                }
                                                _ = v62
                                                _ = v84
                                            }
                                        }
                                    }
                                    _ = v44
                                }
                            case "hello":
                                v103 := GenS6{F0: (float64(1.5) == 1e10), F1: -128}
                                _ = v103
                            }
                        }
                        v40 = (false && (1e10 < 0.0))
                        var v104 float64 = (float64((0.0 + 0.0)) + (1.5 - 1.5))
                        var v105 float32 = 0.0
                        _ = fi39
                        _ = v104
                        _ = v105
                        _ = v40
                    }
                    _ = v38
                } else {
                    if true {
                        v106 := GenS11{F0: 127, F1: v30, F2: fi37}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v106
                    } else {
                        _ = len(v5)
                        var v107 uint32 = 4294967295
                        _ = v107
                    }
                    for fi108 := 0; fi108 < 9; fi108++ {
                        var v109 uint = 0
                        var v110 float32 = 0.0
                        if true {
                            var v111 uint16 = 255
                            var v112 string = v4
                            _ = v111
                            _ = v112
                        } else {
                            var v113 int8 = -1
                            _ = v113
                        }
                        _ = fi108
                        _ = v109
                        _ = v110
                    }
                    var v114 int16 = (v6 << ((65535) & 63))
                    _ = v114
                }
            }
            fmt.Println(v4, v6)
            for fi115 := 0; fi115 < 11; fi115++ {
                v4 = v4
                _ = fi115
            }
            _ = fi37
        }
    } else {
        v116 := GenS4{F0: 857}
        v4 = v4
        _ = v116
    }
    if false {
        switch (v4 + "a") {
        case "a":
            fmt.Println(v6)
        case "foo bar":
            for fi117 := 0; fi117 < 8; fi117++ {
                if len(v3) > 0 {
                    v3[0] = v4
                }
                var v118 uint16 = 42
                var v119 float32 = -2.25
                _ = fi117
                _ = v118
                _ = v119
            }
        case "":
            v6 = v6
        }
        v6 -= -32768
        var v120 float32 = 0.1
        _ = v120
    }
    var v121 uint64 = 42
    _ = len(v1)
    v121 = v121
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v121:", v121)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v121
    _ = v2
    _ = v3
    _ = v30
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
