package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 uint; F1 int32; F2 int32 }
type GenS5 struct { F0 uint32; F1 string; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int) uint8 {
    var v5 uint16 = 270
    if false {
        switch (("hello" + ("a" + ("\n\t" + ("hello" + ("\n\t" + "foo bar"))))) + ("" + "a")) {
        case "foo bar":
            if false {
                var v6 int64 = 140
                _ = v6
            }
        case "hello":
            var v7 string = (("hello" + "hello") + "\n\t")
            _ = v7
        case "a":
            for fi8 := 0; fi8 < 13; fi8++ {
                switch (p0 & (fi8 + (p0 ^ p0))) {
                case 4294967295:
                    if true {
                        if false {
                            fmt.Println(p0, v5)
                            v5 += (v5 % ((42) | 1))
                            p0 = (fi8 / ((4294967295) | 1))
                        } else {
                            if (!true) {
                                v9 := GenS1{F0: 178}
                                v10 := []int{(fi8 + fi8), p0, p0}
                                if false {
                                    fmt.Println(p0)
                                }
                                _ = v10
                                _ = v9
                            } else {
                                switch "" {
                                case "\n\t":
                                    switch (fi8 & (fi8 & fi8)) {
                                    case 0:
                                        var v11 uint16 = v5
                                        _ = v11
                                    case 32768:
                                        var v12 int64 = 0
                                        _ = v12
                                    case 128:
                                        var v13 int = (p0 & (fi8 / ((fi8) | 1)))
                                        _ = v13
                                    }
                                case "hello":
                                    p0 = fi8
                                case "":
                                    var v14 bool = false
                                    _ = v14
                                default:
                                    v5 = v5
                                }
                            }
                        }
                        v15 := GenS2{F0: 93, F1: (true || (true && (!false)))}
                        v5 = v5
                        _ = v15
                    } else {
                        v16 := []int{fi8}
                        _ = v16
                    }
                }
                switch "a" {
                case "":
                    var v17 uint8 = 255
                    _ = v17
                }
                _ = fi8
            }
        default:
            fmt.Println(p0, v5)
        }
    } else {
        var v18 float64 = 0.0
        _ = v18
    }
    var v19 uint = 18446744073709551615
    p0 -= 0
    _ = p0
    _ = v19
    _ = v5
    return uint8(v19)
}

