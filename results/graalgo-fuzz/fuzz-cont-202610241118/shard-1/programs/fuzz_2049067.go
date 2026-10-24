package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS2 struct { F0 int8; F1 int64; F2 int32 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8; F1 int64; F2 int64 }
type GenS6 struct { F0 uint; F1 bool; F2 int8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 int }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 uint; F1 uint16 }
type GenS11 struct { F0 int8; F1 int8; F2 uint64 }
type GenS12 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int, p1 bool, p2 int8) int {
    var v4 string = "hello"
    if false {
        for fi5 := 0; fi5 < 2; fi5++ {
            for fi6 := 0; fi6 < 13; fi6++ {
                p1 = false
                v7 := []int{(-(fi6)), (^(p0)), -430, 0}
                var v8 uint16 = 32767
                v9 := GenS1{F0: 2, F1: 88, F2: -157}
                var v10 uint64 = 18446744073709551615
                _ = fi6
                _ = v10
                _ = v7
                _ = v8
                _ = v9
            }
            switch 32768 {
            case 2147483647:
                v11 := GenS1{F0: 255, F1: 1, F2: -1}
                _ = v11
            case 4294967295:
                v12 := []int{}
                _ = v12
            }
            var v13 uint8 = 255
            _ = fi5
            _ = v13
        }
        v14 := []int{}
        _ = v14
    } else {
        var v15 int32 = 0
        for fi16 := 0; fi16 < 16; fi16++ {
            if p1 {
                for fi17 := 0; fi17 < 2; fi17++ {
                    switch 2147483647 {
                    case 1:
                        v18 := GenS2{F0: (p2 << ((95) & 63)), F1: 146, F2: 42}
                        _ = v18
                    default:
                        p2 = p2
                    }
                    fmt.Println(p2)
                    _ = fi17
                }
            } else {
                if p1 {
                    var v19 uint16 = 228
                    for fi20 := 0; fi20 < 3; fi20++ {
                        p2 = 42
                        v21 := GenS3{F0: 0, F1: int64(v19)}
                        v22 := GenS4{F0: (int(v15) > fi20)}
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS5{F0: 2, F1: -405, F2: -208}
                    _ = v19
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 9; fi24++ {
                p1 = ((42 > 0) && (32767 >= -492))
                _ = fi24
            }
            _ = fi16
        }
        _ = v15
    }
    var v25 uint8 = 234
    p0 += (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v4
    return p0
}
func genFn1(p0 string) uint32 {
    v26 := GenS6{F0: 42, F1: false, F2: 1}
    switch -780 {
    case 0:
        var v27 uint16 = 2
        _ = v27
    case 265:
        fmt.Println(p0)
    }
    var v28 string = ((p0 + (((p0 + "") + p0) + (p0 + p0))) + ("a" + p0))
    if (-992 <= 350) {
        var v29 int = 609
        _ = v29
    }
    _ = p0
    _ = v26
    _ = v28
    return 1
}

func main() {
    v1 := []int{4294967295, -12}
    var v2 int64 = -9223372036854775808
    var v3 int16 = 1
    var v30 uint = 18446744073709551615
    fmt.Println(v30, v30)
    if len(v1) > 0 {
        v1[0] = 42
    }
    switch 104 {
    case 42:
        v2 = (v2 ^ -142)
    case 0:
        var v31 uint = 298
        _ = v31
    case -955:
        v2 = (v2 >> (((v30 & v30)) & 63))
    default:
        switch ("" + ((((("hello" + ("a" + "hello")) + "") + "hello") + "foo bar") + "a")) {
        case "":
            var v32 int = 42
            _ = v32
        case "\n\t":
            fmt.Println(v30, v30)
        default:
            var v33 int16 = 828
            _ = v33
        }
    }
    if false {
        v34 := GenS7{F0: ("\n\t" + "\n\t")}
        _ = v34
    }
    switch ("foo bar" + "a") {
    case "":
        v35 := GenS6{F0: v30, F1: true, F2: -1}
        _ = v35
    case "\n\t":
        for fi36 := 0; fi36 < 13; fi36++ {
            v3 -= 578
            if true {
                if len(v1) > 0 {
                    v1[0] = fi36
                }
                if false {
                    v2 = (v2 << ((v30) & 63))
                    fmt.Println(v2, v3, fi36)
                }
            }
            _ = fi36
        }
    default:
        var v37 float64 = float64(0.0)
        _ = v37
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    v38 := GenS1{F0: 49, F1: 255, F2: 2147483647}
    if true {
        switch "a" {
        case "a":
            for fi39 := 0; fi39 < 3; fi39++ {
                for fi40 := 0; fi40 < 5; fi40++ {
                    var v41 uint16 = 87
                    _ = fi40
                    _ = v41
                }
                for fi42 := 0; fi42 < 11; fi42++ {
                    if (v2 <= -769) {
                        if false {
                            fmt.Println(fi42)
                        } else {
                            if ((v3 / (((v3 / ((v3) | 1))) | 1)) >= (v3 - (v3 / (((v3 / (((v3 | 240)) | 1))) | 1)))) {
                                v3 = v3
                                switch (^(fi42)) {
                                case -9223372036854775808:
                                    var v43 int8 = 127
                                    _ = v43
                                case -1:
                                    fmt.Println(v2)
                                    switch 576 {
                                    case -9223372036854775808:
                                        var v44 uint64 = 44
                                        _ = v44
                                    case 157:
                                        v30 -= (v30 | v30)
                                    case -2147483648:
                                        if (true && true) {
                                            v45 := GenS6{F0: v30, F1: (false && (127 >= -128)), F2: 22}
                                            var v46 uint8 = 135
                                            v47 := GenS3{F0: v30, F1: (-(v2))}
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        }
                                    default:
                                        switch "hello" {
                                        case "a":
                                            switch fi42 {
                                            case 216:
                                                var v48 string = ("a" + ("" + "foo bar"))
                                                _ = v48
                                            }
                                        case "foo bar":
                                            switch "\n\t" {
                                            case "foo bar":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case "a":
                                                v30 -= 18446744073709551615
                                            case "":
                                                v49 := GenS8{F0: 42, F1: 127}
                                                _ = v49
                                            }
                                        default:
                                            var v50 uint8 = 95
                                            _ = v50
                                        }
                                    }
                                    var v51 uint = (^(v30))
                                    _ = v51
                                case 381:
                                    if len(v1) > 0 {
                                        v1[0] = (^(fi42))
                                    }
                                }
                                v2 += (^(v2))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v52 := GenS5{F0: 200, F1: (-(v2)), F2: (v2 * v2)}
                            var v53 int32 = 0
                            _ = v52
                            _ = v53
                        }
                        v1 = append(v1, genFn0(fi39, (0 != 111), 1))
                        var v54 int8 = -122
                        _ = v54
                    } else {
                        if false {
                            switch "" {
                            case "\n\t":
                                if false {
                                    if (65535 >= 65535) {
                                        v55 := GenS8{F0: uint16(v30), F1: genFn0(-186, (!(127 < 44)), -72)}
                                        var v56 uint8 = 99
                                        v2 -= (v2 / ((v2) | 1))
                                        v30 = 247
                                        _ = v55
                                        _ = v56
                                    }
                                } else {
                                    _ = len(v1)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                        }
                        var v57 uint16 = 1
                        if (272 < v2) {
                            v58 := GenS6{F0: (v30 >> (((v30 * v30)) & 63)), F1: false, F2: 42}
                            _ = v58
                        } else {
                            var v59 uint32 = 42
                            if (!(false && true)) {
                                var v60 uint = 6
                                v61 := GenS6{F0: 65535, F1: false, F2: -18}
                                if len(v1) > 0 {
                                    v1[0] = 128
                                }
                                _ = v60
                                _ = v61
                            }
                            v59 += (v59 | (v59 - v59))
                            if true {
                                v30 += 255
                                if (!(((v3 == v3) && false) && false)) {
                                    v62 := GenS2{F0: 100, F1: v2, F2: 2147483647}
                                    fmt.Println(v3, v30, v57)
                                    v63 := GenS9{F0: (v3 << ((v30) & 63)), F1: (1e10 <= float64((3.14159 + 0.1)))}
                                    v59 = uint32(v57)
                                    v2 = (v2 << ((900) & 63))
                                    if false {
                                        var v64 string = ("" + ((("\n\t" + "") + "") + "foo bar"))
                                        switch v64 {
                                        case "":
                                            v65 := GenS3{F0: 758, F1: v2}
                                            _ = v65
                                        case "foo bar":
                                            var v66 uint32 = v59
                                            _ = v66
                                        default:
                                            var v67 int64 = (-(v2))
                                            _ = v67
                                        }
                                        if false {
                                            fmt.Println(fi39, fi39, v2)
                                            if ((v59 * (^(v59))) <= v59) {
                                                var v68 int = (fi39 - (fi42 + (fi39 * 2147483647)))
                                                v30 += (v30 + v30)
                                                _ = v68
                                            }
                                        }
                                        _ = v64
                                    }
                                    _ = v62
                                    _ = v63
                                }
                                fmt.Println(v57, v57, v59)
                            } else {
                                var v69 int8 = -114
                                if (0 <= 373) {
                                    fmt.Println(v3, v2, v2)
                                    var v70 uint32 = (v59 - (v59 | v59))
                                    _ = v70
                                }
                                v71 := GenS10{F0: 65535, F1: v57}
                                fmt.Println(v2)
                                _ = v69
                                _ = v71
                            }
                            _ = v59
                        }
                        _ = v57
                    }
                    v72 := GenS9{F0: v3, F1: false}
                    switch (("" + "hello") + (("hello" + "a") + "a")) {
                    case "hello":
                        v3 = int16(v30)
                    case "foo bar":
                        v73 := GenS11{F0: -128, F1: 4, F2: 0}
                        _ = v73
                    case "\n\t":
                        if (0.0 > 1.5) {
                            if true {
                                switch (fi42 - genFn0(genFn0((fi42 & 2147483647), (-1 >= 0), 42), (((!true) && true) || false), -13)) {
                                case -579:
                                    var v74 int16 = (v3 >> (((v30 / ((1) | 1))) & 63))
                                    _ = v74
                                }
                            } else {
                                var v75 string = ""
                                if (0.0 >= -2.25) {
                                    v76 := GenS9{F0: (v3 >> (((v30 | (v30 / (((v30 / (((v30 % ((v30) | 1))) | 1))) | 1)))) & 63)), F1: false}
                                    v77 := GenS12{F0: uint32(v30), F1: 996}
                                    var v78 uint8 = 128
                                    _ = v76
                                    _ = v77
                                    _ = v78
                                } else {
                                    switch fi39 {
                                    case 4294967295:
                                        var v79 float64 = (0.0 - 1.5)
                                        _ = v79
                                    default:
                                        fmt.Println(fi39)
                                    }
                                }
                                var v80 float64 = 0.0
                                _ = v75
                                _ = v80
                            }
                        }
                    default:
                        v3 = v3
                    }
                    _ = fi42
                    _ = v72
                }
                var v81 string = ("" + ("\n\t" + "hello"))
                _ = fi39
                _ = v81
            }
        case "\n\t":
            var v82 int32 = -940
            _ = v82
        default:
            fmt.Println(v3, v2, v30)
        }
        var v83 uint32 = 980
        _ = v83
    }
    v2 += (v2 / ((v2) | 1))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v38
}
