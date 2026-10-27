package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 1
    var v3 uint32 = 42
    v4 := []string{}
    if (4294967295 < 42) {
        var v5 uint8 = 127
        var v6 string = ""
        _ = v5
        _ = v6
    } else {
        var v7 int = 4294967295
        _ = v7
    }
    if true {
        var v8 int32 = 398
        _ = v8
    }
    var v9 int32 = 42
    if true {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint = v2
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                if false {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v11)
                    if (143 <= 2) {
                        var v12 uint8 = 255
                        _ = v12
                    } else {
                        var v13 int32 = v9
                        _ = v13
                    }
                }
            }
            v1[fi10] = false
            _ = fi10
            _ = v11
        }
        var v14 int32 = v9
        _ = v14
    }
    if false {
        if (1e10 > 0.1) {
            var v15 uint32 = v3
            if (0 <= 4294967295) {
                if ((("" + "\n\t") != "\n\t") || false) {
                    v1[127] = true
                    if false {
                        v4 = append(v4, ("" + ("foo bar" + "")))
                        var v16 int64 = 0
                        _ = v16
                    } else {
                        v3 -= 717
                    }
                    var v17 int32 = v9
                    _ = v17
                }
            } else {
                var v18 int32 = 0
                v2 = (v2 ^ 1)
                if ((3.14159 + 1e10) == (0.1 - (0.0 + -2.25))) {
                    v19 := GenS1{F0: false, F1: -2147483648}
                    v20 := GenS1{F0: (!(false || (((-2.25 / float32(-2.25)) != (0.1 + 0.0)) && ((^(v15)) > (v15 | v3))))), F1: -2147483648}
                    var v21 int8 = -38
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    var v22 uint32 = v15
                    _ = v22
                }
                _ = v18
            }
            v3 -= (v15 | (v15 >> (((v2 + 255)) & 63)))
            _ = v15
        } else {
            v2 = (v2 * 841)
            fmt.Println(v3, v9, v2)
            v3 += v3
        }
    }
    var v23 uint = v2
    v1[-679] = false
    if (829 == -51) {
        if true {
            for fi24 := 0; fi24 < 16; fi24++ {
                var v25 bool = true
                _ = fi24
                _ = v25
            }
            var v26 float64 = 1.5
            _ = v26
        } else {
            if (1 == 4294967295) {
                v27 := GenS1{F0: true, F1: (v9 >> (((v2 / (((v23 % ((18446744073709551615) | 1))) | 1))) & 63))}
                var v28 int64 = 908
                _ = v27
                _ = v28
            } else {
                var v29 string = ("foo bar" + "a")
                v2 = (v23 & (v23 << ((v2) & 63)))
                _ = v29
            }
            v30 := GenS1{F0: false, F1: v9}
            _ = v30
        }
    } else {
        switch 1 {
        case -2147483648:
            v2 += v23
        case 0:
            var v31 uint32 = 0
            _ = v31
        default:
            v32 := GenS1{F0: ((v9 % (((v9 | (^(v9)))) | 1)) > 1), F1: v9}
            _ = v32
        }
        var v33 uint = v2
        _ = v33
    }
    var v34 float32 = (float32(((0.0 / 3.14159) / 1.5)) - -2.25)
    fmt.Println(v2, v34)
    var v35 float32 = (0.1 * v34)
    if (-1 < 859) {
        switch "hello" {
        case "\n\t":
            switch 32768 {
            case 6:
                v36 := GenS2{F0: v2, F1: true}
                _ = v36
            }
        default:
            v34 = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v9
}
