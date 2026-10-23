package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS12 struct { F0 string }
func genFn0(p0 string, p1 uint8) int64 {
    p0 = ""
    v7 := []int{4294967295}
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(fi8)
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 9223372036854775807
    var v3 uint32 = 706
    var v4 int = 42
    var v5 uint16 = 237
    var v6 float64 = 1e10
    v4 = 128
    if (!false) {
        if (false || (255 == 0)) {
            v3 += 2147483648
            v9 := []int{(^(v4)), (v2 * (v2 - (v2 + v4))), (v2 & v2), (v2 << ((0) & 63))}
            if len(v9) > 0 {
                v9[0] = (v4 - v4)
            }
            _ = v9
        }
        if false {
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := []int{fi10, (v2 << ((0) & 63)), (v2 | (-(v4)))}
                _ = fi10
                _ = v11
            }
            if false {
                for fi12 := 0; fi12 < 5; fi12++ {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        if (1 <= 128) {
                            if (("hello" + ("\n\t" + "foo bar")) > "a") {
                                v14 := GenS1{F0: 42}
                                v15 := []int{(fi12 | fi12), fi13}
                                _ = v14
                                _ = v15
                            }
                            var v16 int = (fi12 - fi12)
                            v3 += (v3 << ((275) & 63))
                            _ = v16
                        } else {
                            if (false && (!false)) {
                                v4 -= (fi13 / ((v2) | 1))
                            }
                        }
                        _ = fi13
                    }
                    _ = fi12
                }
                for fi17 := 0; fi17 < 5; fi17++ {
                    v18 := []int{42, fi17, 423}
                    for fi19 := 0; fi19 < 12; fi19++ {
                        fmt.Println(v5, fi19)
                        _ = fi19
                    }
                    fmt.Println(v2, v3, fi17)
                    _ = fi17
                    _ = v18
                }
                var v20 int8 = -1
                _ = v20
            } else {
                if false {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        v6 -= ((v6 - 1e10) * 0.1)
                        for fi22 := 0; fi22 < 9; fi22++ {
                            var v23 uint64 = 759
                            var v24 int32 = -610
                            if (false && false) {
                                var v25 uint8 = 42
                                fmt.Println(v25, v4, v6)
                                _ = v25
                            } else {
                                var v26 int8 = 42
                                switch "foo bar" {
                                case "hello":
                                    switch fi22 {
                                    case -2147483648:
                                        v27 := []int{fi21, fi22}
                                        _ = v27
                                    case 32768:
                                        var v28 uint = 742
                                        _ = v28
                                    case 127:
                                        v3 -= v3
                                    default:
                                        var v29 float32 = (1.5 + -2.25)
                                        _ = v29
                                    }
                                case "":
                                    var v30 uint16 = (v5 << ((18446744073709551615) & 63))
                                    _ = v30
                                case "a":
                                    switch fi21 {
                                    case 32768:
                                        var v31 bool = (!false)
                                        v32 := GenS1{F0: v26}
                                        _ = v31
                                        _ = v32
                                    case 127:
                                        var v33 uint8 = 42
                                        _ = v33
                                    case 2147483647:
                                        switch "" {
                                        case "a":
                                            var v34 int32 = v24
                                            _ = v34
                                        case "":
                                            v35 := GenS1{F0: 100}
                                            _ = v35
                                        default:
                                            var v36 float64 = (0.0 / v6)
                                            _ = v36
                                        }
                                    default:
                                        if (186 <= -204) {
                                            v4 -= (-(v4))
                                            v37 := GenS1{F0: v26}
                                            _ = v37
                                        } else {
                                            v26 += (-(v26))
                                        }
                                    }
                                default:
                                    var v38 uint8 = 2
                                    var v39 uint8 = v38
                                    v40 := GenS2{F0: 42, F1: 149}
                                    var v41 uint8 = (v38 & v39)
                                    v38 += (v39 ^ (^(v38)))
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                    _ = v41
                                }
                                var v42 bool = (!true)
                                _ = v26
                                _ = v42
                            }
                            v2 += 848
                            v4 = int(v23)
                            var v43 bool = true
                            if (!(!v43)) {
                                var v44 string = "a"
                                switch fi21 {
                                case 0:
                                    var v45 int32 = (v24 + (v24 & -373))
                                    _ = v45
                                }
                                _ = v44
                            }
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v43
                        }
                        _ = fi21
                    }
                }
                var v46 int32 = -1
                _ = v46
            }
        } else {
            v2 -= (v4 * v2)
        }
        for fi47 := 0; fi47 < 2; fi47++ {
            v6 = v6
            _ = fi47
        }
    }
    fmt.Println(v3, v2, v2)
    v3 = (^(v3))
    var v48 int8 = 127
    if false {
        switch (((("" + "foo bar") + ("" + "hello")) + "\n\t") + ("hello" + "hello")) {
        case "hello":
            switch (^(v2)) {
            case 127:
                v49 := []int{}
                _ = v49
            }
        case "foo bar":
            for fi50 := 0; fi50 < 1; fi50++ {
                v51 := v1[(v2 - 128)]
                var v52 int16 = 332
                _ = fi50
                _ = v51
                _ = v52
            }
        }
    }
    v5 = (^(v5))
    fmt.Println(v2)
    v3 += v3
    v53 := v1[v2]
    v54 := GenS1{F0: (-(v48))}
    var v55 int8 = (v48 & -128)
    switch "\n\t" {
    case "":
        fmt.Println(v5, v55)
    case "a":
        fmt.Println(v2, v6, v4)
    case "foo bar":
        if true {
            v2 += (v4 << ((65535) & 63))
        } else {
            if (!((true && (((v4 * v4) > (-(v4))) || false)) && true)) {
                for fi56 := 0; fi56 < 7; fi56++ {
                    fmt.Println(v2, v6, fi56)
                    _ = fi56
                }
            } else {
                for fi57 := 0; fi57 < 14; fi57++ {
                    v58 := GenS1{F0: v48}
                    v59 := GenS2{F0: -908, F1: (v5 % (((v5 ^ v5)) | 1))}
                    _ = fi57
                    _ = v58
                    _ = v59
                }
                delete(v1, v4)
                for fi60 := 0; fi60 < 8; fi60++ {
                    var v61 int32 = 2147483647
                    _ = fi60
                    _ = v61
                }
            }
            var v62 float64 = 0.0
            for fi63 := 0; fi63 < 4; fi63++ {
                v64 := GenS1{F0: v55}
                switch -2147483648 {
                case 786:
                    if ((v5 & (v5 | 65535)) == v5) {
                        fmt.Println(v3, v62, v2)
                    } else {
                        if (false && true) {
                            v65 := GenS1{F0: int8(v4)}
                            if (v6 == v62) {
                                for fi66 := 0; fi66 < 20; fi66++ {
                                    switch v53 {
                                    case "hello":
                                        var v67 uint = 529
                                        _ = v67
                                    default:
                                        v68 := []int{(fi63 + (v2 % (((fi66 ^ fi63)) | 1))), (v4 % ((v2) | 1)), int(v48), (fi66 & (fi66 ^ v4))}
                                        _ = v68
                                    }
                                    var v69 uint8 = 1
                                    v70 := []int{42, fi66}
                                    var v71 float32 = 3.14159
                                    _ = fi66
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                }
                                v62 = 3.14159
                            } else {
                                switch v53 {
                                case "\n\t":
                                    for fi72 := 0; fi72 < 1; fi72++ {
                                        v62 = -2.25
                                        v73 := GenS1{F0: (v55 % (((v48 | v48)) | 1))}
                                        v74 := GenS3{F0: 127, F1: -642}
                                        _ = fi72
                                        _ = v73
                                        _ = v74
                                    }
                                default:
                                    v55 += v48
                                }
                                _ = len(v1)
                            }
                            if true {
                                v75 := GenS2{F0: -616, F1: (^(v5))}
                                if (42 == v5) {
                                    v48 -= 127
                                    v53 = v53
                                    v55 = v48
                                } else {
                                    v76 := GenS2{F0: -1, F1: v5}
                                    var v77 bool = (((v53 + (v53 + v53)) + ((v53 + v53) + (("\n\t" + "") + (("hello" + v53) + "foo bar")))) != v53)
                                    var v78 bool = (!(float64((0.0 / -2.25)) >= v6))
                                    _ = v76
                                    _ = v77
                                    _ = v78
                                }
                                v79 := []int{v4, 32768, 128, v2}
                                _ = v75
                                _ = v79
                            } else {
                                v48 -= (v55 % (((-(v48))) | 1))
                            }
                            _ = v65
                        } else {
                            for fi80 := 0; fi80 < 17; fi80++ {
                                fmt.Println(v4)
                                _ = fi80
                            }
                            switch 32767 {
                            case 32768:
                                fmt.Println(v55, v4)
                            case -57:
                                v81 := GenS3{F0: 42, F1: 42}
                                _ = v81
                            case 128:
                                for fi82 := 0; fi82 < 13; fi82++ {
                                    if false {
                                        v83 := GenS2{F0: 974, F1: (v5 / (((v5 << ((18446744073709551615) & 63))) | 1))}
                                        var v84 float32 = float32(1e10)
                                        var v85 int64 = 1
                                        _ = v83
                                        _ = v84
                                        _ = v85
                                    }
                                    if true {
                                        fmt.Println(v55, v4)
                                        if true {
                                            _ = len(v1)
                                            switch "a" {
                                            case "\n\t":
                                                switch v2 {
                                                case 2147483647:
                                                    fmt.Println(v53)
                                                    if false {
                                                        switch ("foo bar" + "foo bar") {
                                                        case "a":
                                                            v86 := v1[(v4 + v2)]
                                                            _ = v86
                                                        case "":
                                                            switch v53 {
                                                            case "a":
                                                                var v87 uint32 = 4294967295
                                                                _ = v87
                                                            case "hello":
                                                                var v88 uint = 18446744073709551615
                                                                _ = v88
                                                            }
                                                            v89 := GenS1{F0: v55}
                                                            var v90 uint64 = 514
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                        var v91 uint16 = (v5 / (((v5 - v5)) | 1))
                                                        _ = v91
                                                    } else {
                                                        var v92 float32 = 3.14159
                                                        switch 426 {
                                                        case -9223372036854775808:
                                                            if false {
                                                                v93 := GenS4{F0: 127}
                                                                _ = v93
                                                            }
                                                        case -2147483648:
                                                            v53 = (v53 + ("" + v53))
                                                        case -1:
                                                            switch 9223372036854775807 {
                                                            case -340:
                                                                var v94 bool = false
                                                                _ = v94
                                                            case 42:
                                                                v95 := GenS5{F0: 2147483647, F1: 19}
                                                                _ = v95
                                                            case 919:
                                                                v3 -= (v3 ^ v3)
                                                            default:
                                                                var v96 int = (fi63 - (^(v4)))
                                                                _ = v96
                                                            }
                                                        }
                                                        fmt.Println(v92)
                                                        _ = len(v1)
                                                        v5 -= (v5 * (^(v5)))
                                                        _ = v92
                                                    }
                                                    var v97 uint32 = (v3 * v3)
                                                    _ = v97
                                                case 1:
                                                    var v98 float64 = (float64(float32(v62)) + 0.0)
                                                    _ = v98
                                                case 42:
                                                    switch fi82 {
                                                    case 584:
                                                        v99 := GenS6{F0: 255, F1: "a"}
                                                        _ = v99
                                                    default:
                                                        var v100 int16 = 881
                                                        _ = v100
                                                    }
                                                default:
                                                    v2 = v2
                                                }
                                            case "a":
                                                var v101 uint = 65535
                                                _ = v101
                                            case "foo bar":
                                                if true {
                                                    v102 := GenS3{F0: (fi82 & 32768), F1: genFn0(("hello" + "hello"), 252)}
                                                    v103 := GenS3{F0: -120, F1: 0}
                                                    _ = v102
                                                    _ = v103
                                                } else {
                                                    var v104 uint16 = (v5 & v5)
                                                    if (1 <= 1000) {
                                                        if (true || false) {
                                                            var v105 uint = 2
                                                            var v106 int32 = -1
                                                            _ = v105
                                                            _ = v106
                                                        } else {
                                                            v107 := GenS7{F0: (166 == 127), F1: v3}
                                                            var v108 uint64 = 4294967295
                                                            _ = v107
                                                            _ = v108
                                                        }
                                                        if true {
                                                            switch v4 {
                                                            case -446:
                                                                fmt.Println(v53, v48, v3)
                                                            case -591:
                                                                if ((true || (18446744073709551615 < 1)) && true) {
                                                                    if (((1 <= uint(v3)) && (true && (v53 < ""))) || (!(0 < 1000))) {
                                                                        if (true && false) {
                                                                            if ((!(fi82 != (v4 + v2))) && (false && (!false))) {
                                                                                v109 := GenS8{F0: -1}
                                                                                v110 := GenS4{F0: (v55 & v55)}
                                                                                if true {
                                                                                    v111 := GenS5{F0: -215, F1: 191}
                                                                                    var v112 string = ("hello" + v53)
                                                                                    _ = v111
                                                                                    _ = v112
                                                                                } else {
                                                                                    v113 := []int{}
                                                                                    v114 := GenS9{F0: 1}
                                                                                    _ = v113
                                                                                    _ = v114
                                                                                }
                                                                                _ = v109
                                                                                _ = v110
                                                                            }
                                                                            if (4294967295 >= 636) {
                                                                                v48 = (v55 & -1)
                                                                                v115 := v1[fi82]
                                                                                _ = v115
                                                                            }
                                                                            v116 := []int{-9223372036854775808, v4}
                                                                            _ = v116
                                                                        }
                                                                        var v117 int8 = 127
                                                                        if (uint(v3) < 1) {
                                                                            v117 = v48
                                                                            switch v53 {
                                                                            case "hello":
                                                                                fmt.Println(v2, v2)
                                                                            case "foo bar":
                                                                                switch 32767 {
                                                                                case 4294967295:
                                                                                    if true {
                                                                                        v104 = (v104 ^ v104)
                                                                                    } else {
                                                                                        if true {
                                                                                            v104 = v104
                                                                                        }
                                                                                        v48 = v48
                                                                                        if true {
                                                                                            var v118 uint8 = 1
                                                                                            _ = v118
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v104 = (v5 % (((v104 << ((500) & 63))) | 1))
                                                                                }
                                                                            case "a":
                                                                                var v119 int8 = v55
                                                                                _ = v119
                                                                            }
                                                                            var v120 string = "hello"
                                                                            _ = v120
                                                                        }
                                                                        switch v53 {
                                                                        case "":
                                                                            v121 := GenS10{F0: ((((((("foo bar" + v53) + "\n\t") + v53) + ("foo bar" + v53)) + v53) + "hello") + v53)}
                                                                            _ = v121
                                                                        case "foo bar":
                                                                            if (-420 != -59) {
                                                                                var v122 bool = true
                                                                                _ = v122
                                                                            } else {
                                                                                var v123 int32 = -726
                                                                                v124 := GenS3{F0: fi82, F1: 1}
                                                                                _ = v123
                                                                                _ = v124
                                                                            }
                                                                        }
                                                                        v3 -= v3
                                                                        if true {
                                                                            v125 := GenS1{F0: v48}
                                                                            fmt.Println(v48, v6)
                                                                            _ = v125
                                                                        } else {
                                                                            v3 += v3
                                                                            v104 = v5
                                                                            if false {
                                                                                switch (v4 & (^(v2))) {
                                                                                case 127:
                                                                                    var v126 uint64 = 42
                                                                                    _ = v126
                                                                                case -241:
                                                                                    var v127 int = fi63
                                                                                    _ = v127
                                                                                case 4294967295:
                                                                                    v53 = (((v53 + (v53 + v53)) + ((v53 + "\n\t") + v53)) + v53)
                                                                                default:
                                                                                    switch 965 {
                                                                                    case 32768:
                                                                                        v4 += v4
                                                                                        v104 += (v104 & (v104 ^ (v5 >> ((2) & 63))))
                                                                                    case 894:
                                                                                        v53 = (v53 + ("hello" + v53))
                                                                                    }
                                                                                }
                                                                                v5 += 32767
                                                                            }
                                                                        }
                                                                        _ = v117
                                                                    } else {
                                                                        v128 := GenS11{F0: (v3 & 4294967295), F1: v5, F2: 32767}
                                                                        _ = v128
                                                                    }
                                                                    v129 := GenS11{F0: 676, F1: 42, F2: -201}
                                                                    _ = v129
                                                                } else {
                                                                    var v130 float64 = float64((float32(v62) * (0.0 - 0.1)))
                                                                    _ = v130
                                                                }
                                                            case -9223372036854775808:
                                                                v5 = (v104 | (v5 >> ((255) & 63)))
                                                            default:
                                                                v131 := GenS3{F0: -192, F1: 9223372036854775807}
                                                                _ = v131
                                                            }
                                                            v3 = v3
                                                            fmt.Println(v104, v3, v53)
                                                            var v132 float32 = (float32(float64(0.0)) - 1.5)
                                                            _ = v132
                                                        } else {
                                                            var v133 float64 = (((v62 - v62) / v6) - float64(0.0))
                                                            fmt.Println(fi82, fi82, v48)
                                                            _ = v133
                                                        }
                                                        var v134 int32 = -1
                                                        _ = v134
                                                    } else {
                                                        delete(v1, (v4 | fi82))
                                                    }
                                                    _ = v104
                                                }
                                            }
                                            v6 += v6
                                            fmt.Println(fi82, v3, v4)
                                            v135 := GenS12{F0: v53}
                                            _ = v135
                                        } else {
                                            if ((^(v3)) <= v3) {
                                                var v136 int = (v4 % ((int(v3)) | 1))
                                                _ = len(v1)
                                                var v137 uint8 = 128
                                                v138 := v1[0]
                                                _ = v136
                                                _ = v137
                                                _ = v138
                                            }
                                            if (2 > 18446744073709551615) {
                                                var v139 uint8 = 71
                                                v140 := GenS5{F0: 0, F1: v139}
                                                var v141 uint32 = v3
                                                _ = v139
                                                _ = v140
                                                _ = v141
                                            } else {
                                                _ = len(v1)
                                                fmt.Println(v48, v55)
                                            }
                                        }
                                    }
                                    if (!(!false)) {
                                        v142 := []int{(fi63 * -275), v4, (v4 << ((42) & 63))}
                                        _ = v142
                                    }
                                    _ = fi82
                                }
                            }
                        }
                        _ = len(v1)
                        fmt.Println(fi63)
                    }
                case 0:
                    var v143 int8 = (-(v55))
                    _ = v143
                default:
                    var v144 int = 252
                    _ = v144
                }
                _ = fi63
                _ = v64
            }
            _ = v62
        }
    }
    for fi145 := 0; fi145 < 5; fi145++ {
        var v146 uint8 = 255
        _ = fi145
        _ = v146
    }
    switch v53 {
    case "hello":
        for fi147 := 0; fi147 < 5; fi147++ {
            v148 := v1[v2]
            v148 = ("hello" + v53)
            _ = fi147
            _ = v148
        }
    case "\n\t":
        var v149 int32 = -2147483648
        _ = v149
    case "foo bar":
        fmt.Println(v48)
    }
    for fi150 := 0; fi150 < 2; fi150++ {
        v53 = v53
        var v151 float32 = 3.14159
        v152 := GenS7{F0: false, F1: v3}
        _ = fi150
        _ = v151
        _ = v152
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
