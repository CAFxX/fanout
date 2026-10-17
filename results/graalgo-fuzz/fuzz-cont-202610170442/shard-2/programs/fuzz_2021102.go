package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint }
type GenS3 struct { F0 string; F1 int8; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 uint; F2 uint32 }
type GenS6 struct { F0 uint16; F1 uint32; F2 uint }
type GenS7 struct { F0 uint16; F1 string; F2 int64 }
func genFn0(p0 uint8, p1 bool) uint32 {
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint16 = 255
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "hello":
        if (2 != 4294967295) {
            var v9 int8 = 42
            _ = v9
        }
    case "":
        var v10 uint16 = 32767
        _ = v10
    }
    v11 := []int{120, -700, 1}
    _ = p0
    _ = p1
    _ = v11
    return 2147483647
}
func genFn1(p0 uint32, p1 int8) uint8 {
    if false {
        var v12 float64 = float64((0.1 * (float32(0.0) * (3.14159 / 3.14159))))
        var v13 uint32 = p0
        fmt.Println(v13, v13, v13)
        _ = v12
        _ = v13
    }
    if true {
        p1 += 53
        if (!(!("a" >= ""))) {
            v14 := []int{68}
            _ = v14
        }
        fmt.Println(p1, p1)
    } else {
        for fi15 := 0; fi15 < 15; fi15++ {
            v16 := []int{fi15, (fi15 | (-(fi15))), fi15, (fi15 + (fi15 * (fi15 % (((fi15 % (((fi15 - 42)) | 1))) | 1))))}
            if true {
                if len(v16) > 0 {
                    v16[0] = fi15
                }
            } else {
                for fi17 := 0; fi17 < 7; fi17++ {
                    p0 -= genFn0(0, true)
                    if ((p0 | p0) >= p0) {
                        if false {
                            _ = len(v16)
                            if (false && false) {
                                p1 = p1
                                var v18 uint = 0
                                p0 = (p0 & p0)
                                _ = v18
                            }
                            var v19 int32 = -771
                            _ = v19
                        }
                    }
                    switch fi15 {
                    case 4294967295:
                        p1 = p1
                    case 127:
                        fmt.Println(fi17, fi17)
                    case 476:
                        p1 -= (p1 >> ((2) & 63))
                    default:
                        switch "\n\t" {
                        case "":
                            switch ("\n\t" + "\n\t") {
                            case "hello":
                                v16 = append(v16, (fi17 & fi15))
                            }
                        case "foo bar":
                            p0 -= p0
                        case "hello":
                            if (float64((float32(1.5) + -2.25)) == 0.1) {
                                p1 = -1
                                var v20 bool = (false || (703 < -2))
                                _ = v20
                            } else {
                                p0 += 44
                                p1 += p1
                                p1 = 1
                            }
                        }
                    }
                    _ = fi17
                }
            }
            _ = fi15
            _ = v16
        }
        var v21 string = "a"
        _ = v21
    }
    v22 := GenS1{F0: 784}
    _ = p0
    _ = p1
    _ = v22
    return 1
}

