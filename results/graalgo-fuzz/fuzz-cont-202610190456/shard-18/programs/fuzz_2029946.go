package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int }
type GenS2 struct { F0 bool; F1 uint32; F2 int }
type GenS3 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 bool; F1 int8; F2 string }
type GenS6 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS7 struct { F0 string; F1 uint16; F2 uint8 }
type GenS8 struct { F0 uint32; F1 string; F2 int32 }
type GenS9 struct { F0 int16; F1 int64 }
type GenS10 struct { F0 uint; F1 uint32 }
type GenS11 struct { F0 bool; F1 int64; F2 int8 }
type GenS12 struct { F0 int8; F1 uint16; F2 uint64 }
func genFn0(p0 uint) uint16 {
    for fi4 := 0; fi4 < 18; fi4++ {
        fmt.Println(p0, p0)
        p0 = p0
        _ = fi4
    }
    switch 32767 {
    case 1:
        p0 += 255
    case -587:
        v5 := []int{-9223372036854775808, -83, 1}
        _ = v5
    default:
        switch -9223372036854775808 {
        case -464:
            if (((0.1 + ((float64((1.5 * 3.14159)) * 3.14159) / -2.25)) + 0.0) == (1e10 + (-2.25 - 0.0))) {
                var v6 float32 = 1.5
                _ = v6
            }
        case -9223372036854775808:
            v7 := GenS1{F0: 1, F1: 32768, F2: 585}
            _ = v7
        case 702:
            v8 := GenS2{F0: false, F1: 2, F2: 32767}
            _ = v8
        default:
            if true {
                v9 := []int{}
                _ = v9
            } else {
                switch "a" {
                case "a":
                    fmt.Println(p0)
                case "\n\t":
                    for fi10 := 0; fi10 < 16; fi10++ {
                        fmt.Println(fi10, fi10)
                        v11 := []int{fi10, -449}
                        p0 = 774
                        _ = fi10
                        _ = v11
                    }
                default:
                    for fi12 := 0; fi12 < 9; fi12++ {
                        v13 := GenS2{F0: false, F1: 1, F2: (fi12 ^ (fi12 << ((65535) & 63)))}
                        var v14 int64 = 42
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                }
            }
        }
    }
    p0 += p0
    _ = p0
    return 217
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{9223372036854775807, 498, 435, 4294967295}
    v3 := []string{("a" + ("" + (("a" + ("hello" + "a")) + "\n\t")))}
    var v15 int8 = 127
    v15 -= 105
    v16 := GenS3{F0: 861, F1: 4294967295, F2: 724}
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            if false {
                var v19 uint64 = 613
                v20 := GenS1{F0: fi17, F1: (fi18 >> ((2) & 63)), F2: 9223372036854775807}
                _ = v19
                _ = v20
            }
            _ = fi18
        }
        if true {
            v21 := GenS4{F0: 0, F1: 1}
            _ = v21
        }
        v22 := GenS1{F0: 200, F1: (fi17 % (((fi17 / (((-(fi17))) | 1))) | 1)), F2: fi17}
        _ = fi17
        _ = v22
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(fi23)
        var v24 float64 = 1e10
        var v25 uint8 = 49
        _ = fi23
        _ = v24
        _ = v25
    }
    switch "a" {
    case "hello":
        v26 := GenS2{F0: false, F1: 836, F2: 42}
        _ = v26
    case "":
        v27 := GenS2{F0: ((-674 == 160) || true), F1: 454, F2: 1}
        _ = v27
    case "a":
        var v28 uint8 = 183
        _ = v28
    }
    var v29 uint16 = 42
    switch -2147483648 {
    case -749:
        var v30 int8 = -1
        _ = v30
    case 112:
        for fi31 := 0; fi31 < 8; fi31++ {
            if (843 != 585) {
                _ = len(v2)
                fmt.Println(v15)
                if false {
                    switch "a" {
                    case "foo bar":
                        for fi32 := 0; fi32 < 9; fi32++ {
                            v15 = v15
                            v15 = v15
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi32
                        }
                    }
                    v1[(("" + ("a" + ("" + ("hello" + (("" + "a") + ("" + ("" + "hello"))))))) + "foo bar")] = true
                }
            }
            v33 := GenS5{F0: true, F1: (v15 | v15), F2: ""}
            _ = fi31
            _ = v33
        }
    case -1:
        v15 += v15
    default:
        var v34 int8 = 1
        _ = v34
    }
    if (-30 >= 1000) {
        for fi35 := 0; fi35 < 4; fi35++ {
            v15 = -1
            var v36 uint16 = v29
            _ = fi35
            _ = v36
        }
        if len(v2) > 0 {
            v2[0] = 560
        }
    } else {
        if false {
            v37 := GenS3{F0: (v29 ^ (v29 % ((65535) | 1))), F1: 0, F2: -2147483648}
            v38 := GenS2{F0: false, F1: 65535, F2: 128}
            v15 += 58
            _ = v37
            _ = v38
        }
        v29 = v29
    }
    v39 := GenS6{F0: 128, F1: 65535, F2: (v29 | (v29 - genFn0(6)))}
    switch ("" + "") {
    case "foo bar":
        if false {
            v15 -= v15
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v40 int16 = 777
            _ = v40
        }
    case "hello":
        fmt.Println(v15, v15)
    default:
        var v41 int16 = 979
        _ = v41
    }
    for fi42 := 0; fi42 < 20; fi42++ {
        v43 := GenS6{F0: 42, F1: v29, F2: 32767}
        for fi44 := 0; fi44 < 5; fi44++ {
            var v45 string = ""
            if false {
                v46 := GenS1{F0: (fi42 | 1000), F1: fi42, F2: (fi44 / ((710) | 1))}
                if false {
                    switch v45 {
                    case "foo bar":
                        switch v45 {
                        case "hello":
                            if false {
                                var v47 uint32 = 2147483647
                                var v48 string = v45
                                switch (fi42 * fi42) {
                                case -929:
                                    var v49 int64 = 42
                                    _ = v49
                                default:
                                    if (0.0 != float32(1.5)) {
                                        v1[v45] = (808 < 212)
                                    }
                                    fmt.Println(v15, fi42)
                                    v47 += (v47 ^ v47)
                                    switch (^(fi44)) {
                                    case 9223372036854775807:
                                        switch (fi42 >> ((42) & 63)) {
                                        case 258:
                                            fmt.Println(fi44, v15, v48)
                                        default:
                                            var v50 uint32 = (v47 & (v47 % ((2147483648) | 1)))
                                            _ = v50
                                        }
                                    case 127:
                                        if (442 != 136) {
                                            v51 := GenS7{F0: v48, F1: genFn0(18446744073709551615), F2: 42}
                                            _ = v51
                                        } else {
                                            v52 := GenS8{F0: (v47 & v47), F1: v45, F2: 0}
                                            if len(v2) > 0 {
                                                v2[0] = 2147483647
                                            }
                                            _ = v52
                                        }
                                    default:
                                        if (-1 < 236) {
                                            delete(v1, "")
                                        } else {
                                            if len(v3) > 0 {
                                                v3[0] = v48
                                            }
                                            var v53 bool = true
                                            v15 += v15
                                            fmt.Println(v53, fi44)
                                            v54 := v1["\n\t"]
                                            _ = v53
                                            _ = v54
                                        }
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                _ = v47
                                _ = v48
                            } else {
                                var v55 string = v45
                                if true {
                                    var v56 uint = 65535
                                    _ = v56
                                }
                                v45 = ""
                                _ = v55
                            }
                        }
                        switch (v45 + v45) {
                        case "foo bar":
                            switch v45 {
                            case "hello":
                                if true {
                                    var v57 int = fi42
                                    fmt.Println(v29)
                                    _ = v57
                                } else {
                                    v1[(v45 + v45)] = (v29 == (v29 - (^(v29))))
                                    var v58 int = (^(fi42))
                                    v59 := GenS1{F0: v58, F1: v58, F2: (v58 | (fi44 / ((-2147483648) | 1)))}
                                    _ = v58
                                    _ = v59
                                }
                            case "\n\t":
                                var v60 uint16 = 921
                                _ = v60
                            default:
                                v61 := GenS8{F0: 2147483648, F1: v45, F2: -36}
                                _ = v61
                            }
                        case "\n\t":
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                        case "hello":
                            var v62 uint = 18446744073709551615
                            v63 := GenS3{F0: (v29 * v29), F1: 2147483648, F2: 2147483647}
                            _ = v62
                            _ = v63
                        default:
                            v64 := GenS8{F0: 874, F1: ("\n\t" + v45), F2: 42}
                            var v65 int16 = 0
                            _ = v64
                            _ = v65
                        }
                        switch v45 {
                        case "hello":
                            v66 := GenS7{F0: v45, F1: 65535, F2: 208}
                            _ = v66
                        case "foo bar":
                            if len(v2) > 0 {
                                v2[0] = (fi42 - 1)
                            }
                        case "a":
                            v15 += (v15 | (-(v15)))
                        }
                    }
                }
                v67 := GenS9{F0: 0, F1: 42}
                _ = v46
                _ = v67
            }
            _ = fi44
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    if (!(true && true)) {
        v2 = append(v2, -610)
        var v68 uint32 = 42
        v29 += genFn0(4294967295)
        _ = v68
    } else {
        v69 := GenS1{F0: 1, F1: 9223372036854775807, F2: 32768}
        if (!false) {
            for fi70 := 0; fi70 < 5; fi70++ {
                _ = len(v1)
                _ = len(v1)
                v29 -= 255
                _ = fi70
            }
        } else {
            var v71 int32 = 0
            switch "a" {
            case "\n\t":
                v72 := GenS10{F0: 42, F1: 65535}
                _ = v72
            default:
                fmt.Println(v29)
            }
            switch "foo bar" {
            case "hello":
                var v73 string = "\n\t"
                _ = v73
            case "":
                for fi74 := 0; fi74 < 9; fi74++ {
                    var v75 uint16 = (^(v29))
                    delete(v1, "foo bar")
                    _ = fi74
                    _ = v75
                }
            default:
                fmt.Println(v71)
            }
            _ = v71
        }
        _ = v69
    }
    _ = len(v3)
    for fi76 := 0; fi76 < 12; fi76++ {
        delete(v1, "")
        switch (^(fi76)) {
        case -285:
            v77 := GenS7{F0: "", F1: (v29 - genFn0(127)), F2: 127}
            _ = v77
        case 389:
            v78 := GenS11{F0: ("" >= "a"), F1: -1, F2: v15}
            _ = v78
        case 385:
            v29 = (v29 / (((v29 * (v29 / ((genFn0(18446744073709551615)) | 1)))) | 1))
        }
        for fi79 := 0; fi79 < 8; fi79++ {
            v80 := GenS12{F0: -6, F1: genFn0(255), F2: 1}
            var v81 uint16 = (v29 * (v29 << ((759) & 63)))
            if (-1 <= 1000) {
                var v82 uint = 42
                _ = v82
            } else {
                v83 := GenS2{F0: true, F1: 94, F2: (fi79 << ((53) & 63))}
                var v84 uint8 = 0
                _ = v83
                _ = v84
            }
            delete(v1, "foo bar")
            v85 := GenS11{F0: (!(false && false)), F1: 332, F2: v15}
            v86 := GenS6{F0: 39, F1: (v81 / ((genFn0(4294967295)) | 1)), F2: (v81 % ((306) | 1))}
            _ = fi79
            _ = v80
            _ = v81
            _ = v85
            _ = v86
        }
        _ = fi76
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v39
}
