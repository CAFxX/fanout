package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool; F1 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 float64 = 1e10
    var v4 uint32 = 672
    v4 -= 295
    for fi5 := 0; fi5 < 20; fi5++ {
        v4 = (v4 % (((v4 | v4)) | 1))
        _ = fi5
    }
    var v6 uint8 = 128
    if ((42 < -9223372036854775808) || true) {
        var v7 uint32 = v4
        var v8 uint8 = 42
        v3 += v3
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -2147483648, F1: 1}
    var v10 int16 = 1
    if false {
        v11 := GenS1{F0: 56, F1: 1}
        for fi12 := 0; fi12 < 3; fi12++ {
            if (1e10 < (3.14159 * 0.0)) {
                v13 := GenS2{F0: (!false), F1: 0}
                _ = len(v2)
                _ = v13
            } else {
                var v14 int8 = -128
                if ((!(!(("a" != "a") || true))) && (56 != v6)) {
                    var v15 bool = (102 != 2)
                    fmt.Println(v4, fi12, v3)
                    var v16 float64 = 3.14159
                    _ = v15
                    _ = v16
                }
                var v17 int8 = v14
                _ = v14
                _ = v17
            }
            fmt.Println(v6)
            switch 587 {
            case 785:
                var v18 uint8 = (^(v6))
                _ = v18
            case 32767:
                fmt.Println(v10, v10, v4)
            }
            _ = fi12
        }
        for fi19 := 0; fi19 < 17; fi19++ {
            fmt.Println(v10, v6, v6)
            if (42 == (v6 | v6)) {
                if true {
                    var v20 int = fi19
                    _ = v20
                }
                if false {
                    for fi21 := 0; fi21 < 11; fi21++ {
                        fmt.Println(v3, v10)
                        var v22 float32 = 0.0
                        var v23 uint16 = 65535
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                    fmt.Println(fi19, v10, v6)
                } else {
                    v6 = uint8(v4)
                    if false {
                        v6 = (v6 + (^(v6)))
                    }
                    var v24 uint8 = 128
                    _ = v24
                }
                v25 := v1[2147483647]
                _ = v25
            } else {
                v6 = 2
                if (1e10 < v3) {
                    v26 := GenS1{F0: fi19, F1: 2147483647}
                    v27 := GenS1{F0: (fi19 + fi19), F1: 2147483647}
                    var v28 uint8 = v6
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    switch fi19 {
                    case 127:
                        fmt.Println(v4)
                    case -9223372036854775808:
                        v10 = (v10 / ((v10) | 1))
                    }
                    v4 = v4
                    v29 := []int{(fi19 << ((0) & 63)), (fi19 - (^(fi19)))}
                    _ = v29
                }
            }
            v30 := GenS2{F0: (((42 != -237) || false) || (((!(true && false)) && false) || false)), F1: 1}
            _ = fi19
            _ = v30
        }
        _ = v11
    }
    fmt.Println(v10, v10, v4)
    var v31 float32 = 0.1
    var v32 int32 = 42
    var v33 uint8 = v6
    v31 = (v31 + v31)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v6
    _ = v9
}
