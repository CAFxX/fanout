package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8; F1 int16; F2 uint }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int32; F1 int64; F2 uint32 }
type GenS10 struct { F0 int8; F1 uint32 }

func main() {
    var v1 int16 = 0
    v2 := make(map[int]uint32)
    v3 := []uint32{42, 47, 840, 2147483647}
    v4 := make(map[string]bool)
    var v5 int16 = v1
    if (!true) {
        switch -2147483648 {
        case 922:
            var v6 int8 = 31
            _ = v6
        case 2147483647:
            v7 := GenS1{F0: 122}
            _ = v7
        case 4294967295:
            switch int(v5) {
            case 2147483647:
                switch 128 {
                case 32768:
                    v8 := GenS2{F0: 2147483647}
                    _ = v8
                }
            case -44:
                for fi9 := 0; fi9 < 13; fi9++ {
                    if (480 >= 24) {
                        v10 := v4["foo bar"]
                        v11 := GenS3{F0: 2, F1: v1, F2: 592}
                        _ = v10
                        _ = v11
                    }
                    _ = fi9
                }
            case -757:
                if true {
                    for fi12 := 0; fi12 < 9; fi12++ {
                        if (false && (((464 != 4294967295) && false) && ((float64(1.5) > 1e10) && (0 >= -1)))) {
                            var v13 float64 = ((((0.0 + (((3.14159 + 0.1) + float64((((0.1 - 1.5) + 1.5) * float32((float64(1e10) / (((((0.1 - 0.0) * ((-2.25 + (1e10 + ((0.1 + (3.14159 / 1.0)) / ((0.0 / 0.1) + 0.1)))) / -2.25)) - 0.0) + (1e10 / 0.1)) - (1.5 + float64(float32((0.1 / 1.5)))))))))) * 3.14159)) + 0.0) * 3.14159) / float64(1e10))
                            var v14 float32 = 1.5
                            _ = v13
                            _ = v14
                        }
                        if true {
                            for fi15 := 0; fi15 < 18; fi15++ {
                                if true {
                                    v16 := GenS4{F0: 255, F1: 0}
                                    switch "hello" {
                                    case "foo bar":
                                        var v17 float32 = 0.0
                                        _ = v17
                                    case "\n\t":
                                        fmt.Println(v5, v5)
                                    case "hello":
                                        v5 = v5
                                        v3 = append(v3, 42)
                                    }
                                    _ = v16
                                } else {
                                    var v18 int64 = 1
                                    if false {
                                        v18 = 0
                                    }
                                    var v19 uint8 = 224
                                    _ = v18
                                    _ = v19
                                }
                                v5 = 729
                                _ = fi15
                            }
                            fmt.Println(v1)
                        }
                        for fi20 := 0; fi20 < 7; fi20++ {
                            v21 := GenS4{F0: 311, F1: 42}
                            _ = fi20
                            _ = v21
                        }
                        _ = fi12
                    }
                    switch "" {
                    case "hello":
                        if (107 == 42) {
                            var v22 int64 = 42
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 18; fi23++ {
                                for fi24 := 0; fi24 < 1; fi24++ {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v25 := v2[fi23]
                                        _ = v25
                                    }
                                    _ = fi24
                                }
                                var v26 float64 = 1.5
                                _ = fi23
                                _ = v26
                            }
                        }
                    case "":
                        var v27 uint16 = 847
                        _ = v27
                    case "\n\t":
                        v28 := GenS5{F0: (18446744073709551615 == 894), F1: true}
                        _ = v28
                    default:
                        if (42 > 386) {
                            v1 -= 1000
                        } else {
                            switch 1 {
                            case -295:
                                v29 := GenS6{F0: 42}
                                _ = v29
                            case -1:
                                v5 -= -800
                            case 2147483647:
                                if false {
                                    var v30 uint64 = 1
                                    if true {
                                        v5 = v5
                                    } else {
                                        var v31 int64 = 9223372036854775807
                                        _ = v31
                                    }
                                    var v32 int = 690
                                    _ = v30
                                    _ = v32
                                } else {
                                    v1 -= (-(v5))
                                    switch 623 {
                                    case 4294967295:
                                        if (!false) {
                                            v33 := GenS7{F0: -9223372036854775808}
                                            fmt.Println(v5)
                                            if (float64(1e10) <= 0.0) {
                                                if true {
                                                    v34 := GenS5{F0: false, F1: (-1 < 127)}
                                                    _ = v34
                                                }
                                                v5 = v1
                                                for fi35 := 0; fi35 < 15; fi35++ {
                                                    for fi36 := 0; fi36 < 7; fi36++ {
                                                        fmt.Println(fi36, v1)
                                                        var v37 float32 = ((1e10 - (1.5 - -2.25)) * (1.5 + 1.5))
                                                        if (100 <= 127) {
                                                            v1 = v5
                                                            v37 = v37
                                                        }
                                                        v1 = v1
                                                        _ = fi36
                                                        _ = v37
                                                    }
                                                    var v38 uint8 = 127
                                                    v39 := GenS8{F0: false}
                                                    _ = fi35
                                                    _ = v38
                                                    _ = v39
                                                }
                                            } else {
                                                var v40 string = "a"
                                                if (-1 > -1) {
                                                    fmt.Println(v40)
                                                } else {
                                                    v41 := GenS9{F0: 42, F1: 269, F2: 2147483647}
                                                    for fi42 := 0; fi42 < 2; fi42++ {
                                                        if (2 >= 65535) {
                                                            v43 := GenS2{F0: (fi42 << ((255) & 63))}
                                                            switch 421 {
                                                            case 128:
                                                                v44 := GenS1{F0: 2}
                                                                _ = v44
                                                            case -375:
                                                                var v45 uint32 = 4294967295
                                                                _ = v45
                                                            case 125:
                                                                _ = len(v3)
                                                            }
                                                            _ = v43
                                                        }
                                                        v46 := GenS4{F0: 42, F1: 416}
                                                        v3 = append(v3, 1)
                                                        _ = fi42
                                                        _ = v46
                                                    }
                                                    if true {
                                                        var v47 int = 127
                                                        _ = v47
                                                    }
                                                    _ = v41
                                                }
                                                v48 := v4[("foo bar" + ("a" + (((v40 + ("foo bar" + "")) + v40) + (((v40 + (v40 + "")) + "\n\t") + "hello"))))]
                                                _ = v40
                                                _ = v48
                                            }
                                            _ = v33
                                        }
                                    }
                                }
                            }
                            var v49 uint8 = 28
                            _ = v49
                        }
                    }
                    var v50 int8 = 100
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 20; fi51++ {
                        v1 -= -32768
                        _ = fi51
                    }
                }
            }
        }
        for fi52 := 0; fi52 < 20; fi52++ {
            switch "foo bar" {
            case "":
                if (true || false) {
                    var v53 int16 = -163
                    switch fi52 {
                    case 32768:
                        fmt.Println(fi52, v53, fi52)
                    }
                    _ = v53
                }
            case "a":
                for fi54 := 0; fi54 < 12; fi54++ {
                    if true {
                        delete(v2, fi54)
                        var v55 uint16 = 1
                        _ = v55
                    }
                    var v56 int32 = 2147483647
                    _ = fi54
                    _ = v56
                }
            }
            _ = fi52
        }
    }
    var v57 float32 = float32(((1e10 - (3.14159 + (float64((-2.25 * (float32(float64(0.0)) * float32(-2.25)))) / 1e10))) - (0.1 - 0.1)))
    var v58 float64 = 0.0
    switch ((("hello" + ("\n\t" + "\n\t")) + ("foo bar" + "foo bar")) + "hello") {
    case "hello":
        var v59 int8 = 100
        _ = v59
    }
    for fi60 := 0; fi60 < 12; fi60++ {
        v4["\n\t"] = (252 != 18446744073709551615)
        _ = fi60
    }
    if (!false) {
        var v61 uint8 = 69
        var v62 uint16 = 0
        var v63 int32 = 0
        _ = v61
        _ = v62
        _ = v63
    }
    fmt.Println(v1, v5)
    v58 -= ((1e10 - ((v58 * v58) + v58)) * (v58 + (v58 + ((v58 + 1.5) - (3.14159 * 3.14159)))))
    _ = len(v4)
    for fi64 := 0; fi64 < 8; fi64++ {
        var v65 float32 = v57
        _ = fi64
        _ = v65
    }
    v66 := GenS10{F0: -67, F1: 0}
    v67 := GenS7{F0: 42}
    if (65535 >= 4294967295) {
        var v68 int8 = -38
        _ = v68
    }
    switch 32767 {
    case 32767:
        var v69 int16 = 1000
        _ = v69
    case 128:
        v2[9223372036854775807] = 182
    case 213:
        var v70 int16 = v5
        _ = v70
    default:
        for fi71 := 0; fi71 < 15; fi71++ {
            v58 -= v58
            _ = fi71
        }
    }
    var v72 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v66
    _ = v67
    _ = v72
}
