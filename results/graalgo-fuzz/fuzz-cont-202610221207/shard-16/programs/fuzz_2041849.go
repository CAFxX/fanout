package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int8; F1 string }
type GenS8 struct { F0 int; F1 int32; F2 int32 }
type GenS9 struct { F0 int16; F1 int; F2 bool }
type GenS10 struct { F0 uint16; F1 uint64; F2 uint }
type GenS11 struct { F0 int8; F1 uint64 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int; F1 int16 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int64 }
type GenS17 struct { F0 uint8; F1 int32; F2 int8 }
type GenS18 struct { F0 bool }
type GenS19 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS20 struct { F0 uint16; F1 string; F2 int }
type GenS21 struct { F0 uint; F1 int; F2 int32 }
type GenS22 struct { F0 int8; F1 uint; F2 string }
type GenS23 struct { F0 uint16; F1 int }
type GenS24 struct { F0 uint32; F1 int; F2 int }
type GenS25 struct { F0 int8; F1 string; F2 int32 }
type GenS26 struct { F0 int8; F1 uint16 }
type GenS27 struct { F0 string; F1 uint64 }
type GenS28 struct { F0 uint64; F1 bool }
type GenS29 struct { F0 int32; F1 bool; F2 uint16 }
type GenS30 struct { F0 int32 }
type GenS31 struct { F0 int8; F1 int8 }
type GenS32 struct { F0 bool; F1 int64; F2 int }
type GenS33 struct { F0 int64; F1 bool; F2 int32 }
type GenS34 struct { F0 uint32 }
type GenS35 struct { F0 string; F1 uint32; F2 int16 }
type GenS36 struct { F0 int; F1 int8 }
type GenS37 struct { F0 int }
type GenS38 struct { F0 int64; F1 uint8; F2 int }
type GenS39 struct { F0 int16; F1 int32 }
type GenS40 struct { F0 uint8 }
type GenS41 struct { F0 int32; F1 uint16; F2 int32 }
type GenS42 struct { F0 bool; F1 uint16; F2 int32 }
func genFn0(p0 int64) int {
    var v4 int = 9223372036854775807
    fmt.Println(v4, v4)
    var v5 uint = 355
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int16 = 1000
        v4 += fi6
        var v8 int8 = -46
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    return 32768
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]bool)
    var v3 int32 = -111
    if true {
        switch ("\n\t" + (((((("foo bar" + (("" + ("foo bar" + "foo bar")) + ("a" + "foo bar"))) + "foo bar") + "foo bar") + "") + "hello") + "")) {
        case "":
            v9 := GenS1{F0: (v1 % ((v1) | 1)), F1: 1, F2: 2}
            _ = v9
        case "\n\t":
            if (!true) {
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 uint = 2
                    if true {
                        var v12 uint64 = 1
                        var v13 string = "\n\t"
                        _ = v12
                        _ = v13
                    } else {
                        if ((v3 | v3) != 1) {
                            var v14 string = ""
                            fmt.Println(v11, v3)
                            var v15 uint = (v11 | 65535)
                            _ = v14
                            _ = v15
                        }
                        var v16 uint8 = 128
                        _ = v16
                    }
                    v1 = 42
                    _ = fi10
                    _ = v11
                }
                var v17 string = "hello"
                var v18 string = v17
                _ = v17
                _ = v18
            }
        case "a":
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS2{F0: v3}
                _ = fi19
                _ = v20
            }
        }
        v1 -= -1
    } else {
        if (true && ((!(false || (65535 == 42))) && (true && true))) {
            switch "" {
            case "":
                v21 := []int{42}
                _ = v21
            case "\n\t":
                v1 -= int64(v3)
            default:
                var v22 uint8 = 42
                _ = v22
            }
            v23 := GenS2{F0: (v3 ^ (v3 % ((v3) | 1)))}
            _ = v23
        } else {
            v1 = (v1 | (-(v1)))
        }
        if ((!false) || (132 >= 0)) {
            var v24 int64 = v1
            v25 := GenS1{F0: 174, F1: int8(v3), F2: 0}
            var v26 float32 = (0.1 + ((((0.0 + float32(((0.0 / 1e10) * 3.14159))) - 1.5) + 0.0) - 3.14159))
            _ = v24
            _ = v25
            _ = v26
        } else {
            fmt.Println(v3)
            if true {
                for fi27 := 0; fi27 < 12; fi27++ {
                    delete(v2, 65535)
                    v3 += (v3 * 0)
                    _ = fi27
                }
            } else {
                if true {
                    if true {
                        var v28 int32 = (v3 % (((v3 & 2147483647)) | 1))
                        if (!false) {
                            var v29 uint64 = 42
                            var v30 bool = ((v1 | v1) != (v1 << ((1) & 63)))
                            for fi31 := 0; fi31 < 18; fi31++ {
                                v28 = 970
                                _ = fi31
                            }
                            _ = v29
                            _ = v30
                        } else {
                            fmt.Println(v1, v1)
                            if (!false) {
                                switch "\n\t" {
                                case "\n\t":
                                    v1 += (v1 % ((-1) | 1))
                                case "foo bar":
                                    switch "hello" {
                                    case "a":
                                        v32 := GenS2{F0: v28}
                                        _ = v32
                                    case "foo bar":
                                        for fi33 := 0; fi33 < 14; fi33++ {
                                            v1 = (v1 | 879)
                                            _ = fi33
                                        }
                                    }
                                case "":
                                    v3 = v3
                                }
                                v34 := GenS3{F0: -1, F1: 0}
                                _ = v34
                            }
                            v3 -= (v3 * v3)
                        }
                        _ = v28
                    } else {
                        if true {
                            if (!(255 >= 42)) {
                                if true {
                                    if true {
                                        for fi35 := 0; fi35 < 17; fi35++ {
                                            if true {
                                                v36 := GenS4{F0: 27, F1: 4294967295}
                                                if false {
                                                    for fi37 := 0; fi37 < 7; fi37++ {
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v38 := GenS5{F0: (-(v1)), F1: 4294967295}
                                                            _ = v38
                                                        case "foo bar":
                                                            v1 -= v1
                                                        default:
                                                            v39 := GenS4{F0: 13, F1: 1}
                                                            _ = v39
                                                        }
                                                        if (127 <= 55) {
                                                            var v40 int32 = (v3 << ((921) & 63))
                                                            switch -9223372036854775808 {
                                                            case 127:
                                                                if (928 != 65535) {
                                                                    v41 := GenS6{F0: false, F1: -541}
                                                                    var v42 uint64 = 18446744073709551615
                                                                    var v43 int = (fi35 << ((4294967295) & 63))
                                                                    var v44 int32 = v40
                                                                    _ = v41
                                                                    _ = v42
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                            case 42:
                                                                v3 += (^(v3))
                                                            case 499:
                                                                fmt.Println(fi37)
                                                            default:
                                                                var v45 string = "hello"
                                                                _ = v45
                                                            }
                                                            var v46 int8 = 0
                                                            switch 1 {
                                                            case -1:
                                                                if (410 != 18446744073709551615) {
                                                                    v1 -= 0
                                                                    if (!false) {
                                                                        v47 := GenS3{F0: -153, F1: 2}
                                                                        var v48 uint64 = 0
                                                                        fmt.Println(v3, v48)
                                                                        _ = v47
                                                                        _ = v48
                                                                    } else {
                                                                        var v49 int32 = (v3 + (v40 * v40))
                                                                        if true {
                                                                            v2[848] = false
                                                                        }
                                                                        _ = v49
                                                                    }
                                                                }
                                                            default:
                                                                if (false || false) {
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        var v50 float64 = 1e10
                                                                        _ = v50
                                                                    case "hello":
                                                                        var v51 uint = 18446744073709551615
                                                                        _ = v51
                                                                    case "foo bar":
                                                                        v52 := v2[65535]
                                                                        _ = v52
                                                                    default:
                                                                        var v53 float64 = 0.0
                                                                        fmt.Println(v1, v40, v40)
                                                                        _ = v53
                                                                    }
                                                                    v54 := GenS7{F0: -86, F1: "\n\t"}
                                                                    if (1 != 0) {
                                                                        var v55 int32 = (v3 ^ -1)
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            if false {
                                                                                v56 := []int{(fi35 >> ((65535) & 63))}
                                                                                _ = v56
                                                                            } else {
                                                                                v55 += (v55 >> ((1) & 63))
                                                                                var v57 int64 = (v1 / ((-9223372036854775808) | 1))
                                                                                v58 := []int{1, fi37, fi37}
                                                                                var v59 uint16 = 0
                                                                                var v60 int64 = (v57 ^ (v1 & v1))
                                                                                _ = v57
                                                                                _ = v58
                                                                                _ = v59
                                                                                _ = v60
                                                                            }
                                                                        case "foo bar":
                                                                            fmt.Println(v55, v40, fi35)
                                                                        case "a":
                                                                            v61 := GenS4{F0: 59, F1: 2}
                                                                            _ = v61
                                                                        default:
                                                                            v3 = (v55 | v55)
                                                                        }
                                                                        var v62 int8 = v46
                                                                        _ = v55
                                                                        _ = v62
                                                                    }
                                                                    v40 = v3
                                                                    if (18446744073709551615 <= 2) {
                                                                        var v63 uint16 = 283
                                                                        _ = v63
                                                                    }
                                                                    if ((("" + ("foo bar" + "hello")) + "hello") >= ("foo bar" + ("hello" + "hello"))) {
                                                                        v3 = (v3 & v40)
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            v40 += 42
                                                                        case "a":
                                                                            v3 = (^(v40))
                                                                        }
                                                                        var v64 int16 = 42
                                                                        _ = v64
                                                                    } else {
                                                                        if (!((fi37 < (fi35 ^ 752)) || (!(v3 != v40)))) {
                                                                            v65 := []int{(fi37 - int(v1)), fi35, genFn0(int64(fi35)), 32767}
                                                                            if (((float32(0.1) * float32(((((0.1 + (0.1 * -2.25)) / 1.0) / 3.14159) * 3.14159))) > float32((1.5 * 1.5))) && (v1 < (v1 & -191))) {
                                                                                if len(v65) >= 2 {
                                                                                    v65 = v65[1:len(v65)]
                                                                                }
                                                                                v1 += (v1 / ((v1) | 1))
                                                                                if (!true) {
                                                                                    if true {
                                                                                        var v66 float64 = -2.25
                                                                                        _ = v66
                                                                                    }
                                                                                    if len(v65) >= 2 {
                                                                                        v65 = v65[1:len(v65)]
                                                                                    }
                                                                                    v67 := v2[804]
                                                                                    _ = v67
                                                                                } else {
                                                                                    v68 := GenS4{F0: 68, F1: 42}
                                                                                    _ = v68
                                                                                }
                                                                                if (!(18446744073709551615 != 2)) {
                                                                                    if (v3 > int32(fi37)) {
                                                                                        var v69 uint = 959
                                                                                        fmt.Println(v3, v1, v46)
                                                                                        v3 = v40
                                                                                        switch (-(fi37)) {
                                                                                        case -691:
                                                                                            switch (("a" + "") + "a") {
                                                                                            case "hello":
                                                                                                if (!(v3 <= (v3 & (-(v3))))) {
                                                                                                    if true {
                                                                                                        if true {
                                                                                                            if true {
                                                                                                                v46 = 39
                                                                                                                if false {
                                                                                                                    v70 := GenS7{F0: (v46 / ((127) | 1)), F1: ("" + "a")}
                                                                                                                    var v71 float32 = 3.14159
                                                                                                                    var v72 int = (fi37 & fi35)
                                                                                                                    _ = v70
                                                                                                                    _ = v71
                                                                                                                    _ = v72
                                                                                                                } else {
                                                                                                                    var v73 int8 = 0
                                                                                                                    v46 -= (v46 & v46)
                                                                                                                    _ = v73
                                                                                                                }
                                                                                                            }
                                                                                                            if (fi37 == (fi37 ^ (fi37 % ((fi35) | 1)))) {
                                                                                                                var v74 float64 = 1.5
                                                                                                                if true {
                                                                                                                    switch "foo bar" {
                                                                                                                    case "foo bar":
                                                                                                                        if true {
                                                                                                                            var v75 uint = v69
                                                                                                                            if true {
                                                                                                                                if (v3 != 2147483647) {
                                                                                                                                    v1 -= v1
                                                                                                                                    v3 += v40
                                                                                                                                } else {
                                                                                                                                    if false {
                                                                                                                                        fmt.Println(v3, v69, v74)
                                                                                                                                        var v76 int = fi37
                                                                                                                                        var v77 float32 = (3.14159 - -2.25)
                                                                                                                                        _ = v76
                                                                                                                                        _ = v77
                                                                                                                                    }
                                                                                                                                    var v78 uint = (v69 ^ v75)
                                                                                                                                    v78 = (v69 + 255)
                                                                                                                                    _ = v78
                                                                                                                                }
                                                                                                                                _ = len(v2)
                                                                                                                            } else {
                                                                                                                                v79 := GenS3{F0: 669, F1: 0}
                                                                                                                                v3 += (v40 & v40)
                                                                                                                                var v80 bool = true
                                                                                                                                _ = v79
                                                                                                                                _ = v80
                                                                                                                            }
                                                                                                                            _ = v75
                                                                                                                        }
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        var v81 float64 = v74
                                                                                                                        _ = v81
                                                                                                                    } else {
                                                                                                                        switch "" {
                                                                                                                        case "\n\t":
                                                                                                                            switch "foo bar" {
                                                                                                                            case "a":
                                                                                                                                switch "hello" {
                                                                                                                                case "a":
                                                                                                                                    var v82 int64 = (v1 & v1)
                                                                                                                                    _ = v82
                                                                                                                                case "foo bar":
                                                                                                                                    var v83 uint = (v69 << (((v69 * v69)) & 63))
                                                                                                                                    _ = v83
                                                                                                                                default:
                                                                                                                                    if (!(((1.5 / 0.1) + 0.0) > 1e10)) {
                                                                                                                                        v65 = append(v65, (fi35 * fi37))
                                                                                                                                    } else {
                                                                                                                                        v84 := GenS8{F0: fi35, F1: -219, F2: (v40 % (((v3 >> (((v69 | (^(v69)))) & 63))) | 1))}
                                                                                                                                        _ = v84
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                if (((v74 / float64(float32(v74))) != 1e10) || true) {
                                                                                                                                    if len(v65) >= 2 {
                                                                                                                                        v65 = v65[1:len(v65)]
                                                                                                                                    }
                                                                                                                                    _ = len(v2)
                                                                                                                                    var v85 uint32 = 4294967295
                                                                                                                                    _ = v85
                                                                                                                                }
                                                                                                                                if false {
                                                                                                                                    var v86 int8 = v46
                                                                                                                                    var v87 int64 = int64(v86)
                                                                                                                                    var v88 float64 = v74
                                                                                                                                    _ = v86
                                                                                                                                    _ = v87
                                                                                                                                    _ = v88
                                                                                                                                } else {
                                                                                                                                    if (42 != 0) {
                                                                                                                                        if (true || true) {
                                                                                                                                            v89 := GenS6{F0: false, F1: -32768}
                                                                                                                                            _ = v89
                                                                                                                                        } else {
                                                                                                                                            var v90 uint8 = 156
                                                                                                                                            if false {
                                                                                                                                                if ((((false || (!true)) || ((((!true) || (0.1 >= (float32((v74 / 1e10)) - 1e10))) && (!true)) || false)) || false) && (!false)) {
                                                                                                                                                    var v91 int32 = v3
                                                                                                                                                    _ = v91
                                                                                                                                                }
                                                                                                                                            } else {
                                                                                                                                                v92 := GenS6{F0: ((v46 + v46) != 100), F1: 0}
                                                                                                                                                _ = v92
                                                                                                                                            }
                                                                                                                                            _ = v90
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    if len(v65) >= 2 {
                                                                                                                                        v65 = v65[1:len(v65)]
                                                                                                                                    }
                                                                                                                                    switch fi37 {
                                                                                                                                    case 2147483647:
                                                                                                                                        v93 := GenS9{F0: 0, F1: (fi35 * fi35), F2: (0 >= 42)}
                                                                                                                                        _ = v93
                                                                                                                                    case 4294967295:
                                                                                                                                        var v94 string = ""
                                                                                                                                        _ = v94
                                                                                                                                    case -2147483648:
                                                                                                                                        v40 = v40
                                                                                                                                    default:
                                                                                                                                        v3 = (-(v40))
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                var v95 uint8 = 1
                                                                                                                                switch "foo bar" {
                                                                                                                                case "a":
                                                                                                                                    var v96 float32 = 0.0
                                                                                                                                    _ = v96
                                                                                                                                default:
                                                                                                                                    if (397 >= 2147483647) {
                                                                                                                                        if len(v65) > 0 {
                                                                                                                                            v65[0] = -818
                                                                                                                                        }
                                                                                                                                        v95 += v95
                                                                                                                                    } else {
                                                                                                                                        var v97 string = (("\n\t" + "") + ("" + ("foo bar" + ("foo bar" + ("" + "foo bar")))))
                                                                                                                                        switch fi35 {
                                                                                                                                        case -818:
                                                                                                                                            fmt.Println(v97, fi37)
                                                                                                                                        case 127:
                                                                                                                                            if false {
                                                                                                                                                fmt.Println(v95)
                                                                                                                                                v1 = v1
                                                                                                                                                v97 = v97
                                                                                                                                            } else {
                                                                                                                                                v1 = v1
                                                                                                                                                v98 := v2[841]
                                                                                                                                                v1 = 1
                                                                                                                                                fmt.Println(v95)
                                                                                                                                                var v99 bool = (!(!true))
                                                                                                                                                var v100 int64 = v1
                                                                                                                                                if (42 != v46) {
                                                                                                                                                    var v101 int32 = v3
                                                                                                                                                    var v102 float32 = (-2.25 - 1.5)
                                                                                                                                                    var v103 int64 = (v100 % ((v1) | 1))
                                                                                                                                                    if false {
                                                                                                                                                        var v104 int64 = v100
                                                                                                                                                        switch fi35 {
                                                                                                                                                        case 128:
                                                                                                                                                            v65 = append(v65, 32767)
                                                                                                                                                        case 0:
                                                                                                                                                            var v105 bool = (true && true)
                                                                                                                                                            _ = v105
                                                                                                                                                        case -2147483648:
                                                                                                                                                            if (246 <= 189) {
                                                                                                                                                                v3 = -1
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        _ = v104
                                                                                                                                                    } else {
                                                                                                                                                        v106 := GenS10{F0: 859, F1: 1, F2: 0}
                                                                                                                                                        _ = v106
                                                                                                                                                    }
                                                                                                                                                    _ = v101
                                                                                                                                                    _ = v102
                                                                                                                                                    _ = v103
                                                                                                                                                } else {
                                                                                                                                                    if len(v65) >= 2 {
                                                                                                                                                        v65 = v65[1:len(v65)]
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                                _ = v100
                                                                                                                                                _ = v98
                                                                                                                                                _ = v99
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            fmt.Println(v1)
                                                                                                                                            v107 := GenS7{F0: v46, F1: v97}
                                                                                                                                            _ = v107
                                                                                                                                        }
                                                                                                                                        _ = v97
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                _ = v95
                                                                                                                            }
                                                                                                                            if false {
                                                                                                                                var v108 string = ("\n\t" + "foo bar")
                                                                                                                                _ = v108
                                                                                                                            } else {
                                                                                                                                switch ("\n\t" + "\n\t") {
                                                                                                                                case "hello":
                                                                                                                                    var v109 uint8 = 27
                                                                                                                                    _ = v109
                                                                                                                                case "a":
                                                                                                                                    v40 = (v40 & v40)
                                                                                                                                    switch (-(fi35)) {
                                                                                                                                    case 128:
                                                                                                                                        v69 += v69
                                                                                                                                    default:
                                                                                                                                        if true {
                                                                                                                                            switch ("foo bar" + "a") {
                                                                                                                                            case "a":
                                                                                                                                                v74 += float64(3.14159)
                                                                                                                                            case "\n\t":
                                                                                                                                                if (!true) {
                                                                                                                                                    fmt.Println(fi35, v40, fi35)
                                                                                                                                                } else {
                                                                                                                                                    v110 := GenS11{F0: 7, F1: 18446744073709551615}
                                                                                                                                                    var v111 int = (fi35 / ((fi35) | 1))
                                                                                                                                                    var v112 int8 = v46
                                                                                                                                                    _ = v110
                                                                                                                                                    _ = v111
                                                                                                                                                    _ = v112
                                                                                                                                                }
                                                                                                                                            default:
                                                                                                                                                if true {
                                                                                                                                                    v46 -= (v46 & (v46 - v46))
                                                                                                                                                    if (!(!false)) {
                                                                                                                                                        v113 := GenS1{F0: (v1 << ((2) & 63)), F1: v46, F2: 2}
                                                                                                                                                        v114 := GenS10{F0: 65535, F1: 4294967295, F2: v69}
                                                                                                                                                        v40 += v40
                                                                                                                                                        _ = v113
                                                                                                                                                        _ = v114
                                                                                                                                                    } else {
                                                                                                                                                        if false {
                                                                                                                                                            v1 = 42
                                                                                                                                                        } else {
                                                                                                                                                            v69 = 255
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    if true {
                                                                                                                                                        v74 = v74
                                                                                                                                                        var v115 bool = false
                                                                                                                                                        v116 := GenS12{F0: (v40 + v3)}
                                                                                                                                                        if (!v115) {
                                                                                                                                                            if (!(!v115)) {
                                                                                                                                                                var v117 float64 = 0.1
                                                                                                                                                                v74 += float64(3.14159)
                                                                                                                                                                var v118 int = fi37
                                                                                                                                                                _ = v117
                                                                                                                                                                _ = v118
                                                                                                                                                            } else {
                                                                                                                                                                v119 := GenS13{F0: -1, F1: -1}
                                                                                                                                                                _ = v119
                                                                                                                                                            }
                                                                                                                                                            v1 = int64(v46)
                                                                                                                                                            delete(v2, 42)
                                                                                                                                                        }
                                                                                                                                                        _ = v115
                                                                                                                                                        _ = v116
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    if (!false) {
                                                                                                                                                        fmt.Println(v69, v46, v40)
                                                                                                                                                    }
                                                                                                                                                    v2[65535] = (1 <= 4294967295)
                                                                                                                                                    v120 := GenS14{F0: 2147483648}
                                                                                                                                                    _ = v120
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            v3 += int32(v1)
                                                                                                                                            var v121 int64 = (v1 % ((v1) | 1))
                                                                                                                                            _ = v121
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    var v122 int16 = 1
                                                                                                                                    _ = v122
                                                                                                                                case "foo bar":
                                                                                                                                    if len(v65) > 0 {
                                                                                                                                        v65[0] = int(v69)
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                var v123 uint8 = 128
                                                                                                                                fmt.Println(v3, v123, v46)
                                                                                                                                _ = v123
                                                                                                                            }
                                                                                                                            var v124 uint64 = 380
                                                                                                                            if true {
                                                                                                                                var v125 int64 = v1
                                                                                                                                v74 = v74
                                                                                                                                _ = v125
                                                                                                                            } else {
                                                                                                                                fmt.Println(v1, fi35)
                                                                                                                                v65 = append(v65, (fi37 & fi35))
                                                                                                                                var v126 uint32 = 2147483648
                                                                                                                                _ = v126
                                                                                                                            }
                                                                                                                            _ = v124
                                                                                                                        }
                                                                                                                        var v127 int64 = v1
                                                                                                                        _ = len(v2)
                                                                                                                        _ = v127
                                                                                                                    }
                                                                                                                }
                                                                                                                v65 = append(v65, (fi35 | (fi35 * fi35)))
                                                                                                                _ = v74
                                                                                                            }
                                                                                                            var v128 string = "foo bar"
                                                                                                            _ = v128
                                                                                                        }
                                                                                                        var v129 int64 = int64(fi35)
                                                                                                        _ = v129
                                                                                                    }
                                                                                                    var v130 float64 = 0.0
                                                                                                    if true {
                                                                                                        v131 := v2[42]
                                                                                                        v69 -= v69
                                                                                                        var v132 uint32 = 2147483648
                                                                                                        _ = v131
                                                                                                        _ = v132
                                                                                                    }
                                                                                                    var v133 uint = v69
                                                                                                    _ = v130
                                                                                                    _ = v133
                                                                                                }
                                                                                            default:
                                                                                                v134 := GenS15{F0: int16(v46)}
                                                                                                _ = v134
                                                                                            }
                                                                                        case 32768:
                                                                                            if ((^(v46)) == 127) {
                                                                                                v1 -= (v1 * (v1 << ((42) & 63)))
                                                                                            }
                                                                                        }
                                                                                        _ = v69
                                                                                    } else {
                                                                                        switch fi35 {
                                                                                        case 32767:
                                                                                            var v135 float32 = (float32((-2.25 / ((float64(1e10) + -2.25) / (0.1 - (0.1 * 1e10))))) * -2.25)
                                                                                            _ = v135
                                                                                        }
                                                                                        switch ("hello" + "hello") {
                                                                                        case "a":
                                                                                            var v136 float32 = -2.25
                                                                                            _ = v136
                                                                                        case "\n\t":
                                                                                            v1 -= v1
                                                                                        }
                                                                                        v137 := GenS4{F0: 1, F1: 18446744073709551615}
                                                                                        _ = v137
                                                                                    }
                                                                                    v2[42] = false
                                                                                    v1 -= v1
                                                                                } else {
                                                                                    switch genFn0(v1) {
                                                                                    case 0:
                                                                                        v40 = (v3 / (((-(v40))) | 1))
                                                                                    case 2147483647:
                                                                                        if len(v65) >= 2 {
                                                                                            v65 = v65[1:len(v65)]
                                                                                        }
                                                                                    default:
                                                                                        switch ("" + (("\n\t" + "") + "foo bar")) {
                                                                                        case "hello":
                                                                                            var v138 string = ("foo bar" + "foo bar")
                                                                                            _ = v138
                                                                                        case "":
                                                                                            v1 += (v1 >> ((0) & 63))
                                                                                        case "foo bar":
                                                                                            v139 := GenS11{F0: v46, F1: 4294967295}
                                                                                            _ = v139
                                                                                        default:
                                                                                            v140 := GenS15{F0: -32768}
                                                                                            _ = v140
                                                                                        }
                                                                                    }
                                                                                    var v141 float32 = (0.0 * -2.25)
                                                                                    _ = v141
                                                                                }
                                                                                if true {
                                                                                    if false {
                                                                                        var v142 string = ""
                                                                                        v143 := GenS10{F0: 397, F1: 23, F2: 42}
                                                                                        var v144 float32 = 1e10
                                                                                        switch genFn0(v1) {
                                                                                        case 1:
                                                                                            var v145 int64 = 1
                                                                                            _ = v145
                                                                                        case -2147483648:
                                                                                            fmt.Println(fi35, v144, v1)
                                                                                        case -9223372036854775808:
                                                                                            fmt.Println(fi37, fi35, v3)
                                                                                        default:
                                                                                            v146 := GenS3{F0: 1000, F1: 1}
                                                                                            _ = v146
                                                                                        }
                                                                                        v3 = (^(v40))
                                                                                        _ = v142
                                                                                        _ = v143
                                                                                        _ = v144
                                                                                    } else {
                                                                                        switch fi37 {
                                                                                        case -1:
                                                                                            v147 := GenS16{F0: (v1 - -445)}
                                                                                            _ = v147
                                                                                        case -543:
                                                                                            fmt.Println(v1, fi37)
                                                                                        case -2147483648:
                                                                                            switch "hello" {
                                                                                            case "hello":
                                                                                                var v148 uint16 = 42
                                                                                                _ = v148
                                                                                            case "\n\t":
                                                                                                var v149 uint16 = 65535
                                                                                                _ = v149
                                                                                            case "":
                                                                                                fmt.Println(fi37, v40)
                                                                                                var v150 uint = 1
                                                                                                v151 := GenS17{F0: 127, F1: (^(v40)), F2: v46}
                                                                                                v3 += v40
                                                                                                v152 := GenS18{F0: true}
                                                                                                var v153 bool = (true || false)
                                                                                                if len(v65) >= 2 {
                                                                                                    v65 = v65[1:len(v65)]
                                                                                                }
                                                                                                _ = v150
                                                                                                _ = v151
                                                                                                _ = v152
                                                                                                _ = v153
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    switch "foo bar" {
                                                                                    case "foo bar":
                                                                                        v46 = v46
                                                                                    default:
                                                                                        var v154 int64 = v1
                                                                                        _ = v154
                                                                                    }
                                                                                    v155 := GenS3{F0: 32767, F1: 59}
                                                                                    _ = len(v2)
                                                                                    v40 = v40
                                                                                    if (3.14159 <= ((0.1 + 0.0) + 0.0)) {
                                                                                        var v156 bool = false
                                                                                        v157 := GenS14{F0: 65535}
                                                                                        if (!v156) {
                                                                                            v46 = v46
                                                                                        } else {
                                                                                            var v158 uint16 = 676
                                                                                            _ = v158
                                                                                        }
                                                                                        _ = v156
                                                                                        _ = v157
                                                                                    } else {
                                                                                        switch "foo bar" {
                                                                                        case "hello":
                                                                                            fmt.Println(v3, v3, fi37)
                                                                                        case "\n\t":
                                                                                            v159 := GenS19{F0: 2, F1: 65535, F2: 65535}
                                                                                            _ = len(v2)
                                                                                            v1 = v1
                                                                                            _ = v159
                                                                                        }
                                                                                        v3 -= v3
                                                                                    }
                                                                                    v40 -= 42
                                                                                    _ = v155
                                                                                }
                                                                            } else {
                                                                                _ = len(v65)
                                                                            }
                                                                            if len(v65) >= 2 {
                                                                                v65 = v65[1:len(v65)]
                                                                            }
                                                                            _ = v65
                                                                        } else {
                                                                            var v160 bool = false
                                                                            _ = v160
                                                                        }
                                                                        var v161 int8 = (v46 % (((-(v46))) | 1))
                                                                        _ = v161
                                                                    }
                                                                    v162 := []int{genFn0(1), (fi35 / (((-(fi35))) | 1)), (fi35 * (-(fi35))), 127}
                                                                    _ = v162
                                                                    _ = v54
                                                                }
                                                            }
                                                            _ = v40
                                                            _ = v46
                                                        } else {
                                                            var v163 float32 = 3.14159
                                                            if false {
                                                                v3 -= v3
                                                                v164 := GenS19{F0: 42, F1: 255, F2: 65535}
                                                                if (!(true && true)) {
                                                                    var v165 int32 = (v3 | v3)
                                                                    if (false || true) {
                                                                        v166 := v2[42]
                                                                        fmt.Println(v163, fi35, v3)
                                                                        _ = v166
                                                                    } else {
                                                                        if (true && false) {
                                                                            v163 -= ((v163 - v163) * v163)
                                                                            var v167 uint32 = 4294967295
                                                                            _ = v167
                                                                        } else {
                                                                            var v168 float32 = ((v163 * v163) + v163)
                                                                            _ = len(v2)
                                                                            _ = v168
                                                                        }
                                                                        var v169 uint8 = 0
                                                                        _ = v169
                                                                    }
                                                                    if (v1 == (v1 | -892)) {
                                                                        v163 -= v163
                                                                        v1 += (^(v1))
                                                                        var v170 uint = 4294967295
                                                                        _ = v170
                                                                    } else {
                                                                        fmt.Println(v1)
                                                                        if (322 > 42) {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                var v171 int = genFn0(1)
                                                                                _ = v171
                                                                            }
                                                                        }
                                                                        switch ((("a" + "\n\t") + "foo bar") + "\n\t") {
                                                                        case "a":
                                                                            var v172 uint16 = 0
                                                                            _ = v172
                                                                        case "\n\t":
                                                                            v173 := []int{fi35, 9223372036854775807, (-(fi35))}
                                                                            _ = v173
                                                                        case "":
                                                                            v174 := []int{fi37, fi37, 94}
                                                                            _ = v174
                                                                        }
                                                                    }
                                                                    var v175 string = "a"
                                                                    switch (^(fi35)) {
                                                                    case 0:
                                                                        if (!(v175 != v175)) {
                                                                            switch (fi35 | (fi37 / ((fi37) | 1))) {
                                                                            case 1:
                                                                                var v176 string = v175
                                                                                _ = v176
                                                                            case -345:
                                                                                v177 := GenS20{F0: 0, F1: v175, F2: (fi35 * fi37)}
                                                                                _ = v177
                                                                            case 42:
                                                                                if true {
                                                                                    v178 := v2[1]
                                                                                    var v179 bool = v178
                                                                                    switch genFn0(v1) {
                                                                                    case 128:
                                                                                        v180 := []int{(fi37 + 4294967295), fi35, fi35}
                                                                                        v181 := GenS21{F0: 65535, F1: fi37, F2: 42}
                                                                                        v165 -= v165
                                                                                        v178 = false
                                                                                        _ = v180
                                                                                        _ = v181
                                                                                    case 127:
                                                                                        if v178 {
                                                                                            switch (((v175 + v175) + v175) + (v175 + (v175 + (("a" + v175) + "\n\t")))) {
                                                                                            case "":
                                                                                                switch genFn0(v1) {
                                                                                                case 677:
                                                                                                    switch "\n\t" {
                                                                                                    case "\n\t":
                                                                                                        v182 := GenS21{F0: 4294967295, F1: fi35, F2: (v3 + 42)}
                                                                                                        v3 = v165
                                                                                                        if true {
                                                                                                            if v178 {
                                                                                                                v179 = false
                                                                                                            }
                                                                                                            if (42 <= 1) {
                                                                                                                var v183 uint32 = 65535
                                                                                                                if v178 {
                                                                                                                    v184 := GenS21{F0: 0, F1: (fi37 & (fi35 + fi37)), F2: v3}
                                                                                                                    v165 = 666
                                                                                                                    fmt.Println(v163, fi35)
                                                                                                                    _ = v184
                                                                                                                } else {
                                                                                                                    if v178 {
                                                                                                                        v185 := GenS9{F0: 165, F1: genFn0((v1 & (v1 * 9223372036854775807))), F2: v179}
                                                                                                                        _ = v185
                                                                                                                    } else {
                                                                                                                        if v179 {
                                                                                                                            var v186 int32 = v165
                                                                                                                            fmt.Println(v179, v165, v175)
                                                                                                                            _ = v186
                                                                                                                        } else {
                                                                                                                            v179 = false
                                                                                                                            if false {
                                                                                                                                switch (fi35 | fi35) {
                                                                                                                                case 128:
                                                                                                                                    if v178 {
                                                                                                                                        v187 := GenS13{F0: (fi35 >> ((79) & 63)), F1: 1000}
                                                                                                                                        _ = v187
                                                                                                                                    } else {
                                                                                                                                        v183 = v183
                                                                                                                                    }
                                                                                                                                case 127:
                                                                                                                                    v188 := GenS22{F0: -128, F1: 65535, F2: ""}
                                                                                                                                    _ = v188
                                                                                                                                case 0:
                                                                                                                                    var v189 uint = 65535
                                                                                                                                    _ = v189
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v190 int = fi37
                                                                                                                            var v191 uint8 = 0
                                                                                                                            v2[v183] = (v178 && (!v179))
                                                                                                                            var v192 uint = 149
                                                                                                                            switch v175 {
                                                                                                                            case "foo bar":
                                                                                                                                if ((102 == -128) || (2 < (v191 << (((v192 / ((v192) | 1))) & 63)))) {
                                                                                                                                    v179 = (fi37 < 128)
                                                                                                                                }
                                                                                                                            case "hello":
                                                                                                                                if v179 {
                                                                                                                                    v193 := GenS16{F0: v1}
                                                                                                                                    _ = v193
                                                                                                                                } else {
                                                                                                                                    v1 += v1
                                                                                                                                    if v179 {
                                                                                                                                        var v194 uint32 = (v183 & (v183 + (v183 ^ 65535)))
                                                                                                                                        _ = v194
                                                                                                                                    } else {
                                                                                                                                        var v195 int8 = 110
                                                                                                                                        _ = v195
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            case "":
                                                                                                                                var v196 int8 = 42
                                                                                                                                _ = v196
                                                                                                                            }
                                                                                                                            v183 = (v183 ^ (^(v183)))
                                                                                                                            _ = v190
                                                                                                                            _ = v191
                                                                                                                            _ = v192
                                                                                                                        }
                                                                                                                        if v178 {
                                                                                                                            v197 := GenS6{F0: false, F1: -1}
                                                                                                                            v175 = ((((v175 + v175) + "\n\t") + (("" + (("hello" + ("hello" + ("foo bar" + v175))) + v175)) + v175)) + (((v175 + (v175 + "")) + (v175 + v175)) + "hello"))
                                                                                                                            _ = v197
                                                                                                                        } else {
                                                                                                                            v198 := GenS9{F0: -32768, F1: fi37, F2: (-405 != -741)}
                                                                                                                            var v199 string = (("foo bar" + v175) + (("\n\t" + v175) + (("a" + v175) + "a")))
                                                                                                                            v199 = v175
                                                                                                                            _ = v198
                                                                                                                            _ = v199
                                                                                                                        }
                                                                                                                        v200 := v2[1]
                                                                                                                        _ = v200
                                                                                                                    }
                                                                                                                }
                                                                                                                v201 := GenS10{F0: 0, F1: 1, F2: 626}
                                                                                                                _ = v183
                                                                                                                _ = v201
                                                                                                            }
                                                                                                            v202 := []int{-119, fi37, genFn0(v1), (fi37 & (fi35 << ((65535) & 63)))}
                                                                                                            _ = v202
                                                                                                        } else {
                                                                                                            var v203 int32 = v165
                                                                                                            var v204 float32 = 1.5
                                                                                                            _ = v203
                                                                                                            _ = v204
                                                                                                        }
                                                                                                        var v205 float64 = 1e10
                                                                                                        v205 -= 1e10
                                                                                                        v206 := []int{0}
                                                                                                        _ = v182
                                                                                                        _ = v205
                                                                                                        _ = v206
                                                                                                    }
                                                                                                case 5:
                                                                                                    v207 := []int{fi37}
                                                                                                    _ = v207
                                                                                                }
                                                                                            case "foo bar":
                                                                                                var v208 int = (fi37 * (fi35 >> ((379) & 63)))
                                                                                                _ = v208
                                                                                            case "a":
                                                                                                switch (-(fi35)) {
                                                                                                case 32768:
                                                                                                    var v209 uint = 4294967295
                                                                                                    if ((v3 <= v165) && v178) {
                                                                                                        var v210 int64 = v1
                                                                                                        var v211 float32 = 0.1
                                                                                                        _ = v210
                                                                                                        _ = v211
                                                                                                    } else {
                                                                                                        v212 := GenS23{F0: 32767, F1: (-(fi35))}
                                                                                                        _ = v212
                                                                                                    }
                                                                                                    var v213 uint8 = 128
                                                                                                    _ = v209
                                                                                                    _ = v213
                                                                                                case 9223372036854775807:
                                                                                                    v179 = (true || v178)
                                                                                                }
                                                                                                v214 := GenS18{F0: (369 == 753)}
                                                                                                _ = v214
                                                                                            default:
                                                                                                switch fi35 {
                                                                                                case -2147483648:
                                                                                                    var v215 uint64 = 260
                                                                                                    _ = v215
                                                                                                case -9223372036854775808:
                                                                                                    switch "foo bar" {
                                                                                                    case "a":
                                                                                                        fmt.Println(v178, v179)
                                                                                                    default:
                                                                                                        v216 := []int{(fi35 - (-(fi37))), 0, 1}
                                                                                                        _ = v216
                                                                                                    }
                                                                                                case 32767:
                                                                                                    var v217 uint8 = 42
                                                                                                    _ = v217
                                                                                                }
                                                                                            }
                                                                                            var v218 float32 = 3.14159
                                                                                            _ = v218
                                                                                        } else {
                                                                                            v165 = v3
                                                                                            var v219 uint64 = 2
                                                                                            v175 = v175
                                                                                            _ = v219
                                                                                        }
                                                                                    default:
                                                                                        v175 = "a"
                                                                                    }
                                                                                    _ = v178
                                                                                    _ = v179
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    if true {
                                                                                        var v220 int32 = (v3 >> ((685) & 63))
                                                                                        _ = v220
                                                                                    }
                                                                                    if true {
                                                                                        var v221 int = 32767
                                                                                        _ = v221
                                                                                    }
                                                                                    if false {
                                                                                        var v222 bool = false
                                                                                        v175 = v175
                                                                                        _ = v222
                                                                                    } else {
                                                                                        switch v175 {
                                                                                        case "\n\t":
                                                                                            if true {
                                                                                                if ((855 > v3) && (!true)) {
                                                                                                    var v223 uint = 386
                                                                                                    v224 := GenS4{F0: 42, F1: 42}
                                                                                                    _ = v223
                                                                                                    _ = v224
                                                                                                } else {
                                                                                                    var v225 uint8 = 42
                                                                                                    if false {
                                                                                                        _ = len(v2)
                                                                                                        if false {
                                                                                                            v163 = float32(1e10)
                                                                                                            var v226 uint16 = 65535
                                                                                                            _ = v226
                                                                                                        }
                                                                                                        if (0 == -1) {
                                                                                                            if true {
                                                                                                                if true {
                                                                                                                    if true {
                                                                                                                        switch v175 {
                                                                                                                        case "hello":
                                                                                                                            v227 := GenS24{F0: 612, F1: (fi35 % ((fi35) | 1)), F2: fi35}
                                                                                                                            _ = v227
                                                                                                                        default:
                                                                                                                            if (127 > 127) {
                                                                                                                                v228 := GenS25{F0: -33, F1: v175, F2: (v165 * v165)}
                                                                                                                                var v229 bool = false
                                                                                                                                _ = v228
                                                                                                                                _ = v229
                                                                                                                            } else {
                                                                                                                                v230 := v2[598]
                                                                                                                                v225 += 6
                                                                                                                                if (18446744073709551615 != 0) {
                                                                                                                                    v231 := GenS13{F0: fi35, F1: 32767}
                                                                                                                                    var v232 uint16 = 32767
                                                                                                                                    v163 = (1e10 - -2.25)
                                                                                                                                    fmt.Println(v163, v1)
                                                                                                                                    var v233 string = v175
                                                                                                                                    var v234 uint32 = 611
                                                                                                                                    _ = v231
                                                                                                                                    _ = v232
                                                                                                                                    _ = v233
                                                                                                                                    _ = v234
                                                                                                                                }
                                                                                                                                _ = v230
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v225 = (v225 >> ((42) & 63))
                                                                                                                    } else {
                                                                                                                        var v235 uint16 = 228
                                                                                                                        _ = v235
                                                                                                                    }
                                                                                                                    switch 2147483647 {
                                                                                                                    case 1:
                                                                                                                        if true {
                                                                                                                            switch v175 {
                                                                                                                            case "a":
                                                                                                                                var v236 uint = 304
                                                                                                                                _ = v236
                                                                                                                            case "\n\t":
                                                                                                                                v165 += -2147483648
                                                                                                                            case "":
                                                                                                                                v1 = v1
                                                                                                                            default:
                                                                                                                                v165 -= v165
                                                                                                                            }
                                                                                                                            _ = len(v2)
                                                                                                                        }
                                                                                                                    case -2147483648:
                                                                                                                        v163 -= 3.14159
                                                                                                                    }
                                                                                                                    v175 = (v175 + (("a" + (("a" + "hello") + (("a" + "a") + v175))) + ("hello" + "\n\t")))
                                                                                                                }
                                                                                                                switch ((v175 + v175) + ((v175 + ((v175 + v175) + ((v175 + ("" + v175)) + "\n\t"))) + ((((v175 + (((v175 + (("foo bar" + "foo bar") + v175)) + "hello") + "a")) + "a") + v175) + v175))) {
                                                                                                                case "":
                                                                                                                    if (!(v225 < 255)) {
                                                                                                                        var v237 uint64 = 0
                                                                                                                        var v238 string = (v175 + v175)
                                                                                                                        var v239 int8 = 100
                                                                                                                        _ = v237
                                                                                                                        _ = v238
                                                                                                                        _ = v239
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    var v240 int = (fi35 / ((-632) | 1))
                                                                                                                    _ = v240
                                                                                                                case "a":
                                                                                                                    v241 := GenS13{F0: fi35, F1: 1}
                                                                                                                    _ = v241
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        v163 -= v163
                                                                                                    }
                                                                                                    _ = v225
                                                                                                }
                                                                                                v242 := GenS23{F0: 976, F1: genFn0(v1)}
                                                                                                _ = v242
                                                                                            } else {
                                                                                                var v243 int32 = v3
                                                                                                v244 := GenS1{F0: (^(v1)), F1: -99, F2: 128}
                                                                                                _ = v243
                                                                                                _ = v244
                                                                                            }
                                                                                        case "a":
                                                                                            v245 := []int{(fi35 % (((-(fi37))) | 1)), fi37}
                                                                                            _ = v245
                                                                                        }
                                                                                        v246 := []int{-722, genFn0(v1), fi37, -288}
                                                                                        if true {
                                                                                            var v247 float32 = ((v163 + -2.25) * (v163 - v163))
                                                                                            var v248 uint64 = 444
                                                                                            _ = v247
                                                                                            _ = v248
                                                                                        }
                                                                                        _ = v246
                                                                                    }
                                                                                }
                                                                            }
                                                                            if (1000 != 42) {
                                                                                v1 = (v1 | (v1 % (((v1 << ((0) & 63))) | 1)))
                                                                                var v249 int16 = -558
                                                                                var v250 uint16 = 1
                                                                                _ = v249
                                                                                _ = v250
                                                                            } else {
                                                                                v251 := []int{-345, (-(fi37))}
                                                                                _ = v251
                                                                            }
                                                                            var v252 float32 = ((v163 * ((float32((float64(v163) / 3.14159)) - 1.5) + -2.25)) - (v163 / v163))
                                                                            if (734 <= 971) {
                                                                                if true {
                                                                                    v253 := GenS26{F0: 127, F1: 1}
                                                                                    v254 := []int{genFn0((v1 >> ((1) & 63))), fi37, genFn0(914), (fi35 + 4294967295)}
                                                                                    var v255 uint32 = 2147483648
                                                                                    _ = v253
                                                                                    _ = v254
                                                                                    _ = v255
                                                                                }
                                                                                var v256 int8 = 100
                                                                                var v257 float64 = 3.14159
                                                                                if true {
                                                                                    if (false || (42 == v165)) {
                                                                                        v3 += 1
                                                                                    }
                                                                                    v258 := GenS6{F0: (!false), F1: 1000}
                                                                                    v1 -= v1
                                                                                    if true {
                                                                                        fmt.Println(v3, v163, v175)
                                                                                    } else {
                                                                                        var v259 uint8 = 0
                                                                                        fmt.Println(v257)
                                                                                        if (1e10 <= 0.1) {
                                                                                            v3 -= v165
                                                                                            var v260 int32 = v3
                                                                                            if true {
                                                                                                var v261 uint32 = 42
                                                                                                v262 := GenS10{F0: 42, F1: 461, F2: 65535}
                                                                                                fmt.Println(v165)
                                                                                                v263 := []int{genFn0((v1 & -1))}
                                                                                                _ = v261
                                                                                                _ = v262
                                                                                                _ = v263
                                                                                            } else {
                                                                                                var v264 uint64 = 1
                                                                                                _ = v264
                                                                                            }
                                                                                            _ = v260
                                                                                        }
                                                                                        _ = v259
                                                                                    }
                                                                                    _ = v258
                                                                                } else {
                                                                                    v265 := GenS22{F0: v256, F1: 65535, F2: v175}
                                                                                    _ = v265
                                                                                }
                                                                                if true {
                                                                                    v266 := GenS11{F0: v256, F1: 1}
                                                                                    v2[2147483648] = (1 < 42)
                                                                                    switch 1 {
                                                                                    case 14:
                                                                                        var v267 int8 = (v256 / ((v256) | 1))
                                                                                        _ = v267
                                                                                    case 32768:
                                                                                        v268 := []int{(^(fi37)), genFn0(v1)}
                                                                                        _ = v268
                                                                                    }
                                                                                    _ = v266
                                                                                } else {
                                                                                    v269 := GenS27{F0: "\n\t", F1: 862}
                                                                                    if false {
                                                                                        if false {
                                                                                            v270 := GenS28{F0: 550, F1: ((((!false) || ((!(-32768 == 0)) || true)) || (((v1 > 42) && ((v256 <= (v256 & v256)) && (v257 < 3.14159))) && false)) && (!false))}
                                                                                            v271 := GenS29{F0: v165, F1: (!((v256 / ((-65) | 1)) != 19)), F2: 255}
                                                                                            _ = v270
                                                                                            _ = v271
                                                                                        } else {
                                                                                            v272 := GenS8{F0: (fi35 * genFn0(v1)), F1: 0, F2: v3}
                                                                                            fmt.Println(fi37, v256, v257)
                                                                                            if (!(((v257 * (1e10 - (0.0 / (float64((v252 + 0.0)) - (v257 / (((v257 + (v257 / 1e10)) * v257) * 0.0)))))) * v257) > float64(v252))) {
                                                                                                if false {
                                                                                                    v252 += v163
                                                                                                } else {
                                                                                                    var v273 int16 = 1
                                                                                                    var v274 bool = true
                                                                                                    v275 := GenS16{F0: int64(fi37)}
                                                                                                    var v276 float64 = v257
                                                                                                    v252 += v163
                                                                                                    _ = v273
                                                                                                    _ = v274
                                                                                                    _ = v275
                                                                                                    _ = v276
                                                                                                }
                                                                                            }
                                                                                            v277 := GenS30{F0: 0}
                                                                                            switch v175 {
                                                                                            case "":
                                                                                                var v278 int32 = v165
                                                                                                _ = v278
                                                                                            }
                                                                                            v3 += (v165 << ((4294967295) & 63))
                                                                                            _ = v272
                                                                                            _ = v277
                                                                                        }
                                                                                        v279 := GenS21{F0: 42, F1: 1, F2: 1}
                                                                                        if false {
                                                                                            var v280 float32 = float32(float64(float32(v257)))
                                                                                            v281 := GenS21{F0: 2, F1: int(v1), F2: 42}
                                                                                            var v282 uint = 1
                                                                                            _ = v280
                                                                                            _ = v281
                                                                                            _ = v282
                                                                                        }
                                                                                        v256 = (v256 - v256)
                                                                                        _ = v279
                                                                                    }
                                                                                    switch v175 {
                                                                                    case "a":
                                                                                        var v283 bool = false
                                                                                        _ = v283
                                                                                    case "hello":
                                                                                        v284 := GenS14{F0: 2}
                                                                                        _ = v284
                                                                                    case "foo bar":
                                                                                        v285 := []int{fi37, fi35, (fi35 >> ((211) & 63))}
                                                                                        _ = v285
                                                                                    }
                                                                                    _ = v269
                                                                                }
                                                                                _ = v256
                                                                                _ = v257
                                                                            }
                                                                            v1 = (v1 >> ((226) & 63))
                                                                            _ = v252
                                                                        } else {
                                                                            var v286 uint16 = 1
                                                                            v287 := v2[42]
                                                                            if (-1 < 0) {
                                                                                var v288 int64 = 0
                                                                                v175 = v175
                                                                                _ = v288
                                                                            } else {
                                                                                v175 = v175
                                                                            }
                                                                            _ = v286
                                                                            _ = v287
                                                                        }
                                                                    }
                                                                    _ = v165
                                                                    _ = v175
                                                                }
                                                                _ = v164
                                                            }
                                                            v1 -= 42
                                                            _ = v163
                                                        }
                                                        var v289 int16 = 42
                                                        _ = fi37
                                                        _ = v289
                                                    }
                                                }
                                                switch (fi35 ^ -2147483648) {
                                                case -129:
                                                    var v290 uint8 = 127
                                                    _ = v290
                                                }
                                                _ = v36
                                            } else {
                                                var v291 float32 = ((0.1 - -2.25) + (((0.1 * 1.5) * 0.0) + ((-2.25 / 1.5) * (1.5 + 1e10))))
                                                for fi292 := 0; fi292 < 15; fi292++ {
                                                    v1 -= int64(fi292)
                                                    v293 := GenS26{F0: 0, F1: 1}
                                                    _ = fi292
                                                    _ = v293
                                                }
                                                _ = v291
                                            }
                                            _ = fi35
                                        }
                                    }
                                    v2[24] = false
                                }
                            }
                            var v294 uint16 = 301
                            _ = v294
                        } else {
                            var v295 int8 = -127
                            for fi296 := 0; fi296 < 3; fi296++ {
                                v295 += 1
                                _ = fi296
                            }
                            fmt.Println(v3)
                            _ = v295
                        }
                        fmt.Println(v1, v1)
                        fmt.Println(v3)
                    }
                    if true {
                        v297 := GenS15{F0: -1}
                        _ = v297
                    }
                    var v298 uint = 978
                    _ = v298
                } else {
                    v299 := []int{-805}
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        v1 -= (v1 >> ((0) & 63))
                    case 32768:
                        for fi300 := 0; fi300 < 3; fi300++ {
                            var v301 int8 = 42
                            var v302 string = "foo bar"
                            v1 += v1
                            _ = fi300
                            _ = v301
                            _ = v302
                        }
                    }
                    if true {
                        var v303 string = "foo bar"
                        if true {
                            v304 := GenS29{F0: v3, F1: true, F2: 0}
                            for fi305 := 0; fi305 < 13; fi305++ {
                                fmt.Println(v1, v1, v303)
                                v306 := GenS31{F0: -128, F1: 1}
                                _ = fi305
                                _ = v306
                            }
                            _ = v304
                        }
                        switch 32767 {
                        case 32767:
                            v307 := GenS32{F0: true, F1: v1, F2: 2147483647}
                            _ = v307
                        default:
                            for fi308 := 0; fi308 < 3; fi308++ {
                                if false {
                                    if (false && true) {
                                        if true {
                                            v309 := v2[4294967295]
                                            var v310 int8 = 0
                                            var v311 uint = 928
                                            _ = v309
                                            _ = v310
                                            _ = v311
                                        } else {
                                            switch v303 {
                                            case "":
                                                v303 = (("hello" + ((v303 + ("\n\t" + ("foo bar" + (v303 + v303)))) + (v303 + ""))) + ((v303 + "a") + (v303 + v303)))
                                            default:
                                                v3 = -57
                                            }
                                        }
                                        for fi312 := 0; fi312 < 8; fi312++ {
                                            v313 := GenS33{F0: (-(v1)), F1: true, F2: (v3 ^ (v3 * v3))}
                                            if false {
                                                v303 = v303
                                                var v314 uint8 = 1
                                                _ = v314
                                            }
                                            _ = fi312
                                            _ = v313
                                        }
                                    }
                                    for fi315 := 0; fi315 < 7; fi315++ {
                                        v316 := GenS34{F0: 2147483648}
                                        v1 -= (v1 >> ((206) & 63))
                                        if (false || true) {
                                            if false {
                                                if (1 != 81) {
                                                    fmt.Println(fi308)
                                                }
                                            } else {
                                                v1 -= v1
                                                switch ((v303 + (("" + (("" + (((v303 + "a") + "\n\t") + v303)) + "a")) + "foo bar")) + v303) {
                                                case "\n\t":
                                                    if (false && false) {
                                                        delete(v2, 614)
                                                    } else {
                                                        var v317 int64 = -741
                                                        var v318 int32 = v3
                                                        if (1 >= 213) {
                                                            switch (fi308 ^ fi308) {
                                                            case -9223372036854775808:
                                                                var v319 int8 = 100
                                                                _ = v319
                                                            case 32767:
                                                                fmt.Println(v1)
                                                            case -2147483648:
                                                                var v320 float64 = float64(1.5)
                                                                _ = v320
                                                            default:
                                                                var v321 float64 = ((0.0 * (1e10 * ((0.1 + 1.5) - (0.1 / (0.1 + 1.5))))) + float64(1.5))
                                                                _ = v321
                                                            }
                                                            v303 = v303
                                                        } else {
                                                            var v322 int8 = 100
                                                            _ = v322
                                                        }
                                                        _ = v317
                                                        _ = v318
                                                    }
                                                case "":
                                                    v323 := GenS18{F0: (!false)}
                                                    _ = v323
                                                case "foo bar":
                                                    if (!(255 == 2)) {
                                                        if (((fi308 / ((fi315) | 1)) <= fi315) || true) {
                                                            v3 = (v3 | (v3 >> ((834) & 63)))
                                                            switch v303 {
                                                            case "a":
                                                                if true {
                                                                    var v324 uint16 = 65535
                                                                    v303 = v303
                                                                    if ((790 > 42) && (4294967295 >= 42)) {
                                                                        v325 := v2[2]
                                                                        fmt.Println(fi315, v1)
                                                                        _ = v325
                                                                    }
                                                                    _ = v324
                                                                }
                                                            }
                                                        } else {
                                                            if true {
                                                                var v326 int32 = (v3 & 0)
                                                                var v327 string = "a"
                                                                _ = v326
                                                                _ = v327
                                                            }
                                                        }
                                                        if false {
                                                            if ((2147483647 < 596) && ((1e10 + float64(1e10)) < (1e10 * (-2.25 - 0.1)))) {
                                                                var v328 int8 = 127
                                                                _ = v328
                                                            } else {
                                                                switch 0 {
                                                                case 2147483647:
                                                                    v303 = (v303 + "\n\t")
                                                                case 814:
                                                                    switch ("a" + ("a" + "hello")) {
                                                                    case "\n\t":
                                                                        v299 = append(v299, genFn0(-1))
                                                                    case "a":
                                                                        if len(v299) > 0 {
                                                                            v299[0] = -2147483648
                                                                        }
                                                                    default:
                                                                        if ((255 <= 1) && false) {
                                                                            var v329 bool = true
                                                                            fmt.Println(v3)
                                                                            _ = v329
                                                                        }
                                                                    }
                                                                case 538:
                                                                    v303 = ""
                                                                }
                                                            }
                                                        } else {
                                                            v1 -= (-(v1))
                                                        }
                                                    } else {
                                                        var v330 bool = ((v3 % (((^(v3))) | 1)) == 659)
                                                        fmt.Println(v330, v3, fi308)
                                                        switch fi308 {
                                                        case 28:
                                                            v1 = v1
                                                        case -1:
                                                            fmt.Println(v1, v303)
                                                        }
                                                        _ = v330
                                                    }
                                                }
                                            }
                                        }
                                        v331 := GenS35{F0: ("foo bar" + (v303 + (("hello" + (v303 + v303)) + v303))), F1: 2147483647, F2: 0}
                                        var v332 bool = true
                                        delete(v2, 2147483648)
                                        if (!(76 > 42)) {
                                            var v333 float32 = ((((float32(0.0) * 0.0) - -2.25) / 1.0) * float32(0.1))
                                            _ = v333
                                        }
                                        _ = fi315
                                        _ = v316
                                        _ = v331
                                        _ = v332
                                    }
                                    var v334 uint8 = 255
                                    _ = v334
                                } else {
                                    switch "\n\t" {
                                    case "foo bar":
                                        for fi335 := 0; fi335 < 5; fi335++ {
                                            if false {
                                                _ = len(v2)
                                                fmt.Println(v303, fi335)
                                            }
                                            _ = fi335
                                        }
                                    case "hello":
                                        switch ((v303 + v303) + "\n\t") {
                                        case "\n\t":
                                            for fi336 := 0; fi336 < 3; fi336++ {
                                                var v337 bool = ((-(fi308)) == (fi336 % ((4294967295) | 1)))
                                                v338 := GenS31{F0: -1, F1: -56}
                                                _ = len(v2)
                                                _ = fi336
                                                _ = v337
                                                _ = v338
                                            }
                                        }
                                    case "\n\t":
                                        if false {
                                            delete(v2, 2147483648)
                                            v339 := GenS29{F0: (v3 % (((-(v3))) | 1)), F1: true, F2: 32767}
                                            var v340 int64 = v1
                                            _ = v339
                                            _ = v340
                                        }
                                    default:
                                        if (false || (!false)) {
                                            if true {
                                                for fi341 := 0; fi341 < 7; fi341++ {
                                                    v1 = (v1 >> ((18446744073709551615) & 63))
                                                    _ = fi341
                                                }
                                                if (!true) {
                                                    v342 := GenS36{F0: 891, F1: 0}
                                                    var v343 uint8 = 110
                                                    delete(v2, 42)
                                                    _ = v342
                                                    _ = v343
                                                } else {
                                                    v344 := GenS7{F0: 27, F1: "foo bar"}
                                                    if true {
                                                        v345 := GenS37{F0: 32768}
                                                        _ = v345
                                                    }
                                                    var v346 int = fi308
                                                    _ = v344
                                                    _ = v346
                                                }
                                            } else {
                                                var v347 uint = 42
                                                _ = v347
                                            }
                                            var v348 int64 = (v1 / ((v1) | 1))
                                            _ = v348
                                        } else {
                                            var v349 int16 = -32768
                                            var v350 string = "a"
                                            _ = v349
                                            _ = v350
                                        }
                                    }
                                    if (200 == 2) {
                                        var v351 int64 = v1
                                        v352 := GenS22{F0: 1, F1: 65535, F2: v303}
                                        _ = v351
                                        _ = v352
                                    } else {
                                        v353 := GenS38{F0: 1, F1: 2, F2: (fi308 | (fi308 / (((fi308 / (((^(fi308))) | 1))) | 1)))}
                                        _ = v353
                                    }
                                }
                                var v354 uint32 = 42
                                switch 9223372036854775807 {
                                case -716:
                                    var v355 int64 = (v1 % ((-632) | 1))
                                    _ = v355
                                case 0:
                                    v356 := GenS39{F0: 1000, F1: 0}
                                    _ = v356
                                case 42:
                                    delete(v2, (v354 + (^(v354))))
                                }
                                _ = fi308
                                _ = v354
                            }
                        }
                        _ = v303
                    } else {
                        var v357 uint8 = 2
                        v358 := GenS34{F0: 52}
                        _ = v357
                        _ = v358
                    }
                    _ = v299
                }
                var v359 int16 = 1000
                for fi360 := 0; fi360 < 20; fi360++ {
                    v361 := []int{}
                    v362 := GenS39{F0: v359, F1: (v3 >> ((4294967295) & 63))}
                    _ = fi360
                    _ = v361
                    _ = v362
                }
                _ = v359
            }
            if true {
                var v363 uint16 = 255
                for fi364 := 0; fi364 < 7; fi364++ {
                    var v365 int64 = (v1 - (v1 % ((42) | 1)))
                    _ = fi364
                    _ = v365
                }
                v1 = (v1 & (v1 + (v1 ^ v1)))
                _ = v363
            }
        }
    }
    _ = len(v2)
    v1 = -1
    fmt.Println(v3, v1)
    if false {
        if (!false) {
            switch 735 {
            case -2147483648:
                v366 := GenS37{F0: 766}
                _ = v366
            case 32767:
                v2[42] = false
            default:
                switch "" {
                case "":
                    v367 := GenS40{F0: 1}
                    _ = v367
                case "a":
                    v368 := GenS40{F0: 2}
                    _ = v368
                default:
                    v369 := v2[960]
                    _ = v369
                }
            }
        } else {
            if false {
                if ((!(false && ((true || false) || (!true)))) || true) {
                    switch 32768 {
                    case 32767:
                        var v370 int64 = 42
                        _ = v370
                    default:
                        var v371 string = "\n\t"
                        _ = v371
                    }
                }
                v3 += 605
                v3 += (v3 & 0)
            }
        }
        var v372 float32 = 1.5
        _ = v372
    }
    if (!true) {
        if ("foo bar" <= "") {
            var v373 float64 = 0.0
            v374 := GenS41{F0: (v3 % (((-(v3))) | 1)), F1: 2, F2: (-(v3))}
            _ = v373
            _ = v374
        }
    } else {
        fmt.Println(v3, v3)
        v375 := []int{128, 481, 128, 2147483647}
        var v376 uint16 = 711
        _ = v375
        _ = v376
    }
    v377 := GenS42{F0: false, F1: 1, F2: v3}
    var v378 string = "foo bar"
    v379 := v2[0]
    v380 := []int{9223372036854775807, 32767, 42, 9223372036854775807}
    if len(v380) > 0 {
        v380[0] = 9223372036854775807
    }
    var v381 uint = 18446744073709551615
    switch "foo bar" {
    case "":
        v378 = ((((v378 + v378) + v378) + v378) + v378)
    default:
        var v382 int32 = int32(v1)
        _ = v382
    }
    if (true && (42 > 0)) {
        var v383 int8 = 0
        _ = v383
    }
    if len(v380) > 0 {
        v380[0] = genFn0(544)
    }
    var v384 bool = (!(127 != 42))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v378:", v378)
    fmt.Println("v379:", v379)
    fmt.Println("slicelen v380:", len(v380), v380)
    fmt.Println("v381:", v381)
    fmt.Println("v384:", v384)
    _ = v1
    _ = v2
    _ = v3
    _ = v377
    _ = v378
    _ = v379
    _ = v380
    _ = v381
    _ = v384
}
