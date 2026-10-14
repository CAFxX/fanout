package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 string; F1 int16; F2 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint32; F1 uint8; F2 uint32 }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 int8 }
type GenS10 struct { F0 int16; F1 uint32; F2 bool }
func genFn0(p0 uint32) uint {
    v4 := GenS1{F0: 1, F1: 0}
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int64 = 556
            if true {
                fmt.Println(fi5, fi5, fi6)
                v8 := GenS1{F0: 255, F1: -128}
                v9 := GenS2{F0: 686, F1: fi6}
                v7 = (v7 << ((255) & 63))
                if (true || (!true)) {
                    switch "hello" {
                    case "\n\t":
                        v7 -= (^(v7))
                    case "foo bar":
                        fmt.Println(fi5, fi6, v7)
                    }
                }
                _ = v8
                _ = v9
            } else {
                fmt.Println(fi5, fi6)
            }
            _ = fi6
            _ = v7
        }
        if true {
            p0 -= p0
            var v10 int8 = 100
            _ = v10
        }
        _ = fi5
    }
    p0 = 204
    _ = p0
    _ = v4
    return 65535
}
func genFn1(p0 uint64) int64 {
    v11 := GenS3{F0: "foo bar", F1: 874, F2: 127}
    if ((("a" + ("foo bar" + "a")) == ("hello" + "")) || true) {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int8 = 100
            var v14 int32 = 0
            _ = fi12
            _ = v13
            _ = v14
        }
        fmt.Println(p0)
    }
    fmt.Println(p0)
    v15 := GenS4{F0: 1}
    v16 := []int{-9223372036854775808, 32768, -9223372036854775808, 42}
    _ = p0
    _ = v11
    _ = v15
    _ = v16
    return 42
}

