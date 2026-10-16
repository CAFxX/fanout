package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 string; F1 int64; F2 uint32 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64; F1 uint; F2 string }
func genFn0(p0 uint64, p1 uint64) uint {
    p1 += (p1 / ((1) | 1))
    v7 := GenS1{F0: 17, F1: uint8(p0)}
    switch 32768 {
    case 0:
        p1 += (p1 * 962)
    case 1:
        if false {
            for fi8 := 0; fi8 < 2; fi8++ {
                for fi9 := 0; fi9 < 19; fi9++ {
                    p1 += 4294967295
                    fmt.Println(p0, fi9, p1)
                    p0 += 0
                    _ = fi9
                }
                p0 += p1
                _ = fi8
            }
        } else {
            p0 -= p1
            if (!true) {
                var v10 int8 = 117
                v11 := []int{-2147483648, 127}
                switch "foo bar" {
                case "foo bar":
                    var v12 string = ("hello" + ("\n\t" + ("hello" + "foo bar")))
                    _ = v12
                case "hello":
                    for fi13 := 0; fi13 < 10; fi13++ {
                        for fi14 := 0; fi14 < 10; fi14++ {
                            if (-584 >= 42) {
                                v15 := GenS2{F0: p0, F1: (fi14 + fi14), F2: (v10 / (((v10 & v10)) | 1))}
                                if true {
                                    _ = len(v11)
                                    v16 := GenS2{F0: p1, F1: fi13, F2: (v10 / ((v10) | 1))}
                                    _ = v16
                                } else {
                                    var v17 float64 = (1e10 * ((1.5 * float64(-2.25)) + (1e10 * 3.14159)))
                                    p1 = 2
                                    if true {
                                        v10 += (v10 << ((585) & 63))
                                        p0 = 141
                                    }
                                    _ = v17
                                }
                                v18 := GenS2{F0: 1, F1: -9223372036854775808, F2: (-(v10))}
                                fmt.Println(p1, v10, fi14)
                                _ = v15
                                _ = v18
                            } else {
                                v11 = append(v11, fi13)
                                if len(v11) > 0 {
                                    v11[0] = fi14
                                }
                                p1 = (p1 >> ((255) & 63))
                            }
                            var v19 float32 = ((0.0 / 1e10) * (((float32(0.1) - (3.14159 / (((float32(float64(float32(((0.1 + (1.5 / 0.1)) * 1e10)))) * 0.0) + -2.25) + (0.1 - 1e10)))) / 1.5) * 0.0))
                            _ = fi14
                            _ = v19
                        }
                        _ = fi13
                    }
                }
                _ = v10
                _ = v11
            }
            var v20 string = "a"
            _ = v20
        }
    case 200:
        p1 -= p0
    default:
        p0 -= p0
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := []int{fi21}
        fmt.Println(fi21, p1)
        var v23 uint32 = 426
        _ = fi21
        _ = v22
        _ = v23
    }
    p1 = (p1 | (p1 | p1))
    _ = p0
    _ = p1
    _ = v7
    return 4294967295
}

