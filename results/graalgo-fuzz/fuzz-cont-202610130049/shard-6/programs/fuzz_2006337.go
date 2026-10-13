package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 uint64 }
type GenS6 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS7 struct { F0 int8; F1 uint16; F2 bool }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]uint32)
    var v3 bool = false
    delete(v2, 0)
    var v4 float64 = ((float64(((1e10 + 3.14159) - 0.0)) + ((-2.25 * 3.14159) - 3.14159)) / ((float64(0.1) + (1.5 / -2.25)) + (-2.25 / 1.5)))
    var v5 int8 = 127
    var v6 uint16 = 473
    var v7 uint16 = v6
    v6 += (v6 + (v6 | (v7 * v7)))
    var v8 int = -91
    switch 32768 {
    case 128:
        v3 = (v7 <= v7)
    case 245:
        if ((1 != v7) && v3) {
            if true {
                if v3 {
                    if v3 {
                        var v9 int = v8
                        for fi10 := 0; fi10 < 9; fi10++ {
                            v11 := GenS1{F0: 42, F1: 716, F2: (v6 << ((65535) & 63))}
                            if (((-2147483648 == (v9 - 4294967295)) && v3) && (v8 >= v8)) {
                                if v3 {
                                    v12 := GenS2{F0: (("" + "foo bar") + "hello"), F1: 0}
                                    _ = v12
                                }
                                v13 := GenS2{F0: (("a" + (("foo bar" + "a") + ((("\n\t" + "foo bar") + "a") + "foo bar"))) + "a"), F1: 9223372036854775807}
                                _ = v13
                            }
                            _ = fi10
                            _ = v11
                        }
                        _ = v9
                    } else {
                        var v14 uint8 = 128
                        var v15 int = v8
                        if v3 {
                            if ((v6 << ((576) & 63)) >= v6) {
                                v16 := GenS3{F0: (^(v6))}
                                var v17 uint = 255
                                _ = v16
                                _ = v17
                            } else {
                                fmt.Println(v6)
                            }
                            fmt.Println(v7, v5, v15)
                        } else {
                            if v3 {
                                fmt.Println(v14)
                                v18 := []int{(v15 * (v15 >> ((1) & 63))), v15}
                                _ = v18
                            } else {
                                if v3 {
                                    _ = len(v2)
                                }
                                fmt.Println(v15, v5, v14)
                                v19 := []int{-1}
                                _ = v19
                            }
                        }
                        _ = v14
                        _ = v15
                    }
                    for fi20 := 0; fi20 < 13; fi20++ {
                        var v21 uint16 = (^(v7))
                        _ = fi20
                        _ = v21
                    }
                } else {
                    v7 -= (v7 & v6)
                }
                fmt.Println(v5, v6)
            } else {
                switch (v8 | (v8 + (-(v8)))) {
                case 1:
                    switch ((("foo bar" + (("a" + "a") + "\n\t")) + "") + "foo bar") {
                    case "":
                        for fi22 := 0; fi22 < 14; fi22++ {
                            v6 += (^(v6))
                            fmt.Println(v6)
                            _ = fi22
                        }
                    case "foo bar":
                        var v23 int = (v8 % ((v8) | 1))
                        _ = v23
                    case "hello":
                        v24 := GenS3{F0: v6}
                        _ = v24
                    default:
                        switch (-(v8)) {
                        case -1:
                            switch 128 {
                            case 32767:
                                v25 := []int{}
                                _ = v25
                            case 32768:
                                var v26 int8 = (v5 | v5)
                                _ = v26
                            case -935:
                                if false {
                                    v27 := GenS4{F0: 4294967295}
                                    if (42 < -1) {
                                        v4 = (0.0 / ((0.0 + ((((v4 * (v4 - v4)) - ((v4 * (float64(1e10) - float64(3.14159))) - (v4 / (float64((0.1 - ((0.1 / -2.25) / float32((float64(3.14159) + v4))))) - 1.5)))) * 0.1) - 1e10)) - v4))
                                        v5 += -128
                                    }
                                    _ = v27
                                } else {
                                    v6 -= 454
                                }
                            default:
                                v7 = v7
                            }
                        }
                    }
                case 32768:
                    for fi28 := 0; fi28 < 16; fi28++ {
                        v29 := []int{(fi28 % (((v8 + 42)) | 1))}
                        v1[fi28] = -2147483648
                        v30 := GenS1{F0: 0, F1: 18446744073709551615, F2: (v6 % ((v7) | 1))}
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                default:
                    if (v7 <= (v6 + (^(v7)))) {
                        fmt.Println(v7)
                        var v31 uint64 = 570
                        fmt.Println(v6, v7)
                        _ = v31
                    } else {
                        var v32 uint64 = 784
                        v33 := []int{v8, v8, v8}
                        _ = v32
                        _ = v33
                    }
                }
                v34 := GenS5{F0: v5, F1: 42}
                if (((0 >= 505) || false) || v3) {
                    v35 := GenS6{F0: 128, F1: -942, F2: 63}
                    fmt.Println(v8)
                    var v36 string = "\n\t"
                    _ = v35
                    _ = v36
                } else {
                    v8 = -2147483648
                    v37 := []int{(v8 >> ((1) & 63)), (v8 ^ 127), -2147483648, int(v5)}
                    var v38 uint32 = 600
                    _ = v37
                    _ = v38
                }
                _ = v34
            }
        } else {
            for fi39 := 0; fi39 < 3; fi39++ {
                _ = len(v2)
                v5 = (-(v5))
                v40 := v2[65535]
                _ = fi39
                _ = v40
            }
        }
    case 127:
        v41 := GenS3{F0: (v7 ^ v6)}
        _ = v41
    }
    _ = len(v2)
    v42 := v2[2]
    v43 := GenS7{F0: 100, F1: 822, F2: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
