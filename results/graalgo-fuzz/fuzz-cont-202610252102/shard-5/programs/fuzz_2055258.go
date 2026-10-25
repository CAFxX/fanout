package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 string; F1 int8; F2 uint16 }
type GenS10 struct { F0 uint16; F1 uint32; F2 uint }
type GenS11 struct { F0 string; F1 int32; F2 string }
type GenS12 struct { F0 uint16; F1 uint8 }
type GenS13 struct { F0 uint; F1 int32; F2 int8 }
type GenS14 struct { F0 uint32; F1 int32 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 bool; F1 uint64 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 int64 }
type GenS20 struct { F0 int64; F1 bool; F2 string }

func main() {
    var v1 string = "\n\t"
    var v2 int64 = 42
    var v3 int8 = -19
    var v4 uint64 = 4294967295
    v5 := make(map[string]string)
    var v6 uint = uint(v4)
    v7 := GenS1{F0: -1}
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v6, v6)
        _ = fi8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int8 = 0
        var v11 string = ("foo bar" + ((v1 + "a") + ((v1 + v1) + (v1 + v1))))
        if (((3.14159 / 3.14159) == float32((1.5 - (-2.25 - (-2.25 / (1e10 * (1.5 + 3.14159))))))) && (!false)) {
            v3 = (v10 + (v3 % ((v3) | 1)))
            var v12 uint16 = 32767
            var v13 float32 = 0.0
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    if false {
        if (-728 > v2) {
            v5[v1] = "foo bar"
            var v14 bool = false
            if v14 {
                v4 = v4
                v15 := []int{942, 468, 32768}
                _ = v15
            } else {
                v16 := GenS2{F0: (v2 * (^(v2)))}
                _ = v16
            }
            _ = v14
        } else {
            if false {
                if true {
                    v17 := GenS3{F0: 128}
                    if (127 <= 2) {
                        var v18 bool = true
                        _ = v18
                    } else {
                        v19 := GenS4{F0: 65535, F1: uint8(v4)}
                        _ = v19
                    }
                    v20 := GenS5{F0: 394}
                    _ = v17
                    _ = v20
                } else {
                    v21 := []int{}
                    _ = v21
                }
                var v22 float32 = ((-2.25 * ((1.5 - 1.5) * float32(float64(0.1)))) - (3.14159 / 3.14159))
                for fi23 := 0; fi23 < 4; fi23++ {
                    switch (fi23 << ((4294967295) & 63)) {
                    case 127:
                        if true {
                            if false {
                                var v24 float32 = (v22 * 3.14159)
                                switch 0 {
                                case 127:
                                    var v25 float32 = (v24 * v24)
                                    _ = v25
                                default:
                                    switch v1 {
                                    case "hello":
                                        v26 := []int{}
                                        _ = v26
                                    case "a":
                                        v27 := []int{fi23, -1}
                                        _ = v27
                                    case "":
                                        switch v1 {
                                        case "a":
                                            for fi28 := 0; fi28 < 18; fi28++ {
                                                fmt.Println(fi23, v1)
                                                if ((v6 + v6) < (^(v6))) {
                                                    var v29 uint = 1
                                                    _ = v29
                                                }
                                                _ = fi28
                                            }
                                        case "foo bar":
                                            v4 = v4
                                        default:
                                            if ((v6 % (((v6 - 2)) | 1)) > 12) {
                                                var v30 float64 = (((0.1 / 1.0) + float64(((float32((0.1 / 0.1)) / v22) + v22))) + -2.25)
                                                _ = v30
                                            } else {
                                                if true {
                                                    var v31 uint32 = 1
                                                    if (v1 != ("foo bar" + v1)) {
                                                        var v32 bool = false
                                                        for fi33 := 0; fi33 < 1; fi33++ {
                                                            var v34 uint16 = 688
                                                            fmt.Println(v22)
                                                            var v35 string = ((v1 + v1) + v1)
                                                            _ = fi33
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        for fi36 := 0; fi36 < 3; fi36++ {
                                                            var v37 uint = v6
                                                            var v38 int16 = 42
                                                            _ = fi36
                                                            _ = v37
                                                            _ = v38
                                                        }
                                                        _ = v32
                                                    } else {
                                                        v31 = v31
                                                    }
                                                    _ = v31
                                                } else {
                                                    if ((3.14159 == (v24 - v24)) && (2 < 131)) {
                                                        fmt.Println(v22, v3)
                                                        switch (((v1 + ((v1 + "hello") + ((v1 + v1) + (v1 + v1)))) + v1) + ("hello" + ((v1 + "a") + v1))) {
                                                        case "a":
                                                            var v39 int32 = 1
                                                            _ = v39
                                                        case "hello":
                                                            delete(v5, ("foo bar" + ("a" + "\n\t")))
                                                        case "foo bar":
                                                            if false {
                                                                var v40 int = int(v3)
                                                                v1 = ((v1 + (v1 + v1)) + v1)
                                                                for fi41 := 0; fi41 < 10; fi41++ {
                                                                    fmt.Println(v4)
                                                                    switch -2147483648 {
                                                                    case 42:
                                                                        var v42 uint = v6
                                                                        _ = v42
                                                                    case 32767:
                                                                        if true {
                                                                            v43 := GenS6{F0: 1}
                                                                            _ = v43
                                                                        }
                                                                    case 2147483647:
                                                                        if true {
                                                                            if (-684 >= -990) {
                                                                                v44 := []int{(^(fi23)), fi41, 1}
                                                                                _ = len(v44)
                                                                                _ = v44
                                                                            }
                                                                            v40 -= 532
                                                                            v22 += (1.5 * v22)
                                                                            v45 := GenS7{F0: 576}
                                                                            _ = v45
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            v46 := []int{32767, 168}
                                                                            var v47 int32 = 768
                                                                            if (!true) {
                                                                                v48 := GenS4{F0: 0, F1: 0}
                                                                                var v49 string = v1
                                                                                var v50 bool = true
                                                                                var v51 uint64 = v4
                                                                                if true {
                                                                                    v52 := GenS8{F0: 4294967295, F1: v50}
                                                                                    v53 := GenS9{F0: "foo bar", F1: -1, F2: 527}
                                                                                    _ = v52
                                                                                    _ = v53
                                                                                } else {
                                                                                    switch v49 {
                                                                                    case "foo bar":
                                                                                        if v50 {
                                                                                            if ((-(v40)) >= (fi23 << (((^(v6))) & 63))) {
                                                                                                _ = len(v46)
                                                                                            } else {
                                                                                                v1 = "\n\t"
                                                                                                fmt.Println(v1, v50)
                                                                                                v54 := v5[(v49 + ("\n\t" + (v1 + v49)))]
                                                                                                _ = v54
                                                                                            }
                                                                                            v50 = (446 <= 654)
                                                                                            switch v49 {
                                                                                            case "hello":
                                                                                                v3 = (v3 ^ (v3 - (v3 ^ v3)))
                                                                                                v3 -= v3
                                                                                                v51 -= (v4 | v4)
                                                                                            case "\n\t":
                                                                                                v55 := GenS5{F0: 255}
                                                                                                _ = v55
                                                                                            case "":
                                                                                                var v56 bool = true
                                                                                                _ = v56
                                                                                            default:
                                                                                                v24 += float32((3.14159 + 0.1))
                                                                                            }
                                                                                        } else {
                                                                                            if true {
                                                                                                var v57 int = fi41
                                                                                                _ = v57
                                                                                            } else {
                                                                                                var v58 uint16 = 1
                                                                                                var v59 uint16 = (^(v58))
                                                                                                _ = v58
                                                                                                _ = v59
                                                                                            }
                                                                                            if v50 {
                                                                                                var v60 int16 = -1
                                                                                                if false {
                                                                                                    if (int64(v3) != 0) {
                                                                                                        v4 = (v51 * v4)
                                                                                                    }
                                                                                                    var v61 uint = (^(v6))
                                                                                                    if true {
                                                                                                        var v62 uint = (v6 | (v61 << ((273) & 63)))
                                                                                                        if ((!true) || true) {
                                                                                                            v40 = fi41
                                                                                                            var v63 uint64 = 2
                                                                                                            _ = v63
                                                                                                        }
                                                                                                        _ = v62
                                                                                                    } else {
                                                                                                        fmt.Println(v24, v6)
                                                                                                        v64 := GenS1{F0: int32(v4)}
                                                                                                        var v65 int8 = -1
                                                                                                        _ = v64
                                                                                                        _ = v65
                                                                                                    }
                                                                                                    _ = v61
                                                                                                }
                                                                                                _ = v60
                                                                                            } else {
                                                                                                v49 = (v49 + (v49 + v1))
                                                                                                switch (v40 & (fi23 / ((-2147483648) | 1))) {
                                                                                                case -2147483648:
                                                                                                    if false {
                                                                                                        switch v49 {
                                                                                                        case "foo bar":
                                                                                                            v46 = append(v46, (v40 % ((-62) | 1)))
                                                                                                        case "hello":
                                                                                                            v50 = v50
                                                                                                        case "":
                                                                                                            switch v1 {
                                                                                                            case "\n\t":
                                                                                                                v40 += (v40 % (((fi41 * fi41)) | 1))
                                                                                                            case "":
                                                                                                                v1 = (v1 + v49)
                                                                                                            case "foo bar":
                                                                                                                if v50 {
                                                                                                                    switch fi23 {
                                                                                                                    case -9223372036854775808:
                                                                                                                        fmt.Println(v6, v49)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v5[v49] = v1
                                                                                                                    if (!false) {
                                                                                                                        _ = len(v5)
                                                                                                                        var v66 int8 = v3
                                                                                                                        v3 = v3
                                                                                                                        _ = v66
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                v4 = v51
                                                                                                            }
                                                                                                            if v50 {
                                                                                                                var v67 uint16 = 32767
                                                                                                                var v68 int = (v40 % ((fi41) | 1))
                                                                                                                _ = v67
                                                                                                                _ = v68
                                                                                                            } else {
                                                                                                                fmt.Println(v50, v49, v2)
                                                                                                                if (((v2 >> ((v6) & 63)) <= (v2 & (v2 % (((v2 - -128)) | 1)))) && v50) {
                                                                                                                    v46 = append(v46, 4294967295)
                                                                                                                    var v69 uint32 = 0
                                                                                                                    _ = v69
                                                                                                                } else {
                                                                                                                    fmt.Println(v40)
                                                                                                                    switch v40 {
                                                                                                                    case 105:
                                                                                                                        var v70 int8 = v3
                                                                                                                        _ = v70
                                                                                                                    case -1:
                                                                                                                        var v71 uint32 = 42
                                                                                                                        switch int(v51) {
                                                                                                                        case 2147483647:
                                                                                                                            fmt.Println(v6, v1)
                                                                                                                        case 128:
                                                                                                                            v72 := GenS4{F0: v71, F1: 42}
                                                                                                                            _ = v72
                                                                                                                        case -720:
                                                                                                                            v1 = "\n\t"
                                                                                                                        }
                                                                                                                        _ = v71
                                                                                                                    case 365:
                                                                                                                        var v73 int8 = (^(v3))
                                                                                                                        _ = v73
                                                                                                                    default:
                                                                                                                        v2 -= (^(v2))
                                                                                                                    }
                                                                                                                    v40 -= int(v4)
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        v74 := GenS10{F0: 42, F1: 2147483648, F2: 1}
                                                                                                        _ = v74
                                                                                                    }
                                                                                                case -1:
                                                                                                    var v75 int16 = 0
                                                                                                    _ = v75
                                                                                                }
                                                                                                v76 := GenS11{F0: "hello", F1: (v47 << ((uint(v4)) & 63)), F2: v49}
                                                                                                v77 := GenS8{F0: v51, F1: (!(v50 && v50))}
                                                                                                var v78 int16 = -592
                                                                                                if false {
                                                                                                    var v79 uint64 = (v51 << (((v6 + (^(v6)))) & 63))
                                                                                                    v80 := GenS12{F0: 255, F1: 54}
                                                                                                    if true {
                                                                                                        v46 = append(v46, (fi23 % (((v40 << ((0) & 63))) | 1)))
                                                                                                        var v81 float32 = v24
                                                                                                        _ = v81
                                                                                                    }
                                                                                                    _ = v79
                                                                                                    _ = v80
                                                                                                } else {
                                                                                                    if true {
                                                                                                        if v50 {
                                                                                                            v40 = (fi23 + 536)
                                                                                                        } else {
                                                                                                            v47 += v47
                                                                                                            var v82 uint = v6
                                                                                                            _ = len(v46)
                                                                                                            _ = v82
                                                                                                        }
                                                                                                        v2 += (v2 - v2)
                                                                                                    } else {
                                                                                                        v49 = v1
                                                                                                        v83 := GenS12{F0: 42, F1: 2}
                                                                                                        var v84 int32 = v47
                                                                                                        _ = v83
                                                                                                        _ = v84
                                                                                                    }
                                                                                                    var v85 int16 = (v78 ^ 1000)
                                                                                                    if (!(!(!(-616 == (v47 << ((v6) & 63)))))) {
                                                                                                        v86 := GenS7{F0: -1}
                                                                                                        var v87 uint = 18446744073709551615
                                                                                                        v46 = append(v46, 320)
                                                                                                        var v88 float64 = (0.1 / (0.1 * 3.14159))
                                                                                                        var v89 float32 = v22
                                                                                                        _ = v86
                                                                                                        _ = v87
                                                                                                        _ = v88
                                                                                                        _ = v89
                                                                                                    }
                                                                                                    _ = v85
                                                                                                }
                                                                                                v24 = v24
                                                                                                fmt.Println(fi23, fi41, v22)
                                                                                                if (v4 == 2) {
                                                                                                    fmt.Println(v78, v3)
                                                                                                    if (v50 && ((v6 ^ (v6 << ((1) & 63))) <= 42)) {
                                                                                                        v90 := GenS9{F0: v49, F1: int8(v47), F2: 32767}
                                                                                                        fmt.Println(v6)
                                                                                                        var v91 int64 = (v2 % (((v2 & v2)) | 1))
                                                                                                        _ = v90
                                                                                                        _ = v91
                                                                                                    } else {
                                                                                                        switch ((v49 + v1) + "foo bar") {
                                                                                                        case "a":
                                                                                                            if true {
                                                                                                                v40 = (v40 * v40)
                                                                                                                if true {
                                                                                                                    v78 = v78
                                                                                                                    if v50 {
                                                                                                                        if v50 {
                                                                                                                            if (1.5 > (((v22 - v22) + 3.14159) + (1.5 - v24))) {
                                                                                                                                fmt.Println(v78, v47)
                                                                                                                                var v92 int = v40
                                                                                                                                var v93 int8 = (v3 - v3)
                                                                                                                                _ = v92
                                                                                                                                _ = v93
                                                                                                                            }
                                                                                                                        } else {
                                                                                                                            fmt.Println(v4, v22)
                                                                                                                            delete(v5, ("" + v1))
                                                                                                                        }
                                                                                                                        if len(v46) >= 2 {
                                                                                                                            v46 = v46[1:len(v46)]
                                                                                                                        }
                                                                                                                        fmt.Println(v4, v50)
                                                                                                                        v2 = -498
                                                                                                                    } else {
                                                                                                                        fmt.Println(v47, v47, v24)
                                                                                                                        _ = len(v5)
                                                                                                                        var v94 uint = (v6 + (v6 << (((v6 & (^(v6)))) & 63)))
                                                                                                                        _ = v94
                                                                                                                    }
                                                                                                                }
                                                                                                                if true {
                                                                                                                    if false {
                                                                                                                        v40 += 0
                                                                                                                    } else {
                                                                                                                        var v95 float32 = v24
                                                                                                                        v96 := GenS5{F0: (v6 / ((255) | 1))}
                                                                                                                        if ((v50 || (v49 == v1)) || (v40 > (fi23 >> (((v6 | (v6 & (v6 * (v6 / (((v6 * (v6 - v6))) | 1)))))) & 63)))) {
                                                                                                                            v97 := GenS13{F0: (v6 | v6), F1: v47, F2: 41}
                                                                                                                            v5["\n\t"] = (v1 + ((v1 + ("" + ("foo bar" + ""))) + v49))
                                                                                                                            if len(v46) >= 2 {
                                                                                                                                v46 = v46[1:len(v46)]
                                                                                                                            }
                                                                                                                            _ = v97
                                                                                                                        } else {
                                                                                                                            switch v49 {
                                                                                                                            case "foo bar":
                                                                                                                                if v50 {
                                                                                                                                    v40 = (^(fi41))
                                                                                                                                }
                                                                                                                            case "":
                                                                                                                                var v98 int32 = v47
                                                                                                                                _ = v98
                                                                                                                            case "hello":
                                                                                                                                v99 := GenS2{F0: (-(v2))}
                                                                                                                                _ = v99
                                                                                                                            }
                                                                                                                        }
                                                                                                                        _ = v95
                                                                                                                        _ = v96
                                                                                                                    }
                                                                                                                    var v100 uint = 151
                                                                                                                    if true {
                                                                                                                        v101 := GenS8{F0: v4, F1: v50}
                                                                                                                        var v102 int64 = (v2 & (v2 - v2))
                                                                                                                        var v103 int = (v40 / ((-844) | 1))
                                                                                                                        _ = v101
                                                                                                                        _ = v102
                                                                                                                        _ = v103
                                                                                                                    }
                                                                                                                    fmt.Println(v51, v3, v100)
                                                                                                                    _ = v100
                                                                                                                }
                                                                                                            } else {
                                                                                                                v5["foo bar"] = ((v1 + (v49 + ("\n\t" + v49))) + "hello")
                                                                                                                v2 = (v2 >> ((v6) & 63))
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            v40 += -9223372036854775808
                                                                                                        }
                                                                                                        v104 := GenS12{F0: 255, F1: 0}
                                                                                                        if (((fi41 / ((fi23) | 1)) < fi41) && true) {
                                                                                                            v51 -= 0
                                                                                                            _ = len(v5)
                                                                                                        }
                                                                                                        _ = v104
                                                                                                    }
                                                                                                    switch int(v78) {
                                                                                                    case 1:
                                                                                                        v22 += (-2.25 - float32(1.5))
                                                                                                    case 32767:
                                                                                                        var v105 string = v1
                                                                                                        _ = v105
                                                                                                    case -248:
                                                                                                        var v106 uint64 = v4
                                                                                                        _ = v106
                                                                                                    }
                                                                                                    _ = len(v46)
                                                                                                } else {
                                                                                                    v107 := GenS12{F0: 162, F1: 2}
                                                                                                    fmt.Println(v50, v51, v40)
                                                                                                    _ = v107
                                                                                                }
                                                                                                _ = v76
                                                                                                _ = v77
                                                                                                _ = v78
                                                                                            }
                                                                                            v108 := GenS14{F0: 65535, F1: 42}
                                                                                            switch "foo bar" {
                                                                                            case "a":
                                                                                                if (3.14159 >= ((((1e10 - float32(0.0)) - (v24 * 1.5)) - -2.25) * v24)) {
                                                                                                    var v109 int16 = 1
                                                                                                    v110 := GenS10{F0: 255, F1: 2, F2: (v6 & (v6 & v6))}
                                                                                                    _ = v109
                                                                                                    _ = v110
                                                                                                }
                                                                                            }
                                                                                            if len(v46) > 0 {
                                                                                                v46[0] = (fi23 | fi23)
                                                                                            }
                                                                                            if len(v46) >= 2 {
                                                                                                v46 = v46[1:len(v46)]
                                                                                            }
                                                                                            _ = v108
                                                                                        }
                                                                                    case "\n\t":
                                                                                        if true {
                                                                                            var v111 float32 = 0.1
                                                                                            v46 = append(v46, -2147483648)
                                                                                            fmt.Println(v6, v22, v51)
                                                                                            _ = v111
                                                                                        } else {
                                                                                            v24 -= 3.14159
                                                                                            fmt.Println(fi41, v51)
                                                                                            var v112 uint16 = 221
                                                                                            var v113 uint8 = 156
                                                                                            _ = v112
                                                                                            _ = v113
                                                                                        }
                                                                                        var v114 int8 = (v3 + int8(v2))
                                                                                        v115 := GenS6{F0: (v2 >> (((v6 & (v6 + v6))) & 63))}
                                                                                        _ = v114
                                                                                        _ = v115
                                                                                    case "a":
                                                                                        v116 := GenS4{F0: 951, F1: 127}
                                                                                        _ = v116
                                                                                    }
                                                                                    _ = len(v5)
                                                                                    if (v22 < (((v24 / (v22 - 3.14159)) + v24) + v24)) {
                                                                                        var v117 uint32 = 0
                                                                                        fmt.Println(fi41, v6, fi23)
                                                                                        _ = v117
                                                                                    } else {
                                                                                        v118 := GenS5{F0: 988}
                                                                                        v1 = ((v49 + "foo bar") + v1)
                                                                                        var v119 float64 = 1e10
                                                                                        _ = v118
                                                                                        _ = v119
                                                                                    }
                                                                                }
                                                                                _ = v48
                                                                                _ = v49
                                                                                _ = v50
                                                                                _ = v51
                                                                            }
                                                                            _ = v46
                                                                            _ = v47
                                                                        } else {
                                                                            if false {
                                                                                if (42 > 512) {
                                                                                    switch (v1 + "\n\t") {
                                                                                    case "":
                                                                                        fmt.Println(v1, v2, v22)
                                                                                    case "\n\t":
                                                                                        v5[v1] = "foo bar"
                                                                                    }
                                                                                    var v120 string = v1
                                                                                    v2 -= v2
                                                                                    _ = v120
                                                                                } else {
                                                                                    var v121 int8 = (v3 % (((^(v3))) | 1))
                                                                                    if ((v3 % ((-1) | 1)) > -128) {
                                                                                        var v122 float64 = -2.25
                                                                                        var v123 int32 = 2147483647
                                                                                        v5[v1] = v1
                                                                                        _ = v122
                                                                                        _ = v123
                                                                                    } else {
                                                                                        var v124 uint32 = 65535
                                                                                        fmt.Println(v3, v6)
                                                                                        _ = v124
                                                                                    }
                                                                                    _ = v121
                                                                                }
                                                                                v5["foo bar"] = v1
                                                                                v125 := GenS15{F0: 42}
                                                                                if false {
                                                                                    if (int(v6) == v40) {
                                                                                        if (-1 != (v3 | (v3 | (v3 ^ (v3 % (((v3 % (((v3 << ((739) & 63))) | 1))) | 1)))))) {
                                                                                            switch "foo bar" {
                                                                                            case "\n\t":
                                                                                                var v126 uint32 = 65535
                                                                                                _ = v126
                                                                                            case "":
                                                                                                v127 := []int{}
                                                                                                _ = v127
                                                                                            case "foo bar":
                                                                                                fmt.Println(fi23, v24)
                                                                                            }
                                                                                        }
                                                                                        var v128 uint16 = 32767
                                                                                        if false {
                                                                                            v129 := v5[((v1 + v1) + "\n\t")]
                                                                                            switch (fi41 % (((-(fi41))) | 1)) {
                                                                                            case 32768:
                                                                                                v2 = -9223372036854775808
                                                                                            case 1:
                                                                                                fmt.Println(v3, v4, v40)
                                                                                            case 32767:
                                                                                                var v130 uint8 = 1
                                                                                                _ = v130
                                                                                            default:
                                                                                                if true {
                                                                                                    v131 := v5["foo bar"]
                                                                                                    v129 = ""
                                                                                                    _ = v131
                                                                                                }
                                                                                            }
                                                                                            v132 := GenS15{F0: (v128 << ((18446744073709551615) & 63))}
                                                                                            _ = v129
                                                                                            _ = v132
                                                                                        }
                                                                                        var v133 int8 = (v3 << (((v6 - (v6 & v6))) & 63))
                                                                                        var v134 uint = v6
                                                                                        _ = len(v5)
                                                                                        _ = v128
                                                                                        _ = v133
                                                                                        _ = v134
                                                                                    }
                                                                                }
                                                                                _ = v125
                                                                            }
                                                                            v135 := []int{fi23, 127, (fi41 % ((fi41) | 1))}
                                                                            v40 = int(v6)
                                                                            _ = v135
                                                                        }
                                                                        v2 += v2
                                                                        v136 := v5[v1]
                                                                        _ = v136
                                                                    }
                                                                    _ = fi41
                                                                }
                                                                _ = v40
                                                            }
                                                        default:
                                                            for fi137 := 0; fi137 < 14; fi137++ {
                                                                v3 += v3
                                                                v138 := GenS7{F0: 2147483647}
                                                                v22 += (0.0 - (((1.5 * v22) / v22) - ((v24 + v24) - 1e10)))
                                                                _ = fi137
                                                                _ = v138
                                                            }
                                                        }
                                                        v139 := GenS11{F0: ("" + v1), F1: -1, F2: ""}
                                                        _ = v139
                                                    } else {
                                                        for fi140 := 0; fi140 < 4; fi140++ {
                                                            v2 = 9223372036854775807
                                                            if ((2 > 4294967295) && false) {
                                                                delete(v5, (v1 + ""))
                                                                v141 := GenS16{F0: 99}
                                                                _ = v141
                                                            } else {
                                                                v142 := []int{(fi140 << ((v6) & 63)), fi140, 32768}
                                                                var v143 int64 = (v2 >> (((v6 & 65535)) & 63))
                                                                _ = v142
                                                                _ = v143
                                                            }
                                                            _ = fi140
                                                        }
                                                        if false {
                                                            switch "foo bar" {
                                                            case "a":
                                                                var v144 int32 = 612
                                                                _ = v144
                                                            }
                                                            var v145 int64 = v2
                                                            fmt.Println(v2, v4, v6)
                                                            _ = v145
                                                        } else {
                                                            v146 := GenS9{F0: v1, F1: 87, F2: 255}
                                                            var v147 int32 = -2147483648
                                                            _ = v146
                                                            _ = v147
                                                        }
                                                        var v148 int = (^(fi23))
                                                        _ = v148
                                                    }
                                                }
                                            }
                                        }
                                    }
                                }
                                _ = v24
                            }
                            switch fi23 {
                            case 575:
                                if (fi23 <= -9223372036854775808) {
                                    v149 := []int{-9223372036854775808, fi23}
                                    _ = v149
                                } else {
                                    for fi150 := 0; fi150 < 2; fi150++ {
                                        v3 -= 100
                                        delete(v5, "")
                                        if true {
                                            if true {
                                                fmt.Println(v22, v22, fi150)
                                            }
                                            v151 := GenS17{F0: (127 < 128), F1: (v4 ^ (v4 / ((0) | 1)))}
                                            v152 := []int{fi23, (fi23 ^ int(v6)), (^(fi150)), 128}
                                            _ = v151
                                            _ = v152
                                        }
                                        var v153 uint16 = 619
                                        _ = fi150
                                        _ = v153
                                    }
                                }
                            default:
                                v154 := GenS17{F0: false, F1: (v4 ^ (^(v4)))}
                                _ = v154
                            }
                        }
                    case 9223372036854775807:
                        var v155 int = -869
                        _ = v155
                    case -743:
                        switch 127 {
                        case 32768:
                            for fi156 := 0; fi156 < 11; fi156++ {
                                var v157 uint64 = v4
                                switch fi23 {
                                case 2147483647:
                                    if (true && true) {
                                        v5[("a" + ("" + (v1 + "\n\t")))] = (((v1 + (v1 + v1)) + v1) + ((v1 + "a") + "a"))
                                    } else {
                                        v3 -= (-(v3))
                                    }
                                case 737:
                                    fmt.Println(v3)
                                case -1:
                                    v3 -= 42
                                }
                                _ = fi156
                                _ = v157
                            }
                        case 42:
                            var v158 uint = (v6 >> (((^(v6))) & 63))
                            _ = v158
                        }
                    default:
                        var v159 uint8 = 100
                        _ = v159
                    }
                    _ = fi23
                }
                _ = v22
            }
            if false {
                v3 -= (v3 >> (((v6 / ((v6) | 1))) & 63))
            }
        }
        v160 := GenS18{F0: 321}
        _ = v160
    } else {
        v6 -= (v6 % (((^(v6))) | 1))
        var v161 uint32 = 2
        if true {
            v162 := GenS19{F0: (v2 / ((v2) | 1))}
            var v163 int16 = -32768
            _ = v162
            _ = v163
        }
        _ = v161
    }
    fmt.Println(v3, v3)
    if false {
        var v164 float64 = 0.1
        var v165 int = -275
        _ = v164
        _ = v165
    } else {
        v166 := GenS20{F0: 42, F1: true, F2: v1}
        v6 += v6
        var v167 int64 = 785
        _ = v166
        _ = v167
    }
    switch "\n\t" {
    case "a":
        var v168 uint = v6
        _ = v168
    }
    if true {
        var v169 uint = v6
        var v170 uint64 = 0
        var v171 float32 = (1.5 - 1.5)
        _ = v169
        _ = v170
        _ = v171
    }
    v172 := []int{-844, 9223372036854775807}
    v1 = (v1 + "foo bar")
    for fi173 := 0; fi173 < 6; fi173++ {
        v2 = v2
        var v174 int64 = (v2 / (((v2 >> ((v6) & 63))) | 1))
        _ = fi173
        _ = v174
    }
    v175 := GenS12{F0: 251, F1: 46}
    var v176 uint8 = 130
    v1 = (v1 + v1)
    delete(v5, v1)
    fmt.Println(v3, v176, v2)
    v6 += (v6 & v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v172:", len(v172), v172)
    fmt.Println("v176:", v176)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v172
    _ = v175
    _ = v176
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
