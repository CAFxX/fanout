package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint; F1 uint8; F2 int8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 int64; F2 int }
type GenS7 struct { F0 int16; F1 int }
type GenS8 struct { F0 int16; F1 string }

func main() {
    v1 := []string{(("\n\t" + "foo bar") + ("a" + "\n\t")), "\n\t"}
    var v2 int8 = -128
    var v3 uint32 = 65535
    var v4 int = 9223372036854775807
    var v5 float64 = 0.0
    if false {
        v5 = v5
        for fi6 := 0; fi6 < 13; fi6++ {
            switch 128 {
            case -1:
                var v7 uint32 = v3
                _ = v7
            default:
                if false {
                    if (!(true || ((-1 <= -930) || (65535 < 635)))) {
                        switch fi6 {
                        case 66:
                            if (0.0 > 0.0) {
                                var v8 uint16 = 1
                                v3 += (v3 + (^(v3)))
                                _ = v8
                            }
                        case -470:
                            if (float32((float64(3.14159) - v5)) != ((float32((0.1 + -2.25)) * 0.1) * ((-2.25 / (3.14159 / (3.14159 * float32(v5)))) / float32(v5)))) {
                                fmt.Println(v3, v2)
                                fmt.Println(v3, v4)
                            }
                        case 4294967295:
                            if false {
                                if false {
                                    if true {
                                        var v9 string = "hello"
                                        _ = v9
                                    }
                                    v1 = append(v1, "\n\t")
                                } else {
                                    var v10 uint = 995
                                    _ = v10
                                }
                                _ = len(v1)
                            }
                        }
                        if ((v2 & 100) >= v2) {
                            v4 += 822
                            v3 += 42
                        } else {
                            var v11 uint8 = 127
                            _ = v11
                        }
                        v12 := GenS1{F0: 127, F1: 32767}
                        _ = v12
                    } else {
                        v1 = append(v1, "foo bar")
                    }
                    if false {
                        v13 := GenS2{F0: (v5 < v5), F1: -53}
                        if (858 < -654) {
                            var v14 uint16 = 65535
                            for fi15 := 0; fi15 < 15; fi15++ {
                                var v16 string = ""
                                switch "a" {
                                case "hello":
                                    switch ((((v16 + "hello") + "a") + ("foo bar" + ("" + "foo bar"))) + v16) {
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "":
                                            v14 = v14
                                        case "foo bar":
                                            var v17 uint16 = v14
                                            _ = v17
                                        case "hello":
                                            _ = len(v1)
                                        }
                                    default:
                                        var v18 float64 = v5
                                        _ = v18
                                    }
                                case "\n\t":
                                    if false {
                                        var v19 int8 = (^(v2))
                                        _ = v19
                                    }
                                }
                                _ = fi15
                                _ = v16
                            }
                            _ = v14
                        } else {
                            var v20 float32 = 0.0
                            v4 = (v4 * 259)
                            _ = v20
                        }
                        fmt.Println(v5, v3, v3)
                        _ = v13
                    } else {
                        v21 := GenS3{F0: "hello"}
                        _ = v21
                    }
                }
            }
            for fi22 := 0; fi22 < 1; fi22++ {
                v23 := GenS1{F0: 255, F1: 494}
                if (fi22 >= fi6) {
                    v2 -= (-(v2))
                    fmt.Println(v5, v5)
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                    fmt.Println(v2)
                }
                var v24 uint8 = 127
                if true {
                    v3 += 65535
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = fi22
                _ = v23
                _ = v24
            }
            fmt.Println(fi6, v2)
            _ = fi6
        }
        v25 := GenS1{F0: 4294967295, F1: 688}
        _ = v25
    } else {
        v26 := GenS4{F0: 18446744073709551615, F1: 212, F2: (v2 / (((v2 * (v2 ^ v2))) | 1))}
        var v27 int32 = -538
        fmt.Println(v2)
        _ = v26
        _ = v27
    }
    v5 = (((v5 + 0.1) - 0.1) * v5)
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 bool = (1 >= 255)
        v30 := GenS5{F0: (42 != 0)}
        v29 = (0 >= 9223372036854775807)
        _ = fi28
        _ = v29
        _ = v30
    }
    v31 := GenS6{F0: (v4 ^ (v4 >> ((18446744073709551615) & 63))), F1: 9223372036854775807, F2: v4}
    v5 -= 0.1
    for fi32 := 0; fi32 < 6; fi32++ {
        fmt.Println(fi32, fi32)
        _ = fi32
    }
    var v33 int = (v4 >> ((255) & 63))
    for fi34 := 0; fi34 < 2; fi34++ {
        _ = len(v1)
        _ = fi34
    }
    v3 = (v3 ^ v3)
    for fi35 := 0; fi35 < 4; fi35++ {
        if false {
            v36 := GenS7{F0: 750, F1: (v4 % ((32768) | 1))}
            v2 -= 100
            _ = v36
        } else {
            v3 = (v3 / ((v3) | 1))
            v5 = float64(((1e10 - 3.14159) + -2.25))
            fmt.Println(v5, v4, v4)
        }
        for fi37 := 0; fi37 < 20; fi37++ {
            _ = len(v1)
            _ = fi37
        }
        v33 = v33
        _ = fi35
    }
    v4 += 2147483647
    var v38 uint32 = v3
    var v39 uint = 255
    if (65535 != uint16(v38)) {
        v33 += 9223372036854775807
        v1 = append(v1, "\n\t")
        v3 = (v3 << ((v39) & 63))
    } else {
        v40 := GenS8{F0: 599, F1: "a"}
        v5 = (float64(1e10) * (1.5 * v5))
        _ = v40
    }
    v41 := GenS7{F0: int16(v33), F1: (v4 ^ (v33 & 864))}
    if false {
        for fi42 := 0; fi42 < 18; fi42++ {
            v4 = fi42
            v43 := GenS1{F0: 255, F1: 409}
            v4 -= (v4 - (^(v33)))
            _ = fi42
            _ = v43
        }
        v5 -= v5
    }
    v1 = append(v1, "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v38
    _ = v39
    _ = v4
    _ = v41
    _ = v5
}
