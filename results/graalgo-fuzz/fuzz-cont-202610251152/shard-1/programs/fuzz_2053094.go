package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 uint16; F1 string; F2 bool }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 int16 }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 bool; F1 string }
type GenS11 struct { F0 uint64; F1 bool; F2 int16 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int; F1 uint8 }
type GenS15 struct { F0 int; F1 int64 }
type GenS16 struct { F0 string; F1 int32; F2 uint }
type GenS17 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS18 struct { F0 bool; F1 uint8 }
type GenS19 struct { F0 uint32; F1 uint32 }
type GenS20 struct { F0 int32; F1 int8 }
type GenS21 struct { F0 string; F1 uint }
type GenS22 struct { F0 bool; F1 bool; F2 uint32 }
type GenS23 struct { F0 int16; F1 int64 }
type GenS24 struct { F0 uint }
type GenS25 struct { F0 int; F1 uint8; F2 uint64 }
type GenS26 struct { F0 uint16; F1 uint8 }
type GenS27 struct { F0 string }
type GenS28 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS29 struct { F0 uint16 }
type GenS30 struct { F0 int }
func genFn0(p0 bool, p1 uint16, p2 uint8) int8 {
    var v4 uint8 = (^(p2))
    v5 := GenS1{F0: 1, F1: "hello"}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -1
}

