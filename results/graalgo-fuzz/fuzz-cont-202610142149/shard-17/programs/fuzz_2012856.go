package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 bool; F2 int }
type GenS3 struct { F0 uint32; F1 uint16; F2 bool }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32 }

func main() {
    var v1 uint16 = 2
    var v2 float32 = 0.1
    var v3 int = 1
    var v4 int8 = 100
    var v5 int8 = (^(v4))
    for fi6 := 0; fi6 < 8; fi6++ {
        v7 := []int{(fi6 * 0), (v3 ^ (v3 << ((42) & 63)))}
        fmt.Println(v2, v2, v1)
        _ = fi6
        _ = v7
    }
    var v8 int = (v3 * (v3 & v3))
    v8 = -9223372036854775808
    fmt.Println(v1)
    for fi9 := 0; fi9 < 15; fi9++ {
        if (569 <= 521) {
            var v10 int8 = (v5 / (((v4 % ((52) | 1))) | 1))
            v11 := []int{}
            v12 := GenS1{F0: (v10 << ((268) & 63))}
            _ = v10
            _ = v11
            _ = v12
        }
        fmt.Println(v3)
        v13 := GenS1{F0: 100}
        _ = fi9
        _ = v13
    }
    switch (v3 * 128) {
    case -806:
        if (v8 >= -910) {
            v14 := GenS2{F0: 683, F1: true, F2: v8}
            for fi15 := 0; fi15 < 8; fi15++ {
                if false {
                    for fi16 := 0; fi16 < 3; fi16++ {
                        if (true && true) {
                            v17 := GenS1{F0: (v4 << ((0) & 63))}
                            var v18 int = (v3 << ((352) & 63))
                            v2 += v2
                            _ = v17
                            _ = v18
                        } else {
                            v19 := GenS3{F0: 497, F1: (v1 % ((0) | 1)), F2: (843 < 2)}
                            _ = v19
                        }
                        _ = fi16
                    }
                    fmt.Println(v5, v2)
                    v5 += (v5 % ((-113) | 1))
                } else {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        v21 := []int{128, (fi15 - (fi15 & fi15)), (^(fi15))}
                        v3 = fi15
                        _ = fi20
                        _ = v21
                    }
                    if false {
                        v22 := GenS4{F0: v1}
                        _ = v22
                    }
                    v5 = (v5 & (v5 ^ v5))
                }
                v8 = fi15
                for fi23 := 0; fi23 < 7; fi23++ {
                    if true {
                        if false {
                            var v24 int64 = -261
                            if false {
                                v1 = 32767
                                fmt.Println(v4)
                                v25 := GenS5{F0: 42, F1: 0}
                                _ = v25
                            } else {
                                if true {
                                    v26 := GenS1{F0: (v5 * (v4 % ((-128) | 1)))}
                                    _ = v26
                                } else {
                                    v27 := []int{}
                                    _ = v27
                                }
                                var v28 uint16 = 1
                                _ = v28
                            }
                            _ = v24
                        } else {
                            v29 := []int{v8, (fi15 % (((fi23 >> ((727) & 63))) | 1)), v8, v8}
                            v1 -= v1
                            fmt.Println(v1)
                            _ = v29
                        }
                        var v30 uint = 18446744073709551615
                        switch 42 {
                        case 32768:
                            v31 := []int{(^(fi23))}
                            _ = v31
                        default:
                            v32 := []int{}
                            _ = v32
                        }
                        _ = v30
                    } else {
                        v4 = (v4 * (v5 - v5))
                        fmt.Println(fi15)
                    }
                    if (1e10 == 0.0) {
                        v33 := GenS6{F0: (^(v1)), F1: 1}
                        v34 := GenS7{F0: -2147483648}
                        _ = v33
                        _ = v34
                    } else {
                        v35 := []int{}
                        v4 -= v5
                        _ = v35
                    }
                    _ = fi23
                }
                _ = fi15
            }
            var v36 string = ""
            _ = v14
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        v2 += (v2 - 0.0)
        switch 128 {
        case 1:
            fmt.Println(v1, v2)
        case -239:
            var v38 int8 = v4
            _ = v38
        }
        _ = fi37
    }
    fmt.Println(v5, v3, v5)
    v3 = v3
    var v39 uint = 18446744073709551615
    v40 := GenS2{F0: -1, F1: (!true), F2: v8}
    var v41 uint32 = 2147483647
    v8 = v8
    v5 = (^(v4))
    fmt.Println(v41, v1, v5)
    v42 := GenS8{F0: 2147483647}
    fmt.Println(v1, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v8
}