func main() {
    var v1 float32 = float32(0.1)
    v2 := make(map[string]bool)
    var v3 float64 = -2.25
    if (true || true) {
        v17 := GenS3{F0: (("" + ((("hello" + ("\n\t" + "foo bar")) + (("\n\t" + (("foo bar" + "hello") + "hello")) + "\n\t")) + "a")) + "foo bar"), F1: 1000, F2: -128}
        switch 569 {
        case -9223372036854775808:
            for fi18 := 0; fi18 < 11; fi18++ {
                switch "" {
                case "foo bar":
                    v19 := []int{(fi18 | (fi18 & (fi18 + (fi18 * 42)))), 690, (fi18 ^ fi18)}
                    _ = v19
                case "":
                    if ((478 != 194) && (9223372036854775807 < -9223372036854775808)) {
                        switch fi18 {
                        case -239:
                            v20 := []int{}
                            _ = v20
                        }
                        if false {
                            var v21 int32 = 42
                            _ = len(v2)
                            _ = v21
                        } else {
                            v22 := GenS3{F0: "\n\t", F1: 406, F2: -1}
                            v23 := GenS5{F0: 65535, F1: 127, F2: 4294967295}
                            _ = v22
                            _ = v23
                        }
                        if true {
                            v24 := []int{}
                            v1 = float32(v3)
                            _ = v24
                        } else {
                            delete(v2, "foo bar")
                        }
                    } else {
                        if false {
                            v25 := GenS6{F0: 860, F1: 222}
                            _ = v25
                        } else {
                            for fi26 := 0; fi26 < 13; fi26++ {
                                var v27 uint8 = 2
                                v3 -= 1e10
                                if true {
                                    if true {
                                        var v28 bool = (!(42 <= 4))
                                        v1 = (v1 / 3.14159)
                                        _ = v28
                                    }
                                    var v29 uint = 0
                                    if ("" != (("" + ((("a" + "foo bar") + ("a" + "hello")) + "hello")) + ("foo bar" + "hello"))) {
                                        if true {
                                            var v30 int64 = 1
                                            v3 += (1e10 * v3)
                                            v31 := GenS2{F0: -483, F1: 128}
                                            _ = v30
                                            _ = v31
                                        }
                                        var v32 uint32 = 370
                                        _ = v32
                                    } else {
                                        fmt.Println(v29, v29)
                                        v3 -= float64((v1 - (((((float32(v3) * v1) - float32(0.1)) / float32(v3)) - v1) / ((v1 + v1) - float32(float64(v1))))))
                                    }
                                    _ = v29
                                } else {
                                    v3 -= 3.14159
                                    v3 -= 0.0
                                    fmt.Println(v27, fi26)
                                }
                                _ = fi26
                                _ = v27
                            }
                        }
                        var v33 float64 = (v3 / -2.25)
                        v34 := GenS2{F0: -498, F1: fi18}
                        _ = v33
                        _ = v34
                    }
                }
                var v35 uint8 = 225
                v36 := GenS7{F0: -47}
                _ = fi18
                _ = v35
                _ = v36
            }
        }
        _ = v17
    } else {
        if false {
            v3 -= v3
            if (834 > 0) {
                for fi37 := 0; fi37 < 15; fi37++ {
                    var v38 uint32 = 0
                    if (!true) {
                        fmt.Println(v38, v3, v1)
                        v3 = (float64(0.1) - 0.1)
                        switch (^(fi37)) {
                        case -1:
                            v1 = (v1 - -2.25)
                        case 32768:
                            var v39 int64 = 1
                            _ = v39
                        case 1:
                            v1 += (v1 + v1)
                        default:
                            v38 = (v38 << ((888) & 63))
                        }
                    }
                    _ = fi37
                    _ = v38
                }
                v40 := v2["a"]
                _ = v40
            } else {
                fmt.Println(v3, v3)
                var v41 string = "hello"
                _ = v41
            }
            if ((3.14159 <= 1e10) || (114 <= 0)) {
                _ = len(v2)
                v1 = (v1 * 1e10)
            } else {
                var v42 int32 = -2147483648
                v1 = float32(((0.1 / v3) - v3))
                _ = v42
            }
        } else {
            v43 := GenS8{F0: 2}
            _ = v43
        }
        for fi44 := 0; fi44 < 9; fi44++ {
            var v45 int64 = genFn1(339)
            var v46 int8 = -82
            _ = fi44
            _ = v45
            _ = v46
        }
        v47 := []int{-2147483648, 858, 0}
        _ = v47
    }
    v1 = -2.25
    fmt.Println(v3, v1)
    v48 := GenS9{F0: 255, F1: 127}
    fmt.Println(v3, v1)
    var v49 int64 = 0
    var v50 int32 = 1
    v51 := GenS9{F0: 31, F1: 1}
    for fi52 := 0; fi52 < 15; fi52++ {
        v53 := GenS10{F0: -32768, F1: 231, F2: true}
        _ = fi52
        _ = v53
    }
    if (true || true) {
        var v54 float32 = v1
        switch 0 {
        case -1:
            switch -2147483648 {
            case -2147483648:
                var v55 int64 = -876
                _ = v55
            case 128:
                var v56 string = "\n\t"
                _ = v56
            case 1:
                for fi57 := 0; fi57 < 14; fi57++ {
                    v54 = v54
                    for fi58 := 0; fi58 < 14; fi58++ {
                        if (!(true && ((fi58 << ((255) & 63)) == -2147483648))) {
                            v54 -= v1
                        }
                        _ = fi58
                    }
                    _ = fi57
                }
            default:
                v59 := []int{32767, 696, 9223372036854775807, 83}
                _ = v59
            }
        default:
            fmt.Println(v49, v3, v3)
        }
        _ = v54
    } else {
        v60 := GenS5{F0: 65535, F1: 51, F2: 1}
        v50 = -1
        var v61 uint = 2
        _ = v60
        _ = v61
    }
    v49 -= v49
    var v62 uint16 = 498
    var v63 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v49
    _ = v50
    _ = v51
    _ = v62
    _ = v63
}
