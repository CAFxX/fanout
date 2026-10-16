package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 uint64; F1 int; F2 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 int32; F2 int64 }
type GenS9 struct { F0 uint8; F1 uint32; F2 int64 }
func genFn0() uint32 {
    for fi4 := 0; fi4 < 7; fi4++ {
        fmt.Println(fi4)
        switch "foo bar" {
        case "":
            var v5 int8 = -19
            _ = v5
        default:
            v6 := GenS1{F0: 42, F1: 42}
            _ = v6
        }
        for fi7 := 0; fi7 < 20; fi7++ {
            switch "\n\t" {
            case "":
                fmt.Println(fi4, fi4)
            }
            switch "a" {
            case "":
                fmt.Println(fi4, fi4)
            default:
                var v8 int16 = -32768
                _ = v8
            }
            var v9 int64 = -1
            v10 := GenS2{F0: 1, F1: 641}
            v9 = v9
            _ = fi7
            _ = v10
            _ = v9
        }
        _ = fi4
    }
    if true {
        v11 := GenS3{F0: 1, F1: 42, F2: 658}
        _ = v11
    }
    v12 := GenS4{F0: -1, F1: 65535}
    _ = v12
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (3.14159 - (0.0 * ((0.0 - 0.0) + 1e10)))
    v3 := make(map[uint32]int)
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 uint = 841
        if false {
            if (fi13 > fi13) {
                for fi15 := 0; fi15 < 14; fi15++ {
                    fmt.Println(fi13, v14, fi13)
                    var v16 float32 = (0.1 - float32(1e10))
                    switch "a" {
                    case "a":
                        v17 := GenS5{F0: 965, F1: (fi15 & (fi15 << (((v14 * v14)) & 63))), F2: (42 != 42)}
                        _ = v17
                    case "hello":
                        v18 := GenS6{F0: (fi13 | fi13)}
                        _ = v18
                    default:
                        fmt.Println(v14, v14)
                    }
                    _ = fi15
                    _ = v16
                }
                for fi19 := 0; fi19 < 8; fi19++ {
                    v14 = (^(v14))
                    _ = fi19
                }
            } else {
                var v20 uint8 = 0
                fmt.Println(v14)
                _ = v20
            }
            v21 := []int{(fi13 / ((fi13) | 1))}
            var v22 uint = 645
            _ = v21
            _ = v22
        }
        _ = fi13
        _ = v14
    }
    v23 := GenS1{F0: 833, F1: 2}
    var v24 int = 32767
    var v25 string = "foo bar"
    var v26 float64 = 3.14159
    for fi27 := 0; fi27 < 3; fi27++ {
        v28 := v3[2147483648]
        v29 := GenS6{F0: (fi27 & 2147483647)}
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 uint = 769
            _ = fi30
            _ = v31
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    v25 = v25
    for fi32 := 0; fi32 < 12; fi32++ {
        v33 := []int{(v24 - (v24 + 4294967295)), 32768, fi32}
        fmt.Println(v25, v2, v24)
        for fi34 := 0; fi34 < 6; fi34++ {
            v35 := GenS7{F0: v25}
            v36 := GenS7{F0: ""}
            if (v25 >= v25) {
                switch 127 {
                case -9223372036854775808:
                    var v37 bool = false
                    _ = v37
                case 32767:
                    var v38 bool = false
                    _ = v38
                case 2147483647:
                    v39 := GenS5{F0: 830, F1: fi32, F2: (!false)}
                    _ = v39
                }
                var v40 uint64 = 18446744073709551615
                var v41 uint = 679
                _ = v40
                _ = v41
            }
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi32
        _ = v33
    }
    if (1 == -163) {
        fmt.Println(v24)
        switch (v24 ^ v24) {
        case 2147483647:
            _ = len(v1)
        case 32768:
            for fi42 := 0; fi42 < 16; fi42++ {
                var v43 uint64 = 1
                v44 := GenS6{F0: 42}
                for fi45 := 0; fi45 < 19; fi45++ {
                    if ((v43 ^ (^(v43))) <= v43) {
                        fmt.Println(v24)
                    }
                    v2 = v2
                    v43 = (v43 + 1)
                    _ = fi45
                }
                _ = fi42
                _ = v43
                _ = v44
            }
        case 1:
            if (!false) {
                var v46 uint = 18446744073709551615
                switch v24 {
                case 1:
                    switch v25 {
                    case "":
                        for fi47 := 0; fi47 < 14; fi47++ {
                            fmt.Println(v46, v46)
                            var v48 uint64 = 0
                            if false {
                                var v49 int = -1
                                if (false || (-332 < -9223372036854775808)) {
                                    v46 -= v46
                                }
                                _ = v49
                            } else {
                                fmt.Println(v46, v46)
                                if (!(!(!(0 > 0)))) {
                                    for fi50 := 0; fi50 < 18; fi50++ {
                                        v51 := GenS4{F0: 42, F1: 49}
                                        v2 += 3.14159
                                        _ = fi50
                                        _ = v51
                                    }
                                    v52 := v1[4294967295]
                                    _ = v52
                                } else {
                                    var v53 uint = v46
                                    if (fi47 == (^(fi47))) {
                                        if true {
                                            v24 = (fi47 * (fi47 * v24))
                                            var v54 uint32 = 212
                                            _ = v54
                                        } else {
                                            for fi55 := 0; fi55 < 2; fi55++ {
                                                fmt.Println(v24, v53)
                                                var v56 bool = (true && ((32767 >= 32767) || true))
                                                v3[65535] = (^(v24))
                                                _ = fi55
                                                _ = v56
                                            }
                                        }
                                        var v57 string = "foo bar"
                                        _ = v57
                                    } else {
                                        v24 = 128
                                        var v58 int64 = 636
                                        _ = v58
                                    }
                                    if true {
                                        _ = len(v1)
                                        fmt.Println(v26)
                                    } else {
                                        for fi59 := 0; fi59 < 13; fi59++ {
                                            var v60 uint32 = uint32(v46)
                                            var v61 uint = v46
                                            _ = fi59
                                            _ = v60
                                            _ = v61
                                        }
                                        v62 := GenS6{F0: (fi47 & (fi47 << ((65535) & 63)))}
                                        _ = v62
                                    }
                                    _ = v53
                                }
                            }
                            _ = fi47
                            _ = v48
                        }
                    case "a":
                        var v63 int16 = -574
                        _ = v63
                    case "\n\t":
                        if ((-2147483648 <= (v24 ^ 42)) || false) {
                            var v64 uint32 = 891
                            _ = v64
                        } else {
                            fmt.Println(v2, v25)
                        }
                    }
                default:
                    var v65 float32 = 0.0
                    _ = v65
                }
                for fi66 := 0; fi66 < 20; fi66++ {
                    for fi67 := 0; fi67 < 3; fi67++ {
                        if (42 >= 0) {
                            var v68 string = v25
                            delete(v1, (fi66 + fi67))
                            v46 = (v46 + v46)
                            v2 = (float32((((v26 * float64((0.0 - 1.5))) * -2.25) - (float64(0.0) / 0.1))) / v2)
                            v69 := GenS7{F0: "\n\t"}
                            _ = v68
                            _ = v69
                        } else {
                            v24 += -2147483648
                            v70 := []int{fi66, (-(fi67))}
                            if ((true && false) && ((1 <= 485) && true)) {
                                v2 = 1.5
                            } else {
                                var v71 int = (fi66 | -633)
                                var v72 bool = (!false)
                                _ = v71
                                _ = v72
                            }
                            if (0 <= 4294967295) {
                                var v73 uint32 = 225
                                _ = v73
                            } else {
                                if len(v70) >= 2 {
                                    v70 = v70[1:len(v70)]
                                }
                            }
                            _ = v70
                        }
                        v26 -= float64(((((1e10 + float32(v26)) / float32(v26)) / v2) * -2.25))
                        var v74 float64 = 1e10
                        _ = fi67
                        _ = v74
                    }
                    var v75 bool = false
                    _ = fi66
                    _ = v75
                }
                _ = v46
            }
        default:
            var v76 uint16 = 270
            _ = v76
        }
    } else {
        switch 32768 {
        case 32767:
            if false {
                v77 := []int{v24, (v24 / (((v24 & v24)) | 1)), 127}
                _ = v77
            } else {
                if false {
                    if (1000 == 42) {
                        switch v24 {
                        case 0:
                            switch v25 {
                            case "\n\t":
                                v2 -= 0.1
                            case "foo bar":
                                v26 -= 3.14159
                            case "a":
                                var v78 uint64 = 961
                                _ = v78
                            }
                        case 9223372036854775807:
                            v25 = (v25 + "")
                        case -302:
                            fmt.Println(v26, v26)
                        default:
                            v26 = 1e10
                        }
                    } else {
                        fmt.Println(v26, v24, v2)
                    }
                } else {
                    v26 = float64((-2.25 * v2))
                    for fi79 := 0; fi79 < 20; fi79++ {
                        if false {
                            v80 := GenS8{F0: 1, F1: 42, F2: 9223372036854775807}
                            v26 = (v26 / float64(float32((((float64(v2) / v26) - -2.25) - (v26 / v26)))))
                            switch "a" {
                            case "a":
                                var v81 uint32 = 0
                                _ = v81
                            default:
                                var v82 bool = (2 < 4294967295)
                                _ = v82
                            }
                            _ = v80
                        }
                        fmt.Println(v25, v24, v26)
                        if (124 <= 127) {
                            for fi83 := 0; fi83 < 4; fi83++ {
                                switch 245 {
                                case 1:
                                    fmt.Println(v2)
                                }
                                v2 += 0.0
                                _ = fi83
                            }
                            switch (("" + ("hello" + v25)) + ("foo bar" + v25)) {
                            case "foo bar":
                                var v84 string = v25
                                _ = v84
                            case "":
                                if false {
                                    fmt.Println(v24, v2, v25)
                                }
                            }
                            if (6 <= -54) {
                                delete(v1, (fi79 << ((0) & 63)))
                            }
                        }
                        _ = fi79
                    }
                    for fi85 := 0; fi85 < 2; fi85++ {
                        fmt.Println(v2)
                        fmt.Println(v26, v25, fi85)
                        v1[(^(v24))] = 2147483647
                        _ = fi85
                    }
                }
                v86 := GenS9{F0: 128, F1: 2147483648, F2: 9223372036854775807}
                _ = v86
            }
        case 463:
            v3[552] = v24
        case 9223372036854775807:
            v26 = v26
        default:
            var v87 uint64 = 42
            _ = v87
        }
        var v88 uint16 = 1
        for fi89 := 0; fi89 < 7; fi89++ {
            delete(v1, 2147483647)
            _ = fi89
        }
        _ = v88
    }
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
