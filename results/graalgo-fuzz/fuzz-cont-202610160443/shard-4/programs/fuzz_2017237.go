package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint8; F1 string; F2 uint64 }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint16; F1 uint32 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int64; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 255
    var v3 uint32 = 305
    v4 := []string{("foo bar" + (("foo bar" + ("foo bar" + "foo bar")) + ("a" + "foo bar")))}
    var v5 uint32 = 2147483648
    var v6 float32 = 1.5
    v7 := GenS1{F0: 0, F1: 2}
    var v8 bool = (((!(-2147483648 != 1)) && (float64((v6 / v6)) < (0.0 + (float64(3.14159) + 1.5)))) && (-204 == 4294967295))
    delete(v1, v5)
    if (-32768 == -1) {
        fmt.Println(v8, v3, v5)
    } else {
        if (!((!false) || v8)) {
            v5 = 517
            v9 := GenS2{F0: v2, F1: ("\n\t" + ""), F2: 1}
            _ = v9
        }
        fmt.Println(v5)
        v10 := GenS2{F0: (v2 + (v2 * v2)), F1: "\n\t", F2: 42}
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 2; fi12++ {
            v8 = false
            fmt.Println(v2, v2)
            var v13 int16 = -32768
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 int = 579
    v6 = 3.14159
    fmt.Println(v8, v2, v2)
    if (!(!false)) {
        v5 = v5
        if v8 {
            v15 := v1[(v5 & (v3 ^ (^(v5))))]
            fmt.Println(v2, v14)
            if len(v4) > 0 {
                v4[0] = (("hello" + "foo bar") + "hello")
            }
            _ = v15
        } else {
            if v8 {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v16 := GenS2{F0: (^(v2)), F1: "\n\t", F2: 910}
                delete(v1, v3)
                _ = v16
            }
            v17 := GenS1{F0: (v14 % ((1) | 1)), F1: 434}
            _ = v17
        }
    } else {
        var v18 float32 = 0.0
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 14; fi19++ {
            if v8 {
                fmt.Println(v3, v8)
                fmt.Println(v2, v8, v8)
            } else {
                var v20 uint16 = 885
                _ = v20
            }
            for fi21 := 0; fi21 < 2; fi21++ {
                v6 = 1e10
                delete(v1, 2147483647)
                _ = fi21
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 uint64 = 18446744073709551615
                v24 := GenS1{F0: fi22, F1: 1}
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi19
        }
        var v25 bool = true
        for fi26 := 0; fi26 < 8; fi26++ {
            switch (fi26 >> ((312) & 63)) {
            case 128:
                _ = len(v4)
            case 9223372036854775807:
                for fi27 := 0; fi27 < 7; fi27++ {
                    v28 := GenS3{F0: v14, F1: int8(v2), F2: 42}
                    v14 = fi26
                    var v29 float32 = v6
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            case 42:
                fmt.Println(v14)
            }
            v1[(v3 >> ((4294967295) & 63))] = (true && true)
            fmt.Println(v5)
            _ = fi26
        }
        _ = v25
    } else {
        for fi30 := 0; fi30 < 8; fi30++ {
            var v31 bool = (-128 < 48)
            v5 = 288
            _ = fi30
            _ = v31
        }
    }
    if v8 {
        v32 := GenS3{F0: (v14 / (((v14 * v14)) | 1)), F1: 42, F2: -968}
        v8 = v8
        if ((!("a" < "hello")) && true) {
            v6 -= -2.25
            v33 := v1[uint32(v2)]
            _ = v33
        } else {
            switch "foo bar" {
            case "a":
                switch (v14 / ((0) | 1)) {
                case -305:
                    if v8 {
                        switch (v14 * (^(v14))) {
                        case -9223372036854775808:
                            switch (v14 + v14) {
                            case -664:
                                switch 681 {
                                case -2147483648:
                                    v4 = append(v4, (("" + "") + "hello"))
                                default:
                                    if (v6 != (float32(-2.25) / -2.25)) {
                                        v34 := GenS3{F0: v14, F1: 42, F2: -251}
                                        _ = v34
                                    }
                                }
                            case 1:
                                v35 := GenS4{F0: 0, F1: -2147483648, F2: 65535}
                                _ = v35
                            }
                        case 2147483647:
                            var v36 int32 = -1
                            _ = v36
                        case 128:
                            if (1 <= 65535) {
                                for fi37 := 0; fi37 < 15; fi37++ {
                                    var v38 uint16 = 494
                                    if true {
                                        if (fi37 != fi37) {
                                            switch 4294967295 {
                                            case 841:
                                                var v39 bool = (!v8)
                                                _ = v39
                                            case 4294967295:
                                                for fi40 := 0; fi40 < 5; fi40++ {
                                                    v5 = v3
                                                    _ = fi40
                                                }
                                            default:
                                                v41 := GenS3{F0: 32768, F1: 79, F2: -835}
                                                _ = v41
                                            }
                                            if false {
                                                if (!(32767 <= 0)) {
                                                    for fi42 := 0; fi42 < 20; fi42++ {
                                                        v43 := GenS5{F0: 4294967295, F1: v8}
                                                        v38 = 494
                                                        _ = fi42
                                                        _ = v43
                                                    }
                                                    v2 = uint8(v38)
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                } else {
                                                    for fi44 := 0; fi44 < 10; fi44++ {
                                                        fmt.Println(fi44)
                                                        v45 := GenS6{F0: 1, F1: 2}
                                                        _ = fi44
                                                        _ = v45
                                                    }
                                                    for fi46 := 0; fi46 < 12; fi46++ {
                                                        switch v14 {
                                                        case 1:
                                                            if (4294967295 <= 18446744073709551615) {
                                                                var v47 string = "hello"
                                                                if ((fi37 % ((fi46) | 1)) < (fi46 - -891)) {
                                                                    v5 = v3
                                                                    v5 = 4294967295
                                                                    var v48 float32 = (-2.25 / 1.0)
                                                                    _ = v48
                                                                } else {
                                                                    var v49 string = v47
                                                                    _ = v49
                                                                }
                                                                v14 = fi46
                                                                var v50 uint = 0
                                                                v51 := GenS5{F0: 4294967295, F1: v8}
                                                                _ = v47
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        case -1:
                                                            if (42 <= 42) {
                                                                var v52 int16 = -1
                                                                _ = v52
                                                            }
                                                        }
                                                        _ = fi46
                                                    }
                                                }
                                                v5 = uint32(v38)
                                            }
                                        }
                                    } else {
                                        _ = len(v4)
                                        if len(v4) > 0 {
                                            v4[0] = "foo bar"
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            switch "hello" {
                                            case "":
                                                switch "\n\t" {
                                                case "hello":
                                                    for fi53 := 0; fi53 < 7; fi53++ {
                                                        var v54 uint32 = v5
                                                        _ = fi53
                                                        _ = v54
                                                    }
                                                case "a":
                                                    for fi55 := 0; fi55 < 12; fi55++ {
                                                        fmt.Println(v6, v38)
                                                        var v56 int64 = 1
                                                        _ = fi55
                                                        _ = v56
                                                    }
                                                }
                                            case "a":
                                                v2 -= v2
                                            case "foo bar":
                                                var v57 int8 = -107
                                                _ = v57
                                            }
                                        case "foo bar":
                                            for fi58 := 0; fi58 < 10; fi58++ {
                                                var v59 int64 = -931
                                                v60 := v1[42]
                                                if (42 <= 42) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                } else {
                                                    switch ("hello" + "a") {
                                                    case "foo bar":
                                                        if len(v4) > 0 {
                                                            v4[0] = ""
                                                        }
                                                    case "a":
                                                        switch "hello" {
                                                        case "":
                                                            v6 -= v6
                                                        case "\n\t":
                                                            if len(v4) > 0 {
                                                                v4[0] = "hello"
                                                            }
                                                        }
                                                    case "\n\t":
                                                        if (!v60) {
                                                            if false {
                                                                if (false || v8) {
                                                                    if v60 {
                                                                        var v61 float64 = (1.5 + 1e10)
                                                                        fmt.Println(v61)
                                                                        v62 := GenS6{F0: -32768, F1: uint64(v3)}
                                                                        var v63 bool = v60
                                                                        var v64 int8 = -1
                                                                        switch "" {
                                                                        case "a":
                                                                            v65 := GenS5{F0: 859, F1: (fi58 == (v14 << ((124) & 63)))}
                                                                            _ = v65
                                                                        }
                                                                        delete(v1, 613)
                                                                        v6 -= v6
                                                                        v66 := GenS7{F0: v38, F1: 74}
                                                                        v1[157] = true
                                                                        var v67 bool = v8
                                                                        _ = v61
                                                                        _ = v62
                                                                        _ = v63
                                                                        _ = v64
                                                                        _ = v66
                                                                        _ = v67
                                                                    }
                                                                    switch (fi58 - fi37) {
                                                                    case -9223372036854775808:
                                                                        var v68 uint8 = (v2 + v2)
                                                                        _ = v68
                                                                    default:
                                                                        fmt.Println(v38, fi58)
                                                                    }
                                                                    v14 += 2147483647
                                                                } else {
                                                                    var v69 int8 = 42
                                                                    var v70 int = (fi37 * 32768)
                                                                    v3 -= (v3 ^ v3)
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                                var v71 int16 = int16(v59)
                                                                v72 := GenS6{F0: (v71 / (((^(v71))) | 1)), F1: 0}
                                                                _ = v71
                                                                _ = v72
                                                            }
                                                            var v73 bool = ((v38 << ((0) & 63)) > (v38 & 32767))
                                                            _ = v73
                                                        }
                                                    }
                                                    v59 = v59
                                                    var v74 int32 = 2147483647
                                                    var v75 uint8 = 42
                                                    _ = v74
                                                    _ = v75
                                                }
                                                _ = fi58
                                                _ = v59
                                                _ = v60
                                            }
                                        default:
                                            v38 += 2
                                        }
                                    }
                                    v14 = v14
                                    _ = fi37
                                    _ = v38
                                }
                                var v76 int8 = 72
                                var v77 int64 = 198
                                _ = v76
                                _ = v77
                            }
                        default:
                            v78 := GenS8{F0: 42}
                            _ = v78
                        }
                        v8 = v8
                    } else {
                        for fi79 := 0; fi79 < 13; fi79++ {
                            switch ("a" + ("hello" + "foo bar")) {
                            case "":
                                fmt.Println(v6)
                            case "\n\t":
                                v6 = (1.5 - 0.1)
                            case "a":
                                v4 = append(v4, ("hello" + "foo bar"))
                            default:
                                fmt.Println(v8)
                            }
                            var v80 string = "foo bar"
                            _ = fi79
                            _ = v80
                        }
                    }
                case 32768:
                    v81 := GenS2{F0: (v2 ^ v2), F1: "\n\t", F2: 192}
                    _ = v81
                case -691:
                    v8 = v8
                default:
                    delete(v1, (v5 - 922))
                }
            default:
                v6 = 1e10
            }
            for fi82 := 0; fi82 < 9; fi82++ {
                v14 -= v14
                _ = fi82
            }
            var v83 uint32 = (^(v5))
            _ = v83
        }
        _ = v32
    }
    v6 = 1e10
    v2 = 0
    v3 = 4294967295
    v5 = (v3 << ((255) & 63))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi84 := 0; fi84 < 10; fi84++ {
        if false {
            var v85 float32 = 0.0
            fmt.Println(v3, v8)
            _ = v85
        } else {
            var v86 string = "foo bar"
            fmt.Println(v86, v3, v8)
            var v87 int16 = 368
            _ = v86
            _ = v87
        }
        v88 := GenS9{F0: 0, F1: (v2 ^ v2)}
        v2 -= (v2 % (((v2 / ((v2) | 1))) | 1))
        _ = fi84
        _ = v88
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
