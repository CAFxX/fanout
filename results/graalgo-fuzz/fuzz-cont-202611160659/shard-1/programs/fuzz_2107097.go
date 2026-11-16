package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 bool; F2 int32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int64; F1 uint; F2 uint16 }
type GenS9 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS10 struct { F0 uint; F1 bool; F2 int8 }
type GenS11 struct { F0 int64; F1 int32; F2 uint }
type GenS12 struct { F0 int32; F1 int64; F2 uint32 }
type GenS13 struct { F0 int8; F1 int8; F2 int8 }
type GenS14 struct { F0 int8; F1 bool }
type GenS15 struct { F0 bool; F1 int8; F2 int16 }
type GenS16 struct { F0 uint8; F1 string }
type GenS17 struct { F0 uint16; F1 int64 }
type GenS18 struct { F0 string; F1 uint8 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint16 }
type GenS21 struct { F0 uint64; F1 int16 }
type GenS22 struct { F0 int; F1 int8; F2 uint }
type GenS23 struct { F0 int32 }
type GenS24 struct { F0 uint64 }
type GenS25 struct { F0 uint64 }
type GenS26 struct { F0 uint8; F1 bool; F2 string }
type GenS27 struct { F0 uint64; F1 uint16 }
type GenS28 struct { F0 uint32; F1 uint64; F2 bool }
type GenS29 struct { F0 uint16 }
type GenS30 struct { F0 int8 }
type GenS31 struct { F0 int16; F1 int32; F2 bool }
type GenS32 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS33 struct { F0 string; F1 uint16; F2 int }
type GenS34 struct { F0 uint32; F1 int; F2 uint }
type GenS35 struct { F0 uint; F1 uint }
type GenS36 struct { F0 uint32; F1 uint16; F2 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 255
    var v3 uint16 = 32767
    var v4 int = 9223372036854775807
    var v5 int32 = 306
    var v6 int16 = 579
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 float32 = -2.25
        for fi9 := 0; fi9 < 16; fi9++ {
            switch (v4 ^ (-(fi9))) {
            case -9223372036854775808:
                if (320 > 65535) {
                    if (false && true) {
                        v10 := v1[42]
                        if (!false) {
                            var v11 uint = v2
                            v12 := []int{(^(fi9))}
                            _ = len(v1)
                            _ = v11
                            _ = v12
                        } else {
                            switch (("\n\t" + "hello") + "hello") {
                            case "hello":
                                var v13 int64 = 1
                                _ = v13
                            case "":
                                v14 := []int{}
                                _ = v14
                            }
                        }
                        var v15 float64 = 1e10
                        _ = v10
                        _ = v15
                    }
                    v16 := GenS1{F0: true, F1: 128, F2: 1}
                    var v17 bool = true
                    _ = v16
                    _ = v17
                }
            case 2147483647:
                v2 = (v2 % ((v2) | 1))
            case 128:
                var v18 uint32 = 448
                _ = v18
            }
            delete(v1, 2)
            _ = fi9
        }
        if true {
            if (false && (0.1 > (float64((v8 - v8)) / 3.14159))) {
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 string = ("a" + "\n\t")
                    v21 := []int{(fi19 / ((127) | 1))}
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                for fi22 := 0; fi22 < 13; fi22++ {
                    v23 := []int{fi22}
                    var v24 string = ""
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            }
            if false {
                for fi25 := 0; fi25 < 1; fi25++ {
                    if false {
                        var v26 uint16 = (v3 - 65535)
                        v27 := GenS1{F0: true, F1: (^(v4)), F2: 18446744073709551615}
                        v2 = 716
                        _ = v26
                        _ = v27
                    } else {
                        if (383 != v2) {
                            fmt.Println(v4, fi7, v4)
                            v28 := []int{39, fi7, fi7}
                            fmt.Println(v8, v3)
                            switch ("foo bar" + "a") {
                            case "foo bar":
                                if false {
                                    v8 -= ((v8 - v8) + v8)
                                } else {
                                    var v29 uint32 = uint32(v3)
                                    if len(v28) >= 2 {
                                        v28 = v28[1:len(v28)]
                                    }
                                    _ = v29
                                }
                            default:
                                fmt.Println(fi25, v4)
                            }
                            _ = v28
                        }
                    }
                    var v30 int16 = v6
                    switch "a" {
                    case "a":
                        var v31 uint64 = 2
                        _ = v31
                    default:
                        if false {
                            var v32 int16 = 1000
                            fmt.Println(v32)
                            v33 := GenS1{F0: false, F1: (fi25 | v4), F2: 4294967295}
                            _ = v32
                            _ = v33
                        } else {
                            var v34 uint = v2
                            var v35 int = (fi7 + -618)
                            _ = len(v1)
                            _ = v34
                            _ = v35
                        }
                        v4 -= v4
                    }
                    v3 -= (v3 & (v3 & (^(v3))))
                    switch "hello" {
                    case "a":
                        v36 := []int{}
                        _ = v36
                    case "":
                        var v37 int8 = -1
                        _ = v37
                    }
                    v38 := GenS1{F0: (!false), F1: (-(fi7)), F2: 42}
                    _ = fi25
                    _ = v30
                    _ = v38
                }
                for fi39 := 0; fi39 < 5; fi39++ {
                    v2 = 4294967295
                    if ((v3 - 2) == 208) {
                        v40 := GenS2{F0: 2, F1: 346}
                        var v41 string = ("foo bar" + "\n\t")
                        _ = v40
                        _ = v41
                    }
                    v8 -= v8
                    if true {
                        v42 := GenS1{F0: true, F1: 2147483647, F2: 544}
                        switch ("\n\t" + ("foo bar" + ("foo bar" + ""))) {
                        case "a":
                            switch fi39 {
                            case -2147483648:
                                var v43 uint64 = 2
                                _ = v43
                            case 4294967295:
                                switch ("hello" + "foo bar") {
                                case "hello":
                                    v8 = 1e10
                                case "":
                                    _ = len(v1)
                                default:
                                    if true {
                                        fmt.Println(v2)
                                        v44 := GenS3{F0: 65535, F1: 237}
                                        fmt.Println(v6, fi7, v5)
                                        _ = v44
                                    } else {
                                        if (true && (1e10 <= 1.5)) {
                                            v3 += v3
                                            v45 := []int{fi7, 4294967295, fi7}
                                            switch fi39 {
                                            case 9223372036854775807:
                                                v5 = (v5 % (((v5 ^ (-(v5)))) | 1))
                                                if len(v45) > 0 {
                                                    v45[0] = fi7
                                                }
                                            default:
                                                if (v3 <= v3) {
                                                    var v46 float32 = v8
                                                    if false {
                                                        var v47 uint16 = v3
                                                        _ = v47
                                                    }
                                                    _ = v46
                                                } else {
                                                    switch ("" + ("hello" + "foo bar")) {
                                                    case "hello":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if (v8 <= 1.5) {
                                                                v4 -= (fi39 / ((358) | 1))
                                                            } else {
                                                                var v48 int = (fi39 * -2147483648)
                                                                var v49 uint16 = (v3 / ((v3) | 1))
                                                                if (1e10 <= (v8 - -2.25)) {
                                                                    v48 = int(v2)
                                                                }
                                                                _ = len(v1)
                                                                v50 := GenS1{F0: (2147483648 >= 4294967295), F1: fi39, F2: 18446744073709551615}
                                                                v2 -= 28
                                                                var v51 uint8 = 42
                                                                v6 = int16(v51)
                                                                switch ("hello" + "hello") {
                                                                case "\n\t":
                                                                    switch (^(fi39)) {
                                                                    case 1:
                                                                        v52 := GenS3{F0: (v2 ^ v2), F1: (^(v51))}
                                                                        _ = v52
                                                                    default:
                                                                        if (!true) {
                                                                            v49 = (v49 - (v49 | v49))
                                                                            switch fi7 {
                                                                            case 257:
                                                                                v53 := GenS1{F0: false, F1: 32768, F2: 241}
                                                                                _ = v53
                                                                            case 1:
                                                                                switch ("" + "foo bar") {
                                                                                case "hello":
                                                                                    v6 = v6
                                                                                case "a":
                                                                                    fmt.Println(v51, v48, v3)
                                                                                default:
                                                                                    if true {
                                                                                        var v54 float32 = 1e10
                                                                                        v51 = 2
                                                                                        _ = v54
                                                                                    } else {
                                                                                        if len(v45) > 0 {
                                                                                            v45[0] = fi7
                                                                                        }
                                                                                    }
                                                                                }
                                                                            case -9223372036854775808:
                                                                                if (!(true && ((!true) && true))) {
                                                                                    fmt.Println(v2, v51, fi39)
                                                                                    _ = len(v1)
                                                                                } else {
                                                                                    switch v48 {
                                                                                    case 948:
                                                                                        var v55 int64 = 42
                                                                                        _ = v55
                                                                                    default:
                                                                                        if false {
                                                                                            switch ("a" + "\n\t") {
                                                                                            case "a":
                                                                                                v56 := GenS1{F0: ((v5 > 1) && (!(!(v5 < 2147483647)))), F1: v48, F2: 1}
                                                                                                _ = v56
                                                                                            default:
                                                                                                v57 := GenS4{F0: ((v49 >> ((v2) & 63)) >= v3)}
                                                                                                v8 -= (v8 / float32(3.14159))
                                                                                                _ = v57
                                                                                            }
                                                                                            _ = len(v45)
                                                                                        }
                                                                                    }
                                                                                    v5 = (^(v5))
                                                                                    _ = len(v1)
                                                                                }
                                                                            default:
                                                                                switch (fi39 % (((fi39 & (v48 | (v48 % ((fi7) | 1))))) | 1)) {
                                                                                case 0:
                                                                                    _ = len(v45)
                                                                                case -2147483648:
                                                                                    var v58 string = "foo bar"
                                                                                    _ = v58
                                                                                case -23:
                                                                                    if (!true) {
                                                                                        var v59 int = v4
                                                                                        v60 := GenS4{F0: false}
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                    } else {
                                                                                        if true {
                                                                                            fmt.Println(v8, v4, fi7)
                                                                                            if false {
                                                                                                fmt.Println(fi7)
                                                                                                switch (v4 - (fi39 + v4)) {
                                                                                                case -405:
                                                                                                    _ = len(v1)
                                                                                                case -651:
                                                                                                    v48 = v48
                                                                                                default:
                                                                                                    fmt.Println(v51)
                                                                                                }
                                                                                                v6 -= v6
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v51 = v51
                                                                                var v61 uint = (v2 * (v2 * (v2 << ((v2) & 63))))
                                                                                var v62 int16 = v6
                                                                                var v63 uint16 = (^(v49))
                                                                                _ = v61
                                                                                _ = v62
                                                                                _ = v63
                                                                            }
                                                                            if (v5 > (v5 ^ v5)) {
                                                                                var v64 string = ("hello" + ("a" + (("" + "") + (("hello" + "hello") + "foo bar"))))
                                                                                v65 := v1[477]
                                                                                v66 := GenS5{F0: 2, F1: false, F2: v5}
                                                                                _ = v64
                                                                                _ = v65
                                                                                _ = v66
                                                                            } else {
                                                                                fmt.Println(v5)
                                                                            }
                                                                        } else {
                                                                            v67 := GenS1{F0: ((v51 >> ((uint(v3)) & 63)) >= v51), F1: 0, F2: 1}
                                                                            var v68 int16 = (v6 << ((v2) & 63))
                                                                            _ = v67
                                                                            _ = v68
                                                                        }
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            v6 = 1
                                                                        case "hello":
                                                                            if true {
                                                                                fmt.Println(fi39, v51, v5)
                                                                                var v69 int64 = -127
                                                                                var v70 float32 = v8
                                                                                _ = v69
                                                                                _ = v70
                                                                            }
                                                                        default:
                                                                            if (!((!true) || true)) {
                                                                                v71 := GenS6{F0: 4294967295}
                                                                                _ = v71
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    if (v8 < (-2.25 / 1e10)) {
                                                                        var v72 int = v4
                                                                        var v73 string = ""
                                                                        var v74 int8 = int8(v2)
                                                                        _ = v72
                                                                        _ = v73
                                                                        _ = v74
                                                                    }
                                                                }
                                                                if false {
                                                                    v75 := GenS7{F0: 494}
                                                                    _ = v75
                                                                } else {
                                                                    var v76 int32 = (v5 + (v5 & int32(v49)))
                                                                    v4 = (v4 | 32767)
                                                                    _ = v76
                                                                }
                                                                _ = v48
                                                                _ = v49
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        case "foo bar":
                                                            if true {
                                                                if (65535 <= 749) {
                                                                    if false {
                                                                        v77 := v1[809]
                                                                        _ = v77
                                                                    } else {
                                                                        var v78 int64 = 1
                                                                        v6 += (v6 >> (((v2 | v2)) & 63))
                                                                        _ = v78
                                                                    }
                                                                    if (!(!true)) {
                                                                        v45 = append(v45, (v4 >> (((v2 + (^(v2)))) & 63)))
                                                                        if (!true) {
                                                                            v79 := GenS1{F0: (((^(v6)) >= (-(v6))) || ((((3.14159 / -2.25) / -2.25) >= 3.14159) || false)), F1: fi7, F2: 234}
                                                                            fmt.Println(v5)
                                                                            v80 := GenS4{F0: ((true && true) || false)}
                                                                            var v81 uint64 = 18446744073709551615
                                                                            switch 0 {
                                                                            case 365:
                                                                                fmt.Println(v5)
                                                                            case 4294967295:
                                                                                v5 -= v5
                                                                            case 128:
                                                                                fmt.Println(fi39, v8)
                                                                            }
                                                                            _ = v79
                                                                            _ = v80
                                                                            _ = v81
                                                                        } else {
                                                                            if true {
                                                                                if (true && false) {
                                                                                    v45 = append(v45, 1)
                                                                                    _ = len(v1)
                                                                                } else {
                                                                                    var v82 int32 = (v5 + (v5 << (((v2 % (((v2 & v2)) | 1))) & 63)))
                                                                                    _ = v82
                                                                                }
                                                                                if true {
                                                                                    if (v5 > v5) {
                                                                                        switch (fi39 << ((uint(v3)) & 63)) {
                                                                                        case -85:
                                                                                            fmt.Println(v5)
                                                                                            if (float64(((v8 * v8) - (v8 * (v8 + (-2.25 / v8))))) <= 1.5) {
                                                                                                switch ((((("foo bar" + (("a" + "hello") + "hello")) + ("hello" + "foo bar")) + "foo bar") + "foo bar") + "hello") {
                                                                                                case "hello":
                                                                                                    v83 := GenS6{F0: (v2 << (((v2 & v2)) & 63))}
                                                                                                    _ = v83
                                                                                                case "\n\t":
                                                                                                    var v84 uint = (v2 & v2)
                                                                                                    _ = v84
                                                                                                }
                                                                                            }
                                                                                        case 127:
                                                                                            v1[2] = (v4 >> ((v2) & 63))
                                                                                        case 32767:
                                                                                            v85 := GenS3{F0: 42, F1: 42}
                                                                                            _ = v85
                                                                                        }
                                                                                        _ = len(v45)
                                                                                    }
                                                                                    v5 = v5
                                                                                }
                                                                                if true {
                                                                                    switch (v4 % ((fi7) | 1)) {
                                                                                    case 1:
                                                                                        if (true || true) {
                                                                                            if (78 > -22) {
                                                                                                fmt.Println(fi39, fi39, fi7)
                                                                                                if false {
                                                                                                    v45 = append(v45, fi7)
                                                                                                    _ = len(v45)
                                                                                                }
                                                                                                var v86 float32 = float32(3.14159)
                                                                                                _ = v86
                                                                                            }
                                                                                            switch "a" {
                                                                                            case "a":
                                                                                                v4 = fi7
                                                                                            }
                                                                                        } else {
                                                                                            if true {
                                                                                                fmt.Println(fi39, v5)
                                                                                            }
                                                                                            if (-128 == 42) {
                                                                                                var v87 int = v4
                                                                                                if (18446744073709551615 < 2) {
                                                                                                    v8 += 1.5
                                                                                                }
                                                                                                _ = v87
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                var v88 uint16 = v3
                                                                                var v89 bool = true
                                                                                _ = len(v45)
                                                                                _ = v88
                                                                                _ = v89
                                                                            }
                                                                            _ = len(v1)
                                                                            v90 := GenS8{F0: -9223372036854775808, F1: (v2 + (^(v2))), F2: (^(v3))}
                                                                            v91 := GenS9{F0: 55, F1: (v3 - 804), F2: 2}
                                                                            _ = v90
                                                                            _ = v91
                                                                        }
                                                                    } else {
                                                                        var v92 int8 = 42
                                                                        fmt.Println(v8, v4)
                                                                        var v93 float32 = v8
                                                                        _ = v92
                                                                        _ = v93
                                                                    }
                                                                    switch (fi39 >> ((v2) & 63)) {
                                                                    case 263:
                                                                        if false {
                                                                            if (432 >= 1) {
                                                                                switch v4 {
                                                                                case 127:
                                                                                    v3 -= (v3 - 32767)
                                                                                default:
                                                                                    v4 -= fi7
                                                                                }
                                                                                v5 += v5
                                                                                fmt.Println(v2)
                                                                                var v94 uint = v2
                                                                                v6 = v6
                                                                                _ = v94
                                                                            }
                                                                            switch 32767 {
                                                                            case -1:
                                                                                fmt.Println(v6, v4, fi39)
                                                                            case -2147483648:
                                                                                if true {
                                                                                    fmt.Println(fi39, fi7)
                                                                                } else {
                                                                                    v5 += (v5 << ((v2) & 63))
                                                                                    v95 := v1[944]
                                                                                    v96 := GenS5{F0: 792, F1: true, F2: v5}
                                                                                    _ = v95
                                                                                    _ = v96
                                                                                }
                                                                            }
                                                                        }
                                                                    case 2147483647:
                                                                        v97 := GenS10{F0: (v2 - (v2 << (((v2 * (v2 * (v2 / ((v2) | 1))))) & 63))), F1: false, F2: 0}
                                                                        _ = v97
                                                                    case -937:
                                                                        v98 := v1[2147483648]
                                                                        _ = v98
                                                                    default:
                                                                        v99 := GenS8{F0: 448, F1: (v2 << ((822) & 63)), F2: v3}
                                                                        _ = v99
                                                                    }
                                                                } else {
                                                                    if (("\n\t" + "foo bar") < (("a" + ("hello" + "")) + (("\n\t" + ("\n\t" + "")) + ("foo bar" + "a")))) {
                                                                        v100 := GenS4{F0: ((!((float32(3.14159) * (1e10 * v8)) < float32(1e10))) || (v6 != (v6 ^ (v6 << (((v2 ^ (v2 | v2))) & 63)))))}
                                                                        _ = v100
                                                                    }
                                                                }
                                                            }
                                                        case "a":
                                                            v101 := GenS11{F0: int64(v3), F1: (v5 | (v5 % (((v5 + 0)) | 1))), F2: 42}
                                                            _ = v101
                                                        default:
                                                            v4 = -9223372036854775808
                                                        }
                                                    }
                                                    fmt.Println(v8, v2, v5)
                                                }
                                            }
                                            _ = v45
                                        }
                                        if (42 >= 883) {
                                            v8 += v8
                                            v5 = v5
                                            v1[4294967295] = (^(fi39))
                                            v102 := GenS12{F0: (-(v5)), F1: 42, F2: 2147483648}
                                            v103 := GenS9{F0: 920, F1: (v3 ^ 1), F2: 2147483647}
                                            switch fi7 {
                                            case 4294967295:
                                                if (!false) {
                                                    if false {
                                                        if (0 > 747) {
                                                            if false {
                                                                switch ("\n\t" + (("hello" + "\n\t") + (("\n\t" + ("foo bar" + ("a" + (("\n\t" + "foo bar") + "")))) + ("" + "foo bar")))) {
                                                                case "\n\t":
                                                                    if false {
                                                                        v5 = v5
                                                                        fmt.Println(fi39, v4, v6)
                                                                    }
                                                                    var v104 int32 = v5
                                                                    _ = v104
                                                                case "":
                                                                    var v105 string = ("" + "a")
                                                                    _ = v105
                                                                case "hello":
                                                                    var v106 int32 = v5
                                                                    _ = v106
                                                                default:
                                                                    var v107 int = fi7
                                                                    var v108 int8 = 0
                                                                    var v109 bool = true
                                                                    _ = v107
                                                                    _ = v108
                                                                    _ = v109
                                                                }
                                                                v4 = 2147483647
                                                                v110 := GenS12{F0: v5, F1: 180, F2: 42}
                                                                _ = v110
                                                            }
                                                            v4 = 32768
                                                            if (!false) {
                                                                var v111 uint8 = 127
                                                                fmt.Println(fi7)
                                                                v6 = v6
                                                                _ = v111
                                                            } else {
                                                                v4 = 4294967295
                                                                v112 := GenS13{F0: 42, F1: -95, F2: 100}
                                                                _ = v112
                                                            }
                                                            switch ("a" + "foo bar") {
                                                            case "":
                                                                v113 := []int{(fi7 + (-(fi7))), fi39, (fi39 * fi39)}
                                                                _ = v113
                                                            case "a":
                                                                if (0 >= v3) {
                                                                    fmt.Println(v5)
                                                                }
                                                            default:
                                                                var v114 int32 = v5
                                                                _ = v114
                                                            }
                                                        }
                                                        v115 := v1[50]
                                                        v5 = v5
                                                        _ = v115
                                                    } else {
                                                        fmt.Println(v5)
                                                        v5 = (v5 | v5)
                                                        v8 = 0.1
                                                        switch (fi7 ^ 811) {
                                                        case 1:
                                                            var v116 uint8 = 127
                                                            var v117 uint16 = v3
                                                            switch (fi39 ^ (fi7 & 128)) {
                                                            case 127:
                                                                v118 := GenS13{F0: 84, F1: -1, F2: -1}
                                                                _ = v118
                                                            case -448:
                                                                v119 := GenS4{F0: false}
                                                                _ = v119
                                                            case -1:
                                                                if (!true) {
                                                                    v116 -= v116
                                                                    switch 128 {
                                                                    case 935:
                                                                        if ((!true) || true) {
                                                                            v120 := GenS14{F0: 0, F1: (!(0 != 65535))}
                                                                            switch (fi39 - fi7) {
                                                                            case 510:
                                                                                v121 := GenS4{F0: (0 <= -1)}
                                                                                _ = v121
                                                                            case 42:
                                                                                switch "a" {
                                                                                case "a":
                                                                                    v122 := []int{(fi7 << ((4294967295) & 63)), v4}
                                                                                    _ = v122
                                                                                case "hello":
                                                                                    _ = len(v1)
                                                                                default:
                                                                                    v123 := GenS15{F0: true, F1: -16, F2: 930}
                                                                                    _ = v123
                                                                                }
                                                                            case 753:
                                                                                var v124 uint64 = uint64(v2)
                                                                                _ = v124
                                                                            default:
                                                                                v117 += 255
                                                                            }
                                                                            v4 -= v4
                                                                            _ = v120
                                                                        }
                                                                    case -2147483648:
                                                                        fmt.Println(fi7)
                                                                    case 0:
                                                                        var v125 string = "hello"
                                                                        _ = v125
                                                                    }
                                                                } else {
                                                                    var v126 int16 = v6
                                                                    _ = len(v1)
                                                                    v2 = 4294967295
                                                                    _ = v126
                                                                }
                                                            }
                                                            _ = v116
                                                            _ = v117
                                                        }
                                                    }
                                                    switch ("hello" + ("a" + "\n\t")) {
                                                    case "\n\t":
                                                        fmt.Println(v2, v3)
                                                    case "a":
                                                        fmt.Println(v6)
                                                    }
                                                    if ((-436 >= 1) || true) {
                                                        var v127 uint16 = (v3 + (^(v3)))
                                                        var v128 int32 = v5
                                                        var v129 bool = (0 != 0)
                                                        _ = v127
                                                        _ = v128
                                                        _ = v129
                                                    } else {
                                                        v3 = (v3 >> ((v2) & 63))
                                                        v3 = 42
                                                        v8 = (0.0 * v8)
                                                    }
                                                }
                                            default:
                                                var v130 int8 = 100
                                                _ = v130
                                            }
                                            var v131 bool = (false && true)
                                            var v132 float32 = 0.1
                                            v5 -= (v5 << ((976) & 63))
                                            v133 := GenS8{F0: 0, F1: (v2 + (v2 % ((329) | 1))), F2: v3}
                                            if v131 {
                                                v2 += v2
                                                switch "\n\t" {
                                                case "":
                                                    v134 := GenS16{F0: 128, F1: ("a" + "a")}
                                                    _ = v134
                                                case "hello":
                                                    if ((v8 - v8) >= 3.14159) {
                                                        var v135 int32 = v5
                                                        var v136 uint64 = 1
                                                        _ = v135
                                                        _ = v136
                                                    } else {
                                                        v137 := GenS17{F0: 2, F1: 538}
                                                        v138 := GenS13{F0: 80, F1: 100, F2: 42}
                                                        if true {
                                                            if v131 {
                                                                v132 = ((((v8 + (1.5 - ((float32((0.0 * -2.25)) / 1.0) - float32(0.1)))) - v132) / 1.0) - v132)
                                                            }
                                                            v139 := GenS9{F0: 0, F1: (v3 / ((v3) | 1)), F2: 2147483648}
                                                            var v140 uint16 = (v3 + v3)
                                                            if (("\n\t" >= "") || ((!v131) && (18446744073709551615 >= (^(v2))))) {
                                                                var v141 uint32 = 2147483647
                                                                if v131 {
                                                                    v142 := GenS18{F0: "foo bar", F1: 42}
                                                                    switch -2147483648 {
                                                                    case 32768:
                                                                        if (!((-1 <= -128) || (!(v3 < (v140 / ((v140) | 1)))))) {
                                                                            if false {
                                                                                if (false || ((true && (false && (887 >= 42))) && (v6 > v6))) {
                                                                                    if ((v141 - 542) < 845) {
                                                                                        if (true || false) {
                                                                                            var v143 int = fi39
                                                                                            v132 += v132
                                                                                            fmt.Println(v8)
                                                                                            _ = v143
                                                                                        } else {
                                                                                            v144 := GenS8{F0: 9223372036854775807, F1: v2, F2: v140}
                                                                                            _ = v144
                                                                                        }
                                                                                    }
                                                                                    if (!v131) {
                                                                                        v8 = float32(3.14159)
                                                                                        var v145 bool = v131
                                                                                        var v146 uint = (^(v2))
                                                                                        _ = v145
                                                                                        _ = v146
                                                                                    } else {
                                                                                        fmt.Println(v4, v2, v4)
                                                                                        v147 := []int{(v4 + (v4 * fi39))}
                                                                                        if true {
                                                                                            fmt.Println(v8)
                                                                                            if true {
                                                                                                v1[1] = fi39
                                                                                                var v148 int8 = -75
                                                                                                _ = v148
                                                                                            }
                                                                                            v140 = v140
                                                                                        } else {
                                                                                            v6 += v6
                                                                                            delete(v1, (^(v141)))
                                                                                            var v149 bool = ((v2 & (v2 * v2)) <= (v2 | v2))
                                                                                            _ = v149
                                                                                        }
                                                                                        var v150 int64 = 9223372036854775807
                                                                                        v1[v141] = (^(v4))
                                                                                        _ = v147
                                                                                        _ = v150
                                                                                    }
                                                                                    v141 += (^(v141))
                                                                                }
                                                                            } else {
                                                                                if (!(42 < 0)) {
                                                                                    if true {
                                                                                        var v151 int8 = 100
                                                                                        var v152 int16 = v6
                                                                                        v1[v141] = fi7
                                                                                        _ = v151
                                                                                        _ = v152
                                                                                    }
                                                                                    v153 := GenS13{F0: -128, F1: -128, F2: 78}
                                                                                    switch "" {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v4, v3)
                                                                                    case "":
                                                                                        fmt.Println(v131)
                                                                                    case "\n\t":
                                                                                        v154 := []int{v4, (-(v4)), -9223372036854775808, (v4 * v4)}
                                                                                        _ = v154
                                                                                    default:
                                                                                        v155 := GenS10{F0: 175, F1: (v3 > v140), F2: 91}
                                                                                        _ = v155
                                                                                    }
                                                                                    _ = v153
                                                                                } else {
                                                                                    var v156 int16 = (v6 / (((^(v6))) | 1))
                                                                                    var v157 float32 = (((((v132 + v132) + v132) / 0.1) - (float32((-2.25 + float64(1e10))) + (v132 / v8))) / (float32(0.0) * ((float32(0.1) * v132) / v8)))
                                                                                    v158 := []int{128}
                                                                                    _ = v156
                                                                                    _ = v157
                                                                                    _ = v158
                                                                                }
                                                                                v4 -= fi7
                                                                            }
                                                                        }
                                                                    case -572:
                                                                        var v159 int64 = 1
                                                                        switch (-(fi7)) {
                                                                        case -126:
                                                                            if (v6 > (v6 * (^(v6)))) {
                                                                                fmt.Println(v141)
                                                                            } else {
                                                                                v160 := GenS6{F0: v2}
                                                                                v131 = v131
                                                                                if (!v131) {
                                                                                    switch (fi39 >> (((v2 * (v2 * v2))) & 63)) {
                                                                                    case -1:
                                                                                        fmt.Println(v4)
                                                                                    default:
                                                                                        v161 := GenS19{F0: -10}
                                                                                        _ = v161
                                                                                    }
                                                                                    v162 := GenS20{F0: v3}
                                                                                    if v131 {
                                                                                        switch (^(fi39)) {
                                                                                        case 9223372036854775807:
                                                                                            switch ("" + ("\n\t" + "a")) {
                                                                                            case "\n\t":
                                                                                                if (v6 != 42) {
                                                                                                    if (v132 > -2.25) {
                                                                                                        var v163 int = 9223372036854775807
                                                                                                        if ((v4 - v4) > (fi7 * v163)) {
                                                                                                            if ((v5 ^ 2147483647) != (v5 >> (((v2 | (v2 / (((v2 * (v2 >> ((v2) & 63)))) | 1)))) & 63))) {
                                                                                                                var v164 uint32 = (v141 * v141)
                                                                                                                switch "" {
                                                                                                                case "":
                                                                                                                    v2 -= v2
                                                                                                                case "a":
                                                                                                                    fmt.Println(v2)
                                                                                                                case "\n\t":
                                                                                                                    switch "" {
                                                                                                                    case "foo bar":
                                                                                                                        fmt.Println(v3, v8, v132)
                                                                                                                    default:
                                                                                                                        _ = len(v1)
                                                                                                                    }
                                                                                                                default:
                                                                                                                    if v131 {
                                                                                                                        v141 -= v141
                                                                                                                        delete(v1, v141)
                                                                                                                        v8 += (v8 - float32(0.0))
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v164
                                                                                                            } else {
                                                                                                                if v131 {
                                                                                                                    fmt.Println(v6, v3)
                                                                                                                    v165 := GenS13{F0: -47, F1: 0, F2: 127}
                                                                                                                    if ((("\n\t" + "") >= ("\n\t" + "foo bar")) && v131) {
                                                                                                                        v166 := GenS21{F0: 42, F1: (v6 - v6)}
                                                                                                                        v159 = int64(v6)
                                                                                                                        switch ("a" + ("\n\t" + "\n\t")) {
                                                                                                                        case "foo bar":
                                                                                                                            v167 := GenS22{F0: fi7, F1: -26, F2: v2}
                                                                                                                            _ = v167
                                                                                                                        case "hello":
                                                                                                                            var v168 int16 = v6
                                                                                                                            _ = v168
                                                                                                                        case "a":
                                                                                                                            _ = len(v1)
                                                                                                                        }
                                                                                                                        _ = v166
                                                                                                                    }
                                                                                                                    _ = v165
                                                                                                                } else {
                                                                                                                    switch -9223372036854775808 {
                                                                                                                    case 32768:
                                                                                                                        fmt.Println(v4)
                                                                                                                    case 323:
                                                                                                                        v141 = 4294967295
                                                                                                                        var v169 uint = (v2 * (v2 >> (((v2 & 18446744073709551615)) & 63)))
                                                                                                                        _ = v169
                                                                                                                    case 2147483647:
                                                                                                                        if (v131 && (((v141 ^ v141) == v141) || false)) {
                                                                                                                            v2 = v2
                                                                                                                            v140 = (v140 + (v3 % (((v140 + v3)) | 1)))
                                                                                                                        } else {
                                                                                                                            v140 -= v140
                                                                                                                            v170 := GenS10{F0: (v2 / ((4294967295) | 1)), F1: v131, F2: 100}
                                                                                                                            _ = v170
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v3 = (v140 << (((^(v2))) & 63))
                                                                                                                }
                                                                                                                if (false && (v4 < (v163 >> ((v2) & 63)))) {
                                                                                                                    if (v2 > (v2 * (v2 | v2))) {
                                                                                                                        if v131 {
                                                                                                                            fmt.Println(fi7, v163)
                                                                                                                            var v171 uint = 65535
                                                                                                                            if ((("foo bar" + "foo bar") + "a") < "hello") {
                                                                                                                                v172 := []int{0, fi39, fi39, v163}
                                                                                                                                _ = v172
                                                                                                                            }
                                                                                                                            v173 := []int{128, (v4 & (^(v163)))}
                                                                                                                            _ = v171
                                                                                                                            _ = v173
                                                                                                                        } else {
                                                                                                                            v3 = (v3 << (((^(v2))) & 63))
                                                                                                                            v174 := GenS23{F0: (v5 & (^(v5)))}
                                                                                                                            _ = v174
                                                                                                                        }
                                                                                                                        var v175 float64 = -2.25
                                                                                                                        _ = v175
                                                                                                                    } else {
                                                                                                                        if ((v140 | v140) < v3) {
                                                                                                                            fmt.Println(v6, fi7)
                                                                                                                            v1[v141] = v163
                                                                                                                            v4 = v163
                                                                                                                        } else {
                                                                                                                            var v176 int16 = v6
                                                                                                                            _ = v176
                                                                                                                        }
                                                                                                                        fmt.Println(v2, v132)
                                                                                                                        v132 = 3.14159
                                                                                                                    }
                                                                                                                    fmt.Println(v132, v132, v132)
                                                                                                                } else {
                                                                                                                    fmt.Println(v2, v140)
                                                                                                                    var v177 int32 = v5
                                                                                                                    _ = v177
                                                                                                                }
                                                                                                                v178 := GenS24{F0: 895}
                                                                                                                _ = v178
                                                                                                            }
                                                                                                            v141 += (v141 | (^(v141)))
                                                                                                            v179 := GenS25{F0: 0}
                                                                                                            var v180 float64 = float64(v8)
                                                                                                            _ = v179
                                                                                                            _ = v180
                                                                                                        } else {
                                                                                                            v1[(^(v141))] = (fi39 ^ v4)
                                                                                                            var v181 uint16 = (v3 + (v3 >> ((v2) & 63)))
                                                                                                            _ = v181
                                                                                                        }
                                                                                                        _ = v163
                                                                                                    }
                                                                                                    if v131 {
                                                                                                        v182 := GenS4{F0: ((^(v3)) == (v3 / (((v140 << (((v2 - (v2 & (v2 * (v2 ^ v2))))) & 63))) | 1)))}
                                                                                                        v183 := GenS16{F0: uint8(v141), F1: (("hello" + ("a" + "a")) + "\n\t")}
                                                                                                        _ = v182
                                                                                                        _ = v183
                                                                                                    } else {
                                                                                                        v184 := v1[v141]
                                                                                                        _ = v184
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        switch v4 {
                                                                                        case 32768:
                                                                                            fmt.Println(v159, v3, v141)
                                                                                        default:
                                                                                            var v185 int64 = -335
                                                                                            _ = v185
                                                                                        }
                                                                                    } else {
                                                                                        v4 = -503
                                                                                        if v131 {
                                                                                            var v186 float32 = v8
                                                                                            if v131 {
                                                                                                switch (^(fi7)) {
                                                                                                case 0:
                                                                                                    v5 = (v5 % ((int32(v159)) | 1))
                                                                                                default:
                                                                                                    v187 := GenS26{F0: 255, F1: (42 >= -21), F2: ("" + "")}
                                                                                                    _ = v187
                                                                                                }
                                                                                            } else {
                                                                                                v131 = v131
                                                                                                var v188 string = ("a" + "\n\t")
                                                                                                fmt.Println(v5)
                                                                                                _ = v188
                                                                                            }
                                                                                            _ = v186
                                                                                        }
                                                                                    }
                                                                                    var v189 float32 = v132
                                                                                    v190 := GenS27{F0: 810, F1: v140}
                                                                                    _ = v162
                                                                                    _ = v189
                                                                                    _ = v190
                                                                                }
                                                                                _ = v160
                                                                            }
                                                                        case -689:
                                                                            v2 += v2
                                                                        case 303:
                                                                            if ((true && (!false)) || true) {
                                                                                delete(v1, (v141 % ((v141) | 1)))
                                                                                var v191 string = ("\n\t" + "a")
                                                                                if true {
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        if v131 {
                                                                                            v192 := GenS12{F0: v5, F1: v159, F2: v141}
                                                                                            _ = v192
                                                                                        } else {
                                                                                            if true {
                                                                                                var v193 float64 = 0.1
                                                                                                _ = v193
                                                                                            }
                                                                                            var v194 bool = (932 < (^(v3)))
                                                                                            _ = v194
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v191 = ""
                                                                                    }
                                                                                    v132 += v8
                                                                                } else {
                                                                                    switch (v191 + (v191 + (v191 + v191))) {
                                                                                    case "hello":
                                                                                        v2 = (v2 / ((65535) | 1))
                                                                                    default:
                                                                                        v131 = (!false)
                                                                                    }
                                                                                    if v131 {
                                                                                        _ = len(v1)
                                                                                        var v195 int8 = 42
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            v159 = v159
                                                                                        case "a":
                                                                                            var v196 float64 = (-2.25 - (1e10 + 1.5))
                                                                                            _ = v196
                                                                                        }
                                                                                        var v197 int32 = -1
                                                                                        v198 := GenS10{F0: (v2 / ((0) | 1)), F1: false, F2: (-(v195))}
                                                                                        _ = v195
                                                                                        _ = v197
                                                                                        _ = v198
                                                                                    }
                                                                                    switch v191 {
                                                                                    case "hello":
                                                                                        v159 += 9223372036854775807
                                                                                        v191 = v191
                                                                                        v199 := []int{fi7, (fi7 >> ((v2) & 63)), -2147483648, int(v2)}
                                                                                        _ = v199
                                                                                    case "a":
                                                                                        v200 := GenS6{F0: (^(v2))}
                                                                                        _ = v200
                                                                                    case "foo bar":
                                                                                        v201 := GenS28{F0: v141, F1: 378, F2: v131}
                                                                                        _ = v201
                                                                                    }
                                                                                }
                                                                                v202 := GenS29{F0: uint16(v2)}
                                                                                var v203 uint8 = 127
                                                                                _ = v191
                                                                                _ = v202
                                                                                _ = v203
                                                                            }
                                                                        }
                                                                        _ = v159
                                                                    default:
                                                                        v131 = false
                                                                    }
                                                                    _ = v142
                                                                } else {
                                                                    if true {
                                                                        var v204 int64 = 9223372036854775807
                                                                        if v131 {
                                                                            if (42 == (-(v6))) {
                                                                                if v131 {
                                                                                    v205 := GenS30{F0: 109}
                                                                                    var v206 float32 = v132
                                                                                    _ = v205
                                                                                    _ = v206
                                                                                } else {
                                                                                    fmt.Println(v204, v132)
                                                                                }
                                                                                switch (fi39 - v4) {
                                                                                case 127:
                                                                                    var v207 uint8 = 42
                                                                                    _ = v207
                                                                                default:
                                                                                    v4 -= (fi7 & (fi39 | (-(v4))))
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v141 -= 1
                                                                        }
                                                                        _ = v204
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            if (false && (true || ((!v131) && (int16(v141) <= (v6 | (v6 % (((v6 << ((0) & 63))) | 1))))))) {
                                                                                var v208 uint16 = (^(v140))
                                                                                if (!(("foo bar" + (("\n\t" + (("a" + "\n\t") + "hello")) + "hello")) < "foo bar")) {
                                                                                    var v209 int8 = 100
                                                                                    var v210 string = "hello"
                                                                                    switch int(v5) {
                                                                                    case 42:
                                                                                        v208 = (v140 ^ v208)
                                                                                    }
                                                                                    _ = v209
                                                                                    _ = v210
                                                                                } else {
                                                                                    v2 += (v2 << (((v2 >> ((159) & 63))) & 63))
                                                                                    delete(v1, 783)
                                                                                }
                                                                                var v211 uint = (v2 * (v2 & (v2 - (^(v2)))))
                                                                                _ = v208
                                                                                _ = v211
                                                                            }
                                                                        }
                                                                        switch (fi7 | (v4 - v4)) {
                                                                        case 128:
                                                                            v212 := []int{}
                                                                            _ = v212
                                                                        case 72:
                                                                            if true {
                                                                                if v131 {
                                                                                    v213 := GenS7{F0: v141}
                                                                                    var v214 int64 = 916
                                                                                    v215 := []int{2147483647, fi7}
                                                                                    fmt.Println(fi39, v2)
                                                                                    var v216 uint16 = v140
                                                                                    if (v140 >= (v140 << (((v2 | v2)) & 63))) {
                                                                                        switch "" {
                                                                                        case "":
                                                                                            v2 = v2
                                                                                        case "foo bar":
                                                                                            v217 := GenS20{F0: v216}
                                                                                            _ = v217
                                                                                        case "a":
                                                                                            v214 = (v214 | v214)
                                                                                        }
                                                                                        var v218 int8 = -12
                                                                                        var v219 int8 = (^(v218))
                                                                                        var v220 int64 = v214
                                                                                        _ = v218
                                                                                        _ = v219
                                                                                        _ = v220
                                                                                    }
                                                                                    var v221 int16 = 42
                                                                                    _ = v213
                                                                                    _ = v214
                                                                                    _ = v215
                                                                                    _ = v216
                                                                                    _ = v221
                                                                                } else {
                                                                                    if v131 {
                                                                                        var v222 uint64 = 506
                                                                                        v141 -= (^(v141))
                                                                                        var v223 int64 = 9223372036854775807
                                                                                        _ = v222
                                                                                        _ = v223
                                                                                    } else {
                                                                                        v6 = 0
                                                                                        v3 -= v3
                                                                                    }
                                                                                }
                                                                                switch 32768 {
                                                                                case 275:
                                                                                    switch ("foo bar" + "foo bar") {
                                                                                    case "":
                                                                                        v3 = v3
                                                                                        switch "" {
                                                                                        case "":
                                                                                            fmt.Println(v141, v2)
                                                                                        case "foo bar":
                                                                                            v224 := GenS26{F0: 42, F1: (127 > 1), F2: ""}
                                                                                            _ = v224
                                                                                        case "hello":
                                                                                            v225 := GenS24{F0: 2}
                                                                                            _ = v225
                                                                                        default:
                                                                                            delete(v1, 4294967295)
                                                                                            if ((^(v6)) > (v6 ^ (v6 * v6))) {
                                                                                                v132 += v132
                                                                                                var v226 float32 = float32(0.1)
                                                                                                _ = v226
                                                                                            }
                                                                                            var v227 uint16 = (v140 / (((v140 / ((v140) | 1))) | 1))
                                                                                            _ = v227
                                                                                        }
                                                                                    case "hello":
                                                                                        if ((1.5 + v8) >= 1.5) {
                                                                                            v228 := []int{fi39, (fi39 / (((fi39 ^ fi7)) | 1)), v4}
                                                                                            var v229 bool = true
                                                                                            var v230 float32 = v8
                                                                                            _ = v228
                                                                                            _ = v229
                                                                                            _ = v230
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v132 = v132
                                                                                    }
                                                                                default:
                                                                                    if (18446744073709551615 > (v2 + 107)) {
                                                                                        if v131 {
                                                                                            var v231 bool = false
                                                                                            _ = v231
                                                                                        } else {
                                                                                            v232 := GenS16{F0: 42, F1: ("\n\t" + "foo bar")}
                                                                                            var v233 uint32 = v141
                                                                                            _ = v232
                                                                                            _ = v233
                                                                                        }
                                                                                        var v234 int64 = -9223372036854775808
                                                                                        _ = v234
                                                                                    } else {
                                                                                        v235 := GenS15{F0: ("foo bar" != "\n\t"), F1: 0, F2: (^(v6))}
                                                                                        _ = v235
                                                                                    }
                                                                                }
                                                                            }
                                                                        case -9223372036854775808:
                                                                            if (v131 || (!true)) {
                                                                                v236 := []int{127}
                                                                                if v131 {
                                                                                    var v237 uint = (v2 % ((v2) | 1))
                                                                                    v238 := GenS31{F0: (v6 << ((2) & 63)), F1: 2147483647, F2: false}
                                                                                    _ = v237
                                                                                    _ = v238
                                                                                } else {
                                                                                    v131 = (-853 >= (v6 - int16(v2)))
                                                                                }
                                                                                if true {
                                                                                    v239 := GenS32{F0: v141, F1: 0, F2: (v3 / ((527) | 1))}
                                                                                    _ = v239
                                                                                }
                                                                                _ = v236
                                                                            } else {
                                                                                v240 := []int{}
                                                                                var v241 uint8 = 127
                                                                                _ = v240
                                                                                _ = v241
                                                                            }
                                                                            v5 += int32(v2)
                                                                            if (!(uint64(v140) < 2)) {
                                                                                var v242 uint32 = v141
                                                                                var v243 int32 = v5
                                                                                _ = v242
                                                                                _ = v243
                                                                            } else {
                                                                                v3 += v140
                                                                                fmt.Println(v5, v5)
                                                                                v244 := GenS33{F0: "", F1: v3, F2: fi7}
                                                                                _ = v244
                                                                            }
                                                                        default:
                                                                            var v245 int = v4
                                                                            _ = v245
                                                                        }
                                                                    }
                                                                }
                                                                _ = v141
                                                            } else {
                                                                var v246 uint8 = 42
                                                                fmt.Println(v8, v132, fi39)
                                                                _ = v246
                                                            }
                                                            switch (^(v4)) {
                                                            case -2147483648:
                                                                var v247 int = int(v2)
                                                                _ = v247
                                                            }
                                                            v132 = 1.5
                                                            _ = v139
                                                            _ = v140
                                                        } else {
                                                            v8 += ((float32(3.14159) / v132) - ((float32(1e10) + 0.1) - v8))
                                                            v132 -= (1.5 / (v8 - (((v8 + v8) * (v8 + 1.5)) - float32((0.1 * (float64(v8) * float64((v132 - v8))))))))
                                                            v248 := GenS16{F0: 127, F1: "a"}
                                                            _ = v248
                                                        }
                                                        _ = v137
                                                        _ = v138
                                                    }
                                                case "a":
                                                    _ = len(v1)
                                                default:
                                                    if (!false) {
                                                        if v131 {
                                                            fmt.Println(v2)
                                                        } else {
                                                            var v249 float32 = (((3.14159 - 0.0) * v8) - float32(0.0))
                                                            var v250 string = "a"
                                                            switch fi39 {
                                                            case 429:
                                                                var v251 float32 = 1.5
                                                                _ = v251
                                                            }
                                                            _ = v249
                                                            _ = v250
                                                        }
                                                        switch fi39 {
                                                        case 2147483647:
                                                            fmt.Println(v6)
                                                        default:
                                                            v131 = true
                                                        }
                                                    }
                                                    if ((v5 + 341) != 1) {
                                                        v252 := []int{-1, (v4 << (((^(v2))) & 63)), -1, (fi39 + 42)}
                                                        v8 -= (v8 * float32((1.5 + (1e10 * ((0.1 + 1e10) - ((3.14159 + ((-2.25 * 1.5) * (float64(v132) / float64((v8 * v8))))) + 3.14159))))))
                                                        var v253 int8 = 100
                                                        _ = v252
                                                        _ = v253
                                                    }
                                                    var v254 uint = 395
                                                    _ = v254
                                                }
                                                var v255 float32 = 0.1
                                                _ = v255
                                            } else {
                                                var v256 float64 = -2.25
                                                if (v5 != (v5 % ((114) | 1))) {
                                                    v257 := v1[42]
                                                    switch (v4 << (((v2 | (v2 >> (((v2 * (^(v2)))) & 63)))) & 63)) {
                                                    case -965:
                                                        fmt.Println(v4, v257)
                                                    default:
                                                        var v258 uint8 = 255
                                                        _ = v258
                                                    }
                                                    var v259 int64 = 42
                                                    v260 := GenS14{F0: -1, F1: v131}
                                                    _ = v257
                                                    _ = v259
                                                    _ = v260
                                                } else {
                                                    if ((v6 << (((v2 % (((v2 | (v2 & v2))) | 1))) & 63)) > v6) {
                                                        if (!v131) {
                                                            var v261 uint8 = 127
                                                            v262 := []int{v4}
                                                            _ = v261
                                                            _ = v262
                                                        }
                                                        if v131 {
                                                            var v263 int64 = 9223372036854775807
                                                            _ = v263
                                                        }
                                                        var v264 int32 = (v5 & v5)
                                                        _ = v264
                                                    }
                                                    v265 := GenS17{F0: v3, F1: 1}
                                                    var v266 bool = v131
                                                    if ((false || v266) && (32767 >= (v3 * v3))) {
                                                        if (!(255 != (^(v3)))) {
                                                            fmt.Println(v5)
                                                            fmt.Println(fi39, v256, fi39)
                                                        } else {
                                                            if (v256 == v256) {
                                                                var v267 uint32 = 283
                                                                _ = v267
                                                            } else {
                                                                if ((!v266) || v266) {
                                                                    v2 += (^(v2))
                                                                    v268 := v1[343]
                                                                    v269 := []int{}
                                                                    _ = v268
                                                                    _ = v269
                                                                }
                                                                var v270 int32 = -2147483648
                                                                v271 := v1[0]
                                                                _ = v270
                                                                _ = v271
                                                            }
                                                        }
                                                    }
                                                    var v272 uint16 = v3
                                                    v131 = ((v6 - v6) != (v6 * (v6 | (^(v6)))))
                                                    if (42 > 2147483648) {
                                                        v273 := GenS7{F0: 42}
                                                        _ = v273
                                                    }
                                                    fmt.Println(v266, v132)
                                                    _ = v265
                                                    _ = v266
                                                    _ = v272
                                                }
                                                v6 = (v6 ^ (v6 & 1000))
                                                switch ("\n\t" + "\n\t") {
                                                case "":
                                                    var v274 uint = v2
                                                    _ = v274
                                                case "a":
                                                    var v275 uint = 42
                                                    _ = v275
                                                default:
                                                    if true {
                                                        v276 := []int{v4}
                                                        if v131 {
                                                            var v277 int = (fi39 / (((fi7 % (((fi7 | 128)) | 1))) | 1))
                                                            var v278 uint = (^(v2))
                                                            fmt.Println(v256, v256)
                                                            var v279 float64 = (float64(v8) + float64((v8 - float32(((v256 + 3.14159) + float64(v132))))))
                                                            _ = len(v1)
                                                            _ = v277
                                                            _ = v278
                                                            _ = v279
                                                        }
                                                        _ = v276
                                                    }
                                                }
                                                _ = v256
                                            }
                                            v2 = v2
                                            _ = v102
                                            _ = v103
                                            _ = v131
                                            _ = v132
                                            _ = v133
                                        } else {
                                            v280 := []int{int(v5), -2147483648}
                                            _ = v280
                                        }
                                        var v281 string = ("\n\t" + "")
                                        _ = v281
                                    }
                                }
                            case -1:
                                v282 := GenS34{F0: 217, F1: (v4 | fi39), F2: 318}
                                _ = v282
                            default:
                                v283 := GenS32{F0: 42, F1: 1, F2: v3}
                                v284 := GenS35{F0: (v2 & v2), F1: (v2 >> ((v2) & 63))}
                                v3 -= v3
                                _ = v283
                                _ = v284
                            }
                        case "":
                            if (32768 == (-(fi7))) {
                                v2 = (v2 / ((v2) | 1))
                                var v285 string = ""
                                v286 := []int{(fi7 / ((-9223372036854775808) | 1)), int(v3), (fi39 - fi39)}
                                var v287 float64 = (0.0 * 3.14159)
                                _ = v285
                                _ = v286
                                _ = v287
                            } else {
                                var v288 uint64 = 42
                                v289 := []int{}
                                _ = v288
                                _ = v289
                            }
                        }
                        v4 = fi7
                        _ = v42
                    }
                    _ = fi39
                }
                for fi290 := 0; fi290 < 12; fi290++ {
                    delete(v1, 2)
                    var v291 string = ""
                    var v292 uint64 = 0
                    _ = fi290
                    _ = v291
                    _ = v292
                }
            } else {
                v4 += -605
            }
        } else {
            v293 := GenS36{F0: 2, F1: 65535, F2: 128}
            fmt.Println(v5, v5)
            _ = v293
        }
        _ = fi7
        _ = v8
    }
    v6 = int16(v3)
    if ((!true) || (!false)) {
        v294 := GenS15{F0: true, F1: -86, F2: (v6 * (v6 + (v6 - 968)))}
        if (255 >= 127) {
            fmt.Println(v6)
            v295 := []int{v4, 941, v4}
            _ = v295
        }
        _ = v294
    }
    v2 = (v2 % ((1) | 1))
    var v296 uint32 = 42
    if ((v296 - v296) < (v296 << ((444) & 63))) {
        var v297 int64 = 0
        _ = v297
    } else {
        for fi298 := 0; fi298 < 14; fi298++ {
            switch v4 {
            case -1:
                var v299 float32 = 0.1
                _ = v299
            }
            _ = fi298
        }
        var v300 float32 = 0.1
        var v301 int16 = (v6 >> ((v2) & 63))
        _ = v300
        _ = v301
    }
    if true {
        var v302 uint8 = 42
        if (false || true) {
            var v303 uint = 42
            fmt.Println(v296)
            _ = v303
        }
        _ = v302
    } else {
        switch v4 {
        case -9223372036854775808:
            v4 -= v4
        }
        fmt.Println(v296)
    }
    v304 := GenS19{F0: -20}
    v2 -= v2
    var v305 int = (v4 ^ (^(v4)))
    switch (v305 | (v4 / ((9223372036854775807) | 1))) {
    case -66:
        for fi306 := 0; fi306 < 17; fi306++ {
            switch "" {
            case "foo bar":
                fmt.Println(fi306)
            }
            switch (v4 & fi306) {
            case 128:
                var v307 int32 = 2147483647
                _ = v307
            case -664:
                for fi308 := 0; fi308 < 8; fi308++ {
                    v5 = (-(v5))
                    _ = fi308
                }
            default:
                if ((0 < -1) || ((false && true) && (1000 >= (v6 % ((v6) | 1))))) {
                    var v309 int64 = -480
                    v305 = int(v2)
                    var v310 uint32 = v296
                    _ = v309
                    _ = v310
                }
            }
            _ = fi306
        }
    case 734:
        var v311 int16 = (^(v6))
        _ = v311
    case 32768:
        for fi312 := 0; fi312 < 5; fi312++ {
            for fi313 := 0; fi313 < 12; fi313++ {
                var v314 uint16 = (v3 - (v3 + 255))
                fmt.Println(v6, v3, v5)
                if (2 == 255) {
                    v5 -= (v5 + v5)
                    switch "hello" {
                    case "a":
                        v315 := []int{v4, (fi312 >> (((v2 / ((v2) | 1))) & 63)), -1}
                        fmt.Println(v4, fi312)
                        _ = v315
                    default:
                        v3 += v314
                    }
                    var v316 bool = true
                    v305 -= (fi313 % (((v305 + v305)) | 1))
                    var v317 int16 = v6
                    _ = v316
                    _ = v317
                }
                var v318 float32 = (float32(float64((1.5 + float32((float64((((1.5 * 1.5) / 0.1) - -2.25)) + 1.5))))) - 0.0)
                _ = fi313
                _ = v314
                _ = v318
            }
            _ = fi312
        }
    default:
        for fi319 := 0; fi319 < 17; fi319++ {
            v2 -= uint(v3)
            for fi320 := 0; fi320 < 1; fi320++ {
                v6 = (-(v6))
                _ = fi320
            }
            var v321 uint16 = v3
            _ = fi319
            _ = v321
        }
    }
    for fi322 := 0; fi322 < 19; fi322++ {
        for fi323 := 0; fi323 < 15; fi323++ {
            fmt.Println(v3)
            v2 += v2
            fmt.Println(fi322, v305, v3)
            _ = fi323
        }
        _ = fi322
    }
    v6 -= (v6 ^ (v6 ^ v6))
    var v324 float64 = (0.0 * (1e10 - 0.1))
    v325 := []int{v4, 9223372036854775807, (v4 << ((255) & 63)), 42}
    if (false && (v2 != v2)) {
        fmt.Println(v305)
        v326 := GenS18{F0: "\n\t", F1: 0}
        _ = v326
    } else {
        if len(v325) >= 2 {
            v325 = v325[1:len(v325)]
        }
        v327 := GenS10{F0: 0, F1: (((v3 >> ((v2) & 63)) != (v3 % ((v3) | 1))) || false), F2: 100}
        _ = v327
    }
    v5 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v296:", v296)
    fmt.Println("v3:", v3)
    fmt.Println("v305:", v305)
    fmt.Println("v324:", v324)
    fmt.Println("slicelen v325:", len(v325), v325)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v296
    _ = v3
    _ = v304
    _ = v305
    _ = v324
    _ = v325
    _ = v4
    _ = v5
    _ = v6
}
