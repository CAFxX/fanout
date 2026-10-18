package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS4 struct { F0 bool; F1 int32; F2 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 uint8; F2 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint32; F1 uint16 }
type GenS12 struct { F0 int32 }
func genFn0() string {
    var v5 uint16 = 65535
    v5 = v5
    if false {
        v6 := []int{4294967295}
        v7 := GenS1{F0: 0, F1: false}
        _ = v6
        _ = v7
    }
    v8 := []int{32767}
    v9 := GenS2{F0: -1}
    _ = v5
    _ = v8
    _ = v9
    return (("foo bar" + (("" + (("hello" + "a") + ("\n\t" + ("\n\t" + "")))) + "foo bar")) + "a")
}
func genFn1(p0 string, p1 int16) uint8 {
    v10 := GenS2{F0: 0}
    if (40 >= -2147483648) {
        var v11 uint8 = 244
        switch 631 {
        case 481:
            v11 = (v11 + 42)
        case -9223372036854775808:
            switch p0 {
            case "\n\t":
                var v12 uint8 = (v11 - v11)
                _ = v12
            case "foo bar":
                for fi13 := 0; fi13 < 18; fi13++ {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        var v15 int = 982
                        _ = fi14
                        _ = v15
                    }
                    _ = fi13
                }
            case "hello":
                var v16 uint16 = 0
                _ = v16
            default:
                v17 := GenS3{F0: 705, F1: 1, F2: 2147483647}
                _ = v17
            }
        }
        fmt.Println(p0)
        _ = v11
    } else {
        fmt.Println(p1, p0)
        for fi18 := 0; fi18 < 12; fi18++ {
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 string = ((p0 + "foo bar") + p0)
                _ = fi19
                _ = v20
            }
            fmt.Println(p0, p1)
            if (false || false) {
                var v21 uint16 = 42
                switch p0 {
                case "foo bar":
                    var v22 uint32 = 4294967295
                    _ = v22
                case "":
                    p1 = p1
                case "hello":
                    v23 := GenS3{F0: (v21 >> ((2) & 63)), F1: 935, F2: 539}
                    _ = v23
                default:
                    for fi24 := 0; fi24 < 14; fi24++ {
                        if (!true) {
                            if true {
                                var v25 int32 = -811
                                v26 := []int{(fi24 % (((fi24 << ((65535) & 63))) | 1)), 947, 32768, (fi24 >> ((255) & 63))}
                                v27 := GenS2{F0: v25}
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            var v28 float64 = 1e10
                            _ = v28
                        } else {
                            p1 += (p1 * p1)
                        }
                        _ = fi24
                    }
                }
                _ = v21
            } else {
                var v29 int8 = 42
                _ = v29
            }
            _ = fi18
        }
    }
    if (-977 < 42) {
        v30 := GenS3{F0: 32767, F1: 0, F2: 65535}
        _ = v30
    }
    p1 = p1
    v31 := GenS4{F0: ((32767 <= 990) || true), F1: -1, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v10
    _ = v31
    return 255
}

