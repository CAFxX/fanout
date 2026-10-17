package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint8; F1 uint; F2 int64 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 uint64; F1 int; F2 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8; F1 string }
type GenS9 struct { F0 int }
func genFn0(p0 uint, p1 uint8, p2 uint16) uint64 {
    var v5 string = ("a" + ("\n\t" + ("foo bar" + "foo bar")))
    fmt.Println(p1, p1, p1)
    v5 = v5
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 823
}

func main() {
    var v1 bool = false
    var v2 uint8 = 0
    var v3 uint32 = 890
    var v4 uint32 = 275
    v6 := []int{419, 42, 127, 42}
    for fi7 := 0; fi7 < 7; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 string = ""
            _ = len(v6)
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 1000, F1: uint(v3)}
        var v11 int = 42
        _ = fi7
        _ = v10
        _ = v11
    }
    var v12 int = 4294967295
    switch ("" + ("\n\t" + ("a" + ("foo bar" + "\n\t")))) {
    case "foo bar":
        v13 := GenS2{F0: v2, F1: 0, F2: 448}
        _ = v13
    case "\n\t":
        v14 := GenS1{F0: 42, F1: 1}
        _ = v14
    case "hello":
        switch ((("" + ("" + "")) + "") + "a") {
        case "\n\t":
            if v1 {
                var v15 int32 = 308
                _ = v15
            }
        case "":
            var v16 uint = 4294967295
            _ = v16
        default:
            for fi17 := 0; fi17 < 8; fi17++ {
                v4 = v3
                if true {
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 int = -1
                        v4 = v4
                        v12 += -9223372036854775808
                        _ = fi18
                        _ = v19
                    }
                    v2 = (v2 ^ (v2 + v2))
                    v3 = (v4 % ((v3) | 1))
                } else {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        v12 = -511
                        switch (v12 % (((fi17 | fi17)) | 1)) {
                        case 32767:
                            if v1 {
                                var v21 int64 = -25
                                _ = v21
                            } else {
                                if ((v1 || ((!(0 >= -694)) && v1)) || v1) {
                                    if (!v1) {
                                        var v22 uint32 = (v3 >> ((42) & 63))
                                        switch (fi17 + v12) {
                                        case -1:
                                            v12 += (v12 & fi17)
                                        default:
                                            if (v1 || (!true)) {
                                                fmt.Println(v22, v3)
                                                var v23 int8 = 0
                                                switch "hello" {
                                                case "hello":
                                                    var v24 string = ("a" + "foo bar")
                                                    _ = v24
                                                case "a":
                                                    if (!true) {
                                                        if v1 {
                                                            fmt.Println(v4, v22)
                                                        }
                                                        _ = len(v6)
                                                        v12 = (fi20 / (((fi20 - -25)) | 1))
                                                    } else {
                                                        var v25 int64 = -1
                                                        if len(v6) > 0 {
                                                            v6[0] = (-(fi17))
                                                        }
                                                        var v26 bool = v1
                                                        var v27 int16 = 285
                                                        var v28 uint32 = (v3 | (v4 ^ 1))
                                                        _ = v25
                                                        _ = v26
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                case "foo bar":
                                                    v23 -= v23
                                                }
                                                _ = v23
                                            }
                                        }
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        _ = v22
                                    } else {
                                        v29 := GenS3{F0: -269, F1: (v2 & 111)}
                                        var v30 int = fi20
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                if (42 >= 255) {
                                    switch "hello" {
                                    case "foo bar":
                                        if v1 {
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v31 int8 = 1
                                                _ = v31
                                            case "":
                                                switch ("" + "") {
                                                case "\n\t":
                                                    fmt.Println(fi20)
                                                default:
                                                    v6 = append(v6, fi17)
                                                }
                                            default:
                                                v12 -= (fi20 + fi20)
                                            }
                                        }
                                        v1 = v1
                                        v1 = (255 < 449)
                                    default:
                                        fmt.Println(v4, fi20, v4)
                                    }
                                    fmt.Println(fi20)
                                    fmt.Println(v1)
                                    v1 = v1
                                    var v32 uint = 18446744073709551615
                                    _ = v32
                                } else {
                                    var v33 uint8 = (v2 & v2)
                                    v33 += (v2 - v33)
                                    _ = v33
                                }
                                var v34 uint8 = 128
                                _ = v34
                            }
                        case -2147483648:
                            switch ("foo bar" + ("hello" + "\n\t")) {
                            case "a":
                                var v35 uint64 = 0
                                _ = v35
                            }
                        case -257:
                            if v1 {
                                v2 -= v2
                                if len(v6) > 0 {
                                    v6[0] = (^(fi17))
                                }
                            }
                        }
                        v12 -= (v12 ^ (v12 & (-(v12))))
                        _ = fi20
                    }
                }
                _ = fi17
            }
        }
    default:
        if true {
            v3 += v4
            switch "a" {
            case "foo bar":
                fmt.Println(v1, v3, v1)
            default:
                v36 := GenS4{F0: 2, F1: (v12 * v12), F2: (^(v12))}
                _ = v36
            }
        } else {
            var v37 int64 = -462
            for fi38 := 0; fi38 < 13; fi38++ {
                v1 = (v1 && true)
                v3 = (v4 - 42)
                for fi39 := 0; fi39 < 15; fi39++ {
                    if false {
                        fmt.Println(fi39, v37)
                        fmt.Println(v1, fi39, v2)
                        if (fi38 != int(v2)) {
                            if (0 == -521) {
                                fmt.Println(v2, v3)
                                v2 = 255
                                v37 -= (v37 % ((1) | 1))
                                var v40 string = "a"
                                if (42 <= 1) {
                                    v4 = (v4 - v3)
                                    if (true || (v2 <= (v2 >> ((uint(v4)) & 63)))) {
                                        v41 := GenS1{F0: 42, F1: 525}
                                        _ = v41
                                    } else {
                                        var v42 uint = 255
                                        v43 := GenS2{F0: (v2 << ((uint(v2)) & 63)), F1: (v42 - v42), F2: (v37 / ((-1) | 1))}
                                        if ((fi39 - 0) >= fi39) {
                                            v37 = v37
                                            if (v1 && (-210 != 1)) {
                                                if (42 < 788) {
                                                    v6 = append(v6, 128)
                                                    if false {
                                                        var v44 string = v40
                                                        v37 = (v37 | (^(v37)))
                                                        _ = v44
                                                    } else {
                                                        var v45 int64 = 0
                                                        var v46 uint = 18446744073709551615
                                                        v47 := GenS1{F0: -672, F1: 65535}
                                                        _ = v45
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                }
                                                v48 := GenS1{F0: 42, F1: v42}
                                                _ = v48
                                            } else {
                                                var v49 float32 = (0.1 / 1.0)
                                                var v50 int8 = -1
                                                _ = v49
                                                _ = v50
                                            }
                                            if len(v6) > 0 {
                                                v6[0] = fi39
                                            }
                                        } else {
                                            v12 = (fi39 << (((v42 | (v42 - v42))) & 63))
                                            var v51 int16 = 42
                                            v52 := GenS5{F0: v40}
                                            _ = v51
                                            _ = v52
                                        }
                                        _ = v42
                                        _ = v43
                                    }
                                    _ = len(v6)
                                }
                                _ = v40
                            }
                            var v53 uint16 = 65535
                            v53 -= v53
                            _ = v53
                        } else {
                            var v54 float32 = 1e10
                            if v1 {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if ((0 <= 1) || ((-795 >= 32767) && v1)) {
                                    v6 = append(v6, (fi38 - fi38))
                                } else {
                                    switch (-(v12)) {
                                    case 42:
                                        switch v12 {
                                        case -694:
                                            var v55 int16 = -1
                                            v3 -= (v3 << ((2) & 63))
                                            if false {
                                                var v56 int64 = v37
                                                _ = v56
                                            }
                                            _ = v55
                                        }
                                    case -384:
                                        v1 = (v1 && v1)
                                    }
                                    v57 := GenS1{F0: 32767, F1: 2}
                                    v12 = fi39
                                    _ = v57
                                }
                                v37 = v37
                            } else {
                                var v58 int = (fi39 * -1)
                                v58 -= 77
                                fmt.Println(v54)
                                _ = v58
                            }
                            _ = v54
                        }
                    } else {
                        if (1.5 >= 3.14159) {
                            var v59 int32 = -1
                            _ = v59
                        } else {
                            v1 = ((v4 * (^(v3))) == v3)
                            if (v1 && true) {
                                if v1 {
                                    fmt.Println(v37, v12)
                                }
                            }
                            var v60 int8 = 100
                            _ = v60
                        }
                        var v61 float64 = 0.1
                        fmt.Println(v2, fi38, v37)
                        _ = v61
                    }
                    v1 = v1
                    var v62 int16 = 42
                    _ = fi39
                    _ = v62
                }
                _ = fi38
            }
            _ = v37
        }
    }
    for fi63 := 0; fi63 < 2; fi63++ {
        for fi64 := 0; fi64 < 9; fi64++ {
            if (-92 >= 127) {
                fmt.Println(v4, fi64, fi63)
            }
            v4 = (v3 << ((42) & 63))
            if (!(3.14159 >= -2.25)) {
                var v65 int16 = 194
                if v1 {
                    var v66 int8 = -30
                    switch "hello" {
                    case "hello":
                        v66 = (v66 + (v66 | v66))
                    default:
                        switch v12 {
                        case -9223372036854775808:
                            var v67 uint = 4294967295
                            _ = v67
                        case 0:
                            if v1 {
                                if (v2 != 64) {
                                    v4 = v3
                                    fmt.Println(v4)
                                    v6 = append(v6, (-(fi64)))
                                }
                            } else {
                                var v68 int = fi64
                                var v69 float64 = 3.14159
                                var v70 int64 = -44
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                        case 128:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        default:
                            switch ("\n\t" + ("\n\t" + ("foo bar" + "\n\t"))) {
                            case "\n\t":
                                if len(v6) > 0 {
                                    v6[0] = fi64
                                }
                            case "hello":
                                v2 -= (^(v2))
                            case "a":
                                v65 = v65
                            default:
                                v6 = append(v6, fi63)
                            }
                        }
                    }
                    _ = v66
                } else {
                    v71 := GenS2{F0: v2, F1: 584, F2: 1}
                    if true {
                        v1 = true
                        var v72 float64 = (1.5 + 3.14159)
                        _ = v72
                    }
                    var v73 int8 = 42
                    v3 = (v4 + (v3 >> ((18446744073709551615) & 63)))
                    _ = v71
                    _ = v73
                }
                _ = v65
            }
            _ = fi64
        }
        v74 := GenS6{F0: -1}
        if true {
            var v75 bool = (1 == 1)
            _ = v75
        } else {
            for fi76 := 0; fi76 < 12; fi76++ {
                if true {
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if len(v6) > 0 {
                        v6[0] = 4294967295
                    }
                    if (v3 <= (v4 + (v3 << ((0) & 63)))) {
                        fmt.Println(fi63, v3)
                        if (!v1) {
                            if (v1 || ("\n\t" != "hello")) {
                                v12 -= (fi63 << ((408) & 63))
                            } else {
                                v2 = v2
                                switch (("foo bar" + "") + ("a" + "hello")) {
                                case "a":
                                    var v77 uint64 = 37
                                    _ = len(v6)
                                    var v78 bool = v1
                                    _ = v77
                                    _ = v78
                                case "\n\t":
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                switch "a" {
                                case "a":
                                    v2 -= (v2 - (v2 % ((98) | 1)))
                                default:
                                    var v79 uint64 = genFn0(1, v2, 32767)
                                    if v1 {
                                        fmt.Println(v2, v3, v79)
                                    } else {
                                        switch 9223372036854775807 {
                                        case 127:
                                            v2 -= (^(v2))
                                        case -279:
                                            switch "" {
                                            case "hello":
                                                switch ("" + "a") {
                                                case "":
                                                    v79 = genFn0(545, v2, 255)
                                                case "\n\t":
                                                    var v80 uint16 = 65535
                                                    _ = v80
                                                default:
                                                    v81 := GenS7{F0: 32767}
                                                    _ = v81
                                                }
                                            case "a":
                                                v6 = append(v6, fi76)
                                            case "":
                                                v1 = (-962 <= 32767)
                                            default:
                                                var v82 uint8 = uint8(v4)
                                                _ = v82
                                            }
                                        }
                                        if (0.1 >= (float64((float32((3.14159 / ((-2.25 + (0.0 / 1.0)) + 3.14159))) * 1e10)) + (1.5 * 1e10))) {
                                            v83 := GenS7{F0: 32767}
                                            switch "foo bar" {
                                            case "foo bar":
                                                v3 = (v3 / ((v4) | 1))
                                            case "hello":
                                                v3 = 65535
                                            }
                                            _ = v83
                                        }
                                        if v1 {
                                            var v84 uint32 = v3
                                            _ = v84
                                        }
                                    }
                                    v85 := GenS8{F0: 127, F1: (("" + "hello") + "foo bar")}
                                    _ = v79
                                    _ = v85
                                }
                            }
                        }
                        v6 = append(v6, fi63)
                        fmt.Println(v12, v4, v1)
                        v1 = (!v1)
                        var v86 uint64 = 506
                        _ = len(v6)
                        _ = v86
                    }
                }
                if (!((v12 > 42) || v1)) {
                    var v87 uint8 = (v2 / (((v2 * v2)) | 1))
                    v12 = int(v4)
                    var v88 float32 = 1.5
                    v6 = append(v6, (fi76 - 32768))
                    _ = v87
                    _ = v88
                }
                _ = fi76
            }
            fmt.Println(v2, fi63, v3)
        }
        _ = fi63
        _ = v74
    }
    if true {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    v89 := GenS4{F0: 608, F1: v12, F2: v12}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi90 := 0; fi90 < 9; fi90++ {
        v91 := GenS9{F0: -1}
        _ = len(v6)
        _ = fi90
        _ = v91
    }
    switch 938 {
    case 32768:
        fmt.Println(v1)
    default:
        _ = len(v6)
    }
    for fi92 := 0; fi92 < 17; fi92++ {
        var v93 float64 = float64(float32(1e10))
        _ = fi92
        _ = v93
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v89
}
