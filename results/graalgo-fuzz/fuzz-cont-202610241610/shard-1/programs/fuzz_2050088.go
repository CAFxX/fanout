package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 float32 = (1.5 / 0.1)
    var v3 uint = 42
    var v4 int16 = 1
    v5 := []uint32{894, 2147483648, 870, 2147483647}
    var v6 int8 = 114
    v2 += (v2 / v2)
    v2 = (v2 - v2)
    var v7 uint = (v3 + v3)
    switch 127 {
    case 0:
        for fi8 := 0; fi8 < 5; fi8++ {
            v2 = float32(0.1)
            _ = fi8
        }
    case 32768:
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: (v6 >> ((v3) & 63))}
            _ = fi9
            _ = v10
        }
    case 61:
        v4 += v4
    default:
        _ = len(v5)
    }
    var v11 int32 = 0
    var v12 uint16 = uint16(v3)
    var v13 int16 = 42
    fmt.Println(v12, v7, v12)
    v14 := GenS1{F0: v6}
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 int32 = v11
        v17 := v1[int(v11)]
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 bool = false
    switch -9223372036854775808 {
    case 128:
        var v19 int32 = v11
        _ = v19
    default:
        for fi20 := 0; fi20 < 6; fi20++ {
            v6 = v6
            v6 -= 127
            _ = fi20
        }
    }
    if v18 {
        if len(v5) > 0 {
            v5[0] = 115
        }
    } else {
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint16 = v12
            var v23 int8 = (-(v6))
            _ = fi21
            _ = v22
            _ = v23
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
