package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 string }
type GenS2 struct { F0 int16; F1 int16; F2 int }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int64; F1 int; F2 int }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int32; F1 uint64; F2 bool }
type GenS8 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS9 struct { F0 int16; F1 uint8; F2 int64 }
type GenS10 struct { F0 uint16; F1 uint16 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int; F1 int8; F2 int64 }
type GenS15 struct { F0 int64; F1 uint }
type GenS16 struct { F0 uint32; F1 bool; F2 int }
func genFn0(p0 int16, p1 uint64, p2 int16) bool {
    v4 := []int{404, -714}
    p0 -= p0
    fmt.Println(p0, p1, p0)
    p1 -= (p1 << ((0) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (1e10 >= 3.14159)
}
func genFn1(p0 int16) uint32 {
    v5 := GenS1{F0: (p0 * (p0 / (((p0 * (p0 + (p0 ^ p0)))) | 1))), F1: 151, F2: ""}
    p0 = p0
    p0 = (p0 & p0)
    v6 := []int{345, 32767, -9223372036854775808, 9223372036854775807}
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    var v7 uint8 = 2
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint32 = 2147483648
        if true {
            fmt.Println(v7, v7, v7)
            var v10 uint8 = 42
            if true {
                delete(v2, fi8)
            } else {
                v11 := GenS2{F0: -1, F1: -733, F2: (fi8 ^ fi8)}
                switch "hello" {
                case "\n\t":
                    var v12 float64 = ((3.14159 / 1.0) + 3.14159)
                    _ = v12
                case "a":
                    v1 -= (v7 ^ (^(v10)))
                }
                _ = v11
            }
            _ = v10
        }
        for fi13 := 0; fi13 < 13; fi13++ {
            _ = len(v3)
            if false {
                fmt.Println(v9)
                var v14 int = fi8
                v2[fi8] = false
                _ = v14
            }
            if (!(false && true)) {
                switch "foo bar" {
                case "":
                    v15 := GenS2{F0: -1, F1: -624, F2: (^(fi8))}
                    _ = v15
                case "foo bar":
                    fmt.Println(v1)
                }
                fmt.Println(fi8, fi8)
            } else {
                if (!false) {
                    switch fi13 {
                    case 2147483647:
                        v16 := []int{(fi8 % ((fi8) | 1))}
                        _ = v16
                    case 32768:
                        if (42 < 69) {
                            v1 += (v1 - (^(v1)))
                        }
                        var v17 bool = true
                        fmt.Println(v9)
                        _ = v17
                    case -9223372036854775808:
                        v18 := GenS3{F0: -111, F1: v9}
                        _ = v18
                    default:
                        if (0.0 >= float64(0.1)) {
                            switch ("" + ("" + (("" + "foo bar") + "a"))) {
                            case "hello":
                                if true {
                                    if true {
                                        v9 = (v9 >> ((65535) & 63))
                                    }
                                    v1 = v1
                                    v9 += 1
                                } else {
                                    var v19 uint32 = v9
                                    var v20 uint8 = 1
                                    fmt.Println(v9)
                                    _ = v19
                                    _ = v20
                                }
                            case "":
                                fmt.Println(fi13)
                            case "a":
                                v7 = 1
                            default:
                                if false {
                                    var v21 int32 = -149
                                    var v22 int64 = 9223372036854775807
                                    if true {
                                        switch (^(fi13)) {
                                        case -9223372036854775808:
                                            fmt.Println(v7, v22, v22)
                                        case 4294967295:
                                            switch 2147483647 {
                                            case 32767:
                                                if (1000 == 0) {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        v23 := GenS3{F0: int8(v21), F1: 1}
                                                        _ = v23
                                                    case "foo bar":
                                                        v1 += v1
                                                    case "a":
                                                        v7 = (v1 << ((1) & 63))
                                                    default:
                                                        v24 := GenS3{F0: 42, F1: v9}
                                                        _ = v24
                                                    }
                                                    if ((-1 < -32768) && true) {
                                                        fmt.Println(fi13)
                                                        fmt.Println(v9, v22)
                                                        var v25 uint = 4294967295
                                                        _ = v25
                                                    } else {
                                                        v9 += (v9 * v9)
                                                        v21 -= 13
                                                    }
                                                } else {
                                                    var v26 int8 = 16
                                                    _ = v26
                                                }
                                            case 0:
                                                if (((3.14159 >= -2.25) && true) || true) {
                                                    var v27 float32 = 1e10
                                                    if false {
                                                        var v28 float32 = 0.1
                                                        var v29 uint32 = genFn1(-32768)
                                                        _ = v28
                                                        _ = v29
                                                    } else {
                                                        switch "" {
                                                        case "":
                                                            delete(v3, "foo bar")
                                                        case "\n\t":
                                                            if true {
                                                                fmt.Println(v22, v21)
                                                                if ((v21 & (-(v21))) <= int32(v9)) {
                                                                    v30 := []int{(-(fi8)), (fi13 | (fi8 << ((18446744073709551615) & 63))), 32768}
                                                                    v9 -= (v9 & 353)
                                                                    var v31 int = -1
                                                                    if true {
                                                                        v32 := GenS3{F0: 1, F1: v9}
                                                                        _ = len(v30)
                                                                        switch 0 {
                                                                        case -31:
                                                                            v22 = 9223372036854775807
                                                                        case 32768:
                                                                            if true {
                                                                                v31 = (^(v31))
                                                                                var v33 bool = (!((v31 & 0) >= v31))
                                                                                _ = v33
                                                                            }
                                                                            if ((1 <= v9) && ((false || (!((^(v21)) < v21))) && false)) {
                                                                                fmt.Println(fi13, v22, v22)
                                                                                var v34 uint = 1
                                                                                _ = len(v2)
                                                                                var v35 int = fi8
                                                                                v21 += (v21 << ((4294967295) & 63))
                                                                                _ = v34
                                                                                _ = v35
                                                                            }
                                                                        }
                                                                        _ = v32
                                                                    }
                                                                    switch ("foo bar" + "a") {
                                                                    case "hello":
                                                                        _ = len(v30)
                                                                    case "foo bar":
                                                                        var v36 uint16 = 32767
                                                                        _ = v36
                                                                    case "":
                                                                        if true {
                                                                            v27 = 3.14159
                                                                            v37 := GenS4{F0: v22, F1: int(v22), F2: (fi13 * (fi8 & (^(fi13))))}
                                                                            _ = v37
                                                                        } else {
                                                                            var v38 uint = 18446744073709551615
                                                                            switch fi8 {
                                                                            case -509:
                                                                                v39 := GenS1{F0: -490, F1: 1, F2: "a"}
                                                                                _ = v39
                                                                            case -888:
                                                                                switch fi13 {
                                                                                case 2147483647:
                                                                                    v7 = v1
                                                                                case 4294967295:
                                                                                    switch fi13 {
                                                                                    case 127:
                                                                                        var v40 float32 = float32((1e10 - (-2.25 - (1.5 / 1.0))))
                                                                                        _ = v40
                                                                                    default:
                                                                                        v2[-9223372036854775808] = ((!true) || true)
                                                                                    }
                                                                                default:
                                                                                    v1 = (v7 % (((v7 / (((v7 % ((v1) | 1))) | 1))) | 1))
                                                                                }
                                                                            case 948:
                                                                                v9 += (v9 / (((v9 | (v9 & v9))) | 1))
                                                                            default:
                                                                                var v41 float64 = 1e10
                                                                                _ = v41
                                                                            }
                                                                            fmt.Println(v7, v9)
                                                                            _ = v38
                                                                        }
                                                                    default:
                                                                        var v42 int32 = v21
                                                                        _ = v42
                                                                    }
                                                                    fmt.Println(v21)
                                                                    _ = v30
                                                                    _ = v31
                                                                }
                                                            }
                                                            fmt.Println(v7, fi13)
                                                            var v43 uint = 2
                                                            _ = v43
                                                        case "foo bar":
                                                            fmt.Println(v9, v7)
                                                        }
                                                        v44 := []int{455, -1, -165}
                                                        fmt.Println(fi13, v1)
                                                        var v45 uint64 = 1
                                                        _ = v44
                                                        _ = v45
                                                    }
                                                    var v46 uint = 416
                                                    _ = v27
                                                    _ = v46
                                                } else {
                                                    v22 = v22
                                                    v47 := GenS5{F0: 0, F1: -1}
                                                    _ = v47
                                                }
                                            case -234:
                                                fmt.Println(v9, v9, v1)
                                            default:
                                                if false {
                                                    if false {
                                                        v9 = (v9 + v9)
                                                    }
                                                    fmt.Println(fi8, v22, v7)
                                                    switch fi13 {
                                                    case -2147483648:
                                                        fmt.Println(fi8)
                                                    case 9223372036854775807:
                                                        var v48 uint64 = 42
                                                        _ = v48
                                                    case 758:
                                                        v7 -= (v7 % ((v7) | 1))
                                                    default:
                                                        v49 := GenS4{F0: -559, F1: (fi8 - 127), F2: fi8}
                                                        _ = v49
                                                    }
                                                } else {
                                                    switch fi13 {
                                                    case 127:
                                                        v22 = v22
                                                    }
                                                    v9 = v9
                                                }
                                            }
                                        case 0:
                                            switch 277 {
                                            case -2147483648:
                                                v50 := GenS2{F0: 42, F1: 308, F2: fi13}
                                                _ = v50
                                            case 2147483647:
                                                v9 -= v9
                                            }
                                        }
                                        fmt.Println(fi13)
                                        v51 := []int{fi13}
                                        _ = v51
                                    } else {
                                        v52 := v3[""]
                                        if true {
                                            delete(v3, ("hello" + ("" + (("hello" + ("\n\t" + ("hello" + (("a" + ("\n\t" + ("foo bar" + "hello"))) + "")))) + "a"))))
                                        } else {
                                            v53 := GenS2{F0: -32768, F1: 1000, F2: fi13}
                                            var v54 int64 = (^(v22))
                                            switch int(v21) {
                                            case 134:
                                                switch fi8 {
                                                case 1:
                                                    v52 = v52
                                                default:
                                                    _ = len(v2)
                                                }
                                            case 106:
                                                v21 = v21
                                            }
                                            var v55 int8 = -128
                                            v56 := []int{}
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                            _ = v56
                                        }
                                        if (!false) {
                                            v9 = (v9 * v9)
                                            var v57 float32 = 0.0
                                            switch (fi13 & fi13) {
                                            case -517:
                                                if (-2.25 == 0.1) {
                                                    if true {
                                                        if (!((v1 + v7) < 254)) {
                                                            var v58 uint64 = 58
                                                            v22 = v22
                                                            if true {
                                                                var v59 int16 = -803
                                                                v52 = (fi13 % (((fi8 * -304)) | 1))
                                                                if ((v22 / ((60) | 1)) < (v22 % ((42) | 1))) {
                                                                    v22 = v22
                                                                    switch v52 {
                                                                    case 32768:
                                                                        var v60 int32 = v21
                                                                        _ = v60
                                                                    case -2147483648:
                                                                        if true {
                                                                            v61 := []int{1, fi8, (-(fi8)), (fi13 | int(v59))}
                                                                            _ = v61
                                                                        } else {
                                                                            if true {
                                                                                var v62 uint32 = (v9 << ((1) & 63))
                                                                                if ("\n\t" > "\n\t") {
                                                                                    v63 := GenS6{F0: -12}
                                                                                    v64 := v2[fi13]
                                                                                    _ = v63
                                                                                    _ = v64
                                                                                }
                                                                                var v65 uint8 = v7
                                                                                _ = v62
                                                                                _ = v65
                                                                            }
                                                                            switch ("a" + "foo bar") {
                                                                            case "hello":
                                                                                var v66 int64 = 9223372036854775807
                                                                                _ = v66
                                                                            case "":
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    var v67 string = (("\n\t" + "hello") + (("foo bar" + (("" + "") + ("a" + "a"))) + ("\n\t" + "hello")))
                                                                                    _ = v67
                                                                                case "\n\t":
                                                                                    switch (-(fi13)) {
                                                                                    case -473:
                                                                                        switch fi8 {
                                                                                        case 128:
                                                                                            fmt.Println(v9, v9)
                                                                                        case 910:
                                                                                            v68 := GenS7{F0: (-(v21)), F1: v58, F2: true}
                                                                                            _ = v68
                                                                                        case 281:
                                                                                            fmt.Println(v21, fi8, v59)
                                                                                        default:
                                                                                            fmt.Println(v52, v58)
                                                                                        }
                                                                                    case -1:
                                                                                        var v69 string = "foo bar"
                                                                                        _ = v69
                                                                                    case 128:
                                                                                        if (!true) {
                                                                                            v1 += v7
                                                                                        } else {
                                                                                            v57 = 0.0
                                                                                            switch (fi8 >> ((65535) & 63)) {
                                                                                            case -108:
                                                                                                fmt.Println(v7, v52, v9)
                                                                                            case 32767:
                                                                                                var v70 int32 = 2147483647
                                                                                                _ = v70
                                                                                            case 42:
                                                                                                v71 := []int{fi13, (fi8 + v52), (fi13 + v52)}
                                                                                                _ = v71
                                                                                            default:
                                                                                                fmt.Println(fi8)
                                                                                            }
                                                                                            fmt.Println(v21, v57)
                                                                                        }
                                                                                        v72 := []int{}
                                                                                        _ = v72
                                                                                    }
                                                                                }
                                                                            case "a":
                                                                                if true {
                                                                                    v73 := GenS1{F0: v59, F1: 2, F2: ("" + "foo bar")}
                                                                                    v58 = (v58 & v58)
                                                                                    var v74 int8 = 100
                                                                                    _ = v73
                                                                                    _ = v74
                                                                                } else {
                                                                                    if ((v58 * uint64(v7)) >= v58) {
                                                                                        if true {
                                                                                            var v75 float64 = 0.0
                                                                                            fmt.Println(v7)
                                                                                            if false {
                                                                                                var v76 float64 = v75
                                                                                                if false {
                                                                                                    v21 += (v21 % ((int32(v9)) | 1))
                                                                                                    var v77 int32 = v21
                                                                                                    v1 -= v1
                                                                                                    _ = v77
                                                                                                }
                                                                                                _ = v76
                                                                                            } else {
                                                                                                v78 := GenS1{F0: (v59 & v59), F1: (v7 << ((65535) & 63)), F2: (("foo bar" + (("hello" + ("" + ("\n\t" + "\n\t"))) + "foo bar")) + "a")}
                                                                                                v21 += (v21 * (-(v21)))
                                                                                                _ = v78
                                                                                            }
                                                                                            _ = v75
                                                                                        } else {
                                                                                            if false {
                                                                                                var v79 uint16 = 188
                                                                                                var v80 bool = (((v57 - 0.1) - ((((((0.0 * v57) - v57) + float32(float64(1e10))) + 0.1) / v57) + 1e10)) <= (float32(0.1) * float32(-2.25)))
                                                                                                var v81 float64 = 0.1
                                                                                                if false {
                                                                                                    var v82 int32 = 2147483647
                                                                                                    _ = v82
                                                                                                } else {
                                                                                                    if (v80 && true) {
                                                                                                        v3["\n\t"] = (fi8 % ((-9223372036854775808) | 1))
                                                                                                        var v83 int16 = v59
                                                                                                        _ = v83
                                                                                                    }
                                                                                                    var v84 int8 = 127
                                                                                                    var v85 int32 = v21
                                                                                                    _ = v84
                                                                                                    _ = v85
                                                                                                }
                                                                                                _ = v79
                                                                                                _ = v80
                                                                                                _ = v81
                                                                                            }
                                                                                            fmt.Println(v9, v9)
                                                                                            v86 := GenS7{F0: (v21 | int32(v7)), F1: 4294967295, F2: true}
                                                                                            _ = v86
                                                                                        }
                                                                                    } else {
                                                                                        v87 := GenS7{F0: v21, F1: v58, F2: ((false && true) && (!(32767 >= 283)))}
                                                                                        v88 := v2[fi8]
                                                                                        var v89 uint = 255
                                                                                        _ = v87
                                                                                        _ = v88
                                                                                        _ = v89
                                                                                    }
                                                                                    v90 := GenS5{F0: 255, F1: 42}
                                                                                    _ = v90
                                                                                }
                                                                            default:
                                                                                if false {
                                                                                    var v91 uint16 = 355
                                                                                    fmt.Println(v52)
                                                                                    var v92 int64 = v22
                                                                                    v93 := GenS8{F0: (v58 | (v58 << ((187) & 63))), F1: v22, F2: v1}
                                                                                    v94 := []int{}
                                                                                    if false {
                                                                                        v3["foo bar"] = (fi13 - (fi8 + 602))
                                                                                        _ = len(v94)
                                                                                    } else {
                                                                                        v22 += v22
                                                                                        v91 += v91
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            switch "\n\t" {
                                                                                            case "foo bar":
                                                                                                if len(v94) >= 2 {
                                                                                                    v94 = v94[1:len(v94)]
                                                                                                }
                                                                                            default:
                                                                                                v91 = (v91 - v91)
                                                                                            }
                                                                                        default:
                                                                                            v92 = 573
                                                                                        }
                                                                                        v95 := GenS1{F0: (v59 + v59), F1: 42, F2: ""}
                                                                                        _ = v95
                                                                                    }
                                                                                    _ = v91
                                                                                    _ = v92
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                } else {
                                                                                    switch ("\n\t" + "foo bar") {
                                                                                    case "foo bar":
                                                                                        if false {
                                                                                            v96 := GenS7{F0: v21, F1: (v58 & (v58 ^ 2)), F2: true}
                                                                                            _ = v96
                                                                                        } else {
                                                                                            v9 += (v9 * 0)
                                                                                            v22 += v22
                                                                                        }
                                                                                    case "":
                                                                                        if true {
                                                                                            v52 = fi8
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v7 -= (v1 - 255)
                                                                                    default:
                                                                                        v97 := GenS3{F0: -1, F1: (v9 * (v9 >> ((1) & 63)))}
                                                                                        fmt.Println(v21, v1, v57)
                                                                                        _ = v97
                                                                                    }
                                                                                    switch 9223372036854775807 {
                                                                                    case -9223372036854775808:
                                                                                        fmt.Println(v9, v21)
                                                                                    case -2147483648:
                                                                                        if true {
                                                                                            var v98 bool = true
                                                                                            _ = v98
                                                                                        }
                                                                                    case 4294967295:
                                                                                        v99 := GenS1{F0: (v59 | (v59 | (v59 / ((v59) | 1)))), F1: (v1 % ((v1) | 1)), F2: ("" + "hello")}
                                                                                        _ = v99
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    case 42:
                                                                        v100 := []int{(v52 + fi13)}
                                                                        _ = v100
                                                                    default:
                                                                        switch "" {
                                                                        case "":
                                                                            _ = len(v2)
                                                                        case "foo bar":
                                                                            var v101 int8 = int8(v59)
                                                                            _ = v101
                                                                        case "a":
                                                                            fmt.Println(v1, v52, fi8)
                                                                            if (v52 >= -1) {
                                                                                v22 += (v22 >> ((18446744073709551615) & 63))
                                                                            }
                                                                        }
                                                                    }
                                                                    v59 = -32768
                                                                } else {
                                                                    v102 := GenS6{F0: 44}
                                                                    v103 := GenS7{F0: -1, F1: v58, F2: ((v22 >> ((42) & 63)) <= v22)}
                                                                    v104 := GenS9{F0: 1000, F1: v1, F2: (v22 & -521)}
                                                                    if (!(!(384 < 292))) {
                                                                        _ = len(v3)
                                                                        fmt.Println(v22, v7, fi8)
                                                                    }
                                                                    var v105 string = "foo bar"
                                                                    switch (fi13 % ((v52) | 1)) {
                                                                    case -691:
                                                                        v106 := GenS9{F0: (v59 + (v59 - (^(v59)))), F1: (v7 | 75), F2: v22}
                                                                        _ = v106
                                                                    case 4294967295:
                                                                        delete(v3, v105)
                                                                    }
                                                                    _ = v102
                                                                    _ = v103
                                                                    _ = v104
                                                                    _ = v105
                                                                }
                                                                switch "hello" {
                                                                case "a":
                                                                    var v107 float32 = 1e10
                                                                    _ = v107
                                                                default:
                                                                    v108 := GenS10{F0: 65535, F1: 255}
                                                                    if (((1e10 / float64(-2.25)) * 3.14159) >= 3.14159) {
                                                                        var v109 uint32 = genFn1(v59)
                                                                        _ = v109
                                                                    } else {
                                                                        var v110 int64 = v22
                                                                        fmt.Println(fi13, v22)
                                                                        v111 := []int{(fi13 % ((0) | 1))}
                                                                        _ = v110
                                                                        _ = v111
                                                                    }
                                                                    if true {
                                                                        v9 -= genFn1((v59 + (^(v59))))
                                                                    } else {
                                                                        v22 = 0
                                                                        switch (v52 & -2147483648) {
                                                                        case -1:
                                                                            var v112 bool = false
                                                                            _ = v112
                                                                        }
                                                                    }
                                                                    _ = v108
                                                                }
                                                                v52 = (fi8 ^ 2147483647)
                                                                v113 := GenS6{F0: 31}
                                                                _ = v113
                                                                _ = v59
                                                            }
                                                            _ = v58
                                                        } else {
                                                            fmt.Println(fi13, v57, fi8)
                                                            v114 := GenS11{F0: (v52 - fi8)}
                                                            if true {
                                                                if true {
                                                                    _ = len(v2)
                                                                } else {
                                                                    var v115 string = "a"
                                                                    var v116 uint8 = v7
                                                                    if (2147483647 != 42) {
                                                                        if true {
                                                                            if ((false || false) || (((1.5 + 3.14159) + 0.0) > 0.0)) {
                                                                                var v117 uint32 = v9
                                                                                v118 := GenS8{F0: 428, F1: v22, F2: v1}
                                                                                v115 = "foo bar"
                                                                                var v119 uint64 = 42
                                                                                v120 := GenS12{F0: 128}
                                                                                _ = v117
                                                                                _ = v118
                                                                                _ = v119
                                                                                _ = v120
                                                                            }
                                                                            v121 := GenS7{F0: v21, F1: 4294967295, F2: (1 <= fi13)}
                                                                            var v122 int = v52
                                                                            _ = v121
                                                                            _ = v122
                                                                        }
                                                                        var v123 uint16 = 1
                                                                        _ = v123
                                                                    }
                                                                    _ = v115
                                                                    _ = v116
                                                                }
                                                                v22 -= 42
                                                            } else {
                                                                var v124 int8 = -91
                                                                v124 -= v124
                                                                switch fi8 {
                                                                case 2147483647:
                                                                    v125 := GenS12{F0: (v1 ^ v1)}
                                                                    _ = v125
                                                                case -478:
                                                                    v57 -= 1.5
                                                                default:
                                                                    if false {
                                                                        v7 = (v7 * (v7 ^ (v7 * v1)))
                                                                        switch ("" + "hello") {
                                                                        case "hello":
                                                                            v126 := GenS2{F0: -543, F1: 1, F2: fi8}
                                                                            _ = v126
                                                                        case "":
                                                                            var v127 int8 = v124
                                                                            _ = v127
                                                                        case "foo bar":
                                                                            switch "a" {
                                                                            case "":
                                                                                v128 := []int{v52, fi8, fi13}
                                                                                _ = v128
                                                                            }
                                                                            var v129 string = ""
                                                                            fmt.Println(v7, v21, v52)
                                                                            _ = v129
                                                                        default:
                                                                            v9 -= (v9 >> ((4294967295) & 63))
                                                                        }
                                                                        var v130 bool = (float32(((1e10 - 0.0) - 0.0)) >= v57)
                                                                        var v131 int64 = v22
                                                                        _ = v130
                                                                        _ = v131
                                                                    }
                                                                }
                                                                if (true || false) {
                                                                    var v132 int64 = (v22 >> ((4294967295) & 63))
                                                                    _ = v132
                                                                }
                                                                var v133 int = (-(fi8))
                                                                v57 -= v57
                                                                _ = v124
                                                                _ = v133
                                                            }
                                                            _ = v114
                                                        }
                                                    }
                                                    v52 -= (fi8 | (fi13 + v52))
                                                }
                                                v21 += (v21 >> ((1) & 63))
                                            case -372:
                                                if true {
                                                    if true {
                                                        v57 = v57
                                                        var v134 uint8 = v1
                                                        _ = v134
                                                    } else {
                                                        switch v52 {
                                                        case 0:
                                                            v135 := GenS13{F0: 0}
                                                            _ = v135
                                                        }
                                                        var v136 string = "a"
                                                        _ = v136
                                                    }
                                                    v137 := GenS8{F0: 656, F1: v22, F2: v7}
                                                    if (1 > -88) {
                                                        v7 += (v7 ^ (v7 ^ 127))
                                                        switch (fi8 << ((1) & 63)) {
                                                        case 4294967295:
                                                            var v138 int32 = 42
                                                            _ = v138
                                                        default:
                                                            v22 = v22
                                                            var v139 uint8 = 255
                                                            v1 += 34
                                                            _ = v139
                                                        }
                                                        var v140 int8 = -62
                                                        _ = v140
                                                    } else {
                                                        v141 := GenS14{F0: fi8, F1: 42, F2: (v22 >> ((65535) & 63))}
                                                        _ = v141
                                                    }
                                                    _ = v137
                                                } else {
                                                    fmt.Println(v52, v7, v7)
                                                    switch (fi8 / (((fi8 * (fi13 % (((^(fi8))) | 1)))) | 1)) {
                                                    case -172:
                                                        v1 = (v1 * 1)
                                                    case -2147483648:
                                                        var v142 int64 = v22
                                                        var v143 float32 = (v57 + (v57 - (v57 / 1.5)))
                                                        fmt.Println(v22, fi13, fi8)
                                                        _ = v142
                                                        _ = v143
                                                    default:
                                                        var v144 uint = 790
                                                        v145 := GenS5{F0: 65535, F1: (-(v21))}
                                                        _ = v144
                                                        _ = v145
                                                    }
                                                    v57 -= 1.5
                                                }
                                            case -712:
                                                fmt.Println(fi13)
                                            }
                                            _ = v57
                                        }
                                        _ = v52
                                    }
                                    _ = v21
                                    _ = v22
                                }
                            }
                        } else {
                            v146 := v3["a"]
                            v147 := GenS4{F0: -1, F1: -9223372036854775808, F2: (fi13 / (((fi8 >> ((0) & 63))) | 1))}
                            var v148 int = 721
                            _ = v146
                            _ = v147
                            _ = v148
                        }
                    }
                    if (1e10 != ((1e10 - 0.1) + (3.14159 + 3.14159))) {
                        v1 = (^(v7))
                        switch "a" {
                        case "a":
                            switch "" {
                            case "foo bar":
                                if false {
                                    var v149 uint = 384
                                    var v150 int8 = 127
                                    _ = v149
                                    _ = v150
                                }
                            case "hello":
                                if (2 != 0) {
                                    switch (fi8 / ((fi13) | 1)) {
                                    case 4294967295:
                                        var v151 bool = false
                                        _ = v151
                                    case -9223372036854775808:
                                        var v152 int16 = 32767
                                        _ = v152
                                    }
                                    if (!true) {
                                        var v153 int32 = 2147483647
                                        fmt.Println(fi8, v7)
                                        _ = v153
                                    }
                                    v7 += (v1 << ((65535) & 63))
                                    v1 = 127
                                }
                            default:
                                switch ("" + "hello") {
                                case "foo bar":
                                    var v154 uint64 = 42
                                    _ = v154
                                case "a":
                                    fmt.Println(fi13)
                                case "":
                                    if true {
                                        v155 := GenS7{F0: 1, F1: 1, F2: (0.0 != 1.5)}
                                        v1 += v1
                                        _ = v155
                                    } else {
                                        var v156 uint16 = 65535
                                        var v157 string = "foo bar"
                                        _ = v156
                                        _ = v157
                                    }
                                    v158 := []int{fi8, fi8}
                                    _ = v158
                                }
                            }
                            v9 -= (v9 / (((v9 << ((0) & 63))) | 1))
                            v9 -= v9
                        default:
                            if (2 < 4294967295) {
                                _ = len(v2)
                            } else {
                                v9 = v9
                                v7 = 0
                                var v159 string = (("" + ("\n\t" + ("hello" + ((("hello" + "\n\t") + "hello") + "\n\t")))) + "")
                                _ = v159
                            }
                        }
                    }
                    fmt.Println(v9, v1)
                } else {
                    if false {
                        v160 := []int{-645}
                        v1 += (^(v7))
                        v161 := v3["\n\t"]
                        _ = v160
                        _ = v161
                    } else {
                        var v162 uint32 = v9
                        v2[fi13] = (-2147483648 < 0)
                        if (fi13 == -831) {
                            v163 := GenS14{F0: (fi13 & fi13), F1: 38, F2: 93}
                            _ = v163
                        }
                        _ = v162
                    }
                }
                var v164 bool = (false && (!(((!false) && false) && (675 == 65535))))
                _ = v164
            }
            _ = fi13
        }
        _ = fi8
        _ = v9
    }
    var v165 int16 = 42
    var v166 int8 = 42
    v1 = 127
    v167 := GenS10{F0: 277, F1: 4}
    if false {
        if ((v166 & v166) > (-(v166))) {
            v165 = (v165 & (v165 | (v165 ^ (v165 % (((v165 & v165)) | 1)))))
            v168 := []int{-409, -544}
            _ = v168
        }
    }
    var v169 int16 = 0
    switch "hello" {
    case "\n\t":
        v7 = (v7 >> ((uint(v7)) & 63))
    case "":
        for fi170 := 0; fi170 < 12; fi170++ {
            if false {
                for fi171 := 0; fi171 < 8; fi171++ {
                    switch "" {
                    case "a":
                        var v172 string = ""
                        _ = v172
                    case "\n\t":
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            v165 = 32767
                        case "":
                            switch (^(fi170)) {
                            case 0:
                                var v173 string = "\n\t"
                                _ = v173
                            case -1:
                                var v174 uint64 = 0
                                _ = v174
                            default:
                                v165 = 320
                            }
                            v169 = (v169 % ((v169) | 1))
                        case "a":
                            v175 := GenS15{F0: 9223372036854775807, F1: 422}
                            _ = v175
                        default:
                            v176 := GenS14{F0: fi170, F1: (v166 ^ v166), F2: 9223372036854775807}
                            var v177 int64 = -1
                            if ((!true) || (32768 < fi171)) {
                                v2[fi171] = (v7 > v7)
                                v1 = v7
                                v169 += -935
                                switch fi171 {
                                case 894:
                                    var v178 int = fi170
                                    _ = v178
                                }
                                var v179 float64 = (1e10 * 1e10)
                                _ = v179
                            }
                            var v180 string = "\n\t"
                            v180 = v180
                            _ = v176
                            _ = v177
                            _ = v180
                        }
                    case "hello":
                        switch fi170 {
                        case -1:
                            var v181 string = ("" + "\n\t")
                            _ = v181
                        case -270:
                            var v182 bool = false
                            _ = v182
                        case -998:
                            if (false || (!false)) {
                                if (3.14159 < (0.0 * 1.5)) {
                                    v166 -= (v166 * (v166 >> ((4294967295) & 63)))
                                    v183 := []int{}
                                    _ = v183
                                }
                            }
                        default:
                            v184 := GenS1{F0: v165, F1: 127, F2: "\n\t"}
                            _ = v184
                        }
                    }
                    v185 := GenS10{F0: 65535, F1: 1}
                    switch (("hello" + "hello") + "foo bar") {
                    case "foo bar":
                        v186 := GenS15{F0: -270, F1: 4294967295}
                        var v187 bool = (-1 == -585)
                        _ = v186
                        _ = v187
                    case "a":
                        v169 = (-(v165))
                    }
                    _ = fi171
                    _ = v185
                }
                if false {
                    v3[("a" + "hello")] = 4294967295
                }
            }
            if (((-2147483648 <= -663) || ((fi170 & (fi170 - fi170)) <= (fi170 >> ((0) & 63)))) || ((1.5 + 1.5) != 3.14159)) {
                if (2 != 277) {
                    for fi188 := 0; fi188 < 20; fi188++ {
                        v1 -= (v1 / ((v7) | 1))
                        _ = fi188
                    }
                    for fi189 := 0; fi189 < 12; fi189++ {
                        fmt.Println(v1)
                        v169 -= (v165 & v165)
                        _ = fi189
                    }
                    v190 := GenS14{F0: -500, F1: (v166 * v166), F2: -620}
                    _ = v190
                }
                switch fi170 {
                case 9223372036854775807:
                    var v191 uint64 = 730
                    _ = v191
                default:
                    var v192 int = fi170
                    _ = v192
                }
                switch (-(fi170)) {
                case -2147483648:
                    switch "" {
                    case "":
                        v169 = (v165 % (((^(v165))) | 1))
                    case "\n\t":
                        fmt.Println(fi170, v7)
                    case "foo bar":
                        for fi193 := 0; fi193 < 4; fi193++ {
                            var v194 int64 = -933
                            var v195 int8 = (v166 + int8(fi193))
                            _ = fi193
                            _ = v194
                            _ = v195
                        }
                    }
                default:
                    v7 += (v7 - v7)
                }
            }
            v165 = v165
            _ = fi170
        }
    default:
        if false {
            var v196 bool = (!true)
            _ = v196
        }
    }
    if false {
        v197 := []int{9223372036854775807, 0, 0}
        _ = v197
    } else {
        delete(v3, "hello")
    }
    v198 := GenS16{F0: 693, F1: (1 != 0), F2: 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v165:", v165)
    fmt.Println("v166:", v166)
    fmt.Println("v169:", v169)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v165
    _ = v166
    _ = v167
    _ = v169
    _ = v198
    _ = v2
    _ = v3
    _ = v7
}
