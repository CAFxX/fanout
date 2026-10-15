package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int; F1 int8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint32; F1 uint8; F2 uint }
type GenS11 struct { F0 int; F1 int16; F2 int }
type GenS12 struct { F0 uint8 }
func genFn0(p0 int64, p1 int32) string {
    v7 := GenS1{F0: (0 <= 127)}
    v8 := GenS2{F0: 2, F1: 2, F2: 42}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return ((("" + "") + "") + (("\n\t" + ("a" + "\n\t")) + "a"))
}
func genFn1(p0 uint, p1 uint64, p2 uint8) bool {
    var v9 string = "foo bar"
    if true {
        for fi10 := 0; fi10 < 3; fi10++ {
            p2 = (p2 / ((p2) | 1))
            _ = fi10
        }
        v9 = v9
    } else {
        v11 := []int{32768, 60}
        fmt.Println(p0, v9, p2)
        var v12 uint8 = p2
        _ = v11
        _ = v12
    }
    p2 = (^(p2))
    for fi13 := 0; fi13 < 5; fi13++ {
        switch v9 {
        case "a":
            v14 := []int{fi13, 682, fi13}
            _ = v14
        default:
            v15 := GenS2{F0: p1, F1: p2, F2: 25}
            _ = v15
        }
        var v16 uint16 = 65535
        _ = fi13
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return true
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 3.14159
    var v3 float32 = (((0.0 + 1e10) / ((float32(1e10) + -2.25) - float32((v2 + v2)))) / -2.25)
    var v4 uint64 = 134
    var v5 bool = (!false)
    var v6 int = 2147483647
    v17 := []int{(v6 | (v6 ^ (^(v6)))), 917, v6}
    _ = len(v1)
    var v18 uint64 = (^(v4))
    switch "foo bar" {
    case "":
        var v19 uint16 = 2
        _ = v19
    case "a":
        _ = len(v1)
    case "foo bar":
        v17 = append(v17, 128)
    default:
        v18 += v4
    }
    var v20 uint8 = 127
    v2 = v2
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    fmt.Println(v20, v18, v18)
    _ = len(v1)
    for fi21 := 0; fi21 < 11; fi21++ {
        if false {
            for fi22 := 0; fi22 < 20; fi22++ {
                if (-1 > 2147483647) {
                    v23 := GenS3{F0: "a"}
                    switch (v6 * fi21) {
                    case -933:
                        switch "hello" {
                        case "foo bar":
                            v17 = append(v17, (fi22 + 2147483647))
                        case "\n\t":
                            v6 = v6
                        }
                    default:
                        var v24 int16 = 493
                        _ = v24
                    }
                    if (-102 <= 0) {
                        switch (^(v6)) {
                        case 32767:
                            var v25 string = ("hello" + "foo bar")
                            _ = v25
                        case 1:
                            v17 = append(v17, (-(fi21)))
                        default:
                            var v26 uint32 = 2147483647
                            _ = v26
                        }
                        var v27 bool = (!v5)
                        v3 -= (v3 * float32(v2))
                        _ = v27
                    }
                    _ = v23
                } else {
                    fmt.Println(v2)
                    var v28 uint16 = 309
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    _ = v28
                }
                if v5 {
                    var v29 uint32 = 2147483647
                    v20 = (v20 / ((127) | 1))
                    _ = v29
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        v5 = (-2147483648 == 42)
                    case "a":
                        v4 = (v4 % ((v18) | 1))
                    }
                }
                if (600 == 2) {
                    v18 = 4294967295
                    v30 := v1[9223372036854775807]
                    v5 = (float64(((v3 / -2.25) / -2.25)) > ((v2 * (-2.25 + float64(v3))) * v2))
                    _ = len(v1)
                    var v31 uint64 = 0
                    _ = v30
                    _ = v31
                } else {
                    if v5 {
                        v2 -= float64(3.14159)
                        var v32 uint32 = 0
                        _ = v32
                    } else {
                        if v5 {
                            if false {
                                fmt.Println(v4, v6)
                                v18 -= v18
                                v33 := GenS3{F0: "foo bar"}
                                _ = v33
                            }
                        } else {
                            if ("a" >= "") {
                                v3 = v3
                                fmt.Println(fi21, v3)
                            }
                            v34 := GenS4{F0: 0}
                            v35 := GenS5{F0: "foo bar", F1: v20, F2: 1}
                            fmt.Println(v6)
                            fmt.Println(v6, v3)
                            _ = v34
                            _ = v35
                        }
                        v36 := GenS4{F0: (v20 % ((v20) | 1))}
                        _ = v36
                    }
                }
                _ = fi22
            }
            var v37 uint16 = 2
            _ = v37
        } else {
            v38 := GenS4{F0: 0}
            _ = v38
        }
        _ = fi21
    }
    if v5 {
        v3 = 1.5
    } else {
        v39 := GenS6{F0: 42}
        if v5 {
            if true {
                delete(v1, (-(v6)))
                v40 := GenS4{F0: v20}
                var v41 int64 = 819
                _ = v40
                _ = v41
            } else {
                if len(v17) > 0 {
                    v17[0] = (v6 / (((-(v6))) | 1))
                }
            }
            switch "" {
            case "hello":
                v1[v6] = "\n\t"
            case "foo bar":
                if (-656 > -1) {
                    for fi42 := 0; fi42 < 12; fi42++ {
                        var v43 int8 = 100
                        if v5 {
                            switch ("" + ("hello" + "foo bar")) {
                            case "a":
                                if v5 {
                                    for fi44 := 0; fi44 < 3; fi44++ {
                                        switch 2147483647 {
                                        case 992:
                                            switch "a" {
                                            case "hello":
                                                var v45 uint8 = (v20 | 2)
                                                _ = v45
                                            case "\n\t":
                                                switch (v6 | 32767) {
                                                case -2147483648:
                                                    v2 += v2
                                                case -658:
                                                    v43 = v43
                                                case 32768:
                                                    if len(v17) > 0 {
                                                        v17[0] = (^(v6))
                                                    }
                                                default:
                                                    var v46 int16 = -32768
                                                    _ = v46
                                                }
                                            }
                                        case -2147483648:
                                            if false {
                                                if (2147483647 >= 801) {
                                                    switch (fi42 | -408) {
                                                    case 32767:
                                                        v47 := GenS5{F0: "\n\t", F1: (v20 & v20), F2: -456}
                                                        switch "a" {
                                                        case "":
                                                            var v48 int64 = 0
                                                            _ = v48
                                                        case "a":
                                                            if ((1e10 - v2) < (0.0 + float64(float32(0.0)))) {
                                                                var v49 int32 = -654
                                                                fmt.Println(v49, v18, v43)
                                                                _ = v49
                                                            } else {
                                                                if (!(!(32767 > 374))) {
                                                                    switch (fi42 * fi44) {
                                                                    case -636:
                                                                        v3 = -2.25
                                                                    case 2147483647:
                                                                        var v50 uint16 = 42
                                                                        _ = v50
                                                                    default:
                                                                        v3 = -2.25
                                                                    }
                                                                    v43 = (v43 + 100)
                                                                    v20 = 0
                                                                    var v51 uint32 = 675
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        if false {
                                                                            var v52 int = fi44
                                                                            _ = v52
                                                                        } else {
                                                                            var v53 uint16 = 332
                                                                            var v54 uint64 = (^(v4))
                                                                            _ = v53
                                                                            _ = v54
                                                                        }
                                                                    case "foo bar":
                                                                        var v55 uint64 = v18
                                                                        _ = v55
                                                                    case "hello":
                                                                        v56 := v1[(fi44 + 32767)]
                                                                        _ = v56
                                                                    }
                                                                    _ = v51
                                                                } else {
                                                                    v18 -= v18
                                                                    fmt.Println(fi44, v6, v43)
                                                                    _ = len(v17)
                                                                }
                                                            }
                                                            fmt.Println(v18, fi44)
                                                            switch (-(v6)) {
                                                            case -1:
                                                                delete(v1, 4294967295)
                                                            case -947:
                                                                v2 -= v2
                                                            case -97:
                                                                if false {
                                                                    v2 += ((v2 / float64(v3)) / v2)
                                                                    fmt.Println(v6, v6, v6)
                                                                    if ((v43 * v43) >= v43) {
                                                                        var v57 float64 = 3.14159
                                                                        v58 := GenS7{F0: 461}
                                                                        v17 = append(v17, int(v43))
                                                                        _ = v57
                                                                        _ = v58
                                                                    } else {
                                                                        _ = len(v1)
                                                                        var v59 int32 = 0
                                                                        fmt.Println(v59)
                                                                        _ = v59
                                                                    }
                                                                } else {
                                                                    fmt.Println(v3, v5, fi44)
                                                                }
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(v3, fi44)
                                                        default:
                                                            var v60 uint = 42
                                                            _ = v60
                                                        }
                                                        _ = v47
                                                    case 42:
                                                        if (1 <= -597) {
                                                            if v5 {
                                                                v43 = v43
                                                                v61 := GenS8{F0: fi42, F1: v43}
                                                                _ = v61
                                                            }
                                                            v62 := GenS7{F0: 65535}
                                                            v18 = (v18 / ((v18) | 1))
                                                            var v63 uint16 = 0
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                    case -436:
                                                        if false {
                                                            if true {
                                                                var v64 int16 = 32767
                                                                switch "hello" {
                                                                case "":
                                                                    if v5 {
                                                                        var v65 uint64 = (v18 / (((v18 >> ((42) & 63))) | 1))
                                                                        _ = v65
                                                                    } else {
                                                                        var v66 string = ("foo bar" + "foo bar")
                                                                        if len(v17) > 0 {
                                                                            v17[0] = fi42
                                                                        }
                                                                        v67 := GenS9{F0: v18}
                                                                        v43 = (v43 + 1)
                                                                        _ = v66
                                                                        _ = v67
                                                                    }
                                                                case "foo bar":
                                                                    if (42 >= 42) {
                                                                        v2 = 0.0
                                                                    }
                                                                case "a":
                                                                    _ = len(v1)
                                                                }
                                                                v68 := v1[(fi44 ^ (v6 >> ((uint(v18)) & 63)))]
                                                                _ = v64
                                                                _ = v68
                                                            }
                                                            if false {
                                                                v4 = v4
                                                                if (2 > 2147483648) {
                                                                    _ = len(v1)
                                                                    var v69 uint8 = 128
                                                                    _ = v69
                                                                } else {
                                                                    var v70 uint = uint(v20)
                                                                    v43 = (v43 - v43)
                                                                    _ = v70
                                                                }
                                                                if true {
                                                                    if len(v17) > 0 {
                                                                        v17[0] = fi42
                                                                    }
                                                                    v71 := GenS4{F0: 255}
                                                                    _ = v71
                                                                }
                                                                fmt.Println(v18)
                                                                if (!false) {
                                                                    var v72 int32 = 42
                                                                    v3 = 0.1
                                                                    v72 += (v72 << ((0) & 63))
                                                                    _ = v72
                                                                } else {
                                                                    switch "hello" {
                                                                    case "a":
                                                                        v18 = v18
                                                                    default:
                                                                        switch "a" {
                                                                        case "":
                                                                            fmt.Println(v6)
                                                                        }
                                                                    }
                                                                    v73 := GenS10{F0: 2, F1: v20, F2: 4294967295}
                                                                    _ = v73
                                                                }
                                                            } else {
                                                                var v74 uint8 = (v20 << ((16) & 63))
                                                                var v75 int8 = v43
                                                                _ = v74
                                                                _ = v75
                                                            }
                                                            if len(v17) > 0 {
                                                                v17[0] = (fi44 / ((160) | 1))
                                                            }
                                                        }
                                                        if ((v2 + v2) != ((-2.25 + v2) / float64((v3 / float32(0.1))))) {
                                                            switch ("" + "") {
                                                            case "foo bar":
                                                                v76 := GenS8{F0: 825, F1: v43}
                                                                _ = v76
                                                            case "":
                                                                v1[(fi42 % ((fi42) | 1))] = (("hello" + ("" + "foo bar")) + "hello")
                                                            }
                                                        }
                                                    }
                                                    v77 := GenS11{F0: fi42, F1: -708, F2: int(v20)}
                                                    _ = v77
                                                }
                                                v78 := GenS12{F0: (v20 - v20)}
                                                _ = v78
                                            }
                                        case 42:
                                            v3 = (v3 + v3)
                                        }
                                        var v79 uint32 = 979
                                        v80 := GenS10{F0: (v79 + v79), F1: (v20 * uint8(v18)), F2: uint(v79)}
                                        fmt.Println(v2, v2)
                                        _ = len(v1)
                                        _ = fi44
                                        _ = v79
                                        _ = v80
                                    }
                                } else {
                                    v3 -= v3
                                    var v81 int32 = -1
                                    if (v4 > (^(v18))) {
                                        if (!v5) {
                                            var v82 bool = v5
                                            _ = v82
                                        } else {
                                            delete(v1, 4294967295)
                                            var v83 float32 = 1e10
                                            _ = v83
                                        }
                                        v20 = (v20 << ((65535) & 63))
                                    }
                                    _ = v81
                                }
                            default:
                                v84 := GenS5{F0: "hello", F1: (^(v20)), F2: 1000}
                                _ = v84
                            }
                            if len(v17) > 0 {
                                v17[0] = (fi42 - (-(v6)))
                            }
                        }
                        _ = fi42
                        _ = v43
                    }
                    fmt.Println(v18, v2, v20)
                }
            }
        }
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
