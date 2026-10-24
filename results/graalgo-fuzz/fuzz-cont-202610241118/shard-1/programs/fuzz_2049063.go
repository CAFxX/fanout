package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint32 = 2147483647
    v2 := []string{}
    v3 := make(map[string]string)
    v4 := []uint32{v1, (v1 * (v1 | 1)), (v1 & 624)}
    v5 := []int{-845, 672}
    v6 := []string{"\n\t", "\n\t"}
    v1 = v1
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            if true {
                delete(v3, "hello")
                fmt.Println(fi8)
                var v9 float32 = (-2.25 / float32(1.5))
                _ = v9
            }
            _ = fi8
        }
        v1 += (v1 / (((v1 % ((v1) | 1))) | 1))
        _ = fi7
    }
    v1 = v1
    var v10 int16 = 42
    var v11 uint16 = uint16(v1)
    for fi12 := 0; fi12 < 9; fi12++ {
        v1 = v1
        v11 = (^(v11))
        _ = fi12
    }
    v1 += 2147483647
    var v13 uint64 = 546
    v4 = append(v4, v1)
    v1 += (v1 & v1)
    v13 = v13
    _ = len(v2)
    v14 := GenS1{F0: false}
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 4; fi16++ {
            v10 += (v10 >> ((985) & 63))
            _ = fi16
        }
        if (-2.25 > (0.0 / 1.5)) {
            fmt.Println(v1)
            var v17 int64 = -825
            var v18 int8 = -8
            _ = v17
            _ = v18
        }
        fmt.Println(v11, v10, v10)
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
