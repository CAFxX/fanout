package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint64) int32 {
    p0 = (p0 * (p0 << ((520) & 63)))
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(fi7)
        for fi8 := 0; fi8 < 10; fi8++ {
            p0 = (p0 / (((p0 + p0)) | 1))
            var v9 int16 = 0
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 15; fi10++ {
            var v11 int = (fi7 % ((fi10) | 1))
            v12 := GenS1{F0: true, F1: 9, F2: 1}
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    p0 = (p0 >> ((530) & 63))
    switch 32767 {
    case 127:
        for fi13 := 0; fi13 < 19; fi13++ {
            fmt.Println(p0, fi13)
            fmt.Println(fi13)
            var v14 uint32 = 65535
            _ = fi13
            _ = v14
        }
    default:
        v15 := []int{-265}
        _ = v15
    }
    _ = p0
    return -1
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = 9223372036854775807
    var v3 string = "\n\t"
    var v4 uint8 = 255
    var v5 int32 = -9
    var v6 uint32 = 42
    var v16 bool = ((0 <= (v1 & (v1 >> ((483) & 63)))) || true)
    v17 := GenS2{F0: 127}
    v1 = (v1 & 467)
    fmt.Println(v2, v2)
    for fi18 := 0; fi18 < 3; fi18++ {
        v5 -= (v5 * (v5 & 2147483647))
        _ = fi18
    }
    v6 = (v6 * v6)
    v6 = uint32(v4)
    v6 -= v6
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
