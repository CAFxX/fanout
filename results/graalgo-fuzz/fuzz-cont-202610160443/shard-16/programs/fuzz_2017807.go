package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint64; F1 int32; F2 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64; F1 uint8; F2 uint16 }

func main() {
    var v1 uint32 = 2
    v2 := []uint32{v1, v1, v1}
    v3 := []string{}
    v4 := make(map[string]bool)
    v5 := make(map[string]string)
    v1 += 701
    var v6 bool = (((true || (!(1e10 >= ((1e10 / 1e10) / 0.1)))) || (v1 > v1)) || (0.0 <= (0.1 * 0.1)))
    var v7 float64 = 3.14159
    if len(v3) > 0 {
        v3[0] = (("a" + "") + ("foo bar" + "a"))
    }
    switch 4294967295 {
    case 9223372036854775807:
        v6 = (!(v7 < float64(0.1)))
    case 1:
        v8 := GenS1{F0: 981, F1: 987, F2: 42}
        _ = v8
    default:
        var v9 float64 = 0.0
        _ = v9
    }
    fmt.Println(v7, v1)
    if false {
        if (!(65535 != 0)) {
            var v10 float32 = 1.5
            _ = v10
        } else {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 uint32 = 292
                var v13 int16 = 570
                v7 -= (v7 - (0.0 + (0.0 * v7)))
                _ = fi11
                _ = v12
                _ = v13
            }
            fmt.Println(v6, v1, v6)
            for fi14 := 0; fi14 < 14; fi14++ {
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 int32 = -2147483648
                    if v6 {
                        var v17 float32 = 0.0
                        delete(v5, ("foo bar" + "\n\t"))
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
        }
        if false {
            if true {
                var v18 uint16 = 32767
                for fi19 := 0; fi19 < 11; fi19++ {
                    switch -567 {
                    case 9223372036854775807:
                        v1 += (v1 ^ 2147483648)
                    }
                    switch (fi19 * (^(fi19))) {
                    case -9223372036854775808:
                        switch ("hello" + "") {
                        case "":
                            v18 = v18
                        }
                    case 42:
                        var v20 uint = 591
                        _ = v20
                    default:
                        if ((!v6) && (0 == 65535)) {
                            v1 = v1
                            v6 = true
                            fmt.Println(v18, v7, v1)
                        }
                    }
                    _ = fi19
                }
                var v21 bool = (!(true || (!(4294967295 != 0))))
                _ = v18
                _ = v21
            } else {
                for fi22 := 0; fi22 < 1; fi22++ {
                    var v23 uint8 = 128
                    var v24 float64 = v7
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                v1 = v1
            }
            for fi25 := 0; fi25 < 6; fi25++ {
                if v6 {
                    v1 -= (v1 * v1)
                    if true {
                        v6 = (!(791 > 2))
                        v26 := GenS1{F0: -459, F1: 459, F2: 1}
                        v27 := GenS2{F0: v1}
                        _ = v26
                        _ = v27
                    }
                    v1 -= v1
                }
                switch "foo bar" {
                case "hello":
                    var v28 uint64 = 18446744073709551615
                    _ = v28
                case "a":
                    v1 = (v1 | (^(v1)))
                default:
                    var v29 uint = 255
                    _ = v29
                }
                _ = fi25
            }
            fmt.Println(v7, v1, v6)
        }
    }
    if v6 {
        v6 = true
        var v30 int64 = 9223372036854775807
        _ = v30
    } else {
        var v31 int8 = 42
        _ = v31
    }
    fmt.Println(v7, v6)
    var v32 uint8 = 127
    fmt.Println(v7, v1)
    v33 := GenS3{F0: 2, F1: 275}
    switch 32767 {
    case -2147483648:
        _ = len(v3)
    case -481:
        switch -9223372036854775808 {
        case 32767:
            v6 = v6
        case -961:
            for fi34 := 0; fi34 < 15; fi34++ {
                for fi35 := 0; fi35 < 9; fi35++ {
                    var v36 int64 = 846
                    v4["\n\t"] = (4294967295 >= 918)
                    _ = fi35
                    _ = v36
                }
                for fi37 := 0; fi37 < 13; fi37++ {
                    v6 = v6
                    _ = fi37
                }
                _ = fi34
            }
        }
    case 970:
        var v38 uint32 = 911
        _ = v38
    default:
        var v39 uint8 = v32
        _ = v39
    }
    if (int64(v1) == 1) {
        fmt.Println(v6)
        var v40 int8 = -1
        if ((0 < 213) && (true || (573 == 2))) {
            v40 = 127
        } else {
            switch ("hello" + "a") {
            case "":
                var v41 int16 = -620
                _ = v41
            case "\n\t":
                v42 := GenS3{F0: uint64(v1), F1: 2147483647}
                _ = v42
            default:
                var v43 float32 = 0.1
                _ = v43
            }
            v44 := GenS4{F0: 1, F1: 42, F2: v6}
            v32 = 128
            _ = v44
        }
        _ = v40
    } else {
        for fi45 := 0; fi45 < 9; fi45++ {
            switch "hello" {
            case "\n\t":
                var v46 uint8 = 255
                _ = v46
            case "foo bar":
                for fi47 := 0; fi47 < 13; fi47++ {
                    v2 = append(v2, v1)
                    switch "" {
                    case "":
                        v48 := GenS4{F0: 18, F1: -1, F2: (-971 <= -492)}
                        _ = v48
                    case "\n\t":
                        fmt.Println(v7, fi47, fi45)
                    }
                    v32 += 1
                    _ = fi47
                }
            default:
                if v6 {
                    v49 := GenS3{F0: 77, F1: 1}
                    switch "hello" {
                    case "foo bar":
                        v1 -= (v1 | v1)
                    default:
                        switch 2147483647 {
                        case 176:
                            var v50 int64 = -9223372036854775808
                            _ = v50
                        case 127:
                            for fi51 := 0; fi51 < 18; fi51++ {
                                var v52 uint64 = 4294967295
                                v7 = (v7 * v7)
                                switch (fi45 % ((fi45) | 1)) {
                                case 1:
                                    if false {
                                        var v53 int16 = 1000
                                        var v54 int16 = v53
                                        var v55 float64 = v7
                                        var v56 uint = 490
                                        v52 = (v52 / (((v52 / ((2) | 1))) | 1))
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                default:
                                    var v57 int = (fi45 * fi51)
                                    _ = v57
                                }
                                _ = fi51
                                _ = v52
                            }
                        default:
                            v3 = append(v3, "a")
                        }
                    }
                    v32 -= v32
                    _ = v49
                } else {
                    v58 := GenS5{F0: 9223372036854775807}
                    switch "hello" {
                    case "hello":
                        if v6 {
                            v7 = float64(0.1)
                            var v59 uint16 = 65535
                            v7 = (v7 - 3.14159)
                            _ = v59
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 2147483648
                            }
                            switch (fi45 | (^(fi45))) {
                            case 931:
                                var v60 int64 = -37
                                _ = v60
                            }
                            v61 := GenS1{F0: 0, F1: 39, F2: 42}
                            _ = v61
                        }
                    case "\n\t":
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                    case "foo bar":
                        _ = len(v2)
                    default:
                        for fi62 := 0; fi62 < 1; fi62++ {
                            if v6 {
                                if (true && v6) {
                                    fmt.Println(fi45)
                                    v6 = (!(!v6))
                                }
                                fmt.Println(fi62)
                                v2 = append(v2, (v1 << ((554) & 63)))
                            } else {
                                fmt.Println(v7, v7, fi62)
                                v7 = ((float64((float32(v7) + 1e10)) / -2.25) / ((v7 * -2.25) * v7))
                            }
                            fmt.Println(v7)
                            var v63 uint32 = 0
                            _ = fi62
                            _ = v63
                        }
                    }
                    for fi64 := 0; fi64 < 19; fi64++ {
                        v65 := GenS6{F0: 9223372036854775807}
                        _ = fi64
                        _ = v65
                    }
                    _ = v58
                }
            }
            v32 = 42
            _ = len(v2)
            _ = fi45
        }
        for fi66 := 0; fi66 < 15; fi66++ {
            if (!(!(42 <= 0))) {
                if (!v6) {
                    v67 := GenS7{F0: (false && (v32 < v32))}
                    _ = v67
                } else {
                    v68 := GenS5{F0: 979}
                    v69 := GenS1{F0: -2147483648, F1: 493, F2: 255}
                    if v6 {
                        if false {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            delete(v4, "a")
                        }
                        v6 = true
                    } else {
                        switch int(v1) {
                        case 42:
                            v70 := v5["foo bar"]
                            _ = v70
                        case -257:
                            v71 := GenS8{F0: (v32 ^ v32)}
                            _ = v71
                        case 32767:
                            v7 = (((float64((-2.25 / -2.25)) - 3.14159) + (((0.0 / 1.0) * 0.0) / v7)) + (v7 + v7))
                        default:
                            if true {
                                v32 = (v32 / (((v32 & v32)) | 1))
                            }
                        }
                        v7 = float64(0.0)
                        var v72 uint64 = 4294967295
                        _ = v72
                    }
                    _ = v68
                    _ = v69
                }
                v73 := GenS1{F0: 0, F1: 4294967295, F2: 808}
                v74 := GenS9{F0: 42, F1: (v32 ^ v32), F2: 499}
                _ = v73
                _ = v74
            } else {
                for fi75 := 0; fi75 < 20; fi75++ {
                    var v76 uint32 = (v1 / (((v1 + (v1 / ((uint32(v32)) | 1)))) | 1))
                    fmt.Println(v32, v76)
                    _ = fi75
                    _ = v76
                }
            }
            _ = fi66
        }
    }
    for fi77 := 0; fi77 < 11; fi77++ {
        v7 += -2.25
        for fi78 := 0; fi78 < 9; fi78++ {
            v79 := GenS8{F0: v32}
            if false {
                v7 -= 0.0
                v80 := GenS6{F0: 42}
                switch ("a" + ("" + ((("\n\t" + "a") + ("\n\t" + "")) + (("\n\t" + "") + "a")))) {
                case "hello":
                    v32 -= (v32 << ((4294967295) & 63))
                case "a":
                    v81 := GenS3{F0: uint64(v32), F1: 42}
                    _ = v81
                default:
                    if len(v3) > 0 {
                        v3[0] = "\n\t"
                    }
                }
                _ = v80
            } else {
                v6 = v6
                fmt.Println(fi78)
            }
            _ = fi78
            _ = v79
        }
        switch ("foo bar" + "hello") {
        case "hello":
            _ = len(v2)
        default:
            v7 -= 3.14159
        }
        _ = fi77
    }
    var v82 uint = 2
    switch 127 {
    case -277:
        v7 = float64(((1e10 / -2.25) * (-2.25 + 1e10)))
    case 42:
        var v83 int16 = 1
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v82
}
