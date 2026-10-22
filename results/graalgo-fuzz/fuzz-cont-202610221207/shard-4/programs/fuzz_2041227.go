package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 string }
type GenS2 struct { F0 string; F1 int16; F2 int64 }

func main() {
    var v1 uint = 2
    var v2 bool = (1.5 != (1e10 / 1.5))
    var v3 uint16 = 531
    v4 := []uint32{1, 561}
    fmt.Println(v2)
    v3 = v3
    fmt.Println(v1, v2, v1)
    v2 = v2
    var v5 int = 9223372036854775807
    if len(v4) > 0 {
        v4[0] = 2
    }
    var v6 uint = 2
    var v7 int32 = int32(v5)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int16 = 1
        var v10 uint32 = 223
        _ = fi8
        _ = v10
        _ = v9
    }
    switch "" {
    case "a":
        var v11 float32 = 1.5
        _ = v11
    case "":
        v1 = v1
    default:
        var v12 float32 = 3.14159
        _ = v12
    }
    v13 := GenS1{F0: 0, F1: v6, F2: ""}
    var v14 int = v5
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v15 uint32 = 4294967295
    switch (("a" + ("hello" + "\n\t")) + ("hello" + "hello")) {
    case "foo bar":
        v7 = (-(v7))
    case "\n\t":
        var v16 uint = v6
        _ = v16
    case "":
        v17 := GenS1{F0: 42, F1: (v1 & 224), F2: "hello"}
        _ = v17
    }
    v18 := GenS2{F0: "hello", F1: 739, F2: 1}
    if v2 {
        var v19 int16 = -1
        _ = v19
    } else {
        v7 += (v7 << (((^(v1))) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
