package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 int32; F2 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32; F1 uint8 }

func main() {
    var v1 string = "a"
    var v2 int = 9223372036854775807
    var v3 uint8 = 127
    v4 := []uint32{}
    v5 := make(map[int]int)
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 12; fi7++ {
            switch (v1 + ("a" + "")) {
            case "hello":
                var v8 int = fi6
                _ = v8
            case "\n\t":
                fmt.Println(v3, fi6)
            }
            if len(v4) > 0 {
                v4[0] = 2
            }
            _ = fi7
        }
        if len(v4) > 0 {
            v4[0] = 0
        }
        if true {
            v3 = 1
            switch "foo bar" {
            case "foo bar":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            for fi9 := 0; fi9 < 15; fi9++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v10 int16 = 0
                if false {
                    fmt.Println(v2, v3, fi6)
                    v10 += 1000
                    switch fi9 {
                    case 42:
                        var v11 int64 = -1
                        _ = v11
                    case -28:
                        v3 -= 128
                    case 127:
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                    default:
                        v12 := GenS1{F0: (-(fi9)), F1: 759, F2: 42}
                        _ = v12
                    }
                }
                _ = fi9
                _ = v10
            }
        }
        _ = fi6
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            if (752 <= -1) {
                if (((!false) && true) || (((127 < 42) || (false && (("" + v1) > "a"))) && ((765 >= -2147483648) || false))) {
                    var v15 int = (fi14 >> ((251) & 63))
                    var v16 bool = (false || true)
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                    if v16 {
                        var v17 float64 = 0.0
                        if true {
                            var v18 int16 = 42
                            switch (v15 * fi14) {
                            case 127:
                                var v19 string = v1
                                _ = v19
                            case 811:
                                v18 = (v18 >> ((65535) & 63))
                            case 9223372036854775807:
                                if v16 {
                                    _ = len(v4)
                                    var v20 uint = 932
                                    v15 = (v15 << ((2) & 63))
                                    _ = v20
                                }
                            }
                            fmt.Println(fi14, v3)
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                            v1 = (("foo bar" + v1) + v1)
                            _ = v18
                        } else {
                            v21 := GenS1{F0: v2, F1: 360, F2: 704}
                            if (304 != 0) {
                                var v22 int16 = 0
                                if v16 {
                                    if v16 {
                                        if v16 {
                                            fmt.Println(v16)
                                        }
                                        fmt.Println(v17)
                                        _ = len(v4)
                                        if (!((v15 / ((fi13) | 1)) == (fi14 / (((v15 & v2)) | 1)))) {
                                            var v23 uint = 1
                                            _ = v23
                                        } else {
                                            var v24 int8 = 78
                                            _ = v24
                                        }
                                        _ = len(v5)
                                    } else {
                                        delete(v5, (-(v2)))
                                    }
                                    v1 = v1
                                }
                                _ = v22
                            } else {
                                var v25 float64 = ((v17 + v17) * 1.5)
                                var v26 bool = false
                                if (0 > -2147483648) {
                                    v17 -= v25
                                    delete(v5, 32768)
                                }
                                _ = v25
                                _ = v26
                            }
                            switch "" {
                            case "a":
                                switch (((((v1 + (v1 + (v1 + "\n\t"))) + (v1 + (((v1 + v1) + v1) + "a"))) + "foo bar") + v1) + v1) {
                                case "":
                                    v27 := GenS1{F0: (^(v15)), F1: 65535, F2: 490}
                                    _ = v27
                                default:
                                    v3 = (v3 | (v3 / (((v3 & (v3 - v3))) | 1)))
                                }
                            }
                            _ = v21
                        }
                        v28 := GenS2{F0: 1}
                        if (0.0 == (3.14159 * 0.1)) {
                            _ = len(v5)
                        }
                        _ = v17
                        _ = v28
                    }
                    _ = v15
                    _ = v16
                } else {
                    if (!true) {
                        v1 = v1
                    }
                    v29 := GenS2{F0: 537}
                    _ = v29
                }
            } else {
                fmt.Println(v3, fi13, fi14)
                fmt.Println(fi13, v1, v3)
                v5[(fi13 + fi13)] = 9223372036854775807
                var v30 int32 = -2147483648
                if (true && false) {
                    if (42 > -1) {
                        if true {
                            v1 = ("\n\t" + v1)
                            if (!(9223372036854775807 <= 0)) {
                                v31 := v5[int(v30)]
                                fmt.Println(v2, v31, v3)
                                if ((585 < 42) && false) {
                                    v1 = "a"
                                } else {
                                    v2 = v2
                                    var v32 int32 = v30
                                    v33 := v5[v31]
                                    _ = v32
                                    _ = v33
                                }
                                v34 := GenS3{F0: -32768, F1: true}
                                fmt.Println(v30, v1, v2)
                                _ = v31
                                _ = v34
                            }
                        } else {
                            if false {
                                var v35 uint64 = 36
                                v35 = (v35 | 490)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v2 = v2
                                switch v1 {
                                case "hello":
                                    fmt.Println(v35, v3)
                                case "a":
                                    v30 -= (v30 * v30)
                                case "":
                                    _ = len(v5)
                                default:
                                    v35 = 4294967295
                                }
                                _ = v35
                            }
                            v36 := GenS4{F0: uint(v3)}
                            _ = v36
                        }
                        fmt.Println(v3, fi13)
                    } else {
                        v3 -= (v3 * 128)
                        if false {
                            v37 := GenS1{F0: (fi13 * (fi14 * fi13)), F1: 2, F2: 160}
                            v30 += (v30 + v30)
                            _ = v37
                        } else {
                            switch (v1 + (v1 + "\n\t")) {
                            case "a":
                                fmt.Println(fi13, v1, v30)
                            }
                        }
                        var v38 float32 = (((1e10 - 0.0) + float32(1e10)) + ((1.5 / (1.5 + (0.0 + 0.0))) + float32(-2.25)))
                        _ = v38
                    }
                    var v39 uint64 = 617
                    _ = v39
                }
                v40 := GenS1{F0: 0, F1: 49, F2: 4294967295}
                _ = v30
                _ = v40
            }
            _ = fi14
        }
        _ = fi13
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        v2 = -887
        for fi42 := 0; fi42 < 10; fi42++ {
            delete(v5, (^(v2)))
            if (!false) {
                var v43 uint64 = 1
                v2 = int(v3)
                _ = v43
            }
            v44 := GenS5{F0: fi42, F1: 223, F2: "a"}
            v2 += 42
            _ = fi42
            _ = v44
        }
        _ = fi41
    }
    if (false || true) {
        for fi45 := 0; fi45 < 1; fi45++ {
            if (9223372036854775807 < -1) {
                v3 += 2
                v1 = (v1 + "\n\t")
                var v46 int = v2
                _ = v46
            } else {
                var v47 uint16 = 306
                for fi48 := 0; fi48 < 13; fi48++ {
                    v1 = v1
                    v47 -= (v47 / (((^(v47))) | 1))
                    fmt.Println(v1, v2, v1)
                    _ = fi48
                }
                _ = v47
            }
            var v49 uint16 = 32767
            var v50 uint = 769
            _ = fi45
            _ = v49
            _ = v50
        }
        fmt.Println(v1, v3)
    } else {
        var v51 float32 = 1e10
        _ = v51
    }
    fmt.Println(v2, v2, v2)
    v52 := GenS6{F0: 0}
    var v53 int64 = 9223372036854775807
    var v54 float32 = 0.0
    fmt.Println(v3, v54, v54)
    v1 = v1
    if (v2 < 32768) {
        v55 := GenS7{F0: ((v1 + ("hello" + ("" + "foo bar"))) + (v1 + (v1 + ("" + ""))))}
        _ = v55
    }
    switch v1 {
    case "foo bar":
        if ((0.0 / 1e10) >= ((1e10 / 0.1) - 0.1)) {
            v56 := GenS8{F0: 42, F1: v3}
            _ = v56
        } else {
            _ = len(v5)
            var v57 uint = 1
            v54 += (v54 * v54)
            _ = v57
        }
    case "\n\t":
        fmt.Println(v54, v54, v3)
    }
    for fi58 := 0; fi58 < 17; fi58++ {
        if (!(v2 > -1)) {
            var v59 int = -2147483648
            _ = v59
        }
        if (v1 < "a") {
            v2 = fi58
            v54 -= (1.5 * (float32((-2.25 / (float64(v54) / 3.14159))) / (3.14159 / v54)))
            for fi60 := 0; fi60 < 1; fi60++ {
                v54 = v54
                _ = fi60
            }
        }
        _ = fi58
    }
    var v61 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v61
}
