package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 string = (("hello" + "\n\t") + "")
    var v3 int = 6
    v4 := make(map[uint32]uint32)
    var v5 int = v3
    var v6 bool = true
    var v7 int64 = 0
    v6 = (-2147483648 <= int32(v5))
    var v8 uint8 = 128
    v7 = (v7 / (((v7 & (v7 & 42))) | 1))
    if v6 {
        v7 = v7
        var v9 int32 = int32(v5)
        _ = v9
    } else {
        v2 = "foo bar"
        v7 = v7
        switch (-(v5)) {
        case -70:
            var v10 int = (v5 ^ v5)
            _ = v10
        case 453:
            switch (^(v5)) {
            case -9223372036854775808:
                v1[917] = v2
            case 1:
                if (!(v2 == "")) {
                    v11 := []int{-1}
                    fmt.Println(v8, v5)
                    if v6 {
                        v12 := GenS1{F0: 1}
                        _ = v12
                    } else {
                        var v13 float64 = 1e10
                        _ = len(v11)
                        var v14 int = (v5 >> ((368) & 63))
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                } else {
                    v8 -= (v8 | (^(v8)))
                    if (!(-858 == 1)) {
                        fmt.Println(v7, v3, v8)
                        v6 = v6
                        v15 := GenS2{F0: (v3 - v5), F1: v6}
                        _ = v15
                    }
                    v16 := GenS1{F0: -1}
                    _ = v16
                }
            case -138:
                var v17 int32 = 42
                _ = v17
            }
        case -254:
            var v18 float32 = 1.5
            _ = v18
        default:
            if (1.5 <= -2.25) {
                var v19 uint64 = 18446744073709551615
                var v20 uint16 = 728
                _ = v19
                _ = v20
            }
        }
    }
    var v21 int16 = 668
    v8 -= v8
    v6 = (!v6)
    v5 += -9223372036854775808
    var v22 int16 = -584
    v3 -= v5
    var v23 int32 = 1
    fmt.Println(v8, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
