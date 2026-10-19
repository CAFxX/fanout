package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int32 }
type GenS2 struct { F0 int; F1 int16; F2 int64 }
type GenS3 struct { F0 int; F1 string; F2 int16 }
type GenS4 struct { F0 uint8; F1 bool; F2 bool }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 uint; F2 uint }
type GenS8 struct { F0 int; F1 bool }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint16; F1 bool }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 bool; F1 uint64 }

func main() {
    v1 := []uint32{2, 2147483647, 848, 42}
    v2 := []int{453, 319, -664}
    var v3 uint32 = 1
    var v4 uint32 = v3
    fmt.Println(v4, v3)
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 float64 = float64(1.5)
        fmt.Println(v4)
        if (0.1 >= 1e10) {
            v6 = ((v6 * v6) - ((3.14159 - float64(3.14159)) / 1.0))
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 uint8 = uint8(v3)
                _ = fi7
                _ = v8
            }
        } else {
            v9 := GenS1{F0: 758, F1: 156, F2: -440}
            v10 := GenS1{F0: 255, F1: 0, F2: 0}
            if ((v3 / ((847) | 1)) > v3) {
                _ = len(v2)
                var v11 uint8 = 113
                v11 = (v11 % ((127) | 1))
                _ = v11
            } else {
                if false {
                    if (!(((-2.25 * float32(v6)) <= -2.25) && ((0 < (v3 + 2147483648)) && ((v6 * v6) > v6)))) {
                        v12 := GenS2{F0: 0, F1: 42, F2: 42}
                        _ = v12
                    }
                    fmt.Println(v4, v6, v6)
                }
                var v13 uint = 628
                for fi14 := 0; fi14 < 12; fi14++ {
                    v13 = v13
                    v1 = append(v1, v4)
                    _ = fi14
                }
                _ = v13
            }
            _ = v10
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v3 = v4
        _ = fi15
    }
    var v16 float32 = 3.14159
    if true {
        if (!true) {
            v16 = 1e10
            for fi17 := 0; fi17 < 14; fi17++ {
                for fi18 := 0; fi18 < 19; fi18++ {
                    v4 += v4
                    var v19 int64 = 0
                    var v20 float32 = ((float32(1e10) + (0.0 + 0.0)) / (v16 / v16))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                fmt.Println(v16, v16)
                _ = fi17
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                if (1 == 0) {
                    for fi22 := 0; fi22 < 20; fi22++ {
                        if (32767 <= 524) {
                            var v23 int = (fi21 - 32767)
                            _ = v23
                        } else {
                            if true {
                                v4 -= v4
                                v3 -= (v4 * v4)
                                v24 := GenS3{F0: (fi21 % ((fi22) | 1)), F1: "", F2: 1}
                                _ = v24
                            } else {
                                var v25 int32 = -2147483648
                                v25 += (v25 | v25)
                                fmt.Println(v4, v25, v25)
                                _ = v25
                            }
                        }
                        _ = fi22
                    }
                    for fi26 := 0; fi26 < 4; fi26++ {
                        switch (fi21 * -1) {
                        case 1:
                            var v27 uint8 = 231
                            _ = v27
                        case -733:
                            v16 -= v16
                        default:
                            var v28 float32 = float32((-2.25 + (((float64(1e10) + 0.0) * 3.14159) / 0.1)))
                            _ = v28
                        }
                        _ = fi26
                    }
                    if true {
                        var v29 int64 = 1
                        v16 += (3.14159 * v16)
                        _ = v29
                    }
                } else {
                    v3 = (v4 | 65535)
                    var v30 uint = 255
                    fmt.Println(fi21, v4, v16)
                    _ = v30
                }
                switch fi21 {
                case 32767:
                    var v31 int32 = -2147483648
                    _ = v31
                case -9223372036854775808:
                    v16 += v16
                case 42:
                    if true {
                        _ = len(v1)
                        var v32 uint = 1
                        if (!(((730 >= -2147483648) && false) && true)) {
                            v1 = append(v1, (v3 % (((v3 & v4)) | 1)))
                        } else {
                            if (!true) {
                                fmt.Println(v3, v32)
                                v33 := GenS4{F0: 2, F1: false, F2: (v3 != 521)}
                                v16 = v16
                                _ = v33
                            } else {
                                if (false || (!false)) {
                                    v32 = (v32 * v32)
                                    if false {
                                        fmt.Println(v16, v32, fi21)
                                        var v34 int = fi21
                                        switch (^(v34)) {
                                        case 127:
                                            if len(v2) > 0 {
                                                v2[0] = (fi21 * (v34 + fi21))
                                            }
                                        case 444:
                                            if (v16 != 0.1) {
                                                fmt.Println(v32)
                                                fmt.Println(v16, v3)
                                                if (true && true) {
                                                    if len(v1) > 0 {
                                                        v1[0] = v3
                                                    }
                                                }
                                            } else {
                                                v35 := GenS5{F0: 2}
                                                var v36 float32 = v16
                                                if (2 == 42) {
                                                    v1 = append(v1, 1)
                                                } else {
                                                    var v37 uint32 = v3
                                                    v38 := GenS5{F0: 42}
                                                    _ = v37
                                                    _ = v38
                                                }
                                                _ = v35
                                                _ = v36
                                            }
                                        case -2147483648:
                                            fmt.Println(v34, v4)
                                        default:
                                            var v39 uint = (^(v32))
                                            _ = v39
                                        }
                                        _ = v34
                                    }
                                    if false {
                                        switch (("foo bar" + (("foo bar" + "") + "")) + "foo bar") {
                                        case "":
                                            v40 := GenS6{F0: -128}
                                            _ = v40
                                        case "hello":
                                            var v41 uint = v32
                                            _ = v41
                                        default:
                                            _ = len(v1)
                                        }
                                        var v42 uint16 = 0
                                        var v43 uint64 = 2
                                        _ = v42
                                        _ = v43
                                    }
                                }
                                if true {
                                    fmt.Println(v16, fi21)
                                } else {
                                    fmt.Println(fi21, v4)
                                }
                                var v44 float64 = 3.14159
                                _ = v44
                            }
                            switch (fi21 % ((fi21) | 1)) {
                            case -1:
                                switch ("a" + "hello") {
                                case "a":
                                    if true {
                                        switch (("\n\t" + "") + "foo bar") {
                                        case "hello":
                                            v3 = 0
                                        case "a":
                                            var v45 uint64 = 18446744073709551615
                                            _ = v45
                                        case "\n\t":
                                            for fi46 := 0; fi46 < 15; fi46++ {
                                                fmt.Println(fi21, fi46, v16)
                                                if true {
                                                    v16 += v16
                                                    var v47 int8 = 100
                                                    if true {
                                                        v4 += (^(v3))
                                                        var v48 int8 = v47
                                                        var v49 bool = (v32 < v32)
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        switch "a" {
                                                        case "a":
                                                            if len(v1) > 0 {
                                                                v1[0] = v3
                                                            }
                                                        case "hello":
                                                            if len(v2) > 0 {
                                                                v2[0] = (fi46 - 9223372036854775807)
                                                            }
                                                        default:
                                                            var v50 uint64 = 18446744073709551615
                                                            _ = v50
                                                        }
                                                        v51 := GenS7{F0: -116, F1: v32, F2: v32}
                                                        _ = v51
                                                    }
                                                    _ = v47
                                                }
                                                _ = fi46
                                            }
                                        default:
                                            v52 := GenS1{F0: 1, F1: v32, F2: -1}
                                            _ = v52
                                        }
                                        if (-2.25 <= (1.5 / 1.0)) {
                                            if (false && (137 <= 42)) {
                                                var v53 uint8 = 2
                                                v54 := GenS6{F0: int8(v53)}
                                                for fi55 := 0; fi55 < 11; fi55++ {
                                                    v56 := GenS8{F0: fi55, F1: false}
                                                    var v57 int32 = -1
                                                    _ = fi55
                                                    _ = v56
                                                    _ = v57
                                                }
                                                _ = v53
                                                _ = v54
                                            } else {
                                                v16 = ((v16 / ((float32((3.14159 * 0.1)) - v16) * (((float32(0.1) * (1.5 / 1.0)) / (1.5 + v16)) + float32((1.5 - 1e10))))) - 0.0)
                                                for fi58 := 0; fi58 < 12; fi58++ {
                                                    v4 = (v4 % ((v3) | 1))
                                                    v59 := GenS2{F0: fi21, F1: -1, F2: 141}
                                                    switch fi21 {
                                                    case 2147483647:
                                                        if len(v2) > 0 {
                                                            v2[0] = (fi58 & fi21)
                                                        }
                                                        if true {
                                                            v1 = append(v1, v4)
                                                        }
                                                        v60 := GenS9{F0: 127}
                                                        _ = v60
                                                    case 4294967295:
                                                        _ = len(v1)
                                                    case 1:
                                                        v61 := GenS2{F0: fi21, F1: -550, F2: 1}
                                                        _ = v61
                                                    }
                                                    v62 := GenS9{F0: 1}
                                                    v3 -= 4294967295
                                                    _ = fi58
                                                    _ = v59
                                                    _ = v62
                                                }
                                            }
                                            v3 -= (v4 | (v3 | 2147483648))
                                        }
                                        var v63 uint32 = v3
                                        _ = v63
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            for fi64 := 0; fi64 < 18; fi64++ {
                                                var v65 uint32 = 909
                                                var v66 float32 = 0.1
                                                _ = fi64
                                                _ = v65
                                                _ = v66
                                            }
                                        case "":
                                            for fi67 := 0; fi67 < 3; fi67++ {
                                                v3 += v4
                                                _ = fi67
                                            }
                                        case "\n\t":
                                            v68 := GenS4{F0: uint8(v4), F1: ((!false) && (100 <= 127)), F2: (!true)}
                                            _ = v68
                                        default:
                                            v3 = (v4 % (((v3 ^ v3)) | 1))
                                        }
                                    }
                                default:
                                    if (!((517 < -2147483648) || false)) {
                                        v69 := GenS10{F0: uint16(v32), F1: true}
                                        _ = v69
                                    } else {
                                        var v70 int64 = -1
                                        var v71 uint = v32
                                        _ = v70
                                        _ = v71
                                    }
                                }
                            }
                            v32 = v32
                        }
                        _ = v32
                    }
                }
                var v72 string = ""
                _ = fi21
                _ = v72
            }
        } else {
            switch 606 {
            case 172:
                v73 := GenS11{F0: 2}
                _ = v73
            case 4294967295:
                var v74 int = 32767
                _ = v74
            case -575:
                fmt.Println(v3, v16, v16)
            default:
                switch "foo bar" {
                case "a":
                    for fi75 := 0; fi75 < 14; fi75++ {
                        for fi76 := 0; fi76 < 4; fi76++ {
                            v4 = v3
                            _ = fi76
                        }
                        _ = fi75
                    }
                case "foo bar":
                    v4 = (^(v3))
                default:
                    switch "hello" {
                    case "foo bar":
                        fmt.Println(v3, v4)
                    case "hello":
                        _ = len(v1)
                    }
                }
            }
        }
        v77 := GenS4{F0: 128, F1: true, F2: false}
        v4 += v3
        _ = v77
    } else {
        v3 = (v3 - (v4 | v3))
    }
    if true {
        for fi78 := 0; fi78 < 10; fi78++ {
            var v79 int16 = -847
            v79 = v79
            if false {
                var v80 bool = true
                v16 -= v16
                v81 := GenS6{F0: int8(v79)}
                _ = v80
                _ = v81
            } else {
                v3 -= 2147483648
                _ = len(v2)
            }
            _ = fi78
            _ = v79
        }
    } else {
        var v82 uint = 65535
        _ = v82
    }
    var v83 int64 = -608
    switch 1 {
    case -975:
        if true {
            for fi84 := 0; fi84 < 3; fi84++ {
                var v85 int8 = -128
                _ = len(v1)
                var v86 string = "foo bar"
                _ = fi84
                _ = v85
                _ = v86
            }
        } else {
            v4 += (v4 << ((4294967295) & 63))
            var v87 int16 = -371
            _ = v87
        }
    }
    switch "" {
    case "foo bar":
        v88 := GenS8{F0: 1, F1: (1 == 136)}
        _ = v88
    case "":
        v89 := GenS6{F0: -1}
        _ = v89
    default:
        fmt.Println(v4, v83, v4)
    }
    var v90 uint16 = 65535
    if false {
        fmt.Println(v4)
    } else {
        v91 := GenS12{F0: false, F1: 2}
        if (!((!((-(v83)) > (v83 - v83))) || (1.5 < (((3.14159 / (v16 / v16)) / -2.25) / -2.25)))) {
            v92 := GenS7{F0: v83, F1: 255, F2: 51}
            _ = v92
        }
        _ = v91
    }
    _ = len(v1)
    var v93 int16 = -32768
    v90 += (v90 << ((42) & 63))
    v3 = v3
    var v94 float64 = 0.0
    var v95 string = ""
    var v96 int16 = (v93 * v93)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    fmt.Println("v90:", v90)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v90
    _ = v93
    _ = v94
    _ = v95
    _ = v96
}
