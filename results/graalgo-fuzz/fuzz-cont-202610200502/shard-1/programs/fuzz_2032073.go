package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint8 = 2
    var v2 int64 = 509
    v3 := make(map[uint32]uint32)
    v2 -= (^(v2))
    fmt.Println(v1)
    for fi4 := 0; fi4 < 20; fi4++ {
        if true {
            var v5 float32 = (-2.25 - 1.5)
            _ = v5
        }
        v1 -= (v1 << ((4294967295) & 63))
        var v6 int32 = 0
        _ = fi4
        _ = v6
    }
    fmt.Println(v1, v2)
    v2 -= v2
    var v7 uint = 42
    _ = len(v3)
    v3[4294967295] = 466
    fmt.Println(v2)
    if true {
        if ((898 < 339) || (!true)) {
            if ((^(v7)) >= (v7 >> ((v7) & 63))) {
                fmt.Println(v7)
                v1 = (v1 | 0)
            } else {
                v8 := GenS1{F0: 1}
                _ = v8
            }
            v7 = (v7 / ((4294967295) | 1))
        }
        switch -1 {
        case 42:
            v1 = (v1 | 127)
        default:
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := v3[42]
                _ = fi9
                _ = v10
            }
        }
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            if true {
                for fi12 := 0; fi12 < 4; fi12++ {
                    fmt.Println(v1)
                    _ = fi12
                }
                v7 += 922
                if (65535 <= 32767) {
                    fmt.Println(v1, v2)
                    _ = len(v3)
                }
            } else {
                v7 = v7
                if (699 <= 65535) {
                    v2 = v2
                    if (-1 >= 113) {
                        var v13 uint16 = 255
                        _ = v13
                    }
                    var v14 uint64 = 2
                    _ = v14
                }
            }
            for fi15 := 0; fi15 < 15; fi15++ {
                if false {
                    fmt.Println(v2)
                    v7 = 4294967295
                    v2 = int64(v7)
                } else {
                    fmt.Println(v2, fi11, v2)
                }
                _ = fi15
            }
            v16 := GenS1{F0: -2147483648}
            _ = fi11
            _ = v16
        }
        v17 := GenS1{F0: 903}
        _ = v17
    }
    var v18 int = 128
    if true {
        switch -2147483648 {
        case -1:
            v19 := GenS2{F0: 42, F1: 2}
            _ = v19
        case -9223372036854775808:
            delete(v3, uint32(v7))
        case 128:
            v7 = v7
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 string = "a"
            if ((true && ((false && false) && (!true))) && true) {
                v3[185] = 598
                var v22 int8 = int8(v2)
                _ = v22
            }
            v3[2] = 0
            _ = fi20
            _ = v21
        }
        var v23 int64 = v2
        _ = v23
    }
    var v24 uint32 = 2147483647
    var v25 int = v18
    v7 += (v7 + (v7 + v7))
    if (810 == 42) {
        if (false && false) {
            var v26 float32 = (0.1 + ((0.0 + 1.5) + 0.1))
            var v27 float64 = 3.14159
            v24 = (^(v24))
            _ = v26
            _ = v27
        } else {
            v28 := GenS1{F0: 42}
            _ = v28
        }
    } else {
        v1 = v1
    }
    v3[4294967295] = uint32(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v7
}
