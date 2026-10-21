package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 string; F1 bool; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int8; F1 uint; F2 int }
func genFn0() int {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 string = "foo bar"
        v9 := GenS1{F0: 40, F1: 450, F2: 42}
        if (fi7 > (-(fi7))) {
            var v10 int16 = 1
            for fi11 := 0; fi11 < 1; fi11++ {
                v8 = v8
                v12 := []int{(fi7 ^ fi7), (fi11 ^ (fi11 >> ((18446744073709551615) & 63))), fi11}
                _ = len(v12)
                if ((v8 < ("" + v8)) || true) {
                    if false {
                        if ("foo bar" <= v8) {
                            var v13 uint32 = 858
                            _ = v13
                        }
                    } else {
                        v10 = 518
                        v10 += v10
                        switch (^(fi11)) {
                        case 2147483647:
                            v10 = v10
                        case -1:
                            if false {
                                fmt.Println(fi7)
                            } else {
                                _ = len(v12)
                                v14 := GenS1{F0: 218, F1: 65535, F2: 56}
                                v8 = (v8 + (v8 + v8))
                                v8 = v8
                                v15 := GenS1{F0: 2, F1: 0, F2: 127}
                                _ = v14
                                _ = v15
                            }
                        default:
                            switch fi7 {
                            case -9223372036854775808:
                                if (!false) {
                                    if true {
                                        v8 = v8
                                        var v16 uint32 = 789
                                        v8 = v8
                                        var v17 uint32 = v16
                                        var v18 int64 = 0
                                        v16 = v16
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    } else {
                                        v12 = append(v12, fi11)
                                    }
                                }
                            case 545:
                                var v19 uint64 = 1
                                _ = v19
                            case 32768:
                                v8 = (v8 + "")
                            default:
                                switch v8 {
                                case "\n\t":
                                    fmt.Println(v10, fi7, v8)
                                case "":
                                    var v20 int32 = 477
                                    _ = v20
                                case "hello":
                                    switch (((v8 + ("" + ("foo bar" + v8))) + "hello") + "a") {
                                    case "foo bar":
                                        v21 := GenS2{F0: 2, F1: 18446744073709551615}
                                        _ = v21
                                    case "hello":
                                        v8 = "\n\t"
                                    case "\n\t":
                                        v10 -= v10
                                    default:
                                        v22 := GenS3{F0: v8, F1: false, F2: 2147483647}
                                        _ = v22
                                    }
                                }
                                var v23 uint32 = 42
                                v24 := GenS2{F0: 0, F1: 1}
                                _ = v23
                                _ = v24
                            }
                        }
                    }
                }
                switch fi11 {
                case 32767:
                    var v25 float64 = float64(float32(1e10))
                    _ = v25
                case 454:
                    v26 := GenS4{F0: v8}
                    _ = v26
                default:
                    if true {
                        var v27 uint8 = 1
                        var v28 int16 = -775
                        var v29 float64 = 1.5
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                }
                _ = fi11
                _ = v12
            }
            _ = v10
        } else {
            if (1 > 0) {
                v30 := GenS5{F0: int8(fi7), F1: 2}
                switch "foo bar" {
                case "":
                    fmt.Println(v8, fi7)
                case "foo bar":
                    if false {
                        switch fi7 {
                        case 565:
                            v31 := GenS4{F0: v8}
                            _ = v31
                        default:
                            fmt.Println(v8, fi7)
                        }
                        v32 := GenS1{F0: 114, F1: 0, F2: 127}
                        _ = v32
                    }
                }
                var v33 int32 = 552
                _ = v30
                _ = v33
            }
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v34 float64 = -2.25
    v35 := GenS3{F0: "", F1: (42 == 65535), F2: 666}
    if true {
        for fi36 := 0; fi36 < 17; fi36++ {
            v34 = ((v34 * 3.14159) / 1.5)
            v37 := []int{(fi36 >> ((1) & 63)), (fi36 / ((1) | 1)), -2147483648}
            v38 := GenS1{F0: 97, F1: 42, F2: 0}
            _ = fi36
            _ = v37
            _ = v38
        }
        var v39 bool = true
        _ = v39
    }
    _ = v34
    _ = v35
    return 1
}

func main() {
    var v1 float32 = 1.5
    var v2 int = 32767
    v3 := []int{875, 32768, v2}
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    v2 -= v2
    v1 += 3.14159
    var v40 int = -9223372036854775808
    v40 = (v2 - genFn0())
    _ = len(v4)
    for fi41 := 0; fi41 < 16; fi41++ {
        var v42 float32 = (3.14159 - v1)
        var v43 int = v2
        _ = fi41
        _ = v42
        _ = v43
    }
    if ((!false) && false) {
        v2 = 49
    }
    switch genFn0() {
    case 32768:
        _ = len(v4)
    case 2147483647:
        v2 = (v40 ^ genFn0())
    }
    var v44 uint32 = 532
    v6[(v44 - v44)] = "hello"
    var v45 int = (v40 - (^(v40)))
    var v46 float32 = (v1 * v1)
    for fi47 := 0; fi47 < 8; fi47++ {
        for fi48 := 0; fi48 < 14; fi48++ {
            v40 = (fi48 | genFn0())
            v45 = 32768
            _ = fi48
        }
        for fi49 := 0; fi49 < 10; fi49++ {
            v3 = append(v3, (v45 >> ((42) & 63)))
            if (false || true) {
                v2 = (v40 ^ genFn0())
                var v50 int32 = -1
                v50 -= (^(v50))
                _ = v50
            } else {
                if (!(131 != 204)) {
                    v51 := GenS5{F0: 0, F1: 1}
                    if (65535 >= 0) {
                        fmt.Println(v2, v46)
                        var v52 string = ""
                        if ((!(!false)) && false) {
                            v53 := GenS6{F0: 879, F1: -32768, F2: (v44 * 71)}
                            fmt.Println(v1, v2)
                            _ = v53
                        }
                        _ = v52
                    }
                    var v54 string = "\n\t"
                    _ = v51
                    _ = v54
                } else {
                    delete(v6, (v44 >> ((0) & 63)))
                }
                fmt.Println(v1)
                v44 = 65535
            }
            v46 -= v1
            _ = fi49
        }
        _ = fi47
    }
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v45)
    case "":
        v45 = -9223372036854775808
    case "a":
        if ((!true) && (1 != -1)) {
            var v55 int32 = 42
            v1 = (((float32((float64(v1) + 0.1)) / float32(1e10)) / ((float32(float64(v1)) - -2.25) - 1e10)) + (float32(0.1) / 3.14159))
            for fi56 := 0; fi56 < 19; fi56++ {
                var v57 uint64 = 784
                if ((!true) && (!(v40 != 32768))) {
                    var v58 int32 = (v55 % (((v55 & (^(v55)))) | 1))
                    _ = v58
                }
                _ = fi56
                _ = v57
            }
            _ = v55
        }
    }
    switch "hello" {
    case "":
        fmt.Println(v46)
    case "\n\t":
        if ((^(v45)) < (v45 / ((32767) | 1))) {
            v59 := GenS7{F0: 0, F1: 247, F2: 0}
            v1 += (v46 + v1)
            if (281 < 255) {
                for fi60 := 0; fi60 < 18; fi60++ {
                    var v61 int8 = -70
                    _ = fi60
                    _ = v61
                }
            }
            _ = v59
        } else {
            switch ("hello" + "\n\t") {
            case "foo bar":
                for fi62 := 0; fi62 < 5; fi62++ {
                    var v63 uint8 = 7
                    if false {
                        var v64 uint8 = 0
                        for fi65 := 0; fi65 < 6; fi65++ {
                            switch (fi65 | genFn0()) {
                            case 1:
                                v66 := v4["hello"]
                                _ = v66
                            case 128:
                                v44 = (v44 >> ((1) & 63))
                            case 4294967295:
                                fmt.Println(v64)
                            }
                            _ = fi65
                        }
                        v63 = v64
                        _ = v64
                    }
                    v45 = 2147483647
                    _ = fi62
                    _ = v63
                }
            case "":
                if false {
                    if false {
                        delete(v6, v44)
                        if (true || (2147483647 > v44)) {
                            v67 := v5["\n\t"]
                            var v68 int32 = -695
                            for fi69 := 0; fi69 < 7; fi69++ {
                                if len(v3) > 0 {
                                    v3[0] = v2
                                }
                                _ = fi69
                            }
                            _ = v67
                            _ = v68
                        } else {
                            switch genFn0() {
                            case 0:
                                v1 += v46
                            }
                        }
                        v3 = append(v3, 9223372036854775807)
                    } else {
                        var v70 string = "foo bar"
                        _ = v70
                    }
                }
            }
            var v71 uint64 = 18446744073709551615
            if true {
                var v72 int64 = 408
                fmt.Println(v46)
                for fi73 := 0; fi73 < 16; fi73++ {
                    v72 -= -9223372036854775808
                    _ = fi73
                }
                _ = v72
            }
            _ = v71
        }
    case "foo bar":
        var v74 string = ("hello" + "")
        _ = v74
    }
    v75 := GenS8{F0: 1, F1: 65535, F2: 52}
    v76 := GenS8{F0: 42, F1: 0, F2: genFn0()}
    v46 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v75
    _ = v76
}
