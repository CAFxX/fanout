package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int16; F1 int; F2 uint }
type GenS5 struct { F0 int16; F1 uint; F2 uint64 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 int; F1 uint16; F2 int8 }

func main() {
    v1 := []int{-615, 9223372036854775807, 127}
    v2 := []string{(("foo bar" + "") + "hello")}
    var v3 uint64 = 1
    var v4 uint32 = 2147483648
    var v5 float64 = 3.14159
    _ = len(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = 2
    v5 = -2.25
    var v6 uint16 = 65535
    for fi7 := 0; fi7 < 12; fi7++ {
        if (2147483647 > 42) {
            var v8 uint = 65535
            var v9 uint8 = 255
            _ = v8
            _ = v9
        } else {
            fmt.Println(fi7, v5)
            v10 := GenS1{F0: uint(v4), F1: 4294967295, F2: 1}
            switch ("a" + "") {
            case "hello":
                if (float32((v5 + 0.1)) < 1e10) {
                    v4 = v4
                    var v11 bool = false
                    _ = v11
                }
            case "a":
                if (18446744073709551615 >= 255) {
                    v5 += float64((3.14159 / 0.1))
                }
            }
            _ = v10
        }
        if false {
            v5 = 0.1
        }
        if ((3.14159 / 3.14159) > (-2.25 - -2.25)) {
            var v12 int64 = -149
            switch (fi7 * fi7) {
            case 9223372036854775807:
                for fi13 := 0; fi13 < 16; fi13++ {
                    v1 = append(v1, (fi13 >> ((1) & 63)))
                    switch "" {
                    case "\n\t":
                        var v14 float32 = ((0.0 + (0.1 * 0.0)) + -2.25)
                        _ = v14
                    case "hello":
                        var v15 int32 = 2147483647
                        _ = v15
                    case "":
                        v6 += v6
                    default:
                        var v16 int = (fi7 & fi7)
                        _ = v16
                    }
                    if ((0.1 / ((-2.25 * (0.1 - 1.5)) + float32(v5))) == -2.25) {
                        v5 += (v5 * ((float64(3.14159) * v5) * (float64((1.5 + 3.14159)) + v5)))
                        var v17 uint8 = 255
                        if false {
                            var v18 int32 = -1
                            var v19 uint64 = (v3 ^ v3)
                            switch (fi13 - (fi13 - (fi7 & fi13))) {
                            case 32767:
                                fmt.Println(v12, v4)
                            case 2147483647:
                                v18 = v18
                            case 886:
                                v12 -= 144
                            default:
                                v12 = (v12 & 9223372036854775807)
                            }
                            v20 := GenS2{F0: 100}
                            fmt.Println(v4)
                            var v21 float32 = 1e10
                            _ = v18
                            _ = v19
                            _ = v20
                            _ = v21
                        } else {
                            _ = len(v2)
                            if (-2.25 != (0.0 * ((1.5 - 0.1) / 1.0))) {
                                v5 -= (v5 * v5)
                                if len(v1) > 0 {
                                    v1[0] = 2147483647
                                }
                            } else {
                                _ = len(v1)
                                fmt.Println(v4, v17)
                                if false {
                                    v12 += (-(v12))
                                    var v22 uint16 = v6
                                    v6 = v22
                                    _ = v22
                                }
                                var v23 int16 = 0
                                _ = v23
                            }
                            if true {
                                fmt.Println(v17, v12, v17)
                            }
                        }
                        _ = v17
                    }
                    v12 -= v12
                    var v24 uint64 = 765
                    var v25 int32 = 1
                    if (0 != fi7) {
                        v5 = 1.5
                    }
                    _ = fi13
                    _ = v24
                    _ = v25
                }
            case 303:
                var v26 uint8 = 2
                _ = v26
            default:
                var v27 int16 = 797
                _ = v27
            }
            fmt.Println(v4)
            _ = v12
        } else {
            if false {
                for fi28 := 0; fi28 < 15; fi28++ {
                    v3 += (v3 | (v3 >> ((4294967295) & 63)))
                    _ = fi28
                }
                var v29 uint32 = (v4 % ((v4) | 1))
                v4 = 2147483648
                _ = v29
            } else {
                if len(v1) > 0 {
                    v1[0] = fi7
                }
            }
            v6 = (^(v6))
        }
        _ = fi7
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 int16 = 42
        v32 := GenS2{F0: 42}
        v3 -= (^(v3))
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := GenS3{F0: 519, F1: 2}
    switch 128 {
    case -2147483648:
        if (!(372 >= 364)) {
            if false {
                v3 = v3
            } else {
                if (!(-497 != 1)) {
                    v34 := GenS4{F0: 1, F1: -9223372036854775808, F2: 461}
                    for fi35 := 0; fi35 < 4; fi35++ {
                        if true {
                            var v36 uint32 = (v4 - v4)
                            var v37 uint8 = 255
                            v38 := GenS3{F0: 4294967295, F1: 0}
                            _ = v36
                            _ = v37
                            _ = v38
                        } else {
                            if (-128 >= -1) {
                                for fi39 := 0; fi39 < 2; fi39++ {
                                    v6 += v6
                                    _ = fi39
                                }
                            }
                            _ = len(v2)
                            v40 := GenS3{F0: 627, F1: 507}
                            _ = v40
                        }
                        var v41 uint64 = (^(v3))
                        _ = fi35
                        _ = v41
                    }
                    _ = v34
                }
            }
            var v42 uint64 = 4294967295
            v42 = 42
            _ = v42
        }
    }
    for fi43 := 0; fi43 < 20; fi43++ {
        var v44 bool = (!false)
        for fi45 := 0; fi45 < 7; fi45++ {
            v4 += 1
            _ = fi45
        }
        _ = fi43
        _ = v44
    }
    if false {
        v4 = v4
        if (772 == 32768) {
            if true {
                v4 -= v4
                v46 := GenS4{F0: 314, F1: 1, F2: 65535}
                _ = v46
            } else {
                v47 := GenS5{F0: 32767, F1: 0, F2: 225}
                switch 140 {
                case 42:
                    if (!false) {
                        v48 := GenS6{F0: 0, F1: (v6 & 42)}
                        _ = len(v2)
                        v5 = v5
                        _ = v48
                    }
                case -2147483648:
                    v2 = append(v2, "a")
                case -1:
                    fmt.Println(v5)
                }
                var v49 string = "\n\t"
                _ = v47
                _ = v49
            }
            v4 += v4
            v50 := GenS7{F0: 0, F1: (v6 << ((0) & 63)), F2: -1}
            _ = v50
        } else {
            v5 = (((float64(1.5) + v5) - -2.25) / 0.1)
            v3 = v3
            if ((127 <= 255) && false) {
                for fi51 := 0; fi51 < 2; fi51++ {
                    v3 = v3
                    v4 += uint32(v3)
                    switch (fi51 ^ (fi51 >> ((0) & 63))) {
                    case 756:
                        v5 -= v5
                    case 51:
                        _ = len(v2)
                    case 0:
                        for fi52 := 0; fi52 < 9; fi52++ {
                            v5 = 0.0
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = ""
                                }
                            } else {
                                if false {
                                    v2 = append(v2, (("" + "foo bar") + "hello"))
                                } else {
                                    v3 += (v3 ^ v3)
                                    var v53 uint64 = (v3 >> ((uint(v4)) & 63))
                                    if false {
                                        v6 = 0
                                        if (0 <= 1) {
                                            v54 := GenS3{F0: 693, F1: 1}
                                            fmt.Println(v53, fi52)
                                            _ = v54
                                        }
                                    }
                                    _ = v53
                                }
                                fmt.Println(v5, fi52)
                                switch "\n\t" {
                                case "foo bar":
                                    v6 -= (^(v6))
                                case "\n\t":
                                    var v55 int16 = 0
                                    _ = v55
                                default:
                                    var v56 float32 = 0.0
                                    _ = v56
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v57 := GenS2{F0: -121}
                                _ = v57
                            }
                            v6 -= 1
                            v3 = (v3 | v3)
                            _ = fi52
                        }
                    }
                    _ = fi51
                }
                for fi58 := 0; fi58 < 19; fi58++ {
                    if (v3 < (v3 * (v3 + (v3 % (((v3 | (v3 ^ (v3 % ((67) | 1))))) | 1))))) {
                        switch fi58 {
                        case -336:
                            switch "hello" {
                            case "a":
                                var v59 int = fi58
                                _ = v59
                            default:
                                v3 = 4294967295
                            }
                        case 2147483647:
                            v1 = append(v1, int(v6))
                        default:
                            fmt.Println(v6)
                        }
                        switch "foo bar" {
                        case "":
                            switch fi58 {
                            case 2147483647:
                                _ = len(v2)
                            }
                        case "hello":
                            fmt.Println(v5, fi58, v3)
                        default:
                            fmt.Println(fi58, v4)
                        }
                        var v60 bool = ((fi58 & fi58) <= fi58)
                        _ = v60
                    }
                    _ = fi58
                }
                v61 := GenS5{F0: 42, F1: 661, F2: (^(v3))}
                _ = v61
            }
        }
    } else {
        for fi62 := 0; fi62 < 1; fi62++ {
            var v63 int = -64
            _ = fi62
            _ = v63
        }
        v5 += ((v5 - float64((float32(-2.25) + (float32(0.1) + float32(v5))))) / 1.5)
        v3 += 0
    }
    switch "\n\t" {
    case "a":
        v6 -= 255
    case "":
        v5 += ((3.14159 / v5) * v5)
    }
    fmt.Println(v4, v4, v3)
    for fi64 := 0; fi64 < 5; fi64++ {
        if true {
            var v65 int = -546
            v5 = v5
            v4 = uint32(v3)
            _ = v65
        }
        _ = fi64
    }
    v66 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: 138}
    fmt.Println(v4, v6, v5)
    v2 = append(v2, ((("" + "") + "foo bar") + "a"))
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v66
}
