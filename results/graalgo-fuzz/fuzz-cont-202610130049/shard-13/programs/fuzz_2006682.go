package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int64; F2 int64 }

func main() {
    var v1 uint16 = 42
    var v2 int = -87
    var v3 uint32 = 2147483647
    var v4 float64 = 1.5
    var v5 uint8 = 255
    v6 := make(map[int]uint32)
    v7 := []int{}
    fmt.Println(v2, v5)
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v8 := v6[v2]
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    switch (("\n\t" + "") + "foo bar") {
    case "foo bar":
        var v9 uint32 = (v8 >> ((902) & 63))
        _ = v9
    case "a":
        if (!(-21 < 2147483647)) {
            delete(v6, (v2 << ((983) & 63)))
            var v10 float64 = v4
            v11 := GenS1{F0: v1, F1: 1}
            _ = v10
            _ = v11
        }
    case "hello":
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v2)
            v13 := GenS2{F0: 42}
            for fi14 := 0; fi14 < 1; fi14++ {
                var v15 int8 = -128
                var v16 uint8 = v5
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    }
    switch (("\n\t" + ((((("" + "\n\t") + "foo bar") + ("foo bar" + "")) + ("" + "hello")) + ("" + "a"))) + (("a" + "a") + "foo bar")) {
    case "":
        switch v2 {
        case 32768:
            var v17 int8 = -128
            _ = v17
        case 760:
            var v18 int8 = -15
            _ = v18
        }
    default:
        switch "hello" {
        case "hello":
            v19 := GenS1{F0: 42, F1: -978}
            _ = v19
        case "":
            var v20 float64 = ((float64(-2.25) + 1.5) / v4)
            _ = v20
        default:
            var v21 uint64 = 2
            _ = v21
        }
    }
    if false {
        v5 -= 1
        if false {
            v4 -= ((0.0 / (v4 + v4)) / v4)
            v22 := GenS2{F0: -32768}
            _ = v22
        }
        fmt.Println(v1)
    }
    v23 := GenS3{F0: (^(v1)), F1: 1, F2: 9223372036854775807}
    v8 -= (^(v8))
    v5 = v5
    v8 -= (v3 ^ v8)
    v1 += v1
    _ = len(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
