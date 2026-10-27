package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int64; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 uint; F2 int16 }
type GenS7 struct { F0 int16; F1 int16 }
func genFn0(p0 uint, p1 bool) int16 {
    p0 = p0
    v6 := GenS1{F0: 42, F1: 42, F2: 65535}
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    return -32768
}

func main() {
    var v1 uint = 311
    v2 := make(map[string]bool)
    var v3 float32 = 0.1
    var v4 int16 = 32767
    v5 := make(map[string]bool)
    v7 := GenS1{F0: -168, F1: 54, F2: 1}
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            fmt.Println(fi8, fi8)
            for fi9 := 0; fi9 < 3; fi9++ {
                switch "foo bar" {
                case "a":
                    fmt.Println(v3)
                default:
                    switch (fi9 * (^(fi8))) {
                    case 2147483647:
                        var v10 string = "hello"
                        _ = v10
                    case 0:
                        v11 := GenS1{F0: 9223372036854775807, F1: 42, F2: 198}
                        _ = v11
                    }
                    var v12 int64 = -9223372036854775808
                    v12 = -1
                    _ = v12
                }
                v13 := GenS1{F0: -1, F1: 127, F2: 1}
                v1 -= 2
                _ = fi9
                _ = v13
            }
        }
        v1 = v1
        _ = fi8
    }
    v3 = float32(0.1)
    v14 := []int{-2147483648, 2147483647, 127}
    fmt.Println(v3, v4, v4)
    v15 := GenS1{F0: 9223372036854775807, F1: -43, F2: 372}
    if ((!true) && ("a" <= (("" + "foo bar") + ""))) {
        _ = len(v14)
        var v16 uint = v1
        if (0 == 983) {
            v3 = (0.1 * ((v3 / 1.5) / 3.14159))
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 float32 = ((v3 / (((v3 / v3) - (v3 + 0.0)) / 1e10)) / float32((-2.25 - ((0.1 - (0.0 - (1.5 - (((3.14159 / 3.14159) + ((1.5 / (0.1 + 1e10)) / float64((v3 - 0.1)))) + -2.25)))) * (1.5 + (0.1 / 3.14159))))))
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 9; fi19++ {
                var v20 uint64 = 18446744073709551615
                _ = fi19
                _ = v20
            }
        } else {
            _ = len(v5)
        }
        _ = v16
    } else {
        v3 -= (v3 + 0.0)
        for fi21 := 0; fi21 < 16; fi21++ {
            v4 = (-(v4))
            v3 = v3
            if false {
                v22 := GenS1{F0: 0, F1: 100, F2: 635}
                v23 := GenS1{F0: 9223372036854775807, F1: 53, F2: 1}
                _ = v22
                _ = v23
            } else {
                switch (fi21 ^ 4294967295) {
                case 0:
                    v3 = ((v3 * (1e10 - v3)) / ((v3 - 3.14159) / (((v3 / 1.5) / v3) / v3)))
                case -9223372036854775808:
                    v24 := GenS1{F0: 561, F1: 127, F2: 346}
                    _ = v24
                }
                v5[((("\n\t" + "\n\t") + "") + "a")] = (true || false)
                _ = len(v5)
            }
            _ = fi21
        }
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        if true {
            v4 = genFn0(v1, (!(false && (0.1 > 1.5))))
            var v26 uint16 = 0
            v5["hello"] = true
            _ = v26
        } else {
            if false {
                v27 := GenS1{F0: 562, F1: -11, F2: 255}
                if false {
                    for fi28 := 0; fi28 < 18; fi28++ {
                        v2["\n\t"] = (false || (!(95 <= 42)))
                        switch fi25 {
                        case 9223372036854775807:
                            if (v3 < ((float32(1.5) * (v3 / (((v3 / (0.1 + v3)) - (v3 + ((((v3 / (0.0 / ((v3 * (v3 * v3)) * 3.14159))) / (1.5 - 3.14159)) * (1.5 + ((((v3 - (0.1 * ((1e10 - 0.0) / v3))) - v3) + 0.0) - (float32(float64(((v3 * v3) - 3.14159))) * 3.14159)))) - v3))) / (v3 * ((float32(1e10) * 3.14159) + (v3 - 0.1)))))) / (v3 / (1e10 / -2.25)))) {
                                fmt.Println(v1, v1, v3)
                                v3 -= 1e10
                                v1 = (v1 + (v1 << (((^(v1))) & 63)))
                            } else {
                                var v29 bool = ((("hello" + "hello") < "hello") && (!true))
                                v30 := GenS1{F0: 42, F1: -30, F2: 2}
                                _ = v29
                                _ = v30
                            }
                        case -1:
                            var v31 uint8 = 1
                            _ = v31
                        case -419:
                            var v32 int8 = -1
                            _ = v32
                        default:
                            if len(v14) > 0 {
                                v14[0] = 183
                            }
                        }
                        _ = fi28
                    }
                    v33 := GenS1{F0: 1, F1: 127, F2: 42}
                    if (0 == 1) {
                        var v34 float32 = float32(((float64(0.1) / 1.0) * 1e10))
                        var v35 int16 = v4
                        var v36 uint = 4294967295
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    _ = v33
                } else {
                    v3 -= float32(1e10)
                    fmt.Println(v1, v1)
                    if (2 > 4294967295) {
                        var v37 string = ""
                        _ = v37
                    }
                }
                v38 := GenS1{F0: 9223372036854775807, F1: -62, F2: 2}
                _ = v27
                _ = v38
            } else {
                var v39 bool = false
                var v40 string = ("a" + ("hello" + ""))
                _ = v39
                _ = v40
            }
            v1 = 65535
            for fi41 := 0; fi41 < 14; fi41++ {
                switch fi41 {
                case -328:
                    var v42 uint64 = 1
                    _ = v42
                case 128:
                    if (!true) {
                        v43 := GenS2{F0: 1, F1: 0}
                        _ = len(v14)
                        switch (-(fi41)) {
                        case 42:
                            delete(v2, "foo bar")
                        case -9223372036854775808:
                            v44 := GenS3{F0: -1, F1: -1, F2: 65535}
                            _ = v44
                        case -2147483648:
                            fmt.Println(v1, v3, v4)
                        }
                        fmt.Println(v3, v4)
                        _ = v43
                    } else {
                        v45 := GenS4{F0: 2, F1: 2147483648}
                        v3 -= (((v3 / 3.14159) / v3) * float32((-2.25 - 1.5)))
                        var v46 uint8 = 128
                        _ = v45
                        _ = v46
                    }
                case -1:
                    fmt.Println(v4)
                default:
                    v3 -= -2.25
                    if true {
                        var v47 uint8 = 1
                        var v48 string = ((("\n\t" + ("" + ("foo bar" + "foo bar"))) + "") + "foo bar")
                        if false {
                            if (((38 <= 42) && true) || (v1 != 65535)) {
                                var v49 bool = (210 >= -1)
                                var v50 int16 = (v4 & v4)
                                if false {
                                    var v51 uint8 = uint8(v1)
                                    _ = v51
                                } else {
                                    if true {
                                        if len(v14) >= 2 {
                                            v14 = v14[1:len(v14)]
                                        }
                                    }
                                    switch fi25 {
                                    case 611:
                                        switch (^(fi25)) {
                                        case 1:
                                            fmt.Println(v50, v1)
                                        }
                                    }
                                }
                                _ = v49
                                _ = v50
                            }
                            switch v48 {
                            case "foo bar":
                                if (65535 > 2147483648) {
                                    v1 = 18446744073709551615
                                    v52 := GenS1{F0: 716, F1: 127, F2: 1}
                                    v53 := GenS5{F0: 0}
                                    v54 := GenS3{F0: 597, F1: 42, F2: 2147483647}
                                    if (152 >= -2147483648) {
                                        v55 := GenS2{F0: 42, F1: 60}
                                        if (v4 != -505) {
                                            fmt.Println(fi25, v1)
                                        } else {
                                            var v56 int16 = (-(v4))
                                            v48 = "\n\t"
                                            if false {
                                                fmt.Println(v47)
                                                v1 = v1
                                                v57 := GenS3{F0: -9223372036854775808, F1: -128, F2: 2147483648}
                                                v1 += 18446744073709551615
                                                var v58 int16 = v56
                                                _ = v57
                                                _ = v58
                                            } else {
                                                fmt.Println(fi25, v48, v48)
                                            }
                                            _ = v56
                                        }
                                        v59 := GenS6{F0: -1, F1: 869, F2: (v4 - v4)}
                                        var v60 float64 = (-2.25 + 1.5)
                                        _ = v55
                                        _ = v59
                                        _ = v60
                                    } else {
                                        v47 = uint8(v1)
                                        if true {
                                            v5[""] = ((-2.25 / (float64(float32(1e10)) * 3.14159)) != 0.1)
                                            v1 -= v1
                                            fmt.Println(v1)
                                        } else {
                                            var v61 int64 = 9223372036854775807
                                            var v62 uint64 = 369
                                            var v63 uint32 = 2147483648
                                            _ = v61
                                            _ = v62
                                            _ = v63
                                        }
                                    }
                                    _ = v52
                                    _ = v53
                                    _ = v54
                                } else {
                                    switch v48 {
                                    case "foo bar":
                                        v64 := GenS5{F0: 255}
                                        v3 -= 1e10
                                        v1 -= (v1 & (v1 & (^(v1))))
                                        _ = v64
                                    }
                                }
                            case "\n\t":
                                var v65 int32 = 1
                                _ = v65
                            case "":
                                v47 = v47
                            }
                        } else {
                            _ = len(v14)
                            v3 = v3
                            fmt.Println(v48)
                        }
                        var v66 uint16 = 65535
                        fmt.Println(fi41, fi25)
                        _ = v47
                        _ = v48
                        _ = v66
                    }
                }
                _ = fi41
            }
        }
        v1 = v1
        v67 := GenS2{F0: 1, F1: 42}
        _ = fi25
        _ = v67
    }
    for fi68 := 0; fi68 < 5; fi68++ {
        v1 += v1
        _ = len(v5)
        var v69 float32 = v3
        _ = fi68
        _ = v69
    }
    if (true && (128 <= -2147483648)) {
        if ((0.0 * 0.1) == 1.5) {
            for fi70 := 0; fi70 < 18; fi70++ {
                v4 += v4
                if (!true) {
                    var v71 int16 = genFn0(18446744073709551615, (true || ((!(((!(((1e10 * (-2.25 * float64(v3))) < (0.0 + 1.5)) || ((-1 <= 0) || true))) || false) && false)) || ((float64(1.5) - -2.25) <= 0.0))))
                    var v72 string = ("foo bar" + ("" + "foo bar"))
                    v72 = ("a" + v72)
                    _ = v71
                    _ = v72
                }
                _ = fi70
            }
        } else {
            v3 -= v3
            v73 := GenS5{F0: 65535}
            _ = v73
        }
        if true {
            var v74 string = (("" + "hello") + "")
            v1 = (v1 << (((^(v1))) & 63))
            v3 = -2.25
            _ = v74
        } else {
            _ = len(v14)
            for fi75 := 0; fi75 < 9; fi75++ {
                fmt.Println(v4)
                for fi76 := 0; fi76 < 17; fi76++ {
                    v3 = v3
                    switch "hello" {
                    case "":
                        var v77 int32 = 1
                        _ = v77
                    }
                    if true {
                        var v78 int8 = -1
                        _ = v78
                    }
                    _ = fi76
                }
                _ = fi75
            }
        }
        for fi79 := 0; fi79 < 17; fi79++ {
            fmt.Println(v4)
            var v80 uint16 = 0
            v81 := GenS1{F0: -76, F1: -128, F2: (v80 % ((2) | 1))}
            _ = fi79
            _ = v80
            _ = v81
        }
    }
    switch -1 {
    case 1:
        v82 := GenS7{F0: 42, F1: (v4 + 80)}
        _ = v82
    case 130:
        v3 = v3
    }
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
