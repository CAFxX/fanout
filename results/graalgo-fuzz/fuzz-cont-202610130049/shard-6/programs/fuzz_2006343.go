package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint64; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8; F1 int32; F2 int32 }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint8 }

func main() {
    var v1 int32 = 2147483647
    var v2 bool = true
    var v3 int8 = -55
    fmt.Println(v2)
    fmt.Println(v2, v3, v2)
    var v4 int = 32767
    if (1 == 9223372036854775807) {
        var v5 uint = 0
        _ = v5
    } else {
        switch v4 {
        case 120:
            for fi6 := 0; fi6 < 20; fi6++ {
                if v2 {
                    var v7 int32 = 0
                    _ = v7
                }
                v8 := GenS1{F0: (^(v3)), F1: v2}
                if (!v2) {
                    for fi9 := 0; fi9 < 5; fi9++ {
                        v10 := []int{}
                        _ = fi9
                        _ = v10
                    }
                    var v11 int16 = 0
                    fmt.Println(fi6, v2, v3)
                    _ = v11
                } else {
                    fmt.Println(v1)
                    var v12 uint64 = 42
                    if (("\n\t" + "a") == "\n\t") {
                        v13 := []int{}
                        v2 = (v2 && (255 <= 42))
                        for fi14 := 0; fi14 < 2; fi14++ {
                            fmt.Println(v4)
                            v12 = v12
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            _ = fi14
                        }
                        _ = v13
                    } else {
                        var v15 string = "a"
                        switch v15 {
                        case "hello":
                            v4 += (fi6 * (-(fi6)))
                        default:
                            v12 -= 2
                        }
                        _ = v15
                    }
                    _ = v12
                }
                _ = fi6
                _ = v8
            }
        case 32767:
            v1 = (v1 | (v1 ^ v1))
        case 1:
            if v2 {
                v16 := GenS2{F0: 42, F1: -9223372036854775808, F2: 1}
                fmt.Println(v4, v4)
                _ = v16
            } else {
                v17 := []int{v4, (v4 + 690)}
                switch 4294967295 {
                case 100:
                    v18 := GenS3{F0: 42, F1: (v1 + v1), F2: v1}
                    _ = v18
                }
                fmt.Println(v1, v1)
                _ = v17
            }
        default:
            v3 -= (v3 % ((100) | 1))
        }
    }
    if false {
        v19 := GenS1{F0: (^(v3)), F1: v2}
        v20 := []int{v4}
        _ = v19
        _ = v20
    } else {
        v21 := []int{v4, (v4 ^ int(v1)), v4, v4}
        _ = v21
    }
    v4 += v4
    v22 := GenS4{F0: 18446744073709551615, F1: 255, F2: 127}
    v23 := GenS1{F0: (-(v3)), F1: true}
    var v24 int32 = v1
    var v25 float32 = 0.0
    v4 = 42
    var v26 int16 = 859
    var v27 uint = 255
    if true {
        if v2 {
            var v28 int64 = 821
            v29 := []int{v4, (v4 + (v4 << ((v27) & 63)))}
            _ = v28
            _ = v29
        }
        v30 := []int{128, (v4 & (v4 / ((9223372036854775807) | 1)))}
        _ = v30
    }
    switch "a" {
    case "foo bar":
        if false {
            v24 = (v24 % (((v1 ^ (v24 >> ((65535) & 63)))) | 1))
        } else {
            v26 = (v26 - 389)
            v31 := GenS1{F0: (v3 / ((0) | 1)), F1: false}
            fmt.Println(v4)
            _ = v31
        }
    case "hello":
        if (v2 || ((v4 % ((-1) | 1)) <= (v4 * (v4 - (-(v4)))))) {
            fmt.Println(v3, v2)
            fmt.Println(v4, v26, v25)
        }
    default:
        v27 = (^(v27))
    }
    var v32 int64 = -1
    var v33 int32 = v1
    var v34 uint = (v27 & 277)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
