package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16; F1 int8 }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int16; F1 bool; F2 uint8 }
type GenS11 struct { F0 uint16 }
func genFn0() int32 {
    var v4 int = 32767
    v4 = v4
    var v5 uint32 = 2147483647
    v6 := []int{v4}
    fmt.Println(v4, v4)
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int8, p1 bool) uint16 {
    var v7 float64 = 1.5
    var v8 int16 = 892
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := []int{fi9}
            _ = fi10
            _ = v11
        }
        v12 := GenS1{F0: "foo bar", F1: 65535, F2: 128}
        if p1 {
            v13 := GenS1{F0: (("" + "") + "a"), F1: 42, F2: 42}
            for fi14 := 0; fi14 < 4; fi14++ {
                v15 := GenS2{F0: 1, F1: (^(p0))}
                v7 -= 0.1
                var v16 int16 = (-(v8))
                v17 := GenS1{F0: (("" + ("hello" + "foo bar")) + "foo bar"), F1: 42, F2: 128}
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            fmt.Println(v8)
            _ = v13
        }
        _ = fi9
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{2}
    var v3 int8 = 0
    if false {
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 int8 = 0
            v20 := GenS3{F0: 0}
            _ = fi18
            _ = v19
            _ = v20
        }
        v3 += (v3 % ((70) | 1))
    } else {
        fmt.Println(v3)
        if true {
            var v21 int = 42
            var v22 uint = 0
            for fi23 := 0; fi23 < 10; fi23++ {
                var v24 bool = (0.1 > float32(1.5))
                _ = fi23
                _ = v24
            }
            _ = v21
            _ = v22
        } else {
            switch 1 {
            case -1:
                for fi25 := 0; fi25 < 8; fi25++ {
                    switch (("hello" + "hello") + ("foo bar" + "\n\t")) {
                    case "hello":
                        v3 += 100
                    default:
                        switch fi25 {
                        case 1:
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                        case 684:
                            var v26 uint = 18446744073709551615
                            _ = v26
                        default:
                            switch "a" {
                            case "":
                                switch "foo bar" {
                                case "foo bar":
                                    v3 = v3
                                case "a":
                                    var v27 int32 = -238
                                    _ = v27
                                case "":
                                    v3 = v3
                                }
                            case "\n\t":
                                _ = len(v2)
                            default:
                                switch (((("a" + "\n\t") + "hello") + "") + "hello") {
                                case "a":
                                    var v28 bool = ("\n\t" > (("a" + "") + ((("foo bar" + ("\n\t" + "foo bar")) + "") + ("a" + ""))))
                                    _ = v28
                                default:
                                    v29 := GenS1{F0: "hello", F1: 854, F2: 127}
                                    _ = v29
                                }
                            }
                        }
                    }
                    switch "hello" {
                    case "\n\t":
                        for fi30 := 0; fi30 < 7; fi30++ {
                            var v31 uint16 = 424
                            v3 = (v3 & (v3 << ((0) & 63)))
                            _ = fi30
                            _ = v31
                        }
                    case "hello":
                        v3 = v3
                    case "foo bar":
                        v3 += (v3 ^ v3)
                    default:
                        if false {
                            for fi32 := 0; fi32 < 17; fi32++ {
                                _ = len(v2)
                                var v33 uint64 = 2
                                _ = fi32
                                _ = v33
                            }
                            v34 := GenS1{F0: ("hello" + ""), F1: 2, F2: 128}
                            _ = len(v1)
                            _ = v34
                        }
                    }
                    v3 = (v3 >> ((255) & 63))
                    _ = fi25
                }
            case 884:
                for fi35 := 0; fi35 < 13; fi35++ {
                    switch (fi35 % ((fi35) | 1)) {
                    case -1:
                        var v36 int = fi35
                        _ = v36
                    case 782:
                        switch (fi35 + (fi35 << ((18446744073709551615) & 63))) {
                        case 903:
                            if len(v2) > 0 {
                                v2[0] = 2147483647
                            }
                        default:
                            delete(v1, 1)
                        }
                    case 32767:
                        var v37 float32 = 0.0
                        _ = v37
                    default:
                        if (0 < -759) {
                            if (false && (-1 > -126)) {
                                var v38 int32 = genFn0()
                                for fi39 := 0; fi39 < 4; fi39++ {
                                    v40 := GenS1{F0: ("foo bar" + "a"), F1: 186, F2: 0}
                                    fmt.Println(fi39, fi35, fi39)
                                    _ = fi39
                                    _ = v40
                                }
                                _ = v38
                            } else {
                                for fi41 := 0; fi41 < 8; fi41++ {
                                    var v42 uint16 = 255
                                    fmt.Println(v42, v3)
                                    switch (-(fi41)) {
                                    case 279:
                                        if (128 == 42) {
                                            v43 := GenS4{F0: 906, F1: false}
                                            fmt.Println(v3)
                                            v44 := v1[0]
                                            var v45 int = 32768
                                            v1[2] = fi41
                                            v42 += v42
                                            switch ("\n\t" + "foo bar") {
                                            case "a":
                                                v46 := GenS5{F0: 18446744073709551615}
                                                _ = v46
                                            case "foo bar":
                                                v45 = -968
                                            }
                                            v45 = (^(v45))
                                            v47 := GenS6{F0: v42, F1: (v3 >> ((1) & 63))}
                                            var v48 bool = false
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                            _ = v47
                                            _ = v48
                                        } else {
                                            if false {
                                                v2 = append(v2, 382)
                                                v49 := GenS7{F0: 0, F1: (fi35 / (((^(fi35))) | 1))}
                                                delete(v1, 291)
                                                if ((!true) && false) {
                                                    var v50 int8 = v3
                                                    v42 += 65535
                                                    v3 += v50
                                                    if (65535 <= 42) {
                                                        if true {
                                                            if true {
                                                                v51 := GenS2{F0: (^(v50)), F1: v3}
                                                                var v52 uint = 175
                                                                _ = v51
                                                                _ = v52
                                                            } else {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    fmt.Println(v42, v3)
                                                                case "a":
                                                                    var v53 uint8 = 1
                                                                    if (((!(!false)) && true) || true) {
                                                                        fmt.Println(v42, v53)
                                                                        fmt.Println(v53, fi35)
                                                                    }
                                                                    _ = v53
                                                                default:
                                                                    var v54 int32 = -2147483648
                                                                    _ = v54
                                                                }
                                                                v55 := GenS8{F0: 42, F1: -1}
                                                                v3 -= int8(fi41)
                                                                _ = v55
                                                            }
                                                            fmt.Println(fi41)
                                                        } else {
                                                            if (1 < 1) {
                                                                v50 = v3
                                                                v56 := GenS6{F0: (v42 % ((v42) | 1)), F1: v50}
                                                                _ = v56
                                                            } else {
                                                                fmt.Println(fi41, fi35)
                                                                v42 = (^(v42))
                                                            }
                                                        }
                                                    }
                                                    _ = v50
                                                }
                                                fmt.Println(fi35, v3, fi35)
                                                _ = v49
                                            } else {
                                                if false {
                                                    if (false || (int16(v3) <= -484)) {
                                                        var v57 int8 = -21
                                                        var v58 float32 = 0.0
                                                        v42 = (^(v42))
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        fmt.Println(v3, v42, fi41)
                                                        if (((float64(1.5) + -2.25) / 1.0) != -2.25) {
                                                            switch -946 {
                                                            case 4294967295:
                                                                var v59 int = fi41
                                                                _ = v59
                                                            case 128:
                                                                if false {
                                                                    fmt.Println(v42)
                                                                    var v60 uint8 = 127
                                                                    _ = v60
                                                                } else {
                                                                    if false {
                                                                        if true {
                                                                            v61 := GenS5{F0: 4294967295}
                                                                            _ = v61
                                                                        }
                                                                    }
                                                                    if (171 == 1000) {
                                                                        var v62 int16 = -32768
                                                                        v1[4294967295] = -544
                                                                        fmt.Println(v3)
                                                                        _ = v62
                                                                    } else {
                                                                        v42 = (v42 % ((v42) | 1))
                                                                    }
                                                                }
                                                            case 0:
                                                                v3 = v3
                                                            }
                                                        } else {
                                                            fmt.Println(fi41, fi35, fi35)
                                                            v63 := GenS9{F0: 1}
                                                            switch (fi35 << ((65535) & 63)) {
                                                            case 127:
                                                                v3 = 1
                                                                v3 += v3
                                                            case 2147483647:
                                                                var v64 float64 = 0.1
                                                                _ = v64
                                                            case 32768:
                                                                v65 := GenS1{F0: "foo bar", F1: v42, F2: 131}
                                                                _ = v65
                                                            }
                                                            _ = v63
                                                        }
                                                    }
                                                    fmt.Println(fi35, v3)
                                                    _ = len(v2)
                                                } else {
                                                    var v66 int64 = 9223372036854775807
                                                    var v67 int16 = 0
                                                    if false {
                                                        if true {
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                v66 = v66
                                                            default:
                                                                v42 += v42
                                                            }
                                                        } else {
                                                            var v68 float32 = 0.1
                                                            var v69 uint32 = 2
                                                            fmt.Println(v68, v3)
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                        fmt.Println(fi35)
                                                    } else {
                                                        _ = len(v2)
                                                        fmt.Println(fi35, fi41)
                                                        if true {
                                                            if true {
                                                                v1[uint32(v42)] = 573
                                                            }
                                                        } else {
                                                            v70 := GenS6{F0: (v42 & genFn1((v3 / (((-(v3))) | 1)), false)), F1: (-(v3))}
                                                            v67 = (^(v67))
                                                            var v71 int32 = 2147483647
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = v66
                                                    _ = v67
                                                }
                                            }
                                            var v72 int64 = 9223372036854775807
                                            if false {
                                                v72 -= (v72 | (v72 * (v72 % (((v72 - (v72 >> ((18446744073709551615) & 63)))) | 1))))
                                            } else {
                                                _ = len(v2)
                                            }
                                            _ = v72
                                        }
                                    case 42:
                                        var v73 bool = false
                                        _ = v73
                                    case 9223372036854775807:
                                        switch "foo bar" {
                                        case "a":
                                            fmt.Println(fi41)
                                        case "":
                                            v74 := GenS8{F0: -2147483648, F1: -1}
                                            _ = v74
                                        case "foo bar":
                                            var v75 int8 = (v3 / ((v3) | 1))
                                            _ = v75
                                        default:
                                            var v76 bool = ("hello" == ("\n\t" + ""))
                                            _ = v76
                                        }
                                    default:
                                        var v77 string = "a"
                                        _ = v77
                                    }
                                    _ = fi41
                                    _ = v42
                                }
                                v3 = 36
                            }
                            v78 := GenS10{F0: 0, F1: true, F2: 181}
                            _ = v78
                        }
                    }
                    _ = fi35
                }
            }
            fmt.Println(v3)
        }
    }
    switch "a" {
    case "hello":
        v3 = -128
    default:
        v3 -= v3
    }
    var v79 uint16 = 1
    v79 -= 0
    v2 = append(v2, 65535)
    switch "" {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = 0
        }
    }
    v80 := v1[2147483648]
    if (!false) {
        if true {
            if (!false) {
                v80 = (v80 & v80)
            }
            delete(v1, 2)
        }
    }
    v80 += 128
    if (!true) {
        var v81 uint64 = 68
        fmt.Println(v81, v81)
        _ = v81
    } else {
        for fi82 := 0; fi82 < 14; fi82++ {
            fmt.Println(v80)
            v79 -= v79
            _ = fi82
        }
        v79 = (v79 | (v79 % ((32767) | 1)))
    }
    for fi83 := 0; fi83 < 9; fi83++ {
        var v84 uint16 = genFn1(v3, false)
        var v85 int64 = 0
        v86 := GenS8{F0: int32(v85), F1: v85}
        _ = fi83
        _ = v84
        _ = v85
        _ = v86
    }
    for fi87 := 0; fi87 < 20; fi87++ {
        v88 := GenS11{F0: 65535}
        switch ("\n\t" + ("\n\t" + "\n\t")) {
        case "\n\t":
            v80 = (^(v80))
        case "hello":
            if true {
                var v89 uint64 = 0
                fmt.Println(v80, v79)
                _ = v89
            }
        }
        var v90 uint8 = 42
        _ = fi87
        _ = v88
        _ = v90
    }
    v91 := v1[1]
    if (true || (870 != 430)) {
        var v92 float32 = 0.1
        fmt.Println(v92, v80)
        _ = v92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v79:", v79)
    fmt.Println("v80:", v80)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v79
    _ = v80
    _ = v91
}
