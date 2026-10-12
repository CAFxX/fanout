package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 int32; F1 bool; F2 int }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 uint32 = 332
    fmt.Println(v3)
    switch 32767 {
    case -74:
        if (!true) {
            v3 += 1
            switch 9223372036854775807 {
            case -839:
                v3 = 0
            case 32768:
                var v4 uint64 = 4294967295
                _ = v4
            default:
                switch -658 {
                case -763:
                    var v5 uint32 = v3
                    _ = v5
                case 992:
                    v3 += v3
                }
            }
        }
    default:
        var v6 uint32 = (v3 << ((1) & 63))
        _ = v6
    }
    v7 := GenS1{F0: 4294967295, F1: 528, F2: -141}
    v8 := GenS2{F0: 42, F1: true, F2: 42}
    var v9 float32 = (0.1 + float32(3.14159))
    var v10 uint16 = 1
    var v11 uint = 1
    v11 = v11
    v12 := v1[1]
    v2[v12] = ""
    v13 := GenS3{F0: v12, F1: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
