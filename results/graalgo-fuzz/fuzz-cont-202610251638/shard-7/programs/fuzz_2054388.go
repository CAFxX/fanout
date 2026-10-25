package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int) int64 {
    v6 := []int{(p0 / (((p0 & p0)) | 1)), (^(p0)), 32767}
    p0 = 0
    if (!(2 > 42)) {
        p0 -= p0
        var v7 int = 42
        _ = v7
    }
    _ = p0
    _ = v6
    return -1
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{1, 0, 4294967295, 2147483647}
    var v3 uint = 4294967295
    v4 := []uint32{1}
    v5 := make(map[int]uint32)
    v8 := GenS1{F0: 128}
    v9 := GenS1{F0: 0}
    v10 := GenS1{F0: 127}
    v2 = append(v2, 2147483647)
    v1 = v1
    v1 += (^(v1))
    var v11 int32 = 42
    for fi12 := 0; fi12 < 19; fi12++ {
        if true {
            v13 := GenS1{F0: 0}
            var v14 uint16 = 0
            for fi15 := 0; fi15 < 3; fi15++ {
                fmt.Println(v1, fi12)
                var v16 float32 = 3.14159
                _ = fi15
                _ = v16
            }
            _ = v13
            _ = v14
        }
        v3 -= 42
        _ = fi12
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
