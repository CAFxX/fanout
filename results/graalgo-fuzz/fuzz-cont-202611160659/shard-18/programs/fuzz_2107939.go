package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint; F1 uint8; F2 int16 }

func main() {
    var v1 uint = 1
    var v2 bool = false
    v3 := []int{421, -879, 32768}
    v4 := make(map[int]int)
    var v5 int16 = 1
    var v6 uint8 = 0
    var v7 int32 = -1
    var v8 float64 = 1e10
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 string = ""
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS1{F0: 538, F1: (v6 >> (((v1 | v1)) & 63))}
            v13 := v4[(fi9 % ((fi11) | 1))]
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    v4[0] = -589
    var v14 uint64 = uint64(v1)
    v15 := GenS1{F0: (v14 >> (((v1 & v1)) & 63)), F1: v6}
    for fi16 := 0; fi16 < 14; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 int8 = 63
            _ = fi17
            _ = v18
        }
        switch fi16 {
        case 507:
            if false {
                for fi19 := 0; fi19 < 20; fi19++ {
                    v6 = (v6 | v6)
                    _ = fi19
                }
                for fi20 := 0; fi20 < 5; fi20++ {
                    var v21 float32 = 0.1
                    var v22 uint32 = 0
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            } else {
                var v23 uint16 = 32767
                var v24 float64 = (((1e10 * ((v8 / 1.5) * v8)) / 1e10) + v8)
                if (float64(1e10) >= v8) {
                    switch (fi16 % ((fi16) | 1)) {
                    case 1:
                        fmt.Println(v23, v8, v2)
                    case 42:
                        v5 += v5
                    case 127:
                        if ((v6 | (v6 + v6)) > 42) {
                            fmt.Println(v5, v23, v6)
                        } else {
                            v3 = append(v3, -298)
                        }
                    }
                    if ((v14 ^ (v14 << ((v1) & 63))) <= (^(v14))) {
                        for fi25 := 0; fi25 < 20; fi25++ {
                            switch (fi16 % ((-1) | 1)) {
                            case -9223372036854775808:
                                v1 = (v1 | (^(v1)))
                            case 42:
                                switch "a" {
                                case "":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    if true {
                                        var v26 uint32 = 42
                                        _ = v26
                                    } else {
                                        var v27 int = fi16
                                        _ = v27
                                    }
                                }
                            case 9223372036854775807:
                                switch fi16 {
                                case 0:
                                    v28 := v4[int(v23)]
                                    _ = v28
                                case 127:
                                    v29 := GenS1{F0: (v14 / (((v14 ^ v14)) | 1)), F1: (v6 & (v6 ^ (v6 >> ((209) & 63))))}
                                    _ = v29
                                }
                            }
                            _ = fi25
                        }
                    }
                    v24 += (0.0 + (3.14159 - (v24 / float64(3.14159))))
                } else {
                    if (!v2) {
                        v30 := GenS1{F0: v14, F1: uint8(v1)}
                        var v31 uint = v1
                        _ = v30
                        _ = v31
                    }
                    if true {
                        v3 = append(v3, (fi16 - (^(fi16))))
                        fmt.Println(v6, v5, fi16)
                    } else {
                        var v32 uint16 = (v23 & (v23 | 255))
                        switch 128 {
                        case 34:
                            var v33 string = "\n\t"
                            _ = v33
                        case -994:
                            for fi34 := 0; fi34 < 12; fi34++ {
                                var v35 int8 = 93
                                _ = fi34
                                _ = v35
                            }
                        case 2147483647:
                            for fi36 := 0; fi36 < 2; fi36++ {
                                switch (fi36 + fi36) {
                                case 32768:
                                    if v2 {
                                        v4[(fi16 * 4294967295)] = fi36
                                        delete(v4, (^(fi16)))
                                    }
                                case 0:
                                    var v37 uint32 = 0
                                    _ = v37
                                case 4294967295:
                                    v38 := GenS2{F0: v2}
                                    _ = v38
                                default:
                                    var v39 uint = 1
                                    _ = v39
                                }
                                _ = fi36
                            }
                        default:
                            v40 := GenS1{F0: 0, F1: (v6 / (((^(v6))) | 1))}
                            _ = v40
                        }
                        var v41 uint = 4294967295
                        _ = v32
                        _ = v41
                    }
                }
                _ = v23
                _ = v24
            }
        case -150:
            if false {
                if true {
                    v3 = append(v3, 42)
                    v1 -= v1
                    fmt.Println(v1, fi16, v14)
                }
            } else {
                v6 = (v6 >> (((v1 * (^(v1)))) & 63))
                v6 -= 1
                var v42 uint32 = 2147483648
                _ = v42
            }
        case 42:
            var v43 uint16 = 919
            _ = v43
        }
        for fi44 := 0; fi44 < 3; fi44++ {
            fmt.Println(v7)
            _ = fi44
        }
        _ = fi16
    }
    v3 = append(v3, 32767)
    if (!(v2 || ((1 >= 2) || (!v2)))) {
        if ((!v2) && v2) {
            if false {
                _ = len(v3)
                for fi45 := 0; fi45 < 16; fi45++ {
                    switch fi45 {
                    case -999:
                        var v46 bool = (v1 != (v1 | 2))
                        _ = v46
                    case 128:
                        var v47 int8 = -118
                        _ = v47
                    case 1:
                        var v48 uint8 = 145
                        _ = v48
                    default:
                        if true {
                            for fi49 := 0; fi49 < 19; fi49++ {
                                var v50 int32 = 709
                                _ = fi49
                                _ = v50
                            }
                        }
                    }
                    var v51 int8 = 101
                    var v52 int8 = v51
                    _ = fi45
                    _ = v51
                    _ = v52
                }
            }
        }
    } else {
        switch 861 {
        case -613:
            fmt.Println(v8)
        case 32767:
            fmt.Println(v6, v14)
        }
        for fi53 := 0; fi53 < 14; fi53++ {
            if true {
                v8 -= 1.5
            }
            v2 = ("a" < ((("\n\t" + "") + "") + "foo bar"))
            v3 = append(v3, -534)
            _ = fi53
        }
        fmt.Println(v14)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = v2
    switch "a" {
    case "a":
        v4[9223372036854775807] = 2147483647
    case "":
        var v54 uint8 = (v6 / ((v6) | 1))
        _ = v54
    case "\n\t":
        if v2 {
            var v55 uint16 = 0
            _ = v55
        }
    }
    for fi56 := 0; fi56 < 12; fi56++ {
        if v2 {
            var v57 uint8 = (v6 * 127)
            if ((v5 - v5) == v5) {
                for fi58 := 0; fi58 < 18; fi58++ {
                    var v59 int16 = (v5 - -740)
                    var v60 uint32 = 4294967295
                    v61 := GenS2{F0: true}
                    v62 := GenS1{F0: (v14 - 1), F1: (v57 / ((v57) | 1))}
                    _ = len(v4)
                    var v63 int64 = 932
                    v64 := GenS3{F0: (-(v63))}
                    v63 = (v63 / (((v63 + -9223372036854775808)) | 1))
                    _ = fi58
                    _ = v59
                    _ = v60
                    _ = v61
                    _ = v62
                    _ = v63
                    _ = v64
                }
                if v2 {
                    if false {
                        fmt.Println(v1, v7)
                        v5 -= (v5 >> (((v1 ^ 2)) & 63))
                    }
                    for fi65 := 0; fi65 < 15; fi65++ {
                        if (fi65 > (fi65 | fi65)) {
                            var v66 float64 = (1.5 - 0.1)
                            if v2 {
                                switch -175 {
                                case 1:
                                    if true {
                                        v4[127] = int(v57)
                                    }
                                    v66 = float64(1e10)
                                    v2 = true
                                    if v2 {
                                        v67 := GenS2{F0: true}
                                        _ = v67
                                    } else {
                                        v3 = append(v3, fi65)
                                        switch fi65 {
                                        case 119:
                                            var v68 int = fi56
                                            _ = v68
                                        case -174:
                                            var v69 uint8 = v57
                                            _ = v69
                                        default:
                                            var v70 uint8 = 42
                                            _ = v70
                                        }
                                        v66 = 3.14159
                                    }
                                    switch (fi56 + -219) {
                                    case 949:
                                        v5 += v5
                                    case 2147483647:
                                        v71 := GenS2{F0: v2}
                                        _ = v71
                                    case 4294967295:
                                        v4[32767] = -859
                                    default:
                                        var v72 bool = (v5 <= v5)
                                        _ = v72
                                    }
                                default:
                                    switch "a" {
                                    case "\n\t":
                                        if ((^(v7)) == (v7 & (v7 | (v7 ^ (v7 | (v7 + 2147483647)))))) {
                                            v8 = 0.1
                                            var v73 uint32 = 4294967295
                                            _ = v73
                                        }
                                    }
                                }
                            }
                            var v74 uint = 18446744073709551615
                            if true {
                                fmt.Println(v14)
                                v7 += v7
                            }
                            fmt.Println(v5, v6, v7)
                            v3 = append(v3, (-(fi56)))
                            v75 := GenS1{F0: v14, F1: v6}
                            _ = v66
                            _ = v74
                            _ = v75
                        }
                        _ = fi65
                    }
                    v76 := GenS4{F0: (fi56 | 32767), F1: -2147483648}
                    _ = v76
                }
            }
            _ = v57
        }
        _ = fi56
    }
    switch -873 {
    case 2147483647:
        v14 = v14
    case 32768:
        for fi77 := 0; fi77 < 1; fi77++ {
            v78 := GenS5{F0: (v1 * 2), F1: v6, F2: v5}
            _ = fi77
            _ = v78
        }
    case 1:
        v14 = 734
    }
    v5 = v5
    switch "a" {
    case "":
        for fi79 := 0; fi79 < 2; fi79++ {
            _ = len(v3)
            v8 = (v8 / v8)
            _ = len(v3)
            _ = fi79
        }
    case "a":
        if (393 > 42) {
            var v80 string = "hello"
            v80 = (v80 + v80)
            var v81 uint32 = 150
            _ = v80
            _ = v81
        }
    case "foo bar":
        var v82 uint64 = v14
        _ = v82
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
