package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 bool; F2 int }
type GenS3 struct { F0 uint16; F1 uint; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int32; F2 int64 }
type GenS6 struct { F0 bool; F1 int8; F2 uint16 }

func main() {
    v1 := []uint32{0}
    var v2 int32 = 0
    var v3 int = 585
    var v4 int8 = 127
    v5 := []uint32{1, 921}
    var v6 uint32 = 2147483648
    v5 = append(v5, 4294967295)
    v6 = 901
    v7 := GenS1{F0: v3}
    v8 := GenS2{F0: int8(v6), F1: false, F2: v3}
    var v9 int16 = 32767
    if len(v1) > 0 {
        v1[0] = v6
    }
    switch "hello" {
    case "":
        var v10 int8 = v4
        _ = v10
    }
    var v11 float32 = ((0.0 / (((float32(((0.0 * 3.14159) - 1.5)) - 0.0) / 1e10) + (0.0 - 0.0))) - (3.14159 / -2.25))
    for fi12 := 0; fi12 < 12; fi12++ {
        v3 += fi12
        var v13 int64 = 9223372036854775807
        if (0 <= 32767) {
            var v14 uint8 = 1
            v3 += -2147483648
            v11 = 1.5
            _ = v14
        } else {
            var v15 uint8 = 0
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    if ("" <= "") {
        v3 = (v3 << ((18446744073709551615) & 63))
        if (((1.5 - (1e10 / (((0.0 / (3.14159 - 1e10)) + (((3.14159 * 1e10) * float64(v11)) / (0.0 - (3.14159 + (1e10 + -2.25))))) - (0.0 + 0.0)))) * 1e10) != (((0.0 - -2.25) - 0.0) / 0.1)) {
            switch (v3 + v3) {
            case 9223372036854775807:
                switch ("hello" + ("\n\t" + ("a" + "foo bar"))) {
                case "a":
                    var v16 int = 32768
                    _ = v16
                }
            case 0:
                if true {
                    v2 = 42
                    var v17 int32 = (v2 & v2)
                    _ = v17
                } else {
                    if (int(v2) >= (-(v3))) {
                        for fi18 := 0; fi18 < 8; fi18++ {
                            v4 = (v4 % (((v4 * v4)) | 1))
                            _ = fi18
                        }
                        v2 += (^(v2))
                    } else {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            switch "hello" {
                            case "hello":
                                fmt.Println(v9)
                            case "":
                                for fi20 := 0; fi20 < 12; fi20++ {
                                    var v21 int64 = -962
                                    _ = fi20
                                    _ = v21
                                }
                            case "a":
                                var v22 int16 = (v9 - 587)
                                _ = v22
                            default:
                                v3 -= int(v9)
                            }
                            _ = fi19
                        }
                        switch "" {
                        case "":
                            for fi23 := 0; fi23 < 8; fi23++ {
                                if true {
                                    var v24 uint64 = 98
                                    var v25 uint16 = 1
                                    _ = v24
                                    _ = v25
                                }
                                var v26 float32 = (1.5 / 1.0)
                                _ = fi23
                                _ = v26
                            }
                        }
                    }
                    var v27 float64 = (1e10 - 0.0)
                    _ = v27
                }
            default:
                var v28 int64 = 42
                _ = v28
            }
            fmt.Println(v9)
            for fi29 := 0; fi29 < 11; fi29++ {
                v30 := GenS3{F0: 2, F1: 1, F2: 65535}
                v2 -= (-(v2))
                _ = fi29
                _ = v30
            }
        }
    } else {
        v11 = (v11 / v11)
    }
    v31 := GenS1{F0: -9223372036854775808}
    var v32 int32 = v2
    v6 = v6
    for fi33 := 0; fi33 < 18; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            switch "foo bar" {
            case "foo bar":
                switch (("a" + "foo bar") + (("hello" + ("hello" + "a")) + "")) {
                case "\n\t":
                    var v35 string = "foo bar"
                    _ = v35
                default:
                    if (79 > (v32 & v2)) {
                        var v36 int16 = (v9 & v9)
                        _ = v36
                    }
                }
            case "\n\t":
                var v37 int16 = 42
                _ = v37
            default:
                var v38 string = "a"
                _ = v38
            }
            _ = fi34
        }
        for fi39 := 0; fi39 < 11; fi39++ {
            v32 += (^(v2))
            if (0.0 >= 1e10) {
                var v40 float64 = -2.25
                switch "hello" {
                case "":
                    var v41 int64 = -1
                    _ = v41
                case "\n\t":
                    if false {
                        fmt.Println(v3, fi39)
                        var v42 int16 = (^(v9))
                        if true {
                            v43 := GenS2{F0: v4, F1: false, F2: fi33}
                            _ = v43
                        }
                        _ = v42
                    } else {
                        v11 += v11
                    }
                case "foo bar":
                    v40 = (((v40 * (3.14159 + v40)) + (0.1 + 0.0)) - v40)
                default:
                    v44 := GenS1{F0: -843}
                    _ = v44
                }
                _ = v40
            }
            var v45 uint16 = 255
            var v46 uint8 = 127
            var v47 uint32 = 140
            v47 += 2147483647
            _ = fi39
            _ = v45
            _ = v46
            _ = v47
        }
        if ((("" + "foo bar") + "") <= "") {
            if ("\n\t" == "\n\t") {
                switch "" {
                case "hello":
                    fmt.Println(v6, v32, v3)
                case "a":
                    v5 = append(v5, 4294967295)
                case "foo bar":
                    v4 -= v4
                }
                v5 = append(v5, (v6 ^ 2147483647))
                if (42 <= 65535) {
                    v48 := GenS4{F0: 65535}
                    v3 += (fi33 / ((fi33) | 1))
                    _ = v48
                } else {
                    v4 = v4
                    v11 = v11
                    if ((true || false) && (!true)) {
                        if (!true) {
                            var v49 int32 = v2
                            var v50 uint16 = 32767
                            _ = v49
                            _ = v50
                        } else {
                            for fi51 := 0; fi51 < 17; fi51++ {
                                if true {
                                    var v52 uint32 = 40
                                    _ = v52
                                } else {
                                    var v53 int = v3
                                    _ = v53
                                }
                                _ = fi51
                            }
                            fmt.Println(v2, v9)
                        }
                    }
                }
            } else {
                v11 = v11
                if true {
                    switch ("a" + "\n\t") {
                    case "hello":
                        v3 -= (^(v3))
                    default:
                        switch 902 {
                        case 170:
                            switch 222 {
                            case -2147483648:
                                var v54 float64 = 0.1
                                _ = v54
                            case 2147483647:
                                fmt.Println(v2, v9, v9)
                            default:
                                if (!((0.0 / (float64(-2.25) * 0.1)) < -2.25)) {
                                    var v55 string = "foo bar"
                                    _ = v55
                                } else {
                                    var v56 uint16 = 0
                                    if true {
                                        fmt.Println(v32, v2)
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    var v57 int16 = (^(v9))
                                    _ = v56
                                    _ = v57
                                }
                            }
                        case 42:
                            if false {
                                var v58 int16 = v9
                                _ = v58
                            } else {
                                v32 += v2
                                var v59 uint16 = 0
                                _ = v59
                            }
                        case 32767:
                            fmt.Println(fi33, v9)
                        }
                    }
                    if (true || (2 < 0)) {
                        for fi60 := 0; fi60 < 18; fi60++ {
                            _ = len(v5)
                            v6 += 383
                            _ = fi60
                        }
                        var v61 uint32 = 2
                        switch "" {
                        case "":
                            fmt.Println(v4)
                        }
                        _ = v61
                    } else {
                        if false {
                            v6 = (v6 + v6)
                            _ = len(v1)
                        }
                        var v62 int32 = (v32 << ((717) & 63))
                        v1 = append(v1, (v6 | v6))
                        _ = v62
                    }
                }
            }
            for fi63 := 0; fi63 < 4; fi63++ {
                _ = len(v5)
                if false {
                    v6 += v6
                } else {
                    v64 := GenS5{F0: 0, F1: (^(v32)), F2: -1}
                    v65 := GenS3{F0: 2, F1: 65535, F2: 42}
                    var v66 int16 = (v9 + (v9 >> ((4294967295) & 63)))
                    _ = v64
                    _ = v65
                    _ = v66
                }
                v9 = (^(v9))
                _ = fi63
            }
        } else {
            for fi67 := 0; fi67 < 4; fi67++ {
                switch v3 {
                case 0:
                    v68 := GenS6{F0: (92 <= v4), F1: (v4 - (v4 >> ((2) & 63))), F2: 255}
                    _ = v68
                case -908:
                    var v69 int = 0
                    _ = v69
                }
                var v70 uint32 = (^(v6))
                var v71 float32 = float32(1.5)
                _ = fi67
                _ = v70
                _ = v71
            }
            var v72 int16 = (v9 & (^(v9)))
            _ = v72
        }
        _ = fi33
    }
    for fi73 := 0; fi73 < 12; fi73++ {
        for fi74 := 0; fi74 < 6; fi74++ {
            v11 = v11
            _ = fi74
        }
        v11 = (1.5 + v11)
        var v75 bool = true
        _ = fi73
        _ = v75
    }
    fmt.Println(v2, v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
