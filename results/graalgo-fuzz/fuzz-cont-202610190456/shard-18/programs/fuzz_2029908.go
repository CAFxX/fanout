package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 string; F1 uint8; F2 int64 }
type GenS4 struct { F0 bool; F1 int8; F2 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int16 }
type GenS8 struct { F0 uint32; F1 uint8 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int16; F1 uint16 }
type GenS12 struct { F0 string; F1 uint8; F2 uint16 }
func genFn0(p0 uint8, p1 uint32) int32 {
    var v4 string = ("hello" + (("foo bar" + "") + "a"))
    fmt.Println(p1, p0)
    v5 := []int{127, -978, -645}
    var v6 string = v4
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1() bool {
    v7 := []int{9223372036854775807, 9223372036854775807}
    var v8 float64 = 3.14159
    if (!(!false)) {
        for fi9 := 0; fi9 < 7; fi9++ {
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 uint8 = 161
                if true {
                    v12 := GenS1{F0: 42}
                    if (!(!(fi9 >= (^(fi9))))) {
                        var v13 int64 = -248
                        v14 := GenS1{F0: -1}
                        v8 = v8
                        if ("\n\t" >= "\n\t") {
                            fmt.Println(v11, v11)
                        }
                        var v15 int32 = 42
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        if false {
                            var v16 int64 = -1
                            var v17 uint64 = 550
                            v16 = v16
                            fmt.Println(v8, v8)
                            _ = v16
                            _ = v17
                        } else {
                            var v18 float64 = v8
                            var v19 float64 = ((-2.25 / 0.1) + (((float64(0.1) + v8) + 0.0) / ((v18 * 3.14159) - float64(1e10))))
                            _ = v18
                            _ = v19
                        }
                        v7 = append(v7, 2147483647)
                        var v20 int = (fi9 * 594)
                        _ = v20
                    }
                    _ = len(v7)
                    var v21 uint16 = 200
                    _ = v12
                    _ = v21
                }
                _ = fi10
                _ = v11
            }
            var v22 uint32 = 517
            v23 := GenS1{F0: 1000}
            _ = fi9
            _ = v22
            _ = v23
        }
    }
    var v24 float64 = v8
    _ = v24
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := []int{-1, -1, -357, 1}
    var v2 uint32 = 2
    v3 := make(map[int]int)
    if (255 >= 42) {
        fmt.Println(v2)
        v2 = (v2 << ((4294967295) & 63))
    }
    var v25 string = "\n\t"
    var v26 uint = 255
    v25 = v25
    var v27 uint32 = v2
    if (true || (!(1.5 != 0.0))) {
        if (42 < 0) {
            if (0.1 > 0.0) {
                switch "\n\t" {
                case "hello":
                    var v28 uint = 405
                    _ = v28
                default:
                    v26 = v26
                }
                v29 := GenS1{F0: int16(v26)}
                _ = v29
            } else {
                v27 = v27
            }
            delete(v3, -2147483648)
        }
        v30 := GenS2{F0: 178, F1: v2}
        if (v26 != (v26 & (v26 * v26))) {
            if false {
                var v31 uint64 = 735
                var v32 int32 = 42
                fmt.Println(v31)
                _ = v31
                _ = v32
            } else {
                for fi33 := 0; fi33 < 16; fi33++ {
                    switch v25 {
                    case "foo bar":
                        var v34 int64 = 877
                        _ = v34
                    case "":
                        var v35 string = (((((v25 + "hello") + v25) + v25) + "a") + "a")
                        _ = v35
                    default:
                        v36 := GenS3{F0: v25, F1: 2, F2: 42}
                        _ = v36
                    }
                    var v37 float32 = 1e10
                    _ = fi33
                    _ = v37
                }
            }
            _ = len(v1)
        } else {
            v38 := GenS4{F0: (v25 >= "\n\t"), F1: 42, F2: 187}
            v39 := GenS4{F0: true, F1: 0, F2: 1}
            fmt.Println(v25)
            _ = v38
            _ = v39
        }
        _ = v30
    }
    v27 = (v27 << ((v26) & 63))
    fmt.Println(v27, v26, v25)
    if (true || true) {
        v27 -= 2147483648
        v40 := GenS5{F0: false}
        _ = v40
    }
    v27 = (v2 % (((^(v2))) | 1))
    for fi41 := 0; fi41 < 17; fi41++ {
        for fi42 := 0; fi42 < 14; fi42++ {
            if (!(-73 <= -1)) {
                v26 -= (v26 / ((v26) | 1))
                if false {
                    var v43 float64 = float64(0.1)
                    _ = v43
                } else {
                    if (42 >= 182) {
                        v44 := v3[fi42]
                        _ = v44
                    } else {
                        delete(v3, (fi42 | fi41))
                        switch (fi42 & (fi41 % ((fi41) | 1))) {
                        case 9223372036854775807:
                            var v45 uint64 = 403
                            _ = v45
                        case 4294967295:
                            v3[fi42] = fi41
                        }
                    }
                    if false {
                        if true {
                            v46 := GenS6{F0: (v2 >> (((v26 >> ((v26) & 63))) & 63))}
                            if false {
                                var v47 string = (("" + v25) + v25)
                                _ = v47
                            } else {
                                v25 = (v25 + "")
                                if false {
                                    if true {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v48 := GenS7{F0: -125, F1: 1}
                                        switch (fi42 % ((197) | 1)) {
                                        case -417:
                                            var v49 float64 = (3.14159 * 3.14159)
                                            _ = v49
                                        case 32768:
                                            v3[(fi42 | (fi42 / (((fi42 | (^(fi42)))) | 1)))] = (fi42 & fi42)
                                        default:
                                            v50 := GenS8{F0: v27, F1: 13}
                                            _ = v50
                                        }
                                        _ = v48
                                    }
                                }
                                var v51 uint16 = 42
                                _ = v51
                            }
                            var v52 float64 = (-2.25 + ((float64(0.0) / -2.25) - (((-2.25 - (0.0 + 1.5)) + (3.14159 / (-2.25 * 3.14159))) * 0.0)))
                            v25 = v25
                            _ = v46
                            _ = v52
                        } else {
                            var v53 string = v25
                            _ = v53
                        }
                    }
                    var v54 uint8 = 42
                    _ = v54
                }
                var v55 float64 = ((-2.25 / 1.0) / 3.14159)
                _ = v55
            } else {
                var v56 uint64 = 4294967295
                switch 4294967295 {
                case 851:
                    if true {
                        _ = len(v3)
                        var v57 int16 = -1
                        _ = v57
                    } else {
                        if (!true) {
                            switch v25 {
                            case "\n\t":
                                var v58 bool = (127 >= 152)
                                _ = v58
                            case "hello":
                                fmt.Println(v25, v56, v27)
                            case "":
                                var v59 int8 = -88
                                _ = v59
                            default:
                                var v60 float32 = float32((0.0 * 1.5))
                                _ = v60
                            }
                            if (363 >= v26) {
                                v61 := GenS9{F0: (v25 + v25), F1: -534}
                                if (42 >= 0) {
                                    var v62 float32 = (0.1 / (1.5 / 1.0))
                                    switch ("" + ("foo bar" + v25)) {
                                    case "\n\t":
                                        if false {
                                            switch v25 {
                                            case "foo bar":
                                                if (1.5 != float64(v62)) {
                                                    _ = len(v3)
                                                }
                                            case "a":
                                                _ = len(v3)
                                            }
                                            var v63 float32 = (3.14159 / (v62 / v62))
                                            _ = v63
                                        }
                                    default:
                                        fmt.Println(v26, v62, v56)
                                    }
                                    v2 = v2
                                    _ = v62
                                } else {
                                    switch v25 {
                                    case "foo bar":
                                        var v64 int64 = 947
                                        _ = v64
                                    case "hello":
                                        _ = len(v1)
                                    default:
                                        if ((v56 + v56) != v56) {
                                            delete(v3, (fi42 / ((fi42) | 1)))
                                        }
                                    }
                                    v26 += (v26 / ((v26) | 1))
                                }
                                _ = v61
                            } else {
                                v1 = append(v1, (-(fi41)))
                            }
                            if (101 < -83) {
                                if true {
                                    if (int16(v2) != -277) {
                                        var v65 int8 = 100
                                        if true {
                                            v2 += 359
                                            _ = len(v1)
                                        } else {
                                            var v66 bool = (-232 < 630)
                                            v26 = v26
                                            v67 := GenS10{F0: (^(v56))}
                                            _ = v66
                                            _ = v67
                                        }
                                        var v68 uint16 = 255
                                        _ = v65
                                        _ = v68
                                    }
                                    v69 := GenS9{F0: "hello", F1: 93}
                                    var v70 uint32 = (v27 >> ((v26) & 63))
                                    _ = v69
                                    _ = v70
                                } else {
                                    v27 = (v2 - (v27 >> ((1) & 63)))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = len(v1)
                                }
                                v56 = v56
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                        } else {
                            v71 := GenS7{F0: -376, F1: 982}
                            if true {
                                v2 = 809
                            }
                            _ = v71
                        }
                        delete(v3, fi42)
                    }
                case -719:
                    v3[0] = 4294967295
                case -9223372036854775808:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    switch 0 {
                    case 32767:
                        v72 := GenS10{F0: (v56 % (((^(v56))) | 1))}
                        _ = v72
                    default:
                        var v73 uint32 = (v2 ^ (v27 - 638))
                        _ = v73
                    }
                }
                _ = v56
            }
            switch v25 {
            case "a":
                switch fi42 {
                case 128:
                    if (42 <= -1) {
                        v74 := GenS7{F0: 2147483647, F1: -1}
                        var v75 string = ""
                        _ = v74
                        _ = v75
                    } else {
                        v3[fi41] = fi42
                        var v76 uint = (v26 + (v26 | (v26 & 42)))
                        _ = v76
                    }
                    v1 = append(v1, fi41)
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            case "hello":
                v77 := GenS5{F0: true}
                _ = v77
            case "foo bar":
                v78 := GenS4{F0: false, F1: 56, F2: 1}
                _ = v78
            }
            _ = fi42
        }
        v79 := GenS3{F0: "hello", F1: uint8(v2), F2: 0}
        v80 := GenS11{F0: 82, F1: 255}
        _ = fi41
        _ = v79
        _ = v80
    }
    var v81 int = -939
    v82 := GenS12{F0: "\n\t", F1: 255, F2: 2}
    var v83 float32 = float32((((1e10 - (((0.1 * ((1e10 + 0.1) - (3.14159 + (((((0.1 * 0.0) / 3.14159) - 1.5) - 1.5) - -2.25)))) - float64(((((1.5 - ((((3.14159 / (0.1 + 1e10)) - 0.0) / 1.0) - float32(float64(((0.1 - -2.25) + float32(1e10)))))) / (((-2.25 + (((-2.25 / 1.0) * 3.14159) - 0.0)) * 0.1) / 1e10)) - (((0.1 - -2.25) - float32(-2.25)) / 1.0)) - (3.14159 * 0.1)))) * (1e10 / 1.0))) / (3.14159 + -2.25)) / 1.0))
    fmt.Println(v81)
    v84 := GenS9{F0: v25, F1: -549}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v81:", v81)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v81
    _ = v82
    _ = v83
    _ = v84
}
