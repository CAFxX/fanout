package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int64; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 string }
type GenS5 struct { F0 uint8; F1 uint32; F2 int }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 string; F1 uint64 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int16; F1 int8; F2 int }
type GenS14 struct { F0 string; F1 int8; F2 uint32 }
type GenS15 struct { F0 uint8; F1 uint32 }
type GenS16 struct { F0 uint; F1 int; F2 string }
type GenS17 struct { F0 string; F1 int8; F2 int }
type GenS18 struct { F0 int64; F1 uint8 }
type GenS19 struct { F0 int16; F1 int; F2 bool }
type GenS20 struct { F0 int8 }
type GenS21 struct { F0 uint }
type GenS22 struct { F0 uint16; F1 bool }
type GenS23 struct { F0 int8; F1 int64; F2 int16 }
type GenS24 struct { F0 bool; F1 int64; F2 int8 }
type GenS25 struct { F0 uint32; F1 uint }
type GenS26 struct { F0 uint64; F1 bool }
type GenS27 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS28 struct { F0 uint32; F1 uint32 }
type GenS29 struct { F0 uint32 }
type GenS30 struct { F0 int32; F1 string; F2 int64 }
func genFn0(p0 int16, p1 int32) uint32 {
    var v6 bool = (((p0 / (((-(p0))) | 1)) < (p0 % ((p0) | 1))) || true)
    v7 := GenS1{F0: 160, F1: (p0 / (((p0 | (p0 * p0))) | 1))}
    var v8 int8 = 100
    v9 := GenS2{F0: 42}
    v10 := []int{-9223372036854775808, -353}
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint8, p1 uint64, p2 uint8) int64 {
    v11 := []int{4294967295, int(p1), -326}
    if true {
        fmt.Println(p2, p0, p2)
        for fi12 := 0; fi12 < 19; fi12++ {
            if false {
                var v13 int = (fi12 % (((^(fi12))) | 1))
                if (true && false) {
                    var v14 uint = 903
                    _ = v14
                }
                _ = v13
            } else {
                var v15 int = (fi12 % ((828) | 1))
                if false {
                    if (65535 > 42) {
                        if len(v11) > 0 {
                            v11[0] = v15
                        }
                    }
                    v15 = -9223372036854775808
                } else {
                    if (0.0 <= (-2.25 - (((0.0 * (1.5 - (((float32(1e10) * 1e10) + 0.1) - ((float32(((1.5 - ((0.0 / (1e10 * 1.5)) * float64(((0.0 - 1.5) + 1.5)))) - ((3.14159 - -2.25) * 1.5))) / 1e10) * -2.25)))) * float32(float64((3.14159 / -2.25)))) + 1.5))) {
                        switch (^(fi12)) {
                        case -188:
                            p1 = (p1 >> ((2) & 63))
                        case 32767:
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        default:
                            p1 += (p1 / (((p1 / (((p1 & (^(p1)))) | 1))) | 1))
                        }
                        var v16 string = ("a" + "foo bar")
                        _ = v16
                    } else {
                        var v17 int32 = 1
                        _ = v17
                    }
                    v18 := GenS2{F0: 595}
                    for fi19 := 0; fi19 < 6; fi19++ {
                        if len(v11) > 0 {
                            v11[0] = fi12
                        }
                        if len(v11) > 0 {
                            v11[0] = v15
                        }
                        p0 = p2
                        _ = fi19
                    }
                    _ = v18
                }
                _ = v15
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                if (-1 <= -88) {
                    var v21 float64 = -2.25
                    var v22 string = "\n\t"
                    if true {
                        _ = len(v11)
                    } else {
                        v23 := GenS3{F0: 42, F1: 9223372036854775807, F2: 9223372036854775807}
                        fmt.Println(v21, p0, v21)
                        _ = v23
                    }
                    _ = v21
                    _ = v22
                }
                if len(v11) > 0 {
                    v11[0] = fi20
                }
                if (921 > (fi20 & (fi12 - fi12))) {
                    switch (fi20 >> ((2) & 63)) {
                    case 1:
                        v24 := GenS1{F0: uint64(p0), F1: 0}
                        if (255 == 128) {
                            p0 = 184
                            var v25 uint64 = p1
                            v26 := GenS1{F0: 18446744073709551615, F1: 0}
                            p1 += (v25 + 18446744073709551615)
                            var v27 uint64 = (p1 & (p1 >> ((42) & 63)))
                            _ = v25
                            _ = v26
                            _ = v27
                        } else {
                            if (68 >= (p2 / ((p0) | 1))) {
                                if false {
                                    p2 = 1
                                    v28 := GenS3{F0: fi20, F1: 42, F2: -733}
                                    p0 -= (p2 * (p0 % ((uint8(p1)) | 1)))
                                    _ = v28
                                } else {
                                    v29 := GenS4{F0: 42, F1: 65535, F2: "hello"}
                                    var v30 int8 = -1
                                    _ = v29
                                    _ = v30
                                }
                                v31 := GenS4{F0: 255, F1: 1, F2: "\n\t"}
                                p0 += 42
                                _ = v31
                            } else {
                                var v32 string = "foo bar"
                                fmt.Println(p1, p2)
                                _ = v32
                            }
                            if len(v11) > 0 {
                                v11[0] = (fi20 | 4294967295)
                            }
                            fmt.Println(p1, fi20)
                        }
                        var v33 int64 = 0
                        v34 := GenS5{F0: (^(p0)), F1: 2147483647, F2: fi20}
                        var v35 int32 = int32(p2)
                        switch fi20 {
                        case -1:
                            if (65535 < 0) {
                                if true {
                                    switch "\n\t" {
                                    case "a":
                                        v36 := GenS6{F0: (p0 ^ (p0 >> ((42) & 63))), F1: p2}
                                        _ = v36
                                    case "\n\t":
                                        fmt.Println(p0)
                                    case "hello":
                                        fmt.Println(fi20, v35, fi20)
                                    default:
                                        var v37 uint = 4294967295
                                        _ = v37
                                    }
                                    v33 -= (-(v33))
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                }
                            } else {
                                var v38 int8 = -32
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                if len(v11) > 0 {
                                    v11[0] = (fi12 % ((fi20) | 1))
                                }
                                _ = v38
                            }
                        case -7:
                            switch (fi12 | int(v35)) {
                            case 722:
                                if false {
                                    v39 := GenS2{F0: genFn0(0, (v35 + 2147483647))}
                                    var v40 uint64 = 235
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 float64 = 3.14159
                                    if (false && false) {
                                        v42 := GenS7{F0: 1, F1: ""}
                                        switch "" {
                                        case "":
                                            switch (^(fi20)) {
                                            case -9223372036854775808:
                                                v43 := GenS3{F0: fi12, F1: (v33 / ((int64(p0)) | 1)), F2: (-(v33))}
                                                _ = v43
                                            case -349:
                                                if (448 == 2) {
                                                    v35 += v35
                                                    fmt.Println(p1, p1)
                                                }
                                            case 4294967295:
                                                _ = len(v11)
                                            }
                                        }
                                        fmt.Println(fi20)
                                        _ = len(v11)
                                        _ = v42
                                    }
                                    _ = v41
                                }
                                switch fi12 {
                                case -9223372036854775808:
                                    v11 = append(v11, (fi12 - (fi20 / (((fi20 * 42)) | 1))))
                                case 127:
                                    p2 -= (p2 << ((0) & 63))
                                    v33 -= (v33 - v33)
                                }
                            case 0:
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                            }
                        case -2147483648:
                            var v44 int64 = (v33 * v33)
                            _ = v44
                        default:
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                        }
                        switch ("\n\t" + "foo bar") {
                        case "a":
                            if true {
                                v45 := GenS8{F0: 798}
                                _ = v45
                            }
                            if (!true) {
                                if true {
                                    v11 = append(v11, 9223372036854775807)
                                    p1 += p1
                                } else {
                                    p1 -= p1
                                }
                                switch fi20 {
                                case 9223372036854775807:
                                    switch ("" + "") {
                                    case "hello":
                                        p0 -= 204
                                    }
                                case 833:
                                    var v46 uint64 = 322
                                    _ = v46
                                case 121:
                                    if (false && true) {
                                        v47 := GenS9{F0: 32767}
                                        var v48 uint64 = (p1 * 688)
                                        _ = v47
                                        _ = v48
                                    }
                                default:
                                    if ((0 >= 2) || false) {
                                        p2 = p2
                                    } else {
                                        if ((true && false) || false) {
                                            if false {
                                                var v49 uint = 1
                                                if ((v49 + (^(v49))) > v49) {
                                                    v50 := GenS7{F0: uint8(p1), F1: ("foo bar" + ("foo bar" + "\n\t"))}
                                                    var v51 int64 = v33
                                                    _ = v50
                                                    _ = v51
                                                }
                                                _ = v49
                                            }
                                            _ = len(v11)
                                            _ = len(v11)
                                            v35 = 59
                                            p0 -= 38
                                        } else {
                                            var v52 uint64 = p1
                                            p2 = uint8(p1)
                                            _ = v52
                                        }
                                        if false {
                                            if ((p2 % ((57) | 1)) != (p2 - p2)) {
                                                p1 = (p1 ^ 644)
                                                if true {
                                                    _ = len(v11)
                                                } else {
                                                    v53 := GenS8{F0: 42}
                                                    var v54 uint8 = uint8(p1)
                                                    if (1 < uint16(v54)) {
                                                        v35 += (^(v35))
                                                        v35 += v35
                                                        p0 -= (p2 >> ((255) & 63))
                                                    }
                                                    _ = v53
                                                    _ = v54
                                                }
                                                var v55 float64 = 0.0
                                                _ = v55
                                            }
                                            fmt.Println(p0, v33, fi12)
                                            p1 += 4294967295
                                        } else {
                                            if false {
                                                v56 := GenS10{F0: (^(v33))}
                                                v11 = append(v11, fi12)
                                                v33 -= (-(v33))
                                                _ = v56
                                            } else {
                                                var v57 int16 = 42
                                                if (!true) {
                                                    v35 = 42
                                                    v11 = append(v11, fi12)
                                                }
                                                _ = v57
                                            }
                                            if true {
                                                fmt.Println(fi20)
                                                switch int(v35) {
                                                case 127:
                                                    v58 := GenS11{F0: ((("a" + ("\n\t" + (((("\n\t" + "foo bar") + "") + "\n\t") + "a"))) + "hello") + "foo bar"), F1: p1}
                                                    _ = v58
                                                case 42:
                                                    var v59 int = int(p0)
                                                    _ = v59
                                                case 9223372036854775807:
                                                    switch "" {
                                                    case "":
                                                        v11 = append(v11, fi12)
                                                    case "a":
                                                        _ = len(v11)
                                                    case "hello":
                                                        if (!true) {
                                                            p0 = (p2 | (p2 - (p2 - p0)))
                                                            fmt.Println(p0)
                                                            fmt.Println(fi20)
                                                            _ = len(v11)
                                                            if len(v11) > 0 {
                                                                v11[0] = (fi12 | 52)
                                                            }
                                                        }
                                                    default:
                                                        if len(v11) >= 2 {
                                                            v11 = v11[1:len(v11)]
                                                        }
                                                    }
                                                default:
                                                    if (42 <= 4294967295) {
                                                        fmt.Println(p1, p2, v33)
                                                        p2 += uint8(p1)
                                                        var v60 float32 = (3.14159 - 0.0)
                                                        v33 += v33
                                                        _ = v60
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(p2, v35, p1)
                                    switch ((("" + ("a" + "foo bar")) + ((("foo bar" + "foo bar") + ("\n\t" + "")) + ("foo bar" + "foo bar"))) + "foo bar") {
                                    case "\n\t":
                                        _ = len(v11)
                                        fmt.Println(p1, v35, fi20)
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                        var v61 uint32 = 2147483648
                                        v62 := GenS12{F0: 32767}
                                        v63 := GenS13{F0: 0, F1: 127, F2: int(v61)}
                                        var v64 uint8 = p2
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                    case "a":
                                        v33 = (v33 << ((42) & 63))
                                    case "hello":
                                        switch "foo bar" {
                                        case "":
                                            if len(v11) >= 2 {
                                                v11 = v11[1:len(v11)]
                                            }
                                            var v65 uint = 0
                                            _ = v65
                                        default:
                                            fmt.Println(v35, p2)
                                        }
                                    default:
                                        v33 = (^(v33))
                                    }
                                }
                                p1 += 284
                            } else {
                                v66 := GenS14{F0: "\n\t", F1: -128, F2: 42}
                                _ = v66
                            }
                        default:
                            v11 = append(v11, 0)
                        }
                        v67 := GenS8{F0: 32767}
                        var v68 int32 = v35
                        switch (fi12 ^ (fi12 & fi12)) {
                        case -9223372036854775808:
                            if true {
                                v69 := GenS15{F0: 61, F1: 2147483647}
                                fmt.Println(v68, v35)
                                var v70 float64 = 1e10
                                _ = v69
                                _ = v70
                            } else {
                                v33 = v33
                                var v71 uint32 = 270
                                v68 += (v35 | -807)
                                _ = v71
                            }
                        case -1:
                            if (!(255 <= 65535)) {
                                switch "\n\t" {
                                case "":
                                    switch -9223372036854775808 {
                                    case 4294967295:
                                        v35 += 0
                                    case 499:
                                        v72 := GenS4{F0: 333, F1: 42, F2: "\n\t"}
                                        var v73 uint8 = (p2 & (p2 | (p2 + p2)))
                                        _ = v72
                                        _ = v73
                                    default:
                                        if ((65535 <= 209) && true) {
                                            var v74 float32 = 0.0
                                            var v75 uint = 4294967295
                                            v68 += int32(p0)
                                            _ = v74
                                            _ = v75
                                        }
                                    }
                                default:
                                    fmt.Println(p0)
                                    fmt.Println(p2, fi20)
                                    var v76 uint32 = 42
                                    _ = v76
                                }
                                var v77 int32 = (v68 << ((65535) & 63))
                                switch (fi12 ^ -10) {
                                case 32768:
                                    v78 := GenS8{F0: 382}
                                    _ = v78
                                case 32767:
                                    if true {
                                        p0 += 255
                                    }
                                default:
                                    var v79 int8 = 42
                                    _ = v79
                                }
                                _ = v77
                            } else {
                                _ = len(v11)
                            }
                        default:
                            switch fi20 {
                            case 42:
                                var v80 uint32 = 4294967295
                                _ = v80
                            }
                            if true {
                                var v81 uint64 = p1
                                var v82 uint64 = (v81 | (^(v81)))
                                fmt.Println(v33, v35, v33)
                                var v83 int16 = 0
                                switch (((("\n\t" + "hello") + "hello") + "hello") + ("" + "\n\t")) {
                                case "\n\t":
                                    v68 -= v68
                                case "foo bar":
                                    var v84 uint64 = v82
                                    _ = v84
                                case "a":
                                    switch "a" {
                                    case "":
                                        fmt.Println(p2, v82)
                                    }
                                }
                                _ = v81
                                _ = v82
                                _ = v83
                            }
                            v11 = append(v11, 42)
                            fmt.Println(p1, fi20)
                            var v85 uint64 = (p1 >> ((18446744073709551615) & 63))
                            _ = v85
                        }
                        if (42 == -10) {
                            switch (fi12 % ((fi20) | 1)) {
                            case 42:
                                fmt.Println(p1)
                            default:
                                if false {
                                    v11 = append(v11, (fi12 << ((0) & 63)))
                                } else {
                                    p0 = (p0 * p0)
                                }
                            }
                        }
                        v86 := GenS16{F0: 303, F1: fi20, F2: ""}
                        _ = v24
                        _ = v33
                        _ = v34
                        _ = v35
                        _ = v67
                        _ = v68
                        _ = v86
                    case 128:
                        if false {
                            v87 := GenS17{F0: ("\n\t" + "hello"), F1: -92, F2: -2147483648}
                            var v88 bool = (!(1e10 <= 1e10))
                            _ = v87
                            _ = v88
                        }
                    default:
                        var v89 float32 = 3.14159
                        _ = v89
                    }
                    p0 -= 168
                } else {
                    var v90 int16 = 734
                    v11 = append(v11, 1)
                    if (true || false) {
                        if true {
                            if (!true) {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                            } else {
                                fmt.Println(fi12)
                                fmt.Println(p0, fi12, v90)
                                switch fi20 {
                                case 42:
                                    v11 = append(v11, 9223372036854775807)
                                case 2147483647:
                                    _ = len(v11)
                                case 498:
                                    if (fi20 != fi20) {
                                        if (1000 <= v90) {
                                            var v91 float32 = (float32((float64(float32(1.5)) * 1.5)) + -2.25)
                                            _ = v91
                                        }
                                        var v92 int8 = -128
                                        var v93 uint8 = (p0 << ((65535) & 63))
                                        if true {
                                            var v94 string = (("\n\t" + "") + "a")
                                            var v95 int16 = v90
                                            _ = len(v11)
                                            _ = v94
                                            _ = v95
                                        } else {
                                            if len(v11) > 0 {
                                                v11[0] = 44
                                            }
                                            if false {
                                                if false {
                                                    fmt.Println(v93)
                                                    _ = len(v11)
                                                    var v96 uint64 = 2
                                                    _ = v96
                                                } else {
                                                    if true {
                                                        var v97 uint8 = p2
                                                        _ = v97
                                                    } else {
                                                        p2 += 1
                                                        var v98 int8 = (v92 + 42)
                                                        var v99 float64 = 3.14159
                                                        _ = v98
                                                        _ = v99
                                                    }
                                                }
                                                v100 := GenS14{F0: "a", F1: (v92 | v92), F2: 96}
                                                _ = v100
                                            }
                                            if ((p0 * (p2 << ((255) & 63))) == (p0 - p0)) {
                                                var v101 int32 = -326
                                                var v102 int8 = v92
                                                if true {
                                                    var v103 float32 = float32(1e10)
                                                    var v104 float64 = (float64(v103) * 0.1)
                                                    fmt.Println(v102, p1, p1)
                                                    _ = v103
                                                    _ = v104
                                                }
                                                _ = v101
                                                _ = v102
                                            }
                                        }
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                        _ = v92
                                        _ = v93
                                    } else {
                                        if len(v11) >= 2 {
                                            v11 = v11[1:len(v11)]
                                        }
                                        var v105 int16 = (v90 ^ -1)
                                        _ = v105
                                    }
                                default:
                                    v106 := GenS16{F0: 2, F1: fi20, F2: ""}
                                    _ = v106
                                }
                                var v107 uint8 = 42
                                if true {
                                    if (float32(((float64(3.14159) / 1.0) - -2.25)) > -2.25) {
                                        v108 := GenS14{F0: "", F1: -1, F2: genFn0((v90 / ((v90) | 1)), 999)}
                                        _ = v108
                                    }
                                    var v109 uint32 = 65535
                                    _ = v109
                                }
                                _ = v107
                            }
                            v90 = v90
                        }
                        if (((((("\n\t" + "hello") + ("foo bar" + (("foo bar" + "foo bar") + "hello"))) + "a") + ("foo bar" + ("a" + (("a" + "a") + "")))) < "\n\t") || (973 > 2147483647)) {
                            _ = len(v11)
                            var v110 int64 = -9223372036854775808
                            _ = v110
                        } else {
                            if (true || false) {
                                var v111 bool = false
                                var v112 float64 = (3.14159 * ((0.0 / -2.25) * float64(1e10)))
                                if v111 {
                                    if v111 {
                                        p0 = (p2 * p2)
                                        fmt.Println(fi12, v90, fi12)
                                        fmt.Println(fi12)
                                        switch fi20 {
                                        case 42:
                                            var v113 uint64 = 4294967295
                                            _ = v113
                                        case -2147483648:
                                            if len(v11) > 0 {
                                                v11[0] = fi12
                                            }
                                        case 102:
                                            v114 := GenS4{F0: 42, F1: 0, F2: "foo bar"}
                                            _ = v114
                                        default:
                                            var v115 int = 128
                                            _ = v115
                                        }
                                    } else {
                                        var v116 uint = 2
                                        _ = v116
                                    }
                                    if (p1 != 2) {
                                        p1 -= p1
                                        if (4294967295 != 127) {
                                            if ((-9223372036854775808 == -9223372036854775808) || (-776 == -647)) {
                                                v90 = 42
                                                p2 = (p2 * (p2 >> ((255) & 63)))
                                                v117 := GenS18{F0: 9223372036854775807, F1: p2}
                                                if len(v11) > 0 {
                                                    v11[0] = (fi20 / (((^(fi20))) | 1))
                                                }
                                                if (-2147483648 != 0) {
                                                    if (1 > -2147483648) {
                                                        if (0 >= 4294967295) {
                                                            v118 := GenS11{F0: "a", F1: (p1 ^ (p1 >> ((550) & 63)))}
                                                            _ = v118
                                                        }
                                                        p0 += (p0 | 53)
                                                    }
                                                    if ((p0 == 95) && (v111 && (uint32(p1) >= 4294967295))) {
                                                        if v111 {
                                                            p1 += p1
                                                        } else {
                                                            v119 := GenS11{F0: "hello", F1: p1}
                                                            _ = v119
                                                        }
                                                        _ = len(v11)
                                                    }
                                                    p1 -= (p1 / ((p1) | 1))
                                                    fmt.Println(v112, fi20)
                                                    if true {
                                                        switch fi20 {
                                                        case 1:
                                                            var v120 int8 = -56
                                                            _ = v120
                                                        case 42:
                                                            fmt.Println(v90, v90)
                                                        default:
                                                            var v121 uint = 42
                                                            _ = v121
                                                        }
                                                    }
                                                } else {
                                                    v11 = append(v11, 32768)
                                                }
                                                switch "" {
                                                case "":
                                                    _ = len(v11)
                                                case "foo bar":
                                                    if (2 <= 42) {
                                                        var v122 float32 = 1.5
                                                        if (18446744073709551615 == 910) {
                                                            p2 = p0
                                                        } else {
                                                            p1 += (p1 % ((p1) | 1))
                                                        }
                                                        if len(v11) >= 2 {
                                                            v11 = v11[1:len(v11)]
                                                        }
                                                        _ = v122
                                                    }
                                                }
                                                _ = v117
                                            }
                                        } else {
                                            v123 := GenS19{F0: 1, F1: -9223372036854775808, F2: (uint8(p1) == (p2 + (^(p2))))}
                                            var v124 uint8 = (p2 << ((390) & 63))
                                            fmt.Println(p0, p1, fi20)
                                            if v111 {
                                                _ = len(v11)
                                            } else {
                                                if len(v11) > 0 {
                                                    v11[0] = 0
                                                }
                                            }
                                            _ = v123
                                            _ = v124
                                        }
                                        switch 42 {
                                        case 42:
                                            var v125 uint32 = 401
                                            _ = v125
                                        default:
                                            fmt.Println(fi20)
                                            var v126 int8 = 100
                                            if (6 == 4294967295) {
                                                v90 = (v90 >> ((4294967295) & 63))
                                                if v111 {
                                                    switch 80 {
                                                    case 42:
                                                        if len(v11) > 0 {
                                                            v11[0] = 615
                                                        }
                                                    case -500:
                                                        var v127 string = ("\n\t" + "hello")
                                                        _ = v127
                                                    }
                                                    _ = len(v11)
                                                } else {
                                                    switch 128 {
                                                    case 128:
                                                        v128 := GenS20{F0: (v126 % ((85) | 1))}
                                                        _ = v128
                                                    case 127:
                                                        var v129 uint = 42
                                                        _ = v129
                                                    }
                                                    fmt.Println(v112)
                                                    if v111 {
                                                        if (v111 || (!v111)) {
                                                            fmt.Println(v90, p0, v111)
                                                            v130 := GenS21{F0: 65535}
                                                            p2 += p2
                                                            _ = v130
                                                        } else {
                                                            v112 = 0.1
                                                            v11 = append(v11, (fi20 / ((32767) | 1)))
                                                            v90 = int16(p0)
                                                        }
                                                        var v131 uint = 0
                                                        var v132 float64 = v112
                                                        _ = v131
                                                        _ = v132
                                                    } else {
                                                        v11 = append(v11, fi20)
                                                        if len(v11) > 0 {
                                                            v11[0] = fi20
                                                        }
                                                    }
                                                }
                                                p0 += 1
                                                v133 := GenS4{F0: 1, F1: 677, F2: "foo bar"}
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                                _ = v133
                                            }
                                            v134 := GenS22{F0: 1, F1: (float64((-2.25 - 0.1)) < v112)}
                                            switch -1 {
                                            case -824:
                                                fmt.Println(v112)
                                            case -2147483648:
                                                if len(v11) >= 2 {
                                                    v11 = v11[1:len(v11)]
                                                }
                                            case 32767:
                                                p2 -= 127
                                                fmt.Println(p2, fi20)
                                            }
                                            _ = v126
                                            _ = v134
                                        }
                                    }
                                    if (!v111) {
                                        var v135 float32 = 1e10
                                        v135 -= (v135 / float32((v112 / 1e10)))
                                        v136 := GenS23{F0: -98, F1: 9223372036854775807, F2: v90}
                                        v137 := GenS24{F0: (!true), F1: -9223372036854775808, F2: 1}
                                        fmt.Println(v135, fi20)
                                        var v138 int = fi20
                                        _ = v135
                                        _ = v136
                                        _ = v137
                                        _ = v138
                                    } else {
                                        switch ("foo bar" + ((("a" + "\n\t") + "\n\t") + "\n\t")) {
                                        case "":
                                            if true {
                                                var v139 uint16 = 255
                                                v140 := GenS6{F0: p2, F1: p2}
                                                if v111 {
                                                    v112 = 3.14159
                                                    var v141 uint64 = (p1 & p1)
                                                    _ = v141
                                                } else {
                                                    v142 := GenS6{F0: (^(p0)), F1: 255}
                                                    fmt.Println(fi20)
                                                    _ = v142
                                                }
                                                _ = v139
                                                _ = v140
                                            }
                                        case "foo bar":
                                            v11 = append(v11, (fi12 / ((438) | 1)))
                                        case "\n\t":
                                            if v111 {
                                                switch (fi12 ^ (fi12 + (fi20 << ((65535) & 63)))) {
                                                case -9223372036854775808:
                                                    if ((-2.25 - v112) >= 1.5) {
                                                        fmt.Println(v111, p1, fi20)
                                                        v11 = append(v11, fi20)
                                                        v143 := GenS7{F0: 0, F1: ("foo bar" + "\n\t")}
                                                        _ = v143
                                                    }
                                                case -819:
                                                    var v144 int64 = 42
                                                    _ = v144
                                                case -1:
                                                    var v145 float32 = -2.25
                                                    _ = v145
                                                }
                                                p2 = p2
                                            }
                                        default:
                                            v90 = (^(v90))
                                        }
                                    }
                                }
                                fmt.Println(p2, fi12)
                                _ = len(v11)
                                _ = v111
                                _ = v112
                            } else {
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                            }
                        }
                    } else {
                        fmt.Println(p2, p1)
                        switch "a" {
                        case "":
                            var v146 uint16 = 2
                            _ = v146
                        case "a":
                            var v147 uint = 255
                            _ = v147
                        }
                    }
                    _ = v90
                }
                switch (^(fi20)) {
                case 2147483647:
                    var v148 bool = (2147483648 <= 1)
                    _ = v148
                case -897:
                    var v149 int8 = -1
                    _ = v149
                default:
                    var v150 bool = (((0.1 - 1.5) * 1e10) != 1e10)
                    _ = v150
                }
                _ = fi20
            }
            _ = fi12
        }
        switch 9223372036854775807 {
        case -1:
            var v151 int16 = 1
            _ = v151
        case 2147483647:
            v152 := GenS25{F0: 42, F1: 0}
            _ = v152
        case 42:
            var v153 bool = (!false)
            _ = v153
        }
    }
    switch 127 {
    case 488:
        var v154 int = -565
        _ = v154
    case 9223372036854775807:
        var v155 int = 4294967295
        _ = v155
    default:
        fmt.Println(p1, p0, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (2147483647 == -2147483648)
    v3 := make(map[uint32]uint32)
    var v4 string = (("\n\t" + "") + "\n\t")
    var v5 int16 = -32768
    for fi156 := 0; fi156 < 10; fi156++ {
        fmt.Println(v4, v2)
        if v2 {
            v157 := GenS26{F0: 0, F1: false}
            _ = v157
        } else {
            if v2 {
                v158 := GenS27{F0: 2147483648, F1: 48, F2: 42}
                v159 := GenS28{F0: 1, F1: 657}
                _ = v158
                _ = v159
            }
            v2 = false
        }
        if v2 {
            v160 := v3[377]
            _ = v160
        } else {
            v2 = v2
            v2 = (!((1e10 >= -2.25) && false))
            switch (fi156 | (fi156 ^ (fi156 + (fi156 - (-(fi156)))))) {
            case -994:
                v5 = (v5 & (v5 - (^(v5))))
            case 936:
                for fi161 := 0; fi161 < 20; fi161++ {
                    var v162 uint = 4294967295
                    var v163 uint32 = 159
                    _ = len(v3)
                    _ = fi161
                    _ = v162
                    _ = v163
                }
            default:
                if (!v2) {
                    v164 := GenS3{F0: fi156, F1: 1, F2: 9223372036854775807}
                    _ = v164
                }
            }
        }
        _ = fi156
    }
    if v2 {
        v4 = "hello"
    }
    fmt.Println(v2, v4)
    var v165 bool = (v2 || ((float32(0.1) / (float32((-2.25 - ((-2.25 + 3.14159) + (float64(0.0) - (0.0 - 0.0))))) / -2.25)) != (3.14159 / 1e10)))
    fmt.Println(v2)
    v166 := GenS29{F0: 0}
    switch 0 {
    case -1:
        v167 := GenS9{F0: 32767}
        _ = v167
    case 42:
        fmt.Println(v5)
    case 27:
        v5 += (^(v5))
    }
    v4 = ("\n\t" + v4)
    v168 := GenS30{F0: -2147483648, F1: "hello", F2: 392}
    if v2 {
        var v169 int16 = v5
        _ = v169
    } else {
        fmt.Println(v5, v5)
    }
    v165 = v165
    if v2 {
        if v2 {
            var v170 string = v4
            if (2147483648 >= 2) {
                v171 := GenS5{F0: 179, F1: 2147483647, F2: -2147483648}
                _ = v171
            } else {
                var v172 uint = 255
                _ = v172
            }
            v4 = v4
            _ = v170
        }
    }
    for fi173 := 0; fi173 < 18; fi173++ {
        for fi174 := 0; fi174 < 15; fi174++ {
            v175 := GenS11{F0: "foo bar", F1: 18446744073709551615}
            _ = fi174
            _ = v175
        }
        _ = fi173
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v165:", v165)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v165
    _ = v166
    _ = v168
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
