package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 bool; F1 string; F2 bool }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8; F1 uint16 }
type GenS9 struct { F0 uint64; F1 bool }
type GenS10 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS11 struct { F0 int; F1 uint16 }
type GenS12 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2147483647
    var v3 int16 = -640
    var v4 uint16 = 1
    var v5 float32 = 0.0
    v6 := GenS1{F0: 42, F1: v3}
    v5 = (((-2.25 + v5) * 0.0) + float32(1.5))
    delete(v1, (v2 * 708))
    switch -128 {
    case -247:
        v7 := GenS1{F0: 9223372036854775807, F1: v3}
        _ = v7
    case -460:
        fmt.Println(v2, v5)
    case 4294967295:
        v5 += 0.1
    }
    v4 -= v4
    var v8 uint64 = 209
    var v9 int32 = -2147483648
    if true {
        if false {
            v10 := []int{-1, -1}
            for fi11 := 0; fi11 < 8; fi11++ {
                _ = len(v10)
                switch "\n\t" {
                case "a":
                    for fi12 := 0; fi12 < 17; fi12++ {
                        switch (((("" + "\n\t") + ("a" + "\n\t")) + "\n\t") + "hello") {
                        case "hello":
                            var v13 float32 = v5
                            _ = v13
                        case "":
                            v14 := GenS1{F0: int64(v9), F1: -1}
                            _ = v14
                        default:
                            var v15 int64 = 858
                            _ = v15
                        }
                        v4 = v4
                        _ = fi12
                    }
                case "\n\t":
                    switch fi11 {
                    case 9223372036854775807:
                        fmt.Println(v8)
                    case -347:
                        v16 := GenS1{F0: -802, F1: (v3 % (((-(v3))) | 1))}
                        _ = v16
                    default:
                        v3 = v3
                    }
                case "":
                    var v17 int32 = int32(v3)
                    _ = v17
                }
                _ = fi11
            }
            _ = v10
        }
    }
    switch ("\n\t" + "hello") {
    case "hello":
        v18 := GenS1{F0: -9223372036854775808, F1: v3}
        _ = v18
    case "a":
        v19 := GenS2{F0: false, F1: "a", F2: (32767 <= v3)}
        _ = v19
    case "foo bar":
        var v20 int32 = -2147483648
        _ = v20
    default:
        var v21 int64 = 0
        _ = v21
    }
    var v22 int = -874
    var v23 float32 = v5
    if (!(100 != 1)) {
        if (!((^(v2)) < (^(v2)))) {
            if (-2.25 != (float64(float32(-2.25)) / 1.5)) {
                switch "" {
                case "\n\t":
                    if ((-2.25 == float32(0.0)) && (false || (false || (!true)))) {
                        for fi24 := 0; fi24 < 8; fi24++ {
                            v2 = (v2 % ((500) | 1))
                            _ = fi24
                        }
                        v2 = 4294967295
                        if true {
                            if (true || ((v8 << ((42) & 63)) >= 4294967295)) {
                                if (-128 != int8(v8)) {
                                    fmt.Println(v23, v2)
                                }
                                v25 := GenS3{F0: (^(v9)), F1: 733}
                                fmt.Println(v2, v5)
                                _ = v25
                            }
                            var v26 bool = (v9 >= v9)
                            if (0.1 > 1.5) {
                                v27 := v1[v2]
                                if true {
                                    v23 = (((v5 / v23) + v23) / (((v23 * v23) / 0.1) * -2.25))
                                    for fi28 := 0; fi28 < 8; fi28++ {
                                        v29 := []int{(fi28 | (-(fi28)))}
                                        switch "" {
                                        case "foo bar":
                                            v30 := v1[v2]
                                            _ = v30
                                        case "":
                                            switch "a" {
                                            case "foo bar":
                                                v23 = v5
                                            case "hello":
                                                if true {
                                                    v31 := GenS4{F0: 0, F1: 152, F2: v4}
                                                    if v27 {
                                                        fmt.Println(fi28, v2, v4)
                                                        v32 := GenS4{F0: int16(v22), F1: 195, F2: v4}
                                                        for fi33 := 0; fi33 < 13; fi33++ {
                                                            if (!(!false)) {
                                                                switch (fi33 & (fi33 * fi33)) {
                                                                case -290:
                                                                    v34 := GenS4{F0: -495, F1: 202, F2: (v4 & (v4 ^ v4))}
                                                                    v26 = false
                                                                    _ = v34
                                                                default:
                                                                    if false {
                                                                        switch (v22 - (fi33 | fi28)) {
                                                                        case 2147483647:
                                                                            if (42 != (v3 & int16(v2))) {
                                                                                var v35 float64 = (-2.25 / (-2.25 + (3.14159 - (float64((0.0 + (float32(((float64(1e10) * (1e10 / 1e10)) * ((1e10 + (((3.14159 - 3.14159) * 1.5) + 1e10)) + -2.25))) + v5))) * (3.14159 - 3.14159)))))
                                                                                v1[(v2 - v2)] = v27
                                                                                v36 := GenS5{F0: 100}
                                                                                v37 := GenS3{F0: -1, F1: 2}
                                                                                switch fi33 {
                                                                                case 900:
                                                                                    v3 = (v3 | 1000)
                                                                                case -895:
                                                                                    if (42 > 42) {
                                                                                        var v38 uint64 = (v8 ^ v8)
                                                                                        _ = v38
                                                                                    } else {
                                                                                        var v39 string = ("foo bar" + "foo bar")
                                                                                        var v40 int = fi33
                                                                                        delete(v1, v2)
                                                                                        _ = v39
                                                                                        _ = v40
                                                                                    }
                                                                                case 127:
                                                                                    if true {
                                                                                        v41 := GenS3{F0: 42, F1: 4294967295}
                                                                                        _ = v41
                                                                                    }
                                                                                default:
                                                                                    var v42 int8 = 1
                                                                                    _ = v42
                                                                                }
                                                                                _ = v35
                                                                                _ = v36
                                                                                _ = v37
                                                                            } else {
                                                                                v27 = false
                                                                                v8 -= (v8 / ((v8) | 1))
                                                                            }
                                                                        case 32767:
                                                                            if true {
                                                                                v43 := GenS6{F0: int64(v2), F1: 127}
                                                                                _ = v43
                                                                            } else {
                                                                                var v44 string = (("" + ("foo bar" + "hello")) + (("foo bar" + "hello") + ("hello" + (("" + (("" + ("" + "")) + ("a" + (("foo bar" + "") + ((((("hello" + (("hello" + "\n\t") + (("a" + "\n\t") + ("\n\t" + "a")))) + "a") + "a") + "") + "foo bar"))))) + ("foo bar" + "\n\t")))))
                                                                                _ = v44
                                                                            }
                                                                        default:
                                                                            var v45 uint = 4294967295
                                                                            _ = v45
                                                                        }
                                                                        fmt.Println(v5, v3, fi28)
                                                                    } else {
                                                                        _ = len(v1)
                                                                        if v26 {
                                                                            v8 += v8
                                                                            v29 = append(v29, (fi33 / ((-9223372036854775808) | 1)))
                                                                            if true {
                                                                                _ = len(v1)
                                                                            }
                                                                        } else {
                                                                            if ((true || true) || (-429 <= -9223372036854775808)) {
                                                                                if true {
                                                                                    fmt.Println(v26, v27)
                                                                                    v23 = v5
                                                                                    switch "a" {
                                                                                    case "hello":
                                                                                        switch (v22 % ((int(v2)) | 1)) {
                                                                                        case 42:
                                                                                            if v27 {
                                                                                                fmt.Println(v5)
                                                                                                if ((-(v9)) <= (v9 / ((v9) | 1))) {
                                                                                                    if len(v29) >= 2 {
                                                                                                        v29 = v29[1:len(v29)]
                                                                                                    }
                                                                                                    fmt.Println(fi28)
                                                                                                } else {
                                                                                                    v46 := v1[v2]
                                                                                                    _ = v46
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            v47 := GenS7{F0: fi33}
                                                                                            _ = v47
                                                                                        }
                                                                                    case "\n\t":
                                                                                        var v48 int16 = v3
                                                                                        _ = v48
                                                                                    case "a":
                                                                                        var v49 uint = 1
                                                                                        _ = v49
                                                                                    default:
                                                                                        v1[(v2 - 373)] = (int64(v22) != 900)
                                                                                    }
                                                                                } else {
                                                                                    switch (v22 ^ v22) {
                                                                                    case 1:
                                                                                        fmt.Println(fi28)
                                                                                    case 9223372036854775807:
                                                                                        var v50 float32 = (float32(float64(1e10)) * v5)
                                                                                        _ = v50
                                                                                    default:
                                                                                        var v51 int = (fi33 - fi33)
                                                                                        _ = v51
                                                                                    }
                                                                                }
                                                                                var v52 float32 = v23
                                                                                _ = v52
                                                                            } else {
                                                                                switch ("hello" + ("a" + (("a" + "a") + "foo bar"))) {
                                                                                case "hello":
                                                                                    v5 += v23
                                                                                case "foo bar":
                                                                                    v53 := GenS6{F0: 0, F1: (fi33 / (((v22 * fi33)) | 1))}
                                                                                    _ = v53
                                                                                case "":
                                                                                    var v54 int16 = v3
                                                                                    _ = v54
                                                                                }
                                                                                v26 = (!true)
                                                                                var v55 bool = true
                                                                                _ = v55
                                                                            }
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                var v56 int32 = v9
                                                                                _ = v56
                                                                            case "":
                                                                                _ = len(v29)
                                                                                v1[2147483647] = (-128 >= 99)
                                                                            case "a":
                                                                                var v57 uint32 = v2
                                                                                _ = v57
                                                                            default:
                                                                                if v26 {
                                                                                    v1[v2] = (v27 || (v26 || v27))
                                                                                } else {
                                                                                    if (523 != (v4 | 530)) {
                                                                                        v9 = (v9 ^ v9)
                                                                                        switch ("\n\t" + ((("foo bar" + "") + "\n\t") + "a")) {
                                                                                        case "\n\t":
                                                                                            fmt.Println(v27, v27, v23)
                                                                                        case "foo bar":
                                                                                            if true {
                                                                                                var v58 uint32 = 1
                                                                                                v4 = v4
                                                                                                _ = v58
                                                                                            } else {
                                                                                                if (-9223372036854775808 < -1) {
                                                                                                    var v59 uint = 42
                                                                                                    _ = v59
                                                                                                } else {
                                                                                                    if true {
                                                                                                        var v60 uint32 = uint32(v4)
                                                                                                        var v61 uint8 = 48
                                                                                                        if v27 {
                                                                                                            v61 = 42
                                                                                                            _ = len(v1)
                                                                                                        } else {
                                                                                                            v2 += 65535
                                                                                                            v61 = v61
                                                                                                            fmt.Println(v4, v61)
                                                                                                            var v62 uint8 = v61
                                                                                                            var v63 uint32 = (v2 >> ((4294967295) & 63))
                                                                                                            var v64 float32 = v5
                                                                                                            v64 = v64
                                                                                                            v65 := v1[v2]
                                                                                                            fmt.Println(v62, v60)
                                                                                                            if true {
                                                                                                                v66 := GenS8{F0: -73, F1: v4}
                                                                                                                v1[(v60 - 2147483647)] = ((v9 << ((65535) & 63)) <= -2147483648)
                                                                                                                _ = v66
                                                                                                            }
                                                                                                            v63 = v2
                                                                                                            fmt.Println(fi28)
                                                                                                            fmt.Println(v2, v3, v61)
                                                                                                            v67 := GenS5{F0: -128}
                                                                                                            fmt.Println(fi28, fi28)
                                                                                                            var v68 int64 = -9223372036854775808
                                                                                                            _ = v62
                                                                                                            _ = v63
                                                                                                            _ = v64
                                                                                                            _ = v65
                                                                                                            _ = v67
                                                                                                            _ = v68
                                                                                                        }
                                                                                                        _ = v60
                                                                                                        _ = v61
                                                                                                    }
                                                                                                }
                                                                                                if v26 {
                                                                                                    _ = len(v29)
                                                                                                    var v69 uint32 = (v2 * 2147483647)
                                                                                                    switch fi33 {
                                                                                                    case 0:
                                                                                                        var v70 int32 = (v9 % (((v9 + (v9 + v9))) | 1))
                                                                                                        _ = v70
                                                                                                    case 32768:
                                                                                                        _ = len(v29)
                                                                                                    default:
                                                                                                        fmt.Println(v5, v27)
                                                                                                    }
                                                                                                    _ = v69
                                                                                                }
                                                                                                if (100 > 42) {
                                                                                                    v26 = (-9223372036854775808 <= -771)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        if v27 {
                                                                                            v4 -= v4
                                                                                        } else {
                                                                                            switch ("foo bar" + "a") {
                                                                                            case "":
                                                                                                v71 := GenS4{F0: v3, F1: 0, F2: 1}
                                                                                                _ = v71
                                                                                            case "foo bar":
                                                                                                v22 = (fi33 / ((fi28) | 1))
                                                                                            }
                                                                                        }
                                                                                        var v72 uint8 = 42
                                                                                        v23 = (v23 / v5)
                                                                                        _ = v72
                                                                                    }
                                                                                    if (uint(v2) > 65535) {
                                                                                        delete(v1, 95)
                                                                                        switch fi28 {
                                                                                        case 2147483647:
                                                                                            v73 := GenS8{F0: -106, F1: (v4 << ((42) & 63))}
                                                                                            _ = v73
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                            v74 := GenS4{F0: v3, F1: 1, F2: v4}
                                                                            _ = v74
                                                                        }
                                                                    }
                                                                }
                                                                v1[v2] = false
                                                                if (-57 <= int8(v2)) {
                                                                    var v75 uint64 = v8
                                                                    if (v26 || false) {
                                                                        var v76 bool = true
                                                                        if v76 {
                                                                            if false {
                                                                                var v77 uint8 = 255
                                                                                var v78 int64 = 1
                                                                                if (int64(v3) < v78) {
                                                                                    switch ("foo bar" + ("hello" + "")) {
                                                                                    case "hello":
                                                                                        if (v4 > 255) {
                                                                                            var v79 string = ("\n\t" + ("a" + "\n\t"))
                                                                                            _ = v79
                                                                                        } else {
                                                                                            v80 := GenS2{F0: (65535 <= 88), F1: ("" + ("\n\t" + "\n\t")), F2: (false || (false || v27))}
                                                                                            var v81 int = v22
                                                                                            _ = v80
                                                                                            _ = v81
                                                                                        }
                                                                                    case "foo bar":
                                                                                        var v82 int8 = -99
                                                                                        _ = v82
                                                                                    default:
                                                                                        var v83 uint32 = v2
                                                                                        v84 := GenS1{F0: 0, F1: -387}
                                                                                        _ = v83
                                                                                        _ = v84
                                                                                    }
                                                                                    var v85 uint64 = (v75 | v75)
                                                                                    switch "a" {
                                                                                    case "":
                                                                                        v86 := GenS1{F0: v78, F1: v3}
                                                                                        _ = v86
                                                                                    case "a":
                                                                                        v2 = v2
                                                                                    case "foo bar":
                                                                                        var v87 int16 = 1
                                                                                        _ = v87
                                                                                    default:
                                                                                        fmt.Println(v26)
                                                                                    }
                                                                                    v5 += 1e10
                                                                                    var v88 uint64 = (^(v8))
                                                                                    _ = v85
                                                                                    _ = v88
                                                                                } else {
                                                                                    v89 := GenS4{F0: 32767, F1: uint8(v8), F2: (v4 & v4)}
                                                                                    var v90 uint8 = v77
                                                                                    v29 = append(v29, fi33)
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                }
                                                                                v4 += v4
                                                                                _ = v77
                                                                                _ = v78
                                                                            } else {
                                                                                v29 = append(v29, 32768)
                                                                                if (fi33 > (fi33 * int(v4))) {
                                                                                    v1[v2] = (49 == 100)
                                                                                } else {
                                                                                    fmt.Println(v9, v5)
                                                                                }
                                                                            }
                                                                            v4 = 2
                                                                            switch "" {
                                                                            case "a":
                                                                                var v91 float64 = (float64(float32(1.5)) - (0.1 - 1e10))
                                                                                _ = v91
                                                                            }
                                                                            var v92 int8 = 0
                                                                            _ = v92
                                                                        }
                                                                        var v93 string = ""
                                                                        var v94 float64 = 1e10
                                                                        _ = v76
                                                                        _ = v93
                                                                        _ = v94
                                                                    } else {
                                                                        v2 -= (v2 | (v2 % ((v2) | 1)))
                                                                    }
                                                                    _ = v75
                                                                }
                                                                switch "hello" {
                                                                case "hello":
                                                                    v95 := GenS9{F0: 18446744073709551615, F1: (false || ((!v26) && v26))}
                                                                    _ = v95
                                                                default:
                                                                    v8 = v8
                                                                }
                                                                v96 := GenS7{F0: 1}
                                                                _ = v96
                                                            }
                                                            switch "hello" {
                                                            case "a":
                                                                var v97 int8 = 77
                                                                _ = v97
                                                            case "":
                                                                var v98 string = ""
                                                                _ = v98
                                                            case "\n\t":
                                                                var v99 uint16 = v4
                                                                _ = v99
                                                            }
                                                            _ = fi33
                                                        }
                                                        _ = v32
                                                    }
                                                    _ = v31
                                                } else {
                                                    v29 = append(v29, 812)
                                                    switch v22 {
                                                    case -9223372036854775808:
                                                        v100 := GenS2{F0: true, F1: "", F2: false}
                                                        _ = v100
                                                    default:
                                                        for fi101 := 0; fi101 < 12; fi101++ {
                                                            v102 := v1[(v2 << ((0) & 63))]
                                                            v103 := GenS10{F0: 42, F1: 65535, F2: 127}
                                                            _ = fi101
                                                            _ = v102
                                                            _ = v103
                                                        }
                                                    }
                                                }
                                            default:
                                                for fi104 := 0; fi104 < 2; fi104++ {
                                                    if true {
                                                        var v105 float32 = 0.1
                                                        v23 = (((float32(0.0) - v105) / ((v23 / 1e10) / v5)) * ((((v105 / (v23 + 0.0)) - v5) + v23) + v23))
                                                        _ = v105
                                                    } else {
                                                        v106 := GenS11{F0: (fi104 - 1), F1: 65535}
                                                        var v107 float64 = -2.25
                                                        _ = v106
                                                        _ = v107
                                                    }
                                                    if len(v29) > 0 {
                                                        v29[0] = fi104
                                                    }
                                                    _ = fi104
                                                }
                                            }
                                        }
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                                _ = v27
                            }
                            _ = v26
                        }
                    }
                case "a":
                    var v108 int16 = 0
                    _ = v108
                }
                v8 += v8
            } else {
                var v109 uint32 = (v2 >> ((0) & 63))
                _ = v109
            }
            var v110 string = "foo bar"
            fmt.Println(v110, v23, v2)
            _ = v110
        } else {
            var v111 int32 = 540
            fmt.Println(v4, v23)
            _ = v111
        }
        v3 = int16(v9)
    } else {
        var v112 float32 = ((v23 * v5) - (v5 - (v5 - (v23 * float32((3.14159 + -2.25))))))
        v113 := v1[(v2 | (v2 << ((0) & 63)))]
        switch 127 {
        case 1:
            v3 = v3
        case 42:
            var v114 float64 = 3.14159
            _ = v114
        default:
            v115 := GenS12{F0: v3, F1: 0}
            _ = v115
        }
        _ = v112
        _ = v113
    }
    var v116 uint32 = (v2 >> ((65535) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v116
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
