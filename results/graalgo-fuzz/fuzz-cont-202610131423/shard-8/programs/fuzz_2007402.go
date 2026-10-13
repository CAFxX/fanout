package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 int16 }
type GenS7 struct { F0 uint32; F1 int16 }
func genFn0(p0 int64, p1 string, p2 int) int {
    v5 := GenS1{F0: true}
    if false {
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"", ("" + ("hello" + "hello")), ("a" + "foo bar"), "\n\t"}
    var v3 int32 = 42
    var v4 uint32 = 42
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int32 = (-(v3))
            v3 += (v3 >> ((0) & 63))
            _ = fi7
            _ = v8
        }
        var v9 uint64 = 42
        _ = fi6
        _ = v9
    }
    var v10 uint32 = 362
    v4 = (v10 ^ (v4 / (((v10 - v4)) | 1)))
    v10 -= 17
    var v11 uint = 255
    v12 := GenS1{F0: (121 > 59)}
    for fi13 := 0; fi13 < 4; fi13++ {
        if true {
            switch fi13 {
            case -2147483648:
                for fi14 := 0; fi14 < 16; fi14++ {
                    v15 := GenS2{F0: 247, F1: 159}
                    _ = fi14
                    _ = v15
                }
            default:
                _ = len(v1)
            }
            v4 = 988
            var v16 uint64 = 18446744073709551615
            _ = v16
        } else {
            if len(v2) > 0 {
                v2[0] = (("hello" + "\n\t") + "\n\t")
            }
            if (!false) {
                var v17 int32 = 1
                v10 = 497
                fmt.Println(v3, fi13)
                _ = v17
            }
            v10 = (^(v4))
        }
        if (!true) {
            v18 := GenS1{F0: (0.1 <= (-2.25 * (float64(0.0) * 1.5)))}
            _ = v18
        }
        switch (fi13 * (fi13 << (((^(v11))) & 63))) {
        case 32768:
            v3 = v3
        case -1:
            v3 = v3
        case 232:
            for fi19 := 0; fi19 < 13; fi19++ {
                if false {
                    v4 = (v4 % ((v10) | 1))
                    fmt.Println(fi13)
                } else {
                    v3 += v3
                    v3 = -1
                    switch "hello" {
                    case "hello":
                        if false {
                            if true {
                                v1[v4] = (2 < 598)
                            } else {
                                var v20 int16 = 32767
                                _ = v20
                            }
                        } else {
                            if (520 != -32768) {
                                v21 := GenS3{F0: v4}
                                v22 := GenS2{F0: (v11 - v11), F1: 42}
                                _ = v21
                                _ = v22
                            } else {
                                if false {
                                    var v23 float32 = ((0.0 + -2.25) * float32(-2.25))
                                    v24 := GenS3{F0: (v4 >> ((v11) & 63))}
                                    _ = v23
                                    _ = v24
                                }
                                v10 += v4
                            }
                        }
                    case "foo bar":
                        if (false || ((-32768 != 0) && true)) {
                            var v25 int16 = 32767
                            v3 = v3
                            _ = v25
                        }
                    }
                    var v26 float64 = ((((((float64(1.5) + 0.1) + 3.14159) * (0.0 - float64(float32((float64(1e10) + ((3.14159 + 1e10) - 1e10)))))) * 1.5) - 0.0) / 1.0)
                    _ = v26
                }
                v27 := GenS2{F0: v11, F1: 0}
                var v28 float32 = ((0.0 / 1.5) - 0.1)
                if len(v2) > 0 {
                    v2[0] = ""
                }
                _ = fi19
                _ = v27
                _ = v28
            }
        }
        _ = fi13
    }
    v3 -= -1
    v29 := GenS2{F0: (v11 * v11), F1: 185}
    for fi30 := 0; fi30 < 19; fi30++ {
        if true {
            v4 = (v4 + (v4 | (v4 << ((0) & 63))))
            for fi31 := 0; fi31 < 20; fi31++ {
                v11 = (v11 * (v11 - uint(v10)))
                var v32 uint16 = 394
                _ = fi31
                _ = v32
            }
            for fi33 := 0; fi33 < 11; fi33++ {
                v3 = v3
                v3 = (v3 & v3)
                switch "a" {
                case "":
                    v34 := GenS3{F0: (v4 + (^(v10)))}
                    _ = v34
                case "hello":
                    var v35 int16 = 42
                    _ = v35
                case "foo bar":
                    var v36 int = genFn0(-303, "", genFn0(-9223372036854775808, "\n\t", int(v3)))
                    _ = v36
                default:
                    if (true && true) {
                        if (65535 != 255) {
                            fmt.Println(fi33, v3)
                        }
                        v4 += uint32(v11)
                        var v37 int = fi33
                        _ = v37
                    }
                }
                var v38 bool = (0.1 != (1e10 + 3.14159))
                fmt.Println(v3, v4)
                _ = fi33
                _ = v38
            }
        }
        _ = fi30
    }
    v2 = append(v2, "hello")
    v10 = v10
    delete(v1, v4)
    if (339 <= 14) {
        switch genFn0(1, ("hello" + "foo bar"), -2147483648) {
        case -574:
            v3 -= (-(v3))
        }
        if ("" > "") {
            for fi39 := 0; fi39 < 19; fi39++ {
                var v40 uint32 = (v4 * 2147483647)
                _ = fi39
                _ = v40
            }
            v11 -= v11
            if (true && false) {
                var v41 uint64 = 18446744073709551615
                v41 += v41
                if (3.14159 != 1.5) {
                    if false {
                        _ = len(v1)
                        v11 -= v11
                        v1[2] = false
                    }
                    if true {
                        var v42 int64 = 122
                        v42 += v42
                        _ = v42
                    }
                    v4 = uint32(v11)
                } else {
                    switch 128 {
                    case -9223372036854775808:
                        v41 = v41
                    case 2147483647:
                        for fi43 := 0; fi43 < 12; fi43++ {
                            if ((!false) && ((!(((v10 == v4) && true) || true)) || (int(v41) > fi43))) {
                                if false {
                                    for fi44 := 0; fi44 < 13; fi44++ {
                                        v3 = 2147483647
                                        if (255 == 797) {
                                            v11 = v11
                                        } else {
                                            var v45 string = ""
                                            _ = v45
                                        }
                                        _ = fi44
                                    }
                                    if (v3 == (v3 + (-(v3)))) {
                                        fmt.Println(fi43)
                                    } else {
                                        if false {
                                            for fi46 := 0; fi46 < 9; fi46++ {
                                                fmt.Println(v41, v4)
                                                _ = fi46
                                            }
                                            if true {
                                                v47 := GenS4{F0: v3, F1: "foo bar"}
                                                var v48 float32 = float32(((0.0 - ((((1.5 + (0.1 + ((0.0 + ((((0.0 / -2.25) / 0.1) - 1.5) + 0.1)) / 1.5))) + -2.25) / (-2.25 - 0.0)) * 3.14159)) * 0.1))
                                                v10 += v10
                                                _ = v47
                                                _ = v48
                                            }
                                        } else {
                                            v4 = v4
                                        }
                                    }
                                    if false {
                                        delete(v1, (v4 / ((v10) | 1)))
                                        var v49 int = 503
                                        switch 18 {
                                        case 127:
                                            if len(v2) > 0 {
                                                v2[0] = ("foo bar" + "\n\t")
                                            }
                                        case 746:
                                            var v50 uint16 = 65535
                                            _ = v50
                                        case 9223372036854775807:
                                            var v51 uint16 = 0
                                            _ = v51
                                        }
                                        _ = v49
                                    }
                                }
                                if false {
                                    for fi52 := 0; fi52 < 5; fi52++ {
                                        v53 := GenS5{F0: 127}
                                        _ = fi52
                                        _ = v53
                                    }
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    if (9223372036854775807 != 1) {
                                        v4 = (v4 / (((v4 * 65535)) | 1))
                                        v4 -= (v4 & v10)
                                        if true {
                                            fmt.Println(v3, v4, v4)
                                        } else {
                                            fmt.Println(v41, v11)
                                        }
                                    } else {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                }
                            }
                            _ = fi43
                        }
                    }
                }
                _ = v41
            } else {
                if false {
                    if (1 <= 0) {
                        for fi54 := 0; fi54 < 4; fi54++ {
                            for fi55 := 0; fi55 < 18; fi55++ {
                                v56 := GenS6{F0: 54, F1: 0}
                                var v57 int = (^(fi55))
                                fmt.Println(v4, v57, fi55)
                                _ = fi55
                                _ = v56
                                _ = v57
                            }
                            var v58 uint16 = 42
                            fmt.Println(v58, v10)
                            _ = fi54
                            _ = v58
                        }
                        v11 -= (^(v11))
                        if false {
                            _ = len(v2)
                            var v59 int16 = 1
                            var v60 uint16 = 516
                            _ = v59
                            _ = v60
                        }
                    }
                    for fi61 := 0; fi61 < 17; fi61++ {
                        for fi62 := 0; fi62 < 2; fi62++ {
                            v11 += 65535
                            _ = fi62
                        }
                        v63 := GenS2{F0: (v11 % ((v11) | 1)), F1: 1}
                        fmt.Println(v4, v3, v11)
                        _ = fi61
                        _ = v63
                    }
                } else {
                    v4 = v10
                    for fi64 := 0; fi64 < 7; fi64++ {
                        var v65 int = fi64
                        var v66 int64 = -984
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                }
            }
        }
        fmt.Println(v3)
    }
    v1[306] = ((!true) && ((670 >= v11) && (-2.25 <= 3.14159)))
    if (v3 > (v3 << (((v11 | v11)) & 63))) {
        fmt.Println(v4, v10)
        fmt.Println(v11, v3, v10)
    } else {
        fmt.Println(v4, v3)
        switch "hello" {
        case "hello":
            var v67 bool = false
            _ = v67
        default:
            v68 := GenS7{F0: 2, F1: 479}
            _ = v68
        }
    }
    v10 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
