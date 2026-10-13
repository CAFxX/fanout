package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int64; F1 int16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int8; F1 int64; F2 bool }
type GenS12 struct { F0 int32; F1 int32 }
type GenS13 struct { F0 bool; F1 int32 }
type GenS14 struct { F0 uint8; F1 string; F2 uint64 }
type GenS15 struct { F0 string; F1 int32 }
type GenS16 struct { F0 uint32 }
func genFn0() uint64 {
    var v7 float32 = 0.0
    switch ("a" + "foo bar") {
    case "a":
        var v8 uint8 = 252
        _ = v8
    case "\n\t":
        for fi9 := 0; fi9 < 20; fi9++ {
            if true {
                if false {
                    var v10 int64 = -1
                    for fi11 := 0; fi11 < 16; fi11++ {
                        var v12 uint = 2
                        _ = fi11
                        _ = v12
                    }
                    _ = v10
                }
                if (812 <= fi9) {
                    var v13 uint8 = 42
                    v14 := []int{562}
                    fmt.Println(fi9, v7)
                    _ = v13
                    _ = v14
                }
            } else {
                for fi15 := 0; fi15 < 9; fi15++ {
                    switch ("foo bar" + "a") {
                    case "":
                        var v16 int32 = 0
                        _ = v16
                    case "a":
                        if (!true) {
                            v17 := GenS1{F0: 1}
                            _ = v17
                        }
                    default:
                        v7 = (v7 * (0.0 - v7))
                    }
                    v7 -= v7
                    _ = fi15
                }
                switch ("foo bar" + "foo bar") {
                case "\n\t":
                    v18 := GenS2{F0: 4294967295, F1: 1000}
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if false {
                            var v20 float32 = (0.1 - 0.0)
                            v20 = (v20 + 0.0)
                            var v21 int8 = -1
                            var v22 int64 = 1
                            v23 := []int{(fi9 - fi9)}
                            _ = v20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        v24 := []int{9223372036854775807}
                        _ = fi19
                        _ = v24
                    }
                }
                v25 := []int{fi9, fi9, 9223372036854775807}
                _ = v25
            }
            switch "foo bar" {
            case "a":
                v26 := GenS3{F0: "foo bar", F1: "a"}
                _ = v26
            case "hello":
                v7 -= v7
            case "foo bar":
                var v27 uint8 = 138
                _ = v27
            }
            _ = fi9
        }
    }
    v7 = v7
    v7 += v7
    _ = v7
    return 777
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]int)
    var v3 int64 = -641
    var v4 int16 = 521
    var v5 string = ("a" + "foo bar")
    v6 := []uint32{}
    fmt.Println(v1, v3)
    switch v5 {
    case "\n\t":
        v5 = v5
    }
    var v28 uint = 4294967295
    if (uint(v1) != 628) {
        for fi29 := 0; fi29 < 7; fi29++ {
            if true {
                v3 -= -259
                v3 += v3
                var v30 int32 = 42
                _ = v30
            } else {
                for fi31 := 0; fi31 < 5; fi31++ {
                    if ((!false) && (-1 >= (v3 + v3))) {
                        v32 := GenS3{F0: ("\n\t" + "hello"), F1: v5}
                        var v33 int32 = 1
                        v28 = v28
                        _ = v32
                        _ = v33
                    } else {
                        if false {
                            v28 = 590
                        }
                        v34 := GenS1{F0: v28}
                        _ = v34
                    }
                    var v35 string = "\n\t"
                    v36 := v2[-625]
                    _ = fi31
                    _ = v35
                    _ = v36
                }
                if true {
                    if (1 > (v4 ^ (^(v4)))) {
                        for fi37 := 0; fi37 < 12; fi37++ {
                            if true {
                                v5 = v5
                                fmt.Println(fi37, v4)
                                var v38 uint8 = 128
                                _ = v38
                            }
                            _ = fi37
                        }
                    } else {
                        for fi39 := 0; fi39 < 20; fi39++ {
                            if (true && false) {
                                v6 = append(v6, uint32(v1))
                            }
                            switch v5 {
                            case "\n\t":
                                var v40 uint = (^(v28))
                                _ = v40
                            default:
                                var v41 bool = true
                                _ = v41
                            }
                            _ = fi39
                        }
                        v42 := GenS4{F0: 100}
                        var v43 int16 = (v4 + (v4 / (((v4 & v4)) | 1)))
                        _ = v42
                        _ = v43
                    }
                    if false {
                        switch v5 {
                        case "foo bar":
                            v44 := GenS5{F0: 2, F1: 127}
                            _ = v44
                        default:
                            _ = len(v2)
                        }
                    }
                } else {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if false {
                        var v45 float32 = (1.5 * 0.0)
                        if true {
                            var v46 float64 = 3.14159
                            v46 = 3.14159
                            if ((!false) && ((v3 / ((v3) | 1)) > (v3 % ((v3) | 1)))) {
                                v1 = v1
                            }
                            _ = v46
                        } else {
                            if false {
                                v47 := GenS6{F0: v4}
                                v28 -= v28
                                _ = v47
                            } else {
                                _ = len(v2)
                                if true {
                                    var v48 float32 = float32(1e10)
                                    if true {
                                        for fi49 := 0; fi49 < 11; fi49++ {
                                            if (!false) {
                                                v48 += 0.1
                                                v50 := v2[fi29]
                                                var v51 float64 = 1.5
                                                if false {
                                                    switch (-(fi29)) {
                                                    case 280:
                                                        var v52 uint8 = 0
                                                        _ = v52
                                                    case 380:
                                                        fmt.Println(v5, v1)
                                                    case 0:
                                                        fmt.Println(v4, v1)
                                                    }
                                                    v53 := GenS3{F0: v5, F1: v5}
                                                    var v54 int32 = 0
                                                    _ = v53
                                                    _ = v54
                                                } else {
                                                    if ((v45 / ((v45 - v48) / v45)) < (v45 / v45)) {
                                                        var v55 float64 = v51
                                                        v56 := GenS2{F0: 648, F1: (v4 - 32767)}
                                                        switch (fi49 << ((v28) & 63)) {
                                                        case -2147483648:
                                                            var v57 uint32 = 2
                                                            _ = v57
                                                        case -843:
                                                            var v58 uint8 = 126
                                                            _ = v58
                                                        }
                                                        _ = v55
                                                        _ = v56
                                                    } else {
                                                        fmt.Println(fi29, fi49)
                                                        delete(v2, fi29)
                                                    }
                                                    v4 -= -34
                                                }
                                                if (127 != 255) {
                                                    v48 = v45
                                                    v51 -= v51
                                                    var v59 uint16 = 65535
                                                    _ = v59
                                                } else {
                                                    var v60 string = v5
                                                    v61 := GenS7{F0: ((v60 + "hello") + ("a" + v5))}
                                                    v3 += v3
                                                    _ = v60
                                                    _ = v61
                                                }
                                                _ = v50
                                                _ = v51
                                            } else {
                                                v1 -= v1
                                            }
                                            if (!false) {
                                                if false {
                                                    v62 := GenS5{F0: 233, F1: 255}
                                                    _ = v62
                                                }
                                                v45 -= v48
                                                var v63 int32 = 1
                                                _ = v63
                                            } else {
                                                fmt.Println(fi49)
                                                v5 = v5
                                            }
                                            switch v5 {
                                            case "a":
                                                var v64 uint32 = 607
                                                _ = v64
                                            case "hello":
                                                if ((true || (1 == 127)) && (true || (0 == 34))) {
                                                    var v65 uint16 = 808
                                                    v66 := GenS1{F0: 0}
                                                    v67 := GenS5{F0: 1, F1: 42}
                                                    _ = v65
                                                    _ = v66
                                                    _ = v67
                                                }
                                                if true {
                                                    v68 := GenS2{F0: 2147483648, F1: v4}
                                                    if (v5 < ((("foo bar" + "") + (v5 + v5)) + v5)) {
                                                        var v69 int64 = v3
                                                        fmt.Println(v1, v69, v69)
                                                        _ = v69
                                                    } else {
                                                        switch 0 {
                                                        case 0:
                                                            var v70 int = 2147483647
                                                            _ = v70
                                                        case 2147483647:
                                                            _ = len(v2)
                                                        case 1:
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v71 int = (fi29 + (fi49 + 128))
                                                                _ = v71
                                                            }
                                                        default:
                                                            v28 += v28
                                                        }
                                                        var v72 uint32 = 948
                                                        _ = v72
                                                    }
                                                    _ = v68
                                                } else {
                                                    switch 9223372036854775807 {
                                                    case 1:
                                                        v73 := GenS8{F0: (v3 >> ((v28) & 63)), F1: v4}
                                                        _ = v73
                                                    case 316:
                                                        switch (("foo bar" + v5) + (v5 + v5)) {
                                                        case "foo bar":
                                                            v5 = v5
                                                        case "a":
                                                            v4 = (v4 << ((v28) & 63))
                                                        }
                                                    default:
                                                        delete(v2, fi49)
                                                    }
                                                    if (!(v5 < v5)) {
                                                        var v74 string = ""
                                                        v75 := GenS9{F0: 32767}
                                                        var v76 int16 = -311
                                                        _ = v74
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                    v48 = (v45 / v48)
                                                    delete(v2, 0)
                                                    fmt.Println(fi49, v5)
                                                }
                                                var v77 float32 = -2.25
                                                _ = v77
                                            case "foo bar":
                                                if len(v6) > 0 {
                                                    v6[0] = 720
                                                }
                                                v45 -= (3.14159 * (((float32((0.0 + 0.0)) * v45) * float32(1.5)) + v48))
                                                fmt.Println(v4, v5, v48)
                                            default:
                                                v78 := GenS10{F0: 127}
                                                _ = v78
                                            }
                                            _ = fi49
                                        }
                                    }
                                    fmt.Println(v4, v48, v45)
                                    _ = v48
                                }
                                for fi79 := 0; fi79 < 6; fi79++ {
                                    fmt.Println(v5, v3, fi29)
                                    switch fi29 {
                                    case -498:
                                        v6 = append(v6, 0)
                                    case -185:
                                        var v80 float32 = -2.25
                                        _ = v80
                                    case -2147483648:
                                        var v81 float64 = 3.14159
                                        _ = v81
                                    default:
                                        var v82 int32 = 1
                                        _ = v82
                                    }
                                    fmt.Println(v45, v3)
                                    v6 = append(v6, 65535)
                                    _ = fi79
                                }
                            }
                        }
                        v5 = ("" + v5)
                        _ = v45
                    } else {
                        if len(v6) > 0 {
                            v6[0] = 0
                        }
                        var v83 uint = v28
                        switch int(v3) {
                        case 42:
                            delete(v2, (fi29 - (fi29 % (((fi29 << ((25) & 63))) | 1))))
                        case -2147483648:
                            v83 = 18446744073709551615
                        case 32768:
                            v84 := GenS9{F0: 809}
                            _ = v84
                        }
                        _ = v83
                    }
                }
                for fi85 := 0; fi85 < 7; fi85++ {
                    var v86 bool = (v4 > v4)
                    v4 = v4
                    var v87 float64 = float64(((-2.25 / (((1.5 / 3.14159) * (0.0 - (0.0 - 0.1))) + ((-2.25 * 0.0) - 0.1))) - 1e10))
                    _ = fi85
                    _ = v86
                    _ = v87
                }
            }
            var v88 float32 = (float32(float64((1e10 * 1.5))) - 1e10)
            switch (("" + "") + v5) {
            case "":
                var v89 string = v5
                _ = v89
            case "\n\t":
                var v90 uint = v28
                _ = v90
            case "hello":
                var v91 int = (fi29 & 127)
                _ = v91
            }
            _ = fi29
            _ = v88
        }
        v1 += (v1 * (v1 / ((42) | 1)))
    } else {
        delete(v2, 32768)
        v92 := v2[-9223372036854775808]
        _ = v92
    }
    v93 := GenS11{F0: 0, F1: 35, F2: true}
    if len(v6) > 0 {
        v6[0] = 557
    }
    switch 127 {
    case -1:
        var v94 uint8 = 255
        _ = v94
    case -793:
        for fi95 := 0; fi95 < 16; fi95++ {
            v1 = (v1 / ((genFn0()) | 1))
            var v96 int = 987
            _ = fi95
            _ = v96
        }
    }
    v97 := GenS12{F0: 675, F1: -1}
    v3 -= 0
    var v98 uint64 = (v1 / ((4294967295) | 1))
    if (!(true || true)) {
        if (-2.25 >= (((0.1 / float32(float64(0.1))) - 0.0) / 1e10)) {
            v99 := GenS8{F0: v3, F1: v4}
            for fi100 := 0; fi100 < 20; fi100++ {
                var v101 int16 = 1
                switch v5 {
                case "\n\t":
                    if true {
                        if (v5 == "a") {
                            var v102 uint64 = v1
                            v103 := GenS3{F0: v5, F1: v5}
                            v28 = (^(v28))
                            _ = v102
                            _ = v103
                        }
                        var v104 float64 = 1e10
                        _ = v104
                    }
                case "a":
                    for fi105 := 0; fi105 < 11; fi105++ {
                        v4 += v101
                        _ = fi105
                    }
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = fi100
                _ = v101
            }
            _ = v99
        } else {
            for fi106 := 0; fi106 < 13; fi106++ {
                v107 := GenS13{F0: (((1.5 - ((0.0 - 1.5) - -2.25)) * 1e10) != float32((((1.5 / (3.14159 - -2.25)) - float64(3.14159)) - (1.5 * 3.14159)))), F1: 0}
                v108 := GenS14{F0: 149, F1: v5, F2: (v98 ^ (^(v1)))}
                v109 := GenS14{F0: 167, F1: "\n\t", F2: v98}
                _ = fi106
                _ = v107
                _ = v108
                _ = v109
            }
            v4 = 1000
            var v110 string = "hello"
            _ = v110
        }
        var v111 uint16 = 65535
        _ = v111
    } else {
        switch 4294967295 {
        case -9223372036854775808:
            var v112 int = -9223372036854775808
            _ = v112
        case -940:
            v113 := GenS3{F0: v5, F1: v5}
            _ = v113
        case -433:
            _ = len(v6)
        }
        for fi114 := 0; fi114 < 13; fi114++ {
            v115 := GenS11{F0: 1, F1: v3, F2: (885 == -1)}
            _ = fi114
            _ = v115
        }
    }
    var v116 uint16 = 32767
    fmt.Println(v98, v3, v1)
    for fi117 := 0; fi117 < 1; fi117++ {
        v118 := GenS15{F0: ("a" + ""), F1: 2147483647}
        if true {
            var v119 int64 = 187
            var v120 int16 = v4
            _ = v119
            _ = v120
        } else {
            var v121 uint64 = v1
            var v122 bool = (int32(v121) >= 42)
            _ = v121
            _ = v122
        }
        switch (fi117 / (((fi117 << ((0) & 63))) | 1)) {
        case -655:
            if false {
                v123 := GenS5{F0: 196, F1: 2}
                _ = v123
            }
        case -9223372036854775808:
            v124 := GenS9{F0: 158}
            _ = v124
        case -1:
            v4 -= 32767
        default:
            if false {
                v125 := GenS16{F0: 1}
                for fi126 := 0; fi126 < 4; fi126++ {
                    var v127 uint32 = 182
                    _ = fi126
                    _ = v127
                }
                v128 := GenS8{F0: (v3 * -1), F1: v4}
                _ = v125
                _ = v128
            } else {
                if false {
                    for fi129 := 0; fi129 < 20; fi129++ {
                        var v130 int8 = -128
                        _ = fi129
                        _ = v130
                    }
                }
                v1 -= 871
            }
        }
        _ = fi117
        _ = v118
    }
    v98 -= (v98 - (^(v1)))
    v6 = append(v6, 2)
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v116
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v93
    _ = v97
    _ = v98
}
