package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 bool; F1 int32 }

func main() {
    var v1 int = 32768
    v2 := make(map[int]int)
    var v3 uint32 = 65535
    for fi4 := 0; fi4 < 9; fi4++ {
        v5 := []int{690, v1, v1, (v1 % (((^(v1))) | 1))}
        v6 := GenS1{F0: v3}
        _ = fi4
        _ = v5
        _ = v6
    }
    delete(v2, v1)
    v7 := GenS1{F0: (v3 + 65535)}
    var v8 int = v1
    v9 := GenS2{F0: uint16(v3), F1: -128}
    for fi10 := 0; fi10 < 14; fi10++ {
        var v11 string = ("foo bar" + ("a" + "hello"))
        v12 := GenS3{F0: true, F1: 620}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint64 = 0
    v8 += (v8 << ((18446744073709551615) & 63))
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 float32 = 0.0
        v3 += v3
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
