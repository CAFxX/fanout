package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int; F1 uint32; F2 string }

func main() {
    var v1 bool = true
    v2 := []uint32{1, 42}
    var v3 uint16 = 255
    var v4 int = -9223372036854775808
    var v5 float32 = 0.1
    v6 := make(map[string]int)
    var v7 int32 = 0
    switch "\n\t" {
    case "foo bar":
        var v8 uint32 = 2147483647
        _ = v8
    case "hello":
        _ = len(v6)
    default:
        v1 = (v1 || (!v1))
    }
    v2 = append(v2, uint32(v3))
    switch "hello" {
    case "hello":
        var v9 float64 = (0.1 / 1.0)
        _ = v9
    }
    var v10 int = (v4 & (v4 ^ v4))
    v7 -= 0
    v10 -= (v10 * 2147483647)
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 uint16 = v3
        switch v10 {
        case 32768:
            v7 = 0
        case -9223372036854775808:
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 int = 373
                if ((false || ((42 <= 42) && (true || false))) || (127 > 100)) {
                    var v15 uint16 = v3
                    if v1 {
                        v16 := GenS1{F0: 219, F1: 42}
                        if v1 {
                            v17 := GenS1{F0: 108, F1: 255}
                            _ = v17
                        }
                        var v18 int16 = -1
                        _ = v16
                        _ = v18
                    }
                    _ = v15
                }
                var v19 float32 = v5
                var v20 uint16 = (v12 & (v3 + 0))
                v1 = (false || (v3 == v3))
                v12 = (v20 & v12)
                if v1 {
                    var v21 float32 = (float32((-2.25 / -2.25)) + 1e10)
                    switch (("" + "\n\t") + (("a" + "a") + "a")) {
                    case "\n\t":
                        if (true || v1) {
                            v20 = v12
                            v22 := GenS2{F0: v14, F1: 0, F2: ("hello" + (((("hello" + "") + "\n\t") + ("a" + "")) + "a"))}
                            _ = v22
                        }
                    }
                    _ = v21
                }
                _ = fi13
                _ = v14
                _ = v19
                _ = v20
            }
        default:
            for fi23 := 0; fi23 < 19; fi23++ {
                v2 = append(v2, 426)
                _ = fi23
            }
        }
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
