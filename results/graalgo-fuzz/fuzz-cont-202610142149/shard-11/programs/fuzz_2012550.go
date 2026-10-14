package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS3 struct { F0 string; F1 uint; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 uint16; F1 uint16; F2 int16 }
func genFn0() uint8 {
    if (!true) {
        v6 := []int{-1, 9223372036854775807, -1}
        var v7 int64 = 1
        for fi8 := 0; fi8 < 19; fi8++ {
            switch "foo bar" {
            case "hello":
                if len(v6) > 0 {
                    v6[0] = fi8
                }
            case "\n\t":
                for fi9 := 0; fi9 < 3; fi9++ {
                    v7 += (v7 * (v7 - v7))
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi9
                }
            default:
                v10 := GenS1{F0: 587, F1: "\n\t", F2: 782}
                _ = v10
            }
            fmt.Println(v7, v7)
            v7 += (v7 | -322)
            _ = fi8
        }
        _ = v6
        _ = v7
    }
    switch 9223372036854775807 {
    case -400:
        var v11 int32 = 507
        _ = v11
    case 588:
        var v12 int = -2147483648
        _ = v12
    }
    var v13 uint64 = 18446744073709551615
    _ = v13
    return 44
}

func main() {
    var v1 uint16 = 65535
    var v2 float64 = 0.0
    v3 := make(map[string]uint32)
    var v4 uint8 = 0
    v5 := []uint32{}
    v14 := GenS1{F0: 186, F1: "hello", F2: 1}
    v4 += (^(v4))
    delete(v3, "")
    if (18 == 18446744073709551615) {
        switch ("foo bar" + ("\n\t" + "a")) {
        case "hello":
            v2 = v2
        case "foo bar":
            _ = len(v3)
        }
    } else {
        v2 += 0.1
    }
    if true {
        v2 += v2
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v15 int64 = -1
        _ = v15
    }
    switch "hello" {
    case "a":
        var v16 uint = 0
        _ = v16
    case "":
        v2 += 1.5
    case "\n\t":
        v17 := GenS1{F0: 4294967295, F1: "hello", F2: 2}
        _ = v17
    default:
        for fi18 := 0; fi18 < 10; fi18++ {
            for fi19 := 0; fi19 < 2; fi19++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi19
            }
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            } else {
                for fi20 := 0; fi20 < 9; fi20++ {
                    switch "" {
                    case "hello":
                        v21 := GenS1{F0: 875, F1: (("a" + "a") + ("\n\t" + "\n\t")), F2: 2}
                        _ = v21
                    default:
                        switch (("hello" + ("\n\t" + "a")) + "hello") {
                        case "":
                            var v22 int16 = 915
                            _ = v22
                        case "a":
                            fmt.Println(v4, v4)
                        case "hello":
                            var v23 uint32 = 4294967295
                            _ = v23
                        }
                    }
                    fmt.Println(v2, v1, v4)
                    switch (fi18 + 42) {
                    case 0:
                        v1 = v1
                    }
                    _ = fi20
                }
                if len(v5) > 0 {
                    v5[0] = 65535
                }
            }
            v2 = float64(0.1)
            _ = fi18
        }
    }
    v1 += (v1 >> ((42) & 63))
    switch 32768 {
    case -9223372036854775808:
        v24 := GenS2{F0: 42, F1: 42, F2: -418}
        _ = v24
    case 42:
        switch "" {
        case "a":
            v25 := GenS1{F0: 2, F1: "foo bar", F2: 108}
            _ = v25
        case "foo bar":
            v26 := GenS2{F0: 600, F1: 18446744073709551615, F2: -9223372036854775808}
            _ = v26
        case "hello":
            v1 += (^(v1))
        default:
            switch "\n\t" {
            case "foo bar":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "hello":
                v4 += v4
            }
        }
    case 1:
        _ = len(v5)
    }
    fmt.Println(v1, v4, v1)
    var v27 int32 = -144
    v28 := GenS1{F0: uint(v1), F1: "a", F2: 54}
    switch ("" + "hello") {
    case "a":
        var v29 string = ("a" + ("foo bar" + "foo bar"))
        _ = v29
    case "foo bar":
        fmt.Println(v2, v1, v27)
    case "\n\t":
        if (v2 < v2) {
            var v30 uint32 = 2
            _ = v30
        } else {
            v2 = -2.25
            v31 := GenS3{F0: "a", F1: 42, F2: (v4 / ((42) | 1))}
            var v32 int64 = -1
            _ = v31
            _ = v32
        }
    }
    v5 = append(v5, 4294967295)
    if true {
        fmt.Println(v4, v1)
        switch ("foo bar" + ("hello" + "foo bar")) {
        case "a":
            if (!false) {
                fmt.Println(v2, v4)
                v27 = 211
            }
        case "foo bar":
            var v33 uint8 = (v4 >> ((0) & 63))
            _ = v33
        }
        var v34 string = ("" + "a")
        _ = v34
    } else {
        var v35 uint64 = 4294967295
        if false {
            _ = len(v5)
            if (1000 < 139) {
                if (!(true || false)) {
                    switch 0 {
                    case 42:
                        fmt.Println(v4, v4, v4)
                    case -460:
                        fmt.Println(v1)
                    case 716:
                        if true {
                            var v36 int = 32767
                            switch ("hello" + (("a" + ("foo bar" + "\n\t")) + "foo bar")) {
                            case "hello":
                                var v37 int32 = v27
                                _ = v37
                            }
                            v38 := GenS4{F0: (^(v35))}
                            _ = v36
                            _ = v38
                        } else {
                            var v39 uint32 = 459
                            switch "foo bar" {
                            case "foo bar":
                                for fi40 := 0; fi40 < 13; fi40++ {
                                    v1 = 42
                                    var v41 string = ((("" + ("hello" + "foo bar")) + "\n\t") + "foo bar")
                                    v35 = v35
                                    _ = fi40
                                    _ = v41
                                }
                            case "hello":
                                v42 := GenS5{F0: -1}
                                _ = v42
                            default:
                                switch 4294967295 {
                                case 127:
                                    var v43 uint16 = v1
                                    _ = v43
                                case -893:
                                    var v44 int8 = -10
                                    _ = v44
                                case -370:
                                    var v45 int8 = int8(v35)
                                    _ = v45
                                default:
                                    var v46 int8 = -1
                                    _ = v46
                                }
                            }
                            _ = v39
                        }
                    }
                    v1 -= (v1 * (v1 ^ v1))
                }
                v1 -= 32767
                var v47 uint8 = (v4 + (v4 ^ (v4 << ((0) & 63))))
                _ = v47
            } else {
                var v48 int8 = -128
                _ = v48
            }
        }
        if false {
            v49 := GenS1{F0: 823, F1: "\n\t", F2: v35}
            if false {
                v5 = append(v5, 747)
                var v50 uint16 = uint16(v4)
                _ = v50
            } else {
                fmt.Println(v2, v4)
                var v51 uint = 2
                _ = v51
            }
            _ = v49
        }
        _ = v35
    }
    var v52 int = -9223372036854775808
    fmt.Println(v1, v4)
    if ((true || true) || false) {
        if (false && ((8 == 0) && (!false))) {
            v53 := GenS6{F0: 185, F1: 255, F2: -1}
            v54 := GenS7{F0: 1, F1: -21}
            for fi55 := 0; fi55 < 12; fi55++ {
                var v56 int32 = (v27 << ((2) & 63))
                if true {
                    switch fi55 {
                    case 9223372036854775807:
                        v52 = (fi55 << ((4294967295) & 63))
                    case 127:
                        delete(v3, "foo bar")
                    case 2147483647:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        v56 -= (v27 << ((42) & 63))
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    fmt.Println(v56, v27, v2)
                }
                _ = fi55
                _ = v56
            }
            _ = v53
            _ = v54
        } else {
            v27 += -54
        }
    } else {
        if (!true) {
            v2 = float64(0.0)
            switch (v52 >> ((18446744073709551615) & 63)) {
            case -752:
                v27 = (-(v27))
            case 2147483647:
                if false {
                    if ((v27 & 2147483647) < v27) {
                        fmt.Println(v4, v2, v2)
                    }
                    var v57 int16 = -183
                    _ = v57
                }
            case 0:
                v58 := GenS1{F0: 65535, F1: "hello", F2: 18446744073709551615}
                _ = v58
            }
            switch "" {
            case "":
                v59 := GenS1{F0: 42, F1: "\n\t", F2: 519}
                _ = v59
            case "hello":
                if true {
                    for fi60 := 0; fi60 < 10; fi60++ {
                        _ = len(v3)
                        _ = fi60
                    }
                }
            }
        } else {
            var v61 uint16 = (v1 + 42)
            if (float64(1.5) <= float64(float32(v2))) {
                v2 -= (v2 * ((float64(3.14159) - v2) / v2))
                v62 := GenS2{F0: uint64(v4), F1: 18446744073709551615, F2: 0}
                v63 := GenS8{F0: v1, F1: (^(v61)), F2: 0}
                _ = v62
                _ = v63
            }
            _ = v61
        }
        v64 := v3["\n\t"]
        _ = v64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
