package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 bool; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 string; F1 int; F2 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int8; F1 int16; F2 int8 }
func genFn0(p0 uint64, p1 int32) bool {
    var v7 uint16 = 1
    for fi8 := 0; fi8 < 13; fi8++ {
        p0 = p0
        _ = fi8
    }
    var v9 uint8 = 1
    v9 = v9
    v10 := GenS1{F0: p0, F1: 42}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v9
    return (!(("foo bar" + ("" + "hello")) > ("\n\t" + "")))
}
func genFn1(p0 uint8) int8 {
    p0 = 0
    switch 0 {
    case 32768:
        v11 := GenS1{F0: 4294967295, F1: -169}
        _ = v11
    case 128:
        var v12 string = "a"
        _ = v12
    default:
        p0 = p0
    }
    p0 -= 1
    _ = p0
    return 100
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 uint64 = 18446744073709551615
    var v4 float32 = 3.14159
    var v5 uint64 = (^(v3))
    var v6 uint32 = 1
    v3 += (v3 ^ 0)
    v13 := GenS2{F0: (false || false), F1: -1, F2: 0}
    switch -141 {
    case 32768:
        if ((v5 >> ((1) & 63)) == 18446744073709551615) {
            v14 := GenS2{F0: (int32(v3) < 1), F1: 9223372036854775807, F2: 2}
            _ = v14
        }
    case 128:
        var v15 uint = 42
        _ = v15
    }
    var v16 uint32 = (v6 - v6)
    v17 := []int{2147483647}
    v16 -= 900
    v18 := GenS1{F0: v3, F1: 0}
    switch ("hello" + "hello") {
    case "":
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := v2[("\n\t" + "a")]
            var v21 uint64 = v5
            _ = fi19
            _ = v20
            _ = v21
        }
    case "foo bar":
        for fi22 := 0; fi22 < 1; fi22++ {
            v16 += 1
            var v23 int = (^(fi22))
            _ = fi22
            _ = v23
        }
    }
    var v24 uint64 = (v5 | v5)
    switch (((("hello" + "a") + "\n\t") + "a") + "hello") {
    case "hello":
        var v25 bool = (!(v4 != float32(1e10)))
        _ = v25
    case "":
        if (!false) {
            if (1e10 < (-2.25 * -2.25)) {
                if true {
                    fmt.Println(v5)
                }
                v26 := GenS3{F0: 681, F1: 0}
                _ = v26
            }
            for fi27 := 0; fi27 < 7; fi27++ {
                delete(v2, "\n\t")
                _ = fi27
            }
            for fi28 := 0; fi28 < 7; fi28++ {
                v17 = append(v17, (fi28 ^ (^(fi28))))
                _ = fi28
            }
        } else {
            for fi29 := 0; fi29 < 16; fi29++ {
                if (!("\n\t" < (("" + ("" + "\n\t")) + ""))) {
                    var v30 uint = 42
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 11; fi31++ {
                        v2["a"] = "foo bar"
                        _ = fi31
                    }
                }
                _ = fi29
            }
        }
    case "\n\t":
        var v32 int32 = 42
        _ = v32
    default:
        var v33 float64 = -2.25
        _ = v33
    }
    v4 = 1e10
    switch "hello" {
    case "foo bar":
        if (((float64(-2.25) > (float64(0.1) / 1e10)) || true) && false) {
            _ = len(v1)
            switch "hello" {
            case "foo bar":
                for fi34 := 0; fi34 < 6; fi34++ {
                    v1[fi34] = (v6 ^ 830)
                    var v35 float32 = (float32((float64((v4 / (((v4 + ((v4 * (v4 + (float32((0.0 + (0.1 + 0.0))) + float32(float64(1e10))))) / 1.0)) + v4) / 1e10))) / -2.25)) * float32(0.0))
                    var v36 uint32 = v6
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case "\n\t":
                var v37 int64 = 1
                _ = v37
            }
            _ = len(v17)
        } else {
            var v38 int = -552
            v4 = v4
            if (!true) {
                if (false || (-91 > -128)) {
                    if true {
                        v39 := GenS3{F0: 40, F1: v16}
                        _ = v39
                    } else {
                        fmt.Println(v6)
                        for fi40 := 0; fi40 < 18; fi40++ {
                            v3 += v5
                            delete(v1, fi40)
                            _ = fi40
                        }
                        if (655 >= 704) {
                            switch v38 {
                            case -683:
                                if false {
                                    if len(v17) > 0 {
                                        v17[0] = int(v5)
                                    }
                                }
                            case -2147483648:
                                var v41 uint32 = 1
                                _ = v41
                            case -9223372036854775808:
                                if (v38 == v38) {
                                    fmt.Println(v16, v3)
                                }
                            }
                            var v42 int = (v38 % (((v38 | (v38 ^ -574))) | 1))
                            for fi43 := 0; fi43 < 20; fi43++ {
                                if false {
                                    v44 := GenS3{F0: 42, F1: (v6 / (((v16 & (v6 >> ((42) & 63)))) | 1))}
                                    v3 = v5
                                    v45 := v1[(fi43 * v38)]
                                    _ = v44
                                    _ = v45
                                } else {
                                    var v46 int8 = -128
                                    v47 := GenS2{F0: true, F1: 42, F2: 200}
                                    _ = v46
                                    _ = v47
                                }
                                v2["foo bar"] = "\n\t"
                                var v48 int8 = 0
                                _ = fi43
                                _ = v48
                            }
                            _ = v42
                        }
                    }
                } else {
                    v6 -= 2147483648
                    fmt.Println(v6)
                }
                if (!false) {
                    switch int(v3) {
                    case -1:
                        for fi49 := 0; fi49 < 17; fi49++ {
                            switch 127 {
                            case 42:
                                var v50 int64 = -9223372036854775808
                                _ = v50
                            case 32767:
                                if (true || false) {
                                    if (false && ("\n\t" != "hello")) {
                                        for fi51 := 0; fi51 < 17; fi51++ {
                                            if ("\n\t" <= ("foo bar" + "")) {
                                                fmt.Println(v4, v6)
                                                v6 = (v6 + (v6 % (((^(v16))) | 1)))
                                                switch "hello" {
                                                case "hello":
                                                    if true {
                                                        v52 := GenS1{F0: 42, F1: v38}
                                                        delete(v2, "hello")
                                                        switch (v38 ^ fi51) {
                                                        case 126:
                                                            _ = len(v17)
                                                        case 2147483647:
                                                            v16 -= v6
                                                        case 4294967295:
                                                            var v53 int = 127
                                                            _ = v53
                                                        default:
                                                            v24 -= 2
                                                        }
                                                        _ = v52
                                                    }
                                                default:
                                                    switch "\n\t" {
                                                    case "a":
                                                        v54 := GenS3{F0: 65535, F1: 2147483648}
                                                        _ = v54
                                                    case "\n\t":
                                                        if true {
                                                            if true {
                                                                if len(v17) > 0 {
                                                                    v17[0] = -9223372036854775808
                                                                }
                                                                v55 := GenS1{F0: 352, F1: v38}
                                                                _ = v55
                                                            }
                                                            _ = len(v17)
                                                        }
                                                    default:
                                                        fmt.Println(v24)
                                                    }
                                                    v5 += v3
                                                    switch (fi51 << ((42) & 63)) {
                                                    case 42:
                                                        _ = len(v1)
                                                        v56 := GenS4{F0: ("\n\t" + ("" + "foo bar")), F1: fi51, F2: 108}
                                                        v16 = (v16 ^ v16)
                                                        _ = v56
                                                    }
                                                }
                                            }
                                            v24 = (v24 << ((260) & 63))
                                            if false {
                                                switch "a" {
                                                case "\n\t":
                                                    if len(v17) > 0 {
                                                        v17[0] = (fi49 - 1)
                                                    }
                                                default:
                                                    v57 := GenS5{F0: 2}
                                                    _ = v57
                                                }
                                                v58 := GenS2{F0: ((1 < v16) && (!(v4 <= float32((0.1 * float64((-2.25 - 1.5))))))), F1: 1, F2: 0}
                                                v59 := GenS6{F0: -1, F1: 820, F2: 32}
                                                _ = v58
                                                _ = v59
                                            }
                                            _ = fi51
                                        }
                                        fmt.Println(v3)
                                    }
                                }
                            case 32768:
                                switch v38 {
                                case 482:
                                    var v60 float64 = float64(float32(0.0))
                                    _ = v60
                                }
                            default:
                                v38 = (v38 * fi49)
                            }
                            _ = fi49
                        }
                    }
                    v38 = 42
                }
            } else {
                for fi61 := 0; fi61 < 2; fi61++ {
                    for fi62 := 0; fi62 < 14; fi62++ {
                        v63 := v2[("\n\t" + "")]
                        v16 = 42
                        _ = fi62
                        _ = v63
                    }
                    if ((v38 ^ (fi61 << ((18446744073709551615) & 63))) <= v38) {
                        if (genFn1(1) <= 127) {
                            switch v38 {
                            case 714:
                                switch v38 {
                                case 496:
                                    if (687 <= 0) {
                                        for fi64 := 0; fi64 < 14; fi64++ {
                                            fmt.Println(v5, v16)
                                            _ = fi64
                                        }
                                    } else {
                                        v65 := GenS5{F0: 1}
                                        v24 = 1
                                        for fi66 := 0; fi66 < 5; fi66++ {
                                            v4 = (v4 - (v4 - (v4 - v4)))
                                            var v67 uint64 = 978
                                            v38 = fi61
                                            _ = fi66
                                            _ = v67
                                        }
                                        _ = v65
                                    }
                                case -2147483648:
                                    fmt.Println(fi61, v3)
                                case 531:
                                    fmt.Println(v16, fi61, v4)
                                }
                            case 32767:
                                var v68 uint64 = (v24 * (v3 / ((544) | 1)))
                                _ = v68
                            case 9223372036854775807:
                                v3 += (v24 * 18446744073709551615)
                            default:
                                v17 = append(v17, (v38 / ((v38) | 1)))
                            }
                            v6 -= (^(v6))
                        }
                    }
                    _ = fi61
                }
            }
            _ = v38
        }
    case "a":
        var v69 uint8 = 255
        _ = v69
    case "":
        v2["hello"] = "hello"
    default:
        switch 128 {
        case 127:
            if false {
                var v70 uint8 = 127
                v71 := GenS1{F0: (v5 / ((332) | 1)), F1: 550}
                _ = v70
                _ = v71
            } else {
                fmt.Println(v6)
                var v72 uint64 = (v24 / ((v5) | 1))
                var v73 uint64 = (v72 % ((v5) | 1))
                _ = v72
                _ = v73
            }
        case 0:
            for fi74 := 0; fi74 < 15; fi74++ {
                v75 := GenS4{F0: "", F1: fi74, F2: v6}
                var v76 int32 = 42
                v3 -= (v3 & uint64(v16))
                _ = fi74
                _ = v75
                _ = v76
            }
        case 148:
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
        }
    }
    var v77 uint = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v77
}
