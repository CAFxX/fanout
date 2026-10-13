package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 bool; F1 uint32; F2 int32 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 bool }
func genFn0(p0 uint16, p1 int64, p2 int64) uint8 {
    v4 := []int{-1, -208, -9223372036854775808, 828}
    var v5 uint64 = 18
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[string]string)
    var v3 string = ("foo bar" + "")
    fmt.Println(v1, v3)
    for fi6 := 0; fi6 < 15; fi6++ {
        switch (v1 & (fi6 | (fi6 % ((fi6) | 1)))) {
        case 938:
            v1 = 32768
        default:
            for fi7 := 0; fi7 < 6; fi7++ {
                v1 = (fi7 ^ -2147483648)
                var v8 bool = false
                _ = fi7
                _ = v8
            }
        }
        v9 := GenS1{F0: 4294967295, F1: -128, F2: 42}
        _ = fi6
        _ = v9
    }
    var v10 float32 = (1.5 / -2.25)
    var v11 uint = 18446744073709551615
    v11 -= v11
    if false {
        var v12 int16 = 0
        switch (-(v1)) {
        case 32767:
            v3 = "a"
        case 127:
            switch "hello" {
            case "hello":
                fmt.Println(v11)
            case "a":
                v11 = (^(v11))
            default:
                var v13 int = (-(v1))
                _ = v13
            }
        case 184:
            fmt.Println(v1)
        default:
            var v14 uint = (^(v11))
            _ = v14
        }
        for fi15 := 0; fi15 < 13; fi15++ {
            v1 = (fi15 ^ v1)
            if false {
                if (672 < 1) {
                    v16 := GenS1{F0: 0, F1: -128, F2: 42}
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        if (false || (2 != 2147483648)) {
                            v18 := GenS2{F0: v11}
                            var v19 uint8 = 42
                            var v20 int32 = -2147483648
                            _ = v18
                            _ = v19
                            _ = v20
                        } else {
                            v10 = v10
                        }
                        v21 := GenS1{F0: 18446744073709551615, F1: 100, F2: 84}
                        v22 := GenS3{F0: 1, F1: -338}
                        _ = fi17
                        _ = v21
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    var v24 int = 128
                    var v25 string = "hello"
                    v26 := GenS4{F0: false, F1: 703, F2: -807}
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
            }
            _ = fi15
        }
        _ = v12
    } else {
        var v27 int32 = 0
        _ = v27
    }
    v3 = ("" + "foo bar")
    v11 -= v11
    v28 := []int{127, v1}
    switch (("" + v3) + ((((v3 + "hello") + (v3 + v3)) + v3) + "")) {
    case "\n\t":
        for fi29 := 0; fi29 < 7; fi29++ {
            v3 = v3
            v11 = (v11 & 255)
            v30 := v2[(("" + "foo bar") + ("foo bar" + v3))]
            _ = fi29
            _ = v30
        }
    case "hello":
        v1 -= (v1 ^ v1)
    }
    if (!false) {
        var v31 uint64 = 1
        v32 := GenS1{F0: 4294967295, F1: 127, F2: v31}
        _ = v31
        _ = v32
    }
    if true {
        switch "a" {
        case "\n\t":
            var v33 uint8 = 24
            _ = v33
        case "foo bar":
            var v34 int8 = -78
            _ = v34
        }
        if true {
            if false {
                _ = len(v2)
                switch v3 {
                case "a":
                    v35 := GenS4{F0: (688 > 4294967295), F1: 890, F2: 42}
                    _ = v35
                case "hello":
                    switch (v3 + v3) {
                    case "a":
                        for fi36 := 0; fi36 < 17; fi36++ {
                            var v37 int = -581
                            var v38 int32 = 1
                            v39 := v2[((((v3 + v3) + v3) + "a") + ((v3 + v3) + "\n\t"))]
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    case "foo bar":
                        if true {
                            v40 := GenS5{F0: 1, F1: v1}
                            _ = v40
                        } else {
                            v41 := GenS2{F0: 2}
                            for fi42 := 0; fi42 < 11; fi42++ {
                                var v43 string = (("\n\t" + (("foo bar" + ((((((v3 + v3) + v3) + v3) + "foo bar") + (v3 + v3)) + "a")) + "foo bar")) + "\n\t")
                                switch (v1 << ((v11) & 63)) {
                                case 229:
                                    for fi44 := 0; fi44 < 13; fi44++ {
                                        fmt.Println(v3, v11, v1)
                                        v45 := GenS5{F0: 1, F1: (-(fi44))}
                                        _ = fi44
                                        _ = v45
                                    }
                                case 572:
                                    fmt.Println(v3, v43)
                                case 727:
                                    switch "a" {
                                    case "foo bar":
                                        v1 = v1
                                    }
                                }
                                v2[(v43 + ((v3 + v43) + ("" + v43)))] = (((v3 + v3) + "") + v3)
                                _ = fi42
                                _ = v43
                            }
                            if false {
                                var v46 int64 = -426
                                for fi47 := 0; fi47 < 9; fi47++ {
                                    fmt.Println(v1)
                                    if (v10 > float32(((1.5 + 3.14159) + -2.25))) {
                                        v48 := GenS3{F0: -70, F1: -1}
                                        if (!true) {
                                            fmt.Println(v1)
                                        }
                                        if true {
                                            for fi49 := 0; fi49 < 2; fi49++ {
                                                if false {
                                                    fmt.Println(fi47)
                                                }
                                                _ = fi49
                                            }
                                        }
                                        _ = v48
                                    } else {
                                        v11 = v11
                                    }
                                    var v50 string = "foo bar"
                                    _ = fi47
                                    _ = v50
                                }
                                _ = v46
                            }
                            _ = v41
                        }
                    case "\n\t":
                        v1 -= 78
                    default:
                        for fi51 := 0; fi51 < 5; fi51++ {
                            switch v3 {
                            case "":
                                if len(v28) > 0 {
                                    v28[0] = -630
                                }
                            case "foo bar":
                                for fi52 := 0; fi52 < 5; fi52++ {
                                    if (0.0 > ((-2.25 * (-2.25 / float64(v10))) * (float64(float32(-2.25)) / 1.5))) {
                                        switch v3 {
                                        case "hello":
                                            v53 := GenS1{F0: 903, F1: 0, F2: 0}
                                            _ = v53
                                        case "":
                                            v54 := GenS6{F0: (127 < -1)}
                                            _ = v54
                                        }
                                        fmt.Println(v3, fi52)
                                        fmt.Println(v11, v1, fi51)
                                    }
                                    _ = fi52
                                }
                            case "\n\t":
                                delete(v2, (v3 + "hello"))
                            default:
                                v10 = v10
                            }
                            _ = fi51
                        }
                    }
                }
            } else {
                switch v3 {
                case "":
                    v55 := GenS5{F0: 2, F1: v1}
                    _ = v55
                }
            }
        }
        if (!true) {
            var v56 uint64 = 4294967295
            for fi57 := 0; fi57 < 6; fi57++ {
                fmt.Println(v1)
                _ = fi57
            }
            _ = v56
        } else {
            var v58 float64 = float64(((v10 - v10) - (0.1 / 1.5)))
            v1 -= v1
            _ = v58
        }
    } else {
        for fi59 := 0; fi59 < 13; fi59++ {
            switch v3 {
            case "a":
                switch (v1 >> (((v11 / (((v11 | v11)) | 1))) & 63)) {
                case 212:
                    var v60 uint32 = 2147483647
                    _ = v60
                case 44:
                    switch "a" {
                    case "hello":
                        fmt.Println(v10)
                    }
                }
            case "":
                fmt.Println(v11)
            }
            _ = fi59
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v3
}
