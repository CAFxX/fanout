package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int16; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS6 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 uint8 }
type GenS11 struct { F0 bool; F1 uint; F2 string }
type GenS12 struct { F0 bool; F1 int32; F2 int64 }
func genFn0(p0 bool, p1 int16, p2 uint64) bool {
    if (0 != 65535) {
        p2 = p2
    } else {
        fmt.Println(p1)
        if p0 {
            v4 := []int{4294967295, -9223372036854775808}
            var v5 uint64 = p2
            _ = v4
            _ = v5
        } else {
            for fi6 := 0; fi6 < 8; fi6++ {
                v7 := GenS1{F0: 307}
                _ = fi6
                _ = v7
            }
            if true {
                for fi8 := 0; fi8 < 19; fi8++ {
                    switch "" {
                    case "":
                        var v9 int64 = 9223372036854775807
                        _ = v9
                    default:
                        v10 := GenS1{F0: 0}
                        _ = v10
                    }
                    _ = fi8
                }
                v11 := GenS2{F0: 1, F1: (p1 * (p1 << ((449) & 63))), F2: p1}
                _ = v11
            } else {
                switch "a" {
                case "":
                    v12 := []int{-168, 32767, 9223372036854775807}
                    _ = v12
                case "a":
                    var v13 int64 = 42
                    _ = v13
                }
                p2 = p2
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := GenS3{F0: (p2 ^ (^(p2))), F1: 65535}
                    var v16 uint16 = 679
                    var v17 int32 = 2147483647
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
            }
        }
    }
    var v18 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    return true
}
func genFn1(p0 uint16) uint16 {
    var v19 int = -1
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 int32 = -2147483648
        fmt.Println(p0, fi20)
        _ = fi20
        _ = v21
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v19
    return (p0 % (((p0 | p0)) | 1))
}

