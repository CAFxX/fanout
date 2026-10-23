package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int64; F1 int; F2 int }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 uint }
type GenS8 struct { F0 string; F1 uint; F2 uint8 }
type GenS9 struct { F0 uint; F1 int8; F2 uint16 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 uint64 }
type GenS12 struct { F0 bool; F1 uint64 }

func main() {
    var v1 uint = 934
    var v2 uint64 = 1
    var v3 int16 = 1
    var v4 int32 = -2147483648
    v4 -= v4
    v5 := GenS1{F0: 2147483648}
    v6 := GenS2{F0: (v2 % ((v2) | 1))}
    v4 -= v4
    v7 := GenS2{F0: (v2 ^ (v2 ^ v2))}
    if true {
        for fi8 := 0; fi8 < 11; fi8++ {
            switch ("hello" + "foo bar") {
            case "foo bar":
                v9 := GenS1{F0: 53}
                _ = v9
            case "a":
                v1 = (v1 * (v1 << (((v1 + v1)) & 63)))
            case "\n\t":
                for fi10 := 0; fi10 < 6; fi10++ {
                    v1 = v1
                    v11 := []int{-9223372036854775808, fi10, fi10, (fi8 & 900)}
                    _ = fi10
                    _ = v11
                }
            }
            if (42 == 9223372036854775807) {
                if ((0.0 / (1e10 / 3.14159)) < 1e10) {
                    for fi12 := 0; fi12 < 11; fi12++ {
                        v2 -= (^(v2))
                        if ((v2 & v2) >= 42) {
                            if (v1 != v1) {
                                if true {
                                    fmt.Println(fi12, v2, v3)
                                }
                                v4 += (v4 * 0)
                            }
                            v13 := GenS3{F0: -40}
                            _ = v13
                        }
                        var v14 string = "a"
                        _ = fi12
                        _ = v14
                    }
                }
                v2 += v2
            } else {
                for fi15 := 0; fi15 < 10; fi15++ {
                    v3 = -103
                    if true {
                        fmt.Println(v2, fi8, fi8)
                    }
                    _ = fi15
                }
                switch "\n\t" {
                case "":
                    for fi16 := 0; fi16 < 12; fi16++ {
                        if false {
                            fmt.Println(fi8, fi8, fi16)
                        } else {
                            var v17 uint = v1
                            switch "" {
                            case "\n\t":
                                v18 := GenS4{F0: 9223372036854775807, F1: (fi8 ^ fi16), F2: fi16}
                                v19 := GenS4{F0: 42, F1: (^(fi8)), F2: 1}
                                v20 := []int{(-(fi16)), fi16, int(v3), fi16}
                                _ = v18
                                _ = v19
                                _ = v20
                            case "":
                                v2 -= v2
                            case "a":
                                v4 = v4
                            default:
                                v1 += 941
                            }
                            _ = v17
                        }
                        fmt.Println(fi8, v3)
                        if false {
                            if (true && true) {
                                v21 := GenS5{F0: 42, F1: 255}
                                if false {
                                    v4 = 669
                                }
                                var v22 int8 = -61
                                _ = v21
                                _ = v22
                            }
                            var v23 uint8 = 244
                            _ = v23
                        }
                        if false {
                            v2 = (v2 ^ (v2 + 4294967295))
                            fmt.Println(v4)
                            if (((float32((-2.25 * 1e10)) + 0.1) + float32(((0.0 + (3.14159 + (-2.25 * float64(0.0)))) * 0.0))) <= -2.25) {
                                v3 = v3
                                v24 := GenS1{F0: 65535}
                                if (!(!((true || false) || ((0 > v2) || (false && ((v2 % ((v2) | 1)) > (v2 % ((v2) | 1)))))))) {
                                    v3 -= v3
                                    var v25 bool = (3.14159 > 1e10)
                                    _ = v25
                                } else {
                                    if (true && false) {
                                        if true {
                                            var v26 uint16 = 42
                                            v26 += v26
                                            _ = v26
                                        }
                                        v27 := GenS1{F0: 65535}
                                        v2 = 506
                                        switch "a" {
                                        case "foo bar":
                                            switch ("foo bar" + "foo bar") {
                                            case "foo bar":
                                                if (!(434 <= int64(v3))) {
                                                    var v28 int64 = 330
                                                    v28 = 9223372036854775807
                                                    _ = v28
                                                } else {
                                                    v29 := GenS1{F0: 2147483647}
                                                    v30 := []int{(^(fi16))}
                                                    _ = v29
                                                    _ = v30
                                                }
                                                v3 -= v3
                                                v4 -= v4
                                            default:
                                                v31 := GenS1{F0: 42}
                                                _ = v31
                                            }
                                        case "":
                                            switch "hello" {
                                            case "":
                                                var v32 uint16 = 802
                                                _ = v32
                                            case "a":
                                                switch "foo bar" {
                                                case "hello":
                                                    var v33 uint16 = 32767
                                                    var v34 int32 = (-(v4))
                                                    v35 := []int{}
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                case "\n\t":
                                                    fmt.Println(v4, v1, fi8)
                                                default:
                                                    v3 -= v3
                                                }
                                            case "foo bar":
                                                var v36 uint16 = 181
                                                _ = v36
                                            }
                                        }
                                        _ = v27
                                    } else {
                                        var v37 int32 = 908
                                        if true {
                                            v38 := GenS6{F0: 0}
                                            _ = v38
                                        }
                                        _ = v37
                                    }
                                }
                                _ = v24
                            }
                        }
                        fmt.Println(v3, fi8, v4)
                        var v39 int64 = -915
                        _ = fi16
                        _ = v39
                    }
                case "\n\t":
                    if ((v1 < 780) && ((v1 * 65535) != v1)) {
                        if (((v2 <= 42) && (!false)) || true) {
                            v40 := []int{-2147483648, fi8}
                            switch "" {
                            case "foo bar":
                                v41 := GenS7{F0: (v2 >> (((v1 >> ((v1) & 63))) & 63)), F1: (v1 + 255)}
                                _ = v41
                            case "a":
                                v4 = v4
                            default:
                                v42 := GenS8{F0: "", F1: v1, F2: 127}
                                _ = v42
                            }
                            _ = v40
                        } else {
                            v43 := []int{fi8}
                            if len(v43) >= 2 {
                                v43 = v43[1:len(v43)]
                            }
                            _ = v43
                        }
                        if (1e10 < (1.5 / 1.0)) {
                            if true {
                                if ((0.1 * 1e10) < (-2.25 * -2.25)) {
                                    v44 := []int{}
                                    if (544 > 32767) {
                                        if false {
                                            var v45 string = "\n\t"
                                            var v46 int8 = 1
                                            _ = v45
                                            _ = v46
                                        }
                                    } else {
                                        var v47 uint = (v1 / (((^(v1))) | 1))
                                        switch fi8 {
                                        case -1:
                                            v48 := GenS9{F0: (v47 ^ (v1 | (v1 | v47))), F1: -1, F2: 1}
                                            _ = v48
                                        case -555:
                                            for fi49 := 0; fi49 < 4; fi49++ {
                                                switch (-(fi8)) {
                                                case 32767:
                                                    if len(v44) > 0 {
                                                        v44[0] = (fi49 >> (((v1 * (v1 | v47))) & 63))
                                                    }
                                                case 32768:
                                                    var v50 int32 = (v4 * 42)
                                                    _ = v50
                                                case 9223372036854775807:
                                                    if false {
                                                        v51 := GenS10{F0: (fi49 % (((fi8 % ((-9223372036854775808) | 1))) | 1))}
                                                        v52 := GenS11{F0: (2 >= 2), F1: v2}
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                default:
                                                    if len(v44) > 0 {
                                                        v44[0] = 647
                                                    }
                                                }
                                                if (!true) {
                                                    v53 := GenS12{F0: true, F1: (v2 * (v2 + (v2 % (((v2 / ((v2) | 1))) | 1))))}
                                                    _ = v53
                                                } else {
                                                    v47 += (^(v1))
                                                    switch fi49 {
                                                    case 42:
                                                        if (-1 <= 1) {
                                                            v1 -= v1
                                                            fmt.Println(v2, v2, v47)
                                                            fmt.Println(fi8)
                                                        } else {
                                                            v2 += (v2 / ((v2) | 1))
                                                        }
                                                    case 128:
                                                        v3 = int16(v4)
                                                    default:
                                                        v47 = v1
                                                    }
                                                    fmt.Println(v47, v47, v2)
                                                }
                                                _ = fi49
                                            }
                                        default:
                                            v54 := GenS4{F0: 9223372036854775807, F1: -543, F2: (fi8 / (((fi8 + (fi8 * 0))) | 1))}
                                            _ = v54
                                        }
                                        _ = v47
                                    }
                                    _ = v44
                                }
                                var v55 uint32 = 2
                                _ = v55
                            }
                        }
                    }
                case "a":
                    var v56 uint64 = (^(v2))
                    _ = v56
                }
                if (("hello" + (("foo bar" + (("" + ("a" + "hello")) + "a")) + "foo bar")) == "foo bar") {
                    v3 = v3
                    v57 := []int{(fi8 | (fi8 - fi8)), 9223372036854775807, (fi8 + -9223372036854775808)}
                    v58 := GenS4{F0: 42, F1: fi8, F2: 0}
                    _ = v57
                    _ = v58
                } else {
                    v59 := []int{fi8}
                    v60 := GenS9{F0: v1, F1: 44, F2: 0}
                    _ = v59
                    _ = v60
                }
            }
            _ = fi8
        }
    } else {
        v61 := []int{32767, -9223372036854775808}
        if true {
            fmt.Println(v3, v4, v2)
        }
        _ = v61
    }
    if (((v2 - 1) <= 0) || (952 > -1)) {
        for fi62 := 0; fi62 < 14; fi62++ {
            if true {
                v2 = (^(v2))
            } else {
                v63 := []int{(fi62 << ((v1) & 63)), (fi62 >> (((v1 << ((v1) & 63))) & 63)), int(v1), fi62}
                _ = v63
            }
            _ = fi62
        }
    } else {
        var v64 uint16 = 32767
        _ = v64
    }
    v65 := []int{-1}
    if false {
        var v66 float32 = 0.1
        _ = v66
    }
    var v67 uint32 = 878
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v67
    _ = v7
}
