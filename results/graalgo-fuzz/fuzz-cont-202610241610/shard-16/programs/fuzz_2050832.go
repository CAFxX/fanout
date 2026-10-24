package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 string; F1 int16; F2 bool }
type GenS3 struct { F0 int32; F1 bool; F2 int64 }
type GenS4 struct { F0 uint; F1 int; F2 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int8 }
func genFn0(p0 int, p1 uint8, p2 int) uint16 {
    if false {
        for fi7 := 0; fi7 < 1; fi7++ {
            if false {
                var v8 float64 = 0.0
                var v9 uint64 = uint64(p1)
                var v10 int16 = 42
                _ = v10
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        if (!(true || true)) {
            for fi11 := 0; fi11 < 10; fi11++ {
                var v12 uint64 = 4294967295
                for fi13 := 0; fi13 < 5; fi13++ {
                    var v14 uint64 = v12
                    var v15 int64 = 42
                    if true {
                        v16 := []int{}
                        p2 -= fi11
                        _ = v16
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = fi11
                _ = v12
            }
        } else {
            v17 := GenS1{F0: 42, F1: 1}
            var v18 int64 = -982
            _ = v17
            _ = v18
        }
    }
    if (118 < 364) {
        var v19 int = (p0 ^ (p2 | p2))
        if false {
            fmt.Println(v19, p1, p2)
            var v20 uint32 = uint32(p1)
            var v21 int = -704
            _ = v20
            _ = v21
        } else {
            if (false || false) {
                var v22 float32 = -2.25
                var v23 string = ""
                v23 = v23
                _ = v22
                _ = v23
            } else {
                if true {
                    if (true && (!((p2 <= (p2 / ((32768) | 1))) && (!false)))) {
                        v19 = (v19 & p2)
                        p0 -= (p2 >> ((1) & 63))
                    } else {
                        fmt.Println(v19, p1)
                        v24 := []int{int(p1), 9223372036854775807, (v19 ^ -468)}
                        var v25 int64 = -1
                        _ = v24
                        _ = v25
                    }
                    var v26 int16 = 1000
                    _ = v26
                } else {
                    if (903 <= 42) {
                        v27 := GenS1{F0: 2147483648, F1: p2}
                        v28 := []int{(v19 << ((0) & 63)), -1}
                        _ = v27
                        _ = v28
                    } else {
                        if true {
                            v19 = (v19 + p0)
                        }
                        for fi29 := 0; fi29 < 16; fi29++ {
                            for fi30 := 0; fi30 < 18; fi30++ {
                                if true {
                                    v31 := GenS2{F0: (("a" + ("a" + "hello")) + ""), F1: 165, F2: (292 != 0)}
                                    var v32 uint8 = 127
                                    _ = v31
                                    _ = v32
                                }
                                var v33 uint16 = 122
                                _ = fi30
                                _ = v33
                            }
                            _ = fi29
                        }
                    }
                    p2 = (-(p0))
                    for fi34 := 0; fi34 < 3; fi34++ {
                        var v35 int = fi34
                        var v36 uint32 = 838
                        if ((1 < 682) || true) {
                            var v37 uint64 = 2
                            v38 := []int{p0, 1, fi34}
                            _ = v37
                            _ = v38
                        } else {
                            if ((-2.25 + 3.14159) < 1.5) {
                                v39 := []int{(v19 * (v35 >> ((0) & 63))), -1, p2, v19}
                                _ = v39
                            }
                            var v40 uint8 = (p1 / ((p1) | 1))
                            p2 = (v19 | -2147483648)
                            _ = v40
                        }
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
            }
        }
        _ = v19
    }
    switch ("foo bar" + (("hello" + "a") + "\n\t")) {
    case "hello":
        if true {
            switch -209 {
            case 32767:
                v41 := []int{(p0 - p2), (p2 & (p0 * (p2 % ((p2) | 1))))}
                _ = v41
            }
            if (false || (true && (0 >= 355))) {
                var v42 uint8 = (p1 / (((p1 % (((p1 << ((719) & 63))) | 1))) | 1))
                v43 := []int{(p0 | p0), (^(p0))}
                if (((!(!false)) && false) && false) {
                    v43 = append(v43, 128)
                    v42 += (p1 & v42)
                    switch 32767 {
                    case -9223372036854775808:
                        v42 = (p1 * 255)
                    case -640:
                        var v44 float32 = ((1e10 * float32((1e10 - 1e10))) - 0.1)
                        _ = v44
                    case 0:
                        for fi45 := 0; fi45 < 18; fi45++ {
                            if false {
                                fmt.Println(p1, fi45, fi45)
                            } else {
                                if len(v43) >= 2 {
                                    v43 = v43[1:len(v43)]
                                }
                            }
                            var v46 uint = 42
                            _ = fi45
                            _ = v46
                        }
                    default:
                        v42 -= (p1 << ((1) & 63))
                    }
                }
                _ = v42
                _ = v43
            }
            v47 := GenS2{F0: "a", F1: 0, F2: false}
            _ = v47
        } else {
            p1 += (p1 & 214)
        }
    }
    v48 := GenS2{F0: ("a" + "\n\t"), F1: -424, F2: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v48
    return 2
}
func genFn1(p0 int16) uint {
    v49 := GenS2{F0: (("foo bar" + "foo bar") + "hello"), F1: (p0 | p0), F2: (!(0 <= 255))}
    fmt.Println(p0)
    _ = p0
    _ = v49
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -2147483648
    var v3 uint8 = 1
    var v4 uint32 = 42
    var v5 uint = 255
    v6 := []string{"foo bar", (("a" + "foo bar") + "a"), ("" + "\n\t"), "foo bar"}
    if (3.14159 >= float32(3.14159)) {
        switch "a" {
        case "":
            var v50 string = ("hello" + "")
            _ = v50
        default:
            var v51 uint = (^(v5))
            _ = v51
        }
        v3 -= (v3 - (v3 % ((v3) | 1)))
    } else {
        var v52 uint8 = 38
        switch 947 {
        case 856:
            fmt.Println(v3)
        default:
            if (100 < 42) {
                v5 -= v5
                _ = len(v1)
            } else {
                v2 += 1
            }
        }
        fmt.Println(v52, v2)
        _ = v52
    }
    switch "foo bar" {
    case "hello":
        var v53 uint16 = 255
        _ = v53
    }
    for fi54 := 0; fi54 < 6; fi54++ {
        for fi55 := 0; fi55 < 8; fi55++ {
            if false {
                if len(v6) > 0 {
                    v6[0] = (("hello" + "foo bar") + "a")
                }
            } else {
                v56 := GenS3{F0: v2, F1: true, F2: 9223372036854775807}
                _ = v56
            }
            if true {
                if ((v5 / (((v5 * genFn1(1000))) | 1)) != v5) {
                    switch fi54 {
                    case -1:
                        v3 = 127
                    case 32768:
                        var v57 uint = v5
                        _ = v57
                    case 1:
                        v3 = (v3 - (^(v3)))
                    default:
                        if true {
                            fmt.Println(fi55, fi54)
                            if true {
                                v2 += v2
                                v58 := v1["hello"]
                                delete(v1, ("hello" + v58))
                                _ = v58
                            } else {
                                switch "" {
                                case "a":
                                    var v59 string = ""
                                    _ = v59
                                case "":
                                    var v60 uint32 = (^(v4))
                                    v5 = (v5 << (((^(v5))) & 63))
                                    _ = v60
                                default:
                                    var v61 int = fi54
                                    _ = v61
                                }
                                if (("\n\t" + "hello") <= "\n\t") {
                                    v62 := GenS4{F0: 637, F1: (fi55 >> (((v5 / (((v5 / ((2) | 1))) | 1))) & 63)), F2: v5}
                                    _ = v62
                                } else {
                                    if ((1000 != 99) || false) {
                                        switch (-(fi54)) {
                                        case 32767:
                                            switch "hello" {
                                            case "hello":
                                                var v63 uint32 = (^(v4))
                                                fmt.Println(v5, fi55, v3)
                                                v5 += (v5 & v5)
                                                _ = v63
                                            default:
                                                v5 += (^(v5))
                                            }
                                        }
                                    } else {
                                        v3 -= 0
                                        var v64 uint32 = uint32(v3)
                                        v3 -= 2
                                        var v65 uint64 = 0
                                        var v66 uint = v5
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                    }
                                    var v67 int8 = -17
                                    _ = v67
                                }
                            }
                            v3 = 156
                        }
                    }
                    v4 = (v4 % ((v4) | 1))
                    var v68 float64 = 1.5
                    _ = v68
                }
                v5 += v5
            } else {
                var v69 uint16 = 1
                v5 += (v5 << ((uint(v4)) & 63))
                _ = v69
            }
            if true {
                switch "a" {
                case "":
                    var v70 int = -9223372036854775808
                    _ = v70
                case "hello":
                    if (!(false || false)) {
                        v1["hello"] = ("" + "")
                        v71 := GenS1{F0: 2, F1: fi54}
                        _ = v71
                    } else {
                        v3 = 128
                    }
                case "a":
                    v2 -= v2
                }
                switch fi55 {
                case 320:
                    if true {
                        v5 -= (^(v5))
                    }
                case 1:
                    var v72 int32 = 63
                    _ = v72
                case 9223372036854775807:
                    var v73 uint8 = v3
                    _ = v73
                }
                var v74 int64 = 0
                _ = v74
            }
            _ = fi55
        }
        v75 := GenS5{F0: 901}
        if false {
            v5 += v5
        }
        _ = fi54
        _ = v75
    }
    v76 := GenS2{F0: "foo bar", F1: 1, F2: true}
    var v77 int = 322
    fmt.Println(v3, v4, v77)
    v78 := GenS5{F0: 112}
    _ = len(v6)
    fmt.Println(v4, v77, v3)
    v3 = 1
    _ = len(v1)
    if true {
        var v79 int64 = 9223372036854775807
        if true {
            v3 += (v3 / ((v3) | 1))
            fmt.Println(v77, v2)
            var v80 int8 = 42
            _ = v80
        } else {
            v81 := GenS3{F0: v2, F1: true, F2: v79}
            if (1 >= 42) {
                fmt.Println(v77, v79, v2)
                v5 -= genFn1(32767)
            }
            _ = v81
        }
        switch (v77 << ((v5) & 63)) {
        case 719:
            v4 += v4
        case -9223372036854775808:
            switch "hello" {
            case "":
                for fi82 := 0; fi82 < 9; fi82++ {
                    switch (-(fi82)) {
                    case 32768:
                        v79 = v79
                    case 42:
                        v77 = v77
                    case 29:
                        if ((-32768 <= -1) && ((fi82 % ((v77) | 1)) == (v77 & (fi82 & -2147483648)))) {
                            _ = len(v6)
                            if true {
                                _ = len(v1)
                            } else {
                                var v83 string = ("foo bar" + "")
                                _ = len(v1)
                                for fi84 := 0; fi84 < 16; fi84++ {
                                    v85 := GenS2{F0: "\n\t", F1: 1000, F2: (v3 <= v3)}
                                    _ = fi84
                                    _ = v85
                                }
                                _ = v83
                            }
                            for fi86 := 0; fi86 < 11; fi86++ {
                                var v87 float64 = (float64(float32(-2.25)) + float64(0.1))
                                _ = fi86
                                _ = v87
                            }
                        }
                    default:
                        for fi88 := 0; fi88 < 18; fi88++ {
                            v4 = v4
                            if len(v6) > 0 {
                                v6[0] = ("\n\t" + "a")
                            }
                            _ = fi88
                        }
                    }
                    switch "\n\t" {
                    case "a":
                        _ = len(v6)
                    }
                    if (!false) {
                        var v89 int8 = 49
                        v6 = append(v6, "a")
                        for fi90 := 0; fi90 < 11; fi90++ {
                            switch (-(fi90)) {
                            case -975:
                                v6 = append(v6, ("" + ("a" + "\n\t")))
                            }
                            v79 = v79
                            _ = fi90
                        }
                        _ = v89
                    } else {
                        var v91 float32 = (3.14159 + 1.5)
                        _ = v91
                    }
                    _ = fi82
                }
            case "\n\t":
                var v92 uint32 = 2147483647
                _ = v92
            default:
                v93 := v1[("\n\t" + "hello")]
                _ = v93
            }
        case 127:
            if true {
                v94 := GenS6{F0: -32768, F1: -16}
                v77 = (v77 % ((42) | 1))
                _ = v94
            } else {
                v95 := v1["foo bar"]
                _ = v95
            }
        }
        _ = v79
    } else {
        var v96 int = 128
        if false {
            if true {
                switch 128 {
                case -2147483648:
                    v97 := GenS5{F0: 42}
                    _ = v97
                case 42:
                    var v98 int = (v96 ^ (v96 % ((775) | 1)))
                    _ = v98
                case 32767:
                    fmt.Println(v2, v77)
                default:
                    switch ("" + "hello") {
                    case "\n\t":
                        var v99 int32 = int32(v4)
                        _ = v99
                    default:
                        switch (v96 / ((9223372036854775807) | 1)) {
                        case 433:
                            fmt.Println(v2, v2)
                        case -9223372036854775808:
                            v100 := GenS2{F0: "\n\t", F1: 0, F2: (375 < v4)}
                            _ = v100
                        }
                    }
                }
                for fi101 := 0; fi101 < 17; fi101++ {
                    fmt.Println(fi101, v5)
                    _ = len(v1)
                    _ = fi101
                }
                fmt.Println(v4, v4)
            }
        }
        if (1 < 0) {
            if true {
                _ = len(v6)
                fmt.Println(v3, v3, v96)
            }
            v5 -= v5
            switch ("a" + ("a" + ("hello" + "hello"))) {
            case "foo bar":
                if (!true) {
                    fmt.Println(v2, v5)
                } else {
                    var v102 int8 = -7
                    fmt.Println(v96)
                    _ = v102
                }
            default:
                if false {
                    fmt.Println(v96, v77, v3)
                    if true {
                        v96 -= 32768
                        fmt.Println(v96, v3, v2)
                        for fi103 := 0; fi103 < 15; fi103++ {
                            var v104 uint16 = 42
                            for fi105 := 0; fi105 < 5; fi105++ {
                                if true {
                                    if false {
                                        var v106 int = (fi103 + (fi103 & fi105))
                                        _ = v106
                                    }
                                    v96 = 9223372036854775807
                                    v5 += (v5 - (v5 ^ v5))
                                    if false {
                                        v107 := GenS4{F0: v5, F1: fi105, F2: 18446744073709551615}
                                        _ = v107
                                    } else {
                                        if true {
                                            v4 = (v4 + 42)
                                            var v108 bool = (true || true)
                                            if v108 {
                                                if (v108 || (0 >= v3)) {
                                                    v3 = 2
                                                } else {
                                                    v4 = v4
                                                }
                                                if len(v6) > 0 {
                                                    v6[0] = ("" + ("\n\t" + "hello"))
                                                }
                                            } else {
                                                fmt.Println(v5, v77)
                                            }
                                            _ = v108
                                        }
                                        v109 := GenS1{F0: (v4 | uint32(v5)), F1: (-(v96))}
                                        _ = v109
                                    }
                                    v110 := v1["\n\t"]
                                    _ = v110
                                }
                                var v111 float32 = 1e10
                                _ = fi105
                                _ = v111
                            }
                            fmt.Println(v4, v4, v104)
                            _ = fi103
                            _ = v104
                        }
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            }
        }
        _ = v96
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v76
    _ = v77
    _ = v78
}
