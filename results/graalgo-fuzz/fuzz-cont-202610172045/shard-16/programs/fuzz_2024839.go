package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 int8; F1 uint; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int32; F2 int8 }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 int; F1 string; F2 uint16 }
func genFn0(p0 int16, p1 int64, p2 int) int32 {
    v6 := GenS1{F0: 32767, F1: 995}
    var v7 int64 = -221
    v8 := []int{4294967295, (p2 & 127), p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 0
}
func genFn1() uint64 {
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 int32 = 2147483647
        if (!false) {
            fmt.Println(v10)
        } else {
            var v11 float64 = 1e10
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS2{F0: ("hello" + ""), F1: 42}
                _ = fi12
                _ = v13
            }
            v14 := GenS2{F0: ("" + "foo bar"), F1: 4294967295}
            _ = v11
            _ = v14
        }
        _ = fi9
        _ = v10
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int64 = -1
        fmt.Println(fi15)
        v17 := GenS2{F0: ("" + (("" + "") + "foo bar")), F1: 501}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float64 = 3.14159
    fmt.Println(v18)
    if (42 == 1) {
        fmt.Println(v18)
        v19 := []int{}
        if (1 != 218) {
            _ = len(v19)
            var v20 bool = false
            for fi21 := 0; fi21 < 4; fi21++ {
                for fi22 := 0; fi22 < 3; fi22++ {
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = fi22
                }
                _ = len(v19)
                _ = fi21
            }
            _ = v20
        } else {
            if true {
                var v23 int32 = 42
                _ = v23
            }
            var v24 uint8 = 255
            if (false && (2 == 255)) {
                v24 = 59
                var v25 int64 = -356
                if (!(((v18 + 3.14159) - 1e10) < v18)) {
                    var v26 int32 = 0
                    _ = v26
                } else {
                    v27 := GenS3{F0: 42, F1: 125, F2: 131}
                    _ = v27
                }
                _ = v25
            }
            _ = v24
        }
        _ = v19
    }
    _ = v18
    return 0
}

