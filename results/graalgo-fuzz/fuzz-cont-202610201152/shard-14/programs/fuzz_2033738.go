package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int64, p1 uint16) int16 {
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(p0)
        _ = fi7
    }
    p0 = (p0 * (p0 + p0))
    for fi8 := 0; fi8 < 13; fi8++ {
        p1 += p1
        var v9 int32 = -2147483648
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 uint8, p1 uint16, p2 uint32) uint32 {
    var v10 uint16 = uint16(p2)
    var v11 int = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (p2 << ((976) & 63))
}

func main() {
    var v1 int8 = -1
    var v2 float64 = 0.0
    v3 := []int{654}
    v4 := make(map[uint32]int)
    var v5 string = "foo bar"
    var v6 int64 = -9223372036854775808
    v1 += (v1 << ((183) & 63))
    var v12 uint = 2
    fmt.Println(v12, v5)
    var v13 uint16 = 0
    v14 := GenS1{F0: v1}
    var v15 int = 32767
    v16 := GenS1{F0: v1}
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            v2 += (v2 * (0.1 * v2))
            _ = fi18
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