func main() {
    v1 := make(map[string]bool)
    var v2 float64 = (-2.25 - -2.25)
    var v3 uint = 18446744073709551615
    v22 := GenS4{F0: 4294967295}
    v23 := GenS4{F0: 127}
    switch 4294967295 {
    case 32767:
        if (0 != 42) {
            v2 = v2
            v24 := GenS2{F0: -775, F1: -32768, F2: 828}
            _ = v24
        } else {
            if true {
                if false {
                    var v25 string = ("a" + "hello")
                    for fi26 := 0; fi26 < 8; fi26++ {
                        for fi27 := 0; fi27 < 8; fi27++ {
                            v28 := GenS5{F0: 65535, F1: -1, F2: 1}
                            if false {
                                v25 = ((("hello" + ("hello" + ((("a" + v25) + v25) + v25))) + v25) + "a")
                                switch fi27 {
                                case -277:
                                    var v29 float32 = float32(((v2 / 3.14159) + 3.14159))
                                    var v30 uint = 860
                                    _ = v29
                                    _ = v30
                                case 9223372036854775807:
                                    v2 = (v2 - ((1.5 * float64(0.1)) * 0.1))
                                }
                            }
                            if (false || (!(-1 > 1))) {
                                var v31 uint8 = 72
                                _ = v31
                            }
                            _ = fi27
                            _ = v28
                        }
                        if ((int(v3) > (-(fi26))) || true) {
                            v32 := GenS6{F0: 42, F1: 1, F2: 309}
                            fmt.Println(v25)
                            var v33 int16 = 1000
                            _ = v32
                            _ = v33
                        } else {
                            var v34 float64 = float64(1.5)
                            v35 := GenS7{F0: 0}
                            for fi36 := 0; fi36 < 7; fi36++ {
                                v25 = v25
                                _ = fi36
                            }
                            _ = v34
                            _ = v35
                        }
                        v37 := []int{(fi26 & (fi26 | (fi26 - fi26)))}
                        _ = fi26
                        _ = v37
                    }
                    v25 = v25
                    _ = v25
                } else {
                    fmt.Println(v3, v3)
                    if true {
                        v2 += 1e10
                    } else {
                        for fi38 := 0; fi38 < 19; fi38++ {
                            switch fi38 {
                            case 42:
                                for fi39 := 0; fi39 < 15; fi39++ {
                                    switch fi38 {
                                    case -1:
                                        v3 -= v3
                                    case -36:
                                        v40 := GenS3{F0: 18446744073709551615, F1: genFn1(genFn1(2))}
                                        _ = v40
                                    }
                                    v41 := GenS7{F0: 127}
                                    _ = fi39
                                    _ = v41
                                }
                            default:
                                fmt.Println(fi38, fi38)
                            }
                            for fi42 := 0; fi42 < 15; fi42++ {
                                v2 = v2
                                if (false && true) {
                                    v43 := []int{}
                                    fmt.Println(fi42, v3)
                                    _ = v43
                                } else {
                                    v2 -= v2
                                }
                                var v44 float32 = -2.25
                                _ = fi42
                                _ = v44
                            }
                            v2 -= v2
                            _ = fi38
                        }
                    }
                    v45 := GenS5{F0: 42, F1: 927, F2: 55}
                    _ = v45
                }
            } else {
                if (-9 == 100) {
                    if false {
                        delete(v1, "a")
                        v46 := []int{42, 135, 457}
                        _ = v46
                    } else {
                        if false {
                            var v47 float64 = 0.1
                            v48 := GenS1{F0: (^(v3))}
                            if true {
                                v49 := GenS8{F0: 2147483647}
                                _ = len(v1)
                                v50 := []int{9223372036854775807, 127}
                                _ = v49
                                _ = v50
                            }
                            _ = v47
                            _ = v48
                        }
                        v51 := GenS8{F0: 311}
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 15; fi52++ {
                        v53 := []int{(fi52 * -732), fi52, -2147483648, -9223372036854775808}
                        var v54 uint32 = 1
                        v3 += 65535
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                    switch ("" + "") {
                    case "foo bar":
                        var v55 int8 = -1
                        _ = v55
                    }
                }
                for fi56 := 0; fi56 < 1; fi56++ {
                    switch "foo bar" {
                    case "foo bar":
                        v1[("" + "foo bar")] = ((fi56 >= fi56) && true)
                    case "hello":
                        var v57 int = (fi56 + fi56)
                        _ = v57
                    case "a":
                        fmt.Println(v3, v3)
                    default:
                        if (42 != 719) {
                            var v58 float32 = 0.1
                            var v59 uint64 = 125
                            v60 := GenS9{F0: (^(v59))}
                            _ = v58
                            _ = v59
                            _ = v60
                        }
                    }
                    _ = fi56
                }
            }
            v3 = 65535
            v3 += v3
        }
    case 42:
        var v61 uint64 = 4294967295
        _ = v61
    case 2147483647:
        switch "hello" {
        case "":
            switch "foo bar" {
            case "foo bar":
                v3 = 0
            }
        }
    }
    var v62 uint8 = 2
    v1[""] = (14 <= 526)
    v63 := []int{4294967295}
    if true {
        var v64 uint = v3
        if len(v63) > 0 {
            v63[0] = -901
        }
        _ = v64
    }
    if (810 == 2147483648) {
        if true {
            v3 += 701
            switch -1 {
            case -9223372036854775808:
                v65 := GenS10{F0: 551, F1: uint8(v3)}
                _ = v65
            case 127:
                v3 -= v3
            case 341:
                if ((100 == 0) || false) {
                    switch 4294967295 {
                    case 128:
                        fmt.Println(v2, v3, v62)
                    case -384:
                        for fi66 := 0; fi66 < 15; fi66++ {
                            var v67 string = (("\n\t" + "\n\t") + "hello")
                            _ = fi66
                            _ = v67
                        }
                    case -2147483648:
                        var v68 uint = v3
                        _ = v68
                    }
                    fmt.Println(v2, v62)
                }
            default:
                var v69 int8 = 28
                _ = v69
            }
        } else {
            if false {
                var v70 uint16 = 526
                switch 0 {
                case 128:
                    if len(v63) >= 2 {
                        v63 = v63[1:len(v63)]
                    }
                case 63:
                    fmt.Println(v70)
                case -1:
                    fmt.Println(v62, v70, v2)
                }
                _ = v70
            } else {
                if false {
                    if ((1.5 != float32(float64(0.0))) || (-2147483648 > 42)) {
                        v3 = 607
                    }
                    v63 = append(v63, 32768)
                }
                v71 := GenS5{F0: 2147483647, F1: 140, F2: 18446744073709551615}
                _ = v71
            }
        }
        if len(v63) >= 2 {
            v63 = v63[1:len(v63)]
        }
        v2 += (0.0 * 0.0)
    } else {
        var v72 uint16 = 42
        if len(v63) > 0 {
            v63[0] = 4294967295
        }
        v62 = (v62 << (((v3 * v3)) & 63))
        _ = v72
    }
    switch 1 {
    case 9223372036854775807:
        for fi73 := 0; fi73 < 16; fi73++ {
            if (true && (!false)) {
                if (111 == 255) {
                    switch fi73 {
                    case -827:
                        v3 += v3
                    default:
                        for fi74 := 0; fi74 < 16; fi74++ {
                            v62 -= v62
                            if (false && (42 <= -52)) {
                                var v75 int16 = 0
                                var v76 int32 = 128
                                _ = v75
                                _ = v76
                            }
                            v77 := GenS11{F0: true, F1: (^(v3)), F2: "hello"}
                            _ = fi74
                            _ = v77
                        }
                    }
                } else {
                    v3 = v3
                    if false {
                        v3 = (v3 >> (((v3 >> (((v3 & v3)) & 63))) & 63))
                    } else {
                        v78 := GenS12{F0: (((-500 < 1000) && false) || true), F1: 364, F2: 879}
                        _ = v78
                    }
                }
                var v79 uint16 = 32767
                _ = v79
            }
            var v80 uint16 = 926
            if len(v63) >= 2 {
                v63 = v63[1:len(v63)]
            }
            _ = fi73
            _ = v80
        }
    case 128:
        v81 := v1["\n\t"]
        _ = v81
    case 0:
        var v82 float64 = (v2 * ((float64((((-2.25 / (0.1 * 0.1)) * -2.25) / 1.0)) + ((v2 + (v2 / 0.1)) * 1.5)) - 1.5))
        _ = v82
    default:
        if len(v63) >= 2 {
            v63 = v63[1:len(v63)]
        }
    }
    for fi83 := 0; fi83 < 7; fi83++ {
        if len(v63) > 0 {
            v63[0] = (fi83 + fi83)
        }
        var v84 uint32 = 926
        _ = len(v63)
        _ = fi83
        _ = v84
    }
    for fi85 := 0; fi85 < 8; fi85++ {
        if len(v63) > 0 {
            v63[0] = -519
        }
        _ = fi85
    }
    v62 = (v62 & (v62 + (v62 / (((v62 / ((211) | 1))) | 1))))
    var v86 uint64 = 749
    v63 = append(v63, 1)
    var v87 int = -9223372036854775808
    if true {
        _ = len(v63)
        switch "hello" {
        case "a":
            v3 = 18446744073709551615
        case "\n\t":
            v2 += ((1e10 - -2.25) + v2)
        case "hello":
            delete(v1, "\n\t")
        }
    } else {
        for fi88 := 0; fi88 < 17; fi88++ {
            v89 := GenS6{F0: (v86 >> (((v3 | v3)) & 63)), F1: 2, F2: 2}
            v86 = v86
            _ = fi88
            _ = v89
        }
        if len(v63) >= 2 {
            v63 = v63[1:len(v63)]
        }
        v87 = -2147483648
    }
    v90 := GenS11{F0: false, F1: v3, F2: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v63:", len(v63), v63)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v62
    _ = v63
    _ = v86
    _ = v87
    _ = v90
}
