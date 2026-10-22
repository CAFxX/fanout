package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS6 struct { F0 int8; F1 uint; F2 uint32 }
type GenS7 struct { F0 int8; F1 int64; F2 uint }
type GenS8 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 int64 }
func genFn0() int16 {
    v5 := []int{}
    v6 := GenS1{F0: 42, F1: (!(1.5 < (-2.25 + float64(1e10))))}
    var v7 uint = 65535
    _ = v5
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int32 = -2147483648
    var v2 bool = false
    v3 := []int{32768}
    v4 := []string{"hello", "a", (("foo bar" + (("hello" + "\n\t") + "\n\t")) + "hello"), "a"}
    fmt.Println(v2, v1)
    for fi8 := 0; fi8 < 6; fi8++ {
        v9 := GenS2{F0: (fi8 | fi8), F1: 127}
        var v10 int8 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 = append(v3, -2147483648)
    switch 127 {
    case 0:
        v11 := GenS1{F0: 32767, F1: ((!true) && v2)}
        _ = v11
    }
    v2 = (!(65535 > 2147483648))
    for fi12 := 0; fi12 < 5; fi12++ {
        switch (fi12 & fi12) {
        case 32768:
            var v13 bool = v2
            _ = v13
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v14 float32 = (-2.25 / (float32(1e10) - -2.25))
        _ = fi12
        _ = v14
    }
    switch 9223372036854775807 {
    case 0:
        if true {
            for fi15 := 0; fi15 < 1; fi15++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if ((!v2) || (718 <= 18446744073709551615)) {
                    if (v2 || v2) {
                        if (65535 >= 581) {
                            v2 = (true && v2)
                        } else {
                            for fi16 := 0; fi16 < 2; fi16++ {
                                var v17 int8 = 54
                                var v18 int16 = 32767
                                var v19 int = (fi16 + fi16)
                                fmt.Println(v17, fi15)
                                if v2 {
                                    v20 := GenS1{F0: 4294967295, F1: (0.1 > ((0.1 - 0.0) * 0.0))}
                                    v21 := GenS2{F0: fi15, F1: v17}
                                    fmt.Println(fi15, fi15, v17)
                                    _ = v20
                                    _ = v21
                                } else {
                                    switch (fi15 + (-(fi16))) {
                                    case 414:
                                        var v22 float64 = (-2.25 - 3.14159)
                                        _ = v22
                                    default:
                                        switch 32768 {
                                        case 9223372036854775807:
                                            fmt.Println(v18, v1, v1)
                                        default:
                                            v23 := GenS3{F0: (v1 & v1), F1: -998}
                                            _ = v23
                                        }
                                    }
                                    v1 += (v1 - -2147483648)
                                }
                                v24 := GenS4{F0: -420}
                                v25 := GenS2{F0: (fi15 ^ (fi15 ^ fi16)), F1: -112}
                                _ = fi16
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v24
                                _ = v25
                            }
                            fmt.Println(fi15)
                        }
                        switch (fi15 | 471) {
                        case 32767:
                            var v26 float64 = float64(0.0)
                            _ = v26
                        case 128:
                            v1 = v1
                        case -2147483648:
                            switch (^(fi15)) {
                            case 32768:
                                fmt.Println(fi15)
                            case 4294967295:
                                if len(v4) > 0 {
                                    v4[0] = ("\n\t" + "")
                                }
                            case -116:
                                v4 = append(v4, "\n\t")
                            default:
                                v1 += 338
                            }
                        default:
                            v3 = append(v3, fi15)
                        }
                        v1 = (v1 - int32(fi15))
                    } else {
                        v1 = (v1 / ((v1) | 1))
                        for fi27 := 0; fi27 < 8; fi27++ {
                            var v28 int64 = -1
                            fmt.Println(v1)
                            _ = fi27
                            _ = v28
                        }
                        if false {
                            if v2 {
                                v3 = append(v3, (fi15 * (fi15 / (((fi15 ^ (fi15 | (-(fi15))))) | 1))))
                                if (true || false) {
                                    if (-2.25 != 0.1) {
                                        if true {
                                            var v29 int16 = 32767
                                            if (42 <= int64(v1)) {
                                                fmt.Println(v1, v29)
                                            }
                                            if (1 <= 2) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v30 uint32 = 2147483647
                                                _ = v30
                                            } else {
                                                var v31 string = "\n\t"
                                                var v32 int64 = -9223372036854775808
                                                fmt.Println(v32)
                                                _ = v31
                                                _ = v32
                                            }
                                            _ = v29
                                        }
                                        switch "foo bar" {
                                        case "a":
                                            v1 -= 1
                                        case "":
                                            var v33 float32 = 1.5
                                            _ = v33
                                        default:
                                            for fi34 := 0; fi34 < 19; fi34++ {
                                                v1 += (v1 / ((int32(fi15)) | 1))
                                                if v2 {
                                                    var v35 bool = false
                                                    var v36 float64 = (0.1 - float64(1e10))
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    var v37 bool = (!(false && true))
                                                    _ = v37
                                                }
                                                if (!(-9223372036854775808 != 9223372036854775807)) {
                                                    v2 = (1 != 2)
                                                    fmt.Println(v2, v2, fi34)
                                                    if true {
                                                        v2 = (100 != -128)
                                                        switch fi15 {
                                                        case -2147483648:
                                                            v2 = false
                                                        default:
                                                            if len(v4) > 0 {
                                                                v4[0] = "hello"
                                                            }
                                                        }
                                                        v1 = (v1 - v1)
                                                        switch fi15 {
                                                        case 4294967295:
                                                            if ((42 > 1) || (!((v2 || v2) || ((fi34 >> ((42) & 63)) != 4294967295)))) {
                                                                if (((255 > 466) || (float32(0.1) <= 1e10)) || true) {
                                                                    if false {
                                                                        var v38 int32 = v1
                                                                        _ = v38
                                                                    } else {
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            switch ("foo bar" + "hello") {
                                                                            case "a":
                                                                                var v39 float64 = 0.1
                                                                                if ((v1 | v1) < v1) {
                                                                                    if false {
                                                                                        var v40 int16 = genFn0()
                                                                                        _ = v40
                                                                                    } else {
                                                                                        v2 = (42 < 42)
                                                                                    }
                                                                                    v41 := GenS5{F0: 83, F1: -108, F2: 2}
                                                                                    v42 := GenS2{F0: fi15, F1: 42}
                                                                                    _ = v41
                                                                                    _ = v42
                                                                                }
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = fi15
                                                                                }
                                                                                _ = v39
                                                                            case "foo bar":
                                                                                var v43 uint16 = 0
                                                                                _ = v43
                                                                            case "":
                                                                                switch 30 {
                                                                                case 28:
                                                                                    if (v2 && (!false)) {
                                                                                        v44 := GenS1{F0: (fi34 * fi34), F1: v2}
                                                                                        _ = v44
                                                                                    } else {
                                                                                        if (!v2) {
                                                                                            var v45 int16 = -32768
                                                                                            var v46 float64 = (float64((1.5 * 0.0)) / 1e10)
                                                                                            if true {
                                                                                                v47 := GenS4{F0: -421}
                                                                                                if (v2 && false) {
                                                                                                    v45 = -899
                                                                                                    v48 := GenS3{F0: (v1 % ((v1) | 1)), F1: 424}
                                                                                                    _ = v48
                                                                                                } else {
                                                                                                    var v49 int8 = 100
                                                                                                    _ = v49
                                                                                                }
                                                                                                _ = len(v4)
                                                                                                var v50 int16 = genFn0()
                                                                                                _ = v47
                                                                                                _ = v50
                                                                                            }
                                                                                            _ = v45
                                                                                            _ = v46
                                                                                        } else {
                                                                                            if v2 {
                                                                                                var v51 int16 = genFn0()
                                                                                                _ = v51
                                                                                            } else {
                                                                                                fmt.Println(v2, fi15, v1)
                                                                                                if v2 {
                                                                                                    var v52 int8 = -128
                                                                                                    _ = v52
                                                                                                } else {
                                                                                                    v1 -= (v1 % (((v1 >> ((2) & 63))) | 1))
                                                                                                    v1 += -1
                                                                                                }
                                                                                                var v53 uint8 = 42
                                                                                                _ = v53
                                                                                            }
                                                                                            v4 = append(v4, "a")
                                                                                            v54 := GenS4{F0: 11}
                                                                                            var v55 float32 = -2.25
                                                                                            _ = v54
                                                                                            _ = v55
                                                                                        }
                                                                                        if ((false || (42 < 319)) || false) {
                                                                                            v2 = (fi15 == (fi15 % ((-1) | 1)))
                                                                                            _ = len(v4)
                                                                                            var v56 uint16 = 65535
                                                                                            _ = v56
                                                                                        } else {
                                                                                            var v57 int32 = 1
                                                                                            var v58 uint32 = 113
                                                                                            _ = v57
                                                                                            _ = v58
                                                                                        }
                                                                                        v59 := GenS5{F0: 0, F1: -1, F2: 42}
                                                                                        _ = v59
                                                                                    }
                                                                                default:
                                                                                    v1 -= v1
                                                                                }
                                                                            default:
                                                                                var v60 uint32 = 0
                                                                                _ = v60
                                                                            }
                                                                        case "":
                                                                            switch ("a" + ("foo bar" + "hello")) {
                                                                            case "foo bar":
                                                                                switch ("" + ((("hello" + "a") + "a") + ("\n\t" + (("" + "") + ("hello" + ((("foo bar" + "") + "hello") + "foo bar")))))) {
                                                                                case "":
                                                                                    _ = len(v4)
                                                                                case "foo bar":
                                                                                    switch (("" + "\n\t") + ("foo bar" + "\n\t")) {
                                                                                    case "\n\t":
                                                                                        var v61 int32 = v1
                                                                                        _ = v61
                                                                                    case "foo bar":
                                                                                        if (!(!(1 <= 1))) {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            var v62 int16 = -1
                                                                                            if v2 {
                                                                                                if v2 {
                                                                                                    var v63 int = -950
                                                                                                    v64 := GenS6{F0: -128, F1: 907, F2: 2147483648}
                                                                                                    v1 -= -123
                                                                                                    _ = v63
                                                                                                    _ = v64
                                                                                                }
                                                                                            } else {
                                                                                                v65 := GenS3{F0: v1, F1: int64(fi34)}
                                                                                                v3 = append(v3, fi34)
                                                                                                var v66 int64 = 371
                                                                                                _ = v65
                                                                                                _ = v66
                                                                                            }
                                                                                            fmt.Println(v1, v1)
                                                                                            if (v2 && true) {
                                                                                                var v67 int32 = (v1 * 1)
                                                                                                v2 = ((100 == 100) && (3.14159 > float32(0.0)))
                                                                                                v68 := GenS4{F0: 9223372036854775807}
                                                                                                _ = v67
                                                                                                _ = v68
                                                                                            } else {
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    if v2 {
                                                                                                        fmt.Println(v1, fi15)
                                                                                                        v3 = append(v3, (fi34 << ((1) & 63)))
                                                                                                    }
                                                                                                    fmt.Println(v62, v1, fi15)
                                                                                                    if v2 {
                                                                                                        var v69 float64 = 3.14159
                                                                                                        _ = v69
                                                                                                    }
                                                                                                case "hello":
                                                                                                    if (true && v2) {
                                                                                                        v62 += -32768
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            var v70 uint8 = 42
                                                                                                            v71 := GenS5{F0: v70, F1: -1, F2: 0}
                                                                                                            _ = v70
                                                                                                            _ = v71
                                                                                                        }
                                                                                                    }
                                                                                                case "a":
                                                                                                    switch "a" {
                                                                                                    case "foo bar":
                                                                                                        if v2 {
                                                                                                            v72 := GenS7{F0: 42, F1: -926, F2: 18446744073709551615}
                                                                                                            v2 = (95 < -95)
                                                                                                            _ = v72
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        if len(v3) >= 2 {
                                                                                                            v3 = v3[1:len(v3)]
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v3 = append(v3, fi15)
                                                                                                }
                                                                                                var v73 uint32 = 0
                                                                                                _ = v73
                                                                                            }
                                                                                            _ = v62
                                                                                        } else {
                                                                                            var v74 bool = (true || false)
                                                                                            fmt.Println(v1)
                                                                                            v74 = (true && (fi34 == 9223372036854775807))
                                                                                            _ = v74
                                                                                        }
                                                                                    default:
                                                                                        v1 -= v1
                                                                                    }
                                                                                }
                                                                            case "a":
                                                                                var v75 int8 = 127
                                                                                _ = v75
                                                                            default:
                                                                                fmt.Println(v2)
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    fmt.Println(v2, v2, fi34)
                                                                    if v2 {
                                                                        if (false && v2) {
                                                                            fmt.Println(fi15)
                                                                            fmt.Println(fi34, fi15)
                                                                            v76 := GenS8{F0: 1, F1: 0, F2: 42}
                                                                            _ = v76
                                                                        }
                                                                        v77 := GenS7{F0: 42, F1: -9223372036854775808, F2: 42}
                                                                        if (!v2) {
                                                                            v1 = v1
                                                                            if ((fi15 % (((fi15 * (fi34 / (((-(fi34))) | 1)))) | 1)) > (^(fi34))) {
                                                                                if ((3.14159 / (3.14159 + 0.0)) > (1e10 / 1e10)) {
                                                                                    v2 = v2
                                                                                    fmt.Println(v2)
                                                                                    v2 = ((0.1 < -2.25) || v2)
                                                                                    if true {
                                                                                        v2 = v2
                                                                                    }
                                                                                } else {
                                                                                    if ((((1.5 - 0.0) * 1e10) != (-2.25 * (float32(3.14159) - 1e10))) && false) {
                                                                                        v2 = (133 < 975)
                                                                                        var v78 float64 = (3.14159 + 1e10)
                                                                                        _ = v78
                                                                                    } else {
                                                                                        var v79 float32 = (0.0 / -2.25)
                                                                                        _ = v79
                                                                                    }
                                                                                }
                                                                                var v80 uint64 = 0
                                                                                _ = v80
                                                                            }
                                                                        } else {
                                                                            if len(v4) > 0 {
                                                                                v4[0] = (((("foo bar" + "") + ("\n\t" + "foo bar")) + "foo bar") + ("\n\t" + ""))
                                                                            }
                                                                            var v81 int = 2147483647
                                                                            v1 = (v1 - -810)
                                                                            _ = v81
                                                                        }
                                                                        var v82 uint16 = 32767
                                                                        if v2 {
                                                                            v1 -= (v1 + (v1 / ((v1) | 1)))
                                                                            if len(v4) > 0 {
                                                                                v4[0] = ""
                                                                            }
                                                                            var v83 float32 = 0.0
                                                                            _ = v83
                                                                        }
                                                                        _ = v77
                                                                        _ = v82
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v1 -= (^(v1))
                                                    }
                                                    v84 := GenS1{F0: fi34, F1: v2}
                                                    v85 := GenS7{F0: -1, F1: 1, F2: 42}
                                                    _ = v84
                                                    _ = v85
                                                }
                                                _ = fi34
                                            }
                                        }
                                    }
                                    v86 := GenS1{F0: fi15, F1: v2}
                                    _ = v86
                                } else {
                                    for fi87 := 0; fi87 < 19; fi87++ {
                                        v88 := GenS5{F0: 255, F1: 0, F2: 266}
                                        if false {
                                            var v89 uint = 255
                                            v2 = (v2 && (uint8(v89) < 1))
                                            _ = v89
                                        } else {
                                            if (!((42 > 0) && v2)) {
                                                var v90 float32 = float32(3.14159)
                                                fmt.Println(v90, v2, fi15)
                                                v91 := GenS6{F0: 1, F1: 18446744073709551615, F2: 479}
                                                _ = v90
                                                _ = v91
                                            } else {
                                                fmt.Println(v2, fi87)
                                                var v92 uint = 2
                                                var v93 bool = v2
                                                _ = v92
                                                _ = v93
                                            }
                                            var v94 bool = v2
                                            _ = v94
                                        }
                                        _ = fi87
                                        _ = v88
                                    }
                                }
                                v95 := GenS5{F0: 128, F1: 580, F2: 240}
                                _ = v95
                            }
                        } else {
                            if (872 >= 0) {
                                var v96 float32 = -2.25
                                _ = v96
                            } else {
                                v1 = 0
                            }
                            var v97 uint64 = 4294967295
                            for fi98 := 0; fi98 < 11; fi98++ {
                                switch 0 {
                                case 423:
                                    v2 = ((42 > 2) || true)
                                }
                                v2 = (!(v2 && v2))
                                _ = fi98
                            }
                            _ = v97
                        }
                    }
                    v2 = v2
                    for fi99 := 0; fi99 < 13; fi99++ {
                        switch (fi99 - 358) {
                        case 4294967295:
                            var v100 uint = 2
                            _ = v100
                        case 32768:
                            if len(v3) > 0 {
                                v3[0] = fi99
                            }
                        }
                        if v2 {
                            v2 = (1.5 == 0.0)
                            if v2 {
                                v101 := GenS5{F0: 0, F1: 42, F2: 1}
                                v2 = v2
                                if ((true && v2) || (169 <= 1)) {
                                    fmt.Println(fi99, v1, v1)
                                    var v102 float64 = -2.25
                                    switch -355 {
                                    case 127:
                                        v103 := GenS7{F0: 108, F1: 0, F2: 2}
                                        _ = v103
                                    case 2147483647:
                                        _ = len(v3)
                                    case 4294967295:
                                        var v104 bool = (1 <= 4294967295)
                                        _ = v104
                                    }
                                    _ = v102
                                }
                                _ = v101
                            }
                            if (-9223372036854775808 <= 293) {
                                _ = len(v4)
                                var v105 uint8 = 2
                                _ = len(v3)
                                _ = v105
                            }
                            var v106 bool = v2
                            _ = v106
                        } else {
                            if v2 {
                                if (((127 > 0) && true) || (3.14159 <= 1.5)) {
                                    v4 = append(v4, "foo bar")
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    v107 := GenS4{F0: 950}
                                    v108 := GenS6{F0: 53, F1: 1, F2: 4294967295}
                                    _ = v107
                                    _ = v108
                                }
                                var v109 uint16 = 1
                                var v110 bool = true
                                v110 = (true && v2)
                                _ = v109
                                _ = v110
                            }
                            fmt.Println(v2)
                        }
                        v2 = false
                        _ = fi99
                    }
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi15
            }
        } else {
            for fi111 := 0; fi111 < 16; fi111++ {
                if false {
                    fmt.Println(fi111, fi111)
                    switch (fi111 << ((4294967295) & 63)) {
                    case 0:
                        if false {
                            v4 = append(v4, (((("hello" + "") + "") + "foo bar") + "hello"))
                        } else {
                            if (genFn0() > 1000) {
                                var v112 bool = (-2.25 == float32(0.1))
                                switch fi111 {
                                case 32768:
                                    for fi113 := 0; fi113 < 9; fi113++ {
                                        v2 = (-791 >= fi113)
                                        var v114 uint8 = 2
                                        _ = fi113
                                        _ = v114
                                    }
                                case 442:
                                    v4 = append(v4, "hello")
                                default:
                                    var v115 string = "hello"
                                    _ = v115
                                }
                                if (3.14159 <= float64(0.1)) {
                                    var v116 int16 = -519
                                    v117 := GenS1{F0: fi111, F1: (false || v2)}
                                    _ = v116
                                    _ = v117
                                }
                                _ = v112
                            }
                        }
                    case -981:
                        fmt.Println(v1, v1, v1)
                    case 4294967295:
                        if v2 {
                            v118 := GenS8{F0: 4294967295, F1: 127, F2: 0}
                            _ = v118
                        }
                    default:
                        switch "\n\t" {
                        case "a":
                            _ = len(v3)
                        case "foo bar":
                            v2 = ((true && false) || v2)
                        case "":
                            switch 9223372036854775807 {
                            case 128:
                                for fi119 := 0; fi119 < 2; fi119++ {
                                    if len(v3) > 0 {
                                        v3[0] = (-(fi111))
                                    }
                                    v1 = (-(v1))
                                    _ = fi119
                                }
                            case -866:
                                if v2 {
                                    var v120 uint64 = 42
                                    _ = v120
                                } else {
                                    switch "a" {
                                    case "foo bar":
                                        for fi121 := 0; fi121 < 6; fi121++ {
                                            v4 = append(v4, ("a" + "hello"))
                                            if (!(!v2)) {
                                                fmt.Println(v1)
                                                if (false || ("" != "")) {
                                                    var v122 int = (fi111 % (((-(fi111))) | 1))
                                                    if v2 {
                                                        v2 = (42 == -195)
                                                    } else {
                                                        var v123 int64 = 694
                                                        _ = v123
                                                    }
                                                    v2 = false
                                                    _ = v122
                                                }
                                            } else {
                                                var v124 int64 = -9223372036854775808
                                                v125 := GenS5{F0: 1, F1: (v124 - 42), F2: 65535}
                                                switch fi121 {
                                                case 412:
                                                    var v126 string = ""
                                                    _ = v126
                                                case -9223372036854775808:
                                                    switch 128 {
                                                    case -2147483648:
                                                        var v127 int = 127
                                                        _ = v127
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                default:
                                                    v124 += v124
                                                }
                                                _ = v124
                                                _ = v125
                                            }
                                            v2 = (v2 && ("\n\t" != ("\n\t" + "foo bar")))
                                            _ = fi121
                                        }
                                    }
                                    if v2 {
                                        v128 := GenS9{F0: 1}
                                        _ = v128
                                    } else {
                                        var v129 uint16 = 32767
                                        _ = v129
                                    }
                                    fmt.Println(fi111)
                                }
                            }
                        default:
                            for fi130 := 0; fi130 < 5; fi130++ {
                                v1 = -712
                                if (!v2) {
                                    v131 := GenS9{F0: 2147483647}
                                    _ = v131
                                }
                                if false {
                                    if (!(!v2)) {
                                        var v132 int32 = v1
                                        var v133 int16 = 32767
                                        v134 := GenS9{F0: 42}
                                        fmt.Println(fi111, v132, v132)
                                        _ = v132
                                        _ = v133
                                        _ = v134
                                    } else {
                                        switch (("hello" + "\n\t") + "foo bar") {
                                        case "\n\t":
                                            var v135 uint8 = 1
                                            _ = v135
                                        case "hello":
                                            var v136 int16 = 1000
                                            _ = v136
                                        }
                                        var v137 int16 = 1000
                                        fmt.Println(fi111, v137)
                                        _ = v137
                                    }
                                }
                                _ = fi130
                            }
                        }
                    }
                    var v138 float64 = (0.1 - 1e10)
                    _ = v138
                } else {
                    fmt.Println(fi111, v2)
                    for fi139 := 0; fi139 < 9; fi139++ {
                        v4 = append(v4, "foo bar")
                        switch (fi139 >> ((2) & 63)) {
                        case 2147483647:
                            v140 := GenS8{F0: 781, F1: 128, F2: 42}
                            if (4294967295 == 65535) {
                                if v2 {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v141 int = fi139
                                    v3 = append(v3, -2147483648)
                                    _ = v141
                                } else {
                                    v2 = true
                                    switch "hello" {
                                    case "hello":
                                        if (((0.1 * (0.1 / (3.14159 / (3.14159 + 3.14159)))) / (0.1 / 0.1)) < float64(((3.14159 / float32(1e10)) / (-2.25 / 1.0)))) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            var v142 uint64 = 2
                                            var v143 int32 = v1
                                            _ = v142
                                            _ = v143
                                        } else {
                                            v1 = (-(v1))
                                            v144 := GenS5{F0: 122, F1: 783, F2: 255}
                                            var v145 uint = 4294967295
                                            _ = v144
                                            _ = v145
                                        }
                                    default:
                                        var v146 uint64 = 222
                                        _ = v146
                                    }
                                    v1 = (v1 | (v1 >> ((18446744073709551615) & 63)))
                                }
                                if v2 {
                                    if (false && (!v2)) {
                                        var v147 uint = 283
                                        _ = v147
                                    } else {
                                        v1 -= (^(v1))
                                        if len(v4) > 0 {
                                            v4[0] = "a"
                                        }
                                        var v148 int = fi111
                                        _ = v148
                                    }
                                    var v149 uint16 = 405
                                    v1 -= 2147483647
                                    _ = v149
                                }
                                var v150 float32 = 3.14159
                                _ = v150
                            } else {
                                v2 = (127 >= 57)
                                v2 = v2
                                fmt.Println(v2, fi139)
                            }
                            _ = v140
                        case 32768:
                            if (42 < 1) {
                                var v151 uint16 = 65535
                                v1 = v1
                                _ = v151
                            }
                        }
                        switch fi111 {
                        case 695:
                            fmt.Println(fi139, fi111, fi111)
                        case 4294967295:
                            if ((true || v2) && v2) {
                                switch ((((("foo bar" + (("\n\t" + "foo bar") + "foo bar")) + (("a" + "a") + (((("\n\t" + "\n\t") + ("hello" + ("\n\t" + ("a" + "")))) + "a") + ("foo bar" + (("" + "foo bar") + ((("foo bar" + "hello") + "foo bar") + "a")))))) + "foo bar") + ("\n\t" + "")) + "\n\t") {
                                case "a":
                                    v2 = v2
                                case "hello":
                                    v2 = v2
                                case "":
                                    fmt.Println(v2)
                                default:
                                    v152 := GenS6{F0: 1, F1: 18446744073709551615, F2: 2147483648}
                                    _ = v152
                                }
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        case 9223372036854775807:
                            if (!(!(0 == -1))) {
                                v2 = (((true && v2) && (289 > 4294967295)) && (!(188 < 1)))
                            }
                        default:
                            var v153 string = "a"
                            _ = v153
                        }
                        _ = fi139
                    }
                }
                var v154 string = "hello"
                for fi155 := 0; fi155 < 1; fi155++ {
                    v154 = v154
                    _ = fi155
                }
                _ = fi111
                _ = v154
            }
        }
    default:
        if (1 <= 9223372036854775807) {
            if v2 {
                v1 = (v1 % ((697) | 1))
                var v156 string = "\n\t"
                _ = v156
            }
        } else {
            v1 -= v1
            v157 := GenS10{F0: -396, F1: 9223372036854775807}
            _ = v157
        }
    }
    var v158 int64 = 743
    var v159 float64 = 1.5
    var v160 bool = (!(4294967295 >= 395))
    fmt.Println("v1:", v1)
    fmt.Println("v158:", v158)
    fmt.Println("v159:", v159)
    fmt.Println("v160:", v160)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v158
    _ = v159
    _ = v160
    _ = v2
    _ = v3
    _ = v4
}