func main() {
    var v1 uint = 67
    v2 := make(map[uint32]string)
    var v3 uint8 = 42
    v4 := make(map[int]string)
    var v32 int16 = -1
    v33 := GenS5{F0: int(v3)}
    v1 -= v1
    if ("\n\t" == ("a" + "\n\t")) {
        var v34 float32 = (1.5 * (-2.25 / ((float32(1.5) + ((float32(((0.0 + ((-2.25 + ((1.5 - float64(float32((-2.25 + 1e10)))) - float64(0.1))) * 1e10)) - float64(((0.0 / 1.5) / ((1e10 * (1.5 / 1e10)) - 0.0))))) + (0.0 / float32(((3.14159 / 3.14159) + (float64(1.5) / (3.14159 / 0.1)))))) + 0.1)) / 1.0)))
        var v35 uint16 = 0
        fmt.Println(v35)
        _ = v34
        _ = v35
    }
    var v36 float64 = 1.5
    var v37 int8 = 1
    var v38 uint64 = 163
    v36 += v36
    if false {
        var v39 int32 = 0
        _ = v39
    }
    if (v1 < 18446744073709551615) {
        for fi40 := 0; fi40 < 14; fi40++ {
            _ = len(v2)
            switch "a" {
            case "\n\t":
                if (true && true) {
                    v41 := []int{-1, (fi40 % ((775) | 1)), (fi40 + 9223372036854775807)}
                    v32 += (v32 | v32)
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 4; fi42++ {
                        var v43 uint32 = 1
                        switch "hello" {
                        case "a":
                            if false {
                                v43 = uint32(v1)
                                switch fi42 {
                                case 485:
                                    if false {
                                        v44 := GenS6{F0: false, F1: 127, F2: (-(v32))}
                                        v45 := GenS3{F0: uint16(v1), F1: (v43 | 2147483648), F2: (v43 << (((v1 >> ((uint(v43)) & 63))) & 63))}
                                        v46 := GenS4{F0: true, F1: 1, F2: (v38 << ((4294967295) & 63))}
                                        switch -2147483648 {
                                        case 821:
                                            var v47 uint16 = 0
                                            _ = v47
                                        default:
                                            v48 := []int{int(v32), -585, 127, fi42}
                                            _ = v48
                                        }
                                        v36 = v36
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                    } else {
                                        var v49 uint64 = v38
                                        v50 := GenS7{F0: (!(1 >= -9223372036854775808))}
                                        _ = v49
                                        _ = v50
                                    }
                                case -9223372036854775808:
                                    v1 = v1
                                case -951:
                                    fmt.Println(v1, v3, v3)
                                default:
                                    v36 = (((3.14159 * 1e10) - v36) + 1e10)
                                }
                            }
                        case "hello":
                            v51 := GenS8{F0: (1.5 < 1e10)}
                            _ = v51
                        case "foo bar":
                            v52 := v4[(fi42 << ((403) & 63))]
                            _ = v52
                        default:
                            var v53 int = -1
                            _ = v53
                        }
                        var v54 int16 = v32
                        _ = fi42
                        _ = v43
                        _ = v54
                    }
                    var v55 uint16 = 343
                    _ = v55
                }
            }
            v32 = (v32 * 32767)
            _ = fi40
        }
        if false {
            fmt.Println(v32, v36)
            var v56 uint32 = 0
            if false {
                fmt.Println(v36)
                var v57 float64 = -2.25
                if false {
                    for fi58 := 0; fi58 < 4; fi58++ {
                        v59 := GenS9{F0: 324}
                        v1 = uint(v56)
                        _ = fi58
                        _ = v59
                    }
                    v60 := GenS7{F0: ((int64(v3) <= 42) && false)}
                    if ((v3 * (v3 + (v3 / (((v3 % ((128) | 1))) | 1)))) > v3) {
                        v61 := []int{}
                        for fi62 := 0; fi62 < 6; fi62++ {
                            v36 -= v57
                            _ = fi62
                        }
                        _ = v61
                    } else {
                        fmt.Println(v37)
                        for fi63 := 0; fi63 < 20; fi63++ {
                            for fi64 := 0; fi64 < 4; fi64++ {
                                fmt.Println(v1)
                                switch fi64 {
                                case -205:
                                    v65 := GenS5{F0: fi64}
                                    _ = v65
                                case 42:
                                    v66 := GenS10{F0: v37}
                                    var v67 int32 = 476
                                    var v68 uint64 = v38
                                    _ = v66
                                    _ = v67
                                    _ = v68
                                case 0:
                                    v36 -= ((v36 * (v57 * v57)) * v57)
                                default:
                                    v37 -= -1
                                }
                                _ = fi64
                            }
                            var v69 int64 = -9223372036854775808
                            _ = fi63
                            _ = v69
                        }
                        v1 -= 255
                    }
                    _ = v60
                }
                _ = v57
            } else {
                v37 += 76
                for fi70 := 0; fi70 < 7; fi70++ {
                    fmt.Println(v37, v3)
                    for fi71 := 0; fi71 < 6; fi71++ {
                        if true {
                            v38 += (v38 >> ((v1) & 63))
                        } else {
                            v72 := GenS3{F0: 32767, F1: (v56 / (((v56 + v56)) | 1)), F2: v56}
                            if (((!(v3 <= (v3 >> (((^(v1))) & 63)))) || false) && (0 < 879)) {
                                switch "\n\t" {
                                case "foo bar":
                                    var v73 float64 = (v36 + ((float64(1.5) / -2.25) * 0.1))
                                    _ = v73
                                }
                            }
                            var v74 uint64 = 186
                            v32 += v32
                            _ = v72
                            _ = v74
                        }
                        fmt.Println(v37, v38, v1)
                        _ = fi71
                    }
                    v56 = 797
                    _ = fi70
                }
            }
            _ = v56
        }
        v32 -= (v32 + v32)
    } else {
        for fi75 := 0; fi75 < 15; fi75++ {
            for fi76 := 0; fi76 < 2; fi76++ {
                delete(v4, (fi76 * -9223372036854775808))
                switch ("" + "") {
                case "a":
                    if ((fi75 & fi75) > (fi75 & int(v32))) {
                        v77 := GenS11{F0: 0, F1: 185}
                        if false {
                            var v78 int32 = -1
                            var v79 bool = (2 != uint16(v38))
                            var v80 int = fi75
                            _ = v78
                            _ = v79
                            _ = v80
                        }
                        fmt.Println(v3)
                        _ = v77
                    }
                case "\n\t":
                    if ("foo bar" != "hello") {
                        v36 += v36
                    }
                }
                var v81 uint8 = v3
                _ = fi76
                _ = v81
            }
            v82 := GenS12{F0: -2147483648}
            _ = fi75
            _ = v82
        }
        v83 := []int{int(v32)}
        _ = v83
    }
    if true {
        var v84 uint32 = 0
        _ = v84
    }
    for fi85 := 0; fi85 < 6; fi85++ {
        var v86 int64 = 0
        for fi87 := 0; fi87 < 10; fi87++ {
            v1 = 1
            _ = fi87
        }
        fmt.Println(v3)
        _ = fi85
        _ = v86
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
