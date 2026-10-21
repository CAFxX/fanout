package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 int64; F1 int16; F2 uint }
type GenS3 struct { F0 uint8; F1 int8; F2 uint64 }

func main() {
    var v1 bool = (false && false)
    var v2 uint32 = 243
    v3 := make(map[string]string)
    v2 -= 42
    v4 := GenS1{F0: 1, F1: 18446744073709551615, F2: -271}
    fmt.Println(v2, v1)
    v5 := GenS1{F0: 42, F1: 266, F2: -1}
    v6 := []int{-2147483648}
    var v7 uint8 = 128
    v2 = (v2 / (((v2 & (v2 % (((v2 >> ((255) & 63))) | 1)))) | 1))
    if (!v1) {
        if v1 {
            v6 = append(v6, 0)
            if (((0.0 - 3.14159) * 1e10) != 0.1) {
                for fi8 := 0; fi8 < 7; fi8++ {
                    if true {
                        v9 := v3["foo bar"]
                        var v10 uint8 = v7
                        for fi11 := 0; fi11 < 5; fi11++ {
                            if v1 {
                                var v12 uint8 = (v7 / ((v10) | 1))
                                _ = v12
                            } else {
                                var v13 int32 = 2147483647
                                if (1 >= 4294967295) {
                                    if ((v10 >> ((1) & 63)) == 42) {
                                        switch ("foo bar" + v9) {
                                        case "":
                                            var v14 int16 = 0
                                            fmt.Println(v9, v7, fi8)
                                            switch (fi11 + (fi8 ^ -471)) {
                                            case 0:
                                                var v15 bool = v1
                                                _ = v15
                                            case 9223372036854775807:
                                                fmt.Println(v10, v7)
                                            case 455:
                                                var v16 float64 = -2.25
                                                _ = v16
                                            default:
                                                v9 = v9
                                            }
                                            _ = v14
                                        case "foo bar":
                                            v10 -= v7
                                        default:
                                            switch ("" + (((v9 + "a") + (v9 + v9)) + v9)) {
                                            case "":
                                                delete(v3, v9)
                                            case "\n\t":
                                                v2 = (v2 | 2147483648)
                                                v17 := GenS1{F0: fi11, F1: 4294967295, F2: 42}
                                                _ = v17
                                            }
                                            v7 = (^(v7))
                                            v13 -= (^(v13))
                                        }
                                        var v18 float32 = 0.0
                                        switch fi8 {
                                        case -556:
                                            var v19 int64 = 558
                                            _ = v19
                                        }
                                        delete(v3, ((("" + v9) + "\n\t") + v9))
                                        v2 += 2
                                        var v20 float64 = 0.1
                                        _ = v18
                                        _ = v20
                                    } else {
                                        v13 += v13
                                        var v21 int16 = -32768
                                        v9 = ((v9 + v9) + ("" + ""))
                                        _ = v21
                                    }
                                }
                                var v22 uint = 1
                                if v1 {
                                    var v23 float64 = ((3.14159 * 3.14159) + 1e10)
                                    v24 := v3[("" + ("foo bar" + (v9 + v9)))]
                                    v6 = append(v6, int(v2))
                                    _ = v23
                                    _ = v24
                                } else {
                                    var v25 string = v9
                                    _ = v25
                                }
                                _ = v13
                                _ = v22
                            }
                            _ = fi11
                        }
                        _ = v10
                        _ = v9
                    } else {
                        v2 = 2147483647
                        var v26 uint16 = 206
                        switch 2147483647 {
                        case -9223372036854775808:
                            if len(v6) > 0 {
                                v6[0] = (fi8 | fi8)
                            }
                        default:
                            v26 -= (v26 / (((^(v26))) | 1))
                        }
                        _ = v26
                    }
                    _ = fi8
                }
                delete(v3, ("hello" + ("hello" + ("foo bar" + "a"))))
                if ((-2.25 + float64(0.0)) <= -2.25) {
                    v27 := GenS1{F0: 216, F1: 4294967295, F2: -9223372036854775808}
                    v28 := GenS2{F0: -9223372036854775808, F1: int16(v2), F2: 0}
                    _ = v27
                    _ = v28
                } else {
                    for fi29 := 0; fi29 < 18; fi29++ {
                        for fi30 := 0; fi30 < 5; fi30++ {
                            v2 = (v2 % (((v2 * 2)) | 1))
                            _ = fi30
                        }
                        _ = fi29
                    }
                    if ((^(v2)) != v2) {
                        var v31 uint16 = 726
                        var v32 float64 = (float64(-2.25) - 1.5)
                        var v33 int8 = 0
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    v2 = (v2 | 4294967295)
                }
            }
            v34 := GenS3{F0: (v7 & uint8(v2)), F1: -103, F2: 42}
            _ = v34
        } else {
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 int16 = 1000
                _ = fi35
                _ = v36
            }
            for fi37 := 0; fi37 < 8; fi37++ {
                var v38 int16 = 798
                _ = fi37
                _ = v38
            }
            v39 := v3["foo bar"]
            _ = v39
        }
    } else {
        if (0 < 537) {
            var v40 uint64 = 713
            _ = v40
        }
    }
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
