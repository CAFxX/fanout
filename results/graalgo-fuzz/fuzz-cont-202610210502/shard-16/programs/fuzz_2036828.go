package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint }
type GenS4 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int8; F1 int; F2 uint16 }
type GenS8 struct { F0 uint8; F1 uint16 }
type GenS9 struct { F0 int16; F1 int32; F2 uint16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint64; F1 int64; F2 int }
type GenS12 struct { F0 string; F1 uint16; F2 uint64 }
type GenS13 struct { F0 bool; F1 int64; F2 uint32 }
type GenS14 struct { F0 int64; F1 int32; F2 int16 }
type GenS15 struct { F0 uint; F1 int32; F2 int16 }
type GenS16 struct { F0 bool; F1 int }
type GenS17 struct { F0 uint64 }

func main() {
    var v1 float64 = 0.1
    var v2 int32 = 2147483647
    v3 := make(map[uint32]string)
    var v4 int8 = 101
    var v5 bool = true
    var v6 uint = 981
    delete(v3, 65)
    v5 = (v1 <= v1)
    var v7 bool = false
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: (-(fi8)), F1: ""}
        _ = fi8
        _ = v9
    }
    if (float32((1e10 * v1)) < float32(v1)) {
        for fi10 := 0; fi10 < 10; fi10++ {
            for fi11 := 0; fi11 < 12; fi11++ {
                if (false || v7) {
                    switch ("foo bar" + (((("" + "\n\t") + ("a" + "hello")) + ("foo bar" + "\n\t")) + ("foo bar" + "foo bar"))) {
                    case "hello":
                        v12 := GenS1{F0: fi11, F1: ("hello" + "a")}
                        v13 := GenS1{F0: fi10, F1: ""}
                        v1 -= v1
                        _ = v12
                        _ = v13
                    case "a":
                        var v14 uint = (^(v6))
                        _ = v14
                    default:
                        v1 = (float64(((0.1 * float32(v1)) / (3.14159 + 1e10))) + 3.14159)
                    }
                    var v15 int64 = -70
                    _ = v15
                } else {
                    var v16 uint32 = 42
                    var v17 float32 = float32(float64(1e10))
                    fmt.Println(fi11)
                    _ = v16
                    _ = v17
                }
                fmt.Println(v7, fi10)
                _ = fi11
            }
            v3[200] = ""
            _ = fi10
        }
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 float32 = -2.25
            var v20 int32 = v2
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    var v21 uint = 867
    var v22 float32 = -2.25
    var v23 float32 = (((v22 + float32(1.5)) / v22) - v22)
    for fi24 := 0; fi24 < 10; fi24++ {
        switch "hello" {
        case "hello":
            var v25 int16 = 294
            _ = v25
        case "foo bar":
            v26 := []int{fi24, fi24, fi24, fi24}
            _ = v26
        default:
            v6 -= (v6 + v21)
        }
        var v27 string = ("foo bar" + "\n\t")
        v27 = v27
        _ = fi24
        _ = v27
    }
    switch -710 {
    case 608:
        for fi28 := 0; fi28 < 20; fi28++ {
            fmt.Println(v4)
            if v7 {
                fmt.Println(v7)
                v29 := GenS2{F0: (v5 && ((!v5) && v7)), F1: 1, F2: 4294967295}
                v30 := GenS1{F0: 42, F1: ("" + "foo bar")}
                _ = v29
                _ = v30
            } else {
                if false {
                    var v31 int8 = v4
                    for fi32 := 0; fi32 < 18; fi32++ {
                        var v33 int32 = (v2 >> ((v6) & 63))
                        if ((v31 + (^(v31))) >= (v31 + 0)) {
                            var v34 float32 = (float32(v1) * (v23 - v23))
                            fmt.Println(fi32, v4, v7)
                            _ = v34
                        }
                        switch "\n\t" {
                        case "hello":
                            var v35 int8 = v4
                            if true {
                                v31 = 0
                            } else {
                                v36 := []int{-9223372036854775808, 0}
                                if len(v36) >= 2 {
                                    v36 = v36[1:len(v36)]
                                }
                                if (!v7) {
                                    var v37 uint32 = 65535
                                    if (v33 < (v33 << ((v21) & 63))) {
                                        fmt.Println(v23)
                                        v2 = -808
                                        v21 -= v21
                                        if ((v7 && false) && ("hello" > "a")) {
                                            var v38 int8 = (v4 >> (((v21 & (v6 * (v21 % ((760) | 1))))) & 63))
                                            v33 = -7
                                            _ = v38
                                        }
                                        var v39 float32 = float32(((((float64((-2.25 / 1.0)) / v1) - (float64((v23 / float32((3.14159 - 0.0)))) + 0.0)) / v1) / float64(v23)))
                                        var v40 int16 = 1
                                        _ = v39
                                        _ = v40
                                    }
                                    _ = v37
                                } else {
                                    switch ("" + ("hello" + "foo bar")) {
                                    case "\n\t":
                                        if v5 {
                                            if v5 {
                                                if (true || true) {
                                                    _ = len(v36)
                                                    v41 := GenS3{F0: v4, F1: 4294967295, F2: 65535}
                                                    var v42 string = "\n\t"
                                                    _ = v41
                                                    _ = v42
                                                }
                                                fmt.Println(v7)
                                                v35 = (v4 ^ 42)
                                            } else {
                                                var v43 uint8 = 42
                                                if v5 {
                                                    var v44 uint32 = 1
                                                    switch (fi28 + fi32) {
                                                    case -2147483648:
                                                        if false {
                                                            fmt.Println(fi28)
                                                        } else {
                                                            switch ("" + "hello") {
                                                            case "":
                                                                switch (fi28 + 4294967295) {
                                                                case 32767:
                                                                    if ((v44 + 4294967295) == v44) {
                                                                        v22 = -2.25
                                                                    }
                                                                case 42:
                                                                    if (v21 == 42) {
                                                                        if false {
                                                                            var v45 int64 = -9223372036854775808
                                                                            var v46 int32 = v2
                                                                            var v47 int = fi28
                                                                            v48 := GenS2{F0: (fi28 != 42), F1: 42, F2: 4294967295}
                                                                            _ = v45
                                                                            _ = v46
                                                                            _ = v47
                                                                            _ = v48
                                                                        } else {
                                                                            v2 = v33
                                                                            var v49 uint16 = 2
                                                                            switch "hello" {
                                                                            case "a":
                                                                                if true {
                                                                                    _ = len(v3)
                                                                                    switch 1 {
                                                                                    case 4294967295:
                                                                                        var v50 float32 = 0.0
                                                                                        _ = v50
                                                                                    case -9223372036854775808:
                                                                                        if (18446744073709551615 != 18446744073709551615) {
                                                                                            var v51 uint8 = v43
                                                                                            v52 := GenS4{F0: 578, F1: (v44 - v44), F2: 116}
                                                                                            if v7 {
                                                                                                if (("a" + ("foo bar" + "\n\t")) <= "") {
                                                                                                    if len(v36) >= 2 {
                                                                                                        v36 = v36[1:len(v36)]
                                                                                                    }
                                                                                                    v5 = v5
                                                                                                }
                                                                                                v53 := GenS4{F0: 1, F1: v44, F2: (v35 - v31)}
                                                                                                var v54 uint = (^(v21))
                                                                                                _ = v53
                                                                                                _ = v54
                                                                                            }
                                                                                            _ = v51
                                                                                            _ = v52
                                                                                        } else {
                                                                                            var v55 string = "foo bar"
                                                                                            if ((^(v43)) <= (v43 / ((v43) | 1))) {
                                                                                                fmt.Println(v33, v55)
                                                                                                v21 = v6
                                                                                            } else {
                                                                                                fmt.Println(v6)
                                                                                                v33 = (^(v2))
                                                                                                var v56 uint64 = 42
                                                                                                if ((-9223372036854775808 > 42) || v7) {
                                                                                                    v7 = v5
                                                                                                    var v57 float64 = float64((v22 / (v22 - float32(v1))))
                                                                                                    _ = v57
                                                                                                }
                                                                                                switch "foo bar" {
                                                                                                case "a":
                                                                                                    v58 := v3[v44]
                                                                                                    _ = v58
                                                                                                }
                                                                                                if (v56 == (v56 % ((v56) | 1))) {
                                                                                                    v6 += v6
                                                                                                    if (!v5) {
                                                                                                        v55 = v55
                                                                                                    }
                                                                                                    if ((v23 - 1.5) != (3.14159 + v22)) {
                                                                                                        v35 = (^(v35))
                                                                                                        v31 -= v35
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            if v5 {
                                                                                                                if ((v44 * 2147483648) != 2147483647) {
                                                                                                                    switch (fi28 * fi32) {
                                                                                                                    case 0:
                                                                                                                        var v59 bool = (v7 && false)
                                                                                                                        _ = v59
                                                                                                                    default:
                                                                                                                        v60 := GenS5{F0: (858 > (v56 ^ 0))}
                                                                                                                        _ = v60
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v61 := GenS6{F0: 32767, F1: int8(v56)}
                                                                                                                    _ = v61
                                                                                                                }
                                                                                                                v35 += v31
                                                                                                            } else {
                                                                                                                v62 := GenS5{F0: v7}
                                                                                                                if v5 {
                                                                                                                    v5 = (v5 && true)
                                                                                                                } else {
                                                                                                                    var v63 float64 = float64(0.1)
                                                                                                                    _ = v63
                                                                                                                }
                                                                                                                if false {
                                                                                                                    var v64 float64 = (v1 / v1)
                                                                                                                    var v65 int = 4294967295
                                                                                                                    var v66 uint8 = (^(v43))
                                                                                                                    var v67 uint64 = v56
                                                                                                                    fmt.Println(v5)
                                                                                                                    _ = v64
                                                                                                                    _ = v65
                                                                                                                    _ = v66
                                                                                                                    _ = v67
                                                                                                                } else {
                                                                                                                    v68 := GenS7{F0: (-(v35)), F1: 128, F2: v49}
                                                                                                                    _ = v68
                                                                                                                }
                                                                                                                _ = v62
                                                                                                            }
                                                                                                            switch "" {
                                                                                                            case "\n\t":
                                                                                                                if ((v31 > 1) || false) {
                                                                                                                    if len(v36) >= 2 {
                                                                                                                        v36 = v36[1:len(v36)]
                                                                                                                    }
                                                                                                                }
                                                                                                            case "hello":
                                                                                                                if (v7 && v5) {
                                                                                                                    v56 -= v56
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        v36 = append(v36, fi32)
                                                                                                                        v35 = v31
                                                                                                                        var v69 float64 = v1
                                                                                                                        _ = v69
                                                                                                                    } else {
                                                                                                                        v21 += (v21 % ((100) | 1))
                                                                                                                    }
                                                                                                                    var v70 string = ((("\n\t" + v55) + v55) + "")
                                                                                                                    if (false || true) {
                                                                                                                        v71 := GenS8{F0: 57, F1: (v49 % ((v49) | 1))}
                                                                                                                        _ = v71
                                                                                                                    } else {
                                                                                                                        var v72 float64 = 0.0
                                                                                                                        v73 := GenS2{F0: v7, F1: (v56 / (((^(v56))) | 1)), F2: (v56 + 2)}
                                                                                                                        v74 := GenS9{F0: 0, F1: (v33 * (v2 >> ((255) & 63))), F2: (v49 | 988)}
                                                                                                                        _ = v72
                                                                                                                        _ = v73
                                                                                                                        _ = v74
                                                                                                                    }
                                                                                                                    _ = v70
                                                                                                                }
                                                                                                            default:
                                                                                                                if len(v36) > 0 {
                                                                                                                    v36[0] = (fi32 / ((fi32) | 1))
                                                                                                                }
                                                                                                            }
                                                                                                            v36 = append(v36, (-(fi32)))
                                                                                                        }
                                                                                                        if (0 == v35) {
                                                                                                            var v75 int = -1
                                                                                                            _ = v75
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                _ = v56
                                                                                            }
                                                                                            var v76 uint32 = uint32(v49)
                                                                                            _ = v55
                                                                                            _ = v76
                                                                                        }
                                                                                    }
                                                                                    switch "a" {
                                                                                    case "\n\t":
                                                                                        if v7 {
                                                                                            v77 := GenS10{F0: -521}
                                                                                            _ = v77
                                                                                        }
                                                                                    case "a":
                                                                                        if (v35 > v31) {
                                                                                            var v78 bool = (!true)
                                                                                            _ = v78
                                                                                        } else {
                                                                                            v36 = append(v36, (fi28 % ((42) | 1)))
                                                                                            v35 -= v31
                                                                                        }
                                                                                    case "":
                                                                                        v79 := GenS11{F0: 1, F1: -9223372036854775808, F2: -948}
                                                                                        _ = v79
                                                                                    default:
                                                                                        v80 := GenS6{F0: 1, F1: (v31 ^ (v35 / ((v4) | 1)))}
                                                                                        _ = v80
                                                                                    }
                                                                                } else {
                                                                                    var v81 string = ""
                                                                                    switch -9223372036854775808 {
                                                                                    case 743:
                                                                                        if v7 {
                                                                                            v36 = append(v36, fi32)
                                                                                            v36 = append(v36, fi28)
                                                                                            switch (("a" + ("hello" + v81)) + v81) {
                                                                                            case "foo bar":
                                                                                                v31 = v4
                                                                                            case "\n\t":
                                                                                                v3[(v44 * 594)] = v81
                                                                                            default:
                                                                                                switch ("\n\t" + ("a" + v81)) {
                                                                                                case "\n\t":
                                                                                                    var v82 int16 = 0
                                                                                                    _ = v82
                                                                                                case "":
                                                                                                    var v83 int = 32767
                                                                                                    if ((v81 + v81) > v81) {
                                                                                                        var v84 float64 = 3.14159
                                                                                                        _ = v84
                                                                                                    }
                                                                                                    _ = v83
                                                                                                case "a":
                                                                                                    fmt.Println(v23, v44, v7)
                                                                                                }
                                                                                            }
                                                                                            if len(v36) >= 2 {
                                                                                                v36 = v36[1:len(v36)]
                                                                                            }
                                                                                            if (v43 >= 73) {
                                                                                                fmt.Println(v6, v23)
                                                                                            } else {
                                                                                                var v85 uint8 = 0
                                                                                                v86 := GenS10{F0: 42}
                                                                                                var v87 bool = false
                                                                                                _ = v85
                                                                                                _ = v86
                                                                                                _ = v87
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        var v88 int64 = -9223372036854775808
                                                                                        _ = v88
                                                                                    }
                                                                                    fmt.Println(v21)
                                                                                    _ = v81
                                                                                }
                                                                            case "foo bar":
                                                                                v89 := GenS4{F0: 4294967295, F1: (v44 << (((v6 | uint(v49))) & 63)), F2: v4}
                                                                                _ = v89
                                                                            case "\n\t":
                                                                                v23 = v22
                                                                            default:
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    var v90 int16 = -32768
                                                                                    _ = v90
                                                                                }
                                                                            }
                                                                            var v91 uint = uint(v49)
                                                                            _ = v49
                                                                            _ = v91
                                                                        }
                                                                        v92 := GenS12{F0: "", F1: 65535, F2: 42}
                                                                        _ = v92
                                                                    }
                                                                case 128:
                                                                    v93 := v3[v44]
                                                                    _ = v93
                                                                default:
                                                                    if v7 {
                                                                        v94 := GenS12{F0: "", F1: 2, F2: 42}
                                                                        var v95 bool = false
                                                                        _ = v94
                                                                        _ = v95
                                                                    }
                                                                }
                                                            case "a":
                                                                if (v7 && (!((!(!v7)) && (1 < 0)))) {
                                                                    _ = len(v36)
                                                                    switch (fi28 << ((0) & 63)) {
                                                                    case 127:
                                                                        var v96 int32 = v33
                                                                        _ = v96
                                                                    case 277:
                                                                        v21 = 4294967295
                                                                    }
                                                                } else {
                                                                    v97 := GenS8{F0: (v43 / ((129) | 1)), F1: 489}
                                                                    _ = v97
                                                                }
                                                                v98 := GenS13{F0: v5, F1: -9223372036854775808, F2: 250}
                                                                _ = v98
                                                            default:
                                                                if ((!v7) || v5) {
                                                                    if false {
                                                                        if v7 {
                                                                            v99 := GenS14{F0: 1, F1: (^(v2)), F2: 42}
                                                                            fmt.Println(v33, v1)
                                                                            _ = v99
                                                                        }
                                                                        var v100 bool = (2 >= 296)
                                                                        _ = v100
                                                                    }
                                                                    if v5 {
                                                                        v21 -= v6
                                                                        var v101 string = ""
                                                                        _ = v101
                                                                    } else {
                                                                        v44 += (v44 & v44)
                                                                    }
                                                                }
                                                                v102 := GenS15{F0: (v21 * (^(v21))), F1: 2147483647, F2: -248}
                                                                _ = v102
                                                            }
                                                            v33 += v33
                                                        }
                                                    case -1:
                                                        v5 = ((v1 / float64((-2.25 - v22))) <= -2.25)
                                                    case 433:
                                                        if len(v36) >= 2 {
                                                            v36 = v36[1:len(v36)]
                                                        }
                                                    }
                                                    _ = v44
                                                } else {
                                                    var v103 int8 = -106
                                                    if ((v6 % ((v6) | 1)) >= (^(v6))) {
                                                        var v104 float32 = v23
                                                        _ = v104
                                                    } else {
                                                        var v105 uint8 = (v43 & 122)
                                                        _ = v105
                                                    }
                                                    _ = v103
                                                }
                                                _ = v43
                                            }
                                            v6 += (^(v21))
                                            fmt.Println(v21, v6, v4)
                                            _ = len(v36)
                                        }
                                    case "a":
                                        if len(v36) >= 2 {
                                            v36 = v36[1:len(v36)]
                                        }
                                    case "foo bar":
                                        v106 := v3[42]
                                        _ = v106
                                    }
                                    v107 := GenS6{F0: 261, F1: v4}
                                    if v5 {
                                        var v108 uint = v6
                                        _ = v108
                                    } else {
                                        v109 := GenS11{F0: 1, F1: 42, F2: 4294967295}
                                        switch fi28 {
                                        case 183:
                                            v5 = false
                                        case -9223372036854775808:
                                            if true {
                                                switch ("foo bar" + ("a" + (("a" + "foo bar") + ("hello" + "\n\t")))) {
                                                case "a":
                                                    v36 = append(v36, 81)
                                                case "":
                                                    _ = len(v3)
                                                }
                                            }
                                        }
                                        switch "foo bar" {
                                        case "foo bar":
                                            _ = len(v3)
                                        case "":
                                            var v110 float64 = 0.0
                                            _ = v110
                                        case "hello":
                                            if v5 {
                                                v111 := GenS13{F0: (1 < 140), F1: -1, F2: 257}
                                                fmt.Println(v33, v5, fi32)
                                                v21 = (v21 ^ v6)
                                                switch fi28 {
                                                case 9223372036854775807:
                                                    v33 = 42
                                                default:
                                                    v112 := v3[2147483648]
                                                    _ = v112
                                                }
                                                v113 := GenS12{F0: "a", F1: 32767, F2: 0}
                                                v4 -= v4
                                                _ = v111
                                                _ = v113
                                            }
                                            _ = len(v36)
                                        default:
                                            delete(v3, 53)
                                        }
                                        switch (fi32 * (fi32 + 275)) {
                                        case 128:
                                            if (true || v5) {
                                                if len(v36) >= 2 {
                                                    v36 = v36[1:len(v36)]
                                                }
                                            } else {
                                                if ((v21 << ((v6) & 63)) > v6) {
                                                    if true {
                                                        var v114 float64 = (float64(v22) * v1)
                                                        _ = v114
                                                    }
                                                }
                                                if true {
                                                    var v115 string = (("hello" + ((("hello" + "") + "foo bar") + ("foo bar" + "foo bar"))) + "hello")
                                                    _ = v115
                                                }
                                                delete(v3, 283)
                                            }
                                        case 127:
                                            switch "a" {
                                            case "\n\t":
                                                var v116 string = ""
                                                v116 = v116
                                                _ = v116
                                            case "a":
                                                if ((704 == 2147483648) || (!v7)) {
                                                    fmt.Println(v7, v35)
                                                    fmt.Println(v1, v2)
                                                } else {
                                                    v117 := GenS10{F0: -32768}
                                                    _ = len(v3)
                                                    var v118 uint16 = 1
                                                    var v119 int = (fi32 << (((v21 % ((586) | 1))) & 63))
                                                    v5 = false
                                                    _ = v117
                                                    _ = v118
                                                    _ = v119
                                                }
                                            }
                                        case 42:
                                            var v120 uint32 = 650
                                            _ = v120
                                        default:
                                            if (42 != 0) {
                                                switch 128 {
                                                case 127:
                                                    if (v7 || (!true)) {
                                                        v1 = 0.1
                                                        if v7 {
                                                            if v7 {
                                                                fmt.Println(v33, v6, v4)
                                                                v121 := GenS1{F0: 416, F1: ("" + ("hello" + ""))}
                                                                fmt.Println(v2, fi28, v4)
                                                                _ = v121
                                                            }
                                                            v33 = -684
                                                            if v7 {
                                                                var v122 uint32 = 2147483648
                                                                fmt.Println(v22, v122, v5)
                                                                _ = v122
                                                            }
                                                        }
                                                    } else {
                                                        v4 = v35
                                                        v6 = (^(v6))
                                                        if false {
                                                            switch fi28 {
                                                            case 15:
                                                                v33 = v2
                                                            case 305:
                                                                v123 := GenS12{F0: "a", F1: 150, F2: 971}
                                                                _ = v123
                                                            default:
                                                                if v7 {
                                                                    var v124 int64 = -889
                                                                    if len(v36) >= 2 {
                                                                        v36 = v36[1:len(v36)]
                                                                    }
                                                                    _ = v124
                                                                } else {
                                                                    var v125 bool = ((v2 << ((4294967295) & 63)) == (-(v33)))
                                                                    switch fi28 {
                                                                    case -953:
                                                                        if len(v36) >= 2 {
                                                                            v36 = v36[1:len(v36)]
                                                                        }
                                                                    case 0:
                                                                        v126 := GenS2{F0: v125, F1: 89, F2: 42}
                                                                        _ = v126
                                                                    case 317:
                                                                        var v127 int16 = 1
                                                                        _ = v127
                                                                    }
                                                                    fmt.Println(v4, v35)
                                                                    fmt.Println(v4)
                                                                    fmt.Println(v2, v23)
                                                                    _ = v125
                                                                }
                                                            }
                                                            var v128 string = "hello"
                                                            _ = v128
                                                        }
                                                    }
                                                case 9223372036854775807:
                                                    v5 = (v35 == v35)
                                                case 128:
                                                    fmt.Println(v22, v35)
                                                }
                                                switch "hello" {
                                                case "":
                                                    v35 += -98
                                                case "hello":
                                                    v129 := GenS1{F0: 32768, F1: "foo bar"}
                                                    _ = v129
                                                }
                                            }
                                        }
                                        if v5 {
                                            var v130 bool = v5
                                            fmt.Println(v33, v2, v23)
                                            v1 = 1.5
                                            _ = v130
                                        }
                                        if v7 {
                                            v131 := v3[42]
                                            _ = v131
                                        }
                                        _ = v109
                                    }
                                    _ = v107
                                }
                                v4 += v35
                                _ = v36
                            }
                            var v132 float32 = 1e10
                            _ = v132
                            _ = v35
                        case "a":
                            if (36 != -9223372036854775808) {
                                v133 := GenS10{F0: -32768}
                                _ = v133
                            }
                        default:
                            if v5 {
                                var v134 int64 = 423
                                if false {
                                    var v135 uint32 = 0
                                    var v136 int = fi28
                                    fmt.Println(v135, v1)
                                    var v137 uint64 = 1
                                    var v138 int32 = int32(v136)
                                    v135 += v135
                                    _ = v135
                                    _ = v136
                                    _ = v137
                                    _ = v138
                                } else {
                                    if v5 {
                                        v5 = true
                                        switch "foo bar" {
                                        case "hello":
                                            v31 -= 127
                                        default:
                                            v139 := v3[2147483648]
                                            _ = v139
                                        }
                                    }
                                    v140 := GenS16{F0: (!v7), F1: (fi32 * 128)}
                                    v141 := GenS16{F0: v7, F1: fi28}
                                    if ((fi28 ^ fi28) >= (fi32 << (((v21 & v6)) & 63))) {
                                        v142 := []int{(fi28 & 1), fi28, (fi32 >> (((v6 & v21)) & 63))}
                                        var v143 int = fi28
                                        if (fi32 == (-(fi32))) {
                                            v144 := GenS17{F0: 641}
                                            var v145 float64 = v1
                                            v146 := GenS13{F0: v7, F1: v134, F2: 1}
                                            _ = v144
                                            _ = v145
                                            _ = v146
                                        }
                                        _ = v142
                                        _ = v143
                                    }
                                    _ = v140
                                    _ = v141
                                }
                                _ = v134
                            }
                        }
                        var v147 bool = v7
                        switch fi32 {
                        case -2147483648:
                            fmt.Println(v4, v147)
                            if ((v1 / v1) != (v1 / 1.0)) {
                                var v148 int32 = (^(v33))
                                _ = len(v3)
                                if ((-1 >= 1) && true) {
                                    v21 = 18446744073709551615
                                    var v149 int64 = -1
                                    _ = v149
                                }
                                var v150 uint16 = 882
                                v151 := GenS16{F0: ((fi32 % (((fi32 >> (((v6 & (v6 - v6))) & 63))) | 1)) == 4294967295), F1: fi32}
                                var v152 string = "foo bar"
                                _ = v148
                                _ = v150
                                _ = v151
                                _ = v152
                            }
                        case 32767:
                            switch fi32 {
                            case 1:
                                v147 = v5
                            case 0:
                                var v153 int32 = -348
                                _ = v153
                            default:
                                fmt.Println(v147, v1, v2)
                                var v154 uint64 = 1
                                fmt.Println(v31, fi28, v147)
                                _ = v154
                            }
                            v3[35] = ""
                            fmt.Println(v7, v31)
                        case 215:
                            var v155 int64 = 9223372036854775807
                            _ = v155
                        }
                        _ = fi32
                        _ = v147
                        _ = v33
                    }
                    _ = v31
                }
                var v156 int8 = (^(v4))
                switch "\n\t" {
                case "foo bar":
                    var v157 uint64 = 0
                    _ = v157
                case "\n\t":
                    v21 += 289
                case "hello":
                    var v158 uint64 = 4294967295
                    _ = v158
                }
                _ = v156
            }
            _ = len(v3)
            _ = fi28
        }
    case 2147483647:
        switch -1 {
        case -2147483648:
            v21 += (v6 * v6)
        case 0:
            var v159 uint = (v21 - v21)
            _ = v159
        default:
            switch 32768 {
            case -2147483648:
                if v7 {
                    v6 = (v21 / (((v21 / ((v21) | 1))) | 1))
                    var v160 float64 = v1
                    _ = v160
                }
            case -469:
                var v161 uint64 = 42
                _ = v161
            case -9223372036854775808:
                fmt.Println(v6, v2)
            }
        }
    default:
        fmt.Println(v2, v23)
    }
    if ((-(v4)) <= 1) {
        v162 := GenS9{F0: int16(v4), F1: 0, F2: 1}
        _ = v162
    } else {
        delete(v3, 4294967295)
    }
    for fi163 := 0; fi163 < 13; fi163++ {
        if (102 <= 87) {
            var v164 bool = false
            var v165 uint64 = uint64(v21)
            fmt.Println(v21)
            _ = v164
            _ = v165
        }
        fmt.Println(v6, v23, v4)
        _ = fi163
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
