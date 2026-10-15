package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int16; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16 }
func genFn0() int {
    var v4 float64 = 3.14159
    switch "hello" {
    case "foo bar":
        v5 := GenS1{F0: 1, F1: -9223372036854775808}
        _ = v5
    case "\n\t":
        var v6 bool = true
        _ = v6
    case "":
        v4 = v4
    default:
        v4 += 1.5
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(fi7, v4)
        switch "foo bar" {
        case "hello":
            v4 += v4
        default:
            for fi8 := 0; fi8 < 17; fi8++ {
                switch "hello" {
                case "a":
                    var v9 int16 = 948
                    _ = v9
                }
                if (90 < 94) {
                    v10 := []int{fi7}
                    _ = v10
                } else {
                    v11 := []int{(-(fi8)), fi7}
                    _ = v11
                }
                var v12 uint = 18446744073709551615
                _ = fi8
                _ = v12
            }
        }
        _ = fi7
    }
    fmt.Println(v4)
    var v13 uint8 = 1
    _ = v13
    _ = v4
    return 1
}
func genFn1() uint16 {
    var v14 string = ("" + "\n\t")
    v14 = ((v14 + "foo bar") + ("foo bar" + ("foo bar" + v14)))
    fmt.Println(v14)
    _ = v14
    return 324
}

