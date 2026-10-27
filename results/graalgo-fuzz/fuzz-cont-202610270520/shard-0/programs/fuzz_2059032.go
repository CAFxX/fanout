package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int; F1 int32; F2 uint64 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int }

func main() {
    var v1 int32 = 2147483647
    var v2 int8 = 42
    var v3 uint = 255
    v4 := []uint32{434, 297}
    var v5 int32 = (^(v1))
    v6 := GenS1{F0: ((("a" + ("a" + ("hello" + (("\n\t" + "hello") + "\n\t")))) + "\n\t") != "foo bar"), F1: v5}
    v1 -= (v1 ^ (v1 - v5))
    v5 += v5
    var v7 uint64 = 0
    if true {
        v2 = (v2 + v2)
        if false {
            v2 += v2
            for fi8 := 0; fi8 < 19; fi8++ {
                var v9 int = 4294967295
                _ = fi8
                _ = v9
            }
            if (128 >= 127) {
                var v10 string = ""
                switch v10 {
                case "\n\t":
                    var v11 float64 = float64(float32(-2.25))
                    _ = v11
                case "":
                    fmt.Println(v7, v7, v10)
                }
                v12 := GenS1{F0: ((^(v1)) <= -2147483648), F1: v1}
                _ = v10
                _ = v12
            }
        }
        _ = len(v4)
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            v4 = append(v4, 42)
            _ = fi13
        }
    }
    if (42 < 42) {
        for fi14 := 0; fi14 < 17; fi14++ {
            if len(v4) > 0 {
                v4[0] = 0
            }
            _ = fi14
        }
        var v15 int32 = (-(v5))
        v4 = append(v4, 273)
        _ = v15
    } else {
        v2 += v2
    }
    v4 = append(v4, 65535)
    for fi16 := 0; fi16 < 4; fi16++ {
        v5 -= (v1 + v1)
        if true {
            var v17 uint64 = v7
            _ = v17
        } else {
            _ = len(v4)
            for fi18 := 0; fi18 < 2; fi18++ {
                if ((v3 + 42) < (v3 >> (((^(v3))) & 63))) {
                    _ = len(v4)
                }
                _ = fi18
            }
        }
        v2 = (v2 + v2)
        _ = fi16
    }
    switch -478 {
    case 9223372036854775807:
        v19 := GenS2{F0: v3, F1: 127}
        _ = v19
    case 547:
        for fi20 := 0; fi20 < 1; fi20++ {
            v4 = append(v4, 42)
            _ = fi20
        }
    case 42:
        fmt.Println(v3, v5)
    default:
        var v21 uint = (v3 / ((1) | 1))
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        if (1.5 < 0.0) {
            switch (fi22 ^ fi22) {
            case 9223372036854775807:
                fmt.Println(v2, v7)
            }
            var v23 uint64 = 338
            _ = v23
        } else {
            fmt.Println(v5, v7)
        }
        if false {
            v1 = v5
        }
        for fi24 := 0; fi24 < 4; fi24++ {
            if len(v4) > 0 {
                v4[0] = uint32(v3)
            }
            _ = fi24
        }
        _ = fi22
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        if (false && ((v7 >> ((v3) & 63)) != v7)) {
            switch "a" {
            case "foo bar":
                for fi26 := 0; fi26 < 20; fi26++ {
                    var v27 uint8 = 2
                    if (!false) {
                        switch (^(fi25)) {
                        case 42:
                            v4 = append(v4, 427)
                        case 128:
                            v7 = 1
                            v27 = (v27 - v27)
                            v28 := GenS2{F0: (v3 / ((v3) | 1)), F1: (v27 >> (((v3 | (^(v3)))) & 63))}
                            _ = v28
                        }
                    } else {
                        if true {
                            if (false || false) {
                                _ = len(v4)
                                if true {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    v27 += v27
                                } else {
                                    if false {
                                        if false {
                                            if false {
                                                v29 := GenS3{F0: v27}
                                                if ((1 != 42) || (false || (!true))) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                _ = v29
                                            }
                                            v7 = 235
                                            v30 := GenS3{F0: v27}
                                            var v31 int = -591
                                            _ = v30
                                            _ = v31
                                        }
                                        switch "foo bar" {
                                        case "":
                                            var v32 uint64 = (v7 ^ (v7 * v7))
                                            _ = v32
                                        case "\n\t":
                                            var v33 uint32 = 23
                                            _ = v33
                                        case "foo bar":
                                            var v34 float64 = float64((3.14159 / float32(((-2.25 / (0.0 + (float64((float32(0.0) - 0.1)) + -2.25))) - 1e10))))
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v35 := GenS4{F0: -1}
                                            _ = v34
                                            _ = v35
                                        default:
                                            v3 = uint(v27)
                                        }
                                        fmt.Println(v2, v3)
                                    } else {
                                        v5 -= (v5 - v5)
                                        v5 = (v5 & (v5 - v1))
                                    }
                                }
                                if (false || true) {
                                    var v36 float32 = float32(-2.25)
                                    if ("" == "") {
                                        v37 := GenS5{F0: (fi26 ^ 933), F1: 45, F2: uint64(v3)}
                                        v38 := GenS6{F0: v1, F1: -908}
                                        _ = v37
                                        _ = v38
                                    } else {
                                        var v39 float64 = (0.1 * -2.25)
                                        if false {
                                            v2 -= v2
                                        }
                                        _ = v39
                                    }
                                    _ = v36
                                }
                                v40 := GenS6{F0: v1, F1: -1}
                                v41 := GenS7{F0: -32768}
                                _ = v40
                                _ = v41
                            }
                            if (-1 < 42) {
                                if true {
                                    var v42 uint8 = v27
                                    var v43 bool = (false || false)
                                    if len(v4) > 0 {
                                        v4[0] = 65535
                                    }
                                    _ = v42
                                    _ = v43
                                } else {
                                    fmt.Println(v7, v27, v5)
                                    v44 := GenS3{F0: (v27 << (((^(v3))) & 63))}
                                    if false {
                                        var v45 float32 = 3.14159
                                        fmt.Println(v45, v5, v7)
                                        var v46 uint64 = 1
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = v44
                                }
                            } else {
                                var v47 uint64 = v7
                                _ = v47
                            }
                        } else {
                            v48 := GenS8{F0: (fi25 / (((^(fi25))) | 1))}
                            var v49 int16 = int16(fi26)
                            fmt.Println(fi25, v1, v3)
                            var v50 float64 = (3.14159 / ((-2.25 / 1.0) + -2.25))
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                    }
                    _ = fi26
                    _ = v27
                }
            }
            var v51 uint = (^(v3))
            _ = v51
        }
        if true {
            switch "a" {
            case "\n\t":
                fmt.Println(v2, v1)
            case "":
                var v52 int8 = v2
                _ = v52
            }
            switch (("foo bar" + "hello") + ("" + ("\n\t" + ("" + "a")))) {
            case "foo bar":
                fmt.Println(v1, v1)
            default:
                v53 := GenS7{F0: 609}
                _ = v53
            }
            if len(v4) > 0 {
                v4[0] = 2
            }
        }
        _ = fi25
    }
    var v54 bool = (false && ("foo bar" >= "hello"))
    var v55 int8 = (v2 * v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
    _ = v7
}
