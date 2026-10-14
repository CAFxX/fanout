package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int16; F2 int16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 string; F2 uint32 }
type GenS7 struct { F0 int8; F1 uint }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 int; F1 uint8; F2 uint }
type GenS11 struct { F0 int32; F1 uint8; F2 uint8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int64; F1 string; F2 bool }
type GenS14 struct { F0 int32; F1 bool }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS18 struct { F0 uint; F1 uint64; F2 int64 }
func genFn0(p0 uint) uint64 {
    fmt.Println(p0)
    fmt.Println(p0)
    v6 := []int{32768}
    var v7 int16 = 1000
    _ = p0
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 string, p1 int16) int {
    p0 = ""
    switch p0 {
    case "\n\t":
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 float32 = 1.5
            fmt.Println(fi8)
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := GenS1{F0: 0, F1: 32767, F2: 2}
                var v12 float32 = (1.5 / v9)
                v12 = ((0.1 / 1.5) + float32((((((1e10 - (((((1e10 / float64(((v9 - 0.0) + (((3.14159 + v9) * v9) - ((v12 - v9) * (((v12 * 3.14159) / v9) + ((v12 * 1.5) - 0.1))))))) - 0.0) - (1.5 - (float64((v9 / (v12 / v9))) - float64(float32((((((float64(v9) + 3.14159) - 0.1) + ((-2.25 * float64(v12)) + (0.1 * 3.14159))) - ((0.0 - float64(float32(3.14159))) * (3.14159 * 1e10))) / -2.25)))))) * 1.5) * float64(((v12 * v9) * ((1.5 - v12) * v12))))) - float64(v9)) * ((1.5 / -2.25) + ((-2.25 - ((0.0 / 3.14159) / (((-2.25 / 1.0) + ((3.14159 / -2.25) - float64(v12))) * 0.1))) - 1.5))) - 3.14159) / 0.1)))
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
    default:
        switch ("\n\t" + p0) {
        case "":
            var v13 uint = 255
            _ = v13
        case "a":
            switch -303 {
            case -746:
                switch p0 {
                case "":
                    p1 -= p1
                case "foo bar":
                    for fi14 := 0; fi14 < 3; fi14++ {
                        var v15 int16 = p1
                        v16 := []int{}
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                }
            default:
                p0 = "a"
            }
        case "hello":
            p0 = (p0 + ((p0 + "a") + "\n\t"))
        default:
            v17 := []int{-106}
            _ = v17
        }
    }
    p0 = ((p0 + (p0 + ((p0 + (("" + (p0 + (((p0 + ("\n\t" + "\n\t")) + (p0 + ("\n\t" + p0))) + p0))) + ((("a" + "") + "foo bar") + p0))) + "hello"))) + "foo bar")
    p0 = p0
    p1 -= -1
    _ = p0
    _ = p1
    return 871
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 189
    var v3 uint8 = (v2 * (^(v2)))
    v4 := []uint32{4294967295}
    v5 := []int{57}
    var v18 int32 = -2147483648
    var v19 uint8 = v2
    for fi20 := 0; fi20 < 8; fi20++ {
        switch "\n\t" {
        case "foo bar":
            _ = len(v5)
        default:
            v21 := GenS2{F0: 0}
            _ = v21
        }
        _ = fi20
    }
    v22 := GenS3{F0: v18, F1: 0, F2: 0}
    fmt.Println(v3, v3, v2)
    var v23 uint32 = 574
    switch "" {
    case "":
        v23 = 42
    case "hello":
        fmt.Println(v18)
    }
    switch 2147483647 {
    case 98:
        v3 = v3
    case 42:
        var v24 float32 = ((1.5 + -2.25) - float32(0.1))
        _ = v24
    }
    v2 = 128
    if ((v23 % (((v23 - (v23 / (((^(v23))) | 1)))) | 1)) < (v23 / (((^(v23))) | 1))) {
        v23 = uint32(v3)
        if (!((((v18 >> ((1) & 63)) <= (-(v18))) || true) || ((false || true) || false))) {
            if len(v4) > 0 {
                v4[0] = v23
            }
            fmt.Println(v2)
        } else {
            var v25 uint32 = 4294967295
            _ = len(v4)
            if (true && true) {
                v19 -= (v19 << ((118) & 63))
                var v26 uint8 = 0
                _ = v26
            }
            _ = v25
        }
    } else {
        switch "foo bar" {
        case "\n\t":
            switch "hello" {
            case "":
                var v27 bool = false
                _ = v27
            case "a":
                if false {
                    v2 += v3
                    var v28 uint8 = 177
                    v28 = (^(v3))
                    _ = v28
                } else {
                    var v29 uint = 42
                    v18 += 0
                    _ = v29
                }
            case "\n\t":
                if (554 != 42) {
                    for fi30 := 0; fi30 < 20; fi30++ {
                        if (!((^(v3)) <= (v19 - (v19 * 0)))) {
                            fmt.Println(v2, v2)
                            _ = len(v1)
                            var v31 int16 = 1000
                            _ = v31
                        }
                        _ = fi30
                    }
                }
            default:
                var v32 uint8 = v2
                _ = v32
            }
        }
        switch 32767 {
        case 127:
            v18 -= (v18 / ((v18) | 1))
        case 640:
            for fi33 := 0; fi33 < 16; fi33++ {
                if false {
                    for fi34 := 0; fi34 < 17; fi34++ {
                        fmt.Println(v19)
                        var v35 uint8 = v3
                        v36 := GenS1{F0: 0, F1: 42, F2: v23}
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    var v37 uint32 = 142
                    var v38 uint = 4294967295
                    _ = v37
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 17; fi39++ {
                        var v40 int64 = 42
                        _ = len(v4)
                        if false {
                            v41 := GenS1{F0: v19, F1: 65535, F2: v23}
                            v18 = -1
                            _ = v41
                        }
                        _ = fi39
                        _ = v40
                    }
                    switch (fi33 & int(v23)) {
                    case -82:
                        v42 := GenS1{F0: (^(v2)), F1: 42, F2: (^(v23))}
                        _ = v42
                    case 127:
                        if false {
                            fmt.Println(v23, fi33, v23)
                            if true {
                                if ((255 == 965) && true) {
                                    if false {
                                        for fi43 := 0; fi43 < 16; fi43++ {
                                            v23 = 505
                                            _ = fi43
                                        }
                                        v19 = v3
                                    } else {
                                        if ((0.0 - 0.1) >= 0.0) {
                                            switch 1 {
                                            case 4294967295:
                                                for fi44 := 0; fi44 < 14; fi44++ {
                                                    var v45 int64 = int64(v23)
                                                    if true {
                                                        if false {
                                                            var v46 int32 = v18
                                                            var v47 int = -262
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                    } else {
                                                        if (true || ((1.5 / float64(1.5)) < 0.1)) {
                                                            if (((v45 >= v45) && ((!(0 > 42)) || true)) || false) {
                                                                var v48 float32 = (((3.14159 - 0.1) * 1.5) / (0.1 / 1.0))
                                                                fmt.Println(v23, v18, v2)
                                                                if false {
                                                                    var v49 bool = true
                                                                    _ = v49
                                                                } else {
                                                                    switch ("a" + "hello") {
                                                                    case "hello":
                                                                        v50 := GenS4{F0: v45}
                                                                        _ = v50
                                                                    case "":
                                                                        if true {
                                                                            fmt.Println(fi33)
                                                                            v51 := v1[9223372036854775807]
                                                                            var v52 int64 = (v45 ^ 377)
                                                                            _ = v51
                                                                            _ = v52
                                                                        } else {
                                                                            if false {
                                                                                fmt.Println(fi44)
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 2
                                                                                }
                                                                                v53 := GenS5{F0: "foo bar"}
                                                                                _ = v53
                                                                            } else {
                                                                                v45 -= v45
                                                                            }
                                                                        }
                                                                    case "foo bar":
                                                                        v19 = (v3 & (v19 + v19))
                                                                    default:
                                                                        v45 = v45
                                                                    }
                                                                    if ((v45 >> ((18446744073709551615) & 63)) > 42) {
                                                                        v18 = v18
                                                                        if false {
                                                                            if false {
                                                                                v19 = 127
                                                                            } else {
                                                                                v54 := GenS6{F0: 0, F1: "", F2: 1}
                                                                                v18 += v18
                                                                                _ = v54
                                                                            }
                                                                        } else {
                                                                            v48 += v48
                                                                            fmt.Println(fi33, v18)
                                                                        }
                                                                    }
                                                                    v55 := GenS5{F0: "a"}
                                                                    _ = v55
                                                                }
                                                                var v56 float64 = (1.5 * 1.5)
                                                                _ = v48
                                                                _ = v56
                                                            } else {
                                                                v57 := GenS7{F0: 0, F1: 2}
                                                                if (!true) {
                                                                    var v58 string = ("hello" + "hello")
                                                                    if false {
                                                                        if ((!(v18 > v18)) || ((fi33 % ((fi44) | 1)) == (fi33 + genFn1(v58, 0)))) {
                                                                            delete(v1, (fi44 << ((18446744073709551615) & 63)))
                                                                            if (false && false) {
                                                                                var v59 int = fi44
                                                                                v23 += 257
                                                                                var v60 float32 = 1.5
                                                                                _ = v59
                                                                                _ = v60
                                                                            }
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                var v61 int = 0
                                                                                _ = v61
                                                                            }
                                                                        } else {
                                                                            var v62 int8 = 0
                                                                            if false {
                                                                                v18 = (v18 + v18)
                                                                                v18 = int32(v62)
                                                                            }
                                                                            _ = v62
                                                                        }
                                                                        var v63 float64 = (1.5 / -2.25)
                                                                        _ = v63
                                                                    } else {
                                                                        v2 = 1
                                                                        fmt.Println(v18)
                                                                    }
                                                                    switch fi44 {
                                                                    case 0:
                                                                        v64 := GenS6{F0: 214, F1: "foo bar", F2: 2147483647}
                                                                        _ = v64
                                                                    }
                                                                    _ = v58
                                                                } else {
                                                                    var v65 float32 = 0.1
                                                                    fmt.Println(fi44, fi33, v3)
                                                                    if false {
                                                                        var v66 float64 = 1.5
                                                                        _ = v66
                                                                    }
                                                                    _ = v65
                                                                }
                                                                _ = len(v4)
                                                                v67 := GenS3{F0: v18, F1: 272, F2: -1}
                                                                if true {
                                                                    var v68 bool = (!false)
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        switch "" {
                                                                        case "hello":
                                                                            var v69 int = (fi44 ^ (^(fi44)))
                                                                            _ = v69
                                                                        case "\n\t":
                                                                            v70 := GenS8{F0: 282}
                                                                            _ = v70
                                                                        default:
                                                                            v19 -= v3
                                                                        }
                                                                    case "":
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    case "\n\t":
                                                                        switch ("\n\t" + "foo bar") {
                                                                        case "a":
                                                                            v2 -= 127
                                                                        case "hello":
                                                                            fmt.Println(v3, v18)
                                                                        case "\n\t":
                                                                            if (!v68) {
                                                                                _ = len(v4)
                                                                                var v71 int16 = -1
                                                                                v72 := GenS9{F0: 4294967295, F1: 0}
                                                                                v18 += v18
                                                                                _ = v71
                                                                                _ = v72
                                                                            } else {
                                                                                switch "hello" {
                                                                                case "":
                                                                                    v73 := GenS10{F0: 128, F1: v3, F2: 538}
                                                                                    var v74 float64 = -2.25
                                                                                    v74 += 0.0
                                                                                    v75 := GenS10{F0: genFn1("a", -468), F1: (v3 + 2), F2: 42}
                                                                                    var v76 uint64 = genFn0(521)
                                                                                    switch (fi44 % (((fi44 + 0)) | 1)) {
                                                                                    case -577:
                                                                                        v23 = v23
                                                                                    case -9223372036854775808:
                                                                                        switch (fi44 | fi44) {
                                                                                        case 356:
                                                                                            if true {
                                                                                                switch "\n\t" {
                                                                                                case "foo bar":
                                                                                                    if (false && true) {
                                                                                                        if v68 {
                                                                                                            delete(v1, (-(fi33)))
                                                                                                            v77 := GenS11{F0: -337, F1: (v3 * 127), F2: (v19 >> ((718) & 63))}
                                                                                                            _ = v77
                                                                                                        }
                                                                                                    }
                                                                                                case "hello":
                                                                                                    fmt.Println(v45, v68, fi44)
                                                                                                case "\n\t":
                                                                                                    switch int(v23) {
                                                                                                    case -9223372036854775808:
                                                                                                        delete(v1, fi33)
                                                                                                        _ = len(v1)
                                                                                                        v45 = (v45 - v45)
                                                                                                        var v78 int64 = -9223372036854775808
                                                                                                        var v79 int64 = (v45 ^ (v78 - v78))
                                                                                                        _ = v78
                                                                                                        _ = v79
                                                                                                    case -1:
                                                                                                        v80 := GenS7{F0: 127, F1: 255}
                                                                                                        _ = v80
                                                                                                    default:
                                                                                                        var v81 int8 = -110
                                                                                                        _ = v81
                                                                                                    }
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = -9223372036854775808
                                                                                                    }
                                                                                                    delete(v1, fi44)
                                                                                                }
                                                                                                v82 := GenS4{F0: (v45 | (v45 - v45))}
                                                                                                var v83 uint32 = (v23 / (((v23 >> ((711) & 63))) | 1))
                                                                                                var v84 uint8 = (^(v2))
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                            }
                                                                                        case 4294967295:
                                                                                            v2 = v3
                                                                                        default:
                                                                                            if true {
                                                                                                delete(v1, fi44)
                                                                                                v5 = append(v5, (fi44 >> ((0) & 63)))
                                                                                                var v85 uint32 = 2147483648
                                                                                                _ = v85
                                                                                            } else {
                                                                                                var v86 int64 = -9223372036854775808
                                                                                                var v87 bool = v68
                                                                                                if (((v45 % (((v45 + (v86 % ((308) | 1)))) | 1)) > (v45 / ((v86) | 1))) && (float32(v74) <= 3.14159)) {
                                                                                                    if (false || v87) {
                                                                                                        v88 := GenS9{F0: 4294967295, F1: 1}
                                                                                                        v19 -= 1
                                                                                                        _ = v88
                                                                                                    } else {
                                                                                                        var v89 int64 = -372
                                                                                                        v90 := GenS2{F0: 0}
                                                                                                        fmt.Println(v76)
                                                                                                        delete(v1, int(v23))
                                                                                                        _ = v89
                                                                                                        _ = v90
                                                                                                    }
                                                                                                }
                                                                                                _ = v86
                                                                                                _ = v87
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v73
                                                                                    _ = v74
                                                                                    _ = v75
                                                                                    _ = v76
                                                                                case "a":
                                                                                    v91 := GenS12{F0: 65535}
                                                                                    fmt.Println(v45, v18)
                                                                                    _ = v91
                                                                                case "hello":
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = genFn1("hello", 978)
                                                                                    }
                                                                                }
                                                                                fmt.Println(v45)
                                                                            }
                                                                        default:
                                                                            v5 = append(v5, fi33)
                                                                        }
                                                                    }
                                                                    var v92 float64 = 1.5
                                                                    _ = v68
                                                                    _ = v92
                                                                }
                                                                switch genFn1("\n\t", 32767) {
                                                                case 32767:
                                                                    var v93 uint = 2
                                                                    var v94 int32 = -1
                                                                    fmt.Println(fi44, v3, v45)
                                                                    _ = v93
                                                                    _ = v94
                                                                case -551:
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        var v95 int = (-(fi33))
                                                                        _ = v95
                                                                    case "hello":
                                                                        delete(v1, fi44)
                                                                    default:
                                                                        var v96 int8 = 1
                                                                        _ = v96
                                                                    }
                                                                default:
                                                                    v19 = v2
                                                                }
                                                                _ = v57
                                                                _ = v67
                                                            }
                                                            v19 -= v2
                                                            if (!false) {
                                                                var v97 uint16 = 243
                                                                var v98 int16 = 1
                                                                fmt.Println(v19)
                                                                _ = v97
                                                                _ = v98
                                                            } else {
                                                                v3 -= v2
                                                                fmt.Println(v23, v19)
                                                                switch (fi33 + (-(fi33))) {
                                                                case -2147483648:
                                                                    var v99 float64 = 0.0
                                                                    _ = v99
                                                                case 32768:
                                                                    v45 = (v45 / ((int64(fi33)) | 1))
                                                                case 42:
                                                                    fmt.Println(fi33, v19, v23)
                                                                }
                                                                switch ("\n\t" + "") {
                                                                case "hello":
                                                                    v100 := GenS7{F0: 127, F1: 42}
                                                                    _ = v100
                                                                case "a":
                                                                    v1[(fi44 >> ((294) & 63))] = ""
                                                                case "foo bar":
                                                                    if ((true && (!false)) && (65535 <= 65535)) {
                                                                        v18 = (v18 ^ (-(v18)))
                                                                        v45 -= v45
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            var v101 float64 = 0.1
                                                                            _ = v101
                                                                        }
                                                                        v102 := GenS3{F0: (^(v18)), F1: 42, F2: 32767}
                                                                        var v103 uint = 0
                                                                        _ = v102
                                                                        _ = v103
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            if ((((3.14159 / (-2.25 * ((float32(-2.25) / 1e10) - 1.5))) - (0.0 + 1e10)) / 3.14159) < 3.14159) {
                                                                v19 -= v2
                                                                fmt.Println(v3)
                                                            }
                                                            if len(v5) > 0 {
                                                                v5[0] = fi44
                                                            }
                                                        }
                                                        v104 := GenS13{F0: v45, F1: "a", F2: false}
                                                        switch "foo bar" {
                                                        case "hello":
                                                            fmt.Println(v2)
                                                        case "a":
                                                            if true {
                                                                if len(v5) > 0 {
                                                                    v5[0] = (fi33 * fi44)
                                                                }
                                                                switch -2147483648 {
                                                                case -676:
                                                                    v105 := GenS14{F0: 2147483647, F1: (!false)}
                                                                    _ = v105
                                                                }
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            }
                                                        case "":
                                                            fmt.Println(v19, fi33, v3)
                                                        default:
                                                            if false {
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    v106 := GenS9{F0: 1, F1: fi44}
                                                                    _ = v106
                                                                default:
                                                                    if ((v18 << ((65535) & 63)) >= v18) {
                                                                        var v107 int64 = 42
                                                                        fmt.Println(v45, v3)
                                                                        _ = v107
                                                                    }
                                                                }
                                                                var v108 int64 = 1
                                                                if true {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    var v109 string = "hello"
                                                                    v110 := GenS15{F0: 32767}
                                                                    _ = v109
                                                                    _ = v110
                                                                } else {
                                                                    if false {
                                                                        if false {
                                                                            var v111 uint = 65535
                                                                            var v112 int16 = 1000
                                                                            v111 += v111
                                                                            _ = v111
                                                                            _ = v112
                                                                        } else {
                                                                            var v113 uint32 = 2147483648
                                                                            var v114 uint = 0
                                                                            _ = v113
                                                                            _ = v114
                                                                        }
                                                                    }
                                                                    v45 = (^(v108))
                                                                }
                                                                v45 -= (v108 >> ((uint(v3)) & 63))
                                                                _ = v108
                                                            } else {
                                                                v3 -= (v2 << ((991) & 63))
                                                                fmt.Println(v2)
                                                            }
                                                        }
                                                        fmt.Println(v45, v2)
                                                        _ = v104
                                                    }
                                                    _ = fi44
                                                    _ = v45
                                                }
                                            case 1:
                                                for fi115 := 0; fi115 < 9; fi115++ {
                                                    var v116 int = -2147483648
                                                    v2 += 52
                                                    var v117 uint32 = v23
                                                    _ = fi115
                                                    _ = v116
                                                    _ = v117
                                                }
                                            case 97:
                                                fmt.Println(v19)
                                            default:
                                                var v118 uint32 = 65535
                                                _ = v118
                                            }
                                            var v119 uint32 = 906
                                            if len(v5) > 0 {
                                                v5[0] = (fi33 | (^(fi33)))
                                            }
                                            _ = v119
                                        }
                                        v23 -= v23
                                    }
                                }
                            } else {
                                switch "hello" {
                                case "hello":
                                    for fi120 := 0; fi120 < 1; fi120++ {
                                        var v121 int8 = 1
                                        fmt.Println(v19, v18, v23)
                                        _ = fi120
                                        _ = v121
                                    }
                                }
                                var v122 uint32 = uint32(v3)
                                var v123 float32 = 1.5
                                _ = v122
                                _ = v123
                            }
                        }
                    case -209:
                        for fi124 := 0; fi124 < 1; fi124++ {
                            if false {
                                var v125 int = fi124
                                v125 -= (v125 * genFn1(("\n\t" + "a"), -32768))
                                var v126 uint8 = v19
                                var v127 float64 = (3.14159 * 1.5)
                                _ = v125
                                _ = v126
                                _ = v127
                            }
                            var v128 int8 = -51
                            _ = fi124
                            _ = v128
                        }
                    }
                }
                v3 -= 1
                var v129 uint64 = 277
                _ = fi33
                _ = v129
            }
        }
    }
    var v130 int16 = -745
    switch 3 {
    case 4294967295:
        if (-626 == (v18 & (v18 - (v18 >> ((4294967295) & 63))))) {
            for fi131 := 0; fi131 < 16; fi131++ {
                v132 := GenS16{F0: 100}
                for fi133 := 0; fi133 < 10; fi133++ {
                    var v134 float64 = 1.5
                    v2 = (v3 * (v2 << ((4294967295) & 63)))
                    _ = fi133
                    _ = v134
                }
                _ = fi131
                _ = v132
            }
        } else {
            v23 += v23
            v135 := GenS17{F0: (v19 - uint8(v23)), F1: 42, F2: 1}
            if (!false) {
                var v136 uint64 = genFn0(503)
                if false {
                    var v137 int32 = v18
                    var v138 uint64 = (v136 & 4294967295)
                    _ = v137
                    _ = v138
                }
                _ = v136
            } else {
                v139 := GenS18{F0: 2, F1: 1, F2: -9223372036854775808}
                _ = v139
            }
            _ = v135
        }
    case -9223372036854775808:
        v140 := GenS9{F0: 255, F1: 442}
        _ = v140
    case 333:
        _ = len(v5)
    default:
        var v141 int16 = v130
        _ = v141
    }
    v19 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v130:", v130)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v130
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
