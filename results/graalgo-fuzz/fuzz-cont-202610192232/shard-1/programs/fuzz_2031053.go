package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint64; F1 int64; F2 int64 }
type GenS3 struct { F0 uint; F1 int16; F2 uint8 }
func genFn0(p0 int32, p1 uint32, p2 uint64) string {
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: 587, F1: 1}
    for fi8 := 0; fi8 < 20; fi8++ {
        v6 -= 42
        v9 := GenS1{F0: 1, F1: 32767}
        v10 := GenS2{F0: 1, F1: v6, F2: v6}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{0, 32768, 32767, 4294967295}
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v6
    _ = v7
    return "a"
}
func genFn1(p0 int, p1 int8, p2 uint16) int32 {
    v12 := GenS1{F0: int16(p1), F1: 32767}
    var v13 int8 = (p1 / ((p1) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 65535
    var v3 float64 = v1
    var v4 uint = 4294967295
    var v5 uint32 = 448
    var v14 float64 = v3
    v15 := []int{42}
    v4 = (^(v4))
    v16 := GenS2{F0: 4294967295, F1: 9223372036854775807, F2: 646}
    if (1 == 255) {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        fmt.Println(v1)
    } else {
        var v17 uint64 = 272
        _ = v17
    }
    var v18 uint8 = 34
    switch 1 {
    case 42:
        var v19 uint = (v4 | (^(v2)))
        _ = v19
    case 1:
        v18 = (v18 * (v18 / (((^(v18))) | 1)))
    case 128:
        if true {
            if (uint8(v4) > v18) {
                var v20 string = ("a" + "")
                _ = v20
            }
        } else {
            var v21 uint16 = 193
            for fi22 := 0; fi22 < 3; fi22++ {
                for fi23 := 0; fi23 < 2; fi23++ {
                    if (true && false) {
                        var v24 uint16 = (^(v21))
                        var v25 float32 = float32(v1)
                        _ = v24
                        _ = v25
                    } else {
                        var v26 int8 = 100
                        _ = v26
                    }
                    v27 := GenS3{F0: (v2 >> (((v4 - uint(v18))) & 63)), F1: 335, F2: v18}
                    _ = fi23
                    _ = v27
                }
                var v28 int32 = 1
                if true {
                    v29 := GenS2{F0: 2, F1: int64(fi22), F2: 0}
                    _ = v29
                }
                _ = fi22
                _ = v28
            }
            _ = v21
        }
    default:
        var v30 int64 = -861
        _ = v30
    }
    if (float32(0.0) > (3.14159 / 1e10)) {
        if ((true || (false && (-48 == 0))) || false) {
            v18 += (v18 | (v18 | (v18 + (v18 | v18))))
            for fi31 := 0; fi31 < 14; fi31++ {
                switch fi31 {
                case 181:
                    if true {
                        for fi32 := 0; fi32 < 14; fi32++ {
                            v33 := GenS2{F0: 18446744073709551615, F1: -9223372036854775808, F2: 1}
                            fmt.Println(v3, v2, v2)
                            var v34 float64 = 3.14159
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        var v35 uint64 = 0
                        v36 := GenS2{F0: v35, F1: 479, F2: -1}
                        _ = v35
                        _ = v36
                    } else {
                        if len(v15) > 0 {
                            v15[0] = (fi31 ^ fi31)
                        }
                    }
                default:
                    fmt.Println(v18)
                }
                _ = fi31
            }
        } else {
            v3 = 0.1
            switch (("\n\t" + "\n\t") + "foo bar") {
            case "hello":
                v14 -= ((v14 * 1e10) * ((v1 * v3) + v14))
            case "a":
                fmt.Println(v1, v4)
            case "":
                switch -99 {
                case 509:
                    v2 = (v4 | v2)
                default:
                    switch "" {
                    case "":
                        var v37 string = "\n\t"
                        _ = v37
                    case "a":
                        fmt.Println(v4)
                    case "\n\t":
                        for fi38 := 0; fi38 < 19; fi38++ {
                            v39 := GenS1{F0: -1, F1: -1}
                            for fi40 := 0; fi40 < 1; fi40++ {
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                v5 += v5
                                _ = fi40
                            }
                            v41 := GenS1{F0: -601, F1: 992}
                            _ = fi38
                            _ = v39
                            _ = v41
                        }
                    }
                }
            }
        }
    }
    v42 := GenS1{F0: 42, F1: -892}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
}
