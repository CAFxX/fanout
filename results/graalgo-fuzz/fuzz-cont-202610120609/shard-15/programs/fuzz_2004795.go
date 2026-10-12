package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int }
func genFn0(p0 int16, p1 uint8, p2 uint16) uint64 {
    var v7 uint = 255
    v7 = (v7 / (((^(v7))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 float32 = (0.1 * (float32((-2.25 / 1.0)) + -2.25))
    v2 := []string{("\n\t" + "a"), ("a" + "hello"), "\n\t"}
    var v3 int16 = -765
    v4 := []int{}
    var v5 uint = 106
    var v6 uint8 = 102
    v8 := GenS1{F0: v6, F1: 4294967295, F2: 9223372036854775807}
    fmt.Println(v5, v3)
    fmt.Println(v3)
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint16 = uint16(v5)
        for fi11 := 0; fi11 < 16; fi11++ {
            v2 = append(v2, "\n\t")
            var v12 uint16 = (v10 << (((v5 % (((v5 << (((v5 & v5)) & 63))) | 1))) & 63))
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    v3 += 1
    switch ("\n\t" + ("\n\t" + "hello")) {
    case "foo bar":
        v5 += (v5 % ((255) | 1))
    case "a":
        var v13 string = "\n\t"
        _ = v13
    case "\n\t":
        var v14 float32 = 1.5
        _ = v14
    default:
        v6 += (v6 - 1)
    }
    var v15 uint8 = 128
    fmt.Println(v15)
    fmt.Println(v1, v15)
    v15 = uint8(v5)
    var v16 float32 = 3.14159
    fmt.Println(v5, v1, v16)
    v17 := GenS1{F0: (v15 << ((v5) & 63)), F1: 4294967295, F2: -680}
    _ = len(v2)
    fmt.Println(v15, v5, v1)
    v1 = float32(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
