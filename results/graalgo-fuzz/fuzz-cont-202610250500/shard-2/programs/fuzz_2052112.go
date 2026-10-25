package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint16; F2 bool }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 bool }
type GenS8 struct { F0 uint8; F1 int16 }
func genFn0(p0 string, p1 uint64, p2 uint64) uint16 {
    v6 := []int{128, -502}
    p1 = (p2 % (((p1 - (p2 >> ((457) & 63)))) | 1))
    if (true && false) {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 uint8 = 42
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS1{F0: (^(p2))}
                v11 := GenS1{F0: (p2 / (((^(p1))) | 1))}
                v12 := GenS2{F0: -2147483648, F1: 42}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi7
            _ = v8
        }
    }
    if false {
        if true {
            v13 := GenS1{F0: 626}
            _ = v13
        } else {
            if true {
                fmt.Println(p0, p2)
                var v14 uint8 = 2
                var v15 int64 = -918
                _ = v14
                _ = v15
            }
            fmt.Println(p1, p2)
        }
        switch -738 {
        case 32767:
            p2 += p1
        case -9223372036854775808:
            p2 += (p2 ^ 846)
        }
    } else {
        switch -2147483648 {
        case 128:
            fmt.Println(p0)
        case 32768:
            _ = len(v6)
        case 127:
            fmt.Println(p2, p0, p0)
        default:
            for fi16 := 0; fi16 < 5; fi16++ {
                v17 := GenS3{F0: 691, F1: 65535, F2: true}
                p0 = "foo bar"
                v18 := GenS1{F0: p2}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        p0 = p0
        fmt.Println(p1, p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 583
}
func genFn1() bool {
    v19 := GenS1{F0: 328}
    var v20 int8 = 127
    var v21 string = "a"
    for fi22 := 0; fi22 < 19; fi22++ {
        v23 := []int{}
        if false {
            v24 := GenS3{F0: 65535, F1: 255, F2: true}
            switch (fi22 % ((fi22) | 1)) {
            case 128:
                v20 = (v20 + 47)
            case 127:
                v21 = ("a" + v21)
            default:
                if (!true) {
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    fmt.Println(v20, fi22, v21)
                }
            }
            var v25 float64 = (0.1 / ((3.14159 * (1.5 * 0.1)) - 3.14159))
            _ = v24
            _ = v25
        } else {
            v21 = v21
        }
        _ = fi22
        _ = v23
    }
    if (-199 >= 508) {
        var v26 float32 = 3.14159
        for fi27 := 0; fi27 < 2; fi27++ {
            if (-975 <= 42) {
                v26 = ((((((v26 * v26) - (v26 - ((0.0 - v26) - v26))) + 3.14159) * 3.14159) - (v26 * v26)) + (v26 * (v26 - float32(1e10))))
                fmt.Println(v20)
            } else {
                if false {
                    for fi28 := 0; fi28 < 18; fi28++ {
                        v29 := GenS4{F0: 42, F1: 2}
                        if false {
                            v30 := GenS5{F0: (1 <= 1), F1: 32767}
                            var v31 uint8 = 255
                            _ = v30
                            _ = v31
                        } else {
                            var v32 float64 = 0.0
                            var v33 uint64 = 1
                            v21 = ("" + "foo bar")
                            _ = v32
                            _ = v33
                        }
                        switch "hello" {
                        case "a":
                            var v34 string = ""
                            _ = v34
                        case "hello":
                            switch "\n\t" {
                            case "":
                                var v35 string = "a"
                                _ = v35
                            }
                        case "":
                            var v36 uint8 = 189
                            _ = v36
                        default:
                            v37 := []int{fi28, fi27, int(v20)}
                            _ = v37
                        }
                        _ = fi28
                        _ = v29
                    }
                }
            }
            _ = fi27
        }
        _ = v26
    }
    _ = v19
    _ = v20
    _ = v21
    return (!false)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 42
    var v3 uint = 255
    var v4 string = ("\n\t" + (("hello" + "hello") + "a"))
    var v5 int16 = 1000
    var v38 uint64 = 18446744073709551615
    var v39 int64 = 42
    fmt.Println(v39, v3)
    v40 := GenS6{F0: v4}
    if true {
        v3 -= (v3 % (((^(v3))) | 1))
    }
    var v41 int8 = -91
    fmt.Println(v4)
    var v42 uint = v3
    v42 = v42
    var v43 int8 = (v41 / ((int8(v5)) | 1))
    switch "" {
    case "\n\t":
        if (v38 > (v38 + 2)) {
            v5 = v5
        } else {
            v44 := []int{32767, 42, -364}
            if (false || true) {
                var v45 int = 9223372036854775807
                _ = v45
            } else {
                if len(v44) > 0 {
                    v44[0] = 9223372036854775807
                }
                var v46 int16 = (v5 & (v5 % ((42) | 1)))
                _ = v46
            }
            _ = v44
        }
    default:
        for fi47 := 0; fi47 < 14; fi47++ {
            v1[2] = (v4 + "")
            _ = fi47
        }
    }
    v2 -= v41
    if false {
        v39 = v39
        fmt.Println(v41)
    }
    fmt.Println(v43, v42, v2)
    v48 := v1[2147483648]
    v2 = (v2 | 127)
    if true {
        switch 128 {
        case -169:
            switch 2147483647 {
            case 2147483647:
                var v49 bool = true
                _ = v49
            case -923:
                v39 = (v39 & (v39 + (-(v39))))
            default:
                if ((true || false) && false) {
                    v50 := []int{-40, 724}
                    var v51 int32 = 2147483647
                    _ = v50
                    _ = v51
                } else {
                    v5 += int16(v42)
                    v3 = 609
                    var v52 bool = ((-(v5)) == (v5 ^ v5))
                    _ = v52
                }
            }
        case -9223372036854775808:
            var v53 int = -2147483648
            _ = v53
        default:
            v54 := GenS5{F0: ((v2 << (((v3 / (((v3 + v42)) | 1))) & 63)) < (v2 % (((-(v43))) | 1))), F1: (v5 / ((v5) | 1))}
            _ = v54
        }
        v55 := []int{128, 1}
        var v56 uint = (^(v3))
        _ = v55
        _ = v56
    }
    if true {
        if false {
            v57 := GenS5{F0: false, F1: v5}
            v58 := []int{834, int(v5), int(v38)}
            var v59 float32 = (-2.25 / (-2.25 + 0.1))
            _ = v57
            _ = v58
            _ = v59
        } else {
            v41 = 100
            switch "hello" {
            case "a":
                var v60 uint64 = v38
                _ = v60
            case "hello":
                var v61 int8 = v43
                _ = v61
            }
            switch 0 {
            case 2147483647:
                switch ("foo bar" + v48) {
                case "":
                    if (v4 < "a") {
                        v3 = 2
                        v2 -= int8(v38)
                    } else {
                        if (!(int32(v3) <= -2147483648)) {
                            delete(v1, 115)
                            v42 += uint(v38)
                            v48 = (v4 + ("hello" + ""))
                        } else {
                            v62 := []int{1, -9223372036854775808, -1, 734}
                            switch (v48 + "") {
                            case "hello":
                                fmt.Println(v5, v38)
                            case "\n\t":
                                for fi63 := 0; fi63 < 19; fi63++ {
                                    v39 = -281
                                    _ = fi63
                                }
                            default:
                                fmt.Println(v42, v4, v38)
                            }
                            v42 = v42
                            _ = v62
                        }
                        v1[962] = ("foo bar" + v48)
                    }
                case "\n\t":
                    switch "a" {
                    case "\n\t":
                        if (4294967295 <= 304) {
                            delete(v1, 76)
                            var v64 uint16 = 982
                            var v65 int = 2147483647
                            _ = v64
                            _ = v65
                        }
                    case "foo bar":
                        if false {
                            for fi66 := 0; fi66 < 3; fi66++ {
                                switch (^(fi66)) {
                                case 42:
                                    for fi67 := 0; fi67 < 11; fi67++ {
                                        v38 += 42
                                        _ = fi67
                                    }
                                case 2147483647:
                                    v2 -= (v43 / ((-1) | 1))
                                case 0:
                                    if true {
                                        switch "\n\t" {
                                        case "foo bar":
                                            v48 = ((((v4 + "a") + v4) + v48) + v48)
                                        case "\n\t":
                                            for fi68 := 0; fi68 < 3; fi68++ {
                                                v69 := GenS5{F0: false, F1: v5}
                                                v70 := GenS1{F0: (v38 % ((v38) | 1))}
                                                v71 := GenS5{F0: true, F1: (-(v5))}
                                                _ = fi68
                                                _ = v69
                                                _ = v70
                                                _ = v71
                                            }
                                        default:
                                            if (v39 >= -245) {
                                                v43 = v2
                                                var v72 uint16 = 649
                                                var v73 int = (fi66 >> ((v3) & 63))
                                                _ = v72
                                                _ = v73
                                            } else {
                                                v74 := GenS4{F0: v38, F1: 42}
                                                if (v2 != (v43 % (((v43 << (((v3 >> ((1) & 63))) & 63))) | 1))) {
                                                    fmt.Println(v48)
                                                    v41 -= 42
                                                } else {
                                                    v75 := []int{(fi66 % (((fi66 ^ -9223372036854775808)) | 1)), fi66}
                                                    v4 = "\n\t"
                                                    _ = v75
                                                }
                                                v76 := GenS7{F0: 9223372036854775807, F1: true}
                                                _ = v74
                                                _ = v76
                                            }
                                        }
                                        fmt.Println(fi66, v41, fi66)
                                    } else {
                                        v48 = (v48 + v4)
                                    }
                                default:
                                    if true {
                                        if (-1 < (v39 % ((95) | 1))) {
                                            v77 := []int{(fi66 + fi66), (-(fi66)), 9223372036854775807, (fi66 ^ (fi66 ^ (fi66 / ((fi66) | 1))))}
                                            for fi78 := 0; fi78 < 7; fi78++ {
                                                v5 = v5
                                                var v79 int = (fi78 << ((v42) & 63))
                                                _ = fi78
                                                _ = v79
                                            }
                                            v5 = int16(v39)
                                            _ = v77
                                        }
                                    }
                                }
                                for fi80 := 0; fi80 < 12; fi80++ {
                                    fmt.Println(v42, v41)
                                    v81 := GenS4{F0: 0, F1: 2147483647}
                                    _ = fi80
                                    _ = v81
                                }
                                _ = fi66
                            }
                        } else {
                            if (!false) {
                                v82 := v1[2147483647]
                                switch v48 {
                                case "foo bar":
                                    var v83 int = 32767
                                    _ = v83
                                case "":
                                    v84 := GenS5{F0: false, F1: int16(v38)}
                                    _ = v84
                                case "a":
                                    for fi85 := 0; fi85 < 9; fi85++ {
                                        fmt.Println(v82, v5, v4)
                                        v39 = v39
                                        v86 := GenS7{F0: v39, F1: false}
                                        _ = fi85
                                        _ = v86
                                    }
                                }
                                if (v42 > v42) {
                                    for fi87 := 0; fi87 < 16; fi87++ {
                                        if false {
                                            fmt.Println(v2, v2)
                                            v4 = "hello"
                                            delete(v1, 583)
                                        }
                                        var v88 int16 = 0
                                        _ = fi87
                                        _ = v88
                                    }
                                }
                                _ = v82
                            } else {
                                fmt.Println(v42)
                                for fi89 := 0; fi89 < 20; fi89++ {
                                    v90 := GenS4{F0: v38, F1: 65535}
                                    v91 := GenS8{F0: 0, F1: (v5 ^ -493)}
                                    _ = fi89
                                    _ = v90
                                    _ = v91
                                }
                                v92 := []int{}
                                _ = v92
                            }
                        }
                    case "":
                        switch ("foo bar" + v48) {
                        case "a":
                            switch (("\n\t" + ((v4 + v48) + v48)) + (("" + "a") + v48)) {
                            case "foo bar":
                                v1[0] = (v4 + v4)
                            case "":
                                v5 -= (^(v5))
                            }
                        default:
                            var v93 float64 = 0.0
                            _ = v93
                        }
                    default:
                        v94 := GenS8{F0: 73, F1: v5}
                        _ = v94
                    }
                case "a":
                    v95 := []int{}
                    _ = v95
                }
            default:
                v96 := v1[42]
                _ = v96
            }
        }
        switch "a" {
        case "":
            v3 = v42
        case "hello":
            for fi97 := 0; fi97 < 11; fi97++ {
                var v98 int8 = -33
                var v99 int16 = (v5 + (v5 / ((v5) | 1)))
                _ = fi97
                _ = v98
                _ = v99
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v48
    _ = v5
}
