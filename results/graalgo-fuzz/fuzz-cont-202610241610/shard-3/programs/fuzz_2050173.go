package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    var v1 uint8 = 1
    var v2 int = 2147483647
    v3 := []uint32{}
    v4 := []uint32{2147483647}
    v5 := make(map[int]string)
    v6 := make(map[uint32]string)
    switch v2 {
    case 2147483647:
        v7 := GenS1{F0: 0, F1: ((("" + ("foo bar" + ("a" + "a"))) + "hello") + (("\n\t" + "hello") + ("\n\t" + "hello"))), F2: (v1 << ((543) & 63))}
        _ = v7
    default:
        fmt.Println(v1, v1)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        var v8 int8 = -54
        v8 = 42
        _ = v8
    }
    switch "\n\t" {
    case "foo bar":
        v2 += (v2 & (v2 >> ((42) & 63)))
    }
    fmt.Println(v1, v2)
    v1 = 128
    if (true && ((0.0 - ((1e10 + 0.0) - (1e10 + 0.0))) >= 1e10)) {
        v1 -= (v1 & (v1 / ((0) | 1)))
        v1 -= 140
        v6[2147483648] = (("foo bar" + "a") + "")
    } else {
        v2 -= v2
        for fi9 := 0; fi9 < 9; fi9++ {
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                var v10 uint64 = 2
                _ = v10
            }
            var v11 int = 2147483647
            v3 = append(v3, 1)
            _ = fi9
            _ = v11
        }
    }
    v2 -= 2147483647
    v1 = v1
    var v12 float64 = (3.14159 / ((0.1 * 1.5) - 0.1))
    for fi13 := 0; fi13 < 3; fi13++ {
        fmt.Println(fi13, fi13, v12)
        for fi14 := 0; fi14 < 8; fi14++ {
            _ = len(v3)
            _ = fi14
        }
        var v15 int8 = 42
        _ = fi13
        _ = v15
    }
    v16 := GenS2{F0: -1, F1: 1}
    fmt.Println(v12, v1)
    v1 -= (^(v1))
    switch (v2 - v2) {
    case -9223372036854775808:
        var v17 string = (("a" + "a") + "\n\t")
        _ = v17
    case 4294967295:
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 uint16 = 32767
            var v20 uint32 = 990
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 uint16 = (v19 >> ((uint(v19)) & 63))
                _ = fi21
                _ = v22
            }
            _ = fi18
            _ = v19
            _ = v20
        }
    case 42:
        var v23 uint = 42
        _ = v23
    default:
        v12 += (3.14159 * 0.1)
    }
    v24 := GenS2{F0: -142, F1: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
