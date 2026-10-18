package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 string; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 uint16; F1 int8 }

func main() {
    v1 := []uint32{42, 954, 2147483647, 1}
    v2 := []uint32{4294967295, 65535}
    v3 := make(map[string]uint32)
    var v4 float64 = 1.5
    v5 := make(map[int]int)
    v6 := GenS1{F0: -1}
    for fi7 := 0; fi7 < 7; fi7++ {
        if (true || true) {
            for fi8 := 0; fi8 < 9; fi8++ {
                v4 = 1.5
                _ = fi8
            }
            v9 := GenS1{F0: 42}
            _ = v9
        } else {
            var v10 bool = (65535 <= 53)
            _ = v10
        }
        fmt.Println(fi7)
        _ = fi7
    }
    v4 -= v4
    var v11 int32 = -482
    if len(v1) > 0 {
        v1[0] = 694
    }
    if ((true && false) && false) {
        for fi12 := 0; fi12 < 17; fi12++ {
            v11 += 1
            if ((-255 != -612) && (false || (446 < -32768))) {
                switch ("\n\t" + "\n\t") {
                case "hello":
                    v3[("\n\t" + "hello")] = 730
                case "a":
                    if (139 >= 2) {
                        if (65535 >= 42) {
                            v4 += v4
                            fmt.Println(v11)
                            if ("hello" > (("\n\t" + "hello") + "")) {
                                if false {
                                    fmt.Println(fi12)
                                }
                                switch "foo bar" {
                                case "\n\t":
                                    v13 := GenS1{F0: -245}
                                    _ = v13
                                case "":
                                    switch 9223372036854775807 {
                                    case -860:
                                        v14 := GenS1{F0: int16(v11)}
                                        _ = v14
                                    }
                                default:
                                    if ((!(0.0 <= 0.1)) && true) {
                                        for fi15 := 0; fi15 < 9; fi15++ {
                                            v11 = v11
                                            _ = fi15
                                        }
                                        v16 := GenS2{F0: 680, F1: ("" + "")}
                                        _ = v16
                                    } else {
                                        fmt.Println(fi12, v4, fi12)
                                        fmt.Println(fi12, fi12)
                                    }
                                }
                            }
                        } else {
                            for fi17 := 0; fi17 < 19; fi17++ {
                                if (!(2 != 0)) {
                                    switch (fi12 + (fi12 & 443)) {
                                    case -1:
                                        var v18 bool = (false && (2147483647 == 0))
                                        _ = v18
                                    case 2147483647:
                                        v11 += (v11 | v11)
                                        fmt.Println(fi17, v11, v11)
                                        var v19 float64 = v4
                                        _ = v19
                                    }
                                    v2 = append(v2, 796)
                                }
                                var v20 uint32 = 2
                                v11 = v11
                                _ = fi17
                                _ = v20
                            }
                            v11 -= v11
                        }
                    } else {
                        if false {
                            var v21 uint16 = 0
                            var v22 string = (("a" + (("" + (("a" + ("\n\t" + "a")) + "a")) + "\n\t")) + "a")
                            fmt.Println(v11, v4, v4)
                            _ = v21
                            _ = v22
                        }
                        delete(v5, 127)
                        if (true || (0 != 127)) {
                            v23 := GenS3{F0: "foo bar", F1: 18446744073709551615, F2: 18446744073709551615}
                            var v24 float64 = v4
                            if false {
                                for fi25 := 0; fi25 < 19; fi25++ {
                                    var v26 uint64 = 42
                                    var v27 int16 = int16(fi25)
                                    _ = len(v5)
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                                for fi28 := 0; fi28 < 14; fi28++ {
                                    fmt.Println(v24, v4)
                                    var v29 int = (fi28 * (fi12 >> ((874) & 63)))
                                    v11 = v11
                                    _ = fi28
                                    _ = v29
                                }
                                switch (fi12 >> ((531) & 63)) {
                                case 42:
                                    v30 := GenS4{F0: 401, F1: ("\n\t" + "hello")}
                                    _ = v30
                                case 909:
                                    switch fi12 {
                                    case 73:
                                        if (703 != 34) {
                                            if (684 < 4294967295) {
                                                var v31 uint8 = 217
                                                v32 := GenS1{F0: 76}
                                                v24 += 0.0
                                                _ = v31
                                                _ = v32
                                            } else {
                                                v5[fi12] = fi12
                                            }
                                        }
                                    case 2147483647:
                                        if ((v11 + -1) == v11) {
                                            switch fi12 {
                                            case 32767:
                                                v33 := GenS2{F0: 2147483647, F1: "foo bar"}
                                                _ = v33
                                            case -739:
                                                _ = len(v5)
                                            case -2147483648:
                                                switch (((((("a" + "") + ("" + "foo bar")) + "hello") + ("a" + "foo bar")) + "\n\t") + ("a" + "a")) {
                                                case "\n\t":
                                                    for fi34 := 0; fi34 < 12; fi34++ {
                                                        v24 -= v24
                                                        _ = fi34
                                                    }
                                                default:
                                                    switch fi12 {
                                                    case 0:
                                                        fmt.Println(v4, v24, fi12)
                                                    case 4294967295:
                                                        var v35 int32 = 2147483647
                                                        _ = v35
                                                    case 32767:
                                                        v11 -= 361
                                                    }
                                                }
                                            }
                                            if true {
                                                var v36 uint8 = 127
                                                v11 -= v11
                                                var v37 uint16 = 42
                                                _ = v36
                                                _ = v37
                                            } else {
                                                if false {
                                                    v38 := GenS1{F0: -1}
                                                    _ = v38
                                                }
                                                if true {
                                                    switch "hello" {
                                                    case "hello":
                                                        for fi39 := 0; fi39 < 6; fi39++ {
                                                            v40 := GenS2{F0: 2, F1: ""}
                                                            v41 := GenS2{F0: 26, F1: ("a" + "foo bar")}
                                                            _ = fi39
                                                            _ = v40
                                                            _ = v41
                                                        }
                                                    }
                                                    v11 = (-(v11))
                                                } else {
                                                    for fi42 := 0; fi42 < 12; fi42++ {
                                                        fmt.Println(v11)
                                                        _ = fi42
                                                    }
                                                }
                                            }
                                            if (65535 != 2) {
                                                for fi43 := 0; fi43 < 16; fi43++ {
                                                    if (4294967295 >= 2147483648) {
                                                        v4 = v4
                                                        switch fi12 {
                                                        case 127:
                                                            v24 += (v4 * (1.5 * (v24 / 1.5)))
                                                        case 9223372036854775807:
                                                            fmt.Println(fi43, v4)
                                                        case 32768:
                                                            var v44 int8 = 100
                                                            _ = v44
                                                        default:
                                                            fmt.Println(fi12, v24, v4)
                                                        }
                                                        var v45 float64 = ((0.0 * (v4 * v4)) - v4)
                                                        _ = v45
                                                    } else {
                                                        fmt.Println(v4, v11)
                                                        v24 = ((v4 + v24) / (v24 * v4))
                                                        v46 := GenS2{F0: 4294967295, F1: "foo bar"}
                                                        _ = v46
                                                    }
                                                    v24 = (v24 + (v24 - 1e10))
                                                    var v47 string = "\n\t"
                                                    _ = fi43
                                                    _ = v47
                                                }
                                                fmt.Println(v11, fi12)
                                            }
                                        } else {
                                            if true {
                                                v4 += v24
                                                if ((65535 >= 2) && (((1e10 * 1.5) <= float64(0.1)) || (127 >= 2))) {
                                                    if true {
                                                        for fi48 := 0; fi48 < 18; fi48++ {
                                                            v49 := GenS4{F0: 32767, F1: "a"}
                                                            _ = fi48
                                                            _ = v49
                                                        }
                                                        fmt.Println(v4, fi12)
                                                    } else {
                                                        for fi50 := 0; fi50 < 20; fi50++ {
                                                            v24 -= 0.0
                                                            if true {
                                                                fmt.Println(fi50, v24, fi12)
                                                                fmt.Println(v24, v24, fi12)
                                                            }
                                                            _ = len(v3)
                                                            _ = fi50
                                                        }
                                                    }
                                                } else {
                                                    var v51 float64 = 3.14159
                                                    var v52 uint8 = 2
                                                    _ = v51
                                                    _ = v52
                                                }
                                                v53 := GenS5{F0: true, F1: 1}
                                                _ = v53
                                            }
                                            v54 := v5[fi12]
                                            for fi55 := 0; fi55 < 15; fi55++ {
                                                var v56 float64 = -2.25
                                                _ = fi55
                                                _ = v56
                                            }
                                            _ = v54
                                        }
                                    }
                                case 715:
                                    v11 = (v11 * (v11 / ((0) | 1)))
                                }
                            }
                            _ = v23
                            _ = v24
                        } else {
                            v11 += v11
                            for fi57 := 0; fi57 < 20; fi57++ {
                                if ((255 <= 912) || (((v4 / 1.0) / v4) > 0.1)) {
                                    var v58 uint32 = 1
                                    v11 += (v11 ^ -2147483648)
                                    if (v58 == v58) {
                                        v59 := GenS6{F0: false, F1: -45}
                                        v4 -= (v4 * (float64(0.0) - (((((v4 + -2.25) - v4) + 0.1) + 1e10) * 1e10)))
                                        _ = v59
                                    } else {
                                        switch "hello" {
                                        case "hello":
                                            if ((("a" + "\n\t") + "") < "") {
                                                v60 := GenS4{F0: 2, F1: ("" + ("foo bar" + "a"))}
                                                if (((!(false || (!false))) || (true && (true && (127 > 127)))) && (1 <= 733)) {
                                                    var v61 uint64 = 1
                                                    if (!true) {
                                                        v62 := GenS4{F0: 32767, F1: "foo bar"}
                                                        v63 := GenS7{F0: 1, F1: ""}
                                                        v64 := GenS7{F0: 127, F1: "foo bar"}
                                                        _ = v62
                                                        _ = v63
                                                        _ = v64
                                                    } else {
                                                        v4 = 1e10
                                                        if len(v1) > 0 {
                                                            v1[0] = (v58 >> ((1) & 63))
                                                        }
                                                        if (732 > -9223372036854775808) {
                                                            var v65 uint32 = 1
                                                            _ = v65
                                                        }
                                                        v11 = 884
                                                        _ = len(v2)
                                                    }
                                                    if (1 > 252) {
                                                        if false {
                                                            var v66 int32 = (v11 ^ v11)
                                                            var v67 float64 = v4
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                    } else {
                                                        v68 := GenS1{F0: -306}
                                                        var v69 bool = true
                                                        _ = v68
                                                        _ = v69
                                                    }
                                                    _ = v61
                                                }
                                                var v70 uint32 = (v58 ^ 666)
                                                _ = v60
                                                _ = v70
                                            } else {
                                                fmt.Println(v4, v11)
                                                fmt.Println(v4)
                                            }
                                        case "a":
                                            fmt.Println(v4)
                                            fmt.Println(v58)
                                        case "":
                                            v71 := GenS2{F0: v58, F1: ("\n\t" + "a")}
                                            _ = v71
                                        default:
                                            switch "a" {
                                            case "foo bar":
                                                switch 32767 {
                                                case 32767:
                                                    fmt.Println(v11)
                                                case -1:
                                                    v11 -= (v11 >> ((98) & 63))
                                                }
                                            case "hello":
                                                var v72 float32 = 3.14159
                                                _ = v72
                                            case "\n\t":
                                                var v73 int16 = 32767
                                                _ = v73
                                            default:
                                                switch fi57 {
                                                case 829:
                                                    v58 = v58
                                                case 1:
                                                    var v74 uint32 = 986
                                                    _ = v74
                                                case 32768:
                                                    if (997 < 9223372036854775807) {
                                                        if ((v58 & 65535) >= v58) {
                                                            var v75 string = ("\n\t" + ("hello" + "a"))
                                                            fmt.Println(v11, fi57)
                                                            _ = v75
                                                        } else {
                                                            v76 := GenS7{F0: 143, F1: ("\n\t" + "a")}
                                                            _ = v76
                                                        }
                                                        var v77 float32 = 0.1
                                                        _ = v77
                                                    }
                                                    switch -392 {
                                                    case 413:
                                                        switch fi57 {
                                                        case -953:
                                                            var v78 uint = 18446744073709551615
                                                            _ = v78
                                                        case -9223372036854775808:
                                                            v79 := v3[""]
                                                            _ = v79
                                                        default:
                                                            v80 := v3["hello"]
                                                            _ = v80
                                                        }
                                                    case 4294967295:
                                                        var v81 uint32 = v58
                                                        _ = v81
                                                    }
                                                    v58 = 863
                                                default:
                                                    fmt.Println(v58, v11)
                                                }
                                                var v82 uint64 = 18446744073709551615
                                                _ = v82
                                            }
                                        }
                                        var v83 uint = 0
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v83
                                    }
                                    _ = v58
                                }
                                _ = fi57
                            }
                            v84 := GenS8{F0: 65535, F1: -128}
                            _ = v84
                        }
                    }
                case "":
                    v85 := GenS4{F0: 255, F1: "foo bar"}
                    _ = v85
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v86 int8 = 1
                v11 -= 0
                _ = v86
            }
            _ = fi12
        }
    }
    v11 = v11
    v11 -= (^(v11))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
