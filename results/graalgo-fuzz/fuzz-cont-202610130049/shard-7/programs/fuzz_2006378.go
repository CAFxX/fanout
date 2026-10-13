package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int16; F1 uint8; F2 uint8 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 int32 = 42
    var v2 uint = 978
    v3 := []string{"\n\t", ("hello" + (("foo bar" + ("a" + "hello")) + "foo bar")), "\n\t"}
    var v4 int = 32767
    v5 := []uint32{1, 4294967295}
    var v6 uint32 = 65535
    var v7 float64 = 1.5
    if (!(42 > -702)) {
        v8 := GenS1{F0: (v4 | -1), F1: ("hello" + ((((((("hello" + ("a" + "")) + ("foo bar" + "foo bar")) + "foo bar") + "") + (("foo bar" + ("" + ("a" + ""))) + "hello")) + "") + ("a" + ("\n\t" + ""))))}
        _ = v8
    }
    v6 -= v6
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 bool = ((v4 + (fi9 << (((v2 % ((4294967295) | 1))) & 63))) <= 42)
        _ = fi9
        _ = v10
    }
    var v11 string = (("\n\t" + "foo bar") + "\n\t")
    switch (("hello" + v11) + ((v11 + v11) + v11)) {
    case "\n\t":
        for fi12 := 0; fi12 < 6; fi12++ {
            if false {
                switch v4 {
                case -595:
                    v13 := GenS2{F0: 9223372036854775807, F1: int16(v6)}
                    _ = v13
                }
                var v14 bool = false
                var v15 float32 = (float32(3.14159) * 0.1)
                _ = v14
                _ = v15
            } else {
                if len(v3) > 0 {
                    v3[0] = v11
                }
                var v16 uint64 = 0
                _ = v16
            }
            fmt.Println(fi12, v4)
            var v17 int8 = -128
            _ = fi12
            _ = v17
        }
    }
    if true {
        for fi18 := 0; fi18 < 13; fi18++ {
            if (!false) {
                for fi19 := 0; fi19 < 8; fi19++ {
                    fmt.Println(v4)
                    if ((!((v7 - 3.14159) > ((((0.1 + float64(1e10)) / v7) / (0.0 / v7)) - v7))) || true) {
                        if true {
                            if (!((671 == 4294967295) && false)) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v20 int16 = -1
                                var v21 int = fi19
                                _ = v20
                                _ = v21
                            }
                        }
                        v5 = append(v5, v6)
                        v4 += (v4 * 42)
                        switch (fi19 / (((fi19 >> ((29) & 63))) | 1)) {
                        case 883:
                            var v22 int32 = v1
                            _ = v22
                        case -457:
                            var v23 uint16 = 2
                            _ = v23
                        }
                        var v24 int64 = 1
                        _ = v24
                    } else {
                        v4 -= -833
                        v11 = "foo bar"
                    }
                    if (2147483647 < 127) {
                        if (1e10 == 1e10) {
                            if (2 < 65535) {
                                v1 -= v1
                                var v25 string = ("\n\t" + ((v11 + "hello") + v11))
                                _ = v25
                            } else {
                                if (!false) {
                                    v26 := GenS1{F0: fi19, F1: (v11 + ("hello" + "foo bar"))}
                                    v5 = append(v5, 1)
                                    _ = v26
                                }
                                _ = len(v3)
                                v27 := GenS1{F0: fi19, F1: (v11 + v11)}
                                _ = v27
                            }
                            switch (fi19 & fi18) {
                            case 32768:
                                v28 := GenS1{F0: 32768, F1: v11}
                                _ = v28
                            default:
                                fmt.Println(v4)
                            }
                            if (!false) {
                                fmt.Println(v11, v11, v4)
                                if (true || false) {
                                    switch (v11 + v11) {
                                    case "":
                                        v5 = append(v5, 2147483647)
                                    case "hello":
                                        var v29 int8 = -1
                                        _ = v29
                                    default:
                                        v30 := GenS2{F0: -9223372036854775808, F1: -38}
                                        _ = v30
                                    }
                                    if false {
                                        v7 -= (3.14159 * float64(1.5))
                                        var v31 int8 = 1
                                        v32 := GenS1{F0: 139, F1: ("\n\t" + (v11 + "a"))}
                                        _ = v31
                                        _ = v32
                                    }
                                } else {
                                    switch (v4 + fi19) {
                                    case 42:
                                        v33 := GenS3{F0: ((v4 <= (v4 ^ (fi19 / ((9223372036854775807) | 1)))) || false)}
                                        _ = v33
                                    case 2147483647:
                                        switch (fi18 >> ((1) & 63)) {
                                        case 604:
                                            v34 := GenS2{F0: 1, F1: -723}
                                            _ = v34
                                        case 2147483647:
                                            v4 = (fi18 << (((v2 + v2)) & 63))
                                        default:
                                            v35 := GenS4{F0: 42, F1: 2147483647}
                                            _ = v35
                                        }
                                    case -1:
                                        _ = len(v5)
                                    }
                                }
                            } else {
                                var v36 uint16 = 89
                                switch v11 {
                                case "a":
                                    var v37 int64 = -1
                                    _ = v37
                                case "hello":
                                    fmt.Println(v6, v6)
                                    var v38 uint32 = v6
                                    fmt.Println(v1, v36, v38)
                                    _ = v38
                                }
                                if ((v4 >> ((v2) & 63)) <= v4) {
                                    switch 415 {
                                    case 32767:
                                        v1 += int32(v6)
                                    default:
                                        v11 = (((v11 + v11) + (v11 + "foo bar")) + "")
                                    }
                                    v1 -= v1
                                    v11 = v11
                                } else {
                                    fmt.Println(v6, v7, v6)
                                }
                                _ = v36
                            }
                            var v39 uint = (v2 / ((v2) | 1))
                            v6 += v6
                            _ = v39
                        }
                    }
                    _ = fi19
                }
            } else {
                var v40 int = fi18
                _ = v40
            }
            v1 += 1
            _ = fi18
        }
        switch v4 {
        case -972:
            fmt.Println(v1, v2)
        default:
            v7 = (1e10 + 3.14159)
        }
        v41 := GenS5{F0: true, F1: -1}
        _ = v41
    }
    switch "" {
    case "foo bar":
        if ((v11 <= (v11 + "a")) && (1 != 2147483647)) {
            for fi42 := 0; fi42 < 17; fi42++ {
                fmt.Println(v4, v1)
                var v43 int32 = 839
                _ = fi42
                _ = v43
            }
            v44 := GenS6{F0: 1000, F1: 127, F2: 193}
            _ = v44
        }
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        if true {
            v2 -= (^(v2))
            var v45 float64 = (0.1 - v7)
            _ = v45
        }
    default:
        v7 -= v7
    }
    v11 = ("a" + ((v11 + v11) + "hello"))
    var v46 uint16 = 903
    v47 := GenS7{F0: v1}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
