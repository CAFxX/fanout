package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int; F2 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16; F1 uint32; F2 int64 }
type GenS9 struct { F0 int32; F1 int8 }
type GenS10 struct { F0 int32; F1 int64 }
type GenS11 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS12 struct { F0 uint; F1 int8; F2 int32 }

func main() {
    var v1 uint = 4294967295
    v2 := make(map[int]bool)
    v3 := []string{}
    var v4 uint32 = 42
    var v5 float64 = -2.25
    var v6 int64 = -1
    switch 32767 {
    case 128:
        if (11 != 111) {
            switch "a" {
            case "":
                if (127 != -46) {
                    if (-9223372036854775808 == -2147483648) {
                        switch (("foo bar" + ("\n\t" + "hello")) + "foo bar") {
                        case "foo bar":
                            var v7 int64 = 42
                            _ = v7
                        default:
                            for fi8 := 0; fi8 < 12; fi8++ {
                                switch fi8 {
                                case -964:
                                    fmt.Println(v6)
                                case 128:
                                    fmt.Println(v1, fi8)
                                case 42:
                                    for fi9 := 0; fi9 < 4; fi9++ {
                                        var v10 int = (fi9 - fi8)
                                        _ = fi9
                                        _ = v10
                                    }
                                }
                                if false {
                                    switch (-(fi8)) {
                                    case 1:
                                        var v11 uint16 = 255
                                        _ = v11
                                    case -67:
                                        var v12 bool = (false && true)
                                        _ = v12
                                    }
                                }
                                for fi13 := 0; fi13 < 8; fi13++ {
                                    v14 := GenS1{F0: 0, F1: (fi13 >> (((v1 & 2)) & 63)), F2: false}
                                    var v15 uint16 = uint16(v4)
                                    v16 := GenS1{F0: 0, F1: (fi8 + (fi13 >> (((v1 - v1)) & 63))), F2: ((v1 * v1) < v1)}
                                    switch -1 {
                                    case 32768:
                                        fmt.Println(v15)
                                    case 1:
                                        v15 = (v15 >> ((775) & 63))
                                    case 42:
                                        switch "foo bar" {
                                        case "\n\t":
                                            var v17 uint8 = 1
                                            _ = v17
                                        case "":
                                            var v18 int = fi13
                                            _ = v18
                                        case "hello":
                                            v19 := GenS1{F0: 255, F1: (fi13 % (((fi8 >> ((v1) & 63))) | 1)), F2: (((0.1 + 3.14159) + 0.1) > -2.25)}
                                            _ = v19
                                        default:
                                            v5 += ((v5 / ((v5 + float64(float32(v5))) * v5)) / (0.1 - (((v5 + 1e10) * float64(1.5)) + v5)))
                                        }
                                    }
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                                _ = fi8
                            }
                        }
                        v1 = v1
                    }
                    v20 := GenS2{F0: -378}
                    _ = v20
                }
            }
            if (-60 == 1) {
                v21 := GenS3{F0: 823, F1: 4294967295, F2: 308}
                if false {
                    v22 := GenS1{F0: 14, F1: 32767, F2: ("\n\t" >= ("" + "\n\t"))}
                    v23 := GenS4{F0: 1}
                    v4 -= v4
                    _ = v22
                    _ = v23
                }
                v1 += (^(v1))
                _ = v21
            }
        }
    default:
        fmt.Println(v4)
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        fmt.Println(v4, v5)
        for fi25 := 0; fi25 < 20; fi25++ {
            var v26 bool = false
            _ = fi25
            _ = v26
        }
        fmt.Println(v6)
        _ = fi24
    }
    switch 1 {
    case 246:
        v27 := GenS1{F0: 42, F1: 127, F2: false}
        _ = v27
    case 267:
        v6 = -1
    case 1:
        v1 = (v1 | v1)
    default:
        v28 := GenS5{F0: v6, F1: 1}
        _ = v28
    }
    switch "" {
    case "hello":
        var v29 int = -202
        _ = v29
    case "foo bar":
        fmt.Println(v1, v6, v6)
    }
    fmt.Println(v6)
    if false {
        var v30 int = -1
        v30 = (-(v30))
        _ = v30
    } else {
        v1 = 18446744073709551615
        switch int(v1) {
        case -234:
            switch "a" {
            case "foo bar":
                if ((v1 % (((^(v1))) | 1)) <= v1) {
                    if true {
                        switch 32768 {
                        case 42:
                            v6 -= v6
                        case 9223372036854775807:
                            delete(v2, 127)
                        case -9223372036854775808:
                            var v31 int64 = v6
                            _ = v31
                        default:
                            var v32 int8 = 100
                            _ = v32
                        }
                        v33 := GenS6{F0: ("a" + "a")}
                        _ = v33
                    } else {
                        if (!(!false)) {
                            var v34 int16 = 698
                            _ = len(v3)
                            _ = v34
                        } else {
                            v35 := GenS7{F0: 0}
                            _ = v35
                        }
                    }
                }
            case "\n\t":
                v6 = (v6 % ((312) | 1))
            case "hello":
                v1 += (v1 * v1)
            }
        default:
            v1 += (^(v1))
        }
    }
    v2[-2147483648] = true
    v36 := GenS8{F0: -1, F1: 0, F2: (v6 & (v6 >> ((1) & 63)))}
    var v37 int = 1
    var v38 uint32 = 2
    switch "a" {
    case "\n\t":
        if false {
            var v39 uint64 = 688
            v40 := GenS9{F0: 1, F1: 97}
            var v41 uint32 = (v4 * v38)
            _ = v39
            _ = v40
            _ = v41
        } else {
            v37 -= (-(v37))
            switch "" {
            case "":
                v5 = (1e10 + (v5 * v5))
            case "a":
                var v42 int32 = 116
                _ = v42
            }
        }
    case "foo bar":
        switch v37 {
        case 746:
            v37 += v37
        case 4294967295:
            if (0 > -32768) {
                for fi43 := 0; fi43 < 15; fi43++ {
                    switch ("a" + (("hello" + "hello") + "a")) {
                    case "hello":
                        if (false || (!false)) {
                            _ = len(v3)
                            if false {
                                for fi44 := 0; fi44 < 12; fi44++ {
                                    var v45 float32 = 1e10
                                    _ = fi44
                                    _ = v45
                                }
                            } else {
                                for fi46 := 0; fi46 < 11; fi46++ {
                                    var v47 int64 = v6
                                    v6 += (v6 / ((1) | 1))
                                    _ = fi46
                                    _ = v47
                                }
                                var v48 int = (v37 - 32767)
                                for fi49 := 0; fi49 < 3; fi49++ {
                                    var v50 uint64 = 42
                                    v1 = (v1 ^ v1)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi49
                                    _ = v50
                                }
                                _ = v48
                            }
                        }
                    case "a":
                        for fi51 := 0; fi51 < 13; fi51++ {
                            var v52 float32 = ((1.5 + 0.1) + (((1.5 * float32(float64(3.14159))) + 3.14159) / 1.0))
                            var v53 int = fi51
                            if true {
                                if false {
                                    v54 := GenS4{F0: 948}
                                    fmt.Println(v4, v37)
                                    _ = v54
                                }
                                if false {
                                    v5 = v5
                                } else {
                                    v55 := GenS6{F0: ""}
                                    if len(v3) > 0 {
                                        v3[0] = (("hello" + "\n\t") + ("foo bar" + ("foo bar" + "a")))
                                    }
                                    v6 += (v6 + 0)
                                    _ = v55
                                }
                            }
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                    case "foo bar":
                        switch -2147483648 {
                        case 9223372036854775807:
                            if ((!true) || false) {
                                v5 -= -2.25
                                switch fi43 {
                                case 127:
                                    if ((^(v1)) == (v1 & (v1 % (((v1 * (^(v1)))) | 1)))) {
                                        v1 -= 42
                                        for fi56 := 0; fi56 < 5; fi56++ {
                                            switch "a" {
                                            case "hello":
                                                v6 = 783
                                            case "foo bar":
                                                var v57 float64 = (v5 + (float64(0.1) + 0.1))
                                                _ = v57
                                            case "a":
                                                var v58 float32 = 1e10
                                                _ = v58
                                            default:
                                                delete(v2, v37)
                                            }
                                            v59 := GenS10{F0: 42, F1: (v6 - (^(v6)))}
                                            var v60 int8 = 100
                                            _ = len(v2)
                                            if (9223372036854775807 == (fi56 * fi56)) {
                                                var v61 uint64 = 2
                                                v6 = (^(v6))
                                                v62 := GenS11{F0: 215, F1: false, F2: 127}
                                                _ = v61
                                                _ = v62
                                            } else {
                                                v63 := GenS11{F0: uint8(v38), F1: ("a" >= "hello"), F2: uint8(v1)}
                                                fmt.Println(v5, v60, fi43)
                                                if true {
                                                    delete(v2, v37)
                                                    var v64 float32 = float32(v5)
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        fmt.Println(v4)
                                                    case "a":
                                                        v38 -= v4
                                                    default:
                                                        v65 := GenS3{F0: 759, F1: (fi43 * fi43), F2: 1}
                                                        _ = v65
                                                    }
                                                    _ = v64
                                                }
                                                _ = v63
                                            }
                                            _ = fi56
                                            _ = v59
                                            _ = v60
                                        }
                                    } else {
                                        _ = len(v2)
                                        switch (-(v37)) {
                                        case 1:
                                            var v66 uint = 928
                                            _ = v66
                                        case 4294967295:
                                            v6 += (v6 | (v6 / ((v6) | 1)))
                                        case 236:
                                            v5 = 0.1
                                        }
                                    }
                                default:
                                    fmt.Println(v5, v4, v4)
                                }
                                for fi67 := 0; fi67 < 14; fi67++ {
                                    if (0.1 == v5) {
                                        v68 := v2[42]
                                        fmt.Println(v4, v5)
                                        _ = v68
                                    } else {
                                        v69 := GenS12{F0: (v1 + (v1 % (((v1 << (((v1 + v1)) & 63))) | 1))), F1: -1, F2: -1}
                                        var v70 int32 = -1
                                        _ = v69
                                        _ = v70
                                    }
                                    v5 = -2.25
                                    if (0 > 2147483647) {
                                        var v71 uint32 = v4
                                        if false {
                                            v6 -= (^(v6))
                                            v37 += v37
                                            var v72 uint16 = 2
                                            _ = v72
                                        } else {
                                            if (v1 <= 255) {
                                                var v73 int32 = 2147483647
                                                v1 = v1
                                                fmt.Println(v73, v5, fi43)
                                                _ = v73
                                            }
                                            var v74 int32 = 0
                                            if false {
                                                v71 += (v71 + (^(v71)))
                                                fmt.Println(v4, v38, fi43)
                                            } else {
                                                var v75 int8 = -1
                                                var v76 string = "\n\t"
                                                var v77 int16 = 42
                                                if false {
                                                    v5 += 1.5
                                                }
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = v75
                                                _ = v76
                                                _ = v77
                                            }
                                            _ = v74
                                        }
                                        var v78 string = "a"
                                        _ = v71
                                        _ = v78
                                    }
                                    _ = fi67
                                }
                            }
                        default:
                            for fi79 := 0; fi79 < 17; fi79++ {
                                fmt.Println(v38, v37)
                                v37 = v37
                                _ = fi79
                            }
                        }
                    default:
                        v6 += v6
                    }
                    _ = fi43
                }
            } else {
                v3 = append(v3, "hello")
                delete(v2, v37)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
