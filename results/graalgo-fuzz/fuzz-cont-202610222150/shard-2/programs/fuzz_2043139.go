package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int32; F2 string }
type GenS4 struct { F0 uint; F1 uint8; F2 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string; F1 uint8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int16; F1 bool; F2 uint8 }
type GenS10 struct { F0 int16; F1 int64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 string; F1 uint32; F2 int64 }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 string; F1 uint16 }

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[string]string)
    var v3 uint16 = 608
    v4 := []int{v1}
    var v5 int32 = 1
    var v6 int = (v1 - (v1 | (v1 % ((-1) | 1))))
    var v7 int8 = 127
    if true {
        fmt.Println(v6, v6, v5)
    }
    if false {
        var v8 string = "\n\t"
        v9 := GenS1{F0: 4294967295}
        _ = v8
        _ = v9
    } else {
        v10 := GenS2{F0: 42}
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS1{F0: 767}
            v13 := GenS3{F0: 210, F1: (v5 | 2147483647), F2: ""}
            v7 = (-(v7))
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 float32 = (0.0 / 0.1)
            if ((0.1 * (v15 / float32(float64(1.5)))) != (v15 / (v15 / ((((v15 / v15) * 1.5) * 0.0) + v15)))) {
                switch (v6 % (((v1 | -1)) | 1)) {
                case -954:
                    for fi16 := 0; fi16 < 15; fi16++ {
                        switch -894 {
                        case 0:
                            if (true || false) {
                                v5 += (v5 | (v5 ^ v5))
                            }
                        case 42:
                            var v17 string = ("foo bar" + "hello")
                            _ = v17
                        case -9223372036854775808:
                            var v18 int = 9223372036854775807
                            _ = v18
                        }
                        switch ("\n\t" + ((("\n\t" + ("\n\t" + ("hello" + "a"))) + "foo bar") + "")) {
                        case "foo bar":
                            if false {
                                if (!(v7 < v7)) {
                                    v7 -= 1
                                }
                                v5 = -112
                            } else {
                                var v19 uint64 = 2
                                fmt.Println(v7, fi16, v3)
                                v20 := v2[("hello" + "a")]
                                _ = v19
                                _ = v20
                            }
                        case "\n\t":
                            if false {
                                var v21 float32 = ((v15 + float32(float64(((0.0 * (v15 * v15)) * v15)))) + float32(0.1))
                                fmt.Println(v1)
                                if (-257 > -32768) {
                                    switch int(v5) {
                                    case 829:
                                        v1 -= v1
                                    case -34:
                                        fmt.Println(v1, v21)
                                    case 2147483647:
                                        switch (fi16 + (fi16 - (^(v6)))) {
                                        case 4294967295:
                                            if false {
                                                v22 := GenS2{F0: 156}
                                                v7 -= (v7 | (^(v7)))
                                                v23 := GenS2{F0: 18446744073709551615}
                                                _ = v22
                                                _ = v23
                                            }
                                        case 76:
                                            v24 := GenS4{F0: 18446744073709551615, F1: 105, F2: true}
                                            _ = v24
                                        case -496:
                                            v21 -= float32(((-2.25 * -2.25) * 3.14159))
                                        default:
                                            var v25 uint16 = 63
                                            _ = v25
                                        }
                                    }
                                    v1 -= v1
                                    var v26 uint = 552
                                    _ = v26
                                } else {
                                    if true {
                                        var v27 int8 = (v7 * v7)
                                        _ = v27
                                    } else {
                                        v4 = append(v4, v1)
                                        v1 -= v6
                                    }
                                    v2["a"] = "foo bar"
                                }
                                _ = v21
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = fi16
                                }
                                var v28 uint64 = 42
                                v29 := GenS5{F0: 1}
                                _ = v28
                                _ = v29
                            }
                            if (881 < 19) {
                                v7 = (v7 << ((892) & 63))
                            } else {
                                v6 += (fi14 % ((fi16) | 1))
                            }
                            _ = len(v2)
                            var v30 uint16 = (v3 | 1)
                            v31 := GenS6{F0: 42}
                            _ = v30
                            _ = v31
                        case "a":
                            if ((("a" + "") <= ("hello" + (("hello" + ("\n\t" + "a")) + "\n\t"))) && (!false)) {
                                v6 -= (v1 - v1)
                                if true {
                                    switch (v6 / ((0) | 1)) {
                                    case 617:
                                        fmt.Println(fi16, v7, v3)
                                    case 870:
                                        fmt.Println(fi14, v3, v7)
                                    default:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v1 = (v1 ^ 366)
                                        var v32 uint32 = 4294967295
                                        switch ("foo bar" + ("" + "")) {
                                        case "\n\t":
                                            v32 += 34
                                        case "":
                                            fmt.Println(v5, v3, fi14)
                                        case "foo bar":
                                            if false {
                                                v6 += (-(v1))
                                                v33 := GenS4{F0: 1, F1: 127, F2: (!true)}
                                                v34 := GenS4{F0: 0, F1: 130, F2: (((v5 - v5) < v5) || (!false))}
                                                _ = v33
                                                _ = v34
                                            }
                                        }
                                        fmt.Println(v7)
                                        _ = v32
                                    }
                                    fmt.Println(v6)
                                }
                                fmt.Println(v15, v1)
                            }
                        default:
                            var v35 uint = 2
                            _ = v35
                        }
                        _ = fi16
                    }
                }
                v15 -= float32(3.14159)
            }
            _ = fi14
            _ = v15
        }
        _ = v10
    }
    v36 := GenS2{F0: 1}
    v4 = append(v4, v1)
    for fi37 := 0; fi37 < 17; fi37++ {
        if true {
            if true {
                v38 := GenS7{F0: ("\n\t" + "a"), F1: 255}
                _ = v38
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch "" {
            case "\n\t":
                v39 := v2[""]
                _ = v39
            case "a":
                var v40 int16 = 1
                _ = v40
            case "hello":
                v6 = (v1 - v1)
            }
        }
        v41 := GenS8{F0: 42}
        v42 := GenS9{F0: 32767, F1: true, F2: 187}
        _ = fi37
        _ = v41
        _ = v42
    }
    v43 := GenS10{F0: 1000, F1: 0}
    var v44 uint = 2
    fmt.Println(v5)
    for fi45 := 0; fi45 < 18; fi45++ {
        if true {
            var v46 uint16 = (v3 + 1)
            if true {
                v3 = (v3 * 1)
            } else {
                fmt.Println(v6, v7)
            }
            v1 = v6
            _ = v46
        } else {
            if true {
                v7 -= v7
                for fi47 := 0; fi47 < 19; fi47++ {
                    if (float32(1.5) > (0.0 * float32(float64((1.5 - ((0.0 / (3.14159 / 1.0)) - (3.14159 - -2.25))))))) {
                        v2[("" + "")] = "hello"
                        v3 -= (^(v3))
                        if true {
                            v5 += (-(v5))
                            fmt.Println(fi45)
                        } else {
                            var v48 int32 = -2147483648
                            v4 = append(v4, (fi45 + (fi47 - v6)))
                            _ = v48
                        }
                    } else {
                        if ((1 != 69) && (v7 != 42)) {
                            if true {
                                v7 -= (^(v7))
                                fmt.Println(v6, v6, v3)
                            } else {
                                var v49 uint64 = 4294967295
                                fmt.Println(v44, v1, fi47)
                                _ = v49
                            }
                        } else {
                            v50 := GenS2{F0: 0}
                            _ = v50
                        }
                    }
                    v51 := GenS11{F0: (v1 | 9223372036854775807)}
                    v4 = append(v4, int(v7))
                    _ = fi47
                    _ = v51
                }
            }
            fmt.Println(v7, v6, v44)
        }
        v4 = append(v4, -9223372036854775808)
        v44 -= (v44 + v44)
        _ = fi45
    }
    switch (v1 | (v1 & (v6 * 32767))) {
    case 42:
        v52 := GenS3{F0: 469, F1: int32(v44), F2: "foo bar"}
        _ = v52
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    switch ("foo bar" + (("\n\t" + "") + ("\n\t" + ""))) {
    case "a":
        switch "hello" {
        case "a":
            v53 := GenS2{F0: uint64(v44)}
            _ = v53
        case "\n\t":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "foo bar":
            switch v6 {
            case -1:
                v54 := GenS9{F0: 1, F1: (true && (-2.25 <= 0.1)), F2: 127}
                _ = v54
            case 4294967295:
                switch v6 {
                case 736:
                    var v55 uint = (^(v44))
                    _ = v55
                case 2147483647:
                    v56 := GenS10{F0: -256, F1: 0}
                    _ = v56
                default:
                    v2[("foo bar" + "foo bar")] = ("foo bar" + "")
                }
            case 128:
                var v57 uint8 = 127
                _ = v57
            default:
                v6 += (^(v1))
            }
        default:
            var v58 int16 = 0
            _ = v58
        }
    case "hello":
        fmt.Println(v7, v3)
    case "":
        if (0.1 <= (-2.25 / 0.1)) {
            fmt.Println(v7)
            switch "a" {
            case "\n\t":
                if true {
                    var v59 uint = v44
                    _ = v59
                } else {
                    var v60 int64 = 1
                    _ = v60
                }
            case "foo bar":
                for fi61 := 0; fi61 < 5; fi61++ {
                    var v62 int16 = -847
                    if true {
                        for fi63 := 0; fi63 < 15; fi63++ {
                            v64 := GenS12{F0: "a", F1: 1, F2: -9223372036854775808}
                            _ = fi63
                            _ = v64
                        }
                        switch v6 {
                        case -19:
                            v3 += v3
                        case -563:
                            for fi65 := 0; fi65 < 17; fi65++ {
                                var v66 int32 = -1
                                _ = fi65
                                _ = v66
                            }
                        case -2147483648:
                            v62 += (v62 | 32767)
                        }
                    }
                    var v67 uint32 = 1
                    _ = fi61
                    _ = v62
                    _ = v67
                }
            case "a":
                var v68 uint8 = 106
                _ = v68
            default:
                for fi69 := 0; fi69 < 19; fi69++ {
                    var v70 int = 128
                    v71 := GenS13{F0: 128}
                    _ = fi69
                    _ = v70
                    _ = v71
                }
            }
        } else {
            v72 := GenS14{F0: "\n\t", F1: (v3 * 1)}
            if false {
                var v73 float32 = (-2.25 / -2.25)
                if len(v4) > 0 {
                    v4[0] = int(v7)
                }
                if ((v5 <= v5) && (!false)) {
                    v7 += (v7 | (v7 >> ((42) & 63)))
                    v73 = ((v73 + v73) + v73)
                }
                _ = v73
            } else {
                fmt.Println(v44, v44, v6)
                for fi74 := 0; fi74 < 2; fi74++ {
                    var v75 int8 = v7
                    _ = fi74
                    _ = v75
                }
                v6 += v1
            }
            _ = v72
        }
    default:
        fmt.Println(v7, v3)
    }
    switch (v1 & v6) {
    case 32768:
        var v76 int8 = v7
        _ = v76
    default:
        var v77 float32 = 3.14159
        _ = v77
    }
    _ = len(v4)
    if (("a" < ("hello" + ("hello" + ("hello" + (("\n\t" + "\n\t") + "foo bar"))))) && false) {
        var v78 int32 = v5
        _ = v78
    } else {
        fmt.Println(v3, v6)
        v44 = (v44 * v44)
    }
    var v79 uint = (^(v44))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v79
}
