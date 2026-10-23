package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint32 }
func genFn0(p0 int, p1 string, p2 uint8) int32 {
    fmt.Println(p2)
    fmt.Println(p2, p1)
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(p1, p2)
        var v7 uint = 18446744073709551615
        _ = fi6
        _ = v7
    }
    if false {
        if false {
            p1 = (p1 + (p1 + p1))
            var v8 uint16 = 503
            for fi9 := 0; fi9 < 4; fi9++ {
                v8 += v8
                v10 := []int{(fi9 ^ int(v8))}
                _ = fi9
                _ = v10
            }
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]int)
    var v3 int32 = 559
    var v4 int8 = -1
    var v5 uint64 = 456
    var v11 int = 128
    switch (v11 >> ((255) & 63)) {
    case 42:
        var v12 uint8 = 42
        _ = v12
    case 32767:
        v3 += (v3 >> ((42) & 63))
    case -730:
        switch "\n\t" {
        case "\n\t":
            v13 := GenS1{F0: -699}
            _ = v13
        default:
            fmt.Println(v5)
        }
    default:
        for fi14 := 0; fi14 < 10; fi14++ {
            switch "" {
            case "hello":
                if (v5 == (v5 * (v5 << ((0) & 63)))) {
                    var v15 int8 = (-(v4))
                    v16 := []int{-927, fi14, (fi14 ^ (fi14 % ((int(v3)) | 1)))}
                    _ = v15
                    _ = v16
                } else {
                    var v17 float32 = 3.14159
                    _ = v17
                }
            case "":
                v18 := GenS2{F0: 2147483647, F1: 277}
                _ = v18
            case "foo bar":
                v4 += (v4 ^ (^(v4)))
            }
            _ = fi14
        }
    }
    var v19 int32 = v3
    if (42 < 42) {
        fmt.Println(v4, v5, v3)
        v20 := []int{}
        _ = v20
    }
    v1 += v1
    switch ("foo bar" + "\n\t") {
    case "":
        for fi21 := 0; fi21 < 5; fi21++ {
            delete(v2, v11)
            fmt.Println(v3, v4, v3)
            _ = fi21
        }
    case "hello":
        var v22 int8 = v4
        _ = v22
    case "\n\t":
        var v23 int32 = -2147483648
        _ = v23
    default:
        if false {
            var v24 int8 = (v4 * (v4 / (((-(v4))) | 1)))
            for fi25 := 0; fi25 < 18; fi25++ {
                v4 = -128
                v4 -= (^(v4))
                _ = fi25
            }
            _ = v24
        }
    }
    if ((!(1e10 != 1.5)) || (false && true)) {
        fmt.Println(v3)
        v26 := GenS3{F0: 42, F1: 239, F2: 576}
        v27 := GenS1{F0: 42}
        _ = v26
        _ = v27
    } else {
        v4 -= (v4 << ((4294967295) & 63))
        var v28 int = (v11 << ((42) & 63))
        _ = v28
    }
    var v29 float64 = -2.25
    v30 := GenS2{F0: 887, F1: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
