package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 uint16; F2 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint8; F1 uint; F2 int32 }
type GenS9 struct { F0 uint32; F1 string; F2 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int64; F1 int16 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int; F1 int8; F2 int32 }
type GenS15 struct { F0 int32; F1 int; F2 int8 }
type GenS16 struct { F0 bool; F1 uint64 }
func genFn0(p0 int16, p1 uint8) uint {
    var v4 int16 = (p0 - (-(p0)))
    if false {
        var v5 int8 = -1
        _ = v5
    }
    v6 := []int{2147483647, 774, -952}
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    return 1
}
func genFn1() uint8 {
    if true {
        if ("" == "\n\t") {
            var v7 int16 = -32768
            if (false || true) {
                var v8 string = (((("\n\t" + "") + "a") + ("foo bar" + "")) + "")
                _ = v8
            } else {
                for fi9 := 0; fi9 < 18; fi9++ {
                    if ((v7 & (v7 | v7)) <= (-(v7))) {
                        v10 := GenS1{F0: (722 != 18446744073709551615), F1: "", F2: fi9}
                        var v11 int8 = 87
                        switch "foo bar" {
                        case "foo bar":
                            for fi12 := 0; fi12 < 4; fi12++ {
                                if true {
                                    switch "" {
                                    case "a":
                                        v11 += (v11 + v11)
                                    case "foo bar":
                                        var v13 int = fi9
                                        _ = v13
                                    case "\n\t":
                                        var v14 int16 = v7
                                        _ = v14
                                    default:
                                        var v15 float32 = (1e10 + (-2.25 * (0.0 - 1.5)))
                                        _ = v15
                                    }
                                } else {
                                    fmt.Println(fi9, fi9, fi12)
                                    var v16 uint64 = 42
                                    var v17 uint = 255
                                    _ = v16
                                    _ = v17
                                }
                                _ = fi12
                            }
                        case "":
                            v18 := []int{}
                            _ = v18
                        default:
                            v11 = v11
                        }
                        _ = v10
                        _ = v11
                    } else {
                        fmt.Println(v7)
                        var v19 int64 = 1
                        v20 := GenS2{F0: 255}
                        _ = v19
                        _ = v20
                    }
                    _ = fi9
                }
                switch (("\n\t" + "a") + ("\n\t" + "foo bar")) {
                case "foo bar":
                    var v21 int64 = 9223372036854775807
                    _ = v21
                case "\n\t":
                    v22 := GenS3{F0: "", F1: 32767}
                    _ = v22
                default:
                    fmt.Println(v7)
                }
                if false {
                    switch "foo bar" {
                    case "foo bar":
                        var v23 uint32 = 42
                        _ = v23
                    case "a":
                        var v24 uint32 = 1
                        _ = v24
                    case "":
                        if false {
                            fmt.Println(v7)
                            v7 = v7
                            v25 := GenS4{F0: 1}
                            _ = v25
                        }
                    default:
                        v7 -= (-(v7))
                    }
                    v7 = (v7 >> ((4294967295) & 63))
                    v26 := GenS3{F0: "hello", F1: -9223372036854775808}
                    _ = v26
                }
            }
            var v27 uint32 = 2
            _ = v27
            _ = v7
        } else {
            v28 := []int{9223372036854775807, -9223372036854775808, 733, 2147483647}
            _ = v28
        }
    } else {
        if true {
            v29 := GenS5{F0: -143, F1: 0, F2: (-141 >= 0)}
            var v30 int8 = 30
            _ = v29
            _ = v30
        } else {
            for fi31 := 0; fi31 < 5; fi31++ {
                var v32 float64 = -2.25
                for fi33 := 0; fi33 < 4; fi33++ {
                    v34 := []int{-807, fi33, (fi31 << ((0) & 63))}
                    _ = fi33
                    _ = v34
                }
                v32 = v32
                _ = fi31
                _ = v32
            }
        }
        var v35 float32 = (1e10 / 1.0)
        _ = v35
    }
    v36 := []int{}
    var v37 int = 25
    var v38 uint8 = 1
    if (682 != 1) {
        if false {
            switch -876 {
            case 42:
                var v39 int16 = 879
                _ = v39
            case 686:
                v40 := GenS6{F0: 930}
                _ = v40
            case -498:
                var v41 int32 = 2147483647
                _ = v41
            }
            if len(v36) > 0 {
                v36[0] = v37
            }
            if (!false) {
                var v42 int16 = int16(v37)
                var v43 uint8 = (v38 | 1)
                _ = v42
                _ = v43
            } else {
                if len(v36) > 0 {
                    v36[0] = (v37 ^ -9223372036854775808)
                }
            }
        } else {
            for fi44 := 0; fi44 < 4; fi44++ {
                var v45 int8 = 0
                _ = fi44
                _ = v45
            }
        }
    }
    _ = v36
    _ = v37
    _ = v38
    return v38
}

