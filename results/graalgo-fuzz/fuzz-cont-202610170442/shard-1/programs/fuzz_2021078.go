package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int32; F2 int }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int16; F1 int16; F2 int8 }

func main() {
    var v1 int = 309
    var v2 bool = false
    var v3 float32 = float32((0.1 + 1.5))
    var v4 uint32 = 2
    var v5 uint8 = 128
    v6 := []uint32{(^(v4)), 42, (v4 ^ v4), v4}
    v6 = append(v6, v4)
    _ = len(v6)
    v7 := GenS1{F0: "\n\t", F1: 1, F2: v4}
    var v8 float64 = 0.0
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            if true {
                fmt.Println(v4, v1)
            }
            fmt.Println(v5)
            v4 -= (v4 - v4)
        }
        _ = fi9
    }
    if false {
        switch "a" {
        case "foo bar":
            _ = len(v6)
        case "hello":
            v4 = v4
        default:
            if false {
                fmt.Println(v5, v1, v4)
            }
        }
        fmt.Println(v8)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        if v2 {
            switch (("" + "") + ("hello" + "foo bar")) {
            case "hello":
                v5 += 79
            case "\n\t":
                if (v2 || v2) {
                    v10 := GenS1{F0: "", F1: 552, F2: (v4 + v4)}
                    fmt.Println(v5)
                    _ = v10
                } else {
                    v11 := GenS2{F0: v5}
                    _ = len(v6)
                    for fi12 := 0; fi12 < 12; fi12++ {
                        v13 := GenS3{F0: (^(v5)), F1: 500, F2: -320}
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                }
            default:
                v4 = v4
            }
            v2 = v2
        } else {
            v14 := GenS4{F0: -1000, F1: v4}
            switch "foo bar" {
            case "\n\t":
                if len(v6) > 0 {
                    v6[0] = (^(v4))
                }
            case "hello":
                if true {
                    switch v1 {
                    case -374:
                        v1 = v1
                    case -706:
                        var v15 int = -152
                        _ = v15
                    default:
                        v16 := GenS3{F0: uint8(v4), F1: 2147483647, F2: (v1 ^ int(v5))}
                        _ = v16
                    }
                    v4 += v4
                    v3 -= (v3 / v3)
                } else {
                    v1 = v1
                    v17 := GenS2{F0: (v5 * (^(v5)))}
                    var v18 uint = 2
                    _ = v17
                    _ = v18
                }
            default:
                switch 832 {
                case -9223372036854775808:
                    for fi19 := 0; fi19 < 12; fi19++ {
                        if v2 {
                            var v20 int32 = -2147483648
                            _ = v20
                        }
                        _ = fi19
                    }
                default:
                    if (825 != 32767) {
                        if len(v6) > 0 {
                            v6[0] = v4
                        }
                    } else {
                        for fi21 := 0; fi21 < 19; fi21++ {
                            var v22 bool = false
                            _ = fi21
                            _ = v22
                        }
                        v5 += v5
                    }
                }
            }
            var v23 int32 = 1
            _ = v14
            _ = v23
        }
        v4 = (v4 * (v4 & 4294967295))
    }
    _ = len(v6)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        v6 = append(v6, 0)
        v3 += v3
        _ = fi24
    }
    v6 = append(v6, 2)
    for fi25 := 0; fi25 < 2; fi25++ {
        if v2 {
            v26 := GenS3{F0: v5, F1: -1, F2: 32768}
            if (v2 && (-350 >= 232)) {
                v5 = v5
            }
            _ = v26
        }
        v2 = v2
        _ = fi25
    }
    var v27 uint16 = 32767
    switch "foo bar" {
    case "\n\t":
        v8 += 1e10
    default:
        fmt.Println(v8)
    }
    v3 -= ((((float32(-2.25) - 0.1) / -2.25) / float32(float64(v3))) * (0.1 * 1e10))
    v3 = 1e10
    v27 = v27
    switch v1 {
    case -2147483648:
        switch "foo bar" {
        case "\n\t":
            switch v1 {
            case -1:
                v6 = append(v6, (^(v4)))
            default:
                var v28 int64 = 42
                _ = v28
            }
        case "":
            v6 = append(v6, v4)
        default:
            for fi29 := 0; fi29 < 14; fi29++ {
                if len(v6) > 0 {
                    v6[0] = (^(v4))
                }
                var v30 uint32 = v4
                v30 = 42
                _ = fi29
                _ = v30
            }
        }
    case 42:
        switch ("" + "") {
        case "":
            switch (-(v1)) {
            case 721:
                if v2 {
                    for fi31 := 0; fi31 < 1; fi31++ {
                        var v32 int8 = 14
                        v3 = 3.14159
                        v8 -= -2.25
                        _ = fi31
                        _ = v32
                    }
                } else {
                    for fi33 := 0; fi33 < 16; fi33++ {
                        var v34 int32 = -1
                        _ = fi33
                        _ = v34
                    }
                    if len(v6) > 0 {
                        v6[0] = v4
                    }
                    v35 := GenS5{F0: 2, F1: 127}
                    _ = v35
                }
            case 32768:
                for fi36 := 0; fi36 < 20; fi36++ {
                    if true {
                        v6 = append(v6, (v4 | (v4 * (v4 | 42))))
                        var v37 float64 = 0.1
                        for fi38 := 0; fi38 < 19; fi38++ {
                            v27 += (v27 % (((^(v27))) | 1))
                            fmt.Println(fi36)
                            var v39 float64 = v8
                            if true {
                                fmt.Println(v37, v1, v1)
                                if len(v6) > 0 {
                                    v6[0] = v4
                                }
                                if v2 {
                                    var v40 float32 = (3.14159 / v3)
                                    _ = v40
                                } else {
                                    if len(v6) > 0 {
                                        v6[0] = 42
                                    }
                                    var v41 uint16 = v27
                                    _ = v41
                                }
                            } else {
                                switch (^(v1)) {
                                case 9223372036854775807:
                                    v42 := GenS2{F0: (^(v5))}
                                    _ = v42
                                case -2147483648:
                                    v37 = (v39 * v39)
                                case 42:
                                    fmt.Println(v2, v27)
                                }
                                var v43 int32 = 797
                                var v44 float64 = -2.25
                                _ = v43
                                _ = v44
                            }
                            _ = fi38
                            _ = v39
                        }
                        _ = v37
                    } else {
                        var v45 bool = v2
                        _ = v45
                    }
                    var v46 int64 = -680
                    for fi47 := 0; fi47 < 14; fi47++ {
                        v6 = append(v6, (v4 / ((0) | 1)))
                        v2 = (false || v2)
                        _ = fi47
                    }
                    _ = fi36
                    _ = v46
                }
            case 1:
                switch ("\n\t" + "foo bar") {
                case "a":
                    if (v4 != v4) {
                        if true {
                            v4 = 2
                            v48 := GenS6{F0: ""}
                            _ = v48
                        }
                        fmt.Println(v4, v3)
                    }
                case "\n\t":
                    v3 += (v3 / 1e10)
                }
            default:
                if (!(v2 && (42 < 18446744073709551615))) {
                    var v49 uint16 = v27
                    v8 -= (-2.25 + float64((((-2.25 + (v3 - (v3 / ((((v3 - v3) - -2.25) / (v3 + float32(((1e10 * (1.5 * v8)) / v8)))) * v3)))) + 0.1) - v3)))
                    var v50 uint = 663
                    _ = v49
                    _ = v50
                } else {
                    v8 = float64(v3)
                }
            }
        case "\n\t":
            v3 = (((3.14159 * v3) - float32(v8)) - 1e10)
        case "hello":
            if (v8 < float64(v3)) {
                if ((v2 && (32767 == -255)) || (!(v2 || (!(-2147483648 != -1))))) {
                    v1 += 366
                } else {
                    _ = len(v6)
                    var v51 uint32 = v4
                    _ = v51
                }
                v52 := GenS7{F0: 506, F1: -1, F2: 100}
                switch ("hello" + ("\n\t" + "\n\t")) {
                case "foo bar":
                    v53 := GenS7{F0: 483, F1: 1000, F2: 0}
                    _ = v53
                }
                _ = v52
            }
        }
    case 32768:
        if v2 {
            var v54 int64 = -693
            for fi55 := 0; fi55 < 12; fi55++ {
                var v56 string = ("a" + "")
                _ = fi55
                _ = v56
            }
            var v57 int8 = 33
            _ = v54
            _ = v57
        } else {
            var v58 float64 = v8
            v5 = (^(v5))
            for fi59 := 0; fi59 < 1; fi59++ {
                if (v27 != v27) {
                    v60 := GenS3{F0: (v5 | (v5 ^ v5)), F1: 470, F2: 4294967295}
                    var v61 float64 = 1.5
                    var v62 int8 = 1
                    _ = v60
                    _ = v61
                    _ = v62
                }
                _ = fi59
            }
            _ = v58
        }
    default:
        v27 = 578
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
