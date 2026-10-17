package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint32; F2 int64 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int32; F1 int64; F2 int64 }
type GenS7 struct { F0 uint8; F1 int8; F2 string }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 bool; F1 int; F2 uint32 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint64; F1 int32; F2 string }
type GenS13 struct { F0 int32; F1 string }
type GenS14 struct { F0 int64; F1 int16; F2 int8 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint32) uint64 {
    v6 := GenS1{F0: (1e10 < float64(0.0)), F1: 115, F2: -1}
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 += p0
        if false {
            if true {
                v8 := GenS2{F0: 248}
                p0 += p0
                v9 := GenS2{F0: 4294967295}
                _ = v8
                _ = v9
            }
            var v10 float32 = 3.14159
            _ = v10
        } else {
            v11 := GenS3{F0: 18446744073709551615, F1: (^(p0))}
            var v12 uint = 249
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    v13 := []int{1, -9223372036854775808}
    fmt.Println(p0)
    _ = p0
    _ = v13
    _ = v6
    return 2
}

func main() {
    var v1 int16 = -697
    v2 := []string{"hello", "\n\t"}
    v3 := make(map[uint32]string)
    var v4 int8 = 1
    var v5 uint16 = 32767
    for fi14 := 0; fi14 < 17; fi14++ {
        switch ("\n\t" + ("a" + ("hello" + "a"))) {
        case "foo bar":
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 uint = 18446744073709551615
                _ = fi15
                _ = v16
            }
        default:
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 bool = true
                if true {
                    v1 = (-(v1))
                } else {
                    v5 = (v5 % ((65535) | 1))
                    switch ("\n\t" + ("foo bar" + ("foo bar" + (("a" + (("" + "a") + "a")) + "")))) {
                    case "":
                        if true {
                            v5 += (^(v5))
                            var v19 uint = 36
                            var v20 uint16 = (v5 >> (((^(v19))) & 63))
                            _ = v19
                            _ = v20
                        } else {
                            switch "a" {
                            case "":
                                v1 -= (v1 >> ((200) & 63))
                            case "\n\t":
                                var v21 uint = 4294967295
                                v4 = (v4 + v4)
                                _ = v21
                            default:
                                v18 = (v1 > 1)
                            }
                        }
                    default:
                        switch "foo bar" {
                        case "":
                            if (v18 && true) {
                                _ = len(v3)
                                if v18 {
                                    if true {
                                        v18 = ((!(((1e10 + -2.25) < ((0.0 + (0.0 - 1e10)) + 0.0)) && (!v18))) && (!(!((^(v5)) <= (^(v5))))))
                                        var v22 bool = false
                                        _ = v22
                                    }
                                    v5 = 255
                                    v1 = v1
                                }
                                if ((v1 | v1) > (v1 << ((42) & 63))) {
                                    _ = len(v3)
                                    switch 921 {
                                    case -2147483648:
                                        v1 += (v1 ^ (v1 << ((1) & 63)))
                                    default:
                                        v23 := GenS1{F0: v18, F1: v4, F2: -1}
                                        _ = v23
                                    }
                                    if true {
                                        var v24 int16 = 1
                                        v25 := GenS3{F0: 4294967295, F1: uint32(v5)}
                                        v26 := GenS3{F0: 1, F1: 2}
                                        if (1 > -1) {
                                            var v27 float32 = (1e10 / 3.14159)
                                            v28 := GenS4{F0: (v24 - (v24 & v24)), F1: 251, F2: int64(fi17)}
                                            v29 := GenS5{F0: ((^(v5)) < 255), F1: (v24 ^ v24)}
                                            var v30 int16 = v1
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        } else {
                                            var v31 bool = true
                                            var v32 uint32 = 0
                                            switch -542 {
                                            case 127:
                                                delete(v3, v32)
                                            case 9223372036854775807:
                                                switch fi14 {
                                                case -556:
                                                    v33 := v3[(v32 / ((v32) | 1))]
                                                    _ = v33
                                                case 495:
                                                    v5 = (v5 + (v5 * v5))
                                                default:
                                                    if v31 {
                                                        v34 := GenS5{F0: v31, F1: (v1 + (v24 << ((65535) & 63)))}
                                                        var v35 uint16 = 2
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                }
                                            case -384:
                                                v32 -= 793
                                            default:
                                                var v36 uint64 = 0
                                                _ = v36
                                            }
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = v24
                                        _ = v25
                                        _ = v26
                                    } else {
                                        var v37 int64 = 9223372036854775807
                                        _ = v37
                                    }
                                    v1 += v1
                                    v1 = v1
                                    v18 = false
                                }
                            }
                        case "hello":
                            _ = len(v2)
                        case "\n\t":
                            v4 += 42
                        }
                    }
                }
                v18 = v18
                _ = fi17
                _ = v18
            }
        }
        if false {
            v4 = v4
            _ = len(v2)
        } else {
            v38 := GenS6{F0: 883, F1: 42, F2: 1}
            _ = v38
        }
        delete(v3, 2147483647)
        _ = fi14
    }
    var v39 int32 = 779
    if (1 >= 2) {
        v5 = v5
    }
    switch "a" {
    case "a":
        if (0.1 != float64(3.14159)) {
            fmt.Println(v4, v4, v1)
            fmt.Println(v5, v5, v5)
            fmt.Println(v39)
        }
    case "foo bar":
        v3[uint32(v5)] = "hello"
    case "hello":
        if (((!(!false)) || ("" > (("foo bar" + ("" + "")) + ("a" + "foo bar")))) || true) {
            fmt.Println(v5, v1)
        }
    }
    switch "foo bar" {
    case "foo bar":
        v39 = 42
    case "":
        fmt.Println(v5)
    case "a":
        fmt.Println(v1)
    }
    v40 := GenS2{F0: 42}
    if (985 >= v39) {
        v39 = (v39 << ((5) & 63))
        var v41 uint32 = 0
        _ = v41
    } else {
        for fi42 := 0; fi42 < 8; fi42++ {
            switch ((("\n\t" + "a") + "") + (("" + "hello") + "foo bar")) {
            case "a":
                for fi43 := 0; fi43 < 13; fi43++ {
                    switch ("" + ("a" + "")) {
                    case "":
                        v44 := GenS3{F0: 182, F1: 2147483648}
                        _ = v44
                    case "hello":
                        if (32767 != (v1 - (v1 << ((883) & 63)))) {
                            if (false && (2 < 128)) {
                                v39 = v39
                            } else {
                                switch (fi43 + 128) {
                                case 0:
                                    v45 := GenS1{F0: (int16(v5) == v1), F1: 100, F2: v39}
                                    _ = v45
                                default:
                                    v4 = (v4 ^ (v4 - v4))
                                }
                                if false {
                                    var v46 int8 = (v4 >> ((255) & 63))
                                    _ = v46
                                }
                            }
                            var v47 uint64 = 42
                            if false {
                                v39 += v39
                                _ = len(v2)
                                var v48 int16 = (v1 ^ -124)
                                _ = v48
                            } else {
                                v49 := GenS7{F0: 1, F1: v4, F2: ("foo bar" + "hello")}
                                if false {
                                    v5 = v5
                                } else {
                                    var v50 float32 = (3.14159 - (0.1 / (1.5 - 1e10)))
                                    _ = v50
                                }
                                if false {
                                    v1 += (^(v1))
                                    var v51 float64 = float64(0.0)
                                    v5 += v5
                                    _ = v51
                                }
                                _ = v49
                            }
                            v47 = (^(v47))
                            switch fi43 {
                            case -893:
                                var v52 int16 = v1
                                _ = v52
                            default:
                                if len(v2) > 0 {
                                    v2[0] = ("a" + "foo bar")
                                }
                            }
                            _ = v47
                        } else {
                            _ = len(v2)
                            var v53 float64 = 1.5
                            fmt.Println(v4, v1, v4)
                            _ = v53
                        }
                    case "a":
                        switch fi42 {
                        case 9223372036854775807:
                            if false {
                                var v54 uint16 = (^(v5))
                                _ = v54
                            } else {
                                delete(v3, 0)
                                v4 = (v4 & v4)
                            }
                        case -1:
                            v39 -= v39
                            if true {
                                if (9223372036854775807 >= -1) {
                                    v5 = v5
                                    v39 = v39
                                }
                                fmt.Println(fi42, v1)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v55 int32 = (-(v39))
                                fmt.Println(v1, v39, fi43)
                                _ = v55
                            } else {
                                if true {
                                    v4 -= (v4 - v4)
                                }
                                switch "" {
                                case "hello":
                                    switch -9223372036854775808 {
                                    case 32768:
                                        fmt.Println(v5, fi43, v4)
                                    case 1:
                                        var v56 int = (fi42 - (-(fi43)))
                                        _ = v56
                                    case -2147483648:
                                        fmt.Println(v39, fi42)
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    fmt.Println(v4)
                                    switch "a" {
                                    case "hello":
                                        v3[4294967295] = "a"
                                    case "foo bar":
                                        v1 = (v1 ^ (v1 + (-(v1))))
                                        if true {
                                            v57 := GenS8{F0: (1 == 450), F1: 2}
                                            v3[2147483648] = ("hello" + (("foo bar" + "foo bar") + "a"))
                                            switch ("foo bar" + "a") {
                                            case "":
                                                v58 := GenS5{F0: (!(776 >= 172)), F1: (v1 >> ((2) & 63))}
                                                _ = v58
                                            }
                                            _ = v57
                                        }
                                    case "a":
                                        switch fi42 {
                                        case 2147483647:
                                            v39 += (v39 >> ((65535) & 63))
                                        }
                                    default:
                                        switch fi42 {
                                        case -476:
                                            v3[42] = (("\n\t" + ("hello" + "\n\t")) + "foo bar")
                                        case 253:
                                            if true {
                                                v1 -= 1000
                                            }
                                        default:
                                            if true {
                                                var v59 int16 = v1
                                                _ = v59
                                            }
                                        }
                                    }
                                }
                                v5 -= (v5 / ((v5) | 1))
                            }
                            if false {
                                v60 := v3[264]
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v5 = v5
                                v61 := GenS3{F0: 53, F1: 65535}
                                var v62 float32 = (3.14159 / 1e10)
                                var v63 int32 = v39
                                v39 = v39
                                _ = v60
                                _ = v61
                                _ = v62
                                _ = v63
                            } else {
                                if (!false) {
                                    v64 := GenS1{F0: false, F1: v4, F2: -1}
                                    if (false || (3.14159 > ((0.1 * ((3.14159 + (0.1 / -2.25)) + 0.1)) / 3.14159))) {
                                        _ = len(v3)
                                        var v65 float64 = (1e10 + 1.5)
                                        _ = v65
                                    }
                                    fmt.Println(v5)
                                    if (true || false) {
                                        var v66 int64 = 9223372036854775807
                                        v67 := GenS1{F0: false, F1: 100, F2: v39}
                                        fmt.Println(v66)
                                        if (42 != 216) {
                                            fmt.Println(fi43, v4)
                                            v39 = (v39 - int32(v1))
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = (("foo bar" + ("foo bar" + "")) + (("a" + "foo bar") + ""))
                                            }
                                            if true {
                                                switch "a" {
                                                case "hello":
                                                    v5 += (^(v5))
                                                case "":
                                                    v4 += 0
                                                case "foo bar":
                                                    var v68 float32 = -2.25
                                                    _ = v68
                                                default:
                                                    var v69 int32 = (^(v39))
                                                    _ = v69
                                                }
                                                var v70 bool = (!true)
                                                v71 := v3[4294967295]
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                switch fi42 {
                                                case -2147483648:
                                                    var v72 int = (^(fi43))
                                                    _ = v72
                                                case 4294967295:
                                                    v73 := GenS4{F0: 42, F1: 2, F2: 42}
                                                    fmt.Println(v1)
                                                    _ = v73
                                                default:
                                                    v74 := GenS5{F0: (v70 && ((!(0.0 < -2.25)) && v70)), F1: v1}
                                                    _ = v74
                                                }
                                                _ = v70
                                                _ = v71
                                            }
                                        }
                                        _ = v66
                                        _ = v67
                                    }
                                    _ = v64
                                } else {
                                    if false {
                                        var v75 int32 = int32(fi42)
                                        switch ("a" + "a") {
                                        case "a":
                                            var v76 float64 = -2.25
                                            _ = v76
                                        case "\n\t":
                                            v5 -= (v5 >> ((65535) & 63))
                                        default:
                                            switch fi42 {
                                            case -2147483648:
                                                var v77 int = (^(fi43))
                                                _ = v77
                                            case -9223372036854775808:
                                                if (!(!false)) {
                                                    var v78 bool = false
                                                    if true {
                                                        _ = len(v3)
                                                    }
                                                    _ = v78
                                                }
                                            }
                                        }
                                        _ = v75
                                    }
                                    var v79 uint = 4294967295
                                    _ = v79
                                }
                                var v80 int64 = 1
                                switch "foo bar" {
                                case "a":
                                    v81 := GenS9{F0: (true || (0 == (v80 / ((v80) | 1)))), F1: fi43, F2: 756}
                                    _ = v81
                                }
                                fmt.Println(fi42)
                                var v82 uint8 = 2
                                _ = v80
                                _ = v82
                            }
                            if false {
                                if (!true) {
                                    v83 := GenS10{F0: v39}
                                    var v84 uint8 = 255
                                    _ = v83
                                    _ = v84
                                } else {
                                    if true {
                                        var v85 int32 = v39
                                        fmt.Println(v39)
                                        _ = v85
                                    } else {
                                        var v86 uint32 = 65535
                                        _ = v86
                                    }
                                    _ = len(v2)
                                    if false {
                                        var v87 float64 = 0.1
                                        switch fi42 {
                                        case -9223372036854775808:
                                            var v88 string = ""
                                            switch 32768 {
                                            case 9223372036854775807:
                                                fmt.Println(fi43, v88, v4)
                                            }
                                            if ((v39 | (v39 ^ (v39 + (v39 % (((-(v39))) | 1))))) <= v39) {
                                                if true {
                                                    var v89 string = v88
                                                    _ = v89
                                                }
                                                v39 -= (^(v39))
                                                v88 = ((v88 + v88) + (v88 + v88))
                                            } else {
                                                switch ("hello" + "a") {
                                                case "a":
                                                    var v90 int32 = 634
                                                    _ = v90
                                                case "":
                                                    if true {
                                                        v91 := v3[42]
                                                        _ = v91
                                                    }
                                                case "\n\t":
                                                    var v92 uint16 = 158
                                                    _ = v92
                                                default:
                                                    if false {
                                                        fmt.Println(v4)
                                                        v93 := GenS4{F0: (v1 | -1), F1: 123, F2: 0}
                                                        _ = v93
                                                    } else {
                                                        var v94 int32 = v39
                                                        v88 = v88
                                                        var v95 uint8 = 255
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                }
                                                v4 += (v4 & 114)
                                                var v96 uint64 = genFn0(42)
                                                _ = v96
                                            }
                                            _ = v88
                                        case 32768:
                                            fmt.Println(v4, fi43)
                                            fmt.Println(v5)
                                        case 32767:
                                            var v97 int32 = (v39 % (((v39 >> ((491) & 63))) | 1))
                                            var v98 string = ((("" + (("a" + "foo bar") + "hello")) + "hello") + "a")
                                            var v99 int32 = v39
                                            _ = v97
                                            _ = v98
                                            _ = v99
                                        default:
                                            if false {
                                                if (205 != 4294967295) {
                                                    var v100 uint8 = 255
                                                    var v101 int16 = v1
                                                    var v102 float32 = 1.5
                                                    _ = v100
                                                    _ = v101
                                                    _ = v102
                                                }
                                                v1 += (v1 | (v1 | v1))
                                                v103 := GenS4{F0: v1, F1: 2147483648, F2: 42}
                                                v87 = v87
                                                _ = v103
                                            }
                                        }
                                        var v104 float32 = 1.5
                                        _ = v104
                                        _ = v87
                                    } else {
                                        var v105 int8 = v4
                                        _ = v105
                                    }
                                }
                                var v106 uint64 = 207
                                var v107 bool = (987 > v5)
                                _ = v106
                                _ = v107
                            } else {
                                var v108 int64 = 0
                                _ = v108
                            }
                            fmt.Println(v39, v1)
                        case -53:
                            if true {
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                                if false {
                                    if ((float32(-2.25) * 1e10) >= 1.5) {
                                        switch "hello" {
                                        case "\n\t":
                                            if (false && (-2.25 == 0.1)) {
                                                v39 = 568
                                                if false {
                                                    v4 -= -128
                                                    var v109 float64 = 1.5
                                                    if (-650 < 436) {
                                                        var v110 uint32 = 2147483648
                                                        v2 = append(v2, "foo bar")
                                                        _ = v110
                                                    } else {
                                                        switch ("hello" + ("hello" + (("" + (((((("hello" + "") + ("hello" + "a")) + "hello") + "foo bar") + "a") + ("a" + ("hello" + ("hello" + "a"))))) + (("" + "a") + ("hello" + "a"))))) {
                                                        case "hello":
                                                            fmt.Println(fi43, v5, v4)
                                                        case "":
                                                            if (42 == 343) {
                                                                switch (("a" + ("" + ("hello" + "hello"))) + "hello") {
                                                                case "foo bar":
                                                                    v111 := GenS10{F0: (v39 << ((65535) & 63))}
                                                                    _ = v111
                                                                case "\n\t":
                                                                    var v112 uint64 = 225
                                                                    _ = v112
                                                                default:
                                                                    delete(v3, 42)
                                                                }
                                                                v113 := GenS10{F0: (v39 << ((42) & 63))}
                                                                v114 := GenS1{F0: (1e10 == 3.14159), F1: 0, F2: v39}
                                                                if true {
                                                                    v2 = append(v2, "")
                                                                    v1 = v1
                                                                    var v115 uint16 = 1
                                                                    _ = v115
                                                                }
                                                                _ = len(v2)
                                                                switch "" {
                                                                case "\n\t":
                                                                    var v116 float64 = (v109 + -2.25)
                                                                    _ = v116
                                                                case "":
                                                                    _ = len(v3)
                                                                case "foo bar":
                                                                    v39 = v39
                                                                default:
                                                                    v4 += 100
                                                                }
                                                                switch (fi42 << ((65535) & 63)) {
                                                                case 2147483647:
                                                                    var v117 uint8 = 42
                                                                    _ = v117
                                                                case -1:
                                                                    var v118 int64 = 9223372036854775807
                                                                    fmt.Println(v1, v1)
                                                                    _ = v118
                                                                case -614:
                                                                    fmt.Println(v5)
                                                                }
                                                                switch (("" + "hello") + "\n\t") {
                                                                case "hello":
                                                                    fmt.Println(v4, v4, v39)
                                                                case "a":
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                case "\n\t":
                                                                    delete(v3, 2147483647)
                                                                default:
                                                                    if false {
                                                                        var v119 uint = 18446744073709551615
                                                                        if true {
                                                                            if ((v1 * (v1 >> (((v119 + 18446744073709551615)) & 63))) >= -32768) {
                                                                                v120 := GenS11{F0: v119}
                                                                                var v121 uint32 = 153
                                                                                fmt.Println(fi43, v39, v121)
                                                                                _ = v120
                                                                                _ = v121
                                                                            }
                                                                            v2 = append(v2, "\n\t")
                                                                            fmt.Println(v39, v39)
                                                                        }
                                                                        v122 := GenS12{F0: 42, F1: (-(v39)), F2: ("a" + (("\n\t" + ("hello" + ((("hello" + "hello") + (("hello" + "hello") + ("" + ("" + "hello")))) + ("hello" + "hello")))) + "a"))}
                                                                        var v123 uint32 = 917
                                                                        _ = v119
                                                                        _ = v122
                                                                        _ = v123
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = "hello"
                                                                        }
                                                                        fmt.Println(fi42, fi42, fi43)
                                                                    }
                                                                }
                                                                _ = v113
                                                                _ = v114
                                                            }
                                                        case "\n\t":
                                                            v124 := GenS13{F0: 2147483647, F1: ""}
                                                            _ = len(v3)
                                                            _ = v124
                                                        default:
                                                            fmt.Println(fi43, v1)
                                                        }
                                                        var v125 int64 = 42
                                                        _ = v125
                                                    }
                                                    _ = v109
                                                } else {
                                                    switch (-(fi43)) {
                                                    case -263:
                                                        var v126 uint = 65535
                                                        fmt.Println(v5, fi43, v39)
                                                        v5 -= (v5 - v5)
                                                        _ = v126
                                                    default:
                                                        v3[65535] = (("foo bar" + "") + "a")
                                                    }
                                                    var v127 float64 = (-2.25 + (3.14159 / 1e10))
                                                    var v128 uint32 = 65535
                                                    _ = v127
                                                    _ = v128
                                                }
                                                if (((((244 == 154) && false) || ((2 > 255) || true)) && (1e10 >= 1e10)) && true) {
                                                    fmt.Println(v39, v39)
                                                    v4 += v4
                                                    var v129 string = "\n\t"
                                                    _ = v129
                                                }
                                            }
                                        }
                                        var v130 bool = false
                                        _ = v130
                                    }
                                    var v131 float32 = (float32(0.0) * 1e10)
                                    v132 := GenS14{F0: 9223372036854775807, F1: (v1 + 32767), F2: (v4 & v4)}
                                    fmt.Println(v1)
                                    if false {
                                        var v133 uint16 = v5
                                        var v134 int32 = (v39 * 367)
                                        if (-9223372036854775808 != -1) {
                                            fmt.Println(v1)
                                        } else {
                                            fmt.Println(v1, v4)
                                            fmt.Println(v1, v1)
                                        }
                                        _ = v133
                                        _ = v134
                                    }
                                    var v135 float64 = 0.1
                                    var v136 int16 = v1
                                    fmt.Println(v39, v39, v5)
                                    _ = v131
                                    _ = v132
                                    _ = v135
                                    _ = v136
                                } else {
                                    v1 = -254
                                    var v137 int64 = -988
                                    _ = v137
                                }
                            }
                        }
                    }
                    var v138 int32 = (v39 / ((v39) | 1))
                    _ = fi43
                    _ = v138
                }
            case "":
                v5 -= v5
            case "hello":
                var v139 string = "a"
                _ = v139
            default:
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    var v140 uint = 1
                    _ = v140
                }
            }
            switch (("" + ("" + "\n\t")) + "hello") {
            case "\n\t":
                for fi141 := 0; fi141 < 2; fi141++ {
                    var v142 uint16 = v5
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    _ = fi141
                    _ = v142
                }
            case "":
                var v143 int16 = (^(v1))
                _ = v143
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi42
        }
    }
    v5 += v5
    v39 += (^(v39))
    if (-2.25 <= float64(1e10)) {
        var v144 string = "hello"
        v145 := GenS15{F0: (4294967295 >= 255)}
        _ = v144
        _ = v145
    }
    v146 := v3[1]
    var v147 int64 = -990
    var v148 int64 = (v147 | -845)
    var v149 string = ((v146 + "hello") + ("a" + v146))
    for fi150 := 0; fi150 < 16; fi150++ {
        var v151 uint64 = 2
        for fi152 := 0; fi152 < 5; fi152++ {
            var v153 uint64 = 1
            fmt.Println(v148, v4, v149)
            v154 := GenS16{F0: true, F1: 111}
            _ = fi152
            _ = v153
            _ = v154
        }
        var v155 int64 = v148
        _ = fi150
        _ = v151
        _ = v155
    }
    v146 = v146
    fmt.Println("v1:", v1)
    fmt.Println("v146:", v146)
    fmt.Println("v147:", v147)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v146
    _ = v147
    _ = v148
    _ = v149
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
