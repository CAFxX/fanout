package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint8; F2 int64 }
type GenS5 struct { F0 string; F1 int }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 int32 }
func genFn0(p0 int16) uint32 {
    if (!(-195 >= 2147483647)) {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 uint8 = 58
            for fi7 := 0; fi7 < 19; fi7++ {
                fmt.Println(v6)
                var v8 int64 = 0
                fmt.Println(p0)
                _ = fi7
                _ = v8
            }
            _ = fi5
            _ = v6
        }
        p0 -= p0
        v9 := []int{32768, 871, 571, 4294967295}
        _ = v9
    } else {
        if ((!true) && ((true && false) && (308 <= 1))) {
            p0 += (p0 * 1000)
            for fi10 := 0; fi10 < 4; fi10++ {
                p0 = p0
                switch ("" + (("foo bar" + (("\n\t" + (("foo bar" + (("a" + "a") + "foo bar")) + "")) + "hello")) + "\n\t")) {
                case "hello":
                    fmt.Println(fi10, p0)
                case "\n\t":
                    var v11 uint32 = 0
                    _ = v11
                case "":
                    switch (fi10 % ((2147483647) | 1)) {
                    case 42:
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi10, fi10)
                        case "":
                            var v12 uint8 = 42
                            _ = v12
                        case "hello":
                            p0 -= (p0 / ((p0) | 1))
                        default:
                            var v13 uint8 = 42
                            _ = v13
                        }
                    default:
                        var v14 uint = 261
                        _ = v14
                    }
                default:
                    var v15 float32 = float32(1e10)
                    _ = v15
                }
                _ = fi10
            }
        }
        v16 := GenS1{F0: "\n\t", F1: false}
        fmt.Println(p0)
        _ = v16
    }
    p0 = 555
    p0 -= (p0 * p0)
    var v17 int8 = -1
    _ = p0
    _ = v17
    return 1
}
func genFn1(p0 int8, p1 bool) uint {
    p0 += p0
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS1{F0: "hello", F1: ((-2147483648 != -1) && (p1 || false))}
            switch fi19 {
            case -9223372036854775808:
                v21 := GenS1{F0: "foo bar", F1: (fi18 == 2147483647)}
                _ = v21
            case -1:
                p1 = (true || (p1 || p1))
            case 4294967295:
                switch fi18 {
                case -1:
                    v22 := GenS1{F0: ("\n\t" + "foo bar"), F1: (4294967295 < 4294967295)}
                    _ = v22
                case 128:
                    switch fi18 {
                    case 381:
                        var v23 uint = 0
                        _ = v23
                    default:
                        if (2 > 405) {
                            var v24 int32 = 0
                            var v25 string = "hello"
                            _ = v24
                            _ = v25
                        }
                        switch (("hello" + "hello") + "\n\t") {
                        case "":
                            if p1 {
                                if (42 == 18446744073709551615) {
                                    v26 := []int{}
                                    fmt.Println(p1, p1)
                                    p0 = 84
                                    v27 := GenS1{F0: ("foo bar" + "hello"), F1: false}
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 uint8 = 128
                                    _ = v28
                                }
                                if p1 {
                                    var v29 int16 = 1
                                    _ = v29
                                }
                            } else {
                                p1 = ((((!true) || (0 > 2)) && (!(65535 == 205))) && p1)
                                v30 := GenS2{F0: (p1 || false), F1: 49}
                                var v31 int32 = -2147483648
                                _ = v30
                                _ = v31
                            }
                        case "a":
                            v32 := GenS1{F0: "foo bar", F1: p1}
                            _ = v32
                        default:
                            p0 += -128
                        }
                    }
                    v33 := GenS3{F0: (fi19 % (((fi19 >> ((1) & 63))) | 1))}
                    v34 := GenS1{F0: "\n\t", F1: p1}
                    _ = v33
                    _ = v34
                case 1:
                    if (-2147483648 < -1) {
                        v35 := []int{(fi18 - (fi19 % ((2147483647) | 1))), fi18, fi18}
                        var v36 bool = ((-2.25 + 3.14159) > 0.1)
                        var v37 uint8 = 1
                        _ = v35
                        _ = v36
                        _ = v37
                    } else {
                        fmt.Println(fi18, p1, p1)
                        p0 = p0
                        v38 := GenS3{F0: fi19}
                        var v39 string = ("" + ("" + ("\n\t" + (("a" + "foo bar") + "\n\t"))))
                        v40 := []int{(fi18 >> ((18446744073709551615) & 63)), fi19}
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                }
            default:
                v41 := []int{}
                _ = v41
            }
            _ = fi19
            _ = v20
        }
        p1 = (32767 == 1)
        _ = fi18
    }
    p0 = (p0 + (p0 ^ p0))
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint8 = 42
    var v2 string = "a"
    v3 := make(map[string]uint32)
    v4 := []string{"hello", v2, v2, v2}
    delete(v3, (((v2 + ((v2 + "foo bar") + v2)) + "hello") + ("foo bar" + (((v2 + ("a" + v2)) + "hello") + (v2 + v2)))))
    if (98 != 0) {
        _ = len(v3)
    } else {
        var v42 float32 = (0.1 + 1.5)
        _ = v42
    }
    v1 -= (^(v1))
    v1 -= v1
    v43 := GenS3{F0: 0}
    if (!(65535 == 255)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (false && true) {
            v3[v2] = 42
            if (((0.1 * 0.0) / (-2.25 / 1.0)) < (float64(1e10) - 0.0)) {
                var v44 uint16 = 255
                v44 = 628
                _ = v44
            } else {
                var v45 uint16 = 587
                v1 = v1
                var v46 int16 = 42
                _ = v45
                _ = v46
            }
        }
        if false {
            for fi47 := 0; fi47 < 4; fi47++ {
                v1 = 1
                if true {
                    if false {
                        v48 := v3["a"]
                        if (!(!(false || false))) {
                            switch ("\n\t" + (v2 + (v2 + ("hello" + "foo bar")))) {
                            case "foo bar":
                                v48 = (v48 % (((v48 & (v48 ^ v48))) | 1))
                            }
                        }
                        _ = v48
                    }
                }
                var v49 string = ("" + v2)
                _ = fi47
                _ = v49
            }
            var v50 int8 = -82
            _ = v50
        } else {
            if (-1 == 1000) {
                for fi51 := 0; fi51 < 15; fi51++ {
                    _ = len(v3)
                    for fi52 := 0; fi52 < 3; fi52++ {
                        fmt.Println(fi52)
                        if true {
                            if (1e10 == (1.5 + 0.1)) {
                                v1 += v1
                            }
                            if (((0.1 * -2.25) > 0.0) && (4294967295 > 0)) {
                                var v53 uint32 = 4294967295
                                _ = v53
                            } else {
                                var v54 int16 = 1000
                                v55 := GenS1{F0: "\n\t", F1: (!true)}
                                v1 = 42
                                _ = v54
                                _ = v55
                            }
                            v56 := GenS2{F0: false, F1: v1}
                            var v57 string = v2
                            var v58 int8 = -128
                            _ = v56
                            _ = v57
                            _ = v58
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        v59 := GenS2{F0: (!(0 != 2147483647)), F1: (v1 + (v1 / (((^(v1))) | 1)))}
                        _ = fi52
                        _ = v59
                    }
                    _ = fi51
                }
            } else {
                if false {
                    v2 = "\n\t"
                    var v60 int32 = -1
                    if true {
                        switch v2 {
                        case "a":
                            v1 += v1
                        case "":
                            for fi61 := 0; fi61 < 10; fi61++ {
                                var v62 float64 = 1.5
                                _ = fi61
                                _ = v62
                            }
                        case "\n\t":
                            _ = len(v3)
                        }
                        var v63 int = -575
                        _ = v63
                    }
                    _ = v60
                }
            }
            fmt.Println(v2, v2)
        }
    } else {
        var v64 uint = genFn1(0, true)
        if (116 < 0) {
            v4 = append(v4, (v2 + ""))
            v65 := GenS4{F0: (v2 + ("" + (((((v2 + (v2 + (v2 + v2))) + ((v2 + v2) + "a")) + v2) + v2) + ("foo bar" + ((v2 + v2) + (v2 + "a")))))), F1: uint8(v64), F2: -50}
            for fi66 := 0; fi66 < 1; fi66++ {
                fmt.Println(v64, v2, v2)
                var v67 uint8 = 113
                switch v2 {
                case "":
                    if (genFn0(-352) <= uint32(v67)) {
                        if (!((0 < 452) || true)) {
                            switch fi66 {
                            case 499:
                                for fi68 := 0; fi68 < 16; fi68++ {
                                    var v69 uint16 = 65535
                                    if false {
                                        v64 = (v64 >> (((^(v64))) & 63))
                                    }
                                    switch v2 {
                                    case "hello":
                                        if len(v4) > 0 {
                                            v4[0] = "a"
                                        }
                                    }
                                    _ = fi68
                                    _ = v69
                                }
                            }
                        }
                        var v70 uint8 = 193
                        _ = v70
                    }
                case "hello":
                    v71 := GenS5{F0: v2, F1: fi66}
                    _ = v71
                case "a":
                    var v72 int16 = -32768
                    _ = v72
                default:
                    for fi73 := 0; fi73 < 16; fi73++ {
                        var v74 int32 = -2147483648
                        if true {
                            if false {
                                if false {
                                    v75 := GenS6{F0: false, F1: (v64 ^ genFn1(1, true))}
                                    _ = v75
                                } else {
                                    v76 := v3[v2]
                                    if false {
                                        fmt.Println(fi73)
                                    }
                                    _ = v76
                                }
                                if (32767 < 47) {
                                    fmt.Println(v67, v64, v64)
                                }
                            }
                            v77 := GenS7{F0: v74}
                            switch 128 {
                            case 42:
                                v2 = (v2 + v2)
                            default:
                                var v78 int8 = 1
                                _ = v78
                            }
                            _ = v77
                        }
                        _ = fi73
                        _ = v74
                    }
                }
                _ = fi66
                _ = v67
            }
            _ = v65
        } else {
            var v79 uint8 = (v1 - v1)
            _ = v79
        }
        _ = v64
    }
    var v80 uint = 255
    for fi81 := 0; fi81 < 16; fi81++ {
        v82 := v3[v2]
        _ = fi81
        _ = v82
    }
    fmt.Println(v2, v80, v2)
    var v83 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v80:", v80)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v80
    _ = v83
}
