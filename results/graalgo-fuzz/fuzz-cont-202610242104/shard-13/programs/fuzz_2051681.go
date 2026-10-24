package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 bool; F1 string; F2 uint64 }
type GenS8 struct { F0 string; F1 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 bool; F1 uint8 }
type GenS12 struct { F0 int }
type GenS13 struct { F0 uint64; F1 int64 }
type GenS14 struct { F0 uint64; F1 int8 }
type GenS15 struct { F0 int8; F1 uint }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int16; F1 int8; F2 uint }
type GenS19 struct { F0 int32; F1 int32; F2 uint16 }
type GenS20 struct { F0 int32 }
type GenS21 struct { F0 int8; F1 bool; F2 int32 }
type GenS22 struct { F0 uint64; F1 int16 }
type GenS23 struct { F0 int32 }
type GenS24 struct { F0 uint; F1 string }
type GenS25 struct { F0 int64 }
type GenS26 struct { F0 bool; F1 uint16 }
type GenS27 struct { F0 int16 }
type GenS28 struct { F0 int16; F1 int16; F2 int }
func genFn0() string {
    v4 := []int{-2147483648, 42, 42, 32768}
    var v5 int32 = 0
    v6 := GenS1{F0: true, F1: -102, F2: 623}
    fmt.Println(v5)
    _ = v4
    _ = v5
    _ = v6
    return ("a" + ("hello" + ""))
}
func genFn1(p0 int8, p1 int16, p2 bool) uint64 {
    for fi7 := 0; fi7 < 20; fi7++ {
        if (p2 || false) {
            p1 -= (p1 - p1)
            var v8 uint8 = 0
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 uint64 = 18446744073709551615
                _ = fi9
                _ = v10
            }
            _ = v8
        } else {
            if ((("\n\t" + "foo bar") >= "\n\t") && p2) {
                var v11 uint32 = 0
                var v12 float32 = ((-2.25 + 1.5) + 3.14159)
                _ = v11
                _ = v12
            } else {
                fmt.Println(p0, p2, p1)
                switch (fi7 << ((969) & 63)) {
                case 4294967295:
                    fmt.Println(p1, p0)
                case -851:
                    var v13 uint64 = 42
                    _ = v13
                case 0:
                    fmt.Println(p0, p1)
                default:
                    fmt.Println(p0)
                }
                p2 = false
            }
            var v14 float32 = ((0.0 + ((1.5 * 3.14159) / float32(1.5))) * 0.0)
            _ = v14
        }
        fmt.Println(p0, p1)
        _ = fi7
    }
    if false {
        if p2 {
            if (2 < 2) {
                v15 := []int{32768, 9223372036854775807}
                for fi16 := 0; fi16 < 9; fi16++ {
                    p0 = p0
                    _ = fi16
                }
                _ = v15
            } else {
                var v17 uint64 = 607
                var v18 bool = true
                _ = v17
                _ = v18
            }
            var v19 uint = 1
            _ = v19
        }
    } else {
        var v20 int32 = -646
        _ = v20
    }
    var v21 int8 = (p0 & 77)
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 0
}

