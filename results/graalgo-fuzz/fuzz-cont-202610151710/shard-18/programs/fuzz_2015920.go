package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int32; F2 int64 }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 int64; F1 uint8; F2 bool }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int8; F1 int64; F2 int16 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 uint32; F2 int64 }
type GenS10 struct { F0 uint8; F1 uint64; F2 string }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int16; F1 int16; F2 int32 }
type GenS13 struct { F0 uint32; F1 uint }
type GenS14 struct { F0 string; F1 int64 }
func genFn0(p0 uint8, p1 uint) uint16 {
    switch "a" {
    case "hello":
        for fi4 := 0; fi4 < 6; fi4++ {
            fmt.Println(p1)
            var v5 string = "foo bar"
            v6 := []int{}
            _ = fi4
            _ = v5
            _ = v6
        }
    case "a":
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := []int{(fi7 / ((-1) | 1)), (fi7 - int(p0))}
            var v9 uint32 = 86
            p0 += 128
            _ = fi7
            _ = v8
            _ = v9
        }
    case "":
        switch -2147483648 {
        case -888:
            var v10 string = "foo bar"
            _ = v10
        }
    }
    v11 := []int{}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v11
    return 1
}
func genFn1(p0 uint32) string {
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 float32 = 0.0
        var v14 uint8 = 1
        var v15 int16 = 1
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int8 = -128
    _ = p0
    _ = v16
    return "foo bar"
}

