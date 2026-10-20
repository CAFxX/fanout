package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16; F1 int32 }
type GenS9 struct { F0 int64; F1 int32; F2 uint8 }
type GenS10 struct { F0 uint32; F1 uint32; F2 bool }
func genFn0() int16 {
    var v7 float64 = 0.0
    v8 := GenS1{F0: 42, F1: 2147483647, F2: 18446744073709551615}
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint16) int8 {
    var v9 uint32 = 65535
    var v10 float64 = 0.0
    switch -674 {
    case -512:
        var v11 string = "\n\t"
        _ = v11
    default:
        fmt.Println(p0, v10, v10)
    }
    v12 := GenS2{F0: 2147483648, F1: 1}
    _ = p0
    _ = v10
    _ = v12
    _ = v9
    return 100
}

func main() {
    var v1 float32 = 0.0
    v2 := []uint32{2147483648, 406}
    var v3 bool = true
    var v4 float64 = ((0.1 * (float64((v1 / 1.0)) / 0.1)) + (3.14159 + 1.5))
    v5 := make(map[string]int)
    var v6 string = ("a" + ("foo bar" + "hello"))
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            v1 -= v1
            _ = fi14
        }
        _ = fi13
    }
    v5[("\n\t" + v6)] = -2147483648
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int8 = 100
        _ = fi15
        _ = v16
    }
    v17 := GenS3{F0: -194, F1: v3}
    v4 = 0.0
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int32 = -1
        _ = fi18
        _ = v19
    }
    if ((558 == 42) || (false && ((v3 && v3) && v3))) {
        fmt.Println(v3, v4, v4)
        for fi20 := 0; fi20 < 11; fi20++ {
            if v3 {
                switch (v6 + v6) {
                case "":
                    if v3 {
                        v6 = v6
                        var v21 uint32 = 253
                        v21 = v21
                        _ = v21
                    } else {
                        var v22 int = -752
                        _ = v22
                    }
                case "\n\t":
                    v6 = v6
                case "a":
                    v4 = v4
                default:
                    fmt.Println(v6, v4)
                }
            } else {
                var v23 uint16 = 1
                for fi24 := 0; fi24 < 11; fi24++ {
                    fmt.Println(v6, v23, v3)
                    var v25 int16 = genFn0()
                    switch v6 {
                    case "hello":
                        v6 = (v6 + (v6 + "a"))
                    case "foo bar":
                        switch 4294967295 {
                        case 144:
                            v26 := GenS3{F0: 0, F1: true}
                            _ = v26
                        }
                    case "":
                        var v27 int32 = -932
                        _ = v27
                    }
                    _ = fi24
                    _ = v25
                }
                v6 = ("\n\t" + v6)
                _ = v23
            }
            v28 := GenS4{F0: -32768}
            _ = fi20
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := GenS4{F0: int16(fi29)}
        switch (^(fi29)) {
        case 893:
            var v31 uint8 = 64
            _ = v31
        default:
            v4 -= v4
        }
        for fi32 := 0; fi32 < 14; fi32++ {
            var v33 bool = (32767 == (fi29 << ((65535) & 63)))
            switch fi32 {
            case 536:
                switch ((v6 + v6) + "\n\t") {
                case "\n\t":
                    if true {
                        var v34 uint32 = 471
                        v35 := GenS5{F0: 0, F1: 42, F2: 2}
                        fmt.Println(v3)
                        _ = v34
                        _ = v35
                    } else {
                        if (-675 <= 0) {
                            var v36 float64 = v4
                            var v37 uint64 = 72
                            v2 = append(v2, 2147483648)
                            v38 := GenS3{F0: -88, F1: (!v3)}
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                    }
                    if ((730 > 1000) && (v3 && (!(v3 || (v1 > (-2.25 / -2.25)))))) {
                        if (("" + "\n\t") == v6) {
                            v33 = false
                        }
                        v39 := GenS5{F0: 42, F1: 822, F2: 491}
                        v33 = v33
                        _ = v39
                    } else {
                        v33 = v33
                        if (4294967295 != 2) {
                            if (!v3) {
                                var v40 int = fi32
                                if v3 {
                                    switch v6 {
                                    case "a":
                                        switch "foo bar" {
                                        case "a":
                                            if len(v2) > 0 {
                                                v2[0] = 263
                                            }
                                        case "hello":
                                            v4 = (v4 * (float64(1e10) / float64(v1)))
                                        case "":
                                            fmt.Println(v6)
                                        default:
                                            v4 = 3.14159
                                        }
                                    }
                                } else {
                                    if (v4 == 0.1) {
                                        var v41 uint16 = 42
                                        var v42 string = v6
                                        _ = v41
                                        _ = v42
                                    } else {
                                        if (0 <= -876) {
                                            if v3 {
                                                fmt.Println(v1, fi29, fi29)
                                            } else {
                                                v33 = ((1 == 858) || (1000 < 32767))
                                                var v43 int8 = 13
                                                var v44 uint64 = 1
                                                _ = v43
                                                _ = v44
                                            }
                                            switch v6 {
                                            case "a":
                                                if (v33 && v3) {
                                                    var v45 int32 = -1
                                                    var v46 uint = 0
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    v3 = v3
                                                    if v3 {
                                                        _ = len(v2)
                                                    } else {
                                                        fmt.Println(fi29, v6, v1)
                                                        if v3 {
                                                            var v47 string = v6
                                                            if (0 != 0) {
                                                                if len(v2) > 0 {
                                                                    v2[0] = 898
                                                                }
                                                                var v48 uint64 = 4294967295
                                                                var v49 bool = (!(v6 > v6))
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                            _ = v47
                                                        } else {
                                                            fmt.Println(v3, v1)
                                                            fmt.Println(fi32, v6)
                                                        }
                                                    }
                                                    var v50 uint8 = 0
                                                    _ = v50
                                                }
                                            case "hello":
                                                fmt.Println(v6, fi32)
                                            default:
                                                v33 = (32 >= 128)
                                            }
                                            if (!false) {
                                                if true {
                                                    var v51 uint16 = 2
                                                    v51 = v51
                                                    _ = v51
                                                } else {
                                                    v3 = v3
                                                    if (42 <= 446) {
                                                        fmt.Println(v40)
                                                    }
                                                    v1 = v1
                                                }
                                                v52 := v5[((v6 + "hello") + ("a" + v6))]
                                                _ = v52
                                            } else {
                                                fmt.Println(v3, v1)
                                                var v53 int = fi29
                                                var v54 int = (v53 - v40)
                                                v55 := GenS5{F0: 1, F1: 32767, F2: 18446744073709551615}
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                            var v56 float32 = (0.0 * (v1 / (0.0 - v1)))
                                            v56 = v56
                                            _ = v56
                                        } else {
                                            v5[(("\n\t" + "foo bar") + v6)] = -1
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    var v57 float64 = (1e10 / v4)
                                    var v58 int32 = 42
                                    _ = v57
                                    _ = v58
                                }
                                _ = v40
                            }
                            v59 := GenS5{F0: 172, F1: 2, F2: 4294967295}
                            _ = v59
                        } else {
                            var v60 uint64 = 18446744073709551615
                            _ = v60
                        }
                        fmt.Println(fi32)
                    }
                    var v61 string = v6
                    _ = v61
                }
            case -1:
                switch (fi29 << ((4294967295) & 63)) {
                case 2147483647:
                    fmt.Println(v1)
                    switch fi32 {
                    case 32767:
                        v3 = v33
                    case 328:
                        v62 := v5[v6]
                        _ = v62
                    case -160:
                        var v63 float32 = v1
                        var v64 int = (fi32 | fi29)
                        switch v64 {
                        case -36:
                            var v65 uint64 = 2
                            _ = v65
                        default:
                            fmt.Println(v4, v4, v4)
                        }
                        switch (v6 + ((v6 + ("\n\t" + "a")) + ((v6 + ((v6 + (((("\n\t" + "hello") + "\n\t") + (v6 + v6)) + v6)) + ("a" + v6))) + v6))) {
                        case "foo bar":
                            v4 -= v4
                        default:
                            var v66 int16 = 0
                            _ = v66
                        }
                        _ = v63
                        _ = v64
                    default:
                        if false {
                            var v67 int32 = 42
                            var v68 bool = (!v33)
                            _ = v67
                            _ = v68
                        } else {
                            var v69 bool = (!((v1 * ((((float32(0.1) + float32(v4)) / 1.0) / (0.0 * (-2.25 + v1))) * (v1 * 0.1))) < 3.14159))
                            v1 += v1
                            var v70 int8 = -128
                            delete(v5, v6)
                            _ = v69
                            _ = v70
                        }
                    }
                    if (3.14159 >= v1) {
                        if v33 {
                            var v71 int16 = -88
                            if (42 == 32767) {
                                if v3 {
                                    switch (v6 + "foo bar") {
                                    case "\n\t":
                                        var v72 int = fi32
                                        _ = v72
                                    case "":
                                        fmt.Println(v71)
                                    default:
                                        var v73 int32 = 0
                                        _ = v73
                                    }
                                }
                                v74 := GenS6{F0: 255}
                                _ = len(v5)
                                _ = v74
                            }
                            fmt.Println(fi32)
                            switch "" {
                            case "\n\t":
                                if (127 != 2) {
                                    v6 = "foo bar"
                                    if (-1 == genFn1(716)) {
                                        if len(v2) > 0 {
                                            v2[0] = 624
                                        }
                                        v75 := GenS2{F0: 1, F1: 268}
                                        _ = v75
                                    }
                                    v76 := GenS6{F0: 924}
                                    _ = v76
                                }
                            case "hello":
                                switch (fi32 * (fi32 >> ((459) & 63))) {
                                case 1:
                                    v6 = "a"
                                    var v77 int32 = 0
                                    v33 = (838 == 608)
                                    _ = v77
                                case 42:
                                    switch -2147483648 {
                                    case 128:
                                        _ = len(v2)
                                        var v78 uint32 = 2
                                        if v3 {
                                            v79 := GenS4{F0: v71}
                                            if v33 {
                                                var v80 int16 = (v71 * v71)
                                                var v81 uint32 = (v78 >> ((1) & 63))
                                                switch (v6 + v6) {
                                                case "\n\t":
                                                    _ = len(v5)
                                                case "":
                                                    v2 = append(v2, v81)
                                                }
                                                _ = v80
                                                _ = v81
                                            } else {
                                                v1 = v1
                                            }
                                            v71 -= (v71 >> ((771) & 63))
                                            _ = v79
                                        } else {
                                            if (v3 && ((v3 || v3) || v33)) {
                                                v1 -= (float32((v4 / v4)) + 0.0)
                                            } else {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v78 = (v78 << ((979) & 63))
                                                switch (("" + "\n\t") + "a") {
                                                case "":
                                                    if (127 < 42) {
                                                        v82 := GenS6{F0: 4294967295}
                                                        var v83 float32 = v1
                                                        switch fi32 {
                                                        case 32768:
                                                            if ((65535 < 140) || v33) {
                                                                v84 := GenS1{F0: 741, F1: -1, F2: 18446744073709551615}
                                                                if v3 {
                                                                    fmt.Println(v1, v1, fi29)
                                                                }
                                                                v33 = v3
                                                                _ = v84
                                                            }
                                                        case 972:
                                                            if v33 {
                                                                v1 -= (v83 + float32(v4))
                                                                var v85 int64 = 1
                                                                v83 = v83
                                                                _ = v85
                                                            } else {
                                                                if (!(1 <= 1)) {
                                                                    var v86 uint64 = 2
                                                                    _ = v86
                                                                } else {
                                                                    switch fi29 {
                                                                    case -994:
                                                                        var v87 uint16 = uint16(v78)
                                                                        _ = v87
                                                                    case -539:
                                                                        if (127 >= 100) {
                                                                            v33 = (!v33)
                                                                        }
                                                                    default:
                                                                        v71 = (v71 + v71)
                                                                    }
                                                                }
                                                                v2 = append(v2, (v78 * 2147483647))
                                                                var v88 int8 = 120
                                                                _ = v88
                                                            }
                                                        case -9223372036854775808:
                                                            var v89 int32 = -2147483648
                                                            _ = v89
                                                        default:
                                                            v5["\n\t"] = (fi32 & 42)
                                                        }
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        if true {
                                                            v90 := GenS5{F0: 1, F1: 1, F2: 0}
                                                            v3 = v33
                                                            var v91 uint8 = 128
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                        var v92 string = v6
                                                        v4 -= 0.1
                                                        var v93 bool = (-128 > -128)
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                case "hello":
                                                    _ = len(v5)
                                                case "foo bar":
                                                    v5[v6] = fi32
                                                default:
                                                    if v3 {
                                                        if (!(0 != 4294967295)) {
                                                            v71 -= (v71 | 32767)
                                                        } else {
                                                            v33 = true
                                                            var v94 int32 = 350
                                                            var v95 uint16 = 42
                                                            if false {
                                                                v96 := v5["foo bar"]
                                                                var v97 int = int(v94)
                                                                _ = v96
                                                                _ = v97
                                                            }
                                                            _ = v94
                                                            _ = v95
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        _ = v78
                                    case -2147483648:
                                        switch "a" {
                                        case "\n\t":
                                            v33 = false
                                        default:
                                            v71 = -719
                                        }
                                    case 4294967295:
                                        var v98 int32 = 42
                                        _ = v98
                                    }
                                }
                            case "foo bar":
                                switch -1 {
                                case 4294967295:
                                    v4 = v4
                                default:
                                    v3 = v33
                                }
                            default:
                                var v99 uint8 = 145
                                _ = v99
                            }
                            if v33 {
                                v100 := GenS7{F0: 42}
                                _ = v100
                            } else {
                                _ = len(v2)
                                if v3 {
                                    if (160 == 1) {
                                        v101 := GenS5{F0: 42, F1: 704, F2: 0}
                                        fmt.Println(v71)
                                        switch "hello" {
                                        case "a":
                                            v102 := GenS8{F0: v71, F1: 42}
                                            _ = v102
                                        }
                                        _ = v101
                                    }
                                    if (v3 || v3) {
                                        var v103 string = "\n\t"
                                        var v104 bool = (42 == 42)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v103
                                        _ = v104
                                    }
                                } else {
                                    v4 = (0.0 + v4)
                                    v33 = v33
                                }
                            }
                            _ = v71
                        }
                        if v3 {
                            if len(v2) > 0 {
                                v2[0] = 1
                            }
                            v105 := GenS9{F0: 147, F1: 1, F2: 2}
                            _ = v105
                        }
                    }
                case 923:
                    if (0 < 319) {
                        fmt.Println(fi29)
                        v106 := GenS9{F0: -263, F1: -2147483648, F2: 1}
                        fmt.Println(fi32)
                        _ = v106
                    } else {
                        v3 = true
                        if v33 {
                            var v107 int32 = 594
                            var v108 int32 = v107
                            _ = len(v5)
                            v109 := GenS10{F0: 2, F1: 49, F2: (v6 > v6)}
                            _ = v107
                            _ = v108
                            _ = v109
                        } else {
                            var v110 uint32 = 790
                            _ = v110
                        }
                        v111 := GenS8{F0: 0, F1: 565}
                        _ = v111
                    }
                default:
                    v112 := v5["foo bar"]
                    _ = v112
                }
            }
            _ = fi32
            _ = v33
        }
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
