package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 int8; F2 bool }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string; F1 bool }
type GenS9 struct { F0 int; F1 uint16 }
type GenS10 struct { F0 int16; F1 int32; F2 int8 }
type GenS11 struct { F0 uint32 }
func genFn0(p0 string, p1 uint16, p2 string) uint64 {
    v6 := GenS1{F0: 65535, F1: -342}
    p2 = p0
    v7 := GenS2{F0: -100}
    var v8 bool = (1 != 1)
    var v9 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int64) uint32 {
    fmt.Println(p0)
    v10 := []int{-9223372036854775808, 32767, -508}
    _ = p0
    _ = v10
    return 2
}

func main() {
    var v1 uint32 = 662
    v2 := make(map[string]string)
    var v3 int64 = -510
    v4 := []uint32{(^(v1)), v1, 4, v1}
    v5 := make(map[uint32]int)
    switch -2147483648 {
    case 32767:
        v11 := GenS3{F0: 9223372036854775807}
        _ = v11
    case 1:
        switch 32767 {
        case 4294967295:
            if (true || true) {
                v12 := GenS3{F0: v3}
                _ = v12
            } else {
                var v13 uint16 = 65535
                for fi14 := 0; fi14 < 11; fi14++ {
                    if (!(-1 != 130)) {
                        var v15 float32 = 3.14159
                        _ = v15
                    } else {
                        fmt.Println(v1, v13, v13)
                        for fi16 := 0; fi16 < 3; fi16++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v17 int8 = 42
                            switch "a" {
                            case "\n\t":
                                var v18 uint64 = 679
                                _ = v18
                            case "hello":
                                switch ("\n\t" + ("foo bar" + "\n\t")) {
                                case "":
                                    v17 = (-(v17))
                                    var v19 int = (^(fi16))
                                    _ = v19
                                default:
                                    v3 -= (v3 ^ v3)
                                }
                            case "a":
                                switch 32768 {
                                case 1:
                                    v20 := GenS3{F0: (v3 << ((2) & 63))}
                                    _ = v20
                                case 912:
                                    if len(v4) > 0 {
                                        v4[0] = 2147483648
                                    }
                                    var v21 int64 = 42
                                    var v22 int16 = 42
                                    _ = v21
                                    _ = v22
                                }
                                v13 += v13
                            default:
                                v17 = (v17 * v17)
                            }
                            _ = fi16
                            _ = v17
                        }
                    }
                    v23 := GenS4{F0: 1}
                    _ = fi14
                    _ = v23
                }
                _ = v13
            }
        case 42:
            v1 = (^(v1))
        case 9223372036854775807:
            var v24 uint32 = genFn1((v3 - v3))
            _ = v24
        }
    }
    var v25 float64 = (((float64(3.14159) / float64(((0.1 + (0.1 + (-2.25 * (((-2.25 + 3.14159) * -2.25) + 3.14159)))) + (float32(1.5) + 0.0)))) / ((3.14159 + -2.25) - 3.14159)) * (1e10 * -2.25))
    v4 = append(v4, v1)
    var v26 uint16 = uint16(v1)
    if (-44 != 87) {
        var v27 uint32 = genFn1((^(v3)))
        _ = v27
    }
    var v28 uint64 = 2
    v1 += v1
    switch "foo bar" {
    case "foo bar":
        for fi29 := 0; fi29 < 10; fi29++ {
            if false {
                fmt.Println(v1, v25)
                var v30 float32 = 0.1
                if true {
                    fmt.Println(v28, fi29, v3)
                    for fi31 := 0; fi31 < 11; fi31++ {
                        var v32 float32 = -2.25
                        if ((0.1 * (((0.1 / v32) / (-2.25 / float32(v25))) * ((0.0 / 1e10) * ((v30 / (3.14159 + v30)) * v32)))) > v30) {
                            v1 -= (v1 | v1)
                            v5[uint32(v28)] = fi29
                            v28 = genFn0(("foo bar" + ""), (v26 + 2), ("foo bar" + "\n\t"))
                        } else {
                            fmt.Println(v3, v1)
                        }
                        _ = len(v5)
                        _ = fi31
                        _ = v32
                    }
                    v30 += (1.5 * (0.1 / (0.0 + -2.25)))
                }
                _ = v30
            } else {
                if true {
                    _ = len(v4)
                    var v33 int32 = -2147483648
                    _ = v33
                }
            }
            if false {
                var v34 int16 = int16(v1)
                v35 := GenS2{F0: 100}
                if ((v28 & v28) >= (v28 ^ 0)) {
                    fmt.Println(v26)
                    var v36 uint32 = 42
                    _ = v36
                }
                _ = v34
                _ = v35
            }
            _ = fi29
        }
    default:
        for fi37 := 0; fi37 < 12; fi37++ {
            for fi38 := 0; fi38 < 3; fi38++ {
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    v39 := GenS5{F0: (v28 - genFn0("foo bar", uint16(v1), (("foo bar" + ("hello" + ("foo bar" + ((("" + "\n\t") + "hello") + "a")))) + "a")))}
                    _ = v39
                default:
                    _ = len(v5)
                }
                _ = fi38
            }
            var v40 float64 = v25
            var v41 uint = 560
            _ = fi37
            _ = v40
            _ = v41
        }
    }
    if true {
        var v42 uint64 = 516
        for fi43 := 0; fi43 < 8; fi43++ {
            switch (fi43 % (((fi43 + fi43)) | 1)) {
            case -1:
                v44 := GenS6{F0: false, F1: -78, F2: (65535 != v1)}
                _ = v44
            }
            v42 = (v42 * (v42 % ((v28) | 1)))
            _ = fi43
        }
        if (true || (false || true)) {
            if false {
                if true {
                    var v45 int16 = 0
                    v42 -= (v28 << ((883) & 63))
                    v2["a"] = ""
                    _ = v45
                } else {
                    for fi46 := 0; fi46 < 9; fi46++ {
                        v28 -= (v42 - 269)
                        v47 := GenS7{F0: 0}
                        for fi48 := 0; fi48 < 11; fi48++ {
                            if true {
                                v3 += (v3 >> ((18446744073709551615) & 63))
                            }
                            v49 := GenS8{F0: "", F1: true}
                            _ = fi48
                            _ = v49
                        }
                        _ = fi46
                        _ = v47
                    }
                    switch 602 {
                    case 32767:
                        var v50 int16 = -1
                        _ = v50
                    case -2147483648:
                        for fi51 := 0; fi51 < 16; fi51++ {
                            switch 2147483647 {
                            case 42:
                                v4 = append(v4, (v1 + (v1 | (v1 * (v1 ^ 959)))))
                            }
                            _ = fi51
                        }
                    default:
                        v26 = 693
                    }
                }
                v52 := GenS9{F0: 32767, F1: (v26 / (((v26 | 2)) | 1))}
                _ = v52
            } else {
                for fi53 := 0; fi53 < 19; fi53++ {
                    var v54 int32 = -1
                    if (true && false) {
                        v42 -= (^(v42))
                        v55 := GenS10{F0: 1, F1: (v54 / ((v54) | 1)), F2: -1}
                        fmt.Println(v3, v54)
                        _ = v55
                    }
                    _ = fi53
                    _ = v54
                }
                var v56 int16 = -32768
                _ = v56
            }
            if (!false) {
                v28 -= v28
                v25 = (0.0 + float64((float32(0.0) * 1.5)))
                v26 = v26
            } else {
                var v57 float64 = v25
                _ = v57
            }
            v26 += v26
        }
        _ = v42
    }
    v58 := GenS11{F0: v1}
    fmt.Println(v25, v25, v1)
    v25 += (v25 + ((float64(3.14159) + (v25 + float64((1e10 * 1.5)))) - (v25 / 1.5)))
    var v59 uint = 29
    var v60 string = ("a" + ("a" + "\n\t"))
    fmt.Println(v26, v59)
    for fi61 := 0; fi61 < 11; fi61++ {
        if true {
            var v62 int = -9223372036854775808
            var v63 int64 = v3
            _ = v62
            _ = v63
        }
        switch -2147483648 {
        case -2147483648:
            var v64 uint = (v59 >> ((1) & 63))
            _ = v64
        }
        var v65 uint16 = (^(v26))
        _ = fi61
        _ = v65
    }
    v26 = 447
    for fi66 := 0; fi66 < 4; fi66++ {
        fmt.Println(v28, v28, v26)
        var v67 bool = (-1 >= 1)
        _ = fi66
        _ = v67
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v60
}