func main() {
    v1 := []uint32{}
    var v2 float32 = 1e10
    var v3 uint = 2
    var v15 uint = v3
    var v16 int32 = 283
    v17 := GenS1{F0: -1, F1: 0}
    v18 := GenS1{F0: (v16 - (v16 + (-(v16)))), F1: -9223372036854775808}
    if false {
        var v19 uint16 = 2
        if (!(!true)) {
            for fi20 := 0; fi20 < 8; fi20++ {
                v3 = (^(v15))
                if true {
                    if len(v1) > 0 {
                        v1[0] = 157
                    }
                }
                _ = fi20
            }
            v19 -= v19
        } else {
            if true {
                for fi21 := 0; fi21 < 15; fi21++ {
                    v19 += (v19 << (((v15 + v3)) & 63))
                    for fi22 := 0; fi22 < 13; fi22++ {
                        if true {
                            fmt.Println(v2)
                            _ = len(v1)
                        }
                        var v23 int16 = 0
                        _ = fi22
                        _ = v23
                    }
                    var v24 int32 = (v16 - 0)
                    _ = fi21
                    _ = v24
                }
                fmt.Println(v19)
                if true {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        switch "hello" {
                        case "a":
                            var v26 uint64 = 2
                            _ = v26
                        default:
                            var v27 int8 = 104
                            _ = v27
                        }
                        v19 = v19
                        _ = fi25
                    }
                    v19 = genFn1()
                }
            }
        }
        _ = v19
    }
    if (v2 >= (0.1 / (-2.25 / (v2 * (((3.14159 - float32((0.0 + 0.1))) + v2) - 3.14159))))) {
        fmt.Println(v16, v15)
        switch -968 {
        case -2147483648:
            if (true || false) {
                switch 128 {
                case 142:
                    for fi28 := 0; fi28 < 1; fi28++ {
                        v2 += v2
                        _ = fi28
                    }
                }
                var v29 uint8 = 2
                if (((float64((v2 + float32(1.5))) + 3.14159) * 3.14159) > (float64(0.1) + ((1e10 / -2.25) / (0.1 + 0.1)))) {
                    if true {
                        if false {
                            if true {
                                v3 = 255
                                var v30 uint64 = 18446744073709551615
                                _ = v30
                            }
                            fmt.Println(v2, v29, v3)
                        }
                        v29 -= (^(v29))
                        for fi31 := 0; fi31 < 2; fi31++ {
                            v2 += v2
                            _ = fi31
                        }
                    }
                }
                _ = v29
            } else {
                var v32 uint64 = 327
                _ = v32
            }
        default:
            switch int(v15) {
            case -1:
                v2 = (((v2 / 3.14159) + v2) * (3.14159 + (1e10 + float32(-2.25))))
            case -2147483648:
                var v33 uint16 = 65535
                _ = v33
            case 756:
                for fi34 := 0; fi34 < 8; fi34++ {
                    v35 := GenS2{F0: -925, F1: -122, F2: 42}
                    for fi36 := 0; fi36 < 13; fi36++ {
                        fmt.Println(fi36, v2, v3)
                        var v37 uint16 = 65535
                        _ = fi36
                        _ = v37
                    }
                    _ = fi34
                    _ = v35
                }
            default:
                if false {
                    if (42 < -32768) {
                        for fi38 := 0; fi38 < 18; fi38++ {
                            fmt.Println(v16, v3)
                            _ = fi38
                        }
                    } else {
                        v15 -= v15
                        for fi39 := 0; fi39 < 5; fi39++ {
                            v16 = v16
                            _ = fi39
                        }
                    }
                    fmt.Println(v2, v16)
                    v40 := GenS2{F0: v16, F1: 1, F2: 2}
                    _ = v40
                }
            }
        }
        for fi41 := 0; fi41 < 7; fi41++ {
            for fi42 := 0; fi42 < 3; fi42++ {
                v3 = v15
                _ = len(v1)
                if true {
                    v43 := GenS3{F0: 2}
                    if true {
                        if true {
                            v44 := GenS2{F0: (v16 & (v16 % ((1) | 1))), F1: 42, F2: 4294967295}
                            if len(v1) > 0 {
                                v1[0] = 221
                            }
                            v2 = 1e10
                            switch ("hello" + "a") {
                            case "hello":
                                v45 := GenS4{F0: v3, F1: -636, F2: 65535}
                                _ = v45
                            }
                            _ = len(v1)
                            _ = v44
                        }
                        var v46 string = ("hello" + ("hello" + "a"))
                        _ = v46
                    } else {
                        var v47 int32 = 2147483647
                        switch 42 {
                        case 128:
                            var v48 float64 = (float64(v2) - (float64(v2) / (3.14159 + float64(((v2 + v2) + 1e10)))))
                            _ = v48
                        case 938:
                            v47 += -1
                        default:
                            var v49 uint = (v15 << ((v15) & 63))
                            _ = v49
                        }
                        _ = v47
                    }
                    fmt.Println(v3)
                    _ = v43
                } else {
                    if (false && false) {
                        var v50 int64 = 0
                        v51 := GenS2{F0: 1, F1: v50, F2: 740}
                        v52 := GenS4{F0: (v15 | (^(v15))), F1: 792, F2: 2}
                        _ = v50
                        _ = v51
                        _ = v52
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v2 = (v2 / v2)
                        v53 := GenS4{F0: v3, F1: 997, F2: 2}
                        _ = v53
                    }
                    switch "hello" {
                    case "a":
                        v15 += (v15 - 1)
                    case "foo bar":
                        if true {
                            v16 = v16
                        }
                    case "hello":
                        v16 += v16
                    }
                    v54 := GenS1{F0: v16, F1: -9223372036854775808}
                    _ = v54
                }
                _ = fi42
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi41
        }
    } else {
        v2 -= 1e10
        if false {
            if len(v1) > 0 {
                v1[0] = 633
            }
            v1 = append(v1, 4294967295)
        } else {
            switch "hello" {
            case "foo bar":
                v15 -= (v15 / ((v15) | 1))
            case "\n\t":
                var v55 int32 = 535
                _ = v55
            case "a":
                v1 = append(v1, 2)
            default:
                for fi56 := 0; fi56 < 7; fi56++ {
                    v3 = (^(v3))
                    var v57 int8 = 72
                    v16 = (v16 - (v16 & v16))
                    _ = fi56
                    _ = v57
                }
            }
            for fi58 := 0; fi58 < 1; fi58++ {
                switch "hello" {
                case "a":
                    if false {
                        switch ("hello" + "") {
                        case "hello":
                            v15 = (^(v15))
                        case "\n\t":
                            v59 := GenS5{F0: 716}
                            _ = v59
                        case "a":
                            v60 := GenS4{F0: v3, F1: -1, F2: 235}
                            _ = v60
                        default:
                            var v61 int64 = 9223372036854775807
                            _ = v61
                        }
                    }
                default:
                    v3 += 275
                }
                switch fi58 {
                case 2147483647:
                    v62 := GenS1{F0: -2147483648, F1: 42}
                    _ = v62
                case -1:
                    var v63 uint16 = 232
                    _ = v63
                }
                _ = fi58
            }
        }
    }
    v15 -= 4294967295
    if false {
        if (770 >= genFn1()) {
            switch "\n\t" {
            case "hello":
                switch 2147483647 {
                case 9223372036854775807:
                    var v64 int32 = (-(v16))
                    _ = v64
                case -587:
                    for fi65 := 0; fi65 < 8; fi65++ {
                        v16 += 1
                        if len(v1) > 0 {
                            v1[0] = 2147483647
                        }
                        _ = fi65
                    }
                }
            }
        } else {
            fmt.Println(v2, v15)
            var v66 bool = true
            for fi67 := 0; fi67 < 19; fi67++ {
                if false {
                    v16 += (v16 / (((v16 ^ (^(v16)))) | 1))
                }
                var v68 string = "foo bar"
                var v69 uint16 = 42
                _ = fi67
                _ = v68
                _ = v69
            }
            _ = v66
        }
        if (-1 >= 0) {
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
            if ((100 == 42) || (616 > 42)) {
                switch 32767 {
                case -595:
                    var v70 bool = false
                    _ = v70
                case 2147483647:
                    if (793 >= genFn0()) {
                        v2 += 1.5
                    } else {
                        v1 = append(v1, 557)
                        v3 -= (v15 / ((0) | 1))
                        v3 = (v15 << (((v3 & v15)) & 63))
                    }
                }
            }
            v71 := GenS2{F0: (v16 * v16), F1: -1, F2: 42}
            _ = v71
        }
    } else {
        for fi72 := 0; fi72 < 18; fi72++ {
            for fi73 := 0; fi73 < 4; fi73++ {
                switch fi73 {
                case 128:
                    var v74 uint64 = 18446744073709551615
                    _ = v74
                case -168:
                    switch (fi73 / ((-902) | 1)) {
                    case -2147483648:
                        var v75 uint8 = 42
                        _ = v75
                    case 2147483647:
                        v76 := GenS6{F0: 0}
                        _ = v76
                    case 686:
                        _ = len(v1)
                    default:
                        fmt.Println(fi73, fi72)
                    }
                }
                v16 -= v16
                v16 -= (v16 * v16)
                _ = fi73
            }
            fmt.Println(fi72, v15, v2)
            if ((-991 <= v16) && (true && false)) {
                var v77 int16 = int16(fi72)
                _ = v77
            }
            _ = fi72
        }
    }
    if true {
        v78 := GenS7{F0: -32768}
        if true {
            fmt.Println(v16)
            for fi79 := 0; fi79 < 14; fi79++ {
                v2 = v2
                _ = fi79
            }
        }
        for fi80 := 0; fi80 < 5; fi80++ {
            v1 = append(v1, 858)
            fmt.Println(v15, fi80)
            _ = fi80
        }
        _ = v78
    } else {
        var v81 float64 = (0.1 - 0.1)
        _ = v81
    }
    v3 = (v3 << ((v3) & 63))
    var v82 int32 = int32(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v82
}
