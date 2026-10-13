package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 int32; F1 int64; F2 int }
type GenS8 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS9 struct { F0 uint64; F1 uint; F2 uint8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int32; F1 uint64; F2 string }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int32; F1 int64 }
type GenS14 struct { F0 int }
func genFn0() int8 {
    v4 := []int{508, 803}
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    var v5 string = "foo bar"
    _ = v4
    _ = v5
    return 52
}

func main() {
    var v1 uint = 65535
    var v2 uint64 = 177
    var v3 bool = (!false)
    switch 9223372036854775807 {
    case 717:
        for fi6 := 0; fi6 < 4; fi6++ {
            fmt.Println(v1)
            var v7 int8 = -128
            _ = fi6
            _ = v7
        }
    default:
        v8 := []int{979, 2147483647, 32767}
        _ = v8
    }
    var v9 bool = (!false)
    if (1.5 != (-2.25 * 1e10)) {
        switch "" {
        case "a":
            v10 := GenS1{F0: "foo bar", F1: 652}
            _ = v10
        case "\n\t":
            var v11 float32 = 0.1
            _ = v11
        }
        switch "a" {
        case "a":
            var v12 uint64 = v2
            _ = v12
        case "foo bar":
            if (42 >= 780) {
                v9 = (-1 > -1)
                v13 := GenS1{F0: "foo bar", F1: 397}
                _ = v13
            } else {
                v14 := GenS2{F0: 255}
                v9 = (-9223372036854775808 <= 42)
                _ = v14
            }
        }
        if (uint16(v1) > 932) {
            for fi15 := 0; fi15 < 17; fi15++ {
                if (v3 && v3) {
                    v2 = (^(v2))
                    fmt.Println(v3)
                    if false {
                        for fi16 := 0; fi16 < 5; fi16++ {
                            switch "a" {
                            case "":
                                var v17 int16 = 42
                                _ = v17
                            case "a":
                                switch "" {
                                case "":
                                    var v18 uint64 = v2
                                    _ = v18
                                }
                            }
                            _ = fi16
                        }
                        v9 = (("" + "a") > ((("foo bar" + "a") + ("hello" + "a")) + ((("a" + "foo bar") + "a") + "a")))
                    } else {
                        if true {
                            var v19 bool = (391 < 803)
                            if true {
                                if (-537 != 32767) {
                                    var v20 uint8 = 0
                                    _ = v20
                                } else {
                                    v9 = (9223372036854775807 == -9223372036854775808)
                                    if (!v19) {
                                        v21 := GenS1{F0: "a", F1: -26}
                                        v3 = v9
                                        v22 := []int{42, (-(fi15))}
                                        _ = v21
                                        _ = v22
                                    }
                                }
                                switch "hello" {
                                case "\n\t":
                                    for fi23 := 0; fi23 < 4; fi23++ {
                                        if true {
                                            if true {
                                                var v24 bool = v9
                                                _ = v24
                                            }
                                            if v19 {
                                                v2 -= 18446744073709551615
                                            } else {
                                                var v25 float64 = (float64(((1e10 * 1e10) + (0.0 + 0.1))) / ((0.0 / 0.1) + 1.5))
                                                if ("a" <= "foo bar") {
                                                    v26 := GenS3{F0: (v2 << (((v1 ^ (v1 * (v1 ^ 255)))) & 63)), F1: 412}
                                                    var v27 string = "hello"
                                                    _ = v26
                                                    _ = v27
                                                }
                                                fmt.Println(v3)
                                                var v28 uint8 = 65
                                                if v19 {
                                                    switch "" {
                                                    case "foo bar":
                                                        if (v1 <= v1) {
                                                            switch "" {
                                                            case "\n\t":
                                                                if (-9223372036854775808 < 1) {
                                                                    var v29 int8 = -1
                                                                    v30 := []int{(fi15 >> ((v1) & 63)), (fi23 ^ fi23), fi15}
                                                                    if v19 {
                                                                        var v31 uint64 = uint64(v1)
                                                                        switch fi23 {
                                                                        case 42:
                                                                            var v32 uint16 = 255
                                                                            _ = v32
                                                                        default:
                                                                            switch 32767 {
                                                                            case -646:
                                                                                v33 := GenS3{F0: uint64(v1), F1: 1}
                                                                                _ = v33
                                                                            case 1:
                                                                                v3 = (!(true || v3))
                                                                            case 42:
                                                                                v34 := GenS4{F0: 0, F1: 65535}
                                                                                _ = v34
                                                                            default:
                                                                                var v35 uint = 4294967295
                                                                                _ = v35
                                                                            }
                                                                        }
                                                                        var v36 uint16 = 619
                                                                        _ = v31
                                                                        _ = v36
                                                                    }
                                                                    fmt.Println(v9, v28)
                                                                    _ = v29
                                                                    _ = v30
                                                                } else {
                                                                    if ((42 != 2147483647) && ("\n\t" == "a")) {
                                                                        v37 := []int{fi23}
                                                                        if v19 {
                                                                            if v19 {
                                                                                if (v28 >= (v28 % ((42) | 1))) {
                                                                                    v28 -= v28
                                                                                    var v38 string = "a"
                                                                                    _ = v38
                                                                                } else {
                                                                                    switch (fi15 % (((-(fi23))) | 1)) {
                                                                                    case -9223372036854775808:
                                                                                        fmt.Println(fi15, v1, v1)
                                                                                    case 32767:
                                                                                        v39 := GenS5{F0: "", F1: 857}
                                                                                        _ = v39
                                                                                    case 2147483647:
                                                                                        if v9 {
                                                                                            fmt.Println(v19, v25)
                                                                                        }
                                                                                    default:
                                                                                        if false {
                                                                                            v28 = v28
                                                                                            v9 = ((v28 ^ v28) >= 179)
                                                                                        }
                                                                                    }
                                                                                    var v40 int = fi23
                                                                                    v28 = (v28 * (v28 ^ v28))
                                                                                    _ = v40
                                                                                }
                                                                            }
                                                                            var v41 float32 = (1e10 + 3.14159)
                                                                            fmt.Println(v1)
                                                                            _ = v41
                                                                        } else {
                                                                            if v19 {
                                                                                if false {
                                                                                    var v42 int64 = 1
                                                                                    if v3 {
                                                                                        fmt.Println(fi23, fi23)
                                                                                        if len(v37) > 0 {
                                                                                            v37[0] = -842
                                                                                        }
                                                                                        v43 := GenS4{F0: 4294967295, F1: 2}
                                                                                        _ = v43
                                                                                    }
                                                                                    switch ("foo bar" + ("a" + ("" + "foo bar"))) {
                                                                                    case "foo bar":
                                                                                        v3 = (((fi23 & fi23) > fi23) || v3)
                                                                                    }
                                                                                    _ = v42
                                                                                }
                                                                            }
                                                                        }
                                                                        v19 = (0 >= 0)
                                                                        _ = v37
                                                                    } else {
                                                                        var v44 uint32 = 42
                                                                        if true {
                                                                            if v9 {
                                                                                if v3 {
                                                                                    if (!false) {
                                                                                        switch fi15 {
                                                                                        case -529:
                                                                                            var v45 int16 = -1
                                                                                            switch fi23 {
                                                                                            case -920:
                                                                                                v46 := []int{(fi15 | 2147483647), fi23}
                                                                                                _ = v46
                                                                                            }
                                                                                            var v47 int64 = -170
                                                                                            v19 = (!(false || true))
                                                                                            _ = v45
                                                                                            _ = v47
                                                                                        case 0:
                                                                                            var v48 uint = v1
                                                                                            _ = v48
                                                                                        case 127:
                                                                                            fmt.Println(v3)
                                                                                        }
                                                                                        if (v9 || true) {
                                                                                            v49 := []int{}
                                                                                            var v50 bool = ("a" <= "a")
                                                                                            var v51 uint = v1
                                                                                            switch "\n\t" {
                                                                                            case "hello":
                                                                                                var v52 float64 = ((float64(0.0) / v25) * v25)
                                                                                                _ = v52
                                                                                            case "\n\t":
                                                                                                var v53 uint8 = v28
                                                                                                _ = v53
                                                                                            case "a":
                                                                                                v1 = 166
                                                                                            default:
                                                                                                if len(v49) > 0 {
                                                                                                    v49[0] = (fi15 + fi15)
                                                                                                }
                                                                                            }
                                                                                            var v54 float64 = (v25 + v25)
                                                                                            fmt.Println(fi15, v51)
                                                                                            fmt.Println(fi23, v28)
                                                                                            fmt.Println(v50, v3, fi15)
                                                                                            v55 := GenS6{F0: 967, F1: 42}
                                                                                            _ = v49
                                                                                            _ = v50
                                                                                            _ = v51
                                                                                            _ = v54
                                                                                            _ = v55
                                                                                        } else {
                                                                                            if v9 {
                                                                                                v56 := []int{(-(fi23)), (fi15 + fi23), fi23}
                                                                                                _ = v56
                                                                                            } else {
                                                                                                var v57 uint16 = 42
                                                                                                v9 = (uint8(v2) < 127)
                                                                                                switch (^(fi15)) {
                                                                                                case 127:
                                                                                                    var v58 uint32 = v44
                                                                                                    _ = v58
                                                                                                case -472:
                                                                                                    if (v9 || v19) {
                                                                                                        if v3 {
                                                                                                            if (3.14159 > (0.1 - (1e10 / (3.14159 / 1.5)))) {
                                                                                                                fmt.Println(fi23)
                                                                                                                var v59 uint8 = v28
                                                                                                                v9 = (63 == 1)
                                                                                                                _ = v59
                                                                                                            }
                                                                                                        }
                                                                                                        v60 := []int{(-(fi23)), (^(fi15)), (fi23 % ((9223372036854775807) | 1)), fi23}
                                                                                                        v1 -= v1
                                                                                                        _ = v60
                                                                                                    }
                                                                                                default:
                                                                                                    v61 := []int{fi23, fi23, -1, (fi23 % (((fi15 ^ (fi23 >> (((v1 - (v1 << ((v1) & 63)))) & 63)))) | 1))}
                                                                                                    _ = v61
                                                                                                }
                                                                                                if v19 {
                                                                                                    v62 := GenS3{F0: v2, F1: 32767}
                                                                                                    v63 := GenS4{F0: (v44 + 2147483647), F1: v57}
                                                                                                    _ = v62
                                                                                                    _ = v63
                                                                                                }
                                                                                                _ = v57
                                                                                            }
                                                                                            fmt.Println(fi15, fi23, v9)
                                                                                        }
                                                                                    } else {
                                                                                        var v64 int = (fi15 | 32768)
                                                                                        var v65 uint64 = 1
                                                                                        if v9 {
                                                                                            fmt.Println(fi23, v19)
                                                                                            var v66 string = ((("\n\t" + "hello") + "hello") + "hello")
                                                                                            _ = v66
                                                                                        }
                                                                                        if (127 == -128) {
                                                                                            switch (fi23 << (((v1 | (v1 * (v1 - v1)))) & 63)) {
                                                                                            case -644:
                                                                                                v25 += float64(0.1)
                                                                                                var v67 float64 = float64(-2.25)
                                                                                                v68 := []int{(fi23 / ((fi15) | 1)), (-(fi15))}
                                                                                                _ = v67
                                                                                                _ = v68
                                                                                            default:
                                                                                                v69 := []int{fi15, (fi15 / ((v64) | 1)), 4294967295, 489}
                                                                                                _ = v69
                                                                                            }
                                                                                            v70 := GenS4{F0: (v44 + (v44 % (((v44 >> ((uint(v44)) & 63))) | 1))), F1: 65535}
                                                                                            v71 := []int{v64}
                                                                                            _ = v70
                                                                                            _ = v71
                                                                                        }
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                    }
                                                                                } else {
                                                                                    var v72 bool = true
                                                                                    switch fi23 {
                                                                                    case 32768:
                                                                                        v1 = 2
                                                                                    default:
                                                                                        var v73 string = "\n\t"
                                                                                        _ = v73
                                                                                    }
                                                                                    _ = v72
                                                                                }
                                                                            }
                                                                            v74 := GenS3{F0: 18446744073709551615, F1: int16(v1)}
                                                                            v75 := []int{fi15, (-(fi23))}
                                                                            _ = v74
                                                                            _ = v75
                                                                        }
                                                                        var v76 bool = v19
                                                                        _ = v44
                                                                        _ = v76
                                                                    }
                                                                    v77 := GenS7{F0: -1, F1: 42, F2: fi15}
                                                                    var v78 int8 = -128
                                                                    _ = v77
                                                                    _ = v78
                                                                }
                                                            case "a":
                                                                var v79 string = "\n\t"
                                                                _ = v79
                                                            case "foo bar":
                                                                v19 = v9
                                                            }
                                                            switch 4294967295 {
                                                            case 42:
                                                                if (0.1 == 3.14159) {
                                                                    switch fi15 {
                                                                    case -9223372036854775808:
                                                                        switch 736 {
                                                                        case 32767:
                                                                            v80 := GenS5{F0: "a", F1: 1}
                                                                            _ = v80
                                                                        case 128:
                                                                            var v81 string = "foo bar"
                                                                            _ = v81
                                                                        case 2147483647:
                                                                            var v82 uint = v1
                                                                            _ = v82
                                                                        default:
                                                                            switch 127 {
                                                                            case -497:
                                                                                v83 := []int{fi15, (fi23 % (((fi23 | fi15)) | 1))}
                                                                                _ = v83
                                                                            case 128:
                                                                                switch 2147483647 {
                                                                                case 128:
                                                                                    fmt.Println(v2, v1)
                                                                                }
                                                                            case 32767:
                                                                                if ((v28 % ((0) | 1)) == v28) {
                                                                                    v28 -= (^(v28))
                                                                                    v25 = 1.5
                                                                                    var v84 int = 4294967295
                                                                                    _ = v84
                                                                                }
                                                                            default:
                                                                                v28 = v28
                                                                            }
                                                                        }
                                                                    case 2147483647:
                                                                        var v85 float64 = (v25 * -2.25)
                                                                        _ = v85
                                                                    case 128:
                                                                        v9 = (true && true)
                                                                    default:
                                                                        v9 = (!false)
                                                                    }
                                                                } else {
                                                                    switch fi15 {
                                                                    case 0:
                                                                        v86 := []int{(fi23 ^ fi15), fi15, int(v28)}
                                                                        _ = v86
                                                                    default:
                                                                        var v87 uint16 = 1
                                                                        _ = v87
                                                                    }
                                                                }
                                                            case -2147483648:
                                                                v88 := []int{(fi15 << ((744) & 63))}
                                                                _ = v88
                                                            default:
                                                                v89 := GenS4{F0: 0, F1: uint16(v28)}
                                                                _ = v89
                                                            }
                                                            v90 := GenS6{F0: -470, F1: 0}
                                                            _ = v90
                                                        } else {
                                                            switch ("foo bar" + ("hello" + "foo bar")) {
                                                            case "\n\t":
                                                                var v91 int8 = 127
                                                                _ = v91
                                                            case "foo bar":
                                                                v3 = (v2 >= v2)
                                                            case "":
                                                                var v92 uint8 = 127
                                                                _ = v92
                                                            }
                                                            v25 += v25
                                                            if (0 != 1) {
                                                                v19 = (!v3)
                                                            }
                                                        }
                                                        v93 := []int{2147483647, fi15}
                                                        _ = v93
                                                    }
                                                    var v94 int64 = 9223372036854775807
                                                    switch 32767 {
                                                    case 4294967295:
                                                        var v95 int8 = int8(v1)
                                                        _ = v95
                                                    }
                                                    _ = v94
                                                } else {
                                                    v96 := []int{(fi23 - fi15), 32768}
                                                    if len(v96) >= 2 {
                                                        v96 = v96[1:len(v96)]
                                                    }
                                                    v9 = (!((1e10 - v25) != (v25 * v25)))
                                                    _ = v96
                                                }
                                                _ = v25
                                                _ = v28
                                            }
                                        }
                                        if v3 {
                                            if (v2 == (v2 >> (((^(v1))) & 63))) {
                                                v97 := []int{fi15}
                                                _ = v97
                                            }
                                            if v3 {
                                                fmt.Println(v9, v1, fi15)
                                                v1 = 255
                                            }
                                        } else {
                                            var v98 int = (fi15 << ((v1) & 63))
                                            _ = v98
                                        }
                                        _ = fi23
                                    }
                                case "":
                                    switch ("foo bar" + ("\n\t" + "foo bar")) {
                                    case "a":
                                        switch fi15 {
                                        case -9223372036854775808:
                                            v99 := GenS8{F0: 65535, F1: 88, F2: 0}
                                            _ = v99
                                        default:
                                            var v100 float64 = float64(0.0)
                                            _ = v100
                                        }
                                    case "hello":
                                        fmt.Println(v1, v3)
                                    case "\n\t":
                                        if (("" < "foo bar") && (false || v3)) {
                                            v9 = (!(v3 && (float32(-2.25) == 1.5)))
                                        }
                                    default:
                                        v101 := []int{127, (fi15 ^ 127), fi15, -1}
                                        _ = v101
                                    }
                                }
                            }
                            v102 := []int{}
                            _ = v102
                            _ = v19
                        } else {
                            fmt.Println(v1, v1, v2)
                        }
                        switch "\n\t" {
                        case "a":
                            v1 += 2
                        case "foo bar":
                            for fi103 := 0; fi103 < 18; fi103++ {
                                var v104 uint16 = 1
                                switch "foo bar" {
                                case "foo bar":
                                    if (!(int64(fi15) == 0)) {
                                        v2 -= v2
                                        var v105 uint16 = v104
                                        v3 = v9
                                        var v106 float32 = (float32((1.5 * 0.1)) - 1e10)
                                        fmt.Println(fi103, v1, v105)
                                        _ = v105
                                        _ = v106
                                    }
                                case "hello":
                                    fmt.Println(v1, v3)
                                case "a":
                                    v9 = true
                                default:
                                    var v107 int64 = 399
                                    _ = v107
                                }
                                _ = fi103
                                _ = v104
                            }
                        case "\n\t":
                            if (v9 || v9) {
                                var v108 float32 = float32(0.0)
                                v2 = (v2 - 2)
                                _ = v108
                            } else {
                                v109 := GenS9{F0: (v2 | 42), F1: v1, F2: 0}
                                if v3 {
                                    switch 2147483647 {
                                    case 128:
                                        switch (fi15 - (fi15 + fi15)) {
                                        case 9223372036854775807:
                                            v110 := GenS5{F0: "\n\t", F1: 42}
                                            _ = v110
                                        case 127:
                                            fmt.Println(v3, v3, v2)
                                        }
                                    case -9223372036854775808:
                                        var v111 int64 = 0
                                        _ = v111
                                    case 127:
                                        var v112 int = 4294967295
                                        _ = v112
                                    default:
                                        v113 := []int{(fi15 / (((^(fi15))) | 1)), fi15, (fi15 | 9223372036854775807)}
                                        _ = v113
                                    }
                                } else {
                                    if (!(v3 && v3)) {
                                        v114 := []int{1}
                                        v115 := GenS10{F0: 241}
                                        v3 = v9
                                        _ = v114
                                        _ = v115
                                    }
                                }
                                _ = v109
                            }
                        default:
                            v116 := []int{}
                            _ = v116
                        }
                    }
                } else {
                    for fi117 := 0; fi117 < 7; fi117++ {
                        v118 := []int{fi15, (fi117 + (fi117 << ((2) & 63))), fi15, 32768}
                        if ((1e10 / 1.0) == -2.25) {
                            v3 = v3
                            v1 = v1
                            v3 = v9
                        } else {
                            v119 := GenS11{F0: 1, F1: v2, F2: ("hello" + "hello")}
                            v120 := GenS12{F0: 100}
                            _ = v119
                            _ = v120
                        }
                        var v121 string = ("" + ("\n\t" + (("\n\t" + "") + ("hello" + "foo bar"))))
                        v122 := GenS3{F0: v2, F1: 1}
                        if (fi15 >= 32768) {
                            v121 = v121
                            v118 = append(v118, fi15)
                        }
                        _ = fi117
                        _ = v118
                        _ = v121
                        _ = v122
                    }
                    v123 := GenS10{F0: v1}
                    var v124 int16 = 670
                    _ = v123
                    _ = v124
                }
                fmt.Println(v9, v9)
                var v125 int16 = -824
                _ = fi15
                _ = v125
            }
            fmt.Println(v3)
        }
    } else {
        v9 = false
        for fi126 := 0; fi126 < 18; fi126++ {
            v127 := GenS12{F0: -86}
            var v128 int = (fi126 | fi126)
            _ = fi126
            _ = v127
            _ = v128
        }
    }
    if (!(0.1 <= ((3.14159 + 1.5) / 0.1))) {
        if (v9 && (-1 < -492)) {
            v129 := GenS13{F0: 0, F1: -466}
            _ = v129
        }
        switch "" {
        case "foo bar":
            for fi130 := 0; fi130 < 15; fi130++ {
                var v131 uint = (v1 - v1)
                _ = fi130
                _ = v131
            }
        default:
            var v132 int = 892
            _ = v132
        }
    } else {
        if (32767 > 657) {
            for fi133 := 0; fi133 < 10; fi133++ {
                var v134 int32 = -2147483648
                _ = fi133
                _ = v134
            }
            if (-9223372036854775808 > 9223372036854775807) {
                var v135 uint64 = v2
                v136 := []int{-391, 32767, 32768}
                var v137 int32 = 1
                _ = v135
                _ = v136
                _ = v137
            }
        } else {
            var v138 int64 = 9223372036854775807
            _ = v138
        }
        v3 = true
    }
    var v139 uint = (^(v1))
    v140 := []int{4294967295, -362, 9223372036854775807, 32767}
    for fi141 := 0; fi141 < 18; fi141++ {
        var v142 int8 = 0
        var v143 uint32 = uint32(v2)
        if (true && v3) {
            var v144 int16 = int16(v139)
            v3 = (!v9)
            var v145 int = fi141
            _ = v144
            _ = v145
        }
        _ = fi141
        _ = v142
        _ = v143
    }
    for fi146 := 0; fi146 < 7; fi146++ {
        v9 = (1.5 == 1e10)
        for fi147 := 0; fi147 < 18; fi147++ {
            var v148 uint16 = 360
            var v149 int32 = 0
            if (v2 == v2) {
                if false {
                    var v150 int32 = -1
                    if len(v140) >= 2 {
                        v140 = v140[1:len(v140)]
                    }
                    var v151 uint32 = 778
                    var v152 int8 = 42
                    v140 = append(v140, (fi146 << ((v139) & 63)))
                    _ = v150
                    _ = v151
                    _ = v152
                }
                v139 += v139
            }
            _ = fi147
            _ = v148
            _ = v149
        }
        for fi153 := 0; fi153 < 1; fi153++ {
            fmt.Println(fi146, v1, v139)
            switch int(v2) {
            case 42:
                var v154 uint32 = 951
                _ = v154
            case -498:
                if true {
                    switch "" {
                    case "hello":
                        v1 += (^(v139))
                        v140 = append(v140, fi146)
                        fmt.Println(fi146, v3)
                    case "a":
                        fmt.Println(fi153)
                    }
                }
            }
            _ = fi153
        }
        _ = fi146
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v139, v2, v2)
    case "hello":
        v3 = v3
    case "foo bar":
        if (1.5 == 3.14159) {
            v140 = append(v140, -2147483648)
            v140 = append(v140, 2147483647)
            v3 = ("a" > "hello")
        } else {
            var v155 int = 32768
            for fi156 := 0; fi156 < 11; fi156++ {
                for fi157 := 0; fi157 < 1; fi157++ {
                    if (v2 >= 18446744073709551615) {
                        v139 += v1
                    }
                    _ = fi157
                }
                _ = fi156
            }
            _ = v155
        }
    }
    var v158 uint8 = 255
    if len(v140) > 0 {
        v140[0] = -2147483648
    }
    if (!v9) {
        var v159 string = "hello"
        v160 := GenS13{F0: -1, F1: 47}
        var v161 bool = (0 >= 0)
        _ = v159
        _ = v160
        _ = v161
    } else {
        fmt.Println(v158)
        var v162 int8 = 1
        _ = v162
    }
    var v163 int16 = -32768
    v3 = (!v9)
    v140 = append(v140, 230)
    for fi164 := 0; fi164 < 7; fi164++ {
        _ = len(v140)
        _ = fi164
    }
    switch 42 {
    case 32768:
        for fi165 := 0; fi165 < 1; fi165++ {
            v140 = append(v140, fi165)
            if len(v140) > 0 {
                v140[0] = 42
            }
            for fi166 := 0; fi166 < 6; fi166++ {
                if false {
                    v1 = v139
                    switch (-(fi166)) {
                    case 88:
                        v9 = ((v158 != (^(v158))) || (3.14159 >= (((3.14159 / 3.14159) / 0.1) * (1e10 + (float32(-2.25) + float32(3.14159))))))
                    case -9223372036854775808:
                        v163 += (v163 % ((v163) | 1))
                    case 2147483647:
                        var v167 uint = v1
                        _ = v167
                    }
                    fmt.Println(v3)
                    switch ("a" + ("\n\t" + "hello")) {
                    case "":
                        fmt.Println(v163)
                    case "a":
                        v9 = true
                    }
                }
                var v168 uint32 = 2147483648
                v169 := GenS10{F0: v1}
                _ = fi166
                _ = v168
                _ = v169
            }
            _ = fi165
        }
    case 42:
        switch 615 {
        case 2147483647:
            for fi170 := 0; fi170 < 8; fi170++ {
                v9 = (true || (v9 || true))
                var v171 uint16 = 2
                if ((uint64(v158) > (v2 * v2)) || (v3 && v3)) {
                    if (!(-128 > genFn0())) {
                        v2 = v2
                        var v172 int16 = v163
                        v171 = 65535
                        _ = v172
                    }
                    v173 := GenS14{F0: (fi170 ^ 926)}
                    _ = v173
                }
                _ = fi170
                _ = v171
            }
        }
    case -2147483648:
        var v174 float32 = 1e10
        _ = v174
    }
    v3 = (true && (!v9))
    fmt.Println("v1:", v1)
    fmt.Println("v139:", v139)
    fmt.Println("slicelen v140:", len(v140), v140)
    fmt.Println("v158:", v158)
    fmt.Println("v163:", v163)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v139
    _ = v140
    _ = v158
    _ = v163
    _ = v2
    _ = v3
    _ = v9
}
