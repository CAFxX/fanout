package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int32; F1 int32 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 uint64 }

func main() {
    var v1 uint64 = 308
    var v2 float64 = (0.1 + 1.5)
    var v3 string = (("\n\t" + "") + ("foo bar" + "a"))
    var v4 uint32 = 4294967295
    v5 := make(map[string]int)
    var v6 int = -1
    v1 = (^(v1))
    var v7 uint64 = v1
    v8 := []int{(v6 & (v6 * 32768)), 9223372036854775807, (v6 >> ((65535) & 63))}
    if (!true) {
        for fi9 := 0; fi9 < 20; fi9++ {
            v3 = v3
            if true {
                var v10 int8 = 1
                fmt.Println(v1, v2, v2)
                _ = v10
            }
            _ = fi9
        }
        var v11 float64 = (float64(-2.25) / float64(0.1))
        var v12 uint = 65535
        _ = v11
        _ = v12
    }
    fmt.Println(v7, v7, v7)
    switch (v6 | v6) {
    case 42:
        fmt.Println(v3, v1, v2)
    case 1:
        if false {
            v13 := v5[(v3 + v3)]
            _ = v13
        } else {
            fmt.Println(v6)
        }
    case -1:
        _ = len(v5)
    default:
        if true {
            for fi14 := 0; fi14 < 12; fi14++ {
                if (!false) {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        var v16 float32 = 1.5
                        _ = fi15
                        _ = v16
                    }
                    switch 857 {
                    case 9223372036854775807:
                        var v17 uint32 = (v4 ^ v4)
                        _ = v17
                    case -1:
                        v2 = float64((float32(-2.25) * float32(1e10)))
                    case -2147483648:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    }
                } else {
                    v18 := GenS1{F0: 208, F1: 42, F2: 42}
                    fmt.Println(v4, v7, v6)
                    v19 := GenS1{F0: 2, F1: -24, F2: -635}
                    _ = v18
                    _ = v19
                }
                switch fi14 {
                case 42:
                    v8 = append(v8, fi14)
                case 261:
                    var v20 int32 = -1
                    _ = v20
                default:
                    v7 += (v7 % (((v1 * 488)) | 1))
                }
                _ = fi14
            }
            v21 := GenS1{F0: 745, F1: -9223372036854775808, F2: -1}
            if (false || true) {
                v7 = 18446744073709551615
                fmt.Println(v4)
            }
            _ = v21
        } else {
            var v22 uint16 = 65535
            delete(v5, v3)
            _ = v22
        }
    }
    switch v6 {
    case -62:
        var v23 uint64 = v1
        _ = v23
    }
    v24 := GenS2{F0: 1}
    for fi25 := 0; fi25 < 20; fi25++ {
        if (1 >= 2) {
            if false {
                if true {
                    var v26 int32 = 2147483647
                    v2 += v2
                    v27 := GenS3{F0: 65535}
                    _ = v26
                    _ = v27
                }
                v7 += (v7 >> ((255) & 63))
                var v28 uint16 = 99
                _ = v28
            } else {
                var v29 uint16 = 155
                v30 := v5["\n\t"]
                _ = v29
                _ = v30
            }
        }
        _ = fi25
    }
    if true {
        switch (v6 + int(v7)) {
        case -9223372036854775808:
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 int64 = -1
                if true {
                    v1 += (v7 | (v1 & 1))
                    for fi33 := 0; fi33 < 7; fi33++ {
                        if true {
                            v4 = v4
                            var v34 uint = 1
                            v5[((v3 + v3) + v3)] = fi33
                            switch (v3 + (v3 + (v3 + (((v3 + v3) + "\n\t") + "\n\t")))) {
                            case "\n\t":
                                var v35 uint16 = uint16(v4)
                                _ = v35
                            case "a":
                                var v36 int16 = 42
                                _ = v36
                            case "foo bar":
                                v37 := GenS4{F0: 657, F1: 2}
                                _ = v37
                            default:
                                fmt.Println(v4, v1, v4)
                            }
                            _ = v34
                        } else {
                            v38 := GenS3{F0: 255}
                            _ = v38
                        }
                        v39 := v5[(("\n\t" + (v3 + v3)) + "hello")]
                        _ = fi33
                        _ = v39
                    }
                    var v40 int8 = 42
                    _ = v40
                }
                _ = fi31
                _ = v32
            }
        case -2147483648:
            var v41 int32 = 967
            _ = v41
        case 32768:
            fmt.Println(v6, v3, v2)
        default:
            for fi42 := 0; fi42 < 15; fi42++ {
                if (((((0.0 / -2.25) + float32(v2)) / 1.0) / 1e10) == 1e10) {
                    v1 += v7
                } else {
                    _ = len(v5)
                    if true {
                        var v43 uint32 = (v4 | v4)
                        _ = v43
                    }
                }
                fmt.Println(v4, v6)
                _ = fi42
            }
        }
    } else {
        v7 -= 4294967295
        var v44 uint64 = (v1 << ((4294967295) & 63))
        _ = v44
    }
    v45 := GenS3{F0: 255}
    for fi46 := 0; fi46 < 6; fi46++ {
        for fi47 := 0; fi47 < 8; fi47++ {
            var v48 int8 = 100
            v8 = append(v8, v6)
            v2 = v2
            _ = fi47
            _ = v48
        }
        v3 = ("foo bar" + "\n\t")
        if (false || true) {
            for fi49 := 0; fi49 < 20; fi49++ {
                var v50 string = v3
                _ = fi49
                _ = v50
            }
            if (!(!false)) {
                v7 += (v1 | (v7 << ((4294967295) & 63)))
                switch (^(fi46)) {
                case 4294967295:
                    switch "foo bar" {
                    case "":
                        for fi51 := 0; fi51 < 3; fi51++ {
                            v52 := GenS5{F0: -1, F1: 1}
                            v1 -= 1
                            var v53 int8 = 1
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                    case "hello":
                        var v54 uint = 0
                        _ = v54
                    case "foo bar":
                        for fi55 := 0; fi55 < 20; fi55++ {
                            var v56 uint = 864
                            v57 := v5[("hello" + "foo bar")]
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                    default:
                        v58 := GenS6{F0: (v4 ^ (^(v4))), F1: 1}
                        _ = v58
                    }
                }
                for fi59 := 0; fi59 < 13; fi59++ {
                    fmt.Println(fi46, v6, v7)
                    var v60 int16 = -32768
                    _ = fi59
                    _ = v60
                }
            } else {
                if false {
                    v4 = v4
                    var v61 int16 = -180
                    v4 = v4
                    _ = v61
                }
                v62 := GenS7{F0: (^(v7))}
                v63 := GenS4{F0: 963, F1: 255}
                _ = v62
                _ = v63
            }
        } else {
            if true {
                v4 -= (v4 / ((248) | 1))
                switch v6 {
                case 0:
                    fmt.Println(fi46, v1, v7)
                }
                v6 += int(v4)
            } else {
                fmt.Println(v2)
            }
            var v64 float32 = 0.0
            _ = v64
        }
        _ = fi46
    }
    if ((0 != 276) || true) {
        switch v6 {
        case 42:
            if (((v1 ^ 1) != 2) || (!(v2 > (v2 * v2)))) {
                switch v3 {
                case "foo bar":
                    if len(v8) > 0 {
                        v8[0] = (v6 / (((v6 << ((65535) & 63))) | 1))
                    }
                }
                if true {
                    if false {
                        for fi65 := 0; fi65 < 15; fi65++ {
                            v7 -= 42
                            _ = fi65
                        }
                    } else {
                        if (2147483647 != int(v4)) {
                            var v66 uint16 = 65535
                            if (!true) {
                                v3 = (("foo bar" + "") + v3)
                                var v67 uint = 2
                                for fi68 := 0; fi68 < 19; fi68++ {
                                    v69 := GenS2{F0: uint16(v4)}
                                    v6 = v6
                                    _ = fi68
                                    _ = v69
                                }
                                _ = v67
                            }
                            v70 := GenS3{F0: 2}
                            _ = v66
                            _ = v70
                        }
                    }
                    var v71 bool = false
                    _ = v71
                }
            } else {
                var v72 uint16 = 2
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                switch v3 {
                case "":
                    fmt.Println(v2, v7, v3)
                case "hello":
                    v6 = v6
                case "foo bar":
                    var v73 uint8 = 127
                    _ = v73
                }
                _ = v72
            }
        case -1:
            if true {
                var v74 int = -401
                var v75 int16 = 42
                _ = v74
                _ = v75
            } else {
                fmt.Println(v6)
                fmt.Println(v7)
            }
        case -884:
            var v76 uint16 = 255
            _ = v76
        }
    }
    fmt.Println(v7, v7)
    var v77 float64 = (float64(0.1) + 1.5)
    v78 := v5[("" + "hello")]
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
    _ = v77
    _ = v78
    _ = v8
}
