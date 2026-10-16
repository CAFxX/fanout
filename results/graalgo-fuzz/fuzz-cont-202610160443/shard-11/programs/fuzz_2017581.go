package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int16; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string; F1 string }
func genFn0(p0 uint32, p1 int8) uint8 {
    fmt.Println(p1)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 uint32 = 65535
        v7 := GenS1{F0: (-(fi5))}
        v8 := GenS1{F0: fi5}
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return uint8(p0)
}

func main() {
    var v1 uint = 255
    var v2 uint64 = 978
    v3 := make(map[int]int)
    var v4 float32 = -2.25
    v9 := []int{128, 42, -2147483648, 32768}
    var v10 int64 = 9223372036854775807
    v11 := GenS2{F0: (42 >= -32768), F1: 32767, F2: -31}
    var v12 bool = (!true)
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int32 = 448
        _ = fi13
        _ = v14
    }
    var v15 uint8 = 189
    if (65535 > 42) {
        fmt.Println(v15, v10)
    }
    v12 = (!true)
    delete(v3, 42)
    v1 = (v1 * uint(v15))
    v3[42] = -9223372036854775808
    if v12 {
        var v16 string = "foo bar"
        v3[-2147483648] = -2147483648
        _ = v16
    }
    var v17 int32 = -910
    if ((v4 / ((v4 - 0.1) - v4)) != ((-2.25 + 3.14159) / v4)) {
        v12 = v12
        var v18 uint = 255
        var v19 string = "a"
        _ = v18
        _ = v19
    } else {
        v20 := GenS2{F0: true, F1: 957, F2: (^(v17))}
        switch "\n\t" {
        case "":
            v15 = v15
        case "hello":
            fmt.Println(v12)
        default:
            var v21 int32 = 42
            _ = v21
        }
        for fi22 := 0; fi22 < 9; fi22++ {
            if v12 {
                fmt.Println(v12)
                var v23 uint8 = v15
                delete(v3, (fi22 >> ((2) & 63)))
                _ = v23
            } else {
                if true {
                    v3[(fi22 ^ fi22)] = 1
                }
                if v12 {
                    v10 += v10
                    fmt.Println(fi22, v1)
                    v4 += v4
                } else {
                    if v12 {
                        if (((float64(((-2.25 * float32(1e10)) + -2.25)) + (1e10 - 0.1)) / 1.0) > 0.0) {
                            for fi24 := 0; fi24 < 3; fi24++ {
                                switch "\n\t" {
                                case "\n\t":
                                    var v25 uint64 = (v2 & v2)
                                    _ = v25
                                case "":
                                    v26 := GenS3{F0: 65535}
                                    _ = v26
                                default:
                                    if v12 {
                                        v2 -= (v2 - (v2 / (((v2 << (((v1 - (v1 >> ((42) & 63)))) & 63))) | 1)))
                                    }
                                }
                                if v12 {
                                    var v27 string = ((((("\n\t" + "hello") + "a") + "a") + "hello") + "hello")
                                    switch ("\n\t" + "\n\t") {
                                    case "\n\t":
                                        var v28 int16 = 42
                                        _ = v28
                                    case "hello":
                                        _ = len(v3)
                                    case "":
                                        if (!v12) {
                                            fmt.Println(fi24)
                                        }
                                        delete(v3, fi24)
                                        v10 -= -887
                                        v3[fi22] = fi24
                                        _ = len(v3)
                                    default:
                                        v2 = 0
                                        v29 := v3[-1]
                                        _ = v29
                                    }
                                    _ = v27
                                }
                                _ = fi24
                            }
                            v30 := GenS1{F0: fi22}
                            _ = v30
                        }
                    } else {
                        for fi31 := 0; fi31 < 20; fi31++ {
                            if (!(!(((!false) && (0.1 > (((3.14159 * -2.25) - (0.1 - 1.5)) * (1.5 - 3.14159)))) || (!false)))) {
                                v1 += v1
                            } else {
                                var v32 bool = (false && (true && (423 != (v2 - v2))))
                                v33 := GenS2{F0: ((!((v10 == (v10 << (((v1 >> ((v1) & 63))) & 63))) && (!v32))) || v32), F1: -32768, F2: 42}
                                fmt.Println(v2)
                                _ = v32
                                _ = v33
                            }
                            if ((("\n\t" + "hello") + "foo bar") > "foo bar") {
                                switch "hello" {
                                case "a":
                                    v34 := GenS1{F0: (fi22 ^ fi31)}
                                    _ = v34
                                case "hello":
                                    if (v17 == (-(v17))) {
                                        v10 = (v10 - 0)
                                        if (v12 || (29 > 108)) {
                                            var v35 int8 = 25
                                            fmt.Println(v17, v15, v15)
                                            if (!false) {
                                                var v36 float64 = float64((v4 - v4))
                                                v9 = append(v9, 127)
                                                _ = v36
                                            } else {
                                                v17 += v17
                                                fmt.Println(v2, v10, v2)
                                                v12 = (-1 < v17)
                                            }
                                            _ = v35
                                        } else {
                                            v2 = (v2 ^ (^(v2)))
                                            if (v10 <= (^(v10))) {
                                                var v37 uint64 = (v2 + (v2 - v2))
                                                if false {
                                                    var v38 int64 = 233
                                                    if (("\n\t" + "") == "hello") {
                                                        fmt.Println(v38, v37, v12)
                                                    } else {
                                                        var v39 uint64 = v37
                                                        _ = v39
                                                    }
                                                    _ = v38
                                                } else {
                                                    fmt.Println(v1)
                                                    if (-32768 != -1) {
                                                        var v40 uint = (v1 / ((338) | 1))
                                                        _ = v40
                                                    }
                                                }
                                                if (fi31 >= -9223372036854775808) {
                                                    delete(v3, fi22)
                                                    v12 = false
                                                } else {
                                                    var v41 uint = v1
                                                    if (true && v12) {
                                                        v42 := GenS3{F0: uint16(v1)}
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v15 = v15
                                                            switch "foo bar" {
                                                            case "":
                                                                var v43 bool = v12
                                                                _ = v43
                                                            case "hello":
                                                                v1 -= 18446744073709551615
                                                            }
                                                            fmt.Println(v17, v1, v17)
                                                        case "foo bar":
                                                            var v44 int16 = -635
                                                            _ = v44
                                                        case "a":
                                                            v17 = (v17 - (v17 - v17))
                                                        }
                                                        _ = v42
                                                    }
                                                    fmt.Println(v12, v10, fi31)
                                                    var v45 uint64 = v37
                                                    _ = v41
                                                    _ = v45
                                                }
                                                _ = v37
                                            }
                                        }
                                        v46 := v3[fi22]
                                        var v47 uint16 = 0
                                        _ = v46
                                        _ = v47
                                    }
                                case "":
                                    switch "\n\t" {
                                    case "":
                                        var v48 uint32 = 197
                                        var v49 uint64 = (^(v2))
                                        _ = v48
                                        _ = v49
                                    default:
                                        if ((!v12) || (2 != 42)) {
                                            if (int64(v1) == v10) {
                                                if v12 {
                                                    v1 += (v1 / ((v1) | 1))
                                                    var v50 uint8 = v15
                                                    if true {
                                                        switch "a" {
                                                        case "":
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                var v51 int32 = 0
                                                                _ = v51
                                                            }
                                                        case "foo bar":
                                                            if v12 {
                                                                v1 = v1
                                                                var v52 uint32 = 438
                                                                _ = v52
                                                            } else {
                                                                if ((v1 * v1) <= 2) {
                                                                    v53 := GenS3{F0: 255}
                                                                    if (!(uint16(v15) < 42)) {
                                                                        var v54 int64 = 0
                                                                        fmt.Println(v4)
                                                                        fmt.Println(v50)
                                                                        v3[(^(fi22))] = (fi31 >> ((255) & 63))
                                                                        _ = v54
                                                                    }
                                                                    _ = v53
                                                                }
                                                            }
                                                        default:
                                                            var v55 float32 = -2.25
                                                            _ = v55
                                                        }
                                                        var v56 string = ("\n\t" + "foo bar")
                                                        _ = len(v3)
                                                        _ = v56
                                                    }
                                                    _ = v50
                                                }
                                            }
                                            if v12 {
                                                if (!v12) {
                                                    fmt.Println(v1, v12)
                                                    if (((float64(1.5) * 1e10) - -2.25) <= (1e10 - 0.0)) {
                                                        switch "a" {
                                                        case "\n\t":
                                                            fmt.Println(v4)
                                                        case "a":
                                                            fmt.Println(v17)
                                                        default:
                                                            v4 = v4
                                                        }
                                                        fmt.Println(fi31)
                                                        var v57 uint8 = (v15 % ((v15) | 1))
                                                        v9 = append(v9, fi22)
                                                        v58 := GenS3{F0: 1}
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        _ = len(v3)
                                                        v17 -= (v17 + v17)
                                                    }
                                                } else {
                                                    fmt.Println(fi22)
                                                    fmt.Println(v1, v10, fi31)
                                                    var v59 int64 = v10
                                                    v59 += v10
                                                    switch "hello" {
                                                    case "hello":
                                                        var v60 bool = v12
                                                        _ = v60
                                                    case "":
                                                        var v61 bool = (3.14159 != 0.1)
                                                        _ = v61
                                                    case "\n\t":
                                                        var v62 int32 = (v17 & v17)
                                                        _ = v62
                                                    }
                                                    _ = v59
                                                }
                                                var v63 uint64 = 662
                                                _ = v63
                                            }
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v64 int64 = (v10 - v10)
                                                _ = v64
                                            default:
                                                var v65 bool = (1000 == 32767)
                                                var v66 uint16 = 32767
                                                _ = v65
                                                _ = v66
                                            }
                                        }
                                    }
                                default:
                                    v4 = v4
                                }
                            } else {
                                v2 -= 1
                            }
                            if v12 {
                                switch "hello" {
                                case "a":
                                    switch "\n\t" {
                                    case "":
                                        var v67 int8 = 100
                                        _ = v67
                                    default:
                                        switch ("hello" + (((("hello" + "hello") + "foo bar") + "hello") + ("\n\t" + "foo bar"))) {
                                        case "a":
                                            var v68 int8 = 1
                                            _ = v68
                                        case "":
                                            var v69 float64 = 0.1
                                            _ = v69
                                        case "hello":
                                            if v12 {
                                                var v70 uint = 2
                                                _ = v70
                                            } else {
                                                var v71 int8 = -101
                                                _ = v71
                                            }
                                        default:
                                            v72 := v3[fi31]
                                            _ = v72
                                        }
                                    }
                                case "hello":
                                    v73 := GenS1{F0: (^(fi22))}
                                    _ = v73
                                default:
                                    if v12 {
                                        v12 = ((3.14159 * 0.0) >= 1.5)
                                        var v74 int32 = v17
                                        _ = v74
                                    }
                                }
                                v75 := GenS1{F0: fi22}
                                fmt.Println(v17)
                                switch "" {
                                case "\n\t":
                                    v3[fi22] = (fi31 ^ int(v15))
                                }
                                _ = v75
                            } else {
                                if (-9223372036854775808 > v10) {
                                    var v76 uint = (^(v1))
                                    _ = len(v3)
                                    _ = v76
                                }
                                var v77 int = fi22
                                delete(v3, -846)
                                if (42 < 4294967295) {
                                    if len(v9) >= 2 {
                                        v9 = v9[1:len(v9)]
                                    }
                                } else {
                                    if (v77 <= 1) {
                                        v9 = append(v9, (v77 ^ v77))
                                        var v78 bool = v12
                                        _ = v78
                                    } else {
                                        v79 := GenS1{F0: (v77 / (((v77 - 367)) | 1))}
                                        _ = v79
                                    }
                                }
                                v9 = append(v9, fi31)
                                v80 := GenS2{F0: v12, F1: 32767, F2: (v17 ^ v17)}
                                v10 += int64(v1)
                                _ = v77
                                _ = v80
                            }
                            v81 := GenS4{F0: 2147483648}
                            _ = fi31
                            _ = v81
                        }
                        if true {
                            if v12 {
                                for fi82 := 0; fi82 < 18; fi82++ {
                                    v83 := GenS5{F0: "hello", F1: ("a" + "hello")}
                                    _ = fi82
                                    _ = v83
                                }
                                v15 += uint8(v1)
                                if v12 {
                                    v2 = (^(v2))
                                    var v84 bool = (((-2.25 / -2.25) / 1.5) <= (1e10 * ((0.1 + (((float64(3.14159) - 1e10) / (0.1 / 1.5)) / 1e10)) / 1.0)))
                                    _ = v84
                                }
                            } else {
                                v3[9223372036854775807] = (fi22 & -633)
                                if len(v9) > 0 {
                                    v9[0] = fi22
                                }
                            }
                            if false {
                                var v85 uint = v1
                                var v86 uint32 = 2147483648
                                _ = v85
                                _ = v86
                            }
                        }
                        fmt.Println(v12, v15, v1)
                    }
                    if true {
                        var v87 uint32 = 2147483647
                        var v88 int8 = 1
                        _ = v87
                        _ = v88
                    } else {
                        var v89 int8 = 72
                        if true {
                            v89 = (v89 >> ((v1) & 63))
                        }
                        _ = v89
                    }
                    var v90 int32 = (v17 % ((v17) | 1))
                    _ = v90
                }
            }
            var v91 uint64 = v2
            _ = fi22
            _ = v91
        }
        _ = v20
    }
    var v92 uint8 = (v15 - v15)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
    _ = v92
}
