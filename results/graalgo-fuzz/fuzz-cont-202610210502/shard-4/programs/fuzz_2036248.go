package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int32; F2 bool }
type GenS5 struct { F0 uint64; F1 uint; F2 bool }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int; F1 int; F2 int }
type GenS11 struct { F0 int; F1 bool }
type GenS12 struct { F0 uint; F1 int32 }
type GenS13 struct { F0 int; F1 uint32 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int32; F1 bool; F2 uint64 }
type GenS16 struct { F0 int64 }

func main() {
    var v1 int64 = 0
    v2 := []uint32{65535, 13, 441, 401}
    var v3 uint = 4294967295
    v4 := make(map[uint32]int)
    var v5 bool = (501 <= -32768)
    var v6 int16 = 1
    if (0 == -1) {
        v2 = append(v2, 2)
        v3 = 1
    } else {
        var v7 uint8 = 2
        _ = v7
    }
    v5 = (-1 == 127)
    if v5 {
        switch 766 {
        case 463:
            v8 := GenS1{F0: "foo bar", F1: "foo bar"}
            _ = v8
        case -1:
            v6 = -1
        }
        v9 := GenS2{F0: int16(v3)}
        if (v6 > v6) {
            v10 := GenS1{F0: "\n\t", F1: ("foo bar" + ("a" + ("\n\t" + "a")))}
            _ = v10
        } else {
            v1 -= int64(v3)
            v6 = v6
            var v11 string = ((("a" + "a") + "") + (("a" + "") + "hello"))
            _ = v11
        }
        _ = v9
    }
    if ((!(0 == 0)) && (int(v6) < 32767)) {
        switch "foo bar" {
        case "a":
            v1 = (v1 % ((-1) | 1))
        default:
            if ((4294967295 != 0) && false) {
                var v12 uint64 = 41
                _ = v12
            } else {
                v5 = v5
            }
        }
        var v13 uint64 = 1
        v2 = append(v2, uint32(v13))
        _ = v13
    } else {
        var v14 float64 = (1.5 + 3.14159)
        for fi15 := 0; fi15 < 20; fi15++ {
            if (!false) {
                fmt.Println(fi15, v5, fi15)
                for fi16 := 0; fi16 < 4; fi16++ {
                    var v17 bool = ((v3 / ((18446744073709551615) | 1)) >= (v3 & 297))
                    if (!(true && (!(!((0.0 - 1.5) >= (float32(v14) * (1e10 * 0.1))))))) {
                        var v18 uint16 = 503
                        if ((!((v18 + (v18 * 2)) > 0)) || v5) {
                            if (!((v3 + (v3 % ((v3) | 1))) >= v3)) {
                                if (-413 <= (fi15 * 128)) {
                                    v18 = (v18 / ((32767) | 1))
                                    switch -331 {
                                    case 2147483647:
                                        if (-101 <= 127) {
                                            if v5 {
                                                var v19 string = ""
                                                if len(v2) > 0 {
                                                    v2[0] = 0
                                                }
                                                _ = v19
                                            }
                                            var v20 uint8 = 255
                                            _ = v20
                                        } else {
                                            var v21 uint8 = 128
                                            v2 = append(v2, 239)
                                            v22 := GenS1{F0: "a", F1: "\n\t"}
                                            v6 = v6
                                            v1 -= v1
                                            v23 := GenS3{F0: 127}
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        }
                                    case 127:
                                        _ = len(v2)
                                    default:
                                        switch fi16 {
                                        case 32767:
                                            v18 = (^(v18))
                                        case 487:
                                            var v24 uint32 = 1
                                            if v17 {
                                                fmt.Println(fi15, v14, v3)
                                            } else {
                                                var v25 int = fi16
                                                if v17 {
                                                    v5 = v5
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v26 float32 = ((0.1 / 0.1) * ((1e10 / 1.0) - -2.25))
                                                        _ = v26
                                                    case "":
                                                        if len(v2) > 0 {
                                                            v2[0] = v24
                                                        }
                                                        var v27 int = (-(v25))
                                                        v1 += (v1 >> ((18446744073709551615) & 63))
                                                        _ = v27
                                                    default:
                                                        switch fi15 {
                                                        case 555:
                                                            switch (fi15 ^ (^(fi16))) {
                                                            case -488:
                                                                if false {
                                                                    v17 = (1e10 != (v14 * (0.0 / (v14 - -2.25))))
                                                                    fmt.Println(fi16, v3, v6)
                                                                    var v28 int = 128
                                                                    _ = v28
                                                                } else {
                                                                    if (false && v5) {
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            v17 = false
                                                                            switch "a" {
                                                                            case "a":
                                                                                var v29 uint64 = uint64(v18)
                                                                                _ = v29
                                                                            case "":
                                                                                if (!(!false)) {
                                                                                    _ = len(v4)
                                                                                    v17 = ((fi16 - (fi16 / (((fi16 & 32767)) | 1))) != 1)
                                                                                    fmt.Println(v25, v24, v1)
                                                                                } else {
                                                                                    var v30 uint8 = 1
                                                                                    _ = v30
                                                                                }
                                                                            default:
                                                                                _ = len(v2)
                                                                            }
                                                                        case "hello":
                                                                            var v31 string = "foo bar"
                                                                            _ = v31
                                                                        case "":
                                                                            fmt.Println(v1, v14)
                                                                        default:
                                                                            if (-649 != v1) {
                                                                                v32 := GenS3{F0: (fi16 << ((18446744073709551615) & 63))}
                                                                                _ = v32
                                                                            }
                                                                        }
                                                                        var v33 bool = (42 != 421)
                                                                        v24 = (v24 | v24)
                                                                        _ = v33
                                                                    } else {
                                                                        var v34 uint32 = 2147483647
                                                                        var v35 int = -2147483648
                                                                        _ = v34
                                                                        _ = v35
                                                                    }
                                                                    switch 1 {
                                                                    case -345:
                                                                        v5 = (v25 <= -2147483648)
                                                                    }
                                                                    var v36 uint64 = 1
                                                                    _ = v36
                                                                }
                                                            case 32767:
                                                                v37 := GenS4{F0: fi15, F1: 1, F2: v17}
                                                                _ = v37
                                                            case 42:
                                                                if len(v2) > 0 {
                                                                    v2[0] = 2
                                                                }
                                                            }
                                                        case 9223372036854775807:
                                                            switch v25 {
                                                            case 32767:
                                                                if (v18 < v18) {
                                                                    fmt.Println(v3, v14)
                                                                    v1 = v1
                                                                }
                                                            case 42:
                                                                if (v5 && ((v3 / ((v3) | 1)) != v3)) {
                                                                    if (!(v25 < (fi16 % (((fi16 % ((fi16) | 1))) | 1)))) {
                                                                        if (18446744073709551615 < 42) {
                                                                            v38 := GenS3{F0: (^(v25))}
                                                                            fmt.Println(v3)
                                                                            var v39 uint16 = v18
                                                                            _ = v38
                                                                            _ = v39
                                                                        }
                                                                        var v40 uint32 = 4294967295
                                                                        var v41 int = (fi15 >> (((v3 / ((v3) | 1))) & 63))
                                                                        v14 = (((1.5 * 1.5) - ((((v14 + ((v14 * v14) + (((0.1 + ((((v14 / (v14 * (v14 / v14))) - v14) - ((v14 - (v14 * ((v14 + v14) * ((float64((-2.25 + 0.1)) * -2.25) + float64((3.14159 + 1.5)))))) * (float64(3.14159) + ((v14 - (v14 * v14)) * v14)))) * (v14 * (float64((float32(1.5) / 1.0)) + v14)))) + ((float64((-2.25 * -2.25)) * v14) + float64(-2.25))) + (1e10 - (0.1 + ((1.5 / ((0.1 / v14) * ((1e10 / v14) + (((v14 / 1.0) - ((1e10 * ((v14 - float64((float32(0.0) / 1e10))) / (1.5 - v14))) + v14)) * 1.5)))) + 0.1)))))) - (v14 + 0.1)) * -2.25) - ((1.5 + 0.0) + 1e10))) - 3.14159)
                                                                        var v42 uint8 = 255
                                                                        _ = v40
                                                                        _ = v41
                                                                        _ = v42
                                                                    } else {
                                                                        if (!(0 >= (-(v25)))) {
                                                                            v43 := GenS2{F0: int16(v18)}
                                                                            v44 := GenS5{F0: 42, F1: v3, F2: (517 == 18446744073709551615)}
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                if ((0.0 / ((1e10 + float32(0.0)) + 1.5)) < 3.14159) {
                                                                                    if v5 {
                                                                                        if v17 {
                                                                                            var v45 uint16 = v18
                                                                                            if (!false) {
                                                                                                v18 += (v18 >> ((uint(v45)) & 63))
                                                                                                v46 := GenS3{F0: fi16}
                                                                                                v6 += 42
                                                                                                v47 := GenS4{F0: 729, F1: -1, F2: (v17 && v17)}
                                                                                                _ = v46
                                                                                                _ = v47
                                                                                            } else {
                                                                                                fmt.Println(v3, v5, v18)
                                                                                            }
                                                                                            v48 := GenS1{F0: "a", F1: ((("a" + "") + "foo bar") + "foo bar")}
                                                                                            _ = v45
                                                                                            _ = v48
                                                                                        } else {
                                                                                            if v5 {
                                                                                                v3 = (v3 / ((v3) | 1))
                                                                                                v17 = v17
                                                                                                switch ("hello" + "a") {
                                                                                                case "\n\t":
                                                                                                    v49 := GenS1{F0: ("" + ""), F1: "hello"}
                                                                                                    _ = v49
                                                                                                case "foo bar":
                                                                                                    var v50 uint32 = v24
                                                                                                    _ = v50
                                                                                                case "":
                                                                                                    fmt.Println(v25, v14)
                                                                                                }
                                                                                            }
                                                                                            v51 := GenS1{F0: ("" + "foo bar"), F1: ((("a" + ("foo bar" + "\n\t")) + "\n\t") + ("a" + ((("foo bar" + "\n\t") + "") + "\n\t")))}
                                                                                            _ = v51
                                                                                        }
                                                                                        if false {
                                                                                            v3 = 255
                                                                                            var v52 bool = v17
                                                                                            v24 += 2147483647
                                                                                            if ((3.14159 - 0.0) <= float32(0.0)) {
                                                                                                fmt.Println(v14, v1)
                                                                                                var v53 uint16 = (v18 / ((v18) | 1))
                                                                                                _ = v53
                                                                                            }
                                                                                            v54 := GenS4{F0: fi15, F1: -965, F2: ((513 > 0) && (-484 == -2147483648))}
                                                                                            _ = v52
                                                                                            _ = v54
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v1, v14)
                                                                                    v55 := GenS5{F0: 4294967295, F1: 916, F2: (((^(v18)) <= v18) || (fi15 >= (fi15 & v25)))}
                                                                                    _ = v55
                                                                                }
                                                                            case "a":
                                                                                fmt.Println(v6)
                                                                            default:
                                                                                fmt.Println(v6)
                                                                            }
                                                                            _ = v43
                                                                            _ = v44
                                                                        }
                                                                        v18 -= (v18 + (v18 ^ uint16(v3)))
                                                                    }
                                                                }
                                                                fmt.Println(v18)
                                                                var v56 uint64 = 42
                                                                _ = v56
                                                            case 128:
                                                                if (!false) {
                                                                    var v57 uint = (v3 & v3)
                                                                    _ = v57
                                                                } else {
                                                                    fmt.Println(v3, v17)
                                                                }
                                                            }
                                                        case -1:
                                                            var v58 float64 = v14
                                                            _ = v58
                                                        default:
                                                            v25 = fi16
                                                        }
                                                    }
                                                } else {
                                                    var v59 uint16 = v18
                                                    v60 := GenS6{F0: -1, F1: v17}
                                                    if (0.1 == 1e10) {
                                                        _ = len(v4)
                                                    } else {
                                                        v14 = v14
                                                    }
                                                    _ = v59
                                                    _ = v60
                                                }
                                                var v61 int64 = v1
                                                if (1e10 == float64(1.5)) {
                                                    v62 := GenS7{F0: true, F1: -9223372036854775808}
                                                    if ((1e10 / 1.0) != (float32(v14) * 1e10)) {
                                                        var v63 float64 = (((float64((((0.1 - ((1e10 / (1e10 / 1.0)) + 0.1)) * ((float32(v14) / 1.5) - 3.14159)) - 1e10)) - 0.0) * (v14 * (v14 - ((1.5 - float64(1e10)) - v14)))) - v14)
                                                        _ = v63
                                                    } else {
                                                        var v64 uint8 = 2
                                                        _ = v64
                                                    }
                                                    _ = v62
                                                }
                                                _ = v25
                                                _ = v61
                                            }
                                            _ = v24
                                        case 127:
                                            fmt.Println(v5, v18)
                                            v65 := GenS8{F0: 1}
                                            if (0 <= 0) {
                                                v66 := GenS7{F0: v17, F1: -1}
                                                _ = v66
                                            } else {
                                                var v67 int64 = 610
                                                _ = v67
                                            }
                                            _ = v65
                                        }
                                    }
                                }
                                fmt.Println(v6)
                                if false {
                                    v68 := GenS9{F0: 4294967295}
                                    v4[2] = (-(fi16))
                                    var v69 uint = (v3 & 966)
                                    if true {
                                        if (!(v17 || ((v14 == 3.14159) || false))) {
                                            var v70 int16 = v6
                                            if v17 {
                                                var v71 int32 = 889
                                                switch fi16 {
                                                case 32767:
                                                    v17 = (!(18446744073709551615 == 2))
                                                    v18 = v18
                                                case -186:
                                                    var v72 int64 = (v1 >> ((v69) & 63))
                                                    _ = v72
                                                case -530:
                                                    if len(v2) > 0 {
                                                        v2[0] = 2
                                                    }
                                                    if ((v71 ^ (v71 % (((^(v71))) | 1))) <= (-(v71))) {
                                                        v18 = 988
                                                        var v73 uint = (^(v3))
                                                        _ = v73
                                                    }
                                                    if v17 {
                                                        var v74 int64 = v1
                                                        _ = v74
                                                    } else {
                                                        if v17 {
                                                            v6 = v6
                                                            if v5 {
                                                                var v75 float64 = v14
                                                                v76 := GenS10{F0: fi15, F1: fi15, F2: (fi16 | int(v18))}
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                        }
                                                        var v77 int = (fi15 % ((42) | 1))
                                                        _ = v77
                                                    }
                                                    v6 = -32768
                                                    v5 = v17
                                                }
                                                _ = v71
                                            }
                                            v14 = ((-2.25 - 0.1) - v14)
                                            _ = v70
                                        }
                                        _ = len(v2)
                                    }
                                    _ = v68
                                    _ = v69
                                } else {
                                    fmt.Println(v6, fi16, v17)
                                    if v17 {
                                        v17 = (2147483647 <= 2147483648)
                                    }
                                }
                            }
                            v4[4294967295] = (fi16 & (fi16 * -2147483648))
                        } else {
                            if v5 {
                                v14 -= (1.5 + v14)
                            } else {
                                v78 := v4[967]
                                v6 -= (v6 / (((^(v6))) | 1))
                                v79 := GenS11{F0: fi16, F1: (!((v6 % (((v6 << (((v3 >> (((v3 + (v3 | v3))) & 63))) & 63))) | 1)) != v6))}
                                _ = v78
                                _ = v79
                            }
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                            switch fi16 {
                            case 0:
                                switch (fi15 - (fi15 / (((-(fi15))) | 1))) {
                                case 1:
                                    v17 = v17
                                case 127:
                                    delete(v4, 65535)
                                }
                            case 595:
                                v2 = append(v2, uint32(v18))
                            case -641:
                                v80 := GenS6{F0: 127, F1: (!v5)}
                                _ = v80
                            default:
                                var v81 string = "a"
                                _ = v81
                            }
                        }
                        _ = v18
                    }
                    v5 = (0 <= 0)
                    _ = fi16
                    _ = v17
                }
            }
            v82 := GenS11{F0: fi15, F1: v5}
            v83 := GenS12{F0: v3, F1: 2147483647}
            _ = fi15
            _ = v82
            _ = v83
        }
        _ = v14
    }
    var v84 uint8 = 1
    v5 = (77 != 541)
    var v85 int64 = 581
    var v86 uint32 = 235
    var v87 string = ("foo bar" + "foo bar")
    for fi88 := 0; fi88 < 11; fi88++ {
        v89 := GenS7{F0: false, F1: 1}
        _ = fi88
        _ = v89
    }
    v87 = (v87 + "\n\t")
    if v5 {
        fmt.Println(v6)
        v87 = ("" + v87)
        _ = len(v4)
    }
    switch ((v87 + "hello") + "a") {
    case "hello":
        if true {
            for fi90 := 0; fi90 < 20; fi90++ {
                v84 = v84
                v5 = v5
                _ = fi90
            }
            v91 := GenS6{F0: 127, F1: (v5 || (0 == 100))}
            v1 += v1
            _ = v91
        }
    case "foo bar":
        if ((^(v3)) < 4294967295) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch (v87 + v87) {
            case "foo bar":
                if (((v6 + -1) >= 42) || true) {
                    v85 = -233
                    var v92 bool = v5
                    _ = v92
                }
            case "\n\t":
                switch 32767 {
                case 32767:
                    var v93 string = v87
                    _ = v93
                default:
                    switch 4294967295 {
                    case 128:
                        v1 = v1
                    case 802:
                        v94 := GenS13{F0: 9223372036854775807, F1: v86}
                        _ = v94
                    case 652:
                        v3 = v3
                    default:
                        var v95 string = "\n\t"
                        _ = v95
                    }
                }
            case "a":
                v87 = v87
            }
            v96 := GenS14{F0: 128}
            _ = v96
        }
    default:
        if (v5 || (18446744073709551615 < 42)) {
            var v97 int8 = -46
            var v98 uint32 = v86
            _ = v97
            _ = v98
        }
    }
    var v99 bool = (((!(4294967295 != 18446744073709551615)) || (42 == (^(v1)))) && false)
    v100 := GenS10{F0: 4294967295, F1: 128, F2: 128}
    switch 127 {
    case 616:
        for fi101 := 0; fi101 < 17; fi101++ {
            v3 = v3
            if (42 <= 42) {
                var v102 int64 = (v85 * v1)
                var v103 uint8 = 139
                _ = len(v2)
                _ = v102
                _ = v103
            } else {
                if (0.1 <= float64(1.5)) {
                    delete(v4, (v86 + v86))
                    v104 := GenS15{F0: 1, F1: v99, F2: 51}
                    for fi105 := 0; fi105 < 20; fi105++ {
                        var v106 int32 = -593
                        var v107 int64 = v1
                        v108 := GenS16{F0: (^(v1))}
                        var v109 int = (fi105 - (-(fi101)))
                        _ = fi105
                        _ = v106
                        _ = v107
                        _ = v108
                        _ = v109
                    }
                    _ = v104
                }
            }
            for fi110 := 0; fi110 < 16; fi110++ {
                var v111 uint = v3
                _ = fi110
                _ = v111
            }
            _ = fi101
        }
    case 128:
        if true {
            v5 = (1.5 < (1e10 * float32(float64(1e10))))
        }
    case 4294967295:
        switch 4294967295 {
        case 637:
            v3 = 381
        default:
            switch "a" {
            case "foo bar":
                v3 = (^(v3))
            case "a":
                _ = len(v4)
            case "":
                fmt.Println(v86)
            }
        }
    }
    var v112 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v112
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v84
    _ = v85
    _ = v86
    _ = v87
    _ = v99
}
