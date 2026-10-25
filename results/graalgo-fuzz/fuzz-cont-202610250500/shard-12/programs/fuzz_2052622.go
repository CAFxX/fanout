package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint16 }
type GenS2 struct { F0 bool; F1 uint64; F2 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 133
    v3 := make(map[int]string)
    v4 := []uint32{2147483647, 906, 65535}
    var v5 uint = 0
    var v6 float64 = 1e10
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := GenS1{F0: 18446744073709551615, F1: ("foo bar" + ("" + (((("a" + "") + "foo bar") + "") + ""))), F2: 32767}
        _ = fi7
        _ = v8
    }
    v6 = (((v6 - (float64(0.1) + v6)) + v6) * 3.14159)
    v2 += v2
    _ = len(v4)
    v6 += (v6 - -2.25)
    v4 = append(v4, 219)
    v9 := GenS2{F0: (!false), F1: 4294967295, F2: (v5 + (v5 ^ (v5 >> ((255) & 63))))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
