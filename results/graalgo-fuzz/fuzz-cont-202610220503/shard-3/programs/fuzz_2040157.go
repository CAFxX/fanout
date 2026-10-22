package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint16; F1 bool; F2 int64 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int64; F1 int32; F2 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 uint16; F2 int64 }
func genFn0(p0 string) uint8 {
    var v4 float64 = 0.0
    switch ("" + ((p0 + p0) + (p0 + p0))) {
    case "a":
        v5 := GenS1{F0: 1, F1: true, F2: 4294967295}
        _ = v5
    case "foo bar":
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := []int{fi6, fi6, fi6, (fi6 % ((-9223372036854775808) | 1))}
            fmt.Println(fi6, fi6, fi6)
            p0 = ("hello" + p0)
            _ = fi6
            _ = v7
        }
    case "":
        var v8 uint16 = 65535
        _ = v8
    }
    _ = p0
    _ = v4
    return 128
}
func genFn1() uint8 {
    if (42 > -248) {
        var v9 uint = 65535
        for fi10 := 0; fi10 < 15; fi10++ {
            if ((false || false) || true) {
                fmt.Println(v9, v9)
                for fi11 := 0; fi11 < 3; fi11++ {
                    switch "foo bar" {
                    case "":
                        var v12 string = "\n\t"
                        v12 = v12
                        _ = v12
                    default:
                        v9 = (v9 % ((255) | 1))
                    }
                    v13 := GenS1{F0: 1, F1: (495 >= 238), F2: v9}
                    _ = fi11
                    _ = v13
                }
            } else {
                if (110 == int8(v9)) {
                    v14 := GenS1{F0: 32767, F1: false, F2: v9}
                    var v15 bool = ((!(634 == -9223372036854775808)) || false)
                    if v15 {
                        var v16 uint16 = 0
                        v9 = v9
                        v17 := []int{}
                        _ = v16
                        _ = v17
                    } else {
                        v18 := GenS2{F0: (fi10 >> (((v9 << ((v9) & 63))) & 63)), F1: 42, F2: 0}
                        v19 := GenS1{F0: uint16(v9), F1: false, F2: v9}
                        switch fi10 {
                        case 906:
                            var v20 int64 = -1
                            _ = v20
                        case 32767:
                            for fi21 := 0; fi21 < 4; fi21++ {
                                v9 += v9
                                switch "a" {
                                case "\n\t":
                                    v22 := GenS3{F0: 65535, F1: "foo bar"}
                                    _ = v22
                                case "a":
                                    v23 := GenS3{F0: 658, F1: ""}
                                    _ = v23
                                default:
                                    v24 := []int{fi10, (fi21 / ((fi10) | 1)), fi21, -2147483648}
                                    fmt.Println(v15)
                                    var v25 float32 = ((3.14159 - 0.1) * (float32(3.14159) / ((0.1 - 1.5) - -2.25)))
                                    var v26 string = "hello"
                                    _ = v24
                                    _ = v25
                                    _ = v26
                                }
                                _ = fi21
                            }
                        case -277:
                            var v27 int8 = 0
                            _ = v27
                        default:
                            v28 := GenS4{F0: 42, F1: ((1000 > -887) || (4294967295 <= 2147483647)), F2: 1}
                            _ = v28
                        }
                        _ = v18
                        _ = v19
                    }
                    _ = v14
                    _ = v15
                } else {
                    switch "" {
                    case "\n\t":
                        for fi29 := 0; fi29 < 10; fi29++ {
                            var v30 uint16 = 22
                            if true {
                                switch "foo bar" {
                                case "\n\t":
                                    var v31 uint16 = 32767
                                    _ = v31
                                case "hello":
                                    switch "hello" {
                                    case "hello":
                                        v32 := GenS2{F0: 42, F1: -2147483648, F2: 342}
                                        _ = v32
                                    }
                                default:
                                    v30 -= (v30 + v30)
                                    fmt.Println(v30)
                                    v33 := GenS2{F0: fi29, F1: -112, F2: 2147483647}
                                    _ = v33
                                }
                            }
                            _ = fi29
                            _ = v30
                        }
                    case "hello":
                        v34 := []int{(fi10 >> ((v9) & 63)), (fi10 - fi10), fi10}
                        _ = v34
                    case "foo bar":
                        if ((false || (true && true)) && false) {
                            var v35 uint8 = 167
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(v9, v35)
                            case "":
                                v36 := []int{(-(fi10)), fi10, fi10}
                                _ = v36
                            }
                            _ = v35
                        } else {
                            var v37 int = (fi10 ^ fi10)
                            v38 := []int{(^(fi10)), -1, (v37 / (((v37 & v37)) | 1))}
                            if len(v38) >= 2 {
                                v38 = v38[1:len(v38)]
                            }
                            _ = v37
                            _ = v38
                        }
                    }
                    for fi39 := 0; fi39 < 13; fi39++ {
                        var v40 uint = 65535
                        _ = fi39
                        _ = v40
                    }
                    v41 := GenS3{F0: 0, F1: ("foo bar" + ("a" + "foo bar"))}
                    _ = v41
                }
            }
            if false {
                var v42 bool = true
                v43 := []int{(fi10 & -463), fi10, 42, (fi10 * fi10)}
                if true {
                    var v44 int = 368
                    v45 := GenS2{F0: (v44 / (((fi10 ^ -153)) | 1)), F1: 148, F2: 42}
                    var v46 float64 = 3.14159
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    v47 := GenS2{F0: fi10, F1: -2147483648, F2: 42}
                    _ = v47
                }
                _ = v42
                _ = v43
            } else {
                v9 += v9
                var v48 uint16 = 0
                _ = v48
            }
            _ = fi10
        }
        if true {
            v9 = v9
            v9 -= 332
            v9 = (v9 / ((616) | 1))
        }
        _ = v9
    }
    var v49 float32 = (float32(0.0) * 1e10)
    var v50 int8 = -1
    if ("a" >= (("hello" + "foo bar") + "")) {
        switch 0 {
        case 127:
            if true {
                v50 = (v50 / (((v50 << ((4294967295) & 63))) | 1))
            }
        case 0:
            v51 := []int{}
            _ = v51
        case -9223372036854775808:
            var v52 uint64 = 42
            _ = v52
        default:
            v53 := GenS1{F0: 255, F1: (!true), F2: 42}
            _ = v53
        }
        v54 := []int{0}
        _ = v54
    }
    _ = v49
    _ = v50
    return 128
}

