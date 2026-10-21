package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int8; F1 int32; F2 string }
func genFn0() int64 {
    var v6 uint32 = 875
    v7 := []int{9223372036854775807, 32768, -9223372036854775808}
    _ = v6
    _ = v7
    return -1
}
func genFn1() int32 {
    var v8 int64 = 9223372036854775807
    v8 -= (v8 * (^(v8)))
    if (0.1 < 0.1) {
        v8 = 42
    } else {
        var v9 int = 639
        _ = v9
    }
    _ = v8
    return 865
}

func main() {
    var v1 int32 = -181
    var v2 int32 = (v1 ^ v1)
    var v3 float64 = (1e10 + float64((1e10 * (0.1 * (float32(1.5) - 3.14159)))))
    var v4 uint8 = 51
    var v5 int32 = v1
    v10 := []int{-544}
    if (((3.14159 + v3) / -2.25) <= 3.14159) {
        v5 = v5
        v11 := GenS1{F0: "\n\t", F1: 513, F2: -564}
        _ = v11
    }
    switch 400 {
    case -1:
        if false {
            _ = len(v10)
        } else {
            switch ("" + ("\n\t" + "\n\t")) {
            case "foo bar":
                fmt.Println(v1, v2)
            }
            v2 = (v1 * (v2 * genFn1()))
        }
    case 1:
        v10 = append(v10, -668)
    default:
        v5 = 1
    }
    var v12 float32 = (1e10 / 3.14159)
    _ = len(v10)
    var v13 int32 = v1
    fmt.Println(v4, v5, v3)
    if (!true) {
        v14 := GenS2{F0: 1, F1: 42}
        if (!(false && true)) {
            var v15 int32 = -1
            for fi16 := 0; fi16 < 6; fi16++ {
                for fi17 := 0; fi17 < 9; fi17++ {
                    if true {
                        if true {
                            var v18 int8 = int8(v4)
                            _ = v18
                        } else {
                            v5 = 0
                            fmt.Println(v2)
                            switch ("a" + "") {
                            case "\n\t":
                                var v19 int32 = (v15 | v5)
                                _ = v19
                            default:
                                if len(v10) > 0 {
                                    v10[0] = (fi17 ^ (fi16 / ((989) | 1)))
                                }
                            }
                        }
                        v20 := GenS2{F0: v4, F1: 1}
                        _ = v20
                    } else {
                        switch (fi16 | fi17) {
                        case 32767:
                            var v21 int64 = genFn0()
                            _ = v21
                        case -9223372036854775808:
                            _ = len(v10)
                        }
                        var v22 float64 = v3
                        v2 += (v15 * 42)
                        _ = v22
                    }
                    v23 := GenS3{F0: 0, F1: v2, F2: "foo bar"}
                    _ = fi17
                    _ = v23
                }
                fmt.Println(v3)
                _ = fi16
            }
            _ = v15
        }
        _ = v14
    }
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v24 float32 = 0.0
    for fi25 := 0; fi25 < 4; fi25++ {
        for fi26 := 0; fi26 < 13; fi26++ {
            var v27 bool = true
            fmt.Println(v3)
            v24 = v24
            var v28 bool = v27
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    if true {
        v2 = v2
    }
    v5 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
