package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int8; F1 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint; F1 int8; F2 bool }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 uint }
func genFn0(p0 int64, p1 uint32, p2 uint32) uint32 {
    p2 -= 65535
    p1 = (^(p2))
    var v4 int16 = 1
    var v5 string = ("" + "foo bar")
    v6 := []int{42, 4294967295, int(p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return p2
}
func genFn1(p0 uint) string {
    var v7 int = -160
    v8 := GenS1{F0: (v7 % ((v7) | 1)), F1: 255, F2: 65535}
    var v9 string = ""
    var v10 int32 = -2147483648
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return ""
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint = 255
    v3 := []uint32{2147483648}
    var v11 int64 = 1
    var v12 float32 = 3.14159
    var v13 float32 = float32(1.5)
    if false {
        v2 -= v2
    } else {
        switch (("a" + ("foo bar" + (("a" + "\n\t") + "\n\t"))) + "") {
        case "":
            if len(v3) > 0 {
                v3[0] = genFn0((-(v11)), 65535, 2)
            }
        case "foo bar":
            v11 -= 42
        case "\n\t":
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 string = ("" + "hello")
                _ = fi14
                _ = v15
            }
        default:
            var v16 uint64 = 1
            _ = v16
        }
        if (((false && false) || true) && ((v13 / -2.25) <= (v13 + v12))) {
            if true {
                if (((v2 >= (v2 / ((v2) | 1))) || true) || true) {
                    var v17 int = 159
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v17
                }
            }
            v11 -= (v11 << ((42) & 63))
        } else {
            switch "foo bar" {
            case "foo bar":
                v12 += (v12 - 0.0)
            case "a":
                _ = len(v1)
            default:
                if (!false) {
                    v11 += -79
                    if false {
                        var v18 int32 = 1
                        switch 32768 {
                        case -70:
                            var v19 int64 = v11
                            _ = v19
                        case 4294967295:
                            if false {
                                v20 := GenS2{F0: true}
                                _ = v20
                            } else {
                                for fi21 := 0; fi21 < 5; fi21++ {
                                    if false {
                                        fmt.Println(v11, v11)
                                        v1 = append(v1, ("foo bar" + "\n\t"))
                                    }
                                    _ = len(v1)
                                    _ = fi21
                                }
                            }
                        }
                        _ = v18
                    }
                    v2 += 112
                }
            }
        }
    }
    var v22 int64 = (v11 + v11)
    v2 += 0
    var v23 float32 = (v13 * (v13 - -2.25))
    for fi24 := 0; fi24 < 8; fi24++ {
        switch fi24 {
        case 0:
            switch "" {
            case "\n\t":
                fmt.Println(v12)
            case "":
                v2 = (^(v2))
            default:
                for fi25 := 0; fi25 < 10; fi25++ {
                    switch -9223372036854775808 {
                    case 1:
                        fmt.Println(v11, v23, v12)
                    case -1:
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    }
                    v22 = v11
                    v13 -= ((3.14159 / ((0.0 - v13) * v12)) + 0.1)
                    _ = fi25
                }
            }
        case 32767:
            switch ("" + "foo bar") {
            case "foo bar":
                if false {
                    switch "\n\t" {
                    case "hello":
                        if len(v1) > 0 {
                            v1[0] = "hello"
                        }
                    default:
                        var v26 float64 = (float64((v13 * float32(0.0))) * 0.1)
                        _ = v26
                    }
                    if true {
                        v22 += v11
                        if false {
                            var v27 uint = 949
                            _ = v27
                        } else {
                            if true {
                                fmt.Println(v12, v11, fi24)
                                if (true && true) {
                                    if (65535 != 981) {
                                        var v28 int64 = (v22 / (((v22 - (v22 | v22))) | 1))
                                        for fi29 := 0; fi29 < 13; fi29++ {
                                            fmt.Println(v23)
                                            _ = fi29
                                        }
                                        _ = v28
                                    } else {
                                        v13 = 1.5
                                        for fi30 := 0; fi30 < 11; fi30++ {
                                            if false {
                                                if (false || false) {
                                                    v31 := GenS3{F0: 107, F1: 65535}
                                                    _ = v31
                                                } else {
                                                    v32 := GenS4{F0: false}
                                                    var v33 string = ("foo bar" + ((("" + "\n\t") + "\n\t") + "hello"))
                                                    if true {
                                                        if len(v3) > 0 {
                                                            v3[0] = genFn0(-89, 2147483647, 65535)
                                                        }
                                                        if true {
                                                            v34 := GenS3{F0: 26, F1: 683}
                                                            v35 := GenS2{F0: true}
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                        var v36 int = 32768
                                                        switch "foo bar" {
                                                        case "a":
                                                            if (v33 == v33) {
                                                                var v37 int16 = 42
                                                                fmt.Println(v2, v22)
                                                                _ = v37
                                                            } else {
                                                                fmt.Println(v12, v2, v2)
                                                                v1 = append(v1, v33)
                                                            }
                                                        case "\n\t":
                                                            v38 := GenS5{F0: 4294967295, F1: false}
                                                            _ = v38
                                                        default:
                                                            fmt.Println(fi24, v2)
                                                        }
                                                        _ = v36
                                                    } else {
                                                        var v39 uint32 = 42
                                                        _ = len(v3)
                                                        _ = v39
                                                    }
                                                    _ = v32
                                                    _ = v33
                                                }
                                                if true {
                                                    v23 = ((v12 / float32(1e10)) * v23)
                                                    if false {
                                                        switch "" {
                                                        case "hello":
                                                            if false {
                                                                if true {
                                                                    v23 = (v12 + (((1e10 + 1.5) * (v23 * v12)) + (float32((((((0.0 + ((3.14159 - (3.14159 + ((0.0 / 3.14159) / (1e10 * 3.14159)))) + 0.0)) + float64(v13)) / 1e10) * 1e10) * 0.1)) / 1.0)))
                                                                } else {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = ((("\n\t" + "foo bar") + "hello") + ("foo bar" + (((("" + "\n\t") + "a") + ("\n\t" + "foo bar")) + ("foo bar" + "a"))))
                                                                    }
                                                                    var v40 bool = (v12 != v13)
                                                                    _ = v40
                                                                }
                                                            } else {
                                                                v41 := GenS6{F0: 100, F1: (("\n\t" + "hello") + "foo bar")}
                                                                v22 = v11
                                                                switch 42 {
                                                                case 128:
                                                                    v23 = 0.0
                                                                }
                                                                switch 752 {
                                                                case -1:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 0
                                                                    }
                                                                }
                                                                _ = v41
                                                            }
                                                        case "":
                                                            v13 += 0.0
                                                        }
                                                        v42 := GenS7{F0: v11}
                                                        v3 = append(v3, 2147483648)
                                                        _ = v42
                                                    }
                                                    fmt.Println(v23)
                                                } else {
                                                    fmt.Println(v13)
                                                }
                                            }
                                            var v43 uint32 = 438
                                            _ = fi30
                                            _ = v43
                                        }
                                        switch ((("a" + "\n\t") + "hello") + ("" + "foo bar")) {
                                        case "hello":
                                            v2 = v2
                                        case "a":
                                            if (true || true) {
                                                v11 = v22
                                                if len(v3) > 0 {
                                                    v3[0] = 20
                                                }
                                                var v44 uint = (v2 >> (((v2 >> ((v2) & 63))) & 63))
                                                _ = v44
                                            }
                                        case "foo bar":
                                            for fi45 := 0; fi45 < 17; fi45++ {
                                                var v46 uint8 = 255
                                                if false {
                                                    var v47 float32 = 1e10
                                                    if false {
                                                        var v48 int = (fi24 + fi45)
                                                        _ = v48
                                                    }
                                                    _ = v47
                                                } else {
                                                    v49 := GenS8{F0: v2, F1: 7, F2: true}
                                                    if len(v3) > 0 {
                                                        v3[0] = 2
                                                    }
                                                    switch (fi45 * (fi45 % ((fi24) | 1))) {
                                                    case 32767:
                                                        _ = len(v1)
                                                    case -941:
                                                        v22 -= 1
                                                    default:
                                                        var v50 uint16 = 255
                                                        v11 = (v22 >> (((v2 & (v2 + (v2 << (((v2 >> ((uint(v50)) & 63))) & 63))))) & 63))
                                                        v22 = 934
                                                        _ = v50
                                                    }
                                                    switch "\n\t" {
                                                    case "a":
                                                        v13 = v13
                                                    }
                                                    _ = v49
                                                }
                                                _ = fi45
                                                _ = v46
                                            }
                                        }
                                    }
                                }
                                v2 = v2
                            }
                            if ((false && true) && (((v22 ^ 799) == (^(v11))) || true)) {
                                var v51 uint8 = 155
                                _ = v51
                            } else {
                                for fi52 := 0; fi52 < 3; fi52++ {
                                    v13 += float32(0.0)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v53 := GenS9{F0: 0, F1: 42}
                                    _ = fi52
                                    _ = v53
                                }
                                v11 = (v11 - v22)
                                var v54 int16 = 1000
                                _ = v54
                            }
                            fmt.Println(v23)
                        }
                        if false {
                            for fi55 := 0; fi55 < 2; fi55++ {
                                if true {
                                    if (0.0 >= 3.14159) {
                                        v22 = 9223372036854775807
                                    } else {
                                        v2 = v2
                                    }
                                    switch (-(fi55)) {
                                    case -9223372036854775808:
                                        var v56 uint64 = 1
                                        _ = v56
                                    case 9223372036854775807:
                                        v11 = (v11 | (v11 << ((v2) & 63)))
                                    default:
                                        var v57 bool = false
                                        _ = v57
                                    }
                                    v58 := GenS10{F0: (v2 % ((536) | 1))}
                                    _ = v58
                                }
                                v23 -= ((1e10 - v23) + (((v13 + 0.0) * v12) + v12))
                                _ = fi55
                            }
                        }
                    }
                    switch (fi24 & fi24) {
                    case 42:
                        var v59 int32 = 42
                        _ = v59
                    case 1:
                        v13 = 0.0
                    }
                }
            case "a":
                v13 = ((-2.25 / float32(-2.25)) / 1.0)
            case "\n\t":
                for fi60 := 0; fi60 < 8; fi60++ {
                    v61 := GenS1{F0: (fi60 ^ (^(fi60))), F1: 107, F2: 65535}
                    fmt.Println(fi24, v12)
                    fmt.Println(v22, v22)
                    if false {
                        fmt.Println(v11, v13, v11)
                        var v62 int16 = 345
                        _ = v62
                    }
                    var v63 int8 = -1
                    _ = fi60
                    _ = v61
                    _ = v63
                }
            }
        case 4294967295:
            if false {
                v64 := GenS9{F0: -15, F1: 42}
                _ = len(v3)
                _ = v64
            } else {
                var v65 string = "hello"
                var v66 int = (fi24 >> ((0) & 63))
                _ = v65
                _ = v66
            }
        }
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        _ = fi24
    }
    for fi67 := 0; fi67 < 20; fi67++ {
        var v68 float32 = (-2.25 / v13)
        _ = fi67
        _ = v68
    }
    var v69 int64 = v22
    var v70 uint8 = 42
    _ = len(v3)
    if (65535 != v2) {
        for fi71 := 0; fi71 < 20; fi71++ {
            for fi72 := 0; fi72 < 10; fi72++ {
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v22)
                case "foo bar":
                    v22 = -1
                }
                _ = fi72
            }
            _ = fi71
        }
    }
    v12 -= float32(((-2.25 / ((float64(3.14159) + float64(1.5)) - -2.25)) - (0.1 - 1.5)))
    var v73 uint64 = 18446744073709551615
    v74 := GenS3{F0: -51, F1: 65535}
    if true {
        fmt.Println(v23)
        v75 := GenS3{F0: -84, F1: 2147483648}
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        _ = v75
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v69
    _ = v70
    _ = v73
    _ = v74
}
