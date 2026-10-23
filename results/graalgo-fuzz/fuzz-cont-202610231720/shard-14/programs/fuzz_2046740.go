package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint = 255
    v2 := []uint32{42, 42, 486, 849}
    var v3 int64 = 1
    var v4 uint64 = 1
    v5 := make(map[uint32]int)
    var v6 float64 = 3.14159
    _ = len(v5)
    for fi7 := 0; fi7 < 16; fi7++ {
        v6 = (v6 + ((v6 * float64(-2.25)) - v6))
        var v8 int32 = 42
        var v9 int32 = (v8 - -138)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 4294967295
    fmt.Println(v10)
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS1{F0: 812, F1: 92}
        _ = fi11
        _ = v12
    }
    switch "\n\t" {
    case "a":
        var v13 int8 = -128
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 7; fi14++ {
            v3 -= v3
            for fi15 := 0; fi15 < 10; fi15++ {
                v16 := GenS1{F0: v10, F1: 1}
                if (1 != 255) {
                    if (true || (!false)) {
                        v4 += (^(v4))
                        v17 := GenS2{F0: (v10 >> (((v1 & (^(v1)))) & 63))}
                        _ = v17
                    }
                }
                if (!((v4 * 2) <= (v4 + (v4 * v4)))) {
                    fmt.Println(v10)
                    switch ("" + ((("\n\t" + ("foo bar" + "\n\t")) + (("" + "\n\t") + (("" + ("\n\t" + "foo bar")) + (("hello" + "a") + ("" + "\n\t"))))) + (((("" + ("hello" + "hello")) + ("a" + (("hello" + "foo bar") + "hello"))) + ("" + "")) + "foo bar"))) {
                    case "a":
                        v1 += v1
                    case "foo bar":
                        var v18 bool = (false && false)
                        _ = v18
                    default:
                        fmt.Println(fi14)
                    }
                    if true {
                        v10 = 0
                        var v19 float32 = 1e10
                        _ = v19
                    } else {
                        v20 := GenS1{F0: (v10 / ((v10) | 1)), F1: uint8(v1)}
                        v21 := GenS2{F0: 42}
                        _ = v20
                        _ = v21
                    }
                }
                _ = fi15
                _ = v16
            }
            v4 = (v4 * 840)
            _ = fi14
        }
    }
    fmt.Println(v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
