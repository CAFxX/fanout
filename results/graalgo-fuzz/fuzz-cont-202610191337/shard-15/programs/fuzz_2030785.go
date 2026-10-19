package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint64 }
type GenS7 struct { F0 string; F1 int; F2 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int64 }

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{}
    var v3 float64 = 1.5
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    v6 := make(map[int]int)
    var v7 uint32 = (v1 + (v1 >> ((134) & 63)))
    var v8 float32 = ((-2.25 - 0.0) * 0.1)
    v9 := v5["\n\t"]
    v10 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: -175}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v7, v7, v9)
    if false {
        var v11 int = 1
        v8 = 0.1
        if true {
            v9 += (^(v11))
        } else {
            for fi12 := 0; fi12 < 4; fi12++ {
                switch ("" + "a") {
                case "\n\t":
                    v13 := GenS1{F0: 4294967295, F1: v1, F2: 1000}
                    _ = v13
                case "foo bar":
                    if false {
                        v11 = (-(fi12))
                        for fi14 := 0; fi14 < 1; fi14++ {
                            if (18446744073709551615 > 872) {
                                if false {
                                    var v15 uint8 = 42
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v1)
                                    default:
                                        var v16 float32 = float32(v3)
                                        _ = v16
                                    }
                                    _ = v15
                                } else {
                                    v17 := GenS2{F0: 18446744073709551615}
                                    var v18 int = fi12
                                    if (v8 >= 0.0) {
                                        fmt.Println(fi14, v1)
                                        if false {
                                            var v19 int64 = 0
                                            _ = v19
                                        } else {
                                            v20 := GenS2{F0: 827}
                                            var v21 uint64 = 18446744073709551615
                                            _ = v20
                                            _ = v21
                                        }
                                        if true {
                                            v22 := GenS1{F0: 42, F1: v1, F2: -199}
                                            var v23 int8 = 0
                                            fmt.Println(fi14, v23, v8)
                                            v11 += (v11 % (((v9 - (v11 - (v11 & (fi14 | v9))))) | 1))
                                            _ = v22
                                            _ = v23
                                        } else {
                                            v18 = 4294967295
                                        }
                                    }
                                    _ = v17
                                    _ = v18
                                }
                                if true {
                                    var v24 float32 = (0.0 / (((v8 - (0.1 / 1e10)) / v8) * (1.5 * v8)))
                                    _ = v24
                                } else {
                                    var v25 uint16 = 81
                                    _ = v25
                                }
                                v8 = (3.14159 / float32(v3))
                                if true {
                                    v7 = v1
                                    var v26 float32 = ((1e10 * ((v8 + (v8 + (v8 + (0.1 + v8)))) - (1e10 + v8))) - (v8 - 3.14159))
                                    if (!false) {
                                        fmt.Println(v8, fi14, v26)
                                        fmt.Println(fi12)
                                    }
                                    _ = v26
                                } else {
                                    v27 := GenS3{F0: -1, F1: 2147483647}
                                    v6[(v11 << ((31) & 63))] = -27
                                    if true {
                                        fmt.Println(v8, fi12)
                                        v11 = fi14
                                        var v28 uint32 = (v1 + v7)
                                        v29 := GenS4{F0: 1}
                                        _ = v28
                                        _ = v29
                                    } else {
                                        switch (("\n\t" + "foo bar") + ("" + ("" + "a"))) {
                                        case "hello":
                                            if true {
                                                v9 = -363
                                                if len(v2) > 0 {
                                                    v2[0] = 4294967295
                                                }
                                            }
                                        default:
                                            v9 = v11
                                        }
                                    }
                                    var v30 int32 = -957
                                    _ = v27
                                    _ = v30
                                }
                            } else {
                                v9 = 32768
                                if (754 == -1) {
                                    switch -963 {
                                    case 1:
                                        _ = len(v2)
                                    case 32767:
                                        fmt.Println(v3)
                                    case -9223372036854775808:
                                        var v31 uint64 = 35
                                        var v32 uint = 155
                                        var v33 int32 = -978
                                        v4["a"] = ((-261 >= (^(v33))) || ("" > ("hello" + "hello")))
                                        if false {
                                            var v34 uint8 = 92
                                            v35 := GenS2{F0: v31}
                                            var v36 uint32 = v1
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                        }
                                        v37 := GenS2{F0: 248}
                                        if true {
                                            v32 = 42
                                            v38 := GenS5{F0: (^(v31)), F1: 1, F2: int64(v32)}
                                            _ = v38
                                        } else {
                                            if (!(305 == v33)) {
                                                var v39 bool = (86 == 1)
                                                _ = v39
                                            }
                                        }
                                        v33 = -1
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                        _ = v37
                                    default:
                                        v40 := GenS1{F0: 4294967295, F1: (v7 % ((2) | 1)), F2: -276}
                                        _ = v40
                                    }
                                    fmt.Println(v9, fi12, v8)
                                    if true {
                                        v41 := GenS6{F0: 126, F1: 2}
                                        var v42 int8 = 100
                                        v42 -= -68
                                        if true {
                                            switch fi14 {
                                            case 32768:
                                                var v43 uint = 18446744073709551615
                                                _ = v43
                                            case 317:
                                                if (true || false) {
                                                    var v44 int32 = 42
                                                    v45 := GenS4{F0: 0}
                                                    var v46 int16 = 1000
                                                    _ = v44
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    v3 += (v3 - v3)
                                                }
                                            case 128:
                                                if len(v2) > 0 {
                                                    v2[0] = v7
                                                }
                                                var v47 float32 = v8
                                                _ = v47
                                            }
                                            var v48 int8 = v42
                                            v49 := GenS2{F0: 892}
                                            v1 = (^(v7))
                                            fmt.Println(v1, v9)
                                            if false {
                                                if false {
                                                    if false {
                                                        if true {
                                                            switch "hello" {
                                                            case "\n\t":
                                                                var v50 int64 = 1
                                                                _ = v50
                                                            case "hello":
                                                                var v51 int16 = -608
                                                                _ = v51
                                                            case "foo bar":
                                                                if false {
                                                                    v3 += v3
                                                                    v4["a"] = (1 >= 42)
                                                                    v7 = v1
                                                                } else {
                                                                    v2 = append(v2, 42)
                                                                }
                                                                v48 = v42
                                                                v1 = (v7 % ((2) | 1))
                                                            }
                                                            var v52 float64 = ((-2.25 - -2.25) / 3.14159)
                                                            var v53 string = (("foo bar" + ("" + "a")) + ("" + ("a" + "hello")))
                                                            var v54 uint = 4294967295
                                                            if ((v52 * (v52 - v52)) <= (float64(1.5) - float64(v8))) {
                                                                var v55 int = 2147483647
                                                                v56 := GenS1{F0: 1, F1: v1, F2: 104}
                                                                delete(v5, ("" + ("foo bar" + v53)))
                                                                _ = v55
                                                                _ = v56
                                                            } else {
                                                                v2 = append(v2, uint32(v54))
                                                            }
                                                            var v57 float32 = 3.14159
                                                            var v58 int16 = 1000
                                                            switch 42 {
                                                            case 32768:
                                                                v2 = append(v2, (v7 * 2))
                                                            }
                                                            fmt.Println(v57, v3)
                                                            v59 := GenS7{F0: (v53 + v53), F1: (-(v9)), F2: 0}
                                                            v9 = fi14
                                                            _ = v52
                                                            _ = v53
                                                            _ = v54
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                    } else {
                                                        if (-2147483648 <= 975) {
                                                            var v60 uint16 = 65535
                                                            var v61 string = ("\n\t" + "")
                                                            _ = v60
                                                            _ = v61
                                                        }
                                                        var v62 int16 = 42
                                                        v7 = 780
                                                        if (196 == 0) {
                                                            v63 := GenS5{F0: 2, F1: 42, F2: -1}
                                                            v11 += (fi12 - v11)
                                                            _ = v63
                                                        } else {
                                                            var v64 int32 = -1
                                                            if true {
                                                                if len(v2) > 0 {
                                                                    v2[0] = (v7 - v7)
                                                                }
                                                                if false {
                                                                    v65 := GenS5{F0: 2, F1: 127, F2: -275}
                                                                    _ = v65
                                                                }
                                                            }
                                                            v66 := GenS6{F0: 255, F1: 897}
                                                            var v67 string = ""
                                                            var v68 int = v9
                                                            var v69 int16 = v62
                                                            _ = v64
                                                            _ = v66
                                                            _ = v67
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                        _ = v62
                                                    }
                                                }
                                            }
                                            _ = v48
                                            _ = v49
                                        }
                                        _ = v41
                                        _ = v42
                                    }
                                }
                                var v70 int8 = 0
                                if false {
                                    v2 = append(v2, 0)
                                    fmt.Println(v1, fi14, v7)
                                    v2 = append(v2, v7)
                                }
                                if true {
                                    var v71 bool = ((0 > 91) && false)
                                    var v72 int = 32768
                                    _ = v71
                                    _ = v72
                                }
                                _ = v70
                            }
                            if false {
                                v3 = 0.1
                                fmt.Println(v8, v9, v9)
                                var v73 uint64 = 42
                                _ = v73
                            } else {
                                if (true && (((v7 - (v1 | v7)) > 2147483648) || true)) {
                                    var v74 bool = true
                                    _ = v74
                                } else {
                                    var v75 int = fi14
                                    switch "a" {
                                    case "a":
                                        v75 += (^(v11))
                                    }
                                    if (true || (((((((-2.25 / v8) - v8) * float32(v3)) > float32(0.1)) || false) && true) || false)) {
                                        v6[(fi14 >> ((2) & 63))] = -1
                                        var v76 int = (v11 * v9)
                                        switch "foo bar" {
                                        case "":
                                            v9 = 110
                                        }
                                        _ = v76
                                    } else {
                                        var v77 float64 = (v3 - 3.14159)
                                        _ = v77
                                    }
                                    _ = v75
                                }
                                switch 4294967295 {
                                case 128:
                                    if (0.1 > v8) {
                                        var v78 float64 = v3
                                        v7 -= 861
                                        _ = v78
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if (827 == 1) {
                                            if false {
                                                var v79 uint8 = 0
                                                _ = v79
                                            }
                                            var v80 uint64 = 1
                                            _ = v80
                                        } else {
                                            switch 1 {
                                            case 0:
                                                v9 = 128
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if false {
                                                        v7 = (v1 ^ 0)
                                                        v81 := GenS7{F0: "", F1: fi14, F2: -45}
                                                        v7 = (v1 / (((v7 - v1)) | 1))
                                                        v6[-848] = fi12
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v82 := v4[("\n\t" + "foo bar")]
                                                            _ = v82
                                                        default:
                                                            switch "hello" {
                                                            case "a":
                                                                v83 := GenS3{F0: -2147483648, F1: 0}
                                                                var v84 float32 = 0.1
                                                                var v85 uint32 = 0
                                                                _ = v83
                                                                _ = v84
                                                                _ = v85
                                                            case "\n\t":
                                                                v86 := GenS5{F0: 42, F1: 127, F2: 0}
                                                                _ = v86
                                                            }
                                                        }
                                                        v87 := GenS4{F0: 18446744073709551615}
                                                        _ = v81
                                                        _ = v87
                                                    } else {
                                                        switch ("a" + ("foo bar" + "a")) {
                                                        case "foo bar":
                                                            var v88 int64 = -1
                                                            _ = v88
                                                        case "hello":
                                                            v1 = (v7 >> ((1) & 63))
                                                        case "":
                                                            v3 += (v3 - -2.25)
                                                        default:
                                                            var v89 int16 = -32768
                                                            _ = v89
                                                        }
                                                        var v90 uint32 = (v7 * (v1 + 4))
                                                        v91 := v5[("" + "foo bar")]
                                                        _ = v90
                                                        _ = v91
                                                    }
                                                case "hello":
                                                    fmt.Println(v9, v1)
                                                    var v92 uint8 = 0
                                                    var v93 uint32 = 1
                                                    _ = v92
                                                    _ = v93
                                                default:
                                                    v94 := GenS4{F0: 2}
                                                    v95 := GenS3{F0: 0, F1: 42}
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if ((9223372036854775807 >= -35) || (true || (v8 >= 1.5))) {
                                                            v96 := v6[v11]
                                                            v97 := GenS3{F0: -2147483648, F1: 0}
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                    case "foo bar":
                                                        v11 = fi14
                                                    default:
                                                        fmt.Println(v1, v8)
                                                        v98 := GenS5{F0: 2, F1: 2, F2: 0}
                                                        var v99 uint32 = (^(v1))
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                    _ = v94
                                                    _ = v95
                                                }
                                            case 316:
                                                if true {
                                                    fmt.Println(v7, v11, v7)
                                                }
                                            case 127:
                                                var v100 float64 = ((float64(0.0) - (v3 - -2.25)) + 3.14159)
                                                if false {
                                                    v1 += (v7 & v1)
                                                }
                                                if false {
                                                    if true {
                                                        if (false && ((0.0 == (-2.25 * float32((v3 - v3)))) && false)) {
                                                            v101 := GenS4{F0: 822}
                                                            switch (v9 >> ((1) & 63)) {
                                                            case 974:
                                                                v102 := GenS3{F0: -56, F1: 0}
                                                                _ = v102
                                                            default:
                                                                switch v9 {
                                                                case 0:
                                                                    var v103 uint64 = 4294967295
                                                                    _ = v103
                                                                case 986:
                                                                    _ = len(v5)
                                                                }
                                                            }
                                                            _ = v101
                                                        }
                                                        v104 := GenS8{F0: 1}
                                                        fmt.Println(v9, v11)
                                                        _ = v104
                                                    } else {
                                                        v8 -= 0.0
                                                        fmt.Println(v11, v7, v100)
                                                        v9 = v11
                                                        v11 = v11
                                                        switch ("hello" + "") {
                                                        case "":
                                                            if true {
                                                                v3 -= (((3.14159 - 0.0) * v100) / (v100 + 0.0))
                                                                v4["\n\t"] = (682 >= -156)
                                                                switch (-(fi14)) {
                                                                case 32768:
                                                                    v100 -= (-2.25 * -2.25)
                                                                default:
                                                                    v8 = (-2.25 + 3.14159)
                                                                }
                                                            }
                                                        default:
                                                            var v105 bool = (!(32767 >= -32768))
                                                            _ = v105
                                                        }
                                                    }
                                                    var v106 uint16 = 1
                                                    v1 += (v7 | (v1 << ((443) & 63)))
                                                    _ = v106
                                                } else {
                                                    v107 := GenS9{F0: -1}
                                                    _ = v107
                                                }
                                                _ = v100
                                            default:
                                                v108 := GenS7{F0: "a", F1: v11, F2: 100}
                                                _ = v108
                                            }
                                            if (false && false) {
                                                switch "hello" {
                                                case "a":
                                                    v11 = (fi14 % ((v9) | 1))
                                                case "foo bar":
                                                    v2 = append(v2, v7)
                                                default:
                                                    fmt.Println(v7, v7, v1)
                                                }
                                                var v109 int = (^(fi12))
                                                switch (^(v109)) {
                                                case 619:
                                                    _ = len(v4)
                                                }
                                                var v110 uint64 = 0
                                                var v111 uint64 = (v110 | v110)
                                                _ = v109
                                                _ = v110
                                                _ = v111
                                            } else {
                                                fmt.Println(fi12, v8)
                                                if (("\n\t" >= ("foo bar" + "a")) || (0 > 0)) {
                                                    v112 := v6[fi14]
                                                    _ = v112
                                                } else {
                                                    var v113 int = v9
                                                    v114 := GenS4{F0: 42}
                                                    fmt.Println(v1)
                                                    _ = v113
                                                    _ = v114
                                                }
                                                v115 := GenS2{F0: 4294967295}
                                                v8 = v8
                                                if false {
                                                    delete(v4, "hello")
                                                    v7 = v1
                                                }
                                                v3 = -2.25
                                                _ = v115
                                            }
                                            v116 := GenS2{F0: 18446744073709551615}
                                            if false {
                                                v7 -= (^(v7))
                                                v117 := v4[("" + "foo bar")]
                                                v117 = v117
                                                _ = v117
                                            }
                                            var v118 uint8 = 2
                                            switch fi14 {
                                            case 128:
                                                v6[(fi12 >> ((2) & 63))] = (-(v11))
                                            case -1:
                                                v6[(-(v9))] = (fi12 + v11)
                                            default:
                                                var v119 float32 = 1e10
                                                _ = v119
                                            }
                                            _ = v116
                                            _ = v118
                                        }
                                        v120 := GenS9{F0: 42}
                                        _ = v120
                                    }
                                case 45:
                                    fmt.Println(v7, v3, v9)
                                }
                                v121 := GenS2{F0: 2}
                                _ = v121
                            }
                            _ = fi14
                        }
                        fmt.Println(v8)
                    }
                case "hello":
                    v8 = 1.5
                }
                _ = fi12
            }
            v9 = 2147483647
        }
        _ = v11
    }
    v9 = v9
    _ = len(v2)
    if true {
        v1 -= v1
    } else {
        var v122 uint32 = (v7 * 4294967295)
        v7 += v1
        _ = v122
    }
    for fi123 := 0; fi123 < 20; fi123++ {
        v3 = float64(0.0)
        v9 = (-(fi123))
        var v124 uint16 = 5
        _ = fi123
        _ = v124
    }
    for fi125 := 0; fi125 < 6; fi125++ {
        delete(v6, fi125)
        _ = fi125
    }
    v126 := GenS2{F0: 18446744073709551615}
    v3 = (1.5 + 0.1)
    switch 166 {
    case 1:
        fmt.Println(v1)
    case 9223372036854775807:
        v9 = (v9 + (v9 & v9))
    default:
        fmt.Println(v7)
    }
    v8 = float32(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
