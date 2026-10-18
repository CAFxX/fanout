package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 int16; F2 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string; F1 uint16; F2 int32 }
type GenS7 struct { F0 int8; F1 string; F2 uint8 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int16; F1 int64 }
type GenS10 struct { F0 uint; F1 bool }
type GenS11 struct { F0 int32; F1 uint; F2 uint8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 bool; F1 int; F2 string }
type GenS15 struct { F0 bool; F1 bool; F2 uint }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint32; F1 int16 }
type GenS18 struct { F0 uint8; F1 string; F2 int64 }
type GenS19 struct { F0 int16 }
func genFn0(p0 int8, p1 int16, p2 int8) uint16 {
    p2 = p0
    var v4 int32 = -1
    if false {
        if false {
            v5 := []int{-35, -2147483648, int(p2), 32768}
            switch "hello" {
            case "hello":
                var v6 uint16 = 173
                _ = v6
            case "a":
                if true {
                    var v7 uint8 = 255
                    _ = v7
                } else {
                    p1 += p1
                    if (0 == 42) {
                        for fi8 := 0; fi8 < 10; fi8++ {
                            var v9 float32 = float32(((3.14159 + 3.14159) + -2.25))
                            for fi10 := 0; fi10 < 18; fi10++ {
                                fmt.Println(p1, v4)
                                var v11 uint64 = 1
                                var v12 uint8 = 127
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                            _ = fi8
                            _ = v9
                        }
                    } else {
                        var v13 uint32 = 2147483647
                        _ = v13
                    }
                }
            default:
                var v14 uint8 = 127
                _ = v14
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v5
        } else {
            var v15 uint8 = 2
            v16 := GenS1{F0: (2147483648 != 2147483648)}
            _ = v15
            _ = v16
        }
        fmt.Println(p1, p0, p1)
    } else {
        v17 := []int{9223372036854775807, 428}
        p2 = 1
        if false {
            p0 = 118
            var v18 uint16 = 2
            _ = v18
        }
        _ = v17
    }
    if (false && false) {
        if (!true) {
            v19 := GenS2{F0: 597}
            var v20 int32 = v4
            _ = v19
            _ = v20
        } else {
            v21 := []int{-2147483648}
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 uint32) int64 {
    var v22 int16 = 1
    var v23 int64 = 9223372036854775807
    p0 -= p0
    _ = p0
    _ = v22
    _ = v23
    return (v23 - v23)
}

func main() {
    v1 := []string{"\n\t"}
    var v2 int16 = -462
    v3 := make(map[int]string)
    switch "foo bar" {
    case "a":
        v2 += 32767
    default:
        if true {
            for fi24 := 0; fi24 < 9; fi24++ {
                v25 := GenS3{F0: false}
                if true {
                    if true {
                        v2 += (v2 << ((1) & 63))
                        var v26 int8 = 42
                        var v27 bool = false
                        _ = v26
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 5; fi28++ {
                            if true {
                                v29 := GenS4{F0: (-(v2)), F1: (v2 << ((42) & 63)), F2: 0}
                                if (true && false) {
                                    if ((v2 ^ v2) <= v2) {
                                        v2 += v2
                                        v2 += (^(v2))
                                    }
                                    v2 += 182
                                }
                                var v30 string = ""
                                _ = v29
                                _ = v30
                            } else {
                                var v31 int8 = 127
                                if ((!(!false)) && ((((float32(0.1) + -2.25) > -2.25) && (!(-2.25 >= 0.0))) || (255 <= 578))) {
                                    fmt.Println(v2)
                                    v31 -= -1
                                    switch "" {
                                    case "\n\t":
                                        var v32 float64 = (0.0 - ((-2.25 + 0.0) * ((0.0 - 1e10) / -2.25)))
                                        fmt.Println(fi28, v2, v2)
                                        var v33 uint = 614
                                        _ = v32
                                        _ = v33
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        v34 := GenS5{F0: 787}
                                        _ = v34
                                    case "a":
                                        v35 := GenS6{F0: "", F1: genFn0(int8(v2), v2, v31), F2: 127}
                                        if len(v1) > 0 {
                                            v1[0] = ((((("hello" + "hello") + ("" + "hello")) + "\n\t") + ("hello" + "\n\t")) + (("" + "a") + ""))
                                        }
                                        fmt.Println(v31, fi24)
                                        _ = v35
                                    }
                                    fmt.Println(fi28)
                                    var v36 float64 = (-2.25 + float64(1e10))
                                    if ((fi28 << ((42) & 63)) >= (fi24 / ((fi28) | 1))) {
                                        v37 := GenS2{F0: 65535}
                                        var v38 bool = false
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = v36
                                }
                                if false {
                                    v39 := v3[fi24]
                                    delete(v3, (-(fi24)))
                                    _ = v39
                                } else {
                                    v40 := GenS7{F0: 127, F1: "foo bar", F2: 42}
                                    if false {
                                        fmt.Println(v2, fi28, v2)
                                    }
                                    _ = v40
                                }
                                _ = v31
                            }
                            var v41 int16 = v2
                            _ = fi28
                            _ = v41
                        }
                        switch fi24 {
                        case -340:
                            switch ("foo bar" + "a") {
                            case "a":
                                switch (fi24 << ((18446744073709551615) & 63)) {
                                case -18:
                                    v42 := GenS4{F0: v2, F1: v2, F2: genFn1(0)}
                                    _ = v42
                                case 54:
                                    v2 -= (v2 * v2)
                                case -9223372036854775808:
                                    switch "hello" {
                                    case "hello":
                                        if (-25 == 42) {
                                            for fi43 := 0; fi43 < 5; fi43++ {
                                                v2 = 858
                                                _ = fi43
                                            }
                                        } else {
                                            var v44 float32 = 0.0
                                            fmt.Println(v2)
                                            _ = v44
                                        }
                                    }
                                }
                            }
                        case 739:
                            if (3.14159 < -2.25) {
                                fmt.Println(v2, fi24)
                                var v45 uint8 = 57
                                if true {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v46 float32 = 0.1
                                    fmt.Println(v2, fi24, v46)
                                    _ = v46
                                }
                                _ = v45
                            } else {
                                switch fi24 {
                                case 795:
                                    var v47 int = (-(fi24))
                                    _ = v47
                                case -232:
                                    v2 -= (^(v2))
                                default:
                                    fmt.Println(fi24, v2)
                                }
                                fmt.Println(fi24)
                            }
                        case 1:
                            fmt.Println(v2, v2)
                        default:
                            v2 = v2
                        }
                    }
                    v48 := GenS4{F0: (^(v2)), F1: (v2 ^ 0), F2: 9223372036854775807}
                    v49 := GenS3{F0: (42 > 65535)}
                    _ = v48
                    _ = v49
                }
                _ = fi24
                _ = v25
            }
            v50 := GenS5{F0: 2}
            _ = v50
        } else {
            fmt.Println(v2)
            v2 += (v2 | v2)
            for fi51 := 0; fi51 < 20; fi51++ {
                fmt.Println(fi51)
                var v52 bool = false
                fmt.Println(v52, v52, v52)
                _ = fi51
                _ = v52
            }
        }
    }
    v53 := v3[-317]
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println(v53)
    if false {
        _ = len(v1)
    }
    if true {
        fmt.Println(v2)
        for fi54 := 0; fi54 < 2; fi54++ {
            var v55 int8 = -1
            var v56 bool = ((true || (84 < 87)) && (-846 == -1))
            v53 = ""
            _ = fi54
            _ = v55
            _ = v56
        }
    } else {
        switch 878 {
        case 9223372036854775807:
            var v57 string = v53
            _ = v57
        case -126:
            _ = len(v1)
        case 127:
            var v58 float64 = 1e10
            _ = v58
        }
        var v59 uint64 = 782
        var v60 int16 = (v2 + (^(v2)))
        _ = v59
        _ = v60
    }
    v53 = v53
    var v61 int = int(v2)
    _ = len(v1)
    v62 := GenS8{F0: 38}
    v63 := GenS9{F0: (v2 * 1), F1: 9223372036854775807}
    switch v53 {
    case "hello":
        if true {
            if ((v61 ^ (v61 * (v61 - (v61 % (((-(v61))) | 1))))) < v61) {
                switch v61 {
                case 0:
                    delete(v3, v61)
                default:
                    var v64 uint8 = 0
                    _ = v64
                }
                v2 = (v2 >> ((42) & 63))
                var v65 uint16 = 32767
                _ = v65
            } else {
                var v66 int64 = -1
                _ = v66
            }
            var v67 uint8 = 2
            _ = v67
        } else {
            switch (v61 / ((v61) | 1)) {
            case 2147483647:
                if false {
                    v2 = (v2 ^ v2)
                }
            default:
                v68 := GenS10{F0: 431, F1: true}
                _ = v68
            }
            for fi69 := 0; fi69 < 19; fi69++ {
                if true {
                    var v70 int64 = 9223372036854775807
                    var v71 bool = (128 <= 1)
                    _ = v70
                    _ = v71
                }
                if (1e10 < 1e10) {
                    var v72 float32 = 1.5
                    v72 = v72
                    if true {
                        v73 := GenS11{F0: -2147483648, F1: 943, F2: 255}
                        _ = v73
                    } else {
                        if (false && (!((true || true) && (!(!true))))) {
                            if false {
                                fmt.Println(v72, v61, v53)
                            }
                            if false {
                                for fi74 := 0; fi74 < 15; fi74++ {
                                    v72 = v72
                                    var v75 uint8 = 0
                                    v76 := GenS12{F0: v75}
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                }
                            }
                            var v77 uint = 4294967295
                            _ = v77
                        } else {
                            if (false && true) {
                                var v78 uint16 = 42
                                var v79 uint16 = v78
                                var v80 int16 = (v2 & 32767)
                                _ = v78
                                _ = v79
                                _ = v80
                            }
                            switch "" {
                            case "hello":
                                for fi81 := 0; fi81 < 7; fi81++ {
                                    if true {
                                        switch -1 {
                                        case 1:
                                            var v82 int8 = -128
                                            _ = v82
                                        case 313:
                                            switch ((("foo bar" + v53) + v53) + "") {
                                            case "hello":
                                                var v83 uint16 = 417
                                                _ = v83
                                            case "\n\t":
                                                var v84 int64 = 715
                                                _ = v84
                                            case "a":
                                                v3[v61] = ((v53 + "hello") + ("" + v53))
                                            default:
                                                delete(v3, -1)
                                            }
                                        case 4294967295:
                                            fmt.Println(fi69)
                                        default:
                                            v1 = append(v1, "\n\t")
                                        }
                                        switch v61 {
                                        case 128:
                                            v61 = -2147483648
                                        case 653:
                                            if ("\n\t" == (((((v53 + (v53 + v53)) + v53) + v53) + v53) + "")) {
                                                var v85 uint64 = 139
                                                _ = v85
                                            }
                                        case -23:
                                            var v86 uint16 = 42
                                            _ = v86
                                        }
                                        switch v53 {
                                        case "\n\t":
                                            v87 := GenS3{F0: (!true)}
                                            _ = v87
                                        case "":
                                            v88 := GenS13{F0: (v2 | 1)}
                                            _ = v88
                                        default:
                                            var v89 float32 = float32(0.0)
                                            _ = v89
                                        }
                                    } else {
                                        v2 = 1
                                    }
                                    _ = fi81
                                }
                            case "":
                                for fi90 := 0; fi90 < 1; fi90++ {
                                    v53 = v53
                                    if true {
                                        if (42 > 0) {
                                            if true {
                                                v3[(fi69 | fi69)] = ("foo bar" + v53)
                                                var v91 int = fi69
                                                if ((-2.25 * 3.14159) >= 3.14159) {
                                                    var v92 float64 = (1.5 * (1.5 + float64(v72)))
                                                    _ = v92
                                                } else {
                                                    var v93 uint32 = 42
                                                    var v94 uint = uint(v93)
                                                    if (((v94 << ((uint(v93)) & 63)) > v94) && (2 < 2)) {
                                                        if false {
                                                            if false {
                                                                if true {
                                                                    v53 = "a"
                                                                    if ((-2.25 * (1e10 + (((1e10 - (((((-2.25 + 0.0) / -2.25) / 1.5) + ((1e10 + float64(((v72 - 0.1) - v72))) + ((-2.25 * 3.14159) * 1e10))) - -2.25)) + (1.5 - -2.25)) + (0.0 - 0.1)))) != 1e10) {
                                                                        v53 = "\n\t"
                                                                    } else {
                                                                        switch v53 {
                                                                        case "a":
                                                                            v72 = 1e10
                                                                        case "hello":
                                                                            var v95 bool = true
                                                                            _ = v95
                                                                        }
                                                                        if false {
                                                                            v2 = (v2 + v2)
                                                                            var v96 uint32 = 1
                                                                            switch -9223372036854775808 {
                                                                            case -1:
                                                                                v72 -= 3.14159
                                                                            case 9223372036854775807:
                                                                                if true {
                                                                                    switch 42 {
                                                                                    case 1:
                                                                                        v97 := GenS3{F0: false}
                                                                                        _ = v97
                                                                                    case -819:
                                                                                        if false {
                                                                                            if false {
                                                                                                if (!true) {
                                                                                                    if (-2147483648 != 471) {
                                                                                                        fmt.Println(v72, v2)
                                                                                                    }
                                                                                                } else {
                                                                                                    v98 := GenS14{F0: (255 > genFn0(42, 42, -1)), F1: v91, F2: v53}
                                                                                                    _ = v98
                                                                                                }
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    var v99 bool = (-128 <= 0)
                                                                                                    _ = v99
                                                                                                case "":
                                                                                                    if (127 > -70) {
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                        v2 = v2
                                                                                                    }
                                                                                                case "hello":
                                                                                                    v100 := GenS1{F0: false}
                                                                                                    _ = v100
                                                                                                default:
                                                                                                    fmt.Println(v96)
                                                                                                }
                                                                                                v101 := GenS15{F0: true, F1: true, F2: v94}
                                                                                                _ = v101
                                                                                            }
                                                                                            v96 = (v93 / ((v96) | 1))
                                                                                            var v102 int16 = -163
                                                                                            _ = v102
                                                                                        }
                                                                                    }
                                                                                    var v103 uint16 = 0
                                                                                    _ = v103
                                                                                } else {
                                                                                    if true {
                                                                                        v93 -= v93
                                                                                    } else {
                                                                                        v104 := GenS16{F0: 108}
                                                                                        if (42 < 32767) {
                                                                                            var v105 float32 = v72
                                                                                            var v106 int64 = 42
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                        v107 := v3[(v61 >> ((v94) & 63))]
                                                                                        _ = v104
                                                                                        _ = v107
                                                                                    }
                                                                                    var v108 uint = v94
                                                                                    fmt.Println(v72, v61, v61)
                                                                                    _ = v108
                                                                                }
                                                                            case 960:
                                                                                v109 := GenS4{F0: v2, F1: v2, F2: 1}
                                                                                _ = v109
                                                                            }
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                } else {
                                                                    v110 := GenS14{F0: ((v94 + (^(v94))) >= (^(v94))), F1: 32767, F2: (v53 + v53)}
                                                                    _ = v110
                                                                }
                                                                fmt.Println(fi90, v93)
                                                            } else {
                                                                v94 -= (^(v94))
                                                                v93 = (v93 << ((v94) & 63))
                                                                v93 = v93
                                                                _ = len(v1)
                                                            }
                                                            v111 := GenS17{F0: v93, F1: 203}
                                                            _ = v111
                                                        } else {
                                                            var v112 string = v53
                                                            _ = v112
                                                        }
                                                        var v113 int32 = 911
                                                        fmt.Println(v53, v94)
                                                        _ = v113
                                                    } else {
                                                        v72 += 1e10
                                                        v93 -= v93
                                                        var v114 uint16 = 0
                                                        _ = v114
                                                    }
                                                    v115 := GenS18{F0: 0, F1: (("" + v53) + "\n\t"), F2: 9223372036854775807}
                                                    _ = v115
                                                    _ = v93
                                                    _ = v94
                                                }
                                                _ = v91
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        } else {
                                            if (0 < 4294967295) {
                                                fmt.Println(fi90, v61, v53)
                                            }
                                            if (!false) {
                                                var v116 uint8 = 1
                                                switch ("hello" + v53) {
                                                case "hello":
                                                    fmt.Println(fi69, v2, v61)
                                                case "":
                                                    switch fi69 {
                                                    case 128:
                                                        v117 := GenS7{F0: 1, F1: v53, F2: 128}
                                                        _ = v117
                                                    case 4294967295:
                                                        v118 := GenS4{F0: v2, F1: -983, F2: 1}
                                                        _ = v118
                                                    case 439:
                                                        if true {
                                                            v119 := GenS10{F0: 4294967295, F1: false}
                                                            _ = v119
                                                        }
                                                    default:
                                                        if true {
                                                            if (v53 < v53) {
                                                                var v120 float32 = (v72 * (-2.25 / v72))
                                                                v121 := GenS6{F0: "foo bar", F1: 65535, F2: -1}
                                                                _ = v120
                                                                _ = v121
                                                            }
                                                            v72 = float32((-2.25 * (0.1 + 3.14159)))
                                                        } else {
                                                            v61 = (fi69 | (-(v61)))
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(fi69, v53)
                                                }
                                                v61 = 42
                                                _ = v116
                                            }
                                            var v122 float64 = float64((0.1 + (float32(float64(v72)) * (v72 + 1e10))))
                                            _ = v122
                                        }
                                        var v123 float32 = ((0.0 * (v72 * (1e10 + v72))) - v72)
                                        v61 -= v61
                                        _ = v123
                                    } else {
                                        var v124 uint8 = 159
                                        switch (fi69 >> ((0) & 63)) {
                                        case -663:
                                            if (2 > 2) {
                                                v124 = v124
                                                switch (fi69 * (fi69 << ((65535) & 63))) {
                                                case 1:
                                                    v72 -= 3.14159
                                                case 682:
                                                    var v125 int8 = 42
                                                    _ = v125
                                                case 127:
                                                    v126 := GenS13{F0: (v2 + (v2 >> ((uint(v124)) & 63)))}
                                                    _ = v126
                                                }
                                                var v127 string = v53
                                                fmt.Println(fi90, v61, v124)
                                                v128 := GenS19{F0: v2}
                                                fmt.Println(fi90)
                                                v127 = ((v127 + "a") + v127)
                                                var v129 float32 = 0.1
                                                v130 := v3[fi90]
                                                var v131 uint = 1
                                                var v132 int8 = 100
                                                v72 -= 1.5
                                                _ = v127
                                                _ = v128
                                                _ = v129
                                                _ = v130
                                                _ = v131
                                                _ = v132
                                            }
                                            var v133 uint8 = (^(v124))
                                            var v134 uint16 = 797
                                            _ = v133
                                            _ = v134
                                        }
                                        fmt.Println(fi69, v2, v124)
                                        _ = v124
                                    }
                                    var v135 uint64 = 42
                                    _ = fi90
                                    _ = v135
                                }
                            default:
                                _ = len(v1)
                            }
                        }
                        var v136 uint8 = 181
                        v137 := GenS1{F0: true}
                        _ = v136
                        _ = v137
                    }
                    _ = v72
                } else {
                    fmt.Println(v2)
                    v138 := GenS4{F0: (v2 >> ((65535) & 63)), F1: (v2 - (v2 % ((-32768) | 1))), F2: 132}
                    _ = v138
                }
                _ = fi69
            }
        }
    default:
        v2 = (v2 / (((v2 * (v2 >> ((42) & 63)))) | 1))
    }
    _ = len(v3)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 -= 32767
    v2 = (v2 + v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v53:", v53)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v53
    _ = v61
    _ = v62
    _ = v63
}
