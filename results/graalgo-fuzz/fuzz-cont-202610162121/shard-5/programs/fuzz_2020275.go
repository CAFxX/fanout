package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint64; F2 int }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint; F1 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool; F1 uint32; F2 int16 }
type GenS10 struct { F0 int; F1 uint64 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int64; F1 string; F2 bool }
func genFn0() int8 {
    v6 := []int{2147483647, -1, 9223372036854775807}
    var v7 uint64 = 18446744073709551615
    v7 = 1
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint64 = 221
    v2 := make(map[string]uint32)
    v3 := []uint32{497, 2147483647, 4294967295, 0}
    var v4 float32 = 0.0
    var v5 float32 = (v4 - ((float32(1e10) * v4) + v4))
    fmt.Println(v1, v4, v4)
    if true {
        if true {
            _ = len(v2)
            v4 = float32(0.0)
            var v8 uint8 = 80
            _ = v8
        }
        if false {
            var v9 float32 = v4
            _ = v9
        }
    }
    if (-1 <= -116) {
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 uint32 = 0
            v4 = v5
            _ = fi10
            _ = v11
        }
    } else {
        switch 32767 {
        case 127:
            v12 := GenS1{F0: 1}
            _ = v12
        case -2147483648:
            switch "hello" {
            case "foo bar":
                var v13 uint = 1
                _ = v13
            case "a":
                v4 -= 0.1
            case "":
                if true {
                    var v14 uint16 = 879
                    var v15 int16 = 1
                    fmt.Println(v5, v4)
                    _ = v14
                    _ = v15
                }
            default:
                v1 -= v1
            }
        case 4294967295:
            fmt.Println(v5, v1, v5)
        default:
            if len(v3) > 0 {
                v3[0] = 586
            }
        }
        var v16 uint8 = 255
        v17 := GenS2{F0: 2147483647}
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        v19 := GenS3{F0: 789, F1: 65535}
        var v20 bool = true
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 uint = 42
        _ = fi21
        _ = v22
    }
    fmt.Println(v1, v4)
    for fi23 := 0; fi23 < 7; fi23++ {
        if (7 > (v1 & (v1 % (((v1 | 2)) | 1)))) {
            if true {
                if (false || false) {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        switch (fi24 << ((1) & 63)) {
                        case 2147483647:
                            if true {
                                var v25 bool = false
                                if true {
                                    fmt.Println(v1)
                                    v26 := v2[("\n\t" + (("foo bar" + ("foo bar" + "\n\t")) + ((("hello" + "\n\t") + ("" + "\n\t")) + ("" + "hello"))))]
                                    fmt.Println(v26, v25, fi23)
                                    _ = v26
                                } else {
                                    fmt.Println(v25, fi23)
                                    if (((("a" + "\n\t") + "foo bar") + "a") == "") {
                                        fmt.Println(v4, fi24)
                                        v1 = (v1 ^ (^(v1)))
                                        var v27 uint = 4294967295
                                        _ = v27
                                    } else {
                                        v28 := GenS1{F0: 1000}
                                        _ = v28
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    fmt.Println(v1, v5, fi24)
                                    v2["a"] = 303
                                }
                                if ((v5 < 1e10) && (-1 != -729)) {
                                    if true {
                                        v5 -= float32(1e10)
                                        var v29 int8 = genFn0()
                                        _ = len(v2)
                                        _ = v29
                                    }
                                } else {
                                    var v30 bool = false
                                    var v31 float64 = -2.25
                                    _ = v30
                                    _ = v31
                                }
                                v5 = 0.0
                                v5 = v5
                                if v25 {
                                    var v32 uint64 = v1
                                    v4 += (v5 - float32(0.1))
                                    v33 := GenS3{F0: 1000, F1: 127}
                                    _ = v32
                                    _ = v33
                                }
                                _ = v25
                            } else {
                                fmt.Println(fi23, v1)
                                if false {
                                    v1 = v1
                                } else {
                                    v5 = (float32(0.0) * (1e10 + float32(3.14159)))
                                    v1 = 4294967295
                                    v1 = (^(v1))
                                }
                            }
                        case 9223372036854775807:
                            v4 -= v5
                        case -1:
                            v34 := GenS4{F0: 234, F1: 267, F2: fi23}
                            _ = v34
                        default:
                            v35 := GenS1{F0: 32767}
                            _ = v35
                        }
                        _ = fi24
                    }
                    v4 -= 0.0
                } else {
                    for fi36 := 0; fi36 < 7; fi36++ {
                        v37 := GenS5{F0: ((("a" + "\n\t") + "hello") + ""), F1: (v1 % ((v1) | 1))}
                        var v38 string = "hello"
                        v5 -= -2.25
                        _ = fi36
                        _ = v37
                        _ = v38
                    }
                }
                for fi39 := 0; fi39 < 3; fi39++ {
                    delete(v2, "a")
                    if false {
                        if (1 >= genFn0()) {
                            _ = len(v2)
                            v4 += (v5 + (v4 * float32(0.0)))
                        }
                    } else {
                        fmt.Println(v5, v4, fi39)
                        if false {
                            v5 += (v5 * (((v5 / float32(0.1)) / v5) * v4))
                        } else {
                            v1 += 1
                            if (!("hello" != "")) {
                                v2["a"] = 112
                                var v40 string = ""
                                var v41 int16 = -1
                                v41 = v41
                                _ = v40
                                _ = v41
                            } else {
                                if (2 == 4294967295) {
                                    if (255 <= 42) {
                                        v42 := GenS2{F0: 418}
                                        _ = v42
                                    } else {
                                        v4 += 0.0
                                        _ = len(v2)
                                        var v43 uint8 = 2
                                        _ = v43
                                    }
                                    v3 = append(v3, 4294967295)
                                    v5 = (-2.25 * (v5 * v5))
                                    var v44 uint32 = 0
                                    var v45 float64 = ((float64(0.0) * ((1.5 + -2.25) / 0.1)) / 1e10)
                                    _ = v44
                                    _ = v45
                                } else {
                                    v46 := GenS2{F0: -2147483648}
                                    switch "hello" {
                                    case "\n\t":
                                        var v47 uint16 = 255
                                        _ = v47
                                    default:
                                        if true {
                                            if true {
                                                if (true && (32767 <= 907)) {
                                                    var v48 float64 = 0.1
                                                    v3 = append(v3, 801)
                                                    v48 -= float64((v4 - 1e10))
                                                    _ = v48
                                                } else {
                                                    v5 = float32(0.1)
                                                    var v49 uint64 = (^(v1))
                                                    _ = v49
                                                }
                                                v50 := GenS6{F0: v1}
                                                _ = v50
                                            }
                                        }
                                    }
                                    v2[""] = 65535
                                    if ((true || (127 >= -1)) || (255 == 255)) {
                                        if (true || (true && false)) {
                                            var v51 int = (-(fi39))
                                            fmt.Println(v5)
                                            _ = v51
                                        } else {
                                            v52 := GenS3{F0: 397, F1: 42}
                                            if (904 >= (fi23 + fi23)) {
                                                var v53 uint = 2
                                                v1 = uint64(v53)
                                                if true {
                                                    if false {
                                                        var v54 string = ("" + "")
                                                        if false {
                                                            fmt.Println(v53)
                                                            v55 := GenS6{F0: v1}
                                                            _ = v55
                                                        } else {
                                                            v56 := GenS4{F0: 42, F1: 2, F2: 128}
                                                            _ = v56
                                                        }
                                                        switch 2147483647 {
                                                        case -130:
                                                            switch (fi23 + fi23) {
                                                            case 4294967295:
                                                                var v57 bool = ((49 > 42) || (2 != (^(v53))))
                                                                _ = v57
                                                            default:
                                                                if true {
                                                                    v54 = v54
                                                                    v54 = "hello"
                                                                } else {
                                                                    _ = len(v2)
                                                                    fmt.Println(fi39, v1, v54)
                                                                }
                                                            }
                                                        case 32768:
                                                            v4 -= v4
                                                        }
                                                        _ = v54
                                                    } else {
                                                        var v58 float32 = v4
                                                        delete(v2, ((("a" + (((("\n\t" + "") + "\n\t") + "a") + "")) + ((("" + "foo bar") + "a") + (("foo bar" + "foo bar") + "a"))) + "a"))
                                                        _ = v58
                                                    }
                                                    delete(v2, "")
                                                    _ = len(v2)
                                                    v5 += v4
                                                } else {
                                                    v1 += (v1 * v1)
                                                    v5 = v4
                                                    var v59 int = 9223372036854775807
                                                    _ = v59
                                                }
                                                _ = v53
                                            } else {
                                                if true {
                                                    switch fi39 {
                                                    case -351:
                                                        v2["\n\t"] = 1
                                                    default:
                                                        if (true && (false && (1e10 == float32((3.14159 - 1e10))))) {
                                                            fmt.Println(fi23, fi23)
                                                        }
                                                    }
                                                } else {
                                                    var v60 uint8 = 0
                                                    var v61 int64 = 9223372036854775807
                                                    _ = v60
                                                    _ = v61
                                                }
                                                fmt.Println(v1)
                                                var v62 uint16 = 210
                                                _ = v62
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = 65535
                                            }
                                            _ = v52
                                        }
                                        var v63 int = fi39
                                        _ = v63
                                    } else {
                                        var v64 string = "a"
                                        v1 = (v1 ^ (v1 | v1))
                                        _ = v64
                                    }
                                    var v65 uint8 = 236
                                    if false {
                                        v5 += 0.0
                                        v1 = v1
                                        v5 = 3.14159
                                        fmt.Println(v65, v4, fi39)
                                        if true {
                                            var v66 bool = false
                                            switch fi23 {
                                            case 571:
                                                v66 = false
                                            case -1:
                                                v67 := GenS5{F0: "hello", F1: (v1 * (v1 << ((724) & 63)))}
                                                _ = v67
                                            case 1:
                                                v66 = v66
                                            default:
                                                if true {
                                                    var v68 float64 = 3.14159
                                                    var v69 bool = (v1 != (^(v1)))
                                                    _ = v68
                                                    _ = v69
                                                }
                                            }
                                            var v70 bool = false
                                            _ = v66
                                            _ = v70
                                        } else {
                                            if (3.14159 == 1.5) {
                                                v65 += (v65 - (v65 * v65))
                                                v1 = 18446744073709551615
                                            } else {
                                                switch (^(fi39)) {
                                                case -9223372036854775808:
                                                    var v71 float32 = v4
                                                    _ = v71
                                                case 392:
                                                    v2["hello"] = 1
                                                }
                                                switch ("foo bar" + "foo bar") {
                                                case "":
                                                    v65 = 42
                                                    v4 = 0.0
                                                    v65 += 1
                                                default:
                                                    v72 := GenS7{F0: 581, F1: 42}
                                                    _ = v72
                                                }
                                                v73 := GenS5{F0: "hello", F1: 4294967295}
                                                fmt.Println(v4, v4)
                                                if false {
                                                    var v74 uint32 = 238
                                                    var v75 bool = (v74 > v74)
                                                    var v76 string = ""
                                                    _ = v74
                                                    _ = v75
                                                    _ = v76
                                                } else {
                                                    v77 := v2["\n\t"]
                                                    _ = v77
                                                }
                                                v1 -= v1
                                                v4 += -2.25
                                                _ = v73
                                            }
                                            v5 = (v4 + v4)
                                            if ((!(true && ((!false) && false))) || true) {
                                                _ = len(v3)
                                            }
                                        }
                                    }
                                    v65 -= 12
                                    var v78 uint8 = v65
                                    _ = v46
                                    _ = v65
                                    _ = v78
                                }
                                var v79 int16 = 1
                                switch -2147483648 {
                                case -2147483648:
                                    var v80 int16 = (v79 | 32767)
                                    _ = v80
                                case 435:
                                    if len(v3) > 0 {
                                        v3[0] = 2147483648
                                    }
                                case 164:
                                    var v81 int64 = 1
                                    _ = v81
                                }
                                var v82 int8 = 100
                                _ = v79
                                _ = v82
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        v1 += (v1 ^ (v1 & 1))
                    }
                    fmt.Println(fi39, v4, v4)
                    _ = fi39
                }
            }
            if (("foo bar" + "hello") < "hello") {
                v5 = (0.0 * v4)
            }
        } else {
            var v83 uint = 287
            for fi84 := 0; fi84 < 18; fi84++ {
                fmt.Println(v4, fi84, v1)
                _ = fi84
            }
            fmt.Println(fi23)
            _ = v83
        }
        fmt.Println(v1, v5, v1)
        if (352 < 1) {
            switch (fi23 - (fi23 >> ((58) & 63))) {
            case -1:
                v85 := GenS8{F0: 1}
                _ = v85
            default:
                if (!(-530 <= 510)) {
                    var v86 string = "a"
                    _ = v86
                }
            }
            switch (-(fi23)) {
            case 1:
                v87 := GenS8{F0: 1000}
                _ = v87
            case 128:
                for fi88 := 0; fi88 < 11; fi88++ {
                    if true {
                        v4 += v4
                        switch fi23 {
                        case 838:
                            v5 += v4
                        case 2147483647:
                            fmt.Println(v5, v1)
                        case 42:
                            v89 := GenS9{F0: ("\n\t" <= "hello"), F1: 0, F2: -32768}
                            _ = v89
                        }
                    }
                    _ = fi88
                }
            default:
                if true {
                    v90 := GenS10{F0: (fi23 % ((fi23) | 1)), F1: v1}
                    _ = v90
                } else {
                    v5 = 1e10
                    v1 -= v1
                    switch "a" {
                    case "":
                        switch fi23 {
                        case 42:
                            switch "\n\t" {
                            case "foo bar":
                                v4 = 0.1
                            case "":
                                switch (("foo bar" + "hello") + "a") {
                                case "hello":
                                    for fi91 := 0; fi91 < 16; fi91++ {
                                        fmt.Println(fi23, v1, v4)
                                        fmt.Println(v4, v4, v1)
                                        var v92 float64 = (3.14159 * 1.5)
                                        _ = fi91
                                        _ = v92
                                    }
                                case "":
                                    v93 := GenS9{F0: true, F1: 730, F2: 32767}
                                    _ = v93
                                }
                            case "\n\t":
                                fmt.Println(fi23, v1)
                            default:
                                var v94 uint16 = 182
                                _ = v94
                            }
                        default:
                            var v95 int32 = -2147483648
                            _ = v95
                        }
                    case "a":
                        fmt.Println(fi23)
                    default:
                        if (65535 != 0) {
                            if len(v3) > 0 {
                                v3[0] = 169
                            }
                            var v96 float32 = -2.25
                            _ = v96
                        } else {
                            if ((-1 <= 511) || false) {
                                switch (fi23 / (((fi23 | (fi23 << ((0) & 63)))) | 1)) {
                                case 42:
                                    var v97 int8 = -121
                                    _ = v97
                                default:
                                    fmt.Println(v4, v1)
                                }
                                var v98 int = fi23
                                var v99 int64 = -814
                                _ = v98
                                _ = v99
                            }
                        }
                    }
                }
            }
        } else {
            v100 := GenS11{F0: 2147483647}
            _ = v100
        }
        _ = fi23
    }
    switch ("hello" + ("" + "")) {
    case "\n\t":
        var v101 bool = (42 == 1)
        _ = v101
    case "hello":
        fmt.Println(v4, v1)
    default:
        v1 = v1
    }
    switch -424 {
    case 32767:
        fmt.Println(v4)
    default:
        fmt.Println(v1)
    }
    switch 127 {
    case 421:
        if (!true) {
            var v102 bool = (false || true)
            var v103 int8 = -128
            if v102 {
                var v104 int = 906
                _ = v104
            }
            _ = v102
            _ = v103
        } else {
            if true {
                _ = len(v3)
                for fi105 := 0; fi105 < 10; fi105++ {
                    for fi106 := 0; fi106 < 18; fi106++ {
                        var v107 float64 = 3.14159
                        v4 -= v4
                        _ = fi106
                        _ = v107
                    }
                    _ = fi105
                }
            }
            v108 := GenS5{F0: "hello", F1: v1}
            _ = v108
        }
    case 460:
        if (("\n\t" < "foo bar") && (1 <= 42)) {
            var v109 uint64 = v1
            v5 = v4
            _ = v109
        } else {
            for fi110 := 0; fi110 < 1; fi110++ {
                v1 = v1
                _ = fi110
            }
            v111 := GenS9{F0: false, F1: 2147483648, F2: 32767}
            var v112 int32 = -1
            _ = v111
            _ = v112
        }
    case 882:
        for fi113 := 0; fi113 < 3; fi113++ {
            for fi114 := 0; fi114 < 19; fi114++ {
                _ = len(v2)
                var v115 float32 = float32((((1.5 - 0.0) - 1e10) * 0.1))
                v2["\n\t"] = 65535
                _ = fi114
                _ = v115
            }
            _ = fi113
        }
    }
    if ((-2.25 > 3.14159) && false) {
        if (47 == 9) {
            var v116 float32 = v5
            _ = len(v2)
            _ = len(v3)
            _ = v116
        } else {
            v117 := GenS12{F0: 0, F1: "foo bar", F2: false}
            _ = v117
        }
    } else {
        fmt.Println(v1)
        var v118 uint8 = 128
        _ = v118
    }
    v5 += 0.1
    v5 -= v5
    v2["foo bar"] = 2147483648
    var v119 bool = (true && true)
    v120 := GenS1{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v119
    _ = v120
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
