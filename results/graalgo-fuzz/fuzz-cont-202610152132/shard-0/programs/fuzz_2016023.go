package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 uint; F1 uint; F2 string }
type GenS8 struct { F0 uint16 }

func main() {
    var v1 uint = 1
    v2 := []uint32{2147483648, 1, 2}
    v3 := []uint32{}
    if (1000 == 627) {
        v2 = append(v2, 2147483647)
        var v4 string = "foo bar"
        if false {
            var v5 int8 = 100
            _ = v5
        }
        _ = v4
    }
    if (("foo bar" + "") != "a") {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 uint16 = 929
            _ = fi6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: (-(fi8)), F1: 2147483647}
            v10 := GenS2{F0: 666, F1: 2, F2: 1}
            for fi11 := 0; fi11 < 19; fi11++ {
                if ((true && (199 < 42)) && true) {
                    v12 := GenS1{F0: fi11, F1: -1}
                    if (true && true) {
                        _ = len(v2)
                        v13 := GenS2{F0: 42, F1: 42, F2: 284}
                        _ = v13
                    } else {
                        v1 = (v1 * (^(v1)))
                        var v14 float32 = 1.5
                        fmt.Println(fi8)
                        _ = v14
                    }
                    v3 = append(v3, 65535)
                    _ = v12
                }
                var v15 uint64 = 410
                switch "foo bar" {
                case "foo bar":
                    if (uint16(v15) != 28) {
                        var v16 int64 = 1
                        v15 += 477
                        switch fi11 {
                        case 42:
                            var v17 uint16 = 32767
                            _ = v17
                        case 4294967295:
                            v18 := GenS3{F0: 0}
                            _ = v18
                        case 32768:
                            var v19 uint64 = uint64(v1)
                            _ = v19
                        }
                        _ = v16
                    } else {
                        switch "a" {
                        case "hello":
                            var v20 uint8 = 129
                            _ = v20
                        case "foo bar":
                            v21 := GenS4{F0: v15, F1: 1, F2: 255}
                            _ = v21
                        default:
                            v15 = v15
                            _ = len(v2)
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v15, fi11)
                            case "hello":
                                var v22 int = fi11
                                _ = v22
                            case "":
                                v23 := GenS5{F0: (true || (((v1 >> ((v1) & 63)) >= (v1 + v1)) || (0 == 65535))), F1: 2147483647, F2: v1}
                                _ = v23
                            }
                        }
                        if len(v3) > 0 {
                            v3[0] = 65535
                        }
                        var v24 uint16 = 42
                        _ = v24
                    }
                case "a":
                    if false {
                        var v25 string = "foo bar"
                        _ = v25
                    } else {
                        var v26 string = (("\n\t" + "\n\t") + ("a" + "\n\t"))
                        v27 := GenS6{F0: (^(v15)), F1: "foo bar"}
                        v28 := GenS6{F0: (v15 + v15), F1: "\n\t"}
                        if (-446 > int(v1)) {
                            v26 = "a"
                            v26 = v26
                            var v29 string = ("a" + (v26 + v26))
                            _ = v29
                        } else {
                            var v30 uint32 = 735
                            fmt.Println(v15)
                            v3 = append(v3, (v30 ^ (v30 >> (((v1 >> (((v1 & (^(v1)))) & 63))) & 63))))
                            _ = v30
                        }
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                default:
                    v31 := GenS3{F0: 0}
                    _ = v31
                }
                if (!true) {
                    if (("a" + ("\n\t" + ("foo bar" + "foo bar"))) != "\n\t") {
                        fmt.Println(v15)
                        v15 += (v15 | v15)
                        if ((4 > 135) || false) {
                            var v32 bool = ((-(fi11)) != (fi11 << ((v1) & 63)))
                            var v33 int8 = -90
                            var v34 uint64 = (^(v15))
                            v3 = append(v3, 2)
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    } else {
                        v15 = v15
                        v35 := GenS5{F0: true, F1: 42, F2: v1}
                        _ = v35
                    }
                    var v36 string = ((("" + ("foo bar" + "a")) + "\n\t") + "a")
                    _ = v36
                } else {
                    v15 -= v15
                    v37 := GenS5{F0: false, F1: 443, F2: v1}
                    _ = v37
                }
                switch (fi11 | 128) {
                case -9223372036854775808:
                    _ = len(v3)
                    var v38 uint8 = 127
                    v38 -= (v38 * 1)
                    v15 += 1
                    _ = v38
                }
                _ = fi11
                _ = v15
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        if true {
            v2 = append(v2, 893)
            if (((1 != 363) || (-1 >= 127)) && (42 < -9223372036854775808)) {
                var v39 float32 = (3.14159 - 3.14159)
                _ = v39
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        v40 := GenS5{F0: true, F1: 727, F2: (v1 + 162)}
        _ = v40
    }
    var v41 uint = v1
    var v42 float32 = float32((0.1 / 1.5))
    var v43 uint32 = 779
    v44 := GenS5{F0: (((v41 >= (v1 - v1)) || true) && false), F1: 42, F2: 18446744073709551615}
    fmt.Println(v43, v43, v42)
    if (772 <= 73) {
        fmt.Println(v1, v41, v41)
        var v45 int8 = 42
        _ = v45
    } else {
        if (false || (!false)) {
            for fi46 := 0; fi46 < 16; fi46++ {
                var v47 int16 = 32767
                var v48 uint = (v1 % ((v1) | 1))
                for fi49 := 0; fi49 < 19; fi49++ {
                    v43 -= v43
                    _ = fi49
                }
                _ = fi46
                _ = v47
                _ = v48
            }
            if (1.5 < 1.5) {
                for fi50 := 0; fi50 < 7; fi50++ {
                    var v51 float32 = float32(1.5)
                    _ = fi50
                    _ = v51
                }
            } else {
                v52 := GenS7{F0: (v41 / ((v41) | 1)), F1: 441, F2: "hello"}
                v41 -= v1
                _ = v52
            }
            var v53 uint32 = (v43 * v43)
            _ = v53
        }
        if true {
            var v54 uint64 = 42
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v54
        } else {
            v42 += v42
            v55 := GenS5{F0: (4294967295 != 18446744073709551615), F1: (^(v43)), F2: v41}
            _ = v55
        }
    }
    fmt.Println(v41)
    fmt.Println(v42)
    if (false && true) {
        v3 = append(v3, v43)
    } else {
        v2 = append(v2, v43)
        var v56 int32 = 0
        for fi57 := 0; fi57 < 8; fi57++ {
            v1 = (^(v41))
            fmt.Println(v41)
            if ((0.0 * 3.14159) != (3.14159 + 1e10)) {
                var v58 int8 = 127
                _ = v58
            }
            _ = fi57
        }
        _ = v56
    }
    v1 = (v41 + v41)
    for fi59 := 0; fi59 < 6; fi59++ {
        v43 = v43
        _ = fi59
    }
    for fi60 := 0; fi60 < 10; fi60++ {
        v43 = 0
        _ = fi60
    }
    var v61 float32 = 1.5
    v62 := GenS8{F0: 255}
    v1 -= v41
    v42 = v61
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v61
    _ = v62
}
