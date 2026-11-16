package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint32; F2 int8 }
type GenS4 struct { F0 int8; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint8) bool {
    v7 := GenS1{F0: 2147483647, F1: ("hello" + ("a" + ("\n\t" + ("" + "foo bar"))))}
    fmt.Println(p0)
    p0 = 179
    var v8 bool = false
    _ = p0
    _ = v7
    _ = v8
    return ((2147483647 <= 863) && (!false))
}
func genFn1(p0 int8, p1 uint32, p2 string) uint {
    p2 = p2
    p2 = p2
    var v9 int64 = -766
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2
}

func main() {
    v1 := []uint32{}
    var v2 int16 = 1
    var v3 int = 432
    var v4 uint64 = 18446744073709551615
    var v5 int64 = -256
    var v6 float32 = ((3.14159 / ((1e10 + 0.1) * ((-2.25 + ((3.14159 / -2.25) + (0.1 / 3.14159))) - -2.25))) + float32(1.5))
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(v6, v2, v6)
        var v11 uint64 = v4
        _ = fi10
        _ = v11
    }
    v5 = (v5 + -1)
    if true {
        var v12 int = v3
        _ = v12
    }
    var v13 int8 = -1
    var v14 float32 = (v6 - v6)
    var v15 int32 = 1
    if (false && true) {
        v6 = (v6 - ((0.1 * v14) / 0.1))
        v16 := GenS1{F0: v15, F1: ("" + ("\n\t" + "foo bar"))}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v16
    } else {
        var v17 int16 = 788
        _ = v17
    }
    var v18 uint64 = (v4 / ((1) | 1))
    if ((false || false) || ((v2 & -32768) >= (v2 % ((1) | 1)))) {
        for fi19 := 0; fi19 < 16; fi19++ {
            if (-110 < v13) {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    v20 := GenS2{F0: v18}
                    _ = v20
                case "hello":
                    if false {
                        if (true && false) {
                            var v21 bool = ((830 == 425) && false)
                            if v21 {
                                for fi22 := 0; fi22 < 16; fi22++ {
                                    if v21 {
                                        v23 := GenS2{F0: v4}
                                        if (-481 != fi19) {
                                            fmt.Println(v21, fi22)
                                            if v21 {
                                                if true {
                                                    if v21 {
                                                        if (v18 >= (v18 + (^(v18)))) {
                                                            v21 = v21
                                                        } else {
                                                            var v24 bool = (v13 == (v13 * v13))
                                                            var v25 string = ""
                                                            _ = v24
                                                            _ = v25
                                                        }
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v1 = append(v1, 1)
                                                    }
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = 42
                                                    }
                                                    if true {
                                                        v13 += int8(fi22)
                                                        v26 := GenS3{F0: 0, F1: 4294967295, F2: (v13 >> ((0) & 63))}
                                                        var v27 int64 = -1
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = 997
                                                }
                                                var v28 uint64 = (v18 >> ((1) & 63))
                                                var v29 float32 = v6
                                                _ = len(v1)
                                                _ = v28
                                                _ = v29
                                            } else {
                                                v15 += v15
                                                v30 := GenS3{F0: (v5 >> ((345) & 63)), F1: 383, F2: (-(v13))}
                                                var v31 float64 = (0.1 + float64(((float32(1e10) * v6) / float32(-2.25))))
                                                _ = v30
                                                _ = v31
                                            }
                                        } else {
                                            v32 := GenS1{F0: v15, F1: "a"}
                                            fmt.Println(v3, v13)
                                            _ = v32
                                        }
                                        if ((float32((1.5 + ((-2.25 * ((float64((((v6 * (v6 / (v14 / 1.0))) + (v14 + (float32(((0.0 * (0.1 / float64(0.1))) + 0.1)) + v14))) / (float32(float64(1.5)) * 0.1))) * 1.5) - 1e10)) * 3.14159))) - (0.0 * ((v6 + v6) + v6))) != (v14 * -2.25)) {
                                            v33 := GenS3{F0: v5, F1: 14, F2: v13}
                                            v21 = false
                                            v34 := GenS2{F0: (^(v18))}
                                            _ = v33
                                            _ = v34
                                        } else {
                                            v13 = v13
                                            fmt.Println(v15, fi19, v21)
                                            v13 += 76
                                        }
                                        _ = v23
                                    } else {
                                        v3 -= 9223372036854775807
                                    }
                                    switch "hello" {
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "":
                                            var v35 bool = v21
                                            _ = v35
                                        case "\n\t":
                                            if (!(v21 && ((-(v3)) != -1))) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v36 int = v3
                                                var v37 float32 = v6
                                                _ = v36
                                                _ = v37
                                            } else {
                                                fmt.Println(v15, v14)
                                                var v38 uint32 = 1
                                                fmt.Println(v21, fi19)
                                                _ = v38
                                            }
                                        }
                                    }
                                    _ = fi22
                                }
                            } else {
                                switch 42 {
                                case -850:
                                    var v39 uint = 42
                                    _ = v39
                                default:
                                    fmt.Println(v3, v18, v4)
                                }
                            }
                            _ = v21
                        } else {
                            for fi40 := 0; fi40 < 9; fi40++ {
                                v4 += 2
                                _ = len(v1)
                                v15 = v15
                                _ = fi40
                            }
                            v1 = append(v1, 65535)
                            for fi41 := 0; fi41 < 11; fi41++ {
                                var v42 uint = 0
                                _ = fi41
                                _ = v42
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                }
            }
            _ = fi19
        }
        v6 -= 0.0
    } else {
        var v43 bool = (((float64((v6 - v14)) * (0.0 * ((1.5 * 0.1) - -2.25))) >= float64((((v14 + v6) / 3.14159) - (v14 + (v6 * float32((((-2.25 - (3.14159 - 3.14159)) * float64(3.14159)) + (3.14159 * (((((0.0 * (-2.25 + (-2.25 / 1.0))) + 1.5) - ((0.1 * 0.1) / float64((v6 + 0.1)))) + -2.25) - ((3.14159 - -2.25) * 1.5)))))))))) || false)
        fmt.Println(v2)
        var v44 int = v3
        _ = v43
        _ = v44
    }
    var v45 int16 = -771
    v46 := GenS4{F0: (v13 * (v13 << ((1) & 63))), F1: 127, F2: v5}
    switch "\n\t" {
    case "a":
        for fi47 := 0; fi47 < 1; fi47++ {
            var v48 float32 = (v14 * v6)
            switch "\n\t" {
            case "\n\t":
                v14 -= float32(float64((0.1 / v6)))
            case "a":
                for fi49 := 0; fi49 < 6; fi49++ {
                    switch "hello" {
                    case "hello":
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                    case "a":
                        v14 -= 3.14159
                    case "":
                        var v50 int64 = v5
                        _ = v50
                    default:
                        var v51 int = (v3 % ((fi49) | 1))
                        _ = v51
                    }
                    var v52 uint = genFn1(42, 2, "")
                    if len(v1) > 0 {
                        v1[0] = 201
                    }
                    _ = fi49
                    _ = v52
                }
            default:
                for fi53 := 0; fi53 < 12; fi53++ {
                    switch -2147483648 {
                    case -1:
                        var v54 uint64 = (^(v18))
                        _ = v54
                    case 2147483647:
                        fmt.Println(v45, v3)
                    case 128:
                        if true {
                            fmt.Println(v13)
                            switch "" {
                            case "a":
                                switch fi53 {
                                case 127:
                                    v55 := GenS3{F0: (v5 >> ((2) & 63)), F1: 4294967295, F2: v13}
                                    _ = v55
                                }
                            case "foo bar":
                                fmt.Println(v4, v3, v18)
                            case "hello":
                                v15 = v15
                            }
                            v4 = (v4 * v18)
                        }
                        var v56 int32 = (^(v15))
                        _ = v56
                    }
                    v57 := GenS2{F0: (v4 * (v18 + (v18 + 4294967295)))}
                    v58 := GenS4{F0: -128, F1: 147, F2: (v5 / (((v5 * 1)) | 1))}
                    switch (("hello" + ((("hello" + "hello") + ("a" + (("foo bar" + ("foo bar" + "hello")) + ("foo bar" + "\n\t")))) + "")) + "hello") {
                    case "":
                        var v59 int64 = (-(v5))
                        _ = v59
                    case "\n\t":
                        switch (fi47 - v3) {
                        case 997:
                            var v60 string = "hello"
                            _ = v60
                        }
                    default:
                        var v61 uint8 = 127
                        _ = v61
                    }
                    var v62 int32 = v15
                    _ = fi53
                    _ = v57
                    _ = v58
                    _ = v62
                }
            }
            _ = fi47
            _ = v48
        }
    case "foo bar":
        switch ("" + "") {
        case "hello":
            v3 = v3
        default:
            var v63 string = "hello"
            _ = v63
        }
    case "":
        switch (v3 % (((v3 & v3)) | 1)) {
        case -648:
            if ((float64(v14) < (0.1 * 3.14159)) && true) {
                var v64 int32 = (v15 >> ((729) & 63))
                _ = v64
            } else {
                for fi65 := 0; fi65 < 7; fi65++ {
                    fmt.Println(v5, v45)
                    fmt.Println(v45)
                    _ = fi65
                }
            }
        }
    }
    switch (("foo bar" + "a") + ("hello" + "hello")) {
    case "a":
        if (v3 >= (^(v3))) {
            for fi66 := 0; fi66 < 20; fi66++ {
                fmt.Println(v5)
                fmt.Println(v14, v6, v5)
                var v67 int64 = v5
                _ = fi66
                _ = v67
            }
            var v68 uint = 255
            _ = v68
        } else {
            if len(v1) > 0 {
                v1[0] = uint32(v4)
            }
            var v69 int8 = int8(v15)
            _ = v69
        }
    case "":
        v70 := GenS4{F0: (v13 << ((42) & 63)), F1: 2, F2: v5}
        _ = v70
    default:
        for fi71 := 0; fi71 < 13; fi71++ {
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if ((!(v5 < v5)) && false) {
                    if false {
                        var v72 string = ((("" + "") + "") + (("" + "hello") + "a"))
                        var v73 int16 = 1000
                        _ = v72
                        _ = v73
                    }
                    v74 := GenS2{F0: v18}
                    var v75 int64 = -9223372036854775808
                    _ = v74
                    _ = v75
                }
                for fi76 := 0; fi76 < 12; fi76++ {
                    _ = len(v1)
                    var v77 int = (fi71 >> ((697) & 63))
                    fmt.Println(v13)
                    _ = fi76
                    _ = v77
                }
            }
            if true {
                if false {
                    var v78 int8 = int8(fi71)
                    var v79 int = (fi71 / (((fi71 & (v3 % (((^(v3))) | 1)))) | 1))
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                    _ = v78
                    _ = v79
                }
                fmt.Println(v13, v5)
                v5 = v5
            } else {
                v14 += v6
                var v80 bool = (1 <= v5)
                fmt.Println(v3)
                _ = v80
            }
            _ = fi71
        }
    }
    for fi81 := 0; fi81 < 11; fi81++ {
        var v82 uint32 = 2147483648
        _ = fi81
        _ = v82
    }
    v14 += v6
    var v83 int64 = (v5 ^ v5)
    for fi84 := 0; fi84 < 13; fi84++ {
        fmt.Println(v5, v45)
        v85 := GenS5{F0: 128, F1: (v13 << ((65535) & 63))}
        _ = fi84
        _ = v85
    }
    v15 = v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v83
}
