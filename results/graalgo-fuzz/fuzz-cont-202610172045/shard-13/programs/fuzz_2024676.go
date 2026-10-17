package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int8; F2 uint8 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint32; F1 string; F2 uint32 }
type GenS7 struct { F0 uint32; F1 int64 }
type GenS8 struct { F0 bool; F1 string; F2 int }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 uint16; F1 uint16 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS14 struct { F0 int16; F1 int16; F2 int64 }
type GenS15 struct { F0 uint32; F1 uint16; F2 bool }
type GenS16 struct { F0 string; F1 string }
type GenS17 struct { F0 int8; F1 uint32 }
type GenS18 struct { F0 int16 }
func genFn0(p0 string, p1 uint) uint64 {
    p1 += 2
    var v7 int8 = 0
    var v8 int64 = -298
    var v9 int8 = (v7 << (((p1 ^ (p1 * (^(p1))))) & 63))
    var v10 int8 = (-(v9))
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 string) uint {
    if false {
        p0 = p0
    } else {
        for fi11 := 0; fi11 < 2; fi11++ {
            for fi12 := 0; fi12 < 18; fi12++ {
                fmt.Println(fi11, p0, p0)
                v13 := GenS1{F0: (int8(fi11) <= 100), F1: fi11}
                _ = fi12
                _ = v13
            }
            p0 = (((p0 + p0) + (p0 + ((p0 + p0) + p0))) + p0)
            v14 := []int{(fi11 + (-(fi11))), (fi11 << ((18446744073709551615) & 63)), (fi11 >> ((145) & 63)), (fi11 << ((65535) & 63))}
            _ = fi11
            _ = v14
        }
        for fi15 := 0; fi15 < 16; fi15++ {
            v16 := []int{(fi15 >> ((42) & 63))}
            _ = fi15
            _ = v16
        }
        var v17 uint64 = 1
        _ = v17
    }
    if false {
        var v18 uint16 = 0
        if (3.14159 < 0.0) {
            v19 := []int{32768, -282, 0, -9223372036854775808}
            v20 := GenS1{F0: false, F1: -98}
            _ = v19
            _ = v20
        } else {
            v18 -= v18
            if ((4294967295 != 65535) || true) {
                if (745 == 32767) {
                    switch 0 {
                    case 127:
                        for fi21 := 0; fi21 < 4; fi21++ {
                            fmt.Println(fi21, p0)
                            _ = fi21
                        }
                    case 862:
                        v22 := GenS2{F0: 304}
                        _ = v22
                    case 9223372036854775807:
                        var v23 int32 = 962
                        _ = v23
                    }
                    if true {
                        for fi24 := 0; fi24 < 14; fi24++ {
                            if false {
                                v25 := []int{-148, 251, (fi24 | fi24), fi24}
                                _ = v25
                            } else {
                                var v26 uint16 = 1
                                if false {
                                    p0 = ("foo bar" + (("foo bar" + "hello") + p0))
                                    switch fi24 {
                                    case -9223372036854775808:
                                        var v27 int64 = 9223372036854775807
                                        _ = v27
                                    case 128:
                                        switch "hello" {
                                        case "foo bar":
                                            v18 -= (v18 >> ((988) & 63))
                                        case "":
                                            switch (fi24 % ((fi24) | 1)) {
                                            case -9223372036854775808:
                                                fmt.Println(v26, fi24, fi24)
                                            }
                                        case "a":
                                            var v28 uint16 = (^(v18))
                                            _ = v28
                                        }
                                    case 1:
                                        if (false || false) {
                                            for fi29 := 0; fi29 < 18; fi29++ {
                                                if (0 >= 852) {
                                                    switch (^(fi24)) {
                                                    case 577:
                                                        switch p0 {
                                                        case "":
                                                            if (574 >= 2) {
                                                                var v30 uint32 = 0
                                                                var v31 uint64 = 319
                                                                _ = v30
                                                                _ = v31
                                                            }
                                                        case "a":
                                                            v32 := GenS3{F0: -884}
                                                            _ = v32
                                                        case "hello":
                                                            v26 = v18
                                                        default:
                                                            v33 := []int{fi29, fi29, (fi29 * (fi24 | fi24))}
                                                            _ = v33
                                                        }
                                                    case 754:
                                                        if false {
                                                            v34 := []int{fi29, 744, fi29}
                                                            if true {
                                                                v18 -= 42
                                                                v26 = v26
                                                                if (!(!true)) {
                                                                    switch ((p0 + (p0 + p0)) + (p0 + p0)) {
                                                                    case "foo bar":
                                                                        v35 := GenS2{F0: 287}
                                                                        _ = v35
                                                                    case "a":
                                                                        if false {
                                                                            switch ("\n\t" + "\n\t") {
                                                                            case "hello":
                                                                                v36 := GenS4{F0: 1, F1: -1, F2: 1}
                                                                                _ = v36
                                                                            case "foo bar":
                                                                                if true {
                                                                                    v18 = 244
                                                                                    fmt.Println(v26, fi29, fi29)
                                                                                } else {
                                                                                    if false {
                                                                                        if true {
                                                                                            v37 := GenS3{F0: 0}
                                                                                            if len(v34) >= 2 {
                                                                                                v34 = v34[1:len(v34)]
                                                                                            }
                                                                                            v26 = (v26 % ((v26) | 1))
                                                                                            _ = v37
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            v38 := GenS4{F0: 0, F1: 75, F2: 0}
                                                                            var v39 bool = ((1.5 / 3.14159) <= (1e10 - 3.14159))
                                                                            v40 := GenS4{F0: -1, F1: -128, F2: 205}
                                                                            _ = v38
                                                                            _ = v39
                                                                            _ = v40
                                                                        }
                                                                    case "hello":
                                                                        if (!(!(!(true && false)))) {
                                                                            switch ("a" + p0) {
                                                                            case "hello":
                                                                                v41 := GenS5{F0: (128 != 1), F1: 846}
                                                                                _ = v41
                                                                            case "\n\t":
                                                                                v18 += 32767
                                                                            case "foo bar":
                                                                                if len(v34) >= 2 {
                                                                                    v34 = v34[1:len(v34)]
                                                                                }
                                                                            }
                                                                            v18 -= (v26 | (v26 - v18))
                                                                            fmt.Println(v26)
                                                                        } else {
                                                                            v34 = append(v34, fi24)
                                                                            if false {
                                                                                fmt.Println(p0, v18)
                                                                                v18 = (v26 - v18)
                                                                                v18 = 255
                                                                            } else {
                                                                                var v42 int8 = -128
                                                                                v34 = append(v34, -820)
                                                                                _ = v42
                                                                            }
                                                                            if (1 == 0) {
                                                                                v43 := GenS4{F0: 0, F1: 100, F2: 42}
                                                                                _ = v43
                                                                            } else {
                                                                                if (!(true || (!(1.5 > (0.1 - (0.0 - ((1.5 * -2.25) + (3.14159 * float64(0.1))))))))) {
                                                                                    v26 = (v26 ^ v18)
                                                                                }
                                                                                var v44 uint32 = 2
                                                                                p0 = (p0 + "hello")
                                                                                v26 = v26
                                                                                if true {
                                                                                    switch (fi29 + fi24) {
                                                                                    case -885:
                                                                                        if false {
                                                                                            v18 = (v18 >> ((65535) & 63))
                                                                                            if false {
                                                                                                var v45 int64 = 42
                                                                                                if (((3.14159 + ((float32(0.1) * ((-2.25 + (float32(0.1) + -2.25)) / -2.25)) + 1e10)) - (1.5 / (1.5 * 0.1))) >= 1e10) {
                                                                                                    var v46 uint = 437
                                                                                                    var v47 int8 = -87
                                                                                                    fmt.Println(v18)
                                                                                                    _ = v46
                                                                                                    _ = v47
                                                                                                } else {
                                                                                                    fmt.Println(p0, v45)
                                                                                                    var v48 string = p0
                                                                                                    _ = v48
                                                                                                }
                                                                                                _ = v45
                                                                                            } else {
                                                                                                var v49 int32 = 42
                                                                                                v50 := GenS2{F0: genFn0(p0, 723)}
                                                                                                switch "hello" {
                                                                                                case "\n\t":
                                                                                                    if len(v34) > 0 {
                                                                                                        v34[0] = (fi29 << ((42) & 63))
                                                                                                    }
                                                                                                }
                                                                                                var v51 uint32 = v44
                                                                                                if true {
                                                                                                    v26 += (^(v18))
                                                                                                    v52 := GenS4{F0: 42, F1: 1, F2: 1}
                                                                                                    _ = v52
                                                                                                }
                                                                                                p0 = p0
                                                                                                switch (fi29 * fi29) {
                                                                                                case 32767:
                                                                                                    switch fi29 {
                                                                                                    case 1:
                                                                                                        var v53 int = (fi29 & fi24)
                                                                                                        _ = v53
                                                                                                    case -9223372036854775808:
                                                                                                        v18 -= v18
                                                                                                    default:
                                                                                                        if (!(v49 <= v49)) {
                                                                                                            fmt.Println(v51, v44, fi29)
                                                                                                        } else {
                                                                                                            var v54 int8 = 127
                                                                                                            v55 := GenS5{F0: (1 >= 127), F1: 0}
                                                                                                            fmt.Println(v18, v18)
                                                                                                            _ = v54
                                                                                                            _ = v55
                                                                                                        }
                                                                                                    }
                                                                                                case 9223372036854775807:
                                                                                                    _ = len(v34)
                                                                                                }
                                                                                                _ = v49
                                                                                                _ = v50
                                                                                                _ = v51
                                                                                            }
                                                                                        } else {
                                                                                            v56 := GenS1{F0: (((!true) || false) && false), F1: 128}
                                                                                            var v57 int64 = int64(fi29)
                                                                                            v57 += (v57 >> ((387) & 63))
                                                                                            _ = v56
                                                                                            _ = v57
                                                                                        }
                                                                                    case 0:
                                                                                        if (!false) {
                                                                                            v26 += (^(v26))
                                                                                        }
                                                                                    }
                                                                                    switch fi29 {
                                                                                    case -2147483648:
                                                                                        if true {
                                                                                            if true {
                                                                                                v58 := GenS2{F0: 42}
                                                                                                fmt.Println(v18, p0, v18)
                                                                                                _ = v58
                                                                                            } else {
                                                                                                p0 = "hello"
                                                                                            }
                                                                                        }
                                                                                    case 32768:
                                                                                        var v59 int32 = 0
                                                                                        _ = v59
                                                                                    case 127:
                                                                                        v18 = 255
                                                                                        p0 = ""
                                                                                        switch (fi29 + (^(fi29))) {
                                                                                        case -426:
                                                                                            _ = len(v34)
                                                                                        case 109:
                                                                                            fmt.Println(v18, v26, fi29)
                                                                                        }
                                                                                    default:
                                                                                        p0 = p0
                                                                                        p0 = "foo bar"
                                                                                    }
                                                                                    fmt.Println(p0, fi24, p0)
                                                                                } else {
                                                                                    var v60 float32 = 0.0
                                                                                    if true {
                                                                                        var v61 uint32 = v44
                                                                                        _ = v61
                                                                                    } else {
                                                                                        if len(v34) >= 2 {
                                                                                            v34 = v34[1:len(v34)]
                                                                                        }
                                                                                        if false {
                                                                                            var v62 float32 = float32((3.14159 - float64(v60)))
                                                                                            var v63 int16 = -32768
                                                                                            if (4294967295 > 18446744073709551615) {
                                                                                                if (-622 > -1) {
                                                                                                    v60 = -2.25
                                                                                                }
                                                                                            } else {
                                                                                                var v64 uint = 539
                                                                                                _ = v64
                                                                                            }
                                                                                            fmt.Println(p0)
                                                                                            _ = v62
                                                                                            _ = v63
                                                                                        }
                                                                                        v65 := GenS6{F0: v44, F1: "\n\t", F2: (v44 / ((v44) | 1))}
                                                                                        _ = v65
                                                                                    }
                                                                                    switch p0 {
                                                                                    case "\n\t":
                                                                                        v66 := GenS6{F0: (v44 | 2147483647), F1: ("\n\t" + "foo bar"), F2: (^(v44))}
                                                                                        _ = v66
                                                                                    case "":
                                                                                        p0 = p0
                                                                                    case "hello":
                                                                                        if true {
                                                                                            var v67 bool = false
                                                                                            var v68 int16 = -713
                                                                                            _ = v67
                                                                                            _ = v68
                                                                                        }
                                                                                    default:
                                                                                        v44 = v44
                                                                                        v69 := GenS7{F0: (v44 | v44), F1: -9223372036854775808}
                                                                                        switch (((("\n\t" + p0) + ("" + "a")) + "foo bar") + p0) {
                                                                                        case "\n\t":
                                                                                            v70 := GenS3{F0: -216}
                                                                                            _ = v70
                                                                                        default:
                                                                                            if false {
                                                                                                switch (("foo bar" + ("" + (p0 + "\n\t"))) + (("foo bar" + "hello") + "hello")) {
                                                                                                case "a":
                                                                                                    if (!false) {
                                                                                                        var v71 uint8 = 128
                                                                                                        if len(v34) > 0 {
                                                                                                            v34[0] = (-(fi29))
                                                                                                        }
                                                                                                        _ = v71
                                                                                                    }
                                                                                                }
                                                                                                var v72 uint32 = (^(v44))
                                                                                                if len(v34) >= 2 {
                                                                                                    v34 = v34[1:len(v34)]
                                                                                                }
                                                                                                _ = v72
                                                                                            } else {
                                                                                                var v73 uint16 = v26
                                                                                                _ = v73
                                                                                            }
                                                                                        }
                                                                                        _ = v69
                                                                                    }
                                                                                    _ = v60
                                                                                }
                                                                                var v74 string = "a"
                                                                                _ = v44
                                                                                _ = v74
                                                                            }
                                                                        }
                                                                        v26 += v18
                                                                    }
                                                                } else {
                                                                    switch p0 {
                                                                    case "":
                                                                        v75 := GenS1{F0: (!false), F1: (^(fi29))}
                                                                        _ = v75
                                                                    }
                                                                    switch (p0 + p0) {
                                                                    case "hello":
                                                                        var v76 uint32 = 1
                                                                        _ = v76
                                                                    case "":
                                                                        p0 = p0
                                                                    case "a":
                                                                        v77 := GenS8{F0: (3.14159 >= 3.14159), F1: (p0 + ("\n\t" + "foo bar")), F2: fi29}
                                                                        var v78 int = (fi24 * (fi24 << ((1) & 63)))
                                                                        v26 += (v26 / ((v26) | 1))
                                                                        v79 := GenS9{F0: 127, F1: -165}
                                                                        fmt.Println(fi24, v26)
                                                                        var v80 float32 = (0.0 * 0.1)
                                                                        v81 := GenS10{F0: (v26 + (v18 ^ (v18 ^ 65535))), F1: (v18 & v18)}
                                                                        var v82 int8 = 90
                                                                        _ = v77
                                                                        _ = v78
                                                                        _ = v79
                                                                        _ = v80
                                                                        _ = v81
                                                                        _ = v82
                                                                    default:
                                                                        fmt.Println(fi24, fi29, fi29)
                                                                    }
                                                                }
                                                                switch "" {
                                                                case "\n\t":
                                                                    if len(v34) >= 2 {
                                                                        v34 = v34[1:len(v34)]
                                                                    }
                                                                case "":
                                                                    var v83 float64 = (1.5 - (((0.0 * -2.25) / 1e10) - 3.14159))
                                                                    _ = v83
                                                                }
                                                                var v84 int32 = 2147483647
                                                                _ = v84
                                                            } else {
                                                                switch (("hello" + p0) + (p0 + p0)) {
                                                                case "":
                                                                    v18 = 32767
                                                                default:
                                                                    if false {
                                                                        if (1.5 <= 1.5) {
                                                                            v18 -= v26
                                                                            if len(v34) >= 2 {
                                                                                v34 = v34[1:len(v34)]
                                                                            }
                                                                            fmt.Println(fi29, v26, v26)
                                                                            var v85 int = fi29
                                                                            v86 := GenS5{F0: ((^(v18)) > 1), F1: 147}
                                                                            v87 := GenS10{F0: (v18 >> ((42) & 63)), F1: 1}
                                                                            _ = v85
                                                                            _ = v86
                                                                            _ = v87
                                                                        } else {
                                                                            var v88 uint16 = 820
                                                                            _ = v88
                                                                        }
                                                                        v26 = (^(v18))
                                                                        v34 = append(v34, (-(fi29)))
                                                                    } else {
                                                                        var v89 uint64 = 42
                                                                        v26 = v26
                                                                        _ = v89
                                                                    }
                                                                }
                                                                if (!false) {
                                                                    if false {
                                                                        fmt.Println(p0, fi29, fi24)
                                                                        if ((299 < 2147483647) && (true && (v18 < (v26 | (v26 + (v18 - (^(v18)))))))) {
                                                                            switch p0 {
                                                                            case "hello":
                                                                                fmt.Println(v26, v26)
                                                                            case "":
                                                                                p0 = (p0 + "\n\t")
                                                                            default:
                                                                                switch (fi29 ^ (-(fi24))) {
                                                                                case -1:
                                                                                    _ = len(v34)
                                                                                case 446:
                                                                                    v18 = v26
                                                                                    p0 = (p0 + "hello")
                                                                                    v90 := GenS11{F0: (^(v18))}
                                                                                    _ = v90
                                                                                case 127:
                                                                                    if ((true || (!true)) || (!false)) {
                                                                                        var v91 int16 = 1000
                                                                                        _ = len(v34)
                                                                                        var v92 int = (fi24 + (fi29 << ((391) & 63)))
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                    } else {
                                                                                        if (-9223372036854775808 >= (fi24 / ((fi24) | 1))) {
                                                                                            var v93 int16 = -32768
                                                                                            fmt.Println(v18)
                                                                                            v94 := GenS7{F0: 2147483648, F1: int64(fi24)}
                                                                                            p0 = (p0 + "hello")
                                                                                            var v95 int16 = (v93 >> ((2) & 63))
                                                                                            v96 := GenS12{F0: (!(!(-206 < -9223372036854775808)))}
                                                                                            switch "foo bar" {
                                                                                            case "\n\t":
                                                                                                if (!true) {
                                                                                                    if len(v34) > 0 {
                                                                                                        v34[0] = fi29
                                                                                                    }
                                                                                                } else {
                                                                                                    if true {
                                                                                                        var v97 int8 = 1
                                                                                                        v98 := GenS13{F0: v93, F1: 140, F2: 4294967295}
                                                                                                        v26 += (^(v26))
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                    } else {
                                                                                                        v93 = v93
                                                                                                        var v99 int32 = -2147483648
                                                                                                        var v100 uint32 = 4294967295
                                                                                                        _ = v100
                                                                                                        _ = v99
                                                                                                    }
                                                                                                    if (!(false && false)) {
                                                                                                        if len(v34) >= 2 {
                                                                                                            v34 = v34[1:len(v34)]
                                                                                                        }
                                                                                                        if false {
                                                                                                            var v101 int = (-(fi29))
                                                                                                            if true {
                                                                                                                v18 -= (v18 | 0)
                                                                                                                var v102 uint = 912
                                                                                                                var v103 uint32 = 2147483647
                                                                                                                _ = v102
                                                                                                                _ = v103
                                                                                                            }
                                                                                                            var v104 int16 = (v93 >> ((18446744073709551615) & 63))
                                                                                                            if len(v34) >= 2 {
                                                                                                                v34 = v34[1:len(v34)]
                                                                                                            }
                                                                                                            if true {
                                                                                                                p0 = p0
                                                                                                            }
                                                                                                            _ = v101
                                                                                                            _ = v104
                                                                                                        }
                                                                                                    }
                                                                                                    switch fi24 {
                                                                                                    case 0:
                                                                                                        var v105 float32 = 1.5
                                                                                                        _ = v105
                                                                                                    case 1:
                                                                                                        v26 += v26
                                                                                                    }
                                                                                                    v95 += (v93 >> ((65535) & 63))
                                                                                                }
                                                                                            default:
                                                                                                v26 -= v18
                                                                                            }
                                                                                            _ = v93
                                                                                            _ = v94
                                                                                            _ = v95
                                                                                            _ = v96
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            p0 = ((((("foo bar" + ((p0 + ("a" + (p0 + "a"))) + "")) + "foo bar") + p0) + (p0 + ((("a" + ("a" + p0)) + ("a" + (p0 + p0))) + "a"))) + ("a" + (p0 + ("" + p0))))
                                                                            fmt.Println(p0, fi29)
                                                                        } else {
                                                                            v26 = 527
                                                                            switch p0 {
                                                                            case "hello":
                                                                                var v106 float64 = 1e10
                                                                                _ = v106
                                                                            case "foo bar":
                                                                                var v107 float32 = -2.25
                                                                                _ = v107
                                                                            case "\n\t":
                                                                                var v108 string = "\n\t"
                                                                                v18 += v26
                                                                                v108 = v108
                                                                                _ = v108
                                                                            }
                                                                            var v109 uint32 = 42
                                                                            _ = v109
                                                                        }
                                                                    } else {
                                                                        fmt.Println(p0, fi24)
                                                                    }
                                                                    if (false || true) {
                                                                        fmt.Println(p0, fi24, fi24)
                                                                        var v110 uint32 = 2
                                                                        switch 1 {
                                                                        case 9223372036854775807:
                                                                            var v111 uint64 = 42
                                                                            _ = v111
                                                                        case 127:
                                                                            var v112 string = p0
                                                                            _ = v112
                                                                        default:
                                                                            fmt.Println(v110, p0, p0)
                                                                        }
                                                                        _ = v110
                                                                    } else {
                                                                        if true {
                                                                            v113 := GenS14{F0: 42, F1: 755, F2: 421}
                                                                            p0 = (p0 + (p0 + (("" + (("\n\t" + (((p0 + "a") + "a") + "a")) + ((p0 + p0) + (((p0 + (p0 + p0)) + (("a" + (p0 + p0)) + p0)) + (p0 + p0))))) + p0)))
                                                                            _ = v113
                                                                        }
                                                                        var v114 int = -1
                                                                        _ = v114
                                                                    }
                                                                    v115 := GenS5{F0: (-2.25 > 0.0), F1: 4294967295}
                                                                    _ = v115
                                                                }
                                                            }
                                                            _ = v34
                                                        }
                                                    default:
                                                        fmt.Println(v18, v26)
                                                    }
                                                    if (929 < 32767) {
                                                        if (!(false || false)) {
                                                            p0 = ((("hello" + "foo bar") + p0) + ((("" + "hello") + "\n\t") + ""))
                                                            var v116 float64 = (3.14159 / float64(0.1))
                                                            v116 = v116
                                                            _ = v116
                                                        } else {
                                                            if (genFn0("", 1) >= 2) {
                                                                var v117 float64 = 1.5
                                                                _ = v117
                                                            } else {
                                                                v26 = v18
                                                                fmt.Println(v26, fi24, v26)
                                                            }
                                                        }
                                                        p0 = "a"
                                                    } else {
                                                        if true {
                                                            var v118 bool = (-529 == -461)
                                                            _ = v118
                                                        } else {
                                                            fmt.Println(fi29, fi29)
                                                            fmt.Println(p0, fi24, v26)
                                                            v119 := []int{}
                                                            _ = v119
                                                        }
                                                        var v120 string = "hello"
                                                        _ = v120
                                                    }
                                                } else {
                                                    switch ("hello" + (p0 + "a")) {
                                                    case "hello":
                                                        v26 = v18
                                                    case "\n\t":
                                                        var v121 uint64 = 0
                                                        _ = v121
                                                    case "":
                                                        v122 := GenS8{F0: false, F1: p0, F2: fi29}
                                                        _ = v122
                                                    }
                                                    v26 = v26
                                                }
                                                _ = fi29
                                            }
                                            if true {
                                                if true {
                                                    v123 := []int{}
                                                    _ = v123
                                                } else {
                                                    for fi124 := 0; fi124 < 6; fi124++ {
                                                        var v125 uint64 = 42
                                                        switch (p0 + p0) {
                                                        case "foo bar":
                                                            v26 = (v26 & v26)
                                                        default:
                                                            switch "" {
                                                            case "hello":
                                                                v125 += genFn0(p0, 529)
                                                                v126 := GenS6{F0: 389, F1: (p0 + p0), F2: 208}
                                                                _ = v126
                                                            }
                                                        }
                                                        var v127 float32 = ((((float32(3.14159) / 1.0) + (1.5 - 1e10)) * -2.25) - float32(((3.14159 - 0.0) - 1.5)))
                                                        _ = fi124
                                                        _ = v125
                                                        _ = v127
                                                    }
                                                    if true {
                                                        v128 := []int{fi24}
                                                        if false {
                                                            var v129 bool = false
                                                            var v130 uint16 = (^(v18))
                                                            if len(v128) >= 2 {
                                                                v128 = v128[1:len(v128)]
                                                            }
                                                            _ = v129
                                                            _ = v130
                                                        } else {
                                                            var v131 uint32 = 2147483647
                                                            if true {
                                                                fmt.Println(fi24, v26)
                                                                if false {
                                                                    v18 = (v26 - 875)
                                                                    if true {
                                                                        var v132 uint64 = 18446744073709551615
                                                                        v26 += (v26 - v18)
                                                                        _ = len(v128)
                                                                        _ = v132
                                                                    } else {
                                                                        if (true && ((fi24 >> ((390) & 63)) != fi24)) {
                                                                            if false {
                                                                                var v133 uint32 = 65535
                                                                                v131 = (v133 * v133)
                                                                                _ = v133
                                                                            }
                                                                        } else {
                                                                            var v134 string = ""
                                                                            p0 = v134
                                                                            _ = v134
                                                                        }
                                                                    }
                                                                } else {
                                                                    fmt.Println(v26, v131, p0)
                                                                    if false {
                                                                        fmt.Println(v26)
                                                                        for fi135 := 0; fi135 < 3; fi135++ {
                                                                            var v136 uint32 = v131
                                                                            var v137 bool = (true || true)
                                                                            _ = fi135
                                                                            _ = v136
                                                                            _ = v137
                                                                        }
                                                                        fmt.Println(v26, fi24, v18)
                                                                    }
                                                                }
                                                                v18 -= (v18 & 348)
                                                            } else {
                                                                for fi138 := 0; fi138 < 10; fi138++ {
                                                                    var v139 uint8 = 255
                                                                    if true {
                                                                        if false {
                                                                            var v140 string = (((p0 + p0) + (p0 + p0)) + (p0 + "hello"))
                                                                            v139 = v139
                                                                            v128 = append(v128, (fi138 | (fi24 & int(v131))))
                                                                            _ = v140
                                                                        } else {
                                                                            if (v139 >= (v139 >> ((uint(v139)) & 63))) {
                                                                                v128 = append(v128, fi24)
                                                                                var v141 uint32 = v131
                                                                                _ = v141
                                                                            } else {
                                                                                switch (fi138 >> ((uint(v131)) & 63)) {
                                                                                case -1:
                                                                                    v142 := GenS15{F0: (^(v131)), F1: 2, F2: (v131 >= 270)}
                                                                                    _ = v142
                                                                                }
                                                                                var v143 uint8 = (v139 >> ((255) & 63))
                                                                                _ = v143
                                                                            }
                                                                        }
                                                                        switch p0 {
                                                                        case "hello":
                                                                            fmt.Println(v139, fi24)
                                                                        default:
                                                                            v144 := GenS10{F0: v26, F1: v26}
                                                                            _ = v144
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v131, fi24)
                                                                    }
                                                                    _ = fi138
                                                                    _ = v139
                                                                }
                                                                var v145 uint8 = 128
                                                                var v146 uint32 = (v131 | (v131 + (^(v131))))
                                                                _ = v145
                                                                _ = v146
                                                            }
                                                            _ = len(v128)
                                                            _ = v131
                                                        }
                                                        _ = v128
                                                    } else {
                                                        p0 = ""
                                                        var v147 uint8 = 128
                                                        _ = v147
                                                    }
                                                }
                                            }
                                        } else {
                                            var v148 int32 = 1
                                            if (true || (!(4294967295 <= 2))) {
                                                switch p0 {
                                                case "foo bar":
                                                    fmt.Println(p0)
                                                }
                                                v26 = v26
                                            } else {
                                                for fi149 := 0; fi149 < 8; fi149++ {
                                                    switch ("foo bar" + ("\n\t" + "")) {
                                                    case "":
                                                        if true {
                                                            var v150 int8 = 127
                                                            if true {
                                                                fmt.Println(v26)
                                                                if true {
                                                                    if (-2.25 > float32(3.14159)) {
                                                                        fmt.Println(v26, v18)
                                                                    }
                                                                    v151 := GenS4{F0: 0, F1: 0, F2: 127}
                                                                    _ = v151
                                                                }
                                                            }
                                                            _ = v150
                                                        } else {
                                                            v152 := []int{fi24, 2147483647}
                                                            if false {
                                                                v26 = v26
                                                                var v153 uint64 = 4294967295
                                                                var v154 float64 = 3.14159
                                                                _ = v153
                                                                _ = v154
                                                            } else {
                                                                v155 := GenS5{F0: true, F1: 532}
                                                                var v156 uint8 = 60
                                                                _ = v155
                                                                _ = v156
                                                            }
                                                            switch "" {
                                                            case "":
                                                                if ((436 >= fi24) || (!(!(int(v26) >= (^(fi149)))))) {
                                                                    p0 = p0
                                                                } else {
                                                                    v18 -= v26
                                                                    v157 := GenS8{F0: (255 >= 0), F1: "", F2: fi24}
                                                                    v26 -= v18
                                                                    _ = v157
                                                                }
                                                            }
                                                            _ = v152
                                                        }
                                                    default:
                                                        v158 := []int{-1, 91}
                                                        v158 = append(v158, -321)
                                                        if true {
                                                            if (!false) {
                                                                if len(v158) >= 2 {
                                                                    v158 = v158[1:len(v158)]
                                                                }
                                                                v159 := GenS13{F0: 0, F1: 127, F2: 2}
                                                                _ = v159
                                                            }
                                                            v26 += v18
                                                        } else {
                                                            if (p0 <= "\n\t") {
                                                                if (v26 <= v26) {
                                                                    if (true && (fi24 == (fi149 % ((fi24) | 1)))) {
                                                                        var v160 int32 = (v148 ^ (v148 % (((v148 & (v148 ^ 2147483647))) | 1)))
                                                                        _ = v160
                                                                    }
                                                                    p0 = ""
                                                                }
                                                                if len(v158) > 0 {
                                                                    v158[0] = fi149
                                                                }
                                                                fmt.Println(fi24)
                                                                var v161 int = (fi24 - 1)
                                                                if false {
                                                                    v158 = append(v158, (fi24 & 248))
                                                                    var v162 int64 = -9223372036854775808
                                                                    _ = v162
                                                                } else {
                                                                    v161 += (fi24 * fi24)
                                                                }
                                                                _ = v161
                                                            }
                                                            if true {
                                                                v163 := GenS16{F0: (((p0 + ("" + "foo bar")) + ("foo bar" + p0)) + p0), F1: ((p0 + p0) + p0)}
                                                                v158 = append(v158, (fi149 & (fi24 >> ((255) & 63))))
                                                                p0 = "a"
                                                                _ = v163
                                                            } else {
                                                                p0 = "foo bar"
                                                                v148 -= 1
                                                                p0 = p0
                                                            }
                                                        }
                                                        _ = v158
                                                    }
                                                    if (592 >= v18) {
                                                        v164 := []int{}
                                                        if len(v164) > 0 {
                                                            v164[0] = (^(fi149))
                                                        }
                                                        var v165 uint = 4294967295
                                                        _ = v164
                                                        _ = v165
                                                    } else {
                                                        v166 := []int{}
                                                        _ = v166
                                                    }
                                                    if false {
                                                        var v167 int8 = int8(v18)
                                                        var v168 int32 = 1
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            fmt.Println(v18, v168)
                                                        }
                                                        if false {
                                                            v169 := GenS3{F0: 822}
                                                            v170 := []int{fi149, fi24}
                                                            _ = v169
                                                            _ = v170
                                                        } else {
                                                            v168 -= (v168 / ((v148) | 1))
                                                        }
                                                        _ = v167
                                                        _ = v168
                                                    } else {
                                                        switch (-(fi24)) {
                                                        case -1:
                                                            switch p0 {
                                                            case "foo bar":
                                                                p0 = "hello"
                                                            case "":
                                                                v171 := GenS13{F0: 0, F1: 2, F2: 521}
                                                                _ = v171
                                                            }
                                                        case 0:
                                                            v18 = (v18 >> ((969) & 63))
                                                        case 2147483647:
                                                            if (!true) {
                                                                fmt.Println(fi24, fi24, v18)
                                                                var v172 uint16 = (^(v18))
                                                                v173 := []int{fi24, (fi24 | (fi24 ^ (fi149 >> ((18446744073709551615) & 63)))), -825, 2147483647}
                                                                _ = v172
                                                                _ = v173
                                                            }
                                                            v174 := GenS15{F0: 2147483647, F1: v18, F2: false}
                                                            if (((v26 == v26) || (-9223372036854775808 != 1)) || false) {
                                                                v26 -= 42
                                                            }
                                                            _ = v174
                                                        }
                                                        fmt.Println(p0, v18)
                                                        if true {
                                                            var v175 uint = 4294967295
                                                            var v176 string = p0
                                                            _ = v175
                                                            _ = v176
                                                        }
                                                    }
                                                    v18 += (v26 - v26)
                                                    _ = fi149
                                                }
                                                var v177 string = "a"
                                                var v178 string = p0
                                                _ = v177
                                                _ = v178
                                            }
                                            _ = v148
                                        }
                                    default:
                                        v179 := GenS3{F0: 42}
                                        _ = v179
                                    }
                                } else {
                                    switch "a" {
                                    case "a":
                                        fmt.Println(v18, v26, p0)
                                    case "":
                                        v26 -= v26
                                    case "\n\t":
                                        switch p0 {
                                        case "foo bar":
                                            var v180 uint8 = 1
                                            _ = v180
                                        default:
                                            var v181 int16 = -32768
                                            _ = v181
                                        }
                                    }
                                    v182 := []int{(fi24 << ((1) & 63))}
                                    var v183 int8 = -128
                                    _ = v182
                                    _ = v183
                                }
                                _ = v26
                            }
                            var v184 bool = (1 == 1)
                            _ = fi24
                            _ = v184
                        }
                        p0 = p0
                        v185 := GenS17{F0: -1, F1: 4294967295}
                        _ = v185
                    } else {
                        v186 := []int{42}
                        v187 := GenS18{F0: 745}
                        _ = v186
                        _ = v187
                    }
                }
            }
        }
        var v188 string = (((p0 + (p0 + p0)) + p0) + p0)
        _ = v18
        _ = v188
    } else {
        if true {
            if ((42 >= 1) || true) {
                for fi189 := 0; fi189 < 3; fi189++ {
                    v190 := GenS13{F0: 521, F1: 83, F2: 1}
                    _ = fi189
                    _ = v190
                }
                v191 := []int{128, -437, -769, 42}
                for fi192 := 0; fi192 < 16; fi192++ {
                    if len(v191) > 0 {
                        v191[0] = -9223372036854775808
                    }
                    _ = fi192
                }
                _ = v191
            }
            p0 = ("" + p0)
        } else {
            var v193 uint16 = 774
            var v194 uint16 = (^(v193))
            var v195 uint16 = v194
            _ = v193
            _ = v194
            _ = v195
        }
        fmt.Println(p0)
        for fi196 := 0; fi196 < 14; fi196++ {
            var v197 string = ("foo bar" + (("a" + ("hello" + p0)) + p0))
            p0 = ("hello" + "a")
            _ = fi196
            _ = v197
        }
    }
    var v198 uint8 = 2
    var v199 int8 = 84
    var v200 int16 = -1
    _ = p0
    _ = v198
    _ = v199
    _ = v200
    return 255
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 0
    var v3 bool = true
    var v4 float32 = 3.14159
    v5 := make(map[string]string)
    v6 := make(map[int]uint32)
    v4 += (v4 / v4)
    v201 := v5[("\n\t" + (("hello" + (("foo bar" + "a") + "a")) + "a"))]
    fmt.Println(v3, v3, v4)
    v4 -= v4
    v201 = (("" + v201) + v201)
    var v202 uint16 = 2
    for fi203 := 0; fi203 < 12; fi203++ {
        v204 := GenS10{F0: 32767, F1: v202}
        var v205 uint = genFn1("")
        _ = fi203
        _ = v204
        _ = v205
    }
    fmt.Println(v3)
    switch v201 {
    case "":
        delete(v6, 32768)
    case "foo bar":
        fmt.Println(v201, v2, v201)
    }
    v3 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v201:", v201)
    fmt.Println("v202:", v202)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v201
    _ = v202
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
