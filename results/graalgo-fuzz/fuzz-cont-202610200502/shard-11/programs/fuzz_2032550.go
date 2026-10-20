package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int8, p2 uint8) uint32 {
    if true {
        v4 := []int{4294967295, 340, 9223372036854775807, 4294967295}
        v4 = append(v4, -207)
        if (true || true) {
            var v5 uint16 = 0
            _ = v5
        }
        _ = v4
    }
    p0 = (p0 >> ((1) & 63))
    v6 := GenS1{F0: 65535, F1: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 65535
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{4294967295, 0}
    var v3 bool = false
    if (100 >= -1) {
        var v7 uint = 255
        v3 = (2147483647 != uint32(v7))
        _ = v7
    }
    v8 := GenS2{F0: 0, F1: (genFn0(392, 127, 2) <= 42)}
    v9 := GenS2{F0: -32768, F1: v3}
    if len(v2) > 0 {
        v2[0] = 65535
    }
    v10 := GenS3{F0: 308}
    v11 := GenS2{F0: -1, F1: ((!v3) || v3)}
    delete(v1, "a")
    if v3 {
        switch 4294967295 {
        case 127:
            for fi12 := 0; fi12 < 14; fi12++ {
                var v13 uint16 = 2
                var v14 int16 = 1
                _ = fi12
                _ = v13
                _ = v14
            }
        default:
            fmt.Println(v3)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
