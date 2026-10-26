package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 bool; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint16; F2 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string; F1 uint32 }
func genFn0(p0 uint16, p1 uint16) uint32 {
    p0 = (^(p0))
    v5 := []int{-9223372036854775808, -9223372036854775808}
    var v6 string = "foo bar"
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 string = ("" + ("foo bar" + "foo bar"))
    var v2 int16 = 0
    var v3 float64 = -2.25
    var v4 int32 = 1
    for fi7 := 0; fi7 < 5; fi7++ {
        fmt.Println(fi7, v4, v1)
        if (v3 <= float64(0.0)) {
            fmt.Println(v3)
            for fi8 := 0; fi8 < 11; fi8++ {
                v1 = (v1 + (v1 + (("foo bar" + (((v1 + "\n\t") + v1) + "")) + "")))
                _ = fi8
            }
        } else {
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 int8 = 42
                v11 := GenS1{F0: 32767}
                fmt.Println(fi9, v2, v1)
                v12 := []int{fi9, 582}
                var v13 int32 = v4
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
                _ = v13
            }
            v14 := []int{(^(fi7)), -2147483648, (fi7 >> ((172) & 63))}
            v4 += v4
            _ = v14
        }
        _ = fi7
    }
    v15 := []int{}
    if (104 == 1) {
        switch v1 {
        case "foo bar":
            var v16 uint = 18446744073709551615
            _ = v16
        case "":
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 bool = true
                _ = fi17
                _ = v18
            }
        }
        var v19 uint32 = 0
        v4 += v4
        _ = v19
    } else {
        v2 += (v2 - 32767)
    }
    fmt.Println(v2)
    if len(v15) > 0 {
        v15[0] = -854
    }
    var v20 uint = 297
    v3 += v3
    switch (((v1 + v1) + ((((v1 + (("foo bar" + (((v1 + v1) + (v1 + ((v1 + v1) + v1))) + ((v1 + v1) + v1))) + "")) + ("hello" + (v1 + (v1 + "\n\t")))) + ((v1 + v1) + "\n\t")) + (v1 + v1))) + ((v1 + v1) + "")) {
    case "foo bar":
        v2 += v2
    case "a":
        v1 = v1
    case "hello":
        if len(v15) > 0 {
            v15[0] = -151
        }
    default:
        if len(v15) > 0 {
            v15[0] = 0
        }
    }
    if true {
        if (v20 > v20) {
            for fi21 := 0; fi21 < 13; fi21++ {
                switch (fi21 * fi21) {
                case 343:
                    v22 := GenS1{F0: 255}
                    _ = v22
                case 32767:
                    v3 = v3
                default:
                    switch (v1 + (((v1 + v1) + v1) + ((v1 + v1) + "a"))) {
                    case "a":
                        if len(v15) > 0 {
                            v15[0] = 128
                        }
                    default:
                        fmt.Println(v2)
                    }
                }
                fmt.Println(v1, v3, v1)
                if ((1 <= 200) && false) {
                    v1 = "foo bar"
                    for fi23 := 0; fi23 < 11; fi23++ {
                        var v24 bool = false
                        v24 = (v24 || (2 >= 0))
                        _ = fi23
                        _ = v24
                    }
                    switch (("a" + v1) + v1) {
                    case "hello":
                        if false {
                            v2 += -251
                            for fi25 := 0; fi25 < 12; fi25++ {
                                v2 = v2
                                var v26 int8 = -1
                                var v27 uint64 = 1
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        } else {
                            v15 = append(v15, 4294967295)
                            if false {
                                v20 = (v20 % ((v20) | 1))
                            } else {
                                var v28 uint64 = 2
                                v28 -= (v28 / ((2) | 1))
                                _ = v28
                            }
                        }
                    case "foo bar":
                        v2 += v2
                    default:
                        switch (fi21 ^ (fi21 ^ -13)) {
                        case -2147483648:
                            if (65535 < 1) {
                                if (true || (float64(0.1) != (v3 + v3))) {
                                    v29 := GenS1{F0: 257}
                                    var v30 bool = false
                                    _ = v29
                                    _ = v30
                                }
                                if (207 != 1) {
                                    for fi31 := 0; fi31 < 12; fi31++ {
                                        v1 = v1
                                        fmt.Println(v20, v20, v3)
                                        fmt.Println(v2, v2, v1)
                                        if (!true) {
                                            v3 = ((v3 * float64((((0.0 + ((float32(v3) + -2.25) / (3.14159 * 3.14159))) * float32((float64(1.5) - float64(((1.5 - 3.14159) * (float32(v3) + float32(1.5))))))) * 1.5))) * float64(((float32(v3) - ((((-2.25 / 1.5) + 0.1) - 0.1) / 3.14159)) * ((-2.25 * float32(v3)) / 1.0))))
                                            var v32 string = ("" + "\n\t")
                                            if (false && false) {
                                                v33 := GenS2{F0: (^(v20)), F1: true, F2: 42}
                                                fmt.Println(v20, v3)
                                                _ = len(v15)
                                                _ = v33
                                            } else {
                                                v34 := GenS3{F0: 2, F1: v20, F2: 18446744073709551615}
                                                v35 := GenS3{F0: 65535, F1: (v20 + (v20 << ((v20) & 63))), F2: 0}
                                                _ = v34
                                                _ = v35
                                            }
                                            if len(v15) > 0 {
                                                v15[0] = (fi21 + fi31)
                                            }
                                            if ((!true) || true) {
                                                var v36 uint16 = 1
                                                _ = v36
                                            }
                                            _ = v32
                                        }
                                        _ = fi31
                                    }
                                    v37 := GenS4{F0: true, F1: 0, F2: ((fi21 ^ 32767) != 101)}
                                    fmt.Println(fi21, v20)
                                    _ = v37
                                }
                                switch v1 {
                                case "hello":
                                    for fi38 := 0; fi38 < 1; fi38++ {
                                        v39 := GenS2{F0: 255, F1: false, F2: 42}
                                        _ = fi38
                                        _ = v39
                                    }
                                case "a":
                                    for fi40 := 0; fi40 < 20; fi40++ {
                                        v2 += v2
                                        fmt.Println(fi40)
                                        v41 := GenS5{F0: true}
                                        _ = fi40
                                        _ = v41
                                    }
                                case "\n\t":
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                default:
                                    switch (((("" + "a") + v1) + v1) + "\n\t") {
                                    case "a":
                                        if len(v15) >= 2 {
                                            v15 = v15[1:len(v15)]
                                        }
                                    case "hello":
                                        if len(v15) > 0 {
                                            v15[0] = (fi21 / ((fi21) | 1))
                                        }
                                    case "\n\t":
                                        var v42 int32 = 42
                                        _ = v42
                                    }
                                }
                            }
                        case 127:
                            var v43 int32 = (^(v4))
                            _ = v43
                        case 42:
                            if true {
                                if (!(fi21 != fi21)) {
                                    if (0.1 > 3.14159) {
                                        fmt.Println(v1, v1)
                                        v3 -= 1.5
                                        fmt.Println(v3)
                                    } else {
                                        v44 := GenS5{F0: (0.0 == (((v3 + v3) - (v3 * float64((3.14159 / 1.0)))) / (1.5 / v3)))}
                                        var v45 int32 = v4
                                        _ = v44
                                        _ = v45
                                    }
                                    for fi46 := 0; fi46 < 4; fi46++ {
                                        if true {
                                            v3 = -2.25
                                            v47 := GenS3{F0: 379, F1: v20, F2: 42}
                                            v1 = "hello"
                                            _ = v47
                                        }
                                        var v48 uint16 = 32767
                                        _ = fi46
                                        _ = v48
                                    }
                                }
                                if ((769 <= 255) && true) {
                                    switch fi21 {
                                    case 237:
                                        fmt.Println(v2, v2, v20)
                                    }
                                    var v49 int32 = (v4 + -2147483648)
                                    _ = v49
                                }
                                switch ("" + (v1 + "\n\t")) {
                                case "foo bar":
                                    v20 += 18446744073709551615
                                case "":
                                    switch (v1 + (("hello" + (v1 + v1)) + "")) {
                                    case "a":
                                        v50 := GenS6{F0: (-(v4))}
                                        _ = v50
                                    case "foo bar":
                                        var v51 float32 = (-2.25 - ((0.1 + 0.0) / 0.1))
                                        _ = v51
                                    }
                                case "a":
                                    switch 127 {
                                    case 128:
                                        v15 = append(v15, (fi21 * fi21))
                                    }
                                }
                            }
                        default:
                            switch (fi21 >> (((v20 ^ (v20 & v20))) & 63)) {
                            case 632:
                                v52 := GenS1{F0: 686}
                                _ = v52
                            case -287:
                                v53 := GenS3{F0: 2, F1: v20, F2: 42}
                                _ = v53
                            case 32768:
                                if ((v4 >= v4) && (2 <= 1)) {
                                    var v54 bool = false
                                    _ = v54
                                } else {
                                    if ((v2 + -404) >= v2) {
                                        var v55 uint16 = 42
                                        fmt.Println(fi21)
                                        _ = v55
                                    } else {
                                        for fi56 := 0; fi56 < 4; fi56++ {
                                            var v57 float32 = (0.0 + float32((v3 + 3.14159)))
                                            switch (fi56 * int(v2)) {
                                            case -606:
                                                switch (v1 + v1) {
                                                case "a":
                                                    switch 475 {
                                                    case 42:
                                                        v1 = v1
                                                    }
                                                case "":
                                                    v58 := GenS4{F0: false, F1: uint16(v20), F2: (0 < 188)}
                                                    _ = v58
                                                default:
                                                    switch 4294967295 {
                                                    case -464:
                                                        v20 += 2
                                                    default:
                                                        v59 := GenS7{F0: v1, F1: 120}
                                                        _ = v59
                                                    }
                                                }
                                            default:
                                                if false {
                                                    v4 = v4
                                                } else {
                                                    var v60 uint = (v20 | v20)
                                                    var v61 bool = (1 >= 42)
                                                    fmt.Println(v20, fi56, v57)
                                                    _ = v60
                                                    _ = v61
                                                }
                                            }
                                            _ = fi56
                                            _ = v57
                                        }
                                    }
                                    v2 = (v2 ^ 32767)
                                }
                            default:
                                v3 += ((0.0 / v3) / 1.5)
                            }
                        }
                    }
                }
                _ = fi21
            }
        }
        v2 -= (v2 & v2)
        var v62 uint64 = 18446744073709551615
        _ = v62
    } else {
        fmt.Println(v2, v20)
        for fi63 := 0; fi63 < 3; fi63++ {
            _ = len(v15)
            if ("foo bar" != (v1 + (v1 + ("a" + v1)))) {
                if true {
                    if (539 <= 2) {
                        v3 -= 0.0
                        var v64 bool = true
                        _ = v64
                    }
                    v1 = "\n\t"
                }
            }
            _ = fi63
        }
    }
    if (1 < 9223372036854775807) {
        v65 := GenS7{F0: v1, F1: genFn0(32767, 0)}
        _ = v65
    }
    v15 = append(v15, 555)
    var v66 float64 = v3
    v15 = append(v15, 731)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v66
}