func main() {
    var v1 string = ("hello" + "a")
    v2 := []int{9223372036854775807, -277}
    v3 := make(map[string]uint32)
    var v4 uint8 = 255
    v5 := make(map[string]bool)
    var v6 float64 = 1.5
    var v23 uint = 661
    var v24 string = "a"
    v24 = v24
    v25 := GenS2{F0: v4, F1: 1, F2: v23}
    v1 = (v24 + v24)
    if true {
        v26 := GenS1{F0: -1}
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 int64 = 242
            if (!(32767 < 32768)) {
                fmt.Println(v23, v23)
                if (true && false) {
                    v29 := GenS1{F0: -2147483648}
                    v6 -= v6
                    v28 -= -1
                    _ = v29
                }
                var v30 int = fi27
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        if false {
            var v31 string = v24
            _ = v31
        } else {
            v3["foo bar"] = genFn0(v4, true)
            switch -2147483648 {
            case 0:
                var v32 float64 = ((0.0 * 1.5) * -2.25)
                _ = v32
            case 32767:
                for fi33 := 0; fi33 < 6; fi33++ {
                    if (fi33 <= 680) {
                        var v34 uint32 = 4294967295
                        if true {
                            v23 += (v23 >> ((v23) & 63))
                        }
                        v24 = v1
                        _ = v34
                    } else {
                        fmt.Println(v1, fi33, v6)
                        v2 = append(v2, fi33)
                        delete(v5, v1)
                    }
                    v24 = (((((v1 + v24) + "a") + v24) + v24) + (("foo bar" + v1) + v24))
                    _ = fi33
                }
            default:
                var v35 uint16 = 676
                _ = v35
            }
        }
        _ = v26
    } else {
        v36 := GenS1{F0: 554}
        v1 = v1
        for fi37 := 0; fi37 < 10; fi37++ {
            if (845 == genFn0((v4 ^ (^(v4))), false)) {
                var v38 int = (fi37 & (fi37 | (fi37 << (((v23 * v23)) & 63))))
                v2 = append(v2, -304)
                v38 = -2147483648
                _ = v38
            } else {
                var v39 int8 = 42
                v40 := GenS1{F0: 1}
                _ = v39
                _ = v40
            }
            v5[v1] = false
            _ = fi37
        }
        _ = v36
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        v42 := GenS2{F0: 128, F1: 829, F2: v23}
        v6 -= float64(1e10)
        v43 := GenS3{F0: "\n\t", F1: -1, F2: 1}
        _ = fi41
        _ = v42
        _ = v43
    }
    v23 -= v23
    if true {
        for fi44 := 0; fi44 < 4; fi44++ {
            if false {
                var v45 uint16 = 765
                v46 := GenS4{F0: (fi44 >> (((v23 >> (((v23 & v23)) & 63))) & 63))}
                _ = len(v2)
                _ = v45
                _ = v46
            } else {
                if true {
                    v1 = "a"
                }
                v47 := GenS1{F0: -1}
                switch ("" + ((v1 + v24) + ((v24 + (((v1 + "") + "\n\t") + (v1 + "hello"))) + v1))) {
                case "hello":
                    fmt.Println(v4)
                case "\n\t":
                    for fi48 := 0; fi48 < 13; fi48++ {
                        fmt.Println(v6, v1)
                        fmt.Println(fi44)
                        switch fi44 {
                        case 247:
                            v24 = v1
                        }
                        _ = fi48
                    }
                case "":
                    v49 := GenS5{F0: 4294967295, F1: 335, F2: 65535}
                    _ = v49
                default:
                    for fi50 := 0; fi50 < 16; fi50++ {
                        v24 = (("a" + v1) + v1)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if true {
                            if false {
                                v51 := GenS6{F0: 1, F1: 2147483647, F2: (v23 - (v23 ^ (v23 >> ((v23) & 63))))}
                                fmt.Println(v24, fi44, fi50)
                                fmt.Println(v1)
                                var v52 int64 = 253
                                var v53 int = -1
                                _ = v51
                                _ = v52
                                _ = v53
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                fmt.Println(v4)
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if len(v2) > 0 {
                                v2[0] = 4294967295
                            }
                        } else {
                            v54 := GenS5{F0: 4294967295, F1: v23, F2: 42}
                            v23 += v23
                            _ = v54
                        }
                        _ = fi50
                    }
                }
                _ = v47
            }
            for fi55 := 0; fi55 < 10; fi55++ {
                var v56 int64 = 900
                _ = fi55
                _ = v56
            }
            _ = fi44
        }
        if false {
            if len(v2) > 0 {
                v2[0] = 545
            }
        }
    } else {
        v23 += 2
    }
    v1 = v24
    if ((!(32768 > 538)) || (-920 <= -9223372036854775808)) {
        for fi57 := 0; fi57 < 18; fi57++ {
            var v58 uint8 = (v4 << (((v23 / ((v23) | 1))) & 63))
            _ = fi57
            _ = v58
        }
        v23 = 255
    } else {
        fmt.Println(v23, v23, v1)
        v2 = append(v2, 32768)
        for fi59 := 0; fi59 < 14; fi59++ {
            fmt.Println(fi59, v4)
            var v60 int64 = 42
            var v61 int = fi59
            _ = fi59
            _ = v60
            _ = v61
        }
    }
    var v62 uint32 = 1
    if (!(genFn1(179, 41) != v4)) {
        fmt.Println(v23, v23)
    } else {
        switch ((v1 + (v1 + v24)) + v24) {
        case "a":
            v6 -= v6
        }
        fmt.Println(v62, v6, v23)
        for fi63 := 0; fi63 < 18; fi63++ {
            for fi64 := 0; fi64 < 11; fi64++ {
                v65 := GenS7{F0: 32767, F1: (((v1 + (("a" + (v1 + "\n\t")) + (("foo bar" + ((v24 + v1) + "hello")) + "foo bar"))) + v24) + v1), F2: 0}
                _ = fi64
                _ = v65
            }
            var v66 uint64 = 412
            v6 = (v6 / v6)
            _ = fi63
            _ = v66
        }
    }
    var v67 string = "a"
    var v68 int = -737
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v67
    _ = v68
}
