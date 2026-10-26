package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int16 = 42
    var v2 uint16 = 1
    var v3 float64 = float64(1e10)
    var v4 int8 = -1
    var v5 float32 = float32(1e10)
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 float32 = v5
        var v8 float64 = v3
        _ = fi6
        _ = v7
        _ = v8
    }
    if (false && (128 >= 255)) {
        v9 := GenS1{F0: -652, F1: 0}
        v10 := GenS1{F0: 1, F1: -1}
        _ = v10
        _ = v9
    }
    if (false || true) {
        fmt.Println(v4, v5, v5)
        var v11 uint16 = 32767
        _ = v11
    }
    v12 := GenS1{F0: 236, F1: 0}
    fmt.Println(v3)
    var v13 float64 = v3
    switch "hello" {
    case "a":
        v2 -= v2
    }
    v4 = -20
    if true {
        for fi14 := 0; fi14 < 7; fi14++ {
            for fi15 := 0; fi15 < 2; fi15++ {
                var v16 int16 = v1
                _ = fi15
                _ = v16
            }
            if true {
                if (18446744073709551615 < 18446744073709551615) {
                    for fi17 := 0; fi17 < 9; fi17++ {
                        var v18 uint8 = 1
                        v18 = (v18 / ((v18) | 1))
                        _ = fi17
                        _ = v18
                    }
                    v2 -= v2
                }
                if (379 <= 139) {
                    if (v4 > (v4 ^ (v4 - v4))) {
                        for fi19 := 0; fi19 < 11; fi19++ {
                            var v20 int = fi19
                            if true {
                                v21 := GenS2{F0: v4}
                                if ((1 != 60) && (!false)) {
                                    switch ((("hello" + ("foo bar" + (("a" + ("a" + ("foo bar" + ("a" + ("\n\t" + "hello"))))) + "foo bar"))) + "hello") + "hello") {
                                    case "foo bar":
                                        fmt.Println(v1, v4, v1)
                                    case "":
                                        switch fi14 {
                                        case 32767:
                                            v13 -= 1e10
                                        case 202:
                                            fmt.Println(fi19, v3)
                                            v22 := []int{fi14, (fi14 >> ((18446744073709551615) & 63)), (^(fi19))}
                                            _ = v22
                                        case -983:
                                            var v23 uint8 = 173
                                            _ = v23
                                        default:
                                            v5 = ((((((v5 * v5) + ((((v5 + v5) + float32(((-2.25 * v3) + v3))) * ((-2.25 / (float32(1.5) / (v5 - float32((float64(v5) + v13))))) - v5)) - (float32(v3) / v5))) * -2.25) - (((1e10 - 0.1) * v5) + float32(float64(v5)))) + v5) - (1.5 + v5))
                                        }
                                    case "hello":
                                        var v24 float32 = (v5 - v5)
                                        _ = v24
                                    default:
                                        switch (^(fi19)) {
                                        case 9223372036854775807:
                                            switch "foo bar" {
                                            case "hello":
                                                if (true && (v4 > (^(v4)))) {
                                                    fmt.Println(v5, v2, v20)
                                                    var v25 int32 = 2147483647
                                                    if (!false) {
                                                        v25 = (^(v25))
                                                        fmt.Println(v20)
                                                    }
                                                    _ = v25
                                                }
                                            case "\n\t":
                                                var v26 int16 = v1
                                                _ = v26
                                            case "foo bar":
                                                v1 += -502
                                            default:
                                                fmt.Println(fi19, fi19)
                                            }
                                        case -84:
                                            var v27 uint8 = 142
                                            _ = v27
                                        case 797:
                                            v13 = float64(((v5 - v5) / 1e10))
                                        default:
                                            var v28 float64 = v13
                                            _ = v28
                                        }
                                    }
                                    var v29 uint32 = 739
                                    _ = v29
                                }
                                _ = v21
                            } else {
                                v30 := []int{fi19}
                                var v31 int8 = 42
                                _ = v30
                                _ = v31
                            }
                            _ = fi19
                            _ = v20
                        }
                        var v32 uint64 = 42
                        fmt.Println(v32, fi14)
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 15; fi33++ {
                        var v34 int16 = 256
                        var v35 int16 = v34
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                } else {
                    v36 := []int{(fi14 ^ fi14), (-(fi14)), fi14, 864}
                    v2 = 53
                    var v37 uint = 255
                    _ = v36
                    _ = v37
                }
                if true {
                    v3 += ((float64(v5) / 1.0) / (v3 * -2.25))
                } else {
                    for fi38 := 0; fi38 < 16; fi38++ {
                        var v39 uint = 1
                        _ = fi38
                        _ = v39
                    }
                }
            }
            v2 += (v2 * 70)
            _ = fi14
        }
        if false {
            var v40 uint16 = (v2 % ((2) | 1))
            fmt.Println(v4)
            _ = v40
        } else {
            switch "" {
            case "hello":
                v4 += 127
            }
            v41 := GenS3{F0: 185}
            _ = v41
        }
        fmt.Println(v3, v13)
    }
    v42 := GenS4{F0: v4}
    switch (("\n\t" + "foo bar") + "hello") {
    case "\n\t":
        v43 := []int{2147483647, -1, 171, 127}
        _ = v43
    case "a":
        var v44 string = "\n\t"
        _ = v44
    case "":
        v45 := []int{32768}
        _ = v45
    default:
        var v46 int64 = 9223372036854775807
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
