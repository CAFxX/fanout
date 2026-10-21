package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 string }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 uint32; F2 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint8; F1 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int8 }

func main() {
    var v1 int16 = 1
    var v2 uint16 = 0
    v3 := []string{"\n\t", (("hello" + ("" + "")) + (("hello" + "") + "foo bar")), ("" + "")}
    v4 := make(map[int]int)
    var v5 uint = 255
    var v6 int8 = 2
    v3 = append(v3, "hello")
    fmt.Println(v1)
    fmt.Println(v5)
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v1)
        var v8 uint16 = v2
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 uint64 = 1
        _ = fi9
        _ = v10
    }
    delete(v4, 4294967295)
    switch (("foo bar" + "") + "foo bar") {
    case "\n\t":
        if (!(true && true)) {
            switch "foo bar" {
            case "a":
                v4[-691] = 127
            }
            if (v2 > (v2 / ((v2) | 1))) {
                v11 := GenS1{F0: -321, F1: 238, F2: "hello"}
                v6 -= int8(v1)
                _ = v11
            } else {
                v5 += (^(v5))
                _ = len(v3)
            }
            v6 = 100
        } else {
            var v12 int64 = int64(v5)
            switch "" {
            case "foo bar":
                var v13 int = -492
                _ = v13
            case "":
                var v14 int8 = 33
                _ = v14
            case "\n\t":
                for fi15 := 0; fi15 < 11; fi15++ {
                    switch "hello" {
                    case "foo bar":
                        _ = len(v3)
                    }
                    var v16 int32 = 0
                    fmt.Println(v6, v16, v12)
                    _ = fi15
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 8; fi17++ {
                if (!((true || true) || (((float32((0.1 - float64(-2.25))) / 1e10) <= 1e10) || true))) {
                    v2 -= 32767
                    fmt.Println(fi17)
                }
                _ = fi17
            }
            _ = v12
        }
    case "hello":
        for fi18 := 0; fi18 < 10; fi18++ {
            v1 += (v1 - int16(fi18))
            _ = fi18
        }
    default:
        var v19 uint32 = 419
        _ = v19
    }
    switch 4294967295 {
    case -9223372036854775808:
        for fi20 := 0; fi20 < 13; fi20++ {
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
            v2 = v2
            _ = fi20
        }
    case -2147483648:
        switch 99 {
        case -2147483648:
            switch "" {
            case "\n\t":
                if (0.0 >= 0.1) {
                    switch "" {
                    case "":
                        fmt.Println(v1, v6, v6)
                    }
                } else {
                    var v21 uint = v5
                    var v22 uint16 = v2
                    v22 += v22
                    _ = v21
                    _ = v22
                }
            case "hello":
                v5 -= (v5 / (((v5 - v5)) | 1))
            case "a":
                if (0 < v1) {
                    if false {
                        switch 87 {
                        case 2147483647:
                            delete(v4, 300)
                        case 32767:
                            _ = len(v4)
                        }
                    }
                } else {
                    v1 -= (v1 << (((v5 / (((v5 - 93)) | 1))) & 63))
                    var v23 int64 = 1
                    for fi24 := 0; fi24 < 2; fi24++ {
                        var v25 int = (fi24 | -989)
                        _ = fi24
                        _ = v25
                    }
                    _ = v23
                }
            }
        default:
            if false {
                switch ("hello" + "") {
                case "hello":
                    for fi26 := 0; fi26 < 1; fi26++ {
                        var v27 int8 = 127
                        _ = fi26
                        _ = v27
                    }
                case "a":
                    if false {
                        var v28 int8 = (-(v6))
                        if (!false) {
                            v4[2147483647] = -9223372036854775808
                        }
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 14; fi29++ {
                            if (float64((0.1 * float32(((1.5 / ((((((1e10 / float64(float32(3.14159))) - 1e10) + (((0.1 + float64(0.0)) + 0.1) / 1e10)) - (0.1 - ((1.5 + -2.25) / (((3.14159 + (((1.5 * 0.0) / 1.5) - 1.5)) / 0.1) * (float64((float32(-2.25) - float32(3.14159))) + 1e10))))) + float64((-2.25 / -2.25))) + (1e10 * 1e10))) + (-2.25 * 0.0))))) > 1.5) {
                                v30 := GenS2{F0: -1, F1: v5}
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    v2 += (v2 + v2)
                                    _ = fi31
                                }
                                _ = v30
                            } else {
                                switch "" {
                                case "hello":
                                    v32 := GenS2{F0: 1, F1: v5}
                                    _ = v32
                                case "":
                                    var v33 uint = (v5 - (v5 * 91))
                                    _ = v33
                                case "\n\t":
                                    delete(v4, fi29)
                                default:
                                    if (float64(0.1) != float64(1e10)) {
                                        var v34 int = -2147483648
                                        _ = v34
                                    } else {
                                        fmt.Println(v2)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        fmt.Println(v2, v5)
                                    }
                                }
                                var v35 uint8 = 42
                                _ = v35
                            }
                            _ = fi29
                        }
                        v36 := GenS3{F0: 255}
                        _ = v36
                    }
                case "":
                    v6 += v6
                }
            }
        }
    case 127:
        v1 += (v1 ^ v1)
    default:
        v1 -= -1
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        _ = len(v3)
        _ = fi37
    }
    var v38 int8 = 0
    v39 := GenS3{F0: v5}
    if false {
        if (4294967295 < 42) {
            if (v2 != v2) {
                for fi40 := 0; fi40 < 11; fi40++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi40
                }
                v3 = append(v3, "foo bar")
                var v41 int8 = v6
                _ = v41
            } else {
                if false {
                    v42 := GenS4{F0: 0, F1: 42, F2: 4294967295}
                    if (42 == 19) {
                        v43 := GenS2{F0: 694, F1: (^(v5))}
                        _ = v43
                    } else {
                        _ = len(v3)
                        v44 := GenS5{F0: (v1 ^ v1)}
                        _ = v44
                    }
                    _ = v42
                }
                switch "foo bar" {
                case "":
                    v45 := GenS4{F0: -9223372036854775808, F1: 2147483648, F2: 184}
                    _ = v45
                case "\n\t":
                    if false {
                        v6 = 90
                    } else {
                        if false {
                            for fi46 := 0; fi46 < 19; fi46++ {
                                if false {
                                    var v47 uint = v5
                                    var v48 string = "a"
                                    v6 -= (v38 | int8(v1))
                                    _ = v47
                                    _ = v48
                                }
                                for fi49 := 0; fi49 < 7; fi49++ {
                                    v3 = append(v3, "hello")
                                    v50 := GenS6{F0: 128, F1: 32768}
                                    _ = fi49
                                    _ = v50
                                }
                                var v51 float32 = (((0.0 + 3.14159) - 3.14159) - 1.5)
                                _ = fi46
                                _ = v51
                            }
                            if (!true) {
                                v52 := GenS4{F0: int(v1), F1: 504, F2: uint64(v2)}
                                _ = v52
                            }
                        } else {
                            v2 = v2
                        }
                        switch 2147483647 {
                        case 564:
                            if false {
                                for fi53 := 0; fi53 < 9; fi53++ {
                                    if ((1.5 + 1e10) < 0.1) {
                                        v2 = v2
                                        v38 = v38
                                        v3 = append(v3, "hello")
                                    } else {
                                        for fi54 := 0; fi54 < 6; fi54++ {
                                            switch ("a" + "a") {
                                            case "hello":
                                                var v55 int8 = (v38 % ((v6) | 1))
                                                v5 -= v5
                                                v55 -= (v6 * (v55 % ((-1) | 1)))
                                                _ = v55
                                            }
                                            if (v2 > (v2 / ((42) | 1))) {
                                                switch fi54 {
                                                case 127:
                                                    v56 := GenS6{F0: 211, F1: 834}
                                                    _ = v56
                                                case 32767:
                                                    v6 = (v6 << ((uint(v2)) & 63))
                                                default:
                                                    v57 := GenS7{F0: (v2 | (^(v2)))}
                                                    _ = v57
                                                }
                                                v2 = v2
                                                if (1 <= 0) {
                                                    var v58 uint16 = v2
                                                    v59 := GenS8{F0: (v38 & 43)}
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    v6 += (^(v38))
                                                    v4[(fi53 ^ 95)] = (fi54 % ((19) | 1))
                                                    v2 = v2
                                                    if (!false) {
                                                        v60 := GenS5{F0: (-(v1))}
                                                        _ = v60
                                                    }
                                                }
                                                var v61 int32 = 2147483647
                                                _ = v61
                                            }
                                            v1 -= (v1 % (((^(v1))) | 1))
                                            v62 := GenS5{F0: -973}
                                            var v63 uint8 = 127
                                            v64 := GenS8{F0: v38}
                                            _ = fi54
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                        }
                                    }
                                    var v65 int8 = (v38 ^ -1)
                                    if (true || true) {
                                        if len(v3) > 0 {
                                            v3[0] = ("foo bar" + "\n\t")
                                        }
                                    }
                                    _ = fi53
                                    _ = v65
                                }
                            } else {
                                v3 = append(v3, ("foo bar" + ("\n\t" + "a")))
                            }
                        case -9223372036854775808:
                            for fi66 := 0; fi66 < 3; fi66++ {
                                fmt.Println(v38)
                                var v67 int16 = (v1 << ((v5) & 63))
                                fmt.Println(v67)
                                _ = fi66
                                _ = v67
                            }
                        case 9223372036854775807:
                            var v68 bool = (!(!false))
                            _ = v68
                        }
                    }
                case "hello":
                    v2 = (v2 << (((^(v5))) & 63))
                default:
                    var v69 uint8 = 255
                    _ = v69
                }
            }
            v70 := GenS7{F0: 0}
            _ = v70
        } else {
            fmt.Println(v5, v2, v5)
            v1 = v1
            v71 := GenS8{F0: (v38 % (((-(v38))) | 1))}
            _ = v71
        }
        v72 := GenS6{F0: 42, F1: -44}
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
