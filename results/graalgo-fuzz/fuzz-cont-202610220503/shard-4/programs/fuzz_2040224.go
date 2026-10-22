package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 bool }
type GenS2 struct { F0 uint32; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16; F1 int64; F2 int32 }
type GenS4 struct { F0 bool; F1 uint; F2 uint32 }
func genFn0(p0 uint64, p1 string) uint8 {
    var v5 uint = 390
    var v6 float64 = 1e10
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{}
    var v3 uint8 = 127
    var v4 int16 = int16(v3)
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v3, v1)
        v8 := GenS1{F0: 625, F1: ("a" + ("a" + "\n\t")), F2: false}
        v9 := GenS2{F0: 214, F1: 0, F2: false}
        _ = fi7
        _ = v8
        _ = v9
    }
    switch "hello" {
    case "foo bar":
        var v10 uint8 = (^(v3))
        _ = v10
    default:
        v1 += (v1 | -204)
    }
    var v11 uint16 = 2
    switch ("\n\t" + ("" + "foo bar")) {
    case "hello":
        v4 = 431
    }
    if ((0 < 41) || (("foo bar" + ("\n\t" + (("foo bar" + (("hello" + (("\n\t" + ("hello" + "foo bar")) + ("hello" + "hello"))) + "a")) + "hello"))) < ((("a" + "foo bar") + "hello") + ("hello" + "foo bar")))) {
        v1 = v1
        v12 := GenS2{F0: 2147483647, F1: 18446744073709551615, F2: (v1 <= 1)}
        _ = v12
    }
    if ((true && true) || (2147483647 == 42)) {
        var v13 float64 = (3.14159 * 1e10)
        fmt.Println(v1, v1)
        for fi14 := 0; fi14 < 5; fi14++ {
            for fi15 := 0; fi15 < 15; fi15++ {
                v2 = append(v2, 700)
                var v16 int = fi15
                _ = fi15
                _ = v16
            }
            v17 := GenS3{F0: 1, F1: 9223372036854775807, F2: 1}
            v13 = v13
            _ = fi14
            _ = v17
        }
        _ = v13
    }
    var v18 uint64 = 2
    fmt.Println(v11, v4, v18)
    for fi19 := 0; fi19 < 13; fi19++ {
        if (("a" + "foo bar") == "") {
            switch "" {
            case "foo bar":
                var v20 int16 = 0
                _ = v20
            case "hello":
                fmt.Println(v11)
            case "\n\t":
                v11 = 42
            default:
                switch "a" {
                case "":
                    var v21 uint = 18446744073709551615
                    _ = v21
                case "\n\t":
                    fmt.Println(v11, v18)
                }
            }
            v18 = v18
        }
        _ = len(v2)
        _ = fi19
    }
    v1 -= (v1 * v4)
    v22 := GenS3{F0: (v11 / ((v11) | 1)), F1: 0, F2: 1}
    switch -570 {
    case 32767:
        switch 593 {
        case 128:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case -1:
        switch ("foo bar" + "") {
        case "foo bar":
            if false {
                var v23 int = -2147483648
                _ = v23
            } else {
                v11 += 421
            }
        }
    case 1:
        var v24 uint32 = 785
        _ = v24
    }
    switch ("foo bar" + ("" + "hello")) {
    case "":
        v4 = (v1 - v4)
    case "hello":
        for fi25 := 0; fi25 < 20; fi25++ {
            if (false && false) {
                for fi26 := 0; fi26 < 2; fi26++ {
                    v18 -= v18
                    var v27 int16 = v1
                    if true {
                        var v28 uint64 = 4294967295
                        if (-98 > -1) {
                            v3 -= v3
                            var v29 int32 = -498
                            v30 := GenS1{F0: -97, F1: ((("\n\t" + "") + "foo bar") + ""), F2: (100 <= -59)}
                            _ = v29
                            _ = v30
                        }
                        _ = v28
                    } else {
                        v11 += (^(v11))
                        fmt.Println(v1, v3, fi25)
                        v3 = (v3 & uint8(v11))
                    }
                    _ = fi26
                    _ = v27
                }
            } else {
                if false {
                    var v31 uint8 = (v3 | (^(v3)))
                    for fi32 := 0; fi32 < 6; fi32++ {
                        if (399 > v1) {
                            v11 = (v11 | v11)
                            fmt.Println(v1)
                            switch (fi32 & fi25) {
                            case 42:
                                var v33 int16 = (v4 / (((v4 / ((v4) | 1))) | 1))
                                _ = v33
                            case -2147483648:
                                switch (("" + "") + "hello") {
                                case "a":
                                    var v34 float32 = (-2.25 / ((3.14159 / 0.1) + 1.5))
                                    _ = v34
                                case "":
                                    v3 -= v3
                                case "\n\t":
                                    v18 += v18
                                default:
                                    fmt.Println(v3, fi32, v31)
                                }
                            case -9223372036854775808:
                                var v35 bool = true
                                _ = v35
                            default:
                                if len(v2) > 0 {
                                    v2[0] = 2147483648
                                }
                            }
                        }
                        _ = fi32
                    }
                    if (3.14159 == 0.1) {
                        for fi36 := 0; fi36 < 8; fi36++ {
                            var v37 int32 = 1
                            v38 := GenS4{F0: ((v37 & int32(v11)) != v37), F1: 194, F2: 65535}
                            v31 = (v31 << ((4294967295) & 63))
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                        if len(v2) > 0 {
                            v2[0] = 339
                        }
                        var v39 uint8 = (v3 & (v31 & genFn0((v18 + v18), "hello")))
                        _ = v39
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v18 = uint64(v11)
                    }
                    _ = v31
                } else {
                    var v40 int8 = 1
                    var v41 float64 = -2.25
                    var v42 int32 = 42
                    _ = v40
                    _ = v41
                    _ = v42
                }
                if true {
                    var v43 int32 = 0
                    _ = v43
                } else {
                    var v44 uint32 = 2147483647
                    _ = v44
                }
            }
            _ = fi25
        }
    case "a":
        v18 -= (v18 | 1)
    default:
        var v45 int32 = -375
        _ = v45
    }
    var v46 float32 = (0.1 * (float32(-2.25) * 0.0))
    v18 += 0
    switch 0 {
    case -474:
        v4 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v46
}
