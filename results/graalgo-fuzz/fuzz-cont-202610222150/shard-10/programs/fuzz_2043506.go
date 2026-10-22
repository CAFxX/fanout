package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 int64; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint16; F2 int }
func genFn0(p0 uint, p1 uint) int {
    var v7 int32 = 0
    v8 := GenS1{F0: 1}
    var v9 float64 = 0.1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 264
}
func genFn1() uint16 {
    if true {
        switch 128 {
        case -711:
            if false {
                var v10 uint8 = 1
                v11 := GenS1{F0: genFn0(791, 255)}
                _ = v10
                _ = v11
            }
        }
    }
    v12 := []int{-910}
    var v13 int32 = -785
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            switch fi14 {
            case 0:
                switch "foo bar" {
                case "foo bar":
                    if false {
                        fmt.Println(fi15)
                        switch genFn0(1, 1) {
                        case 0:
                            if (3.14159 < float64(1.5)) {
                                var v16 uint16 = 32767
                                var v17 int16 = 1
                                _ = v16
                                _ = v17
                            }
                        case -9223372036854775808:
                            v18 := GenS1{F0: genFn0(1, 18446744073709551615)}
                            _ = v18
                        case 452:
                            if false {
                                fmt.Println(fi14, fi14, v13)
                            } else {
                                fmt.Println(v13, v13, fi15)
                                var v19 float64 = -2.25
                                v20 := GenS1{F0: genFn0(1, 65535)}
                                _ = v19
                                _ = v20
                            }
                        }
                    } else {
                        var v21 int16 = -391
                        _ = v21
                    }
                    v22 := GenS1{F0: -1}
                    _ = v22
                case "":
                    switch fi15 {
                    case 1:
                        switch "hello" {
                        case "foo bar":
                            var v23 uint64 = 2
                            _ = v23
                        case "\n\t":
                            var v24 uint64 = 18446744073709551615
                            _ = v24
                        default:
                            v13 = v13
                        }
                    case -9223372036854775808:
                        switch "hello" {
                        case "\n\t":
                            v25 := GenS1{F0: (fi15 << ((42) & 63))}
                            v26 := GenS1{F0: (fi14 + fi14)}
                            var v27 int64 = 42
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                    case 127:
                        if (false && true) {
                            var v28 int = (fi15 << ((2) & 63))
                            v13 = v13
                            _ = v28
                        }
                    default:
                        v13 += 821
                    }
                case "\n\t":
                    v13 = (v13 * (v13 - 42))
                default:
                    var v29 float32 = float32(1.5)
                    _ = v29
                }
            }
            v30 := GenS1{F0: (fi14 >> ((255) & 63))}
            switch (fi14 >> ((2) & 63)) {
            case 600:
                if (-128 != -63) {
                    if len(v12) > 0 {
                        v12[0] = (fi14 | (fi14 * (^(fi14))))
                    }
                } else {
                    if (!(-14 != -1)) {
                        var v31 int32 = (v13 | 1)
                        switch "a" {
                        case "\n\t":
                            v31 -= 1
                        case "":
                            if len(v12) > 0 {
                                v12[0] = fi14
                            }
                        case "foo bar":
                            v31 += v13
                        default:
                            v32 := GenS1{F0: (fi14 | (fi15 % ((-26) | 1)))}
                            _ = v32
                        }
                        switch "a" {
                        case "foo bar":
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(fi14)
                            case "a":
                                v33 := GenS1{F0: fi14}
                                _ = v33
                            case "":
                                v34 := GenS2{F0: ("\n\t" + ("foo bar" + "\n\t")), F1: 127}
                                _ = v34
                            default:
                                v35 := GenS2{F0: "a", F1: 253}
                                _ = v35
                            }
                        case "a":
                            v12 = append(v12, (^(fi15)))
                        case "":
                            fmt.Println(fi14, fi15)
                        default:
                            v36 := GenS3{F0: 32767}
                            _ = v36
                        }
                        _ = v31
                    } else {
                        v13 += v13
                    }
                }
            case -2147483648:
                if false {
                    if (0 != -1) {
                        v13 = v13
                    }
                }
            case 512:
                if true {
                    switch ("" + "") {
                    case "":
                        var v37 uint = 628
                        _ = v37
                    case "foo bar":
                        var v38 int64 = -9223372036854775808
                        _ = v38
                    default:
                        v13 -= -1
                    }
                }
            default:
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
            }
            switch ("a" + ("" + ("hello" + "\n\t"))) {
            case "foo bar":
                v13 = (v13 % (((-(v13))) | 1))
            case "a":
                v13 -= v13
                if (!true) {
                    v12 = append(v12, 1)
                    _ = len(v12)
                    var v39 uint32 = 526
                    _ = v39
                }
                var v40 uint = 568
                _ = v40
            }
            _ = fi15
            _ = v30
        }
        if false {
            var v41 uint32 = 316
            _ = len(v12)
            for fi42 := 0; fi42 < 14; fi42++ {
                var v43 uint = 2
                if false {
                    fmt.Println(v13, v43, v13)
                }
                v43 -= (v43 >> ((v43) & 63))
                _ = fi42
                _ = v43
            }
            _ = v41
        } else {
            var v44 float64 = (float64(3.14159) + 1.5)
            v13 -= (v13 | v13)
            _ = v44
        }
        _ = fi14
    }
    _ = v12
    _ = v13
    return 1
}

