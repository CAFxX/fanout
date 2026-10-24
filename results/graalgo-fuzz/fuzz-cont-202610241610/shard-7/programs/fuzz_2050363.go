package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
func genFn0(p0 int64) uint {
    p0 -= 42
    v5 := GenS1{F0: 255, F1: 42}
    v6 := []int{127, 4294967295}
    if ((-(p0)) <= p0) {
        var v7 int32 = -241
        switch "" {
        case "hello":
            var v8 int64 = (p0 << ((0) & 63))
            _ = v8
        default:
            var v9 int64 = (p0 & p0)
            _ = v9
        }
        _ = v7
    } else {
        p0 -= 1
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    _ = p0
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 int32 = 42
    var v2 uint32 = 0
    v3 := make(map[uint32]bool)
    v4 := []string{"", "a"}
    var v10 string = "\n\t"
    v10 = v10
    if (32767 < 0) {
        v11 := GenS1{F0: 1, F1: v1}
        _ = v11
    }
    _ = len(v3)
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 bool = (!(1 != 2))
        _ = fi12
        _ = v13
    }
    var v14 uint32 = v2
    v1 = 42
    var v15 float32 = 1.5
    var v16 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
