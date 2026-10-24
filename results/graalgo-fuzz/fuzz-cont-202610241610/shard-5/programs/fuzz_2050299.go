package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint16 }
func genFn0(p0 int8, p1 uint32) int64 {
    p1 -= (p1 / (((p1 - p1)) | 1))
    p0 -= p0
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{781, 729, 1, 2}
    var v3 uint8 = 223
    v4 := make(map[string]int)
    v5 := make(map[string]uint32)
    var v6 int16 = -1
    _ = len(v4)
    var v7 uint64 = 2
    v2 = append(v2, 2147483648)
    for fi8 := 0; fi8 < 1; fi8++ {
        switch (("hello" + ("\n\t" + "foo bar")) + ((((("" + ("" + ("hello" + (("hello" + ("a" + "\n\t")) + "\n\t")))) + "") + "foo bar") + "a") + "")) {
        case "\n\t":
            var v9 int32 = 2147483647
            _ = v9
        default:
            var v10 uint16 = uint16(v7)
            _ = v10
        }
        var v11 bool = false
        var v12 int64 = -9223372036854775808
        _ = fi8
        _ = v11
        _ = v12
    }
    fmt.Println(v7)
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: -954, F1: 65535, F2: 0}
        v7 -= v7
        _ = fi13
        _ = v14
    }
    v6 = (v6 - (v6 / ((v6) | 1)))
    v7 -= v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
