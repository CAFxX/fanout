package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS3 struct { F0 uint; F1 int; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64; F1 int }
func genFn0(p0 uint64, p1 uint8) int8 {
    var v4 int16 = 32767
    fmt.Println(p1, p1, p0)
    for fi5 := 0; fi5 < 4; fi5++ {
        if ((false && true) || false) {
            for fi6 := 0; fi6 < 12; fi6++ {
                p1 = 204
                switch -1 {
                case 1:
                    fmt.Println(fi5)
                case -9223372036854775808:
                    fmt.Println(p1, v4)
                case 42:
                    fmt.Println(p0, fi5, p1)
                default:
                    if false {
                        var v7 int64 = 42
                        _ = v7
                    } else {
                        var v8 string = ""
                        v8 = "\n\t"
                        v8 = (v8 + "a")
                        if false {
                            v4 += v4
                            var v9 uint = 411
                            fmt.Println(fi5)
                            _ = v9
                        } else {
                            v10 := []int{(fi5 % ((fi6) | 1))}
                            p1 -= 128
                            _ = v10
                        }
                        v11 := []int{(fi5 & (fi6 << ((718) & 63))), fi6, (fi5 * fi6), -914}
                        _ = v11
                        _ = v8
                    }
                }
                _ = fi6
            }
            p0 = p0
            var v12 uint16 = 255
            _ = v12
        } else {
            var v13 int16 = v4
            fmt.Println(v4, v13, p1)
            switch "a" {
            case "foo bar":
                var v14 float64 = 1.5
                _ = v14
            default:
                v15 := GenS1{F0: -1, F1: 872, F2: -18}
                _ = v15
            }
            _ = v13
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            if (false || true) {
                var v17 bool = true
                var v18 float32 = ((((((1.5 - ((float32(float64(float32((1e10 / (3.14159 - 1.5))))) - 3.14159) / 3.14159)) / 3.14159) - ((3.14159 - (-2.25 - 1.5)) * (0.0 / 1.0))) + (1.5 + (1.5 * 0.1))) / ((((0.1 * 1.5) - 1e10) * (1e10 + (1e10 - (-2.25 / 1.0)))) - (0.1 - (0.1 / 1.0)))) / 0.1)
                var v19 bool = v17
                _ = v17
                _ = v18
                _ = v19
            } else {
                v4 = -32768
                p1 = p1
                p1 = (p1 ^ 1)
            }
            var v20 int16 = -32768
            _ = fi16
            _ = v20
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return 1
}
func genFn1(p0 uint16, p1 int32) int16 {
    switch ("" + "\n\t") {
    case "":
        if (true && (1 >= 919)) {
            p0 = 255
            for fi21 := 0; fi21 < 12; fi21++ {
                var v22 uint16 = 32767
                _ = fi21
                _ = v22
            }
        }
    case "\n\t":
        v23 := []int{0, 0, -431}
        _ = v23
    case "hello":
        for fi24 := 0; fi24 < 15; fi24++ {
            var v25 uint32 = 0
            v26 := []int{(fi24 ^ (fi24 | (^(fi24))))}
            if true {
                if (-92 >= int64(fi24)) {
                    _ = len(v26)
                } else {
                    if len(v26) > 0 {
                        v26[0] = fi24
                    }
                }
                fmt.Println(fi24, p0)
            } else {
                var v27 uint32 = v25
                v25 = v27
                if len(v26) >= 2 {
                    v26 = v26[1:len(v26)]
                }
                _ = v27
            }
            _ = fi24
            _ = v25
            _ = v26
        }
    default:
        v28 := []int{-9223372036854775808, -952}
        _ = v28
    }
    if true {
        p0 -= p0
        var v29 int = -2147483648
        _ = v29
    } else {
        v30 := GenS1{F0: 100, F1: 42, F2: p1}
        if (true && (true && true)) {
            for fi31 := 0; fi31 < 4; fi31++ {
                switch "foo bar" {
                case "foo bar":
                    if true {
                        var v32 uint = 0
                        _ = v32
                    } else {
                        var v33 uint16 = (p0 * 255)
                        _ = v33
                    }
                case "":
                    var v34 uint64 = 18446744073709551615
                    _ = v34
                case "hello":
                    fmt.Println(p0)
                default:
                    switch "hello" {
                    case "hello":
                        if false {
                            switch (("hello" + ("a" + "\n\t")) + "foo bar") {
                            case "foo bar":
                                v35 := GenS2{F0: 2, F1: 237, F2: 127}
                                _ = v35
                            default:
                                if false {
                                    p1 = p1
                                }
                            }
                        }
                    case "foo bar":
                        v36 := GenS3{F0: 149, F1: fi31, F2: (!true)}
                        _ = v36
                    default:
                        v37 := []int{(fi31 | fi31), fi31, (fi31 ^ (fi31 | fi31)), (fi31 ^ fi31)}
                        _ = v37
                    }
                }
                _ = fi31
            }
        }
        v38 := []int{}
        _ = v30
        _ = v38
    }
    _ = p0
    _ = p1
    return 481
}

func main() {
    var v1 bool = false
    v2 := []uint32{2147483648, 42, 42, 65535}
    v3 := make(map[int]string)
    _ = len(v3)
    var v39 float64 = 1e10
    fmt.Println(v1)
    v1 = (-1 <= 9223372036854775807)
    for fi40 := 0; fi40 < 5; fi40++ {
        fmt.Println(v1, v39)
        _ = fi40
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        if (4294967295 != 68) {
            if ((!v1) || (42 != 235)) {
                v42 := GenS4{F0: 2}
                if (((0 <= 22) || v1) && false) {
                    if (0 < 393) {
                        fmt.Println(v39)
                        v1 = false
                        v43 := GenS2{F0: 65535, F1: 42, F2: 42}
                        _ = v43
                    }
                    if true {
                        var v44 string = "\n\t"
                        _ = v44
                    }
                    v2 = append(v2, 2147483647)
                }
                _ = v42
            }
            var v45 uint = 797
            _ = v45
        } else {
            switch fi41 {
            case 590:
                v46 := GenS3{F0: 890, F1: (fi41 - fi41), F2: (190 > 128)}
                _ = v46
            default:
                fmt.Println(v1)
            }
            var v47 bool = v1
            _ = v47
        }
        if true {
            var v48 uint64 = 4294967295
            if false {
                if v1 {
                    var v49 float64 = (((v39 * v39) * 0.0) + v39)
                    if v1 {
                        fmt.Println(v1, v48)
                        var v50 int64 = -1
                        _ = v50
                    }
                    var v51 uint16 = 2
                    _ = v49
                    _ = v51
                } else {
                    delete(v3, fi41)
                    var v52 int8 = 0
                    v48 = v48
                    _ = v52
                }
            } else {
                fmt.Println(v48, v1, fi41)
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    v3[(-(fi41))] = "\n\t"
                case "":
                    v48 = (v48 * 1)
                }
            }
            for fi53 := 0; fi53 < 18; fi53++ {
                v1 = v1
                if true {
                    v54 := GenS4{F0: 466}
                    v1 = v1
                    if (42 >= 42) {
                        var v55 int32 = int32(v48)
                        var v56 uint8 = 128
                        v39 -= v39
                        _ = v55
                        _ = v56
                    }
                    v1 = (!(false || v1))
                    _ = v54
                }
                v1 = v1
                _ = fi53
            }
            _ = v48
        } else {
            v57 := GenS5{F0: -596, F1: (fi41 & 1), F2: (-(fi41))}
            _ = v57
        }
        var v58 int64 = -381
        _ = fi41
        _ = v58
    }
    for fi59 := 0; fi59 < 2; fi59++ {
        if (false && (4294967295 > 18446744073709551615)) {
            v39 = v39
            switch (fi59 - fi59) {
            case 611:
                if false {
                    for fi60 := 0; fi60 < 9; fi60++ {
                        delete(v3, (fi59 / ((fi59) | 1)))
                        v61 := GenS4{F0: 122}
                        if (v1 || (("foo bar" + (("foo bar" + "\n\t") + "a")) == "")) {
                            if true {
                                if false {
                                    v62 := v3[(fi60 << ((1) & 63))]
                                    if len(v2) > 0 {
                                        v2[0] = 693
                                    }
                                    v63 := GenS6{F0: fi60}
                                    _ = v62
                                    _ = v63
                                }
                                v39 = ((v39 / v39) * v39)
                                if (!v1) {
                                    var v64 uint8 = 42
                                    v39 += (((3.14159 / float64(1.5)) - 0.0) / ((1e10 + v39) / (((v39 - (1e10 - 0.0)) / ((v39 + (((v39 * v39) - v39) * -2.25)) - (v39 / float64(1.5)))) + float64((3.14159 - ((1e10 * ((0.0 * 3.14159) / 0.1)) / 3.14159))))))
                                    var v65 int16 = 25
                                    _ = v64
                                    _ = v65
                                } else {
                                    v1 = (!false)
                                }
                            } else {
                                if (1 <= 18446744073709551615) {
                                    if v1 {
                                        v1 = (false && (0 <= 32767))
                                        if (!true) {
                                            var v66 uint16 = 42
                                            fmt.Println(fi60, v1)
                                            v2 = append(v2, 2147483648)
                                            _ = v66
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 65535
                                        }
                                    } else {
                                        v67 := v3[fi59]
                                        _ = v67
                                    }
                                } else {
                                    if v1 {
                                        if (!v1) {
                                            var v68 float64 = (v39 + float64(0.1))
                                            if v1 {
                                                v68 += v68
                                            } else {
                                                v69 := GenS7{F0: 2}
                                                v68 += 1e10
                                                _ = v69
                                            }
                                            v70 := GenS5{F0: 9223372036854775807, F1: (fi60 >> ((255) & 63)), F2: (fi59 >> ((18446744073709551615) & 63))}
                                            _ = v68
                                            _ = v70
                                        } else {
                                            v71 := GenS2{F0: 65535, F1: -32768, F2: 127}
                                            _ = v71
                                        }
                                        switch fi60 {
                                        case 1:
                                            v1 = (!(!(265 >= 2)))
                                        case 42:
                                            v1 = ((fi59 / ((fi60) | 1)) == (^(fi59)))
                                        }
                                    }
                                    var v72 float32 = -2.25
                                    var v73 int64 = 42
                                    _ = v72
                                    _ = v73
                                }
                            }
                        } else {
                            v39 = 1e10
                            v1 = v1
                        }
                        var v74 uint64 = 0
                        if false {
                            var v75 uint = 1
                            _ = v75
                        } else {
                            var v76 float32 = (((-2.25 - float32(v39)) / 1.0) * 0.0)
                            v76 += 3.14159
                            _ = v76
                        }
                        _ = fi60
                        _ = v61
                        _ = v74
                    }
                    var v77 bool = (!v1)
                    _ = v77
                }
            default:
                if true {
                    for fi78 := 0; fi78 < 18; fi78++ {
                        if (52 < 0) {
                            v79 := GenS7{F0: 0}
                            var v80 uint = 678
                            _ = v79
                            _ = v80
                        } else {
                            if (v1 && v1) {
                                v39 = 1e10
                                if true {
                                    var v81 uint16 = 108
                                    fmt.Println(v1)
                                    _ = len(v3)
                                    _ = v81
                                } else {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                if ((4294967295 <= 2) || ((!(9223372036854775807 > -9223372036854775808)) || v1)) {
                                    v39 = 3.14159
                                    v82 := GenS4{F0: 86}
                                    var v83 float32 = 0.1
                                    _ = v82
                                    _ = v83
                                } else {
                                    var v84 uint8 = 187
                                    _ = v84
                                }
                                fmt.Println(fi78, v39)
                            }
                            v85 := GenS3{F0: 65535, F1: 1, F2: v1}
                            v86 := GenS8{F0: 9223372036854775807, F1: (fi78 ^ -9223372036854775808)}
                            _ = v85
                            _ = v86
                        }
                        var v87 bool = v1
                        _ = fi78
                        _ = v87
                    }
                    var v88 uint64 = 42
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v88
                }
            }
        }
        _ = fi59
    }
    var v89 int64 = -9223372036854775808
    v90 := v3[42]
    fmt.Println(v39)
    for fi91 := 0; fi91 < 8; fi91++ {
        _ = len(v3)
        v90 = v90
        for fi92 := 0; fi92 < 19; fi92++ {
            v93 := GenS2{F0: 1, F1: genFn1(22, -2147483648), F2: 127}
            _ = fi92
            _ = v93
        }
        _ = fi91
    }
    delete(v3, -9223372036854775808)
    for fi94 := 0; fi94 < 15; fi94++ {
        var v95 float32 = 0.0
        _ = fi94
        _ = v95
    }
    v39 = float64(0.0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v89
    _ = v90
}
