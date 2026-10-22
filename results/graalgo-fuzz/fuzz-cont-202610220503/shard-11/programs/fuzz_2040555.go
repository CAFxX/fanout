package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint16, p2 int16) int16 {
    v5 := GenS1{F0: (p2 / (((p2 + (p2 / ((p2) | 1)))) | 1))}
    switch ("foo bar" + (("" + "a") + "a")) {
    case "a":
        v6 := []int{545}
        _ = v6
    case "foo bar":
        fmt.Println(p0)
    default:
        var v7 uint64 = 18446744073709551615
        _ = v7
    }
    var v8 int32 = int32(p0)
    var v9 int16 = (-(p2))
    var v10 float64 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v8
    _ = v9
    return p2
}

func main() {
    var v1 int16 = -380
    v2 := []uint32{776}
    var v3 int64 = 42
    var v4 uint16 = 42
    v1 = genFn0(255, v4, v1)
    v4 = (v4 >> ((65535) & 63))
    if ((v4 / (((v4 % (((v4 | (v4 / ((v4) | 1)))) | 1))) | 1)) == v4) {
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 10; fi12++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v4, v3)
                v13 := GenS1{F0: v1}
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        v3 += v3
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v14 int32 = -1
    fmt.Println(v14, v4)
    var v15 float32 = (float32(((float64(1e10) + (1.5 * (-2.25 * float64(-2.25)))) * -2.25)) - 0.0)
    v4 -= (v4 >> ((0) & 63))
    v14 = (v14 & (v14 << ((uint(v4)) & 63)))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
