package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint64; F1 bool; F2 int8 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint = 42
    v4 := make(map[int]bool)
    var v5 uint8 = 202
    v6 := make(map[int]int)
    fmt.Println(v5)
    v7 := GenS1{F0: false, F1: 2147483648}
    var v8 int8 = -1
    var v9 uint64 = 42
    if (3.14159 > -2.25) {
        v8 -= (v8 & (v8 - (v8 & v8)))
    }
    v10 := GenS1{F0: (79 == 0), F1: 42}
    var v11 int32 = -442
    v12 := []int{-1, -2147483648, -9223372036854775808, 9223372036854775807}
    var v13 string = "foo bar"
    v9 -= (v9 & v9)
    var v14 float32 = 0.1
    switch "a" {
    case "a":
        var v15 float32 = v14
        _ = v15
    case "hello":
        v13 = (((v13 + "foo bar") + (v13 + ((v13 + v13) + v13))) + v13)
    case "\n\t":
        if false {
            if (v13 < (v13 + "hello")) {
                if (!(2147483648 <= 539)) {
                    var v16 string = ((v13 + v13) + v13)
                    if false {
                        fmt.Println(v5, v9, v14)
                        v5 = 127
                    }
                    _ = v16
                }
                if (true && true) {
                    v5 = v5
                    for fi17 := 0; fi17 < 15; fi17++ {
                        var v18 float32 = v14
                        v12 = append(v12, fi17)
                        if (true || true) {
                            if true {
                                if ((v8 ^ (v8 | -118)) != (v8 / (((v8 >> ((983) & 63))) | 1))) {
                                    if (2 == 384) {
                                        fmt.Println(v5)
                                        v19 := GenS1{F0: (0.1 > ((-2.25 * (1.5 / 0.1)) / 1.0)), F1: 42}
                                        var v20 int32 = (-(v11))
                                        _ = v19
                                        _ = v20
                                    } else {
                                        v21 := GenS2{F0: v5, F1: 255}
                                        _ = v21
                                    }
                                    switch fi17 {
                                    case -9223372036854775808:
                                        var v22 bool = false
                                        _ = v22
                                    case 2147483647:
                                        if len(v12) > 0 {
                                            v12[0] = (fi17 ^ (fi17 + fi17))
                                        }
                                    case 62:
                                        switch fi17 {
                                        case 217:
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        }
                                    }
                                    var v23 int16 = 174
                                    _ = v23
                                }
                                if (!false) {
                                    for fi24 := 0; fi24 < 1; fi24++ {
                                        fmt.Println(v18, v9, v11)
                                        if ((^(v8)) > 42) {
                                            if len(v12) > 0 {
                                                v12[0] = fi24
                                            }
                                            v14 = 3.14159
                                        }
                                        if true {
                                            if (!false) {
                                                v3 = v3
                                                v12 = append(v12, (^(fi24)))
                                                var v25 string = ("\n\t" + (v13 + v13))
                                                _ = v25
                                            } else {
                                                v8 = (v8 << ((18446744073709551615) & 63))
                                                fmt.Println(v8, v14, fi24)
                                                var v26 int = fi24
                                                _ = v26
                                            }
                                        }
                                        _ = fi24
                                    }
                                } else {
                                    if ((v8 / (((v8 + (v8 % ((v8) | 1)))) | 1)) == 55) {
                                        v27 := GenS2{F0: v5, F1: 727}
                                        _ = v27
                                    }
                                }
                            }
                            var v28 uint8 = 128
                            _ = v28
                        } else {
                            v14 += -2.25
                        }
                        _ = fi17
                        _ = v18
                    }
                } else {
                    switch v13 {
                    case "foo bar":
                        fmt.Println(v14, v3, v14)
                    default:
                        v29 := GenS2{F0: (v5 - 255), F1: 65535}
                        _ = v29
                    }
                }
                for fi30 := 0; fi30 < 19; fi30++ {
                    fmt.Println(v11, v13, v8)
                    v3 += v3
                    _ = fi30
                }
            } else {
                var v31 int16 = 0
                _ = v31
            }
        } else {
            var v32 int = 669
            for fi33 := 0; fi33 < 4; fi33++ {
                var v34 uint64 = v9
                _ = fi33
                _ = v34
            }
            var v35 int16 = int16(v8)
            _ = v32
            _ = v35
        }
    }
    switch int(v11) {
    case -2147483648:
        _ = len(v12)
    case 1:
        var v36 int8 = (v8 >> (((^(v3))) & 63))
        _ = v36
    default:
        v37 := GenS3{F0: (v9 * (v9 & v9)), F1: false, F2: (v8 & 42)}
        _ = v37
    }
    switch v13 {
    case "hello":
        v38 := GenS4{F0: v8}
        _ = v38
    }
    v3 -= 179
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v39 uint8 = (v5 % (((v5 ^ (v5 - 219))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
