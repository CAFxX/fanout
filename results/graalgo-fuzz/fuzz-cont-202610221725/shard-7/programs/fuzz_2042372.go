package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 uint }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 int64; F1 int8; F2 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 uint; F1 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"foo bar", "foo bar"}
    var v3 int32 = -1
    v4 := []int{2147483647, 9223372036854775807, -9223372036854775808, 4294967295}
    fmt.Println(v3)
    _ = len(v2)
    for fi5 := 0; fi5 < 18; fi5++ {
        switch ("foo bar" + "foo bar") {
        case "hello":
            switch fi5 {
            case -708:
                v3 += v3
            case 127:
                if (0.0 != -2.25) {
                    var v6 bool = false
                    for fi7 := 0; fi7 < 4; fi7++ {
                        v8 := GenS1{F0: 2, F1: (v3 >= (v3 - v3)), F2: 4294967295}
                        var v9 uint = 2
                        if len(v4) > 0 {
                            v4[0] = 128
                        }
                        _ = fi7
                        _ = v8
                        _ = v9
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v6
                }
            case 0:
                if true {
                    var v10 uint32 = 42
                    _ = v10
                } else {
                    v3 += v3
                    for fi11 := 0; fi11 < 16; fi11++ {
                        if false {
                            fmt.Println(v3, fi5)
                        } else {
                            v3 += (v3 << ((2) & 63))
                        }
                        v3 = v3
                        fmt.Println(fi5)
                        _ = fi11
                    }
                }
            }
        default:
            switch (("\n\t" + ("foo bar" + "\n\t")) + "a") {
            case "foo bar":
                v4 = append(v4, fi5)
            }
        }
        _ = fi5
    }
    if false {
        delete(v1, 4294967295)
        v3 -= (^(v3))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        switch ("" + ("hello" + "")) {
        case "\n\t":
            v3 = 0
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v12 int64 = 42
    var v13 int8 = -109
    v3 -= (v3 + v3)
    if (42 == 1000) {
        v13 = (v13 + (v13 ^ v13))
    }
    fmt.Println(v13)
    if (false || true) {
        v13 = -121
        v12 = (-(v12))
    } else {
        switch ("hello" + ("foo bar" + "\n\t")) {
        case "\n\t":
            if (!false) {
                v14 := GenS2{F0: 400, F1: 1000}
                v1[2] = false
                v15 := GenS2{F0: -268, F1: -32768}
                _ = v14
                _ = v15
            } else {
                v16 := GenS3{F0: (v12 ^ (^(v12))), F1: -1, F2: v12}
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 int64 = v12
                    if len(v2) > 0 {
                        v2[0] = ("" + "foo bar")
                    }
                    _ = fi17
                    _ = v18
                }
                v3 = v3
                _ = v16
            }
        default:
            if (v12 <= (v12 + (v12 + (v12 / (((v12 & 9223372036854775807)) | 1))))) {
                switch ("a" + "foo bar") {
                case "a":
                    fmt.Println(v3, v3, v12)
                case "":
                    switch "" {
                    case "foo bar":
                        v13 = v13
                    case "hello":
                        v19 := GenS2{F0: 42, F1: -1}
                        _ = v19
                    }
                }
                v3 -= v3
                for fi20 := 0; fi20 < 14; fi20++ {
                    if true {
                        v21 := GenS4{F0: 393}
                        for fi22 := 0; fi22 < 9; fi22++ {
                            v23 := GenS3{F0: v12, F1: v13, F2: (v12 << ((2) & 63))}
                            delete(v1, 65535)
                            _ = fi22
                            _ = v23
                        }
                        _ = v21
                    }
                    v24 := GenS5{F0: -1}
                    _ = fi20
                    _ = v24
                }
            } else {
                if ((v12 / ((v12) | 1)) <= (v12 / ((1) | 1))) {
                    var v25 uint8 = 1
                    v26 := GenS3{F0: v12, F1: 41, F2: (^(v12))}
                    v27 := v1[871]
                    _ = v25
                    _ = v26
                    _ = v27
                }
                var v28 uint16 = 2
                if false {
                    if (((3.14159 + 1e10) / (3.14159 + -2.25)) > (0.1 - 1.5)) {
                        var v29 int32 = 1
                        if (true && ("hello" > "")) {
                            var v30 string = ""
                            _ = v30
                        }
                        v31 := GenS3{F0: (v12 % ((-9223372036854775808) | 1)), F1: 2, F2: (v12 >> ((80) & 63))}
                        _ = v29
                        _ = v31
                    } else {
                        if true {
                            switch 32768 {
                            case 1:
                                var v32 int16 = -1
                                _ = v32
                            default:
                                if (!(int(v12) >= 1)) {
                                    v33 := GenS4{F0: 18446744073709551615}
                                    v12 = (v12 << ((2) & 63))
                                    if false {
                                        v1[2147483647] = (false || true)
                                    }
                                    _ = v33
                                } else {
                                    v34 := GenS6{F0: 440, F1: (("\n\t" + "a") + "")}
                                    var v35 uint8 = 205
                                    _ = v34
                                    _ = v35
                                }
                            }
                        }
                        v3 -= v3
                        if false {
                            var v36 uint64 = 890
                            var v37 int64 = v12
                            if false {
                                var v38 int32 = (v3 >> ((uint(v28)) & 63))
                                if true {
                                    for fi39 := 0; fi39 < 19; fi39++ {
                                        if (!true) {
                                            v12 += (v12 >> ((65535) & 63))
                                        }
                                        _ = fi39
                                    }
                                    fmt.Println(v38, v3, v38)
                                    var v40 int64 = v12
                                    _ = v40
                                } else {
                                    fmt.Println(v38, v28, v38)
                                    for fi41 := 0; fi41 < 9; fi41++ {
                                        if (((^(v28)) > v28) && ((("\n\t" + ("foo bar" + "hello")) + "") < "hello")) {
                                            for fi42 := 0; fi42 < 6; fi42++ {
                                                v43 := GenS3{F0: 9223372036854775807, F1: -31, F2: v12}
                                                var v44 uint32 = 0
                                                v36 -= 4294967295
                                                _ = fi42
                                                _ = v43
                                                _ = v44
                                            }
                                            for fi45 := 0; fi45 < 20; fi45++ {
                                                var v46 int16 = 32767
                                                if false {
                                                    v13 -= v13
                                                } else {
                                                    var v47 uint32 = 4294967295
                                                    v37 = 9223372036854775807
                                                    _ = v47
                                                }
                                                _ = fi45
                                                _ = v46
                                            }
                                            for fi48 := 0; fi48 < 15; fi48++ {
                                                if (v28 > v28) {
                                                    var v49 uint = 1
                                                    _ = v49
                                                }
                                                v50 := v1[643]
                                                _ = fi48
                                                _ = v50
                                            }
                                        } else {
                                            var v51 float64 = 0.1
                                            var v52 uint = 586
                                            var v53 uint64 = uint64(v28)
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        }
                                        for fi54 := 0; fi54 < 4; fi54++ {
                                            fmt.Println(v38, v36, v13)
                                            _ = fi54
                                        }
                                        _ = fi41
                                    }
                                    var v55 int64 = 370
                                    _ = v55
                                }
                                switch -993 {
                                case 9223372036854775807:
                                    var v56 int64 = v12
                                    _ = v56
                                case 127:
                                    v38 = v3
                                case 32768:
                                    v12 += -9223372036854775808
                                default:
                                    for fi57 := 0; fi57 < 8; fi57++ {
                                        var v58 uint = 761
                                        var v59 uint16 = v28
                                        if true {
                                            v28 = v28
                                            for fi60 := 0; fi60 < 1; fi60++ {
                                                v61 := GenS6{F0: 4294967295, F1: "foo bar"}
                                                if (v36 <= v36) {
                                                    v36 += v36
                                                    switch fi57 {
                                                    case -859:
                                                        fmt.Println(fi57, v59, v12)
                                                    }
                                                }
                                                _ = fi60
                                                _ = v61
                                            }
                                        }
                                        _ = fi57
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                _ = v38
                            } else {
                                if true {
                                    for fi62 := 0; fi62 < 16; fi62++ {
                                        var v63 int8 = v13
                                        _ = fi62
                                        _ = v63
                                    }
                                }
                                var v64 float64 = (3.14159 / ((float64(1e10) / 1.5) - -2.25))
                                _ = v64
                            }
                            _ = v36
                            _ = v37
                        }
                    }
                    v12 += v12
                    if (("\n\t" > ("\n\t" + "\n\t")) || false) {
                        for fi65 := 0; fi65 < 4; fi65++ {
                            var v66 float64 = (1e10 + 3.14159)
                            for fi67 := 0; fi67 < 18; fi67++ {
                                fmt.Println(fi65, v12)
                                v66 -= v66
                                fmt.Println(fi67, v66)
                                _ = fi67
                            }
                            for fi68 := 0; fi68 < 9; fi68++ {
                                switch (-(fi65)) {
                                case 0:
                                    v69 := GenS7{F0: 0, F1: 0}
                                    _ = v69
                                case 9223372036854775807:
                                    var v70 int = (-(fi68))
                                    _ = v70
                                default:
                                    if (fi65 > fi68) {
                                        v13 = v13
                                        var v71 uint64 = 18446744073709551615
                                        _ = v71
                                    }
                                }
                                _ = len(v2)
                                var v72 uint64 = 42
                                _ = fi68
                                _ = v72
                            }
                            _ = fi65
                            _ = v66
                        }
                    }
                }
                _ = v28
            }
        }
    }
    if true {
        if (-2147483648 >= -9223372036854775808) {
            if (false || true) {
                v13 -= -1
                v12 += (v12 / (((v12 << ((4294967295) & 63))) | 1))
                var v73 bool = (true || false)
                _ = v73
            } else {
                if len(v4) > 0 {
                    v4[0] = 32767
                }
                if (!((false && (!true)) || false)) {
                    var v74 uint32 = 100
                    var v75 float64 = 3.14159
                    var v76 uint8 = 42
                    _ = v74
                    _ = v75
                    _ = v76
                }
                v77 := GenS7{F0: 593, F1: (-(v3))}
                _ = v77
            }
        } else {
            v12 += v12
        }
    }
    v13 += v13
    var v78 int8 = v13
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v78
}