func main() {
    var v1 string = ((("hello" + "a") + "foo bar") + ("\n\t" + "foo bar"))
    var v2 int64 = 328
    var v3 int32 = -2147483648
    v46 := []int{127, 54}
    switch -81 {
    case 32768:
        v1 = (v1 + v1)
    case 42:
        for fi47 := 0; fi47 < 17; fi47++ {
            var v48 uint64 = 999
            v1 = ("hello" + "\n\t")
            var v49 int16 = 32767
            _ = fi47
            _ = v48
            _ = v49
        }
    }
    var v50 int8 = 0
    if len(v46) >= 2 {
        v46 = v46[1:len(v46)]
    }
    if true {
        v2 -= (v2 % ((80) | 1))
    }
    if false {
        v2 += (v2 ^ 1)
        v51 := GenS7{F0: (v1 + "a")}
        var v52 uint16 = 706
        _ = v51
        _ = v52
    } else {
        var v53 float32 = -2.25
        switch 42 {
        case 2147483647:
            v54 := GenS7{F0: "hello"}
            _ = v54
        case 4294967295:
            v53 = (v53 - 1.5)
        default:
            v2 += 741
        }
        _ = v53
    }
    if len(v46) >= 2 {
        v46 = v46[1:len(v46)]
    }
    if len(v46) > 0 {
        v46[0] = -2147483648
    }
    var v55 float32 = 1e10
    v56 := GenS1{F0: ((false && true) && (1e10 <= 0.1)), F1: v1, F2: 42}
    var v57 int8 = v50
    fmt.Println(v1, v2)
    var v58 float64 = ((-2.25 - ((0.0 - 0.0) / (0.0 - 0.1))) / (((3.14159 - (3.14159 - 3.14159)) - (1e10 / 1.0)) * (0.0 - 3.14159)))
    var v59 uint = 0
    fmt.Println(v2, v59)
    if (false && ("\n\t" <= ((v1 + "") + "hello"))) {
        var v60 int64 = -7
        var v61 int8 = int8(v3)
        var v62 uint64 = 0
        _ = v60
        _ = v61
        _ = v62
    } else {
        if len(v46) >= 2 {
            v46 = v46[1:len(v46)]
        }
        if (v2 >= (v2 << ((255) & 63))) {
            if (("a" + v1) == "foo bar") {
                var v63 float64 = (1e10 * 1.5)
                for fi64 := 0; fi64 < 14; fi64++ {
                    v50 = v50
                    if true {
                        v65 := GenS6{F0: 65535}
                        v46 = append(v46, fi64)
                        for fi66 := 0; fi66 < 14; fi66++ {
                            fmt.Println(v58)
                            _ = fi66
                        }
                        _ = v65
                    } else {
                        var v67 uint16 = 65535
                        _ = v67
                    }
                    _ = fi64
                }
                _ = v63
            }
            v59 = (v59 ^ (v59 ^ 462))
            var v68 string = "a"
            _ = v68
        } else {
            switch 42 {
            case -1:
                switch v1 {
                case "a":
                    for fi69 := 0; fi69 < 19; fi69++ {
                        var v70 uint16 = 32767
                        for fi71 := 0; fi71 < 18; fi71++ {
                            v59 -= (v59 / ((v59) | 1))
                            v72 := GenS5{F0: fi71, F1: v70, F2: ((false && (!(2 < 226))) && (1e10 == (v58 * v58)))}
                            _ = fi71
                            _ = v72
                        }
                        switch ((v1 + (("" + "a") + "hello")) + v1) {
                        case "foo bar":
                            switch v1 {
                            case "foo bar":
                                for fi73 := 0; fi73 < 13; fi73++ {
                                    v3 += (v3 ^ (v3 & 2147483647))
                                    v3 -= -1
                                    switch "" {
                                    case "":
                                        var v74 float32 = v55
                                        _ = v74
                                    }
                                    _ = fi73
                                }
                            case "hello":
                                v75 := GenS8{F0: uint8(v70), F1: (v59 >> (((v59 | 255)) & 63)), F2: (v3 ^ v3)}
                                _ = v75
                            case "":
                                var v76 uint32 = 450
                                _ = v76
                            default:
                                v77 := GenS9{F0: 42, F1: (v1 + "\n\t"), F2: v50}
                                _ = v77
                            }
                        case "a":
                            switch (fi69 - int(v2)) {
                            case 2147483647:
                                fmt.Println(v57, v1, v1)
                            default:
                                for fi78 := 0; fi78 < 9; fi78++ {
                                    v2 -= (v2 * v2)
                                    _ = len(v46)
                                    v46 = append(v46, (fi78 & (-(fi78))))
                                    _ = fi78
                                }
                            }
                        case "hello":
                            for fi79 := 0; fi79 < 1; fi79++ {
                                var v80 float32 = 0.0
                                v50 += 73
                                _ = fi79
                                _ = v80
                            }
                        }
                        _ = fi69
                        _ = v70
                    }
                case "hello":
                    for fi81 := 0; fi81 < 8; fi81++ {
                        if len(v46) >= 2 {
                            v46 = v46[1:len(v46)]
                        }
                        fmt.Println(v2)
                        for fi82 := 0; fi82 < 3; fi82++ {
                            v57 = v57
                            _ = fi82
                        }
                        _ = fi81
                    }
                default:
                    if (true || (!true)) {
                        v1 = v1
                        v2 = v2
                        for fi83 := 0; fi83 < 15; fi83++ {
                            for fi84 := 0; fi84 < 19; fi84++ {
                                if (false && false) {
                                    if true {
                                        if false {
                                            v85 := GenS10{F0: 9223372036854775807}
                                            _ = v85
                                        } else {
                                            var v86 int32 = v3
                                            switch ("" + "foo bar") {
                                            case "hello":
                                                if len(v46) > 0 {
                                                    v46[0] = (fi84 - fi84)
                                                }
                                            case "foo bar":
                                                var v87 int64 = int64(v57)
                                                _ = v87
                                            default:
                                                if false {
                                                    v86 += v86
                                                    v88 := GenS11{F0: int64(v57), F1: 1}
                                                    _ = v88
                                                } else {
                                                    v46 = append(v46, 195)
                                                    var v89 uint32 = 0
                                                    if ((0 == 42) && false) {
                                                        v90 := GenS12{F0: v1}
                                                        _ = v90
                                                    } else {
                                                        v58 = v58
                                                        _ = len(v46)
                                                        switch (fi84 >> ((v59) & 63)) {
                                                        case 2147483647:
                                                            fmt.Println(v50, fi83)
                                                        case 127:
                                                            v46 = append(v46, fi83)
                                                        case -2147483648:
                                                            if false {
                                                                var v91 string = v1
                                                                fmt.Println(v59)
                                                                _ = v91
                                                            } else {
                                                                v1 = ""
                                                                if (0 != 1) {
                                                                    v59 -= (^(v59))
                                                                    fmt.Println(v57, fi84, v3)
                                                                    fmt.Println(v59, v1)
                                                                    if false {
                                                                        v92 := GenS9{F0: v89, F1: v1, F2: (v50 ^ 1)}
                                                                        var v93 int8 = 42
                                                                        v93 = v50
                                                                        _ = v92
                                                                        _ = v93
                                                                    }
                                                                }
                                                                _ = len(v46)
                                                                var v94 uint64 = 559
                                                                v95 := GenS8{F0: 255, F1: (v59 + 1), F2: (v86 ^ (v86 - v3))}
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                        default:
                                                            if (true || false) {
                                                                var v96 bool = false
                                                                var v97 uint16 = 255
                                                                var v98 int64 = (v2 + (v2 + v2))
                                                                _ = v96
                                                                _ = v97
                                                                _ = v98
                                                            }
                                                            fmt.Println(v59)
                                                        }
                                                    }
                                                    _ = v89
                                                }
                                            }
                                            if (!false) {
                                                var v99 float32 = -2.25
                                                switch (fi83 * fi84) {
                                                case -768:
                                                    var v100 uint = v59
                                                    _ = v100
                                                case 1:
                                                    if (!(!true)) {
                                                        v57 += 42
                                                    } else {
                                                        var v101 int = (fi84 | (^(fi83)))
                                                        var v102 uint = 18446744073709551615
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                case -307:
                                                    switch v1 {
                                                    case "foo bar":
                                                        v46 = append(v46, -3)
                                                    case "hello":
                                                        var v103 uint64 = 42
                                                        _ = v103
                                                    case "":
                                                        var v104 int16 = 42
                                                        _ = v104
                                                    default:
                                                        if true {
                                                            v105 := GenS13{F0: 42}
                                                            switch v1 {
                                                            case "\n\t":
                                                                if (!true) {
                                                                    switch fi84 {
                                                                    case -2147483648:
                                                                        var v106 uint64 = 2
                                                                        _ = v106
                                                                    }
                                                                    if (-62 <= (v50 | -1)) {
                                                                        if false {
                                                                            fmt.Println(v55)
                                                                            var v107 float32 = v55
                                                                            var v108 int64 = -1
                                                                            _ = v107
                                                                            _ = v108
                                                                        }
                                                                        v99 = (v55 * 1e10)
                                                                        var v109 int16 = -32768
                                                                        _ = v109
                                                                    } else {
                                                                        v58 += -2.25
                                                                        if (v1 > (("\n\t" + v1) + v1)) {
                                                                            if (231 <= (fi84 | (fi83 / (((fi83 >> ((v59) & 63))) | 1)))) {
                                                                                v1 = (v1 + v1)
                                                                            }
                                                                            v110 := GenS2{F0: (v59 + genFn0(626, 2))}
                                                                            _ = v110
                                                                        }
                                                                        v86 = (v86 * v3)
                                                                        fmt.Println(fi83, v55, v3)
                                                                        var v111 bool = false
                                                                        _ = v111
                                                                    }
                                                                    var v112 float32 = float32((v58 / ((v58 / v58) + v58)))
                                                                    _ = v112
                                                                }
                                                            }
                                                            var v113 int64 = (v2 / (((v2 / ((int64(v50)) | 1))) | 1))
                                                            v114 := GenS14{F0: fi83, F1: (-(v50)), F2: (v86 % (((v3 >> (((v59 << (((v59 << (((v59 ^ (v59 % ((v59) | 1)))) & 63))) & 63))) & 63))) | 1))}
                                                            var v115 uint64 = 2
                                                            _ = v105
                                                            _ = v113
                                                            _ = v114
                                                            _ = v115
                                                        } else {
                                                            if len(v46) > 0 {
                                                                v46[0] = fi84
                                                            }
                                                            if (genFn0(-68, 127) != (v59 << ((2) & 63))) {
                                                                switch (fi84 + fi83) {
                                                                case -1:
                                                                    var v116 uint8 = 127
                                                                    _ = v116
                                                                case 42:
                                                                    if true {
                                                                        v46 = append(v46, (^(fi84)))
                                                                    } else {
                                                                        var v117 uint16 = 255
                                                                        var v118 int = (fi84 - (fi84 / ((fi84) | 1)))
                                                                        if (!(1 <= 2147483647)) {
                                                                            if (!(fi84 == (fi84 / (((v118 % ((36) | 1))) | 1)))) {
                                                                                switch ("" + v1) {
                                                                                case "\n\t":
                                                                                    if len(v46) > 0 {
                                                                                        v46[0] = 370
                                                                                    }
                                                                                }
                                                                                if len(v46) >= 2 {
                                                                                    v46 = v46[1:len(v46)]
                                                                                }
                                                                                v99 = (-2.25 + 0.0)
                                                                                var v119 bool = false
                                                                                switch v1 {
                                                                                case "a":
                                                                                    fmt.Println(v117)
                                                                                case "foo bar":
                                                                                    v120 := GenS15{F0: -160, F1: int(v59), F2: int8(v117)}
                                                                                    _ = v120
                                                                                default:
                                                                                    var v121 int64 = (v2 + v2)
                                                                                    _ = v121
                                                                                }
                                                                                _ = v119
                                                                            }
                                                                            var v122 uint16 = (v117 >> ((18446744073709551615) & 63))
                                                                            _ = v122
                                                                        }
                                                                        _ = v117
                                                                        _ = v118
                                                                    }
                                                                default:
                                                                    fmt.Println(fi83)
                                                                }
                                                                var v123 uint64 = 951
                                                                _ = len(v46)
                                                                _ = v123
                                                            }
                                                        }
                                                    }
                                                }
                                                var v124 float64 = v58
                                                _ = v124
                                                _ = v99
                                            } else {
                                                var v125 string = "foo bar"
                                                _ = v125
                                            }
                                            _ = v86
                                        }
                                        v126 := GenS14{F0: 32767, F1: v57, F2: v3}
                                        fmt.Println(v55, v58)
                                        switch fi83 {
                                        case -1:
                                            fmt.Println(v59, v57, v1)
                                        }
                                        v2 -= (-(v2))
                                        _ = v126
                                    } else {
                                        var v127 int8 = v50
                                        v59 = 670
                                        _ = v127
                                    }
                                    var v128 int64 = 1
                                    _ = v128
                                }
                                _ = fi84
                            }
                            var v129 uint64 = 1
                            _ = fi83
                            _ = v129
                        }
                    }
                }
            case 9223372036854775807:
                v55 = v55
            case 183:
                v130 := GenS6{F0: 65535}
                _ = v130
            default:
                var v131 int32 = (v3 * (v3 - 2147483647))
                _ = v131
            }
            if true {
                v132 := GenS16{F0: false, F1: 192}
                var v133 uint64 = 4294967295
                _ = v132
                _ = v133
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v46
    _ = v50
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
