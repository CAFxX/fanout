package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 bool; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 int16; F2 string }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 string; F1 int32; F2 string }
type GenS9 struct { F0 uint; F1 int32; F2 int }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint64; F1 string }
type GenS12 struct { F0 uint8; F1 int64 }
type GenS13 struct { F0 uint; F1 int8 }
type GenS14 struct { F0 uint32; F1 uint32 }
type GenS15 struct { F0 uint16; F1 int64; F2 uint8 }
type GenS16 struct { F0 int; F1 bool }
type GenS17 struct { F0 uint; F1 string }
type GenS18 struct { F0 int8; F1 int64; F2 int64 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 uint32 }
type GenS22 struct { F0 uint64 }
type GenS23 struct { F0 uint; F1 uint }
type GenS24 struct { F0 uint16 }
func genFn0(p0 uint64) uint16 {
    var v5 uint = 255
    v6 := []int{4294967295, 0, -2147483648, 242}
    fmt.Println(v5, v5)
    p0 -= 18446744073709551615
    if false {
        v5 = 1
        _ = len(v6)
        var v7 int = 128
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return 890
}
func genFn1(p0 uint64, p1 int16, p2 uint16) bool {
    p0 = p0
    var v8 uint = 2
    for fi9 := 0; fi9 < 3; fi9++ {
        v8 = (v8 | v8)
        switch "" {
        case "a":
            if true {
                fmt.Println(v8)
            }
        case "\n\t":
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := []int{fi10}
                _ = fi10
                _ = v11
            }
        case "foo bar":
            v12 := []int{fi9, 32768, fi9}
            _ = v12
        default:
            var v13 uint64 = (p0 ^ (p0 / ((847) | 1)))
            _ = v13
        }
        var v14 float64 = (0.0 + (0.0 * 0.1))
        _ = fi9
        _ = v14
    }
    v15 := []int{-907}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v8
    return true
}

