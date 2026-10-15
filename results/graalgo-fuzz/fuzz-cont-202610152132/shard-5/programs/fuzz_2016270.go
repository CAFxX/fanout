package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 int64; F1 uint16; F2 int64 }
type GenS8 struct { F0 uint8; F1 uint16; F2 int16 }
type GenS9 struct { F0 int16; F1 uint; F2 int }
type GenS10 struct { F0 uint32; F1 int16; F2 uint }
type GenS11 struct { F0 uint16; F1 int; F2 uint32 }
type GenS12 struct { F0 string }
func genFn0() uint8 {
    var v6 string = ("\n\t" + "\n\t")
    v7 := []int{32768}
    var v8 int32 = 2147483647
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v6
    _ = v7
    _ = v8
    return 246
}
func genFn1() string {
    v9 := GenS1{F0: -1, F1: 42, F2: (1 != 612)}
    var v10 float32 = (0.0 / (float32(-2.25) + 1.5))
    _ = v10
    _ = v9
    return (("foo bar" + ("\n\t" + (("a" + "hello") + ("foo bar" + ("a" + ""))))) + "")
}

func main() {
    var v1 float64 = (0.1 / 1e10)
    v2 := []uint32{1, 727}
    v3 := []uint32{2147483648, 2147483647, 906, 885}
    var v4 int8 = 42
    var v5 uint8 = 127
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int64 = 0
        _ = fi11
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 4; fi13++ {
            switch "foo bar" {
            case "foo bar":
                switch fi13 {
                case -9223372036854775808:
                    switch "\n\t" {
                    case "foo bar":
                        v14 := GenS2{F0: fi13}
                        _ = v14
                    default:
                        var v15 int32 = 1
                        _ = v15
                    }
                case 127:
                    v2 = append(v2, 2)
                case -94:
                    switch ((("a" + "foo bar") + "foo bar") + (("foo bar" + "hello") + "\n\t")) {
                    case "a":
                        if (0.0 > (-2.25 / 1.0)) {
                            if (v4 != (v4 % ((v4) | 1))) {
                                if false {
                                    var v16 uint16 = 1
                                    if false {
                                        switch ("" + "foo bar") {
                                        case "\n\t":
                                            fmt.Println(v5)
                                        case "foo bar":
                                            if false {
                                                switch (fi13 << ((565) & 63)) {
                                                case 32767:
                                                    v1 -= 3.14159
                                                }
                                                var v17 uint8 = (^(v5))
                                                _ = v17
                                            } else {
                                                for fi18 := 0; fi18 < 13; fi18++ {
                                                    if true {
                                                        var v19 float32 = 3.14159
                                                        var v20 uint32 = 0
                                                        fmt.Println(v4, v20)
                                                        v21 := GenS3{F0: 0, F1: -2147483648, F2: 2}
                                                        _ = v19
                                                        _ = v20
                                                        _ = v21
                                                    } else {
                                                        var v22 int16 = 0
                                                        _ = v22
                                                    }
                                                    _ = fi18
                                                }
                                                if true {
                                                    for fi23 := 0; fi23 < 6; fi23++ {
                                                        var v24 int32 = -787
                                                        _ = fi23
                                                        _ = v24
                                                    }
                                                    var v25 int8 = -128
                                                    _ = v25
                                                } else {
                                                    if true {
                                                        for fi26 := 0; fi26 < 19; fi26++ {
                                                            v4 = v4
                                                            if len(v3) > 0 {
                                                                v3[0] = 65
                                                            }
                                                            if true {
                                                                switch 127 {
                                                                case 4294967295:
                                                                    if (42 >= v16) {
                                                                        switch fi26 {
                                                                        case -479:
                                                                            if len(v2) > 0 {
                                                                                v2[0] = 1
                                                                            }
                                                                        case 2147483647:
                                                                            switch "a" {
                                                                            case "foo bar":
                                                                                var v27 float32 = 1e10
                                                                                if true {
                                                                                    var v28 bool = (false && ((!(18446744073709551615 != 18446744073709551615)) && (uint8(v16) == 179)))
                                                                                    fmt.Println(v5, v4)
                                                                                    v16 = (v16 >> ((692) & 63))
                                                                                    _ = v28
                                                                                }
                                                                                v29 := GenS4{F0: v4, F1: genFn0()}
                                                                                _ = v27
                                                                                _ = v29
                                                                            case "hello":
                                                                                fmt.Println(v4, v1, fi26)
                                                                                var v30 int16 = int16(fi26)
                                                                                _ = v30
                                                                            case "":
                                                                                fmt.Println(v4, v16, fi13)
                                                                            }
                                                                        default:
                                                                            switch (fi26 & (^(fi13))) {
                                                                            case 872:
                                                                                fmt.Println(v16)
                                                                            case 32768:
                                                                                v5 = (v5 & (v5 ^ genFn0()))
                                                                            default:
                                                                                v31 := GenS1{F0: (fi13 << ((255) & 63)), F1: (v5 + (v5 >> ((42) & 63))), F2: false}
                                                                                _ = v31
                                                                            }
                                                                        }
                                                                        v1 -= v1
                                                                        var v32 int = (^(fi13))
                                                                        _ = v32
                                                                    } else {
                                                                        var v33 float32 = 1.5
                                                                        v1 += v1
                                                                        if false {
                                                                            v33 = ((v33 * v33) / (v33 + v33))
                                                                            v2 = append(v2, 2147483647)
                                                                            var v34 int32 = -753
                                                                            v33 = 3.14159
                                                                            _ = len(v2)
                                                                            _ = v34
                                                                        } else {
                                                                            v35 := GenS5{F0: (v5 * genFn0()), F1: -32768, F2: (v16 ^ v16)}
                                                                            if false {
                                                                                var v36 uint32 = 42
                                                                                _ = v36
                                                                            }
                                                                            _ = v35
                                                                        }
                                                                        if (true && (false || (1000 >= 32767))) {
                                                                            var v37 float32 = (0.1 + v33)
                                                                            v38 := GenS6{F0: 965, F1: 42}
                                                                            _ = v37
                                                                            _ = v38
                                                                        }
                                                                        _ = v33
                                                                    }
                                                                case 594:
                                                                    var v39 uint = 65535
                                                                    _ = v39
                                                                case -2147483648:
                                                                    if (1 != 4294967295) {
                                                                        var v40 uint32 = 65535
                                                                        var v41 int32 = -193
                                                                        v42 := GenS6{F0: -1, F1: 42}
                                                                        _ = v40
                                                                        _ = v41
                                                                        _ = v42
                                                                    } else {
                                                                        v43 := GenS7{F0: -837, F1: (v16 * (v16 * (v16 / ((v16) | 1)))), F2: -1}
                                                                        _ = v43
                                                                    }
                                                                }
                                                            }
                                                            _ = fi26
                                                        }
                                                    }
                                                    v44 := GenS1{F0: fi13, F1: v5, F2: (!false)}
                                                    v5 += v5
                                                    _ = v44
                                                }
                                            }
                                        case "a":
                                            switch "hello" {
                                            case "\n\t":
                                                var v45 string = "foo bar"
                                                _ = v45
                                            }
                                        default:
                                            if true {
                                                fmt.Println(v5)
                                            } else {
                                                var v46 int32 = -1
                                                v46 = 2147483647
                                                fmt.Println(v16, v1, v46)
                                                _ = v46
                                            }
                                        }
                                        if (401 == uint64(v16)) {
                                            fmt.Println(v5)
                                        } else {
                                            for fi47 := 0; fi47 < 20; fi47++ {
                                                var v48 float32 = ((1e10 / (3.14159 * 3.14159)) + (1e10 / ((3.14159 * ((-2.25 - -2.25) / 1.5)) - -2.25)))
                                                if (!true) {
                                                    v48 = (v48 + (float32(1e10) * 0.0))
                                                    if (!(("\n\t" <= "foo bar") && false)) {
                                                        fmt.Println(v48, v5)
                                                        var v49 uint8 = (^(v5))
                                                        _ = v49
                                                    }
                                                    fmt.Println(v16)
                                                    var v50 int8 = (v4 | v4)
                                                    if true {
                                                        v4 += (v4 - -1)
                                                    } else {
                                                        v5 = (v5 % ((v5) | 1))
                                                        var v51 string = ("\n\t" + ("" + "a"))
                                                        _ = len(v2)
                                                        _ = v51
                                                    }
                                                    if ((true || true) && false) {
                                                        v52 := GenS8{F0: v5, F1: v16, F2: 305}
                                                        fmt.Println(v48, v4, fi13)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v52
                                                    } else {
                                                        switch "a" {
                                                        case "foo bar":
                                                            if len(v2) > 0 {
                                                                v2[0] = 114
                                                            }
                                                        }
                                                        var v53 int8 = int8(v16)
                                                        if (!(v16 < v16)) {
                                                            v5 += 42
                                                        } else {
                                                            if len(v3) > 0 {
                                                                v3[0] = 4294967295
                                                            }
                                                        }
                                                        _ = v53
                                                    }
                                                    v2 = append(v2, 2)
                                                    var v54 int8 = (v4 | (^(v4)))
                                                    _ = v50
                                                    _ = v54
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = 2147483648
                                                }
                                                _ = fi47
                                                _ = v48
                                            }
                                            _ = len(v3)
                                        }
                                        if true {
                                            v55 := GenS1{F0: (fi13 | (fi13 & fi13)), F1: (^(v5)), F2: (float32(0.1) < 3.14159)}
                                            _ = v55
                                        } else {
                                            fmt.Println(v1, v1)
                                            switch "\n\t" {
                                            case "a":
                                                v4 -= int8(fi13)
                                            case "\n\t":
                                                fmt.Println(v5)
                                            }
                                        }
                                    }
                                    v5 = 86
                                    _ = v16
                                }
                                v56 := GenS6{F0: 0, F1: 0}
                                v5 += (^(v5))
                                _ = v56
                            }
                            fmt.Println(v5)
                        }
                    }
                }
            default:
                fmt.Println(v5, v5)
            }
            var v57 uint64 = 42
            _ = fi13
            _ = v57
        }
    } else {
        if (42 < 74) {
            switch -2147483648 {
            case -9223372036854775808:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case -42:
                fmt.Println(v5, v4, v5)
            }
            v1 = 0.0
            v58 := GenS9{F0: -1, F1: 793, F2: -43}
            _ = v58
        } else {
            v4 -= (^(v4))
            if false {
                fmt.Println(v4, v5, v4)
                switch "\n\t" {
                case "hello":
                    v5 = (v5 | genFn0())
                }
                var v59 int16 = -100
                _ = v59
            }
            v4 = v4
        }
        if (!true) {
            var v60 uint32 = 4294967295
            fmt.Println(v4, v5)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v60
        } else {
            var v61 int64 = -9223372036854775808
            v2 = append(v2, uint32(v5))
            _ = v61
        }
        var v62 float64 = 0.0
        _ = v62
    }
    v63 := GenS5{F0: v5, F1: 633, F2: 1}
    for fi64 := 0; fi64 < 8; fi64++ {
        var v65 uint64 = 633
        v66 := GenS10{F0: 65535, F1: -1, F2: 976}
        v65 = 382
        _ = fi64
        _ = v65
        _ = v66
    }
    fmt.Println(v5)
    var v67 int32 = 42
    for fi68 := 0; fi68 < 20; fi68++ {
        for fi69 := 0; fi69 < 11; fi69++ {
            var v70 int64 = 0
            v67 -= 2147483647
            var v71 uint = 18446744073709551615
            _ = fi69
            _ = v70
            _ = v71
        }
        for fi72 := 0; fi72 < 19; fi72++ {
            _ = len(v2)
            if true {
                fmt.Println(v5, fi68, v67)
                switch 32768 {
                case 2147483647:
                    fmt.Println(fi72)
                case 4294967295:
                    if false {
                        var v73 int16 = 92
                        if (false && (false && ((v67 + (v67 ^ 0)) <= (-(v67))))) {
                            switch "hello" {
                            case "foo bar":
                                var v74 uint32 = 4294967295
                                _ = v74
                            case "":
                                v75 := GenS3{F0: 4294967295, F1: (v67 << ((4294967295) & 63)), F2: 259}
                                _ = v75
                            }
                            var v76 uint32 = 42
                            _ = v76
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v77 uint32 = 0
                            _ = v77
                        }
                        _ = v73
                    }
                }
                var v78 uint = 18446744073709551615
                _ = v78
            } else {
                v79 := GenS11{F0: 65535, F1: fi68, F2: 2147483647}
                switch 2147483647 {
                case -1:
                    switch "\n\t" {
                    case "hello":
                        v80 := GenS12{F0: "\n\t"}
                        _ = v80
                    case "\n\t":
                        var v81 uint64 = 2
                        _ = v81
                    case "":
                        fmt.Println(fi72)
                    default:
                        var v82 bool = false
                        _ = v82
                    }
                default:
                    fmt.Println(v67)
                }
                v4 = 100
                _ = v79
            }
            _ = fi72
        }
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            v1 = v1
        case "hello":
            if (42 > 1) {
                if true {
                    v1 = 1e10
                    switch fi68 {
                    case 4294967295:
                        var v83 bool = false
                        _ = v83
                    default:
                        v67 = (v67 + (-(v67)))
                    }
                }
                v1 = (v1 - v1)
            }
        case "a":
            _ = len(v3)
        default:
            var v84 uint8 = v5
            _ = v84
        }
        _ = fi68
    }
    v67 += (v67 >> ((255) & 63))
    if (true || (false && false)) {
        var v85 float32 = (float32((float64(1e10) * 1.5)) + (1.5 * 3.14159))
        _ = v85
    } else {
        var v86 int = 0
        _ = v86
    }
    switch "\n\t" {
    case "hello":
        v87 := GenS12{F0: "foo bar"}
        _ = v87
    case "foo bar":
        v4 = -1
    default:
        v1 -= (((v1 + -2.25) - v1) + (0.1 * 0.1))
    }
    var v88 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    fmt.Println("v88:", v88)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v67
    _ = v88
}
