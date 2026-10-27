package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int8; F1 bool; F2 int16 }
func genFn0(p0 uint64, p1 uint32) string {
    var v5 uint64 = p0
    v6 := []int{2147483647, 128, 2147483647}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "hello"
}
func genFn1(p0 string) int8 {
    v7 := []int{127, 368, -1}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint = 156
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        p0 = p0
        _ = fi10
    }
    _ = p0
    _ = v7
    return -1
}

func main() {
    var v1 int8 = 127
    var v2 int32 = 42
    var v3 float32 = float32(3.14159)
    v4 := make(map[string]bool)
    v2 = (v2 ^ v2)
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint8 = 127
        _ = fi11
        _ = v12
    }
    fmt.Println(v1, v2)
    var v13 uint8 = 1
    switch -9223372036854775808 {
    case 4294967295:
        var v14 float32 = v3
        _ = v14
    case 9223372036854775807:
        if true {
            var v15 uint8 = v13
            if (!(711 <= 42)) {
                v16 := GenS1{F0: "\n\t"}
                _ = v16
            } else {
                v17 := GenS2{F0: int32(v1), F1: v13}
                v18 := []int{-1, 0, -9223372036854775808, 1}
                _ = v17
                _ = v18
            }
            v2 = (v2 << ((65535) & 63))
            _ = v15
        } else {
            v3 = (v3 / float32(-2.25))
        }
    }
    var v19 int32 = 1
    v19 = (-(v2))
    var v20 uint = 1
    v21 := GenS3{F0: (v1 | v1), F1: false, F2: int16(v20)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
