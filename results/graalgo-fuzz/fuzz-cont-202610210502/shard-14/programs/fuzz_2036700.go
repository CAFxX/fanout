package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int8, p1 uint8) int32 {
    switch 42 {
    case 9223372036854775807:
        p0 -= (p0 % ((p0) | 1))
    }
    p1 -= (^(p1))
    switch -555 {
    case 1:
        fmt.Println(p0, p0)
    default:
        switch 127 {
        case 177:
            fmt.Println(p1, p0)
        default:
            if (32768 <= -277) {
                for fi6 := 0; fi6 < 16; fi6++ {
                    v7 := []int{}
                    _ = fi6
                    _ = v7
                }
                switch ("foo bar" + "a") {
                case "":
                    p0 = (p0 ^ p0)
                default:
                    p1 += 2
                }
            } else {
                var v8 float32 = (-2.25 / 1.5)
                var v9 int = 0
                var v10 int8 = p0
                _ = v10
                _ = v8
                _ = v9
            }
        }
    }
    p1 -= 42
    var v11 uint16 = 42
    _ = p0
    _ = p1
    _ = v11
    return -1
}

func main() {
    var v1 float64 = (float64(1.5) / 1e10)
    var v2 int16 = 800
    var v3 int16 = v2
    v4 := make(map[uint32]int)
    var v5 uint = 18446744073709551615
    v12 := GenS1{F0: 366}
    _ = len(v4)
    var v13 bool = (2 != 1)
    var v14 uint64 = 0
    v15 := []int{32768, 9223372036854775807, 2147483647, 128}
    var v16 uint64 = (v14 + v14)
    switch 780 {
    case 1:
        fmt.Println(v14, v13, v3)
    case 4294967295:
        v4[4294967295] = 0
    case 128:
        var v17 float32 = 0.1
        _ = v17
    }
    v18 := GenS1{F0: 1}
    v16 = (v14 / ((v16) | 1))
    v19 := GenS1{F0: 255}
    v20 := GenS2{F0: 223}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
