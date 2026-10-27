package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint; F2 uint64 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 int; F1 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{}
    var v3 int16 = 448
    v4 := make(map[uint32]string)
    v5 := make(map[int]uint32)
    v5[-2147483648] = 944
    fmt.Println(v3)
    for fi6 := 0; fi6 < 19; fi6++ {
        switch (fi6 * fi6) {
        case 32767:
            switch fi6 {
            case 128:
                var v7 int = fi6
                _ = v7
            }
        case 2147483647:
            if false {
                v8 := GenS1{F0: 2147483647, F1: 1}
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := GenS2{F0: ((65535 != 0) || false), F1: true}
                    _ = fi9
                    _ = v10
                }
                v3 += (v3 % ((v3) | 1))
                _ = v8
            }
        }
        var v11 float32 = (((float32((0.1 / 1.5)) - 0.1) * (3.14159 - (0.1 + 0.1))) * 3.14159)
        _ = fi6
        _ = v11
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        _ = len(v2)
        switch "hello" {
        case "foo bar":
            v13 := v5[(fi12 ^ fi12)]
            _ = v13
        case "hello":
            switch (^(fi12)) {
            case 4294967295:
                v3 += v3
            case 127:
                v14 := GenS3{F0: -128}
                _ = v14
            case 304:
                v3 = (v3 & -1)
            }
        case "":
            v3 = -991
        }
        _ = fi12
    }
    switch 127 {
    case -2147483648:
        fmt.Println(v3)
    case 32768:
        switch "a" {
        case "foo bar":
            v15 := GenS4{F0: 42, F1: 614, F2: 998}
            _ = v15
        case "hello":
            if true {
                var v16 bool = (!(0 == 65535))
                if (9223372036854775807 > 1) {
                    v3 -= v3
                    var v17 int64 = 0
                    v3 += v3
                    _ = v17
                }
                _ = v16
            }
        default:
            switch -230 {
            case 1:
                var v18 int32 = 0
                _ = v18
            case -902:
                _ = len(v2)
            case 4294967295:
                v3 = (v3 % ((1) | 1))
            }
        }
    case -75:
        for fi19 := 0; fi19 < 20; fi19++ {
            v3 += v3
            for fi20 := 0; fi20 < 12; fi20++ {
                v3 -= v3
                _ = fi20
            }
            _ = fi19
        }
    }
    _ = len(v2)
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 int = fi21
        var v23 uint64 = 2
        _ = fi21
        _ = v22
        _ = v23
    }
    if ((false && true) && ((!(19 < 2)) || true)) {
        var v24 uint = 4294967295
        for fi25 := 0; fi25 < 16; fi25++ {
            var v26 int16 = (v3 & 139)
            v27 := GenS2{F0: (3.14159 != -2.25), F1: true}
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = v24
    }
    v28 := GenS1{F0: 32767, F1: 4294967295}
    if (340 > 65535) {
        v3 += 1
        if ("\n\t" >= ("a" + ("foo bar" + "foo bar"))) {
            for fi29 := 0; fi29 < 16; fi29++ {
                fmt.Println(v3, fi29)
                v30 := GenS2{F0: (100 <= -1), F1: false}
                if (1 == 100) {
                    var v31 uint64 = 637
                    var v32 float64 = -2.25
                    _ = v31
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
            switch ("foo bar" + ("hello" + "foo bar")) {
            case "a":
                v3 = v3
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            if true {
                for fi33 := 0; fi33 < 7; fi33++ {
                    if ((v3 ^ (v3 * -95)) >= (v3 ^ 0)) {
                        for fi34 := 0; fi34 < 12; fi34++ {
                            var v35 uint = 0
                            _ = fi34
                            _ = v35
                        }
                        v3 = 1000
                    }
                    var v36 int64 = -9223372036854775808
                    v37 := v5[(fi33 ^ 302)]
                    _ = fi33
                    _ = v36
                    _ = v37
                }
                for fi38 := 0; fi38 < 8; fi38++ {
                    v39 := v4[0]
                    switch (fi38 - 760) {
                    case 669:
                        fmt.Println(v39)
                    case 42:
                        for fi40 := 0; fi40 < 5; fi40++ {
                            var v41 float64 = float64((3.14159 * (0.0 / 1e10)))
                            fmt.Println(v41)
                            fmt.Println(fi40, v3, v3)
                            _ = fi40
                            _ = v41
                        }
                    case 0:
                        switch v39 {
                        case "a":
                            if (!true) {
                                if false {
                                    if true {
                                        var v42 float64 = 0.1
                                        if true {
                                            v39 = "hello"
                                        }
                                        if (-92 >= 0) {
                                            v39 = "foo bar"
                                        } else {
                                            if true {
                                                v5[-1] = 0
                                                v3 += (v3 | v3)
                                            } else {
                                                if true {
                                                    fmt.Println(v39)
                                                    switch v39 {
                                                    case "hello":
                                                        fmt.Println(v39, fi38, v39)
                                                    case "":
                                                        v43 := GenS1{F0: 2147483647, F1: 2147483647}
                                                        _ = v43
                                                    }
                                                    switch (fi38 / ((fi38) | 1)) {
                                                    case 2147483647:
                                                        var v44 int = (fi38 / ((-514) | 1))
                                                        _ = v44
                                                    case 32767:
                                                        var v45 uint64 = 514
                                                        _ = v45
                                                    default:
                                                        v42 = v42
                                                    }
                                                } else {
                                                    if true {
                                                        v42 = v42
                                                        for fi46 := 0; fi46 < 15; fi46++ {
                                                            v47 := GenS3{F0: 0}
                                                            var v48 bool = (0 <= 242)
                                                            _ = fi46
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        var v49 string = "a"
                                                        _ = v49
                                                    }
                                                    v3 = 1
                                                }
                                            }
                                            v3 += -1
                                            var v50 float64 = v42
                                            _ = v50
                                        }
                                        _ = v42
                                    }
                                    if ((0 == 0) || (1.5 < 0.1)) {
                                        v51 := GenS5{F0: 65535, F1: (v3 % ((888) | 1))}
                                        switch (fi38 / ((fi38) | 1)) {
                                        case 597:
                                            switch "hello" {
                                            case "a":
                                                if len(v2) > 0 {
                                                    v2[0] = 0
                                                }
                                            case "hello":
                                                var v52 int16 = v3
                                                _ = v52
                                            case "":
                                                switch fi38 {
                                                case -2147483648:
                                                    for fi53 := 0; fi53 < 9; fi53++ {
                                                        v3 = v3
                                                        var v54 int32 = 1
                                                        _ = fi53
                                                        _ = v54
                                                    }
                                                case 4294967295:
                                                    v3 -= (v3 >> ((773) & 63))
                                                }
                                            default:
                                                for fi55 := 0; fi55 < 17; fi55++ {
                                                    v3 -= (^(v3))
                                                    var v56 uint16 = 1
                                                    v57 := GenS4{F0: fi38, F1: 42, F2: 971}
                                                    _ = fi55
                                                    _ = v56
                                                    _ = v57
                                                }
                                            }
                                        case 300:
                                            v3 = v3
                                        case 0:
                                            var v58 uint8 = 127
                                            _ = v58
                                        default:
                                            v3 -= (v3 % ((-32768) | 1))
                                        }
                                        _ = v51
                                    }
                                } else {
                                    for fi59 := 0; fi59 < 4; fi59++ {
                                        fmt.Println(v3, fi59, fi59)
                                        fmt.Println(v39, v39)
                                        var v60 bool = false
                                        fmt.Println(v39, v39)
                                        var v61 int64 = 1
                                        fmt.Println(v39)
                                        _ = fi59
                                        _ = v60
                                        _ = v61
                                    }
                                    for fi62 := 0; fi62 < 9; fi62++ {
                                        if ((2 != 91) && (9223372036854775807 == -9223372036854775808)) {
                                            switch fi38 {
                                            case -373:
                                                v39 = v39
                                            }
                                            fmt.Println(fi38)
                                        }
                                        var v63 string = v39
                                        _ = len(v1)
                                        _ = fi62
                                        _ = v63
                                    }
                                }
                                var v64 int = (fi38 + 4294967295)
                                for fi65 := 0; fi65 < 11; fi65++ {
                                    v5[(v64 * fi38)] = 340
                                    v39 = ((v39 + (v39 + "\n\t")) + "\n\t")
                                    _ = fi65
                                }
                                _ = v64
                            } else {
                                v2 = append(v2, 240)
                                if true {
                                    var v66 bool = (42 < 42)
                                    v67 := GenS4{F0: -42, F1: 42, F2: 4294967295}
                                    _ = v66
                                    _ = v67
                                }
                            }
                        case "hello":
                            v3 -= (v3 % ((-1) | 1))
                        }
                    default:
                        for fi68 := 0; fi68 < 3; fi68++ {
                            if false {
                                _ = len(v5)
                            } else {
                                var v69 int32 = -1
                                delete(v4, 42)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v69
                            }
                            v70 := GenS6{F0: fi38, F1: 0}
                            _ = fi68
                            _ = v70
                        }
                    }
                    _ = fi38
                    _ = v39
                }
            } else {
                if false {
                    if false {
                        switch "" {
                        case "hello":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "":
                            v3 -= v3
                        case "\n\t":
                            v3 = (v3 & 1000)
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    v3 = v3
                    v3 = (v3 % (((v3 * v3)) | 1))
                }
                if len(v2) > 0 {
                    v2[0] = 65535
                }
            }
        }
        delete(v1, ("foo bar" + "foo bar"))
    }
    var v71 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v71
}
