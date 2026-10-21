package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint64 = 42
    var v2 uint = 1
    var v3 uint8 = 127
    var v4 int32 = 1
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]uint32)
    var v7 int8 = 127
    v3 = (v3 >> ((65535) & 63))
    var v8 bool = (65535 >= uint16(v3))
    if v8 {
        for fi9 := 0; fi9 < 9; fi9++ {
            for fi10 := 0; fi10 < 18; fi10++ {
                var v11 int64 = -1
                v7 = v7
                _ = fi10
                _ = v11
            }
            v7 = -1
            var v12 int = -515
            _ = fi9
            _ = v12
        }
    }
    var v13 int64 = -264
    if (true || v8) {
        fmt.Println(v7)
        var v14 string = "hello"
        _ = v14
    } else {
        if (v7 >= 127) {
            v7 += (v7 + v7)
            var v15 float64 = (-2.25 * 0.1)
            var v16 string = "\n\t"
            _ = v15
            _ = v16
        } else {
            if true {
                v8 = v8
                v17 := []int{32768, 1}
                _ = v17
            }
            var v18 float64 = 1.5
            _ = v18
        }
        var v19 uint32 = 2147483648
        _ = v19
    }
    v20 := GenS1{F0: (v2 - (v2 & v2)), F1: v3}
    v2 = (v2 + 840)
    switch "\n\t" {
    case "\n\t":
        var v21 uint32 = 2147483647
        _ = v21
    default:
        v8 = ((v2 / (((v2 & (v2 + 255))) | 1)) == (v2 - (v2 & 0)))
    }
    var v22 int64 = -137
    var v23 bool = v8
    fmt.Println(v13, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