func main() {
    var v1 int8 = 1
    var v2 uint32 = 65535
    v3 := make(map[string]bool)
    v4 := []int{0, -2147483648}
    v5 := make(map[string]bool)
    for fi28 := 0; fi28 < 7; fi28++ {
        v1 += (v1 ^ v1)
        _ = fi28
    }
    switch ("" + "foo bar") {
    case "\n\t":
        for fi29 := 0; fi29 < 20; fi29++ {
            if false {
                for fi30 := 0; fi30 < 1; fi30++ {
                    v31 := v5[""]
                    v31 = (v31 || false)
                    _ = fi30
                    _ = v31
                }
                v1 -= v1
                v32 := GenS1{F0: 492, F1: (v2 % ((v2) | 1))}
                _ = v32
            } else {
                fmt.Println(fi29, fi29, fi29)
            }
            if false {
                fmt.Println(v2, v2, fi29)
            } else {
                switch fi29 {
                case -958:
                    for fi33 := 0; fi33 < 3; fi33++ {
                        v34 := GenS1{F0: 197, F1: (v2 - (v2 >> ((255) & 63)))}
                        _ = fi33
                        _ = v34
                    }
                default:
                    if (float32(float64(((3.14159 / 1e10) / (0.1 / (3.14159 * -2.25))))) >= -2.25) {
                        if (false && false) {
                            v1 = (v1 & (^(v1)))
                            fmt.Println(v1, fi29, v2)
                            switch "\n\t" {
                            case "foo bar":
                                v2 -= (v2 | v2)
                            default:
                                v2 = (v2 >> ((568) & 63))
                            }
                        }
                        v35 := GenS4{F0: v2, F1: 0, F2: (v1 - (v1 * v1))}
                        _ = v35
                    } else {
                        switch fi29 {
                        case -824:
                            v2 -= (v2 * v2)
                        case 141:
                            var v36 int = -2147483648
                            _ = v36
                        case 4294967295:
                            v2 = v2
                        default:
                            v37 := GenS2{F0: "", F1: 1}
                            _ = v37
                        }
                        var v38 uint16 = 160
                        v2 = (v2 * v2)
                        _ = v38
                    }
                }
                var v39 uint64 = 18446744073709551615
                if true {
                    fmt.Println(fi29, v39, v39)
                    v39 = 4294967295
                }
                _ = v39
            }
            _ = fi29
        }
    case "hello":
        for fi40 := 0; fi40 < 14; fi40++ {
            fmt.Println(v1)
            _ = fi40
        }
    case "foo bar":
        var v41 int32 = -2147483648
        _ = v41
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        var v43 string = "foo bar"
        v1 = (-(v1))
        _ = fi42
        _ = v43
    }
    v1 = v1
    fmt.Println(v2)
    var v44 uint64 = 4294967295
    var v45 int32 = -2147483648
    v45 = genFn0(-225, 9223372036854775807, 1)
    for fi46 := 0; fi46 < 1; fi46++ {
        if (!(!true)) {
            _ = len(v4)
            var v47 int8 = v1
            _ = v47
        } else {
            v48 := GenS2{F0: "hello", F1: 434}
            _ = v48
        }
        v49 := GenS1{F0: 0, F1: 72}
        var v50 bool = (v45 != (v45 | 42))
        _ = fi46
        _ = v49
        _ = v50
    }
    if ((false && (128 >= 0)) && true) {
        if (1 != 261) {
            v51 := GenS3{F0: v1, F1: 326, F2: v44}
            var v52 bool = false
            _ = v51
            _ = v52
        } else {
            var v53 string = "foo bar"
            if true {
                v1 = (v1 ^ 42)
            }
            _ = v53
        }
        switch ("hello" + "a") {
        case "a":
            v54 := v3["foo bar"]
            _ = v54
        default:
            switch -54 {
            case -2147483648:
                _ = len(v4)
            case -9223372036854775808:
                v55 := GenS4{F0: (v2 << ((651) & 63)), F1: (v45 << ((425) & 63)), F2: (v1 * 69)}
                _ = v55
            case 2147483647:
                _ = len(v5)
            default:
                if (4294967295 >= 666) {
                    fmt.Println(v45)
                } else {
                    switch 128 {
                    case 399:
                        var v56 int = 42
                        _ = v56
                    }
                    if false {
                        v57 := v3["a"]
                        var v58 uint = 255
                        switch "foo bar" {
                        case "hello":
                            if false {
                                var v59 int8 = -128
                                switch "hello" {
                                case "\n\t":
                                    v60 := GenS1{F0: 528, F1: (v2 >> ((18446744073709551615) & 63))}
                                    _ = v60
                                default:
                                    switch 127 {
                                    case 1:
                                        v61 := GenS5{F0: "a", F1: -128}
                                        _ = v61
                                    }
                                }
                                _ = v59
                            } else {
                                v62 := GenS6{F0: 32767, F1: (v58 >> (((v58 ^ 255)) & 63))}
                                v63 := GenS2{F0: (("foo bar" + ("hello" + "\n\t")) + "hello"), F1: 2}
                                v64 := GenS7{F0: 151, F1: "a", F2: 1}
                                _ = v62
                                _ = v63
                                _ = v64
                            }
                        default:
                            v1 -= v1
                        }
                        _ = v57
                        _ = v58
                    } else {
                        var v65 int64 = -897
                        v66 := GenS7{F0: 4294967295, F1: "\n\t", F2: 32767}
                        var v67 uint = 0
                        _ = v65
                        _ = v66
                        _ = v67
                    }
                    for fi68 := 0; fi68 < 5; fi68++ {
                        var v69 int32 = int32(v1)
                        if (89 != 143) {
                            var v70 int64 = 1
                            _ = v70
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            for fi71 := 0; fi71 < 4; fi71++ {
                                v1 = v1
                                fmt.Println(v45, v69, fi71)
                                v1 = -127
                                v2 = (v2 * (v2 & v2))
                                _ = fi71
                            }
                            switch "hello" {
                            case "\n\t":
                                for fi72 := 0; fi72 < 8; fi72++ {
                                    var v73 bool = (3.14159 <= -2.25)
                                    var v74 float64 = 1e10
                                    v44 -= v44
                                    _ = fi72
                                    _ = v73
                                    _ = v74
                                }
                            case "":
                                for fi75 := 0; fi75 < 4; fi75++ {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v76 := GenS6{F0: 32767, F1: 0}
                                    _ = fi75
                                    _ = v76
                                }
                            case "a":
                                var v77 int64 = 81
                                _ = v77
                            }
                        }
                        _ = fi68
                        _ = v69
                    }
                }
            }
        }
    } else {
        v2 = 327
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
}
