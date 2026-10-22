package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int8; F2 uint16 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 int; F1 uint16; F2 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int; F1 int8 }
type GenS10 struct { F0 string; F1 int32; F2 bool }
type GenS11 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS12 struct { F0 uint64; F1 uint8; F2 string }
func genFn0(p0 int16) uint8 {
    var v7 int8 = 1
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{fi9, fi9}
            var v11 int = fi9
            v10 = append(v10, (fi9 / (((fi9 - (fi9 & (fi8 >> ((18446744073709551615) & 63))))) | 1)))
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := []int{(-(fi8)), 4294967295}
        _ = fi8
        _ = v12
    }
    fmt.Println(v7)
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            if (18446744073709551615 < 42) {
                switch (fi14 & 597) {
                case -58:
                    switch fi13 {
                    case 42:
                        if false {
                            fmt.Println(v7, fi13, fi14)
                            if true {
                                var v15 uint64 = 483
                                if false {
                                    var v16 uint16 = 42
                                    _ = v16
                                } else {
                                    switch (-(fi13)) {
                                    case -1:
                                        var v17 uint32 = 2147483647
                                        _ = v17
                                    case 2147483647:
                                        if (false || true) {
                                            var v18 float32 = ((3.14159 - float32(0.0)) * 1e10)
                                            v19 := []int{}
                                            switch "\n\t" {
                                            case "foo bar":
                                                v20 := GenS1{F0: p0}
                                                _ = v20
                                            }
                                            var v21 uint8 = 42
                                            var v22 int = (-(fi14))
                                            _ = v18
                                            _ = v19
                                            _ = v21
                                            _ = v22
                                        }
                                        v15 = (v15 >> ((2) & 63))
                                        v23 := GenS1{F0: (p0 ^ (p0 | p0))}
                                        _ = v23
                                    }
                                    v15 = (v15 * v15)
                                }
                                v24 := []int{(fi13 % ((int(v15)) | 1)), fi14, (fi14 * fi14)}
                                _ = v15
                                _ = v24
                            } else {
                                v7 = (v7 / ((v7) | 1))
                                v25 := GenS1{F0: p0}
                                _ = v25
                            }
                            p0 = (^(p0))
                        } else {
                            p0 += (^(p0))
                            p0 += p0
                            var v26 uint16 = 330
                            _ = v26
                        }
                    }
                case 128:
                    if true {
                        v27 := []int{(fi13 % (((fi13 ^ (fi13 >> ((18446744073709551615) & 63)))) | 1)), (fi13 % ((fi13) | 1)), (fi14 >> ((2) & 63)), fi14}
                        v28 := GenS2{F0: true}
                        _ = v27
                        _ = v28
                    } else {
                        var v29 uint8 = 127
                        var v30 uint16 = 630
                        v30 = 2
                        _ = v29
                        _ = v30
                    }
                default:
                    if true {
                        v31 := []int{(fi14 % ((fi14) | 1)), (fi13 - (fi14 >> ((1) & 63))), fi14}
                        if (0 > 4294967295) {
                            var v32 uint64 = 0
                            var v33 int = (fi14 / ((fi13) | 1))
                            _ = v32
                            _ = v33
                        } else {
                            p0 += p0
                            v7 = (v7 * (-(v7)))
                            v7 += (-(v7))
                        }
                        _ = len(v31)
                        _ = v31
                    } else {
                        v34 := []int{(^(fi13)), (fi14 << ((4294967295) & 63))}
                        switch ("\n\t" + "") {
                        case "foo bar":
                            v7 = int8(p0)
                        case "a":
                            switch fi13 {
                            case -9223372036854775808:
                                var v35 bool = ("a" == "hello")
                                _ = v35
                            case 4294967295:
                                v7 = (-(v7))
                            }
                        }
                        _ = v34
                    }
                }
                if false {
                    p0 = p0
                }
            }
            _ = fi14
        }
        switch fi13 {
        case -1:
            if false {
                var v36 uint32 = 2147483648
                v36 += 2
                _ = v36
            } else {
                for fi37 := 0; fi37 < 13; fi37++ {
                    switch "" {
                    case "":
                        fmt.Println(fi13)
                    case "hello":
                        var v38 uint16 = 32767
                        _ = v38
                    }
                    _ = fi37
                }
            }
        }
        var v39 int16 = p0
        _ = fi13
        _ = v39
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 bool, p1 uint64) uint {
    v40 := []int{9223372036854775807, 675, -2147483648}
    p0 = p0
    var v41 bool = p0
    if len(v40) >= 2 {
        v40 = v40[1:len(v40)]
    }
    _ = p0
    _ = p1
    _ = v40
    _ = v41
    return 65535
}

