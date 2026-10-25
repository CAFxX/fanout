package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int8; F1 uint }
type GenS13 struct { F0 uint16; F1 int64 }
type GenS14 struct { F0 uint16; F1 uint }
type GenS15 struct { F0 uint16; F1 uint64 }
type GenS16 struct { F0 int32; F1 int; F2 int32 }
type GenS17 struct { F0 int8; F1 uint }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 uint64; F1 bool }
type GenS20 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS21 struct { F0 string }
type GenS22 struct { F0 uint32; F1 int32 }
type GenS23 struct { F0 int16; F1 uint8; F2 int32 }
type GenS24 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS25 struct { F0 uint64 }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 uint16 }
type GenS28 struct { F0 int16; F1 int }
type GenS29 struct { F0 int8; F1 bool; F2 int }
func genFn0() uint {
    var v5 uint16 = 42
    fmt.Println(v5)
    if true {
        fmt.Println(v5)
        v6 := GenS1{F0: -93}
        var v7 int16 = -324
        _ = v6
        _ = v7
    } else {
        var v8 string = ((("a" + ("foo bar" + (("foo bar" + "a") + "hello"))) + (("hello" + ("hello" + ("" + "a"))) + "hello")) + (("foo bar" + "hello") + "\n\t"))
        _ = v8
    }
    _ = v5
    return 42
}

