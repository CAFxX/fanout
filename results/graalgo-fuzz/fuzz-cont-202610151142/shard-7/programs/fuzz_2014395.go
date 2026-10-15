package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int16, p2 string) uint16 {
    v7 := GenS1{F0: 676}
    switch -2147483648 {
    case 798:
        var v8 int64 = 9223372036854775807
        _ = v8
    case -2147483648:
        p1 = p1
    case -9223372036854775808:
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 uint64 = 905
    v2 := make(map[int]uint32)
    var v3 int32 = int32(v1)
    var v4 float32 = float32(float64(3.14159))
    v5 := []uint32{1}
    var v6 float32 = v4
    v6 = (float32(1.5) + v4)
    var v9 uint64 = (v1 / ((2) | 1))
    var v10 uint8 = 255
    if true {
        if ((v1 << ((2) & 63)) == v1) {
            v11 := GenS1{F0: 847}
            v12 := v2[4294967295]
            _ = len(v2)
            _ = v11
            _ = v12
        } else {
            v13 := GenS1{F0: 42}
            v1 = v9
            _ = v13
        }
    }
    var v14 uint = 255
    v15 := GenS1{F0: genFn0((v1 / ((v9) | 1)), 478, "a")}
    v16 := GenS1{F0: 0}
    v10 = (v10 % ((119) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
