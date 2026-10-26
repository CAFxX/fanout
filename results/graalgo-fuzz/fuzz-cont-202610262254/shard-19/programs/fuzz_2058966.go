package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }

func main() {
    var v1 float32 = (3.14159 + 0.1)
    var v2 uint = 333
    var v3 int = 32768
    v4 := make(map[uint32]bool)
    var v5 bool = (0 >= -1)
    v1 = 1.5
    var v6 int64 = 42
    v7 := GenS1{F0: 63, F1: 42}
    if (!(!false)) {
        v6 = (v6 >> (((v2 + 865)) & 63))
    } else {
        for fi8 := 0; fi8 < 4; fi8++ {
            v3 -= 32767
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 uint8 = 255
                v4[810] = (int8(v2) != -1)
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    }
    v3 += -168
    var v11 string = ""
    v12 := GenS1{F0: 255, F1: 2147483647}
    fmt.Println(v3)
    var v13 int16 = 56
    v4[65535] = (42 >= 42)
    var v14 bool = v5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