func main() {
    var v1 int8 = 127
    v2 := make(map[int]uint32)
    var v3 uint = 18446744073709551615
    v4 := make(map[string]uint32)
    v9 := GenS1{F0: -1}
    delete(v4, "a")
    v3 = (v3 / (((v3 - 65535)) | 1))
    var v10 uint32 = 2147483648
    var v11 uint16 = 42
    v12 := GenS1{F0: (v1 << ((790) & 63))}
    fmt.Println(v3, v3, v3)
    if false {
        v13 := []int{128, 4294967295}
        switch 9223372036854775807 {
        case 32767:
            if (4294967295 != (v3 * v3)) {
                if len(v13) > 0 {
                    v13[0] = 9223372036854775807
                }
                if (!(true || false)) {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        v13 = append(v13, fi14)
                        var v15 uint = genFn0()
                        _ = fi14
                        _ = v15
                    }
                    v16 := GenS1{F0: v1}
                    _ = v16
                } else {
                    fmt.Println(v11, v10)
                    var v17 uint = v3
                    if false {
                        switch ("a" + "foo bar") {
                        case "foo bar":
                            for fi18 := 0; fi18 < 16; fi18++ {
                                v13 = append(v13, 9223372036854775807)
                                if true {
                                    v17 += (v3 - (v17 & 1))
                                    delete(v2, 4294967295)
                                    v3 -= (v3 - (^(v3)))
                                } else {
                                    var v19 uint32 = v10
                                    _ = v19
                                }
                                v11 -= uint16(v17)
                                _ = fi18
                            }
                        case "":
                            var v20 int8 = v1
                            _ = v20
                        case "hello":
                            var v21 uint32 = (^(v10))
                            _ = v21
                        }
                    } else {
                        if true {
                            for fi22 := 0; fi22 < 4; fi22++ {
                                v1 += (v1 + v1)
                                var v23 int16 = -539
                                _ = fi22
                                _ = v23
                            }
                            var v24 int64 = 9223372036854775807
                            _ = v24
                        } else {
                            if false {
                                v2[int(v10)] = v10
                                fmt.Println(v11, v17)
                            } else {
                                fmt.Println(v11, v1, v3)
                                v25 := GenS1{F0: v1}
                                fmt.Println(v1)
                                _ = v25
                            }
                        }
                    }
                    _ = v17
                }
            }
        case 627:
            _ = len(v2)
        case 9223372036854775807:
            v26 := GenS2{F0: 42, F1: 1, F2: "a"}
            _ = v26
        }
        var v27 uint16 = (v11 >> ((v3) & 63))
        _ = v13
        _ = v27
    }
    switch "\n\t" {
    case "":
        for fi28 := 0; fi28 < 18; fi28++ {
            if false {
                var v29 int64 = int64(fi28)
                v30 := []int{}
                switch fi28 {
                case 9223372036854775807:
                    var v31 int64 = (v29 + v29)
                    _ = v31
                }
                _ = v29
                _ = v30
            }
            v32 := []int{}
            if true {
                v10 -= v10
                for fi33 := 0; fi33 < 13; fi33++ {
                    fmt.Println(v11, fi28, fi28)
                    _ = fi33
                }
                v34 := GenS3{F0: 42, F1: v11, F2: 36}
                _ = v34
            } else {
                for fi35 := 0; fi35 < 19; fi35++ {
                    fmt.Println(v1, v10, fi28)
                    _ = fi35
                }
                v36 := GenS1{F0: (v1 & v1)}
                _ = v36
            }
            _ = fi28
            _ = v32
        }
    default:
        if (true && false) {
            switch 4294967295 {
            case 4294967295:
                var v37 float64 = -2.25
                _ = v37
            case 32767:
                switch int(v11) {
                case -671:
                    if (v10 > v10) {
                        switch "hello" {
                        case "hello":
                            if (53 >= v1) {
                                v38 := GenS4{F0: 42, F1: 9223372036854775807}
                                v1 = (v1 / (((v1 / (((v1 ^ (v1 % ((v1) | 1)))) | 1))) | 1))
                                _ = v38
                            } else {
                                v10 += 4294967295
                            }
                        }
                        if (!true) {
                            if true {
                                if true {
                                    if (3.14159 == 1e10) {
                                        v39 := GenS5{F0: 0}
                                        if (4294967295 >= 42) {
                                            switch -2147483648 {
                                            case 2147483647:
                                                if false {
                                                    var v40 uint = (v3 >> ((genFn0()) & 63))
                                                    _ = v40
                                                } else {
                                                    v10 = v10
                                                    delete(v2, 552)
                                                }
                                            case 1:
                                                v41 := v2[1]
                                                _ = v41
                                            }
                                        }
                                        v1 += v1
                                        _ = v39
                                    } else {
                                        fmt.Println(v11, v1, v1)
                                    }
                                    var v42 bool = true
                                    fmt.Println(v1, v42, v42)
                                    _ = v42
                                }
                            } else {
                                v10 -= (v10 ^ (v10 ^ (v10 + 2147483648)))
                            }
                            var v43 int64 = -1
                            _ = v43
                        } else {
                            v44 := GenS6{F0: v3, F1: 1, F2: 0}
                            for fi45 := 0; fi45 < 17; fi45++ {
                                v46 := []int{}
                                v10 = 2
                                if true {
                                    v10 -= 65535
                                } else {
                                    v47 := GenS7{F0: 255, F1: 2147483647}
                                    switch ("hello" + "foo bar") {
                                    case "hello":
                                        var v48 uint16 = 32767
                                        _ = v48
                                    case "":
                                        var v49 int16 = -1
                                        _ = v49
                                    case "\n\t":
                                        if false {
                                            switch "foo bar" {
                                            case "hello":
                                                for fi50 := 0; fi50 < 14; fi50++ {
                                                    if (1 < 42) {
                                                        if false {
                                                            v3 = (^(v3))
                                                            var v51 uint16 = v11
                                                            _ = v51
                                                        } else {
                                                            v52 := GenS7{F0: 61, F1: 247}
                                                            switch fi45 {
                                                            case 42:
                                                                var v53 float32 = 0.1
                                                                _ = v53
                                                            case 4294967295:
                                                                if true {
                                                                    v10 -= 2147483647
                                                                    var v54 uint32 = v10
                                                                    v10 = (v54 << (((^(v3))) & 63))
                                                                    _ = v54
                                                                } else {
                                                                    if len(v46) > 0 {
                                                                        v46[0] = 4294967295
                                                                    }
                                                                    v55 := GenS8{F0: 42}
                                                                    var v56 uint16 = v11
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                            case 32767:
                                                                v57 := GenS1{F0: (v1 ^ (v1 % ((v1) | 1)))}
                                                                _ = v57
                                                            }
                                                            switch (fi45 + fi45) {
                                                            case -489:
                                                                v58 := GenS2{F0: 96, F1: 4294967295, F2: ("" + "foo bar")}
                                                                _ = v58
                                                            case 0:
                                                                switch 2147483647 {
                                                                case 9223372036854775807:
                                                                    v59 := GenS9{F0: v1}
                                                                    _ = v59
                                                                case 42:
                                                                    v3 -= (v3 << ((v3) & 63))
                                                                default:
                                                                    v11 = (v11 + (v11 / ((65535) | 1)))
                                                                }
                                                            case -83:
                                                                v60 := GenS3{F0: 1, F1: (v11 % ((v11) | 1)), F2: 127}
                                                                _ = v60
                                                            }
                                                            _ = v52
                                                        }
                                                        var v61 uint32 = (v10 << ((2) & 63))
                                                        _ = v61
                                                    }
                                                    var v62 float32 = 1.5
                                                    _ = fi50
                                                    _ = v62
                                                }
                                            default:
                                                var v63 uint = (^(v3))
                                                _ = v63
                                            }
                                            var v64 int64 = -1
                                            var v65 uint8 = 0
                                            _ = v64
                                            _ = v65
                                        } else {
                                            fmt.Println(v3, v3, fi45)
                                            if (((3.14159 + 0.1) / 1.5) > 0.0) {
                                                var v66 uint = v3
                                                v66 += (v66 / ((v66) | 1))
                                                switch (fi45 / (((^(fi45))) | 1)) {
                                                case -290:
                                                    switch (fi45 ^ fi45) {
                                                    case -239:
                                                        v66 -= v3
                                                    case 9223372036854775807:
                                                        if false {
                                                            var v67 bool = false
                                                            for fi68 := 0; fi68 < 10; fi68++ {
                                                                var v69 int8 = v1
                                                                var v70 uint = v3
                                                                _ = len(v4)
                                                                _ = fi68
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            var v71 string = "a"
                                                            _ = v67
                                                            _ = v71
                                                        }
                                                    case 0:
                                                        v72 := GenS10{F0: 0}
                                                        _ = v72
                                                    }
                                                default:
                                                    fmt.Println(fi45, fi45)
                                                }
                                                _ = v66
                                            } else {
                                                var v73 int16 = -663
                                                v1 -= 0
                                                for fi74 := 0; fi74 < 5; fi74++ {
                                                    var v75 float32 = 0.0
                                                    if (!true) {
                                                        if ((("a" + (("\n\t" + "hello") + "hello")) + "hello") < "foo bar") {
                                                            fmt.Println(v1, fi45, fi45)
                                                            v11 -= (v11 ^ 1)
                                                            v73 = v73
                                                        } else {
                                                            if ((-2.25 + 0.1) != (1e10 / -2.25)) {
                                                                var v76 float32 = v75
                                                                var v77 int32 = -618
                                                                _ = v76
                                                                _ = v77
                                                            }
                                                            fmt.Println(v11, v75)
                                                            var v78 uint8 = 184
                                                            _ = v78
                                                        }
                                                        v79 := GenS11{F0: ("" < ("\n\t" + ("" + "\n\t")))}
                                                        _ = v79
                                                    } else {
                                                        if false {
                                                            v73 = (v73 * (v73 >> ((v3) & 63)))
                                                            var v80 float64 = (3.14159 / 1.5)
                                                            var v81 int32 = 2147483647
                                                            _ = v80
                                                            _ = v81
                                                        } else {
                                                            if true {
                                                                if (genFn0() < genFn0()) {
                                                                    if true {
                                                                        v75 += (3.14159 * v75)
                                                                        v82 := GenS12{F0: v1, F1: v3}
                                                                        _ = v82
                                                                    }
                                                                    var v83 uint = 183
                                                                    switch (("" + "a") + "a") {
                                                                    case "hello":
                                                                        var v84 int32 = 0
                                                                        _ = v84
                                                                    case "foo bar":
                                                                        if true {
                                                                            if false {
                                                                                v4["\n\t"] = v10
                                                                            } else {
                                                                                v75 = v75
                                                                                var v85 int16 = v73
                                                                                _ = v85
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v83
                                                                }
                                                                switch ("foo bar" + "\n\t") {
                                                                case "":
                                                                    var v86 int16 = (v73 * -570)
                                                                    _ = v86
                                                                case "hello":
                                                                    v11 += v11
                                                                case "\n\t":
                                                                    v11 = 32767
                                                                default:
                                                                    var v87 uint16 = (v11 % ((v11) | 1))
                                                                    _ = v87
                                                                }
                                                                _ = len(v46)
                                                            }
                                                            switch (fi45 / (((fi45 ^ int(v1))) | 1)) {
                                                            case -9223372036854775808:
                                                                var v88 int8 = (v1 << (((v3 << ((65535) & 63))) & 63))
                                                                _ = v88
                                                            default:
                                                                if false {
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        v3 = v3
                                                                    default:
                                                                        v89 := GenS3{F0: -353, F1: (^(v11)), F2: 128}
                                                                        _ = v89
                                                                    }
                                                                    if (0 > -2147483648) {
                                                                        v4["\n\t"] = v10
                                                                        if true {
                                                                            v46 = append(v46, fi45)
                                                                            var v90 uint32 = 2
                                                                            v91 := GenS10{F0: 172}
                                                                            _ = v90
                                                                            _ = v91
                                                                        }
                                                                        if (702 != 752) {
                                                                            v92 := v4[""]
                                                                            if (42 == 9223372036854775807) {
                                                                                var v93 int8 = (v1 - v1)
                                                                                var v94 bool = true
                                                                                var v95 int32 = 42
                                                                                _ = v93
                                                                                _ = v94
                                                                                _ = v95
                                                                            } else {
                                                                                v96 := GenS6{F0: v3, F1: 163, F2: 1}
                                                                                var v97 uint8 = 128
                                                                                _ = v96
                                                                                _ = v97
                                                                            }
                                                                            switch -2147483648 {
                                                                            case -154:
                                                                                switch 9223372036854775807 {
                                                                                case -1:
                                                                                    var v98 string = ("hello" + "foo bar")
                                                                                    _ = v98
                                                                                }
                                                                            case 42:
                                                                                v75 = v75
                                                                            }
                                                                            _ = v92
                                                                        } else {
                                                                            var v99 uint16 = 65535
                                                                            v99 = 255
                                                                            v11 -= uint16(v10)
                                                                            switch (^(fi74)) {
                                                                            case 105:
                                                                                if false {
                                                                                    v73 = -865
                                                                                } else {
                                                                                    v10 = uint32(v3)
                                                                                    var v100 string = ""
                                                                                    switch ((v100 + v100) + (((v100 + (v100 + v100)) + ("foo bar" + v100)) + "hello")) {
                                                                                    case "a":
                                                                                        var v101 int = 710
                                                                                        _ = v101
                                                                                    case "hello":
                                                                                        if false {
                                                                                            fmt.Println(fi74, v73)
                                                                                            var v102 uint16 = 42
                                                                                            _ = v102
                                                                                        } else {
                                                                                            v103 := GenS1{F0: 42}
                                                                                            fmt.Println(v10)
                                                                                            if false {
                                                                                                delete(v2, int(v11))
                                                                                                v104 := GenS13{F0: 2, F1: 0}
                                                                                                _ = v104
                                                                                            }
                                                                                            fmt.Println(fi74, v100, fi45)
                                                                                            _ = v103
                                                                                        }
                                                                                        if true {
                                                                                            v11 = (v11 / (((v99 + uint16(v10))) | 1))
                                                                                        }
                                                                                        var v105 float32 = v75
                                                                                        _ = v105
                                                                                    }
                                                                                    _ = v100
                                                                                }
                                                                            }
                                                                            _ = v99
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        switch "" {
                                                        case "\n\t":
                                                            v75 -= float32(-2.25)
                                                        default:
                                                            if ((v75 >= ((float32(float64(v75)) / (((v75 * v75) * ((v75 * float32(0.1)) - v75)) * (v75 - v75))) + float32(0.1))) && false) {
                                                                if true {
                                                                    if len(v46) >= 2 {
                                                                        v46 = v46[1:len(v46)]
                                                                    }
                                                                    v3 = v3
                                                                    if len(v46) >= 2 {
                                                                        v46 = v46[1:len(v46)]
                                                                    }
                                                                }
                                                                var v106 float32 = v75
                                                                var v107 uint8 = 42
                                                                _ = v106
                                                                _ = v107
                                                            } else {
                                                                var v108 int32 = 106
                                                                _ = v108
                                                            }
                                                        }
                                                    }
                                                    v1 += (-(v1))
                                                    v109 := GenS9{F0: v1}
                                                    _ = fi74
                                                    _ = v109
                                                    _ = v75
                                                }
                                                _ = v73
                                            }
                                            v1 = (v1 % ((v1) | 1))
                                        }
                                    }
                                    v110 := GenS5{F0: 42}
                                    _ = v110
                                    _ = v47
                                }
                                _ = fi45
                                _ = v46
                            }
                            switch ("\n\t" + "\n\t") {
                            case "":
                                fmt.Println(v11, v3)
                            default:
                                v111 := GenS11{F0: true}
                                _ = v111
                            }
                            _ = v44
                        }
                        v10 = (v10 << ((462) & 63))
                    } else {
                        v11 += (v11 >> ((genFn0()) & 63))
                        v3 = 314
                    }
                default:
                    var v112 float64 = -2.25
                    _ = v112
                }
            default:
                var v113 int8 = int8(v11)
                _ = v113
            }
        }
    }
    delete(v4, ("\n\t" + "a"))
    var v114 uint16 = uint16(v3)
    if true {
        var v115 string = ("" + ("foo bar" + ((("\n\t" + "foo bar") + "foo bar") + "a")))
        if false {
            switch "hello" {
            case "a":
                var v116 uint32 = (^(v10))
                _ = v116
            case "\n\t":
                if false {
                    switch -1 {
                    case -9223372036854775808:
                        v117 := GenS3{F0: -2147483648, F1: 0, F2: 247}
                        _ = v117
                    default:
                        if (!false) {
                            var v118 uint64 = 1
                            if true {
                                v1 += 100
                            }
                            _ = v118
                        } else {
                            for fi119 := 0; fi119 < 8; fi119++ {
                                var v120 int32 = 42
                                var v121 int8 = (^(v1))
                                _ = fi119
                                _ = v120
                                _ = v121
                            }
                        }
                    }
                    var v122 string = "\n\t"
                    _ = v122
                } else {
                    if false {
                        if ((!(!(9223372036854775807 == -9223372036854775808))) || true) {
                            v123 := []int{1, 127}
                            delete(v4, ("a" + ""))
                            _ = v123
                        } else {
                            if (!false) {
                                v10 = (^(v10))
                                v124 := []int{4294967295, 127, 595}
                                for fi125 := 0; fi125 < 7; fi125++ {
                                    v2[320] = 2147483648
                                    switch ((v115 + "foo bar") + v115) {
                                    case "a":
                                        if (uint(v10) <= v3) {
                                            var v126 float64 = (-2.25 + 0.0)
                                            switch v115 {
                                            case "\n\t":
                                                if true {
                                                    fmt.Println(fi125, v114, v3)
                                                } else {
                                                    for fi127 := 0; fi127 < 16; fi127++ {
                                                        v128 := GenS8{F0: 27}
                                                        v126 -= float64(1e10)
                                                        v129 := GenS11{F0: true}
                                                        _ = fi127
                                                        _ = v128
                                                        _ = v129
                                                    }
                                                    v130 := GenS14{F0: (v114 - (v11 << ((1) & 63))), F1: (^(v3))}
                                                    _ = v130
                                                }
                                            default:
                                                if len(v124) >= 2 {
                                                    v124 = v124[1:len(v124)]
                                                }
                                            }
                                            _ = v126
                                        } else {
                                            var v131 uint8 = 42
                                            var v132 string = v115
                                            _ = v131
                                            _ = v132
                                        }
                                    case "foo bar":
                                        v1 += (v1 % ((v1) | 1))
                                    case "":
                                        v133 := GenS3{F0: 42, F1: (v114 << ((genFn0()) & 63)), F2: 2}
                                        _ = v133
                                    }
                                    fmt.Println(v3)
                                    _ = fi125
                                }
                                _ = v124
                            } else {
                                v4[v115] = (v10 >> ((1) & 63))
                                v114 += 65535
                                if true {
                                    for fi134 := 0; fi134 < 15; fi134++ {
                                        v135 := GenS7{F0: 42, F1: 42}
                                        v3 -= 599
                                        v2[9223372036854775807] = 994
                                        _ = fi134
                                        _ = v135
                                    }
                                    delete(v4, (v115 + "hello"))
                                }
                            }
                        }
                        if true {
                            delete(v4, "a")
                        } else {
                            if (1 == 2) {
                                switch v115 {
                                case "hello":
                                    if false {
                                        switch 32767 {
                                        case 128:
                                            for fi136 := 0; fi136 < 12; fi136++ {
                                                var v137 bool = (!(v1 != (^(v1))))
                                                fmt.Println(v115)
                                                if v137 {
                                                    switch "\n\t" {
                                                    case "":
                                                        for fi138 := 0; fi138 < 13; fi138++ {
                                                            if v137 {
                                                                switch ((("foo bar" + v115) + "hello") + (("foo bar" + v115) + (v115 + v115))) {
                                                                case "a":
                                                                    var v139 uint16 = v11
                                                                    _ = v139
                                                                default:
                                                                    v140 := GenS15{F0: (v11 / ((v114) | 1)), F1: uint64(v10)}
                                                                    _ = v140
                                                                }
                                                                var v141 int8 = (v1 << ((v3) & 63))
                                                                v3 = 2
                                                                _ = v141
                                                            }
                                                            _ = fi138
                                                        }
                                                    case "\n\t":
                                                        if v137 {
                                                            for fi142 := 0; fi142 < 5; fi142++ {
                                                                v143 := GenS16{F0: 2147483647, F1: 958, F2: -1}
                                                                _ = fi142
                                                                _ = v143
                                                            }
                                                            v3 = (v3 / ((genFn0()) | 1))
                                                        } else {
                                                            for fi144 := 0; fi144 < 12; fi144++ {
                                                                if v137 {
                                                                    if ((2147483647 < 2147483647) || (65535 == (v3 & (v3 << ((genFn0()) & 63))))) {
                                                                        if ((3.14159 >= (((((float32(1.5) + float32((0.0 / 1.5))) * (0.0 + (float32(((0.0 * 0.0) / 1.5)) + 1e10))) + 1.5) - (3.14159 + (-2.25 - (((float32(3.14159) + 1e10) + (1.5 / 3.14159)) / 1.0)))) + 0.0)) || false) {
                                                                            _ = len(v4)
                                                                            if v137 {
                                                                                v145 := GenS14{F0: (^(v11)), F1: v3}
                                                                                switch 42 {
                                                                                case -9223372036854775808:
                                                                                    v11 = 65535
                                                                                case 4294967295:
                                                                                    var v146 uint32 = (v10 & (v10 << (((v3 | (v3 & (v3 ^ 18446744073709551615)))) & 63)))
                                                                                    _ = v146
                                                                                }
                                                                                _ = v145
                                                                            }
                                                                            v115 = v115
                                                                        }
                                                                        var v147 uint8 = 2
                                                                        _ = v147
                                                                    } else {
                                                                        v148 := v4[v115]
                                                                        v148 = v148
                                                                        _ = v148
                                                                    }
                                                                } else {
                                                                    switch fi144 {
                                                                    case 263:
                                                                        if (127 != 198) {
                                                                            v115 = "\n\t"
                                                                        }
                                                                    }
                                                                    switch fi136 {
                                                                    case 127:
                                                                        v149 := GenS13{F0: uint16(v10), F1: -925}
                                                                        switch 32768 {
                                                                        case 4294967295:
                                                                            v11 -= 42
                                                                        default:
                                                                            var v150 string = v115
                                                                            _ = v150
                                                                        }
                                                                        v1 = -1
                                                                        var v151 bool = (uint64(v3) != 42)
                                                                        _ = v149
                                                                        _ = v151
                                                                    case -795:
                                                                        v10 -= (v10 / ((1) | 1))
                                                                    case 407:
                                                                        v152 := GenS1{F0: (v1 * (v1 | (v1 / ((0) | 1))))}
                                                                        _ = v152
                                                                    default:
                                                                        fmt.Println(v114, fi136, v10)
                                                                    }
                                                                    var v153 uint32 = (v10 * v10)
                                                                    _ = v153
                                                                }
                                                                switch "" {
                                                                case "hello":
                                                                    fmt.Println(fi144, v11)
                                                                }
                                                                v154 := v4[("\n\t" + "hello")]
                                                                var v155 int8 = v1
                                                                _ = fi144
                                                                _ = v154
                                                                _ = v155
                                                            }
                                                            v156 := GenS17{F0: v1, F1: (v3 + 2)}
                                                            v3 -= (v3 << ((genFn0()) & 63))
                                                            _ = v156
                                                        }
                                                    }
                                                    if v137 {
                                                        v157 := GenS18{F0: (v3 << (((v3 * genFn0())) & 63))}
                                                        if (!v137) {
                                                            if v137 {
                                                                delete(v2, (fi136 << ((v3) & 63)))
                                                            }
                                                        }
                                                        for fi158 := 0; fi158 < 4; fi158++ {
                                                            var v159 int64 = 0
                                                            switch (fi158 * fi136) {
                                                            case 972:
                                                                fmt.Println(fi136, v137, v3)
                                                            default:
                                                                v137 = (!(42 > 4294967295))
                                                            }
                                                            _ = fi158
                                                            _ = v159
                                                        }
                                                        _ = v157
                                                    }
                                                    if v137 {
                                                        v160 := v2[fi136]
                                                        v3 = (^(v3))
                                                        _ = v160
                                                    }
                                                } else {
                                                    var v161 uint32 = v10
                                                    if ("" <= (v115 + v115)) {
                                                        v162 := GenS19{F0: 0, F1: (((((float64(1.5) * -2.25) + (float64(3.14159) + 0.0)) * (0.0 + 3.14159)) / ((3.14159 - 0.0) / 3.14159)) == 0.1)}
                                                        _ = v162
                                                    }
                                                    _ = v161
                                                }
                                                _ = fi136
                                                _ = v137
                                            }
                                        case 279:
                                            for fi163 := 0; fi163 < 5; fi163++ {
                                                v11 -= 0
                                                for fi164 := 0; fi164 < 10; fi164++ {
                                                    v165 := GenS20{F0: (v114 | v114), F1: 1, F2: 2}
                                                    var v166 float64 = -2.25
                                                    _ = fi164
                                                    _ = v165
                                                    _ = v166
                                                }
                                                switch (fi163 ^ fi163) {
                                                case 671:
                                                    v11 += (v11 | (^(v114)))
                                                }
                                                _ = fi163
                                            }
                                        default:
                                            v10 += uint32(v3)
                                        }
                                    } else {
                                        fmt.Println(v3, v114, v114)
                                        v114 = v11
                                        v167 := GenS12{F0: v1, F1: genFn0()}
                                        _ = v167
                                    }
                                case "\n\t":
                                    if false {
                                        if false {
                                            v3 = v3
                                            for fi168 := 0; fi168 < 6; fi168++ {
                                                var v169 uint64 = 18446744073709551615
                                                v115 = v115
                                                var v170 string = "a"
                                                _ = fi168
                                                _ = v169
                                                _ = v170
                                            }
                                        }
                                    } else {
                                        switch 9223372036854775807 {
                                        case 0:
                                            var v171 int16 = 108
                                            _ = v171
                                        case 715:
                                            var v172 int32 = 291
                                            _ = v172
                                        }
                                    }
                                default:
                                    for fi173 := 0; fi173 < 11; fi173++ {
                                        if ((18446744073709551615 <= 42) && (0 >= int8(v10))) {
                                            var v174 float32 = (1e10 + 1.5)
                                            v175 := []int{int(v114), fi173, (fi173 + fi173), 4294967295}
                                            v11 = v11
                                            _ = v174
                                            _ = v175
                                        }
                                        if false {
                                            if (!((!(v10 < (v10 * v10))) || false)) {
                                                v176 := []int{(fi173 | 929)}
                                                var v177 int32 = int32(v11)
                                                fmt.Println(v177, v115)
                                                _ = v176
                                                _ = v177
                                            }
                                        }
                                        delete(v4, (v115 + (v115 + "a")))
                                        _ = fi173
                                    }
                                }
                            } else {
                                if true {
                                    for fi178 := 0; fi178 < 14; fi178++ {
                                        v1 += 100
                                        var v179 int64 = 42
                                        if true {
                                            switch fi178 {
                                            case 42:
                                                if true {
                                                    if (v114 != 65535) {
                                                        v1 = v1
                                                    }
                                                    var v180 int64 = (^(v179))
                                                    _ = v180
                                                }
                                            case 522:
                                                switch (-(fi178)) {
                                                case 9223372036854775807:
                                                    var v181 int64 = (v179 | (v179 & v179))
                                                    _ = v181
                                                case 32767:
                                                    var v182 int8 = (-(v1))
                                                    _ = v182
                                                case 32768:
                                                    switch fi178 {
                                                    case 127:
                                                        var v183 int16 = 0
                                                        _ = v183
                                                    case 1:
                                                        var v184 bool = (!true)
                                                        _ = v184
                                                    case 0:
                                                        if true {
                                                            var v185 uint = 527
                                                            var v186 uint16 = 255
                                                            v2[-1] = v10
                                                            _ = v185
                                                            _ = v186
                                                        } else {
                                                            fmt.Println(v114)
                                                        }
                                                    default:
                                                        fmt.Println(v1)
                                                    }
                                                }
                                            }
                                            v187 := GenS19{F0: 441, F1: (42 > genFn0())}
                                            _ = v187
                                        }
                                        _ = fi178
                                        _ = v179
                                    }
                                    v114 += 65535
                                    v188 := GenS19{F0: 0, F1: false}
                                    _ = v188
                                } else {
                                    if (42 <= 128) {
                                        v189 := GenS21{F0: ""}
                                        _ = v189
                                    }
                                }
                                for fi190 := 0; fi190 < 8; fi190++ {
                                    v191 := GenS22{F0: 751, F1: 0}
                                    var v192 int64 = -1
                                    if false {
                                        var v193 float32 = -2.25
                                        var v194 uint16 = 747
                                        _ = v193
                                        _ = v194
                                    }
                                    _ = fi190
                                    _ = v191
                                    _ = v192
                                }
                            }
                            v10 += 1
                        }
                    }
                    if false {
                        var v195 float64 = -2.25
                        v3 -= genFn0()
                        v114 += 32767
                        _ = v195
                    }
                    v114 = (v114 % ((2) | 1))
                }
            }
            switch -922 {
            case 331:
                if true {
                    for fi196 := 0; fi196 < 1; fi196++ {
                        v197 := v4[v115]
                        delete(v4, v115)
                        _ = fi196
                        _ = v197
                    }
                    switch (v115 + v115) {
                    case "\n\t":
                        v114 = 643
                    }
                    v198 := []int{127, 127, -1}
                    _ = v198
                } else {
                    if true {
                        var v199 int64 = -9223372036854775808
                        if (831 >= v114) {
                            _ = len(v4)
                            v200 := GenS5{F0: 855}
                            v201 := []int{1}
                            _ = v200
                            _ = v201
                        }
                        _ = v199
                    }
                    var v202 bool = (1.5 > ((0.0 - ((0.1 / 1e10) - 0.0)) - 1e10))
                    _ = v202
                }
            default:
                _ = len(v2)
            }
            var v203 uint = (v3 | v3)
            _ = v203
        } else {
            for fi204 := 0; fi204 < 2; fi204++ {
                v205 := GenS6{F0: v3, F1: 59, F2: 864}
                switch v115 {
                case "a":
                    v206 := GenS23{F0: -1, F1: 42, F2: -2147483648}
                    _ = v206
                default:
                    if (false || true) {
                        if (1e10 >= (((1.5 + 0.1) - 1.5) / float64((0.1 * 0.1)))) {
                            switch (fi204 ^ 32767) {
                            case 128:
                                v207 := GenS3{F0: 1, F1: 358, F2: 195}
                                _ = v207
                            case 1:
                                fmt.Println(v11, v11, v1)
                            case 32768:
                                var v208 int16 = 374
                                _ = v208
                            default:
                                if (0 >= -2147483648) {
                                    var v209 string = v115
                                    v210 := []int{(fi204 - fi204), -847}
                                    _ = v209
                                    _ = v210
                                }
                            }
                            v211 := GenS24{F0: (v10 >> (((v3 % ((1) | 1))) & 63)), F1: 445, F2: 1}
                            v10 -= (v10 >> ((v3) & 63))
                            _ = v211
                        }
                        var v212 string = ""
                        _ = v212
                    } else {
                        var v213 int64 = 42
                        _ = v213
                    }
                }
                _ = fi204
                _ = v205
            }
            fmt.Println(v115, v115, v115)
        }
        _ = v115
    }
    var v214 bool = (-2147483648 != 42)
    v215 := []int{9223372036854775807, 32768, 32768}
    v3 = 0
    switch 147 {
    case 32768:
        for fi216 := 0; fi216 < 9; fi216++ {
            if ((v3 / ((genFn0()) | 1)) > (v3 | 1)) {
                v217 := GenS2{F0: 9, F1: 4294967295, F2: "a"}
                if v214 {
                    v218 := v2[4294967295]
                    v219 := GenS22{F0: v218, F1: -1}
                    _ = v218
                    _ = v219
                } else {
                    for fi220 := 0; fi220 < 4; fi220++ {
                        if v214 {
                            if v214 {
                                var v221 uint16 = 32767
                                v222 := GenS8{F0: 255}
                                _ = v221
                                _ = v222
                            }
                            switch fi216 {
                            case 128:
                                v114 += (v114 + v11)
                            case -250:
                                v3 -= (v3 ^ v3)
                            case 127:
                                if (1.5 >= 1.5) {
                                    var v223 float32 = (1e10 / 1.0)
                                    v10 = 4294967295
                                    var v224 int = 1
                                    _ = v223
                                    _ = v224
                                }
                            }
                            if len(v215) > 0 {
                                v215[0] = fi220
                            }
                        }
                        if (v214 && (-1 > -9223372036854775808)) {
                            switch (fi216 * fi220) {
                            case -45:
                                if (-1 > 1) {
                                    v225 := GenS18{F0: v3}
                                    var v226 float32 = 1e10
                                    _ = v225
                                    _ = v226
                                } else {
                                    var v227 uint64 = 285
                                    fmt.Println(v214)
                                    v228 := GenS25{F0: (v227 ^ (v227 & v227))}
                                    v215 = append(v215, fi216)
                                    _ = v227
                                    _ = v228
                                }
                            case 32768:
                                var v229 float32 = (0.0 - (1.5 - 1e10))
                                _ = v229
                            default:
                                if (1 <= -1) {
                                    switch ("\n\t" + "\n\t") {
                                    case "\n\t":
                                        fmt.Println(fi216)
                                    case "":
                                        v214 = true
                                    default:
                                        switch fi220 {
                                        case 128:
                                            _ = len(v2)
                                            var v230 uint8 = 255
                                            _ = v230
                                        case -881:
                                            v1 += 1
                                        case 42:
                                            v231 := GenS7{F0: 2, F1: 684}
                                            _ = v231
                                        }
                                    }
                                    v232 := GenS26{F0: (v1 & (v1 & (v1 & (-(v1)))))}
                                    v1 += v1
                                    if (!true) {
                                        var v233 int64 = -9223372036854775808
                                        _ = v233
                                    } else {
                                        var v234 int64 = -1
                                        switch "foo bar" {
                                        case "hello":
                                            v215 = append(v215, fi220)
                                        case "foo bar":
                                            fmt.Println(v214, v114)
                                        default:
                                            v215 = append(v215, (fi220 / ((fi220) | 1)))
                                        }
                                        v235 := GenS27{F0: (v114 | v11)}
                                        _ = v234
                                        _ = v235
                                    }
                                    _ = v232
                                }
                            }
                            var v236 int = (fi216 >> ((255) & 63))
                            _ = v236
                        } else {
                            var v237 bool = (v10 < 2)
                            fmt.Println(fi220, fi216)
                            var v238 bool = ((fi216 <= fi216) || true)
                            _ = v237
                            _ = v238
                        }
                        v239 := GenS28{F0: 0, F1: fi220}
                        var v240 int = fi216
                        _ = fi220
                        _ = v239
                        _ = v240
                    }
                }
                _ = v217
            } else {
                switch fi216 {
                case 160:
                    fmt.Println(v3)
                case 4294967295:
                    var v241 float32 = (0.0 * -2.25)
                    _ = v241
                case 2147483647:
                    var v242 int64 = 664
                    _ = v242
                }
                for fi243 := 0; fi243 < 13; fi243++ {
                    delete(v4, "\n\t")
                    v11 += 959
                    var v244 int16 = -32768
                    _ = fi243
                    _ = v244
                }
            }
            v114 -= (^(v11))
            v245 := GenS29{F0: (v1 >> ((v3) & 63)), F1: (((3.14159 - 3.14159) < 0.0) || v214), F2: (fi216 * (fi216 | int(v3)))}
            _ = fi216
            _ = v245
        }
    case -875:
        if v214 {
            v1 = (v1 & 42)
        } else {
            if v214 {
                delete(v4, "\n\t")
                v246 := GenS8{F0: 12}
                if (false || false) {
                    v10 -= (^(v10))
                    switch "a" {
                    case "hello":
                        fmt.Println(v11, v10)
                    case "":
                        switch 32767 {
                        case -511:
                            v247 := v2[-1]
                            _ = v247
                        default:
                            v10 = v10
                        }
                    default:
                        for fi248 := 0; fi248 < 17; fi248++ {
                            if ((v1 * v1) > (v1 | (v1 | (^(v1))))) {
                                for fi249 := 0; fi249 < 10; fi249++ {
                                    var v250 uint16 = v11
                                    var v251 int8 = v1
                                    _ = fi249
                                    _ = v250
                                    _ = v251
                                }
                                delete(v2, fi248)
                            }
                            _ = fi248
                        }
                    }
                    v11 = 347
                } else {
                    v4["foo bar"] = (v10 % (((v10 ^ v10)) | 1))
                    _ = len(v215)
                    delete(v2, 127)
                }
                _ = v246
            } else {
                v252 := GenS3{F0: -730, F1: uint16(v10), F2: 255}
                v10 = 1
                _ = v252
            }
        }
    case 1:
        v253 := GenS3{F0: -291, F1: (v11 + v11), F2: uint8(v3)}
        _ = v253
    }
    fmt.Println(v114, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v114:", v114)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v214:", v214)
    fmt.Println("slicelen v215:", len(v215), v215)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v114
    _ = v12
    _ = v2
    _ = v214
    _ = v215
    _ = v3
    _ = v4
    _ = v9
}
