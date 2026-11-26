package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int64 }
type GenS2 struct { F0 bool; F1 int64; F2 int8 }
func genFn0(p0 string) uint64 {
    p0 = (p0 + (("foo bar" + "") + p0))
    if ((((p0 + (p0 + ("foo bar" + p0))) + p0) + "\n\t") > p0) {
        p0 = (p0 + "foo bar")
        switch 4294967295 {
        case 42:
            var v5 int8 = 100
            _ = v5
        case -1:
            switch 32767 {
            case 127:
                p0 = ("foo bar" + (p0 + (p0 + p0)))
            default:
                fmt.Println(p0)
            }
        case 1:
            p0 = p0
        }
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := []int{(fi6 * (fi6 >> ((716) & 63))), fi6, fi6}
            fmt.Println(fi6)
            _ = fi6
            _ = v7
        }
    }
    v8 := []int{761}
    v9 := GenS1{F0: 2, F1: 0, F2: 0}
    _ = p0
    _ = v8
    _ = v9
    return 638
}

func main() {
    var v1 uint8 = 255
    var v2 uint16 = 0
    v3 := make(map[uint32]int)
    var v4 int64 = -369
    if true {
        var v10 float32 = float32(1.5)
        var v11 uint64 = 272
        _ = v10
        _ = v11
    }
    if (!true) {
        var v12 uint16 = 32767
        _ = v12
    } else {
        var v13 uint = 246
        fmt.Println(v2, v13)
        v14 := []int{-468, 32768, -1, -2147483648}
        _ = v13
        _ = v14
    }
    switch 127 {
    case 127:
        v2 = 518
    case -9223372036854775808:
        v15 := []int{485, 845}
        _ = v15
    case -1:
        v1 += 6
    }
    var v16 float64 = (float64(1.5) * 0.1)
    var v17 int64 = (v4 >> ((255) & 63))
    fmt.Println(v16)
    v17 -= (v17 % (((^(v4))) | 1))
    v2 -= (v2 >> ((255) & 63))
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 uint = 167
        switch (fi18 / ((-215) | 1)) {
        case 128:
            for fi20 := 0; fi20 < 16; fi20++ {
                switch fi20 {
                case 295:
                    var v21 int8 = 42
                    _ = v21
                }
                _ = fi20
            }
        default:
            if true {
                v17 += 42
                v22 := GenS2{F0: (42 != 0), F1: 9223372036854775807, F2: -128}
                fmt.Println(fi18)
                _ = v22
            }
        }
        v2 += (v2 | v2)
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
