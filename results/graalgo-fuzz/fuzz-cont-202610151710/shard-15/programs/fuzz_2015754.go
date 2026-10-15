package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 uint; F1 uint16; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int64; F1 bool; F2 uint64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 int; F2 int64 }
type GenS9 struct { F0 uint32; F1 uint32 }
type GenS10 struct { F0 bool; F1 uint8 }
type GenS11 struct { F0 int64 }
func genFn0(p0 uint32) int32 {
    for fi4 := 0; fi4 < 4; fi4++ {
        p0 += (p0 ^ (p0 / ((p0) | 1)))
        fmt.Println(fi4, p0)
        _ = fi4
    }
    if false {
        v5 := []int{932, 128, 32767, -10}
        _ = v5
    } else {
        var v6 uint16 = 255
        v7 := GenS1{F0: 1, F1: 0}
        _ = v6
        _ = v7
    }
    _ = p0
    return -47
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 127
    var v3 uint = 42
    for fi8 := 0; fi8 < 15; fi8++ {
        v9 := v1[1]
        v10 := GenS2{F0: 127, F1: (v2 + v2)}
        v11 := GenS1{F0: 208, F1: int16(v2)}
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 uint8 = 127
    fmt.Println(v12)
    var v13 int8 = (^(v2))
    v14 := GenS2{F0: -252, F1: v2}
    var v15 uint32 = 2147483648
    v16 := GenS3{F0: (^(v3)), F1: 486, F2: "hello"}
    if false {
        switch 9223372036854775807 {
        case 0:
            if (1e10 <= 0.1) {
                v2 += -65
            } else {
                v17 := GenS4{F0: 1}
                _ = v17
            }
        case 530:
            v18 := GenS4{F0: v2}
            _ = v18
        default:
            for fi19 := 0; fi19 < 19; fi19++ {
                if true {
                    var v20 float64 = ((-2.25 / 3.14159) * 3.14159)
                    for fi21 := 0; fi21 < 11; fi21++ {
                        var v22 float64 = ((0.0 / v20) / -2.25)
                        v23 := GenS2{F0: (fi21 << (((v3 | v3)) & 63)), F1: 42}
                        switch ("" + "\n\t") {
                        case "hello":
                            if true {
                                v24 := GenS3{F0: v3, F1: 598, F2: ("\n\t" + ((("foo bar" + ("hello" + ("\n\t" + ("a" + "hello")))) + ("\n\t" + "foo bar")) + ("a" + ("\n\t" + "foo bar"))))}
                                _ = v24
                            }
                        case "\n\t":
                            if true {
                                if (true || (0 != -146)) {
                                    v25 := []int{2147483647, fi19, (fi19 % ((fi21) | 1)), fi21}
                                    if true {
                                        var v26 float64 = v20
                                        if (((v3 & 957) >= uint(v12)) || false) {
                                            var v27 int = -97
                                            v20 += v26
                                            _ = v27
                                        } else {
                                            switch fi21 {
                                            case 42:
                                                v26 = float64((float32((float64(1.5) + (float64(0.0) / float64(1e10)))) + -2.25))
                                                v2 = (v13 ^ (v2 + v2))
                                                if (1e10 == 3.14159) {
                                                    switch fi19 {
                                                    case -441:
                                                        v28 := GenS3{F0: v3, F1: 350, F2: ""}
                                                        _ = v28
                                                    case 32767:
                                                        switch "hello" {
                                                        case "\n\t":
                                                            var v29 float64 = float64((3.14159 / 3.14159))
                                                            _ = v29
                                                        case "a":
                                                            var v30 float64 = v22
                                                            _ = v30
                                                        }
                                                    case 127:
                                                        var v31 int = (fi21 ^ fi21)
                                                        _ = v31
                                                    default:
                                                        switch "\n\t" {
                                                        case "hello":
                                                            if len(v25) >= 2 {
                                                                v25 = v25[1:len(v25)]
                                                            }
                                                        default:
                                                            if false {
                                                                v32 := GenS5{F0: 65535, F1: 127}
                                                                _ = v32
                                                            } else {
                                                                var v33 float64 = float64(((0.0 * (float32((v26 * -2.25)) - 1.5)) / 1e10))
                                                                switch (fi19 & (fi21 | -9223372036854775808)) {
                                                                case 2147483647:
                                                                    v12 = (v12 + v12)
                                                                case -227:
                                                                    switch (-(fi19)) {
                                                                    case -1:
                                                                        var v34 int8 = (^(v13))
                                                                        _ = v34
                                                                    }
                                                                case -601:
                                                                    switch ("a" + "") {
                                                                    case "\n\t":
                                                                        var v35 uint64 = 1
                                                                        _ = v35
                                                                    }
                                                                default:
                                                                    var v36 int64 = 42
                                                                    _ = v36
                                                                }
                                                                _ = v33
                                                            }
                                                        }
                                                    }
                                                }
                                            case -9223372036854775808:
                                                _ = len(v25)
                                            case 128:
                                                switch fi21 {
                                                case 33:
                                                    fmt.Println(v3)
                                                case 127:
                                                    v15 += 917
                                                case 0:
                                                    v3 -= 971
                                                default:
                                                    switch (fi19 | 648) {
                                                    case 2147483647:
                                                        v15 = (^(v15))
                                                    }
                                                }
                                            default:
                                                v15 = v15
                                            }
                                        }
                                        _ = v26
                                    }
                                    v3 = 42
                                    _ = v25
                                }
                            } else {
                                var v37 string = ("hello" + "a")
                                _ = v37
                            }
                        default:
                            switch (("\n\t" + ("\n\t" + ("\n\t" + "foo bar"))) + "foo bar") {
                            case "foo bar":
                                var v38 uint16 = 0
                                _ = v38
                            case "a":
                                v39 := GenS1{F0: v12, F1: 0}
                                v3 = 2
                                if false {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v3 -= (^(v3))
                                    }
                                    v40 := GenS4{F0: 42}
                                    fmt.Println(v15, v22)
                                    fmt.Println(v2, v13, fi19)
                                    v41 := GenS5{F0: v3, F1: (v12 ^ v12)}
                                    v42 := GenS1{F0: v12, F1: 0}
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                } else {
                                    v22 -= v20
                                    var v43 uint64 = 288
                                    if false {
                                        fmt.Println(fi21)
                                    } else {
                                        if true {
                                            v44 := v1[340]
                                            _ = v44
                                        } else {
                                            switch (fi21 | fi19) {
                                            case 32767:
                                                switch ("\n\t" + "") {
                                                case "\n\t":
                                                    switch (-(fi19)) {
                                                    case 1:
                                                        v45 := GenS3{F0: v3, F1: 32767, F2: "hello"}
                                                        _ = v45
                                                    case 127:
                                                        switch "\n\t" {
                                                        case "":
                                                            _ = len(v1)
                                                        default:
                                                            fmt.Println(v12, v3)
                                                        }
                                                    case 754:
                                                        v46 := []int{fi19, (fi21 * fi21)}
                                                        _ = v46
                                                    default:
                                                        var v47 string = ("" + "foo bar")
                                                        _ = v47
                                                    }
                                                case "":
                                                    if true {
                                                        var v48 uint16 = uint16(v12)
                                                        var v49 string = ("a" + "foo bar")
                                                        _ = v48
                                                        _ = v49
                                                    }
                                                case "foo bar":
                                                    v50 := GenS3{F0: uint(v43), F1: 417, F2: ("hello" + "\n\t")}
                                                    switch "a" {
                                                    case "foo bar":
                                                        if ((v15 ^ 216) >= (v15 ^ 42)) {
                                                            v51 := []int{fi19, (fi21 + 0), -1}
                                                            if false {
                                                                if false {
                                                                    if false {
                                                                        if false {
                                                                            fmt.Println(fi21, v15, v3)
                                                                            if ((v3 | 18446744073709551615) < v3) {
                                                                                if true {
                                                                                    var v52 int16 = 1000
                                                                                    if false {
                                                                                        var v53 float64 = v22
                                                                                        var v54 int32 = int32(v43)
                                                                                        _ = v53
                                                                                        _ = v54
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        v52 += 42
                                                                                    case "a":
                                                                                        v55 := v1[v15]
                                                                                        _ = v55
                                                                                    case "":
                                                                                        v3 = 116
                                                                                        v56 := v1[v15]
                                                                                        var v57 bool = false
                                                                                        _ = v56
                                                                                        _ = v57
                                                                                    }
                                                                                    _ = v52
                                                                                }
                                                                                v58 := GenS5{F0: 1, F1: 2}
                                                                                switch (fi21 ^ fi19) {
                                                                                case 4294967295:
                                                                                    var v59 int8 = (v2 & (-(v13)))
                                                                                    _ = v59
                                                                                }
                                                                                _ = v58
                                                                            }
                                                                            var v60 float64 = 0.1
                                                                            _ = v60
                                                                        }
                                                                        switch (fi21 / (((fi21 % ((fi19) | 1))) | 1)) {
                                                                        case 32768:
                                                                            fmt.Println(fi21, v2, v20)
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                var v61 int8 = 80
                                                                                _ = v61
                                                                            }
                                                                            if len(v51) >= 2 {
                                                                                v51 = v51[1:len(v51)]
                                                                            }
                                                                            v43 = v43
                                                                        case 7:
                                                                            v3 += (^(v3))
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                if (true || true) {
                                                                    if len(v51) >= 2 {
                                                                        v51 = v51[1:len(v51)]
                                                                    }
                                                                } else {
                                                                    _ = len(v51)
                                                                }
                                                            }
                                                            var v62 uint8 = 0
                                                            _ = v51
                                                            _ = v62
                                                        } else {
                                                            switch ((("hello" + "hello") + ("\n\t" + "hello")) + (((("" + "hello") + "\n\t") + "a") + "")) {
                                                            case "":
                                                                v63 := GenS4{F0: (v2 >> ((v3) & 63))}
                                                                _ = v63
                                                            case "a":
                                                                v64 := GenS6{F0: 1, F1: (1 >= 1), F2: 2}
                                                                _ = v64
                                                            case "foo bar":
                                                                var v65 uint8 = 130
                                                                _ = v65
                                                            }
                                                            v22 += (v20 * 0.0)
                                                        }
                                                    }
                                                    v66 := GenS2{F0: (fi19 >> ((v3) & 63)), F1: v2}
                                                    var v67 uint64 = 42
                                                    _ = v50
                                                    _ = v66
                                                    _ = v67
                                                }
                                            default:
                                                var v68 uint = 18446744073709551615
                                                v69 := GenS2{F0: (-(fi21)), F1: v2}
                                                switch (fi21 % ((fi21) | 1)) {
                                                case -1:
                                                    v13 = v2
                                                }
                                                _ = v68
                                                _ = v69
                                            }
                                            v70 := GenS7{F0: (v2 + (v13 ^ v13))}
                                            _ = v70
                                        }
                                        if (1.5 == (v20 + (((1e10 * 0.0) - 1e10) - (1e10 - ((v22 / ((v22 / v22) * v22)) / 1.5))))) {
                                            switch fi19 {
                                            case 9223372036854775807:
                                                fmt.Println(v12, v43, v12)
                                            case 32768:
                                                var v71 string = (("foo bar" + ("hello" + "\n\t")) + ("foo bar" + ""))
                                                _ = v71
                                            default:
                                                switch (("hello" + ("foo bar" + "\n\t")) + ("foo bar" + "hello")) {
                                                case "\n\t":
                                                    fmt.Println(v15)
                                                case "a":
                                                    var v72 bool = false
                                                    v1[(v15 % ((65535) | 1))] = "a"
                                                    _ = v72
                                                }
                                            }
                                            switch "hello" {
                                            case "hello":
                                                if true {
                                                    fmt.Println(v2)
                                                    var v73 uint64 = (v43 % (((v43 - v43)) | 1))
                                                    v74 := v1[42]
                                                    _ = v73
                                                    _ = v74
                                                }
                                            case "a":
                                                v75 := GenS8{F0: (v12 * v12), F1: fi21, F2: -277}
                                                _ = v75
                                            case "\n\t":
                                                switch (fi19 >> ((v3) & 63)) {
                                                case -744:
                                                    var v76 int64 = -246
                                                    _ = v76
                                                case 977:
                                                    var v77 float64 = 1.5
                                                    _ = v77
                                                default:
                                                    var v78 string = "a"
                                                    _ = v78
                                                }
                                            default:
                                                var v79 float32 = (0.1 + float32((v22 - 1e10)))
                                                _ = v79
                                            }
                                            v2 = v13
                                        }
                                    }
                                    if (!((v2 % (((^(v13))) | 1)) >= v2)) {
                                        if false {
                                            fmt.Println(v43, fi19, fi21)
                                            fmt.Println(v12, v15)
                                            v3 -= v3
                                        } else {
                                            switch "" {
                                            case "":
                                                v13 = (v2 + -96)
                                            case "a":
                                                switch ("a" + "foo bar") {
                                                case "\n\t":
                                                    fmt.Println(v43, v43)
                                                }
                                            default:
                                                _ = len(v1)
                                            }
                                            switch fi21 {
                                            case -1:
                                                var v80 int8 = int8(v43)
                                                _ = v80
                                            case 1:
                                                v12 = v12
                                            case 32768:
                                                if (false || false) {
                                                    if (!((float32(v20) > (1e10 * 1e10)) || false)) {
                                                        fmt.Println(v43, v2, v13)
                                                        delete(v1, 153)
                                                        switch "" {
                                                        case "":
                                                            switch (fi21 | 128) {
                                                            case -9223372036854775808:
                                                                fmt.Println(v22)
                                                            case 42:
                                                                v20 = (v22 / ((v22 + 3.14159) * v20))
                                                            default:
                                                                var v81 string = "a"
                                                                _ = v81
                                                            }
                                                        case "\n\t":
                                                            v82 := []int{-796, 127, 4294967295}
                                                            _ = v82
                                                        default:
                                                            switch fi21 {
                                                            case 127:
                                                                v13 = int8(fi21)
                                                            }
                                                        }
                                                        v83 := GenS2{F0: 2147483647, F1: 127}
                                                        _ = v83
                                                    }
                                                } else {
                                                    var v84 uint8 = 253
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if false {
                                                            v85 := GenS3{F0: (v3 >> (((^(v3))) & 63)), F1: 205, F2: "hello"}
                                                            _ = v85
                                                        } else {
                                                            fmt.Println(v2)
                                                            var v86 uint16 = 2
                                                            var v87 uint8 = 0
                                                            var v88 float32 = 0.0
                                                            v89 := GenS9{F0: v15, F1: 4294967295}
                                                            _ = v86
                                                            _ = v87
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                    case "":
                                                        v15 += v15
                                                    }
                                                    _ = v84
                                                }
                                            }
                                        }
                                    } else {
                                        var v90 bool = true
                                        v15 = (v15 ^ v15)
                                        v91 := GenS3{F0: (^(v3)), F1: 255, F2: "foo bar"}
                                        _ = v90
                                        _ = v91
                                    }
                                    _ = v43
                                }
                                _ = v39
                            }
                        }
                        v12 -= (^(v12))
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    _ = v20
                } else {
                    v92 := GenS7{F0: -72}
                    _ = v92
                }
                switch ("a" + ("hello" + "")) {
                case "a":
                    v12 = v12
                case "hello":
                    v93 := GenS2{F0: (fi19 * (fi19 | (fi19 | (fi19 + (fi19 % (((fi19 - 636)) | 1)))))), F1: (v13 & v2)}
                    _ = v93
                case "foo bar":
                    var v94 int32 = genFn0(2147483647)
                    _ = v94
                }
                if (!true) {
                    var v95 bool = false
                    for fi96 := 0; fi96 < 9; fi96++ {
                        v97 := GenS10{F0: (genFn0((v15 / ((486) | 1))) != 1), F1: v12}
                        v2 = 70
                        v98 := v1[(v15 ^ uint32(v3))]
                        _ = fi96
                        _ = v97
                        _ = v98
                    }
                    _ = v95
                } else {
                    var v99 int64 = -9223372036854775808
                    _ = v99
                }
                _ = fi19
            }
        }
        v100 := []int{4294967295, 69, -2147483648, 863}
        _ = v100
    } else {
        if ((true || true) || true) {
            v15 = uint32(v3)
            delete(v1, uint32(v12))
        }
        for fi101 := 0; fi101 < 1; fi101++ {
            var v102 uint8 = v12
            _ = fi101
            _ = v102
        }
        v103 := GenS11{F0: 365}
        _ = v103
    }
    fmt.Println(v13, v2)
    fmt.Println(v2, v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
