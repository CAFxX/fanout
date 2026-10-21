package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int; F2 uint16 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0() string {
    var v5 uint64 = 26
    v6 := []int{537, -344, -562}
    v5 -= v5
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int32 = 0
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return ""
}
func genFn1(p0 string, p1 uint16) string {
    p1 = (p1 / (((p1 / ((p1) | 1))) | 1))
    if (973 == 707) {
        p0 = p0
    }
    if false {
        var v9 float64 = (float64(3.14159) / float64((0.0 - float32(-2.25))))
        var v10 int = 1
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    v11 := GenS1{F0: (!(("\n\t" > "\n\t") && true))}
    switch 127 {
    case 0:
        v12 := GenS1{F0: true}
        _ = v12
    case -9223372036854775808:
        if (false || ("a" >= "")) {
            var v13 uint32 = 74
            if (!(32768 == 128)) {
                v14 := GenS1{F0: false}
                _ = v14
            } else {
                var v15 string = "hello"
                switch 42 {
                case 4294967295:
                    var v16 uint8 = 127
                    _ = v16
                case -786:
                    var v17 int = 791
                    _ = v17
                case 0:
                    v18 := []int{0, -1, 32768}
                    _ = v18
                default:
                    var v19 int = -9223372036854775808
                    _ = v19
                }
                for fi20 := 0; fi20 < 17; fi20++ {
                    for fi21 := 0; fi21 < 7; fi21++ {
                        if false {
                            var v22 float32 = -2.25
                            v23 := GenS1{F0: (!false)}
                            _ = v22
                            _ = v23
                        }
                        _ = fi21
                    }
                    v24 := GenS1{F0: true}
                    _ = fi20
                    _ = v24
                }
                _ = v15
            }
            v4[128] = (("a" + "") + "foo bar")
            _ = v13
        }
    case -231:
        v25 := []int{42, -888}
        _ = v25
    default:
        v26 := GenS2{F0: 61, F1: -9223372036854775808, F2: 455}
        _ = v26
    }
    var v27 int8 = 100
    switch "a" {
    case "\n\t":
        v28 := GenS3{F0: 42, F1: 1000}
        _ = v28
    case "hello":
        for fi29 := 0; fi29 < 5; fi29++ {
            fmt.Println(fi29)
            v30 := GenS1{F0: true}
            v31 := v4[(-(fi29))]
            _ = fi29
            _ = v30
            _ = v31
        }
    }
    fmt.Println(v27)
    switch "" {
    case "":
        var v32 int = 2147483647
        _ = v32
    default:
        if ((("a" + ("\n\t" + "hello")) >= "foo bar") && (924 < 4294967295)) {
            v27 -= 127
            for fi33 := 0; fi33 < 13; fi33++ {
                switch fi33 {
                case 128:
                    v27 = v27
                case 32768:
                    var v34 int8 = v27
                    _ = v34
                default:
                    switch "hello" {
                    case "\n\t":
                        var v35 uint8 = 234
                        _ = v35
                    default:
                        switch (-(fi33)) {
                        case 496:
                            for fi36 := 0; fi36 < 13; fi36++ {
                                if (v27 <= (v27 - 42)) {
                                    switch fi36 {
                                    case -166:
                                        v27 += (^(v27))
                                    case -8:
                                        if true {
                                            switch 4294967295 {
                                            case 32768:
                                                fmt.Println(v27)
                                            }
                                            var v37 float32 = ((3.14159 - 0.0) + (-2.25 - 0.1))
                                            _ = v37
                                        }
                                    case 728:
                                        v38 := []int{2147483647, (fi36 >> ((18446744073709551615) & 63)), 4294967295, fi33}
                                        _ = v38
                                    default:
                                        if true {
                                            var v39 uint16 = 2
                                            fmt.Println(v27, fi33)
                                            _ = v39
                                        } else {
                                            v27 += v27
                                            fmt.Println(v27)
                                        }
                                    }
                                    v27 -= (v27 << ((695) & 63))
                                } else {
                                    v40 := GenS2{F0: 128, F1: fi33, F2: 1}
                                    var v41 uint16 = 576
                                    _ = v40
                                    _ = v41
                                }
                                var v42 uint32 = 4294967295
                                _ = fi36
                                _ = v42
                            }
                        case 2147483647:
                            if (!(!(65535 < 1))) {
                                switch "a" {
                                case "hello":
                                    v43 := GenS2{F0: 42, F1: (fi33 ^ fi33), F2: 0}
                                    _ = v43
                                case "a":
                                    fmt.Println(fi33)
                                case "\n\t":
                                    fmt.Println(fi33, v27)
                                default:
                                    var v44 int = 1
                                    _ = v44
                                }
                                v27 += v27
                            }
                        default:
                            switch "a" {
                            case "a":
                                v27 = v27
                            case "":
                                delete(v4, fi33)
                            }
                        }
                    }
                }
                _ = fi33
            }
        }
    }
    var v45 uint16 = 0
    fmt.Println(v27, v45)
    fmt.Println(v45, v45)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v45
}
