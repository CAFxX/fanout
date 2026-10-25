package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 int }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 int16 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int16; F1 int }
type GenS12 struct { F0 bool; F1 uint16; F2 int }
type GenS13 struct { F0 uint; F1 uint64 }
type GenS14 struct { F0 int; F1 int8 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint32 }

func main() {
    var v1 uint64 = 0
    var v2 float32 = float32(float64(3.14159))
    var v3 float64 = 1.5
    fmt.Println(v1, v1)
    fmt.Println(v3, v2)
    fmt.Println(v1, v1)
    if false {
        v4 := GenS1{F0: 0, F1: 13, F2: 127}
        _ = v4
    } else {
        fmt.Println(v3, v1)
        if (-1 != -1) {
            if true {
                if (638 <= -1) {
                    v5 := []int{}
                    v6 := GenS1{F0: 2, F1: 1, F2: 1}
                    _ = v5
                    _ = v6
                }
                if (v1 != v1) {
                    fmt.Println(v1, v3, v1)
                    v1 = 0
                }
                v7 := GenS1{F0: 4294967295, F1: 381, F2: -2147483648}
                _ = v7
            } else {
                var v8 string = ((("" + (("" + "\n\t") + ("hello" + ("a" + "\n\t")))) + ("foo bar" + ("a" + (("\n\t" + ("foo bar" + "hello")) + "")))) + "\n\t")
                v9 := GenS2{F0: 593, F1: 42}
                v10 := GenS3{F0: 65535}
                _ = v10
                _ = v8
                _ = v9
            }
        }
        var v11 int16 = -346
        _ = v11
    }
    var v12 bool = true
    if v12 {
        switch "" {
        case "foo bar":
            var v13 uint64 = v1
            _ = v13
        case "hello":
            var v14 uint16 = 42
            _ = v14
        default:
            v1 -= 18446744073709551615
        }
        v1 -= (v1 << ((834) & 63))
        var v15 float64 = 3.14159
        _ = v15
    } else {
        var v16 int32 = -1
        v2 = 1e10
        _ = v16
    }
    switch "\n\t" {
    case "hello":
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 int = (fi17 % ((-2147483648) | 1))
            for fi19 := 0; fi19 < 12; fi19++ {
                if v12 {
                    v20 := []int{fi19, v18, (fi17 | (^(v18)))}
                    _ = v20
                } else {
                    if (0 > 128) {
                        if v12 {
                            var v21 int16 = -32768
                            _ = v21
                        }
                    }
                    switch fi17 {
                    case -726:
                        v3 = ((v3 * (0.1 * 1.5)) - -2.25)
                    case -609:
                        v3 -= v3
                    case -2147483648:
                        var v22 float64 = 1.5
                        v23 := GenS3{F0: 2}
                        _ = v22
                        _ = v23
                    default:
                        v24 := GenS3{F0: 2}
                        _ = v24
                    }
                }
                v25 := GenS2{F0: 0, F1: 0}
                _ = fi19
                _ = v25
            }
            _ = fi17
            _ = v18
        }
    case "foo bar":
        var v26 int8 = 127
        _ = v26
    default:
        var v27 float64 = v3
        _ = v27
    }
    switch 32767 {
    case 9223372036854775807:
        v28 := []int{1}
        _ = v28
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := []int{}
        _ = fi29
        _ = v30
    }
    var v31 uint64 = (v1 % ((v1) | 1))
    v12 = false
    v32 := GenS4{F0: v1, F1: 438, F2: 423}
    v31 += 1
    var v33 uint16 = 814
    if false {
        var v34 int32 = 1
        if v12 {
            if (127 > 100) {
                var v35 uint8 = 2
                _ = v35
            } else {
                var v36 float32 = v2
                if (v34 < (v34 + (^(v34)))) {
                    v37 := GenS5{F0: ("\n\t" + "hello")}
                    if (v12 && false) {
                        v38 := GenS5{F0: "hello"}
                        switch -844 {
                        case -673:
                            v39 := GenS2{F0: v33, F1: uint(v1)}
                            _ = v39
                        }
                        fmt.Println(v2, v31, v1)
                        _ = v38
                    } else {
                        switch "\n\t" {
                        case "":
                            if false {
                                for fi40 := 0; fi40 < 9; fi40++ {
                                    var v41 int64 = -9223372036854775808
                                    fmt.Println(v41, v34)
                                    _ = fi40
                                    _ = v41
                                }
                                if (false && (true && (42 > 949))) {
                                    fmt.Println(v1, v31)
                                    var v42 int32 = v34
                                    for fi43 := 0; fi43 < 5; fi43++ {
                                        v36 = v36
                                        _ = fi43
                                    }
                                    _ = v42
                                }
                                var v44 uint64 = v1
                                _ = v44
                            } else {
                                v12 = (((("hello" + ("a" + ("" + "a"))) + (("\n\t" + ("foo bar" + "hello")) + ("foo bar" + "a"))) + ("a" + "a")) <= "a")
                            }
                        case "\n\t":
                            v45 := []int{-9223372036854775808, 4294967295}
                            _ = v45
                        case "foo bar":
                            if ((100 != int8(v34)) || false) {
                                var v46 int = 560
                                v47 := GenS5{F0: "a"}
                                switch (v46 | 448) {
                                case 582:
                                    var v48 uint16 = (v33 % ((v33) | 1))
                                    _ = v48
                                case 1:
                                    var v49 uint16 = (v33 >> ((546) & 63))
                                    _ = v49
                                case -590:
                                    v46 = (v46 * (v46 * (v46 << ((350) & 63))))
                                }
                                _ = v46
                                _ = v47
                            }
                        default:
                            for fi50 := 0; fi50 < 3; fi50++ {
                                if (!(true || (int(v34) != -310))) {
                                    var v51 int16 = 32767
                                    _ = v51
                                }
                                for fi52 := 0; fi52 < 14; fi52++ {
                                    var v53 string = ""
                                    _ = fi52
                                    _ = v53
                                }
                                _ = fi50
                            }
                        }
                        v54 := GenS1{F0: 1, F1: 0, F2: 127}
                        _ = v54
                    }
                    v55 := []int{1, 128}
                    _ = v37
                    _ = v55
                }
                fmt.Println(v33)
                _ = v36
            }
            var v56 uint8 = 1
            _ = v56
        }
        _ = v34
    }
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "\n\t":
        if ((!(4294967295 != (v31 + (v1 ^ v31)))) && (v12 && (v12 && false))) {
            fmt.Println(v31, v3, v1)
            if (((2 != 0) && v12) || (-66 != 0)) {
                v57 := GenS6{F0: 318, F1: 127}
                for fi58 := 0; fi58 < 18; fi58++ {
                    if v12 {
                        v59 := GenS5{F0: ("foo bar" + "")}
                        var v60 uint8 = uint8(v31)
                        v61 := GenS7{F0: -32768}
                        _ = v59
                        _ = v60
                        _ = v61
                    } else {
                        switch "\n\t" {
                        case "hello":
                            v31 = (^(v1))
                        case "\n\t":
                            v62 := GenS2{F0: v33, F1: 860}
                            _ = v62
                        }
                        var v63 int32 = -963
                        var v64 float32 = v2
                        _ = v63
                        _ = v64
                    }
                    for fi65 := 0; fi65 < 3; fi65++ {
                        switch "hello" {
                        case "":
                            switch ("a" + "\n\t") {
                            case "a":
                                v3 = 1.5
                                fmt.Println(fi65, v31, v3)
                                v12 = false
                                v66 := GenS3{F0: (v33 - (v33 % ((v33) | 1)))}
                                _ = v66
                            case "\n\t":
                                var v67 uint = 65535
                                _ = v67
                            default:
                                if true {
                                    v68 := []int{fi58, fi65, 718, fi58}
                                    v1 += (v31 & (^(v31)))
                                    _ = v68
                                }
                            }
                        default:
                            if ((!(!("a" != ("" + "")))) || true) {
                                fmt.Println(v1)
                                v33 += 255
                                if true {
                                    switch fi65 {
                                    case 4294967295:
                                        if ((1.5 < (0.0 / 3.14159)) && (203 < 171)) {
                                            if v12 {
                                                v31 += (v1 << ((1) & 63))
                                                var v69 float32 = 0.1
                                                _ = v69
                                            }
                                        }
                                    default:
                                        switch fi65 {
                                        case 2147483647:
                                            switch "foo bar" {
                                            case "a":
                                                v70 := GenS3{F0: (v33 | v33)}
                                                _ = v70
                                            case "\n\t":
                                                v71 := []int{fi65, fi58}
                                                _ = v71
                                            case "foo bar":
                                                v2 = 0.0
                                            default:
                                                if v12 {
                                                    if v12 {
                                                        v1 = (v31 / (((v31 << ((2) & 63))) | 1))
                                                        v72 := []int{(fi65 % ((fi58) | 1)), fi58, fi58, fi65}
                                                        v31 += (v31 | v31)
                                                        _ = v72
                                                    }
                                                    switch ("foo bar" + "foo bar") {
                                                    case "foo bar":
                                                        if (180 != v33) {
                                                            if true {
                                                                switch (fi58 % ((int(v31)) | 1)) {
                                                                case -1:
                                                                    v1 = (^(v1))
                                                                case -9223372036854775808:
                                                                    fmt.Println(v31)
                                                                }
                                                                v2 = v2
                                                                var v73 int64 = 34
                                                                if (v33 < (v33 ^ (v33 & (v33 << ((2) & 63))))) {
                                                                    if v12 {
                                                                        var v74 float32 = (float32(1e10) + 3.14159)
                                                                        fmt.Println(v12)
                                                                        v75 := []int{int(v73), (fi58 % ((fi65) | 1)), 127}
                                                                        _ = v74
                                                                        _ = v75
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        var v76 int32 = 2147483647
                                                                        v33 += (v33 * uint16(v1))
                                                                        var v77 uint32 = 2147483648
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                    var v78 uint16 = (^(v33))
                                                                    var v79 float32 = v2
                                                                    _ = v78
                                                                    _ = v79
                                                                }
                                                                _ = v73
                                                            }
                                                            var v80 string = "foo bar"
                                                            if v12 {
                                                                fmt.Println(v2, fi65)
                                                                v81 := GenS7{F0: -666}
                                                                _ = v81
                                                            } else {
                                                                var v82 int64 = -9223372036854775808
                                                                var v83 float32 = v2
                                                                var v84 uint32 = 65535
                                                                _ = v82
                                                                _ = v83
                                                                _ = v84
                                                            }
                                                            var v85 int64 = 1
                                                            switch fi58 {
                                                            case 390:
                                                                v86 := GenS6{F0: 42, F1: 1}
                                                                _ = v86
                                                            default:
                                                                var v87 uint64 = (v31 & (v31 & (v31 & v1)))
                                                                _ = v87
                                                            }
                                                            _ = v80
                                                            _ = v85
                                                        }
                                                    case "\n\t":
                                                        v3 = (1.5 + ((v3 / v3) / v3))
                                                        v88 := []int{-9223372036854775808}
                                                        _ = v88
                                                    case "":
                                                        v1 -= (^(v1))
                                                    default:
                                                        var v89 uint32 = 2
                                                        _ = v89
                                                    }
                                                    var v90 uint16 = v33
                                                    _ = v90
                                                }
                                            }
                                        case -1:
                                            v12 = (v12 && v12)
                                            switch ((("foo bar" + "") + "foo bar") + ("" + "\n\t")) {
                                            case "\n\t":
                                                if ((^(v1)) == 896) {
                                                    var v91 int8 = 100
                                                    switch fi58 {
                                                    case -1:
                                                        v92 := []int{(fi65 / ((32767) | 1)), int(v91), (^(fi58)), (-(fi65))}
                                                        _ = v92
                                                    }
                                                    if v12 {
                                                        v12 = v12
                                                        v93 := GenS6{F0: uint(v33), F1: (v91 / ((int8(v31)) | 1))}
                                                        _ = v93
                                                    } else {
                                                        fmt.Println(v91, v33)
                                                    }
                                                    if ((false && v12) && (847 == 12)) {
                                                        var v94 uint64 = (^(v1))
                                                        v95 := []int{32768, fi58}
                                                        fmt.Println(v12)
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    var v96 float64 = ((float64((float32((float64(v2) + float64(float32((v3 * 0.0))))) + 0.1)) + float64(v2)) + -2.25)
                                                    var v97 uint32 = 2147483647
                                                    _ = v91
                                                    _ = v96
                                                    _ = v97
                                                } else {
                                                    v98 := GenS8{F0: (fi65 % ((fi65) | 1)), F1: -32768}
                                                    _ = v98
                                                }
                                            case "hello":
                                                switch (fi58 / (((-(fi65))) | 1)) {
                                                case 0:
                                                    if v12 {
                                                        var v99 int64 = -9223372036854775808
                                                        v100 := GenS9{F0: (fi65 << ((752) & 63))}
                                                        if v12 {
                                                            switch ("foo bar" + (("a" + ("a" + "\n\t")) + "a")) {
                                                            case "foo bar":
                                                                v33 += v33
                                                            default:
                                                                if v12 {
                                                                    var v101 int16 = 32767
                                                                    switch ("foo bar" + (("" + ("\n\t" + "foo bar")) + ("hello" + "hello"))) {
                                                                    case "\n\t":
                                                                        var v102 uint8 = 2
                                                                        _ = v102
                                                                    case "a":
                                                                        var v103 int16 = (-(v101))
                                                                        _ = v103
                                                                    case "foo bar":
                                                                        fmt.Println(v33)
                                                                    }
                                                                    _ = v101
                                                                }
                                                            }
                                                            var v104 int16 = 32767
                                                            _ = v104
                                                        } else {
                                                            switch ("foo bar" + ("foo bar" + ("" + "a"))) {
                                                            case "a":
                                                                fmt.Println(fi65, v99)
                                                            case "foo bar":
                                                                v105 := []int{(fi58 - -675)}
                                                                _ = v105
                                                            case "hello":
                                                                v106 := []int{-369, (fi58 << ((1) & 63))}
                                                                _ = v106
                                                            default:
                                                                v107 := GenS5{F0: "hello"}
                                                                _ = v107
                                                            }
                                                            var v108 bool = true
                                                            _ = v108
                                                        }
                                                        _ = v100
                                                        _ = v99
                                                    }
                                                case 627:
                                                    v109 := []int{}
                                                    _ = v109
                                                case 1:
                                                    v110 := []int{(fi58 << ((0) & 63)), 60, -2147483648, fi65}
                                                    _ = v110
                                                }
                                            default:
                                                var v111 uint32 = 4294967295
                                                _ = v111
                                            }
                                            v31 -= v31
                                        case 42:
                                            v112 := []int{(fi58 >> ((uint(v33)) & 63))}
                                            fmt.Println(v31)
                                            var v113 int8 = 42
                                            _ = v112
                                            _ = v113
                                        }
                                    }
                                    if true {
                                        v1 -= 4294967295
                                        v114 := []int{}
                                        var v115 int32 = 42
                                        _ = v114
                                        _ = v115
                                    } else {
                                        v116 := GenS8{F0: 2147483647, F1: 543}
                                        if (v12 && ((fi65 ^ (fi65 % (((fi58 - 128)) | 1))) <= fi58)) {
                                            v31 = (v1 / ((uint64(v33)) | 1))
                                            v33 = v33
                                            v117 := GenS10{F0: 64}
                                            _ = v117
                                        } else {
                                            if false {
                                                switch fi65 {
                                                case 9223372036854775807:
                                                    var v118 bool = v12
                                                    _ = v118
                                                }
                                                v119 := []int{}
                                                _ = v119
                                            }
                                            v120 := GenS8{F0: 128, F1: 1000}
                                            v121 := GenS11{F0: int16(v33), F1: -9223372036854775808}
                                            _ = v120
                                            _ = v121
                                        }
                                        _ = v116
                                    }
                                    v122 := []int{(fi58 & (fi58 >> ((2) & 63)))}
                                    if len(v122) >= 2 {
                                        v122 = v122[1:len(v122)]
                                    }
                                    if v12 {
                                        switch (("a" + ("\n\t" + ("foo bar" + "foo bar"))) + "") {
                                        case "":
                                            v123 := GenS1{F0: 131, F1: 2147483647, F2: (^(fi58))}
                                            _ = v123
                                        case "a":
                                            v12 = v12
                                        case "hello":
                                            _ = len(v122)
                                        default:
                                            v1 -= (^(v31))
                                        }
                                        v124 := GenS12{F0: (2147483648 > 4294967295), F1: 855, F2: (fi58 >> ((18446744073709551615) & 63))}
                                        switch fi65 {
                                        case 32768:
                                            var v125 int32 = 2147483647
                                            _ = v125
                                        }
                                        _ = v124
                                    } else {
                                        if len(v122) >= 2 {
                                            v122 = v122[1:len(v122)]
                                        }
                                    }
                                    _ = v122
                                }
                                var v126 uint16 = v33
                                _ = v126
                            } else {
                                if v12 {
                                    if v12 {
                                        fmt.Println(v1)
                                        if false {
                                            fmt.Println(v3, v1, v1)
                                            fmt.Println(v1)
                                        } else {
                                            v12 = v12
                                            v1 = v1
                                        }
                                        v2 += (v2 + ((float32(3.14159) - (v2 - (v2 - float32(float64(float32(1e10)))))) - ((((v2 * -2.25) + (v2 + float32((float64((v2 * float32(((0.0 * v3) / (float64(v2) / v3))))) * ((v3 * v3) + v3))))) - 0.1) + (3.14159 + 1e10))))
                                    } else {
                                        v127 := GenS13{F0: 2, F1: 18446744073709551615}
                                        fmt.Println(fi65, v1, fi65)
                                        _ = v127
                                    }
                                    switch (fi58 * (fi65 * -949)) {
                                    case 32767:
                                        v128 := GenS14{F0: fi58, F1: 48}
                                        _ = v128
                                    case 127:
                                        fmt.Println(v1, v33, v31)
                                    case -2147483648:
                                        if (65535 <= 2) {
                                            var v129 uint8 = 42
                                            v2 = (v2 + float32(((float64(v2) + (v3 - v3)) * v3)))
                                            v130 := []int{fi58, fi65, fi65, (^(fi65))}
                                            _ = v129
                                            _ = v130
                                        }
                                    default:
                                        if (!(!v12)) {
                                            switch fi58 {
                                            case 2147483647:
                                                var v131 int32 = 1
                                                _ = v131
                                            }
                                            v2 += v2
                                        } else {
                                            var v132 int8 = 127
                                            fmt.Println(v3)
                                            _ = v132
                                        }
                                    }
                                }
                                var v133 uint64 = (v1 / ((42) | 1))
                                _ = v133
                            }
                        }
                        _ = fi65
                    }
                    v134 := GenS7{F0: int16(v31)}
                    _ = fi58
                    _ = v134
                }
                for fi135 := 0; fi135 < 15; fi135++ {
                    fmt.Println(v2, v33)
                    _ = fi135
                }
                _ = v57
            }
        } else {
            switch "a" {
            case "\n\t":
                v136 := GenS15{F0: 42}
                _ = v136
            case "foo bar":
                v3 = ((1e10 - ((float64(1.5) * v3) + v3)) - 0.1)
            case "a":
                v137 := GenS16{F0: uint32(v31)}
                _ = v137
            }
            v12 = v12
        }
    case "":
        var v138 int32 = -316
        _ = v138
    case "foo bar":
        v1 += v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
