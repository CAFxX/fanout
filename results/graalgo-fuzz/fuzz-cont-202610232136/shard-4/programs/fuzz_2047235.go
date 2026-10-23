package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 int8; F2 uint16 }
type GenS6 struct { F0 int; F1 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int16 }

func main() {
    var v1 bool = true
    var v2 uint16 = 255
    v3 := []int{-2147483648}
    var v4 uint8 = 42
    var v5 float32 = -2.25
    var v6 float64 = float64(0.1)
    if true {
        var v7 int = 127
        var v8 uint8 = (v4 & (v4 / (((v4 + 223)) | 1)))
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 407, F1: uint64(v2), F2: 778}
        v11 := GenS2{F0: 0, F1: "hello"}
        v1 = v1
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS3{F0: 2}
    if (v1 && v1) {
        for fi13 := 0; fi13 < 10; fi13++ {
            v1 = (2 >= 994)
            switch -2147483648 {
            case 739:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            var v14 int8 = -128
            _ = fi13
            _ = v14
        }
        _ = len(v3)
    }
    v1 = (65535 > 764)
    v5 = 1.5
    if (0.0 <= v5) {
        switch ("foo bar" + "\n\t") {
        case "":
            var v15 bool = (v1 && true)
            _ = v15
        case "foo bar":
            v16 := GenS1{F0: 18446744073709551615, F1: 42, F2: 18446744073709551615}
            _ = v16
        case "hello":
            v17 := GenS3{F0: (v4 ^ uint8(v2))}
            _ = v17
        default:
            v3 = append(v3, 60)
        }
        var v18 uint32 = 0
        _ = v18
    }
    fmt.Println(v5, v2)
    v2 = (^(v2))
    if (float32(v6) != float32(((v6 / v6) * v6))) {
        v19 := GenS4{F0: -8}
        v1 = true
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 uint16 = v2
            if v1 {
                if false {
                    if false {
                        var v22 float32 = (v5 * v5)
                        fmt.Println(v1, fi20)
                        _ = v22
                    } else {
                        v1 = true
                    }
                    if len(v3) > 0 {
                        v3[0] = fi20
                    }
                    switch (fi20 ^ fi20) {
                    case -9223372036854775808:
                        v6 = (float64((v5 / 1.0)) - v6)
                    default:
                        if (2 > 1) {
                            fmt.Println(v4, fi20, v5)
                            if (v4 != (^(v4))) {
                                v23 := GenS5{F0: -32768, F1: -103, F2: v21}
                                switch "hello" {
                                case "a":
                                    if v1 {
                                        var v24 float64 = (float64(v5) / float64(v5))
                                        var v25 uint32 = 1
                                        for fi26 := 0; fi26 < 1; fi26++ {
                                            fmt.Println(v5)
                                            var v27 int = (fi26 ^ fi20)
                                            _ = fi26
                                            _ = v27
                                        }
                                        _ = v24
                                        _ = v25
                                    } else {
                                        v5 -= v5
                                        var v28 int16 = 0
                                        for fi29 := 0; fi29 < 17; fi29++ {
                                            var v30 int32 = 0
                                            var v31 uint16 = (^(v2))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = fi29
                                            _ = v30
                                            _ = v31
                                        }
                                        _ = v28
                                    }
                                case "hello":
                                    v5 += (((0.0 * v5) / 1.5) + v5)
                                }
                                _ = v23
                            }
                        } else {
                            var v32 uint64 = 1
                            fmt.Println(fi20, v5, v1)
                            v33 := GenS6{F0: 32768, F1: v4}
                            _ = v32
                            _ = v33
                        }
                    }
                } else {
                    if (0 == 0) {
                        if v1 {
                            var v34 int8 = 1
                            _ = v34
                        }
                    }
                    if false {
                        fmt.Println(v1, v4, v21)
                        var v35 uint32 = 1
                        if (int32(fi20) == 2147483647) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            for fi36 := 0; fi36 < 8; fi36++ {
                                var v37 uint64 = 18446744073709551615
                                if false {
                                    var v38 int16 = 1
                                    _ = v38
                                }
                                v3 = append(v3, (fi36 / ((fi36) | 1)))
                                _ = fi36
                                _ = v37
                            }
                        } else {
                            _ = len(v3)
                            fmt.Println(v5, v21, v5)
                        }
                        _ = v35
                    }
                }
            }
            for fi39 := 0; fi39 < 20; fi39++ {
                var v40 bool = true
                if (!false) {
                    var v41 uint32 = 65535
                    _ = v41
                } else {
                    if (1 <= 943) {
                        v40 = (!(((v6 * float64(0.0)) + v6) == v6))
                        _ = len(v3)
                        v42 := GenS7{F0: 0}
                        _ = v42
                    } else {
                        v43 := GenS8{F0: 42}
                        _ = v43
                    }
                    var v44 bool = true
                    _ = v44
                }
                v2 = 231
                _ = fi39
                _ = v40
            }
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
