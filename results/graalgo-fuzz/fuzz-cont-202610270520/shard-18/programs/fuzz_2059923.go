package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 int64; F2 int8 }
type GenS6 struct { F0 string; F1 uint8; F2 uint }
type GenS7 struct { F0 int64; F1 uint8; F2 uint }

func main() {
    var v1 uint = 289
    v2 := []uint32{804, 1, 1, 0}
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 2147483648
    if (73 <= -2147483648) {
        v5 := GenS1{F0: 4294967295, F1: 42}
        switch "a" {
        case "":
            v6 := GenS1{F0: (v4 + v4), F1: -9223372036854775808}
            _ = v6
        case "foo bar":
            for fi7 := 0; fi7 < 6; fi7++ {
                for fi8 := 0; fi8 < 1; fi8++ {
                    v9 := v3[42]
                    fmt.Println(v9, v4, fi8)
                    var v10 int32 = -1
                    switch fi7 {
                    case 4294967295:
                        fmt.Println(fi7)
                    case 0:
                        if true {
                            var v11 int16 = 1000
                            var v12 bool = ((true || false) && (true || true))
                            _ = v11
                            _ = v12
                        }
                    case 127:
                        v4 = v9
                    default:
                        v4 += 4294967295
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                switch fi7 {
                case -9223372036854775808:
                    fmt.Println(v1, v1, v4)
                default:
                    var v13 int = fi7
                    _ = v13
                }
                v14 := GenS2{F0: 1, F1: 65535}
                _ = fi7
                _ = v14
            }
        default:
            if false {
                for fi15 := 0; fi15 < 15; fi15++ {
                    v1 = 65535
                    _ = fi15
                }
                for fi16 := 0; fi16 < 6; fi16++ {
                    var v17 string = ("hello" + ("\n\t" + "\n\t"))
                    for fi18 := 0; fi18 < 12; fi18++ {
                        _ = len(v2)
                        _ = fi18
                    }
                    _ = fi16
                    _ = v17
                }
                for fi19 := 0; fi19 < 9; fi19++ {
                    if false {
                        v1 += (v1 & (v1 * v1))
                    }
                    _ = len(v2)
                    for fi20 := 0; fi20 < 3; fi20++ {
                        switch "hello" {
                        case "hello":
                            v4 = (v4 - (v4 - (v4 << (((v1 * (v1 * (v1 * v1)))) & 63))))
                        default:
                            fmt.Println(fi19, v1)
                        }
                        _ = fi20
                    }
                    _ = fi19
                }
            }
        }
        if ((-1 > 1) || true) {
            if (1 <= (v4 % ((2147483648) | 1))) {
                for fi21 := 0; fi21 < 8; fi21++ {
                    fmt.Println(v1)
                    _ = fi21
                }
                fmt.Println(v1, v4)
            } else {
                var v22 int32 = 167
                v23 := GenS3{F0: (-(v22)), F1: ("a" + ("foo bar" + ""))}
                if (-1 >= -524) {
                    v24 := GenS1{F0: (v4 & v4), F1: -1}
                    v22 += v22
                    if false {
                        var v25 int64 = int64(v22)
                        v22 = (v22 & v22)
                        var v26 bool = ("a" > ("foo bar" + "a"))
                        _ = v25
                        _ = v26
                    }
                    _ = v24
                }
                _ = v22
                _ = v23
            }
            fmt.Println(v4, v1)
        } else {
            fmt.Println(v1)
        }
        _ = v5
    } else {
        var v27 int32 = int32(v4)
        if (!false) {
            for fi28 := 0; fi28 < 6; fi28++ {
                v29 := GenS4{F0: -1}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v30 := GenS2{F0: 42, F1: 255}
                _ = fi28
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 15; fi31++ {
                for fi32 := 0; fi32 < 7; fi32++ {
                    v27 = v27
                    _ = fi32
                }
                v4 = v4
                _ = fi31
            }
        }
        _ = v27
    }
    if false {
        if false {
            if false {
                if len(v2) > 0 {
                    v2[0] = 1
                }
                delete(v3, 2)
            }
            v4 = v4
            v1 = v1
        }
        v4 -= 2147483648
        var v33 string = "foo bar"
        _ = v33
    } else {
        var v34 int8 = 100
        _ = v34
    }
    v4 -= v4
    if (false || false) {
        v1 = v1
    } else {
        v35 := GenS2{F0: 143, F1: 0}
        var v36 uint32 = 2
        _ = v35
        _ = v36
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        fmt.Println(v1, v1, v1)
        _ = fi37
    }
    delete(v3, 765)
    var v38 uint8 = 42
    fmt.Println(v4, v38)
    v1 = v1
    switch 9223372036854775807 {
    case -2147483648:
        v39 := GenS3{F0: 2147483647, F1: ((((("" + ("a" + ("" + "\n\t"))) + "hello") + "a") + "\n\t") + "foo bar")}
        _ = v39
    default:
        if true {
            v40 := GenS5{F0: 816, F1: -1, F2: -1}
            fmt.Println(v38, v1, v4)
            _ = v40
        } else {
            if (((0.0 > 3.14159) && (false && false)) || true) {
                var v41 int = -9223372036854775808
                for fi42 := 0; fi42 < 15; fi42++ {
                    v43 := GenS6{F0: ("\n\t" + ("hello" + "\n\t")), F1: 42, F2: (v1 % ((v1) | 1))}
                    v1 = 42
                    _ = fi42
                    _ = v43
                }
                _ = v41
            } else {
                v4 = 209
                var v44 float64 = 3.14159
                switch 32767 {
                case 2147483647:
                    if true {
                        if (!false) {
                            if (-630 < -540) {
                                if false {
                                    var v45 uint = (v1 & v1)
                                    if true {
                                        if len(v2) > 0 {
                                            v2[0] = v4
                                        }
                                        for fi46 := 0; fi46 < 1; fi46++ {
                                            v4 -= v4
                                            var v47 bool = true
                                            _ = fi46
                                            _ = v47
                                        }
                                        var v48 float32 = 3.14159
                                        _ = v48
                                    }
                                    fmt.Println(v45, v44, v4)
                                    _ = v45
                                }
                                if (42 > 0) {
                                    for fi49 := 0; fi49 < 10; fi49++ {
                                        for fi50 := 0; fi50 < 6; fi50++ {
                                            v44 = ((v44 + ((1e10 - (v44 * v44)) * (v44 - v44))) * ((-2.25 + float64(0.1)) / float64(1.5)))
                                            _ = fi50
                                        }
                                        v51 := v3[(v4 + uint32(v1))]
                                        switch (fi49 / (((fi49 / ((-617) | 1))) | 1)) {
                                        case 128:
                                            v44 += ((v44 + 1.5) - float64(1e10))
                                        case -1:
                                            var v52 float64 = 3.14159
                                            _ = v52
                                        }
                                        _ = fi49
                                        _ = v51
                                    }
                                    fmt.Println(v4)
                                    if true {
                                        v44 = (-2.25 * ((-2.25 * v44) * 3.14159))
                                        var v53 uint16 = 695
                                        if len(v2) > 0 {
                                            v2[0] = (v4 * (v4 << ((v1) & 63)))
                                        }
                                        _ = v53
                                    }
                                } else {
                                    var v54 uint = 1
                                    for fi55 := 0; fi55 < 5; fi55++ {
                                        if true {
                                            fmt.Println(v1)
                                            v4 += (v4 + (^(v4)))
                                            fmt.Println(v4)
                                        }
                                        if (218 < uint16(v4)) {
                                            fmt.Println(v38)
                                        } else {
                                            var v56 uint64 = 47
                                            var v57 uint32 = (v4 << ((v1) & 63))
                                            _ = v56
                                            _ = v57
                                        }
                                        _ = len(v3)
                                        _ = fi55
                                    }
                                    _ = v54
                                }
                            } else {
                                _ = len(v2)
                            }
                        }
                        for fi58 := 0; fi58 < 10; fi58++ {
                            if true {
                                for fi59 := 0; fi59 < 6; fi59++ {
                                    if true {
                                        if false {
                                            v60 := v3[(v4 % (((v4 << ((v1) & 63))) | 1))]
                                            _ = v60
                                        } else {
                                            v2 = append(v2, v4)
                                            var v61 uint16 = 2
                                            fmt.Println(v44, fi59)
                                            _ = v61
                                        }
                                        v62 := GenS1{F0: 2147483647, F1: -1}
                                        v63 := GenS2{F0: v38, F1: 1}
                                        _ = v62
                                        _ = v63
                                    }
                                    _ = fi59
                                }
                                for fi64 := 0; fi64 < 3; fi64++ {
                                    if true {
                                        switch 128 {
                                        case 1:
                                            if false {
                                                fmt.Println(v38, fi64, fi64)
                                            }
                                        case 32768:
                                            fmt.Println(v44, fi58, fi58)
                                        case 4294967295:
                                            var v65 uint8 = (v38 << ((999) & 63))
                                            switch (fi64 & -2147483648) {
                                            case -9223372036854775808:
                                                var v66 int16 = 1
                                                _ = v66
                                            case 4294967295:
                                                var v67 int = fi64
                                                _ = v67
                                            case -1:
                                                fmt.Println(v1, v38)
                                            default:
                                                var v68 uint64 = 18446744073709551615
                                                _ = v68
                                            }
                                            if false {
                                                fmt.Println(fi58, v38, v4)
                                                v65 -= (v65 * v65)
                                            } else {
                                                v1 += (v1 ^ v1)
                                                v44 = (v44 / (v44 * ((v44 / -2.25) * v44)))
                                                v44 -= ((v44 + (float64(((3.14159 / 1e10) - 1e10)) * float64(0.1))) / -2.25)
                                                var v69 int8 = 127
                                                var v70 float64 = 3.14159
                                                _ = v69
                                                _ = v70
                                            }
                                            v71 := GenS7{F0: 9223372036854775807, F1: 209, F2: (v1 + (v1 << (((v1 * (v1 | v1))) & 63)))}
                                            _ = v65
                                            _ = v71
                                        default:
                                            if len(v2) > 0 {
                                                v2[0] = v4
                                            }
                                        }
                                    }
                                    _ = fi64
                                }
                                v72 := GenS7{F0: -481, F1: (v38 >> (((v1 + (^(v1)))) & 63)), F2: (v1 << ((4294967295) & 63))}
                                _ = v72
                            } else {
                                v73 := GenS2{F0: (v38 ^ (v38 % ((v38) | 1))), F1: 32767}
                                if (595 != -9223372036854775808) {
                                    var v74 float64 = v44
                                    _ = v74
                                }
                                _ = v73
                            }
                            for fi75 := 0; fi75 < 15; fi75++ {
                                if (-226 > -1) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(v4)
                                }
                                fmt.Println(v1, v1, v38)
                                var v76 uint = 65535
                                var v77 int32 = 904
                                _ = fi75
                                _ = v76
                                _ = v77
                            }
                            _ = fi58
                        }
                        var v78 string = ("hello" + "hello")
                        _ = v78
                    }
                case 128:
                    var v79 int8 = 100
                    _ = v79
                }
                _ = v44
            }
            switch "\n\t" {
            case "a":
                var v80 int32 = -2147483648
                _ = v80
            }
        }
    }
    v4 = (^(v4))
    v81 := GenS6{F0: "a", F1: (v38 & (v38 & v38)), F2: v1}
    v38 -= v38
    v38 = 108
    var v82 bool = ((32767 > 0) || true)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v81
    _ = v82
}
