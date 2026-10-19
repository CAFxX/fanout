package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint16; F1 int64; F2 int64 }
type GenS6 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 uint8; F2 bool }
type GenS9 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint64, p1 uint, p2 int8) int8 {
    v6 := GenS1{F0: 748}
    p2 = (p2 / (((p2 << ((2) & 63))) | 1))
    v7 := GenS1{F0: 65535}
    v8 := []int{-2147483648, 128, 2147483647}
    var v9 uint8 = 149
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return p2
}
func genFn1(p0 string, p1 int) string {
    p0 = ("\n\t" + p0)
    var v10 uint = 494
    for fi11 := 0; fi11 < 7; fi11++ {
        if ((p0 >= (("\n\t" + p0) + p0)) || false) {
            if (-1 == 934) {
                p1 -= -1
            }
            var v12 string = (p0 + p0)
            if (32767 == 243) {
                if (0 != -49) {
                    fmt.Println(fi11, v12)
                    for fi13 := 0; fi13 < 3; fi13++ {
                        var v14 uint32 = 4294967295
                        var v15 uint64 = 18446744073709551615
                        v14 = (^(v14))
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    var v16 int8 = -1
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        p0 = (("" + "") + "foo bar")
                        var v18 uint64 = 2
                        _ = fi17
                        _ = v18
                    }
                    p1 -= (fi11 / (((fi11 | fi11)) | 1))
                    v12 = (v12 + (v12 + (v12 + (v12 + (v12 + "\n\t")))))
                }
                if (p1 > (-(fi11))) {
                    if false {
                        var v19 float32 = float32(0.1)
                        if false {
                            switch 32767 {
                            case 138:
                                v20 := GenS2{F0: 1, F1: 2}
                                _ = v20
                            case -9223372036854775808:
                                switch v12 {
                                case "":
                                    v21 := GenS1{F0: 4294967295}
                                    _ = v21
                                case "a":
                                    fmt.Println(v12)
                                }
                            case 642:
                                fmt.Println(fi11, v19)
                            default:
                                v19 = v19
                            }
                            v22 := GenS3{F0: -32768}
                            v23 := GenS4{F0: 42, F1: -32768}
                            _ = v22
                            _ = v23
                        } else {
                            switch p1 {
                            case 128:
                                v24 := []int{fi11}
                                _ = v24
                            case -554:
                                var v25 bool = false
                                _ = v25
                            default:
                                v12 = "hello"
                            }
                            if (42 >= 1) {
                                p1 -= 669
                                switch "hello" {
                                case "":
                                    v26 := GenS4{F0: 2, F1: 1}
                                    _ = v26
                                }
                                var v27 float64 = 0.1
                                _ = v27
                            } else {
                                switch "hello" {
                                case "a":
                                    v10 -= (^(v10))
                                }
                                var v28 int32 = 2147483647
                                v19 -= v19
                                _ = v28
                            }
                            v10 += 1
                        }
                        v12 = (p0 + "\n\t")
                        _ = v19
                    }
                    for fi29 := 0; fi29 < 3; fi29++ {
                        v30 := []int{p1, (p1 ^ fi29), 32768, (p1 ^ (^(fi11)))}
                        if (255 < 128) {
                            var v31 uint8 = 12
                            if len(v30) > 0 {
                                v30[0] = fi29
                            }
                            _ = v31
                        } else {
                            if len(v30) >= 2 {
                                v30 = v30[1:len(v30)]
                            }
                        }
                        v32 := GenS1{F0: 2147483647}
                        _ = fi29
                        _ = v30
                        _ = v32
                    }
                    fmt.Println(p1, p1)
                } else {
                    var v33 float64 = 0.0
                    _ = v33
                }
                if (-1 == -112) {
                    var v34 uint16 = 1
                    _ = v34
                }
            } else {
                v10 = (v10 % ((v10) | 1))
            }
            _ = v12
        } else {
            var v35 int = (^(fi11))
            var v36 bool = false
            _ = v35
            _ = v36
        }
        p0 = "\n\t"
        fmt.Println(p0, p0)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return p0
}

func main() {
    var v1 int = 32768
    var v2 uint8 = 127
    v3 := make(map[string]string)
    var v4 float64 = 1.5
    var v5 uint32 = 2
    for fi37 := 0; fi37 < 16; fi37++ {
        v5 += 0
        _ = fi37
    }
    var v38 uint64 = 257
    v39 := GenS1{F0: 65535}
    if true {
        v1 -= int(v2)
        if (false && (v1 >= v1)) {
            var v40 float64 = v4
            var v41 uint64 = v38
            for fi42 := 0; fi42 < 15; fi42++ {
                v43 := []int{}
                switch 112 {
                case -1:
                    for fi44 := 0; fi44 < 5; fi44++ {
                        switch (fi42 + -9223372036854775808) {
                        case 4294967295:
                            switch ("hello" + ("" + "a")) {
                            case "foo bar":
                                var v45 uint16 = 1
                                _ = v45
                            default:
                                if ("foo bar" >= "a") {
                                    v43 = append(v43, 874)
                                } else {
                                    v4 -= (((v40 - (((((v40 / (((3.14159 + (0.1 * (((v40 / (v40 - 0.1)) / v40) - v40))) * v40) * v4)) * ((v40 - (float64((0.1 - (float32(v4) - 0.0))) / -2.25)) + float64(1.5))) / (0.1 * ((float64(1e10) * float64(((1.5 - 3.14159) + 0.1))) + v4))) - v40) - v4)) / v40) - v4)
                                }
                            }
                            if false {
                                if true {
                                    fmt.Println(v4)
                                    var v46 bool = ((true || true) || ((-2.25 * 0.0) > 3.14159))
                                    _ = v46
                                }
                            }
                        case 766:
                            if (false || (!(((1.5 - 1e10) == 3.14159) && (!true)))) {
                                v43 = append(v43, (^(fi44)))
                                fmt.Println(v38, v41, fi44)
                            } else {
                                v4 -= 1.5
                                var v47 int64 = 9223372036854775807
                                v40 -= v40
                                _ = v47
                            }
                        case 127:
                            v48 := GenS4{F0: (^(v38)), F1: 0}
                            _ = v48
                        default:
                            _ = len(v3)
                        }
                        var v49 bool = true
                        _ = fi44
                        _ = v49
                    }
                case 1:
                    fmt.Println(v41, v41, fi42)
                }
                _ = fi42
                _ = v43
            }
            _ = v40
            _ = v41
        }
    } else {
        delete(v3, "\n\t")
        for fi50 := 0; fi50 < 8; fi50++ {
            v51 := GenS1{F0: (v5 - (v5 * v5))}
            v2 = v2
            v52 := GenS1{F0: v5}
            _ = fi50
            _ = v51
            _ = v52
        }
    }
    fmt.Println(v5, v38)
    fmt.Println(v1)
    v53 := GenS5{F0: 570, F1: 1, F2: 1}
    v54 := []int{}
    v1 -= 127
    v1 += -1
    switch v1 {
    case 32767:
        for fi55 := 0; fi55 < 3; fi55++ {
            var v56 uint16 = 2
            _ = fi55
            _ = v56
        }
    case -1:
        for fi57 := 0; fi57 < 14; fi57++ {
            v58 := GenS5{F0: 362, F1: 9223372036854775807, F2: 9223372036854775807}
            _ = fi57
            _ = v58
        }
    }
    var v59 int = v1
    switch (v1 >> ((255) & 63)) {
    case -321:
        var v60 uint16 = 255
        _ = v60
    case -2147483648:
        for fi61 := 0; fi61 < 7; fi61++ {
            fmt.Println(v59, fi61, v4)
            fmt.Println(v59, v5)
            _ = fi61
        }
    default:
        v59 = int(v38)
    }
    delete(v3, "hello")
    var v62 int32 = -683
    var v63 uint64 = (v38 + v38)
    var v64 uint64 = uint64(v2)
    for fi65 := 0; fi65 < 17; fi65++ {
        for fi66 := 0; fi66 < 11; fi66++ {
            v59 = fi65
            _ = len(v3)
            _ = fi66
        }
        for fi67 := 0; fi67 < 8; fi67++ {
            if (!((!(v64 == v64)) || true)) {
                if true {
                    var v68 int64 = -9223372036854775808
                    var v69 int16 = 469
                    if (!false) {
                        fmt.Println(fi65, v2, v5)
                    } else {
                        v64 = (v63 << ((1) & 63))
                        switch fi67 {
                        case 553:
                            var v70 uint64 = (v38 & v64)
                            var v71 int8 = genFn0(v63, 618, -1)
                            fmt.Println(v4, v1)
                            _ = v70
                            _ = v71
                        case 1:
                            v54 = append(v54, -1)
                            v68 = (v68 * (v68 - v68))
                        case 850:
                            v72 := GenS6{F0: 861, F1: 2147483647, F2: 42}
                            _ = v72
                        }
                        if (!true) {
                            if false {
                                if (-128 <= -17) {
                                    if false {
                                        v73 := GenS7{F0: (^(v2))}
                                        var v74 int32 = (^(v62))
                                        v59 += 9223372036854775807
                                        _ = v73
                                        _ = v74
                                    }
                                }
                                fmt.Println(v4)
                            } else {
                                v75 := GenS8{F0: "foo bar", F1: v2, F2: (true && false)}
                                v59 += (fi67 - (fi65 % (((v1 - (^(v59)))) | 1)))
                                var v76 string = ("a" + (((("a" + ("foo bar" + ((((("" + ("foo bar" + "hello")) + "a") + (((("hello" + ((("hello" + "") + "") + ((((("foo bar" + ("a" + ((("\n\t" + "hello") + ("\n\t" + ("foo bar" + ""))) + ("a" + "foo bar")))) + "") + ("foo bar" + "a")) + ((("foo bar" + "\n\t") + "hello") + "a")) + "foo bar"))) + "hello") + "a") + ("" + "\n\t"))) + ("\n\t" + "foo bar")) + ("\n\t" + "a")))) + "a") + ("a" + "foo bar")) + ("a" + "foo bar")))
                                _ = v75
                                _ = v76
                            }
                            var v77 uint64 = (v64 / ((v64) | 1))
                            v5 = v5
                            _ = v77
                        } else {
                            v38 -= 0
                            var v78 uint = 0
                            var v79 float32 = 0.1
                            _ = v78
                            _ = v79
                        }
                        v1 = fi65
                    }
                    _ = v68
                    _ = v69
                } else {
                    v62 += (v62 - v62)
                }
                v80 := GenS9{F0: 42, F1: v2}
                switch "a" {
                case "\n\t":
                    if (-9223372036854775808 != -9223372036854775808) {
                        fmt.Println(v64)
                    }
                case "a":
                    var v81 int32 = v62
                    _ = v81
                case "foo bar":
                    v1 = (fi67 | fi67)
                default:
                    var v82 uint8 = (v2 % ((42) | 1))
                    _ = v82
                }
                _ = v80
            }
            if false {
                switch v1 {
                case 42:
                    var v83 int32 = v62
                    _ = v83
                case -9223372036854775808:
                    _ = len(v54)
                case 769:
                    var v84 uint64 = (v38 * v38)
                    _ = v84
                }
                switch "foo bar" {
                case "foo bar":
                    var v85 int = 32767
                    _ = v85
                }
            } else {
                if true {
                    if (false || false) {
                        var v86 string = ("" + "\n\t")
                        if true {
                            v54 = append(v54, (v1 >> ((4294967295) & 63)))
                        } else {
                            fmt.Println(v86)
                        }
                        _ = v86
                    } else {
                        _ = len(v3)
                        v1 = fi65
                        v2 += 1
                        if len(v54) >= 2 {
                            v54 = v54[1:len(v54)]
                        }
                    }
                }
            }
            switch (fi67 & -2147483648) {
            case 42:
                var v87 float32 = 1e10
                _ = v87
            case 127:
                v1 = v59
            case 128:
                var v88 int16 = 1
                _ = v88
            }
            _ = fi67
        }
        _ = fi65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v54:", len(v54), v54)
    fmt.Println("v59:", v59)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v59
    _ = v62
    _ = v63
    _ = v64
}
