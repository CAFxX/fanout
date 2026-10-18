package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int64; F1 int8 }

func main() {
    v1 := []string{"hello", "\n\t", ("hello" + "foo bar")}
    var v2 int32 = 1
    var v3 uint16 = 65535
    var v4 uint16 = v3
    var v5 uint8 = 255
    v6 := GenS1{F0: (v2 / ((-629) | 1))}
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            if true {
                v3 -= 937
                var v9 uint8 = (v5 / ((37) | 1))
                var v10 int = (fi7 * fi7)
                _ = v10
                _ = v9
            }
            var v11 int8 = 0
            _ = fi8
            _ = v11
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            v2 += 0
            if (false && (true || ((!true) || ("" != "a")))) {
                var v13 uint16 = v3
                var v14 uint8 = (v5 % (((v5 + v5)) | 1))
                var v15 uint16 = (v3 | v13)
                _ = v13
                _ = v14
                _ = v15
            } else {
                v16 := GenS1{F0: v2}
                _ = v16
            }
            _ = fi12
        }
        v3 -= v3
        _ = fi7
    }
    var v17 float64 = 0.0
    switch 0 {
    case 32768:
        v5 = 128
    case -995:
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    }
    if true {
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 float64 = v17
            v20 := GenS1{F0: 2147483647}
            _ = fi18
            _ = v19
            _ = v20
        }
        v2 = (v2 + 42)
    }
    v1 = append(v1, ("a" + ((("a" + "") + "hello") + "\n\t")))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v21 int = -9223372036854775808
    v2 = (v2 >> ((294) & 63))
    for fi22 := 0; fi22 < 8; fi22++ {
        fmt.Println(v21, fi22)
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(v17, v3)
            _ = fi23
        }
        var v24 int = v21
        _ = fi22
        _ = v24
    }
    fmt.Println(v21)
    switch (v21 >> ((18446744073709551615) & 63)) {
    case 32768:
        v4 = (v3 >> ((910) & 63))
    case 9223372036854775807:
        if (uint8(v3) >= 1) {
            var v25 uint8 = v5
            v1 = append(v1, "a")
            v4 = v3
            _ = v25
        } else {
            if (((v2 - int32(v21)) >= (-(v2))) || false) {
                switch (v21 | 32767) {
                case 128:
                    for fi26 := 0; fi26 < 12; fi26++ {
                        v21 -= (fi26 << ((4294967295) & 63))
                        fmt.Println(v2, v21, v3)
                        _ = fi26
                    }
                case 596:
                    switch v21 {
                    case 42:
                        if true {
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                var v27 uint8 = v5
                                switch 32768 {
                                case 4294967295:
                                    v28 := GenS2{F0: (!(42 == 2)), F1: 2147483647, F2: v27}
                                    _ = v28
                                default:
                                    switch 9223372036854775807 {
                                    case 42:
                                        v27 = (v27 & 16)
                                    case 0:
                                        switch ("a" + ("hello" + "a")) {
                                        case "":
                                            for fi29 := 0; fi29 < 9; fi29++ {
                                                _ = len(v1)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                var v30 int32 = (-(v2))
                                                _ = fi29
                                                _ = v30
                                            }
                                        }
                                    case 239:
                                        if (((-128 < 75) && ("\n\t" == "foo bar")) && (42 == 65535)) {
                                            v31 := GenS3{F0: v5, F1: true}
                                            _ = v31
                                        }
                                    default:
                                        switch (-(v21)) {
                                        case 338:
                                            if (-156 > 0) {
                                                switch ("hello" + ("\n\t" + ("a" + ("foo bar" + "a")))) {
                                                case "\n\t":
                                                    switch (^(v21)) {
                                                    case 2147483647:
                                                        switch (v21 ^ v21) {
                                                        case 361:
                                                            v1 = append(v1, "\n\t")
                                                        case -803:
                                                            fmt.Println(v3, v21)
                                                        case 128:
                                                            v21 = int(v2)
                                                        }
                                                    case -234:
                                                        v17 = v17
                                                    case 127:
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v32 := GenS1{F0: v2}
                                                            _ = v32
                                                        case "foo bar":
                                                            for fi33 := 0; fi33 < 7; fi33++ {
                                                                switch v21 {
                                                                case -1:
                                                                    var v34 float32 = 1.5
                                                                    _ = v34
                                                                }
                                                                var v35 uint8 = 89
                                                                v36 := GenS4{F0: v2}
                                                                _ = fi33
                                                                _ = v35
                                                                _ = v36
                                                            }
                                                        default:
                                                            v2 += v2
                                                        }
                                                    default:
                                                        v5 += 2
                                                    }
                                                default:
                                                    for fi37 := 0; fi37 < 3; fi37++ {
                                                        var v38 float64 = v17
                                                        if ("\n\t" > "\n\t") {
                                                            for fi39 := 0; fi39 < 11; fi39++ {
                                                                v3 += 42
                                                                fmt.Println(fi39, v5)
                                                                var v40 int32 = (v2 >> ((18446744073709551615) & 63))
                                                                _ = fi39
                                                                _ = v40
                                                            }
                                                            fmt.Println(v5, v17, v4)
                                                            var v41 string = "foo bar"
                                                            _ = v41
                                                        }
                                                        _ = fi37
                                                        _ = v38
                                                    }
                                                }
                                                var v42 bool = (!false)
                                                _ = v42
                                            } else {
                                                if (!("\n\t" == "hello")) {
                                                    fmt.Println(v2)
                                                    var v43 float32 = 0.1
                                                    _ = v43
                                                }
                                                var v44 int64 = -1
                                                var v45 uint16 = 890
                                                _ = v44
                                                _ = v45
                                            }
                                        case 32767:
                                            var v46 float64 = (v17 / float64((0.0 - 1e10)))
                                            _ = v46
                                        default:
                                            v3 -= (v3 % ((255) | 1))
                                        }
                                    }
                                }
                                switch v21 {
                                case -2147483648:
                                    v1 = append(v1, "foo bar")
                                }
                                _ = v27
                            }
                            for fi47 := 0; fi47 < 19; fi47++ {
                                fmt.Println(v3, v21, v3)
                                var v48 float64 = (1e10 * (1e10 - v17))
                                for fi49 := 0; fi49 < 7; fi49++ {
                                    if true {
                                        fmt.Println(v5, v21, v48)
                                    }
                                    var v50 uint32 = 2
                                    v3 -= 32767
                                    _ = fi49
                                    _ = v50
                                }
                                _ = fi47
                                _ = v48
                            }
                        } else {
                            if true {
                                for fi51 := 0; fi51 < 8; fi51++ {
                                    for fi52 := 0; fi52 < 4; fi52++ {
                                        if (2 <= 255) {
                                            var v53 int16 = 1
                                            _ = v53
                                        }
                                        v54 := GenS2{F0: (true && false), F1: 2, F2: (v5 + (v5 / ((v5) | 1)))}
                                        var v55 float32 = 1.5
                                        _ = fi52
                                        _ = v54
                                        _ = v55
                                    }
                                    var v56 int64 = 42
                                    v21 += (-(v21))
                                    _ = fi51
                                    _ = v56
                                }
                            } else {
                                for fi57 := 0; fi57 < 16; fi57++ {
                                    v21 = fi57
                                    if len(v1) > 0 {
                                        v1[0] = "a"
                                    }
                                    _ = fi57
                                }
                            }
                        }
                    }
                }
                fmt.Println(v17, v21)
            }
            for fi58 := 0; fi58 < 17; fi58++ {
                if (!(42 > 946)) {
                    v59 := GenS1{F0: (v2 % ((v2) | 1))}
                    v60 := GenS5{F0: 4294967295}
                    switch "" {
                    case "":
                        v3 -= v4
                    case "hello":
                        var v61 int16 = -32768
                        _ = v61
                    default:
                        if false {
                            switch (fi58 >> ((42) & 63)) {
                            case 42:
                                fmt.Println(fi58, v21, v3)
                            case 1:
                                var v62 uint16 = 255
                                _ = v62
                            case 0:
                                v4 -= v4
                            default:
                                if (255 != v4) {
                                    if false {
                                        v63 := GenS6{F0: (fi58 >> ((18446744073709551615) & 63))}
                                        for fi64 := 0; fi64 < 12; fi64++ {
                                            if false {
                                                _ = len(v1)
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            } else {
                                                fmt.Println(v3, v21, fi64)
                                            }
                                            v65 := GenS1{F0: v2}
                                            fmt.Println(v5, v5, fi64)
                                            v1 = append(v1, "")
                                            _ = fi64
                                            _ = v65
                                        }
                                        _ = v63
                                    }
                                } else {
                                    v1 = append(v1, "\n\t")
                                }
                            }
                            switch (v21 / ((-89) | 1)) {
                            case 4294967295:
                                fmt.Println(v3, v3)
                            case -172:
                                v5 -= uint8(v3)
                            case -283:
                                var v66 uint = 42
                                _ = v66
                            }
                        }
                    }
                    _ = v59
                    _ = v60
                } else {
                    var v67 int16 = 1
                    var v68 string = "foo bar"
                    v69 := GenS7{F0: (1e10 >= ((v17 + v17) + float64(1.5)))}
                    _ = v67
                    _ = v68
                    _ = v69
                }
                for fi70 := 0; fi70 < 10; fi70++ {
                    var v71 bool = (!true)
                    v5 -= (v5 * v5)
                    if v71 {
                        if (3.14159 == 1e10) {
                            switch 0 {
                            case 0:
                                if v71 {
                                    switch (fi70 & 32768) {
                                    case 2147483647:
                                        if true {
                                            v71 = (127 >= -53)
                                            var v72 int8 = -1
                                            _ = v72
                                        }
                                    case -150:
                                        var v73 uint = 30
                                        _ = v73
                                    }
                                    v17 += -2.25
                                    v5 = (v5 ^ v5)
                                } else {
                                    v74 := GenS2{F0: v71, F1: 2147483647, F2: (v5 * (v5 / (((v5 | (v5 * (v5 + (v5 | 42))))) | 1)))}
                                    _ = v74
                                }
                            case 2147483647:
                                v21 = v21
                                fmt.Println(fi58, v17, v2)
                            default:
                                if (false || (!v71)) {
                                    if v71 {
                                        if v71 {
                                            v3 = v4
                                            if v71 {
                                                v17 += (1.5 + v17)
                                            } else {
                                                fmt.Println(v3, v17, v3)
                                                var v75 float64 = 0.1
                                                _ = v75
                                            }
                                            if (!v71) {
                                                var v76 int8 = -128
                                                _ = v76
                                            } else {
                                                fmt.Println(v2, v2)
                                                if v71 {
                                                    var v77 float64 = (((v17 + ((v17 * (v17 / 1.5)) - v17)) / ((1e10 + v17) - (v17 * (v17 * -2.25)))) + 0.1)
                                                    v17 -= v17
                                                    _ = v77
                                                } else {
                                                    v1 = append(v1, "hello")
                                                    if len(v1) > 0 {
                                                        v1[0] = ("hello" + ("a" + ("\n\t" + ("\n\t" + ""))))
                                                    }
                                                }
                                            }
                                        } else {
                                            v21 -= -880
                                        }
                                    } else {
                                        v78 := GenS8{F0: 9223372036854775807, F1: -113}
                                        _ = v78
                                    }
                                } else {
                                    var v79 int = 32768
                                    var v80 uint16 = 255
                                    v81 := GenS4{F0: 2147483647}
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                }
                            }
                        }
                        switch "a" {
                        case "\n\t":
                            v71 = (true || v71)
                        default:
                            var v82 float32 = -2.25
                            var v83 uint = 255
                            _ = v82
                            _ = v83
                        }
                        switch v21 {
                        case -1:
                            if len(v1) > 0 {
                                v1[0] = ("foo bar" + "")
                            }
                        case -2147483648:
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                        case 0:
                            var v84 uint8 = (^(v5))
                            _ = v84
                        }
                    }
                    _ = fi70
                    _ = v71
                }
                v3 = 65535
                _ = fi58
            }
        }
    default:
        v4 += (v3 ^ (v3 % ((v4) | 1)))
    }
    switch (^(v21)) {
    case 42:
        var v85 int64 = -1
        _ = v85
    case -1:
        v86 := GenS2{F0: (!(!false)), F1: 220, F2: v5}
        _ = v86
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
