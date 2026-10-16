package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint8; F2 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS7 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS8 struct { F0 int8; F1 uint32; F2 string }
type GenS9 struct { F0 uint; F1 int }
type GenS10 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{}
    var v3 int64 = 1
    var v4 int64 = v3
    for fi5 := 0; fi5 < 10; fi5++ {
        if true {
            for fi6 := 0; fi6 < 6; fi6++ {
                if true {
                    var v7 uint8 = 42
                    var v8 int16 = 32767
                    switch "\n\t" {
                    case "foo bar":
                        var v9 string = ("foo bar" + ((((("" + "hello") + "hello") + "") + "") + ("foo bar" + "foo bar")))
                        _ = v9
                    case "hello":
                        if true {
                            switch (fi6 ^ fi6) {
                            case 127:
                                v3 = (v3 >> ((679) & 63))
                            case 302:
                                if true {
                                    v10 := GenS1{F0: -607, F1: 47, F2: -2147483648}
                                    v7 += (v7 >> ((65535) & 63))
                                    v11 := GenS1{F0: 772, F1: 1, F2: -2147483648}
                                    _ = v10
                                    _ = v11
                                }
                            case 42:
                                fmt.Println(v7, fi6, v4)
                            default:
                                switch "\n\t" {
                                case "":
                                    var v12 int64 = v3
                                    _ = v12
                                case "a":
                                    var v13 float64 = 3.14159
                                    _ = v13
                                default:
                                    v14 := GenS2{F0: ((v3 | (-(v4))) != v3)}
                                    _ = v14
                                }
                            }
                            if false {
                                v4 += v3
                            }
                        }
                    case "":
                        v15 := v1["foo bar"]
                        _ = v15
                    default:
                        if (151 > 2) {
                            if len(v2) > 0 {
                                v2[0] = 149
                            }
                            v8 += 1
                            var v16 uint8 = 255
                            if true {
                                switch ("hello" + "\n\t") {
                                case "\n\t":
                                    var v17 bool = (0 >= 42)
                                    _ = v17
                                default:
                                    delete(v1, "a")
                                }
                            } else {
                                if true {
                                    if (false || true) {
                                        var v18 float64 = (1e10 * 0.1)
                                        _ = v18
                                    } else {
                                        if (!false) {
                                            if ((!true) && false) {
                                                v19 := GenS1{F0: (fi6 / ((fi5) | 1)), F1: 100, F2: 0}
                                                var v20 float64 = float64((0.0 + 3.14159))
                                                switch fi5 {
                                                case -83:
                                                    v2 = append(v2, 478)
                                                case 32767:
                                                    if false {
                                                        v21 := GenS3{F0: 646, F1: v16, F2: (fi5 * fi6)}
                                                        var v22 float32 = float32(0.1)
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                }
                                                _ = v19
                                                _ = v20
                                            } else {
                                                v23 := GenS3{F0: 231, F1: (v7 | 193), F2: fi6}
                                                _ = v23
                                            }
                                        }
                                    }
                                } else {
                                    fmt.Println(fi5, v3)
                                    switch "\n\t" {
                                    case "\n\t":
                                        v3 += v3
                                    }
                                }
                                v3 -= (v4 >> ((2) & 63))
                                var v24 int8 = -84
                                _ = v24
                            }
                            _ = v16
                        }
                    }
                    _ = v7
                    _ = v8
                } else {
                    fmt.Println(v4, v4)
                }
                v1["a"] = true
                v3 += (-(v4))
                _ = fi6
            }
            var v25 int16 = 43
            if (!false) {
                var v26 uint = 4294967295
                switch 128 {
                case 42:
                    v3 = v3
                }
                _ = v26
            }
            _ = v25
        } else {
            var v27 float32 = ((0.0 - 0.0) - (1.5 * (3.14159 + -2.25)))
            _ = len(v2)
            for fi28 := 0; fi28 < 11; fi28++ {
                v29 := GenS3{F0: 255, F1: 2, F2: (fi5 / (((fi5 * fi5)) | 1))}
                switch int(v4) {
                case -2147483648:
                    var v30 uint = 42
                    _ = v30
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                _ = fi28
                _ = v29
            }
            _ = v27
        }
        _ = fi5
    }
    var v31 int8 = 0
    v32 := v1[("foo bar" + ("foo bar" + ""))]
    if (v32 || (!true)) {
        var v33 uint16 = 460
        for fi34 := 0; fi34 < 20; fi34++ {
            v1["\n\t"] = v32
            v32 = v32
            _ = fi34
        }
        v35 := GenS3{F0: 882, F1: uint8(v33), F2: 32768}
        _ = v33
        _ = v35
    }
    if true {
        v36 := GenS4{F0: 1}
        var v37 int8 = v31
        var v38 int64 = v3
        _ = v36
        _ = v37
        _ = v38
    }
    for fi39 := 0; fi39 < 11; fi39++ {
        v40 := GenS5{F0: 42, F1: 42}
        for fi41 := 0; fi41 < 12; fi41++ {
            switch "foo bar" {
            case "":
                if true {
                    v42 := GenS4{F0: 201}
                    v31 = -125
                    var v43 uint = 18446744073709551615
                    _ = v42
                    _ = v43
                } else {
                    fmt.Println(fi39)
                }
            case "hello":
                fmt.Println(v31)
                v44 := GenS4{F0: 1}
                fmt.Println(fi41)
                delete(v1, "a")
                switch 42 {
                case 32767:
                    if (964 == 42) {
                        switch ("\n\t" + "hello") {
                        case "hello":
                            if v32 {
                                if (((1e10 * ((-2.25 / 0.1) - (3.14159 - 1.5))) * 1e10) >= 0.1) {
                                    v45 := GenS2{F0: false}
                                    var v46 uint64 = 1
                                    switch "hello" {
                                    case "a":
                                        var v47 int8 = (v31 * v31)
                                        _ = v47
                                    case "foo bar":
                                        if (3.14159 > 3.14159) {
                                            if v32 {
                                                var v48 float32 = float32(0.0)
                                                _ = v48
                                            } else {
                                                var v49 uint = 4294967295
                                                var v50 bool = (v32 || (false || (32767 < 127)))
                                                v31 -= (v31 % (((v31 >> ((526) & 63))) | 1))
                                                _ = v49
                                                _ = v50
                                            }
                                            switch fi41 {
                                            case 1:
                                                v4 = (v3 - v3)
                                            case 32767:
                                                v31 = v31
                                            }
                                            if v32 {
                                                v31 = 1
                                            } else {
                                                v51 := GenS1{F0: (fi41 >> ((772) & 63)), F1: -128, F2: 1}
                                                _ = v51
                                            }
                                        } else {
                                            v52 := GenS4{F0: 343}
                                            _ = v52
                                        }
                                    default:
                                        var v53 bool = true
                                        _ = v53
                                    }
                                    _ = v45
                                    _ = v46
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if true {
                                    v4 = (-(v4))
                                    var v54 int8 = (^(v31))
                                    if false {
                                        var v55 int = (^(fi41))
                                        var v56 uint = 1
                                        if v32 {
                                            v57 := GenS6{F0: -9223372036854775808, F1: 32767, F2: 127}
                                            if len(v2) > 0 {
                                                v2[0] = 1
                                            }
                                            v58 := v1["foo bar"]
                                            _ = v57
                                            _ = v58
                                        } else {
                                            v59 := GenS7{F0: v4, F1: 492, F2: 2}
                                            fmt.Println(v3)
                                            _ = v59
                                        }
                                        _ = v55
                                        _ = v56
                                    } else {
                                        if (!false) {
                                            var v60 int64 = v4
                                            fmt.Println(v32, v4)
                                            _ = v60
                                        }
                                        if false {
                                            if v32 {
                                                var v61 uint16 = 32767
                                                var v62 uint8 = 175
                                                fmt.Println(v31, v3)
                                                _ = v61
                                                _ = v62
                                            }
                                            var v63 uint32 = 2
                                            _ = v63
                                        } else {
                                            fmt.Println(v32, v54)
                                            _ = len(v1)
                                            var v64 uint64 = 1
                                            _ = v64
                                        }
                                        if v32 {
                                            v54 = v31
                                            if v32 {
                                                var v65 uint8 = 54
                                                _ = v65
                                            } else {
                                                if (true && (!(!((317 != 2147483648) && v32)))) {
                                                    var v66 uint64 = 42
                                                    var v67 uint = 0
                                                    v66 = v66
                                                    var v68 string = "foo bar"
                                                    var v69 uint8 = 0
                                                    _ = v66
                                                    _ = v67
                                                    _ = v68
                                                    _ = v69
                                                } else {
                                                    fmt.Println(v32, v32, fi39)
                                                    v70 := GenS7{F0: v4, F1: 433, F2: 243}
                                                    v71 := v1["foo bar"]
                                                    _ = v70
                                                    _ = v71
                                                }
                                                v72 := v1[(("foo bar" + "") + "hello")]
                                                var v73 bool = true
                                                _ = v72
                                                _ = v73
                                            }
                                            var v74 string = "hello"
                                            if (((0.0 / ((1.5 / 0.1) - (((1.5 * (0.0 * 1e10)) - float64(1.5)) / 1.0))) * (0.1 / -2.25)) > 3.14159) {
                                                var v75 string = "foo bar"
                                                _ = v75
                                            } else {
                                                var v76 float32 = (-2.25 * (1.5 / 1.0))
                                                _ = v76
                                            }
                                            fmt.Println(v31, v32, v3)
                                            v4 = v3
                                            if (!(!v32)) {
                                                v77 := GenS6{F0: (v3 - -291), F1: 477, F2: 42}
                                                var v78 uint64 = 1
                                                _ = v77
                                                _ = v78
                                            } else {
                                                var v79 string = v74
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v4 = (^(v3))
                                                _ = v79
                                            }
                                            _ = v74
                                        }
                                    }
                                    _ = v54
                                } else {
                                    v3 = (v4 * v4)
                                }
                            } else {
                                v2 = append(v2, 0)
                                switch (fi41 % (((fi41 ^ (fi41 - (fi39 / ((fi41) | 1))))) | 1)) {
                                case 389:
                                    fmt.Println(v31, v32, v4)
                                case 678:
                                    v80 := GenS8{F0: (v31 | (-(v31))), F1: 1, F2: "hello"}
                                    _ = v80
                                }
                            }
                        }
                        v3 = 256
                        var v81 uint32 = 0
                        _ = v81
                    }
                case 444:
                    v82 := GenS9{F0: 2, F1: fi41}
                    _ = v82
                default:
                    _ = len(v1)
                }
                if v32 {
                    switch "\n\t" {
                    case "\n\t":
                        if v32 {
                            v4 += (-(v3))
                            var v83 uint32 = 4294967295
                            switch "\n\t" {
                            case "":
                                var v84 float64 = 0.1
                                _ = v84
                            case "foo bar":
                                fmt.Println(fi39)
                            case "hello":
                                var v85 float64 = 3.14159
                                _ = v85
                            default:
                                v86 := GenS7{F0: (v3 + int64(v31)), F1: 4294967295, F2: 33}
                                _ = v86
                            }
                            _ = v83
                        }
                    default:
                        v2 = append(v2, 2)
                    }
                } else {
                    fmt.Println(v4, v31)
                    fmt.Println(v31, v31)
                    v2 = append(v2, 2147483647)
                }
                _ = v44
            case "a":
                switch "foo bar" {
                case "hello":
                    var v87 string = "a"
                    v32 = (!true)
                    var v88 int16 = 909
                    _ = v87
                    _ = v88
                case "\n\t":
                    if true {
                        fmt.Println(fi41, v4, v3)
                    } else {
                        _ = len(v1)
                        if v32 {
                            var v89 int16 = -1
                            fmt.Println(v32, v32)
                            var v90 uint64 = 4294967295
                            _ = v89
                            _ = v90
                        } else {
                            if (255 != 0) {
                                v31 += v31
                            } else {
                                fmt.Println(v3)
                            }
                            v31 -= v31
                            if false {
                                var v91 float32 = 1e10
                                switch (("hello" + (("" + ("" + "")) + "")) + (((((("" + "\n\t") + "hello") + "\n\t") + "\n\t") + "\n\t") + "a")) {
                                case "a":
                                    v3 -= (-(v4))
                                default:
                                    var v92 uint = 2
                                    _ = v92
                                }
                                var v93 int16 = -32768
                                _ = v91
                                _ = v93
                            }
                            if (v32 || (255 <= 42)) {
                                v2 = append(v2, 1)
                            } else {
                                v94 := GenS7{F0: (v3 + (v4 / (((v4 * (v4 << ((1) & 63)))) | 1))), F1: 1, F2: 255}
                                _ = v94
                            }
                            fmt.Println(v32, v4, v32)
                        }
                        fmt.Println(fi41)
                        if (fi39 != (fi41 - (fi41 % (((^(fi41))) | 1)))) {
                            var v95 int64 = 593
                            if (127 == 2) {
                                var v96 string = "a"
                                if v32 {
                                    v96 = v96
                                    if true {
                                        v31 = -42
                                    } else {
                                        v3 -= v4
                                        switch "a" {
                                        case "foo bar":
                                            v31 -= -1
                                        }
                                        var v97 string = "foo bar"
                                        _ = v97
                                    }
                                    if v32 {
                                        v2 = append(v2, 738)
                                    }
                                } else {
                                    if (v32 && (2 != 42)) {
                                        if v32 {
                                            v95 -= v3
                                        }
                                        v2 = append(v2, 42)
                                    } else {
                                        v31 = v31
                                        var v98 uint32 = 2147483647
                                        _ = v98
                                    }
                                    if true {
                                        switch 1 {
                                        case -1:
                                            v1[(("a" + v96) + v96)] = true
                                        }
                                        var v99 float32 = (((0.1 - (1e10 - float32(float64(1e10)))) / (0.0 + (float32(((float64(1e10) * 1e10) + (3.14159 / float64((float32(3.14159) * (-2.25 / (float32((1.5 + 0.1)) - (1.5 / (0.1 * 1e10))))))))) - -2.25))) + 1.5)
                                        _ = v99
                                    }
                                }
                                _ = v96
                            } else {
                                var v100 bool = v32
                                v95 += v4
                                v101 := GenS5{F0: 47, F1: -32768}
                                _ = v100
                                _ = v101
                            }
                            _ = v95
                        }
                    }
                }
            default:
                fmt.Println(v4, v4, fi41)
            }
            var v102 float32 = 3.14159
            _ = fi41
            _ = v102
        }
        _ = fi39
        _ = v40
    }
    for fi103 := 0; fi103 < 18; fi103++ {
        var v104 uint8 = 255
        v105 := GenS5{F0: 4294967295, F1: -32768}
        if (!(-32768 > 0)) {
            switch (fi103 << ((65535) & 63)) {
            case -308:
                v104 -= (v104 >> ((0) & 63))
            default:
                var v106 uint32 = uint32(v104)
                _ = v106
            }
        } else {
            if true {
                v107 := GenS2{F0: v32}
                switch -408 {
                case -96:
                    v4 = (v3 << ((255) & 63))
                case 2147483647:
                    v108 := GenS4{F0: 2147483647}
                    _ = v108
                }
                var v109 float64 = (((0.1 + 0.1) / ((-2.25 / ((0.1 / 0.1) / (0.0 - -2.25))) - (0.1 * float64(1e10)))) / (float64((float32(3.14159) + 0.1)) + -2.25))
                _ = v107
                _ = v109
            } else {
                for fi110 := 0; fi110 < 2; fi110++ {
                    if true {
                        v1["\n\t"] = ((float32(1e10) >= -2.25) && v32)
                    } else {
                        if v32 {
                            fmt.Println(v32, v31, fi110)
                            v4 = (v4 + 9223372036854775807)
                        }
                    }
                    v111 := v1["a"]
                    v112 := GenS5{F0: 18446744073709551615, F1: -32768}
                    _ = fi110
                    _ = v111
                    _ = v112
                }
                var v113 uint64 = 1
                _ = v113
            }
            var v114 uint32 = 42
            v104 += v104
            _ = v114
        }
        _ = fi103
        _ = v104
        _ = v105
    }
    v115 := GenS8{F0: v31, F1: 4294967295, F2: "\n\t"}
    v116 := GenS10{F0: 32767, F1: (-(v31))}
    v32 = false
    switch "a" {
    case "hello":
        var v117 int16 = 1000
        _ = v117
    case "":
        for fi118 := 0; fi118 < 14; fi118++ {
            fmt.Println(v32, v4, v32)
            _ = fi118
        }
    case "foo bar":
        for fi119 := 0; fi119 < 3; fi119++ {
            var v120 int64 = v3
            _ = fi119
            _ = v120
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v115
    _ = v116
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
