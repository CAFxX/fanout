package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 uint64, p2 uint8) uint64 {
    p2 -= (p2 - 0)
    var v4 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p1
}
func genFn1(p0 string, p1 int, p2 int) string {
    p0 = (p0 + ("foo bar" + p0))
    v5 := GenS1{F0: -138}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p0
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]bool)
    var v3 int = 128
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS2{F0: 1}
        var v8 uint8 = 2
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{}
    var v10 float64 = 3.14159
    v3 += v3
    var v11 int8 = 0
    var v12 bool = (818 < 2)
    switch v3 {
    case 4294967295:
        for fi13 := 0; fi13 < 12; fi13++ {
            if v12 {
                var v14 int = 504
                _ = v14
            } else {
                v15 := GenS1{F0: fi13}
                _ = v15
            }
            v16 := GenS1{F0: (fi13 * fi13)}
            _ = fi13
            _ = v16
        }
    case 601:
        v12 = false
    case 323:
        var v17 float32 = 3.14159
        _ = v17
    default:
        for fi18 := 0; fi18 < 20; fi18++ {
            v12 = ((v1 + (v1 + v1)) >= ((((("" + v1) + (v1 + v1)) + v1) + v1) + v1))
            _ = fi18
        }
    }
    var v19 uint64 = 0
    if v12 {
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int64 = -9223372036854775808
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS1{F0: 32767}
    v23 := GenS2{F0: 32767}
    delete(v2, (v3 | -563))
    v11 += v11
    v12 = v12
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v9
}
