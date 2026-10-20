package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 bool; F2 string }
type GenS7 struct { F0 int32; F1 int8; F2 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint8; F1 int }
func genFn0(p0 uint64, p1 bool) int {
    p0 -= 1
    var v5 int32 = 274
    var v6 int16 = 1000
    for fi7 := 0; fi7 < 8; fi7++ {
        v5 -= v5
        p0 -= p0
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1() int32 {
    if ((!false) && (-2.25 != 0.0)) {
        if (1 == 234) {
            var v8 float32 = 1.5
            switch "foo bar" {
            case "hello":
                switch "\n\t" {
                case "foo bar":
                    v9 := []int{-591, genFn0(4294967295, true), 127, genFn0(2, (127 > 25))}
                    _ = v9
                }
            case "foo bar":
                v8 -= v8
            case "\n\t":
                var v10 bool = false
                _ = v10
            default:
                v8 = 0.1
            }
            _ = v8
        } else {
            var v11 float64 = 0.0
            _ = v11
        }
        var v12 uint = 910
        var v13 uint16 = 32767
        _ = v12
        _ = v13
    }
    var v14 float32 = (-2.25 / 1.5)
    _ = v14
    return 42
}

func main() {
    var v1 bool = (!(128 < 59))
    var v2 int8 = 0
    v3 := make(map[string]bool)
    var v4 int = 9223372036854775807
    if false {
        var v15 float64 = (0.0 + 0.1)
        _ = v15
    }
    v1 = v1
    switch "\n\t" {
    case "hello":
        v16 := GenS1{F0: 65535, F1: 4294967295, F2: 255}
        _ = v16
    default:
        v17 := []int{(v4 << ((1) & 63)), (v4 >> ((42) & 63)), (-(v4)), genFn0(428, (v1 || (202 != 127)))}
        _ = v17
    }
    if v1 {
        switch v4 {
        case 128:
            v18 := []int{}
            _ = v18
        case -847:
            for fi19 := 0; fi19 < 5; fi19++ {
                v4 = fi19
                v20 := GenS2{F0: v2}
                for fi21 := 0; fi21 < 8; fi21++ {
                    var v22 string = "foo bar"
                    var v23 bool = true
                    v24 := v3[v22]
                    switch fi21 {
                    case 982:
                        var v25 bool = v23
                        _ = v25
                    case 113:
                        v24 = v23
                    }
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi19
                _ = v20
            }
        case -845:
            for fi26 := 0; fi26 < 9; fi26++ {
                if (!(127 >= 128)) {
                    var v27 uint64 = 18446744073709551615
                    switch fi26 {
                    case 42:
                        delete(v3, "a")
                    case -9223372036854775808:
                        _ = len(v3)
                    case 9223372036854775807:
                        for fi28 := 0; fi28 < 4; fi28++ {
                            if (0.1 < float64((0.1 + ((3.14159 * -2.25) - ((0.1 * (0.1 * -2.25)) / 1.5))))) {
                                if ((^(fi28)) == -1) {
                                    if true {
                                        if ((-(fi26)) <= fi28) {
                                            var v29 int16 = 1000
                                            v30 := []int{(fi26 ^ v4), (fi26 - (fi28 / ((genFn0(1, (146 >= genFn1()))) | 1)))}
                                            _ = v29
                                            _ = v30
                                        }
                                    } else {
                                        switch v4 {
                                        case -351:
                                            v31 := []int{128, -53, -1, genFn0(v27, ((false || v1) || (((0.0 + 1.5) * 3.14159) > 0.0)))}
                                            _ = v31
                                        }
                                        v32 := GenS2{F0: int8(v27)}
                                        if v1 {
                                            v33 := GenS3{F0: (v4 % ((genFn0(v27, (32767 <= 738))) | 1))}
                                            var v34 uint16 = 255
                                            var v35 string = (("" + "") + "")
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = v32
                                    }
                                    var v36 int8 = 1
                                    _ = v36
                                } else {
                                    var v37 bool = v1
                                    _ = v37
                                }
                                if v1 {
                                    if (-1 != 9223372036854775807) {
                                        v4 += (fi28 | fi28)
                                        if v1 {
                                            var v38 uint64 = 2
                                            _ = v38
                                        } else {
                                            if v1 {
                                                fmt.Println(fi28, v27, v4)
                                            } else {
                                                fmt.Println(fi26)
                                                v39 := GenS2{F0: (^(v2))}
                                                v40 := []int{v4, fi28, (-(fi26))}
                                                _ = v39
                                                _ = v40
                                            }
                                            v2 = (v2 + 18)
                                            if ((-9223372036854775808 < 42) || ((v27 << ((0) & 63)) == (v27 & (^(v27))))) {
                                                switch ("" + "a") {
                                                case "":
                                                    switch genFn0(v27, (42 > 1)) {
                                                    case -9223372036854775808:
                                                        switch "foo bar" {
                                                        case "a":
                                                            var v41 float32 = 3.14159
                                                            _ = v41
                                                        }
                                                        var v42 string = ""
                                                        v43 := GenS4{F0: 255}
                                                        v2 = 0
                                                        if (!v1) {
                                                            v42 = v42
                                                            var v44 int64 = 337
                                                            _ = v44
                                                        } else {
                                                            v45 := GenS3{F0: -1}
                                                            var v46 int16 = -897
                                                            v2 += (-(v2))
                                                            v47 := GenS5{F0: 1}
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        _ = v42
                                                        _ = v43
                                                    case 2147483647:
                                                        v4 = (v4 * -2147483648)
                                                        v1 = (((81 < 2) && v1) && (v1 && v1))
                                                        v1 = (869 != 0)
                                                    case 953:
                                                        var v48 float64 = 3.14159
                                                        _ = v48
                                                    default:
                                                        var v49 float32 = (0.1 / 1.0)
                                                        _ = v49
                                                    }
                                                }
                                                var v50 float64 = 1.5
                                                _ = v50
                                            } else {
                                                var v51 int64 = 42
                                                fmt.Println(v4, v2, fi26)
                                                _ = v51
                                            }
                                        }
                                        var v52 int64 = 42
                                        var v53 int8 = v2
                                        _ = v52
                                        _ = v53
                                    }
                                } else {
                                    v54 := GenS6{F0: 2147483647, F1: true, F2: ""}
                                    v55 := []int{2147483647, (-(v4)), 31}
                                    var v56 string = "hello"
                                    v27 += v27
                                    _ = v54
                                    _ = v55
                                    _ = v56
                                }
                            } else {
                                v57 := GenS5{F0: v27}
                                switch 2147483647 {
                                case 32768:
                                    v58 := GenS7{F0: -42, F1: (v2 << ((452) & 63)), F2: 57}
                                    _ = v58
                                case 77:
                                    v1 = v1
                                case 2147483647:
                                    var v59 float64 = 0.0
                                    _ = v59
                                }
                                var v60 int32 = 2147483647
                                _ = v57
                                _ = v60
                            }
                            var v61 int8 = -128
                            v62 := GenS3{F0: (fi28 + -9223372036854775808)}
                            v63 := GenS3{F0: (v4 | (fi28 / (((v4 << ((42) & 63))) | 1)))}
                            _ = fi28
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                    }
                    var v64 float32 = 1e10
                    _ = v27
                    _ = v64
                } else {
                    var v65 uint16 = 65535
                    v65 = (v65 - v65)
                    switch "\n\t" {
                    case "hello":
                        var v66 int8 = (v2 * v2)
                        _ = v66
                    }
                    _ = v65
                }
                if (!v1) {
                    var v67 int = genFn0(18446744073709551615, (893 > 0))
                    if (!true) {
                        var v68 int = (v67 + (v67 >> ((73) & 63)))
                        _ = v68
                    }
                    _ = v67
                }
                _ = fi26
            }
        default:
            fmt.Println(v2)
        }
        _ = len(v3)
        fmt.Println(v1, v4)
    }
    var v69 uint8 = 128
    var v70 uint64 = 0
    fmt.Println(v4, v1, v1)
    v70 = v70
    if v1 {
        for fi71 := 0; fi71 < 20; fi71++ {
            fmt.Println(v4, fi71, v4)
            _ = fi71
        }
    } else {
        if v1 {
            var v72 int16 = -787
            _ = v72
        }
        v73 := []int{v4}
        switch v4 {
        case 9223372036854775807:
            if (3.14159 <= (0.1 - (((0.1 / 0.1) - 1e10) - -2.25))) {
                v1 = (0 == -297)
            } else {
                var v74 int16 = 32767
                v69 -= (v69 ^ v69)
                v69 -= v69
                _ = v74
            }
        case 0:
            for fi75 := 0; fi75 < 9; fi75++ {
                switch -125 {
                case -124:
                    for fi76 := 0; fi76 < 12; fi76++ {
                        if (100 >= v2) {
                            var v77 int8 = (v2 + -100)
                            _ = v77
                        }
                        fmt.Println(v4, v70, v2)
                        _ = fi76
                    }
                case 9223372036854775807:
                    if (((((1e10 + (-2.25 / (0.1 * -2.25))) - ((1e10 - 0.0) - 0.0)) >= -2.25) && false) || (-90 != 657)) {
                        v78 := v3[("\n\t" + "")]
                        var v79 bool = false
                        _ = v78
                        _ = v79
                    } else {
                        fmt.Println(v70)
                        var v80 int8 = v2
                        _ = v80
                    }
                case -1:
                    var v81 float32 = (1.5 * ((3.14159 * -2.25) * 1e10))
                    _ = v81
                default:
                    switch ("a" + ("foo bar" + "\n\t")) {
                    case "\n\t":
                        if (v1 && (-329 <= 0)) {
                            switch "a" {
                            case "\n\t":
                                fmt.Println(v1)
                            case "hello":
                                v82 := GenS4{F0: 42}
                                _ = v82
                            default:
                                var v83 uint8 = v69
                                _ = v83
                            }
                        }
                    case "a":
                        var v84 uint16 = 627
                        _ = v84
                    case "":
                        var v85 uint8 = (v69 & (v69 & v69))
                        _ = v85
                    default:
                        if (42 == v70) {
                            switch ("a" + ((("hello" + "") + ("\n\t" + "hello")) + "hello")) {
                            case "\n\t":
                                v1 = false
                            case "a":
                                switch (fi75 | v4) {
                                case -468:
                                    if true {
                                        fmt.Println(v2, v70, v69)
                                        if (127 < v69) {
                                            for fi86 := 0; fi86 < 13; fi86++ {
                                                fmt.Println(v2, v69)
                                                if v1 {
                                                    v3["foo bar"] = false
                                                } else {
                                                    v1 = (0 > 954)
                                                }
                                                if ((!v1) || false) {
                                                    fmt.Println(v69, v1, v69)
                                                    if v1 {
                                                        var v87 uint64 = v70
                                                        var v88 uint32 = 4294967295
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        switch "" {
                                                        case "\n\t":
                                                            if len(v73) > 0 {
                                                                v73[0] = fi86
                                                            }
                                                        case "hello":
                                                            fmt.Println(fi75)
                                                        case "a":
                                                            var v89 bool = (!false)
                                                            _ = v89
                                                        default:
                                                            v90 := GenS7{F0: int32(v2), F1: (v2 & v2), F2: v2}
                                                            _ = v90
                                                        }
                                                        if ((v2 ^ -1) <= v2) {
                                                            v69 = v69
                                                        }
                                                    }
                                                    if len(v73) >= 2 {
                                                        v73 = v73[1:len(v73)]
                                                    }
                                                    v70 -= (v70 ^ v70)
                                                }
                                                _ = fi86
                                            }
                                            fmt.Println(v69)
                                            v91 := GenS8{F0: 42}
                                            _ = v91
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                v92 := GenS1{F0: 574, F1: 255, F2: 255}
                                                _ = v92
                                            case "hello":
                                                v2 += v2
                                            default:
                                                v93 := GenS5{F0: v70}
                                                _ = v93
                                            }
                                        }
                                    } else {
                                        v94 := GenS2{F0: v2}
                                        _ = v94
                                    }
                                case 4294967295:
                                    switch fi75 {
                                    case -571:
                                        if (65535 <= 0) {
                                            var v95 int64 = -789
                                            _ = v95
                                        } else {
                                            var v96 int16 = -32768
                                            v96 = int16(v70)
                                            _ = v96
                                        }
                                    case -948:
                                        v1 = v1
                                    }
                                case 228:
                                    switch "\n\t" {
                                    case "a":
                                        var v97 uint64 = (v70 >> ((65535) & 63))
                                        _ = v97
                                    case "foo bar":
                                        switch genFn0(v70, false) {
                                        case 32767:
                                            v98 := GenS5{F0: v70}
                                            _ = v98
                                        case 4294967295:
                                            var v99 uint32 = 2147483647
                                            _ = v99
                                        default:
                                            switch (fi75 * (fi75 >> ((uint(v69)) & 63))) {
                                            case 42:
                                                var v100 int16 = 1000
                                                _ = v100
                                            }
                                        }
                                    case "hello":
                                        v2 += v2
                                    default:
                                        if (v70 >= 2) {
                                            v3["foo bar"] = (397 > 4294967295)
                                        } else {
                                            var v101 int32 = 492
                                            v3["\n\t"] = ("a" < "")
                                            _ = v101
                                        }
                                    }
                                }
                            }
                        } else {
                            var v102 bool = (!v1)
                            _ = len(v3)
                            _ = v102
                        }
                    }
                }
                _ = fi75
            }
        case 128:
            switch (^(v4)) {
            case 824:
                var v103 int = (v4 - (v4 * (v4 / ((-115) | 1))))
                _ = v103
            }
        default:
            fmt.Println(v70, v2, v2)
        }
        _ = v73
    }
    if (!v1) {
        switch (v4 / (((v4 >> ((2) & 63))) | 1)) {
        case -2147483648:
            var v104 string = ("\n\t" + ("\n\t" + "foo bar"))
            _ = v104
        case 9223372036854775807:
            for fi105 := 0; fi105 < 6; fi105++ {
                v70 = 0
                if v1 {
                    if v1 {
                        var v106 float32 = (0.0 + 3.14159)
                        for fi107 := 0; fi107 < 8; fi107++ {
                            var v108 bool = v1
                            _ = fi107
                            _ = v108
                        }
                        _ = v106
                    }
                    var v109 uint64 = 2
                    _ = v109
                }
                fmt.Println(v2)
                _ = fi105
            }
        case 672:
            for fi110 := 0; fi110 < 2; fi110++ {
                for fi111 := 0; fi111 < 6; fi111++ {
                    var v112 float64 = float64(3.14159)
                    var v113 float32 = 3.14159
                    _ = fi111
                    _ = v112
                    _ = v113
                }
                _ = fi110
            }
        }
        if ((v70 + (v70 * (^(v70)))) != (v70 + (v70 & v70))) {
            v114 := []int{(v4 % ((v4) | 1)), (v4 % ((32768) | 1))}
            fmt.Println(v70, v4)
            _ = v114
        }
    } else {
        v115 := v3["hello"]
        for fi116 := 0; fi116 < 9; fi116++ {
            if ((v69 & v69) >= v69) {
                if false {
                    for fi117 := 0; fi117 < 9; fi117++ {
                        v115 = v1
                        if (0 < (v69 % ((uint8(v70)) | 1))) {
                            v118 := GenS3{F0: fi116}
                            v119 := GenS1{F0: 18446744073709551615, F1: 418, F2: 42}
                            fmt.Println(fi117, v1)
                            _ = v118
                            _ = v119
                        }
                        v69 = 42
                        _ = fi117
                    }
                }
                var v120 float32 = ((0.1 * 0.0) - float32((3.14159 - float64(((1.5 + -2.25) / 1.0)))))
                var v121 uint32 = 4294967295
                _ = v120
                _ = v121
            } else {
                if (float64(1e10) != 1.5) {
                    for fi122 := 0; fi122 < 1; fi122++ {
                        v123 := []int{fi116}
                        var v124 uint16 = 42
                        delete(v3, "a")
                        v123 = append(v123, (fi116 >> ((255) & 63)))
                        v125 := GenS7{F0: -808, F1: v2, F2: -110}
                        var v126 float64 = (1e10 + 0.0)
                        _ = fi122
                        _ = v123
                        _ = v124
                        _ = v125
                        _ = v126
                    }
                }
            }
            if (897 < 487) {
                for fi127 := 0; fi127 < 15; fi127++ {
                    v128 := []int{}
                    v129 := v3[("" + "a")]
                    var v130 uint16 = 89
                    _ = fi127
                    _ = v128
                    _ = v129
                    _ = v130
                }
                for fi131 := 0; fi131 < 12; fi131++ {
                    var v132 uint32 = 2147483647
                    switch (fi131 ^ fi116) {
                    case 0:
                        switch 4294967295 {
                        case -2147483648:
                            v133 := GenS9{F0: 42, F1: genFn0((v70 / ((v70) | 1)), v115)}
                            _ = v133
                        case 978:
                            fmt.Println(v132)
                        }
                    case 171:
                        var v134 uint = 255
                        _ = v134
                    }
                    v4 = 2147483647
                    _ = fi131
                    _ = v132
                }
            }
            v70 = (v70 & (^(v70)))
            _ = fi116
        }
        _ = v115
    }
    switch 128 {
    case 9223372036854775807:
        v4 = -932
    }
    var v135 int8 = v2
    switch "\n\t" {
    case "hello":
        v136 := GenS2{F0: -1}
        _ = v136
    default:
        var v137 uint = 2
        _ = v137
    }
    for fi138 := 0; fi138 < 13; fi138++ {
        var v139 uint = 18446744073709551615
        _ = fi138
        _ = v139
    }
    fmt.Println(v2, v69)
    var v140 int16 = 1000
    v70 += (v70 - v70)
    var v141 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v140:", v140)
    fmt.Println("v141:", v141)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v135
    _ = v140
    _ = v141
    _ = v2
    _ = v3
    _ = v4
    _ = v69
    _ = v70
}
