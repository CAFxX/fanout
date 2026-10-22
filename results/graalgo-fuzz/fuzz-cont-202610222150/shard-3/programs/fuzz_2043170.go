package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int8; F1 bool }
type GenS11 struct { F0 uint16; F1 uint8 }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 int32; F1 uint8 }
type GenS14 struct { F0 uint }

func main() {
    var v1 int8 = 1
    var v2 uint16 = 255
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    switch "hello" {
    case "hello":
        var v5 int8 = v1
        _ = v5
    case "":
        v6 := []int{54, 127}
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: 4294967295, F2: ((("\n\t" + ("hello" + "hello")) + "foo bar") + "a")}
    if true {
        v8 := []int{}
        fmt.Println(v1, v2)
        fmt.Println(v1, v1)
        _ = v8
    } else {
        v2 = (v2 / ((v2) | 1))
        v9 := []int{int(v1)}
        _ = v9
    }
    if (18446744073709551615 > 18446744073709551615) {
        v1 += v1
        switch "a" {
        case "hello":
            for fi10 := 0; fi10 < 18; fi10++ {
                if false {
                    if (2 >= 469) {
                        v1 = v1
                        v11 := GenS1{F0: 4294967295, F1: -1000, F2: ""}
                        _ = v11
                    } else {
                        v1 += (v1 - 127)
                    }
                } else {
                    switch fi10 {
                    case 9223372036854775807:
                        var v12 uint8 = 71
                        _ = v12
                    }
                    v1 += v1
                }
                var v13 int8 = v1
                if true {
                    var v14 uint = 65535
                    _ = v14
                }
                _ = fi10
                _ = v13
            }
        case "a":
            switch 634 {
            case 1:
                var v15 uint16 = (v2 + v2)
                _ = v15
            case 32767:
                delete(v4, ("" + ("a" + "a")))
            default:
                if true {
                    var v16 float64 = 3.14159
                    if (-2147483648 >= 42) {
                        var v17 bool = false
                        v18 := GenS1{F0: 510, F1: -1, F2: ("" + ("" + "\n\t"))}
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(v1)
                        v19 := []int{199, -541}
                        _ = v19
                    }
                    v20 := v4["\n\t"]
                    _ = v16
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 15; fi21++ {
                        _ = len(v4)
                        if (!((-1 != fi21) && true)) {
                            v1 += 127
                            if (v2 > v2) {
                                if (38 > 255) {
                                    if ((0 <= 971) || (v1 < 100)) {
                                        var v22 int = fi21
                                        _ = v22
                                    } else {
                                        for fi23 := 0; fi23 < 8; fi23++ {
                                            var v24 uint32 = 1
                                            v24 = 4294967295
                                            v25 := GenS1{F0: uint64(v2), F1: fi23, F2: (("hello" + "a") + "hello")}
                                            _ = fi23
                                            _ = v24
                                            _ = v25
                                        }
                                        var v26 int32 = 450
                                        _ = v26
                                    }
                                    if false {
                                        if false {
                                            if true {
                                                var v27 string = ""
                                                _ = v27
                                            } else {
                                                v1 = -49
                                                v3[("foo bar" + "\n\t")] = (("\n\t" + (("" + "hello") + (("hello" + "") + ""))) + "foo bar")
                                                for fi28 := 0; fi28 < 6; fi28++ {
                                                    if (1 != 4294967295) {
                                                        v2 -= 255
                                                        v2 = v2
                                                        v1 -= v1
                                                        fmt.Println(v1)
                                                        if (455 > 4294967295) {
                                                            v2 -= (^(v2))
                                                            v29 := GenS2{F0: 32767, F1: 65535}
                                                            v1 = 1
                                                            _ = v29
                                                        }
                                                    }
                                                    v30 := GenS3{F0: 9223372036854775807}
                                                    switch "\n\t" {
                                                    case "a":
                                                        v31 := GenS1{F0: 0, F1: fi28, F2: "\n\t"}
                                                        v32 := []int{}
                                                        if (0 > 850) {
                                                            if true {
                                                                var v33 float32 = 0.1
                                                                _ = v33
                                                            } else {
                                                                v2 = 2
                                                                var v34 uint32 = 1
                                                                v34 = (v34 << ((998) & 63))
                                                                v35 := GenS2{F0: 1000, F1: 2}
                                                                _ = v34
                                                                _ = v35
                                                            }
                                                            if false {
                                                                var v36 uint64 = 2
                                                                var v37 uint32 = 65535
                                                                _ = v36
                                                                _ = v37
                                                            } else {
                                                                var v38 int8 = (v1 / ((42) | 1))
                                                                _ = v38
                                                            }
                                                        }
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                    v2 = (v2 & v2)
                                                    var v39 int64 = 9223372036854775807
                                                    var v40 uint32 = 0
                                                    _ = fi28
                                                    _ = v30
                                                    _ = v39
                                                    _ = v40
                                                }
                                            }
                                        }
                                    }
                                    v41 := []int{fi21, 1}
                                    _ = v41
                                } else {
                                    var v42 uint8 = 176
                                    _ = v42
                                }
                                v43 := []int{}
                                if true {
                                    for fi44 := 0; fi44 < 14; fi44++ {
                                        fmt.Println(v2, v1)
                                        _ = fi44
                                    }
                                }
                                _ = v43
                            } else {
                                var v45 float64 = -2.25
                                _ = v45
                            }
                        } else {
                            for fi46 := 0; fi46 < 7; fi46++ {
                                var v47 int32 = 42
                                var v48 uint32 = 42
                                _ = fi46
                                _ = v47
                                _ = v48
                            }
                            fmt.Println(v1, v1, fi21)
                        }
                        _ = len(v3)
                        _ = fi21
                    }
                    for fi49 := 0; fi49 < 10; fi49++ {
                        if ((float32((((((float64(0.1) / 1.0) - 3.14159) / 1.0) / 0.1) - (0.0 * float64(((((0.0 + 1e10) - -2.25) + float32(-2.25)) + float32(0.0)))))) + -2.25) == float32(3.14159)) {
                            var v50 bool = true
                            v2 = 609
                            _ = v50
                        }
                        switch (fi49 * fi49) {
                        case -9223372036854775808:
                            for fi51 := 0; fi51 < 8; fi51++ {
                                v52 := GenS2{F0: 973, F1: 2147483647}
                                v1 -= (v1 + v1)
                                var v53 uint8 = 0
                                fmt.Println(fi49, v1)
                                v54 := GenS3{F0: 512}
                                var v55 uint64 = 18446744073709551615
                                _ = fi51
                                _ = v52
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                        }
                        for fi56 := 0; fi56 < 5; fi56++ {
                            v2 -= v2
                            if false {
                                v2 += v2
                                v57 := GenS4{F0: 1, F1: 1, F2: v2}
                                v1 = v1
                                _ = v57
                            } else {
                                switch 42 {
                                case -675:
                                    if ((v1 - v1) >= v1) {
                                        var v58 float64 = 0.0
                                        fmt.Println(fi49, fi56, v2)
                                        v59 := GenS5{F0: fi56}
                                        v60 := GenS5{F0: (fi49 ^ fi49)}
                                        v61 := []int{fi56, fi56}
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                        _ = v61
                                    }
                                case 127:
                                    v1 += v1
                                default:
                                    v2 = 718
                                }
                                v1 += 0
                                if (true && true) {
                                    var v62 uint8 = 255
                                    v62 -= (v62 >> ((0) & 63))
                                    var v63 int8 = (-(v1))
                                    _ = v62
                                    _ = v63
                                }
                            }
                            v64 := GenS4{F0: -572, F1: 83, F2: 65535}
                            _ = fi56
                            _ = v64
                        }
                        _ = fi49
                    }
                }
            }
        }
        v65 := GenS6{F0: 2}
        _ = v65
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2, v2)
    case "a":
        var v66 float64 = (float64(0.1) * -2.25)
        _ = v66
    default:
        v67 := []int{}
        _ = v67
    }
    if (false && (!false)) {
        var v68 int32 = 0
        v68 += v68
        var v69 uint32 = 0
        _ = v68
        _ = v69
    }
    v2 += (v2 >> ((42) & 63))
    if (!false) {
        v1 += -125
    }
    if (1 == 1) {
        for fi70 := 0; fi70 < 20; fi70++ {
            fmt.Println(v2)
            _ = fi70
        }
        var v71 uint64 = 0
        _ = v71
    }
    v72 := []int{}
    if false {
        var v73 uint64 = 0
        v1 += 42
        if false {
            v1 += 0
            switch "\n\t" {
            case "a":
                v2 = v2
            default:
                var v74 uint = 4294967295
                _ = v74
            }
            v75 := GenS7{F0: ""}
            _ = v75
        } else {
            var v76 int32 = 2147483647
            _ = v76
        }
        _ = v73
    } else {
        switch "hello" {
        case "\n\t":
            v1 += 1
        case "":
            var v77 uint = 65535
            _ = v77
        default:
            for fi78 := 0; fi78 < 10; fi78++ {
                switch (fi78 / ((fi78) | 1)) {
                case 4294967295:
                    if false {
                        if true {
                            v72 = append(v72, -977)
                            fmt.Println(v1, fi78)
                        }
                    } else {
                        var v79 int8 = v1
                        switch fi78 {
                        case 127:
                            v79 += (v1 + (v1 >> ((492) & 63)))
                        case 42:
                            v2 = 323
                        }
                        switch ("hello" + "a") {
                        case "a":
                            if true {
                                for fi80 := 0; fi80 < 11; fi80++ {
                                    v81 := GenS8{F0: 128, F1: (0.0 == ((1.5 * 1e10) * 1.5))}
                                    v2 = (^(v2))
                                    _ = fi80
                                    _ = v81
                                }
                                if false {
                                    for fi82 := 0; fi82 < 6; fi82++ {
                                        fmt.Println(v2, fi82, v79)
                                        _ = fi82
                                    }
                                    _ = len(v72)
                                    if (!true) {
                                        switch "a" {
                                        case "a":
                                            v2 = v2
                                        case "\n\t":
                                            v79 = v1
                                        case "hello":
                                            _ = len(v72)
                                        }
                                    }
                                } else {
                                    v1 += -105
                                    var v83 int = fi78
                                    _ = v83
                                }
                            } else {
                                if (fi78 <= fi78) {
                                    if true {
                                        fmt.Println(v2, v1)
                                        var v84 uint64 = 4294967295
                                        _ = v84
                                    } else {
                                        if false {
                                            if false {
                                                if false {
                                                    var v85 string = (("" + "foo bar") + "a")
                                                    _ = v85
                                                } else {
                                                    v2 = 780
                                                }
                                                if (true || (("" + (("" + (((("a" + (("hello" + "") + "")) + "\n\t") + ("hello" + "")) + "a")) + "a")) > "hello")) {
                                                    fmt.Println(v2, v2, fi78)
                                                    if len(v72) > 0 {
                                                        v72[0] = fi78
                                                    }
                                                    v86 := GenS9{F0: "a"}
                                                    _ = v86
                                                }
                                                var v87 int16 = 42
                                                _ = v87
                                            } else {
                                                for fi88 := 0; fi88 < 13; fi88++ {
                                                    v1 = (v79 / ((v79) | 1))
                                                    _ = fi88
                                                }
                                                var v89 uint = 4294967295
                                                var v90 int16 = 1
                                                _ = v89
                                                _ = v90
                                            }
                                        } else {
                                            if ((-2.25 - 1e10) != 1.5) {
                                                var v91 uint8 = 127
                                                var v92 int32 = -1
                                                _ = v91
                                                _ = v92
                                            } else {
                                                v93 := GenS10{F0: v1, F1: (("hello" < "foo bar") || false)}
                                                v1 -= (v1 + 0)
                                                _ = v93
                                            }
                                        }
                                        switch (fi78 << ((65535) & 63)) {
                                        case 268:
                                            if true {
                                                switch ("a" + "hello") {
                                                case "":
                                                    v2 = v2
                                                case "hello":
                                                    switch "foo bar" {
                                                    case "":
                                                        var v94 int8 = (v79 >> ((0) & 63))
                                                        _ = v94
                                                    case "hello":
                                                        switch ("\n\t" + "") {
                                                        case "foo bar":
                                                            for fi95 := 0; fi95 < 9; fi95++ {
                                                                v79 += v79
                                                                v1 = (v79 << ((4294967295) & 63))
                                                                _ = fi95
                                                            }
                                                        case "hello":
                                                            v96 := GenS10{F0: v79, F1: false}
                                                            _ = v96
                                                        case "a":
                                                            switch 660 {
                                                            case 127:
                                                                fmt.Println(v2)
                                                            case 584:
                                                                v97 := GenS2{F0: -340, F1: 2}
                                                                _ = v97
                                                            }
                                                        }
                                                    case "a":
                                                        v98 := v4[("\n\t" + "foo bar")]
                                                        _ = v98
                                                    default:
                                                        var v99 uint8 = 166
                                                        _ = v99
                                                    }
                                                case "\n\t":
                                                    fmt.Println(fi78, v2)
                                                default:
                                                    var v100 bool = (!((128 == 42) && (0 < uint(v2))))
                                                    _ = v100
                                                }
                                                switch fi78 {
                                                case 2147483647:
                                                    _ = len(v3)
                                                case 32767:
                                                    if (true || false) {
                                                        switch (fi78 >> ((2) & 63)) {
                                                        case 675:
                                                            v101 := GenS2{F0: -32768, F1: 2}
                                                            _ = v101
                                                        case -1:
                                                            var v102 string = "a"
                                                            _ = v102
                                                        }
                                                        v1 = (v79 % ((1) | 1))
                                                    } else {
                                                        v103 := GenS11{F0: v2, F1: 0}
                                                        _ = v103
                                                    }
                                                case -9223372036854775808:
                                                    for fi104 := 0; fi104 < 8; fi104++ {
                                                        var v105 string = "\n\t"
                                                        var v106 int32 = 1
                                                        _ = fi104
                                                        _ = v105
                                                        _ = v106
                                                    }
                                                }
                                                v107 := GenS12{F0: 63}
                                                _ = v107
                                            } else {
                                                var v108 uint16 = v2
                                                _ = v108
                                            }
                                        }
                                        for fi109 := 0; fi109 < 14; fi109++ {
                                            var v110 int64 = 9223372036854775807
                                            switch ((("a" + "") + "") + "a") {
                                            case "hello":
                                                var v111 int16 = -32768
                                                _ = v111
                                            case "foo bar":
                                                if false {
                                                    switch (fi109 ^ -2147483648) {
                                                    case 32768:
                                                        fmt.Println(fi78)
                                                    default:
                                                        v110 += -1
                                                    }
                                                } else {
                                                    var v112 int32 = 2147483647
                                                    _ = v112
                                                }
                                            case "":
                                                var v113 uint64 = 42
                                                _ = v113
                                            default:
                                                switch (("foo bar" + "foo bar") + ("a" + ("foo bar" + "foo bar"))) {
                                                case "hello":
                                                    var v114 uint64 = 2
                                                    _ = v114
                                                default:
                                                    var v115 float64 = 3.14159
                                                    _ = v115
                                                }
                                            }
                                            var v116 bool = (750 > 0)
                                            _ = fi109
                                            _ = v110
                                            _ = v116
                                        }
                                    }
                                    v2 = (^(v2))
                                    for fi117 := 0; fi117 < 12; fi117++ {
                                        switch "a" {
                                        case "":
                                            var v118 uint64 = 42
                                            _ = v118
                                        case "\n\t":
                                            if false {
                                                fmt.Println(v2)
                                                v2 = (v2 ^ (v2 * 65535))
                                                var v119 int = (-(fi78))
                                                _ = v119
                                            }
                                        case "hello":
                                            if ((!true) && false) {
                                                fmt.Println(fi117, v2, fi78)
                                            } else {
                                                var v120 float64 = (1e10 / ((-2.25 / (1.5 * 3.14159)) / (1e10 * (1.5 * float64(float32(1e10))))))
                                                _ = len(v3)
                                                var v121 uint8 = 0
                                                _ = v120
                                                _ = v121
                                            }
                                        }
                                        _ = fi117
                                    }
                                }
                            }
                        case "hello":
                            var v122 uint64 = 0
                            _ = v122
                        default:
                            var v123 float32 = (0.1 - 3.14159)
                            _ = v123
                        }
                        _ = v79
                    }
                }
                v1 -= (v1 + (v1 ^ v1))
                for fi124 := 0; fi124 < 2; fi124++ {
                    v2 = (v2 & (^(v2)))
                    v1 -= -75
                    v125 := GenS13{F0: -2147483648, F1: 127}
                    v1 = (v1 & v1)
                    _ = fi124
                    _ = v125
                }
                _ = fi78
            }
        }
        if true {
            for fi126 := 0; fi126 < 15; fi126++ {
                v1 -= v1
                _ = fi126
            }
        }
        var v127 string = ("\n\t" + "foo bar")
        _ = v127
    }
    v128 := GenS14{F0: 94}
    for fi129 := 0; fi129 < 8; fi129++ {
        if false {
            var v130 string = (("foo bar" + ("hello" + (("hello" + ("foo bar" + "\n\t")) + "hello"))) + "")
            v131 := GenS2{F0: 912, F1: 394}
            v2 = (v2 + (v2 ^ 112))
            _ = v130
            _ = v131
        }
        var v132 uint16 = (v2 - 42)
        _ = fi129
        _ = v132
    }
    v1 += -79
    var v133 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v133:", v133)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v72:", len(v72), v72)
    _ = v1
    _ = v128
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v72
}
