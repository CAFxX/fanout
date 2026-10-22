package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int; F1 uint8; F2 int }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int8; F1 int }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int; F1 bool }
type GenS10 struct { F0 bool; F1 string }
func genFn0() int8 {
    v6 := GenS1{F0: (42 > 127)}
    var v7 uint = 2
    switch 32767 {
    case -2147483648:
        v8 := GenS2{F0: -32768, F1: -1, F2: 149}
        _ = v8
    case 1:
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := []int{fi9}
            v7 += v7
            if len(v10) > 0 {
                v10[0] = (fi9 >> (((v7 ^ (^(v7)))) & 63))
            }
            _ = fi9
            _ = v10
        }
    }
    v11 := []int{32768}
    v12 := GenS3{F0: 255}
    _ = v11
    _ = v12
    _ = v6
    _ = v7
    return 100
}
func genFn1(p0 int64, p1 int16) uint64 {
    var v13 uint = 255
    p1 -= p1
    _ = p0
    _ = p1
    _ = v13
    return 4294967295
}

func main() {
    var v1 string = ""
    v2 := []string{("foo bar" + v1)}
    v3 := make(map[uint32]int)
    var v4 float32 = float32(1.5)
    v5 := make(map[string]uint32)
    var v14 uint8 = 217
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS4{F0: -1, F1: v1}
        if (1 != -128) {
            v1 = "\n\t"
        } else {
            if (v4 != (1.5 / 1.0)) {
                v3[2] = (^(fi15))
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                switch fi17 {
                case -377:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case 128:
                    v18 := GenS5{F0: (-(fi17)), F1: v14, F2: (fi17 | 838)}
                    _ = v18
                }
                _ = fi17
            }
        }
        fmt.Println(v4, v1, fi15)
        _ = fi15
        _ = v16
    }
    var v19 int16 = 1000
    switch ("a" + "a") {
    case "\n\t":
        v1 = "foo bar"
    }
    fmt.Println(v4, v1, v14)
    for fi20 := 0; fi20 < 8; fi20++ {
        fmt.Println(fi20, v14, v4)
        _ = fi20
    }
    v14 -= 128
    switch ((v1 + ((v1 + "") + ("" + ("foo bar" + v1)))) + (v1 + v1)) {
    case "":
        var v21 int64 = 42
        _ = v21
    }
    fmt.Println(v14)
    v14 = 127
    _ = len(v3)
    for fi22 := 0; fi22 < 14; fi22++ {
        v14 = v14
        var v23 uint32 = 2147483648
        _ = fi22
        _ = v23
    }
    v24 := GenS5{F0: -560, F1: v14, F2: 0}
    if (!((1.5 == (-2.25 * 3.14159)) || (357 < 42))) {
        switch "" {
        case "hello":
            switch 2147483647 {
            case -2147483648:
                if false {
                    _ = len(v2)
                    fmt.Println(v14, v1, v1)
                }
            default:
                switch -415 {
                case -415:
                    var v25 float64 = 0.0
                    _ = v25
                case 802:
                    switch v1 {
                    case "a":
                        var v26 float32 = float32(0.0)
                        _ = v26
                    default:
                        switch 0 {
                        case -1:
                            if true {
                                if (!true) {
                                    var v27 int8 = -1
                                    _ = v27
                                } else {
                                    fmt.Println(v1, v19)
                                    switch v1 {
                                    case "":
                                        if (4294967295 <= 0) {
                                            v4 += 3.14159
                                            for fi28 := 0; fi28 < 15; fi28++ {
                                                fmt.Println(fi28)
                                                switch -9223372036854775808 {
                                                case 128:
                                                    switch (v1 + v1) {
                                                    case "\n\t":
                                                        v29 := GenS6{F0: (v14 % ((255) | 1))}
                                                        _ = v29
                                                    case "hello":
                                                        if (-1 < -793) {
                                                            var v30 bool = true
                                                            _ = v30
                                                        } else {
                                                            if true {
                                                                if (383 >= 0) {
                                                                    if (!false) {
                                                                        fmt.Println(fi28, fi28, fi28)
                                                                    }
                                                                    if true {
                                                                        v1 = v1
                                                                        v5[("a" + "")] = 4294967295
                                                                        if false {
                                                                            v2 = append(v2, v1)
                                                                            var v31 uint16 = 32767
                                                                            _ = v31
                                                                        }
                                                                    } else {
                                                                        v32 := GenS2{F0: v19, F1: 901, F2: 1}
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                        _ = v32
                                                                    }
                                                                }
                                                                if len(v2) > 0 {
                                                                    v2[0] = (((v1 + "\n\t") + v1) + (("foo bar" + v1) + (("\n\t" + v1) + v1)))
                                                                }
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            }
                                                            switch v1 {
                                                            case "a":
                                                                var v33 uint8 = v14
                                                                _ = v33
                                                            default:
                                                                var v34 int32 = 2147483647
                                                                _ = v34
                                                            }
                                                            v4 += (v4 - 0.0)
                                                        }
                                                    }
                                                default:
                                                    fmt.Println(v4, v4)
                                                }
                                                if true {
                                                    var v35 int8 = -128
                                                    for fi36 := 0; fi36 < 2; fi36++ {
                                                        v3[2] = (fi28 % ((4294967295) | 1))
                                                        switch 128 {
                                                        case -9223372036854775808:
                                                            if true {
                                                                v4 -= (v4 / 1.0)
                                                                fmt.Println(v4, fi36)
                                                            } else {
                                                                v19 = v19
                                                            }
                                                            v4 += v4
                                                            v14 = 0
                                                        default:
                                                            if (!(false || false)) {
                                                                var v37 int32 = -1
                                                                v38 := v3[1]
                                                                switch ("foo bar" + v1) {
                                                                case "":
                                                                    var v39 int8 = v35
                                                                    _ = v39
                                                                case "hello":
                                                                    var v40 string = ((v1 + "foo bar") + "foo bar")
                                                                    _ = v40
                                                                case "a":
                                                                    var v41 uint8 = v14
                                                                    _ = v41
                                                                default:
                                                                    v19 += v19
                                                                }
                                                                v1 = ("hello" + "a")
                                                                _ = v37
                                                                _ = v38
                                                            } else {
                                                                if (230 == 2) {
                                                                    v42 := GenS4{F0: 835, F1: v1}
                                                                    v14 += v14
                                                                    fmt.Println(v35)
                                                                    _ = v42
                                                                }
                                                                v3[818] = (fi36 * (fi36 >> ((65535) & 63)))
                                                                switch 42 {
                                                                case 128:
                                                                    var v43 int8 = v35
                                                                    _ = v43
                                                                case 956:
                                                                    if true {
                                                                        var v44 uint = 18446744073709551615
                                                                        v19 += 0
                                                                        v45 := GenS2{F0: 0, F1: 1, F2: 18446744073709551615}
                                                                        if true {
                                                                            v35 = v35
                                                                            var v46 int64 = 0
                                                                            v46 = v46
                                                                            _ = v46
                                                                        }
                                                                        _ = v44
                                                                        _ = v45
                                                                    } else {
                                                                        var v47 uint64 = 707
                                                                        _ = v47
                                                                    }
                                                                case -1:
                                                                    switch 398 {
                                                                    case 891:
                                                                        _ = len(v5)
                                                                    case 32767:
                                                                        if (3.14159 == (0.1 * ((0.0 + ((v4 / v4) - (float32((1e10 / (1.5 + 1e10))) + v4))) / v4))) {
                                                                            v19 = 42
                                                                            v14 = (v14 ^ v14)
                                                                            fmt.Println(v19, fi28, fi36)
                                                                        }
                                                                    }
                                                                }
                                                                var v48 int64 = 0
                                                                v49 := GenS7{F0: v35, F1: int(v19)}
                                                                _ = v48
                                                                _ = v49
                                                            }
                                                        }
                                                        _ = fi36
                                                    }
                                                    _ = v35
                                                }
                                                _ = fi28
                                            }
                                        }
                                    case "a":
                                        delete(v5, (v1 + ("" + "\n\t")))
                                    default:
                                        switch v1 {
                                        case "foo bar":
                                            if false {
                                                v1 = ("\n\t" + (v1 + (((v1 + "a") + v1) + v1)))
                                            } else {
                                                v4 = 3.14159
                                                v1 = "foo bar"
                                                var v50 uint = 126
                                                _ = v50
                                            }
                                        case "a":
                                            var v51 int = -9223372036854775808
                                            _ = v51
                                        case "\n\t":
                                            if true {
                                                for fi52 := 0; fi52 < 14; fi52++ {
                                                    var v53 string = v1
                                                    if len(v2) > 0 {
                                                        v2[0] = "\n\t"
                                                    }
                                                    _ = fi52
                                                    _ = v53
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = (((v1 + (("a" + ((v1 + (v1 + ("" + v1))) + v1)) + "")) + v1) + "")
                                                }
                                                switch v1 {
                                                case "foo bar":
                                                    var v54 uint16 = 2
                                                    _ = v54
                                                case "hello":
                                                    for fi55 := 0; fi55 < 3; fi55++ {
                                                        var v56 float32 = (v4 + v4)
                                                        _ = fi55
                                                        _ = v56
                                                    }
                                                }
                                            }
                                        default:
                                            v4 = 3.14159
                                        }
                                    }
                                }
                                v57 := GenS8{F0: (!true)}
                                _ = v57
                            } else {
                                for fi58 := 0; fi58 < 15; fi58++ {
                                    if (!true) {
                                        var v59 uint32 = 2147483647
                                        if len(v2) > 0 {
                                            v2[0] = (v1 + v1)
                                        }
                                        if (v4 <= v4) {
                                            v14 = 255
                                        } else {
                                            if true {
                                                if false {
                                                    var v60 int32 = -2147483648
                                                    for fi61 := 0; fi61 < 20; fi61++ {
                                                        if (123 <= -19) {
                                                            v62 := v5[v1]
                                                            _ = v62
                                                        }
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        if true {
                                                            v59 -= 0
                                                        } else {
                                                            var v63 string = "\n\t"
                                                            var v64 uint32 = 1
                                                            delete(v3, (^(v59)))
                                                            _ = v63
                                                            _ = v64
                                                        }
                                                        _ = fi61
                                                    }
                                                    var v65 uint = 255
                                                    _ = v60
                                                    _ = v65
                                                }
                                            } else {
                                                for fi66 := 0; fi66 < 20; fi66++ {
                                                    switch 0 {
                                                    case 9223372036854775807:
                                                        var v67 uint32 = 4294967295
                                                        _ = v67
                                                    }
                                                    v68 := GenS9{F0: (fi58 >> ((255) & 63)), F1: ((false || ((true || (0.1 < float64(float32(1e10)))) && false)) || true)}
                                                    _ = fi66
                                                    _ = v68
                                                }
                                                fmt.Println(v1, v4, v59)
                                                v2 = append(v2, "foo bar")
                                            }
                                        }
                                        _ = v59
                                    }
                                    fmt.Println(fi58)
                                    v69 := GenS10{F0: ((42 == -691) && true), F1: ("a" + ((("foo bar" + "hello") + "") + v1))}
                                    _ = fi58
                                    _ = v69
                                }
                                var v70 uint16 = 84
                                _ = v70
                            }
                        case 1:
                            var v71 string = ("a" + (v1 + "\n\t"))
                            _ = v71
                        default:
                            for fi72 := 0; fi72 < 18; fi72++ {
                                fmt.Println(v14, v4, fi72)
                                v73 := GenS7{F0: genFn0(), F1: 4294967295}
                                _ = fi72
                                _ = v73
                            }
                        }
                    }
                }
            }
        case "":
            for fi74 := 0; fi74 < 3; fi74++ {
                if false {
                    fmt.Println(v14)
                    fmt.Println(fi74)
                } else {
                    if false {
                        if false {
                            var v75 int16 = 1
                            var v76 int8 = 0
                            var v77 string = "hello"
                            _ = v75
                            _ = v76
                            _ = v77
                        } else {
                            if len(v2) > 0 {
                                v2[0] = v1
                            }
                        }
                        fmt.Println(v14, v1, fi74)
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    if (463 == 1) {
                        fmt.Println(v19, v4)
                    }
                    var v78 int = (-(fi74))
                    _ = v78
                }
                var v79 int32 = 42
                if false {
                    var v80 uint64 = 540
                    var v81 float32 = v4
                    if true {
                        var v82 uint = 318
                        _ = v82
                    }
                    _ = v80
                    _ = v81
                }
                _ = fi74
                _ = v79
            }
        case "\n\t":
            delete(v5, (((("\n\t" + "a") + ("a" + v1)) + "a") + "foo bar"))
        default:
            fmt.Println(v14)
        }
        var v83 uint8 = v14
        _ = v83
    }
    if ((2 > 42) || ((-128 <= -81) && true)) {
        var v84 int64 = 42
        v14 += (v14 * v14)
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
