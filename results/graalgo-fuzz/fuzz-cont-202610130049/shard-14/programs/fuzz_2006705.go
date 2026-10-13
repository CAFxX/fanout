package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int16; F1 int16 }
func genFn0() uint {
    var v7 string = "a"
    var v8 int16 = 42
    var v9 int16 = (v8 + (v8 & (v8 | (v8 << ((65535) & 63)))))
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    v1 := []string{("" + ("a" + "\n\t")), ("hello" + "a"), "\n\t"}
    v2 := []uint32{}
    var v3 uint16 = 991
    v4 := []uint32{42, 42, 753, 42}
    v5 := make(map[int]uint32)
    v6 := make(map[int]bool)
    if (-1 >= 1) {
        switch "\n\t" {
        case "\n\t":
            var v10 bool = true
            _ = v10
        case "":
            fmt.Println(v3)
        case "foo bar":
            v11 := GenS1{F0: -989, F1: 2}
            _ = v11
        }
        v12 := GenS1{F0: 0, F1: 859}
        _ = v12
    } else {
        var v13 int16 = -932
        v14 := GenS2{F0: -880, F1: 0}
        _ = v13
        _ = v14
    }
    var v15 float64 = (1e10 * 0.1)
    if false {
        var v16 uint64 = 247
        var v17 float32 = 0.0
        _ = v16
        _ = v17
    }
    var v18 int8 = 100
    fmt.Println(v3, v18)
    for fi19 := 0; fi19 < 4; fi19++ {
        v3 = (v3 - v3)
        _ = fi19
    }
    if (false || true) {
        var v20 uint8 = 63
        switch (("a" + ("\n\t" + "\n\t")) + "") {
        case "a":
            v18 -= (v18 / (((^(v18))) | 1))
        default:
            var v21 float32 = (((0.1 + (0.0 / 0.1)) * 3.14159) + (-2.25 / (3.14159 / 1.5)))
            _ = v21
        }
        _ = v20
    }
    v15 = v15
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
