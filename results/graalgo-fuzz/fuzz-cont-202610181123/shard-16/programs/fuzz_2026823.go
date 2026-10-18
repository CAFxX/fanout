package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint; F1 int64; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 string; F1 uint; F2 int8 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 int16 }
func genFn0(p0 int8, p1 uint32) uint8 {
    if true {
        v6 := GenS1{F0: (^(p0))}
        if (955 != 639) {
            v7 := GenS1{F0: p0}
            p0 += p0
            var v8 int32 = 1
            _ = v7
            _ = v8
        } else {
            switch 406 {
            case 42:
                v9 := []int{4294967295, -3, 0, 861}
                _ = v9
            case 32767:
                switch "foo bar" {
                case "":
                    for fi10 := 0; fi10 < 16; fi10++ {
                        var v11 uint8 = 255
                        v12 := GenS2{F0: (^(p1))}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case "\n\t":
                    v13 := []int{}
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 1; fi14++ {
                        fmt.Println(p0, p0, fi14)
                        v15 := []int{fi14, fi14}
                        switch "hello" {
                        case "hello":
                            for fi16 := 0; fi16 < 2; fi16++ {
                                p1 = (^(p1))
                                _ = fi16
                            }
                        default:
                            switch (-(fi14)) {
                            case 205:
                                fmt.Println(p0)
                            case -1:
                                var v17 int64 = -1
                                _ = v17
                            }
                        }
                        _ = fi14
                        _ = v15
                    }
                }
            }
            for fi18 := 0; fi18 < 19; fi18++ {
                p1 = p1
                _ = fi18
            }
            if (-128 == p0) {
                for fi19 := 0; fi19 < 1; fi19++ {
                    switch "" {
                    case "a":
                        v20 := GenS3{F0: -142}
                        _ = v20
                    case "":
                        v21 := GenS4{F0: 0, F1: 42}
                        _ = v21
                    case "foo bar":
                        if (67 != 18446744073709551615) {
                            var v22 int64 = 42
                            _ = v22
                        }
                    default:
                        for fi23 := 0; fi23 < 17; fi23++ {
                            v24 := []int{fi23, (fi23 % ((fi19) | 1))}
                            _ = len(v24)
                            v24 = append(v24, 9223372036854775807)
                            _ = fi23
                            _ = v24
                        }
                    }
                    var v25 int32 = 0
                    _ = fi19
                    _ = v25
                }
            }
        }
        _ = v6
    }
    var v26 float64 = (0.1 * (((3.14159 + ((((float64(float32(-2.25)) - 1.5) - float64(1.5)) - (1e10 + 1e10)) - float64(1.5))) * 1.5) * (1e10 - float64(((0.0 - 1.5) - float32((1e10 / 3.14159)))))))
    var v27 uint16 = 453
    for fi28 := 0; fi28 < 19; fi28++ {
        if false {
            switch fi28 {
            case 128:
                v29 := GenS4{F0: 251, F1: -9223372036854775808}
                _ = v29
            case 32768:
                switch ("" + "\n\t") {
                case "hello":
                    for fi30 := 0; fi30 < 12; fi30++ {
                        switch fi30 {
                        case 994:
                            var v31 int = (-(fi30))
                            _ = v31
                        case 32767:
                            var v32 uint = 0
                            _ = v32
                        case 396:
                            fmt.Println(p1, p0, v26)
                        }
                        fmt.Println(p0)
                        if false {
                            v33 := GenS5{F0: 985, F1: -1, F2: 2}
                            _ = v33
                        } else {
                            if false {
                                var v34 int16 = 1000
                                if (1 >= 42) {
                                    switch "foo bar" {
                                    case "foo bar":
                                        if true {
                                            if false {
                                                fmt.Println(fi28, p1, p0)
                                                var v35 float64 = v26
                                                _ = v35
                                            }
                                            var v36 int64 = 9223372036854775807
                                            var v37 string = ""
                                            _ = v36
                                            _ = v37
                                        } else {
                                            v38 := GenS2{F0: (p1 & p1)}
                                            switch "foo bar" {
                                            case "a":
                                                v26 += v26
                                                v34 = (v34 ^ -32768)
                                            case "hello":
                                                v34 -= (v34 | (v34 >> ((0) & 63)))
                                            case "foo bar":
                                                v39 := GenS6{F0: 65535}
                                                _ = v39
                                            default:
                                                switch (fi28 * (fi28 - (fi28 - (fi28 & fi30)))) {
                                                case 32767:
                                                    fmt.Println(fi28)
                                                }
                                            }
                                            v40 := []int{367, fi30, -9223372036854775808, fi28}
                                            _ = v38
                                            _ = v40
                                        }
                                    case "\n\t":
                                        v41 := GenS2{F0: p1}
                                        _ = v41
                                    }
                                    switch "" {
                                    case "\n\t":
                                        if (true && false) {
                                            var v42 float32 = 0.1
                                            _ = v42
                                        } else {
                                            var v43 bool = true
                                            _ = v43
                                        }
                                    case "":
                                        v27 = v27
                                    case "a":
                                        v44 := GenS4{F0: 2, F1: 1}
                                        _ = v44
                                    default:
                                        v45 := []int{-1, -528}
                                        _ = v45
                                    }
                                    fmt.Println(v26, fi28, fi28)
                                }
                                var v46 string = (((("hello" + "hello") + "a") + ("hello" + ("hello" + (("a" + ("hello" + ("" + "foo bar"))) + ("foo bar" + ""))))) + "\n\t")
                                _ = v34
                                _ = v46
                            } else {
                                if true {
                                    switch "hello" {
                                    case "foo bar":
                                        v47 := GenS6{F0: 369}
                                        _ = v47
                                    case "a":
                                        v48 := []int{}
                                        _ = v48
                                    case "":
                                        if true {
                                            var v49 int8 = (^(p0))
                                            if (true && (42 <= 666)) {
                                                v26 = (v26 + 0.1)
                                                v27 = 32767
                                            }
                                            v50 := []int{-1, fi30, (fi30 & (fi28 >> ((901) & 63)))}
                                            _ = v49
                                            _ = v50
                                        }
                                    default:
                                        fmt.Println(fi28, p0)
                                    }
                                } else {
                                    v51 := []int{729, (-(fi28)), (fi30 ^ fi30), (fi30 | fi28)}
                                    switch (-(fi28)) {
                                    case 4294967295:
                                        v26 = (v26 + (v26 * v26))
                                    case 127:
                                        if true {
                                            p0 -= (p0 + p0)
                                        }
                                    }
                                    _ = v51
                                }
                                p1 = p1
                                fmt.Println(p0, fi28, v26)
                            }
                        }
                        switch -739 {
                        case 560:
                            v26 += -2.25
                        case 2147483647:
                            if ((false && true) && ((18446744073709551615 >= 359) || (-449 == 1))) {
                                v52 := []int{(fi30 >> ((355) & 63)), (fi28 & 32768), (fi28 - (fi30 % ((-1) | 1))), -113}
                                if false {
                                    if (fi30 == int(p1)) {
                                        p0 -= 1
                                        var v53 int16 = 42
                                        _ = v53
                                    }
                                    var v54 int32 = -446
                                    if (v26 < 3.14159) {
                                        var v55 string = "foo bar"
                                        _ = v55
                                    } else {
                                        if true {
                                            v56 := GenS7{F0: 1, F1: -2147483648}
                                            _ = v56
                                        } else {
                                            v26 = -2.25
                                            switch "hello" {
                                            case "a":
                                                var v57 int32 = (v54 % (((v54 % ((-1) | 1))) | 1))
                                                _ = v57
                                            case "foo bar":
                                                fmt.Println(v27)
                                                if true {
                                                    if false {
                                                        var v58 int16 = 391
                                                        v59 := GenS6{F0: 2}
                                                        var v60 uint8 = 127
                                                        _ = v58
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        v52 = append(v52, (fi30 * 0))
                                                        var v61 int64 = 638
                                                        _ = v61
                                                    }
                                                }
                                                _ = len(v52)
                                            case "":
                                                var v62 float64 = ((v26 * v26) / v26)
                                                _ = v62
                                            }
                                            if (1 == 1) {
                                                if false {
                                                    var v63 int64 = -9223372036854775808
                                                    _ = v63
                                                } else {
                                                    switch "" {
                                                    case "":
                                                        var v64 uint64 = 4294967295
                                                        _ = v64
                                                    case "hello":
                                                        var v65 int8 = (p0 % ((p0) | 1))
                                                        _ = v65
                                                    }
                                                }
                                                var v66 uint = 0
                                                _ = v66
                                            }
                                        }
                                        v67 := GenS2{F0: (p1 | (p1 & p1))}
                                        _ = v67
                                    }
                                    _ = v54
                                }
                                _ = v52
                            }
                        case 32768:
                            var v68 int64 = 42
                            if false {
                                v27 += (v27 - v27)
                                v69 := []int{}
                                _ = v69
                            }
                            var v70 int64 = 42
                            _ = v68
                            _ = v70
                        default:
                            fmt.Println(fi28)
                            fmt.Println(v26)
                            if (false && false) {
                                v27 -= (v27 >> ((966) & 63))
                                var v71 int8 = p0
                                _ = v71
                            }
                        }
                        _ = fi30
                    }
                case "":
                    var v72 uint32 = (p1 - (p1 * (p1 ^ (p1 + uint32(v27)))))
                    _ = v72
                case "\n\t":
                    v73 := GenS8{F0: -30, F1: 0}
                    _ = v73
                default:
                    var v74 uint = 4294967295
                    _ = v74
                }
            default:
                p1 += 42
            }
            v75 := []int{(fi28 / (((fi28 << ((107) & 63))) | 1))}
            switch "\n\t" {
            case "\n\t":
                if true {
                    for fi76 := 0; fi76 < 17; fi76++ {
                        if len(v75) > 0 {
                            v75[0] = fi28
                        }
                        switch "" {
                        case "a":
                            var v77 bool = false
                            _ = v77
                        }
                        switch fi28 {
                        case 127:
                            var v78 int32 = 1
                            _ = v78
                        case 32767:
                            var v79 uint = 4294967295
                            _ = v79
                        case 635:
                            switch ((("" + ("hello" + "")) + "a") + "\n\t") {
                            case "a":
                                switch fi76 {
                                case 2147483647:
                                    var v80 uint32 = p1
                                    _ = v80
                                default:
                                    v27 = (v27 % (((^(v27))) | 1))
                                }
                            case "foo bar":
                                p1 = p1
                            case "\n\t":
                                switch "\n\t" {
                                case "a":
                                    if len(v75) > 0 {
                                        v75[0] = fi28
                                    }
                                default:
                                    _ = len(v75)
                                }
                            }
                        default:
                            v81 := GenS4{F0: 0, F1: 42}
                            _ = v81
                        }
                        _ = fi76
                    }
                } else {
                    v82 := GenS9{F0: "\n\t", F1: 2, F2: (p0 | 0)}
                    v83 := GenS2{F0: 104}
                    _ = v82
                    _ = v83
                }
            }
            _ = v75
        }
        for fi84 := 0; fi84 < 5; fi84++ {
            v85 := []int{(fi84 * (fi28 * int(p0)))}
            var v86 uint = 164
            _ = fi84
            _ = v85
            _ = v86
        }
        _ = fi28
    }
    _ = p0
    _ = p1
    _ = v26
    _ = v27
    return 127
}

func main() {
    var v1 bool = ((false || (false && (0 < 4294967295))) && (432 >= -9223372036854775808))
    var v2 bool = (42 < 1)
    var v3 bool = v1
    v4 := []string{("" + "\n\t"), "", ("\n\t" + "\n\t"), "hello"}
    v5 := make(map[int]int)
    v3 = false
    var v87 uint16 = 377
    if v2 {
        for fi88 := 0; fi88 < 4; fi88++ {
            _ = len(v5)
            _ = fi88
        }
    }
    if true {
        fmt.Println(v1, v3, v2)
    }
    switch "\n\t" {
    case "hello":
        var v89 int64 = 1
        _ = v89
    }
    fmt.Println(v3, v87, v3)
    if (255 <= v87) {
        var v90 uint = 4294967295
        for fi91 := 0; fi91 < 12; fi91++ {
            var v92 int64 = -213
            if (0 != 0) {
                v4 = append(v4, (("hello" + "") + ((("a" + "a") + "hello") + ("a" + "foo bar"))))
                var v93 int = -361
                var v94 int32 = 42
                _ = v93
                _ = v94
            }
            _ = fi91
            _ = v92
        }
        if (v87 < v87) {
            var v95 uint8 = uint8(v90)
            if v2 {
                var v96 uint16 = (v87 & (v87 / ((v87) | 1)))
                switch 4294967295 {
                case 42:
                    for fi97 := 0; fi97 < 19; fi97++ {
                        if (genFn0(100, 0) <= 128) {
                            v3 = (true && v1)
                        } else {
                            if (v2 && (1e10 <= 3.14159)) {
                                switch fi97 {
                                case -741:
                                    if v1 {
                                        var v98 uint8 = (v95 - 190)
                                        _ = v98
                                    } else {
                                        v96 = (v96 / (((v87 | (v96 / (((v96 / ((v96) | 1))) | 1)))) | 1))
                                        v3 = v2
                                    }
                                }
                            } else {
                                if ((100 != 100) && ((v90 & v90) > (v90 ^ (^(v90))))) {
                                    if false {
                                        v99 := GenS10{F0: (fi97 + fi97)}
                                        var v100 int32 = -2147483648
                                        fmt.Println(v100)
                                        _ = v100
                                        _ = v99
                                    }
                                } else {
                                    if (v2 || v1) {
                                        fmt.Println(v1)
                                    }
                                    v90 = v90
                                    fmt.Println(fi97)
                                }
                                fmt.Println(v87)
                                switch "\n\t" {
                                case "hello":
                                    var v101 float64 = 1e10
                                    _ = v101
                                case "foo bar":
                                    for fi102 := 0; fi102 < 9; fi102++ {
                                        if (0 <= 208) {
                                            fmt.Println(v2, v95, v3)
                                        }
                                        fmt.Println(v2, v2, fi97)
                                        _ = fi102
                                    }
                                }
                            }
                            for fi103 := 0; fi103 < 5; fi103++ {
                                if ((true && (2 <= 4294967295)) || v2) {
                                    v3 = (!(1 == 122))
                                } else {
                                    _ = len(v5)
                                }
                                _ = fi103
                            }
                        }
                        _ = len(v5)
                        _ = fi97
                    }
                case 2147483647:
                    v2 = ((350 == 4294967295) || ((^(v95)) < (v95 | (^(v95)))))
                case 32767:
                    var v104 uint32 = 4294967295
                    _ = v104
                }
                var v105 uint64 = 795
                _ = v105
                _ = v96
            } else {
                var v106 int64 = 9223372036854775807
                _ = v106
            }
            _ = v95
        }
        _ = v90
    } else {
        v3 = (!(0 >= -503))
    }
    var v107 int8 = -128
    if (!true) {
        _ = len(v5)
        v87 = v87
        if (1.5 < 1e10) {
            v108 := GenS11{F0: -208}
            var v109 float64 = -2.25
            v4 = append(v4, "hello")
            _ = v108
            _ = v109
        }
    } else {
        var v110 bool = (42 != 1)
        var v111 float32 = (1.5 / 1.5)
        v112 := GenS8{F0: 707, F1: 18446744073709551615}
        _ = v110
        _ = v111
        _ = v112
    }
    v113 := GenS7{F0: 42, F1: 127}
    var v114 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v107:", v107)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v87:", v87)
    _ = v1
    _ = v107
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v87
}
