package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint32; F2 int16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = -306
    v3 := []int{32768, 790}
    v4 := []int{9223372036854775807, -2147483648, 960, 127}
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]int)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: -32768, F1: 0}
        v9 := GenS1{F0: 1000, F1: 24}
        var v10 float32 = ((1.5 + 1.5) - 3.14159)
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if true {
        fmt.Println(v2)
    }
    v5["foo bar"] = 415
    v2 -= v2
    v11 := GenS2{F0: true}
    var v12 uint32 = 710
    v2 += v2
    if false {
        switch 928 {
        case 4294967295:
            var v13 uint16 = 42
            _ = v13
        case -1:
            if len(v4) > 0 {
                v4[0] = 499
            }
        default:
            for fi14 := 0; fi14 < 11; fi14++ {
                v12 += v12
                if (834 < 0) {
                    var v15 string = ("" + "a")
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 8; fi16++ {
                        if false {
                            var v17 int8 = -1
                            fmt.Println(v2)
                            if (!false) {
                                if (!(459 != 33)) {
                                    v18 := v1[("" + "foo bar")]
                                    _ = v18
                                } else {
                                    switch (fi14 - (^(fi14))) {
                                    case -2147483648:
                                        v12 = (^(v12))
                                    }
                                    v12 -= v12
                                }
                                v12 += v12
                            }
                            v17 -= (v17 | 42)
                            _ = v17
                        }
                        var v19 int32 = v2
                        _ = fi16
                        _ = v19
                    }
                }
                _ = fi14
            }
        }
    } else {
        v2 = v2
        for fi20 := 0; fi20 < 12; fi20++ {
            v2 += (-(v2))
            if (false || (((("\n\t" + "\n\t") + "foo bar") + "foo bar") <= ("hello" + "a"))) {
                switch ((((("foo bar" + "a") + "") + "a") + (("a" + "a") + "foo bar")) + "\n\t") {
                case "hello":
                    var v21 uint64 = 954
                    _ = v21
                case "\n\t":
                    fmt.Println(v2)
                default:
                    if true {
                        if (!(fi20 <= fi20)) {
                            fmt.Println(v2, v2, fi20)
                            for fi22 := 0; fi22 < 5; fi22++ {
                                v12 -= v12
                                v2 = v2
                                _ = fi22
                            }
                            for fi23 := 0; fi23 < 11; fi23++ {
                                fmt.Println(v12, fi23, v2)
                                var v24 int16 = -1
                                switch 32768 {
                                case 823:
                                    fmt.Println(v12, v2, v2)
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                    v2 = 0
                                case -9223372036854775808:
                                    v12 += (v12 ^ v12)
                                default:
                                    var v25 uint64 = 0
                                    _ = v25
                                }
                                _ = fi23
                                _ = v24
                            }
                        }
                        v26 := GenS3{F0: 2147483647, F1: fi20}
                        _ = v26
                    }
                }
                v3 = append(v3, fi20)
                v2 = v2
            } else {
                if (true || false) {
                    fmt.Println(fi20, v12)
                    var v27 uint16 = 65535
                    if false {
                        var v28 uint64 = 2
                        v2 = v2
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 18; fi29++ {
                            _ = len(v3)
                            switch "" {
                            case "\n\t":
                                if ((v12 & 4294967295) < (v12 << ((732) & 63))) {
                                    var v30 float32 = (float32((1.5 * 3.14159)) * (0.0 * -2.25))
                                    var v31 float64 = -2.25
                                    v31 += v31
                                    _ = v30
                                    _ = v31
                                }
                            }
                            _ = fi29
                        }
                        var v32 uint = 1
                        _ = v32
                    }
                    _ = v27
                } else {
                    if (v12 <= v12) {
                        for fi33 := 0; fi33 < 13; fi33++ {
                            fmt.Println(fi20)
                            switch (fi20 * fi33) {
                            case 316:
                                if (!false) {
                                    v2 = (v2 >> ((2) & 63))
                                    fmt.Println(fi20)
                                    v34 := GenS4{F0: "hello"}
                                    _ = v34
                                } else {
                                    _ = len(v3)
                                }
                            }
                            v3 = append(v3, fi33)
                            _ = fi33
                        }
                        switch (fi20 ^ 1) {
                        case 2147483647:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        default:
                            delete(v5, "hello")
                        }
                    }
                }
                for fi35 := 0; fi35 < 13; fi35++ {
                    v36 := GenS1{F0: 42, F1: 127}
                    v12 -= v12
                    v37 := GenS5{F0: fi35, F1: 518, F2: 706}
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
            _ = fi20
        }
    }
    var v38 bool = true
    v1[("a" + ("\n\t" + ("\n\t" + ("foo bar" + ("" + ("\n\t" + ("hello" + "hello")))))))] = ((255 == 255) || (32768 <= -387))
    if v38 {
        if (42 >= 1) {
            v39 := GenS4{F0: "hello"}
            v40 := GenS2{F0: v38}
            fmt.Println(v2, v12)
            _ = v39
            _ = v40
        } else {
            v12 -= 371
            if (v2 != (v2 % ((v2) | 1))) {
                v38 = (0 < 128)
                switch 128 {
                case -2147483648:
                    var v41 float64 = 3.14159
                    _ = v41
                }
            }
        }
        v42 := GenS3{F0: v2, F1: 4294967295}
        _ = v42
    }
    switch 0 {
    case -211:
        var v43 uint8 = 42
        _ = v43
    }
    fmt.Println(v38, v2)
    v2 += -2147483648
    for fi44 := 0; fi44 < 8; fi44++ {
        switch "" {
        case "":
            fmt.Println(v38)
        case "a":
            v12 -= v12
        case "\n\t":
            v45 := GenS2{F0: false}
            _ = v45
        }
        _ = fi44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
