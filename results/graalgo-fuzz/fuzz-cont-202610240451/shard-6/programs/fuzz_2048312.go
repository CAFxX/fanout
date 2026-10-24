package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int16; F1 int16; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 bool; F1 int8; F2 uint16 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64; F1 bool; F2 uint32 }

func main() {
    v1 := []uint32{2147483648, 4294967295, 0}
    v2 := []int{844}
    var v3 uint8 = 0
    v4 := make(map[string]string)
    if (((v3 / (((v3 & 2)) | 1)) == (v3 & v3)) && true) {
        var v5 string = "a"
        _ = v5
    } else {
        v6 := GenS1{F0: true, F1: 0}
        for fi7 := 0; fi7 < 1; fi7++ {
            v3 = (v3 >> ((42) & 63))
            v3 += v3
            _ = fi7
        }
        var v8 uint32 = 4294967295
        _ = v6
        _ = v8
    }
    _ = len(v4)
    for fi9 := 0; fi9 < 10; fi9++ {
        v3 -= 1
        _ = fi9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        v3 = 128
        if (2 == 2) {
            var v11 int16 = 217
            _ = v11
        } else {
            if (!(((^(v3)) != (v3 + v3)) && (!((42 <= -1) || true)))) {
                if ((-2.25 > 1.5) || (-1 >= 0)) {
                    var v12 float64 = 3.14159
                    _ = v12
                }
            }
            v3 = 1
            if (58 == 377) {
                if (!false) {
                    v3 -= v3
                } else {
                    if true {
                        var v13 uint64 = 18446744073709551615
                        v14 := GenS2{F0: int64(fi10), F1: -64}
                        if (("foo bar" + ("" + ((("" + "hello") + "hello") + "\n\t"))) == (("a" + "foo bar") + "")) {
                            _ = len(v4)
                            var v15 int = fi10
                            var v16 uint64 = v13
                            _ = v15
                            _ = v16
                        } else {
                            v17 := GenS2{F0: -1, F1: 112}
                            v3 -= (v3 * 128)
                            _ = v17
                        }
                        _ = v13
                        _ = v14
                    } else {
                        fmt.Println(fi10, v3)
                        for fi18 := 0; fi18 < 9; fi18++ {
                            if true {
                                var v19 bool = (2147483647 == -758)
                                fmt.Println(v3)
                                _ = v19
                            } else {
                                var v20 int8 = -124
                                _ = v20
                            }
                            _ = fi18
                        }
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch "hello" {
                    case "\n\t":
                        v21 := GenS1{F0: (0.0 != 0.0), F1: 2147483647}
                        _ = v21
                    case "":
                        _ = len(v2)
                    case "a":
                        delete(v4, (("hello" + "a") + "a"))
                    default:
                        v22 := GenS3{F0: -1, F1: 1, F2: 344}
                        _ = v22
                    }
                }
                var v23 float32 = (0.1 / 1.5)
                _ = v23
            } else {
                v3 += v3
                switch (fi10 - fi10) {
                case -121:
                    switch "hello" {
                    case "":
                        v3 = v3
                    default:
                        v24 := v4["\n\t"]
                        _ = v24
                    }
                default:
                    for fi25 := 0; fi25 < 5; fi25++ {
                        var v26 int32 = 1
                        _ = fi25
                        _ = v26
                    }
                }
            }
        }
        _ = fi10
    }
    switch "foo bar" {
    case "":
        v27 := GenS4{F0: 42}
        _ = v27
    case "\n\t":
        var v28 int64 = 9223372036854775807
        _ = v28
    default:
        if (-1 > 32768) {
            if (2147483647 == 127) {
                if false {
                    if (-2.25 > (-2.25 / ((0.0 - float64(-2.25)) - 1.5))) {
                        fmt.Println(v3)
                    } else {
                        if (true && (2 >= 207)) {
                            if len(v2) > 0 {
                                v2[0] = -978
                            }
                            if false {
                                switch -9223372036854775808 {
                                case -9223372036854775808:
                                    v29 := GenS2{F0: 9223372036854775807, F1: 1}
                                    _ = v29
                                default:
                                    fmt.Println(v3)
                                }
                                v3 = 136
                                v4[""] = "hello"
                            } else {
                                var v30 int64 = 1
                                fmt.Println(v3)
                                if (!false) {
                                    if true {
                                        for fi31 := 0; fi31 < 14; fi31++ {
                                            if true {
                                                if (false || true) {
                                                    v32 := GenS5{F0: 1, F1: (^(v3)), F2: -384}
                                                    for fi33 := 0; fi33 < 15; fi33++ {
                                                        v34 := GenS4{F0: v30}
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v30 += -901
                                                        _ = fi33
                                                        _ = v34
                                                    }
                                                    _ = v32
                                                } else {
                                                    v4[(((("" + "") + "hello") + ("foo bar" + "")) + "")] = "a"
                                                    v30 += 0
                                                    if true {
                                                        if false {
                                                            var v35 string = ("a" + "hello")
                                                            var v36 int = (^(fi31))
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                    }
                                                }
                                                switch (fi31 / ((fi31) | 1)) {
                                                case 713:
                                                    v3 = v3
                                                case -908:
                                                    _ = len(v1)
                                                case 275:
                                                    var v37 int32 = 42
                                                    _ = v37
                                                }
                                                var v38 uint8 = 71
                                                _ = v38
                                            } else {
                                                if (!(-2147483648 == 593)) {
                                                    var v39 int32 = 2147483647
                                                    v40 := GenS6{F0: 1, F1: v3, F2: (v3 * 128)}
                                                    switch 2147483647 {
                                                    case 9223372036854775807:
                                                        v3 = (v3 * (v3 | v3))
                                                    default:
                                                        var v41 int64 = (v30 % ((-780) | 1))
                                                        _ = v41
                                                    }
                                                    _ = v39
                                                    _ = v40
                                                }
                                                var v42 uint16 = 1
                                                _ = v42
                                            }
                                            v43 := GenS7{F0: (v3 & v3), F1: ""}
                                            var v44 int = fi31
                                            _ = fi31
                                            _ = v43
                                            _ = v44
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = 0
                                        }
                                        _ = len(v2)
                                    } else {
                                        v45 := GenS8{F0: false, F1: 127, F2: 2}
                                        _ = v45
                                    }
                                    fmt.Println(v30, v3)
                                    v30 += (v30 * v30)
                                } else {
                                    if (true || false) {
                                        var v46 int64 = (v30 ^ v30)
                                        for fi47 := 0; fi47 < 4; fi47++ {
                                            fmt.Println(v3, v30)
                                            v30 -= -10
                                            _ = fi47
                                        }
                                        _ = v46
                                    }
                                }
                                _ = v30
                            }
                        }
                        var v48 uint8 = 128
                        _ = v48
                    }
                }
                var v49 int16 = 1000
                for fi50 := 0; fi50 < 11; fi50++ {
                    for fi51 := 0; fi51 < 8; fi51++ {
                        switch "hello" {
                        case "\n\t":
                            v3 = v3
                            var v52 float32 = -2.25
                            _ = v52
                        }
                        var v53 uint16 = 255
                        var v54 float32 = 0.1
                        _ = fi51
                        _ = v53
                        _ = v54
                    }
                    v55 := GenS1{F0: false, F1: -860}
                    _ = fi50
                    _ = v55
                }
                _ = v49
            }
            v56 := GenS7{F0: v3, F1: "a"}
            _ = v56
        } else {
            v57 := GenS4{F0: 42}
            var v58 float64 = float64((0.0 * (0.1 - 0.1)))
            _ = len(v2)
            _ = v57
            _ = v58
        }
    }
    var v59 uint64 = 0
    v60 := GenS8{F0: ((v3 ^ (v3 | 11)) != v3), F1: 0, F2: 302}
    for fi61 := 0; fi61 < 5; fi61++ {
        v59 += (v59 * (v59 >> ((255) & 63)))
        _ = fi61
    }
    switch "" {
    case "a":
        switch ("a" + "") {
        case "hello":
            if true {
                for fi62 := 0; fi62 < 20; fi62++ {
                    v4[((("hello" + "hello") + "hello") + "")] = ("foo bar" + "foo bar")
                    _ = fi62
                }
                v63 := GenS2{F0: -372, F1: 42}
                if (!(v59 == v59)) {
                    v3 = (v3 - 128)
                    if true {
                        switch 2147483647 {
                        case 32768:
                            fmt.Println(v59, v3)
                        case -1:
                            _ = len(v2)
                        case 0:
                            var v64 uint32 = 2
                            _ = v64
                        }
                        switch 809 {
                        case 1:
                            var v65 uint16 = 2
                            _ = v65
                        case -2147483648:
                            var v66 uint8 = v3
                            _ = v66
                        case 591:
                            for fi67 := 0; fi67 < 11; fi67++ {
                                var v68 float64 = ((0.0 - 1.5) / 1.5)
                                delete(v4, "a")
                                v69 := GenS1{F0: (float64(((float32((v68 - v68)) + float32(1.5)) + (((0.0 - float32(v68)) / 0.1) + (0.0 - (0.1 + 1e10))))) <= v68), F1: 42}
                                _ = fi67
                                _ = v68
                                _ = v69
                            }
                        }
                    }
                }
                _ = v63
            }
        case "a":
            v59 += v59
        case "\n\t":
            var v70 uint = 4294967295
            _ = v70
        }
    case "foo bar":
        for fi71 := 0; fi71 < 13; fi71++ {
            var v72 uint16 = 42
            v73 := GenS9{F0: (v72 >> ((423) & 63))}
            _ = fi71
            _ = v72
            _ = v73
        }
    case "\n\t":
        fmt.Println(v3)
    }
    switch 9223372036854775807 {
    case 127:
        fmt.Println(v59, v3)
    case 42:
        for fi74 := 0; fi74 < 16; fi74++ {
            var v75 int64 = -1
            for fi76 := 0; fi76 < 7; fi76++ {
                var v77 int32 = 42
                v59 -= v59
                var v78 int32 = 798
                _ = fi76
                _ = v77
                _ = v78
            }
            var v79 uint64 = 0
            _ = fi74
            _ = v75
            _ = v79
        }
    default:
        for fi80 := 0; fi80 < 16; fi80++ {
            v3 = (^(v3))
            switch (^(fi80)) {
            case 434:
                _ = len(v4)
            case 1:
                v81 := GenS10{F0: -9223372036854775808, F1: ("a" > "hello"), F2: 238}
                _ = v81
            default:
                var v82 uint8 = (v3 >> ((2) & 63))
                _ = v82
            }
            _ = fi80
        }
    }
    fmt.Println(v3, v3)
    v59 = (v59 >> ((65535) & 63))
    v3 = uint8(v59)
    if (1.5 != float64(0.1)) {
        _ = len(v2)
        var v83 int = 4294967295
        _ = v83
    }
    if len(v1) > 0 {
        v1[0] = 170
    }
    for fi84 := 0; fi84 < 16; fi84++ {
        v85 := v4["\n\t"]
        fmt.Println(fi84)
        _ = fi84
        _ = v85
    }
    if (!(908 <= 409)) {
        var v86 bool = (false && false)
        fmt.Println(v86, v86)
        if (!(18446744073709551615 > 255)) {
            fmt.Println(v86, v3)
            var v87 int8 = 1
            _ = v87
        } else {
            v3 -= (^(v3))
            _ = len(v2)
        }
        _ = v86
    } else {
        for fi88 := 0; fi88 < 3; fi88++ {
            switch "a" {
            case "a":
                var v89 string = "a"
                _ = v89
            case "foo bar":
                var v90 int32 = 954
                _ = v90
            case "\n\t":
                for fi91 := 0; fi91 < 16; fi91++ {
                    v3 -= v3
                    switch "foo bar" {
                    case "foo bar":
                        v92 := v4["\n\t"]
                        _ = v92
                    }
                    _ = fi91
                }
            }
            _ = fi88
        }
        fmt.Println(v3, v59)
        v59 = (v59 & 807)
    }
    v59 -= v59
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
}
