package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 int8; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int8; F2 int32 }
type GenS5 struct { F0 uint8; F1 int; F2 uint32 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 uint; F2 uint64 }

func main() {
    var v1 uint32 = 947
    v2 := make(map[int]int)
    var v3 uint = 42
    var v4 int8 = -128
    var v5 int16 = -32768
    var v6 float64 = float64(0.1)
    v1 += (v1 << (((v3 & 2)) & 63))
    var v7 float32 = ((float32(1e10) + 0.1) / 0.1)
    v1 = v1
    var v8 int = 4294967295
    v5 -= 730
    v9 := []int{(v8 << (((v3 << ((v3) & 63))) & 63)), (v8 * v8)}
    var v10 float32 = float32(1e10)
    if (true || (!(-1 >= 42))) {
        v11 := GenS1{F0: ""}
        v2[(^(v8))] = (v8 ^ (v8 ^ v8))
        v9 = append(v9, int(v1))
        _ = v11
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            switch 1 {
            case 32768:
                v1 += v1
            }
            switch v8 {
            case -2147483648:
                v9 = append(v9, 581)
            case 32768:
                v13 := GenS1{F0: "hello"}
                _ = v13
            case -9223372036854775808:
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 int32 = -2147483648
                    _ = len(v2)
                    v16 := GenS2{F0: "", F1: v4, F2: v3}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 11; fi17++ {
                    v3 = v3
                    _ = fi17
                }
            }
            v18 := GenS3{F0: int16(fi12)}
            _ = fi12
            _ = v18
        }
        v19 := GenS4{F0: (128 == 1), F1: 42, F2: 1}
        _ = v19
    }
    if false {
        switch v8 {
        case 32767:
            if (-2.25 != -2.25) {
                var v20 uint8 = 190
                for fi21 := 0; fi21 < 11; fi21++ {
                    for fi22 := 0; fi22 < 1; fi22++ {
                        var v23 uint8 = (^(v20))
                        _ = fi22
                        _ = v23
                    }
                    var v24 bool = (v5 > (v5 >> ((572) & 63)))
                    _ = fi21
                    _ = v24
                }
                var v25 int32 = 217
                _ = v20
                _ = v25
            }
        default:
            var v26 uint8 = 0
            _ = v26
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        var v27 int32 = -717
        _ = v27
    } else {
        var v28 string = "foo bar"
        _ = v28
    }
    fmt.Println(v6)
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 uint8 = 180
        if (0.0 > ((v6 / 0.1) / v6)) {
            v4 += v4
        }
        _ = fi29
        _ = v30
    }
    if (1 == 2) {
        switch "" {
        case "\n\t":
            v4 = int8(v3)
        case "hello":
            var v31 uint = 1
            _ = v31
        default:
            if len(v9) > 0 {
                v9[0] = v8
            }
        }
        if false {
            v7 = (v10 * 3.14159)
            fmt.Println(v6, v3)
        } else {
            var v32 int = (v8 << (((v3 / (((v3 * 255)) | 1))) & 63))
            _ = v32
        }
    } else {
        var v33 uint = (^(v3))
        _ = v33
    }
    if false {
        var v34 int32 = -834
        if ((v3 / ((v3) | 1)) == v3) {
            v3 -= v3
            switch "a" {
            case "foo bar":
                if ((v3 - 159) == v3) {
                    for fi35 := 0; fi35 < 3; fi35++ {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        _ = fi35
                    }
                    var v36 bool = false
                    switch (-(v8)) {
                    case -778:
                        var v37 float32 = v10
                        _ = v37
                    case -9223372036854775808:
                        v1 += (^(v1))
                    default:
                        if true {
                            var v38 string = ("a" + (("\n\t" + "\n\t") + "hello"))
                            var v39 int32 = v34
                            _ = v38
                            _ = v39
                        } else {
                            switch "" {
                            case "\n\t":
                                v40 := GenS5{F0: 0, F1: (-(v8)), F2: 359}
                                _ = v40
                            case "":
                                if (float64((v10 * v10)) == (1e10 / v6)) {
                                    delete(v2, (v8 + -2147483648))
                                    for fi41 := 0; fi41 < 20; fi41++ {
                                        if v36 {
                                            for fi42 := 0; fi42 < 4; fi42++ {
                                                if v36 {
                                                    var v43 int16 = (v5 << (((v3 / ((v3) | 1))) & 63))
                                                    var v44 float32 = v10
                                                    if v36 {
                                                        var v45 float64 = float64(0.1)
                                                        _ = v45
                                                    }
                                                    _ = v43
                                                    _ = v44
                                                }
                                                _ = fi42
                                            }
                                            var v46 float32 = v10
                                            switch (fi41 & 2147483647) {
                                            case -1:
                                                switch ("hello" + ("foo bar" + ("\n\t" + "\n\t"))) {
                                                case "hello":
                                                    delete(v2, 9223372036854775807)
                                                case "\n\t":
                                                    if (!((float64(-2.25) < v6) && false)) {
                                                        if (false || (980 != v5)) {
                                                            v1 = (v1 ^ (v1 * v1))
                                                        } else {
                                                            v47 := GenS1{F0: "hello"}
                                                            v10 += 0.0
                                                            fmt.Println(v8, v5)
                                                            _ = v47
                                                        }
                                                        var v48 uint8 = 1
                                                        fmt.Println(v10)
                                                        _ = v48
                                                    }
                                                case "a":
                                                    var v49 bool = (65535 > v3)
                                                    _ = v49
                                                default:
                                                    for fi50 := 0; fi50 < 16; fi50++ {
                                                        switch (fi50 | fi50) {
                                                        case 128:
                                                            var v51 uint64 = 42
                                                            _ = v51
                                                        }
                                                        _ = fi50
                                                    }
                                                }
                                            case 42:
                                                var v52 int = fi41
                                                _ = v52
                                            case 9223372036854775807:
                                                var v53 float64 = (v6 - 0.0)
                                                _ = v53
                                            }
                                            _ = v46
                                        } else {
                                            v54 := GenS6{F0: v8}
                                            _ = v54
                                        }
                                        var v55 int16 = v5
                                        v56 := GenS7{F0: "a", F1: 237, F2: 18446744073709551615}
                                        _ = fi41
                                        _ = v55
                                        _ = v56
                                    }
                                    fmt.Println(v4, v34)
                                } else {
                                    v9 = append(v9, v8)
                                    var v57 int64 = -9223372036854775808
                                    if ((v36 && ((("" + "a") + "\n\t") >= "hello")) || v36) {
                                        var v58 int32 = 42
                                        v59 := GenS5{F0: 255, F1: v8, F2: 470}
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = v57
                                }
                            case "hello":
                                var v60 int8 = -1
                                _ = v60
                            }
                            var v61 int16 = v5
                            v62 := GenS5{F0: 128, F1: 32768, F2: (v1 & 65535)}
                            _ = v61
                            _ = v62
                        }
                    }
                    _ = v36
                } else {
                    v2[(-(v8))] = v8
                }
            case "hello":
                delete(v2, (v8 / ((v8) | 1)))
            }
            fmt.Println(v5, v10, v8)
        }
        _ = v34
    }
    v63 := GenS7{F0: ((("foo bar" + ("\n\t" + "hello")) + ("foo bar" + (("\n\t" + ("a" + (("foo bar" + ("" + "hello")) + ("foo bar" + "hello")))) + ("foo bar" + "")))) + "a"), F1: v3, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v7
    _ = v8
    _ = v9
}
