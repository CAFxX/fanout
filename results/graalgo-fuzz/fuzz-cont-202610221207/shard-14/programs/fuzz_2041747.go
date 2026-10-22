package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 bool; F2 int8 }
func genFn0(p0 int64, p1 uint64) int8 {
    var v6 int32 = 42
    if (!(((0.1 < 1e10) && (!(!true))) || true)) {
        v7 := GenS1{F0: 255, F1: (false || (2 <= 42))}
        v6 -= (-(v6))
        v8 := []int{32768, 9223372036854775807, 0, -2147483648}
        _ = v7
        _ = v8
    }
    v6 += 2147483647
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := []uint32{0, 42}
    var v2 int8 = -78
    var v3 float64 = 1.5
    var v4 int64 = 1
    var v5 uint32 = 4294967295
    var v9 uint16 = 255
    for fi10 := 0; fi10 < 4; fi10++ {
        fmt.Println(v9)
        var v11 int16 = 0
        v4 -= v4
        _ = fi10
        _ = v11
    }
    fmt.Println(v4)
    var v12 int32 = 42
    fmt.Println(v9, v2, v5)
    fmt.Println(v2, v12, v2)
    v5 = (v5 << ((255) & 63))
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 bool = false
        if true {
            for fi15 := 0; fi15 < 9; fi15++ {
                switch fi15 {
                case 422:
                    v4 = v4
                }
                if true {
                    if v14 {
                        v5 = (^(v5))
                        v5 = 2147483647
                    }
                }
                _ = fi15
            }
            v5 = (^(v5))
            var v16 int32 = (v12 ^ v12)
            _ = v16
        }
        if (!((((float32(v3) - float32(v3)) - 0.0) >= (0.1 * 1e10)) || (v9 == (v9 ^ v9)))) {
            switch "hello" {
            case "foo bar":
                v3 -= ((0.1 / v3) * 0.1)
            case "hello":
                if false {
                    var v17 uint = 42
                    fmt.Println(v2, v14, v9)
                    v1 = append(v1, 65535)
                    _ = v17
                }
            default:
                for fi18 := 0; fi18 < 4; fi18++ {
                    v19 := GenS1{F0: 156, F1: false}
                    _ = fi18
                    _ = v19
                }
            }
            v4 -= 42
            switch (-(fi13)) {
            case 42:
                v20 := GenS2{F0: true}
                _ = v20
            case 128:
                var v21 uint8 = 42
                _ = v21
            case 32768:
                v9 += (v9 % ((v9) | 1))
            default:
                v9 += v9
            }
        }
        _ = fi13
        _ = v14
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            v2 = (v2 + v2)
            _ = fi23
        }
        _ = fi22
    }
    for fi24 := 0; fi24 < 3; fi24++ {
        var v25 bool = ((fi24 % (((-(fi24))) | 1)) == fi24)
        fmt.Println(v4, v2, v5)
        _ = fi24
        _ = v25
    }
    v26 := GenS3{F0: 241, F1: ((v3 - v3) == v3), F2: v2}
    var v27 float64 = (-2.25 / float64((1.5 * 1.5)))
    if (!false) {
        fmt.Println(v5, v27, v4)
        for fi28 := 0; fi28 < 10; fi28++ {
            fmt.Println(v2, v5)
            _ = fi28
        }
        var v29 uint8 = uint8(v9)
        _ = v29
    }
    var v30 int8 = genFn0(v4, 130)
    v3 = v3
    switch 2147483647 {
    case -989:
        v3 = (((0.0 / v3) + ((v27 / -2.25) * -2.25)) / v3)
    case -9223372036854775808:
        if false {
            var v31 bool = true
            var v32 string = "hello"
            v32 = v32
            _ = v31
            _ = v32
        } else {
            v27 += 1.5
            _ = len(v1)
            var v33 uint16 = 255
            _ = v33
        }
    default:
        if len(v1) > 0 {
            v1[0] = 215
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
