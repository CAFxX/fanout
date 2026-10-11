package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 bool; F2 int }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS5 struct { F0 int32; F1 uint8 }
func genFn0(p0 int32) int32 {
    if false {
        p0 = (p0 - p0)
        v5 := []int{108, 32767, 128}
        _ = v5
    }
    switch "a" {
    case "foo bar":
        for fi6 := 0; fi6 < 14; fi6++ {
            p0 = (p0 | (p0 % ((-2147483648) | 1)))
            _ = fi6
        }
    case "\n\t":
        p0 -= 42
    case "a":
        var v7 uint64 = 42
        _ = v7
    default:
        v8 := GenS1{F0: 1, F1: 32767}
        _ = v8
    }
    _ = p0
    return (p0 / ((-1) | 1))
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = ((0.1 - ((1.5 / (3.14159 / (3.14159 * 1e10))) + (3.14159 + (1e10 / 3.14159)))) * (3.14159 + -2.25))
    v3 := make(map[string]int)
    var v4 int = 32767
    fmt.Println(v2, v2)
    v2 = v2
    for fi9 := 0; fi9 < 9; fi9++ {
        v2 -= v2
        var v10 int = 128
        v4 = -571
        _ = fi9
        _ = v10
    }
    switch v4 {
    case 932:
        if false {
            var v11 uint32 = 2147483647
            var v12 int16 = 42
            _ = v11
            _ = v12
        }
    case 32767:
        var v13 uint8 = 128
        _ = v13
    }
    fmt.Println(v2)
    if true {
        v14 := GenS1{F0: 18446744073709551615, F1: -32768}
        v1[(("hello" + "a") + ("foo bar" + ("" + "hello")))] = 937
        _ = v14
    }
    if (128 != 71) {
        v15 := []int{(v4 / (((v4 / (((-(v4))) | 1))) | 1)), (v4 / ((v4) | 1)), v4, (v4 + v4)}
        v4 = v4
        if true {
            if len(v15) >= 2 {
                v15 = v15[1:len(v15)]
            }
        } else {
            var v16 int64 = -1
            _ = v16
        }
        _ = v15
    }
    switch "" {
    case "hello":
        switch "\n\t" {
        case "hello":
            switch "hello" {
            case "a":
                v17 := GenS1{F0: 200, F1: 1000}
                _ = v17
            case "":
                var v18 int64 = 0
                _ = v18
            default:
                var v19 int32 = 419
                _ = v19
            }
        }
    }
    if false {
        if true {
            v20 := GenS1{F0: 18446744073709551615, F1: 951}
            fmt.Println(v4, v4)
            var v21 float64 = 0.1
            _ = v20
            _ = v21
        }
        if false {
            fmt.Println(v2, v4)
            fmt.Println(v4, v2)
            for fi22 := 0; fi22 < 9; fi22++ {
                if (true || true) {
                    var v23 float32 = ((v2 / v2) - -2.25)
                    v24 := []int{}
                    _ = len(v1)
                    _ = v23
                    _ = v24
                } else {
                    v25 := GenS2{F0: 18446744073709551615}
                    v2 = float32((1.5 - -2.25))
                    _ = v25
                }
                fmt.Println(v2, fi22, v2)
                _ = fi22
            }
        } else {
            switch "" {
            case "foo bar":
                v2 += 3.14159
            case "hello":
                for fi26 := 0; fi26 < 11; fi26++ {
                    switch fi26 {
                    case 32767:
                        v4 -= (fi26 + -2147483648)
                    default:
                        v27 := GenS2{F0: 42}
                        _ = v27
                    }
                    v28 := GenS3{F0: 4294967295, F1: false, F2: (v4 * (fi26 / ((v4) | 1)))}
                    switch (("" + "foo bar") + "\n\t") {
                    case "foo bar":
                        v4 += v4
                    case "":
                        for fi29 := 0; fi29 < 17; fi29++ {
                            if ((630 <= 167) || false) {
                                if false {
                                    v4 -= (v4 ^ (v4 ^ 42))
                                    var v30 float64 = (0.0 * 1.5)
                                    var v31 int16 = 880
                                    var v32 float32 = 0.1
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                } else {
                                    var v33 float32 = (v2 + float32(1e10))
                                    fmt.Println(fi29, fi29, v2)
                                    _ = v33
                                }
                            } else {
                                switch (((("hello" + "foo bar") + "foo bar") + ("hello" + "hello")) + ("\n\t" + ("foo bar" + ("\n\t" + "\n\t")))) {
                                case "":
                                    v34 := GenS4{F0: 62, F1: 42, F2: 32767}
                                    _ = v34
                                }
                                v35 := GenS5{F0: 1, F1: 128}
                                var v36 uint64 = 42
                                _ = v35
                                _ = v36
                            }
                            v2 -= float32(-2.25)
                            _ = fi29
                        }
                    case "\n\t":
                        v4 -= v4
                    }
                    _ = fi26
                    _ = v28
                }
            default:
                v37 := []int{1, (-(v4)), v4, 320}
                _ = v37
            }
        }
    } else {
        var v38 int = (^(v4))
        var v39 uint = 690
        v40 := GenS4{F0: 1, F1: 2, F2: 42}
        _ = v38
        _ = v39
        _ = v40
    }
    fmt.Println(v4)
    v41 := v3[("hello" + "\n\t")]
    v42 := []int{}
    switch "a" {
    case "foo bar":
        v43 := GenS2{F0: 18446744073709551615}
        _ = v43
    }
    for fi44 := 0; fi44 < 11; fi44++ {
        v45 := GenS4{F0: 2, F1: 2147483648, F2: 0}
        _ = fi44
        _ = v45
    }
    var v46 bool = false
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v46
}
