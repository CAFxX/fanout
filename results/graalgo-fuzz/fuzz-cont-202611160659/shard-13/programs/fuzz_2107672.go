package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 uint32; F2 bool }
type GenS6 struct { F0 int32; F1 bool; F2 int64 }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint16, p1 uint32, p2 uint64) int64 {
    var v5 int = 661
    p1 = (p1 ^ uint32(p0))
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{int(p1), v5, 32767, v5}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v8 := GenS1{F0: "hello", F1: (("" + "foo bar") + ("hello" + "\n\t"))}
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "\n\t" {
    case "a":
        v9 := GenS1{F0: ("" + ""), F1: (("foo bar" + ("foo bar" + ("" + ("hello" + "a")))) + "")}
        _ = v9
    case "":
        p0 -= p0
    }
    if true {
        var v10 uint64 = (p2 % ((p2) | 1))
        var v11 int8 = 0
        _ = v10
        _ = v11
    } else {
        v12 := []int{-2147483648, v5}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 string) int {
    v13 := []int{}
    for fi14 := 0; fi14 < 17; fi14++ {
        if true {
            if (1000 >= 1) {
                var v15 int32 = 2147483647
                p0 = p0
                _ = v15
            }
            v16 := GenS2{F0: 649, F1: 0, F2: 555}
            p0 = p0
            _ = v16
        }
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 uint64 = 42
            fmt.Println(p0, v18)
            _ = fi17
            _ = v18
        }
        var v19 float32 = 1.5
        _ = fi14
        _ = v19
    }
    var v20 uint64 = 18446744073709551615
    _ = p0
    _ = v13
    _ = v20
    return -109
}

