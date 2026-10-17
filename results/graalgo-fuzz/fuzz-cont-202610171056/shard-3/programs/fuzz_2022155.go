package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int; F2 uint32 }
type GenS3 struct { F0 string; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int8; F1 int16; F2 uint }
type GenS5 struct { F0 int; F1 uint64; F2 int }
type GenS6 struct { F0 int8; F1 string; F2 uint32 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 uint32; F1 int16; F2 int64 }
type GenS9 struct { F0 int32; F1 int64; F2 int8 }
type GenS10 struct { F0 int32; F1 uint8; F2 uint32 }
func genFn0() int64 {
    v5 := GenS1{F0: 198, F1: -32768, F2: 2147483647}
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint64 = 64
            fmt.Println(fi6)
            var v9 uint8 = 159
            if ((fi6 - (-(fi7))) <= (fi7 & fi6)) {
                v9 += (v9 % (((^(v9))) | 1))
                v10 := GenS1{F0: (v9 / (((^(v9))) | 1)), F1: 1, F2: 275}
                _ = v10
            } else {
                switch ("" + ("\n\t" + "foo bar")) {
                case "hello":
                    var v11 float32 = ((float32(3.14159) * 3.14159) - 0.1)
                    _ = v11
                case "foo bar":
                    v12 := []int{}
                    _ = v12
                }
                if (false && (!false)) {
                    v13 := GenS2{F0: 426, F1: (fi6 >> ((255) & 63)), F2: 198}
                    _ = v13
                }
                var v14 bool = true
                _ = v14
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        var v15 uint8 = 2
        _ = fi6
        _ = v15
    }
    _ = v5
    return -885
}
func genFn1(p0 int64, p1 uint16) int64 {
    v16 := GenS1{F0: 42, F1: 0, F2: 65535}
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            if false {
                switch (fi17 >> ((2) & 63)) {
                case 9223372036854775807:
                    var v19 int8 = 1
                    _ = v19
                default:
                    p0 += genFn0()
                }
                v20 := []int{(fi18 + -2147483648), 127, 719}
                _ = v20
            } else {
                v21 := []int{int(p1), -709, -997}
                _ = v21
            }
            if true {
                v22 := []int{9223372036854775807}
                p0 -= (p0 << ((42) & 63))
                if (2147483647 <= -1) {
                    p0 += -536
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                } else {
                    var v23 string = "foo bar"
                    _ = v23
                }
                _ = v22
            }
            v24 := GenS2{F0: 0, F1: (fi17 + 4294967295), F2: 4294967295}
            _ = fi18
            _ = v24
        }
        if (2 >= 42) {
            if (0 >= 900) {
                var v25 uint32 = 2
                _ = v25
            }
            var v26 int64 = p0
            v27 := []int{}
            _ = v26
            _ = v27
        } else {
            var v28 int64 = p0
            _ = v28
        }
        p0 -= (^(p0))
        _ = fi17
    }
    fmt.Println(p1, p1)
    v29 := []int{-9223372036854775808, -117}
    v30 := GenS2{F0: p1, F1: -9223372036854775808, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v16
    _ = v29
    _ = v30
    return 42
}