func main() {
    v1 := []string{(("hello" + "foo bar") + "\n\t"), ((("hello" + "\n\t") + ("foo bar" + "hello")) + ("\n\t" + "a")), "a", ""}
    v2 := []int{223, 128, 465}
    var v3 uint8 = 127
    v4 := make(map[uint32]bool)
    v5 := make(map[string]int)
    v6 := []uint32{707, 478, 42, 1}
    if (1 >= 0) {
        if true {
            if (false || false) {
                v5[(("\n\t" + "foo bar") + ("foo bar" + ((("a" + "\n\t") + "\n\t") + "hello")))] = 128
                var v24 uint16 = 42
                _ = v24
            } else {
                switch 4294967295 {
                case -2147483648:
                    v25 := GenS3{F0: v3, F1: 348}
                    _ = v25
                default:
                    for fi26 := 0; fi26 < 12; fi26++ {
                        for fi27 := 0; fi27 < 11; fi27++ {
                            if true {
                                var v28 float32 = (3.14159 * (1.5 - (3.14159 / 1.5)))
                                v29 := v5["a"]
                                _ = v28
                                _ = v29
                            }
                            if (0.1 != (1.5 / (3.14159 * 1.5))) {
                                var v30 int = (fi26 ^ (fi26 << ((65535) & 63)))
                                var v31 uint16 = 255
                                _ = v30
                                _ = v31
                            } else {
                                var v32 uint8 = (v3 % (((^(v3))) | 1))
                                v32 += (v32 - (v3 * v3))
                                v33 := GenS4{F0: 42}
                                _ = v32
                                _ = v33
                            }
                            v2 = append(v2, (fi27 << ((2) & 63)))
                            _ = fi27
                        }
                        _ = fi26
                    }
                }
                v3 = 42
            }
            v1 = append(v1, "foo bar")
            delete(v5, "hello")
        } else {
            v34 := GenS2{F0: 723, F1: 32768, F2: 95}
            _ = v34
        }
    } else {
        fmt.Println(v3)
    }
    v3 -= (v3 * v3)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    for fi35 := 0; fi35 < 6; fi35++ {
        switch "" {
        case "hello":
            if ((!(!false)) && false) {
                _ = len(v1)
            }
        }
        _ = fi35
    }
    if false {
        v36 := GenS3{F0: v3, F1: 1}
        switch "\n\t" {
        case "foo bar":
            for fi37 := 0; fi37 < 15; fi37++ {
                var v38 float32 = -2.25
                v3 = v3
                var v39 int64 = 1
                _ = fi37
                _ = v38
                _ = v39
            }
        case "":
            var v40 float64 = 3.14159
            _ = v40
        default:
            v3 = (v3 << ((2) & 63))
        }
        v3 -= v3
        _ = v36
    } else {
        if (!false) {
            for fi41 := 0; fi41 < 6; fi41++ {
                switch "" {
                case "a":
                    var v42 int8 = 127
                    _ = v42
                default:
                    v43 := v5[""]
                    _ = v43
                }
                _ = fi41
            }
            v44 := GenS5{F0: 65535}
            _ = v44
        } else {
            v3 -= v3
            for fi45 := 0; fi45 < 9; fi45++ {
                v3 -= 2
                v46 := GenS1{F0: 0, F1: (v3 << ((65535) & 63))}
                _ = fi45
                _ = v46
            }
        }
    }
    for fi47 := 0; fi47 < 16; fi47++ {
        switch "a" {
        case "\n\t":
            v3 -= v3
        default:
            if false {
                v3 += 56
            } else {
                for fi48 := 0; fi48 < 11; fi48++ {
                    fmt.Println(fi47, v3, fi47)
                    var v49 float32 = 1e10
                    v3 -= (^(v3))
                    _ = fi48
                    _ = v49
                }
                v3 = 2
                v50 := GenS6{F0: "foo bar", F1: -1, F2: 0}
                _ = v50
            }
        }
        _ = fi47
    }
    v51 := GenS7{F0: (v3 | v3)}
    if (1.5 < ((float32(0.1) / (1.5 * (1.5 + 1.5))) + (1e10 * (0.0 - float32(float64(1.5)))))) {
        delete(v5, (("a" + "foo bar") + ""))
    }
    v6 = append(v6, 4294967295)
    var v52 uint8 = v3
    v53 := GenS8{F0: -1, F1: 255, F2: ""}
    if false {
        var v54 int32 = 42
        _ = v54
    }
    v52 -= v52
    v55 := GenS3{F0: (v3 | v52), F1: 2}
    v52 = (v52 * (v3 % (((v3 | (v3 | v52))) | 1)))
    var v56 bool = (true || false)
    v57 := GenS4{F0: 0}
    v52 = (v3 - v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v52
    _ = v53
    _ = v55
    _ = v56
    _ = v57
    _ = v6
}
