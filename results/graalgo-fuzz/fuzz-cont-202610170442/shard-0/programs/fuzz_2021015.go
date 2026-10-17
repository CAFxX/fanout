package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    var v1 int8 = -71
    v2 := make(map[string]uint32)
    var v3 uint16 = 2
    var v4 uint16 = 65535
    var v5 int16 = -32768
    var v6 int8 = (v1 | v1)
    v7 := GenS1{F0: (4294967295 != 18446744073709551615), F1: -1}
    delete(v2, "")
    v1 -= 100
    var v8 uint32 = 2147483648
    switch 127 {
    case -1:
        v9 := GenS1{F0: (1.5 <= (float64((float32(-2.25) + 0.1)) / 1.0)), F1: -735}
        _ = v9
    }
    v8 += 42
    delete(v2, ("\n\t" + ""))
    fmt.Println(v5, v8, v1)
    v10 := []int{-2147483648, 403}
    v8 = v8
    v4 = (v3 % (((^(v3))) | 1))
    var v11 int8 = (v1 ^ (v1 | (-(v6))))
    var v12 float64 = 0.1
    var v13 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