func main() {
    v1 := []string{"hello", ""}
    var v2 uint16 = 42
    var v3 uint32 = 1
    v4 := []int{0, 2147483647, 32767}
    var v31 bool = false
    if (32767 != -1) {
        v4 = append(v4, -430)
    }
    if ((v3 + v3) <= v3) {
        if (((!((65535 >= 65535) || v31)) || false) || (0 == 18446744073709551615)) {
            fmt.Println(v2)
            _ = len(v4)
        }
    }
    if v31 {
        switch "\n\t" {
        case "a":
            switch 9223372036854775807 {
            case 901:
                v2 -= (v2 | (v2 % ((16) | 1)))
            default:
                var v32 int = -9223372036854775808
                _ = v32
            }
        case "":
            var v33 uint16 = uint16(v3)
            _ = v33
        }
        var v34 float64 = (0.1 + 1.5)
        for fi35 := 0; fi35 < 6; fi35++ {
            switch ("hello" + "") {
            case "hello":
                v34 -= v34
            case "foo bar":
                var v36 uint16 = v2
                _ = v36
            default:
                var v37 int16 = -467
                _ = v37
            }
            _ = fi35
        }
        _ = v34
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if v31 {
        if v31 {
            var v38 uint = 4294967295
            var v39 string = ""
            if (4294967295 != 597) {
                v38 = 255
                var v40 int32 = 2147483647
                _ = v40
            } else {
                v41 := GenS2{F0: 255, F1: 2147483647, F2: (v3 - (v3 % ((42) | 1)))}
                for fi42 := 0; fi42 < 10; fi42++ {
                    if (391 >= 1) {
                        for fi43 := 0; fi43 < 12; fi43++ {
                            var v44 bool = (v39 >= "\n\t")
                            v45 := GenS2{F0: (v2 / (((v2 * uint16(v38))) | 1)), F1: (fi43 - fi43), F2: (v3 % ((v3) | 1))}
                            if (1 != v2) {
                                switch "hello" {
                                case "":
                                    var v46 int16 = 1
                                    _ = v46
                                case "\n\t":
                                    var v47 int32 = 115
                                    _ = v47
                                case "hello":
                                    _ = len(v1)
                                }
                            } else {
                                v31 = (0.1 < 3.14159)
                                fmt.Println(v31)
                                var v48 int8 = 100
                                _ = v48
                            }
                            if len(v4) > 0 {
                                v4[0] = (fi43 * (fi43 | fi43))
                            }
                            switch (fi43 + fi42) {
                            case -428:
                                v38 = (v38 | v38)
                            case 32768:
                                var v49 int8 = 100
                                _ = v49
                            }
                            if v44 {
                                fmt.Println(v3, v38)
                                v44 = (v39 >= v39)
                                v50 := GenS2{F0: (v2 >> (((v38 | 42)) & 63)), F1: fi42, F2: 65535}
                                _ = v50
                            } else {
                                _ = len(v1)
                            }
                            _ = fi43
                            _ = v44
                            _ = v45
                        }
                    } else {
                        var v51 uint32 = 2
                        var v52 string = v39
                        _ = v51
                        _ = v52
                    }
                    _ = len(v4)
                    if (!(((3.14159 + -2.25) * 0.1) >= -2.25)) {
                        v53 := GenS2{F0: (^(v2)), F1: fi42, F2: (v3 / (((v3 | v3)) | 1))}
                        switch "foo bar" {
                        case "":
                            switch fi42 {
                            case -2147483648:
                                var v54 int16 = 0
                                _ = v54
                            case -355:
                                for fi55 := 0; fi55 < 1; fi55++ {
                                    v2 = v2
                                    switch v39 {
                                    case "\n\t":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "a":
                                        switch 2147483647 {
                                        case 237:
                                            var v56 bool = (18446744073709551615 <= (^(v38)))
                                            _ = v56
                                        case -956:
                                            v38 = v38
                                        case -569:
                                            var v57 string = "foo bar"
                                            _ = v57
                                        }
                                    default:
                                        switch fi42 {
                                        case -1:
                                            fmt.Println(v2, v3)
                                        case 32767:
                                            var v58 string = ("a" + v39)
                                            _ = v58
                                        case 2147483647:
                                            _ = len(v1)
                                        }
                                    }
                                    var v59 uint8 = 42
                                    _ = fi55
                                    _ = v59
                                }
                            default:
                                if (false || true) {
                                    if v31 {
                                        var v60 int16 = 46
                                        for fi61 := 0; fi61 < 19; fi61++ {
                                            var v62 int8 = 42
                                            _ = fi61
                                            _ = v62
                                        }
                                        _ = v60
                                    } else {
                                        var v63 uint16 = 78
                                        v39 = (((v39 + (("" + v39) + "foo bar")) + v39) + v39)
                                        _ = v63
                                    }
                                    for fi64 := 0; fi64 < 18; fi64++ {
                                        v65 := GenS2{F0: v2, F1: fi64, F2: 361}
                                        v66 := GenS3{F0: "foo bar", F1: (v2 / (((v2 * (v2 << (((v38 - 2)) & 63)))) | 1)), F2: 541}
                                        v31 = (42 != -2147483648)
                                        _ = fi64
                                        _ = v65
                                        _ = v66
                                    }
                                    var v67 float64 = ((0.1 * (((float64(0.0) - 3.14159) / ((0.1 * (0.1 - (0.1 - 3.14159))) + 1.5)) - 0.1)) - (0.1 / 1e10))
                                    _ = v67
                                } else {
                                    v39 = v39
                                }
                            }
                        case "foo bar":
                            for fi68 := 0; fi68 < 14; fi68++ {
                                v1 = append(v1, (v39 + "foo bar"))
                                var v69 uint32 = v3
                                if v31 {
                                    if ((fi42 == fi42) || (363 >= -1)) {
                                        if (42 != -32768) {
                                            v70 := GenS4{F0: 1, F1: 0, F2: (^(v38))}
                                            if v31 {
                                                var v71 float32 = (1e10 - 1e10)
                                                var v72 bool = ((v71 * float32(0.1)) == (float32(3.14159) - ((v71 + (v71 - v71)) + v71)))
                                                v39 = v39
                                                _ = v71
                                                _ = v72
                                            } else {
                                                v73 := GenS3{F0: "\n\t", F1: v2, F2: 4294967295}
                                                v3 = (v3 / ((2147483648) | 1))
                                                v38 -= v38
                                                _ = v73
                                            }
                                            fmt.Println(fi42)
                                            _ = v70
                                        }
                                        if v31 {
                                            v74 := GenS3{F0: v39, F1: v2, F2: 927}
                                            fmt.Println(v69)
                                            v75 := GenS5{F0: fi42, F1: 243, F2: (fi42 - fi42)}
                                            _ = len(v1)
                                            var v76 uint64 = 1
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                        }
                                        fmt.Println(fi42)
                                        _ = len(v4)
                                    } else {
                                        var v77 int = (fi68 ^ -1)
                                        var v78 int = fi42
                                        v4 = append(v4, v77)
                                        _ = v77
                                        _ = v78
                                    }
                                }
                                _ = fi68
                                _ = v69
                            }
                        case "a":
                            if (1 == 484) {
                                switch "a" {
                                case "a":
                                    v79 := GenS6{F0: -128, F1: v39, F2: (^(v3))}
                                    _ = v79
                                case "foo bar":
                                    v80 := GenS2{F0: (v2 + 1), F1: 1, F2: (v3 + (v3 << (((v38 ^ (v38 >> ((v38) & 63)))) & 63)))}
                                    _ = v80
                                }
                            } else {
                                switch fi42 {
                                case 32768:
                                    var v81 uint64 = 247
                                    _ = v81
                                case 42:
                                    fmt.Println(v2, v31)
                                case -402:
                                    _ = len(v4)
                                }
                                var v82 string = (v39 + ((v39 + v39) + v39))
                                _ = v82
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v53
                    } else {
                        for fi83 := 0; fi83 < 1; fi83++ {
                            var v84 uint16 = 1
                            _ = fi83
                            _ = v84
                        }
                        var v85 int64 = 0
                        _ = v85
                    }
                    _ = fi42
                }
                v39 = (v39 + v39)
                _ = v41
            }
            _ = v38
            _ = v39
        } else {
            var v86 string = ""
            var v87 int32 = 0
            _ = v86
            _ = v87
        }
        var v88 float32 = -2.25
        _ = v88
    }
    var v89 int8 = -10
    if false {
        v1 = append(v1, "foo bar")
        if (-1 == 1) {
            v90 := GenS6{F0: (v89 % ((-115) | 1)), F1: "foo bar", F2: 396}
            _ = v90
        } else {
            for fi91 := 0; fi91 < 7; fi91++ {
                v92 := GenS7{F0: 1, F1: 0}
                var v93 int = (fi91 >> ((uint(v2)) & 63))
                for fi94 := 0; fi94 < 19; fi94++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi94
                }
                _ = fi91
                _ = v92
                _ = v93
            }
            for fi95 := 0; fi95 < 8; fi95++ {
                if ("foo bar" <= "a") {
                    var v96 uint16 = v2
                    v96 = v96
                    v97 := GenS7{F0: -938, F1: v89}
                    _ = v96
                    _ = v97
                }
                _ = fi95
            }
        }
    } else {
        v98 := GenS5{F0: -246, F1: 552, F2: -2147483648}
        if (v89 > (v89 ^ (v89 / ((v89) | 1)))) {
            fmt.Println(v2)
            v99 := GenS8{F0: v3, F1: 42, F2: -390}
            _ = v99
        } else {
            if false {
                v100 := GenS9{F0: 0, F1: -1, F2: v89}
                v89 = 1
                if (!false) {
                    switch 9223372036854775807 {
                    case 1:
                        var v101 string = "a"
                        _ = v101
                    case 128:
                        switch "\n\t" {
                        case "hello":
                            switch 4294967295 {
                            case 946:
                                if ((v89 + (v89 + 100)) >= int8(v2)) {
                                    switch -574 {
                                    case -1:
                                        v102 := GenS10{F0: 2147483647, F1: 0, F2: (v3 << ((18446744073709551615) & 63))}
                                        _ = v102
                                    }
                                    var v103 int16 = 32767
                                    _ = len(v1)
                                    _ = v103
                                }
                            default:
                                v3 -= 2
                            }
                        case "a":
                            switch "" {
                            case "":
                                for fi104 := 0; fi104 < 18; fi104++ {
                                    var v105 float64 = -2.25
                                    v106 := GenS10{F0: 42, F1: 2, F2: v3}
                                    _ = fi104
                                    _ = v105
                                    _ = v106
                                }
                            }
                        case "":
                            var v107 int = -333
                            _ = v107
                        default:
                            v89 += (v89 % ((v89) | 1))
                        }
                    }
                    if (v2 < 2) {
                        if true {
                            switch 9223372036854775807 {
                            case 128:
                                v3 = 29
                            case 32767:
                                v4 = append(v4, -105)
                            case 0:
                                v89 += v89
                            default:
                                for fi108 := 0; fi108 < 13; fi108++ {
                                    if (v31 && (v31 && (!true))) {
                                        v1 = append(v1, "")
                                    }
                                    _ = len(v1)
                                    _ = fi108
                                }
                            }
                            for fi109 := 0; fi109 < 4; fi109++ {
                                fmt.Println(fi109, v31, v3)
                                _ = fi109
                            }
                            if (!((2 != 42) || (!(!((true && (-9223372036854775808 < -753)) || ((!(!(v31 && (("" + "foo bar") != "a")))) && (!((1e10 / 1.5) > float64(1e10))))))))) {
                                var v110 int32 = 1
                                v111 := GenS4{F0: v89, F1: 0, F2: 150}
                                v3 += (v3 << ((42) & 63))
                                _ = v110
                                _ = v111
                            } else {
                                fmt.Println(v2, v3)
                                v112 := GenS7{F0: -188, F1: v89}
                                v2 = v2
                                _ = v112
                            }
                        }
                    }
                } else {
                    var v113 uint32 = uint32(v2)
                    _ = v113
                }
                _ = v100
            } else {
                _ = len(v1)
                fmt.Println(v89, v3)
            }
            fmt.Println(v2, v2, v89)
        }
        v89 -= v89
        _ = v98
    }
    v2 = (v2 % ((v2) | 1))
    if v31 {
        _ = len(v1)
    } else {
        v3 = v3
        v1 = append(v1, "a")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v89
}
