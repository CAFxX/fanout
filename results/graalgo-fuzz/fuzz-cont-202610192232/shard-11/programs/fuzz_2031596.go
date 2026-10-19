package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 int8; F2 int }
func genFn0(p0 int, p1 uint8, p2 int64) int8 {
    p1 -= p1
    switch 861 {
    case 42:
        switch (p0 % ((p0) | 1)) {
        case -237:
            if (!false) {
                var v5 int32 = 1
                _ = v5
            } else {
                p0 += (p0 % (((p0 * p0)) | 1))
                p2 = p2
            }
        case -9223372036854775808:
            var v6 uint = 151
            _ = v6
        }
    case -9223372036854775808:
        v7 := GenS1{F0: 436, F1: 2147483648}
        _ = v7
    case 851:
        v8 := []int{p0}
        _ = v8
    }
    var v9 uint8 = (p1 - (^(p1)))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -101
}

func main() {
    v1 := []uint32{4294967295, 42, 2, 743}
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]bool)
    var v4 uint = 950
    v4 += 255
    if (1 != -2147483648) {
        fmt.Println(v4)
        v4 = 42
        delete(v3, 4294967295)
    } else {
        fmt.Println(v4)
    }
    if true {
        var v10 uint = v4
        v11 := v3[241]
        v10 = (v4 & v4)
        _ = v10
        _ = v11
    }
    fmt.Println(v4)
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(fi13, v4, v12)
        _ = fi13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int64 = 9223372036854775807
        v16 := GenS2{F0: v12}
        switch (fi14 - 42) {
        case -9223372036854775808:
            fmt.Println(fi14, v12)
        case -1:
            v17 := GenS1{F0: 4294967295, F1: 4294967295}
            _ = v17
        case 42:
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS2{F0: ("foo bar" + "hello")}
                if ((((v12 + "") + v12) + v12) >= v12) {
                    v20 := GenS3{F0: 1}
                    v21 := v2[uint32(v4)]
                    v22 := GenS1{F0: 2, F1: 597}
                    if ((-(v15)) < (^(v15))) {
                        v23 := GenS4{F0: -32768}
                        var v24 uint32 = 0
                        var v25 uint16 = uint16(v4)
                        _ = v23
                        _ = v24
                        _ = v25
                    } else {
                        _ = len(v1)
                    }
                    _ = v20
                    _ = v21
                    _ = v22
                } else {
                    if true {
                        v12 = "\n\t"
                        fmt.Println(v12)
                    } else {
                        v12 = v12
                        v12 = (v12 + v12)
                        v26 := GenS5{F0: 133}
                        var v27 int64 = -1
                        if (false && (255 == 128)) {
                            var v28 bool = (true && false)
                            switch 0 {
                            case 128:
                                v29 := GenS6{F0: 42, F1: 42, F2: 0}
                                _ = v29
                            case -1:
                                v15 = 1
                            default:
                                v28 = v28
                            }
                            v12 = (v12 + v12)
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                            _ = v28
                        }
                        fmt.Println(v12, v15, v4)
                        _ = v26
                        _ = v27
                    }
                }
                _ = fi18
                _ = v19
            }
        default:
            var v30 float32 = (1e10 * 0.1)
            _ = v30
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    switch 32768 {
    case 1:
        v31 := v2[987]
        _ = v31
    case -623:
        v12 = v12
    }
    var v32 int = -267
    var v33 bool = ((((1.5 - float32(-2.25)) * 0.1) * float32(1.5)) >= ((1e10 + (3.14159 * (3.14159 / ((float32(1.5) / (1.5 / 1.0)) * 1.5)))) / 1.5))
    var v34 string = v12
    v34 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
