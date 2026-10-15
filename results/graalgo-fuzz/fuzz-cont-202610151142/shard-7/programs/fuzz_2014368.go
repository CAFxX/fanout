package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64; F1 int16; F2 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 string; F2 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 uint8 }
type GenS8 struct { F0 bool; F1 bool }
type GenS9 struct { F0 int32; F1 int16; F2 string }
type GenS10 struct { F0 uint16; F1 uint16; F2 int }
type GenS11 struct { F0 string; F1 int; F2 int32 }
type GenS12 struct { F0 uint16; F1 uint64; F2 uint32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int16; F1 uint32 }
type GenS17 struct { F0 int; F1 uint64; F2 uint }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 int64; F1 int64; F2 uint64 }
type GenS22 struct { F0 int; F1 string }
type GenS23 struct { F0 int8; F1 int64; F2 string }
type GenS24 struct { F0 uint8; F1 int8 }
type GenS25 struct { F0 uint16; F1 int32; F2 string }
type GenS26 struct { F0 int32; F1 uint }
type GenS27 struct { F0 bool; F1 int32 }
type GenS28 struct { F0 uint32; F1 string; F2 int16 }

func main() {
    v1 := []int{}
    var v2 int8 = 52
    var v3 int = -9223372036854775808
    v4 := make(map[string]int)
    v5 := GenS1{F0: 1}
    v3 -= (v3 << ((430) & 63))
    var v6 uint32 = 65535
    _ = len(v1)
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            if false {
                var v9 float64 = (1.5 - (0.1 + 0.0))
                _ = v9
            } else {
                v4["hello"] = (-(fi7))
                var v10 bool = (v6 != 65535)
                _ = v10
            }
            _ = fi8
        }
        var v11 int8 = 127
        v12 := GenS2{F0: "hello"}
        _ = fi7
        _ = v11
        _ = v12
    }
    v13 := GenS3{F0: 491, F1: 32767, F2: 42}
    if true {
        for fi14 := 0; fi14 < 7; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                v6 -= v6
                if (314 != 65535) {
                    _ = len(v1)
                    switch v3 {
                    case -446:
                        switch ("hello" + "") {
                        case "a":
                            switch -1 {
                            case -9223372036854775808:
                                switch "hello" {
                                case "":
                                    switch 1 {
                                    case 4294967295:
                                        fmt.Println(v3, fi14, v3)
                                    }
                                case "a":
                                    v3 = fi14
                                case "foo bar":
                                    switch (-(fi14)) {
                                    case -9223372036854775808:
                                        v1 = append(v1, fi15)
                                    case 42:
                                        fmt.Println(v6, v3)
                                    case -1:
                                        switch "hello" {
                                        case "\n\t":
                                            if (false || true) {
                                                _ = len(v1)
                                                var v16 int8 = 127
                                                var v17 uint16 = 911
                                                switch "foo bar" {
                                                case "hello":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                case "\n\t":
                                                    v18 := GenS2{F0: ("hello" + ("a" + "hello"))}
                                                    _ = v18
                                                default:
                                                    v19 := GenS3{F0: -1, F1: -1, F2: 0}
                                                    _ = v19
                                                }
                                                var v20 int = (v3 ^ v3)
                                                _ = v16
                                                _ = v17
                                                _ = v20
                                            } else {
                                                switch (("" + "\n\t") + (("hello" + "\n\t") + ("foo bar" + "hello"))) {
                                                case "\n\t":
                                                    v3 = v3
                                                case "":
                                                    var v21 int = (fi15 | (v3 | (-(fi15))))
                                                    if (v3 <= (v21 >> ((797) & 63))) {
                                                        var v22 int8 = (v2 % ((int8(v21)) | 1))
                                                        _ = v22
                                                    }
                                                    var v23 int16 = -32768
                                                    _ = v21
                                                    _ = v23
                                                }
                                                fmt.Println(fi15, v6, fi15)
                                                if false {
                                                    v6 = (^(v6))
                                                    if (0.0 <= 1e10) {
                                                        fmt.Println(v2, fi14)
                                                    } else {
                                                        v24 := GenS4{F0: v3}
                                                        v2 -= v2
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v24
                                                    }
                                                    switch fi14 {
                                                    case 128:
                                                        fmt.Println(fi14)
                                                    case -650:
                                                        v25 := GenS4{F0: (v3 % (((fi15 + v3)) | 1))}
                                                        _ = v25
                                                    }
                                                } else {
                                                    v26 := GenS1{F0: -1}
                                                    switch "hello" {
                                                    case "foo bar":
                                                        v2 = v2
                                                    case "":
                                                        switch (fi14 - 42) {
                                                        case 127:
                                                            v2 = 42
                                                        case -984:
                                                            if ((false || true) || (468 > -1)) {
                                                                v27 := GenS5{F0: 266, F1: "\n\t", F2: (^(v6))}
                                                                v2 -= v2
                                                                if false {
                                                                    if (!(false && false)) {
                                                                        switch ("a" + "") {
                                                                        case "a":
                                                                            v28 := GenS5{F0: -965, F1: "", F2: 42}
                                                                            _ = v28
                                                                        case "\n\t":
                                                                            v2 += v2
                                                                        }
                                                                    } else {
                                                                        _ = len(v1)
                                                                    }
                                                                    fmt.Println(v3, fi14)
                                                                    fmt.Println(fi15, v6)
                                                                } else {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                                switch 881 {
                                                                case 1:
                                                                    v29 := GenS1{F0: 736}
                                                                    _ = v29
                                                                default:
                                                                    if true {
                                                                        v30 := GenS6{F0: (fi14 | v3)}
                                                                        _ = v30
                                                                    }
                                                                }
                                                                _ = v27
                                                            } else {
                                                                v31 := GenS7{F0: 42, F1: 1}
                                                                delete(v4, "a")
                                                                _ = v31
                                                            }
                                                        case 32767:
                                                            if (false && false) {
                                                                v32 := GenS8{F0: (!((1 != 2) || ((("a" + (("hello" + (("" + (("\n\t" + "hello") + ("hello" + ""))) + ("\n\t" + "foo bar"))) + (((("hello" + (((((("a" + (("" + "hello") + ("\n\t" + "foo bar"))) + "\n\t") + "") + "hello") + "a") + "")) + "foo bar") + "") + "hello"))) + "foo bar") >= ("" + "")))), F1: false}
                                                                _ = v32
                                                            }
                                                        }
                                                    }
                                                    _ = v26
                                                }
                                            }
                                        case "foo bar":
                                            v33 := GenS9{F0: 42, F1: -32768, F2: "hello"}
                                            _ = v33
                                        case "hello":
                                            if ((v6 / (((^(v6))) | 1)) <= v6) {
                                                switch "" {
                                                case "hello":
                                                    fmt.Println(v2)
                                                case "foo bar":
                                                    if (!true) {
                                                        if false {
                                                            v34 := GenS1{F0: 42}
                                                            var v35 uint = 0
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        v36 := GenS7{F0: 42, F1: 42}
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        _ = v36
                                                    } else {
                                                        var v37 string = "foo bar"
                                                        _ = v37
                                                    }
                                                default:
                                                    var v38 int16 = 690
                                                    v39 := GenS9{F0: 0, F1: v38, F2: "hello"}
                                                    _ = v38
                                                    _ = v39
                                                }
                                            } else {
                                                var v40 string = "foo bar"
                                                _ = v40
                                            }
                                        default:
                                            var v41 int = fi14
                                            _ = v41
                                        }
                                    }
                                default:
                                    var v42 int16 = 32767
                                    _ = v42
                                }
                                switch ("foo bar" + "foo bar") {
                                case "":
                                    v43 := GenS10{F0: 0, F1: 255, F2: (v3 | fi14)}
                                    _ = v43
                                case "a":
                                    v44 := GenS11{F0: "", F1: (fi15 ^ 9223372036854775807), F2: -2147483648}
                                    _ = v44
                                }
                                v3 = -550
                                v45 := GenS12{F0: 2, F1: 648, F2: v6}
                                fmt.Println(v6, v3)
                                v6 = (^(v6))
                                _ = v45
                            case 1:
                                var v46 string = "foo bar"
                                _ = v46
                            default:
                                var v47 bool = true
                                _ = v47
                            }
                        case "foo bar":
                            v48 := GenS13{F0: 0}
                            _ = v48
                        case "hello":
                            if false {
                                v49 := GenS9{F0: 0, F1: -1, F2: ("foo bar" + "a")}
                                if false {
                                    _ = len(v1)
                                    if len(v1) > 0 {
                                        v1[0] = (fi15 | int(v2))
                                    }
                                    v2 = -28
                                    if ((171 != 4294967295) || (!(((-2.25 * -2.25) + 3.14159) <= float64(0.0)))) {
                                        if (!(false && (("a" >= (("hello" + "") + "foo bar")) || (14 >= 42)))) {
                                            if (0.1 < 3.14159) {
                                                var v50 int32 = 42
                                                fmt.Println(v2, fi14, v2)
                                                switch "hello" {
                                                case "foo bar":
                                                    fmt.Println(v6, v50, v50)
                                                case "a":
                                                    v3 -= fi14
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi14 + fi14)
                                                    }
                                                }
                                                _ = v50
                                            } else {
                                                v2 = (-(v2))
                                            }
                                            delete(v4, "")
                                        } else {
                                            v6 = v6
                                            fmt.Println(v2, v6)
                                        }
                                        var v51 uint64 = 596
                                        v52 := GenS4{F0: (fi14 + v3)}
                                        var v53 uint = 689
                                        v51 -= 4294967295
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    } else {
                                        var v54 uint64 = 4294967295
                                        v55 := GenS13{F0: 0}
                                        v54 += v54
                                        _ = v54
                                        _ = v55
                                    }
                                }
                                switch fi14 {
                                case 42:
                                    fmt.Println(fi15, v3)
                                default:
                                    v4[("a" + "a")] = 144
                                }
                                _ = v49
                            }
                        default:
                            var v56 bool = (true && ((v6 * 842) < (v6 << ((18446744073709551615) & 63))))
                            v57 := v4[("" + "")]
                            _ = v56
                            _ = v57
                        }
                    case -452:
                        switch ("\n\t" + "") {
                        case "a":
                            if (!(true && (true || (true && false)))) {
                                v58 := GenS13{F0: -1}
                                var v59 bool = (((1 != 42) || false) || ((0.1 * -2.25) != -2.25))
                                _ = v58
                                _ = v59
                            }
                        case "":
                            v60 := GenS13{F0: -116}
                            _ = v60
                        case "\n\t":
                            fmt.Println(v6, v2, fi14)
                            switch (v3 * (v3 - fi15)) {
                            case -2147483648:
                                v3 += 32767
                            case 4294967295:
                                var v61 float64 = 1e10
                                _ = v61
                            case 127:
                                switch (-(fi15)) {
                                case 1:
                                    v3 += v3
                                case -652:
                                    v6 = v6
                                }
                            default:
                                switch "\n\t" {
                                case "foo bar":
                                    v62 := GenS14{F0: 1}
                                    _ = v62
                                case "":
                                    v63 := GenS11{F0: "\n\t", F1: fi15, F2: 2147483647}
                                    _ = v63
                                default:
                                    var v64 float32 = -2.25
                                    _ = v64
                                }
                            }
                        }
                    case 128:
                        if false {
                            fmt.Println(v6)
                        }
                    }
                }
                var v65 int32 = 964
                _ = fi15
                _ = v65
            }
            _ = fi14
        }
    } else {
        if (!true) {
            v3 += v3
        } else {
            if (42 >= 0) {
                v3 = 32767
                v66 := v4[""]
                _ = v66
            } else {
                v67 := GenS15{F0: 18446744073709551615}
                var v68 bool = (3.14159 < 1e10)
                _ = v67
                _ = v68
            }
        }
        var v69 uint64 = 874
        _ = v69
    }
    if false {
        if (false && true) {
            v70 := v4[(("" + "hello") + "a")]
            var v71 uint32 = (v6 % (((v6 - v6)) | 1))
            for fi72 := 0; fi72 < 17; fi72++ {
                var v73 uint32 = v6
                fmt.Println(fi72, fi72, v6)
                _ = fi72
                _ = v73
            }
            _ = v70
            _ = v71
        } else {
            for fi74 := 0; fi74 < 4; fi74++ {
                v2 = (^(v2))
                v75 := GenS16{F0: 0, F1: (v6 >> ((2) & 63))}
                if ((-783 <= -1) || true) {
                    v76 := GenS17{F0: fi74, F1: 18446744073709551615, F2: 60}
                    var v77 uint64 = 2
                    v77 = 4294967295
                    _ = v76
                    _ = v77
                } else {
                    switch fi74 {
                    case -401:
                        if true {
                            var v78 string = "hello"
                            _ = v78
                        } else {
                            if true {
                                v3 += 128
                            }
                            v3 += (v3 << ((4294967295) & 63))
                        }
                    default:
                        _ = len(v1)
                    }
                    if (!false) {
                        fmt.Println(v3, v2, v3)
                        v6 -= (v6 + v6)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if (42 >= 42) {
                        v2 = v2
                        v6 = v6
                    } else {
                        fmt.Println(v6)
                        v6 = v6
                        for fi79 := 0; fi79 < 8; fi79++ {
                            v80 := v4["a"]
                            if true {
                                fmt.Println(v80, fi74, fi74)
                                switch (-(v80)) {
                                case 42:
                                    v80 += (-(fi79))
                                case -1:
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        var v81 float64 = (float64((0.0 * (1e10 / (0.1 + (1.5 / 0.1))))) / (1.5 / 1.5))
                                        _ = v81
                                    }
                                case 614:
                                    if true {
                                        var v82 uint16 = 32767
                                        _ = v82
                                    }
                                }
                                v3 = v3
                                if (!(v2 >= v2)) {
                                    fmt.Println(v2, fi74, v2)
                                    if (!(float64(((0.0 * 3.14159) + -2.25)) > float64((-2.25 / (1.5 * -2.25))))) {
                                        switch "\n\t" {
                                        case "":
                                            var v83 uint = 18446744073709551615
                                            _ = v83
                                        case "a":
                                            v84 := GenS18{F0: false}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            fmt.Println(v3)
                                            _ = v84
                                        case "foo bar":
                                            v3 = v3
                                        }
                                        var v85 uint32 = v6
                                        _ = v85
                                    }
                                    if false {
                                        fmt.Println(v2, v6)
                                    } else {
                                        switch (^(v3)) {
                                        case -377:
                                            v80 = int(v2)
                                        case 32768:
                                            var v86 uint8 = 2
                                            _ = v86
                                        case -501:
                                            v2 -= (v2 % ((v2) | 1))
                                            delete(v4, ((("hello" + "foo bar") + ("hello" + ("" + ("" + "foo bar")))) + "hello"))
                                            if (-154 >= -235) {
                                                var v87 int = -940
                                                var v88 float32 = 3.14159
                                                _ = v87
                                                _ = v88
                                            }
                                            var v89 int64 = 42
                                            var v90 string = ("hello" + "hello")
                                            if ((1 >= -746) && false) {
                                                var v91 string = "foo bar"
                                                if (4294967295 >= 1) {
                                                    var v92 string = (((v90 + "hello") + v91) + v90)
                                                    switch ("\n\t" + v90) {
                                                    case "":
                                                        var v93 int64 = (v89 - v89)
                                                        _ = v93
                                                    case "hello":
                                                        switch int(v2) {
                                                        case 107:
                                                            v94 := GenS5{F0: -1, F1: (v91 + v92), F2: v6}
                                                            _ = v94
                                                        case 42:
                                                            if (((((v90 + "") + (v90 + v90)) + "hello") + ((v92 + "hello") + "foo bar")) > v91) {
                                                                var v95 int8 = v2
                                                                fmt.Println(fi74, v3, v95)
                                                                switch fi79 {
                                                                case 4294967295:
                                                                    fmt.Println(v91)
                                                                default:
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                                _ = v95
                                                            } else {
                                                                if true {
                                                                    fmt.Println(v3, v6)
                                                                    v80 = (v80 / ((-151) | 1))
                                                                    fmt.Println(fi79)
                                                                } else {
                                                                    v4[(v92 + "")] = v3
                                                                    v6 = 2
                                                                    if len(v1) > 0 {
                                                                        v1[0] = fi79
                                                                    }
                                                                    var v96 int = v3
                                                                    var v97 int = (fi79 - fi74)
                                                                    v1 = append(v1, 42)
                                                                    _ = v96
                                                                    _ = v97
                                                                }
                                                                var v98 float32 = 0.0
                                                                var v99 uint16 = 120
                                                                var v100 uint = 1
                                                                fmt.Println(v3, fi74)
                                                                _ = v100
                                                                _ = v98
                                                                _ = v99
                                                            }
                                                        case -98:
                                                            if true {
                                                                fmt.Println(v89, v91)
                                                                v101 := GenS19{F0: true}
                                                                _ = v101
                                                            }
                                                        default:
                                                            v102 := GenS15{F0: 52}
                                                            _ = v102
                                                        }
                                                    }
                                                    _ = len(v4)
                                                    _ = v92
                                                } else {
                                                    switch fi74 {
                                                    case 985:
                                                        v103 := GenS20{F0: (^(v6))}
                                                        _ = v103
                                                    default:
                                                        if (v6 < v6) {
                                                            if (!false) {
                                                                var v104 int64 = -9223372036854775808
                                                                switch (v91 + ((v91 + (v91 + (v91 + ""))) + ("a" + ("a" + (v91 + (v91 + "a")))))) {
                                                                case "a":
                                                                    fmt.Println(fi74)
                                                                case "\n\t":
                                                                    v90 = v90
                                                                }
                                                                v3 = (fi79 % ((fi79) | 1))
                                                                _ = v104
                                                            }
                                                            if (!false) {
                                                                if false {
                                                                    v105 := v4[(v90 + "")]
                                                                    _ = v105
                                                                }
                                                                var v106 float32 = 3.14159
                                                                _ = v106
                                                            } else {
                                                                v107 := GenS20{F0: (v6 / ((523) | 1))}
                                                                v108 := GenS8{F0: (!(!(0 == 32767))), F1: false}
                                                                switch v90 {
                                                                case "a":
                                                                    var v109 uint8 = 167
                                                                    _ = v109
                                                                case "":
                                                                    var v110 uint8 = 255
                                                                    _ = v110
                                                                default:
                                                                    var v111 int8 = v2
                                                                    _ = v111
                                                                }
                                                                _ = v107
                                                                _ = v108
                                                            }
                                                            if (!false) {
                                                                fmt.Println(v90)
                                                                if ("foo bar" <= "") {
                                                                    var v112 int = fi79
                                                                    switch ((v90 + "hello") + v91) {
                                                                    case "":
                                                                        if (!true) {
                                                                            if (3.14159 == float32(3.14159)) {
                                                                                var v113 string = v90
                                                                                var v114 uint64 = 469
                                                                                fmt.Println(fi74, v80, v90)
                                                                                if (v89 < v89) {
                                                                                    var v115 uint = 0
                                                                                    var v116 int8 = (v2 << ((42) & 63))
                                                                                    var v117 int = -2147483648
                                                                                    _ = v115
                                                                                    _ = v116
                                                                                    _ = v117
                                                                                } else {
                                                                                    v118 := GenS21{F0: int64(v114), F1: 0, F2: 4294967295}
                                                                                    _ = v118
                                                                                }
                                                                                _ = v113
                                                                                _ = v114
                                                                            } else {
                                                                                v112 -= v112
                                                                            }
                                                                        } else {
                                                                            switch v91 {
                                                                            case "foo bar":
                                                                                v3 = -477
                                                                            case "hello":
                                                                                var v119 int64 = -9223372036854775808
                                                                                _ = v119
                                                                            case "":
                                                                                v4[v91] = -790
                                                                            }
                                                                            _ = len(v4)
                                                                        }
                                                                    case "foo bar":
                                                                        if (true && false) {
                                                                            if (((1.5 * 1e10) < 1e10) || false) {
                                                                                v1 = append(v1, (v80 / ((-2147483648) | 1)))
                                                                                v120 := GenS20{F0: 2147483647}
                                                                                var v121 string = v90
                                                                                _ = v120
                                                                                _ = v121
                                                                            } else {
                                                                                v122 := v4["\n\t"]
                                                                                v123 := v4[(v91 + "hello")]
                                                                                _ = v122
                                                                                _ = v123
                                                                            }
                                                                            v80 = v3
                                                                        }
                                                                    case "hello":
                                                                        var v124 int8 = v2
                                                                        _ = v124
                                                                    }
                                                                    v2 = (v2 >> ((343) & 63))
                                                                    var v125 uint = 255
                                                                    _ = v112
                                                                    _ = v125
                                                                }
                                                                v6 -= (^(v6))
                                                            } else {
                                                                v80 += v80
                                                            }
                                                        }
                                                    }
                                                    v91 = ("\n\t" + "foo bar")
                                                    v3 += 9223372036854775807
                                                }
                                                _ = v91
                                            } else {
                                                v80 = (-(fi79))
                                                if true {
                                                    if (1 > v6) {
                                                        v126 := GenS22{F0: 128, F1: v90}
                                                        _ = len(v1)
                                                        _ = v126
                                                    } else {
                                                        var v127 uint64 = 2
                                                        if false {
                                                            v128 := GenS23{F0: (v2 / (((v2 - (^(v2)))) | 1)), F1: (v89 % (((v89 - 9223372036854775807)) | 1)), F2: v90}
                                                            var v129 string = v90
                                                            _ = v128
                                                            _ = v129
                                                        } else {
                                                            if true {
                                                                v80 = v80
                                                                v2 -= 127
                                                            }
                                                        }
                                                        if false {
                                                            if true {
                                                                switch v3 {
                                                                case 2147483647:
                                                                    fmt.Println(v2, v127)
                                                                default:
                                                                    var v130 int16 = 32767
                                                                    _ = v130
                                                                }
                                                                if false {
                                                                    var v131 int16 = 42
                                                                    v132 := GenS24{F0: 202, F1: -104}
                                                                    var v133 int64 = -887
                                                                    _ = v131
                                                                    _ = v132
                                                                    _ = v133
                                                                } else {
                                                                    var v134 bool = false
                                                                    var v135 int8 = v2
                                                                    _ = v134
                                                                    _ = v135
                                                                }
                                                            } else {
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                v136 := GenS25{F0: 964, F1: -1, F2: (v90 + v90)}
                                                                switch "a" {
                                                                case "foo bar":
                                                                    var v137 uint64 = (v127 - (v127 - 42))
                                                                    _ = v137
                                                                case "":
                                                                    switch fi79 {
                                                                    case 1:
                                                                        if (42 >= 1) {
                                                                            fmt.Println(v3, v3)
                                                                            var v138 float64 = 1.5
                                                                            switch v90 {
                                                                            case "hello":
                                                                                var v139 bool = false
                                                                                _ = v139
                                                                            case "foo bar":
                                                                                v140 := GenS26{F0: -1, F1: uint(v127)}
                                                                                _ = v140
                                                                            case "a":
                                                                                v1 = append(v1, (v3 ^ v3))
                                                                            default:
                                                                                var v141 uint32 = 707
                                                                                _ = v141
                                                                            }
                                                                            _ = v138
                                                                        } else {
                                                                            v89 -= v89
                                                                            var v142 uint64 = v127
                                                                            switch int(v89) {
                                                                            case 127:
                                                                                v143 := v4["hello"]
                                                                                _ = v143
                                                                            case -84:
                                                                                var v144 uint = 212
                                                                                _ = v144
                                                                            default:
                                                                                var v145 int8 = 42
                                                                                _ = v145
                                                                            }
                                                                            _ = v142
                                                                        }
                                                                    default:
                                                                        switch v90 {
                                                                        case "hello":
                                                                            v80 -= (fi79 * -2147483648)
                                                                        case "a":
                                                                            if true {
                                                                                v146 := GenS17{F0: 127, F1: v127, F2: 1}
                                                                                var v147 int64 = (v89 >> ((1) & 63))
                                                                                var v148 float32 = -2.25
                                                                                _ = v146
                                                                                _ = v147
                                                                                _ = v148
                                                                            }
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    if true {
                                                                        var v149 int32 = 719
                                                                        var v150 bool = false
                                                                        v1 = append(v1, v80)
                                                                        var v151 int = (-(v80))
                                                                        _ = v149
                                                                        _ = v150
                                                                        _ = v151
                                                                    } else {
                                                                        v152 := GenS27{F0: true, F1: 0}
                                                                        switch 32768 {
                                                                        case 32767:
                                                                            fmt.Println(v3, v80)
                                                                        }
                                                                        _ = v152
                                                                    }
                                                                default:
                                                                    if (0.0 <= 3.14159) {
                                                                        var v153 uint32 = (v6 * v6)
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 4294967295
                                                                        }
                                                                        v90 = "\n\t"
                                                                        v3 += 0
                                                                        var v154 float32 = ((0.1 / 3.14159) / 0.1)
                                                                        _ = v153
                                                                        _ = v154
                                                                    } else {
                                                                        var v155 int16 = 465
                                                                        v127 += uint64(v6)
                                                                        switch (v3 / ((0) | 1)) {
                                                                        case 145:
                                                                            switch v90 {
                                                                            case "a":
                                                                                v155 -= v155
                                                                            case "foo bar":
                                                                                fmt.Println(v155, fi79)
                                                                            default:
                                                                                v156 := GenS28{F0: v6, F1: v90, F2: v155}
                                                                                _ = v156
                                                                            }
                                                                        case 42:
                                                                            v1 = append(v1, (^(v80)))
                                                                        case 2147483647:
                                                                            if ("hello" == "a") {
                                                                                var v157 float64 = (0.0 - 3.14159)
                                                                                _ = v157
                                                                            }
                                                                        default:
                                                                            v90 = v90
                                                                        }
                                                                        _ = v155
                                                                    }
                                                                }
                                                                _ = v136
                                                            }
                                                            v1 = append(v1, fi74)
                                                            var v158 uint64 = v127
                                                            _ = v158
                                                        }
                                                        _ = v127
                                                    }
                                                    v2 = (v2 << ((255) & 63))
                                                    v159 := GenS14{F0: 1}
                                                    _ = v159
                                                }
                                            }
                                            delete(v4, (v90 + v90))
                                            _ = v89
                                            _ = v90
                                        }
                                    }
                                }
                            } else {
                                switch "a" {
                                case "a":
                                    v2 -= 100
                                case "foo bar":
                                    v3 += v3
                                }
                                v160 := GenS2{F0: ("a" + "")}
                                v3 += v3
                                _ = v160
                            }
                            var v161 string = "foo bar"
                            _ = fi79
                            _ = v161
                            _ = v80
                        }
                    }
                }
                _ = fi74
                _ = v75
            }
            v162 := v4["a"]
            _ = v162
        }
        v163 := v4[(("" + "\n\t") + "foo bar")]
        _ = v163
    }
    if false {
        if true {
            fmt.Println(v3, v6, v6)
            var v164 uint8 = 255
            var v165 int16 = 1000
            _ = v164
            _ = v165
        }
    }
    v3 += (v3 ^ 42)
    for fi166 := 0; fi166 < 15; fi166++ {
        fmt.Println(v6)
        v4["hello"] = (v3 & v3)
        fmt.Println(v6, v3, v3)
        _ = fi166
    }
    var v167 int64 = -556
    v168 := GenS21{F0: (v167 - (-(v167))), F1: (v167 - v167), F2: 349}
    if false {
        for fi169 := 0; fi169 < 9; fi169++ {
            switch "foo bar" {
            case "foo bar":
                var v170 uint32 = 4294967295
                _ = v170
            case "\n\t":
                var v171 string = "foo bar"
                _ = v171
            default:
                for fi172 := 0; fi172 < 18; fi172++ {
                    var v173 float64 = 0.1
                    var v174 uint8 = 42
                    v1 = append(v1, fi169)
                    _ = fi172
                    _ = v173
                    _ = v174
                }
            }
            v3 -= (v3 - (v3 | (fi169 * fi169)))
            _ = fi169
        }
        v1 = append(v1, 42)
        if (18446744073709551615 != 1) {
            for fi175 := 0; fi175 < 18; fi175++ {
                var v176 int64 = (v167 - (-(v167)))
                var v177 bool = (0.1 < 1e10)
                v6 = v6
                _ = fi175
                _ = v176
                _ = v177
            }
            if false {
                v1 = append(v1, v3)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v167:", v167)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v167
    _ = v168
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
