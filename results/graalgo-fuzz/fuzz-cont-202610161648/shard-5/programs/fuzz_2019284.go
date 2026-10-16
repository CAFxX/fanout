package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int32; F1 uint8; F2 string }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 bool; F1 int8; F2 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint16; F1 int16; F2 int16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint8; F1 string }
func genFn0() uint8 {
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 int32 = -980
        var v7 int32 = (v6 % (((v6 - v6)) | 1))
        _ = fi5
        _ = v6
        _ = v7
    }
    v8 := []int{32767}
    var v9 string = "foo bar"
    v8 = append(v8, 127)
    _ = v8
    _ = v9
    return 215
}
func genFn1(p0 int16, p1 uint8) uint8 {
    for fi10 := 0; fi10 < 17; fi10++ {
        p0 += p0
        p1 -= p1
        v11 := GenS1{F0: (fi10 + (fi10 - fi10)), F1: fi10, F2: 417}
        _ = fi10
        _ = v11
    }
    fmt.Println(p1, p1)
    switch (("\n\t" + ("a" + ("\n\t" + "\n\t"))) + (("" + ("foo bar" + (("foo bar" + "foo bar") + "\n\t"))) + ("hello" + "hello"))) {
    case "hello":
        var v12 uint32 = 4294967295
        _ = v12
    case "foo bar":
        if (285 > 627) {
            for fi13 := 0; fi13 < 4; fi13++ {
                for fi14 := 0; fi14 < 16; fi14++ {
                    fmt.Println(p0, p0, p1)
                    var v15 int8 = 127
                    _ = fi14
                    _ = v15
                }
                fmt.Println(p0)
                _ = fi13
            }
            v16 := []int{2147483647, 0, 32768}
            p1 = (p1 - (p1 >> ((4294967295) & 63)))
            _ = v16
        } else {
            v17 := []int{9223372036854775807}
            if false {
                var v18 int32 = 252
                _ = len(v17)
                _ = v18
            } else {
                _ = len(v17)
            }
            _ = v17
        }
    case "a":
        p1 = 0
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 int32 = 0
    var v2 bool = (2 != 2)
    var v3 uint64 = 0
    v4 := make(map[string]int)
    delete(v4, "\n\t")
    var v19 int = -1
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 4; fi21++ {
            v2 = false
            if true {
                v22 := GenS2{F0: 0}
                _ = v22
            } else {
                var v23 uint8 = 2
                _ = v23
            }
            var v24 int8 = 0
            _ = fi21
            _ = v24
        }
        _ = fi20
    }
    fmt.Println(v2)
    var v25 uint16 = 71
    fmt.Println(v2)
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 string = "a"
        if v2 {
            v1 -= -1
        }
        for fi28 := 0; fi28 < 17; fi28++ {
            if v2 {
                if true {
                    var v29 int32 = (-(v1))
                    v30 := []int{32768, fi28, (fi26 - (-(v19))), (v19 * (fi28 % ((fi28) | 1)))}
                    v31 := GenS1{F0: v19, F1: (v19 << ((825) & 63)), F2: 65535}
                    _ = v29
                    _ = v30
                    _ = v31
                }
                switch v27 {
                case "\n\t":
                    if (float32((1.5 + 0.1)) == (1.5 + 3.14159)) {
                        if false {
                            var v32 int32 = 0
                            switch (v19 << ((404) & 63)) {
                            case -2147483648:
                                fmt.Println(v19)
                            default:
                                v32 += (v32 / (((v32 * (v1 - v1))) | 1))
                            }
                            v33 := []int{fi26, (-(fi28))}
                            v3 = v3
                            _ = v32
                            _ = v33
                        } else {
                            v34 := []int{(fi26 | v19)}
                            if (v2 && (-1 < (v1 % ((v1) | 1)))) {
                                v27 = v27
                                v35 := GenS1{F0: fi26, F1: -2147483648, F2: (v25 + (v25 << ((738) & 63)))}
                                _ = len(v4)
                                var v36 int16 = 781
                                _ = v35
                                _ = v36
                            } else {
                                var v37 uint64 = v3
                                _ = v37
                            }
                            var v38 uint8 = 171
                            v3 += 18446744073709551615
                            var v39 int64 = 688
                            if (0.0 == (0.1 / 1.0)) {
                                var v40 uint8 = (v38 & v38)
                                if v2 {
                                    fmt.Println(fi28)
                                    v41 := GenS3{F0: int16(v39), F1: (v40 & (v38 - 42))}
                                    v42 := GenS3{F0: -32768, F1: v38}
                                    _ = v41
                                    _ = v42
                                }
                                fmt.Println(fi28)
                                v43 := GenS1{F0: fi26, F1: (fi28 / ((fi26) | 1)), F2: 49}
                                if (!(!false)) {
                                    var v44 int64 = -9223372036854775808
                                    v25 = v25
                                    _ = v44
                                } else {
                                    var v45 uint = 497
                                    _ = v45
                                }
                                _ = v40
                                _ = v43
                            }
                            var v46 int64 = (v39 + (^(v39)))
                            _ = v34
                            _ = v38
                            _ = v39
                            _ = v46
                        }
                    } else {
                        switch (((v27 + (v27 + "foo bar")) + ("foo bar" + "a")) + "\n\t") {
                        case "":
                            var v47 string = ("hello" + v27)
                            _ = v47
                        case "foo bar":
                            if v2 {
                                switch (v19 + fi26) {
                                case 42:
                                    if v2 {
                                        var v48 float64 = float64(((1e10 + (1e10 / 1.0)) * 3.14159))
                                        v19 -= (v19 >> ((18446744073709551615) & 63))
                                        _ = v48
                                    } else {
                                        var v49 uint = 2
                                        _ = v49
                                    }
                                case -9223372036854775808:
                                    v50 := GenS1{F0: (fi26 ^ (fi26 | v19)), F1: (v19 << ((0) & 63)), F2: (^(v25))}
                                    _ = v50
                                default:
                                    fmt.Println(v19, v3)
                                }
                                var v51 int16 = 868
                                _ = v51
                            }
                        }
                        switch (v19 | -9223372036854775808) {
                        case -753:
                            v25 -= v25
                        case -1:
                            v52 := []int{(fi26 & (fi28 << ((645) & 63)))}
                            _ = v52
                        case -9223372036854775808:
                            var v53 int16 = 32767
                            _ = v53
                        }
                        v25 = v25
                    }
                case "foo bar":
                    var v54 uint64 = (v3 << ((1) & 63))
                    _ = v54
                }
                v55 := GenS4{F0: -2147483648, F1: ("\n\t" + v27)}
                v19 -= 408
                v56 := GenS4{F0: -2147483648, F1: ("a" + (v27 + ((v27 + v27) + "")))}
                _ = v55
                _ = v56
            } else {
                if v2 {
                    v57 := GenS2{F0: -128}
                    v58 := GenS5{F0: v1, F1: 75, F2: v27}
                    var v59 string = (v27 + v27)
                    _ = v57
                    _ = v58
                    _ = v59
                } else {
                    v3 = 1
                    var v60 int16 = 42
                    var v61 float32 = (1.5 - 3.14159)
                    v19 = 4294967295
                    v62 := GenS6{F0: v27}
                    _ = v60
                    _ = v61
                    _ = v62
                }
                if v2 {
                    fmt.Println(v19)
                    switch v27 {
                    case "foo bar":
                        if (!true) {
                            switch (fi28 >> ((472) & 63)) {
                            case -981:
                                switch "a" {
                                case "\n\t":
                                    switch v19 {
                                    case -646:
                                        v1 += -2147483648
                                    case 42:
                                        var v63 string = "hello"
                                        _ = v63
                                    default:
                                        if (127 == 99) {
                                            var v64 bool = (v1 >= (v1 << ((0) & 63)))
                                            v3 -= (v3 ^ (^(v3)))
                                            if false {
                                                switch v19 {
                                                case 1:
                                                    if v2 {
                                                        var v65 uint32 = uint32(v3)
                                                        v65 = v65
                                                        v66 := GenS3{F0: 42, F1: 255}
                                                        _ = v65
                                                        _ = v66
                                                    } else {
                                                        v4[v27] = (v19 ^ fi28)
                                                    }
                                                }
                                            } else {
                                                fmt.Println(v1, fi26, v27)
                                                switch v27 {
                                                case "a":
                                                    var v67 uint64 = (v3 ^ 42)
                                                    _ = v67
                                                }
                                                v68 := GenS1{F0: (v19 / ((int(v3)) | 1)), F1: (fi26 ^ (fi26 ^ v19)), F2: 255}
                                                var v69 uint16 = v25
                                                _ = v68
                                                _ = v69
                                            }
                                            _ = v64
                                        }
                                    }
                                default:
                                    switch v19 {
                                    case -1:
                                        var v70 bool = (v2 || (((v3 / (((^(v3))) | 1)) >= v3) || (9223372036854775807 != 1)))
                                        fmt.Println(v19, v27, v19)
                                        _ = v70
                                    case 9223372036854775807:
                                        fmt.Println(v1)
                                    default:
                                        v71 := []int{(v19 & fi26), (^(v19))}
                                        var v72 float64 = (1e10 / 1.0)
                                        if (3.14159 < 1e10) {
                                            v4[v27] = 32768
                                            switch v27 {
                                            case "a":
                                                v73 := v4[v27]
                                                _ = v73
                                            default:
                                                v27 = v27
                                            }
                                        }
                                        _ = v71
                                        _ = v72
                                    }
                                }
                            case 4294967295:
                                v74 := v4["\n\t"]
                                _ = v74
                            case 32768:
                                v19 -= fi26
                            }
                        } else {
                            if (!true) {
                                if (162 == genFn0()) {
                                    fmt.Println(fi28, v2)
                                    var v75 uint8 = uint8(v25)
                                    _ = v75
                                } else {
                                    fmt.Println(v27, v19, v3)
                                    var v76 int = (v19 << ((0) & 63))
                                    if v2 {
                                        v77 := GenS7{F0: (v1 * -2147483648)}
                                        v19 = (v76 >> ((2) & 63))
                                        _ = v77
                                    }
                                    _ = v76
                                }
                                var v78 uint16 = (v25 % ((0) | 1))
                                fmt.Println(v27, v78)
                                _ = v78
                            } else {
                                v79 := GenS6{F0: (("\n\t" + v27) + v27)}
                                _ = v79
                            }
                        }
                    case "hello":
                        v4[v27] = fi26
                    case "\n\t":
                        fmt.Println(fi26, fi26, v3)
                    }
                    if (false && (false || (!(0 > 970)))) {
                        if false {
                            var v80 int16 = 0
                            var v81 int = fi28
                            v3 = v3
                            var v82 int64 = int64(v80)
                            _ = v80
                            _ = v81
                            _ = v82
                        }
                        switch (fi26 % ((32767) | 1)) {
                        case 881:
                            if v2 {
                                v83 := GenS7{F0: v1}
                                var v84 uint = 65535
                                _ = v83
                                _ = v84
                            }
                        case 127:
                            if (uint32(v3) <= 137) {
                                var v85 uint32 = 164
                                v27 = v27
                                v27 = ((v27 + (v27 + v27)) + v27)
                                _ = v85
                            }
                        case -9223372036854775808:
                            switch (v19 / ((9223372036854775807) | 1)) {
                            case -9223372036854775808:
                                v4[v27] = v19
                            case 128:
                                if false {
                                    if ((v1 - v1) <= (^(v1))) {
                                        v86 := GenS2{F0: 1}
                                        v1 = int32(v25)
                                        if v2 {
                                            if false {
                                                v27 = ""
                                                if ("a" < v27) {
                                                    v3 = (v3 & (v3 | (v3 & (v3 ^ (v3 - v3)))))
                                                    if ((v27 != "a") && v2) {
                                                        var v87 int32 = (^(v1))
                                                        v88 := GenS8{F0: v2, F1: 1, F2: (v25 & v25)}
                                                        switch 4294967295 {
                                                        case 4294967295:
                                                            v19 -= fi28
                                                        default:
                                                            v4[v27] = 0
                                                        }
                                                        fmt.Println(v27, fi28)
                                                        if (uint(v3) != 255) {
                                                            v25 += (v25 - (v25 / (((v25 * 65535)) | 1)))
                                                            if (!true) {
                                                                switch "a" {
                                                                case "hello":
                                                                    v89 := []int{(fi26 / ((42) | 1))}
                                                                    v27 = "hello"
                                                                    _ = v89
                                                                default:
                                                                    fmt.Println(v87, v3)
                                                                }
                                                                var v90 float64 = (0.1 + 1.5)
                                                                _ = v90
                                                            } else {
                                                                if (v2 && true) {
                                                                    var v91 bool = v2
                                                                    v92 := []int{(v19 | v19), 0, fi28, v19}
                                                                    if v91 {
                                                                        var v93 string = v27
                                                                        v94 := v4[(v27 + (v93 + v27))]
                                                                        fmt.Println(v94, v94)
                                                                        v95 := GenS1{F0: (v94 - (fi28 >> ((403) & 63))), F1: 367, F2: 2}
                                                                        v96 := GenS9{F0: -1}
                                                                        _ = v93
                                                                        _ = v94
                                                                        _ = v95
                                                                        _ = v96
                                                                    } else {
                                                                        switch ("" + "") {
                                                                        case "foo bar":
                                                                            switch fi26 {
                                                                            case 0:
                                                                                var v97 bool = (!(-2.25 < 0.1))
                                                                                _ = v97
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v91
                                                                    _ = v92
                                                                }
                                                                var v98 float64 = 0.1
                                                                _ = len(v4)
                                                                _ = v98
                                                            }
                                                        }
                                                        _ = v87
                                                        _ = v88
                                                    } else {
                                                        if (127 > 42) {
                                                            var v99 int = (-(fi26))
                                                            _ = v99
                                                        }
                                                    }
                                                    v4[""] = 972
                                                }
                                                if v2 {
                                                    var v100 int = 32768
                                                    _ = v100
                                                }
                                            }
                                        }
                                        switch -2147483648 {
                                        case 42:
                                            v25 -= 0
                                        }
                                        var v101 int32 = (v1 % ((v1) | 1))
                                        fmt.Println(fi26, v101, v1)
                                        v27 = v27
                                        var v102 int8 = -128
                                        _ = v101
                                        _ = v102
                                        _ = v86
                                    } else {
                                        var v103 int16 = 692
                                        _ = v103
                                    }
                                } else {
                                    v104 := []int{v19}
                                    _ = v104
                                }
                            case 581:
                                if (v2 || v2) {
                                    switch 872 {
                                    case 128:
                                        v2 = (((!v2) && ((0 == -1) || (!true))) || false)
                                    case 32768:
                                        v105 := v4["\n\t"]
                                        _ = v105
                                    case 127:
                                        v106 := GenS4{F0: (v1 ^ 42), F1: ("hello" + (("" + v27) + v27))}
                                        _ = v106
                                    }
                                    v107 := GenS10{F0: (v25 % (((v25 + 1)) | 1)), F1: -1, F2: int16(fi28)}
                                    if v2 {
                                        var v108 int8 = 42
                                        _ = v108
                                    } else {
                                        var v109 uint64 = (v3 / ((v3) | 1))
                                        _ = v109
                                    }
                                    _ = v107
                                } else {
                                    v110 := []int{32767, (v19 ^ (v19 ^ fi26)), 32768, (fi26 + 0)}
                                    _ = v110
                                }
                            }
                        default:
                            var v111 uint = 2
                            _ = v111
                        }
                    } else {
                        v112 := GenS4{F0: v1, F1: v27}
                        v27 = v27
                        _ = v112
                    }
                    v2 = (((-2.25 / 1e10) < -2.25) || (1.5 >= (float32((0.1 / (0.1 / 1.0))) / 0.1)))
                    _ = len(v4)
                    var v113 float32 = 1e10
                    _ = v113
                } else {
                    if true {
                        v4[("a" + v27)] = (-(v19))
                        fmt.Println(fi28, v1, v27)
                        if ((^(v25)) <= (^(v25))) {
                            v27 = v27
                        } else {
                            switch v27 {
                            case "foo bar":
                                v114 := []int{v19}
                                _ = v114
                            }
                        }
                    } else {
                        v115 := GenS11{F0: 112}
                        v2 = false
                        _ = v115
                    }
                }
                v1 = 508
            }
            _ = len(v4)
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    v116 := GenS12{F0: ("foo bar" + "")}
    for fi117 := 0; fi117 < 20; fi117++ {
        v118 := GenS13{F0: 42}
        v19 += (v19 >> ((65535) & 63))
        for fi119 := 0; fi119 < 13; fi119++ {
            v120 := GenS14{F0: genFn0(), F1: ("foo bar" + "\n\t")}
            v1 = v1
            _ = fi119
            _ = v120
        }
        _ = fi117
        _ = v118
    }
    fmt.Println(v2, v19, v19)
    v121 := GenS4{F0: (v1 << ((18446744073709551615) & 63)), F1: "hello"}
    v122 := []int{(-(v19)), 242}
    var v123 float32 = float32((0.0 * -2.25))
    fmt.Println(v123, v2, v3)
    if true {
        var v124 int = 32768
        fmt.Println(v19)
        _ = v124
    }
    v3 += 752
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v122:", len(v122), v122)
    fmt.Println("v123:", v123)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v116
    _ = v121
    _ = v122
    _ = v123
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
