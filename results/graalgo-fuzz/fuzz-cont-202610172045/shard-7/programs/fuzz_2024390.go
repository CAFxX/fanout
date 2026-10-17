package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint32; F1 bool; F2 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 string }
func genFn0(p0 int16) uint16 {
    for fi5 := 0; fi5 < 12; fi5++ {
        p0 += 42
        _ = fi5
    }
    if ((p0 - -32768) != p0) {
        var v6 int32 = -1
        v7 := []int{9223372036854775807, -9223372036854775808, 270, 42}
        _ = len(v7)
        _ = v6
        _ = v7
    }
    var v8 int32 = -1
    _ = p0
    _ = v8
    return 744
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    v3 := make(map[int]int)
    var v4 uint64 = 85
    var v9 int8 = 100
    v4 += (^(v4))
    v10 := GenS1{F0: 42, F1: 376, F2: genFn0(-1)}
    if false {
        var v11 int32 = 42
        fmt.Println(v11)
        _ = v11
    } else {
        switch "" {
        case "\n\t":
            var v12 float32 = 1.5
            _ = v12
        case "foo bar":
            v3[128] = -399
        default:
            var v13 int = 0
            _ = v13
        }
        v4 += 823
        if (0.1 < 0.0) {
            v4 = 1
            v9 += -38
        }
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(v9)
        _ = fi14
    }
    v4 += (^(v4))
    v9 = (v9 << ((4294967295) & 63))
    if true {
        v4 = (^(v4))
    } else {
        v15 := v3[972]
        if false {
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 int16 = -908
                for fi18 := 0; fi18 < 6; fi18++ {
                    fmt.Println(fi16)
                    _ = fi18
                }
                _ = fi16
                _ = v17
            }
        } else {
            v4 += 18446744073709551615
            if (!(-9223372036854775808 >= 0)) {
                v15 = (v15 * (^(v15)))
                switch "hello" {
                case "a":
                    if (!false) {
                        if (1 < 9223372036854775807) {
                            if (42 != 42) {
                                v4 = v4
                            } else {
                                var v19 float32 = -2.25
                                var v20 int32 = 2147483647
                                _ = v19
                                _ = v20
                            }
                            fmt.Println(v15)
                        }
                        var v21 bool = true
                        _ = v21
                    }
                case "\n\t":
                    var v22 float32 = float32((float64(((float32((((((((((1.5 - (1.5 / float64((0.0 - 1.5)))) + 1.5) * ((0.1 * (0.0 / (1e10 / 1e10))) * 3.14159)) - float64((3.14159 - (((0.1 / 0.1) - 1e10) / 1.0)))) - 0.0) / 1.5) / (0.1 - 1.5)) * 0.1) * 1e10)) - (0.0 * ((1e10 * (-2.25 - 1e10)) * (3.14159 * float32((float64(3.14159) + (0.0 + 0.1))))))) * 1e10)) - float64((1.5 + 1e10))))
                    _ = v22
                }
                fmt.Println(v9, v15)
            } else {
                v15 += (v15 % ((v15) | 1))
            }
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := GenS1{F0: -1, F1: 42, F2: 65535}
            for fi25 := 0; fi25 < 20; fi25++ {
                _ = len(v2)
                if true {
                    if (((1.5 / (1e10 - 0.1)) != 0.1) || true) {
                        if len(v2) > 0 {
                            v2[0] = -793
                        }
                        v15 += (v15 ^ fi25)
                    } else {
                        if (-78 != v9) {
                            var v26 bool = (!(false && (v9 >= v9)))
                            if false {
                                var v27 int64 = 42
                                _ = v27
                            }
                            switch "foo bar" {
                            case "\n\t":
                                if (1 <= 440) {
                                    if true {
                                        v15 += fi23
                                        fmt.Println(v4, v4, v4)
                                    }
                                    var v28 uint8 = 1
                                    var v29 string = "foo bar"
                                    if ((v4 ^ (v4 * (v4 - v4))) <= 2) {
                                        if len(v2) > 0 {
                                            v2[0] = (v15 % (((fi25 << ((950) & 63))) | 1))
                                        }
                                    }
                                    switch v29 {
                                    case "foo bar":
                                        if (!v26) {
                                            v28 -= 216
                                        } else {
                                            v15 = (fi25 & 127)
                                        }
                                    default:
                                        var v30 bool = v26
                                        _ = v30
                                    }
                                    _ = v28
                                    _ = v29
                                }
                            case "hello":
                                if true {
                                    fmt.Println(fi25)
                                    if (!v26) {
                                        if v26 {
                                            var v31 int16 = 1
                                            _ = v31
                                        } else {
                                            v26 = (807 <= 2147483647)
                                            v15 += (fi23 / ((fi25) | 1))
                                        }
                                    } else {
                                        var v32 float64 = 1e10
                                        if (v26 || (1 == 65535)) {
                                            if v26 {
                                                var v33 float32 = 0.1
                                                v2 = append(v2, (^(fi25)))
                                                var v34 uint8 = 2
                                                v32 -= v32
                                                _ = v33
                                                _ = v34
                                            }
                                            switch (("hello" + "\n\t") + "hello") {
                                            case "a":
                                                var v35 uint16 = 42
                                                _ = v35
                                            case "":
                                                v15 -= (fi25 ^ (fi25 % (((fi23 / ((v15) | 1))) | 1)))
                                            case "foo bar":
                                                delete(v1, "a")
                                            }
                                            v9 = (-(v9))
                                        } else {
                                            v4 = (v4 & v4)
                                        }
                                        switch (v15 % ((fi23) | 1)) {
                                        case -9223372036854775808:
                                            var v36 float64 = v32
                                            _ = v36
                                        }
                                        var v37 uint32 = 4294967295
                                        _ = v32
                                        _ = v37
                                    }
                                } else {
                                    v2 = append(v2, fi23)
                                    v4 = (^(v4))
                                    var v38 float64 = 3.14159
                                    _ = v38
                                }
                            }
                            _ = v26
                        } else {
                            v4 = 18446744073709551615
                        }
                    }
                    v39 := GenS1{F0: 866, F1: 65535, F2: 495}
                    var v40 float32 = 3.14159
                    _ = v39
                    _ = v40
                }
                var v41 float64 = (1e10 * 1e10)
                _ = fi25
                _ = v41
            }
            _ = fi23
            _ = v24
        }
        _ = v15
    }
    for fi42 := 0; fi42 < 17; fi42++ {
        fmt.Println(fi42)
        _ = fi42
    }
    for fi43 := 0; fi43 < 3; fi43++ {
        v4 += 1
        for fi44 := 0; fi44 < 13; fi44++ {
            v45 := GenS2{F0: 2147483648, F1: (4294967295 != 42), F2: v9}
            v9 -= (v9 | (v9 + (v9 % ((v9) | 1))))
            _ = fi44
            _ = v45
        }
        v46 := GenS2{F0: 2, F1: true, F2: (v9 / ((v9) | 1))}
        _ = fi43
        _ = v46
    }
    v47 := GenS3{F0: true}
    if true {
        var v48 float32 = 1e10
        _ = v48
    }
    if (0.0 < 0.1) {
        var v49 float32 = float32(1e10)
        var v50 int8 = (v9 * v9)
        for fi51 := 0; fi51 < 13; fi51++ {
            var v52 bool = (false && false)
            for fi53 := 0; fi53 < 12; fi53++ {
                var v54 uint16 = 523
                v55 := GenS4{F0: 0}
                var v56 bool = ((!v52) && v52)
                _ = fi53
                _ = v54
                _ = v55
                _ = v56
            }
            _ = fi51
            _ = v52
        }
        _ = v49
        _ = v50
    }
    switch 373 {
    case 32768:
        for fi57 := 0; fi57 < 20; fi57++ {
            for fi58 := 0; fi58 < 11; fi58++ {
                if ((v4 < (v4 >> ((195) & 63))) || false) {
                    v9 -= (-(v9))
                    fmt.Println(fi57)
                    v9 -= -69
                } else {
                    var v59 int32 = -2147483648
                    if (32767 > 922) {
                        switch ("" + ("foo bar" + "\n\t")) {
                        case "":
                            var v60 bool = false
                            _ = v60
                        case "hello":
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = (fi58 + 32767)
                                }
                            } else {
                                v61 := GenS1{F0: 0, F1: 255, F2: genFn0(1)}
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v61
                            }
                        case "foo bar":
                            switch "a" {
                            case "hello":
                                fmt.Println(fi57)
                            default:
                                v62 := GenS4{F0: 42}
                                _ = v62
                            }
                        default:
                            switch (fi57 << ((222) & 63)) {
                            case 32767:
                                v9 += v9
                            case -318:
                                fmt.Println(fi57, fi57, v4)
                            }
                        }
                    }
                    _ = v59
                }
                fmt.Println(v9, fi58, v9)
                fmt.Println(fi57, fi57)
                _ = fi58
            }
            if (2 == 4294967295) {
                v9 += (^(v9))
                var v63 float32 = (0.0 - 1.5)
                _ = len(v3)
                _ = v63
            }
            var v64 int64 = 0
            _ = fi57
            _ = v64
        }
    case 0:
        v65 := GenS4{F0: 255}
        _ = v65
    case 9223372036854775807:
        for fi66 := 0; fi66 < 9; fi66++ {
            for fi67 := 0; fi67 < 20; fi67++ {
                var v68 float32 = 1.5
                if false {
                    v9 = v9
                    if (!(true || false)) {
                        fmt.Println(fi66, v68)
                    }
                    v69 := GenS2{F0: 2147483647, F1: ((32767 <= 0) || ((!false) && (65535 == 0))), F2: (v9 ^ (v9 | (v9 | v9)))}
                    _ = v69
                }
                v9 -= (v9 % ((v9) | 1))
                v70 := GenS5{F0: 1, F1: (((("\n\t" + "foo bar") + (("foo bar" + "a") + "hello")) + ("\n\t" + "a")) + "hello")}
                fmt.Println(v68)
                _ = fi67
                _ = v68
                _ = v70
            }
            fmt.Println(fi66, fi66)
            _ = fi66
        }
    }
    if true {
        if (genFn0(401) > 31) {
            fmt.Println(v4)
            var v71 string = (("\n\t" + "hello") + (("\n\t" + (("foo bar" + "\n\t") + ("foo bar" + ("\n\t" + "hello")))) + "\n\t"))
            _ = v71
        }
        v1["\n\t"] = (3.14159 > 0.1)
        switch "" {
        case "hello":
            for fi72 := 0; fi72 < 19; fi72++ {
                var v73 uint16 = 32767
                _ = fi72
                _ = v73
            }
        case "foo bar":
            var v74 uint = 0
            _ = v74
        default:
            var v75 uint8 = 1
            _ = v75
        }
    } else {
        if false {
            if true {
                var v76 uint = 255
                _ = v76
            } else {
                v77 := GenS1{F0: 42, F1: 65535, F2: 180}
                fmt.Println(v9, v4)
                _ = v77
            }
        } else {
            var v78 int8 = -80
            for fi79 := 0; fi79 < 16; fi79++ {
                for fi80 := 0; fi80 < 13; fi80++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    delete(v3, (fi80 ^ 9223372036854775807))
                    _ = fi80
                }
                if (470 != 0) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    for fi81 := 0; fi81 < 8; fi81++ {
                        if len(v2) > 0 {
                            v2[0] = 128
                        }
                        var v82 uint64 = v4
                        var v83 float32 = (0.0 - ((((-2.25 / 1.0) * (1e10 * 1e10)) - -2.25) - (((0.0 - (((1e10 + 1.5) * 0.0) - ((1.5 / (1e10 + 3.14159)) - (float32(1.5) / 0.1)))) - float32(-2.25)) - 3.14159)))
                        _ = fi81
                        _ = v82
                        _ = v83
                    }
                }
                _ = fi79
            }
            _ = v78
        }
    }
    var v84 string = ("hello" + ("a" + "\n\t"))
    var v85 bool = (true || (0 <= 65535))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v84
    _ = v85
    _ = v9
}
