package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int64 }
type GenS2 struct { F0 string; F1 uint; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint }

func main() {
    var v1 float64 = ((3.14159 / 1.5) - -2.25)
    var v2 uint32 = 481
    v3 := make(map[int]string)
    v2 -= 518
    if false {
        if (false || false) {
            v4 := []int{127, 32767}
            v1 += v1
            fmt.Println(v1, v2)
            _ = v4
        } else {
            var v5 int64 = 42
            for fi6 := 0; fi6 < 4; fi6++ {
                var v7 string = "foo bar"
                _ = fi6
                _ = v7
            }
            _ = v5
        }
        v8 := []int{-144, 9223372036854775807, 32768, 32768}
        _ = v8
    } else {
        fmt.Println(v1, v1)
    }
    switch ("\n\t" + "hello") {
    case "foo bar":
        if ((100 >= 42) && true) {
            _ = len(v3)
        }
    case "hello":
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint8 = 58
            _ = fi9
            _ = v10
        }
    case "a":
        var v11 int64 = -9223372036854775808
        _ = v11
    default:
        var v12 int = 32767
        _ = v12
    }
    var v13 float32 = 3.14159
    v14 := GenS1{F0: 2, F1: 0, F2: 0}
    switch "foo bar" {
    case "hello":
        v1 += float64((v13 + v13))
    }
    v1 = float64(v13)
    v15 := GenS2{F0: "foo bar", F1: 2, F2: 65535}
    if true {
        if (-52 == -128) {
            switch -29 {
            case 667:
                fmt.Println(v1, v13, v2)
            case -442:
                fmt.Println(v1, v1)
            case -968:
                var v16 int64 = -1
                _ = v16
            }
            for fi17 := 0; fi17 < 15; fi17++ {
                v18 := []int{-808, fi17, 334, (fi17 & fi17)}
                v1 = float64(v13)
                _ = fi17
                _ = v18
            }
            if (0 != 946) {
                v1 += (float64(float32((v1 * (v1 * v1)))) - (v1 + v1))
                v19 := GenS1{F0: 2, F1: -1, F2: -9223372036854775808}
                v1 = (v1 * (float64(v13) * (v1 * 1e10)))
                _ = v19
            }
        } else {
            var v20 bool = false
            v21 := v3[127]
            _ = v20
            _ = v21
        }
    } else {
        v13 += (3.14159 + v13)
    }
    fmt.Println(v13, v13, v2)
    if ((511 > 1000) && true) {
        v22 := GenS3{F0: (v2 == 170), F1: 4294967295}
        v1 = v1
        _ = v22
    } else {
        v1 += -2.25
        for fi23 := 0; fi23 < 4; fi23++ {
            v2 = 1
            for fi24 := 0; fi24 < 17; fi24++ {
                switch (fi24 / ((fi24) | 1)) {
                case 964:
                    fmt.Println(v1, v1, fi23)
                case -9223372036854775808:
                    switch fi23 {
                    case 32767:
                        if true {
                            var v25 uint8 = 134
                            switch "a" {
                            case "":
                                var v26 int16 = 1
                                _ = v26
                            case "a":
                                v1 -= 0.0
                            default:
                                if ((v2 * 264) > 2147483648) {
                                    v27 := GenS4{F0: 1, F1: -2147483648}
                                    v28 := GenS5{F0: 968}
                                    _ = v27
                                    _ = v28
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        v25 = (^(v25))
                                    default:
                                        v25 += v25
                                    }
                                    var v29 uint8 = 255
                                    _ = v29
                                }
                            }
                            _ = v25
                        } else {
                            if true {
                                _ = len(v3)
                            } else {
                                v2 += 2147483648
                                switch 32767 {
                                case 254:
                                    fmt.Println(v13, v13, v13)
                                case -9223372036854775808:
                                    _ = len(v3)
                                }
                                v2 = 65535
                            }
                            v30 := GenS2{F0: "", F1: 42, F2: 331}
                            _ = v30
                        }
                    default:
                        v1 = ((-2.25 - v1) - (float64(v13) * 3.14159))
                    }
                }
                var v31 int32 = 42
                _ = fi24
                _ = v31
            }
            var v32 float32 = (0.0 + float32(1.5))
            _ = fi23
            _ = v32
        }
    }
    if false {
        v2 += (^(v2))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
