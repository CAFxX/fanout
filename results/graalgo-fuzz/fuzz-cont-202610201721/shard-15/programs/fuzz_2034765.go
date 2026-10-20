package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS12 struct { F0 bool; F1 uint }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 string; F1 int16; F2 uint32 }
type GenS15 struct { F0 uint16; F1 string }
func genFn0(p0 uint16) int64 {
    v5 := []int{32768, 42, -1, 128}
    v5 = append(v5, 4294967295)
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 string) uint8 {
    v6 := GenS1{F0: 1, F1: -9223372036854775808}
    switch 127 {
    case -1:
        v7 := GenS1{F0: 1, F1: 1}
        _ = v7
    case -597:
        fmt.Println(p0)
    }
    v8 := []int{443, 32767, -582}
    p0 = "hello"
    _ = p0
    _ = v6
    _ = v8
    return 128
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 199
    var v3 uint8 = 127
    var v4 uint8 = 128
    if (true && (1 > 727)) {
        v2 += genFn1("foo bar")
        fmt.Println(v2)
        for fi9 := 0; fi9 < 12; fi9++ {
            for fi10 := 0; fi10 < 7; fi10++ {
                if ((0.0 / 1.5) >= 0.0) {
                    fmt.Println(v3, fi10)
                    var v11 uint64 = 42
                    fmt.Println(v11, v3)
                    if (false || false) {
                        v4 = genFn1("hello")
                    }
                    fmt.Println(v3)
                    _ = v11
                } else {
                    switch "foo bar" {
                    case "foo bar":
                        switch (fi9 / ((fi10) | 1)) {
                        case 1:
                            fmt.Println(fi10, fi10, v2)
                            v2 += (v3 ^ (^(v2)))
                            var v12 uint16 = 65535
                            _ = v12
                        case -17:
                            v3 -= (^(v3))
                        case -935:
                            v13 := GenS2{F0: 8, F1: fi9}
                            _ = v13
                        default:
                            v14 := []int{(fi10 - (fi10 ^ (fi10 * fi9)))}
                            _ = v14
                        }
                        switch fi9 {
                        case 32768:
                            v15 := GenS2{F0: 809, F1: (fi10 + fi9)}
                            _ = v15
                        case 567:
                            fmt.Println(v2, fi9)
                        case -311:
                            v16 := GenS1{F0: 773, F1: (fi10 / ((fi9) | 1))}
                            _ = v16
                        default:
                            if (1 < 2) {
                                fmt.Println(v3)
                                v17 := GenS3{F0: -988, F1: v3}
                                _ = v17
                            } else {
                                v18 := GenS4{F0: 974, F1: "hello"}
                                fmt.Println(v2, v3)
                                switch "" {
                                case "":
                                    if ((^(v2)) >= genFn1("hello")) {
                                        v19 := []int{}
                                        var v20 uint64 = 18446744073709551615
                                        _ = v19
                                        _ = v20
                                    }
                                case "a":
                                    v21 := []int{}
                                    if (((0.1 == float64(3.14159)) || false) && ((!("" >= "hello")) || (false || ("" >= "a")))) {
                                        if len(v21) > 0 {
                                            v21[0] = 571
                                        }
                                        v3 = v3
                                        if (!(1.5 != 0.0)) {
                                            if true {
                                                v22 := v1[1]
                                                _ = v22
                                            } else {
                                                if (true || ((1 <= 40) && (!(true || (75 >= -128))))) {
                                                    var v23 bool = (!false)
                                                    v24 := GenS5{F0: ((!v23) || v23)}
                                                    _ = v23
                                                    _ = v24
                                                }
                                                switch "hello" {
                                                case "":
                                                    var v25 uint = 1
                                                    _ = v25
                                                case "a":
                                                    if (!false) {
                                                        v4 = v3
                                                        fmt.Println(fi9, v3)
                                                        delete(v1, 453)
                                                        v3 -= (v3 ^ 2)
                                                        var v26 int8 = 62
                                                        _ = v26
                                                    } else {
                                                        v21 = append(v21, 0)
                                                        _ = len(v1)
                                                    }
                                                default:
                                                    switch "hello" {
                                                    case "":
                                                        v4 = (v4 % ((genFn1((("" + "hello") + ("a" + "a")))) | 1))
                                                    case "a":
                                                        if false {
                                                            var v27 uint8 = genFn1("")
                                                            v2 -= genFn1("\n\t")
                                                            _ = v27
                                                        } else {
                                                            switch (("a" + "hello") + "foo bar") {
                                                            case "foo bar":
                                                                switch "hello" {
                                                                case "a":
                                                                    fmt.Println(fi10, v2)
                                                                }
                                                            case "a":
                                                                if len(v21) >= 2 {
                                                                    v21 = v21[1:len(v21)]
                                                                }
                                                            case "":
                                                                if false {
                                                                    switch "hello" {
                                                                    case "a":
                                                                        if false {
                                                                            v28 := GenS6{F0: "hello"}
                                                                            var v29 int = (fi10 * 937)
                                                                            _ = v28
                                                                            _ = v29
                                                                        }
                                                                    case "":
                                                                        v30 := GenS3{F0: 1, F1: genFn1("hello")}
                                                                        _ = v30
                                                                    default:
                                                                        var v31 uint8 = (v3 - 228)
                                                                        _ = v31
                                                                    }
                                                                } else {
                                                                    var v32 int64 = -9223372036854775808
                                                                    delete(v1, 4294967295)
                                                                    if ("" >= ("" + ("foo bar" + "a"))) {
                                                                        fmt.Println(fi9)
                                                                    }
                                                                    _ = v32
                                                                }
                                                            default:
                                                                v2 -= v4
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v33 := v1[462]
                                                        _ = v33
                                                    default:
                                                        if (false || ("" == (("" + ("" + "a")) + "\n\t"))) {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                var v34 uint = 255
                                                                _ = v34
                                                            default:
                                                                v21 = append(v21, (-(fi10)))
                                                            }
                                                        } else {
                                                            v4 = v4
                                                            if false {
                                                                var v35 float64 = 0.1
                                                                if ("a" < ("\n\t" + "\n\t")) {
                                                                    if false {
                                                                        var v36 int16 = 42
                                                                        switch (fi10 >> ((18446744073709551615) & 63)) {
                                                                        case -88:
                                                                            var v37 float64 = v35
                                                                            _ = v37
                                                                        default:
                                                                            delete(v1, 4294967295)
                                                                        }
                                                                        if (!(4294967295 != 622)) {
                                                                            v38 := GenS1{F0: 2, F1: fi10}
                                                                            _ = v38
                                                                        } else {
                                                                            v39 := v1[1]
                                                                            if (v39 <= "\n\t") {
                                                                                v21 = append(v21, fi9)
                                                                                v40 := GenS5{F0: false}
                                                                                var v41 int16 = (v36 | (^(v36)))
                                                                                _ = v40
                                                                                _ = v41
                                                                            } else {
                                                                                var v42 float32 = 0.1
                                                                                switch v39 {
                                                                                case "foo bar":
                                                                                    switch (fi9 >> ((914) & 63)) {
                                                                                    case 32768:
                                                                                        v3 = genFn1((("" + (v39 + ("foo bar" + v39))) + ""))
                                                                                    }
                                                                                case "\n\t":
                                                                                    if ((^(v4)) == 231) {
                                                                                        v35 += 0.1
                                                                                    } else {
                                                                                        if false {
                                                                                            v43 := GenS4{F0: (v36 >> ((255) & 63)), F1: "foo bar"}
                                                                                            _ = v43
                                                                                        }
                                                                                        v4 += (v2 >> ((4294967295) & 63))
                                                                                        if (!false) {
                                                                                            v44 := GenS4{F0: (v36 >> ((42) & 63)), F1: "\n\t"}
                                                                                            var v45 uint = 518
                                                                                            _ = v44
                                                                                            _ = v45
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v42
                                                                            }
                                                                            v21 = append(v21, fi10)
                                                                            _ = v39
                                                                        }
                                                                        var v46 uint32 = 65535
                                                                        _ = v36
                                                                        _ = v46
                                                                    } else {
                                                                        if (-760 < -32768) {
                                                                            _ = len(v1)
                                                                            delete(v1, 0)
                                                                        } else {
                                                                            v1[2] = "foo bar"
                                                                            fmt.Println(fi10, v35, v4)
                                                                            if (0.1 < 1e10) {
                                                                                if (42 >= -128) {
                                                                                    if false {
                                                                                        v35 += v35
                                                                                        if (false || false) {
                                                                                            var v47 int64 = 1
                                                                                            if (!false) {
                                                                                                v3 += 54
                                                                                            }
                                                                                            _ = v47
                                                                                        } else {
                                                                                            var v48 int64 = 1
                                                                                            v48 += (v48 + 42)
                                                                                            _ = v48
                                                                                        }
                                                                                        v49 := GenS1{F0: 94, F1: 42}
                                                                                        _ = v49
                                                                                    }
                                                                                } else {
                                                                                    v4 = 114
                                                                                    v50 := GenS1{F0: 0, F1: (^(fi9))}
                                                                                    fmt.Println(v3, v4)
                                                                                    _ = v50
                                                                                }
                                                                                v4 -= (v4 % (((v2 << ((4294967295) & 63))) | 1))
                                                                                var v51 float32 = float32(float64((3.14159 / (float32(v35) + 0.1))))
                                                                                _ = v51
                                                                            } else {
                                                                                fmt.Println(v4, fi10, v35)
                                                                            }
                                                                        }
                                                                        if (42 == 42) {
                                                                            if true {
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    if false {
                                                                                        fmt.Println(fi10)
                                                                                        switch ("hello" + "\n\t") {
                                                                                        case "hello":
                                                                                            if len(v21) >= 2 {
                                                                                                v21 = v21[1:len(v21)]
                                                                                            }
                                                                                        case "":
                                                                                            var v52 float64 = (v35 * v35)
                                                                                            _ = v52
                                                                                        default:
                                                                                            v2 += (v4 | v3)
                                                                                        }
                                                                                        v2 = v4
                                                                                    }
                                                                                default:
                                                                                    var v53 bool = false
                                                                                    _ = v53
                                                                                }
                                                                                fmt.Println(v35)
                                                                                _ = len(v21)
                                                                            }
                                                                            v54 := v1[513]
                                                                            if (0 == 42) {
                                                                                switch fi10 {
                                                                                case -9223372036854775808:
                                                                                    if true {
                                                                                        fmt.Println(v54, v2)
                                                                                        var v55 float32 = 0.0
                                                                                        if ((((v35 >= 1.5) && false) || false) || (false && true)) {
                                                                                            v54 = ""
                                                                                            if (1 >= 42) {
                                                                                                var v56 string = ((v54 + ((((("hello" + (v54 + "a")) + "foo bar") + (v54 + "foo bar")) + v54) + v54)) + v54)
                                                                                                _ = v56
                                                                                            }
                                                                                            if false {
                                                                                                if true {
                                                                                                    v57 := GenS7{F0: genFn1((v54 + ""))}
                                                                                                    fmt.Println(fi10, v55)
                                                                                                    _ = v57
                                                                                                }
                                                                                                if ((!true) || false) {
                                                                                                    switch (fi10 | fi9) {
                                                                                                    case -330:
                                                                                                        var v58 uint8 = (v4 / ((v3) | 1))
                                                                                                        _ = v58
                                                                                                    case 1:
                                                                                                        var v59 uint = 18446744073709551615
                                                                                                        _ = v59
                                                                                                    default:
                                                                                                        var v60 int = fi9
                                                                                                        _ = v60
                                                                                                    }
                                                                                                    v61 := GenS8{F0: 4294967295}
                                                                                                    fmt.Println(v3, v54)
                                                                                                    _ = v61
                                                                                                } else {
                                                                                                    v55 = v55
                                                                                                }
                                                                                                var v62 uint32 = 0
                                                                                                _ = v62
                                                                                            }
                                                                                        }
                                                                                        if len(v21) > 0 {
                                                                                            v21[0] = -722
                                                                                        }
                                                                                        v63 := GenS4{F0: -169, F1: "a"}
                                                                                        _ = v55
                                                                                        _ = v63
                                                                                    } else {
                                                                                        var v64 uint = 1
                                                                                        var v65 uint16 = 255
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                    }
                                                                                default:
                                                                                    v54 = ("" + ((v54 + v54) + (v54 + v54)))
                                                                                }
                                                                                v3 += 42
                                                                            } else {
                                                                                v3 = (v4 | genFn1(v54))
                                                                                v21 = append(v21, fi9)
                                                                            }
                                                                            _ = v54
                                                                        } else {
                                                                            var v66 uint64 = 4294967295
                                                                            switch (fi10 + fi9) {
                                                                            case 2147483647:
                                                                                var v67 uint64 = (v66 ^ (v66 / ((v66) | 1)))
                                                                                _ = v67
                                                                            case 0:
                                                                                v68 := GenS3{F0: 351, F1: v4}
                                                                                _ = v68
                                                                            case 32767:
                                                                                v4 = 79
                                                                            }
                                                                            _ = v66
                                                                        }
                                                                        fmt.Println(v35, v4, v35)
                                                                    }
                                                                    switch (fi9 & (-(fi10))) {
                                                                    case -67:
                                                                        v35 -= v35
                                                                    default:
                                                                        v35 = 1.5
                                                                    }
                                                                } else {
                                                                    if (true || true) {
                                                                        v2 = v4
                                                                        v69 := GenS9{F0: -20, F1: (!(!true))}
                                                                        _ = v69
                                                                    } else {
                                                                        var v70 int64 = 1
                                                                        _ = v70
                                                                    }
                                                                    v71 := GenS4{F0: 32767, F1: "hello"}
                                                                    _ = v71
                                                                }
                                                                _ = v35
                                                            } else {
                                                                var v72 int = -9223372036854775808
                                                                v72 -= fi9
                                                                var v73 uint64 = 815
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            v4 += 2
                                            var v74 float64 = ((float64(float32((-2.25 + 0.1))) + 0.0) * -2.25)
                                            _ = v74
                                        } else {
                                            delete(v1, 2147483648)
                                            switch fi9 {
                                            case 128:
                                                var v75 int8 = -1
                                                _ = v75
                                            case 2147483647:
                                                var v76 int64 = -9223372036854775808
                                                _ = v76
                                            case 4294967295:
                                                if false {
                                                    var v77 uint32 = 4294967295
                                                    _ = v77
                                                }
                                            default:
                                                switch (fi9 * (fi9 * fi9)) {
                                                case 9223372036854775807:
                                                    if false {
                                                        v78 := GenS10{F0: -559}
                                                        _ = v78
                                                    } else {
                                                        v79 := GenS11{F0: (v4 | (^(v4))), F1: 1, F2: 255}
                                                        fmt.Println(v4, v3, fi10)
                                                        _ = v79
                                                    }
                                                case 127:
                                                    fmt.Println(v4)
                                                case 32768:
                                                    if false {
                                                        if false {
                                                            var v80 uint16 = 42
                                                            var v81 uint32 = 521
                                                            _ = v80
                                                            _ = v81
                                                        } else {
                                                            v3 -= (^(v3))
                                                        }
                                                        var v82 bool = (42 <= -42)
                                                        v83 := GenS12{F0: (44 == 42), F1: 4294967295}
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                default:
                                                    if len(v21) >= 2 {
                                                        v21 = v21[1:len(v21)]
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    v21 = append(v21, 2147483647)
                                    _ = v21
                                }
                                _ = v18
                            }
                        }
                        v2 += v3
                        if true {
                            if (!true) {
                                switch fi10 {
                                case 0:
                                    var v84 float64 = 1e10
                                    _ = v84
                                case 127:
                                    v2 += (v3 >> ((42) & 63))
                                default:
                                    var v85 uint8 = genFn1((("foo bar" + ("\n\t" + (("a" + (("a" + "a") + "\n\t")) + ("\n\t" + (("a" + "foo bar") + "a"))))) + "hello"))
                                    if (42 <= 1) {
                                        var v86 int8 = 58
                                        var v87 int = fi9
                                        _ = v86
                                        _ = v87
                                    } else {
                                        v88 := GenS12{F0: (18446744073709551615 != 42), F1: 4294967295}
                                        v4 += (v85 * (v4 | 136))
                                        var v89 float64 = -2.25
                                        _ = v88
                                        _ = v89
                                    }
                                    _ = v85
                                }
                            }
                        } else {
                            v90 := []int{(fi10 >> ((42) & 63)), fi9, fi9, 813}
                            _ = v90
                        }
                    case "\n\t":
                        v2 = (v3 / ((249) | 1))
                    }
                }
                var v91 uint16 = 255
                _ = fi10
                _ = v91
            }
            v92 := []int{fi9}
            _ = fi9
            _ = v92
        }
    }
    if (18446744073709551615 > 664) {
        var v93 uint16 = 0
        var v94 uint16 = (v93 >> ((4294967295) & 63))
        var v95 int32 = -2147483648
        _ = v93
        _ = v94
        _ = v95
    } else {
        var v96 int8 = -128
        v97 := []int{}
        var v98 string = ""
        _ = v96
        _ = v97
        _ = v98
    }
    v3 = v3
    switch -9223372036854775808 {
    case -53:
        var v99 uint8 = (v3 ^ genFn1("hello"))
        _ = v99
    }
    var v100 string = ""
    switch v100 {
    case "hello":
        v101 := GenS13{F0: 9223372036854775807}
        _ = v101
    case "a":
        switch v100 {
        case "a":
            v3 = genFn1(v100)
        case "hello":
            for fi102 := 0; fi102 < 17; fi102++ {
                v103 := GenS14{F0: v100, F1: 1, F2: 2}
                delete(v1, 2147483647)
                _ = fi102
                _ = v103
            }
        }
    case "":
        for fi104 := 0; fi104 < 4; fi104++ {
            if (!(42 != 100)) {
                v100 = "hello"
            } else {
                switch "" {
                case "hello":
                    switch fi104 {
                    case 4294967295:
                        v3 = (v4 % (((v3 | 2)) | 1))
                    case 9223372036854775807:
                        for fi105 := 0; fi105 < 9; fi105++ {
                            delete(v1, 65535)
                            v3 = (v3 % (((v2 + (^(v2)))) | 1))
                            _ = fi105
                        }
                    case 0:
                        var v106 uint64 = 912
                        _ = v106
                    default:
                        var v107 uint64 = 722
                        _ = v107
                    }
                case "a":
                    v4 += v3
                }
                v108 := []int{}
                _ = v108
            }
            var v109 uint8 = 178
            v3 += (^(v4))
            _ = fi104
            _ = v109
        }
    default:
        var v110 float32 = 1e10
        _ = v110
    }
    v2 = 127
    v111 := []int{32768, 42, 0, 127}
    _ = len(v1)
    var v112 uint32 = 0
    fmt.Println(v100, v2)
    for fi113 := 0; fi113 < 8; fi113++ {
        if (-1 > 2147483647) {
            v114 := GenS1{F0: 515, F1: (fi113 | fi113)}
            v115 := GenS15{F0: 65535, F1: ("" + "")}
            _ = v114
            _ = v115
        }
        _ = fi113
    }
    var v116 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v100:", v100)
    fmt.Println("slicelen v111:", len(v111), v111)
    fmt.Println("v112:", v112)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v100
    _ = v111
    _ = v112
    _ = v116
    _ = v2
    _ = v3
    _ = v4
}
