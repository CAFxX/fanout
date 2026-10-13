package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint16; F1 int32 }

func main() {
    var v1 uint16 = 0
    var v2 int64 = -121
    v3 := make(map[int]uint32)
    var v4 uint8 = 42
    v5 := GenS1{F0: (v2 + v2), F1: (v4 * v4), F2: 916}
    var v6 uint16 = v1
    fmt.Println(v2, v4)
    v1 -= (^(v1))
    v7 := []int{-95, 594, 2147483647}
    delete(v3, 32768)
    var v8 int32 = 0
    if false {
        for fi9 := 0; fi9 < 18; fi9++ {
            fmt.Println(v2, v2, fi9)
            _ = fi9
        }
        delete(v3, -1)
    } else {
        var v10 int8 = 1
        switch "a" {
        case "hello":
            if (925 == uint16(v4)) {
                switch 128 {
                case 4294967295:
                    var v11 int = -1
                    _ = v11
                case 2147483647:
                    v3[-2147483648] = 2
                }
                if (v10 >= -127) {
                    fmt.Println(v8, v1)
                    v12 := GenS2{F0: 123, F1: (v2 << ((0) & 63))}
                    _ = v12
                } else {
                    var v13 int16 = 1000
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v15 := GenS2{F0: 95, F1: v2}
                        v13 += v13
                        v16 := GenS2{F0: uint(v4), F1: (v2 >> ((644) & 63))}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    if ((true || false) && (2147483647 >= 42)) {
                        var v17 uint8 = (v4 * v4)
                        v18 := GenS2{F0: 2, F1: v2}
                        v17 += v17
                        _ = v17
                        _ = v18
                    } else {
                        v19 := GenS3{F0: 1}
                        var v20 string = (((("foo bar" + "\n\t") + "foo bar") + ("\n\t" + "hello")) + "foo bar")
                        _ = v19
                        _ = v20
                    }
                    _ = v13
                }
                if (!(false || (true && (!(-32768 >= 32767))))) {
                    switch ("foo bar" + ("hello" + ("hello" + ("hello" + "hello")))) {
                    case "hello":
                        var v21 uint8 = (v4 / (((v4 + (v4 % (((v4 << ((1) & 63))) | 1)))) | 1))
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 1; fi22++ {
                            v6 += (^(v6))
                            _ = len(v3)
                            delete(v3, fi22)
                            _ = fi22
                        }
                    }
                    if false {
                        if (637 != v8) {
                            for fi23 := 0; fi23 < 13; fi23++ {
                                for fi24 := 0; fi24 < 19; fi24++ {
                                    var v25 uint8 = (v4 - (v4 | (v4 ^ 127)))
                                    fmt.Println(fi23)
                                    _ = fi24
                                    _ = v25
                                }
                                _ = fi23
                            }
                            if false {
                                for fi26 := 0; fi26 < 20; fi26++ {
                                    fmt.Println(v10, v6, v8)
                                    v4 = v4
                                    fmt.Println(v1, v4, v8)
                                    _ = fi26
                                }
                                fmt.Println(v2, v10)
                                for fi27 := 0; fi27 < 18; fi27++ {
                                    switch "a" {
                                    case "a":
                                        v28 := GenS2{F0: 0, F1: (v2 - (v2 * (^(v2))))}
                                        _ = v28
                                    case "":
                                        fmt.Println(v1)
                                    default:
                                        v29 := GenS4{F0: ((("" + "") + ("" + ("foo bar" + "\n\t"))) + ("hello" + "foo bar")), F1: (v4 * v4)}
                                        _ = v29
                                    }
                                    _ = fi27
                                }
                            }
                        }
                    } else {
                        v30 := GenS5{F0: v6, F1: -1}
                        _ = v30
                    }
                } else {
                    switch 32768 {
                    case 32768:
                        fmt.Println(v10, v8, v6)
                    case 32767:
                        var v31 int64 = 158
                        _ = v31
                    case 958:
                        for fi32 := 0; fi32 < 13; fi32++ {
                            v33 := v3[-9223372036854775808]
                            if false {
                                switch (fi32 >> ((4294967295) & 63)) {
                                case -2147483648:
                                    v34 := GenS5{F0: 363, F1: v8}
                                    _ = v34
                                case 0:
                                    var v35 uint = 4294967295
                                    _ = v35
                                default:
                                    v1 = (v1 % ((v1) | 1))
                                }
                            } else {
                                for fi36 := 0; fi36 < 2; fi36++ {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                    fmt.Println(v1)
                                    _ = fi36
                                }
                            }
                            _ = fi32
                            _ = v33
                        }
                    }
                }
            }
        case "a":
            for fi37 := 0; fi37 < 1; fi37++ {
                var v38 int64 = 0
                _ = fi37
                _ = v38
            }
        default:
            var v39 int16 = 210
            _ = v39
        }
        _ = v10
    }
    v1 = 32767
    fmt.Println(v6)
    v4 -= v4
    for fi40 := 0; fi40 < 2; fi40++ {
        var v41 int8 = 34
        _ = fi40
        _ = v41
    }
    v2 = (-(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
