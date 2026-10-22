package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 int16; F1 bool; F2 uint8 }
func genFn0(p0 int32, p1 int16, p2 int64) int32 {
    v5 := GenS1{F0: 831, F1: 0, F2: "foo bar"}
    switch "foo bar" {
    case "\n\t":
        v6 := GenS2{F0: 514, F1: true, F2: 127}
        _ = v6
    case "a":
        v7 := GenS1{F0: 529, F1: 2, F2: "hello"}
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p0 >> ((18446744073709551615) & 63))
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]string)
    var v3 int8 = 100
    var v4 uint16 = 255
    v1 = (uint32(v4) >= 47)
    v8 := v2[65535]
    v1 = (-706 <= 42)
    v9 := GenS2{F0: 1000, F1: v1, F2: 127}
    v1 = v1
    if true {
        var v10 int16 = -1
        var v11 uint8 = 0
        _ = v10
        _ = v11
    } else {
        v4 += v4
    }
    v12 := GenS1{F0: 2, F1: 2, F2: "foo bar"}
    if (!((v1 && (!false)) && (!v1))) {
        var v13 int = 4294967295
        _ = v13
    } else {
        fmt.Println(v8)
        fmt.Println(v3, v1)
        var v14 float64 = 0.1
        _ = v14
    }
    fmt.Println(v3, v1, v8)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
