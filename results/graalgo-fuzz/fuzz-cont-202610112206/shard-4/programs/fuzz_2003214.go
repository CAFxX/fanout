package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int8; F2 uint16 }
func genFn0(p0 int16) uint64 {
    p0 = (p0 ^ 0)
    v4 := []int{637}
    _ = p0
    _ = v4
    return 18446744073709551615
}
func genFn1(p0 int, p1 int, p2 uint8) uint32 {
    p1 = (p0 % (((-(p0))) | 1))
    p1 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 127
    var v3 int32 = 713
    var v5 string = ""
    for fi6 := 0; fi6 < 20; fi6++ {
        v5 = v5
        fmt.Println(v2)
        v5 = (v5 + ((v5 + v5) + v5))
        _ = fi6
    }
    if false {
        var v7 uint = 196
        v8 := GenS1{F0: v5}
        var v9 float32 = ((0.0 + 1e10) * 0.0)
        _ = v7
        _ = v8
        _ = v9
    } else {
        v10 := []int{-854, -460, -200, -603}
        for fi11 := 0; fi11 < 11; fi11++ {
            v1[830] = (fi11 - (fi11 << ((255) & 63)))
            v2 = 127
            _ = fi11
        }
        _ = v10
    }
    var v12 uint8 = 247
    fmt.Println(v12)
    v13 := GenS1{F0: "a"}
    switch "a" {
    case "hello":
        var v14 uint = 0
        _ = v14
    case "\n\t":
        for fi15 := 0; fi15 < 18; fi15++ {
            if true {
                var v16 string = ((v5 + "foo bar") + "hello")
                _ = v16
            } else {
                fmt.Println(v12, fi15, v5)
                var v17 float32 = float32(0.0)
                _ = v17
            }
            switch v5 {
            case "hello":
                v18 := GenS2{F0: 4294967295, F1: int8(v12), F2: 42}
                _ = v18
            case "":
                v12 += (v12 * v12)
            case "foo bar":
                fmt.Println(v5, v3, v3)
            default:
                switch (("a" + "foo bar") + v5) {
                case "hello":
                    var v19 float32 = (1.5 + 0.0)
                    _ = v19
                case "a":
                    if false {
                        var v20 int64 = 384
                        delete(v1, 981)
                        v5 = "foo bar"
                        _ = v20
                    }
                default:
                    fmt.Println(fi15, v12)
                }
            }
            fmt.Println(fi15)
            _ = fi15
        }
    default:
        for fi21 := 0; fi21 < 12; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                switch 127 {
                case 1:
                    var v23 int64 = -1
                    _ = v23
                case 4294967295:
                    v1[1] = 9223372036854775807
                case 127:
                    var v24 int8 = 1
                    _ = v24
                default:
                    fmt.Println(fi21)
                }
                v12 = (v12 ^ v12)
                _ = fi22
            }
            _ = fi21
        }
    }
    v12 -= 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
}
