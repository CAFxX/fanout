package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 bool }
type GenS2 struct { F0 int32; F1 bool; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS7 struct { F0 int16; F1 string; F2 uint8 }
type GenS8 struct { F0 int8; F1 int64; F2 uint32 }
func genFn0(p0 uint32, p1 uint16, p2 uint64) int32 {
    v7 := GenS1{F0: -721, F1: -1, F2: true}
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := []int{(fi8 | fi8)}
        p0 -= (p0 + 4294967295)
        _ = fi8
        _ = v9
    }
    if false {
        v10 := []int{}
        _ = v10
    }
    switch "a" {
    case "hello":
        v11 := GenS1{F0: 0, F1: 0, F2: (false && true)}
        _ = v11
    case "":
        for fi12 := 0; fi12 < 14; fi12++ {
            p0 += (p0 | (^(p0)))
            v13 := GenS1{F0: 781, F1: 9223372036854775807, F2: (65535 == 255)}
            _ = fi12
            _ = v13
        }
    }
    var v14 int32 = -336
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v7
    return v14
}
func genFn1(p0 uint64, p1 uint32, p2 uint32) uint32 {
    fmt.Println(p2, p2)
    v15 := []int{1}
    for fi16 := 0; fi16 < 18; fi16++ {
        if (2 >= 128) {
            var v17 uint8 = 1
            _ = v17
        } else {
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 uint8 = 176
                _ = fi18
                _ = v19
            }
            if (0 > 141) {
                p2 = 2147483648
                if (!((fi16 + (fi16 * fi16)) < -10)) {
                    if len(v15) > 0 {
                        v15[0] = fi16
                    }
                }
                switch ("hello" + "\n\t") {
                case "a":
                    v15 = append(v15, fi16)
                default:
                    if len(v15) > 0 {
                        v15[0] = (fi16 ^ 9223372036854775807)
                    }
                }
            }
            if ((!(!true)) || false) {
                var v20 string = ""
                p2 = p1
                var v21 int64 = 42
                _ = v20
                _ = v21
            }
        }
        _ = fi16
    }
    p1 += 793
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return p1
}