func main() {
    v1 := []int{128, 1, 9223372036854775807, 4294967295}
    var v2 uint32 = 70
    v3 := make(map[int]bool)
    var v4 uint8 = 1
    v2 += (v2 | (v2 | (v2 % ((782) | 1))))
    v20 := GenS3{F0: 42, F1: 4294967295}
    for fi21 := 0; fi21 < 5; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            switch "hello" {
            case "a":
                fmt.Println(v2, fi22, fi22)
            case "foo bar":
                switch "a" {
                case "":
                    var v23 int64 = -9223372036854775808
                    _ = v23
                }
            case "hello":
                switch ("a" + "foo bar") {
                case "foo bar":
                    var v24 int32 = 42
                    _ = v24
                case "":
                    fmt.Println(fi22, fi22)
                default:
                    var v25 int32 = -418
                    _ = v25
                }
            default:
                v2 += 0
            }
            _ = fi22
        }
        _ = fi21
    }
    if false {
        v1 = append(v1, 127)
        for fi26 := 0; fi26 < 17; fi26++ {
            v4 += (v4 * v4)
            _ = fi26
        }
        v4 = genFn0(1)
    } else {
        v4 = genFn0(127)
    }
    v3[664] = false
    v4 -= genFn0(4294967295)
    _ = len(v1)
    var v27 uint32 = (v2 ^ (v2 + 2147483648))
    if ((42 != -340) && (32767 >= 42)) {
        if ("a" < "hello") {
            if false {
                v28 := GenS4{F0: 2, F1: 2147483647, F2: -788}
                _ = v28
            } else {
                fmt.Println(v27, v2)
            }
            fmt.Println(v2)
            switch ("foo bar" + "foo bar") {
            case "a":
                for fi29 := 0; fi29 < 12; fi29++ {
                    switch fi29 {
                    case 394:
                        if false {
                            var v30 float32 = -2.25
                            var v31 uint = 610
                            _ = v30
                            _ = v31
                        } else {
                            for fi32 := 0; fi32 < 17; fi32++ {
                                fmt.Println(v2, v4)
                                v33 := GenS1{F0: v4}
                                _ = fi32
                                _ = v33
                            }
                            for fi34 := 0; fi34 < 15; fi34++ {
                                if false {
                                    if (false && true) {
                                        fmt.Println(fi29)
                                        var v35 int32 = 0
                                        _ = v35
                                    }
                                    switch fi29 {
                                    case 995:
                                        switch fi29 {
                                        case 127:
                                            var v36 uint64 = 4294967295
                                            _ = v36
                                        default:
                                            v1 = append(v1, fi29)
                                        }
                                    case -873:
                                        fmt.Println(fi29, fi29)
                                    default:
                                        switch (fi29 & int(v2)) {
                                        case -2147483648:
                                            v2 = 2147483647
                                        case 0:
                                            if true {
                                                v37 := GenS3{F0: -1, F1: (v2 * 2)}
                                                _ = v37
                                            } else {
                                                var v38 int64 = -763
                                                v39 := v3[(fi34 + -935)]
                                                _ = v38
                                                _ = v39
                                            }
                                        case 543:
                                            switch fi29 {
                                            case -2147483648:
                                                v27 += 42
                                            case 32767:
                                                if false {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    switch "hello" {
                                                    case "hello":
                                                        fmt.Println(fi34, v2)
                                                    case "":
                                                        v2 = (^(v2))
                                                    }
                                                    v4 += genFn0((fi29 + (fi34 / ((fi29) | 1))))
                                                    v2 += v27
                                                } else {
                                                    var v40 uint8 = v4
                                                    _ = v40
                                                }
                                            default:
                                                switch (fi34 / ((-1) | 1)) {
                                                case -9223372036854775808:
                                                    fmt.Println(fi34)
                                                case 32768:
                                                    var v41 uint32 = (v2 * v27)
                                                    _ = v41
                                                default:
                                                    v2 += v27
                                                }
                                            }
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if false {
                                                v4 -= v4
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v4 = (v4 + (v4 * 255))
                                            }
                                        }
                                    }
                                    if false {
                                        v3[fi29] = true
                                        var v42 int8 = 127
                                        v43 := v3[127]
                                        if (!("hello" == ((("a" + "foo bar") + "foo bar") + "foo bar"))) {
                                            if ((42 < 1000) && ((genFn0((-(fi34))) != uint8(v27)) && (!(((18446744073709551615 <= 135) || (uint16(v27) > 772)) && v43)))) {
                                                if v43 {
                                                    v43 = v43
                                                } else {
                                                    var v44 int = (fi29 << ((382) & 63))
                                                    _ = v44
                                                }
                                                var v45 int16 = 1
                                                v3[(fi34 % ((127) | 1))] = (0.1 > 1e10)
                                                _ = v45
                                            }
                                            switch ("a" + "\n\t") {
                                            case "hello":
                                                switch "" {
                                                case "\n\t":
                                                    _ = len(v1)
                                                }
                                            }
                                        }
                                        var v46 float32 = (0.1 / 1.0)
                                        _ = v42
                                        _ = v43
                                        _ = v46
                                    } else {
                                        v4 = v4
                                    }
                                } else {
                                    v27 = v27
                                }
                                if true {
                                    v4 = v4
                                    if (!true) {
                                        delete(v3, fi29)
                                    } else {
                                        v27 += (^(v27))
                                    }
                                } else {
                                    if false {
                                        switch ("hello" + (("a" + ("hello" + "a")) + "hello")) {
                                        case "hello":
                                            var v47 uint16 = 255
                                            _ = v47
                                        case "foo bar":
                                            fmt.Println(v2, v2)
                                        case "\n\t":
                                            if false {
                                                var v48 uint64 = 435
                                                _ = v48
                                            }
                                        default:
                                            var v49 int = 1
                                            _ = v49
                                        }
                                        if (2147483647 > -2147483648) {
                                            v2 += 65535
                                            if true {
                                                if (965 >= uint64(v2)) {
                                                    v50 := GenS5{F0: (v2 | v2), F1: ("a" + "a"), F2: (^(fi34))}
                                                    _ = v50
                                                }
                                            }
                                            _ = len(v3)
                                        } else {
                                            fmt.Println(fi34, v4)
                                            v51 := GenS5{F0: (v27 ^ (^(v27))), F1: "\n\t", F2: fi34}
                                            var v52 string = "a"
                                            _ = v51
                                            _ = v52
                                        }
                                    }
                                }
                                _ = fi34
                            }
                        }
                    default:
                        _ = len(v1)
                    }
                    switch fi29 {
                    case 4294967295:
                        v53 := GenS6{F0: fi29}
                        _ = v53
                    case 0:
                        v54 := GenS7{F0: 42, F1: (v2 ^ (v2 % (((^(v27))) | 1)))}
                        _ = v54
                    default:
                        fmt.Println(fi29)
                    }
                    _ = fi29
                }
            case "foo bar":
                if false {
                    fmt.Println(v4, v2, v4)
                    for fi55 := 0; fi55 < 14; fi55++ {
                        v2 -= (v2 << ((1) & 63))
                        switch fi55 {
                        case -76:
                            v56 := GenS7{F0: 42, F1: 388}
                            _ = v56
                        case -2147483648:
                            v4 -= v4
                        default:
                            delete(v3, -184)
                        }
                        delete(v3, 0)
                        _ = fi55
                    }
                }
            default:
                v2 += v2
            }
        } else {
            var v57 float64 = 3.14159
            var v58 uint = 18446744073709551615
            if (!true) {
                if false {
                    v2 -= (^(v2))
                }
                var v59 float32 = 0.1
                _ = v59
            }
            _ = v57
            _ = v58
        }
    }
    v4 = genFn0(int(v2))
    var v60 uint64 = 873
    for fi61 := 0; fi61 < 10; fi61++ {
        if true {
            var v62 int32 = -503
            _ = v62
        }
        v60 = 0
        _ = fi61
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v60
}
