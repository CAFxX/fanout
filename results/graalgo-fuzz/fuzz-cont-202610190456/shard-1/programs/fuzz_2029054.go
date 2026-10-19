package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint; F1 int; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int32; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int16; F1 int16 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 int64; F1 int8 }

func main() {
    var v1 int64 = 42
    v2 := make(map[string]string)
    v3 := []string{"a", "", "\n\t"}
    v4 := make(map[uint32]bool)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint32 = 2
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: "", F1: -630, F2: 18446744073709551615}
    var v8 float32 = 0.0
    var v9 int32 = 2147483647
    fmt.Println(v1, v9, v8)
    v9 = (v9 & -2147483648)
    var v10 bool = (!true)
    v10 = (42 > 132)
    var v11 int32 = v9
    if false {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 bool = ((-(v9)) == v9)
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            v1 -= (v1 >> ((473) & 63))
            v10 = false
            _ = fi14
        }
        var v15 uint8 = 136
        _ = v15
    }
    if v10 {
        v1 = (v1 / (((v1 % ((71) | 1))) | 1))
        if false {
            var v16 int = -9223372036854775808
            _ = v16
        } else {
            var v17 uint16 = 1
            v18 := v2[("\n\t" + "")]
            _ = v17
            _ = v18
        }
    } else {
        var v19 int16 = 0
        if (!v10) {
            v20 := GenS2{F0: 492, F1: 460, F2: 42}
            if (float64(v8) != 1.5) {
                v10 = (v10 && (v8 <= v8))
                switch "hello" {
                case "\n\t":
                    for fi21 := 0; fi21 < 15; fi21++ {
                        switch "\n\t" {
                        case "hello":
                            _ = len(v4)
                        default:
                            var v22 float64 = (1.5 + 1e10)
                            _ = v22
                        }
                        var v23 int32 = -475
                        v24 := GenS3{F0: 1, F1: v9, F2: ("" + ("a" + "foo bar"))}
                        _ = fi21
                        _ = v23
                        _ = v24
                    }
                case "foo bar":
                    var v25 int32 = -1
                    _ = v25
                default:
                    v1 = v1
                }
                switch "hello" {
                case "\n\t":
                    var v26 int64 = (v1 - 42)
                    _ = v26
                default:
                    for fi27 := 0; fi27 < 11; fi27++ {
                        var v28 uint = 65535
                        _ = fi27
                        _ = v28
                    }
                }
            }
            _ = v20
        } else {
            for fi29 := 0; fi29 < 6; fi29++ {
                if ((v1 * 762) != v1) {
                    v30 := GenS3{F0: 42, F1: (v11 >> ((94) & 63)), F2: "\n\t"}
                    _ = v30
                } else {
                    var v31 int64 = v1
                    if (((((0.0 - -2.25) / -2.25) / 0.1) + ((-2.25 / 0.1) + -2.25)) >= (float64(v8) / ((float64((v8 * (v8 * (1.5 * 0.0)))) * -2.25) - (1e10 + 1e10)))) {
                        var v32 int16 = -1
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 9; fi33++ {
                            var v34 float64 = (((1e10 - 1e10) - 0.1) * (((1e10 * 0.0) / 1e10) - (float64((v8 * v8)) * float64(v8))))
                            v1 += v1
                            _ = fi33
                            _ = v34
                        }
                        v35 := GenS3{F0: 2, F1: v9, F2: ("foo bar" + "")}
                        switch int(v1) {
                        case 32767:
                            switch "a" {
                            case "foo bar":
                                switch "a" {
                                case "":
                                    for fi36 := 0; fi36 < 14; fi36++ {
                                        v1 = v1
                                        var v37 int16 = v19
                                        v8 -= float32(1e10)
                                        _ = fi36
                                        _ = v37
                                    }
                                case "a":
                                    v11 += (v11 << ((18446744073709551615) & 63))
                                }
                            case "hello":
                                v10 = false
                            default:
                                v9 -= v11
                            }
                        case 222:
                            v1 = (v31 ^ v1)
                        case 4294967295:
                            if v10 {
                                fmt.Println(v8)
                                var v38 int = (fi29 & (^(fi29)))
                                if false {
                                    v39 := GenS2{F0: 943, F1: v38, F2: 1}
                                    for fi40 := 0; fi40 < 4; fi40++ {
                                        var v41 uint64 = 719
                                        if (!v10) {
                                            var v42 float32 = (v8 * v8)
                                            var v43 uint16 = 65535
                                            _ = v42
                                            _ = v43
                                        } else {
                                            var v44 uint64 = 47
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if (v31 < v31) {
                                                v41 = (v41 / ((0) | 1))
                                                if false {
                                                    fmt.Println(v10)
                                                    v8 += (0.0 * v8)
                                                    fmt.Println(fi29, v38, v8)
                                                    if (4294967295 > -105) {
                                                        v45 := GenS4{F0: (v19 & v19)}
                                                        v11 += -1
                                                        switch "foo bar" {
                                                        case "a":
                                                            v46 := GenS1{F0: "hello", F1: v19, F2: v44}
                                                            _ = v46
                                                        case "\n\t":
                                                            v47 := GenS2{F0: 1, F1: (v38 * -664), F2: 2}
                                                            _ = v47
                                                        }
                                                        if (2 < 50) {
                                                            var v48 int64 = 9223372036854775807
                                                            var v49 uint64 = v44
                                                            var v50 float64 = 1e10
                                                            _ = v48
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        var v51 float32 = ((v8 + (v8 - v8)) + (float32(1e10) - (v8 / 1.0)))
                                                        _ = v45
                                                        _ = v51
                                                    } else {
                                                        v52 := GenS5{F0: 1, F1: v19}
                                                        var v53 uint32 = 1
                                                        fmt.Println(v8, v1, v44)
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                } else {
                                                    switch fi29 {
                                                    case -1:
                                                        v11 = int32(v41)
                                                    }
                                                    switch (("hello" + "hello") + "a") {
                                                    case "":
                                                        v54 := GenS3{F0: 99, F1: (v11 ^ 3), F2: "foo bar"}
                                                        _ = v54
                                                    case "hello":
                                                        var v55 bool = v10
                                                        _ = v55
                                                    case "a":
                                                        switch int(v1) {
                                                        case 4294967295:
                                                            switch ("\n\t" + (("foo bar" + (("hello" + "foo bar") + "")) + "hello")) {
                                                            case "foo bar":
                                                                v56 := GenS4{F0: -123}
                                                                _ = v56
                                                            case "hello":
                                                                switch fi29 {
                                                                case 233:
                                                                    fmt.Println(v9)
                                                                default:
                                                                    fmt.Println(v1, v9)
                                                                }
                                                            case "\n\t":
                                                                var v57 float64 = 1e10
                                                                _ = v57
                                                            }
                                                        case -953:
                                                            v58 := v2[("hello" + "")]
                                                            _ = v58
                                                        default:
                                                            var v59 int64 = 0
                                                            _ = v59
                                                        }
                                                    }
                                                }
                                            }
                                            _ = v44
                                        }
                                        var v60 uint16 = uint16(v41)
                                        _ = len(v4)
                                        var v61 uint16 = uint16(v41)
                                        var v62 uint32 = 2
                                        var v63 string = "foo bar"
                                        var v64 uint32 = v62
                                        _ = fi40
                                        _ = v41
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                    }
                                    _ = v39
                                } else {
                                    fmt.Println(v1, v10, v9)
                                    for fi65 := 0; fi65 < 17; fi65++ {
                                        v31 = v31
                                        if (v38 <= fi65) {
                                            if v10 {
                                                v9 = (v9 >> ((939) & 63))
                                                v9 = -834
                                            }
                                            v9 = v9
                                            switch -192 {
                                            case 9223372036854775807:
                                                switch "a" {
                                                case "\n\t":
                                                    fmt.Println(fi65)
                                                default:
                                                    if v10 {
                                                        if (2 > 2) {
                                                            v19 = v19
                                                        } else {
                                                            v66 := GenS3{F0: 156, F1: v9, F2: ("" + "hello")}
                                                            v19 += 32767
                                                            v67 := GenS1{F0: "a", F1: (^(v19)), F2: 544}
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        v4[0] = false
                                                        var v68 int16 = (-(v19))
                                                        _ = v68
                                                    } else {
                                                        var v69 int32 = (v11 & v9)
                                                        v9 = (v69 / ((2147483647) | 1))
                                                        _ = v69
                                                    }
                                                }
                                            case 127:
                                                v10 = v10
                                            }
                                        } else {
                                            if (!v10) {
                                                v19 = (v19 + (v19 << ((2) & 63)))
                                                v9 = v11
                                            }
                                            v19 = (v19 + v19)
                                            if true {
                                                if (v9 > v9) {
                                                    var v70 float64 = 1.5
                                                    if v10 {
                                                        v71 := GenS6{F0: 1, F1: 0}
                                                        v8 = ((v8 * (float32(float64(v8)) - float32(((v70 * -2.25) - 3.14159)))) * ((float32(float64((0.1 + v8))) / 1.0) + (float32(float64((1.5 / ((float32(v70) * v8) + v8)))) * 0.1)))
                                                        if ((1e10 + v70) != v70) {
                                                            switch ((((("foo bar" + "\n\t") + "a") + "hello") + (("hello" + ("foo bar" + "a")) + "a")) + "foo bar") {
                                                            case "hello":
                                                                var v72 uint64 = 2
                                                                _ = v72
                                                            case "":
                                                                v73 := GenS4{F0: (^(v19))}
                                                                _ = v73
                                                            case "a":
                                                                if v10 {
                                                                    fmt.Println(v38)
                                                                    var v74 float32 = 3.14159
                                                                    _ = v74
                                                                }
                                                            default:
                                                                var v75 uint32 = 42
                                                                _ = v75
                                                            }
                                                            fmt.Println(v8, v1)
                                                            v1 -= (v1 | v31)
                                                        }
                                                        _ = v71
                                                    }
                                                    var v76 int8 = 127
                                                    _ = v70
                                                    _ = v76
                                                } else {
                                                    switch fi29 {
                                                    case -2147483648:
                                                        var v77 float32 = v8
                                                        _ = v77
                                                    case 32767:
                                                        var v78 int8 = 127
                                                        _ = v78
                                                    case 128:
                                                        _ = len(v3)
                                                    default:
                                                        v31 += (v1 % ((v31) | 1))
                                                    }
                                                    if (v1 <= (v1 << ((761) & 63))) {
                                                        if (!false) {
                                                            fmt.Println(v10, v38, v1)
                                                            var v79 bool = (!true)
                                                            if false {
                                                                v11 += v9
                                                                if (3.14159 > 0.0) {
                                                                    v80 := GenS4{F0: v19}
                                                                    _ = len(v3)
                                                                    v1 = v1
                                                                    _ = v80
                                                                }
                                                            }
                                                            _ = v79
                                                        }
                                                        _ = len(v4)
                                                    } else {
                                                        if ((v19 - (v19 - 42)) >= v19) {
                                                            fmt.Println(v10)
                                                            var v81 int = -450
                                                            _ = v81
                                                        } else {
                                                            var v82 int64 = 929
                                                            if false {
                                                                v8 += v8
                                                                fmt.Println(v19)
                                                            } else {
                                                                v83 := GenS7{F0: -1, F1: 127}
                                                                if len(v3) > 0 {
                                                                    v3[0] = "\n\t"
                                                                }
                                                                _ = v83
                                                            }
                                                            v9 = v9
                                                            v84 := GenS6{F0: 42, F1: 927}
                                                            _ = v82
                                                            _ = v84
                                                        }
                                                        if (true && (((0.0 * 0.0) != v8) || v10)) {
                                                            v9 = (v11 << ((65535) & 63))
                                                            if ((0.0 * 0.1) == float64((0.0 + float32((-2.25 - (-2.25 * float64((((1e10 - 3.14159) - 3.14159) + float32(3.14159))))))))) {
                                                                fmt.Println(v19)
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                v11 -= v9
                                                            } else {
                                                                var v85 int64 = 863
                                                                v86 := GenS3{F0: 213, F1: (v9 | (v9 << ((2) & 63))), F2: "\n\t"}
                                                                v10 = v10
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                            var v87 uint = 65535
                                                            _ = v87
                                                        }
                                                        fmt.Println(fi29, fi65)
                                                    }
                                                }
                                                v88 := GenS4{F0: v19}
                                                _ = v88
                                            }
                                        }
                                        var v89 int64 = 1
                                        _ = fi65
                                        _ = v89
                                    }
                                    var v90 uint32 = 2147483647
                                    _ = v90
                                }
                                _ = v38
                            } else {
                                var v91 uint = 255
                                var v92 int8 = 0
                                _ = v91
                                _ = v92
                            }
                        }
                        _ = v35
                    }
                    _ = len(v4)
                    _ = v31
                }
                switch int(v1) {
                case 32768:
                    for fi93 := 0; fi93 < 20; fi93++ {
                        v3 = append(v3, "hello")
                        v11 = 42
                        var v94 int = 1
                        _ = fi93
                        _ = v94
                    }
                case -2147483648:
                    for fi95 := 0; fi95 < 3; fi95++ {
                        v10 = true
                        var v96 int = 9223372036854775807
                        _ = fi95
                        _ = v96
                    }
                case 0:
                    fmt.Println(v1)
                default:
                    for fi97 := 0; fi97 < 4; fi97++ {
                        switch -1 {
                        case 926:
                            v1 = v1
                        case 32768:
                            switch 1 {
                            case 2147483647:
                                if v10 {
                                    if (!false) {
                                        if len(v3) > 0 {
                                            v3[0] = ("\n\t" + "foo bar")
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = ("foo bar" + ("" + ("" + ("" + ("a" + "\n\t")))))
                                        }
                                        var v98 int32 = (-(v11))
                                        _ = v98
                                    } else {
                                        v99 := GenS7{F0: v1, F1: 85}
                                        _ = v99
                                    }
                                    switch ("" + ("hello" + "hello")) {
                                    case "":
                                        var v100 bool = (false && v10)
                                        _ = v100
                                    case "a":
                                        _ = len(v2)
                                    case "foo bar":
                                        v2[""] = ""
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                    }
                                    v19 = -1
                                    var v101 uint64 = 399
                                    _ = len(v3)
                                    v19 -= v19
                                    _ = v101
                                }
                            case -463:
                                var v102 uint8 = 2
                                _ = v102
                            }
                        case -23:
                            if ("" > ("foo bar" + "a")) {
                                var v103 int16 = v19
                                _ = v103
                            } else {
                                var v104 uint32 = 192
                                var v105 float32 = (-2.25 * v8)
                                if true {
                                    var v106 uint16 = 2
                                    var v107 float32 = v105
                                    switch "\n\t" {
                                    case "a":
                                        var v108 int16 = v19
                                        _ = v108
                                    case "\n\t":
                                        v105 = float32(-2.25)
                                    case "hello":
                                        fmt.Println(fi29)
                                        v3 = append(v3, "foo bar")
                                        v4[(v104 << ((18446744073709551615) & 63))] = v10
                                    }
                                    _ = v106
                                    _ = v107
                                } else {
                                    var v109 uint32 = v104
                                    var v110 uint32 = 2147483648
                                    v104 = v104
                                    _ = v109
                                    _ = v110
                                }
                                _ = v104
                                _ = v105
                            }
                        default:
                            var v111 float32 = (((0.1 / float32(-2.25)) - (float32(3.14159) + -2.25)) + (3.14159 + v8))
                            _ = v111
                        }
                        var v112 uint8 = 8
                        _ = fi97
                        _ = v112
                    }
                }
                _ = fi29
            }
            v113 := GenS4{F0: (v19 | v19)}
            _ = v113
        }
        v2["a"] = "hello"
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
