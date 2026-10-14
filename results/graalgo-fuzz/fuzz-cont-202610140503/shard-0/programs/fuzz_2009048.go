package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint, p1 uint32) uint {
    p0 = 65535
    v7 := []int{2147483647}
    p0 = (^(p0))
    _ = len(v7)
    _ = p0
    _ = p1
    _ = v7
    return (p0 - p0)
}

func main() {
    v1 := []int{506, 9223372036854775807, 9223372036854775807}
    var v2 float32 = (1e10 / -2.25)
    v3 := make(map[int]string)
    v4 := []uint32{843, 2147483648, 315, 4294967295}
    v5 := make(map[uint32]uint32)
    v6 := []uint32{929, 2, 49}
    if false {
        if (!true) {
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 string = "\n\t"
                var v10 int = fi8
                var v11 uint = 18446744073709551615
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        }
        v2 = (0.0 - ((v2 + (0.1 - v2)) / (1.5 * (v2 * v2))))
        v2 -= v2
    }
    v2 -= float32(float64(1e10))
    var v12 uint = 2
    var v13 uint32 = 907
    v12 -= (v12 >> ((v12) & 63))
    if false {
        v12 += (^(v12))
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS1{F0: "a"}
            v2 -= v2
            if true {
                var v16 int32 = 2147483647
                for fi17 := 0; fi17 < 20; fi17++ {
                    v2 = 1.5
                    v16 += (v16 % ((-2147483648) | 1))
                    _ = fi17
                }
                _ = v16
            }
            _ = fi14
            _ = v15
        }
    } else {
        fmt.Println(v12)
        v18 := GenS1{F0: ("a" + "foo bar")}
        _ = v18
    }
    var v19 int8 = 5
    var v20 int8 = (-(v19))
    v12 -= (v12 & v12)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
