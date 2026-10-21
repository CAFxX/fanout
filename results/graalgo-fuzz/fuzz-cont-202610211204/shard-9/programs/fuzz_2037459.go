package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint16; F1 int8; F2 uint }
func genFn0() uint64 {
    v6 := GenS1{F0: 65535, F1: false}
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    _ = v6
    return 42
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]bool)
    var v3 int16 = 1000
    var v4 uint = 18446744073709551615
    var v5 string = v1
    fmt.Println(v4)
    if true {
        v5 = "hello"
    }
    switch ((((v5 + v5) + "foo bar") + "hello") + ((v5 + (("" + (v1 + (v5 + v1))) + (v1 + v1))) + "a")) {
    case "":
        var v8 int64 = -1
        _ = v8
    case "foo bar":
        v9 := GenS1{F0: (^(v4)), F1: (42 != -1)}
        _ = v9
    case "hello":
        delete(v2, 887)
    default:
        var v10 uint16 = 2
        _ = v10
    }
    v5 = v1
    v11 := v2[2147483648]
    v5 = "\n\t"
    v12 := GenS2{F0: 255, F1: 2, F2: -1}
    switch -1 {
    case -456:
        for fi13 := 0; fi13 < 13; fi13++ {
            if (2 != 65535) {
                v14 := GenS3{F0: 128, F1: v4}
                _ = v14
            } else {
                v3 -= v3
                if (!(16 == 236)) {
                    if v11 {
                        v11 = false
                    } else {
                        for fi15 := 0; fi15 < 14; fi15++ {
                            v16 := GenS4{F0: 42, F1: 100, F2: (v4 - v4)}
                            _ = fi15
                            _ = v16
                        }
                    }
                    var v17 float32 = float32(0.0)
                    fmt.Println(v5, v4)
                    _ = v17
                } else {
                    var v18 uint64 = 18446744073709551615
                    if v11 {
                        if v11 {
                            var v19 uint = (v4 / ((v4) | 1))
                            var v20 int16 = v3
                            _ = v19
                            _ = v20
                        }
                        v21 := []int{-9223372036854775808, (fi13 << ((v4) & 63))}
                        _ = v21
                    }
                    v1 = v1
                    _ = v18
                }
            }
            fmt.Println(v3, fi13, v11)
            _ = fi13
        }
    case 488:
        if v11 {
            if ((0.1 <= 1e10) || v11) {
                v22 := GenS1{F0: 4294967295, F1: v11}
                _ = v22
            } else {
                fmt.Println(v11)
            }
            fmt.Println(v1)
            delete(v2, 200)
        } else {
            fmt.Println(v11, v1)
            for fi23 := 0; fi23 < 12; fi23++ {
                if v11 {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        if ((2 == 255) && v11) {
                            v4 += 18446744073709551615
                            v25 := GenS4{F0: 42, F1: 0, F2: (^(v4))}
                            v26 := GenS3{F0: 0, F1: (v4 * (v4 - (v4 + 2)))}
                            _ = v25
                            _ = v26
                        }
                        v11 = (1 != 2)
                        var v27 uint8 = 122
                        _ = fi24
                        _ = v27
                    }
                    var v28 int32 = -790
                    if v11 {
                        for fi29 := 0; fi29 < 4; fi29++ {
                            v30 := []int{fi23}
                            v1 = "foo bar"
                            delete(v2, 397)
                            _ = fi29
                            _ = v30
                        }
                        fmt.Println(fi23, v5, v11)
                        var v31 bool = (-161 != -294)
                        _ = v31
                    } else {
                        v5 = (v5 + ((("hello" + ("\n\t" + (v5 + v5))) + (("\n\t" + v1) + "hello")) + (v1 + ("" + ((v5 + v5) + "\n\t")))))
                        v5 = (v5 + (v5 + ((v1 + "foo bar") + v5)))
                    }
                    _ = v28
                }
                v3 -= v3
                _ = fi23
            }
        }
    default:
        var v32 uint16 = 42
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
