package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16; F1 uint; F2 uint16 }

func main() {
    var v1 int32 = 1
    v2 := []uint32{0, 974, 2}
    v3 := []int{-1, 128, 607, 2147483647}
    v4 := make(map[uint32]uint32)
    fmt.Println(v1)
    v5 := GenS1{F0: 707, F1: -9223372036854775808}
    if false {
        fmt.Println(v1)
        var v6 uint16 = 65535
        _ = v6
    }
    var v7 uint16 = 977
    v1 = (^(v1))
    v2 = append(v2, 0)
    _ = len(v3)
    if true {
        if (32767 < 543) {
            var v8 int8 = -83
            switch "hello" {
            case "foo bar":
                var v9 uint8 = 0
                _ = v9
            default:
                v10 := GenS2{F0: -32768, F1: 42, F2: 65535}
                _ = v10
            }
            _ = v8
        }
    }
    if (true && ((true && (!false)) && (962 != 65535))) {
        v1 -= (-(v1))
        v1 += 42
        var v11 string = "hello"
        _ = v11
    } else {
        v1 = (v1 >> ((18446744073709551615) & 63))
    }
    fmt.Println(v1)
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: v7, F1: (fi12 + 0)}
        if (!true) {
            v1 = v1
            v1 = v1
            v1 += v1
        } else {
            var v14 uint8 = 185
            fmt.Println(v14)
            if true {
                fmt.Println(fi12)
            }
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        if (209 <= 42) {
            fmt.Println(v7)
            v7 = v7
            fmt.Println(v7, fi15)
        } else {
            if true {
                if false {
                    v16 := GenS1{F0: (v7 >> ((722) & 63)), F1: (fi15 << ((0) & 63))}
                    _ = v16
                } else {
                    _ = len(v4)
                }
                fmt.Println(v7, v7)
                v1 = int32(v7)
            } else {
                for fi17 := 0; fi17 < 19; fi17++ {
                    if (89 < int8(fi15)) {
                        v18 := GenS1{F0: (^(v7)), F1: (fi15 / (((-(fi17))) | 1))}
                        fmt.Println(v7, fi17)
                        _ = v18
                    } else {
                        v19 := GenS1{F0: (v7 ^ (^(v7))), F1: (fi15 - fi17)}
                        v1 -= -1
                        _ = v19
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi17
                }
                v7 = v7
                for fi20 := 0; fi20 < 11; fi20++ {
                    fmt.Println(v7, fi20, fi15)
                    _ = fi20
                }
            }
        }
        _ = fi15
    }
    fmt.Println(v1, v1)
    v7 = (^(v7))
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 uint64 = 2
        var v23 bool = (!false)
        v24 := GenS2{F0: -32768, F1: 1, F2: (v7 ^ 32767)}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
