package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int64; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 bool }
func genFn0(p0 int16, p1 uint32, p2 bool) int8 {
    switch 127 {
    case 0:
        v6 := []int{32768, 709, 2147483647, -543}
        _ = v6
    default:
        p2 = p2
    }
    var v7 uint64 = 4294967295
    v7 += v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 int8 = -82
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    var v4 uint16 = 32767
    v5 := []int{543, 128, -834, -6}
    fmt.Println(v1, v4)
    if (-1 >= 1) {
        var v8 string = ("hello" + "hello")
        switch (v8 + v8) {
        case "hello":
            v1 = (v1 << ((4294967295) & 63))
        case "foo bar":
            var v9 int16 = -32768
            _ = v9
        }
        if (133 > 2) {
            if ((-1 > -408) && false) {
                v1 += (-(v1))
                if false {
                    v4 += 807
                    delete(v2, 2)
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v1 = genFn0(-1, 767, true)
                    var v10 uint8 = 42
                    _ = v10
                }
                var v11 uint64 = 4294967295
                _ = v11
            }
            fmt.Println(v1)
            v12 := v2[402]
            _ = v12
        } else {
            var v13 uint32 = 2147483648
            switch v8 {
            case "a":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                v14 := GenS1{F0: v8}
                _ = v14
            }
            v4 += v4
            _ = v13
        }
        _ = v8
    }
    var v15 float64 = (-2.25 - 3.14159)
    v16 := GenS2{F0: -1, F1: 42, F2: 2}
    fmt.Println(v1, v4, v1)
    switch "foo bar" {
    case "foo bar":
        v15 -= 3.14159
    case "\n\t":
        v1 = v1
    case "hello":
        v1 = v1
    }
    var v17 float32 = -2.25
    fmt.Println(v15, v17)
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS3{F0: -930}
        _ = fi18
        _ = v19
    }
    if (!false) {
        v3[-2147483648] = 2147483648
        var v20 string = ("hello" + "hello")
        _ = v20
    } else {
        if ((false || true) && false) {
            v15 = float64((v17 / float32(((v15 - v15) / v15))))
        } else {
            v21 := GenS4{F0: -936, F1: (1 < -36)}
            _ = v21
        }
        var v22 uint64 = 42
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
