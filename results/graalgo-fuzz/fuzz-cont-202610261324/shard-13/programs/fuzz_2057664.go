package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 int64; F1 uint; F2 uint }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int64; F1 int; F2 int64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint32; F1 bool; F2 uint }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int16; F1 string; F2 uint16 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 uint8; F1 int8; F2 int64 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint16; F1 int }
type GenS18 struct { F0 bool; F1 bool }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint8 }
type GenS21 struct { F0 int8; F1 int32; F2 uint32 }
type GenS22 struct { F0 bool }
type GenS23 struct { F0 uint8 }
type GenS24 struct { F0 uint32; F1 uint64 }
type GenS25 struct { F0 int; F1 bool }
type GenS26 struct { F0 int32 }
type GenS27 struct { F0 int64; F1 int32 }
type GenS28 struct { F0 int8; F1 int64 }
func genFn0(p0 int, p1 int) uint16 {
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{p0, p0}
        var v9 int8 = -16
        v9 = v9
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 uint = 123
        _ = fi10
        _ = v11
    }
    p0 -= (p0 * (p0 / (((p1 * (p0 + (p0 ^ p1)))) | 1)))
    for fi12 := 0; fi12 < 5; fi12++ {
        switch 0 {
        case -866:
            v13 := GenS1{F0: 128, F1: 4294967295, F2: p0}
            _ = v13
        default:
            v14 := GenS1{F0: 255, F1: 1, F2: 127}
            _ = v14
        }
        p1 = (p0 % ((fi12) | 1))
        _ = fi12
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        if true {
            v16 := GenS1{F0: 228, F1: 604, F2: -2147483648}
            _ = v16
        } else {
            v17 := GenS2{F0: -1, F1: 65535, F2: 29}
            for fi18 := 0; fi18 < 17; fi18++ {
                v19 := GenS2{F0: 532, F1: 1, F2: 192}
                p1 = p0
                v20 := []int{}
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        }
        _ = fi15
    }
    _ = p0
    _ = p1
    return 264
}
func genFn1(p0 uint, p1 string, p2 int) uint8 {
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := GenS3{F0: "", F1: false}
        fmt.Println(p2, p2, p1)
        v23 := []int{}
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 16; fi24++ {
        p1 = p1
        var v25 int32 = 42
        var v26 uint64 = 896
        _ = fi24
        _ = v25
        _ = v26
    }
    var v27 int8 = 42
    if (42 >= 2) {
        p2 += (p2 + p2)
        var v28 string = ""
        _ = v28
    }
    p0 = (p0 << ((p0) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 127
}

func main() {
    var v1 int32 = 42
    v2 := []string{}
    var v3 uint = 65535
    v4 := []string{"", "\n\t"}
    v5 := []uint32{65535, 42}
    var v6 int8 = 42
    for fi29 := 0; fi29 < 15; fi29++ {
        v6 = (v6 >> ((v3) & 63))
        var v30 uint = (v3 * v3)
        var v31 string = "hello"
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int32 = -751
    var v33 string = "a"
    v34 := GenS4{F0: 42}
    fmt.Println(v33)
    for fi35 := 0; fi35 < 9; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            v37 := GenS1{F0: 0, F1: 4294967295, F2: -1}
            _ = fi36
            _ = v37
        }
        _ = fi35
    }
    switch 127 {
    case 2147483647:
        v1 += (v1 << (((v3 & 255)) & 63))
    default:
        var v38 float64 = float64(float32(((0.0 / (1.5 - (1.5 + float64((1e10 - -2.25))))) + ((-2.25 - 0.0) / 1.0))))
        _ = v38
    }
    var v39 uint = (v3 - 1)
    if ((0.0 - (0.0 - (float32(0.0) + float32(0.0)))) >= float32(float64(1.5))) {
        var v40 int32 = (v1 / ((2147483647) | 1))
        if (-2.25 != 1e10) {
            v40 = v32
            for fi41 := 0; fi41 < 16; fi41++ {
                for fi42 := 0; fi42 < 12; fi42++ {
                    v6 -= (-(v6))
                    if (-2.25 <= (-2.25 * (((1e10 / (1.5 + 3.14159)) * float64(float32(float64((0.0 * (0.0 + 3.14159)))))) + (float64(3.14159) / -2.25)))) {
                        var v43 string = v33
                        switch (v33 + "\n\t") {
                        case "":
                            v44 := GenS5{F0: -1}
                            _ = v44
                        case "hello":
                            if (v6 >= v6) {
                                var v45 int16 = 1000
                                v46 := GenS3{F0: v43, F1: (false || (880 == 0))}
                                _ = v45
                                _ = v46
                            }
                        case "a":
                            switch ("foo bar" + (v43 + "hello")) {
                            case "a":
                                var v47 uint8 = 131
                                _ = v47
                            case "":
                                fmt.Println(v1, v33)
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                            case "hello":
                                switch v43 {
                                case "a":
                                    v33 = "a"
                                    if false {
                                        _ = len(v2)
                                        v39 = v3
                                        fmt.Println(v33, v33, fi42)
                                    } else {
                                        v40 += v32
                                    }
                                    if (!(!true)) {
                                        v48 := GenS2{F0: 732, F1: v39, F2: (v39 << ((65535) & 63))}
                                        var v49 int = 32767
                                        v50 := GenS6{F0: v6, F1: 1}
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    } else {
                                        v3 = (v39 & v39)
                                        if (!(genFn0((fi41 - 9223372036854775807), fi42) < 32767)) {
                                            v4 = append(v4, "foo bar")
                                            switch fi41 {
                                            case -903:
                                                v3 = (^(v39))
                                            default:
                                                v51 := GenS7{F0: 4294967295}
                                                _ = v51
                                            }
                                            if (false || (0.0 != 1e10)) {
                                                if false {
                                                    switch v33 {
                                                    case "a":
                                                        var v52 int32 = v32
                                                        _ = v52
                                                    case "\n\t":
                                                        var v53 float64 = ((0.0 - (float64((((0.0 + float32(-2.25)) + 0.0) * ((0.0 * ((0.0 + (1.5 / float32((1.5 + 1.5)))) * 3.14159)) * ((3.14159 + 0.0) * 1.5)))) - -2.25)) / 3.14159)
                                                        _ = v53
                                                    }
                                                    switch fi41 {
                                                    case 0:
                                                        if true {
                                                            if ((float32(float64((-2.25 + 1.5))) - 3.14159) == float32(0.0)) {
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                v39 += v39
                                                                var v54 string = ""
                                                                _ = v54
                                                            } else {
                                                                var v55 uint32 = 2147483648
                                                                fmt.Println(v3, v40, v3)
                                                                if false {
                                                                    if (!(false || false)) {
                                                                        v56 := GenS8{F0: 1, F1: (fi41 ^ -9223372036854775808), F2: 1}
                                                                        var v57 int8 = v6
                                                                        _ = v56
                                                                        _ = v57
                                                                    }
                                                                    switch v43 {
                                                                    case "foo bar":
                                                                        if len(v5) > 0 {
                                                                            v5[0] = 65535
                                                                        }
                                                                    case "hello":
                                                                        v32 = v1
                                                                    case "":
                                                                        if false {
                                                                            var v58 uint16 = 42
                                                                            if false {
                                                                                fmt.Println(v1)
                                                                                if ((v32 & (v1 / ((v1) | 1))) <= v32) {
                                                                                    v40 = -214
                                                                                    _ = len(v4)
                                                                                } else {
                                                                                    v4 = append(v4, v33)
                                                                                    v55 = 2147483648
                                                                                }
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = v33
                                                                                }
                                                                            } else {
                                                                                v43 = v33
                                                                                switch v33 {
                                                                                case "\n\t":
                                                                                    switch ("a" + v33) {
                                                                                    case "\n\t":
                                                                                        var v59 uint = v3
                                                                                        _ = v59
                                                                                    case "":
                                                                                        var v60 uint16 = genFn0((-(fi41)), (fi41 << (((v39 / ((uint(v58)) | 1))) & 63)))
                                                                                        _ = v60
                                                                                    case "foo bar":
                                                                                        var v61 int8 = v6
                                                                                        _ = v61
                                                                                    }
                                                                                case "a":
                                                                                    var v62 uint = v39
                                                                                    _ = v62
                                                                                case "foo bar":
                                                                                    switch v43 {
                                                                                    case "":
                                                                                        v32 -= -1
                                                                                    case "foo bar":
                                                                                        v63 := GenS9{F0: v6}
                                                                                        _ = v63
                                                                                    default:
                                                                                        var v64 int64 = 36
                                                                                        _ = v64
                                                                                    }
                                                                                }
                                                                            }
                                                                            v58 = 851
                                                                            _ = v58
                                                                        }
                                                                    }
                                                                    var v65 int64 = -518
                                                                    _ = v65
                                                                } else {
                                                                    v55 = uint32(v3)
                                                                    v3 = (^(v39))
                                                                    v43 = v43
                                                                }
                                                                var v66 uint8 = 2
                                                                v6 = 100
                                                                fmt.Println(v43, fi41, v55)
                                                                switch ((v33 + v43) + v43) {
                                                                case "hello":
                                                                    v67 := GenS3{F0: v43, F1: ((42 <= -32768) || ((65535 >= 65535) || false))}
                                                                    _ = v67
                                                                case "foo bar":
                                                                    fmt.Println(v43, v1, v66)
                                                                }
                                                                _ = v55
                                                                _ = v66
                                                            }
                                                            v40 = v1
                                                            v4 = append(v4, ("hello" + ""))
                                                        }
                                                        var v68 string = v33
                                                        fmt.Println(v1, v32, fi42)
                                                        _ = v68
                                                    case 567:
                                                        if false {
                                                            fmt.Println(v43)
                                                            v43 = "\n\t"
                                                            fmt.Println(v32, v33)
                                                            v69 := GenS3{F0: ("a" + "\n\t"), F1: true}
                                                            _ = v69
                                                        }
                                                    case 42:
                                                        var v70 uint64 = 467
                                                        _ = v70
                                                    default:
                                                        v71 := GenS6{F0: -128, F1: v3}
                                                        if true {
                                                            v4 = append(v4, (v43 + v33))
                                                            var v72 uint64 = 42
                                                            _ = v72
                                                        } else {
                                                            var v73 uint = 42
                                                            fmt.Println(v1, fi41)
                                                            _ = len(v4)
                                                            switch fi42 {
                                                            case 692:
                                                                switch v43 {
                                                                case "foo bar":
                                                                    v74 := GenS10{F0: 65535, F1: (42 != 2), F2: v39}
                                                                    var v75 int64 = 42
                                                                    v43 = (("" + (v43 + v33)) + (v33 + v43))
                                                                    _ = v74
                                                                    _ = v75
                                                                case "a":
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v76 float32 = (-2.25 + (0.0 * (1.5 + 3.14159)))
                                                                        _ = v76
                                                                    case "\n\t":
                                                                        if false {
                                                                            if ((!((v40 >> (((v3 ^ 18446744073709551615)) & 63)) <= v40)) && (((!(42 != (fi41 / ((32767) | 1)))) && (!((18446744073709551615 == 539) || false))) || (!(!(!(!false)))))) {
                                                                                var v77 int16 = 42
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                _ = v77
                                                                            }
                                                                        } else {
                                                                            var v78 int64 = -743
                                                                            v79 := GenS3{F0: v33, F1: (!((4294967295 != 948) && (v43 > "a")))}
                                                                            switch "a" {
                                                                            case "hello":
                                                                                var v80 int64 = (^(v78))
                                                                                _ = v80
                                                                            case "":
                                                                                switch v43 {
                                                                                case "hello":
                                                                                    v43 = v33
                                                                                case "foo bar":
                                                                                    var v81 bool = (!(!(436 <= (v40 * v1))))
                                                                                    _ = v81
                                                                                case "":
                                                                                    v39 = 18446744073709551615
                                                                                default:
                                                                                    var v82 float32 = 1.5
                                                                                    _ = v82
                                                                                }
                                                                            default:
                                                                                v83 := GenS11{F0: true}
                                                                                _ = v83
                                                                            }
                                                                            v78 -= v78
                                                                            _ = v78
                                                                            _ = v79
                                                                        }
                                                                        if (255 != 42) {
                                                                            v3 -= v3
                                                                        }
                                                                    case "a":
                                                                        var v84 uint8 = 0
                                                                        _ = v84
                                                                    }
                                                                case "\n\t":
                                                                    switch (^(fi42)) {
                                                                    case -2147483648:
                                                                        v3 += v3
                                                                    case 2147483647:
                                                                        switch v43 {
                                                                        case "a":
                                                                            if false {
                                                                                if ((v73 % (((v3 << ((v39) & 63))) | 1)) < (v39 * 0)) {
                                                                                    if (!false) {
                                                                                        if true {
                                                                                            var v85 int8 = v6
                                                                                            var v86 uint32 = 65535
                                                                                            switch 651 {
                                                                                            case 128:
                                                                                                var v87 uint16 = 255
                                                                                                _ = v87
                                                                                            case 0:
                                                                                                var v88 int16 = -32768
                                                                                                _ = v88
                                                                                            }
                                                                                            _ = v85
                                                                                            _ = v86
                                                                                        }
                                                                                        v6 = (v6 % ((v6) | 1))
                                                                                        v39 += v3
                                                                                    } else {
                                                                                        v2 = append(v2, (v33 + "a"))
                                                                                    }
                                                                                    var v89 uint = 39
                                                                                    if false {
                                                                                        if false {
                                                                                            v90 := GenS6{F0: v6, F1: (v89 + 4294967295)}
                                                                                            if true {
                                                                                                var v91 string = (v33 + ("a" + (v43 + v33)))
                                                                                                _ = v91
                                                                                            } else {
                                                                                                var v92 uint = (v39 >> (((v39 ^ v3)) & 63))
                                                                                                var v93 float32 = -2.25
                                                                                                var v94 uint64 = 42
                                                                                                if (true || (1.5 <= 0.0)) {
                                                                                                    fmt.Println(fi41, fi42)
                                                                                                    if (false || (-2.25 > 1.5)) {
                                                                                                        v89 = 1
                                                                                                        v95 := GenS12{F0: -1, F1: "", F2: genFn0(fi42, 1)}
                                                                                                        _ = v95
                                                                                                    }
                                                                                                    v33 = v43
                                                                                                } else {
                                                                                                    if true {
                                                                                                        v1 -= v1
                                                                                                        v96 := GenS9{F0: 47}
                                                                                                        _ = v96
                                                                                                    } else {
                                                                                                        var v97 int8 = 0
                                                                                                        var v98 int64 = 178
                                                                                                        _ = v97
                                                                                                        _ = v98
                                                                                                    }
                                                                                                    var v99 int16 = 32767
                                                                                                    v39 = v39
                                                                                                    v3 = (v92 ^ v3)
                                                                                                    _ = v99
                                                                                                }
                                                                                                switch (v33 + ((v43 + v33) + v33)) {
                                                                                                case "":
                                                                                                    v94 -= v94
                                                                                                case "a":
                                                                                                    var v100 string = (("hello" + v43) + ("a" + "a"))
                                                                                                    _ = v100
                                                                                                case "hello":
                                                                                                    _ = len(v4)
                                                                                                }
                                                                                                _ = v92
                                                                                                _ = v93
                                                                                                _ = v94
                                                                                            }
                                                                                            v3 = 4294967295
                                                                                            _ = v90
                                                                                        }
                                                                                        fmt.Println(v73, fi41)
                                                                                        if (!(0 < 42)) {
                                                                                            if true {
                                                                                                v32 += (v1 << ((v73) & 63))
                                                                                                var v101 bool = (false || true)
                                                                                                v73 -= 0
                                                                                                v5 = append(v5, 0)
                                                                                                _ = v101
                                                                                            }
                                                                                        }
                                                                                        var v102 uint64 = 18446744073709551615
                                                                                        v3 = v89
                                                                                        _ = v102
                                                                                    }
                                                                                    _ = v89
                                                                                }
                                                                            }
                                                                            switch (fi42 / ((fi41) | 1)) {
                                                                            case -84:
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    if (false || (0 != int16(v39))) {
                                                                                        v73 = 65535
                                                                                        if false {
                                                                                            fmt.Println(v6, v43)
                                                                                        }
                                                                                        var v103 uint32 = 2
                                                                                        _ = v103
                                                                                    } else {
                                                                                        v2 = append(v2, v33)
                                                                                    }
                                                                                case "hello":
                                                                                    if (91 <= 0) {
                                                                                        v1 += (v32 & -244)
                                                                                        switch fi41 {
                                                                                        case 9223372036854775807:
                                                                                            v2 = append(v2, v33)
                                                                                        case 442:
                                                                                            switch fi41 {
                                                                                            case 42:
                                                                                                switch fi41 {
                                                                                                case 2147483647:
                                                                                                    var v104 float64 = 1.5
                                                                                                    _ = v104
                                                                                                case 78:
                                                                                                    if false {
                                                                                                        var v105 uint = 1
                                                                                                        v73 += (v105 % (((v3 & (v73 & v3))) | 1))
                                                                                                        v106 := GenS9{F0: (v6 << ((888) & 63))}
                                                                                                        _ = v105
                                                                                                        _ = v106
                                                                                                    } else {
                                                                                                        fmt.Println(fi42, v73)
                                                                                                        v40 = (v40 / (((v40 + -560)) | 1))
                                                                                                    }
                                                                                                default:
                                                                                                    var v107 uint32 = 42
                                                                                                    _ = v107
                                                                                                }
                                                                                            case 127:
                                                                                                if false {
                                                                                                    v108 := GenS13{F0: int16(v6)}
                                                                                                    fmt.Println(v33, v73)
                                                                                                    switch v33 {
                                                                                                    case "a":
                                                                                                        fmt.Println(v39)
                                                                                                    case "hello":
                                                                                                        _ = len(v5)
                                                                                                    }
                                                                                                    _ = v108
                                                                                                }
                                                                                                switch (v33 + v33) {
                                                                                                case "":
                                                                                                    v6 = (v6 - (v6 << ((v3) & 63)))
                                                                                                case "a":
                                                                                                    if ((v43 + v33) >= v33) {
                                                                                                        v109 := GenS13{F0: 32767}
                                                                                                        _ = v109
                                                                                                    } else {
                                                                                                        var v110 float32 = 0.1
                                                                                                        v111 := GenS6{F0: v6, F1: 4294967295}
                                                                                                        var v112 float32 = v110
                                                                                                        _ = v110
                                                                                                        _ = v111
                                                                                                        _ = v112
                                                                                                    }
                                                                                                }
                                                                                                var v113 bool = (v40 < int32(v6))
                                                                                                v113 = (!(-970 != 42))
                                                                                                v3 += v73
                                                                                                _ = v113
                                                                                            }
                                                                                        case 1:
                                                                                            fmt.Println(fi42, v39)
                                                                                        }
                                                                                        if len(v4) > 0 {
                                                                                            v4[0] = "a"
                                                                                        }
                                                                                    }
                                                                                case "a":
                                                                                    var v114 int16 = 0
                                                                                    _ = v114
                                                                                default:
                                                                                    v6 += int8(v40)
                                                                                }
                                                                                v115 := GenS3{F0: v33, F1: true}
                                                                                _ = v115
                                                                            case 886:
                                                                                v116 := GenS7{F0: 2147483647}
                                                                                _ = v116
                                                                            }
                                                                            fmt.Println(v40)
                                                                        case "foo bar":
                                                                            if false {
                                                                                fmt.Println(v32, fi42)
                                                                                if true {
                                                                                    fmt.Println(v1, v73, v3)
                                                                                    fmt.Println(v40, v33, v6)
                                                                                }
                                                                            } else {
                                                                                switch ("" + "foo bar") {
                                                                                case "foo bar":
                                                                                    var v117 uint32 = 65535
                                                                                    _ = v117
                                                                                default:
                                                                                    v118 := GenS11{F0: (0.1 != 0.1)}
                                                                                    _ = v118
                                                                                }
                                                                                v33 = v33
                                                                                var v119 int16 = 1
                                                                                _ = v119
                                                                            }
                                                                        case "":
                                                                            v73 = (v73 >> (((v3 & 65535)) & 63))
                                                                        }
                                                                    case 9223372036854775807:
                                                                        switch v33 {
                                                                        case "hello":
                                                                            v32 = v1
                                                                        case "foo bar":
                                                                            v43 = ("a" + (("hello" + (("\n\t" + v43) + "foo bar")) + ""))
                                                                            v33 = "\n\t"
                                                                            if (((v32 & -434) >= (v1 | (v1 ^ (v40 ^ v40)))) || ((0 != 50) || (-2.25 > (0.0 - (float32(((-2.25 + (1e10 - 0.0)) - 3.14159)) + 0.0))))) {
                                                                                var v120 uint64 = 0
                                                                                v121 := GenS14{F0: 255, F1: v6, F2: -1}
                                                                                v122 := GenS15{F0: 1}
                                                                                _ = v120
                                                                                _ = v121
                                                                                _ = v122
                                                                            }
                                                                        case "":
                                                                            if false {
                                                                                if true {
                                                                                    if false {
                                                                                        v123 := GenS16{F0: fi41}
                                                                                        _ = v123
                                                                                    } else {
                                                                                        v124 := GenS12{F0: 0, F1: v43, F2: 750}
                                                                                        _ = v124
                                                                                    }
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = v33
                                                                                    }
                                                                                    switch fi41 {
                                                                                    case 42:
                                                                                        if false {
                                                                                            v125 := GenS7{F0: 620}
                                                                                            v6 = 42
                                                                                            v5 = append(v5, 1)
                                                                                            v126 := GenS5{F0: -756}
                                                                                            v127 := GenS17{F0: 323, F1: (fi41 >> (((v39 | (v73 | v39))) & 63))}
                                                                                            _ = v125
                                                                                            _ = v126
                                                                                            _ = v127
                                                                                        }
                                                                                        switch "a" {
                                                                                        case "a":
                                                                                            v32 += (-(v32))
                                                                                        default:
                                                                                            if true {
                                                                                                v40 = (-(v32))
                                                                                            }
                                                                                        }
                                                                                        v73 = (^(v3))
                                                                                    case -819:
                                                                                        if (!true) {
                                                                                            var v128 float64 = 1e10
                                                                                            _ = v128
                                                                                        }
                                                                                    case 2147483647:
                                                                                        v40 = int32(v6)
                                                                                    }
                                                                                    v129 := GenS17{F0: 65535, F1: fi41}
                                                                                    v2 = append(v2, ("" + ((((v33 + v33) + v33) + ((v33 + v33) + (v43 + (v33 + (("\n\t" + (v33 + "foo bar")) + v43))))) + "hello")))
                                                                                    var v130 string = ("\n\t" + (v33 + v43))
                                                                                    if true {
                                                                                        v131 := GenS12{F0: 196, F1: v43, F2: 65535}
                                                                                        v132 := GenS18{F0: (4294967295 != (v3 >> ((v73) & 63))), F1: ((9223372036854775807 != -398) || true)}
                                                                                        var v133 string = ("\n\t" + (v33 + v43))
                                                                                        _ = v131
                                                                                        _ = v132
                                                                                        _ = v133
                                                                                    } else {
                                                                                        v134 := GenS17{F0: 255, F1: fi42}
                                                                                        if true {
                                                                                            var v135 float32 = (1.5 - (((0.0 - 0.0) * (-2.25 / (1.5 * 0.1))) + 1e10))
                                                                                            _ = v135
                                                                                        } else {
                                                                                            if true {
                                                                                                v5 = append(v5, uint32(v73))
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = 2
                                                                                                }
                                                                                                switch v43 {
                                                                                                case "hello":
                                                                                                    v130 = v130
                                                                                                case "foo bar":
                                                                                                    switch ("hello" + "\n\t") {
                                                                                                    case "a":
                                                                                                        if false {
                                                                                                            if (!true) {
                                                                                                                v2 = append(v2, (v43 + v130))
                                                                                                                if (-2.25 == 1.5) {
                                                                                                                    switch fi41 {
                                                                                                                    case 0:
                                                                                                                        v3 -= (v73 ^ (v3 - 926))
                                                                                                                    case 42:
                                                                                                                        if true {
                                                                                                                            if ((float32((0.1 + 0.0)) - (1.5 / ((1.5 * -2.25) / 1.5))) == -2.25) {
                                                                                                                                if (-2.25 != ((1.5 / (-2.25 + -2.25)) - 1e10)) {
                                                                                                                                    v2 = append(v2, (v43 + v43))
                                                                                                                                    v33 = v43
                                                                                                                                    v136 := GenS17{F0: 255, F1: fi41}
                                                                                                                                    if (!(true && false)) {
                                                                                                                                        fmt.Println(v40)
                                                                                                                                        v73 += (v3 % ((v39) | 1))
                                                                                                                                        v32 -= v1
                                                                                                                                    } else {
                                                                                                                                        v137 := GenS19{F0: v3}
                                                                                                                                        _ = v137
                                                                                                                                    }
                                                                                                                                    v138 := GenS16{F0: fi41}
                                                                                                                                    _ = v136
                                                                                                                                    _ = v138
                                                                                                                                } else {
                                                                                                                                    v139 := GenS20{F0: 127}
                                                                                                                                    _ = v139
                                                                                                                                }
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            v1 += v40
                                                                                                                            if (!(255 > 1)) {
                                                                                                                                var v140 string = (v33 + v33)
                                                                                                                                _ = v140
                                                                                                                            }
                                                                                                                            if false {
                                                                                                                                var v141 float32 = float32(float64(0.0))
                                                                                                                                _ = v141
                                                                                                                            } else {
                                                                                                                                if false {
                                                                                                                                    if len(v4) > 0 {
                                                                                                                                        v4[0] = (v33 + v130)
                                                                                                                                    }
                                                                                                                                    fmt.Println(v43, v1, v32)
                                                                                                                                    v6 -= (v6 | (-(v6)))
                                                                                                                                    var v142 uint16 = 2
                                                                                                                                    var v143 float64 = (1.5 - 3.14159)
                                                                                                                                    _ = v142
                                                                                                                                    _ = v143
                                                                                                                                } else {
                                                                                                                                    if len(v2) > 0 {
                                                                                                                                        v2[0] = "a"
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                var v144 uint8 = genFn1((^(v39)), "", (fi41 ^ (fi41 * 1)))
                                                                                                                                v145 := GenS8{F0: 0, F1: fi42, F2: int64(fi41)}
                                                                                                                                v146 := GenS7{F0: 42}
                                                                                                                                _ = v144
                                                                                                                                _ = v145
                                                                                                                                _ = v146
                                                                                                                            }
                                                                                                                        }
                                                                                                                    case -747:
                                                                                                                        v32 -= v32
                                                                                                                    default:
                                                                                                                        v3 = 42
                                                                                                                    }
                                                                                                                    switch 0 {
                                                                                                                    case 1:
                                                                                                                        var v147 uint8 = 128
                                                                                                                        _ = v147
                                                                                                                    case 32768:
                                                                                                                        var v148 int32 = 0
                                                                                                                        _ = v148
                                                                                                                    }
                                                                                                                    switch fi42 {
                                                                                                                    case 727:
                                                                                                                        if len(v5) > 0 {
                                                                                                                            v5[0] = 42
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v40 -= -2147483648
                                                                                                                    fmt.Println(fi42)
                                                                                                                    switch v33 {
                                                                                                                    case "\n\t":
                                                                                                                        v149 := GenS11{F0: false}
                                                                                                                        _ = v149
                                                                                                                    case "foo bar":
                                                                                                                        var v150 float64 = ((0.1 / 3.14159) - 1e10)
                                                                                                                        _ = v150
                                                                                                                    }
                                                                                                                    var v151 int16 = 373
                                                                                                                    _ = v151
                                                                                                                }
                                                                                                                fmt.Println(fi41, fi42)
                                                                                                            } else {
                                                                                                                var v152 uint = (^(v3))
                                                                                                                _ = v152
                                                                                                            }
                                                                                                            v73 = v73
                                                                                                        } else {
                                                                                                            _ = len(v2)
                                                                                                            if len(v4) > 0 {
                                                                                                                v4[0] = v43
                                                                                                            }
                                                                                                            switch "" {
                                                                                                            case "a":
                                                                                                                fmt.Println(v40, fi41, v33)
                                                                                                                v40 = (v1 << ((4294967295) & 63))
                                                                                                                if false {
                                                                                                                    var v153 int = (fi42 * fi42)
                                                                                                                    _ = v153
                                                                                                                } else {
                                                                                                                    var v154 bool = ((((!((v32 ^ 0) <= (-(v1)))) || true) && true) && false)
                                                                                                                    if ((847 == 4294967295) && ((v154 && ((v6 & (v6 << (((^(v39))) & 63))) >= (v6 * v6))) || v154)) {
                                                                                                                        v6 = v6
                                                                                                                        if false {
                                                                                                                            fmt.Println(v3)
                                                                                                                        } else {
                                                                                                                            if len(v5) > 0 {
                                                                                                                                v5[0] = 931
                                                                                                                            }
                                                                                                                        }
                                                                                                                        v155 := GenS5{F0: 9223372036854775807}
                                                                                                                        _ = v155
                                                                                                                    } else {
                                                                                                                        var v156 int64 = 842
                                                                                                                        _ = v156
                                                                                                                    }
                                                                                                                    _ = v154
                                                                                                                }
                                                                                                                switch ("" + "") {
                                                                                                                case "a":
                                                                                                                    var v157 bool = (!(42 >= 2))
                                                                                                                    _ = v157
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                v158 := GenS21{F0: 100, F1: (^(v32)), F2: 58}
                                                                                                                v39 -= v3
                                                                                                                _ = v158
                                                                                                            default:
                                                                                                                var v159 uint8 = 120
                                                                                                                _ = v159
                                                                                                            }
                                                                                                            fmt.Println(fi42)
                                                                                                            if (false && false) {
                                                                                                                fmt.Println(v39, fi42)
                                                                                                                if (!false) {
                                                                                                                    if (v39 == 255) {
                                                                                                                        var v160 bool = true
                                                                                                                        _ = v160
                                                                                                                    } else {
                                                                                                                        fmt.Println(v130, v6)
                                                                                                                        if (34 == 184) {
                                                                                                                            if false {
                                                                                                                                var v161 int = (fi42 - fi42)
                                                                                                                                var v162 int32 = v32
                                                                                                                                _ = v161
                                                                                                                                _ = v162
                                                                                                                            }
                                                                                                                            v39 = (^(v39))
                                                                                                                            v130 = v33
                                                                                                                        }
                                                                                                                        var v163 uint16 = 255
                                                                                                                        _ = v163
                                                                                                                    }
                                                                                                                    v43 = v33
                                                                                                                    var v164 uint = v39
                                                                                                                    var v165 int = fi41
                                                                                                                    v165 = fi41
                                                                                                                    _ = v164
                                                                                                                    _ = v165
                                                                                                                }
                                                                                                                var v166 int = 2147483647
                                                                                                                var v167 uint8 = 2
                                                                                                                v130 = v43
                                                                                                                v168 := GenS9{F0: (v6 / (((v6 & v6)) | 1))}
                                                                                                                _ = v166
                                                                                                                _ = v167
                                                                                                                _ = v168
                                                                                                            } else {
                                                                                                                var v169 string = "foo bar"
                                                                                                                v39 += (v3 | (v73 - v39))
                                                                                                                var v170 uint64 = 42
                                                                                                                _ = v169
                                                                                                                _ = v170
                                                                                                            }
                                                                                                        }
                                                                                                    case "\n\t":
                                                                                                        v6 += 100
                                                                                                    case "":
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                    default:
                                                                                                        if false {
                                                                                                            fmt.Println(v3, v39, v43)
                                                                                                            if ("hello" < v33) {
                                                                                                                v39 = v3
                                                                                                                v43 = (((((v33 + v130) + (v43 + ("foo bar" + v33))) + v33) + v43) + v130)
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                case "":
                                                                                                    switch -1 {
                                                                                                    case -1:
                                                                                                        if (((65535 >= 2) && true) && false) {
                                                                                                            if ((!(v43 > (v43 + v130))) || false) {
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                            } else {
                                                                                                                v171 := GenS22{F0: true}
                                                                                                                if (true && (1 == -9223372036854775808)) {
                                                                                                                    switch (v33 + v33) {
                                                                                                                    case "hello":
                                                                                                                        switch v43 {
                                                                                                                        case "\n\t":
                                                                                                                            if len(v2) >= 2 {
                                                                                                                                v2 = v2[1:len(v2)]
                                                                                                                            }
                                                                                                                        case "hello":
                                                                                                                            var v172 int8 = 1
                                                                                                                            _ = v172
                                                                                                                        default:
                                                                                                                            var v173 int16 = -32768
                                                                                                                            _ = v173
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch 127 {
                                                                                                                    case 128:
                                                                                                                        v32 = v32
                                                                                                                    default:
                                                                                                                        var v174 uint16 = 0
                                                                                                                        _ = v174
                                                                                                                    }
                                                                                                                    var v175 uint8 = 255
                                                                                                                    _ = v175
                                                                                                                }
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                                _ = v171
                                                                                                            }
                                                                                                            fmt.Println(fi41, fi41, v39)
                                                                                                            var v176 uint64 = 0
                                                                                                            _ = v176
                                                                                                        } else {
                                                                                                            if ((1.5 > 1e10) && (!true)) {
                                                                                                                var v177 uint64 = 773
                                                                                                                fmt.Println(v73, v6, v39)
                                                                                                                v2 = append(v2, (v130 + ("foo bar" + (v43 + "\n\t"))))
                                                                                                                _ = v177
                                                                                                            }
                                                                                                            v39 = 408
                                                                                                            switch (v130 + v33) {
                                                                                                            case "a":
                                                                                                                if false {
                                                                                                                    fmt.Println(v6, v39)
                                                                                                                    var v178 uint16 = 65535
                                                                                                                    var v179 int = (fi41 ^ (fi42 << (((v3 ^ (v3 * v3))) & 63)))
                                                                                                                    _ = v178
                                                                                                                    _ = v179
                                                                                                                }
                                                                                                            case "\n\t":
                                                                                                                v39 = (v3 - (^(v39)))
                                                                                                            case "":
                                                                                                                fmt.Println(v1, v33)
                                                                                                            default:
                                                                                                                var v180 uint64 = 903
                                                                                                                _ = v180
                                                                                                            }
                                                                                                        }
                                                                                                    case 763:
                                                                                                        if len(v5) > 0 {
                                                                                                            v5[0] = 4294967295
                                                                                                        }
                                                                                                    case 2147483647:
                                                                                                        _ = len(v5)
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                v6 = v6
                                                                                                switch -9223372036854775808 {
                                                                                                case 32768:
                                                                                                    fmt.Println(v32, v6)
                                                                                                case 368:
                                                                                                    v5 = append(v5, 0)
                                                                                                case -427:
                                                                                                    v181 := GenS11{F0: (v1 >= v40)}
                                                                                                    _ = v181
                                                                                                default:
                                                                                                    v39 = (v3 + 42)
                                                                                                }
                                                                                                var v182 bool = (!true)
                                                                                                var v183 uint16 = 42
                                                                                                _ = v182
                                                                                                _ = v183
                                                                                            }
                                                                                        }
                                                                                        v6 = v6
                                                                                        _ = v134
                                                                                    }
                                                                                    if false {
                                                                                        fmt.Println(v32)
                                                                                        switch v130 {
                                                                                        case "\n\t":
                                                                                            var v184 uint16 = 0
                                                                                            _ = v184
                                                                                        default:
                                                                                            var v185 float32 = 0.1
                                                                                            _ = v185
                                                                                        }
                                                                                        v40 -= v1
                                                                                    }
                                                                                    _ = v129
                                                                                    _ = v130
                                                                                } else {
                                                                                    if false {
                                                                                        switch fi42 {
                                                                                        case -919:
                                                                                            v73 += v73
                                                                                        case 128:
                                                                                            switch fi42 {
                                                                                            case 128:
                                                                                                v33 = (v43 + ("foo bar" + v43))
                                                                                            case 9223372036854775807:
                                                                                                v1 += 42
                                                                                            }
                                                                                        }
                                                                                        if (true && (!(false && false))) {
                                                                                            v3 = (v73 + v73)
                                                                                            var v186 int8 = v6
                                                                                            switch (fi42 << (((v73 * 18446744073709551615)) & 63)) {
                                                                                            case 32767:
                                                                                                fmt.Println(v3, v32, v39)
                                                                                            case 342:
                                                                                                if ((v73 >> ((4294967295) & 63)) != 987) {
                                                                                                    switch v43 {
                                                                                                    case "foo bar":
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                    case "hello":
                                                                                                        v39 -= v3
                                                                                                    case "\n\t":
                                                                                                        if true {
                                                                                                            v187 := GenS4{F0: -978}
                                                                                                            _ = v187
                                                                                                        } else {
                                                                                                            v188 := GenS23{F0: 2}
                                                                                                            var v189 uint8 = 128
                                                                                                            switch v33 {
                                                                                                            case "hello":
                                                                                                                var v190 uint32 = 42
                                                                                                                _ = v190
                                                                                                            }
                                                                                                            _ = v188
                                                                                                            _ = v189
                                                                                                        }
                                                                                                    }
                                                                                                    switch ((("" + "foo bar") + "a") + v33) {
                                                                                                    case "a":
                                                                                                        v191 := GenS19{F0: v39}
                                                                                                        _ = v191
                                                                                                    case "hello":
                                                                                                        fmt.Println(fi41, v73, v32)
                                                                                                    case "":
                                                                                                        var v192 int64 = 0
                                                                                                        _ = v192
                                                                                                    }
                                                                                                    fmt.Println(v3, fi41)
                                                                                                    fmt.Println(v39)
                                                                                                    v43 = "\n\t"
                                                                                                    v4 = append(v4, "hello")
                                                                                                } else {
                                                                                                    v193 := GenS24{F0: 42, F1: 514}
                                                                                                    _ = v193
                                                                                                }
                                                                                            case -1:
                                                                                                var v194 uint = v39
                                                                                                _ = v194
                                                                                            }
                                                                                            fmt.Println(v1, v43, v39)
                                                                                            var v195 int = (fi41 + 1)
                                                                                            _ = v186
                                                                                            _ = v195
                                                                                        }
                                                                                        v3 = v73
                                                                                    }
                                                                                }
                                                                                fmt.Println(v40, v1)
                                                                                v73 = 18446744073709551615
                                                                            }
                                                                        default:
                                                                            var v196 bool = (v73 > v3)
                                                                            _ = v196
                                                                        }
                                                                    }
                                                                default:
                                                                    if false {
                                                                        fmt.Println(v33)
                                                                    } else {
                                                                        v73 = (^(v3))
                                                                        var v197 int64 = -9223372036854775808
                                                                        _ = v197
                                                                    }
                                                                    switch "a" {
                                                                    case "a":
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    case "foo bar":
                                                                        fmt.Println(v3)
                                                                    case "\n\t":
                                                                        switch ("foo bar" + "a") {
                                                                        case "foo bar":
                                                                            var v198 uint8 = genFn1(323, "a", 570)
                                                                            _ = v198
                                                                        case "\n\t":
                                                                            if true {
                                                                                var v199 string = ("hello" + v33)
                                                                                var v200 float32 = 3.14159
                                                                                v33 = v199
                                                                                switch (((v199 + "a") + "a") + v43) {
                                                                                case "foo bar":
                                                                                    v73 = (^(v73))
                                                                                default:
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                }
                                                                                _ = v199
                                                                                _ = v200
                                                                            }
                                                                            v4 = append(v4, ("" + (v33 + v43)))
                                                                            v39 = (v73 * (v3 ^ 271))
                                                                        }
                                                                    default:
                                                                        switch ("" + ((v43 + "hello") + (((v33 + v33) + "hello") + v43))) {
                                                                        case "hello":
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            v4 = append(v4, (v43 + v33))
                                                                        case "\n\t":
                                                                            v201 := GenS20{F0: uint8(v73)}
                                                                            _ = v201
                                                                        }
                                                                    }
                                                                    v202 := GenS9{F0: v6}
                                                                    _ = v202
                                                                }
                                                            case 9223372036854775807:
                                                                switch (v33 + "\n\t") {
                                                                case "":
                                                                    fmt.Println(v33, v6)
                                                                }
                                                            default:
                                                                v1 = v1
                                                                v203 := GenS14{F0: 253, F1: (v6 >> (((v39 & (v73 << ((v73) & 63)))) & 63)), F2: -1}
                                                                switch (fi41 >> ((42) & 63)) {
                                                                case 9223372036854775807:
                                                                    v204 := GenS23{F0: genFn1(v39, ("foo bar" + v43), fi42)}
                                                                    if (127 == (v6 | v6)) {
                                                                        if (v43 != "a") {
                                                                            v3 = v3
                                                                            if len(v2) > 0 {
                                                                                v2[0] = v33
                                                                            }
                                                                            switch "" {
                                                                            case "hello":
                                                                                var v205 float64 = 0.0
                                                                                _ = v205
                                                                            default:
                                                                                if true {
                                                                                    var v206 uint16 = 255
                                                                                    v39 += 662
                                                                                    if (0 > -1) {
                                                                                        var v207 int16 = 0
                                                                                        if true {
                                                                                            var v208 uint32 = 337
                                                                                            if true {
                                                                                                switch fi42 {
                                                                                                case 246:
                                                                                                    fmt.Println(fi42)
                                                                                                case -582:
                                                                                                    v209 := GenS5{F0: 833}
                                                                                                    _ = v209
                                                                                                default:
                                                                                                    switch (fi41 & (fi42 ^ 1)) {
                                                                                                    case 0:
                                                                                                        var v210 uint32 = (v208 & v208)
                                                                                                        _ = v210
                                                                                                    case -492:
                                                                                                        var v211 int32 = (v32 * (v32 * (v1 - v40)))
                                                                                                        _ = v211
                                                                                                    case -2147483648:
                                                                                                        v212 := GenS3{F0: (v33 + (v33 + "\n\t")), F1: false}
                                                                                                        _ = v212
                                                                                                    }
                                                                                                }
                                                                                                v208 = v208
                                                                                            }
                                                                                            _ = v208
                                                                                        }
                                                                                        v213 := GenS25{F0: (fi41 >> (((v3 & v73)) & 63)), F1: true}
                                                                                        _ = v207
                                                                                        _ = v213
                                                                                    } else {
                                                                                        if (938 == 42) {
                                                                                            var v214 int = -62
                                                                                            switch (("hello" + "") + "a") {
                                                                                            case "":
                                                                                                var v215 int8 = (v6 >> (((v73 ^ v39)) & 63))
                                                                                                _ = v215
                                                                                            case "\n\t":
                                                                                                var v216 float64 = -2.25
                                                                                                _ = v216
                                                                                            case "foo bar":
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = v43
                                                                                                }
                                                                                                var v217 uint8 = 42
                                                                                                fmt.Println(v3)
                                                                                                _ = v217
                                                                                            default:
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = 2
                                                                                                }
                                                                                            }
                                                                                            _ = v214
                                                                                        }
                                                                                    }
                                                                                    _ = v206
                                                                                } else {
                                                                                    var v218 int8 = (v6 % ((43) | 1))
                                                                                    _ = v218
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v219 := GenS17{F0: uint16(v3), F1: fi41}
                                                                        _ = v219
                                                                    }
                                                                    _ = v204
                                                                case 589:
                                                                    fmt.Println(v3)
                                                                    var v220 uint32 = 42
                                                                    switch "a" {
                                                                    case "a":
                                                                        var v221 uint32 = v220
                                                                        _ = v221
                                                                    case "hello":
                                                                        if (("hello" >= ((v33 + v33) + (("\n\t" + (v43 + "hello")) + v33))) || false) {
                                                                            var v222 string = ("\n\t" + ("a" + v33))
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            _ = v222
                                                                        } else {
                                                                            v220 += v220
                                                                            var v223 int8 = (-(v6))
                                                                            if true {
                                                                                var v224 int16 = 54
                                                                                var v225 uint32 = (^(v220))
                                                                                _ = v224
                                                                                _ = v225
                                                                            } else {
                                                                                fmt.Println(v6, v223, v1)
                                                                                var v226 int = fi42
                                                                                var v227 float32 = 0.0
                                                                                _ = v226
                                                                                _ = v227
                                                                            }
                                                                            _ = v223
                                                                        }
                                                                    case "foo bar":
                                                                        v33 = "foo bar"
                                                                    }
                                                                    _ = v220
                                                                case -1:
                                                                    if false {
                                                                        v228 := GenS26{F0: (v1 % ((509) | 1))}
                                                                        _ = v228
                                                                    } else {
                                                                        switch ((v33 + ((v33 + (v33 + v43)) + v33)) + v33) {
                                                                        case "foo bar":
                                                                            v40 -= 1
                                                                        case "":
                                                                            v6 += v6
                                                                        default:
                                                                            fmt.Println(v33, v73, v39)
                                                                        }
                                                                        var v229 int32 = v32
                                                                        _ = v229
                                                                    }
                                                                default:
                                                                    var v230 int16 = 865
                                                                    _ = v230
                                                                }
                                                                _ = v203
                                                            }
                                                            var v231 uint8 = 128
                                                            var v232 int64 = 9223372036854775807
                                                            fmt.Println(v43, v3)
                                                            v73 -= v73
                                                            var v233 uint = v39
                                                            var v234 bool = (v43 == "hello")
                                                            _ = v231
                                                            _ = v232
                                                            _ = v233
                                                            _ = v234
                                                            _ = v73
                                                        }
                                                        v39 -= v39
                                                        _ = v71
                                                    }
                                                    if ((!true) && (-112 != 1)) {
                                                        if true {
                                                            var v235 int = (-(fi42))
                                                            if len(v5) > 0 {
                                                                v5[0] = 326
                                                            }
                                                            _ = v235
                                                        }
                                                        v236 := GenS27{F0: -1, F1: (v1 * (-(v32)))}
                                                        fmt.Println(v43, v3, v43)
                                                        _ = v236
                                                    }
                                                }
                                                var v237 uint8 = genFn1(18446744073709551615, "hello", (fi41 & fi42))
                                                _ = v237
                                            }
                                        }
                                        v6 = int8(v32)
                                        var v238 int32 = (-(v32))
                                        _ = v238
                                    }
                                case "":
                                    v239 := GenS28{F0: v6, F1: 9223372036854775807}
                                    _ = v239
                                }
                            default:
                                var v240 uint64 = 2
                                _ = v240
                            }
                        default:
                            if len(v5) > 0 {
                                v5[0] = 153
                            }
                        }
                        var v241 int64 = -9223372036854775808
                        v33 = ("a" + "a")
                        _ = v241
                        _ = v43
                    }
                    fmt.Println(v39, v1, v40)
                    _ = fi42
                }
                fmt.Println(v32, v3)
                var v242 int = (-(fi41))
                _ = fi41
                _ = v242
            }
            v33 = v33
        }
        v32 = (-(v40))
        _ = v40
    }
    var v243 float32 = (-2.25 + (-2.25 * float32(-2.25)))
    fmt.Println(v33)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v243:", v243)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v243
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
