package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 uint }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := []int{9223372036854775807, 127, 42, 0}
    var v2 uint16 = 809
    v3 := make(map[string]string)
    var v4 float32 = float32(0.0)
    v5 := make(map[uint32]int)
    v6 := []uint32{}
    fmt.Println(v2)
    var v7 int16 = -1
    switch 32767 {
    case -2147483648:
        v8 := v3[""]
        _ = v8
    case 1:
        var v9 uint32 = 685
        _ = v9
    }
    delete(v5, 790)
    var v10 uint = 950
    v11 := GenS1{F0: false, F1: 130, F2: v10}
    switch "hello" {
    case "":
        switch 9223372036854775807 {
        case 9223372036854775807:
            var v12 int = 128
            _ = v12
        case -9223372036854775808:
            var v13 uint64 = 1
            _ = v13
        case 2147483647:
            if false {
                if (32767 != v2) {
                    v2 -= (v2 & (^(v2)))
                    var v14 uint = (v10 / (((v10 >> ((v10) & 63))) | 1))
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v16 := GenS1{F0: ("foo bar" >= "\n\t"), F1: 0, F2: v10}
                        for fi17 := 0; fi17 < 8; fi17++ {
                            if (true || true) {
                                if true {
                                    if (!false) {
                                        if (false && true) {
                                            if true {
                                                v10 = (v10 << ((2) & 63))
                                                var v18 uint16 = v2
                                                delete(v5, 4294967295)
                                                _ = v18
                                            } else {
                                                fmt.Println(v10)
                                                v19 := GenS1{F0: (fi15 < 0), F1: 972, F2: 2}
                                                v14 -= (v10 >> (((v14 >> ((0) & 63))) & 63))
                                                _ = v19
                                            }
                                        } else {
                                            switch ("a" + "hello") {
                                            case "\n\t":
                                                var v20 float64 = 0.1
                                                _ = v20
                                            default:
                                                v21 := GenS1{F0: false, F1: 4294967295, F2: v14}
                                                _ = v21
                                            }
                                            v22 := GenS2{F0: 364}
                                            var v23 int16 = (v7 << (((v14 + (v14 << (((v14 | 0)) & 63)))) & 63))
                                            _ = v22
                                            _ = v23
                                        }
                                        switch (fi15 >> ((0) & 63)) {
                                        case 0:
                                            if true {
                                                v14 = (v10 / ((v10) | 1))
                                                switch "a" {
                                                case "\n\t":
                                                    fmt.Println(v14, v7, fi15)
                                                case "foo bar":
                                                    _ = len(v6)
                                                default:
                                                    var v24 int32 = -1
                                                    _ = v24
                                                }
                                                if false {
                                                    v7 = (v7 | v7)
                                                    v25 := v5[0]
                                                    _ = v25
                                                }
                                            }
                                        }
                                        switch fi17 {
                                        case 4294967295:
                                            var v26 uint32 = 65535
                                            _ = v26
                                        case -321:
                                            if false {
                                                v4 += -2.25
                                                _ = len(v3)
                                                v27 := GenS2{F0: 42}
                                                _ = v27
                                            }
                                        default:
                                            var v28 uint8 = uint8(v14)
                                            if true {
                                                v7 = (v7 << ((342) & 63))
                                                fmt.Println(v10)
                                            }
                                            _ = v28
                                        }
                                        v4 = 0.0
                                    } else {
                                        v1 = append(v1, (^(fi15)))
                                    }
                                    var v29 int32 = 1
                                    var v30 int32 = v29
                                    v10 -= v10
                                    if (878 <= 549) {
                                        fmt.Println(fi17)
                                        fmt.Println(v10)
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v10 += 65535
                                        fmt.Println(v7)
                                    }
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 float32 = 1e10
                                    _ = v31
                                }
                                v32 := GenS1{F0: true, F1: 2147483648, F2: v14}
                                var v33 int16 = (^(v7))
                                _ = v32
                                _ = v33
                            }
                            var v34 uint = (v10 | (v14 & v10))
                            if false {
                                var v35 int8 = -52
                                var v36 int64 = 1
                                var v37 int32 = -2147483648
                                var v38 uint64 = 434
                                _ = v35
                                _ = v36
                                _ = v37
                                _ = v38
                            } else {
                                var v39 float32 = 1e10
                                var v40 float64 = (1e10 - 3.14159)
                                _ = v39
                                _ = v40
                            }
                            _ = fi17
                            _ = v34
                        }
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                }
            }
        }
    }
    v7 += v7
    fmt.Println(v4, v7, v7)
    var v41 uint64 = 2
    switch 42 {
    case 194:
        v10 = v10
    case 9223372036854775807:
        delete(v5, 318)
    default:
        v7 = (v7 * (v7 | int16(v2)))
    }
    switch "foo bar" {
    case "hello":
        if true {
            var v42 uint32 = uint32(v41)
            _ = v42
        } else {
            v7 = v7
        }
    case "\n\t":
        if (475 >= 42) {
            v43 := GenS2{F0: v41}
            _ = v43
        } else {
            v7 += v7
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
    _ = v7
}
