package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 bool }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint8, p1 uint16) int16 {
    fmt.Println(p1, p1)
    var v4 uint64 = 18446744073709551615
    v4 -= v4
    p1 = (^(p1))
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 string = "foo bar"
        var v7 uint64 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    return -779
}
func genFn1(p0 int) int8 {
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := []int{}
        var v10 string = "foo bar"
        v9 = append(v9, fi8)
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint32 = 613
    p0 += (-(p0))
    var v12 int64 = 42
    _ = p0
    _ = v11
    _ = v12
    return -116
}

func main() {
    var v1 float64 = ((((0.1 / ((1e10 + -2.25) - float64((1.5 * (1.5 / ((float32((1e10 - 0.1)) - -2.25) / 1.0)))))) / (0.1 / (-2.25 / 1e10))) * (-2.25 * (0.1 + -2.25))) - 1e10)
    var v2 float64 = v1
    var v3 bool = true
    v13 := []int{-9223372036854775808, 984}
    if (true || false) {
        for fi14 := 0; fi14 < 18; fi14++ {
            switch fi14 {
            case 32768:
                v15 := GenS1{F0: -72, F1: 255}
                _ = v15
            case -1:
                for fi16 := 0; fi16 < 14; fi16++ {
                    if ((!v3) && (0 < 997)) {
                        var v17 int = (fi16 & (fi16 >> ((133) & 63)))
                        v17 = v17
                        _ = v17
                    } else {
                        var v18 uint = 255
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        switch (fi16 - 9223372036854775807) {
                        case 9223372036854775807:
                            v2 = (float64((1.5 * 3.14159)) / ((v1 / (v1 * (v1 + (float64((3.14159 / 1e10)) - v2)))) + (-2.25 - float64((0.0 / 1.0)))))
                            if v3 {
                                var v19 uint8 = 127
                                if v3 {
                                    switch "\n\t" {
                                    case "hello":
                                        if (1 <= 631) {
                                            fmt.Println(v3, v19)
                                            v19 -= 127
                                            switch "hello" {
                                            case "":
                                                fmt.Println(v18, v19, v2)
                                            case "hello":
                                                var v20 uint32 = 349
                                                _ = v20
                                            }
                                            fmt.Println(fi16, v2, fi16)
                                        }
                                    case "a":
                                        var v21 string = ("hello" + "\n\t")
                                        _ = v21
                                    default:
                                        var v22 int32 = 1
                                        _ = v22
                                    }
                                    fmt.Println(v1, v18, v3)
                                }
                                _ = v19
                            }
                            v18 += (v18 ^ (v18 % ((v18) | 1)))
                            switch fi14 {
                            case 42:
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                            case -9223372036854775808:
                                var v23 int8 = 42
                                _ = v23
                            case -155:
                                v24 := GenS2{F0: genFn1((fi14 << (((v18 | v18)) & 63))), F1: 128, F2: ("" > "\n\t")}
                                _ = v24
                            }
                        case 128:
                            _ = len(v13)
                        case -894:
                            if (1 != -986) {
                                fmt.Println(fi14, fi14)
                                fmt.Println(fi16, v2, fi16)
                                fmt.Println(v3, v18)
                            }
                        }
                        _ = v18
                    }
                    if (v3 || v3) {
                        v13 = append(v13, (fi16 - (fi14 & 4294967295)))
                    } else {
                        fmt.Println(v3)
                        fmt.Println(fi14, v3, fi16)
                    }
                    fmt.Println(fi16)
                    _ = fi16
                }
            case 2147483647:
                switch fi14 {
                case 2147483647:
                    v25 := GenS3{F0: (^(fi14)), F1: 38}
                    _ = v25
                default:
                    v2 = float64(float32(v2))
                }
            default:
                var v26 int16 = 42
                _ = v26
            }
            if v3 {
                v1 -= (0.1 / v2)
                switch (("\n\t" + "hello") + "a") {
                case "hello":
                    for fi27 := 0; fi27 < 6; fi27++ {
                        var v28 uint = 123
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        var v29 int64 = 9223372036854775807
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                default:
                    if false {
                        if len(v13) > 0 {
                            v13[0] = fi14
                        }
                    }
                }
                for fi30 := 0; fi30 < 10; fi30++ {
                    if true {
                        if ((fi30 << ((4294967295) & 63)) >= (fi30 ^ fi30)) {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            v3 = ((1 != 42) && true)
                            v1 -= v1
                        }
                        if (!v3) {
                            var v31 int = fi14
                            _ = v31
                        } else {
                            fmt.Println(v2, fi14)
                            v32 := GenS2{F0: 127, F1: 42, F2: v3}
                            if false {
                                v33 := GenS4{F0: 2147483648, F1: 4294967295, F2: 2147483647}
                                switch (fi14 ^ -347) {
                                case -1:
                                    var v34 int = fi30
                                    _ = v34
                                case 0:
                                    if v3 {
                                        v2 = ((v1 + float64(0.0)) / (3.14159 / (v1 / 1.5)))
                                    }
                                }
                                _ = v33
                            }
                            _ = v32
                        }
                        v3 = ((v2 + float64(1.5)) == (float64(0.0) - -2.25))
                    } else {
                        fmt.Println(v1)
                        switch ("foo bar" + "a") {
                        case "hello":
                            var v35 uint16 = 2
                            _ = v35
                        case "":
                            v1 = (float64(0.1) * ((1.5 - float64(float32(3.14159))) - 1e10))
                        }
                    }
                    var v36 int = 42
                    _ = fi30
                    _ = v36
                }
            }
            var v37 float32 = float32(0.0)
            _ = fi14
            _ = v37
        }
    } else {
        var v38 int64 = 0
        fmt.Println(v2, v3)
        var v39 bool = false
        _ = v38
        _ = v39
    }
    if v3 {
        var v40 int8 = 100
        if (414 < 887) {
            switch "foo bar" {
            case "":
                fmt.Println(v3, v2, v3)
            case "\n\t":
                for fi41 := 0; fi41 < 14; fi41++ {
                    for fi42 := 0; fi42 < 2; fi42++ {
                        if v3 {
                            v43 := GenS2{F0: (v40 + genFn1(0)), F1: 1, F2: v3}
                            switch "hello" {
                            case "hello":
                                var v44 float64 = ((v1 * float64(3.14159)) + ((float64((1e10 / 1.0)) - v2) - 0.1))
                                _ = v44
                            case "foo bar":
                                fmt.Println(v1, fi42)
                            default:
                                var v45 uint32 = 4294967295
                                _ = v45
                            }
                            v46 := GenS2{F0: 100, F1: 2, F2: v3}
                            _ = v43
                            _ = v46
                        } else {
                            _ = len(v13)
                            fmt.Println(v40, v40)
                        }
                        v1 = 0.0
                        _ = fi42
                    }
                    v3 = ("\n\t" > "a")
                    _ = fi41
                }
            case "hello":
                var v47 int32 = 42
                _ = v47
            default:
                var v48 string = "\n\t"
                _ = v48
            }
            for fi49 := 0; fi49 < 10; fi49++ {
                v50 := GenS5{F0: 1}
                _ = fi49
                _ = v50
            }
        }
        _ = v40
    }
    v3 = v3
    for fi51 := 0; fi51 < 17; fi51++ {
        fmt.Println(v2, v2, v1)
        _ = fi51
    }
    var v52 float32 = -2.25
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v53 := GenS6{F0: 32767}
    v3 = (0 == -1)
    if (0 <= 1) {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        if v3 {
            for fi54 := 0; fi54 < 11; fi54++ {
                v55 := GenS2{F0: 48, F1: 42, F2: ((!(-1 > -32768)) || true)}
                _ = fi54
                _ = v55
            }
            var v56 string = "\n\t"
            _ = v56
        } else {
            if v3 {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                if v3 {
                    v52 -= (3.14159 * float32(v2))
                } else {
                    v57 := GenS7{F0: 1, F1: genFn0(0, 32767)}
                    _ = v57
                }
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
            } else {
                _ = len(v13)
                v1 = ((1e10 / ((v1 - (1.5 + 1e10)) - 0.1)) - (v1 / v2))
                var v58 uint32 = 534
                _ = v58
            }
            for fi59 := 0; fi59 < 12; fi59++ {
                v60 := GenS6{F0: 0}
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                switch ("foo bar" + "\n\t") {
                case "hello":
                    for fi61 := 0; fi61 < 20; fi61++ {
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = fi61
                    }
                case "\n\t":
                    for fi62 := 0; fi62 < 12; fi62++ {
                        fmt.Println(fi59, v3, v3)
                        v2 = ((v1 / (float64(0.1) + float64(0.1))) - v2)
                        v63 := GenS6{F0: -1}
                        _ = fi62
                        _ = v63
                    }
                case "foo bar":
                    if len(v13) > 0 {
                        v13[0] = (-(fi59))
                    }
                }
                _ = fi59
                _ = v60
            }
            for fi64 := 0; fi64 < 13; fi64++ {
                for fi65 := 0; fi65 < 4; fi65++ {
                    v66 := GenS2{F0: 0, F1: 0, F2: ((752 != 1) || (!(255 == 2)))}
                    _ = fi65
                    _ = v66
                }
                _ = fi64
            }
        }
        v67 := GenS2{F0: -128, F1: 127, F2: (((1.5 / v1) * float64(v52)) >= float64(((float32((v1 * (float64(0.0) - v1))) * float32((float64(((1.5 / float32(-2.25)) * v52)) * v1))) - float32(((0.1 / ((((((3.14159 * v1) * v1) / 3.14159) - v2) / v1) - 0.1)) - v2)))))}
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v52
    _ = v53
}
