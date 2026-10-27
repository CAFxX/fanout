package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint; F1 int16 }
func genFn0(p0 int64) int8 {
    v7 := []int{9223372036854775807, 128, 4294967295, -676}
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(p0, p0)
        var v9 int64 = -246
        _ = fi8
        _ = v9
    }
    p0 += 42
    if len(v7) > 0 {
        v7[0] = -9223372036854775808
    }
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 bool = (65535 < 0)
    v2 := make(map[uint32]bool)
    var v3 float32 = 3.14159
    v4 := make(map[string]string)
    var v5 uint64 = 18446744073709551615
    var v6 uint8 = 127
    v10 := GenS1{F0: 785}
    var v11 int = -9223372036854775808
    var v12 float32 = v3
    if (true && (!v1)) {
        v13 := v2[589]
        v14 := GenS2{F0: v11, F1: 2}
        v11 += v11
        _ = v13
        _ = v14
    }
    var v15 float32 = v12
    var v16 uint64 = (v5 | v5)
    _ = len(v4)
    var v17 int16 = 1000
    var v18 bool = (v16 < v16)
    v19 := GenS3{F0: 2, F1: 65535}
    v3 = (0.0 + v15)
    switch "\n\t" {
    case "\n\t":
        switch 629 {
        case 127:
            fmt.Println(v15)
        }
    default:
        var v20 float32 = (1e10 / 1.0)
        _ = v20
    }
    if v1 {
        fmt.Println(v12, v16, v6)
        v3 -= (0.1 * float32(1.5))
    } else {
        v21 := GenS4{F0: 728, F1: (v17 - (v17 & (v17 >> ((57) & 63))))}
        switch "hello" {
        case "a":
            v5 -= (v16 | (v16 | (v16 - (v5 | (v5 >> ((255) & 63))))))
        default:
            fmt.Println(v12, v5)
        }
        _ = v21
    }
    var v22 uint = 0
    if false {
        if v18 {
            if (32767 >= v17) {
                v15 += 1e10
                fmt.Println(v16, v15)
            }
        } else {
            var v23 int64 = 0
            v22 = v22
            _ = v23
        }
        v24 := GenS2{F0: v11, F1: 1}
        if (39 <= 34) {
            v17 = (v17 << (((v22 + 910)) & 63))
        } else {
            v25 := GenS2{F0: (-(v11)), F1: 459}
            _ = v25
        }
        _ = v24
    }
    var v26 int32 = -50
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
