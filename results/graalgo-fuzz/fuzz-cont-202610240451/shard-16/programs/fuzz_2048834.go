package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 219
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    var v5 uint64 = 18446744073709551615
    fmt.Println(v2, v5)
    var v6 uint64 = (v5 - v5)
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := []int{fi7, 0, (-(fi7))}
        _ = fi7
        _ = v8
    }
    v5 = 159
    v9 := []int{783, 24, 862, -2147483648}
    v6 = 42
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint16 = 255
        if false {
            for fi12 := 0; fi12 < 14; fi12++ {
                v13 := GenS1{F0: 2147483647, F1: (v6 | (v5 + v6))}
                var v14 uint8 = 42
                if len(v9) > 0 {
                    v9[0] = fi12
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = len(v1)
            v15 := GenS1{F0: 2147483647, F1: 4294967295}
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    var v16 float32 = -2.25
    for fi17 := 0; fi17 < 17; fi17++ {
        for fi18 := 0; fi18 < 15; fi18++ {
            v2 += 0
            _ = fi18
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
