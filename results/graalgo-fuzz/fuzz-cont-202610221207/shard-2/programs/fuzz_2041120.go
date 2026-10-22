package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 int; F1 int64; F2 uint16 }

func main() {
    var v1 int32 = -2147483648
    v2 := []uint32{}
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    if ((0 < 42) && (!true)) {
        v1 = (v1 / (((v1 >> ((440) & 63))) | 1))
        v3["foo bar"] = "a"
    } else {
        v5 := GenS1{F0: 65535, F1: v1, F2: -1}
        _ = v5
    }
    v1 = (^(v1))
    v1 += (v1 | (v1 ^ v1))
    _ = len(v3)
    var v6 string = "foo bar"
    var v7 uint16 = 1
    delete(v4, 4294967295)
    var v8 int32 = v1
    v9 := GenS1{F0: 817, F1: v8, F2: 127}
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(v7, v6)
        v1 -= v1
        _ = fi10
    }
    var v11 uint64 = 18446744073709551615
    switch v6 {
    case "a":
        if false {
            v12 := v3[(v6 + (v6 + (v6 + v6)))]
            v2 = append(v2, 553)
            _ = v12
        }
    case "":
        var v13 int8 = -1
        _ = v13
    }
    switch v6 {
    case "a":
        var v14 float64 = float64(-2.25)
        _ = v14
    case "":
        var v15 int64 = 42
        _ = v15
    }
    switch -913 {
    case 623:
        switch "\n\t" {
        case "a":
            v16 := v3[v6]
            _ = v16
        case "":
            v17 := v4[1]
            _ = v17
        case "foo bar":
            for fi18 := 0; fi18 < 18; fi18++ {
                if (0 >= -86) {
                    switch v6 {
                    case "hello":
                        v6 = v6
                    }
                    v11 += 0
                } else {
                    _ = len(v3)
                    var v19 int8 = 42
                    _ = v19
                }
                if false {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        v2 = append(v2, 2147483647)
                        var v21 int32 = (v1 >> ((42) & 63))
                        var v22 uint16 = (v7 + 514)
                        v22 = v7
                        var v23 float64 = 1.5
                        _ = fi20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(v6, v1, v7)
                    for fi24 := 0; fi24 < 11; fi24++ {
                        var v25 int8 = 1
                        v26 := GenS2{F0: v1, F1: (^(v7)), F2: int(v11)}
                        _ = len(v4)
                        switch int(v7) {
                        case 42:
                            v11 -= 42
                        default:
                            var v27 int8 = v25
                            _ = v27
                        }
                        v1 = (v8 + (v1 - v8))
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
                if false {
                    if false {
                        var v28 uint16 = v7
                        switch "hello" {
                        case "a":
                            var v29 float32 = (-2.25 + 0.0)
                            _ = v29
                        default:
                            if false {
                                var v30 int16 = -661
                                v1 = 0
                                _ = v30
                            } else {
                                for fi31 := 0; fi31 < 10; fi31++ {
                                    switch (fi18 ^ (fi31 | (fi31 >> ((uint(v11)) & 63)))) {
                                    case 2147483647:
                                        v1 -= (v1 & 0)
                                    case 9223372036854775807:
                                        v32 := GenS3{F0: -1}
                                        switch "hello" {
                                        case "":
                                            v11 -= v11
                                        case "a":
                                            fmt.Println(fi31)
                                            if (("\n\t" <= ((((((v6 + v6) + (v6 + v6)) + v6) + v6) + v6) + v6)) && true) {
                                                v33 := GenS4{F0: (^(v11)), F1: 2}
                                                v34 := GenS1{F0: 2147483648, F1: 42, F2: fi31}
                                                _ = v33
                                                _ = v34
                                            } else {
                                                if (!(100 >= 100)) {
                                                    v6 = (v6 + v6)
                                                    switch fi18 {
                                                    case 606:
                                                        v6 = "foo bar"
                                                    case 2147483647:
                                                        var v35 float64 = 0.1
                                                        _ = v35
                                                    case -10:
                                                        if (4294967295 <= 4294967295) {
                                                            v28 = (v28 + v28)
                                                            v28 -= v7
                                                        } else {
                                                            v6 = v6
                                                            var v36 uint64 = v11
                                                            _ = v36
                                                        }
                                                    }
                                                    v7 -= (^(v7))
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    switch ((("a" + ("hello" + ("\n\t" + v6))) + v6) + ("foo bar" + v6)) {
                                                    case "a":
                                                        fmt.Println(v11)
                                                    case "\n\t":
                                                        var v37 uint = 135
                                                        _ = v37
                                                    }
                                                    fmt.Println(v11, fi31, v11)
                                                }
                                                fmt.Println(v7, fi31, v1)
                                                var v38 float32 = 0.0
                                                _ = v38
                                            }
                                        }
                                        v1 += 1
                                        _ = v32
                                    case 32767:
                                        var v39 string = "hello"
                                        _ = v39
                                    }
                                    v4[210] = ""
                                    _ = fi31
                                }
                                var v40 bool = (!false)
                                _ = v40
                            }
                        }
                        fmt.Println(v1, v11)
                        _ = v28
                    }
                    var v41 int8 = -128
                    _ = len(v2)
                    _ = v41
                }
                _ = fi18
            }
        default:
            for fi42 := 0; fi42 < 8; fi42++ {
                v6 = "foo bar"
                _ = fi42
            }
        }
    case 4294967295:
        if (42 >= (^(v7))) {
            v43 := GenS5{F0: 0, F1: (^(v7))}
            v44 := GenS6{F0: -1, F1: -1}
            switch 403 {
            case 9223372036854775807:
                if false {
                    _ = len(v2)
                    if false {
                        for fi45 := 0; fi45 < 6; fi45++ {
                            var v46 int32 = v8
                            var v47 string = "foo bar"
                            _ = fi45
                            _ = v46
                            _ = v47
                        }
                        v48 := GenS2{F0: 1, F1: v7, F2: int(v1)}
                        _ = v48
                    } else {
                        if false {
                            var v49 uint = 18446744073709551615
                            _ = v49
                        }
                    }
                }
            case -2147483648:
                v3[("\n\t" + "foo bar")] = "foo bar"
            case 1:
                for fi50 := 0; fi50 < 17; fi50++ {
                    if (true || false) {
                        fmt.Println(v8)
                        v51 := GenS1{F0: 2147483648, F1: -110, F2: (fi50 / (((fi50 - (fi50 >> ((177) & 63)))) | 1))}
                        v2 = append(v2, 42)
                        _ = v51
                    }
                    if (fi50 > fi50) {
                        v4[2147483647] = v6
                    } else {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = len(v2)
                        if (!false) {
                            v2 = append(v2, 2)
                        } else {
                            for fi52 := 0; fi52 < 7; fi52++ {
                                v8 += -927
                                v53 := GenS1{F0: 663, F1: (v1 / ((925) | 1)), F2: (fi50 | int(v7))}
                                var v54 uint8 = 60
                                if false {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if false {
                                        if (255 > v54) {
                                            if ((-2.25 == ((((((1e10 + float32(0.0)) / 0.1) / -2.25) + ((3.14159 / 0.1) / 1e10)) + 1.5) * (0.1 * ((0.1 * 1.5) * (0.1 / 1.0))))) && (!(!true))) {
                                                var v55 uint32 = 1
                                                _ = v55
                                            } else {
                                                fmt.Println(v11, fi50)
                                                var v56 float32 = (1e10 + 3.14159)
                                                if false {
                                                    if false {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        var v57 string = ""
                                                        _ = v57
                                                    }
                                                }
                                                _ = v56
                                            }
                                            fmt.Println(v1, v54)
                                            v58 := GenS5{F0: v54, F1: 42}
                                            var v59 uint8 = 31
                                            _ = v58
                                            _ = v59
                                        }
                                        v54 -= v54
                                        if false {
                                            switch "\n\t" {
                                            case "\n\t":
                                                v60 := v4[2147483648]
                                                _ = v60
                                            }
                                            v3[(("a" + ("hello" + v6)) + ("hello" + ""))] = v6
                                        } else {
                                            var v61 uint32 = 2147483647
                                            switch (fi50 ^ (fi52 ^ (fi52 << ((65535) & 63)))) {
                                            case 128:
                                                fmt.Println(v54, v1)
                                            }
                                            _ = v61
                                        }
                                    }
                                    if false {
                                        fmt.Println(fi52, v1)
                                    }
                                }
                                var v62 float32 = 0.1
                                _ = fi52
                                _ = v53
                                _ = v54
                                _ = v62
                            }
                        }
                    }
                    _ = fi50
                }
            default:
                var v63 uint = 2
                _ = v63
            }
            _ = v43
            _ = v44
        } else {
            switch v6 {
            case "foo bar":
                fmt.Println(v6, v7)
            default:
                if true {
                    v8 = (v1 % ((v8) | 1))
                    if true {
                        v64 := GenS7{F0: 1, F1: -973, F2: 65535}
                        _ = v64
                    }
                }
            }
        }
    default:
        for fi65 := 0; fi65 < 4; fi65++ {
            if true {
                for fi66 := 0; fi66 < 3; fi66++ {
                    var v67 float32 = 3.14159
                    v68 := GenS1{F0: 2147483647, F1: v1, F2: fi66}
                    _ = fi66
                    _ = v67
                    _ = v68
                }
            }
            _ = fi65
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
