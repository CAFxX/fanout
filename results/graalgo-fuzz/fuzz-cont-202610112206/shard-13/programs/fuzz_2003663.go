package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16; F1 bool; F2 int16 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8; F1 int64; F2 uint }
type GenS13 struct { F0 uint; F1 string }
type GenS14 struct { F0 bool; F1 int32; F2 uint8 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint; F1 uint8; F2 uint }
type GenS18 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS19 struct { F0 uint32; F1 uint32 }
type GenS20 struct { F0 string }
type GenS21 struct { F0 uint16; F1 uint8; F2 bool }
func genFn0() int16 {
    v6 := GenS1{F0: 2}
    var v7 bool = (2 != 957)
    for fi8 := 0; fi8 < 8; fi8++ {
        v7 = v7
        fmt.Println(fi8, v7)
        var v9 uint16 = 255
        _ = fi8
        _ = v9
    }
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint32 = 0
    var v3 uint = 255
    v4 := []uint32{v2, v2}
    v5 := make(map[int]int)
    var v10 float32 = float32(0.0)
    for fi11 := 0; fi11 < 15; fi11++ {
        fmt.Println(fi11, v3)
        if true {
            if (((0 == 255) || true) && ((false && (false || ((-2.25 * 1e10) != 1e10))) && (1.5 <= v10))) {
                var v12 float64 = float64(v10)
                _ = v12
            } else {
                for fi13 := 0; fi13 < 15; fi13++ {
                    v3 = v3
                    _ = fi13
                }
                v14 := GenS2{F0: (fi11 % ((fi11) | 1)), F1: false}
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 uint32 = v2
                    v3 = 791
                    if ((2147483647 >= 854) && (9223372036854775807 == v1)) {
                        v17 := GenS2{F0: fi11, F1: true}
                        switch "hello" {
                        case "\n\t":
                            v2 = (v2 ^ (v2 % ((v2) | 1)))
                        case "hello":
                            v18 := GenS1{F0: (v3 ^ (v3 >> ((2) & 63)))}
                            _ = v18
                        }
                        _ = v17
                    }
                    if false {
                        var v19 int32 = 0
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if (127 <= 221) {
                            v20 := GenS2{F0: (fi15 ^ (fi11 >> ((2) & 63))), F1: (uint16(v16) < 371)}
                            v21 := GenS3{F0: -687, F1: v1}
                            _ = v20
                            _ = v21
                        } else {
                            v2 = v16
                            var v22 int32 = 42
                            var v23 float32 = 1e10
                            _ = v22
                            _ = v23
                        }
                        _ = v19
                    }
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
        } else {
            if false {
                if false {
                    fmt.Println(v10, v2, v1)
                }
                var v24 int = fi11
                var v25 int8 = 0
                _ = v24
                _ = v25
            } else {
                var v26 uint32 = v2
                _ = v26
            }
            v4 = append(v4, (v2 | 2))
            v2 = v2
        }
        _ = fi11
    }
    _ = len(v4)
    if (!true) {
        if true {
            var v27 uint16 = 2
            v28 := GenS3{F0: 1, F1: (v1 ^ (v1 * v1))}
            if true {
                v29 := GenS4{F0: 0}
                _ = v29
            }
            _ = v27
            _ = v28
        } else {
            v3 = (v3 >> ((203) & 63))
        }
        v30 := GenS5{F0: v1}
        if false {
            if false {
                v3 += v3
                if (((3.14159 + (((0.0 + 3.14159) / 1.5) * 1e10)) + (1e10 / (0.1 - (0.0 + float64(v10))))) > 1.5) {
                    v31 := GenS6{F0: 293, F1: true, F2: 0}
                    if (!((^(v2)) <= (v2 & 906))) {
                        for fi32 := 0; fi32 < 14; fi32++ {
                            v2 += (^(v2))
                            _ = fi32
                        }
                        v2 = v2
                    }
                    _ = v31
                }
                v33 := GenS7{F0: 501, F1: 661}
                _ = v33
            }
            if len(v4) > 0 {
                v4[0] = v2
            }
        } else {
            v34 := GenS6{F0: -286, F1: true, F2: 1}
            v1 = (v1 | -1)
            for fi35 := 0; fi35 < 11; fi35++ {
                for fi36 := 0; fi36 < 9; fi36++ {
                    var v37 uint8 = 2
                    if false {
                        var v38 uint32 = 649
                        fmt.Println(v1, v1, fi36)
                        v4 = append(v4, v38)
                        _ = v38
                    } else {
                        v2 = uint32(v3)
                    }
                    v1 -= (v1 & v1)
                    _ = fi36
                    _ = v37
                }
                if true {
                    v39 := GenS2{F0: (-(fi35)), F1: (!(1 != 846))}
                    if false {
                        fmt.Println(v3)
                        if (fi35 != fi35) {
                            for fi40 := 0; fi40 < 6; fi40++ {
                                v41 := GenS8{F0: 2147483647, F1: v2}
                                if len(v4) > 0 {
                                    v4[0] = v2
                                }
                                v3 = (v3 | v3)
                                _ = len(v5)
                                var v42 uint32 = v2
                                v43 := GenS9{F0: 127, F1: 1}
                                _ = fi40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            fmt.Println(v1, v2)
                        } else {
                            if false {
                                var v44 float32 = v10
                                if false {
                                    if false {
                                        if true {
                                            v4 = append(v4, (^(v2)))
                                        } else {
                                            v1 = 42
                                        }
                                        v1 = (v1 << ((v3) & 63))
                                    }
                                    v45 := GenS4{F0: 100}
                                    v10 = v44
                                    _ = v45
                                } else {
                                    if true {
                                        var v46 int64 = 0
                                        _ = v46
                                    }
                                    for fi47 := 0; fi47 < 19; fi47++ {
                                        var v48 int32 = 2147483647
                                        _ = fi47
                                        _ = v48
                                    }
                                    v4 = append(v4, (v2 - (v2 + 292)))
                                }
                                v49 := GenS1{F0: v3}
                                _ = v44
                                _ = v49
                            } else {
                                var v50 uint64 = 4294967295
                                v2 = v2
                                _ = v50
                            }
                            for fi51 := 0; fi51 < 12; fi51++ {
                                v1 -= int64(v3)
                                if (!true) {
                                    v52 := GenS10{F0: 42}
                                    if len(v4) > 0 {
                                        v4[0] = v2
                                    }
                                    if false {
                                        v3 += (v3 / ((v3) | 1))
                                    }
                                    _ = v52
                                }
                                var v53 int16 = 32767
                                _ = fi51
                                _ = v53
                            }
                            switch (-(fi35)) {
                            case 1:
                                for fi54 := 0; fi54 < 3; fi54++ {
                                    v55 := GenS6{F0: -500, F1: false, F2: -122}
                                    switch "\n\t" {
                                    case "":
                                        v56 := GenS11{F0: -171}
                                        _ = v56
                                    default:
                                        v57 := GenS4{F0: -99}
                                        _ = v57
                                    }
                                    var v58 int64 = 9223372036854775807
                                    _ = fi54
                                    _ = v55
                                    _ = v58
                                }
                            case -2147483648:
                                if true {
                                    v3 = (v3 * v3)
                                } else {
                                    var v59 uint32 = (^(v2))
                                    if (1 > 1) {
                                        var v60 float64 = ((0.1 * (((-2.25 + 3.14159) / 1.0) - 0.1)) - (1.5 * float64(v10)))
                                        for fi61 := 0; fi61 < 14; fi61++ {
                                            if len(v4) > 0 {
                                                v4[0] = 65535
                                            }
                                            v62 := GenS12{F0: 245, F1: -99, F2: 874}
                                            _ = fi61
                                            _ = v62
                                        }
                                        _ = v60
                                    } else {
                                        var v63 int64 = (v1 * (v1 & 996))
                                        _ = v63
                                    }
                                    _ = v59
                                }
                            }
                        }
                    } else {
                        _ = len(v4)
                        var v64 uint64 = 467
                        v65 := GenS13{F0: (v3 % (((^(v3))) | 1)), F1: (("hello" + "") + "a")}
                        _ = v64
                        _ = v65
                    }
                    var v66 uint32 = 2147483648
                    _ = v39
                    _ = v66
                } else {
                    for fi67 := 0; fi67 < 2; fi67++ {
                        var v68 uint32 = v2
                        v2 = v68
                        v69 := GenS14{F0: true, F1: -1, F2: 207}
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                    for fi70 := 0; fi70 < 14; fi70++ {
                        var v71 int32 = 958
                        fmt.Println(fi35, fi70)
                        _ = fi70
                        _ = v71
                    }
                    v72 := GenS15{F0: (false && false)}
                    _ = v72
                }
                _ = fi35
            }
            _ = v34
        }
        _ = v30
    } else {
        if ((((-2.25 * (((1.5 + (float64(1e10) - (((0.1 * (-2.25 - (3.14159 * 1.5))) * 0.0) + -2.25))) + (0.0 + (3.14159 * (1.5 * -2.25)))) - ((0.1 + float64(v10)) - 0.1))) * (3.14159 / ((1e10 + (float64(1.5) / (0.0 - 0.1))) / ((-2.25 * ((1e10 * 1e10) / 1.5)) + (1.5 / 1.0))))) * ((-2.25 - 0.1) / 1e10)) != (float64((0.1 / 1.0)) - float64(v10))) {
            v3 += v3
            var v73 float64 = -2.25
            _ = v73
        }
    }
    for fi74 := 0; fi74 < 12; fi74++ {
        v2 += v2
        v75 := GenS4{F0: -76}
        _ = fi74
        _ = v75
    }
    for fi76 := 0; fi76 < 16; fi76++ {
        v77 := v5[fi76]
        var v78 uint64 = 383
        _ = fi76
        _ = v77
        _ = v78
    }
    v1 = v1
    v79 := GenS16{F0: 1}
    var v80 int = 235
    if (2147483648 <= v2) {
        v81 := GenS16{F0: -1}
        if false {
            v4 = append(v4, (^(v2)))
            if ((v80 << ((v3) & 63)) >= 128) {
                var v82 int = 103
                _ = v82
            }
            fmt.Println(v80, v2, v3)
        } else {
            for fi83 := 0; fi83 < 11; fi83++ {
                v3 += 915
                _ = fi83
            }
            v80 = (v80 / (((v80 * v80)) | 1))
            fmt.Println(v1, v10, v10)
        }
        _ = v81
    } else {
        switch "foo bar" {
        case "":
            var v84 int16 = 470
            _ = v84
        case "\n\t":
            delete(v5, (v80 / (((^(v80))) | 1)))
        }
    }
    var v85 string = ""
    if true {
        fmt.Println(v80)
    }
    var v86 uint = v3
    switch v85 {
    case "":
        switch (^(v80)) {
        case 128:
            switch ("hello" + (v85 + v85)) {
            case "hello":
                var v87 int8 = 42
                _ = v87
            case "\n\t":
                var v88 int16 = -1
                _ = v88
            case "foo bar":
                v89 := GenS17{F0: (v3 ^ 255), F1: 255, F2: v3}
                _ = v89
            }
        case -86:
            delete(v5, 1)
        case 516:
            if false {
                v85 = (v85 + ("a" + v85))
                var v90 float64 = ((1e10 / 3.14159) - 1e10)
                if (1 >= (v2 | 834)) {
                    v90 = v90
                    switch v85 {
                    case "foo bar":
                        var v91 int8 = 100
                        _ = v91
                    case "a":
                        v1 += 42
                    case "\n\t":
                        for fi92 := 0; fi92 < 4; fi92++ {
                            var v93 float64 = v90
                            fmt.Println(v1, v1)
                            v85 = (((((v85 + ((((v85 + (v85 + v85)) + v85) + v85) + v85)) + v85) + v85) + v85) + "foo bar")
                            _ = fi92
                            _ = v93
                        }
                    }
                }
                _ = v90
            } else {
                switch 9223372036854775807 {
                case 42:
                    for fi94 := 0; fi94 < 14; fi94++ {
                        v10 += v10
                        if true {
                            var v95 uint16 = 42
                            var v96 bool = true
                            v86 += (v3 ^ 319)
                            _ = v95
                            _ = v96
                        } else {
                            switch (fi94 & (v80 - (^(fi94)))) {
                            case -2147483648:
                                fmt.Println(v3, v1)
                            case 32768:
                                var v97 int16 = 1000
                                _ = v97
                            }
                            var v98 int = v80
                            var v99 uint64 = 4294967295
                            _ = v98
                            _ = v99
                        }
                        _ = fi94
                    }
                case 4294967295:
                    var v100 uint8 = 2
                    _ = v100
                }
                v101 := GenS17{F0: 1, F1: 0, F2: (^(v86))}
                for fi102 := 0; fi102 < 15; fi102++ {
                    fmt.Println(v85)
                    v103 := GenS18{F0: v2, F1: 2, F2: v2}
                    if true {
                        switch v85 {
                        case "\n\t":
                            v2 = (v2 | v2)
                        case "foo bar":
                            if ((v1 << ((837) & 63)) != (v1 % ((v1) | 1))) {
                                var v104 float32 = v10
                                if (!(77 < 1)) {
                                    fmt.Println(v85)
                                    v1 -= v1
                                }
                                _ = v104
                            }
                        case "hello":
                            var v105 string = "foo bar"
                            _ = v105
                        }
                    } else {
                        for fi106 := 0; fi106 < 18; fi106++ {
                            switch "foo bar" {
                            case "a":
                                var v107 float32 = 0.1
                                _ = v107
                            }
                            if len(v4) > 0 {
                                v4[0] = v2
                            }
                            v86 += (v3 & 63)
                            _ = fi106
                        }
                        if false {
                            fmt.Println(v80, fi102, v85)
                            if (true || true) {
                                _ = len(v5)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v85 = ""
                            }
                            for fi108 := 0; fi108 < 13; fi108++ {
                                v109 := GenS4{F0: -128}
                                fmt.Println(v86, v86, v85)
                                _ = fi108
                                _ = v109
                            }
                        } else {
                            switch (fi102 | (-(fi102))) {
                            case -9223372036854775808:
                                v110 := GenS19{F0: (v2 << (((v86 & (v3 ^ v86))) & 63)), F1: (v2 + (v2 % (((v2 * 1)) | 1)))}
                                _ = v110
                            case -925:
                                fmt.Println(v86)
                            }
                        }
                        switch (v85 + "\n\t") {
                        case "hello":
                            fmt.Println(fi102)
                        case "\n\t":
                            for fi111 := 0; fi111 < 17; fi111++ {
                                var v112 int32 = 116
                                _ = fi111
                                _ = v112
                            }
                        default:
                            v113 := GenS14{F0: ((v86 >> ((v3) & 63)) <= (v3 + v86)), F1: 0, F2: 2}
                            _ = v113
                        }
                    }
                    _ = fi102
                    _ = v103
                }
                _ = v101
            }
        }
    default:
        fmt.Println(v1)
    }
    v114 := GenS20{F0: v85}
    if (!(42 >= 252)) {
        var v115 int8 = 127
        for fi116 := 0; fi116 < 10; fi116++ {
            fmt.Println(v3)
            v117 := GenS21{F0: 1, F1: 2, F2: false}
            _ = fi116
            _ = v117
        }
        _ = v115
    }
    fmt.Println(v2, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v80:", v80)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v10
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v79
    _ = v80
    _ = v85
    _ = v86
}