func main() {
    v1 := []string{("a" + "a"), "a", (("" + ("hello" + "")) + "hello"), "a"}
    v2 := make(map[uint32]string)
    var v3 int = 1
    v55 := GenS2{F0: (v3 ^ v3), F1: 42, F2: 65535}
    v56 := GenS4{F0: 65535, F1: (!(!true)), F2: 1}
    v3 -= v3
    for fi57 := 0; fi57 < 2; fi57++ {
        for fi58 := 0; fi58 < 11; fi58++ {
            v3 = (fi57 | (fi58 + fi58))
            var v59 float32 = (-2.25 + -2.25)
            switch -518 {
            case 409:
                v1 = append(v1, ("a" + ""))
            case 127:
                v1 = append(v1, (("foo bar" + "foo bar") + ("a" + "\n\t")))
            case 24:
                v60 := GenS2{F0: (fi57 ^ (fi58 | v3)), F1: 851, F2: 555}
                _ = v60
            }
            _ = fi58
            _ = v59
        }
        switch ((("a" + ("hello" + ("" + "a"))) + "hello") + "foo bar") {
        case "foo bar":
            var v61 float64 = (1.5 / (0.1 / -2.25))
            _ = v61
        case "":
            v3 -= fi57
        case "\n\t":
            switch (("" + "a") + "foo bar") {
            case "foo bar":
                for fi62 := 0; fi62 < 5; fi62++ {
                    v63 := GenS3{F0: 175, F1: ""}
                    fmt.Println(v3, fi62, fi57)
                    var v64 uint = 1
                    if (1 == 1) {
                        v3 = 9223372036854775807
                        if len(v1) > 0 {
                            v1[0] = ""
                        }
                        fmt.Println(v64)
                        _ = len(v1)
                        v65 := GenS1{F0: 65535, F1: (false && (255 == genFn1())), F2: (v64 | v64)}
                        _ = v65
                    } else {
                        v66 := v2[923]
                        v2[412] = "\n\t"
                        _ = len(v2)
                        _ = v66
                    }
                    _ = fi62
                    _ = v63
                    _ = v64
                }
            default:
                if true {
                    var v67 uint64 = 1
                    if true {
                        v3 += (v3 + (v3 * fi57))
                        v3 = v3
                        fmt.Println(v3, fi57)
                    }
                    v68 := GenS5{F0: ((32767 >= -1) || (!true)), F1: 2}
                    _ = v67
                    _ = v68
                } else {
                    v1 = append(v1, "foo bar")
                    if true {
                        v69 := GenS6{F0: 275, F1: 1, F2: 0}
                        _ = v69
                    } else {
                        if false {
                            v3 = v3
                        }
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                        for fi70 := 0; fi70 < 17; fi70++ {
                            fmt.Println(fi57, fi57, fi57)
                            fmt.Println(fi70, v3, v3)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi70
                        }
                    }
                    var v71 uint8 = 191
                    _ = v71
                }
            }
        default:
            var v72 int = v3
            _ = v72
        }
        _ = fi57
    }
    v3 += 128
    var v73 int16 = 32767
    for fi74 := 0; fi74 < 14; fi74++ {
        v3 = v3
        _ = fi74
    }
    fmt.Println(v3, v3)
    fmt.Println(v73)
    for fi75 := 0; fi75 < 13; fi75++ {
        v3 = (v3 << ((484) & 63))
        var v76 uint8 = 42
        _ = fi75
        _ = v76
    }
    v77 := GenS7{F0: "foo bar"}
    switch "a" {
    case "foo bar":
        fmt.Println(v3)
    case "a":
        if (18446744073709551615 < 65535) {
            if (-1 == -381) {
                fmt.Println(v73)
            }
        }
    }
    v73 -= v73
    var v78 string = ("\n\t" + "foo bar")
    for fi79 := 0; fi79 < 6; fi79++ {
        var v80 uint8 = 2
        _ = fi79
        _ = v80
    }
    var v81 int64 = 524
    for fi82 := 0; fi82 < 8; fi82++ {
        var v83 int32 = 2147483647
        switch v3 {
        case 32768:
            v3 += v3
        default:
            var v84 int16 = (v73 - v73)
            _ = v84
        }
        if (-1 == 100) {
            _ = len(v1)
            var v85 uint64 = 570
            _ = v85
        }
        _ = fi82
        _ = v83
    }
    v86 := GenS8{F0: 854, F1: 65535, F2: (v81 >> ((42) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v73:", v73)
    fmt.Println("v78:", v78)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
    _ = v56
    _ = v73
    _ = v77
    _ = v78
    _ = v81
    _ = v86
}
