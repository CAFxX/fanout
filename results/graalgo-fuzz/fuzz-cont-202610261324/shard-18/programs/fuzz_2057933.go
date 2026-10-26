package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 uint8; F1 int32; F2 uint8 }

func main() {
    v1 := make(map[string]int)
    v2 := []int{-1}
    var v3 uint64 = 1
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    var v6 bool = ((v3 << ((255) & 63)) < v3)
    var v7 float32 = float32(((((((0.1 + 3.14159) * 0.1) / 1.0) + (1.5 / ((1.5 / 1.0) / 0.1))) - 1.5) - (((-2.25 + 1e10) / 1e10) - (0.0 + 0.0))))
    v7 = float32(0.0)
    v8 := GenS1{F0: 1, F1: 114}
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(fi9, v3)
        v7 = v7
        v10 := GenS2{F0: 127, F1: -930}
        _ = fi9
        _ = v10
    }
    v3 += v3
    _ = len(v4)
    var v11 uint8 = 2
    if v6 {
        var v12 bool = (65535 != 2)
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 int16 = 464
            _ = fi13
            _ = v14
        }
        _ = v12
    } else {
        if (-33 <= 0) {
            var v15 string = "foo bar"
            if ((!v6) || (868 <= 127)) {
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 uint16 = 1
                    var v18 uint8 = v11
                    switch (fi16 ^ (fi16 | fi16)) {
                    case -110:
                        for fi19 := 0; fi19 < 2; fi19++ {
                            v5[2] = fi19
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            v11 = v11
                            _ = fi19
                        }
                    case 0:
                        var v20 uint64 = 4294967295
                        _ = v20
                    default:
                        fmt.Println(v17, v6, v15)
                    }
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if v6 {
                    v15 = ""
                } else {
                    var v21 float64 = 1e10
                    v1["foo bar"] = 32768
                    _ = v21
                }
            }
            v22 := GenS3{F0: v6, F1: 42}
            _ = v15
            _ = v22
        }
        var v23 bool = v6
        switch -787 {
        case 4294967295:
            if v6 {
                v24 := GenS2{F0: -1, F1: -1}
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 int64 = -439
                    v23 = v6
                    switch (fi25 << ((65535) & 63)) {
                    case -2147483648:
                        for fi27 := 0; fi27 < 18; fi27++ {
                            v3 += v3
                            _ = fi27
                        }
                    case 128:
                        v11 += 2
                    case 42:
                        var v28 bool = (693 < 42)
                        _ = v28
                    }
                    _ = fi25
                    _ = v26
                }
                _ = v24
            } else {
                switch "hello" {
                case "hello":
                    var v29 bool = v23
                    _ = v29
                case "":
                    var v30 string = ("foo bar" + "\n\t")
                    _ = v30
                case "foo bar":
                    var v31 string = "a"
                    _ = v31
                }
                for fi32 := 0; fi32 < 11; fi32++ {
                    if (!true) {
                        var v33 bool = ((^(v11)) == (v11 >> ((391) & 63)))
                        var v34 int64 = -9223372036854775808
                        _ = v33
                        _ = v34
                    }
                    _ = fi32
                }
                var v35 uint8 = 128
                _ = v35
            }
        default:
            var v36 float64 = (3.14159 + ((0.0 + ((-2.25 * 3.14159) * 0.0)) + float64(v7)))
            _ = v36
        }
        _ = v23
    }
    var v37 int64 = 0
    for fi38 := 0; fi38 < 4; fi38++ {
        for fi39 := 0; fi39 < 5; fi39++ {
            v3 = v3
            _ = len(v5)
            var v40 uint8 = v11
            _ = fi39
            _ = v40
        }
        var v41 bool = v6
        switch (fi38 * fi38) {
        case 32767:
            _ = len(v5)
        case 2147483647:
            if ("a" >= "a") {
                if false {
                    var v42 uint32 = 2147483647
                    _ = v42
                } else {
                    v37 = (v37 ^ v37)
                }
                _ = len(v1)
                var v43 uint = 913
                _ = v43
            }
        default:
            if v6 {
                if false {
                    for fi44 := 0; fi44 < 19; fi44++ {
                        if v6 {
                            var v45 uint32 = 0
                            _ = v45
                        }
                        switch "a" {
                        case "":
                            v41 = v41
                        case "hello":
                            fmt.Println(v37, v37, fi44)
                        case "\n\t":
                            fmt.Println(v37)
                        }
                        _ = len(v5)
                        var v46 int8 = 1
                        switch fi44 {
                        case 718:
                            if false {
                                v37 = v37
                                delete(v4, "foo bar")
                                delete(v4, "\n\t")
                            } else {
                                v37 = (v37 + v37)
                            }
                        }
                        delete(v5, 162)
                        _ = fi44
                        _ = v46
                    }
                    switch -709 {
                    case 127:
                        var v47 float32 = 3.14159
                        _ = v47
                    case -863:
                        if false {
                            v6 = false
                        }
                    default:
                        for fi48 := 0; fi48 < 6; fi48++ {
                            var v49 bool = v41
                            delete(v1, "a")
                            _ = fi48
                            _ = v49
                        }
                    }
                } else {
                    v50 := GenS4{F0: (^(v11)), F1: false}
                    _ = v50
                }
                v6 = (v6 || (true || ("hello" > "foo bar")))
                if (1 >= 0) {
                    for fi51 := 0; fi51 < 13; fi51++ {
                        v41 = true
                        _ = fi51
                    }
                    switch ("foo bar" + ("a" + "")) {
                    case "":
                        fmt.Println(v11, v3)
                    }
                }
            } else {
                v52 := GenS5{F0: true, F1: 1}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = v52
            }
        }
        _ = fi38
        _ = v41
    }
    var v53 int64 = (v37 + v37)
    v54 := GenS6{F0: v6, F1: 1}
    for fi55 := 0; fi55 < 2; fi55++ {
        if (v11 != v11) {
            v56 := v5[1]
            _ = v56
        } else {
            if v6 {
                v57 := GenS5{F0: ((!true) || false), F1: 42}
                _ = v57
            }
            switch ("" + ("foo bar" + "a")) {
            case "a":
                var v58 uint16 = 1
                _ = v58
            case "":
                var v59 uint8 = v11
                _ = v59
            case "\n\t":
                for fi60 := 0; fi60 < 16; fi60++ {
                    _ = len(v5)
                    if (float32(0.0) >= (v7 + float32(-2.25))) {
                        switch (("\n\t" + "a") + "foo bar") {
                        case "foo bar":
                            if len(v2) > 0 {
                                v2[0] = (fi55 << ((uint(v11)) & 63))
                            }
                        }
                        v53 += 42
                    }
                    v61 := GenS7{F0: 1, F1: 65535}
                    _ = fi60
                    _ = v61
                }
            }
            for fi62 := 0; fi62 < 14; fi62++ {
                fmt.Println(fi62)
                v63 := v1[("foo bar" + "hello")]
                _ = fi62
                _ = v63
            }
        }
        if (v6 && (!(v37 >= v37))) {
            fmt.Println(v3, v6, v53)
            _ = len(v2)
        } else {
            fmt.Println(v53, v6, v7)
            var v64 uint32 = 2
            for fi65 := 0; fi65 < 8; fi65++ {
                switch 1 {
                case 4294967295:
                    switch "\n\t" {
                    case "\n\t":
                        if (!((v53 % ((v37) | 1)) <= (v37 + v53))) {
                            if true {
                                v66 := GenS3{F0: v6, F1: v53}
                                _ = v66
                            }
                            fmt.Println(v6)
                        } else {
                            v2 = append(v2, (fi65 % ((2147483647) | 1)))
                            v64 += v64
                        }
                        if false {
                            v11 = (v11 + v11)
                            delete(v1, "foo bar")
                            var v67 string = "foo bar"
                            _ = v67
                        } else {
                            if (713 < (v64 | 323)) {
                                var v68 int = -9223372036854775808
                                _ = v68
                            } else {
                                fmt.Println(v6, v6)
                                var v69 int16 = int16(v11)
                                _ = v69
                            }
                            var v70 float32 = 0.1
                            _ = v70
                        }
                        v2 = append(v2, 1)
                    case "":
                        switch "a" {
                        case "\n\t":
                            if (0.0 <= float64((v7 * float32((0.1 * ((0.1 - (0.0 * (0.0 * 1e10))) - 0.1)))))) {
                                var v71 int = 4294967295
                                v3 -= (v3 & uint64(v11))
                                v53 = (^(v53))
                                _ = v71
                            } else {
                                v2 = append(v2, fi55)
                                var v72 string = "\n\t"
                                v73 := v5[(v64 - v64)]
                                v7 -= v7
                                var v74 float64 = (-2.25 - 0.0)
                                _ = v72
                                _ = v73
                                _ = v74
                            }
                            switch (-(fi65)) {
                            case 128:
                                if (!v6) {
                                    v37 = v37
                                }
                            }
                        case "hello":
                            v11 += 2
                            var v75 uint64 = v3
                            _ = v75
                        case "foo bar":
                            if ((^(v11)) == v11) {
                                v76 := GenS5{F0: v6, F1: -398}
                                if (v64 == (^(v64))) {
                                    v77 := GenS4{F0: (^(v11)), F1: (0 <= 715)}
                                    v11 = (v11 ^ 44)
                                    var v78 bool = (v64 <= 42)
                                    _ = v77
                                    _ = v78
                                }
                                fmt.Println(v11, v7, v7)
                                v11 = (v11 << ((4294967295) & 63))
                                _ = v76
                            } else {
                                if (v6 && true) {
                                    var v79 uint8 = (v11 * uint8(v64))
                                    _ = v79
                                }
                                v7 = float32(3.14159)
                                delete(v5, 2147483648)
                            }
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    case "a":
                        var v80 bool = (1 < 1)
                        _ = v80
                    default:
                        v3 -= (v3 | 10)
                    }
                default:
                    var v81 bool = v6
                    var v82 int64 = v37
                    _ = v81
                    _ = v82
                }
                v83 := GenS4{F0: v11, F1: ((!(0 <= uint16(v3))) && true)}
                var v84 int = (fi65 & 2147483647)
                _ = fi65
                _ = v83
                _ = v84
            }
            _ = v64
        }
        if v6 {
            if v6 {
                switch fi55 {
                case -1:
                    v53 = v53
                default:
                    v85 := GenS8{F0: (^(v11)), F1: -2147483648, F2: (v11 >> ((uint(v11)) & 63))}
                    _ = v85
                }
                switch "hello" {
                case "":
                    var v86 uint8 = (v11 ^ (v11 ^ 2))
                    _ = v86
                }
                for fi87 := 0; fi87 < 2; fi87++ {
                    var v88 int32 = 0
                    _ = fi87
                    _ = v88
                }
            }
            var v89 uint64 = v3
            _ = v89
        }
        _ = fi55
    }
    for fi90 := 0; fi90 < 1; fi90++ {
        fmt.Println(v37, v6)
        _ = fi90
    }
    var v91 bool = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
    _ = v7
    _ = v8
    _ = v91
}