func main() {
    v1 := []int{9223372036854775807, 1, 0}
    var v2 uint64 = 2
    v3 := []int{}
    v4 := make(map[uint32]bool)
    fmt.Println(v2)
    fmt.Println(v2)
    for fi21 := 0; fi21 < 6; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            _ = len(v3)
            _ = fi22
        }
        _ = fi21
    }
    v2 = (v2 / (((^(v2))) | 1))
    if (!(4294967295 == 2)) {
        v2 += v2
        if false {
            switch "" {
            case "\n\t":
                var v23 uint64 = 852
                _ = v23
            case "foo bar":
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 uint64 = (v2 - v2)
                    _ = fi24
                    _ = v25
                }
            default:
                for fi26 := 0; fi26 < 20; fi26++ {
                    v27 := GenS2{F0: 634, F1: 381, F2: 1}
                    for fi28 := 0; fi28 < 1; fi28++ {
                        switch ("\n\t" + "a") {
                        case "hello":
                            fmt.Println(fi26, fi26, v2)
                        case "a":
                            var v29 int = fi28
                            _ = v29
                        }
                        _ = fi28
                    }
                    for fi30 := 0; fi30 < 13; fi30++ {
                        v2 += (v2 / ((262) | 1))
                        v2 -= v2
                        _ = fi30
                    }
                    _ = fi26
                    _ = v27
                }
            }
        } else {
            v2 = (v2 * (v2 % ((v2) | 1)))
            var v31 uint32 = 2
            switch -9223372036854775808 {
            case 42:
                if false {
                    v32 := v4[v31]
                    for fi33 := 0; fi33 < 3; fi33++ {
                        v34 := GenS3{F0: 150, F1: true, F2: 48}
                        _ = fi33
                        _ = v34
                    }
                    _ = v32
                }
            case 32767:
                var v35 int = genFn1((("\n\t" + (("foo bar" + "a") + "")) + ("" + "")))
                _ = v35
            case 796:
                if true {
                    if false {
                        if ((v31 >> ((1) & 63)) != v31) {
                            var v36 float64 = -2.25
                            fmt.Println(v31, v31, v36)
                            v2 = v2
                            _ = v36
                        } else {
                            fmt.Println(v2)
                            if (4294967295 > 42) {
                                if (v31 <= (^(v31))) {
                                    v31 -= 2147483647
                                    for fi37 := 0; fi37 < 9; fi37++ {
                                        var v38 uint16 = 0
                                        switch fi37 {
                                        case 127:
                                            v39 := GenS3{F0: 683, F1: (fi37 != fi37), F2: 1}
                                            _ = v39
                                        case 32767:
                                            v38 += uint16(v31)
                                        case 2147483647:
                                            fmt.Println(fi37, v31, v38)
                                        default:
                                            var v40 uint16 = (v38 % (((^(v38))) | 1))
                                            _ = v40
                                        }
                                        _ = fi37
                                        _ = v38
                                    }
                                    v41 := GenS4{F0: 992}
                                    _ = v41
                                }
                                for fi42 := 0; fi42 < 14; fi42++ {
                                    var v43 uint64 = v2
                                    switch 32768 {
                                    case 0:
                                        var v44 int16 = 0
                                        _ = v44
                                    case -969:
                                        delete(v4, v31)
                                    case -1:
                                        var v45 int16 = 663
                                        _ = v45
                                    default:
                                        var v46 int64 = 0
                                        _ = v46
                                    }
                                    for fi47 := 0; fi47 < 1; fi47++ {
                                        if (true || (true && true)) {
                                            v48 := GenS5{F0: false, F1: v31, F2: true}
                                            if false {
                                                v2 += (v2 % (((^(v2))) | 1))
                                                v49 := GenS1{F0: "foo bar", F1: "a"}
                                                var v50 uint64 = 648
                                                _ = v49
                                                _ = v50
                                            }
                                            v51 := GenS1{F0: "hello", F1: (("" + "a") + "")}
                                            v31 += 982
                                            _ = v48
                                            _ = v51
                                        }
                                        var v52 uint = 255
                                        _ = fi47
                                        _ = v52
                                    }
                                    _ = fi42
                                    _ = v43
                                }
                                for fi53 := 0; fi53 < 3; fi53++ {
                                    var v54 uint32 = v31
                                    if true {
                                        for fi55 := 0; fi55 < 13; fi55++ {
                                            v56 := GenS2{F0: 1, F1: 2, F2: v54}
                                            _ = len(v3)
                                            _ = fi55
                                            _ = v56
                                        }
                                    } else {
                                        delete(v4, v54)
                                    }
                                    _ = fi53
                                    _ = v54
                                }
                            } else {
                                var v57 int32 = 42
                                v57 += v57
                                var v58 float64 = 1e10
                                _ = v57
                                _ = v58
                            }
                        }
                    } else {
                        var v59 int = 1
                        v59 = genFn1("hello")
                        v60 := v4[v31]
                        _ = v59
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 5; fi61++ {
                        for fi62 := 0; fi62 < 19; fi62++ {
                            switch "hello" {
                            case "":
                                switch "hello" {
                                case "\n\t":
                                    v31 += v31
                                case "foo bar":
                                    var v63 uint16 = 0
                                    _ = v63
                                case "hello":
                                    if len(v1) > 0 {
                                        v1[0] = (fi62 ^ (fi62 | (fi61 >> ((65535) & 63))))
                                    }
                                default:
                                    var v64 int = (fi61 / (((fi61 & fi61)) | 1))
                                    switch "a" {
                                    case "":
                                        v31 = v31
                                    }
                                    _ = v64
                                }
                            case "a":
                                v65 := GenS6{F0: 208, F1: false, F2: genFn0(42, 447, (v2 | v2))}
                                _ = v65
                            default:
                                fmt.Println(fi61, fi61, fi62)
                            }
                            var v66 uint64 = v2
                            v66 -= (v2 + (v66 << ((1000) & 63)))
                            _ = fi62
                            _ = v66
                        }
                        v2 -= (v2 | v2)
                        fmt.Println(v2)
                        _ = fi61
                    }
                }
            }
            _ = v31
        }
    }
    if (((!(!true)) && ("" != "\n\t")) && false) {
        v67 := GenS7{F0: 210, F1: 32767}
        _ = v67
    }
    var v68 string = "hello"
    if true {
        _ = len(v1)
        var v69 float64 = 1.5
        v68 = "hello"
        _ = v69
    } else {
        if (!false) {
            if (42 == 1) {
                var v70 uint = 4294967295
                _ = v70
            } else {
                v4[42] = false
            }
            if (!false) {
                fmt.Println(v68, v2)
                for fi71 := 0; fi71 < 11; fi71++ {
                    v72 := GenS6{F0: 1, F1: (false && false), F2: 9223372036854775807}
                    _ = fi71
                    _ = v72
                }
            }
        } else {
            for fi73 := 0; fi73 < 7; fi73++ {
                v74 := GenS8{F0: 12, F1: 1}
                if true {
                    if true {
                        v75 := GenS8{F0: 222, F1: 881}
                        _ = v75
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v68 = ("\n\t" + ("hello" + v68))
                        fmt.Println(fi73, v2)
                    }
                }
                _ = fi73
                _ = v74
            }
        }
        var v76 float64 = 1e10
        for fi77 := 0; fi77 < 14; fi77++ {
            var v78 uint = 2
            for fi79 := 0; fi79 < 7; fi79++ {
                v80 := GenS4{F0: 691}
                _ = fi79
                _ = v80
            }
            var v81 bool = false
            _ = fi77
            _ = v78
            _ = v81
        }
        _ = v76
    }
    v82 := GenS4{F0: 2}
    if true {
        if len(v1) > 0 {
            v1[0] = -185
        }
        var v83 int32 = 0
        _ = v83
    }
    fmt.Println(v2)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v68
    _ = v82
}
