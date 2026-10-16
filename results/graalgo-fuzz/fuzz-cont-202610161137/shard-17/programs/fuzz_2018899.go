package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 uint32; F1 bool; F2 string }
type GenS5 struct { F0 int16; F1 int16; F2 int8 }
type GenS6 struct { F0 uint32; F1 int64 }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint16; F1 int64 }
type GenS10 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int32; F1 int8 }
type GenS13 struct { F0 string; F1 int64 }
type GenS14 struct { F0 int64; F1 uint; F2 int }
type GenS15 struct { F0 uint32 }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int16; F1 uint }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint8 }
func genFn0(p0 string, p1 uint32, p2 uint32) int {
    fmt.Println(p2)
    fmt.Println(p1, p0, p1)
    v4 := GenS1{F0: 1}
    if ((!(!false)) || true) {
        fmt.Println(p2, p0)
        var v5 string = ("foo bar" + "hello")
        var v6 uint8 = 150
        _ = v5
        _ = v6
    } else {
        fmt.Println(p1, p0)
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := []int{(fi7 * 128), 551, (fi7 + fi7)}
            _ = fi7
            _ = v8
        }
    }
    var v9 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return -76
}

func main() {
    var v1 uint8 = 2
    v2 := []uint32{2147483648, 79, 2147483648}
    v3 := make(map[uint32]int)
    var v10 float64 = float64(0.1)
    if ((float64(((0.0 - 0.1) * -2.25)) * -2.25) > 3.14159) {
        if false {
            switch 4294967295 {
            case -1:
                if (true || true) {
                    switch "" {
                    case "hello":
                        v10 = 1e10
                    }
                }
            case 32767:
                v1 += 1
            case 9223372036854775807:
                fmt.Println(v10, v10)
            }
        } else {
            if (-1 >= 0) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    fmt.Println(v1, v10, v10)
                    var v12 bool = (((("hello" + "\n\t") + "\n\t") + "hello") > "a")
                    _ = fi11
                    _ = v12
                }
                v10 += 0.0
                if true {
                    v10 = ((3.14159 * 0.1) / 3.14159)
                    if len(v2) > 0 {
                        v2[0] = 2
                    }
                    v13 := GenS2{F0: 873, F1: (v1 ^ 128)}
                    _ = v13
                } else {
                    if (127 != -128) {
                        fmt.Println(v10, v1)
                    } else {
                        for fi14 := 0; fi14 < 14; fi14++ {
                            var v15 int32 = 42
                            v15 += 2147483647
                            v15 += int32(v1)
                            _ = fi14
                            _ = v15
                        }
                        v16 := GenS3{F0: true, F1: 65535}
                        _ = v16
                    }
                    v1 -= (v1 / (((v1 / (((^(v1))) | 1))) | 1))
                    switch -782 {
                    case -324:
                        v17 := GenS4{F0: 659, F1: (false || (0 >= 1)), F2: "\n\t"}
                        _ = v17
                    case -563:
                        if (!true) {
                            var v18 uint = 980
                            _ = v18
                        } else {
                            v19 := GenS5{F0: -32768, F1: 1000, F2: -4}
                            var v20 uint16 = 65535
                            var v21 uint = 341
                            _ = v19
                            _ = v20
                            _ = v21
                        }
                    }
                }
            }
            if ((2147483647 >= 790) || false) {
                if true {
                    fmt.Println(v10, v10)
                }
            } else {
                v22 := GenS2{F0: 0, F1: (v1 * v1)}
                _ = v22
            }
        }
    } else {
        if false {
            var v23 uint8 = (v1 >> ((65535) & 63))
            var v24 float64 = v10
            _ = v23
            _ = v24
        }
        switch "hello" {
        case "":
            for fi25 := 0; fi25 < 15; fi25++ {
                fmt.Println(v1)
                switch fi25 {
                case -9223372036854775808:
                    v2 = append(v2, 302)
                case 1:
                    for fi26 := 0; fi26 < 4; fi26++ {
                        switch "a" {
                        case "hello":
                            switch int(v1) {
                            case 42:
                                v27 := GenS4{F0: uint32(v1), F1: false, F2: ""}
                                _ = v27
                            case 32767:
                                if false {
                                    v1 -= 1
                                    switch ("" + "foo bar") {
                                    case "":
                                        var v28 float32 = 1e10
                                        _ = v28
                                    case "hello":
                                        v10 += v10
                                    }
                                    v10 = 0.1
                                } else {
                                    switch -1 {
                                    case 1:
                                        v29 := GenS6{F0: 1, F1: -619}
                                        _ = v29
                                    case -1:
                                        var v30 uint8 = 1
                                        _ = v30
                                    case 4294967295:
                                        var v31 uint8 = (v1 & v1)
                                        _ = v31
                                    default:
                                        if (("" + "\n\t") != "") {
                                            if false {
                                                fmt.Println(fi26)
                                                _ = len(v3)
                                                v10 = 3.14159
                                                switch ("\n\t" + "a") {
                                                case "foo bar":
                                                    v10 += v10
                                                    var v32 int32 = -2147483648
                                                    _ = v32
                                                default:
                                                    fmt.Println(fi25, v1, v10)
                                                }
                                                v10 -= ((0.1 + v10) * v10)
                                                fmt.Println(v1)
                                                switch "\n\t" {
                                                case "":
                                                    _ = len(v3)
                                                    fmt.Println(fi25, v1)
                                                    v1 += v1
                                                    v33 := GenS5{F0: 1000, F1: -250, F2: 1}
                                                    _ = v33
                                                default:
                                                    var v34 uint = 255
                                                    _ = v34
                                                }
                                                switch (fi25 / (((^(fi26))) | 1)) {
                                                case 0:
                                                    v1 = 223
                                                case 2147483647:
                                                    var v35 uint16 = 42
                                                    _ = v35
                                                case 32768:
                                                    if (!(true || (genFn0("a", 2147483647, 703) > (fi26 + -2147483648)))) {
                                                        v2 = append(v2, 2)
                                                        var v36 uint = uint(v1)
                                                        var v37 int64 = 1
                                                        if true {
                                                            var v38 uint = v36
                                                            var v39 int64 = -1
                                                            switch (fi26 % (((fi26 / (((^(fi25))) | 1))) | 1)) {
                                                            case -881:
                                                                var v40 uint64 = 0
                                                                _ = v40
                                                            case 0:
                                                                fmt.Println(v38, fi25, v38)
                                                            case 192:
                                                                var v41 string = ("\n\t" + "\n\t")
                                                                _ = v41
                                                            default:
                                                                fmt.Println(v1, v39, v36)
                                                            }
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                        switch (^(fi26)) {
                                                        case -9223372036854775808:
                                                            fmt.Println(v36, fi26, fi26)
                                                        default:
                                                            if false {
                                                                switch fi25 {
                                                                case 9223372036854775807:
                                                                    _ = len(v2)
                                                                case 42:
                                                                    delete(v3, 4294967295)
                                                                default:
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v42 := GenS6{F0: 2147483648, F1: (^(v37))}
                                                                        _ = v42
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v37 = v37
                                                        _ = v36
                                                        _ = v37
                                                    } else {
                                                        var v43 uint32 = 628
                                                        v10 -= (v10 / float64((float32((((float64((1e10 / 1.0)) * (float64((((1e10 - (0.0 + (1e10 - 1.5))) + -2.25) - -2.25)) * 0.1)) - (3.14159 / ((v10 / 1.5) / -2.25))) + -2.25)) + -2.25)))
                                                        _ = v43
                                                    }
                                                default:
                                                    if false {
                                                        switch "a" {
                                                        case "hello":
                                                            var v44 uint8 = (v1 % ((42) | 1))
                                                            _ = v44
                                                        case "":
                                                            var v45 float64 = 3.14159
                                                            var v46 float32 = -2.25
                                                            _ = v45
                                                            _ = v46
                                                        case "foo bar":
                                                            switch "" {
                                                            case "hello":
                                                                v3[4294967295] = fi25
                                                            case "foo bar":
                                                                v47 := GenS7{F0: 2147483648, F1: 225}
                                                                _ = v47
                                                            case "a":
                                                                if false {
                                                                    v48 := GenS8{F0: true}
                                                                    var v49 bool = ((^(v1)) != (v1 << ((255) & 63)))
                                                                    var v50 int32 = 0
                                                                    if v49 {
                                                                        v51 := GenS9{F0: 997, F1: 743}
                                                                        _ = v51
                                                                    }
                                                                    if (!(301 < 732)) {
                                                                        fmt.Println(fi25, v49, v10)
                                                                        var v52 int32 = (v50 >> ((65535) & 63))
                                                                        v10 -= v10
                                                                        v2 = append(v2, 42)
                                                                        if ((v52 / (((^(v50))) | 1)) != (v52 | 1)) {
                                                                            v49 = v49
                                                                            v49 = v49
                                                                            _ = len(v2)
                                                                        }
                                                                        _ = v52
                                                                    }
                                                                    v49 = v49
                                                                    if (v1 == (v1 * v1)) {
                                                                        switch genFn0("a", 2147483647, 42) {
                                                                        case 1:
                                                                            v49 = v49
                                                                        case 42:
                                                                            if (!true) {
                                                                                var v53 int64 = -168
                                                                                v49 = v49
                                                                                v54 := GenS4{F0: 2147483648, F1: false, F2: ("foo bar" + ((("\n\t" + "") + "") + ("" + "foo bar")))}
                                                                                _ = v53
                                                                                _ = v54
                                                                            } else {
                                                                                if (v10 == -2.25) {
                                                                                    v49 = v49
                                                                                    v55 := GenS2{F0: 42, F1: 2}
                                                                                    v10 -= 3.14159
                                                                                    _ = v55
                                                                                } else {
                                                                                    v50 -= v50
                                                                                    switch (fi26 / ((fi26) | 1)) {
                                                                                    case -2147483648:
                                                                                        if v49 {
                                                                                            switch (-(fi25)) {
                                                                                            case 9223372036854775807:
                                                                                                var v56 uint32 = 1
                                                                                                _ = v56
                                                                                            case 133:
                                                                                                if (!false) {
                                                                                                    delete(v3, 42)
                                                                                                    switch "a" {
                                                                                                    case "foo bar":
                                                                                                        fmt.Println(v1)
                                                                                                    }
                                                                                                    if v49 {
                                                                                                        var v57 float32 = float32((0.0 * float64((0.0 + 0.0))))
                                                                                                        _ = v57
                                                                                                    } else {
                                                                                                        var v58 string = ("hello" + "hello")
                                                                                                        _ = v58
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                            v10 -= -2.25
                                                                                            v10 += 3.14159
                                                                                        }
                                                                                    case -473:
                                                                                        v50 -= -2147483648
                                                                                    default:
                                                                                        if (!v49) {
                                                                                            if v49 {
                                                                                                if (42 >= 2) {
                                                                                                    if v49 {
                                                                                                        if v49 {
                                                                                                            v59 := v3[65535]
                                                                                                            _ = v59
                                                                                                        } else {
                                                                                                            v49 = (0.1 < 0.0)
                                                                                                            fmt.Println(v49, fi25)
                                                                                                            delete(v3, 42)
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v60 int16 = int16(fi25)
                                                                                                        v50 += v50
                                                                                                        _ = v60
                                                                                                    }
                                                                                                    v50 -= 914
                                                                                                    v50 = (v50 - -882)
                                                                                                }
                                                                                                var v61 uint32 = 42
                                                                                                _ = v61
                                                                                            }
                                                                                            v62 := GenS8{F0: v49}
                                                                                            fmt.Println(v1, fi26)
                                                                                            _ = v62
                                                                                        }
                                                                                    }
                                                                                    if v49 {
                                                                                        v63 := GenS9{F0: 32767, F1: 9223372036854775807}
                                                                                        var v64 uint64 = 4294967295
                                                                                        _ = v63
                                                                                        _ = v64
                                                                                    } else {
                                                                                        v65 := GenS10{F0: (v50 * (v50 >> ((4294967295) & 63))), F1: 286, F2: v1}
                                                                                        v1 += (v1 * (v1 | (v1 - (v1 & 164))))
                                                                                        var v66 bool = v49
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    }
                                                                                    var v67 uint64 = 168
                                                                                    _ = v67
                                                                                }
                                                                                var v68 uint16 = 163
                                                                                v69 := GenS11{F0: 1}
                                                                                _ = v68
                                                                                _ = v69
                                                                            }
                                                                        case 127:
                                                                            fmt.Println(v50, v50)
                                                                        }
                                                                        _ = len(v3)
                                                                        fmt.Println(fi26, v50)
                                                                        var v70 float32 = 1e10
                                                                        if ((!v49) && (9223372036854775807 == 42)) {
                                                                            var v71 bool = (-1 <= 0)
                                                                            v70 = -2.25
                                                                            _ = v71
                                                                        }
                                                                        v1 += 182
                                                                        v1 = v1
                                                                        v72 := GenS12{F0: -1, F1: 127}
                                                                        v73 := GenS3{F0: v49, F1: 3}
                                                                        _ = v70
                                                                        _ = v72
                                                                        _ = v73
                                                                    }
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                            default:
                                                                v74 := GenS4{F0: 972, F1: (!true), F2: ("hello" + "a")}
                                                                _ = v74
                                                            }
                                                        default:
                                                            v1 -= v1
                                                        }
                                                        _ = len(v3)
                                                        _ = len(v2)
                                                    }
                                                }
                                                var v75 string = (("foo bar" + ((("hello" + ("\n\t" + "a")) + (("foo bar" + "foo bar") + "")) + "")) + "foo bar")
                                                if len(v2) > 0 {
                                                    v2[0] = 42
                                                }
                                                _ = v75
                                            } else {
                                                v76 := GenS9{F0: 65535, F1: 1}
                                                switch (fi26 | (fi25 | 127)) {
                                                case -2147483648:
                                                    var v77 int32 = 0
                                                    _ = v77
                                                default:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                _ = v76
                                            }
                                            fmt.Println(v1)
                                            v1 = 255
                                        } else {
                                            fmt.Println(fi25, v10)
                                            delete(v3, 0)
                                            v1 -= 175
                                        }
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = 42
                                    }
                                    if (-1 == 42) {
                                        v1 = v1
                                        if (int32(v1) > int32(fi26)) {
                                            v10 = ((1e10 + float64((0.0 + 0.1))) * ((0.0 * (v10 + 0.0)) / v10))
                                            if ((255 >= 33) || true) {
                                                delete(v3, 0)
                                            }
                                        }
                                    }
                                    v1 -= (v1 + 0)
                                    v10 = v10
                                    if (v1 <= 209) {
                                        var v78 string = "\n\t"
                                        switch (fi25 % (((fi25 >> ((135) & 63))) | 1)) {
                                        case -326:
                                            v1 = 128
                                        case -2147483648:
                                            v1 -= (^(v1))
                                        default:
                                            v78 = "hello"
                                        }
                                        _ = v78
                                    } else {
                                        switch (fi25 / ((fi25) | 1)) {
                                        case 32768:
                                            v10 += v10
                                            if false {
                                                if (((127 <= 0) && (1000 < -32768)) && (!("hello" <= ("\n\t" + "a")))) {
                                                    var v79 int16 = 0
                                                    v80 := v3[42]
                                                    _ = v79
                                                    _ = v80
                                                } else {
                                                    var v81 int = genFn0("", 4294967295, 65535)
                                                    v2 = append(v2, 2147483648)
                                                    if ((42 > 100) && (false && (!(4294967295 < 912)))) {
                                                        var v82 float64 = (v10 + v10)
                                                        var v83 int16 = 783
                                                        _ = len(v3)
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                    _ = v81
                                                }
                                                v10 = (float64(-2.25) - (1e10 / ((v10 / v10) / 1.0)))
                                                v84 := GenS13{F0: "\n\t", F1: -9223372036854775808}
                                                _ = v84
                                            } else {
                                                fmt.Println(fi25, fi26)
                                                var v85 float64 = 0.1
                                                v1 = v1
                                                _ = v85
                                            }
                                            if (1 < 4294967295) {
                                                switch ("foo bar" + "foo bar") {
                                                case "":
                                                    v2 = append(v2, 2147483648)
                                                case "foo bar":
                                                    v1 -= 1
                                                }
                                            } else {
                                                v1 += v1
                                            }
                                        case -1:
                                            delete(v3, 2147483647)
                                        case 127:
                                            if false {
                                                _ = len(v3)
                                                var v86 bool = true
                                                _ = v86
                                            }
                                        default:
                                            var v87 string = "a"
                                            _ = v87
                                        }
                                    }
                                    v10 -= 0.0
                                }
                            case 680:
                                var v88 bool = (880 >= 208)
                                _ = v88
                            }
                        case "":
                            if (!true) {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if ((("a" + (("\n\t" + ("a" + ((("hello" + ((("foo bar" + "foo bar") + "a") + "hello")) + "") + "\n\t"))) + ("" + ""))) != "hello") && true) {
                                    var v89 float32 = (((0.1 * (1e10 - 0.1)) + -2.25) * -2.25)
                                    _ = v89
                                } else {
                                    var v90 float64 = (0.0 * v10)
                                    var v91 float64 = v90
                                    v91 = v90
                                    v1 += 252
                                    if true {
                                        v3[2147483647] = -9223372036854775808
                                    } else {
                                        fmt.Println(fi26)
                                        fmt.Println(fi26, fi26)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        switch ((("foo bar" + "\n\t") + "") + ("hello" + ("foo bar" + ""))) {
                                        case "":
                                            v91 = 1.5
                                        case "hello":
                                            var v92 float32 = 1e10
                                            _ = v92
                                        default:
                                            v10 += (1e10 * -2.25)
                                        }
                                    }
                                    _ = len(v3)
                                    v90 += v91
                                    if false {
                                        var v93 uint8 = v1
                                        var v94 int64 = -1
                                        _ = v93
                                        _ = v94
                                    } else {
                                        v90 = ((v10 / v10) + (v91 / ((v90 / v91) + 0.0)))
                                        v1 -= 128
                                    }
                                    v1 = v1
                                    var v95 int64 = 211
                                    _ = v90
                                    _ = v91
                                    _ = v95
                                }
                                v96 := GenS11{F0: -1}
                                if (!false) {
                                    v97 := GenS1{F0: -128}
                                    _ = v97
                                } else {
                                    if true {
                                        var v98 uint16 = 32767
                                        v1 = (v1 + (v1 ^ v1))
                                        fmt.Println(fi25)
                                        _ = v98
                                    }
                                }
                                _ = v96
                            } else {
                                if true {
                                    switch "a" {
                                    case "\n\t":
                                        var v99 uint8 = v1
                                        _ = v99
                                    case "foo bar":
                                        var v100 int64 = -1
                                        _ = v100
                                    case "a":
                                        v1 -= v1
                                    default:
                                        v101 := v3[656]
                                        _ = v101
                                    }
                                } else {
                                    v1 = (v1 + v1)
                                }
                                v102 := GenS14{F0: -414, F1: 2, F2: -1}
                                if (!(546 > 65535)) {
                                    switch "" {
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "\n\t":
                                            fmt.Println(v10, v1, fi25)
                                        case "foo bar":
                                            switch (fi25 * 0) {
                                            case 1:
                                                var v103 uint64 = 424
                                                switch 42 {
                                                case 1:
                                                    if len(v2) > 0 {
                                                        v2[0] = 492
                                                    }
                                                }
                                                v104 := GenS8{F0: ((("a" + ("a" + "hello")) + "a") <= "")}
                                                _ = v103
                                                _ = v104
                                            case -1:
                                                switch "hello" {
                                                case "foo bar":
                                                    v2 = append(v2, 721)
                                                    switch fi26 {
                                                    case 32768:
                                                        switch (fi25 / ((fi26) | 1)) {
                                                        case 42:
                                                            v105 := GenS11{F0: -114}
                                                            _ = v105
                                                        case -1:
                                                            var v106 int8 = -128
                                                            _ = v106
                                                        default:
                                                            var v107 int32 = -35
                                                            _ = v107
                                                        }
                                                    case 127:
                                                        v1 = 127
                                                    case 153:
                                                        fmt.Println(v10)
                                                    }
                                                    v3[104] = fi26
                                                    v1 -= v1
                                                    var v108 uint = 1
                                                    _ = v108
                                                case "\n\t":
                                                    switch genFn0("", 892, 2) {
                                                    case -9223372036854775808:
                                                        switch fi26 {
                                                        case 1:
                                                            v1 -= (v1 << ((255) & 63))
                                                        case 4294967295:
                                                            if false {
                                                                var v109 uint = 42
                                                                if (fi26 >= fi25) {
                                                                    v110 := GenS4{F0: 1, F1: false, F2: "a"}
                                                                    var v111 bool = (3.14159 != 0.1)
                                                                    var v112 uint8 = (v1 + v1)
                                                                    v1 -= (v1 >> ((v109) & 63))
                                                                    _ = v110
                                                                    _ = v111
                                                                    _ = v112
                                                                }
                                                                _ = v109
                                                            } else {
                                                                v10 = v10
                                                                fmt.Println(fi26, fi25)
                                                            }
                                                        case 42:
                                                            var v113 int = (fi26 / ((fi25) | 1))
                                                            var v114 string = "a"
                                                            fmt.Println(fi25)
                                                            v3[65535] = 4294967295
                                                            var v115 float32 = 1e10
                                                            _ = v113
                                                            _ = v114
                                                            _ = v115
                                                        default:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    case -2147483648:
                                                        v2 = append(v2, 9)
                                                    case 128:
                                                        v10 = 0.1
                                                    }
                                                default:
                                                    delete(v3, 42)
                                                }
                                            case 281:
                                                v116 := GenS13{F0: "\n\t", F1: 42}
                                                _ = v116
                                            default:
                                                if true {
                                                    v3[916] = genFn0((("a" + "\n\t") + "hello"), 967, 968)
                                                }
                                            }
                                        default:
                                            if (1 == 42) {
                                                fmt.Println(fi26)
                                                var v117 int16 = 514
                                                _ = v117
                                            } else {
                                                switch ("hello" + (("a" + "foo bar") + "")) {
                                                case "foo bar":
                                                    _ = len(v3)
                                                case "hello":
                                                    if false {
                                                        v10 = v10
                                                    } else {
                                                        var v118 string = ""
                                                        _ = v118
                                                    }
                                                case "a":
                                                    fmt.Println(v10, v10)
                                                default:
                                                    v119 := GenS14{F0: -1, F1: 65535, F2: fi26}
                                                    _ = v119
                                                }
                                                var v120 int16 = 1
                                                v120 += v120
                                                _ = v120
                                            }
                                            if (false || ((v1 >> ((255) & 63)) < v1)) {
                                                v121 := GenS15{F0: 569}
                                                var v122 uint16 = 48
                                                _ = v121
                                                _ = v122
                                            } else {
                                                v1 = 255
                                                delete(v3, 1)
                                            }
                                            v10 = -2.25
                                        }
                                    case "foo bar":
                                        var v123 int8 = 1
                                        _ = v123
                                    }
                                }
                                _ = v102
                            }
                        case "a":
                            v124 := v3[65535]
                            _ = v124
                        default:
                            v10 = float64(0.1)
                        }
                        v10 = float64(0.0)
                        fmt.Println(fi25, fi26)
                        _ = fi26
                    }
                case -1:
                    fmt.Println(v10, v1, v1)
                }
                var v125 int32 = 472
                _ = fi25
                _ = v125
            }
        case "hello":
            v126 := GenS8{F0: (42 != -1)}
            _ = v126
        default:
            var v127 int32 = 2147483647
            _ = v127
        }
        if false {
            if (992 < -1) {
                _ = len(v2)
                v128 := v3[70]
                _ = v128
            } else {
                v1 = v1
                for fi129 := 0; fi129 < 4; fi129++ {
                    var v130 uint32 = 2147483647
                    _ = fi129
                    _ = v130
                }
            }
            v131 := GenS16{F0: -1}
            v132 := GenS2{F0: 368, F1: 172}
            _ = v131
            _ = v132
        } else {
            if true {
                var v133 uint32 = 779
                switch 659 {
                case 42:
                    switch "foo bar" {
                    case "\n\t":
                        if (!true) {
                            switch ("a" + "a") {
                            case "\n\t":
                                v133 += (^(v133))
                            case "a":
                                var v134 int32 = 290
                                _ = v134
                            default:
                                v133 -= (v133 ^ (v133 % ((uint32(v1)) | 1)))
                            }
                        } else {
                            v1 -= 127
                            v135 := GenS17{F0: 0, F1: 42}
                            _ = v135
                        }
                    case "a":
                        var v136 uint8 = 8
                        _ = v136
                    case "hello":
                        v137 := GenS7{F0: v133, F1: -203}
                        _ = v137
                    default:
                        v1 = v1
                    }
                case 1:
                    var v138 int64 = -9223372036854775808
                    _ = v138
                case -2147483648:
                    var v139 int16 = 32767
                    _ = v139
                }
                _ = v133
            }
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
            switch 965 {
            case -626:
                if (!(0 <= 0)) {
                    v3[630] = 188
                }
            case -673:
                var v140 uint = 265
                _ = v140
            case 127:
                v10 += (((3.14159 * (v10 + (v10 / 1e10))) * 1e10) + v10)
            }
        }
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v141 bool = ("hello" > "\n\t")
        _ = v141
    }
    v142 := GenS15{F0: 42}
    v143 := GenS4{F0: 748, F1: (1 < 100), F2: ("" + "foo bar")}
    fmt.Println(v10, v1)
    if true {
        delete(v3, 1)
    }
    v144 := GenS18{F0: 4294967295}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v145 float32 = 0.0
    for fi146 := 0; fi146 < 18; fi146++ {
        v10 = -2.25
        v10 += (((-2.25 - v10) + ((v10 - 0.1) + -2.25)) - -2.25)
        _ = fi146
    }
    _ = len(v3)
    switch -147 {
    case 42:
        for fi147 := 0; fi147 < 14; fi147++ {
            v148 := GenS19{F0: 65535}
            if true {
                switch "\n\t" {
                case "":
                    var v149 int = genFn0(("\n\t" + "a"), 2147483647, 2147483648)
                    _ = v149
                case "a":
                    if (-1 <= -63) {
                        for fi150 := 0; fi150 < 5; fi150++ {
                            v10 -= 0.1
                            _ = fi150
                        }
                    } else {
                        if (42 == fi147) {
                            if false {
                                v10 -= v10
                                fmt.Println(v10, v1, fi147)
                                v151 := GenS7{F0: 2, F1: 42}
                                _ = v151
                            } else {
                                if true {
                                    v152 := GenS20{F0: 255}
                                    var v153 float32 = -2.25
                                    _ = v152
                                    _ = v153
                                } else {
                                    fmt.Println(fi147)
                                    v10 -= v10
                                    for fi154 := 0; fi154 < 13; fi154++ {
                                        if len(v2) > 0 {
                                            v2[0] = 42
                                        }
                                        v1 += 128
                                        switch fi147 {
                                        case -9223372036854775808:
                                            v155 := GenS17{F0: 83, F1: 813}
                                            _ = v155
                                        default:
                                            switch (fi147 - fi147) {
                                            case 2147483647:
                                                fmt.Println(fi154, fi147)
                                            case 512:
                                                fmt.Println(fi154, v10, v145)
                                            default:
                                                var v156 uint16 = 2
                                                _ = v156
                                            }
                                        }
                                        _ = fi154
                                    }
                                }
                            }
                            fmt.Println(v145, v1)
                        }
                        fmt.Println(v1, v145)
                        _ = len(v2)
                    }
                case "\n\t":
                    v1 = (v1 & (v1 * (^(v1))))
                default:
                    fmt.Println(v10, v145, v10)
                }
                var v157 float64 = v10
                _ = v157
            }
            switch (("" + "\n\t") + "") {
            case "\n\t":
                if true {
                    var v158 int8 = 100
                    v145 += ((1e10 * -2.25) + -2.25)
                    if (0.0 != 0.1) {
                        v145 += (float32(3.14159) / ((v145 / v145) / (((-2.25 / float32(v10)) - float32(v10)) - v145)))
                        var v159 float32 = v145
                        var v160 uint = 0
                        _ = v159
                        _ = v160
                    }
                    _ = v158
                }
            default:
                v10 = 3.14159
            }
            _ = fi147
            _ = v148
        }
    case 4294967295:
        v1 -= v1
    default:
        var v161 int16 = 0
        _ = v161
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v145:", v145)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v142
    _ = v143
    _ = v144
    _ = v145
    _ = v2
    _ = v3
}
