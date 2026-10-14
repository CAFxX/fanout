package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint; F1 int32; F2 uint }
type GenS6 struct { F0 int16; F1 string; F2 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS9 struct { F0 int32; F1 bool; F2 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint16 }
func genFn0(p0 string, p1 string, p2 bool) uint32 {
    var v6 uint64 = 172
    var v7 uint32 = 1
    v8 := GenS1{F0: -37, F1: 951, F2: v7}
    v9 := []int{1, 0, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 2147483648
}
func genFn1() uint32 {
    if (!(1.5 < 1.5)) {
        var v10 float32 = (1e10 + 1.5)
        var v11 uint64 = 1
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 1, F1: 357, F2: 793}
    _ = v12
    return 42
}

func main() {
    var v1 int = 678
    v2 := make(map[string]int)
    var v3 uint32 = 58
    v4 := []uint32{(^(v3))}
    var v5 uint = 325
    v1 -= (^(v1))
    var v13 uint8 = 2
    var v14 uint8 = (v13 & v13)
    _ = len(v2)
    for fi15 := 0; fi15 < 15; fi15++ {
        fmt.Println(v14, v13)
        _ = fi15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        v17 := GenS2{F0: v5, F1: 38, F2: 0}
        var v18 float32 = 0.0
        fmt.Println(v3, v13)
        _ = fi16
        _ = v17
        _ = v18
    }
    if (v1 >= v1) {
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 int8 = 1
            var v21 uint16 = 42
            for fi22 := 0; fi22 < 12; fi22++ {
                if (0.1 < ((0.1 * 3.14159) / 0.1)) {
                    v3 = v3
                }
                v5 = v5
                _ = fi22
            }
            _ = fi19
            _ = v20
            _ = v21
        }
        if true {
            switch (v1 << ((42) & 63)) {
            case 32768:
                v4 = append(v4, 200)
            case 420:
                fmt.Println(v13)
            default:
                var v23 float32 = (-2.25 / -2.25)
                _ = v23
            }
            switch 2147483647 {
            case 325:
                _ = len(v2)
            case 177:
                var v24 int8 = 126
                _ = v24
            default:
                var v25 uint64 = 42
                _ = v25
            }
            fmt.Println(v5, v3)
        } else {
            switch (v1 & v1) {
            case 0:
                for fi26 := 0; fi26 < 1; fi26++ {
                    v27 := GenS3{F0: 32767}
                    for fi28 := 0; fi28 < 7; fi28++ {
                        v29 := GenS4{F0: 494, F1: 0}
                        if (!(!true)) {
                            var v30 int32 = 42
                            v31 := GenS5{F0: 4294967295, F1: 42, F2: v5}
                            _ = v30
                            _ = v31
                        }
                        var v32 int = (fi28 & (v1 | fi26))
                        v33 := GenS3{F0: 1}
                        _ = fi28
                        _ = v29
                        _ = v32
                        _ = v33
                    }
                    _ = fi26
                    _ = v27
                }
            case 127:
                if true {
                    var v34 int = (-(v1))
                    switch v1 {
                    case 32767:
                        var v35 uint32 = genFn1()
                        _ = v35
                    case 155:
                        if true {
                            v36 := GenS1{F0: 12, F1: int32(v5), F2: v3}
                            if (!true) {
                                switch "foo bar" {
                                case "hello":
                                    if (598 != 0) {
                                        var v37 uint = (v5 ^ uint(v14))
                                        if len(v4) > 0 {
                                            v4[0] = (v3 / ((uint32(v5)) | 1))
                                        }
                                        if false {
                                            var v38 float64 = 1e10
                                            _ = v38
                                        }
                                        _ = v37
                                    }
                                }
                            }
                            var v39 string = ""
                            _ = v36
                            _ = v39
                        }
                    default:
                        v40 := GenS6{F0: -1, F1: ("hello" + "foo bar"), F2: (v14 << (((v5 >> (((v5 - v5)) & 63))) & 63))}
                        _ = v40
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v34
                }
            }
            for fi41 := 0; fi41 < 15; fi41++ {
                var v42 int64 = 1
                v5 = uint(v3)
                var v43 int64 = (v42 >> ((0) & 63))
                _ = fi41
                _ = v42
                _ = v43
            }
        }
        if (!true) {
            var v44 uint = v5
            _ = v44
        } else {
            if (((v5 / (((v5 % ((v5) | 1))) | 1)) >= 65535) || (!(!true))) {
                if (!false) {
                    v3 = genFn1()
                    if true {
                        v3 = v3
                        if (!(-1 < 441)) {
                            var v45 int = v1
                            var v46 uint8 = 42
                            v47 := GenS7{F0: 80}
                            _ = v45
                            _ = v46
                            _ = v47
                        } else {
                            if (((v13 >> ((v5) & 63)) != 29) || true) {
                                v48 := GenS7{F0: v5}
                                var v49 int8 = -115
                                if ("\n\t" >= (((("hello" + "") + (("hello" + "a") + ((("\n\t" + "foo bar") + "foo bar") + "a"))) + (("a" + "foo bar") + "foo bar")) + "foo bar")) {
                                    v14 = v13
                                    v3 -= v3
                                } else {
                                    v14 = v14
                                }
                                _ = v48
                                _ = v49
                            }
                        }
                        if false {
                            fmt.Println(v3, v1, v13)
                        } else {
                            var v50 float64 = ((0.1 * ((1e10 + 1e10) - -2.25)) / 3.14159)
                            switch v1 {
                            case 32768:
                                v1 -= int(v3)
                            case 127:
                                var v51 int16 = -1
                                _ = v51
                            default:
                                v2[("" + ("" + "hello"))] = (v1 * (-(v1)))
                            }
                            _ = v50
                        }
                    } else {
                        v52 := GenS1{F0: 100, F1: 1, F2: v3}
                        _ = v52
                    }
                    var v53 uint64 = 18446744073709551615
                    _ = v53
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi54 := 0; fi54 < 18; fi54++ {
                    if false {
                        if false {
                            var v55 int = (fi54 * fi54)
                            var v56 float32 = ((-2.25 / 0.1) - (-2.25 + (3.14159 - (-2.25 / ((0.1 / (1.5 + 0.0)) * (1e10 / ((3.14159 + (1.5 + 0.1)) + 0.0)))))))
                            for fi57 := 0; fi57 < 6; fi57++ {
                                var v58 float64 = 1.5
                                _ = fi57
                                _ = v58
                            }
                            _ = v55
                            _ = v56
                        } else {
                            v59 := GenS2{F0: 0, F1: v3, F2: (v5 >> (((v5 | 18446744073709551615)) & 63))}
                            v13 -= (^(v14))
                            _ = v59
                        }
                        var v60 float64 = (float64(3.14159) * 0.1)
                        _ = v60
                    }
                    if len(v4) > 0 {
                        v4[0] = (v3 << ((v5) & 63))
                    }
                    switch (("hello" + "") + ("a" + "a")) {
                    case "\n\t":
                        if len(v4) > 0 {
                            v4[0] = (v3 / ((0) | 1))
                        }
                    case "":
                        for fi61 := 0; fi61 < 12; fi61++ {
                            if (false || ((!(!(true || (false || false)))) || true)) {
                                v13 = (v14 + (^(v13)))
                                v4 = append(v4, v3)
                                v62 := GenS8{F0: (v3 | (v3 - v3)), F1: (v3 & 42), F2: 42}
                                _ = v62
                            }
                            if false {
                                v63 := GenS8{F0: uint32(v5), F1: genFn0("", "foo bar", true), F2: -1}
                                if true {
                                    fmt.Println(v13, v13)
                                    fmt.Println(v5, v14, v14)
                                } else {
                                    v3 += (v3 | (^(v3)))
                                }
                                _ = v63
                            } else {
                                fmt.Println(v5)
                                v4 = append(v4, (v3 >> ((255) & 63)))
                            }
                            v64 := v2["\n\t"]
                            _ = fi61
                            _ = v64
                        }
                    }
                    _ = fi54
                }
            }
            if (255 > v14) {
                if (v1 == (v1 ^ -2147483648)) {
                    var v65 uint64 = 2
                    _ = len(v4)
                    _ = v65
                } else {
                    v66 := GenS9{F0: 0, F1: (!false), F2: 65535}
                    if (!(v3 != (^(v3)))) {
                        v4 = append(v4, (v3 >> (((v5 - v5)) & 63)))
                        var v67 float64 = 0.0
                        for fi68 := 0; fi68 < 19; fi68++ {
                            var v69 bool = (-1 != 1)
                            v69 = true
                            _ = fi68
                            _ = v69
                        }
                        _ = v67
                    }
                    _ = v66
                }
                for fi70 := 0; fi70 < 18; fi70++ {
                    v1 -= 1
                    if (!(2147483647 > -1)) {
                        fmt.Println(v5, fi70, fi70)
                    }
                    v3 += genFn0(("foo bar" + "foo bar"), "", (70 != 71))
                    _ = fi70
                }
                if (false || true) {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            } else {
                if (false || (!true)) {
                    v71 := GenS2{F0: (v5 % (((v5 | v5)) | 1)), F1: v3, F2: v5}
                    for fi72 := 0; fi72 < 20; fi72++ {
                        v73 := GenS10{F0: true}
                        var v74 bool = ((!false) || true)
                        var v75 int64 = -9223372036854775808
                        _ = fi72
                        _ = v73
                        _ = v74
                        _ = v75
                    }
                    _ = v71
                } else {
                    fmt.Println(v5, v5)
                }
                _ = len(v4)
            }
            for fi76 := 0; fi76 < 6; fi76++ {
                var v77 bool = false
                switch "hello" {
                case "a":
                    for fi78 := 0; fi78 < 12; fi78++ {
                        v77 = (-844 == int16(v5))
                        v79 := GenS4{F0: -574, F1: -851}
                        fmt.Println(v13, v77)
                        _ = fi78
                        _ = v79
                    }
                case "hello":
                    for fi80 := 0; fi80 < 2; fi80++ {
                        v81 := GenS1{F0: -1, F1: -1, F2: (v3 >> ((995) & 63))}
                        v82 := GenS11{F0: -363}
                        v13 = v13
                        var v83 uint64 = 18446744073709551615
                        v84 := GenS12{F0: 32767}
                        var v85 uint16 = 989
                        var v86 uint32 = v3
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v87 int16 = int16(v86)
                        if false {
                            v5 += v5
                        }
                        _ = fi80
                        _ = v81
                        _ = v82
                        _ = v83
                        _ = v84
                        _ = v85
                        _ = v86
                        _ = v87
                    }
                case "foo bar":
                    delete(v2, (("foo bar" + "\n\t") + ("foo bar" + ("" + "foo bar"))))
                default:
                    v77 = false
                }
                _ = fi76
                _ = v77
            }
        }
    } else {
        var v88 int8 = -19
        switch v1 {
        case -772:
            var v89 float32 = (((-2.25 + 1.5) - 0.1) * 1.5)
            _ = v89
        case 9223372036854775807:
            for fi90 := 0; fi90 < 3; fi90++ {
                _ = len(v4)
                if (0 == 304) {
                    var v91 int8 = v88
                    _ = v91
                }
                v3 -= 2
                _ = fi90
            }
        default:
            if true {
                _ = len(v4)
                var v92 string = "hello"
                var v93 uint64 = 456
                _ = v92
                _ = v93
            } else {
                fmt.Println(v13, v3)
                v94 := GenS5{F0: v5, F1: 0, F2: v5}
                _ = v94
            }
        }
        _ = v88
    }
    v95 := GenS13{F0: 369}
    var v96 int64 = 42
    v96 += -969
    v97 := GenS11{F0: 42}
    v4 = append(v4, (v3 / ((v3) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v95
    _ = v96
    _ = v97
}
