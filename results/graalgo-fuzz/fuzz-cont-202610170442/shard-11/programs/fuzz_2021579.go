package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int8; F2 uint64 }
type GenS3 struct { F0 int16; F1 uint64; F2 bool }
type GenS4 struct { F0 bool; F1 int; F2 uint16 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int16; F1 uint8; F2 bool }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 uint16; F2 uint16 }

func main() {
    v1 := []string{((("a" + "") + "a") + "foo bar")}
    var v2 uint32 = 522
    v3 := []string{"foo bar", "a", ("hello" + ("" + "\n\t")), "foo bar"}
    var v4 float32 = (0.0 + float32(0.0))
    var v5 float64 = 1.5
    switch -378 {
    case 9223372036854775807:
        v6 := GenS1{F0: (v2 >> ((1) & 63))}
        _ = v6
    default:
        v7 := GenS2{F0: (!(true && (false && false))), F1: -1, F2: 0}
        _ = v7
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4, v4)
    v1 = append(v1, "foo bar")
    v5 = float64((v4 + (v4 * v4)))
    fmt.Println(v5)
    fmt.Println(v4, v5)
    var v8 string = ("\n\t" + "a")
    switch (((("foo bar" + v8) + "") + ("" + v8)) + v8) {
    case "\n\t":
        var v9 uint8 = 132
        _ = v9
    case "":
        if true {
            for fi10 := 0; fi10 < 18; fi10++ {
                fmt.Println(v5, v8, fi10)
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 uint16 = 604
                    _ = fi11
                    _ = v12
                }
                fmt.Println(fi10)
                _ = fi10
            }
        } else {
            for fi13 := 0; fi13 < 20; fi13++ {
                if ((float64(float32(v5)) <= 0.0) && (true && (v5 != float64((v4 - v4))))) {
                    for fi14 := 0; fi14 < 4; fi14++ {
                        v15 := GenS3{F0: 42, F1: 4294967295, F2: (2 >= 424)}
                        if true {
                            v4 = v4
                            var v16 uint16 = 0
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v16
                        } else {
                            v17 := GenS3{F0: 630, F1: 4294967295, F2: true}
                            v1 = append(v1, v8)
                            _ = v17
                        }
                        v4 += v4
                        _ = fi14
                        _ = v15
                    }
                    if false {
                        if (-110 > -2147483648) {
                            v4 += v4
                            switch (fi13 * fi13) {
                            case 205:
                                for fi18 := 0; fi18 < 16; fi18++ {
                                    v1 = append(v1, v8)
                                    _ = fi18
                                }
                            }
                            for fi19 := 0; fi19 < 9; fi19++ {
                                if (((v4 / 1.5) * v4) > 0.0) {
                                    if (0.0 != v5) {
                                        v2 = (v2 | (v2 % (((^(v2))) | 1)))
                                        fmt.Println(v2)
                                        switch fi19 {
                                        case -9223372036854775808:
                                            v1 = append(v1, "a")
                                        case 42:
                                            v20 := GenS2{F0: ((0 == 42) || (-1 == 32767)), F1: 83, F2: 599}
                                            _ = v20
                                        case 32767:
                                            v21 := GenS4{F0: false, F1: (fi13 + (fi19 | fi13)), F2: 42}
                                            _ = v21
                                        default:
                                            var v22 uint8 = 128
                                            _ = v22
                                        }
                                    } else {
                                        v1 = append(v1, (v8 + v8))
                                    }
                                    v23 := GenS5{F0: 32767, F1: 4294967295}
                                    _ = v23
                                }
                                _ = fi19
                            }
                        } else {
                            var v24 string = (v8 + (v8 + v8))
                            _ = v24
                        }
                        v2 = (v2 - v2)
                    }
                    if ((!false) && true) {
                        v5 = (float64((v4 + v4)) / v5)
                        if len(v3) > 0 {
                            v3[0] = ("foo bar" + (v8 + v8))
                        }
                        v25 := GenS3{F0: 1000, F1: 18446744073709551615, F2: (42 <= 66)}
                        _ = v25
                    }
                }
                v4 = v4
                fmt.Println(v4, v2, v8)
                _ = fi13
            }
        }
    }
    for fi26 := 0; fi26 < 12; fi26++ {
        v8 = "foo bar"
        v2 = v2
        _ = fi26
    }
    var v27 float32 = v4
    v28 := GenS6{F0: 42, F1: 227, F2: (false && false)}
    switch v8 {
    case "hello":
        if false {
            var v29 int16 = 1
            if ((false && true) && (!true)) {
                v27 = 1e10
            } else {
                if (uint16(v2) == 44) {
                    v30 := GenS4{F0: (-75 >= 32768), F1: -1, F2: 2}
                    _ = v30
                } else {
                    v4 = v27
                    var v31 float64 = 3.14159
                    v32 := GenS3{F0: (v29 * (v29 + v29)), F1: 1, F2: true}
                    _ = v31
                    _ = v32
                }
                switch 292 {
                case 1:
                    v33 := GenS7{F0: -319}
                    _ = v33
                }
                var v34 uint8 = 2
                _ = v34
            }
            _ = v29
        }
    case "":
        if (1 == -9223372036854775808) {
            fmt.Println(v2, v27)
            v3 = append(v3, ("a" + "foo bar"))
        }
    }
    fmt.Println(v8)
    var v35 uint32 = v2
    for fi36 := 0; fi36 < 9; fi36++ {
        var v37 uint32 = v35
        switch (fi36 & (fi36 << ((518) & 63))) {
        case 2147483647:
            if true {
                v38 := GenS6{F0: 1000, F1: 255, F2: (0 != -32768)}
                if len(v1) > 0 {
                    v1[0] = ("\n\t" + "hello")
                }
                v35 = 0
                _ = v38
            }
        case 210:
            for fi39 := 0; fi39 < 2; fi39++ {
                var v40 uint8 = 42
                if false {
                    v41 := GenS3{F0: -941, F1: 420, F2: true}
                    _ = v41
                }
                if (!(-1 < 906)) {
                    if (!(0 >= 401)) {
                        if true {
                            v40 = (v40 % (((v40 | 255)) | 1))
                        }
                        var v42 uint64 = 2
                        _ = v42
                    }
                } else {
                    switch "a" {
                    case "a":
                        switch fi39 {
                        case 32767:
                            v2 -= v35
                        case 574:
                            var v43 int32 = 1
                            _ = v43
                        case 160:
                            var v44 uint64 = 0
                            _ = v44
                        }
                    case "foo bar":
                        _ = len(v3)
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    var v45 uint8 = v40
                    _ = v45
                }
                _ = fi39
                _ = v40
            }
        default:
            v46 := GenS8{F0: (v35 * (v2 / (((v2 % (((v2 << ((65535) & 63))) | 1))) | 1))), F1: 255, F2: 32767}
            _ = v46
        }
        _ = fi36
        _ = v37
    }
    var v47 int = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v47
    _ = v5
    _ = v8
}
