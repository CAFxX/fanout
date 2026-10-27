package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 bool; F2 uint }
type GenS4 struct { F0 int64; F1 string }
func genFn0() bool {
    for fi6 := 0; fi6 < 13; fi6++ {
        if (!(-113 < 127)) {
            if (!((!true) && false)) {
                for fi7 := 0; fi7 < 1; fi7++ {
                    v8 := []int{(fi7 & (fi6 ^ fi6)), (-(fi6)), (fi7 * fi6)}
                    _ = fi7
                    _ = v8
                }
            } else {
                var v9 int64 = 9223372036854775807
                _ = v9
            }
            var v10 string = "a"
            v11 := []int{}
            _ = v10
            _ = v11
        }
        if (!true) {
            if (true || (1 != -32768)) {
                for fi12 := 0; fi12 < 7; fi12++ {
                    v13 := []int{}
                    _ = fi12
                    _ = v13
                }
            }
            if true {
                var v14 int8 = 100
                var v15 int = (fi6 - fi6)
                _ = v14
                _ = v15
            }
        } else {
            switch "\n\t" {
            case "\n\t":
                var v16 float32 = (float32(0.1) / float32(((0.1 + 0.1) / (1e10 + 3.14159))))
                _ = v16
            case "":
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 uint64 = 1
                    if ((v18 - v18) <= v18) {
                        switch 42 {
                        case 42:
                            v19 := GenS1{F0: -540}
                            _ = v19
                        }
                    } else {
                        if (!false) {
                            switch "a" {
                            case "\n\t":
                                v20 := []int{(fi6 ^ fi17)}
                                _ = v20
                            case "foo bar":
                                v18 += v18
                            default:
                                switch ("hello" + "foo bar") {
                                case "":
                                    v18 -= (v18 << ((255) & 63))
                                default:
                                    var v21 float64 = (float64(3.14159) + float64((float32(3.14159) + ((0.1 + (1e10 + -2.25)) - ((-2.25 * ((-2.25 - -2.25) * (1e10 / (((float32((float64((-2.25 * 1e10)) + 3.14159)) / (float32(float64(((0.0 + (-2.25 - (3.14159 * (float32(1.5) * (3.14159 + 0.1))))) * (1e10 + (1e10 + 1e10))))) * 1.5)) * 1.5) / (1e10 / (3.14159 - (1.5 + 0.0))))))) + 0.0)))))
                                    var v22 int64 = -1
                                    _ = v21
                                    _ = v22
                                }
                            }
                        }
                        v18 = v18
                    }
                    v23 := []int{}
                    _ = fi17
                    _ = v18
                    _ = v23
                }
            default:
                switch "foo bar" {
                case "\n\t":
                    var v24 uint = 2
                    _ = v24
                case "foo bar":
                    v25 := []int{fi6, (^(fi6)), (^(fi6))}
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 float32 = ((0.0 - 1e10) + float32((((((0.0 / 1e10) / 3.14159) * ((1e10 + 0.1) / (0.1 / (1e10 * ((float64(0.0) + -2.25) - 0.0))))) / (3.14159 * 3.14159)) / -2.25)))
                _ = fi26
                _ = v27
            }
        }
        _ = fi6
    }
    var v28 int = 42
    v28 += (v28 >> ((4294967295) & 63))
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := GenS1{F0: 832}
        fmt.Println(v28)
        _ = fi29
        _ = v30
    }
    _ = v28
    return (100 != 1)
}
func genFn1() uint8 {
    switch 262 {
    case -1:
        v31 := []int{-749}
        _ = v31
    case 1:
        v32 := []int{9223372036854775807, 32768, 32767}
        _ = v32
    }
    var v33 bool = true
    _ = v33
    return 2
}

func main() {
    var v1 bool = (114 > 65535)
    var v2 uint32 = 447
    var v3 int64 = -1
    var v4 int32 = 0
    v5 := make(map[int]string)
    var v34 uint16 = 69
    switch "a" {
    case "foo bar":
        v35 := GenS2{F0: v3}
        _ = v35
    default:
        fmt.Println(v4, v3)
    }
    v1 = (v2 > v2)
    var v36 uint8 = 107
    for fi37 := 0; fi37 < 9; fi37++ {
        for fi38 := 0; fi38 < 16; fi38++ {
            delete(v5, fi37)
            _ = fi38
        }
        var v39 int8 = int8(v4)
        switch (fi37 * int(v34)) {
        case 32767:
            fmt.Println(v3, v39)
        case 1:
            v5[fi37] = ""
        case 42:
            v40 := GenS3{F0: (v4 | -1), F1: (false && true), F2: 255}
            _ = v40
        default:
            v39 -= v39
        }
        _ = fi37
        _ = v39
    }
    if (18446744073709551615 < 2) {
        fmt.Println(v34, v2)
    }
    switch 42 {
    case -9223372036854775808:
        fmt.Println(v4, v1, v2)
    case 32767:
        v41 := []int{}
        _ = v41
    default:
        v42 := GenS4{F0: v3, F1: ""}
        _ = v42
    }
    var v43 uint = 4294967295
    fmt.Println(v36, v43, v34)
    if (0 >= 32767) {
        fmt.Println(v43)
    }
    _ = len(v5)
    v1 = v1
    v4 -= (v4 - (v4 / ((0) | 1)))
    _ = len(v5)
    for fi44 := 0; fi44 < 11; fi44++ {
        v45 := GenS4{F0: v3, F1: (("\n\t" + ("" + "a")) + "hello")}
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v4
    _ = v43
    _ = v5
}
