package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 bool, p1 int32) uint16 {
    var v7 float64 = 0.0
    fmt.Println(p0, p0)
    for fi8 := 0; fi8 < 2; fi8++ {
        switch fi8 {
        case 127:
            for fi9 := 0; fi9 < 20; fi9++ {
                v10 := GenS1{F0: -127}
                _ = fi9
                _ = v10
            }
        case -486:
            var v11 uint16 = 898
            _ = v11
        }
        switch "a" {
        case "":
            v12 := []int{fi8, fi8, (fi8 >> ((65535) & 63))}
            _ = v12
        case "a":
            v13 := GenS1{F0: int8(fi8)}
            _ = v13
        }
        fmt.Println(p1, fi8)
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 255
}

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 2
    v3 := []int{127, 330}
    v4 := make(map[uint32]bool)
    var v5 uint = 1
    var v6 bool = (!true)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 uint32 = uint32(v2)
        v5 += uint(v2)
        _ = fi14
        _ = v15
    }
    var v16 bool = (v6 || v6)
    v2 = (v2 + v2)
    var v17 uint = (v5 * (v5 / ((18446744073709551615) | 1)))
    for fi18 := 0; fi18 < 9; fi18++ {
        fmt.Println(fi18, v6, v5)
        v5 = v17
        _ = fi18
    }
    fmt.Println(v16, v6)
    fmt.Println(v17)
    v2 = (v2 | (v2 / ((v2) | 1)))
    var v19 int32 = 1
    v5 = (v17 << ((v17) & 63))
    for fi20 := 0; fi20 < 9; fi20++ {
        var v21 uint32 = 4294967295
        var v22 bool = (!v16)
        var v23 uint8 = (v2 & (^(v2)))
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    if len(v3) > 0 {
        v3[0] = 163
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
