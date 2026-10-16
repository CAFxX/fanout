package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint64; F2 uint }
type GenS5 struct { F0 int64; F1 int; F2 int32 }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 int; F1 int16; F2 int }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 uint64; F1 int8; F2 int }
func genFn0(p0 uint16, p1 int32, p2 uint32) int8 {
    p2 += 2147483648
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    return -42
}

func main() {
    v1 := []uint32{42, 2, 204, 499}
    var v2 float32 = 0.0
    v3 := []int{1, 4294967295}
    var v4 int64 = -372
    v5 := make(map[string]string)
    v6 := make(map[string]string)
    v4 += 42
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS1{F0: 1}
        if false {
            fmt.Println(fi7, v4, v4)
            v4 = v4
        }
        _ = fi7
        _ = v8
    }
    v2 = (v2 * v2)
    if (127 > 42) {
        fmt.Println(v4)
    }
    v9 := GenS2{F0: (-78 > 127), F1: 65535, F2: 202}
    var v10 float64 = (3.14159 / ((((0.0 / float64(v2)) - 3.14159) / 1e10) - (float64(v2) + 0.0)))
    v11 := GenS2{F0: true, F1: 32767, F2: 42}
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS2{F0: (!(!(76 <= 0))), F1: 65535, F2: 128}
        v4 += 9223372036854775807
        _ = fi12
        _ = v13
    }
    if true {
        for fi14 := 0; fi14 < 13; fi14++ {
            if (-367 > -124) {
                v4 = (v4 / (((^(v4))) | 1))
                if false {
                    if (42 > 885) {
                        delete(v5, (("" + "a") + ""))
                        switch (fi14 + fi14) {
                        case 446:
                            v4 = (v4 << ((326) & 63))
                        }
                        v10 += ((v10 / v10) + v10)
                    } else {
                        var v15 uint32 = 591
                        for fi16 := 0; fi16 < 7; fi16++ {
                            if (!false) {
                                v15 -= 0
                                if true {
                                    var v17 uint16 = 1
                                    var v18 uint = 255
                                    _ = v17
                                    _ = v18
                                }
                            }
                            fmt.Println(v15, v4, v2)
                            v6["\n\t"] = "foo bar"
                            _ = fi16
                        }
                        var v19 int8 = 100
                        _ = v15
                        _ = v19
                    }
                    if false {
                        var v20 int32 = 54
                        _ = v20
                    } else {
                        if (!(!(((((float64(v2) + 0.0) + ((-2.25 + v10) * ((v10 / 1.5) / -2.25))) * -2.25) * 3.14159) == 3.14159))) {
                            if true {
                                delete(v6, "hello")
                                if len(v3) > 0 {
                                    v3[0] = 766
                                }
                            }
                            switch ((("" + "foo bar") + "a") + ((("hello" + "") + "\n\t") + "hello")) {
                            case "foo bar":
                                v5["foo bar"] = "a"
                            case "":
                                for fi21 := 0; fi21 < 14; fi21++ {
                                    v5[""] = ("a" + "hello")
                                    var v22 int16 = 42
                                    v23 := GenS3{F0: 42, F1: 42}
                                    _ = fi21
                                    _ = v22
                                    _ = v23
                                }
                            case "a":
                                if false {
                                    v24 := v6[""]
                                    v2 = v2
                                    for fi25 := 0; fi25 < 3; fi25++ {
                                        if ((v4 + (v4 ^ v4)) > v4) {
                                            v24 = ((v24 + v24) + "\n\t")
                                            v4 = (-(v4))
                                            v26 := GenS4{F0: 32767, F1: 42, F2: 0}
                                            _ = v26
                                        }
                                        v24 = "\n\t"
                                        _ = fi25
                                    }
                                    _ = v24
                                }
                            }
                        }
                    }
                }
            }
            _ = fi14
        }
        var v27 float64 = ((((((v10 + float64(float32(0.1))) * ((1e10 + v10) - (v10 - 1e10))) + v10) + 1.5) / (float64((v2 - v2)) / (float64(1.5) - -2.25))) + v10)
        if (509 == 42) {
            v28 := GenS5{F0: v4, F1: 476, F2: 328}
            if (!((!((((v10 + ((v10 - ((v10 * -2.25) / v10)) / ((v27 - float64(v2)) * v10))) + -2.25) != v27) && false)) && ((((0.1 * v27) / v10) + v27) <= ((((0.1 / 1e10) + 3.14159) / v10) * (v27 / v27))))) {
                var v29 bool = true
                _ = v29
            } else {
                var v30 uint64 = 18446744073709551615
                if false {
                    var v31 float32 = 1.5
                    _ = v31
                } else {
                    var v32 bool = (42 > 1)
                    _ = v32
                }
                for fi33 := 0; fi33 < 7; fi33++ {
                    v4 -= (v4 % ((int64(fi33)) | 1))
                    v34 := GenS6{F0: 165, F1: 313}
                    for fi35 := 0; fi35 < 15; fi35++ {
                        fmt.Println(fi33)
                        _ = fi35
                    }
                    _ = fi33
                    _ = v34
                }
                _ = v30
            }
            var v36 float64 = v27
            _ = v28
            _ = v36
        }
        _ = v27
    }
    var v37 string = "a"
    if ((true || true) || false) {
        var v38 uint64 = 2
        _ = v38
    } else {
        if true {
            v39 := GenS1{F0: 2}
            v40 := GenS3{F0: 22, F1: 1}
            fmt.Println(v2, v4)
            _ = v39
            _ = v40
        } else {
            if (65535 != 42) {
                var v41 uint8 = 1
                _ = v41
            } else {
                v37 = "\n\t"
                v42 := GenS7{F0: -9223372036854775808, F1: 1000, F2: 560}
                if true {
                    var v43 uint32 = 0
                    var v44 float64 = (v10 + (v10 / 1.0))
                    var v45 int8 = 0
                    _ = v43
                    _ = v44
                    _ = v45
                } else {
                    var v46 bool = ((!true) && true)
                    var v47 uint16 = 32767
                    if false {
                        for fi48 := 0; fi48 < 9; fi48++ {
                            var v49 int = (-(fi48))
                            _ = fi48
                            _ = v49
                        }
                        v50 := GenS4{F0: (v47 + v47), F1: 910, F2: 4294967295}
                        for fi51 := 0; fi51 < 3; fi51++ {
                            fmt.Println(fi51, v10, v46)
                            _ = len(v6)
                            _ = fi51
                        }
                        _ = v50
                    }
                    _ = v46
                    _ = v47
                }
                _ = v42
            }
        }
        v52 := GenS8{F0: v4, F1: 1}
        _ = v52
    }
    var v53 float32 = (float32(v10) * v2)
    v54 := GenS3{F0: 255, F1: 1}
    if true {
        for fi55 := 0; fi55 < 17; fi55++ {
            for fi56 := 0; fi56 < 12; fi56++ {
                var v57 uint32 = 1
                v2 = v2
                switch (v37 + (v37 + v37)) {
                case "a":
                    switch 1 {
                    case 9223372036854775807:
                        switch v37 {
                        case "a":
                            fmt.Println(fi56, v2, v57)
                        }
                    }
                    if true {
                        fmt.Println(v4, v57)
                        var v58 int16 = 0
                        _ = v58
                    } else {
                        v3 = append(v3, (fi56 & 556))
                        var v59 uint16 = 589
                        v57 -= v57
                        _ = v59
                    }
                    fmt.Println(fi55)
                    v60 := GenS9{F0: 415, F1: 0, F2: 4294967295}
                    if false {
                        v53 -= (v53 * v2)
                        if false {
                            if ((false && true) && (true || ((v10 - v10) > 0.0))) {
                                if (18446744073709551615 < 18446744073709551615) {
                                    if len(v3) > 0 {
                                        v3[0] = (-(fi55))
                                    }
                                    v10 = float64(v53)
                                    v61 := GenS5{F0: (v4 * v4), F1: fi55, F2: -1}
                                    v2 += v2
                                    _ = v61
                                } else {
                                    switch fi56 {
                                    case -1:
                                        v62 := GenS7{F0: fi55, F1: int16(fi56), F2: fi55}
                                        _ = v62
                                    case 1:
                                        v63 := GenS2{F0: ((1.5 + float32((v10 + 1e10))) == -2.25), F1: 306, F2: uint8(v57)}
                                        _ = v63
                                    }
                                    fmt.Println(v10)
                                }
                            } else {
                                v64 := GenS3{F0: 12, F1: -2147483648}
                                _ = v64
                            }
                        } else {
                            delete(v6, (v37 + ((((((((v37 + "") + v37) + (v37 + ("foo bar" + ("a" + "foo bar")))) + ("\n\t" + "foo bar")) + v37) + v37) + v37) + "hello")))
                            var v65 int32 = -236
                            _ = v65
                        }
                        fmt.Println(v57, v57, fi56)
                    }
                    fmt.Println(v10, v4, v4)
                    _ = v60
                case "foo bar":
                    v66 := GenS3{F0: 127, F1: -1}
                    _ = v66
                case "\n\t":
                    v53 -= float32((((v10 + v10) - (v10 / (1.5 - v10))) - v10))
                default:
                    var v67 uint16 = 690
                    _ = v67
                }
                _ = fi56
                _ = v57
            }
            _ = fi55
        }
        v1 = append(v1, 4294967295)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
    _ = v9
}
