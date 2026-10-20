package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint64; F2 string }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int8; F2 uint32 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 bool; F2 bool }
type GenS10 struct { F0 uint32; F1 string }
type GenS11 struct { F0 uint16; F1 int64 }
type GenS12 struct { F0 uint32; F1 int64 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 int; F1 uint64; F2 int32 }
func genFn0(p0 int16, p1 uint32) int32 {
    v5 := []int{}
    fmt.Println(p1)
    var v6 int64 = -1
    v7 := GenS1{F0: (v6 > v6)}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int8, p1 uint, p2 int8) int {
    v8 := []int{4294967295}
    v8 = append(v8, -1)
    v9 := GenS1{F0: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]bool)
    v3 := []int{1, 42, 0, 382}
    var v4 float64 = v1
    v10 := GenS2{F0: 218, F1: 2, F2: ("hello" + "")}
    for fi11 := 0; fi11 < 8; fi11++ {
        v4 = 1e10
        _ = fi11
    }
    switch 9223372036854775807 {
    case -605:
        var v12 float64 = 1e10
        _ = v12
    case -1:
        var v13 int8 = 42
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        fmt.Println(v4, v1)
        _ = fi14
    }
    if ((v1 + (float64(-2.25) + v1)) >= 1e10) {
        fmt.Println(v1, v4)
    } else {
        v4 = 0.0
        v15 := GenS3{F0: 1, F1: 2147483647}
        _ = v15
    }
    if (879 > 42) {
        switch "a" {
        case "":
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 bool = true
                v1 += float64(3.14159)
                _ = fi16
                _ = v17
            }
        case "foo bar":
            _ = len(v3)
        default:
            var v18 uint8 = 255
            _ = v18
        }
        var v19 bool = true
        v4 -= v4
        _ = v19
    } else {
        if (2 == 1) {
            var v20 uint8 = 127
            if (uint64(v20) > 371) {
                if (!(!((((-2.25 > 1.5) || true) || true) || true))) {
                    if true {
                        v20 += v20
                        v3 = append(v3, genFn1(-29, 18446744073709551615, 100))
                        var v21 uint64 = 421
                        _ = v21
                    }
                    for fi22 := 0; fi22 < 2; fi22++ {
                        for fi23 := 0; fi23 < 1; fi23++ {
                            v20 = (v20 - 210)
                            v24 := GenS4{F0: 0}
                            if false {
                                v25 := GenS5{F0: -32768, F1: -1, F2: 2147483647}
                                v1 -= (((v4 + (3.14159 + v1)) * (((1.5 / v1) * 0.1) + v1)) - v1)
                                _ = v25
                            } else {
                                if ((3.14159 + (((0.0 + (-2.25 * (-2.25 - ((-2.25 * 0.0) + -2.25)))) / 3.14159) + 1.5)) != float32(v1)) {
                                    v20 += (v20 >> ((255) & 63))
                                    switch 9223372036854775807 {
                                    case 2147483647:
                                        switch ((("\n\t" + "a") + "a") + ("" + "foo bar")) {
                                        case "a":
                                            var v26 uint64 = 209
                                            _ = v26
                                        case "\n\t":
                                            fmt.Println(v4)
                                        default:
                                            v4 = ((0.1 / -2.25) * v1)
                                        }
                                        fmt.Println(fi23, fi22)
                                    default:
                                        delete(v2, 42)
                                    }
                                }
                            }
                            v1 = (v1 + v1)
                            v20 = (v20 * v20)
                            v27 := GenS6{F0: 255, F1: 42}
                            v28 := GenS1{F0: true}
                            _ = fi23
                            _ = v24
                            _ = v27
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 20; fi29++ {
                            var v30 uint8 = v20
                            v31 := v2[0]
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        _ = fi22
                    }
                }
                for fi32 := 0; fi32 < 11; fi32++ {
                    var v33 uint32 = 517
                    for fi34 := 0; fi34 < 6; fi34++ {
                        var v35 uint8 = 43
                        _ = fi34
                        _ = v35
                    }
                    for fi36 := 0; fi36 < 10; fi36++ {
                        var v37 int64 = -9223372036854775808
                        _ = fi36
                        _ = v37
                    }
                    _ = fi32
                    _ = v33
                }
                switch "" {
                case "":
                    var v38 int = 32768
                    _ = v38
                }
            } else {
                v1 += float64(-2.25)
            }
            _ = v20
        } else {
            v39 := GenS6{F0: 2, F1: 100}
            _ = v39
        }
        var v40 uint = 65535
        _ = v40
    }
    if ((!false) || ((0 != 2) && (119 < 255))) {
        switch ("foo bar" + ("hello" + "\n\t")) {
        case "\n\t":
            var v41 uint16 = 65535
            _ = v41
        case "a":
            for fi42 := 0; fi42 < 1; fi42++ {
                switch genFn1(0, 18446744073709551615, -128) {
                case 658:
                    switch "a" {
                    case "foo bar":
                        v4 += 1.5
                    case "":
                        var v43 string = "hello"
                        _ = v43
                    default:
                        fmt.Println(fi42)
                    }
                case -299:
                    v44 := GenS3{F0: 32767, F1: 65535}
                    _ = v44
                }
                v2[69] = (-1 >= genFn0(0, 4294967295))
                var v45 uint32 = 442
                _ = fi42
                _ = v45
            }
        default:
            var v46 bool = (-549 != 42)
            _ = v46
        }
        v47 := GenS7{F0: (false && (genFn1(-78, 650, -128) == 127)), F1: 4294967295}
        if true {
            fmt.Println(v1)
            switch 4294967295 {
            case 127:
                var v48 float32 = 0.1
                _ = v48
            default:
                if true {
                    if false {
                        fmt.Println(v4, v4)
                        var v49 float64 = (3.14159 / 3.14159)
                        _ = v49
                    }
                    var v50 string = ""
                    if (false && (-915 <= 1)) {
                        v2[2] = (!(!false))
                    }
                    _ = v50
                }
            }
            v4 = (float64((((3.14159 / 1.0) - (float32(float64(1e10)) / (float32((v1 + v4)) / (1e10 - (float32(-2.25) - float32((-2.25 * v4))))))) * ((((3.14159 - (0.1 + (0.0 / -2.25))) * -2.25) - (1e10 + (1.5 + float32(v4)))) / 1.0))) + v1)
        } else {
            var v51 uint16 = 1
            for fi52 := 0; fi52 < 18; fi52++ {
                if false {
                    v1 += 3.14159
                    if (0.1 >= 1e10) {
                        fmt.Println(v51)
                        if false {
                            v4 -= (float64(-2.25) + (v1 - 1.5))
                            for fi53 := 0; fi53 < 16; fi53++ {
                                var v54 uint64 = 616
                                var v55 uint8 = 1
                                v56 := GenS8{F0: 65535}
                                _ = fi53
                                _ = v54
                                _ = v55
                                _ = v56
                            }
                            v57 := GenS3{F0: 530, F1: 4294967295}
                            _ = v57
                        } else {
                            v58 := GenS6{F0: uint(v51), F1: -128}
                            _ = v58
                        }
                        if true {
                            v59 := GenS9{F0: 97, F1: false, F2: (0 > 4294967295)}
                            _ = v59
                        } else {
                            if false {
                                for fi60 := 0; fi60 < 4; fi60++ {
                                    if true {
                                        v61 := GenS10{F0: 4294967295, F1: "foo bar"}
                                        v62 := GenS8{F0: 117}
                                        v63 := GenS11{F0: (v51 << ((376) & 63)), F1: 42}
                                        v64 := GenS10{F0: 638, F1: "foo bar"}
                                        var v65 int32 = -1
                                        _ = v61
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                    }
                                    v4 += 0.0
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi60
                                }
                                var v66 string = "foo bar"
                                _ = v66
                            }
                            var v67 float32 = 1.5
                            _ = v67
                        }
                    } else {
                        switch ("foo bar" + "a") {
                        case "hello":
                            var v68 string = ""
                            _ = v68
                        case "foo bar":
                            if true {
                                switch "hello" {
                                case "":
                                    for fi69 := 0; fi69 < 12; fi69++ {
                                        v51 -= 32767
                                        switch fi52 {
                                        case 764:
                                            if false {
                                                if false {
                                                    var v70 uint8 = 118
                                                    if (!(v51 >= 1)) {
                                                        var v71 int32 = 1
                                                        var v72 float64 = v4
                                                        var v73 uint8 = v70
                                                        _ = v71
                                                        _ = v72
                                                        _ = v73
                                                    }
                                                    _ = v70
                                                }
                                            } else {
                                                v1 = v1
                                            }
                                        case 128:
                                            var v74 float64 = -2.25
                                            v75 := GenS12{F0: 313, F1: -1}
                                            _ = v74
                                            _ = v75
                                        case 854:
                                            if true {
                                                if (v4 >= 1e10) {
                                                    v3 = append(v3, (fi69 << ((0) & 63)))
                                                } else {
                                                    var v76 int64 = 178
                                                    _ = v76
                                                }
                                            }
                                        default:
                                            v2[958] = (42 <= 2)
                                        }
                                        _ = fi69
                                    }
                                case "hello":
                                    v1 += 0.1
                                case "a":
                                    if len(v3) > 0 {
                                        v3[0] = (fi52 - (^(fi52)))
                                    }
                                }
                            } else {
                                var v77 uint8 = 0
                                _ = v77
                            }
                        case "":
                            v78 := GenS7{F0: false, F1: 42}
                            _ = v78
                        default:
                            for fi79 := 0; fi79 < 10; fi79++ {
                                switch fi79 {
                                case 2147483647:
                                    var v80 string = "hello"
                                    _ = v80
                                case 32768:
                                    fmt.Println(fi52, fi52, v4)
                                default:
                                    if (false && false) {
                                        v81 := GenS1{F0: false}
                                        var v82 float64 = v1
                                        _ = v81
                                        _ = v82
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = -1
                                        }
                                        delete(v2, 250)
                                    }
                                }
                                _ = fi79
                            }
                        }
                        v4 = 3.14159
                        switch "foo bar" {
                        case "\n\t":
                            var v83 int16 = 546
                            _ = v83
                        case "foo bar":
                            if (!true) {
                                var v84 int64 = 1
                                for fi85 := 0; fi85 < 4; fi85++ {
                                    var v86 uint16 = v51
                                    if ((18446744073709551615 <= 4294967295) && (42 != 2)) {
                                        var v87 int32 = 58
                                        var v88 float32 = 1.5
                                        fmt.Println(v84)
                                        _ = v87
                                        _ = v88
                                    }
                                    _ = len(v2)
                                    _ = fi85
                                    _ = v86
                                }
                                if true {
                                    v4 += (-2.25 + v4)
                                } else {
                                    var v89 int = fi52
                                    _ = len(v2)
                                    fmt.Println(v84, v51)
                                    _ = v89
                                }
                                _ = v84
                            } else {
                                if (-967 == 1) {
                                    v90 := GenS13{F0: 117}
                                    _ = v90
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v91 string = ""
                                    _ = v91
                                }
                            }
                        case "hello":
                            v51 = v51
                        }
                    }
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi52
            }
            _ = v51
        }
        _ = v47
    } else {
        v92 := GenS14{F0: 42, F1: 4294967295, F2: 206}
        var v93 int32 = 0
        _ = v92
        _ = v93
    }
    var v94 float64 = ((float64((1.5 * (1.5 / (-2.25 * (1.5 / (3.14159 / float32(v1))))))) + float64(3.14159)) * (((v1 - (1.5 + (((((float64(3.14159) + -2.25) + v1) / 1.0) + (float64(float32((v1 + (v4 / v1)))) + v4)) * 0.1))) * v1) * 1e10))
    if (!(42 < 0)) {
        for fi95 := 0; fi95 < 13; fi95++ {
            _ = len(v2)
            _ = fi95
        }
    } else {
        if (207 <= 86) {
            if true {
                if false {
                    v4 += v4
                    var v96 uint16 = 837
                    _ = v96
                }
            }
        }
        if false {
            v1 += ((v94 / float64(0.1)) / 1.0)
        }
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            v1 = float64(0.1)
        }
    }
    var v97 int32 = -2147483648
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v94:", v94)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v94
    _ = v97
}
