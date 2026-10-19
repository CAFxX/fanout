package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int16; F1 int8; F2 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint8 }
func genFn0() uint {
    if false {
        v6 := []int{-165, -1}
        v7 := GenS1{F0: 13, F1: 384}
        switch 32767 {
        case 9223372036854775807:
            var v8 int64 = 0
            _ = v8
        }
        _ = v6
        _ = v7
    }
    v9 := []int{0}
    var v10 float64 = ((0.0 / 0.1) / 3.14159)
    _ = len(v9)
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 int64 = 42
    var v2 int64 = v1
    v3 := []string{"\n\t"}
    var v4 int16 = 1
    var v5 bool = true
    fmt.Println(v5, v4)
    if v5 {
        v5 = (143 == 42)
        v2 -= v1
    } else {
        v4 -= 1000
    }
    var v11 bool = (v4 != (v4 / (((v4 / ((v4) | 1))) | 1)))
    if v11 {
        for fi12 := 0; fi12 < 15; fi12++ {
            v4 = (v4 * (v4 - (v4 & int16(fi12))))
            v13 := GenS2{F0: (v4 >> ((2) & 63)), F1: 127, F2: 1}
            for fi14 := 0; fi14 < 5; fi14++ {
                if (fi14 < (fi14 % ((-9223372036854775808) | 1))) {
                    v11 = v5
                    v15 := GenS1{F0: 42, F1: fi14}
                    _ = len(v3)
                    _ = v15
                } else {
                    v11 = false
                    fmt.Println(v4, v11)
                    v5 = v11
                }
                _ = fi14
            }
            _ = fi12
            _ = v13
        }
        var v16 float64 = 0.0
        _ = v16
    } else {
        for fi17 := 0; fi17 < 9; fi17++ {
            if false {
                if v5 {
                    v4 -= 1
                    for fi18 := 0; fi18 < 3; fi18++ {
                        var v19 float32 = 3.14159
                        _ = len(v3)
                        if (("foo bar" + "foo bar") == (("\n\t" + "") + ("" + (("foo bar" + "a") + "a")))) {
                            fmt.Println(fi18, v11, v19)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch (^(fi18)) {
                            case -1:
                                v5 = v11
                            case 9223372036854775807:
                                switch fi17 {
                                case 42:
                                    v20 := GenS2{F0: 1000, F1: 100, F2: 2}
                                    _ = v20
                                case 987:
                                    var v21 uint32 = 42
                                    _ = v21
                                case 0:
                                    v22 := GenS2{F0: 0, F1: 100, F2: 279}
                                    _ = v22
                                default:
                                    v1 += v1
                                }
                            case 32768:
                                v19 += ((0.1 / (3.14159 + ((1.5 - v19) + (v19 + v19)))) / v19)
                            }
                            if len(v3) > 0 {
                                v3[0] = ("a" + ("foo bar" + "foo bar"))
                            }
                        } else {
                            v23 := GenS3{F0: fi17}
                            var v24 int8 = -1
                            var v25 uint8 = 128
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        _ = fi18
                        _ = v19
                    }
                } else {
                    switch "" {
                    case "\n\t":
                        for fi26 := 0; fi26 < 8; fi26++ {
                            v2 -= 185
                            v27 := GenS4{F0: 1, F1: 42, F2: 127}
                            v3 = append(v3, "")
                            _ = fi26
                            _ = v27
                        }
                    case "foo bar":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
                fmt.Println(fi17, v11, v5)
            }
            _ = fi17
        }
        fmt.Println(v2)
        var v28 float64 = -2.25
        _ = v28
    }
    v4 -= 1
    switch 128 {
    case -378:
        v2 += v2
    case 1:
        v29 := GenS4{F0: 2147483647, F1: 18446744073709551615, F2: 255}
        _ = v29
    default:
        fmt.Println(v5, v2, v4)
    }
    v30 := GenS4{F0: 4294967295, F1: 533, F2: 201}
    var v31 int16 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
