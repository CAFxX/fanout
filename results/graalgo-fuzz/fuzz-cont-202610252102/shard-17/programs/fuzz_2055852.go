package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 int16 }
type GenS4 struct { F0 int; F1 int; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int64; F1 uint64 }
type GenS7 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS8 struct { F0 uint8; F1 int32; F2 string }
type GenS9 struct { F0 int16; F1 uint8; F2 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int16; F1 int }
type GenS14 struct { F0 uint }
func genFn0() uint16 {
    var v6 int = 9223372036854775807
    v7 := []int{v6, (-(v6)), v6, 32767}
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 int) int32 {
    p1 -= 15
    p0 = (p0 << (((^(p0))) & 63))
    _ = p0
    _ = p1
    return int32(p0)
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -128
    v3 := make(map[uint32]string)
    var v4 int64 = 42
    var v5 bool = (42 == 1)
    var v8 uint16 = 65535
    v9 := GenS1{F0: v5, F1: (^(v8)), F2: (v8 % ((v8) | 1))}
    if (v5 || v5) {
        for fi10 := 0; fi10 < 20; fi10++ {
            v5 = ("" > "\n\t")
            if (!v5) {
                v1[(-(fi10))] = "\n\t"
            }
            _ = fi10
        }
    } else {
        var v11 bool = (-1 > (v2 + (-(v2))))
        v12 := GenS1{F0: false, F1: (v8 % (((v8 & 946)) | 1)), F2: (v8 << ((724) & 63))}
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        switch (fi13 & -1) {
        case 32768:
            for fi14 := 0; fi14 < 20; fi14++ {
                fmt.Println(v2, v5, fi13)
                switch (("a" + ("\n\t" + ("\n\t" + "a"))) + "foo bar") {
                case "a":
                    if (("hello" + ("\n\t" + "a")) > "a") {
                        var v15 int64 = (v4 << ((1) & 63))
                        if v5 {
                            switch (fi13 / (((fi14 << ((802) & 63))) | 1)) {
                            case -286:
                                v5 = false
                            }
                            switch ("a" + "hello") {
                            case "":
                                switch fi13 {
                                case 32768:
                                    var v16 int8 = (v2 & (v2 & v2))
                                    _ = v16
                                case -2147483648:
                                    v17 := []int{}
                                    _ = v17
                                case 42:
                                    var v18 int16 = 535
                                    _ = v18
                                default:
                                    var v19 uint64 = 42
                                    _ = v19
                                }
                            case "a":
                                switch ("a" + "hello") {
                                case "":
                                    v20 := GenS1{F0: v5, F1: 42, F2: 255}
                                    _ = v20
                                case "foo bar":
                                    var v21 int16 = 1000
                                    _ = v21
                                case "a":
                                    var v22 float64 = 1.5
                                    _ = v22
                                default:
                                    var v23 string = ("" + "")
                                    _ = v23
                                }
                            case "hello":
                                var v24 string = ""
                                _ = v24
                            default:
                                fmt.Println(fi13, v8, v4)
                            }
                            v3[2] = "\n\t"
                        } else {
                            v25 := GenS2{F0: 914, F1: 766}
                            fmt.Println(v8, v8)
                            v5 = (v2 == -1)
                            _ = v25
                        }
                        _ = v15
                    }
                case "":
                    var v26 string = "hello"
                    _ = v26
                default:
                    if (2 > 738) {
                        switch (fi13 >> ((332) & 63)) {
                        case -331:
                            switch fi13 {
                            case 32767:
                                if (!v5) {
                                    var v27 uint32 = 215
                                    var v28 int32 = -235
                                    _ = v27
                                    _ = v28
                                } else {
                                    fmt.Println(fi14, fi13, v2)
                                }
                            case 17:
                                fmt.Println(v2, v2, fi14)
                            default:
                                v8 += genFn0()
                            }
                        default:
                            switch (-(fi14)) {
                            case -9223372036854775808:
                                fmt.Println(v2, fi14, fi13)
                            case 32767:
                                v2 += 1
                            case -1:
                                fmt.Println(v5, v2, fi14)
                            }
                        }
                        v8 += v8
                        switch "" {
                        case "hello":
                            if false {
                                v29 := []int{}
                                v8 += v8
                                _ = v29
                            }
                        }
                    }
                }
                var v30 uint32 = 1
                _ = fi14
                _ = v30
            }
        case -9223372036854775808:
            for fi31 := 0; fi31 < 11; fi31++ {
                if (v5 && v5) {
                    if ((v4 - (v4 << ((255) & 63))) >= v4) {
                        v32 := GenS2{F0: 456, F1: 0}
                        fmt.Println(v8, v8, fi13)
                        v33 := GenS2{F0: 664, F1: 620}
                        if (-2.25 >= 0.0) {
                            fmt.Println(v4, v4, fi13)
                            fmt.Println(v4, v8)
                        } else {
                            if ((1.5 == -2.25) && (0 > 1)) {
                                v2 += (v2 ^ v2)
                                v8 = 255
                            } else {
                                if (!(!(988 < 2147483647))) {
                                    v2 = (v2 ^ (v2 / ((v2) | 1)))
                                    v34 := GenS3{F0: (fi31 % ((-1) | 1)), F1: v8, F2: -875}
                                    v5 = v5
                                    _ = v34
                                } else {
                                    v35 := GenS4{F0: 2147483647, F1: fi13, F2: 687}
                                    v36 := GenS3{F0: 32768, F1: (v8 >> ((1) & 63)), F2: -999}
                                    _ = v35
                                    _ = v36
                                }
                                var v37 uint8 = 80
                                _ = v37
                            }
                            if v5 {
                                switch -1 {
                                case 2147483647:
                                    v2 = v2
                                case 9223372036854775807:
                                    var v38 float32 = 0.1
                                    _ = v38
                                case -2147483648:
                                    switch "a" {
                                    case "hello":
                                        switch fi13 {
                                        case -2147483648:
                                            var v39 int64 = int64(v8)
                                            _ = v39
                                        case 127:
                                            v40 := GenS2{F0: 18446744073709551615, F1: 65535}
                                            _ = v40
                                        case 9223372036854775807:
                                            v5 = true
                                        }
                                    case "a":
                                        v8 = (v8 & genFn0())
                                    case "\n\t":
                                        v41 := GenS5{F0: -340}
                                        _ = v41
                                    default:
                                        if (104 >= 255) {
                                            v42 := GenS3{F0: (fi31 ^ fi13), F1: v8, F2: 25}
                                            v43 := []int{fi13, -277, fi13}
                                            _ = v42
                                            _ = v43
                                        } else {
                                            v44 := GenS2{F0: 1, F1: 129}
                                            v4 = 0
                                            var v45 uint16 = 65535
                                            _ = v44
                                            _ = v45
                                        }
                                        fmt.Println(v8, v8, v5)
                                        if (0.1 == float32(0.0)) {
                                            v46 := GenS2{F0: 0, F1: 2}
                                            var v47 int64 = v4
                                            switch fi13 {
                                            case -414:
                                                switch (fi13 - (fi31 & fi31)) {
                                                case 127:
                                                    var v48 int32 = 2147483647
                                                    _ = v48
                                                default:
                                                    v47 = (v4 % ((v47) | 1))
                                                }
                                            case 42:
                                                v5 = (42 < (v8 * (^(v8))))
                                            case 32768:
                                                v49 := []int{}
                                                _ = v49
                                            }
                                            _ = v46
                                            _ = v47
                                        }
                                    }
                                }
                                if v5 {
                                    var v50 bool = v5
                                    delete(v1, fi13)
                                    if (((v4 * (v4 | v4)) < v4) && true) {
                                        fmt.Println(v50, v50)
                                        v51 := GenS3{F0: (fi13 % (((fi13 & fi31)) | 1)), F1: v8, F2: 0}
                                        v2 = (v2 << ((42) & 63))
                                        _ = v51
                                    } else {
                                        v52 := GenS5{F0: 1}
                                        fmt.Println(v8, v8)
                                        var v53 float32 = 1e10
                                        _ = v52
                                        _ = v53
                                    }
                                    fmt.Println(v50, v8)
                                    _ = v50
                                } else {
                                    v54 := GenS6{F0: (-(v4)), F1: 18446744073709551615}
                                    v55 := GenS7{F0: 2147483647, F1: 818, F2: -1}
                                    fmt.Println(fi13, fi13, v5)
                                    _ = v54
                                    _ = v55
                                }
                                if v5 {
                                    var v56 int = (^(fi31))
                                    v57 := GenS5{F0: -32768}
                                    _ = v56
                                    _ = v57
                                }
                            } else {
                                v58 := []int{fi13, (fi13 >> ((1) & 63)), (fi13 ^ (fi13 & (fi31 >> ((150) & 63))))}
                                var v59 int32 = 2147483647
                                _ = v58
                                _ = v59
                            }
                        }
                        var v60 float32 = float32(-2.25)
                        var v61 int = fi13
                        switch (fi31 % ((fi13) | 1)) {
                        case 32767:
                            if true {
                                var v62 bool = (!(3.14159 != 3.14159))
                                fmt.Println(fi31, v61, v61)
                                _ = v62
                            } else {
                                v63 := GenS8{F0: 255, F1: 411, F2: "a"}
                                v64 := GenS9{F0: -478, F1: uint8(v8), F2: v2}
                                v61 -= 128
                                v60 = v60
                                var v65 int = (v61 >> ((0) & 63))
                                switch (v65 / ((v61) | 1)) {
                                case 1:
                                    var v66 uint32 = 2147483648
                                    _ = v66
                                default:
                                    var v67 int64 = (v4 | v4)
                                    _ = v67
                                }
                                var v68 uint32 = 201
                                v69 := []int{}
                                var v70 uint8 = 2
                                _ = v63
                                _ = v64
                                _ = v65
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                            var v71 float64 = (1.5 * 3.14159)
                            v60 -= 0.0
                            _ = v71
                        case 2147483647:
                            var v72 uint = 255
                            _ = v72
                        default:
                            v73 := GenS10{F0: 4294967295}
                            _ = v73
                        }
                        fmt.Println(fi13)
                        v61 = fi31
                        _ = v32
                        _ = v33
                        _ = v60
                        _ = v61
                    } else {
                        var v74 int64 = v4
                        delete(v1, fi13)
                        _ = v74
                    }
                    var v75 uint8 = 0
                    v76 := []int{(fi31 / ((fi31) | 1)), fi13, (-(fi31))}
                    v8 = (v8 + 678)
                    switch fi13 {
                    case -2147483648:
                        var v77 string = "foo bar"
                        _ = v77
                    case 32768:
                        var v78 uint32 = 568
                        _ = v78
                    case 102:
                        var v79 uint16 = (^(v8))
                        _ = v79
                    default:
                        switch fi31 {
                        case 646:
                            v5 = ((!v5) || v5)
                        }
                    }
                    var v80 int64 = (v4 / ((689) | 1))
                    if len(v76) > 0 {
                        v76[0] = fi31
                    }
                    _ = v75
                    _ = v76
                    _ = v80
                }
                _ = fi31
            }
        case -940:
            v81 := GenS7{F0: 42, F1: 2, F2: 32767}
            _ = v81
        }
        v5 = (1 > 42)
        var v82 int16 = 314
        _ = fi13
        _ = v82
    }
    fmt.Println(v2, v5)
    if true {
        if true {
            if false {
                switch 4294967295 {
                case 820:
                    fmt.Println(v4, v2, v8)
                }
                for fi83 := 0; fi83 < 2; fi83++ {
                    if v5 {
                        if (992 >= v8) {
                            v84 := GenS8{F0: 25, F1: -2147483648, F2: "foo bar"}
                            for fi85 := 0; fi85 < 3; fi85++ {
                                if (-2147483648 == 1) {
                                    v2 = (^(v2))
                                    v86 := []int{9223372036854775807, fi83, (fi85 * fi83)}
                                    v2 = 100
                                    _ = v86
                                }
                                switch fi83 {
                                case 9223372036854775807:
                                    fmt.Println(v5, v5)
                                case -9223372036854775808:
                                    var v87 int16 = -32768
                                    _ = v87
                                default:
                                    var v88 uint32 = 4294967295
                                    _ = v88
                                }
                                _ = fi85
                            }
                            if (2147483647 <= 65535) {
                                v89 := GenS7{F0: 0, F1: 910, F2: 1000}
                                var v90 int = (fi83 << ((265) & 63))
                                switch ("\n\t" + "") {
                                case "\n\t":
                                    for fi91 := 0; fi91 < 9; fi91++ {
                                        switch ("" + "hello") {
                                        case "":
                                            var v92 int16 = -32768
                                            _ = v92
                                        case "foo bar":
                                            switch ("a" + "hello") {
                                            case "hello":
                                                fmt.Println(v90, fi91, v5)
                                            case "a":
                                                v8 = 206
                                            default:
                                                v2 += (v2 | (v2 | (v2 - v2)))
                                            }
                                        }
                                        if (42 <= 814) {
                                            v93 := GenS4{F0: fi91, F1: v90, F2: 2147483647}
                                            _ = v93
                                        }
                                        if (255 != 0) {
                                            if ((v4 & (-(v4))) <= v4) {
                                                v90 -= int(v4)
                                                v5 = false
                                                if false {
                                                    var v94 int = (v90 - (v90 / (((fi91 << ((65535) & 63))) | 1)))
                                                    _ = v94
                                                } else {
                                                    var v95 uint = 4294967295
                                                    fmt.Println(fi91, v4)
                                                    _ = v95
                                                }
                                            } else {
                                                v96 := GenS11{F0: v8, F1: 32767, F2: int32(v2)}
                                                if v5 {
                                                    v97 := GenS8{F0: 230, F1: genFn1(2, (fi91 % (((-(fi91))) | 1))), F2: "\n\t"}
                                                    v98 := []int{}
                                                    _ = v97
                                                    _ = v98
                                                }
                                                _ = v96
                                            }
                                        }
                                        _ = fi91
                                    }
                                case "foo bar":
                                    v99 := v1[(fi83 * -1)]
                                    _ = v99
                                }
                                _ = v89
                                _ = v90
                            }
                            _ = v84
                        }
                    } else {
                        switch ("hello" + ("a" + ("foo bar" + "\n\t"))) {
                        case "hello":
                            fmt.Println(v2, v8)
                        default:
                            v100 := GenS9{F0: -565, F1: 128, F2: v2}
                            _ = v100
                        }
                        for fi101 := 0; fi101 < 12; fi101++ {
                            fmt.Println(v2, v2, v5)
                            _ = fi101
                        }
                        var v102 int32 = -1
                        _ = v102
                    }
                    var v103 int64 = v4
                    var v104 string = ("foo bar" + "a")
                    _ = fi83
                    _ = v103
                    _ = v104
                }
            }
        }
        if (!(1 != 32767)) {
            fmt.Println(v2)
            switch 9223372036854775807 {
            case -9223372036854775808:
                v2 += -32
            default:
                v5 = v5
            }
        } else {
            var v105 int64 = -599
            if (!(v5 && (2147483647 <= -2147483648))) {
                var v106 int = -9223372036854775808
                var v107 uint64 = 18446744073709551615
                v108 := GenS12{F0: "foo bar"}
                _ = v106
                _ = v107
                _ = v108
            } else {
                for fi109 := 0; fi109 < 6; fi109++ {
                    v105 -= (v4 | v105)
                    _ = fi109
                }
                if (!(v2 > v2)) {
                    var v110 uint = 4294967295
                    _ = v110
                }
                if ((!v5) && false) {
                    if (529 <= 2) {
                        var v111 uint32 = 1
                        v8 += (v8 << ((65535) & 63))
                        _ = v111
                    } else {
                        var v112 float64 = -2.25
                        v5 = v5
                        _ = v112
                    }
                    var v113 int16 = 1
                    _ = v113
                }
            }
            var v114 int = 9223372036854775807
            _ = v105
            _ = v114
        }
        var v115 uint = 255
        _ = v115
    }
    switch 929 {
    case 42:
        if (((!(!v5)) && v5) || v5) {
            var v116 int16 = 32767
            if v5 {
                var v117 bool = v5
                v117 = (v5 || (!v117))
                var v118 bool = false
                _ = v117
                _ = v118
            }
            _ = v116
        }
    case 32768:
        v119 := GenS13{F0: 32767, F1: 4294967295}
        _ = v119
    }
    var v120 uint = 255
    v121 := GenS14{F0: (v120 + v120)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v120:", v120)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v120
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
