package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"hello", "a", "a", "a"}
    var v3 int64 = -9223372036854775808
    var v4 uint32 = 2
    v5 := []int{-747, -283, 128}
    v6 := GenS1{F0: 1}
    v3 -= (v3 ^ v3)
    v3 = -477
    var v7 uint32 = (v4 ^ (v4 ^ (v4 / ((4294967295) | 1))))
    var v8 float64 = 0.0
    fmt.Println(v7, v7, v4)
    v3 = (-(v3))
    if true {
        switch 4294967295 {
        case 418:
            v7 -= v7
        case 127:
            var v9 uint32 = (v4 % ((2147483647) | 1))
            _ = v9
        case 32768:
            if true {
                v8 += ((v8 - float64((3.14159 / (1.5 - (((0.1 - (3.14159 - float32(-2.25))) - 1.5) / -2.25))))) - (v8 + v8))
                fmt.Println(v7, v8, v3)
            }
        }
        _ = len(v1)
        for fi10 := 0; fi10 < 19; fi10++ {
            v7 = v4
            _ = fi10
        }
    } else {
        _ = len(v1)
        if false {
            v4 = v7
            v3 += (-(v3))
            for fi11 := 0; fi11 < 6; fi11++ {
                fmt.Println(fi11, v4, fi11)
                _ = fi11
            }
        }
    }
    delete(v1, v4)
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 bool = ((fi12 * fi12) == fi12)
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 495}
    var v15 float32 = ((0.1 / 3.14159) * (float32(3.14159) / ((-2.25 * 3.14159) * (((0.1 * 0.0) * (3.14159 * float32(v8))) / ((0.1 + 3.14159) + 3.14159)))))
    v5 = append(v5, 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
