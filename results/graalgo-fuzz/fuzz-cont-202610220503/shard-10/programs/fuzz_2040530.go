package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int16 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint; F1 uint }

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 42
    var v3 int16 = 1000
    var v4 uint16 = 32767
    v5 := []int{128, 128, -970, -887}
    var v6 uint32 = uint32(v4)
    v2 += v2
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v1, v2, v3)
    v7 := GenS1{F0: true, F1: 128, F2: (-(v3))}
    v8 := GenS2{F0: -165, F1: 761}
    var v9 bool = false
    for fi10 := 0; fi10 < 8; fi10++ {
        fmt.Println(v1, v3)
        v11 := GenS3{F0: 648, F1: 65535}
        v12 := GenS1{F0: (-9223372036854775808 != -212), F1: fi10, F2: -1}
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(v4)
    var v13 int64 = 42
    v14 := GenS1{F0: v9, F1: 1, F2: int16(v6)}
    var v15 int64 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
