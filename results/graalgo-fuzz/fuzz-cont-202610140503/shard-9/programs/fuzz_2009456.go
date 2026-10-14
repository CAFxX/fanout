package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 string; F2 int64 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]int)
    var v3 uint32 = 4294967295
    var v4 uint16 = 255
    v5 := make(map[int]int)
    var v6 uint16 = (v4 % ((v4) | 1))
    v7 := []int{980}
    for fi8 := 0; fi8 < 19; fi8++ {
        v1 += (-2.25 - 1e10)
        v4 = v6
        _ = fi8
    }
    v1 = -2.25
    for fi9 := 0; fi9 < 17; fi9++ {
        if true {
            var v10 bool = false
            v3 -= 2147483647
            var v11 bool = v10
            _ = v10
            _ = v11
        }
        v1 = ((1e10 + v1) * v1)
        if true {
            v7 = append(v7, 649)
            var v12 int8 = 0
            v13 := GenS1{F0: int16(v12)}
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    if false {
        if false {
            v14 := GenS1{F0: 42}
            switch "a" {
            case "":
                if false {
                    for fi15 := 0; fi15 < 1; fi15++ {
                        if (!false) {
                            if (32767 != 32767) {
                                var v16 float64 = (0.1 / 3.14159)
                                fmt.Println(v3, fi15)
                                _ = v16
                            }
                            v1 = v1
                        }
                        for fi17 := 0; fi17 < 16; fi17++ {
                            v18 := GenS1{F0: -32768}
                            _ = fi17
                            _ = v18
                        }
                        _ = fi15
                    }
                    if false {
                        v1 += 3.14159
                    }
                } else {
                    fmt.Println(v1)
                }
            case "hello":
                v2[42] = -72
            default:
                var v19 float32 = v1
                _ = v19
            }
            _ = v14
        } else {
            v20 := GenS1{F0: -1}
            switch "a" {
            case "hello":
                fmt.Println(v6)
            default:
                if (false && false) {
                    for fi21 := 0; fi21 < 18; fi21++ {
                        if (603 > 324) {
                            var v22 uint32 = (v3 & (v3 / (((^(v3))) | 1)))
                            if (!false) {
                                if ((70 != 128) || false) {
                                    var v23 string = ""
                                    if (fi21 >= fi21) {
                                        v7 = append(v7, fi21)
                                        _ = len(v7)
                                    } else {
                                        v22 = (v3 % ((19) | 1))
                                        v3 -= v22
                                    }
                                    delete(v2, 1)
                                    _ = v23
                                }
                                v4 = v4
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    var v25 uint32 = v22
                                    _ = fi24
                                    _ = v25
                                }
                            } else {
                                v22 -= 45
                            }
                            for fi26 := 0; fi26 < 7; fi26++ {
                                var v27 uint8 = 2
                                var v28 uint = 42
                                v4 -= v6
                                v29 := GenS2{F0: -1}
                                _ = fi26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            _ = v22
                        }
                        _ = fi21
                    }
                    for fi30 := 0; fi30 < 12; fi30++ {
                        if (2 < 232) {
                            if ("a" <= "foo bar") {
                                v3 += v3
                                var v31 int = int(v6)
                                v4 = (v4 << ((0) & 63))
                                _ = v31
                            } else {
                                _ = len(v7)
                            }
                            v3 -= v3
                        } else {
                            v32 := GenS3{F0: -162, F1: "foo bar", F2: 457}
                            _ = v32
                        }
                        _ = fi30
                    }
                }
            }
            switch "hello" {
            case "\n\t":
                var v33 int32 = 2147483647
                _ = v33
            case "":
                switch -324 {
                case 1:
                    fmt.Println(v3, v6, v6)
                case 2147483647:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                }
            default:
                for fi34 := 0; fi34 < 14; fi34++ {
                    for fi35 := 0; fi35 < 14; fi35++ {
                        var v36 int16 = -676
                        switch 4294967295 {
                        case -261:
                            var v37 uint64 = 4294967295
                            _ = v37
                        case -1:
                            switch (fi35 ^ int(v4)) {
                            case -625:
                                v38 := v2[v3]
                                _ = v38
                            default:
                                v3 -= v3
                            }
                        default:
                            var v39 int = 0
                            _ = v39
                        }
                        _ = fi35
                        _ = v36
                    }
                    _ = fi34
                }
            }
            _ = v20
        }
        for fi40 := 0; fi40 < 1; fi40++ {
            var v41 int = fi40
            for fi42 := 0; fi42 < 12; fi42++ {
                if ((fi40 < fi42) && false) {
                    var v43 int16 = 0
                    var v44 float64 = 0.1
                    v6 = v6
                    _ = v43
                    _ = v44
                } else {
                    var v45 uint64 = 42
                    _ = v45
                }
                _ = fi42
            }
            _ = fi40
            _ = v41
        }
        switch int(v4) {
        case 4294967295:
            v3 -= 465
        case 1:
            var v46 uint = 429
            _ = v46
        case -2147483648:
            var v47 int16 = 1
            _ = v47
        default:
            var v48 int64 = -1
            _ = v48
        }
    }
    var v49 int64 = 9223372036854775807
    v50 := GenS1{F0: 42}
    var v51 float64 = -2.25
    switch "foo bar" {
    case "foo bar":
        _ = len(v7)
    }
    for fi52 := 0; fi52 < 2; fi52++ {
        var v53 float32 = 0.1
        var v54 uint = 790
        v7 = append(v7, fi52)
        _ = fi52
        _ = v53
        _ = v54
    }
    switch "a" {
    case "foo bar":
        if (2 >= 42) {
            var v55 uint8 = 127
            if false {
                switch "foo bar" {
                case "":
                    for fi56 := 0; fi56 < 13; fi56++ {
                        if true {
                            v57 := GenS1{F0: int16(fi56)}
                            var v58 int64 = v49
                            _ = len(v7)
                            _ = v57
                            _ = v58
                        }
                        if (fi56 != (fi56 ^ (-(fi56)))) {
                            for fi59 := 0; fi59 < 19; fi59++ {
                                v60 := GenS1{F0: 1}
                                fmt.Println(v4)
                                var v61 uint = 4294967295
                                _ = fi59
                                _ = v60
                                _ = v61
                            }
                            v62 := GenS1{F0: 32767}
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = v62
                        }
                        if true {
                            v6 += 220
                        } else {
                            v3 = v3
                            v3 -= 2
                        }
                        _ = fi56
                    }
                case "foo bar":
                    if (!true) {
                        var v63 uint64 = 0
                        _ = v63
                    }
                }
                var v64 bool = (false && (-266 <= 32767))
                _ = v64
            }
            switch "a" {
            case "a":
                var v65 int64 = 1
                _ = v65
            case "\n\t":
                var v66 int32 = 2147483647
                _ = v66
            case "foo bar":
                v4 += v6
            default:
                if false {
                    v4 = (^(v6))
                    for fi67 := 0; fi67 < 19; fi67++ {
                        var v68 float32 = (v1 - 1e10)
                        _ = fi67
                        _ = v68
                    }
                } else {
                    for fi69 := 0; fi69 < 14; fi69++ {
                        if (true || (146 >= 18446744073709551615)) {
                            v70 := GenS1{F0: -32768}
                            _ = v70
                        } else {
                            fmt.Println(v4, v55, v55)
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            v4 -= 32767
                        }
                        var v71 int64 = 1
                        v72 := GenS2{F0: -49}
                        _ = fi69
                        _ = v71
                        _ = v72
                    }
                    for fi73 := 0; fi73 < 16; fi73++ {
                        v4 += uint16(v55)
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        _ = fi73
                    }
                }
            }
            _ = v55
        }
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi74 := 0; fi74 < 10; fi74++ {
        var v75 int8 = 1
        if (440 != int16(v49)) {
            v76 := GenS3{F0: 341, F1: "a", F2: 9223372036854775807}
            var v77 uint32 = v3
            _ = v76
            _ = v77
        }
        _ = fi74
        _ = v75
    }
    for fi78 := 0; fi78 < 12; fi78++ {
        v49 = (-(v49))
        fmt.Println(v1, v3, v3)
        var v79 int8 = 0
        _ = fi78
        _ = v79
    }
    var v80 uint32 = v3
    v7 = append(v7, 648)
    if (v1 != (float32(v51) / (-2.25 + (v1 / (0.0 * v1))))) {
        v2[(v80 >> ((617) & 63))] = 371
        v4 = (^(v4))
    } else {
        for fi81 := 0; fi81 < 20; fi81++ {
            fmt.Println(v80, v3, v49)
            _ = fi81
        }
        v4 -= v4
        v51 += (float64(float32(3.14159)) * 1e10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
    _ = v80
}