func main() {
    var v1 uint16 = 0
    var v2 bool = true
    var v3 uint = 1
    var v4 uint16 = v1
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]uint32)
    v1 -= (v1 - v4)
    var v22 uint16 = 307
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 int64 = 0
        v1 -= 642
        switch (fi23 / ((-9223372036854775808) | 1)) {
        case -2147483648:
            v25 := GenS2{F0: -311, F1: (!(2 == 353)), F2: ((v1 - 366) >= 0)}
            _ = v25
        default:
            switch fi23 {
            case 127:
                v24 -= 42
            }
        }
        _ = fi23
        _ = v24
    }
    var v26 uint64 = 560
    v22 += v22
    for fi27 := 0; fi27 < 17; fi27++ {
        if true {
            switch fi27 {
            case -803:
                var v28 int16 = 1000
                _ = v28
            default:
                fmt.Println(v26, v2)
            }
        }
        v26 += v26
        _ = fi27
    }
    switch ("\n\t" + "\n\t") {
    case "hello":
        switch ("hello" + "hello") {
        case "\n\t":
            if v2 {
                if v2 {
                    v22 = (v22 << ((4294967295) & 63))
                    var v29 float64 = (0.0 * 1e10)
                    var v30 float32 = (1e10 + (float32(v29) * (0.0 - float32(float64((-2.25 + -2.25))))))
                    _ = v29
                    _ = v30
                } else {
                    v31 := []int{127, 624, 2147483647, -194}
                    _ = v31
                }
                v3 += (^(v3))
                var v32 int64 = 9223372036854775807
                _ = v32
            } else {
                switch "hello" {
                case "\n\t":
                    v33 := GenS1{F0: genFn0(genFn1(v26, 2, 4294967295), v22, (v26 * (v26 * v26))), F1: 211, F2: true}
                    _ = v33
                case "hello":
                    v1 += v1
                }
                if (-2.25 > (0.1 / 0.1)) {
                    fmt.Println(v26, v26, v4)
                    v34 := GenS2{F0: genFn0(121, uint16(v3), (v26 & 366)), F1: (1.5 >= (0.0 * 1.5)), F2: v2}
                    v35 := GenS1{F0: 1, F1: 9223372036854775807, F2: (1000 > -32768)}
                    _ = v34
                    _ = v35
                } else {
                    v36 := GenS1{F0: -833, F1: -1, F2: (255 != 2)}
                    var v37 float64 = 3.14159
                    _ = v36
                    _ = v37
                }
                v38 := v6[-778]
                _ = v38
            }
        case "hello":
            var v39 uint64 = v26
            _ = v39
        default:
            v40 := []int{-278, 42}
            _ = v40
        }
    default:
        v41 := []int{32768, 2147483647}
        _ = v41
    }
    if ((!((true || (!(!(!(2 >= 255))))) && (0.0 != (((float64((3.14159 / 1e10)) + 3.14159) * 0.1) * (1.5 / 1.0))))) || ((-2.25 * (-2.25 - 1.5)) == (-2.25 * float32(float64(3.14159))))) {
        var v42 int8 = -1
        switch 32767 {
        case -2147483648:
            if true {
                for fi43 := 0; fi43 < 4; fi43++ {
                    var v44 uint = v3
                    _ = fi43
                    _ = v44
                }
                switch 91 {
                case -9223372036854775808:
                    var v45 uint = v3
                    _ = v45
                default:
                    if true {
                        if v2 {
                            v46 := []int{127, 429}
                            _ = v46
                        }
                    }
                }
            }
        case 4294967295:
            v47 := GenS1{F0: 42, F1: 1, F2: (v2 || (925 != 462))}
            _ = v47
        }
        v1 += v22
        _ = v42
    } else {
        fmt.Println(v3, v4, v2)
        if v2 {
            v2 = v2
            switch 32768 {
            case 127:
                var v48 uint = v3
                _ = v48
            case -352:
                v1 -= (v22 / ((v4) | 1))
            case 4294967295:
                delete(v5, 0)
            default:
                v49 := GenS2{F0: -582, F1: false, F2: false}
                _ = v49
            }
            var v50 float32 = 0.0
            _ = v50
        }
    }
    v51 := []int{-9223372036854775808, 42, 32768, 4294967295}
    var v52 string = "\n\t"
    v52 = ((("a" + v52) + v52) + "")
    if (((65 <= 64) || (641 == -41)) || (127 >= -9223372036854775808)) {
        if v2 {
            v52 = v52
            for fi53 := 0; fi53 < 18; fi53++ {
                v54 := GenS2{F0: -1, F1: v2, F2: ((v3 >> ((2) & 63)) < 255)}
                if (!true) {
                    v55 := GenS3{F0: (v26 - (v26 + uint64(v3)))}
                    _ = v55
                } else {
                    for fi56 := 0; fi56 < 18; fi56++ {
                        if v2 {
                            v26 -= uint64(v22)
                        }
                        if v2 {
                            if (65535 != v4) {
                                if v2 {
                                    fmt.Println(v4, v1)
                                } else {
                                    if true {
                                        fmt.Println(fi56, v3)
                                        if true {
                                            v22 = v1
                                            if (!((^(v26)) >= v26)) {
                                                v4 -= v1
                                            }
                                            if ((v3 / ((v3) | 1)) < v3) {
                                                _ = len(v51)
                                                if v2 {
                                                    v57 := GenS3{F0: (v26 + v26)}
                                                    v58 := GenS3{F0: (v26 - 547)}
                                                    switch v52 {
                                                    case "foo bar":
                                                        if (v2 || true) {
                                                            v59 := GenS3{F0: uint64(v22)}
                                                            v60 := GenS3{F0: v26}
                                                            fmt.Println(v22)
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            v61 := GenS1{F0: -2147483648, F1: -9223372036854775808, F2: true}
                                                            _ = v61
                                                        }
                                                        var v62 float32 = 1.5
                                                        _ = v62
                                                    case "":
                                                        var v63 uint = v3
                                                        _ = v63
                                                    }
                                                    _ = v57
                                                    _ = v58
                                                }
                                                var v64 uint = 65535
                                                _ = v64
                                            }
                                        }
                                        var v65 int32 = int32(v3)
                                        switch (-(fi56)) {
                                        case 872:
                                            v3 = (v3 << (((v3 % ((v3) | 1))) & 63))
                                        case 32768:
                                            fmt.Println(v52, v3)
                                            var v66 int8 = 0
                                            _ = v66
                                        case -1:
                                            if v2 {
                                                v67 := GenS1{F0: v65, F1: 42, F2: v2}
                                                var v68 uint64 = (v26 | (^(v26)))
                                                var v69 string = (v52 + ((v52 + ("\n\t" + (v52 + ("hello" + v52)))) + "hello"))
                                                _ = v67
                                                _ = v68
                                                _ = v69
                                            }
                                            switch 0 {
                                            case 1:
                                                var v70 int32 = v65
                                                _ = v70
                                            }
                                            var v71 uint8 = 1
                                            _ = v71
                                        }
                                        var v72 uint8 = 255
                                        _ = v65
                                        _ = v72
                                    } else {
                                        _ = len(v5)
                                        v26 -= 1
                                    }
                                    if (((87 <= 128) && (2147483647 != 65535)) || v2) {
                                        var v73 float64 = (-2.25 / -2.25)
                                        _ = v73
                                    }
                                }
                            } else {
                                var v74 string = ("hello" + (v52 + v52))
                                switch (v52 + "hello") {
                                case "\n\t":
                                    v52 = v74
                                case "":
                                    if true {
                                        v2 = false
                                    }
                                }
                                v3 = (^(v3))
                                _ = v74
                            }
                            _ = len(v5)
                            v75 := GenS4{F0: uint64(v3)}
                            v52 = "foo bar"
                            if (0.1 == 1e10) {
                                v4 += (v1 >> (((v3 & (v3 << ((753) & 63)))) & 63))
                                v52 = v52
                                v22 += (^(v4))
                            } else {
                                v26 -= v26
                            }
                            var v76 uint64 = v26
                            _ = v75
                            _ = v76
                        }
                        if (!((!true) || true)) {
                            v52 = (v52 + v52)
                            if (v4 < (v22 & 65535)) {
                                v26 = (v26 >> ((469) & 63))
                            } else {
                                if v2 {
                                    var v77 uint = (v3 & v3)
                                    v2 = (fi56 <= -1)
                                    if (1 < 1) {
                                        v77 -= (v3 ^ v3)
                                        v78 := v5[2147483647]
                                        var v79 int8 = 100
                                        v1 = (v4 * v1)
                                        _ = v78
                                        _ = v79
                                    }
                                    v22 = (^(v22))
                                    _ = v77
                                } else {
                                    fmt.Println(v26, v1, v3)
                                }
                            }
                            if false {
                                v3 -= v3
                                fmt.Println(v2, v22, v22)
                                v52 = (("hello" + (((("hello" + ((v52 + "hello") + "foo bar")) + v52) + ("a" + (v52 + "hello"))) + v52)) + v52)
                            } else {
                                if v2 {
                                    v80 := GenS5{F0: 0}
                                    _ = v80
                                } else {
                                    var v81 bool = (1 == 83)
                                    _ = v81
                                }
                            }
                        } else {
                            v52 = v52
                        }
                        _ = fi56
                    }
                }
                v82 := v5[2]
                _ = fi53
                _ = v54
                _ = v82
            }
            var v83 uint16 = v1
            _ = v83
        } else {
            var v84 uint16 = v1
            fmt.Println(v84, v84)
            for fi85 := 0; fi85 < 20; fi85++ {
                if len(v51) >= 2 {
                    v51 = v51[1:len(v51)]
                }
                fmt.Println(v2, v4, fi85)
                var v86 float64 = float64(((1e10 * 1e10) / 3.14159))
                _ = fi85
                _ = v86
            }
            _ = v84
        }
        var v87 int = -1
        _ = v87
    } else {
        var v88 int64 = 973
        fmt.Println(v3, v3)
        v5[uint32(v3)] = 0
        _ = v88
    }
    v1 -= v4
    for fi89 := 0; fi89 < 2; fi89++ {
        for fi90 := 0; fi90 < 3; fi90++ {
            if ((genFn0(42, (v22 % (((^(v22))) | 1)), (v26 / (((v26 * 234)) | 1))) < 0) && v2) {
                if v2 {
                    v3 = (^(v3))
                } else {
                    v1 = v22
                    v91 := GenS2{F0: genFn0(946, (v22 | 65535), v26), F1: v2, F2: true}
                    _ = v91
                }
                if len(v51) > 0 {
                    v51[0] = -1
                }
            } else {
                var v92 uint32 = uint32(v26)
                var v93 int8 = 100
                fmt.Println(v3)
                v94 := GenS4{F0: v26}
                _ = v92
                _ = v93
                _ = v94
            }
            v51 = append(v51, fi90)
            v2 = v2
            var v95 int = (fi89 ^ 32768)
            _ = fi90
            _ = v95
        }
        for fi96 := 0; fi96 < 18; fi96++ {
            switch (-(fi89)) {
            case 32767:
                v5[2147483648] = 1
            case 42:
                v97 := GenS1{F0: 1, F1: 1, F2: ((202 >= 1) && v2)}
                _ = v97
            case 4294967295:
                if v2 {
                    fmt.Println(fi89)
                } else {
                    if false {
                        v26 = v26
                        var v98 string = ((("" + v52) + "") + (v52 + v52))
                        _ = v98
                    } else {
                        var v99 uint64 = v26
                        v4 = 501
                        _ = v99
                    }
                }
            default:
                v100 := GenS6{F0: uint16(v26), F1: (v22 % ((v4) | 1)), F2: (^(v26))}
                _ = v100
            }
            var v101 float64 = float64(-2.25)
            switch (v52 + (v52 + "\n\t")) {
            case "foo bar":
                var v102 uint32 = 65535
                fmt.Println(v3, v52, v2)
                v52 = ((("\n\t" + (v52 + (("a" + "\n\t") + (("" + v52) + "foo bar")))) + v52) + v52)
                _ = v102
            }
            if (!v2) {
                if v2 {
                    var v103 uint16 = v22
                    fmt.Println(v101, v4)
                    _ = v103
                } else {
                    v104 := GenS3{F0: (^(v26))}
                    _ = v104
                }
            }
            if v2 {
                var v105 float32 = ((float32(0.0) + float32(-2.25)) / 1.5)
                var v106 float64 = (v101 / v101)
                _ = v105
                _ = v106
            }
            v4 = v4
            v1 = 931
            v2 = ((!false) || (v101 <= v101))
            _ = fi96
            _ = v101
        }
        for fi107 := 0; fi107 < 19; fi107++ {
            if (0 < 1) {
                var v108 int64 = -9223372036854775808
                v22 = (v22 + v4)
                _ = v108
            }
            if ("\n\t" > "\n\t") {
                if ((!(-128 < -128)) || (!v2)) {
                    var v109 uint = 42
                    _ = v109
                } else {
                    switch (fi89 & fi89) {
                    case -1:
                        v3 = v3
                    case 127:
                        switch fi89 {
                        case 42:
                            if (!v2) {
                                var v110 bool = (32767 > 42)
                                if (1 == 2) {
                                    v111 := GenS3{F0: v26}
                                    _ = len(v6)
                                    var v112 int64 = -1
                                    _ = v111
                                    _ = v112
                                } else {
                                    if true {
                                        v113 := GenS7{F0: 42, F1: v52, F2: 2}
                                        v51 = append(v51, fi107)
                                        switch 9223372036854775807 {
                                        case -1:
                                            if len(v51) > 0 {
                                                v51[0] = (^(fi89))
                                            }
                                            switch ("hello" + v52) {
                                            case "hello":
                                                v2 = (v4 >= v22)
                                            }
                                        case 2147483647:
                                            var v114 uint32 = 65535
                                            var v115 bool = (259 < 1)
                                            if len(v51) >= 2 {
                                                v51 = v51[1:len(v51)]
                                            }
                                            _ = v114
                                            _ = v115
                                        default:
                                            var v116 int16 = 42
                                            _ = v116
                                        }
                                        v117 := GenS4{F0: 0}
                                        fmt.Println(v52, v4, v1)
                                        _ = v113
                                        _ = v117
                                    }
                                    v110 = (v26 != (v26 >> (((v3 / ((uint(v4)) | 1))) & 63)))
                                }
                                _ = v110
                            } else {
                                var v118 int8 = 0
                                var v119 uint8 = 0
                                _ = v118
                                _ = v119
                            }
                        case 32767:
                            v120 := GenS8{F0: int8(v26), F1: 0, F2: 42}
                            _ = v120
                        case -2147483648:
                            var v121 int64 = 9223372036854775807
                            _ = v121
                        default:
                            var v122 float64 = -2.25
                            _ = v122
                        }
                    }
                    if v2 {
                        var v123 int64 = -1
                        switch ("a" + (v52 + v52)) {
                        case "a":
                            fmt.Println(v2, v22, v4)
                        case "foo bar":
                            if false {
                                var v124 string = (v52 + v52)
                                v123 -= (v123 >> ((v3) & 63))
                                _ = v124
                            } else {
                                v125 := GenS5{F0: -11}
                                _ = v125
                            }
                        }
                        var v126 int32 = -1
                        _ = v123
                        _ = v126
                    }
                    fmt.Println(v26, v1)
                }
            }
            if (!v2) {
                var v127 float32 = 0.1
                _ = v127
            }
            v26 = v26
            _ = fi107
        }
        _ = fi89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v6
}
