package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 string }

func main() {
    var v1 uint = 255
    v2 := []int{2147483647}
    var v3 int = 42
    var v4 int = (v3 | v3)
    var v5 int16 = 32767
    v6 := GenS1{F0: 909, F1: uint64(v1), F2: 1}
    var v7 int16 = -1
    var v8 uint16 = 0
    var v9 uint = (v1 - v1)
    v10 := GenS2{F0: 18446744073709551615, F1: (-(v4))}
    for fi11 := 0; fi11 < 12; fi11++ {
        switch (v4 ^ int(v9)) {
        case 32768:
            fmt.Println(v7, v8, v9)
        case 4294967295:
            v12 := GenS2{F0: 18446744073709551615, F1: (v3 & (fi11 >> ((18446744073709551615) & 63)))}
            _ = v12
        case -2147483648:
            v8 = v8
        default:
            _ = len(v2)
        }
        var v13 uint64 = 760
        _ = fi11
        _ = v13
    }
    v5 = v5
    switch v3 {
    case -2147483648:
        v9 += (v1 % ((911) | 1))
    case 128:
        v3 = v4
    default:
        v1 = v1
    }
    v5 += (^(v7))
    v14 := GenS3{F0: ("a" + "hello")}
    var v15 uint8 = 42
    var v16 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
