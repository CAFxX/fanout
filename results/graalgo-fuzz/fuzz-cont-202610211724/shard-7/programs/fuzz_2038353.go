package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int32 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 uint8; F2 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 int16; F1 int32; F2 int64 }
type GenS9 struct { F0 int16; F1 uint8 }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint32; F1 bool }

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]int)
    var v3 int64 = 42
    v3 = 1
    var v4 int32 = -2147483648
    v5 := []int{9223372036854775807}
    var v6 float64 = (1.5 - float64(0.0))
    var v7 string = (v1 + v1)
    for fi8 := 0; fi8 < 8; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            switch (v7 + v1) {
            case "\n\t":
                fmt.Println(v3, v1)
            case "":
                if len(v5) > 0 {
                    v5[0] = fi9
                }
            default:
                switch "" {
                case "foo bar":
                    v1 = ("\n\t" + (("\n\t" + v7) + v7))
                case "\n\t":
                    fmt.Println(fi9, v1)
                case "":
                    if (v3 >= -1) {
                        var v10 float64 = v6
                        v1 = ((v7 + v7) + v1)
                        if true {
                            switch v7 {
                            case "a":
                                _ = len(v5)
                            }
                        }
                        _ = v10
                    } else {
                        v11 := GenS1{F0: ((!true) || false), F1: v4, F2: (v4 + v4)}
                        if (!(0 >= (v3 + 1))) {
                            fmt.Println(v1, v7, fi8)
                            delete(v2, (-(fi8)))
                            v4 = v4
                        } else {
                            v5 = append(v5, (fi8 - fi9))
                            var v12 int = 32768
                            v6 = 1e10
                            fmt.Println(fi8, v6, v12)
                            _ = v12
                        }
                        switch v7 {
                        case "\n\t":
                            v2[fi8] = (fi8 + (fi9 & fi9))
                            var v13 int16 = 32767
                            _ = v13
                        case "":
                            if (!((26 < 0) || (128 < 108))) {
                                switch (fi8 + (fi9 + (fi8 << ((0) & 63)))) {
                                case 160:
                                    v14 := GenS2{F0: (-(v3)), F1: 1}
                                    _ = v14
                                }
                                v4 += (-(v4))
                            } else {
                                v15 := GenS1{F0: (1e10 >= 1.5), F1: 42, F2: v4}
                                switch (fi9 % (((fi9 ^ 32767)) | 1)) {
                                case -535:
                                    v4 -= (v4 + v4)
                                    fmt.Println(v7)
                                case -9223372036854775808:
                                    switch -332 {
                                    case -1:
                                        v16 := GenS1{F0: (1 > 127), F1: -1, F2: v4}
                                        _ = v16
                                    case 369:
                                        var v17 uint8 = 1
                                        _ = v17
                                    case 1:
                                        if ("\n\t" <= v1) {
                                            fmt.Println(v1)
                                        } else {
                                            if true {
                                                v3 = (v3 << ((18446744073709551615) & 63))
                                                v4 = 576
                                                fmt.Println(v7, v3, fi8)
                                            }
                                            var v18 int64 = (v3 ^ (v3 & v3))
                                            if true {
                                                v19 := GenS3{F0: v18}
                                                _ = len(v2)
                                                _ = v19
                                            }
                                            _ = v18
                                        }
                                    default:
                                        var v20 uint8 = 42
                                        _ = v20
                                    }
                                case 128:
                                    if len(v5) > 0 {
                                        v5[0] = (fi8 / ((fi8) | 1))
                                    }
                                default:
                                    fmt.Println(v4)
                                }
                                if false {
                                    if false {
                                        v1 = v1
                                        switch (((((v7 + (v1 + "\n\t")) + (v7 + v7)) + "\n\t") + ((v7 + v7) + ("a" + v1))) + v7) {
                                        case "\n\t":
                                            if (v1 <= (v1 + v1)) {
                                                v2[(-(fi8))] = (-(fi9))
                                            }
                                        default:
                                            v21 := GenS4{F0: 0}
                                            _ = v21
                                        }
                                        v3 = v3
                                    }
                                    var v22 int = 0
                                    fmt.Println(fi9, v6, v6)
                                    fmt.Println(fi8, fi9, v22)
                                    _ = v22
                                } else {
                                    v3 = v3
                                }
                                _ = v15
                            }
                        case "hello":
                            if true {
                                v2[fi9] = 84
                                var v23 float64 = -2.25
                                _ = v23
                            }
                            fmt.Println(v7)
                        default:
                            if false {
                                switch fi8 {
                                case 0:
                                    if false {
                                        v7 = (((v1 + v1) + v1) + v1)
                                    }
                                case 665:
                                    v24 := v2[0]
                                    _ = v24
                                }
                                if true {
                                    var v25 int32 = -615
                                    if true {
                                        v2[(fi8 & (fi9 % ((fi8) | 1)))] = fi9
                                        switch "\n\t" {
                                        case "\n\t":
                                            v7 = ""
                                        case "":
                                            var v26 uint = 42
                                            _ = v26
                                        case "a":
                                            if (v3 < (v3 & (v3 | 594))) {
                                                v27 := GenS3{F0: (v3 % (((v3 * v3)) | 1))}
                                                fmt.Println(v3, fi9, v3)
                                                _ = v27
                                            } else {
                                                if false {
                                                    v4 = 42
                                                }
                                            }
                                        default:
                                            var v28 uint32 = 0
                                            _ = v28
                                        }
                                        v29 := GenS3{F0: v3}
                                        _ = v29
                                    }
                                    v4 += -821
                                    _ = v25
                                } else {
                                    switch (fi9 - fi8) {
                                    case -156:
                                        if (!(618 >= 2)) {
                                            if true {
                                                if (-32768 > 852) {
                                                    if true {
                                                        var v30 int16 = -1
                                                        _ = len(v5)
                                                        fmt.Println(fi9)
                                                        var v31 float64 = v6
                                                        v32 := GenS5{F0: -1, F1: 1, F2: 32767}
                                                        _ = v30
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                }
                                                if true {
                                                    var v33 bool = false
                                                    if ((((v7 + v1) + "") + "\n\t") >= v7) {
                                                        if len(v5) > 0 {
                                                            v5[0] = (fi9 | 9223372036854775807)
                                                        }
                                                        if v33 {
                                                            delete(v2, (fi9 / ((-2147483648) | 1)))
                                                            v6 += (v6 / v6)
                                                        }
                                                        fmt.Println(v3)
                                                    } else {
                                                        _ = len(v5)
                                                        fmt.Println(v7)
                                                        fmt.Println(fi9, fi8, v3)
                                                        if (false && (int16(v3) != -32768)) {
                                                            var v34 float32 = -2.25
                                                            _ = v34
                                                        }
                                                    }
                                                    _ = v33
                                                }
                                                v4 -= 2147483647
                                            }
                                            v7 = ""
                                            switch v1 {
                                            case "foo bar":
                                                v4 = (v4 - (v4 ^ (v4 % (((v4 ^ (v4 >> ((652) & 63)))) | 1))))
                                            case "a":
                                                _ = len(v5)
                                            }
                                            v4 = v4
                                        } else {
                                            var v35 uint8 = 0
                                            v6 = v6
                                            _ = v35
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v7, v3, fi9)
                                    case -9223372036854775808:
                                        v36 := GenS6{F0: -1}
                                        _ = v36
                                    }
                                    if (v6 != -2.25) {
                                        if (0 <= 4294967295) {
                                            var v37 string = (("\n\t" + "hello") + v7)
                                            v6 += (v6 / v6)
                                            _ = v37
                                        }
                                        v38 := GenS7{F0: 1, F1: (-1 < 42)}
                                        if (18446744073709551615 != 0) {
                                            if ((!(1 != 65535)) && (1e10 > (1.5 / (0.0 + 3.14159)))) {
                                                var v39 uint16 = 2
                                                _ = v39
                                            }
                                            fmt.Println(v7)
                                            switch -2147483648 {
                                            case 156:
                                                v40 := GenS8{F0: 1, F1: v4, F2: v3}
                                                _ = v40
                                            case -9223372036854775808:
                                                var v41 int32 = 639
                                                var v42 uint8 = 34
                                                fmt.Println(fi8, fi8, v42)
                                                _ = v41
                                                _ = v42
                                            default:
                                                v43 := GenS7{F0: 18446744073709551615, F1: (!(0 <= 18446744073709551615))}
                                                _ = v43
                                            }
                                            var v44 uint16 = 65535
                                            if (true && (!true)) {
                                                switch (fi9 & 858) {
                                                case 127:
                                                    if (637 == v44) {
                                                        v1 = "a"
                                                        switch (-(fi8)) {
                                                        case -136:
                                                            v45 := v2[(fi9 >> ((4294967295) & 63))]
                                                            v46 := GenS4{F0: 679}
                                                            _ = v45
                                                            _ = v46
                                                        default:
                                                            if false {
                                                                v1 = v7
                                                            }
                                                        }
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                default:
                                                    var v47 uint32 = 434
                                                    _ = v47
                                                }
                                            } else {
                                                _ = len(v2)
                                            }
                                            _ = v44
                                        } else {
                                            var v48 float32 = (3.14159 - 3.14159)
                                            _ = v48
                                        }
                                        _ = v38
                                    } else {
                                        fmt.Println(v4, v6, v1)
                                    }
                                    if false {
                                        var v49 uint = 4294967295
                                        v50 := GenS9{F0: -1, F1: 86}
                                        _ = len(v2)
                                        if true {
                                            var v51 string = "hello"
                                            _ = v51
                                        }
                                        _ = v49
                                        _ = v50
                                    }
                                }
                                v2[fi9] = 1
                                fmt.Println(fi9, v4, v6)
                            } else {
                                v52 := GenS10{F0: 1, F1: v4}
                                _ = v52
                            }
                        }
                        fmt.Println(fi8, v4)
                        _ = v11
                    }
                }
            }
            _ = fi9
        }
        v53 := GenS9{F0: 1000, F1: 2}
        _ = fi8
        _ = v53
    }
    if false {
        v54 := GenS11{F0: v1}
        var v55 uint16 = 32767
        for fi56 := 0; fi56 < 10; fi56++ {
            for fi57 := 0; fi57 < 4; fi57++ {
                v7 = (("\n\t" + v7) + ("a" + ""))
                var v58 int64 = v3
                _ = fi57
                _ = v58
            }
            v59 := GenS3{F0: v3}
            if false {
                var v60 string = v1
                for fi61 := 0; fi61 < 11; fi61++ {
                    v62 := GenS12{F0: 127}
                    v63 := GenS13{F0: 255}
                    _ = fi61
                    _ = v62
                    _ = v63
                }
                var v64 float64 = v6
                _ = v60
                _ = v64
            } else {
                for fi65 := 0; fi65 < 15; fi65++ {
                    v55 -= v55
                    _ = fi65
                }
                var v66 int16 = 0
                v5 = append(v5, (fi56 % ((-894) | 1)))
                _ = v66
            }
            _ = fi56
            _ = v59
        }
        _ = v54
        _ = v55
    }
    if true {
        delete(v2, -9223372036854775808)
        switch v7 {
        case "foo bar":
            v67 := GenS14{F0: 2147483648, F1: true}
            _ = v67
        default:
            var v68 uint64 = 850
            _ = v68
        }
        v4 = v4
    } else {
        v6 += 0.0
        v7 = "hello"
    }
    v4 = (v4 | 0)
    fmt.Println(v6, v1)
    var v69 bool = (2 == 502)
    var v70 uint = 0
    if v69 {
        var v71 int = -2147483648
        v7 = v1
        _ = v71
    } else {
        var v72 int16 = 42
        v3 -= 0
        _ = v72
    }
    v7 = (v7 + "")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v7:", v7)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v7
    _ = v70
}
