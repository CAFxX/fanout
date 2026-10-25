package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
type GenS4 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 uint; F1 bool; F2 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 string; F1 uint16; F2 int32 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 uint; F1 uint8 }
type GenS13 struct { F0 bool; F1 uint16; F2 int16 }
type GenS14 struct { F0 int16; F1 string; F2 bool }
func genFn0(p0 uint8, p1 uint16, p2 int) uint64 {
    v5 := GenS1{F0: 127, F1: 132}
    p0 += (p0 - 0)
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(fi6 >> ((1) & 63)), fi6, (-(p2)), 128}
        p2 -= fi6
        v8 := GenS1{F0: 83, F1: 18446744073709551615}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{p2}
    var v10 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v9
    return 0
}

func main() {
    v1 := []int{-9223372036854775808}
    v2 := make(map[string]bool)
    var v3 int = -2147483648
    var v4 float32 = -2.25
    var v11 int32 = 0
    switch (("a" + "") + "a") {
    case "foo bar":
        var v12 int8 = -21
        _ = v12
    case "hello":
        v13 := GenS1{F0: -128, F1: 18446744073709551615}
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 18; fi14++ {
            v3 = fi14
            v4 = 0.1
            var v15 int8 = 0
            _ = fi14
            _ = v15
        }
    }
    v2[((("hello" + "") + "") + ("" + "\n\t"))] = ((v3 > v3) && (9223372036854775807 >= 576))
    for fi16 := 0; fi16 < 16; fi16++ {
        fmt.Println(fi16)
        v4 += float32((float64((float32(((0.0 * 3.14159) + -2.25)) * float32((3.14159 + float64((v4 / v4)))))) - (1e10 / (-2.25 + (float64(v4) - -2.25)))))
        _ = fi16
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        v4 -= (((((v4 - (1.5 - float32(-2.25))) + (-2.25 + (0.0 * v4))) + (v4 * (((v4 + v4) + ((v4 - float32((3.14159 / (0.1 * 0.1)))) / 1.0)) + ((v4 + v4) / v4)))) + (0.0 * (((1.5 * v4) * (-2.25 * ((-2.25 * v4) / v4))) / v4))) / -2.25)
        v11 += 42
        _ = fi17
    }
    var v18 uint = 1
    v11 += 2147483647
    switch 32768 {
    case 4294967295:
        fmt.Println(v11)
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        if ((!(65535 != (v18 / ((2) | 1)))) && false) {
            if (127 <= 42) {
                var v20 int8 = 100
                v3 = (v3 - (fi19 / ((803) | 1)))
                v21 := GenS1{F0: v20, F1: v18}
                _ = v20
                _ = v21
            } else {
                var v22 uint8 = 85
                v23 := GenS2{F0: 9223372036854775807}
                _ = v22
                _ = v23
            }
            fmt.Println(v4, fi19)
        }
        switch (v3 * (v3 / (((fi19 >> ((0) & 63))) | 1))) {
        case 4294967295:
            for fi24 := 0; fi24 < 19; fi24++ {
                _ = len(v2)
                var v25 int16 = 931
                var v26 uint32 = 2
                _ = fi24
                _ = v25
                _ = v26
            }
        case 9223372036854775807:
            if false {
                if ("" > (("foo bar" + "hello") + ("hello" + "foo bar"))) {
                    v2["hello"] = false
                    var v27 uint = v18
                    switch "hello" {
                    case "a":
                        for fi28 := 0; fi28 < 20; fi28++ {
                            v29 := GenS3{F0: false, F1: 383, F2: "hello"}
                            var v30 int64 = 0
                            fmt.Println(v11, fi19, v11)
                            _ = fi28
                            _ = v29
                            _ = v30
                        }
                    default:
                        v18 += (v18 ^ 2)
                    }
                    _ = v27
                } else {
                    switch (v3 & -2147483648) {
                    case 32768:
                        if (true || true) {
                            delete(v2, (((("hello" + "hello") + "foo bar") + ("hello" + (("a" + (("" + "a") + "")) + ""))) + ("foo bar" + "\n\t")))
                            if false {
                                switch fi19 {
                                case 4294967295:
                                    v11 = (v11 | (v11 + v11))
                                case -922:
                                    var v31 bool = (-9223372036854775808 >= 0)
                                    _ = v31
                                case 127:
                                    if true {
                                        var v32 uint8 = 1
                                        if (42 > 496) {
                                            v18 = (v18 ^ v18)
                                            fmt.Println(v32, v11)
                                            v33 := GenS3{F0: true, F1: 60, F2: "hello"}
                                            _ = v33
                                        }
                                        delete(v2, ((("\n\t" + (("\n\t" + ("" + ("" + "a"))) + "a")) + (("" + "\n\t") + "foo bar")) + "a"))
                                        _ = v32
                                    } else {
                                        v3 += (-(v3))
                                        if false {
                                            var v34 uint = 4294967295
                                            fmt.Println(v18, fi19, fi19)
                                            _ = v34
                                        }
                                        var v35 int16 = 129
                                        _ = v35
                                    }
                                default:
                                    v36 := GenS4{F0: 1, F1: 127, F2: 42}
                                    _ = v36
                                }
                                v18 += 268
                                var v37 int32 = 42
                                _ = v37
                            } else {
                                v38 := GenS5{F0: v18}
                                _ = v38
                            }
                        } else {
                            var v39 uint64 = 4294967295
                            var v40 uint8 = uint8(v18)
                            v41 := v2["foo bar"]
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    default:
                        switch v3 {
                        case 127:
                            switch -486 {
                            case -416:
                                var v42 string = ("" + ("\n\t" + "foo bar"))
                                _ = v42
                            case 32767:
                                if true {
                                    fmt.Println(v4, v11)
                                    for fi43 := 0; fi43 < 5; fi43++ {
                                        if (2147483647 > 93) {
                                            v44 := v2["a"]
                                            _ = v44
                                        } else {
                                            v45 := GenS5{F0: 42}
                                            v46 := GenS4{F0: 1, F1: -128, F2: 736}
                                            _ = v45
                                            _ = v46
                                        }
                                        var v47 int32 = v11
                                        if len(v1) > 0 {
                                            v1[0] = fi43
                                        }
                                        delete(v2, "hello")
                                        if false {
                                            v48 := GenS1{F0: 42, F1: (v18 + (^(v18)))}
                                            var v49 int64 = int64(v47)
                                            _ = v48
                                            _ = v49
                                        }
                                        _ = fi43
                                        _ = v47
                                    }
                                }
                            case -2147483648:
                                delete(v2, "a")
                            default:
                                v4 = (v4 + v4)
                            }
                        case 549:
                            var v50 string = ("foo bar" + "")
                            _ = v50
                        case 4294967295:
                            var v51 int32 = -2147483648
                            _ = v51
                        }
                    }
                    for fi52 := 0; fi52 < 2; fi52++ {
                        if true {
                            var v53 uint32 = 2147483648
                            _ = len(v1)
                            _ = v53
                        }
                        v1 = append(v1, fi52)
                        _ = fi52
                    }
                }
                v1 = append(v1, (fi19 % ((fi19) | 1)))
            }
        case -579:
            var v54 uint16 = 215
            _ = v54
        default:
            if true {
                if true {
                    v55 := GenS6{F0: "foo bar"}
                    for fi56 := 0; fi56 < 19; fi56++ {
                        if (!true) {
                            switch "\n\t" {
                            case "a":
                                v57 := GenS7{F0: v11, F1: -1}
                                _ = v57
                            default:
                                fmt.Println(fi56, fi19)
                            }
                            var v58 int16 = 42
                            _ = v58
                        }
                        v3 = v3
                        var v59 int32 = int32(v18)
                        v60 := v2[(("hello" + "") + "foo bar")]
                        _ = fi56
                        _ = v59
                        _ = v60
                    }
                    _ = v55
                }
                v18 = 2
                _ = len(v1)
            } else {
                v4 += float32(0.1)
                v61 := GenS8{F0: 832, F1: false, F2: 42}
                fmt.Println(v11, fi19)
                _ = v61
            }
        }
        for fi62 := 0; fi62 < 1; fi62++ {
            var v63 float64 = 0.0
            var v64 int8 = 100
            if len(v1) > 0 {
                v1[0] = fi19
            }
            _ = fi62
            _ = v63
            _ = v64
        }
        _ = fi19
    }
    for fi65 := 0; fi65 < 19; fi65++ {
        if false {
            switch "foo bar" {
            case "hello":
                switch -9223372036854775808 {
                case -2147483648:
                    for fi66 := 0; fi66 < 14; fi66++ {
                        switch (fi65 >> (((v18 >> (((v18 / ((v18) | 1))) & 63))) & 63)) {
                        case 313:
                            v67 := GenS9{F0: 127}
                            v68 := GenS10{F0: "hello", F1: 65535, F2: v11}
                            var v69 uint16 = 32767
                            _ = v67
                            _ = v68
                            _ = v69
                        }
                        _ = fi66
                    }
                case 32767:
                    v70 := GenS11{F0: 42}
                    _ = v70
                case -1:
                    v71 := GenS12{F0: (^(v18)), F1: 91}
                    _ = v71
                }
            case "a":
                if false {
                    v3 = (fi65 << (((v18 & v18)) & 63))
                    for fi72 := 0; fi72 < 12; fi72++ {
                        v4 += v4
                        if false {
                            if true {
                                var v73 float32 = float32(-2.25)
                                v74 := GenS13{F0: true, F1: 255, F2: 0}
                                _ = v73
                                _ = v74
                            }
                            v4 = v4
                            if (!(-1 == 1)) {
                                if false {
                                    if (887 < 4294967295) {
                                        var v75 uint32 = 83
                                        _ = v75
                                    }
                                } else {
                                    var v76 uint32 = 2147483648
                                    if true {
                                        v77 := GenS7{F0: int32(v76), F1: 42}
                                        var v78 uint64 = genFn0(151, 65535, 40)
                                        if true {
                                            if true {
                                                v78 += (v78 >> (((v18 & 42)) & 63))
                                            }
                                            delete(v2, "")
                                        } else {
                                            var v79 int = 60
                                            fmt.Println(v4, v79, fi65)
                                            _ = v79
                                        }
                                        _ = v77
                                        _ = v78
                                    } else {
                                        _ = len(v2)
                                    }
                                    fmt.Println(v76)
                                    v80 := GenS14{F0: 525, F1: "hello", F2: true}
                                    if len(v1) > 0 {
                                        v1[0] = 32767
                                    }
                                    switch fi65 {
                                    case 32767:
                                        _ = len(v1)
                                    case 32768:
                                        fmt.Println(v11)
                                    case 128:
                                        if ((uint(v76) != 0) || (4294967295 < 0)) {
                                            v81 := GenS5{F0: (v18 / (((v18 % (((v18 + v18)) | 1))) | 1))}
                                            var v82 uint64 = uint64(v76)
                                            _ = v81
                                            _ = v82
                                        } else {
                                            var v83 int8 = -1
                                            v84 := GenS12{F0: v18, F1: 127}
                                            fmt.Println(v11, v18, v83)
                                            _ = v83
                                            _ = v84
                                        }
                                    }
                                    v4 = ((1e10 - v4) * float32(1.5))
                                    v85 := GenS5{F0: 2}
                                    _ = v76
                                    _ = v80
                                    _ = v85
                                }
                            }
                            if (true || true) {
                                v4 = 0.1
                            } else {
                                fmt.Println(fi65, v4)
                                if ((false && (((0.0 < 0.1) || true) && (v18 == v18))) || (((false && true) || (false && (uint32(v18) >= 2147483647))) && false)) {
                                    fmt.Println(v18, fi65, fi65)
                                    var v86 int8 = 109
                                    _ = v86
                                } else {
                                    _ = len(v2)
                                    var v87 uint = v18
                                    _ = v87
                                }
                                switch fi72 {
                                case -2147483648:
                                    v88 := GenS9{F0: -2147483648}
                                    _ = v88
                                }
                            }
                        } else {
                            v3 = 128
                        }
                        _ = fi72
                    }
                    v4 = v4
                }
            case "\n\t":
                v89 := GenS6{F0: "a"}
                _ = v89
            }
            fmt.Println(v11, v18)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v90 int8 = 1
        _ = fi65
        _ = v90
    }
    v91 := GenS11{F0: 2}
    v3 = 0
    switch 1 {
    case -940:
        v11 += (-(v11))
    case 846:
        v18 = 4294967295
    case 793:
        switch "hello" {
        case "hello":
            switch ("" + "hello") {
            case "\n\t":
                v92 := GenS5{F0: (v18 - v18)}
                _ = v92
            case "a":
                v18 = v18
            }
        default:
            fmt.Println(v11)
        }
    }
    if false {
        fmt.Println(v18, v3, v18)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v91
}
