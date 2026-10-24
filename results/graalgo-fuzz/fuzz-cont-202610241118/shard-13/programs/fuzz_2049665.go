package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]bool)
    var v3 int8 = 42
    var v4 uint16 = 32767
    var v5 int = int(v1)
    var v6 string = "hello"
    v7 := GenS1{F0: (!(!true))}
    v4 -= (v4 * 82)
    switch v5 {
    case 1:
        fmt.Println(v4, v6, v1)
    case -2147483648:
        v6 = ((v6 + (v6 + v6)) + (v6 + v6))
    }
    var v8 int8 = 42
    var v9 uint32 = 647
    v1 += int16(v5)
    if (!(721 > 852)) {
        for fi10 := 0; fi10 < 20; fi10++ {
            v11 := GenS2{F0: (v1 * 0), F1: 457}
            _ = fi10
            _ = v11
        }
    }
    v12 := GenS1{F0: ((v3 * (^(v3))) >= (v3 ^ 42))}
    fmt.Println(v5, v4)
    fmt.Println(v3)
    v13 := GenS1{F0: false}
    var v14 bool = false
    var v15 float64 = 3.14159
    v16 := v2[(v9 >> ((285) & 63))]
    v17 := GenS3{F0: 1}
    v4 = (v4 - (^(v4)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
