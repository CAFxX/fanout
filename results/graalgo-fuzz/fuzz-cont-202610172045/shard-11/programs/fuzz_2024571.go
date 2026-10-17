package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32; F1 uint16; F2 bool }
type GenS8 struct { F0 uint32; F1 uint16; F2 int8 }
func genFn0(p0 uint8, p1 int16) uint16 {
    v6 := GenS1{F0: 42, F1: 0}
    v7 := []int{994, 32768}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 55
}
func genFn1() bool {
    var v8 uint8 = 1
    v8 += v8
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int32 = 785
        if (42 > (v8 << ((4294967295) & 63))) {
            for fi11 := 0; fi11 < 13; fi11++ {
                v8 = 115
                _ = fi11
            }
            fmt.Println(v8, v8, v10)
            for fi12 := 0; fi12 < 17; fi12++ {
                v10 = v10
                v13 := []int{(fi12 << ((2) & 63)), (fi12 & (-(fi12))), fi12}
                if false {
                    var v14 float64 = 0.1
                    v8 -= v8
                    v14 = ((0.1 / 1.5) * (v14 + (v14 * -2.25)))
                    _ = v14
                } else {
                    switch fi9 {
                    case -798:
                        _ = len(v13)
                    }
                    if true {
                        _ = len(v13)
                        if true {
                            fmt.Println(fi9, v10)
                            switch "\n\t" {
                            case "foo bar":
                                v13 = append(v13, fi9)
                            case "hello":
                                if false {
                                    fmt.Println(v10)
                                    fmt.Println(v10, v10, v10)
                                }
                            case "a":
                                var v15 int32 = (v10 << ((827) & 63))
                                _ = v15
                            }
                            var v16 float64 = (1.5 - 3.14159)
                            _ = v16
                        } else {
                            if len(v13) > 0 {
                                v13[0] = (fi12 % (((fi9 % (((fi9 * (fi12 | fi9))) | 1))) | 1))
                            }
                            v17 := GenS1{F0: 138, F1: genFn0(42, 6)}
                            fmt.Println(fi12)
                            _ = v17
                        }
                        if false {
                            if false {
                                if true {
                                    if true {
                                        if true {
                                            var v18 uint32 = 548
                                            _ = v18
                                        } else {
                                            v10 = (^(v10))
                                        }
                                        if true {
                                            var v19 int = fi9
                                            v20 := GenS1{F0: 399, F1: 1}
                                            var v21 uint16 = genFn0((v8 >> ((56) & 63)), -1)
                                            _ = v19
                                            _ = v20
                                            _ = v21
                                        } else {
                                            var v22 uint8 = (v8 & 1)
                                            v8 += (v22 + 62)
                                            if (!false) {
                                                var v23 int8 = 42
                                                switch ((("hello" + "a") + (("foo bar" + ((("a" + ("a" + "a")) + ((("" + "foo bar") + "a") + "")) + "")) + (("" + "foo bar") + ("foo bar" + "foo bar")))) + ("" + ("\n\t" + ("foo bar" + (("foo bar" + "a") + ""))))) {
                                                case "hello":
                                                    var v24 int32 = v10
                                                    _ = v24
                                                case "foo bar":
                                                    v10 = (v10 >> ((1) & 63))
                                                case "a":
                                                    v22 = (^(v22))
                                                }
                                                _ = v23
                                            } else {
                                                v8 += (v22 | v22)
                                                var v25 int16 = -32768
                                                var v26 uint16 = 1
                                                _ = v25
                                                _ = v26
                                            }
                                            switch ("foo bar" + (("\n\t" + "foo bar") + (("" + "foo bar") + (("hello" + "foo bar") + "foo bar")))) {
                                            case "":
                                                v8 = v8
                                            case "hello":
                                                switch ("hello" + "\n\t") {
                                                case "\n\t":
                                                    var v27 uint = 302
                                                    _ = v27
                                                case "":
                                                    v28 := GenS1{F0: 415, F1: 255}
                                                    _ = v28
                                                }
                                            case "a":
                                                var v29 uint = 206
                                                v29 = (v29 << ((383) & 63))
                                                switch 128 {
                                                case 42:
                                                    var v30 int32 = -2147483648
                                                    _ = v30
                                                case -9223372036854775808:
                                                    var v31 uint64 = 42
                                                    _ = v31
                                                case 4294967295:
                                                    var v32 uint16 = 0
                                                    _ = v32
                                                }
                                                var v33 bool = (1 != 42)
                                                _ = v29
                                                _ = v33
                                            default:
                                                switch -2147483648 {
                                                case -9223372036854775808:
                                                    fmt.Println(v10)
                                                default:
                                                    switch "hello" {
                                                    case "a":
                                                        var v34 int32 = (^(v10))
                                                        _ = v34
                                                    }
                                                }
                                            }
                                            _ = v22
                                        }
                                    } else {
                                        var v35 bool = ((true && (99 >= 1)) || (0.0 <= (1.5 + 1.5)))
                                        _ = v35
                                    }
                                } else {
                                    v13 = append(v13, fi9)
                                    v36 := GenS2{F0: 4294967295, F1: 258, F2: 32767}
                                    _ = v36
                                }
                            }
                            var v37 int = fi12
                            switch (fi12 * (^(fi9))) {
                            case 127:
                                v38 := GenS1{F0: 18446744073709551615, F1: 65535}
                                switch -235 {
                                case 128:
                                    if (fi12 != (fi12 % ((42) | 1))) {
                                        if true {
                                            v39 := GenS2{F0: 2147483647, F1: 42, F2: int16(v8)}
                                            _ = v39
                                        }
                                        var v40 uint = 255
                                        _ = v40
                                    }
                                case 42:
                                    var v41 int64 = 1
                                    _ = v41
                                case 1:
                                    if (v10 == (^(v10))) {
                                        v42 := GenS2{F0: 42, F1: uint64(v8), F2: 42}
                                        v8 = (v8 & (v8 & 127))
                                        _ = v42
                                    } else {
                                        var v43 int64 = 1
                                        var v44 uint64 = 1
                                        _ = v43
                                        _ = v44
                                    }
                                default:
                                    v13 = append(v13, 81)
                                }
                                _ = v38
                            case -9223372036854775808:
                                if true {
                                    var v45 bool = true
                                    v13 = append(v13, fi9)
                                    var v46 uint = 255
                                    _ = v45
                                    _ = v46
                                }
                            case 4294967295:
                                v47 := GenS3{F0: ((3.14159 == float32((0.0 - 3.14159))) && (281 > 18446744073709551615)), F1: int32(fi12)}
                                _ = v47
                            }
                            _ = v37
                        } else {
                            _ = len(v13)
                            fmt.Println(v10, fi9, fi9)
                        }
                    }
                }
                _ = fi12
                _ = v13
            }
        }
        _ = fi9
        _ = v10
    }
    _ = v8
    return (0 != (v8 & 2))
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 0
    v3 := make(map[int]int)
    v4 := []string{"a"}
    var v5 uint16 = 65535
    v1 -= (v1 ^ v1)
    var v48 float64 = (0.1 / 0.1)
    v4 = append(v4, "")
    var v49 int16 = 0
    if true {
        v50 := GenS4{F0: false, F1: 2}
        v5 = v5
        _ = v50
    } else {
        if false {
            fmt.Println(v2, v49, v48)
        }
        var v51 uint64 = v2
        _ = v51
    }
    var v52 bool = (226 != 124)
    var v53 float32 = 1e10
    if (-1 <= 1) {
        var v54 uint32 = 178
        var v55 int16 = 898
        _ = v54
        _ = v55
    } else {
        for fi56 := 0; fi56 < 16; fi56++ {
            if v52 {
                if (fi56 > 4294967295) {
                    for fi57 := 0; fi57 < 19; fi57++ {
                        var v58 int8 = -1
                        _ = fi57
                        _ = v58
                    }
                    if false {
                        if true {
                            var v59 uint32 = v1
                            _ = v59
                        }
                        for fi60 := 0; fi60 < 14; fi60++ {
                            fmt.Println(v49)
                            fmt.Println(v52)
                            if (97 < 68) {
                                switch "" {
                                case "\n\t":
                                    var v61 int = (fi56 - (^(fi56)))
                                    v2 = (v2 / ((v2) | 1))
                                    _ = v61
                                }
                            }
                            fmt.Println(v49)
                            switch ("foo bar" + "foo bar") {
                            case "\n\t":
                                if v52 {
                                    v62 := GenS1{F0: 556, F1: (v5 ^ v5)}
                                    switch "a" {
                                    case "\n\t":
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v2)
                                        case "hello":
                                            if false {
                                                switch "a" {
                                                case "hello":
                                                    v52 = (!false)
                                                    v1 = v1
                                                case "":
                                                    if false {
                                                        delete(v3, fi56)
                                                        fmt.Println(fi60, v53, v48)
                                                    } else {
                                                        var v63 uint16 = (v5 << ((uint(v5)) & 63))
                                                        switch "a" {
                                                        case "":
                                                            _ = len(v4)
                                                        case "\n\t":
                                                            fmt.Println(v53, v52, v5)
                                                        }
                                                        fmt.Println(fi56, fi60)
                                                        var v64 float32 = -2.25
                                                        v53 = (1.5 + (((((v64 / v53) - v53) + v64) + 0.0) * v53))
                                                        v52 = v52
                                                        var v65 float64 = v48
                                                        var v66 int = fi60
                                                        _ = v63
                                                        _ = v64
                                                        _ = v65
                                                        _ = v66
                                                    }
                                                }
                                                v48 = v48
                                                if (((!((v2 - 2) != v2)) || ((-1 > 0) || (!("\n\t" == "hello")))) && (fi56 != int(v1))) {
                                                    var v67 float64 = v48
                                                    var v68 int32 = 193
                                                    v48 = -2.25
                                                    _ = v67
                                                    _ = v68
                                                }
                                                if (v48 >= v48) {
                                                    var v69 uint16 = 65535
                                                    var v70 uint16 = v5
                                                    v48 = (v48 * v48)
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v2 = v2
                                                    if true {
                                                        v48 -= (v48 * v48)
                                                    }
                                                }
                                                v71 := GenS5{F0: 42, F1: v5}
                                                _ = v71
                                            } else {
                                                v49 = v49
                                                switch (fi56 ^ (fi60 % ((fi56) | 1))) {
                                                case 685:
                                                    v49 -= 0
                                                case -661:
                                                    v72 := v3[(fi60 / (((fi60 / ((fi60) | 1))) | 1))]
                                                    _ = v72
                                                case 9223372036854775807:
                                                    var v73 uint64 = (v2 & v2)
                                                    _ = v73
                                                }
                                            }
                                        case "foo bar":
                                            switch "a" {
                                            case "foo bar":
                                                fmt.Println(v2)
                                            case "\n\t":
                                                v74 := GenS6{F0: 127}
                                                _ = v74
                                            }
                                            v75 := GenS1{F0: 255, F1: v5}
                                            fmt.Println(fi56, v53)
                                            _ = v75
                                        default:
                                            var v76 int = -2147483648
                                            _ = v76
                                        }
                                    case "":
                                        switch ("hello" + "foo bar") {
                                        case "a":
                                            if len(v4) > 0 {
                                                v4[0] = "\n\t"
                                            }
                                        case "hello":
                                            if true {
                                                _ = len(v3)
                                            } else {
                                                v52 = (!false)
                                                v52 = false
                                            }
                                        default:
                                            v77 := GenS7{F0: 228, F1: (v5 & (v5 | (^(v5)))), F2: ((-2.25 / (0.0 + (v53 * v53))) < (3.14159 + ((float32((v48 / v48)) * v53) * (v53 / v53))))}
                                            _ = v77
                                        }
                                    case "a":
                                        if (v52 && ((fi56 & fi56) < (-(fi60)))) {
                                            if v52 {
                                                if (-1 == 9223372036854775807) {
                                                    v2 -= 1
                                                }
                                            }
                                            var v78 uint64 = 848
                                            var v79 uint = 65535
                                            var v80 uint32 = v1
                                            _ = v78
                                            _ = v79
                                            _ = v80
                                        }
                                    default:
                                        if false {
                                            if true {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v81 := GenS8{F0: (^(v1)), F1: v5, F2: -1}
                                                v48 += v48
                                                if (4294967295 != 42) {
                                                    if v52 {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            fmt.Println(v1, fi60, v48)
                                                        default:
                                                            delete(v3, fi60)
                                                        }
                                                        var v82 string = "foo bar"
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v83 uint32 = 1
                                                        _ = v82
                                                        _ = v83
                                                    } else {
                                                        v4 = append(v4, "hello")
                                                        v2 = (v2 + (^(v2)))
                                                    }
                                                    fmt.Println(v52)
                                                    var v84 float32 = v53
                                                    var v85 float32 = 3.14159
                                                    _ = v84
                                                    _ = v85
                                                } else {
                                                    switch ("a" + "a") {
                                                    case "foo bar":
                                                        _ = len(v3)
                                                    default:
                                                        _ = len(v3)
                                                    }
                                                }
                                                _ = v81
                                            } else {
                                                v2 = v2
                                                var v86 uint16 = (v5 % (((v5 >> ((932) & 63))) | 1))
                                                _ = v86
                                            }
                                            fmt.Println(v2, v52)
                                            var v87 float64 = v48
                                            if v52 {
                                                delete(v3, (-(fi60)))
                                            }
                                            _ = v87
                                        }
                                    }
                                    _ = v62
                                }
                            case "":
                                switch "\n\t" {
                                case "hello":
                                    v88 := GenS4{F0: false, F1: 42}
                                    _ = v88
                                }
                            case "foo bar":
                                if (-820 > 0) {
                                    if (0 == -185) {
                                        var v89 float64 = ((v48 / 1.5) - (v48 + (v48 - -2.25)))
                                        _ = v89
                                    } else {
                                        var v90 uint16 = v5
                                        var v91 uint32 = 0
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v90
                                        _ = v91
                                    }
                                }
                            }
                            _ = fi60
                        }
                    } else {
                        fmt.Println(fi56)
                        if (42 >= 410) {
                            if v52 {
                                v92 := GenS8{F0: 42, F1: (v5 - (v5 + v5)), F2: 1}
                                v49 += (v49 - (v49 & 1000))
                                var v93 string = "\n\t"
                                _ = v92
                                _ = v93
                            }
                            v94 := GenS4{F0: v52, F1: 2}
                            v2 = 2
                            _ = v94
                        }
                        var v95 uint16 = (^(v5))
                        _ = v95
                    }
                }
            }
            _ = fi56
        }
        v5 += v5
        var v96 string = "foo bar"
        _ = v96
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v53
}
