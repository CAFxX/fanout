package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS7 struct { F0 int; F1 uint64; F2 uint16 }
type GenS8 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS9 struct { F0 uint; F1 int16 }
type GenS10 struct { F0 int64; F1 int64; F2 uint8 }
type GenS11 struct { F0 int64; F1 int16; F2 string }
type GenS12 struct { F0 uint8; F1 int16 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint8; F1 int; F2 bool }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 int32; F1 uint32; F2 uint }
type GenS19 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS20 struct { F0 int }
type GenS21 struct { F0 int32; F1 int32 }
type GenS22 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS23 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS24 struct { F0 int8; F1 int32 }
type GenS25 struct { F0 uint16 }
type GenS26 struct { F0 bool }
type GenS27 struct { F0 uint8 }
type GenS28 struct { F0 uint64 }
type GenS29 struct { F0 int16; F1 uint16 }
type GenS30 struct { F0 uint64 }
type GenS31 struct { F0 uint64 }
type GenS32 struct { F0 string }
type GenS33 struct { F0 int64 }
type GenS34 struct { F0 uint8 }
type GenS35 struct { F0 uint32; F1 int16 }
type GenS36 struct { F0 int64; F1 uint }
type GenS37 struct { F0 int32; F1 uint64; F2 uint }
type GenS38 struct { F0 int8 }
type GenS39 struct { F0 uint8; F1 uint64 }
type GenS40 struct { F0 int64; F1 int }
type GenS41 struct { F0 int16 }
type GenS42 struct { F0 int64; F1 uint8 }
type GenS43 struct { F0 int }
type GenS44 struct { F0 int16; F1 uint32 }
type GenS45 struct { F0 uint; F1 uint64; F2 bool }
type GenS46 struct { F0 uint64; F1 uint8; F2 int }
type GenS47 struct { F0 int16 }
type GenS48 struct { F0 string; F1 int16; F2 int64 }
type GenS49 struct { F0 int; F1 uint }
type GenS50 struct { F0 int }
type GenS51 struct { F0 int16; F1 int32 }
type GenS52 struct { F0 uint }
type GenS53 struct { F0 int16; F1 int8; F2 int16 }
type GenS54 struct { F0 uint64; F1 int8; F2 bool }
type GenS55 struct { F0 int64 }
type GenS56 struct { F0 uint16; F1 uint; F2 int64 }
type GenS57 struct { F0 bool; F1 string }
type GenS58 struct { F0 int16 }
type GenS59 struct { F0 int32; F1 bool; F2 uint16 }
type GenS60 struct { F0 bool; F1 int32 }
type GenS61 struct { F0 int16; F1 int32; F2 int64 }
type GenS62 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS63 struct { F0 uint8; F1 bool; F2 bool }
type GenS64 struct { F0 int8 }
type GenS65 struct { F0 uint16; F1 int64 }
type GenS66 struct { F0 string; F1 int32 }
type GenS67 struct { F0 uint64; F1 int64 }
type GenS68 struct { F0 bool }
type GenS69 struct { F0 int32; F1 int8; F2 uint8 }
type GenS70 struct { F0 int64; F1 int; F2 int8 }
type GenS71 struct { F0 uint8; F1 bool }
type GenS72 struct { F0 uint64; F1 bool }
type GenS73 struct { F0 uint64 }
type GenS74 struct { F0 int32; F1 string; F2 uint8 }
type GenS75 struct { F0 int; F1 int16; F2 uint16 }
type GenS76 struct { F0 uint16; F1 int8 }
type GenS77 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint16, p1 int16, p2 string) uint16 {
    v5 := []int{}
    for fi6 := 0; fi6 < 13; fi6++ {
        if (!(1000 > 0)) {
            for fi7 := 0; fi7 < 5; fi7++ {
                switch fi7 {
                case 0:
                    p0 = (p0 ^ (p0 % (((p0 << ((255) & 63))) | 1)))
                    var v8 int = fi7
                    switch "\n\t" {
                    case "":
                        if (false && true) {
                            if true {
                                fmt.Println(v8)
                                v9 := GenS1{F0: "foo bar", F1: 318}
                                v10 := GenS1{F0: "hello", F1: 2147483647}
                                switch "a" {
                                case "foo bar":
                                    v8 = (-(v8))
                                    if false {
                                        switch ((p2 + "\n\t") + p2) {
                                        case "hello":
                                            v5 = append(v5, 1)
                                        case "":
                                            if ((^(p0)) >= p0) {
                                                if true {
                                                    var v11 int32 = 2147483647
                                                    if true {
                                                        var v12 uint64 = 0
                                                        if (4294967295 >= (v12 << ((1) & 63))) {
                                                            fmt.Println(v11, v8)
                                                            var v13 uint = uint(v12)
                                                            _ = v13
                                                        } else {
                                                            if len(v5) > 0 {
                                                                v5[0] = 127
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        }
                                                        _ = v12
                                                    } else {
                                                        p2 = p2
                                                        if len(v5) > 0 {
                                                            v5[0] = 1
                                                        }
                                                    }
                                                    switch 2147483647 {
                                                    case 1:
                                                        var v14 int8 = 100
                                                        _ = v14
                                                    case -2147483648:
                                                        var v15 string = p2
                                                        _ = v15
                                                    case 247:
                                                        v5 = append(v5, -2147483648)
                                                        v16 := GenS1{F0: p2, F1: 348}
                                                        v17 := GenS1{F0: p2, F1: 2147483647}
                                                        _ = v16
                                                        _ = v17
                                                    default:
                                                        var v18 float64 = 1e10
                                                        _ = v18
                                                    }
                                                    p2 = ((p2 + p2) + ((p2 + p2) + p2))
                                                    _ = v11
                                                } else {
                                                    fmt.Println(fi6, p1)
                                                    var v19 float64 = 0.1
                                                    _ = v19
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = (-(fi7))
                                                }
                                                if true {
                                                    v20 := GenS2{F0: -82, F1: 32767}
                                                    var v21 int16 = p1
                                                    if (!false) {
                                                        switch fi7 {
                                                        case 9223372036854775807:
                                                            p1 = 42
                                                        case 2147483647:
                                                            var v22 uint = 284
                                                            _ = v22
                                                        case -9223372036854775808:
                                                            var v23 string = ((p2 + (("hello" + p2) + (p2 + ("a" + "")))) + (p2 + ""))
                                                            _ = v23
                                                        }
                                                        if true {
                                                            v24 := GenS1{F0: "hello", F1: 2147483648}
                                                            v25 := GenS3{F0: (^(p0)), F1: v8}
                                                            v5 = append(v5, 4294967295)
                                                            _ = v24
                                                            _ = v25
                                                        }
                                                        var v26 int8 = -69
                                                        switch p2 {
                                                        case "a":
                                                            fmt.Println(fi6)
                                                        case "\n\t":
                                                            v26 = v26
                                                        default:
                                                            if true {
                                                                switch fi6 {
                                                                case 128:
                                                                    if (!(619 <= 255)) {
                                                                        p1 = p1
                                                                        p0 = 1
                                                                        v21 = p1
                                                                    } else {
                                                                        switch fi7 {
                                                                        case -2147483648:
                                                                            var v27 int = fi6
                                                                            _ = v27
                                                                        default:
                                                                            fmt.Println(fi6)
                                                                        }
                                                                        _ = len(v5)
                                                                    }
                                                                }
                                                                var v28 uint16 = 2
                                                                var v29 float32 = 3.14159
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                        }
                                                        var v30 int64 = 318
                                                        v8 -= (fi7 - (fi6 >> ((0) & 63)))
                                                        fmt.Println(p0, v26, p0)
                                                        if (!(1 < 126)) {
                                                            v5 = append(v5, fi7)
                                                            v8 = v8
                                                            fmt.Println(p0, fi7, p1)
                                                        } else {
                                                            if false {
                                                                v5 = append(v5, (v8 | int(p0)))
                                                                if true {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    var v31 uint = 499
                                                                    if (!(!true)) {
                                                                        v32 := GenS1{F0: (((p2 + p2) + ((p2 + p2) + "\n\t")) + ""), F1: 1}
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        _ = v32
                                                                    } else {
                                                                        var v33 int8 = v26
                                                                        _ = v33
                                                                    }
                                                                    _ = v31
                                                                }
                                                                var v34 int32 = -2147483648
                                                                _ = v34
                                                            }
                                                        }
                                                        _ = v26
                                                        _ = v30
                                                    } else {
                                                        if true {
                                                            var v35 float64 = 1.5
                                                            fmt.Println(p1, fi7)
                                                            _ = v35
                                                        }
                                                        switch 671 {
                                                        case 1:
                                                            v21 += (-(v21))
                                                        default:
                                                            var v36 uint64 = 4294967295
                                                            _ = v36
                                                        }
                                                        if true {
                                                            v8 = v8
                                                        }
                                                    }
                                                    _ = v20
                                                    _ = v21
                                                } else {
                                                    if false {
                                                        if (-32768 != 1000) {
                                                            var v37 bool = false
                                                            if (-58 > 1) {
                                                                var v38 uint32 = 65535
                                                                var v39 uint64 = 499
                                                                _ = v38
                                                                _ = v39
                                                            } else {
                                                                p2 = ("" + "")
                                                                var v40 uint = 43
                                                                _ = v40
                                                            }
                                                            fmt.Println(v8)
                                                            _ = v37
                                                        }
                                                    } else {
                                                        if (61 <= 103) {
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            switch "hello" {
                                                            case "a":
                                                                if false {
                                                                    fmt.Println(p0, p1)
                                                                    if len(v5) > 0 {
                                                                        v5[0] = -196
                                                                    }
                                                                    v8 -= (v8 % ((int(p1)) | 1))
                                                                } else {
                                                                    if (0 == 65535) {
                                                                        var v41 uint32 = 42
                                                                        _ = v41
                                                                    }
                                                                    if (42 == 751) {
                                                                        switch (v8 * (fi7 & 127)) {
                                                                        case 32767:
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        case -2147483648:
                                                                            var v42 int = fi6
                                                                            _ = v42
                                                                        case -988:
                                                                            if (!false) {
                                                                                var v43 uint16 = 538
                                                                                v44 := GenS4{F0: 42, F1: 1}
                                                                                v8 -= 0
                                                                                _ = v43
                                                                                _ = v44
                                                                            } else {
                                                                                v45 := GenS5{F0: -1}
                                                                                if ((!((1.5 <= 3.14159) && (!false))) && true) {
                                                                                    var v46 uint64 = 4294967295
                                                                                    _ = v46
                                                                                } else {
                                                                                    var v47 int32 = 42
                                                                                    var v48 int32 = v47
                                                                                    fmt.Println(fi6, v47, fi6)
                                                                                    if false {
                                                                                        v49 := GenS4{F0: (v47 + (-(v47))), F1: uint64(p0)}
                                                                                        _ = v49
                                                                                    } else {
                                                                                        fmt.Println(fi7, v8)
                                                                                        fmt.Println(p1)
                                                                                        v50 := GenS5{F0: -520}
                                                                                        _ = v50
                                                                                    }
                                                                                    var v51 int64 = -689
                                                                                    var v52 int16 = -776
                                                                                    _ = v47
                                                                                    _ = v48
                                                                                    _ = v51
                                                                                    _ = v52
                                                                                }
                                                                                _ = len(v5)
                                                                                _ = v45
                                                                            }
                                                                        }
                                                                        if ((p1 >> ((42) & 63)) != p1) {
                                                                            var v53 uint32 = 214
                                                                            var v54 float64 = 1e10
                                                                            _ = v53
                                                                            _ = v54
                                                                        } else {
                                                                            p1 += (p1 | p1)
                                                                            p1 -= (p1 & p1)
                                                                            var v55 int = (fi7 + fi7)
                                                                            _ = v55
                                                                        }
                                                                        if false {
                                                                            var v56 int64 = -130
                                                                            _ = v56
                                                                        }
                                                                        if (v8 > (v8 & -518)) {
                                                                            _ = len(v5)
                                                                        } else {
                                                                            if (128 < 2) {
                                                                                fmt.Println(p2, p2)
                                                                                fmt.Println(p0, fi6)
                                                                                fmt.Println(fi6, p1)
                                                                            } else {
                                                                                var v57 int16 = 1
                                                                                if (!true) {
                                                                                    fmt.Println(p1)
                                                                                    var v58 uint16 = 267
                                                                                    if (true && true) {
                                                                                        fmt.Println(v57)
                                                                                        if ((true || ((!(!(147 < uint8(p0)))) || ((0.0 * 3.14159) <= 0.1))) && ((!(p2 <= "a")) && (1 == 42))) {
                                                                                            p1 += (p1 & 1000)
                                                                                            v59 := GenS4{F0: 0, F1: 18446744073709551615}
                                                                                            _ = v59
                                                                                        }
                                                                                        var v60 float32 = 1.5
                                                                                        if ((!false) && false) {
                                                                                            v61 := GenS6{F0: 153, F1: 516, F2: 0}
                                                                                            v5 = append(v5, v8)
                                                                                            _ = v61
                                                                                        }
                                                                                        v62 := GenS7{F0: (-(v8)), F1: 18446744073709551615, F2: p0}
                                                                                        v8 = v8
                                                                                        if true {
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = fi7
                                                                                            }
                                                                                            switch p2 {
                                                                                            case "\n\t":
                                                                                                if ((!(1 != (p0 << ((42) & 63)))) && (float32((float64(float32(1e10)) + 0.1)) < 1e10)) {
                                                                                                    fmt.Println(fi7, p1, v60)
                                                                                                }
                                                                                            default:
                                                                                                v63 := GenS5{F0: -1}
                                                                                                _ = v63
                                                                                            }
                                                                                            v64 := GenS8{F0: p0, F1: 127, F2: 0}
                                                                                            _ = v64
                                                                                        } else {
                                                                                            v65 := GenS1{F0: p2, F1: 2147483647}
                                                                                            var v66 int = v8
                                                                                            _ = v65
                                                                                            _ = v66
                                                                                        }
                                                                                        _ = v60
                                                                                        _ = v62
                                                                                    } else {
                                                                                        fmt.Println(p1, p1, v8)
                                                                                        if (344 < 179) {
                                                                                            switch (fi7 + fi6) {
                                                                                            case 32768:
                                                                                                p1 += p1
                                                                                            case 128:
                                                                                                v67 := GenS9{F0: 42, F1: int16(v8)}
                                                                                                if false {
                                                                                                    var v68 int8 = -128
                                                                                                    p0 = p0
                                                                                                    v8 = fi6
                                                                                                    _ = v68
                                                                                                } else {
                                                                                                    v5 = append(v5, fi7)
                                                                                                    var v69 uint64 = 0
                                                                                                    _ = v69
                                                                                                }
                                                                                                _ = v67
                                                                                            case 32767:
                                                                                                var v70 int32 = -1
                                                                                                _ = v70
                                                                                            default:
                                                                                                fmt.Println(p2, fi6)
                                                                                            }
                                                                                            if false {
                                                                                                var v71 uint = 65535
                                                                                                var v72 int8 = 1
                                                                                                if true {
                                                                                                    v73 := GenS4{F0: 422, F1: 2}
                                                                                                    _ = v73
                                                                                                }
                                                                                                _ = v71
                                                                                                _ = v72
                                                                                            }
                                                                                            v74 := GenS10{F0: 1, F1: int64(p1), F2: 2}
                                                                                            _ = v74
                                                                                        }
                                                                                        v75 := GenS4{F0: 0, F1: 1}
                                                                                        _ = v75
                                                                                    }
                                                                                    _ = v58
                                                                                } else {
                                                                                    var v76 int16 = (v57 ^ (v57 | (v57 << ((2) & 63))))
                                                                                    fmt.Println(v76)
                                                                                    v57 = v76
                                                                                    _ = v76
                                                                                }
                                                                                _ = v57
                                                                            }
                                                                            var v77 int16 = (p1 << ((182) & 63))
                                                                            if ("\n\t" > p2) {
                                                                                var v78 int16 = p1
                                                                                _ = v78
                                                                            }
                                                                            _ = v77
                                                                        }
                                                                    }
                                                                    p0 = (p0 - p0)
                                                                    var v79 uint64 = 42
                                                                    _ = v79
                                                                }
                                                                var v80 uint32 = 65535
                                                                fmt.Println(v8, fi7)
                                                                _ = v80
                                                            case "hello":
                                                                fmt.Println(fi7)
                                                            default:
                                                                v81 := GenS11{F0: 0, F1: (^(p1)), F2: p2}
                                                                _ = v81
                                                            }
                                                        } else {
                                                            fmt.Println(v8)
                                                        }
                                                        p2 = (p2 + p2)
                                                    }
                                                    if false {
                                                        p2 = "foo bar"
                                                        var v82 uint64 = 2
                                                        _ = v82
                                                    }
                                                    v83 := GenS5{F0: 1}
                                                    _ = v83
                                                }
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            } else {
                                                v84 := GenS4{F0: -1, F1: 42}
                                                p0 = p0
                                                fmt.Println(fi6)
                                                var v85 int = fi6
                                                fmt.Println(p2)
                                                _ = v84
                                                _ = v85
                                            }
                                        default:
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        v86 := GenS12{F0: 123, F1: p1}
                                        _ = len(v5)
                                        _ = v86
                                    } else {
                                        fmt.Println(fi7)
                                        if false {
                                            v87 := GenS5{F0: -9223372036854775808}
                                            var v88 int16 = p1
                                            _ = v87
                                            _ = v88
                                        }
                                        v8 += -2147483648
                                        fmt.Println(p0)
                                        fmt.Println(fi7, p1)
                                    }
                                case "hello":
                                    p2 = p2
                                case "\n\t":
                                    var v89 float32 = 0.0
                                    v8 = -259
                                    switch "a" {
                                    case "\n\t":
                                        v90 := GenS12{F0: 102, F1: (p1 >> ((1) & 63))}
                                        _ = v90
                                    case "hello":
                                        fmt.Println(p0)
                                    }
                                    _ = v89
                                default:
                                    var v91 uint64 = 2
                                    _ = v91
                                }
                                v92 := GenS10{F0: 0, F1: -974, F2: 5}
                                switch fi6 {
                                case 476:
                                    if (false && (250 >= 0)) {
                                        var v93 int16 = p1
                                        _ = v93
                                    }
                                }
                                switch p2 {
                                case "foo bar":
                                    if false {
                                        var v94 uint8 = 9
                                        _ = v94
                                    } else {
                                        v8 += -117
                                    }
                                    p2 = ("foo bar" + "\n\t")
                                case "":
                                    v95 := GenS13{F0: 2147483647}
                                    _ = v95
                                default:
                                    if (624 != 42) {
                                        switch "a" {
                                        case "\n\t":
                                            if true {
                                                switch (p2 + p2) {
                                                case "a":
                                                    var v96 int = fi7
                                                    if false {
                                                        fmt.Println(v96, v8, fi7)
                                                        switch 32768 {
                                                        case 9223372036854775807:
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        }
                                                        v96 += fi6
                                                    } else {
                                                        var v97 uint8 = 127
                                                        if (9223372036854775807 < 0) {
                                                            v96 -= (v8 >> ((2) & 63))
                                                            if ((-34 != (v96 << ((523) & 63))) || (p1 <= (p1 << ((1) & 63)))) {
                                                                switch (-(fi7)) {
                                                                case 42:
                                                                    if len(v5) > 0 {
                                                                        v5[0] = v8
                                                                    }
                                                                case -1:
                                                                    var v98 float64 = 1e10
                                                                    _ = v98
                                                                case 0:
                                                                    var v99 uint64 = 430
                                                                    _ = v99
                                                                default:
                                                                    v8 -= v96
                                                                }
                                                                var v100 int8 = 100
                                                                var v101 float32 = float32(3.14159)
                                                                _ = v100
                                                                _ = v101
                                                            } else {
                                                                v102 := GenS6{F0: v97, F1: 42, F2: 42}
                                                                _ = v102
                                                            }
                                                            if true {
                                                                fmt.Println(p0, v8)
                                                            }
                                                            v96 += 9223372036854775807
                                                            _ = len(v5)
                                                        } else {
                                                            var v103 int = fi7
                                                            var v104 int64 = 1
                                                            _ = v103
                                                            _ = v104
                                                        }
                                                        switch v8 {
                                                        case -9223372036854775808:
                                                            switch p2 {
                                                            case "\n\t":
                                                                v97 = v97
                                                            case "a":
                                                                switch (-(v96)) {
                                                                case 128:
                                                                    var v105 float32 = 0.0
                                                                    _ = v105
                                                                }
                                                            case "hello":
                                                                if len(v5) > 0 {
                                                                    v5[0] = (v96 * fi6)
                                                                }
                                                            }
                                                        case -2147483648:
                                                            fmt.Println(v96, v97)
                                                        case -44:
                                                            if ((p0 >> ((255) & 63)) < (p0 | p0)) {
                                                                if true {
                                                                    p0 = 2
                                                                    _ = len(v5)
                                                                } else {
                                                                    v96 = 144
                                                                    var v106 string = p2
                                                                    p0 = p0
                                                                    fmt.Println(fi6, v8, p2)
                                                                    if len(v5) > 0 {
                                                                        v5[0] = (v8 % ((fi7) | 1))
                                                                    }
                                                                    _ = v106
                                                                }
                                                                var v107 int32 = -1
                                                                _ = v107
                                                            } else {
                                                                fmt.Println(v96, fi6, fi6)
                                                            }
                                                        default:
                                                            var v108 float64 = (1.5 / ((-2.25 / 0.1) - 0.0))
                                                            _ = v108
                                                        }
                                                        var v109 float32 = -2.25
                                                        _ = v109
                                                        _ = v97
                                                    }
                                                    if (!false) {
                                                        if (false || false) {
                                                            if (v8 == (v8 + v8)) {
                                                                v110 := GenS1{F0: ((p2 + p2) + p2), F1: 647}
                                                                _ = v110
                                                            } else {
                                                                v111 := GenS14{F0: 123, F1: v8, F2: false}
                                                                if true {
                                                                    if true {
                                                                        var v112 uint = 790
                                                                        switch v8 {
                                                                        case -1:
                                                                            v96 -= fi6
                                                                        case 32767:
                                                                            v113 := GenS15{F0: 9223372036854775807}
                                                                            _ = v113
                                                                        case 32768:
                                                                            var v114 uint8 = 50
                                                                            _ = v114
                                                                        }
                                                                        _ = v112
                                                                    } else {
                                                                        if (false || (100 == int8(p1))) {
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                var v115 int64 = 198
                                                                                _ = v115
                                                                            case "hello":
                                                                                p0 = p0
                                                                            default:
                                                                                if (42 >= 2) {
                                                                                    var v116 int8 = 100
                                                                                    p1 -= p1
                                                                                    if false {
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                        p2 = p2
                                                                                    } else {
                                                                                        v116 = 100
                                                                                    }
                                                                                    _ = v116
                                                                                }
                                                                            }
                                                                            switch p2 {
                                                                            case "foo bar":
                                                                                var v117 uint64 = 1
                                                                                _ = v117
                                                                            }
                                                                            v118 := GenS13{F0: -2147483648}
                                                                            _ = v118
                                                                        } else {
                                                                            var v119 int16 = (p1 / ((p1) | 1))
                                                                            _ = v119
                                                                        }
                                                                        if false {
                                                                            var v120 bool = true
                                                                            _ = v120
                                                                        } else {
                                                                            if len(v5) > 0 {
                                                                                v5[0] = v96
                                                                            }
                                                                        }
                                                                    }
                                                                    if (56 != 42) {
                                                                        var v121 uint8 = 237
                                                                        if true {
                                                                            if (p0 >= uint16(v121)) {
                                                                                if len(v5) > 0 {
                                                                                    v5[0] = v8
                                                                                }
                                                                                switch p2 {
                                                                                case "foo bar":
                                                                                    v5 = append(v5, (v96 | fi7))
                                                                                case "hello":
                                                                                    var v122 uint16 = 1
                                                                                    _ = v122
                                                                                case "":
                                                                                    switch p2 {
                                                                                    case "foo bar":
                                                                                        var v123 float64 = 3.14159
                                                                                        _ = v123
                                                                                    case "hello":
                                                                                        var v124 int64 = 0
                                                                                        _ = v124
                                                                                    case "":
                                                                                        p2 = "hello"
                                                                                    }
                                                                                }
                                                                                var v125 int16 = 42
                                                                                _ = v125
                                                                            } else {
                                                                                switch p2 {
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        var v126 uint16 = (p0 | p0)
                                                                                        if (!true) {
                                                                                            switch p2 {
                                                                                            case "\n\t":
                                                                                                p1 -= 1
                                                                                            case "":
                                                                                                var v127 float32 = 1.5
                                                                                                _ = v127
                                                                                            default:
                                                                                                v128 := GenS1{F0: (p2 + p2), F1: 2}
                                                                                                _ = v128
                                                                                            }
                                                                                            if false {
                                                                                                v129 := GenS5{F0: 0}
                                                                                                var v130 int32 = 0
                                                                                                _ = v129
                                                                                                _ = v130
                                                                                            }
                                                                                            p1 = (p1 / ((-538) | 1))
                                                                                        }
                                                                                        fmt.Println(fi7, fi6)
                                                                                        _ = v126
                                                                                    } else {
                                                                                        switch (fi6 + fi7) {
                                                                                        case 0:
                                                                                            fmt.Println(fi6, v121)
                                                                                        }
                                                                                        var v131 uint16 = p0
                                                                                        p1 += p1
                                                                                        _ = v131
                                                                                    }
                                                                                }
                                                                            }
                                                                            v8 += (fi7 >> ((570) & 63))
                                                                        }
                                                                        v121 -= (^(v121))
                                                                        _ = v121
                                                                    }
                                                                }
                                                                _ = v111
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        } else {
                                                            var v132 int32 = 1
                                                            if len(v5) > 0 {
                                                                v5[0] = 32767
                                                            }
                                                            _ = v132
                                                        }
                                                    } else {
                                                        switch ("" + ("a" + (p2 + p2))) {
                                                        case "hello":
                                                            if true {
                                                                _ = len(v5)
                                                                _ = len(v5)
                                                                if false {
                                                                    if false {
                                                                        var v133 int32 = 42
                                                                        _ = v133
                                                                    } else {
                                                                        var v134 float64 = 3.14159
                                                                        if true {
                                                                            p2 = p2
                                                                            var v135 int32 = -739
                                                                            if (false || (670 == uint64(p0))) {
                                                                                fmt.Println(v135)
                                                                                if false {
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        var v136 uint32 = 42
                                                                                        _ = v136
                                                                                    case "hello":
                                                                                        if (p1 < p1) {
                                                                                            if (false && false) {
                                                                                                v137 := GenS5{F0: 1}
                                                                                                if (!true) {
                                                                                                    var v138 int64 = 0
                                                                                                    _ = v138
                                                                                                }
                                                                                                if false {
                                                                                                    var v139 int8 = 127
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = v8
                                                                                                    }
                                                                                                    _ = v139
                                                                                                }
                                                                                                _ = v137
                                                                                            } else {
                                                                                                v134 = (((-2.25 + (0.0 * (v134 - 3.14159))) + ((3.14159 * v134) / -2.25)) - v134)
                                                                                            }
                                                                                            switch (p2 + "") {
                                                                                            case "foo bar":
                                                                                                switch "hello" {
                                                                                                case "foo bar":
                                                                                                    v96 += fi6
                                                                                                case "":
                                                                                                    var v140 uint16 = p0
                                                                                                    _ = v140
                                                                                                case "a":
                                                                                                    var v141 int32 = 42
                                                                                                    _ = v141
                                                                                                }
                                                                                            case "a":
                                                                                                var v142 uint8 = 127
                                                                                                _ = v142
                                                                                            }
                                                                                            fmt.Println(p1, fi7, v134)
                                                                                        } else {
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = fi7
                                                                                            }
                                                                                            if ((fi6 ^ 2147483647) < (fi7 % ((int(p1)) | 1))) {
                                                                                                v143 := GenS11{F0: 0, F1: p1, F2: p2}
                                                                                                if true {
                                                                                                    p1 += (p1 & (p1 << ((18446744073709551615) & 63)))
                                                                                                    var v144 string = (p2 + p2)
                                                                                                    v8 = v8
                                                                                                    v145 := GenS2{F0: 1, F1: p0}
                                                                                                    var v146 bool = true
                                                                                                    _ = v144
                                                                                                    _ = v145
                                                                                                    _ = v146
                                                                                                }
                                                                                                _ = v143
                                                                                            } else {
                                                                                                if true {
                                                                                                    var v147 uint32 = 65535
                                                                                                    _ = v147
                                                                                                } else {
                                                                                                    var v148 int64 = 42
                                                                                                    _ = v148
                                                                                                }
                                                                                            }
                                                                                            switch v8 {
                                                                                            case 1:
                                                                                                if (false || true) {
                                                                                                    fmt.Println(v8, v135)
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                } else {
                                                                                                    _ = len(v5)
                                                                                                    if (1 > 65535) {
                                                                                                        if len(v5) > 0 {
                                                                                                            v5[0] = (v8 >> ((255) & 63))
                                                                                                        }
                                                                                                        v149 := GenS6{F0: 0, F1: 4294967295, F2: 2147483648}
                                                                                                        _ = v149
                                                                                                    }
                                                                                                    v96 += v96
                                                                                                }
                                                                                            default:
                                                                                                switch "" {
                                                                                                case "foo bar":
                                                                                                    var v150 uint8 = 1
                                                                                                    _ = v150
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    case "foo bar":
                                                                                        _ = len(v5)
                                                                                        _ = len(v5)
                                                                                        switch p2 {
                                                                                        case "\n\t":
                                                                                            if false {
                                                                                                v151 := GenS16{F0: 255}
                                                                                                _ = v151
                                                                                            } else {
                                                                                                fmt.Println(v134, v134, v96)
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        switch p2 {
                                                                                        case "a":
                                                                                            switch (v8 * (v8 ^ v8)) {
                                                                                            case 42:
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = 9223372036854775807
                                                                                                }
                                                                                            default:
                                                                                                var v152 uint64 = 18446744073709551615
                                                                                                _ = v152
                                                                                            }
                                                                                        case "hello":
                                                                                            v134 = 0.1
                                                                                        default:
                                                                                            if (128 <= 2) {
                                                                                                if (false && true) {
                                                                                                    v153 := GenS13{F0: (v135 << ((65535) & 63))}
                                                                                                    var v154 uint32 = 106
                                                                                                    _ = v153
                                                                                                    _ = v154
                                                                                                }
                                                                                                p2 = p2
                                                                                                if false {
                                                                                                    var v155 int = fi7
                                                                                                    p0 += 653
                                                                                                    _ = v155
                                                                                                } else {
                                                                                                    v96 = 127
                                                                                                    if (((0 >= 9223372036854775807) || true) || true) {
                                                                                                        v156 := GenS14{F0: 56, F1: (^(v8)), F2: ((v135 * 42) < (v135 >> ((42) & 63)))}
                                                                                                        _ = v156
                                                                                                    } else {
                                                                                                        v157 := GenS6{F0: 2, F1: 689, F2: 1}
                                                                                                        if false {
                                                                                                            v8 += (v96 << ((4294967295) & 63))
                                                                                                            v135 += (v135 / ((2147483647) | 1))
                                                                                                        }
                                                                                                        var v158 uint8 = 127
                                                                                                        _ = v157
                                                                                                        _ = v158
                                                                                                    }
                                                                                                    var v159 float64 = 0.1
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = fi6
                                                                                                    }
                                                                                                    _ = v159
                                                                                                }
                                                                                            } else {
                                                                                                if false {
                                                                                                    v160 := GenS17{F0: fi7}
                                                                                                    if ((true || false) || (100 > -1)) {
                                                                                                        v161 := GenS17{F0: 246}
                                                                                                        var v162 uint64 = 352
                                                                                                        _ = v161
                                                                                                        _ = v162
                                                                                                    } else {
                                                                                                        fmt.Println(fi7, p2)
                                                                                                    }
                                                                                                    _ = v160
                                                                                                }
                                                                                                if true {
                                                                                                    switch p2 {
                                                                                                    case "a":
                                                                                                        fmt.Println(v134, v135, v135)
                                                                                                        fmt.Println(p2, v135, p0)
                                                                                                        v163 := GenS10{F0: 0, F1: 0, F2: 2}
                                                                                                        _ = v163
                                                                                                    default:
                                                                                                        fmt.Println(v135)
                                                                                                    }
                                                                                                    var v164 int64 = 9223372036854775807
                                                                                                    _ = v164
                                                                                                } else {
                                                                                                    var v165 int16 = p1
                                                                                                    _ = v165
                                                                                                }
                                                                                                p2 = p2
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v5 = append(v5, (fi7 + (-(v96))))
                                                                                    var v166 uint8 = 128
                                                                                    _ = v166
                                                                                } else {
                                                                                    if ((!true) && (!((-(v135)) < 1))) {
                                                                                        if (42 > 0) {
                                                                                            var v167 bool = ((v134 + 1e10) < (v134 * v134))
                                                                                            _ = v167
                                                                                        }
                                                                                        if true {
                                                                                            var v168 uint32 = 4294967295
                                                                                            p1 = p1
                                                                                            switch 1 {
                                                                                            case -840:
                                                                                                if false {
                                                                                                    switch v8 {
                                                                                                    case -1:
                                                                                                        var v169 int16 = (p1 << ((42) & 63))
                                                                                                        _ = v169
                                                                                                    }
                                                                                                } else {
                                                                                                    var v170 uint64 = 657
                                                                                                    v5 = append(v5, (v96 & v96))
                                                                                                    _ = v170
                                                                                                }
                                                                                            case -2147483648:
                                                                                                fmt.Println(v134)
                                                                                            default:
                                                                                                var v171 bool = (458 < 4294967295)
                                                                                                _ = v171
                                                                                            }
                                                                                            _ = v168
                                                                                        } else {
                                                                                            switch fi6 {
                                                                                            case 799:
                                                                                                var v172 float32 = ((3.14159 - -2.25) * 0.0)
                                                                                                _ = v172
                                                                                            case 128:
                                                                                                var v173 uint8 = 1
                                                                                                _ = v173
                                                                                            }
                                                                                            switch (v8 ^ v8) {
                                                                                            case -9223372036854775808:
                                                                                                _ = len(v5)
                                                                                            case 9223372036854775807:
                                                                                                p2 = (p2 + p2)
                                                                                                v96 -= v96
                                                                                                if false {
                                                                                                    if true {
                                                                                                        switch p2 {
                                                                                                        case "a":
                                                                                                            switch ((p2 + "a") + "a") {
                                                                                                            case "\n\t":
                                                                                                                if false {
                                                                                                                    var v174 int16 = 1
                                                                                                                    fmt.Println(p2)
                                                                                                                    var v175 uint8 = 2
                                                                                                                    _ = v174
                                                                                                                    _ = v175
                                                                                                                }
                                                                                                            case "a":
                                                                                                                v96 += (v8 / ((-9223372036854775808) | 1))
                                                                                                                var v176 bool = (1 >= 200)
                                                                                                                var v177 bool = false
                                                                                                                _ = v176
                                                                                                                _ = v177
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            var v178 uint8 = 1
                                                                                                            _ = v178
                                                                                                        }
                                                                                                        fmt.Println(v134, p0, v134)
                                                                                                        _ = len(v5)
                                                                                                        if true {
                                                                                                            if true {
                                                                                                                var v179 int32 = v135
                                                                                                                if true {
                                                                                                                    v5 = append(v5, (v96 / ((v8) | 1)))
                                                                                                                    switch v8 {
                                                                                                                    case 281:
                                                                                                                        var v180 float64 = ((((v134 / (v134 / 3.14159)) * v134) - -2.25) - (float64(0.0) / (v134 + ((float64(1e10) - v134) * v134))))
                                                                                                                        _ = v180
                                                                                                                    case 1:
                                                                                                                        var v181 uint64 = 4294967295
                                                                                                                        _ = v181
                                                                                                                    case -654:
                                                                                                                        p2 = "\n\t"
                                                                                                                    }
                                                                                                                    if len(v5) > 0 {
                                                                                                                        v5[0] = (fi6 >> ((1) & 63))
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    p0 -= (p0 | p0)
                                                                                                                }
                                                                                                                _ = v179
                                                                                                            }
                                                                                                        }
                                                                                                        if len(v5) > 0 {
                                                                                                            v5[0] = fi7
                                                                                                        }
                                                                                                    } else {
                                                                                                        v182 := GenS18{F0: 2147483647, F1: 510, F2: 2}
                                                                                                        switch "\n\t" {
                                                                                                        case "hello":
                                                                                                            var v183 int32 = v135
                                                                                                            v135 += v135
                                                                                                            if ((4294967295 > 382) && (847 > 255)) {
                                                                                                                var v184 bool = ((p2 + p2) != "a")
                                                                                                                _ = v184
                                                                                                            }
                                                                                                            _ = v183
                                                                                                        }
                                                                                                        _ = v182
                                                                                                    }
                                                                                                }
                                                                                            case 32767:
                                                                                                switch (v96 * (fi6 << ((452) & 63))) {
                                                                                                case 32767:
                                                                                                    if true {
                                                                                                        v185 := GenS11{F0: 905, F1: (p1 - p1), F2: "\n\t"}
                                                                                                        if false {
                                                                                                            switch -800 {
                                                                                                            case -2147483648:
                                                                                                                if false {
                                                                                                                    if (true && (100 == -1)) {
                                                                                                                        p1 -= 0
                                                                                                                    } else {
                                                                                                                        if len(v5) >= 2 {
                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                        }
                                                                                                                        _ = len(v5)
                                                                                                                    }
                                                                                                                    var v186 float32 = (1e10 * 0.0)
                                                                                                                    v5 = append(v5, 737)
                                                                                                                    p2 = ((p2 + p2) + p2)
                                                                                                                    if ((!false) || (!false)) {
                                                                                                                        fmt.Println(fi6)
                                                                                                                        v186 = v186
                                                                                                                    } else {
                                                                                                                        if true {
                                                                                                                            v135 = v135
                                                                                                                            var v187 int64 = 1
                                                                                                                            var v188 bool = (865 == 1)
                                                                                                                            if v188 {
                                                                                                                                v189 := GenS10{F0: v187, F1: v187, F2: 216}
                                                                                                                                _ = v189
                                                                                                                            }
                                                                                                                            _ = v187
                                                                                                                            _ = v188
                                                                                                                        }
                                                                                                                        var v190 string = p2
                                                                                                                        _ = v190
                                                                                                                    }
                                                                                                                    _ = v186
                                                                                                                } else {
                                                                                                                    if len(v5) >= 2 {
                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                    }
                                                                                                                    if false {
                                                                                                                        var v191 bool = (false || false)
                                                                                                                        _ = v191
                                                                                                                    }
                                                                                                                    switch "\n\t" {
                                                                                                                    case "hello":
                                                                                                                        v96 += (fi7 - (v8 * -2147483648))
                                                                                                                    case "":
                                                                                                                        switch p2 {
                                                                                                                        case "\n\t":
                                                                                                                            var v192 float64 = ((v134 * (v134 / (((v134 / (v134 - v134)) + v134) - 0.1))) * 1.5)
                                                                                                                            _ = v192
                                                                                                                        case "hello":
                                                                                                                            switch 128 {
                                                                                                                            case 881:
                                                                                                                                var v193 int = 0
                                                                                                                                _ = v193
                                                                                                                            default:
                                                                                                                                if (-6 != p1) {
                                                                                                                                    p1 -= p1
                                                                                                                                } else {
                                                                                                                                    v5 = append(v5, v96)
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        if len(v5) > 0 {
                                                                                                                            v5[0] = (v96 >> ((1) & 63))
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                if false {
                                                                                                                    var v194 uint32 = 42
                                                                                                                    if true {
                                                                                                                        fmt.Println(p0, fi7, v194)
                                                                                                                        var v195 int64 = -1
                                                                                                                        _ = v195
                                                                                                                    }
                                                                                                                    p2 = p2
                                                                                                                    _ = v194
                                                                                                                } else {
                                                                                                                    if false {
                                                                                                                        v134 += -2.25
                                                                                                                        if (0 >= 0) {
                                                                                                                            if true {
                                                                                                                                v5 = append(v5, 2147483647)
                                                                                                                                var v196 float32 = 0.0
                                                                                                                                fmt.Println(p2)
                                                                                                                                switch v96 {
                                                                                                                                case 32767:
                                                                                                                                    if true {
                                                                                                                                        if len(v5) >= 2 {
                                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                case -1:
                                                                                                                                    var v197 uint64 = 2
                                                                                                                                    _ = v197
                                                                                                                                case 197:
                                                                                                                                    v196 = v196
                                                                                                                                }
                                                                                                                                p1 = -943
                                                                                                                                _ = v196
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            fmt.Println(fi6, v8)
                                                                                                                        }
                                                                                                                        switch (v8 - 0) {
                                                                                                                        case 1:
                                                                                                                            if len(v5) >= 2 {
                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                            }
                                                                                                                        case -9223372036854775808:
                                                                                                                            fmt.Println(v134, v8, v8)
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch ("" + p2) {
                                                                                                                    case "a":
                                                                                                                        switch (^(v96)) {
                                                                                                                        case -9223372036854775808:
                                                                                                                            var v198 uint64 = 227
                                                                                                                            _ = v198
                                                                                                                        case -2147483648:
                                                                                                                            switch "\n\t" {
                                                                                                                            case "a":
                                                                                                                                v135 = (v135 >> ((466) & 63))
                                                                                                                            default:
                                                                                                                                if false {
                                                                                                                                    p1 = 1
                                                                                                                                    switch v8 {
                                                                                                                                    case 32768:
                                                                                                                                        switch 832 {
                                                                                                                                        case 42:
                                                                                                                                            v199 := GenS19{F0: (v135 >> ((0) & 63)), F1: 1, F2: 4294967295}
                                                                                                                                            _ = v199
                                                                                                                                        }
                                                                                                                                    case 128:
                                                                                                                                        if false {
                                                                                                                                            switch v8 {
                                                                                                                                            case 0:
                                                                                                                                                var v200 bool = (p2 == p2)
                                                                                                                                                _ = v200
                                                                                                                                            case 460:
                                                                                                                                                v201 := GenS20{F0: (v96 ^ v96)}
                                                                                                                                                v5 = append(v5, 128)
                                                                                                                                                v135 -= v135
                                                                                                                                                switch 127 {
                                                                                                                                                case 40:
                                                                                                                                                    switch (p2 + p2) {
                                                                                                                                                    case "foo bar":
                                                                                                                                                        var v202 int = (fi7 | v8)
                                                                                                                                                        _ = v202
                                                                                                                                                    case "a":
                                                                                                                                                        var v203 float64 = (((float64(0.1) + (v134 * v134)) + 0.0) / ((v134 + v134) / (v134 - (v134 - 0.1))))
                                                                                                                                                        _ = v203
                                                                                                                                                    }
                                                                                                                                                default:
                                                                                                                                                    var v204 uint64 = 0
                                                                                                                                                    _ = v204
                                                                                                                                                }
                                                                                                                                                if len(v5) >= 2 {
                                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                                }
                                                                                                                                                fmt.Println(v8)
                                                                                                                                                _ = v201
                                                                                                                                            }
                                                                                                                                            switch ("hello" + ("a" + p2)) {
                                                                                                                                            case "hello":
                                                                                                                                                p2 = ""
                                                                                                                                            default:
                                                                                                                                                v205 := GenS21{F0: v135, F1: int32(p0)}
                                                                                                                                                _ = v205
                                                                                                                                            }
                                                                                                                                        } else {
                                                                                                                                            if true {
                                                                                                                                                if (true || ((p2 + "") <= ((p2 + ("a" + (p2 + p2))) + ("hello" + p2)))) {
                                                                                                                                                    p1 += 1
                                                                                                                                                    switch fi7 {
                                                                                                                                                    case 721:
                                                                                                                                                        p0 -= p0
                                                                                                                                                    }
                                                                                                                                                    if (true || true) {
                                                                                                                                                        if len(v5) > 0 {
                                                                                                                                                            v5[0] = v8
                                                                                                                                                        }
                                                                                                                                                        if true {
                                                                                                                                                            v5 = append(v5, fi6)
                                                                                                                                                            v96 -= -9223372036854775808
                                                                                                                                                            v135 -= v135
                                                                                                                                                        } else {
                                                                                                                                                            v134 -= 0.0
                                                                                                                                                            if (1.5 < 0.1) {
                                                                                                                                                                if len(v5) >= 2 {
                                                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                                                }
                                                                                                                                                                if false {
                                                                                                                                                                    if true {
                                                                                                                                                                        v206 := GenS22{F0: 442, F1: 268, F2: -1}
                                                                                                                                                                        _ = v206
                                                                                                                                                                    } else {
                                                                                                                                                                        v8 -= 128
                                                                                                                                                                        p2 = "\n\t"
                                                                                                                                                                    }
                                                                                                                                                                    p1 += int16(fi7)
                                                                                                                                                                    var v207 float32 = float32(float64(-2.25))
                                                                                                                                                                    _ = v207
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            if (!(-1 >= -9223372036854775808)) {
                                                                                                                                                                v208 := GenS23{F0: 2, F1: (!(true || (424 == -9223372036854775808))), F2: 1}
                                                                                                                                                                v8 = v96
                                                                                                                                                                _ = v208
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        if (true && true) {
                                                                                                                                                            if len(v5) > 0 {
                                                                                                                                                                v5[0] = fi7
                                                                                                                                                            }
                                                                                                                                                        } else {
                                                                                                                                                            switch 333 {
                                                                                                                                                            case -540:
                                                                                                                                                                v96 = v8
                                                                                                                                                            default:
                                                                                                                                                                v134 += (v134 * (1e10 - v134))
                                                                                                                                                            }
                                                                                                                                                            switch (^(v96)) {
                                                                                                                                                            case 68:
                                                                                                                                                                fmt.Println(fi6)
                                                                                                                                                            case -2147483648:
                                                                                                                                                                if ((898 > 689) || (42 >= 9223372036854775807)) {
                                                                                                                                                                    if len(v5) > 0 {
                                                                                                                                                                        v5[0] = (^(v8))
                                                                                                                                                                    }
                                                                                                                                                                    v5 = append(v5, (v96 & (-(v8))))
                                                                                                                                                                    switch (fi7 << ((271) & 63)) {
                                                                                                                                                                    case 42:
                                                                                                                                                                        v209 := GenS24{F0: 127, F1: (v135 & (v135 / (((v135 / (((v135 * v135)) | 1))) | 1)))}
                                                                                                                                                                        _ = v209
                                                                                                                                                                    case 1:
                                                                                                                                                                        if len(v5) > 0 {
                                                                                                                                                                            v5[0] = int(v135)
                                                                                                                                                                        }
                                                                                                                                                                    case 32768:
                                                                                                                                                                        var v210 float64 = (v134 * (((v134 - ((-2.25 / -2.25) + (v134 * float64((0.0 + 0.1))))) - float64(-2.25)) * v134))
                                                                                                                                                                        _ = v210
                                                                                                                                                                    default:
                                                                                                                                                                        switch fi6 {
                                                                                                                                                                        case 1:
                                                                                                                                                                            p0 -= (p0 / (((^(p0))) | 1))
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                } else {
                                                                                                                                                                    var v211 int8 = -23
                                                                                                                                                                    if false {
                                                                                                                                                                        fmt.Println(v134, v96)
                                                                                                                                                                    } else {
                                                                                                                                                                        switch (((((p2 + ((p2 + p2) + p2)) + p2) + ("foo bar" + "hello")) + p2) + p2) {
                                                                                                                                                                        case "foo bar":
                                                                                                                                                                            if len(v5) >= 2 {
                                                                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                        v212 := GenS17{F0: fi7}
                                                                                                                                                                        _ = v212
                                                                                                                                                                    }
                                                                                                                                                                    _ = v211
                                                                                                                                                                }
                                                                                                                                                            case 128:
                                                                                                                                                                var v213 uint = 18446744073709551615
                                                                                                                                                                _ = v213
                                                                                                                                                            }
                                                                                                                                                            v8 = (^(fi7))
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                } else {
                                                                                                                                                    var v214 int16 = p1
                                                                                                                                                    _ = v214
                                                                                                                                                }
                                                                                                                                                _ = len(v5)
                                                                                                                                            } else {
                                                                                                                                                v215 := GenS12{F0: 128, F1: p1}
                                                                                                                                                _ = v215
                                                                                                                                            }
                                                                                                                                            v216 := GenS7{F0: fi7, F1: 370, F2: 1}
                                                                                                                                            var v217 int8 = -43
                                                                                                                                            _ = v216
                                                                                                                                            _ = v217
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        fmt.Println(fi7, fi7)
                                                                                                                                    }
                                                                                                                                } else {
                                                                                                                                    switch (^(fi6)) {
                                                                                                                                    case -970:
                                                                                                                                        if false {
                                                                                                                                            v218 := GenS20{F0: -615}
                                                                                                                                            if len(v5) > 0 {
                                                                                                                                                v5[0] = v96
                                                                                                                                            }
                                                                                                                                            _ = v218
                                                                                                                                        } else {
                                                                                                                                            var v219 uint32 = 4294967295
                                                                                                                                            var v220 uint64 = 4294967295
                                                                                                                                            switch (fi6 * v96) {
                                                                                                                                            case 9223372036854775807:
                                                                                                                                                if (1 <= 42) {
                                                                                                                                                    v134 = (((v134 - ((v134 * (-2.25 - v134)) + (v134 * 0.0))) - (-2.25 / v134)) + v134)
                                                                                                                                                    v221 := GenS25{F0: p0}
                                                                                                                                                    if ((^(v135)) < v135) {
                                                                                                                                                        v222 := GenS14{F0: 1, F1: v8, F2: false}
                                                                                                                                                        _ = v222
                                                                                                                                                    } else {
                                                                                                                                                        v223 := GenS8{F0: (^(p0)), F1: 255, F2: 42}
                                                                                                                                                        var v224 float64 = float64(-2.25)
                                                                                                                                                        _ = v223
                                                                                                                                                        _ = v224
                                                                                                                                                    }
                                                                                                                                                    var v225 int32 = (v135 | v135)
                                                                                                                                                    _ = v221
                                                                                                                                                    _ = v225
                                                                                                                                                } else {
                                                                                                                                                    var v226 int = (fi7 | fi6)
                                                                                                                                                    var v227 float32 = -2.25
                                                                                                                                                    _ = v226
                                                                                                                                                    _ = v227
                                                                                                                                                }
                                                                                                                                            case 32767:
                                                                                                                                                _ = len(v5)
                                                                                                                                            case 860:
                                                                                                                                                if false {
                                                                                                                                                    v8 -= (fi7 * v96)
                                                                                                                                                    if len(v5) > 0 {
                                                                                                                                                        v5[0] = fi7
                                                                                                                                                    }
                                                                                                                                                    if (false || (!false)) {
                                                                                                                                                        if (true || true) {
                                                                                                                                                            if ((!((v135 % ((v135) | 1)) == 235)) && true) {
                                                                                                                                                                p1 -= p1
                                                                                                                                                                v228 := GenS15{F0: -1}
                                                                                                                                                                _ = v228
                                                                                                                                                            } else {
                                                                                                                                                                fmt.Println(v8, v8, fi6)
                                                                                                                                                                if len(v5) >= 2 {
                                                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                    } else {
                                                                                                                                                        var v229 uint32 = v219
                                                                                                                                                        if (false && true) {
                                                                                                                                                            v219 += 232
                                                                                                                                                            if true {
                                                                                                                                                                switch "a" {
                                                                                                                                                                case "\n\t":
                                                                                                                                                                    switch ((("hello" + ("" + (((((p2 + p2) + ("" + "\n\t")) + p2) + "\n\t") + p2))) + (("" + ("\n\t" + p2)) + (p2 + p2))) + "foo bar") {
                                                                                                                                                                    case "\n\t":
                                                                                                                                                                        p0 += (p0 + (p0 - (p0 ^ (p0 ^ 32767))))
                                                                                                                                                                    case "foo bar":
                                                                                                                                                                        fmt.Println(v135)
                                                                                                                                                                    default:
                                                                                                                                                                        if (p1 == (^(p1))) {
                                                                                                                                                                            var v230 bool = (4294967295 >= 18446744073709551615)
                                                                                                                                                                            v96 = v96
                                                                                                                                                                            v231 := GenS11{F0: 9223372036854775807, F1: (p1 - p1), F2: p2}
                                                                                                                                                                            _ = v230
                                                                                                                                                                            _ = v231
                                                                                                                                                                        } else {
                                                                                                                                                                            if ((!(v134 <= v134)) && true) {
                                                                                                                                                                                v5 = append(v5, (v8 << ((1) & 63)))
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                case "foo bar":
                                                                                                                                                                    if len(v5) >= 2 {
                                                                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                                                                    }
                                                                                                                                                                case "hello":
                                                                                                                                                                    fmt.Println(p0, v229, fi6)
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                            v8 -= v96
                                                                                                                                                        }
                                                                                                                                                        _ = v229
                                                                                                                                                    }
                                                                                                                                                    if ((true && false) && false) {
                                                                                                                                                        var v232 uint16 = p0
                                                                                                                                                        switch (v96 / (((v96 * v8)) | 1)) {
                                                                                                                                                        case 2147483647:
                                                                                                                                                            v233 := GenS26{F0: (!false)}
                                                                                                                                                            _ = v233
                                                                                                                                                        case 4294967295:
                                                                                                                                                            fmt.Println(v232)
                                                                                                                                                        case 32767:
                                                                                                                                                            fmt.Println(v96, v96)
                                                                                                                                                        }
                                                                                                                                                        _ = v232
                                                                                                                                                    } else {
                                                                                                                                                        switch 42 {
                                                                                                                                                        case 988:
                                                                                                                                                            switch (fi6 + 419) {
                                                                                                                                                            case 1:
                                                                                                                                                                if (p2 < p2) {
                                                                                                                                                                    var v234 int32 = v135
                                                                                                                                                                    _ = v234
                                                                                                                                                                } else {
                                                                                                                                                                    switch (p2 + (("a" + p2) + ("hello" + p2))) {
                                                                                                                                                                    case "foo bar":
                                                                                                                                                                        var v235 bool = ((!false) && false)
                                                                                                                                                                        _ = v235
                                                                                                                                                                    case "hello":
                                                                                                                                                                        v236 := GenS27{F0: 127}
                                                                                                                                                                        _ = v236
                                                                                                                                                                    case "\n\t":
                                                                                                                                                                        v5 = append(v5, (v8 + v8))
                                                                                                                                                                    }
                                                                                                                                                                    switch "\n\t" {
                                                                                                                                                                    case "hello":
                                                                                                                                                                        v220 -= (v220 ^ (v220 % ((0) | 1)))
                                                                                                                                                                    default:
                                                                                                                                                                        if len(v5) >= 2 {
                                                                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                    if len(v5) >= 2 {
                                                                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            case -830:
                                                                                                                                                                if len(v5) > 0 {
                                                                                                                                                                    v5[0] = v96
                                                                                                                                                                }
                                                                                                                                                            case 218:
                                                                                                                                                                v220 += (v220 & 42)
                                                                                                                                                            default:
                                                                                                                                                                if len(v5) >= 2 {
                                                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        case 9223372036854775807:
                                                                                                                                                            if (v219 >= (^(v219))) {
                                                                                                                                                                var v237 string = (p2 + (("\n\t" + p2) + ((((p2 + (p2 + p2)) + "") + "hello") + (p2 + ((p2 + p2) + "")))))
                                                                                                                                                                _ = v237
                                                                                                                                                            }
                                                                                                                                                        default:
                                                                                                                                                            v238 := GenS18{F0: 810, F1: (v219 % (((v219 & v219)) | 1)), F2: 858}
                                                                                                                                                            _ = v238
                                                                                                                                                        }
                                                                                                                                                        _ = len(v5)
                                                                                                                                                        switch fi7 {
                                                                                                                                                        case 509:
                                                                                                                                                            fmt.Println(p0, fi6)
                                                                                                                                                        case -1:
                                                                                                                                                            var v239 bool = (723 >= 9223372036854775807)
                                                                                                                                                            _ = v239
                                                                                                                                                        }
                                                                                                                                                        var v240 int16 = -1
                                                                                                                                                        _ = v240
                                                                                                                                                    }
                                                                                                                                                    if true {
                                                                                                                                                        var v241 uint32 = 744
                                                                                                                                                        v242 := GenS26{F0: true}
                                                                                                                                                        if true {
                                                                                                                                                            if ((0.0 - float32((v134 / float64(0.1)))) != (float32((v134 - v134)) + 1.5)) {
                                                                                                                                                                v243 := GenS28{F0: 922}
                                                                                                                                                                v244 := GenS4{F0: (v135 & (v135 << ((0) & 63))), F1: v220}
                                                                                                                                                                if true {
                                                                                                                                                                    switch (((p2 + "a") + "hello") + p2) {
                                                                                                                                                                    case "a":
                                                                                                                                                                        var v245 float64 = 0.1
                                                                                                                                                                        _ = v245
                                                                                                                                                                    case "foo bar":
                                                                                                                                                                        v219 += 1
                                                                                                                                                                    default:
                                                                                                                                                                        if false {
                                                                                                                                                                            if (false && true) {
                                                                                                                                                                                v96 -= fi6
                                                                                                                                                                                v135 = (v135 << ((2) & 63))
                                                                                                                                                                                var v246 string = "foo bar"
                                                                                                                                                                                switch 4294967295 {
                                                                                                                                                                                case 4294967295:
                                                                                                                                                                                    v247 := GenS7{F0: -2147483648, F1: 738, F2: (p0 << ((255) & 63))}
                                                                                                                                                                                    _ = v247
                                                                                                                                                                                case 549:
                                                                                                                                                                                    p2 = p2
                                                                                                                                                                                case 127:
                                                                                                                                                                                    var v248 float64 = v134
                                                                                                                                                                                    _ = v248
                                                                                                                                                                                }
                                                                                                                                                                                fmt.Println(v135, v96, v219)
                                                                                                                                                                                _ = v246
                                                                                                                                                                            } else {
                                                                                                                                                                                var v249 int8 = 106
                                                                                                                                                                                if true {
                                                                                                                                                                                    v250 := GenS15{F0: -9223372036854775808}
                                                                                                                                                                                    _ = len(v5)
                                                                                                                                                                                    _ = v250
                                                                                                                                                                                } else {
                                                                                                                                                                                    if len(v5) > 0 {
                                                                                                                                                                                        v5[0] = (v8 / (((fi6 << ((2) & 63))) | 1))
                                                                                                                                                                                    }
                                                                                                                                                                                    v219 += (v241 << ((42) & 63))
                                                                                                                                                                                    if len(v5) > 0 {
                                                                                                                                                                                        v5[0] = 42
                                                                                                                                                                                    }
                                                                                                                                                                                    var v251 bool = (!(-1 != 1))
                                                                                                                                                                                    v252 := GenS4{F0: v135, F1: (v220 + 660)}
                                                                                                                                                                                    _ = v251
                                                                                                                                                                                    _ = v252
                                                                                                                                                                                }
                                                                                                                                                                                if len(v5) >= 2 {
                                                                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                                                                }
                                                                                                                                                                                var v253 float32 = (0.1 * float32(v134))
                                                                                                                                                                                var v254 bool = (int16(v241) == -481)
                                                                                                                                                                                p1 = p1
                                                                                                                                                                                if false {
                                                                                                                                                                                    v253 += v253
                                                                                                                                                                                    if (false || v254) {
                                                                                                                                                                                        if ((^(v135)) >= 2147483647) {
                                                                                                                                                                                            var v255 int8 = (-(v249))
                                                                                                                                                                                            _ = v255
                                                                                                                                                                                        } else {
                                                                                                                                                                                            if ((^(p0)) <= 42) {
                                                                                                                                                                                                fmt.Println(fi6)
                                                                                                                                                                                            } else {
                                                                                                                                                                                                v8 -= (v8 % ((v96) | 1))
                                                                                                                                                                                                fmt.Println(v253, v241, p0)
                                                                                                                                                                                                v256 := GenS19{F0: v135, F1: uint16(v241), F2: (v220 % ((v220) | 1))}
                                                                                                                                                                                                _ = v256
                                                                                                                                                                                            }
                                                                                                                                                                                        }
                                                                                                                                                                                        if ((0.1 > 1.5) || v254) {
                                                                                                                                                                                            v254 = v254
                                                                                                                                                                                        } else {
                                                                                                                                                                                            v8 += v8
                                                                                                                                                                                            var v257 float64 = v134
                                                                                                                                                                                            p1 = (p1 & 28)
                                                                                                                                                                                            _ = v257
                                                                                                                                                                                        }
                                                                                                                                                                                    } else {
                                                                                                                                                                                        v254 = (563 == 42)
                                                                                                                                                                                        v258 := GenS13{F0: v135}
                                                                                                                                                                                        switch v96 {
                                                                                                                                                                                        case 32768:
                                                                                                                                                                                            switch p2 {
                                                                                                                                                                                            case "":
                                                                                                                                                                                                v5 = append(v5, (^(fi6)))
                                                                                                                                                                                            default:
                                                                                                                                                                                                var v259 uint16 = 1
                                                                                                                                                                                                var v260 uint16 = (p0 & (p0 / ((p0) | 1)))
                                                                                                                                                                                                switch (v96 >> ((0) & 63)) {
                                                                                                                                                                                                case -1:
                                                                                                                                                                                                    v241 -= 49
                                                                                                                                                                                                default:
                                                                                                                                                                                                    if v254 {
                                                                                                                                                                                                        fmt.Println(v134, v220, v96)
                                                                                                                                                                                                    }
                                                                                                                                                                                                }
                                                                                                                                                                                                var v261 uint32 = (v219 >> ((18446744073709551615) & 63))
                                                                                                                                                                                                if v254 {
                                                                                                                                                                                                    if (v254 && (!v254)) {
                                                                                                                                                                                                        var v262 uint16 = v260
                                                                                                                                                                                                        if v254 {
                                                                                                                                                                                                            var v263 uint16 = v262
                                                                                                                                                                                                            v241 += v261
                                                                                                                                                                                                            _ = v263
                                                                                                                                                                                                        }
                                                                                                                                                                                                        _ = len(v5)
                                                                                                                                                                                                        _ = v262
                                                                                                                                                                                                    } else {
                                                                                                                                                                                                        fmt.Println(fi6, p1)
                                                                                                                                                                                                        fmt.Println(p1)
                                                                                                                                                                                                        var v264 uint32 = 754
                                                                                                                                                                                                        v265 := GenS29{F0: 0, F1: v259}
                                                                                                                                                                                                        fmt.Println(v134)
                                                                                                                                                                                                        if v254 {
                                                                                                                                                                                                            _ = len(v5)
                                                                                                                                                                                                        } else {
                                                                                                                                                                                                            if v254 {
                                                                                                                                                                                                                switch 0 {
                                                                                                                                                                                                                case 2147483647:
                                                                                                                                                                                                                    var v266 uint = 65535
                                                                                                                                                                                                                    _ = v266
                                                                                                                                                                                                                case 0:
                                                                                                                                                                                                                    v253 += v253
                                                                                                                                                                                                                default:
                                                                                                                                                                                                                    fmt.Println(fi6)
                                                                                                                                                                                                                }
                                                                                                                                                                                                                var v267 string = "hello"
                                                                                                                                                                                                                _ = v267
                                                                                                                                                                                                            }
                                                                                                                                                                                                            var v268 int16 = 42
                                                                                                                                                                                                            if len(v5) >= 2 {
                                                                                                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                                                                                                            }
                                                                                                                                                                                                            _ = v268
                                                                                                                                                                                                        }
                                                                                                                                                                                                        v269 := GenS1{F0: "hello", F1: (v219 | (v261 ^ 406))}
                                                                                                                                                                                                        v96 -= v96
                                                                                                                                                                                                        _ = v264
                                                                                                                                                                                                        _ = v265
                                                                                                                                                                                                        _ = v269
                                                                                                                                                                                                    }
                                                                                                                                                                                                    var v270 int64 = -969
                                                                                                                                                                                                    var v271 string = "a"
                                                                                                                                                                                                    _ = v270
                                                                                                                                                                                                    _ = v271
                                                                                                                                                                                                }
                                                                                                                                                                                                _ = v259
                                                                                                                                                                                                _ = v260
                                                                                                                                                                                                _ = v261
                                                                                                                                                                                            }
                                                                                                                                                                                        default:
                                                                                                                                                                                            if (v254 && v254) {
                                                                                                                                                                                                v134 += v134
                                                                                                                                                                                            } else {
                                                                                                                                                                                                switch int(v135) {
                                                                                                                                                                                                case -899:
                                                                                                                                                                                                    v272 := GenS26{F0: (!v254)}
                                                                                                                                                                                                    _ = v272
                                                                                                                                                                                                default:
                                                                                                                                                                                                    v273 := GenS23{F0: (p0 + p0), F1: ((float64(float32(v134)) != (v134 / v134)) || v254), F2: 4294967295}
                                                                                                                                                                                                    _ = v273
                                                                                                                                                                                                }
                                                                                                                                                                                                var v274 int = (fi7 ^ 128)
                                                                                                                                                                                                _ = v274
                                                                                                                                                                                            }
                                                                                                                                                                                        }
                                                                                                                                                                                        _ = v258
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                                v275 := GenS30{F0: (^(v220))}
                                                                                                                                                                                _ = v249
                                                                                                                                                                                _ = v253
                                                                                                                                                                                _ = v254
                                                                                                                                                                                _ = v275
                                                                                                                                                                            }
                                                                                                                                                                            var v276 int64 = 9223372036854775807
                                                                                                                                                                            var v277 float32 = 0.1
                                                                                                                                                                            _ = v276
                                                                                                                                                                            _ = v277
                                                                                                                                                                        } else {
                                                                                                                                                                            if (!(3.14159 == 1e10)) {
                                                                                                                                                                                v5 = append(v5, (fi6 - (fi7 << ((18446744073709551615) & 63))))
                                                                                                                                                                                if false {
                                                                                                                                                                                    p2 = (p2 + p2)
                                                                                                                                                                                    var v278 uint64 = (^(v220))
                                                                                                                                                                                    if len(v5) > 0 {
                                                                                                                                                                                        v5[0] = (v8 - (fi7 * (v8 ^ (v96 - fi6))))
                                                                                                                                                                                    }
                                                                                                                                                                                    _ = v278
                                                                                                                                                                                } else {
                                                                                                                                                                                    var v279 uint32 = (v241 - v241)
                                                                                                                                                                                    v241 = (v241 << ((718) & 63))
                                                                                                                                                                                    if false {
                                                                                                                                                                                        switch (-(v96)) {
                                                                                                                                                                                        case 483:
                                                                                                                                                                                            fmt.Println(p2, fi7, v96)
                                                                                                                                                                                        case 42:
                                                                                                                                                                                            fmt.Println(p1, v8, v8)
                                                                                                                                                                                        }
                                                                                                                                                                                        fmt.Println(v134, fi7)
                                                                                                                                                                                    } else {
                                                                                                                                                                                        var v280 float32 = 1e10
                                                                                                                                                                                        fmt.Println(v280, v241, v135)
                                                                                                                                                                                        _ = len(v5)
                                                                                                                                                                                        _ = v280
                                                                                                                                                                                    }
                                                                                                                                                                                    _ = v279
                                                                                                                                                                                }
                                                                                                                                                                            } else {
                                                                                                                                                                                v241 += (v241 * v241)
                                                                                                                                                                                if false {
                                                                                                                                                                                    if (18446744073709551615 >= (v220 / ((v220) | 1))) {
                                                                                                                                                                                        p1 += p1
                                                                                                                                                                                        p0 = p0
                                                                                                                                                                                        if (!false) {
                                                                                                                                                                                            if len(v5) >= 2 {
                                                                                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                                                                                            }
                                                                                                                                                                                            v5 = append(v5, fi7)
                                                                                                                                                                                        } else {
                                                                                                                                                                                            if false {
                                                                                                                                                                                                if false {
                                                                                                                                                                                                    var v281 float64 = float64(float32(v134))
                                                                                                                                                                                                    _ = v281
                                                                                                                                                                                                } else {
                                                                                                                                                                                                    v219 -= (v241 | v219)
                                                                                                                                                                                                }
                                                                                                                                                                                                switch (^(v96)) {
                                                                                                                                                                                                case 9223372036854775807:
                                                                                                                                                                                                    _ = len(v5)
                                                                                                                                                                                                case 32767:
                                                                                                                                                                                                    if len(v5) > 0 {
                                                                                                                                                                                                        v5[0] = v96
                                                                                                                                                                                                    }
                                                                                                                                                                                                default:
                                                                                                                                                                                                    switch p2 {
                                                                                                                                                                                                    case "foo bar":
                                                                                                                                                                                                        v282 := GenS19{F0: v135, F1: p0, F2: (v220 | 4294967295)}
                                                                                                                                                                                                        _ = v282
                                                                                                                                                                                                    case "\n\t":
                                                                                                                                                                                                        var v283 string = ""
                                                                                                                                                                                                        _ = v283
                                                                                                                                                                                                    default:
                                                                                                                                                                                                        if true {
                                                                                                                                                                                                            var v284 int16 = (p1 + p1)
                                                                                                                                                                                                            fmt.Println(v8, v135, v241)
                                                                                                                                                                                                            if len(v5) >= 2 {
                                                                                                                                                                                                                v5 = v5[1:len(v5)]
                                                                                                                                                                                                            }
                                                                                                                                                                                                            var v285 float32 = -2.25
                                                                                                                                                                                                            _ = v284
                                                                                                                                                                                                            _ = v285
                                                                                                                                                                                                        } else {
                                                                                                                                                                                                            var v286 uint64 = 0
                                                                                                                                                                                                            fmt.Println(v286, p0)
                                                                                                                                                                                                            var v287 uint16 = uint16(v219)
                                                                                                                                                                                                            _ = v286
                                                                                                                                                                                                            _ = v287
                                                                                                                                                                                                        }
                                                                                                                                                                                                    }
                                                                                                                                                                                                }
                                                                                                                                                                                                switch (fi6 * fi6) {
                                                                                                                                                                                                case 42:
                                                                                                                                                                                                    var v288 int = fi7
                                                                                                                                                                                                    _ = v288
                                                                                                                                                                                                case 279:
                                                                                                                                                                                                    v289 := GenS31{F0: v220}
                                                                                                                                                                                                    _ = v289
                                                                                                                                                                                                case 471:
                                                                                                                                                                                                    if len(v5) > 0 {
                                                                                                                                                                                                        v5[0] = (v8 >> ((1) & 63))
                                                                                                                                                                                                    }
                                                                                                                                                                                                default:
                                                                                                                                                                                                    fmt.Println(v241, v8, v135)
                                                                                                                                                                                                }
                                                                                                                                                                                            } else {
                                                                                                                                                                                                if len(v5) >= 2 {
                                                                                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                                                                                }
                                                                                                                                                                                            }
                                                                                                                                                                                            v290 := GenS8{F0: (p0 | (p0 | p0)), F1: 2, F2: 255}
                                                                                                                                                                                            _ = v290
                                                                                                                                                                                        }
                                                                                                                                                                                        var v291 uint64 = v220
                                                                                                                                                                                        _ = v291
                                                                                                                                                                                    } else {
                                                                                                                                                                                        fmt.Println(p2, v241, p1)
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                } else {
                                                                                                                                                                    var v292 int = v8
                                                                                                                                                                    v292 = (fi6 / (((fi6 | (^(fi6)))) | 1))
                                                                                                                                                                    _ = v292
                                                                                                                                                                }
                                                                                                                                                                _ = v243
                                                                                                                                                                _ = v244
                                                                                                                                                            }
                                                                                                                                                            var v293 int64 = 42
                                                                                                                                                            if false {
                                                                                                                                                                if true {
                                                                                                                                                                    if len(v5) >= 2 {
                                                                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                                                                    }
                                                                                                                                                                    if (v220 < (v220 % (((v220 | v220)) | 1))) {
                                                                                                                                                                        v219 -= (^(v241))
                                                                                                                                                                    }
                                                                                                                                                                    if false {
                                                                                                                                                                        if len(v5) > 0 {
                                                                                                                                                                            v5[0] = (fi7 | (fi6 / ((int(v219)) | 1)))
                                                                                                                                                                        }
                                                                                                                                                                        v294 := GenS25{F0: uint16(v220)}
                                                                                                                                                                        _ = v294
                                                                                                                                                                    } else {
                                                                                                                                                                        v295 := GenS8{F0: (p0 * (p0 * (p0 + p0))), F1: 42, F2: 52}
                                                                                                                                                                        v241 = (v241 << ((18446744073709551615) & 63))
                                                                                                                                                                        if len(v5) >= 2 {
                                                                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                                                                        }
                                                                                                                                                                        var v296 bool = false
                                                                                                                                                                        _ = len(v5)
                                                                                                                                                                        _ = v295
                                                                                                                                                                        _ = v296
                                                                                                                                                                    }
                                                                                                                                                                    switch "a" {
                                                                                                                                                                    case "a":
                                                                                                                                                                        var v297 uint16 = p0
                                                                                                                                                                        fmt.Println(v220)
                                                                                                                                                                        v298 := GenS26{F0: (((^(v220)) < 144) && (!(!(!true))))}
                                                                                                                                                                        _ = v297
                                                                                                                                                                        _ = v298
                                                                                                                                                                    case "foo bar":
                                                                                                                                                                        p1 -= 1
                                                                                                                                                                    }
                                                                                                                                                                } else {
                                                                                                                                                                    v299 := GenS32{F0: "hello"}
                                                                                                                                                                    v5 = append(v5, (v96 * 2147483647))
                                                                                                                                                                    v293 -= (v293 | -972)
                                                                                                                                                                    _ = v299
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                var v300 uint32 = (v219 * (v241 << ((253) & 63)))
                                                                                                                                                                switch int(v293) {
                                                                                                                                                                case -411:
                                                                                                                                                                    v5 = append(v5, (fi6 + int(p1)))
                                                                                                                                                                case 128:
                                                                                                                                                                    var v301 int8 = 60
                                                                                                                                                                    _ = v301
                                                                                                                                                                case 1:
                                                                                                                                                                    v219 = 336
                                                                                                                                                                default:
                                                                                                                                                                    p2 = p2
                                                                                                                                                                    var v302 uint16 = 2
                                                                                                                                                                    v293 = (v293 % ((v293) | 1))
                                                                                                                                                                    _ = v302
                                                                                                                                                                }
                                                                                                                                                                _ = v300
                                                                                                                                                            }
                                                                                                                                                            _ = v293
                                                                                                                                                        } else {
                                                                                                                                                            if len(v5) > 0 {
                                                                                                                                                                v5[0] = (fi6 % ((v96) | 1))
                                                                                                                                                            }
                                                                                                                                                            if true {
                                                                                                                                                                if (0 != 4294967295) {
                                                                                                                                                                    fmt.Println(p2, p0)
                                                                                                                                                                    v303 := GenS33{F0: 0}
                                                                                                                                                                    p0 -= 42
                                                                                                                                                                    _ = v303
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                _ = len(v5)
                                                                                                                                                                var v304 bool = (false && ("" <= "foo bar"))
                                                                                                                                                                p1 += p1
                                                                                                                                                                _ = v304
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        v305 := GenS19{F0: v135, F1: p0, F2: v220}
                                                                                                                                                        _ = v241
                                                                                                                                                        _ = v242
                                                                                                                                                        _ = v305
                                                                                                                                                    } else {
                                                                                                                                                        var v306 int32 = v135
                                                                                                                                                        _ = v306
                                                                                                                                                    }
                                                                                                                                                    switch "" {
                                                                                                                                                    case "hello":
                                                                                                                                                        var v307 uint32 = 65535
                                                                                                                                                        _ = v307
                                                                                                                                                    case "foo bar":
                                                                                                                                                        v219 = 882
                                                                                                                                                    default:
                                                                                                                                                        if len(v5) >= 2 {
                                                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            _ = v219
                                                                                                                                            _ = v220
                                                                                                                                        }
                                                                                                                                    case 574:
                                                                                                                                        switch (v8 >> ((255) & 63)) {
                                                                                                                                        case 128:
                                                                                                                                            var v308 float64 = v134
                                                                                                                                            _ = v308
                                                                                                                                        }
                                                                                                                                    case -115:
                                                                                                                                        fmt.Println(v134, v8, p0)
                                                                                                                                    }
                                                                                                                                    var v309 bool = (false && (!false))
                                                                                                                                    v309 = (!v309)
                                                                                                                                    if len(v5) >= 2 {
                                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                                    }
                                                                                                                                    _ = v309
                                                                                                                                }
                                                                                                                            }
                                                                                                                        case 32767:
                                                                                                                            v310 := GenS22{F0: 995, F1: 883, F2: 1}
                                                                                                                            _ = v310
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        if (("" + (p2 + "")) == p2) {
                                                                                                                            if true {
                                                                                                                                if len(v5) >= 2 {
                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                }
                                                                                                                                v311 := GenS34{F0: 1}
                                                                                                                                if false {
                                                                                                                                    switch 32767 {
                                                                                                                                    case 2147483647:
                                                                                                                                        if false {
                                                                                                                                            switch ("hello" + (p2 + p2)) {
                                                                                                                                            case "":
                                                                                                                                                v135 = -2147483648
                                                                                                                                            }
                                                                                                                                            v312 := GenS10{F0: 42, F1: -1, F2: 1}
                                                                                                                                            if len(v5) > 0 {
                                                                                                                                                v5[0] = v96
                                                                                                                                            }
                                                                                                                                            _ = v312
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        var v313 int = (^(fi7))
                                                                                                                                        _ = v313
                                                                                                                                    }
                                                                                                                                    if false {
                                                                                                                                        if len(v5) > 0 {
                                                                                                                                            v5[0] = (fi6 - fi6)
                                                                                                                                        }
                                                                                                                                        if false {
                                                                                                                                            p2 = "hello"
                                                                                                                                        } else {
                                                                                                                                            var v314 uint16 = p0
                                                                                                                                            var v315 uint16 = p0
                                                                                                                                            _ = v314
                                                                                                                                            _ = v315
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                    switch (fi7 >> ((42) & 63)) {
                                                                                                                                    case 1:
                                                                                                                                        if (42 <= 42) {
                                                                                                                                            p0 -= p0
                                                                                                                                            if ((p0 + 309) == p0) {
                                                                                                                                                p1 = int16(v135)
                                                                                                                                                switch (fi7 >> ((324) & 63)) {
                                                                                                                                                case 128:
                                                                                                                                                    p0 += p0
                                                                                                                                                case 4294967295:
                                                                                                                                                    var v316 int8 = 0
                                                                                                                                                    _ = v316
                                                                                                                                                case -1:
                                                                                                                                                    v134 -= 1e10
                                                                                                                                                default:
                                                                                                                                                    var v317 int = v8
                                                                                                                                                    _ = v317
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            var v318 float64 = v134
                                                                                                                                            v319 := GenS2{F0: 30, F1: 0}
                                                                                                                                            if ((v135 + v135) <= (v135 - v135)) {
                                                                                                                                                if len(v5) > 0 {
                                                                                                                                                    v5[0] = 4294967295
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            _ = v318
                                                                                                                                            _ = v319
                                                                                                                                        }
                                                                                                                                    case 128:
                                                                                                                                        var v320 int16 = (p1 | p1)
                                                                                                                                        _ = v320
                                                                                                                                    case 32767:
                                                                                                                                        if (42 > 0) {
                                                                                                                                            if true {
                                                                                                                                                if false {
                                                                                                                                                    var v321 uint64 = 2
                                                                                                                                                    _ = v321
                                                                                                                                                } else {
                                                                                                                                                    var v322 int = v96
                                                                                                                                                    var v323 string = p2
                                                                                                                                                    v323 = "hello"
                                                                                                                                                    fmt.Println(fi7, v96, fi6)
                                                                                                                                                    if len(v5) > 0 {
                                                                                                                                                        v5[0] = v322
                                                                                                                                                    }
                                                                                                                                                    v323 = "hello"
                                                                                                                                                    if ((fi7 / ((-1) | 1)) > v96) {
                                                                                                                                                        v324 := GenS22{F0: 18446744073709551615, F1: 65535, F2: 1}
                                                                                                                                                        _ = v324
                                                                                                                                                    } else {
                                                                                                                                                        fmt.Println(fi6, v135, p2)
                                                                                                                                                    }
                                                                                                                                                    v322 = (v96 % ((576) | 1))
                                                                                                                                                    _ = v322
                                                                                                                                                    _ = v323
                                                                                                                                                }
                                                                                                                                                if true {
                                                                                                                                                    if false {
                                                                                                                                                        switch v96 {
                                                                                                                                                        case 615:
                                                                                                                                                            switch 1 {
                                                                                                                                                            case -2147483648:
                                                                                                                                                                switch (v8 - 957) {
                                                                                                                                                                case -880:
                                                                                                                                                                    fmt.Println(v96, v96)
                                                                                                                                                                case -236:
                                                                                                                                                                    v134 -= (float64((0.1 / 1.0)) + 1.5)
                                                                                                                                                                    v325 := GenS20{F0: (fi6 * fi6)}
                                                                                                                                                                    _ = v325
                                                                                                                                                                case 128:
                                                                                                                                                                    if (((1 == 4294967295) && (((true || ("hello" > p2)) && (1 <= 617)) || (((3.14159 / float64(float32(1.5))) <= 0.0) || false))) && ((p1 - p1) != p1)) {
                                                                                                                                                                        var v326 int = (v96 - (v8 + (fi6 / (((fi7 << ((4294967295) & 63))) | 1))))
                                                                                                                                                                        if false {
                                                                                                                                                                            v5 = append(v5, 128)
                                                                                                                                                                        }
                                                                                                                                                                        if false {
                                                                                                                                                                            switch 2147483647 {
                                                                                                                                                                            case 127:
                                                                                                                                                                                if (!false) {
                                                                                                                                                                                    if true {
                                                                                                                                                                                        if true {
                                                                                                                                                                                            v135 -= (v135 % ((v135) | 1))
                                                                                                                                                                                            v327 := GenS31{F0: 4294967295}
                                                                                                                                                                                            var v328 int64 = 387
                                                                                                                                                                                            _ = v327
                                                                                                                                                                                            _ = v328
                                                                                                                                                                                        }
                                                                                                                                                                                        if len(v5) >= 2 {
                                                                                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                                                                                        }
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                            fmt.Println(v134, fi6)
                                                                                                                                                                            v329 := GenS35{F0: 4294967295, F1: p1}
                                                                                                                                                                            _ = v329
                                                                                                                                                                        }
                                                                                                                                                                        _ = v326
                                                                                                                                                                    } else {
                                                                                                                                                                        v96 += 648
                                                                                                                                                                        var v330 float64 = 1e10
                                                                                                                                                                        p2 = "hello"
                                                                                                                                                                        _ = v330
                                                                                                                                                                    }
                                                                                                                                                                    if (0.1 <= v134) {
                                                                                                                                                                        v331 := GenS36{F0: 1, F1: 18446744073709551615}
                                                                                                                                                                        _ = v331
                                                                                                                                                                    } else {
                                                                                                                                                                        p1 = (p1 / (((p1 << ((18446744073709551615) & 63))) | 1))
                                                                                                                                                                        v134 = v134
                                                                                                                                                                        var v332 uint = 190
                                                                                                                                                                        v96 = (fi7 * int(v332))
                                                                                                                                                                        _ = v332
                                                                                                                                                                    }
                                                                                                                                                                    p1 = (p1 >> ((255) & 63))
                                                                                                                                                                    v5 = append(v5, v96)
                                                                                                                                                                default:
                                                                                                                                                                    p2 = p2
                                                                                                                                                                }
                                                                                                                                                                v134 += v134
                                                                                                                                                                switch (v96 / ((v8) | 1)) {
                                                                                                                                                                case 1:
                                                                                                                                                                    v333 := GenS18{F0: v135, F1: 929, F2: 18446744073709551615}
                                                                                                                                                                    _ = v333
                                                                                                                                                                case 4294967295:
                                                                                                                                                                    if true {
                                                                                                                                                                        switch (("foo bar" + (p2 + ((p2 + ((p2 + "\n\t") + p2)) + p2))) + "a") {
                                                                                                                                                                        case "\n\t":
                                                                                                                                                                            v334 := GenS5{F0: 1}
                                                                                                                                                                            v335 := GenS37{F0: (v135 % (((v135 + -229)) | 1)), F1: 559, F2: 18446744073709551615}
                                                                                                                                                                            _ = v334
                                                                                                                                                                            _ = v335
                                                                                                                                                                        case "foo bar":
                                                                                                                                                                            switch (p2 + (p2 + ("a" + ((p2 + p2) + p2)))) {
                                                                                                                                                                            case "\n\t":
                                                                                                                                                                                v8 = (v96 + v96)
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                        v134 -= float64(0.1)
                                                                                                                                                                    }
                                                                                                                                                                case -468:
                                                                                                                                                                    var v336 uint16 = p0
                                                                                                                                                                    _ = v336
                                                                                                                                                                }
                                                                                                                                                                fmt.Println(fi7, fi6)
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                        if ((p1 / ((p1) | 1)) <= (p1 + p1)) {
                                                                                                                                                            fmt.Println(p0, p1)
                                                                                                                                                            v5 = append(v5, v8)
                                                                                                                                                        } else {
                                                                                                                                                            _ = len(v5)
                                                                                                                                                            v135 = v135
                                                                                                                                                            if (((1 < -688) && (!true)) && (true && false)) {
                                                                                                                                                                v337 := GenS13{F0: v135}
                                                                                                                                                                var v338 bool = ((1.5 * (-2.25 + float32(0.0))) == ((0.1 + ((-2.25 / 1e10) - -2.25)) * -2.25))
                                                                                                                                                                _ = v337
                                                                                                                                                                _ = v338
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    v96 += (fi7 << ((42) & 63))
                                                                                                                                                } else {
                                                                                                                                                    if (!(65535 == 950)) {
                                                                                                                                                        var v339 int32 = (v135 + v135)
                                                                                                                                                        var v340 int32 = int32(v96)
                                                                                                                                                        v96 = 128
                                                                                                                                                        _ = v339
                                                                                                                                                        _ = v340
                                                                                                                                                    } else {
                                                                                                                                                        if false {
                                                                                                                                                            if len(v5) > 0 {
                                                                                                                                                                v5[0] = (v96 | (fi7 * fi7))
                                                                                                                                                            }
                                                                                                                                                            fmt.Println(fi7, v134)
                                                                                                                                                        } else {
                                                                                                                                                            v5 = append(v5, v8)
                                                                                                                                                            if ((p0 % (((p0 % (((p0 * 42)) | 1))) | 1)) > (p0 - p0)) {
                                                                                                                                                                switch (((p2 + "\n\t") + p2) + ((p2 + "") + ("" + ("\n\t" + (("hello" + p2) + ((p2 + p2) + p2)))))) {
                                                                                                                                                                case "a":
                                                                                                                                                                    v135 += 85
                                                                                                                                                                case "foo bar":
                                                                                                                                                                    v5 = append(v5, (v96 + -1))
                                                                                                                                                                case "\n\t":
                                                                                                                                                                    if len(v5) >= 2 {
                                                                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                                                                    }
                                                                                                                                                                default:
                                                                                                                                                                    v341 := GenS38{F0: 1}
                                                                                                                                                                    _ = v341
                                                                                                                                                                }
                                                                                                                                                                p0 += (^(p0))
                                                                                                                                                            }
                                                                                                                                                            v342 := GenS39{F0: 2, F1: 98}
                                                                                                                                                            _ = v342
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    v343 := GenS37{F0: v135, F1: 4294967295, F2: 1}
                                                                                                                                                    p2 = "foo bar"
                                                                                                                                                    if false {
                                                                                                                                                        var v344 float32 = (1e10 * float32(((0.1 * float64(3.14159)) / -2.25)))
                                                                                                                                                        fmt.Println(fi7, p0, v8)
                                                                                                                                                        fmt.Println(fi7)
                                                                                                                                                        var v345 uint8 = 1
                                                                                                                                                        _ = v344
                                                                                                                                                        _ = v345
                                                                                                                                                    }
                                                                                                                                                    var v346 int8 = -53
                                                                                                                                                    var v347 int = 101
                                                                                                                                                    v348 := GenS40{F0: int64(p1), F1: (fi7 ^ (fi6 << ((154) & 63)))}
                                                                                                                                                    _ = v343
                                                                                                                                                    _ = v346
                                                                                                                                                    _ = v347
                                                                                                                                                    _ = v348
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            _ = len(v5)
                                                                                                                                            fmt.Println(v8)
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        if len(v5) >= 2 {
                                                                                                                                            v5 = v5[1:len(v5)]
                                                                                                                                        }
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                if true {
                                                                                                                                    switch (v8 & fi6) {
                                                                                                                                    case 863:
                                                                                                                                        p0 -= (^(p0))
                                                                                                                                    case -48:
                                                                                                                                        if true {
                                                                                                                                            v349 := GenS8{F0: (p0 & (p0 * 762)), F1: 255, F2: 2}
                                                                                                                                            _ = v349
                                                                                                                                        }
                                                                                                                                    case 1:
                                                                                                                                        if true {
                                                                                                                                            v350 := GenS39{F0: 255, F1: 18446744073709551615}
                                                                                                                                            fmt.Println(fi6)
                                                                                                                                            v5 = append(v5, -2147483648)
                                                                                                                                            _ = v350
                                                                                                                                        } else {
                                                                                                                                            var v351 uint = 255
                                                                                                                                            switch p2 {
                                                                                                                                            case "foo bar":
                                                                                                                                                var v352 uint8 = 41
                                                                                                                                                _ = v352
                                                                                                                                            default:
                                                                                                                                                p0 += p0
                                                                                                                                                var v353 uint64 = 4294967295
                                                                                                                                                _ = v353
                                                                                                                                            }
                                                                                                                                            _ = v351
                                                                                                                                        }
                                                                                                                                    default:
                                                                                                                                        fmt.Println(v134)
                                                                                                                                    }
                                                                                                                                    if len(v5) > 0 {
                                                                                                                                        v5[0] = fi7
                                                                                                                                    }
                                                                                                                                    var v354 int64 = 42
                                                                                                                                    fmt.Println(fi7, v96, v8)
                                                                                                                                    _ = v354
                                                                                                                                }
                                                                                                                                if ((0 > 103) && false) {
                                                                                                                                    p0 = 2
                                                                                                                                    var v355 int8 = 0
                                                                                                                                    if len(v5) >= 2 {
                                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                                    }
                                                                                                                                    _ = v355
                                                                                                                                }
                                                                                                                                _ = v311
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        var v356 int16 = 1
                                                                                                        _ = v185
                                                                                                        _ = v356
                                                                                                    }
                                                                                                case -148:
                                                                                                    fmt.Println(p2)
                                                                                                case 571:
                                                                                                    v357 := GenS15{F0: 9223372036854775807}
                                                                                                    _ = v357
                                                                                                default:
                                                                                                    var v358 float64 = 0.0
                                                                                                    _ = v358
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        switch ((p2 + "\n\t") + p2) {
                                                                                        case "":
                                                                                            p2 = (p2 + p2)
                                                                                        case "\n\t":
                                                                                            fmt.Println(p2, v135)
                                                                                        }
                                                                                    } else {
                                                                                        v135 = (v135 + (v135 | (v135 << ((0) & 63))))
                                                                                    }
                                                                                    if true {
                                                                                        if true {
                                                                                            if len(v5) >= 2 {
                                                                                                v5 = v5[1:len(v5)]
                                                                                            }
                                                                                            v8 = (-(fi6))
                                                                                        } else {
                                                                                            var v359 int8 = 127
                                                                                            var v360 int64 = -223
                                                                                            _ = v359
                                                                                            _ = v360
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(fi7, p0, v8)
                                                                                        var v361 uint64 = 4294967295
                                                                                        _ = v361
                                                                                    }
                                                                                    switch ("" + p2) {
                                                                                    case "foo bar":
                                                                                        p2 = (p2 + "")
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                if len(v5) >= 2 {
                                                                                    v5 = v5[1:len(v5)]
                                                                                }
                                                                            }
                                                                            _ = v135
                                                                        } else {
                                                                            v96 -= (fi6 * fi6)
                                                                            var v362 uint32 = 4294967295
                                                                            if len(v5) > 0 {
                                                                                v5[0] = (^(v8))
                                                                            }
                                                                            v134 += 1e10
                                                                            _ = v362
                                                                        }
                                                                        _ = v134
                                                                    }
                                                                    fmt.Println(p2)
                                                                    v363 := GenS25{F0: (^(p0))}
                                                                    _ = v363
                                                                }
                                                                if (true && (2 < 1)) {
                                                                    v364 := GenS8{F0: 32767, F1: 2, F2: 0}
                                                                    switch (-(fi7)) {
                                                                    case 32767:
                                                                        if false {
                                                                            v8 = fi7
                                                                            fmt.Println(fi7, fi6)
                                                                            var v365 uint8 = 0
                                                                            if (65535 > 363) {
                                                                                fmt.Println(p2)
                                                                            } else {
                                                                                var v366 int = fi7
                                                                                var v367 int16 = 32767
                                                                                _ = v366
                                                                                _ = v367
                                                                            }
                                                                            _ = v365
                                                                        }
                                                                    case 127:
                                                                        switch p2 {
                                                                        case "foo bar":
                                                                            if (!(9223372036854775807 == 9223372036854775807)) {
                                                                                if (1e10 != float32(-2.25)) {
                                                                                    var v368 int8 = -128
                                                                                    _ = v368
                                                                                } else {
                                                                                    v369 := GenS40{F0: 9223372036854775807, F1: v8}
                                                                                    var v370 bool = false
                                                                                    _ = v369
                                                                                    _ = v370
                                                                                }
                                                                                if (383 == 278) {
                                                                                    var v371 int64 = -425
                                                                                    if false {
                                                                                        v372 := GenS41{F0: p1}
                                                                                        if true {
                                                                                            p1 = p1
                                                                                            switch fi7 {
                                                                                            case 364:
                                                                                                fmt.Println(v96, fi7, v96)
                                                                                                var v373 int8 = int8(p1)
                                                                                                _ = v373
                                                                                            case 128:
                                                                                                p2 = p2
                                                                                            case 9223372036854775807:
                                                                                                if true {
                                                                                                    v96 -= (fi7 & v8)
                                                                                                    fmt.Println(p0)
                                                                                                    var v374 int8 = -109
                                                                                                    _ = v374
                                                                                                } else {
                                                                                                    v375 := GenS42{F0: v371, F1: 0}
                                                                                                    var v376 int32 = 2147483647
                                                                                                    if false {
                                                                                                        fmt.Println(p2, v96, fi7)
                                                                                                    }
                                                                                                    _ = v375
                                                                                                    _ = v376
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v377 int32 = -2147483648
                                                                                        _ = v372
                                                                                        _ = v377
                                                                                    }
                                                                                    v378 := GenS33{F0: v371}
                                                                                    _ = v371
                                                                                    _ = v378
                                                                                }
                                                                                v379 := GenS43{F0: 42}
                                                                                _ = v379
                                                                            } else {
                                                                                var v380 float32 = float32(3.14159)
                                                                                _ = v380
                                                                            }
                                                                            if false {
                                                                                var v381 uint = 406
                                                                                fmt.Println(v8, p2)
                                                                                switch p2 {
                                                                                case "hello":
                                                                                    v96 -= int(p0)
                                                                                case "\n\t":
                                                                                    if (!false) {
                                                                                        var v382 string = ("a" + (p2 + p2))
                                                                                        _ = v382
                                                                                    }
                                                                                case "a":
                                                                                    p2 = (p2 + p2)
                                                                                default:
                                                                                    var v383 string = ""
                                                                                    v384 := GenS19{F0: 42, F1: 458, F2: 0}
                                                                                    p2 = v383
                                                                                    v385 := GenS18{F0: 0, F1: 42, F2: v381}
                                                                                    fmt.Println(p2, v8, v383)
                                                                                    v96 -= (fi7 * -2147483648)
                                                                                    var v386 uint = (v381 * (v381 - 0))
                                                                                    _ = v383
                                                                                    _ = v384
                                                                                    _ = v385
                                                                                    _ = v386
                                                                                }
                                                                                var v387 int64 = -899
                                                                                _ = v381
                                                                                _ = v387
                                                                            }
                                                                            v388 := GenS40{F0: -1, F1: (-(fi6))}
                                                                            _ = v388
                                                                        case "":
                                                                            if len(v5) >= 2 {
                                                                                v5 = v5[1:len(v5)]
                                                                            }
                                                                        case "hello":
                                                                            fmt.Println(p1, p1)
                                                                        }
                                                                        v389 := GenS44{F0: p1, F1: 65535}
                                                                        _ = v389
                                                                    default:
                                                                        switch p2 {
                                                                        case "\n\t":
                                                                            v8 -= (v96 * -518)
                                                                        case "foo bar":
                                                                            v96 -= fi6
                                                                        case "hello":
                                                                            p1 -= p1
                                                                        }
                                                                    }
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    _ = v364
                                                                } else {
                                                                    var v390 string = "a"
                                                                    p0 = (p0 ^ (p0 << ((0) & 63)))
                                                                    _ = v390
                                                                }
                                                            } else {
                                                                if true {
                                                                    if (4294967295 != 65535) {
                                                                        v96 += -1
                                                                        var v391 uint16 = 0
                                                                        switch p2 {
                                                                        case "a":
                                                                            var v392 uint8 = 103
                                                                            _ = v392
                                                                        case "":
                                                                            var v393 bool = (((18446744073709551615 != 1) || ((true || true) && (false && (-1 > 2147483647)))) || false)
                                                                            _ = v393
                                                                        default:
                                                                            p1 -= p1
                                                                            p2 = p2
                                                                            var v394 int64 = -9223372036854775808
                                                                            p0 += p0
                                                                            _ = v394
                                                                        }
                                                                        var v395 int32 = 1
                                                                        _ = v391
                                                                        _ = v395
                                                                    }
                                                                    var v396 bool = true
                                                                    var v397 uint8 = 249
                                                                    _ = v396
                                                                    _ = v397
                                                                } else {
                                                                    v398 := GenS39{F0: 0, F1: 981}
                                                                    p0 = (p0 / ((2) | 1))
                                                                    _ = v398
                                                                }
                                                                v96 = -276
                                                                switch (p2 + ((p2 + "foo bar") + "foo bar")) {
                                                                case "\n\t":
                                                                    if len(v5) > 0 {
                                                                        v5[0] = fi7
                                                                    }
                                                                case "hello":
                                                                    if true {
                                                                        var v399 int8 = -128
                                                                        if false {
                                                                            switch fi6 {
                                                                            case -352:
                                                                                var v400 uint32 = 1
                                                                                _ = v400
                                                                            case -338:
                                                                                var v401 uint8 = 42
                                                                                _ = v401
                                                                            default:
                                                                                v402 := GenS45{F0: 2, F1: 1, F2: (!true)}
                                                                                _ = v402
                                                                            }
                                                                            if true {
                                                                                v96 -= v8
                                                                            } else {
                                                                                v5 = append(v5, (v8 >> ((2) & 63)))
                                                                                if true {
                                                                                    if false {
                                                                                        _ = len(v5)
                                                                                        if (3.14159 < (((0.0 - (float64(0.0) / (1.5 * (1.5 / (1e10 * 1e10))))) - (3.14159 * 0.1)) * 0.0)) {
                                                                                            switch fi7 {
                                                                                            case 2147483647:
                                                                                                var v403 uint64 = 18446744073709551615
                                                                                                _ = v403
                                                                                            }
                                                                                        } else {
                                                                                            switch (fi7 & 128) {
                                                                                            case -2147483648:
                                                                                                switch (fi6 - 4294967295) {
                                                                                                case -1:
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = 4294967295
                                                                                                    }
                                                                                                case -2147483648:
                                                                                                    if (false && (p2 != p2)) {
                                                                                                        p1 = (p1 - p1)
                                                                                                        switch "hello" {
                                                                                                        case "a":
                                                                                                            if false {
                                                                                                                p0 = (p0 + (p0 - p0))
                                                                                                                p0 += p0
                                                                                                            } else {
                                                                                                                var v404 int8 = (v399 * (v399 - (v399 + v399)))
                                                                                                                var v405 int64 = 42
                                                                                                                var v406 int64 = (v405 % (((^(v405))) | 1))
                                                                                                                _ = v404
                                                                                                                _ = v405
                                                                                                                _ = v406
                                                                                                            }
                                                                                                            p2 = ((p2 + ((p2 + p2) + p2)) + "\n\t")
                                                                                                            p2 = (("a" + p2) + "\n\t")
                                                                                                        case "foo bar":
                                                                                                            v8 = (v8 >> ((4294967295) & 63))
                                                                                                        case "hello":
                                                                                                            if false {
                                                                                                                if (!false) {
                                                                                                                    _ = len(v5)
                                                                                                                    p1 = p1
                                                                                                                }
                                                                                                            } else {
                                                                                                                p2 = (p2 + (((p2 + p2) + p2) + "hello"))
                                                                                                                if true {
                                                                                                                    var v407 float32 = 0.1
                                                                                                                    _ = v407
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        if true {
                                                                                                                            p1 = 1000
                                                                                                                            v408 := GenS12{F0: 62, F1: p1}
                                                                                                                            _ = v408
                                                                                                                        } else {
                                                                                                                            fmt.Println(v399)
                                                                                                                            switch p2 {
                                                                                                                            case "":
                                                                                                                                if len(v5) >= 2 {
                                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                                }
                                                                                                                            case "a":
                                                                                                                                var v409 int8 = v399
                                                                                                                                p2 = p2
                                                                                                                                var v410 int64 = -189
                                                                                                                                _ = v409
                                                                                                                                _ = v410
                                                                                                                            default:
                                                                                                                                p2 = (p2 + p2)
                                                                                                                            }
                                                                                                                            var v411 uint = 1
                                                                                                                            _ = v411
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                p0 = (^(p0))
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v412 := GenS16{F0: 202}
                                                                                                    _ = v412
                                                                                                }
                                                                                            case -427:
                                                                                                var v413 int8 = (v399 & v399)
                                                                                                _ = v413
                                                                                            case -450:
                                                                                                fmt.Println(v399)
                                                                                            }
                                                                                            if (356 <= 18446744073709551615) {
                                                                                                var v414 uint = 2
                                                                                                _ = v414
                                                                                            } else {
                                                                                                fmt.Println(v8, v8, v96)
                                                                                                fmt.Println(fi7, p2, p1)
                                                                                                var v415 int16 = 42
                                                                                                _ = v415
                                                                                            }
                                                                                            p1 = p1
                                                                                            if true {
                                                                                                var v416 int64 = 22
                                                                                                p0 = p0
                                                                                                _ = v416
                                                                                            } else {
                                                                                                var v417 uint = 388
                                                                                                var v418 int16 = -690
                                                                                                v419 := GenS46{F0: 0, F1: 2, F2: (^(v8))}
                                                                                                _ = v417
                                                                                                _ = v418
                                                                                                _ = v419
                                                                                            }
                                                                                            var v420 int64 = 0
                                                                                            fmt.Println(v8, v96)
                                                                                            v421 := GenS18{F0: 414, F1: 86, F2: 1}
                                                                                            _ = v420
                                                                                            _ = v421
                                                                                        }
                                                                                        var v422 float32 = 0.0
                                                                                        _ = v422
                                                                                    } else {
                                                                                        v423 := GenS5{F0: 42}
                                                                                        _ = v423
                                                                                    }
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = (fi6 << ((2) & 63))
                                                                                    }
                                                                                    _ = len(v5)
                                                                                    fmt.Println(fi6, p2, v399)
                                                                                } else {
                                                                                    v399 += (-(v399))
                                                                                    if (fi6 == (v96 >> ((uint(p0)) & 63))) {
                                                                                        switch "a" {
                                                                                        case "\n\t":
                                                                                            var v424 float32 = 0.0
                                                                                            _ = v424
                                                                                        case "":
                                                                                            var v425 uint32 = 423
                                                                                            _ = v425
                                                                                        }
                                                                                        if false {
                                                                                            switch p2 {
                                                                                            case "\n\t":
                                                                                                v426 := GenS12{F0: 0, F1: p1}
                                                                                                _ = v426
                                                                                            case "a":
                                                                                                _ = len(v5)
                                                                                            }
                                                                                        } else {
                                                                                            v8 = fi6
                                                                                        }
                                                                                    } else {
                                                                                        v5 = append(v5, v96)
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = 4294967295
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        var v427 int = -2147483648
                                                                                        _ = v427
                                                                                    } else {
                                                                                        if true {
                                                                                            switch (fi6 >> ((92) & 63)) {
                                                                                            case -1:
                                                                                                switch 9223372036854775807 {
                                                                                                case 32767:
                                                                                                    switch "hello" {
                                                                                                    case "foo bar":
                                                                                                        p2 = p2
                                                                                                        p2 = "a"
                                                                                                    case "":
                                                                                                        if (((1.5 * ((0.1 + 3.14159) - ((((1.5 / (0.0 + 1.5)) * 1.5) * (1e10 - float64(0.1))) + 0.1))) + (1.5 / float64(-2.25))) >= (-2.25 * 0.0)) {
                                                                                                            var v428 int8 = v399
                                                                                                            var v429 int64 = 938
                                                                                                            v430 := GenS15{F0: (v429 | (v429 << ((0) & 63)))}
                                                                                                            _ = v428
                                                                                                            _ = v429
                                                                                                            _ = v430
                                                                                                        } else {
                                                                                                            var v431 int16 = (p1 ^ (p1 << ((uint(p0)) & 63)))
                                                                                                            var v432 uint = 686
                                                                                                            v399 = -14
                                                                                                            _ = v431
                                                                                                            _ = v432
                                                                                                        }
                                                                                                    default:
                                                                                                        switch p2 {
                                                                                                        case "\n\t":
                                                                                                            var v433 int32 = 2147483647
                                                                                                            _ = v433
                                                                                                        case "hello":
                                                                                                            v8 -= (fi7 - v96)
                                                                                                        case "foo bar":
                                                                                                            v399 = v399
                                                                                                        default:
                                                                                                            if false {
                                                                                                                var v434 uint16 = (^(p0))
                                                                                                                _ = v434
                                                                                                            } else {
                                                                                                                var v435 uint16 = (p0 + (p0 - (^(p0))))
                                                                                                                if (fi6 >= (v96 << ((4294967295) & 63))) {
                                                                                                                    var v436 uint64 = 4294967295
                                                                                                                    v5 = append(v5, v96)
                                                                                                                    _ = v436
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        v437 := GenS15{F0: 42}
                                                                                                                        _ = v437
                                                                                                                    }
                                                                                                                    var v438 uint = 2
                                                                                                                    _ = v438
                                                                                                                }
                                                                                                                var v439 uint8 = 2
                                                                                                                _ = v435
                                                                                                                _ = v439
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                case 0:
                                                                                                    switch "foo bar" {
                                                                                                    case "hello":
                                                                                                        p0 += (^(p0))
                                                                                                    case "\n\t":
                                                                                                        v8 -= 100
                                                                                                    case "":
                                                                                                        if (false || false) {
                                                                                                            v440 := GenS12{F0: 0, F1: 32767}
                                                                                                            v399 = -128
                                                                                                            _ = v440
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    p0 = p0
                                                                                                }
                                                                                            case -2147483648:
                                                                                                fmt.Println(p2)
                                                                                            }
                                                                                        } else {
                                                                                            v441 := GenS47{F0: p1}
                                                                                            _ = v441
                                                                                        }
                                                                                    }
                                                                                }
                                                                                p2 = (("" + (p2 + p2)) + "\n\t")
                                                                            }
                                                                            p0 -= (p0 & (^(p0)))
                                                                        } else {
                                                                            fmt.Println(v96, p1)
                                                                        }
                                                                        _ = v399
                                                                    }
                                                                case "":
                                                                    fmt.Println(v8)
                                                                }
                                                                v442 := GenS48{F0: "", F1: -32768, F2: -1}
                                                                var v443 string = p2
                                                                v8 -= (-(v96))
                                                                _ = v442
                                                                _ = v443
                                                            }
                                                        }
                                                        v8 = 9223372036854775807
                                                    }
                                                    if true {
                                                        fmt.Println(p1)
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        v96 = v8
                                                        fmt.Println(fi6, p1, p0)
                                                    } else {
                                                        if (915 == 4294967295) {
                                                            p2 = p2
                                                        } else {
                                                            var v444 int = fi7
                                                            p1 -= 1000
                                                            var v445 int = -883
                                                            _ = v444
                                                            _ = v445
                                                        }
                                                        if false {
                                                            p1 = p1
                                                            switch (p2 + "hello") {
                                                            case "foo bar":
                                                                p2 = "hello"
                                                            case "a":
                                                                switch (fi7 ^ (fi7 << ((18446744073709551615) & 63))) {
                                                                case 32767:
                                                                    p2 = (p2 + p2)
                                                                }
                                                            case "":
                                                                if true {
                                                                    if false {
                                                                        if len(v5) >= 2 {
                                                                            v5 = v5[1:len(v5)]
                                                                        }
                                                                        var v446 uint64 = 4294967295
                                                                        _ = v446
                                                                    }
                                                                    v8 = (v8 << ((42) & 63))
                                                                    var v447 int64 = 56
                                                                    _ = v447
                                                                } else {
                                                                    p1 -= (p1 + (^(p1)))
                                                                }
                                                            default:
                                                                switch p2 {
                                                                case "":
                                                                    var v448 float64 = 0.0
                                                                    _ = v448
                                                                default:
                                                                    _ = len(v5)
                                                                }
                                                            }
                                                        } else {
                                                            var v449 float64 = 1e10
                                                            switch (v96 / (((fi7 - fi7)) | 1)) {
                                                            case -397:
                                                                var v450 int8 = 0
                                                                p1 -= p1
                                                                _ = v450
                                                            default:
                                                                if false {
                                                                    v451 := GenS13{F0: -2147483648}
                                                                    var v452 int = (fi7 << ((2) & 63))
                                                                    v5 = append(v5, (-(v96)))
                                                                    _ = v451
                                                                    _ = v452
                                                                } else {
                                                                    v453 := GenS24{F0: -102, F1: 42}
                                                                    switch 32768 {
                                                                    case 32767:
                                                                        var v454 int = (fi6 + 128)
                                                                        _ = v454
                                                                    case -1:
                                                                        var v455 string = p2
                                                                        _ = v455
                                                                    case -825:
                                                                        v96 += (fi6 ^ v96)
                                                                    }
                                                                    _ = v453
                                                                }
                                                            }
                                                            fmt.Println(p0, p0)
                                                            _ = v449
                                                        }
                                                        switch (v8 / ((fi7) | 1)) {
                                                        case -2147483648:
                                                            if false {
                                                                var v456 int = 32767
                                                                if (-2147483648 > -803) {
                                                                    switch 1 {
                                                                    case 1:
                                                                        fmt.Println(p2, fi6)
                                                                    case 9223372036854775807:
                                                                        if true {
                                                                            v8 = (v96 & (v8 * v8))
                                                                        } else {
                                                                            var v457 bool = (true || (0 == -610))
                                                                            var v458 int8 = 42
                                                                            _ = v457
                                                                            _ = v458
                                                                        }
                                                                    case -896:
                                                                        var v459 bool = (!(!true))
                                                                        _ = v459
                                                                    default:
                                                                        v8 -= fi6
                                                                    }
                                                                } else {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    if true {
                                                                        switch (v8 & (fi7 - v96)) {
                                                                        case -9223372036854775808:
                                                                            var v460 int16 = 1000
                                                                            _ = v460
                                                                        case 1:
                                                                            if (!(!(812 != 4294967295))) {
                                                                                p0 += 2
                                                                                var v461 int32 = -55
                                                                                _ = v461
                                                                            }
                                                                        }
                                                                    } else {
                                                                        switch 1 {
                                                                        case 32768:
                                                                            p0 = (p0 & (^(p0)))
                                                                        default:
                                                                            if (0.1 < -2.25) {
                                                                                switch (fi6 + fi7) {
                                                                                case 746:
                                                                                    if (!(uint32(p0) != 0)) {
                                                                                        v462 := GenS49{F0: (-(fi6)), F1: 2}
                                                                                        _ = v462
                                                                                    }
                                                                                case 405:
                                                                                    if (1.5 >= 3.14159) {
                                                                                        var v463 int32 = int32(v96)
                                                                                        fmt.Println(v456)
                                                                                        _ = v463
                                                                                    } else {
                                                                                        fmt.Println(fi6, fi7, fi6)
                                                                                        v464 := GenS13{F0: 277}
                                                                                        if true {
                                                                                            p0 += p0
                                                                                            if false {
                                                                                                if true {
                                                                                                    if true {
                                                                                                        var v465 uint32 = 817
                                                                                                        _ = v465
                                                                                                    } else {
                                                                                                        if false {
                                                                                                            switch (-(fi6)) {
                                                                                                            case 128:
                                                                                                                var v466 uint = 1
                                                                                                                _ = v466
                                                                                                            case -9223372036854775808:
                                                                                                                fmt.Println(p2, p2, v96)
                                                                                                            }
                                                                                                            v456 = (v456 << ((406) & 63))
                                                                                                        } else {
                                                                                                            if true {
                                                                                                                _ = len(v5)
                                                                                                            } else {
                                                                                                                var v467 int8 = -11
                                                                                                                p0 = p0
                                                                                                                if (0 == 0) {
                                                                                                                    v468 := GenS50{F0: v8}
                                                                                                                    _ = v468
                                                                                                                }
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                                _ = v467
                                                                                                            }
                                                                                                            var v469 string = p2
                                                                                                            if (1.5 >= float32((0.0 - 1.5))) {
                                                                                                                switch v469 {
                                                                                                                case "a":
                                                                                                                    if len(v5) >= 2 {
                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                    }
                                                                                                                    if len(v5) >= 2 {
                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                    }
                                                                                                                case "\n\t":
                                                                                                                    if false {
                                                                                                                        fmt.Println(v96)
                                                                                                                        switch fi6 {
                                                                                                                        case 4294967295:
                                                                                                                            if (!(42 != 451)) {
                                                                                                                                if false {
                                                                                                                                    v470 := GenS51{F0: p1, F1: 42}
                                                                                                                                    _ = v470
                                                                                                                                }
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            fmt.Println(p1, v96)
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                                p0 += p0
                                                                                                                fmt.Println(p2, fi7, p2)
                                                                                                            } else {
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                                var v471 int16 = 1000
                                                                                                                _ = v471
                                                                                                            }
                                                                                                            _ = v469
                                                                                                        }
                                                                                                    }
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                }
                                                                                                switch ((p2 + p2) + (p2 + "")) {
                                                                                                case "a":
                                                                                                    v456 -= fi7
                                                                                                }
                                                                                                var v472 float64 = (-2.25 / (1e10 / ((-2.25 - (float64(0.0) * 0.0)) * float64(-2.25))))
                                                                                                p0 = (^(p0))
                                                                                                fmt.Println(fi6)
                                                                                                _ = v472
                                                                                            } else {
                                                                                                if (!false) {
                                                                                                    var v473 uint64 = 18446744073709551615
                                                                                                    _ = v473
                                                                                                } else {
                                                                                                    p0 += p0
                                                                                                }
                                                                                            }
                                                                                            var v474 uint = 0
                                                                                            _ = v474
                                                                                        }
                                                                                        _ = v464
                                                                                    }
                                                                                }
                                                                                v475 := GenS52{F0: 255}
                                                                                switch (^(fi6)) {
                                                                                case 9223372036854775807:
                                                                                    switch 42 {
                                                                                    case 434:
                                                                                        var v476 uint = 0
                                                                                        _ = v476
                                                                                    case 32767:
                                                                                        var v477 int = 4294967295
                                                                                        _ = v477
                                                                                    }
                                                                                case -9223372036854775808:
                                                                                    _ = len(v5)
                                                                                    p2 = "hello"
                                                                                    switch p2 {
                                                                                    case "foo bar":
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            var v478 uint16 = p0
                                                                                            _ = v478
                                                                                        case "a":
                                                                                            if true {
                                                                                                if len(v5) > 0 {
                                                                                                    v5[0] = fi6
                                                                                                }
                                                                                                var v479 float32 = ((1e10 / (1.5 / ((1.5 + (1.5 * -2.25)) + (-2.25 - ((1.5 * (0.0 + 3.14159)) - (0.0 / (1e10 - 0.0))))))) * float32((0.1 - 0.1)))
                                                                                                switch p2 {
                                                                                                case "hello":
                                                                                                    if len(v5) > 0 {
                                                                                                        v5[0] = 480
                                                                                                    }
                                                                                                    var v480 float64 = 1e10
                                                                                                    _ = v480
                                                                                                case "":
                                                                                                    var v481 int32 = 2147483647
                                                                                                    _ = v481
                                                                                                case "foo bar":
                                                                                                    v5 = append(v5, 127)
                                                                                                default:
                                                                                                    fmt.Println(v8, v8, p1)
                                                                                                }
                                                                                                _ = v479
                                                                                            } else {
                                                                                                fmt.Println(v8)
                                                                                                var v482 int16 = (p1 * p1)
                                                                                                var v483 uint = 1
                                                                                                v482 = p1
                                                                                                if (!(((1e10 / -2.25) * (-2.25 + 1e10)) <= float64(0.1))) {
                                                                                                    var v484 uint32 = 514
                                                                                                    if false {
                                                                                                        v485 := GenS4{F0: 0, F1: 18446744073709551615}
                                                                                                        _ = v485
                                                                                                    }
                                                                                                    v456 -= v456
                                                                                                    p0 = p0
                                                                                                    _ = v484
                                                                                                }
                                                                                                _ = v482
                                                                                                _ = v483
                                                                                            }
                                                                                            if (!false) {
                                                                                                v486 := GenS53{F0: p1, F1: 84, F2: 32767}
                                                                                                fmt.Println(p2)
                                                                                                _ = v486
                                                                                            } else {
                                                                                                p0 += p0
                                                                                                var v487 int = (v456 * (-(v456)))
                                                                                                v8 -= v8
                                                                                                _ = v487
                                                                                            }
                                                                                        default:
                                                                                            var v488 int8 = -128
                                                                                            _ = v488
                                                                                        }
                                                                                    case "":
                                                                                        if false {
                                                                                            v5 = append(v5, (-(v8)))
                                                                                            v489 := GenS34{F0: 255}
                                                                                            _ = v489
                                                                                        } else {
                                                                                            v490 := GenS49{F0: 32767, F1: 227}
                                                                                            _ = v490
                                                                                        }
                                                                                    default:
                                                                                        if (1 > 0) {
                                                                                            var v491 uint = 780
                                                                                            _ = v491
                                                                                        }
                                                                                    }
                                                                                    if false {
                                                                                        v492 := GenS22{F0: 4294967295, F1: 4294967295, F2: -1}
                                                                                        _ = v492
                                                                                    } else {
                                                                                        var v493 bool = (p2 <= p2)
                                                                                        if false {
                                                                                            switch (("" + "hello") + "\n\t") {
                                                                                            case "hello":
                                                                                                var v494 bool = (1 < 863)
                                                                                                _ = v494
                                                                                            }
                                                                                            p2 = (p2 + (p2 + "hello"))
                                                                                            var v495 uint = 0
                                                                                            switch ("foo bar" + "foo bar") {
                                                                                            case "\n\t":
                                                                                                v5 = append(v5, 9223372036854775807)
                                                                                            }
                                                                                            _ = v495
                                                                                        }
                                                                                        _ = v493
                                                                                    }
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = v96
                                                                                    }
                                                                                    fmt.Println(p0, p2)
                                                                                    if false {
                                                                                        p0 = p0
                                                                                        var v496 float32 = (0.0 / 1.5)
                                                                                        _ = v496
                                                                                    } else {
                                                                                        p2 = ("" + "hello")
                                                                                    }
                                                                                    if false {
                                                                                        switch (fi7 + (fi7 * v8)) {
                                                                                        case -9223372036854775808:
                                                                                            if false {
                                                                                                v497 := GenS1{F0: (p2 + (p2 + p2)), F1: 1}
                                                                                                var v498 uint16 = (p0 - (p0 / ((p0) | 1)))
                                                                                                _ = v497
                                                                                                _ = v498
                                                                                            } else {
                                                                                                p1 = p1
                                                                                                var v499 bool = false
                                                                                                switch p2 {
                                                                                                case "foo bar":
                                                                                                    v96 += (v456 + v8)
                                                                                                case "":
                                                                                                    if len(v5) >= 2 {
                                                                                                        v5 = v5[1:len(v5)]
                                                                                                    }
                                                                                                default:
                                                                                                    v500 := GenS54{F0: 2, F1: 6, F2: (100 < 0)}
                                                                                                    _ = v500
                                                                                                }
                                                                                                _ = v499
                                                                                            }
                                                                                        default:
                                                                                            v501 := GenS1{F0: p2, F1: 2147483648}
                                                                                            _ = v501
                                                                                        }
                                                                                    } else {
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = fi7
                                                                                        }
                                                                                        var v502 string = p2
                                                                                        _ = v502
                                                                                    }
                                                                                }
                                                                                _ = v475
                                                                            }
                                                                        }
                                                                        v503 := GenS19{F0: 1, F1: p0, F2: 42}
                                                                        _ = v503
                                                                    }
                                                                    v504 := GenS55{F0: -1}
                                                                    _ = v504
                                                                }
                                                                fmt.Println(p2, fi7, v96)
                                                                var v505 uint = 289
                                                                _ = v456
                                                                _ = v505
                                                            } else {
                                                                fmt.Println(v96, p0)
                                                                var v506 uint8 = 68
                                                                _ = v506
                                                            }
                                                        case -1:
                                                            var v507 string = (p2 + "")
                                                            v508 := GenS41{F0: (p1 | (-(p1)))}
                                                            if true {
                                                                if true {
                                                                    fmt.Println(fi7, p1)
                                                                    fmt.Println(v507, v96, v96)
                                                                }
                                                            } else {
                                                                var v509 bool = false
                                                                var v510 uint32 = 42
                                                                var v511 int16 = (p1 + 722)
                                                                fmt.Println(fi6, v510)
                                                                _ = v509
                                                                _ = v510
                                                                _ = v511
                                                            }
                                                            _ = v507
                                                            _ = v508
                                                        }
                                                    }
                                                    _ = v96
                                                case "foo bar":
                                                    v8 = fi7
                                                    var v512 int64 = 9223372036854775807
                                                    p1 = p1
                                                    var v513 uint32 = 2147483647
                                                    switch p2 {
                                                    case "":
                                                        var v514 bool = false
                                                        _ = v514
                                                    case "\n\t":
                                                        var v515 int16 = (p1 << ((0) & 63))
                                                        _ = v515
                                                    case "foo bar":
                                                        var v516 int16 = p1
                                                        _ = v516
                                                    }
                                                    fmt.Println(v513, fi6, p0)
                                                    var v517 float32 = 1e10
                                                    var v518 bool = (!(v8 != 32767))
                                                    _ = v512
                                                    _ = v513
                                                    _ = v517
                                                    _ = v518
                                                case "hello":
                                                    if ((p1 % ((1) | 1)) == p1) {
                                                        v5 = append(v5, (v8 & -165))
                                                        var v519 int8 = 42
                                                        _ = v519
                                                    }
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = v8
                                                }
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = (v8 + fi6)
                                                }
                                                if false {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                } else {
                                                    switch (^(fi6)) {
                                                    case -9223372036854775808:
                                                        fmt.Println(fi7)
                                                    }
                                                }
                                                v520 := GenS56{F0: (^(p0)), F1: 107, F2: 1}
                                                _ = v520
                                            }
                                        case "":
                                            if (1 < 897) {
                                                var v521 int = 760
                                                _ = v521
                                            } else {
                                                switch (-(fi7)) {
                                                case 2147483647:
                                                    fmt.Println(p1, fi7, p0)
                                                case 926:
                                                    if true {
                                                        if ((1e10 <= 1.5) || (!(705 < -2147483648))) {
                                                            v5 = append(v5, v8)
                                                        } else {
                                                            v522 := GenS57{F0: true, F1: (p2 + p2)}
                                                            switch 240 {
                                                            case 127:
                                                                _ = len(v5)
                                                            case -102:
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            default:
                                                                if len(v5) > 0 {
                                                                    v5[0] = (fi6 / ((fi6) | 1))
                                                                }
                                                            }
                                                            fmt.Println(v8)
                                                            fmt.Println(p2, p1)
                                                            p0 += p0
                                                            _ = v522
                                                        }
                                                    }
                                                default:
                                                    p1 -= (^(p1))
                                                }
                                            }
                                        default:
                                            if (1 == 2147483647) {
                                                switch (v8 * (fi6 % (((fi6 >> ((167) & 63))) | 1))) {
                                                case 9223372036854775807:
                                                    v8 = fi6
                                                case -958:
                                                    fmt.Println(fi6, p2, p1)
                                                }
                                                fmt.Println(p2)
                                                v523 := GenS58{F0: p1}
                                                if (42 != 2) {
                                                    fmt.Println(fi7, fi7)
                                                    switch 591 {
                                                    case 501:
                                                        var v524 uint16 = p0
                                                        _ = v524
                                                    case 614:
                                                        v525 := GenS53{F0: (p1 + -707), F1: -1, F2: -215}
                                                        _ = v525
                                                    default:
                                                        p1 = (p1 / ((p1) | 1))
                                                    }
                                                    var v526 bool = ((!((v8 % (((v8 << ((65535) & 63))) | 1)) >= (fi6 | 1))) || (!false))
                                                    _ = v526
                                                } else {
                                                    switch -215 {
                                                    case 279:
                                                        v527 := GenS22{F0: 470, F1: 4294967295, F2: -128}
                                                        _ = v527
                                                    case -1:
                                                        if false {
                                                            if true {
                                                                v528 := GenS2{F0: 85, F1: 65535}
                                                                p0 = (^(p0))
                                                                _ = v528
                                                            }
                                                            _ = len(v5)
                                                        }
                                                        var v529 uint16 = p0
                                                        _ = v529
                                                    }
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    v530 := GenS17{F0: -2147483648}
                                                    fmt.Println(p2, p0)
                                                    _ = v530
                                                }
                                                p0 = 1
                                                _ = v523
                                            }
                                        }
                                    }
                                }
                                p0 = (p0 >> ((0) & 63))
                                _ = v10
                                _ = v9
                                _ = v92
                            } else {
                                if (!((0.0 <= 0.1) || true)) {
                                    p2 = (p2 + "a")
                                }
                            }
                            _ = len(v5)
                        }
                    }
                    v531 := GenS20{F0: (^(v8))}
                    _ = v531
                    _ = v8
                }
                fmt.Println(p0, p1)
                p0 += 950
                _ = fi7
            }
            p0 = p0
            p2 = p2
        }
        _ = fi6
    }
    for fi532 := 0; fi532 < 8; fi532++ {
        p1 = p1
        _ = fi532
    }
    for fi533 := 0; fi533 < 14; fi533++ {
        switch -9223372036854775808 {
        case 595:
            if (false && false) {
                var v534 float32 = 0.1
                if false {
                    p0 -= 2
                }
                v535 := GenS35{F0: 1, F1: p1}
                _ = v534
                _ = v535
            } else {
                if false {
                    var v536 uint16 = (p0 ^ p0)
                    for fi537 := 0; fi537 < 10; fi537++ {
                        fmt.Println(p2)
                        var v538 int8 = -47
                        switch 42 {
                        case 2147483647:
                            v538 = v538
                        case 0:
                            switch (fi533 << ((18446744073709551615) & 63)) {
                            case 32767:
                                if (850 == (p1 % (((p1 | p1)) | 1))) {
                                    p2 = ((p2 + p2) + "a")
                                } else {
                                    v539 := GenS59{F0: 0, F1: false, F2: (p0 | v536)}
                                    _ = v539
                                }
                            default:
                                fmt.Println(p2, p2)
                            }
                        case -637:
                            switch p2 {
                            case "\n\t":
                                v540 := GenS60{F0: (!true), F1: -2147483648}
                                _ = v540
                            case "hello":
                                v538 = 127
                            default:
                                var v541 uint = 255
                                _ = v541
                            }
                            var v542 uint64 = 18446744073709551615
                            _ = v542
                        default:
                            if (p2 < "hello") {
                                if (((!false) && true) || (4294967295 >= 1)) {
                                    if ((325 <= 255) && true) {
                                        var v543 uint16 = p0
                                        var v544 int16 = p1
                                        _ = v543
                                        _ = v544
                                    }
                                    v545 := GenS61{F0: p1, F1: -2147483648, F2: -533}
                                    p2 = p2
                                    _ = v545
                                } else {
                                    v546 := GenS45{F0: 2, F1: 765, F2: false}
                                    var v547 int8 = (v538 / (((v538 - (v538 % ((127) | 1)))) | 1))
                                    v548 := GenS62{F0: 2, F1: p0, F2: 1}
                                    var v549 int64 = -9223372036854775808
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    _ = v546
                                    _ = v547
                                    _ = v548
                                    _ = v549
                                }
                                v550 := GenS46{F0: 737, F1: 127, F2: fi533}
                                if ((p2 + p2) > ("hello" + "hello")) {
                                    if (!(false && (!(false || ((fi537 <= fi533) && true))))) {
                                        var v551 uint32 = 604
                                        _ = v551
                                    }
                                    var v552 bool = ((p2 + (("a" + (p2 + p2)) + p2)) != p2)
                                    v553 := GenS63{F0: 128, F1: (18446744073709551615 == 255), F2: false}
                                    _ = v552
                                    _ = v553
                                } else {
                                    fmt.Println(p2, p0, p0)
                                    var v554 int32 = int32(v536)
                                    _ = v554
                                }
                                _ = v550
                            } else {
                                v555 := GenS64{F0: int8(fi533)}
                                _ = v555
                            }
                        }
                        var v556 bool = false
                        p2 = p2
                        if (!((fi533 == fi537) && (!false))) {
                            v538 = (v538 / (((-(v538))) | 1))
                            if len(v5) > 0 {
                                v5[0] = fi537
                            }
                            fmt.Println(fi533, p2, p1)
                        }
                        _ = fi537
                        _ = v538
                        _ = v556
                    }
                    _ = v536
                } else {
                    v557 := GenS65{F0: (p0 % ((42) | 1)), F1: 9223372036854775807}
                    _ = v557
                }
                if false {
                    p2 = "a"
                } else {
                    if false {
                        var v558 uint64 = 18446744073709551615
                        _ = v558
                    } else {
                        var v559 uint16 = (p0 * p0)
                        _ = v559
                    }
                }
            }
        case 0:
            for fi560 := 0; fi560 < 18; fi560++ {
                var v561 uint8 = 0
                v562 := GenS40{F0: -1, F1: 42}
                p1 -= (^(p1))
                _ = fi560
                _ = v561
                _ = v562
            }
        }
        _ = fi533
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p0 << ((2) & 63))
}

func main() {
    var v1 bool = false
    v2 := make(map[int]bool)
    var v3 string = "a"
    v4 := make(map[string]string)
    v3 = v3
    v563 := GenS3{F0: 65535, F1: -513}
    switch v3 {
    case "hello":
        var v564 string = v3
        _ = v564
    case "\n\t":
        var v565 uint8 = 228
        _ = v565
    }
    v566 := v4["a"]
    for fi567 := 0; fi567 < 20; fi567++ {
        var v568 uint8 = 0
        v569 := GenS63{F0: v568, F1: v1, F2: (590 != 374)}
        v1 = (929 == fi567)
        _ = fi567
        _ = v568
        _ = v569
    }
    var v570 int16 = 213
    var v571 uint16 = 42
    switch 2147483647 {
    case 128:
        v566 = "foo bar"
    case 1:
        fmt.Println(v566)
    }
    v3 = (((v566 + (v3 + "")) + (v3 + v566)) + v566)
    switch v566 {
    case "a":
        if v1 {
            v4[v566] = v566
        } else {
            var v572 bool = false
            var v573 string = (v566 + ("hello" + ("\n\t" + "hello")))
            fmt.Println(v571, v1, v570)
            _ = v572
            _ = v573
        }
    }
    v4[""] = (v3 + v3)
    if false {
        v2[32767] = ((4294967295 >= 767) && ((((v571 << ((255) & 63)) <= genFn0(v571, v570, "")) && (4294967295 < 4294967295)) || (0 <= 18446744073709551615)))
        v571 = genFn0(genFn0(255, 1, "\n\t"), (v570 ^ (v570 * (v570 + (v570 | v570)))), (v3 + (v3 + v566)))
        if v1 {
            v574 := GenS66{F0: v566, F1: -2147483648}
            var v575 uint32 = 277
            if true {
                v570 = (^(v570))
                v575 -= (v575 + (^(v575)))
                for fi576 := 0; fi576 < 15; fi576++ {
                    for fi577 := 0; fi577 < 15; fi577++ {
                        v578 := GenS29{F0: int16(v575), F1: 65535}
                        if true {
                            v575 += v575
                        }
                        if (false || (true || (fi576 != (fi576 >> ((18446744073709551615) & 63))))) {
                            v1 = v1
                            var v579 uint32 = 1
                            _ = v579
                        }
                        _ = fi577
                        _ = v578
                    }
                    var v580 bool = false
                    _ = len(v2)
                    _ = fi576
                    _ = v580
                }
            } else {
                if v1 {
                    if v1 {
                        v571 += v571
                    } else {
                        v3 = v3
                    }
                    if (705 >= 628) {
                        fmt.Println(v575)
                        if (-460 != -2147483648) {
                            var v581 int8 = 42
                            for fi582 := 0; fi582 < 14; fi582++ {
                                if true {
                                    var v583 bool = v1
                                    v570 -= (v570 & (-(v570)))
                                    _ = v583
                                }
                                for fi584 := 0; fi584 < 17; fi584++ {
                                    v566 = v3
                                    v570 -= v570
                                    _ = fi584
                                }
                                _ = fi582
                            }
                            _ = v581
                        } else {
                            var v585 float64 = ((float64(0.0) + (-2.25 * 1e10)) + 0.0)
                            _ = v585
                        }
                        if ((-(v570)) != (-(v570))) {
                            if (509 >= -4) {
                                switch "hello" {
                                case "":
                                    switch int(v570) {
                                    case 1:
                                        v586 := []int{32767}
                                        _ = v586
                                    default:
                                        v587 := []int{}
                                        _ = v587
                                    }
                                case "hello":
                                    v588 := []int{}
                                    _ = v588
                                case "\n\t":
                                    switch 231 {
                                    case 0:
                                        v1 = v1
                                    case 301:
                                        var v589 int16 = v570
                                        _ = v589
                                    default:
                                        if ((v570 >> ((0) & 63)) < v570) {
                                            if ((42 < 1) && v1) {
                                                for fi590 := 0; fi590 < 14; fi590++ {
                                                    var v591 int8 = 1
                                                    switch v3 {
                                                    case "\n\t":
                                                        v566 = ""
                                                    case "hello":
                                                        v592 := v4[(("" + v3) + v566)]
                                                        _ = v592
                                                    }
                                                    _ = fi590
                                                    _ = v591
                                                }
                                                v593 := []int{-540, 11, 4294967295, 2147483647}
                                                _ = v593
                                            } else {
                                                for fi594 := 0; fi594 < 11; fi594++ {
                                                    var v595 int64 = 0
                                                    for fi596 := 0; fi596 < 6; fi596++ {
                                                        fmt.Println(fi596, fi596, v570)
                                                        switch (fi596 + (fi596 / (((fi594 | (-(fi594)))) | 1))) {
                                                        case 9223372036854775807:
                                                            var v597 float64 = 3.14159
                                                            _ = v597
                                                        default:
                                                            v598 := GenS30{F0: 4294967295}
                                                            _ = v598
                                                        }
                                                        _ = fi596
                                                    }
                                                    _ = fi594
                                                    _ = v595
                                                }
                                                var v599 int32 = 0
                                                v600 := GenS41{F0: (v570 - v570)}
                                                _ = v599
                                                _ = v600
                                            }
                                            for fi601 := 0; fi601 < 2; fi601++ {
                                                v4["foo bar"] = (("\n\t" + "a") + "foo bar")
                                                fmt.Println(v1, v575, v571)
                                                var v602 int8 = -128
                                                _ = fi601
                                                _ = v602
                                            }
                                            var v603 float32 = float32(0.1)
                                            _ = v603
                                        }
                                    }
                                default:
                                    if v1 {
                                        switch 0 {
                                        case 0:
                                            switch -9223372036854775808 {
                                            case 2147483647:
                                                v604 := []int{}
                                                _ = v604
                                            case 1:
                                                var v605 int8 = -98
                                                _ = v605
                                            }
                                        default:
                                            v606 := GenS48{F0: "", F1: (v570 % (((v570 & (v570 + (v570 | (v570 << ((1) & 63)))))) | 1)), F2: 9223372036854775807}
                                            _ = v606
                                        }
                                    } else {
                                        v2[-825] = true
                                    }
                                }
                            } else {
                                switch (("foo bar" + (v566 + "")) + (v566 + v3)) {
                                case "":
                                    if (0 == (v575 << ((42) & 63))) {
                                        delete(v2, 1)
                                        v570 = 62
                                    } else {
                                        var v607 int8 = 90
                                        _ = v607
                                    }
                                default:
                                    v608 := []int{128, 4294967295}
                                    _ = v608
                                }
                                var v609 int = -235
                                _ = v609
                            }
                            v3 = v3
                            v1 = v1
                        } else {
                            v575 = 322
                        }
                    } else {
                        for fi610 := 0; fi610 < 15; fi610++ {
                            v570 += 32767
                            _ = fi610
                        }
                        if (1 == 0) {
                            for fi611 := 0; fi611 < 3; fi611++ {
                                switch "\n\t" {
                                case "\n\t":
                                    v3 = v566
                                default:
                                    var v612 uint = 255
                                    _ = v612
                                }
                                var v613 uint16 = v571
                                _ = fi611
                                _ = v613
                            }
                        } else {
                            if (v1 && (((v575 << ((4294967295) & 63)) >= (v575 * uint32(v571))) && (1 < -1))) {
                                if (128 <= 42) {
                                    for fi614 := 0; fi614 < 20; fi614++ {
                                        v615 := []int{}
                                        v616 := GenS67{F0: 1, F1: 1}
                                        _ = fi614
                                        _ = v615
                                        _ = v616
                                    }
                                } else {
                                    if v1 {
                                        delete(v2, 9223372036854775807)
                                        switch -9223372036854775808 {
                                        case -137:
                                            var v617 int = -409
                                            _ = v617
                                        }
                                        switch v3 {
                                        case "hello":
                                            var v618 bool = v1
                                            _ = v618
                                        }
                                    } else {
                                        var v619 int8 = -128
                                        for fi620 := 0; fi620 < 11; fi620++ {
                                            v621 := []int{}
                                            _ = fi620
                                            _ = v621
                                        }
                                        v622 := v4[(v3 + "a")]
                                        _ = v619
                                        _ = v622
                                    }
                                }
                            } else {
                                for fi623 := 0; fi623 < 5; fi623++ {
                                    for fi624 := 0; fi624 < 5; fi624++ {
                                        if false {
                                            switch "foo bar" {
                                            case "a":
                                                v625 := []int{fi623, 168}
                                                _ = v625
                                            case "\n\t":
                                                v571 = (v571 | (v571 % ((genFn0((v571 % (((v571 - 470)) | 1)), 42, v566)) | 1)))
                                            case "":
                                                if true {
                                                    var v626 float64 = 1.5
                                                    _ = v626
                                                }
                                            }
                                        } else {
                                            switch (fi624 | fi624) {
                                            case 0:
                                                if (9223372036854775807 < 1) {
                                                    if v1 {
                                                        if ((v1 && (!v1)) || v1) {
                                                            if v1 {
                                                                var v627 uint32 = 0
                                                                _ = v627
                                                            }
                                                            if false {
                                                                if false {
                                                                    v628 := GenS68{F0: (float32(float64(0.0)) < -2.25)}
                                                                    switch v566 {
                                                                    case "a":
                                                                        v1 = (4294967295 != 4294967295)
                                                                    case "hello":
                                                                        v566 = "\n\t"
                                                                    }
                                                                    if v1 {
                                                                        switch ("hello" + ((((((v566 + v3) + (v566 + "")) + v566) + v3) + "") + v3)) {
                                                                        case "":
                                                                            if false {
                                                                                var v629 float32 = 0.1
                                                                                if ((false && v1) && false) {
                                                                                    switch ((v3 + (v3 + "foo bar")) + (v3 + ("hello" + v566))) {
                                                                                    case "\n\t":
                                                                                        v575 -= (v575 * 4294967295)
                                                                                        fmt.Println(fi624, v575)
                                                                                        var v630 uint8 = 128
                                                                                        _ = v630
                                                                                    default:
                                                                                        var v631 int8 = 42
                                                                                        _ = v631
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi623)
                                                                                    switch fi624 {
                                                                                    case -2147483648:
                                                                                        v632 := []int{(fi623 + (^(fi623)))}
                                                                                        _ = v632
                                                                                    case -9223372036854775808:
                                                                                        v633 := GenS49{F0: 4294967295, F1: 18446744073709551615}
                                                                                        _ = v633
                                                                                    }
                                                                                    fmt.Println(fi623)
                                                                                    var v634 float64 = 1.5
                                                                                    _ = v634
                                                                                }
                                                                                v629 = (v629 - v629)
                                                                                _ = v629
                                                                            } else {
                                                                                if ((v570 <= (v570 % ((-324) | 1))) || (-379 < -912)) {
                                                                                    switch fi624 {
                                                                                    case 92:
                                                                                        var v635 uint64 = 18446744073709551615
                                                                                        _ = v635
                                                                                    case 4294967295:
                                                                                        v575 = (v575 - (v575 ^ (v575 | 1)))
                                                                                    default:
                                                                                        fmt.Println(fi623, fi624, v575)
                                                                                    }
                                                                                    switch fi623 {
                                                                                    case -913:
                                                                                        delete(v2, (fi624 & -9223372036854775808))
                                                                                    case -1:
                                                                                        if true {
                                                                                            fmt.Println(v571, v575)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v3 = v3
                                                                            }
                                                                        case "foo bar":
                                                                            v636 := []int{}
                                                                            if v1 {
                                                                                var v637 int = (fi623 % (((^(fi623))) | 1))
                                                                                if (73 <= 42) {
                                                                                    v575 = (v575 ^ 4294967295)
                                                                                    switch v3 {
                                                                                    case "foo bar":
                                                                                        fmt.Println(v637, v637, v570)
                                                                                    }
                                                                                } else {
                                                                                    switch 4294967295 {
                                                                                    case 128:
                                                                                        if v1 {
                                                                                            if true {
                                                                                                var v638 int64 = -9223372036854775808
                                                                                                if v1 {
                                                                                                    if len(v636) >= 2 {
                                                                                                        v636 = v636[1:len(v636)]
                                                                                                    }
                                                                                                    var v639 bool = v1
                                                                                                    if len(v636) > 0 {
                                                                                                        v636[0] = 2147483647
                                                                                                    }
                                                                                                    _ = v639
                                                                                                } else {
                                                                                                    var v640 int8 = 127
                                                                                                    _ = len(v636)
                                                                                                    _ = v640
                                                                                                }
                                                                                                var v641 int64 = 768
                                                                                                var v642 float64 = 1.5
                                                                                                _ = v638
                                                                                                _ = v641
                                                                                                _ = v642
                                                                                            }
                                                                                        }
                                                                                        var v643 uint8 = uint8(v571)
                                                                                        _ = v643
                                                                                    case 9223372036854775807:
                                                                                        if true {
                                                                                            v1 = ((1e10 > 0.1) || v1)
                                                                                            if false {
                                                                                                if (!v1) {
                                                                                                    if ((1e10 > (3.14159 - float64(-2.25))) && (1 > uint64(v571))) {
                                                                                                        var v644 float32 = (0.0 * (float32(3.14159) - -2.25))
                                                                                                        _ = v644
                                                                                                    }
                                                                                                    var v645 float64 = -2.25
                                                                                                    var v646 uint32 = (v575 + (v575 % (((v575 ^ v575)) | 1)))
                                                                                                    _ = v645
                                                                                                    _ = v646
                                                                                                } else {
                                                                                                    v570 -= v570
                                                                                                }
                                                                                                v2[(fi623 - (v637 | 826))] = (!v1)
                                                                                            } else {
                                                                                                var v647 string = "\n\t"
                                                                                                _ = v647
                                                                                            }
                                                                                            v648 := GenS46{F0: 42, F1: 1, F2: (fi624 / ((454) | 1))}
                                                                                            _ = v648
                                                                                        } else {
                                                                                            var v649 uint = 1
                                                                                            _ = v649
                                                                                        }
                                                                                    case 273:
                                                                                        fmt.Println(v571)
                                                                                    default:
                                                                                        if ((v3 + v566) < (((((("foo bar" + v3) + ((v3 + v566) + v3)) + (v3 + (((((("a" + (v3 + "\n\t")) + ((v566 + (v3 + (v566 + v566))) + v566)) + v566) + v3) + ((((v3 + ("foo bar" + (v3 + ((v3 + (v566 + (("hello" + (((v566 + v3) + v3) + (("\n\t" + "foo bar") + v3))) + v566))) + "foo bar")))) + "foo bar") + "\n\t") + "")) + (v566 + "foo bar")))) + v3) + "") + ("" + v566))) {
                                                                                            var v650 uint64 = 901
                                                                                            v651 := GenS31{F0: v650}
                                                                                            fmt.Println(v3)
                                                                                            _ = v650
                                                                                            _ = v651
                                                                                        } else {
                                                                                            v652 := GenS69{F0: 42, F1: 1, F2: 62}
                                                                                            if (!(64 == 0)) {
                                                                                                if v1 {
                                                                                                    switch 32767 {
                                                                                                    case 9223372036854775807:
                                                                                                        v653 := GenS70{F0: -672, F1: (fi623 % (((fi623 / (((v637 & (^(fi624)))) | 1))) | 1)), F2: -17}
                                                                                                        _ = v653
                                                                                                    }
                                                                                                    v2[v637] = (!true)
                                                                                                    if ((425 <= (v571 / ((42) | 1))) || false) {
                                                                                                        var v654 float32 = 3.14159
                                                                                                        _ = v654
                                                                                                    } else {
                                                                                                        v655 := GenS23{F0: 1, F1: true, F2: 0}
                                                                                                        _ = v655
                                                                                                    }
                                                                                                }
                                                                                                if false {
                                                                                                    v1 = ("" <= v566)
                                                                                                    var v656 uint32 = v575
                                                                                                    if v1 {
                                                                                                        v1 = v1
                                                                                                    } else {
                                                                                                        delete(v2, fi623)
                                                                                                        fmt.Println(v656, fi623, v570)
                                                                                                    }
                                                                                                    _ = v656
                                                                                                }
                                                                                                if (!true) {
                                                                                                    if true {
                                                                                                        fmt.Println(v575, v575)
                                                                                                        v566 = ((v566 + v3) + v566)
                                                                                                        if len(v636) > 0 {
                                                                                                            v636[0] = (fi623 + 2147483647)
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v657 int8 = 127
                                                                                                        var v658 bool = ((uint32(v571) > v575) && v1)
                                                                                                        _ = v657
                                                                                                        _ = v658
                                                                                                    }
                                                                                                    switch fi623 {
                                                                                                    case -2147483648:
                                                                                                        var v659 int64 = 411
                                                                                                        _ = v659
                                                                                                    case -241:
                                                                                                        if (0.0 != 0.1) {
                                                                                                            v1 = (0.1 == (0.0 * (-2.25 - 0.0)))
                                                                                                            if len(v636) > 0 {
                                                                                                                v636[0] = fi623
                                                                                                            }
                                                                                                        }
                                                                                                    case -605:
                                                                                                        v660 := GenS64{F0: int8(v571)}
                                                                                                        _ = v660
                                                                                                    default:
                                                                                                        v661 := GenS5{F0: -27}
                                                                                                        _ = v661
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                v3 = (v566 + v566)
                                                                                                _ = len(v636)
                                                                                            }
                                                                                            v570 = v570
                                                                                            switch (v566 + (((((v3 + v3) + "a") + "\n\t") + v3) + (v3 + ("a" + v3)))) {
                                                                                            case "":
                                                                                                v1 = v1
                                                                                                if (v575 >= 346) {
                                                                                                    delete(v4, v3)
                                                                                                } else {
                                                                                                    if (false || v1) {
                                                                                                        fmt.Println(v637, v570)
                                                                                                    } else {
                                                                                                        v575 = (v575 + (v575 << ((311) & 63)))
                                                                                                    }
                                                                                                }
                                                                                                v1 = true
                                                                                                v566 = ((v566 + "foo bar") + "a")
                                                                                            }
                                                                                            _ = v652
                                                                                        }
                                                                                    }
                                                                                    if true {
                                                                                        fmt.Println(v575, v571, fi623)
                                                                                    } else {
                                                                                        if v1 {
                                                                                            fmt.Println(v3)
                                                                                            if v1 {
                                                                                                if v1 {
                                                                                                    if len(v636) >= 2 {
                                                                                                        v636 = v636[1:len(v636)]
                                                                                                    }
                                                                                                    if ((v570 < -1) || v1) {
                                                                                                        fmt.Println(v575)
                                                                                                    }
                                                                                                } else {
                                                                                                    if ((3.14159 > 0.1) || (1 >= 184)) {
                                                                                                        var v662 uint64 = 108
                                                                                                        _ = v662
                                                                                                    } else {
                                                                                                        var v663 int8 = -128
                                                                                                        if len(v636) >= 2 {
                                                                                                            v636 = v636[1:len(v636)]
                                                                                                        }
                                                                                                        _ = v663
                                                                                                    }
                                                                                                    v566 = v3
                                                                                                    switch (v637 + 32767) {
                                                                                                    case 414:
                                                                                                        v571 = (v571 - v571)
                                                                                                    case 388:
                                                                                                        v570 = (v570 << ((42) & 63))
                                                                                                        v570 -= 32767
                                                                                                        if len(v636) >= 2 {
                                                                                                            v636 = v636[1:len(v636)]
                                                                                                        }
                                                                                                        if v1 {
                                                                                                            if true {
                                                                                                                v566 = v566
                                                                                                                var v664 string = "\n\t"
                                                                                                                _ = v664
                                                                                                            }
                                                                                                            if len(v636) >= 2 {
                                                                                                                v636 = v636[1:len(v636)]
                                                                                                            }
                                                                                                        }
                                                                                                        var v665 uint32 = (v575 + v575)
                                                                                                        _ = v665
                                                                                                    }
                                                                                                    if v1 {
                                                                                                        v1 = v1
                                                                                                        v575 = 921
                                                                                                    }
                                                                                                }
                                                                                                v636 = append(v636, fi623)
                                                                                                v636 = append(v636, fi623)
                                                                                            }
                                                                                            var v666 int32 = 604
                                                                                            if true {
                                                                                                if false {
                                                                                                    v566 = v566
                                                                                                } else {
                                                                                                    var v667 int32 = (v666 - -2147483648)
                                                                                                    v668 := GenS71{F0: 42, F1: false}
                                                                                                    _ = v667
                                                                                                    _ = v668
                                                                                                }
                                                                                            }
                                                                                            _ = v666
                                                                                        } else {
                                                                                            v669 := GenS25{F0: 2}
                                                                                            fmt.Println(fi623, v566, v570)
                                                                                            _ = v669
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    switch v3 {
                                                                                    case "":
                                                                                        if ((v571 < (v571 << ((0) & 63))) && (v1 && (-2.25 >= (1.5 - -2.25)))) {
                                                                                            v670 := GenS6{F0: 134, F1: 4294967295, F2: 4294967295}
                                                                                            var v671 string = v3
                                                                                            _ = v670
                                                                                            _ = v671
                                                                                        }
                                                                                    case "hello":
                                                                                        v570 += -764
                                                                                    default:
                                                                                        switch (("\n\t" + v566) + v3) {
                                                                                        case "":
                                                                                            var v672 int64 = 1
                                                                                            _ = v672
                                                                                        case "a":
                                                                                            v570 = v570
                                                                                        default:
                                                                                            var v673 string = ("foo bar" + "\n\t")
                                                                                            _ = v673
                                                                                        }
                                                                                    }
                                                                                    if (683 <= 117) {
                                                                                        fmt.Println(v571, v571, v570)
                                                                                        delete(v4, v3)
                                                                                    }
                                                                                } else {
                                                                                    switch (v3 + v566) {
                                                                                    case "a":
                                                                                        var v674 int = v637
                                                                                        _ = v674
                                                                                    }
                                                                                    if v1 {
                                                                                        v675 := GenS72{F0: 42, F1: (false || v1)}
                                                                                        var v676 uint16 = v571
                                                                                        if (167 == 0) {
                                                                                            var v677 uint8 = 0
                                                                                            v636 = append(v636, fi624)
                                                                                            delete(v2, v637)
                                                                                            v678 := GenS17{F0: 0}
                                                                                            if (47 <= int8(v677)) {
                                                                                                v566 = ((("foo bar" + v566) + ("\n\t" + "\n\t")) + v3)
                                                                                                fmt.Println(fi623, v677, fi623)
                                                                                                switch (v637 | (fi623 << ((65535) & 63))) {
                                                                                                case -956:
                                                                                                    if v1 {
                                                                                                        v679 := GenS44{F0: (-(v570)), F1: 4294967295}
                                                                                                        _ = v679
                                                                                                    }
                                                                                                case -1:
                                                                                                    var v680 float32 = ((float32(((0.0 + 0.1) + 3.14159)) / 3.14159) / (0.0 + (float32(1e10) * 1.5)))
                                                                                                    _ = v680
                                                                                                case 42:
                                                                                                    v566 = (v566 + "\n\t")
                                                                                                    v570 = v570
                                                                                                }
                                                                                                var v681 int = (fi624 << ((886) & 63))
                                                                                                _ = v681
                                                                                            } else {
                                                                                                switch v566 {
                                                                                                case "hello":
                                                                                                    if v1 {
                                                                                                        v682 := v2[(fi624 % ((fi623) | 1))]
                                                                                                        _ = v682
                                                                                                    } else {
                                                                                                        var v683 int32 = 166
                                                                                                        _ = len(v4)
                                                                                                        _ = v683
                                                                                                    }
                                                                                                default:
                                                                                                    var v684 uint16 = v676
                                                                                                    _ = v684
                                                                                                }
                                                                                                if v1 {
                                                                                                    var v685 int64 = 1
                                                                                                    _ = v685
                                                                                                } else {
                                                                                                    if v1 {
                                                                                                        if len(v636) >= 2 {
                                                                                                            v636 = v636[1:len(v636)]
                                                                                                        }
                                                                                                        if v1 {
                                                                                                            var v686 int32 = -2147483648
                                                                                                            switch v3 {
                                                                                                            case "hello":
                                                                                                                if v1 {
                                                                                                                    var v687 int32 = v686
                                                                                                                    _ = v687
                                                                                                                } else {
                                                                                                                    v688 := GenS73{F0: 483}
                                                                                                                    v571 = 32767
                                                                                                                    var v689 string = v3
                                                                                                                    _ = v688
                                                                                                                    _ = v689
                                                                                                                }
                                                                                                            default:
                                                                                                                v690 := v4[v3]
                                                                                                                _ = v690
                                                                                                            }
                                                                                                            var v691 float32 = -2.25
                                                                                                            _ = v686
                                                                                                            _ = v691
                                                                                                        } else {
                                                                                                            v636 = append(v636, fi624)
                                                                                                            v571 = (v676 / ((v571) | 1))
                                                                                                            v692 := GenS47{F0: (v570 >> ((0) & 63))}
                                                                                                            v575 = 870
                                                                                                            _ = v692
                                                                                                        }
                                                                                                    } else {
                                                                                                        fmt.Println(v1, fi623, v637)
                                                                                                        if len(v636) > 0 {
                                                                                                            v636[0] = fi624
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                var v693 uint16 = (v676 & 0)
                                                                                                _ = v693
                                                                                            }
                                                                                            _ = v677
                                                                                            _ = v678
                                                                                        }
                                                                                        _ = v675
                                                                                        _ = v676
                                                                                    } else {
                                                                                        v694 := GenS74{F0: -2147483648, F1: "foo bar", F2: 2}
                                                                                        v1 = v1
                                                                                        _ = v694
                                                                                    }
                                                                                }
                                                                                _ = v637
                                                                            }
                                                                            if v1 {
                                                                                v695 := GenS61{F0: v570, F1: int32(v571), F2: 9223372036854775807}
                                                                                var v696 int32 = 2147483647
                                                                                _ = v695
                                                                                _ = v696
                                                                            }
                                                                            _ = v636
                                                                        }
                                                                        v697 := GenS19{F0: -1, F1: (v571 % (((v571 ^ genFn0((v571 + (v571 & 157)), v570, "\n\t"))) | 1)), F2: 625}
                                                                        _ = v697
                                                                    }
                                                                    _ = v628
                                                                }
                                                                var v698 uint32 = 2
                                                                if v1 {
                                                                    v1 = (!v1)
                                                                } else {
                                                                    fmt.Println(fi624, v698)
                                                                    fmt.Println(v3, v698, v571)
                                                                    switch (v566 + "a") {
                                                                    case "foo bar":
                                                                        v1 = v1
                                                                    case "hello":
                                                                        fmt.Println(v698, v3)
                                                                    case "a":
                                                                        v3 = v566
                                                                    }
                                                                }
                                                                _ = v698
                                                            }
                                                        }
                                                        if ((v1 && true) || (!true)) {
                                                            _ = len(v4)
                                                            switch ("a" + ((v566 + v3) + "a")) {
                                                            case "a":
                                                                if (v570 >= -116) {
                                                                    var v699 int64 = -9223372036854775808
                                                                    fmt.Println(v699, v699)
                                                                    _ = v699
                                                                } else {
                                                                    v700 := GenS75{F0: fi623, F1: (v570 * v570), F2: v571}
                                                                    v575 = 699
                                                                    _ = v700
                                                                }
                                                            }
                                                            var v701 bool = (v566 == "\n\t")
                                                            _ = v701
                                                        }
                                                        v702 := []int{fi624, (fi623 - 32767), (fi624 >> ((18446744073709551615) & 63))}
                                                        switch fi624 {
                                                        case 0:
                                                            fmt.Println(fi624)
                                                        case 464:
                                                            var v703 int16 = v570
                                                            _ = v703
                                                        default:
                                                            v704 := GenS29{F0: (v570 - v570), F1: 32767}
                                                            _ = v704
                                                        }
                                                        _ = v702
                                                    } else {
                                                        v705 := v2[fi623]
                                                        _ = v705
                                                    }
                                                    v706 := GenS76{F0: 255, F1: -1}
                                                    _ = v706
                                                }
                                            }
                                            var v707 float64 = 0.0
                                            _ = v707
                                        }
                                        _ = fi624
                                    }
                                    v1 = (((0.0 - 1e10) * 0.0) <= (((1e10 + -2.25) * (0.0 - -2.25)) / 1.5))
                                    var v708 float32 = (float32(0.0) * (0.0 / 1e10))
                                    _ = fi623
                                    _ = v708
                                }
                            }
                            if false {
                                v709 := GenS77{F0: ((v570 >= -1) && true), F1: 1}
                                _ = v709
                            }
                        }
                    }
                    if true {
                        v571 = genFn0((^(v571)), -32768, ("\n\t" + ((((("hello" + v566) + v566) + "a") + "foo bar") + ((v566 + v566) + v566))))
                        var v710 int = 1
                        _ = v710
                    }
                }
                switch 32767 {
                case -2147483648:
                    if v1 {
                        v2[32768] = (false && v1)
                        fmt.Println(v571)
                        fmt.Println(v3, v3)
                    } else {
                        for fi711 := 0; fi711 < 14; fi711++ {
                            v712 := []int{}
                            _ = fi711
                            _ = v712
                        }
                    }
                }
                v713 := []int{9223372036854775807, 4294967295}
                _ = v713
            }
            _ = v574
            _ = v575
        } else {
            fmt.Println(v570, v1)
        }
    } else {
        _ = len(v2)
    }
    var v714 float64 = 3.14159
    fmt.Println(v714, v566, v571)
    v570 = (v570 * (v570 >> ((1) & 63)))
    var v715 float64 = (float64(-2.25) / (v714 - (1e10 * v714)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v566:", v566)
    fmt.Println("v570:", v570)
    fmt.Println("v571:", v571)
    fmt.Println("v714:", v714)
    fmt.Println("v715:", v715)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v563
    _ = v566
    _ = v570
    _ = v571
    _ = v714
    _ = v715
}
