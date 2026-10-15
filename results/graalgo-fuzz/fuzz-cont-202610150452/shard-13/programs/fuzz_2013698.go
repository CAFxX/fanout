package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 int8; F1 string; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS4 struct { F0 uint16; F1 string }

func main() {
    var v1 int8 = 0
    var v2 float64 = 3.14159
    v3 := make(map[string]uint32)
    var v4 uint32 = 427
    var v5 uint64 = 1
    v3[""] = v4
    v6 := v3["hello"]
    var v7 uint8 = 247
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: 0, F1: (32767 <= 1), F2: -263}
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: (v1 - (-(v1))), F1: "", F2: v5}
    var v11 uint = 345
    v1 += (v1 & (v1 & v1))
    var v12 uint64 = (v5 / ((v5) | 1))
    v7 = (v7 << (((v11 - v11)) & 63))
    v6 += (v6 & (v6 % ((988) | 1)))
    for fi13 := 0; fi13 < 18; fi13++ {
        v5 = (v5 / ((v5) | 1))
        v14 := GenS3{F0: uint(v7), F1: v12, F2: 840}
        v6 -= uint32(v11)
        _ = fi13
        _ = v14
    }
    v5 -= (v5 | v5)
    switch -846 {
    case -700:
        v15 := GenS4{F0: 1, F1: (("foo bar" + ((((("a" + ("\n\t" + "foo bar")) + ("\n\t" + "hello")) + "hello") + "\n\t") + "")) + "\n\t")}
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
