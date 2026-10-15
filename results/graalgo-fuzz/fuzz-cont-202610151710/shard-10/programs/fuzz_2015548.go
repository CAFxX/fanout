package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int64; F1 int; F2 uint32 }
type GenS5 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint32 = 88
    var v2 float64 = 0.1
    var v3 uint64 = 0
    var v4 int = -38
    if (!true) {
        var v5 string = "a"
        var v6 int8 = -1
        _ = v5
        _ = v6
    }
    var v7 uint64 = 42
    v8 := []int{253}
    if (!(2 < 1)) {
        var v9 int16 = 32767
        _ = v9
    }
    v2 = (-2.25 / (v2 - (((float64(3.14159) - v2) / 0.1) / (v2 * (v2 * v2)))))
    if true {
        v4 = v4
        if (float64(((1e10 + 1.5) / 1.5)) != -2.25) {
            switch "foo bar" {
            case "foo bar":
                for fi10 := 0; fi10 < 13; fi10++ {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        switch "hello" {
                        case "\n\t":
                            if true {
                                if (2 <= 2) {
                                    switch 428 {
                                    case 681:
                                        v12 := GenS1{F0: (^(v1)), F1: (v1 + (v1 << ((65535) & 63)))}
                                        _ = v12
                                    case 128:
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                    case 4294967295:
                                        if false {
                                            if (2147483647 != 1) {
                                                v7 -= (v3 % (((^(v3))) | 1))
                                                if false {
                                                    if len(v8) > 0 {
                                                        v8[0] = fi11
                                                    }
                                                    var v13 uint8 = 1
                                                    _ = v13
                                                } else {
                                                    if (832 > -1) {
                                                        _ = len(v8)
                                                        switch "a" {
                                                        case "":
                                                            var v14 float32 = 1.5
                                                            _ = v14
                                                        case "hello":
                                                            switch fi10 {
                                                            case -9223372036854775808:
                                                                if len(v8) >= 2 {
                                                                    v8 = v8[1:len(v8)]
                                                                }
                                                            default:
                                                                var v15 int16 = 261
                                                                if true {
                                                                    v8 = append(v8, fi11)
                                                                    v1 += (^(v1))
                                                                    v7 = (v3 & v3)
                                                                }
                                                                switch ("" + "") {
                                                                case "a":
                                                                    if true {
                                                                        var v16 int = -2147483648
                                                                        if false {
                                                                            var v17 float32 = 1.5
                                                                            v18 := GenS1{F0: 2147483647, F1: v1}
                                                                            _ = v17
                                                                            _ = v18
                                                                        } else {
                                                                            if false {
                                                                                v7 = 4294967295
                                                                                v19 := GenS1{F0: (v1 * (^(v1))), F1: v1}
                                                                                var v20 uint = 388
                                                                                _ = v19
                                                                                _ = v20
                                                                            } else {
                                                                                v8 = append(v8, (v4 + fi10))
                                                                            }
                                                                            v3 = v7
                                                                            var v21 int = (v4 | int(v1))
                                                                            _ = v21
                                                                        }
                                                                        fmt.Println(v3)
                                                                        _ = v16
                                                                    } else {
                                                                        v22 := GenS1{F0: (v1 | 42), F1: v1}
                                                                        switch 4294967295 {
                                                                        case 127:
                                                                            v8 = append(v8, -2147483648)
                                                                            var v23 uint32 = v1
                                                                            var v24 uint16 = 32767
                                                                            var v25 bool = false
                                                                            var v26 float64 = v2
                                                                            v27 := GenS1{F0: (v1 >> ((18446744073709551615) & 63)), F1: v1}
                                                                            if (((float64(1e10) * 1e10) / (3.14159 / (v26 / ((((v2 * 0.0) + 3.14159) / 1.5) + 0.1)))) >= float64(1e10)) {
                                                                                if len(v8) >= 2 {
                                                                                    v8 = v8[1:len(v8)]
                                                                                }
                                                                                if len(v8) > 0 {
                                                                                    v8[0] = 4294967295
                                                                                }
                                                                                v23 += (v23 * (v23 << ((333) & 63)))
                                                                            } else {
                                                                                v28 := GenS2{F0: "foo bar", F1: v23}
                                                                                _ = v28
                                                                            }
                                                                            _ = v23
                                                                            _ = v24
                                                                            _ = v25
                                                                            _ = v26
                                                                            _ = v27
                                                                        }
                                                                        v15 = (v15 + v15)
                                                                        _ = v22
                                                                    }
                                                                case "foo bar":
                                                                    if len(v8) > 0 {
                                                                        v8[0] = (v4 % ((v4) | 1))
                                                                    }
                                                                }
                                                                _ = len(v8)
                                                                v15 += (v15 / ((v15) | 1))
                                                                _ = v15
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(fi10, v4)
                                                        }
                                                    }
                                                    v29 := GenS1{F0: 42, F1: v1}
                                                    var v30 float64 = 0.0
                                                    _ = v29
                                                    _ = v30
                                                }
                                            }
                                            v4 = (fi10 / ((fi10) | 1))
                                        }
                                    default:
                                        v1 = v1
                                    }
                                }
                                var v31 uint16 = 0
                                _ = v31
                            } else {
                                v2 -= v2
                            }
                        case "hello":
                            switch (("" + ("\n\t" + ("\n\t" + ""))) + ("" + ("a" + ""))) {
                            case "":
                                _ = len(v8)
                            case "a":
                                v32 := GenS3{F0: v7, F1: -82}
                                _ = v32
                            case "\n\t":
                                var v33 string = ("hello" + "hello")
                                _ = v33
                            }
                        default:
                            v4 += v4
                        }
                        if (v3 == (^(v7))) {
                            v34 := GenS3{F0: v7, F1: 1}
                            v2 -= (float64(float32(v2)) / ((float64(1.5) / v2) * float64((float32(v2) * 1.5))))
                            v3 += 4294967295
                            _ = v34
                        } else {
                            v35 := GenS4{F0: 9223372036854775807, F1: fi10, F2: 1}
                            if len(v8) > 0 {
                                v8[0] = (^(fi10))
                            }
                            _ = v35
                        }
                        v2 = v2
                        _ = fi11
                    }
                    v1 = (v1 / (((^(v1))) | 1))
                    switch "hello" {
                    case "foo bar":
                        for fi36 := 0; fi36 < 12; fi36++ {
                            v8 = append(v8, fi36)
                            var v37 int = (fi36 ^ -1)
                            if (!true) {
                                if (false && (!((v1 >> ((2) & 63)) <= v1))) {
                                    v4 = fi36
                                    var v38 uint32 = (v1 + (v1 ^ (v1 & v1)))
                                    _ = v38
                                } else {
                                    var v39 int8 = 0
                                    _ = v39
                                }
                            } else {
                                v1 = v1
                                v2 = 1.5
                            }
                            var v40 string = ((("hello" + "foo bar") + "\n\t") + "hello")
                            _ = len(v8)
                            _ = fi36
                            _ = v37
                            _ = v40
                        }
                    }
                    _ = fi10
                }
            case "hello":
                fmt.Println(v2)
            case "":
                v41 := GenS4{F0: -9223372036854775808, F1: (v4 << ((18446744073709551615) & 63)), F2: (^(v1))}
                _ = v41
            }
            for fi42 := 0; fi42 < 12; fi42++ {
                for fi43 := 0; fi43 < 4; fi43++ {
                    switch fi42 {
                    case 0:
                        v44 := GenS5{F0: fi42, F1: (v1 & (v1 >> ((1) & 63)))}
                        _ = v44
                    }
                    fmt.Println(fi42, v3)
                    _ = fi43
                }
                fmt.Println(fi42, v4, v4)
                for fi45 := 0; fi45 < 11; fi45++ {
                    _ = len(v8)
                    _ = fi45
                }
                _ = fi42
            }
        }
    } else {
        var v46 int8 = -56
        v1 -= (v1 ^ (v1 & 190))
        _ = v46
    }
    if true {
        var v47 int16 = -783
        var v48 uint16 = 2
        v1 = v1
        _ = v47
        _ = v48
    }
    if (!(-1 != 47)) {
        for fi49 := 0; fi49 < 12; fi49++ {
            if len(v8) > 0 {
                v8[0] = (fi49 | (v4 * (-(v4))))
            }
            if len(v8) > 0 {
                v8[0] = fi49
            }
            _ = fi49
        }
        if false {
            v3 += 2
            fmt.Println(v7)
        } else {
            if (42 < 1) {
                switch 42 {
                case 32767:
                    var v50 bool = false
                    _ = v50
                case 9223372036854775807:
                    if false {
                        v4 = (v4 + -894)
                        var v51 uint8 = 127
                        _ = v51
                    }
                case 490:
                    v52 := GenS4{F0: int64(v4), F1: 32767, F2: (^(v1))}
                    _ = v52
                }
            } else {
                for fi53 := 0; fi53 < 6; fi53++ {
                    if true {
                        var v54 float32 = ((3.14159 * 1e10) + 0.1)
                        for fi55 := 0; fi55 < 8; fi55++ {
                            _ = len(v8)
                            switch fi53 {
                            case -911:
                                var v56 int64 = 42
                                _ = v56
                            case -726:
                                v57 := GenS1{F0: 1, F1: v1}
                                _ = v57
                            default:
                                v8 = append(v8, (fi53 * v4))
                            }
                            _ = fi55
                        }
                        _ = v54
                    } else {
                        v8 = append(v8, fi53)
                        _ = len(v8)
                    }
                    _ = fi53
                }
            }
            v8 = append(v8, (^(v4)))
            if (false || true) {
                fmt.Println(v2, v7, v7)
                switch ("a" + "hello") {
                case "a":
                    switch "" {
                    case "":
                        v8 = append(v8, v4)
                    default:
                        if len(v8) > 0 {
                            v8[0] = v4
                        }
                    }
                }
            } else {
                var v58 int8 = 127
                _ = v58
            }
        }
        var v59 string = "\n\t"
        _ = v59
    } else {
        v60 := GenS5{F0: 32768, F1: (v1 << ((18446744073709551615) & 63))}
        var v61 uint32 = v1
        switch ((("foo bar" + "\n\t") + "\n\t") + (("" + (("a" + "hello") + "hello")) + "a")) {
        case "a":
            for fi62 := 0; fi62 < 5; fi62++ {
                v7 += (v7 * 1)
                v2 = (-2.25 - 0.1)
                _ = fi62
            }
        }
        _ = v60
        _ = v61
    }
    var v63 int64 = -1
    v8 = append(v8, v4)
    for fi64 := 0; fi64 < 5; fi64++ {
        if false {
            v65 := GenS3{F0: v7, F1: 0}
            var v66 int32 = 2147483647
            var v67 int64 = int64(v1)
            _ = v65
            _ = v66
            _ = v67
        } else {
            v1 = (v1 * v1)
            v7 += 18446744073709551615
            var v68 int = fi64
            _ = v68
        }
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi64
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v7
    _ = v8
}
