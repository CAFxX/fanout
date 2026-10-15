package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int32 = 523
    var v2 int = 128
    var v3 uint = 255
    var v4 int8 = 1
    var v5 int = (v2 | v2)
    var v6 float32 = 1.5
    switch (v5 % ((v2) | 1)) {
    case 127:
        v7 := []int{(v2 | (^(v2)))}
        _ = v7
    }
    v8 := GenS1{F0: 138, F1: 2, F2: 0}
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            switch 445 {
            case -636:
                v11 := []int{}
                _ = v11
            case -592:
                v12 := GenS1{F0: 128, F1: 2, F2: 42}
                _ = v12
            default:
                var v13 uint = (v3 + 673)
                _ = v13
            }
            _ = fi10
        }
        v4 -= (v4 + -128)
        fmt.Println(v3, v1, v4)
        _ = fi9
    }
    v1 = (v1 & v1)
    var v14 string = ""
    fmt.Println(v2, v6, v4)
    switch ((v14 + "hello") + "") {
    case "foo bar":
        for fi15 := 0; fi15 < 3; fi15++ {
            v4 = (v4 << (((^(v3))) & 63))
            v2 -= (^(v2))
            fmt.Println(v14, v1, v4)
            _ = fi15
        }
    default:
        v16 := GenS2{F0: 923, F1: 147, F2: v4}
        _ = v16
    }
    if (0 == 0) {
        switch v14 {
        case "\n\t":
            v17 := GenS1{F0: 2, F1: 42, F2: 18446744073709551615}
            _ = v17
        case "hello":
            var v18 uint32 = 2147483648
            _ = v18
        }
    } else {
        if (!false) {
            v6 = 0.1
        } else {
            if (785 != 255) {
                if false {
                    var v19 int8 = v4
                    if (32767 < -1) {
                        for fi20 := 0; fi20 < 11; fi20++ {
                            if false {
                                var v21 string = "a"
                                if (((v21 + "a") + v21) == ("" + (v21 + "hello"))) {
                                    v22 := []int{v5, v2, 650}
                                    _ = v22
                                }
                                _ = v21
                            } else {
                                fmt.Println(v6, v2, v1)
                                for fi23 := 0; fi23 < 17; fi23++ {
                                    v24 := GenS3{F0: -9223372036854775808}
                                    _ = fi23
                                    _ = v24
                                }
                            }
                            v2 += 2147483647
                            _ = fi20
                        }
                    }
                    _ = v19
                }
                v2 -= -76
            } else {
                v1 -= (v1 & v1)
            }
            switch ("foo bar" + v14) {
            case "":
                if true {
                    var v25 float32 = v6
                    _ = v25
                }
            case "hello":
                switch -1 {
                case 2147483647:
                    v3 = v3
                case 42:
                    for fi26 := 0; fi26 < 12; fi26++ {
                        v27 := []int{1, 32768, (-(v2)), (v2 + fi26)}
                        _ = fi26
                        _ = v27
                    }
                default:
                    var v28 int64 = -9223372036854775808
                    _ = v28
                }
            case "\n\t":
                v14 = ""
            }
        }
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        v6 -= v6
        _ = fi29
    }
    if (!(true && true)) {
        v30 := []int{}
        v31 := GenS2{F0: uint32(v3), F1: 124, F2: v4}
        var v32 int32 = v1
        _ = v30
        _ = v31
        _ = v32
    }
    v6 = float32((0.0 + 0.0))
    fmt.Println(v14, v1, v14)
    v1 = (v1 ^ (v1 * (v1 >> ((v3) & 63))))
    var v33 uint = (v3 ^ v3)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
