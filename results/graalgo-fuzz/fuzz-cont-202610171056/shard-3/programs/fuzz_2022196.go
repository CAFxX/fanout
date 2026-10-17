package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS5 struct { F0 string; F1 bool; F2 int8 }
type GenS6 struct { F0 int16; F1 int16; F2 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint32, p1 string) int8 {
    var v4 int16 = -32768
    p1 = p1
    for fi5 := 0; fi5 < 3; fi5++ {
        var v6 uint64 = 2
        _ = fi5
        _ = v6
    }
    if (4294967295 != 1) {
        if (252 < 0) {
            v7 := GenS1{F0: 1}
            _ = v7
        }
        var v8 float32 = 3.14159
        _ = v8
    }
    var v9 uint = 4294967295
    _ = p0
    _ = p1
    _ = v4
    _ = v9
    return 42
}
func genFn1(p0 int8, p1 string, p2 int) int32 {
    v10 := GenS2{F0: (p1 + ((("a" + p1) + p1) + "")), F1: 18446744073709551615}
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 int32 = 42
            _ = fi12
            _ = v13
        }
        var v14 int8 = p0
        var v15 int64 = 1
        _ = fi11
        _ = v14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2147483647
}

func main() {
    var v1 float32 = 1.5
    var v2 int = 2147483647
    var v3 uint64 = 251
    v16 := GenS3{F0: 584, F1: 2}
    v17 := GenS1{F0: v2}
    for fi18 := 0; fi18 < 3; fi18++ {
        if (!(-961 < 744)) {
            var v19 string = "foo bar"
            for fi20 := 0; fi20 < 15; fi20++ {
                v2 -= fi20
                if (false || false) {
                    v2 = (^(fi18))
                    var v21 int = 9223372036854775807
                    _ = v21
                } else {
                    if false {
                        v1 = v1
                        switch (v2 & 128) {
                        case 338:
                            v3 -= (^(v3))
                        case 4294967295:
                            v22 := []int{fi18, v2, v2}
                            _ = v22
                        }
                        v2 -= int(v3)
                        fmt.Println(v1, v1, fi20)
                        v3 = (v3 & (v3 | 42))
                        v3 -= (v3 & v3)
                        v3 = v3
                    } else {
                        v23 := []int{fi18, (-(fi18))}
                        switch v19 {
                        case "a":
                            v23 = append(v23, v2)
                        case "\n\t":
                            switch "" {
                            case "hello":
                                v24 := GenS4{F0: 13, F1: 32767, F2: int8(fi18)}
                                _ = v24
                            case "":
                                v25 := GenS1{F0: (v2 - v2)}
                                _ = v25
                            case "\n\t":
                                if len(v23) >= 2 {
                                    v23 = v23[1:len(v23)]
                                }
                            default:
                                if false {
                                    if true {
                                        switch (fi18 ^ 32767) {
                                        case 9223372036854775807:
                                            v2 = fi18
                                        case -315:
                                            if (false || false) {
                                                if (true && true) {
                                                    var v26 string = (v19 + (v19 + v19))
                                                    var v27 float32 = (v1 * (0.0 - v1))
                                                    if len(v23) > 0 {
                                                        v23[0] = (fi20 & (fi20 >> ((255) & 63)))
                                                    }
                                                    if (!false) {
                                                        v2 += (-(fi20))
                                                        fmt.Println(v2)
                                                    } else {
                                                        if false {
                                                            v28 := GenS3{F0: 1, F1: 29}
                                                            v26 = "hello"
                                                            if true {
                                                                if false {
                                                                    var v29 uint = 669
                                                                    v3 -= (v3 >> (((v29 ^ (v29 >> ((v29) & 63)))) & 63))
                                                                    fmt.Println(v26, v3, v27)
                                                                    _ = v29
                                                                } else {
                                                                    if false {
                                                                        v30 := GenS2{F0: (v19 + "hello"), F1: 65535}
                                                                        switch (fi18 / ((42) | 1)) {
                                                                        case -1:
                                                                            v27 -= float32(0.0)
                                                                        case 128:
                                                                            fmt.Println(v26, v1)
                                                                        case -9223372036854775808:
                                                                            var v31 int = 4294967295
                                                                            _ = v31
                                                                        default:
                                                                            switch (fi20 << ((4294967295) & 63)) {
                                                                            case -1:
                                                                                if (!false) {
                                                                                    if len(v23) >= 2 {
                                                                                        v23 = v23[1:len(v23)]
                                                                                    }
                                                                                    v19 = "a"
                                                                                }
                                                                            default:
                                                                                var v32 int32 = -1
                                                                                _ = v32
                                                                            }
                                                                        }
                                                                        _ = v30
                                                                    }
                                                                    var v33 int = fi18
                                                                    v23 = append(v23, (fi18 / (((v2 - 42)) | 1)))
                                                                    _ = v33
                                                                }
                                                                fmt.Println(v1)
                                                            }
                                                            _ = v28
                                                        }
                                                        var v34 uint64 = (^(v3))
                                                        if false {
                                                            v35 := GenS5{F0: "foo bar", F1: (32767 <= 673), F2: 1}
                                                            _ = v35
                                                        } else {
                                                            v36 := GenS2{F0: v26, F1: 255}
                                                            _ = v36
                                                        }
                                                        v3 -= (v3 | 1)
                                                        _ = v34
                                                    }
                                                    var v37 uint16 = 42
                                                    var v38 bool = (!((v19 != v19) || ((!(42 > 42)) || false)))
                                                    var v39 uint32 = 2147483647
                                                    v38 = ((1 >= 42) || v38)
                                                    _ = v26
                                                    _ = v27
                                                    _ = v37
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    var v40 int32 = -1
                                                    v2 -= (fi18 - (^(fi18)))
                                                    v41 := GenS5{F0: v19, F1: (!((("hello" + (((v19 + "foo bar") + (((v19 + ("foo bar" + ("a" + "\n\t"))) + v19) + ("a" + "foo bar"))) + v19)) == "a") && (!(255 < 93)))), F2: 100}
                                                    _ = v40
                                                    _ = v41
                                                }
                                            } else {
                                                switch ((("foo bar" + "a") + v19) + v19) {
                                                case "a":
                                                    v42 := GenS6{F0: -32768, F1: -32768, F2: 42}
                                                    _ = v42
                                                case "":
                                                    if (false && false) {
                                                        switch v19 {
                                                        case "a":
                                                            var v43 int = 1
                                                            _ = v43
                                                        case "foo bar":
                                                            v1 += ((v1 - (v1 + 0.0)) * (float32(0.0) / (v1 * v1)))
                                                        }
                                                        fmt.Println(v3, fi18)
                                                        switch (fi18 << ((247) & 63)) {
                                                        case 2147483647:
                                                            var v44 int8 = 0
                                                            _ = v44
                                                        case 32767:
                                                            fmt.Println(v1)
                                                        default:
                                                            v3 += 4294967295
                                                        }
                                                    }
                                                    var v45 uint64 = 2
                                                    switch "foo bar" {
                                                    case "a":
                                                        var v46 string = v19
                                                        var v47 int64 = 42
                                                        v45 -= 208
                                                        _ = v46
                                                        _ = v47
                                                    case "foo bar":
                                                        v3 = 822
                                                    case "":
                                                        _ = len(v23)
                                                    }
                                                    v2 += fi18
                                                    _ = v45
                                                case "\n\t":
                                                    v3 = v3
                                                default:
                                                    var v48 int8 = genFn0(1, (v19 + v19))
                                                    _ = v48
                                                }
                                            }
                                        case 32768:
                                            v49 := GenS6{F0: 1, F1: 42, F2: 281}
                                            _ = v49
                                        }
                                        v19 = (v19 + v19)
                                        var v50 int8 = 0
                                        v51 := GenS7{F0: -1}
                                        fmt.Println(v1, v19, v3)
                                        v52 := GenS1{F0: (fi20 - -2147483648)}
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                    } else {
                                        switch v19 {
                                        case "foo bar":
                                            if len(v23) >= 2 {
                                                v23 = v23[1:len(v23)]
                                            }
                                        case "a":
                                            var v53 int64 = -629
                                            _ = v53
                                        case "hello":
                                            var v54 float64 = (float64(((float32((float64(v1) - float64(v1))) * v1) - (v1 / float32((float64(v1) / (3.14159 / 3.14159)))))) * (0.1 * 1.5))
                                            _ = v54
                                        default:
                                            v2 = (-(fi18))
                                        }
                                    }
                                    var v55 uint8 = 214
                                    _ = v55
                                }
                            }
                        }
                        _ = v23
                    }
                    v2 -= 32768
                    v3 = 2
                    v2 = fi20
                    v56 := []int{-454}
                    fmt.Println(v1)
                    _ = v56
                }
                v57 := GenS4{F0: 2, F1: 1, F2: 1}
                var v58 uint = 1
                var v59 float64 = (((0.0 + 0.1) * 3.14159) - (0.0 - (-2.25 - float64((v1 * (v1 / v1))))))
                _ = fi20
                _ = v57
                _ = v58
                _ = v59
            }
            _ = v19
        }
        fmt.Println(v2, v2)
        _ = fi18
    }
    fmt.Println(v3)
    var v60 string = "a"
    fmt.Println(v2)
    v61 := GenS2{F0: v60, F1: 916}
    var v62 uint64 = (v3 * v3)
    for fi63 := 0; fi63 < 2; fi63++ {
        v3 -= v62
        var v64 int = -9223372036854775808
        _ = fi63
        _ = v64
    }
    if false {
        if (0 <= 42) {
            for fi65 := 0; fi65 < 20; fi65++ {
                var v66 float32 = v1
                _ = fi65
                _ = v66
            }
            switch v60 {
            case "foo bar":
                v67 := GenS2{F0: v60, F1: 1}
                _ = v67
            case "a":
                var v68 uint32 = 2
                _ = v68
            }
            switch (v2 ^ v2) {
            case 42:
                fmt.Println(v62, v1)
            case -538:
                if true {
                    v69 := []int{v2, (v2 * v2)}
                    fmt.Println(v3)
                    _ = v69
                }
            }
        }
        fmt.Println(v3)
    }
    switch v2 {
    case 0:
        for fi70 := 0; fi70 < 19; fi70++ {
            for fi71 := 0; fi71 < 4; fi71++ {
                fmt.Println(v60, fi70, v2)
                var v72 int64 = 0
                v1 -= 3.14159
                _ = fi71
                _ = v72
            }
            _ = fi70
        }
    case 127:
        for fi73 := 0; fi73 < 16; fi73++ {
            switch v60 {
            case "foo bar":
                if (963 != 633) {
                    for fi74 := 0; fi74 < 20; fi74++ {
                        v75 := GenS5{F0: v60, F1: true, F2: 0}
                        var v76 float32 = v1
                        _ = fi74
                        _ = v75
                        _ = v76
                    }
                    v77 := GenS2{F0: "\n\t", F1: 4294967295}
                    _ = v77
                } else {
                    var v78 uint = 4294967295
                    _ = v78
                }
            default:
                var v79 uint64 = 18446744073709551615
                _ = v79
            }
            v1 -= v1
            for fi80 := 0; fi80 < 7; fi80++ {
                if false {
                    v1 = v1
                    v81 := GenS8{F0: 2147483647, F1: -578}
                    v82 := GenS7{F0: 0}
                    _ = v81
                    _ = v82
                }
                _ = fi80
            }
            _ = fi73
        }
    default:
        v83 := GenS9{F0: -9223372036854775808, F1: (^(v2))}
        _ = v83
    }
    v84 := []int{4294967295, v2}
    var v85 int8 = -1
    for fi86 := 0; fi86 < 19; fi86++ {
        var v87 int64 = -696
        if false {
            v88 := GenS8{F0: 0, F1: 32767}
            for fi89 := 0; fi89 < 15; fi89++ {
                v2 = -732
                var v90 uint32 = 492
                fmt.Println(v85, v85, v90)
                _ = fi89
                _ = v90
            }
            switch v60 {
            case "":
                switch (fi86 << ((313) & 63)) {
                case -9223372036854775808:
                    fmt.Println(v85, v87, v85)
                case 170:
                    var v91 uint64 = (v62 | 4294967295)
                    _ = v91
                default:
                    switch 0 {
                    case 42:
                        fmt.Println(v3, v2)
                    case 672:
                        v85 += 1
                    case -1:
                        var v92 int16 = 42
                        _ = v92
                    default:
                        v2 += v2
                    }
                }
            default:
                if false {
                    v2 += (v2 + 0)
                    for fi93 := 0; fi93 < 9; fi93++ {
                        _ = len(v84)
                        _ = fi93
                    }
                    if len(v84) > 0 {
                        v84[0] = (v2 & (fi86 % (((v2 & fi86)) | 1)))
                    }
                }
            }
            _ = v88
        } else {
            v94 := GenS10{F0: 1, F1: 110}
            if (!((127 < 255) && (false || (((v85 / (((^(v85))) | 1)) == genFn0(65535, "hello")) && (!((!false) || (((v2 / ((-2147483648) | 1)) < fi86) && true))))))) {
                fmt.Println(v60, fi86, v85)
                var v95 uint16 = 766
                fmt.Println(v1, fi86, v87)
                _ = v95
            } else {
                for fi96 := 0; fi96 < 11; fi96++ {
                    v2 -= v2
                    var v97 uint32 = 927
                    _ = fi96
                    _ = v97
                }
            }
            v2 = (fi86 % ((-9223372036854775808) | 1))
            _ = v94
        }
        _ = fi86
        _ = v87
    }
    v1 = -2.25
    var v98 bool = (233 > 242)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    fmt.Println("slicelen v84:", len(v84), v84)
    fmt.Println("v85:", v85)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v60
    _ = v61
    _ = v62
    _ = v84
    _ = v85
    _ = v98
}
