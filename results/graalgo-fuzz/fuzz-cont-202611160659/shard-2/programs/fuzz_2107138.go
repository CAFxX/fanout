package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 string; F1 uint; F2 int32 }
type GenS6 struct { F0 int8; F1 int32; F2 int64 }
type GenS7 struct { F0 int8; F1 int64; F2 int8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16 }
func genFn0(p0 uint) uint16 {
    if (0 < 89) {
        var v6 int16 = -1
        v7 := []int{4294967295, 1}
        if true {
            v6 = -724
            var v8 int32 = 0
            v6 -= (-(v6))
            _ = v8
        } else {
            var v9 int32 = -830
            p0 -= (p0 << (((^(p0))) & 63))
            _ = len(v7)
            _ = v9
        }
        _ = v6
        _ = v7
    }
    v10 := GenS1{F0: 65535, F1: 0}
    _ = p0
    _ = v10
    return 42
}

func main() {
    v1 := []uint32{}
    var v2 bool = ((3.14159 != 0.0) && (!true))
    var v3 float64 = float64(0.1)
    v4 := make(map[uint32]bool)
    var v5 uint16 = 42
    v4[591] = v2
    var v11 int64 = 0
    v12 := GenS1{F0: 4294967295, F1: -96}
    v11 = (v11 - (v11 - -9223372036854775808))
    v3 = float64(3.14159)
    v13 := v4[0]
    switch (((("\n\t" + (("hello" + "foo bar") + "foo bar")) + ((("hello" + "a") + "") + "foo bar")) + ("a" + "a")) + "a") {
    case "a":
        v2 = (true && ((v11 < v11) || (-556 == 0)))
    case "foo bar":
        if (("foo bar" + "") > (("hello" + "a") + (("\n\t" + "a") + ("hello" + "\n\t")))) {
            fmt.Println(v13)
        }
    case "":
        _ = len(v4)
    default:
        v14 := GenS2{F0: -613, F1: -1}
        _ = v14
    }
    if ((1e10 + -2.25) <= (((float32((((0.0 + v3) * (float64(1e10) + float64(((1.5 + 3.14159) / ((((0.0 * -2.25) + ((0.0 + (-2.25 * float32(v3))) - float32((v3 + v3)))) / 1.5) - (float32((((v3 * v3) / (v3 * 1e10)) + v3)) + 0.0)))))) * v3)) * float32((float64((0.1 * 1e10)) + v3))) + 3.14159) / 1.5)) {
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            for fi15 := 0; fi15 < 2; fi15++ {
                for fi16 := 0; fi16 < 11; fi16++ {
                    if v13 {
                        switch fi15 {
                        case 9223372036854775807:
                            var v17 uint64 = 42
                            _ = v17
                        }
                        v3 = v3
                        fmt.Println(fi15)
                    } else {
                        v2 = v13
                    }
                    delete(v4, 1)
                    fmt.Println(fi15, fi16, v2)
                    _ = fi16
                }
                if (("foo bar" > "\n\t") || false) {
                    switch (fi15 - int(v5)) {
                    case -2147483648:
                        _ = len(v1)
                    }
                    v18 := GenS3{F0: 1}
                    fmt.Println(v3, v2)
                    _ = v18
                }
                _ = fi15
            }
        default:
            for fi19 := 0; fi19 < 5; fi19++ {
                switch "foo bar" {
                case "hello":
                    for fi20 := 0; fi20 < 18; fi20++ {
                        fmt.Println(v3)
                        switch ("a" + ("hello" + (("\n\t" + ("foo bar" + "\n\t")) + "hello"))) {
                        case "\n\t":
                            switch ("a" + (("a" + "") + "hello")) {
                            case "":
                                v21 := v4[136]
                                _ = v21
                            case "a":
                                v22 := GenS2{F0: -2147483648, F1: 888}
                                _ = v22
                            case "hello":
                                fmt.Println(fi20, v13)
                            }
                        case "hello":
                            v3 -= v3
                        case "foo bar":
                            v23 := GenS2{F0: -1, F1: 42}
                            _ = v23
                        default:
                            var v24 uint64 = 514
                            _ = v24
                        }
                        fmt.Println(fi19, fi19)
                        _ = fi20
                    }
                case "foo bar":
                    v2 = v2
                default:
                    for fi25 := 0; fi25 < 11; fi25++ {
                        v26 := v4[uint32(v5)]
                        _ = fi25
                        _ = v26
                    }
                }
                _ = fi19
            }
        }
        if v13 {
            switch "hello" {
            case "hello":
                for fi27 := 0; fi27 < 19; fi27++ {
                    var v28 int = (^(fi27))
                    v13 = (!(42 > v5))
                    var v29 bool = false
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            case "\n\t":
                v30 := GenS2{F0: -430, F1: 1}
                _ = v30
            case "a":
                v2 = (v13 || (32767 <= int16(v5)))
            default:
                v31 := v4[664]
                _ = v31
            }
            v11 += -1
            v3 = 1.5
        }
    }
    if v2 {
        if v2 {
            v3 += 1e10
        } else {
            v2 = (v2 && (v3 >= 1.5))
            if false {
                for fi32 := 0; fi32 < 15; fi32++ {
                    for fi33 := 0; fi33 < 2; fi33++ {
                        v34 := v4[2]
                        fmt.Println(v3)
                        _ = fi33
                        _ = v34
                    }
                    if (0.0 != 3.14159) {
                        for fi35 := 0; fi35 < 2; fi35++ {
                            v36 := GenS3{F0: 2}
                            _ = fi35
                            _ = v36
                        }
                        fmt.Println(v3, v11)
                    }
                    v4[1] = v13
                    _ = fi32
                }
            } else {
                if (0 == 1) {
                    v1 = append(v1, 464)
                    for fi37 := 0; fi37 < 4; fi37++ {
                        v38 := GenS3{F0: 0}
                        for fi39 := 0; fi39 < 11; fi39++ {
                            v40 := GenS2{F0: 1, F1: -1}
                            v4[65535] = (int(v5) != (fi37 * fi37))
                            v13 = true
                            _ = fi39
                            _ = v40
                        }
                        _ = fi37
                        _ = v38
                    }
                    for fi41 := 0; fi41 < 13; fi41++ {
                        if (255 < 255) {
                            if v13 {
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    v43 := GenS2{F0: 535, F1: -28}
                                    _ = fi42
                                    _ = v43
                                }
                                var v44 string = "foo bar"
                                _ = v44
                            } else {
                                var v45 uint16 = (^(v5))
                                var v46 int8 = 69
                                _ = v45
                                _ = v46
                            }
                            for fi47 := 0; fi47 < 19; fi47++ {
                                if v13 {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v48 uint16 = v5
                                    _ = v48
                                } else {
                                    switch ("foo bar" + "\n\t") {
                                    case "foo bar":
                                        var v49 float32 = 0.1
                                        _ = v49
                                    }
                                }
                                switch ("a" + ("foo bar" + "hello")) {
                                case "":
                                    v13 = true
                                case "\n\t":
                                    v3 += (0.1 * 0.1)
                                default:
                                    if (!v13) {
                                        if (1 > 0) {
                                            if ((fi47 << ((255) & 63)) != -97) {
                                                switch "hello" {
                                                case "\n\t":
                                                    v50 := GenS2{F0: -1, F1: -237}
                                                    _ = v50
                                                case "hello":
                                                    if true {
                                                        _ = len(v1)
                                                        if ((!((1.5 - -2.25) == 0.1)) && (v13 && false)) {
                                                            var v51 uint = 586
                                                            v52 := GenS4{F0: -109, F1: 127}
                                                            if false {
                                                                v13 = v2
                                                                v11 = v11
                                                                if (false && v2) {
                                                                    switch (fi47 % ((fi41) | 1)) {
                                                                    case 555:
                                                                        fmt.Println(v3, v13, fi47)
                                                                    case 32768:
                                                                        _ = len(v1)
                                                                    }
                                                                } else {
                                                                    var v53 uint16 = (^(v5))
                                                                    _ = v53
                                                                }
                                                                _ = len(v4)
                                                            }
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        switch ("hello" + "\n\t") {
                                                        case "":
                                                            if (58 > 1) {
                                                                var v54 int64 = v11
                                                                _ = v54
                                                            }
                                                        case "a":
                                                            var v55 float64 = (0.0 * (float64(0.0) * (v3 * (-2.25 / ((v3 - (((1e10 - 0.0) + 1e10) / float64(1.5))) / v3)))))
                                                            _ = v55
                                                        case "hello":
                                                            if (!(806 <= (v5 + (v5 | 255)))) {
                                                                v11 = v11
                                                                v4[2147483647] = true
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    _ = len(v1)
                                                default:
                                                    v2 = false
                                                }
                                            } else {
                                                v13 = false
                                                var v56 float32 = (0.1 * 1.5)
                                                var v57 uint8 = 128
                                                _ = v56
                                                _ = v57
                                            }
                                            v58 := GenS5{F0: "", F1: 42, F2: -2147483648}
                                            _ = v58
                                        }
                                        v59 := v4[65535]
                                        var v60 int16 = -1
                                        _ = v59
                                        _ = v60
                                    }
                                }
                                if false {
                                    var v61 int = 32768
                                    switch "\n\t" {
                                    case "":
                                        var v62 int8 = 1
                                        _ = v62
                                    case "foo bar":
                                        v63 := GenS6{F0: 0, F1: 0, F2: v11}
                                        _ = v63
                                    default:
                                        fmt.Println(v11)
                                    }
                                    _ = v61
                                } else {
                                    if false {
                                        var v64 int64 = v11
                                        var v65 uint16 = v5
                                        if len(v1) > 0 {
                                            v1[0] = 0
                                        }
                                        _ = v64
                                        _ = v65
                                    }
                                }
                                var v66 float32 = 3.14159
                                v67 := GenS7{F0: 42, F1: v11, F2: 0}
                                v4[4294967295] = (!(v13 || true))
                                if false {
                                    if v13 {
                                        switch fi41 {
                                        case -448:
                                            v4[493] = true
                                        }
                                        v3 -= 1e10
                                        v5 = (v5 * (v5 & (v5 - v5)))
                                    } else {
                                        v68 := GenS8{F0: 209}
                                        var v69 uint8 = 255
                                        v70 := GenS8{F0: 65535}
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                    }
                                    v71 := GenS9{F0: genFn0(4294967295)}
                                    _ = v71
                                } else {
                                    if (!(!false)) {
                                        var v72 int = (-(fi47))
                                        if len(v1) > 0 {
                                            v1[0] = 316
                                        }
                                        if (128 == 42) {
                                            if len(v1) > 0 {
                                                v1[0] = 42
                                            }
                                        }
                                        _ = v72
                                    } else {
                                        v2 = v13
                                        v73 := v4[752]
                                        v11 -= v11
                                        _ = v73
                                    }
                                }
                                _ = fi47
                                _ = v66
                                _ = v67
                            }
                        }
                        if true {
                            if true {
                                v5 += v5
                                var v74 uint32 = 2
                                for fi75 := 0; fi75 < 7; fi75++ {
                                    v5 += v5
                                    _ = fi75
                                }
                                _ = v74
                            }
                        }
                        _ = fi41
                    }
                } else {
                    var v76 int8 = 0
                    _ = v76
                }
                fmt.Println(v5)
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
