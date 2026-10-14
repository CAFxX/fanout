package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 string; F1 int; F2 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 int; F2 int64 }
type GenS8 struct { F0 int8; F1 int; F2 bool }
type GenS9 struct { F0 int; F1 int8 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint; F1 uint }
type GenS13 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS14 struct { F0 uint16; F1 uint64 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 int }
type GenS18 struct { F0 uint8; F1 bool }
type GenS19 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS20 struct { F0 uint8; F1 bool; F2 int32 }
type GenS21 struct { F0 uint; F1 string }
type GenS22 struct { F0 int8; F1 uint }
func genFn0() int8 {
    var v7 uint16 = 42
    if true {
        var v8 int16 = -879
        _ = v8
    }
    var v9 int32 = 906
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint8 = 0
        _ = fi10
        _ = v11
    }
    v7 = v7
    _ = v7
    _ = v9
    return 87
}

func main() {
    v1 := []int{-740, -1, 32768, 9223372036854775807}
    v2 := []int{}
    var v3 string = ""
    v4 := []uint32{0, 216, 1, 65535}
    var v5 int32 = -2147483648
    var v6 uint32 = 42
    fmt.Println(v5, v3)
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v5, v3)
        _ = fi12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        if false {
            fmt.Println(fi13)
            var v14 float64 = 0.0
            if ((v3 > "") || true) {
                if false {
                    for fi15 := 0; fi15 < 8; fi15++ {
                        switch v3 {
                        case "":
                            var v16 bool = (!(127 >= -100))
                            _ = v16
                        case "hello":
                            if len(v4) > 0 {
                                v4[0] = 2
                            }
                            v6 = (v6 + (v6 + (v6 % ((v6) | 1))))
                            if true {
                                v17 := GenS1{F0: v6, F1: 237}
                                _ = len(v1)
                                _ = v17
                            } else {
                                var v18 bool = (((v14 / (float64((0.0 + float32((v14 + (v14 * 0.1))))) / ((((((-2.25 * v14) - v14) + v14) - 0.0) / (1e10 * ((0.1 - v14) * v14))) / v14))) - ((v14 + 1.5) + float64(1e10))) != ((v14 - v14) * v14))
                                if (1000 <= 1) {
                                    fmt.Println(v14, fi13, v14)
                                }
                                _ = v18
                            }
                            if ((v6 % (((^(v6))) | 1)) != 884) {
                                v19 := GenS2{F0: (^(v6)), F1: int16(fi15)}
                                v14 = 1.5
                                if len(v4) > 0 {
                                    v4[0] = 0
                                }
                                _ = v19
                            }
                            switch ((("hello" + (v3 + v3)) + ((v3 + v3) + ((((("a" + v3) + v3) + (((((v3 + "hello") + "\n\t") + v3) + "") + (v3 + v3))) + v3) + v3))) + v3) {
                            case "\n\t":
                                switch (("foo bar" + ("hello" + v3)) + v3) {
                                case "a":
                                    if true {
                                        fmt.Println(fi13, v5, v14)
                                    } else {
                                        fmt.Println(v14)
                                    }
                                case "hello":
                                    if (!false) {
                                        var v20 int16 = -219
                                        _ = len(v2)
                                        var v21 string = (("\n\t" + (v3 + v3)) + "foo bar")
                                        _ = v20
                                        _ = v21
                                    }
                                }
                            case "":
                                v22 := GenS2{F0: 0, F1: 42}
                                v14 = v14
                                var v23 uint16 = 32767
                                _ = v22
                                _ = v23
                            case "hello":
                                v5 = (v5 << ((65535) & 63))
                            }
                            v5 -= 0
                            if (!(!(!(!false)))) {
                                v5 = (v5 << ((859) & 63))
                            }
                            var v24 bool = true
                            _ = v24
                        default:
                            v25 := GenS1{F0: v6, F1: (fi13 - (fi13 * 1))}
                            _ = v25
                        }
                        if ((0.0 >= 0.0) && false) {
                            var v26 uint32 = (v6 << ((42) & 63))
                            v3 = (v3 + v3)
                            _ = v26
                        } else {
                            v27 := GenS1{F0: v6, F1: fi13}
                            var v28 bool = false
                            _ = v27
                            _ = v28
                        }
                        _ = fi15
                    }
                    if len(v2) > 0 {
                        v2[0] = 132
                    }
                    var v29 uint64 = 42
                    _ = v29
                } else {
                    var v30 uint64 = 1
                    var v31 int64 = -1
                    v30 += v30
                    _ = v30
                    _ = v31
                }
                fmt.Println(v5)
            } else {
                if (false && (true && (0 < 274))) {
                    var v32 uint32 = v6
                    _ = v32
                } else {
                    var v33 float64 = -2.25
                    v34 := GenS2{F0: (^(v6)), F1: int16(fi13)}
                    v6 = 738
                    _ = v33
                    _ = v34
                }
            }
            _ = v14
        } else {
            var v35 int64 = -9223372036854775808
            fmt.Println(fi13, v5, v3)
            _ = v35
        }
        v36 := GenS2{F0: v6, F1: 42}
        for fi37 := 0; fi37 < 11; fi37++ {
            v3 = v3
            v5 += (-(v5))
            if true {
                if (1.5 != (1.5 + (-2.25 * 0.1))) {
                    if (65535 <= 1) {
                        var v38 float64 = 1.5
                        v38 += (v38 * 0.1)
                        _ = v38
                    } else {
                        v3 = (v3 + v3)
                    }
                } else {
                    switch "\n\t" {
                    case "a":
                        if (1e10 != (float64((0.0 - 0.1)) * (((1.5 + float64(0.0)) / (0.1 * -2.25)) - (-2.25 * (1.5 * 1e10))))) {
                            v3 = v3
                        } else {
                            fmt.Println(fi13, v5, fi13)
                        }
                    case "":
                        switch v3 {
                        case "foo bar":
                            v6 += v6
                        }
                    }
                    var v39 uint = 18446744073709551615
                    v40 := GenS3{F0: -99, F1: (^(v6))}
                    _ = v39
                    _ = v40
                }
                if false {
                    v6 = v6
                    v41 := GenS2{F0: (v6 ^ (^(v6))), F1: -1}
                    _ = v41
                } else {
                    if len(v4) > 0 {
                        v4[0] = (v6 >> ((255) & 63))
                    }
                    v6 -= 42
                    v3 = v3
                }
            } else {
                if false {
                    v3 = (v3 + (("" + "foo bar") + ("\n\t" + v3)))
                    v42 := GenS2{F0: (v6 + v6), F1: -32768}
                    if (!false) {
                        v5 += int32(fi13)
                        var v43 uint8 = 128
                        _ = v43
                    }
                    v3 = v3
                    var v44 int8 = int8(v6)
                    _ = v42
                    _ = v44
                } else {
                    if true {
                        if len(v2) > 0 {
                            v2[0] = fi13
                        }
                        v5 = (v5 ^ (v5 - 474))
                        switch (fi37 >> ((4294967295) & 63)) {
                        case 1:
                            if true {
                                fmt.Println(fi13, fi37, fi37)
                                v5 = (v5 << ((413) & 63))
                                _ = len(v2)
                                fmt.Println(v3, v5)
                            }
                        case 841:
                            fmt.Println(fi37, fi13, v5)
                        }
                    } else {
                        var v45 int64 = -9223372036854775808
                        _ = v45
                    }
                }
                if true {
                    v6 = (v6 ^ (v6 - 65535))
                } else {
                    if (!(false && (v3 >= "hello"))) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v6 = (^(v6))
                    }
                }
            }
            _ = fi37
        }
        _ = fi13
        _ = v36
    }
    for fi46 := 0; fi46 < 8; fi46++ {
        if ((fi46 / (((fi46 >> ((4294967295) & 63))) | 1)) != -561) {
            var v47 uint8 = 0
            fmt.Println(v47)
            _ = v47
        }
        var v48 float64 = 0.1
        switch fi46 {
        case -525:
            for fi49 := 0; fi49 < 9; fi49++ {
                var v50 string = v3
                var v51 uint8 = 0
                _ = fi49
                _ = v50
                _ = v51
            }
        case 9223372036854775807:
            if (false || (!true)) {
                v52 := GenS1{F0: 2147483647, F1: (fi46 << ((881) & 63))}
                _ = v52
            } else {
                v3 = ""
                for fi53 := 0; fi53 < 20; fi53++ {
                    fmt.Println(v5)
                    var v54 string = "\n\t"
                    v55 := GenS4{F0: 234, F1: 109}
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                switch v3 {
                case "":
                    var v56 int16 = -1
                    _ = v56
                case "\n\t":
                    if (true && (42 >= 0)) {
                        v57 := GenS3{F0: 0, F1: (v6 / (((v6 * 284)) | 1))}
                        var v58 int32 = v5
                        if true {
                            switch fi46 {
                            case -584:
                                var v59 uint16 = 614
                                _ = v59
                            case 127:
                                for fi60 := 0; fi60 < 8; fi60++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi60
                                }
                            case 128:
                                if false {
                                    v4 = append(v4, 2147483648)
                                } else {
                                    if true {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if false {
                                            var v61 uint16 = 65535
                                            fmt.Println(v3, v5)
                                            switch 32767 {
                                            case -2147483648:
                                                switch fi46 {
                                                case 127:
                                                    _ = len(v1)
                                                case 508:
                                                    v58 -= v58
                                                default:
                                                    v62 := GenS1{F0: (v6 + (v6 + (^(v6)))), F1: fi46}
                                                    _ = v62
                                                }
                                            case -9223372036854775808:
                                                v63 := GenS5{F0: v3, F1: fi46, F2: 128}
                                                _ = v63
                                            case 42:
                                                fmt.Println(v48)
                                            }
                                            _ = v61
                                        } else {
                                            var v64 uint64 = 4294967295
                                            for fi65 := 0; fi65 < 17; fi65++ {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v66 int = (fi65 - (fi65 | fi65))
                                                _ = fi65
                                                _ = v66
                                            }
                                            v58 = (v5 & v5)
                                            _ = v64
                                        }
                                    } else {
                                        v4 = append(v4, v6)
                                        var v67 int32 = v58
                                        _ = v67
                                    }
                                    fmt.Println(v6, v3)
                                    if true {
                                        v48 -= v48
                                        _ = len(v2)
                                        var v68 int32 = v5
                                        _ = v68
                                    }
                                }
                            default:
                                var v69 uint8 = 87
                                _ = v69
                            }
                        }
                        _ = v57
                        _ = v58
                    } else {
                        v3 = v3
                        v48 = ((-2.25 / ((v48 * 1.5) * v48)) / v48)
                        if (!false) {
                            var v70 float32 = 1e10
                            if ((65535 != 0) && true) {
                                if (155 >= 194) {
                                    var v71 uint = 0
                                    var v72 float64 = 1.5
                                    for fi73 := 0; fi73 < 9; fi73++ {
                                        v70 = v70
                                        _ = fi73
                                    }
                                    _ = v71
                                    _ = v72
                                }
                                _ = len(v2)
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            _ = v70
                        } else {
                            var v74 bool = (0.0 <= 0.1)
                            switch v3 {
                            case "hello":
                                for fi75 := 0; fi75 < 7; fi75++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    switch (fi75 >> ((972) & 63)) {
                                    case 9223372036854775807:
                                        v48 -= (v48 - -2.25)
                                    case 1:
                                        v6 -= (v6 / ((v6) | 1))
                                    case 0:
                                        v76 := GenS6{F0: -9223372036854775808}
                                        _ = v76
                                    }
                                    v77 := GenS7{F0: 2147483647, F1: fi46, F2: 42}
                                    if v74 {
                                        var v78 bool = false
                                        _ = v78
                                    } else {
                                        if (1 != -32768) {
                                            var v79 uint8 = 42
                                            _ = len(v4)
                                            _ = v79
                                        } else {
                                            switch (fi46 << ((42) & 63)) {
                                            case 0:
                                                if (!(v74 && (v74 || v74))) {
                                                    var v80 int64 = -918
                                                    _ = v80
                                                }
                                            case 128:
                                                var v81 string = (v3 + ("hello" + v3))
                                                _ = v81
                                            default:
                                                switch fi75 {
                                                case 9223372036854775807:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case 0:
                                                    v3 = (v3 + "")
                                                default:
                                                    if v74 {
                                                        v6 += 604
                                                    } else {
                                                        if v74 {
                                                            switch fi46 {
                                                            case 0:
                                                                if (true || (153 >= 255)) {
                                                                    v74 = (!((v5 << ((742) & 63)) != v5))
                                                                    fmt.Println(v6, v74)
                                                                }
                                                            case 1:
                                                                _ = len(v2)
                                                            }
                                                        } else {
                                                            var v82 uint64 = 42
                                                            _ = v82
                                                        }
                                                        if (v74 || true) {
                                                            var v83 string = v3
                                                            if ((v5 > 2147483647) && v74) {
                                                                switch 1 {
                                                                case 42:
                                                                    var v84 int64 = -1
                                                                    _ = v84
                                                                case -9223372036854775808:
                                                                    if v74 {
                                                                        var v85 float32 = 3.14159
                                                                        _ = v85
                                                                    } else {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = (v6 * (v6 + v6))
                                                                        }
                                                                        v83 = "\n\t"
                                                                        v86 := GenS8{F0: 127, F1: -121, F2: true}
                                                                        switch v83 {
                                                                        case "foo bar":
                                                                            v87 := GenS8{F0: 0, F1: (fi75 + 1), F2: (42 > int64(v6))}
                                                                            _ = v87
                                                                        }
                                                                        _ = v86
                                                                    }
                                                                case -427:
                                                                    v4 = append(v4, v6)
                                                                }
                                                                var v88 int8 = -1
                                                                v88 = (-(v88))
                                                                var v89 uint16 = 0
                                                                var v90 uint64 = 0
                                                                _ = v88
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                            _ = v83
                                                        } else {
                                                            v6 -= 256
                                                            switch "a" {
                                                            case "foo bar":
                                                                fmt.Println(fi75, v48, v48)
                                                            default:
                                                                var v91 uint8 = 42
                                                                _ = v91
                                                            }
                                                            v6 -= v6
                                                            if (432 != 0) {
                                                                v74 = v74
                                                                v74 = true
                                                            } else {
                                                                v92 := GenS9{F0: fi75, F1: -119}
                                                                _ = v92
                                                            }
                                                        }
                                                        _ = len(v1)
                                                    }
                                                }
                                            }
                                            if (55 >= 255) {
                                                var v93 bool = (36 != 128)
                                                v93 = (!true)
                                                var v94 uint8 = 42
                                                switch v3 {
                                                case "\n\t":
                                                    var v95 float32 = 1e10
                                                    _ = v95
                                                case "foo bar":
                                                    fmt.Println(v94)
                                                    v96 := GenS4{F0: 128, F1: 65535}
                                                    fmt.Println(v48, v93, v5)
                                                    _ = v96
                                                default:
                                                    var v97 uint16 = 0
                                                    _ = v97
                                                }
                                                fmt.Println(v94)
                                                switch fi75 {
                                                case -2147483648:
                                                    switch fi75 {
                                                    case 1:
                                                        if (!(((!v93) && false) || (1 < v6))) {
                                                            var v98 uint64 = 0
                                                            _ = v98
                                                        } else {
                                                            v99 := GenS10{F0: "hello"}
                                                            _ = v99
                                                        }
                                                    case 42:
                                                        v93 = v93
                                                    case 32767:
                                                        v1 = append(v1, 9223372036854775807)
                                                    default:
                                                        v100 := GenS1{F0: v6, F1: fi46}
                                                        _ = v100
                                                    }
                                                case 0:
                                                    v4 = append(v4, v6)
                                                case -143:
                                                    fmt.Println(fi75, fi46, v93)
                                                }
                                                if (1 >= (^(v94))) {
                                                    v101 := GenS3{F0: 127, F1: v6}
                                                    _ = v101
                                                } else {
                                                    var v102 int32 = v5
                                                    var v103 uint32 = (v6 / ((65535) | 1))
                                                    if (true && (false && (!(-99 >= -128)))) {
                                                        var v104 int = fi75
                                                        var v105 float32 = ((-2.25 / 3.14159) / 1e10)
                                                        if (float32(v48) == v105) {
                                                            switch "hello" {
                                                            case "":
                                                                var v106 bool = (!("\n\t" >= v3))
                                                                _ = v106
                                                            case "hello":
                                                                v74 = (!(-672 > int32(fi75)))
                                                            default:
                                                                if len(v4) > 0 {
                                                                    v4[0] = v103
                                                                }
                                                                v105 += (v105 - (float32((v48 * 0.1)) + v105))
                                                            }
                                                            if v74 {
                                                                if (((((v5 + 661) < v5) && v74) && (false && (false || v74))) || v93) {
                                                                    switch fi75 {
                                                                    case -26:
                                                                        v5 = (-(v5))
                                                                    default:
                                                                        v104 = -9223372036854775808
                                                                    }
                                                                    v6 -= 765
                                                                } else {
                                                                    switch fi75 {
                                                                    case -2147483648:
                                                                        var v107 uint64 = 2
                                                                        _ = v107
                                                                    case 395:
                                                                        var v108 int8 = 1
                                                                        _ = v108
                                                                    }
                                                                    switch v3 {
                                                                    case "foo bar":
                                                                        fmt.Println(fi46)
                                                                    case "":
                                                                        v94 = v94
                                                                    case "\n\t":
                                                                        if false {
                                                                            if v93 {
                                                                                _ = len(v1)
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = v104
                                                                                }
                                                                            } else {
                                                                                if (65535 >= 65535) {
                                                                                    v105 = ((v105 * v105) * -2.25)
                                                                                    v109 := GenS3{F0: 70, F1: (v6 / (((v6 & v6)) | 1))}
                                                                                    _ = v109
                                                                                } else {
                                                                                    v94 = (v94 & (v94 * (v94 ^ v94)))
                                                                                    fmt.Println(v94, v3)
                                                                                }
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (v104 % (((v104 / (((-(fi46))) | 1))) | 1))
                                                                                }
                                                                            }
                                                                            if v74 {
                                                                                if v93 {
                                                                                    if ((v94 & (v94 + 42)) == v94) {
                                                                                        var v110 uint64 = 638
                                                                                        _ = v110
                                                                                    } else {
                                                                                        switch (v104 << ((756) & 63)) {
                                                                                        case 9:
                                                                                            _ = len(v4)
                                                                                        case 2147483647:
                                                                                            fmt.Println(v94, fi46)
                                                                                        default:
                                                                                            if v93 {
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                            } else {
                                                                                                v2 = append(v2, fi75)
                                                                                            }
                                                                                            if v74 {
                                                                                                var v111 float64 = 0.1
                                                                                                _ = v111
                                                                                            } else {
                                                                                                v4 = append(v4, v6)
                                                                                            }
                                                                                            _ = len(v4)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if (-162 == 2147483647) {
                                                                                    var v112 int64 = 393
                                                                                    v104 += fi75
                                                                                    var v113 int = -9223372036854775808
                                                                                    _ = v112
                                                                                    _ = v113
                                                                                } else {
                                                                                    var v114 uint32 = uint32(v94)
                                                                                    var v115 int32 = -812
                                                                                    if true {
                                                                                        fmt.Println(v114, v48)
                                                                                        var v116 uint = 65535
                                                                                        _ = v116
                                                                                    } else {
                                                                                        if ((v115 * 350) >= (-(v115))) {
                                                                                            fmt.Println(v102, v102, v104)
                                                                                        }
                                                                                        _ = len(v4)
                                                                                        v114 -= v114
                                                                                        v48 -= v48
                                                                                    }
                                                                                    _ = v114
                                                                                    _ = v115
                                                                                }
                                                                            }
                                                                            fmt.Println(v3, v104)
                                                                            switch v3 {
                                                                            case "a":
                                                                                v117 := GenS10{F0: (v3 + v3)}
                                                                                _ = v117
                                                                            case "hello":
                                                                                v5 = v5
                                                                            case "":
                                                                                if false {
                                                                                    v104 += fi46
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = fi46
                                                                                    }
                                                                                    fmt.Println(v74, v104)
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v6 = v6
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            switch fi75 {
                                                            case 124:
                                                                var v118 uint64 = 33
                                                                _ = v118
                                                            case -953:
                                                                v104 = (fi75 ^ fi46)
                                                            case 127:
                                                                if true {
                                                                    var v119 uint32 = (v6 * (v6 % ((42) | 1)))
                                                                    _ = v119
                                                                }
                                                            }
                                                        }
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                    v120 := GenS11{F0: 42}
                                                    _ = v102
                                                    _ = v103
                                                    _ = v120
                                                }
                                                var v121 int64 = 0
                                                _ = v121
                                                _ = v93
                                                _ = v94
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = (fi46 << ((42) & 63))
                                                }
                                                v6 -= 2
                                                var v122 bool = v74
                                                _ = v122
                                            }
                                            v123 := GenS8{F0: genFn0(), F1: 0, F2: v74}
                                            switch ((((v3 + ("" + (v3 + "foo bar"))) + v3) + v3) + v3) {
                                            case "\n\t":
                                                if v74 {
                                                    var v124 int = (fi75 | fi75)
                                                    if v74 {
                                                        if (18446744073709551615 != uint64(v6)) {
                                                            var v125 uint8 = 152
                                                            var v126 int32 = v5
                                                            _ = v125
                                                            _ = v126
                                                        }
                                                        if ((1 == 4294967295) || true) {
                                                            if v74 {
                                                                switch (v3 + ("" + "a")) {
                                                                case "\n\t":
                                                                    fmt.Println(v3, v5, v124)
                                                                case "":
                                                                    v127 := GenS5{F0: "", F1: -678, F2: 42}
                                                                    if len(v1) > 0 {
                                                                        v1[0] = -1
                                                                    }
                                                                    _ = v127
                                                                case "foo bar":
                                                                    var v128 uint32 = 189
                                                                    _ = v128
                                                                }
                                                            }
                                                        } else {
                                                            var v129 uint8 = 2
                                                            _ = v129
                                                        }
                                                        v48 -= (1e10 + v48)
                                                    }
                                                    fmt.Println(v6)
                                                    _ = v124
                                                } else {
                                                    if true {
                                                        if (1 > 18446744073709551615) {
                                                            var v130 int32 = (-(v5))
                                                            if v74 {
                                                                v131 := GenS12{F0: 60, F1: 1}
                                                                fmt.Println(v74)
                                                                _ = v131
                                                            }
                                                            if len(v1) > 0 {
                                                                v1[0] = (fi75 / (((^(fi75))) | 1))
                                                            }
                                                            v132 := GenS2{F0: v6, F1: 278}
                                                            fmt.Println(fi75)
                                                            _ = v130
                                                            _ = v132
                                                        }
                                                        if v74 {
                                                            if (v74 && ((0 >= 802) || v74)) {
                                                                if (4294967295 != 42) {
                                                                    if (18446744073709551615 != 1) {
                                                                        v5 = v5
                                                                        fmt.Println(v5)
                                                                        v6 += 65535
                                                                    }
                                                                    v3 = "foo bar"
                                                                }
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                v48 -= v48
                                                                var v133 int16 = 0
                                                                switch v3 {
                                                                case "\n\t":
                                                                    if v74 {
                                                                        v6 -= v6
                                                                    }
                                                                case "":
                                                                    v48 = 1e10
                                                                }
                                                                _ = v133
                                                            } else {
                                                                v6 += (^(v6))
                                                                _ = len(v2)
                                                                v74 = true
                                                            }
                                                            if v74 {
                                                                var v134 uint32 = v6
                                                                v48 = v48
                                                                _ = v134
                                                            } else {
                                                                if false {
                                                                    _ = len(v2)
                                                                    fmt.Println(fi75)
                                                                } else {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = v6
                                                                    }
                                                                    var v135 int16 = -1
                                                                    switch "hello" {
                                                                    case "":
                                                                        fmt.Println(fi46)
                                                                    }
                                                                    _ = v135
                                                                }
                                                                if v74 {
                                                                    var v136 int32 = -2147483648
                                                                    _ = v136
                                                                }
                                                            }
                                                            v48 -= (0.1 + (float64(0.0) / v48))
                                                            v74 = v74
                                                            switch v3 {
                                                            case "foo bar":
                                                                var v137 uint32 = v6
                                                                _ = v137
                                                            case "hello":
                                                                v138 := GenS8{F0: -63, F1: -767, F2: v74}
                                                                _ = v138
                                                            }
                                                            switch v3 {
                                                            case "hello":
                                                                var v139 bool = (65535 >= 1)
                                                                _ = v139
                                                            case "\n\t":
                                                                var v140 int32 = int32(v6)
                                                                v141 := GenS9{F0: (fi75 - fi46), F1: -94}
                                                                _ = v140
                                                                _ = v141
                                                            case "":
                                                                v5 -= -2147483648
                                                            default:
                                                                var v142 float32 = 1.5
                                                                _ = v142
                                                            }
                                                        } else {
                                                            if v74 {
                                                                switch "hello" {
                                                                case "\n\t":
                                                                    var v143 uint16 = 610
                                                                    fmt.Println(fi46)
                                                                    _ = v143
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v144 := GenS13{F0: 9223372036854775807, F1: 122, F2: 2}
                                                        if ((0.1 - 0.0) < 1e10) {
                                                            if (v74 || (((^(fi46)) == fi46) && (v5 > v5))) {
                                                                if v74 {
                                                                    v145 := GenS14{F0: 1, F1: 510}
                                                                    v146 := GenS15{F0: 202}
                                                                    v6 -= (v6 % (((v6 + 973)) | 1))
                                                                    _ = v145
                                                                    _ = v146
                                                                }
                                                                v3 = v3
                                                                var v147 float32 = (0.1 - 1e10)
                                                                fmt.Println(fi46, v74, v74)
                                                                fmt.Println(v3)
                                                                _ = v147
                                                            } else {
                                                                if len(v4) > 0 {
                                                                    v4[0] = (^(v6))
                                                                }
                                                                if v74 {
                                                                    fmt.Println(v6, fi46, v74)
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v148 := GenS1{F0: v6, F1: 1}
                                                                    v149 := GenS16{F0: 255}
                                                                    _ = v148
                                                                    _ = v149
                                                                } else {
                                                                    v150 := GenS17{F0: fi46}
                                                                    _ = v150
                                                                }
                                                            }
                                                        } else {
                                                            var v151 int64 = 0
                                                            fmt.Println(fi46, v5)
                                                            _ = v151
                                                        }
                                                        _ = v144
                                                    }
                                                }
                                            case "":
                                                if v74 {
                                                    v5 += (v5 & (v5 & -1))
                                                    fmt.Println(fi46)
                                                } else {
                                                    v74 = true
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi75 ^ fi46)
                                                    }
                                                    var v152 uint32 = v6
                                                    _ = v152
                                                }
                                            case "a":
                                                v3 = ("hello" + v3)
                                            default:
                                                v153 := GenS9{F0: fi46, F1: -128}
                                                _ = v153
                                            }
                                            _ = v123
                                        }
                                        v48 = -2.25
                                        var v154 int16 = -32768
                                        _ = v154
                                    }
                                    _ = fi75
                                    _ = v77
                                }
                            case "":
                                v48 = v48
                            default:
                                var v155 uint64 = 0
                                _ = v155
                            }
                            _ = v74
                        }
                    }
                default:
                    fmt.Println(v6, v5, v6)
                }
            }
        case 348:
            if (false && (true && false)) {
                v156 := GenS18{F0: 117, F1: (!true)}
                if true {
                    switch "hello" {
                    case "hello":
                        fmt.Println(v5, v5, fi46)
                    case "\n\t":
                        var v157 float64 = v48
                        _ = v157
                    default:
                        switch fi46 {
                        case 852:
                            var v158 int16 = -32768
                            _ = v158
                        case -1:
                            v159 := GenS15{F0: 226}
                            _ = v159
                        case 1:
                            for fi160 := 0; fi160 < 10; fi160++ {
                                v161 := GenS15{F0: 280}
                                v162 := GenS7{F0: (v5 << ((18446744073709551615) & 63)), F1: (fi160 >> ((442) & 63)), F2: -9223372036854775808}
                                fmt.Println(v3, v3)
                                var v163 uint16 = 0
                                v163 = v163
                                _ = fi160
                                _ = v161
                                _ = v162
                                _ = v163
                            }
                        }
                    }
                }
                _ = v156
            }
        default:
            if false {
                for fi164 := 0; fi164 < 8; fi164++ {
                    v165 := GenS3{F0: 127, F1: (v6 | (v6 & 662))}
                    _ = fi164
                    _ = v165
                }
                if false {
                    fmt.Println(v6, v3, v3)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            } else {
                for fi166 := 0; fi166 < 4; fi166++ {
                    var v167 uint8 = 127
                    if false {
                        var v168 bool = true
                        v169 := GenS19{F0: 42, F1: 162, F2: (v167 >> ((65535) & 63))}
                        _ = v168
                        _ = v169
                    } else {
                        v4 = append(v4, 2)
                    }
                    v170 := GenS20{F0: 0, F1: (3.14159 == v48), F2: (v5 / ((-248) | 1))}
                    if len(v2) > 0 {
                        v2[0] = fi46
                    }
                    v3 = "\n\t"
                    var v171 float64 = v48
                    v4 = append(v4, v6)
                    _ = fi166
                    _ = v167
                    _ = v170
                    _ = v171
                }
            }
        }
        _ = fi46
        _ = v48
    }
    var v172 uint = 1
    if (!(false && (!false))) {
        var v173 uint16 = 255
        if (((v173 == 663) || true) && (true || false)) {
            if false {
                v5 = v5
                if (!((v5 | (v5 * v5)) > (v5 * v5))) {
                    for fi174 := 0; fi174 < 5; fi174++ {
                        v175 := GenS21{F0: 812, F1: (v3 + v3)}
                        v176 := GenS22{F0: genFn0(), F1: v172}
                        _ = fi174
                        _ = v175
                        _ = v176
                    }
                    v5 += -999
                    _ = len(v1)
                } else {
                    var v177 int16 = 1
                    fmt.Println(v177)
                    _ = v177
                }
            }
        }
        switch ("foo bar" + v3) {
        case "":
            v2 = append(v2, 128)
        case "foo bar":
            switch 128 {
            case 2147483647:
                v178 := GenS7{F0: v5, F1: 32767, F2: 9223372036854775807}
                _ = v178
            case -1:
                if len(v1) > 0 {
                    v1[0] = -879
                }
            case 9223372036854775807:
                var v179 string = (("hello" + v3) + ((v3 + "hello") + (v3 + v3)))
                _ = v179
            }
        case "a":
            v3 = "a"
        }
        _ = v173
    } else {
        fmt.Println(v3, v6)
    }
    fmt.Println(v3, v3, v3)
    v3 = ((v3 + v3) + v3)
    fmt.Println(v172, v6, v6)
    fmt.Println(v6, v5)
    for fi180 := 0; fi180 < 9; fi180++ {
        switch 128 {
        case 4294967295:
            var v181 int16 = 753
            _ = v181
        }
        for fi182 := 0; fi182 < 16; fi182++ {
            var v183 int64 = 42
            v3 = "a"
            v5 = (v5 - v5)
            _ = fi182
            _ = v183
        }
        _ = fi180
    }
    for fi184 := 0; fi184 < 11; fi184++ {
        var v185 int = (fi184 >> ((2) & 63))
        _ = fi184
        _ = v185
    }
    v186 := GenS22{F0: int8(v5), F1: v172}
    v3 = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v172:", v172)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v172
    _ = v186
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
