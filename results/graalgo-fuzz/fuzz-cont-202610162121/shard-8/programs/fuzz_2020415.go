package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 int16 }
func genFn0() uint64 {
    var v4 uint32 = 529
    for fi5 := 0; fi5 < 4; fi5++ {
        if true {
            v6 := GenS1{F0: (v4 ^ (v4 >> ((694) & 63))), F1: 127}
            _ = v6
        }
        var v7 float64 = -2.25
        fmt.Println(fi5, v4, v4)
        _ = fi5
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        switch ("\n\t" + (((((("" + "hello") + "foo bar") + ("foo bar" + "")) + "hello") + ("foo bar" + "foo bar")) + (("hello" + "\n\t") + ""))) {
        case "hello":
            v4 = v4
        case "":
            var v9 int64 = 9223372036854775807
            _ = v9
        }
        var v10 float32 = (float32((1e10 / 1.5)) * 1.5)
        _ = fi8
        _ = v10
    }
    v4 += v4
    _ = v4
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 128
    var v3 bool = ("" < "foo bar")
    v11 := []int{-824, -9223372036854775808, 42}
    v2 = 42
    var v12 int16 = -795
    if (0 != -366) {
        _ = len(v11)
    }
    if (!(!true)) {
        v13 := v1[uint32(v2)]
        var v14 int16 = v12
        _ = v13
        _ = v14
    }
    fmt.Println(v12, v3)
    v2 -= 150
    if v3 {
        v3 = false
        fmt.Println(v2)
    }
    v15 := GenS1{F0: 0, F1: 64}
    var v16 uint = 481
    v12 = (v12 ^ v12)
    v2 -= 2
    v1[260] = "foo bar"
    v17 := GenS1{F0: uint32(v2), F1: -50}
    v18 := GenS1{F0: 2147483648, F1: 0}
    switch "a" {
    case "hello":
        for fi19 := 0; fi19 < 16; fi19++ {
            v12 += (v12 | (v12 * (v12 - (v12 | 1))))
            _ = fi19
        }
    case "a":
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS2{F0: (v12 << ((255) & 63))}
            if (-9223372036854775808 == int64(v2)) {
                fmt.Println(v2, fi20)
                v22 := GenS1{F0: 1, F1: 62}
                if (v3 || ((0 <= 1) || (177 >= 4294967295))) {
                    v16 = (^(v16))
                }
                _ = v22
            } else {
                v12 += v12
                v16 -= 42
                v23 := GenS2{F0: -1}
                _ = v23
            }
            _ = fi20
            _ = v21
        }
    case "":
        var v24 uint16 = 237
        _ = v24
    }
    v2 -= v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
