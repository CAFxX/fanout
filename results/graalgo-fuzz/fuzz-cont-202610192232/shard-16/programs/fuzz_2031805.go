package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint = 601
    var v2 uint64 = 1
    v3 := []int{-744}
    v4 := make(map[int]uint32)
    for fi5 := 0; fi5 < 7; fi5++ {
        v1 = (v1 % ((2) | 1))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v2 -= (^(v2))
        _ = fi5
    }
    v1 = (v1 / ((v1) | 1))
    var v6 int8 = -128
    delete(v4, 638)
    _ = len(v4)
    var v7 uint32 = 2
    switch (("a" + "a") + "\n\t") {
    case "hello":
        v7 = 2
    }
    if true {
        var v8 uint = (v1 << (((v1 - v1)) & 63))
        if (false && (v1 != (^(v8)))) {
            if (!true) {
                fmt.Println(v8, v1)
            } else {
                v1 += 0
                v8 += v8
                v9 := GenS1{F0: -771, F1: (v6 * v6), F2: false}
                _ = v9
            }
            v1 -= (v1 + (^(v8)))
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            if ((0.0 / 1.0) < ((-2.25 / -2.25) - ((1.5 / 0.1) - float32(((1e10 + -2.25) - (1.5 / 0.1)))))) {
                var v11 int16 = 587
                v12 := GenS2{F0: v6}
                for fi13 := 0; fi13 < 19; fi13++ {
                    if false {
                        fmt.Println(v11, v11)
                        v2 += 1
                        switch -2147483648 {
                        case 452:
                            var v14 uint64 = 4294967295
                            _ = v14
                        case 42:
                            v1 = v1
                        case 128:
                            v6 += -1
                        }
                        if (!true) {
                            delete(v4, 9223372036854775807)
                            switch (fi10 % (((fi10 - (fi10 ^ fi13))) | 1)) {
                            case 42:
                                fmt.Println(v2, v1, v6)
                            case 127:
                                v15 := GenS2{F0: 100}
                                _ = v15
                            default:
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                            }
                            fmt.Println(fi10, v2, v2)
                        } else {
                            var v16 int8 = (v6 % ((v6) | 1))
                            if false {
                                v8 -= 42
                                var v17 float32 = -2.25
                                var v18 int = (fi13 | 32768)
                                _ = v17
                                _ = v18
                            } else {
                                if false {
                                    v3 = append(v3, fi13)
                                    switch fi13 {
                                    case -312:
                                        v19 := GenS1{F0: (fi10 - 395), F1: (v16 % ((42) | 1)), F2: (!(128 == 21))}
                                        _ = v19
                                    default:
                                        v8 -= v1
                                    }
                                }
                                var v20 uint32 = (^(v7))
                                v8 += (v8 ^ (v8 + 2))
                                var v21 uint32 = v7
                                fmt.Println(v6, v7)
                                _ = v20
                                _ = v21
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v16
                        }
                    } else {
                        fmt.Println(v7, v8)
                        _ = len(v3)
                        fmt.Println(fi10, v2, v11)
                    }
                    v7 += v7
                    v22 := GenS1{F0: (fi13 ^ fi13), F1: -128, F2: true}
                    _ = fi13
                    _ = v22
                }
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        _ = v8
    }
    v23 := v4[-2147483648]
    v2 -= 14
    var v24 uint = (v1 ^ v1)
    if ((127 != -1) || (false && (!(false || (-324 >= 0))))) {
        if (v6 >= 0) {
            v25 := GenS3{F0: 42, F1: -485}
            var v26 int16 = -13
            _ = v25
            _ = v26
        }
    } else {
        var v27 uint64 = 2
        var v28 int = -2147483648
        v24 = v1
        _ = v27
        _ = v28
    }
    if (32767 != -9223372036854775808) {
        v29 := GenS4{F0: v2}
        _ = v29
    }
    var v30 int32 = 953
    v31 := GenS3{F0: uint8(v7), F1: 9223372036854775807}
    var v32 uint16 = 42
    var v33 uint16 = 932
    fmt.Println(v30, v7, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v6
    _ = v7
}
