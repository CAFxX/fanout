package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 int16; F1 int16; F2 uint16 }
type GenS7 struct { F0 int; F1 int; F2 uint }
type GenS8 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS9 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS10 struct { F0 uint; F1 uint64 }
type GenS11 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS12 struct { F0 bool }
func genFn0(p0 uint8, p1 int64) int {
    var v5 float64 = (1.5 / 1.0)
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int = fi6
        v8 := GenS1{F0: (v7 + (v7 >> ((509) & 63))), F1: 300}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return 32767
}

func main() {
    var v1 uint = 42
    var v2 float32 = 1.5
    v3 := []string{}
    var v4 uint = (v1 / (((v1 << ((4294967295) & 63))) | 1))
    if (!((v1 < v1) || false)) {
        v2 = v2
    } else {
        v4 += v1
        v2 += -2.25
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    var v9 float32 = 1.5
    v10 := GenS1{F0: 1, F1: 18446744073709551615}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 float32 = float32(-2.25)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 string = ("foo bar" + "hello")
            var v15 uint8 = 214
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        for fi17 := 0; fi17 < 12; fi17++ {
            if (true && false) {
                _ = len(v3)
                v9 -= v2
            } else {
                var v18 uint8 = 1
                _ = v18
            }
            if true {
                switch (fi16 << (((^(v4))) & 63)) {
                case 127:
                    switch (fi17 * fi16) {
                    case -1:
                        v3 = append(v3, (("hello" + ("hello" + "")) + "hello"))
                    }
                case -9223372036854775808:
                    v9 = float32(-2.25)
                case 1:
                    var v19 uint32 = 192
                    _ = v19
                default:
                    switch "" {
                    case "":
                        fmt.Println(fi16, v2)
                    case "foo bar":
                        var v20 uint = 42
                        _ = v20
                    case "hello":
                        if (43 != 1) {
                            if (127 > 100) {
                                var v21 int8 = -1
                                _ = v21
                            } else {
                                switch (fi16 / ((9223372036854775807) | 1)) {
                                case 127:
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v4 = (v4 | (v1 / ((v1) | 1)))
                                    }
                                case 0:
                                    switch "a" {
                                    case "":
                                        v4 = (v1 ^ (v1 >> ((18446744073709551615) & 63)))
                                    case "foo bar":
                                        var v22 int32 = 42
                                        _ = v22
                                    default:
                                        switch 439 {
                                        case 776:
                                            v23 := GenS1{F0: fi16, F1: 1}
                                            _ = v23
                                        case -2147483648:
                                            v24 := GenS2{F0: -300}
                                            _ = v24
                                        case 32768:
                                            var v25 float32 = v2
                                            _ = v25
                                        default:
                                            var v26 int8 = 127
                                            _ = v26
                                        }
                                    }
                                    fmt.Println(v4)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v27 int16 = 42
                                        _ = v27
                                    case "\n\t":
                                        switch ("\n\t" + "hello") {
                                        case "foo bar":
                                            v3 = append(v3, ("" + ""))
                                        case "\n\t":
                                            switch fi17 {
                                            case 129:
                                                if false {
                                                    _ = len(v3)
                                                    var v28 uint16 = 32767
                                                    fmt.Println(v9, v4)
                                                    _ = v28
                                                }
                                            case 128:
                                                if (!(783 < 2147483648)) {
                                                    var v29 uint = (v4 / ((255) | 1))
                                                    _ = v29
                                                } else {
                                                    if true {
                                                        if true {
                                                            fmt.Println(v9)
                                                            var v30 int16 = 0
                                                            v4 -= 0
                                                            v9 = -2.25
                                                            v4 = 22
                                                            _ = v30
                                                        } else {
                                                            _ = len(v3)
                                                            var v31 uint8 = 143
                                                            if true {
                                                                v32 := GenS3{F0: 476, F1: 32767, F2: fi17}
                                                                v2 = v9
                                                                v1 = v4
                                                                _ = v32
                                                            } else {
                                                                v9 = float32(1e10)
                                                                v33 := GenS3{F0: 42, F1: 42, F2: int(v1)}
                                                                v2 -= ((0.0 / (v2 + ((v9 * -2.25) + ((v2 * (float32((0.1 / 1e10)) - ((v2 / (float32(float64(v2)) / 1e10)) / 1e10))) - v9)))) - ((0.0 / float32(1.5)) + v2))
                                                                _ = v33
                                                            }
                                                            fmt.Println(v1, v9, v9)
                                                            var v34 int = fi17
                                                            v3 = append(v3, ((("a" + "") + ("a" + "\n\t")) + "\n\t"))
                                                            _ = v31
                                                            _ = v34
                                                        }
                                                        var v35 uint8 = 111
                                                        _ = v35
                                                    }
                                                }
                                            }
                                        case "":
                                            v9 -= 3.14159
                                        }
                                    default:
                                        var v36 int = genFn0(195, 1)
                                        _ = v36
                                    }
                                default:
                                    if (574 != -1) {
                                        switch (fi17 << ((v4) & 63)) {
                                        case 42:
                                            var v37 string = ("a" + "\n\t")
                                            _ = v37
                                        case 128:
                                            if (203 >= 127) {
                                                if ("a" == ((("a" + "a") + "foo bar") + (("foo bar" + "") + ""))) {
                                                    if ((3.14159 - ((((((3.14159 * 1e10) - 1.5) * -2.25) + 1.5) + 0.1) * 1.5)) == 1.5) {
                                                        if true {
                                                            var v38 uint64 = 4294967295
                                                            var v39 float32 = (1e10 + float32(0.0))
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            v1 = 4294967295
                                                            switch ("\n\t" + "hello") {
                                                            case "hello":
                                                                v9 = float32(1e10)
                                                            default:
                                                                var v40 bool = ((-1 == -1) && true)
                                                                _ = v40
                                                            }
                                                            fmt.Println(v2, v2, v9)
                                                        }
                                                    } else {
                                                        v9 = v2
                                                        v41 := GenS4{F0: 1, F1: 0}
                                                        _ = v41
                                                    }
                                                } else {
                                                    v9 -= 0.0
                                                }
                                                v1 += 4294967295
                                                var v42 float32 = (v9 - 1.5)
                                                _ = v42
                                            } else {
                                                switch ("foo bar" + "a") {
                                                case "a":
                                                    if true {
                                                        var v43 float32 = v2
                                                        var v44 int = (fi17 | fi16)
                                                        _ = v43
                                                        _ = v44
                                                    } else {
                                                        v4 -= (v4 % ((v4) | 1))
                                                        var v45 uint16 = 255
                                                        _ = v45
                                                    }
                                                case "foo bar":
                                                    fmt.Println(v1, fi17, v4)
                                                }
                                                fmt.Println(v2, fi16, v1)
                                                v4 -= 254
                                            }
                                        default:
                                            v1 -= v4
                                        }
                                        var v46 uint8 = 127
                                        _ = v46
                                    }
                                }
                                if (v9 < v9) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    if false {
                                        var v47 float32 = v2
                                        _ = v47
                                    }
                                    var v48 int8 = -1
                                    v9 -= v2
                                    if (true || ((!(791 <= 888)) || (!(0 < 42)))) {
                                        v3 = append(v3, "hello")
                                        if (((3.14159 + (3.14159 - (3.14159 * (-2.25 - 0.0)))) + 1.5) < ((3.14159 / 0.1) * 0.0)) {
                                            _ = len(v3)
                                        }
                                        var v49 uint64 = 1
                                        _ = v49
                                    } else {
                                        _ = len(v3)
                                        var v50 uint64 = 2
                                        v51 := GenS3{F0: (v50 + uint64(v4)), F1: 388, F2: fi17}
                                        _ = v50
                                        _ = v51
                                    }
                                    v2 -= 0.0
                                    _ = v48
                                }
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v2, v1, v2)
                                case "":
                                    fmt.Println(fi17, v2)
                                case "\n\t":
                                    v9 -= v9
                                default:
                                    v52 := GenS5{F0: 127, F1: 2}
                                    _ = v52
                                }
                            }
                            v1 = 4294967295
                            if (1 != 80) {
                                fmt.Println(v9)
                                if true {
                                    fmt.Println(fi17)
                                    v9 -= 0.0
                                }
                                v2 = v2
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v2 -= (3.14159 + (float32((((0.1 - 0.1) - 1.5) - 1.5)) / ((v2 + v9) * v2)))
                                    v1 = (v4 % ((602) | 1))
                                }
                                fmt.Println(v4)
                            } else {
                                if true {
                                    var v53 uint = 0
                                    fmt.Println(v9, v4, v1)
                                    _ = v53
                                }
                                _ = len(v3)
                                v3 = append(v3, "\n\t")
                            }
                        } else {
                            var v54 int32 = -2147483648
                            _ = v54
                        }
                    default:
                        v4 = (v1 + (^(v4)))
                    }
                }
                switch "" {
                case "foo bar":
                    v2 = v2
                case "\n\t":
                    fmt.Println(v2, v1)
                }
            } else {
                if true {
                    switch genFn0(2, 0) {
                    case 168:
                        var v55 int32 = 771
                        _ = v55
                    case 1:
                        var v56 uint32 = 65535
                        _ = v56
                    default:
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                    }
                    var v57 int64 = -9223372036854775808
                    var v58 uint64 = 4294967295
                    _ = v57
                    _ = v58
                } else {
                    if false {
                        if false {
                            var v59 uint16 = 65535
                            var v60 uint64 = 42
                            _ = v59
                            _ = v60
                        } else {
                            var v61 string = "\n\t"
                            _ = v61
                        }
                        if (v9 > (1e10 * v9)) {
                            v9 += (0.0 / float32(((-2.25 * ((((((float64((v9 * (v2 + v9))) + (0.1 * 3.14159)) / 3.14159) + ((((3.14159 * -2.25) / 1.0) + 0.0) / (1e10 - 3.14159))) * float64((v9 / (1e10 / (float32(1.5) / v9))))) - (0.1 / 3.14159)) / 0.1)) * 0.0)))
                            var v62 uint64 = 668
                            _ = v62
                        }
                    }
                    if (!false) {
                        if (!false) {
                            if (false && false) {
                                if len(v3) > 0 {
                                    v3[0] = "\n\t"
                                }
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v63 int = genFn0(2, -231)
                            _ = v63
                        }
                        var v64 uint = (v1 * (v1 - (v4 >> ((v4) & 63))))
                        _ = v64
                    }
                }
            }
            _ = fi17
        }
        for fi65 := 0; fi65 < 6; fi65++ {
            if (0 == 4294967295) {
                _ = len(v3)
                v66 := GenS6{F0: 1000, F1: 42, F2: 65535}
                var v67 uint32 = 2
                switch "foo bar" {
                case "":
                    v68 := GenS7{F0: (fi65 + (fi16 / (((fi65 * (fi65 >> (((v4 / ((4294967295) | 1))) & 63)))) | 1))), F1: (fi65 & -2147483648), F2: (v1 | (v1 % ((v1) | 1)))}
                    _ = v68
                case "\n\t":
                    if true {
                        fmt.Println(fi65, fi16)
                    } else {
                        v69 := GenS8{F0: v4, F1: 0, F2: (v67 | (v67 / ((v67) | 1)))}
                        _ = v69
                    }
                    v9 -= v9
                    var v70 uint32 = v67
                    _ = v70
                case "hello":
                    var v71 int16 = 989
                    _ = v71
                default:
                    fmt.Println(v9)
                }
                v67 = (v67 >> (((v4 - 759)) & 63))
                v2 -= (v2 / float32(3.14159))
                _ = v66
                _ = v67
            } else {
                _ = len(v3)
                var v72 int8 = 57
                _ = v72
            }
            if false {
                if (!false) {
                    switch "hello" {
                    case "a":
                        if true {
                            var v73 int16 = 32767
                            v1 = 0
                            _ = v73
                        }
                    case "hello":
                        var v74 uint32 = 4294967295
                        _ = v74
                    default:
                        v75 := GenS9{F0: 42, F1: 4294967295, F2: 1}
                        v3 = append(v3, "\n\t")
                        switch genFn0(203, 1) {
                        case 132:
                            if (false || (32767 != 510)) {
                                if (24 <= 0) {
                                    if (!((!true) || true)) {
                                        var v76 string = "hello"
                                        _ = v76
                                    } else {
                                        var v77 uint = 976
                                        var v78 int = 0
                                        _ = v77
                                        _ = v78
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    fmt.Println(v4, fi65, fi16)
                                    var v79 float64 = 0.1
                                    v79 += (0.1 * v79)
                                    _ = v79
                                }
                            } else {
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(fi65)
                                case "":
                                    fmt.Println(v2)
                                    var v80 int8 = 73
                                    var v81 uint32 = 0
                                    _ = v80
                                    _ = v81
                                case "hello":
                                    if true {
                                        if (127 >= 1) {
                                            if (v4 == v4) {
                                                var v82 int8 = 126
                                                _ = v82
                                            }
                                            fmt.Println(fi65)
                                        } else {
                                            fmt.Println(v2)
                                            switch (-(fi65)) {
                                            case -30:
                                                v83 := GenS10{F0: (v1 + v1), F1: 588}
                                                _ = v83
                                            case 42:
                                                var v84 uint = v4
                                                _ = v84
                                            case 32767:
                                                v4 = 1
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = (((((("hello" + "hello") + "hello") + ("a" + "a")) + "foo bar") + "") + "a")
                                            }
                                        }
                                    } else {
                                        v4 += (v1 * (v1 | (v4 & v1)))
                                        switch ("hello" + ("foo bar" + (("foo bar" + ("\n\t" + "hello")) + "\n\t"))) {
                                        case "foo bar":
                                            v85 := GenS8{F0: v1, F1: 255, F2: 338}
                                            _ = v85
                                        case "hello":
                                            v9 = 1.5
                                        }
                                        v86 := GenS11{F0: v4, F1: uint8(v4), F2: 42}
                                        fmt.Println(v9, v1)
                                        _ = v86
                                    }
                                }
                                v9 += v9
                                v9 -= v2
                            }
                        default:
                            var v87 bool = false
                            _ = v87
                        }
                        _ = v75
                    }
                    v4 = 255
                } else {
                    if len(v3) > 0 {
                        v3[0] = ("foo bar" + "a")
                    }
                    v4 -= v4
                }
            }
            _ = fi65
        }
        _ = fi16
    }
    var v88 uint32 = 138
    v1 -= (^(v4))
    fmt.Println(v4, v1)
    var v89 int64 = 9223372036854775807
    var v90 bool = (1e10 != (0.0 - -2.25))
    v91 := GenS2{F0: 42}
    for fi92 := 0; fi92 < 20; fi92++ {
        if (!v90) {
            v93 := GenS6{F0: -32768, F1: 869, F2: 2}
            _ = v93
        } else {
            var v94 uint8 = 127
            for fi95 := 0; fi95 < 17; fi95++ {
                switch (fi95 >> (((v1 >> ((18446744073709551615) & 63))) & 63)) {
                case 127:
                    if (false && ((v88 - (v88 | (v88 % ((v88) | 1)))) >= 0)) {
                        v96 := GenS6{F0: 1000, F1: 1, F2: 32767}
                        _ = v96
                    }
                case -707:
                    v94 = (v94 + 93)
                case 2147483647:
                    if len(v3) > 0 {
                        v3[0] = (("foo bar" + "a") + ("hello" + "hello"))
                    }
                default:
                    var v97 uint16 = 0
                    _ = v97
                }
                _ = fi95
            }
            _ = v94
        }
        _ = fi92
    }
    v98 := GenS12{F0: (-16 <= -12)}
    fmt.Println(v9, v88, v89)
    fmt.Println(v89)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v9:", v9)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v88
    _ = v89
    _ = v9
    _ = v90
    _ = v91
    _ = v98
}
