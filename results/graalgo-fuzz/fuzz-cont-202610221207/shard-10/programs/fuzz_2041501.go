package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS9 struct { F0 uint64; F1 int; F2 int }
type GenS10 struct { F0 bool; F1 int8; F2 uint }

func main() {
    var v1 int8 = -1
    var v2 uint = 1
    var v3 float32 = -2.25
    v4 := make(map[string]bool)
    fmt.Println(v2, v2, v3)
    if true {
        v3 += v3
        fmt.Println(v2, v2, v1)
    } else {
        v5 := GenS1{F0: 127, F1: 42, F2: v1}
        if false {
            fmt.Println(v3)
            v6 := GenS1{F0: 128, F1: 514, F2: (v1 >> ((v2) & 63))}
            if (0 < 9223372036854775807) {
                if ((2147483647 >= 2) && true) {
                    if (0 != 177) {
                        if ((-401 != -2147483648) || false) {
                            fmt.Println(v3, v3)
                            var v7 float64 = float64(((0.1 * 1.5) + -2.25))
                            v8 := GenS1{F0: 97, F1: 255, F2: int8(v2)}
                            _ = v7
                            _ = v8
                        }
                    } else {
                        _ = len(v4)
                    }
                } else {
                    v9 := []int{128}
                    v10 := GenS2{F0: v1}
                    var v11 int16 = -1
                    _ = v10
                    _ = v11
                    _ = v9
                }
                if (65535 < 255) {
                    v12 := v4[("hello" + "hello")]
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 7; fi13++ {
                    var v14 int16 = 32767
                    for fi15 := 0; fi15 < 15; fi15++ {
                        var v16 uint16 = 0
                        var v17 bool = (false || (v2 <= (v2 & (v2 - 749))))
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    switch -121 {
                    case 4294967295:
                        for fi18 := 0; fi18 < 13; fi18++ {
                            var v19 uint32 = 2147483648
                            var v20 int = (-(fi13))
                            if false {
                                var v21 int16 = (v14 + int16(v1))
                                if (false && true) {
                                    var v22 uint32 = 991
                                    switch "a" {
                                    case "hello":
                                        var v23 int8 = (v1 | (v1 / ((-1) | 1)))
                                        _ = v23
                                    case "foo bar":
                                        v24 := GenS1{F0: 76, F1: 1, F2: v1}
                                        _ = v24
                                    case "\n\t":
                                        if (uint64(v2) > 42) {
                                            if true {
                                                v25 := []int{(fi18 - 128), (fi18 >> ((753) & 63)), 0, v20}
                                                if true {
                                                    v26 := GenS3{F0: v19, F1: 2}
                                                    _ = v26
                                                } else {
                                                    v27 := GenS3{F0: v22, F1: 793}
                                                    fmt.Println(v1, v20)
                                                    var v28 int32 = 17
                                                    v14 = -32768
                                                    var v29 uint = (v2 / ((v2) | 1))
                                                    if (128 >= 128) {
                                                        if (false || true) {
                                                            v30 := GenS4{F0: 1, F1: v22, F2: (v20 - 0)}
                                                            _ = v30
                                                        } else {
                                                            fmt.Println(v19)
                                                        }
                                                        v31 := GenS3{F0: 42, F1: 88}
                                                        switch 671 {
                                                        case -230:
                                                            v32 := GenS1{F0: 128, F1: 308, F2: 100}
                                                            _ = v32
                                                        case 127:
                                                            delete(v4, "\n\t")
                                                            var v33 int8 = v1
                                                            if false {
                                                                v1 += v33
                                                            }
                                                            v29 -= v2
                                                            _ = v33
                                                        }
                                                        _ = v31
                                                    }
                                                    v34 := GenS2{F0: -1}
                                                    v14 -= (v21 % (((v14 / ((v21) | 1))) | 1))
                                                    _ = v27
                                                    _ = v28
                                                    _ = v29
                                                    _ = v34
                                                }
                                                _ = v25
                                            }
                                            var v35 bool = (true && (false || (42 <= 877)))
                                            _ = v35
                                        } else {
                                            var v36 uint = (v2 % (((v2 % ((4294967295) | 1))) | 1))
                                            v22 += v19
                                            v37 := []int{-9223372036854775808}
                                            _ = v36
                                            _ = v37
                                        }
                                    default:
                                        v38 := GenS2{F0: v1}
                                        _ = v38
                                    }
                                    switch (("hello" + ("\n\t" + (("a" + "") + "foo bar"))) + ("hello" + "")) {
                                    case "hello":
                                        if (!(-1 == v1)) {
                                            v39 := GenS5{F0: 10, F1: v2}
                                            v19 += v22
                                            _ = v39
                                        } else {
                                            v40 := []int{1, (v20 << (((v2 | (v2 * v2))) & 63)), (v20 ^ (-(fi18)))}
                                            if true {
                                                _ = len(v40)
                                                v19 += 2147483647
                                                var v41 int8 = (v1 ^ (-(v1)))
                                                _ = v41
                                            } else {
                                                v4[("\n\t" + "foo bar")] = (false || ((v1 * 0) >= (-(v1))))
                                            }
                                            var v42 bool = true
                                            v43 := GenS6{F0: ("hello" + (("hello" + ("a" + "\n\t")) + "a")), F1: 42}
                                            if len(v40) > 0 {
                                                v40[0] = (-(fi13))
                                            }
                                            _ = v40
                                            _ = v42
                                            _ = v43
                                        }
                                    case "":
                                        fmt.Println(v22, v20, v20)
                                    }
                                    _ = v22
                                }
                                var v44 int64 = -1
                                _ = v21
                                _ = v44
                            }
                            if false {
                                if false {
                                    switch v20 {
                                    case -9223372036854775808:
                                        fmt.Println(v14)
                                    default:
                                        switch "" {
                                        case "":
                                            if (v19 > v19) {
                                                if true {
                                                    if true {
                                                        v14 = (v14 % ((v14) | 1))
                                                    }
                                                }
                                                var v45 int32 = 42
                                                _ = v45
                                            }
                                        case "a":
                                            v46 := []int{fi13, (v20 >> ((v2) & 63)), v20, fi18}
                                            _ = v46
                                        default:
                                            switch (fi18 << (((v2 >> (((v2 / ((v2) | 1))) & 63))) & 63)) {
                                            case -274:
                                                v47 := []int{9223372036854775807, v20, (fi18 & v20), fi13}
                                                _ = v47
                                            }
                                        }
                                    }
                                    fmt.Println(v3, fi18)
                                    var v48 int8 = (v1 >> (((v2 | v2)) & 63))
                                    _ = v48
                                }
                            }
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    case 2147483647:
                        v4["hello"] = false
                    case 62:
                        switch (fi13 % (((fi13 ^ fi13)) | 1)) {
                        case 438:
                            v49 := v4["hello"]
                            _ = v49
                        default:
                            var v50 uint32 = 0
                            _ = v50
                        }
                    default:
                        switch fi13 {
                        case -355:
                            var v51 int64 = 54
                            _ = v51
                        case 42:
                            if true {
                                if true {
                                    v1 += (v1 - (^(v1)))
                                } else {
                                    v52 := []int{fi13, (fi13 | -2147483648), fi13}
                                    if false {
                                        for fi53 := 0; fi53 < 9; fi53++ {
                                            if false {
                                                var v54 int16 = (v14 % ((v14) | 1))
                                                v1 = v1
                                                v3 -= (v3 - float32(0.1))
                                                v1 = (v1 % (((v1 * int8(v54))) | 1))
                                                if (-1 != (v14 << ((v2) & 63))) {
                                                    v3 -= v3
                                                    var v55 uint64 = 654
                                                    fmt.Println(fi53)
                                                    _ = v55
                                                } else {
                                                    var v56 float32 = (0.0 / 1e10)
                                                    v1 -= (-(v1))
                                                    var v57 int32 = -94
                                                    v57 += 729
                                                    _ = v56
                                                    _ = v57
                                                }
                                                if false {
                                                    v4["foo bar"] = true
                                                    fmt.Println(fi53, v2, v2)
                                                    _ = len(v4)
                                                }
                                                var v58 float64 = (3.14159 + 1.5)
                                                if (!true) {
                                                    v3 = (v3 / (1.5 * 1e10))
                                                    if (false || false) {
                                                        if false {
                                                            switch "\n\t" {
                                                            case "":
                                                                _ = len(v4)
                                                            case "a":
                                                                fmt.Println(fi13, fi13, fi53)
                                                            }
                                                            v14 -= v14
                                                            v1 -= v1
                                                        } else {
                                                            v14 += v54
                                                            delete(v4, "a")
                                                            if (!false) {
                                                                var v59 string = "hello"
                                                                _ = v59
                                                            }
                                                        }
                                                        v14 += (v14 % ((v14) | 1))
                                                        fmt.Println(v2, v54)
                                                    } else {
                                                        switch ("\n\t" + ("\n\t" + (((("hello" + ("" + "a")) + "") + "foo bar") + ""))) {
                                                        case "foo bar":
                                                            v60 := GenS4{F0: 966, F1: 65535, F2: fi53}
                                                            _ = v60
                                                        case "a":
                                                            v1 = (v1 * 1)
                                                        }
                                                        v14 = (v14 * 32767)
                                                    }
                                                } else {
                                                    if false {
                                                        var v61 uint8 = 128
                                                        v62 := GenS7{F0: 1}
                                                        v1 += v1
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    v58 = 1e10
                                                }
                                                _ = v54
                                                _ = v58
                                            } else {
                                                var v63 float64 = (float64(v3) - ((1e10 + 1e10) / 1.0))
                                                v1 = (v1 >> ((v2) & 63))
                                                fmt.Println(fi53, v63)
                                                var v64 uint = (v2 | (^(v2)))
                                                switch 128 {
                                                case 4294967295:
                                                    switch ("hello" + "a") {
                                                    case "\n\t":
                                                        var v65 int = 32767
                                                        if (0 < (v1 | v1)) {
                                                            if false {
                                                                if (2147483647 <= -933) {
                                                                    var v66 int64 = 0
                                                                    v63 -= v63
                                                                    _ = v66
                                                                }
                                                            }
                                                            v52 = append(v52, fi13)
                                                            var v67 int8 = (v1 + v1)
                                                            if (-935 < 42) {
                                                                var v68 int = fi13
                                                                _ = v68
                                                            } else {
                                                                v69 := GenS6{F0: "\n\t", F1: 0}
                                                                if (!(float64(1.5) != float64(v3))) {
                                                                    fmt.Println(v65)
                                                                    if (-567 > -9223372036854775808) {
                                                                        if (!("a" <= "")) {
                                                                            var v70 uint = v64
                                                                            v64 += v2
                                                                            var v71 string = ""
                                                                            _ = v70
                                                                            _ = v71
                                                                        }
                                                                        if (877 > 0) {
                                                                            var v72 uint8 = 255
                                                                            v73 := GenS3{F0: 875, F1: 42}
                                                                            _ = v72
                                                                            _ = v73
                                                                        } else {
                                                                            var v74 uint = (v64 >> (((v2 / (((v64 - v2)) | 1))) & 63))
                                                                            v75 := GenS8{F0: -863, F1: 42, F2: 266}
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                        fmt.Println(v65)
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        fmt.Println(v2, v63)
                                                                    case "a":
                                                                        v65 -= fi53
                                                                    case "hello":
                                                                        fmt.Println(fi53, fi53, v67)
                                                                    default:
                                                                        if true {
                                                                            v76 := GenS2{F0: v67}
                                                                            if true {
                                                                                var v77 uint64 = 1
                                                                                _ = len(v4)
                                                                                v77 = (v77 | (v77 << ((v64) & 63)))
                                                                                var v78 uint32 = 586
                                                                                if (((float32(-2.25) * v3) + v3) < (float32(float64(v3)) * ((v3 - v3) + float32((3.14159 * v63))))) {
                                                                                    if (false || true) {
                                                                                        switch fi13 {
                                                                                        case 127:
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                v77 = (v77 - 42)
                                                                                                fmt.Println(v2, fi13, v65)
                                                                                            case "\n\t":
                                                                                                v52 = append(v52, 1)
                                                                                            default:
                                                                                                v79 := GenS1{F0: 250, F1: 397, F2: (v67 + (v1 ^ (v1 * v67)))}
                                                                                                _ = v79
                                                                                            }
                                                                                        case -9223372036854775808:
                                                                                            switch int(v77) {
                                                                                            case 4294967295:
                                                                                                fmt.Println(v77, fi53, fi13)
                                                                                            }
                                                                                        case 0:
                                                                                            var v80 uint32 = v78
                                                                                            _ = v80
                                                                                        }
                                                                                    } else {
                                                                                        if ((!true) || false) {
                                                                                            var v81 uint16 = uint16(v77)
                                                                                            _ = v81
                                                                                        } else {
                                                                                            v3 = float32(v63)
                                                                                            _ = len(v52)
                                                                                            var v82 uint64 = 889
                                                                                            v52 = append(v52, fi53)
                                                                                            _ = v82
                                                                                        }
                                                                                        if false {
                                                                                            var v83 float64 = ((float64((v3 / ((v3 * (v3 / v3)) * v3))) - v63) * float64(v3))
                                                                                            switch (fi53 * (v65 ^ fi13)) {
                                                                                            case 2147483647:
                                                                                                v84 := GenS7{F0: v77}
                                                                                                _ = v84
                                                                                            case 1:
                                                                                                switch "hello" {
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v77)
                                                                                                case "a":
                                                                                                    v83 -= (v83 / (v63 + 3.14159))
                                                                                                default:
                                                                                                    v65 = (fi53 >> ((uint(v78)) & 63))
                                                                                                    v85 := GenS6{F0: "foo bar", F1: -2147483648}
                                                                                                    _ = v85
                                                                                                }
                                                                                            }
                                                                                            var v86 uint32 = 450
                                                                                            _ = v83
                                                                                            _ = v86
                                                                                        }
                                                                                        fmt.Println(fi13)
                                                                                        var v87 uint64 = (v77 >> ((uint(v78)) & 63))
                                                                                        _ = v87
                                                                                    }
                                                                                }
                                                                                _ = v77
                                                                                _ = v78
                                                                            } else {
                                                                                switch -17 {
                                                                                case -1:
                                                                                    var v88 uint8 = 128
                                                                                    _ = v88
                                                                                default:
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        v89 := GenS7{F0: 940}
                                                                                        var v90 bool = false
                                                                                        var v91 int64 = 42
                                                                                        v2 += 22
                                                                                        v14 += (v14 + v14)
                                                                                        _ = v89
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    case "foo bar":
                                                                                        v14 = int16(fi13)
                                                                                    }
                                                                                }
                                                                                delete(v4, "foo bar")
                                                                                v64 = v64
                                                                            }
                                                                            _ = v76
                                                                        } else {
                                                                            if true {
                                                                                switch (fi13 - (v65 ^ 0)) {
                                                                                case 32768:
                                                                                    var v92 uint = (v2 / (((v2 % ((4294967295) | 1))) | 1))
                                                                                    _ = v92
                                                                                case 128:
                                                                                    var v93 int8 = v67
                                                                                    _ = v93
                                                                                case 4294967295:
                                                                                    v63 = 3.14159
                                                                                default:
                                                                                    switch fi53 {
                                                                                    case 32767:
                                                                                        v94 := GenS6{F0: ("" + "foo bar"), F1: 42}
                                                                                        _ = v94
                                                                                    default:
                                                                                        v95 := GenS9{F0: 0, F1: v65, F2: -9223372036854775808}
                                                                                        _ = v95
                                                                                    }
                                                                                }
                                                                                v3 = v3
                                                                                if (!(!((v3 + v3) > 1.5))) {
                                                                                    if ((v64 / ((v2) | 1)) < v64) {
                                                                                        if len(v52) >= 2 {
                                                                                            v52 = v52[1:len(v52)]
                                                                                        }
                                                                                        var v96 uint8 = 128
                                                                                        fmt.Println(fi13, v14)
                                                                                        _ = v96
                                                                                    } else {
                                                                                        v97 := GenS6{F0: "hello", F1: 1}
                                                                                        _ = v97
                                                                                    }
                                                                                    var v98 float64 = -2.25
                                                                                    _ = v98
                                                                                } else {
                                                                                    var v99 int64 = -9223372036854775808
                                                                                    _ = v99
                                                                                }
                                                                            }
                                                                            if true {
                                                                                fmt.Println(v63, v1, v14)
                                                                                v100 := GenS2{F0: 100}
                                                                                _ = v100
                                                                            }
                                                                            v1 -= v1
                                                                        }
                                                                    }
                                                                } else {
                                                                    if (v1 >= v1) {
                                                                        v2 += (v64 % (((v64 << ((516) & 63))) | 1))
                                                                    }
                                                                }
                                                                if (0.1 < (v3 + (3.14159 / (((1e10 * ((1.5 / 1e10) + 1e10)) / float32(v63)) * float32(float64(v3)))))) {
                                                                    v4[""] = ((737 < 516) && (0 > 0))
                                                                    v63 += v63
                                                                } else {
                                                                    if false {
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            var v101 bool = (127 >= 128)
                                                                            fmt.Println(v101)
                                                                            _ = v101
                                                                        case "":
                                                                            fmt.Println(v64)
                                                                            fmt.Println(fi53)
                                                                        default:
                                                                            switch ("\n\t" + ("hello" + ("\n\t" + "a"))) {
                                                                            case "foo bar":
                                                                                v102 := GenS4{F0: 611, F1: 711, F2: fi53}
                                                                                _ = v102
                                                                            case "a":
                                                                                var v103 int = (fi13 | fi13)
                                                                                _ = v103
                                                                            default:
                                                                                fmt.Println(v3, v2)
                                                                            }
                                                                        }
                                                                        var v104 uint16 = 42
                                                                        var v105 uint64 = 4294967295
                                                                        _ = v104
                                                                        _ = v105
                                                                    }
                                                                    if (!("foo bar" > ("" + (("a" + "\n\t") + "\n\t")))) {
                                                                        v3 += (1.5 / v3)
                                                                    } else {
                                                                        var v106 bool = ((true || false) || (-350 < 990))
                                                                        var v107 uint32 = 0
                                                                        if v106 {
                                                                            if v106 {
                                                                                v108 := GenS4{F0: v107, F1: 667, F2: fi13}
                                                                                _ = v108
                                                                            }
                                                                            fmt.Println(v106, v106)
                                                                            v4[("\n\t" + "")] = v106
                                                                        }
                                                                        _ = v106
                                                                        _ = v107
                                                                    }
                                                                }
                                                                v14 -= v14
                                                                var v109 string = (("foo bar" + "a") + "hello")
                                                                _ = v109
                                                                _ = v69
                                                            }
                                                            fmt.Println(fi13, v14)
                                                            _ = v67
                                                        }
                                                        var v110 bool = false
                                                        _ = v110
                                                        _ = v65
                                                    case "hello":
                                                        if true {
                                                            v63 -= v63
                                                            v3 = -2.25
                                                        }
                                                    case "foo bar":
                                                        if ((v2 + 18446744073709551615) == v2) {
                                                            if len(v52) >= 2 {
                                                                v52 = v52[1:len(v52)]
                                                            }
                                                            v2 = 65535
                                                            v111 := GenS10{F0: false, F1: (v1 & (v1 >> ((v2) & 63))), F2: 0}
                                                            _ = len(v52)
                                                            _ = v111
                                                        }
                                                    }
                                                case 32767:
                                                    var v112 int = fi13
                                                    _ = v112
                                                default:
                                                    if true {
                                                        var v113 bool = true
                                                        var v114 uint = 2
                                                        _ = v113
                                                        _ = v114
                                                    } else {
                                                        v2 = 42
                                                        v2 = (v2 % ((v2) | 1))
                                                        var v115 uint32 = 65535
                                                        var v116 int8 = v1
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            v4["\n\t"] = ((((1.5 / v63) / (float64(0.0) + 1.5)) * 3.14159) > v63)
                                                        }
                                                        _ = v115
                                                        _ = v116
                                                    }
                                                }
                                                _ = v63
                                                _ = v64
                                            }
                                            _ = fi53
                                        }
                                        var v117 float32 = (v3 * (((v3 + v3) / v3) - -2.25))
                                        if (0 > 2) {
                                            if ("" >= "a") {
                                                var v118 uint8 = 128
                                                var v119 int16 = (v14 << ((255) & 63))
                                                v52 = append(v52, fi13)
                                                _ = v118
                                                _ = v119
                                            }
                                            v3 = float32((0.0 * 0.0))
                                            v1 = v1
                                        } else {
                                            v120 := GenS10{F0: (fi13 >= fi13), F1: (v1 | (v1 ^ (-(v1)))), F2: 964}
                                            var v121 uint32 = 65535
                                            _ = v120
                                            _ = v121
                                        }
                                        _ = v117
                                    } else {
                                        v122 := GenS10{F0: false, F1: (-(v1)), F2: 1}
                                        _ = v122
                                    }
                                    _ = v52
                                }
                            } else {
                                var v123 int = fi13
                                _ = v123
                            }
                        }
                    }
                    _ = fi13
                    _ = v14
                }
            }
            _ = v6
        } else {
            fmt.Println(v3, v2)
            if (int64(v1) <= 0) {
                var v124 uint16 = 249
                _ = v124
            } else {
                var v125 uint = (v2 + (v2 ^ (v2 - 902)))
                fmt.Println(v3)
                _ = v125
            }
        }
        _ = v5
    }
    for fi126 := 0; fi126 < 14; fi126++ {
        v2 += (^(v2))
        var v127 float32 = v3
        _ = fi126
        _ = v127
    }
    if true {
        var v128 string = "foo bar"
        _ = v128
    }
    var v129 float32 = 3.14159
    if true {
        v130 := []int{-844, 32767, -93, 32768}
        _ = v130
    }
    v4[(((("foo bar" + "hello") + "a") + ("foo bar" + ("hello" + ""))) + ("" + "\n\t"))] = (false && true)
    var v131 float32 = ((float32(-2.25) - float32(-2.25)) / 1.0)
    if false {
        var v132 int64 = 1
        delete(v4, "\n\t")
        if true {
            v1 -= (^(v1))
        } else {
            _ = len(v4)
        }
        _ = v132
    }
    if (!(!(!false))) {
        for fi133 := 0; fi133 < 12; fi133++ {
            if (!false) {
                if (false && (!(!false))) {
                    delete(v4, "\n\t")
                    if false {
                        switch (fi133 * fi133) {
                        case 9223372036854775807:
                            if false {
                                v131 -= v129
                                v131 = float32(0.1)
                            } else {
                                if (1e10 != 0.1) {
                                    v134 := GenS7{F0: 18446744073709551615}
                                    if true {
                                        var v135 bool = true
                                        v129 += v3
                                        _ = v135
                                    }
                                    for fi136 := 0; fi136 < 2; fi136++ {
                                        var v137 bool = false
                                        v138 := GenS6{F0: "hello", F1: 2147483647}
                                        _ = fi136
                                        _ = v137
                                        _ = v138
                                    }
                                    _ = v134
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(v129)
                                    }
                                    var v139 int64 = 0
                                    delete(v4, ("foo bar" + "\n\t"))
                                    _ = v139
                                }
                                if true {
                                    v131 -= ((v3 - 3.14159) - v129)
                                    var v140 int32 = -1
                                    _ = v140
                                }
                            }
                        }
                        v141 := GenS8{F0: int64(v2), F1: 9, F2: 190}
                        v142 := GenS8{F0: 0, F1: 65535, F2: 65535}
                        _ = v141
                        _ = v142
                    } else {
                        v3 += (v129 / (1e10 * v3))
                    }
                    v3 = ((1e10 / v131) * 0.0)
                } else {
                    fmt.Println(v1, v129, v2)
                }
                var v143 bool = ("" != ("\n\t" + ("hello" + "a")))
                v144 := []int{}
                _ = v143
                _ = v144
            }
            v145 := []int{}
            _ = fi133
            _ = v145
        }
        for fi146 := 0; fi146 < 7; fi146++ {
            var v147 uint32 = 4294967295
            _ = fi146
            _ = v147
        }
    } else {
        v1 -= (v1 + v1)
        var v148 uint64 = 0
        _ = v148
    }
    fmt.Println("v1:", v1)
    fmt.Println("v129:", v129)
    fmt.Println("v131:", v131)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v129
    _ = v131
    _ = v2
    _ = v3
    _ = v4
}
