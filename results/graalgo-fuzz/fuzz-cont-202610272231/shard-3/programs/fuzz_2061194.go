package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int64; F2 bool }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]bool)
    var v3 uint = 65535
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    v6 := make(map[uint32]int)
    var v7 uint64 = 123
    var v8 int16 = 32767
    if (21 <= -128) {
        if false {
            var v9 float32 = (1.5 - ((0.1 * (float32(3.14159) - (float32((((0.0 * (0.1 + 3.14159)) + 0.1) * 0.1)) - float32(1.5)))) - 1e10))
            _ = v9
        }
    } else {
        switch 42 {
        case -9223372036854775808:
            v7 -= (v7 >> (((v3 + (v3 ^ (v3 * v3)))) & 63))
        case 32768:
            v10 := GenS1{F0: ((v8 % (((v8 - v8)) | 1)) > v8)}
            _ = v10
        }
    }
    var v11 int = 229
    var v12 uint8 = 174
    v12 = 127
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS2{F0: 42, F1: 42, F2: true}
        delete(v5, (fi13 + (fi13 >> ((2) & 63))))
        v15 := GenS1{F0: ((true && false) || false)}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
