package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
type GenS2 struct { F0 int64; F1 int; F2 bool }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int64; F1 bool; F2 int64 }

func main() {
    var v1 uint64 = 308
    var v2 int64 = -234
    var v3 int32 = 477
    var v4 int64 = v2
    var v5 uint16 = 32767
    var v6 float64 = (float64(-2.25) - 1.5)
    if true {
        if false {
            fmt.Println(v3, v2)
            if true {
                v7 := []int{404, -1, -2147483648, 128}
                for fi8 := 0; fi8 < 10; fi8++ {
                    if true {
                        if false {
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            v9 := GenS1{F0: uint16(v1), F1: -2147483648, F2: (-(v3))}
                            _ = v9
                        } else {
                            if true {
                                var v10 uint16 = 65535
                                _ = v10
                            } else {
                                switch "" {
                                case "":
                                    var v11 int = fi8
                                    _ = v11
                                case "foo bar":
                                    var v12 float32 = 0.0
                                    _ = v12
                                case "hello":
                                    if len(v7) > 0 {
                                        v7[0] = fi8
                                    }
                                }
                                for fi13 := 0; fi13 < 17; fi13++ {
                                    var v14 uint64 = v1
                                    fmt.Println(v4, v3, v5)
                                    if len(v7) > 0 {
                                        v7[0] = fi8
                                    }
                                    _ = fi13
                                    _ = v14
                                }
                            }
                        }
                        var v15 int16 = 1000
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 14; fi16++ {
                            if (65535 > 65535) {
                                v17 := GenS1{F0: v5, F1: 268, F2: v3}
                                fmt.Println(fi8, fi8)
                                v6 -= v6
                                _ = v17
                            } else {
                                v18 := GenS1{F0: (v5 * v5), F1: v3, F2: (v3 << ((2) & 63))}
                                var v19 int8 = -29
                                _ = v18
                                _ = v19
                            }
                            fmt.Println(v4, v4)
                            if len(v7) > 0 {
                                v7[0] = fi8
                            }
                            _ = fi16
                        }
                        switch "foo bar" {
                        case "foo bar":
                            v20 := GenS2{F0: v2, F1: (fi8 << ((144) & 63)), F2: (("hello" + "") <= "hello")}
                            _ = v20
                        case "\n\t":
                            v21 := GenS3{F0: true, F1: 1}
                            _ = v21
                        default:
                            var v22 bool = false
                            _ = v22
                        }
                    }
                    _ = fi8
                }
                var v23 string = "hello"
                _ = v23
                _ = v7
            }
            switch 32767 {
            case 283:
                v24 := GenS1{F0: (^(v5)), F1: v3, F2: (v3 ^ (-(v3)))}
                _ = v24
            case 0:
                for fi25 := 0; fi25 < 14; fi25++ {
                    if true {
                        v26 := []int{(fi25 >> ((42) & 63)), -625, 207}
                        _ = v26
                    }
                    if false {
                        var v27 uint8 = 1
                        if true {
                            var v28 uint = 18446744073709551615
                            _ = v28
                        } else {
                            v6 -= 1e10
                            var v29 uint8 = v27
                            _ = v29
                        }
                        _ = v27
                    }
                    _ = fi25
                }
            case -828:
                for fi30 := 0; fi30 < 18; fi30++ {
                    for fi31 := 0; fi31 < 2; fi31++ {
                        switch "a" {
                        case "hello":
                            if (false || true) {
                                v32 := GenS4{F0: (fi31 == 32767), F1: 2}
                                _ = v32
                            }
                        case "":
                            v33 := []int{}
                            _ = v33
                        case "a":
                            v2 = 970
                        default:
                            v6 = (v6 / (3.14159 - v6))
                        }
                        v4 -= 42
                        fmt.Println(v5, fi31, v1)
                        if (!true) {
                            var v34 bool = ((255 > 42) || true)
                            if false {
                                fmt.Println(v6)
                            }
                            v1 -= v1
                            v35 := []int{(fi31 ^ fi31), fi30, fi30, (fi31 >> ((169) & 63))}
                            if len(v35) >= 2 {
                                v35 = v35[1:len(v35)]
                            }
                            v1 = (^(v1))
                            _ = v34
                            _ = v35
                        } else {
                            switch "" {
                            case "a":
                                var v36 float32 = ((0.1 - 1.5) - 1.5)
                                _ = v36
                            case "foo bar":
                                v37 := GenS1{F0: (v5 - v5), F1: 376, F2: v3}
                                _ = v37
                            }
                            fmt.Println(fi31, v4)
                            v5 = v5
                        }
                        _ = fi31
                    }
                    v38 := []int{0, (-(fi30)), fi30}
                    _ = fi30
                    _ = v38
                }
            }
        } else {
            var v39 uint = 255
            v40 := GenS4{F0: ((-1 >= int8(v5)) && (!false)), F1: 2}
            var v41 int64 = -9223372036854775808
            _ = v39
            _ = v40
            _ = v41
        }
        for fi42 := 0; fi42 < 11; fi42++ {
            for fi43 := 0; fi43 < 4; fi43++ {
                var v44 uint64 = (^(v1))
                switch "hello" {
                case "a":
                    v45 := GenS5{F0: 376, F1: fi42}
                    _ = v45
                case "hello":
                    v5 = 65535
                }
                _ = fi43
                _ = v44
            }
            v4 = 1
            var v46 float64 = float64((0.0 * 1e10))
            _ = fi42
            _ = v46
        }
    } else {
        fmt.Println(v4, v5, v1)
    }
    switch (("" + ("hello" + "\n\t")) + (("foo bar" + "") + "")) {
    case "foo bar":
        var v47 int16 = -840
        _ = v47
    case "a":
        v48 := []int{}
        _ = v48
    case "":
        var v49 int64 = (v2 % ((v4) | 1))
        _ = v49
    }
    v5 -= v5
    if (4294967295 != 65535) {
        if (false && false) {
            var v50 uint = 58
            for fi51 := 0; fi51 < 14; fi51++ {
                for fi52 := 0; fi52 < 5; fi52++ {
                    if ((v50 / (((^(v50))) | 1)) <= (^(v50))) {
                        fmt.Println(v50, v3)
                    }
                    _ = fi52
                }
                v1 -= 0
                _ = fi51
            }
            _ = v50
        }
    } else {
        v3 = v3
    }
    var v53 int16 = -825
    if ((-94 != (v3 * v3)) || ((v6 + v6) <= v6)) {
        v54 := []int{467, 512, 292}
        fmt.Println(v53)
        switch "" {
        case "":
            for fi55 := 0; fi55 < 9; fi55++ {
                for fi56 := 0; fi56 < 9; fi56++ {
                    fmt.Println(fi55)
                    v57 := GenS5{F0: 958, F1: (fi55 << ((42) & 63))}
                    fmt.Println(v1)
                    _ = fi56
                    _ = v57
                }
                var v58 int8 = -1
                v59 := GenS6{F0: (v2 % ((v4) | 1)), F1: ((-2.25 >= -2.25) && ((((v58 >> ((645) & 63)) <= (v58 % ((-128) | 1))) || ((132 == 42) || ((3.14159 >= (-2.25 * (float32(1.5) + 0.0))) || (true && true)))) && (!false))), F2: v4}
                _ = fi55
                _ = v58
                _ = v59
            }
        case "\n\t":
            v1 -= (v1 | (v1 & v1))
        }
        _ = v54
    } else {
        v1 = v1
    }
    var v60 uint64 = v1
    var v61 int16 = (v53 << ((97) & 63))
    for fi62 := 0; fi62 < 2; fi62++ {
        fmt.Println(v53, v53, v53)
        v1 = v60
        switch (fi62 ^ fi62) {
        case 777:
            for fi63 := 0; fi63 < 18; fi63++ {
                var v64 float32 = 0.1
                _ = fi63
                _ = v64
            }
        }
        _ = fi62
    }
    v3 = (v3 << ((461) & 63))
    for fi65 := 0; fi65 < 8; fi65++ {
        v1 -= 0
        v60 -= (v60 >> ((727) & 63))
        v53 = -32768
        _ = fi65
    }
    for fi66 := 0; fi66 < 8; fi66++ {
        v60 -= v60
        _ = fi66
    }
    v67 := []int{-1, int(v2)}
    switch -941 {
    case 87:
        v60 += v60
    case 1:
        switch "foo bar" {
        case "a":
            _ = len(v67)
        case "hello":
            var v68 uint16 = v5
            _ = v68
        }
    case 32767:
        if len(v67) >= 2 {
            v67 = v67[1:len(v67)]
        }
    default:
        switch ("hello" + "a") {
        case "a":
            var v69 string = ("a" + (((("\n\t" + "foo bar") + "hello") + "\n\t") + "hello"))
            _ = v69
        case "\n\t":
            for fi70 := 0; fi70 < 2; fi70++ {
                v67 = append(v67, fi70)
                var v71 uint16 = 599
                v4 = v4
                _ = fi70
                _ = v71
            }
        case "":
            switch "\n\t" {
            case "hello":
                fmt.Println(v4)
            default:
                for fi72 := 0; fi72 < 14; fi72++ {
                    fmt.Println(v6, v1, v1)
                    _ = fi72
                }
            }
        default:
            var v73 int32 = (-(v3))
            _ = v73
        }
    }
    for fi74 := 0; fi74 < 20; fi74++ {
        if len(v67) > 0 {
            v67[0] = fi74
        }
        _ = fi74
    }
    v2 -= (v2 & (-(v2)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v67:", len(v67), v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v60
    _ = v61
    _ = v67
}
