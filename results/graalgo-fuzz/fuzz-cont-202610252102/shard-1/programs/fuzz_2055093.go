package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 32767
    var v3 float32 = 3.14159
    var v4 float64 = 0.1
    var v5 uint = 883
    var v6 int64 = -1
    fmt.Println(v4, v6, v3)
    v6 -= (v6 * (v6 - (v6 ^ 9223372036854775807)))
    v4 -= float64(v3)
    var v7 bool = (-32768 != -32768)
    switch 32767 {
    case -9223372036854775808:
        v8 := []int{4294967295, 2147483647, 32767}
        _ = v8
    case 128:
        v9 := GenS1{F0: 446, F1: 0}
        _ = v9
    }
    v2 -= (v2 ^ 2)
    var v10 uint8 = 11
    var v11 uint32 = 0
    v3 = v3
    if (!false) {
        v7 = true
        if true {
            var v12 int8 = 31
            if (-1 < 9223372036854775807) {
                var v13 int64 = (v6 - int64(v12))
                var v14 uint32 = 2147483647
                switch 385 {
                case 0:
                    var v15 uint16 = 302
                    _ = v15
                case 128:
                    v13 -= (v13 >> ((306) & 63))
                }
                _ = v13
                _ = v14
            }
            _ = v12
        }
    }
    var v16 uint32 = v11
    var v17 uint8 = v10
    _ = len(v1)
    if (v7 && true) {
        v6 = (v6 - v6)
        var v18 string = "hello"
        v19 := GenS1{F0: (v11 - 762), F1: -66}
        _ = v18
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
