package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int32; F2 uint16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint64; F1 int64; F2 int8 }
type GenS6 struct { F0 int; F1 string; F2 int16 }
type GenS7 struct { F0 int64; F1 int16 }
func genFn0(p0 uint8, p1 string) uint16 {
    p1 = "foo bar"
    if (1 == 0) {
        var v7 uint16 = 1
        if false {
            p1 = ((p1 + (p1 + "foo bar")) + p1)
        } else {
            v7 += 282
        }
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        if (42 >= 18446744073709551615) {
            fmt.Println(p0, p0, fi8)
            p0 += (^(p0))
        } else {
            p0 -= p0
        }
        fmt.Println(p0)
        v9 := GenS1{F0: p1}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 string, p1 bool) int16 {
    var v10 uint64 = 1
    var v11 int64 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 32767
}

func main() {
    var v1 float32 = ((((0.1 - 1.5) / 1e10) * ((1.5 / (0.1 / 1.0)) + (1e10 / ((1.5 - 0.0) * 1e10)))) / 0.1)
    var v2 int16 = -32768
    var v3 uint = 42
    var v4 uint32 = 127
    var v5 uint32 = (v4 << (((v3 >> ((0) & 63))) & 63))
    var v6 uint8 = 120
    v2 += v2
    v12 := GenS1{F0: "foo bar"}
    v13 := []int{1, 127}
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS2{F0: (fi14 + fi14), F1: 42, F2: genFn0(v6, "")}
        fmt.Println(v4)
        if len(v13) > 0 {
            v13[0] = (fi14 >> (((v3 << (((v3 | 586)) & 63))) & 63))
        }
        _ = fi14
        _ = v15
    }
    var v16 string = ""
    v2 = v2
    if ((!false) && true) {
        fmt.Println(v16, v16)
        v17 := GenS3{F0: v16, F1: v3}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 20; fi18++ {
            switch fi18 {
            case 4294967295:
                var v19 uint16 = 32767
                _ = v19
            case 663:
                v2 = -32768
            case 128:
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            }
            switch fi18 {
            case 42:
                v20 := GenS4{F0: v2}
                _ = v20
            case -413:
                v6 -= (^(v6))
            }
            _ = fi18
        }
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v6 = uint8(v3)
        switch (v16 + (("foo bar" + (("hello" + v16) + v16)) + ("a" + "a"))) {
        case "\n\t":
            _ = len(v13)
        case "a":
            v22 := GenS3{F0: ("foo bar" + v16), F1: 42}
            _ = v22
        case "foo bar":
            for fi23 := 0; fi23 < 16; fi23++ {
                if (65535 != v3) {
                    var v24 uint8 = v6
                    fmt.Println(v4, v24, fi23)
                    _ = v24
                } else {
                    if (!false) {
                        fmt.Println(v16)
                    }
                    switch v16 {
                    case "foo bar":
                        var v25 int16 = int16(v3)
                        _ = v25
                    }
                }
                var v26 uint64 = 18446744073709551615
                _ = fi23
                _ = v26
            }
        default:
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 string = (v16 + v16)
                var v29 float32 = v1
                v30 := GenS4{F0: int16(v5)}
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        }
        for fi31 := 0; fi31 < 8; fi31++ {
            fmt.Println(v4, v4)
            switch "hello" {
            case "hello":
                var v32 int32 = 2147483647
                _ = v32
            case "":
                fmt.Println(v1, v6, fi21)
            case "a":
                var v33 int16 = genFn1(("hello" + v16), false)
                v4 -= 836
                _ = v33
            default:
                if len(v13) > 0 {
                    v13[0] = int(v6)
                }
            }
            var v34 float32 = ((v1 / 1e10) - (v1 - (-2.25 * v1)))
            _ = fi31
            _ = v34
        }
        _ = fi21
    }
    v35 := GenS5{F0: 4294967295, F1: -9223372036854775808, F2: 100}
    if (2 >= 65535) {
        var v36 uint32 = (v4 >> ((v3) & 63))
        _ = v36
    }
    if (genFn1(v16, ("foo bar" < "hello")) < v2) {
        v6 += (^(v6))
        _ = len(v13)
    }
    for fi37 := 0; fi37 < 6; fi37++ {
        if (int32(v6) <= -327) {
            v38 := GenS5{F0: 271, F1: -1, F2: 0}
            if (false && false) {
                v5 = 302
                switch int(v6) {
                case 4294967295:
                    if false {
                        for fi39 := 0; fi39 < 1; fi39++ {
                            var v40 int32 = 0
                            if (65535 < 833) {
                                v1 = v1
                                var v41 int16 = (v2 >> ((v3) & 63))
                                if true {
                                    fmt.Println(v1)
                                    fmt.Println(fi39)
                                    v16 = "foo bar"
                                } else {
                                    switch fi39 {
                                    case -763:
                                        fmt.Println(v6, v40, v3)
                                    }
                                    var v42 uint8 = 128
                                    _ = v42
                                }
                                _ = v41
                            }
                            v1 = (3.14159 / (v1 + v1))
                            _ = fi39
                            _ = v40
                        }
                    } else {
                        for fi43 := 0; fi43 < 12; fi43++ {
                            v6 = (v6 - (v6 % (((v6 % ((v6) | 1))) | 1)))
                            if (1 > 42) {
                                if (55 == v6) {
                                    v44 := GenS6{F0: (^(fi37)), F1: "foo bar", F2: (v2 - genFn1(("" + "a"), ((!(223 < genFn0((v6 << ((562) & 63)), v16))) || (v1 > v1))))}
                                    _ = v44
                                }
                            }
                            _ = fi43
                        }
                        for fi45 := 0; fi45 < 13; fi45++ {
                            if false {
                                v6 -= v6
                                v2 = (^(v2))
                            } else {
                                v4 += 500
                                if true {
                                    if (2 < 4294967295) {
                                        v6 = v6
                                    } else {
                                        var v46 uint = (v3 >> ((1) & 63))
                                        _ = v46
                                    }
                                    var v47 int64 = 9223372036854775807
                                    _ = v47
                                } else {
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    v13 = append(v13, (fi45 << ((v3) & 63)))
                                }
                                v4 = (v4 + 65535)
                            }
                            if true {
                                v48 := GenS7{F0: 0, F1: genFn1(v16, true)}
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                _ = len(v13)
                                v16 = v16
                                _ = v48
                            } else {
                                var v49 uint32 = 2147483648
                                var v50 uint8 = v6
                                _ = v49
                                _ = v50
                            }
                            _ = fi45
                        }
                    }
                case 9223372036854775807:
                    v6 += uint8(v5)
                }
                v6 = v6
            } else {
                switch (-(fi37)) {
                case 4294967295:
                    if ((0.1 >= (v1 - float32(0.1))) || true) {
                        var v51 uint = v3
                        _ = v51
                    } else {
                        v16 = "a"
                    }
                }
                if false {
                    if true {
                        switch (^(fi37)) {
                        case 127:
                            for fi52 := 0; fi52 < 5; fi52++ {
                                v53 := GenS2{F0: fi52, F1: -930, F2: 0}
                                _ = fi52
                                _ = v53
                            }
                        case 32767:
                            if len(v13) > 0 {
                                v13[0] = (fi37 & 127)
                            }
                        default:
                            if len(v13) > 0 {
                                v13[0] = -262
                            }
                        }
                        for fi54 := 0; fi54 < 7; fi54++ {
                            switch (fi54 & (fi37 - fi54)) {
                            case 130:
                                v3 += uint(v4)
                            case 4294967295:
                                var v55 int32 = 0
                                _ = v55
                            case -70:
                                if true {
                                    v13 = append(v13, fi54)
                                    v13 = append(v13, 9223372036854775807)
                                }
                            default:
                                switch "a" {
                                case "hello":
                                    v56 := GenS4{F0: v2}
                                    _ = v56
                                case "a":
                                    var v57 float32 = 0.0
                                    _ = v57
                                }
                            }
                            fmt.Println(v6)
                            var v58 string = v16
                            _ = fi54
                            _ = v58
                        }
                        if (false && ((581 > genFn0(v6, v16)) && false)) {
                            v3 -= (v3 + v3)
                            var v59 uint8 = (v6 + 2)
                            _ = v59
                        }
                    } else {
                        var v60 int16 = v2
                        if (((((float64((((v1 + 1.5) * v1) * float32(1e10))) * (1e10 - 1.5)) * -2.25) / ((-2.25 - 3.14159) + 3.14159)) / (0.0 - 0.1)) < 0.1) {
                            var v61 float64 = (((((0.0 - (float64(3.14159) * 3.14159)) - (0.0 * ((1.5 / 1.5) / (1.5 * 1.5)))) * 0.1) / 3.14159) / 1.0)
                            _ = v61
                        }
                        _ = v60
                    }
                    var v62 float32 = 3.14159
                    for fi63 := 0; fi63 < 12; fi63++ {
                        fmt.Println(v4)
                        if (false && (false || false)) {
                            if ("hello" == ("hello" + ((v16 + ("" + ("" + v16))) + v16))) {
                                if ((-21 >= 127) || false) {
                                    if false {
                                        var v64 int32 = 2147483647
                                        _ = v64
                                    } else {
                                        fmt.Println(v1)
                                        var v65 float32 = 3.14159
                                        _ = v65
                                    }
                                }
                                fmt.Println(v3, v4, v62)
                                if true {
                                    v1 += v1
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                } else {
                                    if len(v13) > 0 {
                                        v13[0] = (^(fi63))
                                    }
                                    var v66 int16 = v2
                                    switch fi63 {
                                    case -252:
                                        v13 = append(v13, (fi63 - (fi63 / (((fi37 | 32767)) | 1))))
                                    }
                                    _ = v66
                                }
                            }
                            v13 = append(v13, 4294967295)
                            var v67 uint16 = uint16(v6)
                            _ = v67
                        } else {
                            v2 = 0
                        }
                        _ = fi63
                    }
                    _ = v62
                }
                var v68 int32 = -2147483648
                _ = v68
            }
            _ = v38
        }
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
