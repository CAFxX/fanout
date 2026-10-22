package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 string; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int16; F1 int; F2 uint32 }
func genFn0() int64 {
    var v7 int8 = -113
    if true {
        v8 := []int{-9223372036854775808, -2147483648}
        var v9 int = 0
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: ""}
        var v11 int64 = 0
        _ = v10
        _ = v11
    }
    var v12 int = 4294967295
    for fi13 := 0; fi13 < 11; fi13++ {
        v7 += (^(v7))
        v14 := []int{fi13, v12}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := []int{}
        _ = len(v16)
        var v17 int32 = -982
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = v12
    _ = v7
    return 1
}
func genFn1(p0 int32, p1 int8, p2 int16) int {
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := []int{-1, fi18}
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        fmt.Println(fi20, p0)
        v21 := []int{(fi20 ^ (fi20 ^ 0))}
        var v22 int32 = p0
        _ = fi20
        _ = v21
        _ = v22
    }
    if true {
        v23 := []int{-1}
        switch int(p1) {
        case -39:
            var v24 int64 = -9223372036854775808
            _ = v24
        default:
            for fi25 := 0; fi25 < 5; fi25++ {
                switch ("" + (("a" + "foo bar") + ("\n\t" + "foo bar"))) {
                case "":
                    if false {
                        for fi26 := 0; fi26 < 9; fi26++ {
                            if len(v23) > 0 {
                                v23[0] = (fi25 ^ fi26)
                            }
                            fmt.Println(p1, fi26)
                            _ = fi26
                        }
                    } else {
                        fmt.Println(p0, fi25)
                        p1 = (p1 & (p1 - (p1 & (p1 | (p1 % ((104) | 1))))))
                        if false {
                            for fi27 := 0; fi27 < 2; fi27++ {
                                if true {
                                    v28 := GenS1{F0: "hello"}
                                    p1 = p1
                                    _ = v28
                                } else {
                                    switch -248 {
                                    case 0:
                                        v29 := GenS1{F0: ""}
                                        _ = v29
                                    case 4294967295:
                                        fmt.Println(p1, p0)
                                    case -9223372036854775808:
                                        var v30 int = -941
                                        _ = v30
                                    }
                                    if false {
                                        var v31 string = "hello"
                                        v31 = (("hello" + "\n\t") + v31)
                                        p0 -= (-(p0))
                                        _ = v31
                                    }
                                    var v32 int64 = -9223372036854775808
                                    if len(v23) > 0 {
                                        v23[0] = (fi25 + (fi27 - fi27))
                                    }
                                    _ = v32
                                }
                                var v33 uint32 = 4294967295
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                                _ = fi27
                                _ = v33
                            }
                            if len(v23) >= 2 {
                                v23 = v23[1:len(v23)]
                            }
                            fmt.Println(fi25)
                        }
                    }
                default:
                    fmt.Println(fi25, p0, p2)
                }
                v23 = append(v23, fi25)
                var v34 int32 = (p0 - p0)
                _ = fi25
                _ = v34
            }
        }
        for fi35 := 0; fi35 < 2; fi35++ {
            v23 = append(v23, (fi35 << ((445) & 63)))
            p2 = (-(p2))
            _ = fi35
        }
        _ = v23
    }
    v36 := GenS1{F0: ("\n\t" + "")}
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 0.1
    var v3 int = 42
    v4 := make(map[string]string)
    var v5 int8 = 109
    var v6 int = (v3 & (v3 << ((1) & 63)))
    for fi37 := 0; fi37 < 12; fi37++ {
        v6 = genFn1(-317, (-(v5)), 1000)
        _ = fi37
    }
    if false {
        v38 := []int{(v3 | (v3 + genFn1(-2147483648, 1, -32768))), v6}
        if (!false) {
            var v39 int16 = -623
            v3 -= v6
            _ = v39
        }
        var v40 int64 = 1
        _ = v38
        _ = v40
    }
    v2 = 3.14159
    var v41 uint32 = 4294967295
    for fi42 := 0; fi42 < 9; fi42++ {
        for fi43 := 0; fi43 < 11; fi43++ {
            v44 := GenS1{F0: "a"}
            if true {
                if true {
                    var v45 uint8 = 0
                    v46 := GenS2{F0: genFn1(42, (v5 << ((18446744073709551615) & 63)), -32768), F1: (((("foo bar" + "foo bar") + "hello") + "hello") + "hello"), F2: 42}
                    _ = v45
                    _ = v46
                } else {
                    switch "hello" {
                    case "a":
                        v47 := []int{(v3 / ((v6) | 1)), (v6 & (fi43 * (v6 % ((32768) | 1)))), v6}
                        _ = v47
                    case "foo bar":
                        var v48 string = "\n\t"
                        _ = v48
                    }
                    fmt.Println(fi43, fi43)
                }
            } else {
                v2 = v2
            }
            var v49 int = genFn1(356, 127, -1)
            _ = fi43
            _ = v44
            _ = v49
        }
        switch 9223372036854775807 {
        case 0:
            v50 := []int{(v3 << ((792) & 63))}
            _ = v50
        default:
            var v51 int16 = -634
            _ = v51
        }
        _ = fi42
    }
    v52 := GenS3{F0: v5, F1: 255, F2: v41}
    v53 := []int{}
    switch (v3 - (v6 | genFn1(-2147483648, (v5 ^ (v5 + (v5 - v5))), -583))) {
    case 689:
        v1[v41] = (v41 >> ((117) & 63))
    case -515:
        if (4294967295 == 18446744073709551615) {
            for fi54 := 0; fi54 < 5; fi54++ {
                var v55 float64 = 1.5
                fmt.Println(v3)
                _ = fi54
                _ = v55
            }
            for fi56 := 0; fi56 < 15; fi56++ {
                v57 := GenS3{F0: 100, F1: 686, F2: (v41 ^ (^(v41)))}
                _ = fi56
                _ = v57
            }
        }
    default:
        if len(v53) >= 2 {
            v53 = v53[1:len(v53)]
        }
    }
    fmt.Println(v41)
    v58 := GenS1{F0: "\n\t"}
    fmt.Println(v3, v6)
    fmt.Println(v3)
    if len(v53) > 0 {
        v53[0] = -935
    }
    v41 += (v41 << ((4294967295) & 63))
    var v59 string = "a"
    if true {
        switch v59 {
        case "":
            v60 := GenS4{F0: -128, F1: (-9223372036854775808 != 9223372036854775807)}
            _ = v60
        case "foo bar":
            for fi61 := 0; fi61 < 11; fi61++ {
                var v62 int = v3
                if len(v53) >= 2 {
                    v53 = v53[1:len(v53)]
                }
                _ = fi61
                _ = v62
            }
        case "a":
            v41 -= v41
        }
        if false {
            v63 := GenS5{F0: 4294967295}
            _ = v63
        }
    } else {
        switch (v3 ^ (v3 >> ((1) & 63))) {
        case -120:
            if len(v53) >= 2 {
                v53 = v53[1:len(v53)]
            }
        default:
            for fi64 := 0; fi64 < 6; fi64++ {
                v59 = v59
                fmt.Println(v41)
                _ = fi64
            }
        }
        var v65 uint8 = 98
        switch genFn1(42, (-(v5)), 924) {
        case -609:
            var v66 int32 = -2147483648
            _ = v66
        default:
            for fi67 := 0; fi67 < 3; fi67++ {
                for fi68 := 0; fi68 < 9; fi68++ {
                    var v69 bool = (false && (901 > 2))
                    if true {
                        v2 -= v2
                        if v69 {
                            fmt.Println(v3, v6, v65)
                            if (v65 == 0) {
                                v70 := GenS3{F0: (v5 & v5), F1: 65535, F2: v41}
                                if (!(0 != 42)) {
                                    v71 := GenS1{F0: v59}
                                    v2 -= v2
                                    _ = v71
                                } else {
                                    fmt.Println(v3, v59)
                                    v72 := GenS1{F0: (v59 + "")}
                                    _ = v72
                                }
                                if v69 {
                                    v5 -= v5
                                }
                                _ = len(v4)
                                v59 = ("\n\t" + ((v59 + ((v59 + v59) + "\n\t")) + v59))
                                _ = v70
                            }
                            var v73 uint8 = v65
                            if ((-80 == -2147483648) || v69) {
                                fmt.Println(v3, v5, fi67)
                                switch (v59 + (((v59 + "hello") + ((v59 + v59) + v59)) + v59)) {
                                case "\n\t":
                                    v65 = 255
                                }
                                v53 = append(v53, fi68)
                            } else {
                                fmt.Println(v41)
                                switch (v59 + v59) {
                                case "hello":
                                    switch v59 {
                                    case "foo bar":
                                        v73 = 2
                                    }
                                }
                            }
                            v53 = append(v53, (v3 << ((0) & 63)))
                            v74 := GenS2{F0: fi68, F1: v59, F2: 693}
                            switch (("foo bar" + v59) + ("foo bar" + v59)) {
                            case "foo bar":
                                var v75 float64 = ((float64(-2.25) / v2) - v2)
                                _ = v75
                            case "hello":
                                switch genFn1(1, (v5 % ((v5) | 1)), 42) {
                                case 0:
                                    v76 := GenS5{F0: 1}
                                    _ = v76
                                case 42:
                                    fmt.Println(v41, v6)
                                    v77 := GenS6{F0: -2147483648, F1: 4294967295}
                                    _ = v77
                                case 910:
                                    if true {
                                        if (3.14159 == float32((0.1 * (3.14159 - float64(0.1))))) {
                                            v6 = v3
                                            var v78 uint32 = v41
                                            _ = v78
                                        }
                                        var v79 uint32 = 0
                                        _ = v79
                                    } else {
                                        if len(v53) > 0 {
                                            v53[0] = genFn1(-172, v5, -838)
                                        }
                                    }
                                default:
                                    var v80 string = v59
                                    _ = v80
                                }
                            }
                            _ = v73
                            _ = v74
                        }
                        if len(v53) >= 2 {
                            v53 = v53[1:len(v53)]
                        }
                        var v81 int8 = v5
                        fmt.Println(v81, v65)
                        if (!(836 <= -695)) {
                            v82 := GenS1{F0: v59}
                            _ = v82
                        } else {
                            v2 += v2
                            var v83 uint8 = v65
                            if len(v53) > 0 {
                                v53[0] = (v3 - v6)
                            }
                            if v69 {
                                if (v69 && v69) {
                                    v84 := GenS4{F0: -51, F1: (49 == 4294967295)}
                                    if (true || v69) {
                                        var v85 uint8 = v65
                                        _ = v85
                                    } else {
                                        var v86 bool = (255 > uint16(v65))
                                        v87 := GenS1{F0: "hello"}
                                        var v88 float64 = 1.5
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                    }
                                    var v89 string = v59
                                    _ = v84
                                    _ = v89
                                } else {
                                    v90 := GenS7{F0: int16(v41), F1: (v6 * (v3 * v3)), F2: 1}
                                    _ = v90
                                }
                                var v91 bool = v69
                                _ = v91
                            } else {
                                var v92 float32 = ((3.14159 / (((1e10 / -2.25) + 3.14159) / 3.14159)) * -2.25)
                                v41 -= 4294967295
                                _ = v92
                            }
                            _ = v83
                        }
                        var v93 int64 = 0
                        _ = v81
                        _ = v93
                    }
                    _ = fi68
                    _ = v69
                }
                var v94 uint16 = 255
                var v95 uint32 = 2147483647
                _ = fi67
                _ = v94
                _ = v95
            }
        }
        _ = v65
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v52
    _ = v53
    _ = v58
    _ = v59
    _ = v6
}
