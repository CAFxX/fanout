package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint32; F2 bool }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 uint16; F2 int64 }
type GenS8 struct { F0 int32; F1 int64 }
type GenS9 struct { F0 uint8; F1 uint16 }
type GenS10 struct { F0 int8; F1 int }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint; F1 int16 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 int8; F1 uint32; F2 string }
type GenS16 struct { F0 uint32; F1 uint64; F2 int64 }

func main() {
    var v1 string = ""
    var v2 float32 = 1e10
    v3 := make(map[int]bool)
    var v4 uint32 = 72
    switch v1 {
    case "hello":
        var v5 int64 = 42
        _ = v5
    }
    v6 := []int{}
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(v2, fi8)
        v9 := GenS2{F0: (fi8 - -1)}
        _ = fi8
        _ = v9
    }
    if false {
        v10 := GenS3{F0: false}
        v11 := GenS3{F0: (true || (false && true))}
        fmt.Println(v1)
        _ = v10
        _ = v11
    }
    v2 = float32((1.5 / (3.14159 + (0.0 / ((1.5 - 0.1) - (0.1 / 1e10))))))
    for fi12 := 0; fi12 < 3; fi12++ {
        if (false || true) {
            v2 = 0.0
            v1 = (v1 + ("\n\t" + v1))
            v13 := GenS2{F0: (fi12 ^ (fi12 << ((1) & 63)))}
            _ = v13
        } else {
            switch "a" {
            case "a":
                if (2 <= 18446744073709551615) {
                    if false {
                        v2 -= float32((1.5 + 1.5))
                    } else {
                        for fi14 := 0; fi14 < 3; fi14++ {
                            var v15 float64 = -2.25
                            _ = fi14
                            _ = v15
                        }
                    }
                    for fi16 := 0; fi16 < 6; fi16++ {
                        switch 128 {
                        case 0:
                            var v17 bool = false
                            _ = v17
                        default:
                            if (true && (true && false)) {
                                if false {
                                    var v18 uint16 = 1
                                    _ = v18
                                }
                                v2 += 3.14159
                                switch 128 {
                                case 4294967295:
                                    var v19 uint8 = 1
                                    _ = v19
                                case 42:
                                    fmt.Println(v2, fi12)
                                case 128:
                                    delete(v3, (fi16 * fi16))
                                default:
                                    switch fi16 {
                                    case -922:
                                        if (!true) {
                                            v20 := GenS4{F0: 1}
                                            fmt.Println(v4)
                                            v21 := GenS2{F0: (fi16 >> ((4294967295) & 63))}
                                            _ = v20
                                            _ = v21
                                        } else {
                                            if (-128 != -1) {
                                                switch "a" {
                                                case "a":
                                                    if ((!false) || (255 > 65535)) {
                                                        v3[(fi16 + -9223372036854775808)] = (!(32767 <= 353))
                                                        var v22 int = 127
                                                        var v23 int = v22
                                                        var v24 bool = (((1.5 / float64((v2 / (v2 / ((3.14159 * float32((((1e10 / ((float64(((1e10 / (float32(((3.14159 / ((((1.5 - (-2.25 / 1.0)) - 0.1) / ((-2.25 + 1e10) + 0.0)) * 1.5)) / float64(((float32(-2.25) + float32(1e10)) * (v2 + v2))))) * v2)) / v2)) / 0.1) / -2.25)) + 1.5) * float64(v2)))) * v2))))) <= (float64(0.0) * ((0.0 * (((0.0 / 1e10) * ((-2.25 + 0.1) * ((3.14159 + 1.5) + 1e10))) / 0.1)) + 3.14159))) || false)
                                                        var v25 int64 = 9223372036854775807
                                                        _ = v22
                                                        _ = v23
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                }
                                                var v26 int64 = -944
                                                v27 := GenS5{F0: (v26 << ((2) & 63)), F1: v4, F2: (v4 >= (v4 - v4))}
                                                _ = v26
                                                _ = v27
                                            } else {
                                                var v28 int8 = -1
                                                _ = v28
                                            }
                                            v4 -= v4
                                        }
                                    }
                                }
                            } else {
                                if true {
                                    fmt.Println(v2, v2)
                                    if (!true) {
                                        v4 -= 1
                                        if (4294967295 == 4294967295) {
                                            var v29 string = ((v1 + (v1 + v1)) + v1)
                                            _ = v29
                                        }
                                        v30 := GenS6{F0: (v1 + v1)}
                                        _ = v30
                                    }
                                } else {
                                    var v31 uint16 = 42
                                    _ = v31
                                }
                                v2 -= float32(float64(v2))
                                v4 = v4
                            }
                        }
                        var v32 int8 = -45
                        _ = fi16
                        _ = v32
                    }
                    if (false && (!(127 != (fi12 ^ (fi12 % ((fi12) | 1)))))) {
                        var v33 int = 128
                        var v34 int32 = 1
                        switch ((v1 + v1) + v1) {
                        case "hello":
                            v35 := GenS7{F0: -1, F1: 255, F2: 0}
                            _ = v35
                        case "foo bar":
                            if (!true) {
                                var v36 int16 = -870
                                _ = v36
                            }
                        }
                        _ = v33
                        _ = v34
                    } else {
                        v2 = (0.0 + (1.5 + 1e10))
                        if (v2 <= float32((float64(v2) - 3.14159))) {
                            v2 = v2
                            if true {
                                v2 = 1.5
                                v1 = ""
                            } else {
                                v37 := GenS2{F0: (fi12 << ((800) & 63))}
                                var v38 int = (fi12 * (fi12 & (fi12 << ((255) & 63))))
                                _ = v37
                                _ = v38
                            }
                            for fi39 := 0; fi39 < 5; fi39++ {
                                var v40 int64 = -9223372036854775808
                                if true {
                                    var v41 int8 = -1
                                    if true {
                                        v42 := GenS4{F0: uint64(v4)}
                                        v4 += (v4 * (v4 * 65535))
                                        v43 := GenS3{F0: (1 < 255)}
                                        _ = v42
                                        _ = v43
                                    }
                                    fmt.Println(fi39, v4)
                                    _ = v41
                                } else {
                                    v44 := GenS2{F0: (fi12 & fi12)}
                                    var v45 int32 = 1
                                    v45 = v45
                                    _ = v44
                                    _ = v45
                                }
                                switch (^(fi12)) {
                                case -9223372036854775808:
                                    switch 2147483647 {
                                    case -2147483648:
                                        fmt.Println(fi12)
                                    }
                                }
                                _ = fi39
                                _ = v40
                            }
                        } else {
                            v6 = append(v6, 602)
                            if true {
                                for fi46 := 0; fi46 < 4; fi46++ {
                                    v47 := GenS8{F0: -2147483648, F1: -1}
                                    v48 := GenS9{F0: 127, F1: 2}
                                    switch (v1 + "a") {
                                    case "":
                                        v2 = 0.1
                                        v49 := GenS4{F0: 265}
                                        v50 := GenS10{F0: -39, F1: -1}
                                        _ = v49
                                        _ = v50
                                    case "\n\t":
                                        switch v1 {
                                        case "a":
                                            fmt.Println(v4)
                                        case "":
                                            fmt.Println(fi46, fi46)
                                        default:
                                            switch -1 {
                                            case 42:
                                                if false {
                                                    fmt.Println(fi46, v1)
                                                }
                                            case 2147483647:
                                                if (1 < 4294967295) {
                                                    var v51 uint64 = 0
                                                    v4 = (^(v4))
                                                    _ = v51
                                                }
                                            default:
                                                if (229 > 65535) {
                                                    if (!(496 <= 65535)) {
                                                        _ = len(v3)
                                                        v52 := GenS6{F0: v1}
                                                        if (42 == 782) {
                                                            var v53 int32 = 1
                                                            var v54 uint = 1
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        _ = v52
                                                    }
                                                    if len(v6) > 0 {
                                                        v6[0] = fi46
                                                    }
                                                } else {
                                                    v55 := GenS11{F0: (!true)}
                                                    _ = v55
                                                }
                                            }
                                        }
                                    case "a":
                                        v3[(fi46 - 127)] = (!false)
                                    default:
                                        if false {
                                            v4 -= v4
                                            if len(v6) > 0 {
                                                v6[0] = fi46
                                            }
                                        } else {
                                            fmt.Println(v4)
                                            v2 = (v2 - (float32(0.0) / v2))
                                            if true {
                                                v56 := GenS12{F0: false}
                                                var v57 int = fi46
                                                _ = len(v3)
                                                _ = v56
                                                _ = v57
                                            } else {
                                                v58 := GenS13{F0: 848, F1: -507}
                                                switch 2147483647 {
                                                case 1:
                                                    fmt.Println(fi12, fi12)
                                                default:
                                                    v59 := GenS12{F0: (-1 < fi46)}
                                                    _ = v59
                                                }
                                                _ = v58
                                            }
                                        }
                                    }
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                                switch fi12 {
                                case 4294967295:
                                    v1 = "hello"
                                default:
                                    fmt.Println(fi12, v1)
                                }
                                var v60 int = fi12
                                _ = v60
                            } else {
                                v61 := GenS14{F0: 2}
                                _ = v61
                            }
                            v4 = 65535
                        }
                    }
                }
            case "hello":
                v4 -= (v4 >> ((4294967295) & 63))
            default:
                for fi62 := 0; fi62 < 4; fi62++ {
                    fmt.Println(fi12, fi62, v1)
                    v4 += v4
                    _ = fi62
                }
            }
        }
        for fi63 := 0; fi63 < 3; fi63++ {
            switch "hello" {
            case "foo bar":
                v64 := GenS15{F0: 1, F1: 2147483648, F2: v1}
                _ = v64
            case "\n\t":
                fmt.Println(fi12)
            default:
                if (!false) {
                    delete(v3, fi63)
                    if (127 == 42) {
                        var v65 uint = 65535
                        fmt.Println(v2, v4)
                        v1 = "hello"
                        _ = v65
                    }
                }
            }
            _ = fi63
        }
        if len(v6) > 0 {
            v6[0] = fi12
        }
        _ = fi12
    }
    v66 := GenS16{F0: v4, F1: 161, F2: 9223372036854775807}
    v67 := GenS13{F0: 4294967295, F1: 42}
    switch -910 {
    case -244:
        if (1 != 2147483647) {
            fmt.Println(v1, v1)
            v4 -= v4
        } else {
            v1 = "foo bar"
            switch 9223372036854775807 {
            case 4294967295:
                for fi68 := 0; fi68 < 14; fi68++ {
                    if true {
                        var v69 int8 = -95
                        if (v2 > (3.14159 / 1e10)) {
                            v6 = append(v6, 9223372036854775807)
                            v1 = ""
                        }
                        _ = v69
                    } else {
                        var v70 int32 = 195
                        for fi71 := 0; fi71 < 6; fi71++ {
                            v6 = append(v6, fi71)
                            _ = fi71
                        }
                        _ = v70
                    }
                    _ = fi68
                }
            default:
                v72 := GenS14{F0: 1}
                _ = v72
            }
        }
    case -2147483648:
        v1 = v1
    default:
        switch (v1 + ("a" + v1)) {
        case "\n\t":
            v73 := GenS12{F0: true}
            _ = v73
        }
    }
    for fi74 := 0; fi74 < 8; fi74++ {
        v4 -= 874
        _ = fi74
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v66
    _ = v67
    _ = v7
}
