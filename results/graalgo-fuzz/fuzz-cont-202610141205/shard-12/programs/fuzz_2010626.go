package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint16; F1 int32 }

func main() {
    v1 := []uint32{320, 145}
    var v2 int32 = 1
    v3 := make(map[uint32]int)
    var v4 uint32 = 207
    for fi5 := 0; fi5 < 20; fi5++ {
        v4 = 42
        _ = fi5
    }
    v2 -= v2
    v2 -= (^(v2))
    fmt.Println(v4)
    _ = len(v3)
    var v6 float32 = 1e10
    if false {
        var v7 int = -2147483648
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                switch fi8 {
                case 32767:
                    v10 := GenS1{F0: -1, F1: true}
                    if true {
                        var v11 int32 = v2
                        var v12 float32 = float32(float64(((float32(0.0) / v6) + v6)))
                        fmt.Println(fi9)
                        var v13 uint16 = 163
                        v14 := GenS1{F0: -477, F1: (!(false && (true && true)))}
                        var v15 int = 1
                        v16 := GenS2{F0: v4, F1: (false && true)}
                        v2 += v11
                        var v17 uint32 = v4
                        _ = v11
                        _ = v12
                        _ = v13
                        _ = v14
                        _ = v15
                        _ = v16
                        _ = v17
                    } else {
                        var v18 uint64 = 42
                        var v19 float32 = v6
                        v20 := GenS2{F0: (^(v4)), F1: (v18 <= 4294967295)}
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    var v21 uint8 = 128
                    switch ("hello" + "a") {
                    case "foo bar":
                        v22 := GenS3{F0: 166, F1: (v21 * 128), F2: 679}
                        _ = v22
                    }
                    fmt.Println(v2, fi8)
                    v21 -= (v21 & (v21 * (v21 << ((2) & 63))))
                    if (42 <= 37) {
                        var v23 int8 = 100
                        _ = v23
                    }
                    v24 := GenS4{F0: 255, F1: (v2 << ((255) & 63))}
                    _ = v10
                    _ = v21
                    _ = v24
                case 872:
                    v6 = 0.0
                case 42:
                    if false {
                        if false {
                            v4 = (v4 & (^(v4)))
                            if false {
                                var v25 uint64 = 1
                                _ = v25
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = (v4 | (v4 % ((v4) | 1)))
                                }
                            }
                        }
                        var v26 int16 = 0
                        _ = v26
                    } else {
                        fmt.Println(fi9, v4)
                    }
                }
                if true {
                    var v27 float32 = (v6 - v6)
                    switch "hello" {
                    case "foo bar":
                        v28 := GenS4{F0: 2, F1: 2147483647}
                        _ = v28
                    default:
                        var v29 int8 = -1
                        _ = v29
                    }
                    _ = v27
                }
                var v30 int = (fi8 >> ((990) & 63))
                var v31 string = ""
                v32 := v3[(v4 - 4294967295)]
                _ = fi9
                _ = v30
                _ = v31
                _ = v32
            }
            _ = fi8
        }
        var v33 uint64 = 42
        _ = v33
        _ = v7
    }
    v34 := GenS4{F0: uint16(v4), F1: (v2 * 273)}
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v6
}
