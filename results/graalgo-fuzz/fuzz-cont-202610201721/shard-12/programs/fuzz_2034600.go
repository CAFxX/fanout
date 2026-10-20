package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint16; F1 int16 }
type GenS10 struct { F0 string; F1 uint8 }
type GenS11 struct { F0 int8; F1 uint; F2 int16 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 bool }
func genFn0(p0 int) uint8 {
    for fi5 := 0; fi5 < 2; fi5++ {
        v6 := []int{-1, fi5, fi5, (fi5 | p0)}
        fmt.Println(fi5, p0)
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(p0, fi7)
        _ = fi7
    }
    fmt.Println(p0)
    _ = p0
    return 128
}
func genFn1(p0 int16) uint8 {
    if false {
        v8 := []int{0, -2147483648, 9223372036854775807, -2147483648}
        _ = v8
    } else {
        switch 593 {
        case 32767:
            p0 = (-(p0))
        default:
            for fi9 := 0; fi9 < 16; fi9++ {
                if (!true) {
                    p0 -= int16(fi9)
                    p0 = (p0 + (p0 * 1000))
                }
                v10 := GenS1{F0: (("hello" + "") + "")}
                for fi11 := 0; fi11 < 8; fi11++ {
                    switch (fi11 / (((fi9 + fi9)) | 1)) {
                    case 1:
                        var v12 int64 = -1
                        _ = v12
                    default:
                        var v13 int8 = 100
                        _ = v13
                    }
                    var v14 float32 = 1.5
                    if true {
                        fmt.Println(v14, fi9, fi9)
                    }
                    _ = fi11
                    _ = v14
                }
                _ = fi9
                _ = v10
            }
        }
        var v15 int = 32767
        _ = v15
    }
    switch -158 {
    case 9223372036854775807:
        v16 := GenS2{F0: 18446744073709551615}
        _ = v16
    default:
        fmt.Println(p0)
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        fmt.Println(p0, p0)
        var v18 int32 = 928
        p0 += p0
        _ = fi17
        _ = v18
    }
    if (!true) {
        fmt.Println(p0)
        var v19 int8 = 100
        _ = v19
    }
    _ = p0
    return 1
}

