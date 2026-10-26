package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint }
func genFn0(p0 uint32) int {
    p0 -= 112
    switch "hello" {
    case "hello":
        v7 := []int{128, 324, -952}
        _ = v7
    case "foo bar":
        var v8 uint8 = 42
        _ = v8
    default:
        p0 = p0
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            p0 = (p0 >> ((65535) & 63))
            v11 := GenS1{F0: "\n\t", F1: 0, F2: int8(p0)}
            _ = fi10
            _ = v11
        }
        var v12 uint8 = 128
        _ = fi9
        _ = v12
    }
    _ = p0
    return 315
}
func genFn1() uint32 {
    var v13 string = "\n\t"
    v13 = v13
    if true {
        v13 = v13
        fmt.Println(v13)
    }
    _ = v13
    return 65535
}

func main() {
    var v1 float64 = (1.5 / float64(-2.25))
    var v2 int64 = 759
    v3 := make(map[uint32]int)
    var v4 float32 = (float32(v1) - 0.0)
    var v5 float32 = (v4 - v4)
    var v6 uint8 = 191
    v14 := GenS2{F0: (v6 & v6)}
    v15 := []int{32768}
    var v16 string = "foo bar"
    v2 = (-(v2))
    var v17 bool = false
    for fi18 := 0; fi18 < 3; fi18++ {
        v4 -= v4
        var v19 int16 = -757
        if (!(v1 <= (float64(v4) + float64(v4)))) {
            var v20 uint = 2
            for fi21 := 0; fi21 < 11; fi21++ {
                v6 = (^(v6))
                v16 = ("foo bar" + (v16 + v16))
                _ = fi21
            }
            fmt.Println(v20)
            _ = v20
        }
        _ = fi18
        _ = v19
    }
    v15 = append(v15, genFn0(207))
    for fi22 := 0; fi22 < 7; fi22++ {
        switch (fi22 >> ((255) & 63)) {
        case 2147483647:
            v6 += v6
        }
        _ = fi22
    }
    fmt.Println(v16, v1, v1)
    v23 := GenS3{F0: 4294967295, F1: 0}
    v24 := v3[2147483647]
    fmt.Println(v6, v4, v2)
    delete(v3, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
