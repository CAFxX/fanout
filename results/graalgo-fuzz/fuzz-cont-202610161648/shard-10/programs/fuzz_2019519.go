package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int8; F1 bool; F2 bool }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }
type GenS4 struct { F0 string; F1 int64; F2 uint16 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 uint8; F2 int16 }
type GenS8 struct { F0 bool; F1 int }
func genFn0(p0 uint) uint8 {
    var v7 float64 = (0.1 * 3.14159)
    p0 -= p0
    if ((v7 != 1e10) && false) {
        v8 := []int{128, 4294967295}
        _ = v8
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            var v10 uint = 605
            v7 += 1e10
            v11 := GenS1{F0: 0, F1: 73}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = v7
    return 2
}

func main() {
    v1 := []uint32{42, 2, 2147483648}
    var v2 int64 = -9223372036854775808
    var v3 string = ("a" + "")
    var v4 uint16 = 874
    var v5 uint = 0
    var v6 uint32 = 4294967295
    v3 = "hello"
    v12 := GenS1{F0: 4294967295, F1: 255}
    switch "a" {
    case "hello":
        var v13 uint8 = 130
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 float64 = 3.14159
        _ = fi14
        _ = v15
    }
    var v16 uint8 = 255
    if true {
        v4 -= 1
        var v17 int16 = -32768
        v5 = v5
        _ = v17
    } else {
        if false {
            switch v3 {
            case "hello":
                v18 := GenS2{F0: int8(v2), F1: false, F2: true}
                _ = v18
            default:
                for fi19 := 0; fi19 < 20; fi19++ {
                    v6 -= 0
                    if false {
                        var v20 uint32 = uint32(v4)
                        if (1000 == 22) {
                            v2 = v2
                        } else {
                            fmt.Println(v2)
                            if false {
                                if (32767 > 554) {
                                    switch ("a" + (v3 + v3)) {
                                    case "foo bar":
                                        v21 := GenS2{F0: -1, F1: (true || (((v3 > "a") || false) && (!(uint8(v5) <= v16)))), F2: true}
                                        _ = v21
                                    case "":
                                        v1 = append(v1, v20)
                                    case "hello":
                                        fmt.Println(v16)
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    v6 = (v6 ^ 65535)
                                } else {
                                    var v22 uint = (v5 * (^(v5)))
                                    switch (fi19 ^ 42) {
                                    case -1:
                                        var v23 float32 = 3.14159
                                        _ = v23
                                    case 127:
                                        var v24 string = (v3 + "foo bar")
                                        _ = v24
                                    case -9223372036854775808:
                                        v2 += (v2 | (v2 >> (((v5 * 196)) & 63)))
                                    }
                                    _ = v22
                                }
                                var v25 float64 = (-2.25 + -2.25)
                                _ = v25
                            } else {
                                if (!false) {
                                    var v26 bool = true
                                    v20 -= 2
                                    var v27 uint16 = 64
                                    _ = v26
                                    _ = v27
                                } else {
                                    if false {
                                        var v28 int = fi19
                                        _ = v28
                                    } else {
                                        v6 = 968
                                        var v29 uint32 = uint32(v5)
                                        v5 -= (v5 ^ v5)
                                        _ = v29
                                    }
                                }
                                for fi30 := 0; fi30 < 5; fi30++ {
                                    var v31 float32 = ((1.5 * -2.25) / ((1e10 / 1.0) * 1.5))
                                    var v32 int64 = (-(v2))
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                                if (false || (255 == v4)) {
                                    if true {
                                        fmt.Println(v4)
                                        v1 = append(v1, v20)
                                        fmt.Println(v16, v2, fi19)
                                    }
                                    if (!false) {
                                        v1 = append(v1, v6)
                                    } else {
                                        if (false && true) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        } else {
                                            var v33 uint32 = 65535
                                            _ = v33
                                        }
                                        fmt.Println(v16)
                                        fmt.Println(v5, v4, v6)
                                    }
                                } else {
                                    v20 -= uint32(v5)
                                    for fi34 := 0; fi34 < 6; fi34++ {
                                        v35 := GenS3{F0: 2147483647, F1: 32767, F2: v5}
                                        var v36 float64 = (0.0 * float64(1e10))
                                        _ = fi34
                                        _ = v35
                                        _ = v36
                                    }
                                    v37 := GenS1{F0: 368, F1: (v16 % (((v16 | (v16 % ((v16) | 1)))) | 1))}
                                    _ = v37
                                }
                            }
                            fmt.Println(v4, fi19)
                        }
                        fmt.Println(v2, fi19, v5)
                        _ = v20
                    } else {
                        v38 := GenS3{F0: 0, F1: v4, F2: v5}
                        if ((v4 | (v4 ^ v4)) >= uint16(v6)) {
                            fmt.Println(v16, v3)
                            v6 = (v6 + (v6 / ((v6) | 1)))
                        }
                        _ = v38
                    }
                    for fi39 := 0; fi39 < 19; fi39++ {
                        if len(v1) > 0 {
                            v1[0] = (v6 | (v6 >> (((v5 % (((v5 / ((65535) | 1))) | 1))) & 63)))
                        }
                        _ = fi39
                    }
                    _ = fi19
                }
            }
            v40 := GenS1{F0: 4294967295, F1: (v16 >> ((70) & 63))}
            _ = v40
        }
    }
    switch v3 {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 209
        }
    default:
        var v41 int32 = 462
        _ = v41
    }
    if ((false || false) && false) {
        if (v6 >= v6) {
            if (false && (true && true)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v42 string = (("foo bar" + ((((v3 + "foo bar") + v3) + v3) + v3)) + (v3 + v3))
                _ = v42
            } else {
                v4 -= (v4 % ((v4) | 1))
                for fi43 := 0; fi43 < 2; fi43++ {
                    var v44 int = fi43
                    var v45 int64 = (v2 % ((-1) | 1))
                    _ = len(v1)
                    _ = fi43
                    _ = v44
                    _ = v45
                }
            }
            var v46 uint = v5
            if false {
                for fi47 := 0; fi47 < 9; fi47++ {
                    if true {
                        v48 := GenS2{F0: 1, F1: false, F2: false}
                        var v49 bool = false
                        if v49 {
                            if v49 {
                                fmt.Println(v49, v16)
                            } else {
                                var v50 uint64 = 1
                                switch (^(fi47)) {
                                case 814:
                                    v4 = v4
                                case 42:
                                    var v51 int32 = 1
                                    _ = v51
                                case -1:
                                    switch 9223372036854775807 {
                                    case -9223372036854775808:
                                        for fi52 := 0; fi52 < 2; fi52++ {
                                            var v53 int8 = -1
                                            _ = len(v1)
                                            _ = fi52
                                            _ = v53
                                        }
                                    case -2147483648:
                                        switch (("hello" + ((v3 + v3) + "hello")) + v3) {
                                        case "\n\t":
                                            switch (fi47 ^ (fi47 | 539)) {
                                            case 2147483647:
                                                if v49 {
                                                    v2 = (v2 >> (((v46 - (v5 * v46))) & 63))
                                                    var v54 string = "hello"
                                                    _ = v54
                                                }
                                            default:
                                                for fi55 := 0; fi55 < 2; fi55++ {
                                                    v5 += (^(v46))
                                                    v56 := GenS2{F0: 122, F1: v49, F2: ((v16 + (v16 + v16)) > (v16 + v16))}
                                                    _ = fi55
                                                    _ = v56
                                                }
                                            }
                                        case "":
                                            var v57 int32 = 753
                                            _ = v57
                                        case "hello":
                                            if v49 {
                                                v58 := GenS4{F0: v3, F1: int64(fi47), F2: (^(v4))}
                                                var v59 uint32 = (^(v6))
                                                if true {
                                                    switch 593 {
                                                    case 128:
                                                        var v60 int16 = 548
                                                        _ = v60
                                                    case -197:
                                                        v61 := GenS5{F0: 42, F1: -82}
                                                        _ = v61
                                                    case -912:
                                                        var v62 string = "foo bar"
                                                        _ = v62
                                                    }
                                                    var v63 uint = (v5 % ((0) | 1))
                                                    _ = v63
                                                } else {
                                                    v64 := GenS6{F0: -563}
                                                    v59 -= v6
                                                    _ = v64
                                                }
                                                _ = v58
                                                _ = v59
                                            }
                                        default:
                                            var v65 uint16 = (v4 >> ((v46) & 63))
                                            _ = v65
                                        }
                                    }
                                }
                                _ = v50
                            }
                        }
                        _ = v48
                        _ = v49
                    }
                    var v66 uint16 = v4
                    _ = fi47
                    _ = v66
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                for fi67 := 0; fi67 < 4; fi67++ {
                    if (!(!(true && true))) {
                        fmt.Println(v16, v4, fi67)
                        v2 = -9223372036854775808
                    } else {
                        var v68 uint8 = (v16 << ((v46) & 63))
                        var v69 int = (fi67 >> ((v46) & 63))
                        var v70 float64 = (1.5 * 0.0)
                        _ = v68
                        _ = v69
                        _ = v70
                    }
                    _ = fi67
                }
            } else {
                v1 = append(v1, (v6 & (v6 >> ((v5) & 63))))
                v46 = v46
                for fi71 := 0; fi71 < 14; fi71++ {
                    v5 -= (v5 * 18446744073709551615)
                    for fi72 := 0; fi72 < 15; fi72++ {
                        if len(v1) > 0 {
                            v1[0] = 589
                        }
                        _ = fi72
                    }
                    _ = len(v1)
                    _ = fi71
                }
            }
            _ = v46
        } else {
            switch (v3 + "hello") {
            case "\n\t":
                switch v3 {
                case "a":
                    if false {
                        v73 := GenS1{F0: 13, F1: (^(v16))}
                        switch ((v3 + v3) + v3) {
                        case "\n\t":
                            var v74 uint64 = 56
                            _ = v74
                        case "a":
                            v2 = -575
                        case "":
                            v16 += v16
                        }
                        _ = v73
                    }
                case "hello":
                    if (true && true) {
                        for fi75 := 0; fi75 < 4; fi75++ {
                            v3 = v3
                            _ = fi75
                        }
                        v5 += v5
                    } else {
                        if false {
                            v1 = append(v1, v6)
                            if ((v16 - v16) >= genFn0(13)) {
                                var v76 float64 = (0.0 + 1e10)
                                for fi77 := 0; fi77 < 8; fi77++ {
                                    if (0 == (^(v6))) {
                                        var v78 int8 = 127
                                        _ = v78
                                    }
                                    _ = len(v1)
                                    _ = fi77
                                }
                                fmt.Println(v76)
                                _ = v76
                            }
                        } else {
                            switch 128 {
                            case 32767:
                                switch -1 {
                                case 128:
                                    switch "\n\t" {
                                    case "hello":
                                        if (v2 != v2) {
                                            var v79 int16 = 42
                                            _ = v79
                                        }
                                    case "\n\t":
                                        _ = len(v1)
                                    case "a":
                                        for fi80 := 0; fi80 < 4; fi80++ {
                                            var v81 float64 = 1e10
                                            _ = fi80
                                            _ = v81
                                        }
                                    default:
                                        if false {
                                            for fi82 := 0; fi82 < 5; fi82++ {
                                                switch "" {
                                                case "a":
                                                    if len(v1) > 0 {
                                                        v1[0] = 4294967295
                                                    }
                                                case "\n\t":
                                                    v16 += (v16 << ((v5) & 63))
                                                }
                                                var v83 uint8 = (v16 >> ((v5) & 63))
                                                v83 -= (^(v83))
                                                _ = fi82
                                                _ = v83
                                            }
                                        } else {
                                            v84 := GenS1{F0: uint64(v6), F1: 42}
                                            _ = v84
                                        }
                                    }
                                default:
                                    v85 := GenS3{F0: 662, F1: v4, F2: (v5 / (((v5 % ((v5) | 1))) | 1))}
                                    _ = v85
                                }
                            case 2147483647:
                                for fi86 := 0; fi86 < 2; fi86++ {
                                    var v87 int8 = 1
                                    var v88 uint = v5
                                    _ = fi86
                                    _ = v87
                                    _ = v88
                                }
                            default:
                                v3 = (v3 + v3)
                            }
                            v5 = (^(v5))
                            v89 := GenS7{F0: -1, F1: uint8(v5), F2: int16(v16)}
                            _ = v89
                        }
                        var v90 float64 = 1.5
                        _ = v90
                    }
                case "\n\t":
                    for fi91 := 0; fi91 < 20; fi91++ {
                        for fi92 := 0; fi92 < 7; fi92++ {
                            fmt.Println(v16, v2, v16)
                            _ = fi92
                        }
                        var v93 float64 = 0.0
                        _ = fi91
                        _ = v93
                    }
                default:
                    fmt.Println(v3, v2)
                }
            default:
                var v94 uint16 = v4
                _ = v94
            }
        }
        var v95 uint32 = (^(v6))
        v96 := GenS8{F0: true, F1: 448}
        _ = v95
        _ = v96
    }
    for fi97 := 0; fi97 < 17; fi97++ {
        v98 := GenS1{F0: uint64(v16), F1: genFn0((v5 + v5))}
        v16 -= (v16 + genFn0((v5 ^ v5)))
        for fi99 := 0; fi99 < 1; fi99++ {
            fmt.Println(v6, v2)
            v4 -= (^(v4))
            _ = fi99
        }
        _ = fi97
        _ = v98
    }
    var v100 string = v3
    v5 = (^(v5))
    if len(v1) > 0 {
        v1[0] = v6
    }
    if (4294967295 >= v6) {
        switch 9223372036854775807 {
        case 2147483647:
            for fi101 := 0; fi101 < 14; fi101++ {
                var v102 int8 = -7
                _ = fi101
                _ = v102
            }
        }
    } else {
        if (421 >= 18446744073709551615) {
            var v103 int8 = -24
            _ = v103
        }
        var v104 int = -730
        _ = v104
    }
    for fi105 := 0; fi105 < 11; fi105++ {
        var v106 bool = false
        v107 := GenS3{F0: -1, F1: v4, F2: v5}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi105
        _ = v106
        _ = v107
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v100:", v100)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v100
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
