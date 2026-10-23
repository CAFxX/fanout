package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 uint64; F1 int; F2 uint16 }
type GenS6 struct { F0 int8; F1 int64; F2 int }
type GenS7 struct { F0 uint; F1 int; F2 int8 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int; F1 int8 }

func main() {
    var v1 bool = false
    var v2 float64 = (float64((0.0 - -2.25)) * 3.14159)
    v3 := []int{32767}
    var v4 uint16 = 32767
    var v5 uint16 = v4
    var v6 int64 = -9223372036854775808
    fmt.Println(v5, v2)
    var v7 uint32 = 4294967295
    v1 = (!(!((129 != 42) || (((v2 / (v2 * float64((0.0 - 1.5)))) * (v2 * 0.1)) != (((v2 * (1e10 * (v2 / v2))) * (0.1 - v2)) / (v2 - (v2 + (v2 * 0.0))))))))
    fmt.Println(v1)
    for fi8 := 0; fi8 < 12; fi8++ {
        v1 = false
        fmt.Println(v2, v2, v4)
        v2 -= v2
        _ = fi8
    }
    fmt.Println(v4, v2, v5)
    v3 = append(v3, 9223372036854775807)
    if false {
        var v9 int8 = 87
        var v10 int = 4294967295
        fmt.Println(v7, v6)
        _ = v10
        _ = v9
    } else {
        if v1 {
            var v11 float32 = ((0.0 + (1.5 / 3.14159)) - 0.0)
            var v12 uint32 = 430
            _ = v11
            _ = v12
        }
        v1 = v1
    }
    if true {
        v13 := GenS1{F0: 2}
        _ = len(v3)
        var v14 int64 = v6
        _ = v13
        _ = v14
    } else {
        if (v7 > v7) {
            if (v2 < -2.25) {
                v1 = (1e10 == 1e10)
            } else {
                _ = len(v3)
                var v15 int = 127
                _ = v15
            }
            switch "a" {
            case "hello":
                switch "\n\t" {
                case "":
                    for fi16 := 0; fi16 < 3; fi16++ {
                        if (!v1) {
                            var v17 float32 = (1.5 * (0.1 * 0.1))
                            v4 += 0
                            _ = v17
                        } else {
                            fmt.Println(fi16, v6)
                        }
                        if false {
                            for fi18 := 0; fi18 < 19; fi18++ {
                                v19 := GenS2{F0: 32767}
                                v20 := GenS3{F0: 18446744073709551615, F1: 676}
                                if (v6 < (v6 & v6)) {
                                    fmt.Println(fi16, v7, v4)
                                    switch "foo bar" {
                                    case "":
                                        var v21 int = 2147483647
                                        _ = v21
                                    case "hello":
                                        fmt.Println(fi18, v5, v6)
                                    case "\n\t":
                                        v22 := GenS4{F0: v4, F1: 0}
                                        _ = v22
                                    default:
                                        v6 += v6
                                    }
                                    v2 = 0.1
                                }
                                if v1 {
                                    var v23 uint64 = 13
                                    _ = v23
                                } else {
                                    v3 = append(v3, 32767)
                                }
                                if (152 > 0) {
                                    switch ("hello" + "a") {
                                    case "foo bar":
                                        fmt.Println(v5, v6, fi18)
                                    case "a":
                                        v7 = (v7 >> ((613) & 63))
                                        fmt.Println(v6, v7, v1)
                                    }
                                    fmt.Println(v5)
                                    var v24 uint8 = uint8(v7)
                                    v25 := GenS1{F0: uint64(v24)}
                                    v5 = v4
                                    _ = v24
                                    _ = v25
                                }
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            var v26 float64 = (1e10 - v2)
                            v4 -= (v4 % ((134) | 1))
                            _ = v26
                        }
                        if true {
                            v27 := GenS1{F0: 4294967295}
                            _ = v27
                        } else {
                            var v28 uint64 = 4294967295
                            v7 -= v7
                            v29 := GenS2{F0: 32767}
                            _ = v28
                            _ = v29
                        }
                        _ = fi16
                    }
                case "a":
                    switch "hello" {
                    case "foo bar":
                        if ((-27 > -128) && v1) {
                            v30 := GenS1{F0: 1}
                            for fi31 := 0; fi31 < 18; fi31++ {
                                v2 -= (1.5 / 0.1)
                                if v1 {
                                    var v32 int32 = -380
                                    var v33 float64 = (float64(0.1) * ((float64(float32((((0.0 / (((3.14159 * v2) - (0.0 / v2)) + -2.25)) + (v2 + float64((0.0 / 1.0)))) - v2))) - ((1.5 / v2) / (v2 - (float64((3.14159 * 1.5)) * v2)))) / 3.14159))
                                    var v34 uint16 = 65535
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                } else {
                                    var v35 float64 = v2
                                    var v36 bool = (true || (true && v1))
                                    fmt.Println(v35, v4, v36)
                                    _ = v35
                                    _ = v36
                                }
                                _ = fi31
                            }
                            _ = v30
                        } else {
                            fmt.Println(v1)
                        }
                    case "hello":
                        v1 = true
                    case "a":
                        var v37 int32 = 343
                        _ = v37
                    }
                }
            default:
                _ = len(v3)
            }
        } else {
            var v38 uint16 = (v4 >> ((42) & 63))
            _ = v38
        }
        var v39 int64 = -186
        _ = v39
    }
    if len(v3) > 0 {
        v3[0] = 943
    }
    v40 := GenS3{F0: 275, F1: -1}
    v4 = v4
    if v1 {
        v41 := GenS5{F0: 478, F1: 122, F2: 255}
        _ = v41
    } else {
        for fi42 := 0; fi42 < 19; fi42++ {
            if (v7 >= v7) {
                var v43 uint8 = 2
                _ = v43
            } else {
                if false {
                    var v44 uint8 = 244
                    if len(v3) > 0 {
                        v3[0] = 983
                    }
                    v1 = (18446744073709551615 <= 65535)
                    _ = v44
                } else {
                    switch fi42 {
                    case 2147483647:
                        switch "hello" {
                        case "foo bar":
                            if (true || (127 >= -1)) {
                                v6 = v6
                                fmt.Println(v6, v4, v2)
                            }
                        default:
                            v6 = 1
                        }
                    case 128:
                        _ = len(v3)
                    }
                    fmt.Println(v4, v1)
                    if v1 {
                        if v1 {
                            var v45 int8 = 42
                            for fi46 := 0; fi46 < 4; fi46++ {
                                var v47 uint16 = v5
                                fmt.Println(v45, v45, v1)
                                _ = fi46
                                _ = v47
                            }
                            for fi48 := 0; fi48 < 17; fi48++ {
                                v45 = (v45 % (((v45 - (v45 % (((v45 & v45)) | 1)))) | 1))
                                var v49 int32 = -613
                                fmt.Println(v1, v7)
                                _ = fi48
                                _ = v49
                            }
                            _ = v45
                        } else {
                            switch "hello" {
                            case "\n\t":
                                v50 := GenS3{F0: 351, F1: -1}
                                _ = v50
                            case "hello":
                                for fi51 := 0; fi51 < 4; fi51++ {
                                    if false {
                                        var v52 int32 = 1
                                        switch int(v6) {
                                        case 32767:
                                            v4 += 1
                                        case -297:
                                            v52 += v52
                                        default:
                                            if (-2.25 > 0.0) {
                                                if (v6 <= -9223372036854775808) {
                                                    v3 = append(v3, fi51)
                                                    fmt.Println(v5, fi51)
                                                    v1 = (2 <= 4294967295)
                                                } else {
                                                    var v53 int = (fi42 | (fi51 + (fi42 * 32768)))
                                                    if v1 {
                                                        v1 = (!true)
                                                        var v54 int8 = 94
                                                        _ = v54
                                                    } else {
                                                        var v55 uint16 = v5
                                                        if v1 {
                                                            var v56 int16 = -1
                                                            _ = v56
                                                        } else {
                                                            switch ("" + "") {
                                                            case "a":
                                                                fmt.Println(v53, v2, v2)
                                                            case "foo bar":
                                                                if false {
                                                                    if false {
                                                                        v1 = true
                                                                        var v57 uint = 668
                                                                        var v58 float64 = ((((v2 / v2) + (3.14159 / v2)) + float64(1e10)) / (v2 + (float64(float32(v2)) + (v2 / v2))))
                                                                        _ = v57
                                                                        _ = v58
                                                                    }
                                                                }
                                                            default:
                                                                v59 := GenS6{F0: 79, F1: (-(v6)), F2: (fi51 * (fi51 - fi51))}
                                                                _ = v59
                                                            }
                                                            v6 -= (v6 ^ (v6 ^ int64(v7)))
                                                        }
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v60 := GenS2{F0: 86}
                                                            _ = v60
                                                        case "\n\t":
                                                            v6 += int64(v52)
                                                        case "hello":
                                                            var v61 float32 = float32((1e10 - v2))
                                                            _ = v61
                                                        default:
                                                            v62 := GenS1{F0: 0}
                                                            _ = v62
                                                        }
                                                        _ = v55
                                                    }
                                                    _ = v53
                                                }
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = (fi51 * (fi42 >> ((166) & 63)))
                                                }
                                                if v1 {
                                                    v2 += (((v2 + v2) + float64(1.5)) + float64(1.5))
                                                    switch ("hello" + "hello") {
                                                    case "\n\t":
                                                        _ = len(v3)
                                                    default:
                                                        if (1e10 == (((0.1 - (1.5 + 3.14159)) - 1e10) + (-2.25 + 0.1))) {
                                                            fmt.Println(v5)
                                                        } else {
                                                            var v63 float64 = float64(3.14159)
                                                            if (v1 || (!v1)) {
                                                                var v64 uint8 = 2
                                                                fmt.Println(v2, v1)
                                                                v65 := GenS7{F0: 2, F1: (fi42 >> ((4294967295) & 63)), F2: -1}
                                                                _ = v64
                                                                _ = v65
                                                            }
                                                            v66 := GenS8{F0: (v52 & v52)}
                                                            _ = v63
                                                            _ = v66
                                                        }
                                                    }
                                                }
                                                var v67 uint64 = 1
                                                v68 := GenS9{F0: (v52 >> ((315) & 63))}
                                                _ = v67
                                                _ = v68
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v69 uint64 = 42
                                            _ = v69
                                        }
                                        _ = v52
                                    } else {
                                        if (127 != -66) {
                                            v70 := GenS4{F0: v5, F1: 1}
                                            _ = v70
                                        } else {
                                            v6 += -9223372036854775808
                                        }
                                        v3 = append(v3, fi42)
                                    }
                                    _ = fi51
                                }
                            default:
                                if (-1 <= -128) {
                                    v6 = -658
                                    if v1 {
                                        _ = len(v3)
                                    }
                                } else {
                                    var v71 float32 = (1.5 / float32(v2))
                                    switch -15 {
                                    case 128:
                                        if len(v3) > 0 {
                                            v3[0] = fi42
                                        }
                                    case 0:
                                        var v72 string = ("hello" + "a")
                                        _ = v72
                                    }
                                    _ = v71
                                }
                            }
                        }
                    }
                }
            }
            fmt.Println(fi42, v6)
            _ = fi42
        }
        fmt.Println(v1, v1, v2)
        if ((v1 && v1) && (v5 < v5)) {
            v73 := GenS10{F0: 9223372036854775807, F1: 42}
            _ = v73
        } else {
            v5 = (v5 << ((883) & 63))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
}
