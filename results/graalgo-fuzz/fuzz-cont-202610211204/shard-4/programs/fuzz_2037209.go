package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int16; F1 int; F2 int16 }
type GenS4 struct { F0 int32; F1 uint64; F2 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint, p1 int64, p2 string) int8 {
    if (!false) {
        fmt.Println(p0)
        p1 += p1
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    return -41
}
func genFn1(p0 int, p1 int16) uint32 {
    if (4294967295 == 493) {
        p1 = p1
    }
    p0 = (p0 % ((32767) | 1))
    var v5 int = p0
    switch (-(v5)) {
    case 128:
        var v6 uint16 = 0
        _ = v6
    case 818:
        if true {
            v7 := []int{-2147483648, -955, v5}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 bool = false
                _ = fi8
                _ = v9
            }
            var v10 int32 = 0
            _ = v10
        }
    case 127:
        switch v5 {
        case -9223372036854775808:
            var v11 int = (v5 | -741)
            _ = v11
        case -1:
            var v12 uint16 = 255
            _ = v12
        case -2147483648:
            var v13 string = "foo bar"
            _ = v13
        }
    }
    var v14 bool = false
    _ = p0
    _ = p1
    _ = v14
    _ = v5
    return 42
}

func main() {
    var v1 int = 954
    v2 := make(map[uint32]bool)
    var v3 uint16 = 255
    v4 := []int{}
    if (!(true && true)) {
        fmt.Println(v3)
        var v15 uint64 = 0
        for fi16 := 0; fi16 < 11; fi16++ {
            switch ("foo bar" + ("a" + "a")) {
            case "a":
                if len(v4) > 0 {
                    v4[0] = (v1 - (fi16 ^ v1))
                }
            case "\n\t":
                v15 = v15
            case "":
                var v17 uint16 = 65535
                _ = v17
            default:
                fmt.Println(v15, v3)
            }
            _ = fi16
        }
        _ = v15
    }
    v2[42] = false
    if ((!false) || ((-32768 > 1000) && (false && true))) {
        var v18 uint16 = 255
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 int64 = 42
            v18 = 842
            if (1 == 608) {
                if ((271 <= 820) || true) {
                    v3 += v18
                }
                v21 := GenS1{F0: "\n\t", F1: (1.5 > 1e10)}
                _ = v21
            } else {
                v2[genFn1(v1, 324)] = true
                var v22 int8 = genFn0(42, (^(v20)), ((("a" + ("a" + "hello")) + "a") + "\n\t"))
                _ = v22
            }
            _ = fi19
            _ = v20
        }
        var v23 int8 = 1
        _ = v18
        _ = v23
    } else {
        var v24 int8 = 1
        fmt.Println(v3, v1)
        _ = v24
    }
    v1 = v1
    _ = len(v2)
    if (!(-1 == -1)) {
        if true {
            fmt.Println(v1, v1)
            switch "" {
            case "\n\t":
                var v25 uint8 = 0
                _ = v25
            case "hello":
                v26 := GenS1{F0: "\n\t", F1: (((503 == 1) || true) && true)}
                _ = v26
            case "foo bar":
                v3 = (v3 / (((v3 * v3)) | 1))
            default:
                fmt.Println(v3, v1)
            }
            delete(v2, 1)
        } else {
            var v27 int64 = 42
            _ = v27
        }
    } else {
        switch ("hello" + (("" + "\n\t") + "foo bar")) {
        case "a":
            switch -9223372036854775808 {
            case 42:
                v28 := GenS2{F0: -2147483648, F1: "hello"}
                _ = v28
            case -197:
                v29 := GenS2{F0: 790, F1: (("hello" + "") + ("" + ("\n\t" + ("" + "foo bar"))))}
                _ = v29
            case 51:
                v4 = append(v4, -1)
            }
        case "foo bar":
            if false {
                for fi30 := 0; fi30 < 19; fi30++ {
                    var v31 uint8 = 95
                    _ = fi30
                    _ = v31
                }
                var v32 int32 = -2147483648
                _ = v32
            } else {
                v33 := v2[2147483648]
                for fi34 := 0; fi34 < 20; fi34++ {
                    for fi35 := 0; fi35 < 8; fi35++ {
                        v3 -= (v3 - v3)
                        _ = fi35
                    }
                    var v36 uint64 = 155
                    _ = fi34
                    _ = v36
                }
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 uint64 = 708
                    var v39 uint = 2
                    v40 := GenS3{F0: 625, F1: (v1 - (v1 / ((fi37) | 1))), F2: 133}
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = v33
            }
        case "hello":
            v1 -= (v1 + v1)
        default:
            var v41 float32 = (((((-2.25 * (3.14159 / ((((1e10 - ((0.1 * 1.5) * 1e10)) + 1.5) / ((-2.25 * ((1.5 * -2.25) / -2.25)) - 3.14159)) + 1.5))) * 3.14159) - 0.1) - (0.0 + 1e10)) - ((-2.25 / (float32(-2.25) - 0.0)) * (1e10 - 0.1)))
            _ = v41
        }
        v42 := GenS4{F0: 594, F1: 18446744073709551615, F2: (("a" != "\n\t") && (!false))}
        _ = v42
    }
    var v43 uint = 255
    v3 -= (v3 << ((v43) & 63))
    for fi44 := 0; fi44 < 4; fi44++ {
        if false {
            v3 = (v3 * v3)
            if (v1 == fi44) {
                switch 32768 {
                case -9223372036854775808:
                    for fi45 := 0; fi45 < 9; fi45++ {
                        v1 = fi44
                        fmt.Println(v3, v3)
                        v46 := GenS3{F0: 0, F1: v1, F2: 1000}
                        _ = len(v2)
                        v43 = (v43 | (v43 >> ((v43) & 63)))
                        v4 = append(v4, fi45)
                        var v47 float32 = 3.14159
                        v48 := GenS2{F0: -1, F1: ("a" + (("" + "\n\t") + ("hello" + "hello")))}
                        _ = fi45
                        _ = v46
                        _ = v47
                        _ = v48
                    }
                }
            }
            var v49 int64 = 9223372036854775807
            _ = v49
        }
        var v50 int32 = -1
        _ = fi44
        _ = v50
    }
    var v51 int8 = 100
    v1 += (v1 + v1)
    if (int8(v1) != v51) {
        for fi52 := 0; fi52 < 9; fi52++ {
            switch ("" + "") {
            case "\n\t":
                if false {
                    var v53 bool = false
                    if len(v4) > 0 {
                        v4[0] = (fi52 ^ v1)
                    }
                    _ = v53
                }
            case "":
                var v54 int16 = 1000
                _ = v54
            case "a":
                switch v1 {
                case -566:
                    v1 -= (^(fi52))
                case -9223372036854775808:
                    v3 = 4
                case 32768:
                    for fi55 := 0; fi55 < 11; fi55++ {
                        fmt.Println(v43, fi55)
                        var v56 int16 = 42
                        _ = fi55
                        _ = v56
                    }
                default:
                    if ((v43 * (v43 - 42)) < v43) {
                        switch v1 {
                        case 670:
                            if (!false) {
                                delete(v2, uint32(v43))
                                fmt.Println(v51)
                                var v57 bool = (true || true)
                                _ = v57
                            }
                        case 128:
                            var v58 uint16 = v3
                            _ = v58
                        default:
                            v59 := GenS2{F0: 149, F1: "foo bar"}
                            _ = v59
                        }
                        v1 -= 1
                    } else {
                        fmt.Println(v1, fi52)
                        for fi60 := 0; fi60 < 10; fi60++ {
                            switch "a" {
                            case "hello":
                                if true {
                                    v4 = append(v4, (v1 * -671))
                                }
                            case "a":
                                var v61 uint8 = 255
                                _ = v61
                            case "\n\t":
                                _ = len(v4)
                            default:
                                if true {
                                    v51 = (-(v51))
                                } else {
                                    if false {
                                        v51 += v51
                                    }
                                    fmt.Println(v51)
                                }
                            }
                            _ = fi60
                        }
                    }
                }
            }
            for fi62 := 0; fi62 < 10; fi62++ {
                v63 := GenS1{F0: ("" + "a"), F1: false}
                _ = fi62
                _ = v63
            }
            for fi64 := 0; fi64 < 6; fi64++ {
                v65 := GenS5{F0: 9223372036854775807}
                v1 -= fi64
                _ = fi64
                _ = v65
            }
            _ = fi52
        }
    } else {
        if false {
            for fi66 := 0; fi66 < 6; fi66++ {
                v51 -= v51
                if true {
                    if ((v51 << ((v43) & 63)) < v51) {
                        v1 = -9223372036854775808
                    } else {
                        _ = len(v4)
                    }
                    var v67 uint64 = 0
                    _ = v67
                } else {
                    switch "" {
                    case "a":
                        v2[uint32(v43)] = ((v43 / ((255) | 1)) < (v43 % ((v43) | 1)))
                    case "\n\t":
                        if (!true) {
                            var v68 string = "hello"
                            for fi69 := 0; fi69 < 12; fi69++ {
                                if ((((0.0 - ((0.0 + (0.1 / 3.14159)) / 1e10)) - (0.0 / 3.14159)) + 0.1) != 0.1) {
                                    v70 := GenS4{F0: -2147483648, F1: 1, F2: false}
                                    fmt.Println(v68)
                                    var v71 uint32 = 2147483648
                                    _ = v70
                                    _ = v71
                                }
                                _ = len(v4)
                                _ = fi69
                            }
                            _ = v68
                        }
                    case "foo bar":
                        v72 := GenS6{F0: false}
                        _ = v72
                    }
                    v73 := v2[65535]
                    v74 := v2[355]
                    _ = v73
                    _ = v74
                }
                _ = fi66
            }
            if false {
                var v75 int16 = 100
                var v76 uint64 = 1
                _ = v75
                _ = v76
            } else {
                v77 := GenS4{F0: -767, F1: 0, F2: (!false)}
                _ = v77
            }
            v51 -= (v51 & v51)
        } else {
            var v78 int16 = 0
            _ = v78
        }
        fmt.Println(v1)
        for fi79 := 0; fi79 < 20; fi79++ {
            v51 = v51
            v80 := GenS3{F0: 0, F1: (fi79 + v1), F2: -1}
            _ = fi79
            _ = v80
        }
    }
    var v81 int32 = 660
    v43 += 1
    fmt.Println(v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v51:", v51)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v51
    _ = v81
}
