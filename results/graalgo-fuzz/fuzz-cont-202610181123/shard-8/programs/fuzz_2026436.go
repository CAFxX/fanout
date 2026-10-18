package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int32; F1 bool; F2 int8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 uint64; F1 int; F2 bool }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int; F1 uint }
type GenS12 struct { F0 int16; F1 uint64 }
type GenS13 struct { F0 int16; F1 int8 }
type GenS14 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS15 struct { F0 uint64; F1 uint }
type GenS16 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS17 struct { F0 int64; F1 bool; F2 string }
type GenS18 struct { F0 uint; F1 int64 }
type GenS19 struct { F0 uint16; F1 int; F2 int }
type GenS20 struct { F0 int }
type GenS21 struct { F0 int64; F1 uint64 }
func genFn0(p0 uint16, p1 int32, p2 bool) int16 {
    v5 := []int{}
    fmt.Println(p0)
    switch (("foo bar" + ("foo bar" + "a")) + "foo bar") {
    case "\n\t":
        var v6 int = 32767
        _ = v6
    case "":
        var v7 uint64 = 42
        _ = v7
    case "foo bar":
        if false {
            p1 += 42
            v8 := GenS1{F0: p2}
            _ = v8
        }
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v4 uint16 = 65535
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v4 = 1
            switch "hello" {
            case "hello":
                v11 := GenS2{F0: 127}
                _ = v11
            case "foo bar":
                var v12 int32 = 2147483647
                _ = v12
            case "":
                fmt.Println(fi9)
                var v13 int32 = -2147483648
                switch 42 {
                case 4294967295:
                    if (!false) {
                        if ((3.14159 + float64((float32(((float64(-2.25) + 3.14159) - (3.14159 / 1.0))) * v1))) < (1e10 / (((float64(v1) - ((float64((0.1 - (v1 * 0.0))) / (float64(3.14159) + float64(0.0))) * 1.5)) * 0.0) + float64(((3.14159 * (float32(0.0) - float32((3.14159 * 1e10)))) + v1))))) {
                            switch (fi9 / ((513) | 1)) {
                            case -9223372036854775808:
                                var v14 float32 = float32(((0.1 - 0.1) * 0.0))
                                if false {
                                    v15 := GenS3{F0: 2147483648, F1: 2, F2: 42}
                                    fmt.Println(v14, fi10)
                                    _ = v15
                                }
                                switch fi10 {
                                case 127:
                                    v16 := []int{(^(fi10))}
                                    _ = v16
                                case 43:
                                    v17 := []int{fi10, fi9, (-(fi9)), 1}
                                    _ = v17
                                }
                                _ = v14
                            case 127:
                                v13 -= (v13 - v13)
                            default:
                                switch (fi10 + fi10) {
                                case 0:
                                    if ((true || true) && true) {
                                        switch "hello" {
                                        case "\n\t":
                                            var v18 uint64 = 2
                                            _ = v18
                                        case "foo bar":
                                            if true {
                                                var v19 float64 = 1e10
                                                fmt.Println(fi9, v19)
                                                var v20 uint32 = 392
                                                _ = v19
                                                _ = v20
                                            } else {
                                                v21 := GenS3{F0: 2, F1: 0, F2: -401}
                                                fmt.Println(fi10, v1, fi9)
                                                v13 += v13
                                                var v22 uint64 = 358
                                                _ = v21
                                                _ = v22
                                            }
                                        }
                                        if true {
                                            if ((v1 / 1.0) == float32(((0.1 * float64((1e10 / v1))) * (((-2.25 - (3.14159 / (float64((float32(0.1) / float32(3.14159))) / (1e10 + -2.25)))) + 0.0) + -2.25)))) {
                                                var v23 uint32 = 65535
                                                _ = v23
                                            }
                                        }
                                        fmt.Println(v1)
                                        fmt.Println(v1, v4)
                                        switch 32768 {
                                        case -9223372036854775808:
                                            var v24 int64 = -700
                                            _ = v24
                                        default:
                                            v25 := []int{fi10}
                                            _ = v25
                                        }
                                        v26 := GenS1{F0: (237 < v13)}
                                        if true {
                                            if (-445 >= (fi10 | (fi9 << ((4294967295) & 63)))) {
                                                var v27 uint64 = 888
                                                var v28 int16 = 197
                                                var v29 uint16 = (^(v4))
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            } else {
                                                v1 -= (v1 * 1e10)
                                                v30 := GenS4{F0: "hello"}
                                                _ = v30
                                            }
                                            v4 += v4
                                            v13 = -1
                                        }
                                        _ = v26
                                    } else {
                                        v31 := GenS1{F0: (!true)}
                                        delete(v2, int(v4))
                                        if (1 == 499) {
                                            var v32 float64 = 3.14159
                                            v33 := GenS5{F0: -2147483648, F1: (^(v4)), F2: 18446744073709551615}
                                            if false {
                                                if false {
                                                    if ((624 != 32767) && true) {
                                                        if false {
                                                            var v34 int = fi10
                                                            _ = v34
                                                        } else {
                                                            v32 -= v32
                                                            var v35 int64 = 1
                                                            fmt.Println(v35)
                                                            var v36 float64 = (float64(v1) - float64(1.5))
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                        if false {
                                                            var v37 int16 = 0
                                                            v38 := []int{2147483647}
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                    } else {
                                                        v39 := v2[32768]
                                                        switch (v39 >> ((65535) & 63)) {
                                                        case 32767:
                                                            if (false || (976 <= 2147483647)) {
                                                                if false {
                                                                    switch (v39 | fi10) {
                                                                    case -858:
                                                                        switch (("\n\t" + "\n\t") + ("hello" + "hello")) {
                                                                        case "a":
                                                                            var v40 int = fi9
                                                                            _ = v40
                                                                        case "foo bar":
                                                                            switch ("\n\t" + "") {
                                                                            case "":
                                                                                v41 := GenS6{F0: v13, F1: true, F2: 100}
                                                                                _ = v41
                                                                            case "a":
                                                                                fmt.Println(v39, fi9)
                                                                            case "\n\t":
                                                                                var v42 uint64 = 214
                                                                                _ = v42
                                                                            default:
                                                                                v13 -= (v13 + v13)
                                                                            }
                                                                        case "\n\t":
                                                                            v39 -= -2147483648
                                                                        }
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        var v43 float32 = ((v1 - float32((v32 - v32))) * v1)
                                                                        v39 = v39
                                                                        var v44 string = ("a" + ("" + ""))
                                                                        _ = v43
                                                                        _ = v44
                                                                    } else {
                                                                        v45 := GenS4{F0: ""}
                                                                        _ = v45
                                                                    }
                                                                }
                                                                v46 := GenS4{F0: "foo bar"}
                                                                var v47 uint16 = 42
                                                                _ = v46
                                                                _ = v47
                                                            }
                                                        case -1:
                                                            v48 := GenS7{F0: v4}
                                                            _ = v48
                                                        case 128:
                                                            v49 := GenS8{F0: (v4 ^ v4), F1: 118}
                                                            _ = v49
                                                        default:
                                                            fmt.Println(v4, v4, v32)
                                                        }
                                                        v13 = (v13 | int32(v39))
                                                        _ = v39
                                                    }
                                                    if (132 < 194) {
                                                        v4 = v4
                                                    }
                                                    var v50 uint32 = 2147483647
                                                    _ = v50
                                                }
                                                v4 -= v4
                                                v4 = v4
                                                if false {
                                                    v13 = (^(v13))
                                                    v51 := []int{fi10}
                                                    switch 42 {
                                                    case 32768:
                                                        v52 := GenS9{F0: 4294967295, F1: 9223372036854775807, F2: (2 < 503)}
                                                        _ = v52
                                                    case 0:
                                                        v53 := GenS10{F0: -118}
                                                        _ = v53
                                                    }
                                                    _ = v51
                                                } else {
                                                    fmt.Println(fi9, v4, v32)
                                                    v32 += v32
                                                    var v54 int64 = 270
                                                    _ = v54
                                                }
                                            } else {
                                                var v55 uint16 = (v4 ^ v4)
                                                var v56 int8 = 42
                                                _ = v55
                                                _ = v56
                                            }
                                            _ = v32
                                            _ = v33
                                        } else {
                                            fmt.Println(v4, v4)
                                            v57 := []int{fi9}
                                            _ = v57
                                        }
                                        _ = v31
                                    }
                                case 9223372036854775807:
                                    var v58 string = "\n\t"
                                    _ = v58
                                default:
                                    if false {
                                        v13 -= (v13 >> ((605) & 63))
                                        if (true && true) {
                                            v4 += (v4 & (^(v4)))
                                        }
                                    } else {
                                        v1 -= 0.1
                                        switch "\n\t" {
                                        case "foo bar":
                                            v59 := []int{}
                                            _ = v59
                                        case "":
                                            switch fi10 {
                                            case -732:
                                                v60 := GenS2{F0: 42}
                                                _ = v60
                                            case 32767:
                                                v61 := GenS7{F0: (v4 | 65535)}
                                                _ = v61
                                            case 32768:
                                                var v62 uint64 = 1
                                                _ = v62
                                            default:
                                                v4 -= (v4 & 1)
                                            }
                                        default:
                                            if ((v13 ^ 42) > v13) {
                                                v4 += 2
                                                v4 = v4
                                                v63 := GenS11{F0: fi9, F1: 65535}
                                                v64 := GenS12{F0: 32767, F1: 288}
                                                v65 := GenS6{F0: (v13 - -924), F1: (-59 == fi9), F2: -1}
                                                v4 = v4
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            } else {
                                                v66 := GenS13{F0: 0, F1: 127}
                                                fmt.Println(v13)
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v67 := GenS3{F0: 2147483647, F1: 83, F2: genFn0((v4 - v4), 1, true)}
                                                    _ = v67
                                                case "foo bar":
                                                    if true {
                                                        if (v4 >= 32767) {
                                                            fmt.Println(v4, fi10)
                                                            delete(v3, fi9)
                                                            var v68 int16 = genFn0(v4, 1, false)
                                                            _ = v68
                                                        }
                                                        v69 := []int{(fi10 % (((^(fi9))) | 1)), 1, (fi10 / ((fi10) | 1)), (fi10 * fi10)}
                                                        _ = v69
                                                    }
                                                case "hello":
                                                    v13 += v13
                                                }
                                                _ = v66
                                            }
                                        }
                                    }
                                }
                            }
                            v4 += v4
                        }
                        v70 := []int{-2147483648, fi9}
                        switch (fi9 & 32767) {
                        case -602:
                            v4 = (^(v4))
                        default:
                            _ = len(v70)
                        }
                        _ = v70
                    } else {
                        v71 := GenS13{F0: 32767, F1: -17}
                        v13 = v13
                        _ = v71
                    }
                    if true {
                        v72 := GenS14{F0: 0, F1: 2147483647, F2: uint32(v4)}
                        v4 = v4
                        v1 += (((v1 * v1) + v1) - v1)
                        _ = v72
                    }
                    var v73 uint = 255
                    _ = v73
                case 206:
                    v1 += ((float32(0.1) / float32((1e10 + 1e10))) * 1e10)
                }
                if true {
                    fmt.Println(fi9, fi10, v13)
                    if (!true) {
                        switch fi9 {
                        case -91:
                            v74 := []int{237, fi9, (-(fi9))}
                            _ = v74
                        case 42:
                            switch -744 {
                            case 128:
                                v3[fi10] = ""
                            case 127:
                                v1 = (v1 * v1)
                            case 0:
                                var v75 uint32 = 799
                                _ = v75
                            default:
                                fmt.Println(v1, fi10)
                            }
                        }
                        var v76 int16 = -32768
                        _ = v76
                    }
                } else {
                    v1 = 1e10
                    v77 := GenS3{F0: 211, F1: 137, F2: 1}
                    var v78 int8 = -128
                    _ = v77
                    _ = v78
                }
                _ = v13
            default:
                v79 := GenS3{F0: 730, F1: 127, F2: genFn0(v4, 0, (((1 != 2) || false) || (!false)))}
                _ = v79
            }
            if (!true) {
                var v80 int = (fi9 * (^(fi9)))
                _ = v80
            } else {
                v1 = 1e10
                fmt.Println(v1)
                if (!true) {
                    v81 := GenS15{F0: 4294967295, F1: 0}
                    var v82 int32 = -1
                    _ = v81
                    _ = v82
                }
                if (2 < 1) {
                    v83 := []int{int(v4), 32768, fi10}
                    var v84 int = (fi9 - (fi10 ^ (fi10 >> ((255) & 63))))
                    v1 += (float32(0.0) * v1)
                    v85 := GenS15{F0: 411, F1: 854}
                    _ = v83
                    _ = v84
                    _ = v85
                }
                if false {
                    v86 := GenS16{F0: -2147483648, F1: v4, F2: (v4 << ((18446744073709551615) & 63))}
                    if (true && (true || false)) {
                        var v87 int = (fi9 >> ((833) & 63))
                        var v88 int = (fi9 + (fi9 + (v87 % ((fi9) | 1))))
                        if false {
                            v89 := []int{(fi9 | (-(fi10))), -173}
                            v90 := GenS5{F0: 2147483647, F1: 2, F2: 0}
                            v91 := v3[fi9]
                            _ = v89
                            _ = v90
                            _ = v91
                        } else {
                            switch (^(fi10)) {
                            case 128:
                                v92 := []int{(fi9 % (((v88 >> ((478) & 63))) | 1)), fi10}
                                _ = v92
                            case 42:
                                if (2 != 45) {
                                    _ = len(v2)
                                    v93 := GenS17{F0: -606, F1: (false || (true || (((-2.25 * v1) + v1) != (v1 * v1)))), F2: ("\n\t" + "")}
                                    var v94 float64 = ((1.5 + (1e10 * 0.1)) * 3.14159)
                                    _ = v93
                                    _ = v94
                                }
                            case 32767:
                                v95 := GenS3{F0: 65535, F1: 199, F2: -478}
                                _ = v95
                            }
                            v96 := GenS17{F0: -9223372036854775808, F1: true, F2: "foo bar"}
                            _ = v96
                        }
                        _ = v87
                        _ = v88
                    }
                    _ = v86
                } else {
                    delete(v3, fi10)
                    v1 -= ((v1 + (1e10 + (0.0 + (v1 + -2.25)))) * (1e10 + 0.1))
                }
            }
            switch "hello" {
            case "hello":
                v4 -= (v4 - v4)
            case "foo bar":
                v97 := GenS13{F0: 454, F1: 1}
                _ = v97
            case "\n\t":
                switch (fi10 + 32767) {
                case -9223372036854775808:
                    delete(v2, fi10)
                case 0:
                    if (32767 <= -1) {
                        var v98 int8 = 42
                        _ = v98
                    }
                }
                fmt.Println(v4, v4)
                if false {
                    var v99 int = fi10
                    _ = v99
                } else {
                    if false {
                        var v100 uint8 = 0
                        _ = v100
                    } else {
                        var v101 int = (fi9 ^ (fi10 - fi9))
                        _ = v101
                    }
                    v4 = 363
                    var v102 uint32 = 65535
                    _ = v102
                }
            default:
                var v103 int8 = 127
                _ = v103
            }
            _ = fi10
        }
        v104 := GenS18{F0: 0, F1: -1}
        _ = fi9
        _ = v104
    }
    fmt.Println(v4)
    fmt.Println(v1, v4)
    v105 := []int{}
    var v106 uint = 18446744073709551615
    fmt.Println(v106, v1, v1)
    switch "a" {
    case "foo bar":
        if true {
            switch 127 {
            case 1:
                var v107 uint64 = uint64(v4)
                _ = v107
            case 0:
                _ = len(v3)
            case -1:
                switch -487 {
                case 128:
                    delete(v2, 127)
                default:
                    switch 0 {
                    case 42:
                        for fi108 := 0; fi108 < 9; fi108++ {
                            var v109 uint = v106
                            var v110 uint16 = v4
                            _ = fi108
                            _ = v109
                            _ = v110
                        }
                    }
                }
            default:
                v111 := GenS19{F0: 1, F1: 563, F2: -587}
                _ = v111
            }
        }
    case "\n\t":
        v112 := GenS17{F0: 687, F1: ((float32(0.1) + (3.14159 / (v1 + ((v1 - float32(float64(float32(3.14159)))) / float32((1e10 + 1.5)))))) > 0.1), F2: ("hello" + ("a" + ("hello" + "")))}
        _ = v112
    }
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v106)
    case "":
        var v113 uint32 = 0
        _ = v113
    default:
        v114 := GenS17{F0: 1, F1: (!true), F2: ("" + "foo bar")}
        _ = v114
    }
    v115 := GenS20{F0: 9223372036854775807}
    v116 := GenS21{F0: -9223372036854775808, F1: 275}
    delete(v2, 128)
    v105 = append(v105, 1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v105:", len(v105), v105)
    fmt.Println("v106:", v106)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v105
    _ = v106
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v4
}
