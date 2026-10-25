package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 uint; F2 string }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16; F1 int }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint8, p1 bool, p2 int16) string {
    var v7 string = "hello"
    var v8 bool = (18446744073709551615 != 935)
    p0 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return ("foo bar" + v7)
}

func main() {
    var v1 int8 = -1
    var v2 uint = 814
    v3 := []int{32768, 914, -450}
    v4 := []uint32{252, 0, 65535}
    var v5 int8 = v1
    v6 := make(map[int]uint32)
    var v9 int64 = -327
    fmt.Println(v5)
    if false {
        v10 := GenS1{F0: 42}
        _ = v10
    } else {
        if true {
            for fi11 := 0; fi11 < 14; fi11++ {
                v9 = (-(v9))
                var v12 float32 = 1.5
                _ = fi11
                _ = v12
            }
            _ = len(v3)
            switch "hello" {
            case "a":
                for fi13 := 0; fi13 < 7; fi13++ {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        var v15 float32 = float32(0.1)
                        v1 = (v5 / (((v5 / ((-113) | 1))) | 1))
                        switch (fi14 % (((fi14 >> (((^(v2))) & 63))) | 1)) {
                        case 32768:
                            v16 := v6[fi13]
                            v16 = v16
                            var v17 int64 = 42
                            _ = v16
                            _ = v17
                        default:
                            var v18 float32 = float32((float64(v15) / 1.0))
                            _ = v18
                        }
                        _ = fi14
                        _ = v15
                    }
                    v9 = (v9 >> ((v2) & 63))
                    for fi19 := 0; fi19 < 7; fi19++ {
                        switch (fi19 >> ((v2) & 63)) {
                        case 2147483647:
                            var v20 int32 = 463
                            _ = v20
                        default:
                            switch "" {
                            case "foo bar":
                                if (((v9 >> ((v2) & 63)) >= v9) && false) {
                                    v9 += 1
                                    v2 = v2
                                }
                            }
                        }
                        _ = fi19
                    }
                    _ = fi13
                }
            case "foo bar":
                for fi21 := 0; fi21 < 2; fi21++ {
                    var v22 uint8 = 150
                    _ = fi21
                    _ = v22
                }
            case "":
                if false {
                    if true {
                        for fi23 := 0; fi23 < 8; fi23++ {
                            for fi24 := 0; fi24 < 8; fi24++ {
                                v5 = (v5 / ((1) | 1))
                                var v25 uint8 = 128
                                var v26 float32 = float32((0.0 - (((((1e10 * 0.1) + (1e10 / -2.25)) + (1e10 + 0.1)) * (3.14159 + (((((1.5 - (1.5 * 1e10)) / 1.0) * 3.14159) + (1.5 * (((((-2.25 * 1e10) - 3.14159) + (float64((3.14159 * 1e10)) * (1.5 / -2.25))) - 0.0) - 1.5))) * 0.1))) - -2.25)))
                                _ = fi24
                                _ = v25
                                _ = v26
                            }
                            switch -9223372036854775808 {
                            case -914:
                                if false {
                                    if (42 != 1) {
                                        v27 := GenS1{F0: -2147483648}
                                        fmt.Println(v1, v1, fi23)
                                        _ = v27
                                    }
                                    fmt.Println(v5)
                                    var v28 float64 = 0.0
                                    _ = v28
                                } else {
                                    fmt.Println(v2)
                                    v29 := GenS2{F0: 2147483647, F1: fi23}
                                    for fi30 := 0; fi30 < 14; fi30++ {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        if ((v2 / ((v2) | 1)) < v2) {
                                            v31 := GenS1{F0: 2147483647}
                                            var v32 uint32 = 2147483647
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = fi30
                                    }
                                    _ = v29
                                }
                            case 32767:
                                for fi33 := 0; fi33 < 8; fi33++ {
                                    switch "foo bar" {
                                    case "\n\t":
                                        v2 -= (v2 / ((v2) | 1))
                                    default:
                                        v2 -= (v2 - (v2 << (((v2 << ((154) & 63))) & 63)))
                                    }
                                    v4 = append(v4, 124)
                                    _ = fi33
                                }
                            }
                            fmt.Println(v9)
                            _ = fi23
                        }
                        if false {
                            if true {
                                v9 -= v9
                                var v34 bool = true
                                _ = v34
                            }
                            delete(v6, 1)
                            for fi35 := 0; fi35 < 19; fi35++ {
                                v9 += (v9 ^ v9)
                                switch "hello" {
                                case "a":
                                    switch fi35 {
                                    case 42:
                                        v36 := GenS2{F0: 840, F1: fi35}
                                        _ = v36
                                    case 32768:
                                        v6[(-(fi35))] = 4294967295
                                    }
                                case "hello":
                                    var v37 string = ("foo bar" + "foo bar")
                                    _ = v37
                                default:
                                    _ = len(v6)
                                }
                                for fi38 := 0; fi38 < 12; fi38++ {
                                    v1 = (v5 + 0)
                                    if (1 > (v9 | 0)) {
                                        if (!true) {
                                            v39 := GenS3{F0: 4294967295}
                                            switch ("foo bar" + "a") {
                                            case "":
                                                var v40 int32 = -2147483648
                                                _ = v40
                                            case "hello":
                                                var v41 int16 = 918
                                                _ = v41
                                            case "foo bar":
                                                fmt.Println(v1, fi35)
                                            }
                                            delete(v6, (fi38 * fi38))
                                            _ = v39
                                        }
                                        var v42 string = ((("a" + "") + ((("foo bar" + "hello") + "\n\t") + "foo bar")) + "hello")
                                        var v43 int8 = (v5 ^ (v5 << (((v2 / (((v2 - v2)) | 1))) & 63)))
                                        _ = v42
                                        _ = v43
                                    }
                                    fmt.Println(v9, v9)
                                    _ = fi38
                                }
                                _ = fi35
                            }
                        } else {
                            if (!((320 <= 65535) || false)) {
                                if true {
                                    v44 := GenS3{F0: 992}
                                    var v45 float32 = 0.0
                                    _ = v44
                                    _ = v45
                                } else {
                                    for fi46 := 0; fi46 < 15; fi46++ {
                                        v47 := v6[-9223372036854775808]
                                        v48 := GenS1{F0: 1}
                                        _ = fi46
                                        _ = v47
                                        _ = v48
                                    }
                                    for fi49 := 0; fi49 < 17; fi49++ {
                                        v50 := GenS2{F0: 2147483647, F1: (fi49 << ((v2) & 63))}
                                        v51 := GenS2{F0: 2147483647, F1: fi49}
                                        _ = fi49
                                        _ = v50
                                        _ = v51
                                    }
                                    if ((v9 & 0) < int64(v1)) {
                                        switch "hello" {
                                        case "hello":
                                            v1 = (v5 & 42)
                                        case "\n\t":
                                            v6[819] = 2147483648
                                        case "a":
                                            fmt.Println(v1, v2)
                                        }
                                        if (!(2147483647 < -2147483648)) {
                                            if (157 > 128) {
                                                var v52 uint = v2
                                                v52 = (v2 | 42)
                                                v53 := GenS4{F0: (v5 / ((v5) | 1)), F1: -741}
                                                _ = v52
                                                _ = v53
                                            } else {
                                                v3 = append(v3, 194)
                                                v1 -= -128
                                            }
                                        }
                                        v9 = (v9 | v9)
                                    }
                                }
                                if true {
                                    var v54 float32 = 3.14159
                                    _ = v54
                                }
                                v55 := GenS5{F0: v5}
                                _ = v55
                            } else {
                                v56 := GenS1{F0: 904}
                                _ = v56
                            }
                        }
                        for fi57 := 0; fi57 < 2; fi57++ {
                            for fi58 := 0; fi58 < 12; fi58++ {
                                v2 = (v2 - v2)
                                var v59 uint = v2
                                _ = fi58
                                _ = v59
                            }
                            v5 = (-(v5))
                            _ = fi57
                        }
                    }
                }
            default:
                for fi60 := 0; fi60 < 9; fi60++ {
                    for fi61 := 0; fi61 < 7; fi61++ {
                        var v62 int32 = -1
                        _ = fi61
                        _ = v62
                    }
                    if (238 >= 2) {
                        _ = len(v6)
                        if true {
                            v63 := GenS2{F0: 414, F1: (-(fi60))}
                            _ = v63
                        }
                    } else {
                        for fi64 := 0; fi64 < 3; fi64++ {
                            v65 := GenS3{F0: 684}
                            var v66 int16 = 42
                            _ = fi64
                            _ = v65
                            _ = v66
                        }
                    }
                    _ = fi60
                }
            }
        } else {
            fmt.Println(v1, v9)
            var v67 int16 = 0
            var v68 uint64 = 755
            _ = v67
            _ = v68
        }
    }
    _ = len(v6)
    v9 -= (v9 * (v9 & (v9 * (v9 | v9))))
    var v69 int16 = 32767
    for fi70 := 0; fi70 < 19; fi70++ {
        for fi71 := 0; fi71 < 17; fi71++ {
            v72 := GenS6{F0: 4294967295, F1: v2, F2: "\n\t"}
            fmt.Println(v1, v2, v5)
            v3 = append(v3, (fi71 - fi70))
            _ = fi71
            _ = v72
        }
        var v73 uint64 = 408
        _ = fi70
        _ = v73
    }
    v74 := GenS2{F0: 1, F1: 42}
    for fi75 := 0; fi75 < 16; fi75++ {
        v5 += int8(v9)
        if false {
            for fi76 := 0; fi76 < 11; fi76++ {
                v77 := GenS2{F0: 330, F1: fi76}
                _ = fi76
                _ = v77
            }
            if ((^(fi75)) >= (fi75 >> (((v2 >> ((4294967295) & 63))) & 63))) {
                v78 := GenS7{F0: 36, F1: (-(v5))}
                v79 := GenS3{F0: 0}
                _ = v78
                _ = v79
            } else {
                var v80 string = (("a" + "\n\t") + ("foo bar" + ((("a" + "\n\t") + "\n\t") + "foo bar")))
                if (1.5 <= 1.5) {
                    if true {
                        var v81 int8 = v5
                        _ = v81
                    } else {
                        var v82 uint = (v2 >> ((585) & 63))
                        _ = v82
                    }
                    fmt.Println(v9, fi75)
                    v9 = -9223372036854775808
                } else {
                    for fi83 := 0; fi83 < 12; fi83++ {
                        if ((v5 << ((v2) & 63)) <= v5) {
                            v84 := GenS8{F0: int32(v1)}
                            switch ("a" + v80) {
                            case "\n\t":
                                v80 = ""
                            case "hello":
                                _ = len(v6)
                            case "":
                                v5 += (v5 | (v5 / (((v5 / ((v5) | 1))) | 1)))
                            }
                            var v85 uint = (v2 / ((0) | 1))
                            _ = v84
                            _ = v85
                        }
                        if false {
                            var v86 uint32 = 64
                            var v87 string = ("hello" + v80)
                            _ = v86
                            _ = v87
                        } else {
                            var v88 int = 128
                            _ = v88
                        }
                        v69 += -32768
                        v80 = v80
                        _ = fi83
                    }
                }
                _ = v80
            }
            for fi89 := 0; fi89 < 19; fi89++ {
                v5 += v1
                if false {
                    if true {
                        var v90 int8 = (-(v5))
                        _ = v90
                    }
                    if false {
                        v1 = (v1 & 1)
                        v9 -= (v9 & (v9 % (((-(v9))) | 1)))
                        v1 = int8(v9)
                        fmt.Println(v2, v69)
                    } else {
                        fmt.Println(v69)
                        var v91 int = (fi89 + (fi75 / ((fi75) | 1)))
                        delete(v6, (v91 / ((fi89) | 1)))
                        v9 = (v9 + v9)
                        if (!((v2 - v2) <= (v2 * v2))) {
                            var v92 uint64 = 0
                            _ = v92
                        } else {
                            v1 -= v5
                            v2 -= 264
                        }
                        var v93 int8 = (v5 * (-(v1)))
                        v94 := GenS9{F0: 0, F1: (fi75 / ((726) | 1))}
                        var v95 uint32 = 0
                        v96 := GenS9{F0: uint16(v95), F1: (v91 ^ fi89)}
                        if ((("a" > "a") && true) && (!(0 <= 849))) {
                            v91 = (-(v91))
                            v97 := GenS10{F0: 159}
                            v95 = (v95 >> ((v2) & 63))
                            _ = v97
                        }
                        _ = v91
                        _ = v93
                        _ = v94
                        _ = v95
                        _ = v96
                    }
                    var v98 string = "hello"
                    switch v98 {
                    case "a":
                        if ((v9 | 0) != (-(v9))) {
                            v99 := GenS8{F0: 85}
                            _ = v99
                        } else {
                            var v100 uint16 = 65535
                            v101 := GenS11{F0: 611, F1: (^(v100))}
                            v2 += 42
                            _ = v100
                            _ = v101
                        }
                    case "":
                        v102 := GenS3{F0: 2147483648}
                        var v103 uint64 = 934
                        var v104 float32 = (1e10 / 1e10)
                        var v105 int8 = v5
                        fmt.Println(v103, v104, v69)
                        _ = v102
                        _ = v103
                        _ = v104
                        _ = v105
                    default:
                        var v106 uint16 = 255
                        _ = v106
                    }
                    v107 := GenS7{F0: 127, F1: v5}
                    _ = v107
                    _ = v98
                }
                _ = fi89
            }
        }
        _ = fi75
    }
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    switch ("hello" + ("hello" + ("hello" + "\n\t"))) {
    case "foo bar":
        v9 = int64(v2)
    case "\n\t":
        v4 = append(v4, 227)
    case "a":
        var v108 float64 = 0.0
        _ = v108
    }
    _ = len(v3)
    for fi109 := 0; fi109 < 15; fi109++ {
        fmt.Println(v2)
        v9 += v9
        _ = fi109
    }
    for fi110 := 0; fi110 < 15; fi110++ {
        v2 += v2
        _ = fi110
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v69:", v69)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v74
    _ = v9
}
