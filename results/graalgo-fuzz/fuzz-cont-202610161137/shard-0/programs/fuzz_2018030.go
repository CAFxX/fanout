package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 int8; F2 uint }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []string{"hello"}
    var v2 uint32 = 2
    v3 := []uint32{}
    var v4 int32 = 2147483647
    var v5 float32 = 3.14159
    v6 := GenS1{F0: -32768}
    v4 = v4
    var v7 int = 42
    v1 = append(v1, "hello")
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int32 = 0
        v10 := GenS1{F0: -51}
        _ = fi8
        _ = v10
        _ = v9
    }
    if false {
        var v11 uint8 = 1
        var v12 int64 = 42
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint32 = v2
        _ = fi13
        _ = v14
    }
    if true {
        if false {
            if (!true) {
                var v15 int32 = v4
                if false {
                    switch 128 {
                    case -2147483648:
                        v7 = (v7 % ((0) | 1))
                    case 32768:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        var v16 int16 = 32767
                        _ = v16
                    }
                    var v17 int32 = (-(v4))
                    var v18 uint32 = v2
                    _ = v17
                    _ = v18
                }
                _ = v15
            }
        } else {
            if len(v3) > 0 {
                v3[0] = (v2 + v2)
            }
        }
        v5 = float32(0.1)
        v19 := GenS1{F0: 1}
        _ = v19
    } else {
        if (!false) {
            if (42 == 61) {
                v7 += v7
                fmt.Println(v4, v4, v5)
                for fi20 := 0; fi20 < 13; fi20++ {
                    if (((v2 + v2) < v2) || true) {
                        if len(v1) > 0 {
                            v1[0] = ("" + "\n\t")
                        }
                        var v21 bool = (false && true)
                        v7 += v7
                        _ = v21
                    }
                    _ = fi20
                }
            }
        } else {
            var v22 uint16 = uint16(v2)
            var v23 uint = 0
            switch "\n\t" {
            case "a":
                var v24 bool = (1e10 != 0.0)
                _ = v24
            default:
                fmt.Println(v7, v23)
            }
            _ = v22
            _ = v23
        }
    }
    var v25 uint8 = 0
    var v26 int16 = 0
    v2 = v2
    switch 32767 {
    case 2147483647:
        v27 := GenS2{F0: v4}
        _ = v27
    case -1:
        if (((float32((-2.25 / 1.5)) * v5) < (1e10 - float32(3.14159))) && true) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi28 := 0; fi28 < 14; fi28++ {
                var v29 int = v7
                if false {
                    _ = len(v1)
                    switch v7 {
                    case -2147483648:
                        switch "foo bar" {
                        case "":
                            v5 -= v5
                        case "foo bar":
                            var v30 int8 = -125
                            _ = v30
                        case "hello":
                            for fi31 := 0; fi31 < 1; fi31++ {
                                fmt.Println(v4, fi28, v5)
                                _ = fi31
                            }
                        }
                    case 128:
                        if false {
                            var v32 string = "foo bar"
                            _ = v32
                        } else {
                            for fi33 := 0; fi33 < 20; fi33++ {
                                v1 = append(v1, "a")
                                v2 = (^(v2))
                                if true {
                                    fmt.Println(v5, v25)
                                    v34 := GenS2{F0: (-(v4))}
                                    v35 := GenS2{F0: -2147483648}
                                    _ = v34
                                    _ = v35
                                } else {
                                    v25 -= v25
                                    if (0.0 < float32(0.1)) {
                                        v26 = 338
                                        var v36 int = v7
                                        v3 = append(v3, 42)
                                        var v37 uint8 = 197
                                        _ = v36
                                        _ = v37
                                    }
                                    v38 := GenS3{F0: 100, F1: -128, F2: 42}
                                    _ = v38
                                }
                                _ = fi33
                            }
                            v2 += 42
                        }
                    case 127:
                        v4 += (v4 & v4)
                    default:
                        fmt.Println(v25, v2, v29)
                    }
                }
                _ = fi28
                _ = v29
            }
        }
    default:
        var v39 string = (("" + ("\n\t" + ("\n\t" + "hello"))) + "hello")
        _ = v39
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 float32 = (1.5 / v5)
        if len(v3) > 0 {
            v3[0] = (v2 << ((255) & 63))
        }
        for fi42 := 0; fi42 < 14; fi42++ {
            if (1.5 < v5) {
                switch "hello" {
                case "hello":
                    switch (-(v7)) {
                    case 127:
                        fmt.Println(v4, v26)
                    case -2147483648:
                        v26 = v26
                    case -1:
                        if (!false) {
                            v43 := GenS4{F0: 2}
                            v44 := GenS4{F0: 255}
                            _ = v43
                            _ = v44
                        }
                    default:
                        if false {
                            v5 = 1.5
                        } else {
                            if len(v3) > 0 {
                                v3[0] = (v2 & v2)
                            }
                            fmt.Println(v25, v7, v41)
                        }
                    }
                case "a":
                    var v45 int64 = 9223372036854775807
                    _ = v45
                }
                var v46 int32 = v4
                if (false || false) {
                    if true {
                        var v47 string = (("foo bar" + ("foo bar" + "hello")) + ((((("hello" + ("" + "a")) + ("foo bar" + ("hello" + ("" + "hello")))) + "\n\t") + ("a" + "hello")) + ("a" + "a")))
                        v5 -= (float32((float64(-2.25) + 1e10)) - float32(3.14159))
                        switch "foo bar" {
                        case "a":
                            if true {
                                var v48 float32 = v5
                                var v49 float32 = (float32(3.14159) / (((v48 + v5) + v5) * v48))
                                _ = v48
                                _ = v49
                            } else {
                                v25 -= (v25 - v25)
                                fmt.Println(fi40, fi42)
                            }
                        default:
                            _ = len(v3)
                        }
                        if (42 >= 32767) {
                            var v50 int32 = int32(v25)
                            var v51 int16 = (v26 - (v26 * v26))
                            v2 += (^(v2))
                            _ = v50
                            _ = v51
                        }
                        v7 -= int(v26)
                        _ = v47
                    } else {
                        v52 := GenS1{F0: v26}
                        if false {
                            _ = len(v3)
                            v46 += 1
                            var v53 float32 = v5
                            _ = v53
                        }
                        _ = v52
                    }
                } else {
                    fmt.Println(v4)
                    _ = len(v3)
                    if ((2 <= 518) && (((0.0 >= v5) && false) || (((^(v25)) >= 34) && false))) {
                        v7 = v7
                        _ = len(v3)
                        v54 := GenS1{F0: v26}
                        _ = v54
                    } else {
                        if true {
                            v5 = v5
                        } else {
                            v1 = append(v1, "foo bar")
                            _ = len(v1)
                            v55 := GenS3{F0: 127, F1: -118, F2: 42}
                            _ = v55
                        }
                        var v56 float32 = 0.0
                        v2 = 213
                        fmt.Println(v4)
                        _ = v56
                    }
                }
                fmt.Println(v4, v26)
                _ = v46
            }
            fmt.Println(fi40, v41)
            _ = fi42
        }
        _ = fi40
        _ = v41
    }
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
