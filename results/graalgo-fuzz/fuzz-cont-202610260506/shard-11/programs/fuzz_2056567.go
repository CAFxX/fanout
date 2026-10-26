package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int }
type GenS2 struct { F0 uint; F1 int8; F2 uint64 }
type GenS3 struct { F0 int8; F1 uint16; F2 int }
func genFn0(p0 int16, p1 uint8) string {
    v4 := GenS1{F0: 18446744073709551615, F1: 400, F2: 32767}
    p1 -= 2
    for fi5 := 0; fi5 < 14; fi5++ {
        if false {
            fmt.Println(p1)
            var v6 uint8 = 255
            p1 = p1
            _ = v6
        }
        _ = fi5
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = v4
    return ("a" + "a")
}
func genFn1(p0 uint16, p1 int32) int {
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint8 = 92
            if (true || ((100 != 42) && false)) {
                v10 := []int{(fi7 >> ((1) & 63)), (fi7 / (((-(fi7))) | 1)), fi7, (fi8 << ((18446744073709551615) & 63))}
                v11 := GenS2{F0: 0, F1: -90, F2: 1}
                var v12 string = "foo bar"
                _ = v10
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v13 uint32 = 1
    v13 = (v13 >> ((346) & 63))
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v13
    return 1
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]bool)
    var v3 uint32 = 167
    v14 := GenS3{F0: 7, F1: 1, F2: 127}
    var v15 float64 = v1
    v2[0] = (0 >= 65535)
    v16 := []int{-371, genFn1(601, 1)}
    var v17 bool = true
    fmt.Println(v1, v1)
    var v18 uint64 = 4294967295
    fmt.Println(v1, v17, v1)
    if (v3 < (v3 * 2)) {
        var v19 int16 = 1000
        if true {
            var v20 int16 = v19
            var v21 float64 = ((-2.25 + -2.25) + 0.0)
            _ = v20
            _ = v21
        }
        _ = v19
    }
    var v22 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
}
