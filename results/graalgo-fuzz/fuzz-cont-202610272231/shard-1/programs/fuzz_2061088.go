package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint8, p1 int32, p2 uint64) uint16 {
    switch 1 {
    case -1:
        var v7 int8 = -128
        _ = v7
    default:
        p0 = p0
    }
    v8 := []int{-565, 128}
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 65535
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 int32 = 464
    var v5 int16 = 0
    var v6 float64 = 0.1
    v9 := GenS1{F0: "hello", F1: 80}
    var v10 string = "foo bar"
    fmt.Println(v6, v6, v5)
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 float64 = 0.0
        _ = fi11
        _ = v12
    }
    fmt.Println(v10, v10)
    v13 := GenS1{F0: ("" + v10), F1: 13}
    fmt.Println(v10, v5)
    fmt.Println(v4, v4)
    v14 := GenS1{F0: v10, F1: -128}
    v2[2147483647] = 127
    var v15 float64 = (float64(((1.5 + -2.25) / float32(-2.25))) / v6)
    if (!((42 == -128) || (float32(0.0) <= (0.0 + -2.25)))) {
        v16 := GenS2{F0: (v4 - (^(v4)))}
        switch (("" + v10) + v10) {
        case "foo bar":
            v6 -= (v6 / 3.14159)
        case "a":
            fmt.Println(v4, v4, v15)
        default:
            var v17 int8 = 127
            _ = v17
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 int8 = int8(fi18)
            v15 = v6
            _ = fi18
            _ = v19
        }
        _ = v16
    }
    var v20 uint = 255
    var v21 bool = true
    var v22 uint8 = 47
    for fi23 := 0; fi23 < 1; fi23++ {
        v24 := []int{(-(fi23)), (fi23 | (fi23 << ((v20) & 63)))}
        var v25 int32 = v4
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println(v22, v5, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
