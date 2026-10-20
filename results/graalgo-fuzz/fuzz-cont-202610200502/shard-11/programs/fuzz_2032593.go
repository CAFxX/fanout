package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 bool; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 int8; F1 int32; F2 int16 }
type GenS9 struct { F0 int32; F1 uint8 }
type GenS10 struct { F0 int32; F1 uint16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint8 }
func genFn0() int {
    if (0 >= 65535) {
        if true {
            v5 := GenS1{F0: 4294967295}
            _ = v5
        } else {
            v6 := GenS1{F0: 0}
            _ = v6
        }
    } else {
        var v7 uint = 65535
        if (((0 < 127) && false) || true) {
            v7 = v7
            v8 := []int{-1, 32767, 9223372036854775807, 127}
            v7 = (^(v7))
            _ = v8
        } else {
            var v9 float64 = 3.14159
            _ = v9
        }
        _ = v7
    }
    var v10 int16 = -32768
    _ = v10
    return 761
}
func genFn1(p0 bool, p1 uint, p2 int32) uint16 {
    fmt.Println(p2, p2, p0)
    v11 := []int{127, -840, -987, 32767}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 255
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = ((1e10 + 0.0) - float64(1.5))
    var v3 int16 = -603
    var v4 int64 = -16
    if true {
        if (true && true) {
            var v12 float32 = (((1.5 + 3.14159) / (-2.25 / 1.0)) * (0.0 - 0.0))
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            v2 = 3.14159
            if (float64(-2.25) > v2) {
                if (v2 < 1.5) {
                    var v14 bool = false
                    v3 = (v3 & v3)
                    _ = v14
                } else {
                    fmt.Println(v2, fi13)
                    v15 := GenS1{F0: 18446744073709551615}
                    _ = v15
                }
                v2 = v2
            } else {
                v3 = (v3 + v3)
                switch "foo bar" {
                case "\n\t":
                    if true {
                        if false {
                            var v16 uint8 = 128
                            fmt.Println(v3, v4, v3)
                            _ = v16
                        }
                        v3 += (-(v3))
                        var v17 int = (fi13 >> ((1) & 63))
                        _ = v17
                    }
                case "":
                    v18 := GenS2{F0: 42}
                    _ = v18
                default:
                    var v19 uint64 = 187
                    _ = v19
                }
                v20 := GenS2{F0: 32}
                _ = v20
            }
            _ = fi13
        }
    } else {
        var v21 uint = 18446744073709551615
        fmt.Println(v3)
        _ = v21
    }
    v4 -= v4
    v22 := []int{32767, 128, -2147483648}
    fmt.Println(v2, v2)
    switch "hello" {
    case "foo bar":
        v22 = append(v22, genFn0())
    case "\n\t":
        v23 := v1[(("a" + ("hello" + ("hello" + ("a" + ("foo bar" + "a"))))) + "")]
        _ = v23
    default:
        for fi24 := 0; fi24 < 20; fi24++ {
            fmt.Println(v2)
            switch ("foo bar" + (("hello" + ("a" + "a")) + "foo bar")) {
            case "":
                fmt.Println(v4)
            case "hello":
                v25 := v1["foo bar"]
                _ = v25
            case "foo bar":
                var v26 bool = true
                _ = v26
            default:
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            }
            var v27 uint8 = 42
            _ = fi24
            _ = v27
        }
    }
    fmt.Println(v2)
    for fi28 := 0; fi28 < 18; fi28++ {
        if (2147483647 <= 1) {
            var v29 uint32 = 706
            var v30 uint32 = (v29 + v29)
            if (0 == 189) {
                switch 4294967295 {
                case -9223372036854775808:
                    switch "" {
                    case "":
                        for fi31 := 0; fi31 < 1; fi31++ {
                            switch fi31 {
                            case 127:
                                fmt.Println(v2, v29, v30)
                            case -1:
                                switch "" {
                                case "hello":
                                    var v32 string = "foo bar"
                                    _ = v32
                                case "":
                                    delete(v1, "hello")
                                case "\n\t":
                                    fmt.Println(v29)
                                }
                            }
                            v2 -= 0.1
                            v33 := GenS2{F0: 1}
                            if (v3 <= v3) {
                                v1[("hello" + ("a" + ("\n\t" + ("a" + ("\n\t" + "foo bar")))))] = (^(fi28))
                            } else {
                                var v34 uint16 = 0
                                var v35 float64 = -2.25
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                switch fi28 {
                                case 128:
                                    var v36 uint64 = 0
                                    switch -9223372036854775808 {
                                    case 127:
                                        if ((^(v3)) > -1) {
                                            v4 = v4
                                            if true {
                                                v37 := GenS2{F0: 255}
                                                if true {
                                                    var v38 int16 = 1
                                                    if false {
                                                        delete(v1, ("foo bar" + "\n\t"))
                                                    } else {
                                                        switch "" {
                                                        case "a":
                                                            var v39 uint64 = (v36 - 0)
                                                            _ = v39
                                                        case "hello":
                                                            v1["foo bar"] = (fi31 >> ((18446744073709551615) & 63))
                                                        case "\n\t":
                                                            _ = len(v1)
                                                            v30 = (v30 ^ (^(v30)))
                                                        }
                                                        var v40 int64 = -304
                                                        v3 -= 1
                                                        _ = v40
                                                    }
                                                    if true {
                                                        v41 := GenS2{F0: 4294967295}
                                                        fmt.Println(v4, v38)
                                                        v1["\n\t"] = (-(fi31))
                                                        _ = v41
                                                    }
                                                    if false {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v42 int64 = -1
                                                            _ = v42
                                                        case "foo bar":
                                                            fmt.Println(v4, fi31, v38)
                                                        case "":
                                                            var v43 uint = 137
                                                            _ = v43
                                                        }
                                                        switch fi31 {
                                                        case 0:
                                                            var v44 int16 = 1
                                                            _ = v44
                                                        case 127:
                                                            v45 := GenS1{F0: uint64(v34)}
                                                            _ = v45
                                                        case -932:
                                                            v22 = append(v22, (-(fi28)))
                                                            switch "a" {
                                                            case "hello":
                                                                v46 := GenS3{F0: uint(v36), F1: v38}
                                                                _ = v46
                                                            default:
                                                                var v47 int8 = 42
                                                                _ = v47
                                                            }
                                                            delete(v1, "foo bar")
                                                        default:
                                                            v29 -= 2147483648
                                                        }
                                                    } else {
                                                        v22 = append(v22, fi28)
                                                        if false {
                                                            fmt.Println(v38)
                                                        } else {
                                                            switch "foo bar" {
                                                            case "":
                                                                _ = len(v22)
                                                            case "foo bar":
                                                                var v48 uint64 = 42
                                                                _ = v48
                                                            default:
                                                                switch fi31 {
                                                                case 127:
                                                                    v49 := GenS1{F0: (v36 * (v36 << ((65535) & 63)))}
                                                                    _ = v49
                                                                case 664:
                                                                    v50 := GenS1{F0: (^(v36))}
                                                                    _ = v50
                                                                default:
                                                                    _ = len(v1)
                                                                }
                                                            }
                                                            if false {
                                                                if false {
                                                                    fmt.Println(v34, v30, v30)
                                                                } else {
                                                                    v22 = append(v22, fi28)
                                                                }
                                                                var v51 int8 = -128
                                                                var v52 uint = 255
                                                                _ = v51
                                                                _ = v52
                                                            } else {
                                                                v36 -= 4294967295
                                                            }
                                                            if false {
                                                                var v53 int = 9223372036854775807
                                                                var v54 string = "foo bar"
                                                                switch "\n\t" {
                                                                case "a":
                                                                    switch "hello" {
                                                                    case "a":
                                                                        if true {
                                                                            _ = len(v1)
                                                                            v36 = (v36 & (v36 - (^(v36))))
                                                                        }
                                                                    case "foo bar":
                                                                        var v55 float32 = ((3.14159 * 3.14159) - (1.5 * (1.5 * 0.0)))
                                                                        _ = v55
                                                                    }
                                                                }
                                                                _ = v53
                                                                _ = v54
                                                            }
                                                        }
                                                        switch (fi31 % (((fi28 - (fi28 * (fi28 * (fi31 / ((int(v38)) | 1)))))) | 1)) {
                                                        case 303:
                                                            v56 := GenS4{F0: (1 > genFn1(true, 788, 257)), F1: (v34 + 65535), F2: 2147483647}
                                                            _ = v56
                                                        }
                                                    }
                                                    _ = v38
                                                }
                                                _ = v37
                                            }
                                        } else {
                                            fmt.Println(v30, v35, v30)
                                            fmt.Println(v4, v29)
                                        }
                                    case 0:
                                        if true {
                                            v57 := GenS5{F0: v29}
                                            fmt.Println(v4, v2, fi28)
                                            _ = v57
                                        }
                                    case -9223372036854775808:
                                        if (-128 >= 24) {
                                            switch ("hello" + "a") {
                                            case "a":
                                                var v58 float64 = float64(float32(v35))
                                                _ = v58
                                            default:
                                                v4 -= (v4 << ((65535) & 63))
                                            }
                                            _ = len(v1)
                                        } else {
                                            v22 = append(v22, genFn0())
                                            if true {
                                                v35 = (float64(-2.25) - (-2.25 / v35))
                                                v30 = v29
                                                var v59 int = genFn0()
                                                _ = v59
                                            }
                                            if true {
                                                v29 -= (v29 + (v30 % ((v29) | 1)))
                                                v60 := GenS6{F0: 2147483648, F1: 1}
                                                v61 := GenS7{F0: (v34 / ((genFn1((0 <= 128), 2, -2147483648)) | 1)), F1: v3}
                                                _ = v60
                                                _ = v61
                                            } else {
                                                v35 = ((v2 / v35) - ((((v35 - v2) + float64(0.1)) - -2.25) / float64(-2.25)))
                                                delete(v1, "\n\t")
                                                v62 := GenS4{F0: ((v34 * (v34 + (v34 + v34))) > (v34 / ((42) | 1))), F1: genFn1((false && (0 >= 1)), 18446744073709551615, 91), F2: -2147483648}
                                                _ = v62
                                            }
                                        }
                                    }
                                    var v63 string = "\n\t"
                                    _ = v36
                                    _ = v63
                                default:
                                    v64 := v1[((("hello" + (("a" + "foo bar") + "a")) + "a") + "hello")]
                                    _ = v64
                                }
                                _ = v34
                                _ = v35
                            }
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(v4, v4, fi28)
                            case "":
                                fmt.Println(v2)
                            case "foo bar":
                                v22 = append(v22, (fi28 & fi31))
                            default:
                                v4 = v4
                            }
                            _ = fi31
                            _ = v33
                        }
                    case "hello":
                        if (true || (!(v3 <= v3))) {
                            for fi65 := 0; fi65 < 14; fi65++ {
                                if false {
                                    v66 := GenS4{F0: (-2147483648 < 42), F1: 913, F2: 504}
                                    if (!true) {
                                        _ = len(v22)
                                        var v67 int16 = 0
                                        _ = v67
                                    }
                                    _ = v66
                                }
                                _ = fi65
                            }
                            v29 += v29
                        }
                    }
                case 32768:
                    switch "" {
                    case "hello":
                        _ = len(v22)
                    case "a":
                        if false {
                            switch genFn0() {
                            case 127:
                                v2 += 0.1
                            case 32767:
                                fmt.Println(v2, v29)
                            default:
                                v68 := GenS1{F0: 1}
                                _ = v68
                            }
                        } else {
                            switch fi28 {
                            case 4294967295:
                                for fi69 := 0; fi69 < 1; fi69++ {
                                    switch (^(fi28)) {
                                    case 709:
                                        v2 = (v2 - float64(0.0))
                                    }
                                    if true {
                                        v2 += (v2 - v2)
                                        v1[("" + ("hello" + ("a" + "")))] = (fi28 << ((521) & 63))
                                    }
                                    if len(v22) > 0 {
                                        v22[0] = genFn0()
                                    }
                                    var v70 uint = 255
                                    _ = fi69
                                    _ = v70
                                }
                            case 32768:
                                for fi71 := 0; fi71 < 9; fi71++ {
                                    fmt.Println(fi71)
                                    if false {
                                        v2 -= v2
                                        if (!false) {
                                            v72 := GenS8{F0: 42, F1: 1, F2: (v3 >> ((116) & 63))}
                                            v4 = 42
                                            _ = v72
                                        } else {
                                            delete(v1, "")
                                        }
                                        v73 := GenS9{F0: 0, F1: 141}
                                        if ((0.0 + -2.25) < 1.5) {
                                            switch (fi71 >> ((255) & 63)) {
                                            case -2147483648:
                                                fmt.Println(fi28, v3, v2)
                                            case 32768:
                                                if true {
                                                    fmt.Println(v2)
                                                    v4 = v4
                                                    fmt.Println(v2, v29)
                                                }
                                            default:
                                                _ = len(v22)
                                            }
                                            _ = len(v22)
                                        } else {
                                            if false {
                                                switch (fi28 | (fi71 + (-(fi71)))) {
                                                case -9223372036854775808:
                                                    var v74 int32 = 0
                                                    _ = v74
                                                case 2147483647:
                                                    v4 = 0
                                                case -1:
                                                    if (909 != 0) {
                                                        var v75 int8 = 127
                                                        _ = v75
                                                    }
                                                default:
                                                    v76 := GenS1{F0: 2}
                                                    _ = v76
                                                }
                                                var v77 uint8 = 2
                                                _ = v77
                                            }
                                            if (42 < 1) {
                                                v4 -= v4
                                                if (238 != 57) {
                                                    var v78 uint64 = 42
                                                    switch (("\n\t" + ("" + (("a" + "a") + "a"))) + "\n\t") {
                                                    case "a":
                                                        v30 = v30
                                                    }
                                                    var v79 float64 = (((1.5 + v2) * 1e10) * 3.14159)
                                                    v80 := GenS7{F0: 0, F1: -1}
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                } else {
                                                    _ = len(v1)
                                                }
                                                fmt.Println(v29, fi28)
                                            }
                                            if (v29 == v30) {
                                                var v81 uint8 = 42
                                                _ = v81
                                            } else {
                                                switch ((((("foo bar" + ("hello" + ("hello" + "a"))) + "\n\t") + "\n\t") + "\n\t") + "\n\t") {
                                                case "a":
                                                    v82 := GenS1{F0: 42}
                                                    _ = v82
                                                case "":
                                                    if (42 >= 100) {
                                                        v83 := GenS2{F0: 0}
                                                        var v84 float32 = 1.5
                                                        var v85 int = (fi71 >> ((0) & 63))
                                                        _ = v83
                                                        _ = v84
                                                        _ = v85
                                                    }
                                                case "foo bar":
                                                    v4 = -9223372036854775808
                                                }
                                                var v86 int8 = -128
                                                _ = v86
                                            }
                                        }
                                        _ = v73
                                    } else {
                                        switch fi28 {
                                        case -1:
                                            v3 = (v3 - v3)
                                        }
                                        v30 -= (^(v29))
                                    }
                                    v87 := v1[""]
                                    _ = fi71
                                    _ = v87
                                }
                            case -2147483648:
                                fmt.Println(v2, v29)
                            default:
                                var v88 int32 = 838
                                _ = v88
                            }
                            for fi89 := 0; fi89 < 1; fi89++ {
                                switch (("hello" + "") + "\n\t") {
                                case "a":
                                    var v90 int64 = -9223372036854775808
                                    _ = v90
                                }
                                var v91 uint16 = genFn1((-2147483648 <= -418), 18446744073709551615, -1)
                                _ = fi89
                                _ = v91
                            }
                            v29 -= (v29 - (v29 % ((4294967295) | 1)))
                        }
                    case "\n\t":
                        for fi92 := 0; fi92 < 9; fi92++ {
                            if true {
                                if (!(0 == 0)) {
                                    var v93 int32 = 372
                                    _ = v93
                                } else {
                                    var v94 float64 = 1.5
                                    v94 -= float64(0.1)
                                    _ = v94
                                }
                                _ = len(v22)
                                fmt.Println(v2, v2)
                            } else {
                                switch 9223372036854775807 {
                                case -2147483648:
                                    switch (("\n\t" + "foo bar") + "\n\t") {
                                    case "":
                                        switch "a" {
                                        case "a":
                                            v95 := GenS3{F0: 255, F1: v3}
                                            _ = v95
                                        case "\n\t":
                                            var v96 int16 = (v3 / (((v3 & v3)) | 1))
                                            _ = v96
                                        }
                                    default:
                                        var v97 string = (("" + "a") + "")
                                        _ = v97
                                    }
                                default:
                                    switch (fi28 * 2147483647) {
                                    case 42:
                                        var v98 uint32 = (v29 << ((65535) & 63))
                                        _ = v98
                                    case 9223372036854775807:
                                        if (1e10 > (0.0 + (float32(v2) * 1e10))) {
                                            v99 := GenS1{F0: 509}
                                            v100 := GenS10{F0: int32(fi92), F1: genFn1(false, 2, 0)}
                                            _ = len(v22)
                                            _ = v100
                                            _ = v99
                                        } else {
                                            if false {
                                                v101 := GenS10{F0: 1, F1: 1}
                                                v2 = (v2 + (v2 / v2))
                                                if true {
                                                    v2 = ((1e10 + 3.14159) + (0.1 * v2))
                                                    var v102 int32 = 708
                                                    switch "hello" {
                                                    case "hello":
                                                        v103 := v1["hello"]
                                                        _ = v103
                                                    default:
                                                        var v104 uint8 = 221
                                                        _ = v104
                                                    }
                                                    _ = v102
                                                } else {
                                                    fmt.Println(v3, fi28)
                                                    var v105 bool = false
                                                    var v106 int32 = -2147483648
                                                    _ = v105
                                                    _ = v106
                                                }
                                                if true {
                                                    v2 += v2
                                                    if len(v22) > 0 {
                                                        v22[0] = 32767
                                                    }
                                                }
                                                _ = v101
                                            }
                                            v22 = append(v22, fi92)
                                            var v107 string = (((("foo bar" + "") + "a") + ("hello" + "foo bar")) + ("foo bar" + ("" + "a")))
                                            if (genFn0() < fi92) {
                                                v108 := GenS11{F0: (v4 | v4)}
                                                _ = v108
                                            } else {
                                                v4 = (v4 | v4)
                                                fmt.Println(v4, v30, v30)
                                            }
                                            _ = v107
                                        }
                                    case -1:
                                        v109 := GenS12{F0: 1}
                                        _ = v109
                                    default:
                                        if false {
                                            fmt.Println(v29)
                                        } else {
                                            var v110 float64 = ((1.5 / (3.14159 / 0.1)) / (v2 + v2))
                                            _ = v110
                                        }
                                    }
                                }
                                if (42 > 300) {
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                } else {
                                    if true {
                                        var v111 uint64 = 374
                                        switch ((("foo bar" + "a") + ((("hello" + ("a" + ("a" + "a"))) + "a") + "")) + "") {
                                        case "foo bar":
                                            fmt.Println(v2)
                                        case "hello":
                                            var v112 float32 = ((0.0 / 1e10) / (1e10 / 3.14159))
                                            _ = v112
                                        }
                                        _ = v111
                                    } else {
                                        var v113 int16 = (-(v3))
                                        _ = v113
                                    }
                                    switch genFn0() {
                                    case -313:
                                        var v114 int8 = 110
                                        _ = v114
                                    case -826:
                                        v3 += (v3 / ((v3) | 1))
                                    case 736:
                                        v115 := GenS2{F0: 255}
                                        _ = v115
                                    }
                                    if (-89 < 758) {
                                        var v116 int32 = 42
                                        _ = v116
                                    }
                                }
                            }
                            _ = fi92
                        }
                    default:
                        v1["foo bar"] = fi28
                    }
                case 937:
                    var v117 int8 = -1
                    _ = v117
                }
                _ = len(v1)
                v22 = append(v22, genFn0())
            }
            _ = v29
            _ = v30
        } else {
            fmt.Println(fi28)
            var v118 float64 = float64(((0.0 / 1.0) / -2.25))
            _ = v118
        }
        _ = fi28
    }
    v1["\n\t"] = 218
    var v119 uint16 = 2
    var v120 float64 = v2
    v22 = append(v22, 9223372036854775807)
    if (!(65535 == 18446744073709551615)) {
        v121 := v1["hello"]
        switch ("foo bar" + "hello") {
        case "hello":
            var v122 string = "hello"
            _ = v122
        case "a":
            v123 := GenS9{F0: 0, F1: 243}
            _ = v123
        default:
            v119 = v119
        }
        _ = v121
    } else {
        v124 := GenS4{F0: false, F1: v119, F2: 0}
        var v125 string = ""
        _ = v124
        _ = v125
    }
    if false {
        v4 -= v4
        for fi126 := 0; fi126 < 18; fi126++ {
            for fi127 := 0; fi127 < 7; fi127++ {
                var v128 uint32 = 65535
                _ = fi127
                _ = v128
            }
            var v129 uint16 = v119
            _ = fi126
            _ = v129
        }
    }
    v130 := GenS11{F0: v4}
    var v131 bool = ("\n\t" <= ((("hello" + "") + "\n\t") + "foo bar"))
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    if (!(!false)) {
        var v132 int64 = (v4 - v4)
        v131 = (!((("" + "foo bar") + (("\n\t" + "") + "a")) != ""))
        _ = v132
    }
    fmt.Println(v131)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v119:", v119)
    fmt.Println("v120:", v120)
    fmt.Println("v131:", v131)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v119
    _ = v120
    _ = v130
    _ = v131
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
