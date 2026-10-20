package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint; F2 int }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint; F1 uint; F2 int8 }
type GenS5 struct { F0 int64; F1 string; F2 int64 }
type GenS6 struct { F0 int16; F1 uint32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int; F1 int; F2 int32 }
type GenS9 struct { F0 int32; F1 bool; F2 uint16 }
type GenS10 struct { F0 uint8; F1 uint }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int64; F1 uint32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 string; F1 bool }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int16 }
type GenS17 struct { F0 uint8; F1 uint; F2 int64 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v1 = v1
    if false {
        v1 = v1
    }
    if true {
        fmt.Println(v1)
        v4 := []int{}
        _ = v4
    } else {
        var v5 uint = 0
        _ = v5
    }
    if false {
        for fi6 := 0; fi6 < 10; fi6++ {
            v7 := GenS1{F0: 0, F1: 128, F2: 1}
            _ = fi6
            _ = v7
        }
        v8 := []int{4294967295, 246, 128}
        var v9 uint = 4294967295
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            if (65535 <= 255) {
                if false {
                    var v11 float32 = (3.14159 + ((1e10 / 1.5) + (3.14159 / (float32((-2.25 - 0.0)) - 1.5))))
                    for fi12 := 0; fi12 < 20; fi12++ {
                        v11 = (-2.25 * v11)
                        switch 128 {
                        case 128:
                            v11 += (v11 / (-2.25 / v11))
                        case 92:
                            var v13 uint8 = 2
                            _ = v13
                        default:
                            v14 := GenS1{F0: 42, F1: 2, F2: 42}
                            _ = v14
                        }
                        switch (fi10 * (fi10 | (fi10 | 32767))) {
                        case 128:
                            var v15 uint32 = 2147483647
                            _ = v15
                        case 1:
                            if (false && true) {
                                if true {
                                    switch fi10 {
                                    case -440:
                                        switch (v1 + "") {
                                        case "hello":
                                            if true {
                                                v1 = v1
                                                v2[(fi12 - fi12)] = (0.1 <= float64(v11))
                                            } else {
                                                var v16 uint64 = 42
                                                var v17 uint = 1
                                                v18 := GenS2{F0: v16, F1: v17, F2: fi12}
                                                _ = v16
                                                _ = v17
                                                _ = v18
                                            }
                                        }
                                    case 127:
                                        if (!true) {
                                            fmt.Println(fi10, fi12)
                                        } else {
                                            if false {
                                                fmt.Println(fi12, v1)
                                                v19 := []int{127, (-(fi12)), fi12, (fi12 >> ((4294967295) & 63))}
                                                if len(v19) >= 2 {
                                                    v19 = v19[1:len(v19)]
                                                }
                                                _ = v19
                                            } else {
                                                fmt.Println(v1)
                                                if true {
                                                    var v20 int = 32767
                                                    _ = v20
                                                }
                                                v1 = (v1 + "a")
                                            }
                                        }
                                    }
                                    if false {
                                        switch fi10 {
                                        case 1:
                                            v21 := []int{-704, fi12, fi12, -370}
                                            _ = v21
                                        case 0:
                                            v22 := []int{fi12, 127}
                                            _ = v22
                                        case 4294967295:
                                            var v23 bool = true
                                            _ = v23
                                        }
                                        v24 := GenS1{F0: 2, F1: 133, F2: -1}
                                        _ = len(v3)
                                        _ = v24
                                    }
                                    switch v1 {
                                    case "":
                                        if (4294967295 == 18446744073709551615) {
                                            fmt.Println(fi10)
                                            var v25 uint = 18446744073709551615
                                            _ = v25
                                        } else {
                                            switch 1 {
                                            case 42:
                                                v3[983] = 4294967295
                                            case 32768:
                                                switch 782 {
                                                case 0:
                                                    if (!(-2.25 > (v11 * v11))) {
                                                        if (!((!true) || true)) {
                                                            var v26 uint32 = 42
                                                            v1 = v1
                                                            var v27 float64 = 3.14159
                                                            var v28 string = v1
                                                            _ = v26
                                                            _ = v27
                                                            _ = v28
                                                        } else {
                                                            var v29 uint16 = 531
                                                            var v30 float64 = float64(v11)
                                                            var v31 int16 = 42
                                                            _ = v29
                                                            _ = v30
                                                            _ = v31
                                                        }
                                                        switch v1 {
                                                        case "foo bar":
                                                            switch (("foo bar" + ("hello" + v1)) + v1) {
                                                            case "hello":
                                                                switch (fi12 + 2147483647) {
                                                                case 32767:
                                                                    v32 := []int{}
                                                                    _ = v32
                                                                }
                                                            case "":
                                                                var v33 float64 = 1e10
                                                                _ = v33
                                                            case "\n\t":
                                                                v11 = ((v11 - 0.0) + v11)
                                                            }
                                                        }
                                                        switch v1 {
                                                        case "\n\t":
                                                            v1 = (v1 + v1)
                                                        case "foo bar":
                                                            v34 := []int{}
                                                            _ = v34
                                                        case "hello":
                                                            if false {
                                                                if false {
                                                                    var v35 int8 = 0
                                                                    _ = v35
                                                                } else {
                                                                    v1 = ((v1 + v1) + (v1 + v1))
                                                                    switch fi10 {
                                                                    case 32768:
                                                                        v1 = v1
                                                                    }
                                                                }
                                                                var v36 bool = true
                                                                _ = v36
                                                            } else {
                                                                v37 := v2[-562]
                                                                var v38 uint32 = 4294967295
                                                                if true {
                                                                    v38 -= v38
                                                                    v39 := []int{fi10}
                                                                    v1 = "a"
                                                                    _ = v39
                                                                } else {
                                                                    v1 = "a"
                                                                    if ((1.5 + (-2.25 - (1.5 - 0.0))) != -2.25) {
                                                                        var v40 int16 = 1
                                                                        var v41 uint64 = 1
                                                                        var v42 uint = 4294967295
                                                                        v43 := GenS2{F0: v41, F1: v42, F2: fi10}
                                                                        switch (fi10 / (((^(fi12))) | 1)) {
                                                                        case -1:
                                                                            v44 := GenS2{F0: v41, F1: 42, F2: (-(fi12))}
                                                                            var v45 uint64 = (v41 | 18446744073709551615)
                                                                            _ = v44
                                                                            _ = v45
                                                                        case 4294967295:
                                                                            v46 := GenS1{F0: 127, F1: 46, F2: 1}
                                                                            _ = v46
                                                                        }
                                                                        v47 := GenS3{F0: 1, F1: -119}
                                                                        v11 -= (v11 / 0.1)
                                                                        _ = v40
                                                                        _ = v41
                                                                        _ = v42
                                                                        _ = v43
                                                                        _ = v47
                                                                    } else {
                                                                        if true {
                                                                            var v48 int32 = 42
                                                                            var v49 bool = ("a" == "\n\t")
                                                                            _ = v48
                                                                            _ = v49
                                                                        } else {
                                                                            if (((42 != 2) || ((!((false || v37) && false)) && false)) && (v11 < float32(float64(float32(float64((0.0 + ((1.5 * (float32((float64(v11) - (0.0 / (-2.25 * -2.25)))) - v11)) - (float32((((1e10 - float64(((float32((0.1 * 0.1)) - (-2.25 + (v11 + float32(0.0)))) / 1e10))) - 3.14159) - 0.0)) + (v11 * v11)))))))))) {
                                                                                v11 = v11
                                                                            }
                                                                            fmt.Println(fi10)
                                                                            if false {
                                                                                delete(v3, (fi12 / (((fi12 % ((fi12) | 1))) | 1)))
                                                                                if false {
                                                                                    v50 := []int{fi12, fi12, 42}
                                                                                    v51 := v3[fi10]
                                                                                    _ = v50
                                                                                    _ = v51
                                                                                }
                                                                                var v52 int64 = -461
                                                                                _ = v52
                                                                            } else {
                                                                                if v37 {
                                                                                    var v53 int = fi12
                                                                                    v38 -= (v38 | v38)
                                                                                    _ = v53
                                                                                } else {
                                                                                    v54 := []int{fi10}
                                                                                    var v55 uint8 = 188
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                }
                                                                                if (!true) {
                                                                                    v56 := GenS2{F0: 4294967295, F1: 911, F2: (fi10 + fi10)}
                                                                                    if (128 > 41) {
                                                                                        v37 = (818 != 1)
                                                                                    } else {
                                                                                        var v57 int64 = 9223372036854775807
                                                                                        v58 := GenS2{F0: 523, F1: 1, F2: (fi12 + (-(fi10)))}
                                                                                        v59 := []int{(fi12 * (fi10 % ((0) | 1))), (fi12 | fi12)}
                                                                                        _ = v57
                                                                                        _ = v58
                                                                                        _ = v59
                                                                                    }
                                                                                    _ = v56
                                                                                }
                                                                            }
                                                                        }
                                                                        fmt.Println(v11)
                                                                    }
                                                                    fmt.Println(v11, v1, v1)
                                                                }
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                        default:
                                                            v60 := []int{(fi12 - (fi12 | (fi10 + 4294967295))), (fi12 - fi10), (fi10 << ((65535) & 63))}
                                                            _ = v60
                                                        }
                                                    } else {
                                                        v61 := GenS3{F0: 157, F1: -47}
                                                        v1 = v1
                                                        switch v1 {
                                                        case "":
                                                            v1 = (v1 + ((v1 + v1) + ("" + v1)))
                                                        case "\n\t":
                                                            var v62 int8 = 57
                                                            _ = v62
                                                        case "a":
                                                            if (0 > 4294967295) {
                                                                v1 = v1
                                                                v11 = (0.1 + (v11 - float32(1.5)))
                                                            }
                                                        default:
                                                            switch fi12 {
                                                            case 714:
                                                                if false {
                                                                    v1 = (("hello" + "a") + "a")
                                                                    switch fi12 {
                                                                    case 32768:
                                                                        switch fi12 {
                                                                        case 32767:
                                                                            var v63 uint32 = 805
                                                                            _ = v63
                                                                        case -9223372036854775808:
                                                                            v64 := []int{fi12, (fi12 * 32768), (fi12 >> ((326) & 63))}
                                                                            _ = v64
                                                                        case 0:
                                                                            v65 := []int{(fi10 << ((258) & 63)), 1, fi12}
                                                                            v66 := GenS4{F0: 42, F1: 495, F2: 127}
                                                                            _ = v65
                                                                            _ = v66
                                                                        default:
                                                                            _ = len(v2)
                                                                        }
                                                                    default:
                                                                        switch (fi12 - fi12) {
                                                                        case 889:
                                                                            var v67 float32 = (v11 * v11)
                                                                            _ = v67
                                                                        }
                                                                    }
                                                                } else {
                                                                    var v68 float64 = 1e10
                                                                    fmt.Println(fi10, v11, fi12)
                                                                    var v69 int64 = 9223372036854775807
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                            case 32768:
                                                                v70 := []int{fi10, fi12}
                                                                _ = v70
                                                            }
                                                        }
                                                        _ = v61
                                                    }
                                                default:
                                                    v71 := []int{128, (fi12 | fi12), (fi10 + fi12)}
                                                    _ = v71
                                                }
                                            case -2147483648:
                                                fmt.Println(fi12, v1)
                                            }
                                            var v72 int16 = -287
                                            _ = v72
                                        }
                                    case "\n\t":
                                        var v73 bool = false
                                        var v74 int8 = 127
                                        if v73 {
                                            if ((1.5 - 0.1) != 1e10) {
                                                if true {
                                                    if v73 {
                                                        v75 := v2[9223372036854775807]
                                                        var v76 int64 = -9223372036854775808
                                                        switch fi10 {
                                                        case 608:
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v77 string = v1
                                                                _ = v77
                                                            default:
                                                                var v78 uint16 = 839
                                                                _ = v78
                                                            }
                                                            if (0 == 2) {
                                                                switch (fi12 >> ((746) & 63)) {
                                                                case 0:
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v79 uint64 = 995
                                                                        _ = v79
                                                                    case "a":
                                                                        fmt.Println(v11, v74, v76)
                                                                    default:
                                                                        fmt.Println(v1, v75)
                                                                    }
                                                                }
                                                                v3[(-(fi12))] = 2
                                                                if (1 <= 18446744073709551615) {
                                                                    v80 := []int{(fi10 & fi12), (fi12 % ((fi10) | 1)), fi10}
                                                                    var v81 int32 = -2147483648
                                                                    _ = len(v2)
                                                                    _ = v80
                                                                    _ = v81
                                                                } else {
                                                                    var v82 uint = 1
                                                                    _ = v82
                                                                }
                                                                v83 := v2[fi12]
                                                                fmt.Println(v83, v74, v1)
                                                                _ = v83
                                                            }
                                                            var v84 uint = 2
                                                            _ = v84
                                                        }
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                    var v85 int8 = -128
                                                    _ = v85
                                                } else {
                                                    v86 := GenS5{F0: 42, F1: v1, F2: 42}
                                                    if ((float32(0.0) != 0.0) && v73) {
                                                        var v87 int = (fi12 << ((255) & 63))
                                                        v88 := []int{(-(fi10)), fi12, (^(fi10))}
                                                        v73 = v73
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        v73 = v73
                                                        var v89 int16 = 0
                                                        if ((0.1 - -2.25) < 1e10) {
                                                            var v90 uint64 = 510
                                                            v1 = (v1 + v1)
                                                            _ = v90
                                                        } else {
                                                            fmt.Println(fi10, fi12)
                                                        }
                                                        switch 42 {
                                                        case -938:
                                                            var v91 float32 = 0.0
                                                            _ = v91
                                                        case 128:
                                                            var v92 int32 = -1
                                                            v2[fi10] = false
                                                            _ = v92
                                                        }
                                                        _ = v89
                                                    }
                                                    _ = v86
                                                }
                                                if v73 {
                                                    v1 = ("foo bar" + v1)
                                                    v93 := []int{fi10, (fi10 >> ((124) & 63)), (fi10 * fi10)}
                                                    _ = v93
                                                }
                                                switch (v1 + v1) {
                                                case "a":
                                                    v74 -= (^(v74))
                                                case "foo bar":
                                                    var v94 int32 = 42
                                                    _ = v94
                                                case "":
                                                    v74 += v74
                                                }
                                                fmt.Println(v74, fi10)
                                            } else {
                                                fmt.Println(v74)
                                            }
                                            v95 := []int{fi12, 32767}
                                            v96 := GenS6{F0: -431, F1: 872}
                                            _ = v95
                                            _ = v96
                                        }
                                        v1 = v1
                                        _ = v73
                                        _ = v74
                                    }
                                }
                                var v97 uint64 = 559
                                var v98 float64 = (0.0 + ((-2.25 - 1.5) / 3.14159))
                                switch v1 {
                                case "hello":
                                    v99 := []int{}
                                    _ = v99
                                }
                                v100 := GenS4{F0: 42, F1: uint(v97), F2: -1}
                                if false {
                                    switch 0 {
                                    case 128:
                                        v1 = "a"
                                    default:
                                        v101 := GenS7{F0: 258}
                                        _ = v101
                                    }
                                    v102 := GenS8{F0: 128, F1: (fi10 | fi12), F2: -823}
                                    if false {
                                        if (42 > 127) {
                                            v103 := GenS5{F0: -1, F1: v1, F2: -1}
                                            var v104 uint64 = (v97 ^ v97)
                                            var v105 int = 778
                                            _ = v103
                                            _ = v104
                                            _ = v105
                                        }
                                        var v106 int8 = 100
                                        _ = v106
                                    } else {
                                        if ((((float32(v98) - 1.5) * float32(float64(v11))) * v11) > v11) {
                                            if (fi12 < 32768) {
                                                fmt.Println(fi12, v11)
                                                var v107 float32 = (((v11 - (v11 + v11)) / v11) - (1e10 * (v11 / (((1e10 - v11) + float32(float64((3.14159 - (v11 * float32(((v98 / float64((v11 - float32(((0.1 + 1e10) + (v98 - (0.0 / v98))))))) * 3.14159))))))) * v11))))
                                                _ = v107
                                            }
                                            v108 := []int{(fi12 * (fi10 >> ((333) & 63))), (fi10 & 128)}
                                            _ = v108
                                        } else {
                                            if false {
                                                v109 := GenS2{F0: v97, F1: 2, F2: (fi12 % ((0) | 1))}
                                                _ = v109
                                            }
                                            v1 = v1
                                            if (float64(((float32(v98) + float32(v98)) + v11)) < v98) {
                                                if false {
                                                    if ((v97 + (v97 >> ((42) & 63))) > 669) {
                                                        v110 := GenS5{F0: -9223372036854775808, F1: ("" + v1), F2: 1}
                                                        _ = v110
                                                    } else {
                                                        v1 = v1
                                                    }
                                                    var v111 int8 = 127
                                                    _ = v111
                                                } else {
                                                    v112 := []int{(fi12 >> ((65535) & 63)), (fi12 - fi12), 32768, fi10}
                                                    fmt.Println(v11)
                                                    v98 -= -2.25
                                                    v11 = v11
                                                    if len(v112) >= 2 {
                                                        v112 = v112[1:len(v112)]
                                                    }
                                                    if len(v112) > 0 {
                                                        v112[0] = fi12
                                                    }
                                                    _ = v112
                                                }
                                                v113 := GenS9{F0: 934, F1: false, F2: 0}
                                                _ = v113
                                            }
                                        }
                                        v11 = 3.14159
                                        fmt.Println(v97)
                                    }
                                    _ = v102
                                } else {
                                    switch (fi12 % ((fi12) | 1)) {
                                    case 9223372036854775807:
                                        v11 = v11
                                    case 42:
                                        _ = len(v3)
                                    default:
                                        var v114 float32 = (v11 * ((1.5 / v11) + 1e10))
                                        _ = v114
                                    }
                                    v11 -= (v11 - float32(v98))
                                }
                                _ = v100
                                _ = v97
                                _ = v98
                            }
                        default:
                            var v115 uint8 = 2
                            _ = v115
                        }
                        var v116 uint = 77
                        if true {
                            switch fi10 {
                            case 127:
                                delete(v3, (fi10 | fi10))
                            case 370:
                                if false {
                                    switch "foo bar" {
                                    case "hello":
                                        if (!(0 != 819)) {
                                            fmt.Println(fi12, fi12, fi12)
                                            switch fi12 {
                                            case 128:
                                                var v117 string = (v1 + "")
                                                _ = v117
                                            default:
                                                var v118 int16 = 0
                                                _ = v118
                                            }
                                        } else {
                                            fmt.Println(fi12, v11, v116)
                                        }
                                    case "a":
                                        fmt.Println(fi10, fi12)
                                    case "foo bar":
                                        v119 := GenS10{F0: 2, F1: (v116 / ((v116) | 1))}
                                        _ = v119
                                    default:
                                        if true {
                                            switch "\n\t" {
                                            case "foo bar":
                                                v120 := GenS6{F0: 32767, F1: 42}
                                                _ = v120
                                            case "a":
                                                v116 += v116
                                            case "":
                                                var v121 int16 = -596
                                                _ = v121
                                            }
                                            var v122 int64 = 756
                                            _ = v122
                                        } else {
                                            var v123 int8 = -1
                                            v116 = (v116 % ((v116) | 1))
                                            v116 += v116
                                            v1 = ("" + ("a" + ((v1 + v1) + (v1 + ""))))
                                            var v124 float64 = 3.14159
                                            v123 = 12
                                            v2[fi12] = ((((v1 + "hello") + ((v1 + ((("a" + v1) + (v1 + ("\n\t" + v1))) + "a")) + (v1 + "a"))) + ("\n\t" + "hello")) != v1)
                                            _ = v123
                                            _ = v124
                                        }
                                    }
                                    fmt.Println(fi12, fi10, v11)
                                } else {
                                    v125 := GenS1{F0: 46, F1: 255, F2: 1}
                                    v1 = "hello"
                                    if false {
                                        var v126 uint8 = 243
                                        _ = v126
                                    }
                                    _ = v125
                                }
                            case -71:
                                fmt.Println(v1, v1)
                            default:
                                v116 += v116
                            }
                            fmt.Println(v116, fi12, v1)
                        } else {
                            var v127 int = fi12
                            fmt.Println(v127, v116, v116)
                            _ = v127
                        }
                        _ = fi12
                        _ = v116
                    }
                    _ = v11
                } else {
                    for fi128 := 0; fi128 < 13; fi128++ {
                        v1 = v1
                        fmt.Println(fi128, v1)
                        v129 := GenS11{F0: 255}
                        if (true || (true && true)) {
                            v1 = ("foo bar" + v1)
                            v3[(fi10 % (((fi10 | fi128)) | 1))] = 42
                            v130 := GenS12{F0: -385, F1: 471}
                            _ = v130
                        } else {
                            if (4294967295 == 781) {
                                if (!(0 < 2)) {
                                    var v131 int16 = -603
                                    v131 = 1000
                                    var v132 uint16 = 0
                                    v132 -= (v132 & v132)
                                    _ = v131
                                    _ = v132
                                } else {
                                    v133 := []int{}
                                    _ = v133
                                }
                                v134 := GenS11{F0: 1}
                                v135 := GenS7{F0: -371}
                                _ = v134
                                _ = v135
                            } else {
                                v1 = "hello"
                                v136 := []int{(fi10 % ((127) | 1)), 0, fi10, 9223372036854775807}
                                _ = v136
                            }
                        }
                        _ = fi128
                        _ = v129
                    }
                    var v137 uint64 = 1
                    _ = v137
                }
                for fi138 := 0; fi138 < 5; fi138++ {
                    v139 := GenS2{F0: 18446744073709551615, F1: 145, F2: -9223372036854775808}
                    _ = fi138
                    _ = v139
                }
            } else {
                var v140 uint32 = 65535
                for fi141 := 0; fi141 < 4; fi141++ {
                    var v142 int8 = 1
                    v140 -= 65535
                    _ = fi141
                    _ = v142
                }
                for fi143 := 0; fi143 < 12; fi143++ {
                    if (false || (!false)) {
                        v140 += 42
                    } else {
                        v144 := GenS6{F0: 620, F1: 4294967295}
                        switch "foo bar" {
                        case "a":
                            fmt.Println(fi10)
                        default:
                            v2[2147483647] = (false && (!((int8(v140) == 127) || (!(528 != 636)))))
                        }
                        _ = len(v2)
                        switch v1 {
                        case "hello":
                            v140 = 42
                        case "foo bar":
                            var v145 int = 32768
                            _ = v145
                        default:
                            if false {
                                v146 := v3[(fi10 >> ((1) & 63))]
                                if true {
                                    v147 := GenS2{F0: 769, F1: 255, F2: (fi10 | (fi143 ^ (fi143 % ((fi143) | 1))))}
                                    if false {
                                        v148 := []int{403, (fi10 - fi10), 32768}
                                        v149 := GenS13{F0: 1}
                                        _ = v148
                                        _ = v149
                                    } else {
                                        var v150 int8 = 0
                                        v151 := GenS4{F0: 4294967295, F1: 7, F2: -116}
                                        v140 += (v140 - 42)
                                        switch fi143 {
                                        case 32767:
                                            switch (fi10 / ((fi143) | 1)) {
                                            case 42:
                                                var v152 bool = (42 < -9223372036854775808)
                                                _ = v152
                                            case 9223372036854775807:
                                                v153 := GenS6{F0: 1000, F1: v140}
                                                _ = v153
                                            case 32767:
                                                v150 -= v150
                                            }
                                        }
                                        switch v1 {
                                        case "":
                                            v154 := GenS11{F0: 32767}
                                            _ = v154
                                        case "foo bar":
                                            v155 := GenS14{F0: ((v1 + (v1 + (v1 + v1))) + (v1 + v1)), F1: true}
                                            _ = v155
                                        }
                                        _ = v150
                                        _ = v151
                                    }
                                    _ = v147
                                } else {
                                    switch (fi143 % (((-(fi10))) | 1)) {
                                    case 42:
                                        var v156 uint = 767
                                        _ = v156
                                    default:
                                        var v157 int = fi143
                                        _ = v157
                                    }
                                    v140 = v146
                                }
                                v158 := []int{(-(fi10)), -2147483648}
                                _ = v146
                                _ = v158
                            }
                        }
                        _ = v144
                    }
                    v159 := v2[(fi10 & -909)]
                    _ = fi143
                    _ = v159
                }
                _ = v140
            }
            v160 := GenS15{F0: 536}
            if true {
                v161 := GenS16{F0: 1000}
                v162 := GenS5{F0: -1, F1: v1, F2: -9223372036854775808}
                var v163 float32 = 1e10
                _ = v161
                _ = v162
                _ = v163
            } else {
                delete(v2, fi10)
                fmt.Println(v1, v1)
            }
            _ = fi10
            _ = v160
        }
        if (3.14159 <= (0.1 + (0.0 * float64(0.1)))) {
            if ((((42 < -9223372036854775808) && false) || (((-2.25 / -2.25) / (0.1 * (-2.25 * 1.5))) > 1.5)) || (!false)) {
                if false {
                    v1 = "\n\t"
                    fmt.Println(v1)
                    v164 := []int{32768, 32767, 42, 4}
                    _ = v164
                } else {
                    v3[32767] = 207
                    fmt.Println(v1)
                    v165 := []int{32767, 413, 9223372036854775807}
                    _ = v165
                }
                var v166 uint64 = 750
                _ = v166
            } else {
                v1 = "hello"
                var v167 int16 = -197
                var v168 int16 = (v167 >> ((271) & 63))
                _ = v167
                _ = v168
            }
        }
    }
    v169 := GenS7{F0: -1}
    if true {
        switch v1 {
        case "foo bar":
            fmt.Println(v1)
        case "":
            v1 = ("hello" + v1)
        default:
            v170 := GenS14{F0: v1, F1: false}
            _ = v170
        }
    } else {
        for fi171 := 0; fi171 < 4; fi171++ {
            var v172 uint16 = 69
            v1 = v1
            v173 := GenS17{F0: 1, F1: 391, F2: 665}
            _ = fi171
            _ = v172
            _ = v173
        }
    }
    v2[-2147483648] = ((!(true || (-128 != 100))) || false)
    v174 := []int{-2147483648, 42}
    v1 = ("\n\t" + "hello")
    _ = len(v174)
    if len(v174) >= 2 {
        v174 = v174[1:len(v174)]
    }
    v1 = ((v1 + v1) + v1)
    v1 = "a"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v174:", len(v174), v174)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v169
    _ = v174
    _ = v2
    _ = v3
}