func main() {
    var v1 int32 = -2147483648
    var v2 uint64 = 42
    v3 := make(map[int]bool)
    var v4 string = (("foo bar" + "a") + "foo bar")
    v5 := make(map[string]bool)
    var v6 string = ("\n\t" + (("a" + ((v4 + ((v4 + "") + "foo bar")) + v4)) + "a"))
    if (42 > 42) {
        v2 = (v2 ^ v2)
        v6 = ""
        switch ((((("" + "hello") + "hello") + ("" + "foo bar")) + v4) + v6) {
        case "\n\t":
            var v45 int32 = 1
            _ = v45
        case "a":
            switch v6 {
            case "hello":
                if ((2 < 0) || true) {
                    v2 -= v2
                }
            }
        default:
            v2 = v2
        }
    } else {
        var v46 int32 = int32(v2)
        v46 = (v46 - v46)
        v47 := GenS3{F0: 2}
        _ = v46
        _ = v47
    }
    v48 := GenS3{F0: 32767}
    if (true || false) {
        var v49 float32 = (float32(float64(1.5)) / 1.5)
        var v50 int64 = -1
        v4 = ("\n\t" + v4)
        _ = v49
        _ = v50
    }
    v4 = ""
    v51 := []int{}
    if true {
        switch 127 {
        case 1:
            var v52 float64 = 1.5
            _ = v52
        default:
            for fi53 := 0; fi53 < 16; fi53++ {
                for fi54 := 0; fi54 < 10; fi54++ {
                    v2 -= (v2 | (v2 & (v2 >> ((56) & 63))))
                    fmt.Println(fi53)
                    var v55 uint16 = 255
                    v56 := GenS3{F0: (v55 >> ((1) & 63))}
                    if false {
                        if (false && true) {
                            switch (-(fi53)) {
                            case 127:
                                switch (fi54 / ((-1) | 1)) {
                                case 32767:
                                    var v57 string = v6
                                    _ = v57
                                case 325:
                                    var v58 uint8 = 172
                                    _ = v58
                                default:
                                    if (!true) {
                                        fmt.Println(fi53, v6)
                                        if (!false) {
                                            v5[v4] = true
                                            v6 = v4
                                        }
                                    }
                                    if len(v51) >= 2 {
                                        v51 = v51[1:len(v51)]
                                    }
                                    v6 = v6
                                }
                            case 4294967295:
                                v4 = "foo bar"
                            default:
                                var v59 float32 = 1.5
                                _ = v59
                            }
                            if (v6 != (v6 + v6)) {
                                v60 := v5[v6]
                                delete(v3, (fi53 - (-(fi54))))
                                var v61 uint8 = uint8(v2)
                                _ = v60
                                _ = v61
                            } else {
                                if len(v51) > 0 {
                                    v51[0] = int(v2)
                                }
                            }
                            v51 = append(v51, fi54)
                        } else {
                            delete(v5, (v4 + v4))
                            var v62 uint32 = 360
                            _ = v62
                        }
                        v63 := GenS3{F0: 2}
                        if len(v51) > 0 {
                            v51[0] = (-(fi54))
                        }
                        _ = v63
                    }
                    var v64 uint8 = 0
                    _ = fi54
                    _ = v55
                    _ = v56
                    _ = v64
                }
                var v65 float64 = 0.1
                _ = fi53
                _ = v65
            }
        }
    } else {
        var v66 string = ("hello" + "\n\t")
        v67 := GenS2{F0: v6, F1: 1}
        _ = v66
        _ = v67
    }
    fmt.Println(v6, v4, v4)
    if false {
        var v68 float32 = (1.5 - -2.25)
        fmt.Println(v2)
        _ = v68
    }
    _ = len(v51)
    for fi69 := 0; fi69 < 11; fi69++ {
        fmt.Println(v6, v2)
        if ((0 <= 1) || (!(-32768 != -779))) {
            if (true || (((v4 + "foo bar") >= (v4 + (v6 + "foo bar"))) || (true || (128 >= 128)))) {
                v70 := GenS1{F0: 127}
                fmt.Println(v4)
                _ = v70
            }
        }
        v2 = v2
        _ = fi69
    }
    fmt.Println(v1, v2)
    for fi71 := 0; fi71 < 10; fi71++ {
        var v72 uint8 = 165
        switch -1 {
        case 2147483647:
            var v73 int8 = -86
            _ = v73
        case -9223372036854775808:
            v4 = v4
        }
        _ = fi71
        _ = v72
    }
    switch 42 {
    case 2147483647:
        v74 := GenS3{F0: 20}
        _ = v74
    case -9223372036854775808:
        for fi75 := 0; fi75 < 7; fi75++ {
            v1 = (v1 / ((0) | 1))
            switch (("" + ("a" + "\n\t")) + "a") {
            case "a":
                var v76 int16 = -32768
                _ = v76
            case "foo bar":
                _ = len(v51)
            case "":
                var v77 int8 = 92
                _ = v77
            default:
                if (!false) {
                    _ = len(v51)
                    var v78 uint32 = 2
                    delete(v5, v6)
                    _ = v78
                } else {
                    for fi79 := 0; fi79 < 4; fi79++ {
                        var v80 float32 = 1.5
                        v81 := GenS4{F0: -1}
                        if false {
                            var v82 uint64 = v2
                            _ = v82
                        } else {
                            v83 := GenS1{F0: genFn0(2, 450)}
                            _ = v83
                        }
                        var v84 uint16 = 65535
                        switch fi75 {
                        case 0:
                            fmt.Println(fi79)
                        case -310:
                            var v85 uint64 = v2
                            _ = v85
                        case 2147483647:
                            if false {
                                if false {
                                    v1 = v1
                                    var v86 int16 = -23
                                    if false {
                                        _ = len(v5)
                                        var v87 uint = 0
                                        v86 = (^(v86))
                                        if false {
                                            fmt.Println(v6)
                                            switch v6 {
                                            case "":
                                                var v88 string = "\n\t"
                                                v89 := GenS2{F0: v6, F1: 255}
                                                v90 := GenS1{F0: (fi75 >> ((4294967295) & 63))}
                                                if len(v51) > 0 {
                                                    v51[0] = (fi79 << (((v87 * v87)) & 63))
                                                }
                                                _ = len(v51)
                                                _ = v88
                                                _ = v89
                                                _ = v90
                                            case "foo bar":
                                                fmt.Println(v87)
                                            case "\n\t":
                                                v87 = (v87 << (((v87 | v87)) & 63))
                                            }
                                        } else {
                                            v91 := GenS3{F0: (v84 * uint16(v2))}
                                            if false {
                                                var v92 int = -9223372036854775808
                                                var v93 int64 = 185
                                                _ = v92
                                                _ = v93
                                            } else {
                                                if true {
                                                    v86 += -765
                                                    if ((v1 | v1) >= (v1 << (((^(v87))) & 63))) {
                                                        var v94 uint32 = 414
                                                        var v95 int32 = v1
                                                        if (9223372036854775807 > -9223372036854775808) {
                                                            v95 = v95
                                                            v94 = v94
                                                            _ = len(v51)
                                                            switch v4 {
                                                            case "hello":
                                                                v95 = v95
                                                            }
                                                            v96 := GenS5{F0: (v95 << ((v87) & 63)), F1: 42, F2: (v94 << ((42) & 63))}
                                                            _ = v96
                                                        } else {
                                                            v94 -= 65535
                                                            var v97 uint = (^(v87))
                                                            if len(v51) >= 2 {
                                                                v51 = v51[1:len(v51)]
                                                            }
                                                            if (0.0 <= 0.0) {
                                                                var v98 int = int(v84)
                                                                _ = v98
                                                            } else {
                                                                var v99 string = v4
                                                                _ = v99
                                                            }
                                                            _ = v97
                                                        }
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                }
                                                var v100 uint32 = 4294967295
                                                switch (v6 + v6) {
                                                case "":
                                                    switch fi75 {
                                                    case 0:
                                                        v87 = 768
                                                    }
                                                }
                                                _ = v100
                                            }
                                            _ = v91
                                        }
                                        _ = v87
                                    } else {
                                        fmt.Println(v2, v1)
                                    }
                                    _ = v86
                                } else {
                                    if false {
                                        switch "a" {
                                        case "hello":
                                            v84 = (v84 / (((v84 ^ (v84 | v84))) | 1))
                                            var v101 int16 = -1
                                            _ = v101
                                        }
                                        var v102 int8 = -1
                                        var v103 uint8 = 42
                                        _ = v102
                                        _ = v103
                                    } else {
                                        if ((!(true || (false || ((v80 + v80) < ((v80 * (v80 - v80)) * v80))))) || (76 >= 100)) {
                                            var v104 uint16 = (v84 % (((v84 % ((521) | 1))) | 1))
                                            fmt.Println(v80, fi75)
                                            if (!(false && (true || (!false)))) {
                                                v105 := GenS2{F0: (v4 + "foo bar"), F1: 255}
                                                _ = v105
                                            }
                                            _ = v104
                                        } else {
                                            v106 := GenS5{F0: 0, F1: 0, F2: 42}
                                            v107 := GenS3{F0: v84}
                                            if (false || (127 == 42)) {
                                                if true {
                                                    var v108 float32 = (v80 - 0.1)
                                                    _ = v108
                                                }
                                                v109 := GenS3{F0: v84}
                                                var v110 int = genFn0(42, 65535)
                                                _ = v109
                                                _ = v110
                                            } else {
                                                if len(v51) > 0 {
                                                    v51[0] = fi79
                                                }
                                                delete(v5, (v6 + v6))
                                            }
                                            _ = v106
                                            _ = v107
                                        }
                                    }
                                }
                                v84 += v84
                                switch "hello" {
                                case "foo bar":
                                    v111 := GenS5{F0: v1, F1: 42, F2: 2147483648}
                                    if true {
                                        var v112 float32 = v80
                                        _ = v112
                                    } else {
                                        fmt.Println(v84)
                                    }
                                    var v113 int16 = 32767
                                    _ = v111
                                    _ = v113
                                case "\n\t":
                                    var v114 float64 = -2.25
                                    _ = v114
                                case "hello":
                                    if (!true) {
                                        if false {
                                            if len(v51) > 0 {
                                                v51[0] = -488
                                            }
                                        } else {
                                            v1 = v1
                                        }
                                    }
                                default:
                                    var v115 int = fi75
                                    _ = v115
                                }
                                v116 := GenS6{F0: 42, F1: (v84 << ((1) & 63)), F2: genFn0(752, 138)}
                                _ = v116
                            }
                        }
                        _ = fi79
                        _ = v80
                        _ = v81
                        _ = v84
                    }
                }
            }
            v5[v6] = (128 != 127)
            _ = fi75
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v51
    _ = v6
}
