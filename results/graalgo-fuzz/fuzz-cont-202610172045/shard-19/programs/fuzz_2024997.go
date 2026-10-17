package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 int; F1 uint8; F2 string }
type GenS5 struct { F0 bool; F1 uint16; F2 int }
func genFn0() uint32 {
    v4 := GenS1{F0: false, F1: 964, F2: 1}
    for fi5 := 0; fi5 < 8; fi5++ {
        switch (fi5 % (((fi5 ^ 128)) | 1)) {
        case 32767:
            var v6 int64 = 1
            _ = v6
        case 128:
            var v7 float32 = 0.1
            _ = v7
        case 413:
            if true {
                var v8 bool = true
                _ = v8
            } else {
                var v9 uint = 18446744073709551615
                _ = v9
            }
        default:
            var v10 uint8 = 1
            _ = v10
        }
        _ = fi5
    }
    switch "\n\t" {
    case "":
        var v11 int = 4294967295
        _ = v11
    default:
        v12 := []int{-2147483648, 42}
        _ = v12
    }
    var v13 uint64 = 42
    _ = v13
    _ = v4
    return 63
}
func genFn1(p0 int64, p1 bool) uint16 {
    var v14 uint = 974
    if false {
        p0 += p0
    }
    fmt.Println(v14, p0)
    v14 = 65535
    _ = p0
    _ = p1
    _ = v14
    return 178
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 2147483647
    var v3 int8 = -36
    for fi15 := 0; fi15 < 13; fi15++ {
        v16 := GenS1{F0: true, F1: 255, F2: 0}
        switch v1 {
        case "":
            var v17 uint64 = 18446744073709551615
            _ = v17
        case "\n\t":
            v18 := []int{}
            _ = v18
        default:
            for fi19 := 0; fi19 < 3; fi19++ {
                if (18446744073709551615 >= 2) {
                    v20 := GenS1{F0: (true && (!true)), F1: 0, F2: genFn1(-333, (1 < 94))}
                    switch (-(fi19)) {
                    case 9223372036854775807:
                        if false {
                            switch (-(fi19)) {
                            case 42:
                                fmt.Println(v3, v3, fi15)
                            default:
                                if true {
                                    switch "\n\t" {
                                    case "foo bar":
                                        v3 = (v3 % ((v3) | 1))
                                        var v21 int8 = (v3 % ((v3) | 1))
                                        v22 := []int{0, (^(fi15)), v2, fi15}
                                        var v23 float64 = (0.0 * float64((1.5 + 1e10)))
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    case "\n\t":
                                        var v24 int64 = -307
                                        _ = v24
                                    default:
                                        v2 += fi15
                                    }
                                    v2 -= fi19
                                    var v25 uint = 0
                                    _ = v25
                                } else {
                                    var v26 float64 = (0.0 * 3.14159)
                                    v3 = (-(v3))
                                    _ = v26
                                }
                            }
                            v1 = ("foo bar" + "")
                            var v27 uint32 = genFn0()
                            _ = v27
                        }
                    default:
                        fmt.Println(v2, v1)
                    }
                    if (-1 > -2147483648) {
                        fmt.Println(v2, fi15, v2)
                        if true {
                            if (-583 == 1) {
                                if true {
                                    v28 := GenS1{F0: true, F1: 32767, F2: 479}
                                    _ = v28
                                }
                            }
                            v29 := []int{}
                            if (381 <= 678) {
                                var v30 uint64 = 0
                                var v31 bool = true
                                if false {
                                    var v32 uint64 = (v30 + v30)
                                    _ = len(v29)
                                    _ = v32
                                }
                                _ = v30
                                _ = v31
                            }
                            _ = v29
                        }
                        v33 := []int{v2, v2, v2}
                        _ = v33
                    }
                    v3 -= v3
                    if false {
                        v34 := []int{fi15, fi15, (v2 | fi19), 0}
                        _ = v34
                    } else {
                        v35 := GenS1{F0: true, F1: genFn1(150, (true && (!false))), F2: 1}
                        v2 = 128
                        fmt.Println(v2, fi15)
                        fmt.Println(v2, fi15)
                        if (255 > 342) {
                            fmt.Println(fi19, fi15)
                            if (true && (false && true)) {
                                if (-812 > 42) {
                                    if false {
                                        v36 := []int{(v2 & (fi15 >> ((902) & 63))), (-(fi19)), fi15, -546}
                                        if (65535 >= genFn0()) {
                                            var v37 uint64 = 18446744073709551615
                                            fmt.Println(v2, fi15, fi19)
                                            _ = v37
                                        } else {
                                            v38 := GenS1{F0: false, F1: 32767, F2: 65535}
                                            v2 += fi19
                                            v39 := GenS1{F0: (false || (989 >= 4294967295)), F1: 42, F2: 185}
                                            _ = v38
                                            _ = v39
                                        }
                                        var v40 string = ("\n\t" + ((v1 + (v1 + ("foo bar" + "\n\t"))) + (("hello" + "a") + v1)))
                                        _ = v36
                                        _ = v40
                                    }
                                    if false {
                                        v41 := []int{-1, -9223372036854775808, v2}
                                        v2 = fi19
                                        var v42 uint32 = 4294967295
                                        _ = v41
                                        _ = v42
                                    }
                                }
                                v43 := []int{fi15, (fi15 * 2147483647), 0}
                                v3 -= v3
                                _ = v43
                            } else {
                                if ((((!((v1 + "") != ("a" + v1))) || (!false)) || (42 == 9223372036854775807)) && (!(0 <= 0))) {
                                    var v44 float32 = 1.5
                                    v45 := []int{fi19, v2}
                                    _ = v44
                                    _ = v45
                                } else {
                                    v46 := GenS1{F0: (18446744073709551615 != 65535), F1: 0, F2: 821}
                                    _ = v46
                                }
                                var v47 int8 = (v3 >> ((594) & 63))
                                _ = v47
                            }
                        }
                        _ = v35
                    }
                    _ = v20
                }
                v1 = (("hello" + v1) + v1)
                if true {
                    var v48 float32 = 0.1
                    var v49 uint64 = 42
                    v49 = 47
                    _ = v48
                    _ = v49
                } else {
                    if (false || (-2147483648 == -2147483648)) {
                        var v50 uint32 = 4294967295
                        v50 = v50
                        switch 42 {
                        case -729:
                            v50 = v50
                            switch "foo bar" {
                            case "hello":
                                var v51 int8 = v3
                                _ = v51
                            }
                            switch (v1 + "hello") {
                            case "":
                                v52 := []int{fi15, 9223372036854775807, fi15, v2}
                                _ = v52
                            }
                        case 2147483647:
                            v1 = "foo bar"
                        case -1:
                            fmt.Println(v50, v3)
                            v53 := GenS1{F0: true, F1: 101, F2: 65535}
                            if (v50 != (v50 ^ genFn0())) {
                                var v54 bool = false
                                v55 := []int{}
                                v50 += (v50 ^ 2147483648)
                                _ = v54
                                _ = v55
                            } else {
                                v56 := []int{-1, 42, (^(v2)), fi19}
                                fmt.Println(v1, v2, v1)
                                _ = v56
                            }
                            _ = v53
                        }
                        _ = v50
                    } else {
                        fmt.Println(fi15)
                        v1 = v1
                    }
                    switch (v2 | (fi15 << ((1) & 63))) {
                    case 2147483647:
                        switch -711 {
                        case 464:
                            v57 := []int{fi15, 0}
                            _ = v57
                        case 287:
                            v1 = v1
                        }
                    case 42:
                        v2 = 9223372036854775807
                    case -268:
                        switch fi15 {
                        case 1:
                            if (-374 <= 1) {
                                v1 = ("foo bar" + ("a" + ""))
                                v3 += v3
                            } else {
                                switch v1 {
                                case "foo bar":
                                    v58 := []int{fi19, fi19, (fi15 - fi19), fi19}
                                    _ = v58
                                default:
                                    v3 = (v3 % ((v3) | 1))
                                }
                                v59 := []int{int(v3)}
                                _ = v59
                            }
                        case 32767:
                            v3 = (v3 << ((484) & 63))
                        }
                    default:
                        v60 := []int{(^(fi19)), fi15, fi19}
                        _ = v60
                    }
                    v61 := GenS2{F0: (("" + "a") + v1), F1: 16, F2: 169}
                    _ = v61
                }
                _ = fi19
            }
        }
        _ = fi15
        _ = v16
    }
    v2 += (v2 << ((182) & 63))
    var v62 uint8 = 128
    if (121 >= 1) {
        fmt.Println(v3, v2)
    }
    fmt.Println(v2)
    for fi63 := 0; fi63 < 20; fi63++ {
        v2 -= 32768
        v62 -= (^(v62))
        v64 := GenS3{F0: v62, F1: (-(fi63))}
        _ = fi63
        _ = v64
    }
    v62 -= v62
    if true {
        v3 += -73
    } else {
        if (65535 == 255) {
            v3 = v3
            var v65 uint8 = 128
            var v66 string = "\n\t"
            _ = v65
            _ = v66
        }
        for fi67 := 0; fi67 < 15; fi67++ {
            for fi68 := 0; fi68 < 4; fi68++ {
                if ((v1 + (v1 + (v1 + ("\n\t" + (v1 + (v1 + v1)))))) > v1) {
                    v62 = v62
                    if true {
                        v1 = ((((("foo bar" + v1) + (("a" + ("foo bar" + "hello")) + (v1 + v1))) + ((v1 + v1) + v1)) + v1) + "\n\t")
                        switch fi67 {
                        case 2147483647:
                            var v69 float64 = float64(0.0)
                            _ = v69
                        default:
                            fmt.Println(fi67)
                        }
                        v3 = v3
                    }
                    v70 := []int{0, (fi68 >> ((4294967295) & 63))}
                    _ = v70
                }
                v71 := GenS2{F0: "foo bar", F1: 221, F2: v62}
                _ = fi68
                _ = v71
            }
            v72 := []int{(v2 << ((65535) & 63))}
            for fi73 := 0; fi73 < 3; fi73++ {
                v74 := GenS2{F0: v1, F1: v62, F2: (v62 / (((v62 ^ 128)) | 1))}
                _ = fi73
                _ = v74
            }
            _ = fi67
            _ = v72
        }
    }
    v75 := GenS4{F0: -2147483648, F1: (v62 - v62), F2: ""}
    switch (v2 ^ 392) {
    case 73:
        var v76 uint = 88
        _ = v76
    }
    v77 := []int{4294967295, v2, v2}
    var v78 uint16 = 64
    if len(v77) > 0 {
        v77[0] = v2
    }
    for fi79 := 0; fi79 < 16; fi79++ {
        fmt.Println(fi79, v78)
        if ((!false) && (0 <= int64(v3))) {
            if ((true && (true && false)) || true) {
                var v80 uint64 = 0
                v62 = 141
                _ = v80
            }
            v3 = (v3 * (v3 >> ((907) & 63)))
        } else {
            v81 := GenS5{F0: (0 <= 0), F1: 607, F2: v2}
            _ = v81
        }
        fmt.Println(v3, fi79)
        _ = fi79
    }
    var v82 uint8 = (v62 + (v62 * v62))
    if true {
        if false {
            var v83 bool = (true && ((v2 >> ((42) & 63)) > v2))
            v62 -= (^(v62))
            var v84 string = v1
            _ = v83
            _ = v84
        } else {
            v85 := GenS5{F0: (((1.5 / (3.14159 * 0.1)) <= (float32(3.14159) - float32((1.5 / 0.1)))) && (!(!false))), F1: v78, F2: -971}
            var v86 uint8 = (v62 >> ((255) & 63))
            _ = v85
            _ = v86
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v77:", len(v77), v77)
    fmt.Println("v78:", v78)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v62
    _ = v75
    _ = v77
    _ = v78
    _ = v82
}
