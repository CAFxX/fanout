package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 bool; F1 uint16; F2 int32 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int64; F1 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int; F1 int32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint32; F1 int32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 uint64; F1 int8 }
type GenS15 struct { F0 int; F1 int }

func main() {
    var v1 int8 = -128
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    v1 = (v1 - v1)
    if true {
        v1 -= -17
        v1 -= v1
    }
    switch int(v1) {
    case 42:
        v4 := GenS1{F0: 258, F1: v1}
        _ = v4
    case -1:
        var v5 uint32 = 65535
        _ = v5
    }
    v6 := GenS1{F0: 42, F1: 100}
    if true {
        if (987 <= 1) {
            for fi7 := 0; fi7 < 16; fi7++ {
                v1 += 0
                _ = len(v2)
                _ = fi7
            }
            v8 := []int{1, -718}
            v1 = (v1 + v1)
            _ = v8
        }
        if (420 <= 42) {
            var v9 float64 = -2.25
            v1 = (v1 << ((766) & 63))
            v10 := GenS1{F0: 4294967295, F1: (^(v1))}
            _ = v10
            _ = v9
        } else {
            if (true || (!(!true))) {
                switch "" {
                case "\n\t":
                    fmt.Println(v1)
                }
                v11 := GenS2{F0: (1.5 > (3.14159 + 1e10))}
                _ = v11
            }
            if false {
                v1 -= 127
                v1 += (v1 >> ((962) & 63))
                var v12 uint16 = 32767
                _ = v12
            }
        }
        var v13 bool = (!(1 > -283))
        _ = v13
    } else {
        v14 := GenS2{F0: true}
        _ = v14
    }
    fmt.Println(v1)
    if false {
        switch ("foo bar" + ("" + ("hello" + "hello"))) {
        case "a":
            switch "\n\t" {
            case "foo bar":
                v15 := []int{}
                _ = v15
            case "a":
                v16 := GenS3{F0: 2, F1: 0}
                _ = v16
            }
        case "hello":
            for fi17 := 0; fi17 < 19; fi17++ {
                switch fi17 {
                case 4294967295:
                    var v18 int = fi17
                    _ = v18
                case -2147483648:
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v20 := GenS3{F0: 18446744073709551615, F1: (v1 / ((-128) | 1))}
                        var v21 int64 = 42
                        if ((!false) && false) {
                            if (true && (1 < 1)) {
                                v22 := GenS3{F0: 279, F1: 0}
                                _ = v22
                            } else {
                                v21 = int64(v1)
                                if false {
                                    if true {
                                        switch (fi17 - fi19) {
                                        case -9223372036854775808:
                                            var v23 float32 = (((-2.25 * 1.5) * 1e10) * 0.1)
                                            _ = v23
                                        case 32768:
                                            switch "\n\t" {
                                            case "hello":
                                                v24 := GenS2{F0: (64 == 159)}
                                                _ = v24
                                            case "\n\t":
                                                delete(v2, (("foo bar" + "a") + "a"))
                                            case "foo bar":
                                                v25 := GenS4{F0: false, F1: 1, F2: -1}
                                                fmt.Println(fi19, v1, v21)
                                                _ = v25
                                            }
                                        default:
                                            if true {
                                                switch fi17 {
                                                case -77:
                                                    switch "" {
                                                    case "foo bar":
                                                        v26 := GenS5{F0: "\n\t", F1: (v21 << ((474) & 63))}
                                                        _ = v26
                                                    case "hello":
                                                        fmt.Println(fi17, v21)
                                                    }
                                                }
                                                var v27 float64 = ((float64(-2.25) * 3.14159) / (0.1 + -2.25))
                                                v21 = 42
                                                _ = v27
                                            }
                                        }
                                    } else {
                                        switch fi19 {
                                        case 32767:
                                            var v28 uint32 = 536
                                            _ = v28
                                        case 900:
                                            var v29 float32 = float32(0.0)
                                            _ = v29
                                        default:
                                            if (65535 > 1) {
                                                v30 := v2[""]
                                                v31 := []int{}
                                                switch (("\n\t" + "hello") + "") {
                                                case "":
                                                    fmt.Println(v1, fi17)
                                                    if len(v31) > 0 {
                                                        v31[0] = (fi19 >> ((0) & 63))
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v30, fi17)
                                                    v21 = v21
                                                    if true {
                                                        v32 := GenS6{F0: -1, F1: (!(false || v30))}
                                                        v1 -= v1
                                                        _ = v32
                                                    }
                                                case "a":
                                                    var v33 bool = (0 <= 42)
                                                    v31 = append(v31, 1)
                                                    v1 = v1
                                                    _ = v33
                                                default:
                                                    v30 = false
                                                }
                                                _ = v30
                                                _ = v31
                                            } else {
                                                v21 += v21
                                                if (-807 != fi19) {
                                                    v21 += 1
                                                    switch (fi19 + fi17) {
                                                    case 127:
                                                        if false {
                                                            var v34 uint64 = 0
                                                            _ = v34
                                                        }
                                                    case 9223372036854775807:
                                                        if (2147483647 == 104) {
                                                            switch "" {
                                                            case "hello":
                                                                if (false || true) {
                                                                    v1 -= v1
                                                                    var v35 uint = 904
                                                                    var v36 uint16 = 32767
                                                                    var v37 uint8 = 2
                                                                    var v38 int16 = 42
                                                                    _ = v35
                                                                    _ = v36
                                                                    _ = v37
                                                                    _ = v38
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(v21)
                                                                    }
                                                                }
                                                                var v39 uint64 = 42
                                                                var v40 int32 = 42
                                                                switch ("hello" + "") {
                                                                case "a":
                                                                    var v41 float32 = -2.25
                                                                    _ = v41
                                                                default:
                                                                    v39 -= v39
                                                                }
                                                                _ = v39
                                                                _ = v40
                                                            default:
                                                                v21 += v21
                                                            }
                                                            v1 += v1
                                                            v21 -= -917
                                                        }
                                                    }
                                                    if true {
                                                        var v42 string = ("\n\t" + "hello")
                                                        if false {
                                                            fmt.Println(v21, fi19, v42)
                                                            var v43 int = fi19
                                                            if false {
                                                                if false {
                                                                    switch (^(fi19)) {
                                                                    case 0:
                                                                        if true {
                                                                            var v44 int = 431
                                                                            _ = v44
                                                                        } else {
                                                                            v42 = ((v42 + v42) + v42)
                                                                            var v45 int16 = -691
                                                                            if (!(-197 >= 42)) {
                                                                                v21 += v21
                                                                                if false {
                                                                                    v46 := GenS5{F0: v42, F1: 1}
                                                                                    v47 := GenS7{F0: int64(v43)}
                                                                                    _ = v46
                                                                                    _ = v47
                                                                                } else {
                                                                                    var v48 int16 = 25
                                                                                    if ((v45 & (v48 << ((4294967295) & 63))) <= 1000) {
                                                                                        v1 -= (v1 ^ (v1 << ((18446744073709551615) & 63)))
                                                                                        v42 = (v42 + "\n\t")
                                                                                        fmt.Println(fi17, fi19)
                                                                                    } else {
                                                                                        fmt.Println(v48)
                                                                                        if (763 <= 1) {
                                                                                            var v49 int16 = (v45 >> ((42) & 63))
                                                                                            v3[fi19] = true
                                                                                            var v50 uint32 = 2147483647
                                                                                            _ = v49
                                                                                            _ = v50
                                                                                        }
                                                                                    }
                                                                                    var v51 uint = 963
                                                                                    _ = v48
                                                                                    _ = v51
                                                                                }
                                                                                var v52 uint = 331
                                                                                _ = v52
                                                                            } else {
                                                                                v53 := GenS4{F0: (255 <= 244), F1: 42, F2: 929}
                                                                                _ = v53
                                                                            }
                                                                            fmt.Println(v21, v42, v42)
                                                                            _ = v45
                                                                        }
                                                                    default:
                                                                        if ((float32(1.5) / ((1.5 - (1.5 / 1.5)) / 1e10)) >= (1.5 - -2.25)) {
                                                                            v54 := GenS5{F0: ("" + ""), F1: v21}
                                                                            v43 = 602
                                                                            v21 = (v21 << ((65535) & 63))
                                                                            _ = v54
                                                                        }
                                                                    }
                                                                    fmt.Println(v42, fi19, v42)
                                                                    delete(v2, v42)
                                                                }
                                                            }
                                                            var v55 int = (fi17 % ((v43) | 1))
                                                            _ = v43
                                                            _ = v55
                                                        } else {
                                                            v21 = -607
                                                            v56 := GenS8{F0: int64(v1)}
                                                            v21 += (v21 - 9223372036854775807)
                                                            _ = v56
                                                        }
                                                        _ = v42
                                                    } else {
                                                        var v57 float64 = 0.1
                                                        var v58 int16 = 0
                                                        if (("" + "") > "a") {
                                                            var v59 uint32 = 320
                                                            v57 = v57
                                                            var v60 uint8 = 42
                                                            _ = v59
                                                            _ = v60
                                                        } else {
                                                            v57 += float64(3.14159)
                                                            v21 -= (v21 + 0)
                                                            v61 := GenS9{F0: (^(fi17)), F1: 42}
                                                            _ = v61
                                                        }
                                                        _ = v57
                                                        _ = v58
                                                    }
                                                } else {
                                                    v3[fi17] = (fi19 == fi17)
                                                    if (354 < 65535) {
                                                        var v62 string = (("\n\t" + ("hello" + ("a" + "foo bar"))) + ("" + "hello"))
                                                        var v63 int32 = 42
                                                        _ = v62
                                                        _ = v63
                                                    } else {
                                                        switch "a" {
                                                        case "a":
                                                            v64 := GenS4{F0: (1 == 588), F1: 865, F2: 42}
                                                            _ = v64
                                                        }
                                                        fmt.Println(fi17, fi19, fi17)
                                                        v1 = -128
                                                    }
                                                    var v65 uint64 = 1
                                                    _ = v65
                                                }
                                            }
                                            v1 -= v1
                                        }
                                        switch "a" {
                                        case "":
                                            v66 := []int{fi19}
                                            var v67 uint8 = 255
                                            if (float32(1e10) <= (float32(3.14159) - (float32(((1e10 * (((0.0 / 1.0) - ((3.14159 / (3.14159 * 1e10)) + ((float64(1e10) + 1e10) / 0.1))) + -2.25)) - (3.14159 / 1.0))) - float32(3.14159)))) {
                                                var v68 uint = 1
                                                if true {
                                                    var v69 uint = 65535
                                                    fmt.Println(fi19, v69, v67)
                                                    var v70 uint = v69
                                                    _ = len(v2)
                                                    _ = v69
                                                    _ = v70
                                                }
                                                if false {
                                                    if (42 != -32768) {
                                                        switch fi17 {
                                                        case 0:
                                                            v67 += v67
                                                        case 4294967295:
                                                            switch ("foo bar" + "a") {
                                                            case "foo bar":
                                                                v2["a"] = false
                                                            case "a":
                                                                v71 := v2["a"]
                                                                var v72 uint8 = v67
                                                                _ = v71
                                                                _ = v72
                                                            case "\n\t":
                                                                v73 := GenS10{F0: 2}
                                                                _ = v73
                                                            default:
                                                                if len(v66) >= 2 {
                                                                    v66 = v66[1:len(v66)]
                                                                }
                                                            }
                                                        default:
                                                            var v74 string = (((("a" + (("" + "hello") + "foo bar")) + "a") + "hello") + ("a" + "foo bar"))
                                                            _ = v74
                                                        }
                                                        var v75 uint32 = 2147483648
                                                        switch "" {
                                                        case "\n\t":
                                                            if true {
                                                                if (42 == (v75 / (((v75 % ((18) | 1))) | 1))) {
                                                                    if ((((float64(((-2.25 - 1.5) - float32(-2.25))) + 3.14159) * float64(((3.14159 - 1.5) - (float32(0.0) / 1e10)))) - 3.14159) > ((1e10 * (0.1 * float64(3.14159))) / 0.1)) {
                                                                        var v76 uint64 = 700
                                                                        if false {
                                                                            var v77 float64 = float64(float32(0.1))
                                                                            fmt.Println(v67, fi19, v67)
                                                                            _ = v77
                                                                        }
                                                                        var v78 uint = (v68 / ((692) | 1))
                                                                        _ = v76
                                                                        _ = v78
                                                                    } else {
                                                                        v21 = 0
                                                                    }
                                                                }
                                                                v79 := v2["\n\t"]
                                                                var v80 int = -201
                                                                var v81 float64 = -2.25
                                                                v75 -= (^(v75))
                                                                _ = v79
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                        case "a":
                                                            v75 += v75
                                                        case "hello":
                                                            if (0 > (v67 / ((v67) | 1))) {
                                                                v67 -= uint8(v75)
                                                                var v82 uint32 = v75
                                                                _ = v82
                                                            } else {
                                                                var v83 uint8 = (v67 ^ (v67 << ((v68) & 63)))
                                                                v84 := GenS1{F0: (v68 - 758), F1: (v1 | int8(v21))}
                                                                _ = v83
                                                                _ = v84
                                                            }
                                                        default:
                                                            var v85 uint16 = 65535
                                                            _ = v85
                                                        }
                                                        _ = v75
                                                    }
                                                    v68 = v68
                                                    v86 := GenS4{F0: false, F1: 2, F2: 2147483647}
                                                    v87 := GenS3{F0: v68, F1: v1}
                                                    switch "\n\t" {
                                                    case "a":
                                                        switch (fi19 + -9223372036854775808) {
                                                        case 434:
                                                            v88 := GenS5{F0: "a", F1: v21}
                                                            _ = v88
                                                        case -2147483648:
                                                            var v89 string = ""
                                                            var v90 int16 = 0
                                                            _ = v89
                                                            _ = v90
                                                        default:
                                                            if ((v1 >> ((42) & 63)) != v1) {
                                                                switch (fi17 % (((fi17 / (((fi17 >> (((v68 >> ((65535) & 63))) & 63))) | 1))) | 1)) {
                                                                case -505:
                                                                    v68 = v68
                                                                case 0:
                                                                    v68 += v68
                                                                }
                                                                var v91 bool = (true && false)
                                                                _ = v91
                                                            } else {
                                                                switch (fi19 | fi19) {
                                                                case -1:
                                                                    v67 = 1
                                                                case -2147483648:
                                                                    _ = len(v3)
                                                                }
                                                            }
                                                        }
                                                    }
                                                    _ = v86
                                                    _ = v87
                                                } else {
                                                    if true {
                                                        v92 := GenS11{F0: 2147483648, F1: int32(v1)}
                                                        var v93 uint32 = 959
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                }
                                                _ = v68
                                            }
                                            _ = v66
                                            _ = v67
                                        case "foo bar":
                                            fmt.Println(fi19)
                                        default:
                                            v94 := GenS1{F0: 0, F1: (v1 & v1)}
                                            _ = v94
                                        }
                                        v1 = v1
                                        v95 := GenS3{F0: 396, F1: v1}
                                        var v96 uint32 = 2147483647
                                        _ = v95
                                        _ = v96
                                    }
                                    var v97 int64 = (v21 & (v21 & (v21 - 9223372036854775807)))
                                    _ = v97
                                }
                            }
                            switch fi19 {
                            case 0:
                                v98 := GenS12{F0: 0}
                                _ = v98
                            }
                            v21 -= 42
                        } else {
                            if (false || (-728 <= int32(v1))) {
                                fmt.Println(fi17, fi17, v21)
                            }
                            if true {
                                var v99 string = ("foo bar" + "a")
                                _ = v99
                            }
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                default:
                    if true {
                        if (true || ((1 >= 0) || false)) {
                            if ("foo bar" > "foo bar") {
                                var v100 uint8 = 255
                                _ = v100
                            }
                            if (577 < 4294967295) {
                                v101 := GenS8{F0: 42}
                                var v102 uint16 = 314
                                _ = v101
                                _ = v102
                            }
                        } else {
                            var v103 bool = false
                            switch fi17 {
                            case 826:
                                delete(v3, (fi17 >> ((4294967295) & 63)))
                            case -1:
                                var v104 uint16 = 32767
                                _ = v104
                            }
                            for fi105 := 0; fi105 < 13; fi105++ {
                                var v106 int32 = 1
                                fmt.Println(fi105, fi105)
                                v107 := GenS9{F0: fi105, F1: -1}
                                _ = fi105
                                _ = v106
                                _ = v107
                            }
                            _ = v103
                        }
                        v1 -= v1
                    }
                }
                if ((float64(float32((-2.25 * (0.0 - (((3.14159 - 1.5) - -2.25) + (float64(float32((0.0 * float64(((float32(((0.1 - float64(1e10)) * (0.1 + 1.5))) / 1e10) * ((-2.25 + ((1.5 / 1e10) - 0.0)) - 0.0)))))) * ((0.1 + 1.5) / 1.0))))))) * 1e10) <= ((3.14159 * (1.5 + 1e10)) * float64(0.1))) {
                    var v108 float64 = 0.0
                    var v109 float64 = -2.25
                    v110 := GenS13{F0: 896}
                    _ = v108
                    _ = v109
                    _ = v110
                }
                for fi111 := 0; fi111 < 14; fi111++ {
                    var v112 float32 = 1e10
                    var v113 int8 = v1
                    fmt.Println(fi111)
                    v114 := GenS9{F0: -9223372036854775808, F1: -2147483648}
                    _ = fi111
                    _ = v112
                    _ = v113
                    _ = v114
                }
                _ = fi17
            }
        }
        v1 = v1
        v115 := GenS9{F0: 2147483647, F1: -2147483648}
        _ = v115
    }
    v1 += v1
    var v116 bool = (42 > -2147483648)
    if (2 >= 569) {
        for fi117 := 0; fi117 < 19; fi117++ {
            if true {
                var v118 uint16 = 435
                v119 := GenS13{F0: v118}
                _ = v118
                _ = v119
            }
            switch fi117 {
            case 2147483647:
                var v120 uint8 = 127
                _ = v120
            }
            for fi121 := 0; fi121 < 17; fi121++ {
                var v122 int32 = 42
                _ = fi121
                _ = v122
            }
            _ = fi117
        }
    }
    v116 = false
    for fi123 := 0; fi123 < 19; fi123++ {
        v124 := v2[""]
        if false {
            var v125 uint32 = 42
            _ = v125
        }
        _ = fi123
        _ = v124
    }
    if true {
        switch 128 {
        case 32767:
            v116 = false
        case -9223372036854775808:
            var v126 int32 = 1
            _ = v126
        case -909:
            var v127 int8 = v1
            _ = v127
        }
        var v128 int = 4294967295
        if v116 {
            var v129 int = 688
            switch (v129 << ((4294967295) & 63)) {
            case 32767:
                var v130 int = (v129 | (v129 * 32768))
                _ = v130
            case 0:
                fmt.Println(v1, v1)
            case 4294967295:
                if (v129 < 32768) {
                    var v131 uint64 = 78
                    var v132 float64 = (1e10 + 0.0)
                    v133 := GenS14{F0: (v131 / (((v131 * (v131 & v131))) | 1)), F1: v1}
                    _ = v131
                    _ = v132
                    _ = v133
                }
            }
            var v134 uint32 = 2
            _ = v129
            _ = v134
        } else {
            v128 += (^(v128))
        }
        _ = v128
    }
    v1 += (v1 & v1)
    v135 := GenS15{F0: 128, F1: 32767}
    v136 := GenS12{F0: 0}
    v137 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("slicelen v137:", len(v137), v137)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v116
    _ = v135
    _ = v136
    _ = v137
    _ = v2
    _ = v3
    _ = v6
}
