package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int; F2 uint8 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 uint16; F1 uint16 }
type GenS8 struct { F0 int64; F1 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint8; F1 int64 }

func main() {
    var v1 string = "foo bar"
    var v2 uint = 2
    var v3 uint16 = 2
    var v4 string = ((("foo bar" + "hello") + ("foo bar" + ("foo bar" + v1))) + v1)
    v5 := make(map[string]string)
    v6 := make(map[uint32]int)
    v7 := GenS1{F0: v3}
    v8 := GenS2{F0: false}
    if false {
        v4 = v1
    }
    v9 := []int{-9223372036854775808}
    var v10 uint8 = uint8(v2)
    var v11 int = 9223372036854775807
    switch v1 {
    case "foo bar":
        v12 := GenS3{F0: 4294967295}
        _ = v12
    case "":
        v13 := GenS2{F0: false}
        _ = v13
    default:
        v14 := GenS1{F0: (v3 << ((0) & 63))}
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 16; fi16++ {
            switch "" {
            case "\n\t":
                var v17 int32 = 0
                _ = v17
            case "foo bar":
                delete(v6, 2147483648)
            }
            v9 = append(v9, 2147483647)
            v3 += (^(v3))
            if (2 <= 4294967295) {
                var v18 uint8 = v10
                v19 := GenS3{F0: 1}
                _ = v18
                _ = v19
            }
            if (((float32(1.5) + 0.1) <= float32((1.5 * 1e10))) && (-1 >= 42)) {
                var v20 uint16 = (^(v3))
                v4 = "hello"
                var v21 uint16 = 819
                var v22 uint32 = 861
                _ = v20
                _ = v21
                _ = v22
            } else {
                _ = len(v9)
            }
            _ = fi16
        }
        switch "\n\t" {
        case "hello":
            v23 := GenS1{F0: 65535}
            _ = v23
        case "a":
            v24 := GenS1{F0: v3}
            _ = v24
        default:
            for fi25 := 0; fi25 < 2; fi25++ {
                if (!((1e10 / 0.1) == -2.25)) {
                    v26 := GenS4{F0: uint32(v10)}
                    _ = len(v9)
                    v4 = v1
                    if false {
                        v27 := GenS3{F0: 0}
                        _ = v27
                    } else {
                        switch (^(v11)) {
                        case 42:
                            v1 = "foo bar"
                        default:
                            var v28 int64 = -88
                            _ = v28
                        }
                        _ = len(v9)
                    }
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    if ((((1e10 * 0.0) * (3.14159 - (float32((0.0 - 3.14159)) + float32((0.0 * (((float64(0.0) - 1e10) + 1e10) / float64(1e10))))))) / 3.14159) < 0.1) {
                        switch v1 {
                        case "\n\t":
                            if true {
                                fmt.Println(v10, v4, v4)
                                fmt.Println(fi25, v2, v11)
                                var v29 int = (fi15 & -477)
                                v11 = (-(v11))
                                _ = v29
                            }
                            if (float64(-2.25) < (3.14159 / 1.0)) {
                                v30 := v5[v1]
                                _ = v30
                            }
                            v31 := GenS2{F0: ((352 > -32768) && (42 != 9223372036854775807))}
                            _ = v31
                        }
                    }
                    v3 -= (v3 % (((v3 % ((42) | 1))) | 1))
                    v32 := GenS3{F0: 321}
                    switch -2147483648 {
                    case -977:
                        fmt.Println(v3, v4, v4)
                    case -450:
                        fmt.Println(v11, fi25)
                    case 813:
                        v33 := GenS2{F0: (-2147483648 < 0)}
                        _ = v33
                    }
                    _ = v26
                    _ = v32
                } else {
                    if ((v3 & (v3 ^ 1)) < (v3 & (^(v3)))) {
                        var v34 float64 = (1e10 - -2.25)
                        _ = v34
                    } else {
                        v11 = (fi25 | fi25)
                        v11 -= (fi25 ^ (-(fi15)))
                    }
                    var v35 uint64 = 18446744073709551615
                    _ = v35
                }
                v36 := GenS5{F0: (-(v11)), F1: (fi15 >> ((v2) & 63)), F2: v10}
                v37 := GenS6{F0: fi15, F1: uint16(v10)}
                _ = fi25
                _ = v36
                _ = v37
            }
        }
        if true {
            fmt.Println(fi15, v10)
            v38 := GenS5{F0: (fi15 / (((fi15 / ((-2147483648) | 1))) | 1)), F1: v11, F2: v10}
            _ = v38
        }
        _ = fi15
    }
    var v39 float32 = (1e10 - (1.5 * float32(((float64((3.14159 / 0.1)) + (3.14159 + 0.1)) - float64(0.1)))))
    switch (v11 / (((v11 ^ 127)) | 1)) {
    case 127:
        v40 := GenS1{F0: v3}
        _ = v40
    default:
        fmt.Println(v3, v2, v10)
    }
    v10 = v10
    for fi41 := 0; fi41 < 7; fi41++ {
        if (1 >= -1) {
            var v42 int32 = -1
            var v43 int8 = 40
            _ = v42
            _ = v43
        } else {
            for fi44 := 0; fi44 < 4; fi44++ {
                var v45 uint8 = v10
                _ = fi44
                _ = v45
            }
            if true {
                v10 = (v10 ^ v10)
            }
        }
        if (!(!true)) {
            switch (-(fi41)) {
            case 42:
                v46 := GenS7{F0: v3, F1: 246}
                _ = v46
            case 4294967295:
                switch "foo bar" {
                case "hello":
                    v47 := GenS5{F0: 32767, F1: (fi41 / ((724) | 1)), F2: (^(v10))}
                    _ = v47
                case "\n\t":
                    var v48 uint8 = (v10 * 0)
                    _ = v48
                }
            default:
                v49 := GenS7{F0: v3, F1: v3}
                _ = v49
            }
            fmt.Println(v10, v3)
        } else {
            if false {
                var v50 bool = ((("" + (v4 + v1)) != (v4 + "hello")) && (2 > v3))
                var v51 string = ""
                if false {
                    if len(v9) >= 2 {
                        v9 = v9[1:len(v9)]
                    }
                    if (v3 > (v3 >> (((v2 - (v2 ^ (^(v2))))) & 63))) {
                        v11 = (^(fi41))
                        v51 = "hello"
                        for fi52 := 0; fi52 < 2; fi52++ {
                            v53 := GenS5{F0: (fi41 % ((42) | 1)), F1: v11, F2: (v10 ^ (^(v10)))}
                            _ = fi52
                            _ = v53
                        }
                    } else {
                        for fi54 := 0; fi54 < 13; fi54++ {
                            if (1 < 1) {
                                _ = len(v5)
                                v3 = v3
                            } else {
                                switch fi41 {
                                case 581:
                                    var v55 int32 = 42
                                    _ = v55
                                case 42:
                                    v56 := GenS6{F0: 9223372036854775807, F1: 1}
                                    _ = v56
                                }
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            }
                            if v50 {
                                v4 = v4
                                if v50 {
                                    var v57 int8 = 127
                                    if true {
                                        v51 = ("" + "")
                                    } else {
                                        v58 := GenS5{F0: -1, F1: 9223372036854775807, F2: v10}
                                        v10 = v10
                                        if (1.5 >= 1.5) {
                                            v39 = v39
                                            fmt.Println(v1, v11, v11)
                                            if v50 {
                                                v3 = v3
                                                v5[v51] = "hello"
                                                v59 := GenS8{F0: -993, F1: (v11 >= (-(fi54)))}
                                                v3 += (v3 | v3)
                                                fmt.Println(v51)
                                                v60 := GenS5{F0: v11, F1: (fi41 / (((fi41 / (((^(fi54))) | 1))) | 1)), F2: (v10 | 231)}
                                                _ = v59
                                                _ = v60
                                            }
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            if (!(-2.25 < float32((0.0 - 3.14159)))) {
                                                if len(v9) > 0 {
                                                    v9[0] = v11
                                                }
                                                if v50 {
                                                    switch v11 {
                                                    case -9223372036854775808:
                                                        if true {
                                                            v61 := GenS6{F0: 4294967295, F1: 65535}
                                                            _ = v61
                                                        } else {
                                                            var v62 string = v1
                                                            var v63 uint8 = v10
                                                            var v64 uint = (v2 - (v2 % (((v2 % ((65535) | 1))) | 1)))
                                                            _ = v62
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                    case 2147483647:
                                                        var v65 int16 = -1
                                                        _ = v65
                                                    case 9223372036854775807:
                                                        if (v3 <= v3) {
                                                            fmt.Println(v2, v51, v3)
                                                        } else {
                                                            if (v3 == (v3 ^ v3)) {
                                                                switch v4 {
                                                                case "foo bar":
                                                                    switch fi54 {
                                                                    case 128:
                                                                        v66 := GenS3{F0: 2}
                                                                        _ = v66
                                                                    case 9223372036854775807:
                                                                        var v67 float32 = v39
                                                                        _ = v67
                                                                    case 127:
                                                                        v68 := GenS6{F0: v11, F1: 255}
                                                                        _ = v68
                                                                    }
                                                                default:
                                                                    var v69 float64 = 1.5
                                                                    _ = v69
                                                                }
                                                                var v70 uint8 = 255
                                                                _ = v70
                                                            } else {
                                                                var v71 bool = false
                                                                switch v4 {
                                                                case "foo bar":
                                                                    v72 := GenS5{F0: -1, F1: (^(fi54)), F2: 127}
                                                                    _ = v72
                                                                default:
                                                                    if ((v57 == v57) && (true || (!(fi54 >= fi41)))) {
                                                                        if false {
                                                                            switch ((v4 + "a") + (v4 + v1)) {
                                                                            case "\n\t":
                                                                                if (3.14159 >= 3.14159) {
                                                                                    var v73 float32 = v39
                                                                                    if false {
                                                                                        var v74 int = int(v3)
                                                                                        _ = v74
                                                                                    }
                                                                                    switch ((v4 + v51) + "foo bar") {
                                                                                    case "\n\t":
                                                                                        v75 := GenS7{F0: 65535, F1: v3}
                                                                                        _ = v75
                                                                                    }
                                                                                    fmt.Println(v3)
                                                                                    _ = v73
                                                                                } else {
                                                                                    var v76 uint8 = (v10 & 37)
                                                                                    _ = v76
                                                                                }
                                                                            }
                                                                            var v77 int16 = 651
                                                                            _ = len(v9)
                                                                            _ = len(v5)
                                                                            v77 -= (v77 & v77)
                                                                            var v78 uint64 = 4294967295
                                                                            _ = v77
                                                                            _ = v78
                                                                        } else {
                                                                            switch ((v51 + ((((v51 + v51) + "a") + (v51 + v4)) + v51)) + "") {
                                                                            case "foo bar":
                                                                                v79 := v6[2147483648]
                                                                                _ = v79
                                                                            case "\n\t":
                                                                                var v80 int = (^(fi54))
                                                                                _ = v80
                                                                            default:
                                                                                if len(v9) >= 2 {
                                                                                    v9 = v9[1:len(v9)]
                                                                                }
                                                                                fmt.Println(v51, v51)
                                                                                var v81 float32 = ((1.5 + (v39 / v39)) - v39)
                                                                                var v82 string = (v1 + v51)
                                                                                if false {
                                                                                    switch (^(v11)) {
                                                                                    case 32768:
                                                                                        var v83 int16 = 32767
                                                                                        _ = v83
                                                                                    case -1:
                                                                                        var v84 uint = (v2 / (((v2 % ((v2) | 1))) | 1))
                                                                                        _ = v84
                                                                                    default:
                                                                                        switch v4 {
                                                                                        case "a":
                                                                                            fmt.Println(v4)
                                                                                        case "hello":
                                                                                            v85 := GenS6{F0: fi54, F1: uint16(v2)}
                                                                                            _ = v85
                                                                                        case "":
                                                                                            var v86 float64 = 0.0
                                                                                            _ = v86
                                                                                        }
                                                                                    }
                                                                                    if true {
                                                                                        var v87 int16 = 781
                                                                                        _ = v87
                                                                                    }
                                                                                } else {
                                                                                    if v71 {
                                                                                        v10 += 2
                                                                                        var v88 int64 = 9223372036854775807
                                                                                        v81 -= (v81 - (0.1 / v39))
                                                                                        _ = v88
                                                                                    }
                                                                                    fmt.Println(v50, v39, v1)
                                                                                    v89 := GenS2{F0: ((fi54 >> (((v2 << ((v2) & 63))) & 63)) <= v11)}
                                                                                    _ = v89
                                                                                }
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                        }
                                                                        if (!(true && true)) {
                                                                            v11 = (fi54 * v11)
                                                                            var v90 uint8 = 127
                                                                            _ = v90
                                                                        } else {
                                                                            v91 := GenS5{F0: (v11 >> ((v2) & 63)), F1: v11, F2: v10}
                                                                            v11 += int(v57)
                                                                            _ = v91
                                                                        }
                                                                        var v92 int16 = 0
                                                                        switch v1 {
                                                                        case "foo bar":
                                                                            v50 = (!((v2 ^ v2) != (v2 % (((v2 - (v2 + 800))) | 1))))
                                                                        default:
                                                                            v93 := GenS1{F0: (v3 % ((v3) | 1))}
                                                                            _ = v93
                                                                        }
                                                                        _ = len(v9)
                                                                        var v94 uint64 = 18446744073709551615
                                                                        v2 = v2
                                                                        _ = v92
                                                                        _ = v94
                                                                    }
                                                                    switch v11 {
                                                                    case 127:
                                                                        var v95 int32 = 961
                                                                        _ = v95
                                                                    case 32767:
                                                                        fmt.Println(fi54)
                                                                    }
                                                                    v4 = v1
                                                                }
                                                                _ = v71
                                                            }
                                                        }
                                                    default:
                                                        switch (^(fi41)) {
                                                        case 431:
                                                            fmt.Println(v50, v3, fi54)
                                                        case -1:
                                                            switch -823 {
                                                            case 171:
                                                                var v96 string = "\n\t"
                                                                _ = v96
                                                            }
                                                        }
                                                    }
                                                    var v97 int16 = 32767
                                                    var v98 float64 = float64((v39 - -2.25))
                                                    _ = v97
                                                    _ = v98
                                                }
                                                v99 := GenS8{F0: 507, F1: v50}
                                                _ = v99
                                            } else {
                                                _ = len(v5)
                                                var v100 int32 = 1
                                                v101 := v6[1]
                                                v102 := GenS9{F0: (!v50)}
                                                var v103 uint8 = (^(v10))
                                                if v50 {
                                                    if (!(v3 <= v3)) {
                                                        fmt.Println(v2)
                                                        delete(v5, "\n\t")
                                                        var v104 string = v4
                                                        _ = v104
                                                    } else {
                                                        var v105 uint8 = (v10 - uint8(v3))
                                                        v11 -= (fi54 ^ v101)
                                                        v10 -= v105
                                                        _ = v105
                                                    }
                                                }
                                                v106 := GenS10{F0: v103, F1: -1}
                                                if len(v9) > 0 {
                                                    v9[0] = -1
                                                }
                                                _ = v100
                                                _ = v101
                                                _ = v102
                                                _ = v103
                                                _ = v106
                                            }
                                        }
                                        _ = v58
                                    }
                                    var v107 float64 = -2.25
                                    v108 := GenS3{F0: 2}
                                    v6[42] = fi54
                                    _ = v107
                                    _ = v108
                                    _ = v57
                                } else {
                                    v6[2147483648] = -336
                                }
                                fmt.Println(fi41, v1, v4)
                                var v109 uint64 = 4294967295
                                _ = v109
                            } else {
                                v3 = (v3 - (v3 / ((82) | 1)))
                                v39 = -2.25
                            }
                            if v50 {
                                v11 = v11
                                var v110 bool = v50
                                var v111 uint32 = 1
                                _ = v110
                                _ = v111
                            }
                            _ = fi54
                        }
                    }
                    v112 := GenS4{F0: 42}
                    _ = v112
                }
                _ = v50
                _ = v51
            } else {
                if false {
                    var v113 float32 = float32(-2.25)
                    switch v1 {
                    case "":
                        v114 := GenS1{F0: (v3 >> (((v2 | v2)) & 63))}
                        _ = v114
                    case "\n\t":
                        v115 := GenS4{F0: 42}
                        _ = v115
                    default:
                        v5[(v1 + ("\n\t" + ((("\n\t" + ("foo bar" + "hello")) + "\n\t") + "foo bar")))] = "hello"
                    }
                    _ = v113
                }
                v116 := GenS8{F0: -1, F1: true}
                _ = v116
            }
        }
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
