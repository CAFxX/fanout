package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0() int16 {
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 int16 = -32768
            v10 := []int{fi7, (fi8 + (fi8 - fi7)), (-(fi7)), (fi7 >> ((4294967295) & 63))}
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 uint64 = 4294967295
        v11 -= 18446744073709551615
        _ = fi7
        _ = v11
    }
    v12 := []int{180, 42, 42, 32767}
    v13 := GenS1{F0: 31}
    _ = v12
    _ = v13
    return -32768
}
func genFn1(p0 string, p1 uint8, p2 bool) uint8 {
    v14 := GenS2{F0: 255}
    v15 := GenS2{F0: 2}
    var v16 float32 = 3.14159
    fmt.Println(p0)
    var v17 int = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return (^(p1))
}

func main() {
    var v1 bool = (false || false)
    v2 := []int{-1, 128}
    v3 := []uint32{2147483647, 1, 4294967295, 2147483647}
    v4 := make(map[uint32]int)
    var v5 float32 = (0.0 - ((0.0 * 1.5) * (0.0 / 1.0)))
    var v6 string = "hello"
    _ = len(v4)
    v5 += (float32(((1.5 + 3.14159) - -2.25)) / float32(-2.25))
    var v18 int8 = -1
    v6 = (v6 + v6)
    v3 = append(v3, 333)
    v18 -= (v18 & v18)
    var v19 string = "hello"
    fmt.Println(v18)
    if v1 {
        var v20 uint16 = 0
        _ = v20
    } else {
        for fi21 := 0; fi21 < 2; fi21++ {
            v1 = (42 <= -1)
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS1{F0: 128}
                var v24 uint32 = 0
                var v25 float32 = v5
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi21
        }
        for fi26 := 0; fi26 < 18; fi26++ {
            v6 = ((v19 + v19) + v6)
            v18 = 100
            for fi27 := 0; fi27 < 12; fi27++ {
                if false {
                    var v28 int8 = (v18 >> ((1) & 63))
                    _ = v28
                }
                _ = fi27
            }
            _ = fi26
        }
    }
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 int16 = 1
        fmt.Println(v1, v1, v18)
        _ = fi29
        _ = v30
    }
    v4[4294967295] = 128
    var v31 uint = 544
    v18 -= (v18 & (-(v18)))
    v18 -= -1
    v19 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
