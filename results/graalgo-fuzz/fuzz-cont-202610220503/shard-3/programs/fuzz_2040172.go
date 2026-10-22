package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
type GenS3 struct { F0 int32; F1 string; F2 int32 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint32 }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 int8; F1 int16; F2 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int16; F1 int32; F2 int }
type GenS12 struct { F0 int32; F1 uint32; F2 string }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 bool; F1 int8 }
type GenS15 struct { F0 int; F1 int16; F2 uint16 }
type GenS16 struct { F0 int; F1 uint8 }
type GenS17 struct { F0 uint64; F1 int64 }
type GenS18 struct { F0 string; F1 int8 }
type GenS19 struct { F0 int8; F1 int64 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 uint16 }
type GenS22 struct { F0 int8 }
type GenS23 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS24 struct { F0 uint8; F1 uint; F2 int64 }
type GenS25 struct { F0 string; F1 uint16; F2 bool }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 int8; F1 uint; F2 string }
func genFn0(p0 bool, p1 int8, p2 uint16) bool {
    p1 = (p1 << ((2) & 63))
    fmt.Println(p0, p2, p1)
    v7 := GenS1{F0: 834, F1: (-(p1))}
    if false {
        var v8 uint8 = 0
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (9223372036854775807 == 1)
}

func main() {
    var v1 uint16 = 2
    var v2 uint16 = v1
    v3 := []int{-973, -589, 9223372036854775807, 710}
    var v4 int16 = -32768
    v5 := make(map[int]bool)
    v6 := []uint32{2147483647, 1, 2}
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(fi9, v1, v1)
        v1 += (v1 * v2)
        _ = fi9
    }
    var v10 int = 1
    if (true && (float32(0.0) <= (0.1 + (1.5 / 1.0)))) {
        _ = len(v5)
        v11 := GenS1{F0: 0, F1: 1}
        _ = v11
    } else {
        switch 9223372036854775807 {
        case 128:
            v5[(v10 << ((255) & 63))] = false
        default:
            v10 -= (v10 << ((255) & 63))
        }
        if false {
            var v12 bool = true
            switch (v10 | 4294967295) {
            case -1:
                v2 = (v2 & 300)
            }
            _ = v12
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                if false {
                    v14 := GenS2{F0: 2147483647, F1: 65535, F2: "a"}
                    var v15 uint16 = v2
                    _ = v14
                    _ = v15
                }
                _ = fi13
            }
            v16 := GenS3{F0: 2147483647, F1: (("a" + "") + "foo bar"), F2: -1}
            var v17 int32 = 1
            _ = v16
            _ = v17
        }
        fmt.Println(v2, v10, v10)
    }
    var v18 uint = 926
    if true {
        if false {
            fmt.Println(v10, v10, v18)
            v10 = 703
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            for fi20 := 0; fi20 < 17; fi20++ {
                if (v4 <= (v4 % ((v4) | 1))) {
                    v21 := GenS3{F0: 2147483647, F1: "", F2: 42}
                    if true {
                        var v22 uint16 = (v1 >> ((v18) & 63))
                        if len(v3) > 0 {
                            v3[0] = fi19
                        }
                        var v23 int = (^(fi20))
                        _ = v22
                        _ = v23
                    } else {
                        var v24 int8 = 1
                        if true {
                            v24 -= (v24 - v24)
                        } else {
                            if true {
                                var v25 uint16 = (v1 % ((v1) | 1))
                                if false {
                                    var v26 int8 = v24
                                    var v27 string = "\n\t"
                                    var v28 int8 = (v24 / (((v26 >> (((v18 % (((^(v18))) | 1))) & 63))) | 1))
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                _ = v25
                            }
                            switch "\n\t" {
                            case "":
                                if (!(!false)) {
                                    fmt.Println(v1)
                                    var v29 uint64 = 42
                                    v30 := GenS3{F0: -1, F1: "foo bar", F2: -21}
                                    _ = v29
                                    _ = v30
                                }
                            case "hello":
                                v5[fi20] = (!true)
                            }
                            if (!(!false)) {
                                switch ("foo bar" + "a") {
                                case "hello":
                                    v18 = (v18 * 2)
                                case "foo bar":
                                    v31 := v5[-229]
                                    _ = v31
                                case "a":
                                    v10 = (v10 / (((fi20 << (((v18 | (v18 | (v18 - v18)))) & 63))) | 1))
                                }
                                switch "" {
                                case "":
                                    v32 := GenS1{F0: 4294967295, F1: (v24 / ((42) | 1))}
                                    _ = v32
                                case "a":
                                    var v33 bool = (v18 <= (^(v18)))
                                    _ = v33
                                case "\n\t":
                                    v10 = 0
                                }
                                switch ((("hello" + "a") + ("hello" + ("a" + ("hello" + "")))) + "") {
                                case "hello":
                                    v34 := GenS2{F0: 312, F1: 0, F2: ("" + "foo bar")}
                                    _ = v34
                                case "\n\t":
                                    fmt.Println(v1)
                                }
                            } else {
                                v10 = 128
                                if ("" <= "hello") {
                                    fmt.Println(v24, v4, v10)
                                    switch ("a" + "\n\t") {
                                    case "\n\t":
                                        switch -9223372036854775808 {
                                        case 0:
                                            if len(v6) > 0 {
                                                v6[0] = 2147483647
                                            }
                                        case -678:
                                            fmt.Println(v2, v4)
                                        }
                                    case "":
                                        if (((("\n\t" + "") == "") || ((84 <= 18446744073709551615) || (!(false || false)))) || (127 == 1)) {
                                            if true {
                                                v1 = v1
                                                if (!true) {
                                                    v35 := GenS1{F0: 18446744073709551615, F1: (v24 | (v24 - 127))}
                                                    v1 = (v1 >> (((v18 ^ 65535)) & 63))
                                                    fmt.Println(v18)
                                                    _ = v35
                                                }
                                            } else {
                                                v2 += (v1 & uint16(v18))
                                                if true {
                                                    v1 += v1
                                                    v36 := v5[(v10 << (((v18 * 255)) & 63))]
                                                    v37 := GenS4{F0: v24, F1: 9223372036854775807}
                                                    switch "hello" {
                                                    case "\n\t":
                                                        v38 := GenS3{F0: -51, F1: "hello", F2: 541}
                                                        _ = v38
                                                    case "a":
                                                        v18 = (^(v18))
                                                        switch "foo bar" {
                                                        case "hello":
                                                            v1 = (^(v2))
                                                        case "a":
                                                            var v39 float64 = ((0.0 * (-2.25 / 3.14159)) / (1e10 / 1.5))
                                                            _ = v39
                                                        case "\n\t":
                                                            var v40 string = "\n\t"
                                                            _ = v40
                                                        default:
                                                            if false {
                                                                if (!true) {
                                                                    v41 := v5[fi20]
                                                                    _ = v41
                                                                } else {
                                                                    v42 := GenS4{F0: int8(fi20), F1: 42}
                                                                    var v43 uint8 = 227
                                                                    var v44 int32 = 2147483647
                                                                    _ = v42
                                                                    _ = v43
                                                                    _ = v44
                                                                }
                                                                var v45 float32 = 3.14159
                                                                _ = v45
                                                            }
                                                        }
                                                        switch "a" {
                                                        case "\n\t":
                                                            switch "" {
                                                            case "\n\t":
                                                                var v46 uint16 = (v2 ^ 42)
                                                                _ = v46
                                                            case "foo bar":
                                                                switch 1 {
                                                                case 127:
                                                                    var v47 float32 = 1e10
                                                                    _ = v47
                                                                case -9223372036854775808:
                                                                    v2 = 65535
                                                                case -138:
                                                                    if len(v6) > 0 {
                                                                        v6[0] = 462
                                                                    }
                                                                default:
                                                                    v36 = v36
                                                                }
                                                            case "":
                                                                var v48 uint16 = v2
                                                                _ = v48
                                                            default:
                                                                var v49 uint16 = (v2 + 511)
                                                                _ = v49
                                                            }
                                                        default:
                                                            v50 := GenS5{F0: (v36 && (-30 > 1))}
                                                            _ = v50
                                                        }
                                                    default:
                                                        var v51 int64 = -1
                                                        _ = v51
                                                    }
                                                    var v52 int32 = 2147483647
                                                    v53 := GenS6{F0: 1, F1: 209}
                                                    if len(v6) > 0 {
                                                        v6[0] = 4294967295
                                                    }
                                                    _ = v36
                                                    _ = v37
                                                    _ = v52
                                                    _ = v53
                                                }
                                            }
                                            v54 := GenS7{F0: 42, F1: true}
                                            v55 := GenS2{F0: 581, F1: 139, F2: ""}
                                            _ = v54
                                            _ = v55
                                        } else {
                                            var v56 uint64 = 1
                                            _ = v56
                                        }
                                    }
                                    v5[int(v4)] = (256 != 18446744073709551615)
                                } else {
                                    switch "" {
                                    case "":
                                        if (!(false && (int16(v18) < v4))) {
                                            if true {
                                                if true {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    if false {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            var v57 uint16 = (^(v2))
                                                            _ = v57
                                                        default:
                                                            v18 = v18
                                                        }
                                                    } else {
                                                        v4 -= -1
                                                        if (!false) {
                                                            v5[(fi19 << (((v18 << ((v18) & 63))) & 63))] = false
                                                            fmt.Println(v1, v2, v10)
                                                            v4 -= v4
                                                        } else {
                                                            var v58 uint32 = 1
                                                            v5[(fi19 % ((int(v2)) | 1))] = false
                                                            v2 += (v1 >> ((18446744073709551615) & 63))
                                                            _ = v58
                                                        }
                                                    }
                                                    if true {
                                                        v59 := GenS8{F0: v24, F1: v4, F2: v4}
                                                        fmt.Println(fi20, fi19)
                                                        if false {
                                                            switch v10 {
                                                            case -9223372036854775808:
                                                                if (!false) {
                                                                    fmt.Println(fi20)
                                                                } else {
                                                                    v18 += v18
                                                                    if true {
                                                                        switch (("a" + ("a" + ("foo bar" + "hello"))) + ("foo bar" + "")) {
                                                                        case "":
                                                                            v24 = v24
                                                                        default:
                                                                            switch fi19 {
                                                                            case 9223372036854775807:
                                                                                fmt.Println(v2, v18, fi19)
                                                                            default:
                                                                                if false {
                                                                                    v6 = append(v6, 42)
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    delete(v5, fi19)
                                                                                    if (-97 >= 784) {
                                                                                        v60 := GenS9{F0: 188}
                                                                                        if ((((("hello" + ("a" + "")) + ("foo bar" + "foo bar")) + ("a" + "hello")) + ("hello" + ("a" + "foo bar"))) < "foo bar") {
                                                                                            v3 = append(v3, fi19)
                                                                                            if false {
                                                                                                if (((v1 / ((v2) | 1)) == (v1 | 288)) && false) {
                                                                                                    var v61 float32 = (float32(-2.25) + 3.14159)
                                                                                                    _ = v61
                                                                                                }
                                                                                                var v62 float32 = (((0.0 + 0.1) * 1e10) / 3.14159)
                                                                                                _ = len(v3)
                                                                                                _ = v62
                                                                                            } else {
                                                                                                v24 += (v24 << (((v18 * 17)) & 63))
                                                                                                v63 := GenS10{F0: -87}
                                                                                                _ = v63
                                                                                            }
                                                                                            v64 := v5[(fi20 + (v10 / ((72) | 1)))]
                                                                                            _ = v64
                                                                                        }
                                                                                        v65 := GenS7{F0: 2147483648, F1: (!(false && (v2 != v1)))}
                                                                                        var v66 bool = (false && true)
                                                                                        _ = v60
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    }
                                                                                    fmt.Println(v10, v2)
                                                                                case "\n\t":
                                                                                    _ = len(v3)
                                                                                }
                                                                            }
                                                                        }
                                                                        v67 := GenS3{F0: 2147483647, F1: "", F2: 5}
                                                                        if true {
                                                                            if ((false || true) && false) {
                                                                                v1 = v2
                                                                                v2 += v1
                                                                            } else {
                                                                                v3 = append(v3, v10)
                                                                            }
                                                                        }
                                                                        _ = v67
                                                                    } else {
                                                                        fmt.Println(v1, fi19, v24)
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            switch "a" {
                                                                            case "":
                                                                                switch ("foo bar" + ((("\n\t" + "a") + ("\n\t" + "foo bar")) + "hello")) {
                                                                                case "\n\t":
                                                                                    if true {
                                                                                        var v68 int32 = -297
                                                                                        var v69 uint16 = v2
                                                                                        v1 = 32767
                                                                                        _ = v68
                                                                                        _ = v69
                                                                                    } else {
                                                                                        fmt.Println(v24, v18, v1)
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = fi19
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    if true {
                                                                                        if (true || (true || true)) {
                                                                                            if false {
                                                                                                var v70 int = 946
                                                                                                v71 := GenS6{F0: 602, F1: 1}
                                                                                                if true {
                                                                                                    if true {
                                                                                                        if (1e10 == 0.0) {
                                                                                                            if (65535 != (^(v18))) {
                                                                                                                var v72 uint16 = uint16(v18)
                                                                                                                _ = v72
                                                                                                            }
                                                                                                            if (v70 == (v70 * v10)) {
                                                                                                                var v73 uint8 = 30
                                                                                                                v1 -= (^(v1))
                                                                                                                _ = v73
                                                                                                            } else {
                                                                                                                fmt.Println(fi19, fi20)
                                                                                                            }
                                                                                                        } else {
                                                                                                            v24 = (v24 & v24)
                                                                                                        }
                                                                                                    } else {
                                                                                                        v70 -= (fi20 & fi20)
                                                                                                        if (v1 == v2) {
                                                                                                            fmt.Println(fi19, v70, v1)
                                                                                                            var v74 int16 = -426
                                                                                                            _ = v74
                                                                                                        } else {
                                                                                                            _ = len(v3)
                                                                                                            v2 += v1
                                                                                                            v75 := GenS11{F0: 1000, F1: 2147483647, F2: (fi19 + v10)}
                                                                                                            _ = v75
                                                                                                        }
                                                                                                    }
                                                                                                    if true {
                                                                                                        if true {
                                                                                                            if false {
                                                                                                                switch -1 {
                                                                                                                case 42:
                                                                                                                    v76 := GenS12{F0: -936, F1: 697, F2: ("" + ("foo bar" + (("" + "foo bar") + "hello")))}
                                                                                                                    _ = v76
                                                                                                                }
                                                                                                                fmt.Println(v24, v2, fi20)
                                                                                                                _ = len(v6)
                                                                                                            }
                                                                                                            v24 += -128
                                                                                                            v77 := GenS13{F0: (v2 & (v1 - 914))}
                                                                                                            switch "a" {
                                                                                                            case "\n\t":
                                                                                                                if false {
                                                                                                                    v1 -= 495
                                                                                                                    if false {
                                                                                                                        v1 += (v1 & (v1 >> ((v18) & 63)))
                                                                                                                    }
                                                                                                                }
                                                                                                            case "a":
                                                                                                                switch (fi19 | 4294967295) {
                                                                                                                case 913:
                                                                                                                    var v78 uint32 = 4294967295
                                                                                                                    _ = v78
                                                                                                                }
                                                                                                            case "":
                                                                                                                delete(v5, (v70 >> (((v18 - (^(v18)))) & 63)))
                                                                                                            default:
                                                                                                                var v79 int16 = v4
                                                                                                                _ = v79
                                                                                                            }
                                                                                                            _ = v77
                                                                                                        }
                                                                                                        v80 := GenS14{F0: (!((-(v4)) <= (v4 - (v4 | (^(v4)))))), F1: v24}
                                                                                                        var v81 int = 32768
                                                                                                        if true {
                                                                                                            var v82 int16 = v4
                                                                                                            _ = v82
                                                                                                        } else {
                                                                                                            if false {
                                                                                                                v4 = int16(v10)
                                                                                                            } else {
                                                                                                                var v83 string = "foo bar"
                                                                                                                _ = v83
                                                                                                            }
                                                                                                            v18 = v18
                                                                                                            var v84 uint32 = 0
                                                                                                            _ = v84
                                                                                                        }
                                                                                                        fmt.Println(fi20, v1)
                                                                                                        delete(v5, v10)
                                                                                                        v2 = v1
                                                                                                        _ = v80
                                                                                                        _ = v81
                                                                                                    }
                                                                                                } else {
                                                                                                    var v85 bool = (((!((v24 * (v24 >> (((v18 | (v18 ^ v18))) & 63))) < v24)) && false) && false)
                                                                                                    var v86 int16 = 42
                                                                                                    _ = v85
                                                                                                    _ = v86
                                                                                                }
                                                                                                if true {
                                                                                                    var v87 int = (v10 * fi20)
                                                                                                    switch "a" {
                                                                                                    case "hello":
                                                                                                        switch (v10 & int(v18)) {
                                                                                                        case 42:
                                                                                                            v88 := GenS15{F0: 32768, F1: v4, F2: 261}
                                                                                                            v89 := GenS16{F0: fi19, F1: 127}
                                                                                                            _ = v88
                                                                                                            _ = v89
                                                                                                        case 860:
                                                                                                            v90 := GenS17{F0: 2, F1: 0}
                                                                                                            _ = v90
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        var v91 uint = (v18 ^ (v18 / ((v18) | 1)))
                                                                                                        _ = v91
                                                                                                    default:
                                                                                                        var v92 uint = (^(v18))
                                                                                                        _ = v92
                                                                                                    }
                                                                                                    _ = v87
                                                                                                } else {
                                                                                                    v1 = 0
                                                                                                    var v93 float32 = -2.25
                                                                                                    switch 2147483647 {
                                                                                                    case 34:
                                                                                                        _ = len(v5)
                                                                                                    case 128:
                                                                                                        var v94 uint16 = v1
                                                                                                        _ = v94
                                                                                                    }
                                                                                                    _ = v93
                                                                                                }
                                                                                                if (v2 > 0) {
                                                                                                    switch -921 {
                                                                                                    case 497:
                                                                                                        var v95 uint16 = (v1 - v2)
                                                                                                        _ = v95
                                                                                                    }
                                                                                                    v4 = v4
                                                                                                    var v96 uint16 = v1
                                                                                                    _ = v96
                                                                                                } else {
                                                                                                    switch "" {
                                                                                                    case "foo bar":
                                                                                                        switch ("a" + "\n\t") {
                                                                                                        case "a":
                                                                                                            v70 += v70
                                                                                                        case "hello":
                                                                                                            if (v24 != int8(v10)) {
                                                                                                                _ = len(v5)
                                                                                                            } else {
                                                                                                                v10 = (v70 << ((v18) & 63))
                                                                                                            }
                                                                                                        default:
                                                                                                            fmt.Println(v18, v70)
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v1, v1)
                                                                                                    v97 := GenS9{F0: 37}
                                                                                                    _ = v97
                                                                                                }
                                                                                                switch ("foo bar" + "hello") {
                                                                                                case "":
                                                                                                    v4 = v4
                                                                                                case "hello":
                                                                                                    v10 -= fi20
                                                                                                case "\n\t":
                                                                                                    v4 -= v4
                                                                                                }
                                                                                                v98 := GenS4{F0: (v24 - (v24 | v24)), F1: 912}
                                                                                                _ = v70
                                                                                                _ = v71
                                                                                                _ = v98
                                                                                            } else {
                                                                                                if (v18 != (v18 << ((304) & 63))) {
                                                                                                    var v99 float32 = ((0.0 + 0.1) + (float32(((-2.25 - 1.5) + 3.14159)) / float32(float64((3.14159 - (((float32(((1e10 - 1e10) / (1e10 * 0.1))) / 1.5) + 0.0) - (0.1 / 1.5)))))))
                                                                                                    v100 := GenS18{F0: "hello", F1: (-(v24))}
                                                                                                    _ = v100
                                                                                                    _ = v99
                                                                                                }
                                                                                                var v101 uint8 = 1
                                                                                                _ = v101
                                                                                            }
                                                                                            if (!false) {
                                                                                                v24 += (v24 >> ((2) & 63))
                                                                                                v24 -= (v24 ^ (-(v24)))
                                                                                            } else {
                                                                                                _ = len(v6)
                                                                                            }
                                                                                            var v102 int = (fi20 / (((fi20 ^ (^(fi19)))) | 1))
                                                                                            v102 = v10
                                                                                            _ = v102
                                                                                        } else {
                                                                                            v10 += -909
                                                                                            var v103 int8 = (v24 << (((v18 / ((v18) | 1))) & 63))
                                                                                            v18 = 772
                                                                                            _ = v103
                                                                                        }
                                                                                        var v104 uint32 = 2147483647
                                                                                        if true {
                                                                                            var v105 uint16 = (v1 - 32767)
                                                                                            fmt.Println(v2, v104, v18)
                                                                                            var v106 float32 = (0.1 + (float32(0.1) + 3.14159))
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                        _ = v104
                                                                                    } else {
                                                                                        switch v10 {
                                                                                        case 226:
                                                                                            if false {
                                                                                                v3 = append(v3, v10)
                                                                                                v24 = (v24 << (((v18 & 128)) & 63))
                                                                                                switch "a" {
                                                                                                case "\n\t":
                                                                                                    delete(v5, (fi19 + (fi20 + (v10 / ((fi20) | 1)))))
                                                                                                default:
                                                                                                    if (true || true) {
                                                                                                        var v107 uint8 = 42
                                                                                                        v108 := GenS7{F0: 0, F1: true}
                                                                                                        _ = v107
                                                                                                        _ = v108
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                delete(v5, 128)
                                                                                            }
                                                                                        case -395:
                                                                                            v24 -= int8(fi20)
                                                                                        case -293:
                                                                                            var v109 float64 = (3.14159 + -2.25)
                                                                                            if false {
                                                                                                fmt.Println(v1)
                                                                                            }
                                                                                            v110 := v5[fi20]
                                                                                            var v111 uint32 = 278
                                                                                            v6 = append(v6, (v111 - 65535))
                                                                                            v1 -= v1
                                                                                            fmt.Println(v1, fi20, v10)
                                                                                            if (!v110) {
                                                                                                v1 += v1
                                                                                            } else {
                                                                                                v24 += 0
                                                                                                if v110 {
                                                                                                    var v112 int8 = (v24 / ((v24) | 1))
                                                                                                    _ = v112
                                                                                                } else {
                                                                                                    var v113 int64 = 9223372036854775807
                                                                                                    if true {
                                                                                                        v18 = (^(v18))
                                                                                                        v114 := GenS9{F0: 127}
                                                                                                        v115 := GenS5{F0: v110}
                                                                                                        _ = v114
                                                                                                        _ = v115
                                                                                                    }
                                                                                                    switch fi20 {
                                                                                                    case 32767:
                                                                                                        v116 := GenS19{F0: (^(v24)), F1: v113}
                                                                                                        _ = v116
                                                                                                    case 0:
                                                                                                        v111 = uint32(v2)
                                                                                                    default:
                                                                                                        if (!(v18 > (v18 & v18))) {
                                                                                                            switch (^(fi20)) {
                                                                                                            case 841:
                                                                                                                var v117 uint16 = 42
                                                                                                                _ = v117
                                                                                                            case 2147483647:
                                                                                                                delete(v5, (fi20 | (fi20 ^ (v10 & fi20))))
                                                                                                            default:
                                                                                                                if v110 {
                                                                                                                    v118 := GenS12{F0: 42, F1: (v111 / ((v111) | 1)), F2: "foo bar"}
                                                                                                                    v1 = 2
                                                                                                                    v1 = (^(v1))
                                                                                                                    v1 = (v2 / (((v2 - v2)) | 1))
                                                                                                                    _ = len(v5)
                                                                                                                    fmt.Println(v110, v2)
                                                                                                                    v2 += v2
                                                                                                                    _ = len(v6)
                                                                                                                    var v119 uint32 = (v111 | (v111 << (((v18 * v18)) & 63)))
                                                                                                                    v18 = (v18 - v18)
                                                                                                                    _ = v118
                                                                                                                    _ = v119
                                                                                                                } else {
                                                                                                                    var v120 uint16 = (v2 - 2)
                                                                                                                    _ = len(v5)
                                                                                                                    var v121 int32 = 1
                                                                                                                    var v122 int64 = v113
                                                                                                                    _ = len(v3)
                                                                                                                    _ = v120
                                                                                                                    _ = v121
                                                                                                                    _ = v122
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            switch "\n\t" {
                                                                                                            case "\n\t":
                                                                                                                var v123 float64 = ((3.14159 / v109) / v109)
                                                                                                                _ = v123
                                                                                                            case "":
                                                                                                                if (!(18446744073709551615 >= 0)) {
                                                                                                                    if ((v111 + (v111 * (v111 + (v111 % ((65535) | 1))))) > v111) {
                                                                                                                        fmt.Println(fi20)
                                                                                                                        var v124 uint16 = v1
                                                                                                                        _ = v124
                                                                                                                    } else {
                                                                                                                        v111 -= v111
                                                                                                                        v125 := GenS11{F0: (v4 * (v4 << ((959) & 63))), F1: 725, F2: 204}
                                                                                                                        _ = v125
                                                                                                                    }
                                                                                                                    v18 = 958
                                                                                                                } else {
                                                                                                                    v126 := GenS2{F0: 2, F1: (v111 * (v111 - (v111 >> (((v18 << ((v18) & 63))) & 63)))), F2: "\n\t"}
                                                                                                                    if false {
                                                                                                                        switch "" {
                                                                                                                        case "hello":
                                                                                                                            var v127 int16 = 575
                                                                                                                            _ = v127
                                                                                                                        }
                                                                                                                        delete(v5, fi20)
                                                                                                                        v4 = v4
                                                                                                                        v128 := GenS19{F0: 100, F1: (-(v113))}
                                                                                                                        v24 = v24
                                                                                                                        _ = v128
                                                                                                                    }
                                                                                                                    var v129 uint16 = 962
                                                                                                                    if len(v6) >= 2 {
                                                                                                                        v6 = v6[1:len(v6)]
                                                                                                                    }
                                                                                                                    _ = v126
                                                                                                                    _ = v129
                                                                                                                }
                                                                                                            case "a":
                                                                                                                var v130 int8 = v24
                                                                                                                _ = v130
                                                                                                            }
                                                                                                            v131 := GenS12{F0: -2147483648, F1: 120, F2: ""}
                                                                                                            v1 = v1
                                                                                                            _ = v131
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v113
                                                                                                }
                                                                                            }
                                                                                            delete(v5, (^(v10)))
                                                                                            _ = v109
                                                                                            _ = v110
                                                                                            _ = v111
                                                                                        default:
                                                                                            if (!true) {
                                                                                                if false {
                                                                                                    v4 += (v4 ^ v4)
                                                                                                    if (127 != 105) {
                                                                                                        v4 += v4
                                                                                                        var v132 int32 = 0
                                                                                                        _ = v132
                                                                                                    }
                                                                                                } else {
                                                                                                    fmt.Println(v4, v4, v10)
                                                                                                    v24 = (v24 & (-(v24)))
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(v4)
                                                                                    }
                                                                                case "":
                                                                                    fmt.Println(v24)
                                                                                default:
                                                                                    v1 = 0
                                                                                }
                                                                            }
                                                                        case "a":
                                                                            fmt.Println(v4)
                                                                        case "\n\t":
                                                                            if (0 == -2147483648) {
                                                                                switch ("" + "\n\t") {
                                                                                case "\n\t":
                                                                                    fmt.Println(v1, v1)
                                                                                default:
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = 323
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                delete(v5, (v10 / ((412) | 1)))
                                                                                var v133 int8 = (v24 ^ (^(v24)))
                                                                                _ = v133
                                                                            }
                                                                        }
                                                                        if ("\n\t" == "") {
                                                                            if false {
                                                                                fmt.Println(v18, v1)
                                                                            } else {
                                                                                v2 += (^(v2))
                                                                                if (!false) {
                                                                                    delete(v5, (fi20 | fi19))
                                                                                    switch (fi20 % ((int(v18)) | 1)) {
                                                                                    case -337:
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            var v134 int = (v10 * (^(v10)))
                                                                                            _ = v134
                                                                                        default:
                                                                                            v135 := GenS9{F0: 82}
                                                                                            _ = v135
                                                                                        }
                                                                                    case 32767:
                                                                                        var v136 uint = (v18 ^ v18)
                                                                                        _ = v136
                                                                                    case 129:
                                                                                        v18 = (v18 >> ((v18) & 63))
                                                                                    }
                                                                                    v137 := GenS13{F0: (v1 >> ((v18) & 63))}
                                                                                    _ = v137
                                                                                }
                                                                            }
                                                                            var v138 int = (-(fi19))
                                                                            _ = v138
                                                                        }
                                                                    }
                                                                    var v139 int16 = (v4 >> (((v18 | v18)) & 63))
                                                                    if true {
                                                                        v2 -= 255
                                                                    }
                                                                    switch ("hello" + ("foo bar" + "a")) {
                                                                    case "a":
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    case "foo bar":
                                                                        v140 := GenS13{F0: v2}
                                                                        _ = v140
                                                                    case "":
                                                                        var v141 int16 = -1
                                                                        _ = v141
                                                                    default:
                                                                        if false {
                                                                            v24 -= -78
                                                                            var v142 int16 = (v4 - (v139 - (v4 ^ (-(v139)))))
                                                                            var v143 uint16 = (v2 % (((v2 * v1)) | 1))
                                                                            _ = v142
                                                                            _ = v143
                                                                        } else {
                                                                            v144 := GenS16{F0: fi19, F1: 42}
                                                                            var v145 uint = 311
                                                                            _ = len(v6)
                                                                            if true {
                                                                                v139 -= v139
                                                                                v18 -= v18
                                                                            } else {
                                                                                v18 += 4294967295
                                                                                if ((1e10 - (-2.25 / 1.5)) >= (0.0 / 3.14159)) {
                                                                                    v1 = (^(v2))
                                                                                    fmt.Println(v10)
                                                                                    var v146 uint64 = 0
                                                                                    _ = v146
                                                                                }
                                                                                v24 += (-(v24))
                                                                            }
                                                                            _ = v144
                                                                            _ = v145
                                                                        }
                                                                    }
                                                                    _ = v139
                                                                }
                                                            case 636:
                                                                if true {
                                                                    v18 += (v18 + 1)
                                                                } else {
                                                                    if (42 != 42) {
                                                                        var v147 int64 = -1
                                                                        var v148 bool = false
                                                                        _ = v147
                                                                        _ = v148
                                                                    }
                                                                    v149 := GenS20{F0: 1}
                                                                    _ = v149
                                                                }
                                                            case -790:
                                                                switch "a" {
                                                                case "hello":
                                                                    var v150 float32 = (0.0 + ((0.0 / 1e10) * -2.25))
                                                                    _ = v150
                                                                default:
                                                                    v18 += v18
                                                                }
                                                            }
                                                            switch -2147483648 {
                                                            case -2147483648:
                                                                v10 -= (fi19 ^ int(v2))
                                                            case -1:
                                                                var v151 float64 = -2.25
                                                                _ = v151
                                                            }
                                                        }
                                                        var v152 uint64 = 18446744073709551615
                                                        if false {
                                                            v153 := GenS21{F0: v2}
                                                            _ = v153
                                                        } else {
                                                            if (true && true) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = -9223372036854775808
                                                                }
                                                                v152 -= (^(v152))
                                                                switch 4294967295 {
                                                                case 42:
                                                                    if ((128 >= 127) && (2 <= 67)) {
                                                                        v154 := v5[v10]
                                                                        if len(v6) > 0 {
                                                                            v6[0] = 2
                                                                        }
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            fmt.Println(v24, v10)
                                                                        case "":
                                                                            v154 = v154
                                                                        default:
                                                                            var v155 int8 = (v24 | -128)
                                                                            v156 := GenS4{F0: (v24 - v155), F1: -9223372036854775808}
                                                                            v157 := GenS10{F0: v155}
                                                                            _ = v155
                                                                            _ = v156
                                                                            _ = v157
                                                                        }
                                                                        _ = v154
                                                                    }
                                                                case 32768:
                                                                    var v158 float64 = (1e10 - 1.5)
                                                                    switch (-(v10)) {
                                                                    case -715:
                                                                        var v159 int16 = (v4 >> (((v18 % (((^(v18))) | 1))) & 63))
                                                                        _ = v159
                                                                    case 483:
                                                                        if ((v24 ^ v24) > 100) {
                                                                            v158 = v158
                                                                            if false {
                                                                                var v160 int8 = (v24 >> (((v18 | (v18 - v18))) & 63))
                                                                                fmt.Println(v4, v160, v158)
                                                                                _ = v160
                                                                            } else {
                                                                                v24 += int8(v18)
                                                                                if false {
                                                                                    fmt.Println(v152)
                                                                                } else {
                                                                                    fmt.Println(v24)
                                                                                    var v161 float64 = ((v158 + 0.0) / float64(1e10))
                                                                                    _ = v161
                                                                                }
                                                                                var v162 uint64 = 4294967295
                                                                                _ = v162
                                                                            }
                                                                            v158 -= (float64(0.1) + ((3.14159 * (v158 / v158)) * -2.25))
                                                                            var v163 float32 = -2.25
                                                                            if true {
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                fmt.Println(v152, v152, v1)
                                                                            }
                                                                            _ = v163
                                                                        } else {
                                                                            var v164 int8 = (^(v24))
                                                                            switch 22 {
                                                                            case 351:
                                                                                v10 = (fi19 + fi20)
                                                                            case -364:
                                                                                v165 := GenS17{F0: v152, F1: -9223372036854775808}
                                                                                _ = v165
                                                                            default:
                                                                                var v166 int32 = 0
                                                                                _ = v166
                                                                            }
                                                                            _ = v164
                                                                        }
                                                                    case 4:
                                                                        switch ("" + "hello") {
                                                                        case "":
                                                                            if ((v152 < (v152 - v152)) && false) {
                                                                                v167 := GenS22{F0: -128}
                                                                                _ = len(v6)
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                _ = v167
                                                                            }
                                                                        case "foo bar":
                                                                            var v168 uint64 = 35
                                                                            _ = v168
                                                                        }
                                                                    default:
                                                                        var v169 uint32 = 587
                                                                        _ = v169
                                                                    }
                                                                    _ = v158
                                                                default:
                                                                    v2 -= v1
                                                                }
                                                            } else {
                                                                fmt.Println(v1, v18)
                                                                v10 = (fi20 ^ (v10 | (fi19 ^ (fi20 % ((fi19) | 1)))))
                                                            }
                                                            v18 = v18
                                                            _ = len(v6)
                                                        }
                                                        _ = v152
                                                        _ = v59
                                                    } else {
                                                        var v170 int32 = 684
                                                        fmt.Println(v4, v24)
                                                        _ = v170
                                                    }
                                                    if (v18 > 42) {
                                                        var v171 int = 32768
                                                        _ = v171
                                                    }
                                                } else {
                                                    v1 = (v2 ^ (v2 - v1))
                                                    v10 = (v10 + fi19)
                                                }
                                                _ = len(v6)
                                            }
                                            var v172 int32 = 42
                                            _ = v172
                                        } else {
                                            v24 += 1
                                        }
                                    case "a":
                                        if true {
                                            if (2 != 127) {
                                                fmt.Println(v10, fi19, fi19)
                                                fmt.Println(v4, v2, fi20)
                                                fmt.Println(v4, v4)
                                            }
                                            fmt.Println(fi20)
                                            v1 += 32767
                                            fmt.Println(v2, v2)
                                            v3 = append(v3, v10)
                                            v10 += (fi20 % (((fi19 << (((v18 | v18)) & 63))) | 1))
                                            v2 += (v1 ^ v2)
                                        } else {
                                            var v173 int32 = 0
                                            if (true && (uint32(v1) <= 42)) {
                                                if (v24 < v24) {
                                                    fmt.Println(v4, fi19, v10)
                                                    fmt.Println(fi19)
                                                }
                                                switch "a" {
                                                case "":
                                                    var v174 int16 = (v4 * v4)
                                                    _ = v174
                                                case "foo bar":
                                                    delete(v5, fi19)
                                                }
                                                if true {
                                                    switch fi20 {
                                                    case -2147483648:
                                                        if ((42 != v24) || true) {
                                                            v1 += v1
                                                            v18 += (^(v18))
                                                            switch fi19 {
                                                            case -9223372036854775808:
                                                                if (816 > 42) {
                                                                    if ((true || true) || (!((494 > uint32(v2)) && true))) {
                                                                        var v175 int8 = (-(v24))
                                                                        var v176 int = (fi20 ^ fi19)
                                                                        var v177 int32 = (v173 / (((-(v173))) | 1))
                                                                        _ = v175
                                                                        _ = v176
                                                                        _ = v177
                                                                    } else {
                                                                        delete(v5, 127)
                                                                        var v178 int = (fi19 >> (((^(v18))) & 63))
                                                                        v173 += (v173 / ((v173) | 1))
                                                                        _ = v178
                                                                    }
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                } else {
                                                                    v173 = v173
                                                                }
                                                            case 32767:
                                                                var v179 int16 = v4
                                                                _ = v179
                                                            case -948:
                                                                if true {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = -196
                                                                    }
                                                                    switch fi19 {
                                                                    case 0:
                                                                        switch int(v2) {
                                                                        case 42:
                                                                            if true {
                                                                                var v180 bool = false
                                                                                v4 += -505
                                                                                _ = v180
                                                                            } else {
                                                                                if ((-9223372036854775808 < 694) && false) {
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                    v181 := GenS23{F0: 301, F1: v18, F2: 127}
                                                                                    _ = v181
                                                                                } else {
                                                                                    v2 = 65535
                                                                                    if (911 != 989) {
                                                                                        var v182 uint = (v18 << (((v18 + 774)) & 63))
                                                                                        _ = v182
                                                                                    }
                                                                                    var v183 int = (fi19 >> (((v18 + v18)) & 63))
                                                                                    _ = v183
                                                                                }
                                                                            }
                                                                        }
                                                                        v18 = (v18 / ((689) | 1))
                                                                    case 4294967295:
                                                                        fmt.Println(v18, v24, v10)
                                                                    case 798:
                                                                        var v184 int32 = v173
                                                                        _ = v184
                                                                    default:
                                                                        var v185 uint = v18
                                                                        _ = v185
                                                                    }
                                                                    fmt.Println(v4)
                                                                }
                                                            default:
                                                                var v186 int = (fi19 >> ((v18) & 63))
                                                                _ = v186
                                                            }
                                                            v187 := GenS24{F0: 1, F1: v18, F2: 1}
                                                            _ = v187
                                                        } else {
                                                            var v188 int = (fi19 + 9223372036854775807)
                                                            fmt.Println(v24, v2, v10)
                                                            v189 := GenS25{F0: "\n\t", F1: v2, F2: true}
                                                            _ = v188
                                                            _ = v189
                                                        }
                                                    case 0:
                                                        var v190 bool = true
                                                        _ = v190
                                                    default:
                                                        fmt.Println(v173)
                                                    }
                                                    v24 = v24
                                                    fmt.Println(fi20, v2)
                                                } else {
                                                    v4 = v4
                                                    if ((((^(v2)) != v2) && false) || true) {
                                                        var v191 bool = (1.5 >= 3.14159)
                                                        v18 += v18
                                                        _ = v191
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = v10
                                                    }
                                                }
                                            }
                                            var v192 uint = 771
                                            var v193 int32 = 251
                                            var v194 int8 = (v24 + v24)
                                            _ = v173
                                            _ = v192
                                            _ = v193
                                            _ = v194
                                        }
                                    }
                                    var v195 int16 = v4
                                    delete(v5, (fi19 & (v10 / ((128) | 1))))
                                    _ = v195
                                }
                            }
                            v2 = (v2 | (v1 * v2))
                            v196 := GenS23{F0: 2, F1: v18, F2: 127}
                            v24 -= (-(v24))
                            if (!true) {
                                var v197 int = -2147483648
                                fmt.Println(v1, v1)
                                if true {
                                    switch "foo bar" {
                                    case "\n\t":
                                        _ = len(v5)
                                    case "foo bar":
                                        if false {
                                            delete(v5, v197)
                                            v5[-2147483648] = false
                                        }
                                    case "hello":
                                        delete(v5, (fi20 | (fi19 ^ fi20)))
                                    default:
                                        var v198 int8 = 100
                                        var v199 float32 = 0.1
                                        _ = v198
                                        _ = v199
                                    }
                                } else {
                                    v200 := GenS16{F0: 32768, F1: 42}
                                    _ = v200
                                }
                                _ = v197
                            }
                            _ = v196
                        }
                        _ = v24
                    }
                    switch -751 {
                    case 32768:
                        var v201 uint16 = v1
                        _ = v201
                    case 127:
                        var v202 int64 = -9223372036854775808
                        _ = v202
                    default:
                        v203 := GenS2{F0: 2147483648, F1: 2147483648, F2: "foo bar"}
                        _ = v203
                    }
                    _ = v21
                } else {
                    v204 := GenS26{F0: 0}
                    var v205 uint = (^(v18))
                    if (9223372036854775807 >= -9223372036854775808) {
                        v1 -= 1
                    } else {
                        v2 += (v2 * 0)
                    }
                    v1 = v1
                    var v206 float64 = 1e10
                    v207 := GenS27{F0: -128, F1: 2, F2: ("foo bar" + "hello")}
                    v10 = 128
                    fmt.Println(v4, v4)
                    _ = v204
                    _ = v205
                    _ = v206
                    _ = v207
                }
                var v208 int16 = v4
                _ = fi20
                _ = v208
            }
            _ = fi19
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi209 := 0; fi209 < 13; fi209++ {
            for fi210 := 0; fi210 < 1; fi210++ {
                v2 = v1
                if (true && true) {
                    switch (fi210 & fi209) {
                    case 4294967295:
                        if (!false) {
                            var v211 int16 = (v4 ^ v4)
                            var v212 uint = (^(v18))
                            _ = v211
                            _ = v212
                        } else {
                            v213 := GenS12{F0: -335, F1: 2147483647, F2: "hello"}
                            _ = len(v6)
                            if true {
                                v4 = (v4 << ((65535) & 63))
                                v214 := v5[v10]
                                _ = v214
                            }
                            _ = v213
                        }
                    default:
                        var v215 uint64 = 42
                        _ = v215
                    }
                } else {
                    var v216 float32 = (1.5 * 3.14159)
                    if (!true) {
                        var v217 uint64 = 42
                        var v218 string = "a"
                        v1 -= v1
                        _ = v217
                        _ = v218
                    }
                    v219 := GenS23{F0: 0, F1: 0, F2: 255}
                    _ = v216
                    _ = v219
                }
                _ = fi210
            }
            var v220 bool = false
            _ = fi209
            _ = v220
        }
        if (1 == (v1 >> ((v18) & 63))) {
            for fi221 := 0; fi221 < 10; fi221++ {
                v4 += (v4 ^ v4)
                _ = fi221
            }
        } else {
            v10 += 4294967295
            for fi222 := 0; fi222 < 14; fi222++ {
                fmt.Println(v18)
                var v223 int16 = v4
                switch "foo bar" {
                case "hello":
                    _ = len(v5)
                default:
                    v10 += 9223372036854775807
                }
                _ = fi222
                _ = v223
            }
        }
    }
    var v224 int32 = 1
    v1 -= 255
    var v225 uint = 735
    v1 -= v2
    var v226 int8 = 100
    var v227 float64 = (float64(float32(-2.25)) + (0.1 * 1.5))
    v5[1] = ((v224 - (-(v224))) < 980)
    switch ("a" + "hello") {
    case "a":
        fmt.Println(v224, v4, v225)
    case "hello":
        v18 -= v18
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v224:", v224)
    fmt.Println("v225:", v225)
    fmt.Println("v226:", v226)
    fmt.Println("v227:", v227)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v224
    _ = v225
    _ = v226
    _ = v227
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
