package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int32; F1 int16; F2 int16 }
type GenS3 struct { F0 bool; F1 uint16; F2 int16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 uint16 }

func main() {
    var v1 bool = (!false)
    v2 := []uint32{}
    var v3 bool = true
    v4 := make(map[string]string)
    var v5 uint8 = 224
    var v6 int8 = 127
    v7 := GenS1{F0: 4294967295, F1: 42}
    var v8 uint8 = (v5 << ((255) & 63))
    switch ("foo bar" + ((("\n\t" + "foo bar") + "\n\t") + ("\n\t" + "a"))) {
    case "hello":
        var v9 int8 = v6
        _ = v9
    case "a":
        switch ("foo bar" + "\n\t") {
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            var v10 float64 = (0.1 / 1.0)
            _ = v10
        }
    case "":
        var v11 int16 = -1
        _ = v11
    default:
        var v12 bool = ((!(25 < 42)) && (!v3))
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v1, v1)
        _ = fi13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 uint64 = 0
        _ = fi14
        _ = v15
    }
    switch -2147483648 {
    case -371:
        switch 4294967295 {
        case 367:
            for fi16 := 0; fi16 < 8; fi16++ {
                v3 = (!false)
                fmt.Println(v3)
                _ = fi16
            }
        case -987:
            switch -9223372036854775808 {
            case -367:
                var v17 int32 = 1
                _ = v17
            case 1:
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 float32 = (1e10 * float32(0.0))
                    _ = fi18
                    _ = v19
                }
            case 128:
                v1 = (v1 || v1)
            default:
                v1 = false
            }
        }
    default:
        for fi20 := 0; fi20 < 18; fi20++ {
            v3 = false
            _ = fi20
        }
    }
    if v1 {
        v21 := GenS1{F0: 1, F1: 577}
        v6 = (v6 - v6)
        _ = v21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 uint16 = 32767
        switch ("foo bar" + ("" + ("hello" + "\n\t"))) {
        case "\n\t":
            if false {
                v6 = (v6 >> ((498) & 63))
                var v24 float64 = 0.1
                if true {
                    v4["\n\t"] = ""
                    v25 := GenS2{F0: -1, F1: 936, F2: -590}
                    for fi26 := 0; fi26 < 13; fi26++ {
                        delete(v4, "foo bar")
                        v27 := v4["a"]
                        _ = fi26
                        _ = v27
                    }
                    _ = v25
                } else {
                    v28 := GenS1{F0: 2, F1: 1}
                    v3 = v1
                    _ = v28
                }
                _ = v24
            } else {
                v23 = v23
                var v29 int8 = v6
                _ = v29
            }
        }
        v1 = ((v23 >> ((971) & 63)) != (v23 * 880))
        _ = fi22
        _ = v23
    }
    var v30 uint32 = 767
    v31 := GenS3{F0: (!v3), F1: 115, F2: 42}
    v32 := GenS1{F0: v30, F1: 483}
    if (false || v3) {
        if (v3 && v3) {
            fmt.Println(v3)
            switch "hello" {
            case "foo bar":
                var v33 uint16 = 65535
                _ = v33
            case "a":
                if (607 <= 788) {
                    v34 := GenS4{F0: 1000, F1: v6}
                    var v35 int8 = v6
                    if v1 {
                        if true {
                            for fi36 := 0; fi36 < 20; fi36++ {
                                var v37 uint32 = v30
                                _ = len(v4)
                                switch (fi36 + fi36) {
                                case -1:
                                    if (!(0 == 42)) {
                                        var v38 uint64 = 679
                                        _ = v38
                                    }
                                default:
                                    for fi39 := 0; fi39 < 4; fi39++ {
                                        var v40 int = -2147483648
                                        fmt.Println(v6, v37)
                                        _ = fi39
                                        _ = v40
                                    }
                                }
                                _ = fi36
                                _ = v37
                            }
                        }
                        fmt.Println(v35, v8, v6)
                        switch "" {
                        case "foo bar":
                            v4[("" + "")] = ("hello" + ("\n\t" + "foo bar"))
                        default:
                            var v41 float32 = 3.14159
                            _ = v41
                        }
                    } else {
                        var v42 uint = 18446744073709551615
                        var v43 uint8 = v5
                        var v44 float64 = (3.14159 + ((0.1 + 3.14159) - ((3.14159 - (0.1 * (1.5 - float64((3.14159 - ((0.1 + 3.14159) + float32(-2.25))))))) * 1.5)))
                        _ = v42
                        _ = v43
                        _ = v44
                    }
                    _ = v34
                    _ = v35
                } else {
                    for fi45 := 0; fi45 < 8; fi45++ {
                        v5 -= v8
                        _ = fi45
                    }
                }
            case "\n\t":
                for fi46 := 0; fi46 < 8; fi46++ {
                    if false {
                        v47 := GenS5{F0: -187, F1: fi46}
                        _ = v47
                    } else {
                        var v48 uint16 = 2
                        _ = v48
                    }
                    _ = fi46
                }
            default:
                if (!v3) {
                    for fi49 := 0; fi49 < 8; fi49++ {
                        if false {
                            v1 = v1
                            if (0.0 > (-2.25 / (float32(((0.1 / ((float64(0.1) - 0.0) - ((1.5 * (1e10 * float64(0.0))) / (0.1 - 0.0)))) + 0.1)) - 0.0))) {
                                var v50 int = (fi49 + 9223372036854775807)
                                _ = v50
                            } else {
                                if (!(1 != v8)) {
                                    if v1 {
                                        if (false || (883 <= -173)) {
                                            if v3 {
                                                for fi51 := 0; fi51 < 13; fi51++ {
                                                    var v52 int32 = 1
                                                    _ = fi51
                                                    _ = v52
                                                }
                                            } else {
                                                switch "a" {
                                                case "foo bar":
                                                    v2 = append(v2, 994)
                                                case "a":
                                                    v2 = append(v2, 541)
                                                }
                                            }
                                            v53 := GenS6{F0: -1, F1: 1}
                                            v4[""] = ("\n\t" + ("\n\t" + "foo bar"))
                                            _ = v53
                                        } else {
                                            delete(v4, "a")
                                            v54 := GenS1{F0: (v30 - 42), F1: 9223372036854775807}
                                            _ = v54
                                        }
                                        for fi55 := 0; fi55 < 20; fi55++ {
                                            if len(v2) > 0 {
                                                v2[0] = v30
                                            }
                                            switch "\n\t" {
                                            case "foo bar":
                                                v56 := GenS1{F0: 0, F1: 0}
                                                _ = v56
                                            }
                                            _ = fi55
                                        }
                                        switch "a" {
                                        case "a":
                                            var v57 bool = (!v1)
                                            _ = v57
                                        default:
                                            v58 := GenS4{F0: 0, F1: (v6 ^ v6)}
                                            _ = v58
                                        }
                                    } else {
                                        v6 -= v6
                                        var v59 int8 = 127
                                        _ = v59
                                    }
                                    var v60 int32 = 1
                                    _ = v60
                                }
                                v1 = ("foo bar" > "")
                            }
                            v61 := GenS7{F0: 1}
                            _ = v61
                        } else {
                            switch ("" + "\n\t") {
                            case "\n\t":
                                for fi62 := 0; fi62 < 10; fi62++ {
                                    v8 = uint8(v30)
                                    var v63 int16 = -32768
                                    v6 = v6
                                    var v64 float64 = (3.14159 + 1.5)
                                    var v65 int32 = 42
                                    _ = fi62
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                            case "hello":
                                for fi66 := 0; fi66 < 14; fi66++ {
                                    var v67 uint8 = (v8 + v5)
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = fi66
                                    _ = v67
                                }
                            case "":
                                switch 128 {
                                case 2147483647:
                                    if true {
                                        if true {
                                            for fi68 := 0; fi68 < 9; fi68++ {
                                                v6 = -38
                                                var v69 uint64 = 4294967295
                                                fmt.Println(v8, v5, v5)
                                                _ = fi68
                                                _ = v69
                                            }
                                        }
                                        if false {
                                            if v3 {
                                                fmt.Println(v30)
                                                var v70 int16 = 42
                                                var v71 int32 = int32(v30)
                                                _ = v70
                                                _ = v71
                                            }
                                        }
                                    } else {
                                        v4["foo bar"] = "foo bar"
                                        var v72 int8 = v6
                                        if (0.0 >= ((float64(1.5) - 1.5) - float64(1e10))) {
                                            var v73 int64 = 321
                                            _ = v73
                                        }
                                        _ = v72
                                    }
                                case -9223372036854775808:
                                    for fi74 := 0; fi74 < 19; fi74++ {
                                        if v3 {
                                            var v75 bool = (18446744073709551615 == 0)
                                            v76 := GenS2{F0: 890, F1: 0, F2: -493}
                                            _ = v75
                                            _ = v76
                                        }
                                        var v77 int32 = -1
                                        _ = fi74
                                        _ = v77
                                    }
                                case 6:
                                    v78 := GenS8{F0: (v30 << ((0) & 63))}
                                    _ = v78
                                }
                            }
                            if (((-2.25 - ((1.5 * -2.25) - (1e10 - 1e10))) - ((float32((0.1 / float64((((float32(-2.25) * (3.14159 * (float32((float64((0.1 + (0.1 + 0.0))) + float64(0.1))) - float32((3.14159 - ((1e10 * -2.25) * (1.5 / -2.25))))))) - (((3.14159 + 1.5) - float32((0.0 * 3.14159))) / 1.0)) * 1e10)))) + (0.0 / (0.1 + 0.1))) - 1e10)) >= (3.14159 - float32(0.1))) {
                                for fi79 := 0; fi79 < 20; fi79++ {
                                    switch fi79 {
                                    case 658:
                                        v6 = v6
                                    case 42:
                                        var v80 uint32 = (v30 << ((1) & 63))
                                        fmt.Println(v80, fi79)
                                        v81 := GenS9{F0: 519, F1: 32767}
                                        _ = v80
                                        _ = v81
                                    case 595:
                                        fmt.Println(v1)
                                    }
                                    fmt.Println(fi49)
                                    _ = fi79
                                }
                                v1 = v1
                                v1 = false
                            }
                            v82 := GenS7{F0: 32767}
                            _ = v82
                        }
                        _ = fi49
                    }
                    for fi83 := 0; fi83 < 14; fi83++ {
                        var v84 int16 = 32767
                        _ = fi83
                        _ = v84
                    }
                }
            }
        } else {
            v85 := GenS6{F0: 297, F1: -334}
            if ((^(v8)) >= v8) {
                v6 = v6
            } else {
                switch "hello" {
                case "\n\t":
                    switch 95 {
                    case -9223372036854775808:
                        delete(v4, ("foo bar" + (("" + "foo bar") + "\n\t")))
                    case 865:
                        var v86 uint16 = 85
                        _ = v86
                    }
                case "a":
                    _ = len(v2)
                }
                if (((18446744073709551615 > 0) || v1) && (0.0 == 1.5)) {
                    v5 = (v5 % ((v8) | 1))
                    var v87 bool = (!(3.14159 <= 1.5))
                    _ = v87
                }
                if false {
                    v88 := GenS9{F0: 9223372036854775807, F1: 255}
                    v4["a"] = ""
                    fmt.Println(v1, v5)
                    _ = v88
                } else {
                    for fi89 := 0; fi89 < 4; fi89++ {
                        v6 = (v6 & (v6 - -128))
                        if v1 {
                            for fi90 := 0; fi90 < 8; fi90++ {
                                if v1 {
                                    if (0 == 42) {
                                        _ = len(v2)
                                    }
                                }
                                var v91 uint32 = 42
                                _ = fi90
                                _ = v91
                            }
                            if (int16(fi89) == -473) {
                                var v92 uint64 = 718
                                delete(v4, "hello")
                                _ = v92
                            }
                            switch 4294967295 {
                            case -2147483648:
                                v93 := v4["foo bar"]
                                _ = v93
                            case 342:
                                for fi94 := 0; fi94 < 7; fi94++ {
                                    var v95 int16 = int16(v30)
                                    _ = fi94
                                    _ = v95
                                }
                            default:
                                var v96 int16 = 1
                                _ = v96
                            }
                        }
                        _ = fi89
                    }
                }
            }
            for fi97 := 0; fi97 < 13; fi97++ {
                var v98 int16 = 1
                _ = fi97
                _ = v98
            }
            _ = v85
        }
    } else {
        var v99 bool = v1
        v100 := GenS2{F0: -2147483648, F1: 24, F2: 1}
        _ = v100
        _ = v99
    }
    for fi101 := 0; fi101 < 8; fi101++ {
        if (true && v1) {
            v6 = -1
        }
        _ = fi101
    }
    var v102 int16 = -184
    var v103 bool = false
    for fi104 := 0; fi104 < 11; fi104++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi104
    }
    var v105 int32 = 42
    if (v3 && v103) {
        v102 = v102
        var v106 string = "foo bar"
        _ = v106
    }
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v105:", v105)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v102
    _ = v103
    _ = v105
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