func main() {
    var v1 string = ""
    var v2 int = -9223372036854775808
    v3 := make(map[string]uint32)
    var v4 uint16 = 0
    var v5 uint64 = 0
    v6 := []int{v2, v2}
    var v42 float32 = 1.5
    if ((true || false) || ((!((false && (!true)) || (v42 >= float32(((3.14159 * (float64((float32((3.14159 + 0.1)) + v42)) / (3.14159 - -2.25))) * -2.25))))) && false)) {
        v43 := GenS3{F0: 1, F1: 1, F2: v4}
        _ = v43
    } else {
        v44 := GenS1{F0: 32767}
        if len(v6) > 0 {
            v6[0] = (^(v2))
        }
        _ = v44
    }
    var v45 bool = ((false && (-1 < 127)) || true)
    fmt.Println(v2, v2, v4)
    var v46 float32 = ((float32((0.0 / 1e10)) * 0.1) * ((((1.5 * v42) - v42) * float32(-2.25)) * (v42 + v42)))
    v5 -= v5
    var v47 float64 = 3.14159
    switch v2 {
    case 1:
        for fi48 := 0; fi48 < 13; fi48++ {
            if ((!v45) || ((0.0 == v47) && (v2 >= (-(v2))))) {
                v46 = float32((v47 + (1.5 / 3.14159)))
                if (-1 < -2147483648) {
                    for fi49 := 0; fi49 < 14; fi49++ {
                        var v50 float64 = (float64(float32(0.1)) * v47)
                        var v51 string = (v1 + (v1 + "\n\t"))
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 5; fi52++ {
                        v2 += v2
                        if (v1 != "hello") {
                            var v53 int = 32768
                            var v54 bool = v45
                            _ = v53
                            _ = v54
                        } else {
                            _ = len(v3)
                            var v55 float32 = v42
                            switch (v1 + (v1 + v1)) {
                            case "a":
                                if v45 {
                                    v42 += float32((((v47 - 3.14159) - v47) * (v47 / float64((1.5 + 1e10)))))
                                    delete(v3, ("\n\t" + v1))
                                    if v45 {
                                        v42 += float32(float64(0.0))
                                        var v56 uint = genFn1(false, v5)
                                        _ = v56
                                    } else {
                                        v2 = fi52
                                    }
                                    var v57 bool = v45
                                    _ = v57
                                }
                            }
                            _ = v55
                        }
                        v58 := GenS1{F0: 42}
                        if v45 {
                            switch (v1 + v1) {
                            case "hello":
                                _ = len(v3)
                            default:
                                v59 := GenS3{F0: int32(v4), F1: 38, F2: v4}
                                _ = v59
                            }
                            var v60 uint32 = 42
                            switch "a" {
                            case "foo bar":
                                fmt.Println(v1)
                            }
                            _ = v60
                        }
                        _ = fi52
                        _ = v58
                    }
                }
                v61 := GenS4{F0: v1, F1: 26}
                _ = v61
            } else {
                _ = len(v3)
                for fi62 := 0; fi62 < 6; fi62++ {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi62
                }
                fmt.Println(v1)
            }
            _ = fi48
        }
    case -694:
        for fi63 := 0; fi63 < 8; fi63++ {
            if v45 {
                for fi64 := 0; fi64 < 2; fi64++ {
                    v3["hello"] = 309
                    delete(v3, v1)
                    _ = fi64
                }
            }
            if false {
                if (v47 == 3.14159) {
                    var v65 uint = 18446744073709551615
                    v46 = (v46 / 1e10)
                    _ = v65
                } else {
                    for fi66 := 0; fi66 < 14; fi66++ {
                        var v67 int = v2
                        if (v1 > ((("foo bar" + "foo bar") + v1) + (v1 + ("\n\t" + "hello")))) {
                            switch (fi63 | v2) {
                            case -2147483648:
                                if (419 >= v5) {
                                    fmt.Println(v45, v45, v67)
                                    if v45 {
                                        _ = len(v3)
                                        var v68 float64 = v47
                                        v1 = v1
                                        _ = v68
                                    }
                                } else {
                                    v69 := GenS5{F0: 0}
                                    fmt.Println(v47, v5, v47)
                                    switch v1 {
                                    case "hello":
                                        if (v1 >= v1) {
                                            var v70 uint8 = 127
                                            var v71 uint16 = (v4 % ((v4) | 1))
                                            _ = v70
                                            _ = v71
                                        }
                                    case "foo bar":
                                        v67 += (-(v2))
                                    default:
                                        var v72 uint32 = 42
                                        _ = v72
                                    }
                                    _ = v69
                                }
                            case 128:
                                switch v1 {
                                case "\n\t":
                                    switch v1 {
                                    case "\n\t":
                                        v73 := GenS6{F0: 2, F1: (true && ((v5 < (v5 | v5)) && v45))}
                                        v74 := v3["a"]
                                        if ((^(v4)) < (v4 << ((964) & 63))) {
                                            switch v67 {
                                            case 0:
                                                v4 = 65535
                                            default:
                                                if v45 {
                                                    v2 = (v2 & fi63)
                                                    var v75 int32 = 42
                                                    _ = v75
                                                } else {
                                                    var v76 uint = 2
                                                    v74 = v74
                                                    _ = v76
                                                }
                                            }
                                            delete(v3, (("a" + "hello") + v1))
                                            if (-128 < 0) {
                                                switch (v1 + "foo bar") {
                                                case "\n\t":
                                                    var v77 float64 = v47
                                                    _ = v77
                                                case "":
                                                    if v45 {
                                                        switch v1 {
                                                        case "":
                                                            v74 -= 811
                                                        case "a":
                                                            var v78 int16 = 0
                                                            _ = v78
                                                        }
                                                        var v79 float32 = v42
                                                        var v80 string = v1
                                                        _ = v79
                                                        _ = v80
                                                    } else {
                                                        var v81 uint16 = 256
                                                        v82 := GenS2{F0: (!v45)}
                                                        v83 := v3["hello"]
                                                        _ = len(v6)
                                                        var v84 uint = 0
                                                        _ = v81
                                                        _ = v82
                                                        _ = v83
                                                        _ = v84
                                                    }
                                                case "hello":
                                                    if v45 {
                                                        v67 -= 0
                                                        v85 := GenS7{F0: (fi63 & (fi63 - 971)), F1: 42, F2: uint32(v5)}
                                                        switch fi66 {
                                                        case 127:
                                                            fmt.Println(v67)
                                                        case -416:
                                                            if v45 {
                                                                v42 = 0.1
                                                                v6 = append(v6, -1)
                                                                if (1 != -381) {
                                                                    v5 = v5
                                                                }
                                                            } else {
                                                                v86 := GenS8{F0: v5}
                                                                _ = v86
                                                            }
                                                        case 2147483647:
                                                            fmt.Println(v2, v2)
                                                        default:
                                                            v42 += v42
                                                        }
                                                        _ = v85
                                                    } else {
                                                        if ((v5 > 1) && v45) {
                                                            var v87 int32 = 1
                                                            var v88 uint8 = genFn0(42)
                                                            switch ("" + v1) {
                                                            case "\n\t":
                                                                fmt.Println(v88, v46)
                                                            case "a":
                                                                v89 := GenS2{F0: true}
                                                                _ = v89
                                                            case "":
                                                                var v90 uint8 = v88
                                                                _ = v90
                                                            default:
                                                                fmt.Println(v4, v42)
                                                            }
                                                            _ = v87
                                                            _ = v88
                                                        } else {
                                                            if v45 {
                                                                v5 = (v5 ^ v5)
                                                                v42 += v46
                                                            } else {
                                                                var v91 uint = 18446744073709551615
                                                                _ = v91
                                                            }
                                                            v47 = -2.25
                                                            switch fi63 {
                                                            case -191:
                                                                if (835 != 1) {
                                                                    var v92 int = (^(fi66))
                                                                    var v93 int64 = 1
                                                                    _ = v92
                                                                    _ = v93
                                                                }
                                                            default:
                                                                switch v1 {
                                                                case "\n\t":
                                                                    switch ("" + ("" + v1)) {
                                                                    case "foo bar":
                                                                        if v45 {
                                                                            switch -325 {
                                                                            case 0:
                                                                                v47 = (v47 + 0.0)
                                                                            }
                                                                            v94 := GenS5{F0: 598}
                                                                            _ = v94
                                                                        } else {
                                                                            var v95 uint32 = (v74 >> ((167) & 63))
                                                                            fmt.Println(v1)
                                                                            _ = v95
                                                                        }
                                                                    case "hello":
                                                                        v1 = v1
                                                                    case "":
                                                                        var v96 int16 = 0
                                                                        _ = v96
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        if ("a" <= v1) {
                                                            v6 = append(v6, (fi63 * v67))
                                                            var v97 bool = (((v67 & (-(v67))) > 32767) || (v45 && v45))
                                                            v6 = append(v6, v2)
                                                            _ = v97
                                                        }
                                                    }
                                                }
                                                var v98 bool = (!(-2.25 != v47))
                                                if len(v6) > 0 {
                                                    v6[0] = 32767
                                                }
                                                var v99 uint8 = 42
                                                v45 = v98
                                                var v100 int8 = 53
                                                _ = v100
                                                _ = v98
                                                _ = v99
                                            } else {
                                                var v101 uint = 65535
                                                var v102 uint64 = 0
                                                _ = v101
                                                _ = v102
                                            }
                                            var v103 bool = false
                                            fmt.Println(fi66)
                                            v104 := GenS6{F0: 691, F1: v45}
                                            _ = v103
                                            _ = v104
                                        }
                                        _ = v73
                                        _ = v74
                                    case "foo bar":
                                        if v45 {
                                            if (!v45) {
                                                var v105 uint32 = 362
                                                var v106 uint32 = v105
                                                v106 -= (^(v106))
                                                _ = v105
                                                _ = v106
                                            }
                                        } else {
                                            var v107 float32 = v42
                                            var v108 float64 = (v47 / 1.5)
                                            _ = v107
                                            _ = v108
                                        }
                                    case "":
                                        fmt.Println(v47, v46)
                                        v46 = 1e10
                                    default:
                                        var v109 int16 = -105
                                        _ = v109
                                    }
                                case "foo bar":
                                    v47 += (1.5 * ((v47 / v47) + v47))
                                case "hello":
                                    v67 = (fi63 | 1)
                                default:
                                    v110 := GenS9{F0: (v2 * fi63), F1: -1}
                                    _ = v110
                                }
                            }
                        } else {
                            if v45 {
                                v2 += fi66
                            }
                            v45 = v45
                            v46 += (float32(v47) * (v42 - (v46 - 1e10)))
                        }
                        _ = fi66
                        _ = v67
                    }
                    if v45 {
                        v6 = append(v6, (^(fi63)))
                        var v111 int64 = -795
                        fmt.Println(v42, v2, v5)
                        _ = v111
                    }
                    switch v1 {
                    case "\n\t":
                        for fi112 := 0; fi112 < 4; fi112++ {
                            if (!v45) {
                                fmt.Println(v5, v2, v47)
                                _ = len(v3)
                            } else {
                                if (1 < -321) {
                                    v2 = (fi112 >> ((651) & 63))
                                    if v45 {
                                        if (false || true) {
                                            var v113 uint8 = genFn0(-1)
                                            if len(v6) > 0 {
                                                v6[0] = -947
                                            }
                                            _ = v113
                                        }
                                        v114 := v3[(v1 + v1)]
                                        var v115 int8 = -5
                                        _ = v114
                                        _ = v115
                                    }
                                    v116 := GenS1{F0: 32767}
                                    _ = v116
                                } else {
                                    if len(v6) > 0 {
                                        v6[0] = (fi63 / ((128) | 1))
                                    }
                                }
                                var v117 int16 = -810
                                switch 628 {
                                case -2147483648:
                                    var v118 int16 = -1
                                    _ = v118
                                }
                                _ = v117
                            }
                            var v119 string = ((v1 + ("\n\t" + ("foo bar" + "\n\t"))) + "\n\t")
                            _ = fi112
                            _ = v119
                        }
                    case "hello":
                        for fi120 := 0; fi120 < 19; fi120++ {
                            v121 := GenS10{F0: "\n\t", F1: 1, F2: v45}
                            switch v1 {
                            case "hello":
                                _ = len(v6)
                            case "\n\t":
                                v5 = (^(v5))
                            }
                            var v122 uint8 = 42
                            _ = fi120
                            _ = v121
                            _ = v122
                        }
                    case "a":
                        v46 = 3.14159
                    default:
                        var v123 uint = 2
                        _ = v123
                    }
                }
                v3["a"] = 434
            } else {
                var v124 string = ""
                _ = v124
            }
            _ = fi63
        }
    default:
        fmt.Println(v4)
    }
    for fi125 := 0; fi125 < 19; fi125++ {
        if (4294967295 < 646) {
            for fi126 := 0; fi126 < 1; fi126++ {
                if true {
                    v127 := GenS11{F0: 255, F1: 234, F2: (v4 % ((22) | 1))}
                    v128 := GenS2{F0: true}
                    v4 = 32767
                    _ = v127
                    _ = v128
                } else {
                    v129 := GenS12{F0: 4294967295, F1: 232, F2: ((v1 + v1) + "\n\t")}
                    v130 := v3[("hello" + (((v1 + ((v1 + "foo bar") + ((v1 + v1) + v1))) + v1) + "a"))]
                    _ = v129
                    _ = v130
                }
                v46 += ((0.0 + 0.0) - v42)
                var v131 int16 = -879
                _ = fi126
                _ = v131
            }
            var v132 int64 = -1
            _ = v132
        } else {
            fmt.Println(v4)
            v1 = ("\n\t" + "foo bar")
            var v133 int32 = 42
            _ = v133
        }
        var v134 uint32 = 2147483648
        if (v2 <= v2) {
            fmt.Println(v4, v46)
            v1 = v1
        }
        _ = fi125
        _ = v134
    }
    fmt.Println(v4, v4)
    fmt.Println(v5)
    var v135 float32 = (float32(v47) / (((v46 / 1.5) + (1.5 / (3.14159 - (1e10 - (v42 / v46))))) - (float32(0.1) + ((float32(float64(v42)) / v46) * (0.0 / 1e10)))))
    var v136 uint = 988
    v137 := GenS8{F0: v5}
    fmt.Println(v46)
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v135
    _ = v136
    _ = v137
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
