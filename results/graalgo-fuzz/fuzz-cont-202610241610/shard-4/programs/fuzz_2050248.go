package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int64; F2 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint16; F1 int32; F2 int16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint32; F1 int16 }

func main() {
    v1 := []uint32{}
    var v2 int8 = -128
    v3 := make(map[int]uint32)
    var v4 int8 = (v2 << ((32) & 63))
    delete(v3, 42)
    v5 := GenS1{F0: ("hello" + "\n\t"), F1: 42}
    v6 := GenS2{F0: 65535, F1: (42 <= 1000)}
    for fi7 := 0; fi7 < 8; fi7++ {
        if true {
            if len(v1) > 0 {
                v1[0] = 695
            }
            for fi8 := 0; fi8 < 3; fi8++ {
                v2 += v2
                v2 += v4
                fmt.Println(v2, fi7, fi8)
                if (2 < 486) {
                    if (671 != 42) {
                        fmt.Println(v4, fi7)
                        var v9 float64 = float64(0.0)
                        _ = v9
                    }
                } else {
                    v10 := GenS3{F0: 1, F1: 2147483647}
                    var v11 int32 = -2147483648
                    _ = v10
                    _ = v11
                }
                if true {
                    switch "a" {
                    case "":
                        v2 += v4
                    case "\n\t":
                        switch 127 {
                        case 128:
                            var v12 uint16 = 328
                            _ = v12
                        case 32768:
                            v4 += v2
                            switch fi7 {
                            case 946:
                                if true {
                                    if false {
                                        v2 = 127
                                        switch -485 {
                                        case 1:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            switch fi7 {
                                            case 42:
                                                v2 = (-(v2))
                                            case -789:
                                                if false {
                                                    fmt.Println(v4, fi7)
                                                    if ((-(v4)) < (^(v4))) {
                                                        var v13 string = (("foo bar" + (((("\n\t" + "hello") + "hello") + (("foo bar" + "\n\t") + "\n\t")) + "")) + "\n\t")
                                                        switch fi7 {
                                                        case 9223372036854775807:
                                                            if true {
                                                                if true {
                                                                    fmt.Println(v2, v4, v13)
                                                                    var v14 int = fi7
                                                                    _ = v14
                                                                }
                                                                var v15 uint = 999
                                                                if (v13 != (("" + v13) + v13)) {
                                                                    var v16 int64 = 0
                                                                    v16 = -925
                                                                    fmt.Println(fi8)
                                                                    _ = v16
                                                                }
                                                                _ = v15
                                                            }
                                                        case -9223372036854775808:
                                                            fmt.Println(fi7)
                                                            if true {
                                                                var v17 int = 4294967295
                                                                var v18 string = "hello"
                                                                _ = v17
                                                                _ = v18
                                                            } else {
                                                                var v19 uint32 = 0
                                                                switch "\n\t" {
                                                                case "":
                                                                    if true {
                                                                        switch (fi7 ^ (fi8 & fi7)) {
                                                                        case 2147483647:
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (v19 & v19)
                                                                            }
                                                                        case 32767:
                                                                            v3[(-(fi7))] = v19
                                                                        case 128:
                                                                            switch v13 {
                                                                            case "":
                                                                                v19 = v19
                                                                                fmt.Println(fi7, v13)
                                                                            case "hello":
                                                                                v20 := GenS1{F0: (v13 + v13), F1: 452}
                                                                                _ = v20
                                                                            default:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 42
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if (!false) {
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            fmt.Println(v13, v13, fi7)
                                                                        }
                                                                        v4 += v2
                                                                        _ = len(v1)
                                                                    }
                                                                case "foo bar":
                                                                    v21 := GenS1{F0: v13, F1: 2}
                                                                    _ = v21
                                                                case "a":
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        fmt.Println(v2, v19, v2)
                                                                    }
                                                                }
                                                                _ = v19
                                                            }
                                                        default:
                                                            if (!(fi8 <= fi8)) {
                                                                fmt.Println(fi7, fi8, v4)
                                                                v22 := GenS4{F0: 871}
                                                                var v23 uint = 826
                                                                _ = v22
                                                                _ = v23
                                                            } else {
                                                                switch fi7 {
                                                                case 600:
                                                                    v2 = v4
                                                                }
                                                            }
                                                        }
                                                        var v24 uint64 = 42
                                                        _ = v13
                                                        _ = v24
                                                    } else {
                                                        var v25 uint8 = 1
                                                        v26 := GenS5{F0: 0, F1: 1, F2: 1}
                                                        var v27 uint8 = (v25 / ((1) | 1))
                                                        v28 := GenS2{F0: 252, F1: false}
                                                        v1 = append(v1, 4294967295)
                                                        _ = v25
                                                        _ = v26
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                    v29 := GenS3{F0: 581, F1: 0}
                                                    _ = v29
                                                } else {
                                                    var v30 uint8 = 2
                                                    v2 = 66
                                                    fmt.Println(v30, v2, v2)
                                                    _ = v30
                                                }
                                            }
                                        }
                                        if (false && true) {
                                            var v31 float32 = 3.14159
                                            v2 -= (v2 << ((0) & 63))
                                            if (691 >= fi7) {
                                                var v32 uint64 = 0
                                                _ = v32
                                            }
                                            _ = v31
                                        }
                                    } else {
                                        v2 -= (v4 + 0)
                                        v33 := v3[2147483647]
                                        _ = v33
                                    }
                                    var v34 int = fi7
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(v4)
                                    default:
                                        var v35 string = ("\n\t" + ("\n\t" + "hello"))
                                        _ = v35
                                    }
                                    _ = v34
                                } else {
                                    v36 := GenS4{F0: 1}
                                    v37 := GenS2{F0: 749, F1: (-255 != -524)}
                                    if ("foo bar" > "foo bar") {
                                        v3[(fi7 / ((-895) | 1))] = 42
                                        if (v4 < 1) {
                                            if true {
                                                var v38 bool = false
                                                _ = v38
                                            } else {
                                                var v39 int64 = -42
                                                _ = v39
                                            }
                                            var v40 int8 = v2
                                            _ = v40
                                        }
                                        v4 = (v4 / (((^(v2))) | 1))
                                    } else {
                                        v41 := GenS3{F0: 0, F1: -1}
                                        var v42 uint32 = 82
                                        _ = v41
                                        _ = v42
                                    }
                                    var v43 uint8 = 42
                                    _ = v36
                                    _ = v37
                                    _ = v43
                                }
                            default:
                                var v44 string = "\n\t"
                                _ = v44
                            }
                        case 9223372036854775807:
                            switch "a" {
                            case "foo bar":
                                var v45 int8 = (v2 ^ v4)
                                _ = v45
                            default:
                                var v46 float32 = 0.1
                                _ = v46
                            }
                        }
                    case "a":
                        switch "\n\t" {
                        case "":
                            if ((42 <= 1) && (!(0 == 2))) {
                                var v47 int32 = -119
                                _ = v47
                            }
                            var v48 float32 = float32(((-2.25 / 1.0) - 3.14159))
                            switch 128 {
                            case 67:
                                _ = len(v3)
                            case 32767:
                                switch "a" {
                                case "":
                                    if (!((v48 < (v48 + -2.25)) && true)) {
                                        var v49 uint64 = 4294967295
                                        var v50 float32 = v48
                                        var v51 int32 = 1
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    }
                                case "foo bar":
                                    v52 := GenS5{F0: 1, F1: -9223372036854775808, F2: 0}
                                    _ = v52
                                default:
                                    var v53 int64 = 9223372036854775807
                                    _ = v53
                                }
                            }
                            var v54 uint = 514
                            v55 := GenS6{F0: 128}
                            _ = len(v3)
                            _ = v48
                            _ = v54
                            _ = v55
                        default:
                            v56 := GenS7{F0: 32767, F1: 326, F2: -32768}
                            _ = v56
                        }
                    default:
                        switch ("foo bar" + "a") {
                        case "foo bar":
                            switch "hello" {
                            case "\n\t":
                                if false {
                                    v2 -= int8(fi7)
                                }
                            default:
                                v57 := GenS2{F0: 0, F1: true}
                                if true {
                                    v3[9223372036854775807] = 2147483647
                                    v58 := GenS7{F0: 255, F1: 0, F2: -32768}
                                    _ = v58
                                }
                                v1 = append(v1, 65535)
                                _ = v57
                            }
                        }
                        v2 -= v2
                        if ((1 > 1) || true) {
                            fmt.Println(v2, v4, fi7)
                        }
                    }
                    v59 := GenS4{F0: 519}
                    _ = v59
                } else {
                    if ((true || false) || (((fi8 != (^(fi7))) && ((!false) || (!(42 > 2147483648)))) && true)) {
                        if true {
                            v60 := GenS8{F0: 668}
                            v61 := GenS3{F0: 1, F1: 1}
                            if true {
                                var v62 float32 = (1.5 + -2.25)
                                var v63 int64 = 982
                                fmt.Println(fi8, v62)
                                _ = v62
                                _ = v63
                            }
                            _ = v60
                            _ = v61
                        } else {
                            delete(v3, 917)
                            switch (fi7 & 32768) {
                            case 9223372036854775807:
                                var v64 uint8 = 128
                                _ = v64
                            }
                        }
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v2 -= v2
                        if false {
                            v4 += (v4 ^ (v2 & (v4 ^ -45)))
                            if true {
                                fmt.Println(v4)
                                v3[int(v4)] = 869
                                var v65 uint8 = 255
                                _ = v65
                            } else {
                                _ = len(v3)
                                var v66 uint8 = 128
                                _ = v66
                            }
                            switch 1 {
                            case 156:
                                var v67 bool = (466 > -1)
                                _ = v67
                            case 32767:
                                switch "a" {
                                case "":
                                    if false {
                                        var v68 int32 = 1
                                        fmt.Println(fi8)
                                        if true {
                                            if (0 != 127) {
                                                if (!((v4 - 127) == v4)) {
                                                    var v69 uint16 = 1
                                                    _ = v69
                                                } else {
                                                    v4 += v4
                                                    v68 -= -1
                                                    var v70 uint64 = 0
                                                    _ = v70
                                                }
                                                var v71 int16 = 32767
                                                _ = v71
                                            }
                                        }
                                        _ = v68
                                    } else {
                                        var v72 uint32 = 2
                                        _ = v72
                                    }
                                case "\n\t":
                                    if (!true) {
                                        if ((-(v4)) <= v4) {
                                            var v73 uint16 = 255
                                            _ = v73
                                        } else {
                                            var v74 int32 = 1
                                            var v75 int64 = 0
                                            _ = v74
                                            _ = v75
                                        }
                                        var v76 int32 = -260
                                        _ = v76
                                    } else {
                                        var v77 uint = 42
                                        v78 := GenS1{F0: (("\n\t" + "foo bar") + "a"), F1: 18446744073709551615}
                                        if true {
                                            switch fi7 {
                                            case 32767:
                                                var v79 float32 = 1.5
                                                _ = v79
                                            }
                                            switch "a" {
                                            case "a":
                                                v80 := GenS5{F0: 344, F1: 9223372036854775807, F2: 2}
                                                _ = v80
                                            }
                                            if false {
                                                v2 = (^(v2))
                                                var v81 uint = v77
                                                v77 -= v77
                                                _ = v81
                                            } else {
                                                _ = len(v1)
                                                if false {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            }
                                        }
                                        _ = v77
                                        _ = v78
                                    }
                                case "a":
                                    v82 := GenS9{F0: 0, F1: -691, F2: 927}
                                    _ = v82
                                }
                            case -9223372036854775808:
                                if true {
                                    v83 := GenS8{F0: 32767}
                                    _ = v83
                                } else {
                                    var v84 uint8 = 190
                                    var v85 int8 = -104
                                    _ = v84
                                    _ = v85
                                }
                            default:
                                var v86 int8 = v4
                                _ = v86
                            }
                        } else {
                            switch "" {
                            case "\n\t":
                                var v87 bool = true
                                _ = v87
                            default:
                                v4 = v2
                                var v88 uint = 1
                                _ = v88
                            }
                            fmt.Println(v4)
                        }
                    }
                    v89 := GenS4{F0: 255}
                    v4 += 127
                    _ = v89
                }
                _ = fi8
            }
            var v90 int8 = (v2 * (v4 & (v2 ^ v4)))
            _ = v90
        } else {
            var v91 float32 = -2.25
            _ = v91
        }
        switch ((("hello" + "hello") + "") + "hello") {
        case "\n\t":
            v4 -= v2
        case "a":
            _ = len(v3)
        }
        v4 = v4
        _ = fi7
    }
    v92 := GenS10{F0: -12}
    if false {
        switch -9223372036854775808 {
        case -2147483648:
            v2 = -128
        case 32768:
            switch "hello" {
            case "a":
                v93 := GenS11{F0: 2147483648, F1: -32768}
                _ = v93
            }
        case 9223372036854775807:
            var v94 bool = false
            _ = v94
        }
        switch "" {
        case "a":
            var v95 int8 = v4
            _ = v95
        default:
            var v96 float32 = 1.5
            _ = v96
        }
        fmt.Println(v4, v2)
    } else {
        var v97 int32 = 42
        _ = v97
    }
    switch 1 {
    case 562:
        v2 = v4
    case -269:
        var v98 uint64 = 0
        _ = v98
    case 9223372036854775807:
        v4 += v4
    }
    if true {
        var v99 uint16 = 65535
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        switch "hello" {
        case "foo bar":
            fmt.Println(v4, v4)
        case "hello":
            _ = len(v1)
        }
        _ = v99
    }
    var v100 int16 = 0
    v2 = v4
    if (1 == 128) {
        v2 = 1
        var v101 float64 = (1.5 + (float64(1.5) * ((3.14159 + 1.5) / (1.5 - -2.25))))
        _ = v101
    } else {
        switch "" {
        case "hello":
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        }
        fmt.Println(v100, v4)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v92
}
