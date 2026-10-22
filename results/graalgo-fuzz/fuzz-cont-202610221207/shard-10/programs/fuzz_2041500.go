package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint16; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS7 struct { F0 uint; F1 uint32; F2 int8 }
func genFn0(p0 bool) int {
    var v6 string = "a"
    v6 = (v6 + v6)
    if p0 {
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: 1, F1: 675}
            _ = fi7
            _ = v8
        }
        v6 = "\n\t"
        v9 := GenS1{F0: 42, F1: 292}
        _ = v9
    }
    switch v6 {
    case "hello":
        for fi10 := 0; fi10 < 14; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := GenS2{F0: 219, F1: 845, F2: 319}
                _ = fi11
                _ = v12
            }
            v6 = (v6 + ("foo bar" + ""))
            if true {
                var v13 int = -345
                _ = v13
            } else {
                switch "hello" {
                case "foo bar":
                    if (!p0) {
                        v6 = v6
                        if (!false) {
                            v6 = v6
                            for fi14 := 0; fi14 < 15; fi14++ {
                                var v15 int32 = 866
                                v16 := GenS3{F0: p0, F1: 356, F2: 879}
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            v17 := GenS1{F0: -218, F1: 42}
                            _ = v17
                        } else {
                            v18 := []int{128, -825, fi10, (fi10 ^ (fi10 + (fi10 & fi10)))}
                            _ = v18
                        }
                        v19 := GenS2{F0: 128, F1: 255, F2: 918}
                        _ = v19
                    }
                }
            }
            _ = fi10
        }
    default:
        v6 = ""
    }
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 4294967295
}
func genFn1(p0 int, p1 uint64, p2 uint8) int8 {
    if ((p1 / (((p1 - p1)) | 1)) >= (p1 >> ((378) & 63))) {
        v20 := []int{p0}
        _ = v20
    } else {
        if false {
            p0 = 127
            if (false || (!false)) {
                p2 -= 42
            }
        }
        fmt.Println(p1, p2)
        p1 += (p1 % ((p1) | 1))
    }
    v21 := []int{(p0 - genFn0((!false))), p0, 9223372036854775807}
    var v22 uint64 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return 100
}

func main() {
    var v1 int64 = 15
    var v2 float64 = (-2.25 * 1.5)
    v3 := []string{"\n\t", "", "foo bar"}
    v4 := make(map[int]int)
    var v5 int32 = 151
    if false {
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 int16 = -1
            _ = fi23
            _ = v24
        }
        switch "a" {
        case "":
            var v25 string = "a"
            _ = v25
        default:
            switch 847 {
            case 1:
                v26 := v4[genFn0(true)]
                _ = v26
            default:
                v5 += -924
            }
        }
    } else {
        var v27 uint32 = 48
        _ = len(v4)
        _ = v27
    }
    switch 32768 {
    case -114:
        if true {
            for fi28 := 0; fi28 < 14; fi28++ {
                var v29 uint16 = 231
                _ = fi28
                _ = v29
            }
            v1 += 1
        } else {
            var v30 int16 = -768
            delete(v4, -1)
            _ = v30
        }
    case 127:
        delete(v4, 548)
    default:
        v31 := GenS4{F0: "\n\t"}
        _ = v31
    }
    v32 := GenS1{F0: 1, F1: 32767}
    var v33 int64 = (v1 | (v1 * (v1 - 0)))
    switch "hello" {
    case "foo bar":
        v33 = 1
    }
    v33 -= -1
    switch 783 {
    case -127:
        if true {
            if len(v3) > 0 {
                v3[0] = (("\n\t" + "foo bar") + "\n\t")
            }
            v34 := v4[402]
            _ = v34
        } else {
            var v35 float64 = v2
            _ = v35
        }
    }
    v36 := GenS1{F0: -1, F1: 677}
    v37 := GenS5{F0: ("\n\t" + "a"), F1: -891}
    v1 = v1
    var v38 float32 = (float32(v2) * ((float32(((v2 * 1.5) / (float64(float32(v2)) - 1.5))) / (0.0 - -2.25)) + 1.5))
    if (!true) {
        fmt.Println(v33, v33, v2)
        v3 = append(v3, "\n\t")
        for fi39 := 0; fi39 < 6; fi39++ {
            if true {
                v40 := GenS1{F0: v5, F1: 255}
                v38 += (v38 / 1.0)
                var v41 int16 = 0
                _ = v40
                _ = v41
            }
            v4[(fi39 << ((99) & 63))] = fi39
            _ = fi39
        }
    } else {
        v42 := GenS3{F0: (36 <= 0), F1: 2, F2: -32768}
        v1 = v1
        _ = v42
    }
    switch -59 {
    case -735:
        var v43 int16 = 32767
        _ = v43
    case 0:
        switch int(v1) {
        case 32767:
            var v44 int32 = -147
            _ = v44
        }
    case 32767:
        v4[42] = -175
    default:
        var v45 uint32 = 2147483648
        _ = v45
    }
    switch "a" {
    case "\n\t":
        switch "a" {
        case "":
            if false {
                if true {
                    var v46 int16 = -560
                    var v47 float64 = float64(0.1)
                    _ = v46
                    _ = v47
                } else {
                    var v48 uint32 = 2147483648
                    v5 = 2147483647
                    switch int(v5) {
                    case 127:
                        delete(v4, 32767)
                    default:
                        v2 = (v2 + (v2 / v2))
                    }
                    _ = v48
                }
            } else {
                if (!true) {
                    if (v5 > v5) {
                        var v49 uint8 = 175
                        _ = v49
                    }
                }
                v50 := GenS6{F0: 32767, F1: 42, F2: 0}
                var v51 string = ("\n\t" + "")
                _ = v50
                _ = v51
            }
        case "foo bar":
            if true {
                switch 2147483647 {
                case 1:
                    for fi52 := 0; fi52 < 1; fi52++ {
                        var v53 uint8 = 0
                        v54 := GenS2{F0: 42, F1: 2, F2: 2147483648}
                        var v55 uint32 = 2
                        _ = fi52
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                }
                v33 += (-(v1))
            }
        default:
            _ = len(v4)
        }
    case "hello":
        v38 -= 1e10
    }
    var v56 uint8 = 255
    for fi57 := 0; fi57 < 20; fi57++ {
        v58 := v4[fi57]
        if (1 >= 18446744073709551615) {
            v56 = v56
            if true {
                v59 := GenS7{F0: 255, F1: 637, F2: 42}
                fmt.Println(v2, v38)
                v58 -= fi57
                _ = v59
            }
        } else {
            if true {
                for fi60 := 0; fi60 < 17; fi60++ {
                    v61 := GenS1{F0: (^(v5)), F1: 0}
                    var v62 uint = 684
                    _ = fi60
                    _ = v61
                    _ = v62
                }
                v4[32767] = (fi57 ^ v58)
            } else {
                for fi63 := 0; fi63 < 15; fi63++ {
                    v64 := v4[fi63]
                    _ = fi63
                    _ = v64
                }
                v1 = (v33 + (v1 | -9223372036854775808))
                fmt.Println(v2, v2)
            }
        }
        _ = fi57
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v56
}
