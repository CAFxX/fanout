package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS3 struct { F0 string; F1 int64; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint16; F1 uint8; F2 string }
type GenS6 struct { F0 int16; F1 bool }
type GenS7 struct { F0 uint; F1 int; F2 string }
func genFn0() int16 {
    for fi4 := 0; fi4 < 10; fi4++ {
        v5 := GenS1{F0: 100, F1: ""}
        v6 := GenS1{F0: 1, F1: "\n\t"}
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 float64 = 0.1
    _ = v7
    return 1000
}
func genFn1() uint16 {
    var v8 int = -2147483648
    if true {
        if (1 >= 255) {
            v9 := []int{v8, v8}
            var v10 int32 = -356
            _ = v10
            _ = v9
        }
        if ((!(0.0 > (-2.25 * 3.14159))) && true) {
            var v11 uint32 = 2147483647
            _ = v11
        }
        var v12 float32 = (1.5 * ((3.14159 / 1e10) * 3.14159))
        _ = v12
    }
    _ = v8
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int = 9223372036854775807
    var v3 uint64 = 1
    if (!true) {
        if false {
            v13 := GenS2{F0: (v3 | (v3 % ((2) | 1))), F1: 2, F2: v1}
            _ = v13
        } else {
            if ((!false) && false) {
                fmt.Println(v1)
                if (!(0.0 > (1e10 * 1.5))) {
                    v14 := []int{v2, (v2 - v2)}
                    fmt.Println(v2, v1)
                    _ = v14
                }
            } else {
                switch "hello" {
                case "hello":
                    v1 += v1
                case "a":
                    for fi15 := 0; fi15 < 19; fi15++ {
                        fmt.Println(v3)
                        v16 := []int{}
                        v1 += (^(v1))
                        _ = fi15
                        _ = v16
                    }
                case "":
                    v17 := []int{127, 32768, (v2 - v2)}
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 17; fi18++ {
                        v1 = -9223372036854775808
                        if false {
                            v3 -= v3
                            for fi19 := 0; fi19 < 1; fi19++ {
                                var v20 float32 = 0.0
                                v20 -= 1.5
                                _ = fi19
                                _ = v20
                            }
                        } else {
                            for fi21 := 0; fi21 < 15; fi21++ {
                                var v22 int32 = 0
                                if ((v3 & v3) != (^(v3))) {
                                    v2 = (^(fi18))
                                    if true {
                                        v23 := GenS3{F0: "foo bar", F1: -400, F2: 0}
                                        v2 = fi18
                                        fmt.Println(fi18, v1, v2)
                                        _ = v23
                                    }
                                } else {
                                    v3 = v3
                                    fmt.Println(v22)
                                    var v24 int16 = genFn0()
                                    _ = v24
                                }
                                _ = fi21
                                _ = v22
                            }
                            switch "a" {
                            case "hello":
                                for fi25 := 0; fi25 < 10; fi25++ {
                                    if true {
                                        v2 += 245
                                    }
                                    switch (-(fi25)) {
                                    case 2147483647:
                                        v26 := GenS2{F0: (v3 + 673), F1: 4294967295, F2: v1}
                                        _ = v26
                                    default:
                                        v1 = (v1 * v1)
                                    }
                                    _ = fi25
                                }
                            case "\n\t":
                                for fi27 := 0; fi27 < 12; fi27++ {
                                    var v28 uint64 = v3
                                    _ = fi27
                                    _ = v28
                                }
                            default:
                                if (!((!true) || false)) {
                                    switch 1 {
                                    case -17:
                                        v1 = 0
                                    case 1:
                                        var v29 int32 = -1
                                        _ = v29
                                    }
                                    switch (fi18 | (-(fi18))) {
                                    case -989:
                                        v30 := GenS1{F0: -9, F1: "a"}
                                        _ = v30
                                    default:
                                        switch "hello" {
                                        case "":
                                            for fi31 := 0; fi31 < 15; fi31++ {
                                                v32 := GenS2{F0: (v3 | v3), F1: 368, F2: 162}
                                                switch fi18 {
                                                case 32768:
                                                    v33 := GenS1{F0: 1, F1: "a"}
                                                    _ = v33
                                                default:
                                                    v34 := GenS2{F0: 42, F1: 991, F2: (v1 % ((-1) | 1))}
                                                    _ = v34
                                                }
                                                v35 := GenS4{F0: -1, F1: 2147483647, F2: 420}
                                                _ = fi31
                                                _ = v32
                                                _ = v35
                                            }
                                        case "a":
                                            switch 127 {
                                            case -737:
                                                switch (("" + "\n\t") + "hello") {
                                                case "":
                                                    if ((v3 == (v3 | 0)) && true) {
                                                        v1 += v1
                                                        fmt.Println(fi18, fi18, v3)
                                                        switch "hello" {
                                                        case "a":
                                                            switch "hello" {
                                                            case "a":
                                                                v36 := []int{(fi18 | (v2 << ((255) & 63))), (fi18 >> ((255) & 63)), fi18, -9223372036854775808}
                                                                _ = v36
                                                            }
                                                        case "":
                                                            var v37 uint32 = 428
                                                            _ = v37
                                                        default:
                                                            if false {
                                                                v38 := []int{}
                                                                _ = v38
                                                            } else {
                                                                v39 := []int{(fi18 & v2), v2, -40, (fi18 ^ (v2 ^ v2))}
                                                                for fi40 := 0; fi40 < 19; fi40++ {
                                                                    v3 = 463
                                                                    if true {
                                                                        v3 = (v3 >> ((65535) & 63))
                                                                        v1 = (-(v1))
                                                                        v39 = append(v39, 32767)
                                                                    } else {
                                                                        v2 -= 42
                                                                        v41 := GenS5{F0: 160, F1: 1, F2: ("" + ("hello" + "foo bar"))}
                                                                        fmt.Println(fi40, fi18)
                                                                        v39 = append(v39, fi18)
                                                                        _ = v41
                                                                    }
                                                                    if len(v39) >= 2 {
                                                                        v39 = v39[1:len(v39)]
                                                                    }
                                                                    _ = fi40
                                                                }
                                                                _ = v39
                                                            }
                                                        }
                                                    } else {
                                                        for fi42 := 0; fi42 < 11; fi42++ {
                                                            v2 = fi18
                                                            switch "\n\t" {
                                                            case "":
                                                                v43 := GenS5{F0: 32767, F1: 42, F2: "\n\t"}
                                                                _ = v43
                                                            }
                                                            var v44 float64 = 0.0
                                                            _ = fi42
                                                            _ = v44
                                                        }
                                                        v1 = v1
                                                        var v45 int64 = -696
                                                        _ = v45
                                                    }
                                                }
                                            case 32767:
                                                var v46 uint = 1
                                                _ = v46
                                            case 127:
                                                v2 = (v2 % ((fi18) | 1))
                                            }
                                        }
                                    }
                                    for fi47 := 0; fi47 < 6; fi47++ {
                                        v48 := GenS6{F0: -118, F1: (2147483647 != -2147483648)}
                                        _ = fi47
                                        _ = v48
                                    }
                                } else {
                                    if false {
                                        var v49 float64 = 0.1
                                        fmt.Println(v1, v49)
                                        _ = v49
                                    }
                                }
                            }
                            var v50 int64 = 715
                            _ = v50
                        }
                        v51 := GenS4{F0: genFn0(), F1: 4294967295, F2: genFn0()}
                        _ = fi18
                        _ = v51
                    }
                }
            }
            fmt.Println(v1, v2, v2)
            switch "" {
            case "hello":
                var v52 float32 = -2.25
                _ = v52
            }
        }
        var v53 float32 = 1.5
        _ = v53
    }
    for fi54 := 0; fi54 < 1; fi54++ {
        switch "hello" {
        case "a":
            v3 += (^(v3))
        case "hello":
            v55 := GenS1{F0: 0, F1: "hello"}
            _ = v55
        default:
            v3 += v3
        }
        v56 := []int{v2}
        _ = fi54
        _ = v56
    }
    for fi57 := 0; fi57 < 3; fi57++ {
        for fi58 := 0; fi58 < 3; fi58++ {
            fmt.Println(fi57, fi57)
            _ = fi58
        }
        _ = fi57
    }
    v59 := GenS6{F0: 1, F1: (734 > 2147483648)}
    if false {
        if false {
            v2 = v2
            v60 := []int{}
            _ = len(v60)
            _ = v60
        }
    }
    v61 := GenS7{F0: 42, F1: v2, F2: (("a" + "") + "")}
    if false {
        v62 := []int{(v2 % ((v2) | 1))}
        if len(v62) >= 2 {
            v62 = v62[1:len(v62)]
        }
        fmt.Println(v1, v3)
        _ = v62
    }
    fmt.Println(v2, v1, v3)
    var v63 int16 = -382
    var v64 uint16 = 1
    switch v2 {
    case 1:
        fmt.Println(v64, v2, v63)
    case -159:
        var v65 string = "foo bar"
        _ = v65
    case 25:
        v66 := GenS5{F0: (v64 | (v64 ^ (v64 << ((65535) & 63)))), F1: 127, F2: "a"}
        _ = v66
    default:
        switch (v2 + v2) {
        case 440:
            v67 := GenS2{F0: v3, F1: 2147483648, F2: 9223372036854775807}
            _ = v67
        default:
            var v68 int32 = 0
            _ = v68
        }
    }
    fmt.Println(v3, v1, v63)
    v69 := []int{}
    v63 += -869
    for fi70 := 0; fi70 < 12; fi70++ {
        switch "a" {
        case "\n\t":
            v2 -= 9223372036854775807
        case "foo bar":
            if false {
                fmt.Println(v3, v64, v63)
                for fi71 := 0; fi71 < 20; fi71++ {
                    if len(v69) > 0 {
                        v69[0] = 853
                    }
                    if ((v2 > (-(fi70))) && true) {
                        v1 = v1
                        v2 = fi70
                        fmt.Println(v2, v63)
                        var v72 int64 = (v1 * 9223372036854775807)
                        _ = v72
                    }
                    _ = fi71
                }
            } else {
                fmt.Println(fi70, fi70)
            }
        case "hello":
            v3 -= 18446744073709551615
        default:
            var v73 uint32 = 65535
            _ = v73
        }
        _ = fi70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v69:", len(v69), v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v59
    _ = v61
    _ = v63
    _ = v64
    _ = v69
}
