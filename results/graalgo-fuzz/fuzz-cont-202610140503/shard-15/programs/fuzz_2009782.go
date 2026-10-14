package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int32; F2 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 string; F2 string }
type GenS6 struct { F0 uint; F1 int16; F2 int64 }
type GenS7 struct { F0 int16 }
func genFn0(p0 int64) int16 {
    p0 += (p0 | 0)
    v5 := []int{536, -9223372036854775808, -9223372036854775808}
    p0 = p0
    p0 -= p0
    _ = p0
    _ = v5
    return 921
}
func genFn1(p0 string, p1 uint8, p2 bool) bool {
    switch p0 {
    case "":
        if p2 {
            switch -1 {
            case 9223372036854775807:
                var v6 int = 459
                _ = v6
            case 604:
                var v7 uint8 = p1
                _ = v7
            case 32767:
                p1 = p1
            }
            p2 = true
            p1 = 128
        } else {
            fmt.Println(p0, p1)
            if p2 {
                v8 := GenS1{F0: "hello", F1: (p1 + (^(p1))), F2: 42}
                var v9 int64 = 1
                var v10 uint64 = 0
                _ = v10
                _ = v8
                _ = v9
            }
        }
    }
    v11 := []int{-655, -481}
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            p1 = 127
            switch fi12 {
            case -20:
                if p2 {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    switch fi12 {
                    case 2147483647:
                        var v14 int8 = int8(p1)
                        _ = v14
                    case 106:
                        if false {
                            if len(v11) > 0 {
                                v11[0] = (fi13 + fi13)
                            }
                            if (-99 != -170) {
                                v15 := GenS2{F0: 2, F1: 0, F2: 2147483647}
                                p2 = (!(true && (!(0 > 32767))))
                                fmt.Println(fi12)
                                var v16 bool = ("a" == ((p0 + (p0 + (p0 + (p0 + "a")))) + (p0 + ("foo bar" + p0))))
                                var v17 float64 = (0.1 - -2.25)
                                _ = v15
                                _ = v16
                                _ = v17
                            } else {
                                var v18 int16 = -1
                                _ = v18
                            }
                        } else {
                            switch p0 {
                            case "foo bar":
                                p0 = p0
                            case "\n\t":
                                v19 := GenS1{F0: "\n\t", F1: (p1 * (p1 << ((980) & 63))), F2: 769}
                                _ = v19
                            default:
                                if ((65535 < 1) || p2) {
                                    var v20 uint32 = 42
                                    fmt.Println(v20)
                                    _ = v20
                                }
                            }
                            if len(v11) > 0 {
                                v11[0] = fi13
                            }
                            v21 := GenS1{F0: ("foo bar" + p0), F1: (p1 + 70), F2: 42}
                            _ = v21
                        }
                    }
                    p0 = p0
                    var v22 bool = ((!p2) && p2)
                    _ = v22
                }
            }
            _ = len(v11)
            _ = fi13
        }
        if p2 {
            switch (fi12 << ((319) & 63)) {
            case -9223372036854775808:
                if (p0 > ((((p0 + p0) + p0) + "\n\t") + p0)) {
                    v23 := GenS1{F0: p0, F1: 0, F2: 255}
                    fmt.Println(p0)
                    _ = v23
                }
            case -859:
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            case 4294967295:
                switch "foo bar" {
                case "\n\t":
                    var v24 int = fi12
                    _ = v24
                case "a":
                    p0 = (p0 + (p0 + p0))
                default:
                    var v25 uint32 = 93
                    _ = v25
                }
            }
        }
        if p2 {
            for fi26 := 0; fi26 < 19; fi26++ {
                if len(v11) > 0 {
                    v11[0] = fi26
                }
                switch "" {
                case "":
                    if len(v11) > 0 {
                        v11[0] = fi12
                    }
                }
                _ = fi26
            }
            if ((p2 || p2) || true) {
                v27 := GenS2{F0: 0, F1: -2147483648, F2: 330}
                for fi28 := 0; fi28 < 14; fi28++ {
                    var v29 float32 = float32(1.5)
                    _ = fi28
                    _ = v29
                }
                _ = v27
            } else {
                v11 = append(v11, (fi12 ^ int(p1)))
                v30 := GenS3{F0: 65535}
                _ = v30
            }
            for fi31 := 0; fi31 < 1; fi31++ {
                _ = len(v11)
                if len(v11) > 0 {
                    v11[0] = fi12
                }
                _ = fi31
            }
        }
        _ = fi12
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        for fi33 := 0; fi33 < 3; fi33++ {
            var v34 bool = (-2.25 < (0.0 + 1e10))
            _ = fi33
            _ = v34
        }
        for fi35 := 0; fi35 < 12; fi35++ {
            fmt.Println(fi32)
            v36 := GenS1{F0: p0, F1: (p1 % ((p1) | 1)), F2: 32767}
            _ = fi35
            _ = v36
        }
        _ = fi32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return p2
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 string = "a"
    var v4 int64 = 9223372036854775807
    v37 := GenS1{F0: v3, F1: 128, F2: 337}
    switch 32767 {
    case -528:
        v38 := GenS2{F0: 2147483647, F1: 2147483647, F2: 0}
        _ = v38
    default:
        if (false || (!false)) {
            if (-88 == 270) {
                for fi39 := 0; fi39 < 17; fi39++ {
                    var v40 uint16 = 255
                    v4 = v4
                    if (!(4294967295 > 1)) {
                        v41 := GenS4{F0: v4}
                        _ = v41
                    }
                    _ = fi39
                    _ = v40
                }
                var v42 uint64 = 2
                fmt.Println(v3, v3)
                _ = v42
            }
            v3 = (v3 + (v3 + v3))
        } else {
            v4 -= (v4 - (^(v4)))
            for fi43 := 0; fi43 < 9; fi43++ {
                v3 = v3
                v44 := []int{(fi43 - 32768), fi43}
                v3 = ""
                _ = fi43
                _ = v44
            }
            v45 := GenS3{F0: 42}
            _ = v45
        }
    }
    switch -196 {
    case 0:
        for fi46 := 0; fi46 < 11; fi46++ {
            fmt.Println(v3, v3)
            _ = fi46
        }
    default:
        v4 = v4
    }
    switch v3 {
    case "\n\t":
        switch "foo bar" {
        case "hello":
            if ((true || (false || false)) && true) {
                var v47 int32 = 2147483647
                _ = v47
            } else {
                v3 = v3
                v3 = (("\n\t" + v3) + "")
            }
        case "":
            v4 += (v4 | (v4 & (v4 + (v4 - v4))))
        case "a":
            switch 331 {
            case -9223372036854775808:
                v48 := GenS1{F0: v3, F1: 157, F2: 91}
                _ = v48
            case 471:
                for fi49 := 0; fi49 < 7; fi49++ {
                    for fi50 := 0; fi50 < 6; fi50++ {
                        fmt.Println(fi49)
                        _ = fi50
                    }
                    var v51 int64 = v4
                    _ = fi49
                    _ = v51
                }
            case 32767:
                var v52 int32 = 42
                _ = v52
            default:
                v1[42] = false
            }
        }
    case "hello":
        var v53 uint32 = 2147483648
        _ = v53
    default:
        var v54 string = (v3 + "a")
        _ = v54
    }
    for fi55 := 0; fi55 < 20; fi55++ {
        v4 = -1
        v56 := GenS2{F0: 631, F1: 0, F2: 42}
        if true {
            v3 = ("foo bar" + v3)
        }
        _ = fi55
        _ = v56
    }
    for fi57 := 0; fi57 < 13; fi57++ {
        v4 += (v4 % ((v4) | 1))
        fmt.Println(v4, v4, v4)
        var v58 int64 = -78
        _ = fi57
        _ = v58
    }
    v3 = ""
    switch 868 {
    case 4294967295:
        v3 = ""
    case -9223372036854775808:
        var v59 int64 = v4
        _ = v59
    default:
        fmt.Println(v4, v4)
    }
    v60 := GenS2{F0: 4294967295, F1: -2147483648, F2: -627}
    switch 42 {
    case 239:
        fmt.Println(v4, v4)
    case 32767:
        var v61 uint = 416
        _ = v61
    default:
        if false {
            fmt.Println(v3, v3)
            for fi62 := 0; fi62 < 1; fi62++ {
                v3 = v3
                if false {
                    var v63 uint64 = 18446744073709551615
                    v64 := GenS5{F0: (fi62 != fi62), F1: v3, F2: "hello"}
                    _ = v63
                    _ = v64
                } else {
                    if (0 > 42) {
                        v65 := GenS6{F0: 255, F1: genFn0((v4 ^ (^(v4)))), F2: v4}
                        delete(v1, fi62)
                        _ = v65
                    } else {
                        fmt.Println(fi62, v3)
                    }
                    if (96 > 870) {
                        v66 := GenS7{F0: 42}
                        _ = v66
                    }
                    if (1e10 > -2.25) {
                        v1[fi62] = true
                        v67 := []int{fi62, (fi62 + 4294967295), fi62, (^(fi62))}
                        _ = v67
                    } else {
                        var v68 uint16 = 65535
                        _ = v68
                    }
                }
                _ = fi62
            }
        } else {
            for fi69 := 0; fi69 < 19; fi69++ {
                var v70 int = fi69
                _ = fi69
                _ = v70
            }
            v4 = (v4 * v4)
        }
    }
    if (-2.25 == (float64(0.1) * 1e10)) {
        var v71 uint8 = 210
        delete(v2, 1)
        fmt.Println(v71, v3)
        _ = v71
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v60
}
