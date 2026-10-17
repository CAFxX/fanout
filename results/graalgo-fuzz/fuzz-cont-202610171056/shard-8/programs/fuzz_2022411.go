package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 int64; F1 int32; F2 uint16 }
type GenS10 struct { F0 uint64; F1 string; F2 string }

func main() {
    var v1 int32 = 1
    v2 := []string{"foo bar", "\n\t"}
    v3 := []int{12}
    var v4 int = 1
    var v5 int16 = 0
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v5, fi7, fi7)
        v1 = (v1 | 0)
        v1 = v1
        _ = fi7
    }
    _ = len(v2)
    var v8 int = v4
    v5 -= -32768
    if false {
        fmt.Println(v5)
        v9 := GenS1{F0: 155, F1: v1, F2: ("foo bar" + (("hello" + "a") + ""))}
        if (32767 < 32767) {
            switch (v8 / (((v8 / ((0) | 1))) | 1)) {
            case 128:
                if (0 <= 255) {
                    _ = len(v3)
                    var v10 int8 = 1
                    if (2 >= 97) {
                        if true {
                            v8 -= int(v10)
                            for fi11 := 0; fi11 < 11; fi11++ {
                                fmt.Println(v5)
                                switch "foo bar" {
                                case "":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "\n\t":
                                    if (42 > fi11) {
                                        _ = len(v6)
                                    } else {
                                        v5 -= 1
                                        for fi12 := 0; fi12 < 19; fi12++ {
                                            v3 = append(v3, (fi12 >> ((2) & 63)))
                                            v13 := GenS1{F0: 0, F1: -2147483648, F2: ""}
                                            _ = fi12
                                            _ = v13
                                        }
                                    }
                                case "a":
                                    for fi14 := 0; fi14 < 9; fi14++ {
                                        fmt.Println(v10, v1, v1)
                                        v10 += (v10 % (((v10 / ((v10) | 1))) | 1))
                                        var v15 uint64 = 133
                                        _ = fi14
                                        _ = v15
                                    }
                                }
                                _ = fi11
                            }
                            switch "" {
                            case "\n\t":
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = v8
                                    }
                                    var v16 uint32 = 2
                                    var v17 uint64 = 1
                                    _ = v16
                                    _ = v17
                                }
                            case "hello":
                                switch (^(v8)) {
                                case 32768:
                                    v1 = v1
                                case 2147483647:
                                    for fi18 := 0; fi18 < 17; fi18++ {
                                        v6[(fi18 ^ (v4 >> ((42) & 63)))] = 743
                                        _ = fi18
                                    }
                                }
                            default:
                                v10 -= v10
                            }
                        } else {
                            fmt.Println(v5)
                        }
                        v8 += (-(v4))
                        var v19 uint8 = 1
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 16; fi20++ {
                            if ((v10 & v10) <= 1) {
                                fmt.Println(fi20, fi20)
                            } else {
                                fmt.Println(v1, v8)
                            }
                            v21 := GenS1{F0: 128, F1: -2147483648, F2: ("a" + "")}
                            var v22 float32 = 3.14159
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        if false {
                            if true {
                                for fi23 := 0; fi23 < 12; fi23++ {
                                    fmt.Println(v1, fi23, fi23)
                                    _ = fi23
                                }
                                v5 = v5
                            } else {
                                fmt.Println(v4, v10, v1)
                                if true {
                                    v10 += v10
                                }
                                v8 = 128
                            }
                        }
                        var v24 string = "hello"
                        _ = v24
                    }
                    _ = v10
                }
            case -1:
                v5 += -32768
            case 1:
                var v25 int16 = 0
                _ = v25
            }
            var v26 float32 = 3.14159
            _ = v26
        }
        _ = v9
    } else {
        v8 -= (v8 << ((0) & 63))
        switch "a" {
        case "hello":
            v5 = -1
        case "foo bar":
            v27 := GenS2{F0: 0, F1: 2}
            _ = v27
        default:
            for fi28 := 0; fi28 < 20; fi28++ {
                v1 += v1
                _ = fi28
            }
        }
        var v29 uint16 = 2
        _ = v29
    }
    delete(v6, (v4 >> ((4294967295) & 63)))
    var v30 uint = 65535
    if false {
        switch 0 {
        case 127:
            switch "" {
            case "":
                v4 -= v4
            case "\n\t":
                if true {
                    if (!(4294967295 < 42)) {
                        _ = len(v3)
                        v5 += v5
                        v8 = (^(v8))
                    }
                    v6[int(v5)] = 4294967295
                }
            default:
                v31 := GenS2{F0: -1, F1: 310}
                _ = v31
            }
        case -1:
            v32 := GenS3{F0: 1}
            _ = v32
        case 32767:
            for fi33 := 0; fi33 < 12; fi33++ {
                for fi34 := 0; fi34 < 7; fi34++ {
                    if (127 > 1) {
                        v8 = 1
                    }
                    _ = fi34
                }
                v35 := GenS4{F0: 1}
                var v36 float32 = 1e10
                _ = fi33
                _ = v35
                _ = v36
            }
        }
    }
    var v37 int32 = 380
    if false {
        if (!((!false) && ((v1 < (v37 % (((v37 - v1)) | 1))) && (!true)))) {
            v38 := GenS5{F0: 4294967295, F1: -503}
            var v39 float64 = (((1e10 + float64(((((3.14159 * 0.1) - (0.1 * (-2.25 * ((1.5 / -2.25) * 0.1)))) + 0.1) / 1e10))) / 1.5) + 1e10)
            _ = v38
            _ = v39
        } else {
            if false {
                v4 -= (v4 | (v4 - (v8 * v4)))
                switch ((("hello" + ("" + "")) + (("foo bar" + "hello") + "\n\t")) + "a") {
                case "":
                    switch v8 {
                    case 1:
                        if (!(!false)) {
                            v40 := GenS6{F0: 246}
                            v8 = (-(v8))
                            _ = v40
                        }
                    }
                case "a":
                    if (-783 != (v5 | (v5 | (v5 / (((v5 % (((v5 ^ (^(v5)))) | 1))) | 1))))) {
                        for fi41 := 0; fi41 < 1; fi41++ {
                            for fi42 := 0; fi42 < 14; fi42++ {
                                fmt.Println(v37, fi42)
                                v5 = v5
                                if (!false) {
                                    if true {
                                        v1 = 0
                                        v1 -= v1
                                    }
                                    fmt.Println(fi41, fi41, v1)
                                    v1 += v1
                                }
                                v43 := GenS7{F0: true, F1: 127}
                                _ = fi42
                                _ = v43
                            }
                            v1 += (v37 & 794)
                            _ = fi41
                        }
                        v5 = v5
                        var v44 float64 = (0.1 / (1e10 - 0.0))
                        _ = v44
                    }
                case "\n\t":
                    var v45 float32 = -2.25
                    _ = v45
                }
            } else {
                v4 += (v4 + v4)
                var v46 float32 = 1e10
                _ = v46
            }
            v8 = v4
            if false {
                for fi47 := 0; fi47 < 12; fi47++ {
                    fmt.Println(v1, v8)
                    _ = fi47
                }
                fmt.Println(v30, v4)
                var v48 string = "hello"
                _ = v48
            } else {
                for fi49 := 0; fi49 < 16; fi49++ {
                    v37 += v37
                    _ = fi49
                }
                if (!false) {
                    v5 = v5
                    switch (("hello" + "\n\t") + "foo bar") {
                    case "a":
                        var v50 int16 = v5
                        _ = v50
                    default:
                        v6[(v8 / (((v4 >> (((v30 << ((v30) & 63))) & 63))) | 1))] = 42
                    }
                } else {
                    var v51 int8 = 54
                    switch (v8 + (v4 / ((v4) | 1))) {
                    case -640:
                        v52 := GenS4{F0: -835}
                        _ = v52
                    default:
                        switch (v8 % (((v8 + (v4 % ((v8) | 1)))) | 1)) {
                        case 32767:
                            var v53 uint8 = 0
                            _ = v53
                        case -587:
                            switch "\n\t" {
                            case "":
                                var v54 int = v4
                                _ = v54
                            case "hello":
                                if ((^(v4)) < (v4 << (((v30 ^ v30)) & 63))) {
                                    for fi55 := 0; fi55 < 2; fi55++ {
                                        for fi56 := 0; fi56 < 17; fi56++ {
                                            _ = len(v2)
                                            if (125 >= 0) {
                                                if (!(false || true)) {
                                                    _ = len(v2)
                                                    if (!((v1 << (((v30 >> ((v30) & 63))) & 63)) != v1)) {
                                                        if (2 == v30) {
                                                            var v57 int16 = (v5 / (((v5 / ((-32768) | 1))) | 1))
                                                            _ = v57
                                                        }
                                                        v5 = (v5 + v5)
                                                    } else {
                                                        v58 := GenS8{F0: 255, F1: (v5 / ((-1) | 1))}
                                                        fmt.Println(fi55, v1)
                                                        var v59 uint32 = 881
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                    if true {
                                                        var v60 int8 = -128
                                                        _ = v60
                                                    }
                                                    if (2147483648 < 65535) {
                                                        v61 := GenS2{F0: 42, F1: 2147483647}
                                                        _ = v61
                                                    } else {
                                                        v62 := GenS9{F0: -173, F1: (v37 ^ 1), F2: 255}
                                                        v8 = (fi55 & v4)
                                                        if ((v30 % (((v30 & (v30 >> ((v30) & 63)))) | 1)) <= 255) {
                                                            if (32767 >= 1) {
                                                                v51 += v51
                                                                var v63 uint = 0
                                                                v1 += (^(v1))
                                                                _ = v63
                                                            } else {
                                                                fmt.Println(v8, v37)
                                                            }
                                                            var v64 string = "hello"
                                                            _ = len(v2)
                                                            _ = v64
                                                        } else {
                                                            v4 += (v4 - v4)
                                                        }
                                                        _ = v62
                                                    }
                                                    var v65 bool = (!true)
                                                    _ = v65
                                                } else {
                                                    var v66 float32 = 3.14159
                                                    var v67 uint64 = 18446744073709551615
                                                    _ = v66
                                                    _ = v67
                                                }
                                            }
                                            _ = fi56
                                        }
                                        if ((v30 / ((v30) | 1)) == (v30 * v30)) {
                                            for fi68 := 0; fi68 < 17; fi68++ {
                                                var v69 int8 = v51
                                                _ = fi68
                                                _ = v69
                                            }
                                        }
                                        _ = fi55
                                    }
                                    v70 := GenS6{F0: 128}
                                    _ = v70
                                }
                            case "\n\t":
                                v5 = (v5 ^ (v5 | 1))
                            default:
                                v71 := GenS7{F0: false, F1: 64}
                                _ = v71
                            }
                        case 9223372036854775807:
                            if false {
                                var v72 float32 = -2.25
                                v5 = (v5 >> ((v30) & 63))
                                for fi73 := 0; fi73 < 19; fi73++ {
                                    for fi74 := 0; fi74 < 13; fi74++ {
                                        var v75 int64 = 1
                                        fmt.Println(v72, v30)
                                        if true {
                                            v5 = (v5 << ((v30) & 63))
                                            var v76 float64 = 1e10
                                            _ = v76
                                        }
                                        _ = fi74
                                        _ = v75
                                    }
                                    var v77 float32 = float32(0.0)
                                    for fi78 := 0; fi78 < 20; fi78++ {
                                        var v79 bool = true
                                        var v80 float64 = 3.14159
                                        if (((v51 << ((1) & 63)) > -68) && v79) {
                                            v81 := GenS10{F0: 18446744073709551615, F1: "\n\t", F2: ""}
                                            var v82 bool = (v79 || v79)
                                            var v83 int = fi78
                                            _ = v81
                                            _ = v82
                                            _ = v83
                                        } else {
                                            var v84 int8 = (v51 | -1)
                                            switch (fi73 + fi73) {
                                            case 4294967295:
                                                var v85 float64 = float64(0.0)
                                                _ = v85
                                            case -64:
                                                v8 = fi73
                                            case 32767:
                                                v86 := GenS4{F0: 9223372036854775807}
                                                _ = v86
                                            }
                                            var v87 int64 = -9223372036854775808
                                            fmt.Println(fi73, v1, fi78)
                                            v88 := GenS10{F0: 1, F1: "hello", F2: "\n\t"}
                                            _ = v84
                                            _ = v87
                                            _ = v88
                                        }
                                        _ = fi78
                                        _ = v79
                                        _ = v80
                                    }
                                    _ = fi73
                                    _ = v77
                                }
                                _ = v72
                            }
                        default:
                            var v89 uint32 = 0
                            _ = v89
                        }
                    }
                    var v90 uint32 = 2
                    _ = v51
                    _ = v90
                }
                delete(v6, (v8 + (^(v8))))
            }
        }
        for fi91 := 0; fi91 < 19; fi91++ {
            v37 = (v1 - (v1 % ((v1) | 1)))
            _ = fi91
        }
    }
    v92 := GenS10{F0: 42, F1: (("" + "foo bar") + ("\n\t" + ("foo bar" + ""))), F2: "foo bar"}
    switch "\n\t" {
    case "foo bar":
        var v93 int8 = 1
        _ = v93
    case "\n\t":
        var v94 string = ("\n\t" + ("a" + ""))
        _ = v94
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v92
}
