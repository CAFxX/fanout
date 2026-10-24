package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int8; F1 string; F2 int }
func genFn0(p0 int32, p1 int, p2 uint16) int16 {
    var v6 uint = 0
    if false {
        fmt.Println(v6)
        v6 = v6
    } else {
        if false {
            for fi7 := 0; fi7 < 4; fi7++ {
                v6 = (^(v6))
                p2 = (p2 + (p2 % ((p2) | 1)))
                p0 -= 1
                _ = fi7
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 0
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 0
    v3 := make(map[int]uint32)
    var v4 uint8 = 60
    var v5 int16 = v2
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 bool = false
        v2 -= (v2 | genFn0(-2147483648, -1, 2))
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 42
    for fi11 := 0; fi11 < 13; fi11++ {
        if true {
            v10 += 0
        }
        if false {
            var v12 int64 = 42
            for fi13 := 0; fi13 < 20; fi13++ {
                if true {
                    if false {
                        v1[4294967295] = 4294967295
                    } else {
                        v2 -= (v2 + (v2 % ((genFn0(0, int(v12), 0)) | 1)))
                    }
                }
                v12 = (-(v12))
                fmt.Println(v4)
                _ = fi13
            }
            _ = v12
        }
        var v14 uint64 = 4294967295
        _ = fi11
        _ = v14
    }
    v15 := GenS1{F0: 1, F1: 989}
    v16 := GenS2{F0: 42}
    if ((766 < 2) || ((1.5 - 1e10) >= (0.1 * float32((1.5 - (3.14159 - 0.0)))))) {
        fmt.Println(v2, v5, v10)
        fmt.Println(v2, v2, v5)
    } else {
        var v17 float64 = 3.14159
        _ = v17
    }
    var v18 float64 = 3.14159
    v19 := []int{128, -104}
    for fi20 := 0; fi20 < 1; fi20++ {
        v21 := GenS1{F0: 2, F1: 581}
        v22 := v1[(fi20 + -2147483648)]
        fmt.Println(v5, v4)
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint = 2
    var v24 int64 = 0
    switch "hello" {
    case "foo bar":
        if true {
            switch -1 {
            case -9223372036854775808:
                if true {
                    v25 := GenS2{F0: int8(v23)}
                    v19 = append(v19, 100)
                    _ = v25
                }
            case 4294967295:
                switch "\n\t" {
                case "foo bar":
                    switch 32767 {
                    case 0:
                        var v26 int64 = v24
                        _ = v26
                    case -2147483648:
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                    default:
                        for fi27 := 0; fi27 < 4; fi27++ {
                            v28 := v1[(fi27 - fi27)]
                            _ = fi27
                            _ = v28
                        }
                    }
                case "hello":
                    var v29 int8 = 0
                    _ = v29
                default:
                    if (4294967295 > uint64(v4)) {
                        switch (("\n\t" + "hello") + "\n\t") {
                        case "a":
                            if true {
                                switch "hello" {
                                case "a":
                                    var v30 bool = ((false && (!(!(((("a" + "foo bar") + "") + "foo bar") == "hello")))) && true)
                                    _ = v30
                                case "hello":
                                    v23 = 65535
                                case "":
                                    _ = len(v19)
                                default:
                                    fmt.Println(v4)
                                }
                                if true {
                                    if false {
                                        var v31 float32 = 3.14159
                                        _ = v31
                                    }
                                    var v32 uint64 = 925
                                    _ = v32
                                }
                                var v33 int8 = -1
                                _ = v33
                            }
                        default:
                            fmt.Println(v5, v18, v2)
                        }
                        v2 = 299
                        v34 := GenS3{F0: (v4 ^ (v10 ^ (v4 << (((^(v23))) & 63)))), F1: 836}
                        _ = v34
                    }
                }
            default:
                if (100 < -128) {
                    if false {
                        for fi35 := 0; fi35 < 9; fi35++ {
                            v5 = 0
                            for fi36 := 0; fi36 < 17; fi36++ {
                                if true {
                                    fmt.Println(v4)
                                }
                                _ = len(v1)
                                _ = fi36
                            }
                            _ = fi35
                        }
                        v4 = (^(v10))
                    } else {
                        var v37 uint64 = 2
                        _ = v37
                    }
                }
            }
            v38 := GenS3{F0: v10, F1: 2}
            _ = v38
        }
    default:
        for fi39 := 0; fi39 < 5; fi39++ {
            if (2147483647 != 42) {
                v5 += (^(v5))
                switch (("\n\t" + "foo bar") + (((("foo bar" + (((("foo bar" + (("\n\t" + "") + ("a" + "\n\t"))) + "foo bar") + "\n\t") + ("" + ("\n\t" + (("hello" + ("hello" + "foo bar")) + (("a" + ((("" + (("a" + ("hello" + ("" + "foo bar"))) + (("hello" + "foo bar") + "\n\t"))) + (("foo bar" + "") + "a")) + "a")) + "foo bar")))))) + (("foo bar" + "foo bar") + (((("hello" + ("\n\t" + ("" + "\n\t"))) + "hello") + "hello") + "a"))) + ("hello" + "")) + "")) {
                case "":
                    delete(v1, fi39)
                }
            } else {
                v3[-33] = 2
                switch (fi39 * fi39) {
                case -1:
                    v40 := GenS1{F0: 32767, F1: 1}
                    _ = v40
                }
                _ = len(v19)
            }
            for fi41 := 0; fi41 < 8; fi41++ {
                var v42 int32 = 42
                v43 := GenS3{F0: (v10 >> (((v23 ^ v23)) & 63)), F1: 1}
                v24 -= v24
                if ((v42 | (v42 % ((v42) | 1))) <= 638) {
                    if true {
                        v44 := GenS4{F0: int16(v24), F1: 434, F2: 0}
                        v45 := v1[fi41]
                        var v46 uint = v23
                        _ = v44
                        _ = v45
                        _ = v46
                    }
                    v19 = append(v19, (fi39 ^ fi41))
                    v4 += (v10 % ((v10) | 1))
                }
                _ = fi41
                _ = v42
                _ = v43
            }
            if false {
                switch "hello" {
                case "foo bar":
                    var v47 float32 = ((1.5 / (float32((((1.5 + float64(3.14159)) - v18) + float64(float32(((float64(((3.14159 / 0.1) / float32(v18))) / (v18 + v18)) / v18))))) / (1.5 / 1.0))) * -2.25)
                    _ = v47
                case "\n\t":
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                default:
                    if ((v24 - v24) < int64(v23)) {
                        var v48 int64 = (v24 | (v24 >> ((v23) & 63)))
                        if len(v19) > 0 {
                            v19[0] = 0
                        }
                        _ = v48
                    }
                }
                v24 -= 9223372036854775807
            }
            _ = fi39
        }
    }
    if ((^(v23)) != (v23 - v23)) {
        var v49 float64 = (float64((1.5 + (0.1 - 1e10))) + (((((v18 / v18) / (v18 / 1e10)) / ((float64((1e10 + (3.14159 + float32(float64(float32(v18)))))) / (v18 - (float64(0.0) + 0.1))) + v18)) - (v18 - ((v18 - float64(1.5)) / 0.1))) / (1.5 - v18)))
        _ = v49
    }
    if true {
        if false {
            switch 32767 {
            case 0:
                v4 -= 1
            }
            var v50 uint32 = 176
            v10 += (v4 % ((v4) | 1))
            _ = v50
        } else {
            v51 := GenS5{F0: 773}
            _ = v51
        }
    } else {
        switch -189 {
        case 9223372036854775807:
            v4 = (^(v10))
        case 10:
            switch 128 {
            case -9223372036854775808:
                if len(v19) > 0 {
                    v19[0] = -974
                }
            case 32767:
                var v52 uint = (^(v23))
                _ = v52
            case 4294967295:
                var v53 int8 = -128
                _ = v53
            }
        case 4294967295:
            if (v5 <= genFn0(2147483647, 128, 382)) {
                if true {
                    v54 := GenS2{F0: 0}
                    _ = v54
                } else {
                    var v55 string = "hello"
                    _ = v55
                }
                v2 = v5
            } else {
                if (4294967295 >= 798) {
                    for fi56 := 0; fi56 < 1; fi56++ {
                        for fi57 := 0; fi57 < 19; fi57++ {
                            var v58 uint32 = 718
                            v18 = 1.5
                            _ = fi57
                            _ = v58
                        }
                        var v59 int8 = 0
                        _ = fi56
                        _ = v59
                    }
                    v60 := GenS6{F0: 100, F1: ((("a" + "hello") + ("" + "a")) + "a"), F2: 32768}
                    _ = v60
                }
                v61 := GenS6{F0: 0, F1: "a", F2: 1}
                if false {
                    v3[32767] = 2147483647
                } else {
                    v23 = 65535
                }
                _ = v61
            }
        default:
            for fi62 := 0; fi62 < 13; fi62++ {
                v3[(fi62 | fi62)] = 953
                _ = fi62
            }
        }
    }
    fmt.Println(v5, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
