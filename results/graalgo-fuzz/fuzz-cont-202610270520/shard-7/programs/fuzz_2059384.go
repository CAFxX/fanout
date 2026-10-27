package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 string; F2 int32 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint16, p1 bool) uint16 {
    fmt.Println(p1, p0)
    p1 = (65535 != 1)
    p1 = ((p0 / ((p0) | 1)) <= (p0 % (((p0 << ((uint(p0)) & 63))) | 1)))
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := []int{-9223372036854775808, fi6, (fi6 & (fi6 % (((-(fi6))) | 1)))}
        fmt.Println(fi6, fi6)
        _ = fi6
        _ = v7
    }
    switch 32767 {
    case 32768:
        v8 := []int{4294967295, -1, 32767}
        _ = v8
    case 127:
        switch 171 {
        case 9223372036854775807:
            if (0.1 <= 1.5) {
                for fi9 := 0; fi9 < 9; fi9++ {
                    if ((fi9 + 32767) <= fi9) {
                        v10 := []int{(fi9 & fi9), (fi9 & fi9), fi9}
                        p1 = false
                        _ = v10
                    }
                    p1 = p1
                    var v11 string = ""
                    _ = fi9
                    _ = v11
                }
                p0 -= p0
            }
        }
    case 0:
        var v12 int32 = -1
        _ = v12
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int64 = 42
    v2 := []string{(("\n\t" + "") + "foo bar")}
    v3 := []uint32{4294967295, 2147483647, 1}
    v4 := make(map[uint32]bool)
    var v5 float64 = (0.0 + 3.14159)
    v1 = v1
    for fi13 := 0; fi13 < 20; fi13++ {
        fmt.Println(v1, v1)
        for fi14 := 0; fi14 < 12; fi14++ {
            if false {
                var v15 bool = (!false)
                v15 = v15
                _ = v15
            }
            switch "a" {
            case "hello":
                _ = len(v2)
            case "foo bar":
                var v16 uint32 = 42
                _ = v16
            case "\n\t":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                if ((127 == -128) && ("foo bar" <= ("hello" + "hello"))) {
                    var v17 int32 = 192
                    _ = v17
                } else {
                    v1 -= v1
                    if true {
                        if false {
                            var v18 uint32 = 2147483648
                            if true {
                                var v19 int16 = -1
                                _ = v19
                            } else {
                                fmt.Println(v1, v5, v5)
                                v4[2147483648] = false
                                var v20 uint64 = 4294967295
                                _ = v20
                            }
                            var v21 int32 = 1
                            _ = v18
                            _ = v21
                        } else {
                            var v22 uint32 = 2147483647
                            var v23 string = ("" + ("foo bar" + ("hello" + "")))
                            switch (v23 + v23) {
                            case "":
                                delete(v4, v22)
                            default:
                                var v24 string = "foo bar"
                                _ = v24
                            }
                            v25 := GenS1{F0: 828, F1: 128, F2: uint8(v22)}
                            _ = v22
                            _ = v23
                            _ = v25
                        }
                        fmt.Println(fi14, v1, fi14)
                        fmt.Println(fi13, v5, v1)
                        var v26 int = fi14
                        _ = v26
                    } else {
                        v5 = v5
                    }
                    if (((float64((-2.25 / (float32(((v5 + 3.14159) + 1e10)) * (1.5 - 1.5)))) - 1e10) + (float64(1e10) + v5)) > float64(0.0)) {
                        if (!(2 >= 4294967295)) {
                            v5 = v5
                            if false {
                                fmt.Println(fi14)
                                v27 := GenS2{F0: (^(fi13))}
                                switch "hello" {
                                case "\n\t":
                                    v28 := GenS1{F0: 42, F1: 0, F2: 38}
                                    _ = v28
                                case "foo bar":
                                    var v29 int = fi13
                                    _ = v29
                                case "":
                                    delete(v4, 2)
                                default:
                                    v5 -= 0.0
                                }
                                _ = v27
                            } else {
                                v3 = append(v3, 170)
                                var v30 float32 = float32((3.14159 / (0.1 - (v5 - v5))))
                                switch (fi13 << ((18446744073709551615) & 63)) {
                                case -209:
                                    if false {
                                        var v31 string = "a"
                                        _ = v31
                                    } else {
                                        switch "a" {
                                        case "":
                                            v32 := v4[42]
                                            _ = v32
                                        case "foo bar":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            _ = len(v2)
                                        }
                                        v30 += float32((0.0 * 0.1))
                                    }
                                    v33 := GenS1{F0: 42, F1: 44, F2: 2}
                                    fmt.Println(v30)
                                    _ = v33
                                case 32767:
                                    v34 := GenS3{F0: 1, F1: "foo bar", F2: 42}
                                    var v35 uint = 255
                                    _ = v34
                                    _ = v35
                                case 1:
                                    v2 = append(v2, "foo bar")
                                default:
                                    var v36 uint = 42
                                    _ = v36
                                }
                                _ = v30
                            }
                        }
                        var v37 int32 = 42
                        var v38 int8 = -1
                        _ = v37
                        _ = v38
                    }
                }
                v39 := GenS1{F0: 18446744073709551615, F1: 127, F2: 128}
                var v40 int32 = 502
                var v41 uint = 1
                _ = v39
                _ = v40
                _ = v41
            }
            _ = fi14
        }
        _ = fi13
    }
    var v42 int64 = (-(v1))
    switch 2147483647 {
    case -9223372036854775808:
        v4[42] = (true && false)
    case 128:
        var v43 float64 = (v5 / v5)
        _ = v43
    default:
        fmt.Println(v42, v42)
    }
    fmt.Println(v1, v1)
    var v44 float64 = (1.5 * v5)
    v45 := GenS4{F0: 42, F1: 2147483648}
    v3 = append(v3, 2)
    v46 := GenS5{F0: -710}
    v5 = 3.14159
    fmt.Println(v42, v1, v5)
    if ((-(v1)) == (v1 + 0)) {
        switch 42 {
        case 1:
            if false {
                var v47 int64 = 9223372036854775807
                if false {
                    for fi48 := 0; fi48 < 12; fi48++ {
                        if (!true) {
                            v49 := GenS2{F0: (fi48 % ((fi48) | 1))}
                            v50 := GenS5{F0: -32768}
                            if (!true) {
                                var v51 float32 = float32(v44)
                                v52 := v4[626]
                                if true {
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                }
                                _ = v51
                                _ = v52
                            }
                            _ = v49
                            _ = v50
                        } else {
                            fmt.Println(v42)
                        }
                        v5 += ((v5 / v5) / v44)
                        switch (^(fi48)) {
                        case 32768:
                            v1 += v1
                        case -1:
                            switch (-(fi48)) {
                            case 2147483647:
                                if (0.0 < (3.14159 * 1.5)) {
                                    var v53 uint = 194
                                    _ = v53
                                }
                            }
                        }
                        _ = fi48
                    }
                    v54 := GenS3{F0: 186, F1: "a", F2: -168}
                    _ = v54
                } else {
                    for fi55 := 0; fi55 < 20; fi55++ {
                        switch (fi55 << ((4294967295) & 63)) {
                        case -886:
                            fmt.Println(v42, v5)
                        case -358:
                            v56 := GenS5{F0: 0}
                            _ = v56
                        case 9223372036854775807:
                            _ = len(v4)
                        }
                        var v57 float64 = v44
                        if ("foo bar" <= ("a" + (("\n\t" + "a") + "a"))) {
                            var v58 int64 = -1
                            _ = v58
                        } else {
                            v59 := GenS6{F0: true}
                            _ = v59
                        }
                        _ = fi55
                        _ = v57
                    }
                    if (-545 != 319) {
                        fmt.Println(v47)
                    }
                    var v60 uint32 = 2147483647
                    _ = v60
                }
                for fi61 := 0; fi61 < 9; fi61++ {
                    var v62 int64 = (v42 >> ((2) & 63))
                    switch fi61 {
                    case -152:
                        for fi63 := 0; fi63 < 5; fi63++ {
                            v47 -= (v47 ^ (v62 >> ((18446744073709551615) & 63)))
                            var v64 bool = (false || false)
                            var v65 int8 = -128
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                    case 737:
                        v1 += (v47 >> ((107) & 63))
                    default:
                        _ = len(v4)
                    }
                    _ = fi61
                    _ = v62
                }
                _ = v47
            }
        case 128:
            v44 = v44
        case 32768:
            v66 := v4[42]
            _ = v66
        default:
            v42 -= v1
        }
        var v67 int16 = 178
        _ = v67
    }
    for fi68 := 0; fi68 < 19; fi68++ {
        fmt.Println(fi68)
        _ = fi68
    }
    var v69 float32 = ((1e10 * -2.25) * 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v69
}
