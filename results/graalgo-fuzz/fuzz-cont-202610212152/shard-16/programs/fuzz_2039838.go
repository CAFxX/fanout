package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 int64; F2 int16 }
type GenS6 struct { F0 uint16; F1 int64; F2 int64 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 bool; F1 uint8 }
type GenS10 struct { F0 int; F1 uint64 }
func genFn0(p0 uint) uint {
    for fi4 := 0; fi4 < 11; fi4++ {
        p0 = p0
        for fi5 := 0; fi5 < 7; fi5++ {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(p0, fi4, fi5)
            case "hello":
                v6 := GenS1{F0: -815}
                _ = v6
            case "":
                v7 := []int{1, fi5, (fi5 >> ((283) & 63))}
                _ = v7
            default:
                fmt.Println(p0, fi4, p0)
            }
            p0 = 1
            var v8 uint32 = 233
            v9 := GenS1{F0: -9223372036854775808}
            _ = fi5
            _ = v8
            _ = v9
        }
        var v10 uint = (p0 + 65535)
        _ = fi4
        _ = v10
    }
    p0 = p0
    v11 := []int{-951, 718, 4294967295, 207}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    _ = p0
    _ = v11
    return 18446744073709551615
}

func main() {
    var v1 string = ((("" + "") + "") + "foo bar")
    var v2 int8 = 1
    var v3 string = "hello"
    v12 := []int{}
    switch -873 {
    case 4294967295:
        v3 = "a"
    case 87:
        switch "hello" {
        case "\n\t":
            v1 = (v1 + "a")
        case "a":
            var v13 int16 = 185
            _ = v13
        default:
            if (!(32767 > 1000)) {
                var v14 int8 = (-(v2))
                for fi15 := 0; fi15 < 11; fi15++ {
                    v16 := GenS1{F0: 9223372036854775807}
                    var v17 uint16 = 660
                    for fi18 := 0; fi18 < 14; fi18++ {
                        v3 = (v1 + ("" + v3))
                        if false {
                            v2 = v2
                            v3 = (v3 + v3)
                            v19 := GenS2{F0: v17, F1: (v14 / ((v2) | 1)), F2: -1}
                            if false {
                                fmt.Println(v17)
                                if true {
                                    _ = len(v12)
                                    if len(v12) > 0 {
                                        v12[0] = (fi18 * fi18)
                                    }
                                    var v20 string = (("" + v3) + ((v1 + v1) + v3))
                                    _ = v20
                                } else {
                                    v1 = (v1 + "hello")
                                    if false {
                                        v12 = append(v12, fi15)
                                        var v21 uint32 = 2147483647
                                        v17 += (v17 + v17)
                                        v1 = ""
                                        v21 += (v21 + v21)
                                        _ = v21
                                    }
                                }
                                v22 := GenS2{F0: 32767, F1: v2, F2: fi18}
                                _ = v22
                            }
                            _ = v19
                        } else {
                            var v23 uint64 = 2
                            fmt.Println(v3, v17, v23)
                            _ = v23
                        }
                        if len(v12) > 0 {
                            v12[0] = fi15
                        }
                        _ = fi18
                    }
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v24 := GenS3{F0: "a"}
                _ = v14
                _ = v24
            }
        }
    default:
        v25 := GenS4{F0: 0}
        _ = v25
    }
    v2 -= 0
    v1 = "foo bar"
    if false {
        if (!(238 == -234)) {
            if (862 <= 1) {
                if false {
                    fmt.Println(v3, v2)
                }
            }
            var v26 float32 = 0.1
            var v27 bool = (1 > -584)
            _ = v26
            _ = v27
        } else {
            var v28 int16 = -32768
            var v29 uint8 = 1
            _ = v28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 11; fi30++ {
        if true {
            var v31 uint32 = 2147483647
            _ = v31
        } else {
            v32 := GenS5{F0: 1, F1: -632, F2: 1}
            _ = v32
        }
        switch (fi30 * fi30) {
        case 1:
            v2 -= 108
        default:
            v3 = "a"
        }
        v3 = v3
        _ = fi30
    }
    var v33 int64 = 0
    var v34 uint8 = 255
    if (true || true) {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        fmt.Println(v34, v34)
        v35 := GenS6{F0: 32767, F1: int64(v34), F2: (-(v33))}
        _ = v35
    }
    v33 += (v33 | 9223372036854775807)
    var v36 uint = 42
    var v37 float32 = ((0.0 / (((0.1 - 3.14159) + -2.25) * (float32(3.14159) - (1e10 * 1e10)))) * 0.1)
    if (!true) {
        var v38 int16 = 83
        v39 := GenS7{F0: 42}
        for fi40 := 0; fi40 < 6; fi40++ {
            fmt.Println(v38, fi40, v37)
            _ = fi40
        }
        _ = v38
        _ = v39
    }
    var v41 int32 = -178
    for fi42 := 0; fi42 < 2; fi42++ {
        if ((-2.25 != float64(-2.25)) || false) {
            if (((1.5 - float64(v37)) < 0.0) || true) {
                v41 = v41
                v43 := GenS6{F0: 2, F1: (v33 >> (((v36 << ((v36) & 63))) & 63)), F2: 178}
                _ = v43
            } else {
                if len(v12) > 0 {
                    v12[0] = (fi42 << (((v36 % ((v36) | 1))) & 63))
                }
                for fi44 := 0; fi44 < 13; fi44++ {
                    v2 = 1
                    var v45 string = v1
                    _ = fi44
                    _ = v45
                }
                for fi46 := 0; fi46 < 1; fi46++ {
                    v3 = "hello"
                    if (!(((v34 + v34) >= (^(v34))) || (!false))) {
                        v36 -= v36
                    }
                    _ = fi46
                }
            }
        } else {
            if true {
                v47 := GenS4{F0: 253}
                var v48 int16 = 1
                fmt.Println(v2, v48)
                _ = v47
                _ = v48
            } else {
                var v49 int32 = v41
                if (!(145 < (v41 & v49))) {
                    fmt.Println(v1, v2)
                } else {
                    if (!false) {
                        if len(v12) > 0 {
                            v12[0] = fi42
                        }
                        fmt.Println(v41, v49, v41)
                        for fi50 := 0; fi50 < 13; fi50++ {
                            var v51 int64 = (v33 | (-(v33)))
                            _ = fi50
                            _ = v51
                        }
                    } else {
                        v49 = 176
                        for fi52 := 0; fi52 < 15; fi52++ {
                            v49 += (v41 >> ((v36) & 63))
                            v3 = v3
                            _ = fi52
                        }
                    }
                    var v53 int = fi42
                    _ = v53
                }
                if true {
                    v54 := GenS8{F0: 0}
                    for fi55 := 0; fi55 < 13; fi55++ {
                        if ((v49 | int32(v34)) == v49) {
                            v56 := GenS1{F0: v33}
                            v49 = (^(v49))
                            _ = v56
                        }
                        v57 := GenS7{F0: (v2 << (((v36 / (((v36 - (v36 % ((2) | 1)))) | 1))) & 63))}
                        _ = fi55
                        _ = v57
                    }
                    _ = v54
                }
                _ = v49
            }
        }
        if false {
            if false {
                var v58 int16 = 42
                _ = v58
            }
            var v59 int32 = (v41 - (v41 / ((-2147483648) | 1)))
            _ = v59
        } else {
            v60 := GenS8{F0: 4294967295}
            fmt.Println(v37, v36)
            _ = v60
        }
        _ = fi42
    }
    if (805 >= 1) {
        v61 := GenS9{F0: (493 < 32767), F1: (v34 & (v34 & (v34 >> (((v36 << ((1) & 63))) & 63))))}
        v62 := GenS10{F0: 9223372036854775807, F1: 0}
        _ = v61
        _ = v62
    }
    switch -907 {
    case 0:
        _ = len(v12)
    case -2147483648:
        if ((v41 << (((^(v36))) & 63)) == (v41 * 1)) {
            var v63 float32 = (v37 + v37)
            _ = v63
        } else {
            if false {
                v64 := GenS4{F0: 2147483647}
                v37 += (0.1 * v37)
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = v64
            }
            if false {
                var v65 int64 = v33
                if true {
                    var v66 float32 = float32((float64((v37 * (v37 + (v37 - (float32(float64((float32(3.14159) + (1e10 - (v37 + ((v37 - v37) + v37)))))) + 1.5))))) + ((1e10 - ((((0.1 - float64((0.0 - 1.5))) - 3.14159) + 0.0) - 0.1)) * 0.0)))
                    v12 = append(v12, 32768)
                    _ = v66
                } else {
                    if false {
                        switch "hello" {
                        case "\n\t":
                            v34 += 0
                        case "":
                            var v67 int64 = v65
                            _ = v67
                        }
                        fmt.Println(v36, v33)
                        if true {
                            switch 356 {
                            case -334:
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            case -986:
                                v12 = append(v12, 42)
                            case -104:
                                v41 = (v41 | (^(v41)))
                            default:
                                var v68 bool = (false && true)
                                _ = v68
                            }
                            var v69 uint = (v36 % ((4294967295) | 1))
                            if true {
                                v70 := GenS10{F0: 42, F1: 4294967295}
                                _ = v70
                            } else {
                                switch 32767 {
                                case 2147483647:
                                    for fi71 := 0; fi71 < 12; fi71++ {
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                        _ = fi71
                                    }
                                case -2147483648:
                                    var v72 int = 2147483647
                                    _ = v72
                                case 9223372036854775807:
                                    v12 = append(v12, 531)
                                }
                                v73 := GenS6{F0: 42, F1: (-(v33)), F2: -1}
                                _ = v73
                            }
                            _ = v69
                        }
                    }
                    fmt.Println(v2)
                    var v74 string = ""
                    _ = v74
                }
                _ = v65
            } else {
                _ = len(v12)
                var v75 int64 = v33
                _ = v75
            }
            var v76 uint32 = 732
            _ = v76
        }
    case 127:
        if len(v12) > 0 {
            v12[0] = -1
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v41
}
