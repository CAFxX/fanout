package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS7 struct { F0 uint64; F1 int }
type GenS8 struct { F0 int8; F1 int16; F2 uint }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int; F1 uint32; F2 int8 }
type GenS11 struct { F0 int8; F1 int16; F2 int8 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint64; F1 int16 }
type GenS14 struct { F0 int32; F1 uint; F2 int16 }
type GenS15 struct { F0 bool; F1 int16 }
type GenS16 struct { F0 int32; F1 uint32; F2 uint }
type GenS17 struct { F0 string; F1 uint16; F2 int16 }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS20 struct { F0 int16; F1 int; F2 uint8 }
type GenS21 struct { F0 int64 }
type GenS22 struct { F0 uint16; F1 bool; F2 uint }
type GenS23 struct { F0 int; F1 uint32 }
type GenS24 struct { F0 uint8; F1 uint; F2 string }
type GenS25 struct { F0 uint16 }
type GenS26 struct { F0 bool }
type GenS27 struct { F0 int; F1 string }
type GenS28 struct { F0 uint32; F1 uint8 }
type GenS29 struct { F0 int16; F1 uint32 }
type GenS30 struct { F0 string; F1 uint32; F2 uint }
type GenS31 struct { F0 int64; F1 int16 }
type GenS32 struct { F0 int16; F1 uint32; F2 string }
type GenS33 struct { F0 int8; F1 int16; F2 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 42
    v3 := make(map[string]uint32)
    v4 := []uint32{}
    v5 := []int{127}
    var v6 uint16 = 32767
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(v6, fi7)
        if (fi7 != fi7) {
            if len(v5) > 0 {
                v5[0] = fi7
            }
            for fi8 := 0; fi8 < 11; fi8++ {
                v6 = 984
                if len(v5) > 0 {
                    v5[0] = fi7
                }
                _ = fi8
            }
        }
        var v9 int32 = 0
        _ = fi7
        _ = v9
    }
    if (834 == -2147483648) {
        for fi10 := 0; fi10 < 16; fi10++ {
            v2 += (v2 | (^(v2)))
            switch ("foo bar" + ("\n\t" + "foo bar")) {
            case "hello":
                if (!false) {
                    var v11 int = (fi10 % (((fi10 % ((32767) | 1))) | 1))
                    v12 := v3["hello"]
                    var v13 bool = ((v2 & v2) < v2)
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    v2 = (v2 / (((^(v2))) | 1))
                    var v14 int = (fi10 | 42)
                    _ = v14
                }
            default:
                var v15 bool = true
                _ = v15
            }
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS1{F0: 1}
                var v18 uint8 = 127
                fmt.Println(v18)
                switch "\n\t" {
                case "\n\t":
                    v19 := GenS1{F0: 100}
                    _ = v19
                default:
                    if (false || true) {
                        var v20 float32 = 1e10
                        v21 := v3[((("\n\t" + ("hello" + "")) + ("foo bar" + ("" + ("\n\t" + "\n\t")))) + "a")]
                        if false {
                            v21 += 0
                            var v22 int32 = 2147483647
                            _ = v22
                        }
                        fmt.Println(fi10, v18)
                        var v23 uint8 = v18
                        _ = v20
                        _ = v21
                        _ = v23
                    }
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi10
        }
        if true {
            var v24 int32 = 2147483647
            _ = v24
        }
        fmt.Println(v6, v6)
    }
    v6 += (v6 - v6)
    if true {
        v6 += (v6 + v6)
        if (!(9223372036854775807 <= -218)) {
            _ = len(v5)
            v6 += 65535
            v25 := GenS1{F0: 100}
            _ = v25
        } else {
            for fi26 := 0; fi26 < 8; fi26++ {
                v2 += 4294967295
                _ = fi26
            }
            _ = len(v4)
        }
    } else {
        fmt.Println(v2)
        for fi27 := 0; fi27 < 7; fi27++ {
            for fi28 := 0; fi28 < 9; fi28++ {
                v6 += v6
                _ = fi28
            }
            v6 += v6
            _ = fi27
        }
        v29 := GenS2{F0: 577}
        _ = v29
    }
    delete(v1, "hello")
    fmt.Println(v6)
    for fi30 := 0; fi30 < 11; fi30++ {
        var v31 float64 = 0.0
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        for fi32 := 0; fi32 < 6; fi32++ {
            v33 := GenS3{F0: -1}
            _ = fi32
            _ = v33
        }
        _ = fi30
        _ = v31
    }
    if (65535 != 2) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v34 uint32 = 0
        var v35 float32 = 0.1
        _ = v34
        _ = v35
    }
    if true {
        var v36 uint32 = 2147483648
        if len(v5) > 0 {
            v5[0] = 128
        }
        if ((v36 >> ((0) & 63)) <= (v36 << ((65535) & 63))) {
            v36 = (v36 << ((255) & 63))
            var v37 int16 = 0
            for fi38 := 0; fi38 < 20; fi38++ {
                _ = len(v3)
                _ = fi38
            }
            _ = v37
        }
        _ = v36
    }
    if (false && (2 < 0)) {
        var v39 uint16 = v6
        v40 := GenS1{F0: 42}
        _ = v39
        _ = v40
    } else {
        fmt.Println(v2, v2)
        v41 := GenS4{F0: 0}
        switch "\n\t" {
        case "a":
            v6 = (v6 + 418)
        case "\n\t":
            var v42 float32 = (1.5 - 1e10)
            _ = v42
        case "foo bar":
            for fi43 := 0; fi43 < 13; fi43++ {
                for fi44 := 0; fi44 < 8; fi44++ {
                    if true {
                        var v45 string = "foo bar"
                        switch (fi43 >> ((1) & 63)) {
                        case 0:
                            var v46 int64 = 9223372036854775807
                            _ = v46
                        default:
                            switch 9223372036854775807 {
                            case -907:
                                v2 = (^(v2))
                            case 0:
                                var v47 bool = ((-2.25 >= (1.5 + (-2.25 * 3.14159))) && false)
                                _ = v47
                            default:
                                var v48 uint8 = 128
                                _ = v48
                            }
                        }
                        v49 := GenS5{F0: 42}
                        _ = v45
                        _ = v49
                    } else {
                        var v50 uint64 = v2
                        if true {
                            if (false && ((^(v2)) > (v2 * (^(v50))))) {
                                var v51 int16 = 42
                                var v52 bool = false
                                _ = v51
                                _ = v52
                            }
                            v2 = (v50 + (^(v50)))
                            fmt.Println(v6)
                        } else {
                            var v53 int64 = -9223372036854775808
                            _ = v53
                        }
                        _ = v50
                    }
                    if len(v5) > 0 {
                        v5[0] = (fi43 - 4294967295)
                    }
                    var v54 uint64 = v2
                    _ = fi44
                    _ = v54
                }
                if true {
                    if (false || (344 < 42)) {
                        if (true || true) {
                            for fi55 := 0; fi55 < 20; fi55++ {
                                switch fi43 {
                                case 773:
                                    v6 += v6
                                case 32768:
                                    var v56 uint64 = 18446744073709551615
                                    _ = v56
                                case 32767:
                                    if (true && (false && (!("hello" != "a")))) {
                                        var v57 int = (fi43 % ((9223372036854775807) | 1))
                                        _ = v57
                                    } else {
                                        var v58 uint64 = (^(v2))
                                        switch "foo bar" {
                                        case "a":
                                            _ = len(v3)
                                        case "":
                                            if false {
                                                v2 = 4294967295
                                                v58 += (v2 >> ((4294967295) & 63))
                                                v59 := GenS3{F0: 42}
                                                var v60 bool = false
                                                _ = v59
                                                _ = v60
                                            } else {
                                                if (1.5 == -2.25) {
                                                    v6 += v6
                                                    if (255 < 0) {
                                                        if false {
                                                            var v61 uint8 = 127
                                                            fmt.Println(v61, v6, v58)
                                                            fmt.Println(v58, fi55, v6)
                                                            _ = v61
                                                        } else {
                                                            if false {
                                                                if true {
                                                                    if ((false && false) && false) {
                                                                        var v62 float32 = (-2.25 - -2.25)
                                                                        _ = v62
                                                                    } else {
                                                                        var v63 float64 = (float64(3.14159) - 0.1)
                                                                        _ = v63
                                                                    }
                                                                } else {
                                                                    switch (("a" + "hello") + ("foo bar" + "")) {
                                                                    case "":
                                                                        v1[((("\n\t" + ("foo bar" + ("" + "foo bar"))) + "") + ("hello" + ("hello" + "")))] = 1
                                                                    case "foo bar":
                                                                        if ((2147483648 != 743) || (0 <= 127)) {
                                                                            fmt.Println(v2)
                                                                        } else {
                                                                            v2 = 320
                                                                        }
                                                                    default:
                                                                        v2 = 2
                                                                    }
                                                                }
                                                            } else {
                                                                var v64 int32 = 1
                                                                v3[("" + "\n\t")] = 2147483647
                                                                _ = v64
                                                            }
                                                            switch 0 {
                                                            case 2147483647:
                                                                if false {
                                                                    v58 -= (^(v58))
                                                                    var v65 float64 = (-2.25 * 1e10)
                                                                    var v66 float64 = 0.0
                                                                    _ = v65
                                                                    _ = v66
                                                                } else {
                                                                    if true {
                                                                        fmt.Println(v6)
                                                                        switch 957 {
                                                                        case 266:
                                                                            v2 -= (v2 + (v58 | v2))
                                                                        case -1:
                                                                            v58 = v58
                                                                        case -9223372036854775808:
                                                                            var v67 int8 = 1
                                                                            _ = v67
                                                                        }
                                                                        var v68 float64 = (float64(1e10) + (1.5 - ((-2.25 / 1.0) - ((3.14159 * (-2.25 * float64((((1.5 - (-2.25 - float32(((0.0 * (float64((3.14159 / 1.0)) - ((1e10 * (1.5 - -2.25)) - -2.25))) + 0.0)))) + ((0.0 * -2.25) + (float32(1e10) + -2.25))) / 0.1)))) - ((0.0 * (3.14159 * 0.1)) - 1e10)))))
                                                                        _ = v68
                                                                    } else {
                                                                        if (false || (false || (0 < 1))) {
                                                                            switch ("foo bar" + ("hello" + "a")) {
                                                                            case "\n\t":
                                                                                v69 := GenS3{F0: -128}
                                                                                _ = v69
                                                                            }
                                                                            v6 = 2
                                                                        } else {
                                                                            v70 := GenS5{F0: 115}
                                                                            if false {
                                                                                if (1 < -205) {
                                                                                    if false {
                                                                                        if false {
                                                                                            v6 += v6
                                                                                            fmt.Println(fi55, v58, fi55)
                                                                                        }
                                                                                        v71 := v1["\n\t"]
                                                                                        _ = v71
                                                                                    } else {
                                                                                        fmt.Println(fi43)
                                                                                        var v72 int = (fi55 / (((fi55 + fi55)) | 1))
                                                                                        switch v72 {
                                                                                        case -9223372036854775808:
                                                                                            v2 += v58
                                                                                        case 139:
                                                                                            fmt.Println(v2)
                                                                                        case 32768:
                                                                                            if false {
                                                                                                var v73 float64 = 0.0
                                                                                                v72 = fi55
                                                                                                _ = v73
                                                                                            }
                                                                                        default:
                                                                                            switch fi43 {
                                                                                            case 445:
                                                                                                fmt.Println(fi55)
                                                                                                if (!true) {
                                                                                                    fmt.Println(fi43, v72, v72)
                                                                                                    var v74 string = "hello"
                                                                                                    v75 := GenS1{F0: 127}
                                                                                                    _ = v74
                                                                                                    _ = v75
                                                                                                } else {
                                                                                                    v2 = (v2 % (((v58 - (^(v58)))) | 1))
                                                                                                    if false {
                                                                                                        var v76 int64 = 42
                                                                                                        v2 = 542
                                                                                                        _ = v76
                                                                                                    } else {
                                                                                                        v1["a"] = 65535
                                                                                                        var v77 int32 = -1
                                                                                                        _ = v77
                                                                                                    }
                                                                                                    if (2 <= 724) {
                                                                                                        if (((v6 % ((v6) | 1)) <= 415) && (!((!true) && false))) {
                                                                                                            v78 := GenS1{F0: 127}
                                                                                                            var v79 float64 = 1e10
                                                                                                            if (true || true) {
                                                                                                                v80 := GenS4{F0: -32768}
                                                                                                                if ((!true) || false) {
                                                                                                                    v58 = 1
                                                                                                                    if false {
                                                                                                                        switch "\n\t" {
                                                                                                                        case "":
                                                                                                                            if false {
                                                                                                                                switch "\n\t" {
                                                                                                                                case "\n\t":
                                                                                                                                    v81 := GenS6{F0: 628, F1: v6, F2: 31}
                                                                                                                                    _ = v81
                                                                                                                                case "hello":
                                                                                                                                    v2 = (v2 ^ 557)
                                                                                                                                case "":
                                                                                                                                    if (1 > 4294967295) {
                                                                                                                                        v82 := GenS5{F0: -32768}
                                                                                                                                        if true {
                                                                                                                                            var v83 uint8 = 131
                                                                                                                                            v84 := GenS4{F0: 32767}
                                                                                                                                            v85 := GenS2{F0: 1}
                                                                                                                                            v79 = float64(1e10)
                                                                                                                                            if len(v4) > 0 {
                                                                                                                                                v4[0] = 1
                                                                                                                                            }
                                                                                                                                            _ = v83
                                                                                                                                            _ = v84
                                                                                                                                            _ = v85
                                                                                                                                        } else {
                                                                                                                                            v2 += v2
                                                                                                                                            v86 := v1["foo bar"]
                                                                                                                                            fmt.Println(v58, fi55)
                                                                                                                                            var v87 string = "hello"
                                                                                                                                            _ = v86
                                                                                                                                            _ = v87
                                                                                                                                        }
                                                                                                                                        if (!((1.5 + (v79 + v79)) <= 0.0)) {
                                                                                                                                            fmt.Println(v72, fi55)
                                                                                                                                            v72 += v72
                                                                                                                                        }
                                                                                                                                        if (v79 != 0.1) {
                                                                                                                                            v6 = (v6 ^ (v6 / (((v6 + (v6 >> ((0) & 63)))) | 1)))
                                                                                                                                            var v88 uint16 = 286
                                                                                                                                            _ = v88
                                                                                                                                        }
                                                                                                                                        _ = v82
                                                                                                                                    } else {
                                                                                                                                        v3["a"] = 817
                                                                                                                                    }
                                                                                                                                    var v89 string = "foo bar"
                                                                                                                                    switch (v89 + v89) {
                                                                                                                                    case "a":
                                                                                                                                        if (v89 != v89) {
                                                                                                                                            v2 = 1
                                                                                                                                        } else {
                                                                                                                                            fmt.Println(v58)
                                                                                                                                            delete(v1, v89)
                                                                                                                                        }
                                                                                                                                    case "hello":
                                                                                                                                        fmt.Println(fi55, v89, v79)
                                                                                                                                    case "foo bar":
                                                                                                                                        var v90 float64 = ((v79 * 0.0) / (v79 / float64(1.5)))
                                                                                                                                        _ = v90
                                                                                                                                    }
                                                                                                                                    _ = v89
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                fmt.Println(v2)
                                                                                                                                var v91 int8 = 100
                                                                                                                                _ = v91
                                                                                                                            }
                                                                                                                        case "hello":
                                                                                                                            var v92 uint64 = v2
                                                                                                                            _ = v92
                                                                                                                        case "\n\t":
                                                                                                                            var v93 int8 = 0
                                                                                                                            _ = v93
                                                                                                                        default:
                                                                                                                            fmt.Println(v2)
                                                                                                                        }
                                                                                                                        var v94 int16 = -32768
                                                                                                                        if (!((v79 * 1e10) > 0.1)) {
                                                                                                                            if false {
                                                                                                                                var v95 bool = (3.14159 >= (v79 + v79))
                                                                                                                                var v96 uint32 = 854
                                                                                                                                _ = v95
                                                                                                                                _ = v96
                                                                                                                            } else {
                                                                                                                                var v97 int64 = -9223372036854775808
                                                                                                                                var v98 uint16 = (v6 & 1)
                                                                                                                                v79 = (v79 + v79)
                                                                                                                                _ = v97
                                                                                                                                _ = v98
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            if ((!true) || (!true)) {
                                                                                                                                fmt.Println(v79)
                                                                                                                                v2 += v2
                                                                                                                                if false {
                                                                                                                                    v58 = v58
                                                                                                                                    var v99 uint = 0
                                                                                                                                    if len(v4) > 0 {
                                                                                                                                        v4[0] = 1
                                                                                                                                    }
                                                                                                                                    _ = v99
                                                                                                                                }
                                                                                                                                fmt.Println(fi43, v6, fi55)
                                                                                                                            } else {
                                                                                                                                var v100 uint32 = 65535
                                                                                                                                v101 := GenS5{F0: v94}
                                                                                                                                _ = len(v5)
                                                                                                                                _ = v100
                                                                                                                                _ = v101
                                                                                                                            }
                                                                                                                            if true {
                                                                                                                                v2 += v58
                                                                                                                            }
                                                                                                                        }
                                                                                                                        if len(v5) >= 2 {
                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                        }
                                                                                                                        _ = v94
                                                                                                                    } else {
                                                                                                                        if false {
                                                                                                                            var v102 int8 = -1
                                                                                                                            _ = v102
                                                                                                                        } else {
                                                                                                                            v103 := v1["\n\t"]
                                                                                                                            var v104 uint = 144
                                                                                                                            _ = v103
                                                                                                                            _ = v104
                                                                                                                        }
                                                                                                                        v72 += v72
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v105 := v1[""]
                                                                                                                    _ = v105
                                                                                                                }
                                                                                                                if true {
                                                                                                                    if true {
                                                                                                                        if true {
                                                                                                                            var v106 float64 = 1.5
                                                                                                                            if ("hello" != "foo bar") {
                                                                                                                                _ = len(v3)
                                                                                                                                v6 = (v6 | (v6 + v6))
                                                                                                                                switch "\n\t" {
                                                                                                                                case "a":
                                                                                                                                    v58 -= v58
                                                                                                                                case "\n\t":
                                                                                                                                    v6 = (v6 + 42)
                                                                                                                                case "":
                                                                                                                                    fmt.Println(v106, v2)
                                                                                                                                default:
                                                                                                                                    var v107 float64 = ((1e10 + 0.0) + v79)
                                                                                                                                    _ = v107
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = v106
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        var v108 uint32 = 1
                                                                                                                        if false {
                                                                                                                            var v109 int32 = -2147483648
                                                                                                                            _ = v109
                                                                                                                        }
                                                                                                                        if len(v5) >= 2 {
                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                        }
                                                                                                                        _ = v108
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    var v110 uint = 65535
                                                                                                                    v2 = v2
                                                                                                                    v58 = 2
                                                                                                                    _ = v110
                                                                                                                }
                                                                                                                _ = v80
                                                                                                            }
                                                                                                            if true {
                                                                                                                var v111 int64 = 9223372036854775807
                                                                                                                v112 := GenS4{F0: -32768}
                                                                                                                v113 := GenS7{F0: 4294967295, F1: (fi43 | (fi55 - -731))}
                                                                                                                var v114 string = "hello"
                                                                                                                if true {
                                                                                                                    switch (-(fi55)) {
                                                                                                                    case 2147483647:
                                                                                                                        var v115 int32 = -2147483648
                                                                                                                        _ = v115
                                                                                                                    case 128:
                                                                                                                        v116 := GenS5{F0: 1000}
                                                                                                                        _ = v116
                                                                                                                    case 711:
                                                                                                                        if (!true) {
                                                                                                                            v117 := GenS7{F0: v2, F1: (fi43 * (v72 / ((v72) | 1)))}
                                                                                                                            v118 := GenS6{F0: (v6 >> ((42) & 63)), F1: uint16(v2), F2: 127}
                                                                                                                            v114 = v114
                                                                                                                            _ = v117
                                                                                                                            _ = v118
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v119 uint = 757
                                                                                                                        _ = v119
                                                                                                                    }
                                                                                                                    fmt.Println(v111)
                                                                                                                    v2 += v2
                                                                                                                } else {
                                                                                                                    v120 := v3[((v114 + v114) + "hello")]
                                                                                                                    var v121 int8 = 0
                                                                                                                    _ = v120
                                                                                                                    _ = v121
                                                                                                                }
                                                                                                                v122 := GenS8{F0: -6, F1: 32767, F2: 689}
                                                                                                                _ = v111
                                                                                                                _ = v112
                                                                                                                _ = v113
                                                                                                                _ = v114
                                                                                                                _ = v122
                                                                                                            } else {
                                                                                                                var v123 uint = 324
                                                                                                                v123 -= v123
                                                                                                                var v124 int8 = 34
                                                                                                                _ = v123
                                                                                                                _ = v124
                                                                                                            }
                                                                                                            v125 := GenS6{F0: v6, F1: (v6 | v6), F2: -41}
                                                                                                            var v126 uint8 = 127
                                                                                                            _ = v125
                                                                                                            _ = v126
                                                                                                            _ = v78
                                                                                                            _ = v79
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                v58 = (v58 & v58)
                                                                                                            }
                                                                                                            fmt.Println(v6)
                                                                                                            var v127 uint64 = v2
                                                                                                            _ = v127
                                                                                                        }
                                                                                                        if (0 > 648) {
                                                                                                            var v128 int32 = 0
                                                                                                            switch ((("\n\t" + "") + ("\n\t" + ("hello" + ""))) + ((("" + "a") + "") + "hello")) {
                                                                                                            case "a":
                                                                                                                _ = len(v3)
                                                                                                            case "\n\t":
                                                                                                                var v129 int = (^(fi43))
                                                                                                                _ = v129
                                                                                                            }
                                                                                                            _ = v128
                                                                                                        } else {
                                                                                                            delete(v3, "hello")
                                                                                                        }
                                                                                                    }
                                                                                                    v6 -= (v6 >> ((286) & 63))
                                                                                                    if false {
                                                                                                        switch 0 {
                                                                                                        case -982:
                                                                                                            if (!((("a" + "") + "\n\t") >= "")) {
                                                                                                                var v130 uint = 65535
                                                                                                                v2 = 851
                                                                                                                _ = v130
                                                                                                            } else {
                                                                                                                if (!((!(-9223372036854775808 <= 497)) && false)) {
                                                                                                                    v131 := GenS5{F0: 32767}
                                                                                                                    _ = v131
                                                                                                                }
                                                                                                                var v132 int = 128
                                                                                                                var v133 uint32 = 2147483648
                                                                                                                var v134 uint32 = v133
                                                                                                                var v135 uint16 = 63
                                                                                                                _ = v132
                                                                                                                _ = v133
                                                                                                                _ = v134
                                                                                                                _ = v135
                                                                                                            }
                                                                                                        case 32767:
                                                                                                            v2 += v2
                                                                                                        default:
                                                                                                            v2 = 0
                                                                                                        }
                                                                                                        v136 := GenS9{F0: 0}
                                                                                                        _ = v136
                                                                                                    }
                                                                                                }
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = (fi55 + (fi43 | (v72 & (v72 >> ((2) & 63)))))
                                                                                                }
                                                                                            case 94:
                                                                                                if ((false || ((!true) && (uint64(v6) == 42))) || (-1 != 100)) {
                                                                                                    if ((v2 | (v58 >> ((18446744073709551615) & 63))) != v58) {
                                                                                                        v6 += (v6 >> ((657) & 63))
                                                                                                        v2 = v2
                                                                                                    }
                                                                                                } else {
                                                                                                    v137 := GenS10{F0: -2147483648, F1: 65535, F2: 127}
                                                                                                    v3["a"] = 458
                                                                                                    _ = v137
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v72
                                                                                    }
                                                                                    v6 = (v6 + 695)
                                                                                } else {
                                                                                    var v138 uint32 = 555
                                                                                    v58 += (v2 / (((v58 | v58)) | 1))
                                                                                    _ = v138
                                                                                }
                                                                                v2 -= (v58 ^ v2)
                                                                            } else {
                                                                                v2 = (v2 & v58)
                                                                            }
                                                                            _ = v70
                                                                        }
                                                                        fmt.Println(fi43, fi43)
                                                                        switch (fi43 + int(v58)) {
                                                                        case 4294967295:
                                                                            switch (fi55 | fi55) {
                                                                            case -9223372036854775808:
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = 2
                                                                                }
                                                                                var v139 int = (fi43 ^ (fi43 - 9223372036854775807))
                                                                                switch 128 {
                                                                                case -2147483648:
                                                                                    delete(v1, "hello")
                                                                                case 73:
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                case 264:
                                                                                    v140 := v1["\n\t"]
                                                                                    _ = v140
                                                                                default:
                                                                                    if false {
                                                                                        switch (-(fi43)) {
                                                                                        case 127:
                                                                                            v6 = v6
                                                                                        case -357:
                                                                                            var v141 float32 = 1.5
                                                                                            _ = v141
                                                                                        case 32767:
                                                                                            var v142 int32 = 42
                                                                                            _ = v142
                                                                                        }
                                                                                        _ = len(v1)
                                                                                    }
                                                                                    if ((0.0 <= (0.0 / (3.14159 / (3.14159 * (-2.25 * 1.5))))) || true) {
                                                                                        var v143 float32 = (0.1 * 0.0)
                                                                                        fmt.Println(v6, v143, fi55)
                                                                                        v58 = v58
                                                                                        _ = v143
                                                                                    }
                                                                                }
                                                                                switch ("a" + "") {
                                                                                case "":
                                                                                    v144 := GenS4{F0: -501}
                                                                                    _ = v144
                                                                                case "hello":
                                                                                    var v145 float32 = ((-2.25 - 1.5) - 0.0)
                                                                                    _ = v145
                                                                                case "a":
                                                                                    fmt.Println(v2, v58, fi43)
                                                                                }
                                                                                fmt.Println(fi43)
                                                                                _ = v139
                                                                            case 1:
                                                                                v58 -= v58
                                                                            default:
                                                                                var v146 int = (fi43 / ((2147483647) | 1))
                                                                                _ = v146
                                                                            }
                                                                        case -735:
                                                                            switch fi43 {
                                                                            case 1:
                                                                                var v147 uint64 = 340
                                                                                _ = v147
                                                                            case -2147483648:
                                                                                var v148 int8 = -128
                                                                                _ = v148
                                                                            case 4294967295:
                                                                                var v149 uint8 = 0
                                                                                _ = v149
                                                                            }
                                                                        case -1:
                                                                            var v150 uint = 1
                                                                            _ = v150
                                                                        }
                                                                    }
                                                                    if ((3.14159 - float32((0.0 - 0.1))) != (0.1 / 1e10)) {
                                                                        if true {
                                                                            var v151 uint = 42
                                                                            fmt.Println(v58, v151, v6)
                                                                            _ = v151
                                                                        } else {
                                                                            v6 -= v6
                                                                            switch fi43 {
                                                                            case -871:
                                                                                var v152 float64 = 1.5
                                                                                _ = v152
                                                                            case 32767:
                                                                                fmt.Println(v6, fi55)
                                                                            default:
                                                                                var v153 string = "hello"
                                                                                _ = v153
                                                                            }
                                                                            _ = len(v5)
                                                                            fmt.Println(fi43)
                                                                        }
                                                                    } else {
                                                                        if ((0 != -1) || false) {
                                                                            v154 := GenS11{F0: int8(fi55), F1: -907, F2: -128}
                                                                            v4 = append(v4, uint32(v6))
                                                                            v6 = (v6 % (((v6 * (v6 / ((42) | 1)))) | 1))
                                                                            _ = v154
                                                                        }
                                                                        v58 -= 18446744073709551615
                                                                    }
                                                                    v3["\n\t"] = 0
                                                                }
                                                            case -807:
                                                                if false {
                                                                    v3["a"] = 2
                                                                    var v155 bool = (49 == 893)
                                                                    _ = v155
                                                                } else {
                                                                    var v156 uint32 = 4294967295
                                                                    _ = v156
                                                                }
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    v6 = v6
                                                                case "":
                                                                    switch (fi43 ^ (fi43 >> ((0) & 63))) {
                                                                    case -2147483648:
                                                                        if true {
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                            fmt.Println(v2)
                                                                            var v157 int = fi55
                                                                            v158 := GenS5{F0: -32768}
                                                                            fmt.Println(v157)
                                                                            _ = v157
                                                                            _ = v158
                                                                        }
                                                                    case 4294967295:
                                                                        v159 := GenS10{F0: 1, F1: 624, F2: 42}
                                                                        switch "" {
                                                                        case "hello":
                                                                            if true {
                                                                                fmt.Println(v2)
                                                                                v160 := GenS12{F0: v6}
                                                                                v58 = (v58 << ((255) & 63))
                                                                                _ = v160
                                                                            } else {
                                                                                var v161 bool = false
                                                                                _ = v161
                                                                            }
                                                                        case "":
                                                                            v162 := GenS10{F0: (fi43 + (fi43 & fi55)), F1: 2, F2: 1}
                                                                            _ = v162
                                                                        case "foo bar":
                                                                            var v163 int32 = 2147483647
                                                                            _ = v163
                                                                        default:
                                                                            var v164 uint = 2
                                                                            _ = v164
                                                                        }
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v3[("" + ("" + "foo bar"))] = 395
                                                                        default:
                                                                            if true {
                                                                                var v165 int = fi43
                                                                                switch "hello" {
                                                                                case "a":
                                                                                    v58 = v2
                                                                                    if (0 <= 236) {
                                                                                        switch "a" {
                                                                                        case "":
                                                                                            var v166 int8 = -128
                                                                                            _ = v166
                                                                                        case "\n\t":
                                                                                            v4 = append(v4, 1)
                                                                                        }
                                                                                        var v167 float32 = 1.5
                                                                                        v167 = v167
                                                                                        _ = v167
                                                                                    }
                                                                                    if ((v58 * 42) > 500) {
                                                                                        v168 := GenS13{F0: v2, F1: 779}
                                                                                        v6 += 255
                                                                                        var v169 string = ""
                                                                                        if false {
                                                                                            switch (fi55 << ((979) & 63)) {
                                                                                            case 32767:
                                                                                                _ = len(v1)
                                                                                            case -2147483648:
                                                                                                var v170 uint = 4294967295
                                                                                                _ = v170
                                                                                            case 128:
                                                                                                var v171 int64 = 0
                                                                                                _ = v171
                                                                                            }
                                                                                            if ((0.1 / 1e10) != 0.1) {
                                                                                                fmt.Println(v6, fi43, fi43)
                                                                                                if false {
                                                                                                    var v172 uint16 = v6
                                                                                                    v173 := GenS4{F0: 42}
                                                                                                    fmt.Println(fi43)
                                                                                                    _ = v172
                                                                                                    _ = v173
                                                                                                } else {
                                                                                                    v165 -= (fi55 << ((1) & 63))
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v165, v165)
                                                                                        }
                                                                                        _ = len(v1)
                                                                                        _ = v168
                                                                                        _ = v169
                                                                                    } else {
                                                                                        var v174 uint32 = 2147483648
                                                                                        var v175 float64 = (float64(0.1) * (0.0 * (((float64((1.5 + 1.5)) + 0.1) - (1e10 / ((((3.14159 * (0.0 + (float64(0.1) * 0.1))) + float64(-2.25)) / (((-2.25 - (3.14159 - (((0.0 - 0.1) - 0.0) * 0.1))) + 0.1) * 3.14159)) / ((0.0 - 3.14159) - float64((float32((0.0 * 0.1)) * (float32(1.5) / 1.0))))))) / 3.14159)))
                                                                                        v6 -= 1
                                                                                        _ = v174
                                                                                        _ = v175
                                                                                    }
                                                                                case "hello":
                                                                                    fmt.Println(v2)
                                                                                case "foo bar":
                                                                                    var v176 uint16 = 1
                                                                                    _ = v176
                                                                                default:
                                                                                    var v177 string = ("a" + "\n\t")
                                                                                    _ = v177
                                                                                }
                                                                                if (0.0 > (0.1 - -2.25)) {
                                                                                    var v178 int = 9223372036854775807
                                                                                    fmt.Println(v165, v2, v58)
                                                                                    v165 -= 77
                                                                                    _ = v178
                                                                                } else {
                                                                                    if ("" < "a") {
                                                                                        v6 = v6
                                                                                        var v179 int32 = 42
                                                                                        var v180 int16 = -32768
                                                                                        v181 := GenS14{F0: 2147483647, F1: 2, F2: (^(v180))}
                                                                                        _ = v179
                                                                                        _ = v180
                                                                                        _ = v181
                                                                                    } else {
                                                                                        if (!((0.1 + (0.0 / 0.1)) == (1e10 / -2.25))) {
                                                                                            var v182 bool = true
                                                                                            switch (v165 * fi43) {
                                                                                            case -530:
                                                                                                var v183 uint32 = 65535
                                                                                                _ = v183
                                                                                            }
                                                                                            v184 := GenS3{F0: 127}
                                                                                            _ = v182
                                                                                            _ = v184
                                                                                        } else {
                                                                                            var v185 int32 = -550
                                                                                            var v186 int16 = 1000
                                                                                            v187 := GenS1{F0: -1}
                                                                                            _ = v185
                                                                                            _ = v186
                                                                                            _ = v187
                                                                                        }
                                                                                        v4 = append(v4, 472)
                                                                                        v165 = -186
                                                                                    }
                                                                                }
                                                                                _ = v165
                                                                            }
                                                                        }
                                                                        v2 = (v2 << ((255) & 63))
                                                                        _ = v159
                                                                    case -596:
                                                                        v2 = v2
                                                                    default:
                                                                        if true {
                                                                            var v188 uint = 65535
                                                                            _ = v188
                                                                        } else {
                                                                            _ = len(v1)
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    switch (fi43 & fi55) {
                                                                    case 128:
                                                                        v189 := GenS14{F0: 421, F1: 1, F2: 0}
                                                                        _ = v189
                                                                    case 2147483647:
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 1
                                                                        }
                                                                    }
                                                                default:
                                                                    var v190 int32 = -785
                                                                    _ = v190
                                                                }
                                                            }
                                                            v6 = v6
                                                        }
                                                        _ = len(v5)
                                                        v58 = (^(v2))
                                                    }
                                                    v191 := GenS15{F0: false, F1: 866}
                                                    _ = v191
                                                }
                                                var v192 int32 = 1
                                                var v193 uint8 = 15
                                                v194 := v1[("hello" + "a")]
                                                _ = v192
                                                _ = v193
                                                _ = v194
                                            }
                                        case "hello":
                                            switch fi55 {
                                            case -828:
                                                v1["hello"] = 973
                                            case -2147483648:
                                                var v195 uint16 = (v6 ^ 973)
                                                _ = v195
                                            case 32767:
                                                v58 -= 18446744073709551615
                                            }
                                        }
                                        v2 += 2
                                        _ = len(v1)
                                        _ = v58
                                    }
                                    switch fi43 {
                                    case 0:
                                        switch "\n\t" {
                                        case "a":
                                            if true {
                                                var v196 int8 = -128
                                                v197 := GenS16{F0: 0, F1: 2147483647, F2: 65535}
                                                var v198 uint16 = v6
                                                var v199 bool = false
                                                switch "\n\t" {
                                                case "":
                                                    var v200 uint32 = 9
                                                    _ = v200
                                                case "hello":
                                                    var v201 uint8 = 127
                                                    _ = v201
                                                case "a":
                                                    switch 32768 {
                                                    case 762:
                                                        var v202 float32 = 1.5
                                                        _ = v202
                                                    case -690:
                                                        v203 := GenS17{F0: "a", F1: (^(v6)), F2: 1}
                                                        _ = v203
                                                    default:
                                                        if (((^(v198)) < 861) || v199) {
                                                            switch (("" + "a") + ("foo bar" + ("" + "hello"))) {
                                                            case "":
                                                                var v204 uint = 255
                                                                _ = v204
                                                            case "hello":
                                                                v198 = (^(v6))
                                                            }
                                                            v205 := GenS18{F0: -1}
                                                            if (("hello" > "") && true) {
                                                                if ((!v199) || v199) {
                                                                    v206 := GenS19{F0: -293, F1: 0, F2: v198}
                                                                    _ = v206
                                                                } else {
                                                                    if (true || (!(544 >= 42))) {
                                                                        if (-32768 == 1000) {
                                                                            v198 += 203
                                                                            var v207 float32 = 0.1
                                                                            _ = v207
                                                                        }
                                                                        var v208 int64 = 868
                                                                        if true {
                                                                            var v209 int16 = -129
                                                                            switch ("a" + ("a" + "\n\t")) {
                                                                            case "\n\t":
                                                                                v196 = v196
                                                                            case "a":
                                                                                if v199 {
                                                                                    if v199 {
                                                                                        var v210 int64 = v208
                                                                                        if ("hello" > "\n\t") {
                                                                                            v198 -= (^(v6))
                                                                                        } else {
                                                                                            var v211 bool = (128 < fi43)
                                                                                            v198 = (v198 ^ (v198 * 0))
                                                                                            _ = v211
                                                                                        }
                                                                                        v3[(("\n\t" + ("a" + ("hello" + ""))) + "foo bar")] = 998
                                                                                        v208 -= 42
                                                                                        _ = v210
                                                                                    } else {
                                                                                        v212 := GenS9{F0: 255}
                                                                                        v213 := GenS20{F0: (v209 / ((v209) | 1)), F1: fi43, F2: uint8(v198)}
                                                                                        if (0 == v198) {
                                                                                            var v214 float64 = float64(3.14159)
                                                                                            var v215 uint64 = v2
                                                                                            _ = v214
                                                                                            _ = v215
                                                                                        } else {
                                                                                            switch int(v196) {
                                                                                            case -1:
                                                                                                v208 += int64(v2)
                                                                                            default:
                                                                                                if (!(65535 != 18446744073709551615)) {
                                                                                                    v216 := GenS17{F0: "hello", F1: 255, F2: v209}
                                                                                                    delete(v3, "hello")
                                                                                                    if false {
                                                                                                        fmt.Println(v209, v209, v196)
                                                                                                    }
                                                                                                    _ = v216
                                                                                                }
                                                                                            }
                                                                                            var v217 uint = 255
                                                                                            delete(v1, "hello")
                                                                                            _ = v217
                                                                                        }
                                                                                        _ = v212
                                                                                        _ = v213
                                                                                    }
                                                                                } else {
                                                                                    v209 = -32768
                                                                                    fmt.Println(fi43, v6)
                                                                                    if v199 {
                                                                                        if ((fi43 << ((734) & 63)) != (fi43 | -2147483648)) {
                                                                                            v218 := GenS21{F0: v208}
                                                                                            switch (fi43 * 42) {
                                                                                            case 42:
                                                                                                if v199 {
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                } else {
                                                                                                    if v199 {
                                                                                                        fmt.Println(fi55, v208, v208)
                                                                                                    } else {
                                                                                                        delete(v3, "foo bar")
                                                                                                        if v199 {
                                                                                                            var v219 uint = 2
                                                                                                            _ = v219
                                                                                                        } else {
                                                                                                            v199 = v199
                                                                                                            var v220 float64 = 1.5
                                                                                                            var v221 int64 = v208
                                                                                                            _ = v220
                                                                                                            _ = v221
                                                                                                        }
                                                                                                        fmt.Println(fi55, v209)
                                                                                                        fmt.Println(v209, fi43, fi43)
                                                                                                        v222 := GenS22{F0: (v198 ^ (v198 >> ((344) & 63))), F1: (-1 <= 79), F2: 67}
                                                                                                        _ = v222
                                                                                                    }
                                                                                                }
                                                                                                var v223 int8 = v196
                                                                                                _ = v223
                                                                                            default:
                                                                                                switch -1 {
                                                                                                case -117:
                                                                                                    if v199 {
                                                                                                        v208 -= v208
                                                                                                    }
                                                                                                case 0:
                                                                                                    var v224 bool = ((0 < 255) || ((0.0 + 1e10) != ((-2.25 + (1.5 * float64((-2.25 - (0.0 - (((-2.25 / 1.0) - -2.25) + 0.1)))))) * 3.14159)))
                                                                                                    _ = v224
                                                                                                }
                                                                                            }
                                                                                            _ = v218
                                                                                        } else {
                                                                                            var v225 float32 = -2.25
                                                                                            v226 := v3[("a" + "foo bar")]
                                                                                            _ = v225
                                                                                            _ = v226
                                                                                        }
                                                                                        v198 -= 2
                                                                                        var v227 uint32 = uint32(v198)
                                                                                        _ = v227
                                                                                    } else {
                                                                                        var v228 int64 = 9223372036854775807
                                                                                        fmt.Println(v208, v2, fi55)
                                                                                        v229 := GenS23{F0: fi55, F1: 1}
                                                                                        _ = v228
                                                                                        _ = v229
                                                                                    }
                                                                                }
                                                                            }
                                                                            if len(v4) > 0 {
                                                                                v4[0] = uint32(v6)
                                                                            }
                                                                            _ = v209
                                                                        } else {
                                                                            v196 = 94
                                                                            if true {
                                                                                v208 -= (v208 >> ((2) & 63))
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                            } else {
                                                                                v230 := GenS24{F0: 174, F1: 42, F2: ""}
                                                                                v199 = ((v198 + v198) >= v6)
                                                                                _ = v230
                                                                            }
                                                                            v231 := GenS25{F0: (v198 / ((65535) | 1))}
                                                                            v232 := GenS26{F0: v199}
                                                                            _ = v231
                                                                            _ = v232
                                                                        }
                                                                        _ = v208
                                                                    } else {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = 955
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            _ = v205
                                                        } else {
                                                            if true {
                                                                var v233 bool = ((v2 & (^(v2))) == v2)
                                                                fmt.Println(v199)
                                                                var v234 int64 = 1
                                                                _ = v233
                                                                _ = v234
                                                            } else {
                                                                var v235 uint = 804
                                                                var v236 int16 = -1
                                                                if len(v4) > 0 {
                                                                    v4[0] = 65535
                                                                }
                                                                _ = v235
                                                                _ = v236
                                                            }
                                                            v6 = (v198 * 252)
                                                        }
                                                    }
                                                    v6 = (v6 % (((^(v198))) | 1))
                                                    var v237 bool = v199
                                                    _ = v237
                                                }
                                                v2 = 2
                                                fmt.Println(v6)
                                                fmt.Println(v198)
                                                v2 -= v2
                                                _ = v196
                                                _ = v197
                                                _ = v198
                                                _ = v199
                                            } else {
                                                if (!true) {
                                                    var v238 int32 = -1
                                                    switch ("a" + "") {
                                                    case "\n\t":
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                    case "":
                                                        if (v238 > (v238 >> ((18446744073709551615) & 63))) {
                                                            v239 := GenS24{F0: 186, F1: 65535, F2: ("foo bar" + "\n\t")}
                                                            _ = v239
                                                        } else {
                                                            v240 := GenS15{F0: false, F1: -1}
                                                            _ = v240
                                                        }
                                                    }
                                                    _ = v238
                                                } else {
                                                    v2 -= v2
                                                }
                                                v1[(("\n\t" + "\n\t") + ("a" + ("" + "")))] = 559
                                                var v241 float64 = 0.0
                                                _ = v241
                                            }
                                        }
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                default:
                                    var v242 int = fi43
                                    _ = v242
                                }
                                fmt.Println(v2, fi55, v6)
                                _ = fi55
                            }
                            var v243 bool = true
                            switch fi43 {
                            case 519:
                                v2 = v2
                            case -526:
                                if ("foo bar" <= ((("a" + ("foo bar" + "\n\t")) + "foo bar") + "hello")) {
                                    var v244 uint64 = (v2 << ((18446744073709551615) & 63))
                                    if v243 {
                                        if v243 {
                                            for fi245 := 0; fi245 < 6; fi245++ {
                                                v246 := GenS27{F0: 9223372036854775807, F1: "a"}
                                                var v247 int = (fi43 ^ 32767)
                                                _ = fi245
                                                _ = v246
                                                _ = v247
                                            }
                                        } else {
                                            if (293 != 453) {
                                                v243 = true
                                                v244 = (v2 << ((949) & 63))
                                                _ = len(v5)
                                            }
                                            var v248 uint8 = 54
                                            _ = v248
                                        }
                                        var v249 uint32 = 42
                                        v243 = (v243 || (1000 <= -654))
                                        _ = v249
                                    } else {
                                        for fi250 := 0; fi250 < 3; fi250++ {
                                            if true {
                                                var v251 int8 = 0
                                                _ = v251
                                            } else {
                                                switch 2147483647 {
                                                case 342:
                                                    v252 := GenS23{F0: (-(fi43)), F1: 1}
                                                    _ = v252
                                                case 42:
                                                    var v253 uint16 = v6
                                                    _ = v253
                                                case 4294967295:
                                                    var v254 int8 = 127
                                                    _ = v254
                                                default:
                                                    v2 = 18446744073709551615
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    var v255 bool = true
                                                    _ = v255
                                                default:
                                                    _ = len(v5)
                                                }
                                            }
                                            _ = fi250
                                        }
                                        v4 = append(v4, 500)
                                        switch (fi43 << ((65535) & 63)) {
                                        case 0:
                                            switch "a" {
                                            case "":
                                                if v243 {
                                                    for fi256 := 0; fi256 < 13; fi256++ {
                                                        v244 = 0
                                                        var v257 uint32 = 2147483647
                                                        var v258 int = fi43
                                                        _ = fi256
                                                        _ = v257
                                                        _ = v258
                                                    }
                                                    var v259 float32 = 0.0
                                                    _ = v259
                                                }
                                            }
                                        case -9223372036854775808:
                                            v244 -= v244
                                        default:
                                            for fi260 := 0; fi260 < 10; fi260++ {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v244 = (v244 - v244)
                                                _ = fi260
                                            }
                                        }
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        if (1 != 4294967295) {
                                            if (!v243) {
                                                switch ("foo bar" + "a") {
                                                case "":
                                                    fmt.Println(v243)
                                                case "hello":
                                                    for fi261 := 0; fi261 < 20; fi261++ {
                                                        v262 := GenS1{F0: int8(fi43)}
                                                        v263 := GenS24{F0: 128, F1: 18446744073709551615, F2: ("" + "\n\t")}
                                                        _ = fi261
                                                        _ = v262
                                                        _ = v263
                                                    }
                                                }
                                                switch fi43 {
                                                case -2147483648:
                                                    v244 = v2
                                                case -452:
                                                    if (18446744073709551615 == 255) {
                                                        v244 = 25
                                                        var v264 uint32 = 2147483647
                                                        _ = v264
                                                    } else {
                                                        delete(v3, "foo bar")
                                                        if (69 < 66) {
                                                            v6 = v6
                                                            v244 = (v244 ^ (v244 ^ v2))
                                                        }
                                                    }
                                                default:
                                                    switch (fi43 * 4294967295) {
                                                    case -9223372036854775808:
                                                        if true {
                                                            var v265 uint16 = (^(v6))
                                                            _ = v265
                                                        }
                                                    case 179:
                                                        var v266 string = "\n\t"
                                                        _ = v266
                                                    default:
                                                        switch 32768 {
                                                        case 128:
                                                            for fi267 := 0; fi267 < 6; fi267++ {
                                                                fmt.Println(fi267, v6)
                                                                _ = fi267
                                                            }
                                                        case 32768:
                                                            for fi268 := 0; fi268 < 5; fi268++ {
                                                                v243 = v243
                                                                v244 = 222
                                                                v269 := GenS28{F0: 2147483648, F1: 1}
                                                                if v243 {
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    fmt.Println(fi268, v244)
                                                                } else {
                                                                    var v270 uint64 = 4294967295
                                                                    v6 -= v6
                                                                    v270 -= 2
                                                                    _ = v270
                                                                }
                                                                var v271 int64 = 723
                                                                _ = fi268
                                                                _ = v269
                                                                _ = v271
                                                            }
                                                        case 132:
                                                            switch (("\n\t" + (("foo bar" + ("" + "")) + "a")) + "a") {
                                                            case "":
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            case "hello":
                                                                var v272 int8 = 0
                                                                _ = v272
                                                            case "\n\t":
                                                                fmt.Println(v244, fi43, v243)
                                                            }
                                                        default:
                                                            v273 := GenS29{F0: -32768, F1: 65535}
                                                            _ = v273
                                                        }
                                                    }
                                                }
                                                var v274 int32 = -2147483648
                                                _ = v274
                                            } else {
                                                v243 = v243
                                                var v275 bool = v243
                                                if len(v5) > 0 {
                                                    v5[0] = fi43
                                                }
                                                _ = v275
                                            }
                                            if ((v6 >> ((uint(v6)) & 63)) > v6) {
                                                v243 = false
                                                for fi276 := 0; fi276 < 16; fi276++ {
                                                    var v277 int = 128
                                                    var v278 bool = (v243 && v243)
                                                    _ = fi276
                                                    _ = v277
                                                    _ = v278
                                                }
                                            } else {
                                                v279 := GenS30{F0: "foo bar", F1: 0, F2: 680}
                                                if ((v2 / ((v2) | 1)) != 0) {
                                                    v1[("foo bar" + ("foo bar" + "foo bar"))] = 1
                                                    v2 += (^(v244))
                                                    fmt.Println(v244, v243, v6)
                                                }
                                                var v280 uint64 = (v2 - v244)
                                                _ = v279
                                                _ = v280
                                            }
                                            v281 := GenS31{F0: 9223372036854775807, F1: -1}
                                            _ = v281
                                        } else {
                                            var v282 string = ""
                                            fmt.Println(fi43)
                                            _ = v282
                                        }
                                    default:
                                        v283 := GenS27{F0: fi43, F1: "hello"}
                                        _ = v283
                                    }
                                    _ = v244
                                } else {
                                    v6 = 32767
                                }
                            case 613:
                                if v243 {
                                    var v284 string = "foo bar"
                                    if (!((float32((3.14159 + ((3.14159 * 0.0) * 1.5))) / 0.1) >= 0.0)) {
                                        var v285 int = (fi43 >> ((2) & 63))
                                        v285 -= (fi43 << ((4294967295) & 63))
                                        if (true || (v284 > "hello")) {
                                            if false {
                                                v285 = (v285 >> ((2) & 63))
                                                for fi286 := 0; fi286 < 20; fi286++ {
                                                    v2 = (v2 % (((v2 & (v2 ^ 2))) | 1))
                                                    switch (fi286 - 2147483647) {
                                                    case 33:
                                                        fmt.Println(v284, v285, v243)
                                                        v5 = append(v5, 32767)
                                                    case -9223372036854775808:
                                                        var v287 bool = false
                                                        _ = v287
                                                    case -1:
                                                        fmt.Println(v284, fi286)
                                                    default:
                                                        var v288 uint64 = (^(v2))
                                                        _ = v288
                                                    }
                                                    v289 := GenS12{F0: 255}
                                                    _ = fi286
                                                    _ = v289
                                                }
                                            }
                                        } else {
                                            if v243 {
                                                for fi290 := 0; fi290 < 6; fi290++ {
                                                    v6 -= 65535
                                                    _ = fi290
                                                }
                                            }
                                        }
                                        _ = v285
                                    } else {
                                        var v291 uint64 = v2
                                        _ = v291
                                    }
                                    v243 = v243
                                    _ = v284
                                }
                            default:
                                for fi292 := 0; fi292 < 2; fi292++ {
                                    fmt.Println(v2, v6)
                                    if v243 {
                                        var v293 uint = 638
                                        if (int8(v6) >= 0) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            v293 += v293
                                            v6 -= v6
                                            var v294 int8 = 42
                                            _ = v294
                                        }
                                        v295 := GenS9{F0: 2}
                                        _ = v293
                                        _ = v295
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    if ((("a" + "") + "foo bar") >= (("\n\t" + ("foo bar" + "foo bar")) + "hello")) {
                                        if (-32768 >= 716) {
                                            var v296 string = ("foo bar" + ("foo bar" + "a"))
                                            var v297 int16 = -32768
                                            var v298 float32 = float32(1.5)
                                            var v299 int32 = -510
                                            _ = v296
                                            _ = v297
                                            _ = v298
                                            _ = v299
                                        }
                                        v6 += v6
                                    } else {
                                        var v300 int16 = -1
                                        if false {
                                            if v243 {
                                                var v301 uint32 = 65535
                                                v300 = 42
                                                _ = v301
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                        }
                                        v302 := GenS5{F0: 176}
                                        if v243 {
                                            v300 = -32768
                                            fmt.Println(fi292, v300)
                                        }
                                        _ = v300
                                        _ = v302
                                    }
                                    v243 = (!(!(2147483647 <= 2147483647)))
                                    _ = fi292
                                }
                            }
                            _ = v243
                        } else {
                            fmt.Println(v6, v6)
                            v6 += v6
                            if (0 == 1) {
                                v6 = 2
                                v303 := GenS32{F0: 400, F1: 2147483647, F2: ""}
                                for fi304 := 0; fi304 < 2; fi304++ {
                                    v305 := GenS33{F0: 42, F1: 32767, F2: "a"}
                                    _ = fi304
                                    _ = v305
                                }
                                _ = v303
                            } else {
                                fmt.Println(fi43, fi43)
                                var v306 uint8 = 128
                                var v307 int16 = -32768
                                _ = v306
                                _ = v307
                            }
                        }
                        if len(v5) > 0 {
                            v5[0] = (fi43 + fi43)
                        }
                    }
                }
                _ = len(v4)
                _ = fi43
            }
        }
        _ = v41
    }
    v2 = 252
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
