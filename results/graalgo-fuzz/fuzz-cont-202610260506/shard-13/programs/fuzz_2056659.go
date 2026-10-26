package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 uint8 = 2
    v2 := []int{-320, 1}
    var v3 int8 = 1
    var v4 int64 = 187
    v5 := make(map[uint32]uint32)
    var v6 int8 = -128
    v3 = -67
    if true {
        var v7 uint32 = 603
        var v8 string = "foo bar"
        _ = v7
        _ = v8
    }
    v1 -= v1
    if true {
        v9 := GenS1{F0: 1, F1: 797, F2: -345}
        _ = v9
    }
    var v10 int32 = -391
    v11 := GenS1{F0: 42, F1: 2, F2: 9223372036854775807}
    for fi12 := 0; fi12 < 11; fi12++ {
        switch -1 {
        case 104:
            switch (^(fi12)) {
            case 32767:
                switch fi12 {
                case 434:
                    v13 := GenS2{F0: "hello", F1: (v10 - v10)}
                    _ = v13
                case 32767:
                    for fi14 := 0; fi14 < 2; fi14++ {
                        switch fi14 {
                        case 9223372036854775807:
                            if true {
                                var v15 uint16 = 2
                                _ = v15
                            }
                        default:
                            v16 := GenS3{F0: 599}
                            _ = v16
                        }
                        if ((676 > 4294967295) || false) {
                            _ = len(v2)
                        }
                        _ = fi14
                    }
                case -462:
                    for fi17 := 0; fi17 < 17; fi17++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi17
                    }
                default:
                    fmt.Println(fi12, v6, v10)
                }
            case 128:
                v18 := GenS4{F0: 65535, F1: 4294967295}
                _ = v18
            default:
                for fi19 := 0; fi19 < 19; fi19++ {
                    if (-1 >= v4) {
                        var v20 int64 = v4
                        var v21 uint16 = 2
                        _ = v20
                        _ = v21
                    } else {
                        v1 = 128
                        var v22 int64 = -1
                        if (!true) {
                            if true {
                                var v23 int = -743
                                if (1 != (v1 & (v1 / ((v1) | 1)))) {
                                    v6 -= int8(fi12)
                                    var v24 bool = (false && (46 < 4294967295))
                                    _ = v24
                                } else {
                                    if (1e10 > 1e10) {
                                        _ = len(v2)
                                        fmt.Println(v3, v4)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    var v25 int8 = 67
                                    _ = v25
                                }
                                _ = v23
                            }
                            fmt.Println(fi19, v22)
                            var v26 uint64 = 42
                            v22 = (v22 >> ((uint(v1)) & 63))
                            v22 = 1
                            _ = v26
                        } else {
                            var v27 int8 = v3
                            _ = v27
                        }
                        _ = v22
                    }
                    _ = len(v5)
                    if ((0 < 654) && false) {
                        v1 -= (v1 << ((1000) & 63))
                        v3 = (^(v6))
                    } else {
                        v6 -= (v6 | (v3 + (v6 % ((v3) | 1))))
                        _ = len(v5)
                        v5[2147483647] = 2147483647
                        v6 += v6
                        v10 = (^(v10))
                    }
                    _ = fi19
                }
            }
        case 127:
            v28 := GenS3{F0: 32767}
            _ = v28
        case 1:
            v4 = 42
        }
        var v29 float64 = 0.1
        v6 = int8(fi12)
        _ = fi12
        _ = v29
    }
    _ = len(v5)
    v5[573] = 759
    v4 += int64(v10)
    if true {
        if true {
            switch ((("hello" + ("a" + (("foo bar" + ("a" + "foo bar")) + (("\n\t" + ("" + "hello")) + "a")))) + "foo bar") + "hello") {
            case "a":
                var v30 float64 = 1e10
                _ = v30
            default:
                for fi31 := 0; fi31 < 2; fi31++ {
                    var v32 uint32 = 0
                    _ = fi31
                    _ = v32
                }
            }
            v2 = append(v2, 2147483647)
        }
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v4 = (-(v4))
        case "foo bar":
            v10 = 0
        }
    }
    v2 = append(v2, -2147483648)
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
