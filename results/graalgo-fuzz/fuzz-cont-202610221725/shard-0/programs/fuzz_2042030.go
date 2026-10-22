package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int64; F1 int32 }
type GenS8 struct { F0 uint; F1 int16; F2 int16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int32; F1 int32 }
func genFn0(p0 int64, p1 int, p2 int32) int64 {
    fmt.Println(p2)
    for fi6 := 0; fi6 < 11; fi6++ {
        p2 += p2
        _ = fi6
    }
    switch p1 {
    case -87:
        p2 += 1
    case -1:
        for fi7 := 0; fi7 < 20; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 int32 = p2
                var v10 bool = false
                v11 := []int{(fi8 ^ fi7), p1, -1}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi7
        }
    case -139:
        v12 := GenS1{F0: 42, F1: -73, F2: 1}
        _ = v12
    default:
        fmt.Println(p0, p0, p0)
    }
    switch "foo bar" {
    case "a":
        var v13 int16 = -624
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 * p0)
}
func genFn1(p0 int16) uint16 {
    v14 := []int{}
    if (1e10 <= 0.0) {
        v15 := GenS2{F0: -1}
        v16 := GenS3{F0: ((!false) && ((1.5 / (-2.25 / -2.25)) > (1e10 * 1e10))), F1: genFn0(-1, 1, -1), F2: 42}
        if (!((!true) && (0 > 1))) {
            fmt.Println(p0)
            p0 = (p0 & p0)
            var v17 bool = (-1 <= 187)
            _ = v17
        } else {
            v14 = append(v14, 32767)
            v18 := GenS4{F0: 86}
            _ = v18
        }
        _ = v15
        _ = v16
    } else {
        for fi19 := 0; fi19 < 5; fi19++ {
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = fi19
        }
        if (800 != -239) {
            p0 += (p0 & (-(p0)))
        }
    }
    var v20 uint32 = 42
    v21 := GenS1{F0: 128, F1: 0, F2: (p0 - 525)}
    v14 = append(v14, 42)
    _ = p0
    _ = v14
    _ = v20
    _ = v21
    return 32767
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]bool)
    var v3 uint8 = 128
    var v4 int64 = 0
    var v5 uint32 = 42
    v22 := GenS2{F0: -322}
    if (288 < 42) {
        var v23 uint = 18446744073709551615
        for fi24 := 0; fi24 < 8; fi24++ {
            v4 += (-(v4))
            var v25 int8 = 127
            _ = fi24
            _ = v25
        }
        v26 := GenS5{F0: "\n\t"}
        _ = v23
        _ = v26
    }
    var v27 float64 = 1e10
    fmt.Println(v5, v3, v27)
    v1 = v1
    v5 += 701
    for fi28 := 0; fi28 < 12; fi28++ {
        fmt.Println(v27, v3, v4)
        _ = fi28
    }
    v29 := GenS6{F0: 934}
    if true {
        delete(v2, -933)
        v2[-2147483648] = ((v3 + uint8(v5)) < (^(v3)))
    }
    v27 = -2.25
    for fi30 := 0; fi30 < 16; fi30++ {
        if (genFn1(42) <= 255) {
            v5 = v5
        }
        if (false && false) {
            for fi31 := 0; fi31 < 20; fi31++ {
                delete(v2, 42)
                v2[(fi30 / (((fi31 ^ (fi31 / ((fi30) | 1)))) | 1))] = ((-(fi30)) > (fi31 + 2147483647))
                if (genFn1(-95) >= genFn1(1000)) {
                    var v32 float32 = v1
                    fmt.Println(v4, v5)
                    if true {
                        var v33 uint16 = genFn1(32767)
                        fmt.Println(v32)
                        _ = v33
                    } else {
                        var v34 uint8 = 128
                        switch "\n\t" {
                        case "a":
                            var v35 bool = false
                            _ = v35
                        case "hello":
                            var v36 int = fi30
                            _ = v36
                        default:
                            switch "foo bar" {
                            case "hello":
                                v4 += (v4 + -163)
                            case "":
                                var v37 int32 = int32(v4)
                                switch ("a" + "\n\t") {
                                case "":
                                    v5 = v5
                                case "foo bar":
                                    if false {
                                        v38 := []int{43, (fi31 - (fi31 + fi30)), (fi31 | 32767), (fi30 << ((42) & 63))}
                                        var v39 int = -424
                                        v27 = v27
                                        _ = v38
                                        _ = v39
                                    }
                                    v40 := GenS4{F0: v3}
                                    var v41 uint32 = (v5 - (v5 + 427))
                                    delete(v2, (fi31 | (^(fi31))))
                                    _ = v40
                                    _ = v41
                                case "a":
                                    if true {
                                        var v42 float64 = (v27 + ((((v27 - -2.25) / v27) - 0.1) + v27))
                                        var v43 int8 = 100
                                        _ = v42
                                        _ = v43
                                    }
                                default:
                                    v44 := []int{fi30, fi31}
                                    _ = v44
                                }
                                v5 -= 65535
                                _ = v37
                            case "\n\t":
                                var v45 uint64 = 1
                                _ = v45
                            default:
                                if (float32(3.14159) < (1e10 / (float32(-2.25) + v1))) {
                                    if (false && (fi31 > 32768)) {
                                        var v46 bool = false
                                        _ = v46
                                    }
                                }
                            }
                        }
                        _ = v34
                    }
                    _ = v32
                } else {
                    v47 := GenS2{F0: 1000}
                    _ = v47
                }
                _ = fi31
            }
            for fi48 := 0; fi48 < 14; fi48++ {
                switch "a" {
                case "hello":
                    fmt.Println(v1, v4)
                case "a":
                    v2[(-(fi48))] = (false && true)
                case "":
                    switch "" {
                    case "hello":
                        var v49 int8 = -128
                        var v50 uint32 = v5
                        var v51 bool = (-711 > 544)
                        fmt.Println(v4, v4, v49)
                        var v52 uint32 = (v50 & v50)
                        switch ((("\n\t" + ("" + "\n\t")) + (("a" + (("foo bar" + "") + (("foo bar" + "\n\t") + "hello"))) + ("a" + "\n\t"))) + "") {
                        case "":
                            var v53 float64 = (v27 / float64(float32((v27 / float64(((-2.25 - v1) * 1e10))))))
                            v52 = (v50 / ((2147483647) | 1))
                            v3 += (^(v3))
                            _ = v53
                        case "\n\t":
                            v54 := GenS5{F0: "foo bar"}
                            _ = v54
                        case "a":
                            if false {
                                v55 := GenS3{F0: (!false), F1: genFn0((v4 + genFn0((v4 + 42), int(v4), -2147483648)), fi30, 382), F2: 255}
                                v56 := []int{fi30, fi48}
                                var v57 uint = 65535
                                _ = v55
                                _ = v56
                                _ = v57
                            }
                        default:
                            var v58 uint64 = 18446744073709551615
                            _ = v58
                        }
                        v59 := []int{fi48, fi30, 9223372036854775807, 9223372036854775807}
                        _ = v49
                        _ = v50
                        _ = v51
                        _ = v52
                        _ = v59
                    case "foo bar":
                        v27 -= ((v27 + 0.0) / ((((1e10 + float64(v1)) - 0.1) + v27) / ((-2.25 - v27) * ((float64(v1) - v27) / (1e10 - 0.0)))))
                    case "a":
                        if true {
                            var v60 uint = 781
                            v61 := GenS6{F0: (v60 + v60)}
                            switch "\n\t" {
                            case "foo bar":
                                v62 := GenS4{F0: 255}
                                v63 := GenS4{F0: 42}
                                v64 := v2[fi30]
                                _ = v62
                                _ = v63
                                _ = v64
                            case "a":
                                v60 = v60
                            case "hello":
                                if (v5 > v5) {
                                    v65 := GenS7{F0: int64(fi48), F1: 1}
                                    v66 := GenS8{F0: uint(v3), F1: 0, F2: 1}
                                    _ = v65
                                    _ = v66
                                } else {
                                    fmt.Println(v27, v3)
                                    if true {
                                        v67 := GenS2{F0: 42}
                                        _ = v67
                                    } else {
                                        if true {
                                            v68 := GenS6{F0: 1}
                                            v5 = 0
                                            fmt.Println(v5, v5)
                                            _ = v68
                                        }
                                        switch int(v3) {
                                        case -379:
                                            fmt.Println(v27, fi30)
                                        case 261:
                                            if true {
                                                fmt.Println(fi30, v3)
                                            }
                                            v69 := []int{}
                                            if (!false) {
                                                var v70 string = "\n\t"
                                                if true {
                                                    var v71 uint16 = 42
                                                    v72 := GenS9{F0: -806}
                                                    var v73 int32 = 1
                                                    _ = v71
                                                    _ = v72
                                                    _ = v73
                                                }
                                                _ = v70
                                            }
                                            if false {
                                                fmt.Println(v3)
                                            }
                                            v74 := GenS1{F0: uint8(v60), F1: 0, F2: 32767}
                                            v75 := v2[fi30]
                                            var v76 uint64 = 472
                                            _ = v69
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                        case 4294967295:
                                            if (v5 != v5) {
                                                v77 := GenS9{F0: (^(v4))}
                                                _ = v77
                                            } else {
                                                var v78 uint8 = v3
                                                _ = v78
                                            }
                                            var v79 float32 = (1e10 * -2.25)
                                            switch "\n\t" {
                                            case "hello":
                                                var v80 float32 = ((3.14159 * 3.14159) / 3.14159)
                                                _ = v80
                                            case "\n\t":
                                                if false {
                                                    switch ("hello" + "a") {
                                                    case "hello":
                                                        var v81 int8 = -12
                                                        _ = v81
                                                    default:
                                                        v82 := GenS1{F0: v3, F1: 1, F2: 1}
                                                        _ = v82
                                                    }
                                                    var v83 int32 = 2147483647
                                                    if (v5 < v5) {
                                                        delete(v2, (fi30 + fi48))
                                                    }
                                                    _ = v83
                                                }
                                            }
                                            _ = v79
                                        default:
                                            v84 := []int{42, (-(fi30)), (-(fi30))}
                                            _ = v84
                                        }
                                        var v85 uint32 = (^(v5))
                                        v4 = (v4 + v4)
                                        v86 := GenS10{F0: -292, F1: -2147483648}
                                        v85 += v5
                                        var v87 int8 = 127
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                    }
                                }
                            }
                            _ = v60
                            _ = v61
                        }
                    }
                default:
                    var v88 uint = 0
                    _ = v88
                }
                _ = fi48
            }
        }
        var v89 int16 = 1000
        _ = fi30
        _ = v89
    }
    fmt.Println(v5, v1, v3)
    var v90 float32 = ((((v1 / (v1 * (v1 / (v1 + ((float32(v27) / (v1 + (v1 - ((((1.5 / 1.0) * 3.14159) - (float32(v27) + v1)) / (-2.25 * ((v1 + v1) / ((v1 / v1) + float32(v27)))))))) / (((1.5 / 1e10) * v1) / float32((float64(v1) + float64(3.14159))))))))) + (1.5 * 0.0)) - 0.0) / 0.1)
    if true {
        v91 := GenS1{F0: v3, F1: 1, F2: 1}
        delete(v2, 2147483647)
        _ = v91
    } else {
        v92 := []int{}
        _ = v92
    }
    if true {
        if (-9223372036854775808 != 0) {
            v93 := []int{185}
            _ = v93
        }
        for fi94 := 0; fi94 < 5; fi94++ {
            v90 = v1
            var v95 uint = 99
            _ = fi94
            _ = v95
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v90
}
