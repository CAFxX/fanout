package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint8; F1 int8; F2 int8 }
type GenS3 struct { F0 int8; F1 int8; F2 bool }

func main() {
    var v1 uint16 = 42
    var v2 uint = 334
    v3 := make(map[string]int)
    var v4 int16 = 1000
    var v5 uint = (v2 * v2)
    var v6 string = "foo bar"
    v5 = 1
    var v7 float32 = ((0.0 * (0.0 - 0.1)) - (float32((1e10 - float64((3.14159 + 1.5)))) + float32(1.5)))
    for fi8 := 0; fi8 < 7; fi8++ {
        v6 = "hello"
        _ = fi8
    }
    v9 := GenS1{F0: v1, F1: (v1 * v1), F2: -1}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int8 = 42
        v12 := GenS1{F0: v1, F1: 864, F2: 1}
        _ = fi10
        _ = v11
        _ = v12
    }
    switch (((v6 + v6) + "a") + (("hello" + v6) + "hello")) {
    case "foo bar":
        v6 = v6
    case "\n\t":
        v3[v6] = -9223372036854775808
    default:
        for fi13 := 0; fi13 < 19; fi13++ {
            v14 := GenS2{F0: uint8(v5), F1: 42, F2: 1}
            _ = fi13
            _ = v14
        }
    }
    var v15 int = 0
    var v16 int8 = -1
    v4 += 42
    if false {
        var v17 uint32 = 42
        if (v7 == v7) {
            var v18 int = 2147483647
            _ = v18
        } else {
            v19 := []int{-648, v15, v15, v15}
            v1 -= v1
            if (true || false) {
                _ = len(v19)
                if false {
                    for fi20 := 0; fi20 < 18; fi20++ {
                        var v21 int64 = 9223372036854775807
                        _ = fi20
                        _ = v21
                    }
                    var v22 bool = true
                    _ = v22
                }
                for fi23 := 0; fi23 < 2; fi23++ {
                    if false {
                        fmt.Println(v17, v4, v7)
                        if false {
                            var v24 int = (^(fi23))
                            _ = v24
                        } else {
                            if true {
                                v25 := GenS2{F0: 42, F1: v16, F2: (v16 << (((v5 + v2)) & 63))}
                                for fi26 := 0; fi26 < 17; fi26++ {
                                    var v27 int = 42
                                    fmt.Println(fi23, fi26, v16)
                                    _ = fi26
                                    _ = v27
                                }
                                _ = v25
                            }
                            v28 := GenS1{F0: (v1 ^ (v1 ^ v1)), F1: (v1 << (((v5 >> (((v5 >> (((v5 % (((v2 & (v2 >> ((v5) & 63)))) | 1))) & 63))) & 63))) & 63)), F2: -2147483648}
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 8; fi29++ {
                            var v30 bool = (v7 < (-2.25 * v7))
                            switch (v15 + fi29) {
                            case 9223372036854775807:
                                var v31 int16 = (v4 | v4)
                                _ = v31
                            case 128:
                                v32 := v3[v6]
                                _ = v32
                            case 663:
                                if false {
                                    var v33 uint8 = 60
                                    if false {
                                        delete(v3, "foo bar")
                                    }
                                    _ = v33
                                }
                            }
                            var v34 int16 = (v4 >> (((v5 * (v5 | 89))) & 63))
                            if true {
                                v6 = (v6 + v6)
                                fmt.Println(v5, v34)
                                v35 := v3[v6]
                                _ = v35
                            } else {
                                v36 := GenS2{F0: 42, F1: (v16 ^ v16), F2: (v16 * v16)}
                                fmt.Println(v2, v5, v17)
                                _ = v36
                            }
                            v16 = -128
                            v15 += fi29
                            if (!true) {
                                v7 = (((1e10 * (3.14159 * v7)) / v7) + float32(1.5))
                                v37 := GenS3{F0: 127, F1: (v16 << ((v2) & 63)), F2: false}
                                var v38 uint64 = 42
                                _ = v37
                                _ = v38
                            } else {
                                var v39 uint32 = v17
                                _ = v39
                            }
                            _ = fi29
                            _ = v30
                            _ = v34
                        }
                    } else {
                        var v40 uint16 = 65535
                        var v41 int = 9223372036854775807
                        _ = v40
                        _ = v41
                    }
                    _ = fi23
                }
            } else {
                v42 := GenS3{F0: (-(v16)), F1: 100, F2: ((uint8(v1) > 1) && false)}
                _ = v42
            }
            _ = v19
        }
        v43 := GenS3{F0: v16, F1: (^(v16)), F2: (v6 < v6)}
        _ = v17
        _ = v43
    }
    v4 = v4
    var v44 uint16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
