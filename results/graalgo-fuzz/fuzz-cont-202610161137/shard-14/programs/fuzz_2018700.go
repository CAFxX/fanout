package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 string; F1 uint64; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 uint8; F2 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 bool }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{950, 128, 42, 127}
    var v3 int16 = 42
    v4 := make(map[int]int)
    var v5 int8 = 100
    if (0.0 < 0.0) {
        v3 += v3
        fmt.Println(v3, v3)
    } else {
        var v6 int = 760
        _ = v6
    }
    if true {
        fmt.Println(v5)
        var v7 float64 = 1e10
        if true {
            v8 := v4[-703]
            var v9 bool = (!false)
            var v10 uint8 = 0
            _ = v10
            _ = v8
            _ = v9
        } else {
            for fi11 := 0; fi11 < 4; fi11++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v7 = float64(-2.25)
                delete(v1, 2147483648)
                _ = fi11
            }
            v4[0] = 128
        }
        _ = v7
    } else {
        if (!false) {
            for fi12 := 0; fi12 < 6; fi12++ {
                delete(v1, 0)
                var v13 uint8 = 128
                switch int(v5) {
                case -525:
                    for fi14 := 0; fi14 < 8; fi14++ {
                        if true {
                            v15 := GenS1{F0: v3, F1: -769}
                            _ = v15
                        }
                        fmt.Println(v5)
                        fmt.Println(v5, fi12, fi14)
                        _ = fi14
                    }
                case 127:
                    v16 := GenS2{F0: "a", F1: 2, F2: 2147483647}
                    _ = v16
                case 128:
                    if len(v2) > 0 {
                        v2[0] = (fi12 + (fi12 & fi12))
                    }
                }
                _ = fi12
                _ = v13
            }
            v17 := GenS3{F0: "foo bar"}
            switch 127 {
            case 0:
                var v18 int16 = (v3 * int16(v5))
                _ = v18
            case -336:
                for fi19 := 0; fi19 < 1; fi19++ {
                    fmt.Println(fi19)
                    if true {
                        var v20 uint = 450
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v20
                    }
                    var v21 bool = (0 <= 0)
                    _ = fi19
                    _ = v21
                }
            case 32767:
                v5 = v5
            default:
                var v22 uint8 = 216
                _ = v22
            }
            _ = v17
        } else {
            var v23 int64 = -1
            v3 = v3
            v3 -= (v3 ^ int16(v5))
            _ = v23
        }
        var v24 float64 = 3.14159
        v5 = 85
        _ = v24
    }
    var v25 uint8 = 255
    if ((v3 + (v3 - (v3 / ((42) | 1)))) < 0) {
        v5 -= v5
        var v26 int8 = (v5 | -4)
        switch "\n\t" {
        case "a":
            delete(v4, 32767)
        default:
            var v27 uint32 = 447
            _ = v27
        }
        _ = v26
    }
    var v28 float64 = ((0.0 / 1.5) * -2.25)
    if ((true && true) || (!(!(!(!(42 > 186)))))) {
        v28 = (v28 + v28)
    }
    var v29 int64 = 9223372036854775807
    var v30 uint32 = 2147483648
    var v31 int16 = (v3 >> ((4294967295) & 63))
    switch -685 {
    case 1:
        var v32 uint16 = 1
        _ = v32
    case 199:
        v33 := GenS4{F0: (v31 >> ((1) & 63)), F1: v29}
        _ = v33
    default:
        var v34 int8 = (v5 - (^(v5)))
        _ = v34
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        v36 := GenS1{F0: v31, F1: -35}
        _ = fi35
        _ = v36
    }
    switch -879 {
    case 32767:
        var v37 int = 128
        _ = v37
    case -405:
        switch "a" {
        case "foo bar":
            var v38 string = "foo bar"
            _ = v38
        case "a":
            var v39 float32 = 1.5
            _ = v39
        default:
            for fi40 := 0; fi40 < 18; fi40++ {
                v5 = -128
                _ = fi40
            }
        }
    }
    v41 := GenS5{F0: (v25 ^ v25)}
    v29 = 0
    for fi42 := 0; fi42 < 5; fi42++ {
        var v43 bool = (128 >= fi42)
        if (float64(3.14159) <= v28) {
            v44 := GenS5{F0: 84}
            _ = v44
        }
        for fi45 := 0; fi45 < 13; fi45++ {
            v43 = (!(true || (-32768 < -878)))
            v5 += v5
            var v46 float32 = -2.25
            _ = fi45
            _ = v46
        }
        _ = fi42
        _ = v43
    }
    switch -96 {
    case 32768:
        var v47 uint64 = 4294967295
        _ = v47
    case -919:
        if len(v2) > 0 {
            v2[0] = 32768
        }
    case 127:
        switch "" {
        case "":
            v3 = int16(v30)
        default:
            var v48 uint16 = 255
            _ = v48
        }
    default:
        if (1 == 1) {
            v49 := GenS6{F0: (!true), F1: (v25 + v25), F2: (v5 << ((968) & 63))}
            if (-32768 <= 32767) {
                v50 := GenS7{F0: (v31 >> ((4294967295) & 63))}
                _ = v50
            }
            switch -196 {
            case 2147483647:
                if true {
                    if false {
                        var v51 bool = (-625 == 2147483647)
                        v52 := GenS8{F0: true}
                        v4[238] = -2147483648
                        _ = v51
                        _ = v52
                    }
                    var v53 float64 = (float64(float32(-2.25)) / v28)
                    var v54 float64 = ((((0.1 * ((v53 / (v53 + v53)) + v28)) + 1.5) - v28) * v53)
                    _ = v53
                    _ = v54
                } else {
                    v2 = append(v2, 32768)
                }
            case 32768:
                switch (("hello" + "foo bar") + "a") {
                case "foo bar":
                    for fi55 := 0; fi55 < 8; fi55++ {
                        if ((false && (1 == 0)) || false) {
                            v2 = append(v2, (^(fi55)))
                        } else {
                            if false {
                                if true {
                                    v5 -= v5
                                    v3 = 42
                                }
                                var v56 int16 = v31
                                _ = v56
                            } else {
                                v28 += 3.14159
                                v5 -= int8(v25)
                                var v57 bool = (!(fi55 > fi55))
                                _ = v57
                            }
                        }
                        _ = fi55
                    }
                case "\n\t":
                    if (1.5 < (((0.1 - float32(v28)) * (0.0 * -2.25)) / 1e10)) {
                        _ = len(v1)
                        v28 -= float64((((0.0 / ((-2.25 / 3.14159) + float32(v28))) / (0.1 - (0.0 - float32((((v28 - 1e10) - v28) * (float64(1e10) / -2.25)))))) * -2.25))
                        var v58 uint = 255
                        _ = v58
                    }
                }
            case 127:
                fmt.Println(v31)
            }
            _ = v49
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v41
    _ = v5
}
