package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint32; F1 string; F2 int32 }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 bool; F1 bool }

func main() {
    var v1 bool = false
    v2 := []int{32768}
    var v3 float32 = float32(3.14159)
    v4 := make(map[uint32]int)
    var v5 uint32 = 65535
    var v6 bool = ((0.1 + -2.25) != (v3 - v3))
    var v7 bool = v1
    var v8 float64 = float64((v3 - float32(0.0)))
    fmt.Println(v7, v8, v7)
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            _ = len(v2)
            v5 += (v5 * (v5 | (v5 >> ((65535) & 63))))
            v7 = true
            if true {
                switch "\n\t" {
                case "a":
                    var v11 uint8 = 42
                    _ = v11
                case "foo bar":
                    if (675 == 893) {
                        v12 := GenS1{F0: 190, F1: (fi10 % ((39) | 1))}
                        v6 = ((-1 >= -36) || (-79 != 1000))
                        var v13 bool = false
                        _ = v12
                        _ = v13
                    } else {
                        v1 = (2 >= 840)
                        v6 = false
                        var v14 int = 459
                        _ = v14
                    }
                    v5 = v5
                    if v7 {
                        v15 := v4[v5]
                        _ = v15
                    } else {
                        v4[4294967295] = 128
                        v16 := GenS1{F0: 2, F1: 2147483647}
                        switch (-(fi10)) {
                        case 0:
                            switch (fi10 | (fi10 * fi10)) {
                            case 0:
                                delete(v4, v5)
                            default:
                                var v17 uint8 = 0
                                _ = v17
                            }
                        case 1:
                            v18 := GenS1{F0: 218, F1: fi9}
                            _ = v18
                        case 32767:
                            v7 = (!((1e10 > (((float64(0.0) - v8) + v8) + float64(((v3 * float32((v8 / v8))) - 0.0)))) && false))
                        default:
                            switch 1 {
                            case 2147483647:
                                v5 = (v5 >> ((489) & 63))
                            }
                        }
                        _ = v16
                    }
                case "hello":
                    var v19 int32 = 371
                    _ = v19
                default:
                    switch ("" + "\n\t") {
                    case "hello":
                        fmt.Println(v7, fi10)
                    case "\n\t":
                        v20 := GenS2{F0: (v5 & 943), F1: "", F2: -2147483648}
                        _ = v20
                    default:
                        if ((0 <= 0) || (v5 != v5)) {
                            fmt.Println(v1)
                            v5 -= v5
                        }
                    }
                }
                v21 := GenS3{F0: (v7 || ((!((!false) || v7)) || true)), F1: true}
                _ = v21
            } else {
                if true {
                    v6 = (!v1)
                    var v22 string = "hello"
                    if (!(!v1)) {
                        if v7 {
                            v23 := GenS4{F0: (v5 * (^(v5))), F1: -605}
                            v5 += 1
                            _ = v23
                        } else {
                            v8 = 0.0
                        }
                        var v24 uint = 42
                        _ = v24
                    } else {
                        if v7 {
                            var v25 uint64 = 2
                            v7 = true
                            fmt.Println(v22)
                            _ = v25
                        }
                        v4[v5] = (fi9 >> ((2) & 63))
                        if (v1 || (!false)) {
                            v8 = (float64(v3) * 0.0)
                        }
                    }
                    _ = v22
                }
            }
            v26 := GenS5{F0: v7, F1: false}
            _ = fi10
            _ = v26
        }
        _ = fi9
    }
    var v27 uint = 42
    for fi28 := 0; fi28 < 15; fi28++ {
        var v29 int32 = 1
        _ = fi28
        _ = v29
    }
    v27 = v27
    for fi30 := 0; fi30 < 9; fi30++ {
        switch "hello" {
        case "\n\t":
            v27 = uint(v5)
        case "hello":
            var v31 bool = ((v27 - 42) < (v27 & (v27 | v27)))
            _ = v31
        case "":
            v32 := GenS5{F0: (-865 < 32767), F1: (100 > 127)}
            _ = v32
        }
        _ = fi30
    }
    v5 = 4294967295
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
