package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int8 }

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 uint = 65535
    var v5 int64 = (v1 >> ((v4) & 63))
    var v6 float32 = (1.5 / 0.1)
    var v7 uint16 = 214
    for fi8 := 0; fi8 < 2; fi8++ {
        switch "\n\t" {
        case "a":
            fmt.Println(fi8, v6, v1)
        case "hello":
            for fi9 := 0; fi9 < 18; fi9++ {
                v4 = v4
                v6 = -2.25
                _ = fi9
            }
        case "":
            var v10 uint8 = 255
            _ = v10
        default:
            v7 = (v7 + v7)
        }
        _ = fi8
    }
    v6 += ((v6 - v6) - v6)
    v11 := v2[991]
    if (!false) {
        v5 -= -9223372036854775808
    }
    var v12 uint8 = 22
    fmt.Println(v5)
    if ((v7 + v7) == (v7 - v7)) {
        fmt.Println(v4)
        fmt.Println(v1, v5, v6)
    } else {
        v11 = (v11 << (((v4 & 1)) & 63))
    }
    if (529 > 42) {
        v13 := []int{127, 4294967295, 42}
        if (true || false) {
            _ = len(v2)
        } else {
            v5 += (v1 ^ (^(v1)))
            v12 += v12
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS1{F0: v7, F1: (false && true), F2: int8(v12)}
                if ((-32768 <= 1) || ((32767 == fi14) || false)) {
                    v7 -= 1
                    switch (fi14 & fi14) {
                    case -401:
                        v11 = (v11 - 2147483648)
                    }
                    for fi16 := 0; fi16 < 1; fi16++ {
                        v6 -= 1.5
                        _ = fi16
                    }
                }
                if false {
                    var v17 uint8 = 127
                    var v18 float32 = float32(0.1)
                    var v19 float32 = -2.25
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v6 = v6
                        v21 := GenS1{F0: uint16(v4), F1: true, F2: 0}
                        _ = fi20
                        _ = v21
                    }
                    v5 -= 1
                    var v22 string = ("" + ("hello" + "\n\t"))
                    _ = v22
                }
                _ = fi14
                _ = v15
            }
        }
        _ = v13
    } else {
        if (v1 <= v1) {
            for fi23 := 0; fi23 < 15; fi23++ {
                var v24 uint64 = 1
                v25 := []int{fi23, int(v1), (fi23 << ((v4) & 63))}
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 int64 = -572
                _ = fi26
                _ = v27
            }
            v2[465] = uint32(v4)
        } else {
            var v28 string = ("hello" + "")
            fmt.Println(v5)
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
