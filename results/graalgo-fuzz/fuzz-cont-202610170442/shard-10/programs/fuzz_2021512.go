package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 uint }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int16; F1 int64; F2 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 uint; F1 int8 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 string; F1 uint64; F2 uint32 }
type GenS9 struct { F0 uint8 }

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]bool)
    v3 := []string{"a", ("hello" + "foo bar"), "a", "\n\t"}
    var v4 int32 = -352
    var v5 float32 = (0.1 * ((-2.25 + float32(3.14159)) - 1.5))
    var v6 string = ""
    var v7 int64 = 0
    if true {
        if true {
            if ((!true) && (true && true)) {
                if (!false) {
                    switch int(v7) {
                    case 42:
                        v4 = v4
                    case 2147483647:
                        delete(v2, 1)
                    default:
                        _ = len(v2)
                    }
                    for fi8 := 0; fi8 < 11; fi8++ {
                        var v9 int = fi8
                        var v10 int16 = -32768
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                } else {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        v7 += (^(v7))
                        v7 -= (v7 | 42)
                        _ = fi11
                    }
                    delete(v2, 2)
                    for fi12 := 0; fi12 < 19; fi12++ {
                        if (!(-1 != fi12)) {
                            var v13 uint64 = 18446744073709551615
                            _ = v13
                        }
                        _ = len(v3)
                        for fi14 := 0; fi14 < 4; fi14++ {
                            if false {
                                switch fi14 {
                                case 42:
                                    switch ((((("" + v6) + v6) + v6) + v6) + v6) {
                                    case "":
                                        var v15 uint8 = 1
                                        _ = v15
                                    case "\n\t":
                                        delete(v2, 4294967295)
                                    }
                                case -2147483648:
                                    if (!(!(1 > 255))) {
                                        v16 := GenS1{F0: v1, F1: ("hello" + v6), F2: 2}
                                        v17 := GenS2{F0: true, F1: 1}
                                        v6 = "\n\t"
                                        _ = v16
                                        _ = v17
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v7 -= (v7 | (v7 % ((1) | 1)))
                                        var v18 uint = 65535
                                        _ = v18
                                    }
                                }
                                v7 = v7
                                if (!(true && (false || ((!(!true)) && false)))) {
                                    v1 += 42
                                }
                            } else {
                                v6 = "hello"
                            }
                            v4 = -2147483648
                            if len(v3) > 0 {
                                v3[0] = v6
                            }
                            _ = fi14
                        }
                        _ = fi12
                    }
                }
            } else {
                if false {
                    var v19 int16 = 75
                    _ = v19
                } else {
                    v20 := GenS3{F0: -1, F1: (v7 + 42), F2: (v1 / (((v1 - (v1 - 1))) | 1))}
                    _ = v20
                }
                switch "hello" {
                case "\n\t":
                    v5 += (float32((3.14159 - (1e10 + 1e10))) + v5)
                case "foo bar":
                    v6 = ("a" + "a")
                case "":
                    var v21 uint8 = 127
                    _ = v21
                }
                var v22 int16 = -32768
                _ = v22
            }
        } else {
            if true {
                var v23 int = -732
                _ = v23
            }
        }
        if (206 != 50) {
            var v24 int8 = (v1 >> ((2) & 63))
            _ = v24
        }
        v1 += (v1 >> ((65535) & 63))
    } else {
        delete(v2, 1)
        for fi25 := 0; fi25 < 4; fi25++ {
            v26 := GenS4{F0: 128}
            _ = fi25
            _ = v26
        }
        var v27 int32 = -1
        _ = v27
    }
    v6 = v6
    for fi28 := 0; fi28 < 7; fi28++ {
        if (!(!(false && ((-(fi28)) == (fi28 ^ (fi28 ^ (fi28 - fi28))))))) {
            v6 = (v6 + v6)
            v29 := GenS5{F0: v4, F1: 18446744073709551615}
            v30 := GenS6{F0: 4294967295, F1: v1}
            _ = v29
            _ = v30
        } else {
            for fi31 := 0; fi31 < 19; fi31++ {
                var v32 float64 = 1.5
                var v33 uint = 805
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        if (!false) {
            var v34 float64 = float64(v5)
            _ = v34
        }
        v6 = v6
        _ = fi28
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        if (!true) {
            var v36 int64 = (v7 / ((1) | 1))
            fmt.Println(v1)
            _ = v36
        }
        v37 := GenS3{F0: 0, F1: v7, F2: (v1 + 1)}
        _ = fi35
        _ = v37
    }
    if (!(!false)) {
        v38 := v2[0]
        var v39 uint16 = 845
        _ = v38
        _ = v39
    }
    if (-1 > v1) {
        for fi40 := 0; fi40 < 4; fi40++ {
            v41 := GenS2{F0: true, F1: 643}
            _ = fi40
            _ = v41
        }
        var v42 uint32 = 2
        v5 = (v5 - 0.0)
        _ = v42
    } else {
        v43 := GenS5{F0: -998, F1: 14}
        switch "\n\t" {
        case "\n\t":
            v44 := GenS4{F0: 42}
            _ = v44
        case "hello":
            if (42 > 0) {
                switch 9223372036854775807 {
                case -1:
                    var v45 uint64 = 155
                    _ = v45
                case 42:
                    var v46 string = v6
                    _ = v46
                }
            } else {
                var v47 int8 = (v1 * 100)
                v47 = (v47 / (((v47 | (v47 * v1))) | 1))
                v6 = "foo bar"
                _ = v47
            }
        default:
            if (4294967295 >= 2) {
                switch "foo bar" {
                case "\n\t":
                    switch -336 {
                    case 1:
                        var v48 string = v6
                        _ = v48
                    case -538:
                        v1 = v1
                    }
                case "a":
                    switch "\n\t" {
                    case "hello":
                        for fi49 := 0; fi49 < 20; fi49++ {
                            v2[42] = true
                            v50 := v2[591]
                            delete(v2, 2)
                            _ = fi49
                            _ = v50
                        }
                    case "foo bar":
                        for fi51 := 0; fi51 < 10; fi51++ {
                            switch -973 {
                            case -2147483648:
                                for fi52 := 0; fi52 < 17; fi52++ {
                                    fmt.Println(v1, v6)
                                    if (!(287 != 2)) {
                                        fmt.Println(fi52, fi52)
                                    }
                                    if (!true) {
                                        v1 -= -91
                                        v53 := GenS4{F0: 127}
                                        _ = v53
                                    }
                                    var v54 bool = (255 != 2)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi52
                                    _ = v54
                                }
                            }
                            _ = fi51
                        }
                    default:
                        for fi55 := 0; fi55 < 8; fi55++ {
                            switch (fi55 + 600) {
                            case -2147483648:
                                if false {
                                    switch (fi55 >> ((941) & 63)) {
                                    case 9223372036854775807:
                                        if false {
                                            v1 = (-(v1))
                                            fmt.Println(fi55)
                                            var v56 uint32 = 2147483648
                                            _ = v56
                                        } else {
                                            v5 = 0.1
                                            delete(v2, 4294967295)
                                            v5 -= 3.14159
                                        }
                                    case -852:
                                        v6 = v6
                                    }
                                }
                            default:
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = v6
                                    }
                                    var v57 string = "\n\t"
                                    _ = v57
                                } else {
                                    var v58 uint8 = 0
                                    v1 += v1
                                    switch v6 {
                                    case "":
                                        v59 := GenS7{F0: int64(v4), F1: v1}
                                        _ = v59
                                    case "a":
                                        fmt.Println(v7, fi55)
                                    case "\n\t":
                                        var v60 uint32 = 143
                                        _ = v60
                                    default:
                                        switch "\n\t" {
                                        case "hello":
                                            v7 = (v7 | v7)
                                        case "foo bar":
                                            switch "" {
                                            case "\n\t":
                                                if false {
                                                    if len(v3) > 0 {
                                                        v3[0] = ("a" + "\n\t")
                                                    }
                                                } else {
                                                    switch 2147483647 {
                                                    case 2147483647:
                                                        for fi61 := 0; fi61 < 9; fi61++ {
                                                            var v62 uint64 = 18446744073709551615
                                                            v63 := GenS8{F0: v6, F1: v62, F2: 681}
                                                            _ = fi61
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                    default:
                                                        if false {
                                                            v64 := GenS9{F0: v58}
                                                            var v65 uint64 = 219
                                                            v2[4294967295] = true
                                                            _ = v64
                                                            _ = v65
                                                        } else {
                                                            var v66 int64 = 9223372036854775807
                                                            _ = v66
                                                        }
                                                    }
                                                    var v67 int32 = 0
                                                    _ = v67
                                                }
                                            case "":
                                                if (127 <= (fi55 / ((fi55) | 1))) {
                                                    v68 := GenS5{F0: v4, F1: 437}
                                                    var v69 uint = 42
                                                    _ = v68
                                                    _ = v69
                                                } else {
                                                    for fi70 := 0; fi70 < 3; fi70++ {
                                                        _ = len(v2)
                                                        _ = fi70
                                                    }
                                                    var v71 uint = 255
                                                    _ = v71
                                                }
                                            case "hello":
                                                if (v5 != 0.0) {
                                                    var v72 bool = (153 == uint(v58))
                                                    fmt.Println(v58)
                                                    _ = v72
                                                } else {
                                                    if len(v3) > 0 {
                                                        v3[0] = ((v6 + ("" + v6)) + v6)
                                                    }
                                                }
                                            }
                                        default:
                                            v58 = 127
                                        }
                                    }
                                    _ = v58
                                }
                            }
                            _ = fi55
                        }
                    }
                case "":
                    v73 := GenS6{F0: 65535, F1: int8(v7)}
                    _ = v73
                default:
                    fmt.Println(v1, v5, v4)
                }
            } else {
                var v74 int16 = 42
                v6 = v6
                v1 -= int8(v4)
                _ = v74
            }
        }
        v75 := GenS6{F0: 2, F1: -70}
        _ = v43
        _ = v75
    }
    switch ((v6 + ("\n\t" + "foo bar")) + (v6 + v6)) {
    case "\n\t":
        if (42 != 65535) {
            v2[2147483648] = true
        }
    default:
        v2[652] = (v6 < v6)
    }
    var v76 int8 = -1
    var v77 uint32 = 1
    var v78 int8 = (v1 + v76)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v76
    _ = v77
    _ = v78
}
