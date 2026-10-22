package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int16 = 0
    v2 := []uint32{2147483648, 569}
    var v3 uint16 = 65535
    v4 := make(map[uint32]bool)
    var v5 float32 = (float32(1e10) - (1.5 / 1.5))
    v6 := GenS1{F0: v1, F1: (!(!(!true))), F2: -128}
    v2 = append(v2, 2147483647)
    var v7 uint16 = v3
    for fi8 := 0; fi8 < 13; fi8++ {
        switch ("foo bar" + "a") {
        case "a":
            v5 -= -2.25
        case "":
            for fi9 := 0; fi9 < 16; fi9++ {
                switch 32767 {
                case 665:
                    v1 -= 259
                case 42:
                    v7 += (^(v7))
                default:
                    switch 2147483647 {
                    case 42:
                        v5 = v5
                        fmt.Println(fi9, fi8)
                    }
                    v3 += v7
                    switch -9223372036854775808 {
                    case -233:
                        var v10 int16 = 0
                        _ = v10
                    case 127:
                        if true {
                            delete(v4, 974)
                            var v11 int16 = v1
                            var v12 uint8 = 127
                            _ = v11
                            _ = v12
                        } else {
                            if (true || true) {
                                switch ("foo bar" + "") {
                                case "a":
                                    if (255 < v7) {
                                        v13 := v4[843]
                                        _ = v13
                                    }
                                case "":
                                    v2 = append(v2, 0)
                                }
                                v7 = v3
                                var v14 float32 = 1e10
                                _ = v14
                            }
                            var v15 int32 = -1
                            var v16 uint16 = (v7 << ((65535) & 63))
                            _ = v15
                            _ = v16
                        }
                        if true {
                            v5 += v5
                        } else {
                            fmt.Println(fi9, v3, v3)
                            v1 -= (v1 - v1)
                        }
                        v1 = (v1 + v1)
                    case 508:
                        if (-2147483648 < 0) {
                            delete(v4, 2147483648)
                            v1 = v1
                        }
                    default:
                        var v17 int8 = 101
                        if (v5 >= (v5 / v5)) {
                            v17 -= (v17 ^ int8(v1))
                            delete(v4, 1)
                        }
                        if true {
                            delete(v4, 42)
                            fmt.Println(fi9, v17)
                        }
                        v4[2] = (456 < -1)
                        _ = len(v4)
                        switch "hello" {
                        case "hello":
                            if true {
                                v18 := GenS1{F0: 1, F1: ((-166 < -883) && (42 < 127)), F2: v17}
                                var v19 int32 = 1
                                if (858 == fi8) {
                                    fmt.Println(fi9, v19)
                                } else {
                                    if (fi9 < fi8) {
                                        v19 = -662
                                        v2 = append(v2, 793)
                                        v20 := GenS1{F0: -1, F1: (true && ((3.14159 >= (((3.14159 / 0.1) + 3.14159) * -2.25)) || (!(!true)))), F2: (v17 - v17)}
                                        _ = v20
                                    } else {
                                        v21 := GenS2{F0: 127, F1: 42}
                                        var v22 int64 = 9223372036854775807
                                        fmt.Println(v19, v22)
                                        _ = v21
                                        _ = v22
                                    }
                                    _ = len(v2)
                                }
                                v23 := GenS2{F0: 42, F1: 1}
                                _ = v18
                                _ = v19
                                _ = v23
                            }
                        case "foo bar":
                            v24 := GenS1{F0: (v1 / (((v1 >> ((18446744073709551615) & 63))) | 1)), F1: false, F2: 100}
                            _ = v24
                        }
                        _ = v17
                    }
                }
                _ = fi9
            }
        case "foo bar":
            v25 := GenS1{F0: v1, F1: (91 >= 2), F2: int8(fi8)}
            _ = v25
        }
        fmt.Println(v1)
        _ = fi8
    }
    if true {
        var v26 uint32 = 4294967295
        fmt.Println(v1, v26, v5)
        for fi27 := 0; fi27 < 3; fi27++ {
            _ = len(v4)
            fmt.Println(v5, v5)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi27
        }
        _ = v26
    } else {
        var v28 float32 = 1e10
        var v29 float64 = (1e10 / 1e10)
        var v30 int16 = v1
        _ = v28
        _ = v29
        _ = v30
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v31 float64 = (float64(float32(float64(3.14159))) * (3.14159 - 0.0))
    fmt.Println(v1, v5)
    fmt.Println(v3, v31, v7)
    v32 := GenS3{F0: -1}
    var v33 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
