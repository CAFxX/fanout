package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint; F1 int8; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint32 }
type GenS6 struct { F0 bool }

func main() {
    var v1 uint64 = 517
    v2 := []uint32{258, 220}
    v3 := make(map[uint32]bool)
    var v4 float64 = (-2.25 + -2.25)
    var v5 uint64 = 54
    var v6 bool = (630 != 443)
    for fi7 := 0; fi7 < 18; fi7++ {
        v2 = append(v2, 1)
        v3[65535] = (v6 || (v6 && v6))
        v4 -= ((v4 * v4) - v4)
        _ = fi7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        if v6 {
            v9 := GenS1{F0: -128, F1: 1}
            fmt.Println(v1)
            var v10 uint = 18446744073709551615
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    switch "hello" {
    case "\n\t":
        for fi11 := 0; fi11 < 8; fi11++ {
            switch -321 {
            case 4294967295:
                v3[2] = (false || v6)
            case 0:
                switch ((("foo bar" + "foo bar") + ("hello" + "a")) + ("" + "\n\t")) {
                case "\n\t":
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                case "foo bar":
                    v1 = (v1 * (v5 / (((v5 % (((^(v1))) | 1))) | 1)))
                }
            case -585:
                fmt.Println(v1)
            }
            var v12 int = fi11
            v13 := v3[31]
            _ = fi11
            _ = v12
            _ = v13
        }
    case "":
        v14 := GenS1{F0: 1, F1: 9223372036854775807}
        _ = v14
    default:
        if (0 > -299) {
            var v15 int = -2147483648
            if v6 {
                _ = len(v3)
                var v16 uint32 = 969
                _ = v16
            }
            _ = v15
        } else {
            v17 := GenS1{F0: -128, F1: -1}
            v18 := GenS1{F0: -128, F1: 9223372036854775807}
            _ = v17
            _ = v18
        }
    }
    var v19 string = ("" + "hello")
    var v20 int = 0
    v21 := GenS1{F0: 0, F1: 1}
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 10; fi23++ {
            v1 = 0
            v6 = v6
            _ = fi23
        }
        for fi24 := 0; fi24 < 6; fi24++ {
            v25 := GenS2{F0: 137}
            if v6 {
                v1 = 803
                if (86 == -84) {
                    var v26 string = v19
                    var v27 int = (fi22 + (-(v20)))
                    _ = v26
                    _ = v27
                }
                if len(v2) > 0 {
                    v2[0] = 4294967295
                }
            } else {
                v4 = v4
                if (v6 || (2 == 2)) {
                    switch (fi24 + fi24) {
                    case 4294967295:
                        fmt.Println(fi24, v20, fi24)
                    default:
                        if (0 >= -32768) {
                            _ = len(v2)
                            switch (v19 + v19) {
                            case "\n\t":
                                v2 = append(v2, 278)
                            default:
                                v19 = ("a" + "hello")
                                v28 := GenS1{F0: -128, F1: 530}
                                if (-32 == -128) {
                                    v5 = 2
                                    v6 = v6
                                } else {
                                    var v29 uint32 = 997
                                    var v30 bool = true
                                    _ = v29
                                    _ = v30
                                }
                                _ = v28
                            }
                        }
                    }
                    v5 += 18446744073709551615
                } else {
                    var v31 int64 = -199
                    v32 := GenS1{F0: 127, F1: (v31 / (((v31 - (v31 + (v31 << ((2) & 63))))) | 1))}
                    v33 := GenS2{F0: 187}
                    _ = v31
                    _ = v32
                    _ = v33
                }
                switch (v20 + (^(fi22))) {
                case 0:
                    v19 = (v19 + ((v19 + v19) + v19))
                case 1:
                    if (-1 <= -9223372036854775808) {
                        if ((!(0 == -2147483648)) && (v6 || false)) {
                            switch (v19 + (v19 + (v19 + v19))) {
                            case "\n\t":
                                fmt.Println(fi24, v1)
                            default:
                                v34 := GenS2{F0: 42}
                                _ = v34
                            }
                            fmt.Println(v20)
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        v19 = (v19 + v19)
                    } else {
                        var v35 uint64 = v1
                        _ = v35
                    }
                default:
                    v5 = v5
                    if false {
                        v4 -= v4
                    }
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                }
                var v36 float32 = 0.0
                v37 := GenS3{F0: (v19 + ((("a" + ("" + "")) + v19) + (v19 + ("foo bar" + (((v19 + v19) + v19) + (v19 + "hello")))))), F1: 42, F2: 1}
                _ = v36
                _ = v37
            }
            if (0.1 != (float64(0.0) / 1.0)) {
                switch v20 {
                case -9223372036854775808:
                    var v38 float32 = -2.25
                    _ = v38
                case -312:
                    v5 -= (^(v1))
                default:
                    v20 -= (v20 / (((fi22 / ((-9223372036854775808) | 1))) | 1))
                }
                if ((!(42 < -711)) || true) {
                    switch fi22 {
                    case 4294967295:
                        var v39 float32 = 3.14159
                        _ = v39
                    case 2147483647:
                        v1 = (v1 >> ((655) & 63))
                    case -9223372036854775808:
                        var v40 int8 = int8(fi22)
                        _ = v40
                    }
                    switch (v19 + v19) {
                    case "a":
                        v5 = v5
                    case "":
                        switch v19 {
                        case "a":
                            v1 = (v1 ^ (v5 | v1))
                        case "hello":
                            var v41 float32 = float32(v4)
                            _ = v41
                        case "foo bar":
                            if v6 {
                                if (!(v6 || v6)) {
                                    switch (((((v19 + v19) + "\n\t") + "a") + v19) + "\n\t") {
                                    case "a":
                                        v4 = ((float64(float32(1.5)) / 3.14159) - ((1.5 - (v4 + (v4 - v4))) + v4))
                                    case "foo bar":
                                        fmt.Println(v6, v19, fi24)
                                    default:
                                        if (1 <= 994) {
                                            v42 := GenS3{F0: ((v19 + "") + v19), F1: 42, F2: 0}
                                            _ = v42
                                        } else {
                                            var v43 string = (("\n\t" + v19) + v19)
                                            var v44 float32 = (((1e10 - 1e10) / (-2.25 / 0.1)) * 3.14159)
                                            v43 = "foo bar"
                                            v44 = 0.0
                                            _ = v43
                                            _ = v44
                                        }
                                    }
                                    v45 := GenS2{F0: 1}
                                    var v46 int = (fi22 * 32768)
                                    _ = v45
                                    _ = v46
                                }
                                var v47 uint8 = 230
                                _ = v47
                            }
                        default:
                            var v48 uint = 255
                            _ = v48
                        }
                    case "hello":
                        if v6 {
                            var v49 float64 = v4
                            _ = v49
                        }
                    }
                }
            }
            _ = fi24
            _ = v25
        }
        fmt.Println(v4, v1, v4)
        _ = fi22
    }
    for fi50 := 0; fi50 < 18; fi50++ {
        v51 := GenS4{F0: 712, F1: 127, F2: 127}
        _ = fi50
        _ = v51
    }
    fmt.Println(v5, v20)
    fmt.Println(v19, v6)
    var v52 float32 = -2.25
    v53 := GenS5{F0: -1, F1: 60}
    var v54 string = (v19 + (v19 + (((v19 + "hello") + v19) + "hello")))
    v55 := GenS6{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
