package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint }
type GenS4 struct { F0 int32; F1 uint32; F2 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 uint8; F2 int8 }
type GenS7 struct { F0 int16; F1 uint32 }

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 int64 = -54
    v5 := make(map[int]bool)
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 14; fi7++ {
        switch (fi7 % (((fi7 ^ fi7)) | 1)) {
        case 32768:
            switch ("a" + ("foo bar" + "a")) {
            case "a":
                switch "\n\t" {
                case "\n\t":
                    v8 := []int{fi7, (fi7 >> ((18446744073709551615) & 63))}
                    _ = v8
                case "a":
                    fmt.Println(v4)
                default:
                    fmt.Println(v1)
                }
            default:
                v4 = int64(v1)
            }
        case -9223372036854775808:
            fmt.Println(v4)
        default:
            v9 := GenS1{F0: int16(fi7), F1: (fi7 & (fi7 + (fi7 / (((fi7 * fi7)) | 1)))), F2: 1}
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        if (!false) {
            v11 := GenS2{F0: 0, F1: 0}
            var v12 uint8 = 128
            if false {
                v13 := GenS1{F0: -307, F1: (^(fi10)), F2: 4294967295}
                if (v12 == 148) {
                    for fi14 := 0; fi14 < 1; fi14++ {
                        if false {
                            var v15 bool = (false || true)
                            v15 = ("\n\t" > "hello")
                            _ = v15
                        } else {
                            var v16 int16 = 1
                            _ = v16
                        }
                        _ = fi14
                    }
                    fmt.Println(v4)
                } else {
                    v12 = v12
                    v17 := GenS2{F0: -767, F1: 921}
                    _ = v17
                }
                for fi18 := 0; fi18 < 20; fi18++ {
                    switch "a" {
                    case "a":
                        switch (^(fi10)) {
                        case -160:
                            if true {
                                v19 := v2[2147483648]
                                var v20 uint8 = v12
                                _ = v19
                                _ = v20
                            } else {
                                v4 += -1
                            }
                            switch (fi10 ^ (fi18 << ((6) & 63))) {
                            case -352:
                                fmt.Println(fi10)
                            case -2147483648:
                                var v21 uint32 = 2147483647
                                _ = v21
                            }
                        case 9223372036854775807:
                            if false {
                                v22 := GenS1{F0: -379, F1: fi18, F2: 2}
                                delete(v2, uint32(v12))
                                switch (fi10 % ((fi18) | 1)) {
                                case 32767:
                                    fmt.Println(v4, v4, v12)
                                case 323:
                                    var v23 uint8 = 2
                                    _ = v23
                                case -482:
                                    var v24 int16 = 0
                                    fmt.Println(v4, v24, fi10)
                                    _ = v24
                                default:
                                    if true {
                                        v4 = int64(v1)
                                        v1 -= v1
                                        v12 += v12
                                    }
                                }
                                _ = v22
                            }
                        case -250:
                            var v25 bool = (42 == 255)
                            _ = v25
                        default:
                            switch (fi10 & fi10) {
                            case -399:
                                var v26 int32 = 1
                                _ = v26
                            case 9223372036854775807:
                                fmt.Println(v1, v4, v1)
                            }
                        }
                    default:
                        var v27 bool = ((((float32((((((3.14159 + 0.0) + (3.14159 * (1.5 / float64((float32(0.1) / (0.1 / (1e10 / 1.0))))))) * 1e10) / 1e10) / 0.1)) - (0.1 - 3.14159)) * 0.1) / 3.14159) < 0.1)
                        _ = v27
                    }
                    var v28 string = ""
                    _ = fi18
                    _ = v28
                }
                _ = v13
            } else {
                var v29 uint64 = 4294967295
                fmt.Println(v12, fi10, v1)
                _ = v29
            }
            _ = v11
            _ = v12
        } else {
            var v30 string = ("a" + "hello")
            switch 4294967295 {
            case -2147483648:
                if (false && true) {
                    var v31 uint32 = 4294967295
                    fmt.Println(v4, v1)
                    _ = v31
                } else {
                    v2[42] = fi10
                    v32 := GenS3{F0: 128, F1: v1, F2: 59}
                    var v33 bool = ((-(fi10)) >= fi10)
                    _ = v32
                    _ = v33
                }
            default:
                v34 := GenS4{F0: -2147483648, F1: 2147483648, F2: 0}
                _ = v34
            }
            switch (fi10 ^ -986) {
            case 32768:
                switch (fi10 * 32768) {
                case 46:
                    if (0.1 == -2.25) {
                        var v35 int = fi10
                        _ = v35
                    } else {
                        var v36 uint8 = 255
                        _ = v36
                    }
                case 128:
                    var v37 float64 = (1.5 * 1e10)
                    _ = v37
                case 0:
                    switch "" {
                    case "\n\t":
                        v38 := v2[2147483648]
                        _ = v38
                    default:
                        var v39 string = v30
                        _ = v39
                    }
                }
            default:
                v1 += 113
            }
            _ = v30
        }
        v40 := []int{}
        _ = fi10
        _ = v40
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        if true {
            if (true || ((false || ((170 >= 32767) || false)) || false)) {
                v1 -= v1
                for fi42 := 0; fi42 < 16; fi42++ {
                    var v43 uint32 = 4294967295
                    v44 := GenS5{F0: (v1 ^ (-(v1)))}
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                switch (fi41 << ((42) & 63)) {
                case -9223372036854775808:
                    switch "" {
                    case "":
                        var v45 int8 = (v1 >> ((0) & 63))
                        _ = v45
                    }
                case 592:
                    if true {
                        var v46 uint32 = 2147483647
                        v47 := []int{}
                        _ = v46
                        _ = v47
                    }
                case -1:
                    var v48 int = fi41
                    _ = v48
                default:
                    var v49 uint64 = 250
                    _ = v49
                }
            } else {
                var v50 int32 = 1
                _ = v50
            }
            delete(v6, (fi41 % ((fi41) | 1)))
        }
        _ = fi41
    }
    fmt.Println(v1, v1)
    v51 := []int{127, -9223372036854775808, 127}
    for fi52 := 0; fi52 < 18; fi52++ {
        delete(v6, fi52)
        v53 := GenS6{F0: (((("foo bar" + ("hello" + ("foo bar" + "hello"))) + "hello") + "hello") + "hello"), F1: 128, F2: v1}
        _ = fi52
        _ = v53
    }
    var v54 int8 = 27
    var v55 string = ("\n\t" + ("\n\t" + "\n\t"))
    fmt.Println(v55)
    v56 := GenS7{F0: 0, F1: 2147483648}
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
