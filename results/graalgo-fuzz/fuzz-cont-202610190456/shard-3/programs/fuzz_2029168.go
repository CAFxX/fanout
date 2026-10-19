package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int64; F2 int16 }
type GenS3 struct { F0 int8; F1 string; F2 int32 }
type GenS4 struct { F0 string; F1 int; F2 uint8 }
type GenS5 struct { F0 uint; F1 uint; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int16; F1 int8; F2 int8 }
type GenS8 struct { F0 uint64; F1 int }

func main() {
    v1 := []int{829, 2147483647}
    var v2 uint32 = 2
    v3 := make(map[int]int)
    var v4 bool = (1.5 > 0.1)
    v4 = ((^(v2)) <= 892)
    fmt.Println(v2, v2)
    if v4 {
        for fi5 := 0; fi5 < 12; fi5++ {
            switch fi5 {
            case 42:
                v6 := GenS1{F0: v2}
                _ = v6
            default:
                if len(v1) > 0 {
                    v1[0] = fi5
                }
            }
            _ = fi5
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            if true {
                switch "" {
                case "a":
                    var v8 uint16 = 545
                    _ = v8
                }
                v4 = false
            } else {
                if true {
                    switch (fi7 >> ((2) & 63)) {
                    case 127:
                        fmt.Println(v2, v2, fi7)
                    case 9223372036854775807:
                        for fi9 := 0; fi9 < 16; fi9++ {
                            var v10 float64 = 1.5
                            var v11 int64 = 530
                            fmt.Println(v2, fi7)
                            _ = fi9
                            _ = v10
                            _ = v11
                        }
                    case 1:
                        v4 = ((-361 < 42) || (v4 && v4))
                    default:
                        fmt.Println(v4, v4, v4)
                    }
                    switch "\n\t" {
                    case "hello":
                        if (15 < 255) {
                            for fi12 := 0; fi12 < 3; fi12++ {
                                fmt.Println(fi7)
                                _ = fi12
                            }
                            v1 = append(v1, int(v2))
                        } else {
                            switch "a" {
                            case "a":
                                v13 := GenS1{F0: 42}
                                _ = v13
                            case "":
                                v14 := GenS2{F0: 42, F1: -9223372036854775808, F2: -1}
                                _ = v14
                            case "foo bar":
                                if v4 {
                                    if (!(127 >= 1)) {
                                        if (42 == -456) {
                                            var v15 bool = (-512 <= 1)
                                            _ = v15
                                        } else {
                                            fmt.Println(v2, v4)
                                            v2 = (v2 ^ (v2 & v2))
                                            v1 = append(v1, fi7)
                                        }
                                        v16 := GenS3{F0: 1, F1: "hello", F2: -237}
                                        var v17 int64 = int64(v2)
                                        _ = v16
                                        _ = v17
                                    } else {
                                        var v18 int8 = 34
                                        for fi19 := 0; fi19 < 5; fi19++ {
                                            v2 -= v2
                                            _ = fi19
                                        }
                                        v20 := GenS3{F0: v18, F1: "\n\t", F2: 278}
                                        _ = v18
                                        _ = v20
                                    }
                                    if v4 {
                                        v4 = (4294967295 > uint(v2))
                                        for fi21 := 0; fi21 < 9; fi21++ {
                                            v2 = 0
                                            v2 += (v2 / ((v2) | 1))
                                            v4 = ((3.14159 - 1.5) >= 1e10)
                                            _ = fi21
                                        }
                                        v22 := GenS1{F0: (^(v2))}
                                        _ = v22
                                    }
                                    if v4 {
                                        if (-1 >= 2147483647) {
                                            fmt.Println(v2, fi7, fi7)
                                        } else {
                                            var v23 uint8 = 1
                                            v3[int(v23)] = fi7
                                            var v24 int32 = 711
                                            _ = v23
                                            _ = v24
                                        }
                                        v25 := GenS1{F0: (v2 - (v2 >> ((18446744073709551615) & 63)))}
                                        if (true && (false || (2147483647 >= 2147483647))) {
                                            v26 := GenS3{F0: 100, F1: "hello", F2: 453}
                                            switch "foo bar" {
                                            case "hello":
                                                for fi27 := 0; fi27 < 11; fi27++ {
                                                    v4 = v4
                                                    if false {
                                                        v4 = (18446744073709551615 <= 2)
                                                    }
                                                    _ = fi27
                                                }
                                            case "foo bar":
                                                for fi28 := 0; fi28 < 10; fi28++ {
                                                    _ = len(v1)
                                                    _ = fi28
                                                }
                                            case "\n\t":
                                                v29 := GenS4{F0: (("foo bar" + "a") + "foo bar"), F1: (fi7 / ((int(v2)) | 1)), F2: 127}
                                                _ = v29
                                            }
                                            _ = v26
                                        } else {
                                            v4 = v4
                                        }
                                        _ = v25
                                    }
                                } else {
                                    fmt.Println(v4)
                                    v2 = 2
                                }
                            }
                        }
                    }
                    v4 = (0 != 255)
                } else {
                    var v30 int32 = 1
                    _ = len(v1)
                    _ = v30
                }
            }
            v4 = (!((3.14159 - (3.14159 + 3.14159)) >= (0.1 + (1e10 * (-2.25 - ((1e10 - -2.25) - 1.5))))))
            switch ("foo bar" + "foo bar") {
            case "":
                v1 = append(v1, (fi7 + -9223372036854775808))
            case "a":
                v1 = append(v1, (-(fi7)))
            case "foo bar":
                v31 := GenS5{F0: 42, F1: 42, F2: 2}
                _ = v31
            default:
                for fi32 := 0; fi32 < 12; fi32++ {
                    fmt.Println(fi32)
                    switch -807 {
                    case -2147483648:
                        var v33 uint64 = 18446744073709551615
                        _ = v33
                    case 939:
                        v34 := GenS3{F0: 42, F1: "foo bar", F2: 2147483647}
                        _ = v34
                    }
                    _ = fi32
                }
            }
            _ = fi7
        }
        var v35 uint32 = v2
        _ = v35
    }
    if ((false && v4) || (!(0 >= 1))) {
        if (!v4) {
            fmt.Println(v4, v4)
        } else {
            var v36 int8 = 127
            v36 = int8(v2)
            var v37 float64 = -2.25
            _ = v36
            _ = v37
        }
        v4 = (1 >= 42)
    }
    v4 = (132 == 65535)
    v2 = v2
    if (-9223372036854775808 != 1) {
        if (127 >= 1) {
            for fi38 := 0; fi38 < 3; fi38++ {
                for fi39 := 0; fi39 < 12; fi39++ {
                    fmt.Println(fi38, v2, v4)
                    var v40 uint8 = 194
                    switch (fi38 % (((fi38 + fi38)) | 1)) {
                    case 465:
                        var v41 float32 = -2.25
                        _ = v41
                    case 2147483647:
                        var v42 uint8 = v40
                        _ = v42
                    case -206:
                        var v43 int = fi38
                        _ = v43
                    }
                    v2 = 470
                    _ = fi39
                    _ = v40
                }
                for fi44 := 0; fi44 < 12; fi44++ {
                    v3[(fi44 * fi44)] = fi38
                    _ = fi44
                }
                switch 456 {
                case -1:
                    for fi45 := 0; fi45 < 9; fi45++ {
                        fmt.Println(fi38, fi38, fi45)
                        _ = fi45
                    }
                case 0:
                    if (1 != 1) {
                        switch fi38 {
                        case 922:
                            var v46 int8 = -128
                            _ = v46
                        case 1:
                            var v47 float32 = 1e10
                            _ = v47
                        }
                        var v48 int32 = 42
                        _ = v48
                    } else {
                        if (128 < uint8(v2)) {
                            v4 = false
                            if true {
                                if (227 > 128) {
                                    v4 = (1 != 0)
                                }
                                switch fi38 {
                                case -2147483648:
                                    if v4 {
                                        if ((18446744073709551615 == 65535) && false) {
                                            fmt.Println(fi38, fi38, fi38)
                                            v4 = false
                                        } else {
                                            var v49 float32 = 0.1
                                            if ((-734 >= -1) && (!v4)) {
                                                var v50 uint = 1
                                                _ = v50
                                            }
                                            _ = v49
                                        }
                                        v1 = append(v1, fi38)
                                        fmt.Println(fi38)
                                    } else {
                                        var v51 float64 = (-2.25 * (1.5 + 3.14159))
                                        v4 = (42 >= 0)
                                        for fi52 := 0; fi52 < 18; fi52++ {
                                            v53 := GenS4{F0: "foo bar", F1: (fi52 >> ((4294967295) & 63)), F2: 2}
                                            _ = fi52
                                            _ = v53
                                        }
                                        _ = v51
                                    }
                                default:
                                    switch "hello" {
                                    case "a":
                                        for fi54 := 0; fi54 < 16; fi54++ {
                                            v4 = (!v4)
                                            var v55 int32 = -2147483648
                                            if v4 {
                                                v2 -= v2
                                                fmt.Println(fi38, v55, v4)
                                                var v56 int = fi54
                                                if v4 {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    v2 = (v2 ^ 2)
                                                    v3[(v56 >> ((655) & 63))] = (v56 % ((v56) | 1))
                                                }
                                                _ = v56
                                            } else {
                                                var v57 float32 = 0.1
                                                _ = v57
                                            }
                                            _ = fi54
                                            _ = v55
                                        }
                                    case "":
                                        v58 := v3[(fi38 % ((9223372036854775807) | 1))]
                                        _ = v58
                                    default:
                                        var v59 bool = ((!((v2 & (v2 << ((18446744073709551615) & 63))) <= v2)) && (-1 != 1))
                                        _ = v59
                                    }
                                }
                            }
                        } else {
                            var v60 uint8 = 217
                            _ = v60
                        }
                    }
                default:
                    for fi61 := 0; fi61 < 10; fi61++ {
                        if len(v1) > 0 {
                            v1[0] = 32767
                        }
                        _ = fi61
                    }
                }
                _ = fi38
            }
            for fi62 := 0; fi62 < 11; fi62++ {
                for fi63 := 0; fi63 < 9; fi63++ {
                    v1 = append(v1, (fi62 * 128))
                    _ = fi63
                }
                if (1 <= 255) {
                    var v64 float32 = 0.1
                    var v65 bool = ("\n\t" != ("" + ("foo bar" + ("foo bar" + "a"))))
                    var v66 int16 = 577
                    _ = v64
                    _ = v65
                    _ = v66
                } else {
                    switch (fi62 >> ((1) & 63)) {
                    case -2147483648:
                        _ = len(v1)
                    }
                }
                _ = fi62
            }
        }
    }
    v67 := GenS6{F0: 1}
    v2 -= (v2 | 42)
    switch "\n\t" {
    case "":
        if v4 {
            delete(v3, 42)
            fmt.Println(v2)
            v2 -= (v2 << ((531) & 63))
        }
    case "hello":
        fmt.Println(v2, v2)
    case "\n\t":
        var v68 bool = (true && false)
        _ = v68
    }
    for fi69 := 0; fi69 < 20; fi69++ {
        v2 = v2
        if v4 {
            for fi70 := 0; fi70 < 17; fi70++ {
                _ = len(v1)
                var v71 bool = true
                _ = fi70
                _ = v71
            }
            fmt.Println(v4, v4, v4)
            fmt.Println(v4, v2, v2)
        } else {
            var v72 uint16 = 37
            fmt.Println(v72)
            v73 := GenS7{F0: 1000, F1: 100, F2: -115}
            _ = v72
            _ = v73
        }
        _ = fi69
    }
    if len(v1) > 0 {
        v1[0] = 32768
    }
    var v74 float32 = 3.14159
    fmt.Println(v74)
    if ((32767 > 971) && (1e10 > (0.0 * v74))) {
        var v75 uint32 = v2
        var v76 uint64 = 2
        switch 9223372036854775807 {
        case 260:
            var v77 float32 = (((float32(1e10) * -2.25) / (((v74 + v74) / float32((-2.25 + 1e10))) + v74)) / v74)
            _ = v77
        case 128:
            v78 := GenS8{F0: v76, F1: -111}
            _ = v78
        }
        _ = v75
        _ = v76
    } else {
        var v79 int64 = -760
        _ = v79
    }
    var v80 float64 = -2.25
    var v81 bool = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v74:", v74)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v74
    _ = v80
    _ = v81
}
