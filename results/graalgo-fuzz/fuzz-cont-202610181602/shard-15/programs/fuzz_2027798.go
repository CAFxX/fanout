package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 bool }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint16; F1 int16 }

func main() {
    var v1 int = 32767
    v2 := make(map[int]int)
    var v3 uint = 521
    v4 := make(map[uint32]bool)
    var v5 uint8 = 128
    v1 = v1
    var v6 string = "foo bar"
    v7 := []int{(v1 & 474), -2147483648}
    v3 = (v3 << (((v3 & 606)) & 63))
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: "", F1: -1, F2: true}
        _ = fi8
        _ = v9
    }
    v7 = append(v7, v1)
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint64 = 872
        fmt.Println(v3)
        v6 = v6
        _ = fi10
        _ = v11
    }
    v1 = v1
    v12 := GenS2{F0: (-(v1)), F1: v6}
    var v13 uint32 = 2
    if len(v7) > 0 {
        v7[0] = v1
    }
    switch (v6 + ((((v6 + v6) + "a") + (v6 + "")) + "")) {
    case "foo bar":
        if ((!true) || (((float64((1e10 + 0.1)) == (-2.25 / 1.5)) || (false && true)) && true)) {
            v5 += (v5 << ((232) & 63))
        }
    default:
        _ = len(v4)
    }
    v5 -= v5
    var v14 int32 = -2147483648
    v15 := GenS3{F0: 0, F1: 108}
    var v16 int32 = (v14 | v14)
    v16 = 994
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
