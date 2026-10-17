package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 int32; F1 uint64; F2 int64 }
type GenS5 struct { F0 int32; F1 int8 }
type GenS6 struct { F0 uint8; F1 uint8 }
func genFn0(p0 int, p1 uint8, p2 uint32) uint8 {
    if false {
        fmt.Println(p1)
        v4 := GenS1{F0: 0, F1: 2147483647, F2: uint16(p1)}
        _ = v4
    } else {
        for fi5 := 0; fi5 < 1; fi5++ {
            if false {
                if true {
                    v6 := []int{}
                    _ = v6
                } else {
                    var v7 int16 = 1000
                    _ = v7
                }
            }
            _ = fi5
        }
        fmt.Println(p0)
        if (0 != 42) {
            p0 = 0
            p1 = 42
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 bool = true
                v10 := GenS2{F0: (p1 % (((p1 >> ((4294967295) & 63))) | 1)), F1: 4294967295}
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            p0 = p0
            if true {
                var v11 int64 = 9223372036854775807
                v11 = (v11 | v11)
                var v12 float64 = 3.14159
                _ = v11
                _ = v12
            }
        }
    }
    fmt.Println(p0, p2, p1)
    v13 := GenS2{F0: (p1 % (((p1 / ((p1) | 1))) | 1)), F1: 62}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return p1
}
func genFn1(p0 uint32, p1 int64) string {
    v14 := GenS3{F0: 0, F1: -2147483648}
    p0 = 293
    _ = p0
    _ = p1
    _ = v14
    return "hello"
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{}
    var v3 uint8 = 1
    if false {
        v3 = v3
        for fi15 := 0; fi15 < 5; fi15++ {
            if false {
                var v16 string = "\n\t"
                v17 := GenS2{F0: v3, F1: 134}
                _ = v16
                _ = v17
            } else {
                var v18 int16 = 1
                _ = v18
            }
            if true {
                v2 = append(v2, -2147483648)
                switch fi15 {
                case -309:
                    if len(v2) > 0 {
                        v2[0] = fi15
                    }
                default:
                    switch fi15 {
                    case 32768:
                        v19 := GenS3{F0: 0, F1: -2147483648}
                        _ = v19
                    default:
                        switch 1 {
                        case 520:
                            if (!true) {
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    v3 -= genFn0(127, v3, 42)
                                    _ = fi20
                                }
                                v21 := GenS1{F0: -1, F1: 1, F2: 1}
                                _ = v21
                            } else {
                                switch (fi15 >> ((244) & 63)) {
                                case -2147483648:
                                    v3 += (v3 + v3)
                                case 4294967295:
                                    if (0 < 18446744073709551615) {
                                        fmt.Println(fi15)
                                        v22 := GenS3{F0: 2, F1: -1}
                                        v23 := GenS2{F0: v3, F1: 822}
                                        _ = v22
                                        _ = v23
                                    }
                                }
                                for fi24 := 0; fi24 < 4; fi24++ {
                                    switch fi15 {
                                    case 4294967295:
                                        _ = len(v2)
                                    case -2147483648:
                                        if (1 == 124) {
                                            switch "" {
                                            case "foo bar":
                                                fmt.Println(fi24, fi24, fi15)
                                            default:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            if ("foo bar" <= ("\n\t" + ("\n\t" + ("a" + (("a" + "foo bar") + ("\n\t" + "")))))) {
                                                v3 = genFn0(32768, 42, 65535)
                                            } else {
                                                delete(v1, (fi15 * (fi15 + fi24)))
                                            }
                                        } else {
                                            var v25 uint32 = 2
                                            _ = v25
                                        }
                                    default:
                                        fmt.Println(fi15, v3)
                                    }
                                    _ = fi24
                                }
                            }
                        case 128:
                            switch 128 {
                            case 127:
                                v3 += v3
                            case -224:
                                var v26 uint32 = 4294967295
                                _ = v26
                            case 32768:
                                var v27 int16 = -1
                                _ = v27
                            default:
                                if (!(!true)) {
                                    if true {
                                        switch ("a" + "\n\t") {
                                        case "foo bar":
                                            v3 -= genFn0(fi15, (v3 % (((v3 % ((163) | 1))) | 1)), 4294967295)
                                        }
                                    } else {
                                        var v28 int32 = 42
                                        _ = v28
                                    }
                                    var v29 bool = false
                                    v29 = (127 > -1)
                                    _ = v29
                                } else {
                                    var v30 bool = ((("foo bar" != "") || true) || true)
                                    var v31 float32 = 0.1
                                    _ = v30
                                    _ = v31
                                }
                            }
                        default:
                            v3 += 255
                        }
                    }
                }
                v3 -= (v3 / (((^(v3))) | 1))
            }
            v2 = append(v2, fi15)
            _ = fi15
        }
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        var v33 int8 = 42
        _ = fi32
        _ = v33
    }
    fmt.Println(v3)
    var v34 int64 = 42
    if true {
        var v35 uint64 = 783
        switch -577 {
        case -2147483648:
            if true {
                for fi36 := 0; fi36 < 13; fi36++ {
                    if ((^(v35)) == (^(v35))) {
                        delete(v1, int(v3))
                        switch ("a" + "foo bar") {
                        case "hello":
                            for fi37 := 0; fi37 < 6; fi37++ {
                                v38 := v1[fi37]
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch "" {
                                case "a":
                                    switch v38 {
                                    case -2147483648:
                                        var v39 uint8 = v3
                                        _ = v39
                                    case 9223372036854775807:
                                        v40 := v1[(v38 & 128)]
                                        _ = v40
                                    case 521:
                                        delete(v1, (fi37 | fi36))
                                    }
                                case "\n\t":
                                    fmt.Println(fi36, v35, v35)
                                case "hello":
                                    v3 += 1
                                }
                                _ = fi37
                                _ = v38
                            }
                        }
                        var v41 int8 = -1
                        _ = v41
                    } else {
                        v3 -= v3
                        if true {
                            var v42 int16 = 42
                            v34 = (v34 * 1)
                            _ = v42
                        }
                        v3 -= v3
                    }
                    fmt.Println(fi36)
                    v35 = 428
                    _ = fi36
                }
                var v43 int8 = -1
                _ = v43
            }
        }
        _ = v35
    } else {
        for fi44 := 0; fi44 < 1; fi44++ {
            if true {
                _ = len(v2)
                if (0 == (v34 | v34)) {
                    for fi45 := 0; fi45 < 8; fi45++ {
                        var v46 int16 = 32767
                        _ = fi45
                        _ = v46
                    }
                    for fi47 := 0; fi47 < 9; fi47++ {
                        var v48 int = int(v3)
                        _ = fi47
                        _ = v48
                    }
                    if true {
                        var v49 uint8 = v3
                        v50 := GenS4{F0: -1, F1: 904, F2: 740}
                        if true {
                            if (!((false || true) || false)) {
                                v51 := GenS4{F0: 0, F1: uint64(v49), F2: (v34 * (v34 | (v34 % (((v34 | v34)) | 1))))}
                                fmt.Println(v34, v3, v49)
                                _ = v51
                            } else {
                                v52 := GenS2{F0: genFn0(fi44, (v49 + (^(v3))), uint32(v3)), F1: 586}
                                switch 906 {
                                case -920:
                                    var v53 uint = 65535
                                    _ = v53
                                }
                                _ = v52
                            }
                            _ = len(v2)
                        }
                        _ = v49
                        _ = v50
                    } else {
                        v54 := GenS2{F0: 0, F1: 0}
                        _ = v54
                    }
                }
            }
            _ = fi44
        }
    }
    v3 = (^(v3))
    v55 := v1[32768]
    var v56 int = (v55 % ((v55) | 1))
    for fi57 := 0; fi57 < 9; fi57++ {
        _ = len(v2)
        var v58 uint8 = (v3 & v3)
        _ = fi57
        _ = v58
    }
    for fi59 := 0; fi59 < 4; fi59++ {
        switch "a" {
        case "hello":
            for fi60 := 0; fi60 < 9; fi60++ {
                v61 := GenS4{F0: -459, F1: 0, F2: v34}
                if (!((true && ((-1 < 1) || false)) || false)) {
                    var v62 int64 = v34
                    var v63 float32 = -2.25
                    _ = v62
                    _ = v63
                } else {
                    switch "hello" {
                    case "foo bar":
                        v64 := GenS5{F0: -505, F1: 65}
                        _ = v64
                    }
                    if ((("foo bar" + ("hello" + "foo bar")) + "\n\t") >= ("" + "hello")) {
                        var v65 uint32 = 773
                        _ = v65
                    }
                    v3 = (v3 / ((128) | 1))
                }
                var v66 float64 = float64(0.1)
                _ = fi60
                _ = v61
                _ = v66
            }
        case "a":
            for fi67 := 0; fi67 < 6; fi67++ {
                switch -2147483648 {
                case -746:
                    var v68 uint = 42
                    _ = v68
                case 127:
                    v55 -= v55
                default:
                    var v69 float32 = 0.1
                    _ = v69
                }
                _ = fi67
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi70 := 0; fi70 < 9; fi70++ {
            if true {
                var v71 int16 = 126
                _ = v71
            }
            switch ("" + "") {
            case "a":
                fmt.Println(v55, fi59)
            case "hello":
                v34 = -9223372036854775808
            case "\n\t":
                var v72 float32 = (1.5 - 1e10)
                var v73 int = (v56 | -903)
                if (42 != 17) {
                    v1[-410] = (v56 * v56)
                    v56 = v73
                    var v74 uint64 = uint64(v3)
                    _ = v74
                } else {
                    if (!((107 == 278) && true)) {
                        var v75 float32 = v72
                        var v76 int16 = 32767
                        v3 = v3
                        _ = v75
                        _ = v76
                    }
                }
                if (!true) {
                    if (!true) {
                        if (true || false) {
                            var v77 string = ("\n\t" + "hello")
                            if true {
                                switch v77 {
                                case "foo bar":
                                    fmt.Println(v34)
                                default:
                                    delete(v1, -1)
                                }
                                switch (v77 + "hello") {
                                case "":
                                    v78 := GenS4{F0: 2147483647, F1: 1, F2: v34}
                                    _ = v78
                                case "\n\t":
                                    v56 += (fi59 << ((764) & 63))
                                default:
                                    v79 := GenS2{F0: v3, F1: 1}
                                    _ = v79
                                }
                                v72 = (v72 / v72)
                            } else {
                                v34 = (-(v34))
                                var v80 float64 = (1e10 - 1.5)
                                _ = v80
                            }
                            var v81 int8 = 101
                            var v82 int16 = 32767
                            var v83 int32 = 202
                            _ = v77
                            _ = v81
                            _ = v82
                            _ = v83
                        } else {
                            var v84 int = (fi59 ^ 127)
                            _ = v84
                        }
                        v85 := v1[(fi59 * v56)]
                        _ = v85
                    }
                    delete(v1, v56)
                    if (42 >= -515) {
                        fmt.Println(v3)
                        v3 -= 128
                        _ = len(v2)
                    }
                }
                v86 := GenS6{F0: v3, F1: v3}
                var v87 int = v73
                _ = v72
                _ = v73
                _ = v86
                _ = v87
            }
            var v88 string = ("a" + (("a" + (("" + "a") + "hello")) + ""))
            _ = fi70
            _ = v88
        }
        _ = fi59
    }
    var v89 int16 = 0
    for fi90 := 0; fi90 < 18; fi90++ {
        switch "\n\t" {
        case "a":
            v91 := GenS5{F0: 978, F1: -1}
            _ = v91
        case "foo bar":
            if true {
                v34 = (v34 - v34)
            }
        case "\n\t":
            v92 := GenS4{F0: -433, F1: 4294967295, F2: v34}
            _ = v92
        }
        v89 = (^(v89))
        for fi93 := 0; fi93 < 4; fi93++ {
            var v94 int64 = (v34 * (v34 >> ((871) & 63)))
            _ = fi93
            _ = v94
        }
        _ = fi90
    }
    var v95 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v89:", v89)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v55
    _ = v56
    _ = v89
    _ = v95
}