func main() {
    v1 := []int{}
    var v2 float32 = float32(((0.0 + (1.5 - float64(1e10))) * 3.14159))
    var v3 int32 = -2147483648
    fmt.Println(v3)
    fmt.Println(v2, v2)
    var v17 uint32 = 2
    v18 := GenS1{F0: -1}
    if (false && (v3 > v3)) {
        var v19 uint16 = 42
        _ = v19
    } else {
        v20 := GenS2{F0: 22, F1: v17}
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 uint16 = 255
            switch (fi21 / (((fi21 >> ((1) & 63))) | 1)) {
            case 735:
                v22 = (v22 % ((v22) | 1))
            default:
                if true {
                    var v23 string = ""
                    var v24 float32 = v2
                    var v25 float64 = (3.14159 / 0.1)
                    _ = v23
                    _ = v24
                    _ = v25
                }
            }
            fmt.Println(v22, v22, v3)
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    switch (("foo bar" + "hello") + "") {
    case "foo bar":
        var v26 float64 = (0.1 - 3.14159)
        _ = v26
    case "a":
        v17 = (v17 | 0)
    case "":
        v27 := GenS3{F0: 2147483648, F1: (v3 - v3), F2: 42}
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        fmt.Println(fi28, v3)
        for fi29 := 0; fi29 < 10; fi29++ {
            if (527 == 520) {
                v17 = 947
                var v30 uint8 = 127
                var v31 float32 = 1e10
                _ = v30
                _ = v31
            } else {
                var v32 int32 = -1
                if (-629 > 1) {
                    v32 = v32
                    fmt.Println(fi29)
                    fmt.Println(fi29, fi28)
                } else {
                    switch "\n\t" {
                    case "hello":
                        switch fi29 {
                        case 127:
                            v3 = v3
                        case 271:
                            v33 := GenS4{F0: -1, F1: v32}
                            _ = v33
                        case -764:
                            fmt.Println(fi28, v2, fi28)
                        }
                    }
                }
                _ = v32
            }
            if (255 < 42) {
                fmt.Println(fi28)
                fmt.Println(v3)
            }
            var v34 float32 = 0.0
            switch "foo bar" {
            case "hello":
                switch -944 {
                case -2147483648:
                    v35 := GenS5{F0: 0, F1: 0, F2: (1 < 42)}
                    _ = v35
                case 127:
                    var v36 int8 = 56
                    _ = v36
                case -154:
                    var v37 float32 = float32((3.14159 / 1.0))
                    _ = v37
                }
            case "a":
                if (-2.25 < (3.14159 * 1.5)) {
                    if ((float32(-2.25) / float32(((-2.25 + (1.5 * 1e10)) - -2.25))) == ((v2 - v34) / v34)) {
                        var v38 int = 540
                        var v39 uint = 477
                        v3 = v3
                        _ = v38
                        _ = v39
                    } else {
                        if ("a" >= (("\n\t" + ("\n\t" + "foo bar")) + "foo bar")) {
                            v40 := GenS4{F0: 32767, F1: v3}
                            _ = v40
                        }
                        v3 = (v3 ^ v3)
                    }
                }
            case "":
                if (!false) {
                    switch (fi29 >> ((1) & 63)) {
                    case -379:
                        var v41 uint8 = 130
                        _ = v41
                    }
                    if len(v1) > 0 {
                        v1[0] = fi28
                    }
                    var v42 uint = 2
                    if (4294967295 >= fi28) {
                        var v43 bool = (genFn0(128, 255) >= 255)
                        switch "" {
                        case "a":
                            v44 := GenS1{F0: 0}
                            _ = v44
                        case "\n\t":
                            if false {
                                if true {
                                    if true {
                                        var v45 string = ("a" + "")
                                        _ = v45
                                    } else {
                                        switch 649 {
                                        case -9223372036854775808:
                                            fmt.Println(v3)
                                        case -288:
                                            if (127 <= uint8(v17)) {
                                                v17 += (v17 * v17)
                                                var v46 uint8 = 2
                                                v1 = append(v1, int(v46))
                                                var v47 int = 32768
                                                var v48 uint32 = v17
                                                _ = v46
                                                _ = v47
                                                _ = v48
                                            }
                                        case -1:
                                            if false {
                                                if (!(1 < -9223372036854775808)) {
                                                    v42 += 18446744073709551615
                                                    var v49 float32 = ((-2.25 * v2) * 0.0)
                                                    if len(v1) > 0 {
                                                        v1[0] = -2147483648
                                                    }
                                                    _ = v49
                                                } else {
                                                    fmt.Println(v42, v3)
                                                }
                                                fmt.Println(v43)
                                            } else {
                                                var v50 bool = (42 >= 1)
                                                var v51 int8 = -2
                                                if v50 {
                                                    var v52 int8 = (v51 % ((v51) | 1))
                                                    v2 -= v34
                                                    _ = v52
                                                }
                                                _ = v50
                                                _ = v51
                                            }
                                        default:
                                            if v43 {
                                                v17 -= (v17 * (v17 * (v17 >> (((v42 ^ (v42 + v42))) & 63))))
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v53 := GenS6{F0: 179, F1: 4294967295}
                                                var v54 int = (fi28 | 128)
                                                _ = v53
                                                _ = v54
                                            } else {
                                                fmt.Println(v17)
                                                v55 := GenS5{F0: 1, F1: 142, F2: (v17 > v17)}
                                                _ = v55
                                            }
                                        }
                                        _ = len(v1)
                                        switch (("hello" + (("hello" + ("a" + "a")) + ("foo bar" + ("\n\t" + "hello")))) + ("foo bar" + "foo bar")) {
                                        case "foo bar":
                                            var v56 int = (fi28 << ((v42) & 63))
                                            _ = v56
                                        }
                                    }
                                    switch ("" + ("\n\t" + "\n\t")) {
                                    case "hello":
                                        fmt.Println(v34)
                                    default:
                                        var v57 string = ("" + "foo bar")
                                        _ = v57
                                    }
                                } else {
                                    _ = len(v1)
                                    var v58 int8 = -65
                                    _ = v58
                                }
                                var v59 string = "hello"
                                var v60 string = (((("hello" + v59) + v59) + v59) + "a")
                                _ = v59
                                _ = v60
                            }
                        case "hello":
                            _ = len(v1)
                            fmt.Println(fi29, v42, v3)
                        default:
                            v43 = true
                        }
                        _ = v43
                    }
                    v1 = append(v1, fi29)
                    v42 -= (^(v42))
                    var v61 int32 = (v3 + 1)
                    v62 := GenS7{F0: 42, F1: 759, F2: 1000}
                    var v63 int8 = -113
                    _ = v42
                    _ = v61
                    _ = v62
                    _ = v63
                }
                var v64 int8 = 94
                _ = v64
            }
            fmt.Println(fi28)
            switch "\n\t" {
            case "a":
                var v65 bool = true
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v65
            case "hello":
                v34 = (float32(1.5) * (((float32(1.5) - float32(0.0)) / -2.25) * v2))
                v66 := GenS8{F0: true}
                var v67 uint = 18446744073709551615
                if true {
                    fmt.Println(v2, v34, v67)
                }
                switch (fi29 - (fi28 + fi28)) {
                case 0:
                    v68 := GenS9{F0: 32767, F1: 1, F2: int64(fi29)}
                    var v69 uint64 = 1
                    _ = v68
                    _ = v69
                case -9223372036854775808:
                    var v70 uint = 4294967295
                    if true {
                        v3 = 42
                        switch fi29 {
                        case 128:
                            var v71 float32 = v2
                            _ = v71
                        case -2147483648:
                            if true {
                                switch "" {
                                case "hello":
                                    if false {
                                        if (2 >= 4294967295) {
                                            if ((^(v17)) == (v17 & 701)) {
                                                v72 := GenS10{F0: 42, F1: 386, F2: "a"}
                                                _ = v72
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = 30
                                            }
                                        } else {
                                            switch (("\n\t" + "a") + "hello") {
                                            case "foo bar":
                                                switch "" {
                                                case "hello":
                                                    if ((v67 ^ 326) != v67) {
                                                        var v73 int32 = 36
                                                        v17 = (^(v17))
                                                        _ = v73
                                                    } else {
                                                        var v74 string = "hello"
                                                        var v75 string = "a"
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                case "a":
                                                    if false {
                                                        if false {
                                                            var v76 int64 = 0
                                                            v76 -= 553
                                                            var v77 string = "foo bar"
                                                            _ = v76
                                                            _ = v77
                                                        } else {
                                                            v34 += float32(-2.25)
                                                            fmt.Println(v17, fi29)
                                                            v70 = (v67 | (v67 | v70))
                                                        }
                                                        var v78 uint16 = 0
                                                        _ = v78
                                                    }
                                                }
                                            case "a":
                                                v67 = v67
                                                switch "foo bar" {
                                                case "hello":
                                                    v2 = v34
                                                    fmt.Println(fi28, fi29, v70)
                                                    if len(v1) > 0 {
                                                        v1[0] = -209
                                                    }
                                                case "a":
                                                    switch "\n\t" {
                                                    case "":
                                                        fmt.Println(v70, v2)
                                                    case "a":
                                                        if (v17 > (^(v17))) {
                                                            if (!true) {
                                                                var v79 string = "\n\t"
                                                                if (!true) {
                                                                    var v80 int64 = 1
                                                                    _ = v80
                                                                }
                                                                _ = v79
                                                            }
                                                            _ = len(v1)
                                                            if ((v17 == (v17 ^ (^(v17)))) || ((^(v70)) <= 1)) {
                                                                _ = len(v1)
                                                                var v81 int8 = -58
                                                                _ = v81
                                                            }
                                                        }
                                                    }
                                                    v34 = v34
                                                    switch "" {
                                                    case "foo bar":
                                                        v2 += 0.1
                                                    case "a":
                                                        if false {
                                                            if (v67 == v67) {
                                                                var v82 uint16 = 255
                                                                _ = v82
                                                            }
                                                            v83 := GenS11{F0: ""}
                                                            var v84 int64 = 9223372036854775807
                                                            _ = v83
                                                            _ = v84
                                                        }
                                                    default:
                                                        var v85 int16 = 42
                                                        _ = v85
                                                    }
                                                case "":
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                _ = len(v1)
                                            case "hello":
                                                if true {
                                                    var v86 float32 = v34
                                                    _ = v86
                                                } else {
                                                    var v87 uint32 = 42
                                                    _ = v87
                                                }
                                            }
                                            fmt.Println(v70, v3)
                                        }
                                    } else {
                                        switch 1 {
                                        case 73:
                                            _ = len(v1)
                                        default:
                                            if true {
                                                var v88 uint = (v67 % ((v67) | 1))
                                                var v89 uint64 = 0
                                                var v90 int16 = 0
                                                _ = v88
                                                _ = v89
                                                _ = v90
                                            } else {
                                                v2 = v34
                                                v91 := GenS12{F0: -81, F1: 1000, F2: int32(v17)}
                                                v34 = v34
                                                _ = v91
                                            }
                                        }
                                    }
                                }
                                var v92 int8 = -128
                                _ = v92
                            }
                            var v93 uint8 = 0
                            _ = v93
                        case 887:
                            var v94 int64 = 0
                            _ = v94
                        }
                        switch "hello" {
                        case "hello":
                            var v95 float64 = 1e10
                            _ = v95
                        default:
                            var v96 uint8 = 1
                            _ = v96
                        }
                        if (32 > -1) {
                            v97 := GenS13{F0: (^(v17)), F1: (^(v67))}
                            v3 += v3
                            var v98 int16 = -518
                            _ = v97
                            _ = v98
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        switch ("foo bar" + "foo bar") {
                        case "a":
                            fmt.Println(fi28)
                        case "":
                            v17 = v17
                        case "foo bar":
                            switch 9223372036854775807 {
                            case 9223372036854775807:
                                v70 = (v67 - v70)
                            case 32768:
                                v1 = append(v1, fi29)
                            case -2147483648:
                                fmt.Println(v34, v17)
                            default:
                                v34 -= v2
                            }
                        default:
                            v99 := GenS14{F0: "hello", F1: 1}
                            _ = v99
                        }
                        var v100 int64 = 1
                        _ = v100
                    }
                    var v101 uint16 = 589
                    _ = v101
                    _ = v70
                }
                _ = v66
                _ = v67
            default:
                fmt.Println(v34, fi28, v17)
            }
            v3 -= (v3 | (v3 / ((1) | 1)))
            _ = fi29
            _ = v34
        }
        for fi102 := 0; fi102 < 1; fi102++ {
            var v103 uint16 = 1
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi102
            _ = v103
        }
        _ = fi28
    }
    v3 += 0
    v3 -= 42
    fmt.Println(v2, v17, v17)
    v17 -= 42
    for fi104 := 0; fi104 < 19; fi104++ {
        var v105 int32 = (v3 ^ (v3 ^ -1))
        _ = fi104
        _ = v105
    }
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v106 := GenS1{F0: 907}
        _ = v106
    }
    v1 = append(v1, -394)
    if (!true) {
        var v107 int16 = 1
        _ = v107
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
