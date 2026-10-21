package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 string; F1 int16; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 int16; F2 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 uint32; F1 int32; F2 int64 }
func genFn0(p0 int) int16 {
    if (!true) {
        v4 := GenS1{F0: 128, F1: 888}
        _ = v4
    } else {
        if (556 < 65535) {
            p0 += 128
            p0 = (p0 & (p0 * p0))
            v5 := GenS2{F0: 2147483648, F1: 2, F2: 128}
            _ = v5
        }
    }
    p0 = -805
    if (!false) {
        var v6 float32 = (((0.1 / 0.1) / -2.25) / 1.0)
        _ = v6
    }
    _ = p0
    return 32767
}
func genFn1(p0 int, p1 uint64) uint {
    var v7 int8 = 49
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            if false {
                switch (p0 - fi8) {
                case 1:
                    var v10 uint64 = (p1 & p1)
                    _ = v10
                case 4294967295:
                    switch (p0 & fi9) {
                    case 128:
                        p0 += (fi9 ^ (fi8 * 0))
                    default:
                        v7 = (-(v7))
                    }
                default:
                    p1 = (p1 - p1)
                }
                v11 := []int{(fi9 & fi9), fi9}
                fmt.Println(p0, fi9)
                fmt.Println(fi9, v7, v7)
                _ = v11
            } else {
                switch (-(p0)) {
                case -805:
                    var v12 string = ""
                    _ = v12
                case 9223372036854775807:
                    if false {
                        switch "a" {
                        case "foo bar":
                            v13 := []int{int(p1), p0, -9223372036854775808}
                            _ = v13
                        case "a":
                            var v14 uint16 = 597
                            _ = v14
                        default:
                            v15 := GenS3{F0: true, F1: 1}
                            _ = v15
                        }
                        v16 := []int{p0, p0}
                        var v17 uint32 = 581
                        _ = v16
                        _ = v17
                    }
                default:
                    p1 -= p1
                }
                v18 := GenS1{F0: 244, F1: 1}
                fmt.Println(p0)
                p0 -= (fi8 & fi9)
                _ = v18
            }
            v19 := []int{fi8, fi8}
            _ = fi9
            _ = v19
        }
        var v20 int = p0
        _ = fi8
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint16 = 32767
    var v3 float64 = 3.14159
    switch 1 {
    case 32767:
        var v21 bool = false
        _ = v21
    case -1:
        v22 := []int{-157, 2147483647}
        _ = v22
    default:
        delete(v1, 906)
    }
    v23 := []int{-415, 128}
    delete(v1, 2)
    var v24 uint = 18446744073709551615
    for fi25 := 0; fi25 < 7; fi25++ {
        if false {
            if (!(1.5 >= 0.0)) {
                fmt.Println(v2)
            }
            for fi26 := 0; fi26 < 19; fi26++ {
                v1[42] = "\n\t"
                v1[621] = ("hello" + "\n\t")
                v2 += v2
                _ = fi26
            }
            switch fi25 {
            case 9223372036854775807:
                var v27 int64 = 317
                _ = v27
            case 0:
                if (0 == 0) {
                    for fi28 := 0; fi28 < 20; fi28++ {
                        var v29 int = (fi25 << ((188) & 63))
                        if (1 != 0) {
                            fmt.Println(v3, fi25, v3)
                        }
                        _ = fi28
                        _ = v29
                    }
                    v30 := v1[681]
                    _ = v30
                } else {
                    var v31 float32 = (((((0.0 + 3.14159) * -2.25) * 1e10) * 0.1) + -2.25)
                    var v32 int64 = -9223372036854775808
                    _ = v31
                    _ = v32
                }
            case 2147483647:
                if false {
                    var v33 bool = (false || (!(!false)))
                    if ((195 != 97) || ((v33 || (((uint32(v2) == 65535) && v33) && v33)) && false)) {
                        v33 = v33
                    } else {
                        v3 -= (v3 * float64(-2.25))
                        v24 -= (v24 % ((genFn1(fi25, 636)) | 1))
                        if v33 {
                            v34 := GenS1{F0: 149, F1: (v2 << ((v24) & 63))}
                            fmt.Println(v33, v2)
                            _ = v34
                        } else {
                            switch (fi25 % ((fi25) | 1)) {
                            case 75:
                                var v35 uint32 = 829
                                _ = v35
                            default:
                                for fi36 := 0; fi36 < 4; fi36++ {
                                    v23 = append(v23, (fi25 >> ((genFn1(-979, 280)) & 63)))
                                    var v37 uint64 = 4294967295
                                    fmt.Println(v3, v33)
                                    _ = fi36
                                    _ = v37
                                }
                            }
                        }
                    }
                    _ = v33
                }
            default:
                switch (^(fi25)) {
                case -901:
                    delete(v1, 42)
                case -1:
                    v1[uint32(v24)] = (("\n\t" + "hello") + "a")
                default:
                    v23 = append(v23, fi25)
                }
            }
        }
        if (0.1 <= 0.1) {
            _ = len(v1)
            var v38 bool = (253 >= 0)
            _ = len(v23)
            _ = v38
        }
        for fi39 := 0; fi39 < 18; fi39++ {
            if true {
                if true {
                    if (42 == 42) {
                        if false {
                            v40 := GenS4{F0: ("\n\t" + "a"), F1: -32768, F2: fi39}
                            var v41 string = "a"
                            switch v41 {
                            case "\n\t":
                                v24 = v24
                            case "a":
                                fmt.Println(v2, fi25)
                            case "hello":
                                if true {
                                    var v42 float32 = 0.1
                                    _ = v42
                                } else {
                                    var v43 float32 = (((0.1 * (1e10 - ((0.0 / 1e10) * 1.5))) * 0.0) - 0.1)
                                    v41 = v41
                                    switch v41 {
                                    case "\n\t":
                                        fmt.Println(v3, v43)
                                    case "a":
                                        fmt.Println(v24, fi25)
                                    }
                                    _ = v43
                                }
                            }
                            _ = v40
                            _ = v41
                        } else {
                            var v44 uint = (^(v24))
                            if (-1 <= 28) {
                                switch "" {
                                case "\n\t":
                                    fmt.Println(fi25, v44, v3)
                                case "foo bar":
                                    fmt.Println(v2, fi25)
                                case "":
                                    var v45 bool = true
                                    _ = v45
                                }
                                var v46 string = ("\n\t" + "foo bar")
                                v47 := GenS1{F0: 250, F1: v2}
                                v48 := GenS5{F0: genFn0((fi25 / (((fi25 * fi39)) | 1)))}
                                v49 := GenS4{F0: v46, F1: 0, F2: -849}
                                _ = v46
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            var v50 int16 = genFn0((fi39 >> ((genFn1(fi25, 42)) & 63)))
                            v44 = (v24 - genFn1(792, 4294967295))
                            v2 += v2
                            switch (fi39 + (fi25 % ((fi39) | 1))) {
                            case -534:
                                v44 = (v44 | genFn1(fi39, 42))
                            case 32767:
                                var v51 string = "foo bar"
                                _ = v51
                            case 531:
                                if (151 < 18446744073709551615) {
                                    if false {
                                        fmt.Println(fi39)
                                        if ((0 >= 587) && (v2 != (v2 & (v2 + v2)))) {
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                            _ = len(v1)
                                            var v52 uint8 = 255
                                            v53 := GenS6{F0: 2147483647, F1: (-(v50)), F2: (v24 % ((2) | 1))}
                                            _ = v52
                                            _ = v53
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                switch fi25 {
                                                case -2147483648:
                                                    v2 = v2
                                                case 0:
                                                    switch "" {
                                                    case "\n\t":
                                                        if false {
                                                            var v54 int = (^(fi39))
                                                            var v55 int32 = -1
                                                            _ = v54
                                                            _ = v55
                                                        }
                                                    case "hello":
                                                        var v56 uint16 = (v2 ^ v2)
                                                        _ = v56
                                                    case "":
                                                        fmt.Println(v2, v44)
                                                    default:
                                                        var v57 float64 = v3
                                                        _ = v57
                                                    }
                                                case -1:
                                                    v58 := GenS1{F0: 133, F1: v2}
                                                    _ = v58
                                                default:
                                                    if len(v23) >= 2 {
                                                        v23 = v23[1:len(v23)]
                                                    }
                                                    if false {
                                                        switch "hello" {
                                                        case "hello":
                                                            var v59 int16 = -1
                                                            _ = v59
                                                        case "a":
                                                            if ((-9223372036854775808 <= 837) && (0.0 >= float32((3.14159 + ((v3 * v3) / v3))))) {
                                                                switch (fi39 << (((v24 - (v24 * v44))) & 63)) {
                                                                case -849:
                                                                    var v60 float32 = 1.5
                                                                    _ = v60
                                                                case -2147483648:
                                                                    var v61 float64 = v3
                                                                    _ = v61
                                                                case -231:
                                                                    v24 += (v44 ^ genFn1(127, 771))
                                                                }
                                                            } else {
                                                                var v62 uint8 = uint8(v44)
                                                                fmt.Println(v24, fi39)
                                                                _ = v62
                                                            }
                                                        case "foo bar":
                                                            fmt.Println(fi25, v2)
                                                        default:
                                                            switch (fi39 >> (((v24 ^ (v44 - (v24 ^ 18446744073709551615)))) & 63)) {
                                                            case 16:
                                                                v2 = v2
                                                            case 42:
                                                                v2 -= 42
                                                            }
                                                        }
                                                        v3 -= (0.1 * (float64(3.14159) * 0.0))
                                                        var v63 bool = (!(-77 != -83))
                                                        _ = v63
                                                    }
                                                }
                                            default:
                                                var v64 uint32 = 1
                                                _ = v64
                                            }
                                            switch (-(fi39)) {
                                            case 32767:
                                                if ((^(v2)) >= (v2 | v2)) {
                                                    _ = len(v23)
                                                }
                                            }
                                        }
                                    }
                                    var v65 int = fi25
                                    _ = v65
                                }
                            }
                            var v66 bool = ((!(false && (!(genFn1(32768, 0) > (v44 | (v44 / ((v44) | 1))))))) && false)
                            switch -1 {
                            case 0:
                                v66 = v66
                            case 45:
                                switch (fi25 + fi25) {
                                case 42:
                                    var v67 uint32 = 1
                                    _ = v67
                                case -977:
                                    v3 -= (1e10 + v3)
                                case 962:
                                    if v66 {
                                        var v68 uint8 = 1
                                        if v66 {
                                            v69 := v1[2147483648]
                                            _ = v69
                                        } else {
                                            if v66 {
                                                v70 := GenS7{F0: 0}
                                                _ = v70
                                            } else {
                                                var v71 int64 = 9223372036854775807
                                                _ = v71
                                            }
                                        }
                                        v72 := GenS5{F0: genFn0(32767)}
                                        fmt.Println(fi39, v68, fi39)
                                        _ = v68
                                        _ = v72
                                    }
                                }
                            }
                            v73 := GenS8{F0: "foo bar", F1: 0}
                            v44 += (^(v24))
                            _ = v44
                            _ = v50
                            _ = v66
                            _ = v73
                        }
                        switch "" {
                        case "foo bar":
                            v3 += (((((v3 + (v3 - 0.1)) - (3.14159 * v3)) - 1e10) - (v3 / -2.25)) + v3)
                        case "a":
                            v24 = (v24 ^ 4294967295)
                        }
                        var v74 uint32 = 42
                        v74 += (v74 % (((v74 & v74)) | 1))
                        _ = v74
                    }
                    v24 = v24
                    switch "hello" {
                    case "hello":
                        fmt.Println(v3)
                    default:
                        v23 = append(v23, fi39)
                    }
                    v3 -= v3
                } else {
                    fmt.Println(v24)
                }
            }
            _ = len(v23)
            _ = fi39
        }
        _ = fi25
    }
    var v75 int32 = 1
    switch -53 {
    case 67:
        v24 -= v24
    }
    var v76 bool = false
    v77 := GenS7{F0: (^(v75))}
    v24 -= (^(v24))
    v24 = v24
    if v76 {
        for fi78 := 0; fi78 < 4; fi78++ {
            for fi79 := 0; fi79 < 3; fi79++ {
                v3 += v3
                _ = fi79
            }
            _ = fi78
        }
    } else {
        var v80 int64 = -1
        _ = v80
    }
    var v81 uint32 = uint32(v2)
    v82 := GenS9{F0: v81, F1: -2147483648, F2: -9223372036854775808}
    switch 0 {
    case 9223372036854775807:
        switch 9223372036854775807 {
        case 127:
            _ = len(v1)
        }
    case -243:
        fmt.Println(v2, v75)
    case 0:
        v76 = ((v75 + (^(v75))) != v75)
    default:
        var v83 string = "hello"
        _ = v83
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v75
    _ = v76
    _ = v77
    _ = v81
    _ = v82
}