func main() {
    var v1 bool = false
    var v2 uint64 = 42
    var v3 uint64 = (v2 + v2)
    v22 := GenS1{F0: true, F1: -95, F2: 327}
    if ((-2.25 / ((float64(((3.14159 * 0.1) * float32(1.5))) * -2.25) - ((0.1 * -2.25) + 0.1))) > (0.1 / 1e10)) {
        if v1 {
            fmt.Println(v1, v1)
            v2 += genFn1(95, 1000, (0.1 != ((float64(float32((0.1 * (-2.25 - ((float64(((1.5 / 1.0) * 0.1)) - ((float64((3.14159 + 0.1)) / 1.0) + 0.0)) * 1e10))))) - (0.0 - ((((3.14159 + float64(1.5)) - float64(1e10)) * 1e10) - (0.0 - (1.5 - -2.25))))) * 1.5)))
            v23 := GenS2{F0: 594, F1: (^(v2))}
            _ = v23
        }
        var v24 int64 = 0
        _ = v24
    }
    v2 = 4294967295
    v25 := []int{127, 32767}
    v1 = v1
    for fi26 := 0; fi26 < 19; fi26++ {
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 float32 = (1.5 - (0.1 / 0.1))
            var v29 bool = v1
            var v30 int = fi27
            _ = fi27
            _ = v28
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 1; fi31++ {
            if (2 >= 2) {
                fmt.Println(fi26, v2, v2)
                if v1 {
                    switch "" {
                    case "":
                        if (6 >= 0) {
                            switch fi26 {
                            case 186:
                                v3 += v2
                            }
                            fmt.Println(v3, v1, fi31)
                            if len(v25) > 0 {
                                v25[0] = (-(fi26))
                            }
                        } else {
                            v2 = v2
                            if (42 > 0) {
                                if len(v25) >= 2 {
                                    v25 = v25[1:len(v25)]
                                }
                            } else {
                                if len(v25) >= 2 {
                                    v25 = v25[1:len(v25)]
                                }
                                fmt.Println(v1)
                            }
                            switch "hello" {
                            case "foo bar":
                                if ((v3 * (^(v3))) >= v2) {
                                    if false {
                                        switch fi31 {
                                        case 952:
                                            var v32 uint64 = genFn1(0, -209, (!((v3 / ((v2) | 1)) == (v2 - v3))))
                                            _ = v32
                                        case 826:
                                            v25 = append(v25, (fi26 ^ fi31))
                                            var v33 float32 = (1e10 + (0.1 * ((3.14159 + ((1.5 + 0.1) * 0.1)) / 1e10)))
                                            v33 = 0.0
                                            _ = v33
                                        }
                                        switch fi26 {
                                        case 638:
                                            var v34 string = (("foo bar" + "hello") + ("a" + "hello"))
                                            var v35 int = (-(fi31))
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                    var v36 int16 = 939
                                    fmt.Println(v2)
                                    _ = v36
                                } else {
                                    var v37 int32 = 445
                                    var v38 uint = 65535
                                    if len(v25) > 0 {
                                        v25[0] = (^(fi31))
                                    }
                                    _ = v37
                                    _ = v38
                                }
                                switch "\n\t" {
                                case "":
                                    v39 := GenS3{F0: 1, F1: (true && (((100 >= 10) && v1) || (0 >= 4294967295)))}
                                    _ = v39
                                case "a":
                                    var v40 bool = v1
                                    _ = v40
                                case "hello":
                                    if ("" > "hello") {
                                        v41 := GenS4{F0: 1}
                                        if true {
                                            v2 += 18446744073709551615
                                            var v42 uint8 = 170
                                            if v1 {
                                                v2 += v2
                                                switch "foo bar" {
                                                case "a":
                                                    var v43 uint64 = v2
                                                    _ = v43
                                                case "":
                                                    v42 = v42
                                                default:
                                                    v42 += (v42 / ((v42) | 1))
                                                }
                                            } else {
                                                if true {
                                                    if (287 <= 802) {
                                                        var v44 float64 = (1e10 + float64(float32(0.0)))
                                                        _ = v44
                                                    }
                                                } else {
                                                    var v45 int64 = 9223372036854775807
                                                    v42 -= 2
                                                    v46 := GenS5{F0: v45}
                                                    _ = v45
                                                    _ = v46
                                                }
                                                fmt.Println(v3)
                                                if v1 {
                                                    var v47 uint64 = (v2 & v2)
                                                    var v48 float64 = 3.14159
                                                    v48 += v48
                                                    fmt.Println(fi26)
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    _ = len(v25)
                                                    v49 := GenS6{F0: -104, F1: 0}
                                                    v50 := GenS7{F0: (537 >= 2147483648), F1: "\n\t", F2: v3}
                                                    v25 = append(v25, fi26)
                                                    v1 = (-1 >= -32768)
                                                    _ = v49
                                                    _ = v50
                                                }
                                                switch ("hello" + ((("" + "hello") + "a") + "a")) {
                                                case "hello":
                                                    switch (^(fi31)) {
                                                    case -2147483648:
                                                        v51 := GenS8{F0: ("\n\t" + "\n\t"), F1: ""}
                                                        _ = v51
                                                    case 29:
                                                        if (0 < 1) {
                                                            var v52 int8 = 100
                                                            if (!(1 > 1)) {
                                                                v3 -= v3
                                                                v42 = v42
                                                                v1 = (-9223372036854775808 < 1)
                                                                v53 := GenS7{F0: (255 < 2), F1: "hello", F2: genFn1((^(v52)), -592, (323 == 502))}
                                                                switch fi31 {
                                                                case 2147483647:
                                                                    v1 = (v3 >= (v2 - v2))
                                                                case 754:
                                                                    v3 -= 4294967295
                                                                }
                                                                _ = v53
                                                            } else {
                                                                fmt.Println(v42, v52, v52)
                                                                var v54 int32 = 0
                                                                _ = v54
                                                            }
                                                            _ = v52
                                                        } else {
                                                            var v55 uint16 = 957
                                                            fmt.Println(v2, v42, v42)
                                                            _ = v55
                                                        }
                                                    default:
                                                        switch "a" {
                                                        case "foo bar":
                                                            var v56 int8 = -36
                                                            _ = v56
                                                        case "a":
                                                            fmt.Println(v1)
                                                            fmt.Println(v2, v1)
                                                            var v57 uint32 = 19
                                                            _ = v57
                                                        case "":
                                                            v42 = (v42 << ((18446744073709551615) & 63))
                                                        }
                                                        if (351 <= 907) {
                                                            var v58 uint32 = 1
                                                            var v59 uint16 = 42
                                                            if v1 {
                                                                v60 := GenS2{F0: 105, F1: (v3 - 0)}
                                                                if len(v25) >= 2 {
                                                                    v25 = v25[1:len(v25)]
                                                                }
                                                                if (v59 > 42) {
                                                                    if true {
                                                                        v59 += (v59 >> ((uint(v59)) & 63))
                                                                    }
                                                                }
                                                                _ = v60
                                                            }
                                                            if false {
                                                                v61 := GenS5{F0: -1}
                                                                var v62 int = (fi26 << ((896) & 63))
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                var v63 int64 = 42
                                                                _ = v63
                                                            }
                                                            _ = v58
                                                            _ = v59
                                                        } else {
                                                            v3 = genFn1(-1, -669, true)
                                                        }
                                                    }
                                                case "foo bar":
                                                    var v64 uint8 = v42
                                                    if (-403 > -32768) {
                                                        fmt.Println(v1, v1, fi26)
                                                        v42 += v64
                                                        _ = len(v25)
                                                    } else {
                                                        var v65 float32 = (3.14159 - (float32((1.5 / ((1e10 * float64(3.14159)) / 1.0))) * 1.5))
                                                        fmt.Println(v42, v2, v2)
                                                        _ = v65
                                                    }
                                                    switch 42 {
                                                    case 0:
                                                        var v66 bool = v1
                                                        _ = v66
                                                    }
                                                    _ = v64
                                                default:
                                                    v67 := GenS9{F0: (-1 >= -1)}
                                                    v42 = uint8(v2)
                                                    _ = v67
                                                }
                                                v68 := GenS4{F0: v42}
                                                fmt.Println(v3)
                                                switch "hello" {
                                                case "a":
                                                    v69 := GenS10{F0: v1, F1: false}
                                                    _ = v69
                                                case "hello":
                                                    if false {
                                                        switch fi31 {
                                                        case 48:
                                                            v70 := GenS11{F0: v1, F1: 1}
                                                            var v71 float64 = 1.5
                                                            v1 = v1
                                                            _ = v70
                                                            _ = v71
                                                        case -2147483648:
                                                            v25 = append(v25, -243)
                                                        default:
                                                            if v1 {
                                                                if (v42 >= v42) {
                                                                    v2 = (v2 + genFn1(-1, -346, v1))
                                                                    if (!true) {
                                                                        var v72 int = fi31
                                                                        fmt.Println(fi31)
                                                                        if len(v25) > 0 {
                                                                            v25[0] = v72
                                                                        }
                                                                        _ = v72
                                                                    } else {
                                                                        var v73 uint8 = v42
                                                                        _ = v73
                                                                    }
                                                                    v25 = append(v25, -1)
                                                                } else {
                                                                    switch (("a" + "a") + ("a" + (("a" + "foo bar") + "foo bar"))) {
                                                                    case "hello":
                                                                        if true {
                                                                            if v1 {
                                                                                v74 := GenS4{F0: 31}
                                                                                v75 := GenS12{F0: fi31}
                                                                                _ = v74
                                                                                _ = v75
                                                                            }
                                                                            fmt.Println(v2)
                                                                            v42 = (v42 & uint8(v2))
                                                                            var v76 int64 = -1
                                                                            v77 := GenS8{F0: ("a" + "a"), F1: "a"}
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                    default:
                                                                        v78 := GenS13{F0: (^(v2)), F1: -973}
                                                                        v1 = v1
                                                                        _ = v78
                                                                    }
                                                                }
                                                            } else {
                                                                if (0 == -815) {
                                                                    var v79 int32 = 51
                                                                    v80 := GenS14{F0: (^(v3)), F1: 0}
                                                                    switch 9223372036854775807 {
                                                                    case -535:
                                                                        v25 = append(v25, 1)
                                                                    case -1:
                                                                        var v81 bool = (("hello" + ("foo bar" + "foo bar")) < "a")
                                                                        _ = v81
                                                                    default:
                                                                        var v82 int = int(v42)
                                                                        var v83 uint8 = v42
                                                                        if (false && (0 > 18446744073709551615)) {
                                                                            if len(v25) > 0 {
                                                                                v25[0] = (fi31 - (fi31 * (fi31 + (fi26 * (v82 / ((802) | 1))))))
                                                                            }
                                                                            v2 -= v2
                                                                            fmt.Println(fi31, v3, v1)
                                                                        }
                                                                        _ = v82
                                                                        _ = v83
                                                                    }
                                                                    _ = v79
                                                                    _ = v80
                                                                } else {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            if (!v1) {
                                                                                v84 := GenS15{F0: 77, F1: 337}
                                                                                var v85 float32 = (3.14159 / 1.0)
                                                                                _ = len(v25)
                                                                                var v86 uint32 = 62
                                                                                var v87 uint = 42
                                                                                _ = v84
                                                                                _ = v85
                                                                                _ = v86
                                                                                _ = v87
                                                                            }
                                                                        default:
                                                                            v88 := GenS6{F0: -113, F1: 1}
                                                                            _ = v88
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        fmt.Println(v42, v3, fi31)
                                                        var v89 uint32 = 0
                                                        _ = v89
                                                    } else {
                                                        switch 954 {
                                                        case -1:
                                                            if false {
                                                                if ((0.1 - ((0.1 / 1.5) - 3.14159)) > -2.25) {
                                                                    v90 := GenS16{F0: 2147483647}
                                                                    v1 = true
                                                                    _ = v90
                                                                }
                                                            }
                                                        case 4294967295:
                                                            v91 := GenS6{F0: 42, F1: -9223372036854775808}
                                                            _ = v91
                                                        case 649:
                                                            var v92 int64 = -9223372036854775808
                                                            _ = v92
                                                        }
                                                        fmt.Println(v2, v3)
                                                        var v93 int64 = 42
                                                        _ = v93
                                                    }
                                                case "\n\t":
                                                    v25 = append(v25, (^(fi26)))
                                                default:
                                                    var v94 float64 = ((((1.5 - float64((-2.25 / 1e10))) / ((3.14159 + 0.1) - 1e10)) * 0.1) / float64(0.1))
                                                    _ = v94
                                                }
                                                _ = v68
                                            }
                                            _ = v42
                                        }
                                        _ = len(v25)
                                        _ = v41
                                    } else {
                                        var v95 uint16 = 0
                                        v96 := GenS16{F0: -1}
                                        switch 9223372036854775807 {
                                        case -1:
                                            switch "a" {
                                            case "hello":
                                                v95 += (^(v95))
                                            case "foo bar":
                                                var v97 int8 = -83
                                                _ = v97
                                            }
                                            var v98 uint32 = 65535
                                            _ = v98
                                        case 122:
                                            switch 128 {
                                            case 42:
                                                v99 := GenS17{F0: 100}
                                                _ = v99
                                            case 8:
                                                var v100 uint32 = 4294967295
                                                _ = v100
                                            case -184:
                                                v101 := GenS18{F0: 42, F1: -128, F2: 2}
                                                var v102 int16 = 0
                                                _ = v101
                                                _ = v102
                                            }
                                        default:
                                            v25 = append(v25, (fi31 & fi31))
                                        }
                                        fmt.Println(v3)
                                        fmt.Println(v2, fi31, fi31)
                                        _ = v95
                                        _ = v96
                                    }
                                default:
                                    v3 = (v3 % ((18446744073709551615) | 1))
                                }
                            case "":
                                v3 = (v2 ^ (^(v3)))
                                v1 = false
                                v103 := GenS19{F0: -991, F1: -2147483648, F2: 507}
                                _ = v103
                            }
                        }
                        var v104 uint = 717
                        var v105 uint = (^(v104))
                        _ = v104
                        _ = v105
                    case "\n\t":
                        var v106 float64 = (0.1 * 0.0)
                        _ = v106
                    }
                    v3 -= (v2 + (v2 | v3))
                    v107 := GenS20{F0: 2147483647}
                    _ = v107
                } else {
                    var v108 int16 = 32767
                    _ = v108
                }
                var v109 float64 = 0.1
                var v110 uint64 = 42
                var v111 float32 = float32(((v109 / (v109 + 0.0)) / ((v109 * -2.25) * ((float64(float32(float64(-2.25))) + 3.14159) + 3.14159))))
                _ = v109
                _ = v110
                _ = v111
            }
            fmt.Println(v1, fi26)
            var v112 uint = 1
            _ = fi31
            _ = v112
        }
        _ = fi26
    }
    v113 := GenS21{F0: -128, F1: (42 < 0), F2: -2147483648}
    v25 = append(v25, -1)
    fmt.Println(v3, v2, v3)
    v2 = v2
    if false {
        v25 = append(v25, 0)
    } else {
        for fi114 := 0; fi114 < 9; fi114++ {
            v3 = v2
            for fi115 := 0; fi115 < 7; fi115++ {
                var v116 uint64 = (^(v3))
                switch (fi114 * (fi114 + 0)) {
                case -2147483648:
                    var v117 float32 = 3.14159
                    _ = v117
                default:
                    if v1 {
                        var v118 uint = 255
                        v25 = append(v25, (fi114 % (((fi114 << ((v118) & 63))) | 1)))
                        _ = v118
                    }
                }
                v119 := GenS19{F0: 0, F1: 1, F2: 258}
                _ = fi115
                _ = v116
                _ = v119
            }
            var v120 uint32 = 68
            _ = fi114
            _ = v120
        }
    }
    if false {
        var v121 int64 = -9223372036854775808
        fmt.Println(v2, v3)
        if (true || (!v1)) {
            for fi122 := 0; fi122 < 1; fi122++ {
                for fi123 := 0; fi123 < 13; fi123++ {
                    fmt.Println(v2, v121)
                    v1 = (478 == v3)
                    v1 = (458 <= 1)
                    v124 := GenS4{F0: 255}
                    _ = fi123
                    _ = v124
                }
                _ = fi122
            }
        }
        _ = v121
    } else {
        if (956 != 42) {
            var v125 float32 = float32((0.0 * ((3.14159 + 0.1) + 1e10)))
            _ = v125
        } else {
            var v126 uint = 4294967295
            for fi127 := 0; fi127 < 15; fi127++ {
                v1 = v1
                v128 := GenS4{F0: 2}
                v2 = 304
                _ = fi127
                _ = v128
            }
            _ = v126
        }
        if (v1 && (((((1.5 / -2.25) + -2.25) != float64(0.0)) && v1) || v1)) {
            v3 = (v3 - v2)
            v129 := GenS13{F0: (^(v3)), F1: -591}
            _ = v129
        } else {
            v2 = (v3 ^ (v3 - 0))
        }
        if true {
            var v130 float64 = ((0.1 + float64(0.1)) / (float64(3.14159) / (1e10 * 1e10)))
            _ = v130
        } else {
            fmt.Println(v2, v1, v1)
            var v131 bool = ("a" < "\n\t")
            if (-527 > 0) {
                fmt.Println(v131)
            } else {
                fmt.Println(v1)
            }
            _ = v131
        }
    }
    switch ("hello" + "") {
    case "":
        fmt.Println(v3)
    default:
        if v1 {
            v2 += (v3 & v3)
            for fi132 := 0; fi132 < 18; fi132++ {
                if false {
                    for fi133 := 0; fi133 < 6; fi133++ {
                        v1 = v1
                        if false {
                            v3 -= 123
                            var v134 uint64 = v2
                            switch ("hello" + "\n\t") {
                            case "hello":
                                var v135 uint = 1
                                _ = v135
                            }
                            _ = v134
                        } else {
                            v3 -= (v2 ^ v2)
                        }
                        _ = fi133
                    }
                    switch "foo bar" {
                    case "":
                        if v1 {
                            if v1 {
                                if v1 {
                                    v25 = append(v25, (fi132 & (fi132 | -767)))
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    for fi136 := 0; fi136 < 14; fi136++ {
                                        v1 = false
                                        if v1 {
                                            v137 := GenS22{F0: 918, F1: 1}
                                            var v138 uint32 = 2147483647
                                            _ = v137
                                            _ = v138
                                        }
                                        switch fi132 {
                                        case 4294967295:
                                            v139 := GenS9{F0: v1}
                                            _ = v139
                                        default:
                                            var v140 int64 = -9223372036854775808
                                            _ = v140
                                        }
                                        _ = fi136
                                    }
                                case "":
                                    var v141 int8 = 110
                                    _ = v141
                                default:
                                    for fi142 := 0; fi142 < 19; fi142++ {
                                        if ((1 != -1) && v1) {
                                            if (!v1) {
                                                if v1 {
                                                    fmt.Println(v1, v2)
                                                    _ = len(v25)
                                                    v3 -= (v2 << ((255) & 63))
                                                } else {
                                                    fmt.Println(fi142)
                                                    fmt.Println(v3, v3)
                                                }
                                                var v143 int8 = -128
                                                fmt.Println(fi132, v1, v143)
                                                v144 := GenS16{F0: 1}
                                                _ = v143
                                                _ = v144
                                            }
                                            v145 := GenS19{F0: 2147483647, F1: 42, F2: 42}
                                            if v1 {
                                                v146 := GenS23{F0: -1}
                                                if false {
                                                    if v1 {
                                                        v2 = v3
                                                        switch "hello" {
                                                        case "foo bar":
                                                            var v147 int32 = 1
                                                            _ = v147
                                                        case "\n\t":
                                                            if len(v25) > 0 {
                                                                v25[0] = (fi132 - (-(fi142)))
                                                            }
                                                        default:
                                                            if len(v25) > 0 {
                                                                v25[0] = (fi132 * (fi142 * fi142))
                                                            }
                                                        }
                                                    }
                                                    var v148 uint = 4294967295
                                                    _ = v148
                                                }
                                                v1 = (200 < 216)
                                                _ = v146
                                            }
                                            _ = v145
                                        }
                                        switch ("hello" + "foo bar") {
                                        case "":
                                            fmt.Println(fi142, fi142, v2)
                                        default:
                                            v149 := GenS24{F0: 0, F1: "a"}
                                            _ = v149
                                        }
                                        var v150 string = "hello"
                                        _ = fi142
                                        _ = v150
                                    }
                                }
                                if (true && (187 != 2)) {
                                    var v151 uint64 = v2
                                    for fi152 := 0; fi152 < 11; fi152++ {
                                        v2 -= v3
                                        var v153 uint = 241
                                        _ = fi152
                                        _ = v153
                                    }
                                    var v154 int64 = -347
                                    _ = v151
                                    _ = v154
                                } else {
                                    v2 -= v3
                                    for fi155 := 0; fi155 < 17; fi155++ {
                                        if len(v25) >= 2 {
                                            v25 = v25[1:len(v25)]
                                        }
                                        _ = fi155
                                    }
                                }
                            } else {
                                switch "\n\t" {
                                case "":
                                    v156 := GenS25{F0: -248}
                                    _ = v156
                                case "a":
                                    v157 := GenS7{F0: v1, F1: "", F2: 577}
                                    _ = v157
                                case "hello":
                                    for fi158 := 0; fi158 < 9; fi158++ {
                                        fmt.Println(v3, v2, fi132)
                                        v1 = v1
                                        var v159 float64 = 0.1
                                        _ = fi158
                                        _ = v159
                                    }
                                }
                            }
                            if false {
                                switch "a" {
                                case "\n\t":
                                    var v160 float32 = (1.5 + ((-2.25 * 1e10) * ((1e10 * ((-2.25 - ((((((0.0 / 1.0) - 1e10) * 0.1) * 1.5) + -2.25) + ((-2.25 + (-2.25 * (1.5 - 1e10))) * 3.14159))) + float32(0.0))) + 1.5)))
                                    _ = v160
                                case "foo bar":
                                    v1 = (-436 != 2147483647)
                                default:
                                    if len(v25) > 0 {
                                        v25[0] = 9223372036854775807
                                    }
                                }
                                var v161 uint64 = genFn1(127, 42, (42 != 42))
                                _ = v161
                            }
                            v25 = append(v25, fi132)
                        }
                    }
                } else {
                    _ = len(v25)
                    var v162 uint16 = 42
                    v163 := GenS4{F0: 143}
                    _ = v162
                    _ = v163
                }
                _ = fi132
            }
            var v164 uint32 = 662
            _ = v164
        }
    }
    v25 = append(v25, 32767)
    switch ("\n\t" + (("hello" + (("a" + "") + "")) + (("" + "a") + "a"))) {
    case "":
        for fi165 := 0; fi165 < 16; fi165++ {
            v25 = append(v25, (fi165 + (fi165 >> ((723) & 63))))
            _ = fi165
        }
    case "\n\t":
        var v166 uint16 = 961
        _ = v166
    case "a":
        if v1 {
            switch "a" {
            case "hello":
                for fi167 := 0; fi167 < 7; fi167++ {
                    var v168 float64 = (((0.1 * -2.25) + float64(0.1)) * float64(1.5))
                    _ = fi167
                    _ = v168
                }
            default:
                v169 := GenS26{F0: v1, F1: 255}
                _ = v169
            }
            if (4294967295 < (v3 & (v2 - (v3 * (^(v2)))))) {
                v2 -= (v2 >> ((447) & 63))
            }
            v2 += v2
        } else {
            if ((3.14159 / 1.0) >= float64(1e10)) {
                if (v1 || (!true)) {
                    v3 = 42
                } else {
                    fmt.Println(v2, v2, v1)
                }
                var v170 string = "a"
                if (v2 > v3) {
                    v171 := GenS27{F0: 436}
                    _ = v171
                } else {
                    switch "a" {
                    case "":
                        var v172 int16 = 0
                        _ = v172
                    case "hello":
                        v173 := GenS4{F0: 1}
                        _ = v173
                    case "foo bar":
                        v25 = append(v25, 32767)
                    default:
                        for fi174 := 0; fi174 < 6; fi174++ {
                            _ = len(v25)
                            if (174 <= 1) {
                                var v175 uint16 = 1
                                _ = v175
                            }
                            _ = fi174
                        }
                    }
                }
                _ = v170
            }
            if len(v25) >= 2 {
                v25 = v25[1:len(v25)]
            }
            if (393 == 18446744073709551615) {
                var v176 int32 = -2147483648
                for fi177 := 0; fi177 < 13; fi177++ {
                    v176 = v176
                    for fi178 := 0; fi178 < 18; fi178++ {
                        fmt.Println(v3)
                        v179 := GenS28{F0: 1, F1: 1000, F2: 4294967295}
                        var v180 int64 = 42
                        _ = fi178
                        _ = v179
                        _ = v180
                    }
                    v3 = v3
                    _ = fi177
                }
                _ = v176
            } else {
                if v1 {
                    fmt.Println(v2, v1)
                    if v1 {
                        if len(v25) >= 2 {
                            v25 = v25[1:len(v25)]
                        }
                        fmt.Println(v2, v1)
                        v2 -= 42
                    } else {
                        if v1 {
                            switch "\n\t" {
                            case "a":
                                v3 -= 4294967295
                            case "hello":
                                var v181 bool = true
                                _ = v181
                            }
                        }
                        switch "a" {
                        case "a":
                            switch ("foo bar" + (("foo bar" + "") + "hello")) {
                            case "a":
                                v3 = (v3 << ((18446744073709551615) & 63))
                            }
                        case "hello":
                            fmt.Println(v1, v1)
                        case "":
                            _ = len(v25)
                        default:
                            fmt.Println(v3, v2, v2)
                        }
                    }
                } else {
                    v2 -= genFn1(42, 983, false)
                }
            }
        }
    default:
        var v182 int16 = -653
        _ = v182
    }
    v25 = append(v25, 32768)
    if len(v25) >= 2 {
        v25 = v25[1:len(v25)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v113
    _ = v2
    _ = v22
    _ = v25
    _ = v3
}
