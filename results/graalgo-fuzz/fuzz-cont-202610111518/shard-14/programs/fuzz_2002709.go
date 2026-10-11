package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 uint }
type GenS3 struct { F0 int; F1 string; F2 bool }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 bool; F1 int; F2 uint }
type GenS8 struct { F0 uint8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]bool)
    var v3 uint8 = 42
    var v4 float32 = 1e10
    var v5 bool = (!(!(!true)))
    var v6 float64 = float64((0.0 / float32(1e10)))
    var v7 uint64 = 2
    for fi8 := 0; fi8 < 13; fi8++ {
        v6 = float64(v4)
        switch (fi8 * 42) {
        case -492:
            v9 := GenS1{F0: 339, F1: 42}
            _ = v9
        case 127:
            var v10 uint8 = 127
            _ = v10
        }
        switch (("" + "hello") + "\n\t") {
        case "\n\t":
            v6 -= (((v6 - (v6 - float64((v4 + v4)))) * (v6 / v6)) / (v6 + (v6 * (v6 - v6))))
        default:
            var v11 string = "\n\t"
            _ = v11
        }
        _ = fi8
    }
    if (42 < -9223372036854775808) {
        for fi12 := 0; fi12 < 8; fi12++ {
            switch "\n\t" {
            case "hello":
                delete(v2, fi12)
            default:
                v6 = v6
            }
            var v13 int32 = 1
            for fi14 := 0; fi14 < 17; fi14++ {
                if true {
                    v15 := GenS2{F0: -1, F1: -1, F2: 333}
                    _ = v15
                } else {
                    var v16 int64 = 370
                    var v17 int = (fi14 << ((uint(v3)) & 63))
                    v18 := GenS2{F0: 1, F1: 42, F2: 1}
                    _ = v16
                    _ = v17
                    _ = v18
                }
                fmt.Println(fi14, v7, v6)
                switch fi12 {
                case 32767:
                    v19 := GenS3{F0: fi14, F1: ("a" + ("" + "foo bar")), F2: ((v4 + float32(float64(v4))) < (v4 - (v4 - v4)))}
                    _ = v19
                default:
                    var v20 uint16 = 0
                    _ = v20
                }
                fmt.Println(v13, v3)
                delete(v2, 775)
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
        switch -1 {
        case -2:
            if v5 {
                if (127 != -16) {
                    if v5 {
                        for fi21 := 0; fi21 < 4; fi21++ {
                            v22 := GenS4{F0: (v3 & (v3 - (v3 >> ((634) & 63)))), F1: 1}
                            _ = fi21
                            _ = v22
                        }
                        if (1 >= int64(v3)) {
                            var v23 uint8 = v3
                            _ = v23
                        }
                        if v5 {
                            for fi24 := 0; fi24 < 13; fi24++ {
                                _ = len(v2)
                                switch fi24 {
                                case 501:
                                    for fi25 := 0; fi25 < 17; fi25++ {
                                        v26 := GenS4{F0: v3, F1: 42}
                                        v5 = true
                                        _ = fi25
                                        _ = v26
                                    }
                                default:
                                    v27 := []int{-9223372036854775808, (fi24 << ((952) & 63))}
                                    _ = v27
                                }
                                _ = fi24
                            }
                            v28 := v1[-9223372036854775808]
                            _ = v28
                        } else {
                            var v29 uint16 = 2
                            _ = v29
                        }
                    }
                } else {
                    switch "hello" {
                    case "foo bar":
                        var v30 bool = (212 != 2)
                        _ = v30
                    case "\n\t":
                        var v31 string = ""
                        _ = v31
                    default:
                        var v32 uint64 = (v7 | (v7 | v7))
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 11; fi33++ {
                        for fi34 := 0; fi34 < 5; fi34++ {
                            if v5 {
                                fmt.Println(v5, v4, v7)
                            }
                            switch "hello" {
                            case "":
                                v35 := GenS5{F0: (786 == 255), F1: v7, F2: 32767}
                                _ = v35
                            case "hello":
                                var v36 bool = v5
                                _ = v36
                            case "foo bar":
                                if v5 {
                                    v4 += 1e10
                                    v37 := GenS1{F0: 4294967295, F1: 228}
                                    _ = v37
                                }
                            default:
                                if ("foo bar" < "\n\t") {
                                    v5 = (!v5)
                                }
                            }
                            switch "foo bar" {
                            case "a":
                                delete(v1, 320)
                            }
                            switch ("hello" + "") {
                            case "hello":
                                if v5 {
                                    v5 = false
                                } else {
                                    if true {
                                        delete(v1, (fi34 - 22))
                                        v7 -= v7
                                        if (!false) {
                                            if (454 >= 1) {
                                                if true {
                                                    if (0 != 96) {
                                                        switch fi33 {
                                                        case 42:
                                                            var v38 uint8 = (v3 | (v3 | 255))
                                                            _ = v38
                                                        case 2147483647:
                                                            var v39 string = "a"
                                                            _ = v39
                                                        case 32768:
                                                            var v40 uint = 640
                                                            _ = v40
                                                        }
                                                        v3 = (v3 >> ((1) & 63))
                                                    } else {
                                                        v1[fi33] = (-(fi33))
                                                        var v41 int16 = -164
                                                        _ = v41
                                                    }
                                                } else {
                                                    var v42 int = 127
                                                    var v43 string = ("hello" + "foo bar")
                                                    delete(v2, -1)
                                                    _ = v42
                                                    _ = v43
                                                }
                                                var v44 bool = (((v4 + 1e10) < v4) && (1000 >= 1000))
                                                _ = v44
                                            } else {
                                                v45 := GenS5{F0: false, F1: (^(v7)), F2: 42}
                                                v6 = (float64(((v4 + 1.5) + v4)) + v6)
                                                if (v5 && (v5 && ((v6 + float64(0.1)) == (v6 * v6)))) {
                                                    fmt.Println(v7, v7, v4)
                                                    switch ("a" + "hello") {
                                                    case "":
                                                        v6 += (3.14159 + v6)
                                                        var v46 int = fi33
                                                        var v47 float64 = v6
                                                        _ = v46
                                                        _ = v47
                                                    default:
                                                        v5 = (-9223372036854775808 >= 782)
                                                    }
                                                    v48 := GenS6{F0: 971, F1: (v7 * v7)}
                                                    _ = v48
                                                } else {
                                                    v49 := GenS2{F0: -1, F1: 42, F2: 619}
                                                    v5 = ((0.1 > float64(v4)) || (0 >= 209))
                                                    switch "" {
                                                    case "\n\t":
                                                        fmt.Println(v6, v4, fi33)
                                                    default:
                                                        if v5 {
                                                            if v5 {
                                                                var v50 int = fi34
                                                                _ = len(v2)
                                                                v7 += (v7 - 1)
                                                                _ = v50
                                                            }
                                                        }
                                                    }
                                                    _ = v49
                                                }
                                                _ = v45
                                            }
                                            if (!(385 > 976)) {
                                                v3 = 21
                                                var v51 uint32 = 0
                                                v52 := []int{(fi33 - (fi34 | 32767))}
                                                _ = v51
                                                _ = v52
                                            }
                                        } else {
                                            var v53 uint8 = (^(v3))
                                            switch "a" {
                                            case "foo bar":
                                                switch int(v3) {
                                                case -2147483648:
                                                    v54 := GenS7{F0: ("a" == ("" + "")), F1: 0, F2: 255}
                                                    _ = v54
                                                }
                                            }
                                            _ = v53
                                        }
                                    }
                                    switch (fi34 << ((42) & 63)) {
                                    case 949:
                                        var v55 uint = 115
                                        _ = v55
                                    case -1:
                                        v4 = -2.25
                                    }
                                    v4 = float32(v6)
                                }
                            case "":
                                v7 += v7
                            case "foo bar":
                                v56 := GenS4{F0: v3, F1: 0}
                                _ = v56
                            }
                            v5 = v5
                            _ = fi34
                        }
                        _ = fi33
                    }
                    var v57 int = 32767
                    _ = v57
                }
                for fi58 := 0; fi58 < 14; fi58++ {
                    v59 := GenS8{F0: v3}
                    _ = fi58
                    _ = v59
                }
            } else {
                v60 := []int{32767, 127, -9223372036854775808}
                if v5 {
                    for fi61 := 0; fi61 < 7; fi61++ {
                        for fi62 := 0; fi62 < 8; fi62++ {
                            v1[fi62] = (fi61 * (fi62 >> ((518) & 63)))
                            _ = fi62
                        }
                        var v63 int = fi61
                        if (42 == 1) {
                            v4 += (((v4 * 1e10) - v4) - ((((v4 - 1e10) * 1e10) * (v4 * ((((v4 + (v4 - float32((((((v6 / 1.0) / float64(v4)) + (1e10 * (float64(((3.14159 + (v4 + v4)) - (v4 * 3.14159))) + 1e10))) - (0.1 - v6)) * (v6 - float64(v4)))))) * (1.5 - float32(((v6 + (v6 / float64(1e10))) - 3.14159)))) * float32(v6)) - 0.0))) / float32(0.1)))
                            fmt.Println(fi61, v63)
                            if len(v60) >= 2 {
                                v60 = v60[1:len(v60)]
                            }
                        }
                        _ = fi61
                        _ = v63
                    }
                    switch "\n\t" {
                    case "a":
                        v5 = v5
                    default:
                        if false {
                            v6 -= v6
                        }
                    }
                    for fi64 := 0; fi64 < 16; fi64++ {
                        delete(v1, 42)
                        _ = fi64
                    }
                } else {
                    v65 := GenS3{F0: -1, F1: "hello", F2: false}
                    _ = v65
                }
                _ = v60
            }
        case -588:
            if false {
                v66 := []int{150, 4294967295, -9223372036854775808}
                v66 = append(v66, -9223372036854775808)
                switch 32768 {
                case 728:
                    var v67 float64 = (float64(v4) * ((((v6 - ((v6 - v6) - float64(v4))) * (float64(3.14159) + v6)) / v6) + float64(v4)))
                    _ = v67
                case 2147483647:
                    var v68 uint = 1
                    _ = v68
                }
                _ = v66
            }
        default:
            if (!(!v5)) {
                fmt.Println(v7, v5)
            }
        }
    }
    v69 := GenS2{F0: -392, F1: 910, F2: 18446744073709551615}
    for fi70 := 0; fi70 < 2; fi70++ {
        switch fi70 {
        case 1:
            var v71 bool = (!false)
            _ = v71
        }
        _ = fi70
    }
    v72 := []int{}
    fmt.Println(v3, v3, v3)
    if ((v4 + v4) > -2.25) {
        _ = len(v1)
        switch "foo bar" {
        case "foo bar":
            switch 1 {
            case -961:
                var v73 uint64 = 1
                _ = v73
            case -880:
                v4 += 1e10
            case 124:
                v6 = (1.5 / v6)
            default:
                for fi74 := 0; fi74 < 11; fi74++ {
                    if v5 {
                        var v75 int64 = 0
                        _ = v75
                    }
                    _ = fi74
                }
            }
        case "\n\t":
            v76 := GenS7{F0: ((true || v5) && v5), F1: -2147483648, F2: 65535}
            _ = v76
        case "":
            v4 += 0.1
        default:
            var v77 uint8 = v3
            _ = v77
        }
    } else {
        v72 = append(v72, 32767)
    }
    v78 := GenS7{F0: (-297 == 1), F1: 4294967295, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v72:", len(v72), v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v7
    _ = v72
    _ = v78
}
