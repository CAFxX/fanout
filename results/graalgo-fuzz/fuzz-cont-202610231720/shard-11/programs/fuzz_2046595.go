package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 int64; F2 int }
type GenS4 struct { F0 int; F1 string; F2 uint32 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{4294967295}
    var v3 float32 = 0.1
    var v4 bool = (!(-356 != 9223372036854775807))
    var v5 int = -695
    v6 := []uint32{}
    if (v5 < v5) {
        v7 := GenS1{F0: -49, F1: 32767, F2: ("" + ("foo bar" + ("hello" + "")))}
        _ = v7
    }
    fmt.Println(v5, v3)
    if false {
        if (1 <= 42) {
            var v8 string = ((("foo bar" + "hello") + "a") + ("hello" + "foo bar"))
            v9 := v1[4294967295]
            _ = v8
            _ = v9
        } else {
            if ((!v4) || ((0.1 / 3.14159) >= ((((3.14159 / (0.1 - 0.0)) + 1.5) - 0.0) + 3.14159))) {
                var v10 uint64 = 982
                var v11 int64 = -771
                for fi12 := 0; fi12 < 8; fi12++ {
                    if v4 {
                        if v4 {
                            var v13 uint = 0
                            fmt.Println(v4)
                            v14 := v1[4294967295]
                            _ = v13
                            _ = v14
                        }
                        fmt.Println(v10)
                        if v4 {
                            fmt.Println(v4, v11, v3)
                            v6 = append(v6, 653)
                            var v15 int8 = 1
                            _ = v15
                        } else {
                            _ = len(v2)
                            for fi16 := 0; fi16 < 15; fi16++ {
                                v4 = true
                                if v4 {
                                    var v17 int = (fi16 - (-(fi12)))
                                    v11 = (v11 << ((169) & 63))
                                    fmt.Println(fi12, v17)
                                    var v18 float32 = float32(1e10)
                                    fmt.Println(v5, v10, fi12)
                                    var v19 int16 = 0
                                    var v20 float64 = 0.1
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                } else {
                                    var v21 uint32 = 2147483648
                                    _ = v21
                                }
                                _ = fi16
                            }
                            var v22 float64 = 1.5
                            _ = v22
                        }
                    }
                    _ = fi12
                }
                _ = v10
                _ = v11
            }
            var v23 int16 = 0
            fmt.Println(v23)
            _ = v23
        }
        switch (("hello" + ("\n\t" + "")) + "") {
        case "hello":
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 uint16 = 0
                _ = fi24
                _ = v25
            }
        case "a":
            var v26 int64 = -9223372036854775808
            _ = v26
        case "":
            for fi27 := 0; fi27 < 12; fi27++ {
                var v28 int64 = 1
                v29 := GenS2{F0: 980}
                _ = fi27
                _ = v28
                _ = v29
            }
        default:
            var v30 string = (("hello" + (("" + "") + "hello")) + "foo bar")
            _ = v30
        }
        if v4 {
            v31 := GenS1{F0: -1, F1: -742, F2: "foo bar"}
            _ = v31
        }
    }
    switch (v5 << ((255) & 63)) {
    case 1:
        for fi32 := 0; fi32 < 8; fi32++ {
            switch fi32 {
            case 33:
                var v33 bool = (127 < 2)
                _ = v33
            case 4294967295:
                v3 = (1.5 + v3)
            default:
                for fi34 := 0; fi34 < 17; fi34++ {
                    v5 = (fi34 >> ((18446744073709551615) & 63))
                    v4 = (1 < -573)
                    var v35 uint64 = 42
                    _ = len(v6)
                    fmt.Println(fi32)
                    var v36 string = "hello"
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
            v37 := GenS2{F0: 2}
            if v4 {
                v38 := GenS3{F0: 1, F1: -9223372036854775808, F2: (^(fi32))}
                _ = v38
            }
            _ = fi32
            _ = v37
        }
    case 262:
        if v4 {
            v39 := GenS2{F0: 97}
            if (4294967295 != 827) {
                fmt.Println(v5, v3, v3)
                var v40 int64 = 135
                _ = v40
            }
            var v41 uint = 428
            _ = v39
            _ = v41
        } else {
            v1[0] = (((("foo bar" + "") + ("\n\t" + "foo bar")) + "foo bar") + ("\n\t" + "foo bar"))
            for fi42 := 0; fi42 < 6; fi42++ {
                if (false || (!v4)) {
                    v4 = v4
                }
                var v43 int = 0
                _ = fi42
                _ = v43
            }
            switch (-(v5)) {
            case -362:
                switch ("" + "foo bar") {
                case "a":
                    v44 := GenS2{F0: 42}
                    _ = v44
                }
            case -578:
                var v45 int8 = -128
                _ = v45
            default:
                v3 = v3
            }
        }
    case 2147483647:
        var v46 uint = 386
        _ = v46
    }
    if v4 {
        if (2 >= 4294967295) {
            v5 -= v5
            var v47 float64 = 3.14159
            v47 = v47
            _ = v47
        } else {
            delete(v1, 771)
            v48 := GenS1{F0: 1, F1: 42, F2: ""}
            _ = v48
        }
    }
    switch (v5 >> ((4294967295) & 63)) {
    case -631:
        if v4 {
            var v49 uint16 = 264
            v5 += (^(v5))
            v50 := v1[4294967295]
            _ = v49
            _ = v50
        }
    case 394:
        if (2 <= 18446744073709551615) {
            v5 += 0
        } else {
            if (v4 && ((42 > 0) && false)) {
                var v51 int = (v5 & 127)
                _ = v51
            } else {
                _ = len(v6)
                for fi52 := 0; fi52 < 19; fi52++ {
                    if (-593 <= 42) {
                        v53 := GenS4{F0: fi52, F1: "a", F2: 2147483648}
                        if ("" != "") {
                            for fi54 := 0; fi54 < 7; fi54++ {
                                var v55 uint16 = 255
                                _ = fi54
                                _ = v55
                            }
                        }
                        if false {
                            if (true || ((1 <= -648) && (154 > 902))) {
                                switch "" {
                                case "a":
                                    for fi56 := 0; fi56 < 8; fi56++ {
                                        switch "" {
                                        case "\n\t":
                                            fmt.Println(fi52, fi56)
                                        default:
                                            v3 -= (v3 - v3)
                                        }
                                        _ = fi56
                                    }
                                default:
                                    v57 := GenS4{F0: (v5 ^ (fi52 >> ((1) & 63))), F1: ("foo bar" + "\n\t"), F2: 1}
                                    _ = v57
                                }
                                for fi58 := 0; fi58 < 20; fi58++ {
                                    if false {
                                        v6 = append(v6, 2)
                                        var v59 int = 655
                                        _ = v59
                                    } else {
                                        v60 := GenS4{F0: (^(fi52)), F1: ("\n\t" + ("" + "a")), F2: 2147483647}
                                        var v61 int64 = 42
                                        var v62 int = (v5 / ((fi52) | 1))
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                    }
                                    switch -795 {
                                    case 750:
                                        v63 := GenS5{F0: v4, F1: 1}
                                        _ = v63
                                    case -270:
                                        v4 = v4
                                    default:
                                        v4 = (2 <= 18446744073709551615)
                                    }
                                    var v64 int8 = 127
                                    switch "" {
                                    case "a":
                                        fmt.Println(v4, v5, v4)
                                    case "":
                                        _ = len(v1)
                                    default:
                                        v65 := GenS6{F0: 802, F1: -990}
                                        _ = v65
                                    }
                                    if true {
                                        v4 = true
                                        var v66 int64 = 1
                                        v64 += (v64 & v64)
                                        _ = v66
                                    }
                                    _ = fi58
                                    _ = v64
                                }
                            }
                        }
                        _ = v53
                    } else {
                        fmt.Println(v4)
                        v4 = v4
                        v67 := GenS3{F0: 85, F1: -468, F2: 1}
                        _ = v67
                    }
                    _ = fi52
                }
            }
            v4 = v4
        }
    case -968:
        var v68 int8 = 100
        _ = v68
    }
    switch "a" {
    case "":
        var v69 int32 = 88
        _ = v69
    case "foo bar":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    for fi70 := 0; fi70 < 2; fi70++ {
        var v71 float64 = -2.25
        var v72 uint32 = 2147483647
        if v4 {
            v72 -= v72
            var v73 uint8 = 1
            _ = v73
        }
        _ = fi70
        _ = v71
        _ = v72
    }
    var v74 bool = ((!v4) && true)
    if v4 {
        var v75 uint = 228
        for fi76 := 0; fi76 < 11; fi76++ {
            var v77 uint = (v75 ^ (v75 % ((42) | 1)))
            fmt.Println(v4)
            _ = fi76
            _ = v77
        }
        var v78 int8 = 42
        _ = v75
        _ = v78
    } else {
        v5 = v5
        if (42 < 100) {
            v3 -= v3
            for fi79 := 0; fi79 < 4; fi79++ {
                for fi80 := 0; fi80 < 14; fi80++ {
                    var v81 uint64 = 29
                    _ = fi80
                    _ = v81
                }
                var v82 uint32 = 643
                var v83 string = ""
                _ = fi79
                _ = v82
                _ = v83
            }
        } else {
            for fi84 := 0; fi84 < 8; fi84++ {
                var v85 float64 = -2.25
                var v86 int = (-(fi84))
                v4 = false
                _ = fi84
                _ = v85
                _ = v86
            }
            v74 = v4
            var v87 uint16 = 2
            _ = v87
        }
        v5 = 127
    }
    if (v74 || (2 <= 763)) {
        switch "hello" {
        case "":
            if v74 {
                if v4 {
                    v88 := GenS7{F0: -96}
                    _ = v88
                }
                fmt.Println(v74)
            }
        default:
            for fi89 := 0; fi89 < 8; fi89++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v3 = float32(0.0)
                switch fi89 {
                case -708:
                    v74 = v4
                }
                _ = fi89
            }
        }
        for fi90 := 0; fi90 < 6; fi90++ {
            var v91 int64 = 1
            var v92 int16 = 42
            _ = fi90
            _ = v91
            _ = v92
        }
        v3 = float32(0.1)
    } else {
        delete(v1, 119)
    }
    if len(v2) > 0 {
        v2[0] = v5
    }
    var v93 uint = 0
    fmt.Println(v4, v74)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v94 uint = v93
    var v95 int64 = 0
    for fi96 := 0; fi96 < 1; fi96++ {
        for fi97 := 0; fi97 < 3; fi97++ {
            fmt.Println(v93, fi96, v74)
            fmt.Println(v3, v93, fi97)
            _ = fi97
        }
        _ = fi96
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v74:", v74)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v93
    _ = v94
    _ = v95
}
