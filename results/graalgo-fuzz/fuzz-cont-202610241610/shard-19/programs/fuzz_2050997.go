package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint16; F1 int8; F2 int16 }
type GenS4 struct { F0 uint16; F1 int64; F2 int }
type GenS5 struct { F0 int64 }

func main() {
    var v1 int64 = -979
    var v2 bool = (42 <= -32768)
    v3 := []string{"a"}
    var v4 uint32 = 2147483647
    v2 = (!true)
    if len(v3) > 0 {
        v3[0] = ""
    }
    v2 = v2
    v4 = v4
    if false {
        v3 = append(v3, "")
    } else {
        for fi5 := 0; fi5 < 11; fi5++ {
            if (1.5 == (-2.25 * (float32(1e10) - 0.0))) {
                v6 := GenS1{F0: (fi5 / ((fi5) | 1)), F1: "\n\t", F2: (-(fi5))}
                _ = v6
            } else {
                v4 -= (v4 ^ (v4 - v4))
                if (v4 != (v4 - (^(v4)))) {
                    var v7 string = "a"
                    switch fi5 {
                    case 4294967295:
                        var v8 int32 = -2147483648
                        _ = v8
                    case -665:
                        if true {
                            fmt.Println(v4, v7, v7)
                            v3 = append(v3, "")
                            if false {
                                fmt.Println(v2)
                                if v2 {
                                    for fi9 := 0; fi9 < 7; fi9++ {
                                        if (v4 != v4) {
                                            var v10 int16 = 132
                                            v7 = v7
                                            switch fi5 {
                                            case -351:
                                                var v11 uint8 = 127
                                                _ = v11
                                            default:
                                                v2 = (!(720 != -1))
                                            }
                                            _ = v10
                                        }
                                        v12 := GenS2{F0: 2, F1: 0}
                                        _ = fi9
                                        _ = v12
                                    }
                                    v2 = ((v4 | v4) == v4)
                                    if (18446744073709551615 != 0) {
                                        fmt.Println(v1, fi5)
                                    } else {
                                        var v13 int = fi5
                                        _ = v13
                                    }
                                } else {
                                    switch (v7 + v7) {
                                    case "foo bar":
                                        v3 = append(v3, "")
                                    case "\n\t":
                                        var v14 uint = 1
                                        _ = v14
                                    default:
                                        for fi15 := 0; fi15 < 1; fi15++ {
                                            if (42 != 42) {
                                                v16 := GenS2{F0: uint64(v4), F1: -51}
                                                v17 := GenS3{F0: 65535, F1: 100, F2: 0}
                                                _ = v16
                                                _ = v17
                                            }
                                            switch fi15 {
                                            case 127:
                                                v18 := GenS2{F0: 286, F1: -1}
                                                _ = v18
                                            default:
                                                v2 = (-46 <= 100)
                                                switch v7 {
                                                case "foo bar":
                                                    var v19 float32 = 1.5
                                                    _ = v19
                                                case "":
                                                    var v20 float64 = ((((float64(0.1) + ((-2.25 / 1.0) - 0.0)) / -2.25) + 0.1) / (3.14159 + 3.14159))
                                                    _ = v20
                                                default:
                                                    _ = len(v3)
                                                }
                                            }
                                            _ = len(v3)
                                            v1 = 835
                                            v1 = v1
                                            v2 = ((v1 >> ((255) & 63)) >= v1)
                                            _ = fi15
                                        }
                                    }
                                }
                                for fi21 := 0; fi21 < 17; fi21++ {
                                    fmt.Println(v4, v4, v1)
                                    _ = fi21
                                }
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch fi5 {
                                case 9223372036854775807:
                                    v2 = (v2 || (255 >= 1))
                                case -1:
                                    v4 = 65535
                                case 32768:
                                    for fi22 := 0; fi22 < 10; fi22++ {
                                        if (1.5 >= 1e10) {
                                            var v23 float64 = 0.0
                                            if (((-2.25 * float32(((float64(1e10) - ((((0.1 * float64((((3.14159 / (3.14159 - 1.5)) + (float32(v23) / 1.0)) / 0.1))) + float64((0.1 / (1.5 * -2.25)))) / float64(3.14159)) + v23)) + v23))) - float32(-2.25)) <= 0.0) {
                                                var v24 int32 = 899
                                                _ = v24
                                            } else {
                                                v23 = v23
                                            }
                                            var v25 uint32 = 2
                                            _ = v23
                                            _ = v25
                                        }
                                        var v26 int8 = -128
                                        _ = fi22
                                        _ = v26
                                    }
                                default:
                                    v27 := GenS2{F0: 0, F1: -128}
                                    _ = v27
                                }
                            }
                        } else {
                            var v28 uint32 = (^(v4))
                            v2 = (18446744073709551615 >= 377)
                            _ = v28
                        }
                    }
                    _ = v7
                } else {
                    v3 = append(v3, (("hello" + "\n\t") + "hello"))
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            v2 = ((-2.25 - ((-2.25 / float64((float32(((-2.25 * (3.14159 + (float64((3.14159 / ((1e10 - 0.1) / 1.0))) * (1.5 - 3.14159)))) - 0.1)) * 3.14159))) - ((((3.14159 - 1e10) * 3.14159) * -2.25) - (-2.25 - (float64((3.14159 * 3.14159)) * float64(1e10)))))) > (0.1 * (1.5 - 1e10)))
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
            _ = fi5
        }
        var v29 uint32 = v4
        _ = v29
    }
    if (!((v1 & v1) < (v1 & 148))) {
        v3 = append(v3, "a")
        if true {
            if v2 {
                v30 := GenS3{F0: 0, F1: int8(v4), F2: 0}
                v1 = v1
                fmt.Println(v4)
                _ = v30
            } else {
                v31 := GenS1{F0: -2147483648, F1: "\n\t", F2: 128}
                if (32767 != 2) {
                    var v32 int8 = 42
                    _ = v32
                }
                var v33 int32 = 0
                _ = v31
                _ = v33
            }
            v2 = (v2 && (v1 <= v1))
            v34 := GenS2{F0: 1, F1: 127}
            _ = v34
        }
        var v35 bool = false
        _ = v35
    } else {
        v2 = (!(2 >= 127))
        switch 32767 {
        case 2147483647:
            var v36 int = -993
            _ = v36
        case -1:
            v37 := GenS1{F0: 1, F1: "foo bar", F2: 127}
            _ = v37
        default:
            v1 = v1
        }
        v2 = v2
    }
    v38 := GenS1{F0: 128, F1: "\n\t", F2: -112}
    var v39 uint16 = 337
    if (0.1 == 0.0) {
        for fi40 := 0; fi40 < 1; fi40++ {
            var v41 uint = 495
            _ = fi40
            _ = v41
        }
        switch "foo bar" {
        case "\n\t":
            for fi42 := 0; fi42 < 5; fi42++ {
                fmt.Println(fi42, v39)
                for fi43 := 0; fi43 < 5; fi43++ {
                    fmt.Println(v4)
                    _ = fi43
                }
                _ = fi42
            }
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            v1 += v1
        }
    } else {
        for fi44 := 0; fi44 < 17; fi44++ {
            if len(v3) > 0 {
                v3[0] = (("" + ("a" + "hello")) + "")
            }
            _ = fi44
        }
        v2 = v2
        var v45 uint8 = 255
        _ = v45
    }
    for fi46 := 0; fi46 < 13; fi46++ {
        var v47 uint32 = v4
        _ = fi46
        _ = v47
    }
    switch 9223372036854775807 {
    case 127:
        v4 -= v4
    case -1:
        v48 := GenS4{F0: 137, F1: (-(v1)), F2: 1}
        _ = v48
    }
    var v49 bool = (!(!true))
    _ = len(v3)
    v50 := GenS2{F0: 4294967295, F1: 127}
    if v2 {
        v49 = (1e10 > 3.14159)
        fmt.Println(v4, v1, v49)
    }
    v51 := GenS5{F0: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v49
    _ = v50
    _ = v51
}
