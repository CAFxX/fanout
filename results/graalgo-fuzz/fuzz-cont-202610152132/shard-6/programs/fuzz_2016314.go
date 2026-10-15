package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16, p1 int32, p2 int16) uint16 {
    p1 = (p1 << ((42) & 63))
    p1 -= (p1 | p1)
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 int32 = 196
    var v2 string = "foo bar"
    v3 := []int{}
    v4 := make(map[string]uint32)
    var v5 uint64 = 42
    if len(v3) > 0 {
        v3[0] = 127
    }
    v6 := GenS1{F0: 503, F1: 0, F2: 42}
    v7 := GenS1{F0: 0, F1: 2147483647, F2: 192}
    delete(v4, (("" + "foo bar") + "a"))
    var v8 float64 = 3.14159
    if true {
        fmt.Println(v2, v1)
    }
    var v9 float32 = (1.5 + 1e10)
    v2 = "a"
    fmt.Println(v8, v9, v8)
    for fi10 := 0; fi10 < 16; fi10++ {
        v4[v2] = 2147483648
        if ((1 < 0) || true) {
            v9 -= ((float32((v8 + float64((((1.5 / v9) / 1.0) + (-2.25 / v9))))) / float32(float64(1e10))) + v9)
            v11 := GenS2{F0: -66}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 11; fi12++ {
                v3 = append(v3, int(v5))
                v2 = v2
                _ = fi12
            }
            v1 += 2147483647
        }
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