func main() {
    v1 := []string{((("hello" + ("" + "foo bar")) + "foo bar") + "foo bar"), "hello", ("" + "hello"), "foo bar"}
    v2 := make(map[uint32]bool)
    var v3 uint8 = 2
    var v4 uint16 = 981
    v3 = (v3 << ((545) & 63))
    var v20 int = 32768
    _ = len(v2)
    v20 -= (v20 >> ((652) & 63))
    v1 = append(v1, "\n\t")
    for fi21 := 0; fi21 < 14; fi21++ {
        v3 = (^(v3))
        _ = fi21
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            delete(v2, 155)
            var v24 uint64 = 42
            _ = fi23
            _ = v24
        }
        var v25 int32 = 2147483647
        _ = fi22
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        v27 := GenS3{F0: 0, F1: v3}
        v28 := GenS1{F0: "foo bar"}
        if ((v3 >> ((242) & 63)) < 255) {
            var v29 int8 = 42
            var v30 uint64 = 0
            _ = v29
            _ = v30
        } else {
            v20 += 759
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    var v31 uint = 65535
    var v32 int8 = -1
    fmt.Println(v31, v32, v31)
    for fi33 := 0; fi33 < 18; fi33++ {
        fmt.Println(v3)
        if false {
            v34 := GenS4{F0: 128}
            _ = v34
        } else {
            if false {
                if true {
                    v2[985] = false
                    var v35 uint64 = 42
                    var v36 uint = 0
                    _ = v35
                    _ = v36
                } else {
                    v37 := GenS5{F0: 4294967295, F1: v3}
                    for fi38 := 0; fi38 < 11; fi38++ {
                        var v39 uint = 65535
                        fmt.Println(v32, v3)
                        v2[2147483648] = false
                        _ = fi38
                        _ = v39
                    }
                    _ = v37
                }
            } else {
                for fi40 := 0; fi40 < 12; fi40++ {
                    fmt.Println(v20, v20)
                    _ = fi40
                }
            }
            v32 -= (v32 + (v32 << ((v31) & 63)))
        }
        switch (("a" + ("a" + (("" + ("foo bar" + ((("foo bar" + "foo bar") + "a") + "a"))) + "\n\t"))) + "") {
        case "hello":
            for fi41 := 0; fi41 < 1; fi41++ {
                if true {
                    v20 = fi33
                    v3 = v3
                    v31 += (v31 & (v31 ^ v31))
                    if (!(9223372036854775807 != -401)) {
                        var v42 float64 = (1.5 + 0.1)
                        if true {
                            if true {
                                delete(v2, 42)
                            } else {
                                switch (v20 / (((fi33 | v20)) | 1)) {
                                case -36:
                                    if ((true && (!(!true))) || false) {
                                        if false {
                                            if false {
                                                var v43 string = "hello"
                                                var v44 int16 = int16(v4)
                                                _ = v43
                                                _ = v44
                                            }
                                        } else {
                                            v4 = (v4 << ((4294967295) & 63))
                                            var v45 uint = (v31 << (((v31 << ((v31) & 63))) & 63))
                                            var v46 string = ((("\n\t" + "\n\t") + "a") + "")
                                            _ = v45
                                            _ = v46
                                        }
                                        switch fi33 {
                                        case -2147483648:
                                            var v47 uint32 = 796
                                            _ = v47
                                        case 4294967295:
                                            if false {
                                                delete(v2, 614)
                                            }
                                            v48 := GenS6{F0: genFn1(1000)}
                                            v1 = append(v1, "\n\t")
                                            _ = v48
                                        default:
                                            v3 -= (^(v3))
                                        }
                                    }
                                }
                                if (87 == -983) {
                                    _ = len(v2)
                                    if (false || (v32 >= 100)) {
                                        v31 -= v31
                                        var v49 uint32 = 438
                                        var v50 int64 = 42
                                        _ = v49
                                        _ = v50
                                    } else {
                                        if false {
                                            var v51 int32 = -2147483648
                                            _ = v51
                                        } else {
                                            v20 = v20
                                        }
                                        fmt.Println(v20, v20)
                                    }
                                } else {
                                    fmt.Println(v3)
                                    if (!true) {
                                        var v52 uint32 = 4294967295
                                        switch ("foo bar" + ((("hello" + "") + "\n\t") + "hello")) {
                                        case "foo bar":
                                            fmt.Println(v52, fi33, v3)
                                        case "hello":
                                            v53 := GenS7{F0: 42, F1: 0}
                                            v54 := GenS8{F0: 42}
                                            v55 := GenS9{F0: 65535, F1: 0}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        case "":
                                            switch "a" {
                                            case "":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case "foo bar":
                                                var v56 string = "\n\t"
                                                _ = v56
                                            case "a":
                                                fmt.Println(fi33, v52, v4)
                                                v32 = (-(v32))
                                                v4 += (v4 >> (((v31 + v31)) & 63))
                                                var v57 int16 = -586
                                                switch -2147483648 {
                                                case 32767:
                                                    var v58 uint16 = (v4 * (v4 * 117))
                                                    var v59 uint8 = genFn0((fi33 << ((v31) & 63)))
                                                    _ = v58
                                                    _ = v59
                                                case 2147483647:
                                                    if ((v52 << ((v31) & 63)) == 429) {
                                                        v3 = 128
                                                    } else {
                                                        _ = len(v2)
                                                        v1 = append(v1, ("a" + "\n\t"))
                                                        var v60 uint8 = 127
                                                        v61 := GenS10{F0: "hello", F1: genFn1(v57)}
                                                        _ = v60
                                                        _ = v61
                                                    }
                                                case 1:
                                                    delete(v2, (v52 / (((v52 + (^(v52)))) | 1)))
                                                }
                                                _ = v57
                                            default:
                                                if (false && false) {
                                                    var v62 uint8 = (^(v3))
                                                    var v63 bool = ("hello" > "hello")
                                                    _ = v62
                                                    _ = v63
                                                } else {
                                                    if true {
                                                        fmt.Println(v3)
                                                        v42 -= v42
                                                        if false {
                                                            v4 = v4
                                                        } else {
                                                            v32 -= v32
                                                            v64 := GenS6{F0: v3}
                                                            var v65 int = (fi33 % (((fi41 & (fi33 | -294))) | 1))
                                                            v4 = 32767
                                                            _ = v64
                                                            _ = v65
                                                        }
                                                    } else {
                                                        v20 = (-(fi33))
                                                    }
                                                    if true {
                                                        var v66 uint64 = 412
                                                        v67 := GenS1{F0: ""}
                                                        _ = v66
                                                        _ = v67
                                                    }
                                                }
                                                v42 -= v42
                                                var v68 uint16 = (^(v4))
                                                _ = v68
                                            }
                                        default:
                                            if true {
                                                var v69 bool = false
                                                v31 = (v31 >> (((v31 | uint(v3))) & 63))
                                                _ = len(v1)
                                                switch fi33 {
                                                case 128:
                                                    var v70 int16 = 42
                                                    _ = v70
                                                case 916:
                                                    var v71 uint8 = genFn0((fi41 & (fi33 % (((v20 >> (((^(v31))) & 63))) | 1))))
                                                    _ = v71
                                                default:
                                                    v72 := GenS11{F0: (v32 & 100), F1: (v31 | (v31 << ((v31) & 63))), F2: -32768}
                                                    _ = v72
                                                }
                                                v31 -= v31
                                                if v69 {
                                                    v42 -= (1e10 - 0.0)
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        var v73 uint16 = (v4 / (((v4 + 2)) | 1))
                                                        _ = v73
                                                    case "foo bar":
                                                        fmt.Println(fi41, v52, fi33)
                                                    case "a":
                                                        v74 := GenS12{F0: (v32 ^ (v32 / (((v32 - v32)) | 1)))}
                                                        _ = v74
                                                    default:
                                                        v32 -= 1
                                                    }
                                                    switch (^(v20)) {
                                                    case -959:
                                                        var v75 uint8 = v3
                                                        _ = v75
                                                    case 0:
                                                        v42 -= 1.5
                                                    case 9223372036854775807:
                                                        _ = len(v1)
                                                    }
                                                    var v76 uint8 = (^(v3))
                                                    _ = v76
                                                } else {
                                                    var v77 string = ""
                                                    v52 -= v52
                                                    var v78 uint32 = v52
                                                    fmt.Println(v31, v4)
                                                    v79 := GenS13{F0: 4294967295}
                                                    var v80 int32 = 2147483647
                                                    v81 := GenS14{F0: v69}
                                                    _ = v77
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                    _ = v81
                                                }
                                                _ = v69
                                            }
                                        }
                                        fmt.Println(v3, v52)
                                        _ = v52
                                    }
                                }
                                v32 += v32
                            }
                        }
                        _ = v42
                    }
                    v2[88] = false
                }
                switch "a" {
                case "hello":
                    v82 := GenS7{F0: 1, F1: 0}
                    _ = v82
                case "\n\t":
                    v3 -= 2
                case "a":
                    var v83 bool = (0 != 2)
                    _ = v83
                }
                _ = fi41
            }
        case "\n\t":
            var v84 int32 = 42
            _ = v84
        case "":
            _ = len(v2)
        }
        _ = fi33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
