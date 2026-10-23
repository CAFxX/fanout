package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 int; F1 uint8; F2 uint32 }
type GenS6 struct { F0 int32; F1 uint; F2 int16 }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS11 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS12 struct { F0 uint8; F1 int }
type GenS13 struct { F0 uint64; F1 int16; F2 int }
type GenS14 struct { F0 bool; F1 string; F2 int16 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int16; F1 int64; F2 uint32 }
type GenS17 struct { F0 bool }
func genFn0(p0 uint32) uint8 {
    v5 := GenS1{F0: true, F1: ((1e10 + -2.25) > -2.25), F2: 0}
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int16 = 1000
        if true {
            var v8 int32 = -956
            v9 := []int{fi6, 1}
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 6; fi10++ {
                var v11 int16 = (-(v7))
                v7 -= (-(v11))
                if false {
                    v11 += v7
                }
                _ = fi10
                _ = v11
            }
            for fi12 := 0; fi12 < 5; fi12++ {
                v7 += (v7 ^ (v7 << ((18446744073709551615) & 63)))
                switch (fi12 | fi6) {
                case 128:
                    var v13 uint8 = 127
                    _ = v13
                case 0:
                    switch ("foo bar" + ("hello" + ("a" + "\n\t"))) {
                    case "":
                        v14 := []int{fi12, 0, (^(fi12))}
                        _ = v14
                    case "foo bar":
                        var v15 int32 = int32(p0)
                        _ = v15
                    default:
                        var v16 int8 = 87
                        _ = v16
                    }
                }
                v17 := []int{fi12, (fi6 / (((fi6 + fi6)) | 1)), 139, (fi6 ^ 2147483647)}
                _ = fi12
                _ = v17
            }
        }
        var v18 uint = 858
        _ = fi6
        _ = v18
        _ = v7
    }
    if true {
        fmt.Println(p0)
        fmt.Println(p0)
        p0 -= p0
    }
    v19 := GenS2{F0: 211, F1: 34}
    if false {
        p0 += 2147483648
    }
    _ = p0
    _ = v19
    _ = v5
    return 128
}
func genFn1() uint8 {
    if (!(true || false)) {
        var v20 bool = ((!(42 != 234)) && false)
        v21 := []int{}
        _ = v20
        _ = v21
    }
    var v22 uint = 18446744073709551615
    v23 := []int{-9223372036854775808, 0}
    for fi24 := 0; fi24 < 2; fi24++ {
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
        if len(v23) > 0 {
            v23[0] = (fi24 * (fi24 & int(v22)))
        }
        _ = fi24
    }
    v25 := GenS3{F0: 0}
    _ = v22
    _ = v23
    _ = v25
    return 0
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[string]uint32)
    v3 := []int{-922, 0, 9223372036854775807, 127}
    var v4 float64 = float64(v1)
    v26 := GenS2{F0: 127, F1: 128}
    v27 := GenS3{F0: -1}
    for fi28 := 0; fi28 < 11; fi28++ {
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 float32 = -2.25
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "\n\t" {
            case "\n\t":
                v4 = v4
            case "hello":
                var v31 int = fi28
                _ = v31
            }
            _ = fi29
            _ = v30
        }
        v1 = (v1 * (0.1 / (-2.25 - (float32(1e10) / v1))))
        switch "a" {
        case "\n\t":
            if false {
                var v32 int32 = 1
                _ = v32
            } else {
                v1 = 1e10
                if false {
                    var v33 uint8 = 42
                    switch "" {
                    case "a":
                        var v34 string = "foo bar"
                        _ = v34
                    }
                    var v35 int = fi28
                    _ = v33
                    _ = v35
                } else {
                    delete(v2, (("\n\t" + ("\n\t" + ((("a" + "") + ("a" + "")) + ("" + "foo bar")))) + ""))
                }
            }
        case "foo bar":
            var v36 uint32 = 1
            _ = v36
        case "a":
            for fi37 := 0; fi37 < 2; fi37++ {
                if true {
                    if true {
                        v1 -= (((3.14159 * float32(v4)) * 1.5) / ((1.5 + (0.0 * float32(1.5))) - (v1 + float32((float64((1.5 * v1)) / 1e10)))))
                        var v38 uint32 = 42
                        _ = v38
                    } else {
                        var v39 uint8 = 127
                        var v40 uint8 = 98
                        v1 = (1.5 / ((v1 / (((v1 / v1) * (v1 * v1)) / v1)) + v1))
                        _ = v39
                        _ = v40
                    }
                    if true {
                        if (2 > 1) {
                            var v41 int64 = -1
                            switch ("\n\t" + "foo bar") {
                            case "\n\t":
                                v4 = 0.0
                                if ((fi37 % (((fi28 ^ (fi37 ^ -1))) | 1)) == fi37) {
                                    var v42 uint16 = 765
                                    v3 = append(v3, -2147483648)
                                    _ = v42
                                } else {
                                    var v43 bool = (41 == -112)
                                    _ = v43
                                }
                            }
                            var v44 int8 = 0
                            var v45 uint64 = 1
                            delete(v2, "")
                            _ = v41
                            _ = v44
                            _ = v45
                        } else {
                            var v46 int8 = 42
                            _ = v46
                        }
                        var v47 float64 = (0.0 * 1.5)
                        _ = v47
                    } else {
                        fmt.Println(fi37, fi37, fi37)
                        v48 := GenS4{F0: -1, F1: 325}
                        if (!(-9223372036854775808 >= -9223372036854775808)) {
                            if false {
                                v4 -= (((0.1 + v4) * 3.14159) / 1e10)
                                var v49 bool = (!(1 > 236))
                                _ = v49
                            } else {
                                v2["\n\t"] = 1
                            }
                            fmt.Println(fi28, fi37, v1)
                            switch -479 {
                            case 657:
                                if false {
                                    var v50 uint = 4294967295
                                    v51 := GenS3{F0: int64(v50)}
                                    v52 := GenS2{F0: 135, F1: 42}
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                } else {
                                    var v53 int16 = -32768
                                    _ = v53
                                }
                            case -219:
                                v54 := GenS4{F0: 509, F1: 464}
                                _ = v54
                            }
                        } else {
                            if (false && true) {
                                v55 := GenS5{F0: (fi37 ^ 246), F1: 141, F2: 42}
                                var v56 string = ("hello" + ("a" + "hello"))
                                v3 = append(v3, 127)
                                _ = v55
                                _ = v56
                            } else {
                                if true {
                                    var v57 string = ((("\n\t" + (((("a" + "\n\t") + "hello") + "a") + "\n\t")) + "") + ("foo bar" + (("foo bar" + (("foo bar" + "hello") + "foo bar")) + "")))
                                    if true {
                                        if (!((!(42 >= 0)) || false)) {
                                            v57 = (v57 + ("" + "\n\t"))
                                        } else {
                                            if (!(!((3.14159 * v4) < ((3.14159 - (v4 / 1e10)) + (1e10 + v4))))) {
                                                var v58 int32 = 689
                                                var v59 uint32 = 4294967295
                                                v57 = ("foo bar" + "a")
                                                if (836 < uint(v59)) {
                                                    switch (fi28 + (fi28 ^ fi37)) {
                                                    case 32768:
                                                        v60 := GenS6{F0: (v58 & (-(v58))), F1: 0, F2: -32768}
                                                        _ = v60
                                                    case 9223372036854775807:
                                                        var v61 uint = 0
                                                        _ = v61
                                                    case 127:
                                                        if true {
                                                            switch (v57 + v57) {
                                                            case "a":
                                                                var v62 uint16 = 363
                                                                _ = v62
                                                            case "hello":
                                                                var v63 int8 = 0
                                                                _ = v63
                                                            case "":
                                                                switch "\n\t" {
                                                                case "a":
                                                                    _ = len(v3)
                                                                case "\n\t":
                                                                    if (v4 == (0.1 * ((v4 * v4) / ((((v4 / float64(v1)) - (1e10 / ((3.14159 - -2.25) + v4))) * v4) - (float64(v1) / v4))))) {
                                                                        switch v57 {
                                                                        case "":
                                                                            if len(v3) > 0 {
                                                                                v3[0] = -9223372036854775808
                                                                            }
                                                                        case "foo bar":
                                                                            var v64 uint32 = (v59 >> ((255) & 63))
                                                                            _ = v64
                                                                        case "\n\t":
                                                                            switch (v57 + "foo bar") {
                                                                            case "foo bar":
                                                                                fmt.Println(v58)
                                                                            }
                                                                        default:
                                                                            switch (v57 + "\n\t") {
                                                                            case "foo bar":
                                                                                v2[v57] = v59
                                                                            case "a":
                                                                                v1 += -2.25
                                                                            default:
                                                                                var v65 int16 = 0
                                                                                _ = v65
                                                                            }
                                                                            fmt.Println(v58)
                                                                        }
                                                                    } else {
                                                                        fmt.Println(v1, v57, v58)
                                                                        v57 = v57
                                                                    }
                                                                case "":
                                                                    v57 = ("foo bar" + "a")
                                                                }
                                                            default:
                                                                v66 := v2["\n\t"]
                                                                _ = v66
                                                            }
                                                            var v67 int32 = -1
                                                            fmt.Println(v1)
                                                            _ = v67
                                                        } else {
                                                            switch (v57 + "foo bar") {
                                                            case "":
                                                                var v68 uint = 4294967295
                                                                _ = v68
                                                            case "\n\t":
                                                                var v69 int8 = int8(v59)
                                                                _ = v69
                                                            case "foo bar":
                                                                if false {
                                                                    var v70 int = fi28
                                                                    _ = v70
                                                                } else {
                                                                    var v71 bool = (v58 <= v58)
                                                                    var v72 int32 = v58
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                            }
                                                            var v73 bool = (v59 == (v59 << ((1) & 63)))
                                                            var v74 float64 = ((v4 + v4) * (v4 + v4))
                                                            var v75 string = v57
                                                            v76 := GenS5{F0: fi28, F1: 82, F2: 1}
                                                            _ = v73
                                                            _ = v74
                                                            _ = v75
                                                            _ = v76
                                                        }
                                                    default:
                                                        v58 -= v58
                                                    }
                                                    v77 := GenS7{F0: 1000, F1: (v57 + v57)}
                                                    var v78 uint16 = 442
                                                    _ = v77
                                                    _ = v78
                                                } else {
                                                    switch ((v57 + (v57 + "a")) + "hello") {
                                                    case "\n\t":
                                                        if true {
                                                            if true {
                                                                v57 = v57
                                                                fmt.Println(v58)
                                                            } else {
                                                                var v79 float64 = v4
                                                                var v80 uint32 = v59
                                                                var v81 int16 = 32767
                                                                v4 -= 1e10
                                                                _ = v79
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                if true {
                                                                    if true {
                                                                        v57 = (v57 + v57)
                                                                        fmt.Println(v59, fi37, v58)
                                                                        if (3.14159 == (1.5 - (v4 + v4))) {
                                                                            fmt.Println(v58, v1)
                                                                            var v82 int16 = 1000
                                                                            _ = v82
                                                                        }
                                                                    }
                                                                    var v83 int8 = 0
                                                                    _ = v83
                                                                }
                                                            case "":
                                                                v84 := GenS8{F0: 752, F1: -1, F2: 688}
                                                                _ = v84
                                                            case "hello":
                                                                if len(v3) > 0 {
                                                                    v3[0] = (fi28 >> ((519) & 63))
                                                                }
                                                            default:
                                                                if (0.0 <= ((v1 * 0.0) + (float32(v4) * v1))) {
                                                                    if (!(435 != 1)) {
                                                                        v1 = ((0.1 + 1e10) * v1)
                                                                        v59 -= (v59 + 42)
                                                                    }
                                                                    fmt.Println(fi37)
                                                                    v85 := GenS9{F0: 9223372036854775807}
                                                                    var v86 uint32 = (^(v59))
                                                                    _ = v85
                                                                    _ = v86
                                                                }
                                                            }
                                                            if false {
                                                                var v87 uint64 = 404
                                                                var v88 int64 = 371
                                                                _ = v87
                                                                _ = v88
                                                            } else {
                                                                v58 = (v58 & (v58 + -269))
                                                            }
                                                            fmt.Println(fi37)
                                                            v89 := GenS2{F0: 255, F1: 42}
                                                            _ = v89
                                                        }
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        fmt.Println(fi37, v59, v4)
                                                    case "hello":
                                                        if true {
                                                            switch "foo bar" {
                                                            case "":
                                                                switch (fi28 - fi37) {
                                                                case 4294967295:
                                                                    v3 = append(v3, fi37)
                                                                    v90 := GenS2{F0: 127, F1: 42}
                                                                    _ = v90
                                                                case 2147483647:
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        var v91 int16 = 452
                                                                        _ = v91
                                                                    case "hello":
                                                                        switch v57 {
                                                                        case "a":
                                                                            v1 = (1.5 + float32(float64(((v1 * v1) * float32((v4 - 0.1))))))
                                                                        default:
                                                                            v59 -= 65535
                                                                        }
                                                                    case "foo bar":
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                    default:
                                                                        v58 -= (^(v58))
                                                                    }
                                                                default:
                                                                    switch 127 {
                                                                    case 11:
                                                                        v57 = v57
                                                                    default:
                                                                        v92 := GenS4{F0: (v58 - v58), F1: 42}
                                                                        if (-2147483648 != v58) {
                                                                            var v93 uint64 = 42
                                                                            if true {
                                                                                _ = len(v3)
                                                                            } else {
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = (fi37 | (fi28 * int(v93)))
                                                                                }
                                                                                if (2 > 274) {
                                                                                    v59 = (v59 >> ((1) & 63))
                                                                                    var v94 uint8 = 128
                                                                                    if false {
                                                                                        v58 = (v58 - (-(v58)))
                                                                                        if true {
                                                                                            v1 = 3.14159
                                                                                            var v95 uint16 = 0
                                                                                            _ = v95
                                                                                        }
                                                                                    } else {
                                                                                        if false {
                                                                                            fmt.Println(v94, v94, v57)
                                                                                            var v96 string = (v57 + v57)
                                                                                            _ = v96
                                                                                        } else {
                                                                                            fmt.Println(v57, v93)
                                                                                        }
                                                                                        var v97 bool = true
                                                                                        switch (fi37 % ((4294967295) | 1)) {
                                                                                        case -1:
                                                                                            v98 := GenS2{F0: genFn0(v59), F1: 103}
                                                                                            _ = v98
                                                                                        }
                                                                                        if ((488 <= 946) && (!true)) {
                                                                                            if false {
                                                                                                v97 = (((v1 + v1) / 3.14159) < v1)
                                                                                            }
                                                                                        }
                                                                                        if (fi28 < (fi37 - (-(fi37)))) {
                                                                                            var v99 int = fi37
                                                                                            if true {
                                                                                                v100 := GenS10{F0: 543, F1: (^(v59)), F2: 415}
                                                                                                var v101 uint = 0
                                                                                                _ = v100
                                                                                                _ = v101
                                                                                            } else {
                                                                                                v1 += float32((v4 / ((float64(1.5) + (float64(0.0) + v4)) / (1e10 - (v4 + (1.5 * 0.1))))))
                                                                                                if v97 {
                                                                                                    if (!false) {
                                                                                                        if (float64(v1) <= (v4 + 3.14159)) {
                                                                                                            v58 = (v58 * v58)
                                                                                                            var v102 uint16 = 42
                                                                                                            if (-128 > 26) {
                                                                                                                var v103 int16 = -1
                                                                                                                var v104 uint8 = uint8(v59)
                                                                                                                _ = v103
                                                                                                                _ = v104
                                                                                                            } else {
                                                                                                                v1 = -2.25
                                                                                                                var v105 int32 = v58
                                                                                                                _ = v105
                                                                                                            }
                                                                                                            _ = v102
                                                                                                        }
                                                                                                    }
                                                                                                    var v106 bool = ((v97 || (("a" + v57) == v57)) || true)
                                                                                                    var v107 bool = v106
                                                                                                    v108 := GenS8{F0: v93, F1: 1, F2: v93}
                                                                                                    var v109 int16 = -339
                                                                                                    _ = v106
                                                                                                    _ = v107
                                                                                                    _ = v108
                                                                                                    _ = v109
                                                                                                }
                                                                                            }
                                                                                            _ = v99
                                                                                        }
                                                                                        v110 := GenS4{F0: v58, F1: 42}
                                                                                        v58 = v58
                                                                                        if ((v94 & (v94 << ((42) & 63))) > v94) {
                                                                                            v97 = v97
                                                                                            if (!true) {
                                                                                                var v111 int64 = 446
                                                                                                _ = v111
                                                                                            }
                                                                                            var v112 float64 = (0.1 - 0.0)
                                                                                            _ = v112
                                                                                        } else {
                                                                                            v1 += 1.5
                                                                                            v113 := GenS5{F0: (^(fi37)), F1: v94, F2: 2147483648}
                                                                                            _ = v113
                                                                                        }
                                                                                        _ = v110
                                                                                        _ = v97
                                                                                    }
                                                                                    v114 := GenS11{F0: 80, F1: v94, F2: 874}
                                                                                    _ = v114
                                                                                    _ = v94
                                                                                }
                                                                                v115 := GenS12{F0: 42, F1: fi37}
                                                                                _ = v115
                                                                            }
                                                                            switch fi37 {
                                                                            case 32767:
                                                                                if true {
                                                                                    fmt.Println(v4, v1, v93)
                                                                                    switch (("hello" + v57) + "hello") {
                                                                                    case "foo bar":
                                                                                        var v116 int64 = -1
                                                                                        _ = v116
                                                                                    case "":
                                                                                        var v117 uint16 = 280
                                                                                        _ = v117
                                                                                    default:
                                                                                        v4 = v4
                                                                                    }
                                                                                }
                                                                                v57 = v57
                                                                            default:
                                                                                var v118 int8 = 109
                                                                                _ = v118
                                                                            }
                                                                            _ = v93
                                                                        } else {
                                                                            fmt.Println(v58)
                                                                            v119 := GenS13{F0: 415, F1: 1000, F2: fi37}
                                                                            v4 = (v4 + v4)
                                                                            v120 := GenS9{F0: 9223372036854775807}
                                                                            _ = v119
                                                                            _ = v120
                                                                        }
                                                                        _ = v92
                                                                    }
                                                                }
                                                            case "\n\t":
                                                                switch "\n\t" {
                                                                case "a":
                                                                    v121 := GenS3{F0: 1}
                                                                    _ = v121
                                                                }
                                                            default:
                                                                v1 += 1.5
                                                            }
                                                            fmt.Println(v57)
                                                        }
                                                    case "":
                                                        var v122 uint = 65535
                                                        _ = v122
                                                    default:
                                                        v123 := v2[(("hello" + (v57 + "hello")) + "hello")]
                                                        _ = v123
                                                    }
                                                    fmt.Println(v1, v57)
                                                }
                                                _ = v58
                                                _ = v59
                                            }
                                            v124 := GenS3{F0: 0}
                                            _ = v124
                                        }
                                    } else {
                                        if false {
                                            v125 := GenS13{F0: 127, F1: 1000, F2: fi28}
                                            var v126 uint32 = 449
                                            _ = v125
                                            _ = v126
                                        }
                                        _ = len(v2)
                                        v127 := GenS14{F0: (fi28 == fi28), F1: ("\n\t" + "\n\t"), F2: 32767}
                                        _ = v127
                                    }
                                    var v128 int64 = -9223372036854775808
                                    v4 = (v4 + 0.0)
                                    var v129 int8 = -93
                                    if true {
                                        if false {
                                            var v130 bool = false
                                            v128 += (v128 ^ (v128 & (v128 >> ((4294967295) & 63))))
                                            _ = v130
                                        }
                                        v128 += v128
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        var v131 uint32 = 2147483647
                                        v128 = v128
                                        switch (v57 + (("a" + v57) + "hello")) {
                                        case "":
                                            var v132 int16 = 1000
                                            _ = v132
                                        case "hello":
                                            delete(v2, "foo bar")
                                        default:
                                            var v133 int8 = (v129 / ((v129) | 1))
                                            _ = v133
                                        }
                                        _ = v131
                                    }
                                    var v134 int64 = v128
                                    v134 = (v128 >> ((693) & 63))
                                    _ = v128
                                    _ = v129
                                    _ = v134
                                    _ = v57
                                }
                                v2[(((("" + (("" + "") + "foo bar")) + "foo bar") + "\n\t") + "\n\t")] = 4294967295
                                _ = len(v3)
                            }
                            switch 32768 {
                            case -505:
                                v135 := GenS6{F0: -1, F1: 234, F2: -279}
                                _ = v135
                            default:
                                if (-2.25 != -2.25) {
                                    _ = len(v2)
                                }
                            }
                        }
                        var v136 float32 = (v1 * -2.25)
                        var v137 float32 = (v1 - float32(v4))
                        _ = v136
                        _ = v137
                        _ = v48
                    }
                } else {
                    var v138 uint64 = 1
                    _ = v138
                }
                v139 := v2["foo bar"]
                if (-750 >= -2147483648) {
                    v1 = float32(v4)
                    v4 = (0.0 / (v4 * v4))
                    v140 := GenS15{F0: -128}
                    _ = v140
                }
                _ = fi37
                _ = v139
            }
        default:
            v4 = v4
        }
        _ = fi28
    }
    if true {
        v141 := GenS7{F0: -1, F1: "\n\t"}
        if len(v3) > 0 {
            v3[0] = 32767
        }
        v3 = append(v3, -927)
        _ = v141
    }
    if false {
        if (true || (128 < -9223372036854775808)) {
            if (float32((v4 + v4)) == (3.14159 - float32(float64(3.14159)))) {
                if (65535 != 420) {
                    _ = len(v2)
                    v1 = (((float32(v4) + v1) / 1.5) / 1e10)
                }
                if (false || true) {
                    v1 = v1
                    v142 := GenS6{F0: 0, F1: 2, F2: 32767}
                    _ = v142
                }
                var v143 int16 = 1000
                _ = v143
            }
            if len(v3) > 0 {
                v3[0] = -2147483648
            }
            v4 += float64(1e10)
        }
        switch -2147483648 {
        case 4294967295:
            v1 = float32(float64(float32((v4 * 1e10))))
        case 185:
            var v144 int16 = 42
            _ = v144
        default:
            v4 = float64(v1)
        }
        for fi145 := 0; fi145 < 10; fi145++ {
            fmt.Println(v4, v4, v1)
            if len(v3) > 0 {
                v3[0] = (-(fi145))
            }
            v4 = -2.25
            _ = fi145
        }
    } else {
        v1 = (v1 * v1)
        for fi146 := 0; fi146 < 19; fi146++ {
            if len(v3) > 0 {
                v3[0] = -859
            }
            v1 += float32((v4 + 0.0))
            _ = len(v3)
            _ = fi146
        }
        for fi147 := 0; fi147 < 12; fi147++ {
            for fi148 := 0; fi148 < 10; fi148++ {
                switch 32768 {
                case 0:
                    var v149 float64 = v4
                    _ = v149
                case -1:
                    if true {
                        v150 := GenS8{F0: 4294967295, F1: 1, F2: 0}
                        _ = v150
                    } else {
                        if ((!(!(((!(((1 == 1000) || (((!false) && false) || (-83 >= 42))) && ((((v4 * float64(v1)) * float64(((3.14159 / 1.0) - (0.0 - ((1e10 - v1) - (0.1 / 1e10)))))) / float64(float32(v4))) <= 3.14159))) || false) && (1 < 1)))) && false) {
                            v4 = 1e10
                        } else {
                            v3 = append(v3, fi148)
                            v151 := GenS15{F0: 34}
                            v1 += v1
                            _ = v151
                        }
                    }
                default:
                    switch fi148 {
                    case 1:
                        switch "a" {
                        case "":
                            var v152 float32 = 1e10
                            _ = v152
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    case 0:
                        var v153 bool = true
                        _ = v153
                    case -613:
                        v154 := GenS16{F0: -32768, F1: -827, F2: 2147483647}
                        _ = v154
                    default:
                        v4 = float64(1e10)
                    }
                }
                if ((fi148 >> ((255) & 63)) >= -1) {
                    fmt.Println(fi148, fi147)
                } else {
                    var v155 uint32 = 42
                    _ = v155
                }
                _ = fi148
            }
            _ = fi147
        }
    }
    v4 -= 0.0
    fmt.Println(v1)
    v1 = (1.5 / v1)
    v1 -= v1
    v156 := GenS1{F0: false, F1: (32767 != 544), F2: 128}
    v157 := GenS17{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v156
    _ = v157
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
