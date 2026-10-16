package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint }
type GenS7 struct { F0 int64; F1 bool; F2 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint; F1 int8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint16; F1 int64; F2 int8 }
type GenS12 struct { F0 int64; F1 int16; F2 uint64 }
type GenS13 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS14 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 uint; F1 int16; F2 uint }
type GenS17 struct { F0 bool; F1 uint8 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 int16; F1 int32; F2 int }

func main() {
    v1 := make(map[int]int)
    var v2 int8 = 42
    var v3 int32 = 0
    var v4 uint = 18446744073709551615
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := GenS1{F0: (v2 * (v2 & v2))}
        v7 := GenS2{F0: v4}
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: (v4 >> ((2) & 63))}
    var v9 int16 = -32768
    fmt.Println(v2)
    var v10 bool = false
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 12; fi12++ {
            var v13 int = (fi12 ^ fi11)
            switch "\n\t" {
            case "a":
                switch "hello" {
                case "hello":
                    fmt.Println(v2)
                case "foo bar":
                    v14 := GenS3{F0: 2, F1: 277}
                    _ = v14
                default:
                    var v15 bool = (true && v10)
                    _ = v15
                }
            case "\n\t":
                switch ("foo bar" + "") {
                case "\n\t":
                    if ((0.1 * -2.25) >= float32(1e10)) {
                        v4 = v4
                        fmt.Println(v10, v10, v2)
                        var v16 int8 = v2
                        if v10 {
                            switch (^(fi11)) {
                            case -9223372036854775808:
                                v17 := GenS4{F0: -35}
                                _ = v17
                            case 4294967295:
                                switch fi12 {
                                case -306:
                                    v3 += (^(v3))
                                default:
                                    v13 -= (-(fi12))
                                }
                            case 127:
                                if false {
                                    fmt.Println(v3, v10, v3)
                                    if v10 {
                                        fmt.Println(v9)
                                    }
                                } else {
                                    var v18 int = 128
                                    _ = v18
                                }
                            }
                            var v19 uint8 = 127
                            switch fi12 {
                            case 580:
                                if v10 {
                                    v4 += v4
                                    v9 -= -875
                                }
                            case 9223372036854775807:
                                var v20 uint32 = 2147483647
                                _ = v20
                            }
                            _ = v19
                        } else {
                            fmt.Println(v3, v9)
                        }
                        var v21 float32 = -2.25
                        _ = v16
                        _ = v21
                    } else {
                        if (v2 != v2) {
                            v22 := GenS2{F0: 535}
                            var v23 bool = v10
                            v1[32767] = 0
                            _ = v22
                            _ = v23
                        } else {
                            var v24 float64 = float64(float32(1e10))
                            v24 = (0.1 * (3.14159 + (-2.25 - (v24 - v24))))
                            var v25 int = v13
                            _ = v24
                            _ = v25
                        }
                        switch "a" {
                        case "":
                            if (v13 >= 4294967295) {
                                switch fi12 {
                                case -947:
                                    v4 += (v4 % ((305) | 1))
                                    v26 := []int{v13, (^(v13)), int(v4), (v13 ^ (^(fi12)))}
                                    v3 += -96
                                    _ = v26
                                case -2147483648:
                                    var v27 string = "hello"
                                    v28 := GenS2{F0: 535}
                                    fmt.Println(fi11, v10, fi11)
                                    switch v13 {
                                    case -2147483648:
                                        var v29 string = v27
                                        _ = v29
                                    case 128:
                                        switch ("foo bar" + (((("a" + ((("a" + v27) + (v27 + ("\n\t" + v27))) + v27)) + v27) + "\n\t") + "foo bar")) {
                                        case "a":
                                            v9 -= v9
                                            v30 := GenS4{F0: v9}
                                            _ = v30
                                        case "hello":
                                            v31 := []int{-2147483648, (v13 & v13)}
                                            _ = v31
                                        }
                                    case -9223372036854775808:
                                        v32 := []int{fi12, fi12, (fi11 << ((v4) & 63))}
                                        _ = v32
                                    }
                                    _ = v27
                                    _ = v28
                                case 464:
                                    var v33 bool = false
                                    _ = v33
                                default:
                                    var v34 int8 = v2
                                    var v35 uint8 = 42
                                    var v36 float32 = 1e10
                                    if v10 {
                                        var v37 int8 = (v2 >> ((147) & 63))
                                        switch fi12 {
                                        case -398:
                                            v38 := GenS5{F0: 2, F1: v4, F2: 838}
                                            _ = v38
                                        case -2147483648:
                                            var v39 uint16 = 42
                                            _ = v39
                                        case 1:
                                            v40 := []int{1, 32767, v13, (-(v13))}
                                            _ = v40
                                        default:
                                            fmt.Println(v2, v34, fi12)
                                        }
                                        _ = v37
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            if v10 {
                                                if (v10 && (true && (!((v34 - 100) >= (v34 ^ (v2 + v34)))))) {
                                                    if (v10 || (!true)) {
                                                        var v41 uint64 = 1
                                                        switch (v13 + -344) {
                                                        case 4294967295:
                                                            v42 := []int{}
                                                            _ = v42
                                                        }
                                                        _ = v41
                                                    }
                                                    fmt.Println(fi12, v13)
                                                }
                                            } else {
                                                if ((fi12 == -1) && (("hello" + "foo bar") < "")) {
                                                    v34 = (v34 ^ (-(v34)))
                                                    if (!(!(v10 && (("a" + "") >= ("hello" + "a"))))) {
                                                        var v43 int = v13
                                                        _ = v43
                                                    } else {
                                                        v34 = v34
                                                        if true {
                                                            v13 = fi12
                                                        }
                                                        v34 -= v2
                                                    }
                                                    v4 = v4
                                                } else {
                                                    v44 := []int{(fi12 >> (((v4 + 539)) & 63)), v13, (-(v13)), v13}
                                                    _ = v44
                                                }
                                                if ("hello" == "") {
                                                    if (36 <= (^(v35))) {
                                                        v45 := GenS1{F0: -1}
                                                        var v46 bool = (v3 >= v3)
                                                        _ = v45
                                                        _ = v46
                                                    } else {
                                                        switch "" {
                                                        case "foo bar":
                                                            v47 := []int{(^(fi12)), 9223372036854775807, fi12, (fi11 * v13)}
                                                            _ = v47
                                                        case "":
                                                            v35 += v35
                                                        case "hello":
                                                            if false {
                                                                fmt.Println(v13, v2, v35)
                                                                if true {
                                                                    fmt.Println(v3, v34)
                                                                    fmt.Println(v34)
                                                                }
                                                                switch (v13 ^ (fi12 + 128)) {
                                                                case 0:
                                                                    var v48 int8 = v2
                                                                    _ = v48
                                                                default:
                                                                    v36 += v36
                                                                }
                                                            } else {
                                                                if ((^(v9)) > v9) {
                                                                    var v49 uint = v4
                                                                    _ = v49
                                                                } else {
                                                                    var v50 uint = (v4 % ((4294967295) | 1))
                                                                    _ = v50
                                                                }
                                                                if (v3 >= v3) {
                                                                    v34 = (v2 | v2)
                                                                    v51 := GenS6{F0: v35, F1: 964, F2: (^(v4))}
                                                                    v52 := GenS6{F0: 92, F1: uint16(v4), F2: v4}
                                                                    _ = v51
                                                                    _ = v52
                                                                }
                                                            }
                                                        default:
                                                            v53 := []int{32768}
                                                            _ = v53
                                                        }
                                                    }
                                                    v54 := []int{fi11, int(v3), -1}
                                                    v35 -= (v35 * v35)
                                                    _ = v54
                                                } else {
                                                    fmt.Println(v3)
                                                    var v55 uint32 = 177
                                                    fmt.Println(v10)
                                                    _ = v55
                                                }
                                                v56 := GenS5{F0: 65535, F1: (^(v4)), F2: 255}
                                                _ = v56
                                            }
                                        }
                                        if (!v10) {
                                            v57 := []int{9223372036854775807, (fi12 * fi12), (fi11 / (((fi11 >> (((^(v4))) & 63))) | 1))}
                                            v58 := GenS2{F0: (v4 / ((v4) | 1))}
                                            if (!(408 > 621)) {
                                                var v59 int = fi12
                                                v60 := GenS7{F0: 0, F1: (!(-225 > -1)), F2: 65535}
                                                _ = v59
                                                _ = v60
                                            } else {
                                                v13 += fi11
                                                switch fi11 {
                                                case 2147483647:
                                                    _ = len(v57)
                                                }
                                                _ = len(v1)
                                            }
                                            if v10 {
                                                v61 := v1[fi11]
                                                switch -1 {
                                                case 828:
                                                    var v62 uint32 = 2147483648
                                                    var v63 bool = true
                                                    var v64 uint8 = v35
                                                    var v65 int8 = (-(v2))
                                                    if len(v57) > 0 {
                                                        v57[0] = (-(fi11))
                                                    }
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                    _ = v65
                                                }
                                                v66 := GenS8{F0: 0}
                                                v4 += (^(v4))
                                                var v67 uint64 = 975
                                                _ = v61
                                                _ = v66
                                                _ = v67
                                            }
                                            fmt.Println(fi11, v4)
                                            v35 = v35
                                            v68 := GenS9{F0: v4, F1: v2}
                                            fmt.Println(v3, fi12, v9)
                                            _ = v57
                                            _ = v58
                                            _ = v68
                                        } else {
                                            if v10 {
                                                v35 = v35
                                                v69 := []int{v13, v13, (fi12 << ((4294967295) & 63))}
                                                _ = v69
                                            } else {
                                                v10 = (((v3 / ((-351) | 1)) > int32(v2)) || v10)
                                            }
                                        }
                                        v70 := []int{fi11, 2147483647, (-(fi12))}
                                        _ = v70
                                    }
                                    v2 = (^(v2))
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                var v71 uint64 = 552
                                v72 := []int{-1, (v13 + -2147483648), (fi11 & fi11)}
                                _ = v71
                                _ = v72
                            } else {
                                fmt.Println(v10, v2)
                                var v73 uint32 = 807
                                if true {
                                    if ((!("" < "foo bar")) && true) {
                                        v73 = uint32(v4)
                                    } else {
                                        if ("a" < "") {
                                            var v74 bool = (1.5 != (((3.14159 / 1.0) / (((0.0 + 1e10) + (0.1 + ((0.0 / (1.5 - (1e10 + float32((((float64(-2.25) - 3.14159) / 1.5) + float64((float32(float64(0.1)) * 0.0))))))) + (1e10 / 0.1)))) / float32((-2.25 - (3.14159 - 1.5))))) - 1.5))
                                            _ = v74
                                        }
                                        v75 := GenS10{F0: 0}
                                        var v76 int64 = -1
                                        _ = v75
                                        _ = v76
                                    }
                                    if false {
                                        if v10 {
                                            v77 := []int{fi11, -722, 1, (fi11 & -293)}
                                            if ((-(v3)) >= 635) {
                                                switch -56 {
                                                case -1:
                                                    v3 -= (v3 & (v3 * (v3 | int32(v13))))
                                                case 4294967295:
                                                    var v78 int = -2147483648
                                                    _ = v78
                                                case -9223372036854775808:
                                                    var v79 float32 = 0.0
                                                    _ = v79
                                                }
                                                switch "" {
                                                case "a":
                                                    v80 := GenS8{F0: 0}
                                                    _ = v80
                                                default:
                                                    v13 = (fi11 / ((-1) | 1))
                                                }
                                            }
                                            if true {
                                                var v81 int64 = -1
                                                v73 = 782
                                                v10 = (!(!v10))
                                                var v82 int16 = 975
                                                v83 := GenS11{F0: 783, F1: -827, F2: (v2 % (((-(v2))) | 1))}
                                                _ = v81
                                                _ = v82
                                                _ = v83
                                            }
                                            var v84 uint8 = 93
                                            _ = v77
                                            _ = v84
                                        } else {
                                            switch "" {
                                            case "a":
                                                v85 := GenS12{F0: 9223372036854775807, F1: (v9 << ((v4) & 63)), F2: 4294967295}
                                                _ = v85
                                            default:
                                                v86 := GenS8{F0: 1}
                                                _ = v86
                                            }
                                            _ = len(v1)
                                            v87 := GenS13{F0: 962, F1: 65535, F2: 974}
                                            _ = v87
                                        }
                                        fmt.Println(v9, fi11)
                                    } else {
                                        var v88 uint16 = 255
                                        var v89 uint8 = 42
                                        switch (fi11 ^ 0) {
                                        case 4294967295:
                                            if ((((0.1 - 0.0) * (1.5 - 0.1)) * -2.25) > -2.25) {
                                                if false {
                                                    v1[v13] = -1
                                                }
                                                v90 := []int{fi11, -551}
                                                if false {
                                                    v88 -= v88
                                                    if v10 {
                                                        v91 := v1[(fi12 * (fi11 << (((v4 << ((795) & 63))) & 63)))]
                                                        _ = v91
                                                    } else {
                                                        v10 = ((true && v10) && ((-2.25 == float32((3.14159 + -2.25))) && false))
                                                    }
                                                    v89 += 128
                                                }
                                                var v92 uint64 = 355
                                                v10 = (v10 || v10)
                                                var v93 uint32 = (v73 % ((v73) | 1))
                                                if len(v90) >= 2 {
                                                    v90 = v90[1:len(v90)]
                                                }
                                                _ = v90
                                                _ = v92
                                                _ = v93
                                            } else {
                                                var v94 int32 = (^(v3))
                                                var v95 uint64 = 4294967295
                                                var v96 int32 = (v3 + 312)
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                            }
                                            v3 += (v3 >> ((v4) & 63))
                                            if (v88 == 2) {
                                                v97 := []int{v13, 32767, fi11}
                                                if true {
                                                    v89 += 128
                                                    if ((-(fi11)) >= -9223372036854775808) {
                                                        if ("\n\t" > "\n\t") {
                                                            var v98 string = "hello"
                                                            switch fi12 {
                                                            case 387:
                                                                if len(v97) > 0 {
                                                                    v97[0] = fi12
                                                                }
                                                            default:
                                                                switch (v98 + v98) {
                                                                case "a":
                                                                    delete(v1, (v13 - fi11))
                                                                case "hello":
                                                                    v2 = (v2 ^ (v2 ^ v2))
                                                                }
                                                            }
                                                            if v10 {
                                                                switch 32768 {
                                                                case 42:
                                                                    var v99 string = "hello"
                                                                    _ = v99
                                                                case 32768:
                                                                    v100 := GenS13{F0: -1, F1: v73, F2: v88}
                                                                    _ = v100
                                                                default:
                                                                    switch v13 {
                                                                    case 609:
                                                                        var v101 int32 = v3
                                                                        _ = v101
                                                                    case -2147483648:
                                                                        if len(v97) >= 2 {
                                                                            v97 = v97[1:len(v97)]
                                                                        }
                                                                    default:
                                                                        v102 := GenS14{F0: (v2 - v2), F1: v89, F2: (^(v89))}
                                                                        _ = v102
                                                                    }
                                                                }
                                                                var v103 float64 = 1.5
                                                                var v104 int8 = (^(v2))
                                                                fmt.Println(fi11, v103)
                                                                delete(v1, (v13 - int(v104)))
                                                                _ = v103
                                                                _ = v104
                                                            } else {
                                                                v105 := GenS3{F0: 1, F1: (v88 - (^(v88)))}
                                                                switch v98 {
                                                                case "a":
                                                                    if (v9 < (v9 - -631)) {
                                                                        var v106 int8 = (-(v2))
                                                                        _ = v106
                                                                    } else {
                                                                        v107 := GenS4{F0: (-(v9))}
                                                                        switch (fi11 + (v13 % ((int(v2)) | 1))) {
                                                                        case 700:
                                                                            fmt.Println(v10)
                                                                        default:
                                                                            switch ("a" + (v98 + v98)) {
                                                                            case "\n\t":
                                                                                if (v89 < v89) {
                                                                                    v108 := GenS8{F0: 0}
                                                                                    if (v10 && v10) {
                                                                                        v73 -= 2147483648
                                                                                        v109 := GenS15{F0: fi12}
                                                                                        v10 = ((v10 || v10) || v10)
                                                                                        fmt.Println(fi11, v89, v4)
                                                                                        if (v98 > ("a" + "")) {
                                                                                            if (!(!((^(v89)) != v89))) {
                                                                                                v88 -= (v88 / ((851) | 1))
                                                                                            }
                                                                                        }
                                                                                        var v110 float64 = -2.25
                                                                                        _ = v109
                                                                                        _ = v110
                                                                                    } else {
                                                                                        v111 := GenS16{F0: (v4 * (v4 >> ((762) & 63))), F1: (v9 + 32767), F2: v4}
                                                                                        var v112 uint8 = (v89 & (v89 / ((v89) | 1)))
                                                                                        fmt.Println(v112, v112)
                                                                                        _ = v111
                                                                                        _ = v112
                                                                                    }
                                                                                    fmt.Println(v3, v13, fi11)
                                                                                    if v10 {
                                                                                        v97 = append(v97, 127)
                                                                                        var v113 string = (v98 + ("foo bar" + "hello"))
                                                                                        if (v3 < v3) {
                                                                                            v98 = v98
                                                                                        } else {
                                                                                            v97 = append(v97, v13)
                                                                                            if v10 {
                                                                                                v88 = v88
                                                                                                var v114 uint32 = (v73 / ((2147483647) | 1))
                                                                                                _ = v114
                                                                                            }
                                                                                        }
                                                                                        _ = v113
                                                                                    } else {
                                                                                        switch fi12 {
                                                                                        case 761:
                                                                                            if true {
                                                                                                var v115 float32 = float32((-2.25 / (float64((3.14159 - 1.5)) / (0.1 * (1.5 * 0.1)))))
                                                                                                var v116 float32 = ((v115 * 3.14159) - float32(3.14159))
                                                                                                var v117 int16 = (v9 - (v9 << ((v4) & 63)))
                                                                                                _ = v115
                                                                                                _ = v116
                                                                                                _ = v117
                                                                                            } else {
                                                                                                fmt.Println(v9)
                                                                                                if (!v10) {
                                                                                                    fmt.Println(v13, v73, v2)
                                                                                                }
                                                                                                var v118 uint8 = v89
                                                                                                if true {
                                                                                                    v4 = v4
                                                                                                }
                                                                                                var v119 string = v98
                                                                                                var v120 float32 = 1e10
                                                                                                var v121 string = v98
                                                                                                var v122 int8 = int8(v3)
                                                                                                v118 -= v118
                                                                                                _ = v118
                                                                                                _ = v119
                                                                                                _ = v120
                                                                                                _ = v121
                                                                                                _ = v122
                                                                                            }
                                                                                        default:
                                                                                            v1[int(v89)] = fi11
                                                                                            var v123 bool = (!(v98 < "a"))
                                                                                            _ = v123
                                                                                        }
                                                                                        v97 = append(v97, (fi11 << ((1) & 63)))
                                                                                    }
                                                                                    _ = v108
                                                                                } else {
                                                                                    v73 = v73
                                                                                }
                                                                            case "":
                                                                                var v124 string = v98
                                                                                _ = v124
                                                                            case "foo bar":
                                                                                v98 = (("\n\t" + "") + v98)
                                                                            default:
                                                                                var v125 float64 = 0.1
                                                                                _ = v125
                                                                            }
                                                                        }
                                                                        var v126 int8 = (v2 & v2)
                                                                        if len(v97) > 0 {
                                                                            v97[0] = -2147483648
                                                                        }
                                                                        fmt.Println(v98, fi11, v88)
                                                                        v73 -= 2
                                                                        v127 := GenS15{F0: 32768}
                                                                        _ = v107
                                                                        _ = v126
                                                                        _ = v127
                                                                    }
                                                                case "hello":
                                                                    if false {
                                                                        if (((((v98 + v98) + v98) + "hello") < v98) && v10) {
                                                                            v128 := GenS9{F0: (^(v4)), F1: v2}
                                                                            _ = len(v1)
                                                                            v73 += 385
                                                                            _ = v128
                                                                        }
                                                                        _ = len(v1)
                                                                    } else {
                                                                        v73 = (v73 * v73)
                                                                    }
                                                                }
                                                                _ = v105
                                                            }
                                                            v9 -= v9
                                                            _ = v98
                                                        }
                                                        v129 := GenS7{F0: 9223372036854775807, F1: (!v10), F2: 0}
                                                        v73 -= v73
                                                        _ = v129
                                                    } else {
                                                        var v130 float64 = float64(0.1)
                                                        _ = v130
                                                    }
                                                    var v131 int = fi11
                                                    _ = v131
                                                }
                                                delete(v1, v13)
                                                _ = v97
                                            }
                                        case 32768:
                                            var v132 int8 = (v2 ^ -1)
                                            _ = v132
                                        default:
                                            switch (fi12 / ((fi11) | 1)) {
                                            case -9223372036854775808:
                                                fmt.Println(v13)
                                            }
                                        }
                                        if true {
                                            var v133 uint8 = v89
                                            _ = v133
                                        } else {
                                            v88 = (v88 >> (((v4 / ((976) | 1))) & 63))
                                            var v134 uint = uint(v88)
                                            v88 += v88
                                            _ = v134
                                        }
                                        if ("foo bar" <= "") {
                                            v135 := GenS17{F0: false, F1: v89}
                                            _ = v135
                                        } else {
                                            v73 -= (v73 * v73)
                                        }
                                        v136 := GenS6{F0: 128, F1: v88, F2: v4}
                                        _ = v136
                                        _ = v88
                                        _ = v89
                                    }
                                    switch (^(fi11)) {
                                    case 128:
                                        var v137 int = (fi12 / ((v13) | 1))
                                        _ = v137
                                    case 807:
                                        switch (fi12 / ((v13) | 1)) {
                                        case -242:
                                            var v138 float64 = ((0.1 * 1.5) * 0.1)
                                            _ = v138
                                        default:
                                            var v139 int = (fi11 << ((v4) & 63))
                                            var v140 int8 = -77
                                            _ = v139
                                            _ = v140
                                        }
                                    }
                                }
                                _ = v73
                            }
                        case "\n\t":
                            var v141 int16 = (v9 + (v9 * -32768))
                            v10 = (1 < 2147483648)
                            var v142 int32 = (v3 & (v3 << ((v4) & 63)))
                            _ = v141
                            _ = v142
                        case "foo bar":
                            v13 = v13
                        default:
                            fmt.Println(v4, v3, v4)
                        }
                        if true {
                            if (220 > 1) {
                                v143 := v1[v13]
                                _ = v143
                            } else {
                                v3 -= v3
                                switch "hello" {
                                case "\n\t":
                                    var v144 int8 = (-(v2))
                                    _ = v144
                                default:
                                    v3 += (v3 ^ (v3 ^ v3))
                                }
                                var v145 uint16 = 2
                                _ = v145
                            }
                        } else {
                            var v146 int16 = (v9 + v9)
                            var v147 int64 = 1
                            if v10 {
                                var v148 uint8 = 127
                                _ = v148
                            } else {
                                v1[fi11] = (fi11 * int(v4))
                            }
                            v4 = v4
                            var v149 int32 = 0
                            _ = v146
                            _ = v147
                            _ = v149
                        }
                    }
                case "a":
                    v150 := GenS5{F0: 2, F1: (v4 & (v4 | v4)), F2: 851}
                    _ = v150
                default:
                    fmt.Println(v9, v13)
                }
            }
            v151 := GenS7{F0: -741, F1: (!v10), F2: 2147483647}
            _ = len(v1)
            _ = fi12
            _ = v13
            _ = v151
        }
        fmt.Println(v2)
        var v152 uint32 = 65535
        _ = fi11
        _ = v152
    }
    for fi153 := 0; fi153 < 19; fi153++ {
        var v154 int16 = v9
        v9 -= (v154 % ((v9) | 1))
        _ = fi153
        _ = v154
    }
    v9 = v9
    for fi155 := 0; fi155 < 10; fi155++ {
        v4 = v4
        _ = fi155
    }
    var v156 int64 = 869
    switch (("\n\t" + "a") + "hello") {
    case "hello":
        for fi157 := 0; fi157 < 10; fi157++ {
            v158 := GenS7{F0: int64(v3), F1: ((^(v4)) >= 850), F2: 1}
            _ = fi157
            _ = v158
        }
    case "":
        v159 := GenS9{F0: 823, F1: 42}
        _ = v159
    }
    v160 := GenS9{F0: (v4 | (v4 >> ((v4) & 63))), F1: (v2 | (-(v2)))}
    v1[-1] = 0
    if false {
        var v161 uint8 = 255
        v162 := []int{}
        _ = v161
        _ = v162
    } else {
        if true {
            v163 := GenS18{F0: v9}
            for fi164 := 0; fi164 < 20; fi164++ {
                if (false && (32767 == 42)) {
                    var v165 int16 = v9
                    v3 = (^(v3))
                    for fi166 := 0; fi166 < 4; fi166++ {
                        v167 := []int{(fi164 / ((fi166) | 1)), -746, (fi166 - fi164), -286}
                        if len(v167) > 0 {
                            v167[0] = 2147483647
                        }
                        _ = fi166
                        _ = v167
                    }
                    _ = v165
                }
                if true {
                    v9 = (v9 - (v9 & v9))
                } else {
                    fmt.Println(v3)
                    fmt.Println(v10, fi164)
                    for fi168 := 0; fi168 < 7; fi168++ {
                        v156 = (v156 ^ (v156 | v156))
                        var v169 uint32 = 0
                        fmt.Println(v10, v4, v2)
                        _ = fi168
                        _ = v169
                    }
                }
                switch fi164 {
                case -707:
                    var v170 string = ""
                    _ = v170
                case 128:
                    var v171 int64 = (v156 & v156)
                    _ = v171
                case 2147483647:
                    for fi172 := 0; fi172 < 2; fi172++ {
                        if v10 {
                            var v173 uint32 = 160
                            if v10 {
                                v174 := GenS2{F0: (v4 | (v4 * v4))}
                                _ = v174
                            }
                            v175 := []int{(^(fi164)), fi172, (fi164 % (((fi172 << (((v4 + v4)) & 63))) | 1))}
                            fmt.Println(v2, v10)
                            var v176 uint32 = v173
                            _ = v173
                            _ = v175
                            _ = v176
                        } else {
                            if (0.1 <= 3.14159) {
                                var v177 uint8 = 127
                                _ = v177
                            }
                        }
                        _ = fi172
                    }
                }
                _ = fi164
            }
            v9 = 1
            _ = v163
        } else {
            v10 = (v9 >= (v9 * 120))
            for fi178 := 0; fi178 < 2; fi178++ {
                v1[128] = fi178
                _ = fi178
            }
        }
        v1[32768] = -1
    }
    v10 = (32767 == 675)
    v3 = (v3 - (^(v3)))
    if (v10 && (127 < -1)) {
        for fi179 := 0; fi179 < 2; fi179++ {
            v180 := GenS19{F0: 1, F1: v3, F2: (fi179 ^ fi179)}
            var v181 float64 = 3.14159
            var v182 uint64 = 557
            _ = fi179
            _ = v180
            _ = v181
            _ = v182
        }
        v2 = (v2 ^ v2)
    }
    v156 = (v156 / (((v156 << (((v4 / ((v4) | 1))) & 63))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v156:", v156)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v156
    _ = v160
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
