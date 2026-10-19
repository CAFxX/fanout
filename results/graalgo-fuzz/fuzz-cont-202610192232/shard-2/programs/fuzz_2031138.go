package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 int32 }
type GenS7 struct { F0 uint8; F1 bool }

func main() {
    var v1 uint64 = 908
    v2 := make(map[string]int)
    var v3 int32 = -2147483648
    v4 := []int{}
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 bool = false
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    fmt.Println(v3)
    switch -9223372036854775808 {
    case 9223372036854775807:
        fmt.Println(v1, v3)
    case 42:
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 string = "foo bar"
            var v10 string = ((v9 + v9) + "")
            v11 := GenS1{F0: (v1 + 42), F1: 2}
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    case 1:
        var v12 uint = 42
        _ = v12
    default:
        v13 := GenS1{F0: 18446744073709551615, F1: 2}
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        v15 := GenS2{F0: -32768, F1: 65535}
        fmt.Println(fi14)
        fmt.Println(v1)
        _ = fi14
        _ = v15
    }
    var v16 uint32 = 437
    var v17 int64 = -9223372036854775808
    if false {
        fmt.Println(v17, v16)
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        fmt.Println(v1, v1, fi18)
        if false {
            for fi19 := 0; fi19 < 6; fi19++ {
                if false {
                    v20 := GenS1{F0: uint64(v16), F1: 149}
                    if true {
                        fmt.Println(v16, v17, fi19)
                        if false {
                            if (1 < 1000) {
                                v16 += v16
                                v21 := GenS2{F0: 32767, F1: (^(v16))}
                                _ = v21
                            }
                            if (false && ((^(v1)) >= v1)) {
                                var v22 uint8 = 211
                                if (false || false) {
                                    var v23 int16 = int16(v17)
                                    v17 -= v17
                                    _ = v23
                                }
                                _ = v22
                            }
                            if false {
                                if (65535 != 32767) {
                                    v4 = append(v4, (-(fi18)))
                                } else {
                                    v17 -= (-(v17))
                                    var v24 float64 = (1.5 + ((float64(0.0) / 1e10) * (((1.5 - 3.14159) / -2.25) + float64((3.14159 - 3.14159)))))
                                    if false {
                                        var v25 int64 = 0
                                        v26 := GenS3{F0: (42 >= 2), F1: 190}
                                        _ = v25
                                        _ = v26
                                    } else {
                                        if (true && false) {
                                            v27 := v2["foo bar"]
                                            _ = v27
                                        }
                                        var v28 string = "hello"
                                        delete(v2, ("\n\t" + "foo bar"))
                                        _ = v28
                                    }
                                    var v29 uint16 = 816
                                    _ = v24
                                    _ = v29
                                }
                                v30 := GenS3{F0: (1.5 <= float32((3.14159 * (0.0 * 0.1)))), F1: 200}
                                v31 := GenS4{F0: 397, F1: 32767}
                                _ = v30
                                _ = v31
                            }
                        } else {
                            v16 = uint32(v1)
                            v16 = 42
                        }
                        v4 = append(v4, 128)
                    } else {
                        if true {
                            switch "hello" {
                            case "foo bar":
                                v32 := v2[("\n\t" + ("foo bar" + "hello"))]
                                _ = v32
                            case "\n\t":
                                switch fi19 {
                                case -95:
                                    v3 = (v3 - -914)
                                case 1:
                                    fmt.Println(fi18)
                                    v3 = (v3 * (v3 ^ (-(v3))))
                                    v17 = v17
                                case -523:
                                    switch "a" {
                                    case "\n\t":
                                        switch fi19 {
                                        case 127:
                                            v16 = (v16 % ((446) | 1))
                                        default:
                                            if len(v4) > 0 {
                                                v4[0] = (fi18 / (((fi19 * fi19)) | 1))
                                            }
                                        }
                                        var v33 int = 480
                                        _ = v33
                                    }
                                default:
                                    switch (fi18 * (fi19 % ((fi19) | 1))) {
                                    case 846:
                                        v34 := GenS5{F0: -1}
                                        _ = v34
                                    case 722:
                                        fmt.Println(fi19)
                                    case 729:
                                        v3 = (v3 ^ v3)
                                    }
                                }
                            default:
                                v35 := GenS4{F0: 18446744073709551615, F1: 919}
                                _ = v35
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            var v36 int64 = v17
                            fmt.Println(v36)
                            var v37 int16 = 42
                            _ = v36
                            _ = v37
                        }
                        v4 = append(v4, 0)
                    }
                    var v38 uint32 = (v16 & v16)
                    if false {
                        var v39 float32 = -2.25
                        v3 = v3
                        v40 := GenS5{F0: -128}
                        _ = v39
                        _ = v40
                    } else {
                        if (!(!((v1 == v1) && true))) {
                            v3 = v3
                            v3 = v3
                            v3 = -2147483648
                        } else {
                            v3 = (v3 >> ((42) & 63))
                        }
                        if (!true) {
                            var v41 int32 = v3
                            _ = v41
                        }
                    }
                    switch ("a" + "\n\t") {
                    case "hello":
                        v4 = append(v4, (fi18 & (fi18 + (fi18 ^ fi18))))
                    case "a":
                        fmt.Println(v3, v1)
                    default:
                        switch (fi18 >> ((2) & 63)) {
                        case 128:
                            v1 += (v1 & v1)
                        case 9223372036854775807:
                            v42 := GenS4{F0: 4294967295, F1: 32767}
                            _ = v42
                        case 744:
                            switch ("hello" + "hello") {
                            case "\n\t":
                                switch fi19 {
                                case 1:
                                    v38 -= 1
                                default:
                                    var v43 int32 = 2147483647
                                    switch (-(fi18)) {
                                    case 42:
                                        switch 534 {
                                        case 128:
                                            v3 = (v43 - (-(v43)))
                                        case 4294967295:
                                            if (true || (true && true)) {
                                                var v44 uint64 = (v1 % ((v1) | 1))
                                                v45 := v2["hello"]
                                                var v46 uint64 = (^(v44))
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            } else {
                                                fmt.Println(v16, v43, v17)
                                                var v47 uint = 65535
                                                _ = v47
                                            }
                                        default:
                                            v43 += -418
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v1, fi18)
                                    }
                                    _ = v43
                                }
                            case "":
                                var v48 int16 = -32768
                                _ = v48
                            default:
                                var v49 float64 = 1e10
                                v17 -= (v17 % (((v17 - 9223372036854775807)) | 1))
                                _ = v49
                            }
                        default:
                            if true {
                                if false {
                                    var v50 uint = 782
                                    _ = v50
                                }
                            } else {
                                var v51 uint = 1
                                fmt.Println(v17)
                                _ = v51
                            }
                        }
                    }
                    _ = v20
                    _ = v38
                }
                v16 = 323
                _ = fi19
            }
        }
        v1 += 18446744073709551615
        _ = fi18
    }
    var v52 int = int(v16)
    _ = len(v4)
    v1 = v1
    if false {
        fmt.Println(v16, v52)
        if true {
            var v53 float32 = (-2.25 - (0.1 - (float32(0.0) + (1.5 * 1.5))))
            _ = v53
        } else {
            var v54 string = "a"
            v55 := v2[("a" + (v54 + ""))]
            v56 := GenS4{F0: 2, F1: 788}
            _ = v54
            _ = v55
            _ = v56
        }
        switch "foo bar" {
        case "\n\t":
            switch "a" {
            case "a":
                switch ("" + "\n\t") {
                case "a":
                    var v57 int16 = 1
                    _ = v57
                case "\n\t":
                    for fi58 := 0; fi58 < 2; fi58++ {
                        var v59 uint32 = v16
                        var v60 uint = 42
                        v1 += uint64(v16)
                        _ = fi58
                        _ = v59
                        _ = v60
                    }
                default:
                    v3 -= v3
                }
            case "":
                v61 := GenS2{F0: 1000, F1: (^(v16))}
                _ = v61
            default:
                var v62 uint64 = (v1 ^ (v1 >> ((42) & 63)))
                _ = v62
            }
        case "a":
            var v63 bool = (!(1 <= -116))
            _ = v63
        default:
            v64 := GenS6{F0: 16, F1: v3}
            _ = v64
        }
    } else {
        v17 -= (v17 & (v17 | (v17 ^ v17)))
        var v65 bool = (!true)
        for fi66 := 0; fi66 < 16; fi66++ {
            for fi67 := 0; fi67 < 19; fi67++ {
                var v68 bool = false
                v17 -= 9223372036854775807
                var v69 float64 = ((1.5 * -2.25) - ((0.0 * 0.1) - 1e10))
                _ = fi67
                _ = v68
                _ = v69
            }
            fmt.Println(v16, v65, v3)
            if ((42 == 42) && (uint32(v1) != (^(v16)))) {
                for fi70 := 0; fi70 < 9; fi70++ {
                    v71 := GenS7{F0: 42, F1: v65}
                    v1 -= v1
                    fmt.Println(v1, fi66, v65)
                    _ = fi70
                    _ = v71
                }
            } else {
                v17 += v17
                v72 := GenS3{F0: (!v65), F1: 247}
                v2[(("" + ("\n\t" + ("\n\t" + "foo bar"))) + (("\n\t" + "hello") + ""))] = (^(v52))
                _ = v72
            }
            _ = fi66
        }
        _ = v65
    }
    for fi73 := 0; fi73 < 12; fi73++ {
        v3 = 42
        _ = fi73
    }
    v74 := GenS1{F0: 18446744073709551615, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v74
}
