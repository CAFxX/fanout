package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 string }
func genFn0() int8 {
    if (42 > 42) {
        var v4 int32 = 421
        v4 += (v4 % ((v4) | 1))
        v4 -= -2147483648
        _ = v4
    }
    var v5 float32 = float32(1.5)
    fmt.Println(v5)
    _ = v5
    return 0
}

func main() {
    var v1 int8 = 66
    var v2 int = -369
    v3 := make(map[uint32]string)
    if false {
        _ = len(v3)
        v2 = 227
        v2 += 32767
    } else {
        if (("\n\t" + "a") > "") {
            if (!true) {
                var v6 int16 = 32767
                fmt.Println(v6, v1, v6)
                var v7 float32 = 0.0
                _ = v6
                _ = v7
            } else {
                for fi8 := 0; fi8 < 11; fi8++ {
                    for fi9 := 0; fi9 < 12; fi9++ {
                        fmt.Println(fi9, v2)
                        v2 = 190
                        if true {
                            if true {
                                var v10 float64 = -2.25
                                if (true && (true && true)) {
                                    v3[311] = "a"
                                    v1 += v1
                                } else {
                                    if true {
                                        v11 := GenS1{F0: int64(v1)}
                                        _ = v11
                                    } else {
                                        if (4294967295 >= 0) {
                                            var v12 int = (fi8 >> ((42) & 63))
                                            _ = v12
                                        } else {
                                            v10 -= (float64(1e10) * ((v10 * float64(1e10)) * (3.14159 - 3.14159)))
                                            switch (("\n\t" + "\n\t") + "a") {
                                            case "hello":
                                                v13 := v3[2147483647]
                                                _ = v13
                                            case "foo bar":
                                                fmt.Println(fi9, v2)
                                            case "\n\t":
                                                v14 := GenS1{F0: 9223372036854775807}
                                                switch "hello" {
                                                case "a":
                                                    v15 := []int{(fi8 >> ((4294967295) & 63)), -918, fi9, fi9}
                                                    _ = v15
                                                default:
                                                    v16 := GenS2{F0: 2, F1: false, F2: 65535}
                                                    _ = v16
                                                }
                                                v10 = v10
                                                var v17 float32 = ((float32(0.1) - 1e10) - -2.25)
                                                if (false && true) {
                                                    switch "hello" {
                                                    case "\n\t":
                                                        fmt.Println(v1)
                                                    }
                                                    v10 = v10
                                                } else {
                                                    var v18 uint64 = 0
                                                    _ = v18
                                                }
                                                _ = v14
                                                _ = v17
                                            default:
                                                v10 -= (v10 - (v10 + v10))
                                            }
                                        }
                                        if (127 < 117) {
                                            v19 := GenS3{F0: 1}
                                            switch (fi9 * (-(fi9))) {
                                            case -1:
                                                fmt.Println(v2, fi8, fi9)
                                            case -2147483648:
                                                v20 := []int{(fi9 << ((65535) & 63))}
                                                _ = v20
                                            case 302:
                                                v10 = (1.5 + (0.1 * -2.25))
                                            default:
                                                v1 -= -85
                                            }
                                            _ = v19
                                        } else {
                                            fmt.Println(fi9, v2, fi9)
                                            var v21 bool = false
                                            var v22 float32 = 1.5
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                    v10 += (v10 / v10)
                                }
                                _ = v10
                            }
                        } else {
                            v2 = -9223372036854775808
                            v2 = (fi8 >> ((255) & 63))
                        }
                        _ = fi9
                    }
                    v1 += v1
                    for fi23 := 0; fi23 < 2; fi23++ {
                        fmt.Println(fi8, fi8, fi23)
                        _ = fi23
                    }
                    _ = fi8
                }
                var v24 float32 = ((float32(0.1) / 3.14159) * (0.1 * 1e10))
                _ = v24
            }
        }
        v25 := []int{9223372036854775807, v2, (^(v2)), 0}
        for fi26 := 0; fi26 < 4; fi26++ {
            if len(v25) > 0 {
                v25[0] = v2
            }
            _ = len(v25)
            fmt.Println(v1, v1)
            _ = fi26
        }
        _ = v25
    }
    fmt.Println(v2, v1)
    var v27 uint16 = 2
    v28 := []int{v2, (v2 / ((v2) | 1)), v2}
    v1 -= -128
    for fi29 := 0; fi29 < 10; fi29++ {
        if len(v28) > 0 {
            v28[0] = -1
        }
        v27 -= v27
        _ = fi29
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 string = "a"
        _ = fi30
        _ = v31
    }
    switch ((("\n\t" + "hello") + ("" + ("foo bar" + ""))) + "hello") {
    case "hello":
        v32 := GenS4{F0: false, F1: "foo bar"}
        _ = v32
    case "a":
        v1 += (v1 / (((v1 >> ((2) & 63))) | 1))
    default:
        var v33 int16 = 119
        _ = v33
    }
    v34 := GenS1{F0: 983}
    v35 := GenS3{F0: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
}
