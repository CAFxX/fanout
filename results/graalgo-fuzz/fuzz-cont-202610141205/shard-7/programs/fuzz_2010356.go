package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint }
type GenS3 struct { F0 uint64; F1 int8 }

func main() {
    var v1 uint64 = 2
    var v2 float32 = 3.14159
    v3 := []uint32{850}
    v4 := make(map[int]int)
    var v5 float32 = v2
    v6 := make(map[string]bool)
    fmt.Println(v5)
    fmt.Println(v1)
    var v7 float64 = 1.5
    switch "" {
    case "":
        var v8 int64 = -9223372036854775808
        _ = v8
    default:
        switch ("foo bar" + "") {
        case "":
            fmt.Println(v5)
        case "hello":
            v5 = 1.5
        default:
            delete(v4, 21)
        }
    }
    switch -9223372036854775808 {
    case 32767:
        v2 -= float32(v7)
    case 221:
        v5 -= float32(3.14159)
    default:
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(v7, v1, v7)
            fmt.Println(fi9, v5, fi9)
            _ = fi9
        }
    }
    v4[32768] = -2147483648
    var v10 uint8 = 255
    switch ("\n\t" + "hello") {
    case "\n\t":
        for fi11 := 0; fi11 < 1; fi11++ {
            delete(v6, "")
            v3 = append(v3, 375)
            _ = fi11
        }
    case "":
        if (true && (!((!(!true)) && ((-127 != 473) && true)))) {
            if false {
                switch 0 {
                case 128:
                    if true {
                        v7 += ((0.1 / 1.0) / (v7 / (float64((((((v5 / v2) * 0.1) * (v2 - v2)) - v5) + -2.25)) / v7)))
                    } else {
                        if true {
                            v10 = uint8(v1)
                        } else {
                            v1 += (v1 + (v1 - v1))
                        }
                        for fi12 := 0; fi12 < 4; fi12++ {
                            v13 := v4[fi12]
                            _ = fi12
                            _ = v13
                        }
                    }
                case -1:
                    fmt.Println(v10)
                case 1:
                    v4[32768] = 794
                default:
                    var v14 uint64 = 717
                    _ = v14
                }
                for fi15 := 0; fi15 < 14; fi15++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch "foo bar" {
                    case "foo bar":
                        v5 += v5
                    default:
                        v10 = v10
                    }
                    _ = fi15
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        } else {
            var v16 bool = (65535 <= 2)
            for fi17 := 0; fi17 < 5; fi17++ {
                v2 = v2
                v18 := GenS1{F0: -1}
                for fi19 := 0; fi19 < 18; fi19++ {
                    if v16 {
                        switch (("foo bar" + ("hello" + "")) + "") {
                        case "":
                            if true {
                                switch ((("\n\t" + "\n\t") + "foo bar") + "\n\t") {
                                case "hello":
                                    v1 = v1
                                default:
                                    _ = len(v3)
                                }
                            } else {
                                if ((("" + ("\n\t" + "")) + "foo bar") > "a") {
                                    v1 = 2
                                    var v20 float64 = 3.14159
                                    if v16 {
                                        v16 = v16
                                    } else {
                                        var v21 bool = (!((v2 < v2) && v16))
                                        if (v21 && true) {
                                            var v22 int16 = 1
                                            v23 := GenS2{F0: 645, F1: (v1 / (((v1 & uint64(v10))) | 1)), F2: 18446744073709551615}
                                            var v24 bool = (0 < 0)
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = v21
                                    }
                                    _ = len(v3)
                                    v10 -= v10
                                    _ = v20
                                } else {
                                    switch 0 {
                                    case 9223372036854775807:
                                        v5 = v2
                                    default:
                                        if v16 {
                                            v16 = v16
                                            if (!((("foo bar" + "") + "hello") != "hello")) {
                                                var v25 float64 = v7
                                                v26 := v4[fi17]
                                                _ = v25
                                                _ = v26
                                            } else {
                                                v16 = (!(-82 != 808))
                                                v16 = v16
                                            }
                                            var v27 bool = (!(v1 < (v1 ^ 0)))
                                            _ = v27
                                        }
                                    }
                                    v2 += (1.5 + float32(v7))
                                }
                                var v28 bool = (-9223372036854775808 > 0)
                                var v29 uint16 = 2
                                _ = v28
                                _ = v29
                            }
                        case "hello":
                            v5 -= 1.5
                        case "foo bar":
                            var v30 uint8 = (^(v10))
                            _ = v30
                        }
                    } else {
                        var v31 int = (-(fi19))
                        _ = v31
                    }
                    fmt.Println(fi17, v10)
                    if true {
                        v16 = (v7 != ((v7 + ((v7 + v7) - (((v7 + ((1e10 / ((v7 / ((v7 + float64(float32((0.1 + 1e10)))) - -2.25)) * v7)) + ((1e10 * (0.0 / (v7 + v7))) * -2.25))) * ((v7 * v7) * v7)) + v7))) / v7))
                        v32 := GenS2{F0: 531, F1: 18446744073709551615, F2: 18446744073709551615}
                        v4[(^(fi19))] = (fi17 - -69)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v32
                    } else {
                        v5 -= (v2 - (0.0 + v2))
                    }
                    _ = len(v6)
                    _ = fi19
                }
                _ = fi17
                _ = v18
            }
            var v33 int = 1
            _ = v16
            _ = v33
        }
    case "a":
        _ = len(v4)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 128 {
    case 42:
        _ = len(v3)
    case -556:
        v34 := GenS1{F0: -872}
        _ = v34
    case 127:
        var v35 uint = 146
        _ = v35
    }
    if ((2 == 2) && true) {
        if false {
            var v36 int = -9223372036854775808
            v3 = append(v3, 2147483648)
            _ = v36
        } else {
            var v37 int16 = 536
            _ = v37
        }
        _ = len(v3)
    } else {
        if ((2 <= 255) || (false && true)) {
            switch "foo bar" {
            case "hello":
                var v38 int = 32767
                _ = v38
            case "a":
                if (127 < -43) {
                    v5 += (float32(0.1) - float32((v7 * v7)))
                    for fi39 := 0; fi39 < 12; fi39++ {
                        if false {
                            v2 += 0.1
                            var v40 float64 = (float64(v5) + 1.5)
                            _ = v40
                        } else {
                            delete(v4, fi39)
                        }
                        _ = fi39
                    }
                }
            }
        } else {
            if false {
                for fi41 := 0; fi41 < 2; fi41++ {
                    if ((((!true) && false) && ((0.1 - (v7 / v7)) > 0.0)) && false) {
                        _ = len(v3)
                        switch "foo bar" {
                        case "hello":
                            fmt.Println(v7, v1, v10)
                        case "foo bar":
                            switch 2147483647 {
                            case 32768:
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    v6[""] = false
                                    var v43 int32 = int32(v10)
                                    _ = fi42
                                    _ = v43
                                }
                            case 0:
                                fmt.Println(v1, v2, fi41)
                            default:
                                var v44 float32 = 0.1
                                _ = v44
                            }
                        case "a":
                            if false {
                                _ = len(v3)
                            }
                        }
                    }
                    _ = fi41
                }
            } else {
                if (1 != 358) {
                    if (!((!false) && (0 < 2147483647))) {
                        v45 := GenS2{F0: uint16(v1), F1: v1, F2: 2}
                        var v46 bool = (false && ((v1 - v1) > (v1 << ((80) & 63))))
                        if true {
                            var v47 float32 = (v2 * 0.0)
                            v1 += uint64(v10)
                            _ = v47
                        }
                        _ = v45
                        _ = v46
                    }
                } else {
                    v48 := v6["\n\t"]
                    _ = len(v3)
                    _ = v48
                }
                switch 4294967295 {
                case 1:
                    fmt.Println(v7, v10)
                case -657:
                    v49 := GenS3{F0: v1, F1: 100}
                    _ = v49
                default:
                    v5 = (float32(-2.25) / v5)
                }
            }
            fmt.Println(v7, v1, v10)
        }
    }
    v50 := GenS3{F0: uint64(v10), F1: 1}
    var v51 int8 = 0
    if (!true) {
        for fi52 := 0; fi52 < 11; fi52++ {
            if (!(v51 > v51)) {
                fmt.Println(v10)
                var v53 float32 = (v5 / 1e10)
                var v54 uint64 = (^(v1))
                _ = v53
                _ = v54
            }
            for fi55 := 0; fi55 < 19; fi55++ {
                v51 += -84
                _ = fi55
            }
            for fi56 := 0; fi56 < 11; fi56++ {
                _ = len(v4)
                var v57 uint8 = v10
                _ = fi56
                _ = v57
            }
            _ = fi52
        }
        v5 = ((v2 / v2) / v2)
    } else {
        fmt.Println(v51, v1)
        for fi58 := 0; fi58 < 6; fi58++ {
            fmt.Println(v10)
            var v59 int16 = 1
            _ = fi58
            _ = v59
        }
        for fi60 := 0; fi60 < 3; fi60++ {
            if false {
                var v61 int32 = -2147483648
                if (2147483647 >= 232) {
                    var v62 uint64 = (v1 & 1)
                    _ = v62
                } else {
                    var v63 int16 = -1
                    for fi64 := 0; fi64 < 19; fi64++ {
                        switch "foo bar" {
                        case "hello":
                            if len(v3) > 0 {
                                v3[0] = 2
                            }
                        }
                        if false {
                            var v65 int16 = v63
                            _ = v65
                        } else {
                            var v66 int64 = 0
                            var v67 int = 42
                            var v68 uint8 = (v10 >> ((98) & 63))
                            _ = v66
                            _ = v67
                            _ = v68
                        }
                        _ = fi64
                    }
                    var v69 int32 = v61
                    _ = v63
                    _ = v69
                }
                v10 += (v10 >> ((42) & 63))
                _ = v61
            } else {
                var v70 uint8 = v10
                var v71 uint32 = 65535
                v3 = append(v3, (v71 + (v71 | 199)))
                _ = v70
                _ = v71
            }
            switch ("\n\t" + (("hello" + "") + "")) {
            case "":
                v10 = (v10 ^ 2)
            case "\n\t":
                if (1 == -1) {
                    v72 := GenS2{F0: 65535, F1: 4294967295, F2: 4294967295}
                    fmt.Println(fi60, v51)
                    _ = v72
                }
            default:
                v73 := GenS2{F0: 2, F1: v1, F2: 131}
                _ = v73
            }
            _ = fi60
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
}
