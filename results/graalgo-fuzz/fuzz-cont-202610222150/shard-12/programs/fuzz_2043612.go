package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 string; F1 uint; F2 uint }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int16; F1 int16; F2 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 int64; F2 bool }
type GenS8 struct { F0 string; F1 int32; F2 int8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 bool }
func genFn0(p0 int32) uint {
    v5 := []int{4294967295}
    var v6 uint32 = 2147483648
    v7 := GenS1{F0: -1, F1: 1000}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1() uint32 {
    var v8 uint32 = 65535
    v9 := []int{}
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            if true {
                v8 = v8
            }
            var v12 uint64 = 42
            switch "hello" {
            case "\n\t":
                var v13 int16 = 1
                _ = v13
            }
            _ = fi11
            _ = v12
        }
        for fi14 := 0; fi14 < 7; fi14++ {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            v8 = v8
            _ = fi14
        }
        _ = fi10
    }
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{332, 2147483647}
    v3 := []int{-679, -9223372036854775808, 9223372036854775807}
    v4 := make(map[int]uint32)
    if (uint(v1) >= 1) {
        for fi15 := 0; fi15 < 13; fi15++ {
            if true {
                if true {
                    v1 = (v1 + v1)
                    var v16 uint16 = 1
                    _ = v16
                }
                v1 += (v1 % ((809) | 1))
            } else {
                v1 = (v1 | 702)
                delete(v4, (^(fi15)))
                v17 := GenS1{F0: 9223372036854775807, F1: 1000}
                _ = v17
            }
            if false {
                var v18 int64 = 867
                _ = v18
            } else {
                var v19 uint32 = 2
                var v20 float32 = 0.1
                var v21 uint16 = 899
                _ = v19
                _ = v20
                _ = v21
            }
            if true {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(fi15, v1)
                case "":
                    if (fi15 != (fi15 % (((fi15 / ((127) | 1))) | 1))) {
                        if true {
                            var v22 uint64 = 1
                            _ = v22
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1 -= (v1 % ((v1) | 1))
                    }
                case "hello":
                    for fi23 := 0; fi23 < 2; fi23++ {
                        var v24 uint16 = 52
                        var v25 int32 = -1
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                v1 = (v1 / ((18446744073709551615) | 1))
            }
            _ = fi15
        }
        var v26 int8 = -116
        _ = v26
    } else {
        fmt.Println(v1)
    }
    delete(v4, -1000)
    var v27 string = "hello"
    if (127 > 0) {
        var v28 int = 9223372036854775807
        _ = v28
    } else {
        var v29 uint64 = v1
        switch 0 {
        case 127:
            var v30 int16 = -429
            _ = v30
        case -44:
            fmt.Println(v27)
        case 2147483647:
            v1 -= (^(v1))
        }
        _ = v29
    }
    var v31 int32 = 768
    var v32 uint16 = 2
    var v33 uint8 = 1
    if (-390 >= v31) {
        if (-780 == 1) {
            var v34 float32 = ((0.1 - 0.1) + (float32(1.5) + 3.14159))
            v27 = v27
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v34
        }
        if false {
            fmt.Println(v33)
        } else {
            if ((v33 ^ v33) >= (v33 | (v33 ^ (v33 % ((v33) | 1))))) {
                v33 -= v33
                v35 := GenS1{F0: int64(v31), F1: -750}
                _ = v35
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        if (true || true) {
            for fi36 := 0; fi36 < 14; fi36++ {
                v3 = append(v3, 42)
                _ = fi36
            }
            var v37 int8 = -1
            _ = v37
        } else {
            switch ("\n\t" + "") {
            case "":
                var v38 uint16 = (v32 / ((v32) | 1))
                _ = v38
            case "a":
                var v39 int64 = 91
                _ = v39
            case "hello":
                var v40 uint32 = 2
                _ = v40
            }
            if (253 > (v32 ^ 255)) {
                for fi41 := 0; fi41 < 15; fi41++ {
                    if (127 < fi41) {
                        if false {
                            var v42 uint16 = (v32 & 255)
                            var v43 int8 = 100
                            _ = v42
                            _ = v43
                        } else {
                            _ = len(v4)
                            _ = len(v4)
                            fmt.Println(v32, v33, v1)
                        }
                        switch v27 {
                        case "hello":
                            var v44 uint64 = 350
                            _ = v44
                        case "\n\t":
                            for fi45 := 0; fi45 < 15; fi45++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(v27)
                                v46 := GenS1{F0: 412, F1: -32768}
                                _ = fi45
                                _ = v46
                            }
                        case "":
                            switch fi41 {
                            case 129:
                                if ((true && (42 >= 1)) && (true || false)) {
                                    switch fi41 {
                                    case -2147483648:
                                        var v47 uint64 = v1
                                        _ = v47
                                    default:
                                        v32 -= v32
                                    }
                                    if (fi41 < (-(fi41))) {
                                        v4[(fi41 - (fi41 >> ((518) & 63)))] = 440
                                        switch fi41 {
                                        case 127:
                                            for fi48 := 0; fi48 < 10; fi48++ {
                                                fmt.Println(fi48, v1, fi48)
                                                v1 = 23
                                                _ = fi48
                                            }
                                        case 615:
                                            v33 = 128
                                        }
                                        v49 := GenS2{F0: (v27 + v27), F1: 519, F2: genFn0(v31)}
                                        _ = v49
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = 932
                                        }
                                        for fi50 := 0; fi50 < 8; fi50++ {
                                            fmt.Println(fi41, v27)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = fi50
                                        }
                                    }
                                }
                            case 32767:
                                if (449 == 0) {
                                    _ = len(v3)
                                    var v51 bool = ((!true) || (v27 == "a"))
                                    switch fi41 {
                                    case -2147483648:
                                        for fi52 := 0; fi52 < 13; fi52++ {
                                            var v53 uint16 = v32
                                            var v54 uint64 = 869
                                            v31 += v31
                                            if v51 {
                                                if v51 {
                                                    v55 := GenS3{F0: -560, F1: ""}
                                                    var v56 int32 = v31
                                                    var v57 uint16 = v32
                                                    var v58 int32 = v31
                                                    if (v56 <= -362) {
                                                        _ = len(v2)
                                                        var v59 float64 = -2.25
                                                        v53 = (^(v53))
                                                        _ = v59
                                                    } else {
                                                        v31 += v56
                                                        var v60 int64 = -1
                                                        fmt.Println(v1, v58, v32)
                                                        _ = v60
                                                    }
                                                    var v61 float32 = -2.25
                                                    fmt.Println(v61, v58)
                                                    _ = v55
                                                    _ = v56
                                                    _ = v57
                                                    _ = v58
                                                    _ = v61
                                                }
                                                fmt.Println(fi52, v31, v32)
                                                fmt.Println(v31)
                                            }
                                            v62 := GenS3{F0: 0, F1: "\n\t"}
                                            _ = fi52
                                            _ = v53
                                            _ = v54
                                            _ = v62
                                        }
                                    case 0:
                                        v1 = 1
                                    default:
                                        v63 := GenS4{F0: 32767, F1: -128}
                                        _ = v63
                                    }
                                    _ = v51
                                }
                            case 891:
                                switch fi41 {
                                case -316:
                                    var v64 int8 = -128
                                    _ = v64
                                case -725:
                                    fmt.Println(v33, v33)
                                case -423:
                                    var v65 float32 = (3.14159 * 1.5)
                                    _ = v65
                                }
                            }
                        }
                    } else {
                        for fi66 := 0; fi66 < 5; fi66++ {
                            if false {
                                v32 -= v32
                                var v67 uint64 = v1
                                v4[fi41] = uint32(v32)
                                switch (fi66 / (((-(fi41))) | 1)) {
                                case 9223372036854775807:
                                    var v68 string = ""
                                    _ = v68
                                case -2147483648:
                                    var v69 int16 = 1000
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    var v70 int16 = v69
                                    _ = v69
                                    _ = v70
                                case 179:
                                    if len(v2) > 0 {
                                        v2[0] = genFn1()
                                    }
                                    v71 := GenS4{F0: -964, F1: -1}
                                    v72 := GenS3{F0: 42, F1: "\n\t"}
                                    _ = v71
                                    _ = v72
                                }
                                fmt.Println(fi66)
                                v27 = v27
                                _ = v67
                            }
                            var v73 uint = 65535
                            _ = fi66
                            _ = v73
                        }
                        _ = len(v2)
                    }
                    var v74 uint8 = 127
                    _ = fi41
                    _ = v74
                }
                v33 += 229
                for fi75 := 0; fi75 < 18; fi75++ {
                    v27 = v27
                    v27 = ""
                    _ = fi75
                }
            } else {
                v27 = v27
                if false {
                    v76 := GenS5{F0: -1, F1: 1, F2: -32768}
                    for fi77 := 0; fi77 < 5; fi77++ {
                        switch 42 {
                        case -81:
                            for fi78 := 0; fi78 < 20; fi78++ {
                                if (false && true) {
                                    fmt.Println(fi78, fi78)
                                }
                                _ = fi78
                            }
                        case 265:
                            var v79 float64 = 0.1
                            _ = v79
                        case 377:
                            v33 = v33
                        }
                        fmt.Println(v33, v33, v32)
                        for fi80 := 0; fi80 < 3; fi80++ {
                            var v81 bool = true
                            v82 := GenS6{F0: (v33 ^ 0)}
                            _ = fi80
                            _ = v81
                            _ = v82
                        }
                        _ = fi77
                    }
                    var v83 uint8 = 42
                    _ = v76
                    _ = v83
                } else {
                    v32 = (v32 + v32)
                    v31 = 0
                }
            }
            fmt.Println(v33)
        }
    }
    var v84 uint32 = 4294967295
    switch 32768 {
    case -230:
        _ = len(v4)
    case -177:
        if false {
            var v85 uint = 42
            for fi86 := 0; fi86 < 2; fi86++ {
                switch fi86 {
                case 4294967295:
                    switch "foo bar" {
                    case "":
                        var v87 bool = ((3.14159 < 0.1) && true)
                        _ = v87
                    case "foo bar":
                        var v88 float32 = ((1e10 / 1e10) - 1.5)
                        _ = v88
                    case "a":
                        v31 -= 0
                    default:
                        if false {
                            for fi89 := 0; fi89 < 16; fi89++ {
                                fmt.Println(fi86, v85, v85)
                                switch 502 {
                                case 2147483647:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case 534:
                                    v27 = (v27 + v27)
                                case 9223372036854775807:
                                    if (uint32(v32) >= v84) {
                                        var v90 int8 = int8(v85)
                                        if (!(v1 < 566)) {
                                            switch -609 {
                                            case 996:
                                                if (42 > -577) {
                                                    var v91 uint64 = (v1 / ((uint64(v33)) | 1))
                                                    v92 := GenS7{F0: (v31 + v31), F1: 42, F2: (uint64(v33) >= 771)}
                                                    var v93 int = -9223372036854775808
                                                    _ = v91
                                                    _ = v92
                                                    _ = v93
                                                }
                                            case -298:
                                                if (v27 < (((("\n\t" + v27) + v27) + v27) + "a")) {
                                                    v94 := GenS6{F0: (v33 >> (((v85 >> ((genFn0(-1)) & 63))) & 63))}
                                                    if true {
                                                        var v95 uint8 = uint8(v32)
                                                        fmt.Println(v27)
                                                        if (true || false) {
                                                            v96 := v4[int(v90)]
                                                            if (32767 <= 2) {
                                                                var v97 float64 = float64(0.0)
                                                                fmt.Println(v96)
                                                                var v98 int64 = int64(v31)
                                                                _ = v97
                                                                _ = v98
                                                            }
                                                            v90 = (v90 >> ((v85) & 63))
                                                            _ = v96
                                                        }
                                                        _ = v95
                                                    } else {
                                                        delete(v4, (fi86 * fi86))
                                                    }
                                                    _ = v94
                                                } else {
                                                    v4[-317] = genFn1()
                                                    var v99 bool = (true || (448 == 1))
                                                    _ = v99
                                                }
                                            case -510:
                                                var v100 int16 = 1000
                                                _ = v100
                                            }
                                        } else {
                                            if ((((true || true) || true) || (uint8(v1) >= v33)) || (int8(v1) >= (v90 << (((v85 * v85)) & 63)))) {
                                                v101 := GenS8{F0: "foo bar", F1: v31, F2: 0}
                                                fmt.Println(v1, v90, v85)
                                                _ = v101
                                            } else {
                                                var v102 uint64 = v1
                                                if (v27 != ("\n\t" + ((v27 + (((v27 + ((((v27 + (("a" + ("" + "\n\t")) + v27)) + v27) + "\n\t") + (((v27 + v27) + "foo bar") + v27))) + "hello") + v27)) + v27))) {
                                                    v33 -= (v33 * 2)
                                                    v103 := GenS9{F0: 967}
                                                    _ = v103
                                                }
                                                _ = v102
                                            }
                                            fmt.Println(v32, fi89)
                                            if false {
                                                v33 = v33
                                                v85 += 815
                                                v104 := GenS3{F0: 223, F1: v27}
                                                _ = v104
                                            } else {
                                                if ((v31 < -108) && ((-1 <= -85) || true)) {
                                                    var v105 int64 = 0
                                                    switch (("a" + v27) + v27) {
                                                    case "hello":
                                                        if true {
                                                            v32 += v32
                                                            v84 = uint32(v33)
                                                            fmt.Println(fi89, v33)
                                                        }
                                                    case "":
                                                        fmt.Println(fi86)
                                                    }
                                                    v90 = v90
                                                    var v106 uint = (v85 * uint(v84))
                                                    var v107 uint = genFn0((v31 / ((int32(v106)) | 1)))
                                                    _ = v105
                                                    _ = v106
                                                    _ = v107
                                                }
                                            }
                                        }
                                        if false {
                                            if false {
                                                fmt.Println(v1)
                                            } else {
                                                if (v31 > 197) {
                                                    fmt.Println(fi86)
                                                    v108 := GenS6{F0: (v33 - 255)}
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v108
                                                } else {
                                                    if (0 >= -69) {
                                                        v33 += (v33 << ((v85) & 63))
                                                        if true {
                                                            _ = len(v3)
                                                        }
                                                    }
                                                    v4[(fi86 & 1)] = v84
                                                }
                                                fmt.Println(v90)
                                            }
                                        }
                                        _ = v90
                                    }
                                }
                                v109 := GenS10{F0: (v33 & (v33 - (v33 ^ 127)))}
                                _ = fi89
                                _ = v109
                            }
                        }
                    }
                case -1:
                    for fi110 := 0; fi110 < 14; fi110++ {
                        var v111 float64 = 0.0
                        _ = len(v2)
                        v111 -= v111
                        _ = fi110
                        _ = v111
                    }
                default:
                    v33 += (v33 ^ (v33 * (^(v33))))
                }
                for fi112 := 0; fi112 < 4; fi112++ {
                    var v113 uint = (v85 << ((0) & 63))
                    _ = fi112
                    _ = v113
                }
                _ = fi86
            }
            _ = v85
        } else {
            if false {
                v33 -= 42
                var v114 int32 = v31
                switch 1 {
                case -2147483648:
                    v84 = genFn1()
                case 32767:
                    if false {
                        var v115 uint64 = (v1 | (v1 & (v1 * v1)))
                        _ = v115
                    }
                case 2147483647:
                    if len(v2) > 0 {
                        v2[0] = (v84 | (v84 >> ((18446744073709551615) & 63)))
                    }
                }
                _ = v114
            } else {
                fmt.Println(v84, v27, v31)
                fmt.Println(v33, v27, v84)
            }
            v32 -= (v32 - v32)
            v31 = v31
        }
    }
    v116 := GenS11{F0: (!false)}
    switch -2147483648 {
    case 127:
        fmt.Println(v31)
    case 32768:
        var v117 int8 = -116
        _ = v117
    }
    v118 := GenS2{F0: "a", F1: 42, F2: 440}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v84:", v84)
    _ = v1
    _ = v116
    _ = v118
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v84
}
