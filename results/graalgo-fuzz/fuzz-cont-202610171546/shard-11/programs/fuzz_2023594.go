package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint16; F1 int8; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint; F2 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS10 struct { F0 string; F1 int }
type GenS11 struct { F0 uint64; F1 uint8 }
type GenS12 struct { F0 int64; F1 uint8 }
type GenS13 struct { F0 string; F1 int8; F2 string }
type GenS14 struct { F0 bool; F1 string }

func main() {
    var v1 float32 = ((0.1 * -2.25) / (0.0 + (3.14159 + 3.14159)))
    var v2 uint64 = 570
    var v3 int = -603
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    var v6 uint8 = 2
    v7 := GenS1{F0: -361}
    fmt.Println(v2, v6)
    var v8 int8 = 0
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS2{F0: 255, F1: (fi9 * v3)}
        fmt.Println(v1)
        var v11 int = v3
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 65535, F1: (^(v3))}
    v13 := []int{(^(v3)), (v3 << ((494) & 63))}
    if true {
        if (!(110 <= v2)) {
            fmt.Println(v1, v2)
            if true {
                var v14 string = "hello"
                _ = v14
            }
        }
    } else {
        fmt.Println(v8, v6)
    }
    v15 := v4[(v3 >> ((1) & 63))]
    switch v3 {
    case 4294967295:
        v16 := GenS3{F0: -128, F1: 255}
        _ = v16
    case 841:
        var v17 uint16 = 0
        _ = v17
    case 1:
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(fi18, fi18)
            if v15 {
                if (v1 == float32((-2.25 * 0.1))) {
                    v19 := GenS2{F0: 0, F1: v3}
                    for fi20 := 0; fi20 < 12; fi20++ {
                        switch 583 {
                        case -90:
                            var v21 float32 = 3.14159
                            _ = v21
                        case -30:
                            var v22 int8 = v8
                            _ = v22
                        case 128:
                            var v23 uint16 = 514
                            _ = v23
                        default:
                            switch (v3 / (((fi18 - (fi18 & 986))) | 1)) {
                            case -17:
                                if true {
                                    if len(v13) > 0 {
                                        v13[0] = 42
                                    }
                                    v24 := GenS1{F0: 1000}
                                    _ = v24
                                }
                            default:
                                if (0.1 > (v1 * v1)) {
                                    if (((-2147483648 <= 2147483647) && v15) && v15) {
                                        v25 := GenS4{F0: 65535, F1: v8, F2: 9223372036854775807}
                                        if false {
                                            switch (fi20 - fi20) {
                                            case 0:
                                                v26 := GenS5{F0: 255, F1: 18446744073709551615, F2: fi20}
                                                _ = v26
                                            case 554:
                                                if (1 == 0) {
                                                    _ = len(v4)
                                                    if v15 {
                                                        delete(v4, fi20)
                                                        if (550 > 42) {
                                                            v8 += (v8 - v8)
                                                            v6 -= v6
                                                        }
                                                        v4[v3] = v15
                                                    } else {
                                                        if v15 {
                                                            var v27 uint64 = v2
                                                            v27 = v27
                                                            var v28 bool = false
                                                            switch -9223372036854775808 {
                                                            case -1:
                                                                v3 += 32768
                                                            default:
                                                                fmt.Println(v6)
                                                            }
                                                            v15 = true
                                                            _ = v27
                                                            _ = v28
                                                        } else {
                                                            var v29 uint32 = 412
                                                            _ = v29
                                                        }
                                                    }
                                                    v6 += v6
                                                }
                                            }
                                        } else {
                                            if true {
                                                if v15 {
                                                    v15 = ((2147483647 < 468) || true)
                                                }
                                                if v15 {
                                                    v4[fi18] = true
                                                    var v30 float64 = -2.25
                                                    _ = v30
                                                }
                                                switch "a" {
                                                case "foo bar":
                                                    v1 += ((float32(1e10) - v1) * 0.1)
                                                case "hello":
                                                    var v31 uint16 = 1
                                                    _ = v31
                                                default:
                                                    switch (fi18 >> ((uint(v6)) & 63)) {
                                                    case 42:
                                                        v32 := GenS6{F0: ((v8 << ((2) & 63)) >= 109)}
                                                        _ = v32
                                                    case -9223372036854775808:
                                                        _ = len(v5)
                                                    default:
                                                        delete(v5, 65535)
                                                    }
                                                }
                                                var v33 uint8 = (v6 << ((18446744073709551615) & 63))
                                                _ = v33
                                            } else {
                                                v34 := GenS5{F0: 65535, F1: 42, F2: (^(fi18))}
                                                v35 := GenS4{F0: 1, F1: v8, F2: 9223372036854775807}
                                                var v36 int32 = 42
                                                v6 += (v6 - (v6 ^ 255))
                                                var v37 uint16 = 978
                                                _ = v34
                                                _ = v35
                                                _ = v36
                                                _ = v37
                                            }
                                            v2 += (v2 & v2)
                                        }
                                        _ = v25
                                    } else {
                                        var v38 int = -856
                                        _ = v38
                                    }
                                    switch (fi18 | fi20) {
                                    case -9223372036854775808:
                                        if v15 {
                                            v1 -= v1
                                            v8 = v8
                                            v39 := GenS7{F0: int16(v2), F1: -9223372036854775808}
                                            _ = v39
                                        } else {
                                            v8 = v8
                                            v6 -= uint8(v2)
                                        }
                                    }
                                    if len(v13) > 0 {
                                        v13[0] = (fi20 << ((255) & 63))
                                    }
                                } else {
                                    if v15 {
                                        if len(v13) > 0 {
                                            v13[0] = v3
                                        }
                                        switch (fi20 % ((32767) | 1)) {
                                        case -350:
                                            if ("" <= "") {
                                                if false {
                                                    var v40 uint32 = 903
                                                    var v41 int = fi18
                                                    if len(v13) >= 2 {
                                                        v13 = v13[1:len(v13)]
                                                    }
                                                    _ = v40
                                                    _ = v41
                                                }
                                                switch "" {
                                                case "\n\t":
                                                    switch "a" {
                                                    case "\n\t":
                                                        switch v3 {
                                                        case 128:
                                                            v1 -= v1
                                                        case 0:
                                                            v8 = v8
                                                        default:
                                                            if (!(((v1 - v1) - (3.14159 / v1)) < 0.0)) {
                                                                var v42 bool = true
                                                                var v43 int8 = (v8 ^ (v8 % ((-72) | 1)))
                                                                switch "a" {
                                                                case "hello":
                                                                    fmt.Println(v3)
                                                                case "foo bar":
                                                                    v8 -= 71
                                                                }
                                                                _ = v42
                                                                _ = v43
                                                            }
                                                        }
                                                    default:
                                                        v1 = (1.5 + (float32(1e10) + (v1 * v1)))
                                                    }
                                                default:
                                                    var v44 int = 634
                                                    _ = v44
                                                }
                                            }
                                        case 343:
                                            _ = len(v4)
                                        case 42:
                                            if false {
                                                if v15 {
                                                    v1 -= (v1 - float32(0.0))
                                                    if (1 > 4294967295) {
                                                        if (!v15) {
                                                            v45 := GenS4{F0: 255, F1: v8, F2: 56}
                                                            _ = v45
                                                        }
                                                    } else {
                                                        if ("a" != "a") {
                                                            var v46 uint = 988
                                                            _ = v46
                                                        } else {
                                                            v47 := GenS3{F0: 0, F1: 70}
                                                            v8 -= (v8 << ((0) & 63))
                                                            if v15 {
                                                                var v48 string = "\n\t"
                                                                if v15 {
                                                                    switch ((v48 + v48) + "") {
                                                                    case "foo bar":
                                                                        _ = len(v13)
                                                                    case "hello":
                                                                        v13 = append(v13, (v3 | (v3 / ((fi20) | 1))))
                                                                    }
                                                                    var v49 int = fi20
                                                                    if (v48 <= (v48 + v48)) {
                                                                        if v15 {
                                                                            switch ("" + (((v48 + "") + "") + "hello")) {
                                                                            case "foo bar":
                                                                                v6 = v6
                                                                            case "hello":
                                                                                var v50 int8 = (v8 >> ((2) & 63))
                                                                                _ = v50
                                                                            }
                                                                            if (42 > 902) {
                                                                                if (((!(!v15)) || v15) && (false && (-9223372036854775808 != -9223372036854775808))) {
                                                                                    v51 := GenS8{F0: -1, F1: 988}
                                                                                    v3 -= 128
                                                                                    v2 += v2
                                                                                    v2 -= 4294967295
                                                                                    switch v49 {
                                                                                    case 1:
                                                                                        switch v3 {
                                                                                        case -9223372036854775808:
                                                                                            fmt.Println(v49, v2, v1)
                                                                                        case 4294967295:
                                                                                            var v52 int = (fi18 * (^(v3)))
                                                                                            _ = v52
                                                                                        case 32767:
                                                                                            if v15 {
                                                                                                if true {
                                                                                                    v4[-282] = (3.14159 < 1e10)
                                                                                                } else {
                                                                                                    delete(v5, 38)
                                                                                                    _ = len(v13)
                                                                                                    var v53 string = v48
                                                                                                    _ = v53
                                                                                                }
                                                                                                delete(v4, 1)
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        switch v48 {
                                                                                        case "\n\t":
                                                                                            v54 := v5[119]
                                                                                            _ = v54
                                                                                        }
                                                                                    }
                                                                                    var v55 int = fi18
                                                                                    _ = v51
                                                                                    _ = v55
                                                                                }
                                                                            } else {
                                                                                switch fi20 {
                                                                                case 128:
                                                                                    if v15 {
                                                                                        v3 += v49
                                                                                        fmt.Println(v48)
                                                                                        if false {
                                                                                            var v56 int8 = -128
                                                                                            var v57 int8 = v56
                                                                                            _ = v56
                                                                                            _ = v57
                                                                                        } else {
                                                                                            v6 -= (v6 | (v6 % (((^(v6))) | 1)))
                                                                                            if v15 {
                                                                                                var v58 float64 = float64(float32(1.5))
                                                                                                _ = v58
                                                                                            } else {
                                                                                                switch fi18 {
                                                                                                case 0:
                                                                                                    var v59 uint = 255
                                                                                                    _ = v59
                                                                                                case 128:
                                                                                                    v15 = v15
                                                                                                case 42:
                                                                                                    var v60 int = (fi18 % ((v49) | 1))
                                                                                                    switch (v48 + (v48 + "\n\t")) {
                                                                                                    case "":
                                                                                                        v49 = fi18
                                                                                                    default:
                                                                                                        if v15 {
                                                                                                            if true {
                                                                                                                fmt.Println(v48, v15, v15)
                                                                                                                switch 32768 {
                                                                                                                case -764:
                                                                                                                    var v61 int16 = 0
                                                                                                                    _ = v61
                                                                                                                case 128:
                                                                                                                    fmt.Println(v60, v3, v3)
                                                                                                                default:
                                                                                                                    var v62 int16 = 875
                                                                                                                    _ = v62
                                                                                                                }
                                                                                                                var v63 uint = 2
                                                                                                                if len(v13) > 0 {
                                                                                                                    v13[0] = v60
                                                                                                                }
                                                                                                                _ = v63
                                                                                                            } else {
                                                                                                                v3 = (fi20 << ((926) & 63))
                                                                                                                fmt.Println(v6)
                                                                                                            }
                                                                                                            if false {
                                                                                                                v13 = append(v13, -9223372036854775808)
                                                                                                            } else {
                                                                                                                v64 := GenS9{F0: 1, F1: v2, F2: 4294967295}
                                                                                                                _ = v64
                                                                                                            }
                                                                                                            fmt.Println(v49)
                                                                                                        } else {
                                                                                                            v6 += 102
                                                                                                        }
                                                                                                    }
                                                                                                    if (!(!(!v15))) {
                                                                                                        var v65 uint64 = (v2 - v2)
                                                                                                        if (9223372036854775807 <= 719) {
                                                                                                            v66 := GenS5{F0: 255, F1: 2, F2: 1}
                                                                                                            v13 = append(v13, 42)
                                                                                                            _ = v66
                                                                                                        } else {
                                                                                                            v67 := GenS1{F0: 1000}
                                                                                                            v68 := GenS2{F0: uint16(v65), F1: -1}
                                                                                                            _ = v67
                                                                                                            _ = v68
                                                                                                        }
                                                                                                        _ = v65
                                                                                                    } else {
                                                                                                        v6 += (v6 << ((2) & 63))
                                                                                                        v3 -= fi18
                                                                                                        var v69 string = ("foo bar" + (v48 + ("" + v48)))
                                                                                                        _ = v69
                                                                                                    }
                                                                                                    _ = v60
                                                                                                }
                                                                                                var v70 int16 = 372
                                                                                                if (true && ((v6 ^ 42) != v6)) {
                                                                                                    if (!v15) {
                                                                                                        switch -2147483648 {
                                                                                                        case -9223372036854775808:
                                                                                                            var v71 uint32 = 1
                                                                                                            _ = v71
                                                                                                        }
                                                                                                    } else {
                                                                                                        if v15 {
                                                                                                            if true {
                                                                                                                if v15 {
                                                                                                                    v72 := v4[(v49 + -1)]
                                                                                                                    v4[fi18] = (float32(0.1) == (1e10 + (v1 + v1)))
                                                                                                                    var v73 uint32 = 31
                                                                                                                    _ = v72
                                                                                                                    _ = v73
                                                                                                                }
                                                                                                                var v74 bool = (!v15)
                                                                                                                v13 = append(v13, v3)
                                                                                                                v74 = (!((true && v15) || ((v8 << ((255) & 63)) <= (-(v8)))))
                                                                                                                if (!v15) {
                                                                                                                    v70 = 0
                                                                                                                    fmt.Println(fi20, v74)
                                                                                                                }
                                                                                                                switch v48 {
                                                                                                                case "hello":
                                                                                                                    fmt.Println(v74, v74)
                                                                                                                case "\n\t":
                                                                                                                    var v75 string = (("foo bar" + ("a" + "\n\t")) + "\n\t")
                                                                                                                    _ = v75
                                                                                                                }
                                                                                                                _ = v74
                                                                                                            }
                                                                                                        }
                                                                                                        var v76 uint8 = (v6 << ((2) & 63))
                                                                                                        switch (fi20 - (v49 + v3)) {
                                                                                                        case -2147483648:
                                                                                                            switch int(v70) {
                                                                                                            case -1:
                                                                                                                fmt.Println(v8)
                                                                                                            case 127:
                                                                                                                fmt.Println(v3)
                                                                                                            }
                                                                                                        case -9223372036854775808:
                                                                                                            switch -184 {
                                                                                                            case 32767:
                                                                                                                var v77 float32 = v1
                                                                                                                _ = v77
                                                                                                            case 770:
                                                                                                                v13 = append(v13, (fi18 & (v49 ^ fi18)))
                                                                                                            default:
                                                                                                                var v78 int64 = 9223372036854775807
                                                                                                                _ = v78
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v76
                                                                                                    }
                                                                                                    if true {
                                                                                                        var v79 int8 = v8
                                                                                                        if v15 {
                                                                                                            var v80 uint = 987
                                                                                                            _ = v80
                                                                                                        }
                                                                                                        var v81 uint32 = 42
                                                                                                        if true {
                                                                                                            _ = len(v4)
                                                                                                            switch 954 {
                                                                                                            case 128:
                                                                                                                fmt.Println(v3, v70)
                                                                                                            case 905:
                                                                                                                v79 = v8
                                                                                                            case 9223372036854775807:
                                                                                                                v8 += v79
                                                                                                            }
                                                                                                        } else {
                                                                                                            if (2147483647 > -814) {
                                                                                                                var v82 float32 = (v1 - (0.0 + v1))
                                                                                                                _ = v82
                                                                                                            }
                                                                                                            var v83 int16 = (v70 ^ v70)
                                                                                                            v15 = v15
                                                                                                            _ = v83
                                                                                                        }
                                                                                                        v84 := GenS8{F0: int32(v79), F1: (v81 & v81)}
                                                                                                        v81 = (v81 % ((42) | 1))
                                                                                                        var v85 string = ("\n\t" + "a")
                                                                                                        _ = v79
                                                                                                        _ = v81
                                                                                                        _ = v84
                                                                                                        _ = v85
                                                                                                    } else {
                                                                                                        var v86 int32 = 2147483647
                                                                                                        var v87 uint64 = (v2 << ((4294967295) & 63))
                                                                                                        switch (((v48 + "foo bar") + "foo bar") + "\n\t") {
                                                                                                        case "hello":
                                                                                                            v88 := GenS9{F0: 2147483647, F1: v87, F2: (^(v87))}
                                                                                                            _ = v88
                                                                                                        }
                                                                                                        _ = v86
                                                                                                        _ = v87
                                                                                                    }
                                                                                                }
                                                                                                var v89 int16 = (v70 ^ 42)
                                                                                                _ = v70
                                                                                                _ = v89
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v3 = int(v8)
                                                                                }
                                                                            }
                                                                            var v90 uint32 = 2147483647
                                                                            _ = v90
                                                                        } else {
                                                                            v91 := GenS8{F0: 0, F1: 42}
                                                                            var v92 bool = v15
                                                                            _ = v91
                                                                            _ = v92
                                                                        }
                                                                        v5[869] = (v3 << ((1) & 63))
                                                                    } else {
                                                                        if v15 {
                                                                            v49 -= -102
                                                                            switch v49 {
                                                                            case 127:
                                                                                v3 = (v49 << ((499) & 63))
                                                                            case 32768:
                                                                                if (!(255 == 255)) {
                                                                                    if (!v15) {
                                                                                        if (v2 != v2) {
                                                                                            v93 := GenS10{F0: (v48 + "\n\t"), F1: (^(fi18))}
                                                                                            v8 += (v8 & (v8 + (v8 / ((-11) | 1))))
                                                                                            _ = v93
                                                                                        } else {
                                                                                            if (-361 <= -871) {
                                                                                                var v94 string = "foo bar"
                                                                                                _ = v94
                                                                                            } else {
                                                                                                if v15 {
                                                                                                    _ = len(v13)
                                                                                                    if (32767 != 8) {
                                                                                                        var v95 int32 = -894
                                                                                                        _ = v95
                                                                                                    }
                                                                                                    if (255 != 255) {
                                                                                                        if (!false) {
                                                                                                            if false {
                                                                                                                v3 = (v3 / (((-(v3))) | 1))
                                                                                                                var v96 int = v49
                                                                                                                var v97 float32 = ((v1 / (float32((0.1 + 1e10)) * (float32((0.0 + float64(float32(float64(float32(0.1)))))) + (1e10 / 1e10)))) + ((v1 + -2.25) * v1))
                                                                                                                _ = v96
                                                                                                                _ = v97
                                                                                                            }
                                                                                                            var v98 int64 = -181
                                                                                                            _ = v98
                                                                                                        }
                                                                                                        if v15 {
                                                                                                            switch (v49 + int(v6)) {
                                                                                                            case 32768:
                                                                                                                v8 += v8
                                                                                                            }
                                                                                                        }
                                                                                                        v1 = (1.5 + ((((v1 - ((v1 - v1) + (v1 + (1e10 / v1)))) + v1) / (v1 / float32(1.5))) - (v1 + ((float32((((float64(-2.25) + (0.0 - (1e10 * ((0.0 - float64(v1)) * -2.25)))) + ((1e10 / 0.1) - 3.14159)) / (float64((1e10 / 3.14159)) / float64((3.14159 + (v1 * (3.14159 - -2.25))))))) / 1e10) + float32((0.0 - ((1.5 / ((float64(0.0) - (((1e10 + 1.5) - (1.5 / 3.14159)) * (float64(float32(-2.25)) / 1.5))) + 1e10)) * (1e10 - float64(1.5)))))))))
                                                                                                    }
                                                                                                }
                                                                                                fmt.Println(v48, v49, v15)
                                                                                                v99 := GenS10{F0: v48, F1: 127}
                                                                                                if false {
                                                                                                    fmt.Println(v8, v6)
                                                                                                    v49 = fi18
                                                                                                }
                                                                                                v100 := GenS11{F0: (^(v2)), F1: v6}
                                                                                                _ = v100
                                                                                                _ = v99
                                                                                            }
                                                                                        }
                                                                                        v1 = v1
                                                                                    }
                                                                                    v101 := GenS9{F0: 1, F1: 1, F2: v2}
                                                                                    _ = v101
                                                                                } else {
                                                                                    var v102 bool = v15
                                                                                    if v15 {
                                                                                        var v103 uint64 = (^(v2))
                                                                                        _ = v103
                                                                                    }
                                                                                    v13 = append(v13, (v49 % ((v49) | 1)))
                                                                                    var v104 uint32 = 0
                                                                                    _ = v102
                                                                                    _ = v104
                                                                                }
                                                                            case 0:
                                                                                switch ("" + (v48 + (v48 + "\n\t"))) {
                                                                                case "hello":
                                                                                    if (v15 && ((1 <= (v8 % ((v8) | 1))) && (v8 > v8))) {
                                                                                        switch v48 {
                                                                                        case "hello":
                                                                                            fmt.Println(v1, v1, v8)
                                                                                        case "\n\t":
                                                                                            v105 := GenS8{F0: -378, F1: 65535}
                                                                                            _ = v105
                                                                                        default:
                                                                                            var v106 float64 = 1.5
                                                                                            _ = v106
                                                                                        }
                                                                                        fmt.Println(v15, v3)
                                                                                    }
                                                                                case "foo bar":
                                                                                    fmt.Println(v2)
                                                                                case "":
                                                                                    v107 := GenS6{F0: v15}
                                                                                    _ = v107
                                                                                default:
                                                                                    var v108 float64 = -2.25
                                                                                    _ = v108
                                                                                }
                                                                            }
                                                                            if v15 {
                                                                                if v15 {
                                                                                    v13 = append(v13, (v49 * (v49 / ((fi18) | 1))))
                                                                                    var v109 bool = v15
                                                                                    _ = v109
                                                                                } else {
                                                                                    v15 = v15
                                                                                }
                                                                            }
                                                                        }
                                                                        v110 := GenS11{F0: 18446744073709551615, F1: v6}
                                                                        _ = v110
                                                                    }
                                                                    v49 -= 127
                                                                    _ = v49
                                                                } else {
                                                                    v111 := GenS8{F0: 944, F1: 2}
                                                                    if ((v48 + "a") >= v48) {
                                                                        var v112 uint = 2
                                                                        var v113 uint16 = 0
                                                                        if v15 {
                                                                            if len(v13) >= 2 {
                                                                                v13 = v13[1:len(v13)]
                                                                            }
                                                                            if v15 {
                                                                                switch ("" + "\n\t") {
                                                                                case "\n\t":
                                                                                    v6 += v6
                                                                                default:
                                                                                    if len(v13) >= 2 {
                                                                                        v13 = v13[1:len(v13)]
                                                                                    }
                                                                                }
                                                                                var v114 string = v48
                                                                                _ = v114
                                                                            }
                                                                        }
                                                                        v112 += v112
                                                                        _ = v112
                                                                        _ = v113
                                                                    } else {
                                                                        var v115 uint8 = v6
                                                                        var v116 uint64 = v2
                                                                        _ = v115
                                                                        _ = v116
                                                                    }
                                                                    _ = v111
                                                                }
                                                                var v117 int8 = v8
                                                                _ = v117
                                                                _ = v48
                                                            } else {
                                                                var v118 int8 = (v8 / (((v8 >> ((42) & 63))) | 1))
                                                                switch (("\n\t" + ("\n\t" + "")) + "foo bar") {
                                                                case "hello":
                                                                    switch (("\n\t" + "\n\t") + (((("a" + "a") + ("a" + "")) + (("a" + ("" + "\n\t")) + "")) + "foo bar")) {
                                                                    case "a":
                                                                        v119 := GenS12{F0: 42, F1: v6}
                                                                        _ = v119
                                                                    case "":
                                                                        if v15 {
                                                                            fmt.Println(v118)
                                                                            var v120 uint64 = (^(v2))
                                                                            var v121 uint = 1
                                                                            _ = v120
                                                                            _ = v121
                                                                        } else {
                                                                            var v122 int64 = 9223372036854775807
                                                                            switch -1 {
                                                                            case -9223372036854775808:
                                                                                var v123 float32 = (v1 * 1.5)
                                                                                _ = v123
                                                                            case 0:
                                                                                if true {
                                                                                    var v124 int = (fi18 | -9223372036854775808)
                                                                                    v125 := GenS13{F0: "\n\t", F1: (v118 - (v8 | (-(v8)))), F2: "\n\t"}
                                                                                    v124 -= -301
                                                                                    _ = v124
                                                                                    _ = v125
                                                                                }
                                                                            case -1:
                                                                                v15 = true
                                                                            default:
                                                                                var v126 int16 = 42
                                                                                _ = v126
                                                                            }
                                                                            delete(v4, 4294967295)
                                                                            _ = v122
                                                                        }
                                                                    case "\n\t":
                                                                        v1 += (v1 + 1.5)
                                                                    default:
                                                                        fmt.Println(v2, fi20)
                                                                    }
                                                                default:
                                                                    fmt.Println(fi20)
                                                                }
                                                                _ = v118
                                                            }
                                                            _ = v47
                                                        }
                                                        fmt.Println(v8)
                                                    }
                                                    if ((v2 % (((^(v2))) | 1)) >= v2) {
                                                        v2 = v2
                                                    } else {
                                                        v3 -= 9223372036854775807
                                                        v1 = (float32(0.1) - v1)
                                                        v127 := GenS14{F0: v15, F1: ""}
                                                        _ = v127
                                                    }
                                                }
                                                switch 0 {
                                                case 32768:
                                                    v1 -= v1
                                                case 128:
                                                    v128 := GenS4{F0: 65535, F1: (v8 - int8(v3)), F2: 1}
                                                    _ = v128
                                                case 127:
                                                    fmt.Println(v8, v3, v15)
                                                }
                                                v2 -= (v2 * 2)
                                            }
                                        default:
                                            v1 += v1
                                        }
                                    }
                                    fmt.Println(v6, fi18, fi18)
                                    v1 -= (v1 - ((3.14159 / v1) / v1))
                                    var v129 uint8 = (v6 ^ (v6 + v6))
                                    _ = v129
                                }
                            }
                        }
                        fmt.Println(v2, fi20)
                        fmt.Println(fi20, fi18, v1)
                        _ = fi20
                    }
                    _ = v19
                } else {
                    var v130 int16 = -1
                    _ = len(v5)
                    var v131 uint16 = 65535
                    _ = v130
                    _ = v131
                }
                v132 := GenS4{F0: 0, F1: 1, F2: 42}
                _ = v132
            } else {
                v2 += v2
            }
            _ = len(v13)
            _ = fi18
        }
    }
    delete(v5, 2147483648)
    v6 = (v6 % ((v6) | 1))
    var v133 float32 = (1.5 + 3.14159)
    for fi134 := 0; fi134 < 18; fi134++ {
        if (!(!(2147483647 < -882))) {
            var v135 int32 = -1
            _ = v135
        }
        fmt.Println(v15, fi134, v8)
        v136 := v4[(v3 >> ((255) & 63))]
        _ = fi134
        _ = v136
    }
    if (((1e10 / v1) / -2.25) != (-2.25 * (v133 / v133))) {
        var v137 int32 = 0
        v1 += v133
        v15 = v15
        _ = v137
    } else {
        for fi138 := 0; fi138 < 8; fi138++ {
            delete(v4, (v3 >> ((255) & 63)))
            if v15 {
                var v139 string = "foo bar"
                _ = v139
            }
            v2 -= 876
            _ = fi138
        }
        fmt.Println(v8)
        for fi140 := 0; fi140 < 3; fi140++ {
            v141 := GenS4{F0: 1, F1: v8, F2: -985}
            var v142 int8 = v8
            if len(v13) > 0 {
                v13[0] = (v3 | -663)
            }
            _ = fi140
            _ = v141
            _ = v142
        }
    }
    if v15 {
        v133 += 1e10
        v5[2147483647] = v3
        switch "foo bar" {
        case "hello":
            fmt.Println(v15, v6)
        }
    }
    var v143 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v133:", v133)
    fmt.Println("v143:", v143)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v133
    _ = v143
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
