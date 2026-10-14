package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 bool }
type GenS2 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint8; F2 int8 }
type GenS5 struct { F0 int16; F1 int8; F2 int }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int64; F1 int32; F2 string }
type GenS10 struct { F0 int8; F1 int8 }
func genFn0(p0 uint32) bool {
    v5 := []int{127, 2147483647}
    if true {
        p0 -= (^(p0))
        switch 1 {
        case -1:
            if ("hello" < ("a" + "foo bar")) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
        case 0:
            var v6 uint16 = 1
            _ = v6
        default:
            _ = len(v5)
        }
    } else {
        _ = len(v5)
        p0 = (p0 % ((p0) | 1))
        p0 = (p0 ^ p0)
    }
    p0 -= (p0 << ((493) & 63))
    fmt.Println(p0)
    _ = p0
    _ = v5
    return false
}
func genFn1() int16 {
    var v7 bool = (121 != 1000)
    v8 := []int{-773}
    v7 = v7
    if ((255 >= 1) && (1e10 < 3.14159)) {
        v8 = append(v8, -133)
        v7 = ((("" + (("hello" + "") + ("foo bar" + "hello"))) != ("foo bar" + "a")) || v7)
        if (18446744073709551615 > 4294967295) {
            v9 := GenS1{F0: -2147483648, F1: 0, F2: v7}
            _ = v9
        } else {
            var v10 int8 = 42
            v8 = append(v8, int(v10))
            _ = v10
        }
    }
    if len(v8) > 0 {
        v8[0] = 32767
    }
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 897
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    v11 := GenS2{F0: 789, F1: 36, F2: 1}
    for fi12 := 0; fi12 < 10; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := GenS3{F0: 935}
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    if (-2147483648 < 1) {
        v15 := GenS4{F0: 2, F1: 134, F2: 1}
        fmt.Println(v2)
        switch "" {
        case "":
            v2 += 1
        case "hello":
            if (!((0.1 + ((0.0 * ((0.1 * (((0.1 * (0.1 - -2.25)) + 0.1) * 0.0)) * 1.5)) / ((0.0 + 0.0) - (0.0 + 1.5)))) <= 1.5)) {
                switch "" {
                case "":
                    v2 = v2
                case "\n\t":
                    for fi16 := 0; fi16 < 20; fi16++ {
                        v2 += v2
                        fmt.Println(v2)
                        _ = fi16
                    }
                case "a":
                    v2 += v2
                }
                var v17 int64 = -834
                v17 = (v17 << ((172) & 63))
                _ = v17
            } else {
                v2 -= (v2 * v2)
                v2 = (v2 - genFn1())
            }
        case "\n\t":
            switch 1 {
            case -2147483648:
                for fi18 := 0; fi18 < 17; fi18++ {
                    for fi19 := 0; fi19 < 1; fi19++ {
                        var v20 bool = false
                        var v21 uint32 = 359
                        if (187 == 888) {
                            fmt.Println(v21, fi18, v20)
                            delete(v4, v21)
                            if ("foo bar" != "") {
                                _ = len(v3)
                            }
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v22 := GenS5{F0: (v2 >> ((255) & 63)), F1: 127, F2: fi18}
                    _ = fi18
                    _ = v22
                }
            case 2147483647:
                var v23 uint8 = 1
                _ = v23
            case 127:
                for fi24 := 0; fi24 < 14; fi24++ {
                    if true {
                        var v25 uint32 = 1
                        switch "" {
                        case "foo bar":
                            v2 += int16(v25)
                        case "a":
                            v26 := []int{64}
                            _ = v26
                        case "":
                            if (1 == 2) {
                                for fi27 := 0; fi27 < 11; fi27++ {
                                    if true {
                                        var v28 uint32 = v25
                                        v29 := GenS4{F0: 415, F1: 42, F2: -10}
                                        v25 -= v28
                                        v30 := GenS6{F0: 0, F1: (((("foo bar" + (("foo bar" + "a") + (("a" + "hello") + "foo bar"))) + "\n\t") + ("foo bar" + ("hello" + "a"))) + ("" + "hello"))}
                                        switch ("\n\t" + "hello") {
                                        case "a":
                                            v1[(fi27 << ((18446744073709551615) & 63))] = (fi24 | (-(fi27)))
                                        }
                                        _ = v28
                                        _ = v29
                                        _ = v30
                                    }
                                    fmt.Println(v2, fi24, fi27)
                                    v25 += v25
                                    _ = fi27
                                }
                                var v31 int = (fi24 + fi24)
                                v25 += 543
                                _ = v31
                            }
                        }
                        var v32 uint16 = 32767
                        _ = v25
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v2 += -824
                        _ = fi33
                    }
                    switch 9223372036854775807 {
                    case 32767:
                        fmt.Println(fi24, fi24)
                    }
                    _ = fi24
                }
            default:
                var v34 int8 = 1
                _ = v34
            }
        }
        _ = v15
    }
    if true {
        v2 += (v2 * (v2 ^ 0))
        fmt.Println(v2)
    } else {
        v2 = (v2 / (((v2 / (((v2 ^ genFn1())) | 1))) | 1))
        v35 := []int{0}
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 float64 = -2.25
            _ = fi36
            _ = v37
        }
        _ = v35
    }
    v38 := []int{2147483647, 127}
    switch -2147483648 {
    case -2147483648:
        var v39 bool = (516 == 788)
        _ = v39
    case 2147483647:
        v2 = (^(v2))
    case 400:
        v40 := GenS2{F0: 0, F1: 2, F2: 127}
        _ = v40
    }
    for fi41 := 0; fi41 < 3; fi41++ {
        var v42 uint = 255
        var v43 int = int(v2)
        _ = fi41
        _ = v42
        _ = v43
    }
    v44 := v4[2147483648]
    fmt.Println(v2, v44)
    v45 := GenS7{F0: 982}
    var v46 float64 = 1.5
    switch "foo bar" {
    case "hello":
        if true {
            var v47 uint8 = 125
            _ = v47
        } else {
            if (v46 <= v46) {
                v48 := v4[42]
                fmt.Println(v48, v2)
                switch v44 {
                case -959:
                    for fi49 := 0; fi49 < 10; fi49++ {
                        var v50 int8 = 42
                        v51 := GenS8{F0: v44, F1: false}
                        v48 -= v44
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                case 813:
                    if (!true) {
                        v52 := GenS9{F0: -9223372036854775808, F1: -165, F2: "hello"}
                        v2 = v2
                        v38 = append(v38, -1)
                        _ = v52
                    } else {
                        v2 -= (-(v2))
                        v48 = (v48 | (v48 - v48))
                    }
                case 71:
                    for fi53 := 0; fi53 < 20; fi53++ {
                        if false {
                            fmt.Println(v48, v2, v48)
                            switch "foo bar" {
                            case "hello":
                                var v54 int16 = v2
                                _ = v54
                            case "foo bar":
                                v2 += 0
                            default:
                                if ((true || (!true)) && false) {
                                    var v55 int = (v44 & v48)
                                    _ = v55
                                } else {
                                    v2 += (-(v2))
                                    switch (fi53 - 1) {
                                    case -9223372036854775808:
                                        if len(v38) > 0 {
                                            v38[0] = v48
                                        }
                                    case 42:
                                        switch "" {
                                        case "foo bar":
                                            v56 := GenS3{F0: 125}
                                            _ = v56
                                        case "hello":
                                            _ = len(v1)
                                        case "":
                                            for fi57 := 0; fi57 < 17; fi57++ {
                                                if ("" != (("a" + "hello") + ("a" + "foo bar"))) {
                                                    if len(v38) > 0 {
                                                        v38[0] = (fi53 & fi57)
                                                    }
                                                } else {
                                                    switch fi57 {
                                                    case 0:
                                                        v38 = append(v38, v48)
                                                    case -582:
                                                        if true {
                                                            var v58 uint16 = 32767
                                                            _ = v58
                                                        } else {
                                                            var v59 bool = false
                                                            fmt.Println(v48)
                                                            _ = v59
                                                        }
                                                    default:
                                                        var v60 uint32 = 474
                                                        _ = v60
                                                    }
                                                }
                                                var v61 uint32 = 2147483647
                                                v46 -= v46
                                                _ = fi57
                                                _ = v61
                                            }
                                        default:
                                            v46 = float64(-2.25)
                                        }
                                    case -424:
                                        for fi62 := 0; fi62 < 19; fi62++ {
                                            if false {
                                                if (false || false) {
                                                    var v63 int8 = 60
                                                    fmt.Println(v44, v46)
                                                    var v64 float32 = (float32(v46) + 3.14159)
                                                    _ = v63
                                                    _ = v64
                                                } else {
                                                    var v65 string = ("a" + "\n\t")
                                                    var v66 uint64 = 182
                                                    _ = v65
                                                    _ = v66
                                                }
                                            } else {
                                                var v67 bool = false
                                                v68 := GenS10{F0: 1, F1: 119}
                                                _ = v67
                                                _ = v68
                                            }
                                            _ = fi62
                                        }
                                    default:
                                        v48 += (v44 & (^(v48)))
                                    }
                                    var v69 float32 = 1.5
                                    _ = v69
                                }
                            }
                        } else {
                            if len(v38) > 0 {
                                v38[0] = (v44 - v48)
                            }
                            for fi70 := 0; fi70 < 17; fi70++ {
                                var v71 int32 = 0
                                fmt.Println(v46, v46, v48)
                                _ = fi70
                                _ = v71
                            }
                            var v72 uint8 = 2
                            _ = v72
                        }
                        var v73 uint64 = 290
                        var v74 string = "\n\t"
                        _ = fi53
                        _ = v73
                        _ = v74
                    }
                default:
                    v46 = 0.0
                }
                _ = v48
            }
            var v75 uint32 = 4
            fmt.Println(v2)
            _ = v75
        }
    case "foo bar":
        for fi76 := 0; fi76 < 20; fi76++ {
            var v77 uint32 = 865
            _ = fi76
            _ = v77
        }
    case "a":
        var v78 uint32 = 1
        _ = v78
    }
    v79 := GenS3{F0: 2147483648}
    for fi80 := 0; fi80 < 14; fi80++ {
        for fi81 := 0; fi81 < 10; fi81++ {
            v46 += v46
            _ = fi81
        }
        v1[(v44 * (^(fi80)))] = (fi80 | v44)
        var v82 int16 = (v2 / ((1000) | 1))
        _ = fi80
        _ = v82
    }
    fmt.Println(v46)
    v44 -= (-(v44))
    var v83 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v79
    _ = v83
}
