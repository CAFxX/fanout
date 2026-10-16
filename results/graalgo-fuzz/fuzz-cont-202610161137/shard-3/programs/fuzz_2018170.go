package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 bool }

func main() {
    var v1 uint = 999
    v2 := make(map[uint32]int)
    var v3 int = 42
    v4 := []string{"hello", (("hello" + "a") + "")}
    v5 := []uint32{}
    var v6 string = "foo bar"
    delete(v2, 42)
    switch (v6 + v6) {
    case "foo bar":
        v6 = (v6 + ("a" + (v6 + "")))
    case "\n\t":
        var v7 uint = 1
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: 1000}
            if len(v5) > 0 {
                v5[0] = 385
            }
            _ = fi8
            _ = v9
        }
    default:
        v6 = "a"
    }
    v1 -= v1
    for fi10 := 0; fi10 < 9; fi10++ {
        _ = len(v4)
        _ = fi10
    }
    if (!false) {
        switch v3 {
        case 474:
            v1 = (v1 / ((v1) | 1))
        case -913:
            if len(v4) > 0 {
                v4[0] = ""
            }
        case 161:
            if false {
                if false {
                    if (42 != 1000) {
                        if true {
                            v11 := GenS2{F0: v6, F1: v1}
                            _ = v11
                        }
                    }
                    var v12 uint32 = 114
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 18; fi13++ {
                    v1 -= 65535
                    for fi14 := 0; fi14 < 6; fi14++ {
                        if true {
                            fmt.Println(fi13, fi13, fi14)
                        } else {
                            switch (v6 + (((("foo bar" + ((("\n\t" + v6) + v6) + (("\n\t" + "a") + v6))) + v6) + v6) + ((v6 + v6) + "foo bar"))) {
                            case "foo bar":
                                v1 -= (v1 ^ 983)
                            case "hello":
                                var v15 uint16 = 764
                                _ = v15
                            case "":
                                if true {
                                    var v16 int = -9223372036854775808
                                    v5 = append(v5, 0)
                                    _ = v16
                                }
                            }
                            switch (v6 + v6) {
                            case "":
                                v1 = (v1 / ((v1) | 1))
                            case "a":
                                v6 = v6
                                if false {
                                    v17 := v2[2147483647]
                                    _ = len(v4)
                                    v2[4294967295] = (v3 / (((v17 - fi14)) | 1))
                                    _ = v17
                                }
                            case "\n\t":
                                fmt.Println(fi13, v1, v1)
                            }
                            var v18 int8 = -1
                            _ = v18
                        }
                        v19 := GenS3{F0: 1}
                        _ = fi14
                        _ = v19
                    }
                    _ = fi13
                }
                fmt.Println(v6, v6)
                for fi20 := 0; fi20 < 18; fi20++ {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        _ = len(v5)
                        v1 = v1
                        _ = fi21
                    }
                    _ = fi20
                }
            }
        default:
            var v22 float32 = (float32(-2.25) * 1.5)
            _ = v22
        }
        switch "hello" {
        case "hello":
            _ = len(v5)
        case "a":
            v3 = v3
        case "":
            var v23 float32 = (1.5 / 1.5)
            _ = v23
        }
        v6 = v6
    } else {
        var v24 uint8 = 42
        v3 += int(v24)
        v25 := GenS2{F0: v6, F1: 18446744073709551615}
        _ = v24
        _ = v25
    }
    switch "\n\t" {
    case "foo bar":
        v1 = v1
    case "":
        if false {
            switch v3 {
            case 2147483647:
                if (-1 == 62) {
                    v2[942] = v3
                    v1 = v1
                    for fi26 := 0; fi26 < 16; fi26++ {
                        var v27 bool = true
                        _ = fi26
                        _ = v27
                    }
                } else {
                    fmt.Println(v3, v6)
                }
            default:
                var v28 int8 = -126
                _ = v28
            }
            if false {
                switch ("foo bar" + (v6 + v6)) {
                case "foo bar":
                    var v29 uint64 = 2
                    _ = v29
                case "hello":
                    switch "" {
                    case "\n\t":
                        v4 = append(v4, v6)
                    case "foo bar":
                        fmt.Println(v3, v3)
                    default:
                        if len(v5) > 0 {
                            v5[0] = 405
                        }
                    }
                default:
                    v30 := GenS4{F0: 42, F1: true}
                    _ = v30
                }
                v2[1] = (v3 << ((v1) & 63))
            }
            var v31 float32 = 0.0
            _ = v31
        }
    case "a":
        var v32 uint16 = 2
        _ = v32
    default:
        v1 += (v1 ^ v1)
    }
    var v33 bool = false
    v34 := GenS1{F0: -1}
    if (0 >= uint32(v1)) {
        var v35 int32 = 708
        for fi36 := 0; fi36 < 2; fi36++ {
            if ((!(0 != 0)) && v33) {
                fmt.Println(v1, v3, v3)
                v1 -= 18446744073709551615
            }
            v1 = v1
            _ = fi36
        }
        v37 := GenS4{F0: 95, F1: false}
        _ = v35
        _ = v37
    } else {
        switch v3 {
        case 2147483647:
            var v38 float32 = ((0.0 + (0.1 + (1.5 - 3.14159))) - (1e10 / 1.5))
            _ = v38
        case 4294967295:
            v39 := GenS5{F0: false}
            _ = v39
        default:
            delete(v2, 65535)
        }
        for fi40 := 0; fi40 < 4; fi40++ {
            v41 := GenS1{F0: -82}
            _ = fi40
            _ = v41
        }
    }
    switch v3 {
    case 1:
        var v42 bool = (42 < 0)
        _ = v42
    }
    var v43 int = (-(v3))
    if len(v4) > 0 {
        v4[0] = ""
    }
    var v44 uint64 = uint64(v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
