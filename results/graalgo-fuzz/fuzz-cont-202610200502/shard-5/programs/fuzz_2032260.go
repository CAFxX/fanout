package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string; F1 int64; F2 uint }
type GenS4 struct { F0 uint; F1 uint32 }
type GenS5 struct { F0 int64; F1 int8; F2 int }
type GenS6 struct { F0 string; F1 int32; F2 int32 }
func genFn0(p0 bool) int32 {
    var v5 int32 = 467
    switch "foo bar" {
    case "\n\t":
        v6 := []int{-2147483648, -40, -506}
        _ = v6
    case "hello":
        fmt.Println(v5, v5)
    default:
        v7 := []int{-346, -409, int(v5)}
        _ = v7
    }
    var v8 int64 = -9223372036854775808
    _ = p0
    _ = v5
    _ = v8
    return (^(v5))
}
func genFn1() uint {
    var v9 string = "\n\t"
    v9 = "a"
    var v10 int32 = 871
    if (!(0 >= 0)) {
        v10 = genFn0(false)
        v9 = ("\n\t" + v9)
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v12 := []int{(^(fi11)), (fi11 % (((-(fi11))) | 1)), 9223372036854775807}
        _ = fi11
        _ = v12
    }
    _ = v10
    _ = v9
    return 255
}

func main() {
    var v1 float64 = float64(float32(1.5))
    var v2 bool = true
    var v3 uint8 = 56
    var v4 int8 = -31
    fmt.Println(v2, v2)
    fmt.Println(v1, v4, v4)
    var v13 int64 = -501
    var v14 float32 = (0.1 - 0.1)
    v15 := GenS1{F0: ("foo bar" + "a"), F1: 1, F2: 680}
    v14 = ((((v14 - float32(1e10)) / v14) / ((0.1 * (v14 * (1e10 + -2.25))) * (0.0 * -2.25))) - (v14 - 3.14159))
    var v16 int = 312
    v17 := []int{}
    fmt.Println(v1, v16, v16)
    if v2 {
        if (v14 > 0.0) {
            v1 = v1
            if false {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                v18 := GenS2{F0: 42}
                for fi19 := 0; fi19 < 5; fi19++ {
                    v13 += (-(v13))
                    _ = fi19
                }
                _ = v18
            } else {
                v20 := GenS3{F0: "foo bar", F1: (v13 & (v13 - 0)), F2: 1}
                _ = v20
            }
        } else {
            var v21 float64 = (v1 * float64((1.5 - v14)))
            if ((32767 < 0) || v2) {
                if v2 {
                    v22 := GenS4{F0: 4294967295, F1: 0}
                    _ = v22
                } else {
                    var v23 uint32 = 4294967295
                    switch (v16 + (v16 / (((v16 ^ (v16 ^ v16))) | 1))) {
                    case 0:
                        switch "hello" {
                        case "\n\t":
                            v17 = append(v17, v16)
                        case "foo bar":
                            v24 := GenS5{F0: -9223372036854775808, F1: (v4 | (v4 - v4)), F2: (v16 >> ((4294967295) & 63))}
                            _ = v24
                        default:
                            for fi25 := 0; fi25 < 17; fi25++ {
                                var v26 uint8 = 0
                                v1 += v21
                                _ = fi25
                                _ = v26
                            }
                        }
                    case -511:
                        for fi27 := 0; fi27 < 9; fi27++ {
                            switch v16 {
                            case 4294967295:
                                v28 := GenS1{F0: "foo bar", F1: 2, F2: 0}
                                _ = v28
                            }
                            v17 = append(v17, int(v4))
                            _ = fi27
                        }
                    case 41:
                        fmt.Println(v2)
                    }
                    v29 := GenS6{F0: "a", F1: 212, F2: 399}
                    _ = v23
                    _ = v29
                }
            } else {
                v3 += 1
                if false {
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    fmt.Println(v1, v16, v4)
                }
            }
            _ = v21
        }
        var v30 uint8 = v3
        fmt.Println(v3)
        _ = v30
    }
    if (288 == 2) {
        if false {
            if ((18446744073709551615 == 42) || (true && false)) {
                var v31 uint8 = 101
                _ = v31
            } else {
                var v32 int64 = v13
                v33 := GenS3{F0: ("hello" + ("\n\t" + "a")), F1: (v32 % ((v13) | 1)), F2: 2}
                _ = v32
                _ = v33
            }
            if len(v17) >= 2 {
                v17 = v17[1:len(v17)]
            }
            v34 := GenS5{F0: v13, F1: v4, F2: (v16 % ((int(v4)) | 1))}
            _ = v34
        }
        v17 = append(v17, 0)
        if len(v17) > 0 {
            v17[0] = 930
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