func main() {
    v1 := []int{128, 402}
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    if false {
        if (-9223372036854775808 >= -1) {
            v6 := v2[42]
            var v7 uint8 = 2
            switch 0 {
            case 13:
                v6 = (-(v6))
            default:
                v7 = (v7 / (((^(v7))) | 1))
            }
            _ = v6
            _ = v7
        }
        switch "a" {
        case "\n\t":
            var v8 float64 = 1.5
            _ = v8
        case "":
            if (1 >= 90) {
                for fi9 := 0; fi9 < 5; fi9++ {
                    if true {
                        var v10 int64 = -1
                        v11 := GenS1{F0: (v10 % ((v10) | 1)), F1: "foo bar"}
                        var v12 uint32 = 2147483647
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 17; fi13++ {
                        var v14 bool = false
                        fmt.Println(fi9, fi13)
                        v1 = append(v1, (fi9 * (^(fi9))))
                        v15 := GenS2{F0: (351 >= 0), F1: genFn0(false, 758, 245)}
                        switch (("a" + "a") + ((((("" + (("" + ("foo bar" + "foo bar")) + "a")) + "") + (("" + "") + "\n\t")) + "foo bar") + "\n\t")) {
                        case "":
                            var v16 int64 = int64(fi9)
                            _ = v16
                        case "\n\t":
                            var v17 uint = 2
                            _ = v17
                        case "hello":
                            v14 = v14
                            var v18 int32 = 42
                            v14 = v14
                            _ = v18
                        default:
                            fmt.Println(fi13, fi9, v14)
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    var v19 float64 = (0.1 - 0.1)
                    _ = fi9
                    _ = v19
                }
                for fi20 := 0; fi20 < 11; fi20++ {
                    switch "hello" {
                    case "a":
                        _ = len(v1)
                    default:
                        var v21 float32 = 0.1
                        _ = v21
                    }
                    _ = fi20
                }
            } else {
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 uint = 2
                    v23 += v23
                    fmt.Println(fi22, fi22)
                    _ = fi22
                    _ = v23
                }
                var v24 uint64 = 1
                _ = v24
            }
        default:
            if (("foo bar" + ("" + "a")) >= (("hello" + "") + "foo bar")) {
                v25 := GenS1{F0: 9223372036854775807, F1: ("hello" + "a")}
                switch -715 {
                case -1:
                    var v26 int8 = 1
                    _ = v26
                }
                v27 := GenS3{F0: (false && (127 < -661)), F1: 204}
                _ = v25
                _ = v27
            } else {
                var v28 uint = 255
                for fi29 := 0; fi29 < 8; fi29++ {
                    if ((1 == 94) && (1 > 110)) {
                        var v30 float32 = ((0.0 * 1.5) + float32(3.14159))
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 11; fi31++ {
                        var v32 int16 = -32768
                        _ = fi31
                        _ = v32
                    }
                    _ = fi29
                }
                switch "a" {
                case "foo bar":
                    v28 += (v28 & 615)
                case "a":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "\n\t":
                    if true {
                        var v33 uint = (^(v28))
                        _ = v33
                    } else {
                        var v34 int32 = 0
                        for fi35 := 0; fi35 < 19; fi35++ {
                            if false {
                                if false {
                                    for fi36 := 0; fi36 < 12; fi36++ {
                                        var v37 uint64 = 1
                                        fmt.Println(fi36, v37, v34)
                                        switch fi35 {
                                        case 0:
                                            var v38 float32 = 1.5
                                            _ = v38
                                        case 32767:
                                            if (!true) {
                                                switch "" {
                                                case "a":
                                                    v39 := GenS1{F0: 467, F1: "hello"}
                                                    _ = v39
                                                case "\n\t":
                                                    switch (fi35 - (fi35 | fi35)) {
                                                    case -1:
                                                        switch ("a" + "\n\t") {
                                                        case "":
                                                            var v40 int32 = (v34 | (v34 + -1))
                                                            _ = v40
                                                        case "hello":
                                                            switch fi36 {
                                                            case 908:
                                                                var v41 bool = ((v28 * 226) != (v28 & 959))
                                                                _ = v41
                                                            case -621:
                                                                v42 := GenS3{F0: false, F1: 1}
                                                                var v43 bool = (true || false)
                                                                if true {
                                                                    v43 = v43
                                                                    if false {
                                                                        v44 := GenS2{F0: (4294967295 < (fi36 * (fi36 % (((fi35 | 2147483647)) | 1)))), F1: 127}
                                                                        _ = v44
                                                                    } else {
                                                                        switch (("hello" + (((("a" + ("\n\t" + ("" + ("foo bar" + ("" + ""))))) + "a") + "a") + "hello")) + "hello") {
                                                                        case "hello":
                                                                            switch fi35 {
                                                                            case 2147483647:
                                                                                var v45 string = ""
                                                                                _ = v45
                                                                            default:
                                                                                if false {
                                                                                    delete(v3, fi36)
                                                                                    v46 := GenS2{F0: (96 == 2), F1: -98}
                                                                                    _ = v46
                                                                                }
                                                                            }
                                                                            if v43 {
                                                                                v1 = append(v1, fi35)
                                                                                var v47 bool = (112 <= 209)
                                                                                _ = len(v1)
                                                                                _ = v47
                                                                            }
                                                                        case "a":
                                                                            v43 = ((v28 ^ v28) >= v28)
                                                                        }
                                                                        if v43 {
                                                                            switch (fi36 ^ (fi35 >> (((v28 / ((0) | 1))) & 63))) {
                                                                            case 42:
                                                                                fmt.Println(fi36)
                                                                            case 32768:
                                                                                var v48 int16 = -379
                                                                                _ = v48
                                                                            }
                                                                            if (519 == 418) {
                                                                                if (127 != 4) {
                                                                                    v43 = ("hello" > "foo bar")
                                                                                }
                                                                                v37 = 471
                                                                                v34 = (-(v34))
                                                                            }
                                                                        }
                                                                    }
                                                                    v43 = v43
                                                                } else {
                                                                    if ((^(fi35)) >= fi36) {
                                                                        v49 := GenS3{F0: v43, F1: 27}
                                                                        v43 = (float32((-2.25 * 1e10)) == 1.5)
                                                                        _ = v49
                                                                    } else {
                                                                        var v50 string = "a"
                                                                        var v51 uint = v28
                                                                        var v52 int64 = 913
                                                                        switch fi35 {
                                                                        case 64:
                                                                            var v53 int32 = -2147483648
                                                                            _ = v53
                                                                        case 32768:
                                                                            switch ((v50 + (v50 + v50)) + ("hello" + ((v50 + v50) + v50))) {
                                                                            case "foo bar":
                                                                                switch "a" {
                                                                                case "":
                                                                                    if false {
                                                                                        if true {
                                                                                            v54 := v2[2]
                                                                                            var v55 float64 = (0.0 - (3.14159 * 3.14159))
                                                                                            if (float64((3.14159 - 0.1)) != v55) {
                                                                                                var v56 int = 32767
                                                                                                switch "a" {
                                                                                                case "foo bar":
                                                                                                    v51 += v28
                                                                                                case "":
                                                                                                    fmt.Println(v54)
                                                                                                default:
                                                                                                    switch ((v50 + v50) + ("hello" + (v50 + v50))) {
                                                                                                    case "\n\t":
                                                                                                        switch (fi35 + (fi35 >> (((^(v28))) & 63))) {
                                                                                                        case 127:
                                                                                                            if v43 {
                                                                                                                var v57 uint16 = 1
                                                                                                                _ = v57
                                                                                                            } else {
                                                                                                                var v58 bool = (2 < 2147483648)
                                                                                                                var v59 bool = ((v37 + (v37 & v37)) == v37)
                                                                                                                if v43 {
                                                                                                                    if ((v56 % (((fi35 * (v56 / ((127) | 1)))) | 1)) != fi35) {
                                                                                                                        delete(v3, int(v34))
                                                                                                                        fmt.Println(v34)
                                                                                                                    } else {
                                                                                                                        v60 := GenS1{F0: v52, F1: (v50 + v50)}
                                                                                                                        if v59 {
                                                                                                                            var v61 uint8 = 1
                                                                                                                            if len(v1) > 0 {
                                                                                                                                v1[0] = (fi36 + (fi35 % ((fi35) | 1)))
                                                                                                                            }
                                                                                                                            if ((v56 * fi35) > fi35) {
                                                                                                                                v58 = false
                                                                                                                                var v62 float64 = ((0.1 / 1.0) - v55)
                                                                                                                                var v63 uint16 = 306
                                                                                                                                _ = v62
                                                                                                                                _ = v63
                                                                                                                            }
                                                                                                                            _ = v61
                                                                                                                        } else {
                                                                                                                            if ((v34 >> ((v28) & 63)) != v34) {
                                                                                                                                var v64 uint16 = 42
                                                                                                                                v56 -= fi36
                                                                                                                                v65 := GenS2{F0: v59, F1: -128}
                                                                                                                                _ = v64
                                                                                                                                _ = v65
                                                                                                                            } else {
                                                                                                                                switch "" {
                                                                                                                                case "hello":
                                                                                                                                    v28 += v51
                                                                                                                                    var v66 float32 = -2.25
                                                                                                                                    _ = len(v1)
                                                                                                                                    _ = v66
                                                                                                                                }
                                                                                                                                if ((v37 / ((905) | 1)) >= 4294967295) {
                                                                                                                                    v56 = v54
                                                                                                                                    v34 += v34
                                                                                                                                    switch (fi36 << ((42) & 63)) {
                                                                                                                                    case 128:
                                                                                                                                        v56 = (^(fi36))
                                                                                                                                    default:
                                                                                                                                        v1 = append(v1, v54)
                                                                                                                                    }
                                                                                                                                    if v58 {
                                                                                                                                        switch fi36 {
                                                                                                                                        case -2147483648:
                                                                                                                                            v67 := GenS4{F0: uint8(v28), F1: (!(v43 && v59))}
                                                                                                                                            _ = v67
                                                                                                                                        case -9223372036854775808:
                                                                                                                                            switch "foo bar" {
                                                                                                                                            case "":
                                                                                                                                                delete(v2, 797)
                                                                                                                                            case "foo bar":
                                                                                                                                                var v68 int32 = v34
                                                                                                                                                _ = v68
                                                                                                                                            case "a":
                                                                                                                                                v69 := GenS2{F0: (!v59), F1: genFn0(((v37 * 818) == (^(v37))), 1, 1)}
                                                                                                                                                _ = v69
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            v70 := v3[32767]
                                                                                                                                            _ = v70
                                                                                                                                        }
                                                                                                                                        v71 := GenS4{F0: 128, F1: v43}
                                                                                                                                        v50 = "a"
                                                                                                                                        _ = v71
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(v56, v56, v55)
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    var v72 uint64 = 0
                                                                                                                                    _ = v72
                                                                                                                                }
                                                                                                                                if len(v1) >= 2 {
                                                                                                                                    v1 = v1[1:len(v1)]
                                                                                                                                }
                                                                                                                            }
                                                                                                                            v73 := GenS2{F0: (float32(float64((1.5 - (3.14159 - 1e10)))) <= 3.14159), F1: genFn0(true, 42, 20)}
                                                                                                                            var v74 int64 = (v52 / ((v52) | 1))
                                                                                                                            _ = v73
                                                                                                                            _ = v74
                                                                                                                        }
                                                                                                                        if (!true) {
                                                                                                                            v54 += fi35
                                                                                                                            var v75 uint32 = 2
                                                                                                                            _ = v75
                                                                                                                        }
                                                                                                                        _ = v60
                                                                                                                    }
                                                                                                                    if (36 > 127) {
                                                                                                                        var v76 uint16 = 255
                                                                                                                        fmt.Println(v54, v76)
                                                                                                                        if (!true) {
                                                                                                                            if (2 == (^(v28))) {
                                                                                                                                fmt.Println(v43, fi36, v55)
                                                                                                                                var v77 uint64 = 2
                                                                                                                                _ = v77
                                                                                                                            } else {
                                                                                                                                fmt.Println(v51, v56)
                                                                                                                                v58 = false
                                                                                                                                var v78 string = "foo bar"
                                                                                                                                _ = v78
                                                                                                                            }
                                                                                                                            v76 += (v76 << ((v28) & 63))
                                                                                                                            if v43 {
                                                                                                                                if len(v1) > 0 {
                                                                                                                                    v1[0] = 1
                                                                                                                                }
                                                                                                                                fmt.Println(v52, fi35, v43)
                                                                                                                            }
                                                                                                                            _ = len(v1)
                                                                                                                            var v79 int16 = int16(v28)
                                                                                                                            _ = v79
                                                                                                                        } else {
                                                                                                                            if v59 {
                                                                                                                                v80 := GenS5{F0: v76, F1: (("hello" + ((v50 + v50) + (v50 + v50))) + ("a" + ((v50 + ("\n\t" + v50)) + v50))), F2: ((v37 | v37) < uint64(v76))}
                                                                                                                                switch v50 {
                                                                                                                                case "hello":
                                                                                                                                    v56 = v56
                                                                                                                                    v1 = append(v1, (fi36 / ((v54) | 1)))
                                                                                                                                default:
                                                                                                                                    var v81 uint = (v28 >> (((v51 * (v28 * v28))) & 63))
                                                                                                                                    _ = v81
                                                                                                                                }
                                                                                                                                if v58 {
                                                                                                                                    v82 := GenS2{F0: (float32(float64(0.0)) < 0.0), F1: genFn0((3.14159 <= ((v55 / v55) - (1e10 / 1.0))), (v76 - v76), 43)}
                                                                                                                                    var v83 uint = (^(v28))
                                                                                                                                    _ = v82
                                                                                                                                    _ = v83
                                                                                                                                }
                                                                                                                                _ = v80
                                                                                                                            } else {
                                                                                                                                var v84 uint16 = 1
                                                                                                                                _ = v84
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                v3[(fi35 + (-(v56)))] = (v50 + "a")
                                                                                                                            } else {
                                                                                                                                v56 += v56
                                                                                                                                if v43 {
                                                                                                                                    v1 = append(v1, 4294967295)
                                                                                                                                } else {
                                                                                                                                    var v85 bool = true
                                                                                                                                    var v86 float64 = 3.14159
                                                                                                                                    var v87 uint = v51
                                                                                                                                    if (!v59) {
                                                                                                                                        var v88 float32 = 0.0
                                                                                                                                        var v89 uint = 18446744073709551615
                                                                                                                                        v1 = append(v1, (-(fi35)))
                                                                                                                                        v90 := GenS5{F0: (v76 & (v76 * uint16(v37))), F1: "\n\t", F2: (true && true)}
                                                                                                                                        if len(v1) >= 2 {
                                                                                                                                            v1 = v1[1:len(v1)]
                                                                                                                                        }
                                                                                                                                        if v58 {
                                                                                                                                            fmt.Println(v87)
                                                                                                                                            switch fi35 {
                                                                                                                                            case -45:
                                                                                                                                                if v59 {
                                                                                                                                                    v37 -= (v37 * 18446744073709551615)
                                                                                                                                                }
                                                                                                                                            case 110:
                                                                                                                                                if ((0.1 - v55) <= 1.5) {
                                                                                                                                                    var v91 uint64 = 1
                                                                                                                                                    switch v50 {
                                                                                                                                                    case "\n\t":
                                                                                                                                                        if v85 {
                                                                                                                                                            v91 = (v91 & v37)
                                                                                                                                                            switch fi35 {
                                                                                                                                                            case 32767:
                                                                                                                                                                v92 := GenS6{F0: v37, F1: (v50 + v50)}
                                                                                                                                                                var v93 int16 = 32767
                                                                                                                                                                _ = v92
                                                                                                                                                                _ = v93
                                                                                                                                                            default:
                                                                                                                                                                v89 = v87
                                                                                                                                                            }
                                                                                                                                                            v34 = (v34 ^ (v34 - v34))
                                                                                                                                                        }
                                                                                                                                                    default:
                                                                                                                                                        v94 := GenS7{F0: int(v52)}
                                                                                                                                                        if true {
                                                                                                                                                            var v95 int64 = (^(v52))
                                                                                                                                                            _ = v95
                                                                                                                                                        }
                                                                                                                                                        _ = v94
                                                                                                                                                    }
                                                                                                                                                    _ = v91
                                                                                                                                                } else {
                                                                                                                                                    v96 := GenS2{F0: v85, F1: genFn0((!(42 != -128)), (v76 >> (((v51 | 434)) & 63)), 255)}
                                                                                                                                                    _ = v96
                                                                                                                                                }
                                                                                                                                                v97 := GenS6{F0: (^(v37)), F1: v50}
                                                                                                                                                _ = v97
                                                                                                                                            default:
                                                                                                                                                v98 := GenS8{F0: 42, F1: 32767}
                                                                                                                                                _ = v98
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        var v99 int16 = 1
                                                                                                                                        _ = v88
                                                                                                                                        _ = v89
                                                                                                                                        _ = v90
                                                                                                                                        _ = v99
                                                                                                                                    } else {
                                                                                                                                        switch 797 {
                                                                                                                                        case 1:
                                                                                                                                            if true {
                                                                                                                                                fmt.Println(fi36, v59, v43)
                                                                                                                                                var v100 int8 = -128
                                                                                                                                                _ = v100
                                                                                                                                            }
                                                                                                                                        case 32767:
                                                                                                                                            delete(v2, 2)
                                                                                                                                        }
                                                                                                                                        v101 := GenS8{F0: 113, F1: 0}
                                                                                                                                        _ = v101
                                                                                                                                    }
                                                                                                                                    _ = len(v1)
                                                                                                                                    _ = v85
                                                                                                                                    _ = v86
                                                                                                                                    _ = v87
                                                                                                                                }
                                                                                                                                if v59 {
                                                                                                                                    var v102 int8 = -1
                                                                                                                                    if (!(255 >= 182)) {
                                                                                                                                        if (((!true) && (((3.14159 / ((v55 * float64(3.14159)) * (v55 * ((3.14159 + ((v55 + v55) * float64(3.14159))) - -2.25)))) + v55) != (v55 + v55))) && (v50 < "a")) {
                                                                                                                                            v59 = ((v76 >> (((v28 | v51)) & 63)) > 65)
                                                                                                                                            switch "" {
                                                                                                                                            case "foo bar":
                                                                                                                                                switch v50 {
                                                                                                                                                case "hello":
                                                                                                                                                    v2[4294967295] = v56
                                                                                                                                                case "a":
                                                                                                                                                    var v103 int16 = 212
                                                                                                                                                    _ = v103
                                                                                                                                                default:
                                                                                                                                                    v104 := GenS7{F0: fi36}
                                                                                                                                                    _ = v104
                                                                                                                                                }
                                                                                                                                                v105 := GenS9{F0: (fi36 & -833), F1: v50}
                                                                                                                                                if true {
                                                                                                                                                    v2[901] = fi36
                                                                                                                                                    fmt.Println(v59)
                                                                                                                                                }
                                                                                                                                                var v106 uint64 = (v37 >> ((18446744073709551615) & 63))
                                                                                                                                                _ = v105
                                                                                                                                                _ = v106
                                                                                                                                            case "":
                                                                                                                                                var v107 uint8 = 255
                                                                                                                                                _ = v107
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                        fmt.Println(v34, v28, v54)
                                                                                                                                    } else {
                                                                                                                                        if v43 {
                                                                                                                                            v1 = append(v1, (v54 % (((v56 & fi35)) | 1)))
                                                                                                                                        } else {
                                                                                                                                            v34 = v34
                                                                                                                                            var v108 int64 = (v52 + v52)
                                                                                                                                            fmt.Println(v76, v102, fi36)
                                                                                                                                            _ = v108
                                                                                                                                        }
                                                                                                                                        v109 := GenS10{F0: ((float64(-2.25) - ((v55 / (v55 * v55)) + v55)) <= v55), F1: v50}
                                                                                                                                        v102 += (v102 & (v102 % ((v102) | 1)))
                                                                                                                                        _ = v109
                                                                                                                                    }
                                                                                                                                    v110 := GenS11{F0: v37, F1: v59, F2: 0}
                                                                                                                                    _ = v102
                                                                                                                                    _ = v110
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v76
                                                                                                                    } else {
                                                                                                                        if v58 {
                                                                                                                            v51 = (^(v51))
                                                                                                                            v56 += fi35
                                                                                                                            var v111 uint64 = (v37 * v37)
                                                                                                                            _ = v111
                                                                                                                        } else {
                                                                                                                            if ((-1 > 738) || (4294967295 >= 72)) {
                                                                                                                                var v112 float64 = (float64(1e10) * v55)
                                                                                                                                _ = v112
                                                                                                                            }
                                                                                                                            var v113 int = fi36
                                                                                                                            _ = v113
                                                                                                                        }
                                                                                                                        if (v34 != v34) {
                                                                                                                            var v114 int64 = (v52 & v52)
                                                                                                                            _ = v114
                                                                                                                        }
                                                                                                                        var v115 uint64 = 4294967295
                                                                                                                        switch (^(v54)) {
                                                                                                                        case 0:
                                                                                                                            switch "a" {
                                                                                                                            case "a":
                                                                                                                                v1 = append(v1, (fi35 * (fi36 | (fi36 | v56))))
                                                                                                                            case "":
                                                                                                                                fmt.Println(fi36, v28, v50)
                                                                                                                            }
                                                                                                                        case -2147483648:
                                                                                                                            var v116 float64 = (v55 * 3.14159)
                                                                                                                            _ = v116
                                                                                                                        }
                                                                                                                        v58 = (v115 <= v115)
                                                                                                                        _ = v115
                                                                                                                    }
                                                                                                                    var v117 bool = v59
                                                                                                                    _ = v117
                                                                                                                } else {
                                                                                                                    v59 = v58
                                                                                                                    switch "\n\t" {
                                                                                                                    case "foo bar":
                                                                                                                        var v118 float64 = v55
                                                                                                                        _ = v118
                                                                                                                    default:
                                                                                                                        _ = len(v1)
                                                                                                                    }
                                                                                                                }
                                                                                                                _ = v58
                                                                                                                _ = v59
                                                                                                            }
                                                                                                        case 9223372036854775807:
                                                                                                            v51 -= (v28 & 255)
                                                                                                        default:
                                                                                                            var v119 int = int(v52)
                                                                                                            _ = v119
                                                                                                        }
                                                                                                    case "a":
                                                                                                        var v120 int64 = (v52 << (((^(v28))) & 63))
                                                                                                        _ = v120
                                                                                                    default:
                                                                                                        if (-37 == -128) {
                                                                                                            fmt.Println(fi35)
                                                                                                            switch fi35 {
                                                                                                            case 127:
                                                                                                                v34 += (v34 << ((v28) & 63))
                                                                                                            case 1:
                                                                                                                v56 += fi36
                                                                                                            case -776:
                                                                                                                fmt.Println(v56, v51, v52)
                                                                                                            default:
                                                                                                                fmt.Println(v54, v54)
                                                                                                            }
                                                                                                            var v121 uint16 = uint16(v37)
                                                                                                            _ = v121
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                if v43 {
                                                                                                    v37 -= 4294967295
                                                                                                    var v122 uint16 = 255
                                                                                                    var v123 int64 = (v52 | v52)
                                                                                                    switch ((v50 + v50) + (v50 + v50)) {
                                                                                                    case "a":
                                                                                                        v28 = (v28 / (((v51 & v51)) | 1))
                                                                                                    case "hello":
                                                                                                        v124 := GenS9{F0: (^(fi36)), F1: "foo bar"}
                                                                                                        _ = v124
                                                                                                    default:
                                                                                                        v3[v54] = v50
                                                                                                    }
                                                                                                    _ = v122
                                                                                                    _ = v123
                                                                                                } else {
                                                                                                    switch (-(fi36)) {
                                                                                                    case -9223372036854775808:
                                                                                                        _ = len(v3)
                                                                                                    case 32768:
                                                                                                        var v125 uint64 = v37
                                                                                                        _ = v125
                                                                                                    default:
                                                                                                        var v126 uint16 = 290
                                                                                                        _ = v126
                                                                                                    }
                                                                                                    var v127 int8 = 60
                                                                                                    v3[(fi36 & -9223372036854775808)] = "a"
                                                                                                    _ = v127
                                                                                                }
                                                                                                _ = v56
                                                                                            } else {
                                                                                                v54 = (fi35 << ((v51) & 63))
                                                                                                v43 = (!((^(v37)) <= v37))
                                                                                            }
                                                                                            _ = v54
                                                                                            _ = v55
                                                                                        }
                                                                                    } else {
                                                                                        var v128 uint32 = 42
                                                                                        var v129 uint = (v51 - v28)
                                                                                        switch (v50 + (v50 + (v50 + ((((v50 + (v50 + v50)) + (v50 + "")) + "") + ("" + v50))))) {
                                                                                        case "\n\t":
                                                                                            var v130 uint64 = (v37 | (v37 << ((v28) & 63)))
                                                                                            _ = v130
                                                                                        case "":
                                                                                            fmt.Println(v52)
                                                                                        default:
                                                                                            v131 := GenS12{F0: uint(v128)}
                                                                                            _ = v131
                                                                                        }
                                                                                        _ = v128
                                                                                        _ = v129
                                                                                    }
                                                                                default:
                                                                                    var v132 int16 = 1000
                                                                                    _ = v132
                                                                                }
                                                                            }
                                                                        case 987:
                                                                            if (v37 < (v37 & v37)) {
                                                                                var v133 string = v50
                                                                                _ = v133
                                                                            } else {
                                                                                fmt.Println(v43, v43)
                                                                                if (65535 > 65535) {
                                                                                    v134 := GenS13{F0: 42}
                                                                                    v135 := GenS8{F0: 0, F1: 740}
                                                                                    _ = v134
                                                                                    _ = v135
                                                                                } else {
                                                                                    var v136 int = (fi35 / ((4294967295) | 1))
                                                                                    var v137 uint64 = 2
                                                                                    var v138 float32 = (-2.25 + float32(3.14159))
                                                                                    v139 := GenS14{F0: 127, F1: 42}
                                                                                    if (v34 == -1) {
                                                                                        v137 = 42
                                                                                        v140 := GenS12{F0: v51}
                                                                                        _ = v140
                                                                                    } else {
                                                                                        var v141 uint64 = (v37 * v37)
                                                                                        _ = len(v1)
                                                                                        _ = v141
                                                                                    }
                                                                                    var v142 uint64 = (^(v37))
                                                                                    switch ((v50 + ("" + v50)) + v50) {
                                                                                    case "hello":
                                                                                        v143 := GenS6{F0: (v142 << ((2) & 63)), F1: v50}
                                                                                        _ = v143
                                                                                    case "foo bar":
                                                                                        v34 -= -990
                                                                                    case "":
                                                                                        _ = len(v1)
                                                                                    default:
                                                                                        v50 = ((("hello" + (v50 + v50)) + ("hello" + "")) + ((v50 + "foo bar") + v50))
                                                                                    }
                                                                                    v2[4294967295] = 1
                                                                                    _ = v136
                                                                                    _ = v137
                                                                                    _ = v138
                                                                                    _ = v139
                                                                                    _ = v142
                                                                                }
                                                                            }
                                                                        }
                                                                        var v144 uint16 = 42
                                                                        _ = v144
                                                                        _ = v50
                                                                        _ = v51
                                                                        _ = v52
                                                                    }
                                                                    switch "a" {
                                                                    case "foo bar":
                                                                        v43 = false
                                                                    case "a":
                                                                        fmt.Println(v37)
                                                                    }
                                                                    if (((3.14159 + -2.25) != 0.0) || (!(1 == -9223372036854775808))) {
                                                                        v145 := v2[2147483648]
                                                                        _ = v145
                                                                    }
                                                                }
                                                                _ = v42
                                                                _ = v43
                                                            case -9223372036854775808:
                                                                var v146 int = (fi36 & -725)
                                                                _ = v146
                                                            }
                                                        default:
                                                            switch int(v28) {
                                                            case 42:
                                                                var v147 uint32 = 42
                                                                _ = v147
                                                            case 2147483647:
                                                                v34 -= int32(v28)
                                                            case 32768:
                                                                var v148 int8 = 1
                                                                _ = v148
                                                            }
                                                        }
                                                    case 0:
                                                        var v149 int8 = 0
                                                        if ((v34 & 1) == (v34 + v34)) {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                v28 -= 86
                                                            case "foo bar":
                                                                fmt.Println(v149, v34, v37)
                                                            case "hello":
                                                                fmt.Println(v149, v37)
                                                            }
                                                            fmt.Println(v37)
                                                            switch (fi35 | int(v149)) {
                                                            case 273:
                                                                v150 := GenS6{F0: v37, F1: "a"}
                                                                _ = v150
                                                            default:
                                                                if true {
                                                                    switch "\n\t" {
                                                                    case "\n\t":
                                                                        if false {
                                                                            var v151 uint = v28
                                                                            var v152 int32 = (v34 * v34)
                                                                            v153 := GenS4{F0: 128, F1: ((v34 >= v34) && ("a" == ("foo bar" + "\n\t")))}
                                                                            _ = v151
                                                                            _ = v152
                                                                            _ = v153
                                                                        } else {
                                                                            v149 -= v149
                                                                            var v154 uint64 = (^(v37))
                                                                            _ = v154
                                                                        }
                                                                    case "foo bar":
                                                                        v155 := GenS15{F0: (fi36 + fi35), F1: 0}
                                                                        _ = v155
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            v156 := GenS5{F0: uint16(v28), F1: (("a" + (("foo bar" + (("hello" + "hello") + ("hello" + "\n\t"))) + (("foo bar" + ("" + "foo bar")) + (("\n\t" + ((("\n\t" + "hello") + "foo bar") + ("hello" + ""))) + "foo bar")))) + ""), F2: true}
                                                            _ = v156
                                                        }
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v157 string = "foo bar"
                                                            _ = v157
                                                        }
                                                        _ = v149
                                                    case 9223372036854775807:
                                                        if false {
                                                            if len(v1) > 0 {
                                                                v1[0] = (^(fi35))
                                                            }
                                                            var v158 int = (fi36 - (-(fi35)))
                                                            var v159 string = "\n\t"
                                                            v158 = (fi36 / ((fi35) | 1))
                                                            var v160 float32 = 0.0
                                                            _ = len(v1)
                                                            _ = v158
                                                            _ = v159
                                                            _ = v160
                                                        } else {
                                                            v28 -= v28
                                                        }
                                                    default:
                                                        v37 -= v37
                                                    }
                                                case "hello":
                                                    if true {
                                                        switch 32767 {
                                                        case -959:
                                                            v37 = v37
                                                        case -2147483648:
                                                            v28 += (v28 ^ (v28 + v28))
                                                        }
                                                        var v161 uint = (v28 ^ (v28 ^ (v28 - (v28 ^ v28))))
                                                        _ = v161
                                                    }
                                                default:
                                                    var v162 string = ""
                                                    _ = v162
                                                }
                                                v163 := GenS16{F0: "foo bar", F1: v34, F2: (^(v28))}
                                                v164 := GenS17{F0: v37, F1: 0, F2: 2147483648}
                                                _ = v163
                                                _ = v164
                                            } else {
                                                fmt.Println(fi35)
                                                var v165 float64 = (0.0 - float64(float32((3.14159 * 1e10))))
                                                v166 := GenS6{F0: v37, F1: ("hello" + ("\n\t" + "\n\t"))}
                                                _ = v165
                                                _ = v166
                                            }
                                        case -326:
                                            v28 = 2
                                            fmt.Println(v37, v34)
                                        }
                                        _ = fi36
                                        _ = v37
                                    }
                                }
                            } else {
                                delete(v3, (fi35 & fi35))
                            }
                            v3[127] = ("a" + "hello")
                            _ = fi35
                        }
                        if (((v28 % (((v28 - v28)) | 1)) > (v28 ^ 0)) && true) {
                            for fi167 := 0; fi167 < 19; fi167++ {
                                v168 := GenS14{F0: fi167, F1: 0}
                                _ = fi167
                                _ = v168
                            }
                        } else {
                            if (192 >= 102) {
                                v169 := GenS18{F0: false, F1: 204}
                                if true {
                                    v2[399] = -9223372036854775808
                                    switch "a" {
                                    case "a":
                                        switch -2147483648 {
                                        case 128:
                                            var v170 int16 = -409
                                            _ = v170
                                        case -1:
                                            var v171 float64 = (0.0 * 0.1)
                                            _ = v171
                                        default:
                                            v172 := GenS19{F0: 65535, F1: 828}
                                            _ = v172
                                        }
                                    case "hello":
                                        var v173 int32 = 341
                                        _ = v173
                                    case "":
                                        if (2147483647 <= (v34 ^ v34)) {
                                            var v174 int = 0
                                            if ((false || true) || false) {
                                                _ = len(v3)
                                                _ = len(v2)
                                            }
                                            v34 += v34
                                            _ = v174
                                        } else {
                                            var v175 uint8 = 3
                                            _ = v175
                                        }
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        if ((false || ((true || false) && true)) && false) {
                                            switch "hello" {
                                            case "a":
                                                for fi176 := 0; fi176 < 11; fi176++ {
                                                    v1 = append(v1, 32767)
                                                    for fi177 := 0; fi177 < 6; fi177++ {
                                                        if (1.5 >= 0.1) {
                                                            v34 = -885
                                                        }
                                                        _ = fi177
                                                    }
                                                    v34 -= (v34 ^ (v34 & -780))
                                                    _ = fi176
                                                }
                                            default:
                                                var v178 uint16 = 65535
                                                _ = v178
                                            }
                                        } else {
                                            var v179 uint8 = 127
                                            v1 = append(v1, -985)
                                            _ = v179
                                        }
                                    case "a":
                                        var v180 uint16 = 42
                                        _ = v180
                                    default:
                                        var v181 int32 = v34
                                        _ = v181
                                    }
                                }
                                _ = v169
                            } else {
                                v28 = 1
                                for fi182 := 0; fi182 < 19; fi182++ {
                                    v183 := GenS20{F0: (v34 & (v34 - 42)), F1: 127}
                                    fmt.Println(v34, v28, fi182)
                                    v184 := GenS5{F0: 163, F1: "foo bar", F2: true}
                                    _ = fi182
                                    _ = v183
                                    _ = v184
                                }
                                fmt.Println(v34)
                            }
                            for fi185 := 0; fi185 < 12; fi185++ {
                                var v186 uint8 = 128
                                var v187 int64 = -1
                                _ = fi185
                                _ = v186
                                _ = v187
                            }
                            for fi188 := 0; fi188 < 15; fi188++ {
                                for fi189 := 0; fi189 < 7; fi189++ {
                                    delete(v2, 62)
                                    fmt.Println(v28, fi189)
                                    _ = fi189
                                }
                                v190 := GenS2{F0: true, F1: -56}
                                v191 := GenS21{F0: "a", F1: (v28 & v28)}
                                _ = fi188
                                _ = v190
                                _ = v191
                            }
                        }
                        _ = v34
                    }
                }
                _ = v28
            }
        }
    }
    v1 = append(v1, -356)
    if false {
        var v192 bool = (920 < -9223372036854775808)
        _ = v192
    } else {
        _ = len(v3)
        if ((65535 <= 860) && (!(!(32767 != 255)))) {
            switch "hello" {
            case "a":
                var v193 bool = (42 <= 18446744073709551615)
                _ = v193
            case "\n\t":
                var v194 uint16 = 815
                _ = v194
            default:
                if true {
                    v195 := v3[32768]
                    for fi196 := 0; fi196 < 7; fi196++ {
                        if (true || true) {
                            if (false && true) {
                                if false {
                                    var v197 uint32 = 42
                                    v198 := GenS3{F0: false, F1: 81}
                                    v199 := v2[v197]
                                    _ = v197
                                    _ = v198
                                    _ = v199
                                }
                                for fi200 := 0; fi200 < 13; fi200++ {
                                    if (42 >= 2) {
                                        var v201 int = fi200
                                        if false {
                                            v202 := GenS22{F0: (-32768 != -272), F1: true, F2: 91}
                                            _ = v202
                                        }
                                        var v203 uint32 = 4294967295
                                        _ = v201
                                        _ = v203
                                    }
                                    var v204 uint = 613
                                    if (32767 < (fi200 - (fi196 % (((^(fi200))) | 1)))) {
                                        delete(v3, (fi196 ^ (fi200 ^ 9223372036854775807)))
                                        switch fi196 {
                                        case 4294967295:
                                            var v205 string = v195
                                            _ = v205
                                        case -329:
                                            if (v204 < v204) {
                                                v204 = (v204 ^ 65535)
                                                v2[1] = fi196
                                            } else {
                                                delete(v3, fi196)
                                                v195 = v195
                                                v206 := v3[(fi200 << ((v204) & 63))]
                                                _ = v206
                                            }
                                        case 9223372036854775807:
                                            if (v195 == "a") {
                                                switch (v195 + v195) {
                                                case "":
                                                    v207 := GenS7{F0: (^(fi196))}
                                                    _ = v207
                                                case "a":
                                                    switch (-(fi200)) {
                                                    case -2147483648:
                                                        var v208 string = v195
                                                        _ = v208
                                                    default:
                                                        fmt.Println(v195)
                                                    }
                                                }
                                                if true {
                                                    fmt.Println(v195)
                                                    var v209 int = 855
                                                    if ((false || (v195 <= v195)) && (127 > 242)) {
                                                        var v210 uint8 = 128
                                                        if (fi196 > (-(v209))) {
                                                            v211 := GenS23{F0: 32767, F1: -9223372036854775808}
                                                            var v212 float64 = 0.1
                                                            _ = v211
                                                            _ = v212
                                                        }
                                                        _ = v210
                                                    } else {
                                                        if true {
                                                            fmt.Println(fi196)
                                                            v213 := GenS22{F0: (205 < 255), F1: true, F2: 2147483648}
                                                            switch (fi200 - fi196) {
                                                            case -179:
                                                                v214 := v2[2147483648]
                                                                _ = v214
                                                            case 127:
                                                                _ = len(v1)
                                                            case 0:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            }
                                                            v215 := v2[2147483647]
                                                            _ = v213
                                                            _ = v215
                                                        } else {
                                                            if false {
                                                                if false {
                                                                    var v216 int16 = 0
                                                                    _ = v216
                                                                }
                                                                v217 := GenS24{F0: v204}
                                                                _ = v217
                                                            }
                                                            v3[-9223372036854775808] = ("hello" + ((("a" + "") + v195) + "\n\t"))
                                                            var v218 uint64 = 280
                                                            _ = v218
                                                        }
                                                        fmt.Println(v209, fi196)
                                                    }
                                                    _ = v209
                                                } else {
                                                    var v219 uint16 = 255
                                                    v195 = v195
                                                    v219 -= 255
                                                    v219 -= (v219 - v219)
                                                    v204 -= v204
                                                    _ = v219
                                                }
                                                v220 := v3[fi196]
                                                v221 := GenS25{F0: fi200, F1: 128, F2: 18446744073709551615}
                                                if true {
                                                    var v222 uint = v204
                                                    v220 = "foo bar"
                                                    _ = v222
                                                } else {
                                                    var v223 int = fi196
                                                    if false {
                                                        if false {
                                                            v224 := GenS23{F0: 3, F1: -1}
                                                            v225 := GenS1{F0: 42, F1: v220}
                                                            var v226 int8 = 100
                                                            v204 += 65535
                                                            var v227 int32 = 2147483647
                                                            _ = v224
                                                            _ = v225
                                                            _ = v226
                                                            _ = v227
                                                        }
                                                    } else {
                                                        var v228 uint64 = 18446744073709551615
                                                        fmt.Println(fi200)
                                                        _ = v228
                                                    }
                                                    fmt.Println(fi196, fi196, v204)
                                                    var v229 float64 = (0.1 / ((1e10 - 0.0) + ((3.14159 - 0.0) * 0.0)))
                                                    var v230 int = fi200
                                                    _ = v223
                                                    _ = v229
                                                    _ = v230
                                                }
                                                if false {
                                                    var v231 uint64 = 75
                                                    v195 = v220
                                                    delete(v2, 263)
                                                    var v232 bool = (fi200 == (^(fi200)))
                                                    _ = v231
                                                    _ = v232
                                                } else {
                                                    switch (v195 + v195) {
                                                    case "foo bar":
                                                        fmt.Println(v195, fi196, fi196)
                                                    }
                                                    v233 := GenS25{F0: (fi196 ^ (fi200 | (fi196 ^ (-(fi200))))), F1: uint8(v204), F2: 1}
                                                    v234 := GenS6{F0: 1, F1: v220}
                                                    v220 = ""
                                                    _ = v233
                                                    _ = v234
                                                }
                                                v220 = "hello"
                                                _ = v220
                                                _ = v221
                                            }
                                        }
                                    }
                                    _ = fi200
                                    _ = v204
                                }
                                var v235 float32 = -2.25
                                _ = v235
                            }
                            for fi236 := 0; fi236 < 19; fi236++ {
                                v237 := GenS20{F0: -866, F1: 127}
                                switch fi196 {
                                case -2147483648:
                                    if (!(!(!true))) {
                                        v195 = v195
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                case 32768:
                                    if true {
                                        v238 := GenS26{F0: 42, F1: 42}
                                        switch v195 {
                                        case "foo bar":
                                            if false {
                                                v195 = v195
                                                v195 = ("hello" + v195)
                                            } else {
                                                v1 = append(v1, 4294967295)
                                                if true {
                                                    v239 := GenS27{F0: "foo bar"}
                                                    var v240 bool = true
                                                    v240 = (!false)
                                                    _ = v239
                                                    _ = v240
                                                }
                                                v195 = (v195 + (v195 + "hello"))
                                            }
                                        case "hello":
                                            fmt.Println(fi236, fi236, fi196)
                                        case "\n\t":
                                            v195 = ((v195 + ((v195 + v195) + (v195 + "a"))) + v195)
                                        }
                                        v2[1] = (fi236 | (fi196 << ((4294967295) & 63)))
                                        _ = v238
                                    } else {
                                        v195 = v195
                                        v241 := GenS28{F0: 35, F1: (!(128 <= 42)), F2: 692}
                                        _ = len(v3)
                                        _ = v241
                                    }
                                case 330:
                                    fmt.Println(fi196, fi236)
                                default:
                                    if (!(true || true)) {
                                        if (!true) {
                                            var v242 uint32 = 65535
                                            switch (v195 + "") {
                                            case "":
                                                v243 := GenS20{F0: 680, F1: genFn0((false || true), 32767, 42)}
                                                _ = v243
                                            case "\n\t":
                                                var v244 int32 = 502
                                                _ = v244
                                            default:
                                                if true {
                                                    v242 = 873
                                                } else {
                                                    v245 := GenS29{F0: 879}
                                                    _ = v245
                                                }
                                            }
                                            _ = v242
                                        } else {
                                            fmt.Println(fi236, fi236)
                                            v195 = ("" + v195)
                                            var v246 uint = 2
                                            _ = v246
                                        }
                                    } else {
                                        if (false && (false || (true && true))) {
                                            if true {
                                                if true {
                                                    switch (v195 + v195) {
                                                    case "foo bar":
                                                        v195 = v195
                                                    default:
                                                        switch v195 {
                                                        case "foo bar":
                                                            v195 = (v195 + ("hello" + v195))
                                                            if false {
                                                                var v247 int16 = -32768
                                                                _ = v247
                                                            } else {
                                                                fmt.Println(fi196, fi236)
                                                                switch fi196 {
                                                                case 9223372036854775807:
                                                                    v195 = v195
                                                                case 454:
                                                                    v195 = (v195 + v195)
                                                                default:
                                                                    var v248 int = (fi236 + fi196)
                                                                    _ = v248
                                                                }
                                                            }
                                                        case "\n\t":
                                                            v195 = ""
                                                        case "":
                                                            v249 := GenS9{F0: (fi196 >> ((18446744073709551615) & 63)), F1: (v195 + "\n\t")}
                                                            _ = v249
                                                        }
                                                        v195 = "foo bar"
                                                    }
                                                    v195 = v195
                                                } else {
                                                    switch "a" {
                                                    case "hello":
                                                        var v250 bool = (-2147483648 > 0)
                                                        _ = v250
                                                    default:
                                                        if true {
                                                            var v251 int64 = -9223372036854775808
                                                            var v252 float64 = 3.14159
                                                            if (32767 != 255) {
                                                                v252 += float64(((0.0 / 3.14159) - float32((1e10 / v252))))
                                                            }
                                                            _ = v251
                                                            _ = v252
                                                        } else {
                                                            v2[712] = (fi196 << ((0) & 63))
                                                        }
                                                    }
                                                    v195 = ("foo bar" + v195)
                                                    if true {
                                                        v195 = v195
                                                        fmt.Println(fi236, fi196, fi236)
                                                    }
                                                }
                                                switch v195 {
                                                case "hello":
                                                    if len(v1) > 0 {
                                                        v1[0] = fi236
                                                    }
                                                    v253 := GenS2{F0: true, F1: -60}
                                                    fmt.Println(fi236, fi196, fi196)
                                                    if len(v1) > 0 {
                                                        v1[0] = fi236
                                                    }
                                                    _ = v253
                                                }
                                            } else {
                                                v195 = v195
                                                v195 = (v195 + ("\n\t" + v195))
                                                var v254 float64 = 3.14159
                                                _ = v254
                                            }
                                            var v255 uint = 42
                                            v1 = append(v1, (fi196 << (((^(v255))) & 63)))
                                            _ = v255
                                        } else {
                                            if true {
                                                v195 = (("\n\t" + (v195 + ("foo bar" + ("a" + "a")))) + v195)
                                                var v256 int32 = -1
                                                _ = v256
                                            } else {
                                                var v257 uint64 = 2
                                                v195 = v195
                                                var v258 int = (fi196 & (fi196 * fi236))
                                                _ = v257
                                                _ = v258
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = fi236
                                            }
                                        }
                                        var v259 bool = true
                                        var v260 bool = (1.5 == 0.0)
                                        _ = v259
                                        _ = v260
                                    }
                                }
                                _ = fi236
                                _ = v237
                            }
                        }
                        _ = fi196
                    }
                    v261 := GenS30{F0: -2147483648}
                    _ = v195
                    _ = v261
                } else {
                    var v262 string = ("a" + ("" + "\n\t"))
                    v262 = "hello"
                    _ = v262
                }
            }
        }
    }
    var v263 uint32 = 0
    var v264 string = ("\n\t" + "hello")
    v265 := GenS5{F0: 32767, F1: (v264 + ((v264 + ("a" + (v264 + (v264 + "")))) + v264)), F2: true}
    switch -1 {
    case -944:
        for fi266 := 0; fi266 < 19; fi266++ {
            switch fi266 {
            case -1:
                fmt.Println(v263, v263)
            case -222:
                for fi267 := 0; fi267 < 9; fi267++ {
                    v268 := GenS8{F0: 0, F1: -1}
                    fmt.Println(fi266)
                    var v269 uint = 0
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v270 int16 = 32767
                    fmt.Println(fi267, v263)
                    var v271 uint64 = uint64(v269)
                    _ = fi267
                    _ = v268
                    _ = v269
                    _ = v270
                    _ = v271
                }
            default:
                var v272 float32 = float32(1.5)
                _ = v272
            }
            v263 = v263
            var v273 int32 = 42
            _ = fi266
            _ = v273
        }
    default:
        var v274 uint16 = 65535
        _ = v274
    }
    v263 += v263
    var v275 int = 9223372036854775807
    v275 = (v275 - v275)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v263:", v263)
    fmt.Println("v264:", v264)
    fmt.Println("v275:", v275)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v263
    _ = v264
    _ = v265
    _ = v275
    _ = v3
}
