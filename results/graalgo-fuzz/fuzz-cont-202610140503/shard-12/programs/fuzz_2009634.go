package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
func genFn0(p0 uint8, p1 uint32, p2 bool) uint {
    var v5 int64 = -1
    v6 := GenS1{F0: 1, F1: 2}
    var v7 string = (("" + "a") + ("a" + ""))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint8) uint32 {
    p0 = (p0 - p0)
    v8 := []int{-1, 2147483647}
    var v9 int8 = 100
    _ = p0
    _ = v8
    _ = v9
    return 631
}

func main() {
    var v1 float64 = (((float64(0.0) - (0.0 * ((3.14159 * -2.25) * (1.5 + (-2.25 + (-2.25 + (-2.25 / (((-2.25 / 1.0) + 1.5) * 0.1)))))))) + 0.0) / (0.1 / 0.1))
    var v2 bool = false
    v3 := make(map[uint32]string)
    var v4 bool = (2 == 4294967295)
    var v10 bool = v4
    switch -82 {
    case 825:
        v1 += 0.1
    case -243:
        v1 -= (float64((0.0 + (-2.25 - 0.1))) - v1)
    case 1:
        var v11 float32 = 1.5
        _ = v11
    default:
        delete(v3, 0)
    }
    v2 = true
    v4 = ((32767 == 1) || v4)
    switch "\n\t" {
    case "hello":
        var v12 int16 = 1
        _ = v12
    case "foo bar":
        v13 := v3[2]
        _ = v13
    }
    v14 := []int{773, -1, 9223372036854775807}
    var v15 int32 = 0
    var v16 uint32 = 372
    for fi17 := 0; fi17 < 17; fi17++ {
        fmt.Println(v10, v10, fi17)
        v16 = (v16 ^ 0)
        for fi18 := 0; fi18 < 20; fi18++ {
            v4 = v10
            switch (-(fi18)) {
            case 840:
                v19 := GenS1{F0: 997, F1: 4294967295}
                _ = v19
            case 128:
                _ = len(v3)
            }
            var v20 uint32 = v16
            _ = fi18
            _ = v20
        }
        _ = fi17
    }
    var v21 float64 = v1
    v15 = v15
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
