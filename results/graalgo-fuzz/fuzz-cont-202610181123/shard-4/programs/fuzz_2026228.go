package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 string }
type GenS5 struct { F0 string; F1 int16; F2 uint }
type GenS6 struct { F0 uint8; F1 uint; F2 int }
type GenS7 struct { F0 bool; F1 uint; F2 int64 }
func genFn0() int {
    var v6 string = "foo bar"
    switch 711 {
    case 32767:
        v6 = "hello"
    case -48:
        v7 := []int{9223372036854775807, -9223372036854775808}
        _ = v7
    case 2147483647:
        v6 = "a"
    default:
        v6 = (v6 + "a")
    }
    v6 = ("" + "a")
    fmt.Println(v6)
    _ = v6
    return 128
}

func main() {
    var v1 int16 = 42
    v2 := []int{240, -9223372036854775808, 127}
    var v3 uint = 255
    var v4 int16 = v1
    v5 := make(map[int]bool)
    v4 = -299
    for fi8 := 0; fi8 < 10; fi8++ {
        switch (fi8 << (((v3 - (v3 & v3))) & 63)) {
        case 32767:
            for fi9 := 0; fi9 < 13; fi9++ {
                if len(v2) > 0 {
                    v2[0] = genFn0()
                }
                switch fi8 {
                case 42:
                    fmt.Println(v1, fi9)
                default:
                    v2 = append(v2, genFn0())
                }
                var v10 int32 = 870
                _ = fi9
                _ = v10
            }
        case -1:
            if ((3.14159 * (1e10 + ((1.5 + 1e10) + float32(0.1)))) < float32(float64(3.14159))) {
                var v11 uint = 18446744073709551615
                v3 = 4294967295
                var v12 uint = v3
                _ = v11
                _ = v12
            } else {
                v13 := v5[genFn0()]
                _ = v13
            }
        default:
            for fi14 := 0; fi14 < 16; fi14++ {
                switch 42 {
                case -332:
                    if (-835 >= -814) {
                        v15 := GenS1{F0: 2147483647}
                        if (1 != 1) {
                            v3 = 2
                            if (117 >= 127) {
                                delete(v5, 128)
                            }
                        } else {
                            var v16 float64 = float64((1.5 / 1.5))
                            fmt.Println(fi8)
                            _ = v16
                        }
                        _ = v15
                    } else {
                        v17 := GenS1{F0: 2147483647}
                        v18 := GenS1{F0: -392}
                        v19 := GenS1{F0: -2147483648}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                case 32768:
                    v20 := GenS1{F0: -69}
                    _ = v20
                default:
                    var v21 int = 128
                    _ = v21
                }
                var v22 float32 = 1e10
                _ = fi14
                _ = v22
            }
        }
        v23 := GenS2{F0: (v3 * (^(v3))), F1: -128}
        v5[(fi8 * (fi8 ^ (fi8 >> ((v3) & 63))))] = ((^(v4)) > (v1 ^ v4))
        _ = fi8
        _ = v23
    }
    if false {
        var v24 bool = false
        if (!v24) {
            v25 := GenS1{F0: 0}
            v24 = (3.14159 < 0.1)
            switch "a" {
            case "hello":
                if (v24 && v24) {
                    v5[-456] = ("hello" > "a")
                    for fi26 := 0; fi26 < 2; fi26++ {
                        var v27 int16 = (v1 ^ 42)
                        if ((42 < -687) || v24) {
                            for fi28 := 0; fi28 < 19; fi28++ {
                                v29 := GenS2{F0: (v3 ^ (v3 - 18446744073709551615)), F1: 0}
                                _ = fi28
                                _ = v29
                            }
                            for fi30 := 0; fi30 < 18; fi30++ {
                                fmt.Println(v3)
                                _ = fi30
                            }
                            var v31 float32 = (-2.25 - (3.14159 / 1.5))
                            _ = v31
                        } else {
                            v32 := GenS1{F0: 2147483647}
                            v4 = v27
                            _ = v32
                        }
                        _ = fi26
                        _ = v27
                    }
                } else {
                    switch genFn0() {
                    case 2147483647:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -619:
                        var v33 int16 = -32768
                        _ = v33
                    default:
                        for fi34 := 0; fi34 < 14; fi34++ {
                            for fi35 := 0; fi35 < 15; fi35++ {
                                fmt.Println(v24, v24, fi34)
                                _ = fi35
                            }
                            v36 := GenS3{F0: 4294967295}
                            v37 := GenS2{F0: (v3 | 18446744073709551615), F1: 1}
                            _ = fi34
                            _ = v36
                            _ = v37
                        }
                    }
                    v3 = (v3 << ((v3) & 63))
                    if v24 {
                        var v38 uint8 = 255
                        v1 = (v1 & v1)
                        var v39 uint16 = 1
                        _ = v38
                        _ = v39
                    } else {
                        v3 = (v3 * (v3 >> ((1) & 63)))
                    }
                }
            }
            _ = v25
        }
        _ = v24
    } else {
        v40 := GenS4{F0: 42, F1: "\n\t"}
        _ = len(v2)
        _ = v40
    }
    v5[32767] = (0 != 634)
    if ((((1.5 - ((float64(1.5) / (0.1 + -2.25)) - 3.14159)) * float64(((-2.25 * 0.1) + (((1e10 - 3.14159) + 1e10) - 1.5)))) * 0.0) >= 0.1) {
        v41 := v5[141]
        v42 := GenS3{F0: 18446744073709551615}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v41
        _ = v42
    }
    switch "\n\t" {
    case "foo bar":
        var v43 float32 = 0.1
        _ = v43
    case "a":
        v4 = -871
    }
    switch (("\n\t" + ("foo bar" + ("\n\t" + ""))) + "foo bar") {
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "hello":
        fmt.Println(v4, v3, v4)
    case "\n\t":
        if false {
            v1 -= 42
        } else {
            for fi44 := 0; fi44 < 9; fi44++ {
                v1 = v1
                _ = fi44
            }
            if (true && (!((v3 >> ((770) & 63)) >= 18446744073709551615))) {
                v4 -= v1
            }
            if (-9223372036854775808 == 419) {
                for fi45 := 0; fi45 < 9; fi45++ {
                    v3 -= 0
                    v2 = append(v2, -2147483648)
                    if (65535 != (v3 & v3)) {
                        fmt.Println(v4, fi45, v3)
                    }
                    _ = fi45
                }
                var v46 bool = ((!false) || (!false))
                _ = v46
            }
        }
    default:
        fmt.Println(v4, v3, v1)
    }
    for fi47 := 0; fi47 < 6; fi47++ {
        v48 := GenS2{F0: v3, F1: 0}
        var v49 uint32 = 513
        var v50 float64 = 3.14159
        _ = fi47
        _ = v48
        _ = v49
        _ = v50
    }
    for fi51 := 0; fi51 < 12; fi51++ {
        for fi52 := 0; fi52 < 4; fi52++ {
            switch 127 {
            case -2147483648:
                v53 := GenS3{F0: 693}
                _ = v53
            }
            v4 = v4
            _ = len(v2)
            switch "\n\t" {
            case "foo bar":
                _ = len(v2)
            }
            _ = fi52
        }
        if true {
            if ((true && (4294967295 <= 1)) || (!false)) {
                if false {
                    var v54 bool = false
                    v55 := GenS5{F0: "hello", F1: (v4 >> ((v3) & 63)), F2: v3}
                    for fi56 := 0; fi56 < 2; fi56++ {
                        v2 = append(v2, (fi56 - fi56))
                        if v54 {
                            v54 = v54
                            var v57 int8 = 0
                            v57 = -128
                            var v58 int32 = -2147483648
                            _ = v57
                            _ = v58
                        } else {
                            var v59 int32 = -407
                            var v60 float64 = 0.1
                            _ = v59
                            _ = v60
                        }
                        fmt.Println(v3, v4)
                        _ = fi56
                    }
                    _ = v54
                    _ = v55
                } else {
                    var v61 float64 = 1.5
                    _ = v61
                }
                var v62 float32 = (float32(0.1) / 1.0)
                var v63 bool = (true && ((true || (((127 == -1) || true) || true)) && false))
                _ = v62
                _ = v63
            }
        }
        var v64 bool = false
        _ = fi51
        _ = v64
    }
    for fi65 := 0; fi65 < 9; fi65++ {
        v66 := GenS6{F0: 61, F1: 4294967295, F2: genFn0()}
        v1 += 32767
        var v67 string = ""
        _ = fi65
        _ = v66
        _ = v67
    }
    if (false && ("foo bar" == "hello")) {
        v1 = (v1 | (v4 << (((v3 & (^(v3)))) & 63)))
    } else {
        var v68 uint8 = 128
        switch "\n\t" {
        case "hello":
            switch "\n\t" {
            case "foo bar":
                if (346 >= 42) {
                    var v69 uint16 = 65535
                    _ = v69
                }
            case "hello":
                var v70 int16 = -1
                _ = v70
            default:
                var v71 uint32 = 42
                _ = v71
            }
        }
        if len(v2) > 0 {
            v2[0] = -2147483648
        }
        _ = v68
    }
    switch 287 {
    case 42:
        if (-85 < 0) {
            delete(v5, 128)
            v3 += (v3 % (((v3 * v3)) | 1))
            v1 = v1
        } else {
            v72 := GenS4{F0: -38, F1: "\n\t"}
            v5[128] = (false || true)
            _ = v72
        }
    case -1:
        v73 := GenS7{F0: (2 >= 763), F1: v3, F2: 0}
        _ = v73
    default:
        v4 = v1
    }
    var v74 int16 = v1
    v74 -= (v1 >> (((v3 ^ (v3 / ((65535) | 1)))) & 63))
    var v75 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
    _ = v75
}
