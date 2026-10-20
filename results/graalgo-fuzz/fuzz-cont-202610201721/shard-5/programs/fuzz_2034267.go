package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint64; F2 bool }
type GenS4 struct { F0 int8; F1 bool; F2 int }
type GenS5 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int8 = -128
    var v2 uint = 1
    var v3 uint32 = 42
    v4 := make(map[uint32]int)
    var v5 float32 = 1e10
    v6 := GenS1{F0: 42, F1: -1, F2: -1}
    var v7 float32 = v5
    var v8 int16 = -988
    if true {
        var v9 bool = true
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 7; fi11++ {
                v12 := GenS1{F0: 9223372036854775807, F1: -1, F2: -1}
                _ = fi11
                _ = v12
            }
            if v9 {
                for fi13 := 0; fi13 < 5; fi13++ {
                    v1 = v1
                    var v14 uint64 = 0
                    v9 = false
                    _ = fi13
                    _ = v14
                }
                v15 := GenS1{F0: 9223372036854775807, F1: 2147483647, F2: 42}
                v16 := []int{(fi10 >> (((^(v2))) & 63))}
                _ = v15
                _ = v16
            } else {
                if (v9 || v9) {
                    fmt.Println(v5)
                    v17 := GenS1{F0: 542, F1: 0, F2: 2147483647}
                    v7 = (v5 + v7)
                    _ = v17
                }
                for fi18 := 0; fi18 < 5; fi18++ {
                    var v19 string = ("\n\t" + ("foo bar" + ("" + "a")))
                    v2 = v2
                    _ = fi18
                    _ = v19
                }
            }
            _ = fi10
        }
        v3 = v3
        _ = v9
    } else {
        if false {
            fmt.Println(v7)
            v1 = 0
            v20 := []int{}
            _ = v20
        } else {
            v21 := GenS2{F0: (v1 >> (((v2 % (((v2 << (((^(v2))) & 63))) | 1))) & 63))}
            var v22 uint64 = 4294967295
            v23 := []int{-1}
            _ = v21
            _ = v22
            _ = v23
        }
    }
    if (-9223372036854775808 < 9223372036854775807) {
        v24 := GenS3{F0: (v1 | (v1 - v1)), F1: 438, F2: (!false)}
        for fi25 := 0; fi25 < 10; fi25++ {
            v3 = uint32(v2)
            switch fi25 {
            case -1:
                if (v8 != (v8 << (((v2 * (v2 << ((0) & 63)))) & 63))) {
                    if ((fi25 & 66) != fi25) {
                        var v26 uint64 = 0
                        switch fi25 {
                        case 446:
                            switch fi25 {
                            case 111:
                                v27 := GenS2{F0: 0}
                                _ = v27
                            case 1:
                                if false {
                                    v1 = v1
                                }
                            case -248:
                                v1 = v1
                            default:
                                v7 += 0.1
                            }
                        case 4294967295:
                            v26 = uint64(v2)
                        default:
                            fmt.Println(v7, v8)
                        }
                        _ = v26
                    } else {
                        fmt.Println(v1, v7, v3)
                        fmt.Println(v2)
                        fmt.Println(v3)
                    }
                } else {
                    v28 := GenS4{F0: v1, F1: (true || false), F2: 4294967295}
                    var v29 int = (fi25 & fi25)
                    v30 := GenS5{F0: 2147483647, F1: 38}
                    _ = v28
                    _ = v29
                    _ = v30
                }
            case 297:
                if false {
                    v5 = 3.14159
                }
            }
            _ = fi25
        }
        fmt.Println(v5, v8)
        _ = v24
    }
    v31 := GenS5{F0: -2147483648, F1: 255}
    v32 := GenS3{F0: (v1 / (((v1 % ((v1) | 1))) | 1)), F1: uint64(v2), F2: (v8 < 42)}
    v33 := []int{-696}
    fmt.Println(v5, v2, v5)
    v7 += v7
    fmt.Println(v1, v8, v2)
    var v34 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
