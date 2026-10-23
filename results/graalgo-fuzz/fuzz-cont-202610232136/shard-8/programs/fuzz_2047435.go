package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 int; F1 uint }
func genFn0(p0 uint, p1 int8, p2 int32) uint8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := []int{(fi6 | 708), fi6}
        _ = fi6
        _ = v7
    }
    var v8 uint = (p0 | p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{"\n\t", "hello", "a", "\n\t"}
    v3 := []uint32{0, 0}
    var v4 int32 = 1
    var v5 int = 2147483647
    _ = len(v2)
    var v9 uint16 = 32767
    var v10 bool = false
    if true {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 uint16 = 1
            v13 := GenS1{F0: v12, F1: 0}
            var v14 int64 = v1
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
    }
    var v15 int16 = 565
    v16 := GenS2{F0: 32767, F1: 436}
    if (false || v10) {
        fmt.Println(v9, v9, v1)
    }
    if (v10 && false) {
        var v17 uint32 = 42
        _ = v17
    } else {
        switch ((("foo bar" + "hello") + "") + "foo bar") {
        case "hello":
            var v18 float64 = 1.5
            _ = v18
        case "\n\t":
            for fi19 := 0; fi19 < 8; fi19++ {
                fmt.Println(v5, fi19)
                fmt.Println(v15, v4)
                _ = fi19
            }
        default:
            v20 := GenS2{F0: int(v4), F1: 646}
            _ = v20
        }
        v5 = (v5 / (((v5 >> ((2) & 63))) | 1))
        if (float64(0.1) > 1.5) {
            fmt.Println(v9, v4, v4)
        }
    }
    var v21 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
