package main

import (
    "fmt"
)

func genFn0(p0 uint32) int32 {
    v7 := []int{2147483647, -1, 32767, -9223372036854775808}
    v7 = append(v7, 9223372036854775807)
    if len(v7) > 0 {
        v7[0] = 1
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        p0 += (p0 >> ((162) & 63))
        _ = fi8
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 uint32, p1 uint16) uint32 {
    p0 -= p0
    v9 := []int{128, 4294967295, 128}
    p1 += 1
    _ = p0
    _ = p1
    _ = v9
    return p0
}

func main() {
    var v1 uint16 = 0
    var v2 int32 = -1
    var v3 uint64 = 336
    v4 := make(map[uint32]bool)
    var v5 uint32 = 65535
    v6 := make(map[int]uint32)
    v2 = v2
    for fi10 := 0; fi10 < 1; fi10++ {
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    if (127 > 188) {
        v12 := []int{-2147483648, -536, 254, 127}
        _ = v12
    }
    var v13 int16 = 0
    v13 = (v13 << ((580) & 63))
    var v14 bool = (!(true || false))
    switch 489 {
    case -803:
        var v15 int16 = v13
        _ = v15
    }
    var v16 bool = (false && v14)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
