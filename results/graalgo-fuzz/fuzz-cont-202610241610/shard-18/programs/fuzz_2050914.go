package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int64; F1 uint64; F2 bool }
type GenS5 struct { F0 string; F1 uint8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    var v3 int = 603
    v4 := []uint32{514}
    switch v3 {
    case 989:
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 7; fi6++ {
                v7 := GenS1{F0: 1, F1: 4294967295, F2: 46}
                v8 := v2[65535]
                fmt.Println(v8, v3)
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi5
        }
    case 9223372036854775807:
        v9 := GenS1{F0: 127, F1: 18446744073709551615, F2: 42}
        _ = v9
    }
    v3 += 32768
    if (((2 == 253) && (176 >= 255)) || false) {
        _ = len(v4)
        switch ("" + ("\n\t" + (("a" + "a") + "a"))) {
        case "hello":
            switch v3 {
            case 4294967295:
                fmt.Println(v3)
            }
        default:
            v3 -= (v3 + (v3 % ((127) | 1)))
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v3 = -9223372036854775808
        var v10 uint = 4294967295
        _ = v10
    }
    v11 := GenS2{F0: (-(v3))}
    if false {
        var v12 string = ("a" + "foo bar")
        _ = v12
    } else {
        v3 = v3
        v2[862] = v3
        var v13 uint16 = 255
        _ = v13
    }
    if false {
        var v14 string = ""
        _ = v14
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := v2[4294967295]
            _ = fi15
            _ = v16
        }
    }
    _ = len(v4)
    switch "hello" {
    case "\n\t":
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := GenS1{F0: 255, F1: 527, F2: 0}
            _ = fi17
            _ = v18
        }
    case "":
        for fi19 := 0; fi19 < 20; fi19++ {
            v3 += fi19
            if true {
                if true {
                    v3 -= (fi19 | -570)
                    for fi20 := 0; fi20 < 15; fi20++ {
                        v3 = 127
                        _ = fi20
                    }
                    var v21 int16 = 32767
                    _ = v21
                } else {
                    var v22 int = (fi19 >> ((1) & 63))
                    v4 = append(v4, 2)
                    switch fi19 {
                    case -9223372036854775808:
                        v22 = -517
                    default:
                        if (0.1 >= (0.1 - (((0.0 + 0.1) + 1.5) * ((1.5 - 0.0) + -2.25)))) {
                            v3 = v22
                            switch "foo bar" {
                            case "":
                                fmt.Println(v3, v22)
                            case "\n\t":
                                if ((1e10 - -2.25) >= float64(-2.25)) {
                                    var v23 bool = true
                                    var v24 uint8 = 128
                                    for fi25 := 0; fi25 < 15; fi25++ {
                                        var v26 float32 = 3.14159
                                        var v27 int32 = 1
                                        var v28 string = "hello"
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = v23
                                    _ = v24
                                }
                            case "hello":
                                var v29 uint16 = 65535
                                _ = v29
                            default:
                                var v30 uint = 12
                                _ = v30
                            }
                        } else {
                            var v31 int = v3
                            _ = v31
                        }
                    }
                    _ = v22
                }
                v32 := GenS2{F0: (v3 / (((v3 - (v3 | v3))) | 1))}
                _ = v32
            } else {
                for fi33 := 0; fi33 < 15; fi33++ {
                    if true {
                        var v34 int8 = 0
                        _ = len(v2)
                        _ = v34
                    }
                    v35 := GenS3{F0: 640, F1: (v3 >> ((42) & 63))}
                    _ = fi33
                    _ = v35
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                for fi36 := 0; fi36 < 4; fi36++ {
                    fmt.Println(fi36, v3, fi36)
                    _ = fi36
                }
            }
            fmt.Println(fi19, fi19)
            _ = fi19
        }
    case "foo bar":
        if false {
            fmt.Println(v3)
            v37 := GenS4{F0: -100, F1: 801, F2: true}
            _ = v37
        }
    default:
        var v38 float32 = (float32(1e10) + 1.5)
        _ = v38
    }
    if (127 <= 53) {
        delete(v2, 2147483648)
        if true {
            switch v3 {
            case -649:
                delete(v1, 42)
            case -571:
                switch (("" + "hello") + ("" + "\n\t")) {
                case "":
                    switch (v3 % (((v3 >> ((4294967295) & 63))) | 1)) {
                    case -2147483648:
                        _ = len(v2)
                    default:
                        for fi39 := 0; fi39 < 6; fi39++ {
                            fmt.Println(v3)
                            _ = fi39
                        }
                    }
                case "\n\t":
                    if ("foo bar" <= "hello") {
                        for fi40 := 0; fi40 < 15; fi40++ {
                            switch "\n\t" {
                            case "foo bar":
                                v3 = v3
                            case "\n\t":
                                v3 -= v3
                            default:
                                if false {
                                    v41 := GenS3{F0: -516, F1: fi40}
                                    var v42 string = ("a" + ("a" + "a"))
                                    fmt.Println(v3)
                                    _ = v41
                                    _ = v42
                                }
                            }
                            if ((3.14159 > (1.5 * 1e10)) || true) {
                                for fi43 := 0; fi43 < 18; fi43++ {
                                    var v44 float32 = (0.0 - (0.0 - (1e10 * 1.5)))
                                    v45 := GenS1{F0: 128, F1: 834, F2: 1}
                                    _ = fi43
                                    _ = v44
                                    _ = v45
                                }
                                for fi46 := 0; fi46 < 19; fi46++ {
                                    v3 -= v3
                                    var v47 uint = 18446744073709551615
                                    var v48 float32 = -2.25
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                                for fi49 := 0; fi49 < 7; fi49++ {
                                    v50 := GenS4{F0: -1, F1: 1, F2: (float32(-2.25) == -2.25)}
                                    _ = fi49
                                    _ = v50
                                }
                            }
                            var v51 float32 = (float32(1e10) / 3.14159)
                            _ = fi40
                            _ = v51
                        }
                        fmt.Println(v3)
                    } else {
                        for fi52 := 0; fi52 < 10; fi52++ {
                            fmt.Println(v3)
                            _ = fi52
                        }
                    }
                case "a":
                    for fi53 := 0; fi53 < 17; fi53++ {
                        switch -144 {
                        case 784:
                            var v54 int32 = 639
                            _ = v54
                        default:
                            v3 += fi53
                        }
                        switch "hello" {
                        case "":
                            for fi55 := 0; fi55 < 13; fi55++ {
                                switch 9223372036854775807 {
                                case 9223372036854775807:
                                    var v56 uint8 = 1
                                    _ = v56
                                }
                                v3 += (fi53 + v3)
                                if false {
                                    v3 = fi53
                                    var v57 int8 = 42
                                    _ = v57
                                }
                                _ = fi55
                            }
                        default:
                            if (!(42 <= 0)) {
                                if false {
                                    var v58 uint32 = 2
                                    _ = v58
                                }
                                v2[65535] = (v3 % (((fi53 % ((0) | 1))) | 1))
                                if false {
                                    v3 += fi53
                                    if len(v4) > 0 {
                                        v4[0] = 352
                                    }
                                    var v59 float64 = 0.1
                                    _ = v59
                                }
                            }
                        }
                        var v60 int32 = 2147483647
                        _ = fi53
                        _ = v60
                    }
                default:
                    v3 += v3
                }
            default:
                switch "a" {
                case "a":
                    v3 -= v3
                case "":
                    v3 = -593
                case "\n\t":
                    v2[2147483648] = (v3 * -91)
                default:
                    v61 := GenS3{F0: -32768, F1: v3}
                    _ = v61
                }
            }
            if true {
                switch v3 {
                case -459:
                    if (-1 < 0) {
                        for fi62 := 0; fi62 < 6; fi62++ {
                            var v63 string = ""
                            _ = fi62
                            _ = v63
                        }
                        v64 := GenS5{F0: "a", F1: 42}
                        fmt.Println(v3)
                        _ = v64
                    }
                case 32768:
                    v4 = append(v4, 1)
                }
            }
        }
    }
    for fi65 := 0; fi65 < 10; fi65++ {
        if (!(!true)) {
            var v66 int64 = 42
            _ = v66
        }
        if false {
            var v67 int = -1
            _ = v67
        } else {
            v68 := GenS6{F0: v3}
            v3 += (v3 | (v3 | (fi65 - (v3 - (v3 & 589)))))
            _ = v68
        }
        if (3.14159 > float64(1.5)) {
            v3 += v3
            for fi69 := 0; fi69 < 8; fi69++ {
                if false {
                    v4 = append(v4, 0)
                } else {
                    if (true || false) {
                        v3 += 128
                        v70 := GenS7{F0: -32768, F1: 128}
                        v71 := GenS8{F0: 42}
                        _ = v70
                        _ = v71
                    }
                    v3 -= (fi69 ^ v3)
                    var v72 int64 = 1
                    _ = v72
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi69
            }
        } else {
            switch ("a" + (("" + "") + ("foo bar" + ("a" + ("\n\t" + ((((("\n\t" + "") + "\n\t") + "hello") + (((("a" + ("a" + "a")) + ((("\n\t" + ((("foo bar" + "hello") + "\n\t") + ("a" + "\n\t"))) + ("\n\t" + "foo bar")) + "")) + "a") + ("a" + "hello"))) + "")))))) {
            case "":
                v3 -= -1
            case "\n\t":
                if ((!(!(127 > 40))) || (("" + ("" + ("" + "\n\t"))) != "\n\t")) {
                    var v73 uint64 = 625
                    var v74 int64 = -1
                    _ = v73
                    _ = v74
                }
            case "foo bar":
                if (-11 != -74) {
                    fmt.Println(fi65, fi65)
                    var v75 int = fi65
                    if (42 != 467) {
                        switch (v75 + v75) {
                        case 1:
                            for fi76 := 0; fi76 < 5; fi76++ {
                                var v77 uint8 = 219
                                if true {
                                    v4 = append(v4, 2)
                                    v77 -= v77
                                    if false {
                                        v78 := GenS8{F0: 0}
                                        var v79 int8 = -85
                                        _ = v78
                                        _ = v79
                                    }
                                } else {
                                    v4 = append(v4, 4294967295)
                                    var v80 int16 = 0
                                    _ = v80
                                }
                                v3 -= 9223372036854775807
                                v81 := GenS8{F0: 0}
                                v82 := GenS5{F0: ((("hello" + "\n\t") + ("\n\t" + "a")) + "hello"), F1: v77}
                                _ = fi76
                                _ = v77
                                _ = v81
                                _ = v82
                            }
                        }
                        v3 += (fi65 & -9223372036854775808)
                    }
                    _ = v75
                }
            default:
                var v83 uint64 = 272
                _ = v83
            }
            v84 := GenS9{F0: 127}
            _ = v84
        }
        _ = fi65
    }
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