func main() {
    var v1 float64 = float64((0.0 + (0.1 - (0.0 * ((3.14159 - -2.25) + 3.14159)))))
    var v2 uint16 = 2
    v3 := []uint32{}
    var v4 bool = false
    switch -1 {
    case 42:
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    default:
        var v16 int16 = -1
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 uint16 = v2
        if len(v3) > 0 {
            v3[0] = 0
        }
        if false {
            fmt.Println(v2)
            for fi19 := 0; fi19 < 19; fi19++ {
                v2 += v18
                _ = fi19
            }
        } else {
            if false {
                if (42 < -9223372036854775808) {
                    switch ("a" + ("\n\t" + ("hello" + "a"))) {
                    case "\n\t":
                        if (v4 || (127 > 213)) {
                            for fi20 := 0; fi20 < 6; fi20++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(v4, v1)
                                _ = fi20
                            }
                        } else {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                }
                var v21 int16 = int16(v18)
                _ = v21
            } else {
                var v22 int32 = -2147483648
                fmt.Println(v22, v1)
                var v23 uint8 = 223
                _ = v22
                _ = v23
            }
            if (!v4) {
                switch (fi17 % ((fi17) | 1)) {
                case 32768:
                    for fi24 := 0; fi24 < 11; fi24++ {
                        v25 := GenS1{F0: 0, F1: 1}
                        _ = fi24
                        _ = v25
                    }
                case 127:
                    switch "" {
                    case "":
                        v2 -= (v2 ^ (^(v2)))
                    case "a":
                        v26 := GenS2{F0: 1}
                        _ = v26
                    }
                default:
                    for fi27 := 0; fi27 < 9; fi27++ {
                        switch fi27 {
                        case 4294967295:
                            v2 = 2
                        case -2147483648:
                            fmt.Println(fi17, fi17, fi17)
                        default:
                            var v28 int64 = 42
                            _ = v28
                        }
                        _ = fi27
                    }
                }
            }
            for fi29 := 0; fi29 < 9; fi29++ {
                if (v2 != v2) {
                    v30 := GenS3{F0: 147, F1: v4, F2: 255}
                    _ = v30
                } else {
                    v18 += (v2 << ((104) & 63))
                }
                _ = fi29
            }
        }
        _ = fi17
        _ = v18
    }
    if (127 < 42) {
        var v31 float32 = float32(v1)
        var v32 int32 = 42
        if (120 != -16) {
            v33 := GenS4{F0: 32768}
            v2 = genFn0(42)
            _ = v33
        }
        _ = v31
        _ = v32
    } else {
        fmt.Println(v2, v1, v1)
    }
    var v34 bool = true
    v35 := GenS4{F0: 348}
    if ((v4 || (!((v2 >> ((42) & 63)) != (v2 + genFn0(506))))) || v34) {
        for fi36 := 0; fi36 < 20; fi36++ {
            var v37 float32 = (0.1 / 1.0)
            fmt.Println(v34, v34, v34)
            _ = fi36
            _ = v37
        }
        v4 = (!false)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v38 bool = (0 > 42)
    switch -2147483648 {
    case 128:
        var v39 float32 = 0.1
        _ = v39
    default:
        var v40 int16 = 42
        _ = v40
    }
    switch "foo bar" {
    case "a":
        var v41 uint = 4294967295
        _ = v41
    case "foo bar":
        var v42 int = 9223372036854775807
        _ = v42
    }
    if v34 {
        v1 += (v1 * v1)
        for fi43 := 0; fi43 < 17; fi43++ {
            v4 = false
            for fi44 := 0; fi44 < 19; fi44++ {
                if v34 {
                    if (255 != 107) {
                        if v4 {
                            fmt.Println(v38, v1, fi43)
                            var v45 float64 = float64((((0.0 * 3.14159) + -2.25) + 0.1))
                            _ = v45
                        }
                        var v46 uint8 = 127
                        if v4 {
                            if (0.0 < (1.5 * (1e10 * (0.0 - 3.14159)))) {
                                if false {
                                    v46 = 128
                                } else {
                                    v47 := GenS4{F0: fi43}
                                    var v48 int32 = 2147483647
                                    _ = v47
                                    _ = v48
                                }
                                v46 = v46
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                v49 := GenS5{F0: 95, F1: -1, F2: "foo bar"}
                                switch (("a" + "a") + ((("a" + "hello") + "a") + "a")) {
                                case "foo bar":
                                    if (((3.14159 + ((1e10 - 0.0) - 1e10)) >= -2.25) || (9223372036854775807 >= -1)) {
                                        v50 := GenS6{F0: 18446744073709551615, F1: "\n\t"}
                                        _ = v50
                                    }
                                case "":
                                    switch fi43 {
                                    case 18:
                                        if (0 == 0) {
                                            v1 += 3.14159
                                            if true {
                                                var v51 float64 = v1
                                                if v34 {
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v52 uint = 255
                                                        _ = v52
                                                    default:
                                                        var v53 uint32 = 42
                                                        _ = v53
                                                    }
                                                    var v54 int = int(v46)
                                                    _ = v54
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                v55 := GenS7{F0: fi43, F1: ((((("a" + "") + "") + "") + "a") + "hello")}
                                                _ = v51
                                                _ = v55
                                            }
                                            fmt.Println(v4, v34, fi43)
                                        }
                                    case 1:
                                        if (false || (1 >= 2)) {
                                            _ = len(v3)
                                            var v56 float32 = 0.1
                                            v2 -= (v2 | 132)
                                            _ = v56
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = 2147483647
                                            }
                                            if v34 {
                                                var v57 int = (-(fi43))
                                                fmt.Println(v57, fi44)
                                                var v58 int32 = 42
                                                _ = v57
                                                _ = v58
                                            } else {
                                                var v59 uint32 = 145
                                                v1 -= float64(3.14159)
                                                _ = len(v3)
                                                _ = v59
                                            }
                                            fmt.Println(v34, fi43, fi44)
                                        }
                                    case -9223372036854775808:
                                        v3 = append(v3, 1)
                                    default:
                                        if (true && true) {
                                            if (v4 && v4) {
                                                var v60 uint8 = v46
                                                v61 := GenS4{F0: 4294967295}
                                                v46 += (v60 % ((v46) | 1))
                                                _ = v60
                                                _ = v61
                                            } else {
                                                v4 = v34
                                            }
                                            fmt.Println(v1, fi44, fi43)
                                        } else {
                                            switch ("" + ("a" + "foo bar")) {
                                            case "a":
                                                v34 = v4
                                                var v62 int16 = -32768
                                                var v63 int32 = 393
                                                _ = v62
                                                _ = v63
                                            case "foo bar":
                                                _ = len(v3)
                                            case "":
                                                v3 = append(v3, 2147483648)
                                            }
                                        }
                                    }
                                case "hello":
                                    switch (fi44 / ((127) | 1)) {
                                    case 4294967295:
                                        v3 = append(v3, 452)
                                    case -643:
                                        var v64 int = fi44
                                        _ = v64
                                    }
                                default:
                                    v1 -= float64(((0.0 / (0.0 + ((float32((v1 - (v1 - v1))) * ((0.1 * float32((-2.25 * (1e10 / float64(1e10))))) + 3.14159)) + 1e10))) / (0.0 - float32(v1))))
                                }
                                fmt.Println(v2, fi43, v34)
                                if v34 {
                                    switch "hello" {
                                    case "hello":
                                        var v65 int8 = -128
                                        _ = v65
                                    case "":
                                        var v66 uint16 = 42
                                        _ = v66
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = 0
                                        }
                                        if v38 {
                                            var v67 uint64 = 627
                                            v68 := GenS8{F0: ("hello" + "a"), F1: -30, F2: ("\n\t" + "hello")}
                                            v69 := GenS3{F0: 9223372036854775807, F1: true, F2: 127}
                                            _ = v67
                                            _ = v68
                                            _ = v69
                                        }
                                        if v38 {
                                            var v70 uint8 = v46
                                            var v71 uint16 = 32767
                                            _ = v70
                                            _ = v71
                                        } else {
                                            switch "" {
                                            case "hello":
                                                var v72 int32 = -141
                                                _ = v72
                                            }
                                            v4 = ((1.5 - (0.1 + 3.14159)) < 1e10)
                                        }
                                    }
                                    var v73 bool = (504 <= 922)
                                    if v34 {
                                        switch "foo bar" {
                                        case "":
                                            if ((-591 <= 0) && v4) {
                                                if v73 {
                                                    var v74 uint16 = v2
                                                    if true {
                                                        var v75 bool = (190 > 18446744073709551615)
                                                        switch 0 {
                                                        case -198:
                                                            switch ((("a" + "foo bar") + "a") + "a") {
                                                            case "a":
                                                                v76 := GenS9{F0: 255, F1: 2147483647, F2: (^(fi44))}
                                                                fmt.Println(v2, v73)
                                                                v77 := GenS1{F0: (v46 % ((v46) | 1)), F1: (v46 & (v46 << ((4294967295) & 63)))}
                                                                _ = v76
                                                                _ = v77
                                                            case "\n\t":
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            default:
                                                                _ = len(v3)
                                                            }
                                                        case -9223372036854775808:
                                                            v3 = append(v3, 1)
                                                        case 0:
                                                            if (v34 && (!(65535 >= 42))) {
                                                                switch int(v2) {
                                                                case 127:
                                                                    var v78 uint32 = 4294967295
                                                                    _ = v78
                                                                case -9223372036854775808:
                                                                    v46 += uint8(v74)
                                                                default:
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                                if true {
                                                                    _ = len(v3)
                                                                }
                                                                switch (-(fi43)) {
                                                                case -1:
                                                                    v2 += (v74 * (v74 >> ((255) & 63)))
                                                                case -961:
                                                                    v34 = v75
                                                                    var v79 int32 = 42
                                                                    if v73 {
                                                                        if false {
                                                                            var v80 bool = v4
                                                                            if false {
                                                                                var v81 string = "hello"
                                                                                v34 = ((v80 && v4) || (((2 > 2) || ((1e10 - 3.14159) < 1e10)) && v4))
                                                                                _ = v81
                                                                            } else {
                                                                                v2 -= (v2 + (v2 & (v74 + 489)))
                                                                            }
                                                                            _ = v80
                                                                        } else {
                                                                            switch (("" + ((("a" + "") + ("a" + ("foo bar" + (("\n\t" + "a") + "")))) + "\n\t")) + ("" + ("hello" + ("" + "")))) {
                                                                            case "hello":
                                                                                if (0 == 489) {
                                                                                    if (genFn0(2) > v74) {
                                                                                        fmt.Println(fi44, v79, fi43)
                                                                                        var v82 uint32 = 1
                                                                                        fmt.Println(v1)
                                                                                        _ = v82
                                                                                    } else {
                                                                                        v83 := GenS6{F0: 1, F1: "a"}
                                                                                        _ = v83
                                                                                    }
                                                                                    var v84 bool = true
                                                                                    fmt.Println(v1, v1, v75)
                                                                                    _ = v84
                                                                                }
                                                                            case "":
                                                                                var v85 int16 = 1
                                                                                _ = v85
                                                                            case "foo bar":
                                                                                v4 = (-128 >= 1)
                                                                            default:
                                                                                var v86 uint32 = 756
                                                                                _ = v86
                                                                            }
                                                                            fmt.Println(v1)
                                                                        }
                                                                        v34 = (v1 >= float64(1.5))
                                                                        switch (fi44 << ((uint(v74)) & 63)) {
                                                                        case 128:
                                                                            var v87 bool = false
                                                                            _ = v87
                                                                        case 32767:
                                                                            v74 = v2
                                                                        default:
                                                                            if false {
                                                                                fmt.Println(v1, v1, v79)
                                                                                fmt.Println(v46)
                                                                                if v38 {
                                                                                    var v88 float32 = -2.25
                                                                                    if (1 < -128) {
                                                                                        var v89 int16 = 0
                                                                                        if (v1 > (0.0 / v1)) {
                                                                                            v90 := GenS9{F0: 177, F1: v79, F2: (fi44 ^ fi43)}
                                                                                            fmt.Println(v75, v4)
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            _ = v90
                                                                                        } else {
                                                                                            fmt.Println(v75, v1)
                                                                                        }
                                                                                        v3 = append(v3, 818)
                                                                                        _ = v89
                                                                                    } else {
                                                                                        var v91 uint = 234
                                                                                        v4 = (false || true)
                                                                                        var v92 string = "\n\t"
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                    }
                                                                                    v93 := GenS8{F0: ("foo bar" + ("a" + "hello")), F1: (v79 >> ((18446744073709551615) & 63)), F2: "a"}
                                                                                    _ = v88
                                                                                    _ = v93
                                                                                } else {
                                                                                    var v94 int16 = 1000
                                                                                    switch 0 {
                                                                                    case -709:
                                                                                        switch ("" + "foo bar") {
                                                                                        case "\n\t":
                                                                                            if ((^(v79)) <= (v79 * int32(v46))) {
                                                                                                if v4 {
                                                                                                    if false {
                                                                                                        var v95 bool = (-128 < 16)
                                                                                                        if (!(v94 >= (v94 - v94))) {
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = 2147483647
                                                                                                            }
                                                                                                            if (9223372036854775807 == 1) {
                                                                                                                v46 += 42
                                                                                                                v38 = v75
                                                                                                                if v34 {
                                                                                                                    v79 = v79
                                                                                                                } else {
                                                                                                                    fmt.Println(v38, fi44, v73)
                                                                                                                    if (1 == (v79 ^ (v79 >> ((65535) & 63)))) {
                                                                                                                        if (v4 || false) {
                                                                                                                            if (-128 < 100) {
                                                                                                                                v96 := GenS10{F0: 100}
                                                                                                                                fmt.Println(v74)
                                                                                                                                _ = v96
                                                                                                                            } else {
                                                                                                                                var v97 bool = v95
                                                                                                                                v98 := GenS6{F0: 636, F1: "a"}
                                                                                                                                if (v34 && ((v2 / (((v74 << ((910) & 63))) | 1)) > (v2 + genFn0(0)))) {
                                                                                                                                    switch "" {
                                                                                                                                    case "foo bar":
                                                                                                                                        _ = len(v3)
                                                                                                                                    case "":
                                                                                                                                        var v99 uint64 = 228
                                                                                                                                        _ = v99
                                                                                                                                    default:
                                                                                                                                        switch ("" + ("foo bar" + ("hello" + "a"))) {
                                                                                                                                        case "a":
                                                                                                                                            v100 := GenS11{F0: 326, F1: "a"}
                                                                                                                                            fmt.Println(v94)
                                                                                                                                            _ = v100
                                                                                                                                        case "hello":
                                                                                                                                            v38 = ((v75 && false) && v75)
                                                                                                                                        case "foo bar":
                                                                                                                                            var v101 int16 = -560
                                                                                                                                            _ = v101
                                                                                                                                        default:
                                                                                                                                            var v102 int64 = 9223372036854775807
                                                                                                                                            _ = v102
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    if len(v3) >= 2 {
                                                                                                                                        v3 = v3[1:len(v3)]
                                                                                                                                    }
                                                                                                                                    var v103 uint16 = v2
                                                                                                                                    _ = v103
                                                                                                                                }
                                                                                                                                _ = len(v3)
                                                                                                                                v104 := GenS10{F0: -67}
                                                                                                                                _ = v104
                                                                                                                                _ = v97
                                                                                                                                _ = v98
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                v105 := GenS12{F0: (v46 * 255), F1: -181}
                                                                                                                                var v106 uint16 = genFn0(216)
                                                                                                                                v34 = true
                                                                                                                                _ = v105
                                                                                                                                _ = v106
                                                                                                                            }
                                                                                                                            var v107 string = ("a" + ("" + "hello"))
                                                                                                                            _ = v107
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            _ = len(v3)
                                                                                                        }
                                                                                                        _ = v95
                                                                                                    }
                                                                                                    if (!((-934 != v94) || (((v2 << ((1) & 63)) > (v74 << ((255) & 63))) || (!(1 <= 1))))) {
                                                                                                        if len(v3) > 0 {
                                                                                                            v3[0] = 173
                                                                                                        }
                                                                                                        v74 = v2
                                                                                                        if ((v4 && v75) || (!true)) {
                                                                                                            if (1 <= int8(v2)) {
                                                                                                                if (!(v79 >= v79)) {
                                                                                                                    var v108 string = ""
                                                                                                                    if v34 {
                                                                                                                        var v109 int32 = v79
                                                                                                                        v75 = v4
                                                                                                                        var v110 uint16 = (^(v74))
                                                                                                                        _ = v109
                                                                                                                        _ = v110
                                                                                                                    }
                                                                                                                    if (v46 != (v46 ^ v46)) {
                                                                                                                        v111 := GenS13{F0: 4294967295, F1: 1}
                                                                                                                        v3 = append(v3, 2)
                                                                                                                        _ = v111
                                                                                                                    }
                                                                                                                    _ = v108
                                                                                                                }
                                                                                                                v74 -= (v2 * v74)
                                                                                                            } else {
                                                                                                                fmt.Println(v1)
                                                                                                                v38 = false
                                                                                                            }
                                                                                                            var v112 int8 = -90
                                                                                                            if v4 {
                                                                                                                v75 = (false || (v73 || false))
                                                                                                            } else {
                                                                                                                var v113 uint16 = v74
                                                                                                                var v114 int16 = (-(v94))
                                                                                                                _ = v113
                                                                                                                _ = v114
                                                                                                            }
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = 4294967295
                                                                                                            }
                                                                                                            var v115 uint8 = v46
                                                                                                            _ = v112
                                                                                                            _ = v115
                                                                                                        } else {
                                                                                                            if v38 {
                                                                                                                var v116 int8 = 1
                                                                                                                _ = v116
                                                                                                            } else {
                                                                                                                var v117 int8 = -128
                                                                                                                switch "a" {
                                                                                                                case "foo bar":
                                                                                                                    if v38 {
                                                                                                                        var v118 uint = 2
                                                                                                                        _ = v118
                                                                                                                    } else {
                                                                                                                        v3 = append(v3, 65535)
                                                                                                                        v46 = (^(v46))
                                                                                                                    }
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v74, fi43, fi43)
                                                                                                                }
                                                                                                                fmt.Println(fi43)
                                                                                                                _ = v117
                                                                                                            }
                                                                                                            if true {
                                                                                                                var v119 uint = 65535
                                                                                                                var v120 int64 = 42
                                                                                                                _ = v119
                                                                                                                _ = v120
                                                                                                            }
                                                                                                            if len(v3) > 0 {
                                                                                                                v3[0] = 42
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        _ = len(v3)
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v121 int16 = v94
                                                                                                _ = v121
                                                                                            }
                                                                                        case "foo bar":
                                                                                            if (4294967295 <= 1) {
                                                                                                fmt.Println(v46, v2)
                                                                                                var v122 string = "a"
                                                                                                var v123 uint8 = (v46 << ((1) & 63))
                                                                                                _ = v122
                                                                                                _ = v123
                                                                                            }
                                                                                        case "":
                                                                                            fmt.Println(v34)
                                                                                        default:
                                                                                            var v124 uint32 = 42
                                                                                            v125 := GenS14{F0: v124, F1: 0}
                                                                                            _ = v124
                                                                                            _ = v125
                                                                                        }
                                                                                    case 9223372036854775807:
                                                                                        var v126 int32 = -2147483648
                                                                                        _ = v126
                                                                                    }
                                                                                    _ = v94
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        _ = len(v3)
                                                                        var v127 int32 = 340
                                                                        _ = v127
                                                                    }
                                                                    _ = v79
                                                                default:
                                                                    var v128 int64 = 9223372036854775807
                                                                    _ = v128
                                                                }
                                                            } else {
                                                                if ((true && v4) && v4) {
                                                                    v129 := GenS10{F0: 1}
                                                                    _ = v129
                                                                }
                                                                v4 = (!false)
                                                            }
                                                            if (true || ((fi43 == 1) && (v1 < (v1 * 1e10)))) {
                                                                fmt.Println(fi43, v74)
                                                                v130 := GenS15{F0: v74, F1: -855, F2: v46}
                                                                var v131 uint32 = 588
                                                                if (961 == -1) {
                                                                    var v132 float32 = ((0.0 - 1.5) / float32(v1))
                                                                    v133 := GenS16{F0: 32767, F1: (fi44 < fi44)}
                                                                    var v134 uint64 = 0
                                                                    _ = v132
                                                                    _ = v133
                                                                    _ = v134
                                                                } else {
                                                                    v2 = (v74 - genFn0(2))
                                                                }
                                                                _ = v130
                                                                _ = v131
                                                            } else {
                                                                var v135 int16 = -89
                                                                if (v46 > (v46 | v46)) {
                                                                    v136 := GenS13{F0: 1, F1: 100}
                                                                    var v137 float64 = (v1 + v1)
                                                                    _ = v136
                                                                    _ = v137
                                                                } else {
                                                                    v138 := GenS17{F0: 65535, F1: "a"}
                                                                    _ = v138
                                                                }
                                                                _ = v135
                                                            }
                                                            v2 = (v2 * (v2 - v74))
                                                            switch "a" {
                                                            case "a":
                                                                v74 -= v2
                                                            }
                                                            v46 += (v46 % ((127) | 1))
                                                        }
                                                        var v139 int = 110
                                                        _ = v139
                                                        _ = v75
                                                    }
                                                    _ = v74
                                                } else {
                                                    var v140 uint64 = 2
                                                    fmt.Println(v73, v140)
                                                    if v38 {
                                                        switch "" {
                                                        case "\n\t":
                                                            v140 = (v140 / ((2) | 1))
                                                        default:
                                                            if ((!v73) && v34) {
                                                                var v141 uint = 65535
                                                                _ = v141
                                                            }
                                                        }
                                                        if v38 {
                                                            var v142 uint64 = 4294967295
                                                            _ = v142
                                                        }
                                                    }
                                                    _ = v140
                                                }
                                                v2 = (v2 >> ((881) & 63))
                                                v143 := GenS18{F0: 1, F1: -9223372036854775808, F2: -1}
                                                _ = v143
                                            } else {
                                                var v144 uint32 = 1
                                                v1 += 1.5
                                                switch (("foo bar" + "") + "") {
                                                case "a":
                                                    var v145 uint64 = 720
                                                    _ = v145
                                                case "\n\t":
                                                    v146 := GenS2{F0: 0}
                                                    _ = v146
                                                default:
                                                    fmt.Println(v46, v2, v46)
                                                }
                                                _ = v144
                                            }
                                        default:
                                            v1 -= ((-2.25 + (v1 * v1)) + 3.14159)
                                        }
                                    }
                                    _ = v73
                                }
                                if (true || true) {
                                    v34 = false
                                    if (((v1 - (float64(3.14159) - float64((1e10 * 3.14159)))) != v1) && true) {
                                        fmt.Println(v46, v34, v34)
                                        switch 4294967295 {
                                        case 9223372036854775807:
                                            var v147 string = (("hello" + "foo bar") + "\n\t")
                                            _ = v147
                                        case -1:
                                            var v148 int16 = -32768
                                            _ = v148
                                        case -9223372036854775808:
                                            switch (fi43 - (fi43 % (((fi44 / ((fi44) | 1))) | 1))) {
                                            case 2147483647:
                                                if false {
                                                    var v149 int16 = -81
                                                    var v150 int8 = -34
                                                    _ = v149
                                                    _ = v150
                                                }
                                            case -523:
                                                v151 := GenS11{F0: 2, F1: "foo bar"}
                                                _ = v151
                                            }
                                        default:
                                            switch "" {
                                            case "hello":
                                                var v152 int = (fi43 & 669)
                                                _ = v152
                                            case "a":
                                                v153 := GenS19{F0: (v2 ^ genFn0(1))}
                                                _ = v153
                                            default:
                                                v154 := GenS20{F0: 2}
                                                _ = v154
                                            }
                                        }
                                    }
                                }
                                v155 := GenS21{F0: 934}
                                if ((v46 % ((43) | 1)) > 161) {
                                    v1 = 1.5
                                } else {
                                    v38 = (!false)
                                    if (18446744073709551615 != 18446744073709551615) {
                                        switch "hello" {
                                        case "a":
                                            var v156 uint = 682
                                            _ = v156
                                        case "":
                                            switch 4294967295 {
                                            case -1:
                                                if (-1 < -9223372036854775808) {
                                                    var v157 int16 = 1000
                                                    var v158 bool = (!(false || v38))
                                                    _ = v157
                                                    _ = v158
                                                }
                                            case 868:
                                                v159 := GenS22{F0: 204}
                                                _ = v159
                                            case 128:
                                                if len(v3) > 0 {
                                                    v3[0] = 625
                                                }
                                            }
                                        }
                                        if v34 {
                                            var v160 uint64 = 919
                                            var v161 bool = (0 <= 255)
                                            if len(v3) > 0 {
                                                v3[0] = 2
                                            }
                                            if ((fi44 + (^(fi44))) > 297) {
                                                v162 := GenS14{F0: 1, F1: 2147483648}
                                                fmt.Println(v4, v34, v46)
                                                _ = v162
                                            } else {
                                                if (!((127 > -44) || (1 == 1000))) {
                                                    var v163 string = "\n\t"
                                                    v34 = ((genFn0(42) >= 65535) && false)
                                                    _ = v163
                                                }
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch (fi43 * fi43) {
                                                    case 42:
                                                        v38 = true
                                                    case 380:
                                                        var v164 float64 = v1
                                                        _ = v164
                                                    case 4294967295:
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    }
                                                case "a":
                                                    fmt.Println(v4, v46)
                                                }
                                            }
                                            switch ("a" + "") {
                                            case "\n\t":
                                                v1 += ((1.5 * v1) / (-2.25 - 3.14159))
                                            case "a":
                                                var v165 int = (fi44 & fi43)
                                                _ = v165
                                            case "hello":
                                                var v166 uint16 = v2
                                                _ = v166
                                            default:
                                                var v167 float64 = ((v1 + ((v1 - (1e10 + v1)) + ((float64(3.14159) / float64(1.5)) - -2.25))) / ((0.0 - (v1 - float64(3.14159))) - (v1 + 1e10)))
                                                var v168 int64 = int64(v46)
                                                v161 = (!v161)
                                                _ = v167
                                                _ = v168
                                            }
                                            _ = v160
                                            _ = v161
                                        }
                                        v2 = 2
                                    }
                                    v169 := GenS15{F0: 29, F1: 586, F2: 127}
                                    _ = v169
                                }
                                _ = v155
                                _ = v49
                            }
                        }
                        _ = v46
                    }
                    fmt.Println(v34)
                }
                _ = fi44
            }
            _ = fi43
        }
    } else {
        switch -1 {
        case -1:
            var v170 uint32 = 25
            _ = v170
        }
        var v171 bool = ((v4 || (128 != 255)) || v34)
        switch ("a" + "a") {
        case "foo bar":
            if false {
                fmt.Println(v34, v171)
                var v172 int16 = 32767
                fmt.Println(v2)
                _ = v172
            }
        case "":
            var v173 uint64 = 0
            _ = v173
        case "a":
            v2 += (v2 / ((42) | 1))
        default:
            for fi174 := 0; fi174 < 18; fi174++ {
                var v175 float32 = 0.0
                v4 = v171
                _ = fi174
                _ = v175
            }
        }
        _ = v171
    }
    v38 = (!(v38 && v4))
    fmt.Println(v34, v38, v34)
    switch "foo bar" {
    case "a":
        v176 := GenS23{F0: 822, F1: 18446744073709551615}
        _ = v176
    case "\n\t":
        v177 := GenS24{F0: (v2 % (((v2 ^ 255)) | 1))}
        _ = v177
    case "foo bar":
        var v178 string = "foo bar"
        _ = v178
    }
    v3 = append(v3, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v4
}
