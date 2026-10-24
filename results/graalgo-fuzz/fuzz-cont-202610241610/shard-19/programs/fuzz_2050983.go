package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint; F1 int; F2 uint64 }
type GenS4 struct { F0 int32; F1 int8 }

func main() {
    var v1 uint8 = 0
    var v2 string = "a"
    v3 := make(map[string]string)
    var v4 uint32 = 430
    if (false && false) {
        var v5 float64 = ((0.0 - float64(1.5)) * -2.25)
        _ = v5
    }
    v4 += 0
    v6 := GenS1{F0: 0, F1: -1, F2: "hello"}
    v1 = (v1 ^ (v1 << ((18446744073709551615) & 63)))
    var v7 int16 = -1
    var v8 uint8 = (v1 >> ((481) & 63))
    v2 = ((v2 + "foo bar") + v2)
    v7 -= (v7 & 42)
    if ((float64((-2.25 - (3.14159 / ((((((0.0 / -2.25) - float32(0.1)) - (0.1 - (-2.25 - 0.1))) * 0.1) + -2.25) + (0.0 + -2.25))))) + 1.5) > float64((0.1 - (float32(1.5) + 3.14159)))) {
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := []int{(fi9 ^ fi9), (fi9 % ((9223372036854775807) | 1))}
            for fi11 := 0; fi11 < 9; fi11++ {
                v12 := GenS1{F0: v8, F1: 42, F2: v2}
                if false {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    switch -870 {
                    case 131:
                        _ = len(v3)
                    case -760:
                        if true {
                            var v13 string = v2
                            _ = v13
                        }
                        if len(v10) > 0 {
                            v10[0] = fi11
                        }
                        var v14 float64 = (3.14159 / (float64(-2.25) / 1.0))
                        _ = v14
                    }
                    var v15 uint64 = 42
                    _ = v15
                }
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        switch 9223372036854775807 {
        case -2147483648:
            var v16 int8 = 127
            _ = v16
        case 1:
            var v17 int16 = v7
            _ = v17
        }
        fmt.Println(v2)
    } else {
        var v18 int32 = -2147483648
        var v19 int = -9223372036854775808
        if (v2 <= v2) {
            v20 := GenS1{F0: uint8(v4), F1: v18, F2: ((((v2 + v2) + (v2 + v2)) + "foo bar") + v2)}
            v19 = -124
            v18 -= -287
            _ = v20
        }
        _ = v18
        _ = v19
    }
    v21 := GenS2{F0: 0, F1: v7}
    v4 = (^(v4))
    if true {
        v7 += v7
    } else {
        for fi22 := 0; fi22 < 8; fi22++ {
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := GenS3{F0: 1, F1: (fi22 % ((int(v4)) | 1)), F2: 18446744073709551615}
                v7 = (v7 + (v7 * (v7 ^ (-(v7)))))
                switch v2 {
                case "foo bar":
                    v25 := GenS4{F0: -261, F1: 6}
                    _ = v25
                default:
                    v8 -= v1
                    fmt.Println(v4)
                }
                _ = fi23
                _ = v24
            }
            v7 = (v7 / ((737) | 1))
            _ = fi22
        }
    }
    fmt.Println(v2, v7)
    v26 := GenS1{F0: v8, F1: 42, F2: (v2 + (v2 + v2))}
    if false {
        var v27 int8 = 42
        v28 := GenS2{F0: 0, F1: (v7 + v7)}
        v29 := []int{42, 0, -454, 0}
        _ = v27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
