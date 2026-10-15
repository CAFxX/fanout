package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int) int64 {
    var v5 int64 = 42
    v5 += v5
    _ = p0
    _ = v5
    return v5
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 443
    var v3 uint32 = 4294967295
    var v4 uint16 = 65535
    v6 := GenS1{F0: true}
    v7 := GenS1{F0: true}
    v8 := GenS2{F0: 65535}
    var v9 uint16 = (v4 + (v4 * (v4 * (v4 >> (((v2 / (((^(v2))) | 1))) & 63)))))
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            v12 := []int{(fi10 + (fi10 >> (((v2 - v2)) & 63))), (fi11 * fi10), fi11, fi11}
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        v4 -= (^(v4))
        _ = fi13
    }
    v14 := GenS1{F0: false}
    v15 := GenS1{F0: (true || (1 >= 1))}
    fmt.Println(v3, v9, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
