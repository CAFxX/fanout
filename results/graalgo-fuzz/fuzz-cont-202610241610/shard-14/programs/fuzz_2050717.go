package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 string; F1 uint16; F2 string }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint16; F1 bool; F2 int16 }
type GenS6 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS7 struct { F0 string; F1 bool; F2 int32 }
type GenS8 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS9 struct { F0 uint8; F1 string; F2 uint64 }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 int16; F1 uint8; F2 string }
func genFn0(p0 uint16, p1 string) uint8 {
    p1 = "foo bar"
    v5 := GenS1{F0: p1, F1: -62}
    if true {
        switch -9223372036854775808 {
        case 9223372036854775807:
            v6 := GenS1{F0: p1, F1: -128}
            _ = v6
        case 232:
            p1 = p1
        }
        v7 := []int{782, 928}
        switch 32767 {
        case 296:
            _ = len(v7)
        case -457:
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case 0:
            for fi8 := 0; fi8 < 13; fi8++ {
                p0 = (p0 >> ((288) & 63))
                v9 := GenS1{F0: p1, F1: 127}
                _ = fi8
                _ = v9
            }
        default:
            p1 = "\n\t"
        }
        _ = v7
    } else {
        var v10 float64 = 3.14159
        p0 = (p0 % (((p0 % (((p0 ^ p0)) | 1))) | 1))
        var v11 float64 = v10
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v5
    return 255
}

func main() {
    var v1 uint8 = 62
    var v2 int32 = 0
    var v3 uint32 = 695
    v4 := []uint32{(^(v3)), (v3 * (v3 * v3))}
    var v12 float32 = ((-2.25 * ((0.1 / ((1e10 - ((3.14159 / (1.5 * ((1.5 * (((-2.25 + -2.25) - -2.25) * 1e10)) / (((-2.25 / (((-2.25 / 1.0) / (float32(1e10) / 3.14159)) / 1e10)) - 3.14159) + 1.5)))) + 0.0)) / 1.5)) * (0.1 - 0.1))) / 0.1)
    v13 := GenS2{F0: 255, F1: 100}
    if (!(1e10 == (1e10 + (-2.25 * 1e10)))) {
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS2{F0: 42, F1: 0}
            if (float64(((v12 + (v12 + v12)) * 3.14159)) == float64((0.1 - ((v12 / v12) - (v12 * float32(1e10)))))) {
                if (32767 >= 1) {
                    v3 += 42
                } else {
                    v16 := GenS3{F0: "foo bar", F1: 255, F2: (("\n\t" + "\n\t") + "\n\t")}
                    for fi17 := 0; fi17 < 10; fi17++ {
                        var v18 uint64 = 965
                        switch ("" + ((("a" + "\n\t") + "") + ("hello" + "hello"))) {
                        case "hello":
                            v12 += 0.1
                        case "\n\t":
                            switch "hello" {
                            case "a":
                                v19 := GenS3{F0: ("foo bar" + "foo bar"), F1: 0, F2: "\n\t"}
                                _ = v19
                            case "hello":
                                v20 := GenS3{F0: "\n\t", F1: 265, F2: "hello"}
                                _ = v20
                            case "\n\t":
                                var v21 string = "\n\t"
                                _ = v21
                            }
                        case "":
                            var v22 int32 = (-(v2))
                            v1 = (v1 >> ((4294967295) & 63))
                            _ = v22
                        default:
                            if (!(42 == 0)) {
                                v23 := GenS2{F0: uint(v1), F1: -1}
                                _ = v23
                            } else {
                                v24 := GenS4{F0: (^(v18)), F1: (true || true)}
                                _ = v24
                            }
                        }
                        switch (fi17 * fi17) {
                        case -943:
                            if (-2147483648 == (v2 / ((v2) | 1))) {
                                v2 += v2
                                switch (^(fi14)) {
                                case 42:
                                    fmt.Println(v1)
                                case -776:
                                    v25 := GenS2{F0: uint(v3), F1: 0}
                                    _ = v25
                                case 1:
                                    if false {
                                        if ((true && (!true)) && false) {
                                            fmt.Println(v1, v2, v18)
                                        } else {
                                            var v26 int8 = 87
                                            _ = len(v4)
                                            _ = v26
                                        }
                                        v12 += v12
                                        v3 -= v3
                                    }
                                }
                                if false {
                                    v2 -= (v2 >> ((905) & 63))
                                    if false {
                                        fmt.Println(v1)
                                    } else {
                                        var v27 int64 = 42
                                        fmt.Println(v3, v18, v1)
                                        switch "\n\t" {
                                        case "a":
                                            _ = len(v4)
                                        default:
                                            v28 := GenS3{F0: "\n\t", F1: 1, F2: "foo bar"}
                                            _ = v28
                                        }
                                        _ = v27
                                    }
                                    v29 := GenS4{F0: v18, F1: ((((("a" + ("foo bar" + "foo bar")) + ("" + ("foo bar" + ""))) + ("foo bar" + "a")) + ("\n\t" + (("hello" + ("\n\t" + "")) + (("foo bar" + "\n\t") + ("hello" + "\n\t"))))) == "foo bar")}
                                    _ = v29
                                }
                                fmt.Println(v3, v18, fi14)
                                if false {
                                    switch "" {
                                    case "hello":
                                        var v30 int64 = 9223372036854775807
                                        _ = v30
                                    case "":
                                        var v31 uint64 = uint64(v3)
                                        _ = v31
                                    case "\n\t":
                                        v1 = (v1 * 128)
                                    }
                                    if true {
                                        v3 = 494
                                    } else {
                                        if (genFn0(42, "") > (v1 / ((255) | 1))) {
                                            var v32 int32 = -808
                                            v33 := GenS1{F0: "\n\t", F1: -90}
                                            var v34 uint32 = (v3 * (^(v3)))
                                            _ = v32
                                            _ = v33
                                            _ = v34
                                        } else {
                                            var v35 int16 = 1000
                                            _ = v35
                                        }
                                        v12 = float32(0.0)
                                        var v36 float64 = (-2.25 + 1.5)
                                        _ = v36
                                    }
                                } else {
                                    if true {
                                        v4 = append(v4, v3)
                                    }
                                    v37 := GenS5{F0: 255, F1: (true || ((-512 < 32767) || (true || false))), F2: -446}
                                    v12 = 3.14159
                                    _ = v37
                                }
                            }
                        case 4294967295:
                            switch "hello" {
                            case "foo bar":
                                v3 -= (v3 % ((uint32(v18)) | 1))
                            default:
                                if len(v4) > 0 {
                                    v4[0] = v3
                                }
                            }
                        case -18:
                            var v38 string = "\n\t"
                            v38 = v38
                            v1 -= v1
                            _ = v38
                        }
                        _ = fi17
                        _ = v18
                    }
                    _ = v16
                }
                v2 += v2
                for fi39 := 0; fi39 < 13; fi39++ {
                    if false {
                        fmt.Println(v3, v12)
                    }
                    var v40 uint8 = (v1 - (v1 & (v1 - v1)))
                    _ = fi39
                    _ = v40
                }
            } else {
                v1 -= genFn0(0, ("hello" + (("hello" + "foo bar") + "")))
                var v41 float32 = 1.5
                var v42 int16 = 1000
                _ = v41
                _ = v42
            }
            fmt.Println(v2)
            _ = fi14
            _ = v15
        }
        for fi43 := 0; fi43 < 6; fi43++ {
            var v44 string = (("foo bar" + "foo bar") + ("hello" + "foo bar"))
            switch 32767 {
            case 0:
                if false {
                    v45 := GenS5{F0: 65535, F1: true, F2: 0}
                    _ = v45
                } else {
                    if (!false) {
                        var v46 int64 = -271
                        for fi47 := 0; fi47 < 18; fi47++ {
                            v48 := GenS6{F0: -1, F1: v3, F2: 65535}
                            var v49 int16 = 1
                            if (v2 == v2) {
                                if (true && false) {
                                    var v50 float32 = (1.5 * ((1.5 * float32(0.0)) + -2.25))
                                    _ = v50
                                }
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = (v3 & v3)
                                }
                                v2 = (v2 >> ((uint(v1)) & 63))
                                fmt.Println(v2)
                            }
                            _ = fi47
                            _ = v48
                            _ = v49
                        }
                        _ = v46
                    }
                    var v51 uint64 = 4294967295
                    v4 = append(v4, v3)
                    _ = v51
                }
            case 127:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                var v52 uint = 533
                _ = v52
            }
            _ = fi43
            _ = v44
        }
        v53 := GenS7{F0: "foo bar", F1: (2 == 2), F2: (v2 - v2)}
        _ = v53
    } else {
        switch 4294967295 {
        case 42:
            for fi54 := 0; fi54 < 14; fi54++ {
                _ = len(v4)
                _ = len(v4)
                v55 := GenS5{F0: 2, F1: (!(false || false)), F2: 0}
                _ = fi54
                _ = v55
            }
        default:
            if (216 < 979) {
                v56 := GenS8{F0: uint16(v3), F1: true, F2: 4294967295}
                _ = v56
            }
        }
        var v57 uint8 = (v1 + (^(v1)))
        _ = v57
    }
    if false {
        v58 := GenS1{F0: (((((("a" + ("foo bar" + "")) + ((((((("foo bar" + "hello") + "\n\t") + "a") + ("\n\t" + "\n\t")) + "foo bar") + "hello") + "a")) + "hello") + "a") + (("a" + "") + (("\n\t" + "hello") + "foo bar"))) + ("foo bar" + ((((("a" + "hello") + ("foo bar" + "")) + "hello") + "\n\t") + "a"))), F1: -128}
        if false {
            v59 := GenS8{F0: 0, F1: false, F2: 42}
            v2 += (v2 ^ (v2 & v2))
            var v60 uint64 = 299
            _ = v59
            _ = v60
        } else {
            v3 = v3
        }
        v12 += (((v12 - v12) / v12) / 1.0)
        _ = v58
    } else {
        var v61 uint = 733
        for fi62 := 0; fi62 < 6; fi62++ {
            v63 := GenS5{F0: 786, F1: (false && (false && ((0.0 - (float64(v12) * 1e10)) != ((3.14159 + -2.25) / 1.5)))), F2: 1}
            fmt.Println(v12, v1)
            _ = fi62
            _ = v63
        }
        _ = v61
    }
    if false {
        var v64 int = -785
        for fi65 := 0; fi65 < 6; fi65++ {
            var v66 uint = 692
            _ = fi65
            _ = v66
        }
        _ = v64
    } else {
        switch ("a" + ("foo bar" + "")) {
        case "hello":
            if (((0.0 * 1e10) * (((1e10 * 1e10) + float32(-2.25)) * (v12 / v12))) <= v12) {
                v67 := GenS8{F0: 650, F1: (752 == 18446744073709551615), F2: 18446744073709551615}
                var v68 float32 = v12
                if (float64(v12) == 0.0) {
                    var v69 uint = 65535
                    v68 += 3.14159
                    _ = v69
                }
                _ = v67
                _ = v68
            }
        }
        for fi70 := 0; fi70 < 13; fi70++ {
            v71 := GenS9{F0: v1, F1: "hello", F2: 0}
            v2 += v2
            v1 += (v1 >> ((2) & 63))
            _ = fi70
            _ = v71
        }
    }
    if (65535 <= 311) {
        v1 -= (v1 & genFn0(2, "foo bar"))
        var v72 int16 = 515
        for fi73 := 0; fi73 < 12; fi73++ {
            fmt.Println(v1, fi73)
            _ = fi73
        }
        _ = v72
    }
    var v74 uint64 = 0
    if (!false) {
        v75 := GenS10{F0: -601, F1: 42}
        v74 += 0
        _ = v75
    } else {
        for fi76 := 0; fi76 < 9; fi76++ {
            v12 = 3.14159
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if false {
                var v77 uint64 = 0
                _ = v77
            } else {
                v74 = 1
                if (0 >= 761) {
                    var v78 int16 = int16(v2)
                    for fi79 := 0; fi79 < 15; fi79++ {
                        v12 += (float32(0.1) / v12)
                        _ = fi79
                    }
                    v3 = uint32(v1)
                    _ = v78
                }
                fmt.Println(fi76, v2)
            }
            _ = fi76
        }
        v80 := GenS11{F0: -1, F1: (v1 - (v1 >> ((uint(v3)) & 63))), F2: "hello"}
        v74 -= v74
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v74
}
