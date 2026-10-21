package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int; F1 uint8; F2 uint }
func genFn0() int16 {
    if true {
        if (0 <= 32767) {
            v6 := []int{}
            if (65535 == 42) {
                var v7 bool = true
                if len(v6) > 0 {
                    v6[0] = -9223372036854775808
                }
                if ((3.14159 / 1.0) <= float32(float64(0.0))) {
                    switch 709 {
                    case 32767:
                        if v7 {
                            if true {
                                var v8 float64 = (((float64((((0.1 + float32(((0.1 * 1e10) - -2.25))) - ((3.14159 + (1e10 / (1.5 - (0.0 - 0.1)))) + (((3.14159 / float32(-2.25)) / -2.25) * 0.1))) * 0.0)) / float64((1e10 - 3.14159))) + 3.14159) / (1.5 + 1.5))
                                var v9 uint32 = 1
                                var v10 float64 = v8
                                _ = v10
                                _ = v8
                                _ = v9
                            } else {
                                switch 126 {
                                case 9223372036854775807:
                                    fmt.Println(v7)
                                }
                                v7 = (0 != -135)
                            }
                            if false {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v11 := GenS1{F0: -9223372036854775808}
                                for fi12 := 0; fi12 < 16; fi12++ {
                                    v7 = v7
                                    _ = fi12
                                }
                                _ = v11
                            } else {
                                for fi13 := 0; fi13 < 11; fi13++ {
                                    if (v7 || true) {
                                        if false {
                                            v7 = (!(!v7))
                                        }
                                        v7 = (v7 && (32768 != fi13))
                                    } else {
                                        fmt.Println(v7, v7)
                                        if (18446744073709551615 < 42) {
                                            v14 := GenS2{F0: "hello"}
                                            _ = v14
                                        } else {
                                            fmt.Println(fi13, fi13)
                                        }
                                    }
                                    v7 = (1 == int16(fi13))
                                    v7 = v7
                                    _ = fi13
                                }
                            }
                        }
                    case -9223372036854775808:
                        var v15 bool = (113 == -98)
                        _ = v15
                    }
                }
                _ = v7
            }
            var v16 string = "\n\t"
            _ = v16
            _ = v6
        }
        var v17 int = 0
        var v18 float32 = -2.25
        _ = v17
        _ = v18
    }
    v19 := GenS3{F0: -2147483648, F1: 42, F2: -937}
    _ = v19
    return -32768
}
func genFn1(p0 uint16, p1 uint8, p2 int) int8 {
    fmt.Println(p2, p1, p2)
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS4{F0: 2, F1: 365}
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v23 := GenS5{F0: fi22, F1: (p1 & 42), F2: 255}
        _ = fi22
        _ = v23
    }
    fmt.Println(p2)
    var v24 uint64 = 396
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    return 42
}

func main() {
    var v1 uint16 = 69
    var v2 uint32 = 2147483647
    var v3 uint = uint(v2)
    v4 := make(map[string]int)
    v5 := make(map[uint32]uint32)
    var v25 int32 = 1
    var v26 uint16 = (^(v1))
    v25 -= 42
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 bool = (!(2147483647 > v2))
        var v29 int32 = (-(v25))
        if ((v2 ^ 707) == 28) {
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 int32 = (v25 - 42)
                v32 := []int{}
                _ = fi30
                _ = v31
                _ = v32
            }
            if (!v28) {
                if (!(v28 || (!v28))) {
                    switch ((("hello" + "\n\t") + ("\n\t" + ("a" + ""))) + ("hello" + "hello")) {
                    case "\n\t":
                        v5[2147483648] = v2
                    case "foo bar":
                        if v28 {
                            v28 = (!v28)
                            for fi33 := 0; fi33 < 9; fi33++ {
                                v34 := GenS2{F0: "hello"}
                                var v35 float64 = -2.25
                                var v36 uint16 = (v1 / ((255) | 1))
                                _ = fi33
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            v25 += 1
                        } else {
                            for fi37 := 0; fi37 < 5; fi37++ {
                                v38 := []int{(fi37 >> ((v3) & 63)), (fi27 >> ((843) & 63)), 1, 127}
                                _ = fi37
                                _ = v38
                            }
                            switch (fi27 - fi27) {
                            case 1:
                                for fi39 := 0; fi39 < 13; fi39++ {
                                    v1 = (v26 | (v26 * 0))
                                    _ = fi39
                                }
                            }
                        }
                    case "hello":
                        v28 = ((genFn0() == 42) || ((^(v2)) != v2))
                    default:
                        v40 := GenS4{F0: (v2 | 434), F1: -1}
                        _ = v40
                    }
                    fmt.Println(v26)
                }
            }
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    var v41 uint16 = v26
    if true {
        v42 := []int{-2147483648, -9223372036854775808, -1}
        _ = v42
    }
    v26 = v26
    var v43 bool = (!true)
    for fi44 := 0; fi44 < 17; fi44++ {
        var v45 uint32 = v2
        var v46 uint16 = 0
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v5
}
