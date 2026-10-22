package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int; F2 bool }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 string; F1 int64; F2 int64 }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint16; F1 bool }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 string }

func main() {
    v1 := []int{128, 32768, 944, -2147483648}
    v2 := make(map[string]string)
    var v3 uint64 = 339
    var v4 int64 = 1
    var v5 uint = 18446744073709551615
    var v6 float32 = 1e10
    var v7 int16 = 384
    if (((v6 + (v6 - (v6 - (v6 / v6)))) == (v6 * v6)) || (true || false)) {
        v7 = v7
    } else {
        var v8 uint = (v5 >> ((18446744073709551615) & 63))
        _ = v8
    }
    switch 160 {
    case 832:
        var v9 uint64 = 695
        _ = v9
    }
    var v10 uint16 = 32767
    v6 += v6
    v5 -= v5
    if true {
        var v11 uint16 = (v10 / (((^(v10))) | 1))
        _ = v11
    } else {
        switch -664 {
        case -327:
            _ = len(v1)
        default:
            if true {
                var v12 bool = true
                v7 -= (v7 / ((505) | 1))
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            for fi14 := 0; fi14 < 2; fi14++ {
                fmt.Println(v5)
                v15 := GenS1{F0: (v4 * v4), F1: 128}
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        v3 += 534
    }
    var v16 uint = 65535
    v7 += v7
    switch 127 {
    case -9223372036854775808:
        switch ("hello" + "hello") {
        case "a":
            if false {
                for fi17 := 0; fi17 < 6; fi17++ {
                    var v18 int16 = v7
                    var v19 uint16 = (^(v10))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 int16 = (v7 * v7)
                v4 = (v4 >> ((v16) & 63))
                _ = v20
            } else {
                v16 = (^(v5))
                switch "\n\t" {
                case "hello":
                    var v21 uint8 = 127
                    _ = v21
                case "\n\t":
                    if true {
                        if false {
                            var v22 uint = 650
                            for fi23 := 0; fi23 < 13; fi23++ {
                                v24 := GenS2{F0: (fi23 + fi23)}
                                fmt.Println(v16, fi23, v22)
                                fmt.Println(v22, v22)
                                _ = fi23
                                _ = v24
                            }
                            _ = v22
                        }
                        if false {
                            switch -1 {
                            case 2147483647:
                                v25 := GenS1{F0: v4, F1: 0}
                                _ = v25
                            case 128:
                                v26 := GenS1{F0: v4, F1: 128}
                                _ = v26
                            case 277:
                                fmt.Println(v3)
                            }
                            var v27 uint16 = (^(v10))
                            if ("a" >= ("" + "\n\t")) {
                                for fi28 := 0; fi28 < 4; fi28++ {
                                    v6 -= v6
                                    _ = len(v1)
                                    v3 -= 12
                                    _ = fi28
                                }
                            } else {
                                for fi29 := 0; fi29 < 2; fi29++ {
                                    v30 := GenS1{F0: (v4 & (v4 - v4)), F1: 199}
                                    var v31 uint64 = v3
                                    _ = fi29
                                    _ = v30
                                    _ = v31
                                }
                                var v32 uint32 = 526
                                fmt.Println(v3)
                                _ = v32
                            }
                            _ = v27
                        } else {
                            for fi33 := 0; fi33 < 19; fi33++ {
                                for fi34 := 0; fi34 < 3; fi34++ {
                                    delete(v2, ("foo bar" + ""))
                                    v35 := GenS2{F0: fi34}
                                    v7 = v7
                                    v5 += 42
                                    if (true && true) {
                                        switch (((("foo bar" + "\n\t") + ("" + ("" + ("a" + ("a" + "\n\t"))))) + ("" + ("hello" + "a"))) + "foo bar") {
                                        case "hello":
                                            v36 := GenS3{F0: 254, F1: int(v10), F2: false}
                                            var v37 int8 = -22
                                            if (((!(!true)) || true) && ((-2.25 * float32((3.14159 / 0.1))) >= (float32(((-2.25 / -2.25) / 1e10)) - 0.0))) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                if (16 < (fi34 * fi33)) {
                                                    var v38 uint64 = v3
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v38
                                                }
                                            }
                                            _ = v36
                                            _ = v37
                                        case "\n\t":
                                            v39 := GenS3{F0: 127, F1: (fi34 % (((fi33 - (-(fi33)))) | 1)), F2: ((v4 * 42) != v4)}
                                            _ = v39
                                        case "a":
                                            var v40 bool = (-9223372036854775808 >= v4)
                                            _ = v40
                                        }
                                    }
                                    fmt.Println(v6, v6)
                                    _ = fi34
                                    _ = v35
                                }
                                var v41 bool = true
                                _ = fi33
                                _ = v41
                            }
                        }
                        v4 += 42
                    }
                case "":
                    if false {
                        for fi42 := 0; fi42 < 15; fi42++ {
                            var v43 uint64 = v3
                            fmt.Println(v5)
                            var v44 float64 = (((1e10 / 1.0) - 0.0) + (1e10 / 1e10))
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        v7 = v7
                    } else {
                        for fi45 := 0; fi45 < 1; fi45++ {
                            for fi46 := 0; fi46 < 19; fi46++ {
                                if true {
                                    v47 := GenS4{F0: 4294967295, F1: true}
                                    v7 += (v7 >> (((v5 * v5)) & 63))
                                    _ = v47
                                } else {
                                    var v48 uint64 = (^(v3))
                                    v49 := GenS5{F0: 57, F1: (v10 ^ 255)}
                                    _ = v48
                                    _ = v49
                                }
                                _ = fi46
                            }
                            v3 -= 733
                            _ = fi45
                        }
                        if (-128 >= int8(v3)) {
                            switch -2147483648 {
                            case 734:
                                if (v4 < 1) {
                                    fmt.Println(v7, v5)
                                    v50 := GenS2{F0: 1}
                                    _ = v50
                                }
                            case -609:
                                fmt.Println(v3, v16)
                            case 1:
                                v5 += 42
                            default:
                                switch 32767 {
                                case -1:
                                    if true {
                                        var v51 uint8 = 1
                                        var v52 uint8 = 1
                                        if true {
                                            if ((v7 % ((v7) | 1)) <= v7) {
                                                switch 128 {
                                                case 750:
                                                    var v53 int32 = -696
                                                    _ = v53
                                                default:
                                                    if true {
                                                        if (true && (true && true)) {
                                                            for fi54 := 0; fi54 < 6; fi54++ {
                                                                v52 += (^(v51))
                                                                v4 = (v4 - v4)
                                                                v52 -= 255
                                                                _ = fi54
                                                            }
                                                        }
                                                        var v55 uint64 = 0
                                                        var v56 string = "foo bar"
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                }
                                                delete(v2, "")
                                            } else {
                                                var v57 uint16 = v10
                                                var v58 uint16 = (v10 % ((v10) | 1))
                                                _ = v57
                                                _ = v58
                                            }
                                            if false {
                                                if true {
                                                    var v59 uint = (v16 << ((uint(v10)) & 63))
                                                    v60 := GenS6{F0: "\n\t", F1: -1, F2: (v4 * v4)}
                                                    _ = v59
                                                    _ = v60
                                                } else {
                                                    var v61 int8 = -128
                                                    var v62 uint = v5
                                                    _ = v61
                                                    _ = v62
                                                }
                                                var v63 int16 = (v7 >> (((v16 ^ v16)) & 63))
                                                for fi64 := 0; fi64 < 1; fi64++ {
                                                    v51 -= uint8(v5)
                                                    switch "a" {
                                                    case "foo bar":
                                                        v65 := GenS6{F0: ("foo bar" + "a"), F1: (v4 & (v4 ^ v4)), F2: int64(v10)}
                                                        _ = v65
                                                    default:
                                                        for fi66 := 0; fi66 < 18; fi66++ {
                                                            var v67 int32 = 42
                                                            v68 := v2["hello"]
                                                            _ = fi66
                                                            _ = v67
                                                            _ = v68
                                                        }
                                                    }
                                                    var v69 uint32 = 42
                                                    _ = fi64
                                                    _ = v69
                                                }
                                                _ = v63
                                            } else {
                                                if (int8(v4) > -1) {
                                                    if true {
                                                        if false {
                                                            if false {
                                                                v10 += v10
                                                                if (!(!((v51 + v52) != 141))) {
                                                                    if true {
                                                                        _ = len(v2)
                                                                    } else {
                                                                        if false {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 1
                                                                            }
                                                                            var v70 int16 = v7
                                                                            _ = v70
                                                                        } else {
                                                                            var v71 int8 = 42
                                                                            v2["\n\t"] = ("" + "hello")
                                                                            _ = v71
                                                                        }
                                                                    }
                                                                }
                                                                if true {
                                                                    var v72 int16 = 1
                                                                    var v73 float32 = (1e10 + 1e10)
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                            }
                                                            fmt.Println(v3, v4, v5)
                                                            v52 += uint8(v3)
                                                        }
                                                        var v74 int16 = 634
                                                        _ = v74
                                                    }
                                                    switch -173 {
                                                    case 9223372036854775807:
                                                        v10 = (v10 % ((v10) | 1))
                                                    case 4294967295:
                                                        v7 += (v7 << ((v16) & 63))
                                                    case 128:
                                                        var v75 bool = (v6 > 0.1)
                                                        _ = v75
                                                    }
                                                    var v76 uint64 = (^(v3))
                                                    _ = v76
                                                } else {
                                                    fmt.Println(v7)
                                                    v52 = (^(v51))
                                                    v51 = (v51 % (((v51 ^ 0)) | 1))
                                                }
                                                v3 = 2
                                            }
                                        }
                                        _ = v51
                                        _ = v52
                                    } else {
                                        var v77 int16 = 8
                                        var v78 int16 = (^(v77))
                                        _ = v77
                                        _ = v78
                                    }
                                case -922:
                                    if (v4 < (^(v4))) {
                                        if (("\n\t" >= "foo bar") && (!false)) {
                                            v79 := GenS7{F0: int32(v3), F1: 0}
                                            var v80 float32 = v6
                                            var v81 int8 = 127
                                            _ = v79
                                            _ = v80
                                            _ = v81
                                        } else {
                                            if (-2.25 >= -2.25) {
                                                v10 = v10
                                                v16 = v16
                                            } else {
                                                for fi82 := 0; fi82 < 7; fi82++ {
                                                    for fi83 := 0; fi83 < 12; fi83++ {
                                                        var v84 bool = (953 >= 42)
                                                        v2[("\n\t" + "hello")] = ""
                                                        _ = fi83
                                                        _ = v84
                                                    }
                                                    _ = fi82
                                                }
                                                for fi85 := 0; fi85 < 13; fi85++ {
                                                    switch fi85 {
                                                    case -9223372036854775808:
                                                        v3 += 297
                                                    case -2147483648:
                                                        var v86 uint = uint(v10)
                                                        _ = v86
                                                    case 661:
                                                        v7 += int16(v4)
                                                    default:
                                                        for fi87 := 0; fi87 < 16; fi87++ {
                                                            switch (("a" + (((("\n\t" + "foo bar") + "foo bar") + "a") + "")) + ("\n\t" + "a")) {
                                                            case "foo bar":
                                                                v88 := GenS8{F0: v5}
                                                                _ = v88
                                                            case "hello":
                                                                var v89 uint64 = (^(v3))
                                                                _ = v89
                                                            default:
                                                                var v90 uint = v5
                                                                _ = v90
                                                            }
                                                            v6 += (v6 - (v6 * 0.0))
                                                            _ = fi87
                                                        }
                                                    }
                                                    switch (fi85 - (fi85 & (fi85 ^ (fi85 ^ -2147483648)))) {
                                                    case 2147483647:
                                                        var v91 uint32 = 65535
                                                        _ = v91
                                                    case -2147483648:
                                                        v92 := GenS9{F0: ((1.5 != 1e10) || false)}
                                                        _ = v92
                                                    case 1:
                                                        v93 := GenS1{F0: (v4 >> ((4294967295) & 63)), F1: 35}
                                                        _ = v93
                                                    }
                                                    _ = len(v1)
                                                    _ = fi85
                                                }
                                                var v94 int16 = (v7 | (v7 & v7))
                                                _ = v94
                                            }
                                            if (1.5 == float64(((v6 + -2.25) / v6))) {
                                                var v95 bool = ("a" > ("\n\t" + "\n\t"))
                                                var v96 uint8 = 0
                                                _ = v95
                                                _ = v96
                                            } else {
                                                var v97 int = -1
                                                _ = v97
                                            }
                                        }
                                        if (103 != 102) {
                                            _ = len(v1)
                                            switch "foo bar" {
                                            case "":
                                                v3 = v3
                                            case "hello":
                                                for fi98 := 0; fi98 < 13; fi98++ {
                                                    v99 := GenS10{F0: 2, F1: ((fi98 << ((42) & 63)) >= fi98)}
                                                    if ((239 < -2147483648) || ((((0.1 + ((1.5 + (1.5 + 3.14159)) * 0.0)) - 0.1) + (1e10 / -2.25)) <= (float64(1.5) * 1e10))) {
                                                        switch "a" {
                                                        case "foo bar":
                                                            if false {
                                                                v2[("foo bar" + "\n\t")] = ("hello" + "foo bar")
                                                                v100 := GenS6{F0: "foo bar", F1: v4, F2: (v4 * -1)}
                                                                var v101 int8 = -115
                                                                _ = v100
                                                                _ = v101
                                                            } else {
                                                                switch "foo bar" {
                                                                case "a":
                                                                    var v102 int16 = v7
                                                                    _ = v102
                                                                }
                                                                v6 += 1.5
                                                            }
                                                        default:
                                                            if (true && (3.14159 < 1e10)) {
                                                                fmt.Println(v16, v3)
                                                                fmt.Println(v5)
                                                                v10 += v10
                                                            }
                                                        }
                                                        v103 := GenS11{F0: v5}
                                                        fmt.Println(fi98, fi98, v4)
                                                        _ = v103
                                                    } else {
                                                        v104 := GenS12{F0: (("a" + ("" + ((("\n\t" + ((("hello" + "\n\t") + "a") + "")) + "\n\t") + "hello"))) + ("" + "hello"))}
                                                        _ = v104
                                                    }
                                                    v4 = v4
                                                    _ = fi98
                                                    _ = v99
                                                }
                                            case "foo bar":
                                                if (((0.1 * 0.0) / 0.1) < 3.14159) {
                                                    v5 += 18446744073709551615
                                                    switch "hello" {
                                                    case "a":
                                                        fmt.Println(v5)
                                                    case "\n\t":
                                                        fmt.Println(v3, v7, v16)
                                                    default:
                                                        v16 = v16
                                                    }
                                                    v16 = v16
                                                }
                                            default:
                                                v3 -= (^(v3))
                                            }
                                        }
                                    }
                                default:
                                    fmt.Println(v6)
                                }
                            }
                            for fi105 := 0; fi105 < 4; fi105++ {
                                var v106 int = fi105
                                _ = fi105
                                _ = v106
                            }
                            var v107 uint16 = 255
                            _ = v107
                        } else {
                            v108 := GenS3{F0: 0, F1: 4294967295, F2: true}
                            v16 = (^(v5))
                            if (true && (false || false)) {
                                var v109 float64 = ((float64(v6) * (1e10 / 1e10)) / 1.0)
                                if false {
                                    var v110 int8 = -128
                                    _ = v110
                                }
                                v5 = 42
                                _ = v109
                            }
                            _ = v108
                        }
                    }
                }
            }
        }
    }
    if true {
        v7 -= v7
    }
    var v111 uint64 = 919
    v3 = v3
    v111 = (^(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v111:", v111)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v111
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
