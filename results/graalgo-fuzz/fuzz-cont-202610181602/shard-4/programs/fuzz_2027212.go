package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 float64 = ((1e10 * -2.25) - 3.14159)
    if ((true && false) && false) {
        fmt.Println(v3)
    } else {
        if (1.5 != (v3 + (v3 + v3))) {
            v3 = (v3 / 0.1)
        } else {
            var v4 uint32 = 0
            for fi5 := 0; fi5 < 13; fi5++ {
                var v6 float32 = float32(v3)
                _ = fi5
                _ = v6
            }
            v7 := GenS1{F0: 32767, F1: (v4 % ((v4) | 1))}
            _ = v4
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 1000, F1: 2147483647}
        _ = fi8
        _ = v9
    }
    v3 = v3
    v10 := v2[2]
    if (!(127 <= 0)) {
        if false {
            for fi11 := 0; fi11 < 17; fi11++ {
                v12 := GenS2{F0: v10, F1: 32767}
                v13 := GenS1{F0: 588, F1: 2}
                var v14 uint32 = 952
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            var v15 uint = 0
            _ = v15
        }
        v1[65535] = v10
        fmt.Println(v3, v3)
    } else {
        v16 := v2[1]
        for fi17 := 0; fi17 < 15; fi17++ {
            if ((!false) || false) {
                v18 := GenS1{F0: 42, F1: 65535}
                _ = v18
            }
            if v16 {
                switch (fi17 / ((fi17) | 1)) {
                case 58:
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v20 := GenS2{F0: true, F1: 32767}
                        v16 = (1.5 >= 0.1)
                        _ = fi19
                        _ = v20
                    }
                default:
                    v16 = ("" > "\n\t")
                }
                var v21 int8 = 127
                _ = v21
            } else {
                var v22 uint16 = 0
                fmt.Println(v3)
                _ = v22
            }
            _ = fi17
        }
        _ = v16
    }
    v23 := GenS3{F0: 0}
    v24 := []int{-458, 127, -2147483648}
    if false {
        v3 += v3
        v1[65535] = v10
        if len(v24) > 0 {
            v24[0] = 0
        }
    } else {
        switch -886 {
        case -926:
            for fi25 := 0; fi25 < 8; fi25++ {
                if (729 > 0) {
                    fmt.Println(v10)
                } else {
                    if (1 > 893) {
                        v3 += (3.14159 - v3)
                        var v26 int8 = -69
                        _ = v26
                    }
                    v3 -= (3.14159 * 3.14159)
                    v10 = (v10 && (2 == 694))
                }
                _ = fi25
            }
        default:
            var v27 uint8 = 42
            _ = v27
        }
        v28 := GenS2{F0: (2 != 2), F1: 42}
        _ = v28
    }
    fmt.Println(v3, v10)
    v29 := GenS4{F0: ("" + "\n\t"), F1: (((false || (850 == 42)) && (0 < 42)) || (42 >= 760))}
    v30 := GenS1{F0: 1, F1: 4294967295}
    fmt.Println(v3)
    v31 := GenS4{F0: ("hello" + ("" + "\n\t")), F1: (!((("\n\t" + "") + "foo bar") < "hello"))}
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    var v32 uint = 2
    if false {
        v33 := GenS3{F0: 1}
        if true {
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 bool = (!v10)
                _ = fi34
                _ = v35
            }
        }
        _ = v33
    } else {
        for fi36 := 0; fi36 < 20; fi36++ {
            for fi37 := 0; fi37 < 14; fi37++ {
                v10 = ((v10 && (1 != 416)) || true)
                _ = fi37
            }
            _ = fi36
        }
        if v10 {
            var v38 float32 = 1.5
            var v39 string = (("hello" + "a") + "")
            if (v10 && (276 == 1)) {
                if (("a" + v39) != "foo bar") {
                    switch 0 {
                    case 1:
                        for fi40 := 0; fi40 < 7; fi40++ {
                            if v10 {
                                var v41 uint = (v32 + v32)
                                switch fi40 {
                                case 307:
                                    var v42 uint64 = 4294967295
                                    _ = v42
                                default:
                                    v43 := GenS1{F0: 0, F1: 652}
                                    _ = v43
                                }
                                _ = v41
                            } else {
                                v10 = (2147483648 < 1)
                                v32 -= (v32 ^ (v32 >> (((v32 | (v32 - 913))) & 63)))
                            }
                            _ = fi40
                        }
                    case 128:
                        switch v39 {
                        case "foo bar":
                            fmt.Println(v38, v38)
                        default:
                            switch 42 {
                            case -9223372036854775808:
                                var v44 uint32 = 1
                                _ = v44
                            case 32768:
                                var v45 int64 = 0
                                _ = v45
                            }
                        }
                    case 37:
                        var v46 uint64 = 2
                        _ = v46
                    default:
                        var v47 bool = (0 > 99)
                        _ = v47
                    }
                    var v48 int16 = 42
                    v3 = v3
                    _ = v48
                } else {
                    for fi49 := 0; fi49 < 8; fi49++ {
                        if (-1 < 1) {
                            switch (fi49 * (fi49 << ((v32) & 63))) {
                            case 4294967295:
                                for fi50 := 0; fi50 < 10; fi50++ {
                                    var v51 uint64 = 1
                                    var v52 uint8 = 255
                                    _ = fi50
                                    _ = v51
                                    _ = v52
                                }
                            default:
                                if true {
                                    fmt.Println(v32, v39)
                                    var v53 int64 = 1
                                    _ = v53
                                }
                            }
                            var v54 uint32 = 4294967295
                            if (-9223372036854775808 > -1) {
                                for fi55 := 0; fi55 < 1; fi55++ {
                                    var v56 int16 = 33
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    v24 = append(v24, (fi49 / ((-270) | 1)))
                                    _ = fi55
                                    _ = v56
                                }
                            } else {
                                for fi57 := 0; fi57 < 14; fi57++ {
                                    v58 := GenS4{F0: v39, F1: ((v54 & uint32(v32)) < (v54 ^ v54))}
                                    if v10 {
                                        fmt.Println(fi57, v10)
                                        var v59 float32 = float32(float64(v38))
                                        _ = v59
                                    }
                                    v60 := GenS4{F0: ("\n\t" + v39), F1: ("hello" < ("\n\t" + ("a" + "a")))}
                                    _ = fi57
                                    _ = v58
                                    _ = v60
                                }
                                _ = len(v2)
                                if false {
                                    fmt.Println(v38, v39, v38)
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    var v61 int16 = 1000
                                    _ = v61
                                } else {
                                    var v62 uint32 = (v54 / ((v54) | 1))
                                    delete(v2, v62)
                                    _ = v62
                                }
                            }
                            _ = v54
                        }
                        _ = fi49
                    }
                    for fi63 := 0; fi63 < 18; fi63++ {
                        switch v39 {
                        case "hello":
                            var v64 uint8 = 243
                            _ = v64
                        }
                        if (!false) {
                            v65 := GenS4{F0: v39, F1: (!((-(fi63)) >= (fi63 & fi63)))}
                            if ((100 >= 5) || false) {
                                fmt.Println(fi63, v3)
                                switch fi63 {
                                case -59:
                                    if (!((v10 && (!(!true))) || v10)) {
                                        for fi66 := 0; fi66 < 20; fi66++ {
                                            var v67 int64 = 0
                                            v38 -= v38
                                            _ = fi66
                                            _ = v67
                                        }
                                        for fi68 := 0; fi68 < 18; fi68++ {
                                            v3 -= -2.25
                                            _ = fi68
                                        }
                                        v10 = ((v32 >> ((v32) & 63)) < v32)
                                    }
                                case 604:
                                    v39 = "foo bar"
                                }
                                switch (fi63 >> ((308) & 63)) {
                                case 482:
                                    switch fi63 {
                                    case -301:
                                        var v69 uint8 = 230
                                        _ = v69
                                    case 32767:
                                        var v70 int8 = 127
                                        _ = v70
                                    }
                                case -9223372036854775808:
                                    switch "a" {
                                    case "foo bar":
                                        fmt.Println(v10)
                                    default:
                                        v71 := GenS5{F0: fi63, F1: -1}
                                        _ = v71
                                    }
                                }
                            }
                            _ = v65
                        }
                        v3 = 1.5
                        _ = fi63
                    }
                    if (v10 && true) {
                        v2[2147483648] = true
                        fmt.Println(v39, v38)
                        var v72 float32 = ((v38 / (v38 - ((v38 - 1.5) + v38))) + ((float32(v3) + float32(v3)) / v38))
                        _ = v72
                    } else {
                        if len(v24) > 0 {
                            v24[0] = -268
                        }
                        if len(v24) > 0 {
                            v24[0] = -1
                        }
                        _ = len(v1)
                    }
                }
                var v73 int = 1
                _ = v73
            }
            _ = v38
            _ = v39
        } else {
            if v10 {
                if true {
                    switch "\n\t" {
                    case "":
                        v32 = 65535
                    default:
                        if (1.5 >= (-2.25 + 0.0)) {
                            v3 += v3
                        } else {
                            v74 := GenS6{F0: 65535}
                            _ = v74
                        }
                    }
                } else {
                    switch int(v32) {
                    case 9223372036854775807:
                        v10 = false
                    case 1:
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                    case 32767:
                        switch int(v32) {
                        case 662:
                            var v75 uint16 = 420
                            _ = v75
                        case 2147483647:
                            fmt.Println(v3)
                        case 1:
                            v32 += v32
                        default:
                            for fi76 := 0; fi76 < 4; fi76++ {
                                v77 := GenS5{F0: 32768, F1: -28}
                                if (true || v10) {
                                    for fi78 := 0; fi78 < 14; fi78++ {
                                        if true {
                                            v3 -= ((-2.25 * v3) * 0.0)
                                            if v10 {
                                                v79 := v2[2]
                                                _ = v79
                                            }
                                        } else {
                                            var v80 int64 = -966
                                            _ = v80
                                        }
                                        _ = fi78
                                    }
                                }
                                fmt.Println(fi76)
                                _ = fi76
                                _ = v77
                            }
                        }
                    default:
                        v81 := GenS3{F0: 4294967295}
                        _ = v81
                    }
                    var v82 uint = 1
                    for fi83 := 0; fi83 < 18; fi83++ {
                        fmt.Println(fi83)
                        _ = fi83
                    }
                    _ = v82
                }
            } else {
                var v84 uint = v32
                _ = v84
            }
            for fi85 := 0; fi85 < 13; fi85++ {
                if (-1 < -9223372036854775808) {
                    if ((v3 / v3) == v3) {
                        var v86 int8 = -1
                        for fi87 := 0; fi87 < 6; fi87++ {
                            fmt.Println(v3, v32)
                            v3 = (((1e10 - v3) + v3) * float64(1.5))
                            var v88 uint32 = 356
                            _ = fi87
                            _ = v88
                        }
                        _ = v86
                    } else {
                        v10 = (((-2.25 * ((v3 + v3) + v3)) * (((1.5 / -2.25) * v3) - -2.25)) <= (v3 * v3))
                        var v89 int64 = -1
                        _ = v89
                    }
                    switch ("" + "foo bar") {
                    case "foo bar":
                        v32 = v32
                    case "":
                        v10 = (!((!v10) || ((true || v10) && (-32768 >= -1))))
                    case "a":
                        v3 = ((1.5 * (3.14159 + ((v3 + v3) + (1.5 / v3)))) - (v3 + v3))
                    default:
                        v32 -= v32
                    }
                    if (!v10) {
                        v10 = false
                    }
                }
                _ = fi85
            }
        }
    }
    v3 += float64(float32(float64(1.5)))
    for fi90 := 0; fi90 < 13; fi90++ {
        v2[4294967295] = (151 == 9223372036854775807)
        for fi91 := 0; fi91 < 13; fi91++ {
            var v92 int16 = -64
            var v93 uint16 = 32767
            _ = fi91
            _ = v92
            _ = v93
        }
        if (42 <= 2147483647) {
            switch "a" {
            case "a":
                v32 += v32
            case "foo bar":
                v94 := v2[42]
                _ = v94
            }
            for fi95 := 0; fi95 < 15; fi95++ {
                v96 := GenS7{F0: 18446744073709551615}
                _ = fi95
                _ = v96
            }
        }
        _ = fi90
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
