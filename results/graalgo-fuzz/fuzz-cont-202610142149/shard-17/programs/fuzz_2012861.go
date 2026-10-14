package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint16; F1 int16; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 int16; F2 int32 }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint64; F1 uint; F2 int8 }
type GenS8 struct { F0 int64; F1 string; F2 int }
func genFn0(p0 uint64, p1 int64) int64 {
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: 127, F1: p0, F2: 0}
        switch fi4 {
        case 1:
            v6 := []int{(fi4 >> ((65535) & 63)), (fi4 % ((fi4) | 1))}
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 0, F1: 18446744073709551615, F2: 128}
            _ = fi7
            _ = v8
        }
        _ = fi4
        _ = v5
    }
    switch "hello" {
    case "a":
        var v9 int = 32767
        _ = v9
    case "hello":
        p1 = p1
    }
    _ = p0
    _ = p1
    return (p1 * (^(p1)))
}
func genFn1(p0 uint16, p1 int32, p2 string) uint8 {
    v10 := []int{128, 127, 9223372036854775807, -2147483648}
    v11 := GenS1{F0: 1, F1: 42, F2: 2}
    p1 += 846
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int = -9223372036854775808
    var v2 float32 = 3.14159
    v3 := make(map[string]int)
    v12 := GenS1{F0: genFn1(198, -1, ""), F1: 1, F2: 255}
    v2 -= (float32(1e10) / float32(1.5))
    switch 0 {
    case 0:
        for fi13 := 0; fi13 < 13; fi13++ {
            v2 -= (v2 - 1e10)
            fmt.Println(v2)
            _ = fi13
        }
    case 128:
        v2 -= (0.0 / (v2 - 0.1))
    case 127:
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int64 = genFn0(1, 540)
            var v16 uint8 = 1
            var v17 int32 = -132
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
    }
    var v18 uint32 = 2147483647
    v18 -= v18
    v1 = 2147483647
    fmt.Println(v1, v2, v2)
    if true {
        if false {
            for fi19 := 0; fi19 < 10; fi19++ {
                for fi20 := 0; fi20 < 19; fi20++ {
                    v2 = (-2.25 + (v2 / 1e10))
                    if true {
                        v21 := []int{}
                        if (!true) {
                            var v22 int8 = 100
                            _ = v22
                        } else {
                            if false {
                                fmt.Println(fi19)
                                if (!true) {
                                    fmt.Println(fi19, v2, v18)
                                    v2 = 1.5
                                } else {
                                    fmt.Println(v2, v2)
                                    var v23 int64 = 42
                                    _ = v23
                                }
                                v1 += (fi20 ^ fi20)
                            } else {
                                v24 := GenS1{F0: 1, F1: 883, F2: 174}
                                if (false || true) {
                                    if false {
                                        v18 += 4294967295
                                    }
                                }
                                v18 = v18
                                _ = v24
                            }
                        }
                        var v25 float64 = 0.1
                        _ = v21
                        _ = v25
                    } else {
                        var v26 int8 = -1
                        switch "a" {
                        case "":
                            if false {
                                fmt.Println(v26, v26, fi19)
                                if (true || true) {
                                    var v27 uint16 = 42
                                    _ = v27
                                }
                            } else {
                                if true {
                                    v28 := GenS2{F0: 255, F1: -128}
                                    var v29 uint8 = 211
                                    fmt.Println(v26, v2, fi20)
                                    _ = v28
                                    _ = v29
                                }
                                switch "a" {
                                case "foo bar":
                                    if (42 <= 54) {
                                        v26 = v26
                                    }
                                case "\n\t":
                                    v30 := []int{-910, 2147483647}
                                    _ = v30
                                }
                                v31 := GenS3{F0: 65535, F1: -411, F2: true}
                                v2 = -2.25
                                _ = v31
                            }
                        case "a":
                            switch "" {
                            case "":
                                if (-1 == 948) {
                                    v32 := GenS3{F0: 42, F1: -32768, F2: (v26 <= v26)}
                                    var v33 int32 = 0
                                    fmt.Println(v2, v2, v18)
                                    _ = v32
                                    _ = v33
                                }
                            case "a":
                                v18 = v18
                            }
                        case "\n\t":
                            var v34 uint8 = genFn1(436, -1, "a")
                            _ = v34
                        }
                        if true {
                            v18 -= (v18 + (^(v18)))
                            v1 = fi19
                        }
                        v2 += (1e10 / (0.1 * ((v2 / 1e10) + v2)))
                        var v35 uint32 = v18
                        _ = v26
                        _ = v35
                    }
                    _ = fi20
                }
                v36 := GenS3{F0: 42, F1: -385, F2: (!true)}
                if (v18 < v18) {
                    fmt.Println(v1, v2, v1)
                    v1 = 42
                    var v37 uint16 = 42
                    _ = v37
                }
                _ = fi19
                _ = v36
            }
            var v38 uint = 0
            v39 := GenS4{F0: v38}
            _ = v38
            _ = v39
        } else {
            switch v1 {
            case 293:
                var v40 uint8 = 128
                _ = v40
            default:
                v41 := GenS5{F0: -1, F1: 114, F2: 110}
                _ = v41
            }
        }
    } else {
        fmt.Println(v1)
        v42 := GenS6{F0: "\n\t", F1: -311}
        v18 = v18
        _ = v42
    }
    var v43 int16 = -32768
    if true {
        if true {
            if false {
                v44 := GenS7{F0: 42, F1: 207, F2: -1}
                _ = v44
            } else {
                v45 := []int{451, (-(v1)), 32768}
                var v46 int16 = v43
                _ = v45
                _ = v46
            }
        } else {
            v2 = (3.14159 / float32(float64(-2.25)))
            var v47 bool = false
            _ = v47
        }
        v1 += 9223372036854775807
    }
    for fi48 := 0; fi48 < 1; fi48++ {
        v1 += 32767
        _ = fi48
    }
    v49 := GenS8{F0: 1, F1: ("\n\t" + (("\n\t" + "hello") + ("\n\t" + ((((("" + "foo bar") + "a") + ((("\n\t" + "foo bar") + ("" + "hello")) + "foo bar")) + "\n\t") + "a")))), F2: v1}
    var v50 int64 = -735
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v43:", v43)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v43
    _ = v49
    _ = v50
}
