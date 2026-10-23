package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int32; F1 int; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint16; F2 int32 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 int16; F1 bool; F2 uint64 }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 int16; F1 int16; F2 uint32 }
type GenS12 struct { F0 uint16; F1 uint32; F2 int32 }
func genFn0(p0 uint8, p1 int, p2 int16) uint16 {
    if (false && false) {
        fmt.Println(p0, p1)
        v7 := []int{}
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(fi8, p0)
        var v9 string = ("" + (("hello" + (((("a" + ("\n\t" + "foo bar")) + "") + "a") + (((("foo bar" + "") + ("" + "foo bar")) + ("" + "a")) + ""))) + "a"))
        _ = fi8
        _ = v9
    }
    v10 := []int{(p1 >> ((2) & 63)), p1, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 617
}

func main() {
    var v1 int = 32768
    var v2 float64 = 0.0
    var v3 uint32 = 2147483648
    var v4 int64 = -1
    var v5 int = int(v4)
    v6 := make(map[int]bool)
    var v11 uint = 255
    v12 := []int{v5}
    _ = len(v6)
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int = v1
        for fi15 := 0; fi15 < 6; fi15++ {
            switch (v5 << ((42) & 63)) {
            case 42:
                v5 += -142
            }
            var v16 uint64 = 4294967295
            _ = fi15
            _ = v16
        }
        v4 = (v4 / ((v4) | 1))
        _ = fi13
        _ = v14
    }
    var v17 uint = v11
    if true {
        for fi18 := 0; fi18 < 20; fi18++ {
            for fi19 := 0; fi19 < 10; fi19++ {
                if false {
                    if ((410 > 1) && true) {
                        var v20 uint16 = 791
                        v3 += v3
                        _ = v20
                    }
                    if false {
                        v21 := GenS1{F0: (v3 - v3), F1: "a"}
                        var v22 int8 = 0
                        var v23 uint64 = 18446744073709551615
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    v24 := GenS2{F0: "hello", F1: 255}
                    _ = v24
                }
                if ((65535 != genFn0(255, int(v3), -1)) && true) {
                    fmt.Println(v11, v1)
                    var v25 bool = false
                    v26 := GenS3{F0: -1, F1: 32767, F2: v3}
                    _ = v25
                    _ = v26
                } else {
                    if (!true) {
                        if (true && true) {
                            if len(v12) > 0 {
                                v12[0] = (v5 * (v1 & v1))
                            }
                            if false {
                                v3 -= (v3 + v3)
                                var v27 bool = false
                                switch "a" {
                                case "a":
                                    if (true && v27) {
                                        v17 += (v17 % ((42) | 1))
                                        if (v27 && v27) {
                                            if true {
                                                fmt.Println(v2, v3, v27)
                                                var v28 int8 = 60
                                                if true {
                                                    var v29 uint16 = genFn0(163, (v1 << ((42) & 63)), -1)
                                                    fmt.Println(v27, v28, fi18)
                                                    switch "" {
                                                    case "hello":
                                                        v28 += (v28 - 77)
                                                    case "foo bar":
                                                        v30 := v6[fi18]
                                                        _ = v30
                                                    case "a":
                                                        fmt.Println(v17, v29, v2)
                                                    }
                                                    _ = v29
                                                } else {
                                                    if v27 {
                                                        v31 := GenS4{F0: (^(v11)), F1: 2, F2: -414}
                                                        _ = v31
                                                    }
                                                }
                                                _ = v28
                                            } else {
                                                var v32 uint16 = 32767
                                                if v27 {
                                                    var v33 uint = 96
                                                    var v34 uint32 = uint32(v32)
                                                    var v35 float32 = 1e10
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    if false {
                                                        fmt.Println(v11, v3)
                                                        var v36 uint = (v11 * 18446744073709551615)
                                                        _ = v36
                                                    } else {
                                                        var v37 uint32 = v3
                                                        _ = v37
                                                    }
                                                }
                                                if (true || v27) {
                                                    var v38 string = (((("hello" + (("a" + "a") + "foo bar")) + "hello") + "a") + "\n\t")
                                                    _ = v38
                                                }
                                                _ = v32
                                            }
                                            v1 += (v5 % (((fi18 | fi18)) | 1))
                                            fmt.Println(v1, fi18)
                                        } else {
                                            var v39 float32 = 0.0
                                            v40 := v6[v5]
                                            v6[-9223372036854775808] = ((v1 & v1) == 753)
                                            v2 = -2.25
                                            v41 := GenS4{F0: 885, F1: 1, F2: 1}
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                        v5 = v5
                                    } else {
                                        if len(v12) > 0 {
                                            v12[0] = 32767
                                        }
                                        var v42 int64 = (v4 % (((v4 / ((v4) | 1))) | 1))
                                        _ = v42
                                    }
                                case "foo bar":
                                    var v43 bool = false
                                    _ = v43
                                default:
                                    if (-128 > 100) {
                                        v5 += v5
                                        if v27 {
                                            v44 := GenS2{F0: "foo bar", F1: 433}
                                            v45 := GenS5{F0: 127, F1: true}
                                            _ = v44
                                            _ = v45
                                        }
                                        v46 := GenS6{F0: (v17 % ((v17) | 1))}
                                        _ = v46
                                    } else {
                                        var v47 int64 = (v4 & (v4 & 730))
                                        v48 := GenS6{F0: v17}
                                        v5 = (fi18 & (fi19 % ((fi19) | 1)))
                                        _ = v47
                                        _ = v48
                                    }
                                }
                                _ = v27
                            } else {
                                if (("a" > "a") && true) {
                                    v11 = 394
                                    if true {
                                        var v49 float32 = 1.5
                                        v4 += (v4 + v4)
                                        _ = v49
                                    } else {
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        switch (^(v1)) {
                                        case 32768:
                                            if ((0 > 127) || false) {
                                                fmt.Println(fi18)
                                                switch -2147483648 {
                                                case -822:
                                                    if (1 == 2147483647) {
                                                        if ((590 != (v11 ^ v17)) && (false && true)) {
                                                            v17 = v11
                                                            switch (fi19 & 127) {
                                                            case 1:
                                                                var v50 string = "a"
                                                                _ = v50
                                                            }
                                                        }
                                                        switch (v5 << (((v17 << (((^(v17))) & 63))) & 63)) {
                                                        case -703:
                                                            v3 = v3
                                                        }
                                                        _ = len(v12)
                                                    }
                                                case 325:
                                                    fmt.Println(v3, v1, v5)
                                                case -2147483648:
                                                    v51 := GenS7{F0: false}
                                                    _ = v51
                                                }
                                                var v52 bool = (((fi18 * fi19) >= 412) || ((true || false) && (1.5 < 3.14159)))
                                                _ = v52
                                            } else {
                                                fmt.Println(fi18, v3, v11)
                                                v53 := GenS8{F0: 1, F1: 100}
                                                switch ("a" + "\n\t") {
                                                case "a":
                                                    v1 += -2147483648
                                                case "":
                                                    delete(v6, (-(v1)))
                                                }
                                                _ = v53
                                            }
                                        }
                                        v1 = (fi19 + v1)
                                    }
                                }
                                var v54 int32 = 42
                                _ = v54
                            }
                        } else {
                            var v55 int64 = (v4 - 42)
                            _ = v55
                        }
                        v5 -= v5
                        fmt.Println(v17, v5)
                    }
                    v17 = (^(v11))
                    v3 -= (^(v3))
                }
                _ = fi19
            }
            _ = fi18
        }
        var v56 float32 = (-2.25 + -2.25)
        _ = len(v12)
        _ = v56
    }
    _ = len(v6)
    if false {
        var v57 int = (v5 + 4294967295)
        _ = v57
    } else {
        v17 -= v11
        for fi58 := 0; fi58 < 14; fi58++ {
            if true {
                var v59 int32 = -857
                var v60 uint32 = (^(v3))
                _ = v59
                _ = v60
            }
            v4 = (v4 * v4)
            _ = fi58
        }
        v2 += float64(((float32(float64(1e10)) - 1e10) - float32(float64(3.14159))))
    }
    v17 = v17
    for fi61 := 0; fi61 < 3; fi61++ {
        for fi62 := 0; fi62 < 10; fi62++ {
            if len(v12) > 0 {
                v12[0] = (-(v1))
            }
            _ = fi62
        }
        var v63 uint64 = 845
        if (v11 <= 65535) {
            _ = len(v6)
            fmt.Println(fi61, v2, v3)
        } else {
            v6[int(v63)] = false
        }
        _ = fi61
        _ = v63
    }
    if len(v12) > 0 {
        v12[0] = 9223372036854775807
    }
    v64 := GenS7{F0: (true || false)}
    for fi65 := 0; fi65 < 17; fi65++ {
        switch v5 {
        case -828:
            fmt.Println(v4, v1, v5)
        case 598:
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        case 127:
            if true {
                fmt.Println(v4)
                var v66 uint16 = 591
                switch (v1 << (((v17 / ((v17) | 1))) & 63)) {
                case 479:
                    if (!false) {
                        for fi67 := 0; fi67 < 5; fi67++ {
                            v5 -= (-(fi67))
                            if true {
                                switch (fi65 * (fi67 + (fi67 & fi67))) {
                                case -663:
                                    var v68 uint64 = 340
                                    _ = v68
                                case -1:
                                    v17 = (^(v17))
                                default:
                                    if (false && (!(v4 == (v4 / ((1) | 1))))) {
                                        switch ("foo bar" + ("" + "hello")) {
                                        case "hello":
                                            var v69 uint8 = 0
                                            _ = v69
                                        case "\n\t":
                                            if true {
                                                v70 := GenS9{F0: 0, F1: (!true), F2: 724}
                                                if len(v12) > 0 {
                                                    v12[0] = (v5 * (v1 << (((v11 + v11)) & 63)))
                                                }
                                                v12 = append(v12, (v5 >> ((v17) & 63)))
                                                _ = v70
                                            }
                                        default:
                                            _ = len(v12)
                                        }
                                        v4 = v4
                                    }
                                }
                            }
                            switch "\n\t" {
                            case "foo bar":
                                var v71 int32 = 0
                                var v72 int16 = 0
                                _ = v71
                                _ = v72
                            }
                            _ = fi67
                        }
                        v73 := GenS3{F0: 2147483647, F1: 2147483647, F2: v3}
                        v11 -= 813
                        _ = v73
                    } else {
                        var v74 string = "a"
                        switch v74 {
                        case "foo bar":
                            _ = len(v12)
                        case "a":
                            var v75 int8 = -128
                            _ = v75
                        case "":
                            for fi76 := 0; fi76 < 20; fi76++ {
                                var v77 int32 = 2147483647
                                switch (fi65 ^ 4294967295) {
                                case 128:
                                    v5 = 9223372036854775807
                                    switch (v74 + (("\n\t" + v74) + (v74 + v74))) {
                                    case "a":
                                        if (-1 != -118) {
                                            var v78 int8 = -1
                                            _ = v78
                                        }
                                    case "foo bar":
                                        var v79 float64 = v2
                                        _ = v79
                                    }
                                    v80 := GenS1{F0: uint32(v66), F1: v74}
                                    _ = v80
                                case -2147483648:
                                    v81 := GenS1{F0: (v3 + v3), F1: v74}
                                    _ = v81
                                case 32768:
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                }
                                _ = fi76
                                _ = v77
                            }
                        default:
                            if true {
                                var v82 int = (v1 / ((v5) | 1))
                                _ = v82
                            } else {
                                for fi83 := 0; fi83 < 1; fi83++ {
                                    var v84 uint = v11
                                    _ = fi83
                                    _ = v84
                                }
                                if false {
                                    for fi85 := 0; fi85 < 19; fi85++ {
                                        if (900 > 2) {
                                            var v86 uint64 = 42
                                            if len(v12) > 0 {
                                                v12[0] = (v1 | fi65)
                                            }
                                            _ = v86
                                        }
                                        v5 = int(v66)
                                        _ = fi85
                                    }
                                } else {
                                    switch v74 {
                                    case "a":
                                        v87 := GenS10{F0: -986, F1: v11}
                                        _ = v87
                                    case "foo bar":
                                        if len(v12) > 0 {
                                            v12[0] = 32767
                                        }
                                    case "hello":
                                        var v88 uint8 = 1
                                        _ = v88
                                    }
                                    var v89 string = v74
                                    v2 -= v2
                                    _ = v89
                                }
                            }
                        }
                        v3 -= 422
                        _ = v74
                    }
                case 4294967295:
                    switch "" {
                    case "":
                        v90 := GenS2{F0: "\n\t", F1: 2}
                        _ = v90
                    case "\n\t":
                        v91 := GenS11{F0: 42, F1: 0, F2: v3}
                        _ = v91
                    case "foo bar":
                        v5 = (^(fi65))
                    default:
                        for fi92 := 0; fi92 < 18; fi92++ {
                            v4 += (v4 << (((v17 >> (((^(v11))) & 63))) & 63))
                            v66 += (^(v66))
                            fmt.Println(fi65)
                            _ = fi92
                        }
                    }
                case -9223372036854775808:
                    if (v3 != (v3 + v3)) {
                        switch v5 {
                        case -1:
                            v93 := GenS9{F0: 0, F1: false, F2: 0}
                            _ = v93
                        case 32767:
                            if false {
                                v17 = (v17 + 972)
                            } else {
                                v3 = (^(v3))
                                switch -664 {
                                case 42:
                                    v94 := GenS7{F0: (42 < 0)}
                                    _ = v94
                                default:
                                    v66 += genFn0(uint8(v11), 2147483647, 42)
                                }
                                for fi95 := 0; fi95 < 18; fi95++ {
                                    v4 = v4
                                    switch fi65 {
                                    case 127:
                                        v17 = v17
                                    case -2147483648:
                                        var v96 string = ""
                                        _ = v96
                                    case 128:
                                        switch "a" {
                                        case "foo bar":
                                            v97 := GenS12{F0: v66, F1: v3, F2: 1}
                                            _ = v97
                                        default:
                                            v98 := v6[(fi65 - fi95)]
                                            _ = v98
                                        }
                                    default:
                                        v99 := GenS7{F0: true}
                                        _ = v99
                                    }
                                    var v100 uint32 = v3
                                    _ = fi95
                                    _ = v100
                                }
                            }
                        case 127:
                            v4 = (v4 + v4)
                        }
                        for fi101 := 0; fi101 < 18; fi101++ {
                            fmt.Println(v3, v66)
                            v2 += 1.5
                            var v102 int64 = 0
                            _ = fi101
                            _ = v102
                        }
                        var v103 uint32 = v3
                        _ = v103
                    }
                default:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                }
                _ = v66
            }
        default:
            for fi104 := 0; fi104 < 10; fi104++ {
                var v105 int64 = v4
                fmt.Println(v4)
                _ = fi104
                _ = v105
            }
        }
        if (0 > 65535) {
            var v106 uint = (^(v11))
            _ = v106
        } else {
            v5 += v5
        }
        fmt.Println(v3, v2)
        _ = fi65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
