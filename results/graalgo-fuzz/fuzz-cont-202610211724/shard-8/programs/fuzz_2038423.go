package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint64; F2 string }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64; F1 int64 }
type GenS10 struct { F0 uint32; F1 bool; F2 uint }
type GenS11 struct { F0 int64 }
func genFn0() int64 {
    var v5 string = ""
    v6 := []int{1}
    var v7 string = ((("\n\t" + ("\n\t" + v5)) + v5) + "hello")
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1() uint {
    if true {
        switch 936 {
        case 42:
            v8 := []int{9223372036854775807, 1}
            _ = v8
        case -312:
            switch ((("" + "\n\t") + "") + "\n\t") {
            case "":
                v9 := GenS1{F0: -33, F1: 972, F2: -1}
                _ = v9
            case "hello":
                v10 := []int{1, 2147483647}
                _ = v10
            default:
                switch ("foo bar" + "") {
                case "":
                    v11 := []int{9223372036854775807}
                    _ = v11
                case "hello":
                    for fi12 := 0; fi12 < 16; fi12++ {
                        var v13 string = ("hello" + "a")
                        v14 := GenS1{F0: 0, F1: 1, F2: genFn0()}
                        for fi15 := 0; fi15 < 16; fi15++ {
                            if (65535 < 255) {
                                var v16 string = (("hello" + ("foo bar" + v13)) + v13)
                                _ = v16
                            } else {
                                v13 = v13
                            }
                            v13 = ("" + (v13 + ""))
                            _ = fi15
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                default:
                    var v17 uint64 = 2
                    _ = v17
                }
            }
        case 896:
            var v18 int64 = 275
            _ = v18
        }
        if (true || false) {
            var v19 int32 = -302
            var v20 string = ""
            switch "a" {
            case "foo bar":
                v19 -= 726
            case "a":
                var v21 uint16 = 0
                _ = v21
            case "":
                var v22 int64 = -1
                _ = v22
            }
            _ = v19
            _ = v20
        } else {
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 int16 = 42
                v24 = v24
                fmt.Println(fi23, v24)
                _ = fi23
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := GenS1{F0: -1, F1: 0, F2: -775}
            var v27 string = "hello"
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    v28 := GenS2{F0: -112}
    var v29 bool = (!(-2.25 != float32((-2.25 + (1.5 + 0.0)))))
    if true {
        switch "a" {
        case "\n\t":
            if false {
                if (4294967295 == 1) {
                    var v30 uint8 = 208
                    for fi31 := 0; fi31 < 8; fi31++ {
                        fmt.Println(v29, v29, v29)
                        v30 = (v30 / ((v30) | 1))
                        _ = fi31
                    }
                    if (((-2.25 - 3.14159) / (0.1 + 0.1)) >= 0.0) {
                        v32 := GenS2{F0: 0}
                        _ = v32
                    }
                    _ = v30
                }
                fmt.Println(v29)
            }
        case "a":
            v33 := []int{-2147483648}
            _ = v33
        }
        for fi34 := 0; fi34 < 20; fi34++ {
            for fi35 := 0; fi35 < 6; fi35++ {
                var v36 int64 = 1
                switch fi35 {
                case 128:
                    if v29 {
                        switch "a" {
                        case "hello":
                            var v37 int8 = -88
                            _ = v37
                        }
                        v38 := []int{(fi34 | -115), (fi35 << ((494) & 63))}
                        v39 := GenS1{F0: -1, F1: -32768, F2: (v36 + v36)}
                        _ = v38
                        _ = v39
                    } else {
                        v40 := GenS1{F0: 42, F1: 1000, F2: v36}
                        _ = v40
                    }
                case 32768:
                    v41 := GenS1{F0: 71, F1: -32768, F2: (v36 >> ((2) & 63))}
                    _ = v41
                case 773:
                    var v42 int8 = -1
                    _ = v42
                default:
                    var v43 uint32 = 2
                    _ = v43
                }
                v36 += v36
                v44 := GenS1{F0: 42, F1: 215, F2: v36}
                v36 = v36
                _ = fi35
                _ = v36
                _ = v44
            }
            _ = fi34
        }
    }
    _ = v28
    _ = v29
    return 1
}

func main() {
    var v1 float64 = 1e10
    var v2 int64 = -631
    var v3 int64 = v2
    v4 := []int{}
    v3 -= 114
    v3 += v3
    var v45 uint16 = 1
    switch "a" {
    case "a":
        var v46 uint64 = 1
        _ = v46
    case "foo bar":
        if (!true) {
            v1 += 1.5
        } else {
            if (false && (true || (0 != 348))) {
                if (2 < 2147483648) {
                    for fi47 := 0; fi47 < 10; fi47++ {
                        v1 += ((v1 * 0.0) * v1)
                        for fi48 := 0; fi48 < 4; fi48++ {
                            v2 += int64(fi48)
                            v49 := GenS1{F0: -27, F1: 1, F2: (v3 % ((v2) | 1))}
                            _ = fi48
                            _ = v49
                        }
                        v50 := GenS2{F0: -1}
                        _ = fi47
                        _ = v50
                    }
                }
                v2 = v3
            }
            for fi51 := 0; fi51 < 6; fi51++ {
                for fi52 := 0; fi52 < 7; fi52++ {
                    v53 := GenS2{F0: 1}
                    v1 = 1.5
                    if false {
                        v54 := GenS1{F0: 127, F1: -120, F2: v2}
                        if (v45 > v45) {
                            if (false || ((v45 - v45) >= (v45 - (v45 >> ((4294967295) & 63))))) {
                                v55 := GenS1{F0: 1, F1: 32767, F2: (v3 << ((65535) & 63))}
                                _ = v55
                            }
                        } else {
                            v56 := GenS2{F0: -117}
                            _ = v56
                        }
                        var v57 int = -491
                        v58 := GenS2{F0: 127}
                        var v59 int8 = -1
                        var v60 string = ""
                        v2 = v3
                        _ = v54
                        _ = v57
                        _ = v58
                        _ = v59
                        _ = v60
                    } else {
                        v2 -= v3
                    }
                    _ = fi52
                    _ = v53
                }
                _ = fi51
            }
            for fi61 := 0; fi61 < 18; fi61++ {
                v45 = (^(v45))
                var v62 uint8 = 232
                for fi63 := 0; fi63 < 3; fi63++ {
                    var v64 uint32 = 179
                    _ = fi63
                    _ = v64
                }
                _ = fi61
                _ = v62
            }
        }
    default:
        if ("a" >= "a") {
            if len(v4) > 0 {
                v4[0] = 2147483647
            }
        } else {
            var v65 int = -9
            _ = v65
        }
    }
    v66 := GenS1{F0: 42, F1: 42, F2: (^(v3))}
    v2 -= (v3 + 221)
    v67 := GenS1{F0: 127, F1: 494, F2: genFn0()}
    switch "a" {
    case "\n\t":
        v2 = genFn0()
    }
    switch "hello" {
    case "":
        if ((v3 % ((v2) | 1)) <= 189) {
            if (2147483647 > 9223372036854775807) {
                var v68 uint64 = 1
                var v69 float32 = 3.14159
                _ = v68
                _ = v69
            } else {
                v70 := GenS3{F0: (v2 ^ -9223372036854775808), F1: 18446744073709551615}
                v71 := GenS4{F0: genFn1(), F1: 42, F2: "foo bar"}
                if ((false && (true || true)) || false) {
                    switch "" {
                    case "hello":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "a":
                        for fi72 := 0; fi72 < 11; fi72++ {
                            for fi73 := 0; fi73 < 8; fi73++ {
                                v74 := GenS1{F0: 42, F1: 1000, F2: (^(v2))}
                                switch (^(fi72)) {
                                case 32768:
                                    v2 = v2
                                case 42:
                                    fmt.Println(fi73, v2, v3)
                                default:
                                    v45 = (v45 + v45)
                                }
                                v75 := GenS5{F0: true, F1: 1000}
                                _ = fi73
                                _ = v74
                                _ = v75
                            }
                            _ = fi72
                        }
                    case "":
                        if len(v4) > 0 {
                            v4[0] = -973
                        }
                    }
                    if ((42 >= 18446744073709551615) && true) {
                        v1 -= ((1.5 * v1) - v1)
                    }
                } else {
                    var v76 bool = (-460 == 42)
                    v4 = append(v4, 689)
                    fmt.Println(v3)
                    _ = v76
                }
                _ = v70
                _ = v71
            }
            if true {
                fmt.Println(v2, v2)
                if false {
                    if (float64((1.5 + (0.0 / 1.5))) <= -2.25) {
                        if true {
                            fmt.Println(v1)
                            fmt.Println(v45, v45, v1)
                        }
                        v77 := GenS6{F0: 145}
                        v1 = v1
                        _ = v77
                    } else {
                        switch -2147483648 {
                        case 42:
                            v4 = append(v4, 124)
                        case 9223372036854775807:
                            if (!true) {
                                var v78 float32 = float32((v1 / -2.25))
                                v78 = 3.14159
                                _ = v78
                            }
                        case 2147483647:
                            fmt.Println(v3, v45)
                        default:
                            switch "\n\t" {
                            case "a":
                                v2 = v2
                            case "hello":
                                _ = len(v4)
                            case "":
                                switch 9223372036854775807 {
                                case 32768:
                                    var v79 uint32 = 42
                                    _ = v79
                                case 128:
                                    var v80 uint = genFn1()
                                    _ = v80
                                case 0:
                                    var v81 float32 = 1e10
                                    _ = v81
                                default:
                                    fmt.Println(v3)
                                }
                            default:
                                fmt.Println(v1, v1, v2)
                            }
                        }
                        for fi82 := 0; fi82 < 13; fi82++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            fmt.Println(v2, v1)
                            _ = fi82
                        }
                        v83 := GenS3{F0: (v2 % ((genFn0()) | 1)), F1: 42}
                        _ = v83
                    }
                    for fi84 := 0; fi84 < 14; fi84++ {
                        var v85 bool = (!false)
                        for fi86 := 0; fi86 < 13; fi86++ {
                            _ = len(v4)
                            v4 = append(v4, (^(fi84)))
                            var v87 string = "a"
                            _ = fi86
                            _ = v87
                        }
                        _ = fi84
                        _ = v85
                    }
                    switch 42 {
                    case 1:
                        v3 = (v3 << ((0) & 63))
                    case 4294967295:
                        var v88 int32 = 1
                        _ = v88
                    }
                } else {
                    fmt.Println(v3, v45, v45)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch -172 {
                    case -2147483648:
                        v3 = 1
                    case 32768:
                        if false {
                            if (81 == 0) {
                                v89 := GenS7{F0: -794}
                                _ = v89
                            }
                        } else {
                            switch 0 {
                            case 32768:
                                for fi90 := 0; fi90 < 16; fi90++ {
                                    v2 -= (^(v2))
                                    _ = fi90
                                }
                            default:
                                var v91 int32 = -964
                                _ = v91
                            }
                            v1 -= float64(0.0)
                            var v92 int = 0
                            _ = v92
                        }
                    }
                }
                v3 = (v2 | (v3 - (v3 * v2)))
            } else {
                var v93 uint16 = (v45 << ((42) & 63))
                switch "\n\t" {
                case "hello":
                    v93 = (v45 + (v45 + 255))
                case "\n\t":
                    var v94 int64 = genFn0()
                    _ = v94
                default:
                    v95 := GenS8{F0: 0}
                    _ = v95
                }
                v96 := GenS4{F0: 210, F1: 2, F2: "hello"}
                _ = v93
                _ = v96
            }
        }
    case "hello":
        for fi97 := 0; fi97 < 16; fi97++ {
            for fi98 := 0; fi98 < 20; fi98++ {
                v2 -= 42
                if true {
                    if len(v4) > 0 {
                        v4[0] = (fi97 << ((65535) & 63))
                    }
                    var v99 float32 = ((-2.25 / float32(float64(-2.25))) + 1e10)
                    switch fi97 {
                    case 158:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v100 := GenS5{F0: (-47 <= 42), F1: 0}
                        switch -402 {
                        case 42:
                            var v101 bool = true
                            _ = v101
                        case 9223372036854775807:
                            var v102 int8 = 127
                            _ = v102
                        case -240:
                            fmt.Println(v99)
                        }
                        _ = v100
                    case 127:
                        v99 -= 0.0
                    case 237:
                        var v103 uint16 = 2
                        _ = v103
                    default:
                        v104 := GenS9{F0: 1, F1: v2}
                        _ = v104
                    }
                    var v105 uint32 = 2147483647
                    v106 := GenS10{F0: v105, F1: (!true), F2: 4294967295}
                    switch "hello" {
                    case "a":
                        v2 -= (-(v2))
                    case "hello":
                        _ = len(v4)
                    case "":
                        switch -568 {
                        case 32767:
                            if len(v4) > 0 {
                                v4[0] = (fi97 * fi98)
                            }
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    }
                    _ = v105
                    _ = v106
                    _ = v99
                } else {
                    switch (fi98 & 32768) {
                    case 0:
                        v1 -= v1
                    case 461:
                        var v107 bool = true
                        _ = v107
                    }
                }
                var v108 bool = (2147483647 == -35)
                var v109 uint32 = 65535
                fmt.Println(v109)
                _ = fi98
                _ = v108
                _ = v109
            }
            _ = fi97
        }
    }
    fmt.Println(v45)
    _ = len(v4)
    for fi110 := 0; fi110 < 19; fi110++ {
        _ = len(v4)
        var v111 uint64 = 1
        _ = fi110
        _ = v111
    }
    for fi112 := 0; fi112 < 11; fi112++ {
        var v113 uint16 = (^(v45))
        for fi114 := 0; fi114 < 4; fi114++ {
            var v115 int16 = -32768
            if true {
                v116 := GenS11{F0: (v2 << ((77) & 63))}
                v117 := GenS11{F0: (v3 - (^(v3)))}
                var v118 bool = (!true)
                _ = v116
                _ = v117
                _ = v118
            }
            var v119 uint64 = 1
            _ = fi114
            _ = v115
            _ = v119
        }
        v2 = (v2 * (v2 % (((-(v2))) | 1)))
        _ = fi112
        _ = v113
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v66
    _ = v67
}
