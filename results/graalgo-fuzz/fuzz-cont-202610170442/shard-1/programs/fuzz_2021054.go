package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 int16; F2 int64 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 uint16; F1 string; F2 uint8 }
type GenS9 struct { F0 int; F1 string }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int8 }
type GenS15 struct { F0 uint; F1 string }
type GenS16 struct { F0 uint; F1 int64 }
type GenS17 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{}
    var v3 uint64 = 4294967295
    v4 := GenS1{F0: -128, F1: 0}
    var v5 uint32 = 42
    var v6 int = -9223372036854775808
    for fi7 := 0; fi7 < 9; fi7++ {
        switch v6 {
        case 128:
            v6 -= (v6 + 0)
        case 1:
            v6 -= v6
        default:
            switch "a" {
            case "hello":
                _ = len(v1)
            case "":
                for fi8 := 0; fi8 < 8; fi8++ {
                    v9 := GenS1{F0: 127, F1: (v3 / (((v3 * (v3 / ((4294967295) | 1)))) | 1))}
                    _ = fi8
                    _ = v9
                }
            case "\n\t":
                v10 := GenS1{F0: 1, F1: 2}
                _ = v10
            default:
                var v11 uint32 = 844
                _ = v11
            }
        }
        switch "a" {
        case "\n\t":
            for fi12 := 0; fi12 < 11; fi12++ {
                _ = len(v1)
                _ = fi12
            }
        case "foo bar":
            switch "hello" {
            case "hello":
                v5 -= 811
            case "a":
                fmt.Println(v6, v6)
            default:
                v1[(v5 ^ v5)] = v5
            }
        case "":
            for fi13 := 0; fi13 < 6; fi13++ {
                v14 := GenS2{F0: 1, F1: (v3 - v3)}
                switch "\n\t" {
                case "\n\t":
                    if ((true && (!((3.14159 != ((((-2.25 - 1e10) + ((((((((((0.1 + (1.5 / (1.5 + 0.1))) - (0.0 * 1.5)) - (1e10 / (-2.25 + -2.25))) * float64((0.0 + (-2.25 - 3.14159)))) * ((float64(1e10) / 1.5) * ((-2.25 + (((float64(-2.25) / 1.5) / 1.0) + ((float64((float32(((float64(((float32(float64(1e10)) * 0.0) * float32((1e10 + 0.0)))) / (float64((-2.25 + 0.0)) * 1e10)) / (float64((-2.25 - 0.1)) + (0.1 - 1.5)))) * -2.25)) * 0.1) + ((3.14159 * 0.1) * ((1e10 - 1.5) * 3.14159))))) / 1.0))) / ((1.5 - 1e10) + (((((0.1 + 0.1) + 0.0) * float64((0.0 + 0.1))) + (0.1 * 3.14159)) * 0.0))) / 0.1) - ((float64(3.14159) - (0.1 + (0.1 + 1.5))) / 1.0)) - 1.5) - 0.0)) * 1.5) * float64(-2.25))) && (!true)))) && true) {
                        if true {
                            v6 = (v6 + -1)
                            var v15 int = 941
                            _ = v15
                        }
                    }
                default:
                    var v16 float64 = -2.25
                    _ = v16
                }
                v17 := GenS3{F0: "foo bar"}
                _ = fi13
                _ = v14
                _ = v17
            }
        default:
            if (0.0 == 1.5) {
                switch 0 {
                case 32767:
                    if true {
                        v18 := GenS1{F0: 122, F1: v3}
                        switch "hello" {
                        case "":
                            var v19 float64 = 0.0
                            _ = v19
                        default:
                            var v20 int32 = 0
                            _ = v20
                        }
                        if false {
                            v21 := GenS1{F0: 0, F1: (v3 << ((935) & 63))}
                            _ = v21
                        }
                        _ = v18
                    } else {
                        var v22 int = 32767
                        if ((false || (int16(v3) != 42)) || false) {
                            if false {
                                for fi23 := 0; fi23 < 19; fi23++ {
                                    var v24 uint16 = 1
                                    switch fi7 {
                                    case 4294967295:
                                        switch "a" {
                                        case "foo bar":
                                            v25 := GenS1{F0: 42, F1: (v3 >> ((255) & 63))}
                                            _ = v25
                                        case "a":
                                            var v26 uint = 236
                                            _ = v26
                                        case "\n\t":
                                            v27 := GenS3{F0: "a"}
                                            var v28 uint64 = 42
                                            var v29 bool = false
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        default:
                                            var v30 int16 = -140
                                            _ = v30
                                        }
                                    case 128:
                                        v5 += v5
                                    case 0:
                                        switch "hello" {
                                        case "a":
                                            v6 = fi23
                                        case "foo bar":
                                            fmt.Println(v24, v6, v22)
                                        case "":
                                            var v31 uint = 4294967295
                                            _ = v31
                                        }
                                    }
                                    if (-2147483648 != -2147483648) {
                                        delete(v1, (v5 ^ v5))
                                        v32 := v1[(v5 << ((312) & 63))]
                                        v6 = (fi23 + (^(fi7)))
                                        _ = v32
                                    } else {
                                        var v33 int64 = int64(v6)
                                        v3 = 42
                                        if false {
                                            if (18446744073709551615 > 2) {
                                                if len(v2) > 0 {
                                                    v2[0] = 180
                                                }
                                                var v34 int = (fi23 - v22)
                                                var v35 int = fi7
                                                _ = v34
                                                _ = v35
                                            }
                                            v24 -= (v24 * (v24 - (v24 & (v24 ^ 0))))
                                            v24 = 2
                                        } else {
                                            fmt.Println(fi23)
                                        }
                                        _ = v33
                                    }
                                    _ = fi23
                                    _ = v24
                                }
                            } else {
                                var v36 int8 = 42
                                _ = v36
                            }
                            var v37 uint8 = 0
                            if len(v2) > 0 {
                                v2[0] = v5
                            }
                            _ = v37
                        }
                        v38 := v1[4294967295]
                        _ = v22
                        _ = v38
                    }
                case 543:
                    if (1.5 != (-2.25 / 1.0)) {
                        if false {
                            v39 := GenS4{F0: 4294967295, F1: -81, F2: -676}
                            var v40 int = fi7
                            v5 -= 951
                            _ = v39
                            _ = v40
                        } else {
                            var v41 int8 = 1
                            _ = v41
                        }
                        for fi42 := 0; fi42 < 1; fi42++ {
                            var v43 uint8 = 127
                            if ((^(fi42)) > v6) {
                                v43 += (^(v43))
                                var v44 int8 = 127
                                var v45 float64 = (0.1 / float64(float32(((3.14159 - -2.25) + -2.25))))
                                _ = v44
                                _ = v45
                            } else {
                                var v46 uint16 = 1
                                v43 = v43
                                v46 += v46
                                _ = v46
                            }
                            _ = fi42
                            _ = v43
                        }
                    } else {
                        v3 = (v3 | 0)
                        for fi47 := 0; fi47 < 19; fi47++ {
                            v5 -= 42
                            _ = fi47
                        }
                        if (3.14159 == (0.1 + 3.14159)) {
                            var v48 int16 = 801
                            var v49 uint64 = v3
                            _ = v48
                            _ = v49
                        }
                    }
                case 127:
                    for fi50 := 0; fi50 < 19; fi50++ {
                        v51 := GenS2{F0: 2147483647, F1: (v3 >> ((38) & 63))}
                        v5 = v5
                        _ = fi50
                        _ = v51
                    }
                }
                v52 := GenS4{F0: 200, F1: 32767, F2: -9223372036854775808}
                fmt.Println(v6, v6, v3)
                _ = v52
            }
        }
        if (-128 <= 26) {
            for fi53 := 0; fi53 < 1; fi53++ {
                var v54 uint8 = 90
                if ("a" < (("" + "\n\t") + "foo bar")) {
                    v1[v5] = v5
                    v55 := GenS5{F0: 499, F1: 0}
                    _ = v55
                }
                var v56 int32 = 42
                _ = fi53
                _ = v54
                _ = v56
            }
            switch fi7 {
            case 0:
                if true {
                    v5 = 2147483647
                    v57 := GenS6{F0: 2}
                    if true {
                        var v58 string = "a"
                        switch (^(fi7)) {
                        case 32767:
                            v1[v5] = (v5 | (v5 ^ 2))
                        case 128:
                            var v59 uint16 = uint16(v5)
                            _ = v59
                        case -949:
                            v58 = v58
                        default:
                            v60 := GenS7{F0: (v5 | v5), F1: 9}
                            _ = v60
                        }
                        _ = len(v2)
                        _ = v58
                    }
                    _ = v57
                }
            default:
                v5 = (v5 * v5)
            }
            fmt.Println(v5)
        } else {
            for fi61 := 0; fi61 < 19; fi61++ {
                fmt.Println(v6)
                v5 = 126
                _ = fi61
            }
            var v62 float64 = 1.5
            switch ("\n\t" + ("a" + "hello")) {
            case "":
                v62 = v62
            default:
                v63 := GenS5{F0: 4294967295, F1: -9223372036854775808}
                _ = v63
            }
            _ = v62
        }
        _ = fi7
    }
    for fi64 := 0; fi64 < 16; fi64++ {
        fmt.Println(v6, v6, fi64)
        _ = fi64
    }
    v6 -= (v6 + -452)
    if false {
        if true {
            fmt.Println(v6, v3, v5)
            for fi65 := 0; fi65 < 12; fi65++ {
                var v66 int32 = -1
                if ((255 >= 255) || true) {
                    v66 += v66
                    switch -510 {
                    case -121:
                        var v67 int8 = -1
                        _ = v67
                    case 128:
                        for fi68 := 0; fi68 < 6; fi68++ {
                            if false {
                                if ((false || (0 >= 844)) && (v66 == v66)) {
                                    v5 += v5
                                    var v69 uint8 = uint8(v3)
                                    _ = v69
                                } else {
                                    v5 = (v5 | (v5 ^ v5))
                                    fmt.Println(v3, v3, v66)
                                }
                            } else {
                                fmt.Println(v66, v66, v66)
                                var v70 uint64 = 1
                                _ = v70
                            }
                            _ = fi68
                        }
                    case 0:
                        v71 := GenS8{F0: 1, F1: "a", F2: 1}
                        _ = v71
                    default:
                        v72 := GenS8{F0: 42, F1: ("hello" + ""), F2: 128}
                        _ = v72
                    }
                    var v73 int32 = 42
                    _ = v73
                } else {
                    var v74 bool = ((v66 | (v66 % ((v66) | 1))) != v66)
                    v3 = 18446744073709551615
                    _ = v74
                }
                for fi75 := 0; fi75 < 20; fi75++ {
                    v66 -= 0
                    _ = fi75
                }
                _ = fi65
                _ = v66
            }
            v5 -= 2147483647
        } else {
            v76 := GenS9{F0: (v6 + v6), F1: ("" + "a")}
            if true {
                if false {
                    fmt.Println(v5, v3)
                    if (v6 == int(v3)) {
                        v77 := GenS10{F0: 42}
                        v6 -= v6
                        v3 += (v3 & (v3 | 0))
                        _ = v77
                    }
                    v6 = 9223372036854775807
                }
                switch "" {
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = (v5 | 0)
                    }
                case "":
                    for fi78 := 0; fi78 < 13; fi78++ {
                        var v79 uint32 = (v5 | 4294967295)
                        v80 := GenS4{F0: 42, F1: 42, F2: 9223372036854775807}
                        if false {
                            for fi81 := 0; fi81 < 16; fi81++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v82 := GenS3{F0: "\n\t"}
                                var v83 int32 = 42
                                _ = fi81
                                _ = v82
                                _ = v83
                            }
                        } else {
                            var v84 int = -572
                            _ = v84
                        }
                        _ = fi78
                        _ = v79
                        _ = v80
                    }
                default:
                    for fi85 := 0; fi85 < 2; fi85++ {
                        switch "a" {
                        case "\n\t":
                            switch fi85 {
                            case -2147483648:
                                var v86 uint32 = v5
                                _ = v86
                            case 9223372036854775807:
                                var v87 uint = 2
                                _ = v87
                            case 127:
                                var v88 int = (v6 * (^(fi85)))
                                _ = v88
                            default:
                                fmt.Println(v5, v3)
                            }
                        case "foo bar":
                            switch 2147483647 {
                            case 128:
                                for fi89 := 0; fi89 < 16; fi89++ {
                                    var v90 int16 = 1000
                                    _ = fi89
                                    _ = v90
                                }
                            case 545:
                                v91 := GenS11{F0: -118}
                                _ = v91
                            case 268:
                                v92 := GenS12{F0: 456, F1: 651, F2: 100}
                                _ = v92
                            }
                        default:
                            for fi93 := 0; fi93 < 11; fi93++ {
                                fmt.Println(v5)
                                _ = fi93
                            }
                        }
                        for fi94 := 0; fi94 < 1; fi94++ {
                            if true {
                                var v95 int64 = -1
                                _ = v95
                            }
                            switch "a" {
                            case "hello":
                                var v96 float32 = ((-2.25 * float32(-2.25)) - (1.5 + 1.5))
                                _ = v96
                            case "a":
                                var v97 int16 = -581
                                _ = v97
                            case "\n\t":
                                fmt.Println(v5)
                            default:
                                var v98 uint64 = 4294967295
                                _ = v98
                            }
                            _ = fi94
                        }
                        _ = fi85
                    }
                }
            }
            v99 := GenS13{F0: (v3 >> ((18446744073709551615) & 63))}
            _ = v76
            _ = v99
        }
        v100 := GenS14{F0: -1}
        if (true || true) {
            v5 += (^(v5))
            v5 += 4294967295
            if (30 >= 128) {
                if (0 > 127) {
                    v6 = (v6 << ((4294967295) & 63))
                    fmt.Println(v3)
                    var v101 int8 = 100
                    _ = v101
                }
                var v102 uint32 = 65535
                _ = v102
            }
        }
        _ = v100
    }
    var v103 uint16 = 388
    switch "" {
    case "":
        var v104 int64 = 863
        _ = v104
    case "\n\t":
        v105 := GenS15{F0: 1, F1: ""}
        _ = v105
    case "foo bar":
        v106 := GenS16{F0: uint(v3), F1: -1}
        _ = v106
    }
    for fi107 := 0; fi107 < 7; fi107++ {
        var v108 int32 = -747
        v109 := GenS17{F0: 42, F1: -703}
        _ = fi107
        _ = v108
        _ = v109
    }
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v103:", v103)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v103
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
